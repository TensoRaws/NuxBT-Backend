// Code generated by go generate; DO NOT EDIT.
// Code generated by go generate; DO NOT EDIT.
// Code generated by go generate; DO NOT EDIT.
package code

// codeToString use a map to store the string representation of Code
var codeToString = map[Code]string{
	InternalError:                                    "Internal error",
	UnknownError:                                     "Unknown error",
	AuthErrorTokenHasBeenBlacklisted:                 "Auth error token has been blacklisted",
	AuthErrorTokenIsInvalid:                          "Auth error token is invalid",
	AuthErrorNoPermission:                            "Auth error no permission",
	RequestErrorInvalidParams:                        "Request error invalid params",
	DatabaseErrorRecordCreateFailed:                  "Database error record create failed",
	DatabaseErrorRecordNotFound:                      "Database error record not found",
	DatabaseErrorRecordPatchFailed:                   "Database error record patch failed",
	UserErrorRegisterNotAllowed:                      "User error register not allowed",
	UserErrorInvalidUsername:                         "User error invalid username",
	UserErrorInvalidPassword:                         "User error invalid password",
	UserErrorInvalidEmail:                            "User error invalid email",
	UserErrorInvitationCodeEligibilityTimeNotReached: "User error invitation code eligibility time not reached",
	UserErrorInvalidInvitationCode:                   "User error invalid invitation code",
	UserErrorInvitationCodeHasReachedLimit:           "User error invitation code has reached limit",
}
