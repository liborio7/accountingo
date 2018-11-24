package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type PostresOpt struct {
	ConnStr  string
	MaxConns int
}

func Postgres(o *PostresOpt) *sql.DB {
	db, err := sql.Open("postgres", o.ConnStr)
	if err != nil {
		log.Panicf("unable to obtain db connection: %v", o.ConnStr)
	}
	db.SetMaxIdleConns(o.MaxConns)
	db.SetMaxOpenConns(o.MaxConns)
	return db
}
