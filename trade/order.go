package trade

import (
	"encoding/json"

	"github.com/lourkeur/gokex/rest"
)

type OrderSpec struct {
	InstId        string `json:"instId"`
	TradeMode     string `json:"tdMode"`
	Currency      string `json:"ccy"`
	ClientOrderId string `json:"clOrdId"`
	Tag           string `json:"tag"`
	Side          string `json:"side"`
	PositionSide  string `json:"posSide"`
	OrderType     string `json:"ordType"`
	Quantity      string `json:"sz"`
	Price         string `json:"px"`
	ReduceOnly    bool   `json:"reduceOnly"`
	QuantityType  string `json:"tgtCcy"`
}

type OrderResult struct {
	OrderId       string `json:"ordId"`
	ClientOrderId string `json:"clOrdId"`
	Code          string `json:"sCode"`
	Msg           string `json:"sMsg"`
}

func Order(h rest.Handle, spec *OrderSpec) ([]OrderResult, error) {
	env, err := h.Post("trade/order", spec)
	if env != nil {
		res := make([]OrderResult, len(env.Data))
		for i := range env.Data {
			if err := json.Unmarshal(env.Data[i], &res[i]); err != nil {
				return nil, err
			}
		}
		return res, err
	}
	return nil, err
}
