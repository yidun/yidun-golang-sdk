---
# Spec Flow 审核人配置
# 提交技术方案或代码评审时会自动从此文件读取审核人

# 技术方案审核人 (用于 submit_tech_scheme.py)
tech_scheme_reviewers:
  - jinxiaotian01

# 代码评审审核人 (用于 submit_code_review.py)
code_reviewers:
  - jinxiaotian01

# 通用审核人 (如果未指定特定类型，使用此配置)
reviewers:
  - jinxiaotian01
---

# 审核人配置说明

本文件用于配置 Yidun Golang SDK 提交评审时的审核人。

## 配置格式

支持三种审核人类型：

1. **tech_scheme_reviewers** - 技术方案评审的审核人
2. **code_reviewers** - 代码评审的审核人
3. **reviewers** - 通用审核人（如果未配置特定类型，则使用此配置）

## 使用说明

1. 修改上方 YAML 部分的邮箱地址为实际的审核人
2. 邮箱格式必须是完整的 `username`
3. 可以配置多个审核人，每行一个

## 优先级

- 提交技术方案时：优先使用 `tech_scheme_reviewers`，如果未配置则使用 `reviewers`
- 提交代码评审时：优先使用 `code_reviewers`，如果未配置则使用 `reviewers`
