package transaction

import (
	"bwa_golang/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	User       user.User
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
