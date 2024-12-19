package process_test

import (
	"testing"

	extserv "github.com/nellystanford/sistema-de-bilhetagem/internal/extserv/catalogue"
	"github.com/nellystanford/sistema-de-bilhetagem/internal/usecase/process"
	"github.com/stretchr/testify/assert"
)

func TestCalculateAmount(t *testing.T) {
	t.Run("Should calculate amount spent correctly from MB to GB", func(t *testing.T) {
		// Arrange
		input := process.Input{
			TenantID:   "102",
			Product:    "product_sku",
			UsedAmount: "100",
			UseUnity:   "MB",
		}

		clientContract := extserv.Catalogue{
			Tenant:     input.TenantID,
			ProductSKU: input.Product,
			Price:      100,
			UseUnit:    "GB",
		}

		// Act
		value, err := process.CalculateSpentValue(input, clientContract)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, float64(10), value)
	})

	t.Run("Should calculate amount spent correctly from MB to GB", func(t *testing.T) {
		// Arrange
		input := process.Input{
			TenantID:   "102",
			Product:    "product_sku",
			UsedAmount: "1",
			UseUnity:   "GB",
		}

		clientContract := extserv.Catalogue{
			Tenant:     input.TenantID,
			ProductSKU: input.Product,
			Price:      10,
			UseUnit:    "MB",
		}

		// Act
		value, err := process.CalculateSpentValue(input, clientContract)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, float64(100), value)
	})
}
