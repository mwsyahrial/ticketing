package transaction

type TransactionFormatter struct {
	ID        int       `json:"id"`
	UserId    int		`json:"user_id"`
	TicketId  int		`json:"ticket_id"`
	Name      string    `json:"name"`
	Event	  string    `json:"event"`
	Quantity  int		`json:"quantity"`
	Status    string	`json:"status"`
	Price	  int 		`json:"price"`
	TotalAmount int     `json:"total_amount"`

}

func FormatTransaction(transaction Transaction, name, event string) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:    transaction.ID,
		UserId: transaction.UserId,
		TicketId: transaction.TicketId,
		Name: name,
		Event: event,
		Quantity: transaction.Quantity,
		Status: transaction.Status,
		Price: transaction.Price,
		TotalAmount: transaction.TotalAmount,
	}

	return formatter
}