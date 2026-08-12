# SYMBIVELA

**人类-智能体协作工作空间 · 人类主权控制面**

[English](README.md)

SYMBIVELA 让"人类主权"在机器为主体的企业中真正落地。它是一个持久化工作空间：人类在这里设定目标、审阅机器生成的计划、审批高风险动作、处理异常、干预自主执行、验收结果，并跨智能体、工作流、机器人与数字孪生核查证据。

> **一句话定位：** SYMBIVELA 让人类能够跨智能体、工作流、机器人与数字孪生设定目标、审阅机器生成的计划、审批高风险动作、解决异常、干预执行、验收成果并检查证据。

---

## 为什么需要 SYMBIVELA

在机器为主体的企业中，稀缺资源不再是算力，而是**人类的注意力、判断力与最终责任**。当智能体、工作流、机器人与孪生的数量超过人类时，聊天窗口不是控制面。SYMBIVELA 把机器规模运行压缩为人类规模的**目标、决策、异常、干预、成果与证据**。

| 旧假设 | 自治系统现实 | SYMBIVELA |
| --- | --- | --- |
| 人类是主要运行者 | 机器远多于人类 | 注意力压缩与异常优先操作 |
| 会话是基本单位 | 目标与执行跨系统、跨天持续存在 | 持久化工作空间与跨运行时时间线 |
| 人在环上 = 审批按钮 | 还包括目标设定、计划修订、干预、接管、价值冲突 | 完整的人类治理闭环 |
| 日志越多越透明 | 事件洪流掩盖责任 | 渐进式披露与证据聚合 |

## SYMBIVELA 解决什么问题

- **理解（Understand）** —— 将目标、计划、风险与证据压缩为可决策视图。
- **决策（Decide）** —— 在正确时间把正确决定交给应负责的人。
- **控制（Control）** —— 在授权边界内审批、约束、干预、接管或停止活动。
- **追责（Account）** —— 保留从目标到成果的可重建责任链。

## 关键能力

- **目标工作空间** —— 带成功标准、约束、负责人与监督模式的版本化目标。
- **计划审阅** —— 来自 ORCHADYN 的只读计划投影，含影响与修订上下文。
- **审批中心** —— 呈现 AEGIVELA 审批门禁的风险、范围、替代方案与证据。
- **注意力与异常** —— 去重、优先级排序的注意力项；认领/解决/升级生命周期。
- **干预控制** —— 请求、AEGIVELA 重新授权与运行时执行，支持暂停、取消、接管、急停。
- **交接** —— 明确的责任转移，含接受、拒绝与升级。
- **成果与证据** —— 成果验收与责任链组合（目标 → 计划 → 决策 → 成果）。

## 仓库模型

| 仓库 | 许可证 | 角色 |
| --- | --- | --- |
| `symbivela-open` | Apache-2.0 | 公共契约、API schema、示例、SDK 与经批准的核心二进制 —— **本仓库** |
| `symbivela` | AGPL-3.0 | 治理内核的开源核心实现 |
| `symbivela-ee` | 企业版 | 企业功能与内部产品治理文档 |

SYMBIVELA 是**控制表面，而非控制权威**。它拥有工作空间体验与人类交互记录；计划、授权、执行、世界状态与源证据仍由各自权威产品（ORCHADYN、AEGIVELA、PRAXOVELA、RHEOVELA、KINETOVELA、ONTOVELA、TEKMOVELA、Harmovela）持有。

## 本仓库：SYMBIVELA Open

面向集成方的公开 Apache-2.0 表面。

### 契约

- `contracts/openapi.yaml` —— 治理内核 HTTP API v1（受信任上下文头、幂等、工作空间角色）。
- `contracts/governance-kernel-v1.schema.json` —— 核心人类治理对象（Workspace、AttentionItem、HumanDecision、Handoff、Goal）。
- `contracts/orchadyn-plan-projection-v1.schema.json` —— 只读计划投影契约。
- `contracts/aegivela-approval-request-v1.schema.json` —— 审批门禁投影契约。
- `contracts/harmovela-event-envelope-v1.schema.json` —— 运行时事件信封契约。
- `contracts/intervention-request-v1.schema.json` —— 干预请求契约。

### 示例

- `examples/governance-kernel/` —— 受信任上下文 API 用法。

## 快速开始

```bash
# 从 core 仓库本地运行开源核心 API 与 PostgreSQL
docker compose up --build

# 部署前显式执行迁移
DATABASE_URL="postgres://..." go run ./backend/cmd/symbivela-migrate
```

## 文档

- 产品定位、架构、路线图：`symbivela-ee/docs/`
- 公共契约与示例：本仓库。

## 许可证

Apache-2.0。`symbivela` 中的核心实现为 AGPL-3.0；`symbivela-ee` 中的企业扩展使用企业版许可证。
