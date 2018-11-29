package db

import (
	"context"
	"database/sql"
	"github.com/gocraft/dbr"
	"github.com/rs/zerolog/log"
)

type Service struct {
	*dbr.Connection
}

func (s *Service) OpenSession() *dbr.Session {
	return s.NewSession(&TracingEventReceiver{})
}

func (s *Service) OpenTransaction(sess *dbr.Session, opts *sql.TxOptions) (*dbr.Tx, error) {
	tx, err := sess.BeginTx(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

type TracingEventReceiver struct {
	dbr.NullEventReceiver
}

func (er *TracingEventReceiver) SpanStart(ctx context.Context, eventName, query string) context.Context {
	log.Ctx(ctx).Info().Msgf("DB :: %s", query)
	return ctx
}

func (er *TracingEventReceiver) SpanError(ctx context.Context, err error) {
	log.Ctx(ctx).Error().Msg(err.Error())
}

func (er *TracingEventReceiver) SpanFinish(ctx context.Context) {}
