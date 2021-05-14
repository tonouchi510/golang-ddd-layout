package circles

import (
	"fmt"
)

type CircleName string

func NewName(value string) (CircleName, error) {
	if value == "" {
		return "", fmt.Errorf("ValueError: CircleNameが空です。")
	}
	if len(value) < 3 {
		return "", fmt.Errorf("ValueError: サークル名は3文字以上です。")
	}
	if len(value) > 20 {
		return "", fmt.Errorf("ValueError: サークル名は20文字以下です。")
	}
	name := CircleName(value)
	return name, nil
}

func (n CircleName) Equals(other CircleName) (bool, error) {
	return n == other, nil
}
