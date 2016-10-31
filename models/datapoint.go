package models

import "encoding/json"

type (
	Datapoint struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}
)

func (d *Datapoint) UnmarshalJSON(b []byte) error {
	var datapointsArray []float64
	err := json.Unmarshal(b, &datapointsArray)

	if err != nil {
		return err
	}

	d.X = datapointsArray[0]
	d.Y = datapointsArray[1]

	return nil
}
