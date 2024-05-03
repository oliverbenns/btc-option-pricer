package bybit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type GetHistoricalVolatilityProps struct {
	Category  string
	BaseCoin  string
	Period    int
	StartTime int
	EndTime   int
}

type GetHistoricalVolatilityResult struct {
	RetCode  int                                   `json:"retCode"`
	RetMsg   string                                `json:"retMsg"`
	Category string                                `json:"category"`
	Result   []GetHistoricalVolatilityResultResult `json:"result"`
}

type GetHistoricalVolatilityResultResult struct {
	Period int    `json:"period"`
	Value  string `json:"value"`
	Time   string `json:"time"`
}

func (c *Client) GetHistoricalVolatility(props *GetHistoricalVolatilityProps) (*GetHistoricalVolatilityResult, error) {
	path := buildVolatilityQuery(props)
	url := fmt.Sprintf("%s/v5/market/historical-volatility%s", c.baseURL, path)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var resp GetHistoricalVolatilityResult
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func buildVolatilityQuery(props *GetHistoricalVolatilityProps) string {
	parts := []string{}

	if props.Category != "" {
		parts = append(parts, fmt.Sprintf("category=%s", props.Category))
	}
	if props.BaseCoin != "" {
		parts = append(parts, fmt.Sprintf("baseCoin=%s", props.BaseCoin))
	}
	if props.Period != 0 {
		parts = append(parts, fmt.Sprintf("period=%d", props.Period))
	}
	if props.StartTime != 0 {
		parts = append(parts, fmt.Sprintf("startTime=%d", props.StartTime))
	}
	if props.EndTime != 0 {
		parts = append(parts, fmt.Sprintf("endTime=%d", props.EndTime))
	}

	query := strings.Join(parts, "&")

	if query == "" {
		return ""
	}

	return "?" + query
}
