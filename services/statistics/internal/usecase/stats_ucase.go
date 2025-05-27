package usecase

import (
	"github.com/KassymbekovTimur/UTTMS/statistics/internal/model"
	"github.com/KassymbekovTimur/UTTMS/statistics/internal/repository"
)

// StatsUsecase contains business-logic
type StatsUsecase struct {
	repo repository.Repository
}

// NewStatsUsecase ccreates usecase
func NewStatsUsecase(r repository.Repository) *StatsUsecase {
	return &StatsUsecase{repo: r}
}

func (u *StatsUsecase) HandleOrderEvent(evt model.OrderEvent) error {
	return u.repo.SaveOrderEvent(evt)
}

func (u *StatsUsecase) HandleUserEvent(evt model.UserEvent) error {
	return u.repo.SaveUserEvent(evt)
}

func (u *StatsUsecase) GetUserOrderStats(userID string) (int, map[string]int, error) {
	total, err := u.repo.CountOrdersByUser(userID)
	if err != nil {
		return 0, nil, err
	}
	byHour, err := u.repo.CountOrdersByHour(userID)
	return total, byHour, err
}

func (u *StatsUsecase) GetUserStats() (int, int, error) {
	total, err := u.repo.CountTotalUsers()
	if err != nil {
		return 0, 0, err
	}
	reg, err := u.repo.CountRegisteredUsers()
	return total, reg, err
}
