package error

type Service struct {
	message string
}
func (err Service) Error() string {
	return err.message
}

func NewServiceError(message string) error {
	return Service{message: message}
}
