# Work Projection Example

COORDIVELA publishes one latest work-projection snapshot per work item. SYMBIVELA stores the snapshot and its source references; COORDIVELA remains authoritative for coordination state.

```json
{
  "work_item_id": "wi-1",
  "goal_ref": "symbivela:goal:goal-1",
  "workspace_id": "ws-1",
  "repository": "repo-a",
  "state": "executing",
  "owner": {"kind": "agent", "ref": "agent:alice"},
  "coordination_plan_summary": {"version": 2, "objective": "Ship the coordination bridge"},
  "claim_count": 2,
  "open_finding_count": 1,
  "blocking_findings": [{"id": "f-1", "kind": "path_conflict", "severity": "blocking"}],
  "evidence_refs": ["tekmovela:evidence:abc"],
  "version": 7,
  "updated_at": "2026-10-23T05:59:58Z",
  "source_system": "coordivela",
  "source_ref": "evt_proj_abc"
}
```

Publish the snapshot with a COORDIVELA service token:

```bash
curl -X POST http://localhost:8080/v1/work-projections \
  -H "Content-Type: application/json" \
  -H "X-SYMBIVELA-Tenant: supply-chain" \
  -H "X-SYMBIVELA-Service: coordivela" \
  -H "X-SYMBIVELA-Token: $SERVICE_TOKEN" \
  -d @payload.json
```

Read the latest snapshot for a work item:

```bash
curl http://localhost:8080/v1/work-projections/wi-1 \
  -H "X-SYMBIVELA-Tenant: supply-chain" \
  -H "X-SYMBIVELA-Actor: approver-a"
```

A version that is stale or equal to the stored version is ignored. The `source_ref` makes replay of the same projection source idempotent.
