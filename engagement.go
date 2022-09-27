package defectdojo

import (
	"context"
	"errors"
	"fmt"
)

const (
	engagementAPIBase = "/api/v2/engagements/"
)

type Engagement struct {
	Id                         int      `json:"id,omitempty" url:"id,omitempty"`
	Tags                       []string `json:"tags,omitempty" url:"tags,omitempty"`
	Name                       string   `json:"name,omitempty" url:"name,omitempty"`
	Description                string   `json:"description,omitempty" url:"description,omitempty"`
	Version                    string   `json:"version,omitempty" url:"version,omitempty"`
	FirstContacted             string   `json:"first_contacted,omitempty" url:"first_contacted,omitempty"`
	TargetStart                string   `json:"target_start" url:"target_start"`
	TargetEnd                  string   `json:"target_end" url:"target_end"`
	Reason                     string   `json:"reason,omitempty" url:"reason,omitempty"`
	Tracker                    string   `json:"tracker,omitempty" url:"tracker,omitempty"`
	TestStrategy               string   `json:"test_strategy,omitempty" url:"test_strategy,omitempty"`
	ThreatModel                bool     `json:"threat_model,omitempty" url:"threat_model,omitempty"`
	APITest                    bool     `json:"api_test,omitempty" url:"api_test,omitempty"`
	PenTest                    bool     `json:"pen_test,omitempty" url:"pen_test,omitempty"`
	CheckList                  bool     `json:"check_list,omitempty" url:"check_list,omitempty"`
	Status                     string   `json:"status,omitempty" url:"status,omitempty"`
	EngagementType             string   `json:"engagement_type,omitempty" url:"engagement_type,omitempty"`
	BuildID                    string   `json:"build_id,omitempty" url:"build_id,omitempty"`
	CommitHash                 string   `json:"commit_hash,omitempty" url:"commit_hash,omitempty"`
	BranchTag                  string   `json:"branch_tag,omitempty" url:"branch_tag,omitempty"`
	SourceCodeManagementURI    string   `json:"source_code_management_uri,omitempty" url:"source_code_management_uri,omitempty"`
	DeduplicationOnEngagement  bool     `json:"deduplication_on_engagement,omitempty" url:"deduplication_on_engagement,omitempty"`
	Lead                       int      `json:"lead,omitempty" url:"lead,omitempty"`
	Requestor                  int      `json:"requester,omitempty" url:"requester,omitempty"`
	Preset                     int      `json:"preset,omitempty" url:"preset,omitempty"`
	ReportType                 int      `json:"report_type,omitempty" url:"report_type,omitempty"`
	Product                    int      `json:"product" url:"product"`
	BuildServer                int      `json:"build_server,omitempty" url:"build_server,omitempty"`
	SourceCodeManagementServer int      `json:"source_code_management_server,omitempty" url:"source_code_management_server,omitempty"`
	OrchestrationEngine        int      `json:"orchestration_engine,omitempty" url:"orchestration_engine,omitempty"`
	Notes                      []Note   `json:"notes,omitempty" url:"notes,omitempty"`
	Files                      []File   `json:"files,omitempty" url:"files,omitempty"`
	RiskAcceptance             []int    `json:"risk_acceptance,omitempty" url:"risk_acceptance,omitempty"`
}

type Note struct {
	Id       int           `json:"id,omitempty" url:"id,omitempty"`
	Author   User          `json:"author,omitempty" url:"author,omitempty"`
	Editor   User          `json:"editor,omitempty" url:"editor,omitempty"`
	History  []NoteHistory `json:"history,omitempty" url:"history,omitempty"`
	Entry    string        `json:"entry,omitempty" url:"entry,omitempty"`
	Date     string        `json:"date,omitempty" url:"date,omitempty"`
	Private  bool          `json:"private,omitempty" url:"private,omitempty"`
	Edited   bool          `json:"edited,omitempty" url:"edited,omitempty"`
	EditTime string        `json:"edit_time,omitempty" url:"edit_time,omitempty"`
	NoteType int           `json:"note_type,omitempty" url:"note_type,omitempty"`
}

type NoteHistory struct {
	Id            int    `json:"id,omitempty" url:"id,omitempty"`
	CurrentEditor User   `json:"current_editor,omitempty" url:"current_editor,omitempty"`
	Data          string `json:"data,omitempty" url:"data,omitempty"`
	Time          string `json:"time,omitempty" url:"time,omitempty"`
	NoteType      int    `json:"note_type,omitempty" url:"note_type,omitempty"`
}

type File struct {
	Id    int    `json:"id,omitempty" url:"id,omitempty"`
	File  string `json:"file,omitempty" url:"file,omitempty"`
	Title string `json:"title,omitempty" url:"title,omitempty"`
}

func (d *DefectDojoAPI) GetEngagements(ctx context.Context, engagement *Engagement, options *RequestOptions) (*PaginatedList[Engagement], error) {
	if options.Limit == 0 {
		return nil, ErrorInvalidOptions
	}

	out := &PaginatedList[Engagement]{}
	err := d.get(ctx, engagementAPIBase, options, engagement, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) AddEngagement(ctx context.Context, engagement *Engagement) (*Engagement, error) {
	out := &Engagement{}
	err := d.post(ctx, engagementAPIBase, engagement, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) UpdateEngagement(ctx context.Context, engagement *Engagement) (*Engagement, error) {
	if engagement.Id == 0 && engagement.Product == 0 {
		return nil, errors.New("[defectdojo/UpdateEngagement] error: cannot update product with blank ids")
	}

	out := &Engagement{}
	err := d.patch(ctx, fmt.Sprintf("%s%v/", engagementAPIBase, engagement.Id), engagement, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DefectDojoAPI) RemoveEngagement(ctx context.Context, engagement *Engagement) error {
	return fmt.Errorf("not implemented")
}
