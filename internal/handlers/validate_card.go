package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/J3imip/card-validator/internal/requests"
	"github.com/J3imip/card-validator/internal/types"
	"github.com/J3imip/card-validator/internal/utils"
	"github.com/google/jsonapi"
)

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCard(r)
	if err != nil {
		utils.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
			Status: strconv.Itoa(http.StatusBadRequest),
		})
		return
	}

	now := time.Now()
	utils.Render(w, types.ValidationResultResponse{
		Data: types.ValidationResult{
			Key: types.Key{
				Type: "validate-card",
				ID:   "0",
			},
			Attributes: types.ValidationResultAttributes{
				Valid: request.Data.Attributes.ExpirationYear > uint16(now.Year()) ||
					(request.Data.Attributes.ExpirationYear == uint16(now.Year()) &&
						request.Data.Attributes.ExpirationMonth > uint8(now.Month())),
			},
		},
		Included: json.RawMessage("{}"),
	})
}
