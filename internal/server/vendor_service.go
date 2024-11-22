package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezacse/go-micro/internal/dberrors"
	"github.com/rezacse/go-micro/internal/entities"
)

func (s *EchoServer) GetAllVendors(ctx echo.Context) error {

	vendors, err := s.DB.GetAllVendors(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, vendors)
}


func (s *EchoServer) AddVendor(ctx echo.Context) error {
	vendor := new(entities.Vendor)
	if err := ctx.Bind(vendor); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	vendor, err := s.DB.AddVendor(ctx.Request().Context(), vendor)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, vendor)
}

func (s *EchoServer) GetVendorById(ctx echo.Context) error {
	id := ctx.Param("id")
	customer, err := s.DB.GetVendorById(ctx.Request().Context(), id)
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