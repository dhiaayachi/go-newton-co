package newton

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const baseUrl = "https://api.newton.co/v1"

type ActionType string

const (
	DEPOSIT    ActionType = "DEPOSIT"
	WITHDRAWAL ActionType = "WITHDRAWAL"
	TRANSACT   ActionType = "TRANSACT"
)

type Newton struct {
	clientId     string
	clientSecret string
}

type Args struct {
	Key   string
	Value string
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
		Tick       float64 `json:"tick"`
	}
}

type GetMaxTradeAmountsResp struct {
	MaxTradeAmounts map[string]struct {
		Buy       float64 `json:"buy"`
		Sell       float64 `json:"sell"`
	}
}

type GetApplicableFeesResp struct {
	Fees struct {
		Maker       float64 `json:"maker"`
		Taker       float64 `json:"taker"`
	}
}

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

type OrdersHistoryResp struct {
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

// Public API
///////////////////////////////////////////////////////////////////////////////////////////////////
func (n *Newton) doPublicQuery(path string, method string, args []Args, body string) (*http.Response, error) {
	url := baseUrl + path

	req, _ := http.NewRequest(method, url, nil)
	q := req.URL.Query()
	for _, a := range args {
		q.Add(a.Key, a.Value)
	}
	req.URL.RawQuery = q.Encode()
	if method != http.MethodGet {
		_, err := req.Body.Read([]byte(body))
		if err != nil {
			return nil, err
		}
		req.Header.Add("content-type", "application/json")
	}

	res, err := http.DefaultClient.Do(req)

	return res, err
}

func (n *Newton) GetTickSizes() (*GetTickSizesResp, error) {
	res, err := n.doPublicQuery("/order/tick-sizes", http.MethodGet, []Args{}, "")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request failed :: %d", res.StatusCode))
	}

	body, _ := ioutil.ReadAll(res.Body)

	var resp GetTickSizesResp
	err = json.Unmarshal(body, &resp.Ticks)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (n *Newton) GetMaximumTradeAmounts() (*GetMaxTradeAmountsResp, error) {
	res, err := n.doPublicQuery("/order/maximums", http.MethodGet, []Args{}, "")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request failed :: %d", res.StatusCode))
	}

	body, _ := ioutil.ReadAll(res.Body)

	var resp GetMaxTradeAmountsResp
	err = json.Unmarshal(body, &resp.MaxTradeAmounts)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (n *Newton) GetApplicableFees() (*GetApplicableFeesResp, error) {
	res, err := n.doPublicQuery("/fees", http.MethodGet, []Args{}, "")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request failed :: %d", res.StatusCode))
	}

	body, _ := ioutil.ReadAll(res.Body)

	var resp GetApplicableFeesResp
	err = json.Unmarshal(body, &resp.Fees)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (n *Newton) GetSymbols(baseAsset, quoteAsset string) (*GetSymbolsResp, error) {
	args := []Args{}
	if baseAsset != "" {
		args = append(args, Args{Key: "base_asset", Value: baseAsset})
	}

	if quoteAsset != "" {
		args = append(args, Args{Key: "quote_asset", Value: quoteAsset})
	}

	res, err := n.doPublicQuery("/symbols", http.MethodGet, args, "")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request failed :: %d", res.StatusCode))
	}

	body, _ := ioutil.ReadAll(res.Body)

	var resp GetSymbolsResp
	err = json.Unmarshal(body, &resp.Symbols)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Private API
///////////////////////////////////////////////////////////////////////////////////////////////////
func (n *Newton) doPrivateQuery(path string, method string, args []Args, body string) (*http.Response, error) {
	url := baseUrl + path

	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	q := req.URL.Query()
	for _, a := range args {
		q.Add(a.Key, a.Value)
	}
	req.URL.RawQuery = q.Encode()
	if method != http.MethodGet {
		req.Header.Add("content-type", "application/json")
	}
	err := n.sign(req)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)

	return res, err
}

func (n *Newton) Balances(asset string) (*BalancesResp, error) {

	a := make([]Args, 1)

	a[0].Key = "asset"
	a[0].Value = asset
	res, err := n.doPrivateQuery("/balances", http.MethodGet, a, "")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request failed :: %d", res.StatusCode))
	}

	body, _ := ioutil.ReadAll(res.Body)

	var b BalancesResp
	err = json.Unmarshal(body, &b.Balances)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (n *Newton) Actions(actionType ActionType, limit int, offset int, startDate int64, endDate int64) (*ActionsResp, error) {

	a := make([]Args, 5)

	a[0].Key = "action_type"
	a[0].Value = string(actionType)

	a[1].Key = "end_date"
	a[1].Value = strconv.FormatInt(endDate, 10)

	a[2].Key = "limit"
	a[2].Value = strconv.Itoa(limit)

	a[3].Key = "offset"
	a[3].Value = strconv.Itoa(offset)

	a[4].Key = "start_date"
	a[4].Value = strconv.FormatInt(startDate, 10)

	res, err := n.doPrivateQuery("/actions", http.MethodGet, a, "")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request failed :: %d", res.StatusCode))
	}

	var r ActionsResp
	err = json.Unmarshal(body, &r.Actions)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (n *Newton) OrdersHistory(limit int, offset int, startDate int64, endDate int64, symbol string, timeInForce string) (*OrdersHistoryResp, error) {

	a := make([]Args, 10)

	a[1].Key = "end_date"
	a[1].Value = strconv.FormatInt(endDate, 10)

	a[2].Key = "limit"
	a[2].Value = strconv.Itoa(limit)

	a[3].Key = "offset"
	a[3].Value = strconv.Itoa(offset)

	a[4].Key = "start_date"
	a[4].Value = strconv.FormatInt(startDate, 10)

	a[5].Key = "symbol"
	a[5].Value = symbol

	a[6].Key = "time_in_force"
	a[6].Value = timeInForce

	res, err := n.doPrivateQuery("/order/history", http.MethodGet, a, "")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request failed :: %d", res.StatusCode))
	}

	var r OrdersHistoryResp
	err = json.Unmarshal(body, &r.OrdersHistory)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (n *Newton) OpenOrders(limit int, offset int, symbol string, timeInForce string) (*OpenOrdersResp, error) {

	a := make([]Args, 10)

	a[1].Key = "limit"
	a[1].Value = strconv.Itoa(limit)

	a[2].Key = "offset"
	a[2].Value = strconv.Itoa(offset)

	a[3].Key = "symbol"
	a[3].Value = symbol

	a[4].Key = "time_in_force"
	a[4].Value = timeInForce

	res, err := n.doPrivateQuery("/order/history", http.MethodGet, a, "")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request failed :: %d", res.StatusCode))
	}

	body, _ := ioutil.ReadAll(res.Body)

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

	res, err := n.doPrivateQuery("/order/new", http.MethodPost, nil, string(b))
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error:%s", err.Error())
		}
	}()
	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, errors.New(fmt.Sprintf("request failed :: %d %s", res.StatusCode, body))
	}

	body, _ := ioutil.ReadAll(res.Body)

	var r OpenOrdersResp
	err = json.Unmarshal(body, &r.OpenOrders)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
