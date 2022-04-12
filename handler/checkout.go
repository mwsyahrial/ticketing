package handler

import (
	"net/http"
	"ticketing/checkout"
	"ticketing/helper"
	"ticketing/ticket"
	"ticketing/user"

	"github.com/gin-gonic/gin"
)

type checkoutHandler struct {
	checkoutService checkout.Service
	ticketService ticket.Service
	userService user.Service
}

func NewCheckoutHandler(checkoutService checkout.Service, ticketService ticket.Service, userService user.Service) *checkoutHandler {
	return &checkoutHandler{checkoutService, ticketService, userService}
}

func (h *checkoutHandler) AddCheckout(c *gin.Context) {
	var input checkout.CheckoutInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add checkout failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCheckout, currentTicket, currentUser, err := h.checkoutService.AddCheckout(input, h.ticketService, h.userService)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ticketPrice := currentTicket.Price

	totalAmount := ticketPrice * newCheckout.Quantity

	formatter := checkout.FormatCheckout(newCheckout, ticketPrice, totalAmount, currentUser.Name, currentTicket.Event)

	response := helper.APIResponse("Checkout has been added", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *checkoutHandler) GetCheckout(c *gin.Context) {

	var input checkout.IdCheckoutInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Get checkout failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentCheckout, currentTicket, currentUser, err := h.checkoutService.GetCheckoutByID(input.ID, h.ticketService, h.userService)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ticketPrice := currentTicket.Price

	totalAmount := ticketPrice * currentCheckout.Quantity

	formatter := checkout.FormatCheckout(currentCheckout, ticketPrice, totalAmount, currentUser.Name, currentTicket.Event)

	response := helper.APIResponse("Get checkout data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *checkoutHandler) GetCheckouts(c *gin.Context) {
	var input checkout.CheckoutsInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Get all checkout failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	allCheckout, allTicket, currentUser, err := h.checkoutService.GetAllCheckouts(input.UserId, h.ticketService, h.userService)

	if err != nil {
		response := helper.APIResponse("Get all checkout failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var formatters []checkout.CheckoutFormatter
	for _, i := range allCheckout {
		var currentTicket ticket.Ticket
		for _, j := range allTicket {
			if j.ID == i.TicketId {
				currentTicket = j
			}
		}

		ticketPrice := currentTicket.Price
	
		totalAmount := ticketPrice * i.Quantity
	
		formatter := checkout.FormatCheckout(i, ticketPrice, totalAmount, currentUser.Name, currentTicket.Event)

		formatters = append(formatters,formatter)
	}

	response := helper.APIResponse("Get checkout data", http.StatusOK, "success", formatters)

	c.JSON(http.StatusOK, response)

}


func (h *checkoutHandler) DeleteCheckout(c *gin.Context) {

	var input checkout.IdCheckoutInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Delete checkout failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedCheckout, err := h.checkoutService.DeleteCheckout(input.ID)

	_ = deletedCheckout

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete checkout data", http.StatusOK, "success", nil)

	c.JSON(http.StatusOK, response)

}
