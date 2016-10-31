package models

type (
	Target struct {
		Name       string      `json:"target"`
		Datapoints []Datapoint `json:"datapoints"`
	}
)
