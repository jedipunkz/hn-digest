---
source: "https://github.com/AgentBull/ankole"
hn_url: "https://news.ycombinator.com/item?id=48896706"
title: "Show HN: Ankole – Claude Tag open source alternative"
article_title: "GitHub - AgentBull/ankole: Ankole is an open-source, self-hosted AgentOS for shared AI colleagues. Claude Tag open source alternative. · GitHub"
author: "borisding1994"
captured_at: "2026-07-13T19:12:48Z"
capture_tool: "hn-digest"
hn_id: 48896706
score: 2
comments: 0
posted_at: "2026-07-13T18:26:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ankole – Claude Tag open source alternative

- HN: [48896706](https://news.ycombinator.com/item?id=48896706)
- Source: [github.com](https://github.com/AgentBull/ankole)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T18:26:08Z

## Translation

タイトル: 表示 HN: Ankole – クロード タグのオープンソース代替案
記事タイトル: GitHub - AgentBull/ankole: Ankole は、共有 AI 同僚向けのオープンソースの自己ホスト型 AgentOS です。クロードタグのオープンソース代替品。 · GitHub
説明: Ankole は、共有 AI 同僚向けのオープンソースの自己ホスト型 AgentOS です。クロードタグのオープンソース代替品。 - エージェントブル/アンコール

記事本文:
GitHub - AgentBull/ankole: Ankole は、共有 AI 同僚向けのオープンソースの自己ホスト型 AgentOS です。クロードタグのオープンソース代替品。 · GitHub
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
読み込み中にエラーが発生しました。このページをリロードしてください。
エージェントブル
/
アンコール
公共
そうではない

化
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
77 コミット 77 コミット .agents/ スキル .agents/ スキル .devcontainer .devcontainer .github/ workflows .github/ workflows app app docs docs libs libs plugins plugins tools tools .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .oxfmtrc.json .oxfmtrc.json .oxlintrc.json .oxlintrc.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTEXT.md CONTEXT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.ja.md README.ja.md README.md README.md README.zh-Hans.md README.zh-Hans.md SECURITY.md SECURITY.md ankole.code-workspace ankole.code-workspace bun.lock bun.lock bunfig.toml bunfig.toml knip.config.ts knip.config.ts konsistent.json konsistent.json package.json package.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.jsonturbo.jsonturbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Ankole - 共有 AI 同僚向けの Open AgentOS
どう違うのか・製品の形状・アクターのランタイム・アーキテクチャー・現状・開発
Ankole は、共有 AI 同僚向けの自己ホスト型 AgentOS です。 1 つのインストール、多数のエージェント、実際の責任 - 管理するインフラストラクチャ上で。
AI の作業をプライベート チャット ボックスからチャネル、リポジトリ、スケジュール、ダッシュボード、内部システム、長期にわたるプロジェクトのコンテキストなど、すでに作業が行われている場所に移動します。 Ankole エージェントは独自の ID、メモリ、権限、ツール、ワークスペース、および責任境界を持っているため、1 回限りのメッセージに応答するだけでなく、継続的な作業を所有できます。
Claude Tag は有用な公開参照点です。AI を S にタグ付けします。

スレッドが不足している場合は、スレッドに共有コンテキストを読み取らせ、整理ツールを使用し、チャネルのコンテキストを記憶し、作業に時間がかかる場合はフォローアップします。 Ankole は、そのパターンのより広範なオープン バージョンをターゲットとしています。Slack だけでなく、Claude だけでもなく、1 人のエージェントだけでもなく、ベンダー所有のコンテキストでもありません。
Ankole は、単なる答えではなく、所有者が必要な仕事に適しています。 Ankole の優れた役割には、コードのマージ、レポートの発送、顧客の問題の処理、アラートのトリアージ、市場の変化の認識、バックログの削減など、目に見える結果が伴います。
プライベートチャットではなく、デフォルトで共有されます。エージェントはチームに表示されるチャネルとプロバイダー コンテキストに参加します。複数の人間が同じ作業を観察し、操作し、継続することができます。
即時的な慣例ではなく、永続的なアイデンティティ。人間とエージェントは権限付与と監査証跡を持つプリンシパルであるため、承認は実行時の懸念事項となります。
リクエスト/レスポンスではなく、長時間実行されるアクター セッション。セッションは、コンテキストに応じて起動、信号の受信、チェックポイント、ストリームの進行状況、休止状態、および回復します。
ベンダーがホストするコンテキストではなく、オペレーターが所有するコンテキスト。メモリ、構成、認証情報、監査は、セルフホスト型インストールのインフラストラクチャ内に存在します。
ライブコントロールと永続的な真実、どちらか一方ではありません。 ZeroMQ RuntimeFabric はライブ アクター/ワーカー/RPC トラフィックを伝送しますが、PostgreSQL はリプレイ、フェンス、最終コミットのソースのままです。
多くの情報源。 IM、Webhook、スケジュールされたリマインダー、内部システム、および将来のプロバイダー アダプターはすべて、正規化された信号入力になります。
多くのエージェント。 1 つのインストールで、さまざまなミッション、アクセス、ツール、メモリ、送信 ID を持つ複数のエージェントをホストできます。
セッションアクター。長時間実行される実行ユニットは、actor_id = {agent_id, session_id} です。セッションは、コンテキスト、ワークスペースの状態、ステアリング、キャンセル、および回復が交わる場所です。
所有されたコンテキスト。会話、モデルターン、要約、シグナルプロジェクト

イオン、決定、修正、および将来のドメイン記録はインフラストラクチャ内に存在します。
オペレーター制御。アクセス、構成、プラグインのアクティブ化、アクターのリース、送信ボックスの副作用、および監査サーフェスは、インストール オペレーターに属します。
Ankole はこれらのワークフローを自然にする必要があります。
コーディング エージェントは問題を監視し、バグを再現し、コードを変更し、ドラフト PR を開き、人間の決定がまだ必要なものを報告します。
カスタマー サクセス エージェントは、共有グループ チャットを観察し、重要な事実を記録し、作業状況を更新し、必要な場合にのみ非公開でエスカレーションします。
調査エージェントは、市場、政策、競合他社、社内メモを監視し、重要な変更があった場合にフォローアップします。
QA エージェントはテストのバックログを処理し、証拠を収集し、レビューに十分なコンテキストを備えた障害を引き渡します。
運用エージェントは、危険なアクションを実行する前に、アラートを監視し、ランブックを準備し、承認を求めます。
よくあるパターンは、「この質問に答えてください」ではありません。それは、「この席を保持し、利用可能なコンテキストを使用し、結果によって判断される」というものです。
Ankole は、長時間実行される AI 作業のためのアクター指向のランタイムです。各アクティブ セッションは、アドレス指定可能な仮想アクターです。エージェントが単なる HTTP リクエストまたはキュー ジョブであるかのように振る舞うことなく、起動、メッセージの受信、チェックポイント、ストリームの進行状況、休止状態、回復、および継続を行うことができます。
ランタイムは、次の 5 つの技術的賭けを中心に構築されています。
AI 作業用のバーチャル アクター。セッションは、アドレス、メールボックス、ライフサイクル、および回復パスを備えたステートフルな作業 ID であり、緩やかなバックグラウンド作業ではありません。
障害ドメインとしての OTP 監視ツリー。 1 つのエージェントがハング、タイムアウト、またはクラッシュした場合、Ankole はデプロイメント全体の障害に発展させることなく、そのブランチを分離または再起動できます。
ライブ制御用の ZeroMQ Activation Fabric。ウェイクアップ、ステアリング、チェックポイント、ストリーミング、バックプレッシャーは低緯度を通過します。

エージェントがまだ動作している間に、ency ルーティング層を削除します。
実行基盤としてのエージェントコンピュータ。 LLM ループ、ツール、ファイル、ターミナル状態、ストリーミング出力は、ワークスペースに近い Bun + TypeScript コンピューター内で実行されます。
回復と監査のための耐久性のある台帳。メールボックス、ターン、リマインダー、決定、コミットされた副作用はプロセスよりも長く存続します。ストリーミングは進歩です。献身的な仕事は真実です。
ユーザーとオペレーターにとっての約束はシンプルです。エージェントは数時間または数日間作業し、実行中に新しい入力を受け取り、独立して失敗し、状況に応じて回復し、副作用に責任を負い続けることができます。ランタイム引数のより長いバージョンは、「OTP がマルチエージェント オーケストレーションの優れたランタイムである理由」にあります。
これが技術的な賭けです。長期にわたる作業 ID にはアクター モデル、障害セマンティクスには OTP、ライブ アクティベーションには ZeroMQ、ローカル実行にはエージェント コンピューターが必要です。 Ankole は、チャットボット バックエンドというよりも、AI 作業用の分散オペレーティング システムに近いものです。
フローチャート LR
プロバイダー["チャット / Webhook / スケジュール"] --> SG["シグナルゲートウェイ"]
コンソール["Web UI / オペレーター API"] --> CP["コントロール プレーン<br/>フェニックス / OTP"]
SG --> CP
CP --> PG[("PostgreSQL<br/>永続的な真実")]
CP <-->|"RuntimeFabric<br/>ライブ ルーティング"| Worker["エージェント コンピュータ<br/>ブン / TypeScript ワーカー"]
ワーカー --> ツール[「ツール<br/>ブラウザ / ターミナル / ファイル / モデル呼び出し」]
CP --> カーネル["Rust カーネル<br/>AuthZ / ランタイム プリミティブ"]
読み込み中
高いレベルで:
SignalsGateway はプロバイダーのイングレスを受け入れ、それを永続的なアクター イベントに正規化します。
コントロール プレーンは、永続的な状態、アクター オーケストレーション、構成、ID、および承認を所有します。
RuntimeFabric はアクター、ワーカー、RPC レーンを接続して ZeroMQ 経由でライブ実行を行いますが、PostgreSQL は引き続きリプレイ、フェンス、リコンシリエーション、および最終コミットの永続的なソースとして機能します。
エージェントコンピュータexe

孤立したワーカー コンテナ内のキュートなターンとツール。
PostgreSQL は、受け入れられたイベント、状態、フェンス、最終コミットの永続的な記録を保持します。
Ankole は初期のエンジニアリング ディストリビューションであり、洗練されたエンドユーザー製品やホスト型 SaaS ではありません。以下のサブシステムは、現在このリポジトリに動作するコードとして存在しています。正直な注意点は、Vaporware ではなく、洗練さと API の安定性です。
このリポジトリは、アクティブな Ankole コントロール プレーンおよびランタイム ワークスペースです。これはまだエンジニアリングディストリビューションであり、洗練されたエンドユーザーリリースではありません。
app/control_plane - Principal/AuthZ、AppConfigure、セットアップ、コンソール、プラグイン レジストリ、I18n、SignalsGateway、アクター ランタイム、RuntimeFabric、および PostgreSQL 所有の永続状態用の Phoenix/OTP コントロール プレーン。
app/kernel - 暗号化、識別子、電話/JWT ヘルパー、AuthZ 評価、protobuf エンベロープ、および ZeroMQ RuntimeFabric トランスポート用に Elixir と Bun によってロードされた共有 Rust 基盤。
app/agent_computer - ローカル LLM ループ、プロバイダー アダプター、ツール、スキル読み込み、ファイル、ターミナル状態、ワーカー デーモン用の Bun + TypeScript エージェント コンピューター ワーカー。
app/webapps - Phoenix 静的シェルに組み込まれた、認証、セットアップ、コンソール サーフェス用の Vite + React フロントエンド アプリケーション。
app/library - 組み込みのエージェント スキルとスターター テンプレート ( MISSION.md や SOUL.md など)。
app/locales - コントロール プレーンとブラウザ サーフェスによって使用される共有 TOML 翻訳カタログ。
libs/uikit - Ankole Web アプリの共有 UI プリミティブ。
libs/feishu_openapi - ローカル Lark/Feishu OpenAPI クライアント ライブラリ。
libs/slack_openapi - ローカル Slack Web API、ソケット モード、および OIDC クライアント ライブラリ。
Internals/plugins - リポジトリに保持されますが、パブリック プラグイン境界としては提示されない、プライベートのファーストパーティ プロバイダー/プラグイン コード。
tools/devkit - ローカル サービス、アプリ データベース ヘルパー、C 用のワークスペース自動化

ode の生成と分析。
docs/design-docs - プリンシパル ID、認可、構成、I18n、プラグイン、RuntimeFabric、SignalsGateway、およびプロバイダー アダプターに関する現在の設計ドキュメント。
RuntimeFabric は、コントロール プレーンからワーカーへのライブ ファブリックです。アクター トラフィック、制限付き RPC、ワーカー ファイル フレームは ZeroMQ 経由で伝送されますが、PostgreSQL は引き続き永続的なリプレイ、フェンス、リコンシリエーション、および最終コミットのソースとなります。 SignalsGateway はプロバイダー入力層です。外部チャット、Webhook、およびプロバイダー イベントは、ソース ファクトを実行状態に変えることなくアクター イベントになります。
Ankole のデフォルトは、ワークスペース スクリプトの場合は Bun、コントロール プレーンの場合は Elixir/Phoenix です。
バンインストール
# ローカル サポート サービスとワークスペース ヘルパー
パンキット --ヘルプ
bun サービス:開始
パン サービス:ステータス
# コントロールプレーン
bun コントロール プレーン:セットアップ
bun コントロール プレーン:dev
bun コントロール プレーン:テスト
# エージェント コンピューターのコンテナー イメージとテスト
docker build -f app/agent_computer/Dockerfile -t ankole-agent-computer:0.1.0 。
bun エージェント コンピュータ:テスト
bun エージェント-コンピュータ:タイプチェック
# その他のパンパッケージ
bun Web アプリ:ビルド
bun feishu-openapi:test
エージェント コンピューターは、Linux コンテナー ランタイムとして実行されるように設計されています。強いために
bubblewrap コマンドの分離、Doc の実行

[切り捨てられた]

## Original Extract

Ankole is an open-source, self-hosted AgentOS for shared AI colleagues. Claude Tag open source alternative. - AgentBull/ankole

GitHub - AgentBull/ankole: Ankole is an open-source, self-hosted AgentOS for shared AI colleagues. Claude Tag open source alternative. · GitHub
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
AgentBull
/
ankole
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
77 Commits 77 Commits .agents/ skills .agents/ skills .devcontainer .devcontainer .github/ workflows .github/ workflows app app docs docs libs libs plugins plugins tools tools .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .oxfmtrc.json .oxfmtrc.json .oxlintrc.json .oxlintrc.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTEXT.md CONTEXT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.ja.md README.ja.md README.md README.md README.zh-Hans.md README.zh-Hans.md SECURITY.md SECURITY.md ankole.code-workspace ankole.code-workspace bun.lock bun.lock bunfig.toml bunfig.toml knip.config.ts knip.config.ts konsistent.json konsistent.json package.json package.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json turbo.json turbo.json View all files Repository files navigation
Ankole - Open AgentOS for Shared AI Colleagues
How it's different · Product shape · Actor runtime · Architecture · Current status · Development
Ankole is a self-hosted AgentOS for shared AI colleagues. One installation, many agents, real responsibilities — on infrastructure you control.
It moves AI work out of a private chat box and into the places where work already happens: channels, repositories, schedules, dashboards, internal systems, and long-running project context. An Ankole agent has its own identity, memory, permissions, tools, workspace, and responsibility boundary — so it can own ongoing work , not just answer a one-off message.
Claude Tag is a useful public reference point: tag an AI into a Slack thread, let it read the shared context, use organization tools, remember channel context, and follow up when work takes time. Ankole targets the broader open version of that pattern: not only Slack, not only Claude, not only one agent, and not vendor-owned context.
Ankole is for work that needs an owner, not just an answer. A good Ankole role has a visible result: code merged, a report shipped, a customer issue handled, an alert triaged, a market change noticed, or a backlog worked down.
Shared by default, not private chat. Agents join team-visible channels and provider contexts; multiple humans can observe, steer, and continue the same work.
Durable identity, not a prompt convention. Humans and agents are Principals with permission grants and audit trails, so authorization is a runtime concern.
Long-running actor sessions, not request/response. Sessions wake, receive signals, checkpoint, stream progress, hibernate, and recover with context.
Operator-owned context, not vendor-hosted. Memory, configuration, credentials, and audit live in your infrastructure on a self-hosted installation.
Live control plus durable truth, not one or the other. ZeroMQ RuntimeFabric carries live actor/worker/RPC traffic while PostgreSQL remains the source of replay, fences, and final commits.
Many sources. IM, webhooks, scheduled reminders, internal systems, and future provider adapters all become normalized signal input.
Many agents. One installation can host multiple agents with different missions, access, tools, memory, and outbound identities.
Session actors. The long-running execution unit is actor_id = {agent_id, session_id} . A session is where context, workspace state, steering, cancellation, and recovery meet.
Owned context. Conversations, model turns, summaries, signal projections, decisions, corrections, and future domain records live in your infrastructure.
Operator control. Access, configuration, plugin activation, actor leases, outbox side effects, and audit surfaces belong to the installation operator.
Ankole should make these workflows natural:
A coding agent watches an issue, reproduces the bug, changes code, opens a draft PR, and reports what still needs a human decision.
A customer-success agent observes a shared group chat, records the important facts, updates work state, and escalates privately only when needed.
A research agent monitors markets, policy, competitors, and internal notes, then follows up when a change matters.
A QA agent works through a test backlog, gathers evidence, and hands off failures with enough context for review.
An operations agent watches alerts, prepares a runbook, and asks for approval before taking risky action.
The common pattern is not "answer this question." It is "hold this seat, use the available context, and be judged by the result."
Ankole is an actor-oriented runtime for long-running AI work. Each active session is an addressable virtual actor: it can wake, receive messages, checkpoint, stream progress, hibernate, recover, and continue without pretending an agent is just an HTTP request or a queue job.
The runtime is built around five technical bets:
Virtual Actors for AI work. A session is a stateful work identity with an address, mailbox, lifecycle, and recovery path, not loose background work.
OTP Supervision Trees as failure domains. If one agent hangs, times out, or crashes, Ankole can isolate or restart that branch without turning it into a deployment-wide failure.
ZeroMQ Activation Fabric for live control. Wakeups, steering, checkpoints, streaming, and backpressure move through a low-latency routing layer while the agent is still working.
Agent Computer as the execution substrate. The LLM loop, tools, files, terminal state, and streaming output run inside a Bun + TypeScript computer close to the workspace.
Durable Ledger for recovery and audit. Mailboxes, turns, reminders, decisions, and committed side effects outlive processes. Streaming is progress; committed work is truth.
For users and operators, the promise is simple: agents can work for hours or days, receive new input while running, fail independently, recover with context, and keep their side effects accountable. A longer version of the runtime argument is in Why OTP Is a Better Runtime for Multi-Agent Orchestration .
That is the technical bet: actor model for long-lived work identity, OTP for failure semantics, ZeroMQ for live activation, and Agent Computer for local execution. Ankole is closer to a distributed operating system for AI work than a chatbot backend.
flowchart LR
Providers["Chats / webhooks / schedules"] --> SG["SignalsGateway"]
Console["Web UI / operator APIs"] --> CP["Control Plane<br/>Phoenix / OTP"]
SG --> CP
CP --> PG[("PostgreSQL<br/>durable truth")]
CP <-->|"RuntimeFabric<br/>live routing"| Worker["Agent Computer<br/>Bun / TypeScript worker"]
Worker --> Tools["Tools<br/>browser / terminal / files / model calls"]
CP --> Kernel["Rust Kernel<br/>AuthZ / runtime primitives"]
Loading
At a high level:
SignalsGateway accepts provider ingress and normalizes it into durable actor events.
Control Plane owns durable state, actor orchestration, configuration, identity, and authorization.
RuntimeFabric connects actors, workers, and RPC lanes for live execution over ZeroMQ while PostgreSQL remains the durable source of replay, fences, reconciliation, and final commits.
Agent Computer executes turns and tools in an isolated worker container.
PostgreSQL remains the durable record for accepted events, state, fences, and final commits.
Ankole is an early engineering distribution, not a polished end-user product or hosted SaaS. The subsystems below exist as working code in this repository today — the honest caveat is polish and API stability, not vaporware.
This repository is the active Ankole control-plane and runtime workspace. It is still an engineering distribution, not a polished end-user release.
app/control_plane - Phoenix/OTP control plane for Principal/AuthZ, AppConfigure, setup, console, plugin registry, I18n, SignalsGateway, actor runtime, RuntimeFabric, and PostgreSQL-owned durable state.
app/kernel - shared Rust foundation loaded by Elixir and Bun for crypto, identifiers, phone/JWT helpers, AuthZ evaluation, protobuf envelopes, and ZeroMQ RuntimeFabric transport.
app/agent_computer - Bun + TypeScript Agent Computer worker for the local LLM loop, provider adapters, tools, skill loading, files, terminal state, and worker daemon.
app/webapps - Vite + React frontend applications for auth, setup, and console surfaces, built into the Phoenix static shell.
app/library - built-in agent skills and starter templates such as MISSION.md and SOUL.md .
app/locales - shared TOML translation catalogs consumed by the control plane and browser surfaces.
libs/uikit - shared UI primitives for Ankole webapps.
libs/feishu_openapi - local Lark/Feishu OpenAPI client library.
libs/slack_openapi - local Slack Web API, Socket Mode, and OIDC client library.
internals/plugins - private first-party provider/plugin code that is kept with the repo but not presented as the public plugin boundary.
tools/devkit - workspace automation for local services, app database helpers, code generation, and analysis.
docs/design-docs - current design documents for principal identity, authorization, configuration, I18n, plugins, RuntimeFabric, SignalsGateway, and provider adapters.
RuntimeFabric is the live control-plane-to-worker fabric. It carries actor traffic, bounded RPC, and worker-file frames over ZeroMQ while PostgreSQL remains the source of durable replay, fences, reconciliation, and final commits. SignalsGateway is the provider-ingress layer: external chats, webhooks, and provider events become actor events without turning source facts into execution state.
Ankole defaults to Bun for workspace scripts and Elixir/Phoenix for the control plane.
bun install
# Local support services and workspace helpers
bun kit --help
bun services:start
bun services:status
# Control plane
bun control-plane:setup
bun control-plane:dev
bun control-plane:test
# Agent Computer container image and tests
docker build -f app/agent_computer/Dockerfile -t ankole-agent-computer:0.1.0 .
bun agent-computer:test
bun agent-computer:type-check
# Other Bun packages
bun webapps:build
bun feishu-openapi:test
Agent Computer is designed to run as a Linux container runtime. For strong
bubblewrap command isolation, run Doc

[truncated]
