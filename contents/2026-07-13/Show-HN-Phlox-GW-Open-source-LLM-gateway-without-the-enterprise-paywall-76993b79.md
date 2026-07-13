---
source: "https://robert-mcdermott.github.io/phlox-gw/"
hn_url: "https://news.ycombinator.com/item?id=48898189"
title: "Show HN: Phlox-GW – Open-source LLM gateway without the enterprise paywall"
article_title: "Phlox-GW — The Enterprise LLM Gateway Without the Enterprise Paywall"
author: "mcdermott"
captured_at: "2026-07-13T20:50:51Z"
capture_tool: "hn-digest"
hn_id: 48898189
score: 2
comments: 0
posted_at: "2026-07-13T20:16:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Phlox-GW – Open-source LLM gateway without the enterprise paywall

- HN: [48898189](https://news.ycombinator.com/item?id=48898189)
- Source: [robert-mcdermott.github.io](https://robert-mcdermott.github.io/phlox-gw/)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T20:16:14Z

## Translation

タイトル: Show HN: Phlox-GW – エンタープライズ ペイウォールのないオープンソース LLM ゲートウェイ
記事のタイトル: Phlox-GW — エンタープライズ ペイウォールを使用しないエンタープライズ LLM ゲートウェイ
説明: Phlox-GW は、完全にオープンソースのセルフホスト型 LLM ゲートウェイです。OpenAI、Anthropic、Azure、Bedrock、Gemini、および OpenAI 互換プロバイダー向けの SSO、コスト計算、予算、レート制限、ガードレール、クラスタリング、およびプロトコル変換を備えています。 Apache-2.0。永久無料。

記事本文:
Phlox-GW LLM ゲートウェイ
なぜ無料なのか
特長
プロトコル
ガバナンス
クイックスタート
インストール
GitHub
最新の安定リリースが利用可能 · Apache-2.0 · エンタープライズ層なし。これまで。
エンタープライズ LLM ゲートウェイ
エンタープライズペイウォールなしで
Phlox-GW は、ユーザー、アプリ、およびすべての LLM プロバイダー間のセルフホステッド ゲートウェイおよびガバナンス層です。
クラウドでもローカルでも。 SSO、コスト計算、予算執行、ガードレール、監査ログ、クラスタリング —
他のゲートウェイがエンタープライズライセンスの後ろにロックしている機能は、ここでは永久に無料です。
GitHub で見る
単一の Go バイナリ · SQLite または Postgres · macOS / Linux / WSL / Windows
phlox-gw · 管理 → 操作
1 つのゲートウェイ、すべてのプロバイダー、クラウドまたはローカル
OpenAI
人間的
Azure OpenAI
Azure AI Foundry (クロード)
AWS の基盤
Google ジェミニ
オープンルーター
LiteLLM
オラマ
vLLM
LMスタジオ
任意の OpenAI 互換 API
オープンソースの誓約
「オープンソース」にはアスタリスクを付けないでください
ほとんどのオープンソース LLM ゲートウェイは、プロキシとペイウォールを公開してガバナンスを実現します。特徴は、
実際には、責任を持って LLM を実行する必要がある組織が、エンタープライズ ライセンスを必要とすることになります。
Phlox-GW はすべてを Apache-2.0 で出荷します。
これはプロジェクトの原則であり、発売プロモーションではありません。すべての機能は完全に出荷されます。
Apache-2.0 でのオープンソース。エンタープライズ ライセンスの背後にある同等の製品の機能は、ここでは永久に無料のままです。
ユーザーとモデルの間のすべて
モデルルートを公開し、それを使用できるユーザーを制御し、価格を設定し、予算とレート制限を適用します。
1 つの管理コンソールからチャージバックの使用状況をレポートします。
任意の OIDC ID プロバイダーによるブラウザ SSO。最初のログイン時の自動プロビジョニング、チャージバックのための部門請求マッピング、およびグループベースの管理者ロール マッピング。ローカルアカウントも機能します。
追加専用からのユーザーごとおよび部門ごとの毎月のチャージバック

使用量台帳。100 万トークンあたりのモデルごとの米ドル価格が記載されています。財務用の CSV ダウンロード、自動化用の JSON API。
アラートだけでなく、ユーザー、部門、API キーごとの毎月の予算をハードリミットで適用します。バーンダウン ビューには、支出、残りの予算、予測される月末の実行率が表示されます。
ユーザー、部門、プロバイダー、モデル、API キーごとの RPM および TPM 制御。組織全体を抑制することなく、ノイズの多い 1 つのスクリプトを鈍化させます。
chat/default などの安定したルート エイリアス、順序付けされたフォールバック、再試行、試行ごとのタイムアウト、健全性を考慮したルーティング、重み付けされたトラフィック分割 - クライアントにアクセスせずにプロバイダーを交換します。
電子メール、電話番号、SSN、クレジット カード、API キー、カスタム正規表現パターンを検出します。保存前にポリシーをテストするための管理プレビュー ツールを使用して、入出力を編集またはブロックします。
セルフサービスのキーの生成、ローテーション、有効期限切れ。キーごとのモデルの許可リスト、予算、レート制限を備えた管理キー インベントリ。
コスト、トークン、リクエスト、エラー、レイテンシーに関するオペレーション ダッシュボードに加え、既存の監視スタックの Prometheus メトリクスと OpenTelemetry トレース。
すべてのログイン、管理アクション、および API キーのライフサイクル イベントがアクター、ターゲット、詳細、ソース IP とともに記録されます。
デフォルトでは、プロンプト テキスト、応答本文、画像バイト、ツール コンテンツ、シークレットを保存せずに、メタデータ検索と CSV エクスポートをリクエストします。
任意のモデル ルートを通じてテスト チャット メッセージを送信し、プロバイダーとモデルをエンドツーエンドで検証します。API キーは必要ありません。
Ed25519 で署名され、サニタイズされた設定をエクスポートして、レビュー、バックアップ、環境のプロモーションを行うことができます。機密情報や認証情報が箱から出ることはありません。
誰が何を使ったかを正確に把握し、オーバーランを発生前に阻止します
すべてのリクエストは、ユーザー、部門、モデル、トークン数、コストを含む追加専用台帳に格納されます。
財務部門は毎月のチャージバック レポートを受け取ります。アドミ

実際に施行する予算を取得します。
CSV ダウンロードと金融システム統合用の JSON API を使用した、部門およびユーザーごとの毎月のチャージバック。
ハードバジェットの強制 - ユーザー、部門、またはキーが上限に達すると、リクエストは停止します。驚くような請求書はありません。
バーンダウンと予測 - 現在の支出、残りの予算、1 日の平均、スコープごとの予測される月末の実行率。
ローカル モデルでも正確です。プロバイダーが使用状況データを省略しても、トークンはメモリ内で推定されるため、コストが黙ってゼロになることはありません。
クラスタに成長する単一のバイナリ
1 つのプロセスと SQLite から開始します。外部依存関係はありません。 1 つのノードでは不十分な場合は、
複数のゲートウェイをロード バランサーの背後にある共有 Postgres に向けて続行します。
単一バイナリのシンプルさ — ダッシュボード、API、ゲートウェイは 1 つの Go 実行可能ファイルに同梱されています。 SQLite がデフォルトです。
明示的クラスター モード — 移行ロック、ノード ハートビート、および準備状況チェックを備えたクラスター postgres。
クラスター管理 UI — すべてのノードのステータス、ハートビート、データベース ターゲットを一目で確認できます。
ロードバランサに優しい — 稼働状態には /health 、データベースとハートビートの準備には /ready です。
Ops-ready — systemd、macOS、および Windows のガイダンス、リバース プロキシ パターン、および単一ホストのデモ クラスター スクリプト。
どのクライアントでも。任意のプロバイダー。フロックス-GWが翻訳します。
OpenAI プロトコル クライアントは、Bedrock 上の Claude を呼び出すことができます。 Anthropic-protocol クライアント — クロード コードを含む —
Azure 上の GPT または Ollama 上のローカル モデルを呼び出すことができます。ストリーミング、ツール呼び出し、画像が含まれます。
6 つのパスはすべて、ツール呼び出しとイメージ マッピングを使用して、ストリーミングおよび非ストリーミング リクエストをサポートします。
正確なコスト計算のための使用状況のキャプチャ。推論モデルの癖 (max_completion_tokens など)
拒否されたサンプリング パラメータなど）が自動的に検出されて再試行されるため、古いクライアントは引き続き動作します

新しいモデルを備えたキング。
1 つの ANTHROPIC_BASE_URL を使用してゲートウェイで Claude Code をポイントし、それを Claude、GPT、Bedrock、またはローカル モデルにルーティングします。ストリーミングとツールの使用は完全に変換され、anthropic-beta ヘッダーはプロバイダーごとにフィルターされるため、Azure AI Foundry のような厳密なエンドポイントは要求を拒否しません。
プロバイダーごとに、SDK 認証情報チェーン (SSO プロファイル、インスタンスおよびタスクロール)、保存されたアクセスキー、または Bedrock API キーを使用して AWS Bedrock に対して認証します。管理するベース URL はありません。地域を選択するだけです。
OpenAI または Anthropic SDK は、base_url と API キーを変更することで機能します。モデル名は安定したルート ID であるため、プロバイダーを交換してもアプリケーション コードには触れません。
ガードレール、制限、および紙の軌跡
ポリシーはリクエスト パスで適用されますが、Wiki には文書化されておらず、期待されています。
電子メール アドレス、電話番号、SSN、クレジット カード、API キーの組み込み検出器 —
さらに、組織の識別子 (MRN、従業員 ID、内部トークン) 用のカスタム RE2 正規表現パターンも含まれます。
プロバイダーに何かが送信される前に、入力を編集またはブロックします。
応答がクライアントに届く前に出力を編集またはブロックします。
コミットする前にプレビューする — 管理パネルで直接、サンプル テキストに対してドラフト パターンをテストします。
メモリ内で検査 — コンテンツは実行中にスキャンされ、保存されることはありません。
組織図に一致するレート制限
必要に応じて、1 分あたりのリクエスト数と 1 分あたりのトークン数の上限を設定します。
ユーザー、部門、プロバイダー、モデルルート、または個別の API キー。
共有容量を保護する — 部門による高価なフロンティア モデルの使用を制限します。
暴走を阻止する - 1 つの不正なスクリプトが全員のプロバイダー クォータを使い果たすことはできません。
サーキット ブレーカー — 障害が発生し続けるプロバイダーは、回復するまで自動的にローテーションから外されます。
バイナリが 1 つ。ラウ前に検証済み

インチ。
macOS、Linux、または WSL でチェックサム検証シェル インストーラーを使用するか、ネイティブ Windows で PowerShell を使用してインストールします。ダッシュボード、API、ゲートウェイがバンドルされています。
インストーラーは、オペレーティング システムとアーキテクチャに適したリリースを選択し、そのチェックサムを検証して、 phlox-gw としてインストールします。
最新のネイティブ実行可能ファイルをダウンロードし、公開されている SHA-256 マニフェストと照合して検証し、現在のディレクトリから起動します。この例は x86-64 をターゲットとしています。 Windows ARM64 では、アセット名の amd64 を arm64 に変更します。
設計により透過的: どちらのインストール パスでも、ダウンロードされたバイナリがリリース チェックサム マニフェストと照合して検証されます。シェル インストーラーは実行前に読み取ることができ、PowerShell の手順は上に完全に示されています。リリース バイナリはまだコード署名または公証されていません。
インストールからLLMトラフィックの管理まで数分で完了
Kubernetes、サイドカー フリート、ライセンス キーはありません。バイナリ 1 つとブラウザ 1 つ。
# macOS、Linux、または WSL
$カール --proto '=https' --tlsv1.2 -fsSL \
https://raw.githubusercontent.com/robert-mcdermott/phlox-gw/main/install.sh |しー
# 専用のローカルデータディレクトリで実行
$ mkdir -p "$HOME/.local/share/phlox-gw"
$ cd "$HOME/.local/share/phlox-gw"
$ "$HOME/.local/bin/phlox-gw"
# → http://127.0.0.1:8080 でサインインします
# OpenAI のように呼び出します…
$カール http://127.0.0.1:8080 /v1/chat/completions \
-H "認可: Bearer pgw-sk-your-key" \
-d '{"モデル":"チャット/デフォルト","メッセージ":
[{"役割":"ユーザー","コンテンツ":"こんにちは"}]}'
# …または Anthropic のようなもの — 同じゲートウェイ、同じキー
$カール http://127.0.0.1:8080 /anthropic/v1/messages \
-H "x-api-key: pgw-sk-your-key" \
-H "人間バージョン: 2023-06-01" \
-d '{"モデル":"チャット/デフォルト","最大トークン":256,
"メッセージ":[{"ロール":"ユーザー","コンテンツ":"Hello"}]}'
インストールして実行する macOS、Linux でチェックサム検証シェル インストーラーを使用する

、WSL、またはネイティブ Windows 上の PowerShell ワークフロー。 SQLite はインフラストラクチャなしで開始できることを意味します。
プロバイダーとモデルを追加する OpenAI、Anthropic、Azure、Bedrock、Gemini、またはローカルの Ollama/vLLM を指定します。キーは環境変数内に存在できます。保存されたシークレットは書き込み専用であり、ブラウザーに返されることはありません。
価格、予算、モデルごとの 100 万トークンあたりの米ドルの制限、ユーザー/部門/キーごとの月間上限、任意のスコープでの RPM および TPM 制限を設定します。
プレイグラウンドで検証する 単一のクライアントが接続する前に、管理 UI からすべてのルートをテストします。API キーは必要ありません。
アプリをそこにポイントします。 OpenAI または Anthropic SDK のベース URL を交換します。既存のコードは引き続き動作します。ガバナンスがオンになります。
今夜は LLM トラフィックを管理しましょう。
予算とデータを維持しましょう。
自己ホスト型で、デフォルトでプライベートであり、細字部分なしで無料です。それが製品に含まれている場合、それはあなたのものです。

## Original Extract

Phlox-GW is a fully open-source, self-hosted LLM gateway: SSO, cost accounting, budgets, rate limits, guardrails, clustering, and protocol translation for OpenAI, Anthropic, Azure, Bedrock, Gemini, and any OpenAI-compatible provider. Apache-2.0. Free forever.

Phlox-GW LLM Gateway
Why Free
Features
Protocols
Governance
Quick Start
Install
GitHub
Latest stable release available · Apache-2.0 · No enterprise tier. Ever.
The enterprise LLM gateway
without the enterprise paywall
Phlox-GW is a self-hosted gateway and governance layer between your users, apps, and every LLM provider —
cloud or local. SSO, cost accounting, budget enforcement, guardrails, audit logs, and clustering —
the features other gateways lock behind enterprise licenses are free here, forever.
View on GitHub
single Go binary · SQLite or Postgres · macOS / Linux / WSL / Windows
phlox-gw · Admin → Operations
One gateway · every provider · cloud or local
OpenAI
Anthropic
Azure OpenAI
Azure AI Foundry (Claude)
AWS Bedrock
Google Gemini
OpenRouter
LiteLLM
Ollama
vLLM
LM Studio
any OpenAI-compatible API
The open-source pledge
“Open source” shouldn’t come with an asterisk
Most open-source LLM gateways publish the proxy and paywall the governance. The features an
organization actually needs to run LLMs responsibly end up behind an enterprise license.
Phlox-GW ships everything under Apache-2.0.
This is a project principle, not a launch promo: every feature ships fully
open source under Apache-2.0. Capabilities that comparable products gate behind enterprise licenses stay free here — permanently.
Everything between your users and your models
Publish model routes, control who can use them, attach prices, enforce budgets and rate limits,
and report usage for chargeback — from one admin console.
Browser SSO with any OIDC identity provider. Auto-provisioning on first login, department claim mapping for chargeback, and group-based admin role mapping. Local accounts work too.
Per-user and per-department monthly chargeback from an append-only usage ledger, with per-model USD pricing per million tokens. CSV download for finance, JSON API for automation.
Monthly budgets per user, department, and API key with hard-limit enforcement — not just alerts. Burn-down views show spend, remaining budget, and projected month-end run rate.
RPM and TPM controls by user, department, provider, model, and API key. Blunt one noisy script without throttling the whole org.
Stable route aliases like chat/default , ordered fallbacks, retries, per-attempt timeouts, health-aware routing, and weighted traffic splits — swap providers without touching clients.
Detect emails, phone numbers, SSNs, credit cards, API keys, and custom regex patterns. Redact or block on input and output, with an admin preview tool for testing policies before saving.
Self-service key minting, rotation, and expiration. Admin key inventory with per-key model allowlists, budgets, and rate limits.
Operations dashboards for cost, tokens, requests, errors, and latency, plus Prometheus metrics and OpenTelemetry traces for your existing monitoring stack.
Every login, admin action, and API-key lifecycle event recorded with actor, target, details, and source IP.
Request metadata search and CSV export without storing prompt text, response bodies, image bytes, tool contents, or secrets — by default.
Send test chat messages through any model route to validate providers and models end to end — no API key required.
Ed25519-signed, sanitized configuration export for review, backup, and environment promotion — secrets and credentials never leave the box.
Know exactly who spent what — and stop overruns before they happen
Every request lands in an append-only ledger with user, department, model, token counts, and cost.
Finance gets a monthly chargeback report; admins get budgets that actually enforce.
Monthly chargeback by department and user, with CSV download and a JSON API for financial-system integration.
Hard budget enforcement — when a user, department, or key hits its cap, requests stop. No surprise invoices.
Burn-down & projections — current spend, remaining budget, daily average, and projected month-end run rate per scope.
Accurate even for local models — when a provider omits usage data, tokens are estimated in memory so cost never silently drops to zero.
A single binary that grows into a cluster
Start with one process and SQLite — zero external dependencies. When one node isn’t enough,
point multiple gateways at a shared Postgres behind a load balancer and keep going.
Single-binary simplicity — the dashboard, API, and gateway ship in one Go executable. SQLite is the default.
Explicit cluster mode — cluster-postgres with migration locking, node heartbeats, and readiness checks.
Cluster admin UI — see every node’s status, heartbeat, and database target at a glance.
Load-balancer friendly — /health for liveness and /ready for database-and-heartbeat readiness.
Ops-ready — systemd, macOS, and Windows guidance, reverse-proxy patterns, and a single-host demo cluster script.
Any client. Any provider. Phlox-GW translates.
OpenAI-protocol clients can call Claude on Bedrock. Anthropic-protocol clients — including Claude Code —
can call GPT on Azure or a local model on Ollama. Streaming, tool calls, and images included.
All six paths support streaming and non-streaming requests, with tool-call and image mapping,
and usage capture for accurate cost accounting. Reasoning-model quirks (like max_completion_tokens
and rejected sampling parameters) are detected and retried automatically, so old clients keep working with new models.
Point Claude Code at the gateway with one ANTHROPIC_BASE_URL — and route it to Claude, GPT, Bedrock, or a local model. Streaming and tool use are fully translated, and anthropic-beta headers are filtered per provider so strict endpoints like Azure AI Foundry don't reject requests.
Authenticate to AWS Bedrock with the SDK credential chain (SSO profiles, instance and task roles), stored access keys, or a Bedrock API key — per provider. No base URL to manage; just pick a region.
Any OpenAI or Anthropic SDK works by changing base_url and the API key. Model names are your stable route IDs, so swapping providers never touches application code.
Guardrails, limits, and a paper trail
Policy is enforced in the request path — not documented in a wiki and hoped for.
Built-in detectors for email addresses, phone numbers, SSNs, credit cards, and API keys —
plus custom RE2 regex patterns for your organization’s identifiers (MRNs, employee IDs, internal tokens).
Redact or block on input before anything is sent to a provider.
Redact or block on output before a response reaches the client.
Preview before you commit — test draft patterns against sample text right in the admin panel.
Inspected in memory — content is scanned in-flight and never stored.
Rate limits that match your org chart
Set requests-per-minute and tokens-per-minute ceilings wherever they make sense:
a user, a department, a provider, a model route, or an individual API key.
Protect shared capacity — cap a department’s use of an expensive frontier model.
Contain runaways — one misbehaving script can’t exhaust the provider quota for everyone.
Circuit breakers — providers that keep failing are automatically taken out of rotation until they recover.
One binary. Verified before launch.
Use the checksum-verifying shell installer on macOS, Linux, or WSL, or install with PowerShell on native Windows. The dashboard, API, and gateway are bundled together.
The installer selects the correct release for your operating system and architecture, verifies its checksum, and installs it as phlox-gw .
Download the latest native executable, verify it against the published SHA-256 manifest, and launch it from your current directory. This example targets x86-64; on Windows ARM64, change amd64 to arm64 in the asset name.
Transparent by design: both installation paths verify the downloaded binary against the release checksum manifest. The shell installer is readable before execution, and the PowerShell steps are fully shown above. Release binaries are not yet code-signed or notarized.
From install to governed LLM traffic in minutes
No Kubernetes, no sidecar fleet, no license key. One binary and a browser.
# macOS, Linux, or WSL
$ curl --proto '=https' --tlsv1.2 -fsSL \
https://raw.githubusercontent.com/robert-mcdermott/phlox-gw/main/install.sh | sh
# run in a dedicated local data directory
$ mkdir -p "$HOME/.local/share/phlox-gw"
$ cd "$HOME/.local/share/phlox-gw"
$ "$HOME/.local/bin/phlox-gw"
# → sign in at http://127.0.0.1:8080
# call it like OpenAI…
$ curl http://127.0.0.1:8080 /v1/chat/completions \
-H "Authorization: Bearer pgw-sk-your-key" \
-d '{"model":"chat/default","messages":
[{"role":"user","content":"Hello"}]}'
# …or like Anthropic — same gateway, same key
$ curl http://127.0.0.1:8080 /anthropic/v1/messages \
-H "x-api-key: pgw-sk-your-key" \
-H "anthropic-version: 2023-06-01" \
-d '{"model":"chat/default","max_tokens":256,
"messages":[{"role":"user","content":"Hello"}]}'
Install and run Use the checksum-verifying shell installer on macOS, Linux, or WSL, or the PowerShell workflow on native Windows. SQLite means zero infrastructure to start.
Add providers & models Point at OpenAI, Anthropic, Azure, Bedrock, Gemini, or your local Ollama/vLLM. Keys can live in env vars — stored secrets are write-only and never returned to the browser.
Set prices, budgets & limits USD per million tokens per model, monthly caps per user/department/key, RPM & TPM limits at any scope.
Validate in the playground Test every route from the admin UI before a single client connects — no API key needed.
Point your apps at it Swap the base URL in any OpenAI or Anthropic SDK. Existing code keeps working; governance turns on.
Govern your LLM traffic tonight.
Keep your budget — and your data.
Self-hosted, private by default, and free without fine print. If it’s in the product, it’s yours.
