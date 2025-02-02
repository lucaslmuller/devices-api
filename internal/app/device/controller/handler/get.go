package handler

import (
	"net/http"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/utils"
)

func (api api) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, _, _, err := utils.ParseRequest[any](r, []string{})
	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	output, err := api.Endpoints.GetAll(*ctx)
	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(ctx, w, http.StatusOK, output)
}
func (api api) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx, input, params, err := utils.ParseRequest[dto.GetDeviceByIDInput](r, []string{"id"})
	input.ID = params["id"]

	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	output, validationErr, err := api.Endpoints.GetByID(*ctx, input)

	if validationErr != nil {
		utils.RespondWithError(ctx, w, http.StatusBadRequest, validationErr.Error())
		return
	}

	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	if output == nil {
		utils.RespondWithError(ctx, w, http.StatusNotFound, "Device not found")
		return
	}

	utils.RespondWithJSON(ctx, w, http.StatusOK, output)
}
func (api api) GetByBrand(w http.ResponseWriter, r *http.Request) {
	ctx, input, params, err := utils.ParseRequest[dto.GetDeviceByBrandInput](r, []string{"brand"})
	input.Brand = params["brand"]

	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, err.Error())

		return
	}

	output, validationErr, err := api.Endpoints.GetByBrand(*ctx, input)

	if validationErr != nil {
		utils.RespondWithError(ctx, w, http.StatusBadRequest, validationErr.Error())
		return
	}

	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(ctx, w, http.StatusOK, output)
}
func (api api) GetByState(w http.ResponseWriter, r *http.Request) {
	ctx, input, params, err := utils.ParseRequest[dto.GetDeviceByStateInput](r, []string{"state"})
	input.State = params["state"]

	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	output, validationErr, err := api.Endpoints.GetByState(*ctx, input)

	if validationErr != nil {
		utils.RespondWithError(ctx, w, http.StatusBadRequest, validationErr.Error())
		return
	}

	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(ctx, w, http.StatusOK, output)
}
