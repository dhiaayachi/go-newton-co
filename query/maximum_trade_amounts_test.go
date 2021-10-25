package query_test

import (
	"net/http"
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestMaximumTradeAmountsGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MaximumTradeAmounts{}

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())
	g.Expect(actualBody).Should(gomega.BeEquivalentTo(query.EMPTY_BODY))
}

func TestMaximumTradeAmountsGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MaximumTradeAmounts{}

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodGet))
}

func TestMaximumTradeAmountsGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MaximumTradeAmounts{}

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.MaximumTradeAmountsPath))
}

func TestMaximumTradeAmountsGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MaximumTradeAmounts{}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestMaximumTradeAmountsIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MaximumTradeAmounts{}

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
