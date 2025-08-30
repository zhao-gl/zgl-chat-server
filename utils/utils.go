package utils

import (
	"crypto/rand"
	"fmt"
)

func UUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		panic(err)
	}
	// 设置版本号 (RFC 4122 版本 4)
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	// 设置变体 (RFC 4122)
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	uuidStr := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	return uuidStr
}
