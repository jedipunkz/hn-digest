---
source: "https://github.com/sammysltd/makerchecker"
hn_url: "https://news.ycombinator.com/item?id=48571738"
title: "Show HN: Stop your AI agents from approving their own work"
article_title: "GitHub - sammysltd/makerchecker: Roles, segregation of duties, and tamper-evident audit for the AI agents you already run. Open-source, self-hosted. · GitHub"
author: "smashini"
captured_at: "2026-06-17T16:21:59Z"
capture_tool: "hn-digest"
hn_id: 48571738
score: 2
comments: 1
posted_at: "2026-06-17T15:22:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Stop your AI agents from approving their own work

- HN: [48571738](https://news.ycombinator.com/item?id=48571738)
- Source: [github.com](https://github.com/sammysltd/makerchecker)
- Score: 2
- Comments: 1
- Posted: 2026-06-17T15:22:23Z

## Translation

タイトル: HN を表示: AI エージェントが自分の作業を承認しないようにする
記事のタイトル: GitHub - sammysltd/makerchecker: すでに実行している AI エージェントの役割、職務の分離、改ざん防止監査。オープンソース、自己ホスト型。 · GitHub
説明: すでに実行している AI エージェントの役割、職務の分離、および改ざん明示的な監査。オープンソース、自己ホスト型。 - sammysltd/makerchecker

記事本文:
GitHub - sammysltd/makerchecker: すでに実行している AI エージェントの役割、職務の分離、改ざん防止監査。オープンソース、自己ホスト型。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
サミー株式会社
/
メーカーチェッカー
公共
通知
通知を変更するにはサインインする必要があります

フィクション設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット .githooks .githooks .github .github docs docs サンプル サンプル ops ops パッケージ パッケージ スクリプト スクリプト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml .nvmrc .nvmrc CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス LICENSING.md LICENSING.md 通知 通知 README.md README.md SECURITY.md SECURITY.md docker-compose.hardened.yml docker-compose.hardened.yml docker-compose.yml docker-compose.yml eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json turbo.jsonturbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントがお金を動かしました。誰もそれを承認しませんでした。
MakerChecker は、構造的な強制と人間の承認を通じて AI エージェントを管理する自己ホスト型ソフトウェアです。構造的強制は、パス内に人間がいない状態でマシンの速度で実行されます。エージェントはロールを通じてのみ動作し、そのロールに付与されたスキルのみを実行し（デフォルトで拒否され、正確なバージョンに固定されます）、その制限を超えることはできず、おそらく自身の作業を承認することはできません。人間の承認は、ルールによって指名された人の署名が必要な、いくつかの高リスクのアクションに対して保留されます。すべてのアクションは、ハッシュチェーンされ、Ed25519 で署名された監査ログにコミットされ、誰もがオフラインで検証できます。 1 行変更すると、そこで検証が中断されます。
エージェントは既存のフレームワークで実行し続けます。 MakerChecker は、その前にあるチェックポイントであり、その後ろにあるレコード、つまり Postgres 上の Fastify サーバーです。エージェントはフローとして接続します (MakerChecker がステップとゲートを実行します)、または

プロキシ セッション (MakerChecker は、フレームワークが実行するツール呼び出しを承認および記録します)。どちらも同じ監査チェーンを作成します。
ここは新しいですか？オペレーター → クイックスタート 。インテグレーター → 統合 。セキュリティレビュー担当者 → docs/security-model.md 。検査者 → docs/audit-spec.md 。
ライブデモ : エージェントは許可を超過したり自身の作業を承認したりすることがブロックされ、実行の監査チェーンはオフラインで検証され、指名された人間はルールで要求されている場合にのみサインオフします。サインアップはありません。
付与。ロールを正確なスキル バージョンにバインドします。他には何も実行されません。
チェック。すべてのツール呼び出しは最初にゲートに到達します。許可がない、制限を超えている、または SoD 制約に違反しているため、ツール本体が実行される前に拒否されます。
ゲート。高リスクのステップは、指定された人間の承認を待ちます。要求者は自分自身を承認することはできません。
記録。状態の変化とツール呼び出しは同じトランザクション内の監査チェーンにコミットされ、各イベントはハッシュによって最後まで連鎖されます。
すべての拒否には名前が付けられ、監査されます。
ドッカーの構成
Postgres とサーバーはポート 3000 で起動します。最初にブートするとデモがシードされ、管理者キーと役員キーが出力されます。これらをログからコピーします。
実稼働 (非デモ) デプロイメントでは、何もシードされません。ノード dist/cli.js bootstrap-admin --email <e> --name <n> を使用して、最初の管理者とその API キーを明示的にミントします (1 回出力されます)。 「新規デプロイメントの最初の管理者」を参照してください。
メーカーチェッカー制約を備えた現金調整フローがシードされ、準備が整いました。
import H= ' authorization: Bearer mk_... ' # ログからの管理キー
# フローをトリガーする
curl -X POST localhost:3000/api/flows/daily-cash-reconciliation/runs -H " $H " -H ' content-type: application/json ' -d ' {} '
# 保留中の承認ゲートを検査する
カールローカルホスト:3000/api/approvals -H " $H "
# ゲートを承認する
curl -X POST localhost:3000/api/approvals/ < id > /decion -H " $H " -H ' content-type: a

アプリケーション/json ' \
-d ' {"決定":"承認","理由":"例外が解決されました"} '
# 監査チェーンを検証する
カールローカルホスト:3000/api/audit/verify -H " $H "
完全なローカル セットアップと実際のモデルでの実行は、 docs/quickstart.md にあります。
クイックスタートは Postgres 所有者として接続し、追加のみの監査トリガーを無効にします。侵害されたアプリ資格情報に対する改ざん防止のために、 docker-compose.hardened.yml を使用してサーバーを非所有者ロールとして実行します (ウォークスルー)。
Turborepo を使用した pnpm ワークスペース。サーバーはAGPL-3.0です。 SDK とコネクタは Apache-2.0 であるため、クローズド ソース コードに埋め込むことができます。
管理されるプリミティブ (エージェント、ロール、スキル、トリガー、フロー、実行/監査) は Postgres によってサポートされ、バージョン管理されています。 docs/concepts.md を参照してください。
プロキシ セッションを開き、各ツールをラップします。すべての呼び出しで proxy.check が実行され (拒否すると、ツールの実行前に GovernanceDeniedError がスローされます)、ツールが実行され、出力が記録されます。高リスクのスキルはプロキシ パスでは拒否されます。彼らにはフローゲートが必要です。
"@makerchecker/sdk" からインポート { createClient , gownedTool , GovernanceDeniedError } ;
const client = createClient({baseUrl : "http://localhost:3000" , apiKey : "mk_..." } ) ;
const {セッション} = クライアントを待ちます。プロキシ 。 openSession ( { ラベル : "recon-run" } ) ;
const match = 管理ツール (
クライアント、
セッション。 ID 、
"recon-preparer" , // ロール付与が評価される登録済みエージェント
"txn-match@1" , // skillRef: name @version
(入力 : { ステートメント : 不明 [ ] ; 元帳 : 不明 [ ] } ) => matchTxns (入力 ) 、
) ;
一致を待ちます ({ ステートメント , 台帳 } ) ; // 拒否された場合は GovernanceDeniedError をスローします
クライアントを待ちます。プロキシ 。 closeSession (セッション.id) ;
コネクタはツールの名前、説明、スキーマを保持します。
LangChain — GovernmentLangChainTool は DynamicStructuredTool を返します。例/接続を参照

tors/langchain 。
クロード エージェント SDK — gownClaudeTool は、 createSdkMcpServer の SdkMcpToolDefinition を返します。 「packages/connector-claude-agent」 を参照してください。
Python (CrewAI、LangChain、LlamaIndex、AutoGen) — create_client の次に、governed_tool ; pip インストール メーカーチェッカー 。 「packages/sdk-python」を参照してください。
すべての状態遷移は、状態書き込みと同じトランザクションで監査イベントを発行します。各イベントのハッシュは、イベントの RFC 8785 正規 JSON ( seq を除く) 上の SHA-256 であり、インスタンスに関連付けられたジェネシス イベントから prev_hash を通じてチェーンされます。行を変更すると、その行で再計算が中断されます。
GET /api/audit/verify はチェーンをたどり、ブレーク時に { ok, count, headHash } 、または { ok: false, failedSeq,reason } を返します。 CLI は、ライブ チェーンと署名されたバンドルをオフラインで検証します。
# 実行中のデータベースに対して検証します
docker compose exec サーバー ノード dist/cli.js 監査検証
# 署名されたバンドルをエクスポートし、データベースなしで検証します
docker compose exec サーバー ノード dist/cli.js 監査エクスポート --outbundle.json
ノード dist/cli.js 監査検証バンドル --inbundle.json
node dist/cli.js Audit verify-bundle --inbundle.json --keyinstance.pub # キーを固定する
バンドルは Ed25519 で署名されており、チェーンの再計算に必要なマニフェストが含まれています。形式は、任意の言語で再実装できるように docs/audit-spec.md で指定されています。監査レポート --run <id> は、自己完結型の HTML 実行レポートを作成します。 Audit access-review は、role/grant/SoD レビューをレンダリングします (/api/reports/access-review にもあります)。
メーカーチェッカー1.0。サーバー、Web、共有、統合、検証のパスは、CI での Postgres に対する単体テストと統合テストでカバーされます。
1.0 には、ドラッグ アンド ドロップ フロー ビルダー、SSO/SAML、またはマルチテナンシーがありません。フロー定義は JSON/YAML 型で指定されます。
サーバー、Web、および共有は AGPL-3.0 ( LICENSE ) です。 SDK、コネクタ、サンプルは次のとおりです。

e Apache-2.0 — コピーレフトなしでクローズドソース製品としてインポートおよび出荷します。 AGPL-3.0 を使用できない組織は、商用ライセンスを利用できます: hello@makerchecker.ai 。詳細: LICENSING.md 。
貢献者: COTRIBUTING.md · セキュリティ: SECURITY.md · 行動規範: CODE_OF_CONDUCT.md · 変更ログ: CHANGELOG.md
すでに実行している AI エージェントの役割、職務の分離、および改ざん明示的な監査。オープンソース、自己ホスト型。
AGPL-3.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v1.0.0
最新の
2026 年 6 月 16 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Roles, segregation of duties, and tamper-evident audit for the AI agents you already run. Open-source, self-hosted. - sammysltd/makerchecker

GitHub - sammysltd/makerchecker: Roles, segregation of duties, and tamper-evident audit for the AI agents you already run. Open-source, self-hosted. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
sammysltd
/
makerchecker
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .githooks .githooks .github .github docs docs examples examples ops ops packages packages scripts scripts .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml .nvmrc .nvmrc CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LICENSING.md LICENSING.md NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md docker-compose.hardened.yml docker-compose.hardened.yml docker-compose.yml docker-compose.yml eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json turbo.json turbo.json View all files Repository files navigation
Your AI agent moved the money. No one approved it.
MakerChecker is self-hosted software that governs AI agents through structural enforcement and human approvals . Structural enforcement runs at machine speed with no human in the path: an agent acts only through a role , runs only the skills its role was granted (deny by default, pinned to an exact version), cannot exceed its limits, and provably cannot approve its own work. Human approval is reserved for the few high-risk actions where a rule requires a named person to sign. Every action commits to a hash-chained, Ed25519-signed audit log that anyone verifies offline. Change one row and verification breaks at it.
Your agents keep running in their existing framework. MakerChecker is the checkpoint in front of them and the record behind them: a Fastify server on Postgres. Agents connect as a flow (MakerChecker runs the steps and gates) or a proxy session (MakerChecker authorizes and records tool calls your framework executes). Both write the same audit chain.
New here? Operator → Quickstart . Integrator → Integration . Security reviewer → docs/security-model.md . Examiner → docs/audit-spec.md .
Live demo : an agent is blocked from exceeding its grant and from approving its own work, the run's audit chain verifies offline, and a named human signs off only where a rule requires it. No signup.
Grant. Bind a role to exact skill versions. Nothing else runs.
Check. Every tool call hits the gate first. No grant, over a limit, or against an SoD constraint, and it is denied before the tool body runs.
Gate. High-risk steps wait for named human approval. The requester cannot approve their own.
Record. State changes and tool calls commit to the audit chain in the same transaction, each event chained to the last by hash.
Every refusal is named and audited:
docker compose up
Postgres and the server come up on port 3000. First boot seeds the demo and prints an admin key and an officer key — copy them from the logs.
On a production (non-demo) deployment nothing is seeded; mint the first admin and its API key explicitly with node dist/cli.js bootstrap-admin --email <e> --name <n> (printed once). See First admin on a fresh deployment .
A cash-reconciliation flow with a maker-checker constraint is seeded and ready:
export H= ' authorization: Bearer mk_... ' # admin key from the logs
# Trigger the flow
curl -X POST localhost:3000/api/flows/daily-cash-reconciliation/runs -H " $H " -H ' content-type: application/json ' -d ' {} '
# Inspect the pending approval gate
curl localhost:3000/api/approvals -H " $H "
# Approve the gate
curl -X POST localhost:3000/api/approvals/ < id > /decision -H " $H " -H ' content-type: application/json ' \
-d ' {"decision":"approved","reason":"Exceptions resolved"} '
# Verify the audit chain
curl localhost:3000/api/audit/verify -H " $H "
Full local setup, and running with real models, is in docs/quickstart.md .
The quickstart connects as the Postgres owner, which disables the append-only audit triggers. For tamper-resistance against a compromised app credential, run the server as a non-owner role with docker-compose.hardened.yml ( walkthrough ).
pnpm workspaces with Turborepo. The server is AGPL-3.0; the SDKs and connectors are Apache-2.0, so you can embed them in closed-source code.
The governed primitives — Agent, Role, Skill, Trigger, Flow, Run/Audit — are Postgres-backed and versioned. See docs/concepts.md .
Open a proxy session, then wrap each tool. Every call runs proxy.check (a deny throws GovernanceDeniedError before the tool runs), executes the tool, then records the output. High-risk skills are refused on the proxy path; they need a flow gate.
import { createClient , governedTool , GovernanceDeniedError } from "@makerchecker/sdk" ;
const client = createClient ( { baseUrl : "http://localhost:3000" , apiKey : "mk_..." } ) ;
const { session } = await client . proxy . openSession ( { label : "recon-run" } ) ;
const match = governedTool (
client ,
session . id ,
"recon-preparer" , // registered agent whose role grants are evaluated
"txn-match@1" , // skillRef: name @version
( input : { statement : unknown [ ] ; ledger : unknown [ ] } ) => matchTxns ( input ) ,
) ;
await match ( { statement , ledger } ) ; // throws GovernanceDeniedError if denied
await client . proxy . closeSession ( session . id ) ;
Connectors keep your tool's name, description, and schema:
LangChain — governLangChainTool returns a DynamicStructuredTool . See examples/connectors/langchain .
Claude Agent SDK — governClaudeTool returns an SdkMcpToolDefinition for createSdkMcpServer . See packages/connector-claude-agent .
Python (CrewAI, LangChain, LlamaIndex, AutoGen) — create_client then governed_tool ; pip install makerchecker . See packages/sdk-python .
Every state transition emits an audit event in the same transaction as the state write. Each event's hash is SHA-256 over the RFC 8785 canonical JSON of the event (excluding seq ), chained through prev_hash from a genesis event tied to the instance. Change any row and recomputation breaks at it.
GET /api/audit/verify walks the chain and returns { ok, count, headHash } , or { ok: false, failedSeq, reason } on a break. The CLI verifies the live chain and signed bundles offline:
# verify against the running database
docker compose exec server node dist/cli.js audit verify
# export a signed bundle, then verify it with no database
docker compose exec server node dist/cli.js audit export --out bundle.json
node dist/cli.js audit verify-bundle --in bundle.json
node dist/cli.js audit verify-bundle --in bundle.json --key instance.pub # pin the key
Bundles are Ed25519-signed and carry the manifest needed to recompute the chain. The format is specified in docs/audit-spec.md for reimplementation in any language. audit report --run <id> builds a self-contained HTML run report; audit access-review renders the role/grant/SoD review (also at /api/reports/access-review ).
MakerChecker 1.0. The server, web, shared, integration, and verification paths are covered by unit and integration tests against Postgres in CI.
1.0 has no drag-and-drop flow builder, SSO/SAML, or multi-tenancy. Flow definitions are typed JSON/YAML.
Server, web, and shared are AGPL-3.0 ( LICENSE ). The SDKs, connectors, and examples are Apache-2.0 — import and ship them in closed-source products without copyleft. A commercial license is available for organizations that cannot use AGPL-3.0: hello@makerchecker.ai . Details: LICENSING.md .
Contributing: CONTRIBUTING.md · Security: SECURITY.md · Code of Conduct: CODE_OF_CONDUCT.md · Changelog: CHANGELOG.md
Roles, segregation of duties, and tamper-evident audit for the AI agents you already run. Open-source, self-hosted.
AGPL-3.0 license
Code of conduct
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v1.0.0
Latest
Jun 16, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
