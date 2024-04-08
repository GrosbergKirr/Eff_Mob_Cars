package api

import (
	"Eff_Mob/models"
	"log/slog"
)

type Cars interface {
	GetCarInfo(log *slog.Logger, request string) ([]models.CarInfo, error)
	DeleteCar(log *slog.Logger, regNum string) error
	SaveCar(log *slog.Logger, infos []models.CarInfo) error
	UpdateCar(log *slog.Logger, query string) error
}
