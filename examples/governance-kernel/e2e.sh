#!/usr/bin/env bash
# End-to-end smoke of the Governance Kernel API with curl.
# Usage: ./e2e.sh [BASE_URL] [TENANT] [ACTOR]
set -euo pipefail

BASE="${1:-http://localhost:8080}"
TENANT="${2:-supply-chain}"
ACTOR="${3:-approver-a}"

hdr=(-H "X-SYMBIVELA-Tenant: $TENANT" -H "X-SYMBIVELA-Actor: $ACTOR")

echo "== workspace =="
curl -sf -X POST "$BASE/v1/workspaces" "${hdr[@]}" -H "Idempotency-Key: e2e-ws" \
  -d '{"workspace_id":"e2e","name":"E2E","owner_id":"'"$ACTOR"'"}' >/dev/null

echo "== goal =="
curl -sf -X POST "$BASE/v1/goals" "${hdr[@]}" -H "Idempotency-Key: e2e-goal" \
  -d '{"goal_id":"g1","workspace_id":"e2e","title":"Restock parts","owner_id":"'"$ACTOR"'","status":"active"}' >/dev/null

echo "== decision =="
curl -sf -X POST "$BASE/v1/decisions" "${hdr[@]}" -H "Idempotency-Key: e2e-decision" \
  -d '{"decision_id":"d1","workspace_id":"e2e","request_ref":"approval://1","decision":"approve","context_digest":"sha256:ctx"}' >/dev/null

echo "== my-work =="
curl -sf "$BASE/v1/my-work" "${hdr[@]}"

echo ""
echo "== evidence =="
curl -sf "$BASE/v1/evidence/g1" "${hdr[@]}"

echo ""
echo "e2e ok"
