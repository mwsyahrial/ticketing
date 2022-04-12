package transaction

type CreateTransactionInput struct {
	CheckoutId int `json:"checkout_id" binding:"required"`
}

type IdTransactionInput struct {
	ID int `json:"id" binding:"required"`
}

type TransactionsInput struct {
	UserId int `json:"user_id" binding:"required"`
}
