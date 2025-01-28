package repositories

import (
	"database/sql"
	"fmt"
)

type AccountRepo interface {
	GetPointByMobileNumber(mobileNumber string) (*AccountResponse, error)
	CreateAccount(tx *sql.Tx, reqBody InsertAccount) (int64, error)
}

type accountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) AccountRepo {
	return &accountRepo{db: db}
}

func (r *accountRepo) GetPointByMobileNumber(mobileNumber string) (*AccountResponse, error) {
	query := `SELECT a.thai_id, a.mobile_number, a.email, a.name, a.address, 
		       a.sub_district, a.district, a.province, a.zip_code, 
		       p.balance
		FROM accounts a
		INNER JOIN points p ON a.thai_id = p.thai_id
		WHERE FIND_IN_SET(?, a.mobile_number) > 0`

	var account AccountResponse
	err := r.db.QueryRow(query, mobileNumber).Scan(
		&account.ThaiID,
		&account.MobileNumber,
		&account.Email,
		&account.Name,
		&account.Address,
		&account.SubDistrict,
		&account.District,
		&account.Province,
		&account.ZipCode,
		&account.Balance,
	)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to get point: %w", err)
	}
	return &account, nil
}

func (r *accountRepo) CreateAccount(tx *sql.Tx, reqBody InsertAccount) (int64, error) {
	query := `INSERT INTO accounts (thai_id, mobile_number, email, name, address, sub_district, district, province, zip_code, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`
	result, err := tx.Exec(query,
		reqBody.ThaiID,
		reqBody.MobileNumber,
		reqBody.Email,
		reqBody.Name,
		reqBody.Address,
		reqBody.SubDistrict,
		reqBody.District,
		reqBody.Province,
		reqBody.ZipCode,
	)
	if err != nil {
		return 0, fmt.Errorf("failed insert account error: %w", err)
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to fetch last insert ID: %w", err)
	}

	return insertID, nil
}
