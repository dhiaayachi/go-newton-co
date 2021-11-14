package query_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

var (
	cancelOrderBody query.CancelOrderBody = query.CancelOrderBody{
		OrderId:   "8c1be04a-813d-4abd-8b15-9367686afaed",
	}
	validCancelOrder query.CancelOrder = query.CancelOrder{
		Body: cancelOrderBody,
	}
)

func TestCancelOrderGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validCancelOrder

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())

	var actualBodyParsed query.CancelOrderBody
	err = json.Unmarshal([]byte(actualBody), &actualBodyParsed)
	g.Expect(err).Should(gomega.BeNil())

	g.Expect(actualBodyParsed).Should(gomega.BeEquivalentTo(cancelOrderBody))
}

func TestCancelOrderGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.CancelOrder{}

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodPost))
}

func TestCancelOrderGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.CancelOrder{}

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.CancelOrderPath))
}

func TestCancelOrderGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validCancelOrder

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestCancelOrderGetResponse(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.CancelOrder{}

	response := sut.GetResponse()

	g.Expect(reflect.TypeOf(response)).Should(gomega.Equal(reflect.TypeOf(&query.CancelOrderResponse{})))
}	

func TestCancelOrderIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validCancelOrder

	g.Expect(sut.IsPublic()).Should(gomega.BeFalse())
}
