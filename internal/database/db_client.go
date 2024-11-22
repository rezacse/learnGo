package database

import (
	"context"
	"fmt"
	"time"

	"github.com/rezacse/go-micro/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool
	GetAllCustomers(ctx context.Context, emailAddress string) ([]entities.Customer, error)
	AddCustomer(ctx context.Context, customer *entities.Customer) (*entities.Customer, error)
	GetCustomerById(ctx context.Context, id string) (*entities.Customer, error) 
	UpdateCustomerById(ctx context.Context, customer *entities.Customer) (*entities.Customer, error)
	DeleteCustomerById(ctx context.Context, id string) error

	GetAllProducts(ctx context.Context, vendorID string) ([]entities.Product, error)
	AddProduct(ctx context.Context, customer *entities.Product) (*entities.Product, error)
	GetProductById(ctx context.Context, id string) (*entities.Product, error)

	GetAllVendors(ctx context.Context) ([]entities.Vendor, error)
	AddVendor(ctx context.Context, customer *entities.Vendor) (*entities.Vendor, error)
	GetVendorById(ctx context.Context, id string) (*entities.Vendor, error)

	GetAllServices(ctx context.Context) ([]entities.Service, error)
	AddService(ctx context.Context, customer *entities.Service) (*entities.Service, error)
	GetServiceById(ctx context.Context, id string) (*entities.Service, error)
	
}

type Client struct {
	DB *gorm.DB
}

func NewDtabaseClient() (DatabaseClient , error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
	 "localhost", "postgres", "postgres", "postgres", 5432, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config {
		NamingStrategy: schema.NamingStrategy { TablePrefix: "wisdom."},
		NowFunc: func() time.Time { return time.Now().UTC() },
		QueryFields: true,
	})

	if err != nil {
		return nil, err
	}

	client := Client { DB: db }

	return client, nil
}

func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}

	if ready == "1" {
		return true
	}

	return false
}