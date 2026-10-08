package sync

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

func TestImageCheckResponseUnmarshalLlmExtension(t *testing.T) {
	const extension = ` {"riskCode":"A01","score":0.98} `
	const responseJSON = `{"result":[{"antispam":{"labels":[{"subLabels":[{"details":{"llm":{"extension":" {\"riskCode\":\"A01\",\"score\":0.98} "}}}] }]}}]}`

	var response ImageCheckResponse
	if err := json.Unmarshal([]byte(responseJSON), &response); err != nil {
		t.Fatalf("unmarshal image check response: %v", err)
	}
	if response.Result == nil || len(*response.Result) != 1 {
		t.Fatalf("unexpected result: %#v", response.Result)
	}
	antispam := (*response.Result)[0].Antispam
	if antispam == nil || antispam.Labels == nil || len(*antispam.Labels) != 1 {
		t.Fatalf("unexpected antispam labels: %#v", antispam)
	}
	subLabels := (*antispam.Labels)[0].SubLabels
	if subLabels == nil || len(*subLabels) != 1 {
		t.Fatalf("unexpected subLabels: %#v", subLabels)
	}
	details := (*subLabels)[0].Details
	if details == nil || details.Llm == nil || details.Llm.Extension == nil {
		t.Fatalf("unexpected llm details: %#v", details)
	}
	if got := *details.Llm.Extension; got != extension {
		t.Fatalf("extension mismatch: got %q, want %q", got, extension)
	}
	if details.Llm.Keyword != nil {
		t.Fatalf("extension-only response should keep keyword nil, got %q", *details.Llm.Keyword)
	}
}
