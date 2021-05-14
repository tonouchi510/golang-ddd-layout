package users

import "fmt"

type UserId string

func NewUserId(value string) (UserId, error) {
	// ここにルールを指定することで、不正なデータを作成できないようにする
	if value == "" {
		return "", fmt.Errorf("ValueError: UserIdが空です。")
	}
	id := UserId(value)
	return id, nil
}
