package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

type Storage struct {
	Db *sql.DB
}

func New(storagePath string) (*Storage, error) {

	db, err := sql.Open("postgres", storagePath)

	if err != nil {
		return nil, err
	}

	//cr, err := db.Exec("create table if not exists car(regNum varchar primary key , mark varchar, model varchar, year integer, owner integer)")
	//cr1, err := db.Exec("create table if not exists people(id varchar primary key, name varchar , surname varchar, patronymic varchar)")
	err = goose.Up(db, "db/migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to up migration %e", err)
	}
	return &Storage{Db: db}, nil
}
