# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added - 图片检测响应字段补齐

#### ImageV5Result 新增字段
- `LlmCheckInfo` ([]LlmCheckInfo) - 大模型检测结果列表

#### ImageV5AntispamResp 新增字段
- `SuggestionLevel` (int) - 建议级别
- `Censor` (string) - 审核人
- `OverAllMarkDesc` (string) - 整体审核备注
- `DetailMarks` ([]DetailMark) - 细节标注列表
- `Remark` (string) - 审核备注
- `CensorExtension` (CensorExtension) - 人审扩展字段
- `HitType` (int) - 命中策略类型
- `HitSource` (int) - 特征添加来源
- `StrategyType` (int) - 策略类型（1:公有策略 2:私有策略）
- `HitResult` (string) - 命中结果

#### ImageV5OcrResp 新增字段
- `FrameSize` (int) - OCR 分帧数

#### 新增结构体
- `LlmCheckInfo` - 大模型检测结果
  - `Label` (string) - 模型标签
  - `Explain` (string) - 模型解释
  - `Rate` (float64) - 模型分数
  - `ModelIdentifier` (string) - 大模型标识
- `DetailMark` - 细节标注结构体
  - `Position` ([]MarkPoint) - 标注位置
  - `CensorLabels` ([]CensorLabelInfo) - 标注标签列表
  - `Desc` (string) - 标注备注
- `MarkPoint` - 标注点坐标
  - `X` (float32) - x 坐标
  - `Y` (float32) - y 坐标
- `CensorExtension` - 人审扩展字段
  - `QualityInspectionTaskId` (string) - 质检任务ID

### Fixed
- 修复 Go SDK 响应字段与 API 不一致的问题 (#SF-166-261)
- 修正 `SuggestionRiskLevel` 和 `PublicOpinionInfo` 字段命名（大写开头导出）
- 修正 `CustomLabels` 字段的 JSON tag（添加 omitempty）
