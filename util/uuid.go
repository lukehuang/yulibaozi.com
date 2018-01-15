package util

import (
	uuid "github.com/satori/go.uuid"
)

// GetUUID 生成uuid
func GetUUID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
