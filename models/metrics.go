package models

type (
	Metric struct {
		Text          string `json:"text"`
		Expandable    int    `json:"expandable"`
		ID            string `json:"id"`
		AllowChildren int    `json:"allowChildren"`
		Leaf          int    `json:"leaf"`
	}
)

func (m *Metric) IsLeaf() bool {
	return m.Leaf == 1
}

func (m *Metric) AllowsChildren() bool {
	return m.AllowChildren == 1
}

func (m *Metric) IsExpandable() bool {
	return m.Expandable == 1
}
