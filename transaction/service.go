package transaction

import (
	"fmt"
	"ticketing/checkout"
	"ticketing/helper"
	"ticketing/ticket"
	"ticketing/user"
	"time"
)

type Service interface {
	AddTransaction(input CreateTransactionInput, t ticket.Service, u user.Service, c checkout.Service) (Transaction, ticket.Ticket, user.User, error)
	GetTransactionByID(ID int, t ticket.Service, u user.Service) (Transaction, ticket.Ticket, user.User, error)
	GetAllTransactions(ID int, t ticket.Service, u user.Service) ([]Transaction, []ticket.Ticket, user.User, error)
	DeleteTransaction(ID int) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewTransactionService(repository Repository) *service {
	return &service{repository}
}


func (s *service) AddTransaction(input CreateTransactionInput, t ticket.Service, u user.Service, c checkout.Service) (Transaction, ticket.Ticket, user.User, error) {
	currentCheckout, currentTicket, currentUser, err := c.GetCheckoutByID(input.CheckoutId, t, u)

	if err != nil {
		return Transaction{}, currentTicket, currentUser, fmt.Errorf("checkout not found, id : %d", input.CheckoutId)
	}

	if currentCheckout.Timeout.Before(time.Now()){
		return Transaction{}, currentTicket, currentUser, fmt.Errorf("checkout expired, id : %d", input.CheckoutId)
	}

	if currentTicket.Quota < currentCheckout.Quantity{
		return Transaction{}, currentTicket, currentUser, fmt.Errorf("quantity exceed available quota, id : %d", input.CheckoutId)
	}

	transaction := Transaction{}
	transaction.UserId = currentUser.ID
	transaction.TicketId = currentTicket.ID
	transaction.Quantity = currentCheckout.Quantity
	transaction.Price = currentTicket.Price
	transaction.TotalAmount =  currentTicket.Price * currentCheckout.Quantity
	transaction.Status = "PAID" 

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, currentTicket, currentUser, err
	}

	//update checkout
	updatedCheckout, err := c.DeleteCheckout(input.CheckoutId)

	_ = updatedCheckout

	if err != nil {
		return newTransaction, currentTicket, currentUser, err
	}

	//update ticket
	ticket := ticket.UpdateTicketInput{}
	ticket.ID = currentTicket.ID
	ticket.Event = currentTicket.Event
	ticket.Schedule = helper.DateToString(currentTicket.Schedule)
	ticket.Quota = currentTicket.Quota - currentCheckout.Quantity
	ticket.Price = currentTicket.Price
	updatedTicket, err := t.UpdateTicket(ticket)

	if err != nil {
		return newTransaction, updatedTicket, currentUser, err
	}

	return newTransaction, updatedTicket, currentUser, nil
}

func (s *service) GetTransactionByID(ID int, t ticket.Service, u user.Service) (Transaction, ticket.Ticket, user.User, error) {
	transaction, err := s.repository.FindByID(ID)
	if err != nil {
		return transaction, ticket.Ticket{}, user.User{}, err
	}

	if transaction.ID == 0 {
		return transaction, ticket.Ticket{}, user.User{}, fmt.Errorf("Transaction not found, id : %d", ID)
	}

	currentTicket, err := t.GetTicketByID(transaction.TicketId)

	if err != nil {
		return transaction, currentTicket, user.User{}, err
	}

	currentUser, err := u.GetUserByID(transaction.UserId)
	
	if err != nil {
		return transaction, currentTicket, currentUser, err
	}

	return transaction, currentTicket, currentUser, nil
}

func (s *service) GetAllTransactions(ID int, t ticket.Service, u user.Service) ([]Transaction, []ticket.Ticket, user.User, error) {
	transactions, err := s.repository.FindAll(ID)
	if err != nil {
		return transactions, []ticket.Ticket{}, user.User{}, err
	}

	tickets, err := t.GetAllTickets()

	if err != nil {
		return transactions, tickets, user.User{}, err
	}

	currentUser, err := u.GetUserByID(ID)
	
	if err != nil {
		return transactions, tickets, currentUser, err
	}


	return transactions, tickets, currentUser, nil
}


func (s *service) DeleteTransaction(ID int) (Transaction, error) {
	deletedTransaction, err := s.repository.FindByID(ID)
	if err != nil {
		return deletedTransaction, err
	}

	if deletedTransaction.ID == 0 {
		return deletedTransaction, fmt.Errorf("Transaction not found, id : %d", ID)
	}

	deletedTransaction.IsDeleted = true

	deletedTransaction, err = s.repository.Update(deletedTransaction)

	if err != nil {
		return deletedTransaction, err
	}

	return deletedTransaction, nil
}
