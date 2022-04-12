package handler

import (
	"net/http"
	"ticketing/helper"
	"ticketing/ticket"

	"github.com/gin-gonic/gin"
)

type ticketHandler struct {
	ticketService ticket.Service
}

func NewTicketHandler(ticketService ticket.Service) *ticketHandler {
	return &ticketHandler{ticketService}
}

func (h *ticketHandler) AddTicket(c *gin.Context) {
	var input ticket.TicketInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add ticket failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTicket, err := h.ticketService.AddTicket(input)

	if err != nil {
		response := helper.APIResponse("Add ticket failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := ticket.FormatTicket(newTicket)

	response := helper.APIResponse("Ticket has been added", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *ticketHandler) GetTicket(c *gin.Context) {

	var input ticket.IdTicketInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Get ticket failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentTicket, err := h.ticketService.GetTicketByID(input.ID)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := ticket.FormatTicket(currentTicket)

	response := helper.APIResponse("Get ticket data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *ticketHandler) GetTickets(c *gin.Context) {
	var input struct{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Get all ticket failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	allTicket, err := h.ticketService.GetAllTickets()

	if err != nil {
		response := helper.APIResponse("Get all ticket failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var formatters []ticket.TicketFormatter
	for _, i := range allTicket {
		formatter := ticket.FormatTicket(i)
		formatters = append(formatters,formatter)
	}

	response := helper.APIResponse("Get ticket data", http.StatusOK, "success", formatters)

	c.JSON(http.StatusOK, response)

}

func (h *ticketHandler) UpdateTicket(c *gin.Context) {

	var input ticket.UpdateTicketInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update ticket failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedTicket, err := h.ticketService.UpdateTicket(input)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := ticket.FormatTicket(updatedTicket)

	response := helper.APIResponse("Update ticket data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *ticketHandler) DeleteTicket(c *gin.Context) {

	var input ticket.IdTicketInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Delete ticket failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedTicket, err := h.ticketService.DeleteTicket(input.ID)

	_ = deletedTicket

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete ticket data", http.StatusOK, "success", nil)

	c.JSON(http.StatusOK, response)

}
