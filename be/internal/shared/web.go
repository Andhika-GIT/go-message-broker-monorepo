package shared

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
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
		Success: true,
		Message: message,
		Data:    data,
	}

	_ = json.NewEncoder(w).Encode(res)
}

func SendJsonErrorResponse(w http.ResponseWriter, err error, data any) {
	if e, ok := err.(*Error); ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(e.Code)

		res := APIResponse{
			Status:  e.Code,
			Success: false,
			Message: e.Message,
			Data:    data,
		}

		_ = json.NewEncoder(w).Encode(res)
	}
}
