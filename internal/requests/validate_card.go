package requests

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/J3imip/card-validator/internal/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func NewCard(r *http.Request) (*types.CardResponse, error) {
	var request types.CardResponse

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, validation.Errors{
			"/": errors.New("invalid request"),
		}.Filter()
	}

	return &request, validateNewCard(request)
}

func validateNewCard(request types.CardResponse) error {
	return validation.Errors{
		"/data/attributes/card_number": validation.Validate(
			request.Data.Attributes.CardNumber,
			validation.Required,
			// Validate credit card number
			is.CreditCard,
		),
		"/data/attributes/expiration_month": validation.Validate(
			request.Data.Attributes.ExpirationMonth,
			validation.Required,
			validation.By(func(value interface{}) error {
				month, ok := value.(uint8)
				if !ok {
					return errors.New("invalid expiration month")
				}

				if month < 1 || month > 12 {
					return errors.New("invalid expiration month")
				}

				return nil
			}),
		),
		"/data/attributes/expiration_year": validation.Validate(
			request.Data.Attributes.ExpirationYear,
			validation.Required,
		),
	}.Filter()
}
