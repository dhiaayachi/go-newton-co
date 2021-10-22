package query_test

import (
	"strconv"
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestOpenOrdersQueryNoFilter(t *testing.T) {
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

func TestOpenOrdersQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	limit := 1
	offset := 0
	symbol := "BTC_USDC"
	timeInForce := query.IOC

	sut := &query.OpenOrders{
		limit,
		offset,
		symbol,
		timeInForce,
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(4))

	g.Expect(parameters).Should(gomega.ContainElements(
		gomega.BeEquivalentTo(query.Parameter{string(query.Limit), strconv.Itoa(limit)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.Offset), strconv.Itoa(offset)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.Symbol), symbol}),
		gomega.BeEquivalentTo(query.Parameter{string(query.TimeInForce), string(timeInForce)}),
	))
}
