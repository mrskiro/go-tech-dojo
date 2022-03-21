package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, params interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&params)
	if err != nil {
		return err
	}
	return nil
}
