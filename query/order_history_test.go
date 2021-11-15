package query_test

import (
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

var (
	startDate = time.Now().Unix()
	endDate = startDate + 1
	validOrderHistory = query.OrderHistory{
		limit,
		offset,
		startDate,
		endDate,
		symbol,
		timeInForce,
	}
)

func TestOrderHistoryGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validOrderHistory

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())
	g.Expect(actualBody).Should(gomega.BeEquivalentTo(query.EMPTY_BODY))
}

func TestOrderHistoryGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validOrderHistory

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodGet))
}

func TestOrderHistoryGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validOrderHistory

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.OrderHistoryPath))
}

func TestOrdersHistoryGetParametersNoFilter(t *testing.T) {
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

func TestOrdersHistoryGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validOrderHistory

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

func TestOrderHistoryGetResponse(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.OrderHistory{}

	response := sut.GetResponse()

	g.Expect(reflect.TypeOf(response)).Should(gomega.Equal(reflect.TypeOf(&query.OrderHistoryResponse{})))
}	

func TestOrderHistoryIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validOrderHistory

	g.Expect(sut.IsPublic()).Should(gomega.BeFalse())
}
