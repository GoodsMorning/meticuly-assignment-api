package model

type Stock struct {
	Resources []Resource      `json:"resources"`
	Materials []MaterialStock `json:"materials"`
}
