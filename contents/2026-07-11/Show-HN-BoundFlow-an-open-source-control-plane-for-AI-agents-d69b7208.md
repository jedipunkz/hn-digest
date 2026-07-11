---
source: "https://github.com/boundflow/boundflow"
hn_url: "https://news.ycombinator.com/item?id=48875888"
title: "Show HN: BoundFlow – an open-source control plane for AI agents"
article_title: "GitHub - boundflow/boundflow: Open-source control plane for AI agents that take real actions: policy-dictated lifecycle controls, approvals, durable execution, audit receipts, and observability. · GitHub"
author: "alama24"
captured_at: "2026-07-11T21:40:07Z"
capture_tool: "hn-digest"
hn_id: 48875888
score: 1
comments: 0
posted_at: "2026-07-11T21:07:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: BoundFlow – an open-source control plane for AI agents

- HN: [48875888](https://news.ycombinator.com/item?id=48875888)
- Source: [github.com](https://github.com/boundflow/boundflow)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T21:07:11Z

## Translation

タイトル: Show HN: BoundFlow – AI エージェント用のオープンソース コントロール プレーン
記事のタイトル: GitHub - バウンドフロー/バウンドフロー: 実際のアクションを実行する AI エージェント用のオープンソース コントロール プレーン: ポリシーに基づくライフサイクル制御、承認、永続的な実行、監査の受信、および可観測性。 · GitHub
説明: 実際のアクションを実行する AI エージェント用のオープンソース コントロール プレーン: ポリシーに基づいたライフサイクル制御、承認、永続的な実行、監査の受信、可観測性。 - バウンドフロー/バウンドフロー

記事本文:
GitHub -boundflow/boundflow: 実際のアクションを実行する AI エージェント用のオープンソース コントロール プレーン: ポリシーに基づくライフサイクル制御、承認、永続的な実行、監査の受信、および可観測性。 · GitHub
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
ロード中にエラーが発生しました

してる。このページをリロードしてください。
バウンドフロー
/
バウンドフロー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
53 コミット 53 コミット .github/ workflows .github/ workflows cmd/boundflow cmd/boundflow docs docs gen/boundflow/ v1 gen/boundflow/ v1 内部 内部移行 移行 proto/boundflow/ v1 proto/boundflow/ v1 sdk/ python sdk/ python .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile QUICKSTART.md QUICKSTART.md README.md README.md buf.gen.yaml buf.gen.yaml buf.yaml buf.yaml docker-compose.dist.yml docker-compose.dist.yml docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum mkdocs.yml mkdocs.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
無人で実行する LLM エージェントとワークフローの運用レイヤー - コントロール プレーンによって適用される、コスト キャップ、承認ゲート、自己修復ポリシー。
パブリック プレビュー (1.0 より前)。エンジンは完成しており、Go、mock-LLM、
およびライブ LLM テスト スイートが含まれていますが、外部ツールを使用して実稼働環境でまだ実行されていません。
ユーザー。 gRPC protobuf を含む API は 1.0 より前に変更される可能性があります。探しています
早期採用者および設計パートナーの場合: お問い合わせください。
BoundFlow は、長時間実行されるステートフルなエージェント ワークフローを実行し、ガードレールを強制します
エージェントを無人で実行する前に必要なこと: 実行ごとのコスト上限、モデルの自動切り替えをオンにする
コスト/ループ ポリシー、機密性の高いアクションの前の人間による承認ゲート、ツール呼び出し
制限、再試行、クールダウン、バージョン管理されたロールバック。エージェントを書いて、
クリーンな非同期 SDK に対するワークフロー。コントロール プレーンのスケジュール、ディスパッチ、
彼らを統治している。
で

リファレンスは持ち込みです - エージェントはあなた自身の人間性を持ってクロードに電話します
キー、ワーカー内で実行されます。バックエンドはそれを決して認識せず、トークンの料金を支払うこともありません。
キー、データ、トークンの使用量は、ワイヤーの側に残ります。
実際: 実行あたり最大 $0.25 を費やす可能性があるサポート優先順位付けワークフローでは、
払い戻しを行う前に人間の承認を得て、払い戻しを行う場合は Haiku にダウングレードします。
コストが急上昇し、起動すると最新の正常なバージョンに自動的にロールバックされる
失敗します。エージェント コードにはそのようなロジックは存在しません。それをポリシーとして宣言します。
コントロール プレーンはそれを強制し、耐久性があり、クエリ可能な監査ログをすべてのログに保存します。
承認と政策決定。
BoundFlow は、プロンプト フレームワークでも、推論プロバイダーでも、エージェント ビルダーでもありません。
これは、構築するエージェントの周囲の運用層です。
バックエンド — オープン ソース (Apache-2.0)、コンテナとして自己ホスト可能。
Python SDK — オープンソース (MIT)、pip installboundflow 。
ドキュメント — 概念、ガバナンス、デプロイメント、および API リファレンス ( docs/ )。
BoundFlow クラウド — 自己ホストしたくないですか?マネージド ホスティング (早期アクセス) — 以下を参照してください。
実際のアクションを実行するエージェントには、次のような場合に実際のアクションを実行するコントロール プレーンが必要です。
彼らは間違ってしまいます。ほとんどのツールはエージェントを監視します。 BoundFlow は両方に介入します
レベル。エージェント側: 支出に上限を設け、実行中にモデルを交換します。で
ワークフロー : 人間による承認のための危険なステップをゲートし、冷却し、元の状態にロールバックします。
正常なバージョンを使用するか、完全に一時停止します。ワークフローを認識するだけでなく、
エージェント対応 — モデル呼び出しだけでなく、ワークフロー全体を実行するためです。
各実行のスケジュール設定、ステップ間での状態の保持、障害からの回復、
エージェントを、システム内の 1 つの操作として、ライフサイクルを通じて実行します。
耐久性に優れた複数段階のプロセスをエンドツーエンドで所有しています。
エージェントが無人で実行される瞬間に、次のような答えが必要になります。

ループしたらどうなるの？もしも
50ドルかかりますか？取り返しのつかないことが起こりそうになったらどうしますか？どのモデルにするべきか
を使用していますが、いつ変更する必要がありますか? BoundFlow は、代わりにこれらのポリシーを作成します。
コード:
ポリシーはサーバー側 (ライフサイクル) で評価され、SDK 側 (ランタイム) で適用されます。
呼び出しごとのメトリクス (コスト、トークン、LLM 呼び出し、ツールごとの数/失敗数) を使用します。
実行ごとに収集されます。
BoundFlow バックエンドはコントロール プレーンです。自己ホストするか、上で実行します。
バウンドフロークラウド。いずれの場合でも、ワーカーは gRPC 経由でワーカーに接続し、
Anthropic キーを備えた実際のエージェントが環境内に存在します。バックエンドのスケジュール、
ディスパッチ、管理、監査を行い、キーや推論トラフィックを決して見ることはありません。
┌─────────────┐ gRPC ┌─────────────┐
│ クライアント / SDK │ ───────────▶ │ │
━━━━━━━━━┘ 呼び出し・承認 │ BoundFlow バックエンド │
·クエリ │ (コントロールプレーン) │
┌─────────────┐ gRPC ストリーム │ │
│ あなたのワーカー │ ◀───────────▶ │ スケジュール・派遣 │
│ エージェント+ツールの実行 │ 起動/結果 │ ・管理・監査 │
│ API キーを使用して │ └───────────┘
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
内部では、バックエンドは 3 つのプロセス モード (サーバー、スケジューラー、
ワーカー ) 1 つのバイナリ共有 Postgres をオフにする — を参照
完全な内訳とライフサイクルについては docs/concepts.md
州。
fromboundflow import AgentDefinition 、 BoundFlowWorker 、 Complete 、 ControlPlaneClient 、 Work

フロー構成
バウンドフローから。 anthropic_client インポート AnthropicLlmClient
worker = BoundFlowWorker ( llm = AnthropicLlmClient (...)) # エンドポイント + env からのキー
@ ワーカー 。ワークフロー (「トリアージ」、バージョン = 1)
非同期デフォルトトリアージ (ctx):
ctx 。 add_context ( "チケット" , "..." )
ctx を待ちます。 run_agent ( エージェント定義 (
名前 = "アナリスト" 、モデル = "claude-haiku-4-5" 、
system_prompt = "問題を診断します。" 、output_schema = { "概要" : { "タイプ" : "文字列" }},
))
完了を返す ()
LangChain 経由で独自のプロバイダーを導入します。ツール呼び出しの LangChain チャットをラップする
LangChainLlmClient のモデルとガバナンスは同一です — OpenAI、Google、
Bedrock と LangChain の残りのエコシステムは同じコスト上限、モデルの下で実行されます。
ポリシーと承認ゲート:
from langchain_anthropic import ChatAnthropic # または ChatOpenAI、ChatVertexAI、...
バウンドフローから。 langchain_client インポート LangChainLlmClient
worker = BoundFlowWorker (llm = LangChainLlmClient (ChatAnthropic (model = "claude-haiku-4-5" )))
pip install "boundflow[langchain]" でインストールします。参照してください
boundflow.examples.langchain_adapter
実行可能なエンドツーエンドの例については、
BoundFlow によって管理される LangGraph とオーケストレーションします。 LangGraph エージェントを構築する
ctx.run_agent を呼び出すノードを含むワークフロー内のグラフなので、LangGraph
はルーティングを所有し、BoundFlow はエージェントのすべてのステップとワークフローを管理します。
全体。 「統合と実行可能ファイル」を参照してください。
boundflow.examples.langgraph_workflow 。
ワークフローはマルチステップでステートフルです: 操作は人間に代わって実行できます
決定または後続の操作に連鎖し、ワークフローはその時点から再開されます。
left off — 後ろにゲートされたブランチが実行されるまで、元に戻せないものは何も実行されません。
fromboundflow import AwaitApproval 、 Next 、 Complete
@ ワーカー 。ワークフロー ( "払い戻し" 、バージョン = 1 )
非同期デフォルト返金 (ctx):
ctx を待ちます。走る

_agent (アナリスト) # ステップ 1: リクエストの理由
return AwaitApproval ( # park — まだ元に戻せないものは何もありません
on_approve = 次へ ( "issue_refund" , ctx . context ),
on_reject = 完了 ()、
justification = 「5,000 ドルの払い戻しを承認しますか?」 、
)
@ ワーカー 。操作 ( "refund" 、 "issue_refund" ) # ステップ 2: 人間の承認後にのみ実行されます
非同期デフォルト issue_refund ( ctx ):
... # デリケートな行為、現在は認可されています
完了を返す ()
ガバナンスはコントロール プレーンから適用されます - 3 つのレイヤー、実行ごとの上限に基づいて適用されます
自己修復バージョンのロールバックへ:
バウンドフローインポートから (
RuntimePolicy 、 AgentRule 、 AgentMetric 、 Op 、 SetModel 、
WorkflowRule 、 WorkflowMetric 、 SetVersion 、
)
# 1. 実行時 — すべての実行中*に適用されるハードキャップ:
cp を待ちます。 set_agent_runtime_policy ( wf . id , "analyst" , RuntimePolicy ( max_cost_usd = 0.25 ))
# 2. エージェントのライフサイクル — 実行後、コストが高くなる傾向にある場合はモデルをダウングレードします。
cp を待ちます。 set_agent_lifecycle_policy ( wf . id , "アナリスト" , [
AgentRule ( metric = AgentMetric . COST_USD 、 op = Op . GT 、 しきい値 = 0.20 、 window = 5 、
アクション = SetModel (値 = "claude-haiku-4-5" ))、
])
# 3. ワークフローのライフサイクル — 失敗が繰り返された後、ワークフロー全体をロールバックする
# 既知の正常なバージョンに自動的に移行します。
cp を待ちます。 set_workflow_lifecycle_policy ( wf . id , [
WorkflowRule (メトリック = WorkflowMetric . NUM_FAILURES 、しきい値 = 3 、
アクション = SetVersion (ターゲット = 1 ))、
])
ワークフロー ルールでは、ワークフローを一時停止したり、代わりにクールダウンに置いたりすることもできます。
ロールバックします。実行可能な例については、sdk/python/boundflow/examples/ を参照してください。
管理されたエージェントを数分で実行できます。完全なウォークスルー: QUICKSTART.md 。
# 1. データベースのパスワードを設定します (任意の強力な秘密)
echo " BOUNDFLOW_DB_PASSWORD= $( openssl rand -hex 16 ) " > .env
# 2. バックエンドの起動 (Postgres + サーバー + スケジュール)

ラー + 労働者)
docker compose -f docker-compose.dist.yml up -d
# 3. API キーをプロビジョニングする
docker compose -f docker-compose.dist.yml run --rm サーバー -mode=provision -name=me
import BOUNDFLOW_API_KEY= < 印刷されたキー >
# 4. SDK をインストールし、Anthropic キーを持参します
pip インストール バウンドフロー
import ANTHROPIC_API_KEY= <あなたのキー>
#5. ガバナンスの下で実際のエージェントを実行する
python -mboundflow.examples.hello
次に、バンドルされているサンプルを調べます。
python -mboundflow.examples.approval_gate # 人間参加型サインオフ
これは、boundflow CLI (SDK とともにインストールされます) から管理および観察します。
boundflow ワークフロー リスト # ワークフローとその状態
boundflow ワークフローの実行 < id > # 実行とその結果 · スクリプト用の --json
可観測性
可観測性は第一級であり、OpenTelemetry ネイティブです。独自のフォーマットはありません。
ロックインがないため、すでに実行しているテレメトリ スタックに直接接続できます。 2
レイヤー: 実行トレース (独自のバックエンドにエクスポートする実行テレメトリ) と
ガバナンス監査ログ (決定、サーバー側に保持され、クエリ可能)。
トレースを実行します。すべての操作は、OperationTrace (操作 → エージェント → トークンの使用法と完全なプロンプト/応答コンテンツを含む llm/ツール ツリー) をプラガブル オブジェクトに出力します。
あなたが所有するシンク。組み込み: LoggingTraceSink 、 JsonlFileTraceSink 、および
OTelTraceSink : OpenTelemetry GenAI セマンティック規約にマッピングされ、出荷されます
OTLP を介して任意のバックエンド (Jaeger、Tempo、Langfuse、Phoenix など) に広がります。すべての操作
1 回の実行で、trace_id を共有します。
バウンドフローインポートからBoundFlowWorker
バウンドフローから。トレースインポート OTelTraceSink
ワーカー = BoundFlowWorker (llm = ...、trace_sink = OTelTraceSink (トレーサー))
実行可能ファイルについては、sdk/python/examples/otel/ を参照してください。
○

[切り捨てられた]

## Original Extract

Open-source control plane for AI agents that take real actions: policy-dictated lifecycle controls, approvals, durable execution, audit receipts, and observability. - boundflow/boundflow

GitHub - boundflow/boundflow: Open-source control plane for AI agents that take real actions: policy-dictated lifecycle controls, approvals, durable execution, audit receipts, and observability. · GitHub
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
boundflow
/
boundflow
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
53 Commits 53 Commits .github/ workflows .github/ workflows cmd/ boundflow cmd/ boundflow docs docs gen/ boundflow/ v1 gen/ boundflow/ v1 internal internal migrations migrations proto/ boundflow/ v1 proto/ boundflow/ v1 sdk/ python sdk/ python .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile QUICKSTART.md QUICKSTART.md README.md README.md buf.gen.yaml buf.gen.yaml buf.yaml buf.yaml docker-compose.dist.yml docker-compose.dist.yml docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum mkdocs.yml mkdocs.yml View all files Repository files navigation
The operational layer for the LLM agents and workflows you run unattended — cost caps, approval gates, and self-healing policy, enforced by a control plane.
Public preview (pre-1.0). The engine is complete and covered by Go, mock-LLM,
and live-LLM test suites, but it hasn't yet been run in production with external
users. APIs — including the gRPC protobufs — may change before 1.0. We're looking
for early adopters and design partners: reach out .
BoundFlow runs long-running, stateful agent workflows and enforces the guardrails
you'll want before running agents unattended: per-run cost caps , automatic model switching on
cost/loop policies, human approval gates before sensitive actions, tool-call
limits, retries, cooldowns, and versioned rollbacks. You write agents and
workflows against a clean async SDK; the control plane schedules, dispatches, and
governs them.
Inference is bring-your-own — your agents call Claude with your own Anthropic
key, running in your worker. The backend never sees it and never pays for tokens.
Your keys, your data, and your token spend stay on your side of the wire.
In practice: a support-triage workflow that may spend up to $0.25/run , must
get a human's sign-off before issuing a refund, downgrades to Haiku when
costs spike, and auto-rolls-back to the last good version if it starts
failing — none of that logic living in your agent code. You declare it as policy;
the control plane enforces it and keeps a durable, queryable audit log of every
approval and policy decision.
BoundFlow is not a prompt framework, an inference provider, or an agent-builder —
it's the operational layer around the agents you build.
Backend — open source (Apache-2.0), self-hostable as a container.
Python SDK — open source (MIT), pip install boundflow .
Docs — concepts, governance, deployment, and API reference in docs/ .
BoundFlow Cloud — prefer not to self-host? Managed hosting (early access) — see below .
Agents that take real actions need a control plane that takes real action when
they go wrong. Most tools watch your agents; BoundFlow intervenes — at both
levels. On the agent : cap its spend, swap its model mid-run. On the
workflow : gate a risky step for human sign-off, cool it down, roll it back to a
known-good version, or pause it outright. It's workflow-aware, not just
agent-aware — because it runs the whole workflow, not just the model call:
scheduling each run, carrying state across steps, recovering from failures, and
driving it through its lifecycle, with the agent as just one operation inside a
durable, multi-step process it owns end to end.
The moment agents run unattended you need answers to: What if it loops? What if
it spends $50? What if it's about to do something irreversible? Which model should
it use, and when should that change? BoundFlow makes those policies instead of
code:
Policies are evaluated server-side (lifecycle) and enforced SDK-side (runtime),
with per-invocation metrics — cost, tokens, LLM calls, per-tool counts/failures —
collected on every run.
The BoundFlow backend is the control plane — self-host it, or run it on
BoundFlow Cloud. Either way, your worker connects to it over gRPC and runs the
actual agents, with your Anthropic key, in your environment; the backend schedules,
dispatches, governs, and audits, and never sees your key or your inference traffic.
┌─────────────────────┐ gRPC ┌────────────────────────┐
│ Your client / SDK │ ───────────────▶ │ │
└─────────────────────┘ invoke·approve │ BoundFlow backend │
·query │ (control plane) │
┌─────────────────────┐ gRPC stream │ │
│ Your worker │ ◀──────────────▶ │ schedules·dispatches │
│ runs agents+tools │ launch/result │ ·governs·audits │
│ with your API key │ └────────────────────────┘
└─────────────────────┘
Under the hood the backend runs as three process modes ( server , scheduler ,
worker ) off one binary sharing Postgres — see
docs/concepts.md for the full breakdown and the lifecycle
states.
from boundflow import AgentDefinition , BoundFlowWorker , Complete , ControlPlaneClient , WorkflowConfig
from boundflow . anthropic_client import AnthropicLlmClient
worker = BoundFlowWorker ( llm = AnthropicLlmClient (...)) # endpoints + key from env
@ worker . workflow ( "triage" , version = 1 )
async def triage ( ctx ):
ctx . add_context ( "ticket" , "..." )
await ctx . run_agent ( AgentDefinition (
name = "analyst" , model = "claude-haiku-4-5" ,
system_prompt = "Diagnose the issue." , output_schema = { "summary" : { "type" : "string" }},
))
return Complete ()
Bring your own provider via LangChain. Wrap any tool-calling LangChain chat
model in LangChainLlmClient and the governance is identical — OpenAI, Google,
Bedrock, and the rest of LangChain's ecosystem run under the same cost caps, model
policies, and approval gates:
from langchain_anthropic import ChatAnthropic # or ChatOpenAI, ChatVertexAI, ...
from boundflow . langchain_client import LangChainLlmClient
worker = BoundFlowWorker ( llm = LangChainLlmClient ( ChatAnthropic ( model = "claude-haiku-4-5" )))
Install with pip install "boundflow[langchain]" ; see
boundflow.examples.langchain_adapter
for a runnable end-to-end example.
Orchestrate with LangGraph, governed by BoundFlow. Build a LangGraph agent
graph inside a workflow with its nodes calling ctx.run_agent , so LangGraph
owns the routing while BoundFlow governs every agent step and the workflow as a
whole. See Integrations and the runnable
boundflow.examples.langgraph_workflow .
Workflows are multi-step and stateful : an operation can park for a human
decision or chain into a follow-on operation, and the workflow resumes where it
left off — nothing irreversible runs until the branch it's gated behind does.
from boundflow import AwaitApproval , Next , Complete
@ worker . workflow ( "refund" , version = 1 )
async def refund ( ctx ):
await ctx . run_agent ( analyst ) # step 1: reason about the request
return AwaitApproval ( # park — nothing irreversible yet
on_approve = Next ( "issue_refund" , ctx . context ),
on_reject = Complete (),
justification = "Approve the $5,000 refund?" ,
)
@ worker . operation ( "refund" , "issue_refund" ) # step 2: runs only after a human approves
async def issue_refund ( ctx ):
... # the sensitive action, now sanctioned
return Complete ()
Governance is applied from the control plane — three layers, from a per-run cap
to self-healing version rollback:
from boundflow import (
RuntimePolicy , AgentRule , AgentMetric , Op , SetModel ,
WorkflowRule , WorkflowMetric , SetVersion ,
)
# 1. Runtime — a hard cap enforced *during* every run:
await cp . set_agent_runtime_policy ( wf . id , "analyst" , RuntimePolicy ( max_cost_usd = 0.25 ))
# 2. Agent lifecycle — after runs, downgrade the model if cost trends high:
await cp . set_agent_lifecycle_policy ( wf . id , "analyst" , [
AgentRule ( metric = AgentMetric . COST_USD , op = Op . GT , threshold = 0.20 , window = 5 ,
action = SetModel ( value = "claude-haiku-4-5" )),
])
# 3. Workflow lifecycle — after repeated failures, roll the whole workflow back
# to a known-good version automatically:
await cp . set_workflow_lifecycle_policy ( wf . id , [
WorkflowRule ( metric = WorkflowMetric . NUM_FAILURES , threshold = 3 ,
action = SetVersion ( target = 1 )),
])
Workflow rules can also Pause a workflow or put it on Cooldown instead of
rolling back. See sdk/python/boundflow/examples/ for runnable examples.
Get a governed agent running in a few minutes. Full walkthrough: QUICKSTART.md .
# 1. Set a database password (any strong secret)
echo " BOUNDFLOW_DB_PASSWORD= $( openssl rand -hex 16 ) " > .env
# 2. Start the backend (Postgres + server + scheduler + worker)
docker compose -f docker-compose.dist.yml up -d
# 3. Provision an API key
docker compose -f docker-compose.dist.yml run --rm server -mode=provision -name=me
export BOUNDFLOW_API_KEY= < printed key >
# 4. Install the SDK and bring your Anthropic key
pip install boundflow
export ANTHROPIC_API_KEY= < your key >
# 5. Run a real agent under governance
python -m boundflow.examples.hello
Then explore the bundled examples:
python -m boundflow.examples.approval_gate # human-in-the-loop sign-off
Manage and observe it from the boundflow CLI (installed with the SDK):
boundflow workflow list # your workflows and their state
boundflow workflow runs < id > # runs and their outcomes · --json for scripting
Observability
Observability is first-class and OpenTelemetry-native — no proprietary format,
no lock-in, so it plugs straight into the telemetry stack you already run. Two
layers: run traces (execution telemetry you export to your own backend) and a
governance audit log (decisions, kept server-side and queryable).
Run traces. Every operation emits an OperationTrace — the operation → agent → llm/tool tree with token usage and full prompt/response content — to a pluggable
sink you own. Built-ins: LoggingTraceSink , JsonlFileTraceSink , and
OTelTraceSink , which maps onto OpenTelemetry GenAI semantic conventions and ships
spans over OTLP to any backend (Jaeger, Tempo, Langfuse, Phoenix, …); all operations
of one run share a trace_id .
from boundflow import BoundFlowWorker
from boundflow . trace import OTelTraceSink
worker = BoundFlowWorker ( llm = ..., trace_sink = OTelTraceSink ( tracer ))
See sdk/python/examples/otel/ for a runnable
O

[truncated]
