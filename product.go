package defectdojo

import (
	"context"
	"errors"
	"fmt"
	"regexp"
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

type PaginatedProductList struct {
	Count    int        // Number of Results
	Next     string     // URL to next set of results
	Previous string     // URL to previous set of results
	Results  []*Product // List of Product results
	//lint:ignore U1000 required field
	prefetch interface{} // Prefetch data, currently unsupported
}

func (d *DefectDojoAPI) GetProducts(ctx context.Context, product *Product, options *RequestOptions) (*PaginatedProductList, error) {
	out := &PaginatedProductList{}
	err := d.get(ctx, productAPIBase, options, product, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) GetProductByGitURL(ctx context.Context, gitURL string, gitURLMetaField string) (*Product, error) {
	product := &Product{}
	regex, err := regexp.Compile("[^/:]+/[^/]+$")
	if err != nil {
		return nil, err
	}
	match := regex.FindString(gitURL)
	if match == "" {
		return nil, fmt.Errorf("failed to extract org/repo from git url: %s", gitURL)
	}

	product.Name = match

	out := &PaginatedProductList{}
	err = d.get(ctx, productAPIBase, DefaultRequestOptions, product, out)
	if err != nil {
		return nil, err
	}
	for i, prod := range out.Results {
		for _, meta := range prod.ProductMeta {
			if meta.Name == gitURLMetaField && meta.Value == gitURL {
				return out.Results[i], nil
			}
		}
	}
	return nil, fmt.Errorf("no repositories had metadata field \"%s\" with a value of \"%s\"", gitURLMetaField, gitURL)
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

	return nil
}
