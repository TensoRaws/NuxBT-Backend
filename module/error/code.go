package error

type Code uint32

const (
	// DO NOT EDIT
	// gen code start
	InternalError Code = 10000 + iota
	UnknownError
	InvalidParams
	InvalidToken
	InvalidUserName
	// gen code end
	// DO NOT EDIT
)

func (c Code) String() string {
	if str, ok := codeToString[c]; ok {
		return str
	}
	return "Unknown ErrorCode"
}
