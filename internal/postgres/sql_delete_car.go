package postgres

import (
	"log/slog"
)

func (s *Storage) DeleteCar(log *slog.Logger, regNum string) error {
	delStmt, err := s.Db.Prepare("delete from car where regnum = $1")

	if err != nil {
		log.Warn("delete prepare mistake\n%s", err)
		return err
	}

	delExec, err := delStmt.Exec(regNum)

	if err != nil {
		log.Warn("delete exec mistake%s", err)
		return err
	}
	_ = delExec
	return nil
}
