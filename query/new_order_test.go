package query_test

import (
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestNewOrderGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	body := query.NewOrderBody{
		OrderType:   "LIMIT",
		TimeInForce: query.IOC,
		Side:        "BUY",
		Symbol:      "BTC_USDC",
		Price:       10.0,
		Quantity:    2.5,
	}

	sut := &query.NewOrder{
		Body: body,
	}

	g.Expect(sut.GetBody()).Should(gomega.BeEquivalentTo(&body))
}

func TestNewOrderGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.NewOrder{
		Body: query.NewOrderBody{
			OrderType:   "LIMIT",
			TimeInForce: query.IOC,
			Side:        "BUY",
			Symbol:      "BTC_USDC",
			Price:       10.0,
			Quantity:    2.5,
		},
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestNewOrderIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.NewOrder{}

	g.Expect(sut.IsPublic()).Should(gomega.BeFalse())
}
