package controller

import (
	"context"
	"net/http"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
)

type IDeviceAPI interface {
	GetByID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByBrand(w http.ResponseWriter, r *http.Request)
	GetByState(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type DeviceEndpoints struct {
	Get    IGetEndpoints
	Create ICreateEndpoints
	Update IUpdateEndpoints
	Delete IDeleteEndpoints
}

type IGetEndpoints interface {
	ByID(ctx context.Context, input *dto.GetDeviceByIDInput) (output *dto.Output, validationErr error, internalErr error)
	ByBrand(ctx context.Context, input *dto.GetDeviceByBrandInput) (output []dto.Output, validationErr error, internalErr error)
	ByState(ctx context.Context, input *dto.GetDeviceByStateInput) (output []dto.Output, validationErr error, internalErr error)
	All(ctx context.Context) (output []dto.Output, err error)
}
type ICreateEndpoints interface {
	Device(ctx context.Context, input *dto.CreateDeviceInput) (output *dto.Output, validationErr error, internalErr error)
}
type IDeleteEndpoints interface {
	Device(ctx context.Context, input *dto.DeleteDeviceInput) (validationErr error, internalErr error)
}
type IUpdateEndpoints interface {
	Device(ctx context.Context, input *dto.UpdateDeviceInput) (output *dto.Output, validationErr error, internalErr error)
}
