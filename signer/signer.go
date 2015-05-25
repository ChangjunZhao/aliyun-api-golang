package signer

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

//签名接口
type signer interface {
	Sign(message string) (string, error)
	Debug(enabled bool)
}

func NewSigner(accessKeySecret string) *SHA1Signer {
	return &SHA1Signer{accessKeySecret: accessKeySecret}
}

//SHA1签名
type SHA1Signer struct {
	accessKeySecret string
	debug           bool
}

func (s *SHA1Signer) Debug(enabled bool) {
	s.debug = enabled
}

func (s *SHA1Signer) Sign(message string) (string, error) {
	hashfun := hmac.New(sha1.New, []byte(s.accessKeySecret))
	hashfun.Write([]byte(message))
	rawSignature := hashfun.Sum(nil)
	base64signature := base64.StdEncoding.EncodeToString(rawSignature)
	if s.debug {
		fmt.Println("Base64 signature:", base64signature)
	}
	return base64signature, nil
}
