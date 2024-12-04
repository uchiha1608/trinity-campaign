package models

import "time"

type Campaign struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	MaxUsersAllowed int       `json:"max_allowed_user"`
	RegisteredUsers int       `json:"registered_users"`
}
