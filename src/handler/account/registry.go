package account

import (
	"tpconsulting/src/usecase/account"

	"github.com/labstack/echo/v4"
)

type AccountHandler interface {
	CreateAccount(c echo.Context) error
	GetPointByMobileNumber(c echo.Context) error
}

type accountHandler struct {
	accountUseCase account.AccountUseCase
}

func NewAccountUseCase(accountUseCase account.AccountUseCase) AccountHandler {
	return &accountHandler{
		accountUseCase: accountUseCase,
	}
}
