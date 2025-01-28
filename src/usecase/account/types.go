package account

type CreateAccountRequest struct {
	ThaiID       int    `json:"thai_id" required:"true"`
	MobileNumber []string `json:"mobile_number" required:"true"`
	Email        string   `json:"email" required:"true"`
	Name         string   `json:"name" required:"true"`
	Address      string   `json:"address" required:"true"`
	SubDistrict  string   `json:"sub_district" required:"true"`
	District     string   `json:"district" required:"true"`
	Province     string   `json:"province" required:"true"`
	ZipCode      string   `json:"zip_code" required:"true"`
}


type GetPointByMobileNumberResponse struct {
	ThaiID       int `json:"thai_id"`
	PointBalance int64 `json:"point_balance"`
}