package extserv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetClientContract(tenantID string, product string) (Catalogue, error) {
	url := "http://localhost:3000/mock-endpoint"

	// should send tenant id and product sku as query string
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error while requesting client contract:", err)
		return Catalogue{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error while reading response body:", err)
		return Catalogue{}, err
	}

	var response Catalogue
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("error while decoding json:", err)
		return Catalogue{}, err
	}

	fmt.Println("client contract retrieved successfully:")
	fmt.Printf("tenant: %s\n", response.Tenant)
	fmt.Printf("product: %s\n", response.ProductSKU)
	fmt.Printf("price: %f\n", response.Price)
	fmt.Printf("product: %s\n", response.UseUnit)

	return response, nil
}
