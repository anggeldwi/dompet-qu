package handler

import "dompet-qu/features/user"

type UserRequest struct {
	FullName    string `json:"full_name" form:"full_name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Image       string `json:"image" form:"image"`
	Role        string `json:"role" form:"role"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Password    string `json:"password" form:"password"`
}

func RequestToCore(input UserRequest) user.Core {
	return user.Core{
		FullName:    input.FullName,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
		Password:    input.Password,
		Image:       input.Image,
		Role:        input.Role,
	}
}
