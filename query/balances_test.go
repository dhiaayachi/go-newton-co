package query_test

import (
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestBalancesQueryNoFilter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.Balances{
		query.NO_FILTER,
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestBalancesQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	asset := "BTC"

	sut := &query.Balances{
		asset,
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(1))

	g.Expect(parameters).Should(gomega.ContainElements(
		gomega.BeEquivalentTo(query.Parameter{string(query.Asset), asset}),
	))
}

func TestBalancesIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.Balances{
		query.NO_FILTER,
	}

	g.Expect(sut.IsPublic()).Should(gomega.BeFalse())
}
