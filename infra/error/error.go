package apperror

// ErrorManager error manager interface
type ErrorManager interface {
	Wrap(err error, code int) error
	LogMessage(err error) string
	Code(err error) int
}
