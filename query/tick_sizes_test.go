package query_test

import (
	"net/http"
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestTickSizesGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.TickSizes{}

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())
	g.Expect(actualBody).Should(gomega.BeEquivalentTo(query.EMPTY_BODY))
}

func TestTickSizesGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.TickSizes{}

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodGet))
}

func TestTickSizesGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.TickSizes{}

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.TickSizesPath))
}

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
