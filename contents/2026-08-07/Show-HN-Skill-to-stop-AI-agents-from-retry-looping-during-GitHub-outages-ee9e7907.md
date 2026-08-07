---
source: "https://github.com/tbille/check-github-status"
hn_url: "https://news.ycombinator.com/item?id=49206945"
title: "Show HN: Skill to stop AI agents from retry-looping during GitHub outages"
article_title: "GitHub - tbille/check-github-status: Agent Skill: checks githubstatus.com before an AI agent retries a failed GitHub action, to tell real outages apart from local bugs. · GitHub"
author: "totostache"
captured_at: "2026-08-07T08:07:08Z"
capture_tool: "hn-digest"
hn_id: 49206945
score: 1
comments: 0
posted_at: "2026-08-07T07:14:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Skill to stop AI agents from retry-looping during GitHub outages

- HN: [49206945](https://news.ycombinator.com/item?id=49206945)
- Source: [github.com](https://github.com/tbille/check-github-status)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T07:14:35Z

## Translation

タイトル: HN を表示: GitHub の停止中に AI エージェントの再試行ループを停止するスキル
記事のタイトル: GitHub - tbille/check-github-status: エージェント スキル: AI エージェントが失敗した GitHub アクションを再試行する前に githubstatus.com をチェックして、ローカルのバグと実際の停止を区別します。 · GitHub
説明: エージェント スキル: AI エージェントが失敗した GitHub アクションを再試行する前に githubstatus.com をチェックして、ローカルのバグと実際の停止を区別します。 - tbille/check-github-status
HN テキスト: GitHub がダウンすることは今やエンジニアリングの一部です。私はエージェントが失敗時に再試行ループの代わりに GitHub のステータスをチェックできるスキルを作成しました。

記事本文:
GitHub - tbille/check-github-status: エージェント スキル: AI エージェントが失敗した GitHub アクションを再試行する前に githubstatus.com をチェックして、ローカルのバグと実際の停止を区別します。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
トゥビル
/
チェック-github-status
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows ライセンス ライセンス README.md README.md SKILL.md SKILL.md すべて表示

ファイル リポジトリ ファイルのナビゲーション
AIコーディングエージェントのループを阻止するエージェントスキル
GitHub 関連のコマンドが失敗した場合に再試行します — チェックすることで
まず githubstatus.com で障害が発生しているかどうかを確認します。
実際には GitHub のせいです。
GitHub の停止 (アクションの実行停止、PR の同期/マージの失敗、API 5xx エラー)
多くの場合、地元のバグのように見えます。 git プッシュを再試行し続けるエージェント、gh
コマンドを実行したり、アクションを再実行したりすると、問題を追跡するのに多くの時間が費やされる可能性があります。
そこに。このスキルは、2 回目の再試行の前に GitHub のステータス API をチェックし、
エージェント (およびあなた) は、ローカルでデバッグを続けるか、停止が起こるまで待つかどうかを決定します。
npx スキルは tbille/check-github-status を追加します
何をするのか
GitHub 関連の障害 ( git Push/pull/fetch/clone、gh CLI エラー、
アクションの失敗、マージされない PR、GitHub API エラーなど）、次のように実行されます。
カール -s https://www.githubstatus.com/api/v2/summary.json
そして、その応答を使用して次のことを決定します。
GitHub の機能が低下している → 再試行を停止し、影響を受けるコンポーネント/インシデントを報告します。
待つことを提案します。
GitHub は大丈夫です → そう言って、ローカルな原因 (認証、レート制限、
ワークフロー構文など)。
完全な手順については、SKILL.md を参照してください。
エージェント スキル: AI エージェントが失敗した GitHub アクションを再試行する前に githubstatus.com をチェックし、ローカルのバグと実際の停止を区別します。
Readme MIT ライセンス アクティビティ スター
フォーク数 0 レポート リポジトリの寄稿者
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agent Skill: checks githubstatus.com before an AI agent retries a failed GitHub action, to tell real outages apart from local bugs. - tbille/check-github-status

GitHub going down is part of engineering now, I made a skill that lets your agent check GitHub's status on failure instead of retry-looping.

GitHub - tbille/check-github-status: Agent Skill: checks githubstatus.com before an AI agent retries a failed GitHub action, to tell real outages apart from local bugs. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
tbille
/
check-github-status
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows LICENSE LICENSE README.md README.md SKILL.md SKILL.md View all files Repository files navigation
An Agent Skill that stops AI coding agents from looping on
retries when a GitHub-related command fails — by checking
githubstatus.com first to see whether the failure
is actually GitHub's fault.
GitHub outages (Actions runners down, PRs failing to sync/merge, API 5xx errors)
often look like local bugs. An agent that just keeps retrying a git push , gh
command, or Actions re-run can burn a lot of time chasing a problem that isn't
there. This skill checks GitHub's status API before a second retry, and tells the
agent (and you) whether to keep debugging locally or just wait out an outage.
npx skills add tbille/check-github-status
What it does
On any GitHub-related failure ( git push/pull/fetch/clone , gh CLI errors,
Actions failures, PRs that won't merge, GitHub API errors), it runs:
curl -s https://www.githubstatus.com/api/v2/summary.json
and uses the response to decide:
GitHub is degraded → stop retrying, report the affected component/incident,
suggest waiting.
GitHub is fine → say so, and move on to local causes (auth, rate limits,
workflow syntax, etc.).
See SKILL.md for the full instructions.
Agent Skill: checks githubstatus.com before an AI agent retries a failed GitHub action, to tell real outages apart from local bugs.
Readme MIT license Activity Stars
0 forks Report repository Contributors
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
