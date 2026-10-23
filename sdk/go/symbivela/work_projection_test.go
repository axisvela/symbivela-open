package symbivela

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const workProjectionFixture = `{
  "work_item_id": "wi-1",
  "goal_ref": "symbivela:goal:goal-1",
  "workspace_id": "ws-1",
  "repository": "repo-a",
  "state": "executing",
  "owner": {"kind": "agent", "ref": "agent:alice"},
  "dependency_summary": {"count": 2, "work_item_ids": ["wi-0", "wi-2"]},
  "coordination_plan_summary": {"version": 2, "objective": "Ship the coordination bridge"},
  "claim_count": 2,
  "open_finding_count": 1,
  "blocking_findings": [{"id": "f-1", "kind": "path_conflict", "severity": "blocking"}],
  "evidence_refs": ["tekmovela:evidence:abc"],
  "version": 7,
  "updated_at": "2026-10-23T05:59:58Z",
  "source_system": "coordivela",
  "source_ref": "evt_proj_abc"
}`

func TestWorkProjectionReadPaths(t *testing.T) {
	var got []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = append(got, r.URL.RequestURI())
		if r.URL.Path == "/v1/work-projections" {
			_, _ = w.Write([]byte("[" + workProjectionFixture + "]"))
			return
		}
		_, _ = w.Write([]byte(workProjectionFixture))
	}))
	defer server.Close()

	assertProjection := func(t *testing.T, projection WorkProjection) {
		t.Helper()
		if projection.WorkItemID != "wi-1" || projection.WorkspaceID != "ws-1" || projection.State != "executing" || projection.Version != 7 {
			t.Fatalf("projection=%+v", projection)
		}
		if projection.Owner == nil || projection.Owner.Kind != "agent" || projection.Owner.Ref != "agent:alice" {
			t.Fatalf("owner=%+v", projection.Owner)
		}
		if projection.DependencySummary == nil || projection.DependencySummary.Count != 2 || len(projection.DependencySummary.WorkItemIDs) != 2 || projection.DependencySummary.WorkItemIDs[0] != "wi-0" {
			t.Fatalf("dependency_summary=%+v", projection.DependencySummary)
		}
		if projection.CoordinationPlanSummary == nil || projection.CoordinationPlanSummary.Version != 2 || projection.CoordinationPlanSummary.Objective != "Ship the coordination bridge" {
			t.Fatalf("coordination_plan_summary=%+v", projection.CoordinationPlanSummary)
		}
		if len(projection.BlockingFindings) != 1 || projection.BlockingFindings[0].ID != "f-1" || projection.BlockingFindings[0].Kind != "path_conflict" || projection.BlockingFindings[0].Severity != "blocking" {
			t.Fatalf("blocking_findings=%+v", projection.BlockingFindings)
		}
		if len(projection.EvidenceRefs) != 1 || projection.EvidenceRefs[0] != "tekmovela:evidence:abc" {
			t.Fatalf("evidence_refs=%v", projection.EvidenceRefs)
		}
	}

	client := New(server.URL, "t1", "a1")
	projection, err := client.GetWorkProjection(context.Background(), "wi-1")
	if err != nil {
		t.Fatal(err)
	}
	assertProjection(t, projection)
	projections, err := client.ListWorkProjections(context.Background(), "ws-1")
	if err != nil {
		t.Fatal(err)
	}
	if len(projections) != 1 {
		t.Fatalf("projections=%v", projections)
	}
	assertProjection(t, projections[0])
	if len(got) != 2 || got[0] != "/v1/work-projections/wi-1" || got[1] != "/v1/work-projections?workspace_id=ws-1" {
		t.Fatalf("paths=%v", got)
	}
}

func TestWorkProjectionContractSchemas(t *testing.T) {
	cases := []struct {
		name      string
		path      string
		required  []string
		forbidden []string
	}{
		{
			name:     "response",
			path:     "../../../contracts/work-projection-v1.schema.json",
			required: []string{"work_item_id", "workspace_id", "state", "version", "updated_at", "source_system", "source_ref", "claim_count", "open_finding_count"},
		},
		{
			name:      "request",
			path:      "../../../contracts/work-projection-request-v1.schema.json",
			required:  []string{"work_item_id", "workspace_id", "state", "version", "updated_at", "source_system", "source_ref"},
			forbidden: []string{"tenant_id", "received_at"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var schema struct {
				Required   []string       `json:"required"`
				Properties map[string]any `json:"properties"`
			}
			data, err := os.ReadFile(tc.path)
			if err != nil {
				t.Fatalf("contract not readable: %v", err)
			}
			if err := json.Unmarshal(data, &schema); err != nil {
				t.Fatal(err)
			}
			required := map[string]bool{}
			for _, field := range schema.Required {
				required[field] = true
			}
			for _, field := range tc.required {
				if _, ok := schema.Properties[field]; !ok {
					t.Fatalf("schema missing property %s", field)
				}
				if !required[field] {
					t.Fatalf("schema does not require %s", field)
				}
			}
			for _, field := range tc.forbidden {
				if _, ok := schema.Properties[field]; ok {
					t.Fatalf("schema must not include property %s", field)
				}
			}
		})
	}
}
