package models

import "encoding/json"

type (
	Datapoint struct {
		Value     float64 `json:"value"`
		Timestamp int32   `json:"timestamp"`
	}
)

func (d *Datapoint) UnmarshalJSON(b []byte) error {
	var datapointsArray []float64
	err := json.Unmarshal(b, &datapointsArray)

	if err != nil {
		return err
	}

	d.Value = datapointsArray[0]
	d.Timestamp = int32(datapointsArray[1])
	return nil
}
