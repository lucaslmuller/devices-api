package handler

import (
	"net/http"

	"github.com/lucaslmuller/technical-test/internal/app/device/controller/dto"
	"github.com/lucaslmuller/technical-test/internal/utils"
)

func (api *api) Update(w http.ResponseWriter, r *http.Request) {
	ctx, input, params, err := utils.ParseRequest[dto.UpdateDeviceInput](r, []string{"id"})

	input.ID = params["id"]

	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	output, validationErr, err := api.Endpoints.Update(*ctx, input)

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
