---
source: "https://redqueen.sh/"
hn_url: "https://news.ycombinator.com/item?id=48462957"
title: "Red Queen – deterministic orchestration for AI coding agents"
article_title: "Red Queen — The Jenkins for Your AI Coding Workflow"
author: "odyth"
captured_at: "2026-06-09T18:53:41Z"
capture_tool: "hn-digest"
hn_id: 48462957
score: 1
comments: 0
posted_at: "2026-06-09T16:08:43Z"
tags:
  - hacker-news
  - translated
---

# Red Queen – deterministic orchestration for AI coding agents

- HN: [48462957](https://news.ycombinator.com/item?id=48462957)
- Source: [redqueen.sh](https://redqueen.sh/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T16:08:43Z

## Translation

タイトル: Red Queen – AI コーディング エージェントの決定論的オーケストレーション
記事のタイトル: Red Queen — AI コーディング ワークフローのための Jenkins
説明: クロード・コードにとってのレッド・クイーンは、ジェンキンスがバッシングすることと同じである。設定可能なヒューマン ゲートを使用して、仕様→計画レビュー→コード→レビュー→テストを通じてチケットを往復させる、確定的なゼロトークン ステート マシン。オープンソース。自己ホスト型。 Node.js。

記事本文:
AI コーディング ワークフローのための Jenkins。
「クロード・コードにとってのレッド・クイーンは、ジェンキンスがバッシングすることと同じだ。」
決定論的なステートマシンがチケットを往復させます
仕様→計画レビュー→コード→レビュー→テスト→人によるレビュー。
フェーズ、スキル、ゲートは YAML で構成します。オーケストレーター
それ自体は AI トークンを消費せず、ワーカー間で作業をやり取りするだけです
そして設定したゲートで止まります。ビルドサーバーのように決定的。
コーディング エージェントはすでに存在します。あなたにないものは決定論的なパイプラインです
仕様、計画レビュー、コード、レビュー、テスト、ヒューマンゲートを通じてチケットを往復させるため
あなたが見ていない間に。 Red Queen はそのパイプラインです。 YAML で構成可能、監査可能、
また、ルーティングにトークンを消費しないステート マシン上に構築されています。
コーディネートではなく、コードに対して料金を支払います
ステート マシンは純粋な決定論的ロジックであり、ルーティング層に LLM はありません。費やしたすべてのトークンは、次に何をするかを決定するのではなく、実際のコード作業に使用されます。
各フェーズ (仕様作成者、コーダー、レビュー担当者、テスター、コメントハンドラー) は、専用のプロンプトを個別に実行します。キッチンシンクを集中的にビートします。
レビュー チェックポイントは第一級のワークフロー構成であり、後付けではありません。 redqueen.yaml で追加、削除、並べ替えを行います。
Jira および GitHub の問題との双方向同期。作業はトラッカーからパイプラインを通って流れ、戻ってきます。ロードマップ上は直線的。
失敗したフェーズは最大 3 回再試行され、その後は人間のゲートにエスカレーションされます。無限ループはありません。トークンの請求書が暴走することはありません。
すぐに使用できるポーリング調整機能を備えた、即時応答のためのオプションの Webhook サーバー。外部キューは必要ありません。
スキルはマークダウン プロンプト テンプレートです。 SKILL.md を .redqueen/skills/<name>/ にドロップすると、あなたのバージョンがビルトインよりも優先されます。
すべての統合は、IssueTracker または SourceControl インテルを実装します。

表面。新しいトラッカー = 新しいアダプターであり、コアの変更ではありません。
バックグラウンド オーケストレーターは自動化されたフェーズを通じてチケットを処理し続け、キーボードに戻ったときのために人間のゲートで作業を保留します。
あなたはチケットを提出します。クロードは仕様書を書きます。自動化された計画レビューのスコア
それはブロッカーと曖昧さのためです。人間の門で承認します。クロード
コードを記述して PR を開きます。別のクロードがそれをレビューします。別のテスト
それ。最終的な PR を確認してマージします。どの門からでも飛び込むか、あるいは門を破ってください。
完全にゲートアウト。
フェーズは動的であり、コードには組み込まれず、 redqueen.yaml で定義されます。
セキュリティ審査ゲートを追加します。人間による仕様レビューを削除して、
plan-review は、 SkipSpecReviewIfReady を使用してクリーンな仕様を自動プロモートします。
グラフはあなたが形作ります。
Red Queen はコーディング エージェントの上に位置します。コーディング エージェントを置き換えたり、ホストしたりするのではなく、
またはそのスキルを管理します。これらのいずれかを見たことがあれば、レッドクイーンは
カテゴリーの違うもの。
エージェントホストとオープンコードではない
Red Queen はクロード CLI サブプロセスを生成し、それらをフェーズを通して駆動します。これはエージェント ランタイムではなくディスパッチャーです。
スキルマネージャーではなく、agentskills.io
私たちの「スキル」は、パイプラインのフェーズ プロンプト テンプレートです。これらは、agentskills.io 形式に従っているため移植可能ですが、Red Queen は一般的なユーザー スキルのためのレジストリやローダーではありません。
仕様フレームワークと OpenSpec ではない
永続的な PRD やセッション間の計画レイヤーはありません。 Red Queen はパイプラインを通じてチケットを実行します。これは仕様フレームワークを補完するものであり、置き換えるものではありません。
インタラクティブ コーディング ツールではなく、Claude Code ではありません、エイダー
Red Queen はバックグラウンドで実行され、ユーザーが見ていない間にチケットをフェーズごとに移動します。キーボードでのペアプログラミングには、Claude Code または Aider を使用してください。
ゼロから最初の AI 主導のチケットまで、約 15 分で作成されます。
Claude Code をインストールし、 clude --version を実行します。レッドクイーンディ

クロード コード ワーカーをパッチするため、CLI にログインする必要があります。
Red Queen をインストールします。
npm install -g redqueen
ターゲット リポジトリに対するコンテンツ、問題、プル リクエスト、ワークフロー、メタデータ権限を含む GitHub PAT を github.com/settings/personal-access-tokens/new で生成します。
リポジトリ内で初期化します。
レッドクイーン init -y
トークンを .env に追加します。
GITHUB_PAT=ghp_xxxxxxxxxxxxxxxxxxxxx
オーケストレーターを開始します。
レッドクイーンスタート
http://127.0.0.1:4400 — ライブ ダッシュボードを開きます。
GitHub でテストの問題を開き、rq:phase:spec-writing ラベルを適用します。位相の動きを観察してください。
完全なドキュメント、アダプター ガイド、構成リファレンスは、
GitHub の README 。
Red Queen は以下のツールに代わるものではありません。ツールに委任するのはマネージャーです。
今日はクロード コードを発送します。アダプター パターンは、他のコーディング CLI でのスワップは、書き換えではなく構成変更であることを意味します。
採用する前に尋ねる質問。
AI コーディング エージェントのための決定論的なオーケストレーター。孤立した AI ワーカーを派遣するステートマシンです
仕様の作成、コーディング、レビュー、テスト、人間によるレビューという完全なソフトウェア開発ライフサイクルを通じて。
オーケストレーターはルーティングに AI トークンをまったく使用しません。すべての決定は、LLM ではなく純粋な決定論的ロジックによって行われます。
Red Queen は、Jenkins がバッシングするのと同じクロード コードを行うことになります。コーディング エージェントの上に座って、チケットを送り届けます。
決定論的なステートマシン上のフェーズを介して。 Devin、Cline、Aider はあなたと対話するコーディング エージェントです
直接的に。 Red Queen は、構成可能なヒューマン ゲートを備えた、バックグラウンドでそれらを送信するパイプラインです。
フェーズの間。ルーティング層には LLM がないため、ルーティングにはトークンのコストがゼロで、完全にデバッグ可能です。
いいえ、Red Queen はクロード CLI サブプロセスを生成し、フェーズ グラフを通じてそれらを駆動します。
エージェント ランタイムや汎用スキルではなく、ディスパッチャー

l レジストリ。私たちの「スキル」はフェーズプロンプトです
パイプラインのテンプレート。彼らはagentskills.ioをフォローしています
フロントマター形式なので、互換性のあるクライアントに移植できますが、Red Queen 自体は読み込まれません。
ユーザーがインストールしたスキルを管理します。
OpenSpec は仕様/PRD フレームワークです - 永続的な計画
セッション間のアーティファクト。 Red Queen は実行パイプラインです。1 つのチケットが入力され、PR が出力されます。それらは補完的であり、
競争力がない。 OpenSpec で作成されたチケットを Red Queen パイプラインに確実にフィードできます。
今日は、そうです。レッドクイーンはクロード・コードの労働者を派遣する。アダプター パターンにより、次のサポートが追加されます。
他のコーディング CLI は簡単であり、それはロードマップにあります。
オーケストレーション自体は無料で、ルーティングに消費されるトークンはゼロです。トークンはクロードのサブスクリプションです
または、API は実際のコード作業に対してすでに料金を請求しています。
はい。 Red Queen は、ラップトップまたは所有するサーバー上でローカル Node.js 24+ プロセスとして実行されます。ホスト型サービスなし
が必要です。コードは、構成した AI CLI にのみ送信されます。 Red Queen 自体は HTTPS 経由で GitHub または Jira と通信します。
はい — project.buildCommand および testCommand のスコープは実行ごと、およびモジュールごとのコマンドです
設定可能です。
フェーズは最大 3 回再試行され、その後人間のゲートにエスカレートします。無限ループや暴走請求はありません。
はい — redqueen.yaml を編集します。ゲートは構成であり、ハードコーディングされません。
はい。 SKILL.md を .redqueen/skills/<name>/ にドロップすると、組み込みよりも優先されます。
GitHub リポジトリのスキル契約を参照してください。
GitHub Issues、Jira、および GitHub (ソース管理) は本日出荷されます。リニアはロードマップに載っています。
v0.1のプレビュー版です。それを使用し、バグを報告し、荒削りなエッジを予期してください。API サーフェスは構築するのに十分な安定性を備えています。
はい — MIT ライセンスを取得しており、公開で開発されています。
github.com/odyth/red-queen 。
LLM を使用したオーケストレーションを停止します。
彼らに命令し始めてください。
オープンソース。自己ホスト型。

決定論的。リポジトリにスターを付け、タイヤを蹴り、バグをファイルします。

## Original Extract

Red Queen is to Claude Code what Jenkins is to bash. A deterministic, zero-token state machine that shuttles tickets through spec → plan review → code → review → test with configurable human gates. Open source. Self-hosted. Node.js.

The Jenkins for your AI coding workflow.
“Red Queen is to Claude Code what Jenkins is to bash.”
A deterministic state machine shuttles tickets through
spec → plan review → code → review → test → human review .
You configure the phases, skills, and gates in YAML. The orchestrator
itself spends zero AI tokens — it just shuttles work between workers
and stops at the gates you set. Deterministic, like a build server.
You already have a coding agent. What you don't have is a deterministic pipeline
to shuttle tickets through it — spec, plan review, code, review, test, human gates —
while you're not looking. Red Queen is that pipeline. Configurable in YAML, auditable,
and built on a state machine that spends zero tokens on routing.
You pay for code, not coordination
The state machine is pure deterministic logic — no LLM in the routing layer. Every token you spend goes to actual code work, not deciding what to do next.
Each phase — spec writer, coder, reviewer, tester, comment-handler — runs a purpose-built prompt in isolation. Focused beats kitchen-sink.
Review checkpoints are first-class workflow configuration, not an afterthought. Add them, remove them, reorder them in redqueen.yaml .
Bidirectional sync with Jira and GitHub Issues. Work flows from your tracker through the pipeline and back. Linear on the roadmap.
Failed phases retry up to three times, then escalate to a human gate. No infinite loops. No runaway token bills.
Optional webhook server for instant response, with a polling reconciler that works out of the box. No external queue required.
Skills are markdown prompt templates. Drop SKILL.md into .redqueen/skills/<name>/ and your version wins over the built-in.
Every integration implements the IssueTracker or SourceControl interface. New trackers = new adapter, never a core change.
Background orchestrator keeps tickets moving through automated phases and parks work at human gates for when you're back at the keyboard.
You file a ticket. Claude writes a spec. An automated plan review scores
it for blockers and ambiguity. You approve at the human gate. Claude
writes the code and opens a PR. Another Claude reviews it. Another tests
it. You review the final PR and merge. Jump in at any gate, or rip the
gates out entirely.
Phases are dynamic — defined in redqueen.yaml , not baked into the code.
Add a security-review gate. Remove the human spec-review and let
plan-review auto-promote clean specs with skipSpecReviewIfReady .
The graph is yours to shape.
Red Queen sits above your coding agent — it doesn't replace it, host it,
or manage its skills. If you've seen one of these, Red Queen is a
different category of thing.
Not an agent host vs. opencode
Red Queen spawns claude CLI subprocesses and drives them through phases. It's the dispatcher, not the agent runtime.
Not a skill manager vs. agentskills.io
Our “skills” are phase prompt templates for the pipeline. They follow the agentskills.io format so they're portable, but Red Queen isn't a registry or loader for generic user skills.
Not a spec framework vs. OpenSpec
No persistent PRDs, no cross-session planning layer. Red Queen executes a ticket through a pipeline — complementary to spec frameworks, not a replacement.
Not an interactive coding tool vs. Claude Code, Aider
Red Queen runs in the background and moves tickets through phases while you're not looking. Use Claude Code or Aider for pair-programming at the keyboard.
From zero to your first AI-driven ticket in about 15 minutes.
Install Claude Code and run claude --version . Red Queen dispatches Claude Code workers, so the CLI must be logged in.
Install Red Queen:
npm install -g redqueen
Generate a GitHub PAT at github.com/settings/personal-access-tokens/new with Contents, Issues, Pull requests, Workflows, Metadata permissions on your target repo.
Initialize inside your repo:
redqueen init -y
Add your token to .env :
GITHUB_PAT=ghp_xxxxxxxxxxxxxxxxxxxx
Start the orchestrator:
redqueen start
Open http://127.0.0.1:4400 — the live dashboard.
Open a test issue in GitHub and apply the rq:phase:spec-writing label. Watch the phases move.
Full docs, adapter guides, and configuration reference live in the
GitHub README .
Red Queen isn't a replacement for the tools below — it's the manager that delegates to them.
It dispatches Claude Code today; the adapter pattern means swapping in other coding CLIs is a config change, not a rewrite.
The questions you'd ask before adopting.
A deterministic orchestrator for AI coding agents. It's a state machine that dispatches isolated AI workers
through a complete software development lifecycle: spec writing, coding, review, testing, and human review.
The orchestrator spends zero AI tokens on routing — all decisions are made by pure deterministic logic, not an LLM.
Red Queen is to Claude Code what Jenkins is to bash — it sits above the coding agent and shuttles tickets
through phases on a deterministic state machine. Devin, Cline, and Aider are coding agents you interact with
directly. Red Queen is the pipeline that dispatches them in the background, with configurable human gates
between phases. No LLM in the routing layer, so routing costs zero tokens and is fully debuggable.
No. Red Queen spawns claude CLI subprocesses and drives them through a phase graph — it's a
dispatcher, not an agent runtime or a generic skill registry. Our “skills” are phase prompt
templates for the pipeline; they follow the agentskills.io
frontmatter format so they're portable to any compatible client, but Red Queen itself does not load or
manage user-installed skills.
OpenSpec is a spec / PRD framework — persistent planning
artifacts across sessions. Red Queen is an execution pipeline — one ticket in, PR out. They're complementary,
not competitive. You can absolutely feed OpenSpec-authored tickets into a Red Queen pipeline.
Today, yes. Red Queen dispatches Claude Code workers. The adapter pattern makes adding support for
other coding CLIs straightforward, and that's on the roadmap.
The orchestration itself is free — zero tokens spent on routing. Tokens are what your Claude subscription
or API already charges for the actual code work.
Yes. Red Queen runs as a local Node.js 24+ process on your laptop or any server you own. No hosted service
is required. Your code is only sent to the AI CLI you configure; Red Queen itself talks to GitHub or Jira over HTTPS.
Yes — project.buildCommand and testCommand scope per-run, and per-module commands
are configurable.
Phases retry up to three times and then escalate to a human gate. No infinite loops, no runaway bills.
Yes — edit redqueen.yaml . Gates are configuration, not hardcoded.
Yes. Drop a SKILL.md at .redqueen/skills/<name>/ and it wins over the built-in.
See the skills contract in the GitHub repository.
GitHub Issues, Jira, and GitHub (source control) ship today. Linear is on the roadmap.
It's a v0.1 preview. Use it, file bugs, expect rough edges — the API surface is stable enough to build on.
Yes — MIT-licensed, developed in the open at
github.com/odyth/red-queen .
Stop orchestrating with LLMs.
Start commanding them.
Open source. Self-hosted. Deterministic. Star the repo, kick the tires, file bugs.
