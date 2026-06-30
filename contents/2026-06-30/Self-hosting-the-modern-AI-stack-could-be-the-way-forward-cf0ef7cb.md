---
source: "https://github.com/raiyanyahya/llmaker"
hn_url: "https://news.ycombinator.com/item?id=48726917"
title: "Self hosting the modern AI stack could be the way forward"
article_title: "GitHub - raiyanyahya/llmaker: Self-host the modern LLM stack. · GitHub"
author: "sleepynoodle"
captured_at: "2026-06-30T00:30:43Z"
capture_tool: "hn-digest"
hn_id: 48726917
score: 3
comments: 0
posted_at: "2026-06-29T23:58:53Z"
tags:
  - hacker-news
  - translated
---

# Self hosting the modern AI stack could be the way forward

- HN: [48726917](https://news.ycombinator.com/item?id=48726917)
- Source: [github.com](https://github.com/raiyanyahya/llmaker)
- Score: 3
- Comments: 0
- Posted: 2026-06-29T23:58:53Z

## Translation

タイトル: 最新の AI スタックを自己ホストすることが前進の可能性がある
記事のタイトル: GitHub - raiyanyahya/llmaker: 最新の LLM スタックを自己ホストします。 · GitHub
説明: 最新の LLM スタックを自己ホストします。 GitHub でアカウントを作成して、raiyanyahya/llmaker の開発に貢献してください。

記事本文:
GitHub - raiyanyahya/llmaker: 最新の LLM スタックを自己ホストします。 · GitHub
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
ライヤニャヒヤ
/
イルメーカー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルへ移動 コード

「その他のアクション」メニューを開く フォルダーとファイル
35 コミット 35 コミット .github/ workflows .github/ workflows エージェント エージェント cmd/ llmaker cmd/ llmaker ドキュメント ドキュメントの例 例 ファサード ファサード 画像 画像 内部 内部スクリプト スクリプト .dockerignore .dockerignore .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
最新の LLM スタックを自己ホストします。
llmaker は、完全な最新の LLM スタックを実行するためのオープンソース プラットフォームです。
独自のインフラストラクチャ - 大規模な言語モデル、ベクトル データベース、埋め込み、
キャッシュ、可観測性、組み込みの取得層とエージェント層 - プロビジョニングされ、
ネットワーク化され、単一のコマンドから実稼働環境が形成されます。
プライベート検索を強化したチャットボット、FAQ アシスタント、およびレコメンデーションを構築する
エンジンは現地で。サードパーティの API キーはありません。マシンからデータが流出することはありません。
クイックスタート · llmaker を選ぶ理由 · スタック · エージェント · アーキテクチャ · CLI · ロードマップ
モデルをローカルで実行するのは簡単です。アプリケーションの出荷は行われません。プロダクション
検索システムにはベクトル データベース、埋め込みサービス、キャッシュ層が必要です。
オーケストレーション層と可観測性 - それぞれがコンテナ化され、ネットワーク化され、
他のものを検出するように構成されています。それが経常税であることをまとめると、
docker run フラグ、脆弱な Compose ファイル、および数百行のフレームワーク接着剤。
llmaker はその税金を取り除きます。 1 つの CLI でスタック全体をプライベート ネットワーク上にプロビジョニングします
ライブステータス、ログ、リソースとして単一のフリートとして運用します。
すべてのモデルとサービスにわたるダッシュボード。スタックは宣言的であり、
調整可能 ( apply --prune )、モデルは OpenAI 互換、および取得
箱から出してトレースされます。単一モデルから完全なアプリケーションまで:
# ── 完全なアプリケーション スタックを構築する ──

───────────────────
llmaker スタックアップ アシスタント # 1 つのコマンド → ローカル モデル上のプライベート ChatGPT スタイルの UI
llmaker stack init rag # …または編集するスタックをスキャフォールディングしてから適用します。
llmaker 適用 # アシスタント · 音声 · ラグ · リサーチ · コード · チャットボット · よくある質問 · 推奨 · SQL
# ── …または単一モデルを実行する (OpenAI 互換) ────────
llmaker up --model llama3:8b # ローカル エンドポイント — 明示的、またはプリセット:
llmaker アップチャット # チャット · コード · 小さな · 埋め込み · ビジョン
llmaker chat < name > # ターミナルでテストします
llmaker open < name > # 組み込みの Web UI を開きます
# ── …またはサービスごとにスタックをアラカルトで構成する ───
llmaker サービス カタログ # 利用可能なものを参照する
llmaker サービス qdrant # ベクトル データベースを追加 → qdrant:6333
llmaker サービスが redis # キャッシュ / メモリを追加 → redis:6379
llmaker サービス追加 langfuse # 可観測性 → langfuse:3000
# ── フリートの運用 ──── ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─
llmaker ls # すべてのモデル + サービス、1 つのビュー (--json)
llmaker トップ # ライブ リソース ダッシュボード (TUI)
llmaker status <名前> # ゲージ、ロードされたモデル、エンドポイント
llmaker logs < name > -f # 任意のコンテナからログをストリーミングします
llmaker pull misstral --on chat # 進行状況に応じてモデルをダウンロードする
llmaker 停止 / 開始 / rm # ライフサイクル管理
# ── エージェントの API または任意の OpenAI クライアントを使用します ──────
AGENT= $( llmaker service ls --json | jq -r ' .[]|select(.service=="agent").url ' )
curl " $AGENT /api/ingest " -F file=@handbook.pdf # 知識を追加
カール " $AGENT /api/chat " -d ' {"質問":"返金ポリシー?"} ' # grounde

d 答え + ソース
curl " $AGENT /api/recommend " -d ' {"like":["sku1","sku2"]} ' # セマンティックな推奨事項
すべてはプライベート ネットワーク上に配置され、そこで各コンテナが他のコンテナを検出します。
名前による - Compose ファイルやグルー コードはありません。
厳選された完全なスタック
モデルとその周囲のインフラストラクチャ (ベクトル データベース (Qdrant、Chroma、pgvector、Weaviate)、Redis、埋め込み、Open WebUI、n8n、Flowise、Whisper、Langfuse) を 1 つのバージョン管理されたカタログから取得します。
自動サービス検出
すべてのモデルとサービスはプライベート Docker ネットワークに参加し、名前によって解決されます。アプリケーションは IP 配線なしで chat:8080 および qdrant:6333 に到達します。
検索およびツールエージェントが組み込まれています
FastAPI + LangGraph サービス: 書き換え → 取得 → 再ランク付け → 生成 (マルチターン、MMR)、ツール呼び出しループ (計算機、ナレッジ ベース、自己ホスト型 Web 検索、SQL)、およびセマンティック推奨 API。
デフォルトでの可観測性
すべてのインスタンスはスクレイピング用に Prometheus /metrics (リクエスト、トークン/秒、CPU/RAM/GPU) を公開し、RAG スタックには Langfuse が同梱されています。セットアップなしですべてのクエリ (検索ヒットとスコア、モデルとトークンの使用状況) がトレースされます。
測定可能な品質
評価ハーネス ( /api/eval ) は、回答の根拠性、関連性、正確性を LLM 判定者によって評価します。取得品質は、推測ではなく変更全体を追跡できます。
RAG以上のもの
要約 (長いドキュメントのマップリデュース)、構造化 JSON 抽出、音声テキスト変換 (Whisper) のためのファーストクラスのエンドポイントに加え、オプションの Redis を利用した会話メモリ。
宣言的、調和的
スタックを 1 つのファイルで定義します。 llmaker apply は、依存関係の順序で目的の状態にします。 --prune は宣言されなくなったものを削除します。
OpenAI対応
各モデルは、1 つのコントラクトの背後で安定した /v1/* API (チャット、補完、埋め込み、ストリーミング) を公開します — Oll

現在、ama はこれを実行しており、llama.cpp バックエンドが進行中です。
設計によりプライベート化
コンテナーはデフォルトで 127.0.0.1 にバインドされます。ドキュメント、埋め込み、トレースがインフラストラクチャから離れることはありません。トークンごとのコストやベンダー ロックインはありません。
操作可能
単一の静的 Go バイナリ、ドリフトする状態ファイルのないラベル付きコンテナ モデル、あらゆる場所に --json 出力、ライブ トップ ダッシュボード。
LLM スタックをセルフホストする理由は何ですか?
データの所有権。独自の文書、顧客データ、プロンプトはそのまま残ります
あなたが制御するハードウェア。サードパーティ API には何も送信されません。
組み立て税はかかりません。ベクター DB、エンベディング、キャッシュ、エージェント、トレースが付属
手動で管理する Compose ファイルとしてではなく、事前に統合され、ネットワーク化されています。
予測可能なコスト。推論と取得はすでに構築されているインフラストラクチャ上で実行されます
支払う。トークンごとの課金やレート制限はありません。
携帯性。同じ stack.yaml がラップトップ、CI ランナー、または
サーバー。アプリケーションに触れることなく、モデルまたはベクトル データベースを交換します。
Docker が必要です。その後、llmaker Doctor を実行して環境を検証します。
# ビルド済みバイナリ (Linux / macOS)
カール -fsSL https://raw.githubusercontent.com/raiyanyahya/llmaker/master/scripts/install.sh |しー
# Go ツールチェーン
github.com/raiyanyahya/llmaker/cmd/llmaker@latest をインストールしてください
# ソースから
git clone https://github.com/raiyanyahya/llmaker && cd llmaker && make build
Homebrew および winget パッケージはロードマップにあります。エージェント イメージは、レジストリに公開されるまで、make image-agent を使用してローカルに構築されます。
完全な取得拡張生成スタックをプロビジョニングして実行します。
llmaker スタックアップアシスタント # scaffold + 1 ステップで適用 (アシスタントにはエージェントイメージは必要ありません)
llmaker stack init rag # stack.yaml を生成 (アシスタント | 音声 | ラグ | 研究 | コード | チャットボット | よくある質問 | 推奨 | SQL)
make image-agent # エージェントイメージを構築します

e (エージェントを含むスタック)
llmaker apply -f stack.yaml # スタックをプロビジョニングします — モデル + サービス、ネットワーク化
llmaker ls # モデルとサービスを 1 つのビューで検査する
エージェント エンドポイントを解決して使用します。
AGENT= $( llmaker service ls --json | jq -r ' .[] | select(.service=="agent").url ' )
curl " $AGENT /api/ingest " -F file=@handbook.pdf # ドキュメントを取り込む
curl " $AGENT /api/chat " -d ' {"question":"…","history":[],"top_k":4} ' # クエリ、ソースあり
llmaker は個別のモデルも実行します。これは、ローカルのモデルを公開する最も簡単な方法です。
OpenAI 互換エンドポイント:
llmaker up --model llama3:8b # モデル インスタンスをプロビジョニングします
openaiインポートからOpenAI
client = OpenAI (base_url = "http://127.0.0.1:11500/v1" 、api_key = "不要")
クライアント。チャット 。完成品。作成 (モデル = "llama3:8b" ,
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : "Hello" }])
スタック
スタックは、モデルとその周りのサービスが一緒にプロビジョニングされたものです。足場
llmaker stack up <name> を使用して 1 つのステップで 1 つを実行するか、
stack.yaml を llmaker stack init <name> で編集し、それを適用します
llメーカー適用。
カタログのエージェントは、FastAPI + LangGraph サービス (agent/) です。
ベア モデルとベクターはアプリケーションに保存されます。の標準サービスです
ネットワーク。他のネットワークを名前で検出するように環境によって構成されます。
明示的なグラフとしての取得 — 書き換え → 取得 → 再ランク付け → 生成 :
rewrite — マルチターン履歴をスタンドアロンのクエリに折りたたむため、
コンテキスト (「いつリリースされましたか?」) に依存するフォローアップで解決します。
正しく。モデルは、解決する履歴がある場合にのみ呼び出されます。
取得 — クエリを埋め込み、ベクトルから候補セットを取得します
店。
再ランク付け — 最大周辺関連性を適用します
関連性があり、冗長性のないコンテキストに対応します。
生成 — t を生成します

彼はその文脈と会話から答えます。
POST /api/ingest マルチパート ファイルまたはテキスト → チャンク、埋め込み、保存
POST /api/chat { 質問、履歴?、top_k?、session_id? } → 答え + ソース
POST /api/agent { 質問、履歴?、セッション ID? } → ツールを使用した応答 + ツール呼び出し
POST /api/summarize { テキスト、指示?、最大単語数? } → 概要 (長いテキストの場合はマップリデュース)
POST /api/extract { text,fields: { name: description } } → まさにそれらのキーを含む JSON
POST /api/transcribe マルチパート オーディオ ファイル → { text } (ウィスパー サービスが必要)
POST /api/eval { ケース: [{ 質問、参照? }] } → 採点された解答 + 概要
POST /api/items { items: [{ ID、テキスト、メタデータ? }] } → 推奨事項のインデックス
POST /api/recommend { query } or { like: [id, …] } → ランク付けされたアイテム
ツール呼び出し。取得を超えて、/api/agent はツール呼び出しループを実行します。
モデルは、どのツール (計算機、ナレッジベース) を呼び出すかを決定します。
(ツールとしての取得)、現在時刻、自己ホスト型 Web 検索
(SearXNG、有料 API なし)、およびオプションの読み取り専用 SQL ツールを使用できます。
データベース — そしてループは答えが得られるまでそれらを実行します。応答
実行されたすべてのツール呼び出しが含まれます。ツールの追加は 1 つのエントリです
エージェント/app/tools.py 。
トレース。ラグ スタックは Langfuse をプロビジョニングし、エージェントはすべてのクエリをトレースします。
設定なしで、各リクエスト (RAG またはツール使用) が
取得、ツール、生成のステップを使用してトレースします。トレースは、
テンプレートを使用し、それ以外の場合は 2 つの環境変数を介してオプトインされます。
評価。 /api/eval が実行される

[切り捨てられた]

## Original Extract

Self-host the modern LLM stack. Contribute to raiyanyahya/llmaker development by creating an account on GitHub.

GitHub - raiyanyahya/llmaker: Self-host the modern LLM stack. · GitHub
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
raiyanyahya
/
llmaker
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .github/ workflows .github/ workflows agent agent cmd/ llmaker cmd/ llmaker docs docs examples examples facade facade images images internal internal scripts scripts .dockerignore .dockerignore .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Self-host the modern LLM stack.
llmaker is an open-source platform for running the complete modern LLM stack on
your own infrastructure — large language models, vector databases, embeddings,
caching, observability, and a built-in retrieval & agent layer — provisioned,
networked, and production-shaped from a single command.
Build private retrieval-augmented chatbots, FAQ assistants, and recommendation
engines locally. No third-party API keys. No data leaving your machine.
Quickstart · Why llmaker · Stacks · The agent · Architecture · CLI · Roadmap
Running a model locally is easy. Shipping an application is not. A production
retrieval system needs a vector database, an embeddings service, a caching layer,
an orchestration layer, and observability — each containerized, networked, and
configured to discover the others. Assembling that is a recurring tax: a sprawl of
docker run flags, a brittle Compose file, and hundreds of lines of framework glue.
llmaker removes that tax. One CLI provisions the entire stack on a private network
and operates it as a single fleet — live status, logs, and a resource
dashboard across every model and service. Stacks are declarative and
reconcilable ( apply --prune ), models are OpenAI-compatible , and retrieval
is traced out of the box . From a single model to a complete application:
# ── Build a complete application stack ──────────────────────────
llmaker stack up assistant # one command → a private ChatGPT-style UI over a local model
llmaker stack init rag # …or scaffold any stack to edit, then apply it:
llmaker apply # assistant · voice · rag · research · code · chatbot · faq · recommend · sql
# ── …or run a single model (OpenAI-compatible) ──────────────────
llmaker up --model llama3:8b # a local endpoint — explicit, or a preset:
llmaker up chat # chat · code · small · embed · vision
llmaker chat < name > # test it in the terminal
llmaker open < name > # open its built-in web UI
# ── …or compose the stack à la carte, service by service ────────
llmaker service catalog # browse what's available
llmaker service add qdrant # vector database → qdrant:6333
llmaker service add redis # cache / memory → redis:6379
llmaker service add langfuse # observability → langfuse:3000
# ── Operate the fleet ───────────────────────────────────────────
llmaker ls # every model + service, one view (--json)
llmaker top # live resource dashboard (TUI)
llmaker status < name > # gauges, loaded models, endpoints
llmaker logs < name > -f # stream logs from any container
llmaker pull mistral --on chat # download a model with progress
llmaker stop / start / rm # lifecycle management
# ── Consume it — the agent's API, or any OpenAI client ──────────
AGENT= $( llmaker service ls --json | jq -r ' .[]|select(.service=="agent").url ' )
curl " $AGENT /api/ingest " -F file=@handbook.pdf # add knowledge
curl " $AGENT /api/chat " -d ' {"question":"refund policy?"} ' # grounded answer + sources
curl " $AGENT /api/recommend " -d ' {"like":["sku1","sku2"]} ' # semantic recommendations
Everything lands on a private network where each container discovers the others
by name — no Compose file and no glue code.
The complete stack, curated
Models and the infrastructure around them — vector databases (Qdrant, Chroma, pgvector, Weaviate), Redis, embeddings, Open WebUI, n8n, Flowise, Whisper, Langfuse — from one versioned catalog.
Automatic service discovery
Every model and service joins a private Docker network and resolves by name. Your application reaches chat:8080 and qdrant:6333 with zero IP wiring.
A retrieval & tool agent, built in
A FastAPI + LangGraph service: rewrite → retrieve → rerank → generate (multi-turn, MMR), a tool-calling loop (calculator, knowledge base, self-hosted web search, SQL), and a semantic recommendation API.
Observability by default
Every instance exposes Prometheus /metrics (requests, tokens/sec, CPU/RAM/GPU) for scraping, and the RAG stack ships Langfuse — every query traced (retrieval hits and scores, model and token usage) with no setup.
Measurable quality
An evaluation harness ( /api/eval ) grades answers for groundedness, relevance, and correctness with an LLM judge — retrieval quality you can track across changes, not guess at.
More than RAG
First-class endpoints for summarization (map-reduce over long docs), structured JSON extraction, and speech-to-text (Whisper), plus optional Redis-backed conversation memory.
Declarative, reconcilable
Define your stack in one file. llmaker apply brings it to the desired state in dependency order; --prune removes what's no longer declared.
OpenAI-compatible
Each model exposes a stable /v1/* API (chat, completions, embeddings, streaming) behind one contract — Ollama runs it today, with a llama.cpp backend in progress .
Private by design
Containers bind to 127.0.0.1 by default. Your documents, embeddings, and traces never leave your infrastructure. No per-token cost, no vendor lock-in.
Operable
A single static Go binary, a labeled-container model with no state file to drift, --json output everywhere, and a live top dashboard.
Why self-host your LLM stack?
Data ownership. Proprietary documents, customer data, and prompts stay on
hardware you control. Nothing is sent to a third-party API.
No assembly tax. The vector DB, embeddings, cache, agent, and tracing come
pre-integrated and networked — not as a Compose file you maintain by hand.
Predictable cost. Inference and retrieval run on infrastructure you already
pay for. No per-token billing, no rate limits.
Portability. The same stack.yaml runs on a laptop, a CI runner, or a
server. Swap the model or the vector database without touching your application.
Requires Docker . Run llmaker doctor afterward to validate your environment.
# Prebuilt binary (Linux / macOS)
curl -fsSL https://raw.githubusercontent.com/raiyanyahya/llmaker/master/scripts/install.sh | sh
# Go toolchain
go install github.com/raiyanyahya/llmaker/cmd/llmaker@latest
# From source
git clone https://github.com/raiyanyahya/llmaker && cd llmaker && make build
Homebrew and winget packages are on the roadmap . The agent image is built locally with make image-agent until it is published to a registry.
Provision and run a complete retrieval-augmented generation stack:
llmaker stack up assistant # scaffold + apply in one step (assistant needs no agent image)
llmaker stack init rag # generate stack.yaml (assistant | voice | rag | research | code | chatbot | faq | recommend | sql)
make image-agent # build the agent image once (stacks that include the agent)
llmaker apply -f stack.yaml # provision the stack — model + services, networked
llmaker ls # inspect models and services in one view
Resolve the agent endpoint and use it:
AGENT= $( llmaker service ls --json | jq -r ' .[] | select(.service=="agent").url ' )
curl " $AGENT /api/ingest " -F file=@handbook.pdf # ingest documents
curl " $AGENT /api/chat " -d ' {"question":"…","history":[],"top_k":4} ' # query, with sources
llmaker also runs individual models — the easiest way to expose a local,
OpenAI-compatible endpoint:
llmaker up --model llama3:8b # provision a model instance
from openai import OpenAI
client = OpenAI ( base_url = "http://127.0.0.1:11500/v1" , api_key = "not-needed" )
client . chat . completions . create ( model = "llama3:8b" ,
messages = [{ "role" : "user" , "content" : "Hello" }])
Stacks
A stack is a model plus the services around it, provisioned together. Scaffold
and run one in a single step with llmaker stack up <name> , or generate a
stack.yaml to edit with llmaker stack init <name> and apply it with
llmaker apply .
The catalog's agent is a FastAPI + LangGraph service ( agent/ ) that turns a
bare model and vector store into an application. It is a standard service on the
network, configured by environment to discover the others by name.
Retrieval as an explicit graph — rewrite → retrieve → rerank → generate :
rewrite — collapses multi-turn history into a standalone query, so
follow-ups that depend on context ("and when was it released?") resolve
correctly. The model is only invoked when there is history to resolve.
retrieve — embeds the query and retrieves a candidate set from the vector
store.
rerank — applies Maximal Marginal Relevance
for relevant, non-redundant context.
generate — produces the answer from that context and the conversation.
POST /api/ingest multipart file or text → chunk, embed, store
POST /api/chat { question, history?, top_k?, session_id? } → answer + sources
POST /api/agent { question, history?, session_id? } → tool-using answer + tool calls
POST /api/summarize { text, instructions?, max_words? } → summary (map-reduce for long text)
POST /api/extract { text, fields: { name: description } } → JSON with exactly those keys
POST /api/transcribe multipart audio file → { text } (needs a whisper service)
POST /api/eval { cases: [{ question, reference? }] } → graded answers + summary
POST /api/items { items: [{ id, text, metadata? }] } → index for recommendation
POST /api/recommend { query } or { like: [id, …] } → ranked items
Tool calling. Beyond retrieval, /api/agent runs a tool-calling loop where
the model decides which tools to invoke — a calculator , the knowledge base
(retrieval as a tool), the current time , a self-hosted web search
(SearXNG, no paid API), and an optional read-only SQL tool over your
database — and the loop executes them until it has an answer. The response
includes every tool call it made. Adding a tool is one entry in
agent/app/tools.py .
Tracing. The rag stack provisions Langfuse and the agent traces every query
to it, with zero configuration — each request (RAG or tool-using) appears as a
trace with its retrieval, tool, and generation steps. Tracing is enabled by the
template and is otherwise opt-in via two environment variables.
Evaluation. /api/eval runs

[truncated]
