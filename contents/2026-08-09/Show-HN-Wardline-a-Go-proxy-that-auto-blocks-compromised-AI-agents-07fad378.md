---
source: "https://github.com/kabirnarang39/wardline"
hn_url: "https://news.ycombinator.com/item?id=49235863"
title: "Show HN: Wardline, a Go proxy that auto-blocks compromised AI agents"
article_title: "GitHub - kabirnarang39/wardline: Open source control-plane proxy for AI agents: identity, policy, budget, audit for MCP and beyond · GitHub"
author: "kabirnarang39"
captured_at: "2026-08-09T21:22:07Z"
capture_tool: "hn-digest"
hn_id: 49235863
score: 2
comments: 0
posted_at: "2026-08-09T21:03:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Wardline, a Go proxy that auto-blocks compromised AI agents

- HN: [49235863](https://news.ycombinator.com/item?id=49235863)
- Source: [github.com](https://github.com/kabirnarang39/wardline)
- Score: 2
- Comments: 0
- Posted: 2026-08-09T21:03:50Z

## Translation

タイトル: Show HN: Wardline、侵害された AI エージェントを自動ブロックする Go プロキシ
記事のタイトル: GitHub - kabirnarang39/wardline: AI エージェント用のオープンソース コントロール プレーン プロキシ: MCP 以降の ID、ポリシー、予算、監査 · GitHub
説明: AI エージェント用のオープンソース コントロール プレーン プロキシ: MCP などの ID、ポリシー、予算、監査 - kabirnarang39/wardline

記事本文:
GitHub - kabirnarang39/wardline: AI エージェント用のオープンソース コントロール プレーン プロキシ: MCP などの ID、ポリシー、予算、監査 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
カビルナラン39
/
ワードライン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
573 コミット 573 コミット .github .github charts/ warline charts/ warline cmd/ warline cmd/ warline デモ デモ docs-site docs-site docs docs 内部 内部

サイト サイト .dockerignore .dockerignore .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md credentials.yaml.example credentials.yaml.example go.mod go.mod go.sum go.sum ポリシー.cedar.example ポリシー.cedar.example ポリシー.rego.example ポリシー.rego.example ポリシー.yaml.example ポリシー.yaml.example rbac.yaml.example rbac.yaml.example warline.yaml.example warline.yaml.example すべてのファイルを表示 リポジトリ ファイルのナビゲーション
侵害された AI エージェントを自動ブロックするコントロール プレーン プロキシ — 1 つの静的 Go バイナリ内。
Wardline は、AI エージェントと、AI エージェントが呼び出すすべてのもの (MCP サーバー、ツール、gRPC アップストリーム) の間に位置し、アイデンティティ、ポリシー、予算、監査を強制するオープンソース プロキシです。統計的異常検出により、侵害されたエージェントをリアルタイムでブロックします。攻撃用のルールは作成されず、人間も関与しません。 1 つの静的 Go バイナリ。起動するデータベース、IdP、またはサイドカーがありません。
make Demon # 模擬 MCP サーバー + Wardline を起動し、上記のシナリオを実行します
組み込みの読み取り専用ダッシュボードでも同じことが実行されます。ブロック、そのブロックを引き起こした異常、その背後にあるポリシーが表示されます。
AI エージェント、CLI/IDE、アプリなどの呼び出し元は、Wardline 経由でのみ MCP/gRPC アップストリームに到達します。Wardline は、プロセス内で ID、ポリシー、予算、異常検出を適用し、あらゆる決定を監査証跡に書き込みます。
リアルタイム異常自動ブロック
アラートだけではない 4 つの自己ベースライン ヒューリスティック (レート スパイク、新しいツール、拒否レート スパイク、およびウェルフォード アルゴリズムによる ml_score Z スコアの組み合わせ - トレーニング データなし、外部モデルなし): auto_block は、フラグが設定された ID の呼び出しを拒否します。

s は制限付き TTL の場合。ログ行ではなく強制です。
3 つのポリシー バックエンド、1 つのバイナリ
静的 YAML、組み込み OPA/Rego、組み込み AWS Cedar — 外部プロセスやネットワーク ホップなしで、単一のpolicy_backend 設定キーによって切り替えられます。
アイデンティティとアクセス
リフレッシュ トークンと JWKS ローテーションを使用した有効期間の短い RS256 JWT 発行、OIDC / mTLS-SPIFFE ブートストラップ、Kubernetes スタイルの RBAC、SCIM 2.0 プロビジョニング、およびエンドツーエンドのテナント分離。
予算と料金の管理
ID ごとおよびテナントごとの 2 段階のレート制限。通話を続行するには両方をクリアする必要があります。
コンプライアンスと監査
構造化された JSON 監査証跡、ワードラインのエクスポート証拠 (チェックサム、監査人用の RSA 署名可能なバンドル)、構成可能な保持、および観察されたトラフィックからスターター ホワイトリストを生成するワードラインの推論ポリシー。
フェデレーションと可観測性
署名済み、仮名化された異常概要のインスタンス間相関。 OpenTelemetry トレース。ライブ Web ダッシュボード。 Postgres 上で状態を共有する HA マルチレプリカ展開。
# ソースから (常に動作)
go build -o warline ./cmd/wardline
# または、公開されたマルチ アーキテクチャ イメージ (タグ付けされたリリースごとに構築) をプルします
docker pull ghcr.io/kabirnarang39/wardline:latest
./wardline validate-policy --filepolicy.yaml.example
./wardline validate-config --config warline.yaml.example
./wardlineserve --config warline.yaml.example
実際の MCP サーバーの上流を指します (簡単なテストのために、 python3 -m http.server 9000 を実行するまで、プロキシされた呼び出し 502s)。すべてのリクエストには X-Wardline-Identity ヘッダーが含まれます。ポリシーは、その値と MCP ツール名に基づいて一致します。
カール -X POST http://localhost:8080 \
-H " X-Wardline-Identity: エージェント-abc123 " \
-H " Content-Type: application/json " \
-d ' {"jsonrpc":"2.0","method":"tools/call","params":{"name":"read_file"}} '
事前構築済みバイナリ (linux/darwin/windows)

· amd64/arm64) およびマルチ アーキテクチャ イメージは、リリースおよび GHCR を介してすべての v* タグで出荷されます。
完全なドキュメント、機能ごとの設計ノート、正直な既知の制限事項はドキュメント サイトにあります。
はじめに — インストール、クイックスタート、構成
概念 — アーキテクチャ、ポリシー バックエンド、アイデンティティ、監査
機能 — すべての機能の詳細
デプロイ - Docker、Helm、HA、可観測性
フレームワークの統合 — LangChain、LlamaIndex、OpenAI Agents SDK、CrewAI、生の MCP
以下のすべては、internal/features/ で出荷され、テスト可能です。 v0.1 ベースライン (プロキシ + ポリシー + 監査) は常にオンです。それ以外はすべて設定フラグによって制御されます。
マーケティング数値ではなく、 go test -bench で再現可能です。 BenchmarkDecider_Decide (デフォルトの YAML バックエンド、Apple Silicon): 10 ルールで ~33 ns / 0 割り当て、1000 ルールで ~2.4 µs。 ml_score の誤検知の主張は、TestDetector_MLScore_FalsePositiveRateOnSteadyTraffic によって回帰保護されています (定常トラフィック、予算 < 2% での誤検知 0% をアサートします)。
ダッシュボードと X-Wardline-Identity ヘッダーはデフォルトでは認証されていません。実際のセキュリティ値を得るには、credential_issuance や rbac と組み合わせます。すべてのオプション機能はデフォルトで出荷され、フェールクローズされます。 Wardline は起動時に、まだ有効な安全でないデフォルトごとに警告をログに記録するため、姿勢が沈黙することはありません。 SECURITY.md に従って脆弱性を報告してください。
すぐに使用すると、プロキシはポリシーで失敗して閉じられますが、ID とダッシュボードは開いています。実際の展開では、以下をオンにします。
特徴：
credential_issuance : true # X-Wardline-Identity を信頼する代わりに署名付きベアラー トークンを検証します
rbac : true # 実際の権限に基づいてダッシュボードと管理者アクションをゲートします
credential_issuance がオンの場合、スプーフィング可能なヘッダーは RS256 ベアラー トークン検証によって置き換えられます。 rbac がオンの場合、ダッシュボードの読み取りビューとミューテーションが再表示されます。

承認された身元を要求します。異常検出は突然の不正行為を検出しますが、緩やかな変化は検出しません (既知の制限を参照)。したがって、ハードフロアとして明示的なポリシー + 予算制限を維持してください。
ワードラインは若く、動きが速い。すべての機能のドキュメント ページでは、機能が何をするのか、何をしないのかについて意図的に率直に記載されています。特に異常検出アプローチと脅威モデルに関するフィードバック、問題点、貢献を歓迎します。
CONTRIBUTING.md および CODE_OF_CONDUCT.md を参照してください。アーキテクチャとエンジニアリングの規則は CLAUDE.md に文書化されています。ロードマップはドキュメントに記載されています。
AI エージェント用のオープンソース コントロール プレーン プロキシ: ID、ポリシー、予算、MCP などの監査
kabirnarang39.github.io/wardline/ トピック
Readme Apache-2.0 ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open source control-plane proxy for AI agents: identity, policy, budget, audit for MCP and beyond - kabirnarang39/wardline

GitHub - kabirnarang39/wardline: Open source control-plane proxy for AI agents: identity, policy, budget, audit for MCP and beyond · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
kabirnarang39
/
wardline
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
573 Commits 573 Commits .github .github charts/ wardline charts/ wardline cmd/ wardline cmd/ wardline demo demo docs-site docs-site docs docs internal internal site site .dockerignore .dockerignore .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md credentials.yaml.example credentials.yaml.example go.mod go.mod go.sum go.sum policy.cedar.example policy.cedar.example policy.rego.example policy.rego.example policy.yaml.example policy.yaml.example rbac.yaml.example rbac.yaml.example wardline.yaml.example wardline.yaml.example View all files Repository files navigation
The control-plane proxy that auto-blocks compromised AI agents — in one static Go binary.
Wardline is an open-source proxy that sits between your AI agents and everything they call (MCP servers, tools, gRPC upstreams) and enforces identity, policy, budget, and audit — with statistical anomaly detection that blocks a compromised agent in real time , no rule written for the attack and no human in the loop. One static Go binary; no database, IdP, or sidecar to start.
make demo # spins up a mock MCP server + Wardline and runs the scenario above
The same run in the built-in read-only dashboard — the block, the anomaly that triggered it, and the policy behind it:
Any caller — an AI agent, a CLI/IDE, or an app — reaches its MCP/gRPC upstreams only through Wardline, which applies identity, policy, budget, and anomaly detection in-process and writes every decision to the audit trail.
Real-time anomaly auto-block
Four self-baselining heuristics (rate spike, novel tool, deny-rate spike, and a combined ml_score z-score via Welford's algorithm — no training data, no external model) that don't just alert: auto_block rejects a flagged identity's calls for a bounded TTL. Enforcement, not a log line.
Three policy backends, one binary
Static YAML, embedded OPA/Rego, and embedded AWS Cedar — switched by a single policy_backend config key, with no external process and no network hop.
Identity & access
Short-lived RS256 JWT issuance with refresh tokens and JWKS rotation, OIDC / mTLS-SPIFFE bootstrap, Kubernetes-style RBAC, SCIM 2.0 provisioning, and end-to-end tenant isolation.
Budget & rate control
Two-tier per-identity and per-tenant rate limits — both must clear for a call to proceed.
Compliance & audit
Structured JSON audit trail, wardline export-evidence (checksummed, RSA-signable bundle for an auditor), configurable retention, and wardline infer-policy to generate a starter allow-list from observed traffic.
Federation & observability
Cross-instance correlation over signed, pseudonymized anomaly summaries; OpenTelemetry tracing; a live web dashboard; and HA multi-replica deployment with shared state over Postgres.
# From source (always works)
go build -o wardline ./cmd/wardline
# Or pull the published multi-arch image (built for each tagged release)
docker pull ghcr.io/kabirnarang39/wardline:latest
./wardline validate-policy --file policy.yaml.example
./wardline validate-config --config wardline.yaml.example
./wardline serve --config wardline.yaml.example
Point upstream at a real MCP server (a proxied call 502s until you do — for a quick test, python3 -m http.server 9000 ). Every request carries an X-Wardline-Identity header; policy matches on that value plus the MCP tool name:
curl -X POST http://localhost:8080 \
-H " X-Wardline-Identity: agent-abc123 " \
-H " Content-Type: application/json " \
-d ' {"jsonrpc":"2.0","method":"tools/call","params":{"name":"read_file"}} '
Prebuilt binaries (linux/darwin/windows · amd64/arm64) and multi-arch images ship on every v* tag via Releases and GHCR .
Full docs, per-feature design notes, and honest known-limitations live on the docs site:
Getting Started — install, quickstart, configuration
Concepts — architecture, policy backends, identity, audit
Features — every capability in depth
Deployment — Docker, Helm, HA, observability
Framework integrations — LangChain, LlamaIndex, OpenAI Agents SDK, CrewAI, raw MCP
Everything below is shipped and testable under internal/features/ . The v0.1 baseline (proxy + policy + audit) is always on; everything else is gated by a config flag.
Reproducible with go test -bench , not marketing numbers. BenchmarkDecider_Decide (default YAML backend, Apple Silicon): ~33 ns / 0 allocations at 10 rules, ~2.4 µs at 1000 rules. The ml_score false-positive claim is regression-guarded by TestDetector_MLScore_FalsePositiveRateOnSteadyTraffic (asserts 0% false positives on steady traffic, budget < 2%).
The dashboard and the X-Wardline-Identity header are unauthenticated by default — pair with credential_issuance and/or rbac for real security value. Every optional capability ships off by default and fails closed. On startup Wardline logs a WARN for each insecure default still in effect, so the posture is never silent. Report vulnerabilities per SECURITY.md .
Out of the box the proxy fails closed on policy , but identity and the dashboard are open. For any real deployment, turn on:
features :
credential_issuance : true # verify a signed bearer token instead of trusting X-Wardline-Identity
rbac : true # gate the dashboard and admin actions on real permissions
With credential_issuance on, the spoofable header is replaced by RS256 bearer-token verification; with rbac on, dashboard read views and mutations require an authorized identity. Anomaly detection catches abrupt abuse but not low-and-slow ramps (see its known limitations ), so keep explicit policy + budget limits as the hard floor.
Wardline is young and moving fast. Every feature's docs page is deliberately blunt about what it does and doesn't do. Feedback, issues, and contributions are welcome — especially on the anomaly-detection approach and threat model.
See CONTRIBUTING.md and CODE_OF_CONDUCT.md . Architecture and engineering conventions are documented in CLAUDE.md ; the roadmap lives in the docs .
Open source control-plane proxy for AI agents: identity, policy, budget, audit for MCP and beyond
kabirnarang39.github.io/wardline/ Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
