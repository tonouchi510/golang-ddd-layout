package circles

type CircleService struct {
	circleRepository ICircleRepository
}

func (s CircleService) New(circleRepository ICircleRepository) {
	s.circleRepository = circleRepository
}

func (s CircleService) Exists(circle Circle) (bool, error) {
	duplicated, err := s.circleRepository.FindByName(circle.name)
	if err != nil {
		return false, err
	}
	return duplicated != nil, nil
}
