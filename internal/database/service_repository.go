package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rezacse/go-micro/internal/dberrors"
	entities "github.com/rezacse/go-micro/internal/entities"
	"gorm.io/gorm"
)

func (c Client) GetAllServices(ctx context.Context) ([]entities.Service, error) {
	var services []entities.Service

	result := c.DB.WithContext(ctx).
	Find(&services)

	return services, result.Error
}

func (c Client) AddService(ctx context.Context, service *entities.Service) (*entities.Service, error) {
	service.ServiceID = uuid.NewString()
	res := c.DB.WithContext(ctx).Create(&service)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, res.Error
	}
	return service, nil
}

func (c Client) GetServiceById(ctx context.Context, id string) (*entities.Service, error) {
	service := entities.Service{}

	res := c.DB.WithContext(ctx).
	Where(entities.Service{ServiceID: id}).
	First(&service)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "service", ID: id}
		}
		return nil, res.Error
	}
	return &service, nil
}