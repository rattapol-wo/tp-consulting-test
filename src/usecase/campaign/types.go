package campaign

type CampaignAddPointRequest struct {
	Mobile       string `json:"mobile"`
	CampaignCode string `json:"campaign_code"`
	Channel      string `json:"channel`
}

type CampaignCreateCampaignRequest struct {
	CampaignCode string `json:"campaign_code"`
	PointAction  string `json:"point_action"`
	CampaignName string `json:"campaign_name"`
	Description  string `json:"description"`
	Provision    string `json:"provision"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

type CampaignAddPointResponse struct {
	Balance int64 `json:"balance"`
}
