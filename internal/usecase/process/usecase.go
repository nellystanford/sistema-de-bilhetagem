package process

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/nellystanford/sistema-de-bilhetagem/internal/entity"
	extserv "github.com/nellystanford/sistema-de-bilhetagem/internal/extserv/catalogue"
)

var usageUnit map[string]int = map[string]int{
	"TB": 12,
	"GB": 9,
	"MB": 6,
	"KB": 3,
}

func ProcessMessage(input Input) (entity.TotalCost, error) {
	clientContract, err := extserv.GetClientContract(input.TenantID, input.Product)
	if err != nil {
		return entity.TotalCost{}, err
	}

	spentValue, err := CalculateSpentValue(input, clientContract)
	if err != nil {
		return entity.TotalCost{}, fmt.Errorf("error calculating spent value difference: %w", err)
	}

	return entity.TotalCost{
		Tenant:      input.TenantID,
		SpentAmount: spentValue,
		Product:     input.Product,
		Date:        time.Now(),
	}, nil
}

func CalculateSpentValue(input Input, clientContract extserv.Catalogue) (float64, error) {
	usedAmount, err := strconv.ParseFloat(input.UsedAmount, 64)
	if err != nil {
		fmt.Println("invalid used amount: ", input.UsedAmount)
		return 0, fmt.Errorf("error converting used amount to float value: %w", err)
	}

	convertedAmount, err := convertUnit(usedAmount, input.UseUnity, clientContract.UseUnit)
	if err != nil {
		return 0, fmt.Errorf("error calculating power difference: %w", err)
	}

	return clientContract.Price * convertedAmount, nil
}

func convertUnit(value float64, fromUnit string, toUnit string) (float64, error) {
	fromPower, exists := usageUnit[fromUnit]
	if !exists {
		return 0, errors.New("invalid input unit provided")
	}

	toPower, exists := usageUnit[toUnit]
	if !exists {
		return 0, errors.New("invalid output unit provided")
	}

	powerDifference := fromPower - toPower

	scaledValue := value * math.Pow(10, float64(powerDifference))
	return scaledValue, nil
}
