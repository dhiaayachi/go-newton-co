package query_test

import (
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestMinimumTradeAmountGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MinimumTradeAmount{}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestMinimumTradeAmountIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MinimumTradeAmount{}

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
