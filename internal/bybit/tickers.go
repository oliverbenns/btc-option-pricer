package bybit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type GetTickersProps struct {
	Category string
	Symbol   string
	BaseCoin string
	ExpDate  string
}

type GetTickersResult struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     GetTickersResultResult `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type GetTickersResultResult struct {
	Category string                         `json:"category"`
	List     []GetTickersResultResultTicker `json:"list"`
}

type GetTickersResultResultTicker struct {
	Symbol          string  `json:"symbol"`
	Bid1Price       string  `json:"bid1Price"`
	Bid1Size        string  `json:"bid1Size"`
	Ask1Price       string  `json:"ask1Price"`
	Ask1Size        string  `json:"ask1Size"`
	LastPrice       string  `json:"lastPrice"`
	PrevPrice24h    string  `json:"prevPrice24h"`
	Price24hPcnt    string  `json:"price24hPcnt"`
	HighPrice24h    string  `json:"highPrice24h"`
	LowPrice24h     string  `json:"lowPrice24h"`
	USDIndexPrice   *string `json:"usdIndexPrice,omitempty"`
	MarkPrice       *string `json:"markPrice,omitempty"`
	IndexPrice      *string `json:"indexPrice,omitempty"`
	MarkIv          *string `json:"markIv,omitempty"`
	UnderlyingPrice *string `json:"underlyingPrice,omitempty"`
	OpenInterest    *string `json:"openInterest,omitempty"`
	Turnover24h     *string `json:"turnover24h,omitempty"`
	Volume24h       *string `json:"volume24h,omitempty"`
	TotalVolume     *string `json:"totalVolume,omitempty"`
	TotalTurnover   *string `json:"totalTurnover,omitempty"`
	Delta           *string `json:"delta,omitempty"`
	Gamma           *string `json:"gamma,omitempty"`
	Vega            *string `json:"vega,omitempty"`
}

func buildQuery(props *GetTickersProps) string {
	parts := []string{}

	if props.Category != "" {
		parts = append(parts, fmt.Sprintf("category=%s", props.Category))
	}
	if props.Symbol != "" {
		parts = append(parts, fmt.Sprintf("symbol=%s", props.Symbol))
	}
	if props.BaseCoin != "" {
		parts = append(parts, fmt.Sprintf("baseCoin=%s", props.BaseCoin))
	}
	if props.ExpDate != "" {
		parts = append(parts, fmt.Sprintf("expDate=%s", props.ExpDate))
	}

	query := strings.Join(parts, "&")

	if query == "" {
		return ""
	}

	return "?" + query
}

func (c *Client) GetTickers(props *GetTickersProps) (*GetTickersResult, error) {
	path := buildQuery(props)
	url := fmt.Sprintf("%s/v5/market/tickers%s", c.baseURL, path)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var priceResp GetTickersResult
	if err := json.NewDecoder(response.Body).Decode(&priceResp); err != nil {
		return nil, err
	}

	return &priceResp, nil
}
