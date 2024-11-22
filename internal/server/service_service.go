package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezacse/go-micro/internal/dberrors"
	"github.com/rezacse/go-micro/internal/entities"
)

func (s *EchoServer) GetAllServices(ctx echo.Context) error {

	services, err := s.DB.GetAllServices(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, services)
}


func (s *EchoServer) AddService(ctx echo.Context) error {
	service := new(entities.Service)
	if err := ctx.Bind(service); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	service, err := s.DB.AddService(ctx.Request().Context(), service)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, service)
}


func (s *EchoServer) GetServiceById(ctx echo.Context) error {
	id := ctx.Param("id")
	customer, err := s.DB.GetServiceById(ctx.Request().Context(), id)
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