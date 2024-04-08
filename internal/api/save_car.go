package api

import (
	"Eff_Mob/internal/api/client"
	"Eff_Mob/models"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func NewCarSaver(log *slog.Logger, car Cars) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req models.SaveNums

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Debug("fail to decode json", http.StatusBadRequest)
		}
		resp := client.Client(log, req.RegNums)

		err = car.SaveCar(log, resp)
		if err != nil {
			log.Debug("fail to save to DB: ", http.StatusBadRequest)
		}
	}
}
