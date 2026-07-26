---
source: "https://axtary.com"
hn_url: "https://news.ycombinator.com/item?id=49053543"
title: "Show HN: Axtary – Content Authorization for AI Agents"
article_title: "Axtary | Content Authorization for AI Agents"
author: "Axtary"
captured_at: "2026-07-26T01:51:19Z"
capture_tool: "hn-digest"
hn_id: 49053543
score: 1
comments: 0
posted_at: "2026-07-26T01:02:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Axtary – Content Authorization for AI Agents

- HN: [49053543](https://news.ycombinator.com/item?id=49053543)
- Source: [axtary.com](https://axtary.com)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T01:02:04Z

## Translation

タイトル: 表示 HN: Axtary – AI エージェントのコンテンツ認証
記事タイトル: アクスタリー | AI エージェントのコンテンツ認証
説明: Axtary は、AI エージェントが GitHub、Slack、MCP ツール、または運用 API にアクセスする前に、正確なツール呼び出し、ペイロード、ポリシー、承認をチェックします。

記事本文:
アクスタリー | AI エージェントのコンテンツ認証 Axtary ランタイム デモ ポリシー パッケージ ブログ ドキュメント サインイン アクセスのリクエスト コンテンツ
認可
AIエージェント向け。
Axtary は、コネクタが実行される前に、正確な diff、メッセージ、クエリ、またはツールのペイロードをチェックします。日常的な行動はポリシーに従います。リスクの高いアクションには、その正確なペイロードの承認が必要です。
エージェントが間違いを犯したり侵害された場合、その権限は承認されたアクションに限定されたままになります。すべての試みは記録されます。
Axtary の決定: allowed action_pass_id : ap_01JAXTARY Expires_in : 10m payload_hash : sha256:7f32… ホット パス トークンはチャネルを承認します。 Axtary はコンテンツを承認します。
Axtary は、コード変更、インフラストラクチャ パス、データ アクセス、メッセージ、チケット、ドキュメント、MCP ツールなど、セキュリティを重視したエンジニアリングとエージェントの運用を管理します。強制はエージェントのそばで実行されるため、プロバイダーの資格情報はローカルに残ります。日常的なアクションは決定的なポリシーに従います。リスクの高いアクションには、正確なペイロードの承認が必要です。
アクションを正規化する アクター、タスク、リソース、制約、および正確なペイロード ハッシュを 1 つのアクション レコードにキャプチャします。ポリシーを評価して、実行前にパス、ファイル、テスト、ツール定義、および運用に影響を与えるルールを適用します。 ActionPass を発行します。 承認されたアクション、ペイロード ハッシュ、ポリシー バージョン、制約、および有効期限に署名します。結果を記録する 決定、合格、実行結果、およびトレース参照を検証可能な台帳に書き込みます。ペイロード バインディング ペイロードを変更します。検証が失敗します。
人間の承認は、レビューされたペイロード ハッシュに関連付けられます。承認後にペイロードが変更された場合、プロバイダーが呼び出される前にアダプター側の検証によって不一致が拒否されます。承認および提示されたハッシュは台帳に記録されます。
署名されたパスは、承認されたアクションにバインドされます。台帳には検証記録が保存されます。
lを再現します

ローカル: axtary 実行ワークフロー github-pr-review --real --tamper
承認は概要ではなくペイロードにバインドされます。
ActionPass は、セキュリティをレビューし、SDK、プロキシ、MCP ラッパー全体で使用できるように設計されています。人間の承認は正規化されたアクションとペイロード ハッシュに署名し、レビュー後に承認が変更されるのを防ぎます。その結果、ツールを使用するための広範な許可ではなく、特定のアクションに対する許可が得られます。
github.pull_requests.create（ブランチ、ファイル、パス、テスト制約あり）slack.chat.postMessage（チャネルスコープと受信者ステップアップ）プロジェクト、担当者、ステータス、フィールド制約ありlinear.issue.updateプロジェクト、バケット、リージョン、プレフィックスでスコープ設定されたAWSとGCP読み取りサーバーIDとツール定義ハッシュにバインドされたmcp.tool.callルート、結果、バイト、トラバーサル制限付きdocs.search/readブロックされたシークレットおよび環境パスを使用した github.contents.read/write 非破壊的な ID 証拠によるコネクタの準備状況チェック 正規化されたアクション {
"action_pass_id": "ap_01JAXTARY",
"agent_id": "agent:codex-prod",
"human_owner": "user:reviewer@company.com",
"intent": "AXT-418 の PR を開く",
"ツール": "github.pull_requests.create",
"リソース": "リポジトリ:会社/ウェブアプリ",
「制約」: {
"ベースブランチ": "メイン",
"max_files_changed": 12、
"blocked_paths": ["infra/prod/**", ".env*"],
"requires_tests": true
}、
"expires_in": "10m",
"ペイロード_ハッシュ": "sha256:7f32...",
"ポリシー": "シダー+レゴ:パス",
"ledger_hash": "sha256:b9a1..."
Cedar 互換のポリシー許可 (
プリンシパル == エージェント::"codex-prod",
アクション == アクション::"github.pull_requests.create",
resource == リポジトリ::"会社/ウェブアプリ"
) のとき {
context.intent.task_id == "AXT-418" &&
context.payload.max_files_changed <= 12 &&
!context.payload.touches_production
};コネクタとランタイム
MCP サーバーとサポートされているネイティブ コネクタを管理します。

チームがすでに使用しているランタイム内で、同じポリシー、ActionPass、台帳を通じて。
モデル コンテキスト プロトコル GitHub GitHub Slack Slack Linear Linear Jira Jira Sentry Sentry PostgreSQL PostgreSQL AWS AWS Google Cloud Google Cloud Google Drive Google Drive Microsoft (近日登場) Microsoft (近日登場) Okta (近日登場) Okta (近日登場) Auth0 (近日登場) Auth0 (近日登場) ランタイム Anthropic Anthropic OpenAI OpenAI Cursor Cursor ランタイム パッケージ 強制実行エージェントの横で。チームはホストされたコントロール プレーンを通じて調整します。
エージェントが実行される場所に Axtary をインストールします。 SDK とローカル プロキシは、ツールを実行する前にアクション ポリシーを適用します。ホストされたアプリは、承認、ポリシー、監査エクスポートを調整します。
ActionPass のドラフトとベリファイアは公開成果物です。仕様とベリファイアのガイドをお読みください。
npm i -g @axtary/cli && axtary init && axtary デモ
施行と認証情報はローカルのままです。ホストされたコントロール プレーンは、アクション データ パスに入ることなく、チーム ポリシーとレビューを調整します。
ポリシー:
ギットハブ:
プルリクエスト:
必須ベースブランチ: main
変更された最大ファイル数: 12
拒否パスプレフィックス: [.env, Secrets/]
stepUpPathPrefixes: [infra/prod/, billing/]
必要なテスト: true
緩み:
メッセージ:
allowedChannels: ["#axtary-dev"] Apache-2.0 · npm 上の v0.5.0
@axtary/actionpass 署名付き認可アーティファクト 正規化されたアクション、ペイロード ハッシュ、ポリシー バージョン、制約、および有効期限をポータブル パスに署名します。 npm i @axtary/actionpass @axtary/policy 決定論的ポリシー エンジン 正規化されたアクションをローカルで評価し、同じ決定コンテキストを Cedar および OPA 入力にマッピングします。 npm i @axtary/policy @axtary/proxy ローカル施行プロキシ ローカル施行ポイントでのポリシー評価、ActionPass 発行、実行チェック、台帳記録を組み合わせます。 npm i @axtary/proxy @axtary/ledger 検証可能なアクション

n 元帳 署名された証明書と独立して検証可能な包含および一貫性の証明を備えた追加専用のハッシュ チェーンを維持します。 npm i @axtary/ledger @axtary/approvals ペイロードにバインドされた承認 人間の各決定を、レビューのために提示されたアクションおよびペイロード ハッシュにバインドします。 npm i @axtary/approvals @axtary/mcp MCP 来歴制御 サーバー ID、ツール スキーマ、定義ハッシュ、および実行前の承認状態を記録します。 npm i @axtary/mcp @axtary/adapters スコープ付きコネクタ アダプタ サポートされているコネクタ アクション全体にプロバイダー固有のスコープと証拠チェックを適用します。 npm i @axtary/adapters @axtary/cli runtime CLI ローカル プロキシ、コネクタ ワークフロー、検証コマンド、および動作チェックをエージェントのそばで実行します。 npm i -g @axtary/cli @axtary/config 型付き設定 サポートされていない、または不完全な設定の型付きエラーを含む axtary.yml をロードして検証します。 npm i @axtary/config Axtary ツール、MCP サーバー、実稼働システムにわたるエージェント アクションのコンテンツ認証。
すでにアカウントをお持ちですか?サインイン

## Original Extract

Axtary checks the exact tool call, payload, policy, and approval before an AI agent touches GitHub, Slack, MCP tools, or production APIs.

Axtary | Content Authorization for AI Agents Axtary Runtime Demo Policy Packages Blog Docs Sign in Request access Content
authorization
for AI agents.
Axtary checks the exact diff, message, query, or tool payload before a connector executes. Routine actions follow policy; higher-risk actions require approval of that exact payload.
If an agent is mistaken or compromised, its authority remains limited to the approved action. Every attempt is recorded.
Axtary decision: allow action_pass_id : ap_01JAXTARY expires_in : 10m payload_hash : sha256:7f32… Hot path Tokens authorize channels. Axtary authorizes content.
Axtary governs security-sensitive engineering and agent operations, including code changes, infrastructure paths, data access, messages, tickets, documents, and MCP tools. Enforcement runs beside the agent so provider credentials remain local. Routine actions follow deterministic policy. Higher-risk actions require approval of the exact payload.
Normalize the action Capture the actor, task, resource, constraints, and exact payload hash in one action record. Evaluate policy Apply path, file, test, tool-definition, and production-impact rules before execution. Issue the ActionPass Sign the approved action, payload hash, policy version, constraints, and expiry. Record the outcome Write the decision, pass, execution result, and trace reference to a verifiable ledger. Payload binding Change the payload. Verification fails.
Human approval is bound to the reviewed payload hash. If the payload changes after approval, adapter-side verification rejects the mismatch before the provider is called. The approved and presented hashes are recorded in the ledger.
The signed pass is bound to the approved action. The ledger preserves the verification record.
Reproduce locally: axtary run workflow github-pr-review --real --tamper
Approval is bound to the payload, not a summary.
ActionPass is designed for security review and use across SDKs, proxies, and MCP wrappers. Human approval signs the normalized action and payload hash, preventing authorization from changing after review. The result is authorization for a specific action, rather than broad permission to use a tool.
github.pull_requests.create with branch, file, path, and test constraints slack.chat.postMessage with channel scope and recipient step-up linear.issue.update with project, assignee, status, and field constraints AWS and GCP reads scoped by project, bucket, region, and prefix mcp.tool.call bound to server identity and tool definition hash docs.search/read with root, result, byte, and traversal limits github.contents.read/write with blocked secret and environment paths Connector readiness checks with non-destructive identity evidence Normalized action {
"action_pass_id": "ap_01JAXTARY",
"agent_id": "agent:codex-prod",
"human_owner": "user:reviewer@company.com",
"intent": "Open a PR for AXT-418",
"tool": "github.pull_requests.create",
"resource": "repo:company/web-app",
"constraints": {
"base_branch": "main",
"max_files_changed": 12,
"blocked_paths": ["infra/prod/**", ".env*"],
"requires_tests": true
},
"expires_in": "10m",
"payload_hash": "sha256:7f32...",
"policy": "cedar+rego:pass",
"ledger_hash": "sha256:b9a1..."
} Cedar-compatible policy permit (
principal == Agent::"codex-prod",
action == Action::"github.pull_requests.create",
resource == Repo::"company/web-app"
) when {
context.intent.task_id == "AXT-418" &&
context.payload.max_files_changed <= 12 &&
!context.payload.touches_production
}; Connectors and runtimes
Govern MCP servers and supported native connectors through the same policy, ActionPass, and ledger, inside the runtimes your teams already use.
Model Context Protocol GitHub GitHub Slack Slack Linear Linear Jira Jira Sentry Sentry PostgreSQL PostgreSQL AWS AWS Google Cloud Google Cloud Google Drive Google Drive Microsoft (coming soon) Microsoft (coming soon) Okta (coming soon) Okta (coming soon) Auth0 (coming soon) Auth0 (coming soon) Runtimes Anthropic Anthropic OpenAI OpenAI Cursor Cursor Runtime packages Enforcement runs beside the agent. Teams coordinate through the hosted control plane.
Install Axtary where agents run. SDKs and the local proxy enforce action policy before tools execute. The hosted app coordinates approvals, policies, and audit exports.
The ActionPass draft and verifier are public artifacts: read the spec and verifier guide .
npm i -g @axtary/cli && axtary init && axtary demo
Enforcement and credentials remain local. The hosted control plane coordinates team policy and review without entering the action data path.
policy:
github:
pullRequests:
requiredBaseBranch: main
maxFilesChanged: 12
denyPathPrefixes: [.env, secrets/]
stepUpPathPrefixes: [infra/prod/, billing/]
requiresTests: true
slack:
messages:
allowedChannels: ["#axtary-dev"] Apache-2.0 · v0.5.0 on npm
@axtary/actionpass signed authorization artifact Signs the normalized action, payload hash, policy version, constraints, and expiry into a portable pass. npm i @axtary/actionpass @axtary/policy deterministic policy engine Evaluates normalized actions locally and maps the same decision context to Cedar and OPA inputs. npm i @axtary/policy @axtary/proxy local enforcement proxy Combines policy evaluation, ActionPass issuance, execution checks, and ledger recording at the local enforcement point. npm i @axtary/proxy @axtary/ledger verifiable action ledger Maintains an append-only hash chain with signed attestations and independently verifiable inclusion and consistency proofs. npm i @axtary/ledger @axtary/approvals payload-bound approvals Binds each human decision to the action and payload hash presented for review. npm i @axtary/approvals @axtary/mcp MCP provenance controls Records server identity, tool schema, definition hash, and approval state before execution. npm i @axtary/mcp @axtary/adapters scoped connector adapters Applies provider-specific scope and evidence checks across supported connector actions. npm i @axtary/adapters @axtary/cli runtime CLI Runs the local proxy, connector workflows, verification commands, and operational checks beside the agent. npm i -g @axtary/cli @axtary/config typed configuration Loads and validates axtary.yml with typed errors for unsupported or incomplete configuration. npm i @axtary/config Axtary Content authorization for agent actions across tools, MCP servers, and production systems.
Already have an account? Sign in
