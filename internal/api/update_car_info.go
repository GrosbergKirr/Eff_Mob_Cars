package api

import (
	"Eff_Mob/models"
	"Eff_Mob/tools"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func NewCarUpdater(log *slog.Logger, car Cars) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req models.CarInfo

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Debug("fail to decode json", http.StatusBadRequest)
		}

		// Set query for DB
		query := tools.QueryFilterUpdateCar(w, req)

		// Update info in DB
		err = car.UpdateCar(log, query)
		if err != nil {
			log.Debug("failed to update car info,", err)
		}

	}
}
