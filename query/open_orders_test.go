package query_test

import (
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)
const(
	symbol = "BTC_USDC"
	timeInForce = query.IOC
)

var validOpenOrders query.OpenOrders = query.OpenOrders{
	limit,
	offset,
	symbol,
	timeInForce,
}

func TestOpenOrdersGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validOpenOrders

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())
	g.Expect(actualBody).Should(gomega.BeEquivalentTo(query.EMPTY_BODY))
}

func TestOpenOrdersGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validOpenOrders

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodGet))
}

func TestOpenOrdersGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validOpenOrders

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.OpenOrdersPath))
}

func TestOpenOrdersGetParametersNoFilter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.OpenOrders{
		query.ANY,
		query.ANY,
		query.NO_FILTER,
		query.NO_FILTER_VALUE,
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestOpenOrdersGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validOpenOrders

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(4))

	g.Expect(parameters).Should(gomega.ContainElements(
		gomega.BeEquivalentTo(query.Parameter{string(query.Limit), strconv.Itoa(limit)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.Offset), strconv.Itoa(offset)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.Symbol), symbol}),
		gomega.BeEquivalentTo(query.Parameter{string(query.TimeInForce), string(timeInForce)}),
	))
}

func TestOpenOrdersGetResponse(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.OpenOrders{}

	response := sut.GetResponse()

	g.Expect(reflect.TypeOf(response)).Should(gomega.Equal(reflect.TypeOf(&query.OpenOrdersResponse{})))
}	

func TestOpenOrdersIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validOpenOrders

	g.Expect(sut.IsPublic()).Should(gomega.BeFalse())
}
