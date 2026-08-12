# SYMBIVELA

**Human–Agent Collaboration Workspace · Human Sovereignty Control Surface**

[简体中文](README.zh-CN.md)

SYMBIVELA makes human sovereignty operational in machine-majority enterprises. It is a persistent workspace where humans set goals, review machine-generated plans, approve high-risk actions, resolve exceptions, intervene in autonomous execution, accept outcomes, and inspect evidence across agents, workflows, robots, and digital twins.

> **One-liner:** SYMBIVELA lets people set goals, review machine-generated plans, approve high-risk actions, resolve exceptions, intervene in execution, accept outcomes and inspect evidence across agents, workflows, robots and twins.

---

## Why SYMBIVELA

In a machine-majority enterprise the scarce resource is no longer compute — it is **human attention, judgment, and final accountability**. When agents, workflows, robots, and twins outnumber people, a chat window is not a control surface. SYMBIVELA compresses machine-scale operations into human-scale **Goals, Decisions, Exceptions, Interventions, Outcomes, and Evidence**.

| Old assumption | Autonomous reality | SYMBIVELA |
| --- | --- | --- |
| Humans run everything | Machines far outnumber humans | Attention compression and exception-first operation |
| A conversation is the unit | Goals and executions persist across systems and days | Persistent workspace and cross-runtime timeline |
| Human-in-the-loop = approval buttons | Also goal setting, plan amendment, intervention, takeover, value conflicts | A complete Human Governance Loop |
| More logs means more transparency | Event floods hide accountability | Progressive disclosure and evidence bundles |

## What SYMBIVELA Solves

- **Understand** — compress goals, plans, risk, and evidence into decision-ready views.
- **Decide** — route the right decision to the accountable person at the right time.
- **Control** — authorize, constrain, intervene, take over, or stop activity inside valid authority.
- **Account** — preserve a reconstructable responsibility chain from goal through outcome.

## Key Capabilities

- **Goal workspace** — versioned goals with success criteria, constraints, owner, and supervision mode.
- **Plan review** — read-only ORCHADYN plan projections with impact and revision context.
- **Approval center** — AEGIVELA approval gates presented with risk, scope, alternatives, and evidence.
- **Attention & exceptions** — deduplicated, prioritized attention items; claim / resolve / escalate lifecycle.
- **Intervention control** — request, AEGIVELA re-authorization, and runtime execution for pause, cancel, takeover, and emergency stop.
- **Handoffs** — explicit responsibility transfer with acceptance, rejection, and escalation.
- **Outcome & evidence** — outcome acceptance and a composed responsibility chain (goal → plan → decisions → outcome).

## Repository Model

| Repository | License | Role |
| --- | --- | --- |
| `symbivela-open` | Apache-2.0 | Public contracts, API schemas, examples, SDKs, and approved core binaries — **this repository** |
| `symbivela` | AGPL-3.0 | Open-core implementation of the Governance Kernel |
| `symbivela-ee` | Enterprise | Enterprise features and internal product governance documents |

SYMBIVELA is a **control surface, not a control authority**. It owns the workspace experience and human interaction records. Plans, grants, execution, world state, and source evidence remain authoritative in their owning products (ORCHADYN, AEGIVELA, PRAXOVELA, RHEOVELA, KINETOVELA, ONTOVELA, TEKMOVELA, Harmovela).

## This Repository: SYMBIVELA Open

Public, Apache-2.0 surface for integrating with SYMBIVELA.

### Contracts

- `contracts/openapi.yaml` — Governance Kernel HTTP API v1 (trusted-context headers, idempotency, workspace roles).
- `contracts/governance-kernel-v1.schema.json` — core human-governance objects (Workspace, AttentionItem, HumanDecision, Handoff, Goal).
- `contracts/orchadyn-plan-projection-v1.schema.json` — read-only plan projection contract.
- `contracts/aegivela-approval-request-v1.schema.json` — approval gate projection contract.
- `contracts/harmovela-event-envelope-v1.schema.json` — runtime event envelope contract.
- `contracts/intervention-request-v1.schema.json` — intervention request contract.

### Examples

- `examples/governance-kernel/` — trusted-context API usage.

## Quick Start

```bash
# Run the open-core API and PostgreSQL locally (from the core repository)
docker compose up --build

# Apply migrations explicitly before deploying
DATABASE_URL="postgres://..." go run ./backend/cmd/symbivela-migrate
```

## Documentation

- Product positioning, architecture, roadmap: `symbivela-ee/docs/`
- Public contracts and examples: this repository.

## License

Apache-2.0. The core implementation in `symbivela` is AGPL-3.0; enterprise extensions in `symbivela-ee` are under the Enterprise License.
