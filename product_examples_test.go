package defectdojo_test

var exampleGetProductResponse = `{
	"count": 123,
	"next": "http://api.example.org/accounts/?offset=400&limit=100",
	"previous": "http://api.example.org/accounts/?offset=200&limit=100",
	"results": [
	  {
		"id": 123,
		"findings_count": 0,
		"findings_list": [
		  0
		],
		"tags": [
		  "string"
		],
		"product_meta": [
		  {
			"name": "string",
			"value": "string"
		  }
		],
		"name": "name",
		"description": "description",
		"created": "2021-08-05T23:53:00.069Z",
		"prod_numeric_grade": 0,
		"business_criticality": "very high",
		"platform": "web service",
		"lifecycle": "construction",
		"origin": "third party library",
		"user_records": 0,
		"revenue": "string",
		"external_audience": true,
		"internet_accessible": true,
		"enable_simple_risk_acceptance": true,
		"enable_full_risk_acceptance": true,
		"product_manager": 0,
		"technical_contact": 0,
		"team_manager": 0,
		"prod_type": 0,
		"members": [
		  0
		],
		"authorization_groups": [
		  0
		],
		"regulations": [
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
