package newton_test

import (
	"os"
	"testing"
	"time"

	"github.com/dhiaayachi/go-newton-co"
	"github.com/dhiaayachi/go-newton-co/query"
	"github.com/onsi/gomega"
)

const (
	mockClientId       = "mock_id"
	mock_client_secret = "mock_secret"
	mockServerURL      = "https://stoplight.io/mocks/newton/newton-api-docs/431375"
)

func TestNewNewton(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.New(mockClientId, mock_client_secret)

	g.Expect(sut.ClientId).Should(gomega.Equal(mockClientId))
	g.Expect(sut.ClientSecret).Should(gomega.Equal(mock_client_secret))
	g.Expect(sut.BaseUrl).Should(gomega.Equal(newton.BaseURL))
}

// Public API
///////////////////////////////////////////////////////////////////////////////////////////////////
func TestGetTickSizes(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	_, err := sut.Do(&query.TickSizes{})

	g.Expect(err).Should(gomega.BeNil())
}

func TestGetMaximumTradeAmounts(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	_, err := sut.Do(&query.MaximumTradeAmounts{})

	g.Expect(err).Should(gomega.BeNil())
}

func TestGetApplicableFees(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	_, err := sut.Do(&query.ApplicableFees{})

	g.Expect(err).Should(gomega.BeNil())
}

func TestSymbolsNoQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	q := &query.Symbols{
		BaseAsset:  query.NO_FILTER,
		QuoteAsset: query.NO_FILTER,
	}

	_, err := sut.Do(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestSymbolsWithQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	q := &query.Symbols{
		BaseAsset:  "BTC",
		QuoteAsset: "ETH",
	}

	_, err := sut.Do(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestHealthCheck(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	_, err := sut.Do(&query.HealthCheck{})

	g.Expect(err).Should(gomega.BeNil())
}

func TestGetMinTradeAmounts(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	_, err := sut.Do(&query.MinimumTradeAmounts{})

	g.Expect(err).Should(gomega.BeNil())
}

// Private API
///////////////////////////////////////////////////////////////////////////////////////////////////
func TestBalances(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	q := &query.Balances{Asset: "BTC"}
	_, err := sut.Do(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestBalancesNoQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	q := &query.Balances{Asset: query.NO_FILTER}
	_, err := sut.Do(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestActions(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	q := &query.Actions{
		ActionType: query.DEPOSIT,
		Limit:      1,
		Offset:     0,
		StartDate:  time.Date(2020, 01, 01, 00, 00, 00, 00, time.Local).Unix(),
		EndDate:    time.Date(2020, 01, 02, 00, 00, 00, 00, time.Local).Unix(),
	}

	_, err := sut.Do(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestActionsNoQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	req := &query.Actions{
		ActionType: query.ActionType(query.NO_FILTER),
		Limit:      int(query.ANY),
		Offset:     int(query.ANY),
		StartDate:  int64(query.ANY),
		EndDate:    int64(query.ANY),
	}

	_, err := sut.Do(req)

	g.Expect(err).Should(gomega.BeNil())
}

func TestOrderHistory(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	q := &query.OrderHistory{
		Limit:       1,
		Offset:      0,
		StartDate:   time.Date(2020, 01, 01, 00, 00, 00, 00, time.Local).Unix(),
		EndDate:     time.Date(2020, 01, 01, 01, 00, 00, 00, time.Local).Unix(),
		Symbol:      "BTC_USDC",
		TimeInForce: "IOC",
	}

	_, err := sut.Do(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestOpenOrders(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	q := &query.OpenOrders{
		Limit:       1,
		Offset:      0,
		Symbol:      "BTC_USDC",
		TimeInForce: "IOC"}

	_, err := sut.Do(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestNewOrder(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	q := &query.NewOrder{
		Body: query.NewOrderBody{
			OrderType:   "LIMIT",
			TimeInForce: query.IOC,
			Side:        "BUY",
			Symbol:      "BTC_USDC",
			Price:       1000,
			Quantity:    0.0001,
		},
	}

	_, err := sut.Do(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestCancelOrder(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := newton.Newton{
		mockClientId,
		mock_client_secret,
		mockServerURL,
	}

	q := &query.CancelOrder{
		Body: query.CancelOrderBody{
			OrderId: "test",
		},
	}

	_, err := sut.Do(q)

	g.Expect(err).Should(gomega.BeNil())
}

// Newton's mock server does not utilize authnetication headers.
// Therefore, to fully test the API, we must test any private production endpoint
// with real credentials.
func TestAuthentication(t *testing.T) {
	if os.Getenv("TEST_AUTH") != "true" {
		t.Skip("Skipping authentication test in production environment.")
	}
	g := gomega.NewGomegaWithT(t)

	productionClientId := os.Getenv("CLIENT_ID")
	productionClientSecret := os.Getenv("CLIENT_SECRET")

	sut := newton.Newton{
		productionClientId,
		productionClientSecret,
		newton.BaseURL, // production URL
	}

	q := &query.Balances{Asset: "BTC"}
	_, err := sut.Do(q)

	g.Expect(err).Should(gomega.BeNil())
}
