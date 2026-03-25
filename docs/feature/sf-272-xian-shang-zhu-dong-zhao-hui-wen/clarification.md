# 前置澄清记录

## 澄清时间
2026-03-16 10:40:00

## 澄清内容

| # | 问题 | 用户回复 |
|---|------|---------|
| 1 | 此次修复的范围是？ | 仅 Go SDK (Recommended) - 仅修复 Go SDK 的字段缺失 |
| 2 | 需要修复哪些优先级的问题？ | 全部 P0-P3 - 修复全部 18 个问题，包括类型差异 |
| 3 | 是否需要保持向后兼容性？ | 需要向后兼容 - 新增字段不破坏现有代码 |
| 4 | Go SDK 字段命名不一致如何处理？ | 保持现状 - 保持 Go SDK 现有命名（如 relatedContents） |

## 补充说明

基于用户确认，本次技术方案将：

1. **仅修复 Go SDK**：不涉及 Java SDK 的修改
2. **全量修复**：修复所有 P0-P3 问题，共 18 个问题
3. **向后兼容**：所有新增字段使用指针类型（可选字段），确保不破坏现有代码
4. **保持现有命名**：对于已存在的字段名差异（如 relatedContents），保持不变

## 需要修复的问题列表

### P0 - 严重缺失（整个模块缺失） - 3个
- P0-001: Go SDK缺少 riskControl 整个模块
- P0-002: Go SDK缺少 similarText 模块
- P0-003: Go SDK缺少 underage 模块

### P1 - 高优先级字段缺失 - 4个
- P1-001: antispam.censor 字段缺失（Java SDK 问题，Go SDK 无）
- P1-002: antispam.relateContent 字段缺失（保持现有 relatedContents 命名）
- P1-003: antispam.hitSource 字段缺失（保持现有 hitSources 命名）
- P1-004: aigcPrompt.proxyAnswerType 字段缺失

### P2 - 中优先级字段缺失 - 6个
- P2-001: antispam.hitType 字段缺失
- P2-002: antispam.strategyType 字段缺失
- P2-003: antispam.hitResult 字段缺失
- P2-004: labels[].hitType 字段缺失（Go SDK 已有，无需修复）
- P2-005: censorLabels[].depth 字段缺失（Go SDK 已有，无需修复）
- P2-006: libInfos[].strategyGroupName 和 strategyGroupId 字段缺失

### P3 - 低优先级/类型差异 - 5个
- P3-001: aigcPrompt.details Go SDK类型为指针切片，需要确认是否调整
- P3-002: userRisk 内部字段缺失（后端内部字段，不需要透出）
- P3-003~P3-005: 类型差异（int vs Integer），Go SDK 不涉及

## 实际需要修复的问题

经过筛选，Go SDK 实际需要修复：

**P0 (3个)**:
- riskControl 模块
- similarText 模块
- underage 模块

**P1 (2个，排除Java专属)**:
- aigcPrompt.proxyAnswerType

**P2 (4个，排除Go已有)**:
- antispam.hitType
- antispam.strategyType
- antispam.hitResult
- libInfos[].strategyGroupName/strategyGroupId

**P3 (1个)**:
- 确认 aigcPrompt.details 类型是否需要调整

**总计：约 10 个实际修复项**
