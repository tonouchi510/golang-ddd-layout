package circles

type ICircleQueryRepository interface {
	ListCircleSummaries(page int, size int) (CircleGetSummariesResult, error)
}
