package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezacse/go-micro/internal/dberrors"
	"github.com/rezacse/go-micro/internal/entities"
)

func (s *EchoServer) GetAllCustomers(ctx echo.Context) error {
	email := ctx.QueryParam("email")

	customers, err := s.DB.GetAllCustomers(ctx.Request().Context(), email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, customers)
}

func (s *EchoServer) AddCustomer(ctx echo.Context) error {
	customer := new(entities.Customer)
	if err := ctx.Bind(customer); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	customer, err := s.DB.AddCustomer(ctx.Request().Context(), customer)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, customer)
}


func (s *EchoServer) GetCustomerById(ctx echo.Context) error {
	id := ctx.Param("id")
	customer, err := s.DB.GetCustomerById(ctx.Request().Context(), id)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, customer)
}


func (s *EchoServer) UpdateCustomerById(ctx echo.Context) error {
	customer := new(entities.Customer)
	if err := ctx.Bind(customer); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	customer.CustomerID = ctx.Param("id")
	customer, err := s.DB.UpdateCustomerById(ctx.Request().Context(), customer)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, customer)
}

func (s *EchoServer) DeleteCustomerById(ctx echo.Context) error {
	customerID := ctx.Param("id")
	err := s.DB.DeleteCustomerById(ctx.Request().Context(), customerID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, nil)
}