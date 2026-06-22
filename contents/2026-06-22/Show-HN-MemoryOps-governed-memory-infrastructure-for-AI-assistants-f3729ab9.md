---
source: "https://github.com/patibandlavenkatamanideep/memoryops-ai"
hn_url: "https://news.ycombinator.com/item?id=48626034"
title: "Show HN: MemoryOps – governed memory infrastructure for AI assistants"
article_title: "GitHub - patibandlavenkatamanideep/memoryops-ai: Enterprise-shaped memory governance layer for AI assistants: typed capture, policy evaluation, hybrid retrieval, deletion guarantees, tenant isolation, provenance, and audit. · GitHub"
author: "pvmanideep20"
captured_at: "2026-06-22T05:46:58Z"
capture_tool: "hn-digest"
hn_id: 48626034
score: 2
comments: 0
posted_at: "2026-06-22T05:26:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: MemoryOps – governed memory infrastructure for AI assistants

- HN: [48626034](https://news.ycombinator.com/item?id=48626034)
- Source: [github.com](https://github.com/patibandlavenkatamanideep/memoryops-ai)
- Score: 2
- Comments: 0
- Posted: 2026-06-22T05:26:04Z

## Translation

タイトル: Show HN: MemoryOps – AI アシスタント向けの管理されたメモリ インフラストラクチャ
記事のタイトル: GitHub - patibandlavenkatamanideep/memoryops-ai: AI アシスタント向けのエンタープライズ型メモリ ガバナンス レイヤー: 型付きキャプチャ、ポリシー評価、ハイブリッド取得、削除保証、テナント分離、来歴、監査。 · GitHub
説明: AI アシスタント向けのエンタープライズ型メモリ ガバナンス レイヤー: 型付きキャプチャ、ポリシー評価、ハイブリッド取得、削除保証、テナント分離、来歴、監査。 - patibandlavenkatamanideep/memoryops-ai
HN テキスト: 削除圧縮、ベクターパージ検証、テナント分離、監査証拠を備えた AI アシスタント用の管理されたメモリ層を構築しました。

記事本文:
GitHub - patibandlavenkatamanideep/memoryops-ai: AI アシスタント向けのエンタープライズ型メモリ ガバナンス レイヤー: 型付きキャプチャ、ポリシー評価、ハイブリッド取得、削除保証、テナント分離、来歴、監査。 · GitHub
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
パティバン

ドラベンカタマニディープ
/
メモリーオプス-ai
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
patibandlavenkatamanideep/memoryops-ai
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット .github/ workflows .github/ workflows .hermes/ skill .hermes/ skill apps/ web apps/ web docs docs evals evals infra infra 鉄道スクリプト スクリプトサービス services .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CLAUDE_ENTERPRISE.md CLAUDE_ENTERPRISE.md CONTRIBUTING.md CONTRIBUTING.md README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
MemoryOps AI は、AI アシスタント向けのエンタープライズ型のループエンジニアリングされたメモリ ガバナンス レイヤーです。
キャプチャ、ポリシー評価、型付きストレージを備えた ChatGPT スタイルのメモリ ライフサイクルを実装します。
ハイブリッド検索、忘却の制御、監査可能性、およびテナントの分離。
ほとんどのデモでは、メモリをベクトル データベースとして扱います。 MemoryOps AI はメモリを管理された状態として扱います。
キャッチフレーズ: AI アシスタントのためのエンタープライズ メモリ ガバナンス。
主な主張: メモリはデータベースではありません。記憶は、何を決定するかを決定する管理された意思決定システムです。
情報は将来に引き継ぐのに十分な価値があります。
ほとんどの AI の「メモリ」デモでは次のことが行われます。
チャットメッセージ → ベクターデータベース → 後で取得
MemoryOps AI は次のことを行います。
書き込みパス
メッセージ → エクストラクター → エバリュエーター / ポリシー ブローカー → 書き込みサービス → 型付きメモリ ストア → 監査ログ
読み取りパス
メッセージ → レトリーバー → ランカー → コンテキストコンポーザー → 応答 LLM
背景
ディケイジョブ→リフレクションエージェント→コンフリクトリゾルバー→圧縮ワーカー
横断平面
セキュリティ、ガバナンス、可観測性、評価

使用・信頼性
システムが実証する必要がある 5 つの動詞:
キャプチャ → 保存 → 取得 → 更新 → 忘れ (ガバナンスは 5 つすべてをラップします)
フローチャート LR
M["チャット メッセージ"] --> GW["ゲートウェイ"]
GW --> EX["エクストラクター"] --> PB["ポリシーブローカー"] --> WS["書き込みサービス"] --> ST[("型付きストア")]
GW --> RT["Retriever"] --> RK["Ranker"] --> CC["Context Composer"] --> RESP["Response"]
PB --> AUD[["監査ログ (追加のみ)"]]
WS --> オーストラリアドル
ST-。背景 .-> BG["減衰・反射・葛藤・圧縮"]
読み込み中
その他の図 (システム アーキテクチャ、ライフサイクル ステート マシン、リクエスト シーケンス) は次のとおりです。
docs/architecture.md にあります。
これらは交渉の余地がなく、コードとテストで強制されます。
テナントの分離 — ユーザー A のメモリがユーザー B または別のテナントに返されることはありません。
削除保証 — 削除された記憶は二度と復元されません。
来歴 — 保存されたすべての記憶は、そのソース メッセージ/文書/手入力にまで遡ります。
グレースフル デグラデーション — 取得の失敗によって応答の生成がブロックされることはありません。
Policy-before-storage — 安全でないコンテンツや秘密のようなコンテンツは、ストアに到達する前にフィルタリングされます。
一時チャット — 一時セッションではメモリの書き込みや取得は行われません。
監査可能性 — すべてのメモリ ライフサイクル イベントは、追加のみの監査イベントを生成します。
説明可能性 — システムは、どの記憶が反応に影響を与えたかを示すことができます。
型付き記憶 — エピソード記憶、意味記憶、手順記憶、プロジェクト記憶、知識記憶、システム記憶は異なります。
評価 — メモリの品質は、手動検査だけでなく、ゴールデン セットを通じてテストできます。
完全な設計と各不変条件の場所については、docs/architecture.md を参照してください。
施行された。
メモリオプス-ai/
apps/web/ Next.js フロントエンド (チャット、思い出、ガバナンス、監査、ループ、管理、アーキテクチャ)
services/api/ FastAPI バックエンド (ゲートウェイ、エクストラクター、ポリシー ブローカー、書き込み/読み取りパス、監査)
サービス/を

rker/ バックグラウンド ジョブ (減衰、反映、競合解決、圧縮)
パッケージ/共有/共有タイプ
infra/db/Postgres + pgvector の移行とシード
infra/adr/ アーキテクチャ決定記録
インフラ/可観測性/ OpenTelemetry / メトリクスのメモ
evals/Golden + adversarial ケースと eval ランナー
ドキュメント/アーキテクチャ、セキュリティ、ガバナンス、ロールアウト、デモスクリプト
docker-compose.yml
クイックスタート
オプション A — API のみ、インフラなし (最速)
API にはメモリ内リポジトリが付属しているため、Postgres を使用せずに書き込みパスとテストを実行できます。
CD サービス/API
python -m venv .venv && ソース .venv/bin/activate
pip install -r 要件.txt
エクスポート MEMORYOPS_STORAGE=メモリ # デフォルト;インメモリストアを使用します
uvicorn app.main:app --reload --port 8000
# http://localhost:8000/docs を開きます
不変テスト スイートを実行します。
CD サービス/API
pip install -r 要件-dev.txt
pytest -q
実行中の API (またはインプロセス) に対して評価ハーネスを実行します。
cd 評価
Python run_evals.py
オプション B — Docker Compose を使用したフルスタック
cp .env.example .env
docker 構成 --build
# web → http://localhost:3000
# API → http://localhost:8000/docs
# db → localhost:5432 (postgres/pgvector)
#redis→ローカルホスト:6379
Compose は、最初の起動時に infra/db/migrations から移行を実行し、
MEMORYOPS_STORAGE=API の postgres。
取得には、交換可能な埋め込みプロバイダーが使用されます。デフォルトは決定的です。
オフライン スタブ - API キーは必要ない - ため、テストとデモは再現可能です。
エクスポート MEMORYOPS_EMBEDDING_PROVIDER=スタブ # デフォルト;決定的、キーなし
# オプションの実埋め込み:
エクスポート MEMORYOPS_EMBEDDING_PROVIDER=openai
エクスポート OPENAI_API_KEY=sk-...
エクスポート OPENAI_EMBEDDING_MODEL=text-embedding-3-small
未構成または障害のあるプロバイダーはスタブに低下し、クエリ埋め込みが行われます。
失敗すると検索がキーワードに低下します

-only ( retrieval_mode="fallback" )。
抽出と競合検出はプロバイダー中立の LLM 層を通じて実行されます。
(app/llm/)。デフォルトは確定的なオフライン スタブであり、API キーはありません。
動作は再現可能であり、テストはネットワークに触れることがありません。オプションの OpenAI、
Anthropic アダプターと Gemini アダプターは、キーが設定されている場合にのみ使用されます。
エクスポート MEMORYOPS_LLM_PROVIDER=スタブ # デフォルト;決定的、キーなし
# オプションの実際のプロバイダー (キーが存在する場合にのみ使用されます):
エクスポート MEMORYOPS_LLM_PROVIDER=人類
エクスポート ANTHROPIC_API_KEY=... ANTHROPIC_MODEL=claude-haiku-4-5-20251001
# また: openai (OPENAI_API_KEY/OPENAI_MODEL)、gemini (GEMINI_API_KEY/GEMINI_MODEL)
import MEMORYOPS_LLM_FALLBACK_TO_HEURISTIC=true # 無効な JSON / 失敗 → ヒューリスティック
LLM 出力は勧告です。決定論的ポリシー ブローカーは抽出後に実行されます。
権威を保ちます — モデルはポリシーをオーバーライドすることはできません。
コンテンツはまだブロックされています。 docs/provider-llm-adapters.md を参照してください。
docs/structured-memory-intelligence.md 、
および ADR-008 。
実行中の Postgres に対して行レベルのセキュリティが適用されていることを確認します。
python scripts/check_rls_policies.py # 到達可能な DB がない場合はきれいにスキップします
フロントエンド
CD アプリ/ウェブ
npmインストール
npm run dev # http://localhost:3000
フロントエンドは NEXT_PUBLIC_API_URL (デフォルトは http://localhost:8000 ) を読み取ります。
導入 — 鉄道のみ (v0.3.2)
MemoryOps は鉄道にのみ展開されます。ヴェルセルパスはありません。一本の鉄道
プロジェクト (memoryops-ai) は 5 つのサービスを実行します。
ビルド/デプロイは、 Railway/ の下の config-as-code です。ドキュメント:
docs/deployment/railway.md — トポロジ、順序、ロールバック
docs/deployment/railway-env.md — 環境変数マトリックス
docs/deployment/railway-smoke-test.md — デプロイ後のチェック
Python スクリプト/railway_smoke_test.py \
--api-url https://memoryops-api.up.railway.app \
--web-url https://memoryop

s-web.up.railway.app
現在機能しているもの (フェーズ 0 + フェーズ 1)
完全な設計スパイン: README、アーキテクチャ/セキュリティ/ガバナンス/ロールアウト ドキュメント、5 つの ADR、DB スキーマ。
FastAPI 書き込みパス: Gateway → Extractor → Policy Broker → Write Service → Memory Store → Audit 。
ヒューリスティック抽出機能 + ポリシー ブローカー (API キーなしで動作します);プラグ可能な LLM アダプター インターフェイス。
型付き記憶の分類、重要性/信頼性/感度スコアリング、来歴のキャプチャ。
ポリシーの決定: SAVE 、 PENDING_APPROVAL 、 BLOCK 、 DROP_LOW_UTILITY 、 UPDATE_EXISTING 、 MERGE_WITH_EXISTING 。
シークレット/PII 検出は、保存する前に API キーと認証情報をブロックします。
すべてのライフサイクル イベントに対する追加専用の監査ログ。
一時的なチャットにより、読み取りと書き込みの両方がショートします。
メモリ ダッシュボード + 管理/監査 + アーキテクチャ ページ (フロントエンド スケルトン)。
インバリアント テスト スイート + 評価ハーネス スキャフォールディング。
ループエンジニアリングレイヤー (v0.3.1)
MemoryOps は、受動的なストアではなく、一連の管理されたループとしてメモリをモデル化します。
各ループには、明示的な状態、ポリシー ゲート、監査イベント、フォールバック動作、および
証拠要件。ループ定義は services/api/app/loops/ 、loop にあります
実行/イベントは /api/loops を通じて公開され、フロントエンドには Loops ページが含まれます。
docs/loop-engineering.md を参照してください。
docs/loop-contracts.md 、および
docs/release-loop.md 。
トークン圧縮層 (v0.2.1)
MemoryOps はオプションのヘッドルームをサポートします
コンテキスト圧縮層。圧縮はポリシーチェック、ガバナンスの後に実行されます。
フィルタリング、コンテキスト合成、および合成されたコンテキスト ブロックのみ —
生のユーザーメッセージやポリシーブローカーの前では決してありません。トークンが減ります
MemoryOps の不変条件 (来歴、削除) を保持しながら LLM に送信されます。
保証、テナントの分離、一時的なチャット動作、説明可能性メタデータ）。
デフォルトではオフになっており、depe ではありません

ndency — アプリは何もしなくても実行されます
headroom-ai がインストールされており、圧縮エラーが発生しても安全に分解されます。
圧縮されていないコンテキスト。
pip install " headroom-ai[all] " # オプション
import MEMORYOPS_CONTEXT_COMPRESSION=ヘッドルーム # デフォルト: なし
各チャット応答には、推定トークンが保存された圧縮ブロックが含まれています。
圧縮率。 docs/token-compression.md を参照してください。
docs/integrations/headroom.md 、および
ADR-007 。ヘッドルームは Apache-2.0 です。
MemoryOps はアダプターを介してこれを統合し、そのソースをベンダーしません。
v0.3 で動作するもの (リアルデータレイヤー)
交換可能な埋め込みプロバイダー ( app/embeddings/ ): 決定論的なオフライン スタブ + オプションの OpenAI。
ハイブリッド検索: pgvector cosine ( search_candidates ) + キーワードの重複、ランカーによってブレンドされます。
メモリごとのスコアの内訳 + 応答取得モード (ハイブリッド / フォールバック / なし)。
Postgres の行レベル セキュリティを強制しました (移行 004 、 FORCE + テナント ポリシー + セッション GUC)。
拡張された evals (セマンティック / キーワード / アーカイブ / スコアの内訳) + 新しいテスト。 RLS テストは DB で保護されています。
v0.4 で動作するもの (プロバイダー LLM アダプター)
プロバイダー中立の LLM レイヤー ( app/llm/ ): 決定的な StubProvider のデフォルト +
オプションの OpenAI/Anthropic/Gemini アダプター。 MEMORYOPS_LLM_PROVIDER によって選択されます。
構造化メモリ インテリジェンス : スキーマ検証済みの抽出 + 最小限の競合
検出、プロ付き

[切り捨てられた]

## Original Extract

Enterprise-shaped memory governance layer for AI assistants: typed capture, policy evaluation, hybrid retrieval, deletion guarantees, tenant isolation, provenance, and audit. - patibandlavenkatamanideep/memoryops-ai

I built a governed memory layer for AI assistants with deletion compaction, vector purge verification, tenant isolation, and audit evidence.

GitHub - patibandlavenkatamanideep/memoryops-ai: Enterprise-shaped memory governance layer for AI assistants: typed capture, policy evaluation, hybrid retrieval, deletion guarantees, tenant isolation, provenance, and audit. · GitHub
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
patibandlavenkatamanideep
/
memoryops-ai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
patibandlavenkatamanideep/memoryops-ai
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits .github/ workflows .github/ workflows .hermes/ skills .hermes/ skills apps/ web apps/ web docs docs evals evals infra infra railway railway scripts scripts services services .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CLAUDE_ENTERPRISE.md CLAUDE_ENTERPRISE.md CONTRIBUTING.md CONTRIBUTING.md README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml View all files Repository files navigation
MemoryOps AI is an enterprise-shaped, loop-engineered memory governance layer for AI assistants.
It implements a ChatGPT-style memory lifecycle with capture, policy evaluation, typed storage,
hybrid retrieval, controlled forgetting, auditability, and tenant isolation.
Most demos treat memory as a vector database. MemoryOps AI treats memory as governed state .
Tagline: Enterprise memory governance for AI assistants.
Core claim: Memory is not a database. Memory is a governed decision system that decides what
information is valuable enough to carry into the future.
Most AI "memory" demos do this:
chat message → vector database → retrieve later
MemoryOps AI does this:
WRITE PATH
Message → Extractor → Evaluator / Policy Broker → Write Service → Typed Memory Stores → Audit Log
READ PATH
Message → Retriever → Ranker → Context Composer → Response LLM
BACKGROUND
Decay Job → Reflection Agent → Conflict Resolver → Compression Worker
CROSS-CUTTING PLANES
Security · Governance · Observability · Evaluation · Reliability
The five verbs the system must demonstrate:
Capture → Store → Retrieve → Update → Forget (Governance wraps all five)
flowchart LR
M["chat message"] --> GW["Gateway"]
GW --> EX["Extractor"] --> PB["Policy Broker"] --> WS["Write Service"] --> ST[("Typed Store")]
GW --> RT["Retriever"] --> RK["Ranker"] --> CC["Context Composer"] --> RESP["Response"]
PB --> AUD[["Audit Log (append-only)"]]
WS --> AUD
ST -. background .-> BG["Decay · Reflection · Conflict · Compression"]
Loading
More diagrams (system architecture, lifecycle state machine, request sequence) are
in docs/architecture.md .
These are non-negotiable and are enforced in code and tests.
Tenant isolation — User A's memory is never returned to User B or another tenant.
Deletion guarantee — Deleted memories are never retrieved again.
Provenance — Every stored memory traces back to its source message/document/manual input.
Graceful degradation — Retrieval failure never blocks response generation.
Policy-before-storage — Unsafe / secret-like content is filtered before it reaches the store.
Temporary chat — Temporary sessions never write or retrieve memory.
Auditability — Every memory lifecycle event produces an append-only audit event.
Explainability — The system can show which memories affected a response.
Typed memory — Episodic, semantic, procedural, project, knowledge, system memories differ.
Evaluation — Memory quality is testable through a golden set, not just manual inspection.
See docs/architecture.md for the full design and where each invariant is
enforced.
memoryops-ai/
apps/web/ Next.js frontend (chat, memories, governance, audit, loops, admin, architecture)
services/api/ FastAPI backend (gateway, extractor, policy broker, write/read path, audit)
services/worker/ Background jobs (decay, reflection, conflict resolution, compression)
packages/shared/ Shared types
infra/db/ Postgres + pgvector migrations and seed
infra/adr/ Architecture Decision Records
infra/observability/ OpenTelemetry / metrics notes
evals/ Golden + adversarial cases and the eval runner
docs/ architecture, security, governance, rollout, demo-script
docker-compose.yml
Quickstart
Option A — API only, no infra (fastest)
The API ships with an in-memory repository so you can run the write path and tests without Postgres.
cd services/api
python -m venv .venv && source .venv/bin/activate
pip install -r requirements.txt
export MEMORYOPS_STORAGE=memory # default; uses in-memory store
uvicorn app.main:app --reload --port 8000
# open http://localhost:8000/docs
Run the invariant test suite:
cd services/api
pip install -r requirements-dev.txt
pytest -q
Run the eval harness against a running API (or in-process):
cd evals
python run_evals.py
Option B — Full stack with Docker Compose
cp .env.example .env
docker compose up --build
# web → http://localhost:3000
# api → http://localhost:8000/docs
# db → localhost:5432 (postgres/pgvector)
# redis→ localhost:6379
Compose runs migrations from infra/db/migrations on first boot and sets
MEMORYOPS_STORAGE=postgres for the API.
Retrieval uses a swappable embedding provider. The default is a deterministic,
offline stub — no API key required — so tests and demos are reproducible.
export MEMORYOPS_EMBEDDING_PROVIDER=stub # default; deterministic, no key
# optional real embeddings:
export MEMORYOPS_EMBEDDING_PROVIDER=openai
export OPENAI_API_KEY=sk-...
export OPENAI_EMBEDDING_MODEL=text-embedding-3-small
An unconfigured or failing provider degrades to the stub, and a query-embedding
failure degrades retrieval to keyword-only ( retrieval_mode="fallback" ).
Extraction and conflict detection run through a provider-neutral LLM layer
( app/llm/ ). The default is a deterministic, offline stub — no API key — so
behavior is reproducible and tests never touch the network. Optional OpenAI,
Anthropic, and Gemini adapters are used only when their key is set.
export MEMORYOPS_LLM_PROVIDER=stub # default; deterministic, no key
# optional real providers (used only when the key is present):
export MEMORYOPS_LLM_PROVIDER=anthropic
export ANTHROPIC_API_KEY=... ANTHROPIC_MODEL=claude-haiku-4-5-20251001
# also: openai (OPENAI_API_KEY/OPENAI_MODEL), gemini (GEMINI_API_KEY/GEMINI_MODEL)
export MEMORYOPS_LLM_FALLBACK_TO_HEURISTIC=true # invalid JSON / failure → heuristic
LLM output is advisory : the deterministic policy broker runs after extraction
and stays authoritative — a model can never override policy, and secret-like
content is still blocked. See docs/provider-llm-adapters.md ,
docs/structured-memory-intelligence.md ,
and ADR-008 .
Verify enforced Row-Level Security against a running Postgres:
python scripts/check_rls_policies.py # SKIPs cleanly if no DB is reachable
Frontend
cd apps/web
npm install
npm run dev # http://localhost:3000
The frontend reads NEXT_PUBLIC_API_URL (defaults to http://localhost:8000 ).
Deployment — Railway only (v0.3.2)
MemoryOps deploys to Railway only . There is no Vercel path. One Railway
project ( memoryops-ai ) runs five services:
Build/deploy is config-as-code under railway/ . Docs:
docs/deployment/railway.md — topology, order, rollback
docs/deployment/railway-env.md — env var matrix
docs/deployment/railway-smoke-test.md — post-deploy checks
python scripts/railway_smoke_test.py \
--api-url https://memoryops-api.up.railway.app \
--web-url https://memoryops-web.up.railway.app
What works today (Phase 0 + Phase 1)
Full design spine: README, architecture/security/governance/rollout docs, 5 ADRs, DB schema.
FastAPI write path: Gateway → Extractor → Policy Broker → Write Service → Memory Store → Audit .
Heuristic extractor + policy broker (works with no API keys ); pluggable LLM adapter interface.
Typed memory classification, importance/confidence/sensitivity scoring, provenance capture.
Policy decisions: SAVE , PENDING_APPROVAL , BLOCK , DROP_LOW_UTILITY , UPDATE_EXISTING , MERGE_WITH_EXISTING .
Secret / PII detection blocks API keys and credentials before storage.
Append-only audit log for every lifecycle event.
Temporary chat short-circuits both read and write.
Memory dashboard + admin/audit + architecture pages (frontend skeleton).
Invariant test suite + eval harness scaffolding.
Loop Engineering Layer (v0.3.1)
MemoryOps models memory as a set of governed loops rather than a passive store.
Each loop has explicit states, policy gates, audit events, fallback behavior, and
evidence requirements. Loop definitions live in services/api/app/loops/ , loop
runs/events are exposed through /api/loops , and the frontend includes a Loops page.
See docs/loop-engineering.md ,
docs/loop-contracts.md , and
docs/release-loop.md .
Token Compression Layer (v0.2.1)
MemoryOps supports an optional Headroom -powered
context compression layer. Compression runs after policy checks, governance
filtering, and context composition, and only on the composed context block —
never the raw user message and never before the policy broker. It reduces tokens
sent to the LLM while preserving MemoryOps invariants (provenance, deletion
guarantee, tenant isolation, temporary-chat behavior, explainability metadata).
It is off by default and not a dependency — the app runs without
headroom-ai installed, and any compression failure degrades safely to the
uncompressed context.
pip install " headroom-ai[all] " # optional
export MEMORYOPS_CONTEXT_COMPRESSION=headroom # default: none
Each chat response carries a compression block with estimated tokens saved and
the compression ratio. See docs/token-compression.md ,
docs/integrations/headroom.md , and
ADR-007 . Headroom is Apache-2.0;
MemoryOps integrates it via an adapter and does not vendor its source.
What works as of v0.3 (real data layer)
Swappable embedding provider ( app/embeddings/ ): deterministic offline stub + optional OpenAI.
Hybrid retrieval : pgvector cosine ( search_candidates ) + keyword overlap, blended by the ranker.
Per-memory score_breakdown + response retrieval_mode ( hybrid / fallback / none ).
Enforced Postgres Row-Level Security (migration 004 , FORCE + tenant policy + session GUC).
Expanded evals (semantic / keyword / archived / score-breakdown) + new tests; RLS test is DB-guarded.
What works as of v0.4 (provider LLM adapters)
Provider-neutral LLM layer ( app/llm/ ): deterministic StubProvider default +
optional OpenAI/Anthropic/Gemini adapters, selected by MEMORYOPS_LLM_PROVIDER .
Structured memory intelligence : schema-validated extraction + minimal conflict
detection, with pro

[truncated]
