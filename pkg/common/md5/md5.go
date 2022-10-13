package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(value string) string {
	result := md5.Sum([]byte(value))
	return hex.EncodeToString(result[:])
}
