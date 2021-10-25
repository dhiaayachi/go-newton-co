package query_test

import (
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

func TestMinimumTradeAmountGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MinimumTradeAmounts{}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestMinimumTradeAmountIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MinimumTradeAmounts{}

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
