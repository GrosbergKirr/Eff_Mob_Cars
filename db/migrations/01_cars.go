package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up, Down)
}

func Up(tx *sql.Tx) error {
	_, err := tx.Exec("create table if not exists car(regNum varchar primary key ,mark varchar, model varchar,year integer, owner integer);")
	if err != nil {
		return err
	}
	return nil
}

func Down(tx *sql.Tx) error {
	_, err := tx.Exec("create table if not exists people(id varchar primary key, name varchar , surname varchar, patronymic varchar);")
	if err != nil {
		return err
	}
	return nil
}
