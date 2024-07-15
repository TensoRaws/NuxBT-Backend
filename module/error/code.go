package error

type Code uint32

const (
	OK Code = 10000 + iota
	Canceled
	Unknown
	InvalidArgument
	DeadlineExceeded
	AlreadyExists
	PermissionDenied
	ResourceExhausted
	FailedPrecondition
	Aborted
	OutOfRange
	Unimplemented
	Internal
	Unavailable
	DataLoss
	Unauthenticated
)
