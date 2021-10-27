package query_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

const asset = "BTC"
var validBalances query.Balances = query.Balances{
	asset,
}

func TestBalancesGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validBalances

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())
	g.Expect(actualBody).Should(gomega.BeEquivalentTo(query.EMPTY_BODY))
}

func TestBalancesGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validBalances

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodGet))
}

func TestBalancesGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.ApplicableFees{}

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.ApplicableFeesPath))
}

func TestBalancesGetParametersNoFilter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.Balances{
		query.NO_FILTER,
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestBalancesGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validBalances

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(1))

	g.Expect(parameters).Should(gomega.ContainElements(
		gomega.BeEquivalentTo(query.Parameter{string(query.Asset), asset}),
	))
}

func TestBalancesGetResponse(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.Balances{}

	response := sut.GetResponse()

	g.Expect(reflect.TypeOf(response)).Should(gomega.Equal(reflect.TypeOf(&query.BalancesResponse{})))
}	

func TestBalancesIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validBalances

	g.Expect(sut.IsPublic()).Should(gomega.BeFalse())
}
