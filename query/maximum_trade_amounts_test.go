package query_test

import (
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestMaximumTradeAmountsGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.MaximumTradeAmounts{}
	
	g.Expect(sut.GetBody()).Should(gomega.BeNil())
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
