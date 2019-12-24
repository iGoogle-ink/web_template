package pay

func (s *Service) SendMessage(msg string) error {
	err := s.producer.Publish([]byte(msg))
	return err
}
