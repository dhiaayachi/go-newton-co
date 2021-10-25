package query_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

const (
	limit                       = 1
	offset                      = 2
	actionType query.ActionType = query.DEPOSIT
)

var (
	startTime    = time.Now().Unix()
	endTime      = startTime + 1
	validActions = query.Actions{
		actionType,
		limit,
		offset,
		startTime,
		endTime,
	}
)

func TestActionsGetBody(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validActions

	actualBody, err := sut.GetBody()
	g.Expect(err).Should(gomega.BeNil())
	g.Expect(actualBody).Should(gomega.BeEquivalentTo(query.EMPTY_BODY))
}

func TestActionsGetMethod(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validActions

	g.Expect(sut.GetMethod()).Should(gomega.Equal(http.MethodGet))
}

func TestActionsGetPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validActions

	g.Expect(sut.GetPath()).Should(gomega.Equal(query.ActionsPath))
}

func TestActionsGetParametersNoFilter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &query.Actions{
		query.ActionType(query.NO_FILTER),
		int(query.ANY),
		int(query.ANY),
		int64(query.ANY),
		int64(query.ANY),
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(0))
}

func TestActionsGetParameters(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validActions

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(5))

	g.Expect(parameters).Should(gomega.ContainElements(
		gomega.BeEquivalentTo(query.Parameter{string(query.ActionTypeKey), string(actionType)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.Limit), strconv.Itoa(limit)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.Offset), strconv.Itoa(offset)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.StartDate), strconv.FormatInt(startTime, 10)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.EndDate), strconv.FormatInt(endTime, 10)}),
	))
}

func TestActionsIsPublic(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	sut := &validActions

	g.Expect(sut.IsPublic()).Should(gomega.BeFalse())
}
