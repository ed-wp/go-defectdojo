package defectdojo_test

import (
	"context"
	"testing"

	"github.com/ed-wp/go-defectdojo"
	"github.com/stretchr/testify/require"
)

func Test_GetProductTypes(t *testing.T) {
	mockAPIServer := NewMockAPIServer(t, "secret", 200, exampleGetProductTypeResponse)
	config := defectdojo.APIConfig{
		Host:     mockAPIServer.URL,
		APIToken: "secret",
	}
	api := defectdojo.New(config)

	ctx := context.Background()
	prod := &defectdojo.ProductType{
		Name:        "productTypeName",
		Description: "productTypeDescription",
	}
	pgl, err := api.GetProductTypes(ctx, prod, defectdojo.DefaultRequestOptions)

	require.NoError(t, err)
	require.Len(t, pgl.Results, 2)
	require.Equal(t, 123, pgl.Results[0].Id)
	require.Equal(t, "productTypeName", pgl.Results[0].Name)
	require.Equal(t, "productTypeDescription", pgl.Results[0].Description)
	require.Equal(t, 124, pgl.Results[1].Id)
	require.Equal(t, "productTypeName2", pgl.Results[1].Name)
	require.Equal(t, "productTypeDescription2", pgl.Results[1].Description)
	for i, prod := range pgl.Results {
		t.Logf("[%d] Id: %d Name: %s\n", i, prod.Id, prod.Name)
	}
	mockAPIServer.Close()
}

func Test_GetProductTypes400(t *testing.T) {

	mockAPIServer := NewMockAPIServer(t, "secret", 400, "{}")
	config := defectdojo.APIConfig{
		Host:     mockAPIServer.URL,
		APIToken: "secret",
	}
	api := defectdojo.New(config)

	ctx := context.Background()
	prod := &defectdojo.ProductType{
		Name:        "productTypeName",
		Description: "productTypeDescription",
	}
	_, err := api.GetProductTypes(ctx, prod, defectdojo.DefaultRequestOptions)

	require.NotNil(t, err)
	mockAPIServer.Close()
}
