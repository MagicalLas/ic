package impression

type Service struct{}

func (s Service) createTrackingID() (int64, error) {
	return 0, nil
}

func (s Service) imp() error {
	return nil
}

func gc() error {
	return nil
}

func checkMemoryFragmentation() error {
	return nil
}
