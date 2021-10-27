package query_test

import (
	"net/http"
	"reflect"
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

func TestHealthCheckGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.HealthCheck{}

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodGet))
}

func TestHealthCheckGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.HealthCheck{}

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.HealthCheckPath))
}

func TestHealthCheckGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.HealthCheck{}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestHealthCheckGetResponse(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.HealthCheck{}

	response := sut.GetResponse()

	g.Expect(reflect.TypeOf(response)).Should(gomega.BeNil())
}	

func TestHealthCheckIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.HealthCheck{}

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
