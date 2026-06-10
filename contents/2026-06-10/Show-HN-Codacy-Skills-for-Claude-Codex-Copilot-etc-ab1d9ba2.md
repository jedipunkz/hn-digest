---
source: "https://github.com/codacy/codacy-cloud-cli"
hn_url: "https://news.ycombinator.com/item?id=48474980"
title: "Show HN: Codacy Skills for Claude, Codex, Copilot, etc."
article_title: "GitHub - codacy/codacy-cloud-cli: A command-line tool to interact with Codacy Cloud directly from your terminal. · GitHub"
author: "claudiacsf"
captured_at: "2026-06-10T13:19:33Z"
capture_tool: "hn-digest"
hn_id: 48474980
score: 1
comments: 0
posted_at: "2026-06-10T11:57:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Codacy Skills for Claude, Codex, Copilot, etc.

- HN: [48474980](https://news.ycombinator.com/item?id=48474980)
- Source: [github.com](https://github.com/codacy/codacy-cloud-cli)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T11:57:04Z

## Translation

タイトル: HN を表示: クロード、コーデックス、副操縦士などのコーダシー スキル
記事のタイトル: GitHub - codacy/codacy-cloud-cli: ターミナルから Codacy Cloud と直接対話するためのコマンドライン ツール。 · GitHub
説明: ターミナルから直接 Codacy Cloud と対話するためのコマンドライン ツール。 - GitHub - codacy/codacy-cloud-cli: ターミナルから直接 Codacy Cloud と対話するためのコマンドライン ツール。
HN テキスト: Codacy は、エンジニア チームが AI で生成されたコードに対してコーディング標準を適用するのに役立つコード品質およびセキュリティ プラットフォームです。彼らはエージェント スキルとクラウド CLI をリリースしたばかりです。これにより、クロードなどが UI を使用せずに日常のタスクを処理できるようになります。プロンプト: 「PR 42 はゲートに失敗しています。実際のものを修正し、テストを追加し、理由を付けて誤検知を無視し、再実行してください。」分析はサーバー側で実行され、エージェント トークンは消費されません。

記事本文:
GitHub - codacy/codacy-cloud-cli: ターミナルから直接 Codacy Cloud と対話するためのコマンドライン ツール。 · GitHub
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
コーダシー
/
コーダシークラウド-cli
公共
通知
通知設定を変更するにはサインインする必要があります
追加ナビ

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット .changeset .changeset .claude/ コマンド .claude/ コマンド .codacy .codacy .github/ workflows .github/ workflows SPECS SPECS src src .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md ライセンス ライセンス README.md README.md TODO.md TODO.md package-lock.json package-lock.json package.json package.json tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json vitest.config.mts vitest.config.mts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ターミナルから直接 Codacy Cloud と対話するためのコマンドライン ツール。 Node.js と TypeScript で構築されています。
npm install -g " @codacy/codacy-cloud-cli "
ソースから
git clone https://github.com/alerizzo/codacy-cloud-cli.git
cd codacy-cloud-cli
npmインストール
npm ビルドを実行する
npmリンク
認証
対話的にログインします (推奨):
コーダシーログイン
または、CODACY_API_TOKEN 環境変数を設定します。
CODACY_API_TOKEN=あなたのトークンをここにエクスポートします
トークンは、[Codacy] > [マイ アカウント] > [アクセス管理] > [API トークン] ( リンク ) から取得できます。
login コマンドは、暗号化されたトークンを ~/.codacy/credentials に保存します。両方が存在する場合、環境変数は保存された資格情報よりも優先されます。
codacy < コマンド > [オプション]
codacy < command > --help # コマンドの詳細な使用法
グローバルオプション
オプション
説明
-o, --output <フォーマット>
出力形式: テーブル (デフォルト) または json
-V、--バージョン
バージョンを表示
-h、--ヘルプ
ヘルプを表示する
リポジトリの自動検出
git リポジトリ内でコマンドを実行すると、CLI は元のリモート URL からプロバイダ、組織、およびリポジトリを自動的に検出します。これは、これらの引数を完全にスキップできることを意味します。
# GitHub リ内

po — プロバイダー/組織/リポジトリを自動検出します
コーダシーの問題
コードのプルリクエスト 42
コーダシーツール
# または明示的に指定する
コーダシーの問題 gh my-org my-repo
サポートされているプロバイダー: GitHub ( gh )、GitLab ( gl )、Bitbucket ( bb )。
コマンド
説明
ログイン
API トークンを保存して Codacy で認証する
ログアウト
保存されている Codacy API トークンを削除する
情報
認証されたユーザー情報とその組織を表示します
リポジトリ <プロバイダ> <組織>
組織のリポジトリを一覧表示する
リポジトリ [プロバイダ] [組織] [リポジトリ]
リポジトリのメトリクスを表示するか、追加/削除/フォロー/フォロー解除/再分析します (オプションで結果を待ちます)
問題 [プロバイダ] [組織] [リポジトリ]
フィルターを使用してリポジトリ内の課題を検索する
問題 [プロバイダー] [組織] [リポジトリ] <id>
単一の問題の詳細を表示するか、無視するか無視しない
調査結果 [プロバイダー] [組織] [リポジトリ]
リポジトリまたは組織のセキュリティ調査結果を表示する
<プロバイダー> <組織> <ID> を見つけます
単一のセキュリティ検出結果の詳細を表示するか、無視するか無視します
プルリクエスト [プロバイダー] [組織] [リポジトリ] <pr>
PR 分析、問題、差分カバレッジ、および変更されたファイルを表示します。または再分析します (オプションで結果を待ちます)
ツール [プロバイダー] [組織] [リポジトリ]
リポジトリ用に構成されたリスト分析ツール
ツール [プロバイダー] [組織] [リポジトリ] <ツール>
分析ツールを有効化、無効化、または構成する
パターン [プロバイダー] [組織] [リポジトリ] <ツール>
ツールのパターンをリストするか、一括で有効化/無効化します。
パターン [プロバイダー] [組織] [リポジトリ] <ツール> <id>
パターンを表示するか、パターンのパラメータを有効化、無効化、または設定します
codacy <command> --help を実行すると、コマンドの完全な引数とオプションの詳細が表示されます。
npm start -- < コマンド > # 開発モードで実行
npm test # テストを実行する
npm run type-check # 発行せずに型チェックを行う
npm run build # 実稼働用のビルド
npm run update-api # 自動生成された API クライアントを更新します
CI/CD
CI : メインおよび

PRについて。 Node.js 18、20、22 にわたってビルドとテストを行います。
リリース : 自動バージョン管理と npm パブリッシングに変更セットを使用します。
変更を行う場合は、 npx changeset を実行し、変更内容を説明します ( patch 、minor 、または Major を選択します)。
生成された .changeset/*.md ファイルを PR に含めます
CI では、すべての PR に変更セットが含まれるように強制します (ドキュメントや CI など、バージョンのバンプを必要としない変更には npx changeset --empty を使用します)。
PR が main にマージされると、リリース ワークフローはバージョンを上げて CHANGELOG.md を更新する「chore: version package」PR を自動的に作成します。
PR が npm に発行したものを来歴とマージする
前提条件 : NPM_TOKEN シークレットが GitHub リポジトリ設定で構成されている必要があります。
ターミナルから直接 Codacy Cloud と対話するためのコマンドライン ツール。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
9
v1.2.0
最新の
2026 年 6 月 3 日
+ 8 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A command-line tool to interact with Codacy Cloud directly from your terminal. - GitHub - codacy/codacy-cloud-cli: A command-line tool to interact with Codacy Cloud directly from your terminal.

Codacy is a code quality and security platform that helps eng teams enforce coding standards against their AI generated code. They just launched agent skills and a cloud CLI, which allows Claude etc. to handle everyday tasks without using the UI E.g. prompt: "PR 42 is failing the gate, fix what's real, add the tests, ignore the false positives with a reason, re-run." The analysis runs server-side and burns no agent tokens.

GitHub - codacy/codacy-cloud-cli: A command-line tool to interact with Codacy Cloud directly from your terminal. · GitHub
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
codacy
/
codacy-cloud-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits .changeset .changeset .claude/ commands .claude/ commands .codacy .codacy .github/ workflows .github/ workflows SPECS SPECS src src .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md LICENSE LICENSE README.md README.md TODO.md TODO.md package-lock.json package-lock.json package.json package.json tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json vitest.config.mts vitest.config.mts View all files Repository files navigation
A command-line tool to interact with Codacy Cloud directly from your terminal. Built with Node.js and TypeScript.
npm install -g " @codacy/codacy-cloud-cli "
From source
git clone https://github.com/alerizzo/codacy-cloud-cli.git
cd codacy-cloud-cli
npm install
npm run build
npm link
Authentication
Log in interactively (recommended):
codacy login
Or set the CODACY_API_TOKEN environment variable:
export CODACY_API_TOKEN=your-token-here
You can get a token from Codacy > My Account > Access Management > API Tokens ( link ).
The login command stores the token encrypted at ~/.codacy/credentials . The environment variable takes precedence over stored credentials when both are present.
codacy < command > [options]
codacy < command > --help # Detailed usage for any command
Global Options
Option
Description
-o, --output <format>
Output format: table (default) or json
-V, --version
Show version
-h, --help
Show help
Repository Auto-Detection
When you run a command inside a git repository, the CLI automatically detects the provider , organization , and repository from the origin remote URL. This means you can skip those arguments entirely:
# Inside a GitHub repo — auto-detects provider/org/repo
codacy issues
codacy pull-request 42
codacy tools
# Or specify them explicitly
codacy issues gh my-org my-repo
Supported providers: GitHub ( gh ), GitLab ( gl ), Bitbucket ( bb ).
Command
Description
login
Authenticate with Codacy by storing your API token
logout
Remove stored Codacy API token
info
Show authenticated user info and their organizations
repositories <provider> <org>
List repositories for an organization
repository [provider] [org] [repo]
Show metrics for a repository, or add/remove/follow/unfollow/reanalyze it (optionally waiting for results)
issues [provider] [org] [repo]
Search issues in a repository with filters
issue [provider] [org] [repo] <id>
Show details for a single issue, or ignore/unignore it
findings [provider] [org] [repo]
Show security findings for a repository or organization
finding <provider> <org> <id>
Show details for a single security finding, or ignore/unignore it
pull-request [provider] [org] [repo] <pr>
Show PR analysis, issues, diff coverage, and changed files; or reanalyze it (optionally waiting for results)
tools [provider] [org] [repo]
List analysis tools configured for a repository
tool [provider] [org] [repo] <tool>
Enable, disable, or configure an analysis tool
patterns [provider] [org] [repo] <tool>
List patterns for a tool, or bulk enable/disable them
pattern [provider] [org] [repo] <tool> <id>
Show a pattern, or enable, disable, or set parameters for it
Run codacy <command> --help for full argument and option details for any command.
npm start -- < command > # Run in development mode
npm test # Run tests
npm run type-check # Type-check without emitting
npm run build # Build for production
npm run update-api # Update the auto-generated API client
CI/CD
CI : Runs on every push to main and on PRs. Builds and tests across Node.js 18, 20, and 22.
Release : Uses changesets for automated versioning and npm publishing.
When making changes, run npx changeset and describe your change (select patch , minor , or major )
Include the generated .changeset/*.md file in your PR
CI enforces that every PR includes a changeset (use npx changeset --empty for changes that don't need a version bump, like docs or CI)
When PRs are merged to main , the release workflow automatically creates a "chore: version packages" PR that bumps the version and updates CHANGELOG.md
Merging that PR publishes to npm with provenance
Prerequisite : An NPM_TOKEN secret must be configured in the GitHub repository settings.
A command-line tool to interact with Codacy Cloud directly from your terminal.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
9
v1.2.0
Latest
Jun 3, 2026
+ 8 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
