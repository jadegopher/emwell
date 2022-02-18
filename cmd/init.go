package main

import (
	"database/sql"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	httpHandlers "emwell/internal/api/http/handlers"
	"emwell/internal/api/telegram"
	"emwell/internal/api/telegram/consumer"
	"emwell/internal/api/telegram/handlers"
	"emwell/internal/api/telegram/handlers/daily_routine"
	"emwell/internal/api/telegram/handlers/daily_routine/rates"
	"emwell/internal/api/telegram/handlers/start"
	"emwell/internal/api/telegram/handlers/statistic"
	"emwell/internal/api/telegram/handlers/statistic/getter"
	"emwell/internal/api/telegram/handlers/unknown"
	"emwell/internal/api/telegram/middlewares"
	"emwell/internal/api/telegram/middlewares/register"
	"emwell/internal/api/telegram/sender"
	"emwell/internal/config"
	"emwell/internal/core/diary"
	"emwell/internal/core/diary/entites"
	diaryRepo "emwell/internal/core/diary/repository/psql"
	"emwell/internal/core/link"
	"emwell/internal/core/link/repository/kv"
	"emwell/internal/core/timer"
	"emwell/internal/core/user"
	"emwell/internal/core/user/repository/psql"
	"emwell/internal/logger"
)

type container struct {
	services services
	config   config.Config
}

type services struct {
	logger *logger.Logger
	diary  *diary.Diary
	link   *link.Service
}

func (c *container) InitTelegramBot(wg *sync.WaitGroup) (_ *telegram.Telegram, err error) {
	c.config, err = config.GetConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", c.config.PostgreSQL.MasterDSN)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	c.services.logger, err = logger.NewLogger()
	if err != nil {
		return nil, err
	}

	conn, err := consumer.NewEventConsumer(c.services.logger, c.config.Telegram.Token, wg)
	if err != nil {
		return nil, err
	}

	replySender, err := sender.NewReplySender(c.services.logger, c.config.Telegram.Token)
	if err != nil {
		return nil, err
	}

	c.services.diary = diary.NewDiary(c.services.logger, diaryRepo.NewRepository(db))
	c.services.link = link.NewLinkService(c.services.logger, c.config.EmWell.URL, &timer.Timer{}, kv.NewLinkKVStorage())

	tg, err := telegram.NewTelegramBotAPI(
		c.services.logger,
		conn,
		[]middlewares.Middleware{
			register.NewMiddleware(c.services.logger, user.NewManager(c.services.logger, psql.NewRepository(db))),
		},
		c.initHandlers(),
		replySender,
	)
	if err != nil {
		return nil, err
	}

	return tg, nil
}

func (c *container) initHandlers() []handlers.Handler {
	return []handlers.Handler{
		start.NewMenuHandler(),
		daily_routine.NewDailyRoutineHandler(),
		rates.NewDailyRoutineEmotionalHandler(
			daily_routine.DailyRoutineWorst,
			entites.WorstEmotionalRate,
			"Какой ужас...",
			c.services.diary,
		),
		rates.NewDailyRoutineEmotionalHandler(
			daily_routine.DailyRoutineWorse,
			entites.WorseEmotionalRate,
			"Эх(",
			c.services.diary,
		),
		rates.NewDailyRoutineEmotionalHandler(
			daily_routine.DailyRoutineBad,
			entites.BadEmotionalRate,
			"Печально(",
			c.services.diary,
		),
		rates.NewDailyRoutineEmotionalHandler(
			daily_routine.DailyRoutineNeutral,
			entites.NeutralEmotionalRate,
			"Нуу у всех бывает такое",
			c.services.diary,
		),
		rates.NewDailyRoutineEmotionalHandler(
			daily_routine.DailyRoutineGood,
			entites.GoodEmotionalRate,
			"Отлично, рад за тебя)",
			c.services.diary,
		),
		rates.NewDailyRoutineEmotionalHandler(
			daily_routine.DailyRoutineBetter,
			entites.BetterEmotionalRate,
			"Ура! Так держать!",
			c.services.diary,
		),
		rates.NewDailyRoutineEmotionalHandler(
			daily_routine.DailyRoutineBest,
			entites.BestEmotionalRate,
			"Превосходно)",
			c.services.diary,
		),
		statistic.NewStatisticHandler(),
		getter.NewStatisticGetterHandler(
			c.services.diary,
			c.services.link,
			statistic.EmotionalStatisticsForWeek,
			c.config.EmWell.URL,
			0, 0, -7,
		),
		getter.NewStatisticGetterHandler(
			c.services.diary,
			c.services.link,
			statistic.EmotionalStatisticsForMonth,
			c.config.EmWell.URL,
			0, -1, 0,
		),
		unknown.NewHandler(),
	}
}

func (c *container) InitHttpHandlers() http.Handler {
	r := mux.NewRouter()
	h := httpHandlers.NewHandler(c.services.logger, c.services.link)
	r.HandleFunc("/emotional/statistics", h.GetEmotionalStatistics)
	return r
}
