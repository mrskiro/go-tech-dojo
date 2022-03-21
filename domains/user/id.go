package user

import "github.com/google/uuid"

type ID string

func GenID() string {
	uuid := uuid.New()
	return uuid.String()
}

func (id ID) String() string {
	return string(id)
}
