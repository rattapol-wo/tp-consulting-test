package repositories

import (
	"database/sql"
	"fmt"
	"time"
)

type CampaignRepo interface {
	CreateCampaign(reqBody InsertCampaign) (int64, error)
	GetCampaign(campaignCode string) (*Campaign, error)
}

type campaignRepo struct {
	db *sql.DB
}

func NewCampaignRepo(db *sql.DB) CampaignRepo {
	return &campaignRepo{db: db}
}

func (r *campaignRepo) CreateCampaign(reqBody InsertCampaign) (int64, error) {
	query := `INSERT INTO campaigns (campaign_code, point_action, campaign_name, description, provision, start_date, end_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`
	result, err := r.db.Exec(query,
		reqBody.CampaignCode,
		reqBody.PointAction,
		reqBody.CampaignName,
		reqBody.Description,
		reqBody.Provision,
		reqBody.StartDate,
		reqBody.EndDAte,
	)
	if err != nil {
		return 0, fmt.Errorf("Failed insert campaign error: %w", err)
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Failed to fetch last insert ID: %w", err)
	}

	return insertID, nil
}

func (r *campaignRepo) GetCampaign(campaignCode string) (*Campaign, error) {
	query := `SELECT * FROM campaigns WHERE campaign_code = ?`

	var campaign Campaign
	var startDate, endDate, createdAt, updatedAt []uint8

	err := r.db.QueryRow(query, campaignCode).Scan(
		&campaign.CampaignCode,
		&campaign.PointAction,
		&campaign.CampaignName,
		&campaign.Description,
		&campaign.Provision,
		&startDate,
		&endDate,
		&createdAt,
		&updatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to get campaign: %w", err)
	}

	campaign.StartDate, err = parseTime(startDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %w", err)
	}

	campaign.EndDate, err = parseTime(endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %w", err)
	}

	campaign.CreatedAt, err = parseTime(createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse created at: %w", err)
	}

	campaign.UpdatedAt, err = parseTime(updatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse updated at: %w", err)
	}

	return &campaign, nil
}

func parseTime(data []uint8) (time.Time, error) {
	if len(data) == 0 {
		return time.Time{}, nil
	}
	
	str := string(data)
	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}