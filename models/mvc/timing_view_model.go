package mvc

import (
	"time"
)

type TimingViewModel struct {
	Id            int       `json:"id"`
	ClientProject string    `json:"client_project"`
	Client        string    `json:"client"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	StartDateTime time.Time `json:"start_date_time"`
	EndDateTime   time.Time `json:"end_date_time"`
	Status        string    `json:"status"`
}
