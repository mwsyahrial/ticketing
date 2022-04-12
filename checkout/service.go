package checkout

import (
	"fmt"
	"ticketing/ticket"
	"ticketing/user"
	"time"
)

type Service interface {
	AddCheckout(input CheckoutInput, t ticket.Service, u user.Service) (Checkout, ticket.Ticket, user.User, error)
	GetCheckoutByID(ID int, t ticket.Service, u user.Service) (Checkout, ticket.Ticket, user.User, error)
	GetAllCheckouts(ID int, t ticket.Service, u user.Service) ([]Checkout, []ticket.Ticket, user.User, error)
	DeleteCheckout(ID int) (Checkout, error)
}

type service struct {
	repository Repository
}

func NewCheckoutService(repository Repository) *service {
	return &service{repository}
}


func (s *service) AddCheckout(input CheckoutInput, t ticket.Service, u user.Service) (Checkout, ticket.Ticket, user.User, error) {
	checkout := Checkout{}
	checkout.UserId = input.UserId
	checkout.TicketId = input.TicketId
	checkout.Quantity = input.Quantity
	checkout.Timeout = time.Now().Add(time.Hour * 1)

	newCheckout, err := s.repository.Save(checkout)
	if err != nil {
		return newCheckout, ticket.Ticket{}, user.User{}, err
	}


	currentTicket, err := t.GetTicketByID(newCheckout.TicketId)

	if err != nil {
		return newCheckout, currentTicket, user.User{}, err
	}

	//check ticket quota
	if (input.Quantity > currentTicket.Quota){
		return newCheckout, currentTicket, user.User{}, fmt.Errorf("input quantity exceed available quota")
	}

	currentUser, err := u.GetUserByID(newCheckout.UserId)
	
	if err != nil {
		return newCheckout, currentTicket, currentUser, err
	}
	
	return newCheckout, currentTicket, currentUser, nil
}

func (s *service) GetCheckoutByID(ID int, t ticket.Service, u user.Service) (Checkout, ticket.Ticket, user.User, error) {
	checkout, err := s.repository.FindByID(ID)
	if err != nil {
		return checkout, ticket.Ticket{}, user.User{}, err
	}

	if checkout.ID == 0 {
		return checkout, ticket.Ticket{}, user.User{}, fmt.Errorf("Checkout not found, id : %d", ID)
	}

	currentTicket, err := t.GetTicketByID(checkout.TicketId)

	if err != nil {
		return checkout, currentTicket, user.User{}, err
	}

	currentUser, err := u.GetUserByID(checkout.UserId)
	
	if err != nil {
		return checkout, currentTicket, currentUser, err
	}

	return checkout, currentTicket, currentUser, nil
}

func (s *service) GetAllCheckouts(ID int, t ticket.Service, u user.Service) ([]Checkout, []ticket.Ticket, user.User, error) {
	checkouts, err := s.repository.FindAll(ID)
	if err != nil {
		return checkouts, []ticket.Ticket{}, user.User{}, err
	}

	tickets, err := t.GetAllTickets()

	if err != nil {
		return checkouts, tickets, user.User{}, err
	}

	currentUser, err := u.GetUserByID(ID)
	
	if err != nil {
		return checkouts, tickets, currentUser, err
	}


	return checkouts, tickets, currentUser, nil
}


func (s *service) DeleteCheckout(ID int) (Checkout, error) {
	deletedCheckout, err := s.repository.FindByID(ID)
	if err != nil {
		return deletedCheckout, err
	}

	if deletedCheckout.ID == 0 {
		return deletedCheckout, fmt.Errorf("Checkout not found, id : %d", ID)
	}

	deletedCheckout.IsDeleted = true

	deletedCheckout, err = s.repository.Update(deletedCheckout)

	if err != nil {
		return deletedCheckout, err
	}

	return deletedCheckout, nil
}
