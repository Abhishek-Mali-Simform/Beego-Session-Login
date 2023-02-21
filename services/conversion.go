package services

import (
	"crypto/md5"
	"encoding/hex"
)

func ConvertToMD5(values string) (hash string) {
	hasher := md5.New()
	hasher.Write([]byte(values))
	hash = hex.EncodeToString(hasher.Sum(nil))
	return
}
