# 需求文档

> 创建人: jinxiaotian01
> 创建时间: 2026-03-13 10:02:55
> 版本: v1

---

缺失字段详见附件

<!DOCTYPE html><html lang="zh-CN"><head><meta charset="UTF-8"><title>SDK字段对比报告 - 文本检测(Text)回调接口</title><style>body{font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;line-height:1.6;max-width:1400px;margin:0 auto;padding:20px;color:#333}h1{font-size:2em;border-bottom:2px solid #eaecef;padding-bottom:.3em}h2{font-size:1.5em;border-bottom:1px solid #eaecef;padding-bottom:.3em}h3{font-size:1.25em}h4{font-size:1em}code{background:#f6f8fa;padding:2px 6px;border-radius:3px;font-family:monospace;font-size:.9em}pre{background:#f6f8fa;padding:16px;border-radius:6px;overflow-x:auto}pre code{background:transparent;padding:0}table{border-collapse:collapse;width:100%;margin:16px 0;display:block;overflow-x:auto;font-size:.9em}th,td{border:1px solid #dfe2e5;padding:8px 13px;text-align:left}th{background:#f6f8fa;font-weight:600}tr:nth-child(even){background:#f9f9f9}hr{border:0;border-top:2px solid #eaecef;margin:24px 0}a{color:#0366d6}li{margin:4px 0}p{margin:8px 0}blockquote{border-left:4px solid #dfe2e5;padding-left:16px;color:#6a737d;margin:16px 0}</style></head><body><h1>SDK字段对比报告 - 文本检测(Text)回调接口</h1>
<br>
<p><strong>生成时间:</strong> 2026-03-10 21:39:13</p>
<br>
<p><strong>对比规则:</strong> ignore_case=false, ignore_field_order=true, ignore_annotations=true, compare_field_types=true, recursive_compare=true, max_recursion_depth=15</p>
<br>
<hr>
<br>
<h2>一、结构概览</h2>
<br>
<h3>1.1 顶层响应结构</h3>
<br>
<pre><code>
Java后端:    TextCheckResultBodyV5 (extends TextCheckResultBody)
Java SDK:   TextCallBackResponse (extends CommonResponse) -&gt; result: List&lt;TextCheckResult&gt;
Go SDK:     TextCallBackResponse -&gt; *types.CommonResponse + Result []*single.TextCheckResult
</code></pre>
<br>
<h3>1.2 字段树形结构</h3>
<br>
<pre><code>
TextCheckResult / TextCheckResultBodyV5
├── antispam (Antispam / TextCheckAntispamResultV5)
│   ├── taskId: String
│   ├── dataId: String
│   ├── label: Integer/int
│   ├── secondLabel: String
│   ├── thirdLabel: String
│   ├── riskDescription: String
│   ├── suggestion: Integer/int
│   ├── suggestionLevel: Integer
│   ├── suggestionRiskLevel: Integer
│   ├── publicOpinionInfo: String
│   ├── customAction: Integer
│   ├── resultType: Integer
│   ├── censorType: Integer
│   ├── callback: String
│   ├── censorSource: Integer
│   ├── censorRound: Integer
│   ├── censorTime: Long/int64
│   ├── isRelatedHit: Boolean/bool
│   ├── relatedHitType: Integer
│   ├── remark: String
│   ├── filteredContent: String
│   ├── mergeHints: List&lt;String&gt;
│   ├── status: Integer
│   ├── censor: String (后端/Go SDK)
│   ├── relateContent/relatedContents: String
│   ├── hitSource/hitSources: Integer
│   ├── hitType: Integer (后端独有)
│   ├── strategyType: Integer (后端独有)
│   ├── hitResult: String (后端独有)
│   ├── censorLabels: List&lt;CensorLabel/CensorLabelInfo&gt;
│   │   ├── code: String
│   │   ├── name: String
│   │   ├── desc: String
│   │   ├── customCode: String
│   │   ├── parentLabelId: String
│   │   └── depth: Integer
│   ├── strategyVersions: List&lt;StrategyVersion/TextVersionResp&gt;
│   │   ├── label: Integer
│   │   └── version: String
│   ├── customLabels: List&lt;CustomLabel/CustomLabelInfo&gt;
│   │   ├── name: String
│   │   ├── code: String
│   │   └── depth: Integer
│   ├── censorExtension: CensorExtension
│   │   ├── qualityInspectionTaskId: String
│   │   └── inspTaskCreateTime: Long/int64
│   └── labels: List&lt;AntispamLabel/TextCheckResultLabelV5&gt;
│       ├── label: Integer/int
│       ├── level: Integer/int
│       ├── rate: Double/float64
│       ├── hitType: Integer (SDK独有)
│       └── subLabels: List&lt;AntispamSubLabel/TextSubLabelV5&gt;
│           ├── subLabel: String
│           ├── subLabelDepth: Integer
│           ├── secondLabel: String
│           ├── thirdLabel: String
│           ├── riskDescription: String
│           ├── suggestionRiskLevel: Integer
│           ├── rate: Double/float64
│           ├── politicalSentiment: Integer
│           ├── isRelatedLabel: Boolean/bool
│           └── details: AntispamSubLabelDetail/TextCheckResultLabelDetailV5
│               ├── keywords: List&lt;AntispamSubLabelDetailKeyword&gt;
│               │   ├── word: String
│               │   ├── strategyGroupName: String
│               │   └── strategyGroupId: Long/int64
│               ├── libInfos: List&lt;AntispamSubLabelDetailLibInfo/TextCheckResultLabelDetailLibInfoV5&gt;
│               │   ├── type: Integer
│               │   ├── entity: String
│               │   ├── releaseTime: Long/int64
│               │   ├── strategyGroupName: String (后端独有)
│               │   └── strategyGroupId: Long (后端独有)
│               ├── rules: List&lt;AntispamSubLabelDetailRule&gt;
│               │   └── name: String
│               ├── anticheat: AntispamSubLabelDetailAnticheat
│               │   └── type: Integer
│               └── hitInfos: List&lt;AntispamSubLabelDetailHitInfo&gt;
│                   ├── value: String
│                   ├── type: Integer
│                   └── positions: List&lt;AntispamSubLabelDetailHitInfoPosition&gt;
│                       ├── fieldName: String
│                       ├── startPos: Integer
│                       └── endPos: Integer
├── emotionAnalysis (EmotionAnalysis / TextCheckEmotionAnalysisResultV5)
│   ├── taskId: String
│   ├── dataId: String
│   └── details: List&lt;EmotionAnalysisDetail&gt;
│       ├── positiveProb: Double/float64
│       ├── negativeProb: Double/float64
│       └── sentiment: String
├── anticheat (Anticheat / TextCheckAnticheatResultV5)
│   ├── taskId: String
│   ├── dataId: String
│   └── details: List&lt;AnticheatDetail&gt;
│       ├── suggestion: Integer
│       └── hitInfos: List&lt;AnticheatDetailHitInfo&gt;
│           ├── hitType: Integer
│           └── hitMsg: String
├── userRisk (UserRisk / TextCheckUserRiskResultV5)
│   ├── taskId: String
│   ├── dataId: String
│   ├── details: List&lt;UserRiskDetail&gt;
│   │   ├── account: String
│   │   ├── accountLevel: Integer
│   │   └── acDetails: List&lt;UserRiskAcDetail&gt;
│   │       ├── riskType: String
│   │       ├── riskLevel: Integer
│   │       └── riskScore: Double/float64
│   ├── accountDetail: RiskPortraitDetail.Detail (后端独有，内部字段)
│   ├── phoneDetail: RiskPortraitDetail.Detail (后端独有，内部字段)
│   └── ipDetail: RiskPortraitDetail.Detail (后端独有，内部字段)
├── language (Language / TextCheckLanguageResultV5)
│   ├── taskId: String
│   ├── dataId: String
│   └── details: List&lt;LanguageDetail&gt;
│       └── type: String
├── riskControl (RiskControl / TextCheckRiskControlResultV5)
│   ├── taskId: String
│   ├── dataId: String
│   └── details: List&lt;RiskControlDetail&gt;
│       ├── riskLevel: Integer
│       └── hitInfos: List&lt;HitInfo/RiskControlCheckerResult.HitInfo&gt;
│           ├── type: String
│           ├── name: String
│           └── desc: String
├── aigcPrompt (AigcPrompt / AigcPromptAnalysisResultV5)
│   ├── taskId: String
│   ├── dataId: String
│   ├── proxyAnswerType: Integer (后端独有)
│   └── details: List&lt;AigcPromptDetail&gt;
│       ├── type: Integer
│       ├── answer: String
│       ├── source: Integer
│       ├── libId: String
│       └── answerId: String
├── llmCheckInfo (LlmCheckInfo / LlmCheckResultV5)
│   ├── taskId: String
│   ├── dataId: String
│   └── details: List&lt;LlmCheckInfoDetail&gt;
│       ├── modelIdentifier: String
│       ├── label: String
│       └── explain: String
├── similarText (SimilarTextResultV5) [后端独有]
│   ├── taskId: String
│   ├── dataId: String
│   └── details: List&lt;SimilarTextResultDetailV5&gt;
│       ├── matchDataId: String
│       └── rate: Double
└── underage (TextCheckUnderageResultV5) [后端独有]
    ├── taskId: String
    ├── dataId: String
    └── details: List&lt;TextCheckUnderageDetailV5&gt;
        └── type: Integer
</code></pre>
<br>
<hr>
<br>
<h2>二、顶层响应结构对比</h2>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>(响应根)</td>
<td>TextCheckResultBodyV5</td>
<td>TextCallBackResponse</td>
<td>TextCallBackResponse</td>
<td>-</td>
<td>命名不同</td>
</tr>
<tr>
<td>code</td>
<td>无(继承自BaseResponse)</td>
<td>int (CommonResponse)</td>
<td>int (CommonResponse)</td>
<td>SDK一致</td>
<td>后端无此字段</td>
</tr>
<tr>
<td>msg</td>
<td>无</td>
<td>String (CommonResponse)</td>
<td>string (CommonResponse)</td>
<td>SDK一致</td>
<td>后端无此字段</td>
</tr>
<tr>
<td>result</td>
<td>无(自身即result)</td>
<td>List\<TextCheckResult\></td>
<td>[]*TextCheckResult</td>
<td>SDK一致</td>
<td>后端是单体对象</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h2>三、主模块字段对比</h2>
<br>
<h3>3.1 TextCheckResult 顶层字段</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>antispam</td>
<td>TextCheckAntispamResultV5</td>
<td>Antispam (内部类)</td>
<td>*Antispam</td>
<td>一致</td>
<td>类型名不同但结构对应</td>
</tr>
<tr>
<td>emotionAnalysis</td>
<td>TextCheckEmotionAnalysisResultV5</td>
<td>EmotionAnalysis (内部类)</td>
<td>*EmotionAnalysis</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>anticheat</td>
<td>TextCheckAnticheatResultV5</td>
<td>Anticheat (内部类)</td>
<td>*Anticheat</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>userRisk</td>
<td>TextCheckUserRiskResultV5</td>
<td>UserRisk (内部类)</td>
<td>*UserRisk</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>language</td>
<td>TextCheckLanguageResultV5</td>
<td>Language (内部类)</td>
<td>*Language</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>riskControl</td>
<td>TextCheckRiskControlResultV5</td>
<td>RiskControl (内部类)</td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>Go SDK缺少riskControl字段</strong></td>
</tr>
<tr>
<td>aigcPrompt</td>
<td>AigcPromptAnalysisResultV5</td>
<td>AigcPrompt (内部类)</td>
<td>*AigcPrompt</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>llmCheckInfo</td>
<td>LlmCheckResultV5</td>
<td>LlmCheckInfo (内部类)</td>
<td>*LlmCheckInfo</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>similarText</td>
<td>SimilarTextResultV5</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>Java SDK和Go SDK均缺少</strong></td>
</tr>
<tr>
<td>underage</td>
<td>TextCheckUnderageResultV5</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>Java SDK和Go SDK均缺少</strong></td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.2 antispam 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>antispam.taskId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.dataId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.label</td>
<td>int</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>后端基本类型vs包装类型</td>
</tr>
<tr>
<td>antispam.secondLabel</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.thirdLabel</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.riskDescription</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.suggestion</td>
<td>int</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.suggestionLevel</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.suggestionRiskLevel</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.publicOpinionInfo</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.customAction</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.resultType</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.censorType</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.callback</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.censorSource</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.censorRound</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.censorTime</td>
<td>Long</td>
<td>Long</td>
<td>*int64</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.isRelatedHit</td>
<td>Boolean</td>
<td>Boolean</td>
<td>*bool</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.relatedHitType</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.remark</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.filteredContent</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.mergeHints</td>
<td>List\<String\></td>
<td>List\<String\></td>
<td>[]*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.status</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>antispam.censor</td>
<td>String</td>
<td><strong>缺失</strong></td>
<td>*string</td>
<td><strong>不一致</strong></td>
<td><strong>Java SDK缺少censor字段</strong></td>
</tr>
<tr>
<td>antispam.relateContent</td>
<td>String</td>
<td><strong>缺失</strong></td>
<td>*string (relatedContents)</td>
<td><strong>不一致</strong></td>
<td><strong>Java SDK缺失；Go字段名relatedContents</strong></td>
</tr>
<tr>
<td>antispam.hitSource</td>
<td>Integer</td>
<td><strong>缺失</strong></td>
<td>*int (hitSources)</td>
<td><strong>不一致</strong></td>
<td><strong>Java SDK缺失；Go字段名hitSources</strong></td>
</tr>
<tr>
<td>antispam.hitType</td>
<td>Integer</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>SDK均缺少，后端独有</strong></td>
</tr>
<tr>
<td>antispam.strategyType</td>
<td>Integer</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>SDK均缺少，后端独有</strong></td>
</tr>
<tr>
<td>antispam.hitResult</td>
<td>String</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>SDK均缺少，后端独有</strong></td>
</tr>
<tr>
<td>antispam.censorLabels</td>
<td>List\<CensorLabelInfo\></td>
<td>List\<CensorLabel\></td>
<td>[]*CensorLabel</td>
<td>一致(结构)</td>
<td>类名不同，见3.2.1</td>
</tr>
<tr>
<td>antispam.strategyVersions</td>
<td>List\<TextVersionResp\></td>
<td>List\<StrategyVersion\></td>
<td>[]*StrategyVersion</td>
<td>一致(结构)</td>
<td>类名不同，见3.2.2</td>
</tr>
<tr>
<td>antispam.customLabels</td>
<td>List\<CustomLabelInfo\></td>
<td>List\<CustomLabel\></td>
<td>[]*CustomLabel</td>
<td>一致(结构)</td>
<td>类名不同，见3.2.3</td>
</tr>
<tr>
<td>antispam.censorExtension</td>
<td>CensorExtension</td>
<td>CensorExtension</td>
<td>*CensorExtension</td>
<td>一致(结构)</td>
<td>见3.2.4</td>
</tr>
<tr>
<td>antispam.labels</td>
<td>List\<TextCheckResultLabelV5\></td>
<td>List\<AntispamLabel\></td>
<td>[]*AntispamLabel</td>
<td>一致(结构)</td>
<td>见3.3</td>
</tr>
</tbody></table>
<br>
<h4>3.2.1 censorLabels 内部字段对比</h4>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>censorLabels[].code</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>censorLabels[].name</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>censorLabels[].desc</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>censorLabels[].customCode</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>censorLabels[].parentLabelId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>censorLabels[].depth</td>
<td>Integer</td>
<td><strong>缺失</strong></td>
<td>*int</td>
<td><strong>不一致</strong></td>
<td><strong>Java SDK缺少depth字段</strong></td>
</tr>
</tbody></table>
<br>
<h4>3.2.2 strategyVersions 内部字段对比</h4>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>strategyVersions[].label</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>strategyVersions[].version</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<h4>3.2.3 customLabels 内部字段对比</h4>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>customLabels[].name</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>customLabels[].code</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>customLabels[].depth</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<h4>3.2.4 censorExtension 内部字段对比</h4>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>censorExtension.qualityInspectionTaskId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>censorExtension.inspTaskCreateTime</td>
<td>Long</td>
<td>Long</td>
<td>*int64</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.3 antispam.labels (AntispamLabel) 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>labels[].label</td>
<td>int</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>labels[].level</td>
<td>int</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>labels[].rate</td>
<td>Double</td>
<td>Double</td>
<td>*float64</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>labels[].hitType</td>
<td>Integer</td>
<td><strong>缺失</strong></td>
<td>*int</td>
<td><strong>不一致</strong></td>
<td><strong>Java SDK缺少hitType字段</strong></td>
</tr>
<tr>
<td>labels[].subLabels</td>
<td>List\<TextSubLabelV5\></td>
<td>List\<AntispamSubLabel\></td>
<td>[]*AntispamSubLabel</td>
<td>一致(结构)</td>
<td>见3.4</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.4 antispam.labels[].subLabels (AntispamSubLabel) 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>subLabels[].subLabel</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>subLabels[].subLabelDepth</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>subLabels[].secondLabel</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>subLabels[].thirdLabel</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>subLabels[].riskDescription</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>subLabels[].suggestionRiskLevel</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>subLabels[].rate</td>
<td>Double</td>
<td>Double</td>
<td>*float64</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>subLabels[].politicalSentiment</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>subLabels[].isRelatedLabel</td>
<td>Boolean</td>
<td>Boolean</td>
<td>*bool</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>subLabels[].censorStatus</td>
<td>Integer(@JsonIgnore)</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td>一致</td>
<td>后端@JsonIgnore不输出，SDK无需</td>
</tr>
<tr>
<td>subLabels[].spamType</td>
<td>Integer(@JsonIgnore)</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td>一致</td>
<td>后端@JsonIgnore不输出，SDK无需</td>
</tr>
<tr>
<td>subLabels[].details</td>
<td>TextCheckResultLabelDetailV5</td>
<td>AntispamSubLabelDetail</td>
<td>*AntispamSubLabelDetail</td>
<td>一致(结构)</td>
<td>见3.5</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.5 subLabels[].details (AntispamSubLabelDetail) 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>details.keywords</td>
<td>List\<TextCheckResultLabelDetailKeywordV5\></td>
<td>List\<AntispamSubLabelDetailKeyword\></td>
<td>[]*AntispamSubLabelDetailKeyword</td>
<td>一致(结构)</td>
<td>见3.5.1</td>
</tr>
<tr>
<td>details.libInfos</td>
<td>List\<TextCheckResultLabelDetailLibInfoV5\></td>
<td>List\<AntispamSubLabelDetailLibInfo\></td>
<td>[]*AntispamSubLabelDetailLibInfo</td>
<td>部分不一致</td>
<td>见3.5.2</td>
</tr>
<tr>
<td>details.rules</td>
<td>List\<TextCheckResultLabelDetailRuleV5\></td>
<td>List\<AntispamSubLabelDetailRule\></td>
<td>[]*AntispamSubLabelDetailRule</td>
<td>一致</td>
<td>见3.5.3</td>
</tr>
<tr>
<td>details.anticheat</td>
<td>TextCheckResultLabelDetailAnticheatV5</td>
<td>AntispamSubLabelDetailAnticheat</td>
<td>*AntispamSubLabelDetailAnticheat</td>
<td>一致</td>
<td>见3.5.4</td>
</tr>
<tr>
<td>details.hitInfos</td>
<td>List\<TextCheckResultLabelDetailHitInfoV5\></td>
<td>List\<AntispamSubLabelDetailHitInfo\></td>
<td>[]*AntispamSubLabelDetailHitInfo</td>
<td>一致(结构)</td>
<td>见3.5.5</td>
</tr>
</tbody></table>
<br>
<h4>3.5.1 details.keywords 内部字段对比</h4>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>keywords[].word</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>keywords[].strategyGroupName</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>keywords[].strategyGroupId</td>
<td>Long</td>
<td>Long</td>
<td>*int64</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<h4>3.5.2 details.libInfos 内部字段对比</h4>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>libInfos[].type</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>libInfos[].entity</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>libInfos[].releaseTime</td>
<td>Long</td>
<td>Long</td>
<td>*int64</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>libInfos[].strategyGroupName</td>
<td>String</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>SDK均缺少，后端独有</strong></td>
</tr>
<tr>
<td>libInfos[].strategyGroupId</td>
<td>Long</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>SDK均缺少，后端独有</strong></td>
</tr>
</tbody></table>
<br>
<h4>3.5.3 details.rules 内部字段对比</h4>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>rules[].name</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<h4>3.5.4 details.anticheat 内部字段对比</h4>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>anticheat.type</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<h4>3.5.5 details.hitInfos 内部字段对比</h4>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>hitInfos[].value</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>hitInfos[].type</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>hitInfos[].positions</td>
<td>List\<TextCheckResultLabelDetailHitInfoPositionV5\></td>
<td>List\<AntispamSubLabelDetailHitInfoPosition\></td>
<td>[]*AntispamSubLabelDetailHitInfoPosition</td>
<td>一致(结构)</td>
<td>见3.5.5.1</td>
</tr>
</tbody></table>
<br>
<h5>3.5.5.1 details.hitInfos[].positions 内部字段对比</h5>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>positions[].fieldName</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>positions[].startPos</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>positions[].endPos</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.6 emotionAnalysis 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>emotionAnalysis.taskId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>emotionAnalysis.dataId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>emotionAnalysis.details</td>
<td>List\<TextCheckResultEmotionAnalysisDetailV5\></td>
<td>List\<EmotionAnalysisDetail\></td>
<td>[]*EmotionAnalysisDetail</td>
<td>一致(结构)</td>
<td>-</td>
</tr>
<tr>
<td>emotionAnalysis.details[].positiveProb</td>
<td>Double</td>
<td>Double</td>
<td>*float64</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>emotionAnalysis.details[].negativeProb</td>
<td>Double</td>
<td>Double</td>
<td>*float64</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>emotionAnalysis.details[].sentiment</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.7 anticheat 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>anticheat.taskId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>anticheat.dataId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>anticheat.details</td>
<td>List\<TextCheckResultAnticheatDetailV5\></td>
<td>List\<AnticheatDetail\></td>
<td>[]*AnticheatDetail</td>
<td>一致(结构)</td>
<td>-</td>
</tr>
<tr>
<td>anticheat.details[].suggestion</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>anticheat.details[].hitInfos</td>
<td>List\<TextCheckResultAnticheatDetailHitInfoV5\></td>
<td>List\<AnticheatDetailHitInfo\></td>
<td>[]*AnticheatDetailHitInfo</td>
<td>一致(结构)</td>
<td>-</td>
</tr>
<tr>
<td>anticheat.details[].hitInfos[].hitType</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>anticheat.details[].hitInfos[].hitMsg</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.8 userRisk 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>userRisk.taskId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>userRisk.dataId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>userRisk.details</td>
<td>List\<TextCheckResultUserRiskDetailV5\></td>
<td>List\<UserRiskDetail\></td>
<td>[]*UserRiskDetail</td>
<td>一致(结构)</td>
<td>-</td>
</tr>
<tr>
<td>userRisk.details[].account</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>userRisk.details[].accountLevel</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>userRisk.details[].acDetails</td>
<td>List\<AcDetail\></td>
<td>List\<UserRiskAcDetail\></td>
<td>[]*UserRiskAcDetail</td>
<td>一致(结构)</td>
<td>类名不同</td>
</tr>
<tr>
<td>userRisk.details[].acDetails[].riskType</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>userRisk.details[].acDetails[].riskLevel</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>userRisk.details[].acDetails[].riskScore</td>
<td>Double</td>
<td>Double</td>
<td>*float64</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>userRisk.accountDetail</td>
<td>RiskPortraitDetail.Detail</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>后端内部字段，SDK不透出</strong></td>
</tr>
<tr>
<td>userRisk.phoneDetail</td>
<td>RiskPortraitDetail.Detail</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>后端内部字段，SDK不透出</strong></td>
</tr>
<tr>
<td>userRisk.ipDetail</td>
<td>RiskPortraitDetail.Detail</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>后端内部字段，SDK不透出</strong></td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.9 language 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>language.taskId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>language.dataId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>language.details</td>
<td>List\<TextCheckResultLanguageDetailV5\></td>
<td>List\<LanguageDetail\></td>
<td>[]*LanguageDetail</td>
<td>一致(结构)</td>
<td>-</td>
</tr>
<tr>
<td>language.details[].type</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.10 riskControl 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>riskControl</td>
<td>TextCheckRiskControlResultV5</td>
<td>RiskControl</td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>Go SDK缺少整个riskControl字段</strong></td>
</tr>
<tr>
<td>riskControl.taskId</td>
<td>String</td>
<td>String</td>
<td>-</td>
<td>后端/Java SDK一致</td>
<td>-</td>
</tr>
<tr>
<td>riskControl.dataId</td>
<td>String</td>
<td>String</td>
<td>-</td>
<td>后端/Java SDK一致</td>
<td>-</td>
</tr>
<tr>
<td>riskControl.details</td>
<td>List\<TextCheckResultRiskControlDetailV5\></td>
<td>List\<RiskControlDetail\></td>
<td>-</td>
<td>后端/Java SDK一致(结构)</td>
<td>-</td>
</tr>
<tr>
<td>riskControl.details[].riskLevel</td>
<td>Integer</td>
<td>Integer</td>
<td>-</td>
<td>后端/Java SDK一致</td>
<td>-</td>
</tr>
<tr>
<td>riskControl.details[].hitInfos</td>
<td>List\<RiskControlCheckerResult.HitInfo\></td>
<td>List\<HitInfo\></td>
<td>-</td>
<td>后端/Java SDK一致(结构)</td>
<td>HitInfo来自不同包</td>
</tr>
<tr>
<td>riskControl.details[].hitInfos[].type</td>
<td>String</td>
<td>String</td>
<td>-</td>
<td>后端/Java SDK一致</td>
<td>-</td>
</tr>
<tr>
<td>riskControl.details[].hitInfos[].name</td>
<td>String</td>
<td>String</td>
<td>-</td>
<td>后端/Java SDK一致</td>
<td>-</td>
</tr>
<tr>
<td>riskControl.details[].hitInfos[].desc</td>
<td>String</td>
<td>String</td>
<td>-</td>
<td>后端/Java SDK一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.11 aigcPrompt 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>aigcPrompt.taskId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>aigcPrompt.dataId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>aigcPrompt.proxyAnswerType</td>
<td>Integer</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>SDK均缺少，后端独有</strong></td>
</tr>
<tr>
<td>aigcPrompt.details</td>
<td>List\<AigcPromptAnalysisResultDetailV5\></td>
<td>List\<AigcPromptDetail\></td>
<td>*[]*AigcPromptDetail</td>
<td><strong>类型差异</strong></td>
<td><strong>Go SDK为指针切片，Java为切片</strong></td>
</tr>
<tr>
<td>aigcPrompt.details[].type</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>aigcPrompt.details[].answer</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>aigcPrompt.details[].source</td>
<td>Integer</td>
<td>Integer</td>
<td>*int</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>aigcPrompt.details[].libId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>aigcPrompt.details[].answerId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.12 llmCheckInfo 字段对比</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>llmCheckInfo.taskId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>llmCheckInfo.dataId</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>llmCheckInfo.details</td>
<td>List\<LlmCheckResultDetailV5\></td>
<td>List\<LlmCheckInfoDetail\></td>
<td>[]*LlmCheckInfoDetail</td>
<td>一致(结构)</td>
<td>-</td>
</tr>
<tr>
<td>llmCheckInfo.details[].modelIdentifier</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>llmCheckInfo.details[].label</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
<tr>
<td>llmCheckInfo.details[].explain</td>
<td>String</td>
<td>String</td>
<td>*string</td>
<td>一致</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.13 similarText 字段对比 (后端独有)</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>similarText</td>
<td>SimilarTextResultV5</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>Java SDK和Go SDK均缺少整个模块</strong></td>
</tr>
<tr>
<td>similarText.taskId</td>
<td>String</td>
<td>-</td>
<td>-</td>
<td>-</td>
<td>-</td>
</tr>
<tr>
<td>similarText.dataId</td>
<td>String</td>
<td>-</td>
<td>-</td>
<td>-</td>
<td>-</td>
</tr>
<tr>
<td>similarText.details[].matchDataId</td>
<td>String</td>
<td>-</td>
<td>-</td>
<td>-</td>
<td>-</td>
</tr>
<tr>
<td>similarText.details[].rate</td>
<td>Double</td>
<td>-</td>
<td>-</td>
<td>-</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h3>3.14 underage 字段对比 (后端独有)</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>字段路径</th>
<th>Java后端类型</th>
<th>Java SDK类型</th>
<th>Go SDK类型</th>
<th>一致性</th>
<th>备注</th>
</tr></thead><tbody>
<tr>
<td>underage</td>
<td>TextCheckUnderageResultV5</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td><strong>不一致</strong></td>
<td><strong>Java SDK和Go SDK均缺少整个模块</strong></td>
</tr>
<tr>
<td>underage.taskId</td>
<td>String</td>
<td>-</td>
<td>-</td>
<td>-</td>
<td>-</td>
</tr>
<tr>
<td>underage.dataId</td>
<td>String</td>
<td>-</td>
<td>-</td>
<td>-</td>
<td>-</td>
</tr>
<tr>
<td>underage.details[].type</td>
<td>Integer</td>
<td>-</td>
<td>-</td>
<td>-</td>
<td>-</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h2>四、问题汇总</h2>
<br>
<h3>P0 - 严重缺失（整个模块缺失）</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>问题编号</th>
<th>问题描述</th>
<th>影响范围</th>
<th>后端字段路径</th>
<th>SDK状态</th>
</tr></thead><tbody>
<tr>
<td>P0-001</td>
<td>Go SDK缺少riskControl整个模块</td>
<td>Go SDK</td>
<td>TextCheckResult.riskControl</td>
<td>Go SDK缺失</td>
</tr>
<tr>
<td>P0-002</td>
<td>Java SDK和Go SDK均缺少similarText模块</td>
<td>两个SDK</td>
<td>TextCheckResultBodyV5.similarText</td>
<td>双SDK缺失</td>
</tr>
<tr>
<td>P0-003</td>
<td>Java SDK和Go SDK均缺少underage模块</td>
<td>两个SDK</td>
<td>TextCheckResultBodyV5.underage</td>
<td>双SDK缺失</td>
</tr>
</tbody></table>
<br>
<h3>P1 - 高优先级字段缺失</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>问题编号</th>
<th>问题描述</th>
<th>影响范围</th>
<th>后端字段路径</th>
<th>SDK状态</th>
</tr></thead><tbody>
<tr>
<td>P1-001</td>
<td>Java SDK缺少antispam.censor字段</td>
<td>Java SDK</td>
<td>antispam.censor</td>
<td>Java SDK缺失</td>
</tr>
<tr>
<td>P1-002</td>
<td>antispam.relateContent字段Java SDK缺失，Go SDK字段名不一致(relatedContents)</td>
<td>两个SDK</td>
<td>antispam.relateContent</td>
<td>Java SDK缺失，Go SDK字段名为relatedContents</td>
</tr>
<tr>
<td>P1-003</td>
<td>antispam.hitSource字段Java SDK缺失，Go SDK字段名不一致(hitSources)</td>
<td>两个SDK</td>
<td>antispam.hitSource</td>
<td>Java SDK缺失，Go SDK字段名为hitSources</td>
</tr>
<tr>
<td>P1-004</td>
<td>aigcPrompt.proxyAnswerType字段SDK均缺失</td>
<td>两个SDK</td>
<td>aigcPrompt.proxyAnswerType</td>
<td>双SDK缺失</td>
</tr>
</tbody></table>
<br>
<h3>P2 - 中优先级字段缺失</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>问题编号</th>
<th>问题描述</th>
<th>影响范围</th>
<th>后端字段路径</th>
<th>SDK状态</th>
</tr></thead><tbody>
<tr>
<td>P2-001</td>
<td>antispam.hitType字段SDK均缺失</td>
<td>两个SDK</td>
<td>antispam.hitType</td>
<td>双SDK缺失(后端内部字段)</td>
</tr>
<tr>
<td>P2-002</td>
<td>antispam.strategyType字段SDK均缺失</td>
<td>两个SDK</td>
<td>antispam.strategyType</td>
<td>双SDK缺失(后端内部字段)</td>
</tr>
<tr>
<td>P2-003</td>
<td>antispam.hitResult字段SDK均缺失</td>
<td>两个SDK</td>
<td>antispam.hitResult</td>
<td>双SDK缺失(后端内部字段)</td>
</tr>
<tr>
<td>P2-004</td>
<td>labels[].hitType字段Java SDK缺失</td>
<td>Java SDK</td>
<td>antispam.labels[].hitType</td>
<td>Java SDK缺失，Go SDK有</td>
</tr>
<tr>
<td>P2-005</td>
<td>censorLabels[].depth字段Java SDK缺失</td>
<td>Java SDK</td>
<td>antispam.censorLabels[].depth</td>
<td>Java SDK缺失，Go SDK有</td>
</tr>
<tr>
<td>P2-006</td>
<td>libInfos[].strategyGroupName和strategyGroupId SDK均缺失</td>
<td>两个SDK</td>
<td>antispam.labels[].subLabels[].details.libInfos[].strategyGroupName/strategyGroupId</td>
<td>双SDK缺失</td>
</tr>
</tbody></table>
<br>
<h3>P3 - 低优先级/类型差异</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>问题编号</th>
<th>问题描述</th>
<th>影响范围</th>
<th>字段路径</th>
<th>说明</th>
</tr></thead><tbody>
<tr>
<td>P3-001</td>
<td>aigcPrompt.details Go SDK类型为指针切片(*[]*AigcPromptDetail) vs Java的切片</td>
<td>Go SDK</td>
<td>aigcPrompt.details</td>
<td>Go SDK使用指针切片，JSON序列化行为可能不同</td>
</tr>
<tr>
<td>P3-002</td>
<td>userRisk.accountDetail/phoneDetail/ipDetail后端有但SDK无</td>
<td>两个SDK</td>
<td>userRisk.accountDetail/phoneDetail/ipDetail</td>
<td>后端内部字段，SDK不透出，属正常设计</td>
</tr>
<tr>
<td>P3-003</td>
<td>antispam.label后端为int基本类型，SDK为Integer包装类型</td>
<td>后端vs SDK</td>
<td>antispam.label</td>
<td>Java类型差异，功能等价</td>
</tr>
<tr>
<td>P3-004</td>
<td>antispam.suggestion后端为int基本类型，SDK为Integer包装类型</td>
<td>后端vs SDK</td>
<td>antispam.suggestion</td>
<td>Java类型差异，功能等价</td>
</tr>
<tr>
<td>P3-005</td>
<td>labels[].label/level后端为int基本类型，SDK为Integer包装类型</td>
<td>后端vs SDK</td>
<td>labels[].label/level</td>
<td>Java类型差异，功能等价</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h2>五、统计信息</h2>
<br>
<h3>5.1 字段总量统计</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>统计维度</th>
<th>数量</th>
</tr></thead><tbody>
<tr>
<td>后端总字段数(含嵌套，约)</td>
<td>~115</td>
</tr>
<tr>
<td>Java SDK总字段数</td>
<td>~100</td>
</tr>
<tr>
<td>Go SDK总字段数</td>
<td>~98</td>
</tr>
<tr>
<td>三方完全一致字段数</td>
<td>~75</td>
</tr>
<tr>
<td>存在差异字段数</td>
<td>~40</td>
</tr>
</tbody></table>
<br>
<h3>5.2 问题数量统计</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>优先级</th>
<th>数量</th>
<th>说明</th>
</tr></thead><tbody>
<tr>
<td>P0 (严重缺失)</td>
<td>3</td>
<td>整个模块缺失</td>
</tr>
<tr>
<td>P1 (高优先级)</td>
<td>4</td>
<td>重要字段缺失或字段名不一致</td>
</tr>
<tr>
<td>P2 (中优先级)</td>
<td>6</td>
<td>一般字段缺失</td>
</tr>
<tr>
<td>P3 (低优先级)</td>
<td>5</td>
<td>类型差异/设计差异</td>
</tr>
<tr>
<td><strong>合计</strong></td>
<td><strong>18</strong></td>
<td>-</td>
</tr>
</tbody></table>
<br>
<h3>5.3 模块一致性概览</h3>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>模块</th>
<th>后端</th>
<th>Java SDK</th>
<th>Go SDK</th>
<th>一致状态</th>
</tr></thead><tbody>
<tr>
<td>antispam</td>
<td>存在</td>
<td>存在</td>
<td>存在</td>
<td>部分差异(P1-001,P1-002,P1-003,P2-001~003,P2-004,P2-005,P2-006)</td>
</tr>
<tr>
<td>emotionAnalysis</td>
<td>存在</td>
<td>存在</td>
<td>存在</td>
<td>完全一致</td>
</tr>
<tr>
<td>anticheat</td>
<td>存在</td>
<td>存在</td>
<td>存在</td>
<td>完全一致</td>
</tr>
<tr>
<td>userRisk</td>
<td>存在</td>
<td>存在</td>
<td>存在</td>
<td>实质一致(后端内部字段除外)</td>
</tr>
<tr>
<td>language</td>
<td>存在</td>
<td>存在</td>
<td>存在</td>
<td>完全一致</td>
</tr>
<tr>
<td>riskControl</td>
<td>存在</td>
<td>存在</td>
<td><strong>缺失</strong></td>
<td>P0-001</td>
</tr>
<tr>
<td>aigcPrompt</td>
<td>存在</td>
<td>存在</td>
<td>存在</td>
<td>部分差异(P1-004,P3-001)</td>
</tr>
<tr>
<td>llmCheckInfo</td>
<td>存在</td>
<td>存在</td>
<td>存在</td>
<td>完全一致</td>
</tr>
<tr>
<td>similarText</td>
<td>存在</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td>P0-002</td>
</tr>
<tr>
<td>underage</td>
<td>存在</td>
<td><strong>缺失</strong></td>
<td><strong>缺失</strong></td>
<td>P0-003</td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<h2>六、源文件信息</h2>
<br>
<table border="1" cellpadding="8" cellspacing="0">
<thead><tr>
<th>角色</th>
<th>文件路径</th>
</tr></thead><tbody>
<tr>
<td>Java后端</td>
<td><code>/Users/admin/Documents/coding/antispam-business-text/common/src/main/java/com/netease/is/antispam/text/common/bo/check/result/v5/TextCheckResultBodyV5.java</code></td>
</tr>
<tr>
<td>Java SDK</td>
<td><code>/Users/admin/Documents/coding/yidun-java-sdk/yidun-java-sdk/src/main/java/com/netease/yidun/sdk/antispam/text/v5/callback/response/TextCallBackResponse.java</code></td>
</tr>
<tr>
<td>Go SDK</td>
<td><code>/Users/admin/Documents/coding/yidun-golang-sdk/yidun/service/antispam/text/v5/callback/text_callback_response.go</code></td>
</tr>
</tbody></table>
<br>
<hr>
<br>
<p>*报告生成工具: Claude Code SDK字段对比任务*</p>
<br></body></html>