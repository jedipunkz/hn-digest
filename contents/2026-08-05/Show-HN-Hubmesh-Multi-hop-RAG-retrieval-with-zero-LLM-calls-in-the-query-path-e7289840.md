---
source: "https://github.com/DemigodDSK/hubmesh"
hn_url: "https://news.ycombinator.com/item?id=49177513"
title: "Show HN: Hubmesh – Multi-hop RAG retrieval with zero LLM calls in the query path"
article_title: "GitHub - DemigodDSK/hubmesh: Centrality-aware GraphRAG retrieval planner — drop-in layer over any vector DB. Zero LLM in the query path; MCP server included. · GitHub"
author: "dsk_7699"
captured_at: "2026-08-05T01:41:22Z"
capture_tool: "hn-digest"
hn_id: 49177513
score: 1
comments: 0
posted_at: "2026-08-05T01:27:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Hubmesh – Multi-hop RAG retrieval with zero LLM calls in the query path

- HN: [49177513](https://news.ycombinator.com/item?id=49177513)
- Source: [github.com](https://github.com/DemigodDSK/hubmesh)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T01:27:25Z

## Translation

タイトル: HN の表示: Hubmesh – クエリ パス内の LLM 呼び出しがゼロのマルチホップ RAG 取得
記事のタイトル: GitHub - DemigodDSK/hubmesh: 中心性を意識した GraphRAG 取得プランナー — 任意のベクター DB 上のドロップイン層。クエリ パス内の LLM がゼロ。 MCPサーバーが含まれています。 · GitHub
説明: 中心性を意識した GraphRAG 取得プランナー — 任意のベクター DB 上のドロップイン層。クエリ パス内の LLM がゼロ。 MCPサーバーが含まれています。 - DemigodDSK/ハブメッシュ

記事本文:
GitHub - DemigodDSK/hubmesh: 中心性を意識した GraphRAG 取得プランナー — 任意のベクター DB 上のドロップイン層。クエリ パス内の LLM がゼロ。 MCPサーバーが含まれています。 · GitHub
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
デミゴッドDSK
/
ハブメッシュ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
31 コミット 31 コミット .github/ workflows .github/ workflows ベンチマーク ベンチマーク ドキュメント ドキュメント サンプル サンプル src/ ハブメッシュ src/ ハブメッシュ

テスト テスト .gitignore .gitignore BENCHMARKS.md BENCHMARKS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml server.json server.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
中心性を意識した GraphRAG 取得プランナー。任意のベクター DB 上のドロップイン レイヤ。
Hubmesh は、既存の RAG 品質に加えてマルチホップ RAG 品質を向上させる Python ライブラリです。
ベクトルデータベース。インフラストラクチャを置き換えるのではなく、インフラストラクチャとインフラストラクチャの間にスマート プランナーを追加します。
ベクトルDBとLLM。
単純なベクトルの取得 (「埋め込みクエリ、コサイン類似度による上位 K の取得」) がマルチホップで失敗する
「Slack を買収した会社の創設者はどこで生まれたのですか?」などの質問。の
正解するには、最も単一のエンティティを取得するのではなく、推論パスに沿ってエンティティを取得する必要があります。
同様の商品。
GraphRAG と HippoRAG は、ナレッジに対して小さなパーソナライズされた PageRank を実行することを示しました。
クエリ時のグラフにより、マルチホップ検索が大幅に改善されます。ハブメッシュが拡張します
この行には 2 つの貢献があります。
複数コンポーネントのシードの選択。生のクエリによって PPR シードを選択する代わりに
類似性 (特徴の重複が多い場合に間違ったコミュニティ シードを選択する)、シードは
クエリの関連性、構造的適合性、および
報道範囲の多様性。
予算を考慮したコンテキスト パッキング。関連するエンティティがスコア付けされたら、それらをパックします。
明示的なカバレッジと冗長性制御を備えた LLM のコンテキスト ウィンドウ
top-k を切り捨てるだけです。
複数コンポーネントのスコアリング パターンは NNSI フレームワークから採用されています。
SDN トポロジ用 (Naidu Dsk、ICOMP'25 — 登場予定)
最適化。ここでは取得計画のために再利用されます。
インメモリ (テスト、小さなコーパス)
Hubmesh インポート プランナーから
ハブメッシュから。アダプターが InMemoryStore をインポートする
embed = ... # 呼び出し可能: text -> np.ndarray
ドキュメント =

[...] # ドキュメント、文字列、辞書のリスト
ストア = InMemoryStore 。 from_documents ( docs , embed = embed )
planner = プランナー (store = ストア、embed = 埋め込み)
結果 = プランナー。取得 (クエリ = "..." 、top_k = 10、budget_tokens = 4000)
Qdrant アダプター (製品版)
Hubmesh インポート プランナーから
ハブメッシュから。アダプターのインポート QdrantStore
ストア = QdrantStore 。 from_documents ( docs ) # メモリ内
ストア = QdrantStore 。 from_documents ( docs , path = "./qdrant_data" ) # ディスク上
ストア = QdrantStore 。 from_documents ( docs , url = "http://localhost:6333" ) # リモート
planner = プランナー (store = ストア、embed = 埋め込み)
結果 = プランナー。取得 (クエリ = "..." 、top_k = 10)
クロマアダプター
ハブメッシュから。アダプターが ChromaStore をインポートする
ストア = ChromaStore 。 from_documents ( docs ) # 一時的
ストア = ChromaStore 。 from_documents ( docs ,persist_directory = "./chroma_data" )
ストア = ChromaStore 。 from_documents ( docs 、ホスト = "localhost" 、ポート = 8000 )
マルチホップ/KGモード
ハブメッシュから。 kg インポート build_entity_kg
輸入スペーシー
nlp = スペーシー。ロード ( "en_core_web_sm" )
kg = build_entity_kg ( docs 、 nlp = nlp )
planner = プランナー (store = ストア、kg = kg、nlp = nlp)
結果 = プランナー。 retrieve ( query = "Slack を買収した会社の創設者はどこで生まれましたか?" ,
top_k = 10 、budget_tokens = 4000 )
# RetrievalResult には、各ドキュメントが返された理由を示す推論パスが含まれます
結果のパスの場合。推論：
print ( f" スコア= { パス . スコア :.3f } { ' → ' . join ( パス . ノード ID ) } " )
LLM抽出KG（スペイシーより濃厚）
ハブメッシュから。 kg_llm インポート build_entity_kg_llm
ハブメッシュから。 entity_linker import EmbeddingLinker 、 make_st_embedder
def llm (プロンプト): # プロバイダーに依存しない — 独自のプロバイダーを使用する
return your_llm_call (プロンプト)
kg = build_entity_kg_llm (ドキュメント、llm = llm、

キャッシュパス = "kg_cache.json" )
# オプション: クロスドキュメントエンティティ重複排除 — spaCy パスと同じリンカープロトコル
kg = build_entity_kg_llm (ドキュメント、llm = llm、cache_path = "kg_cache.json" ,
リンカー = EmbeddingLinker ( embed = make_st_embedder ()))
planner = プランナー (store = 店舗、kg = kg)
エンティティリンクの改善
ハブメッシュから。 kg インポート build_entity_kg
ハブメッシュから。 entity_linker import EmbeddingLinker 、 make_st_embedder
# クラスター表面のバリエーション: "米国" / "米国" / "USA" → 1 つのエンティティ
リンカー = EmbeddingLinker ( embed = make_st_embedder ()、しきい値 = 0.82 )
kg = build_entity_kg (ドキュメント、リンカー = リンカー)
反復マルチホップ: エージェントが運転できるようにする
r1 = プランナー。取得 (クエリ = 質問、top_k = 5)
# エージェントは r1 を読み取り、ブリッジ エンティティを特定し、ホップ 2 をそこに向けます。
r2 = プランナー。取得 (
クエリ = 質問、top_k = 5、
seed_entities = [ "Nimbus Analytics" ], # クエリ自身のシードとマージされる
exclude_docs = [ s .博士。 r1 の s の ID。ソース ]、 # 消費されたドキュメントを再取得しません
）
シードのメンションはエイリアス インデックスを通じて解決されるため、フリーテキストのエンティティ名が使用されます。
仕事。クエリ パスは決定的で LLM フリーのままです — 計画
発信者の中には知性が宿っています。
MCP サーバー: ハブメッシュを任意のエージェントに接続します
pip インストール「ハブメッシュ[mcp]」
python -m spacy ダウンロード en_core_web_sm
{ "mcpServers" : { "hubmesh" : { "コマンド" : " Hubmesh-mcp " }}}
プランナーを標準入出力上の決定論的なオペレーター ツールとして公開します。
Index_corpus 、retrieve (上記のようにシードステアリング可能)、resolve_entities 、
entity_neighbors 、 path_between 、 get_document 、graph_stats 、
list_corpora 。エージェントは解決者です。質問を分解し、
各ホップを読み取り、次のホップを狙います。サーバーは次のように応答します
LLM 呼び出しがゼロの場合はミリ秒。コーパスはプレーンな JSON/NPZ として保持されます。
~/.hubmesh/corpora の下にあります。

サーバーは、バックグラウンドでモデルと永続化されたコーパスをウォームアップします。
起動 (最初の実行で約 5 ～ 10 秒) なので、ツール呼び出しは最初から高速に保たれます —
厳密なタイムアウトのコネクタ クライアント (Perplexity など) に関連します。
Web ベースのコネクタ クライアントの場合、ゲートウェイなしで SSE をネイティブに提供します
必要なプロセス:
Hubmesh-mcp --transport sse --port 8000 --allow-tunnel
ngrok http 8000 # https://<your-url>/sse をコネクタに貼り付けます
トンネル フィールド ノート (ライブ Perplexity 統合から): ngrok は動作します
(無料利用枠が含まれています); Cloudflared クイック トンネルの SSE ボディのバッファー
そしてツール呼び出しをハングします。ここではスーパーゲートウェイは不要なのでクラッシュします
再接続時。 --allow-tunnel はトンネルの転送されたホストを受け入れます
ヘッダー - これがないと、プロキシされたリクエストは 421 Misdirected Request を受け取ります。
完全なフィールドレポート — セットアップ、エラーデコーダー、9/9 テストバッテリー実行
Perplexity と、推論モデルの動作に関する 2 つの調査結果 —
docs/perplexity.md にあります。
ハブメッシュからインポート chunk_by_sentences 、 chunk_documents
チャンク = chunk_documents (
[{ "id" : "doc1" , "text" : 長文 }, ...],
戦略 = "文" 、ターゲットトークン = 200 、
）
# その後、チャンクを埋め込み、通常どおりインデックスを作成します
インストール
pip インストールハブメッシュ # コア
pip install " Hubmesh[qdrant] " # Qdrant アダプター
pip install " Hubmesh[chroma] " # Chroma アダプター
pip install " Hubmesh[kg] " # エンティティリンクされた KG (spaCy)
pip install " Hubmesh[linker] " # 埋め込みベースのエンティティ リンカー
pip install " Hubmesh[all] " # すべて
python -m spacy download en_core_web_sm # KG モードに必要
デザイン
クエリ → ファーストパス ANN → 誘導サブグラフ → 多成分スコアリング
↓ ↓
コミュニティアンカー → パーソナライズされた PageRank
↓ ↓
└─────→ ランキング → 予算に応じたパッキング → コンテキスト
各層は独立してテスト可能であり、交換可能です。アダプターは既存のベクターをラップします
DB そうそう

移住する必要はありません。
見出し: マルチホップ QA で、hubmesh の KG モードが両方のナイーブ コサインを破る
回収と同じ KG を使用する HippoRAG スタイルの PPR のみのアブレーション、
すべてのホップの深さで。
v0.4.0 のデフォルト (エイリアスインデックス付きシード + NNSI-KG) で測定されたすべての行
収束;ベンチマークでコミットされたアブレーション JSON/ )。開示:
収束は最上位の精度と引き換えに深さの再現を実現します。つまり、recall@2 は次のとおりです。
フル開発では -0.75 ポイント vs ナイーブ (n が小さい場合はディップ ≤0.5)。もしあなたが
top_k=2 で取得し、 use_convergence=False を設定します。マルチシード
クエリのコストは約 1.5 ～ 1.8 倍 (それでも LLM トークンはゼロ、決定的)。
同じ KG での PPR のみのアブレーションとの比較: N=500 で HotpotQA で +29.8 ポイント
(v0.2.0 で測定) — 複数コンポーネントのスコアリングが機能しています。
ただ「グラフがある」だけではありません。
完全な N=7405 HotpotQA 開発者: ハブメッシュは 75.2% に達しました。
リコール@10 対単純コサインの 69.3% (リコール@5 で +4.21 ポイント;
再現率@2 -0.75、上記で開示）。
レイテンシ: 7K ノード KG で平均 ~22 ミリ秒 / クエリあたり 26 ミリ秒 p95 (PPR 後)
マトリックス キャッシュ); 66K パラグラフのフル開発スケールでクエリあたり約 3 秒
v0.4 コンバージェンスがオン。
完全な方法論、アブレーション、
ホップごとの内訳と、これで何が証明され、何が証明されないかについてのメモ。
python ベンチマーク/run_hotpotqa.py --n 500 --kg
Python ベンチマーク/run_musique.py --n 300 --kg
python benchmarks/profile_query.py # レイテンシ プロファイル
ステータス
プレアルファ版 (v0.4.0)。コアアルゴリズムが実装および検証されています。のアダプター
インメモリ、Qdrant、Chroma。 spaCy NER とエンティティリンクされた KG の両方
LLM ベースの抽出 (両方ともリンカー認識)。エイリアスインデックス付きエンティティ解決。
NNSI-KG スコアリング (マルチソース コンバージェンスがデフォルトでオン、ハブ割引 PPR)
オプトイン);エージェント駆動の反復マルチホップ (seed_entities 経由) /
除外ドキュメント ; MCP オペレーター サーバー ( Hubmesh-mcp 、ネイティブ SSE)
JSON/NPZ コーパスの永続化

エンス;ドキュメントのチャンク化。推論パス
説明。 PPR キャッシュ遅延の最適化。松ぼっくり / pgvector / Weaviate アダプター
追加のマルチホップ ベンチマークは次のように追跡されます。
良い最初の問題。
複数コンポーネントのスコアリング パターンは、ネットワーク ノードの重要度から適応されます。
インデックス (NNSI) フレームワークが導入されました
Naidu Dsk、「グラフに基づいてネットワーク トポロジを改善するためのフレームワーク」
Software-Defined Networking の理論」、第 26 回国際会議
インターネット コンピューティング & IoT (ICOMP'25)、ラスベガス、2025 年 7 月 — 議事録
現れること。ここでは、SDN トポロジーの最適化から取得まで再利用されています。
計画中。
中心性を意識した GraphRAG 取得プランナー — 任意のベクター DB 上のドロップイン層。クエリ パス内の LLM がゼロ。 MCPサーバーが含まれています。
pypi.org/project/hubmesh/ トピック
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Centrality-aware GraphRAG retrieval planner — drop-in layer over any vector DB. Zero LLM in the query path; MCP server included. - DemigodDSK/hubmesh

GitHub - DemigodDSK/hubmesh: Centrality-aware GraphRAG retrieval planner — drop-in layer over any vector DB. Zero LLM in the query path; MCP server included. · GitHub
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
DemigodDSK
/
hubmesh
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
31 Commits 31 Commits .github/ workflows .github/ workflows benchmarks benchmarks docs docs examples examples src/ hubmesh src/ hubmesh tests tests .gitignore .gitignore BENCHMARKS.md BENCHMARKS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml server.json server.json View all files Repository files navigation
Centrality-aware GraphRAG retrieval planner. Drop-in layer over any vector DB.
hubmesh is a Python library that improves multi-hop RAG quality on top of an existing
vector database. You don't replace your infrastructure — you add a smart planner between
your vector DB and your LLM.
Naive vector retrieval ("embed query, get top-k by cosine similarity") fails on multi-hop
questions like "Where was the founder of the company that acquired Slack born?" The
correct answer requires retrieving entities along a reasoning path, not the single most
similar item.
GraphRAG and HippoRAG showed that running a small Personalized PageRank over a knowledge
graph at query time can substantially improve multi-hop retrieval. hubmesh extends
that line with two contributions:
Multi-component seed selection. Instead of picking PPR seeds by raw query
similarity (which picks wrong-community seeds at high feature overlap), seeds are
chosen by a multi-component score combining query relevance, structural fit, and
coverage diversity.
Budget-aware context packing. Once relevant entities are scored, pack them into
the LLM's context window with explicit coverage and redundancy control rather than
just truncating top-k.
The multi-component scoring pattern is adapted from the NNSI framework
(Naidu Dsk, ICOMP'25 — to appear) for SDN topology
optimization, repurposed here for retrieval planning.
In-memory (testing, small corpora)
from hubmesh import Planner
from hubmesh . adapters import InMemoryStore
embed = ... # callable: text -> np.ndarray
docs = [...] # list of Document or strings or dicts
store = InMemoryStore . from_documents ( docs , embed = embed )
planner = Planner ( store = store , embed = embed )
result = planner . retrieve ( query = "..." , top_k = 10 , budget_tokens = 4000 )
Qdrant adapter (production)
from hubmesh import Planner
from hubmesh . adapters import QdrantStore
store = QdrantStore . from_documents ( docs ) # in-memory
store = QdrantStore . from_documents ( docs , path = "./qdrant_data" ) # on-disk
store = QdrantStore . from_documents ( docs , url = "http://localhost:6333" ) # remote
planner = Planner ( store = store , embed = embed )
result = planner . retrieve ( query = "..." , top_k = 10 )
Chroma adapter
from hubmesh . adapters import ChromaStore
store = ChromaStore . from_documents ( docs ) # ephemeral
store = ChromaStore . from_documents ( docs , persist_directory = "./chroma_data" )
store = ChromaStore . from_documents ( docs , host = "localhost" , port = 8000 )
Multi-hop / KG mode
from hubmesh . kg import build_entity_kg
import spacy
nlp = spacy . load ( "en_core_web_sm" )
kg = build_entity_kg ( docs , nlp = nlp )
planner = Planner ( store = store , kg = kg , nlp = nlp )
result = planner . retrieve ( query = "Where was the founder of the company that bought Slack born?" ,
top_k = 10 , budget_tokens = 4000 )
# RetrievalResult includes reasoning paths showing why each doc was returned
for path in result . reasoning :
print ( f" score= { path . score :.3f } { ' → ' . join ( path . node_ids ) } " )
LLM-extracted KG (richer than spaCy)
from hubmesh . kg_llm import build_entity_kg_llm
from hubmesh . entity_linker import EmbeddingLinker , make_st_embedder
def llm ( prompt ): # provider-agnostic — bring your own
return your_llm_call ( prompt )
kg = build_entity_kg_llm ( docs , llm = llm , cache_path = "kg_cache.json" )
# optional: cross-document entity dedup — same Linker protocol as the spaCy path
kg = build_entity_kg_llm ( docs , llm = llm , cache_path = "kg_cache.json" ,
linker = EmbeddingLinker ( embed = make_st_embedder ()))
planner = Planner ( store = store , kg = kg )
Better entity linking
from hubmesh . kg import build_entity_kg
from hubmesh . entity_linker import EmbeddingLinker , make_st_embedder
# Cluster surface variations: "United States" / "U.S." / "USA" → one entity
linker = EmbeddingLinker ( embed = make_st_embedder (), threshold = 0.82 )
kg = build_entity_kg ( docs , linker = linker )
Iterative multi-hop: let your agent drive
r1 = planner . retrieve ( query = question , top_k = 5 )
# your agent reads r1, spots the bridge entity, then aims hop 2 at it:
r2 = planner . retrieve (
query = question , top_k = 5 ,
seed_entities = [ "Nimbus Analytics" ], # merged with the query's own seeds
exclude_docs = [ s . doc . id for s in r1 . sources ], # don't re-retrieve consumed docs
)
Seed mentions resolve through the alias index, so free-text entity names
work. The query path stays deterministic and LLM-free — the planning
intelligence lives in the caller.
MCP server: plug hubmesh into any agent
pip install " hubmesh[mcp] "
python -m spacy download en_core_web_sm
{ "mcpServers" : { "hubmesh" : { "command" : " hubmesh-mcp " }}}
Exposes the planner as deterministic operator tools over stdio —
index_corpus , retrieve (seed-steerable, as above), resolve_entities ,
entity_neighbors , path_between , get_document , graph_stats ,
list_corpora . Your agent is the solver: it decomposes the question,
reads each hop, and aims the next one; the server answers in
milliseconds with zero LLM calls. Corpora persist as plain JSON/NPZ
under ~/.hubmesh/corpora .
The server warms up models and persisted corpora in the background at
launch (~5-10s on first run), so tool calls stay fast from the start —
relevant for strict-timeout connector clients (Perplexity, etc.).
For web-based connector clients, serve SSE natively — no gateway
process needed:
hubmesh-mcp --transport sse --port 8000 --allow-tunnel
ngrok http 8000 # paste https://<your-url>/sse into the connector
Tunnel field notes (from a live Perplexity integration): ngrok works
(free tier included); cloudflared quick tunnels buffer SSE bodies
and hang tool calls; supergateway is unnecessary here and crashes
on reconnect. --allow-tunnel accepts the tunnel's forwarded Host
header — without it, proxied requests get 421 Misdirected Request.
Full field report — setup, error decoder, a 9/9 test battery run
through Perplexity, and two findings about reasoning-model behaviour —
in docs/perplexity.md .
from hubmesh import chunk_by_sentences , chunk_documents
chunks = chunk_documents (
[{ "id" : "doc1" , "text" : long_text }, ...],
strategy = "sentences" , target_tokens = 200 ,
)
# Then embed chunks and index normally
Installation
pip install hubmesh # core
pip install " hubmesh[qdrant] " # Qdrant adapter
pip install " hubmesh[chroma] " # Chroma adapter
pip install " hubmesh[kg] " # entity-linked KG (spaCy)
pip install " hubmesh[linker] " # embedding-based entity linker
pip install " hubmesh[all] " # everything
python -m spacy download en_core_web_sm # required for KG mode
Design
query → first-pass ANN → induced subgraph → multi-component scoring
↓ ↓
community anchoring → Personalized PageRank
↓ ↓
└─────→ ranking → budget-aware packing → context
Each layer is independently testable and replaceable. Adapters wrap your existing vector
DB so you don't have to migrate.
Headline: on multi-hop QA, hubmesh's KG mode beats both naive cosine
retrieval and a HippoRAG-style PPR-only ablation that uses the same KG,
at every hop depth.
All rows measured with v0.4.0 defaults (alias-indexed seeds + NNSI-KG
convergence; ablation JSONs committed in benchmarks/ ). Disclosed:
convergence trades top-rank precision for depth recall — recall@2 is
−0.75 pts vs naive on full dev (dips ≤0.5 at smaller n); if you
retrieve with top_k=2 , set use_convergence=False . Multi-seed
queries cost ~1.5–1.8× (still zero LLM tokens, deterministic).
vs PPR-only ablation on the same KG: +29.8 pts on HotpotQA at N=500
(measured on v0.2.0) — the multi-component scoring is doing the work,
not just "having a graph."
On the full N=7405 HotpotQA dev: hubmesh hits 75.2% supporting-fact
recall@10 vs naive cosine's 69.3% (+4.21 pts at recall@5;
recall@2 −0.75, disclosed above).
Latency: ~22 ms mean / 26 ms p95 per query on a 7K-node KG (after PPR
matrix caching); ~3 s/query at the 66K-paragraph full-dev scale with
v0.4 convergence on.
See BENCHMARKS.md for the full methodology, ablations,
per-hop breakdown, and notes on what this proves and doesn't.
python benchmarks/run_hotpotqa.py --n 500 --kg
python benchmarks/run_musique.py --n 300 --kg
python benchmarks/profile_query.py # latency profile
Status
Pre-alpha (v0.4.0). Core algorithms implemented and validated; adapters for
in-memory, Qdrant, and Chroma; entity-linked KG with both spaCy NER and
LLM-based extraction (both linker-aware); alias-indexed entity resolution;
NNSI-KG scoring (multi-source convergence default-on, hub-discounted PPR
opt-in); agent-driven iterative multi-hop via seed_entities /
exclude_docs ; MCP operator server ( hubmesh-mcp , native SSE) with
JSON/NPZ corpus persistence; document chunking; reasoning-path
explanation; PPR-cache latency optimisation. Pinecone / pgvector / Weaviate adapters
and additional multi-hop benchmarks are tracked as
good first issues .
The multi-component scoring pattern is adapted from the Network Node Significance
Index (NNSI) framework introduced in
Naidu Dsk, "A Framework for Improving Network Topology Based on Graph
Theory in Software-Defined Networking", 26th International Conference on
Internet Computing & IoT (ICOMP'25), Las Vegas, July 2025 — proceedings
to appear. Repurposed here from SDN topology optimization to retrieval
planning.
Centrality-aware GraphRAG retrieval planner — drop-in layer over any vector DB. Zero LLM in the query path; MCP server included.
pypi.org/project/hubmesh/ Topics
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
