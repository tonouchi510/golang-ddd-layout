package circles

import "fmt"

/* 値オブジェクト */
type CircleId string

// idを作成する場合は、この関数を呼ぶ
func NewID(value string) (CircleId, error) {
	// ここにルールを指定することで、不正なデータを作成できないようにする
	if value == "" {
		return "", fmt.Errorf("ValueError: CircleIdが空です。")
	}
	id := CircleId(value)
	return id, nil
}
