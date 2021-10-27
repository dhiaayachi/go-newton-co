package query_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestMinimumTradeAmountsGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MinimumTradeAmounts{}

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())
	g.Expect(actualBody).Should(gomega.BeEquivalentTo(query.EMPTY_BODY))
}

func TestMinimumTradeAmountsGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MinimumTradeAmounts{}

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodGet))
}

func TestMinimumTradeAmountsGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MinimumTradeAmounts{}

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.MinimumTradeAmountsPath))
}

func TestMinimumTradeAmountGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MinimumTradeAmounts{}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestMinimumTradeAmountsGetResponse(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MinimumTradeAmounts{}

	response := sut.GetResponse()

	g.Expect(reflect.TypeOf(response)).Should(gomega.Equal(reflect.TypeOf(&query.MinimumTradeAmountsResponse{})))
}

func TestMinimumTradeAmountIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MinimumTradeAmounts{}

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
