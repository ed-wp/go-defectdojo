package defectdojo

import (
	"context"
	"errors"
	"fmt"
)

const (
	endpointAPIBase = "/api/v2/endpoints/"
)

type Endpoint struct {
	Id             int      `json:"id,omitempty" url:"id,omitempty"`
	Tags           []string `json:"tags,omitempty" url:"tags,omitempty"`
	Protocol       string   `json:"protocol,omitempty" url:"protocol,omitempty"`
	UserInfo       string   `json:"userinfo,omitempty" url:"userinfo,omitempty"`
	Host           string   `json:"host,omitempty" url:"host,omitempty"`
	Port           int      `json:"port,omitempty" url:"port,omitempty"`
	Path           string   `json:"path,omitempty" url:"path,omitempty"`
	Query          string   `json:"query,omitempty" url:"query,omitempty"`
	Fragment       string   `json:"fragment,omitempty" url:"fragment,omitempty"`
	Mitigated      bool     `json:"mitigated,omitempty" url:"mitigated,omitempty"`
	Product        int      `json:"product,omitempty" url:"product,omitempty"`
	EndpointParams []int    `json:"endpoint_params,omitempty" url:"endpoint_params,omitempty"`
	EndpointStatus []int    `json:"endpoint_status,omitempty" url:"endpoint_status,omitempty"`
}

func (d *DefectDojoAPI) GetEndpoints(ctx context.Context, endpoint *Endpoint, options *RequestOptions) (*PaginatedList[Endpoint], error) {
	if options.Limit == 0 {
		return nil, ErrorInvalidOptions
	}

	out := &PaginatedList[Endpoint]{}
	err := d.get(ctx, endpointAPIBase, options, endpoint, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) AddEndpoint(ctx context.Context, endpoint *Endpoint) (*Endpoint, error) {
	out := &Endpoint{}
	err := d.post(ctx, endpointAPIBase, endpoint, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) UpdateEndpoint(ctx context.Context, endpoint *Endpoint) (*Endpoint, error) {
	if endpoint.Id == 0 && endpoint.Product == 0 {
		return nil, errors.New("[defectdojo/UpdateEndpoint] error: cannot update product with blank ids")
	}

	out := &Endpoint{}
	err := d.patch(ctx, fmt.Sprintf("%s%v/", endpointAPIBase, endpoint.Id), endpoint, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) RemoveEndpoint(ctx context.Context, endpoint *Endpoint) error {

	return nil
}
