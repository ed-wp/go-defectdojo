package defectdojo

import (
	"context"
	"errors"
	"fmt"
	"log"
)

const (
	productTypeAPIBase = "/api/v2/product_types/"
)

type ProductType struct {
	Id                  int    `json:"id,omitempty" url:"id,omitempty"`
	Name                string `json:"name,omitempty" url:"name,omitempty"`
	Description         string `json:"description,omitempty" url:"description,omitempty"`
	CriticalProduct     bool   `json:"critical_product,omitempty" url:"critical_product,omitempty"`
	KeyProduct          bool   `json:"key_product,omitempty" url:"key_product,omitempty"`
	Updated             string `json:"updated,omitempty" url:"updated,omitempty"`
	Created             string `json:"created,omitempty" url:"created,omitempty"`
	Members             []int  `json:"members,omitempty" url:"members,omitempty"`
	AuthorizationGroups []int  `json:"authorization_groups,omitempty" url:"authorization_groups,omitempty"`
}

type PaginatedProductTypeList struct {
	Count    int            // Number of Results
	Next     string         // URL to next set of results
	Previous string         // URL to previous set of results
	Results  []*ProductType // List of ProductType results
	//lint:ignore U1000 required field
	prefetch interface{} // Prefetch data, currently unsupported
}

func (d *DefectDojoAPI) GetProductTypes(ctx context.Context, productType *ProductType, options *RequestOptions) (*PaginatedProductTypeList, error) {
	out := &PaginatedProductTypeList{}
	err := d.get(ctx, productTypeAPIBase, options, productType, out)
	if err != nil {
		return nil, err
	}
	if d.verbose {
		s := "Product Types:"
		for _, t := range out.Results {
			s = s + " " + t.Name
		}
		log.Println(s)
	}
	return out, nil
}

func (d *DefectDojoAPI) AddProductType(ctx context.Context, productType *ProductType) (*ProductType, error) {
	out := &ProductType{}
	err := d.post(ctx, productTypeAPIBase, productType, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) UpdateProductType(ctx context.Context, productType *ProductType) (*ProductType, error) {
	if productType.Id == 0 && productType.Created == "" {
		return nil, errors.New("[defectdojo/UpdateProductType] error: cannot update product type with blank ids")
	}

	out := &ProductType{}
	err := d.patch(ctx, fmt.Sprintf("%s%v/", productTypeAPIBase, productType.Id), productType, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
