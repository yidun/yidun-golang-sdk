package types

type BizPostJsonRequest struct {
	*PostJsonRequest
	BusinessId string
}

func NewBizPostJsonRequest(businessId string) *BizPostJsonRequest {
	bizPostJsonRequest := BizPostJsonRequest{
		PostJsonRequest: NewPostJsonRequest(),
		BusinessId:      businessId,
	}
	return &bizPostJsonRequest
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *BizPostJsonRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.PostJsonRequest.GetBusinessCustomSignParams()
	parentParams["businessId"] = r.BusinessId
	return parentParams
}
