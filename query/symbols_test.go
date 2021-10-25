package query_test

import (
	"net/http"
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

const (
	baseAsset  = "BTC"
	quoteAsset = "ETH"
)

var validSymbols = query.Symbols{
	baseAsset,
	quoteAsset,
}

func TestSymbolsGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validSymbols

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())
	g.Expect(actualBody).Should(gomega.BeEquivalentTo(query.EMPTY_BODY))
}

func TestSymbolsGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validSymbols

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodGet))
}

func TestSymbolsGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validSymbols

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.SymbolsPath))
}

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

	sut := &validSymbols

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(2))

	g.Expect(parameters).Should(gomega.ContainElements(
		gomega.BeEquivalentTo(query.Parameter{string(query.BaseAsset), baseAsset}),
		gomega.BeEquivalentTo(query.Parameter{string(query.QuoteAsset), quoteAsset}),
	))
}

func TestSymbolsIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validSymbols

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
