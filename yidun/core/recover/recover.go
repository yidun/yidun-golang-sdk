package recover

type RequestRecover interface {
	// 当前request类是否支持恢复
	IsSupport(request interface{}) bool

	DoRecover(message RecoverMessage, responseClass interface{}) bool

	HandleFallbackResponse(data interface{})
}

type RecoverMessage struct {
	Clazz    string
	SecretId string
	Message  string
}
