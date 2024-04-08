package main

import (
	"Eff_Mob/internal/api"
	"Eff_Mob/internal/config"
	"Eff_Mob/internal/logger"
	"Eff_Mob/internal/postgres"
	"Eff_Mob/internal/server"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"os"
)

func main() {
	//Init config and logger

	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	log.Info("init config & logger success")

	// Init storage
	storage, err := postgres.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage. If you didn't set DB password, do it in local.yaml", err)
		os.Exit(1)
	}
	_ = storage

	log.Info("init storage success")

	// Set router
	router := chi.NewRouter()

	//Set handel's paths
	router.Get("/show", api.NewCarGetter(log, storage))
	router.Delete("/delete", api.NewDeleter(log, storage))
	router.Post("/save", api.NewCarSaver(log, storage))
	router.Post("/update", api.NewCarUpdater(log, storage))

	log.Info("starting server", slog.String("address", cfg.Address))

	//Run server
	server.ServerRun(log, cfg, router)

}
