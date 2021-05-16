package circles

type CircleQueryService struct {
	repo ICircleQueryRepository
}

func NewCircleQueryService(
	repo ICircleQueryRepository,
) (CircleQueryService, error) {
	s := CircleQueryService{
		repo: repo,
	}
	return s, nil
}

func (s CircleQueryService) GetSummaries(command CircleGetSummariesCommand) (CircleGetSummariesResult, error) {
	page := command.Page
	size := command.Size
	summaries, err := s.repo.ListCircleSummaries(page, size)
	if err != nil {
		return nil, err
	}
	return summaries, nil
}
