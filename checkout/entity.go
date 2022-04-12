package checkout

import (
	"time"

)

type Checkout struct {
	ID        int
	UserId    int
	TicketId  int
	Quantity  int
	Timeout   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
}
