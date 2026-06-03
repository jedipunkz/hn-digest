---
source: "https://github.com/Farfield-Dev/deep-review"
hn_url: "https://news.ycombinator.com/item?id=48371737"
title: "Show HN: Claude Code plugin for deep multi-agent code reviews"
article_title: "GitHub - Farfield-Dev/deep-review: The bug-finding recipe Farfield uses in production, packaged as a Claude Code skill. Free + MIT. · GitHub"
author: "hemezh"
captured_at: "2026-06-03T00:42:12Z"
capture_tool: "hn-digest"
hn_id: 48371737
score: 1
comments: 0
posted_at: "2026-06-02T15:42:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude Code plugin for deep multi-agent code reviews

- HN: [48371737](https://news.ycombinator.com/item?id=48371737)
- Source: [github.com](https://github.com/Farfield-Dev/deep-review)
- Score: 1
- Comments: 0
- Posted: 2026-06-02T15:42:31Z

## Translation

タイトル: HN を表示: マルチエージェント コードを詳細にレビューするための Claude Code プラグイン
記事のタイトル: GitHub - Farfield-Dev/deep-review: Farfield が運用環境で使用するバグ発見レシピ。クロード コード スキルとしてパッケージ化されています。無料 + MIT。 · GitHub
説明: Farfield が運用環境で使用するバグ発見レシピ。クロード コード スキルとしてパッケージ化されています。無料 + MIT。 - Farfield-開発/詳細レビュー

記事本文:
GitHub - Farfield-Dev/deep-review: Farfield が運用環境で使用するバグ発見レシピ。クロード コード スキルとしてパッケージ化されています。無料 + MIT。 · GitHub
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
ファーフィールド開発者
/
ディープレビュー
公共
通知
通知設定を変更するにはサインインする必要があります

GS
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット .claude-plugin .claude-plugin plugins/ deep-review plugins/ deep-review .gitignore .gitignore LICENSE LICENSE MAINTAINING.md MAINTAINING.md README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
リポジトリの上級エンジニア レベルのコード レビューを行う Claude Code プラグイン。 5 つのフェーズ、並列サブエージェント、敵対的検証。所見.md を書き込みます。
/プラグイン マーケットプレイス追加 Farfield-Dev/deep-review
/プラグインのインストール deep-review@deep-review
/reload-プラグイン
( /reload-plugins は、現在のセッションに新しくインストールされたスキルを登録します。これをスキップすると、「不明なコマンド」で /deep-review:run エラーが発生します。)
このリポジトリの main から自動更新されます。
レビューしたいリポジトリ内:
/ディープレビュー:実行
--fast は、チームの意図と製品スキャンのフェーズをスキップします (より安く、より速く、より狭い)。
出力は ./.deep-review/ の下に中間成果物とともに ./findings.md に配置されます (gitignore を使用します)。
各検出結果には、重大度、影響カテゴリ (REVENUE_LEAK、SUPPORT_BURDEN、DATA_LOSS、SECURITY_BREACH、COMPLIANCE、PROD_INCIDENT、BRAND_DAMAGE)、ファイル + 行、1 段落の根本原因、検証アーティファクト、および修正方向があります。影響カテゴリのいずれにもマッピングされない調査結果は削除され、すべての調査結果はレポートに掲載される前に 100% の信頼度の敵対的チェックに合格します。中程度のバックエンドでは、50 ではなく 0 ～ 10 の結果が期待されます。
独自の Anthropic キーの場合、10 ～ 50,000 LOC バックエンドでのコールド レビューあたり 5 ～ 25 ドル。安価なフェーズでは Sonnet (アーキテクチャ マップ、チームの意図)、重いフェーズでは Opus (アクション トレース、検証)。 --高速はその約 3 分の 1 です。
コールドスタート。実行間のメモリがありません。
実際のワークフローを備えたバックエンドに最適です。静的サイトと純粋な UI リポジトリは表面に現れません

h.
意見があり、網羅的ではありません。 30 個のおそらくバグよりも 3 個の鉄壁の発見を好みます。
Farfield が運用環境で使用するバグ発見レシピ。クロード コード スキルとしてパッケージ化されています。無料 + MIT。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

The bug-finding recipe Farfield uses in production, packaged as a Claude Code skill. Free + MIT. - Farfield-Dev/deep-review

GitHub - Farfield-Dev/deep-review: The bug-finding recipe Farfield uses in production, packaged as a Claude Code skill. Free + MIT. · GitHub
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
Farfield-Dev
/
deep-review
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .claude-plugin .claude-plugin plugins/ deep-review plugins/ deep-review .gitignore .gitignore LICENSE LICENSE MAINTAINING.md MAINTAINING.md README.md README.md View all files Repository files navigation
A Claude Code plugin that does a senior-engineer-grade code review of a repo. Five phases, parallel sub-agents, adversarial validation. Writes findings.md .
/plugin marketplace add Farfield-Dev/deep-review
/plugin install deep-review@deep-review
/reload-plugins
( /reload-plugins registers the freshly installed skill in the current session. Skip it and /deep-review:run errors with "Unknown command".)
Auto-updates from this repo's main .
In the repo you want reviewed:
/deep-review:run
--fast skips the team-intent and product-scan phases (cheaper, faster, narrower).
Output lands in ./findings.md with intermediate artifacts under ./.deep-review/ (gitignore it).
Each finding has a severity, an impact category (REVENUE_LEAK, SUPPORT_BURDEN, DATA_LOSS, SECURITY_BREACH, COMPLIANCE, PROD_INCIDENT, BRAND_DAMAGE), a file + line, a one-paragraph root cause, a verification artifact, and a fix direction. Findings that don't map to one of the impact categories get dropped, and every finding passes a 100%-confidence adversarial check before it lands in the report. Expect 0–10 findings on a medium backend, not 50.
On your own Anthropic key, $5–$25 per cold review on a 10–50k LOC backend. Sonnet on the cheap phases (architecture map, team intent), Opus on the heavy ones (action traces, validation). --fast is about a third of that.
Cold-start. No memory between runs.
Best on backends with real workflows. Static sites and pure UI repos don't surface much.
Opinionated, not exhaustive. Prefers 3 ironclad findings to 30 maybe-bugs.
The bug-finding recipe Farfield uses in production, packaged as a Claude Code skill. Free + MIT.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
