package callback

import (
	"encoding/json"
	"testing"
)

func TestMediaCallbackSubLabelLevel(t *testing.T) {
	const payload = `{"result":[{"antispam":{"evidences":{` +
		`"texts":[{"labels":[{"label":500,"subLabels":[{"subLabel":"500087","level":2}]}]}],` +
		`"images":[{"labels":[{"label":500,"subLabels":[{"subLabel":"50000","level":1}]}]}]` +
		`}}}]}`

	var response MediaCallbackResponseV2
	if err := json.Unmarshal([]byte(payload), &response); err != nil {
		t.Fatalf("unmarshal media callback response: %v", err)
	}
	if len(response.Result) != 1 || response.Result[0].Antispam == nil ||
		response.Result[0].Antispam.Evidences == nil {
		t.Fatalf("unexpected callback result: %#v", response.Result)
	}

	evidences := response.Result[0].Antispam.Evidences
	textSubLabels := evidences.Texts[0].Labels[0].SubLabels
	imageSubLabels := evidences.Images[0].Labels[0].SubLabels
	if textSubLabels[0].Level == nil || *textSubLabels[0].Level != 2 {
		t.Fatalf("unexpected text subLabel level: %#v", textSubLabels[0].Level)
	}
	if imageSubLabels[0].Level == nil || *imageSubLabels[0].Level != 1 {
		t.Fatalf("unexpected image subLabel level: %#v", imageSubLabels[0].Level)
	}
}

func TestMediaCallbackMissingSubLabelLevel(t *testing.T) {
	const payload = `{"result":[{"antispam":{"evidences":{` +
		`"texts":[{"labels":[{"label":500,"subLabels":[{"subLabel":"500087"}]}]}]` +
		`}}}]}`

	var response MediaCallbackResponseV2
	if err := json.Unmarshal([]byte(payload), &response); err != nil {
		t.Fatalf("unmarshal media callback response: %v", err)
	}
	subLabel := response.Result[0].Antispam.Evidences.Texts[0].Labels[0].SubLabels[0]
	if subLabel.Level != nil {
		t.Fatalf("missing level should remain nil, got %d", *subLabel.Level)
	}
}
