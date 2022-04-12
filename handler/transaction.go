package handler

import (
	"net/http"
	"ticketing/checkout"
	"ticketing/helper"
	"ticketing/ticket"
	"ticketing/transaction"
	"ticketing/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
	ticketService ticket.Service
	userService user.Service
	checkoutService checkout.Service
}

func NewTransactionHandler(transactionService transaction.Service, ticketService ticket.Service, 
		userService user.Service, checkoutService checkout.Service) *transactionHandler {
	return &transactionHandler{transactionService, ticketService, userService, checkoutService}
}

func (h *transactionHandler) AddTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add transaction failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTransaction, currentTicket, currentUser, err := h.transactionService.AddTransaction(input, h.ticketService, h.userService, h.checkoutService)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatTransaction(newTransaction, currentUser.Name, currentTicket.Event)

	response := helper.APIResponse("Transaction has been added", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetTransaction(c *gin.Context) {

	var input transaction.IdTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Get transaction failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentTransaction, currentTicket, currentUser, err := h.transactionService.GetTransactionByID(input.ID, h.ticketService, h.userService)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatTransaction(currentTransaction, currentUser.Name, currentTicket.Event)

	response := helper.APIResponse("Get transaction data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *transactionHandler) GetTransactions(c *gin.Context) {
	var input transaction.TransactionsInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Get all transaction failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	allTransaction, allTicket, currentUser, err := h.transactionService.GetAllTransactions(input.UserId, h.ticketService, h.userService)

	if err != nil {
		response := helper.APIResponse("Get all transaction failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var formatters []transaction.TransactionFormatter
	for _, i := range allTransaction {
		var currentTicket ticket.Ticket
		for _, j := range allTicket {
			if j.ID == i.TicketId {
				currentTicket = j
			}
		}

		formatter := transaction.FormatTransaction(i, currentUser.Name, currentTicket.Event)

		formatters = append(formatters,formatter)
	}

	response := helper.APIResponse("Get transaction data", http.StatusOK, "success", formatters)

	c.JSON(http.StatusOK, response)

}


func (h *transactionHandler) DeleteTransaction(c *gin.Context) {

	var input transaction.IdTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Delete transaction failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedTransaction, err := h.transactionService.DeleteTransaction(input.ID)

	_ = deletedTransaction

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete transaction data", http.StatusOK, "success", nil)

	c.JSON(http.StatusOK, response)

}
