package http

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func CreateByConfig(config *HttpClientConfig) *CustomClient {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(config.ConnectionTimeoutMs) * time.Millisecond,
			KeepAlive: time.Duration(config.ConnectionKeepAliveMillis) * time.Millisecond,
		}).DialContext,
		TLSHandshakeTimeout: time.Duration(config.ConnectionTimeoutMs) * time.Millisecond,
		MaxIdleConns:        config.MaxConnectionCount,
		MaxIdleConnsPerHost: config.MaxConnectionCountPerRoute,
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(config.ResponseTimeoutMs) * time.Millisecond,
	}
	userAgent := fmt.Sprintf("%s/%s;%s", SDKVersion, DefaultSDKVersion, HttpCLientVersion)
	return &CustomClient{
		client:    httpClient,
		userAgent: userAgent,
	}
}
