package apperror

// ErrorManager
type ErrorManager interface {
	Wrap(err error, code int) error
	LogMessage(err error) string
	Code(err error) int
}
