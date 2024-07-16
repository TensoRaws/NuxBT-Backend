package code

type Code uint32

// 请勿修改，常量之间不要有空行！
// DO NOT EDIT, no empty line between constants!
const (
	// DO NOT EDIT
	// gen code start
	InternalError Code = 10000 + iota
	UnknownError
	// AuthError 鉴权错误
	AuthErrorTokenHasBeenBlacklisted
	AuthErrorTokenIsInvalid
	// RequestError 请求错误
	RequestErrorInvalidParams
	// DatabaseError 数据库错误
	DatabaseErrorRecordCreateFailed
	DatabaseErrorRecordNotFound
	DatabaseErrorRecordPatchFailed
	// UserError 用户侧错误
	UserErrorRegisterNotAllowed
	UserErrorInvalidUsername
	UserErrorInvalidPassword
	UserErrorInvalidEmail
	UserErrorInvalidInvitationCode
	UserErrorInvitationCodeHasReachedLimit
	// gen code end
	// DO NOT EDIT
)

func (c Code) String() string {
	if str, ok := codeToString[c]; ok {
		return str
	}
	return "Unknown ErrorCode"
}
