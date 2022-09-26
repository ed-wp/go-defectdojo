package defectdojo_test

import (
	"context"
	"testing"

	"github.com/ed-wp/go-defectdojo"
	"github.com/stretchr/testify/require"
)

func Test_GetMetadata(t *testing.T) {

	mockAPIServer := NewMockAPIServer(t, "secret", 200, exampleGetMetadataResponse)
	config := defectdojo.APIConfig{
		Host:     mockAPIServer.URL,
		APIToken: "secret",
	}
	api := defectdojo.New(config)

	ctx := context.Background()
	meta := &defectdojo.Metadata{
		Name:  "metadataName",
		Value: "metadataValue",
	}
	pgl, err := api.GetMetadatas(ctx, meta, defectdojo.DefaultRequestOptions)

	require.NoError(t, err)
	require.Len(t, pgl.Results, 1)
	require.Equal(t, 123, pgl.Results[0].Id)
	require.Equal(t, "metadataName", pgl.Results[0].Name)
	require.Equal(t, "metadataValue", pgl.Results[0].Value)
	for i, meta := range pgl.Results {
		t.Logf("[%d] Id: %d Name: %s\n", i, meta.Id, meta.Name)
	}
	mockAPIServer.Close()
}

func Test_GetMetadata400(t *testing.T) {

	mockAPIServer := NewMockAPIServer(t, "secret", 400, "{}")
	config := defectdojo.APIConfig{
		Host:     mockAPIServer.URL,
		APIToken: "secret",
	}
	api := defectdojo.New(config)

	ctx := context.Background()
	meta := &defectdojo.Metadata{
		Name:  "metadataName",
		Value: "metadataValue",
	}
	_, err := api.GetMetadatas(ctx, meta, defectdojo.DefaultRequestOptions)

	require.NotNil(t, err)
	mockAPIServer.Close()
}
