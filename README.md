# go-defectdojo
Library to simplify interacting with the [DefectDojo API](https://github.com/DefectDojo/django-DefectDojo)

## API

API Config
```
type APIConfig struct {
	Host     string       // DefectDojo Server, https://example.org
	APIToken string       // DefectDojo V2 API Token
	Client   *http.Client // Optional, can provide a custom HTTP Client, defaults to http.DefaultClient
	Verbose  bool         // Prints requests and stack traces on API errors
}
```

Example:
```
config := &defectdojo.APIConfig{
    Host:     "https://demo.defectdojo.org",
    APIToken: "do-not-hardcode-me",
}
api := defectdojo.New(config)
```

### Pagination
DefectDojo API is paginated, so you must provide Offsets and Limits for any Get requests.

These offsets and limits as passed into a get function using the RequestOptions struct:
```
type RequestOptions struct {
	Offset int
	Limit  int
}
```

The API returns a pagination wrapper around the results:
```
type PaginatedList[L any] struct {
	Count    int    // Number of Results
	Next     string // URL to next set of results
	Previous string // URL to previous set of results
	Results  []*L
}
```

For example, if you'd like to iterate the results and print them out:
```
for i, v := range pgl.Results {
    fmt.Printf("[%d] Result: %#v\n", i, v)
}
```

Optionally, for simple requests, you can use the default options:
```
var DefaultRequestOptions = &RequestOptions{
	Offset: 0,
	Limit:  100,
}
```

## Endpoint

```
endpoint := &defectdojo.Endpoint{
	Host:        "host",
	Product: 0,
}
```

### GetEndpoints

```
// Search for products matching provided Host and Product ID
pgl, err := api.GetEndpoints(ctx, endpoint, defectdojo.DefaultRequestOptions)

// check for errors before using the result
if err != nil {
    ...
}
```


## Engagement

```
engagement := &defectdojo.Engagement{
	Name:        "name",
	Description: "description",
}
```

### GetEngagements

```
// Search for engagements matching provided Name and Description
pgl, err := api.GetEngagements(ctx, engagement, defectdojo.DefaultRequestOptions)

// check for errors before using the result
if err != nil {
    ...
}
```

## Finding

```
finding := &defectdojo.Finding{
	Title:        "title",
	Description: "description",
}
```

### GetFindings

```
// Search for findings matching provided Title and Description
pgl, err := api.GetFindings(ctx, finding, defectdojo.DefaultRequestOptions)

// check for errors before using the result
if err != nil {
    ...
}
```

## ImportScan

```
scan := &defectdojo.ImportScan{
	ScanType:   "scanType",
	Engagement: 123,
	Tags:       []string{"tag1", "tag2"},
}
```

### ImportScan

```
// Import a scan with a scan report (json, sarif, etc)
scan, err := api.ImportScan(ctx, scan, scanReport)
```

## Metadata

```
metadata := &defectdojo.Metadata{
	Name:    "name",
	Value:   "value",
}
```

### GetMetadatas

```
// Search for metadata matching provided Name and Value
pgl, err := api.GetMetadatas(ctx, metadata, defectdojo.DefaultRequestOptions)

// check for errors before using the result
if err != nil {
    ...
}
```

## Product

```
product := &defectdojo.Product{
	Name:        "name",
	Description: "description",
}
```

### GetProducts

```
// Search for products matching provided Name and Description
pgl, err := api.GetProducts(ctx, product, defectdojo.DefaultRequestOptions)

// check for errors before using the result
if err != nil {
    ...
}
```

## ProductType

```
productType := &defectdojo.ProductType{
	Name:        "name",
	Description: "description",
}
```

### GetProductTypes

```
// Search for product types matching provided Name and Description
pgl, err := api.GetProductTypes(ctx, product, defectdojo.DefaultRequestOptions)

// check for errors before using the result
if err != nil {
    ...
}
```

## User

Not yet supported
