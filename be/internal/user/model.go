package user

type UserImport struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

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
