package envidence

import (
	"encoding/json"
	"testing"
)

func TestMediaLlmEvidenceUnmarshalProductExtension(t *testing.T) {
	const extension = `{"category":"饰品/玄学","QA":"1"}`
	const evidenceJSON = `{"llms":[{"subLabel":"400017","keyword":"","extension":"{\"category\":\"饰品/玄学\",\"QA\":\"1\"}"}]}`

	var evidence MediaAntispamEvidence
	if err := json.Unmarshal([]byte(evidenceJSON), &evidence); err != nil {
		t.Fatalf("unmarshal product llm evidence: %v", err)
	}

	if len(evidence.Llms) != 1 || evidence.Llms[0].Extension == nil {
		t.Fatalf("unexpected product llm evidence: %#v", evidence.Llms)
	}
	if got := *evidence.Llms[0].Extension; got != extension {
		t.Fatalf("extension mismatch: got %q, want %q", got, extension)
	}
}
