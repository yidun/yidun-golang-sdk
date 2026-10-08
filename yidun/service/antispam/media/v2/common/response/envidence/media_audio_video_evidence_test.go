package envidence

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestMediaAudioVideoEvidenceUnmarshalPicCount(t *testing.T) {
	tests := []struct {
		name     string
		payload  string
		expected *int64
	}{
		{
			name:     "positive pic count",
			payload:  `{"picCount":37}`,
			expected: int64Pointer(37),
		},
		{
			name:     "zero pic count",
			payload:  `{"picCount":0}`,
			expected: int64Pointer(0),
		},
		{
			name:     "missing pic count",
			payload:  `{}`,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var evidence MediaAudioVideoEvidence
			if err := json.Unmarshal([]byte(test.payload), &evidence); err != nil {
				t.Fatalf("unmarshal evidence: %v", err)
			}

			if test.expected == nil {
				if evidence.PicCount != nil {
					t.Fatalf("expected nil picCount, got %d", *evidence.PicCount)
				}
				return
			}
			if evidence.PicCount == nil {
				t.Fatal("expected picCount, got nil")
			}
			if *evidence.PicCount != *test.expected {
				t.Fatalf("expected picCount %d, got %d", *test.expected, *evidence.PicCount)
			}
		})
	}
}

func TestMediaAudioVideoEvidenceMarshalOmitsNilPicCount(t *testing.T) {
	payload, err := json.Marshal(MediaAudioVideoEvidence{})
	if err != nil {
		t.Fatalf("marshal evidence: %v", err)
	}

	if strings.Contains(string(payload), `"picCount"`) {
		t.Fatalf("expected nil picCount to be omitted, got %s", payload)
	}
}

func int64Pointer(value int64) *int64 {
	return &value
}
