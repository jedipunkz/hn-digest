---
source: "https://github.com/Ethereal-Agents/arxiv-scholar"
hn_url: "https://news.ycombinator.com/item?id=48554802"
title: "Show HN: ArXiv Scholar – An Open-Source RAG System for AI Research Papers"
article_title: "GitHub - Ethereal-Agents/arxiv-scholar: Retrieval system over the arXiv corpus · GitHub"
author: "dubeyaayush07"
captured_at: "2026-06-16T13:56:11Z"
capture_tool: "hn-digest"
hn_id: 48554802
score: 1
comments: 0
posted_at: "2026-06-16T13:14:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ArXiv Scholar – An Open-Source RAG System for AI Research Papers

- HN: [48554802](https://news.ycombinator.com/item?id=48554802)
- Source: [github.com](https://github.com/Ethereal-Agents/arxiv-scholar)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T13:14:41Z

## Translation

タイトル: Show HN: ArXiv Scholar – AI 研究論文用のオープンソース RAG システム
記事タイトル: GitHub - Ethereal-Agents/arxiv-scholar: arXiv コーパス上の検索システム · GitHub
説明: arXiv コーパス上の検索システム。 GitHub でアカウントを作成して、Ethereal-Agents/arxiv-scholar の開発に貢献してください。
HN テキスト: 検索してみてください: https://ethereal-agents.space/search.html 技術ブログ: https://ethereal-agents.space/blog/launching-arxiv-scholar.h... 検索品質、ユーザー エクスペリエンス、全体的なアプローチに関するフィードバックをお待ちしています。

記事本文:
GitHub - Ethereal-Agents/arxiv-scholar: arXiv コーパス上の検索システム · GitHub
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
エーテルエージェント
/
arxiv-学者
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラン

hes タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
68 コミット 68 コミット .github/ workflows .github/ workflows colab colab configs configs data data docs docs Notebook Notebook scripts scripts src/ arxiv_scholar src/ arxiv_scholar テスト テスト .gitattributes .gitattributes .gitignore .gitignore Dockerfile Dockerfile README.md README.md app.py app.py docker-compose.yml docker-compose.yml main.py main.py pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
タイトル
アルクシブ学者
絵文字
📚
色から
青
色に
紫
SDK
港湾労働者
固定された
偽
ArXiv 研究者
AI エンジニアリング研究用の高性能検索拡張生成 (RAG) システム。
ArXiv Scholar は、arXiv から学術論文を取り込み、解析し、チャンクしてハイブリッド ベクトル データベースに埋め込むエンドツーエンドのパイプラインで、科学文書に対する高速なセマンティック検索を可能にします。完全なアーキテクチャ制御と透過的な障害モードを実現するために、高レベルの抽象化フレームワーク (LangChain なし) を使用せずに最初から構築されています。
ステータス: 約 5,600 件の AI エンジニアリング論文が Qdrant Cloud にインデックス付けされています。ストリーミング API ライブ。エージェント駆動の取り込みによるスケーリングが計画されています。
フローチャート TD
サブグラフ取り込みパイプライン
A[ArxivUnifiedEngine] -->|GCS から PDF をダウンロード| B(ローカルディレクトリリーダー/GCSBucketリーダー)
B -->|収益書類| C(レイアウト対応チャンカー)
C -->|チャンクを生成| D{埋め込みサービス}
D -->|密なベクトル| E[FastEmbedEmbedder<br/><small>BAAI/bge-m3</small>]
D -->|スパースベクトル| F[SparseBM25Embedder<br/><small>Qdrant/bm25</small>]
E --> G[(Qdrant Cloud)]
F --> G
終わり
サブグラフ取得パイプライン
H[ユーザー クエリ] --> I{ML クエリ ルーター}
私 -->|シンプル| J[ダイレクトハイブリッド検索]
I -->|複雑/比較| K[LLM 分解 + メタデータ抽出]
私 -->|概念的| L[HyDE - 仮説的なドキュメントの埋め込み

ディン]
K -->|サブクエリ + フィルター| J
L -->|生成された要約| J
J -->|密 + 疎フェッチ| G
G -->|最小-最大正規化 + 加重融合| N[最終結果]
終わり
読み込み中
再ランキングに関する注意: クロスエンコーダーの再ランカー ( jina-reranker-v1-tiny-en ) はコードベースに実装されていますが、デフォルトでは無効になっています ( USE_RERANKER=False )。ベンチマーク中に、リランカーが現在のコーパス サイズのパフォーマンス低下を引き起こしたため、オフになりました。コードは、将来の大規模な評価のために保持されます。
ステージ
コンポーネント
説明
ダウンロード
Arxiv統合エンジン
構成可能なバッチでパブリック arxiv-dataset GCS バケットから PDF をストリーミングします。 YYMM フォルダー間で再開可能でクラッシュセーフな取り込みを行うために、JSON カーソル ( current_month 、 last_file ) を維持します。
解析中
LocalDirectoryReader / GCSBucketReader
PyMuPDF を介して PDF から生のテキストを抽出します。重複排除のために SHA-256 ハッシュを計算し、正規表現を使用してファイル名から arXiv ID を抽出します。 GCS リーダーは、サーバーレス展開の場合は完全にインメモリで動作します。
チャンク化
レイアウトを意識したチャンカー
Docling を使用して PDF レイアウト (ヘッダー、段落、表) を視覚的に解析し、意味的にグループ化されたチャンクを生成します。ブロックが大きすぎる場合、または Docling が使用できない場合は、SlidingWindowChunker にフォールバックします。
埋め込み
文TransformerEmbedder + SparseBM25Embedder
SentenceTransformers (PyTorch) を介して密なベクトル (BAAI/bge-m3) を生成し、FastEmbed (ONNX) を介して疎な BM25 ベクトルを同時に生成します。
ストレージ
QdrantVectorStore
決定論的な UUID-v5 ポイント ID を持つチャンクを Qdrant Cloud に更新/挿入します。テスト用にクラウド モード (URL + API キー) とインメモリ モードの両方をサポートします。
検索
ハイブリッドレトリバー
密と疎の結果を個別に取得し、最小-最大正規化を適用し、スコアを構成可能な重みと融合します (デフォルト: 密 = 1.0、疎 = 0.3)。
主な特長
ハイブリッド検索 — 巣を結合します

スパース BM25 キーワード マッチングを使用したセマンティック埋め込みは、重み付けされた最小-最大正規化を介して融合され、いずれかの方法単独よりも優れた再現率を実現します。
インテリジェントなクエリ ルーティング — ハイブリッド ML + ヒューリスティック ルーター (<1ms) は、受信クエリを Direct、Decompose、または HyDE パスに分類します。 ML 幻覚なしで保証されたメタデータ ルーティング (年フィルタリングなど) を実現する正規表現ベースのハード オーバーライドに加え、HyDE に自動ルーティングするショート クエリ検出が含まれます。
LLM を利用したクエリ分解 — 複雑なクエリは、LLM から JSON 経由で抽出された厳密なメタデータ フィルター (出版年など) を使用してアトミックなサブクエリに分割されます。フィルターは Qdrant プリフェッチ レベルでネイティブに適用されます。
動的なコンピューティング予算設定 — 分解によるサブクエリは同時にフェッチされ、グローバルな重複排除とスコアリングの前にプールされます。フェッチ バジェットはサブクエリ全体に動的に割り当てられます。
レイアウトを意識した PDF 解析 — Docling ベースの視覚的な文書理解により、単純なテキスト分割ではなく、学術論文の意味構造 (セクション、表、方程式) が保持されます。
クラッシュセーフなバッチ取り込み — カーソルベースの状態管理により、大規模な取り込み実行全体にわたって、パイプラインが障害が発生した正確な時点から再開できます。
ストリーミング API — Server-Sent Events (SSE) を備えた FastAPI エンドポイントは、取得したソースと LLM で合成された応答をトークンごとにストリームします。
arxiv-学者/
§── main.py # フルインジェストパイプラインオーケストレーター
§── app.py # Streamlit チャット UI
§── configs/
│ └─ config.py # env-var に基づく一元的な設定
§── src/arxiv_scholar/
│ §── schema.py # コアデータモデル（ドキュメント、チャンク）
│ §── api/
│ │ §── schema.py # REST API リクエスト/レスポンスモデル (SSE イベント)
│ │ ━──server.py # FastAPI ストリーミング エンドポイント (POST

/api/v1/クエリ)
│ §── チャンク化/
│ │ §──base.py # 抽象的な BaseChunker インターフェイス
│ │ §──layout.py # Docling ベースのレイアウト対応チャンカー
│ │ └── slide_window.py # 固定サイズのスライディング ウィンドウ フォールバック チャンカー
│ §── ダウンロード/
│ │ └── arxiv_ingestion.py # カーソル状態を含む GCS 支援の PDF ダウンローダー
│ §── 埋め込み/
│ │ §──base.py # 抽象的な BaseEmbedder インターフェイス
│ │ §── fastembed_embedder.py # ONNX CPU エンベッダー (密 + スパース BM25)
│ │ └─ st_embedder.py # SentenceTransformer エンベッダー (GPU)
│ §── 摂取/
│ │ §──base.py # 抽象的な DocumentReader インターフェイス
│ │ §── local.py # ローカルファイルシステム PDF リーダー (PyMuPDF)
│ │ └── gcs.py # インメモリ GCS バケットリーダー (サーバーレス)
│ §── llm/
│ │ ━─ service.py # LLM クライアント (分解、HyDE、合成)
│ §── 検索/
│ │ §── retrieval.py # ウェイトフュージョンを備えたハイブリッドレトリーバー
│ │ §── Orchestrator.py # オーケストレーターのクエリ (ルート → 取得 → ヒューズ)
│ │ └── router.py # ML + ヒューリスティッククエリルーター
│ └── 保管/
│ §──base.py # 抽象的な BaseVectorStore インターフェイス
│ └─ qdrant_store.py # Qdrant クライアント (upsert、検索、ハイブリッド検索)
§── スクリプト/
│ §──generate_arxiv_manifest.py # 用紙選択基準とマニフェストジェネレーター
│ §── download_qdrant.sh # Qdrant バイナリインストーラー
│ §──generate_eval_dataset.py # 評価用データセット生成器
│ §── run_benchmarks.py # 取得ベンチマークランナー
│ └── ... # さまざまな取り込みおよびインポート ユーティリティ
§── コラボ/
│ §──batch_gcs_to_drive.py # Colab スクリプト: GCS からのバッチダウンロード
│ ━──generate_embedded_dataset.py #

Colab スクリプト: Qdrant Cloud に埋め込んでプッシュする
§── ノートブック/ # 開発とテスト用の Jupyter ノートブック
§─
[切り捨てられた]
レイヤー
テクノロジー
目的
高密度埋め込み
SentenceTransformers 経由の BAAI/bge-m3 (PyTorch)
意味ベクトル
スパース埋め込み
FastEmbed 経由の Qdrant/bm25
キーワード マッチング用の BM25 用語頻度ベクトル
ベクターデータベース
Qdrant クラウド
サーバー側のクエリバッチ処理を備えたハイブリッドストレージ
PDF の解析
PyMuPDF + ドクリング
テキスト抽出とレイアウトを意識したチャンク化
API
FastAPI + ユビコーン
ストリーミング SSE エンドポイント
LLM
設定可能（OpenAI互換API）
クエリ分解、HyDE生成、回答合成
クエリルーター
scikit-learn + 正規表現ヒューリスティック
ミリ秒未満のクエリ分類
オーケストレーション
純粋な Python (LangChain なし)
完全なアーキテクチャ制御
はじめに
# リポジトリのクローンを作成します
git クローン https://github.com/dubeyaayush07/arxiv-scholar.git
cd arxiv-scholar
# 仮想環境を作成し、依存関係をインストールする
uv venv && ソース .venv/bin/activate
uv pip install -e 。
環境変数
.env ファイルを作成するか、次のものをエクスポートします。
# 必須 — Qdrant クラウド接続
export QDRANT_URL= " your_qdrant_cloud_url "
import QDRANT_API_KEY= " your_qdrant_api_key "
import QDRANT_COLLECTION= " Arxiv-Scholar "
# LLM 機能 (分解、HyDE、応答合成) に必要
export LLM_API_KEY= " your_key_here "
import LLM_BASE_URL= " https://generative language.googleapis.com/v1beta/openai/ " # または任意の OpenAI 互換エンドポイント
エクスポート LLM_MODEL= " クロード-俳句-4-5 "
# オプションのオーバーライド (デフォルトを表示)
import EMBEDDING_BACKEND= " fastembed " # または "sentence-transformers"
エクスポート EMBEDDING_MODEL= " BAAI/bge-m3 "
エクスポート SPARSE_EMBEDDING_MODEL= " Qdrant/bm25 "
import USE_RERANKER= " False " # 無効 — パフォーマンスの低下を引き起こします
エクスポート DENSE_WEIGHT= " 0.6

」
エクスポート SPARSE_WEIGHT= " 0.4 "
# ローカル Qdrant の場合 (クラウドの代替)
エクスポート QDRANT_HOST= " ローカルホスト "
エクスポート QDRANT_PORT= " 6333 "
使用法
完全な取り込みパイプラインは main.py に実装されています。 arXiv GCS バケットから PDF をダウンロードし、Docling で解析し、チャンク化し、デュアル エンベディングを生成し、Qdrant に更新/挿入します。
# 試行実行 (2 つの PDF をダウンロード、メモリ内で Qdrant を処理)
python main.py --trial
# 実稼働実行 (継続的なバッチ取り込み)
Python main.py
実際のデータの取り込み方法: ローカル コンピューティングの制約のため、最初の約 5,600 件の紙のコーパスは Google Colab 経由で取り込まれました。 colab/batch_gcs_to_drive.py スクリプトは GCS から PDF をバッチ処理し、colab/generate_embedded_dataset.py はエンベディングを生成して Qdrant Cloud にプッシュします。正確な論文選択基準 (キーワード グループ、ドメインの除外、ゴールデン ターム) については、scripts/generate_arxiv_manifest.py を参照してください。
API サーバーは src/arxiv_scholar/api/server.py に実装されています。これは、LLM を介してクエリをルーティングし、結果を取得し、回答を合成するストリーミング SSE エンドポイントを公開します。
ライブホスト型エンドポイント (Hugging Face Spaces):
curl -N -X POST " https://trinetra-dev-arxiv-scholar.hf.space/api/v1/query " \
-H " Content-Type: application/json " \
-d ' {
"query": "対照学習とは何ですか?",
「制限」: 5、
"use_reranker": false
} '
ローカルで実行します。
# サーバーを起動します
uvicorn src.arxiv_scholar.api.server:app --reload
# または Docker 経由
docker build -t arxiv-scholar 。
docker run -p 7860:7860 --env-file .env arxiv-scholar
# ローカルインスタンスにクエリを実行します
curl -N -X POST http://localhost:8000/api/v1/query \
-H " コンテンツ タイプ: ap

[切り捨てられた]

## Original Extract

Retrieval system over the arXiv corpus. Contribute to Ethereal-Agents/arxiv-scholar development by creating an account on GitHub.

Try Search: https://ethereal-agents.space/search.html Technical Blog: https://ethereal-agents.space/blog/launching-arxiv-scholar.h... We'd love feedback on the retrieval quality, user experience, and overall approach.

GitHub - Ethereal-Agents/arxiv-scholar: Retrieval system over the arXiv corpus · GitHub
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
Ethereal-Agents
/
arxiv-scholar
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
68 Commits 68 Commits .github/ workflows .github/ workflows colab colab configs configs data data docs docs notebooks notebooks scripts scripts src/ arxiv_scholar src/ arxiv_scholar tests tests .gitattributes .gitattributes .gitignore .gitignore Dockerfile Dockerfile README.md README.md app.py app.py docker-compose.yml docker-compose.yml main.py main.py pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
title
Arxiv Scholar
emoji
📚
colorFrom
blue
colorTo
purple
sdk
docker
pinned
false
ArXiv Scholar
A high-performance Retrieval-Augmented Generation (RAG) system for AI Engineering research.
ArXiv Scholar is an end-to-end pipeline that ingests, parses, chunks, and embeds academic papers from arXiv into a hybrid vector database — enabling fast semantic search over scientific documents. Built from scratch without high-level abstraction frameworks (no LangChain) for full architectural control and transparent failure modes.
Status: ~5,600 AI Engineering papers indexed on Qdrant Cloud. Streaming API live. Scaling via agent-driven ingestion planned.
flowchart TD
subgraph Ingestion Pipeline
A[ArxivUnifiedEngine] -->|Downloads PDFs from GCS| B(LocalDirectoryReader / GCSBucketReader)
B -->|Yields Documents| C(LayoutAwareChunker)
C -->|Produces Chunks| D{Embedding Service}
D -->|Dense Vectors| E[FastEmbedEmbedder<br/><small>BAAI/bge-m3</small>]
D -->|Sparse Vectors| F[SparseBM25Embedder<br/><small>Qdrant/bm25</small>]
E --> G[(Qdrant Cloud)]
F --> G
end
subgraph Retrieval Pipeline
H[User Query] --> I{ML Query Router}
I -->|Simple| J[Direct Hybrid Search]
I -->|Complex/Comparative| K[LLM Decomposition + Metadata Extraction]
I -->|Conceptual| L[HyDE - Hypothetical Document Embedding]
K -->|Sub-queries + Filters| J
L -->|Generated Abstract| J
J -->|Dense + Sparse Fetch| G
G -->|Min-Max Normalized + Weighted Fusion| N[Final Results]
end
Loading
Note on reranking: A cross-encoder reranker ( jina-reranker-v1-tiny-en ) is implemented in the codebase but is disabled by default ( USE_RERANKER=False ). During benchmarking, the reranker caused performance degradation on the current corpus size and was turned off. The code is retained for future evaluation at larger scale.
Stage
Component
Description
Download
ArxivUnifiedEngine
Streams PDFs from the public arxiv-dataset GCS bucket in configurable batches. Maintains a JSON cursor ( current_month , last_file ) for resumable, crash-safe ingestion across YYMM folders.
Parsing
LocalDirectoryReader / GCSBucketReader
Extracts raw text from PDFs via PyMuPDF. Computes SHA-256 hashes for deduplication and extracts arXiv IDs from filenames using regex. GCS reader operates fully in-memory for serverless deployments.
Chunking
LayoutAwareChunker
Uses Docling to visually parse PDF layouts (headers, paragraphs, tables) and produce semantically grouped chunks. Falls back to SlidingWindowChunker for oversized blocks or when Docling is unavailable.
Embedding
SentenceTransformerEmbedder + SparseBM25Embedder
Generates dense vectors (BAAI/bge-m3) via SentenceTransformers (PyTorch) and sparse BM25 vectors via FastEmbed (ONNX) concurrently.
Storage
QdrantVectorStore
Upserts chunks with deterministic UUID-v5 point IDs to Qdrant Cloud. Supports both cloud mode (URL + API key) and in-memory mode for testing.
Retrieval
HybridRetriever
Fetches dense and sparse results independently, applies min-max normalization, and fuses scores with configurable weights (default: dense=1.0, sparse=0.3).
Key Features
Hybrid Search — Combines dense semantic embeddings with sparse BM25 keyword matching, fused via weighted min-max normalization for superior recall over either method alone.
Intelligent Query Routing — A hybrid ML + heuristic router (<1ms) classifies incoming queries into Direct, Decompose, or HyDE paths. Includes regex-based Hard Overrides for guaranteed metadata routing (e.g., year filtering) without ML hallucinations, plus short-query detection that auto-routes to HyDE.
LLM-Powered Query Decomposition — Complex queries are split into atomic sub-queries with strict metadata filters (e.g., publication year) extracted via JSON from an LLM. Filters are applied natively at the Qdrant Prefetch level.
Dynamic Compute Budgeting — Sub-queries from decomposition are fetched concurrently and pooled before global deduplication and scoring. The fetch budget is dynamically allocated across sub-queries.
Layout-Aware PDF Parsing — Docling-based visual document understanding preserves the semantic structure of academic papers (sections, tables, equations) instead of naive text splitting.
Crash-Safe Batch Ingestion — Cursor-based state management allows the pipeline to resume from the exact point of failure across large ingestion runs.
Streaming API — FastAPI endpoint with Server-Sent Events (SSE) streams retrieved sources and LLM-synthesized answers token-by-token.
arxiv-scholar/
├── main.py # Full ingestion pipeline orchestrator
├── app.py # Streamlit chat UI
├── configs/
│ └── config.py # Centralized env-var-backed configuration
├── src/arxiv_scholar/
│ ├── schema.py # Core data models (Document, Chunk)
│ ├── api/
│ │ ├── schema.py # REST API request/response models (SSE events)
│ │ └── server.py # FastAPI streaming endpoint (POST /api/v1/query)
│ ├── chunking/
│ │ ├── base.py # Abstract BaseChunker interface
│ │ ├── layout.py # Docling-based layout-aware chunker
│ │ └── sliding_window.py # Fixed-size sliding window fallback chunker
│ ├── download/
│ │ └── arxiv_ingestion.py # GCS-backed PDF downloader with cursor state
│ ├── embedding/
│ │ ├── base.py # Abstract BaseEmbedder interface
│ │ ├── fastembed_embedder.py # ONNX CPU embedder (dense + sparse BM25)
│ │ └── st_embedder.py # SentenceTransformer embedder (GPU)
│ ├── ingestion/
│ │ ├── base.py # Abstract DocumentReader interface
│ │ ├── local.py # Local filesystem PDF reader (PyMuPDF)
│ │ └── gcs.py # In-memory GCS bucket reader (serverless)
│ ├── llm/
│ │ └── service.py # LLM client (decomposition, HyDE, synthesis)
│ ├── retrieval/
│ │ ├── retrieval.py # Hybrid retriever with weighted fusion
│ │ ├── orchestrator.py # Query orchestrator (routes → retrieves → fuses)
│ │ └── router.py # ML + heuristic query router
│ └── storage/
│ ├── base.py # Abstract BaseVectorStore interface
│ └── qdrant_store.py # Qdrant client (upsert, search, hybrid search)
├── scripts/
│ ├── generate_arxiv_manifest.py # Paper selection criteria & manifest generator
│ ├── download_qdrant.sh # Qdrant binary installer
│ ├── generate_eval_dataset.py # Evaluation dataset generator
│ ├── run_benchmarks.py # Retrieval benchmark runner
│ └── ... # Various ingestion and import utilities
├── colab/
│ ├── batch_gcs_to_drive.py # Colab script: batch download from GCS
│ └── generate_embedded_dataset.py # Colab script: embed and push to Qdrant Cloud
├── notebooks/ # Jupyter notebooks for development & testing
├─
[truncated]
Layer
Technology
Purpose
Dense Embedding
BAAI/bge-m3 via SentenceTransformers (PyTorch)
Semantic vectors
Sparse Embedding
Qdrant/bm25 via FastEmbed
BM25 term-frequency vectors for keyword matching
Vector Database
Qdrant Cloud
Hybrid storage with server-side query batching
PDF Parsing
PyMuPDF + Docling
Text extraction and layout-aware chunking
API
FastAPI + Uvicorn
Streaming SSE endpoint
LLM
Configurable (OpenAI-compatible API)
Query decomposition, HyDE generation, answer synthesis
Query Router
scikit-learn + regex heuristics
Sub-millisecond query classification
Orchestration
Pure Python (no LangChain)
Full architectural control
Getting Started
# Clone the repository
git clone https://github.com/dubeyaayush07/arxiv-scholar.git
cd arxiv-scholar
# Create a virtual environment and install dependencies
uv venv && source .venv/bin/activate
uv pip install -e .
Environment Variables
Create a .env file or export the following:
# Required — Qdrant Cloud connection
export QDRANT_URL= " your_qdrant_cloud_url "
export QDRANT_API_KEY= " your_qdrant_api_key "
export QDRANT_COLLECTION= " Arxiv-Scholar "
# Required for LLM features (decomposition, HyDE, answer synthesis)
export LLM_API_KEY= " your_key_here "
export LLM_BASE_URL= " https://generativelanguage.googleapis.com/v1beta/openai/ " # or any OpenAI-compatible endpoint
export LLM_MODEL= " claude-haiku-4-5 "
# Optional overrides (defaults shown)
export EMBEDDING_BACKEND= " fastembed " # or "sentence-transformers"
export EMBEDDING_MODEL= " BAAI/bge-m3 "
export SPARSE_EMBEDDING_MODEL= " Qdrant/bm25 "
export USE_RERANKER= " False " # Disabled — causes performance degradation
export DENSE_WEIGHT= " 0.6 "
export SPARSE_WEIGHT= " 0.4 "
# For local Qdrant (alternative to cloud)
export QDRANT_HOST= " localhost "
export QDRANT_PORT= " 6333 "
Usage
The full ingestion pipeline is implemented in main.py . It downloads PDFs from the arXiv GCS bucket, parses them with Docling, chunks them, generates dual embeddings, and upserts to Qdrant.
# Trial run (downloads 2 PDFs, processes in-memory Qdrant)
python main.py --trial
# Production run (continuous batch ingestion)
python main.py
How we actually ingested data: Due to local compute constraints, the initial ~5,600 paper corpus was ingested via Google Colab. The colab/batch_gcs_to_drive.py script batches PDFs from GCS, and colab/generate_embedded_dataset.py generates embeddings and pushes them to Qdrant Cloud. See scripts/generate_arxiv_manifest.py for the exact paper selection criteria (keyword groups, domain exclusions, golden terms).
The API server is implemented in src/arxiv_scholar/api/server.py . It exposes a streaming SSE endpoint that routes queries, retrieves results, and synthesizes answers via LLM.
Live hosted endpoint (Hugging Face Spaces):
curl -N -X POST " https://trinetra-dev-arxiv-scholar.hf.space/api/v1/query " \
-H " Content-Type: application/json " \
-d ' {
"query": "What is contrastive learning?",
"limit": 5,
"use_reranker": false
} '
Run locally:
# Start the server
uvicorn src.arxiv_scholar.api.server:app --reload
# Or via Docker
docker build -t arxiv-scholar .
docker run -p 7860:7860 --env-file .env arxiv-scholar
# Query your local instance
curl -N -X POST http://localhost:8000/api/v1/query \
-H " Content-Type: ap

[truncated]
