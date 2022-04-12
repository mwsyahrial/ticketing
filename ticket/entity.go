package ticket

import (
	"time"

)

type Ticket struct {
	ID        int
	Event     string
	Schedule  time.Time
	Quota	  int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
}
