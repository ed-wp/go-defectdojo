package defectdojo_test

var exampleGetMetadataResponse = `{
	"count": 123,
	"next": "http://api.example.org/accounts/?offset=400&limit=100",
	"previous": "http://api.example.org/accounts/?offset=200&limit=100",
	"results": [
	  {
		"id": 123,
		"product": 456,
		"endpoint": 789,
		"finding": 111,
		"name": "metadataName",
		"value": "metadataValue"
	  }
	]
  }`
