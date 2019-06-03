package http_transport

import (
	"encoding/json"
	"lookup/internal/errs"
	"net/http"
)

func ParseJsonBody(request *http.Request, d interface{}) error {
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&d)
	if err != nil {
		return errs.BadRequest("invalid JSON request")
	}
	return nil
}
