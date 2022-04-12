package user

import (
	"time"
)

type User struct {
	ID        int
	//User Name
	//in: string
	Name      string
	//User Email
	//in: string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
}
