package repositories

import "time"

type Account struct {
	ThaiID       int       `json:"thai_id" required:"true"`
	MobileNumber string    `json:"mobile_number" required:"true"`
	Email        string    `json:"email" required:"true"`
	Name         string    `json:"name" required:"true"`
	Address      string    `json:"address" required:"true"`
	SubDistrict  string    `json:"sub_district" required:"true"`
	District     string    `json:"district" required:"true"`
	Province     string    `json:"province" required:"true"`
	ZipCode      string    `json:"zip_code" required:"true"`
	CreatedAt    time.Time `json:"created_at" required:"true"`
	UpdatedAt    time.Time `json:"updated_at" required:"true"`
}

type AccountResponse struct {
	ThaiID       int       `json:"thai_id" required:"true"`
	MobileNumber string    `json:"mobile_number" required:"true"`
	Email        string    `json:"email" required:"true"`
	Name         string    `json:"name" required:"true"`
	Address      string    `json:"address" required:"true"`
	SubDistrict  string    `json:"sub_district" required:"true"`
	District     string    `json:"district" required:"true"`
	Province     string    `json:"province" required:"true"`
	ZipCode      string    `json:"zip_code" required:"true"`
	Balance      int64     `json:"balance" required:"true"`
}

type InsertAccount struct {
	ThaiID       int    `json:"thai_id" required:"true"`
	MobileNumber string `json:"mobile_number" required:"true"`
	Email        string `json:"email" required:"true"`
	Name         string `json:"name" required:"true"`
	Address      string `json:"address" required:"true"`
	SubDistrict  string `json:"sub_district" required:"true"`
	District     string `json:"district" required:"true"`
	Province     string `json:"province" required:"true"`
	ZipCode      string `json:"zip_code" required:"true"`
}

type Campaign struct {
	CampaignCode string    `json:"campaign_code" required:"true"`
	PointAction  string    `json:"point_action" required:"true"`
	CampaignName string    `json:"campaign_name" required:"true"`
	Description  string    `json:"description" required:"true"`
	Provision    string    `json:"provision" required:"true"`
	StartDate    time.Time `json:"start_date" required:"true"`
	EndDate      time.Time `json:"end_date" required:"true"`
	CreatedAt    time.Time `json:"created_at" required:"true"`
	UpdatedAt    time.Time `json:"updated_at" required:"true"`
}

type InsertCampaign struct {
	CampaignCode string    `json:"campaign_code" required:"true"`
	PointAction  string    `json:"point_action" required:"true"`
	CampaignName string    `json:"campaign_name" required:"true"`
	Description  string    `json:"description" required:"true"`
	Provision    string    `json:"provision" required:"true"`
	StartDate    time.Time `json:"start_date" required:"true"`
	EndDAte      time.Time `json:"end_date" required:"true"`
}

type Point struct {
	PointID   int64     `json:"point_id" required:"true"`
	Balance   int64     `json:"balance" required:"true"`
	ThaiID    int64     `json:"thai_id" required:"true"`
	CreatedAt time.Time `json:"created_at" required:"true"`
	UpdatedAt time.Time `json:"updated_at" required:"true"`
}
