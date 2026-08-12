# SYMBIVELA SDKs

Apache-2.0 typed clients for the Governance Kernel API. Both send trusted-context
`X-SYMBIVELA-Tenant` / `X-SYMBIVELA-Actor` headers and an `Idempotency-Key` on
mutating calls, matching the OpenAPI contract in `contracts/openapi.yaml`.

## Go

```bash
cd sdk/go
go test ./...
```

```go
import "github.com/axisrobo/symbivela-open/sdk/go/symbivela"

client := symbivela.New("http://localhost:8080", "supply-chain", "approver-a")
err := client.CreateGoal(ctx, "g1", "ws", "Restock parts", "approver-a", "goal-1")
err = client.RecordDecision(ctx, "d1", "ws", "approval://1", "approve", "sha256:ctx", "decision-1")
attention, _ := client.ListAttention(ctx, "ws")
work, _ := client.MyWork(ctx)
```

## TypeScript

```bash
cd sdk/typescript
npm install
npm run build
```

```ts
import { SymbivelaClient } from "./src/client";

const client = new SymbivelaClient("http://localhost:8080", "supply-chain", "approver-a");
await client.createGoal("g1", "ws", "Restock parts", "approver-a", "goal-1");
await client.recordDecision("d1", "ws", "approval://1", "approve", "sha256:ctx", "decision-1");
const attention = await client.listAttention("ws");
const work = await client.myWork();
```

The SDK surface intentionally mirrors the public contract; see `contracts/openapi.yaml`
for the complete API.
