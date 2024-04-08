package types

import "encoding/json"

type ValidateCard struct {
	Attributes CardAttributes `json:"attributes"`
}

type CardAttributes struct {
	CardNumber      string `json:"card_number"`
	ExpirationMonth uint8  `json:"expiration_month"`
	ExpirationYear  uint16 `json:"expiration_year"`
}

type CardResponse struct {
	Data     ValidateCard    `json:"data"`
	Included json.RawMessage `json:"included"`
}
