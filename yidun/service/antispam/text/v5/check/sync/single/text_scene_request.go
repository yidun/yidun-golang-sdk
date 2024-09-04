package single

/**
 * 文本检测请求基类
 */
type TextSceneRequest struct {
	DataId                *string   `json:"dataId,omitempty"`                             // 数据唯一标识，能够根据该值定位到该条数据
	Content               *string   `json:"content,omitempty"`                            // 用户发表内容，建议对内容中json、表情符、HTML标签、UBB标签等做过滤，只传递纯文本，以减少误判概率
	Title                 *string   `json:"title,omitempty"`                              // 内容标题，适用于贴子、博客的文章标题等场景，建议抄送，辅助机审策略精准调优
	DataType              *int      `json:"dataType,omitempty"`                           // 大小限制 子数据类型，与易盾内容安全服务约定即可
	Callback              *string   `json:"callback,omitempty"`                           // 数据回调参数，调用方根据业务情况自行设计，当调用文本结果获取接口-轮询模式或文本结果获取接口-推送模式时，该接口会原样返回该字段。
	PublishTime           *int64    `json:"publishTime,omitempty"`                        // 发表时间，UNIX 时间戳(毫秒值)
	CallbackUrl           *string   `json:"callbackUrl,omitempty"`                        // 人工审核结果回调通知到客户的URL
	Category              *string   `json:"category,omitempty"`                           // 来源，用于展示渠道名称，应用名称等
	Account               *string   `json:"account,omitempty"`                            // 用户唯一标识
	Phone                 *string   `json:"phone,omitempty"`                              // 手机号
	Nickname              *string   `json:"nickname,omitempty"`                           // 用户昵称
    CustomLangCodes       *string   `json:"customLangCodes,omitempty"`                    // 客户指定的语种类型 如果客户有传入则直接作为语种类型，不再进行检测; ","连接的字符串 语种类型列表参考{<a href="https://support.dun.163.com/documents/588434200783982592?docId=891095487301275648"/a>}
	Gender                *int      `json:"gender,omitempty"`                             // 性别，0: 未知，1: 男性，2: 女性
	Age                   *int      `json:"age,omitempty"`                                // 年龄，0: 未知
	Level                 *int      `json:"level,omitempty"`                              // 用户等级，0: 未知，1: 低，2: 中，3: 高
	RegisterTime          *int64    `json:"registerTime,omitempty"`                       // 注册时间，UNIX 时间戳(毫秒值)
	FriendNum             *int64    `json:"friendNum,omitempty"`                          // 好友数
	FansNum               *int64    `json:"fansNum,omitempty"`                            // 粉丝数
	IsPremiumUse          *int      `json:"isPremiumUse,omitempty"`                       // 是否付费用户，0为默认值，1为付费
	Role                  *string   `json:"role,omitempty" validate:"max=32"`             // 用户类型角色
	DeviceType            *int      `json:"deviceType,omitempty"`                         // 用户设备id的类型
	DeviceId              *string   `json:"deviceId,omitempty" validate:"max=128"`        // 用户设备 id
	Mac                   *string   `json:"mac,omitempty" validate:"max=64"`              // 用户设备mac信息
	Imei                  *string   `json:"imei,omitempty" validate:"max=64"`             // 国际移动设备识别码
	Idfa                  *string   `json:"idfa,omitempty" validate:"max=64"`             // iOS设备标识码
	Idfv                  *string   `json:"idfv,omitempty" validate:"max=64"`             // iOS设备标识码
	AppVersion            *string   `json:"appVersion,omitempty" validate:"max=32"`       // APP版本号
	ReceiveUid            *string   `json:"receiveUid,omitempty" validate:"max=64"`       // 接受消息的用户标识
	Relationship          *int      `json:"relationship,omitempty"`                       // 收发消息者好友关系，1接收人关注发送人，2发送人关注接收人，3互相关注，4互未关注
	GroupId               *string   `json:"groupId,omitempty" validate:"max=32"`          // 群聊id
	RoomId                *string   `json:"roomId,omitempty" validate:"max=32"`           // 聊天室/直播/游戏房间
	Topic                 *string   `json:"topic,omitempty" validate:"max=128"`           // 文章/帖子id
	CommentId             *string   `json:"commentId,omitempty" validate:"max=32"`        // 主评论id
	CommodityId           *string   `json:"commodityId,omitempty" validate:"max=32"`      // 商品id
	Ip                    *string   `json:"ip,omitempty" validate:"max=128"`              // 用户IP地址
	ReceiveIp             *string   `json:"receiveIp,omitempty" validate:"max=128"`       // 接收者IP地址
	RelatedKeys           *[]string `json:"relatedKeys,omitempty"`                        // 上下文关联key列表
	ExtStr1               *string   `json:"extStr1,omitempty" validate:"max=128"`         // 自定义扩展参数1
	ExtStr2               *string   `json:"extStr2,omitempty" validate:"max=128"`         // 自定义扩展参数2
	ExtLon1               *int64    `json:"extLon1,omitempty"`                            // 预留扩展long字段1
	ExtLon2               *int64    `json:"extLon2,omitempty"`                            // 预留扩展long字段2
	RiskControlToken      *string   `json:"riskControlToken,omitempty"`                   // 智能风控token
	RiskControlBusinessId *string   `json:"riskControlBusinessId,omitempty"`              // 智能风控businessId
	CensorExt             *string   `json:"censorExt,omitempty" validate:"max=1024"`      // 人审扩展字段，用于人审调度中心的规则匹配
	Extension             *string   `json:"extension,omitempty"`                          // 预留扩展字段
	SubProduct            *string   `json:"subProduct,omitempty"`                         // 子产品
	RelateContext         *string   `json:"relateContext,omitempty" validate:"max=10000"` // 关联内容
}

// 构建request
func NewTextSceneRequest() *TextSceneRequest {
	sceneRequest := &TextSceneRequest{}
	return sceneRequest
}

// 设置Account
func (t *TextSceneRequest) SetAccount(account string) {
	t.Account = &account
}

// 设置Age
func (t *TextSceneRequest) SetAge(age int) {
	t.Age = &age
}

// 设置AppVersion
func (t *TextSceneRequest) SetAppVersion(appVersion string) {
	t.AppVersion = &appVersion
}

// 设置Callback
func (t *TextSceneRequest) SetCallback(callback string) {
	t.Callback = &callback
}

// 设置CallbackUrl
func (t *TextSceneRequest) SetCallbackURL(callbackURL string) {
	t.CallbackUrl = &callbackURL
}

// 设置Category
func (t *TextSceneRequest) SetCategory(category string) {
	t.Category = &category
}

// 设置CensorExt
func (t *TextSceneRequest) SetCensorExt(censorExt string) {
	t.CensorExt = &censorExt
}

// 设置CommentId
func (t *TextSceneRequest) SetCommentID(commentID string) {
	t.CommentId = &commentID
}

// 设置CommodityId
func (t *TextSceneRequest) SetCommodityID(commodityID string) {
	t.CommodityId = &commodityID
}

// 设置Content
func (t *TextSceneRequest) SetContent(content string) {
	t.Content = &content
}

// 设置DataId
func (t *TextSceneRequest) SetDataID(dataID string) {
	t.DataId = &dataID
}

// 设置DataType
func (t *TextSceneRequest) SetDataType(dataType int) {
	t.DataType = &dataType
}

// 设置DeviceId
func (t *TextSceneRequest) SetDeviceID(deviceID string) {
	t.DeviceId = &deviceID
}

// 设置DeviceType
func (t *TextSceneRequest) SetDeviceType(deviceType int) {
	t.DeviceType = &deviceType
}

// 设置ExtLon1
func (t *TextSceneRequest) SetExtLon1(extLon1 int64) {
	t.ExtLon1 = &extLon1
}

// 设置ExtLon2
func (t *TextSceneRequest) SetExtLon2(extLon2 int64) {
	t.ExtLon2 = &extLon2
}

// 设置ExtStr1
func (t *TextSceneRequest) SetExtStr1(extStr1 string) {
	t.ExtStr1 = &extStr1
}

// 设置ExtStr2
func (t *TextSceneRequest) SetExtStr2(extStr2 string) {
	t.ExtStr2 = &extStr2
}

// 设置FansNum
func (t *TextSceneRequest) SetFansNum(fansNum int64) {
	t.FansNum = &fansNum
}

// 设置FriendNum
func (t *TextSceneRequest) SetFriendNum(friendNum int64) {
	t.FriendNum = &friendNum
}

// 设置Gender
func (t *TextSceneRequest) SetGender(gender int) {
	t.Gender = &gender
}

// 设置GroupId
func (t *TextSceneRequest) SetGroupID(groupID string) {
	t.GroupId = &groupID
}

// 设置Idfa
func (t *TextSceneRequest) SetIdfa(idfa string) {
	t.Idfa = &idfa
}

// 设置Idfv
func (t *TextSceneRequest) SetIdfv(idfv string) {
	t.Idfv = &idfv
}

// 设置Imei
func (t *TextSceneRequest) SetImei(imei string) {
	t.Imei = &imei
}

// 设置Ip
func (t *TextSceneRequest) SetIP(ip string) {
	t.Ip = &ip
}

// 设置ReceiveIp
func (t *TextSceneRequest) SetReceiveIp(receiveIp string) {
	t.ReceiveIp = &receiveIp
}

// 设置IsPremiumUse
func (t *TextSceneRequest) SetIsPremiumUse(isPremiumUse int) {
	t.IsPremiumUse = &isPremiumUse
}

// 设置Level
func (t *TextSceneRequest) SetLevel(level int) {
	t.Level = &level
}

// 设置Mac
func (t *TextSceneRequest) SetMac(mac string) {
	t.Mac = &mac
}

// 设置Nickname
func (t *TextSceneRequest) SetNickname(nickname string) {
	t.Nickname = &nickname
}

// 设置CustomLangCodes
func (t *TextSceneRequest) SetCustomLangCodes(customLangCodes string) {
    t.CustomLangCodes = &customLangCodes
}

// 设置Phone
func (t *TextSceneRequest) SetPhone(phone string) {
	t.Phone = &phone
}

// 设置PublishTime
func (t *TextSceneRequest) SetPublishTime(publishTime int64) {
	t.PublishTime = &publishTime
}

// 设置ReceiveUid
func (t *TextSceneRequest) SetReceiveUID(receiveUID string) {
	t.ReceiveUid = &receiveUID
}

// 设置RegisterTime
func (t *TextSceneRequest) SetRegisterTime(registerTime int64) {
	t.RegisterTime = &registerTime
}

// 设置RelatedKeys
func (t *TextSceneRequest) SetRelatedKeys(relatedKeys []string) {
	t.RelatedKeys = &relatedKeys
}

// 设置Relationship
func (t *TextSceneRequest) SetRelationship(relationship int) {
	t.Relationship = &relationship
}

// 设置RiskControlBusinessId
func (t *TextSceneRequest) SetRiskControlBusinessID(riskControlBusinessID string) {
	t.RiskControlBusinessId = &riskControlBusinessID
}

// 设置RiskControlToken
func (t *TextSceneRequest) SetRiskControlToken(riskControlToken string) {
	t.RiskControlToken = &riskControlToken
}

// 设置Role
func (t *TextSceneRequest) SetRole(role string) {
	t.Role = &role
}

// 设置RoomId
func (t *TextSceneRequest) SetRoomID(roomID string) {
	t.RoomId = &roomID
}

// 设置Title
func (t *TextSceneRequest) SetTitle(title string) {
	t.Title = &title
}

// 设置Topic
func (t *TextSceneRequest) SetTopic(topic string) {
	t.Topic = &topic
}

// 设置Extension
func (t *TextSceneRequest) SetExtension(extension string) {
	t.Extension = &extension
}

// 设置SubProduct
func (t *TextSceneRequest) SetSubProduct(subProduct string) {
	t.SubProduct = &subProduct
}

// 设置RelateContext
func (t *TextSceneRequest) SetRelateContext(relateContext string) {
	t.RelateContext = &relateContext
}
