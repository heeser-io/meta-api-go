package main

import (
	v1 "github.com/heeser-io/meta-api-go/v1"
)

func main() {
	c := v1.WithAPIKey("f107c987abd3fcf2d3bde73ac9b6f8d1e7ca68d2a534c67c")
	// res, err := c.Auth.GenerateToken(&v1.GenerateTokenParams{
	// 	UserID:    "4830e158-a3c3-4e63-872f-a01d2febc23e",
	// 	ProjectID: "763c76bf-f2d3-4b90-93bf-b20b63ab6709",
	// 	Auth: map[string]interface{}{
	// 		"groups": []string{"admin"},
	// 	},
	// 	Meta: map[string]interface{}{
	// 		"channel": []string{"a"},
	// 	},
	// })

	c.Event.Dispatch(&v1.DispatchEventParams{
		Data: v1.SocketData{
			Event: "test-event",
			Data: map[string]interface{}{
				"test": "hallo",
			},
			MetaSelector: []v1.MetaSelector{
				{
					Operator: v1.OP_INCLUDES,
					Path:     "auth.groups",
					Value:    "admin",
				},
			},
		},
	})
}
