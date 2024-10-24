/**
 ******************************************************************************
 * @file           : salt.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/26
 ******************************************************************************
 */

package tools

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSalt() (string, error) {
	salt := make([]byte, 16) // 生成16字节的随机盐值
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}
