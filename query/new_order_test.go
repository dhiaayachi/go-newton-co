package query_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

var (
	body query.NewOrderBody = query.NewOrderBody{
	OrderType:   "LIMIT",
	TimeInForce: query.IOC,
	Side:        "BUY",
	Symbol:      "BTC_USDC",
	Price:       10.0,
	Quantity:    2.5,
	}
	validNewOrder query.NewOrder = query.NewOrder{
		Body: body,
	}
)

func TestNewOrderGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validNewOrder

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())

	var actualBodyParsed query.NewOrderBody
	err = json.Unmarshal([]byte(actualBody), &actualBodyParsed)
	g.Expect(err).Should(gomega.BeNil())

	g.Expect(actualBodyParsed).Should(gomega.BeEquivalentTo(body))
}

func TestNewOrderGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.NewOrder{}

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodPost))
}

func TestNewOrderGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.NewOrder{}

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.NewOrderPath))
}

func TestNewOrderGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validNewOrder

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestNewOrderIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validNewOrder

	g.Expect(sut.IsPublic()).Should(gomega.BeFalse())
}
