package query_test

import (
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestSymbolsGetParametersNoFilter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.Symbols{
		query.NO_FILTER,
		query.NO_FILTER,
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestSymbolsGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	baseAsset := "BTC"
	quoteAsset := "ETH"

	sut := &query.Symbols{
		baseAsset,
		quoteAsset,
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(2))

	g.Expect(parameters).Should(gomega.ContainElements(
		gomega.BeEquivalentTo(query.Parameter{string(query.BaseAsset), baseAsset}),
		gomega.BeEquivalentTo(query.Parameter{string(query.QuoteAsset), quoteAsset}),
	))
}

func TestSymbolsIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.Symbols{
		query.NO_FILTER,
		query.NO_FILTER,
	}

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
