# go-defectdojo
Library to simplify interacting with the [DefectDojo API](https://github.com/DefectDojo/django-DefectDojo)

## Configuration

API Config
```
type APIConfig struct {
	Host     string       // DefectDojo Server, ex: https://demo.defectdojo.org
	APIToken string       // DefectDojo V2 API Token
	Client   *http.Client // Optional, can provide a custom HTTP Client, defaults to http.DefaultClient
	Verbose  bool         // Prints requests and stack traces on API errors
}
```

See examples for usage. 

## Pagination
DefectDojo API is paginated, so you must provide Offsets and Limits for any Get requests.
Failure to do so may lead to truncated or missing results.

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

There are additional helper functions for PaginatedLists, 
`HasNext()` which tells you if there are additional results
and `NextRequestOptions()` which will generate the next RequestOptions 
with the correct limit/offset. 

You can check out the [get-all-products example](examples/get-all-products) for an approach to handling pagination.

For simple requests, you can use the default options:
```
var DefaultRequestOptions = &RequestOptions{
	Offset: 0,
	Limit:  100,
}
```

or your just create your own:
```
options := &RequestOptions{
	Offset: 0,
	Limit:  50,
}
```

But be careful modifying this as large limit sizes can tie up resources on the defectdojo server.


## Examples
Provided are two examples, [get-products](examples/get-products) and [get-all-products](examples/get-all-products) in the [examples](examples) folder.

### get-products
This example shows basic library use and the dangers of pagination.

### get-all-products
This example builds on the first, and shows an approach to handling pagination.

## Search / Filtering
The struct passed into any Get... function are search parameters.
The above examples, we use `defectdojo.Product{}`, which means we are searching for anything.

For instance:
```
p := &defectdojo.Product{
	Name: "name-i-am-looking-for",
}
products, err = api.GetProducts(ctx, p, options)
```
will return products with `name-i-am-looking-for` in their name.

## API Docs
This library supports [Endpoint](#endpoint), [Engagement](#engagement), [Finding](#finding), [Metadata](#metadata), [Product](#product), [ProductType](#producttype), and [Test](#test)
with `Get`, `Add`, and `Update`.

`Remove` functionality is not implemented yet.

### Endpoint

```
endpoint := &defectdojo.Endpoint{
	Host:        "host",
	Product: 0,
}
```

#### GetEndpoints

```
// Search for products matching provided Host and Product ID
pgl, err := api.GetEndpoints(ctx, endpoint, defectdojo.DefaultRequestOptions)

// Check for errors before using the result
if err != nil {
    ...
}
```

### Engagement

```
engagement := &defectdojo.Engagement{
	Name:        "name",
	Description: "description",
}
```

#### GetEngagements

```
// Search for engagements matching provided Name and Description
pgl, err := api.GetEngagements(ctx, engagement, defectdojo.DefaultRequestOptions)

// Check for errors before using the result
if err != nil {
    ...
}
```

### Finding

```
finding := &defectdojo.Finding{
	Title:        "title",
	Description: "description",
}
```

#### GetFindings

```
// Search for findings matching provided Title and Description
pgl, err := api.GetFindings(ctx, finding, defectdojo.DefaultRequestOptions)

// Check for errors before using the result
if err != nil {
    ...
}
```

### ImportScan

```
scan := &defectdojo.ImportScan{
	ScanType:   "scanType",
	Engagement: 123,
	Tags:       []string{"tag1", "tag2"},
}
```

#### ImportScan

```
// Import a scan with a scan report (json, sarif, etc)
scan, err := api.ImportScan(ctx, scan, scanReport)
```

### Metadata

```
metadata := &defectdojo.Metadata{
	Name:    "name",
	Value:   "value",
}
```

#### GetMetadatas

```
// Search for metadata matching provided Name and Value
pgl, err := api.GetMetadatas(ctx, metadata, options)

// Check for errors before using the result
if err != nil {
    ...
}
```

### Product

```
product := &defectdojo.Product{
	Name:        "name",
	Description: "description",
}
```

#### GetProducts

```
// Search for products matching provided Name and Description
products, err := api.GetProducts(ctx, product, options)

// Check for errors before using the result
if err != nil {
    ...
}
// Use products.Results 
```

### ProductType

```
productType := &defectdojo.ProductType{
	Name:        "name",
	Description: "description",
}
```

#### GetProductTypes

```
// Search for product types matching provided Name and Description
productTypes, err := api.GetProductTypes(ctx, productType, options)

// Check for errors before using the result
if err != nil {
    ...
}
// Use productTypes.Results
```

### Test
```
productType := &defectdojo.Test{
	Title:       "title",
	Description: "description",
}
```

#### GetTests

```
// Search for product types matching provided Name and Description
tests, err := api.GetTests(ctx, product, options)

// Check for errors before using the result
if err != nil {
    ...
}
// Use tests.Results
```

### User

Not yet implemented.
