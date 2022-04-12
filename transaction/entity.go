package transaction

import (
	"time"

)

type Transaction struct {
	ID        int
	UserId    int
	TicketId  int
	Quantity  int
	Price	  int
	TotalAmount int
	Status 	  string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
}
