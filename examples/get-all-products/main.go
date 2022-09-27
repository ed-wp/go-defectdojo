package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ed-wp/go-defectdojo"
)

/*
1. Go to https://defectdojo.github.io/django-DefectDojo/getting_started/demo/
2. Use credentials to login to https://demo.defectdojo.org
3. Get API Token from top right menu
4. export DEFECTDOJO_HOST=https://demo.defectdojo.org
5. export DEFECTDOJO_TOKEN=api-token-from-step-3-without-Token-prefix
6. go run main.go
*/

func main() {
	config := defectdojo.APIConfig{
		Host:     os.Getenv("DEFECTDOJO_HOST"),
		APIToken: os.Getenv("DEFECTDOJO_TOKEN"),
	}
	api := defectdojo.New(config)
	options := &defectdojo.RequestOptions{
		Offset: 0,
		Limit:  5,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	// search/filter
	p := &defectdojo.Product{}

	for options != nil {
		products, err := api.GetProducts(ctx, p, options)
		if err != nil {
			panic(err) // used for brevity
		}

		for _, p := range products.Results {
			fmt.Printf("[Id: %d] Name: %s\n", p.Id, p.Name)
		}

		if !products.HasNext() {
			break
		}

		// fetch new RequestOptions, e.g. update offset
		options, err = products.NextRequestOptions()
		if err != nil {
			panic(err) // used for brevity
		}
	}
}
