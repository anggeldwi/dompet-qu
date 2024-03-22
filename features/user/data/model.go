package data

import (
	"dompet-qu/features/user"

	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	FullName    string
	Email       string `gorm:"unique"`
	Password    string
	PhoneNumber string `gorm:"unique"`
	Image       string
	Role        string `gorm:"not null"`
	Balance     int    `gorm:"default:0"`
}

func CoreToModel(input user.Core) User {
	return User{
		FullName:    input.FullName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
		Role:        input.Role,
		Balance:     input.Balance,
	}
}

func (u User) ModelToCore() user.Core {
	return user.Core{
		ID:          u.ID,
		FullName:    u.FullName,
		Email:       u.Email,
		Password:    u.Password,
		PhoneNumber: u.PhoneNumber,
		Role:        u.Role,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
