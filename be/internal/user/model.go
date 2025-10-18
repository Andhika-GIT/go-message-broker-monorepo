package user

import "time"

type UserImport struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserResponse struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	CreatedAt   time.Time `json:"created_at"`
}

func ConvertToUsersResponse(users []User) []UserResponse {
	var usersResp []UserResponse

	for _, user := range users {
		resp := UserResponse{
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			CreatedAt:   user.CreatedAt,
		}

		usersResp = append(usersResp, resp)
	}

	return usersResp
}
