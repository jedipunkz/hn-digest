---
source: "https://github.com/mgranados/screenshotter"
hn_url: "https://news.ycombinator.com/item?id=48346481"
title: "Show HN: Compress your screenshots for AI coding agents"
article_title: "GitHub - mgranados/screenshotter: Small utility to compress screenshots in macos and copy to clipboard · GitHub"
author: "mgranados"
captured_at: "2026-06-01T00:28:45Z"
capture_tool: "hn-digest"
hn_id: 48346481
score: 2
comments: 0
posted_at: "2026-05-31T15:27:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Compress your screenshots for AI coding agents

- HN: [48346481](https://news.ycombinator.com/item?id=48346481)
- Source: [github.com](https://github.com/mgranados/screenshotter)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T15:27:38Z

## Translation

タイトル: HN の表示: AI コーディング エージェント用にスクリーンショットを圧縮する
記事のタイトル: GitHub - mgranados/screenshotter: MacOS でスクリーンショットを圧縮してクリップボードにコピーする小さなユーティリティ · GitHub
説明: MacOS でスクリーンショットを圧縮し、クリップボードにコピーする小さなユーティリティ - mgranados/screenshotter
HN テキスト: AI エージェント インターフェイスに貼り付ける前に、スクリーンショットを圧縮して帯域幅を節約するためのローカル ユーティリティ (macOS)。これにより帯域幅とコストが節約されることを願っています。

記事本文:
GitHub - mgranados/screenshotter: MacOS でスクリーンショットを圧縮し、クリップボードにコピーするための小さなユーティリティ · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ムグラナドス
/
スクリーンショット撮影者
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github .github bin bin docs docs 例/ラッパーの例/ラッパーの拡張機能/スクリーンショットの拡張機能/スクリーンショットのスクリプト スクリプトのスキル/スクリーンショットのスキル/スクリーンショット .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェントのローカル macOS スクリーンショット。
スクリーンショットを撮ります。スクリーンショット ツールはそれをローカルで最適化し、クリップボードにコピーします。
git clone https://github.com/mgranados/screenshotter.git
CD スクリーンショット
ノード bin/screenshoter.mjs ドクター
オプション:
mkdir -p ~ /.local/bin
ln -sf " $PWD /bin/screenshotter.mjs " ~ /.local/bin/screenshotter
使用する
スクリーンショット ウォッチ --verbose
Cmd+Shift+3 または Cmd+Shift+4 でスクリーンショットを撮り、 Cmd+V で Codex、Claude、または別のエージェントに貼り付けます。
pi をインストールします。 -l
次に、 /screenshoter を実行します。
サイズ
オリジナル
デフォルト
保存されたサイズ
帯域幅の節約 / 1k
プロ ディスプレイ XDR 6016x3384
5.48MB
0.89MB
93%
5.0GB
16インチMacBook Pro 3456x2234
1.86MB
0.83MB
89%
1.6GB
14インチMacBook Pro 3024x1964
2.34MB
0.75MB
91%
2.1GB
ウィンドウ 1920x1200
1.04MB
0.40MB
81%
0.8GB
ウィンドウ 1440x900
0.63MB
0.38MB
68%
0.4GB
最近の 5 つのスクリーンショットの平均。デフォルトでは可読性が維持されます。 API イメージ トークンのコストを下げるために、トークン モードのサイズが変更されます。ダウンスケールのデフォルトは、Apple Vision のテキスト可読性ベンチマークでチェックされます。
アップロード帯域幅: 多くの場合、2 ～ 5 MB -> <1 MB 。
貼り付け/送信の遅延: Codex または Claude が取り込む画像データが少なくなります。
ローカル ストレージ: 最適化されたコピーのサイズは小さくなります。
信頼性: アタッチメントの制限に達する可能性が低くなります。

バイトあたりの可読性: 高い次元を維持しながら効率的にエンコードします。
スクリーンショット ウォッチ -- プロファイルの可読性 # デフォルト
スクリーンショット ウォッチ -- バランスのとれたプロファイル
スクリーンショット ウォッチ -- プロファイル トークン
pi の場合: /screenshotter readability 、 /screenshotter Balanced 、または /screenshotter token 。
スクリーンショット ウォッチ --verbose
スクリーンショット クリップ --target codex-app
スクリーンショット クロード-アプリ --verbose
スクリーンショットの準備-最新 --ターゲットマニュアル --json
スクリーンショット クレーム --ターゲット マニュアル --json
スクリーンショット ベンチ --最新 20 --トークン --json
スクリーンショットの医者
MCP、実験的:
codex mcp スクリーンショットを追加 -- スクリーンショット mcp-server
クロード mcp スクリーンショットを追加 -- スクリーンショット mcp-server
シンボリックリンクがありません:
ノード bin/screenshotter.mjs watch --verbose
Verbose を実行すると、JSONL ログが次の場所に書き込まれます。
~/ライブラリ/アプリケーション サポート/スクリーンショット/ログ/イベント.jsonl
ライセンス
MacOS でスクリーンショットを圧縮し、クリップボードにコピーするための小さなユーティリティ
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

Small utility to compress screenshots in macos and copy to clipboard - mgranados/screenshotter

Local util (macOS) to compress the screenshots and save some bandwidth before pasting on ai agent interfaces, hopefully this saves some bandwidth and costs.

GitHub - mgranados/screenshotter: Small utility to compress screenshots in macos and copy to clipboard · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
mgranados
/
screenshotter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github .github bin bin docs docs examples/ wrappers examples/ wrappers extensions/ screenshotter extensions/ screenshotter scripts scripts skills/ screenshotter skills/ screenshotter .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Local macOS screenshots for coding agents.
Take a screenshot. screenshotter optimizes it locally and copies it to your clipboard.
git clone https://github.com/mgranados/screenshotter.git
cd screenshotter
node bin/screenshotter.mjs doctor
Optional:
mkdir -p ~ /.local/bin
ln -sf " $PWD /bin/screenshotter.mjs " ~ /.local/bin/screenshotter
Use
screenshotter watch --verbose
Take a screenshot with Cmd+Shift+3 or Cmd+Shift+4 , then paste into Codex, Claude, or another agent with Cmd+V .
pi install . -l
Then run /screenshotter on .
Size
Original
Default
Size saved
Bandwidth saved / 1k
Pro Display XDR 6016x3384
5.48 MB
0.89 MB
93%
5.0 GB
16in MacBook Pro 3456x2234
1.86 MB
0.83 MB
89%
1.6 GB
14in MacBook Pro 3024x1964
2.34 MB
0.75 MB
91%
2.1 GB
Window 1920x1200
1.04 MB
0.40 MB
81%
0.8 GB
Window 1440x900
0.63 MB
0.38 MB
68%
0.4 GB
Average from 5 recent screenshots. Default preserves readability. Token mode resizes for lower API image-token cost. Downscale defaults are checked with Apple Vision text-readability benchmarks.
Upload bandwidth: often 2-5 MB -> <1 MB .
Paste/send latency: less image data for Codex or Claude to ingest.
Local storage: optimized copies are smaller.
Reliability: less likely to hit attachment limits.
Readability per byte: efficient encoding while keeping dimensions high.
screenshotter watch --profile readability # default
screenshotter watch --profile balanced
screenshotter watch --profile token
In pi: /screenshotter readability , /screenshotter balanced , or /screenshotter token .
screenshotter watch --verbose
screenshotter clip --target codex-app
screenshotter claude-app --verbose
screenshotter prepare-latest --target manual --json
screenshotter claim --target manual --json
screenshotter bench --latest 20 --tokens --json
screenshotter doctor
MCP, experimental:
codex mcp add screenshotter -- screenshotter mcp-server
claude mcp add screenshotter -- screenshotter mcp-server
No symlink:
node bin/screenshotter.mjs watch --verbose
Verbose runs write JSONL logs to:
~/Library/Application Support/screenshotter/logs/events.jsonl
License
Small utility to compress screenshots in macos and copy to clipboard
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
