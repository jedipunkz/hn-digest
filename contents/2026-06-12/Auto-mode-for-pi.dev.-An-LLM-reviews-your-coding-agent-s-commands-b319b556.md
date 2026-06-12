---
source: "https://github.com/vinzenzu/pi-auto-reviewer"
hn_url: "https://news.ycombinator.com/item?id=48501272"
title: "Auto mode for pi.dev. An LLM reviews your coding agent's commands"
article_title: "GitHub - vinzenzu/pi-auto-reviewer: Auto-review commands your Pi agent (pi.dev) wants to execute · GitHub"
author: "vinzenzu"
captured_at: "2026-06-12T10:35:57Z"
capture_tool: "hn-digest"
hn_id: 48501272
score: 1
comments: 2
posted_at: "2026-06-12T08:08:56Z"
tags:
  - hacker-news
  - translated
---

# Auto mode for pi.dev. An LLM reviews your coding agent's commands

- HN: [48501272](https://news.ycombinator.com/item?id=48501272)
- Source: [github.com](https://github.com/vinzenzu/pi-auto-reviewer)
- Score: 1
- Comments: 2
- Posted: 2026-06-12T08:08:56Z

## Translation

タイトル: pi.dev の自動モード。 LLM はコーディング エージェントのコマンドをレビューします
記事のタイトル: GitHub - vinzenzu/pi-auto-reviewer: Pi エージェント (pi.dev) が実行したいコマンドを自動レビューする · GitHub
説明: Pi エージェント (pi.dev) が実行したい自動レビュー コマンド - vinzenzu/pi-auto-reviewer

記事本文:
GitHub - vinzenzu/pi-auto-reviewer: Pi エージェント (pi.dev) が実行したいコマンドを自動レビューする · GitHub
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
ヴィンゼンズ
/
pi-auto-reviewer
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
C

頌歌
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット ライセンス ライセンス README.md README.md auto-reviewer.ts auto-reviewer.ts package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
pi エージェントが実行する bash コマンドを自動的にレビューします。Codex の「自動レビュー」や Claude Code の「自動モード」に似ています。
エージェントが実行するすべての bash コマンドは、次の 3 つの層を通過します。
コマンドが Tier 3 に分類される場合、サブエージェント LLM はプロジェクト コンテキストを使用してコマンドをレビューし、許可またはブロックを決定します。
cp auto-reviewer.ts ~ /.pi/agent/extensions/
npm経由
pi インストール npm:pi-auto-reviewer
単一プロジェクト
拡張機能をプロジェクトにコピーします。
cp auto-reviewer.ts .pi/extensions/
プロジェクトが信頼されている場合、Pi は .pi/extensions/ 内の拡張子を自動検出します。
pi -e ./auto-reviewer.ts
使用法
インストールすると自動的に動作するため、設定は必要ありません。エージェントが実行しようとするすべての bash コマンドがレビューされます。
安全なコマンド (層 1) は目に見える遅延なく実行されます。
危険なコマンド (Tier 2) はブロックされ、その理由を説明する通知が表示されます。
レビュー担当者 LLM が決定する間、他のすべて (Tier 3) は一時停止します。 「Reviewing: <command>...」というステータス メッセージが表示されます。
許可されている場合: コマンドが実行され、自動レビューアーが表示されます: ✓ <reason>
ブロックされた場合: コマンドは拒否され、自動レビューア: ✗ <reason> が表示されます。
レビューアが失敗した場合 (タイムアウト、エラー): 対話的に手動で許可または拒否するよう求められます。
印刷モード ( pi -p ) または JSON モードでは、フォールバックする UI がないため、Tier 3 コマンドはデフォルトでブロックされます。
auto-reviewer.ts の AUTO_PERMITTED 配列と AUTO_BLOCKED 配列を編集して、パターンを追加または削除します。 buildReviewPrompt() を編集して、レビュー担当者 LLM の決定方法を変更します。
Pi エージェント (pi.dev) が実行したいコマンドを自動レビューします
そこw

ロード中のエラーとして。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Auto-review commands your Pi agent (pi.dev) wants to execute - vinzenzu/pi-auto-reviewer

GitHub - vinzenzu/pi-auto-reviewer: Auto-review commands your Pi agent (pi.dev) wants to execute · GitHub
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
vinzenzu
/
pi-auto-reviewer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits LICENSE LICENSE README.md README.md auto-reviewer.ts auto-reviewer.ts package.json package.json View all files Repository files navigation
Automatically review bash commands that your pi agent wants to execute - akin to Codex "Auto-review" and Claude Code "auto mode".
Every bash command the agent wants to run goes through three tiers:
When a command falls into Tier 3 , a subagent LLM reviews the command with project context and decides ALLOW or BLOCK.
cp auto-reviewer.ts ~ /.pi/agent/extensions/
Via npm
pi install npm:pi-auto-reviewer
Single project
Copy the extension into your project:
cp auto-reviewer.ts .pi/extensions/
Pi auto-discovers extensions in .pi/extensions/ when the project is trusted.
pi -e ./auto-reviewer.ts
Usage
Once installed, it works automatically - no configuration needed. Every bash command the agent tries to run will be reviewed.
Safe commands (Tier 1) run without any visible delay.
Dangerous commands (Tier 2) are blocked with a notification explaining why.
Everything else (Tier 3) pauses briefly while the reviewer LLM decides. You'll see a status message: Reviewing: <command>...
If allowed : the command runs and you see Auto-reviewer: ✓ <reason>
If blocked : the command is refused and you see Auto-reviewer: ✗ <reason>
If the reviewer fails (timeout, error): you're prompted interactively to allow or deny manually.
In print mode ( pi -p ) or JSON mode, Tier 3 commands are blocked by default since there's no UI to fall back on.
Edit AUTO_PERMITTED and AUTO_BLOCKED arrays in auto-reviewer.ts to add or remove patterns. Edit buildReviewPrompt() to change how the reviewer LLM decides.
Auto-review commands your Pi agent (pi.dev) wants to execute
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
