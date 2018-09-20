package api_test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/cdempsie/payments-example/api"
	api_test "github.com/cdempsie/payments-example/test"
)

// TestFromJSON tests that the sample response can be decoded to the ListHolder struct without error.
func TestFromJSON(t *testing.T) {
	dec := json.NewDecoder(strings.NewReader(api_test.ListResponse))
	err := dec.Decode(&api.ListHolder{})
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}
}

// TestToJSONAndBack tests the roundtrip from JSON to struct and back.
func TestToJSONAndBack(t *testing.T) {
	dec := json.NewDecoder(strings.NewReader(api_test.ListResponse))
	listHolder := &api.ListHolder{}
	err := dec.Decode(listHolder)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(listHolder)
	if err != nil {
		t.Fatalf("Failed to encode JSON: %v", err)
	}
}
