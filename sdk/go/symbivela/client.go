// Package symbivela is a typed client for the SYMBIVELA Governance Kernel API.
package symbivela

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Client talks to a SYMBIVELA core with trusted-context tenant/actor identity.
type Client struct {
	base   string
	tenant string
	actor  string
	hc     *http.Client
}

func New(base, tenant, actor string) *Client {
	return &Client{base: base, tenant: tenant, actor: actor, hc: &http.Client{Timeout: 20 * time.Second}}
}

type Attention struct {
	AttentionID string `json:"attention_id"`
	WorkspaceID string `json:"workspace_id"`
	Kind        string `json:"kind"`
	Status      string `json:"status"`
	Priority    int    `json:"priority"`
}

type Evidence struct {
	SubjectRef    string            `json:"subject_ref"`
	Decisions     []json.RawMessage `json:"decisions"`
	Interventions []json.RawMessage `json:"interventions"`
	Artifacts     []json.RawMessage `json:"artifacts"`
}

type MyWork struct {
	WorkspaceCount  int               `json:"workspace_count"`
	Attention       []Attention       `json:"attention"`
	Approvals       []json.RawMessage `json:"approvals"`
	Handoffs        []json.RawMessage `json:"handoffs"`
	WorkProjections []WorkProjection  `json:"work_projections"`
}

type WorkProjectionOwner struct {
	Kind string `json:"kind"`
	Ref  string `json:"ref"`
}

type WorkProjectionDependencySummary struct {
	Count       int      `json:"count"`
	WorkItemIDs []string `json:"work_item_ids"`
}

type WorkProjectionPlanSummary struct {
	Version   int64  `json:"version"`
	Objective string `json:"objective"`
}

type WorkProjectionFinding struct {
	ID       string `json:"id"`
	Kind     string `json:"kind"`
	Severity string `json:"severity"`
}

type WorkProjection struct {
	TenantID                string                           `json:"tenant_id,omitempty"`
	WorkItemID              string                           `json:"work_item_id"`
	GoalRef                 string                           `json:"goal_ref,omitempty"`
	WorkspaceID             string                           `json:"workspace_id"`
	Repository              string                           `json:"repository,omitempty"`
	State                   string                           `json:"state"`
	Owner                   *WorkProjectionOwner             `json:"owner,omitempty"`
	DependencySummary       *WorkProjectionDependencySummary `json:"dependency_summary,omitempty"`
	CoordinationPlanSummary *WorkProjectionPlanSummary       `json:"coordination_plan_summary,omitempty"`
	ClaimCount              int                              `json:"claim_count"`
	OpenFindingCount        int                              `json:"open_finding_count"`
	BlockingFindings        []WorkProjectionFinding          `json:"blocking_findings,omitempty"`
	EvidenceRefs            []string                         `json:"evidence_refs,omitempty"`
	Version                 int64                            `json:"version"`
	UpdatedAt               string                           `json:"updated_at"`
	SourceSystem            string                           `json:"source_system"`
	SourceRef               string                           `json:"source_ref"`
	ReceivedAt              string                           `json:"received_at,omitempty"`
}

// Do issues a raw request with trusted-context headers and decodes the JSON body.
func (c *Client) Do(ctx context.Context, method, path string, body any, idemKey string, out any) error {
	var reader io.Reader
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(ctx, method, c.base+path, reader)
	if err != nil {
		return err
	}
	req.Header.Set("X-SYMBIVELA-Tenant", c.tenant)
	req.Header.Set("X-SYMBIVELA-Actor", c.actor)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if idemKey != "" {
		req.Header.Set("Idempotency-Key", idemKey)
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("http %d: %s", resp.StatusCode, string(payload))
	}
	if out != nil {
		return json.Unmarshal(payload, out)
	}
	return nil
}

func (c *Client) CreateGoal(ctx context.Context, goalID, workspaceID, title, ownerID string, idemKey string) error {
	return c.Do(ctx, http.MethodPost, "/v1/goals", map[string]any{
		"goal_id": goalID, "workspace_id": workspaceID, "title": title, "owner_id": ownerID, "status": "active",
	}, idemKey, nil)
}

func (c *Client) RecordDecision(ctx context.Context, decisionID, workspaceID, requestRef, decision, contextDigest, idemKey string) error {
	return c.Do(ctx, http.MethodPost, "/v1/decisions", map[string]any{
		"decision_id": decisionID, "workspace_id": workspaceID, "request_ref": requestRef,
		"decision": decision, "context_digest": contextDigest,
	}, idemKey, nil)
}

func (c *Client) ListAttention(ctx context.Context, workspaceID string) ([]Attention, error) {
	var items []Attention
	err := c.Do(ctx, http.MethodGet, "/v1/attention?workspace_id="+url.QueryEscape(workspaceID), nil, "", &items)
	return items, err
}

func (c *Client) ComposeEvidence(ctx context.Context, goalID string) (Evidence, error) {
	var evidence Evidence
	err := c.Do(ctx, http.MethodGet, "/v1/evidence/"+url.PathEscape(goalID), nil, "", &evidence)
	return evidence, err
}

func (c *Client) MyWork(ctx context.Context) (MyWork, error) {
	var work MyWork
	err := c.Do(ctx, http.MethodGet, "/v1/my-work", nil, "", &work)
	return work, err
}

func (c *Client) GetWorkProjection(ctx context.Context, workItemID string) (WorkProjection, error) {
	var projection WorkProjection
	err := c.Do(ctx, http.MethodGet, "/v1/work-projections/"+url.PathEscape(workItemID), nil, "", &projection)
	return projection, err
}

func (c *Client) ListWorkProjections(ctx context.Context, workspaceID string) ([]WorkProjection, error) {
	var projections []WorkProjection
	err := c.Do(ctx, http.MethodGet, "/v1/work-projections?workspace_id="+url.QueryEscape(workspaceID), nil, "", &projections)
	return projections, err
}
