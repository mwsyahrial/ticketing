package ticket


type TicketInput struct {
	Event    string    `json:"event" binding:"required"`
	Schedule string    `json:"schedule" binding:"required"`
	Quota    int	   `json:"quota" binding:"required"`
	Price    int	   `json:"price" binding:"required"`
}

type UpdateTicketInput struct {
	ID       int       `json:"id" binding:"required"`
	Event    string    `json:"event" binding:"required"`
	Schedule string    `json:"schedule" binding:"required"`
	Quota    int	   `json:"quota" binding:"required"`
	Price    int	   `json:"price" binding:"required"`
}

type IdTicketInput struct {
	ID int `json:"id" binding:"required"`
}
