---
source: "https://github.com/faisalhussain-devs/MindCache/tree/collapsed_tree"
hn_url: "https://news.ycombinator.com/item?id=48997473"
title: "Show HN: MindCache – An open-source agentic memory system for LLMs"
article_title: "GitHub - faisalhussain-devs/MindCache · GitHub"
author: "Faisal9876"
captured_at: "2026-07-21T20:14:17Z"
capture_tool: "hn-digest"
hn_id: 48997473
score: 1
comments: 0
posted_at: "2026-07-21T20:03:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: MindCache – An open-source agentic memory system for LLMs

- HN: [48997473](https://news.ycombinator.com/item?id=48997473)
- Source: [github.com](https://github.com/faisalhussain-devs/MindCache/tree/collapsed_tree)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T20:03:59Z

## Translation

タイトル: Show HN: MindCache – LLM 用のオープンソース エージェント メモリ システム
記事のタイトル: GitHub - faisalhussain-devs/MindCache · GitHub
説明: GitHub でアカウントを作成して、faisalhussain-devs/MindCache の開発に貢献します。

記事本文:
GitHub - faisalhussain-devs/MindCache · GitHub
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
アラートを閉じる
{{ メッセージ }}
ファイサルフセイン開発者
/
マインドキャッシュ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
折りたたまれた_ツリー

ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
73 コミット 73 コミット BEAM BEAM eval eval 例 例 mincache mincache .gitignore .gitignore BLOG.pdf BLOG.pdf README.md README.md pyproject.toml pyproject.toml required.txt required.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM エージェント用の構造化長期メモリ SDK
MindCache は、運用グレードの LLM エージェント向けに設計された構造化された長期メモリ エンジンです。フラット ベクトル検索データベースや基本的なサマリー ストアとは異なり、MindCache は明示的な意思決定追跡を備えた自己再構築型の階層メモリ オントロジーを維持し、BEAM QA ベンチマークでフラット検索システムよりも優れたパフォーマンスを発揮します。
📖 設計ストーリーを読む: アーキテクチャ設計の詳細については、ブログ投稿「MindCache の構築: 長期 AI のためのエージェントティック メモリ システムの設計」を参照してください。
pip インストールマインドキャッシュ
2. 使用方法
MindCache から MindCache をインポート
# クライアントを初期化します (デフォルトでは SQLite を使用)
mc = マインドキャッシュ (
db_path = "my_memory.db" ,
プロバイダー = "ジェミニ" 、
モデル名 = "ジェミニ-2.5-フラッシュ"
)
# 1. 会話ターンの取り込み (ミリ秒単位でキューに書き込みます)
job_id = mc 。追加([
{ "role" : "user" , "content" : "バックエンド開発には Python と FastAPI、DB には Postgres を好みます。" }、
{ "role" : "assistant" , "content" : "わかりました! Python、FastAPI、Postgres の好みを覚えておきます。" }
]、user_id = "アリス" )
# 2. 保留キューの処理 (抽出を実行し、ツリーとインデックスを更新します)
マック。 process_queue ( user_id = "アリス" )
# 3. 今後のプロンプトに関連するコンテキストを取得する
コンテキスト = mc 。 search ( "アリスの好みのデータベースは何ですか?" , user_id = "alice" )
印刷 (コンテキスト)
🛡️ 主要なアーキテクチャの差別化要因 (なぜ MindCache なのか?)
ほとんどのメモリ システムは、長期記憶を unstr のフラット プールとして扱います。

構築された埋め込みベクトル。 MindCache は、推論エージェント向けに最適化された、構造化された自己組織化メモリ アーキテクチャを導入します。
🔹 フェーズ 1 — インジェスト パイプライン
生の会話はキューに入れられ、埋め込まれ、バッチで処理されます。グラウンデッド ルーティングは、抽出前に既存のツリー パスを制約としてクエリすることによって実行されます。 Decision State Analyzer と再編成プロセスは、後でバッチ タスクとして実行されます。
フローチャート TD
クラス定義ダークフィル:#313244、カラー:#cdd6f4、ストローク:#585b70
classDef llm 塗りつぶし:#cba6f7、カラー:#11111b、ストローク:#11111b
classDef ツリーの塗りつぶし:#89dceb、カラー:#11111b、ストローク:#11111b
classDef 黄色の塗りつぶし:#f9e2af、色:#11111b、ストローク:#11111b
classDef 緑の塗りつぶし:#a6e3a1、色:#11111b、ストローク:#11111b
classDef アナライザー 塗りつぶし:#f38ba8、色:#11111b、ストローク:#11111b
A["💬 Raw Turn"]:::dark --> B{"🤖 覚えておく価値はありますか?"}:::llm
B -- いいえ --> C["無視して戻る"]:::dark
B -- はい --> D["🌲 上位 K 個のトピック パスを取得"]:::tree
D --> E["スマート接地ルーティング"]:::llm
E --> F["🤖 思い出の抽出と分類"]:::llm
F --> G["トピック ツリー パスに保存"]:::tree
G --> H{"バッチは完了しましたか?"}:::dark
H -- はい --> I["🔍 意思決定状態の分析"]:::analyzer
I --> J{"⏱️ REORG_THRESHOLD を超えましたか?"}::: yellow
J -- いいえ --> K["キャッシュを更新して永続化"]:::green
J -- はい --> L["✂️ ツリーの再編成\n(ライデン分割 / 分割 / マージ)"]::: yellow
L --> M{"要約は有効ですか?"}::: yellow
M -- はい --> N["📜 デルタサマリーの生成"]:::llm
M -- いいえ --> K
N --> K
読み込み中
🔹 フェーズ 2 — 動的ツリーのライフサイクル
トピック ツリーは決して静的ではありません。取り込み数が再編成のしきい値を超えると、バックグラウンド プロセスによってグラフが再編成され、増分ボトムアップ集計が作成されます。
フローチャート LR
classDef ツリーの塗りつぶし:#cba6f7、カラー:#11111b、ストローク:#11111b
classDef reorg fill:#f9e2af,color:#1111

1b、ストローク:#11111b
classDef 概要 塗りつぶし:#a6e3a1、色:#11111b、ストローク:#11111b
classDef トリガーの塗りつぶし:#313244、カラー:#cdd6f4、ストローク:#585b70
T["🌲 動的トピック ツリー\n(階層的オントロジー)"]:::tree
R1["✂️ オーバーロードされたリーフ ノードを分割する"]:::reorg
R2["🔗 スパース兄弟ノードをマージ"]:::reorg
R3["⬆️ パスの昇格/降格"]:::reorg
R4["🌐 ライデン グラフの分割"]:::reorg
S["📜 増分デルタ集計\n(ボトムアップデルタロールアップ)"]:::summary
TH["⏱️ しきい値トリガー\n(TriadBlock 数)"]:::トリガー
TH --> T
T --> R1 & R2 & R3 & R4
R1 & R2 & R3 & R4 --> T
T -- 「新葉の思い出」 --> S
S -- 「親に保存された概要」 --> T
読み込み中
🔹 フェーズ 3 — マルチステージハイブリッド検索エンジン
検索クエリは、スコア生成と RRF、クエリ分類、意思決定の拡張、ディレクティブの組み立てといった多段階のプリファレンスに基づいたパイプラインを通じて処理されます。クロスエンコーダーの再ランカーは評価され、削除されました。これは、最大 23 倍の遅延コストでわずかな精度の向上を実現しました。
フローチャート TD
クラス定義ダークフィル:#313244、カラー:#cdd6f4、ストローク:#585b70
classDef llm 塗りつぶし:#cba6f7、カラー:#11111b、ストローク:#11111b
classDef 検索塗りつぶし:#89b4fa、カラー:#11111b、ストローク:#11111b
classDef アンカー塗りつぶし:#f38ba8、カラー:#11111b、ストローク:#11111b
classDef 緑の塗りつぶし:#a6e3a1、色:#11111b、ストローク:#11111b
Q["🔎ユーザークエリ"]:::dark --> A["コンテキストブリッジと埋め込み"]:::dark
A --> B["スコア (ベクター + BM25)"]:::検索
B --> C["相互順位融合 RRF"]:::検索
C --> D["🏷️ クエリを分類\n(広範 vs 事実 / 時間的)"]:::llm
D --> E["分割された候補の取得"]:::検索
E --> F["メモリ プール\n(ユーザー / 知識 / エピソード / 意思決定)"]:::検索
D -->|"広範囲のクエリの場合"| G["📜 サマリー プール"]:::検索
F & G --> H["⚖️ トップ 3 の決定を取得"]:::anchor
H --> I["BM25 アンカー拡張"]:::アンカー
私 --> J[

「重複候補を削除」]:::検索
F & J --> K["📋 ディレクティブの挿入\n(プロファイル、最新性、棄権)"]:::anchor
K --> OUT["📤 最終コンテキスト → LLM プロンプト"]:::green
読み込み中
1. 4 つの構造化メモリ タイプ
MindCache は、一般的なチャンクを保存するのではなく、情報をセマンティック タイプに分離します。
📝 エピソード : インタラクティブな会話イベント、マイルストーン、セッションのコンテキスト。
🧠 知識 : 事実の記述、概念、および事実に関する領域の知識。
👤 ユーザー : アイデンティティの事実、現在の好み、および行動の属性。
⚖️ 決定 : ユーザーによって行われた明示的な選択、指示、および解決。
意思決定は時間の経過とともに進化します。 MindCache は、バックグラウンドで意思決定状態アナライザーを実行して、競合する意思決定または更新された意思決定を特定します。次のように状態を追跡し、フラグを立てます。
Active : 現在のバインディングの選択。
Superseded : 新しい決定によって上書きされます (標準コンテキストからは自動的に抑制されますが、履歴として保存されます)。
条件付き : 未解決の条件に依存します。
Rejected : ユーザーがこのオプションに反対することを明示的に決定しました。
3. 意思決定に基づく BM25 の拡張
決定は意味論的なアンカーとして機能します。検索では、MindCache はクエリに関連する上位の意思決定記憶を特定します。次に、それらの決定を拡張アンカーとして使用して、関連するエピソード、知識、ユーザーの記憶を取得し、高度に調整された背景コンテキストを提供します。
4. スマートな取り込みと接地されたルーティング
トピック ツリーの断片化 (LLM が重複したトピック パスや、FastAPI や FastAPI バックエンドなどのわずかに異なるトピック パスを作成する場合) を防ぐために、MindCache はグラウンデッド ルーティングを実行します。 LLM エクストラクターを呼び出す前に、既存のツリーから上位 K 個の最も近いパスをクエリし、それらを LLM の分類をガイドするための基礎制約として渡します。
5. 自己再構成する階層型トピックツリー
MindCache がメモリを整理します

木にオリエスします。ただし、静的な階層構造とは異なり、ツリーは生きており、バックグラウンドで自身を再編成します。
大きくなりすぎる、または幅が広すぎる葉のノードを分割します。
意味上の重複がある疎な兄弟ブランチをマージします。
ノードを昇格/降格してパスの深さを最適化します。
6. 増分デルタ要約
概要は広範なクエリを解決するためにボトムアップで構築されます。 MindCache はこれを段階的に実行します。新しいメモリがリーフ ノードに追加されると、既存のサマリー コンテキストと結合されたデルタ (最後のロールアップ以降の新しいエントリ) のみが処理され、LLM API トークンの使用量が最小限に抑えられます。
7. BM25 形態学的エイリアシング (Union-Find)
標準の BM25 語彙検索では、単語の語形変化 (例: データベースとデータベース、実行と実行) が見つかりません。 MindCache は、Union-Find データ構造と spaCy 見出し語化を組み合わせることで、この問題を解決します。基本語とその語形変化を 1 つの同値クラスにまとめ、クエリ中にメモリ内でそれらのインデックスを組み合わせて最大限の再現率を実現します。
MindCache は、次のような複数段階の検索シーケンスを使用します。
パーティション化された RRF : メモリ タイプごとにベクトル検索 + BM25 を個別に実行し、相互ランク融合を介してそれらをマージします。
適応型クエリ分類 : 事実クエリと大まかな概要クエリを検出します。キーワードが欠落しており、ベクトルの類似性が失敗する広範な概要クエリの場合、RAPTOR スタイルの階層的な要約が動的に取得されます。
Decision-Anchor Expansion : ファンアウトして、最上位のデシジョンアンカーを囲むコンテキストを取得します。
ディレクティブ挿入 : 分類に基づいて厳密なディレクティブ (棄権、競合解決、ユーザー プロファイル) を使用してプロンプト コンテキストを定式化します。
MindCache は、1M および 10M のトークン コンテキスト ウィンドウでマルチセッションの会話履歴にわたる取得をストレス テストする長期記憶 QA ベンチマークである BEAM でベンチマークされました。
セットアップ: ハイブリッドベクター + BM2

5 回の取得、クロスエンコーダーの再ランカーなし (個別に評価 - 遅延コストが約 23 倍で無視できる程度の精度向上)、要約なし。平均取得遅延: ~1.08 秒。
100 万件のコンテキスト — 10 件の会話、それぞれ最大 20 件の質問
1,000 万のコンテキスト — 5 つの会話、それぞれ最大 20 の質問
10M での会話 2 には、ベンチマーク設計の上限に達する 3 ～ 4 つのベンチマーク エッジ ケース (Lambada スタイルのマルチホップの質問) が含まれていました。それらを除くと、その会話の精度は最大 94% です。
全体: 両方のコンテキスト スケールにわたって 300 の質問が評価されました。
以下の表は、MindCache の結果を、BEAM ベンチマークに関する Mem0 の調査から公式に公開されたベンチマーク スコアと比較しています。
実際のほぼすべての失敗は、両方のコンテキスト サイズにわたって、イベントの順序付けと要約タイプの質問という 2 つのカテゴリに分類されます。どちらも、多数のサポート メモリ チャンクを同時に取得する必要がある広範なクエリです。関連する証拠が多くのノードに分散している場合、上位 K 検索プールが必要なすべての事実を常にカバーするとは限りません。
場合によっては、マルチセッションの追跡、時間的順序付け、および知識の更新で軽微な障害が発生することがあります (1 ～ 2 回)。事実に基づいた具体的な質問 (知識の検索、ユーザーの好み、意思決定) には、正確かつ一貫して回答します。
この評価では概要が有効になっていませんでした。有効になっていないからではありません。

[切り捨てられた]

## Original Extract

Contribute to faisalhussain-devs/MindCache development by creating an account on GitHub.

GitHub - faisalhussain-devs/MindCache · GitHub
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
faisalhussain-devs
/
MindCache
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
collapsed_tree Branches Tags Go to file Code Open more actions menu Folders and files
73 Commits 73 Commits BEAM BEAM eval eval examples examples mindcache mindcache .gitignore .gitignore BLOG.pdf BLOG.pdf README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
Structured Long-Term Memory SDK for LLM Agents
MindCache is a structured long-term memory engine designed for production-grade LLM agents. Unlike flat vector search databases or basic summary stores, MindCache maintains a self-restructuring hierarchical memory ontology with explicit decision tracking, outperforming flat retrieval systems on the BEAM QA benchmark.
📖 Read the Design Story: Detailed details of the architecture design are available in the blog post Building MindCache: Designing an Agentic Memory System for Long-Term AI .
pip install mindcache
2. Usage
from mindcache import MindCache
# Initialize client (SQLite-backed by default)
mc = MindCache (
db_path = "my_memory.db" ,
provider = "gemini" ,
model_name = "gemini-2.5-flash"
)
# 1. Ingest conversation turns (writes to queue in milliseconds)
job_id = mc . add ([
{ "role" : "user" , "content" : "I prefer Python and FastAPI for backend development, and Postgres for DB." },
{ "role" : "assistant" , "content" : "Got it! I will remember your preference for Python, FastAPI, and Postgres." }
], user_id = "alice" )
# 2. Process pending queue (runs extraction, updates tree and indices)
mc . process_queue ( user_id = "alice" )
# 3. Retrieve relevant context for a future prompt
context = mc . search ( "What is Alice's preferred database?" , user_id = "alice" )
print ( context )
🛡️ Key Architectural Differentiators (Why MindCache?)
Most memory systems treat long-term memories as a flat pool of unstructured embedding vectors. MindCache introduces a structured, self-organizing memory architecture optimized for reasoning agents.
🔹 Phase 1 — Ingestion Pipeline
Raw conversations are queued, embedded, and processed in batches. Grounded Routing is performed by querying existing tree paths as constraints before extraction. The Decision State Analyzer and reorganization processes execute later as batch tasks.
flowchart TD
classDef dark fill:#313244,color:#cdd6f4,stroke:#585b70
classDef llm fill:#cba6f7,color:#11111b,stroke:#11111b
classDef tree fill:#89dceb,color:#11111b,stroke:#11111b
classDef yellow fill:#f9e2af,color:#11111b,stroke:#11111b
classDef green fill:#a6e3a1,color:#11111b,stroke:#11111b
classDef analyzer fill:#f38ba8,color:#11111b,stroke:#11111b
A["💬 Raw Turn"]:::dark --> B{"🤖 Worth Remembering?"}:::llm
B -- No --> C["Ignore & Return"]:::dark
B -- Yes --> D["🌲 Fetch Top-K Topic Paths"]:::tree
D --> E["Smart Grounded Routing"]:::llm
E --> F["🤖 Extract & Classify Memories"]:::llm
F --> G["Save to Topic Tree Path"]:::tree
G --> H{"Batch Complete?"}:::dark
H -- Yes --> I["🔍 Analyze Decision States"]:::analyzer
I --> J{"⏱️ Crossed REORG_THRESHOLD?"}:::yellow
J -- No --> K["Refresh Cache & Persist"]:::green
J -- Yes --> L["✂️ Reorganize Tree\n(Leiden Partitioning / Split / Merge)"]:::yellow
L --> M{"Summarization Enabled?"}:::yellow
M -- Yes --> N["📜 Generate Delta Summaries"]:::llm
M -- No --> K
N --> K
Loading
🔹 Phase 2 — Dynamic Tree Lifecycle
The topic tree is never static. When the ingestion count crosses the reorganization threshold, a background process reorganizes the graph and builds incremental bottom-up summaries.
flowchart LR
classDef tree fill:#cba6f7,color:#11111b,stroke:#11111b
classDef reorg fill:#f9e2af,color:#11111b,stroke:#11111b
classDef summary fill:#a6e3a1,color:#11111b,stroke:#11111b
classDef trigger fill:#313244,color:#cdd6f4,stroke:#585b70
T["🌲 Dynamic Topic Tree\n(Hierarchical ontology)"]:::tree
R1["✂️ Split overloaded leaf nodes"]:::reorg
R2["🔗 Merge sparse sibling nodes"]:::reorg
R3["⬆️ Promote / Demote paths"]:::reorg
R4["🌐 Leiden Graph Partitioning"]:::reorg
S["📜 Incremental Delta Summaries\n(Bottom-up delta rollups)"]:::summary
TH["⏱️ Threshold Trigger\n(TriadBlock Count)"]:::trigger
TH --> T
T --> R1 & R2 & R3 & R4
R1 & R2 & R3 & R4 --> T
T -- "New leaf memories" --> S
S -- "Summary stored at parent" --> T
Loading
🔹 Phase 3 — Multi-Stage Hybrid Retrieval Engine
Search queries are processed through a multi-stage preference-anchored pipeline: score generation & RRF, query classification, decision expansion, and directive assembly. The cross-encoder re-ranker was evaluated and removed — it offered marginal accuracy gains at a ~23× latency cost.
flowchart TD
classDef dark fill:#313244,color:#cdd6f4,stroke:#585b70
classDef llm fill:#cba6f7,color:#11111b,stroke:#11111b
classDef search fill:#89b4fa,color:#11111b,stroke:#11111b
classDef anchor fill:#f38ba8,color:#11111b,stroke:#11111b
classDef green fill:#a6e3a1,color:#11111b,stroke:#11111b
Q["🔎 User Query"]:::dark --> A["Context Bridge & Embed"]:::dark
A --> B["Score (Vector + BM25)"]:::search
B --> C["Reciprocal Rank Fusion RRF"]:::search
C --> D["🏷️ Classify Query\n(Broad vs Fact / Temporal)"]:::llm
D --> E["Partitioned Candidate Retrieval"]:::search
E --> F["Memory Pools\n(User / Knowledge / Episodic / Decision)"]:::search
D -->|"If Broad Query"| G["📜 Summaries Pool"]:::search
F & G --> H["⚖️ Get Top 3 Decisions"]:::anchor
H --> I["BM25 Anchor Expansion"]:::anchor
I --> J["Deduplicate Candidates"]:::search
F & J --> K["📋 Inject Directives\n(Profile, Recency, Abstention)"]:::anchor
K --> OUT["📤 Final Context → LLM Prompt"]:::green
Loading
1. Four Structured Memory Types
Rather than storing generic chunks, MindCache segregates information into semantic types:
📝 Episodic : Interactive conversation events, milestones, and session context.
🧠 Knowledge : Fact statements, concepts, and factual domain knowledge.
👤 User : Identity facts, standing preferences, and behavioral attributes.
⚖️ Decision : Explicit choices, directives, and resolutions made by the user.
Decisions evolve over time. MindCache runs a background Decision State Analyzer to identify conflicting or updated decisions. It tracks and flags states as:
Active : Current binding choice.
Superseded : Overridden by a newer decision (automatically suppressed from standard context but preserved for history).
Conditional : Depends on unresolved conditions.
Rejected : User explicitly decided against this option.
3. Decision-Anchored BM25 Expansion
Decisions act as semantic anchors. In retrieval, MindCache identifies the top-ranked Decision memories relevant to the query. It then uses those decisions as expansion anchors to retrieve related episodic, knowledge, and user memories, providing highly aligned background context.
4. Smart Ingestion & Grounded Routing
To prevent the topic tree from fragmenting (where the LLM creates duplicate or slightly different topic paths like FastAPI and FastAPI backend ), MindCache performs Grounded Routing . Before calling the LLM extractor, we query the top-K closest paths from the existing tree and pass them as grounding constraints to guide the LLM's classification.
5. Self-Restructuring Hierarchical Topic Tree
MindCache organizes memories into a tree. However, unlike static hierarchical structures, the tree is living and reorganizes itself in the background:
Splits leaf nodes that grow too large or broad.
Merges sparse sibling branches with semantic overlap.
Promotes / Demotes nodes to optimize path depth.
6. Incremental Delta Summarization
Summaries are built bottom-up to resolve broad queries. MindCache does this incrementally . When new memories are added to a leaf node, we only process the delta (new entries since the last rollup) combined with the existing summary context, keeping LLM API token usage minimal.
7. BM25 Morphological Aliasing (Union-Find)
Standard BM25 lexical search misses word inflections (e.g. database vs databases , run vs running ). MindCache solves this by coupling a Union-Find data structure with spaCy lemmatization. It collapses base words and their inflections into a single equivalence class, combining their indexes in-memory during queries for maximum recall.
MindCache uses a multi-stage search sequence:
Partitioned RRF : Performs vector search + BM25 independently per memory type and merges them via Reciprocal Rank Fusion.
Adaptive Query Classification : Detects Fact vs Broad Overview queries. For broad overview queries where keywords are missing and vector similarity fails, it dynamically pulls RAPTOR-style hierarchical summaries.
Decision-Anchor Expansion : Fans out to retrieve context surrounding top decision anchors.
Directive Injection : Formulates the prompt context using strict directives (Abstention, Conflict Resolution, User Profile) based on classification.
MindCache was benchmarked on BEAM — a long-term memory QA benchmark that stress-tests retrieval across multi-session conversation histories at 1M and 10M token context windows.
Setup: Hybrid Vector + BM25 retrieval, no cross-encoder re-ranker (evaluated separately — negligible accuracy gain at ~23× latency cost), no summarization. Average retrieval latency: ~1.08s .
1 Million context — 10 conversations, ~20 questions each
10 Million context — 5 conversations, ~20 questions each
Conversation 2 at 10M included 3–4 benchmark edge cases (Lambada-style multi-hop questions) that sit at the hard ceiling of the benchmark design. Excluding those, accuracy for that conversation is ~94%.
Overall: 300 questions evaluated across both context scales.
The table below compares MindCache's results against the officially published benchmark scores from Mem0's research on the BEAM benchmark:
Nearly all real failures — across both context sizes — fall into two categories: event ordering and summarization-type questions. Both are broad queries that require retrieving many supporting memory chunks simultaneously. When the relevant evidence is spread across many nodes, the top-K retrieval pool doesn't always cover every required fact.
Occasionally, minor failures occur in multi-session tracking, temporal ordering, and knowledge updates (1 or 2 times). Factual, specific questions (knowledge lookups, user preferences, decisions) answer correctly and consistently.
Summaries were not enabled in this evaluation — not because they don'

[truncated]
