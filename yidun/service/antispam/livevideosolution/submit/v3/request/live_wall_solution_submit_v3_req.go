package request

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	videoCheck "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/check/v4/request"
)

// LiveWallSolutionSubmitV3Req represents live wall solution submit version 3 request.
type LiveWallSolutionSubmitV3Req struct {
	*types.PostFormRequest
	Url                *string                              `json:"url,omitempty"`
	DataId             *string                              `json:"dataId,omitempty"`
	Ip                 *string                              `json:"ip,omitempty"`
	Account            *string                              `json:"account,omitempty"`
	DeviceId           *string                              `json:"deviceId,omitempty"`
	DeviceType         *int                                 `json:"deviceType,omitempty"`
	Age                *string                              `json:"age,omitempty"`
	LabourUnion        *string                              `json:"labourUnion,omitempty"`
	OperationManager   *string                              `json:"operationManager,omitempty"`
	Callback           *string                              `json:"callback,omitempty"`
	CallbackUrl        *string                              `json:"callbackUrl,omitempty"`
	AccountLevel       *string                              `json:"accountLevel,omitempty"`
	AccountName        *string                              `json:"accountName,omitempty"`
	RoomId             *string                              `json:"roomId,omitempty"`
	Title              *string                              `json:"title,omitempty"`
	ScFrequency        *string                              `json:"scFrequency,omitempty"`
	DetectType         *int                                 `json:"detectType,omitempty"`
	UniqueKey          *string                              `json:"uniqueKey,omitempty"`
	BackgroundImage    *string                              `json:"backgroundImage,omitempty"`
	Cover              *string                              `json:"cover,omitempty"`
	Photo              *string                              `json:"photo,omitempty"`
	PublishTime        *int64                               `json:"publishTime,omitempty"`
	LiveLink           *string                              `json:"liveLink,omitempty"`
	AdvancedFrequency  *videoCheck.AdvancedFrequencyRequest `json:"advancedFrequency,omitempty"`
	ScreenMode         *int                                 `json:"screenMode,omitempty"`
	CheckLanguageCode  *string                              `json:"checkLanguageCode,omitempty"`
	Gender             *int                                 `json:"gender,omitempty"`
	CheckSpeakerIds    *[]string                            `json:"checkSpeakerIds,omitempty"`
	NoCheckSpeakerIds  *[]string                            `json:"noCheckSpeakerIds,omitempty"`
	ExtLon1            *int64                               `json:"extLon1,omitempty"`
	ExtLon2            *int64                               `json:"extLon2,omitempty"`
	ExtStr1            *string                              `json:"extStr1,omitempty"`
	ExtStr2            *string                              `json:"extStr2,omitempty"`
	Extension          *string                              `json:"extension,omitempty"`
	ScreenShotStrategy *int                                 `json:"screenShotStrategy,omitempty"`
	SubProduct         *string                              `json:"subProduct,omitempty"`
	AccountInfo        *string                              `json:"accountInfo,omitempty"`
}

// LiveWallSolutionSubmitV3Req 创建一个 LiveAudioCallbackV4Req 实例
func NewLiveWallSolutionSubmitV3Req() *LiveWallSolutionSubmitV3Req {
	var request = &LiveWallSolutionSubmitV3Req{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("liveVideoSolutionCheck")
	request.SetUriPattern("/v3/livewallsolution/check")
	request.SetVersion("v3")
	return request
}

// GetCustomSignParams 获取自定义签名参数
func (l *LiveWallSolutionSubmitV3Req) GetBusinessCustomSignParams() map[string]string {
	m := l.PostFormRequest.GetBusinessCustomSignParams()
	if l.Account != nil {
		m["account"] = *l.Account
	}
	if l.AccountLevel != nil {
		m["accountLevel"] = *l.AccountLevel
	}
	if l.AccountName != nil {
		m["accountName"] = *l.AccountName
	}
	if l.Age != nil {
		m["age"] = *l.Age
	}
	if l.BackgroundImage != nil {
		m["backgroundImage"] = *l.BackgroundImage
	}
	if l.Callback != nil {
		m["callback"] = *l.Callback
	}
	if l.CallbackUrl != nil {
		m["callbackUrl"] = *l.CallbackUrl
	}
	if l.Cover != nil {
		m["cover"] = *l.Cover
	}
	if l.DataId != nil {
		m["dataId"] = *l.DataId
	}
	if l.DetectType != nil {
		m["detectType"] = strconv.Itoa(*l.DetectType)
	}
	if l.DeviceId != nil {
		m["deviceId"] = *l.DeviceId
	}
	if l.DeviceType != nil {
		m["deviceType"] = strconv.Itoa(*l.DeviceType)
	}
	if l.Ip != nil {
		m["ip"] = *l.Ip
	}
	if l.LabourUnion != nil {
		m["labourUnion"] = *l.LabourUnion
	}
	if l.OperationManager != nil {
		m["operationManager"] = *l.OperationManager
	}
	if l.Photo != nil {
		m["photo"] = *l.Photo
	}
	if l.PublishTime != nil {
		m["publishTime"] = strconv.FormatInt(*l.PublishTime, 10)
	}
	if l.RoomId != nil {
		m["roomId"] = *l.RoomId
	}
	if l.ScFrequency != nil {
		m["scFrequency"] = *l.ScFrequency
	}
	if l.Title != nil {
		m["title"] = *l.Title
	}
	if l.UniqueKey != nil {
		m["uniqueKey"] = *l.UniqueKey
	}
	if l.Url != nil {
		m["url"] = *l.Url
	}
	if l.AdvancedFrequency != nil {
		// Assuming you have json package imported as: "encoding/json"
		advancedFrequencyJson, _ := json.Marshal(l.AdvancedFrequency)
		m["advancedFrequency"] = string(advancedFrequencyJson)
	}
	if l.LiveLink != nil {
		m["liveLink"] = *l.LiveLink
	}
	if l.ScreenMode != nil {
		m["screenMode"] = strconv.Itoa(*l.ScreenMode)
	}
	if l.CheckLanguageCode != nil {
		m["checkLanguageCode"] = *l.CheckLanguageCode
	}
	if l.Gender != nil {
		m["gender"] = strconv.Itoa(*l.Gender)
	}
	checkSpeakerIds := l.CheckSpeakerIds
	if checkSpeakerIds != nil && len(*checkSpeakerIds) > 0 {
		m["checkSpeakerIds"] = strings.Join(*l.CheckSpeakerIds, ",")
	}
	// Convert NoCheckSpeakerIds to comma-separated string
	noCheckSpeakerIds := l.NoCheckSpeakerIds
	if noCheckSpeakerIds != nil && len(*noCheckSpeakerIds) > 0 {
		m["noCheckSpeakerIds"] = strings.Join(*l.NoCheckSpeakerIds, ",")
	}
	if l.ExtLon1 != nil {
		m["extLon1"] = strconv.FormatInt(*l.ExtLon1, 10)
	}
	if l.ExtLon2 != nil {
		m["extLon2"] = strconv.FormatInt(*l.ExtLon2, 10)
	}
	if l.ExtStr1 != nil {
		m["extStr1"] = *l.ExtStr1
	}
	if l.ExtStr2 != nil {
		m["extStr2"] = *l.ExtStr2
	}
	if l.ScreenShotStrategy != nil {
		m["screenShotStrategy"] = strconv.Itoa(*l.ScreenShotStrategy)
	}
	if l.SubProduct != nil {
		m["subProduct"] = *l.SubProduct
	}
	if l.AccountInfo != nil {
		m["accountInfo"] = *l.AccountInfo
	}
	return m
}

// SetUrl sets the Url.
func (l *LiveWallSolutionSubmitV3Req) SetUrl(url string) {
	l.Url = &url
}

// SetDataId sets the DataId.
func (l *LiveWallSolutionSubmitV3Req) SetDataId(dataId string) {
	l.DataId = &dataId
}

// SetIp sets the Ip.
func (l *LiveWallSolutionSubmitV3Req) SetIp(ip string) {
	l.Ip = &ip
}

// SetAccount sets the Account.
func (l *LiveWallSolutionSubmitV3Req) SetAccount(account string) {
	l.Account = &account
}

// SetDeviceId sets the DeviceId.
func (l *LiveWallSolutionSubmitV3Req) SetDeviceId(deviceId string) {
	l.DeviceId = &deviceId
}

// SetDeviceType sets the DeviceType.
func (l *LiveWallSolutionSubmitV3Req) SetDeviceType(deviceType int) {
	l.DeviceType = &deviceType
}

// SetAge sets the Age.
func (l *LiveWallSolutionSubmitV3Req) SetAge(age string) {
	l.Age = &age
}

// SetLabourUnion sets the LabourUnion.
func (l *LiveWallSolutionSubmitV3Req) SetLabourUnion(labourUnion string) {
	l.LabourUnion = &labourUnion
}

// SetOperationManager sets the OperationManager.
func (l *LiveWallSolutionSubmitV3Req) SetOperationManager(operationManager string) {
	l.OperationManager = &operationManager
}

// SetCallback sets the Callback.
func (l *LiveWallSolutionSubmitV3Req) SetCallback(callback string) {
	l.Callback = &callback
}

// SetCallbackUrl sets the CallbackUrl.
func (l *LiveWallSolutionSubmitV3Req) SetCallbackUrl(callbackUrl string) {
	l.CallbackUrl = &callbackUrl
}

// SetAccountLevel sets the AccountLevel.
func (l *LiveWallSolutionSubmitV3Req) SetAccountLevel(accountLevel string) {
	l.AccountLevel = &accountLevel
}

// SetAccountName sets the AccountName.
func (l *LiveWallSolutionSubmitV3Req) SetAccountName(accountName string) {
	l.AccountName = &accountName
}

// SetRoomId sets the RoomId.
func (l *LiveWallSolutionSubmitV3Req) SetRoomId(roomId string) {
	l.RoomId = &roomId
}

// SetTitle sets the Title.
func (l *LiveWallSolutionSubmitV3Req) SetTitle(title string) {
	l.Title = &title
}

// SetScFrequency sets the ScFrequency.
func (l *LiveWallSolutionSubmitV3Req) SetScFrequency(scFrequency string) {
	l.ScFrequency = &scFrequency
}

// SetDetectType sets the DetectType.
func (l *LiveWallSolutionSubmitV3Req) SetDetectType(detectType int) {
	l.DetectType = &detectType
}

// SetUniqueKey sets the UniqueKey.
func (l *LiveWallSolutionSubmitV3Req) SetUniqueKey(uniqueKey string) {
	l.UniqueKey = &uniqueKey
}

// SetBackgroundImage sets the BackgroundImage.
func (l *LiveWallSolutionSubmitV3Req) SetBackgroundImage(backgroundImage string) {
	l.BackgroundImage = &backgroundImage
}

// SetCover sets the Cover.
func (l *LiveWallSolutionSubmitV3Req) SetCover(cover string) {
	l.Cover = &cover
}

// SetPhoto sets the Photo.
func (l *LiveWallSolutionSubmitV3Req) SetPhoto(photo string) {
	l.Photo = &photo
}

// SetPublishTime sets the PublishTime.
func (l *LiveWallSolutionSubmitV3Req) SetPublishTime(publishTime int64) {
	l.PublishTime = &publishTime
}

// SetLiveLink sets the LiveLink.
func (l *LiveWallSolutionSubmitV3Req) SetLiveLink(liveLink string) {
	l.LiveLink = &liveLink
}

// SetAdvancedFrequency sets the AdvancedFrequency.
func (l *LiveWallSolutionSubmitV3Req) SetAdvancedFrequency(advancedFrequency videoCheck.AdvancedFrequencyRequest) {
	l.AdvancedFrequency = &advancedFrequency
}

// SetScreenMode sets the ScreenMode.
func (l *LiveWallSolutionSubmitV3Req) SetScreenMode(screenMode int) {
	l.ScreenMode = &screenMode
}

// SetCheckLanguageCode sets the CheckLanguageCode.
func (l *LiveWallSolutionSubmitV3Req) SetCheckLanguageCode(checkLanguageCode string) {
	l.CheckLanguageCode = &checkLanguageCode
}

// SetGender sets the Gender.
// SetGender sets the Gender.
func (l *LiveWallSolutionSubmitV3Req) SetGender(gender int) {
	l.Gender = &gender
}

// SetCheckSpeakerIds sets the CheckSpeakerIds.
func (l *LiveWallSolutionSubmitV3Req) SetCheckSpeakerIds(checkSpeakerIds []string) {
	l.CheckSpeakerIds = &checkSpeakerIds
}

// SetNoCheckSpeakerIds sets the NoCheckSpeakerIds.
func (l *LiveWallSolutionSubmitV3Req) SetNoCheckSpeakerIds(noCheckSpeakerIds []string) {
	l.NoCheckSpeakerIds = &noCheckSpeakerIds
}

// SetExtLon1 sets the ExtLon1.
func (l *LiveWallSolutionSubmitV3Req) SetExtLon1(extLon1 int64) {
	l.ExtLon1 = &extLon1
}

// SetExtLon2 sets the ExtLon2.
func (l *LiveWallSolutionSubmitV3Req) SetExtLon2(extLon2 int64) {
	l.ExtLon2 = &extLon2
}

// SetExtStr1 sets the ExtStr1.
func (l *LiveWallSolutionSubmitV3Req) SetExtStr1(extStr1 string) {
	l.ExtStr1 = &extStr1
}

// SetExtStr2 sets the ExtStr2.
func (l *LiveWallSolutionSubmitV3Req) SetExtStr2(extStr2 string) {
	l.ExtStr2 = &extStr2
}

// SetExtension sets the Extension.
func (l *LiveWallSolutionSubmitV3Req) SetExtension(extension string) {
	l.Extension = &extension
}

// SetScreenShotStrategy sets the ScreenShotStrategy.
func (l *LiveWallSolutionSubmitV3Req) SetScreenShotStrategy(screenShotStrategy int) {
	l.ScreenShotStrategy = &screenShotStrategy
}

// SetSubProduct sets the SubProduct.
func (l *LiveWallSolutionSubmitV3Req) SetSubProduct(subProduct string) {
	l.SubProduct = &subProduct
}

// SetAccountInfo sets the AccountInfo.
func (l *LiveWallSolutionSubmitV3Req) SetAccountInfo(accountInfo string) {
	l.AccountInfo = &accountInfo
}
