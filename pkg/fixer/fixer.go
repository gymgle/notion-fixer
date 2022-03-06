package fixer

import (
	"context"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type GetRatesResponse struct {
	Success bool               `json:"success"`
	Rates   map[string]float64 `json:"rates"`
}

const uri = "http://data.fixer.io/api/latest"

func GetRates(ctx context.Context, accessKey, base string) (map[string]float64, error) {
	client := resty.New()

	req, err := client.R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"access_key": accessKey,
			"base":       base,
		}).
		SetResult(&GetRatesResponse{}).
		Get(uri)

	if err != nil {
		return nil, err
	}

	res := req.Result().(*GetRatesResponse)
	log.Println("url:", req.Request.URL)
	log.Println("result:", req.Request.Result)

	if res.Success {
		return res.Rates, nil
	}

	return nil, fmt.Errorf("fetch fixer error %#v", res)
}
