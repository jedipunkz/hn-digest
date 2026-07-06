---
source: "https://github.com/CURSED-ME/loopers-oss"
hn_url: "https://news.ycombinator.com/item?id=48804865"
title: "Show HN: Loopers – Open-source fail-closed firewall for AI agent runtimes"
article_title: "GitHub - CURSED-ME/loopers-oss: A simple reverse proxy that can save you from unexpected AI Bills · GitHub"
author: "varad-khoriya"
captured_at: "2026-07-06T15:02:15Z"
capture_tool: "hn-digest"
hn_id: 48804865
score: 1
comments: 0
posted_at: "2026-07-06T14:10:45Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Loopers – Open-source fail-closed firewall for AI agent runtimes

- HN: [48804865](https://news.ycombinator.com/item?id=48804865)
- Source: [github.com](https://github.com/CURSED-ME/loopers-oss)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T14:10:45Z

## Translation

タイトル: Show HN: Loopers – AI エージェント ランタイム用のオープンソースのフェールクローズ ファイアウォール
記事のタイトル: GitHub - CURSED-ME/loopers-oss: 予期せぬ AI 請求から救えるシンプルなリバース プロキシ · GitHub
説明: 予期せぬ AI 請求から身を守ることができるシンプルなリバース プロキシ - CURSED-ME/loopers-oss

記事本文:
GitHub - CURSED-ME/loopers-oss: 予期せぬ AI 請求から身を守ることができるシンプルなリバース プロキシ · GitHub
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
呪われた私
/
ルーパーズ-OSS
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
249 コミット 249 コミット .github .github ドキュメント ドキュメント cmd cmd デモ デモ ドキュメント ドキュメント 例/ポリシー 例/ポリシー grafana grafana helm/ ルーパーズ helm/ ルーパーズ 内部 内部 pkg/ api pkg/ api sdk sdk .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml .release-please-manifest.json .release-please-manifest.json AGENT_README.md AGENT_README.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.goreleaser Dockerfile.goreleaser ライセンス ライセンス Makefile Makefile README.md README.md SECURITY.md SECURITY.md docker-compose.demo.yml docker-compose.demo.yml docker-compose.yml docker-compose.yml Final_results.md Final_results.md go.mod go.mod go.sum go.sum loopers.example.yaml loopers.example.yaml 価格.yaml 価格.yaml release-please-config.json release-please-config.json runs.json run.json test-config.yaml test-config.yaml test_output.txt test_output.txt test_viper.go test_viper.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Loopers – エージェント時代のファイアウォール
予算を超えてしまう前にループを断ち切りましょう。
AIエージェントかLLMか？技術的なコンテキストについては、AI に最適化された AGENT_README.md をお読みください。
Loopers は、エージェント時代のベアメタル、遅延ゼロのファイアウォールです。 500 を超える AI モデル (OpenAI、Anthropic、Gemini、Groq、Ollama、vLLM などの 14 のプロバイダー) に加え、あらゆる OpenAI 互換エンドポイントのリクエストをネイティブにインターセプトして、トークンの過剰使用を防ぎ、エージェントの暴走を阻止し、LLM ジャッキングのような壊滅的な請求ショックから保護します。
Looper を最速で最も安全な AI ファイアウォールにするためのアップデートを常に提供しています。完全な構成の詳細については、次のサイトを参照してください。

ルーパーのドキュメント 。
Local Policy Engine (OPA/Rego) (最新) : Rego 言語を使用して、きめ細かい属性ベースのアクセス制御 (ABAC) ポリシーを作成します。組み込みの Open Policy Agent は、ホットリロード サポートを使用して、すべてのリクエスト (LLM 呼び出しと MCP ツール呼び出しの両方) を .rego ファイルに対して評価します。プロキシを再起動することなく、破壊的なツールをブロックし、チームごとにモデルを制限し、環境レベルのガバナンスを強制します。 『ポリシー エンジン ガイド』を参照してください。
エージェント ID とキー メタデータ: --agent-name 、 --owner 、 --allowed-tools 、 --allowed-providers 、および --tags を使用して、すべてのプロキシ キーに豊富な ID を付加します。このメタデータは、ポリシー評価、OpenTelemetry スパン、および完全な監査証跡のセキュリティ イベントに流れ込みます。
CrewAI および AutoGen フレームワーク アダプター : CrewAI および Microsoft AutoGen 用のネイティブ Python SDK アダプター。 get_loopers_crewai_llm() または get_loopers_autogen_config() を追加して、すべてのエージェント LLM トラフィックを設定なしで Loopers プロキシ経由でルーティングします。 「フレームワーク アダプター ガイド」を参照してください。
キーごとのレート制限 : アトミックな Redis Lua スクリプトを利用したスライディング ウィンドウ レート リミッタ。設定で rate_limit.requests_per_minut を設定して、予算制限とは関係なく、キーごとのリクエスト速度を制限します。
モデル コンテキスト プロトコル (MCP) ガバナンス: ルーパーは MCP トラフィックをネイティブに管理します。機能には、透過的な JSON-RPC 2.0 プロキシ、ツールごとの予算の強制 (例: Snowflake クエリごとに 0.05 ドル)、無限ループを停止する決定論的なツール呼び出しサーキット ブレーカー、厳密な Blast Radius (横方向の移動) 防止が含まれます。 MCP セットアップ ガイドを参照して、2 分で開始してください。
セキュリティ イベントと OpenTelemetry : バジェット/ループ ブロックの LLM セキュリティ ペイロードの OWASP Top 10 を出力し、EU AI 法準拠向けに設計されたスマート サンプリング プロセッサによる W3C OTLP トレースをサポートします。
ループ検出エンジン v1.1

: 自律エージェント向けの決定論的かつファジーなサーキット ブレーカー。多態性/変異プロンプトを捕捉するための Bi-Gram Jaccard 類似性マッチング、速度リミッター、TOCTOU セーフな施行のためのストール検出機能を備えています。
汎用 OpenAI 互換エンドポイント : 独自のプロバイダーを導入してください!予算全体の適用を維持しながら、トラフィックを vLLM、ローカル Llama.cpp、またはカスタム プロキシにルーティングします。
動的価格設定 Fetcher : ハンズフリーのトークン アカウンティング。 Looper は、リモート JSON エンドポイントからリアルタイムのトークン価格を自動的に同期します。
エンタープライズ グレードのセキュリティ: 専用の分離された管理ポート ( /metrics )、厳格な TLS 強制、およびベアメタル展開のための安全な Redis 構成。
LangChain、LlamaIndex、CrewAI、および AutoGen アダプター: ネイティブ エージェント フレームワークでセッション ID を簡単に挿入し、ステップ数を追跡するドロップイン Python SDK ラッパー ( ChatLoopers 、 LoopersLLM 、 get_loopers_crewai_llm 、 get_loopers_autogen_config )。
自律エージェントがループに陥ったり、API キーが侵害されたりすると、数分で数千ドルを消費する可能性があります。 Loopers はアラートやダッシュボードではなく、キルスイッチです。
アトミック正確性保証 : 単一の Redis Lua トランザクションでチェックを実行し、極端な同時実行下での TOCTOU 競合状態を防ぎます。
ゼロストレージ セキュリティ モデル: パススルー アーキテクチャ。 API キーは、リクエストのライフサイクル中のみメモリ内に保持されます。ディスク/データベースへの永続性がゼロになり、データ侵害の影響を受けなくなります。
ミリ秒未満のオーバーヘッド : httputil.ReverseProxy と Redis を使用して Go で記述されており、リクエスト パスに追加される遅延はわずか 1 ～ 2 ミリ秒です。コールド スタートやストリーミング パフォーマンスのブロックはありません。
フェイルクローズ保証 : Redis またはプロキシがダウンした場合はフェイルクローズし、リクエストを即座にブロックしてウォレットを保護します。
ミッドストリーム カットオフ : ストリーミング サーブをインターセプトします

r-Sent Events (SSE) 応答、リアルタイムでトークンをカウントし、制限を超えた場合は接続を即座に切断します。
Loopers は、単純な可観測性よりも絶対的なセキュリティと正確性を優先する、高性能インフラストラクチャ レベルのファイアウォールとして特別に設計されています。
パフォーマンスのベンチマーク (エピソード 1)
Loopers は、予算執行に支障をきたすことなく、大量の同時トラフィックの急増に対処できるように設計されています。 LiteLLM などの Python/FastAPI 代替製品に対する最新の LLM ファイアウォール ベンチマークで、Loopers は次のことを実証しました。
最終ベンチマーク結果 で、生データと方法論に関する完全な詳細を読んでください。
プロバイダー
モデル名
ストリーミング
非ストリーミング
予算の執行
トークンカウンティング
OpenAI
gpt-4o 、 gpt-4o-mini など
サポートされています
サポートされています
サポートされています
サポートあり (tiktoken)
人間的
クロード-3-5-ソネット など
サポートされています
サポートされています
サポートされています
サポートあり (countTokens API)
Google ジェミニ
gemini-2.5-フラッシュ など
サポートされています
サポートされています
サポートされています
サポートあり (countTokens API)
AWS の基盤
クロード/岩盤上のラマ
サポートされています
サポートされています
サポートされています
サポートあり (モデルトークナイザー)
Azure OpenAI
Azure 上の GPT モデル
サポートされています
サポートされています
サポートされています
サポートあり (tiktoken)
ミストラルAI
ミストラルラージ など
サポートされています
サポートされています
サポートされています
サポートあり (tiktoken)
グロク
Groq のラマ 3 など
サポートされています
サポートされています
サポートされています
サポートあり (tiktoken)
コヒア
command-r など
サポートされています
サポートされています
サポートされています
サポートあり (モデルトークナイザー)
ディープシーク
ディープシークチャットなど
サポートされています
サポートされています
サポートされています
サポートあり (tiktoken)
一緒に
ラマ3 on Togetherなど
サポートされています
サポートされています
サポートされています
サポートあり (tiktoken)
ジェネリック (BYO)
任意の OpenAI 互換モデル (LM Studio、LocalAI、OpenRouter など)
サポートされています
サポートされています
サポートされています
サポートあり (tiktoken)
MCPサーバー
任意の JSON-RPC 2.0 MCP サーバー
該当なし
該当なし
サポートあり (ツールごとのコスト)
該当なし
60秒で試してみる

(行く必要はありません)
実際の API キーに触れる前に Looper が暴走リクエストをブロックする様子を見てみたいですか?自己完結型のデモを開始します。
git clone https://github.com/CURSED-ME/loopers-oss.git
CD ルーパー-OSS
docker-compose -f docker-compose.demo.yml アップ
ブートストラップ コンテナのログで、準備ができているcurlコマンドを確認してください。デモでは模擬 O​​penAI サーバーを使用するため、実際のクレジットは使用されません。
カール -sSL https://github.com/CURSED-ME/loopers-oss/releases/latest/download/loopers_Linux_x86_64.tar.gz | tar -xz && sudo mv ルーパー /usr/local/bin/
または、Docker イメージを直接プルします。
docker pull ghcr.io/cursed-me/loopers:latest
または、ウィザードを介して初期化します (Go が必要です)。
github.com/CURSED-ME/loopers-oss/cmd/loopers init を実行します。
ステップ 2: プロキシをスピンアップする
docker-compose up -d
ステップ 3: キーの作成と予算の構成
OpenAI の API プロキシ キーを生成します。
docker-compose exec loopers /app/loopers key create --name my-app-key --provider openai
生成された生のキー ( lp-xxx ) とそのハッシュに注目してください。
キー ハッシュの 5 つの詳細な時間枠にわたって予算制限を設定します。
docker-compose exec loopers /app/loopers 予算セット < KEY_HASH > \
--分 0.50 \
-- 時給 2.00 \
--毎日 10.00 \
--毎週 50.00 \
--月額 150.00
5 つのウィンドウ ( -- minutes 、 --hourly 、 --daily 、 --weekly 、 --monthly ) はすべてオプションであり、自由に組み合わせることができます。最初に制限に達した場合が勝ちとなり、リクエストがブロックされます。
ステップ 4: リクエストをルーティングする
公式 SDK または生の cURL のいずれかを使用して、Loopers プロキシ経由で API 呼び出しを行います。
curl -X POST http://localhost:8080/openai/v1/chat/completions \
-H " 認証: ベアラー <RAW_LP_KEY> " \
-H " X-Loopers プロバイダー キー: <YOUR_REAL_OPENAI_KEY> " \
-H " Content-Type: application/json " \
-d ' {
"モデル": "gpt-4o-mini",
"messages": [{"role": "user", "content": "Hello, Loopers!"}]
} '
CLI リファレンス
コマンド
説明

オプション
ルーパーの初期化
対話型ウィザード — ルーパーズ.yaml と docker-compose.yml を生成します
ルーパーがサーブする
プロキシサーバーを起動する
ルーパー キー作成 --name <n> --provider <p>
新しいプロキシ キーを作成します (プロバイダー: openai 、 anthropic 、 gemini 、 bedrock 、 azure 、 misstral 、 groq 、 cohere 、 deepseek 、 together 、 ollama 、 Fireworks 、 xai 、 vllm )。オプションの ID フラグをサポートします: --agent-name 、 --owner 、 --allowed-tools 、 --allowed-providers 、 --tags
ルーパーキーリスト
すべてのプロキシ キーとメタデータをリストする
ルーパー キーは <ハッシュ> を取り消します
ハッシュによるキーの取り消し
ルーパーの予算セット <ハッシュ> [フラグ]
予算制限を設定する ( -- minutes 、 --hourly 、 --daily 、 --weekly 、 --monthly )
ルーパーの予算ステータス <ハッシュ>
現在の支出とキーの制限を表示する
ゼロ SDK 統合
SDK ラッパーを使用できない場合は、Looper を指すようにクライアントを構成し、ほとんどの SDK で利用可能なdefault_headers パラメーターを使用して必要な HTTP ヘッダーを挿入することで、標準の OpenAI 互換クライアントを使用できます。
OSをインポートする
openaiインポートからOpenAI
クライアント = OpenAI (
base_url = "http://localhost:8080/openai/v1" , # Loopers プロキシを介したルート
api_key = "lp-xxx" , # Loopers プロキシ キー
デフォルトヘッダー = {
「X-Loopers-Provider-Key」: os 。環境。 get ( "OPENAI_API_K

[切り捨てられた]

## Original Extract

A simple reverse proxy that can save you from unexpected AI Bills - CURSED-ME/loopers-oss

GitHub - CURSED-ME/loopers-oss: A simple reverse proxy that can save you from unexpected AI Bills · GitHub
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
CURSED-ME
/
loopers-oss
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
249 Commits 249 Commits .github .github Documentation Documentation cmd cmd demo demo docs docs examples/ policies examples/ policies grafana grafana helm/ loopers helm/ loopers internal internal pkg/ api pkg/ api sdk sdk .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml .release-please-manifest.json .release-please-manifest.json AGENT_README.md AGENT_README.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.goreleaser Dockerfile.goreleaser LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md docker-compose.demo.yml docker-compose.demo.yml docker-compose.yml docker-compose.yml final_results.md final_results.md go.mod go.mod go.sum go.sum loopers.example.yaml loopers.example.yaml pricing.yaml pricing.yaml release-please-config.json release-please-config.json runs.json runs.json test-config.yaml test-config.yaml test_output.txt test_output.txt test_viper.go test_viper.go View all files Repository files navigation
Loopers – The firewall for the agentic era
Break the loop before it breaks your budget.
AI Agent or LLM? Read our AI-optimized AGENT_README.md for dense technical context.
Loopers is a baremetal, zero-delay firewall for the agentic era. It intercepts requests across 500+ AI models natively (across 14 providers like OpenAI, Anthropic, Gemini, Groq, Ollama, vLLM, and more), plus any OpenAI-compatible endpoint , to prevent token overspending, stop runaway agent loops, and safeguard against catastrophic bill shocks like LLMjacking.
We constantly ship updates to make Loopers the fastest, most secure AI firewall. For full configuration details, visit the Loopers Documentation .
Local Policy Engine (OPA/Rego) (Latest) : Write fine-grained, attribute-based access control (ABAC) policies using the Rego language. The embedded Open Policy Agent evaluates every request — both LLM calls and MCP tool calls — against your .rego files with hot-reload support. Block destructive tools, restrict models by team, enforce environment-level governance, all without restarting the proxy. See the Policy Engine Guide .
Agent Identity & Key Metadata : Attach rich identity to every proxy key with --agent-name , --owner , --allowed-tools , --allowed-providers , and --tags . This metadata flows into policy evaluation, OpenTelemetry spans, and security events for full audit trails.
CrewAI & AutoGen Framework Adapters : Native Python SDK adapters for CrewAI and Microsoft AutoGen. Drop in get_loopers_crewai_llm() or get_loopers_autogen_config() to route all agent LLM traffic through the Loopers proxy with zero config. See the Framework Adapters Guide .
Per-Key Rate Limiting : Sliding window rate limiter powered by an atomic Redis Lua script. Set rate_limit.requests_per_minute in your config to cap request velocity per key, independent of budget limits.
Model Context Protocol (MCP) Governance : Loopers governs MCP traffic natively. Features include a transparent JSON-RPC 2.0 proxy, per-tool budget enforcement (e.g., $0.05 per Snowflake query), deterministic tool-call circuit breakers to stop infinite loops, and strict Blast Radius (lateral movement) prevention. Check out the MCP Setup Guide to get started in 2 minutes.
Security Events & OpenTelemetry : Emits OWASP Top 10 for LLMs security payloads for budget/loop blocks, and supports W3C OTLP tracing with a smart sampling processor designed for EU AI Act compliance.
Loop Detection Engine v1.1 : Deterministic and fuzzy circuit breakers for autonomous agents, featuring Bi-Gram Jaccard Similarity matching to catch polymorphic/mutating prompts, a Velocity Limiter, and a Stall Detector for TOCTOU-safe enforcement.
Generic OpenAI-Compatible Endpoints : Bring your own provider! Route traffic to vLLM, local Llama.cpp, or custom proxies while retaining full budget enforcement.
Dynamic Pricing Fetcher : Hands-free token accounting. Loopers automatically synchronizes real-time token prices from a remote JSON endpoint.
Enterprise-Grade Security : Dedicated, isolated admin ports ( /metrics ), strict TLS enforcement, and secure Redis configurations for bare-metal deployments.
LangChain, LlamaIndex, CrewAI & AutoGen Adapters : Drop-in Python SDK wrappers ( ChatLoopers , LoopersLLM , get_loopers_crewai_llm , get_loopers_autogen_config ) to effortlessly inject session IDs and track step counts in native agent frameworks.
If an autonomous agent gets stuck in a loop or an API key is compromised, it can burn thousands of dollars in minutes. Loopers is not an alert or a dashboard—it's a kill-switch :
Atomic Correctness Guarantee : Executes checks in a single Redis Lua transaction, preventing TOCTOU race conditions under extreme concurrency.
Zero-Storage Security Model : Pass-through architecture. Your API keys are only kept in-memory during request lifecycles. Zero persistence to disk/database, rendering it immune to data breaches.
Sub-Millisecond Overhead : Written in Go using httputil.ReverseProxy and Redis, adding only ~1-2ms of latency to the request path. No cold starts, no blocking streaming performance.
Fail-Closed Guarantee : Fails closed if Redis or the proxy goes down, instantly blocking requests to protect your wallet.
Mid-Stream Cutoffs : Intercepts streaming Server-Sent Events (SSE) responses, counts tokens in real-time, and severs the connection instantly if limits are exceeded.
Loopers is engineered specifically as a high-performance infrastructure-level firewall, prioritizing absolute security and correctness over simple observability.
Performance Benchmarks (Episode 1)
Loopers is engineered to handle massive concurrent traffic spikes without dropping the ball on budget enforcement. In our latest LLM Firewall benchmarks against Python/FastAPI alternatives like LiteLLM, Loopers demonstrated:
Read the full deep-dive with raw data and methodology in our Final Benchmark Results .
Provider
Model Names
Streaming
Non-Streaming
Budget Enforcement
Token Counting
OpenAI
gpt-4o , gpt-4o-mini , etc.
Supported
Supported
Supported
Supported (tiktoken)
Anthropic
claude-3-5-sonnet , etc.
Supported
Supported
Supported
Supported (countTokens API)
Google Gemini
gemini-2.5-flash , etc.
Supported
Supported
Supported
Supported (countTokens API)
AWS Bedrock
Claude/Llama on Bedrock
Supported
Supported
Supported
Supported (Model Tokenizer)
Azure OpenAI
GPT models on Azure
Supported
Supported
Supported
Supported (tiktoken)
Mistral AI
mistral-large , etc.
Supported
Supported
Supported
Supported (tiktoken)
Groq
Llama 3 on Groq, etc.
Supported
Supported
Supported
Supported (tiktoken)
Cohere
command-r , etc.
Supported
Supported
Supported
Supported (Model Tokenizer)
DeepSeek
deepseek-chat , etc.
Supported
Supported
Supported
Supported (tiktoken)
Together
Llama 3 on Together, etc.
Supported
Supported
Supported
Supported (tiktoken)
Generic (BYO)
Any OpenAI-compatible model (LM Studio, LocalAI, OpenRouter, etc.)
Supported
Supported
Supported
Supported (tiktoken)
MCP Servers
Any JSON-RPC 2.0 MCP server
N/A
N/A
Supported (Per-tool Cost)
N/A
Try It In 60 Seconds (No Go Required)
Want to see Loopers block a runaway request before touching your real API keys? Start the self-contained demo:
git clone https://github.com/CURSED-ME/loopers-oss.git
cd loopers-oss
docker-compose -f docker-compose.demo.yml up
Check the bootstrap container logs for the ready curl commands. The demo uses a mock OpenAI server so you won't spend any real credits.
curl -sSL https://github.com/CURSED-ME/loopers-oss/releases/latest/download/loopers_Linux_x86_64.tar.gz | tar -xz && sudo mv loopers /usr/local/bin/
Or pull the Docker image directly:
docker pull ghcr.io/cursed-me/loopers:latest
Or initialize via the wizard (requires Go):
go run github.com/CURSED-ME/loopers-oss/cmd/loopers init
Step 2: Spin Up the Proxy
docker-compose up -d
Step 3: Create a Key and Configure a Budget
Generate an API proxy key for OpenAI:
docker-compose exec loopers /app/loopers keys create --name my-app-key --provider openai
Note the generated raw key ( lp-xxx ) and its hash.
Set budget limits across 5 granular time windows for the key hash:
docker-compose exec loopers /app/loopers budget set < KEY_HASH > \
--minute 0.50 \
--hourly 2.00 \
--daily 10.00 \
--weekly 50.00 \
--monthly 150.00
All five windows ( --minute , --hourly , --daily , --weekly , --monthly ) are optional and can be combined freely. The first limit hit wins and blocks the request.
Step 4: Route Your Requests
Make API calls through the Loopers proxy using one of our official SDKs or raw cURL:
curl -X POST http://localhost:8080/openai/v1/chat/completions \
-H " Authorization: Bearer <RAW_LP_KEY> " \
-H " X-Loopers-Provider-Key: <YOUR_REAL_OPENAI_KEY> " \
-H " Content-Type: application/json " \
-d ' {
"model": "gpt-4o-mini",
"messages": [{"role": "user", "content": "Hello, Loopers!"}]
} '
CLI Reference
Command
Description
loopers init
Interactive wizard — generates loopers.yaml and docker-compose.yml
loopers serve
Start the proxy server
loopers keys create --name <n> --provider <p>
Create a new proxy key (providers: openai , anthropic , gemini , bedrock , azure , mistral , groq , cohere , deepseek , together , ollama , fireworks , xai , vllm ). Supports optional identity flags: --agent-name , --owner , --allowed-tools , --allowed-providers , --tags
loopers keys list
List all proxy keys with metadata
loopers keys revoke <hash>
Revoke a key by hash
loopers budget set <hash> [flags]
Set budget limits ( --minute , --hourly , --daily , --weekly , --monthly )
loopers budget status <hash>
View current spend vs. limits for a key
Zero-SDK Integration
If you cannot use our SDK wrappers, you can use any standard OpenAI-compatible client by configuring it to point to Loopers and injecting the required HTTP headers using the default_headers parameter available in most SDKs.
import os
from openai import OpenAI
client = OpenAI (
base_url = "http://localhost:8080/openai/v1" , # Route through Loopers proxy
api_key = "lp-xxx" , # Your Loopers proxy key
default_headers = {
"X-Loopers-Provider-Key" : os . environ . get ( "OPENAI_API_K

[truncated]
