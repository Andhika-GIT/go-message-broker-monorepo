package shared

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ReadJsonRequest(request *http.Request, requestData any) error {
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(requestData)

	if err != nil {
		return err
	}

	return nil
}

func SendJsonResponse(w http.ResponseWriter, statusCode int, message string, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	res := APIResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}

	_ = json.NewEncoder(w).Encode(res)
}
