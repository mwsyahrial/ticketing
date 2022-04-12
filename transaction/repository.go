package transaction

import "gorm.io/gorm"

type Repository interface {
	Save(transaction Transaction) (Transaction, error)
	FindByID(ID int) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	FindAll(ID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindByID(ID int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Where("id = ? AND is_deleted = ?", ID, false).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Updates(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindAll(ID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Where("user_id = ? AND is_deleted = ?", ID, false).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
