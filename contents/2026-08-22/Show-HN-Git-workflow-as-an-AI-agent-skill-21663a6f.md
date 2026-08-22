---
source: "https://github.com/musoyangrigor/gitx-skill"
hn_url: "https://news.ycombinator.com/item?id=49401487"
title: "Show HN: Git workflow as an AI-agent skill"
article_title: "GitHub - musoyangrigor/gitx-skill: GitX is a portable AI-agent skill for creating clean Git commits, tagged branches, and safe pushes across Codex, Claude Code, Cursor, and other skills-compatible agents. · GitHub"
image: "https://opengraph.githubassets.com/83ef055a8724c1e2afe33ad4a765d3bfae43daa8e5ff57cec9bc8bcad5a34cf8/musoyangrigor/gitx-skill"
author: "MusoyanGrigor"
captured_at: "2026-08-22T17:12:18Z"
capture_tool: "hn-digest"
hn_id: 49401487
score: 1
comments: 0
posted_at: "2026-08-22T16:50:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Git workflow as an AI-agent skill

- HN: [49401487](https://news.ycombinator.com/item?id=49401487)
- Source: [github.com](https://github.com/musoyangrigor/gitx-skill)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T16:50:30Z

## Translation

タイトル: HN の表示: AI エージェント スキルとしての Git ワークフロー
記事タイトル: GitHub - musoyangrigor/gitx-skill: GitX は、クリーンな Git コミット、タグ付きブランチ、Codex、Claude Code、Cursor、およびその他のスキル互換エージェント間での安全なプッシュを作成するためのポータブル AI エージェント スキルです。 · GitHub
説明: GitX は、クリーンな Git コミット、タグ付きブランチ、Codex、Claude Code、Cursor、およびその他のスキル互換エージェント間での安全なプッシュを作成するためのポータブル AI エージェント スキルです。 - musoyangrigor/gitx-skill

記事本文:
GitHub - musoyangrigor/gitx-skill: GitX は、クリーンな Git コミット、タグ付きブランチ、Codex、Claude Code、Cursor、およびその他のスキル互換エージェント間での安全なプッシュを作成するためのポータブル AI エージェント スキルです。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ムソヤンリゴール
/
gitx スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
39 コミット 39 コミット フォルダーとファイル
gitx gitx ライセンス ライセンス README.md README.md すべて表示

iles リポジトリ ファイルのナビゲーション
AI によって生成された乱雑な変更を、クリーンで安全な Git 履歴に変換します。
GitX は、AI コーディング エージェント向けの移植可能な Git ワークフロー スキルであり、ワーキング ツリーの変更を論理的な従来のコミットに変換し、ブランチ、プロジェクト チェック、安全なプルおよびプッシュ ワークフロー、GitHub プル リクエストと問題、シークレット スキャン、マージとリベースの競合解決、リポジトリのステータスと履歴、およびエージェント スキル互換ツール全体でのコミット計画を処理します。
npx スキル追加 musoyangrigor/gitx-skill --skill gitx
スキル CLI は、選択したサポートされている AI エージェントのスキルを構成します。インストール後に新しいエージェント セッションを開始します。
裸の $gitx 呼び出しはすぐにリポジトリを検査し、スマート コミットを実行します。 GitX は、助けを求めた場合にのみコマンド リストを表示します。
コマンド
説明
$gitx
変更を検査し、スマートコミットを作成します。
$gitx本体
便利なコミット本文を含むスマート コミットを作成します。
$gitx ブランチ
feat/ や fix/ などの推論されたプレフィックスを持つブランチを作成して切り替えます。
$gitx ブランチ修正/トークンリフレッシュ
名前付きブランチを作成して切り替えます。
$gitx ブランチチェック
推論されたブランチを作成し、チェックを実行してからコミットします。
$gitx プル
現在のブランチの更新を安全にプルします。
$gitx プッシュ
現在のブランチをoriginにプッシュします。必要に応じて上流を作成します。
$gitx pr
生成されたタイトルと本文を使用して、 GitHub プル リクエストを Origin のデフォルト ブランチに作成します。
$gitx 開発前
現在のブランチから build への GitHub プル リクエストを作成します。
$gitx の問題 トークンの有効期限が切れるとログインが失敗する
生成されたタイトルと本文を使用して GitHub の問題を作成します。
$gitx 問題 123
GitHub の問題 #123 を読み、要求された修正を現在の作業ツリーに実装します。
$gitx 解決
進行中のマージまたはリベースの競合を解決します。
$gitx チェック
関連するチェックを実行してから、スマート コミットを作成します。
$gitx ステータス
ショー担当者

何も変わらずオシトリー状態。
$gitx ツリー
コンパクトな Git 履歴ツリー、ブランチ、同期、PR、および作業ツリー情報を表示します。
$gitx スキャン
変更と Git 履歴をスキャンして、公開されたシークレットと機密ファイルを探します。何も変更せずに、プロジェクトのセキュリティ状態、調査結果、肯定的な点、推奨されるアクション、および機密漏洩の評価を要約します。
$gitx プラン
何も変更せずにコミット グループとメッセージをプレビューします。
$gitx タイプの修正
修正タイプでスマート コミットを作成します。
$gitx スコープ認証
認証スコープを使用してスマート コミットを作成します。
$gitx ファイル README.md package.json
指定したファイルのみをコミットします。
$gitx 修正
最新のコミットを修正する前に質問してください。
GitX が複数の論理コミット グループを見つけた場合、実際の数のコミットを作成するか 1 つのコミットを作成するかを尋ねます。
gitx issue の場合、123 や #123 などの数字のみが GitX に既存の問題を実装するように指示します。 「修正」や「更新」などの単語を使用した説明など、その他のテキストは新しい GitHub 問題を作成します。
AI エージェントのスキル インターフェイスを通じて GitX を呼び出し、同じコマンド ワードを使用します。例: gitx plan 、 gitx check 、または gitx Branch fix/token-refresh 。
GitX は、移植可能な SKILL.md エージェント スキル形式に従っています。
GitX は、クリーンな Git コミット、タグ付きブランチ、Codex、Claude Code、Cursor、およびその他のスキル互換エージェント間での安全なプッシュを作成するためのポータブル AI エージェント スキルです。
Readme MIT ライセンス アクティビティ スター
5 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

GitX is a portable AI-agent skill for creating clean Git commits, tagged branches, and safe pushes across Codex, Claude Code, Cursor, and other skills-compatible agents. - musoyangrigor/gitx-skill

GitHub - musoyangrigor/gitx-skill: GitX is a portable AI-agent skill for creating clean Git commits, tagged branches, and safe pushes across Codex, Claude Code, Cursor, and other skills-compatible agents. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
musoyangrigor
/
gitx-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
39 Commits 39 Commits Folders and files
gitx gitx LICENSE LICENSE README.md README.md View all files Repository files navigation
Turn messy AI-generated changes into clean, safe Git history.
GitX is a portable Git workflow skill for AI coding agents that turns working-tree changes into logical Conventional Commits and handles branches, project checks, safe pull and push workflows, GitHub pull requests and issues, secret scanning, merge and rebase conflict resolution, repository status and history, and commit planning across Agent Skills-compatible tools.
npx skills add musoyangrigor/gitx-skill --skill gitx
The Skills CLI configures the skill for the selected supported AI agent. Start a new agent session after installation.
A bare $gitx invocation immediately inspects the repository and runs Smart commit. GitX shows its command list only when you ask for help.
Command
Description
$gitx
Inspect changes and create a smart commit.
$gitx body
Create a smart commit with a useful commit body.
$gitx branch
Create and switch to a branch with an inferred prefix, such as feat/ or fix/ .
$gitx branch fix/token-refresh
Create and switch to the named branch.
$gitx branch check
Create an inferred branch, run checks, then commit.
$gitx pull
Safely pull updates for the current branch.
$gitx push
Push the current branch to origin ; create its upstream if needed.
$gitx pr
Create a GitHub pull request into origin 's default branch with a generated title and body.
$gitx pr develop
Create a GitHub pull request from the current branch into develop .
$gitx issue Login fails after token expiry
Create a GitHub issue with a generated title and body.
$gitx issue 123
Read GitHub issue #123 and implement the requested fix in the current working tree.
$gitx resolve
Resolve an in-progress merge or rebase conflict.
$gitx check
Run relevant checks, then create a smart commit.
$gitx status
Show repository status without changing anything.
$gitx tree
Show a compact Git history tree, branch, sync, PR, and working-tree information.
$gitx scan
Scan changes and Git history for exposed secrets and sensitive files; summarize the project’s security state, findings, positives, recommended actions, and a secret-exposure rating without modifying anything.
$gitx plan
Preview commit groups and messages without changing anything.
$gitx type fix
Create a smart commit with the fix type.
$gitx scope auth
Create a smart commit with the auth scope.
$gitx files README.md package.json
Commit only the specified files.
$gitx amend
Ask before amending the most recent commit.
If GitX finds several logical commit groups, it asks whether to create the real number of commits or one commit.
For gitx issue , only a number such as 123 or #123 tells GitX to implement an existing issue. Any other text creates a new GitHub issue, including descriptions that use words such as “fix” or “update.”
Invoke GitX through your AI agent's skill interface, then use the same command words. For example: gitx plan , gitx check , or gitx branch fix/token-refresh .
GitX follows the portable SKILL.md Agent Skills format.
GitX is a portable AI-agent skill for creating clean Git commits, tagged branches, and safe pushes across Codex, Claude Code, Cursor, and other skills-compatible agents.
Readme MIT license Activity Stars
5 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
