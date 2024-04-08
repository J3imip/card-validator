package types

import "encoding/json"

type Key struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type ValidationResultAttributes struct {
	Valid bool `json:"valid"`
}

type ValidationResult struct {
	Key
	Attributes ValidationResultAttributes `json:"attributes"`
}

type ValidationResultResponse struct {
	Data     ValidationResult `json:"data"`
	Included json.RawMessage  `json:"included"`
}
