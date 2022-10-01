package ViewModels

type RuleViewModel struct {
	Routes      []RouteViewModel `json:"routes"`
	Airlines    []string         `json:"airlines"`
	Agencies    []string         `json:"agencies"`
	Suppliers   []string         `json:"suppliers"`
	AmountType  string           `json:"amountType"`
	AmountValue float64          `json:"amountValue"`
}
