package response

// VideoCallbackFaceV4Response 点播视频回调人脸识别结构
type VideoCallbackFaceV4Response struct {
	TaskID   *string                    `json:"taskId,omitempty"`
	DataID   *string                    `json:"dataId,omitempty"`
	Pictures []*VideoDataFaceV4Response `json:"pictures,omitempty"`
}

type VideoDataFaceV4Response struct {
	URL       *string `json:"url,omitempty"`
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	// face检测结果
	Details   []*VideoDataFaceDetailBo `json:"details,omitempty"`
	PictureID *string                  `json:"pictureId,omitempty"`
}

type VideoDataFaceDetailBo struct {
	// 人脸数量
	FaceNumber *int `json:"faceNumber,omitempty"`
	// 是否有伪造人脸 1：有 0：无
	Deepfake *int `json:"deepfake,omitempty"`
	// ocr文本详情
	FaceContents []*FaceLineContent `json:"faceContents,omitempty"`
}

type FaceLineContent struct {
	// 人名
	Name *string `json:"name,omitempty"`
	// 坐标左上一个 右下一个
	X1 *float32 `json:"x1,omitempty"`
	Y1 *float32 `json:"y1,omitempty"`
	X2 *float32 `json:"x2,omitempty"`
	Y2 *float32 `json:"y2,omitempty"`
	// cartoon（动漫脸）、normal（真人脸）
	Type *string `json:"type,omitempty"`
	// star（明星）、normal（非明星）
	Category *string `json:"category,omitempty"`
	// 性别
	Gender *string `json:"gender,omitempty"`
	// 年龄
	Age *int `json:"age,omitempty"`
	// 人脸大小占比
	SizeRatio *string `json:"sizeRatio,omitempty"`
	// 人脸颜值分类
	BeautyScore *float32 `json:"beautyScore,omitempty"`
	// 人脸情绪
	Expression *string `json:"expression,omitempty"`
	// 人脸遮挡
	MaskType *string `json:"maskType,omitempty"`
	// 遮挡分数
	MaskScore *float32 `json:"maskScore,omitempty"`
	// 人脸戴眼镜
	Glasses *string `json:"glasses,omitempty"`
	// 成长阶段，对应枚举FaceAgeRangeEnum
	GrowthStage *string `json:"growthStage,omitempty"`
}
