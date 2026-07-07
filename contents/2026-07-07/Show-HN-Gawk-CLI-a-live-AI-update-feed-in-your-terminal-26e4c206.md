---
source: "https://github.com/Neelagiri65/gawk-cli"
hn_url: "https://news.ycombinator.com/item?id=48823561"
title: "Show HN: Gawk CLI – a live AI update feed in your terminal"
article_title: "GitHub - Neelagiri65/gawk-cli: The gawk.dev live AI-ecosystem feed in your terminal — every number traces to its source · GitHub"
author: "gawkdev"
captured_at: "2026-07-07T21:15:48Z"
capture_tool: "hn-digest"
hn_id: 48823561
score: 1
comments: 0
posted_at: "2026-07-07T20:51:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Gawk CLI – a live AI update feed in your terminal

- HN: [48823561](https://news.ycombinator.com/item?id=48823561)
- Source: [github.com](https://github.com/Neelagiri65/gawk-cli)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T20:51:06Z

## Translation

タイトル: HN を表示: Gawk CLI – 端末内のライブ AI アップデート フィード
記事のタイトル: GitHub - Neelagiri65/gawk-cli: 端末内の gawk.dev ライブ AI エコシステム フィード - すべての数値がそのソースを追跡 · GitHub
説明: ターミナル内の gawk.dev ライブ AI エコシステム フィード - すべての数値がそのソースを追跡します - Neelagiri65/gawk-cli
HN テキスト: 「最速のモデル」「最も使用されている SDK」 誰が言ったのですか? LinkedIn で手を振ったり、AI がスロップしたりすることなく、AI 更新フィードが必要だったので、それを端末に組み込みました。 npx gawk-cli 何がエキサイティングかというと、出典が引用されているということです。 CLI でいくつかのフィードを作成しましたが、どのフィードをそこに置く価値があるのか​​知りたいと思っています。

記事本文:
GitHub - Neelagiri65/gawk-cli: ターミナル内の gawk.dev ライブ AI エコシステム フィード - すべての数値がそのソースを追跡 · GitHub
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
ニーラギリ65
/
gawk-cli
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット bin bin lib lib test test ライセンス ライセンス README.md README.md package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
ターミナル内の gawk.dev ライブ AI エコシステム フィード。
npx gawk-cli # 現時点での厳選されたワイヤー
npx gawk-cli モデル # OpenRouter トップ週間ランキング
npx gawk-cli tools # ツールの健全性、ベンダー宣言
npx gawk-cli sdk # レジストリ カウンタからの SDK の導入
フラグ: --limit N · --type TOOL_ALERT · --json (生の API 応答、逐語的) · --no-color 。
ダッシュボードと同じです。すべての数値は公開ソースを追跡します。各カードには、ソース名、年齢、URL が印刷されます。 CLI は、gawk のパブリック読み取り API 上のシン クライアントです。CLI は、表示される数値の算術以外の何も導き出しません。ソースが沈黙しているかダウンしている場合は、正直に沈黙し、決して発明することはありません。
依存関係ゼロ。ノード ≥ 18。認証なし、追跡なし (使用量はサーバー側で匿名集計としてのみカウントされます。ハッシュ化され、ID はカウントされません)。
グローバルにインストールされる ( npm i -g gawk-cli )、コマンドは gawk-cli です。システム上の GNU awk に属する gawk ではなく、意図的にです。
ターミナル内の gawk.dev ライブ AI エコシステム フィード - すべての数値がそのソースを追跡します
読み込み中にエラーが発生しました。このページをリロードしてください。
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

The gawk.dev live AI-ecosystem feed in your terminal — every number traces to its source - Neelagiri65/gawk-cli

"Fastest model" "Most used SDK" Says who? I wanted the AI update feed without the hand waving and without the ai slops in linkedin, so I built it into my terminal. npx gawk-cli what's exciting is.. it is source cited. I have built some feeds in CLI but curious to know which feeds are worth having there.

GitHub - Neelagiri65/gawk-cli: The gawk.dev live AI-ecosystem feed in your terminal — every number traces to its source · GitHub
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
Neelagiri65
/
gawk-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits bin bin lib lib test test LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
The gawk.dev live AI-ecosystem feed in your terminal.
npx gawk-cli # the curated wire, right now
npx gawk-cli models # OpenRouter top-weekly rankings
npx gawk-cli tools # tool health, vendor-declared
npx gawk-cli sdk # SDK adoption from registry counters
Flags: --limit N · --type TOOL_ALERT · --json (raw API response, verbatim) · --no-color .
Same as the dashboard: every number traces to a public source. Each card prints its source name, age, and URL. The CLI is a thin client over gawk's public read API — it derives nothing beyond arithmetic on numbers it also shows, and if a source is quiet or down you get honest quiet, never invention.
Zero dependencies. Node ≥ 18. No auth, no tracking (usage is counted server-side as anonymous aggregates only — hashed, never identities).
Installed globally ( npm i -g gawk-cli ), the command is gawk-cli — deliberately not gawk , which belongs to GNU awk on your system.
The gawk.dev live AI-ecosystem feed in your terminal — every number traces to its source
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
