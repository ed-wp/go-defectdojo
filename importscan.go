package defectdojo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"reflect"
	"strconv"
	"strings"
)

const (
	importScanAPIBase = "/api/v2/import-scan/"
)

type ImportScan struct {
	ScanDate         string   `json:"scan_date,omitempty" url:"scan_date,omitempty"`
	MinimumSeverity  string   `json:"minimum_severity,omitempty" url:"minimum_severity,omitempty"`
	Active           bool     `json:"active,omitempty" url:"active,omitempty"`
	Verified         bool     `json:"verified,omitempty" url:"verified,omitempty"`
	ScanType         string   `json:"scan_type,omitempty" url:"scan_type,omitempty"`
	EndpointToAdd    int      `json:"endpoint_to_add,omitempty" url:"endpoint_to_add,omitempty"`
	File             string   `json:"file,omitempty" url:"file,omitempty"`
	Engagement       int      `json:"engagement,omitempty" url:"engagement,omitempty"`
	Lead             int      `json:"lead,omitempty" url:"lead,omitempty"`
	Tags             []string `json:"tags,omitempty" url:"tags,omitempty"`
	CloseOldFindings bool     `json:"close_old_findings,omitempty" url:"close_old_findings,omitempty"`
	PushToJira       bool     `json:"push_to_jira,omitempty" url:"push_to_jira,omitempty"`
	Environment      string   `json:"environment,omitempty" url:"environment,omitempty"`
	Version          string   `json:"version,omitempty" url:"version,omitempty"`
	BuildID          string   `json:"build_id,omitempty" url:"build_id,omitempty"`
	BranchTag        string   `json:"branch_tag,omitempty" url:"branch_tag,omitempty"`
	CommitHash       string   `json:"commit_hash,omitempty" url:"commit_hash,omitempty"`
	Test             int      `json:"test,omitempty" url:"test,omitempty"`
	GroupBy          string   `json:"group_by,omitempty" url:"group_by,omitempty"`
}

// taken from go standard library - src/mime/multipart/writer.go
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func (d *DefectDojoAPI) ImportScan(ctx context.Context, importScan *ImportScan, scanData string) (*ImportScan, error) {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	if importScan.File == "" {
		importScan.File = "scan"
	}

	partHeader := textproto.MIMEHeader{}
	partHeader.Set("Content-Disposition", "form-data; name=\"file\"; filename=\""+escapeQuotes(importScan.File)+"\"")
	partHeader.Set("Content-Type", "application/octet-stream")
	part, err := writer.CreatePart(partHeader)
	if err != nil {
		return nil, err
	}
	_, err = part.Write([]byte(scanData))
	if err != nil {
		return nil, err
	}

	// Iterate the fields in the struct
	v := reflect.ValueOf(*importScan)
	vt := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// Check if the struct value is empty per go stdlib json encoding rules
		if !isEmptyValue(v.Field(i)) {
			partHeader = textproto.MIMEHeader{}
			// Take the json field from the struct
			k := strings.Split(vt.Field(i).Tag.Get("json"), ",")[0]
			partHeader.Set("Content-Disposition", "form-data; name=\""+k+"\"")
			part, err = writer.CreatePart(partHeader)
			if err != nil {
				return nil, err
			}

			vv := v.Field(i)
			vvt := v.Field(i).Type()
			switch vvt.Kind() {
			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.String:
				_, err = part.Write([]byte(fmt.Sprintf("%v", vv.Interface())))
				if err != nil {
					return nil, err
				}
			case reflect.Array, reflect.Slice:
				// The API expects separate parts for multiple values
				// For example sending =1,2,3 ends up with ["1,2,3"] not ["1", "2", "3"]
				// We need to send tag=1, tag=2, tag=3 all in separate parts

				// iterate the array or slice
				n := vv.Len()
				for i := 0; i < n; i++ {
					if i > 0 {
						// if we have more than 1 item, we will need a separate header
						// e.g. --boundary\nContent-Disposition: form-data; name=...
						part, err = writer.CreatePart(partHeader)
						if err != nil {
							return nil, err
						}
					}
					_, err = part.Write([]byte(fmt.Sprintf("%v", vv.Index(i))))
					if err != nil {
						return nil, err
					}
				}
			default:
				return nil, errors.New("unsupported type (" + vvt.Kind().String() + ") in struct")
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.host+importScanAPIBase, buf)
	if err != nil {
		return nil, err
	}
	// Multipart content type must include the boundary between each part
	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary())
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", fmt.Sprintf("%d", buf.Len()))

	// Add authorization header
	req.Header.Add("Authorization", "Token "+d.apiToken)
	res, err := d.client.Do(req)
	req.Header.Set("Authorization", "Token <redacted>")
	if err != nil {
		d.errorWithStackTrace(err)
		return nil, err
	}

	// POST will return 201 on success, non-201 on error
	if res.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(res.Body)
		err = errors.New("unexpected HTTP Response. URL: " + res.Request.URL.String() +
			" Expected: " + strconv.Itoa(http.StatusCreated) +
			" Got: " + strconv.Itoa(res.StatusCode) +
			" Body: " + string(b))
		d.errorWithStackTrace(err)
		return nil, err
	}

	// API will respond with an ImportScan json object
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// Unmarshal into an importscan struct
	out := &ImportScan{}
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// taken from go standard library - src/encoding/json/encode.go
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
