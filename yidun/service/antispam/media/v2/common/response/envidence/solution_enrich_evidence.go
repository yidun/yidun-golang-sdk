package envidence

/**
 * 解决方案维度-失败集合
 */
type SolutionEnrichEvidence struct {
	FailedUnits []*FailedUnit `json:"failedUnits,omitempty"`
}

func (s *SolutionEnrichEvidence) GetFailedUnits() []*FailedUnit {
	return s.FailedUnits
}

func (s *SolutionEnrichEvidence) SetFailedUnits(failedUnits []*FailedUnit) {
	s.FailedUnits = failedUnits
}

// 失败原因
type FailedUnit struct {
	DataId        *string `json:"dataId,omitempty"`
	FailureReason *int    `json:"failureReason,omitempty"`
}

func (f *FailedUnit) GetDataId() *string {
	return f.DataId
}

func (f *FailedUnit) SetDataId(dataId *string) {
	f.DataId = dataId
}

func (f *FailedUnit) GetFailureReason() *int {
	return f.FailureReason
}

func (f *FailedUnit) SetFailureReason(failureReason *int) {
	f.FailureReason = failureReason
}
