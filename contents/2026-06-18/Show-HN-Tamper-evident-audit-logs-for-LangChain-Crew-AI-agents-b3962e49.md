---
source: "https://github.com/Providex-AI/rootsign"
hn_url: "https://news.ycombinator.com/item?id=48588177"
title: "Show HN: Tamper-evident audit logs for LangChain/Crew AI agents"
article_title: "GitHub - Providex-AI/rootsign: Rootsign is an open-source tamper-evident decision and action provenance logging library for AI agents · GitHub"
author: "oabolade"
captured_at: "2026-06-18T18:56:32Z"
capture_tool: "hn-digest"
hn_id: 48588177
score: 2
comments: 0
posted_at: "2026-06-18T16:52:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tamper-evident audit logs for LangChain/Crew AI agents

- HN: [48588177](https://news.ycombinator.com/item?id=48588177)
- Source: [github.com](https://github.com/Providex-AI/rootsign)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T16:52:51Z

## Translation

タイトル: HN を表示: LangChain/Crew AI エージェントの改ざん防止監査ログ
記事のタイトル: GitHub - Providex-AI/rootsign: Rootsign は、AI エージェント用のオープンソースの改ざん明示決定およびアクション来歴ログ ライブラリです · GitHub
説明: Rootsign は、AI エージェント用のオープンソースの改ざん明示決定およびアクションの出所ログ ライブラリです - Providex-AI/rootsign
HN テキスト: 最近、LangChain および CrewAI エージェント パイプラインを計測する際に同じ問題に遭遇しています。ツールの呼び出しで何か問題が発生した場合、エージェントが何をどの順序で実行したか、ログが変更されたかどうかを証明する方法はありませんでした。 LangSmith や Langfuse などの可観測性プラットフォームは、エージェントの動作、トークン、コストの最適化には優れていますが、法的に防御可能で監査可能なアーティファクトは生成しません。これが、CrewAI および LangGraph エージェントを計測し、CrewAI および LangGraph エージェントの暗号監査ログを生成する SDK である RootSign を構築することにした理由です。 RootSign は、すべてのツール呼び出しに暗号化ハッシュ チェーンを追加します。レコードが事後的に変更された場合、「rootsign verify」がそれを検出します。機能: - セッション内のすべてのアクション レコードにわたる SHA-256 ハッシュ チェーン
- 特定のエージェントのアクションに対する承認レコードを備えた人間参加型チェックポイント
- ハッシュ前に編集された PII (標準の StandardPIIConfig)
- LangGraph および CrewAI と連携 — AutoGen は近日公開予定
- ローカルファースト (Postgres + Timescale) — クラウド依存なし (まだ) できないこと: コンプライアンス ダッシュボード、クラウド バックエンド、ポリシー エンジン、すべてがロードマップにあります。 Github リポジトリで試してみてください。貢献やフィードバックはいつでも歓迎されます。

記事本文:
GitHub - Providex-AI/rootsign: Rootsign は、AI エージェント用のオープンソースの改ざん明示的決定およびアクション来歴ログ ライブラリです · GitHub
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
Providex-AI
/
ルートサイン
公共
通知
通知設定を変更するにはサインインする必要があります

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
32 コミット 32 コミット .github .github .vscode .vscode docs docs サンプル サンプル ルートサイン ルートサイン スクリプト スクリプト テスト テスト .env.example .env.example .gitignore .gitignore CLA.md CLA.md COTRIBUTING.md COTRIBUTING.md CONTRIBUTORS.md CONTRIBUTORS.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md alembic.ini alembic.ini docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
実稼働 AI エージェントの改ざん防止出所ログ。
RootSign は Providex AI 製品であり、Providex AI Agent Accountability Platform のエージェント キャプチャ レイヤーです。
AI エージェントが運用環境でアクション (ツールの呼び出し、API の実行、データベースへの書き込み) を実行する場合、組み込みの監査証跡はありません。何か問題が発生した場合（誤った払い戻し、PII 記録の漏洩、不正な展開）、エージェントが何をどの順序で誰の許可に基づいて実行したか、または記録が事後に改ざんされたかどうかを証明する方法はありません。
RootSign はこれを解決します。各エージェントのアクションは、前のアクションの SHA-256 ハッシュ (レコードの改ざんを証明する暗号化ハッシュ チェーン) を含むアクション レコードとしてキャプチャされます。事後にレコードを変更すると、rootsign verify がそれを検出します。
コンプライアンスグレードの監査証跡。エージェント コードに一切の変更を加えません。
フェーズ 1 MVP — v0.1.1。 LangGraph + CrewAI の統合、rootsign verify CLI、PII 編集、人間参加型チェックポイント、およびオプトイン意思決定キャプチャ (PRD-19 / ADR-008) はすべて出荷されています。
Python 3.11 または 3.12 を推奨します。 RootSign 自体は 3.11 以降をサポートしていますが、[crewai] エクストラは現在 3.13/3.14 ホイールで遅れています。 「一致する分布がありません」を選択した場合

crewai の場合は、Python 3.12 に切り替えて再インストールします。
pip install rootsign[langgraph]
PostgreSQL + TimescaleDB をローカルで起動し、スキーマを適用します。
rootsign-admin start-db # docker run timescale/timescaledb:latest-pg16
rootsign-admin init # アレンビックアップグレードヘッド
start-db は単一の Docker 実行をラップするため、リポジトリのクローンを作成する必要はありません。クローンを作成した場合、 docker-compose up -d db が同等の開発者パスになります。どちらも同じ rootsign-timescaledb コンテナー名と rootsign_pgdata ボリュームを再利用します。両方ではなく、どちらかを選択してください。
2. エージェントを登録します (1 回限りのセットアップ)
非同期をインポートする
rootsign import register_agent 、 AgentEnvironment 、 AgentRiskTier 、 AgentFramework から
エージェント = asyncio 。 run ( register_agent (
名前 = "私の請求書エージェント" ,
所有者 = "プラットフォーム チーム" ,
環境 = エージェント環境 。生産、
リスク層 = エージェントリスク層 。高い、
フレームワーク = AgentFramework 。ランググラフ、
))
print ( エージェント . エージェント ID )
3. ツールを計測する
ルートサインをインポートする
ルートサインから LocalIngestClient をインポート
ルートサインから。データベースのインポート AsyncSessionLocal
langchain_core から。ツールインポートツール
ランググラフから。事前構築済みインポートツールノード
@ツール
def send_invoice (customer_id : str , amount : float ) -> str :
"""顧客に請求書を送信します。"""
「送信済み」を返す
async def run_graph (agent_id):
db として AsyncSessionLocal () を使用した非同期:
client = LocalIngestClient ( db = db )
rootsign と非同期。セッション (agent_id = Agent_id 、 client = client ) として ctx :
tools = ルートサイン 。 Wrap_tools ([ send_invoice ]、ctx = ctx、client = client)
tool_node = ツールノード (ツール)
# ...通常どおりグラフを構築して実行します
データベースを待ちます。コミット()
すべてのツール呼び出しにより、改ざんが明らかなアクション レコードがハッシュ チェーン上に生成されるようになりました。
$ rootsign verify 660e8400-e29b-41d4-a716-446655440001
有効 ✓ — 3 レコード、チェーンは無傷
セッション: 660e84

00-e29b-41d4-a716-446655440001
終了コードは、有効の場合は 0、改ざんの場合は 1 です。オフライン JSONL セッション ファイルには --local <path.jsonl> を使用します (DB は必要ありません)。
レコードが変更された場合、検証者は壊れたリンクに次の名前を付けます。
$ rootsign verify 660e8400-e29b-41d4-a716-446655440001
改ざん ✗ — レコード #2 でチェーンが壊れました
詳細: アクション < action_id > の self_hash の不一致
セッション: 660e8400-e29b-41d4-a716-446655440001
警告: このセッション ログは改ざんされている可能性があります。
バージョン マトリックスと統合メモについては、docs/framework-support.md を参照してください。完全な実行可能な LangGraph サンプル (ReAct エージェント、3 つのインストルメント化ツール、OpenAI 支援) は、examples/langgraph-invoice-agent にあります。
各ツールを呼び出す前に理由を記録します。これはフェーズ 2 セッションの再生の基礎です。デフォルトではオフです。 ROOTSIGN_CAPTURE_DECISIONS=true を使用して意図的にオプトインします。
OSをインポートする
OS 。環境 [ "ROOTSIGN_CAPTURE_DECISIONS" ] = "true"
rootsign と非同期。セッション (agent_id = エージェント . エージェント ID 、クライアント = クライアント) として ctx :
# ツールを呼び出す前にエージェントが決定した内容を記録します。
ctx を待ちます。記録_決定 (
selected_action = "send_invoice" ,
reasoning_summary = "保険契約内の金額。受信者は確認済みです。" 、
信頼度 = 0.97 、
ingest_client = クライアント 、
)
tools = ルートサイン 。 Wrap_tools ([ send_invoice ]、ctx = ctx、client = client)
ツール [ 0 ] を待ちます。 ainvoke ({ "customer_id" : "acme" , "amount" : 1500.0 })
# Action レコードには、上記の推論にリンクする Decision_id が含まれるようになりました。
深さは、 ROOTSIGN_REASONING_DEPTH を介して推論をどの程度保持するかを制御します。
フラグがオフのときに ctx.record_decion() を呼び出すと、サイレントな no-op が行われます。呼び出しサイトで条件を指定せずに、キャプチャ オンおよびキャプチャ オフの環境で安全に出荷できます。 1 つの決定は 1 つの行動につながります。保留中のスロットは単一であり、次のツール呼び出し後にクリアされます。

それはうめぇ。決定はハッシュ チェーンに含まれていません (ADR-008) — verify_chain は変更されていません。
CrewAI の統合も同じ形で、構築時にツール リストをラップします。
pip install rootsign[crewai]
ルートサインをインポートする
crewai輸入代理店より
クルーアイから。ツールインポートツール
@ツール ( "send_invoice" )
def send_invoice (customer_id : str , amount : float ) -> str :
"""顧客に請求書を送信します。"""
「送信済み」を返す
async def run_crew (agent_id):
db として AsyncSessionLocal () を使用した非同期:
client = LocalIngestClient ( db = db )
rootsign と非同期。セッション (agent_id = Agent_id 、 client = client ) として ctx :
ラップされた = ルートサイン。 Wrap_crewai_tools (
[ send_invoice ]、ctx = ctx、client = クライアント
)
エージェント = エージェント (
役割 = "請求書作成アシスタント" ,
目標 = "請求書の送信" ,
ツール = ラップされた、
)
# ...通常どおり乗組員を実行します
データベースを待ちます。コミット()
CrewAI 0.28 、 0.40 、および 1.x に対してテスト済み (CI マトリックスを参照)。
リスクの高い行動は人間の判断に基づいて制御できます。 require_approval=True を @rootsign.trace に渡すと、誰かが CLI 経由で承認するまで、SDK はツールの実行をブロックします。
ルートサインをインポートする
@ ルートサイン 。トレース (
ingest_client = クライアント 、
セッションコンテキスト = ctx 、
require_approval = True 、
timeout_秒 = 300 、 # 5 分
)
async def Wire_transfer (アカウント: str 、金額: float ) -> str :
# これは人間が承認した後にのみ実行されます。
returnexecute_transfer (アカウント、金額)
Wire_transfer(...) が呼び出されると、SDK は authorization_status='pending' を指定して ACTION_RECORD を挿入し、待機します。オペレーターが別の端末から承認 (または拒否) します。
$ rootsign accept --list
承認待ち (1):
< アクション ID > Wire_transfer session= < セッション ID > submit= < タイムスタンプ >
$ rootsign accept < action-id > --reason " 顧客と確認済み "
✓ アクション < action-id > が承認されました。
装飾された機能

正常に戻ります。拒否 ( --reject ) により HiTLRejectedError が発生します。 5 分間のタイムアウトでは HiTLTimeoutError が発生し、アクションの承認ステータスは 'timed_out' ( 'human_rejected' とは異なる端末フォレンジック状態) になります。
設計の理論的根拠 (ポーリング ループ、タイムアウト セマンティクス、競合耐性) については、ADR-007 を参照してください。
RedactionConfig はハッシュ化の前に実行されるため、保存された input_hash / Output_hash 値には PII 信号が含まれません。 Three ready-to-use configs:
rootsign からインポート StandardPIIConfig 、 FinancialPIIConfig 、 HealthcarePIIConfig
# 標準: 電子メール、電話、米国 SSN、クレジット カード、英国 NI 番号
redaction = StandardPIIConfig ()
tools = ルートサイン 。ラップツール (
[ send_invoice ], ctx = ctx , client = client ,
redaction_config = redaction ,
)
FinancialPIIConfig はアカウント/ルーティング/IBAN パターンを追加します。 HealthcarePIIConfig adds MRN / NPI / DOB.それぞれは、サブクラス化せずにドメイン固有のパターンの extra_rules={...} を受け入れます。 ADR-006 を参照してください。
@rootsign.trace は呼び出し可能なツールをラップし、呼び出しごとに ACTION_RECORD エンベロープを発行します。 LangGraph BaseTool および CrewAI ツールは自動的に検出されます。
LocalIngestClient は、v0.1.x のインプロセス取り込みパスです。ホストされたバックエンドの HttpIngestClient はフェーズ 2 に移行します。
ハッシュ チェーンはセッションごとです。各アクションには prev_action_hash が含まれるため、チェーンを再構築すると事後の変更が検出されます。
HiTLCheckpoint は、サイクルごとに独自の DB セッションを開く非同期ポーリング ループです。ループ バインディングの理論的根拠については、ADR-007 を参照してください。
Storage is PostgreSQL 16 + TimescaleDB 2.14. The actions table is a hypertable; the chain stays intact across chunks.
フェーズ 2 クラウド バックエンド — HttpIngestClient + ホスト型コンプライアンス ダッシュボード。 LocalIngestClient のドロップイン代替品が利用可能になりました。
HiTL 用 Web UI — CLI ではなくブラウザから保留中のアクションを承認/拒否します。
オートジー

n 統合 — CrewAI と同じダックタイピング形状。
現在のロードマップについては、GitHub の問題を参照してください。
寄付を歓迎します。開発セットアップ、コーディング標準、PR プロセスについては、CONTRIBUTING.md を参照してください。投稿を送信すると、CLA に同意したことになります。
オープンソース コミュニティ チャネルと Discord が近日公開されます。現時点では、GitHub Issues がバグを報告し、機能を提案するための正規の場所です。
Apache License 2.0 — 「ライセンス」と「通知」を参照してください。
脆弱性を報告するには、 SECURITY.md を参照してください。 GitHub の公開問題を開かないでください。
Rootsign は、AI エージェント向けのオープンソースの改ざん明示決定およびアクションの出所ログ ライブラリです。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
3
v0.1.2 — LangGraph ToolCall パスの JSON セーフな正規化
最新の
2026 年 6 月 18 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Rootsign is an open-source tamper-evident decision and action provenance logging library for AI agents - Providex-AI/rootsign

I have been running into the same problems recently instrumenting my LangChain and CrewAI agent pipelines. If something goes wrong with a tool call, there was no way to way to prove what the agent did, in what order, and whether the logs have been modified. Observability platforms like LangSmith and Langfuse are great at optimizing for agent behavior, tokens and costs but they do not produce legally defensible and auditable artifacts. That's the reason why I decided to build RootSign, an SDK that instruments your CrewAI and LangGraph agents and produces cryptographic audit logs for CrewAI and LangGraph agents. RootSign adds a cryptographic hash chain to every tool call. If any record is modified after the fact, "rootsign verify" detects it. What it does: - SHA-256 hash chain across every Action record in a session
- Human-in-the-loop checkpoints with Approval records for certain agent actions
- PII redacted before hashing (StandardPIIConfig out of the box)
- Works with LangGraph and CrewAI — AutoGen coming soon
- Local first (Postgres + Timescale) — no cloud dependency What it doesn't do (yet): compliance dashboard, cloud backend, policy engine, all on the roadmap. Please try it out on the Github repo, contributions and feedback are always welcome.

GitHub - Providex-AI/rootsign: Rootsign is an open-source tamper-evident decision and action provenance logging library for AI agents · GitHub
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
Providex-AI
/
rootsign
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
32 Commits 32 Commits .github .github .vscode .vscode docs docs examples examples rootsign rootsign scripts scripts tests tests .env.example .env.example .gitignore .gitignore CLA.md CLA.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTORS.md CONTRIBUTORS.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md alembic.ini alembic.ini docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Tamper-evident provenance logging for production AI agents.
RootSign is a Providex AI product — the agent capture layer of the Providex AI Agent Accountability Platform.
When AI agents take actions in production — calling tools, hitting APIs, writing to databases — there is no built-in audit trail. If something goes wrong (a wrong refund, a leaked PII record, a malformed deployment), there is no way to prove what the agent did, in what order, on whose authorization, or whether the record has been tampered with after the fact.
RootSign solves this. Each agent action is captured as an Action record containing a SHA-256 hash of the previous action — a cryptographic hash chain that makes the record tamper-evident. Modify any record after the fact and rootsign verify detects it.
Compliance-grade audit trails. Zero changes to your agent code.
Phase 1 MVP — v0.1.1. LangGraph + CrewAI integrations, rootsign verify CLI, PII redaction, human-in-the-loop checkpoints, and opt-in decision capture (PRD-19 / ADR-008) are all shipping.
Python 3.11 or 3.12 recommended. RootSign itself supports 3.11+, but the [crewai] extra currently lags on 3.13/3.14 wheels. If you hit No matching distribution found for crewai , switch to Python 3.12 and reinstall.
pip install rootsign[langgraph]
Start PostgreSQL + TimescaleDB locally and apply the schema:
rootsign-admin start-db # docker run timescale/timescaledb:latest-pg16
rootsign-admin init # alembic upgrade head
start-db wraps a single docker run so you don't need to clone the repo. If you have cloned it, docker-compose up -d db is the equivalent developer path. Both reuse the same rootsign-timescaledb container name and rootsign_pgdata volume — pick either, not both.
2. Register your agent (one-time setup)
import asyncio
from rootsign import register_agent , AgentEnvironment , AgentRiskTier , AgentFramework
agent = asyncio . run ( register_agent (
name = "my-invoice-agent" ,
owner = "platform-team" ,
environment = AgentEnvironment . PRODUCTION ,
risk_tier = AgentRiskTier . HIGH ,
framework = AgentFramework . LANGGRAPH ,
))
print ( agent . agent_id )
3. Instrument your tools
import rootsign
from rootsign import LocalIngestClient
from rootsign . database import AsyncSessionLocal
from langchain_core . tools import tool
from langgraph . prebuilt import ToolNode
@ tool
def send_invoice ( customer_id : str , amount : float ) -> str :
"""Send an invoice to a customer."""
return "sent"
async def run_graph ( agent_id ):
async with AsyncSessionLocal () as db :
client = LocalIngestClient ( db = db )
async with rootsign . session ( agent_id = agent_id , client = client ) as ctx :
tools = rootsign . wrap_tools ([ send_invoice ], ctx = ctx , client = client )
tool_node = ToolNode ( tools )
# ...build and run your graph as normal
await db . commit ()
Every tool call now produces a tamper-evident Action record on the hash chain.
$ rootsign verify 660e8400-e29b-41d4-a716-446655440001
VALID ✓ — 3 records, chain intact
Session: 660e8400-e29b-41d4-a716-446655440001
Exit code is 0 for VALID, 1 for TAMPERED. Use --local <path.jsonl> for offline JSONL session files (no DB required).
If a record was modified, the verifier names the broken link:
$ rootsign verify 660e8400-e29b-41d4-a716-446655440001
TAMPERED ✗ — chain broken at record # 2
Detail: self_hash mismatch on action < action_id >
Session: 660e8400-e29b-41d4-a716-446655440001
WARNING: This session log may have been tampered with.
See docs/framework-support.md for the version matrix and integration notes. A full runnable LangGraph example (ReAct agent, three instrumented tools, OpenAI-backed) lives in examples/langgraph-invoice-agent .
Record the why before each tool call — foundational for Phase 2 session replay. Off by default; opt in deliberately with ROOTSIGN_CAPTURE_DECISIONS=true .
import os
os . environ [ "ROOTSIGN_CAPTURE_DECISIONS" ] = "true"
async with rootsign . session ( agent_id = agent . agent_id , client = client ) as ctx :
# Record what the agent decided before calling the tool.
await ctx . record_decision (
selected_action = "send_invoice" ,
reasoning_summary = "Amount within policy; recipient verified." ,
confidence = 0.97 ,
ingest_client = client ,
)
tools = rootsign . wrap_tools ([ send_invoice ], ctx = ctx , client = client )
await tools [ 0 ]. ainvoke ({ "customer_id" : "acme" , "amount" : 1500.0 })
# The Action record now carries decision_id linking it to the reasoning above.
Depth controls how much reasoning is persisted, via ROOTSIGN_REASONING_DEPTH :
Calling ctx.record_decision() when the flag is off is a silent no-op — safe to ship in capture-on and capture-off environments without conditionals at the call site. One Decision links to one Action ; the pending slot is single and cleared after the next tool call consumes it. Decisions are not in the hash chain (ADR-008) — verify_chain is unchanged.
CrewAI integration is the same shape — wrap the tool list at construction time.
pip install rootsign[crewai]
import rootsign
from crewai import Agent
from crewai . tools import tool
@ tool ( "send_invoice" )
def send_invoice ( customer_id : str , amount : float ) -> str :
"""Send an invoice to a customer."""
return "sent"
async def run_crew ( agent_id ):
async with AsyncSessionLocal () as db :
client = LocalIngestClient ( db = db )
async with rootsign . session ( agent_id = agent_id , client = client ) as ctx :
wrapped = rootsign . wrap_crewai_tools (
[ send_invoice ], ctx = ctx , client = client
)
agent = Agent (
role = "Invoicing assistant" ,
goal = "Send invoices" ,
tools = wrapped ,
)
# ...run your crew as normal
await db . commit ()
Tested against CrewAI 0.28 , 0.40 , and 1.x (see CI matrix ).
High-risk actions can be gated on a human decision. Pass require_approval=True to @rootsign.trace and the SDK blocks the tool from running until someone approves it via the CLI.
import rootsign
@ rootsign . trace (
ingest_client = client ,
session_context = ctx ,
require_approval = True ,
timeout_seconds = 300 , # 5 minutes
)
async def wire_transfer ( account : str , amount : float ) -> str :
# This runs ONLY after a human approves.
return execute_transfer ( account , amount )
When wire_transfer(...) is called, the SDK inserts an ACTION_RECORD with authorization_status='pending' and waits. An operator approves (or rejects) from another terminal:
$ rootsign approve --list
Pending approvals (1):
< action-id > wire_transfer session= < session-id > submitted= < timestamp >
$ rootsign approve < action-id > --reason " Verified with customer "
✓ Action < action-id > approved.
The decorated function returns normally. Rejection ( --reject ) raises HiTLRejectedError ; a 5-minute timeout raises HiTLTimeoutError and the action's authorization status becomes 'timed_out' (a terminal forensic state distinct from 'human_rejected' ).
See ADR-007 for the design rationale (poll loop, timeout semantics, race tolerance).
RedactionConfig runs before hashing, so stored input_hash / output_hash values carry no PII signal. Three ready-to-use configs:
from rootsign import StandardPIIConfig , FinancialPIIConfig , HealthcarePIIConfig
# Standard: email, phone, US SSN, credit card, UK NI number
redaction = StandardPIIConfig ()
tools = rootsign . wrap_tools (
[ send_invoice ], ctx = ctx , client = client ,
redaction_config = redaction ,
)
FinancialPIIConfig adds account / routing / IBAN patterns; HealthcarePIIConfig adds MRN / NPI / DOB. Each accepts extra_rules={...} for domain-specific patterns without subclassing. See ADR-006 .
@rootsign.trace wraps a tool callable and emits an ACTION_RECORD envelope per call. LangGraph BaseTool and CrewAI tools are detected automatically.
LocalIngestClient is the in-process ingest path for v0.1.x. A HttpIngestClient for the hosted backend lands in Phase 2.
Hash chain is per-session: each Action carries prev_action_hash so reconstructing the chain detects any after-the-fact modification.
HiTLCheckpoint is an async poll loop that opens its own DB session per cycle — see ADR-007 for the loop-binding rationale.
Storage is PostgreSQL 16 + TimescaleDB 2.14. The actions table is a hypertable; the chain stays intact across chunks.
Phase 2 cloud backend — HttpIngestClient + hosted compliance dashboard. Drop-in replacement for LocalIngestClient once available.
Web UI for HiTL — approve/reject pending actions from a browser instead of the CLI.
AutoGen integration — same duck-typing shape as CrewAI.
Watch the GitHub Issues for the active roadmap.
We welcome contributions. See CONTRIBUTING.md for development setup, coding standards, and the PR process. By submitting a contribution, you agree to the CLA .
Open-source community channels and Discord coming soon — for now, GitHub Issues is the canonical place to file bugs and propose features.
Apache License 2.0 — see LICENSE and NOTICE .
To report a vulnerability, see SECURITY.md . Do not open a public GitHub issue.
Rootsign is an open-source tamper-evident decision and action provenance logging library for AI agents
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
3
v0.1.2 — JSON-safe normalization for LangGraph ToolCall path
Latest
Jun 18, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
