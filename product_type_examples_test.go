package defectdojo_test

var exampleGetProductTypeResponse = `{
	"count": 123,
	"next": "http://api.example.org/accounts/?offset=400&limit=100",
	"previous": "http://api.example.org/accounts/?offset=200&limit=100",
	"results": [
	  {
		"id": 123,
		"name": "productTypeName",
		"description": "productTypeDescription",
		"critical_product": true,
		"key_product": true,
		"updated": "2021-08-06T06:09:15.952Z",
		"created": "2021-08-06T06:09:15.952Z",
		"members": [
		  0
		],
		"authorization_groups": [
		  0
		],
		"prefetch": {
		  "authorization_groups": {
			"additionalProp1": {
			  "id": 0,
			  "name": "string",
			  "description": "string",
			  "users": [
				0
			  ],
			  "prefetch": {
				"product_groups": {},
				"product_type_groups": {}
			  }
		    }
		  }
	    }
      },
	  {
		"id": 124,
		"name": "productTypeName2",
		"description": "productTypeDescription2",
		"critical_product": true,
		"key_product": true,
		"updated": "2021-08-06T06:09:15.952Z",
		"created": "2021-08-06T06:09:15.952Z",
		"members": [
		  0
		],
		"authorization_groups": [
		  0
		],
		"prefetch": {
		  "authorization_groups": {
			"additionalProp1": {
			  "id": 0,
			  "name": "string",
			  "description": "string",
			  "users": [
				0
			  ],
			  "prefetch": {
				"product_groups": {},
				"product_type_groups": {}
			  }
		    }
		  }
	    }
      }
    ],
   "prefetch" : {}
}`
