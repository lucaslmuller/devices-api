package handler

import (
	"net/http"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/utils"
)

func (api *api) Create(w http.ResponseWriter, r *http.Request) {
	ctx, input, _, err := utils.ParseRequest[dto.CreateDeviceInput](r, []string{})

	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	output, validationErr, err := api.Endpoints.Create(*ctx, input)

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
