package defectdojo

import (
	"context"
	"errors"
	"fmt"
)

const (
	productAPIBase = "/api/v2/products/"
)

type Product struct {
	Id                         int         `json:"id,omitempty" url:"id,omitempty"`
	Tags                       []string    `json:"tags,omitempty" url:"tags,omitempty"`
	Name                       string      `json:"name" url:"name"`
	Description                string      `json:"description,omitempty" url:"description,omitempty"`
	ProdNumericGrade           int         `json:"prod_numeric_grade,omitempty" url:"prod_numeric_grade,omitempty"`
	BusinessCriticality        string      `json:"business_criticality,omitempty" url:"business_criticality,omitempty"`
	Platform                   string      `json:"platform,omitempty" url:"platform,omitempty"`
	Lifecycle                  string      `json:"lifecycle,omitempty" url:"lifecycle,omitempty"`
	Origin                     string      `json:"origin,omitempty" url:"origin,omitempty"`
	UserRecords                int         `json:"user_records,omitempty" url:"user_records,omitempty"`
	Revenue                    string      `json:"revenue,omitempty" url:"revenue,omitempty"`
	ExternalAudience           bool        `json:"external_audience,omitempty" url:"external_audience,omitempty"`
	InternetAccessible         bool        `json:"internet_accessible,omitempty" url:"internet_accessible,omitempty"`
	EnableSimpleRiskAcceptance bool        `json:"enable_simple_risk_acceptance,omitempty" url:"enable_simple_risk_acceptance,omitempty"`
	EnableFullRiskAcceptance   bool        `json:"enable_full_risk_acceptance,omitempty" url:"enable_full_risk_acceptance,omitempty"`
	ProductManager             int         `json:"product_manager,omitempty" url:"product_manager,omitempty"`
	TechnicalContact           int         `json:"technical_contact,omitempty" url:"technical_contact,omitempty"`
	TeamManager                int         `json:"team_manager,omitempty" url:"team_manager,omitempty"`
	ProdType                   int         `json:"prod_type,omitempty" url:"prod_type,omitempty"`
	Regulations                []int       `json:"regulations,omitempty" url:"regulations,omitempty"`
	ProductMeta                []*Metadata `json:"product_meta,omitempty" url:"product_meta,omitempty"`
	Prefetch                   interface{} `json:"prefetch,omitempty" url:"prefetch,omitempty"`
}

func (d *DefectDojoAPI) GetProducts(ctx context.Context, product *Product, options *RequestOptions) (*PaginatedList[Product], error) {
	out := &PaginatedList[Product]{}
	err := d.get(ctx, productAPIBase, options, product, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// beta, using generics, use with caution
func (d *DefectDojoAPI) GetAllProducts(ctx context.Context, product *Product, options *RequestOptions) (*PaginatedList[Product], error) {
	out := &PaginatedList[Product]{}
	err := d.getAll(ctx, productAPIBase, options, product, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) AddProduct(ctx context.Context, product *Product) (*Product, error) {
	out := &Product{}
	err := d.post(ctx, productAPIBase, product, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) UpdateProduct(ctx context.Context, product *Product) (*Product, error) {
	if product.Id == 0 && product.ProdType == 0 {
		return nil, errors.New("[defectdojo/UpdateProduct] error: cannot update product with blank ids")
	}

	out := &Product{}
	err := d.patch(ctx, fmt.Sprintf("%s%v/", productAPIBase, product.Id), product, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) RemoveProduct(ctx context.Context, product *Product) error {
	return fmt.Errorf("not implemented")
}
