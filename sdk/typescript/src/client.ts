export interface Attention {
  attention_id: string;
  workspace_id: string;
  kind: string;
  status: string;
  priority: number;
}

export interface MyWork {
  workspace_count: number;
  attention: Attention[];
  approvals: unknown[];
  handoffs: unknown[];
  work_projections: WorkProjection[];
}

export interface WorkProjectionOwner {
  kind: string;
  ref: string;
}

export interface WorkProjectionDependencySummary {
  count: number;
  work_item_ids: string[];
}

export interface WorkProjectionPlanSummary {
  version: number;
  objective: string;
}

export interface WorkProjectionFinding {
  id: string;
  kind: string;
  severity: string;
}

export interface WorkProjection {
  tenant_id?: string;
  work_item_id: string;
  goal_ref?: string;
  workspace_id: string;
  repository?: string;
  state: string;
  owner?: WorkProjectionOwner;
  dependency_summary?: WorkProjectionDependencySummary;
  coordination_plan_summary?: WorkProjectionPlanSummary;
  claim_count: number;
  open_finding_count: number;
  blocking_findings?: WorkProjectionFinding[];
  evidence_refs?: string[];
  version: number;
  updated_at: string;
  source_system: string;
  source_ref: string;
  received_at?: string;
}

/** Typed client for the SYMBIVELA Governance Kernel API. */
export class SymbivelaClient {
  constructor(
    private readonly base: string,
    private readonly tenant: string,
    private readonly actor: string,
  ) {}

  private async request(
    method: string,
    path: string,
    body?: unknown,
    idemKey?: string,
  ): Promise<Response> {
    const headers: Record<string, string> = {
      "X-SYMBIVELA-Tenant": this.tenant,
      "X-SYMBIVELA-Actor": this.actor,
    };
    if (body !== undefined) headers["Content-Type"] = "application/json";
    if (idemKey) headers["Idempotency-Key"] = idemKey;
    const response = await fetch(this.base + path, {
      method,
      headers,
      body: body === undefined ? undefined : JSON.stringify(body),
    });
    if (response.status >= 300) {
      throw new Error(`SYMBIVELA http ${response.status}: ${await response.text()}`);
    }
    return response;
  }

  createGoal(goalId: string, workspaceId: string, title: string, ownerId: string, idemKey: string): Promise<Response> {
    return this.request("POST", "/v1/goals", {
      goal_id: goalId,
      workspace_id: workspaceId,
      title,
      owner_id: ownerId,
      status: "active",
    }, idemKey);
  }

  recordDecision(decisionId: string, workspaceId: string, requestRef: string, decision: string, contextDigest: string, idemKey: string): Promise<Response> {
    return this.request("POST", "/v1/decisions", {
      decision_id: decisionId,
      workspace_id: workspaceId,
      request_ref: requestRef,
      decision,
      context_digest: contextDigest,
    }, idemKey);
  }

  async listAttention(workspaceId: string): Promise<Attention[]> {
    const response = await this.request("GET", `/v1/attention?workspace_id=${encodeURIComponent(workspaceId)}`);
    return (await response.json()) as Attention[];
  }

  async myWork(): Promise<MyWork> {
    const response = await this.request("GET", "/v1/my-work");
    return (await response.json()) as MyWork;
  }

  async getWorkProjection(workItemId: string): Promise<WorkProjection> {
    const response = await this.request("GET", `/v1/work-projections/${encodeURIComponent(workItemId)}`);
    return (await response.json()) as WorkProjection;
  }

  async listWorkProjections(workspaceId: string): Promise<WorkProjection[]> {
    const response = await this.request("GET", `/v1/work-projections?workspace_id=${encodeURIComponent(workspaceId)}`);
    return (await response.json()) as WorkProjection[];
  }
}
