package query_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/dhiaayachi/go-newton-co/query"

	"github.com/onsi/gomega"
)

func TestActionsQueryNoFilter(t *testing.T) {
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

func TestActionsQuery(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	actionType := query.DEPOSIT
	limit := 1
	offset := 2
	startTime := time.Now().Unix()
	endTime := startTime + 1

	sut := &query.Actions{
		query.DEPOSIT,
		limit,
		offset,
		startTime,
		endTime,
	}

	parameters := sut.GetParameters()

	g.Expect(len(parameters)).Should(gomega.Equal(5))

	g.Expect(parameters).Should(gomega.ContainElements(
		gomega.BeEquivalentTo(query.Parameter{string(query.ActionTypeKey), string(actionType)}),
		gomega.BeEquivalentTo(query.Parameter{string(query.Limit), strconv.Itoa(limit)}) ,
		gomega.BeEquivalentTo(query.Parameter{string(query.Offset), strconv.Itoa(offset)}) ,
		gomega.BeEquivalentTo(query.Parameter{string(query.StartDate), strconv.FormatInt(startTime, 10)}) ,
		gomega.BeEquivalentTo(query.Parameter{string(query.EndDate), strconv.FormatInt(endTime, 10)}) ,
	))
}
