---
source: "https://github.com/geludobre/sovereign-agentops"
hn_url: "https://news.ycombinator.com/item?id=48875223"
title: "Sovereign AgentOps – Self-hosted constitutional AI governance for MCP agents"
article_title: "GitHub - geludobre/sovereign-agentops: Sovereign AgentOps — Community Edition: Runtime-agnostic governed agent execution with real Ed25519 crypto · GitHub"
author: "geludobre"
captured_at: "2026-07-11T19:59:40Z"
capture_tool: "hn-digest"
hn_id: 48875223
score: 1
comments: 0
posted_at: "2026-07-11T19:52:10Z"
tags:
  - hacker-news
  - translated
---

# Sovereign AgentOps – Self-hosted constitutional AI governance for MCP agents

- HN: [48875223](https://news.ycombinator.com/item?id=48875223)
- Source: [github.com](https://github.com/geludobre/sovereign-agentops)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T19:52:10Z

## Translation

タイトル: Sovereign AgentOps – MCP エージェント向けの自己ホスト型憲法上の AI ガバナンス
記事のタイトル: GitHub - geludobre/sovereign-agentops: Sovereign AgentOps — Community Edition: 本物の Ed25519 暗号を使用したランタイムに依存しない管理されたエージェントの実行 · GitHub
説明: Sovereign AgentOps — Community Edition: 実際の Ed25519 暗号を使用した、ランタイムに依存しない管理型エージェントの実行 - geludobre/sovereign-agentops

記事本文:
GitHub - geludobre/sovereign-agentops: Sovereign AgentOps — Community Edition: 本物の Ed25519 暗号を使用したランタイムに依存しない管理されたエージェントの実行 · GitHub
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
ゲルドブレ
/
ソブリンエージェントトップ
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
27 コミット 27 コミット .github/ workflows .github/ workflows cli cli docs docs 例 例 テスト テスト ツール ツール .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile LICENSE.community LICENSE.community Makefile Makefile README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自律型デジタル組織プラットフォーム — コミュニティ版
ランタイムに依存しない管理されたエージェントの実行 - デモ版と評価版
# PyPI (スタンドアロン CLI + MCP サーバー)
pip インストール ソブリン-エージェント-コミュニティ
# ドッカー
docker pull geludobre/sovereign-agentops-community
# ソースから
git clone https://github.com/geludobre/sovereign-agentops.git
cd コミュニティ && pip install -e 。
# 完全なドキュメント
# → https://geludobre.github.io/sovereign-agenttops/
これは何ですか？
AI エージェント用のセルフホスト型 MCP ガバナンス サーバー - Community Edition
憲法上の統治、Ed25519 署名の監査証跡、および
EU AI 法に準拠したコンプライアンス ツール。 7 つの MCP ツールを提供します。
ポリシーの適用、Ed25519 署名付きの実行受領書、ワークスペース パスのジェイリング、
ローカル モデル ルーティング — シミュレーションではなく実際の暗号化を使用します。デプロイ
完全にオフライン、エアギャップ、または外部からの影響がない独自のファイアウォールの内側
依存関係。
完全なエンタープライズ プラットフォーム (91 個の MCP ツール、486 個のエンドポイント、88 個の Web UI タブ、
憲法上の統治、サービス カタログ、コンプライアンスの自動化、フリート
コマンド、エージェント所有の資産、デジタル ツイン シミュレーション、および自律インシデント
解像度）私

商用ライセンスの下で利用可能です。
#1. 依存関係をインストールする
pip インストール暗号化 > =41.0.0
# 2. MCP ガバナンス サーバーを実行します (標準入出力上の JSON-RPC)
python3ツール/mcp_server.py
別の端末で、次のコマンドを送信します。
# 利用可能なツールをリストする
echo ' {"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}} ' | python3ツール/mcp_server.py
# コマンドがポリシーによって許可されているかどうかを確認する
echo ' {"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"demo_policy_check","arguments":{"command":"git Push Origin main","path":"/workspace/repo"}}} ' | python3ツール/mcp_server.py
# 本物の Ed25519 で実行受領書に署名する
echo ' {"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"demo_receipt_sign","arguments":{"action":"deploy","target":"web-app"}}} ' | python3ツール/mcp_server.py
ドッカー
docker 構成 --build
サーバーはコンテナ内で stdio モードで実行されます。標準入力経由でコマンドを送信します。
echo ' {"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}} ' \
| docker exec -i Agentops-demo python3 tools/mcp_server.py
ツール
ツール
説明
デモ_ポリシー_チェック
コマンド + パスがポリシーによって許可/ブロックされているかどうかを評価する
デモ_領収書_サイン
Ed25519 署名付き実行受領書を作成する
デモ_受信_検証
領収書の Ed25519 署名を確認する
デモモデルルート
どのローカル LLM エンドポイントがプロンプトを処理するかを表示する
デモ_ワークスペース_ジェイル
パスがワークスペースのルート内に制限されているかどうかを確認する
デモ監査ログ
メモリ内の署名付き監査ログから最近のエントリを返します。
デモサーバー情報
サーバーのメタデータ: バージョン、公開キー、暗号化バックエンド
受領確認
すべての署名済みレシートには、サーバーの Ed25519 公開キーが含まれています。付属のものを使用してください
確認するための CLI ツール:
# サーバーは起動時 (stderr) およびすべての署名済みレシートに公開キーを出力します。
python3 cli/receipt-verify.py --verify pa

th/to/receipt.json
サーバーは最初の実行時に新しい Ed25519 キーペアを生成し、次の場所に保存します。
~/.config/agenttops/ed25519_private.key (権限 0o400 )。
なぜ自律型デジタル組織プラットフォームなのか?
規制された環境で AI コーディング エージェントを導入するチームは、何を制御する必要があるか
エージェントが実行できるコマンド、暗号化された証拠を使用してすべてのアクションを監査し、操作できる
完全にオフライン。このコミュニティ エディションでは、実際の暗号通貨を使用したコンセプトを実証します。
そして実際のポリシーの適用 — エンタープライズ プラットフォームは完全な運用環境を提供します。
積み重ねる。
特徴
コミュニティ
エンタープライズ
ポリシーの適用 (許可/ブロック/検出)
7つのデモツール
完全なプロダクションスタック
Ed25519 署名付き実行受領書
はい
はい
受信確認 CLI
はい
はい
ワークスペースパス刑務所
デモ
完全な施行
モデルルーティングヒューリスティック
デモ
15 以上のプロバイダー ルーティング
署名付き監査ログ
インメモリ
SQLite ハッシュチェーン
MCPツール
7つのデモツール
91 ツール (91 静的 + 0 プラグイン)
憲法上の統治
いいえ
6層（コーテックス、ポリシー、自律性、メモリ、フェデレーション）
ウェブUI
いいえ
91 個以上のタブ (Flask + SocketIO)
PostgreSQL / Redis の永続化
いいえ
はい
エージェントのオーケストレーション
いいえ
機能レジストリ、メッセージバス、DAG
サービスカタログ
いいえ
自動検出、ヘルス、SLA
可観測性ハブ
いいえ
LLM トレース、SLO、コスト分析
コンプライアンスの自動化 (SOC2/HIPAA/PCI)
いいえ
加重スコアリング、PDF レポート、SoD
艦隊司令部
いいえ
マルチインスタンス、ハートビート、条約ゲート型
エージェント所有の資産
いいえ
トークン/データ/コンピューティング/アーティファクト/認証情報
デジタルツインシミュレータ
いいえ
What-if分析、シナリオシミュレーション
自律的なインシデント解決
いいえ
ML診断、自動解決、パターン学習
自主交渉条約
いいえ
6 種類の条約、二重署名
SSO (SAML/OIDC)
いいえ
はい
RBAC カスタムロール
いいえ
はい
高可用性クラスター
いいえ
はい
商用ライセンス
アパッチ2.0
エンタープライズ EULA
ドキュメント
アーキテクチャの概要 — ガバナンス層

デザイン
セキュリティ モデル — 多層防御アーキテクチャ
貢献ガイド — 貢献方法
Community Edition は、Enterprise Edition を備えた Apache 2.0 でのオープンソースです
除外条項。詳細については、LICENSE.community を参照してください。
Enterprise Edition の機能には、FinBridge からの商用ライセンスが必要です。
MCP プロトコルに基づいて構築されています。ロックインではなくガバナンス。
📖 ドキュメントを読む — geludobre.github.io/sovereign-agenttops/
💬 企業向けのお問い合わせ — Enterprise Edition
📦 PyPI — pip インストール sovereign-agentops-community
🐳 Docker Hub — docker pull geludobre/sovereign-agentops-community
Sovereign AgentOps — Community Edition: 本物の Ed25519 暗号を使用した、ランタイムに依存しない管理されたエージェントの実行
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v2.0.0
最新の
2026 年 7 月 9 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Sovereign AgentOps — Community Edition: Runtime-agnostic governed agent execution with real Ed25519 crypto - geludobre/sovereign-agentops

GitHub - geludobre/sovereign-agentops: Sovereign AgentOps — Community Edition: Runtime-agnostic governed agent execution with real Ed25519 crypto · GitHub
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
geludobre
/
sovereign-agentops
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits .github/ workflows .github/ workflows cli cli docs docs examples examples tests tests tools tools .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE.community LICENSE.community Makefile Makefile README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml View all files Repository files navigation
Autonomous Digital Organization Platform — Community Edition
Runtime-agnostic governed agent execution — demo and evaluation edition
# PyPI (standalone CLI + MCP server)
pip install sovereign-agentops-community
# Docker
docker pull geludobre/sovereign-agentops-community
# From source
git clone https://github.com/geludobre/sovereign-agentops.git
cd community && pip install -e .
# Full documentation
# → https://geludobre.github.io/sovereign-agentops/
What is this?
A self-hosted MCP governance server for AI agents — the Community Edition
demonstrates constitutional governance , Ed25519-signed audit trail , and
EU AI Act -aligned compliance tooling. It provides 7 MCP tools that showcase
policy enforcement, Ed25519-signed execution receipts, workspace path jailing,
and local model routing — with real cryptography , not simulations. Deploy
fully offline, air-gapped, or behind your own firewall with zero external
dependencies.
The full Enterprise platform (91 MCP tools, 486 endpoints, 88 web UI tabs,
constitutional governance, service catalogue, compliance automation, fleet
command, agent-owned assets, digital twin simulation, and autonomous incident
resolution) is available under a commercial license.
# 1. Install dependencies
pip install cryptography > =41.0.0
# 2. Run the MCP governance server (JSON-RPC over stdio)
python3 tools/mcp_server.py
In another terminal, send it commands:
# List available tools
echo ' {"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}} ' | python3 tools/mcp_server.py
# Check if a command is allowed by policy
echo ' {"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"demo_policy_check","arguments":{"command":"git push origin main","path":"/workspace/repo"}}} ' | python3 tools/mcp_server.py
# Sign an execution receipt with real Ed25519
echo ' {"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"demo_receipt_sign","arguments":{"action":"deploy","target":"web-app"}}} ' | python3 tools/mcp_server.py
Docker
docker compose up --build
The server runs in stdio mode inside the container. Send commands via stdin:
echo ' {"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}} ' \
| docker exec -i agentops-demo python3 tools/mcp_server.py
Tools
Tool
Description
demo_policy_check
Evaluate whether a command + path is allowed/blocked by policy
demo_receipt_sign
Create an Ed25519-signed execution receipt
demo_receipt_verify
Verify a receipt's Ed25519 signature
demo_model_route
Show which local LLM endpoint would handle a prompt
demo_workspace_jail
Check whether a path is confined within a workspace root
demo_audit_log
Return recent entries from the in-memory signed audit log
demo_server_info
Server metadata: version, public key, crypto backend
Receipt Verification
Every signed receipt includes the server's Ed25519 public key. Use the included
CLI tool to verify:
# The server prints its public key on startup (stderr) and in every signed receipt
python3 cli/receipt-verify.py --verify path/to/receipt.json
The server generates a fresh Ed25519 keypair on first run, stored at
~/.config/agentops/ed25519_private.key (permissions 0o400 ).
Why Autonomous Digital Organization Platform?
Teams adopting AI coding agents in regulated environments need to control what
commands agents can run, audit every action with cryptographic proof, and operate
entirely offline. This community edition demonstrates the concept with real crypto
and real policy enforcement — the Enterprise platform delivers the full production
stack.
Feature
Community
Enterprise
Policy enforcement (allow/block/detect)
7 demo tools
Full production stack
Ed25519-signed execution receipts
Yes
Yes
Receipt verification CLI
Yes
Yes
Workspace path jail
Demo
Full enforcement
Model routing heuristic
Demo
15+ provider routing
Signed audit log
In-memory
SQLite hash-chain
MCP tools
7 demo tools
91 tools (91 static + 0 plugins)
Constitutional governance
No
6-layer (Cortex, Policy, Autonomy, Memory, Federation)
Web UI
No
91+ tabs (Flask + SocketIO)
PostgreSQL / Redis persistence
No
Yes
Agent orchestration
No
Capability registry, message bus, DAG
Service catalogue
No
Auto-discovery, health, SLA
Observability hub
No
LLM tracing, SLOs, cost analytics
Compliance automation (SOC2/HIPAA/PCI)
No
Weighted scoring, PDF reports, SoD
Fleet command
No
Multi-instance, heartbeat, treaty-gated
Agent-owned assets
No
Token/data/compute/artifact/credential
Digital twin simulator
No
What-if analysis, scenario sim
Autonomous incident resolution
No
ML diagnosis, auto-resolve, pattern learning
Self-negotiating treaties
No
6 treaty types, dual-signed
SSO (SAML/OIDC)
No
Yes
RBAC custom roles
No
Yes
High-availability cluster
No
Yes
Commercial license
Apache 2.0
Enterprise EULA
Documentation
Architecture Overview — governance layer design
Security Model — defense-in-depth architecture
Contributing Guide — how to contribute
Community Edition is open source under Apache 2.0 with an Enterprise Edition
exclusion clause. See LICENSE.community for details.
Enterprise Edition features require a commercial license from FinBridge.
Built on the MCP protocol. Governance, not lock-in.
📖 Read the docs — geludobre.github.io/sovereign-agentops/
💬 Enterprise inquiries — Enterprise Edition
📦 PyPI — pip install sovereign-agentops-community
🐳 Docker Hub — docker pull geludobre/sovereign-agentops-community
Sovereign AgentOps — Community Edition: Runtime-agnostic governed agent execution with real Ed25519 crypto
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v2.0.0
Latest
Jul 9, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
