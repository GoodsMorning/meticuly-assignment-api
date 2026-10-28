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


type StepType string

const (
	StepMachining1  StepType = "Machining-1"
	StepMachining2  StepType = "Machining-2"
	StepAssembly    StepType = "Assembly"
	StepQualityCheck StepType = "Quality Check"
	StepPackaging   StepType = "Packaging"
)

type MaterialRequirement struct {
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Numerator   int    `json:"numerator"`
	Denominator int    `json:"denominator"`
}

type StepTemplate struct {
	ID          StepType             `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Duration    int                  `json:"duration"`
	Resources   []string             `json:"resources"`
	Materials   []MaterialRequirement `json:"materials"`
}