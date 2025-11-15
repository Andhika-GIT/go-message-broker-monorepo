package shared

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type APIResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type PaginationRequest struct {
	Page    int `json:"page" query:"page"`
	PerPage int `json:"per_page" query:"per_page"`
}

type Paginated[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
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

func GetPaginationParams(r *http.Request) *PaginationRequest {
	query := r.URL.Query()

	pageStr := query.Get("page")
	perPageStr := query.Get("per_page")

	page, err := strconv.Atoi(pageStr)

	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)

	if err != nil || page < 1 {
		perPage = 1
	}

	return &PaginationRequest{
		Page:    page,
		PerPage: perPage,
	}

}
