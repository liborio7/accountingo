package payment

import "github.com/satori/go.uuid"

type Model struct {
	id            uuid.UUID
	fromAccountId uuid.UUID
	toAccountId   uuid.UUID
	currency      string
	amount        int64
}
