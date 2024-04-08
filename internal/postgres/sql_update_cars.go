package postgres

import (
	"log/slog"
)

func (s *Storage) UpdateCar(log *slog.Logger, query string) error {
	updateStmt, err := s.Db.Prepare(query)
	if err != nil {
		log.Warn("change exec mistake%s", err)
		return err
	}

	updateExec, err := updateStmt.Exec()
	if err != nil {
		log.Warn("change exec mistake%s", err)
		return err
	}

	_ = updateExec
	return nil
}
