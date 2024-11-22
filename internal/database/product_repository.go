package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rezacse/go-micro/internal/dberrors"
	entities "github.com/rezacse/go-micro/internal/entities"
	"gorm.io/gorm"
)

func (c Client) GetAllProducts(ctx context.Context, vendorID string) ([]entities.Product, error) {
	var products []entities.Product

	result := c.DB.WithContext(ctx).
	Where(entities.Product{VendorID: vendorID}).
	Find(&products)

	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	product.ProductID = uuid.NewString()
	res := c.DB.WithContext(ctx).Create(&product)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, res.Error
	}
	return product, nil
}

func (c Client) GetProductById(ctx context.Context, id string) (*entities.Product, error) {
	product := entities.Product{}

	res := c.DB.WithContext(ctx).
	Where(entities.Product{ProductID: id}).
	First(&product)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "product", ID: id}
		}
		return nil, res.Error
	}
	return &product, nil
}