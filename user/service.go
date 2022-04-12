package user

import (
	"fmt"
)

type Service interface {
	AddUser(input UserInput) (User, error)
	GetUserByID(ID int) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(input UpdateUserInput) (User, error)
	DeleteUser(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AddUser(input UserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, fmt.Errorf("User not found, id : %d", ID)
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) UpdateUser(input UpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, fmt.Errorf("User not found, id : %d", input.ID)
	}

	user.Name = input.Name
	user.Email = input.Email

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) DeleteUser(ID int) (User, error) {
	deletedUser, err := s.repository.FindByID(ID)
	if err != nil {
		return deletedUser, err
	}

	if deletedUser.ID == 0 {
		return deletedUser, fmt.Errorf("User not found, id : %d", ID)
	}

	deletedUser.IsDeleted = true

	deletedUser, err = s.repository.Update(deletedUser)

	if err != nil {
		return deletedUser, err
	}

	return deletedUser, nil
}
