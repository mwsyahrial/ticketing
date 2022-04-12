package checkout

import "gorm.io/gorm"

type Repository interface {
	Save(checkout Checkout) (Checkout, error)
	FindByID(ID int) (Checkout, error)
	Update(checkout Checkout) (Checkout, error)
	FindAll(ID int) ([]Checkout, error)
}

type repository struct {
	db *gorm.DB
}

func NewCheckoutRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(checkout Checkout) (Checkout, error) {
	err := r.db.Create(&checkout).Error
	if err != nil {
		return checkout, err
	}

	return checkout, nil
}

func (r *repository) FindByID(ID int) (Checkout, error) {
	var checkout Checkout
	err := r.db.Where("id = ? AND is_deleted = ?", ID, false).Find(&checkout).Error
	if err != nil {
		return checkout, err
	}

	return checkout, nil
}

func (r *repository) Update(checkout Checkout) (Checkout, error) {
	err := r.db.Updates(&checkout).Error

	if err != nil {
		return checkout, err
	}

	return checkout, nil
}

func (r *repository) FindAll(ID int) ([]Checkout, error) {
	var checkouts []Checkout

	err := r.db.Where("user_id = ? AND is_deleted = ?", ID, false).Find(&checkouts).Error
	if err != nil {
		return checkouts, err
	}

	return checkouts, nil
}
