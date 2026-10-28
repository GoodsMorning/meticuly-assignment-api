package model

type ResourceStatus string

const (
	ResourceStatusBusy    ResourceStatus = "busy"
	ResourceStatusIdle    ResourceStatus = "idle"
	ResourceStatusBlocked ResourceStatus = "blocked"
)

type Resource struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Status      ResourceStatus `json:"status"`
	Utilization float64        `json:"utilization"`
	Orders      []string       `json:"orders"`
	Note        string         `json:"note"`
}
