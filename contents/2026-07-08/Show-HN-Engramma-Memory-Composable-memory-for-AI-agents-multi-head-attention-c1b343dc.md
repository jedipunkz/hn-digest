---
source: "https://github.com/engramma-ai/engramma-memory"
hn_url: "https://news.ycombinator.com/item?id=48832313"
title: "Show HN: Engramma Memory – Composable memory for AI agents(multi-head attention)"
article_title: "GitHub - engramma-ai/engramma-memory: A composable memory engine for AI agents. Unlike vector databases that only retrieve, Engramma natively composes, generalizes, and adapts powered by a hybrid architecture combining exact kNN, Hopfield energy networks, and multi-head attention. One pip install. Z\n[truncated]"
author: "riscoss"
captured_at: "2026-07-08T15:12:04Z"
capture_tool: "hn-digest"
hn_id: 48832313
score: 1
comments: 0
posted_at: "2026-07-08T14:18:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Engramma Memory – Composable memory for AI agents(multi-head attention)

- HN: [48832313](https://news.ycombinator.com/item?id=48832313)
- Source: [github.com](https://github.com/engramma-ai/engramma-memory)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T14:18:48Z

## Translation

タイトル: Show HN: Engramma Memory – AI エージェント用のコンポーザブル メモリ (マルチヘッド アテンション)
記事のタイトル: GitHub - engramma-ai/engramma-memory: AI エージェント用のコンポーザブル メモリ エンジン。検索のみを行うベクトル データベースとは異なり、Engramma は、正確な kNN、ホップフィールド エネルギー ネットワーク、およびマルチヘッド アテンションを組み合わせたハイブリッド アーキテクチャを活用して、ネイティブに構成、一般化、適応します。 1 つの pip インストール。 Z
[切り捨てられた]
説明: AI エージェント用のコンポーザブル メモリ エンジン。検索のみを行うベクトル データベースとは異なり、Engramma は、正確な kNN、ホップフィールド エネルギー ネットワーク、およびマルチヘッド アテンションを組み合わせたハイブリッド アーキテクチャを活用して、ネイティブに構成、一般化、適応します。 1 つの pip インストール。 NumPy を超える依存関係はありません。製品
[切り捨てられた]

記事本文:
GitHub - engramma-ai/engramma-memory: AI エージェント用のコンポーザブル メモリ エンジン。検索のみを行うベクトル データベースとは異なり、Engramma は、正確な kNN、ホップフィールド エネルギー ネットワーク、およびマルチヘッド アテンションを組み合わせたハイブリッド アーキテクチャを活用して、ネイティブに構成、一般化、適応します。 1 つの pip インストール。 NumPy を超える依存関係はありません。 Engramma Cloud で本番環境に対応。 · GitHub
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
別のタブまたはウィンドウでサインアウトしました。リ

ロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
エングラマアイ
/
エングラマメモリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github/ workflows .github/ workflows docs docs engramma_memory engramma_memory 例 例 ノートブック ノートブック real_tests real_tests テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
単に検索するだけではなく、思考する AI 用のメモリ エンジン。
構成。一般化。因果推論。 1 pip でインストールできます。
ベクトルデータベースが取得します。エングラマが構成します。
エージェントは「Python と機械学習について何を知っていますか?」と尋ねます。
ChromaDB は 2 つの別々の結果を返します。 Engramma は 1 つの融合された回答を返します。
今日のすべての AI メモリ システムは単なる検索です。最も近いベクトルを見つけてそれを返します。それは記憶ではなく検索エンジンです。実メモリはさらに多くのことを行います。
🧩 作文 — 「X と Y の交点は何ですか?」 → 一つの一貫した答え
🧠 一般化 — ノイズの多い入力でも適切なパターンがトリガーされる
📈 適応 — 頻繁にアクセスされるパターンがより強力になります
🗑️ 忘れる — 時代遅れのパターンは自然に衰退します
Engramma は 4 つすべてを実行します。 3 行のコードで。
pip インストール engramma-memory
基本的な使い方
numpyをnpとしてインポート
engramma_memory から EngrammaMemory をインポート
# ローカルの純粋なメモリ内エンジンを初期化する
mem = エングラマメモリ

( dim = 256 、バックエンド = "ローカル" )
# 知識を蓄える
メモリ。 store ( key = embedding_a 、 value = "Python はプログラミング言語です" )
メモリ。 store ( key = embedding_b , value = "機械学習は学習にデータを使用します" )
# 取得 — 3 つの経路にわたるスマートなルーティング
結果 = mem 。取得 (クエリ埋め込み)
# Compose — キラー機能 ⚡
ブレンド = mem . compose ([ embedding_a , embedding_b ])
# ネイティブのマルチヘッド アテンション フュージョンは一貫した応答を返します
注記
それだけです。設定ファイルはありません。ドッカーはありません。 API キーはありません。ただしびれる。
🆚 Vector DB を使用しないのはなぜですか?
従来のベクトル データベースを使用して概念を組み合わせる場合、複数の結果を取得して手動で結合する必要があります。 Engramma はネイティブ合成でこれを解決します。
# ChromaDB / Pinecone / FAISS — 2 つの別々の結果が得られます
結果_a = db 。 query ( key_a ) # 「Python は言語です...」
結果_b = db 。 query ( key_b ) # "ML はデータを使用して学習します..."
＃さてどうする？それらを平均すると？連結しますか?祈る？
Blend = ( result_a + result_b ) / 2 # 意味のない算術
エングラマのやり方
# Engramma — 融合された答えが 1 つ得られます
ブレンド = mem . compose ([ key_a , key_b ])
# 各ヘッドの専門分野: あるものは A を思い出し、あるものは B を思い出す → コヒーレントフュージョン
特徴
従来のベクター DB
エングラマ
🔍 最近傍検索
✅
✅
🧩 ネイティブ構成
❌
✅
🧠 ソフト一般化 (ホップフィールド)
❌
✅
🔀 アダプティブルーティング
❌
✅
📉 重要度に基づく立ち退き
❌
✅
🍂徐々に忘れていく
❌
✅
📦 依存関係ゼロ
❌
✅ (numpy のみ)
🏗️ 仕組み
Engramma は、マルチパスウェイ アーキテクチャを使用して、タスクに基づいてクエリをインテリジェントにルーティングします。
グラフTD
Q[クエリ] --> 正確[正確な kNN メモリ]
Q --> エネルギー[エネルギーホップフィールド]
Q --> MHA[マルチヘッド アテンション]
Exact --> Router[Confidence Router<br/>学習した重み]
エネルギー -->

ルーター
MHA --> ルーター
ルーター --> 最良[最良の結果]
読み込み中
正確な記憶 — 重要度スコアリングを備えた kNN による完全な記憶
エネルギーメモリ — 温度スケールのホップフィールドダイナミクスによるソフトな一般化
マルチヘッドの注意 — 各ヘッドは異なるパターンに注意を払います → ネイティブの構成
Confidence Router — どのパスウェイがどのクエリ タイプを処理するかを学習します
すべての学習はローカル (ヘビアン) です。バックプロパゲーションはありません。 GPUは必要ありません。純粋なNumPy。
Engramma は、ほんの少しの素の速度と引き換えに、作曲能力を大幅に向上させます。
Engramma は、生のスピードと引き換えに合成機能を備えています。数百万のベクトルでの純粋な最近傍の場合は、FAISS を使用します。記憶力を使って考える必要がある AI エージェントには、Engramma を使用します。
Engramma は既存の AI スタックに直接組み込まれます。
engramma_memory から。統合。 langchain import EngrammaLangChainMemory
メモリ = EngrammaLangChainMemory ( dim = 256 、 embed_fn = fn )
ラマインデックス
engramma_memory から。統合。 llamaindexインポートEngrammaRetriever
retriever = EngrammaRetriever ( dim = 256 、embed_fn = fn )
OpenAI アシスタント
engramma_memory から。統合。 openai_assistants インポート engramma_tool_settings
tools = engramma_tool_settings ()
ファストAPI
engramma_memory から。統合。 fastapiインポートcreate_memory_router
アプリ。 include_router ( create_memory_router ( dim = 256 ))
☁️エングラマクラウド
同じ API。制限はありません。 43 のプレミアム機能。
# ローカル (無料、オープンソース、1000 パターンに制限)
mem = EngrammaMemory ( dim = 256 、バックエンド = "ローカル" )
# クラウド (無制限、永続的、インテリジェント) — 同じコードを 1 行変更!
mem = EngrammaMemory ( dim = 256 、バックエンド = "cloud" 、api_key = "nx_live_..." )
重要
特徴
ローカル（無料）
クラウド
🗃️ 最大パターン
1,000
無制限
💾 ストレージ
RAMのみ
段階的（温/温/冷）
⚖️構成

重み
等しいのみ
カスタム分数 (0.0 ～ 1.0)
🛡️ 持続性
なし（処理中）
耐久性 + スナップショット
🧠 ルーティング
信頼に基づく
アクティブ推論 + phi_B
🔗 因果推論
—
DAG の発見 + 介入
🚨 異常検出
—
3つの安全システム
🔮 時間的予測
—
グレンジャー因果関係 + プリフェッチ
💬 テキストインターフェース
—
HDC トークナイザー (埋め込みは必要ありません)
🔬 説明可能性
—
完全な XAI ダッシュボード
クラウド機能のハイライト
因果推論と安全性
# 因果構造を発見する
グラフ = mem 。 get_causal_graph ()
# 「A を変更すると、B はどうなりますか?」
効果 = mem 。予測原因効果 (原因キー = a 、効果キー = b )
# 部分構成 (50/50 だけではない)
ブレンド = mem . compose_fractional ( a , b , alpha = 0.7 )
# 危険な OOD 構成を自動ブロックする
メモリ。 Enable_anomaly_protection (有効 = True)
テキストの記憶と説明可能性
# 自然言語で保存 (埋め込みは必要ありません!)
メモリ。 store_text ( "ユーザーは JS よりも Python を好む" )
# 自然言語によるクエリ
結果 = mem 。 query_text ( "ユーザーはどの言語を好みますか?" )
# 結果が返された理由を理解する
説明＝メモ。説明（質問）
# { 経路: "注意"、自信: ...、注意マップ: [...] }
非同期のサポート
運用非同期フレームワーク (FastAPI など) の場合:
engramma_memory からのインポート EngrammaMemoryAsync
EngrammaMemoryAsync ( dim = 256 、 backend = "cloud" 、 api_key = "..." ) を mem として使用する非同期:
私を待ってください。ストア (キー = 埋め込み、値 = データ)
結果 = 私を待ってください。クエリ (埋め込み、top_k = 5)
🤝 貢献しています
寄付を歓迎します!詳細については、CONTRIBUTING.md を参照してください。
git clone https://github.com/engramma-ai/engramma-memory.git
CD エングラマメモリー
pip install -e " .[dev] "
pytest # 39 テスト、すべて緑色
生産の準備はできていますか?
ローカルは永久に無料です。を打つと、

壁 — 1000 のパターン、持続性なし、因果推論なし — クラウドは 1 行先にあります。
MITライセンス •
ドキュメント •
ギットハブ •
問題点
AI エージェント用のコンポーザブル メモリ エンジン。検索のみを行うベクトル データベースとは異なり、Engramma は、正確な kNN、ホップフィールド エネルギー ネットワーク、およびマルチヘッド アテンションを組み合わせたハイブリッド アーキテクチャを活用して、ネイティブに構成、一般化、適応します。 1 つの pip インストール。 NumPy を超える依存関係はありません。 Engramma Cloud で本番環境に対応。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A composable memory engine for AI agents. Unlike vector databases that only retrieve, Engramma natively composes, generalizes, and adapts powered by a hybrid architecture combining exact kNN, Hopfield energy networks, and multi-head attention. One pip install. Zero dependencies beyond NumPy. Product
[truncated]

GitHub - engramma-ai/engramma-memory: A composable memory engine for AI agents. Unlike vector databases that only retrieve, Engramma natively composes, generalizes, and adapts powered by a hybrid architecture combining exact kNN, Hopfield energy networks, and multi-head attention. One pip install. Zero dependencies beyond NumPy. Production-ready with Engramma Cloud. · GitHub
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
engramma-ai
/
engramma-memory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows docs docs engramma_memory engramma_memory examples examples notebooks notebooks real_tests real_tests tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml View all files Repository files navigation
The memory engine for AI that thinks — not just retrieves.
Composition. Generalization. Causal reasoning. One pip install away.
Vector databases retrieve. Engramma composes .
Your agent asks: "What do you know about Python AND machine learning?"
ChromaDB returns two separate results. Engramma returns one fused answer .
Every AI memory system today is just retrieval — find the nearest vector, return it. That's a search engine, not a memory. Real memory does more:
🧩 Composes — "What's the intersection of X and Y?" → a single coherent answer
🧠 Generalizes — noisy input still triggers the right pattern
📈 Adapts — frequently accessed patterns become stronger
🗑️ Forgets — outdated patterns decay naturally
Engramma does all four. In 3 lines of code.
pip install engramma-memory
Basic Usage
import numpy as np
from engramma_memory import EngrammaMemory
# Initialize local, purely in-memory engine
mem = EngrammaMemory ( dim = 256 , backend = "local" )
# Store knowledge
mem . store ( key = embedding_a , value = "Python is a programming language" )
mem . store ( key = embedding_b , value = "Machine learning uses data to learn" )
# Retrieve — smart routing across 3 pathways
result = mem . retrieve ( query_embedding )
# Compose — the killer feature ⚡
blend = mem . compose ([ embedding_a , embedding_b ])
# Native multi-head attention fusion returns a coherent response
Note
That's it. No config files. No Docker. No API keys. Just numpy .
🆚 Why Not Just Use a Vector DB?
When combining concepts using traditional vector databases, you are forced to retrieve multiple results and manually piece them together. Engramma solves this with native composition .
# ChromaDB / Pinecone / FAISS — you get TWO separate results
result_a = db . query ( key_a ) # "Python is a language..."
result_b = db . query ( key_b ) # "ML uses data to learn..."
# Now what? Average them? Concatenate? Pray?
blend = ( result_a + result_b ) / 2 # Meaningless arithmetic
The Engramma Way
# Engramma — you get ONE fused answer
blend = mem . compose ([ key_a , key_b ])
# Each head specializes: some recall A, some recall B → coherent fusion
Feature
Traditional Vector DBs
Engramma
🔍 Nearest-neighbor search
✅
✅
🧩 Native composition
❌
✅
🧠 Soft generalization (Hopfield)
❌
✅
🔀 Adaptive routing
❌
✅
📉 Importance-based eviction
❌
✅
🍂 Gradual forgetting
❌
✅
📦 Zero dependencies
❌
✅ (numpy only)
🏗️ How It Works
Engramma uses a multi-pathway architecture to route queries intelligently based on the task.
graph TD
Q[Query] --> Exact[Exact kNN Memory]
Q --> Energy[Energy Hopfield]
Q --> MHA[Multi-Head Attention]
Exact --> Router[Confidence Router<br/>learned weights]
Energy --> Router
MHA --> Router
Router --> Best[Best Result]
Loading
Exact Memory — perfect recall via kNN with importance scoring
Energy Memory — soft generalization via temperature-scaled Hopfield dynamics
Multi-Head Attention — each head attends to different patterns → native composition
Confidence Router — learns which pathway handles which query type
All learning is local (Hebbian). No backpropagation. No GPU required. Pure NumPy.
Engramma trades a tiny bit of raw speed for massive gains in composition capability.
Engramma trades raw speed for composition capability. For pure nearest-neighbor at millions of vectors, use FAISS. For AI agents that need to think with their memory, use Engramma.
Engramma drops right into your existing AI stack.
from engramma_memory . integrations . langchain import EngrammaLangChainMemory
memory = EngrammaLangChainMemory ( dim = 256 , embed_fn = fn )
LlamaIndex
from engramma_memory . integrations . llamaindex import EngrammaRetriever
retriever = EngrammaRetriever ( dim = 256 , embed_fn = fn )
OpenAI Assistants
from engramma_memory . integrations . openai_assistants import engramma_tool_definitions
tools = engramma_tool_definitions ()
FastAPI
from engramma_memory . integrations . fastapi import create_memory_router
app . include_router ( create_memory_router ( dim = 256 ))
☁️ Engramma Cloud
Same API. No limits. 43 premium capabilities.
# Local (free, open source, limited to 1000 patterns)
mem = EngrammaMemory ( dim = 256 , backend = "local" )
# Cloud (unlimited, persistent, intelligent) — same code, one line change!
mem = EngrammaMemory ( dim = 256 , backend = "cloud" , api_key = "nx_live_..." )
Important
Feature
Local (free)
Cloud
🗃️ Max patterns
1,000
Unlimited
💾 Storage
RAM only
Tiered (hot/warm/cold)
⚖️ Composition weights
Equal only
Custom fractional (0.0–1.0)
🛡️ Persistence
None (in-process)
Durable + snapshots
🧠 Routing
Confidence-based
Active Inference + phi_B
🔗 Causal reasoning
—
DAG discovery + interventions
🚨 Anomaly detection
—
3-regime safety system
🔮 Temporal prediction
—
Granger causality + prefetch
💬 Text interface
—
HDC tokenizer (no embeddings needed)
🔬 Explainability
—
Full XAI dashboard
Cloud Feature Highlights
Causal Reasoning & Safety
# Discover causal structure
graph = mem . get_causal_graph ()
# "If I change A, what happens to B?"
effect = mem . predict_causal_effect ( cause_key = a , effect_key = b )
# Fractional composition (not just 50/50)
blend = mem . compose_fractional ( a , b , alpha = 0.7 )
# Auto-block risky OOD compositions
mem . enable_anomaly_protection ( enabled = True )
Text Memory & Explainability
# Store with natural language (no embeddings needed!)
mem . store_text ( "User prefers Python over JS" )
# Query with natural language
results = mem . query_text ( "what language does the user prefer?" )
# Understand WHY a result was returned
explanation = mem . explain ( query )
# { pathway: "attention", confidence: ..., attention_map: [...] }
Async Support
For production async frameworks (FastAPI, etc.):
from engramma_memory import EngrammaMemoryAsync
async with EngrammaMemoryAsync ( dim = 256 , backend = "cloud" , api_key = "..." ) as mem :
await mem . store ( key = embedding , value = data )
results = await mem . query ( embedding , top_k = 5 )
🤝 Contributing
We welcome contributions! See CONTRIBUTING.md for details.
git clone https://github.com/engramma-ai/engramma-memory.git
cd engramma-memory
pip install -e " .[dev] "
pytest # 39 tests, all green
Ready for production?
Local is free forever. When you hit the wall — 1000 patterns, no persistence, no causal reasoning — Cloud is one line away.
MIT License •
Documentation •
GitHub •
Issues
A composable memory engine for AI agents. Unlike vector databases that only retrieve, Engramma natively composes, generalizes, and adapts powered by a hybrid architecture combining exact kNN, Hopfield energy networks, and multi-head attention. One pip install. Zero dependencies beyond NumPy. Production-ready with Engramma Cloud.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
