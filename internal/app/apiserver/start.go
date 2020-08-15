package apiserver

import (
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/shitikovkirill/auth-service/internal/app/config"
	"github.com/shitikovkirill/auth-service/internal/app/dependencies/db"
	"github.com/shitikovkirill/auth-service/internal/app/dependencies/logger"
	"github.com/shitikovkirill/auth-service/internal/app/dependencies/redis"
	"github.com/shitikovkirill/auth-service/internal/app/service/realservice"
)

// Start server ...
func Start(config *config.Config) error {
	// Load additional dependencies
	logs := logger.Get()

	if len(config.SentryDSN) > 0 {
		if err := sentry.Init(sentry.ClientOptions{Dsn: config.SentryDSN}); err != nil {
			logs.Fatalf("sentry.Init: %s", err)
		}
		logs.Info("Sentry running")
		sentry.CaptureMessage("Online-payments running...")
		defer sentry.Flush(2 * time.Second)
	}

	if err := redis.Load(config.Redis); err != nil {
		return err
	}

	store, err := db.GetStore(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer store.Close()

	srv := newServer(realservice.New(store))

	logs.Infof("Server running on address %v", config.BindAddr)
	return http.ListenAndServe(config.BindAddr, srv)
}
