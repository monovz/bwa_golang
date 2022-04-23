package transaction

import (
	"bwa_golang/campaign"
	"bwa_golang/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	User       user.User
	Campaign   campaign.Campaign
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
