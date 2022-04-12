package checkout

import (
	"ticketing/helper"
)

type CheckoutFormatter struct {
	ID        int       `json:"id"`
	UserId    int		`json:"user_id"`
	TicketId  int		`json:"ticket_id"`
	Name      string    `json:"name"`
	Event	  string    `json:"event"`
	Quantity  int		`json:"quantity"`
	Timeout   string	`json:"timeout"`
	Price	  int 		`json:"price"`
	TotalAmount int     `json:"total_amount"`

}

func FormatCheckout(checkout Checkout, price, totalAmount int, name, event string) CheckoutFormatter {
	formatter := CheckoutFormatter{
		ID:    checkout.ID,
		UserId: checkout.UserId,
		TicketId: checkout.TicketId,
		Name: name,
		Event: event,
		Quantity: checkout.Quantity,
		Timeout: helper.DateToString(checkout.Timeout),
		Price: price,
		TotalAmount: totalAmount,
	}

	return formatter
}