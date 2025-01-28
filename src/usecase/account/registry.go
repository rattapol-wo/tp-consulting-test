package account

import (
	"database/sql"
	"tpconsulting/src/repositories"
)

type AccountUseCase interface {
	CreateAccount(value CreateAccountRequest) (int64, *repositories.AccountResponse, error)
	GetPointByMobileNumber(mobileNumber string) (int64, *GetPointByMobileNumberResponse, error)
}

type accountUseCase struct {
	db          *sql.DB
	accountRepo repositories.AccountRepo
	pointRepo   repositories.PointRepo
}

func NewAccountUseCase(db *sql.DB, accountRepo repositories.AccountRepo, pointRepo repositories.PointRepo) AccountUseCase {
    return &accountUseCase{
        db:          db,
        accountRepo: accountRepo,
        pointRepo:   pointRepo,
    }
}
