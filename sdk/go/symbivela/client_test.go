package symbivela

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClientDrivesCoreFlow(t *testing.T) {
	var calls []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls = append(calls, r.Method+" "+r.URL.Path)
		if r.Header.Get("X-SYMBIVELA-Tenant") != "tenant-a" || r.Header.Get("X-SYMBIVELA-Actor") != "approver-a" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		switch r.URL.Path {
		case "/v1/goals":
			w.WriteHeader(http.StatusCreated)
		case "/v1/decisions":
			w.WriteHeader(http.StatusCreated)
		case "/v1/attention":
			_, _ = w.Write([]byte(`[]`))
		case "/v1/my-work":
			_, _ = w.Write([]byte(`{"workspace_count":1,"attention":[],"approvals":[],"handoffs":[]}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := New(server.URL, "tenant-a", "approver-a")
	if err := client.CreateGoal(context.Background(), "g1", "w1", "Restock", "approver-a", "k1"); err != nil {
		t.Fatal(err)
	}
	if err := client.RecordDecision(context.Background(), "d1", "w1", "approval://1", "approve", "sha256:x", "k2"); err != nil {
		t.Fatal(err)
	}
	attention, err := client.ListAttention(context.Background(), "w1")
	if err != nil || len(attention) != 0 {
		t.Fatalf("attention = %+v, err = %v", attention, err)
	}
	work, err := client.MyWork(context.Background())
	if err != nil || work.WorkspaceCount != 1 {
		t.Fatalf("work = %+v, err = %v", work, err)
	}
	if len(calls) != 4 || !strings.Contains(strings.Join(calls, ","), "POST /v1/goals") {
		t.Fatalf("calls = %+v", calls)
	}
}
