package user

import "time"

type Core struct {
	ID          uint
	FullName    string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	Image       string
	Role        string
	Balance     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserDataInterface interface {
	Insert(input Core) error
	Login(phoneNumber, password string) (data *Core, err error)
	SelectById(userIdLogin int) (*Core, error)
	Update(userIdLogin int, input Core) error
	Delete(userIdLogin int) error
}

type UserServiceInterface interface {
	Insert(input Core) error
	Login(phoneNumber, password string) (data *Core, token string, err error)
	SelectById(userIdLogin int) (*Core, error)
	Update(userIdLogin int, input Core) error
	Delete(userIdLogin int) error
}
