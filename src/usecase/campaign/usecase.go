package campaign

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"tpconsulting/src/repositories"
)

func (u *campaignUseCase) CampaignAddPoint(reqBody CampaignAddPointRequest) (int64, *CampaignAddPointResponse, error) {

	campaign, err := u.campaignRepo.GetCampaign(reqBody.CampaignCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, nil, fmt.Errorf("campaign not found for campaign code: %s", reqBody.CampaignCode)
		}
		return http.StatusBadRequest, nil, err
	}
	account, err := u.accountRepo.GetPointByMobileNumber(reqBody.Mobile)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, nil, fmt.Errorf("account not found for mobile: %s", reqBody.Mobile)
		}
		return http.StatusBadRequest, nil, err
	}

	newBalance, err := u.pointRepo.UpdateBalance(account.ThaiID, campaign.PointAction)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK,
		&CampaignAddPointResponse{
			Balance: *newBalance,
		},
		nil
}

func (u *campaignUseCase) CreateCampaign(reqBody CampaignCreateCampaignRequest) (int64, error) {

	fmt.Print("reqBody.CampaignCode : ", reqBody.CampaignCode)
	exists, err := u.campaignRepo.GetCampaign(reqBody.CampaignCode)
	if err != nil && err != sql.ErrNoRows {
		return http.StatusBadRequest, err
	}

	if exists != nil {
		return http.StatusBadRequest, fmt.Errorf("campaign code already exists")
	}

	if reqBody.PointAction != "A" && reqBody.PointAction != "D" && reqBody.PointAction != "N" {
		return http.StatusBadRequest, fmt.Errorf("invalid point action type: %v", reqBody.PointAction)
	}

	body, err := transformDataCreateCampaign(reqBody)
	if err != nil {
		return http.StatusBadRequest, err
	}

	_, err = u.campaignRepo.CreateCampaign(*body)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func transformDataCreateCampaign(value CampaignCreateCampaignRequest) (*repositories.InsertCampaign, error) {

	layout := "02-01-2006"
	parsedStartDate, err := time.Parse(layout, value.StartDate)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return nil, err
	}

	parsedEndDate, err := time.Parse(layout, value.EndDate)
	if err != nil {
		fmt.Println("Error parsing end date:", err)
		return nil, err
	}

	return &repositories.InsertCampaign{
		CampaignCode: value.CampaignCode,
		PointAction:  value.PointAction,
		CampaignName: value.CampaignName,
		Description:  value.Description,
		Provision:    value.Provision,
		StartDate:    parsedStartDate,
		EndDAte:      parsedEndDate,
	}, nil
}
