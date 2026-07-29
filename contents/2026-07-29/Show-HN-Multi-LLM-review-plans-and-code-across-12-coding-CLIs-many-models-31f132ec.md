---
source: "https://github.com/beastlabai/multi-llm-plugin"
hn_url: "https://news.ycombinator.com/item?id=49101866"
title: "Show HN: Multi-LLM – review plans and code across 12 coding CLIs, many models"
article_title: "GitHub - beastlabai/multi-llm-plugin: Multi-LLM orchestration plugin for Claude Code: parallel plan review, task generation, implementation, and code review across multiple LLM providers. · GitHub"
author: "beastlabai"
captured_at: "2026-07-29T20:09:42Z"
capture_tool: "hn-digest"
hn_id: 49101866
score: 2
comments: 0
posted_at: "2026-07-29T19:26:37Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Multi-LLM – review plans and code across 12 coding CLIs, many models

- HN: [49101866](https://news.ycombinator.com/item?id=49101866)
- Source: [github.com](https://github.com/beastlabai/multi-llm-plugin)
- Score: 2
- Comments: 0
- Posted: 2026-07-29T19:26:37Z

## Translation

タイトル: HN の表示: マルチ LLM – 12 のコーディング CLI、多くのモデルにわたる計画とコードを確認する
記事のタイトル: GitHub - Beastlabai/multi-llm-plugin: クロード コード用のマルチ LLM オーケストレーション プラグイン: 複数の LLM プロバイダーにわたる並列計画レビュー、タスク生成、実装、およびコード レビュー。 · GitHub
説明: Claude Code 用のマルチ LLM オーケストレーション プラグイン: 複数の LLM プロバイダーにわたる並列計画レビュー、タスク生成、実装、およびコード レビュー。 - Beastlabai/multi-llm-plugin

記事本文:
GitHub - Beastlabai/multi-llm-plugin: クロード コード用のマルチ LLM オーケストレーション プラグイン: 複数の LLM プロバイダーにわたる並列計画レビュー、タスク生成、実装、およびコード レビュー。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ビーストラバ

私は
/
マルチLLMプラグイン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
45 コミット 45 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows skill/ multi-llm skill/ multi-llm .gitignore .gitignore ライセンス ライセンス README.md README.md TODO.md TODO.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
multi-llm - ハーネスを利用する
あなたのコードベースのための群衆の知恵。このプラグインを使用すると、複数の AI コーディング ツールと LLM で同じ作業を並行して実行し、そのフィードバックを統合して、単一のモデルでは見逃される可能性のあるバグ、盲点、改善点を見つけることができ、計画、実装、コード レビューを改善できます。
多様性は 2 つの軸で機能します。異なる LLM モデル (Opus、GPT、Gemini、Grok、Composer など) を取得するだけでなく、異なるコード ハーネス (Codex、OpenCode、Cursor Agent、Gemini CLI、Grok Build、Cline、goose、Aider、Antigravity CLI、Pi など) も取得できます。それぞれが独自のプロンプト、ツール、コンテキスト収集動作を備えています。同じモデルを実行している 2 つのハーネスでも異なる結果が明らかになる可能性があるため、両方の軸を組み合わせることで群衆の範囲がさらに広がります。
📺 実際のプラグインのチュートリアルについては、ビデオチュートリアルをご覧ください。
計画レビュー、タスク生成、実装、およびコード レビューを複数の LLM プロバイダ間で並行して調整します。または、すべてのモデルに計画に関する同じ自由記述の質問をして、1 つの統合された回答を得ることができます。
multi-llm は、単一のワークフローを複数の CLI ベースの LLM ( Cursor Agent 、 Gemini CLI 、 Grok Build 、 Codex 、 OpenCode 、 Kilocode 、 Cline 、 goose 、 Aider 、 Antigravity CLI 、 Pi 、および Claude Code 自体) に展開し、その提案を検証して統合し、 Claude Code str を渡します。

結果を適用するための指示を作成します。オーケストレーターはコードを直接変更することはありません。オーケストレーターは、Claude Code が独自のツールを通じて実行する JSON を生成するため、すべての変更がレビュー可能になります。
/plugin マーケットプレイス add Beastlabai/multi-llm-plugin
/プラグインインストール multi-llm@beastlabai
/reload-プラグイン
/リロードスキル
/multi-llm:multi-llm --init
以前にマーケットプレイスを追加したことがありますか (別のプロジェクトなど)? /plugin Marketplace add は既存のマーケットプレイスを再フェッチしません。ローカルにキャッシュされたコピーをサイレントに再利用するため、/plugin install では古いバージョンが提供されます。代わりに「更新」に従ってください。
次に、スキルを呼び出します (プラグイン スキルは名前空間が指定された plugin:skill です)。
/multi-llm:multi-llm プラン/my-feature.md
クロードは、「複数のモデルで計画をレビューする」、「複数の LLM コード レビューを実行する」などを要求したときに、スキルを自動的にロードすることもできます。
マーケットプレイスのキャッシュは、明示的に要求した場合にのみ更新されます。/plugin マーケットプレイスの更新は、実際に GitHub から最新のコミットを取得するステップです。次に、プラグインを更新 (またはインストール) してリロードします。
/プラグイン マーケットプレイスの更新
# ユーザースコープとしてインストールした場合は「--scope local」を削除します
！クロード プラグイン更新 multi-llm@beastlabai --scope local
/reload-プラグイン
/リロードスキル
/multi-llm:multi-llm --init --force
プロバイダーの構成 (最初の使用前に必要)
使用するコード ハーネスとモデルに一致するように、skills/multi-llm/providers.yaml を手動で編集します。同梱されたファイルは開始点のみです。インストールした CLI のプロバイダーを保持し、残りを削除し、defaults.models (およびオプションで Quick_models ) を実際に実行する Provider:model ペアに設定します。この手順を行わないと、デフォルトの呼び出しで、設定していない可能性のあるハーネスまたはモデルが呼び出されます。形式と使用可能なキーについては、「プロバイダー」を参照してください。
ファイルはどこにありますか

/plugin 経由でインストールした後? (このリポジトリを複製するのではなく) プラグインをインストールした場合、providers.yaml はインストールされたプラグイン内、Claude Code プラグイン ディレクトリの下に存在します (通常は ~/.claude/plugins/<marketplace>/multi-llm/skills/multi-llm/providers.yaml )。これを開く最も簡単な方法は、クロード コードに「multi-llm Providers.yaml を開く」ように要求することです。スキルは独自のインストール場所を ( ${CLAUDE_SKILL_DIR} 経由で) 解決するため、クロードはその場でファイルを読み取って編集できます。
要件
なぜ
PATH 上の UV
オーケストレーターは、UV 実行 Python スクリプトとして実行されます。 uv は、最初の使用時に依存関係 (pyyaml のみ) をインストールします。モードの開始時に uv が欠落している場合 ( --init 、 --review-plan 、 --review-code など)、スキルはそれを検出し、何かを実行する前にそれをインストールすることを提案します。インストール場所とインストーラーは OS ごとに異なります。Linux/macOS では、公式インストーラー (curl -LsSf https://astral.sh/uv/install.sh | sh ) は uv を ~/.local/bin/uv に置きます。 macOS では、Homebrew は /opt/homebrew/bin/uv (Apple Silicon) または /usr/local/bin/uv (Intel) にインストールされ、MacPorts は /opt/local/bin/uv にインストールされます。 Windows では、公式インストーラーは PowerShell ( powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex" ) であり、 %USERPROFILE%\.local\bin\uv.exe を作成します。これは、Git Bash では ~/.local/bin/uv.exe ( .exe に注意してください) として認識されます。 winget install --id=astral-sh.uv -e および pipx install uv も機能します。
Windows のみ: Git for Windows (Git Bash)
Windows では必須の前提条件。スキルの命令セット全体とその許可 ( allowed-tools ) ルールは Bash のみです。 Claude Code の PowerShell ツール フォールバック (Git for Windows が存在しない場合に使用される) は明示的にサポートされていません。このスキルは、PowerShell 同等の機能を試行するのではなく、インストール ポインターによってすぐに失敗します。 Git for Windows はバンドルされた coreutil も提供します

■ コマンド ブロックは ( ls 、 grep 、 diff 、 tail 、 dirname 、 ...) に依存しますが、これらは PowerShell/cmd には存在しません。 realpath はそれらの 1 つではないことに注意してください。プラン パスはそのままオーケストレーターに渡され ( $(realpath ...) でラップされることはありません)、オーケストレーターはそれらを OS ネイティブの絶対パス自体に解決します。
1 つ以上のプロバイダー CLI
実際に実行するモデルの CLI のみが必要です。 「プロバイダー」を参照してください。
このプラグインには独自の pyproject.toml / uv.lock が同梱されているため、Python を手動でセットアップする必要はありません。モードが初めて実行されるときに uv が環境を解決します。
計画は、機能仕様、リファクタリングの概要、設計ドキュメントなど、加えたい変更を記述する単なるマークダウン ファイルです。あなたがそれを書き（またはクロードに下書きさせて）、それをマルチ llm で指します。必須のテンプレートはありません。計画が具体的であればあるほど、レビューはより有用になります。ほとんどのワークフローは計画から始まりますが、1 つの例外があります。 --ask では、質問のコンテキストをモデルに与えるための計画のみが必要です。
渡すパス (例: plan/my-feature.md ) は常に file です。そこから、multi-llm は同じフォルダー内に兄弟出力ディレクトリを派生します。このディレクトリは、拡張子を除いたプラン ファイルにちなんだ名前が付けられます。つまり、plans/my-feature.md によって plan/my-feature/ が生成されます。すべてのワークフローの状態と生成されたアーティファクトはそこに存在します。
この README とスキル ドキュメント全体で、{plan}/... はその派生ディレクトリを指します。すべてがプランの隣に保存されるため、ワークフローは完全に移植可能です。プラン フォルダーを移動またはコミットすると、その状態も一緒に移動します。
/multi-llm:multi-llm [mode-flag] <plan_path> [options] で呼び出します。モード フラグがない場合、--review-plan がデフォルトになります。
--full フェーズをエンドツーエンドで連鎖させます。計画のレビュー -> 提案の適用 -> タスクの生成 -> タスクのレビュー -> タスクの提案の適用 -> 実装 -> コードのレビュー -> コードの修正の適用。アプリごとに一時停止します

ニーズと人間の意思決定の結果を承認できるようにするためのステップを追加します。 --yes (エイリアス --non-interactive ) を渡して、パイプライン全体を完全に無人で実行します。プロンプトなし: 非対話型のモデル選択、クロードがすべての必要な人間の意思決定項目を決定し、レビュー タスクが自動的に実行されます。より細かく制御するには、代わりに --no-confirm および/または --claude-decide (以下を参照) を追加してください。
コード修正の適用には 2 つの方法があります。 --apply-code-fixes をスタンドアロン パスとして実行して、以前のレビューの修正を適用します。これは、人間による判断が必要な項目 (プロンプト、サルベージ、HTML バッジ) を処理するフェーズです。または、レビュー自体中に明らかに有効な修正のみをインラインで適用するには、--apply-fixes を --review-code の実行に追加します。
モード フラグは、順序付けされたパイプライン (まさに --full チェーン) として実行するように設計されています。レビューの各ステップでは、作業が複数の LLM に並行して行われます。各適用ステップでは、Claude Code が変更を書き戻し、ニーズと人間による決定の結果を確認するために一時停止します。
フローチャート TD
START(["plan.md"]) -.-> FULL["--full"]:::full
完全に - 。 「すべてのステップを順番に実行します」 .-> DONE(["done"])
サブグラフ PLAN [「実装計画」]
RP["--review-plan"]:::review --> AS["--apply-suggestions"]:::apply
終わり
サブグラフ TASKS ["実装タスク"]
GT["--generate-tasks"]:::gen --> RT["--review-tasks"]:::review --> ATS["--apply-task-suggestions"]:::apply
終わり
サブグラフ IMPLEMENT [「実装」]
IMP["--implement"]:::gen
終わり
サブグラフ CODE [「コードレビュー」]
RC["--review-code"]:::review --> ACF["--apply-code-fixes"]:::apply
終わり
スタート --> RP
AS --> GT
ATS --> IMP
IMP --> RC
ACF --> 完了
classDef レビュー fill:#dbeafe、ストローク:#3b82f6、color:#1e3a8a;
classDef gen fill:#ede9fe、ストローク:#8b5cf6、color:#4c1d95;
classDef 適用 fill:#ffedd5、ストローク:#f97316、color:#7c2d12;
classDef 完全な塗りつぶし:#dcfce7、ストローク:#22c55e、色:#14532d;
読み込み中

🔵 複数のモデルを並行してクエリする手順を確認してください
🟠 ステップを適用すると、Claude Code が変更を書き込み、ニーズと人間の決定項目のために一時停止します
🟣 ステップの生成と実装でタスクを生成し、コードを作成します
🟢 --full は 1 つのコマンドでフロー全体を実行します
チェーン全体を実行する必要はありません。すべてのステップは、共有された {plan}/state.json を読み書きするスタンドアロンのエントリ ポイントでもあるため、どのフェーズでも開始、停止、再開できます。いくつかのモードはパイプラインの外側にあります - --ask (計画に関する読み取り専用 Q&A)、--status (進行状況と提案される次のステップを表示)、--init (プロジェクトごとの構成を書き込む) - 一方、--full は単にエンドツーエンドでフロー全体を実行します。
# 設定されたデフォルト LLM を使用してプランを確認する
/multi-llm:multi-llm プラン/my-feature.md
# 複数のプロバイダーからモデルを明示的に選択します
/multi-llm:multi-llm --review-plan plan/my-feature.md --models codex:gpt-5.5-extra-highcursor-agent:composer-2.5
# クイックレビュー (事前に選択された少数のモデルのサブセット)
/multi-llm:multi-llm --review-plan プラン/my-feature.md --quick
# 計画に照らしてコードの変更をレビューする
/multi-llm:multi-llm --review-code plan/my-feature.md
# 設定済みのデフォルト LLM を使用して計画を完全に手動でレビュー/実装/コードレビューします
/マルチ-llm:マルチ-ll

[切り捨てられた]

## Original Extract

Multi-LLM orchestration plugin for Claude Code: parallel plan review, task generation, implementation, and code review across multiple LLM providers. - beastlabai/multi-llm-plugin

GitHub - beastlabai/multi-llm-plugin: Multi-LLM orchestration plugin for Claude Code: parallel plan review, task generation, implementation, and code review across multiple LLM providers. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
beastlabai
/
multi-llm-plugin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
45 Commits 45 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows skills/ multi-llm skills/ multi-llm .gitignore .gitignore LICENSE LICENSE README.md README.md TODO.md TODO.md View all files Repository files navigation
multi-llm - Harness the harnesses
Wisdom of crowds for your codebase. Use this plugin to improve planning, implementation, and code review by running the same work through multiple AI coding tools and LLMs in parallel - then consolidating their feedback so you catch bugs, blind spots, and improvements a single model might miss.
The diversity works on two axes: not only do you get different LLM models (Opus, GPT, Gemini, Grok, Composer, ...), you also get different code harnesses (Codex, OpenCode, Cursor Agent, Gemini CLI, Grok Build, Cline, goose, Aider, Antigravity CLI, Pi, ...) - each with its own prompting, tooling, and context-gathering behavior. Two harnesses running the same model can still surface different findings, so combining both axes widens the crowd further.
📺 Watch the video tutorial for a walkthrough of the plugin in action.
Orchestrate plan reviews, task generation, implementation, and code reviews across multiple LLM providers in parallel - or ask every model the same free-text question about a plan and get one consolidated answer.
multi-llm fans a single workflow out to several CLI-based LLMs ( Cursor Agent , Gemini CLI , Grok Build , Codex , OpenCode , Kilocode , Cline , goose , Aider , Antigravity CLI , Pi , and Claude Code itself), validates and consolidates their suggestions, and hands Claude Code structured instructions to apply the results. The orchestrators never modify your code directly - they produce JSON that Claude Code executes through its own tools, so every change stays reviewable.
/plugin marketplace add beastlabai/multi-llm-plugin
/plugin install multi-llm@beastlabai
/reload-plugins
/reload-skills
/multi-llm:multi-llm --init
Already added the marketplace before (e.g. in another project)? /plugin marketplace add does not refetch an existing marketplace - it silently reuses the locally cached copy, so /plugin install will give you a stale version. Follow Updating instead.
Then invoke the skill (plugin skills are namespaced plugin:skill ):
/multi-llm:multi-llm plans/my-feature.md
Claude can also load the skill automatically when you ask it to "review my plan with multiple models", "run a multi-LLM code review", etc.
The marketplace cache only refreshes when you explicitly ask it to - /plugin marketplace update is the step that actually pulls the latest commits from GitHub. Then update (or install) the plugin and reload:
/plugin marketplace update beastlabai
# remove "--scope local" if you installed it as user scope
! claude plugin update multi-llm@beastlabai --scope local
/reload-plugins
/reload-skills
/multi-llm:multi-llm --init --force
Configure providers (required before first use)
Edit skills/multi-llm/providers.yaml manually to match the code harnesses and models you want to use. The shipped file is a starting point only - keep the providers whose CLIs you have installed, remove the rest, and set defaults.models (and optionally quick_models ) to the provider:model pairs you actually run. Without this step, default invocations will call harnesses or models you may not have configured. See Providers for the format and available keys.
Where is the file after installing via /plugin ? If you installed the plugin (rather than cloning this repo), providers.yaml lives inside the installed plugin, under your Claude Code plugins directory - typically ~/.claude/plugins/<marketplace>/multi-llm/skills/multi-llm/providers.yaml . The quickest way to open it is to ask Claude Code to "open the multi-llm providers.yaml ": the skill resolves its own install location (via ${CLAUDE_SKILL_DIR} ), so Claude can read and edit the file in place.
Requirement
Why
uv on your PATH
The orchestrators run as uv run Python scripts. uv installs their dependencies (just pyyaml ) on first use. If uv is missing when a mode starts ( --init , --review-plan , --review-code , ...), the skill detects that and offers to install it before running anything. Install location and installer differ per OS: on Linux/macOS the official installer ( curl -LsSf https://astral.sh/uv/install.sh | sh ) puts uv in ~/.local/bin/uv ; on macOS Homebrew installs to /opt/homebrew/bin/uv (Apple Silicon) or /usr/local/bin/uv (Intel), and MacPorts to /opt/local/bin/uv ; on Windows the official installer is PowerShell ( powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex" ) and creates %USERPROFILE%\.local\bin\uv.exe — which Git Bash sees as ~/.local/bin/uv.exe (note the .exe ); winget install --id=astral-sh.uv -e and pipx install uv also work.
Windows only: Git for Windows (Git Bash)
Hard prerequisite on Windows. The skill's entire instruction set and its permission ( allowed-tools ) rules are Bash-only; Claude Code's PowerShell-tool fallback (used when Git for Windows is absent) is explicitly unsupported — the skill fails fast with an install pointer rather than attempting PowerShell equivalents. Git for Windows also supplies the bundled coreutils the command blocks depend on ( ls , grep , diff , tail , dirname , ...), which are absent from PowerShell/cmd. Note that realpath is not one of them: plan paths are passed to the orchestrators as-is (never wrapped in $(realpath ...) ), and the orchestrators resolve them to OS-native absolute paths themselves.
One or more provider CLIs
You only need the CLIs for the models you actually run. See Providers .
The plugin ships its own pyproject.toml / uv.lock , so no manual Python setup is required - uv resolves the environment the first time a mode runs.
A plan is just a markdown file describing the change you want to make - a feature spec, refactor outline, or design doc. You write it (or have Claude draft it), then point multi-llm at it. There's no required template; the more concrete the plan, the more useful the reviews. Most workflows start from a plan, with one exception: --ask only needs a plan to give the models context for your question.
The path you pass (e.g. plans/my-feature.md ) is always a file . From it, multi-llm derives a sibling output directory in the same folder, named after the plan file without its extension - so plans/my-feature.md produces plans/my-feature/ . All workflow state and generated artifacts live there:
Throughout this README and the skill docs, {plan}/... refers to that derived directory. Because everything is stored next to the plan, a workflow is fully portable - move or commit the plan folder and its state travels with it.
Invoke with /multi-llm:multi-llm [mode-flag] <plan_path> [options] . With no mode flag, --review-plan is the default.
--full chains the phases end to end: review-plan -> apply-suggestions -> generate-tasks -> review-tasks -> apply-task-suggestions -> implement -> review-code -> apply-code-fixes. It pauses at each apply step to let you approve needs-human-decision findings; pass --yes (alias --non-interactive ) to run the whole pipeline fully unattended — zero prompts: non-interactive model selection, Claude decides every needs-human-decision item, and review-tasks runs automatically. For finer control, add just --no-confirm and/or --claude-decide (see below ) instead.
Applying code fixes can happen two ways. Run --apply-code-fixes as a standalone pass to apply a previous review's fixes - this is the phase that handles needs-human-decision items (prompts, salvage, HTML badges). Or, to apply only the clearly-valid fixes inline during the review itself, add --apply-fixes to a --review-code run.
The mode flags are designed to run as an ordered pipeline (exactly what --full chains for you). Each review step fans the work out to multiple LLMs in parallel; each apply step is where Claude Code writes the changes back and pauses for any needs-human-decision findings.
flowchart TD
START(["plan.md"]) -.-> FULL["--full"]:::full
FULL -. "runs every step in order" .-> DONE(["done"])
subgraph PLAN ["Implementation Plan"]
RP["--review-plan"]:::review --> AS["--apply-suggestions"]:::apply
end
subgraph TASKS ["Implementation Tasks"]
GT["--generate-tasks"]:::gen --> RT["--review-tasks"]:::review --> ATS["--apply-task-suggestions"]:::apply
end
subgraph IMPLEMENT ["Implement"]
IMP["--implement"]:::gen
end
subgraph CODE ["Code Review"]
RC["--review-code"]:::review --> ACF["--apply-code-fixes"]:::apply
end
START --> RP
AS --> GT
ATS --> IMP
IMP --> RC
ACF --> DONE
classDef review fill:#dbeafe,stroke:#3b82f6,color:#1e3a8a;
classDef gen fill:#ede9fe,stroke:#8b5cf6,color:#4c1d95;
classDef apply fill:#ffedd5,stroke:#f97316,color:#7c2d12;
classDef full fill:#dcfce7,stroke:#22c55e,color:#14532d;
Loading
🔵 Review steps query multiple models in parallel
🟠 Apply steps let Claude Code write changes and pause for needs-human-decision items
🟣 Generate & implement steps produce tasks and write code
🟢 --full runs the entire flow with one command
You don't have to run the whole chain: every step is also a standalone entry point that reads/writes the shared {plan}/state.json , so you can start, stop, and resume at any phase. A few modes sit outside the pipeline - --ask (read-only Q&A about a plan), --status (show progress and the suggested next step), and --init (write a per-project config) - while --full simply runs the entire flow above end to end.
# Review a plan using the configured default LLMs
/multi-llm:multi-llm plans/my-feature.md
# Pick models explicitly, from several providers
/multi-llm:multi-llm --review-plan plans/my-feature.md --models codex:gpt-5.5-extra-high cursor-agent:composer-2.5
# Quick review (fewer, preselected subset of models)
/multi-llm:multi-llm --review-plan plans/my-feature.md --quick
# Review code changes against the plan
/multi-llm:multi-llm --review-code plans/my-feature.md
# Review/implement/code-review a plan using the configured default LLMs completely hands-off
/multi-llm:multi-ll

[truncated]
