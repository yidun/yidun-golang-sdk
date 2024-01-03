package test

import (
	"testing"

	"github.com/yidun/yidun-golang-sdk/yidun/core/endpoint"
)

func TestReadEndpoint(t *testing.T) {
	resolver := endpoint.GetEndpointResolver()
	s, err := resolver.Resolve("irisk", "cn-hangzhou")
	if err != nil {
		t.Error(err)
	}
	print(s)
}
