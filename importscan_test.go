package defectdojo_test

import (
	"context"
	"testing"

	"github.com/ed-wp/go-defectdojo"
	"github.com/stretchr/testify/require"
)

func Test_ImportScan(t *testing.T) {

	mockAPIServer := NewMockAPIServer(t, "secret", 201, exampleImportScanResponse)
	config := defectdojo.APIConfig{
		Host:     mockAPIServer.URL,
		APIToken: "secret",
	}
	api := defectdojo.New(config)

	ctx := context.Background()
	scan := &defectdojo.ImportScan{
		ScanType:   "scanType",
		Engagement: 123,
		Tags:       []string{"tag1", "tag2"},
	}
	ret, err := api.ImportScan(ctx, scan, exampleImportScanGitLeaksReport)

	require.NoError(t, err)
	require.Equal(t, "scanType", ret.ScanType)
	require.Equal(t, 123, ret.Engagement)
	mockAPIServer.Close()
}

func Test_ImportScan400(t *testing.T) {

	mockAPIServer := NewMockAPIServer(t, "secret", 400, "{}")
	config := defectdojo.APIConfig{
		Host:     mockAPIServer.URL,
		APIToken: "secret",
	}
	api := defectdojo.New(config)

	ctx := context.Background()
	scan := &defectdojo.ImportScan{
		ScanType:   "scanType",
		Engagement: 123,
	}
	_, err := api.ImportScan(ctx, scan, exampleImportScanGitLeaksReport)

	require.NotNil(t, err)
	mockAPIServer.Close()
}
