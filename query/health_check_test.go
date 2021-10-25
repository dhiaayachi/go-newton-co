package query_test

import (
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestHealthCheckGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.HealthCheck{}
	
	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())
	g.Expect(actualBody).Should(gomega.BeEquivalentTo(query.EMPTY_BODY))
}

func TestHealthCheckGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.HealthCheck{}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestHealthCheckIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.HealthCheck{}

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
