package types

type BizPostFormRequest struct {
	*PostFormRequest
	BusinessId string
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *BizPostFormRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.PostFormRequest.GetBusinessCustomSignParams()
	parentParams["businessId"] = r.BusinessId
	return parentParams
}

func NewBizPostFormRequest(businessId string) *BizPostFormRequest {
	return &BizPostFormRequest{
		PostFormRequest: NewPostFormRequest(),
		BusinessId:      businessId,
	}
}

// 解决方案没有businessId
func NewBizPostFormRequestWithoutBizId() *BizPostFormRequest {
	return &BizPostFormRequest{
		PostFormRequest: NewPostFormRequest(),
	}
}
