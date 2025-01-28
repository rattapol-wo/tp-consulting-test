package repositories

import (
	"database/sql"
	"fmt"
)

type PointRepo interface {
	CreatePoint(tx *sql.Tx, thaiID int) (int64, error)
	UpdateBalance(thaiID int, actionType string)  (*int64, error)
}

type pointRepo struct {
	db *sql.DB
}

func NewPointRepo(db *sql.DB) PointRepo {
 return &pointRepo{db: db}
}

func (r *pointRepo) CreatePoint(tx *sql.Tx, thaiID int) (int64, error) {
	query := `INSERT INTO points (thai_id, created_at, updated_at) VALUES (?, NOW(), NOW())`
    result, err := tx.Exec(query,
		thaiID,
	)

	if err != nil {
		return 0, fmt.Errorf("failed insert point error: %w", err)
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to fetch last insert ID: %w", err)
	}

	return insertID, nil
}

func (u *pointRepo) UpdateBalance(thaiID int, actionType string) (*int64, error) {

    query := `
        UPDATE points
        SET balance = 
            CASE 
                WHEN ? = 'A' THEN balance + 1
                WHEN ? = 'D' THEN balance - 1
                ELSE balance
            END,
            updated_at = NOW()
        WHERE thai_id = ?;`

    _, err := u.db.Exec(query, actionType, actionType, thaiID)
    if err != nil {
        return nil, fmt.Errorf("failed to update balance: %w", err)
    }

	var newBalance int64
    selectQuery := `SELECT balance FROM points WHERE thai_id = ?;`
    err = u.db.QueryRow(selectQuery, thaiID).Scan(&newBalance)
    if err != nil {
        return nil , fmt.Errorf("failed to get updated balance: %w", err)
    }

    return &newBalance, nil
}
