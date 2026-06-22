---
source: "https://robert-mcdermott.github.io/phlox/"
hn_url: "https://news.ycombinator.com/item?id=48632312"
title: "Phlox: A full-featured AI platform you own"
article_title: "Phlox — A full-featured, self-hostable AI platform"
author: "mcdermott"
captured_at: "2026-06-22T16:33:13Z"
capture_tool: "hn-digest"
hn_id: 48632312
score: 1
comments: 0
posted_at: "2026-06-22T16:24:14Z"
tags:
  - hacker-news
  - translated
---

# Phlox: A full-featured AI platform you own

- HN: [48632312](https://news.ycombinator.com/item?id=48632312)
- Source: [robert-mcdermott.github.io](https://robert-mcdermott.github.io/phlox/)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T16:24:14Z

## Translation

タイトル: Phlox: あなたが所有するフル機能の AI プラットフォーム
記事のタイトル: Phlox — フル機能を備えた自己ホスト可能な AI プラットフォーム
説明: Phlox は、エージェント ハーネス、ドキュメント RAG、コード実行、OpenAI 互換ゲートウェイ、ユーザーごとのコスト計算、および MCP 統合を備えたフル機能の自己ホスト型 AI プラットフォームであり、あらゆるモデル プロバイダー (AWS Bedrock または完全なローカル モデルを含む OpenAI 互換エンドポイント) 上で実行されます。

記事本文:
フロックス
特長
エージェント
知識
セキュリティ
プラットフォーム
建築
クイックスタート
GitHub
自己ホスト可能 · オープンソース · Apache 2.0
フル機能の AI プラットフォーム
あなたが実際に所有しています。
Phlox は自己ホスト型 AI プラットフォームです。ハーネス、ドキュメント RAG、
コード実行、OpenAI 互換ゲートウェイ、ユーザーごとのコスト計算 - 実行中
あらゆるモデルプロバイダー上で: AWS Bedrock または任意の OpenAI 互換
完全にローカルなモデルを含むエンドポイント。
GitHub で見る
クイックスタート →
轢かれる
AWS の基盤
ご自身のモデルをご持参ください。または、すべてをローカルで実行します。
名前付きプロバイダー プロファイルは、AWS Bedrock とその他のあらゆるサービスをカバーします。
OpenAI 互換エンドポイント — OpenAI、LiteLLM、またはローカル ランタイム。ポイントフロックス
Ollama、LM Studio、または vLLM とスタック全体 (チャットと RAG 埋め込み)
クラウド API キーなしでオフラインで実行されます。組み込みの接続を使用して、プロファイルをライブで切り替えます
テスター。
プロバイダー プロファイルを好きなだけ定義し、即座に切り替えることができます
ローカル埋め込み (例: nomic-embed-text ) は RAG を完全にオフラインに保ちます
プロファイル、価格設定、制限をライブで編集 - サーバーの再起動は必要ありません
デフォルトプロファイル : ローカルオラマ
プロフィール:
ローカルオラマ :
タイプ：オープンアイ
レーベル：「オラマ（ローカル）」
エンドポイント: http://localhost:11434/v1
api_key : ollam # Ollama によって無視される
モデル：qwen3.6:35b
サポートツール: true
すべてを 1 つのアプリで
単なるチャットボックスではない完全なアシスタント
Phlox は、自分でつなぎ合わせていた部分をバンドルしており、それぞれが自己ホストされており、制御下にあります。
会話履歴の名前変更、削除、検索、エクスポート。メッセージの編集/再生成、強調表示されたコピー可能なコードによるマークダウン、さらに人魚図と LaTeX 数学。
このモデルは、ファイルシステム、シェル、Python/ノードの実行、ドキュメント検索に加えて、計画、サブエージェント、メモリ、チェックポイントなどのツールをループで使用し、すべてサンドボックス化された環境で実行します。

アークスペース。
機密性の高いツールを一時停止し、承認または拒否してから再開します。実行状態は保持されるため、承認は切断されても存続します。
キャプチャされた出力とインライン成果物を使用してコードを実行します。 [ワークスペース ファイル] パネルでは、エージェントが作成したすべてのものを参照してダウンロードできます。
PDF、DOCX、TXT、MD、またはコードをアップロードします。再ランキングと引用を使用した Qdrant の高密度 + 疎のハイブリッド検索。グローバルまたは会話ごとにスコープが設定されます。オフラインでも動作します。
プロンプトごとのコンポーザーの切り替えにより web_search (ゼロ構成 DDGS または SearXNG) が公開されるため、エージェントはページをフェッチする前に現在のソースを検出できます。
永続的なファクトは保存され、チャット間で意味的に呼び戻されるため、アシスタントは会話ごとにあなたを覚えています。
ビジョン モデルのメッセージに画像を添付し、画像コンテンツ パーツとしてプロバイダーに永続化および再生されます。
UI から Model Context Protocol サーバーに接続します。彼らのツールはモデルのツールセットに自動的に結合され、コードは必要ありません。
ユーザーごとの API キーを作成し、/v1/chat/completions 経由で任意の OpenAI SDK から Phlox を呼び出します。同じユーザーごとのコスト計算を使用します。
UI でのメッセージごとのトークンとコストに加え、月×ユーザー×部門×モデルごとの管理者チャージバック ビューと、財務用の CSV エクスポートが可能です。
デフォルトではフロックス ダーク、ライト、フレッド ハッチ、ハッチ ナイト、サンドストーンなど - CSS 変数トークン システムを介して瞬時に切り替えられます。
「ツールを呼び出すチャット」ではない、本物のエージェント
各ターン、モデルは、検査、スナップショット、ロールバックが可能な会話ごとのサンドボックス化されたワークスペース内で、ツールの呼び出し、計画、回復というループで動作します。
ファイルシステム ( read_file 、 write_file 、 edit_file 、 glob 、 grep )、 run_shell 、execute_python /execute_node、および search_documents — タスクが完了するまでモデルが駆動する 1 つの統合ツール サーフェス。
update_todos は計画を目に見える状態に保ちます。 spawn_subagent は nes を実行します

ted は、同じワークスペース内でスコープ指定されたツールセットを持つ一時的なエージェントであり、レポートを返します。
save_memory はチャット間で永続的なファクトを保持します。すべてのワークスペースは、ツールを変更した後にワンクリックで復元できる自動スナップショットを作成する git リポジトリです。
すべてのツールには自動 / 質問 / 拒否ポリシーがあります。ループは ask で一時停止し、その状態を保持し、決定後にステートレスに再開します。
ドキュメントは正しい方法で検索されました
PDF、Office ドキュメント、マークダウン、またはソース コードをアップロードします。 Phlox は解析、チャンク、および
それらを Qdrant に埋め込み、真のハイブリッド検索で取得します。
RRF と融合された、チャンクごとの密な意味ベクトルと疎な語彙ベクトル。
再ランク付けされ、モデルに引用するように指示された番号付き引用を返します。
グローバルな知識ベースまたは会話ごとのドキュメントの範囲設定
依存関係のないスパース ベクターとリランカーは完全にオフラインで動作します
SQLite は信頼できる情報源であり続けます — インデックスはいつでも再構築できます
会話間の記憶により、あらゆる場面で永続的な事実が思い出されます
マルチユーザー、孤立、責任
認証はデフォルトでオンになっており、データの範囲はユーザーごとに厳密に設定され、すべての機密ツールは制御する権限ゲートの背後で実行されます。
ローカル アカウント (bcrypt + JWT) または Microsoft Entra ID SSO。ユーザー/管理者の役割、ユーザーごとの厳密な分離 — 管理者はアカウントを管理しますが、他のユーザーのコンテンツを読み取ることはできません。
CPU、メモリ、PID 制限とネットワーク分離を備えた一時的な Podman/Docker コンテナでエージェント コードを実行するか、信頼できるシングル ユーザーが使用する場合は高速ローカル サブプロセスを実行します。
各ツールは auto 、 ask 、またはdeny です。変更ツールと実行ツールはデフォルトで ask します。エージェント モードの切り替えは 1 回のターンを自動承認します。
管理パネルからプロバイダー プロファイル、モデル価格、復元力、サンドボックス制限を編集します。再起動せずに適用されます。 API キーは書き込み専用でマスクされています。
リクエストごとの構造化ログ、オプションの OpenTelemetr

y トレースシーム、および耐久性のある台帳でのターンごとのトークン/コストのキャプチャ。
耐久性のある使用状況台帳は、追跡するアカウントの存続期間を超えて存続します。つまり、離脱したユーザーの費用は、アカウントが削除された後も引き続き請求されます。月×ユーザー×部門×機種別の利用状況、CSV出力可能。
チーム全体のための LLM ゲートウェイとコスト台帳
チャットを超えて、Phlox は、ユーザーごとの API キー、ライブ モデルの価格設定、部門レベルのチャージバックを備えた OpenAI 互換ゲートウェイであり、チャット アプリを共有インフラストラクチャに変えるガバナンス層です。
FastAPI バックエンドは、LLM オーケストレーション、エージェント ハーネス、MCP、RAG、コード実行、認証、SQLite 永続性を処理します。 React + Vite フロントエンドは、リッチなストリーミング UI をレンダリングします。
Zustand ストア — ライブストリーミングアセンブリ
/api/chat の SSE ストリーム パーサー
ツールカード、推論、インラインアーティファクト
再開可能なエージェント ループ + ツール レジストリ
許可ゲート — セキュリティの継ぎ目
プロバイダー: OpenAI 互換および Bedrock
RAG・サンドボックス・ワークスペース・MCP
SQLite の信頼できるソース + Qdrant インデックス
開発環境では、Vite は /api を FastAPI にプロキシします。運用環境では、FastAPI はフロントエンド/dist から構築された SPA を提供します。つまり、1 つのコマンドで全体を実行できます。
8つのテーマ、瞬時に切り替え
セマンティック CSS 変数トークン レイヤーは、テーマを再構築せずに変更できることを意味し、独自のテーマを追加するのは 2 つの小さな編集だけです。
前提条件: Python 3.11 以降 (uv 付き)、Node 18 以降、およびモデル プロバイダー - ローカルの Ollama が最も簡単です。
# バックエンドから/
UV同期
cp config.yml.example config.yml # プロバイダーを設定します
uv run uvicorn app.main:app --reload --port 8000
2 · フロントエンド
コピー
# フロントエンド/から、別のターミナル
npmインストール
npm 実行開発
# http://localhost:5173 を開きます
Windows では、両方を ./scripts/dev.ps1 で実行します。 macOS/Linux では ./scripts/dev.sh 。
認証はシードされた管理者/管理者でデフォルトでオンになります -
共有する前にそれを変更し、実際の jwt_secret を設定します

アクセス。
システム マップ、リクエスト ライフサイクル、モジュール ガイド — ここから始めてください。
Tier 1 ～ 5 全体で何が行われ、次に何が行われるか。
ローカル アカウント、ロール、分離、および Entra ID SSO セットアップ。
ローカルでのコード実行と Podman/Docker コンテナー コードの実行。
トークンの使用量/コスト、構造化ログ、OpenTelemetry トレース。
OpenAI 互換キーと /v1/* エンドポイント。
モデル コンテキスト プロトコル サーバーの接続。
テーマトークンシステムと新しいテーマの追加。
今すぐ独自の AI アシスタントをセルフホストしましょう
Apache 2.0 でのオープンソース。クローンを作成し、モデルにポイントして実行します。
フル機能を備えた自己ホスト可能な AI プラットフォーム。

## Original Extract

Phlox is a full-featured, self-hostable AI platform with an agentic harness, document RAG, code execution, an OpenAI-compatible gateway, per-user cost accounting, and MCP integration — running over any model provider: AWS Bedrock or any OpenAI-compatible endpoint, including fully local models.

Phlox
Features
Agent
Knowledge
Security
Platform
Architecture
Quick start
GitHub
Self-hostable · Open source · Apache 2.0
A full-featured AI platform
you actually own.
Phlox is a self-hostable AI platform — an agentic tool-using harness, document RAG,
code execution, an OpenAI-compatible gateway, and per-user cost accounting — running
over any model provider: AWS Bedrock or any OpenAI-compatible
endpoint, including fully local models.
View on GitHub
Quick start →
Runs over
AWS Bedrock
Bring your own model. Or run it all locally.
Named provider profiles cover AWS Bedrock and any
OpenAI-compatible endpoint — OpenAI, LiteLLM, or a local runtime. Point Phlox at
Ollama, LM Studio, or vLLM and the whole stack — chat and RAG embeddings —
runs offline with no cloud API key. Switch profiles live, with a built-in connection
tester.
Define as many provider profiles as you like, switch between them instantly
Local embeddings (e.g. nomic-embed-text ) keep RAG fully offline
Edit profiles, pricing, and limits live — no server restart required
default_profile : local-ollama
profiles :
local-ollama :
type : openai
label : "Ollama (local)"
endpoint : http://localhost:11434/v1
api_key : ollama # ignored by Ollama
model : qwen3.6:35b
supports_tools : true
Everything in one app
A complete assistant, not just a chat box
Phlox bundles the pieces you'd otherwise stitch together yourself — each one self-hosted and under your control.
Conversation history with rename, delete, search & export. Edit/regenerate messages, markdown with highlighted, copyable code, plus Mermaid diagrams and LaTeX math.
The model uses tools in a loop — filesystem, shell, Python/Node execution, document search, plus planning, sub-agents, memory, and checkpoints — all in a sandboxed workspace.
Pause on sensitive tools, approve or deny, then resume. The run state is persisted, so approvals survive disconnects.
Run code with captured output and inline artifacts. A Workspace Files panel lets you browse and download everything the agent created.
Upload PDF, DOCX, TXT, MD, or code. Hybrid dense + sparse search over Qdrant with reranking and citations, scoped globally or per conversation. Works offline.
A per-prompt composer toggle exposes web_search (zero-config ddgs or SearXNG) so the agent can discover current sources before fetching pages.
Durable facts are saved and semantically recalled across chats, so the assistant remembers you from one conversation to the next.
Attach images to messages for vision models, persisted and replayed into the provider as image content parts.
Connect Model Context Protocol servers from the UI; their tools join the model's toolset automatically, no code required.
Mint per-user API keys and call Phlox from any OpenAI SDK via /v1/chat/completions — with the same per-user cost accounting.
Per-message token and cost in the UI, plus an admin chargeback view by month × user × department × model, with CSV export for finance.
Phlox Dark by default, with Light, Fred Hutch, Hutch Night, Sandstone and more — instant switching via a CSS-variable token system.
A real agent, not "chat that calls tools"
Each turn, the model works in a loop — calling tools, planning, and recovering — inside a per-conversation sandboxed workspace you can inspect, snapshot, and roll back.
Filesystem ( read_file , write_file , edit_file , glob , grep ), run_shell , execute_python / execute_node , and search_documents — one unified tool surface the model drives until the task is done.
update_todos keeps a visible plan; spawn_subagent runs a nested, ephemeral agent with a scoped toolset in the same workspace and returns a report.
save_memory persists durable facts across chats. Every workspace is a git repo that auto-snapshots after mutating tools, with one-click restore.
Every tool has an auto / ask / deny policy. The loop pauses on ask , persists its state, and resumes statelessly after you decide.
Your documents, searched the right way
Upload PDFs, Office docs, markdown, or source code. Phlox parses, chunks, and
embeds them into Qdrant , then retrieves with true hybrid search —
a dense semantic vector and a sparse lexical vector per chunk, fused with RRF and
reranked, returning numbered citations the model is instructed to cite.
Global knowledge base or per-conversation document scoping
Dependency-free sparse vectors and reranker work fully offline
SQLite stays the source of truth — the index can always be rebuilt
Cross-conversation memory recalls durable facts into every turn
Multi-user, isolated, and accountable
Auth is on by default, data is scoped strictly per user, and every sensitive tool runs behind a permission gate you control.
Local accounts (bcrypt + JWT) or Microsoft Entra ID SSO. user / admin roles, strict per-user isolation — admins manage accounts but can't read others' content.
Run agent code in an ephemeral Podman/Docker container with CPU, memory, and PID limits plus network isolation — or a fast local subprocess for trusted single-user use.
Each tool is auto , ask , or deny . Mutating and execution tools default to ask ; an Agent-mode toggle auto-approves for a single turn.
Edit provider profiles, model pricing, resilience, and sandbox limits from an admin panel — applied without a restart. API keys are write-only and masked.
Per-request structured logs, an optional OpenTelemetry tracing seam, and per-turn token/cost capture in a durable ledger.
A durable usage ledger outlives the accounts it tracks — a departed user's costs stay billable after their account is deleted. Usage by month × user × department × model, CSV-exportable.
An LLM gateway and cost ledger for the whole team
Beyond chat, Phlox is an OpenAI-compatible gateway with per-user API keys, live model pricing, and department-level chargeback — the governance layer that turns a chat app into shared infrastructure.
A FastAPI backend handles LLM orchestration, the agent harness, MCP, RAG, code execution, auth, and SQLite persistence. A React + Vite frontend renders the rich, streaming UI.
Zustand store — live streaming assembly
SSE stream parser for /api/chat
Tool cards, reasoning, inline artifacts
Resumable agent loop + tool registry
Permission gate — the security seam
Providers: OpenAI-compatible & Bedrock
RAG · sandbox · workspace · MCP
SQLite source of truth + Qdrant index
In dev, Vite proxies /api to FastAPI. In production, FastAPI serves the built SPA from frontend/dist — one command to run the whole thing.
Eight themes, instant switching
A semantic CSS-variable token layer means themes change with no rebuild — and adding your own is two small edits.
Prerequisites: Python 3.11+ with uv , Node 18+, and a model provider — a local Ollama is the easiest.
# from backend/
uv sync
cp config.yml.example config.yml # set your provider
uv run uvicorn app.main:app --reload --port 8000
2 · Frontend
Copy
# from frontend/, separate terminal
npm install
npm run dev
# open http://localhost:5173
On Windows run both with ./scripts/dev.ps1 ; on macOS/Linux ./scripts/dev.sh .
Auth is on by default with a seeded admin / admin —
change it and set a real jwt_secret before sharing access.
System map, request lifecycle, module guide — start here.
What's done and what's next across Tiers 1–5.
Local accounts, roles, isolation, and Entra ID SSO setup.
Local vs Podman/Docker container code execution.
Token usage/cost, structured logs, OpenTelemetry tracing.
OpenAI-compatible keys and /v1/* endpoints.
Connecting Model Context Protocol servers.
The theme token system and adding new themes.
Self-host your own AI assistant today
Open source under Apache 2.0. Clone it, point it at a model, and run.
A full-featured, self-hostable AI platform.
