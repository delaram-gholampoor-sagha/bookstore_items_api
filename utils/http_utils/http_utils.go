package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_utils-go/rest_errors"
)

func RespondJson(w http.ResponseWriter, StatusCode int, body interface{}) {

	w.Header().Set("Content-type", "applicaiton/json")
	w.WriteHeader(StatusCode)
	json.NewEncoder(w).Encode(body)

}

func RespondError(w http.ResponseWriter, err rest_errors.RestErr) {
	RespondJson(w, err.Status(), err)

}
