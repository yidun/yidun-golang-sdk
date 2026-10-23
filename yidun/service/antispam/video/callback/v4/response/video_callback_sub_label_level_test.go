package response

import (
	"encoding/json"
	"testing"
)

func TestVideoCallbackSubLabelLevel(t *testing.T) {
	const payload = `{"result":[{"antispam":{"pictures":[{"labels":[{"label":200,"level":2,` +
		`"subLabels":[{"subLabel":200001,"level":1}]}]}]}}]}`

	var response VideoCallbackV4Response
	if err := json.Unmarshal([]byte(payload), &response); err != nil {
		t.Fatalf("unmarshal video callback response: %v", err)
	}
	subLabel := (*response.Result)[0].Antispam.Pictures[0].Labels[0].SubLabels[0]
	if subLabel.Level == nil || *subLabel.Level != 1 {
		t.Fatalf("unexpected video subLabel level: %#v", subLabel.Level)
	}
}

func TestVideoCallbackMissingSubLabelLevel(t *testing.T) {
	const payload = `{"result":[{"antispam":{"pictures":[{"labels":[{"subLabels":[{` +
		`"subLabel":200001}]}]}]}}]}`

	var response VideoCallbackV4Response
	if err := json.Unmarshal([]byte(payload), &response); err != nil {
		t.Fatalf("unmarshal video callback response: %v", err)
	}
	subLabel := (*response.Result)[0].Antispam.Pictures[0].Labels[0].SubLabels[0]
	if subLabel.Level != nil {
		t.Fatalf("missing level should remain nil, got %d", *subLabel.Level)
	}
}
