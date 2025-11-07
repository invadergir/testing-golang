package main

type DataWrapper[T any] struct {
	Data T `json:"data"`
}

// type DataWrapper struct {
// 	Data TokenBalances `json:"data"`
// }

type TokenBalances struct {
	TokenBalances []TokenBalance `json:"tokenBalances"`
}

type TokenBalance struct {
	Amount     string `json:"amount"`
	UpdateDate string `json:"updateDate"`
}

// func Test_jsonParsing(t *testing.T) {
// 	var assert = makeAsserter(t)

// 	// var jsonString = `{"data": {
// 	// 	"tokenBalances": [ {
// 	// 			"amount" : "100",
// 	// 			"updateDate" : "today"
// 	// 		}, {
// 	// 			"amount" : "200",
// 	// 			"updateDate" : "tomorrow"
// 	// 		}
// 	// 	]
// 	// }}`

// 	var jsonString = `{"data": {
// 		"tokenBalances": [ {
// 				"amount" : "100"
// 			}, {
// 				"amount" : "200"
// 			}
// 		]
// 	}}`

// 	var data DataWrapper[TokenBalances] = DataWrapper[TokenBalances]{}
// 	//	var data DataWrapper = DataWrapper{}
// 	var err = json.NewDecoder(strings.NewReader(jsonString)).Decode(&data)
// 	assert.Equal(nil, err)
// 	assert.Equal(3, len(data.Data.TokenBalances))
// 	assert.Equal(data.Data.TokenBalances[0].Amount, "100")
// 	assert.Equal(data.Data.TokenBalances[1].Amount, "200")
// }
