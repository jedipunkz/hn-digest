---
source: "https://github.com/malvads/Slopper"
hn_url: "https://news.ycombinator.com/item?id=48426225"
title: "Slopper GitHub Action: Fighting AI Slop Contributions on Open Source Projects"
article_title: "GitHub - malvads/Slopper: Keep AI slop out of your pull requests. · GitHub"
author: "malvads"
captured_at: "2026-06-06T18:33:02Z"
capture_tool: "hn-digest"
hn_id: 48426225
score: 3
comments: 0
posted_at: "2026-06-06T15:58:07Z"
tags:
  - hacker-news
  - translated
---

# Slopper GitHub Action: Fighting AI Slop Contributions on Open Source Projects

- HN: [48426225](https://news.ycombinator.com/item?id=48426225)
- Source: [github.com](https://github.com/malvads/Slopper)
- Score: 3
- Comments: 0
- Posted: 2026-06-06T15:58:07Z

## Translation

タイトル: Slopper GitHub アクション: オープンソース プロジェクトにおける AI のスロップ貢献との戦い
記事のタイトル: GitHub - malvads/Slopper: AI スロップをプル リクエストから除外します。 · GitHub
説明: プル リクエストから AI スロップを排除します。 GitHub でアカウントを作成して、malvads/Slopper の開発に貢献してください。

記事本文:
GitHub - malvads/Slopper: プル リクエストから AI スロップを排除します。 · GitHub
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
マルバッド
/
だらしない人
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用します

ect このアクションを既存のワークフローに追加するか、新しいワークフローを作成します マーケットプレイスで表示する メイン ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット .claude .claude .github/ workflows .github/ workflows __tests__ __tests__ アセット アセット バッジ バッジ dist dist 例 例 src src .gitignore .gitignore .slopper_risky_users .slopper_risky_users CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md action.yml action.yml jest.config.js jest.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
非常に実験的です。荒削りな部分、誤検知、重大な変更が予想されます。
オープンソースのメンテナーは、見た目はきれいでも何も追加しない、AI によって生成されたプル リクエストに溺れています。カール、Linux カーネル、Godot、Node.js はすべて攻撃を受けています。洗練された説明、CI 合格、実質価値ゼロ。
Slopper は、決定論的ヒューリスティックを使用してすべての PR をスコアリングして、1 つの質問に答える GitHub アクションです。この PR は実際に価値をもたらしますか? API キーやコストは不要で、ワークフロー ファイルは 1 つです。
存在しないバグのファントム修正
クリティカル パスに触れる不必要なリファクタリング
明白なことを再説明するドキュメント
数十の無関係なリポジトリに送信するスプレーアンドプレイアカウント
重要インフラにおけるレピュテーションファーミング
スコア PR 0 ～ 10 — ヒューリスティック シグナルからの決定論的なリスク スコアなので、どこにレビュー時間を費やすべきかがわかります
6 つのヒューリスティック シグナル (コメント密度、スロップ語彙、冗長識別子、ドキュメント文字列の肥大化、ボイラープレート比率、構造パターン) を使用して AI 生成コードを検出します。
GitHub 全体の投稿者をプロファイル (アカウントの年齢、PR ボリューム、マージ率、スプレー スコア) して、全期間にわたって低品質の PR をショットガンしているアカウントを特定します。

数十のリポジトリ
コミュニティが管理するブロックリストから危険なアカウントにフラグを立て、リアルタイムで更新されます
構成可能なしきい値に基づいて自動終了、自動承認、またはレビューのリクエストを行う
常習犯の禁止 — メンテナーはアカウントを永久にブロックするコメント/怠け者の報告を行う
信頼できる寄稿者を保証します — /slopper vouch は、今後その著者のすべての分析をスキップします
AI はオプション — デフォルトでは完全に決定論的ヒューリスティックに基づいて実行されます。必要に応じて、より詳細な分析のために AI プロバイダー ( openai 、 anthropic 、 vertex 、 groq 、 gemini ) を追加します。
名前 : スラッパー
に:
プルリクエスト:
タイプ: [オープン、同期、再オープン]
問題のコメント:
タイプ: [作成済み]
ジョブ:
分析します:
実行: ubuntu-最新
権限:
内容：書く
プルリクエスト : 書き込み
手順:
- 使用: malvads/slopper@v1
付き:
github トークン: ${{ Secrets.GITHUB_TOKEN }}
それだけです。 API キーは必要ありません — Slopper は、すぐに使える決定論的ヒューリスティックを使用して PR をスコアリングします。
AI を活用したより詳細な分析を行うには、プロバイダーを追加します。
- 使用: malvads/slopper@v1
付き:
aiプロバイダー:「ジェミニ」
gemini-API-key : ${{ Secrets.GEMINI_API_KEY }}
github トークン: ${{ Secrets.GITHUB_TOKEN }}
その他のセットアップ (厳密モード、マージ ゲーティング、マルチプロバイダー) については、examples/ を参照してください。
決定論的モード (デフォルト) では、リスク スコアは次から導出されます。
AI プロバイダーが構成されている場合、スコアは代わりに AI 分析から得られ、コミット メッセージ、コードの差分、および動作信号からのより豊富なコンテキストが含まれます。
リポジトリのルートに .slopper ファイルを作成します。すべてのフィールドはオプションです。
保証された：
- 信頼できる貢献者
- 依存ボット[ボット]
禁止:
- 既知のスロップアカウント
アクション:
auto_close :
有効 : true
しきい値 : 9
auto_approve :
有効 : false
しきい値 : 2
auto_request_review :
有効 : true
しきい値 : 6
査読者：
- シニアメンテナ
しきい値:
低い：2
中：5
高：8
ラベル_t

しきい値:
aiの可能性: 70
ai_おそらく : 40
スプレースコア : 60
new_account_days : 30
activity_burst_prs : 10
activity_burst_days : 7
マージ比率サスペクト : 0.4
security_review_score : 6
疑わしいスコア : 8
スプレー重み :
リポジトリ: 40
ボリューム：30
マージ率 : 20
アカウント年齢 : 10
無視するパス:
- " *.md "
- " docs/** "
- 「 *.lock 」
ルール：
require_description : true
require_linked_issue : false
max_files_changed : 0
コマンド
すべてのラベルは決定的であり、AI がそれらを選択することはありません。
後続のワークフロー ステップで使用します。
プル リクエストから AI のスロップを排除します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.0.1: マーケットプレイス公開のアクション名と説明を修正
最新の
2026 年 6 月 6 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Keep AI slop out of your pull requests. Contribute to malvads/Slopper development by creating an account on GitHub.

GitHub - malvads/Slopper: Keep AI slop out of your pull requests. · GitHub
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
malvads
/
Slopper
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits .claude .claude .github/ workflows .github/ workflows __tests__ __tests__ assets assets badges badges dist dist examples examples src src .gitignore .gitignore .slopper_risky_users .slopper_risky_users CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md action.yml action.yml jest.config.js jest.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Very experimental — expect rough edges, false positives, and breaking changes.
Open source maintainers are drowning in AI-generated pull requests that look clean but add nothing. curl, the Linux kernel, Godot, Node.js — they've all been hit. Polished descriptions, passing CI, and zero real value.
Slopper is a GitHub Action that scores every PR using deterministic heuristics to answer one question: does this PR actually add value? No API keys, no cost, one workflow file.
Phantom fixes for bugs that don't exist
Unnecessary refactoring that touches critical paths
Documentation that restates the obvious
Spray-and-pray accounts submitting to dozens of unrelated repos
Reputation farming in critical infrastructure
Scores PRs 0–10 — deterministic risk score from heuristic signals, so you know where to spend review time
Detects AI-generated code using 6 heuristic signals (comment density, slop vocabulary, verbose identifiers, docstring bloat, boilerplate ratio, structural patterns)
Profiles contributors across GitHub — account age, PR volume, merge ratio, spray score — to spot accounts that shotgun low-quality PRs across dozens of repos
Flags risky accounts from a community-maintained blocklist , updated in real time
Auto-closes, auto-approves, or requests review based on configurable thresholds
Bans repeat offenders — maintainers comment /slopper report to permanently block an account
Vouches trusted contributors — /slopper vouch skips all analysis for that author going forward
AI optional — runs fully on deterministic heuristics by default. Add an AI provider ( openai , anthropic , vertex , groq , gemini ) for deeper analysis when you want it
name : Slopper
on :
pull_request :
types : [opened, synchronize, reopened]
issue_comment :
types : [created]
jobs :
analyze :
runs-on : ubuntu-latest
permissions :
contents : write
pull-requests : write
steps :
- uses : malvads/slopper@v1
with :
github-token : ${{ secrets.GITHUB_TOKEN }}
That's it. No API keys needed — Slopper scores PRs using deterministic heuristics out of the box.
For deeper AI-powered analysis, add a provider:
- uses : malvads/slopper@v1
with :
ai-provider : ' gemini '
gemini-api-key : ${{ secrets.GEMINI_API_KEY }}
github-token : ${{ secrets.GITHUB_TOKEN }}
See examples/ for more setups (strict mode, merge gating, multi-provider).
In deterministic mode (default), the risk score is derived from:
When an AI provider is configured, the score comes from the AI analysis instead, with richer context from commit messages, code diffs, and behavioral signals.
Create a .slopper file in your repo root. Every field is optional:
vouched :
- trusted-contributor
- dependabot[bot]
banned :
- known-slop-account
actions :
auto_close :
enabled : true
threshold : 9
auto_approve :
enabled : false
threshold : 2
auto_request_review :
enabled : true
threshold : 6
reviewers :
- senior-maintainer
thresholds :
low : 2
medium : 5
high : 8
label_thresholds :
ai_likely : 70
ai_possibly : 40
spray_score : 60
new_account_days : 30
activity_burst_prs : 10
activity_burst_days : 7
merge_ratio_suspect : 0.4
security_review_score : 6
suspicious_score : 8
spray_weights :
repos : 40
volume : 30
merge_ratio : 20
account_age : 10
ignore_paths :
- " *.md "
- " docs/** "
- " *.lock "
rules :
require_description : true
require_linked_issue : false
max_files_changed : 0
Commands
All labels are deterministic — the AI never picks them.
Use in subsequent workflow steps:
Keep AI slop out of your pull requests.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.0.1: Fix action name and description for Marketplace publishing
Latest
Jun 6, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
