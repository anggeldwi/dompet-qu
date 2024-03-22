package handler

import "dompet-qu/features/user"

type UserResponseLogin struct {
	FullName string `json:"full_name" form:"full_name"`
	Role     string `json:"role" form:"role"`
	Token    string `json:"token" form:"token"`
}

type UserResponse struct {
	FullName    string `json:"full_name" form:"full_name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Email       string `json:"email" form:"email"`
	Image       string `json:"image" form:"image"`
	Balance     int    `json:"balance" form:"balance"`
}

func CoreToResponseLogin(data *user.Core, token string) UserResponseLogin {
	var result = UserResponseLogin{
		FullName: data.FullName,
		Role:     data.Role,
		Token:    token,
	}
	return result
}

func CoreToResponseUser(data *user.Core) UserResponse {
	var result = UserResponse{
		FullName:    data.FullName,
		PhoneNumber: data.PhoneNumber,
		Email:       data.Email,
		Image:       data.Image,
		Balance:     data.Balance,
	}
	return result
}
