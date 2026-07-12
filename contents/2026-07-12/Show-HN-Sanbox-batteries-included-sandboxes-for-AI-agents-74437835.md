---
source: "https://sanbox.cloud"
hn_url: "https://news.ycombinator.com/item?id=48879908"
title: "Show HN: Sanbox, batteries included sandboxes for AI agents"
article_title: "Sanbox: Sandboxes for AI agents"
author: "oryx1729"
captured_at: "2026-07-12T11:05:48Z"
capture_tool: "hn-digest"
hn_id: 48879908
score: 2
comments: 0
posted_at: "2026-07-12T10:16:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sanbox, batteries included sandboxes for AI agents

- HN: [48879908](https://news.ycombinator.com/item?id=48879908)
- Source: [sanbox.cloud](https://sanbox.cloud)
- Score: 2
- Comments: 0
- Posted: 2026-07-12T10:16:21Z

## Translation

タイトル: 表示 HN: サンボックス、バッテリーを含む AI エージェント用サンドボックス
記事のタイトル: Sanbox: AI エージェント用のサンドボックス
説明: MicroVM 分離、永続的なファイルシステム、ライブ実行イベント、セルフホスティング オプションを備えた再開可能なエージェント サンドボックス。
HN テキスト: こんにちは、HN。私たちは、分離された再開可能なサンドボックスで AI エージェントを実行するためのプラットフォームである Sanbox を構築しています。 OpenCode SDK をハーネスとして使用し、再利用可能なテンプレートをサポートし、Codex、Claude Code、カーソル、CI、または端末で動作する CLI を備えています。各サンドボックスには、MicroVM の分離、永続的なファイルシステム、および実行イベントのライブ証跡があります。セキュリティ/コンプライアンスのために必要な場合は自己ホストすることもできます。ネットワーク ACL、シークレット管理、LLM コストの追跡と可観測性を追加するロードマップにあります。特定のユースケース向けにカスタム統合やトリガーを喜んで構築します。たとえば、データ変更時にデータベース エッジ機能からサンドボックスをスピンアップしたり、複雑なワークフロー処理のために Sanbox を電子メールの受信箱に接続したりできます。まだ時期尚早です。ここからサインアップして試してみてください: https://sanbox.cloud 私のメールアドレスもここの私のプロフィールにあります。

記事本文:
コンテンツにスキップ
サンボックス
サンラボ
AI エージェントを実行する
孤立したサンドボックス。
CLI から並列実行を起動します。それぞれの実行は、永続的なファイル システムを備えた分離された microVM 内で行われます。サンボックス
アーティファクト、エージェントの状態、会話履歴を保存しながら、ツールの呼び出しとモデルのアクティビティを記録します。
そのため、後で任意のサンドボックスを再開できます。
サンボックス
コントロールプレーン
F 財務 GPT 5.6
18 件のレポートをチェックしました 03:12
C 契約リスク GLM-OCR
12 個の条項にフラグが立てられました 02:47
M マーケット スキャン キミ 2.7
36 個のソースを読み取りました 04:08
S セキュリティ体制 GLM 5.2
8 つのコントロールをチェックしました 01:54
目標: 買収を評価します。並行して実行します。再開可能な MicroVM。
NPM を使用してインストールする
npm install -g @sanlabs/sanbox-cli
01 / ラン境界
プロンプトは
セキュリティ境界線。
エージェントには、ファイル、ツール、認証情報、状態、およびネットワーク アクセスが必要です。 Sanbox はこれらの入力を
リソースとネットワーク制限が構成された専用の MicroVM。
ファイルシステムは、アーティファクト、エージェントの状態、会話履歴を含むスナップショットを作成されます。サンドボックスを停止し、
実行記録を検査し、後で同じ状態から再開します。
CLI から並列実行を開始して監視します。
ターミナル、CI ジョブ、Codex、または Claude Code から作業をディスパッチします。運転を停止することなく、取り外しと再取り付けが可能です。
エージェントネイティブのオンボーディング sanbox init は、エージェントをコーディングするための適切な命令を追加します。
ポータブルな実行バンドル どのタスクとファイルがマシンから流出するかを正確に検査します。
再開可能なワークスペース エージェントの状態と会話履歴をそのままにして、保存されたファイルシステムのスナップショットから再開します。
再開可能なMicroVM
実行ごとに。
ランナー、モデル、CPU、メモリ、タイムアウト、およびネットワーク ポリシーを選択します。 Sanbox はファイルシステムのスナップショットを作成します。
アーティファクト、エージェントの状態、会話履歴を保存できるため、停止した実行を中断したところから再開できます。
00:02 harness.input.ready 入力ドシエが抽出されました
00:05 Agent.turn.started エージェントのターンが開始されました
00

:11 Agent.tool.completed rg "refreshToken" src/
00:16 Agent.file.changed src/auth/session.ts を変更しました
00:24 エージェント.使用法.更新 1,834 入力 · 412 出力
00:31 snapshot.saved ファイルシステム + エージェント状態が保持されました
コーディング エージェント、ターミナル、または CI ジョブからの並行作業を委任します。各実行に必要なファイルのみをパッケージ化します。
現在、OpenCode ランナーでサポートされているモデルからお選びいただけます。より多くのエージェントとエンドポイントに対応する安定したアダプター契約が付いています。
反復可能なエージェント作業のために、承認された OpenCode イメージ、モデル、制限、および Web ツール ポリシーを登録します。
専用のゲスト カーネルと明示的な CPU およびメモリ制限を使用して、各サンドボックスを独自の MicroVM で実行します。
明示的な宛先許可、プライベート IP ブロック、および実行に関連付けられた決定によるデフォルト拒否の出力。
有効期間の長い認証情報をサンドボックスの外に保管します。実行に必要な機能と寿命のみを発行します。
再開可能なファイルシステムのスナップショット
ファイル、アーティファクト、エージェントの状態、会話履歴を保持し、後で同じサンドボックスを再開します。
クラウドまたはハードウェアに導入します。
コントロール プレーン、MicroVM ランナー、ファイル システムのスナップショット、認証情報、実行履歴を、チームが運用するインフラストラクチャ内に保持します。
EU インフラストラクチャ上での専用導入をお客様と一緒に運用します。
EU 内での専用 Sanbox の展開。
コントロール プレーン、MicroVM ランナー、および永続的なファイル システム スナップショットは、ドイツの専用インフラストラクチャ上で実行されます。
入力バンドル、ランナー設定、イベント ストリーム、ファイル システム スナップショット、および出力を同じ実行 ID に保ちます。
CLI から並列実行を開始します。
ターミナル、CI、またはコーディング エージェントから 1 つのタスクまたは制限された並列タスク リストを起動します。各実行に必要なファイルのみを含めます。
それぞれの実行に専用のゲスト カーネル、明示的なリソース制限、および独自の永続ファイル システムを与えます。
好みのモデルで OpenCode を使用し、再利用可能にします

ツール、権限、リソースのデフォルトのテンプレート。
実行ごとに CPU、メモリ、タイムアウト、ワークスペース保持、および承認されたランナーのデフォルトを設定します。
エージェントのターン、ツールの呼び出し、ファイルの変更、モデルの使用状況、ログ、エラーが発生したときにストリーミングします。
アーティファクト、エージェントの状態、会話履歴を含むファイルシステムを復元し、同じサンドボックスを続行します。
Sanbox をどこで使用するかを教えてください。アカウントを設定し、24 時間以内にログイン手順を送信します。
これらの詳細は、アカウントを設定し、ログイン手順を送信するためにのみ使用されます。
アカウントを設定し、24 時間以内にログイン手順をメールで送信します。

## Original Extract

Resumable agent sandboxes with MicroVM isolation, persistent filesystems, live run events, and a self-hosting option.

Hi HN, We are building Sanbox, a platform for running AI agents in isolated and resumable sandboxes. We use the OpenCode SDK as the harness, support reusable templates, and have a CLI that works with Codex, Claude Code, Cursor, CI, or your terminal. Each sandbox has MicroVM isolation, a persistent filesystem, and live trail of run events. Can also self-host if required for security/compliance. It's on the roadmap to add network ACL, secrets managements, LLM cost tracking & observability. We are also happy to build custom integrations or triggers for specific use cases. For example, spinning up a sandbox from a database edge function when data changes, or connecting Sanbox to an email inbox for complex workflow processing. It is still early, you can sign up to try here: https://sanbox.cloud My email is also on my profile here.

Skip to content
San box
by Sanlabs
Run AI agents in
isolated sandboxes.
Launch parallel runs from the CLI, each in an isolated microVM with a persistent filesystem. Sanbox
records tool calls and model activity while preserving artifacts, agent state, and conversation history,
so you can resume any sandbox later.
SANBOX
CONTROL PLANE
F Financials GPT 5.6
18 reports checked 03:12
C Contract risks GLM-OCR
12 clauses flagged 02:47
M Market scan Kimi 2.7
36 sources read 04:08
S Security posture GLM 5.2
8 controls checked 01:54
Goal: evaluate an acquisition. Parallel runs. Resumable MicroVMs.
INSTALL WITH NPM
npm install -g @sanlabs/sanbox-cli
01 / Run boundary
A prompt is not a
security boundary.
Agents need files, tools, credentials, state, and network access. Sanbox places those inputs in a
dedicated MicroVM with configured resource and network limits.
The filesystem is snapshotted with artifacts, agent state, and conversation history. Stop the sandbox,
inspect the run record, and resume it from the same state later.
Start and watch parallel runs from the CLI.
Dispatch work from a terminal, CI job, Codex, or Claude Code. Detach and reattach without stopping the run.
Agent-native onboarding sanbox init adds the right instructions for coding agents.
Portable run bundles Inspect exactly which task and files leave your machine.
Resumable workspaces Resume from a saved filesystem snapshot with agent state and conversation history intact.
A resumable MicroVM
for every run.
Choose the runner, model, CPU, memory, timeout, and network policy. Sanbox snapshots the filesystem with
artifacts, agent state, and conversation history so a stopped run can resume where it left off.
00:02 harness.input.ready Input dossier extracted
00:05 agent.turn.started Agent turn started
00:11 agent.tool.completed rg "refreshToken" src/
00:16 agent.file.changed Modified src/auth/session.ts
00:24 agent.usage.updated 1,834 input · 412 output
00:31 snapshot.saved Filesystem + agent state persisted
Delegate parallel work from any coding agent, terminal, or CI job. Package only the files each run needs.
Choose from supported models in the OpenCode runner today, with a stable adapter contract for more agents and endpoints.
Register an approved OpenCode image, model, limits, and web-tool policy for repeatable agent work.
Run each sandbox in its own MicroVM with a dedicated guest kernel and explicit CPU and memory limits.
Default-deny egress with explicit destination grants, private-IP blocking, and decisions tied back to a run.
Keep long-lived credentials outside the sandbox. Issue only the capability and lifetime a run requires.
Resumable filesystem snapshots
Persist files, artifacts, agent state, and conversation history, then resume the same sandbox later.
Deploy in your cloud or on your hardware.
Keep the control plane, MicroVM runners, filesystem snapshots, credentials, and run history inside infrastructure your team operates.
Dedicated deployment on EU infrastructure, operated with you.
Dedicated Sanbox deployment in the EU.
The control plane, MicroVM runners, and persistent filesystem snapshots run on dedicated infrastructure in Germany.
Keep the input bundle, runner settings, event stream, filesystem snapshot, and outputs under the same run ID.
Start parallel runs from the CLI.
Launch one task or a bounded parallel task list from your terminal, CI, or coding agent. Include only the files each run needs.
Give each run a dedicated guest kernel, explicit resource limits, and its own persistent filesystem.
Use OpenCode with your preferred model and reusable templates for tools, permissions, and resource defaults.
Set CPU, memory, timeout, workspace retention, and approved runner defaults for each run.
Stream agent turns, tool calls, file changes, model usage, logs, and errors as they happen.
Restore the filesystem with artifacts, agent state, and conversation history, then continue the same sandbox.
Tell us where you’ll use Sanbox. We’ll set up your account and send login instructions within 24 hours.
We’ll use these details only to set up your account and send login instructions.
We’ll set up your account and email login instructions to within 24 hours.
