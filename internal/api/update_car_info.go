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
			log.Warn("fail to decode json", http.StatusBadRequest)
		}
		query := tools.QueryFilterUpdateCar(w, req)

		err = car.UpdateCar(log, query)
		if err != nil {
			log.Warn("failed to update car info,", err)
		}

	}
}
