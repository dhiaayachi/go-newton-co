package query_test

import (
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestTickSizesGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.TickSizes{}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestTickSizesIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.TickSizes{}

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
