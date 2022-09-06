package defectdojo_test

import (
	"context"
	"testing"

	"github.com/ed-wp/go-defectdojo"
	"github.com/stretchr/testify/require"
)

func Test_GetProducts(t *testing.T) {

	mockAPIServer := NewMockAPIServer(t, "secret", 200, exampleGetProductResponse)
	config := defectdojo.APIConfig{
		Host:     mockAPIServer.URL,
		APIToken: "secret",
	}
	api := defectdojo.New(config)

	ctx := context.Background()
	prod := &defectdojo.Product{
		Name:        "name",
		Description: "description",
	}
	pgl, err := api.GetProducts(ctx, prod, defectdojo.DefaultRequestOptions)

	require.NoError(t, err)
	require.Len(t, pgl.Results, 1)
	require.Equal(t, 123, pgl.Results[0].Id)
	require.Equal(t, "name", pgl.Results[0].Name)
	require.Equal(t, "description", pgl.Results[0].Description)
	for i, prod := range pgl.Results {
		t.Logf("[%d] Id: %d Name: %s\n", i, prod.Id, prod.Name)
	}
	mockAPIServer.Close()
}

func Test_GetProducts400(t *testing.T) {

	mockAPIServer := NewMockAPIServer(t, "secret", 400, "{}")
	config := defectdojo.APIConfig{
		Host:     mockAPIServer.URL,
		APIToken: "secret",
	}
	api := defectdojo.New(config)

	ctx := context.Background()
	prod := &defectdojo.Product{
		Name:        "name",
		Description: "description",
	}
	_, err := api.GetProducts(ctx, prod, defectdojo.DefaultRequestOptions)

	require.NotNil(t, err)
	mockAPIServer.Close()
}
