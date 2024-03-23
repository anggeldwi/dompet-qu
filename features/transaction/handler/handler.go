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
		} else {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert top up", nil))
		}
	}

	result := CoreToResponseTopUp(payment)

	return c.JSON(http.StatusOK, responses.WebResponse("success insert top up", result))
}

func (handler *TransactionHandler) Transfer(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newTransfer := TransferRequest{}
	errBind := c.Bind(&newTransfer)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data transfer not valid", nil))
	}

	transactionCore := RequestToCoreTransfer(newTransfer, uint(userIdLogin), newTransfer.PhoneNumber)

	if transactionCore.JenisTransaksi == "" {
		transactionCore.JenisTransaksi = "transfer"
	}
	transfer, errInsert := handler.transactionService.Transfer(userIdLogin, newTransfer.PhoneNumber, transactionCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	result := CoreToResponseTransfer(transfer)

	return c.JSON(http.StatusOK, responses.WebResponse("success insert transfer", result))
}
