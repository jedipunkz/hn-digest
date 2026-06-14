---
source: "https://github.com/obra/superpowers"
hn_url: "https://news.ycombinator.com/item?id=48524771"
title: "Superpowers for Claude, Codex etc."
article_title: "GitHub - obra/superpowers: An agentic skills framework & software development methodology that works. · GitHub"
author: "wood_spirit"
captured_at: "2026-06-14T07:42:57Z"
capture_tool: "hn-digest"
hn_id: 48524771
score: 1
comments: 0
posted_at: "2026-06-14T06:40:32Z"
tags:
  - hacker-news
  - translated
---

# Superpowers for Claude, Codex etc.

- HN: [48524771](https://news.ycombinator.com/item?id=48524771)
- Source: [github.com](https://github.com/obra/superpowers)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T06:40:32Z

## Translation

タイトル: クロードの超大国、コーデックスなど。
記事のタイトル: GitHub - obra/superpowers: 機能するエージェント スキル フレームワークとソフトウェア開発方法論。 · GitHub
説明: 機能するエージェント スキル フレームワークとソフトウェア開発方法論。 - オブラ/超大国

記事本文:
GitHub - obra/superpowers: 機能するエージェント スキル フレームワークとソフトウェア開発方法論。 · GitHub
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
オブラ
/
超大国
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
441 コミット 441 コミット .claude-plugin .claude-plugin .codex-plugin .codex-plugin .cursor-plugin .cursor-plugin .github .github .opencode .opencode アセット アセット ドキュメント ドキュメント フック フック スクリプト スクリプト スキル スキル テスト テスト .gitattributes .gitattributes .gitignore .gitignore .version-bump.json .version-bump.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md GEMINI.md GEMINI.md ライセンス ライセンス README.md README.md リリースノート.md RELEASE-NOTES.md gemini-extension.json gemini-extension.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Superpowers は、コーディング エージェント向けの完全なソフトウェア開発手法であり、一連の構成可能なスキルと、エージェントがそれらを確実に使用できるようにするためのいくつかの初期指示に基づいて構築されています。
エージェントにスーパーパワーを与えます: Claude Code 、 Codex CLI 、 Codex App 、 Factory Droid 、 Gemini CLI 、 OpenCode 、 Cursor 、 GitHub Copilot CLI 。
それはコーディング エージェントを起動した瞬間から始まります。何かを構築していることが分かると、すぐにコードを書こうとするわけではありません。代わりに、一歩下がって、あなたが実際に何をしようとしているのかを尋ねます。
会話から仕様を引き出すと、それを実際に読んで理解できる程度に短い単位で表示します。
あなたが設計に同意した後、エージェントは、センスが悪く、判断力がなく、プロジェクトのコンテキストを持たず、テストを嫌がる熱心な若手エンジニアが従うのに十分明確な実装計画をまとめます。真の赤/緑 TDD、YAGNI (You Aren't Gonna Need It)、および DRY を強調します。
次に、「go」と言うと、サブエージェント駆動の開発プロセスが開始されます。

エージェントに各エンジニアリング タスクを実行させ、作業を検査およびレビューして、作業を続行します。クロードが、あなたが立てた計画から逸脱することなく、一度に数時間自主的に作業できることは珍しいことではありません。
他にもたくさんありますが、これがシステムの中核です。スキルは自動的に発動するため、特別なことをする必要はありません。あなたのコーディングエージェントはスーパーパワーを持っています。
Superpowers があなたがお金を稼ぐ活動をするのに役立ち、その気になれば、私のオープンソース作品のスポンサーになることを検討していただければ大変感謝します。
ハーネスにより取り付けが異なります。複数を使用する場合は、Superpowers をそれぞれに個別にインストールします。
Superpowers は、公式の Claude プラグイン マーケットプレイスから入手できます
Anthropic の公式マーケットプレイスからプラグインをインストールします。
/プラグインインストール superpowers@claude-plugins-official
Superpowers マーケットプレイスは、Superpowers およびクロード コード用のその他の関連プラグインを提供します。
/plugin Marketplace add obra/superpowers-marketplace
このマーケットプレイスからプラグインをインストールします。
/プラグインインストール superpowers@superpowers-marketplace
Superpowers は、公式 Codex プラグイン マーケットプレイスから入手できます。
プラグイン検索インターフェイスを開きます。
Superpowers は、公式 Codex プラグイン マーケットプレイスから入手できます。
Codex アプリで、サイドバーの [プラグイン] をクリックします。
「コーディング」セクションに「スーパーパワー」が表示されるはずです。
[スーパーパワー] の横にある [+] をクリックし、プロンプトに従います。
droid プラグイン マーケットプレイスに https://github.com/obra/superpowers を追加します
droid プラグインのインストール superpowers@superpowers
Gemini 拡張機能のインストール https://github.com/obra/superpowers
gemini 拡張機能がスーパーパワーを更新します
OpenCode は独自のプラグインのインストールを使用します。 Superpowers を個別にインストールする場合でも、
すでに別のハーネスで使用しています。
インスタを取得してフォローする

https://raw.githubusercontent.com/obra/superpowers/refs/heads/main/.opencode/INSTALL.md からの抜粋
詳細ドキュメント: docs/README.opencode.md
Cursor Agent チャットで、マーケットプレイスからインストールします。
または、プラグイン マーケットプレイスで「superpowers」を検索してください。
copilot プラグイン マーケットプレイス、obra/superpowers-marketplace を追加
copilot プラグインをインストール superpowers@superpowers-marketplace
ブレーンストーミング - コードを記述する前にアクティブ化します。質問を通じて大まかなアイデアを洗練し、代替案を検討し、検証のためにデザインをセクションに分けて提示します。設計ドキュメントを保存します。
using-git-worktrees - 設計の承認後にアクティブ化されます。新しいブランチに分離されたワークスペースを作成し、プロジェクトのセットアップを実行し、クリーンなテスト ベースラインを検証します。
writing-plans - 承認されたデザインでアクティブ化されます。仕事を一口サイズのタスクに分割します (それぞれ 2 ～ 5 分)。すべてのタスクには、正確なファイル パス、完全なコード、検証手順が含まれています。
サブエージェント駆動開発または計画の実行 - 計画とともにアクティブ化されます。 2 段階のレビュー (仕様準拠、次にコード品質) を使用してタスクごとに新しいサブエージェントをディスパッチするか、人間によるチェックポイントを使用してバッチで実行します。
テスト駆動開発 - 実装中にアクティブ化されます。 RED-GREEN-REFACTOR を強制します。失敗するテストを作成し、失敗するのを監視し、最小限のコードを作成し、合格するのを監視し、コミットします。テスト前に書かれたコードを削除します。
requesting-code-review - タスク間でアクティブ化されます。計画に照らしてレビューし、重大度ごとに問題を報告します。重大な問題が進行を妨げます。
finish-a-development-branch - タスクが完了するとアクティブ化されます。テストを検証し、オプション (マージ/PR/保持/破棄) を提示し、ワークツリーをクリーンアップします。
エージェントはタスクの前に、関連するスキルをチェックします。提案ではなく必須のワークフロー。
テスト駆動開発 - RED-GREEN-REFACTOR サイクル (テストのアンチパターン参照を含む)
体系的なデバッグ - 4 段階の根本原因プロセス (根本原因を含む)

トレース、多層防御、条件ベースの待機技術)
完了前検証 - 実際に修正されていることを確認する
ブレーンストーミング - ソクラテス的なデザインの洗練
計画の作成 - 詳細な実装計画
プランの実行 - チェックポイントを使用したバッチ実行
Patching-Parallel-agents - 同時サブエージェント ワークフロー
requesting-code-review - 事前レビューチェックリスト
受信コードレビュー - フィードバックへの応答
using-git-worktrees - 並列開発ブランチ
開発ブランチの終了 - マージ/PR 決定ワークフロー
サブエージェント駆動開発 - 2 段階のレビュー (仕様準拠、次にコード品質) による高速イテレーション
ライティングスキル - ベストプラクティスに従って新しいスキルを作成します (テスト方法を含む)
using-superpowers - スキル システムの紹介
テスト駆動開発 - 常に最初にテストを作成します
アドホックではなく体系的 - 推測ではなくプロセス
複雑さの軽減 - 主な目標としてのシンプルさ
主張よりも証拠 - 成功を宣言する前に検証する
オリジナルのリリース発表を読んでください。
スーパーパワーへの一般的な貢献プロセスは以下のとおりです。通常、新しいスキルの投稿は受け付けないこと、およびスキルの更新は、サポートされているすべてのコーディング エージェントで機能する必要があることに注意してください。
新しいスキルや変更されたスキルを作成およびテストするには、ライティング スキル スキルに従ってください。
プル リクエスト テンプレートに必ず記入して、PR を送信してください。
完全なガイドについては、skills/writing-skills/SKILL.md を参照してください。
Superpower の更新はコーディング エージェントに多少依存しますが、多くの場合は自動的に行われます。
MIT ライセンス - 詳細については LICENSE ファイルを参照してください
Superpowers は、Jesse Vincent と Prime Radiant の残りのメンバーによって構築されています。
Discord : コミュニティのサポート、質問、Superpowers で構築しているものを共有するために参加してください
問題: https://github.com/obra/superpowers/issue

s
リリースのお知らせ : 新しいバージョンに関する通知を受け取るにはサインアップしてください
効果的なエージェント スキル フレームワークとソフトウェア開発方法論。
読み込み中にエラーが発生しました。このページをリロードしてください。
20.2k
フォーク
レポートリポジトリ
リリース
5
v5.1.0
最新の
2026 年 5 月 4 日
+ 4 リリース
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An agentic skills framework & software development methodology that works. - obra/superpowers

GitHub - obra/superpowers: An agentic skills framework & software development methodology that works. · GitHub
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
obra
/
superpowers
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
441 Commits 441 Commits .claude-plugin .claude-plugin .codex-plugin .codex-plugin .cursor-plugin .cursor-plugin .github .github .opencode .opencode assets assets docs docs hooks hooks scripts scripts skills skills tests tests .gitattributes .gitattributes .gitignore .gitignore .version-bump.json .version-bump.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md GEMINI.md GEMINI.md LICENSE LICENSE README.md README.md RELEASE-NOTES.md RELEASE-NOTES.md gemini-extension.json gemini-extension.json package.json package.json View all files Repository files navigation
Superpowers is a complete software development methodology for your coding agents, built on top of a set of composable skills and some initial instructions that make sure your agent uses them.
Give your agent Superpowers: Claude Code , Codex CLI , Codex App , Factory Droid , Gemini CLI , OpenCode , Cursor , GitHub Copilot CLI .
It starts from the moment you fire up your coding agent. As soon as it sees that you're building something, it doesn't just jump into trying to write code. Instead, it steps back and asks you what you're really trying to do.
Once it's teased a spec out of the conversation, it shows it to you in chunks short enough to actually read and digest.
After you've signed off on the design, your agent puts together an implementation plan that's clear enough for an enthusiastic junior engineer with poor taste, no judgement, no project context, and an aversion to testing to follow. It emphasizes true red/green TDD, YAGNI (You Aren't Gonna Need It), and DRY.
Next up, once you say "go", it launches a subagent-driven-development process, having agents work through each engineering task, inspecting and reviewing their work, and continuing forward. It's not uncommon for Claude to be able to work autonomously for a couple hours at a time without deviating from the plan you put together.
There's a bunch more to it, but that's the core of the system. And because the skills trigger automatically, you don't need to do anything special. Your coding agent just has Superpowers.
If Superpowers has helped you do stuff that makes money and you are so inclined, I'd greatly appreciate it if you'd consider sponsoring my opensource work .
Installation differs by harness. If you use more than one, install Superpowers separately for each one.
Superpowers is available via the official Claude plugin marketplace
Install the plugin from Anthropic's official marketplace:
/plugin install superpowers@claude-plugins-official
The Superpowers marketplace provides Superpowers and some other related plugins for Claude Code.
/plugin marketplace add obra/superpowers-marketplace
Install the plugin from this marketplace:
/plugin install superpowers@superpowers-marketplace
Superpowers is available via the official Codex plugin marketplace .
Open the plugin search interface:
Superpowers is available via the official Codex plugin marketplace .
In the Codex app, click on Plugins in the sidebar.
You should see Superpowers in the Coding section.
Click the + next to Superpowers and follow the prompts.
droid plugin marketplace add https://github.com/obra/superpowers
droid plugin install superpowers@superpowers
gemini extensions install https://github.com/obra/superpowers
gemini extensions update superpowers
OpenCode uses its own plugin install; install Superpowers separately even if you
already use it in another harness.
Fetch and follow instructions from https://raw.githubusercontent.com/obra/superpowers/refs/heads/main/.opencode/INSTALL.md
Detailed docs: docs/README.opencode.md
In Cursor Agent chat, install from marketplace:
Or search for "superpowers" in the plugin marketplace.
copilot plugin marketplace add obra/superpowers-marketplace
copilot plugin install superpowers@superpowers-marketplace
brainstorming - Activates before writing code. Refines rough ideas through questions, explores alternatives, presents design in sections for validation. Saves design document.
using-git-worktrees - Activates after design approval. Creates isolated workspace on new branch, runs project setup, verifies clean test baseline.
writing-plans - Activates with approved design. Breaks work into bite-sized tasks (2-5 minutes each). Every task has exact file paths, complete code, verification steps.
subagent-driven-development or executing-plans - Activates with plan. Dispatches fresh subagent per task with two-stage review (spec compliance, then code quality), or executes in batches with human checkpoints.
test-driven-development - Activates during implementation. Enforces RED-GREEN-REFACTOR: write failing test, watch it fail, write minimal code, watch it pass, commit. Deletes code written before tests.
requesting-code-review - Activates between tasks. Reviews against plan, reports issues by severity. Critical issues block progress.
finishing-a-development-branch - Activates when tasks complete. Verifies tests, presents options (merge/PR/keep/discard), cleans up worktree.
The agent checks for relevant skills before any task. Mandatory workflows, not suggestions.
test-driven-development - RED-GREEN-REFACTOR cycle (includes testing anti-patterns reference)
systematic-debugging - 4-phase root cause process (includes root-cause-tracing, defense-in-depth, condition-based-waiting techniques)
verification-before-completion - Ensure it's actually fixed
brainstorming - Socratic design refinement
writing-plans - Detailed implementation plans
executing-plans - Batch execution with checkpoints
dispatching-parallel-agents - Concurrent subagent workflows
requesting-code-review - Pre-review checklist
receiving-code-review - Responding to feedback
using-git-worktrees - Parallel development branches
finishing-a-development-branch - Merge/PR decision workflow
subagent-driven-development - Fast iteration with two-stage review (spec compliance, then code quality)
writing-skills - Create new skills following best practices (includes testing methodology)
using-superpowers - Introduction to the skills system
Test-Driven Development - Write tests first, always
Systematic over ad-hoc - Process over guessing
Complexity reduction - Simplicity as primary goal
Evidence over claims - Verify before declaring success
Read the original release announcement .
The general contribution process for Superpowers is below. Keep in mind that we don't generally accept contributions of new skills and that any updates to skills must work across all of the coding agents we support.
Follow the writing-skills skill for creating and testing new and modified skills
Submit a PR, being sure to fill in the pull request template.
See skills/writing-skills/SKILL.md for the complete guide.
Superpowers updates are somewhat coding-agent dependent, but are often automatic.
MIT License - see LICENSE file for details
Superpowers is built by Jesse Vincent and the rest of the folks at Prime Radiant .
Discord : Join us for community support, questions, and sharing what you're building with Superpowers
Issues : https://github.com/obra/superpowers/issues
Release announcements : Sign up to get notified about new versions
An agentic skills framework & software development methodology that works.
There was an error while loading. Please reload this page .
20.2k
forks
Report repository
Releases
5
v5.1.0
Latest
May 4, 2026
+ 4 releases
Sponsor this project
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
