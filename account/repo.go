package account

import (
	"context"
	"errors"
	"github.com/liborio7/accountingo/cache"
	"github.com/liborio7/accountingo/db"
	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"
	"time"
)

type Repo struct {
	db *db.Service
	ch *cache.Service
}

func NewRepo(db *db.Service, ch *cache.Service) *Repo {
	return &Repo{db: db, ch: ch}
}

func (r *Repo) Insert(ctx context.Context, model *Model) error {
	if uuid.Equal(uuid.Nil, model.Id) {
		model.Id = uuid.NewV4()
	}
	now := time.Now()
	model.InsertDate = now
	model.InsertMillis = now.UnixNano() / 1e6
	stm := r.db.OpenSession().InsertInto("account").
		Columns("id", "version", "insert_date", "insert_millis", "name", "surname", "currency", "amount").
		Record(model)
	result, err := stm.ExecContext(ctx)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("no rows affected")
	}
	log.Ctx(ctx).Info().Msgf("inserted: %+v", model)
	return nil
}

func (r *Repo) LoadById(ctx context.Context, dest *Model, id *uuid.UUID) error {
	stm := r.db.OpenSession().
		Select("*").
		From("account").
		Where("id = ?", *id)
	if err := stm.LoadOneContext(ctx, dest); err != nil {
		return err
	}
	log.Ctx(ctx).Info().Msgf("loaded: %+v", dest)
	return nil
}

func (r *Repo) Load(ctx context.Context, dest *[]Model, startingAfter *uint64, limit *uint64) error {
	stm := r.db.OpenSession().
		Select("*").
		From("account")
	if startingAfter != nil {
		stm = stm.Where("insert_millis > ?", startingAfter)
	}
	if limit != nil {
		stm = stm.Limit(*limit)
	} else {
		stm = stm.Limit(20)
	}
	n, err := stm.LoadContext(ctx, dest)
	if err != nil {
		return err
	}
	log.Ctx(ctx).Info().Msgf("loaded #%d entries", n)
	log.Ctx(ctx).Debug().Msgf("loaded: %+v", dest)
	return nil
}
