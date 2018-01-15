package util

import "testing"

func TestGetUUID(t *testing.T) {
	str, err := GetUUID()
	if err != nil {
		t.Error("错误:", err)
	} else {
		t.Error(str)
	}
}
