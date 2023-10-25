package main

import (
	"errors"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/wesleycremonini/go-event-based-ddd/internal/event"
	"github.com/wesleycremonini/go-event-based-ddd/internal/psql"
	"github.com/wesleycremonini/go-event-based-ddd/internal/workforce"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	err := run()
	if err != nil {
		zap.L().Fatal(err.Error(), zap.ByteString("debug_stack", debug.Stack()))
	}
}

func run() error {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	var cfg config
	cfg.db.dsn = os.Getenv("DB_DSN")

	logger, err := newLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()

	db, err := psql.New(cfg.db.dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	app := &application{
		config:       cfg,
		middlewares:  Middlewares{users: psql.NewUserRepository()},
		eventService: event.NewOrderEventService(psql.NewCompanyRepository(), psql.NewCustomerRepository()),
		eventQueue:   workforce.NewDispatcher(10, 120),
	}

	zap.L().Info("server started")
	err = app.server().ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	zap.L().Info("server stopped")
	event.WG.Wait()
	return nil
}

func newLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncoderConfig.FunctionKey = "func"
	config.OutputPaths = []string{"stderr"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.EncoderConfig.LevelKey = zapcore.OmitKey
	config.EncoderConfig.LineEnding = "\n\n"
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(logger)

	return logger, nil
}

type config struct {
	db struct {
		dsn string
	}
}

type application struct {
	config       config
	middlewares  Middlewares
	eventService workforce.Service
	eventQueue   workforce.Queue
}
