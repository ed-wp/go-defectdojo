package defectdojo

import (
	"context"
	"errors"
	"fmt"
)

const (
	findingAPIBase = "/api/v2/findings/"
)

type Finding struct {
	Id                      int      `json:"id,omitempty" url:"id,omitempty"`
	Test                    int      `json:"test,omitempty" url:"test,omitempty"`
	ThreadId                int      `json:"thread_id,omitempty" url:"thread_id,omitempty"`
	FoundBy                 []int    `json:"found_by,omitempty" url:"found_by,omitempty"`
	Url                     string   `json:"url,omitempty" url:"url,omitempty"`
	Tags                    []string `json:"tags,omitempty" url:"tags,omitempty"`
	PushToJira              bool     `json:"push_to_jira,omitempty" url:"push_to_jira,omitempty"`
	Title                   string   `json:"title,omitempty" url:"title,omitempty"`
	Date                    string   `json:"date,omitempty" url:"date,omitempty"`
	SlaStartDate            string   `json:"sla_start_date,omitempty" url:"sla_start_date,omitempty"`
	Cwe                     int      `json:"cwe,omitempty" url:"cwe,omitempty"`
	Cve                     string   `json:"cve,omitempty" url:"cve,omitempty"`
	Cvssv3                  string   `json:"cvssv3,omitempty" url:"cvssv3,omitempty"`
	Cvssv3Score             float32  `json:"cvssv3_score,omitempty" url:"cvssv3_score,omitempty"`
	Severity                string   `json:"severity" url:"severity"`
	Description             string   `json:"description" url:"description"`
	Mitigation              string   `json:"mitigation,omitempty" url:"mitigation,omitempty"`
	Impact                  string   `json:"impact,omitempty" url:"impact,omitempty"`
	StepsToReproduce        string   `json:"steps_to_reproduce,omitempty" url:"steps_to_reproduce,omitempty"`
	SeverityJustification   string   `json:"severity_justification,omitempty" url:"severity_justification,omitempty"`
	References              string   `json:"references,omitempty" url:"references,omitempty"`
	IsTemplate              bool     `json:"is_template,omitempty" url:"is_template,omitempty"`
	Active                  bool     `json:"active,omitempty" url:"active,omitempty"`
	Verified                bool     `json:"verified,omitempty" url:"verified,omitempty"`
	FalseP                  bool     `json:"false_p,omitempty" url:"false_p,omitempty"`
	Duplicate               bool     `json:"duplicate,omitempty" url:"duplicate,omitempty"`
	OutOfScope              bool     `json:"out_of_scope,omitempty" url:"out_of_scope,omitempty"`
	RiskAccepted            bool     `json:"risk_accepted,omitempty" url:"risk_accepted,omitempty"`
	UnderReview             bool     `json:"under_review,omitempty" url:"under_review,omitempty"`
	UnderDefectReview       bool     `json:"under_defect_review,omitempty" url:"under_defect_review,omitempty"`
	IsMitigated             bool     `json:"is_mitigated,omitempty" url:"is_mitigated,omitempty"`
	NumericalSeverity       string   `json:"numerical_severity" url:"numerical_severity"`
	Line                    int      `json:"line,omitempty" url:"line,omitempty"`
	FilePath                string   `json:"file_path,omitempty" url:"file_path,omitempty"`
	ComponentName           string   `json:"component_name,omitempty" url:"component_name,omitempty"`
	ComponentVersion        string   `json:"component_version,omitempty" url:"component_version,omitempty"`
	StaticFinding           bool     `json:"static_finding,omitempty" url:"static_finding,omitempty"`
	DynamicFinding          bool     `json:"dynamic_finding,omitempty" url:"dynamic_finding,omitempty"`
	UniqueIdFromTool        string   `json:"unique_id_from_tool,omitempty" url:"unique_id_from_tool,omitempty"`
	VulnIdFromTool          string   `json:"vuln_id_from_tool,omitempty" url:"vuln_id_from_tool,omitempty"`
	SastSourceObject        string   `json:"sast_source_object,omitempty" url:"sast_source_object,omitempty"`
	SastSinkObject          string   `json:"sast_sink_object,omitempty" url:"sast_sink_object,omitempty"`
	SastSourceLine          int      `json:"sast_source_line,omitempty" url:"sast_source_line,omitempty"`
	SastSourceFilePath      string   `json:"sast_source_file_path,omitempty" url:"sast_source_file_path,omitempty"`
	Nb_occurences           int      `json:"nb_occurences,omitempty" url:"nb_occurences,omitempty"`
	PublishDate             string   `json:"publish_date,omitempty" url:"publish_date,omitempty"`
	ReviewRequestedBy       int      `json:"review_requested_by,omitempty" url:"review_requested_by,omitempty"`
	DefectReviewRequestedBy int      `json:"defect_review_requested_by,omitempty" url:"defect_review_requested_by,omitempty"`
	SonarqubeIssue          int      `json:"sonarqube_issue,omitempty" url:"sonarqube_issue,omitempty"`
	Endpoints               []int    `json:"endpoints,omitempty" url:"endpoints,omitempty"`
	EndpointStatus          []int    `json:"endpoint_status,omitempty" url:"endpoint_status,omitempty"`
	Reviewers               []int    `json:"reviewers,omitempty" url:"reviewers,omitempty"`
}

type FindingGroup struct {
	Id        int       `json:"id,omitempty" url:"id,omitempty"`
	Name      string    `json:"name,omitempty" url:"name,omitempty"`
	Test      int       `json:"test,omitempty" url:"test,omitempty"`
	JiraIssue JiraIssue `json:"jira_issue,omitempty" url:"jira_issue,omitempty"`
}

type JiraIssue struct {
	Id           int    `json:"id,omitempty" url:"id,omitempty"`
	Url          string `json:"url,omitempty" url:"url,omitempty"`
	JiraId       string `json:"jira_id,omitempty" url:"jira_id,omitempty"`
	JiraKey      string `json:"jira_key,omitempty" url:"jira_key,omitempty"`
	JiraCreation string `json:"jira_creation,omitempty" url:"jira_creation,omitempty"`
	JiraChange   string `json:"jira_change,omitempty" url:"jira_change,omitempty"`
	JiraProject  int    `json:"jira_project,omitempty" url:"jira_project,omitempty"`
	Finding      int    `json:"finding,omitempty" url:"finding,omitempty"`
	Engagement   int    `json:"engagement,omitempty" url:"engagement,omitempty"`
	FindingGroup int    `json:"finding_group,omitempty" url:"finding_group,omitempty"`
}

func (d *DefectDojoAPI) GetFindings(ctx context.Context, finding *Finding, options *RequestOptions) (*PaginatedList[Finding], error) {
	if options.Limit == 0 {
		return nil, ErrorInvalidOptions
	}

	out := &PaginatedList[Finding]{}
	err := d.get(ctx, findingAPIBase, options, finding, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) AddFinding(ctx context.Context, finding *Finding) (*Finding, error) {
	out := &Finding{}
	err := d.post(ctx, findingAPIBase, finding, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) UpdateFinding(ctx context.Context, finding *Finding) (*Finding, error) {
	if finding.Id == 0 && finding.Test == 0 {
		return nil, errors.New("[defectdojo/UpdateFinding] error: cannot update finding with blank ids")
	}

	out := &Finding{}
	err := d.patch(ctx, fmt.Sprintf("%s%v/", findingAPIBase, finding.Id), finding, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) RemoveFinding(ctx context.Context, finding *Finding) error {
	return fmt.Errorf("not implemented")
}
