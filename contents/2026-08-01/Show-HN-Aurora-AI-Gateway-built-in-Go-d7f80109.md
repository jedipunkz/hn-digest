---
source: "https://github.com/aurorallm/aurora"
hn_url: "https://news.ycombinator.com/item?id=49134502"
title: "Show HN: Aurora – AI Gateway built in Go"
article_title: "GitHub - aurorallm/aurora: The Fastest enterprise AI gateway — route LLM traffic across OpenAI, Anthropic, Gemini, Groq & 30+ providers via a single API. Self-hosted, no vendor lock-in. ( 55x faster than litellm ) · GitHub"
author: "gurveer51"
captured_at: "2026-08-01T14:25:25Z"
capture_tool: "hn-digest"
hn_id: 49134502
score: 4
comments: 0
posted_at: "2026-08-01T13:56:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Aurora – AI Gateway built in Go

- HN: [49134502](https://news.ycombinator.com/item?id=49134502)
- Source: [github.com](https://github.com/aurorallm/aurora)
- Score: 4
- Comments: 0
- Posted: 2026-08-01T13:56:31Z

## Translation

タイトル: 表示 HN: Aurora – Go に組み込まれた AI ゲートウェイ
記事のタイトル: GitHub - aurorallm/aurora: 最速のエンタープライズ AI ゲートウェイ — 単一の API を介して、OpenAI、Anthropic、Gemini、Groq、および 30 以上のプロバイダー間で LLM トラフィックをルーティングします。自己ホスト型でベンダーロックインなし。 ( litellm より 55 倍高速) · GitHub
説明: 最速のエンタープライズ AI ゲートウェイ — 単一の API を介して、OpenAI、Anthropic、Gemini、Groq、および 30 以上のプロバイダー間で LLM トラフィックをルーティングします。自己ホスト型でベンダーロックインなし。 ( litellm より 55 倍高速) - aurorallm/aurora
HN テキスト: すべての Ai プロジェクトには、同じインフラストラクチャ (API キー、プロバイダー SDK、再試行ロジック、コスト追跡) が必要です。インフラ層である必要があります。
Aurora は、そのレイヤーを処理する Go ゲートウェイであり、バイナリは最大 15 MB で、依存関係はありません。 OpenaAI、Anthropic、Gemini などへのリクエストをプロキシし、形式間の変換が可能です。OpenAi SDK を使用すると、Aurora が自動的に Anthropic のワイヤ形式に変換します。バックエンドに形式変換コードがありません。役に立ちます:
- 顧客/チームごとのコストを追跡する
- コードを変更せずにモデルとプロバイダーを切り替えます
- リクエストとレスポンスのフローをより簡単にデバッグできます
- 正確かつセマンティックなキャッシュによりコストを削減します 他のゲートウェイとの違い:
- プロバイダーモデルを自動検出します
- 自動フェイルオーバースイッチを備えたプロバイダープール
- 組み込みのガーディアル、エクスポートによる監査ログ、分析、管理ダッシュボード
- モデルのエイリアス、地域展開によるレイテンシーの削減。
- インフラ内で直接使用でき、管理エンドポイントですべてを制御し、ユーザーオンボードで管理 API キーを作成するなどのことができます。 OSS バージョン (Apache 2.0) にはすべてが含まれています。 Enterprise では、SSO、RBAC、テナント分離が追加されます。 npm install -g iaurora Web サイト: https://aurorallm.online フィードバックは大歓迎です。

記事本文:
GitHub - aurorallm/aurora: 最速のエンタープライズ AI ゲートウェイ — 単一の API を介して、OpenAI、Anthropic、Gemini、Groq、および 30 以上のプロバイダー間で LLM トラフィックをルーティングします。自己ホスト型でベンダーロックインなし。 ( litellm より 55 倍高速) · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
を却下する

警告します
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
オーロラルム
/
オーロラ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .github/ workflows .github/ workflows apps/ aurora apps/ aurora configs configs 構成 構成 ダッシュボード UI ダッシュボード UI docs-assets docs-assets docs docs ハートビートワーカー ハートビートワーカー Helm Helm 内部 内部監視 モニタリング .dockerignore .dockerignore .goreleaser.yaml .goreleaser.yaml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md artifacthub-repo.yml artifacthub-repo.yml docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Aurora - 最速の AI ゲートウェイ
AI プロバイダーごとに 1 つの API。自己ホスト型。ベンダーロックインはありません。
14 のプロバイダー タイプ • OpenAI および Anthropic との互換性 • Go • Apache 2.0 • 生の速度を重視して構築
Aurora はアプリと LLM プロバイダーの間に位置します。アプリは標準の OpenAI または Anthropic SDK を使用してリクエストを送信します。Aurora は、設定したプロバイダーにリクエストをルーティングします。 1 つの形式ですべてを処理できるため、プロバイダー固有の形式について心配する必要はありません。
# 以前: ハードコードされたプロバイダー
client = OpenAI (base_url = "https://api.openai.com/v1" 、api_key = "sk-..." )
# 変更後: Aurora Gateway
client = OpenAI (base_url = "http://localhost:8080/v1" 、api_key = "your-aurora-key")
SDKの変更はありません。形式の変更はありません。 Base_url を交換するだけです。
14 のプロバイダー タイプ — OpenAI、Anthropic、Gemini、Groq、DeepSeek、OpenRouter、xAI、Z.ai、MiniMax、Azure OpenAI、Oracle

、オラマ、vLLM、ジナ
自動検出 — API キーを環境変数、再起動、プロバイダーとして設定し、そのすべてのモデルが自動的に表示されます
プロバイダー プール — 複数のキー/エンドポイントをグループ化し、ラウンドロビンまたは加重分散による負荷分散、健全性を認識したフェイルオーバー
モデルのエイリアス — ゲートウェイ全体で任意のモデルの名前を変更/カスタム識別子に再マッピング
モデルのオーバーライド - ユーザー パスごとに特定のモデルを有効または無効にし、ダッシュボードまたは user_pricing.yaml を介して永続化します。
フォールバック — 5xx/429 での自動フェイルオーバー、または失敗したプロバイダーとモデルをバックアップにマッピングする手動ルール (構成または外部 JSON から)
回復力 — ジッターを伴う指数関数的なバックオフ、プロバイダーごとのサーキット ブレーカー (クローズ→オープン→ハーフオープン)、グローバルな再試行/サーキット ブレーカー設定のプロバイダーごとのオーバーライド
複数のインスタンス - OPENAI_EAST_API_KEY と OPENAI_WEST_API_KEY を別のプロバイダーとして実行します
カスタムベース URL — プロバイダーのエンドポイント (企業プロキシ、地域エンドポイント) をオーバーライドします。
パススルー — /p/{provider}/* (チャットの完了だけでなく) 完全なアップストリーム API アクセス用。どのプロバイダー タイプがパススルー ルートを取得するかをフィルターする
構成主導のワークフロー — リクエストごとのルーティング、キャッシュ、ガードレール、監査、使用状況、予算、およびフォールバック動作は、永続化されたワークフロー ドキュメントによって制御されます。
OpenAI 互換 — /v1/chat/completions 、 /v1/embeddings 、 /v1/rerank 、 /v1/models 、 /v1/files 、 /v1/batches
レスポンス API — 完全な CRUD、キャンセル、入力アイテム、コンパクトを含む /v1/responses
Anthropic 互換 — /v1/messages 、 /v1/messages/count_tokens (ネイティブ Anthropic ワイヤー形式); /v1/messages でのオプションの専用 Ingress
ストリーミング — すべてのエンドポイントの SSE ストリーミング、エンドツーエンドで維持
エイリアスのみ保持モード — 生のプロバイダー モデルを /v1/models から非表示にし、エイリアス名のみを公開します
構成されたプロバイダー モデル モード - フォールバック (リストに追加)

自動検出されるモデル）または許可リスト（明示的にリストされたモデルのみを提供）
正確なキャッシュ — リクエストに応じた SHA-256 ハッシュ一致、Redis ベースの非同期書き込み
セマンティック キャッシュ — 設定可能なしきい値によるベクトル類似性、Qdrant、pgvector、Pinecone、Weaviate をサポート
プロンプト キャッシュ —cache_control を Anthropic/OpenAI/Gemini ネイティブ プロンプト キャッシュに転送します。構成可能なモード ( auto 、 manual 、 off )、コンポーネントの切り替え、および最小トークンしきい値
モデル レジストリ キャッシュ - ローカル ファイル システム + Redis、オフラインでも安全。フィールドごとのユーザー価格設定の上書きによるベンダー提供の JSON スナップショットをサポートします
マスターキー — トップレベルのゲートウェイ認証
マネージド API キー — 範囲指定、レート制限、キーごとのモデル認証、使用状況統計
レート制限 — インメモリまたは Redis によるキーごとのレート制限
PII 編集 - 電子メール、電話、SSN、クレジット カードの検出とマスキング
即時注入ブロック - 注入の試行を検出してブロックします。
システム プロンプトの保護 — システム プロンプトを挿入、上書き、または修飾します。
正規表現ブロック — ブロックまたはサニタイズアクションによるカスタムパターンマッチング
長さ制限 — リクエストに対する文字数/トークン数の強制
LLM ベースの変更 — 補助 LLM 呼び出し (匿名化、カスタム プロンプト) を介してメッセージ コンテンツを書き換えるガードレール
ガードレールの方向と順序 — プロバイダーのディスパッチ前 (入力)、応答後 (出力)、またはその両方で実行します。同じ順序のガードレールが並行して走る
バッチ ガードレール — 構成されたガードレールを /v1/batches リクエストのインライン項目に適用します
監査ログ - 完全なリクエスト/レスポンスのキャプチャ、バッファリングされた書き込み、設定可能な保持期間 (ボディ/ヘッダーのログ、バッファ サイズ、フラッシュ間隔)、ライブ SSE ストリーム
使用状況分析 - モデルごとのトークンカウント、コスト追跡、モデル/ユーザーパス別の日次集計、価格再計算アクション
プロメテウスのメトリクス — aurora_req

uests_total 、 aurora_request_duration_seconds 、 aurora_requests_in_flight 、およびゲートウェイ フェーズのタイミング
管理者ダッシュボード — Go バイナリに組み込まれた React SPA: プロバイダー、プール、モデル、エイリアス、ガードレール、キャッシュ、使用状況、監査、認証キー、ワークフロー、コンソール、プレイグラウンド
pprof エンドポイント — /debug/pprof/* でランタイム プロファイリングを実行します (ヒープ、ゴルーチン、ミューテックス、ブロック、スレッド作成)
構造化ログ - 構成可能な形式 (JSON/テキスト)、レベル (デバッグ/情報/警告/エラー)、ソース情報、サービス メタデータ
トークン セーバー — ポリシー主導の出力圧縮 (プロファイル: concise、caveman、ultra、wenyan)。包含/除外フィルターを介して特定のモデル/プロバイダーに範囲を絞ります。構成可能なエラー時の動作 (許可/ブロック)
価格管理 - モデルごとの価格の上書き、再計算、インポート/エクスポート
使用量の予算 — キーごとの使用量の追跡と制限、ワークフロー機能フラグによるリクエストごとの予算の適用
単一バイナリ — npm install -g iaurora または docker pull aurorahq/aurora
CLI — aurora init 、 aurora モデル sync/diff/show 、 aurora update 、 aurora uninstall
CLI ツール API — CLI 構成同期用の管理 REST エンドポイント、個別にゲート制御
Swagger ドキュメント — /swagger/index.html (ビルドタグゲート付き)
構成プロファイル — ローカル、ローカル電源、およびチーム展開用の事前構築された構成
3 層構成 — コードのデフォルト → config.yaml → 環境変数 (環境変数が優先)
Helm チャート — 事前に構築された Helm チャートを使用して Kubernetes にデプロイします
Docker Compose — 完全なインフラストラクチャ スタック: Redis、PostgreSQL、Qdrant、Prometheus、Grafana
Grafana ダッシュボード — リクエスト レート、エラー、レイテンシ、実行中のリクエスト、モデルごとの内訳を表示する事前設定されたパネル
60 秒以内に AI トラフィックのルーティングを開始します。
npm install -g iaurora
mkdir 私のゲートウェイ && cd 私のゲートウェイ
aurora init # config.yaml、.env、data/ を作成します
プロバイダー キーを .env に設定します。
＃─

─ 必須 ───────────────────
AURORA_MASTER_KEY = "あなたの安全なキー"
# ── プロバイダー API キー (少なくとも 1 つ) ──────────
OPENAI_API_KEY = " sk-... "
ANTHROPIC_API_KEY = " sk-ant-... "
GEMINI_API_KEY = " ... "
GROQ_API_KEY = " gsk_... "
DEEPSEEK_API_KEY = " ... "
OPENROUTER_API_KEY = " ... "
XAI_API_KEY = " ... "
ZAI_API_KEY = " ... "
MINIMAX_API_KEY = " ... "
AZURE_API_KEY = " ... "
ORACLE_API_KEY = " ... "
OLLAMA_API_KEY = " ... "
VLLM_API_KEY = " ... "
JINA_API_KEY = " ... "
# ── オプション機能の切り替え (有効にするには true を設定) ───
LOGGING_ENABLED = true # ストレージへのログの監査
METRICS_ENABLED = true # Prometheus /metrics エンドポイント
GUARDRAILS_ENABLED = true # コンテンツ安全フィルター
TOKEN_SAVER_ENABLED = true # トークンの使用を削減するための出力圧縮
# ── 本番ストレージ ────── ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─
# STORAGE_TYPE=postgresql
# POSTGRES_URL=postgres://user:pass@localhost:5432/aurora
# ── REDIS CACHE (モデルキャッシュ + レスポンスキャッシュ) ──────
# REDIS_URL=redis://localhost:6379
# RESPONSE_CACHE_SIMPLE_ENABLED=true
オーロラ
オプション B — インライン環境変数 (.env は必要ありません)
Linux / macOS
AURORA_MASTER_KEY=あなたの安全なキー \
OPENAI_API_KEY=sk-... \
ANTHROPIC_API_KEY=sk-ant-... \
GEMINI_API_KEY=... \
GROQ_API_KEY=gsk_... \
DEEPSEEK_API_KEY=... \
OPENROUTER_API_KEY=... \
XAI_API_KEY=... \
ZAI_API_KEY=... \
MINIMAX_API_KEY=... \
AZURE_API_KEY=... \
ORACLE_API_KEY=... \
OLLAMA_API_KEY=... \
VLLM_API_KEY=... \
JINA_API_KEY=... \
LOGGING_ENABLED=true \
METRICS_ENABLED=true \
ガードレール_ENA

BLED=true \
TOKEN_SAVER_ENABLED=true \
オーロラ
Windows PowerShell
$ env: AURORA_MASTER_KEY = " あなたの安全なキー " ; `
$ env: OPENAI_API_KEY = " sk-... " ; `
$ env: ANTHROPIC_API_KEY = " sk-ant-... " ; `
$ env:GEMINI_API_KEY = " ... " ; `
$ env: GROQ_API_KEY = " gsk_... " ; `
$ env: DEEPSEEK_API_KEY = " ... " ; `
$ env: OPENROUTER_API_KEY = " ... " ; `
$ env: XAI_API_KEY = " ... " ; `
$ env: ZAI_API_KEY = " ... " ; `
$ env: MINIMAX_API_KEY = " ... " ; `
$ env: AZURE_API_KEY = " ... " ; `
$ env: ORACLE_API_KEY = " ... " ; `
$ env: OLLAMA_API_KEY = " ... " ; `
$ env: VLLM_API_KEY = " ... " ; `
$ env: JINA_API_KEY = " ... " ; `
$ env: LOGGING_ENABLED = " true " ; `
$ env: METRICS_ENABLED = " true " ; `
$ env: GUARDRAILS_ENABLED = " true " ; `
$ env: TOKEN_SAVER_ENABLED = " true " ; `
オーロラ
Windows CMD
set AURORA_MASTER_KEY = 安全なキー ^
&& set OPENAI_API_KEY = sk-... ^
&& set ANTHROPIC_API_KEY = sk-ant-... ^
&& set GEMINI_API_KEY = ... ^
&& GROQ_API_KEY = gsk_... を設定します ^
&& set DEEPSEEK_API_KEY = ... ^
&& OPENROUTER_API_KEY を設定 = ... ^
&& set XAI_API_KEY = ... ^
&& set ZAI_API_KEY = ... ^
&& set MINIMAX_API_KEY = ... ^
&& AZURE_API_KEY を設定 = ... ^
&& ORACLE_API_KEY を設定 = ... ^
&& OLLAMA_API_KEY を設定 = ... ^
&& set VLLM_API_KEY = ... ^
&& set JINA_API_KEY = ... ^
&& set LOGGING_ENABLED = true ^
&& set METRICS_ENABLED = true ^
&& GUARDRAILS_ENABLED = true を設定します ^
&& TOKEN_SAVER_ENABLED = true を設定します ^
&& オーロラ
オプション C — Docker
docker run -d --name aurora -p 8080:8080 \
-e AURORA_MASTER_KEY= " あなたの安全なキー " \
-e OPENAI_API_KEY= " sk-... " \
-e ANTHROPIC_API_KEY= " sk-ant-... "

[切り捨てられた]

## Original Extract

The Fastest enterprise AI gateway — route LLM traffic across OpenAI, Anthropic, Gemini, Groq & 30+ providers via a single API. Self-hosted, no vendor lock-in. ( 55x faster than litellm ) - aurorallm/aurora

Every Ai project needs the same infra - API keys, providers SDKs, retry logic, cost tracking. it should be infra layer.
Aurora is a Go gateway that handles that layer, ~15MB binary, no dependencies. It proxies requests to OpenaAI, Anthropic, Gemini etc, and can convert between formats - use the OpenAi SDK and Aurora translates to Anthropic's wire format automatically. No format conversion code in backend. It helps:
- track costs per customer/team
- switch models and providers without changing code
- debug requests and responses flows more easily
- reduce costs with exact and semantic caching How its different than other gateways:
- auto-discovers providers models
- provider pools with automatic failover switch
- built in guardials, audit logging with export, analytics, admin dashboard
- model aliases, regional deployment to reduce latency.
- can be used in infra directly and admin endpoints let control everything, creating managed api keys on user onboard and more. The oss version (Apache 2.0) includes everyhting. Enterprise adds SSO, RBAC, and tenant isolation. npm install -g iaurora website: https://aurorallm.online feedbacks are welcome!

GitHub - aurorallm/aurora: The Fastest enterprise AI gateway — route LLM traffic across OpenAI, Anthropic, Gemini, Groq & 30+ providers via a single API. Self-hosted, no vendor lock-in. ( 55x faster than litellm ) · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Uh oh!
There was an error while loading. Please reload this page .
aurorallm
/
aurora
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .github/ workflows .github/ workflows apps/ aurora apps/ aurora configs configs configuration configuration dashboard-ui dashboard-ui docs-assets docs-assets docs docs heartbeat-worker heartbeat-worker helm helm internal internal monitoring monitoring .dockerignore .dockerignore .goreleaser.yaml .goreleaser.yaml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md artifacthub-repo.yml artifacthub-repo.yml docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum package.json package.json View all files Repository files navigation
Aurora - The Fastest AI Gateway
One API for every AI provider. Self-hosted. No vendor lock-in.
14 provider types • OpenAI & Anthropic compatible • Go • Apache 2.0 • Built for raw speed
Aurora sits between your app and LLM providers. Your app sends requests using the standard OpenAI or Anthropic SDK — Aurora routes them to whichever provider you've configured. One format handles everything — you dont need to worry about provider-specific formats.
# Before: hardcoded provider
client = OpenAI ( base_url = "https://api.openai.com/v1" , api_key = "sk-..." )
# After: Aurora Gateway
client = OpenAI ( base_url = "http://localhost:8080/v1" , api_key = "your-aurora-key" )
No SDK changes. No format changes. Just swap the base_url .
14 provider types — OpenAI, Anthropic, Gemini, Groq, DeepSeek, OpenRouter, xAI, Z.ai, MiniMax, Azure OpenAI, Oracle, Ollama, vLLM, Jina
Auto-discovery — set an API key as an env var, restart, provider + all its models appear automatically
Provider pools — group multiple keys/endpoints, load-balance with round-robin or weighted distribution, health-aware failover
Model aliases — rename/remap any model to a custom identifier across the entire gateway
Model overrides — enable or disable specific models per user path, persisted via dashboard or user_pricing.yaml
Fallback — automatic failover on 5xx/429, or manual rules (from config or external JSON) mapping failed provider+model to backups
Resilience — exponential backoff with jitter, circuit breaker per provider (closed → open → half-open), per-provider override of global retry/circuit-breaker settings
Multiple instances — run OPENAI_EAST_API_KEY and OPENAI_WEST_API_KEY as separate providers
Custom base URLs — override any provider's endpoint (corporate proxies, regional endpoints)
Passthrough — /p/{provider}/* for full upstream API access (not just chat completions); filter which provider types get passthrough routes
Config-driven workflows — per-request routing, caching, guardrail, audit, usage, budget, and fallback behavior controlled by persisted workflow documents
OpenAI-compatible — /v1/chat/completions , /v1/embeddings , /v1/rerank , /v1/models , /v1/files , /v1/batches
Responses API — /v1/responses with full CRUD, cancel, input items, compact
Anthropic-compatible — /v1/messages , /v1/messages/count_tokens (native Anthropic wire format); optional dedicated ingress at /v1/messages
Streaming — SSE streaming for all endpoints, preserved end-to-end
Keep-only-aliases mode — hide raw provider models from /v1/models and expose only aliased names
Configured provider models mode — fallback (add listed models to auto-discovered) or allowlist (only serve explicitly listed models)
Exact cache — SHA-256 hash match on request, Redis-backed, async writes
Semantic cache — vector similarity with configurable threshold, supports Qdrant, pgvector, Pinecone, Weaviate
Prompt cache — forwards cache_control to Anthropic/OpenAI/Gemini native prompt caching; configurable modes ( auto , manual , off ), component toggles, and minimum token threshold
Model registry cache — local filesystem + Redis, offline-safe; supports vendored JSON snapshots with per-field user pricing overrides
Master key — top-level gateway auth
Managed API keys — scoped, rate-limited, per-key model authorization, usage stats
Rate limiting — per-key rate limiting backed by in-memory or Redis
PII redaction — email, phone, SSN, credit card detection and masking
Prompt injection blocking — detects and blocks injection attempts
System prompt protection — inject, override, or decorate system prompts
Regex blocking — custom pattern matching with block or sanitize actions
Length limits — character/token count enforcement on requests
LLM-based altering — guardrail that rewrites message content via an auxiliary LLM call (anonymization, custom prompts)
Guardrail direction & ordering — run before provider dispatch ( input ), after response ( output ), or both; same-order guardrails run in parallel
Batch guardrails — apply configured guardrails to inline items in /v1/batches requests
Audit logging — full request/response capture, buffered writes, configurable retention (body/header logging, buffer size, flush interval), live SSE stream
Usage analytics — per-model token counting, cost tracking, daily aggregation by model/user-path, pricing recalculation action
Prometheus metrics — aurora_requests_total , aurora_request_duration_seconds , aurora_requests_in_flight , plus gateway phase timing
Admin dashboard — React SPA built into the Go binary: providers, pools, models, aliases, guardrails, cache, usage, audit, auth keys, workflows, console, playground
pprof endpoints — Go runtime profiling at /debug/pprof/* (heap, goroutine, mutex, block, threadcreate)
Structured logging — configurable format (JSON/text), level (debug/info/warn/error), source info, service metadata
Token saver — policy-driven output compression (profiles: concise, caveman, ultra, wenyan); scoped to specific models/providers via include/exclude filters; configurable on-error behavior (allow/block)
Pricing management — per-model pricing overrides, recalculation, import/export
Usage budgets — per-key usage tracking and limits, per-request budget enforcement via workflow feature flags
Single binary — npm install -g iaurora or docker pull aurorahq/aurora
CLI — aurora init , aurora models sync/diff/show , aurora update , aurora uninstall
CLI tools API — admin REST endpoints for CLI configuration sync, gated separately
Swagger docs — /swagger/index.html (build-tag gated)
Config profiles — pre-built configs for local, local-power, and team deployments
3-layer config — code defaults → config.yaml → env vars (env vars win)
Helm chart — deploy on Kubernetes with pre-built Helm chart
Docker Compose — full infrastructure stack: Redis, PostgreSQL, Qdrant, Prometheus, Grafana
Grafana dashboard — pre-configured panels for request rate, errors, latency, in-flight requests, per-model breakdown
Start routing AI traffic in 60 seconds.
npm install -g iaurora
mkdir my-gateway && cd my-gateway
aurora init # creates config.yaml, .env, data/
Set your provider keys in .env :
# ── REQUIRED ──────────────────────────────────────────────
AURORA_MASTER_KEY = " your-secure-key "
# ── PROVIDER API KEYS (at least one) ─────────────────────
OPENAI_API_KEY = " sk-... "
ANTHROPIC_API_KEY = " sk-ant-... "
GEMINI_API_KEY = " ... "
GROQ_API_KEY = " gsk_... "
DEEPSEEK_API_KEY = " ... "
OPENROUTER_API_KEY = " ... "
XAI_API_KEY = " ... "
ZAI_API_KEY = " ... "
MINIMAX_API_KEY = " ... "
AZURE_API_KEY = " ... "
ORACLE_API_KEY = " ... "
OLLAMA_API_KEY = " ... "
VLLM_API_KEY = " ... "
JINA_API_KEY = " ... "
# ── OPTIONAL FEATURE TOGGLES (set true to enable) ────────
LOGGING_ENABLED = true # Audit logging to storage
METRICS_ENABLED = true # Prometheus /metrics endpoint
GUARDRAILS_ENABLED = true # Content safety filters
TOKEN_SAVER_ENABLED = true # Output compression to cut token use
# ── PRODUCTION STORAGE ───────────────────────────────────
# STORAGE_TYPE=postgresql
# POSTGRES_URL=postgres://user:pass@localhost:5432/aurora
# ── REDIS CACHE (model cache + response cache) ──────────
# REDIS_URL=redis://localhost:6379
# RESPONSE_CACHE_SIMPLE_ENABLED=true
aurora
Option B — inline env vars (no .env needed)
Linux / macOS
AURORA_MASTER_KEY=your-secure-key \
OPENAI_API_KEY=sk-... \
ANTHROPIC_API_KEY=sk-ant-... \
GEMINI_API_KEY=... \
GROQ_API_KEY=gsk_... \
DEEPSEEK_API_KEY=... \
OPENROUTER_API_KEY=... \
XAI_API_KEY=... \
ZAI_API_KEY=... \
MINIMAX_API_KEY=... \
AZURE_API_KEY=... \
ORACLE_API_KEY=... \
OLLAMA_API_KEY=... \
VLLM_API_KEY=... \
JINA_API_KEY=... \
LOGGING_ENABLED=true \
METRICS_ENABLED=true \
GUARDRAILS_ENABLED=true \
TOKEN_SAVER_ENABLED=true \
aurora
Windows PowerShell
$ env: AURORA_MASTER_KEY = " your-secure-key " ; `
$ env: OPENAI_API_KEY = " sk-... " ; `
$ env: ANTHROPIC_API_KEY = " sk-ant-... " ; `
$ env: GEMINI_API_KEY = " ... " ; `
$ env: GROQ_API_KEY = " gsk_... " ; `
$ env: DEEPSEEK_API_KEY = " ... " ; `
$ env: OPENROUTER_API_KEY = " ... " ; `
$ env: XAI_API_KEY = " ... " ; `
$ env: ZAI_API_KEY = " ... " ; `
$ env: MINIMAX_API_KEY = " ... " ; `
$ env: AZURE_API_KEY = " ... " ; `
$ env: ORACLE_API_KEY = " ... " ; `
$ env: OLLAMA_API_KEY = " ... " ; `
$ env: VLLM_API_KEY = " ... " ; `
$ env: JINA_API_KEY = " ... " ; `
$ env: LOGGING_ENABLED = " true " ; `
$ env: METRICS_ENABLED = " true " ; `
$ env: GUARDRAILS_ENABLED = " true " ; `
$ env: TOKEN_SAVER_ENABLED = " true " ; `
aurora
Windows CMD
set AURORA_MASTER_KEY = your-secure-key ^
&& set OPENAI_API_KEY = sk-... ^
&& set ANTHROPIC_API_KEY = sk-ant-... ^
&& set GEMINI_API_KEY = ... ^
&& set GROQ_API_KEY = gsk_... ^
&& set DEEPSEEK_API_KEY = ... ^
&& set OPENROUTER_API_KEY = ... ^
&& set XAI_API_KEY = ... ^
&& set ZAI_API_KEY = ... ^
&& set MINIMAX_API_KEY = ... ^
&& set AZURE_API_KEY = ... ^
&& set ORACLE_API_KEY = ... ^
&& set OLLAMA_API_KEY = ... ^
&& set VLLM_API_KEY = ... ^
&& set JINA_API_KEY = ... ^
&& set LOGGING_ENABLED = true ^
&& set METRICS_ENABLED = true ^
&& set GUARDRAILS_ENABLED = true ^
&& set TOKEN_SAVER_ENABLED = true ^
&& aurora
Option C — Docker
docker run -d --name aurora -p 8080:8080 \
-e AURORA_MASTER_KEY= " your-secure-key " \
-e OPENAI_API_KEY= " sk-... " \
-e ANTHROPIC_API_KEY= " sk-ant-... "

[truncated]
