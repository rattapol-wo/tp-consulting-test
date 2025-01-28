package account

import (
	"net/http"
	"tpconsulting/src/usecase/account"

	"github.com/labstack/echo/v4"
)

func (h *accountHandler) CreateAccount(c echo.Context) error {
	var reqBody account.CreateAccountRequest

	// Bind JSON body struct
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"description": "Invalid request body",
			"error":       err.Error(),
		})
	}
	responseCode, _, err := h.accountUseCase.CreateAccount(reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"description": "Failed to create account",
			"error":       err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":        responseCode,
		"description": "Successfully create account",
	})
}

func (h *accountHandler) GetPointByMobileNumber(c echo.Context) error {

	mobileNumber := c.QueryParam("mobileNumber")

	if mobileNumber == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":        http.StatusBadRequest,
			"description": "Missing mobile number",
		})
	}

	responseCode, response, err := h.accountUseCase.GetPointByMobileNumber(mobileNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"description": "Failed to get point by mobile number",
			"error":       err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        responseCode,
		"description": "Successfully retrieved point balance",
		"balance":     response.PointBalance,
	})
}
