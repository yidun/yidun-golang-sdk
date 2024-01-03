package submit

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// 文档检测提交请求
type FileSubmitV2Request struct {
	*types.BizPostFormRequest
	//数据ID
	DataId *string `json:"dataId,omitempty" validate:"max=128"`
	// 子数据类型，与易盾内容安全服务约定即可
	DataType *int `json:"dataType,omitempty"`
	// 文件名称/标题
	FileName *string `json:"fileName,omitempty"`
	// 文件类型
	FileType *int `json:"fileType,omitempty"`
	// 文档URL
	URL *string `json:"url,omitempty" validate:"max=512"`
	// 文档内容, 不能和url同时为空，也不能和url同时有值
	Content *string `json:"content,omitempty" validate:"max=100000"`
	// 检测标记 1: 仅检测文本， 2: 仅检测图片， 3：检测文本+图片， 7:检测文本+图片+img标签中的图片url，8：仅检测音视频，15：检测文本+图片+音视频，默认值为7
	CheckFlag *int `json:"checkFlag,omitempty"`
	// 用户唯一标识，如果无需登录则为空
	Account *string `json:"account,omitempty" validate:"max=128"`
	// 用户IP地址
	IP *string `json:"ip,omitempty" validate:"max=128"`
	// 数据回调参数，调用方根据业务情况自行设计
	Callback *string `json:"callback,omitempty" validate:"max=512"`
	// 异步检测结果回调通知的URL
	CallbackURL *string `json:"callbackUrl,omitempty" validate:"max=1024"`
}

// NewFileSubmitV2Request 初始化FileSubmitV2Request对象
func NewFileSubmitV2Request() *FileSubmitV2Request {
	defaultCheckFlag := 7
	var request = &FileSubmitV2Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
		CheckFlag:          &defaultCheckFlag,
	}
	request.SetProductCode("file")
	request.SetUriPattern("/v2/file/submit")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v2.0")
	return request
}

func (f *FileSubmitV2Request) SetDataId(dataId string) {
	f.DataId = &dataId
}

func (f *FileSubmitV2Request) SetDataType(dataType int) {
	f.DataType = &dataType
}

func (f *FileSubmitV2Request) SetFileName(fileName string) {
	f.FileName = &fileName
}

func (f *FileSubmitV2Request) SetFileType(fileType int) {
	f.FileType = &fileType
}

func (f *FileSubmitV2Request) SetURL(url string) {
	f.URL = &url
}

func (f *FileSubmitV2Request) SetContent(content string) {
	f.Content = &content
}

func (f *FileSubmitV2Request) SetCheckFlag(checkFlag int) {
	f.CheckFlag = &checkFlag
}

func (f *FileSubmitV2Request) SetAccount(account string) {
	f.Account = &account
}

func (f *FileSubmitV2Request) SetIP(IP string) {
	f.IP = &IP
}

func (f *FileSubmitV2Request) SetCallback(callback string) {
	f.Callback = &callback
}

func (f *FileSubmitV2Request) SetCallbackURL(callbackURL string) {
	f.CallbackURL = &callbackURL
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (f *FileSubmitV2Request) GetBusinessCustomSignParams() map[string]string {
	result := f.PostFormRequest.GetBusinessCustomSignParams()

	if f.DataId != nil {
		result["dataId"] = *f.DataId
	}
	if f.DataType != nil {
		result["dataType"] = strconv.Itoa(*f.DataType)
	}
	if f.FileName != nil {
		result["fileName"] = *f.FileName
	}
	if f.FileType != nil {
		result["fileType"] = strconv.Itoa(*f.FileType)
	}
	if f.URL != nil {
		result["url"] = *f.URL
	}
	if f.Content != nil {
		result["content"] = *f.Content
	}
	if f.CheckFlag != nil {
		result["checkFlag"] = strconv.Itoa(*f.CheckFlag)
	}
	if f.Account != nil {
		result["account"] = *f.Account
	}
	if f.IP != nil {
		result["ip"] = *f.IP
	}
	if f.Callback != nil {
		result["callback"] = *f.Callback
	}
	if f.CallbackURL != nil {
		result["callbackUrl"] = *f.CallbackURL
	}
	return result
}
