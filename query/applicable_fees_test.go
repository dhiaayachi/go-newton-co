package query_test

import (
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestApplicableFeesGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.ApplicableFees{}

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())
	g.Expect(actualBody).Should(gomega.BeEquivalentTo(query.EMPTY_BODY))
}

func TestApplicableFeesGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.ApplicableFees{}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestApplicableFeesIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.ApplicableFees{}

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
