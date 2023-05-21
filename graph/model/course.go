package model

type Course struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	Category    []*Category `json:"category,omitempty"`
}
