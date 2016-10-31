package models

type (
	Target struct {
		Name          string      `json:"target"`
		RawDatapoints [][]float64 `json:"datapoints"`
	}
)

func (t *Target) Datapoints() []Datapoint {
	var datapoints = make([]Datapoint, len(t.RawDatapoints))

	for index, rawDatapoint := range t.RawDatapoints {
		datapoints[index] = Datapoint{
			X: rawDatapoint[0],
			Y: rawDatapoint[1],
		}
	}

	return datapoints
}
