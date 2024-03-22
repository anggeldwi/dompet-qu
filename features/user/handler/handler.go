package handler

import (
	"dompet-qu/app/middlewares"
	"dompet-qu/features/user"
	"dompet-qu/utils/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
	// cld         cloudinary.CloudinaryUploaderInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
		// cld:         cloudinaryUploader,
	}
}

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	userCore := RequestToCore(newUser)
	errInsert := handler.userService.Insert(userCore)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "Error 1062 (23000): Duplicate entry") {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error register data. "+errInsert.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error register data. "+errInsert.Error(), nil))
		}
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success register data", nil))
}

func (handler *UserHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}
	result, token, err := handler.userService.Login(reqData.PhoneNumber, reqData.Password)
	if err != nil {
		if strings.Contains(err.Error(), "email wajib diisi.") {
			return c.JSON(http.StatusUnauthorized, responses.WebResponse("error login. "+err.Error(), nil))
		} else if strings.Contains(err.Error(), "password wajib diisi.") {
			return c.JSON(http.StatusUnauthorized, responses.WebResponse("error login. "+err.Error(), nil))
		} else if strings.Contains(err.Error(), "password tidak sesuai.") {
			return c.JSON(http.StatusUnauthorized, responses.WebResponse("error login. "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusUnauthorized, responses.WebResponse("error login. "+err.Error(), nil))
		}
	}
	var responseData = CoreToResponseLogin(result, token)
	return c.JSON(http.StatusOK, responses.WebResponse("success login", responseData))
}

func (handler *UserHandler) GetUser(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)

	result, errSelect := handler.userService.SelectById(userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	var userResult = CoreToResponseUser(result)
	return c.JSON(http.StatusOK, responses.WebResponse("success read data", userResult))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	// Extract user ID from JWT token
	userID := middlewares.ExtractTokenUserId(c)
	// log.Println("UserID:", userID)

	// Check if the user ID is valid
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("unauthorized", nil))
	}

	// Use the extracted user ID from the token
	idParam := userID

	// Bind the request data
	var userData = UserRequest{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	// Convert request data to core model
	userCore := RequestToCore(userData)

	// Call the service to update the user profile
	err := handler.userService.Update(idParam, userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)

	errDelete := handler.userService.Delete(userIdLogin)
	if errDelete != nil {
		if strings.Contains(errDelete.Error(), "error record not found") {
			return c.JSON(http.StatusNotFound, responses.WebResponse("error delete data. "+errDelete.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data. "+errDelete.Error(), nil))
		}
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}
