package handler

import (
	"net/http"
	"startup-gin/campaign"
	"startup-gin/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		if err != nil {
			response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	formatter := campaign.FormatCampaigns(campaigns)

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
