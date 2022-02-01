package main

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"

	"emwell/internal/api/telegram"
	"emwell/internal/api/telegram/consumer"
	"emwell/internal/api/telegram/handlers"
	"emwell/internal/api/telegram/handlers/daily_routine"
	"emwell/internal/api/telegram/handlers/start"
	"emwell/internal/api/telegram/handlers/unknown"
	"emwell/internal/api/telegram/middlewares"
	"emwell/internal/api/telegram/middlewares/register"
	"emwell/internal/api/telegram/sender"
	"emwell/internal/config"
	"emwell/internal/logger"
	"emwell/internal/user"
	"emwell/internal/user/repository/psql"
)

func initTelegramBot(wg *sync.WaitGroup) (*telegram.Telegram, error) {
	c, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", c.PostgreSQL.MasterDSN)
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

	conn, err := consumer.NewEventConsumer(log, c.Telegram.Token, wg)
	if err != nil {
		return nil, err
	}

	replySender, err := sender.NewReplySender(log, c.Telegram.Token)
	if err != nil {
		return nil, err
	}

	tg, err := telegram.NewTelegramBotAPI(
		log,
		conn,
		[]middlewares.Middleware{
			register.NewMiddleware(log, user.NewManager(log, psql.NewRepository(db))),
		},
		[]handlers.Handler{
			start.NewMenuHandler(),
			daily_routine.NewDailyRoutineHandler(),
			unknown.NewHandler(),
		},
		replySender,
	)
	if err != nil {
		return nil, err
	}

	return tg, nil
}
