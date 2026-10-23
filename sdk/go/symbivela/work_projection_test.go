package symbivela

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestWorkProjectionReadPaths(t *testing.T) {
	var got []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = append(got, r.URL.RequestURI())
		if r.URL.Path == "/v1/work-projections" {
			_ = json.NewEncoder(w).Encode([]WorkProjection{{WorkItemID: "wi-1", WorkspaceID: "ws-1", State: "executing", Version: 2}})
			return
		}
		_ = json.NewEncoder(w).Encode(WorkProjection{WorkItemID: "wi-1", WorkspaceID: "ws-1", State: "executing", Version: 2})
	}))
	defer server.Close()

	client := New(server.URL, "t1", "a1")
	if _, err := client.GetWorkProjection(context.Background(), "wi-1"); err != nil {
		t.Fatal(err)
	}
	if _, err := client.ListWorkProjections(context.Background(), "ws-1"); err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0] != "/v1/work-projections/wi-1" || got[1] != "/v1/work-projections?workspace_id=ws-1" {
		t.Fatalf("paths=%v", got)
	}
}

func TestWorkProjectionContractSchema(t *testing.T) {
	var schema struct {
		Required   []string       `json:"required"`
		Properties map[string]any `json:"properties"`
	}
	data, err := os.ReadFile("../../../contracts/work-projection-v1.schema.json")
	if err != nil {
		t.Skipf("contract not readable: %v", err)
	}
	if err := json.Unmarshal(data, &schema); err != nil {
		t.Fatal(err)
	}
	if len(schema.Required) == 0 {
		t.Fatal("schema has no required fields")
	}
	if _, ok := schema.Properties["work_item_id"]; !ok {
		t.Fatal("schema missing work_item_id")
	}
}
