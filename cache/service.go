package cache

import (
	"fmt"
)

type Service struct {
	c Cache
}

func NewService(c Cache) *Service {
	return &Service{c: c}
}

func (s *Service) SetKey(m Model) error {
	fmt.Println("insert", m)
	if err := s.c.SetKey(m); err != nil {
		return err
	}
	fmt.Println("inserted", m)
	return nil
}

func (s *Service) GetKey(m Model) error {
	fmt.Println("load", m)
	if err := s.c.GetKey(m); err != nil {
		return err
	}
	fmt.Println("loaded", m)
	return nil
}
