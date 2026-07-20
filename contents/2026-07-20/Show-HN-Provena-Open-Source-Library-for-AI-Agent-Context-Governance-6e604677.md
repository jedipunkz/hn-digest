---
source: "https://github.com/rajfirke/provena"
hn_url: "https://news.ycombinator.com/item?id=48984491"
title: "Show HN: Provena: Open-Source Library for AI Agent Context Governance"
article_title: "GitHub - rajfirke/provena: Context governance for agentic AI — tamper-evident audit trails, provenance validation, EU AI Act compliance. 6 framework adapters, MCP server, PostgreSQL, policy engine. · GitHub"
author: "rfirke"
captured_at: "2026-07-20T21:01:40Z"
capture_tool: "hn-digest"
hn_id: 48984491
score: 1
comments: 0
posted_at: "2026-07-20T20:31:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Provena: Open-Source Library for AI Agent Context Governance

- HN: [48984491](https://news.ycombinator.com/item?id=48984491)
- Source: [github.com](https://github.com/rajfirke/provena)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T20:31:34Z

## Translation

タイトル: HN を表示: Provena: AI エージェント コンテキスト ガバナンスのオープンソース ライブラリ
記事のタイトル: GitHub - rajfirke/provena: エージェント AI のコンテキスト ガバナンス — 改ざん防止監査証跡、出所検証、EU AI 法の準拠。 6 つのフレームワーク アダプター、MCP サーバー、PostgreSQL、ポリシー エンジン。 · GitHub
説明: エージェント AI のコンテキスト ガバナンス - 改ざん明示的な監査証跡、来歴検証、EU AI 法のコンプライアンス。 6 つのフレームワーク アダプター、MCP サーバー、PostgreSQL、ポリシー エンジン。 - ラージフィルケ/プロヴェナ
HN本文：投稿者募集中！ 7 人の寄稿者がおり、11 件の「良好な初号」ラベルと 17 件の「ヘルプ募集」ラベルがあります。コード例がすべて網羅されています。プロジェクト:
コンテキスト入力層を管理する人は誰もいません。エージェントの行動 (Microsoft AGT)、エージェントの発言 (Guardrails AI)、およびエージェントのコミュニケーション方法 (NeMo) を管理するツールが存在します。しかし、コンテキスト ウィンドウに流入するデータ (取得結果、ツール出力、エージェント メッセージ) は完全に管理されていません。それはProvenaが修正するものです

記事本文:
GitHub - rajfirke/provena: エージェント AI のコンテキスト ガバナンス — 改ざん防止監査証跡、来歴検証、EU AI 法のコンプライアンス。 6 つのフレームワーク アダプター、MCP サーバー、PostgreSQL、ポリシー エンジン。 · GitHub
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
{{混乱

年齢 }}
ラージフィルケ
/
プロヴェナ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
43 コミット 43 コミット .github .github docs docs サンプル サンプル スクリプト スクリプト src/rovena src/provena テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE PROMOTION.md PROMOTION.md README.md README.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml givena.example.toml givena.example.toml givena.example.yaml givena.example.yaml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Context governance for agentic AI systems. (MVP)
AI エージェントは、6 つの異なるソースからのデータに基づいて決定を下しました。
Can you tell me which ones? Can you prove the data wasn't tampered with?
Can you verify it was still current?
Provena は、Python の 3 行で改ざん明示的な監査証跡を AI エージェントのコンテキスト パイプラインに追加します。
from provena import ContextTrail
トレイル = ContextTrail ()
@トレイル。 track ( source = "retriever" )
def 検索 (クエリ):
リターンレトリーバー。検索（クエリ）
search() へのすべての呼び出しは、SHA-256 コンテンツ ハッシュ、出所検証、
そして改ざんを検出するハッシュチェーン監査証跡。
AGT governs what agents DO. Guardrails AI governs what agents SAY. Provena governs what agents KNOW.
コンテキスト入力レイヤーを管理する既存のツールはありません。 Provena fills this gap with:
改ざん防止監査証跡 - オプションの HMAC 署名を使用した SHA-256 ハッシュ チェーン (マークル スタイル) ログ記録
来歴の検証 - コンテキストに適切なソース メタデータ (有効/欠落/不完全) が含まれていることを検証します。
鮮度チェック — メタデータを介して古いコンテキストを検出します。

imestamp と正規表現の時間的検出 (FRESH / STALE / UNKNOWN)
任意のコンテキスト ソース - RAG リトリーバー、ツール出力、エージェント メッセージ、メモリ リコール、MCP リソース
1ms 未満のオーバーヘッド — 純粋な Python、ML モデル、ONNX、モデルのダウンロードなし
依存関係なし - コア ライブラリは Python 標準ライブラリのみを使用します
pip install provena # core (依存関係なし)
pip install provena[cli] # + CLI ツール (クリック、リッチ)
pip install provena[otel] # + OpenTelemetry エクスポート
pip install provena[langchain] # + LangChain アダプター
pip install provena[llamaindex] # + LlamaIndex アダプター
pip install provena[all] # すべて
クイックスタート
from provena import ContextTrail 、 ProvenanceMetadata
from datetime import datetime 、タイムゾーン
トレイル = ContextTrail ( storage_path = "audit.db" )
# コンテキストを生成する関数を追跡する
@トレイル。トラック (ソース = "レトリバー")
def 検索 (クエリ):
リターンレトリーバー。検索（クエリ）
@トレイル。トラック (ソース = "tool:pricing_api" )
def get_price (product_id):
API を返します。 get ( f"/price/ { product_id } " )
# 来歴メタデータを使用した手動ロギング
トレイル。ログ (
content = "エンタープライズ プランの料金は月額 499 ドルです。" 、
ソース = "ツール:価格設定_api" ,
来歴 = ProvenanceMetadata (
source_url = "https://api.example.com/pricing" ,
created_at = 日時 。現在 (タイムゾーン .utc)、
）、
)
# 監査証跡が改ざんされていないことを確認する
評決 = 証跡 。 verify_chain ()
print ( f"チェーンは無傷です: { 判定 . 無傷 } " )
print ( f"合計レコード: { verdict . total_records } " )
CLI
pip install provena[cli] を使用してインストールし、次のようにします。
# ハッシュチェーンの整合性を検証する
provena --db Audit.db verify
# PASS — チェーンは無傷 (42 レコードが検証済み)
# 監査ログをクエリする
provena --db Audit.db Audit --sourceretriever --format json
# ガバナンスレポートを生成する
プロヴェナ --db Audit.db レポート -

-テキストのフォーマット
# 簡単な概要
provena --db Audit.db の概要
HMAC 署名付き証跡の場合は、 --signing-key を渡すか、 PROVENA_SIGNING_KEY を設定します。
プロヴェナから。統合。 langchain import ProvenaCallback
チェーン = RetrievalQA 。 from_chain_type (
llm = llm 、
レトリーバー = レトリーバー、
コールバック = [ ProvenaCallback ( トレイル = トレイル )]、
)
ラマインデックス
プロヴェナから。統合。 llamaindex import Provenaポストプロセッサ
クエリエンジン = インデックス 。 as_query_engine (
node_postprocessors = [ProvenaPostprocessor (トレイル = トレイル)]
)
オープンテレメトリ
トレイル = コンテキストトレイル (
storage_path = "audit.db" ,
otel_enabled = True 、
otel_service_name = "私のエージェント" ,
)
# すべての log() 呼び出しは、ガバナンス属性を持つ OTel スパンを出力するようになりました
建築
あなたのアプリケーション
│
│ レトリーバー ──┐
│ ツールコール ──┤
│ エージェントメッセージ ──┼──► ContextTrail ──► LLM コンテキストウィンドウ
│ 記憶 ──┤ │
│ MCP ──┘ │
│ ┌─────┴─────────┐
│ │ ProvenanceValidator │
│ │ フレッシュネスチェッカー │
│ │ ハッシュチェーン (SHA-256) │
│ │ SQLite バックエンド │
│ │ OTel エクスポーター │
│ ━━━━━━━━━━━━┘
コンプライアンス
Provena は、EU AI 法の要件 (施行: 2026 年 8 月 2 日) に直接対応しています。
OWASP ASI06 (メモリおよびコンテキスト ポイズニング) にも対処します。
寄付を歓迎します！開発セットアップ、アーキテクチャ ガイド、PR プロセスについては、CONTRIBUTING.md を参照してください。
参加する前に当社の行動規範をお読みください。
エージェント AI のコンテキスト ガバナンス - 改ざん明示的な監査証跡、来歴検証、EU AI 法のコンプライアンス。 6 つのフレームワーク アダプター、MCP サーバー、PostgreSQL、ポリシー エンジン。
読み込み中にエラーが発生しました。このページをリロードしてください。
4
フォーク
レポートリポジトリ
リリース
7
v0.15.0
最新の
2026 年 7 月 20 日
+6

リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Context governance for agentic AI — tamper-evident audit trails, provenance validation, EU AI Act compliance. 6 framework adapters, MCP server, PostgreSQL, policy engine. - rajfirke/provena

Looking for contributors! We have 7 contributors and We have 11 "good first issue" labels and 17 "help wanted" issues. All well-scoped with code examples. Project:
Nobody governs the context input layer. Tools exist to govern what agents DO (Microsoft AGT), what agents SAY (Guardrails AI), and how agents COMMUNICATE (NeMo). But the data flowing INTO the context window — retriever results, tool outputs, agent messages — is completely ungoverned. That's what Provena fixes

GitHub - rajfirke/provena: Context governance for agentic AI — tamper-evident audit trails, provenance validation, EU AI Act compliance. 6 framework adapters, MCP server, PostgreSQL, policy engine. · GitHub
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
rajfirke
/
provena
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
43 Commits 43 Commits .github .github docs docs examples examples scripts scripts src/ provena src/ provena tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PROMOTION.md PROMOTION.md README.md README.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml provena.example.toml provena.example.toml provena.example.yaml provena.example.yaml pyproject.toml pyproject.toml View all files Repository files navigation
Context governance for agentic AI systems. (MVP)
Your AI agent just made a decision based on data from 6 different sources.
Can you tell me which ones? Can you prove the data wasn't tampered with?
Can you verify it was still current?
Provena adds tamper-evident audit trails to any AI agent's context pipeline — in 3 lines of Python.
from provena import ContextTrail
trail = ContextTrail ()
@ trail . track ( source = "retriever" )
def search ( query ):
return retriever . search ( query )
Every call to search() is now logged with a SHA-256 content hash, provenance validation,
and a hash-chained audit trail that detects tampering.
AGT governs what agents DO. Guardrails AI governs what agents SAY. Provena governs what agents KNOW.
No existing tool governs the context input layer. Provena fills this gap with:
Tamper-evident audit trails — SHA-256 hash-chained (Merkle-style) logging with optional HMAC signing
Provenance validation — Verify that context carries proper source metadata (VALID / MISSING / INCOMPLETE)
Freshness checking — Detect stale context via metadata timestamps and regex temporal detection (FRESH / STALE / UNKNOWN)
Any context source — RAG retrievers, tool outputs, agent messages, memory recalls, MCP resources
Sub-1ms overhead — Pure Python, no ML models, no ONNX, no model downloads
Zero dependencies — Core library uses only the Python standard library
pip install provena # core (zero dependencies)
pip install provena[cli] # + CLI tools (click, rich)
pip install provena[otel] # + OpenTelemetry export
pip install provena[langchain] # + LangChain adapter
pip install provena[llamaindex] # + LlamaIndex adapter
pip install provena[all] # everything
Quick Start
from provena import ContextTrail , ProvenanceMetadata
from datetime import datetime , timezone
trail = ContextTrail ( storage_path = "audit.db" )
# Track any function that produces context
@ trail . track ( source = "retriever" )
def search ( query ):
return retriever . search ( query )
@ trail . track ( source = "tool:pricing_api" )
def get_price ( product_id ):
return api . get ( f"/price/ { product_id } " )
# Manual logging with provenance metadata
trail . log (
content = "The enterprise plan costs $499/month." ,
source = "tool:pricing_api" ,
provenance = ProvenanceMetadata (
source_url = "https://api.example.com/pricing" ,
created_at = datetime . now ( timezone . utc ),
),
)
# Verify the audit trail hasn't been tampered with
verdict = trail . verify_chain ()
print ( f"Chain intact: { verdict . intact } " )
print ( f"Total records: { verdict . total_records } " )
CLI
Install with pip install provena[cli] , then:
# Verify hash chain integrity
provena --db audit.db verify
# PASS — Chain intact (42 records verified)
# Query the audit log
provena --db audit.db audit --source retriever --format json
# Generate a governance report
provena --db audit.db report --format text
# Quick summary
provena --db audit.db summary
For HMAC-signed trails, pass --signing-key or set PROVENA_SIGNING_KEY .
from provena . integrations . langchain import ProvenaCallback
chain = RetrievalQA . from_chain_type (
llm = llm ,
retriever = retriever ,
callbacks = [ ProvenaCallback ( trail = trail )],
)
LlamaIndex
from provena . integrations . llamaindex import ProvenaPostprocessor
query_engine = index . as_query_engine (
node_postprocessors = [ ProvenaPostprocessor ( trail = trail )]
)
OpenTelemetry
trail = ContextTrail (
storage_path = "audit.db" ,
otel_enabled = True ,
otel_service_name = "my-agent" ,
)
# Every log() call now emits an OTel span with governance attributes
Architecture
Your Application
│
│ Retriever ──┐
│ Tool Call ──┤
│ Agent Msg ──┼──► ContextTrail ──► LLM Context Window
│ Memory ──┤ │
│ MCP ──┘ │
│ ┌─────┴──────────────┐
│ │ ProvenanceValidator │
│ │ FreshnessChecker │
│ │ HashChain (SHA-256) │
│ │ SQLite Backend │
│ │ OTel Exporter │
│ └────────────────────┘
Compliance
Provena maps directly to EU AI Act requirements (enforcement: August 2, 2026):
Also addresses OWASP ASI06 (Memory & Context Poisoning).
We welcome contributions! See CONTRIBUTING.md for development setup, architecture guide, and PR process.
Please read our Code of Conduct before participating.
Context governance for agentic AI — tamper-evident audit trails, provenance validation, EU AI Act compliance. 6 framework adapters, MCP server, PostgreSQL, policy engine.
There was an error while loading. Please reload this page .
4
forks
Report repository
Releases
7
v0.15.0
Latest
Jul 20, 2026
+ 6 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
