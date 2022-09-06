package defectdojo

import (
	"context"
	"errors"
	"fmt"
)

const (
	metadataAPIBase = "/api/v2/metadata/"
)

type Metadata struct {
	Id       int    `json:"id,omitempty" url:"id,omitempty"`
	Product  int    `json:"product,omitempty" url:"product,omitempty"`
	Endpoint int    `json:"endpoint,omitempty" url:"endpoint,omitempty"`
	Finding  int    `json:"finding,omitempty" url:"finding,omitempty"`
	Name     string `json:"name" url:"name"`
	Value    string `json:"value" url:"value"`
}

type PaginatedMetadataList struct {
	Count    int         // Number of Results
	Next     string      // URL to next set of results
	Previous string      // URL to previous set of results
	Results  []*Metadata // List of Metadata results
	//lint:ignore U1000 required field
	prefetch interface{} // Prefetch data, currently unsupported
}

func (d *DefectDojoAPI) GetMetadatas(ctx context.Context, metadata *Metadata, options *RequestOptions) (*PaginatedMetadataList, error) {
	out := &PaginatedMetadataList{}
	err := d.get(ctx, metadataAPIBase, options, metadata, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) AddMetadata(ctx context.Context, metadata *Metadata) (*Metadata, error) {
	out := &Metadata{}
	err := d.post(ctx, metadataAPIBase, metadata, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) UpdateMetadata(ctx context.Context, metadata *Metadata) (*Metadata, error) {
	if metadata.Id == 0 && metadata.Product == 0 {
		return nil, errors.New("[defectdojo/UpdateMetadata] error: cannot update product with blank ids")
	}

	out := &Metadata{}
	err := d.patch(ctx, fmt.Sprintf("%s%v/", metadataAPIBase, metadata.Id), metadata, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) RemoveMetadata(ctx context.Context, metadata *Metadata) error {

	return nil
}
