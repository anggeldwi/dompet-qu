package handler

import (
	"dompet-qu/app/middlewares"
	"dompet-qu/features/transaction"
	"dompet-qu/utils/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService transaction.TransactionServiceInterface
}

func New(ts transaction.TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{
		transactionService: ts,
	}
}

func (handler *TransactionHandler) TopUp(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newTopup := TopUpRequest{}
	errBind := c.Bind(&newTopup)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data top up not valid", nil))
	}

	transactionCore := RequestToCoreTopUp(newTopup, uint(userIdLogin))
	if transactionCore.JenisTransaksi == "" {
		transactionCore.JenisTransaksi = "top up"
	}
	payment, errInsert := handler.transactionService.TopUp(userIdLogin, transactionCore)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "bank is required") {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("bank is required", nil))
			// } else if strings.Contains(errInsert.Error(), "phone number is required") {
			// 	return c.JSON(http.StatusBadRequest, responses.WebResponse("phone number is required", nil))
			// } else if strings.Contains(errInsert.Error(), "greeting is required") {
			// 	return c.JSON(http.StatusBadRequest, responses.WebResponse("greeting is required", nil))
			// } else if strings.Contains(errInsert.Error(), "full name number is required") {
			// 	return c.JSON(http.StatusBadRequest, responses.WebResponse("full name number is required", nil))
			// } else if strings.Contains(errInsert.Error(), "email is required") {
			// 	return c.JSON(http.StatusBadRequest, responses.WebResponse("email is required", nil))
			// } else if strings.Contains(errInsert.Error(), "booking date is required") {
			// 	return c.JSON(http.StatusBadRequest, responses.WebResponse("booking date is required", nil))
			// } else if strings.Contains(errInsert.Error(), "maaf, anda tidak bisa menggunakan voucher ini karena total pembayaran anda terlalu rendah") {
			// 	return c.JSON(http.StatusBadRequest, responses.WebResponse("maaf, anda tidak bisa menggunakan voucher ini karena total pembayaran anda terlalu rendah", nil))
			// } else if strings.Contains(errInsert.Error(), "user has already used this voucher") {
			// 	return c.JSON(http.StatusConflict, responses.WebResponse("user has already used this voucher", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert booking", nil))
		}
	}

	result := CoreToResponseTopUp(payment)

	return c.JSON(http.StatusOK, responses.WebResponse("success insert booking", result))
}
