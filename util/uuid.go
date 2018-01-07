package util

import (
	uuid "github.com/satori/go.uuid"
)

// GetUUID 生成uuid
func GetUUID() string {
	return uuid.NewV4().String()
}
