package query

import "github.com/google/go-querystring/query"

type GetMetrics struct {
	Target string `url:"target,omitempty"`
	Format string `url:"format,omitempty"`
}

func (r *GetMetrics) String() (string, error) {
	values, error := query.Values(r)
	if error != nil {
		return "", error
	}

	return values.Encode(), nil
}
