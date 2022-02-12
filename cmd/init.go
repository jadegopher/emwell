package main

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"

	"emwell/internal/api/telegram"
	"emwell/internal/api/telegram/consumer"
	"emwell/internal/api/telegram/handlers"
	"emwell/internal/api/telegram/handlers/daily_routine"
	"emwell/internal/api/telegram/handlers/daily_routine/rates"
	"emwell/internal/api/telegram/handlers/start"
	"emwell/internal/api/telegram/handlers/unknown"
	"emwell/internal/api/telegram/middlewares"
	"emwell/internal/api/telegram/middlewares/register"
	"emwell/internal/api/telegram/sender"
	"emwell/internal/config"
	"emwell/internal/core/diary"
	"emwell/internal/core/diary/entites"
	diaryRepo "emwell/internal/core/diary/repository/psql"
	"emwell/internal/core/user"
	"emwell/internal/core/user/repository/psql"
	"emwell/internal/logger"
)

type container struct {
	services services
}

type services struct {
	diary *diary.Diary
}

func (c *container) InitTelegramBot(wg *sync.WaitGroup) (*telegram.Telegram, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", cfg.PostgreSQL.MasterDSN)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	conn, err := consumer.NewEventConsumer(log, cfg.Telegram.Token, wg)
	if err != nil {
		return nil, err
	}

	replySender, err := sender.NewReplySender(log, cfg.Telegram.Token)
	if err != nil {
		return nil, err
	}

	c.services.diary = diary.NewDiary(log, diaryRepo.NewRepository(db))

	tg, err := telegram.NewTelegramBotAPI(
		log,
		conn,
		[]middlewares.Middleware{
			register.NewMiddleware(log, user.NewManager(log, psql.NewRepository(db))),
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
		unknown.NewHandler(),
	}
}
