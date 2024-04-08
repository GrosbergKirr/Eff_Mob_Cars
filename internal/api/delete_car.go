package api

import (
	"Eff_Mob/models"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func NewDeleter(log *slog.Logger, del Cars) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req models.RegNumCar

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Debug("fail to decode json", http.StatusBadRequest)
		}

		err = del.DeleteCar(log, req.RegNum)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	}
}
