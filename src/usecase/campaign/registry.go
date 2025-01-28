package campaign

import "tpconsulting/src/repositories"

type CampaignUseCase interface {
	CampaignAddPoint(reqBody CampaignAddPointRequest) (int64, *CampaignAddPointResponse, error)
	CreateCampaign(reqBody CampaignCreateCampaignRequest) (int64, error)
}

type campaignUseCase struct {
	accountRepo  repositories.AccountRepo
	pointRepo    repositories.PointRepo
	campaignRepo repositories.CampaignRepo
}

func NewCampaignUseCase(accountRepo repositories.AccountRepo, pointRepo repositories.PointRepo, campaignRepo repositories.CampaignRepo) *campaignUseCase {
	return &campaignUseCase{
		accountRepo:  accountRepo,
		pointRepo:    pointRepo,
		campaignRepo: campaignRepo,
	}
}
