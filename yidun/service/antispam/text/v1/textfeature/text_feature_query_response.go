package textfeature

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// TextFeatureQueryResponse 文本特征查询响应
type TextFeatureQueryResponse struct {
	*types.CommonResponse
	Result Result `json:"result"`
}

// Result 查询结果
type Result struct {
	Count int64             `json:"count"`
	Rows  []TextFeatureInfo `json:"rows"`
}

// TextFeatureInfo 文本特征信息
type TextFeatureInfo struct {
	Uuid        string `json:"uuid"`
	Content     string `json:"content"`
	FeatureType int    `json:"featureType"`
	Level       int    `json:"level"`
	MatchType   int    `json:"matchType"`
	SpamType    int    `json:"spamType"`
	SubLabel    string `json:"subLabel"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	CreateTime  int64  `json:"createTime"`
	EditTime    int64  `json:"editTime"`
	Scope       int    `json:"scope"`
}
