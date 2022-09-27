package defectdojo

import (
	"errors"
	"net/url"
	"strconv"
)

type PaginatedLister interface {
	NextRequestOptions() (*RequestOptions, error)
	SaveResults()
	RestoreResults()
	HasNext() bool
}

type PaginatedList[L any] struct {
	Count    int    // Number of Results
	Next     string // URL to next set of results
	Previous string // URL to previous set of results
	Results  []*L
	results  []*L
	//lint:ignore U1000 required field
	prefetch interface{} // Prefetch data, currently unsupported
}

// beta feature
// a trick/hack to persist results
// avoids json.unmarshal clobbering the results on multiple invocations
// should probably be a cleaner way to do this
func (p *PaginatedList[L]) SaveResults() {
	p.results = append(p.results, p.Results...)
}

func (p *PaginatedList[L]) RestoreResults() {
	p.Results = p.results
}

func (p *PaginatedList[L]) HasNext() bool {
	return p.Next != ""
}

func (p *PaginatedList[L]) NextRequestOptions() (*RequestOptions, error) {
	if p.Next != "" {
		return nil, errors.New("no additional results available")
	}
	u, err := url.Parse(p.Next)
	if err != nil {
		return nil, err
	}
	values := u.Query()
	o := values.Get("offset")
	offset, err := strconv.Atoi(o)
	if err != nil {
		return nil, err
	}

	l := values.Get("limit")
	limit, err := strconv.Atoi(l)
	if err != nil {
		return nil, err
	}
	return &RequestOptions{Offset: offset, Limit: limit}, nil
}
