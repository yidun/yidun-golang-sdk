package response

type LabelQueryInfo struct {
    Label     *int        `json:"label,omitempty"` // label
    Name      *string     `json:"name,omitempty"` // name
    SubLabels []*SubLabel `json:"subLabels,omitempty"` // subLabels
}

type SubLabel struct {
    BusinessId   *string       `json:"businessId,omitempty"`   // businessId
    BusinessType *[]int        `json:"businessType,omitempty"` // businessType
    CreateTime   *int64        `json:"createTime,omitempty"`   // createTime
    Code         *string       `json:"code,omitempty"`         // code
    Depth        *int          `json:"depth,omitempty"`        // depth
    Name         *string       `json:"name,omitempty"`         // name
    Type         *int          `json:"type,omitempty"`         // type
    SubLabels    []*SubLabel   `json:"subLabels,omitempty"`    // subLabels
}
