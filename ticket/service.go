package ticket

import (
	"fmt"
	"ticketing/helper"
)

type Service interface {
	AddTicket(input TicketInput) (Ticket, error)
	GetTicketByID(ID int) (Ticket, error)
	GetAllTickets() ([]Ticket, error)
	UpdateTicket(input UpdateTicketInput) (Ticket, error)
	DeleteTicket(ID int) (Ticket, error)
}

type service struct {
	repository Repository
}

func NewTicketService(repository Repository) *service {
	return &service{repository}
}


func (s *service) AddTicket(input TicketInput) (Ticket, error) {
	ticket := Ticket{}
	ticket.Event = input.Event
	ticket.Schedule = helper.StringToDate(input.Schedule)
	ticket.Quota = input.Quota
	ticket.Price = input.Price

	newTicket, err := s.repository.Save(ticket)
	if err != nil {
		return newTicket, err
	}

	return newTicket, nil
}

func (s *service) GetTicketByID(ID int) (Ticket, error) {
	ticket, err := s.repository.FindByID(ID)
	if err != nil {
		return ticket, err
	}

	if ticket.ID == 0 {
		return ticket, fmt.Errorf("Ticket not found, id : %d", ID)
	}

	return ticket, nil
}

func (s *service) GetAllTickets() ([]Ticket, error) {
	tickets, err := s.repository.FindAll()
	if err != nil {
		return tickets, err
	}

	return tickets, nil
}

func (s *service) UpdateTicket(input UpdateTicketInput) (Ticket, error) {
	ticket, err := s.repository.FindByID(input.ID)
	if err != nil {
		return ticket, err
	}

	if ticket.ID == 0 {
		return ticket, fmt.Errorf("Ticket not found, id : %d", input.ID)
	}

	ticket.Event = input.Event
	ticket.Schedule = helper.StringToDate(input.Schedule)
	ticket.Quota = input.Quota
	ticket.Price = input.Price

	updatedTicket, err := s.repository.Update(ticket)
	if err != nil {
		return updatedTicket, err
	}

	return updatedTicket, nil
}

func (s *service) DeleteTicket(ID int) (Ticket, error) {
	deletedTicket, err := s.repository.FindByID(ID)
	if err != nil {
		return deletedTicket, err
	}

	if deletedTicket.ID == 0 {
		return deletedTicket, fmt.Errorf("Ticket not found, id : %d", ID)
	}

	deletedTicket.IsDeleted = true

	deletedTicket, err = s.repository.Update(deletedTicket)

	if err != nil {
		return deletedTicket, err
	}

	return deletedTicket, nil
}
