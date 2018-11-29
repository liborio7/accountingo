package account

import (
	"github.com/satori/go.uuid"
	"time"
)

type Model struct {
	Id           uuid.UUID `db:"id" id:"true"`
	Version      uint8     `db:"version"`
	InsertDate   time.Time `db:"insert_date"`
	InsertMillis int64     `db:"insert_millis"`
	Name         string    `db:"name"`
	Surname      string    `db:"surname"`
	Currency     string    `db:"currency"`
	Amount       int64     `db:"amount"`
}
