package tools

import (
	"crypto/md5"
	"github.com/anaskhan96/go-password-encoder"
)

func GenerateMd5(psd string) (encodedPsd, salt string) {
	options := &password.Options{SaltLen: 10, Iterations: 10, KeyLen: 30, HashFunction: md5.New}
	salt, encodedPsd = password.Encode(psd, options)
	return
}

func VerifyMd5(encodePsd, psd, salt string) bool {
	options := &password.Options{SaltLen: 10, Iterations: 10, KeyLen: 30, HashFunction: md5.New}
	check := password.Verify(psd, salt, encodePsd, options)
	return check
}
