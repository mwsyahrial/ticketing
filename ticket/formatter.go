package ticket

import (
	"ticketing/helper"
)

type TicketFormatter struct {
	ID       int       `json:"id"`
	Event    string    `json:"event"`
	Schedule string    `json:"schedule"`
	Quota    int       `json:"quota"`
	Price    int       `json:"price"`
}

func FormatTicket(ticket Ticket) TicketFormatter {
	formatter := TicketFormatter{
		ID:    ticket.ID,
		Event:  ticket.Event,
		Schedule: helper.DateToString(ticket.Schedule),
		Quota: ticket.Quota,
		Price: ticket.Price,
	}

	return formatter
}