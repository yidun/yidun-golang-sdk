package auth

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm3"
)

type SignatureMethod string

const (
	MD5    SignatureMethod = "MD5"
	SHA1   SignatureMethod = "SHA1"
	SHA256 SignatureMethod = "SHA256"
	SM3    SignatureMethod = "SM3"
)

var methodMap = make(map[string]SignatureMethod)

func init() {
	methodMap["MD5"] = MD5
	methodMap["SHA1"] = SHA1
	methodMap["SHA256"] = SHA256
	methodMap["SM3"] = SM3
}

type SignCalculator func(str string) string

func (s SignatureMethod) calcSign(str string) string {
	signCalculator := getSignCalculator(s)
	return signCalculator(str)
}

func getSignCalculator(s SignatureMethod) SignCalculator {
	switch s {
	case MD5:
		return func(str string) string {
			hash := md5.Sum([]byte(str))
			return hex.EncodeToString(hash[:])
		}
	case SHA1:
		return func(str string) string {
			hash := sha1.Sum([]byte(str))
			return hex.EncodeToString(hash[:])
		}
	case SHA256:
		return func(str string) string {
			hash := sha256.Sum256([]byte(str))
			return hex.EncodeToString(hash[:])
		}
	case SM3:
		return func(str string) string {
			hash := sm3.Sm3Sum([]byte(str))
			return hex.EncodeToString(hash[:])
		}
	default:
		panic("Unsupported signature method: " + string(s))
	}
}
