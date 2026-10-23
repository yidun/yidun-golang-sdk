package response

import (
	"encoding/json"
	"testing"
)

func TestAudioCallbackSubLabelLevel(t *testing.T) {
	const payload = `{"result":[{"antispam":{"segments":[{"labels":[{"label":200,"level":2,` +
		`"subLabels":[{"subLabel":"200001","level":2}]}]}]}}]}`

	var response AudioCallBackV4Response
	if err := json.Unmarshal([]byte(payload), &response); err != nil {
		t.Fatalf("unmarshal audio callback response: %v", err)
	}
	subLabel := response.Result[0].Antispam.Segments[0].Labels[0].SubLabels[0]
	if subLabel.Level == nil || *subLabel.Level != 2 {
		t.Fatalf("unexpected audio subLabel level: %#v", subLabel.Level)
	}
}

func TestAudioCallbackMissingSubLabelLevel(t *testing.T) {
	const payload = `{"result":[{"antispam":{"segments":[{"labels":[{"subLabels":[{"subLabel":"200001"}]}]}]}}]}`

	var response AudioCallBackV4Response
	if err := json.Unmarshal([]byte(payload), &response); err != nil {
		t.Fatalf("unmarshal audio callback response: %v", err)
	}
	subLabel := response.Result[0].Antispam.Segments[0].Labels[0].SubLabels[0]
	if subLabel.Level != nil {
		t.Fatalf("missing level should remain nil, got %d", *subLabel.Level)
	}
}
