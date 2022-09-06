package defectdojo_test

var exampleImportScanResponse = `{
	"scan_date": "2021-08-17",
	"minimum_severity": "Info",
	"active": true,
	"verified": true,
	"scan_type": "scanType",
	"endpoint_to_add": 0,
	"file": "string",
	"engagement": 123,
	"lead": 0,
	"tags": [
	  "string"
	],
	"close_old_findings": false,
	"push_to_jira": false,
	"environment": "string",
	"version": "string",
	"build_id": "string",
	"branch_tag": "string",
	"commit_hash": "string",
	"test": 0,
	"group_by": "component_name"
  }`

var exampleImportScanGitLeaksReport = `[
 {
  "line": "  api_key: example # Risk.",
  "offender": "api_key: example",
  "commit": "example",
  "repo": ".",
  "rule": "Generic Credential",
  "commitMessage": "Merge branch 'dev'\n",
  "author": "example",
  "email": "git@example.com",
  "file": "example/.example",
  "date": "2020-02-29T18:36:02-08:00",
  "tags": "key, API, generic"
 },
 {
  "line": "  api_key: example # Sensitive.",
  "offender": "api_key: example",
  "commit": "example",
  "repo": ".",
  "rule": "Generic Credential",
  "commitMessage": "Merge branch 'dev'\n",
  "author": "example",
  "email": "example@example.com",
  "file": "example/.example",
  "date": "2020-02-29T18:36:02-08:00",
  "tags": "key, API, generic"
 }
]`
