package helper

import (
	"crypto/md5"
	"fmt"
)

func GetMD5Hash(text string) (string, error) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(text)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func BaiShiSign(test string, key string) (string, error) {
	test = test + key
	return GetMD5Hash(test)
}
