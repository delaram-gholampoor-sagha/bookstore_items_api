package queries

// {
// 	"equals" : [
// 		{
// 			"field" : "status",
// 			"value" : "pending"
// 		},
// 		{
// 			"filed" : "seller",
// 			"value" : 1
// 		}
// 	]
// }

// {
// 	"any_equals" : [
// 		{
// 			"field" : "status",
// 			"value" : "pending"
// 		},
// 		{
// 			"filed" : "seller",
// 			"value" : 1
// 		}
// 	]
// }

type EsQuery struct {
	Equals []FieldValue `json:"equals"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
