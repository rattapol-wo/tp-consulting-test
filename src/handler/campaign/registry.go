package campaign

import (
	"tpconsulting/src/usecase/campaign"

	"github.com/labstack/echo/v4"
)

type CampaignHandler interface {
	CreateCampaign(c echo.Context) error
	CampaignAddPoint(c echo.Context) error
}

type campaignHandler struct {
	campaignUseCase campaign.CampaignUseCase
}

func NewCampaignUseCase(campaignUseCase campaign.CampaignUseCase) CampaignHandler {
	return &campaignHandler{
		campaignUseCase: campaignUseCase,
	}
}
