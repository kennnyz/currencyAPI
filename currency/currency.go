package currency

type Currency struct {
	CurrencyFrom string  `json:"currency_from,omitempty"`
	CurrencyTo   string  `json:"currency_to,omitempty"`
	ExchangeRate float64 `json:"exchange_rate,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}
