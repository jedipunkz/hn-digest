---
source: "https://github.com/simoneloru/legioni"
hn_url: "https://news.ycombinator.com/item?id=48547121"
title: "Show HN: Legioni – a bunch of AI agents always with you"
article_title: "GitHub - simoneloru/legioni: A portable, maturing team of AI coding agents. My name is Legioni, for we are many. · GitHub"
author: "leonvonblut"
captured_at: "2026-06-15T21:54:49Z"
capture_tool: "hn-digest"
hn_id: 48547121
score: 1
comments: 0
posted_at: "2026-06-15T21:16:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Legioni – a bunch of AI agents always with you

- HN: [48547121](https://news.ycombinator.com/item?id=48547121)
- Source: [github.com](https://github.com/simoneloru/legioni)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T21:16:23Z

## Translation

タイトル: 番組 HN: Legioni – たくさんの AI エージェントがいつもあなたと一緒
記事のタイトル: GitHub - simoneloru/legioni: AI コーディング エージェントのポータブルで成熟したチーム。私の名前はレジョーニです、私たちはたくさんいるからです。 · GitHub
説明: AI コーディング エージェントのポータブルで成熟したチーム。私の名前はレジョーニです、私たちはたくさんいるからです。 - シモネロール/レジョーニ

記事本文:
GitHub - simoneloru/legioni: AI コーディング エージェントのポータブルで成熟したチーム。私の名前はレジョーニです、私たちはたくさんいるからです。 · GitHub
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
シモネロール
/
レジョーニ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .github/ workflows .github/ workflows bin bindefaultsdefaults docs docs src src .gitignore .gitignore LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルナビゲーション
お客様の作業を調整し、あらゆるタスクから学習する AI コーディング エージェントのチーム。
# ワンタイムインストール (プロモート、インストール、アップグレードチームに必要)
npm install -g Legioni
プロジェクトの cd
Legioni init # スタックを検出し、エージェントをコンパイルします
または、何もインストールせずに:
プロジェクトの cd
npxレギオーニ初期化
次に、オープンコードを開始し、タスクを入力します。
オープンコード
@orchestratorはテストを含むtruncate(text, max_len, suffix='...')関数を追加します
それだけです。オーケストレーターは作業を計画し、専門のエージェントに委任し、テストが合格するまでループします。見てください。
オーケストレーター → アーキテクト → plan.md
→ 実装者 → コード + テスト
→ 審査員 → 合否
→ テストストラテジスト → エッジケース + フルスイート
← 完了: コードが作成され、テストされ、レビューされます
セッション後、エージェントは学んだことから教訓を提案します。
Legioni Promotion # 各レッスンを確認し、承認または拒否します
Legioni install # 承認されたレッスンでエージェントを再コンパイルする
次のタスク、これらのレッスンはアクティブです。チームは回を重ねるごとに良くなっていきます。
プロジェクト
言語
テスト
何が起こったのか
Apache Commons 圧縮
Java / Maven
1890年
すべて合格し、レビュー担当者が実際の mvn テストを実行しました
Apache Commons テキスト
Java / Maven
1890年
すべて合格、オーケストレーターは破損した既存のコードを修正、レビュー担当者は検証済み
切り捨て関数
Python / パイテスト
13
アーキテクトの仕様には算術エラーがあり、レビュー担当者がそれを発見し、実装者が修正しました
スラッギファイ + Unicode
Python / パイテスト
50
実装者はノルディック文字を見逃し、レビュー担当者

失敗しました。実装者はサイクル 2 で修正しました。レッスンはプロモートされました。次のプロジェクト、最初のパスでキャッチされました。
どのように見えるか
$ レジョーニ初期化
プロジェクト偵察を実行中...完了
→ .legioni/project.md
コンパイルチーム → opencode エージェント ... 完了
→ ~/.config/opencode/agent/architect.md
→ ~/.config/opencode/agent/implementer.md
→ ~/.config/opencode/agent/orchestrator.md
→ ~/.config/opencode/agent/reviewer.md
→ ~/.config/opencode/agent/test-strategist.md
→ ~/.config/opencode/agent/db-expert.md
Legioniの初期化が完了しました。
実際のプロジェクトでは、recon がスタックを検出します。
# .legioni/project.md (自動生成)
# # スタック
- 言語 : Java
- フレームワーク : Maven
# # コマンド
- ビルド: ` mvn コンパイル `
- テスト: ` mvn テスト `
- ターゲットテスト: ` mvn test -Dtest=<TestClass> `
セッション後、エージェントはレッスン候補をステージングします。
$ レジョーニ プロモーション
.legioni/lessons.staging.*.md からステージングされたレッスンを読み取ります ...
─────────────────────
役割: レビュアー スラグ: [nordic-char-limitation-in-nfkd]
─────────────────────
状況: Unicode 正規化実装のレビュー
純粋な NFKD + ASCII エンコーディングを使用しました。
決定: ø、Ø、æ、Æ に値がないため、失敗としてフラグが立てられます。
NFKD 分解 - 完全に削除されます。
理由: NFKD の前に手動の音訳テーブルが必要です。
─────────────────────
宣伝しますか？ [y/n/q]y
→ ~/.legioni/lesson に昇格

s/reviewer/[nordic-...].md
─────────────────────
役割: オーケストレーター スラグ: [タスクの概要-精度]
─────────────────────
状況: コードベースには Hex クラスがありませんでした - 完全な
クラスを最初から作成する必要がありました。
決定: 承認を得て詳細なタスク概要を作成しました
null の処理と 16 進数のアルファベットをカバーする基準。
─────────────────────
宣伝しますか？ [y/n/q] n
コマンド
コマンド
何をするのか
レギオーニ初期化
セットアップ。チーム ストアの足場を築き、プロバイダーを選択し、スタックを検出し、エージェントをコンパイルします。
レギオーニのインストール
レッスンをプロモートするか構成を変更した後、エージェントを再コンパイルします。
レジョーニのアップデート
スタックを再検出して再コンパイルします。プロジェクトが変更されたときに使用します。
レジョーニのプロモーション
段階的なレッスン候補をインタラクティブにレビューします。
レジョーニアップグレードチーム
デフォルトをチームストアと比較し、変更されたロールをアップグレードします。
Legioni 設定セットプロバイダー
モデルプロバイダーを変更します (対話型メニュー)。
Legioni config set-model <役割> <モデル>
1 つのロールのモデルをオーバーライドします。
レジョーニ構成リスト
現在のプロバイダーとモデルの割り当てを表示します。
仕組み
Legioni init は、ロール定義を ~/.legioni/roles/ にコピーします。ストアはマシン間で移植可能です。
Legioni install は各ロールを読み取り、プロモートされたレッスンを追加し、モデルのオーバーライドを適用して、エージェント ファイルを ~/.config/opencode/agent/ に書き込みます。
Legioni init はスタックも検出し、 o に登録された .legioni/project.md を書き込みます。

ペンコード.json 。
セッション中に、エージェントはワークスペースの成果物 (計画、レビュー、テスト結果) を作成し、レッスン候補をステージングします。
Legioni Promotion を使用すると、レッスンをレビューして宣伝できます。これらは次回のコンパイル時にエージェント プロンプトに挿入されます。
Legioni アップグレード チームは、ストアを Legioni リリースの新しいデフォルトと同期します。
ポータブルで成熟した AI コーディング エージェントのチーム。私の名前はレジョーニです、私たちはたくさんいるからです。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
v0.4.4
最新の
2026 年 6 月 15 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A portable, maturing team of AI coding agents. My name is Legioni, for we are many. - simoneloru/legioni

GitHub - simoneloru/legioni: A portable, maturing team of AI coding agents. My name is Legioni, for we are many. · GitHub
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
simoneloru
/
legioni
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github/ workflows .github/ workflows bin bin defaults defaults docs docs src src .gitignore .gitignore LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A team of AI coding agents that coordinates your work and learns from every task.
# one-time install (needed for promote, install, upgrade-team)
npm install -g legioni
cd your-project
legioni init # detects stack, compiles agents
Or without installing anything:
cd your-project
npx legioni init
Then start opencode and type your task:
opencode
@orchestrator add a truncate(text, max_len, suffix='...') function with tests
That's it. The orchestrator plans the work, delegates to specialist agents, and loops until tests pass. You watch.
orchestrator → architect → plan.md
→ implementer → code + tests
→ reviewer → passes or fails
→ test-strategist → edge cases + full suite
← done: code is written, tested, and reviewed
After the session, the agents propose lessons from what they learned:
legioni promote # review each lesson, approve or reject
legioni install # recompile agents with the approved lessons
Next task, those lessons are active. The team gets better each time.
Project
Language
Tests
What happened
Apache Commons Compress
Java / Maven
1890
All pass, reviewer ran real mvn test
Apache Commons Text
Java / Maven
1890
All pass, orchestrator fixed corrupted existing code, reviewer verified
truncate function
Python / pytest
13
Architect's spec had arithmetic errors, reviewer caught them, implementer fixed
slugify + Unicode
Python / pytest
50
Implementer missed Nordic letters, reviewer failed it, implementer fixed on cycle 2. Lesson promoted. Next project, caught on first pass.
What it looks like
$ legioni init
Running project recon ... done
→ .legioni/project.md
Compiling team → opencode agents ... done
→ ~/.config/opencode/agent/architect.md
→ ~/.config/opencode/agent/implementer.md
→ ~/.config/opencode/agent/orchestrator.md
→ ~/.config/opencode/agent/reviewer.md
→ ~/.config/opencode/agent/test-strategist.md
→ ~/.config/opencode/agent/db-expert.md
legioni init complete.
On a real project, recon detects your stack:
# .legioni/project.md (auto-generated)
# # Stack
- Language : Java
- Framework : Maven
# # Commands
- Build : ` mvn compile `
- Test : ` mvn test `
- Targeted test : ` mvn test -Dtest=<TestClass> `
After a session, agents stage lesson candidates:
$ legioni promote
Reading staged lessons from .legioni/lessons.staging.*.md ...
────────────────────────────────────────────────────────────
Role: reviewer Slug: [nordic-char-limitation-in-nfkd]
────────────────────────────────────────────────────────────
Situation: Reviewing a Unicode normalization implementation
that used pure NFKD + ASCII encoding.
Decision: Flagged as failure because ø, Ø, æ, Æ have no
NFKD decomposition — they get dropped entirely.
Why: A manual transliteration table is needed before NFKD.
────────────────────────────────────────────────────────────
Promote? [y/n/q] y
→ Promoted to ~/.legioni/lessons/reviewer/[nordic-...].md
────────────────────────────────────────────────────────────
Role: orchestrator Slug: [task-brief-precision]
────────────────────────────────────────────────────────────
Situation: The codebase had no Hex class — the complete
class had to be created from scratch.
Decision: Wrote a detailed task brief with acceptance
criteria covering null handling and hex alphabet.
────────────────────────────────────────────────────────────
Promote? [y/n/q] n
Commands
Command
What it does
legioni init
Setup. Scaffolds team store, picks provider, detects stack, compiles agents.
legioni install
Recompile agents after promoting lessons or changing config.
legioni update
Re-detect stack and recompile. Use when the project changed.
legioni promote
Review staged lesson candidates interactively.
legioni upgrade-team
Diff defaults against your team store and upgrade changed roles.
legioni config set-provider
Change model provider (interactive menu).
legioni config set-model <role> <model>
Override the model for one role.
legioni config list
Show current provider and model assignments.
How it works
legioni init copies role definitions into ~/.legioni/roles/ . The store is portable across machines.
legioni install reads each role, appends promoted lessons, applies model overrides, writes agent files to ~/.config/opencode/agent/ .
legioni init also detects your stack and writes .legioni/project.md , registered in opencode.json .
During a session, agents write workspace artifacts (plan, review, test results) and stage lesson candidates.
legioni promote lets you review and promote lessons. They get injected into agent prompts on next compile.
legioni upgrade-team syncs your store with newer defaults from legioni releases.
A portable, maturing team of AI coding agents. My name is Legioni, for we are many.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
v0.4.4
Latest
Jun 15, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
