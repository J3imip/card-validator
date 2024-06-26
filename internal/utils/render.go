package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
)

func RenderErr(w http.ResponseWriter, errs ...*jsonapi.ErrorObject) {
	if len(errs) == 0 {
		panic("expected non-empty errors slice")
	}

	status, err := strconv.ParseInt(errs[0].Status, 10, 64)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", jsonapi.MediaType)
	w.WriteHeader(int(status))
	jsonapi.MarshalErrors(w, errs)
}

func Render(w http.ResponseWriter, res interface{}) {
	w.Header().Set("content-type", jsonapi.MediaType)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}
}
