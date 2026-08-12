# Governance Kernel API Example

The trusted gateway supplies tenant and actor headers. A client does not place either value in the JSON body.

```bash
curl -X POST http://localhost:8080/v1/decisions \
  -H "Content-Type: application/json" \
  -H "X-SYMBIVELA-Tenant: supply-chain" \
  -H "X-SYMBIVELA-Actor: procurement-approver" \
  -H "Idempotency-Key: decision-001" \
  -d '{"decision_id":"decision-001","workspace_id":"inventory","request_ref":"approval://purchase/42","decision":"approve","context_digest":"sha256:decision-context"}'
```

This records a human decision only. The decision must still be submitted to AEGIVELA before an executable grant exists. For handoffs, the trusted `X-SYMBIVELA-Actor` becomes `from_owner_id`; clients provide only the receiving owner.
