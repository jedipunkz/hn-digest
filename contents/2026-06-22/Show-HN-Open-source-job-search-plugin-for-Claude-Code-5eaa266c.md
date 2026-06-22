---
source: "https://github.com/agent-data/job-search"
hn_url: "https://news.ycombinator.com/item?id=48633099"
title: "Show HN: Open-source job search plugin for Claude Code"
article_title: "GitHub - agent-data/job-search: Turn Claude Code into your job search assistant. · GitHub"
author: "jb_hn"
captured_at: "2026-06-22T18:23:11Z"
capture_tool: "hn-digest"
hn_id: 48633099
score: 2
comments: 0
posted_at: "2026-06-22T17:26:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source job search plugin for Claude Code

- HN: [48633099](https://news.ycombinator.com/item?id=48633099)
- Source: [github.com](https://github.com/agent-data/job-search)
- Score: 2
- Comments: 0
- Posted: 2026-06-22T17:26:27Z

## Translation

タイトル: Show HN: Claude Code 用のオープンソース求人検索プラグイン
記事のタイトル: GitHub - エージェント データ/求人検索: クロード コードを求人検索アシスタントに変えます。 · GitHub
説明: クロード コードを就職活動アシスタントにしましょう。 GitHub でアカウントを作成して、エージェント データ/求人検索の開発に貢献します。
HN テキスト: HN さん、Claude Code を就職支援アシスタントに変える、私がまとめたオープンソースのプラグインとスキルを共有したいと思いました。目標は、興味深く関連性のある仕事を見つけるのに必要な時間を最小限に抑えることです。仕組み: 1. リポジトリを複製し、プラグインをインストールし、`/job-search` を実行します。 2. クロードは、興味のあるロールを理解し、設定をローカルに保存するためにいくつかの質問をします。 3. 次に、Claude はライブ求人情報 (現在は LinkedIn Jobs) を取得し、投稿をユーザーの設定と比較し、関連する投稿のダイジェストを生成します。 4. [オプション] クロードは、スケジュール (毎日など) に基づいて検索を実行し、時間の経過とともに好みに一致する新しい投稿を表示します。ロードマップ上の機能: - ATS プラットフォームのサポート (例: Workday、Ashby、Greenhouse) - 追加エージェントのサポート (例: Codex、Cursor) - クロードに履歴書と職務記述書を比較させて「適合性」を推定させるスキル - 特定の職務記述書とよりよく一致するようにクロードに履歴書の編集を推奨させるスキル リポジトリ: https://github.com/agent-data/job-search (MIT ライセンス) 重要: プラグインはエージェントデータを使用します(構造化 Web データにアクセスするためのツール) を内部で使用してライブ求人情報を取得できるため、開始するには API キーが必要です (クレジット カードは必要ありません)。このプロジェクトはまだ初期段階ですが、より便利にする方法についてのフィードバックをお待ちしています。

記事本文:
GitHub - エージェントデータ/求人検索: Claude Code を求人検索アシスタントに変えます。 · GitHub
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
エージェントデータ
/
仕事探し
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店 Ta

gs ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
176 コミット 176 コミット .claude-plugin .claude-plugin .claude/ エージェント .claude/ エージェント .github .github docs docs サンプル サンプル スクリプト スクリプト 共有/ 参照 共有/ 参照 スキル スキル テンプレート テンプレート テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md TESTING.md TESTING.md TODOS.md TODOS.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Job Search は、Claude Code を求職アシスタントに変えるプラグインです。あなたが望むものを説明すると、クロードは新しい求人情報を取得し、あなたの好みに照らしてそれぞれを判断し、マッチのランク付けされたダイジェストを生成します。
プラグインをインストールして /job-search を実行します
クロードは、あなたが興味を持っている役割を理解し、設定をローカルに保存するためにいくつかの質問をします。
次に、Claude はライブ求人情報を取得し、投稿をユーザーの設定と比較し、関連する投稿のみを含むダイジェストを生成します。
オプションで、Claude はスケジュール (毎日など) で検索を実行し、時間の経過とともに好みに一致する新しい投稿を表示することもできます。
example/sample-digest.md のダイジェストの例を参照してください。
エージェント データ CLI — ジョブ データ ソース。 Agent-data.motie.dev で API キーを生成し ([プロファイル] → [API キー])、 AGENT_DATA_API_KEY=mtk_… をエクスポートし (または ~/.agent-data/config.json に保存し)、agent-data whoami で確認します。
注: 現在、agent-data は次のソースから求人情報を提供しています: LinkedIn 求人。
エージェントデータ API キーを設定します。 Agent-data.motie.dev (プロファイル → API キー) で 1 つを取得し、エクスポートします。
エクスポート AGENT_DATA_API_KEY=mtk_…
Claude Code を起動し、ローカル クローンをマーケットプレイスとして登録します。
/プラグインマーケットプレイス

/path/to/job-search を追加
/plugin install job-search@agent-data
就職活動を始めましょう。 /job-search を実行します
就職活動 — 玄関口: 新人研修、ステータス、自宅の様子。
job-preference-interview — 平易な英語での希望の概要を作成します。
job-search-run — 1 つのヘッドレス検索パス。これがスケジュールで実行されることです。
Evaluate-job-fit — 貼り付けた単一の投稿を判断します。
job-search-agent — クロードがシステムの設定、拡張、またはトラブルシューティングを行うために手に入れた操作マニュアル。
リポジトリのクローンを作成し、インストール パスを選択します。
永続的 (推奨)。ローカル クローンをマーケットプレイスとして登録し、以下をインストールします。
/plugin Marketplace add /path/to/job-search
/plugin install job-search@agent-data
1 つのセッションで、インストールは不要です。 --plugin-dir 起動フラグを使用してクロード コードを起動します。
クロード --plugin-dir /path/to/job-search
インストール後、フロント ドア スラッシュ コマンドを実行するか、単に必要なことを言います。
/求人検索:求人検索
貢献する
AI エージェントを使用してリポジトリを構築または探索しますか?アーキテクチャ、設計信念、計画のマップである AGENTS.md から始めます。開発ワークフローについては CONTRIBUTING.md を、テスト ハーネスについては TESTING.md を参照してください。
Claude Code を就職活動アシスタントにしましょう。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Turn Claude Code into your job search assistant. Contribute to agent-data/job-search development by creating an account on GitHub.

Hey HN, I wanted to share an open source plugin + skills I put together that turn Claude Code into a job search assistant. The goal is to minimize the amount of time needed to find interesting and relevant jobs. How it works: 1. Clone the repo, install the plugin, and run `/job-search` 2. Claude will ask a few questions to understand the roles you’re interested in and save your preferences locally. 3. Claude will then pull live job postings (currently, LinkedIn Jobs), compare posts against your preferences, and generate a digest of the posts that are relevant. 4. [Optionally] Claude will also run your search on a schedule (e.g., daily) to surface new posts matching your preferences over time. Features on the roadmap: - ATS platform support (e.g., Workday, Ashby, Greenhouse) - Support for additional agents (e.g., Codex, Cursor) - Skills for having Claude compare resumes to job descriptions to estimate “fit” - Skills for having Claude recommend resume edits to better align with a given job description Repo: https://github.com/agent-data/job-search (MIT license) Important: the plugin uses agent-data (a tool for accessing structured web data) under the hood to pull live job postings, so you’ll need an API key to get started (no credit card required). The project is still pretty early but would love to get feedback on ways to make it more useful!

GitHub - agent-data/job-search: Turn Claude Code into your job search assistant. · GitHub
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
agent-data
/
job-search
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
176 Commits 176 Commits .claude-plugin .claude-plugin .claude/ agents .claude/ agents .github .github docs docs examples examples scripts scripts shared/ references shared/ references skills skills templates templates tests tests .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md TESTING.md TESTING.md TODOS.md TODOS.md View all files Repository files navigation
Job Search is a plugin that turns Claude Code into a job search assistant. Describe what you want, and Claude pulls fresh job postings, judges each against your preferences, and generates a ranked digest of the matches.
Install the plugin and run /job-search
Claude will ask a few questions to understand the roles you’re interested in and save your preferences locally.
Claude will then pull live job postings, compare posts against your preferences, and generate a digest with only the posts that are relevant.
Optionally Claude can also run your search on a schedule (e.g., daily) to surface new posts matching your preferences over time.
See an example digest in examples/sample-digest.md .
The agent-data CLI — the job-data source. Generate an API key at agent-data.motie.dev (Profile → API Key), then export AGENT_DATA_API_KEY=mtk_… (or save it to ~/.agent-data/config.json ) and verify with agent-data whoami .
Note: agent-data currently provides job postings from the following sources: LinkedIn Jobs.
Set your agent-data API key. Grab one at agent-data.motie.dev (Profile → API Key), then export it:
export AGENT_DATA_API_KEY=mtk_…
Launch Claude Code then register the local clone as a marketplace:
/plugin marketplace add /path/to/job-search
/plugin install job-search@agent-data
Kick off your job search. Run /job-search
job-search — the front door: onboarding, status, and your home view.
job-preference-interview — builds your plain-English preferences brief.
job-search-run — one headless search pass; this is what the schedule runs.
evaluate-job-fit — judges a single posting you paste in.
job-search-agent — the operator manual Claude reaches for to configure, extend, or troubleshoot the system.
Clone the repo, then pick an install path.
Persistent (recommended). Register the local clone as a marketplace, then install:
/plugin marketplace add /path/to/job-search
/plugin install job-search@agent-data
One session, no install. Launch Claude Code with the --plugin-dir launch flag:
claude --plugin-dir /path/to/job-search
After installing, run the front door slash command, or just say what you want:
/job-search:job-search
Contributing
Building on or exploring the repo with an AI agent? Start at AGENTS.md , the map of the architecture, design beliefs, and plans. See CONTRIBUTING.md for the dev workflow and TESTING.md for the test harness.
Turn Claude Code into your job search assistant.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
