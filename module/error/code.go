package error

type Code uint32

const (
	// gen code start
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
	// gen code end
)

func (c Code) String() string {
	if str, ok := codeToString[c]; ok {
		return str
	}
	return "Unknown ErrorCode"
}
