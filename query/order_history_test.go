package query_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestOrdersHistoryQueryNoFilter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.OrderHistory{
		query.ANY,
		query.ANY,
		int64(query.ANY),
		int64(query.ANY),
		query.NO_FILTER,
		query.NO_FILTER_VALUE,
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestOrdersHistoryQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	limit := 1
	offset := 0
	startDate := time.Now().Unix()
	endDate := startDate + 1
	symbol := "BTC_USDC"
	timeInForce := query.IOC

	sut := &query.OrderHistory{
		limit,
		offset,
		startDate,
		endDate,
		symbol,
		timeInForce,
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(6))

	g.Expect(parameters).Should(gomega.ContainElements(
		gomega.BeEquivalentTo(query.Parameter{string(query.Limit), strconv.Itoa(limit)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.Offset), strconv.Itoa(offset)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.StartDate), strconv.FormatInt(startDate, 10)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.EndDate), strconv.FormatInt(endDate, 10)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.Symbol), symbol}),
		gomega.BeEquivalentTo(query.Parameter{string(query.TimeInForce), string(timeInForce)}),
	))
}
