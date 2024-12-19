package contract

type ConsumptionMessage struct {
	TenantID   string `json:"tenant_id"`
	Product    string `json:"product"`
	UsedAmount string `json:"used_amount"`
	UseUnity   string `json:"use_unity"`
}
