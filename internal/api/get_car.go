package api

import (
	"Eff_Mob/models"
	"Eff_Mob/tools"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func NewCarGetter(log *slog.Logger, car Cars) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req models.Car

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Debug("fail to decode json", http.StatusBadRequest)
		}

		log.Info("Get request")
		page := r.URL.Query().Get("page")

		// Set query for DB
		query := tools.QueryFilterSelect(req, page)

		// Get info from DB
		c, err := car.GetCarInfo(log, query)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			GetCarRespOK(w, r, c)
		}
		log.Info("Get Response")

	}
}

// Response func

func GetCarRespOK(w http.ResponseWriter, r *http.Request, carInfo []models.CarInfo) {
	render.JSON(w, r, carInfo)
}
