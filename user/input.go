package user

type UserInput struct{
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
}

type UpdateUserInput struct{
	ID         int    `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
}

type IdUserInput struct{
	ID         int    `json:"id" binding:"required"`
}
