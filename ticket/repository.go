package ticket

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(ticket Ticket) (Ticket, error)
	FindByID(ID int) (Ticket, error)
	Update(ticket Ticket) (Ticket, error)
	FindAll() ([]Ticket, error)
}

type repository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(ticket Ticket) (Ticket, error) {
	err := r.db.Create(&ticket).Error
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (r *repository) FindByID(ID int) (Ticket, error) {
	var ticket Ticket
	err := r.db.Where("id = ? AND is_deleted = ?", ID, false).Find(&ticket).Error
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (r *repository) Update(ticket Ticket) (Ticket, error) {
	err := r.db.Model(&ticket).Updates(map[string]interface{}{
		"ID": ticket.ID, 
		"Event": ticket.Event, 
		"Schedule": ticket.Schedule,
		"Quota"	 : ticket.Quota,
		"Price"  : ticket.Price,
		"IsDeleted" : ticket.IsDeleted,
		}).Error

	if err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (r *repository) FindAll() ([]Ticket, error) {
	var tickets []Ticket

	err := r.db.Where("is_deleted = ?", false).Find(&tickets).Error
	if err != nil {
		return tickets, err
	}

	return tickets, nil
}
