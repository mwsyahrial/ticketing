package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
	}

	return formatter
}