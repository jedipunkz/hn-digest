---
source: "https://github.com/HtooTayZa/sawtooth-memory"
hn_url: "https://news.ycombinator.com/item?id=48422850"
title: "Sawtooth – An async, multi-tiered memory framework for LLM agents"
article_title: "GitHub - HtooTayZa/sawtooth-memory: Async hierarchical memory middleware for LLM agents. · GitHub"
author: "Jackmann_01"
captured_at: "2026-06-06T09:45:12Z"
capture_tool: "hn-digest"
hn_id: 48422850
score: 1
comments: 0
posted_at: "2026-06-06T08:53:03Z"
tags:
  - hacker-news
  - translated
---

# Sawtooth – An async, multi-tiered memory framework for LLM agents

- HN: [48422850](https://news.ycombinator.com/item?id=48422850)
- Source: [github.com](https://github.com/HtooTayZa/sawtooth-memory)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T08:53:03Z

## Translation

タイトル: Sawtooth – LLM エージェント用の非同期多層メモリ フレームワーク
記事のタイトル: GitHub - HtooTayZa/sawtooth-memory: LLM エージェント用の非同期階層メモリ ミドルウェア。 · GitHub
説明: LLM エージェント用の非同期階層メモリ ミドルウェア。 - HtooTayZa/ノコギリメモリ

記事本文:
GitHub - HtooTayZa/sawtooth-memory: LLM エージェント用の非同期階層メモリ ミドルウェア。 · GitHub
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
トゥーテイザ
/
ノコギリ歯状メモリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン

ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
74 コミット 74 コミット .github .github .idea .idea 例 例 sawtooth_memory sawtooth_memory スクリプト スクリプト テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml BENCHMARKS.md BENCHMARKS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DOCUMENTATION.md DOCUMENTATION.md LICENSE.md LICENSE.md README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM エージェント用の高性能でノンブロッキングの階層メモリ フレームワーク。
標準の LLM メモリ システム (LangChain の ConversationsummaryMemory など) は、メイン アプリケーション スレッドで会話履歴を順次処理します。ユーザーがメッセージを送信するたびに、システムが LLM による新しい履歴概要の生成を待機している間、アプリケーション全体がフリーズします。さらに、これらのサマリーは「途中で失われた」幻覚効果に悩まされ、トークンを保存するために特定の UUID、名前、またはルールが頻繁に削除されます。
Sawtooth メモリは、この遅延とデータ損失を排除します。ユーザーのメッセージを即座に保存し、ミリ秒以内に制御をアプリケーションに戻し、大量の要約を非同期のバックグラウンド ワーカーにオフロードします。幻覚を防ぐために、要約する前に重要な事実を不変の台帳に抽出します。
アーキテクチャの詳細、包括的な API 仕様、高度なライフサイクル構成については、公式ドキュメントを参照してください。
詳細なアーキテクチャと API リファレンスを表示 (DOCUMENTATION.md)
1. ノンブロッキング実行モデル
標準メモリ (ブロッキング) ノコギリ波メモリ (非同期)
────────────────

────────
[ アプリケーション ] [ アプリケーション ]
│ │
▼ ▼
[ コンテキストの保存 ] [ ContextManager ]
│ │
▼ --------------------------------------------------------┐ (即時返却)
[ LLM の要約 ] ▼ ▼
(アプリは 5 ～ 10 秒間フリーズします) [ 次のユーザーターン ] [ バックグラウンドワーカー ]
│ │
▼ ▼
[次のユーザーターン] [LLM の要約]
2. 階層型メモリスタック
エージェントが応答する準備が整うと、Sawtooth は個別のレイヤーから最適化されたコンテキスト ペイロードをつなぎ合わせて、重要な事実が要約されないようにします。
エージェントループ
│
▼
┌─────────────┐
│ コンテキストマネージャー │
│ ┌───────────┐ │
│ │ L0 システム │ │ 不変のペルソナ + ツールのスキーマ
│ │ L2 アーカイブ │ │ 圧縮された物語記憶
│ │ L1.5 エンティティ │ │ 正確な ID、パス、UUID
│ │ L1 仕事中 │ │ 最近の生の会話
│ ━━━━━━━━┘ │
━─────┬─────┘
│
▼
build_prompt()
│
▼
LLM API
パフォーマンスのベンチマーク
圧縮をバックグラウンドに移動することにより、Sawtooth は 100% のリコール精度を維持しながら、メインスレッドでのレイテンシの大幅な削減を実現します。
ローカル GPU ベンチマーク (NVIDIA RTX 5060 | モデル: phi4-mini | 20 メッセージ会話)
完全な方法論、クラウドの比較、再現性の手順については、「パフォーマンス ベンチマークを読む」をご覧ください。
pip install sawtooth-memory
クラウドプロバイダーのオプションの依存関係:
pip インストール langchain-openai langchain-anthropic langchain-google-genai
クイックスタート
ContextManager を初期化し、バックグラウンド ワーカーに面倒な作業を処理させます。ノコギリ歯は、ローカル エアギャップ モデル (Ollama) および CL と普遍的に互換性があります。

API を使用します。
非同期をインポートする
sawtooth_memory からインポート ContextManager 、 ContextManagerConfig
sawtooth_memory から。構成インポート OllamaConfig
非同期デフォルトメイン():
config = ContextManagerConfig (
ソフトリミットトークン = 1000 、
ハードリミットトークン = 2000 、
ollama = OllamaConfig (base_url = "http://localhost:11434" 、モデル = "phi4" )
)
cm として ContextManager ( system_prompt = "あなたは役に立つアシスタントです。" 、 config = config ) と非同期:
# 1. メッセージを即座に取り込む (メインスレッドがブロックされることはありません)
cmを待ちます。 add_message ( "user" , "私のトランザクション ID は txn_998877_alpha" )
cmを待ちます。 add_message ( "assistant" , "トランザクション ID をメモしました。" )
# 2. メイン LLM に送信する最適化されたプロンプトを構築する
プロンプト = cm 。ビルドプロンプト()
印刷 (プロンプト)
if __name__ == "__main__" :
非同期 。実行 (メイン ())
2. 説明可能性トレースを思い出す
Sawtooth は、確定的な監査証跡を提供することで、エージェント メモリの「ブラック ボックス」を排除します。メモリ システムにクエリを実行すると、プロンプトにファクトが保持された理由を正確に確認できます。
トレース = cm 。説明プロンプト ()
jsonをインポートする
print ( json . dumps ( トレース , インデント = 2 ))
出力:
{
"system_prompt" : " あなたは親切なアシスタントです。 " ,
"l2_summary_lineage" : [
" ユーザーがルーターのトラブルシューティングを開始しました。 " 、
「ユーザーが指定した MAC アドレス。」
]、
"l1_5_entities" : [
{
"key" : " user_transaction_id " 、
"値" : " txn_998877_alpha " ,
"origin" : " L1.5 明示的命令を介してアンカー "
}
]、
"l1_active_messages" : 4 、
"total_tokens" : 342
}
3. 統合: LangGraph
Sawtooth はネイティブの SawtoothMemorySaver アダプターを提供し、LangGraph アーキテクチャのドロップイン チェックポイント作成機能の代替として機能します。
ランググラフから。グラフインポートStateGraph
sawtooth_memory から。統合。 langgraph インポート SawtoothMemorySaver
graph_builder = StateGraph (状態)
# ... ノードを追加し、

エッジ ...
メモリセーバー = SawtoothMemorySaver (cm)
グラフ = グラフ_ビルダー 。コンパイル (チェックポイント = メモリセーバー)
ロードマップ
非同期バックグラウンドワーカー
ローカル (Ollama) とクラウドの互換性
フェーズ 2: 可観測性とテレメトリ
永続的な JSONL 監査ジャーナル
パフォーマンスベンチマークハーネス
フェーズ 3: 高度なアーキテクチャ (次へ)
マルチエージェント メモリ プーリング (共有コンテキスト状態)
セマンティック ベクター L3 アーカイブ メモリ (RAG 統合)
分散デプロイメント用の Redis/Postgres アダプター
プルリクエストも歓迎します。テスト スイートを実行してコードの品質を確保する方法に関するガイドラインについては、CONTRIBUTING.md を参照してください。
このプロジェクトは MIT ライセンスに基づいてライセンスされています。詳細については、LICENSE.md ファイルを参照してください。
LLM エージェント用の非同期階層メモリ ミドルウェア。
pypi.org/project/sawtooth-memory/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
0.2.0: 非同期階層メモリと説明可能性トレース
最新の
2026 年 6 月 6 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Async hierarchical memory middleware for LLM agents. - HtooTayZa/sawtooth-memory

GitHub - HtooTayZa/sawtooth-memory: Async hierarchical memory middleware for LLM agents. · GitHub
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
HtooTayZa
/
sawtooth-memory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
74 Commits 74 Commits .github .github .idea .idea examples examples sawtooth_memory sawtooth_memory scripts scripts tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml BENCHMARKS.md BENCHMARKS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DOCUMENTATION.md DOCUMENTATION.md LICENSE.md LICENSE.md README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
A high-performance, non-blocking hierarchical memory framework for LLM Agents.
Standard LLM memory systems (like LangChain's ConversationSummaryMemory ) process conversation history sequentially on the main application thread. Every time a user sends a message, the entire application freezes while the system waits for an LLM to generate a new historical summary. Furthermore, these summaries suffer from the "Lost in the Middle" hallucination effect, frequently deleting specific UUIDs, names, or rules to save tokens.
Sawtooth Memory eliminates this latency and data loss. It immediately stores the user's message and returns control to the application in milliseconds, offloading the heavy summarization to an asynchronous background worker. To prevent hallucinations, it extracts critical facts into an immutable ledger before summarizing.
For deep architectural deep-dives, comprehensive API specifications, and advanced lifecycle configurations, please refer to the official documentation:
View Detailed Architecture & API Reference (DOCUMENTATION.md)
1. The Non-Blocking Execution Model
Standard Memory (Blocking) Sawtooth Memory (Async)
────────────────────────── ───────────────────────
[ Application ] [ Application ]
│ │
▼ ▼
[ Save Context ] [ ContextManager ]
│ │
▼ ├───────────────────┐ (Instant Return)
[ LLM Summarizes ] ▼ ▼
(App freezes for 5-10s) [ Next User Turn ] [ Background Worker ]
│ │
▼ ▼
[ Next User Turn ] [ LLM Summarizes ]
2. The Hierarchical Memory Stack
When your agent is ready to respond, Sawtooth stitches together an optimized context payload from distinct layers, ensuring critical facts are never summarized away.
Agent Loop
│
▼
┌─────────────────────┐
│ ContextManager │
│ ┌───────────────┐ │
│ │ L0 System │ │ immutable persona + tool schemas
│ │ L2 Archive │ │ compressed narrative memory
│ │ L1.5 Entities │ │ exact IDs, paths, UUIDs
│ │ L1 Working │ │ recent raw conversation
│ └───────────────┘ │
└──────────┬──────────┘
│
▼
build_prompt()
│
▼
LLM API
Performance Benchmarks
By moving compression to the background, Sawtooth achieves massive latency reductions on the main thread while maintaining 100% recall accuracy.
Local GPU Benchmark (NVIDIA RTX 5060 | Model: phi4-mini | 20-Message Conversation)
For full methodology, cloud comparisons, and reproducibility steps, view our Read the Performance Benchmarks .
pip install sawtooth-memory
Optional dependencies for cloud providers:
pip install langchain-openai langchain-anthropic langchain-google-genai
Quickstart
Initialize the ContextManager and let the background worker handle the heavy lifting. Sawtooth is universally compatible with local air-gapped models (Ollama) and cloud APIs.
import asyncio
from sawtooth_memory import ContextManager , ContextManagerConfig
from sawtooth_memory . config import OllamaConfig
async def main ():
config = ContextManagerConfig (
soft_limit_tokens = 1000 ,
hard_limit_tokens = 2000 ,
ollama = OllamaConfig ( base_url = "http://localhost:11434" , model = "phi4" )
)
async with ContextManager ( system_prompt = "You are a helpful assistant." , config = config ) as cm :
# 1. Instantly ingest messages (Main thread is never blocked)
await cm . add_message ( "user" , "My transaction ID is txn_998877_alpha" )
await cm . add_message ( "assistant" , "I have noted your transaction ID." )
# 2. Build the optimized prompt to send to your main LLM
prompt = cm . build_prompt ()
print ( prompt )
if __name__ == "__main__" :
asyncio . run ( main ())
2. Recall Explainability Traces
Sawtooth eliminates the "black-box" of agent memory by providing deterministic audit trails. You can query the memory system to see exactly why a fact was retained in the prompt.
trace = cm . explain_prompt ()
import json
print ( json . dumps ( trace , indent = 2 ))
Output:
{
"system_prompt" : " You are a helpful assistant. " ,
"l2_summary_lineage" : [
" User initiated troubleshooting for router. " ,
" User provided MAC address. "
],
"l1_5_entities" : [
{
"key" : " user_transaction_id " ,
"value" : " txn_998877_alpha " ,
"origin" : " Anchored via L1.5 explicit instruction "
}
],
"l1_active_messages" : 4 ,
"total_tokens" : 342
}
3. Integrations: LangGraph
Sawtooth provides a native SawtoothMemorySaver adapter, acting as a drop-in checkpointer replacement for LangGraph architectures.
from langgraph . graph import StateGraph
from sawtooth_memory . integrations . langgraph import SawtoothMemorySaver
graph_builder = StateGraph ( State )
# ... add nodes and edges ...
memory_saver = SawtoothMemorySaver ( cm )
graph = graph_builder . compile ( checkpointer = memory_saver )
Roadmap
Asynchronous Background Worker
Local (Ollama) & Cloud compatibility
Phase 2: Observability & Telemetry
Persistent JSONL Auditing Journal
Performance Benchmarking Harness
Phase 3: Advanced Architectures (Up Next)
Multi-Agent Memory Pooling (Shared contextual state)
Semantic Vector L3 Archival Memory (RAG integration)
Redis/Postgres Adapter for Distributed Deployments
We welcome pull requests. See our CONTRIBUTING.md for guidelines on how to run the test suite and ensure code quality.
This project is licensed under the MIT License - see the LICENSE.md file for details.
Async hierarchical memory middleware for LLM agents.
pypi.org/project/sawtooth-memory/
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
0.2.0: Asynchronous Hierarchical Memory & Explainability Traces
Latest
Jun 6, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
