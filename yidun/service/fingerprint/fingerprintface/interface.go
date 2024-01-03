package authiface

import "github.com/yidun/yidun-golang-sdk/yidun/service/fingerprint"

type FINGERPRINTAPI interface {
	Query(request *fingerprint.DeviceQueryResult) (response fingerprint.FingerprintResponse, err error)
}

var _FINGERPRINTAPI = (*fingerprint.FingerprintClient)(nil)
