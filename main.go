package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ticketing/checkout"
	"ticketing/handler"
	"ticketing/ticket"
	"ticketing/transaction"
	"ticketing/user"

	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {
	//configure mysql server
	dsn := "root:admin@tcp(127.0.0.1:3306)/ticketing?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	ticketRepository := ticket.NewTicketRepository(db)
	ticketService := ticket.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService)

	checkoutRepository := checkout.NewCheckoutRepository(db)
	checkoutService := checkout.NewCheckoutService(checkoutRepository)
	checkoutHandler := handler.NewCheckoutHandler(checkoutService, ticketService, userService)

	transactionRepository := transaction.NewTransactionRepository(db)
	transactionService := transaction.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService, ticketService, userService, checkoutService)

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/users", userHandler.GetUsers)
	router.POST("/user", userHandler.GetUser)
	router.POST("/user/create", userHandler.AddUser)
	router.POST("/user/update", userHandler.UpdateUser)
	router.POST("/user/delete", userHandler.DeleteUser)

	router.POST("/tickets", ticketHandler.GetTickets)
	router.POST("/ticket", ticketHandler.GetTicket)
	router.POST("/ticket/create", ticketHandler.AddTicket)
	router.POST("/ticket/update", ticketHandler.UpdateTicket)
	router.POST("/ticket/delete", ticketHandler.DeleteTicket)

	router.POST("/checkouts", checkoutHandler.GetCheckouts)
	router.POST("/checkout", checkoutHandler.GetCheckout)
	router.POST("/checkout/create", checkoutHandler.AddCheckout)
	router.POST("/checkout/delete", checkoutHandler.DeleteCheckout)

	router.POST("/transactions", transactionHandler.GetTransactions)
	router.POST("/transaction", transactionHandler.GetTransaction)
	router.POST("/transaction/create", transactionHandler.AddTransaction)
	router.POST("/transaction/delete", transactionHandler.DeleteTransaction)

	router.Run()


}
