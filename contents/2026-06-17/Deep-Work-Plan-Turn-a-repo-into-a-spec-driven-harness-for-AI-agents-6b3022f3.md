---
source: "https://deepworkplan.com/"
hn_url: "https://news.ycombinator.com/item?id=48569526"
title: "Deep Work Plan – Turn a repo into a spec-driven harness for AI agents"
article_title: "Deep Work Plan — structured execution for AI coding agents"
author: "xergioalex"
captured_at: "2026-06-17T13:23:23Z"
capture_tool: "hn-digest"
hn_id: 48569526
score: 2
comments: 0
posted_at: "2026-06-17T12:34:13Z"
tags:
  - hacker-news
  - translated
---

# Deep Work Plan – Turn a repo into a spec-driven harness for AI agents

- HN: [48569526](https://news.ycombinator.com/item?id=48569526)
- Source: [deepworkplan.com](https://deepworkplan.com/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T12:34:13Z

## Translation

タイトル: Deep Work Plan – リポジトリを AI エージェント向けの仕様主導のハーネスに変える
記事のタイトル: Deep Work Plan — AI コーディング エージェントの構造化された実行
説明: モデルよりもコンテキストが重要です。 Deep Work Plan は、あらゆるリポジトリを、コーディング エージェントが長期的な作業を完了できる構造化された環境に変えます。

記事本文:
Deep Work Plan — ユーザーが選択したテーマに応じて、以下の AI コーディング エージェントの構造化された実行。 -->
コンテンツにスキップ
Deep Work Plan 方法論 · 仕様 · キット 方法論 仕様キットの例 Trust GitHub JP
オープンな方法論 · MIT · エージェントに依存しない
モデルは重要です。コンテキストの方が重要です。
Deep Work Plan は、あらゆるリポジトリを構造化された環境 (コンテキスト、ガードレール、耐久性のあるプラン) に変えます。そこでは、コーディング エージェントが正確に実行され、長期にわたる作業が完了します。
init.md プロンプトをコピーし、コーディング エージェント (クロード コード、カーソル、コーデックスなど) に貼り付けて、リポジトリを AI ファーストにします。
Deep Work Plan は、リポジトリ自体がハーネスとなる仕様駆動型の開発です。
AI コーディング エージェントは、短期間で非常に効果的です。移行、新しいサブシステム、数十のファイルにわたるリファクタリングなど、長期的な作業では、状況が漂い、コンテキストが埋まり、決定事項が忘れられ、数時間かかるタスクが途中で放棄されます。
Deep Work Plan は仕様駆動型開発で答えます。計画は永続的な信頼できる情報源であり、エージェントは明示的な受け入れ基準と検証ゲートに基づいて実行されます。ドリフトが終了しても、作業は検証可能のままであり、エージェントはセッションをまたいで作業を再開できます。
また、ハーネス工学によりポータブルに作られています。エージェント ハーネスは、コンテキスト、ツール、制御ループ、ガードレール、再開可能な状態といったモデルの周囲の足場であり、信頼性を高めます。 Deep Work Plan は、そのハーネスをリポジトリ自体 (AGENTS.md、ドキュメント、.agents/ スキル ホーム、DWP スキル) にインストールするため、どのエージェントでも任意のリポジトリを試行できます。 Dailybot で生まれ、数か月間実戦テストされ、DailybotHQ/ディープワークプラン スキルとしてリリースされました。
任意のリポジトリを指定します。その理由は、コピー＆ペーストしないことです。
オンボーディング フローはリポジトリの

実際の言語、フレームワーク、パッケージ マネージャー、および検証コマンドを使用して、そのリポジトリに適合したアーティファクトを生成します。汎用スタブは失敗として扱われます。
01 スタックとアーキタイプに関する理由
マニフェスト、フォルダー レイアウト、CI を読み取り、実際のテスト、lint、ビルド コマンドを推測し、リポジトリを個別のリポジトリまたはオーケストレーター ハブとして分類します。
02 AGENTS.md、docs/、およびモジュールごとのドキュメントを生成します
推論された AGENTS.md、分類された docs/ 階層、および各主要モジュール内の README と docs/ - プレースホルダーではなく、リポジトリの実際のコマンドで埋められています。
03 .claude から .agents へのシンボリックリンクを使用して .agents/ を足場にします
クロスエージェントの .agents/ ディレクトリ (スキル、エージェント、コマンド) および .claude から .agents へのシンボリックリンクは、CLAUDE.md を AGENTS.md にミラーリングするため、すべてのツールが 1 つの信頼できる情報源を読み取ります。
04 DWP スキルとスキャフォールドをインストールします。.dwp/
Deep Work Plan スキルを接続し、計画とドラフト用に gitignored .dwp/ フォルダーを作成し、必要に応じて devcontainer サポートなどのオプトイン アドオンを階層化します。
一つの指示。残りはリポジトリが行います。
インストール方法を選択したり、テンプレートをコピーしたりする必要はありません。あなたはエージェントに一行だけ手渡します。スキル (再利用可能なエンジン) をインストールし、リポジトリをそれに適応させます。
deepworkplan.com/init.md にあるオンボーディング プロンプトと、リンク先の方法論、仕様、キット (採用しようとしている標準) を読み取ります。
02 Deep Work Planスキルをインストールします
スキルはエンジンであり、どのリポジトリでも同様です。 1 つのコマンドで、ルーターとそのサブスキル (作成、実行、洗練、再開、ステータス、検証、オンボード、作成者) をクロード コード、カーソル、コーデックス、ジェミニ、およびコパイロットに取り込みます。
実際のスタックについて推論します。決してコピー＆ペーストしないでください。AGENTS.md、分類されたドキュメント/ツリー、モジュールごとの README、理由が書き込まれます。

d .agents/ kit、および gitignored .dwp/。リポジトリがハーネスになります。
あらゆるタスクに対して長期的な詳細作業計画を生成し、明示的な受け入れ基準、検証ゲート、再開可能な状態を使用して、自律的に数時間にわたって段階的に実行します。
スキルはどこにでも同じようにインストールされます。適応されるのはあなたのリポジトリ、つまりスタック用に生成された AGENTS.md、ドキュメント、および推論された .agents/ キットです。この分割により、この方法論は 1 回限りの足場ではなく、再利用可能な標準になっています。
エージェントが自律的に動作するために必要なものすべて。
1 回の実行でアトミックにコミットされます。すべての出力はマークダウンであり、すべての変更は監査可能です。
リポジトリのルートにある AGENTS.md
プレースホルダーを含むテンプレートではなく、リポジトリの実際のスタック、コマンド、構造に基づいて推論されます。 CLAUDE.md は AGENTS.md にシンボリックリンクされています。
分類されたドキュメント/モジュールごとのドキュメント
アーキテクチャ、セットアップ、標準、トラブルシューティングに加えて、コードベースから生成された各主要モジュール内の README とドキュメント。
.agents/ と .claude から .agents へのシンボリックリンク
クロスエージェント .agents/ ディレクトリ (スキル、エージェント、コマンド) と .claude から .agents へのシンボリックリンクにより、すべてのツールが 1 つの信頼できる情報源を読み取ります。
Deep Work Plan スキルがインストールされました
作成、実行、調整、再開、ステータス、検証、オンボード、作成 - リポジトリごとのコピーなしで、エージェントは単一のスキル パックとして利用できます。
/dwp-verify は仕様に対する客観的な合否レポートを生成するため、「AI ファースト」は主張ではなく検証され、計画のたびに再検証可能です。
オンボーディングでは、リポジトリを個別のリポジトリ (一般的なケース)、またはリポジトリ全体で子プランを調整するオーケストレーター ハブとして分類します。
リポジトリが成長するリビング キット
author サブスキル (skill-create、agent-create) により、リポジトリが独自のスキルを進化させます。

、エージェント、およびコマンド。 dependency-upgrade などのオプトイン メンテナンス アドオンは、それ自体を最新の状態に保つのに役立ちます。
デーモンも外部状態もありません。計画とドラフトは gitignored .dwp/ フォルダーに置かれ、コンテキストがオーバーフローした後でも、あらゆるタスクは git のみから再開されます。
すでに使用しているエージェントと連携します。
1 つの方法論で、多数のアダプターを使用できます。 Markdown はフレームワークを何も結合しません。Markdown を読み取るすべてのエージェントはディープ ワーク プランを実行できます。
ネイティブ WebFetch およびスラッシュ コマンドを使用したリファレンス実装。
フルアダプター。 WebFetch がゲートされている場合は、オフライン バンドルを使用します。
オフラインバンドルを推奨。ルールは .codex/ にインストールされます。
完全なアダプター — dwp-* コマンドは、AGENTS.md および # プロシージャを介して実行されます。
ネイティブ WebFetch を備えた Gemini 2.5 Pro 以降が必要です。
オープンソース。 AGENTS.md をネイティブに読み取り、 # コマンド経由で dwp-* を実行します。
ルールと # コマンド プロシージャによって、完全なディープ ワーク プラン ループが駆動されます。
オープンソース。マークダウン ルールと # コマンドは dwp-* ステップごとに実行されます。
ネイティブ コマンド サーフェスを備えた完全なアダプター。
重要なスタックのプリセットを推論します。
これらは推論の補助であり、テンプレートではありません。オンボーディングはリポジトリの実際のマニフェストを読み取り、スタックごとに適応します。プリセットをブラインド コピーすることはありません。 Monorepos はモジュールごとのドキュメントを取得します。
TypeScript · Node Express · Fastify
TypeScript · Lambda サーバーレス · SAM
個々のリポジトリ、またはオーケストレーター ハブ。
アーキタイプのオンボーディングフォーク。ほとんどのリポジトリは個別のリポジトリです。ハブは、多くのリポジトリにわたる子ディープ ワーク プランを調整します。この方法論では、両方をファーストクラスとして扱います。
1 つのプライマリ スタック、独自の検証コマンド、およびモジュールごとのドキュメントを備えた単一のコードベース。デフォルト — リポジトリが明らかにハブでない限り、オンボーディングはそれを想定します。
たとえば、Django API、Vue アプリ、TypeScript Lambda サービスなどです。
調整リポジトリ

オーケストレーター マニフェストを介してサブリポジトリ全体で作業を調整し、それぞれが独自のリポジトリにコミットする子プランと、境界ルールとナビゲーション インデックスを生成します。
たとえば、5 つの製品リポジトリを調整するハブです。
別のレイヤー。競合するのではなく、補完的です。
Deep Work Plan は別の足場ではありません。これは、仕様主導型または足場ツールの下にある方法論レイヤーであり、数時間にわたる自律実行に焦点を当てています。
Dailybot によって構築されました。分散チーム向けの非同期スタンドアップを支援する企業です。内部的には、Deep Work Plans を使用して、Django、Vue、TypeScript Lambda、Astro エージェントにまたがる運用リポジトリを操縦可能にしました。数か月間実稼働環境で使用した後、MIT の下でこの方法論をオープンソース化しました。
エージェントに 1 行を渡す (/init.md を指定する) と、リポジトリが AI ファーストになります。スキル、スタックに関する理由がインストールされ、完全な AGENTS.md 階層がコミットされます。そこから、何時間も自律的に実行されるディープ ワーク プランを作成して実行します。
MIT ライセンス · ゼロ テレメトリ · gitignored .dwp/ フォルダーに出力します。

## Original Extract

Context matters more than models. Deep Work Plan turns any repository into a structured environment where any coding agent finishes long-horizon work.

Deep Work Plan — structured execution for AI coding agents below to whichever theme the user picked. -->
Skip to content
Deep Work Plan Methodology · Spec · Kit Methodology Spec Kit Examples Trust GitHub EN
Open methodology · MIT · Agent-agnostic
Models matter. Context matters more.
Deep Work Plan turns any repository into a structured environment — context, guardrails, and a durable plan — where any coding agent executes with precision and finishes long-horizon work.
Copy the init.md prompt and paste it into your coding agent — Claude Code, Cursor, Codex, or any other — to make any repository AI-first.
Deep Work Plan is spec-driven development where the repository itself becomes the harness.
AI coding agents are remarkably effective in short bursts. On long-horizon work — a migration, a new subsystem, a refactor across dozens of files — they drift: context fills up, decisions are forgotten, and multi-hour tasks are abandoned halfway through.
Deep Work Plan answers with spec-driven development: the plan is the durable source of truth, and agents execute against explicit acceptance criteria and validation gates. Drift drops, the work stays verifiable, and any agent can resume it across sessions.
It is also harness engineering made portable. An agent harness is the scaffolding around a model — context, tools, control loop, guardrails, resumable state — that makes it reliable. Deep Work Plan installs that harness into the repository itself (AGENTS.md, docs, the .agents/ skills home, the DWP skill), so any agent can pilot any repo. Born at Dailybot, battle-tested for months, and released as the DailybotHQ/deepworkplan-skill.
Point it at any repository. It reasons — it does not copy-paste.
The onboarding flow inspects your repository's actual languages, frameworks, package manager, and validation commands, then generates artifacts adapted to that repository. A generic stub is treated as a failure.
01 Reasons about your stack and archetype
Reads manifests, folder layout, and CI to infer the real test, lint, and build commands, then classifies the repository as an individual repo or an orchestrator hub.
02 Generates AGENTS.md, docs/, and per-module docs
A reasoned AGENTS.md, a categorized docs/ hierarchy, and a README plus docs/ inside each major module — filled with your repository's real commands, not placeholders.
03 Scaffolds .agents/ with the .claude to .agents symlink
A cross-agent .agents/ directory (skills, agents, commands) and the .claude to .agents symlink, mirroring CLAUDE.md to AGENTS.md, so every tool reads one source of truth.
04 Installs the DWP skill and scaffolds .dwp/
Wires the Deep Work Plan skill and creates the gitignored .dwp/ folder for plans and drafts, then optionally layers opt-in addons such as devcontainer support.
One instruction. The repository does the rest.
You do not pick an install method or copy a template. You hand your agent one line; it installs the skill — the reusable engine — and adapts your repository to it.
It reads the onboarding prompt at deepworkplan.com/init.md and the methodology, specification, and kit it links to — the standard it is about to adopt.
02 It installs the Deep Work Plan skill
The skill is the engine — the same in every repository. One command pulls in the router and its sub-skills (create, execute, refine, resume, status, verify, onboard, author) for Claude Code, Cursor, Codex, Gemini, and Copilot.
Reasoning about your real stack — never copy-pasting — it writes AGENTS.md, a categorized docs/ tree, per-module READMEs, a reasoned .agents/ kit, and a gitignored .dwp/. Your repository becomes the harness.
Generate long-horizon Deep Work Plans for any task and run them step by step, with explicit acceptance criteria, validation gates, and resumable state — autonomously, for hours.
The skill is installed identically everywhere; what is adapted is your repository — the AGENTS.md, docs, and reasoned .agents/ kit generated for your stack. That split is what makes the methodology a reusable standard rather than a one-off scaffold.
Everything your agent needs to work autonomously.
One run, committed atomically. Every output is Markdown and every change is auditable.
AGENTS.md at the repository root
Reasoned from your repository's actual stack, commands, and structure — not a template with placeholders. CLAUDE.md is symlinked to AGENTS.md.
Categorized docs/ and per-module docs
Architecture, setup, standards, and troubleshooting — plus a README and docs/ inside each major module, generated from your codebase.
.agents/ with the .claude to .agents symlink
A cross-agent .agents/ directory (skills, agents, commands) with the .claude to .agents symlink so every tool reads one source of truth.
The Deep Work Plan skill, installed
create, execute, refine, resume, status, verify, onboard, and author — available to your agent as a single skill pack, with no per-repository copy.
/dwp-verify produces an objective pass/fail report against the specification, so "AI-first" is verified, not asserted — and re-verifiable after every plan.
Onboarding classifies your repository as an individual repo (the common case) or an orchestrator hub that coordinates child plans across repositories.
A living kit your repository grows
The author sub-skill (skill-create, agent-create) lets the repository evolve its own skills, agents, and commands; opt-in maintenance add-ons such as dependency-upgrade help it keep itself up to date.
No daemon and no external state. Plans and drafts land in a gitignored .dwp/ folder, and any task resumes from git alone — even after context overflows.
Works with the agent you already use.
One methodology, many adapters. Markdown couples the framework to nothing — every agent that reads Markdown can run a Deep Work Plan.
Reference implementation, with native WebFetch and slash commands.
Full adapter. Use the offline bundle if WebFetch is gated.
Offline bundle recommended; rules installed under .codex/.
Full adapter — the dwp-* commands run via AGENTS.md and # procedures.
Requires Gemini 2.5 Pro or newer, with native WebFetch.
Open source. Reads AGENTS.md natively and runs dwp-* via # commands.
Rules plus # command procedures drive the full Deep Work Plan loop.
Open source. Markdown rules and # commands run every dwp-* step.
Full adapter with a native command surface.
Reasoning presets for the stacks that matter.
These are reasoning aids, not templates. Onboarding reads your repository's real manifests and adapts per stack — it never blind-copies a preset. Monorepos get per-module docs.
TypeScript · Node Express · Fastify
TypeScript · Lambda Serverless · SAM
Individual repository, or orchestrator hub.
Onboarding forks on the archetype. Most repositories are individual repos. A hub coordinates child Deep Work Plans across many repositories. The methodology handles both as first-class.
A single codebase with one primary stack, its own validation commands, and per-module docs. The default — onboarding assumes it unless the repository is clearly a hub.
For example, a Django API, a Vue app, or a TypeScript Lambda service.
A coordination repository that orchestrates work across sub-repositories via an orchestrator manifest, spawning child plans that each commit in their own repository, plus boundary rules and a navigation index.
For example, a hub coordinating five product repositories.
A different layer. Complementary, not competing.
Deep Work Plan is not another scaffolder. It is the methodology layer underneath any spec-driven or scaffolding tool, focused on multi-hour autonomous runs.
Built by Dailybot — the company behind asynchronous standups for distributed teams. Internally we used Deep Work Plans to make production repositories spanning Django, Vue, TypeScript Lambda, and Astro agent-pilotable. After months of production use, we open-sourced the methodology under MIT.
Hand your agent one line — point it at /init.md — and it makes your repository AI-first: it installs the skill, reasons about your stack, and commits a complete AGENTS.md hierarchy. From there you create and execute Deep Work Plans that run autonomously for hours.
MIT-licensed · zero telemetry · outputs to a gitignored .dwp/ folder.
