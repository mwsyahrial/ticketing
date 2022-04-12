package checkout


type CheckoutInput struct {
	UserId    int		`json:"user_id" binding:"required"`
	TicketId  int		`json:"ticket_id" binding:"required"`
	Quantity	  int	`json:"quantity" binding:"required"`
}


type IdCheckoutInput struct {
	ID int `json:"id" binding:"required"`
}

type CheckoutsInput struct {
	UserId int `json:"user_id" binding:"required"`
}
