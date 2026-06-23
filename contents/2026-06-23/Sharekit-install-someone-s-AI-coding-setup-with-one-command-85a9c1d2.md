---
source: "https://github.com/LucasSantana-Dev/sharekit"
hn_url: "https://news.ycombinator.com/item?id=48647118"
title: "Sharekit – install someone's AI coding setup with one command"
article_title: "GitHub - LucasSantana-Dev/sharekit: Share your AI coding setup — CLAUDE.md, skills, cursorrules — with one command. install / preview / rollback. · GitHub"
author: "LukSantana"
captured_at: "2026-06-23T16:05:11Z"
capture_tool: "hn-digest"
hn_id: 48647118
score: 1
comments: 0
posted_at: "2026-06-23T16:03:10Z"
tags:
  - hacker-news
  - translated
---

# Sharekit – install someone's AI coding setup with one command

- HN: [48647118](https://news.ycombinator.com/item?id=48647118)
- Source: [github.com](https://github.com/LucasSantana-Dev/sharekit)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T16:03:10Z

## Translation

タイトル: Sharekit – 1 つのコマンドで誰かの AI コーディング セットアップをインストールする
記事のタイトル: GitHub - LucasSantana-Dev/sharekit: AI コーディング セットアップ (CLAUDE.md、スキル、カーソルルール) を 1 つのコマンドで共有します。インストール/プレビュー/ロールバック。 · GitHub
説明: AI コーディング設定 (CLAUDE.md、スキル、カーソルルール) を 1 つのコマンドで共有します。インストール/プレビュー/ロールバック。 - LucasSantana-Dev/シェアキット

記事本文:
GitHub - LucasSantana-Dev/sharekit: AI コーディング セットアップ (CLAUDE.md、スキル、カーソルルール) を 1 つのコマンドで共有します。インストール/プレビュー/ロールバック。 · GitHub
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
ルーカス・サンタナ・デヴ
/
シェアキット
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
25 コミット 25 コミット .claude/ backlog .claude/ backlog .github/ workflows .github/ workflows docs docs sharekit-profile/ .claude sharekit-profile/ .claude src src test test .gitignore .gitignore .prettierrc .prettierrc CHANGELOG.md CHANGELOG.mdライセンス ライセンス README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング設定 (CLAUDE.md、スキル、カーソルルール、ドットファイル) を誰とでも共有します。 1 つのコマンドでインストールし、もう 1 つでロールバックします。
npx @lucassantana/sharekit install <github-user>
github.com/<github-user>/sharekit-profile からプロファイルを取得し、変更 (色 + 数) をプレビューし、確認を求め、上書きするファイルをバックアップして適用します。 sharekit rollback <github-user> を使用して元に戻します。
npx @lucassantana/sharekitreview < github-user > # 何が変更されるかを確認し、何も適用しない
npx @lucassantana/sharekit rollback < github-user > # 最後のバックアップを復元します
プロフィールを発見する
GitHub はレジストリです。公開されたプロファイル ( sharekit-profile という名前のリポジトリ) を見つけます。
npx @lucassantana/sharekit search # 公開されているすべてのプロファイルをリストする
npx @lucassantana/sharekit 検索反応 # キーワードでフィルターする
各結果には、インストールするためのワンライナーが示されています。
最新のプロファイルの代わりに、プロファイルの特定のタグまたはブランチをインストールします。
npx @lucassantana/sharekit install < github-user > @v1.0 # タグ
npx @lucassantana/sharekit install < github-user > @stable # ブランチ
プレーン インストール <github-user> は常にプロファイルのデフォルト ブランチ (HEAD) を追跡します。
プロファイルの .claude/settings.json はシェル コマンドを実行するフックを定義できるため、n

デフォルトでインストールされます。プレビューでフラグが付けられ、スキップされます。レビュー後にオプトインするには、 --include-hooks を渡します (記述される前に 2 回目の明示的な確認が表示されます)。
npx @lucassantana/sharekit install < github-user > --include-hooks
自分のプロフィールを公開する
次の構造で sharekit-profile という名前の GitHub リポジトリを作成します。
シェアキットプロファイル/
§── sharekit.toml
§── クロード/ (→ ~/.claude/)
§── カーソル/ (→ ~/.cursor/)
└── 共有/ (→ ~/)
sharekit.toml の例:
【プロフィール】
名前 = " 私のセットアップ "
バージョン = " 1.0 "
description = " カスタムスキルを使用したクロード + カーソル設定 "
サブディレクトリは対応するルートにミラーリングされます。claude/ のファイルは ~/.claude/ に、cursor/ のファイルは ~/.cursor/ に、shared/ のファイルは ~/ に移動します。
sharekit init [skill...] を実行して ~/.claude からプロファイルをスキャフォールディングします。CLAUDE.md と名前付きスキルをすぐにプッシュできる sharekit-profile/ にコピーします。
プロファイルをパブリック GitHub リポジトリにプッシュする前に、sharekit スキャンを使用してシークレットを確認します。
npx @lucassantana/sharekit scan ./sharekit-profile
スキャナーは、秘密キー、AWS/GitHub/Slack/Google API トークン、ベアラー トークン、機密環境変数 ( SECRET|TOKEN|PASSWORD|API_KEY|APIKEY|ACCESS_KEY )、およびホームパス リークを検出します。重大度の高い検出結果 (キー/トークン) は、ゼロ以外の exit で init および scan コマンドをブロックします。中程度/低度の検出結果 (環境変数名、パスなど) は警告のみを表示します。
手動で調査結果を確認して編集した場合は、 --force でブロックをオーバーライドします。
npx @lucassantana/sharekit scan ./sharekit-profile --force
git Push の前に、必ずプロファイルと .gitignore の機密ファイルを確認してください。スキャナーはベストエフォート型です。あなたには、本当の秘密が漏洩しないようにする責任があります。
フックが自動的にインストールされることはありません。フックを含む設定 ( .claude/settings.json

) はプレビューでフラグが付けられ、スキップされます。確認後、手動でマージしてください。
申請する前にプレビューしてください。 sharekit プレビューには、書き込み前の正確な差分 (新規/変更/未変更の数とパス)、つまりトラスト ゲートが表示されます。
すべてがバックアップされています。適用する前に、sharekit は変更されたファイルを ~/.sharekit/backups/<user>-<timestamp>/ に保存します。ロールバックするとそれらが復元されます。
まだレジストリ/検出がありません — GitHub がレジストリです。 sharekit-profile リポジトリをトピック別に検索するか、規則に従って検索します。
ファイルコピーのみ。マルチツールのマージロジックはありません。 TOML、テキスト、バイナリ ファイルはそのままコピーされます。共有キットのプレビューを使用して競合を特定します。
AI コーディング設定 (CLAUDE.md、スキル、カーソルルール) を 1 つのコマンドで共有します。インストール/プレビュー/ロールバック。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.3.0 — プロファイルの検出
最新の
2026 年 6 月 23 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Share your AI coding setup — CLAUDE.md, skills, cursorrules — with one command. install / preview / rollback. - LucasSantana-Dev/sharekit

GitHub - LucasSantana-Dev/sharekit: Share your AI coding setup — CLAUDE.md, skills, cursorrules — with one command. install / preview / rollback. · GitHub
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
LucasSantana-Dev
/
sharekit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
25 Commits 25 Commits .claude/ backlog .claude/ backlog .github/ workflows .github/ workflows docs docs sharekit-profile/ .claude sharekit-profile/ .claude src src test test .gitignore .gitignore .prettierrc .prettierrc CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Share your AI coding setup — CLAUDE.md, skills, cursorrules, and dotfiles — with anyone. One command to install, one to rollback.
npx @lucassantana/sharekit install < github-user >
Fetches a profile from github.com/<github-user>/sharekit-profile , previews the changes (colors + counts), asks for confirmation, backs up any files it will overwrite, and applies them. Undo with sharekit rollback <github-user> .
npx @lucassantana/sharekit preview < github-user > # see what would change, apply nothing
npx @lucassantana/sharekit rollback < github-user > # restore the last backup
Discover profiles
GitHub is the registry — find published profiles (any repo named sharekit-profile ):
npx @lucassantana/sharekit search # list all published profiles
npx @lucassantana/sharekit search react # filter by keyword
Each result shows the one-liner to install it.
Install a specific tag or branch of a profile instead of the latest:
npx @lucassantana/sharekit install < github-user > @v1.0 # a tag
npx @lucassantana/sharekit install < github-user > @stable # a branch
Plain install <github-user> always tracks the profile's default branch (HEAD).
A profile's .claude/settings.json can define hooks that run shell commands, so it is never installed by default — it's flagged in the preview and skipped. To opt in after reviewing it, pass --include-hooks (you'll get a second explicit confirmation before it's written):
npx @lucassantana/sharekit install < github-user > --include-hooks
Publish your own profile
Create a GitHub repo named sharekit-profile with this structure:
sharekit-profile/
├── sharekit.toml
├── claude/ (→ ~/.claude/)
├── cursor/ (→ ~/.cursor/)
└── shared/ (→ ~/)
sharekit.toml example:
[ profile ]
name = " My Setup "
version = " 1.0 "
description = " Claude + Cursor config with custom skills "
Subdirectories mirror into their corresponding roots: files in claude/ go to ~/.claude/ , files in cursor/ to ~/.cursor/ , and files in shared/ to ~/ .
Run sharekit init [skill...] to scaffold a profile from your ~/.claude — copies your CLAUDE.md and any named skills into a ready-to-push sharekit-profile/ .
Before pushing your profile to a public GitHub repo, use sharekit scan to check for secrets:
npx @lucassantana/sharekit scan ./sharekit-profile
The scanner detects private keys, AWS/GitHub/Slack/Google API tokens, bearer tokens, sensitive environment variables ( SECRET|TOKEN|PASSWORD|API_KEY|APIKEY|ACCESS_KEY ), and home-path leaks. High-severity findings (keys/tokens) block the init and scan commands with a non-zero exit . Medium/low findings (e.g., env-var names, paths) warn only.
Override a block with --force if you've manually reviewed and redacted the findings:
npx @lucassantana/sharekit scan ./sharekit-profile --force
Always review the profile and .gitignore sensitive files before git push . The scanner is best-effort; you are responsible for ensuring no real secrets escape.
Hooks are never auto-installed. Settings containing hooks ( .claude/settings.json ) are flagged in the preview and skipped. Merge them manually after reviewing.
Preview before applying. sharekit preview shows the exact diff (new/changed/unchanged counts and paths) — trust gate before any write.
Everything is backed up. Before applying, sharekit saves changed files to ~/.sharekit/backups/<user>-<timestamp>/ . Rollback restores them.
No registry/discovery yet — GitHub is the registry. Search sharekit-profile repos by topic or follow the convention.
File-copy only; no multi-tool merge logic. TOML, text, and binary files are copied as-is. Use sharekit preview to spot conflicts.
Share your AI coding setup — CLAUDE.md, skills, cursorrules — with one command. install / preview / rollback.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.3.0 — discover profiles
Latest
Jun 23, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
