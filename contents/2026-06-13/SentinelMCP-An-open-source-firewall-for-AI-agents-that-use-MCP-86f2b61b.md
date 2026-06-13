---
source: "https://github.com/technosiveuk-ui/SentinelMCP"
hn_url: "https://news.ycombinator.com/item?id=48511652"
title: "SentinelMCP – An open-source firewall for AI agents that use MCP"
article_title: "GitHub - technosiveuk-ui/SentinelMCP: The Open-Source MCP Firewall & Security Gateway for AI Agents. Inspect, redact, and control tool calls. Proxy mode (universal) and Inline SDK mode (Go). · GitHub"
author: "technsoive"
captured_at: "2026-06-13T04:34:16Z"
capture_tool: "hn-digest"
hn_id: 48511652
score: 2
comments: 0
posted_at: "2026-06-13T01:51:11Z"
tags:
  - hacker-news
  - translated
---

# SentinelMCP – An open-source firewall for AI agents that use MCP

- HN: [48511652](https://news.ycombinator.com/item?id=48511652)
- Source: [github.com](https://github.com/technosiveuk-ui/SentinelMCP)
- Score: 2
- Comments: 0
- Posted: 2026-06-13T01:51:11Z

## Translation

タイトル: SentinelMCP – MCP を使用する AI エージェント用のオープンソース ファイアウォール
記事のタイトル: GitHub - technosiveuk-ui/SentinelMCP: AI エージェント用のオープンソース MCP ファイアウォールとセキュリティ ゲートウェイ。ツール呼び出しを検査、編集、および制御します。プロキシ モード (ユニバーサル) とインライン SDK モード (Go)。 · GitHub
説明: AI エージェント用のオープンソース MCP ファイアウォールおよびセキュリティ ゲートウェイ。ツール呼び出しを検査、編集、および制御します。プロキシ モード (ユニバーサル) とインライン SDK モード (Go)。 - technosiveuk-ui/SentinelMCP

記事本文:
GitHub - technosiveuk-ui/SentinelMCP: AI エージェント用のオープンソース MCP ファイアウォールおよびセキュリティ ゲートウェイ。ツール呼び出しを検査、編集、および制御します。プロキシ モード (ユニバーサル) とインライン SDK モード (Go)。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
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
テクノシブクウイ
/
センチネルMCP
公共
そうではない

化
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット .github .github アダプター アダプター アセット アセット cmd cmd config config docs docs example/ inline-sdk Example/ inline-sdk ゲートウェイ ゲートウェイ SDK sdk .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.testclient Dockerfile.testclient Dockerfile.upstream Dockerfile.upstream LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum すべてのファイルを表示リポジトリ ファイルのナビゲーション
AI エージェント用のオープンソース MCP セキュリティ ゲートウェイ
SentinelMCP は現在アルファ版です。プロジェクトは現在開発中であり、API、構成形式、グラフの動作は将来のリリースで予告なく変更される可能性があります。
製造上の警告: このソフトウェアは現状のまま提供されます。私たちはセキュリティに努めていますが、徹底したテストを行わずに、高度に規制された運用環境では、アルファ段階のプロキシだけが唯一の防御手段となるべきではありません。ご自身の責任でご使用ください。
私たちは早期導入者とフィードバックを積極的に求めています。問題が発生した場合、提案がある場合、または貢献したい場合は、問題を開いてください。あなたの意見がロードマップを直接形作ります。
SentinelMCP は、実行時に AI エージェント ツール呼び出しを保護するモデル コンテキスト プロトコル (MCP) のセキュリティ強制エンジンです。 AI エージェントと、AI エージェントが呼び出すツールの間に位置し、検査、ポリシーの適用、PII/秘密の編集、監査ログを提供します。
ランタイムネイティブのグラフ オーケストレーション、割り込み/再開機能、ミリ秒未満の強制を備えた組み込みの高性能 Go

— カスタム配管なしで人間参加型の承認ワークフローを実現します。
SentinelMCP は、アーキテクチャに応じて 2 つの方法で導入できます。
HTTP/SSE MCP トラフィックをインターセプトするスタンドアロンのサイドカー バイナリ。あらゆる言語 (Python、TypeScript、Go など) およびあらゆるエージェント フレームワークで動作します。コードの変更は不要です。MCP トラフィックをプロキシ経由でルーティングするだけです。
AI エージェント (任意の言語) → SentinelMCP プロキシ → MCP サーバー
2. インライン SDK モード (ネイティブ化)
アプリケーションに直接インポートされた Go ライブラリ。ゼロコピー グラフ オーケストレーションを使用して、MCP ツール呼び出しをメモリ内でラップします。ミリ秒未満のレイテンシと詳細なコンテキスト認識を提供します。ネットワーク ホップや個別のプロセスは必要ありません。
AI アプリケーション → SentinelMCP SDK (インプロセス) → MCP サーバーに移動
このデュアルモード アーキテクチャが SentinelMCP の主要な差別化要因です。Permit や Envoy などの競合他社はネットワーク プロキシのみを提供しています。 Inline SDK を使用すると、Go アプリケーションはわずかなレイテンシで同じセキュリティを適用できます。
MCP プロキシ (サイドカー) — 任意の MCP クライアント用のドロップイン HTTP/SSE プロキシ。既存のエージェント フレームワークに対して透過的です。
インライン SDK (Go) — インプロセス強制のためのネイティブ Go モジュール。許可パス上のミリ秒未満のオーバーヘッド (19μs p99)。
ポリシー エンジン — ホットリロードを備えたローカル YAML ベースのポリシー定義。リスク レベル (低/中/高) は、強制アクション (許可/編集/ブロック/中断) にマッピングされます。
データ損失防止 (DLP) — ツールの引数と応答における正規表現ベースの PII とシークレットの編集。 6 つの組み込みパターン (秘密キー、パスワード、API キー、クレジット カード、SSN、電子メール) に加えて、カスタム正規表現のサポート。
Human-in-the-Loop (HITL) — 汎用 Webhook 経由でリスクの高いツール呼び出しを中断し、ローカル API エンドポイント経由で再開します。耐久性のある割り込み/再開のための、BoltDB を利用したチェックポイント。
監査ログ — 構造化された JSON ログ

SIEM パイプラインの標準出力およびネイティブ OpenTelemetry (OTel) 統合。
状態管理 — プロセスの再起動後の永続的な割り込み/再開のための、BoltDB を利用したチェックポイント。
🏗️ アーキテクチャと拡張性
SentinelMCP は、保守性、テスト容易性、シームレスなオープンコア エクスペリエンスを保証するために、懸念事項を厳密に分離して構築されています。
ゲートウェイ/パッケージは、基盤となる Eino フレームワークまたは MCP トランスポート ライブラリへの依存関係を持たずに、すべてのコア インターフェイスとドメイン タイプ (ポリシー、DLP、リスク、監査) を定義します。 adapter/eino/ パッケージは、Eino タイプをインポートする唯一のパッケージです。
このアーキテクチャ上の境界には、次の 3 つの大きな利点があります。
フレームワークに依存しない: 基礎となるオーケストレーション フレームワークを変更する必要がある場合、必要なのは新しいアダプターのみです。コアのセキュリティ エンジンはそのまま残ります。
独立してテスト可能: コア ドメイン ロジック (ポリシー評価、DLP 編集、リスク分析) は、MCP サーバーや LLM グラフを起動することなく、きれいに単体テストできます。
アーキテクチャ的に強化されたオープンコア: エンタープライズ実装 (Teams/Slack HITL、Nightfall DLP、コントロール プレーン API) は、インターフェイスを介して GatewayConfig に簡単に交換できます。この OSS コードベースへの変更は必要ありません。
1. Docker でサイドカーを開始します。
docker run --name Sentinlmcp \
-p 8080:8080 -p 9090:9090 \
-v ./policies.yaml:/etc/sentinlmcp/config.yaml \
ghcr.io/technosiveuk-ui/sentinlmcp:最新
ポート 8080 は MCP プロキシです。 9090 は管理 API (ヘルスチェックおよび高リスク割り込みによって使用される承認再開エンドポイント) です。 -d を追加して切り離しを実行し、 docker logs -f Sentinlmcp でログを追跡します。
2. ポリシー (policy.yaml) を定義します。
スキーマのバージョン: " 1.0 "
グローバル :
default_risk : 明示的なエントリのないツールのリスクは低い # (より厳密なデフォルトの場合は中を設定)
redaction_mask : " ***編集済み*** "
サイドカー

:
listen_addr : " :8080 "
health_addr : " :9090 "
トランスポート: streamable_http
upstream_servers : # MCP ツール サーバー — 少なくとも 1 つが必要です
- 名前 : マイツール
URL : " http://localhost:3001/mcp "
ツール:
読み取りファイル:
リスク : 低 # 許可 — 引数と応答は引き続き DLP スキャンされます
redact_patterns : [パスワード、API_KEY]
書き込みファイル:
リスク : 中 # 編集 — 機密性の高い引数フィールドはツールの実行前にマスクされます
実行コマンド:
リスク: 高 # 割り込み — 人間の承認のために一時停止し、その後実行またはブロックします
require_approval : true
Approved_reason : " シェルコマンドはシステム状態を変更できます "
" db_* " : # glob パターンがサポートされています
リスク：高い
require_approval : true
dlp_patterns : # 6 つのパターンが組み込まれています (PRIVATE_KEY、PASSWORD、API_KEY、
INTERNAL_HOSTNAME : # CREDIT_CARD、SSN、EMAIL);必要に応じてここで独自に定義します
正規表現: ' \b[a-z]+\.internal\.company\.com\b '
タイプ:pii
3. MCP クライアントをプロキシに向けます。
http://localhost:8080/mcp
プロキシは標準の Streamable HTTP MCP を話すため、Python、TypeScript、Go などのクライアントは SDK を変更せずにその URL に接続します。 upstream_servers を独自の MCP ツール サーバーに指定すると、プロキシは呼び出しごとに上記のポリシーを適用します。
安全なツールは、小規模なビルダー API を使用してインプロセスで呼び出します。サイドカーやネットワーク ホップはなく、ミリ秒未満のオーバーヘッドはありません。完全なガイド: docs/INLINE-SDK.md 。
v0.2.0以降が必要です。インライン SDK ( sdk/package) は v0.2.0 で出荷されました。
パッケージメイン
インポート(
「コンテキスト」
「fmt」
「ログ」
「github.com/technosiveuk-ui/sentinlmcp/gateway」
「github.com/technosiveuk-ui/sentinlmcp/sdk」
)
関数メイン () {
// プレーンな Go 関数を安全なツールとして登録します。
実行者:= SDK 。 NewFuncInvoker()。
Register ( "read_file" , func ( _ context . Context , args map [ string ] any ) ( string , error ) {
fmtを返します。 Sprintf ( "%v の内容" , args

[ "パス" ])、nil
})。
Register ( "exec_command" , func ( _ context . Context , args map [ string ] any ) ( string , error ) {
fmtを返します。 Sprintf ( "ran %v" , args [ "cmd" ])、nil
})
// パイプラインを構築します。 StrictDefaults は、未知のツールを秘匿化するようにルーティングします。
パイプライン、エラー:= SDK。新しい (呼び出し元)。
WithRisk ( "exec_command" 、ゲートウェイ . RiskHigh )。 // 承認のために中断します
StrictDefaults ()。
ビルド ()
エラーの場合 != nil {
ログ。致命的 (エラー)
}
// すべての呼び出しが検査され、DLP スキャン、ポリシー チェック、および監査が行われます。
結果、エラー:= パイプライン。 Run ( context . Background (), "read_file" , map [ string ] any {
"パス" : "/etc/config.yaml" ,
})
エラーの場合 != nil {
ログ。致命的 (エラー)
}
fmt 。 Println ( "結果:" , 結果 )
}
完全な許可/編集/割り込みの例を実行します。
./examples/inline-sdk を実行してください
仕組み
フローチャート LR
A[AI エージェント] --> S[センチネルMCP]
S --> I[検査ツール呼び出し]
I --> D{政策決定}
D -->|許可| E[ツールの実行]
D -->|編集| R[PII を編集] --> E
D -->|ブロック| X[エラーを返す]
D -->|割り込み| H[人間の承認]
H -->|承認済み| E
H -->|拒否| ×
E --> O[応答の検査]
O --> A
E --> MCP[MCPサーバー]
MCP --> O
読み込み中
パイプラインの流れ:
検査 — ツールの引数をシリアル化し、DLP スキャンを実行し、ポリシー DB でリスク レベルを検索します。
ポリシーの決定 — リスク別のルート: 低→許可、中→秘匿化、高→承認のために中断、またはブロック
実行 — 上流の MCP ツールを呼び出します (該当する場合は編集された引数を使用)
応答の検査 — ツールの応答を DLP スキャンし、検出結果を編集し、構造化された監査イベントを発行します。
センチネルCP/
§── ゲートウェイ/ # コアドメイン (ZERO Eino インポート)
│ §── types.go # GatewayContext、AuditEvent、DLPFinding、RiskLevel、Decision
│ §──graph.go # Pipeline、ToolInvoker、GatewayConfig、MetricsRecorder
│ §──policy.go # ポリシーインターフェイス + DefaultPol

氷っぽい
│ §── redact.go # DLPScanner, Redactor + RegexDLPScanner + DefaultRedactor
│ §── dlp_multi.go # MultiDLPScanner (正規表現 + 外部の合成)
│ §── dlp_external.go # DLPEndpoint インターフェース + ExternalDLPScanner アダプター
│ §──riskdb.go # RiskDB インターフェース + YAMLRiskDB
│ §── Audit.go # AuditEmitter インターフェース + Stdout/FileAuditEmitter
│ §── Audit_sink.go # AuditSink + WriterAuditSink + CompositeAuditEmitter
│ §── metrics.go # MetricsRecorder インターフェース + NOPMetricsRecorder
│ └── webhook_approval.go # WebhookApprovalProvider + CLIApprovalProvider
§── sdk/ # インライン SDK: ゲートウェイ + アダプター上の人間工学に基づいたビルダー
│ §── builder.go # ビルダー API (New、With*、StrictDefaults、Build)
│ §── func_invoker.go # FuncInvoker — 安全なプレーン Go 関数
│ └──defaults.go # 便利なコンストラクター (BuiltinDLP、StdoutAudit、…)
§──adapter/eino/ # Eino フレームワーク アダプター (Eino インポートを含むパッケージのみ)
│ §──graph.go # BuildGraph() → 割り込み/再開付きの 3 ノード グラフ
│ §── otel_metrics.go # OTelMetricsRecorder (カウンター + ヒストグラム)
│ §── otel_audit_emitter.go # OTelAuditEmitter (監査 → スパンイベント)
│ └─bolt_checkpoint.go #BoltCheckPointStore
§── アダプター/サイドカー/ # サイドカー プロキシ アダプター
│ §── invoker.go # SidecarInvoker (上流の MCP サーバーへのルート)
│ └── Discovery.go # DiscoverTools (上流 MCP サーバー)

[切り捨てられた]

## Original Extract

The Open-Source MCP Firewall & Security Gateway for AI Agents. Inspect, redact, and control tool calls. Proxy mode (universal) and Inline SDK mode (Go). - technosiveuk-ui/SentinelMCP

GitHub - technosiveuk-ui/SentinelMCP: The Open-Source MCP Firewall & Security Gateway for AI Agents. Inspect, redact, and control tool calls. Proxy mode (universal) and Inline SDK mode (Go). · GitHub
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
technosiveuk-ui
/
SentinelMCP
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits .github .github adapter adapter assets assets cmd cmd config config docs docs examples/ inline-sdk examples/ inline-sdk gateway gateway sdk sdk .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.testclient Dockerfile.testclient Dockerfile.upstream Dockerfile.upstream LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum View all files Repository files navigation
The Open-Source MCP Security Gateway for AI Agents
SentinelMCP is currently in Alpha . The project is under active development and APIs, configuration formats, and graph behaviors may change in future releases without advance notice.
Production warning: This software is provided as-is. While we strive for security, an Alpha-stage proxy should not be your sole defense in a highly regulated production environment without thorough testing. Use at your own risk.
We actively seek early adopters and feedback. If you encounter issues, have suggestions, or want to contribute — please open an issue . Your input directly shapes the roadmap.
SentinelMCP is a security enforcement engine for the Model Context Protocol (MCP) that secures AI agent tool calls at runtime. It provides inspection, policy enforcement, PII/secret redaction, and audit logging — sitting between your AI agents and the tools they invoke.
Built in high-performance Go with runtime-native graph orchestration, interrupt/resume capabilities, and sub-millisecond enforcement — enabling human-in-the-loop approval workflows without custom plumbing.
SentinelMCP can be deployed in two ways, depending on your architecture:
A standalone sidecar binary that intercepts HTTP/SSE MCP traffic. Works with any language (Python, TypeScript, Go, etc.) and any agent framework. Zero code changes required — just route your MCP traffic through the proxy.
AI Agent (any language) → SentinelMCP Proxy → MCP Server
2. Inline SDK Mode (Go Native)
A Go library imported directly into your application. Wraps MCP tool calls in-memory with zero-copy graph orchestration. Provides sub-millisecond latency and deep context awareness — no network hop, no separate process.
Go AI Application → SentinelMCP SDK (in-process) → MCP Server
This dual-mode architecture is SentinelMCP's key differentiator: competitors like Permit or Envoy only offer network proxies. With the Inline SDK, Go applications get the same security enforcement at a fraction of the latency.
MCP Proxy (Sidecar) — Drop-in HTTP/SSE proxy for any MCP client. Transparent to existing agent frameworks.
Inline SDK (Go) — Native Go module for in-process enforcement. Sub-millisecond overhead on the Allow path (19μs p99).
Policy Engine — Local YAML-based policy definitions with hot-reloading. Risk levels (low/medium/high) map to enforcement actions (allow/redact/block/interrupt).
Data Loss Prevention (DLP) — Regex-based PII and secret redaction in tool arguments and responses. 6 built-in patterns (private keys, passwords, API keys, credit cards, SSNs, emails) plus custom regex support.
Human-in-the-Loop (HITL) — Interrupt high-risk tool calls via generic Webhooks and resume via a local API endpoint. BoltDB-backed checkpoints for durable interrupt/resume.
Audit Logging — Structured JSON logging to stdout and native OpenTelemetry (OTel) integration for SIEM pipelines.
State Management — BoltDB-backed checkpoints for durable interrupt/resume across process restarts.
🏗️ Architecture & Extensibility
SentinelMCP is built with a strict separation of concerns to ensure maintainability, testability, and a seamless Open-Core experience.
The gateway/ package defines all core interfaces and domain types (Policy, DLP, Risk, Audit) with zero dependencies on the underlying Eino framework or MCP transport libraries. The adapter/eino/ package is the only package that imports Eino types.
This architectural boundary provides three major benefits:
Framework Agnostic: If the underlying orchestration framework ever needs to change, only a new adapter is required; the core security engine remains untouched.
Independently Testable: Core domain logic (policy evaluation, DLP redaction, risk analysis) can be unit-tested cleanly without spinning up MCP servers or LLM graphs.
Architecturally Enforced Open-Core: Enterprise implementations (Teams/Slack HITL, Nightfall DLP, Control Plane APIs) simply swap into the GatewayConfig via interfaces. Zero changes to this OSS codebase are required.
1. Start the sidecar with Docker:
docker run --name sentinelmcp \
-p 8080:8080 -p 9090:9090 \
-v ./policies.yaml:/etc/sentinelmcp/config.yaml \
ghcr.io/technosiveuk-ui/sentinelmcp:latest
Port 8080 is the MCP proxy; 9090 is the admin API (health checks and the approval-resume endpoint used by high-risk interrupts). Add -d to run detached, then follow logs with docker logs -f sentinelmcp .
2. Define your policy ( policies.yaml ):
schema_version : " 1.0 "
global :
default_risk : low # risk for tools with no explicit entry (set medium for a stricter default)
redaction_mask : " ***REDACTED*** "
sidecar :
listen_addr : " :8080 "
health_addr : " :9090 "
transport : streamable_http
upstream_servers : # your MCP tool servers — at least one is required
- name : my-tools
url : " http://localhost:3001/mcp "
tools :
read_file :
risk : low # allow — args and response are still DLP-scanned
redact_patterns : [PASSWORD, API_KEY]
write_file :
risk : medium # redact — sensitive argument fields are masked before the tool runs
exec_command :
risk : high # interrupt — pauses for human approval, then runs or blocks
require_approval : true
approval_reason : " Shell commands can modify system state "
" db_* " : # glob patterns are supported
risk : high
require_approval : true
dlp_patterns : # six patterns are built in (PRIVATE_KEY, PASSWORD, API_KEY,
INTERNAL_HOSTNAME : # CREDIT_CARD, SSN, EMAIL); define your own here as needed
regex : ' \b[a-z]+\.internal\.company\.com\b '
type : pii
3. Point your MCP client at the proxy:
http://localhost:8080/mcp
The proxy speaks standard Streamable HTTP MCP, so any client — Python, TypeScript, or Go — connects to that URL with no SDK changes. Point upstream_servers at your own MCP tool server(s) and the proxy enforces the policy above on every call.
Secure tool calls in-process with a small builder API — no sidecar, no network hop, sub-millisecond overhead. Full guide: docs/INLINE-SDK.md .
Requires v0.2.0+. The Inline SDK ( sdk/ package) shipped in v0.2.0.
package main
import (
"context"
"fmt"
"log"
"github.com/technosiveuk-ui/sentinelmcp/gateway"
"github.com/technosiveuk-ui/sentinelmcp/sdk"
)
func main () {
// Register plain Go functions as secured tools.
invoker := sdk . NewFuncInvoker ().
Register ( "read_file" , func ( _ context. Context , args map [ string ] any ) ( string , error ) {
return fmt . Sprintf ( "contents of %v" , args [ "path" ]), nil
}).
Register ( "exec_command" , func ( _ context. Context , args map [ string ] any ) ( string , error ) {
return fmt . Sprintf ( "ran %v" , args [ "cmd" ]), nil
})
// Build the pipeline. StrictDefaults routes unknown tools to redact.
pipeline , err := sdk . New ( invoker ).
WithRisk ( "exec_command" , gateway . RiskHigh ). // interrupt for approval
StrictDefaults ().
Build ()
if err != nil {
log . Fatal ( err )
}
// Every call is inspected, DLP-scanned, policy-checked, and audited.
result , err := pipeline . Run ( context . Background (), "read_file" , map [ string ] any {
"path" : "/etc/config.yaml" ,
})
if err != nil {
log . Fatal ( err )
}
fmt . Println ( "Result:" , result )
}
Run the complete allow / redact / interrupt example:
go run ./examples/inline-sdk
How It Works
flowchart LR
A[AI Agent] --> S[SentinelMCP]
S --> I[Inspect Tool Call]
I --> D{Policy Decision}
D -->|Allow| E[Execute Tool]
D -->|Redact| R[Redact PII] --> E
D -->|Block| X[Return Error]
D -->|Interrupt| H[Human Approval]
H -->|Approved| E
H -->|Denied| X
E --> O[Inspect Response]
O --> A
E --> MCP[MCP Server]
MCP --> O
Loading
Pipeline flow:
Inspect — Serialize tool arguments, run DLP scanning, look up risk level in policy DB
Policy Decision — Route by risk: low→allow, medium→redact, high→interrupt for approval, or block
Execute — Call the upstream MCP tool (with redacted arguments if applicable)
Inspect Response — DLP-scan the tool response, redact findings, emit structured audit event
sentinelmcp/
├── gateway/ # Core domain (ZERO Eino imports)
│ ├── types.go # GatewayContext, AuditEvent, DLPFinding, RiskLevel, Decision
│ ├── graph.go # Pipeline, ToolInvoker, GatewayConfig, MetricsRecorder
│ ├── policy.go # Policy interface + DefaultPolicy
│ ├── redact.go # DLPScanner, Redactor + RegexDLPScanner + DefaultRedactor
│ ├── dlp_multi.go # MultiDLPScanner (compose regex + external)
│ ├── dlp_external.go # DLPEndpoint interface + ExternalDLPScanner adapter
│ ├── riskdb.go # RiskDB interface + YAMLRiskDB
│ ├── audit.go # AuditEmitter interface + Stdout/FileAuditEmitter
│ ├── audit_sink.go # AuditSink + WriterAuditSink + CompositeAuditEmitter
│ ├── metrics.go # MetricsRecorder interface + NOPMetricsRecorder
│ └── webhook_approval.go # WebhookApprovalProvider + CLIApprovalProvider
├── sdk/ # Inline SDK: ergonomic builder over gateway + adapter
│ ├── builder.go # Builder API (New, With*, StrictDefaults, Build)
│ ├── func_invoker.go # FuncInvoker — secure plain Go functions
│ └── defaults.go # Convenience constructors (BuiltinDLP, StdoutAudit, …)
├── adapter/eino/ # Eino framework adapter (ONLY package with Eino imports)
│ ├── graph.go # BuildGraph() → 3-node graph with interrupt/resume
│ ├── otel_metrics.go # OTelMetricsRecorder (counters + histograms)
│ ├── otel_audit_emitter.go # OTelAuditEmitter (audit → span events)
│ └── bolt_checkpoint.go # BoltCheckPointStore
├── adapter/sidecar/ # Sidecar proxy adapters
│ ├── invoker.go # SidecarInvoker (routes to upstream MCP servers)
│ └── discovery.go # DiscoverTools (upstream MCP server

[truncated]
