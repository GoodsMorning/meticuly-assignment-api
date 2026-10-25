package model

type Resource struct {
	ID   string       `json:"id"`
	Name string       `json:"name"`
	Type ResourceType `json:"type"`
}

type ResourceType string

const (
	ResourceTypeMachine ResourceType = "machine"
	ResourceTypeStation ResourceType = "station"
)