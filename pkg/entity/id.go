package entity

import (
	"errors"

	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, errors.New("ID parsing incorrectly")
	}
	return id, nil
}
