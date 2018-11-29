package db

import (
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type PostresOpt struct {
	ConnStr  string
	MaxConns int
}

func Postgres(o *PostresOpt) *Service {
	conn, err := dbr.Open("postgres", o.ConnStr, nil)
	if err != nil {
		log.Panic().Msgf("unable to obtain db connection: %+v", o.ConnStr)
	}
	conn.SetMaxIdleConns(o.MaxConns)
	conn.SetMaxOpenConns(o.MaxConns)
	return &Service{conn}
}
