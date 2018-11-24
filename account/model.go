package account

import (
	"github.com/satori/go.uuid"
	"time"
)

type Model struct {
	Id           uuid.UUID
	Name         string
	Surname      string
	Currency     string
	Amount       int64
	InsertDate   time.Time
	InsertMillis uint64
	Version      uint8
}
