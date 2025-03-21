package entities

import "time"

type Job struct {
	ID               string
	OutputBucketPath string
	Status           string
	Video            *VideoEntity
	Error            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
