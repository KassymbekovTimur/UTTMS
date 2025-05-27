package usecase

import (
	"github.com/KassymbekovTimur/UTTMS/participant/internal/model"
	"github.com/KassymbekovTimur/UTTMS/participant/internal/repository"
	"github.com/google/uuid"
)

type ParticipantUsecase struct {
	cache repository.Repository
	db    repository.Repository
}

func NewParticipantUsecase(cache, db repository.Repository) *ParticipantUsecase {
	return &ParticipantUsecase{cache: cache, db: db}
}

func (u *ParticipantUsecase) Register(name, email string) (*model.Participant, error) {
	id := uuid.New().String()
	participant := &model.Participant{
		ID:          id,
		Name:        name,
		Email:       email,
		ScheduleIDs: []string{},
		Status:      "pending",
	}
	if err := u.db.Save(participant); err != nil {
		return nil, err
	}
	_ = u.cache.Save(participant)
	return participant, nil
}

func (u *ParticipantUsecase) ConfirmEmail(token string) error {
	p, err := u.db.FindByID(token)
	if err != nil {
		return err
	}
	p.Status = "active"
	if err := u.db.Update(p); err != nil {
		return err
	}
	_ = u.cache.Save(p)
	return nil
}

func (u *ParticipantUsecase) JoinSchedule(participantID, scheduleID string) error {
	p, err := u.db.FindByID(participantID)
	if err != nil {
		return err
	}
	for _, id := range p.ScheduleIDs {
		if id == scheduleID {
			return nil
		}
	}
	p.ScheduleIDs = append(p.ScheduleIDs, scheduleID)
	if err := u.db.Update(p); err != nil {
		return err
	}
	_ = u.cache.Save(p)
	return nil
}

func (u *ParticipantUsecase) LeaveSchedule(participantID, scheduleID string) error {
	p, err := u.db.FindByID(participantID)
	if err != nil {
		return err
	}
	newList := []string{}
	for _, id := range p.ScheduleIDs {
		if id != scheduleID {
			newList = append(newList, id)
		}
	}
	p.ScheduleIDs = newList
	if err := u.db.Update(p); err != nil {
		return err
	}
	_ = u.cache.Save(p)
	return nil
}

func (u *ParticipantUsecase) GetParticipant(id string) (*model.Participant, error) {
	p, err := u.cache.FindByID(id)
	if err == nil {
		return p, nil
	}
	p, err = u.db.FindByID(id)
	if err != nil {
		return nil, err
	}
	_ = u.cache.Save(p)
	return p, nil
}

func (u *ParticipantUsecase) ListParticipants(scheduleID string) ([]*model.Participant, error) {
	list, err := u.cache.FindByScheduleID(scheduleID)
	if err == nil && len(list) > 0 {
		return list, nil
	}
	list, err = u.db.FindByScheduleID(scheduleID)
	if err != nil {
		return nil, err
	}
	return list, nil
}
