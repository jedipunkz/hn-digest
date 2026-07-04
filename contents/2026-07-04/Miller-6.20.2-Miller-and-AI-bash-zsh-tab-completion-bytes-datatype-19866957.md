---
source: "https://github.com/johnkerl/miller/releases/tag/v6.20.2"
hn_url: "https://news.ycombinator.com/item?id=48788362"
title: "Miller 6.20.2: Miller and AI, bash/zsh tab-completion, bytes datatype"
article_title: "Release Miller 6.20.2: Miller and AI, bash/zsh tab-completion, bytes datatype · johnkerl/miller · GitHub"
author: "john_kerl"
captured_at: "2026-07-04T20:44:15Z"
capture_tool: "hn-digest"
hn_id: 48788362
score: 2
comments: 0
posted_at: "2026-07-04T19:50:17Z"
tags:
  - hacker-news
  - translated
---

# Miller 6.20.2: Miller and AI, bash/zsh tab-completion, bytes datatype

- HN: [48788362](https://news.ycombinator.com/item?id=48788362)
- Source: [github.com](https://github.com/johnkerl/miller/releases/tag/v6.20.2)
- Score: 2
- Comments: 0
- Posted: 2026-07-04T19:50:17Z

## Translation

タイトル: Miller 6.20.2: Miller と AI、bash/zsh タブ補完、バイト データ型
記事のタイトル: Miller 6.20.2 のリリース: Miller と AI、bash/zsh タブ補完、バイト データ型 · johnkerl/miller · GitHub
説明: Miller は、CSV、TSV、表形式 JSON などの名前インデックス付きデータの awk、sed、cut、join、sort に似ています - Miller 6.20.2 のリリース: Miller と AI、bash/zsh タブ補完、バイト データ型 · johnkerl/miller

記事本文:
Miller 6.20.2 のリリース: Miller と AI、bash/zsh タブ補完、バイト データ型 · johnkerl/miller · GitHub
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
ジョンケル
/
製粉業者
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ミラー

6.20.2: Miller と AI、bash/zsh タブ補完、バイト データ型
フィルター
読み込み中
申し訳ありませんが、問題が発生しました。
読み込み中にエラーが発生しました。このページをリロードしてください。
v6.20.2
05190f3
6.20.2 の新機能
ミラー氏は AI エージェントを一流のユーザーとして扱うようになりました
bash と zsh にはシェルのタブ補完があります
新しいバイト データ型と、base64 エンコード/デコード サポートが追加されました。
(注: バージョン 6.20.0 と 6.20.1 にはエラーがあったため、リリースされていません。)
ミラー氏は現在、AI エージェントを一流のユーザーとして扱っています。エージェントがミラーをうまく動かすために必要なものすべて --
存在するものを発見し、データの形状を学習し、式を実行する前に検証します。
エラーからの回復 -- mlr バイナリの組み込み構造化機能となり、パッケージ化されました。
1 行のセットアップによる MCP を話すエージェント (Claude Code、Claude Desktop、Cursor など)。あなたの選択
#クロード
mlr スキルのインストール ~/.claude/skills/miller
# コーデックスとジェミニ
mlr スキルのインストール ~/.agents/skills/miller
または
#クロード
クロード mcp ミラーを追加 -- mlr mcp
# コーデックス
codex mcp add miller -- mlr mcp
#ジェミニ
ジェミニ mcp 追加ミラー mlr mcp
クイック スタートについては、新しい Miller と AI のドキュメント ページを参照してください。
https://miller.readthedocs.io/en/latest/ai
https://miller.readthedocs.io/en/latest/agent-skill
https://miller.readthedocs.io/en/latest/mcp-server
https://miller.readthedocs.io/en/latest/ai-support
機械可読ヘルプ カタログ、mlr help --as-json : #2099 、 #2106 の @johnkerl による
#2107 の @johnkerl による JSON 機能インデックスとどの意図ルーターの mlr
#2111 の @johnkerl による列挙値セットを使用した構造化された動詞オプション
#2113 の @johnkerl によるヒントと意図的な提案を含む構造化エラー出力 --errors-json
DSL 検証/ドライラン、mlr put --explain / mlr filter --explain by @johnkerl (#2131)
新しい mlr description 動詞: フィールド名、タイプ、カーディナリティ、および

#2132 の @johnkerl による 1 回のパスでのデータの値ドメイン
#2133 の @johnkerl による --no-shell サンドボックス フラグを備えた新しい mlr mcp MCP サーバーおよびエージェント プレイブック
#2103 の @johnkerl によるロードマップ ドキュメント。 #2134 の @johnkerl による「Miller と AI」ドキュメント ページ
bash および zsh のシェルタブ補完
https://miller.readthedocs.io/en/latest/shell-completion/
#2096 の @johnkerl による bash と zsh のシェル タブ補完
DSL のファーストクラスのバイト型
https://miller.readthedocs.io/en/latest/reference-main-data-types/
https://miller.readthedocs.io/en/latest/reference-dsl-builtin-functions/#base64_encode
https://miller.readthedocs.io/en/latest/reference-dsl-builtin-functions/#base64_decode
#2122 の @johnkerl による b"..." リテラルと Base64/hex コーデックを使用した新しいバイト型
#2071 の @farnoy による新しい tail -n +N および head -n -N オプション
#2100 の @johnkerl による長年の期限切れの --md フラグ
@teo-tsirpanis による Windows-arm64 リリース (#2127)
#2129 の @johnkerl によるマスクされた unset-on-array エラー パスを修正 (govet lint の発見とともに)
#2123 の @dashitongzhi による CSV および JSON のトラブルシューティングのヒント、#2128 の @johnkerl によるページリファクタリング
#2115 の @dashitongzhi による時間変換スレッドの安全性を文書化します。
#2095 の @johnkerl による issue-2084 の perf mods の perf ドキュメント内の PNG グラフィックス
#2076 の @dashitongzhi による golangci-lint CI ワークフロー
#2108 、 #2110 、 #2112 、 #2130 で @johnkerl により staticcheck と errcheck をゼロにする lint 修正
#2121 の @johnkerl により pkg/ からデッドコードを削除します
#2119 の @johnkerl と #2101 、 #2102 、 #2104 、 #2105 、 #2114 、 #2116 、 #2118 、 #2124 、 #2125 、 #2126 の @dependabot による CI/アクションの更新
#2117 、 #2120 の @dependabot による依存関係のバンプ
@farnoy は #2071 で最初の投稿を行いました
@dashitongzhi は #2076 で最初の投稿を行いました
@teo-tsirpanis は #2127 で最初の投稿を行いました
完全な変更ログ:

v6.19.0...v6.20.0
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Miller is like awk, sed, cut, join, and sort for name-indexed data such as CSV, TSV, and tabular JSON - Release Miller 6.20.2: Miller and AI, bash/zsh tab-completion, bytes datatype · johnkerl/miller

Release Miller 6.20.2: Miller and AI, bash/zsh tab-completion, bytes datatype · johnkerl/miller · GitHub
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
johnkerl
/
miller
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Miller 6.20.2: Miller and AI, bash/zsh tab-completion, bytes datatype
Filter
Loading
Sorry, something went wrong.
There was an error while loading. Please reload this page .
v6.20.2
05190f3
What's new in 6.20.2
Miller now treats AI agents as first-class users
There is shell tab-completion for bash and zsh
There is a new bytes datatype, along with base64 encode/decode support
(Note: versions 6.20.0 and 6.20.1 were in error and are not released.)
Miller now treats AI agents as first-class users. Everything an agent needs to drive Miller well --
discovering what exists, learning your data's shape, validating expressions before running them, and
recovering from errors -- is now a built-in, structured feature of the mlr binary, packaged for
MCP-speaking agents (Claude Code, Claude Desktop, Cursor, ...) via a one-line setup. Your choice of
# Claude
mlr skill install ~/.claude/skills/miller
# Codex and Gemini
mlr skill install ~/.agents/skills/miller
or
# Claude
claude mcp add miller -- mlr mcp
# Codex
codex mcp add miller -- mlr mcp
# Gemini
gemini mcp add miller mlr mcp
See the new Miller and AI docs page for the quick start.
https://miller.readthedocs.io/en/latest/ai
https://miller.readthedocs.io/en/latest/agent-skill
https://miller.readthedocs.io/en/latest/mcp-server
https://miller.readthedocs.io/en/latest/ai-support
Machine-readable help catalog, mlr help --as-json : by @johnkerl in #2099 , #2106
JSON capability index and mlr which intent router by @johnkerl in #2107
Structured verb options, with enum value-sets by @johnkerl in #2111
Structured error output, --errors-json , with hints and did-you-mean suggestions by @johnkerl in #2113
DSL validate/dry-run, mlr put --explain / mlr filter --explain by @johnkerl in #2131
New mlr describe verb: field names, types, cardinality, and value domains of your data in one pass by @johnkerl in #2132
New mlr mcp MCP server and agent playbook, with a --no-shell sandbox flag by @johnkerl in #2133
Roadmap doc by @johnkerl in #2103 ; "Miller and AI" docs page by @johnkerl in #2134
Shell tab-completion for bash and zsh
https://miller.readthedocs.io/en/latest/shell-completion/
Shell tab-completion for bash and zsh by @johnkerl in #2096
A first-class bytes type in the DSL
https://miller.readthedocs.io/en/latest/reference-main-data-types/
https://miller.readthedocs.io/en/latest/reference-dsl-builtin-functions/#base64_encode
https://miller.readthedocs.io/en/latest/reference-dsl-builtin-functions/#base64_decode
New bytes type with b"..." literals and base64/hex codecs by @johnkerl in #2122
New tail -n +N and head -n -N options by @farnoy in #2071
Long-overdue --md flag by @johnkerl in #2100
Windows-arm64 releases by @teo-tsirpanis in #2127
Fix masked unset-on-array error path (along with govet lint findings) by @johnkerl in #2129
CSV and JSON troubleshooting tips by @dashitongzhi in #2123 , with page-refactor by @johnkerl in #2128
Document time-conversion thread safety by @dashitongzhi in #2115
PNG graphics in perf docs for issue-2084 perf mods by @johnkerl in #2095
golangci-lint CI workflow by @dashitongzhi in #2076
Lint fixes bringing staticcheck and errcheck to zero by @johnkerl in #2108 , #2110 , #2112 , #2130
Strip dead code from pkg/ by @johnkerl in #2121
CI/actions updates by @johnkerl in #2119 and by @dependabot in #2101 , #2102 , #2104 , #2105 , #2114 , #2116 , #2118 , #2124 , #2125 , #2126
Dependency bumps by @dependabot in #2117 , #2120
@farnoy made their first contribution in #2071
@dashitongzhi made their first contribution in #2076
@teo-tsirpanis made their first contribution in #2127
Full Changelog : v6.19.0...v6.20.0
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
