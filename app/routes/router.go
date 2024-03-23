package routes

import (
	"dompet-qu/app/middlewares"
	ud "dompet-qu/features/user/data"
	uh "dompet-qu/features/user/handler"
	us "dompet-qu/features/user/service"

	td "dompet-qu/features/transaction/data"
	th "dompet-qu/features/transaction/handler"
	ts "dompet-qu/features/transaction/service"

	"dompet-qu/utils/encrypts"
	"dompet-qu/utils/externalapi"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hash := encrypts.NewHashService()
	// cloudinaryUploader := cloudinary.New()
	midtrans := externalapi.New()

	userData := ud.New(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService)

	transactionData := td.New(db, midtrans)
	transactionService := ts.New(transactionData)
	transactionHandlerAPI := th.New(transactionService)

	// define routes/ endpoint USERS
	e.POST("/users", userHandlerAPI.RegisterUser)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	// define routes/ endpoint transactions
	e.POST("/topup", transactionHandlerAPI.TopUp, middlewares.JWTMiddleware())
	e.POST("/transfer", transactionHandlerAPI.Transfer, middlewares.JWTMiddleware())
}
