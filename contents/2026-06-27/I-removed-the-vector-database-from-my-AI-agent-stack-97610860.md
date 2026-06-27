---
source: "https://github.com/usemoss/moss"
hn_url: "https://news.ycombinator.com/item?id=48693631"
title: "I removed the vector database from my AI agent stack"
article_title: "GitHub - usemoss/moss: The retrieval layer for production AI systems. Lightning-fast (<10ms) search without vector databases. Built for browser, edge, on-device, and cloud. · GitHub"
author: "philosopherr"
captured_at: "2026-06-27T00:30:48Z"
capture_tool: "hn-digest"
hn_id: 48693631
score: 1
comments: 0
posted_at: "2026-06-27T00:05:54Z"
tags:
  - hacker-news
  - translated
---

# I removed the vector database from my AI agent stack

- HN: [48693631](https://news.ycombinator.com/item?id=48693631)
- Source: [github.com](https://github.com/usemoss/moss)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T00:05:54Z

## Translation

タイトル: AI エージェント スタックからベクター データベースを削除しました
記事のタイトル: GitHub - usemoss/moss: 実稼働 AI システムの検索レイヤー。ベクトルデータベースを使用しない超高速 (<10ms) 検索。ブラウザ、エッジ、オンデバイス、クラウド向けに構築されています。 · GitHub
説明: 実稼働 AI システムの取得レイヤー。ベクトルデータベースを使用しない超高速 (<10ms) 検索。ブラウザ、エッジ、オンデバイス、クラウド向けに構築されています。 - usemoss/モス

記事本文:
GitHub - usemoss/moss: 本番 AI システムの取得レイヤー。ベクトルデータベースを使用しない超高速 (<10ms) 検索。ブラウザ、エッジ、オンデバイス、クラウド向けに構築されています。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。これをリロードしてください

のページ。
ユースモス
/
苔
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
211 コミット 211 コミット .github .github アプリ アプリ アセット アセット ベンチマーク ベンチマークの例 例 moss-live-labs moss-live-labs moss-workshop/ starter moss-workshop/ starter パッケージ パッケージ スクリプト スクリプト sdks sdks .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE Package.swift Package.swift README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ウェブサイト · ドキュメント · Discord · ブログ
Moss は、会話型 AI エージェント向けに構築された 10 ミリ秒未満のセマンティック検索ランタイムです。ハイブリッド取得 (セマンティック + キーワード検索)、組み込み埋め込み、メタデータ フィルタリング、ブラウザ内で実行される WebAssembly ビルドのすべてが、アプリケーションに埋め込まれた単一の SDK から行われます。
ホット パス上にネットワーク ホップはありません。調整するクラスターがありません。 SDK を Moss Cloud に向け、インデックスをロードし、10 ミリ秒未満でクエリを実行します。 Python、TypeScript、Elixir、および C.
始める前に: moss.dev で project_id と project_key にサインアップしてください - 無料利用枠が利用可能です。
以下のスニペットには、Python 3.10 以降または Node.js 20 以降が必要です。
pip インストールモス
moss インポート MossClient 、QueryOptions から
client = MossClient ( "your_project_id" , "your_project_key" )
# インデックスを作成してドキュメントを追加する
クライアントを待ちます。 create_index ( "サポートドキュメント" , [
{ "id" : "1" , "text" : "返金は 3 ～ 5 営業日以内に処理されます。" }、
{ "id" : "2" , "text" : "注文はダッシュボードで追跡できます。" }、
{ "id" : "3" , "テックス

t" : "24時間年中無休のライブチャットサポートを提供しています。" },
])
# ロードとクエリ — 結果は 10 ミリ秒未満
クライアントを待ちます。 load_index ( "サポートドキュメント" )
results = クライアントを待ちます。クエリ ( "support-docs" 、 "払い戻しにはどのくらい時間がかかりますか?" 、 QueryOptions ( top_k = 3 ))
結果のドキュメントの場合。ドキュメント:
print ( f"[ { doc . core :.3f } ] { doc . text } " ) # {results.time_taken_ms}ms で返されます
TypeScript
npm インストール @moss-dev/moss
"@moss-dev/moss" から { MossClient } をインポートします。
const client = new MossClient ( "your_project_id" , "your_project_key" ) ;
// インデックスを作成し、ドキュメントを追加します
クライアントを待ちます。 createIndex ( "サポートドキュメント" , [
{ id : "1" 、 text : "返金は 3 ～ 5 営業日以内に処理されます。" } 、
{ id : "2" 、 text : "注文はダッシュボードで追跡できます。" } 、
{ id : "3" 、 text : "24 時間 365 日のライブチャット サポートを提供します。" } 、
]);
// ロードとクエリ - 結果は 10 ミリ秒未満
クライアントを待ちます。 loadIndex ( "サポートドキュメント" ) ;
const results = クライアントを待ちます。 query ("support-docs" , "払い戻しにはどのくらい時間がかかりますか?" , { topK : 3 } ) ;
結果。ドキュメント 。 forEach ( ( doc ) => {
コンソール。 log ( `[ ${ doc . スコア . toFixed ( 3 ) } ] ${ doc . text } ` ) ; // ${results.timeTakenInMs}ms で返されます
} ) ;
なぜモスなのか？
ほとんどの検索スタックは、リモートのベクトル データベースを呼び出します。往復だけでも 200 ～ 500 ミリ秒かかりますが、これはリアルタイムの会話を中断するのに十分です。
Moss はプロセス内で検索と埋め込みを実行します。ホット パスにはネットワーク ホップがないため、クエリのレイテンシは 1 桁に収まります。これは、取得がレイテンシ バジェットから消えるほどの速さです。音声ボット、副操縦士、または人間と会話するエージェントを構築している場合、それが生きていると感じられるツールと遅れていると感じられるツールの違いです。
100,000 ドキュメント、750 測定クエリ、top_k=5 でのエンドツーエンド クエリ レイテンシー (埋め込み + 検索)。 MacBook Proでテスト済み

(M4 プロ、24GB)。
Moss には測定に埋め込みが含まれています。競合他社は外部埋め込みサービス ( modal ) を使用しています。 Pinecone と Qdrant はクラウド検索を使用します。
Moss はデータベースではありません。それは検索ランタイムです。クラスターを管理したり、HNSW パラメーターを調整したり、シャーディングについて心配したりする必要はありません。ドキュメントのインデックスを作成し、ランタイムにロードして、クエリを実行します。それでおしまい。
10 ミリ秒未満のセマンティック検索 - ベンチマークでは 1 桁のミリ秒 p99
ハイブリッド検索 - 単一クエリ内のセマンティック + キーワード
組み込みの埋め込みモデル - OpenAI キーは必要ありません (または独自のキーを持参することもできます)
メタデータのフィルタリング - $eq 、 $and 、 $in 、 $near 演算子
ブラウザでも実行 - サーバーを使用しないクライアント側のセマンティック検索用に別の WebAssembly SDK ( @moss-dev/moss-web )
データベース コネクタ - SQLite、MongoDB、MySQL、および Supabase から直接取り込みます (packages/moss-data-connector/)
CLI - ターミナルからのインデックスとクエリの管理 (packages/moss-cli/)
SDK - Python (3.10 以降)、TypeScript / Node.js (20 以降)、Elixir、および C ( libmoss )
フレームワークの統合 - LangChain、DSPy、LlamaIndex、Pipecat、LiveKit、Vapi、ElementalLabs、Strands Agents
このリポジトリには、プロジェクトに直接コピーできる実際のサンプルが含まれています。
例/
§── python/ # Python SDK サンプル
│ §──load_and_query_sample.py
│ §── 総合サンプル.py
│ §──custom_embedding_sample.py
│ ━──metadata_filtering.py
§── python-classification/ # 分類例
§── javascript/ # TypeScript SDK サンプル
│ §──load_and_query_sample.ts
│ §── 総合サンプル.ts
│ └──custom_embedding_sample.ts
§── javascript-web/ # ブラウザ / WASM SDK サンプル
§── c/# C SDK サンプル (libmoss)
§── go/ # Go SDK サンプル
§── voice-agents/ # エンドツーエンドの音声エージェント (アンビエント + マルチエージェント)
│ §──

airline-pnr/ # アンビエント検索; PNR ごとの Moss インデックス、コール中にスワップ
│ └── 住宅ローン貸与/ # セッション状態を共有するマルチエージェント フロー
└── クックブック/ # フレームワーク統合
§── langchain/ # LangChain レトリーバー
§── dspy/ # DSPy モジュール
§── crewai/ # CrewAI 統合
§── haystack/ # ヘイスタック・レトリーバー
§── autogen/ # AutoGen の統合
§── mastra/ #マストラレトリバー
§── pydantic-ai/ # Pydantic AI 統合
━── daytona/ # デイトナサンドボックス例
アプリ/
§── next-js/ # Next.js セマンティック検索 UI
§── Pipecat-moss/ # Moss 取得を備えた Pipecat 音声エージェント
§── vapi-moss/ # Moss 検索を備えた Vapi 音声エージェント
§── elevenlabs-moss/ # Moss 検索を備えた elevenLabs 音声エージェント
§── livekit-moss-vercel/ # Vercel 上の LiveKit 音声エージェント
§── agora-moss/ # Moss 検索を備えた Agora 会話型 AI MCP サーバー
§── moss-llamaindex/ # LlamaIndex RAG バックエンド + フロントエンド
§── moss-bun/ # Bun ランタイム例
└── docker/ # Docker 化された例 (ECS/K8s パターン)
moss-live-labs/ # 実験ゾーン: プロトタイプとコミュニティデモ
§── python/ # 最小限の Python クイックスタート + 高度なクエリの例
§── typescript/ # 最小限の TypeScript クイックスタート + 高度なクエリの例
§── 例/ # より大規模な実験 (画像検索、音声エージェント)
│ §── voice-agent/ # LiveKit + Moss 音声アシスタント
│ §── Advanced-voice-agent/ # PDF ナレッジ ベース上に構築されたペルソナ インパーソネーター
│ └── image-search/ # FastAPI + COCO を介した React 画像検索
└── コミュニティデモ/ # コミュニティ貢献
[切り捨てられた]
CD の例/Python
pip install -r 要件.txt
cp ../../.env.example .env # 資格情報を追加します
Pythonのload_and_query_sample.py
TypeScript 試験を実行する

お願いします
cd サンプル/JavaScript
npmインストール
cp ../../.env.example .env # 資格情報を追加します
npx tsx ロードアンドクエリ_サンプル.ts
Next.js アプリを実行する
cd apps/next-js
npmインストール
cp ../../.env.example .env # 資格情報を追加します
npm run dev # http://localhost:3000 を開きます
Pipecat 音声エージェントを実行する
10 ミリ秒未満の取得は、実際に会話を続けるカスタマー サポート エージェントである Pipecat のリアルタイム音声パイプラインに接続されています。
cd apps/pipecat-moss/pipecat-quickstart
# セットアップと Pipecat Cloud の展開については README を参照してください
完全にローカルな音声エージェントを実行する (Ollama + Moss + Pipecat)
プライバシー最優先の音声 AI スタック: LLM 推論用の Ollama、取得用の Moss、リアルタイム オーディオ用の Pipecat - LLM と取得の両方がマシン上で実行されます。
cd apps/pipecat-moss/ollama-local
ドッカーの構成
完全な API リファレンス: docs.moss.dev 。
フレームワーク
ステータス
例
ラングチェーン
利用可能
例/クックブック/langchain/
DSPy
利用可能
例/クックブック/dspy/
ラマインデックス
利用可能
apps/moss-llamaindex/
CrewAI
利用可能
例/クックブック/crewai/
自動生成
利用可能
例/クックブック/autogen/
干し草の山
利用可能
例/クックブック/干し草の山/
マストラ
利用可能
例/クックブック/マストラ/
ピダンティックAI
利用可能
例/クックブック/pydantic-ai/
パイプキャット
利用可能
apps/pipecat-moss/
ライブキット
利用可能
apps/livekit-moss-vercel/
バピ
利用可能
アプリ/vapi-moss/
イレブンラボ
利用可能
apps/elevenlabs-moss/
アゴラ
利用可能
アプリ/agora-moss/
ストランドエージェント
利用可能
パッケージ/ストランド-エージェント-モス/
Next.js
利用可能
apps/next-js/
ヴァイトプレス
利用可能
パッケージ/vitepress-plugin-moss/
Vercel AI SDK
利用可能
パッケージ/vercel-sdk/
建築
Moss Cloud - 取り込み、ドキュメントの埋め込み、保存、配布を処理します。プロジェクト ID とキーを使用して SDK を指定します。
インデックス - ドキュメントとそのベクトルを 1 つのアーティファクトとしてパッケージ化し、

モスクラウドにあります。
ランタイム - アプリケーションに埋め込まれます。 HTTPS 経由でインデックスを取得し、メモリに保持して、クエリをローカルで処理します。
インデックスがロードされると、クエリはプロセスから離れることはありません。これが 10 ミリ秒未満のレイテンシの原因です。ドキュメントの変更は Moss Cloud を介して行われ、ランタイムは同期を保ちます。
サーバー側 - moss (Python) および @moss-dev/moss (Node.js 20+) はバックエンドにランタイムを埋め込みます。エージェントがサーバー上で実行されている場合にこれを使用します。
ブラウザ - @moss-dev/moss-web は、インデックスをダウンロードしてクエリを完全にクライアント側で実行する WebAssembly ビルドであり、サーバーは必要ありません。これは、静的サイト、ブラウザ拡張機能、オフラインファーストのアプリに使用します。 example/javascript-web/ を参照してください。
完全な Python SDK ソース コードは sdks/python/ で入手できます。
コミュニティが最も大きな影響を与えることができるのは次のとおりです。
新しい SDK バインディング — Swift、Go、Elixir など
フレームワークの統合 — CrewAI、Haystack、AutoGen
リランキングのサポート - クロスエンコーダーのリランカーをプラグイン
ドキュメント解析コネクタ - PDF、DOCX、HTML、Markdown の取り込み
例とチュートリアル — Moss で何かを構築する場合は、ぜひ紹介したいと思います
セットアップ手順については貢献ガイドを、計画内容についてはロードマップを参照してください。
開始するには、「良好な最初の問題」とラベル付けされた問題を確認してください。
Discord — 質問したり、構築しているものを共有したりできます
GitHub の問題 — バグレポートと機能リクエスト
Twitter — お知らせと最新情報
BSD 2 条項ライセンス — このリポジトリの SDK、サンプル、および統合は完全にオープン ソースです。
プロ向けの検索レイヤー

[切り捨てられた]

## Original Extract

The retrieval layer for production AI systems. Lightning-fast (<10ms) search without vector databases. Built for browser, edge, on-device, and cloud. - usemoss/moss

GitHub - usemoss/moss: The retrieval layer for production AI systems. Lightning-fast (<10ms) search without vector databases. Built for browser, edge, on-device, and cloud. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
usemoss
/
moss
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
211 Commits 211 Commits .github .github apps apps assets assets benchmarks benchmarks examples examples moss-live-labs moss-live-labs moss-workshop/ starter moss-workshop/ starter packages packages scripts scripts sdks sdks .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Package.swift Package.swift README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Website · Docs · Discord · Blog
Moss is a sub-10 ms semantic search runtime built for Conversational AI agents. Hybrid retrieval (semantic + Keyword Search), built-in embeddings, metadata filtering, and a WebAssembly build that runs in the browser - all from a single SDK that embeds in your application.
No network hop on the hot path. No clusters to tune. Point the SDK at Moss Cloud, load your index, and query it in under 10 ms . Python, TypeScript, Elixir, and C.
Before you start: sign up at moss.dev for project_id and project_key - free tier available.
The snippets below need Python 3.10+ or Node.js 20+.
pip install moss
from moss import MossClient , QueryOptions
client = MossClient ( "your_project_id" , "your_project_key" )
# Create an index and add documents
await client . create_index ( "support-docs" , [
{ "id" : "1" , "text" : "Refunds are processed within 3-5 business days." },
{ "id" : "2" , "text" : "You can track your order on the dashboard." },
{ "id" : "3" , "text" : "We offer 24/7 live chat support." },
])
# Load and query — results in <10 ms
await client . load_index ( "support-docs" )
results = await client . query ( "support-docs" , "how long do refunds take?" , QueryOptions ( top_k = 3 ))
for doc in results . docs :
print ( f"[ { doc . score :.3f } ] { doc . text } " ) # Returned in {results.time_taken_ms}ms
TypeScript
npm install @moss-dev/moss
import { MossClient } from "@moss-dev/moss" ;
const client = new MossClient ( "your_project_id" , "your_project_key" ) ;
// Create an index and add documents
await client . createIndex ( "support-docs" , [
{ id : "1" , text : "Refunds are processed within 3-5 business days." } ,
{ id : "2" , text : "You can track your order on the dashboard." } ,
{ id : "3" , text : "We offer 24/7 live chat support." } ,
] ) ;
// Load and query — results in <10 ms
await client . loadIndex ( "support-docs" ) ;
const results = await client . query ( "support-docs" , "how long do refunds take?" , { topK : 3 } ) ;
results . docs . forEach ( ( doc ) => {
console . log ( `[ ${ doc . score . toFixed ( 3 ) } ] ${ doc . text } ` ) ; // Returned in ${results.timeTakenInMs}ms
} ) ;
Why Moss?
Most retrieval stacks call out to a remote vector database. The round trip alone runs 200–500 ms - enough to break a real-time conversation.
Moss runs search and embedding inside your process. There's no network hop on the hot path, so query latency lands in the single digits - fast enough that retrieval disappears from the latency budget. If you're building a voice bot, a copilot, or any agent that talks to humans, that's the difference between a tool that feels alive and one that feels laggy.
End-to-end query latency (embedding + search) on 100,000 documents, 750 measured queries, top_k=5. Tested with Macbook pro (M4 Pro, 24GB).
Moss includes embedding in the measurement — competitors use an external embedding service ( modal ). Pinecone and Qdrant use cloud search.
Moss isn't a database! It's a search runtime . You don't manage clusters, tune HNSW parameters, or worry about sharding. You index documents, load them into the runtime, and query. That's it.
Sub-10 ms semantic search - single-digit-ms p99 in our benchmarks
Hybrid search - semantic + keyword in a single query
Built-in embedding models - no OpenAI key required (or bring your own)
Metadata filtering - $eq , $and , $in , $near operators
Runs in the browser too - separate WebAssembly SDK ( @moss-dev/moss-web ) for client-side semantic search with no server
Database connectors - ingest directly from SQLite, MongoDB, MySQL, and Supabase ( packages/moss-data-connector/ )
CLI - manage indexes and query from the terminal ( packages/moss-cli/ )
SDKs - Python (3.10+), TypeScript / Node.js (20+), Elixir, and C ( libmoss )
Framework integrations - LangChain, DSPy, LlamaIndex, Pipecat, LiveKit, Vapi, ElevenLabs, Strands Agents
This repo contains working examples you can copy straight into your project:
examples/
├── python/ # Python SDK samples
│ ├── load_and_query_sample.py
│ ├── comprehensive_sample.py
│ ├── custom_embedding_sample.py
│ └── metadata_filtering.py
├── python-classification/ # Classification example
├── javascript/ # TypeScript SDK samples
│ ├── load_and_query_sample.ts
│ ├── comprehensive_sample.ts
│ └── custom_embedding_sample.ts
├── javascript-web/ # Browser / WASM SDK samples
├── c/ # C SDK samples (libmoss)
├── go/ # Go SDK samples
├── voice-agents/ # End-to-end voice agents (ambient + multi-agent)
│ ├── airline-pnr/ # Ambient retrieval; per-PNR Moss indexes, swap mid-call
│ └── mortgage-lending/ # Multi-agent flow with shared session state
└── cookbook/ # Framework integrations
├── langchain/ # LangChain retriever
├── dspy/ # DSPy module
├── crewai/ # CrewAI integration
├── haystack/ # Haystack retriever
├── autogen/ # AutoGen integration
├── mastra/ # Mastra retriever
├── pydantic-ai/ # Pydantic AI integration
└── daytona/ # Daytona sandbox example
apps/
├── next-js/ # Next.js semantic search UI
├── pipecat-moss/ # Pipecat voice agent with Moss retrieval
├── vapi-moss/ # Vapi voice agent with Moss retrieval
├── elevenlabs-moss/ # ElevenLabs voice agent with Moss retrieval
├── livekit-moss-vercel/ # LiveKit voice agent on Vercel
├── agora-moss/ # Agora Conversational AI MCP server with Moss retrieval
├── moss-llamaindex/ # LlamaIndex RAG backend + frontend
├── moss-bun/ # Bun runtime example
└── docker/ # Dockerized examples (ECS/K8s pattern)
moss-live-labs/ # Experimental zone: prototypes and community demos
├── python/ # Minimal Python quickstart + advanced query example
├── typescript/ # Minimal TypeScript quickstart + advanced query example
├── examples/ # Larger experiments (image search, voice agents)
│ ├── voice-agent/ # LiveKit + Moss voice assistant
│ ├── advanced-voice-agent/ # Persona impersonator built on a PDF knowledge base
│ └── image-search/ # FastAPI + React image search over COCO
└── community-demos/ # Community-contributed
[truncated]
cd examples/python
pip install -r requirements.txt
cp ../../.env.example .env # Add your credentials
python load_and_query_sample.py
Run the TypeScript examples
cd examples/javascript
npm install
cp ../../.env.example .env # Add your credentials
npx tsx load_and_query_sample.ts
Run the Next.js app
cd apps/next-js
npm install
cp ../../.env.example .env # Add your credentials
npm run dev # Open http://localhost:3000
Run the Pipecat voice agent
Sub-10 ms retrieval plugged into Pipecat's real-time voice pipeline — a customer support agent that actually keeps up with conversation.
cd apps/pipecat-moss/pipecat-quickstart
# See README for setup and Pipecat Cloud deployment
Run the fully-local voice agent (Ollama + Moss + Pipecat)
A privacy-first voice AI stack: Ollama for LLM inference, Moss for retrieval, Pipecat for real-time audio - the LLM and retrieval both run on your machine.
cd apps/pipecat-moss/ollama-local
docker compose up
Full API reference: docs.moss.dev .
Framework
Status
Example
LangChain
Available
examples/cookbook/langchain/
DSPy
Available
examples/cookbook/dspy/
LlamaIndex
Available
apps/moss-llamaindex/
CrewAI
Available
examples/cookbook/crewai/
AutoGen
Available
examples/cookbook/autogen/
Haystack
Available
examples/cookbook/haystack/
Mastra
Available
examples/cookbook/mastra/
Pydantic AI
Available
examples/cookbook/pydantic-ai/
Pipecat
Available
apps/pipecat-moss/
LiveKit
Available
apps/livekit-moss-vercel/
Vapi
Available
apps/vapi-moss/
ElevenLabs
Available
apps/elevenlabs-moss/
Agora
Available
apps/agora-moss/
Strands Agents
Available
packages/strands-agents-moss/
Next.js
Available
apps/next-js/
VitePress
Available
packages/vitepress-plugin-moss/
Vercel AI SDK
Available
packages/vercel-sdk/
Architecture
Moss Cloud - handles ingestion, document embedding, storage, and distribution. Point the SDK at it with a project ID and key.
Index - your documents and their vectors, packaged as a single artifact that lives on Moss Cloud.
Runtime - embedded in your application. It pulls indexes over HTTPS, holds them in memory, and serves queries locally.
Once an index is loaded, queries don't leave your process - that's where the sub-10 ms latency comes from. Document changes flow through Moss Cloud and the runtime stays in sync.
Server-side - moss (Python) and @moss-dev/moss (Node.js 20+) embed the runtime in your backend. Use this when your agent runs on a server.
Browser - @moss-dev/moss-web is a WebAssembly build that downloads the index and runs queries entirely client-side, no server required. Use this for static sites, browser extensions, and offline-first apps. See examples/javascript-web/ .
Full Python SDK source code is available at sdks/python/ .
Here's where the community can have the most impact:
New SDK bindings — Swift, Go, Elixir,...
Framework integrations — CrewAI, Haystack, AutoGen
Reranking support — plug in cross-encoder rerankers
Doc-parsing connectors — PDF, DOCX, HTML, Markdown ingestion
Examples and tutorials — if you build something with Moss, we'd love to feature it
See our Contributing Guide for setup instructions and our Roadmap for what's planned.
Check out issues labeled good first issue to get started.
Discord — ask questions, share what you're building
GitHub Issues — bug reports and feature requests
Twitter — announcements and updates
BSD 2-Clause License — the SDKs, examples, and integrations in this repo are fully open source.
The retrieval layer for pro

[truncated]
