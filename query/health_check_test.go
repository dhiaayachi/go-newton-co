package query_test

import (
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

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
