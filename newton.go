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

func (n *Newton) Do(query query.Query) (*http.Response, error) {
	body, err := query.GetBody()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(
		query.GetMethod(),
		baseUrl + query.GetPath(), 
		bytes.NewBuffer(body))
	q := req.URL.Query()
	for _, a := range query.GetParameters() {
		q.Add(a.Key, a.Value)
	}
	req.URL.RawQuery = q.Encode()
	if query.GetMethod() != http.MethodGet {
		req.Header.Add("content-type", "application/json")
	}

	if !query.IsPublic() {
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
	res, err := n.Do(query)
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
	query := &query.MaximumTradeAmounts{}
	res, err := n.Do(query)
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
	query := &query.ApplicableFees{}
	res, err := n.Do(query)
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
	res, err := n.Do(query)
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
	query := &query.HealthCheck{}
	res, err := n.Do(query)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed :: %d", res.StatusCode)
	}

	return nil
}

func (n *Newton) MinimumTradeAmount() (*GetMinTradeAmountsResp, error) {
	query := &query.MinimumTradeAmounts{}
	res, err := n.Do(query)
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
	res, err := n.Do(query)
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
	res, err := n.Do(query)
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
	res, err := n.Do(query)
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
	res, err := n.Do(query)
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

func (n *Newton) NewOrder(query query.Query) (*OpenOrdersResp, error) {
	res, err := n.Do(query)
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
