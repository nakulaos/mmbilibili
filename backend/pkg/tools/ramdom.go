package tools

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomToken() string {
	bytes := make([]byte, 36) // 生成16字节的随机数
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
