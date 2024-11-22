package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rezacse/go-micro/internal/dberrors"
	"github.com/rezacse/go-micro/internal/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]entities.Customer, error) {
	var customers []entities.Customer

	result := c.DB.WithContext(ctx).
	Where(entities.Customer{Email: emailAddress}).
	Find(&customers)

	return customers, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *entities.Customer) (*entities.Customer, error) {
	customer.CustomerID = uuid.NewString()
	res := c.DB.WithContext(ctx).Create(&customer)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, res.Error
	}
	return customer, nil
}

func (c Client) GetCustomerById(ctx context.Context, id string) (*entities.Customer, error) {
	customer := entities.Customer{}

	res := c.DB.WithContext(ctx).
	Where(entities.Customer{CustomerID: id}).
	First(&customer)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "customer", ID: id}
		}
		return nil, res.Error
	}
	return &customer, nil
}

func (c Client) UpdateCustomerById(ctx context.Context, customer *entities.Customer) (*entities.Customer, error) {
	var customers []entities.Customer

	res := c.DB.WithContext(ctx).
	Model(&customers).
	Clauses(clause.Returning{}).
	Where(&entities.Customer{CustomerID: customer.CustomerID}).
	Updates(entities.Customer{
		FirstName: customer.FirstName,
		LastName: customer.LastName,
		Phone: customer.Phone,
		Address: customer.Address,
	})
	
	if res.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{ Entity: "customer", ID: customer.CustomerID }
	}
	
	return &customers[0], nil
	//return customer, nil
}


func (c Client) DeleteCustomerById(ctx context.Context, id string) error {
	return c.DB.WithContext(ctx).Delete(entities.Customer{CustomerID: id}).Error
}