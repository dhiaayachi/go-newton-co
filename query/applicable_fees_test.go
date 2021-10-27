package query_test

import (
	"net/http"
	"reflect"
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

func TestApplicableFeesGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.ApplicableFees{}

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodGet))
}

func TestApplicableFeesGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.ApplicableFees{}

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.ApplicableFeesPath))
}

func TestApplicableFeesGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.ApplicableFees{}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestApplicableFeesGetResponse(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.ApplicableFees{}

	response := sut.GetResponse()

	g.Expect(reflect.TypeOf(response)).Should(gomega.Equal(reflect.TypeOf(&query.ApplicableFeesResponse{})))
}	

func TestApplicableFeesIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.ApplicableFees{}

	g.Expect(sut.IsPublic()).Should(gomega.BeTrue())
}
