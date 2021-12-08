package campaign

import (
	"fmt"
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].Filename
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}
	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                   `json:"id"`
	Name             string                `json:"name"`
	ShortDescription string                `json:"short_description"`
	Description      string                `json:"description"`
	ImageUrl         string                `json:"image_url"`
	GoalAmount       int                   `json:"goal_amount"`
	CurrentAmount    int                   `json:"current_amount"`
	UserID           int                   `json:"user_id"`
	Slug             string                `json:"slug"`
	Perks            []string              `json:"perks"`
	User             CampaignUserFormatter `json:"user"`
	Images           []CampaignImageFormatter
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCamapaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.ImageUrl = ""

	// karena belum tentu index 0 adalah is_primary
	// if len(campaign.CampaignImages) > 0 {
	// 	campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	// }

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormatter.Perks = perks

	// userFormatter

	campaignUserFormatter := CampaignUserFormatter{}
	campaignUser := campaign.User
	campaignUserFormatter.Name = campaignUser.Name
	campaignUserFormatter.ImageUrl = campaignUser.AvatarFileName

	campaignDetailFormatter.User = campaignUserFormatter

	// imageFormatter
	campaignImagesFormatter := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		fmt.Println(image)
		imageFormatter := CampaignImageFormatter{}
		imageFormatter.ImageUrl = image.Filename
		imageFormatter.IsPrimary = false

		if image.IsPrimary == 1 {
			imageFormatter.IsPrimary = true
			campaignDetailFormatter.ImageUrl = image.Filename
		}

		campaignImagesFormatter = append(campaignImagesFormatter, imageFormatter)
	}

	campaignDetailFormatter.Images = campaignImagesFormatter

	return campaignDetailFormatter
}
