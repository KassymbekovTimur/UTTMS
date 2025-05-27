package model

import "time"

// OrderEvent  describes the event of creating an order
type OrderEvent struct {
	UserID    string
	OrderID   string
	Timestamp time.Time
}

// UserEvent describes actions with the user
type UserEvent struct {
	UserID    string
	Action    string //i.e. "created", "deleted" etc.
	Timestamp time.Time
}
