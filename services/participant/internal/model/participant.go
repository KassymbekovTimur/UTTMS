package model

type Participant struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	ScheduleIDs []string `json:"schedule_ids"`
	Status      string   `json:"status"`
}
