package repository

import "github.com/KassymbekovTimur/UTTMS/participant/internal/model"

type Repository interface {
	Save(participant *model.Participant) error
	FindByID(id string) (*model.Participant, error)
	FindByScheduleID(scheduleid string) ([]*model.Participant, error)
	Update(participant *model.Participant) error
	Delete(id string) error
}
