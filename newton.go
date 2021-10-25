package newton

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dhiaayachi/go-newton-co/query"
)

const baseUrl = "https://api.newton.co/v1"

type Newton struct {
	clientId     string
	clientSecret string
}

type NewOrderReq struct {
	OrderType   string  `json:"order_type"`
	TimeInForce string  `json:"time_in_force"`
	Side        string  `json:"side"`
	Symbol      string  `json:"symbol"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"quantity"`
}

type GetTickSizesResp struct {
	Ticks map[string]struct {
		Tick float64 `json:"tick"`
	}
}

type GetMaxTradeAmountsResp struct {
	TradeAmounts map[string]struct {
		Buy  float64 `json:"buy"`
		Sell float64 `json:"sell"`
	}
}

type GetApplicableFeesResp struct {
	Fees struct {
		Maker float64 `json:"maker"`
		Taker float64 `json:"taker"`
	}
}

type GetMinTradeAmountsResp GetMaxTradeAmountsResp

type GetSymbolsResp struct {
	Symbols []string
}

type BalancesResp struct {
	Balances map[string]float64
}

type OpenOrdersResp struct {
	OpenOrders []struct {
		OrderID      string    `json:"order_id"`
		Symbol       string    `json:"symbol"`
		Quantity     int       `json:"quantity"`
		Price        float64   `json:"price"`
		DateCreated  time.Time `json:"date_created"`
		OrderType    string    `json:"order_type"`
		TimeInForce  string    `json:"time_in_force"`
		Side         string    `json:"side"`
		QuantityLeft float64   `json:"quantity_left"`
		ExpiryTime   time.Time `json:"expiry_time"`
	}
}

type OrderHistoryResp struct {
	OrdersHistory []struct {
		OrderID      string    `json:"order_id"`
		Symbol       string    `json:"symbol"`
		Quantity     int       `json:"quantity"`
		Price        float64   `json:"price"`
		Status       string    `json:"status"`
		DateCreated  time.Time `json:"date_created"`
		DateExecuted string    `json:"date_executed"`
		OrderType    string    `json:"order_type"`
		TimeInForce  string    `json:"time_in_force"`
		Side         string    `json:"side"`
		ExpiryTime   time.Time `json:"expiry_time,omitempty"`
	}
}

type ActionsResp struct {
	Actions []struct {
		Type        string    `json:"type"`
		Asset       string    `json:"asset"`
		Quantity    float64   `json:"quantity"`
		Status      string    `json:"status"`
		DateCreated time.Time `json:"date_created"`
		Price       float64   `json:"price,omitempty"`
	}
}

func New(ClientId string, ClientSecret string) *Newton {
	return &Newton{ClientId, ClientSecret}
}

func (n *Newton) sign(req *http.Request) error {
	var (
		method      string
		contentType string
		apiPath     string
		hashedBody  string
		currentTime string
	)

	method = req.Method
	if req.Method != http.MethodGet {
		contentType = "application/json"
	}

	apiPath = req.URL.Path
	if req.Body != http.NoBody {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return fmt.Errorf("read all body: %w", err)
		}

		hash := sha256.Sum256(b)
		hashedBody = fmt.Sprintf("%x", hash[:])

		req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	}

	currentTime = strconv.FormatInt(time.Now().Unix(), 10)

	toJoin := []string{
		method,
		contentType,
		apiPath,
		hashedBody,
		currentTime,
	}

	raw := strings.Join(toJoin, ":")

	mac := hmac.New(sha256.New, []byte(n.clientSecret))
	if _, err := mac.Write([]byte(raw)); err != nil {
		return fmt.Errorf("mac write: %w", err)
	}

	signedBase64 := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	req.Header.Add("NewtonDate", currentTime)
	req.Header.Add("NewtonAPIAuth", n.clientId+":"+signedBase64)

	return nil
}

func (n *Newton) doQuery(path string, method string, parameters []query.Parameter, isPublic bool, body string) (*http.Response, error) {
	url := baseUrl + path

	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	q := req.URL.Query()
	for _, a := range parameters {
		q.Add(a.Key, a.Value)
	}
	req.URL.RawQuery = q.Encode()
	if method != http.MethodGet {
		req.Header.Add("content-type", "application/json")
	}

	if !isPublic {
		err := n.sign(req)
		if err != nil {
			return nil, err
		}
	}

	res, err := http.DefaultClient.Do(req)

	return res, err
}

func (n *Newton) parseResponse(res *http.Response) ([]byte, error) {
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed :: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed :: %d :: %s", res.StatusCode, body)
	}

	return body, nil
}

// Public API
///////////////////////////////////////////////////////////////////////////////////////////////////
func (n *Newton) TickSizes() (*GetTickSizesResp, error) {
	query := &query.TickSizes{}
	res, err := n.doQuery("/order/tick-sizes", http.MethodGet, query.GetParameters(), query.IsPublic(), "")
	if err != nil {
		return nil, err
	}
	
	body, err := n.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var resp GetTickSizesResp
	err = json.Unmarshal(body, &resp.Ticks)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (n *Newton) MaximumTradeAmounts() (*GetMaxTradeAmountsResp, error) {
	res, err := n.doQuery("/order/maximums", http.MethodGet, []query.Parameter{}, true, "")
	if err != nil {
		return nil, err
	}

	body, err := n.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var resp GetMaxTradeAmountsResp
	err = json.Unmarshal(body, &resp.TradeAmounts)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (n *Newton) ApplicableFees() (*GetApplicableFeesResp, error) {
	res, err := n.doQuery("/fees", http.MethodGet, []query.Parameter{}, true, "")
	if err != nil {
		return nil, err
	}

	body, err := n.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var resp GetApplicableFeesResp
	err = json.Unmarshal(body, &resp.Fees)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (n *Newton) Symbols(query query.Query) (*GetSymbolsResp, error) {
	res, err := n.doQuery("/symbols", http.MethodGet, query.GetParameters(), true, "")
	if err != nil {
		return nil, err
	}

	body, err := n.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var resp GetSymbolsResp
	err = json.Unmarshal(body, &resp.Symbols)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (n *Newton) HealthCheck() error {
	res, err := n.doQuery("/health-check", http.MethodGet, []query.Parameter{}, true, "")
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed :: %d", res.StatusCode)
	}

	return nil
}

func (n *Newton) MinimumTradeAmount() (*GetMinTradeAmountsResp, error) {
	res, err := n.doQuery("/order/minimums", http.MethodGet, []query.Parameter{}, true, "")
	if err != nil {
		return nil, err
	}

	body, err := n.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var resp GetMinTradeAmountsResp
	err = json.Unmarshal(body, &resp.TradeAmounts)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Private API
///////////////////////////////////////////////////////////////////////////////////////////////////
func (n *Newton) Balances(query query.Query) (*BalancesResp, error) {
	res, err := n.doQuery("/balances", http.MethodGet, query.GetParameters(), query.IsPublic(), "")
	if err != nil {
		return nil, err
	}

	body, err := n.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var b BalancesResp
	err = json.Unmarshal(body, &b.Balances)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (n *Newton) Actions(query query.Query) (*ActionsResp, error) {
	res, err := n.doQuery("/actions", http.MethodGet, query.GetParameters(), false, "")
	if err != nil {
		return nil, err
	}

	body, err := n.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var r ActionsResp
	err = json.Unmarshal(body, &r.Actions)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (n *Newton) OrderHistory(query query.Query) (*OrderHistoryResp, error) {
	res, err := n.doQuery("/order/history", http.MethodGet, query.GetParameters(), false, "")
	if err != nil {
		return nil, err
	}

	body, err := n.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var r OrderHistoryResp
	err = json.Unmarshal(body, &r.OrdersHistory)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (n *Newton) OpenOrders(query query.Query) (*OpenOrdersResp, error) {
	res, err := n.doQuery("/order/history", http.MethodGet, query.GetParameters(), false, "")
	if err != nil {
		return nil, err
	}

	body, err := n.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var r OpenOrdersResp
	err = json.Unmarshal(body, &r.OpenOrders)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (n *Newton) NewOrder(orderType string, timeInForce string, side string, symbol string, price float64, quantity float64) (*OpenOrdersResp, error) {

	order := NewOrderReq{orderType, timeInForce, side, symbol, price, quantity}

	b, err := json.Marshal(&order)
	if err != nil {
		return nil, err
	}

	res, err := n.doQuery("/order/new", http.MethodPost, nil, false, string(b))
	if err != nil {
		return nil, err
	}
	
	body, err := n.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var r OpenOrdersResp
	err = json.Unmarshal(body, &r.OpenOrders)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
