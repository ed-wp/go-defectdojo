package defectdojo

import (
	"context"
	"errors"
	"fmt"
)

const (
	testAPIBase = "/api/v2/tests/"
)

type Test struct {
	Id              int            `json:"id,omitempty" url:"id,omitempty"`
	Tags            []string       `json:"tags,omitempty" url:"tags,omitempty"`
	TestTypeName    string         `json:"test_type_name,omitempty" url:"test_type_name,omitempty"`
	FindingGroups   []FindingGroup `json:"finding_groups,omitempty" url:"finding_groups,omitempty"`
	Title           string         `json:"title,omitempty" url:"title,omitempty"`
	Description     string         `json:"description,omitempty" url:"description,omitempty"`
	TargetStart     string         `json:"target_start,omitempty" url:"target_start,omitempty"`
	TargetEnd       string         `json:"target_end,omitempty" url:"target_end,omitempty"`
	EstimatedTime   string         `json:"estimated_time,omitempty" url:"estimated_time,omitempty"`
	ActualTime      string         `json:"actual_time,omitempty" url:"actual_time,omitempty"`
	PercentComplete int            `json:"percent_complete,omitempty" url:"percent_complete,omitempty"`
	Updated         string         `json:"updated,omitempty" url:"updated,omitempty"`
	Created         string         `json:"created,omitempty" url:"created,omitempty"`
	Version         string         `json:"version,omitempty" url:"version,omitempty"`
	BuildId         string         `json:"build_id,omitempty" url:"build_id,omitempty"`
	CommitHash      string         `json:"commit_hash,omitempty" url:"commit_hash,omitempty"`
	BranchTag       string         `json:"branch_tag,omitempty" url:"branch_tag,omitempty"`
	Engagement      int            `json:"engagement,omitempty" url:"engagement,omitempty"`
	Lead            int            `json:"lead,omitempty" url:"lead,omitempty"`
	TestType        int            `json:"test_type,omitempty" url:"test_type,omitempty"`
	Environment     int            `json:"environment,omitempty" url:"environment,omitempty"`
	Notes           []Note         `json:"notes,omitempty" url:"notes,omitempty"`
	Files           []File         `json:"files,omitempty" url:"files,omitempty"`
}

type PaginatedTestList struct {
	Count    int     // Number of Results
	Next     string  // URL to next set of results
	Previous string  // URL to previous set of results
	Results  []*Test // List of Test results
	//lint:ignore U1000 required field
	prefetch interface{} // Prefetch data, currently unsupported
}

func (d *DefectDojoAPI) GetTests(ctx context.Context, test *Test, options *RequestOptions) (*PaginatedTestList, error) {
	out := &PaginatedTestList{}
	err := d.get(ctx, testAPIBase, options, test, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) AddTest(ctx context.Context, test *Test) (*Test, error) {
	out := &Test{}
	err := d.post(ctx, testAPIBase, test, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) UpdateTest(ctx context.Context, test *Test) (*Test, error) {
	if test.Id == 0 && test.Created == "" {
		return nil, errors.New("[defectdojo/UpdateTest] error: cannot update product with blank ids")
	}

	out := &Test{}
	err := d.patch(ctx, fmt.Sprintf("%s%v/", testAPIBase, test.Id), test, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) RemoveTest(ctx context.Context, test *Test) error {

	return nil
}
