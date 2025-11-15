package user

import (
	"net/http"

	"gorm.io/gorm"
)

func ConvertToUsersResponse(users []User) []UserResponse {
	var usersResp []UserResponse

	for _, user := range users {
		resp := UserResponse{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		}

		usersResp = append(usersResp, resp)
	}

	return usersResp
}

func BindUserFilterFromRequest(r *http.Request) *UserFilter {
	return &UserFilter{
		Name:        r.URL.Query().Get("name"),
		Email:       r.URL.Query().Get("email"),
		PhoneNumber: r.URL.Query().Get("phone_number"),
		Search:      r.URL.Query().Get("search"),
	}
}

func FilterUserQuery(filter *UserFilter, query *gorm.DB) *gorm.DB {

	if filter.Search != "" {
		query = query.Where(
			"LOWER(name) ILIKE LOWER(?) OR LOWER(email) ILIKE LOWER(?) OR LOWER(phone_number) ILIKE LOWER(?)",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}
	if filter.PhoneNumber != "" {
		query = query.Where("phone_number = ?", filter.PhoneNumber)
	}

	return query
}
