package repository

import "github.com/KassymbekovTimur/UTTMS/statistics/internal/model"

// Repository describes an event holder interface
type Repository interface {
	SaveOrderEvent(evt model.OrderEvent) error
	SaveUserEvent(evt model.UserEvent) error
	CountOrdersByUser(userID string) (int, error)
	CountOrdersByHour(userID string) (map[string]int, error)
	CountTotalUsers() (int, error)
	CountRegisteredUsers() (int, error)
}
