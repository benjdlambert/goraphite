package query

import "github.com/google/go-querystring/query"

type FindMetrics struct {
	Query  string `url:"query,omitempty"`
	Format string `url:"format,omitempty"`
}

func (f *FindMetrics) String() (string, error) {
	values, error := query.Values(f)
	if error != nil {
		return "", error
	}

	return values.Encode(), nil
}
