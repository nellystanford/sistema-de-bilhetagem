package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponsePayload struct {
	Tenant     string  `json:"tenant"`
	ProductSKU string  `json:"product"`
	Price      float64 `json:"price"`
	UseUnit    string  `json:"use_unit"`
}

func main() {
	http.HandleFunc("/mock-endpoint", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := ResponsePayload{
			Tenant:     "102",
			ProductSKU: "product_sku",
			Price:      100,
			UseUnit:    "GB",
		}

		json.NewEncoder(w).Encode(response)
	})

	port := ":3000"
	fmt.Printf("mock server available on port %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("error while initializing server:", err)
	}
}
