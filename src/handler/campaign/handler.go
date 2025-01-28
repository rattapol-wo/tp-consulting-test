package campaign

import (
	"net/http"
	"tpconsulting/src/usecase/campaign"

	"github.com/labstack/echo/v4"
)

func (h *campaignHandler) CreateCampaign(c echo.Context) error {
	var reqBody campaign.CampaignCreateCampaignRequest

	// Bind JSON body struct
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"description": "Invalid request body",
			"error":       err.Error(),
		})
	}

	_, err := h.campaignUseCase.CreateCampaign(reqBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"description": "Failed to create campaign",
			"error":       err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":        http.StatusCreated,
		"description": "Successfully create campaign",
	})
}

func (h *campaignHandler) CampaignAddPoint(c echo.Context) error {
	var reqBody campaign.CampaignAddPointRequest

	// Bind JSON body struct
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	responseCode, response, err := h.campaignUseCase.CampaignAddPoint(reqBody)
	if err != nil {
		return c.JSON(int(responseCode), map[string]string{
			"message": "Failed to campaign add point",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":          200,
		"description":   "Successfully added points",
		"point_balance": response.Balance,
	})
}
