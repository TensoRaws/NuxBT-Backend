package util

import "fmt"

func CheckUsername(username string) error {
	l := len([]rune(username))
	if l > 20 {
		return fmt.Errorf("username too long")
	}
	if l < 2 {
		return fmt.Errorf("username too short")
	}
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	adminWords := []string{"admin", "root", "administrator", "管理员", "超级管理员", "版主", "站长", "moderator"}
	for _, word := range adminWords {
		if username == word {
			return fmt.Errorf("username cannot be %s", word)
		}
	}
	// TODO: more checks，如检查敏感词

	return nil
}
