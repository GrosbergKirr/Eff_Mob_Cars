package api

import (
	"Eff_Mob/models"
	"Eff_Mob/tools"
	"fmt"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func NewCarGetter(log *slog.Logger, car Cars) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req models.Car

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Warn("fail to decode json", http.StatusBadRequest)
		}

		log.Info("Get request:")
		page := r.URL.Query().Get("page")

		query := tools.QueryFilterSelect(req, page)
		fmt.Println(query)

		c, err := car.GetCarInfo(log, query)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			GetCarRespOK(w, r, c)
		}

	}
}

func GetCarRespOK(w http.ResponseWriter, r *http.Request, carInfo []models.CarInfo) {
	render.JSON(w, r, carInfo)
}
