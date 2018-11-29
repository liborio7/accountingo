package cache

import (
	"context"
	"github.com/rs/zerolog/log"
)

type client interface {
	SetKey(context.Context, Model) error
	GetKey(context.Context, Model) error
}

type Service struct {
	c client
}

func (s *Service) SetKey(ctx context.Context, m Model) error {
	log.Ctx(ctx).Info().Msgf("set key %+v", m)
	if err := s.c.SetKey(ctx, m); err != nil {
		return err
	}
	log.Ctx(ctx).Debug().Msgf("key set %+v", m)
	return nil
}

func (s *Service) GetKey(ctx context.Context, m Model) error {
	log.Ctx(ctx).Info().Msgf("get key %+v", m)
	if err := s.c.GetKey(ctx, m); err != nil {
		return err
	}
	log.Ctx(ctx).Debug().Msgf("key get %+v", m)
	return nil
}
