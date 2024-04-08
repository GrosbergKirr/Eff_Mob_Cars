package server

import (
	"Eff_Mob/internal/config"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func ServerRun(log *slog.Logger, cfg *config.Config, router chi.Router) {
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Error("failed to start server", err)
	} else {
		log.Info("server started")
	}
}
