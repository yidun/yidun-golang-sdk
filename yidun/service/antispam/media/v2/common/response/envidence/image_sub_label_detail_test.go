package envidence

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestLlmKeyWordLocationsJSON(t *testing.T) {
	const payload = `{"keyword":"建党100周年特别版","locations":[{"keyword":"建党100周年特别版","x1":0.1094,"y1":0.7299,"x2":0.5135,"y2":0.7839}]}`

	var result LlmKeyWord
	if err := json.Unmarshal([]byte(payload), &result); err != nil {
		t.Fatalf("unmarshal LLM locations: %v", err)
	}
	if result.Keyword == nil || *result.Keyword != "建党100周年特别版" {
		t.Fatalf("unexpected keyword: %v", result.Keyword)
	}
	if len(result.Locations) != 1 || result.Locations[0] == nil {
		t.Fatalf("unexpected locations: %#v", result.Locations)
	}
	location := result.Locations[0]
	if location.Keyword == nil || *location.Keyword != "建党100周年特别版" ||
		location.X1 == nil || location.Y1 == nil || location.X2 == nil || location.Y2 == nil {
		t.Fatalf("unexpected location: %#v", location)
	}

	keywordOnly, err := json.Marshal(LlmKeyWord{Keyword: result.Keyword})
	if err != nil {
		t.Fatalf("marshal keyword-only LLM result: %v", err)
	}
	if strings.Contains(string(keywordOnly), "locations") {
		t.Fatalf("nil locations should be omitted: %s", keywordOnly)
	}
}

func TestMediaImageEvidenceUnmarshalLlmExtension(t *testing.T) {
	const extension = ` {"riskCode":"A01","score":0.98} `
	const evidenceJSON = `{"images":[{"status":2,"labels":[{"label":100,"subLabels":[{"subLabel":"100001","details":{"llm":{"extension":" {\"riskCode\":\"A01\",\"score\":0.98} "}}}]}]}]}`

	var evidence MediaAntispamEvidence
	if err := json.Unmarshal([]byte(evidenceJSON), &evidence); err != nil {
		t.Fatalf("unmarshal media image evidence: %v", err)
	}
	if len(evidence.Images) != 1 || len(evidence.Images[0].Labels) != 1 {
		t.Fatalf("unexpected image evidence: %#v", evidence.Images)
	}
	subLabels := evidence.Images[0].Labels[0].SubLabels
	if len(subLabels) != 1 || subLabels[0].Details == nil || subLabels[0].Details.Llm == nil {
		t.Fatalf("unexpected image subLabels: %#v", subLabels)
	}
	llm := subLabels[0].Details.Llm
	if llm.Extension == nil {
		t.Fatal("llm extension should not be nil")
	}
	if got := *llm.Extension; got != extension {
		t.Fatalf("extension mismatch: got %q, want %q", got, extension)
	}
	if llm.Keyword != nil {
		t.Fatalf("extension-only response should keep keyword nil, got %q", *llm.Keyword)
	}
}
