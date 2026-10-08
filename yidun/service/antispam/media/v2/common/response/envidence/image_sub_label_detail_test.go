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
