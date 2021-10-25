package newton

import (
	"os"
	"testing"
	"time"

	"github.com/dhiaayachi/go-newton-co/query"
	"github.com/onsi/gomega"
)

func getSecrets() (string, string) {
	return os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET")
}

// Public API
///////////////////////////////////////////////////////////////////////////////////////////////////
func TestGetTickSizes(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	_, err := sut.TickSizes()

	g.Expect(err).Should(gomega.BeNil())
}

func TestGetMaximumTradeAmounts(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	_, err := sut.MaximumTradeAmounts()

	g.Expect(err).Should(gomega.BeNil())
}

func TestGetApplicableFees(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	_, err := sut.ApplicableFees()

	g.Expect(err).Should(gomega.BeNil())
}

func TestSymbolsNoQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	q := &query.Symbols{
		BaseAsset: query.NO_FILTER,
		QuoteAsset: query.NO_FILTER,
	}

	_, err := sut.Symbols(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestSymbolsWithQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	q := &query.Symbols{
		BaseAsset: "BTC",
		QuoteAsset: "ETH",
	}

	_, err := sut.Symbols(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestHealthCheck(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	err := sut.HealthCheck()

	g.Expect(err).Should(gomega.BeNil())
}

func TestGetMinTradeAmounts(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	_, err := sut.MinimumTradeAmount()

	g.Expect(err).Should(gomega.BeNil())
}

// Private API
///////////////////////////////////////////////////////////////////////////////////////////////////
func TestBalances(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	q := &query.Balances{Asset: "BTC"}
	_, err := sut.Balances(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestBalancesNoQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	q := &query.Balances{Asset: query.NO_FILTER}
	_, err := sut.Balances(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestActions(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	q := &query.Actions{
		ActionType: query.DEPOSIT,
		Limit:      1,
		Offset:     0,
		StartDate:  time.Date(2020, 01, 01, 00, 00, 00, 00, time.Local).Unix(),
		EndDate:    time.Date(2020, 01, 02, 00, 00, 00, 00, time.Local).Unix(),
	}

	_, err := sut.Actions(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestActionsNoQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	req := &query.Actions{
		ActionType: query.ActionType(query.NO_FILTER),
		Limit:      int(query.ANY),
		Offset:     int(query.ANY),
		StartDate:  int64(query.ANY),
		EndDate:    int64(query.ANY),
	}

	_, err := sut.Actions(req)

	g.Expect(err).Should(gomega.BeNil())
}

func TestOrderHistory(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	q := &query.OrderHistory{
		Limit: 1, 
		Offset: 0,
		StartDate: time.Date(2020, 01, 01, 00, 00, 00, 00, time.Local).Unix(),
		EndDate: time.Date(2020, 01, 01, 01, 00, 00, 00, time.Local).Unix(),
		Symbol: "BTC_USDC",
		TimeInForce: "IOC",
	}

	_, err := sut.OrderHistory(q)

	g.Expect(err).Should(gomega.BeNil())
}

func TestOpenOrders(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ClientId, ClientSecret := getSecrets()
	sut := New(ClientId, ClientSecret)

	q := &query.OpenOrders{
		Limit:       1,
		Offset:      0,
		Symbol:      "BTC_USDC",
		TimeInForce: "IOC"}

	_, err := sut.OpenOrders(q)

	g.Expect(err).Should(gomega.BeNil())
}
