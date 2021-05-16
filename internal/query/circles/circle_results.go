package circles

type CircleSummaryData struct {
	CircleId  string
	OwnerName string
}

func NewCircleSummaryData(circleId string, ownerName string) (CircleSummaryData, error) {
	d := CircleSummaryData{
		CircleId:  circleId,
		OwnerName: ownerName,
	}
	return d, nil
}

type CircleGetSummariesResult []CircleSummaryData

func NewCircleGetSummariesResult(summaries []CircleSummaryData) (CircleGetSummariesResult, error) {
	r := CircleGetSummariesResult(summaries)
	return r, nil
}
