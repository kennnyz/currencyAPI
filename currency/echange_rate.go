package currency

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// получаем обменный курс с openexchangerates.org

type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}

func getRate(c Currency) float64 {
	resp, err := http.Get("https://openexchangerates.org/api/latest.json?app_id=922b16e1a6b54d13a7c3dda92baf95d1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var exchangeRates ExchangeRates
	err = json.Unmarshal(body, &exchangeRates)
	if err != nil {
		return 0
	}

	rateB := exchangeRates.Rates[c.CurrencyFrom]
	rateA := exchangeRates.Rates[c.CurrencyTo]

	fmt.Println("rateA:", rateA, "rateB:", rateB)

	return rateA / rateB
}
