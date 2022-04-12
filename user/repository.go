package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
	FindAll() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ? AND is_deleted = ?", ID, false).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Updates(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Where("is_deleted = ?", false).Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}
