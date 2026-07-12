---
source: "https://github.com/dotbrt/dejaview"
hn_url: "https://news.ycombinator.com/item?id=48879501"
title: "Show HN: DejaView – Terminal dashboard for all your Claude Code sessions"
article_title: "GitHub - dotbrt/dejaview: TUI to browse and resume Claude Code sessions across projects · GitHub"
author: "thisbrt"
captured_at: "2026-07-12T09:52:43Z"
capture_tool: "hn-digest"
hn_id: 48879501
score: 1
comments: 0
posted_at: "2026-07-12T08:52:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: DejaView – Terminal dashboard for all your Claude Code sessions

- HN: [48879501](https://news.ycombinator.com/item?id=48879501)
- Source: [github.com](https://github.com/dotbrt/dejaview)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T08:52:40Z

## Translation

タイトル: HN を表示: DejaView – すべてのクロード コード セッションのターミナル ダッシュボード
記事のタイトル: GitHub - dotbrt/dejaview: プロジェクト全体でクロード コード セッションを参照および再開するための TUI · GitHub
説明: プロジェクト全体でクロード コード セッションを参照および再開するための TUI - dotbrt/dejaview
HN テキスト: こんにちは、HN!これは私にとって初めてのオープンソース プロジェクトです。これは特別なことではありませんが、ちょっと面倒です。claude --resume は、セッションが存在したディレクトリを覚えている場合にのみ機能します。dejaview は、~/.claude/projects/ 内のトランスクリプトを読み取り、プロジェクトごとにセッションをグループ化し、Enter キーで任意のセッションに戻ることができる TUI です。読み取り専用、ローカル専用、構成なし。テキストで構築されています。
試してみてください: uvx dejaview フィードバックは大歓迎です!

記事本文:
GitHub - dotbrt/dejaview: プロジェクト全体でクロード コード セッションを参照および再開するための TUI · GitHub
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
ドットブレット
/
デジャビュー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店T

ags ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows 資産 アセット アセット dejaview dejaview テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイル ナビゲーション
マシン上のすべてのプロジェクトにわたる、すべてのクロード コード セッションのターミナル ダッシュボード。
どこでもクロードコードを使っていますね。ここではリファクタリング、あそこではサイドプロジェクト、昼食後に終わらせると誓ったバグフィックス - 3 日前。 claude --resume は、どのディレクトリにいたかを覚えている場合にのみ役に立ちます。残りは cd と考古学を推測することです。
dejaview は、忘れ去られたセッションの山をダッシュボードに変換します。
すべてを一度に表示 - 触れたすべてのプロジェクトを最新順に並べ替え、14 日間のアクティビティ スパークラインを表示します。
どこから中断したかがわかります - 各セッションには、タイトル、最後のプロンプト、開始方法、期間、git ブランチが表示されます。
すぐに戻ってください - Enter キーを押すと、正しいディレクトリ、正しいセッションで再開されます。
ゼロセットアップ - 設定もデーモンもアカウントも必要ありません。実行するとセッションが見つかります。
セッションの詳細 - 最後のプロンプト、セッションの開始方法、すぐに貼り付けられる再開コマンド:
インストールせずに試してみてください ( uv ):
uv ツール install dejaview # または: pipx install dejaview
またはリポジトリから直接:
uv ツールをインストール git+https://github.com/dotbrt/dejaview
セッションは ~/.claude/projects から読み取られます。あなたが他の場所に住んでいる場合:
dejaview --dir /path/to/projects
キー
キー
アクション
入力してください
選択したセッションを再開します (またはペインにジャンプします)
c
再開コマンドをクリップボードにコピー
r
トランスクリプトを再スキャンする
タブ
ペインの切り替え
Esc
プロジェクトペインに戻る
q
やめる
仕組み
読み取り専用、ローカル専用。 ~/.claude/projects/ 内の *.jsonl トランスクリプトを解析し、人間のプロンプトをカウントします (フック、サイドチェーン、システムは無視します)。

m ノイズ)、作業ディレクトリごとにセッションをグループ化します。何も書き込まれず、マシンから何も残されません。
TUI によるプロジェクト全体のクロード コード セッションの参照と再開
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

TUI to browse and resume Claude Code sessions across projects - dotbrt/dejaview

Hi HN! This is my first ever open-source project. It's nothing fancy, but it scratches an itch: claude --resume only works if you remember which directory the session was in. dejaview is a TUI that reads the transcripts in ~/.claude/projects/, groups sessions by project, and lets you jump back into any of them with Enter. Read-only, local-only, no config. Built with Textual.
Try it: uvx dejaview Feedback is very welcome!

GitHub - dotbrt/dejaview: TUI to browse and resume Claude Code sessions across projects · GitHub
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
dotbrt
/
dejaview
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows assets assets dejaview dejaview tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Terminal dashboard for all your Claude Code sessions, across every project on your machine.
You use Claude Code everywhere. A refactor here, a side project there, a bugfix you swore you'd finish after lunch - three days ago. claude --resume only helps if you remember which directory you were in. The rest is cd -and-guess archaeology.
dejaview turns that pile of forgotten sessions into a dashboard:
See everything at once - every project you've touched, sorted by recency, with a 14-day activity sparkline
Know where you left off - each session shows its title, your last prompt, how it started, duration, and git branch
Jump straight back in - hit Enter and you're resumed, in the right directory, on the right session
Zero setup - no config, no daemon, no account. Run it and it finds your sessions.
Drilling into a session - last prompt, how it started, and the ready-to-paste resume command:
Try it without installing ( uv ):
uv tool install dejaview # or: pipx install dejaview
Or straight from the repo:
uv tool install git+https://github.com/dotbrt/dejaview
Sessions are read from ~/.claude/projects . If yours live elsewhere:
dejaview --dir /path/to/projects
Keys
Key
Action
Enter
resume selected session (or jump to pane)
c
copy resume command to clipboard
r
rescan transcripts
Tab
switch pane
Esc
back to projects pane
q
quit
How it works
Read-only, local-only. It parses the *.jsonl transcripts in ~/.claude/projects/ , counts human prompts (ignoring hooks, sidechains, and system noise), and groups sessions by working directory. Nothing is written and nothing leaves your machine.
TUI to browse and resume Claude Code sessions across projects
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
