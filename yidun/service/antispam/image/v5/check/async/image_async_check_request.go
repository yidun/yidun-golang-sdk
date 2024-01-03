package async

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check"
)

type ImageV5AsyncCheckRequest struct {
	*check.ImageV5CheckRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"`
}

// For ImageV5AsyncCheckRequest
func NewImageV5AsyncCheckRequest(businessId string) *ImageV5AsyncCheckRequest {
	var checkSceneRequest = check.NewImageCheckSceneRequest(businessId)
	var request = &check.ImageV5CheckRequest{
		ImageCheckSceneRequest: checkSceneRequest,
	}
	checkSceneRequest.SetProductCode("imageCheck")
	checkSceneRequest.SetUriPattern("/v5/image/asyncCheck")
	checkSceneRequest.SetVersion("v5.2")
	var asyncCheckRequest = &ImageV5AsyncCheckRequest{
		ImageV5CheckRequest: request,
	}
	return asyncCheckRequest
}

func (r *ImageV5AsyncCheckRequest) SetImageV5CheckRequest(imageV5CheckRequest *check.ImageV5CheckRequest) {
	r.ImageV5CheckRequest = imageV5CheckRequest
}
