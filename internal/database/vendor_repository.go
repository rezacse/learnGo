package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rezacse/go-micro/internal/dberrors"
	entities "github.com/rezacse/go-micro/internal/entities"
	"gorm.io/gorm"
)

func (c Client) GetAllVendors(ctx context.Context) ([]entities.Vendor, error) {
	var vendors []entities.Vendor

	result := c.DB.WithContext(ctx).
	Find(&vendors)

	return vendors, result.Error
}

func (c Client) AddVendor(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error) {
	vendor.VendorID = uuid.NewString()
	res := c.DB.WithContext(ctx).Create(&vendor)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, res.Error
	}
	return vendor, nil
}

func (c Client) GetVendorById(ctx context.Context, id string) (*entities.Vendor, error) {
	vendor := entities.Vendor{}

	res := c.DB.WithContext(ctx).
	Where(entities.Vendor{VendorID: id}).
	First(&vendor)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "vendor", ID: id}
		}
		return nil, res.Error
	}
	return &vendor, nil
}