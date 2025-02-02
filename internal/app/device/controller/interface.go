package controller

import (
	"net/http"
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
