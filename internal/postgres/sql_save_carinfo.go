package postgres

import (
	"Eff_Mob/models"
	"log/slog"
)

func (s *Storage) SaveCar(log *slog.Logger, infos []models.CarInfo) error {
	for i := range infos {

		saveStmt, err := s.Db.Prepare("INSERT into car VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING")

		if err != nil {
			log.Warn("insert prepare mistake\n%s", err)
			return err
		}
		saveExec, err := saveStmt.Exec(infos[i].RegNum, infos[i].Mark, infos[i].Model, infos[i].Year)

		if err != nil {
			log.Warn("insert exec mistake%s", err)
			return err
		}

		_ = saveExec

		saveOwnerStmt, err := s.Db.Prepare("INSERT into people VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING")

		if err != nil {
			log.Warn("insert prepare mistake\n%s", err)
			return err
		}
		saveOwnerExec, err := saveOwnerStmt.Exec(infos[i].CurrentOwner.Id, infos[i].CurrentOwner.Name,
			infos[i].CurrentOwner.Surname, infos[i].CurrentOwner.Patronymic)

		if err != nil {
			log.Warn("insert exec mistake%s", err)
			return err
		}

		_ = saveOwnerExec

	}
	return nil
}
