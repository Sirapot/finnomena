package coin

type Response struct {
	Status bool    `json:"status"`
	Error  string  `json:"error"`
	Data   []Stock `json:"data"`
}

type Stock struct {
	MasterId         string  `json:"mstar_id"`
	ThailandFundCode string  `json:"thailand_fund_code"`
	NavReturn        float64 `json:"nav_return"`
	Nav              float64 `json:"nav"`
	NavDate          string  `json:"nav_date"`
	AvgReturn        float32 `json:"avg_return"`
}
