package account

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"tpconsulting/src/repositories"
)

func (u *accountUseCase) CreateAccount(value CreateAccountRequest) (int64, *repositories.AccountResponse, error) {

	for _, mobileNumber := range value.MobileNumber {
		exists, err := u.accountRepo.GetPointByMobileNumber(mobileNumber)
		if err != nil && err != sql.ErrNoRows {
			return http.StatusBadRequest, nil, err
		}
		if exists != nil {
			return http.StatusBadRequest, exists, fmt.Errorf("account already exists")
		}
	}

	tx, err := u.db.Begin()
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	mobileString := strings.Join(value.MobileNumber, ",")

	_, err = u.accountRepo.CreateAccount(tx, transformDataInsert(value, mobileString))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	pointID, err := u.pointRepo.CreatePoint(tx, value.ThaiID)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	tx.Commit()

	return http.StatusCreated, transformDataResp(value, value.ThaiID, pointID, mobileString), nil
}

func (u *accountUseCase) GetPointByMobileNumber(mobileNumber string) (int64, *GetPointByMobileNumberResponse, error) {

	resp, err := u.accountRepo.GetPointByMobileNumber(mobileNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, nil, fmt.Errorf("account not found with mobile number: %s", mobileNumber)
		}
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, transformDataGetPointResp(resp.Balance, resp.ThaiID), nil
}

func transformDataInsert(value CreateAccountRequest, mobileNumber string) repositories.InsertAccount {
	
	return repositories.InsertAccount{
		ThaiID:       value.ThaiID,
		MobileNumber: mobileNumber,
		Email:        value.Email,
		Name:         value.Name,
		Address:      value.Address,
		SubDistrict:  value.SubDistrict,
		District:     value.District,
		Province:     value.Province,
		ZipCode:      value.ZipCode,
	}
}

func transformDataResp(value CreateAccountRequest, accountID int, pointID int64, mobileNumber string) *repositories.AccountResponse {
	return &repositories.AccountResponse{
		ThaiID:       accountID,
		MobileNumber: mobileNumber,
		Email:        value.Email,
		Name:         value.Name,
		Address:      value.Address,
		SubDistrict:  value.SubDistrict,
		District:     value.District,
		Province:     value.Province,
		ZipCode:      value.ZipCode,
	}
}

func transformDataGetPointResp(pointBalance int64, thaiID int) *GetPointByMobileNumberResponse {
	return &GetPointByMobileNumberResponse{
		ThaiID:       thaiID,
        PointBalance: pointBalance,
	}
}