package users

import "fmt"

type UserName string

func NewUserName(value string) (UserName, error) {
	if value == "" {
		return "", fmt.Errorf("ValueError: UserNameが空です。")
	}
	if len(value) < 3 {
		return "", fmt.Errorf("ValueError: ユーザ名は3文字以上です。")
	}
	name := UserName(value)
	return name, nil
}
