package model

import "time"

type OrderStep struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Sequence   int        `json:"sequence"`
	ResourceID string     `json:"resourceId"`
	Status     StepStatus `json:"status"`
	StartTime  *time.Time `json:"startTime,omitempty"`
	EndTime    *time.Time `json:"endTime,omitempty"`
}

type StepStatus string

const (
	StepStatusPending    StepStatus = "pending"
	StepStatusInProgress StepStatus = "in_progress"
	StepStatusCompleted  StepStatus = "completed"
	StepStatusSkipped    StepStatus = "skipped"
)