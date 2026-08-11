---
source: "https://github.com/stared/trmnl-x-claude-codex"
hn_url: "https://news.ycombinator.com/item?id=49255546"
title: "Show HN: Claude Code and Codex usage screen for TRMNL X e-ink"
article_title: "GitHub - stared/trmnl-x-claude-codex: TRMNL X screen for Claude Code and Codex vibe coders - use as it is or fork to your taste · GitHub"
author: "stared"
captured_at: "2026-08-11T09:51:33Z"
capture_tool: "hn-digest"
hn_id: 49255546
score: 1
comments: 0
posted_at: "2026-08-11T09:44:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude Code and Codex usage screen for TRMNL X e-ink

- HN: [49255546](https://news.ycombinator.com/item?id=49255546)
- Source: [github.com](https://github.com/stared/trmnl-x-claude-codex)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T09:44:03Z

## Translation

タイトル: TRMNL X e-ink の HN: クロード コードとコーデックスの使用画面を表示します。
記事のタイトル: GitHub - stared/trmnl-x-claude-codex: Claude Code および Codex バイブコーダー用の TRMNL X 画面 - そのまま使用するか、好みに合わせてフォークしてください · GitHub
説明: Claude Code および Codex バイブコーダー用の TRMNL X 画面 - そのまま使用するか、好みに応じてフォークして使用します - stared/trmnl-x-claude-codex
HN テキスト: 「思い出にお金がかかる」 - クロード コードに月額わずか 200 ドル支払っていますが、1 週間に 2000 ドル相当のトークンを費やすのは簡単です。 TRMLN X を取得したばかりなので (大好きですが、所属はしていません)、制限だけでなく、1 日ごと、プロジェクトのレコードごとに表示することにしました。

記事本文:
GitHub - stared/trmnl-x-claude-codex: Claude Code および Codex バイブコーダー用の TRMNL X 画面 - そのまま使用するか、好みに合わせてフォークしてください · GitHub
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
見つめた
/
trmnl-x-クロード-コーデックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .gitignore .gitignore ライセンス ライセンス README.md README.md com.pmigdal.trmnl-usage.plist com.pmigdal.trmnl-usage.plist config.json.exampl

e config.json.example Push_usage.py Push_usage.py template.liquid template.liquid trmnl-x-claude-codex-photo.jpg trmnl-x-claude-codex-photo.jpg すべてのファイルを表示 リポジトリ ファイルのナビゲーション
TRMNL X — クロードコード + コーデックス使用画面
TRMNL プライベート プラグインの e-ink ダッシュボード: レート
クロード コードとコーデックスの制限、エージェントごとの 1 日あたりの API 相当ドル、および
上位のプロジェクト — 10 分ごとに Mac からプッシュされます。
データ ソース (サーバーなし、すべてがローカル状態を読み取ります):
クロードの制限 + 計画 — 文書化されていない api.anthropic.com/api/oauth/usage
エンドポイント、macOS キーチェーンからの OAuth トークン。非公式;レート制限があるため、
10 分以内にポーリングしないでください。
コーデックスの制限 — コーデックス アプリサーバー JSON-RPC。 CLI が認証を処理します。
$ / 日とトッププロジェクト — ccusage via
pnpm dlx 、およびプロジェクト名用の ~/.claude/projects/ JSONL
日ごとのスパークライン。
trmnl.com で、Webhook 戦略を使用してプライベート プラグインを作成します。コピー
そのURL。
cp config.json.example config.json 、URLを貼り付けます（ config.json は
gitignored — UUID を使用すると、誰でも画面にプッシュできます)。
template.liquid をプラグインのマークアップ エディターに貼り付けます。
テスト: uv run --no-project Push_usage.py --dry-run を実行し、フラグを付けません。
スケジュールをインストールします (マシンに一致するように plist 内のパスを編集します)。
cp com.pmigdal.trmnl-usage.plist ~ /Library/LaunchAgents/
launchctl ブートストラップ gui/ $( id -u ) ~ /Library/LaunchAgents/com.pmigdal.trmnl-usage.plist
ログ: /tmp/trmnl-usage.log 。フェッチの失敗も警告ストリップとして表示されます
デバイス上でゼロ以外の値で終了します。
キー
デフォルト
webhook_url
—
必須 (または TRMNL_WEBHOOK_URL 環境変数)
コスト_日
7
コストチャートとスパークラインの日数
トッププロジェクト
6
プロジェクト行
chart_max_px
230
最も高いグラフのバー、ピクセル
注意事項
このテンプレートは、TRMNL X の論理ビューポート (~936×702 @ 2 倍の密度) をターゲットとしています。
使用率 90% 以上のタイルが反転します

黒に。
TRMNL Webhook の上限: 12 プッシュ/時間、2 KB ペイロード。
Mac がスリープしている間、画面が古くなります。「更新済み」スタンプを確認してください。
Claude Code および Codex バイブコーダー用の TRMNL X 画面 - そのまま使用するか、好みに応じてフォークして使用します
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

TRMNL X screen for Claude Code and Codex vibe coders - use as it is or fork to your taste - stared/trmnl-x-claude-codex

"Memento costly" - while I pay just $200/month for Claude Code, it is easy to spend $2000 worth in tokens a week. Since I just got TRMLN X (I adore it, and I am not affiliated), I decided to show it - not only limits, but per day and per project record.

GitHub - stared/trmnl-x-claude-codex: TRMNL X screen for Claude Code and Codex vibe coders - use as it is or fork to your taste · GitHub
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
stared
/
trmnl-x-claude-codex
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .gitignore .gitignore LICENSE LICENSE README.md README.md com.pmigdal.trmnl-usage.plist com.pmigdal.trmnl-usage.plist config.json.example config.json.example push_usage.py push_usage.py template.liquid template.liquid trmnl-x-claude-codex-photo.jpg trmnl-x-claude-codex-photo.jpg View all files Repository files navigation
TRMNL X — Claude Code + Codex usage screen
An e-ink dashboard for a TRMNL private plugin: rate
limits for Claude Code and Codex, API-equivalent $ per day per agent, and
your top projects — pushed from your Mac every 10 minutes.
Data sources (no servers, everything reads local state):
Claude limits + plan — undocumented api.anthropic.com/api/oauth/usage
endpoint, OAuth token from the macOS Keychain. Unofficial; rate-limited, so
don't poll faster than ~10 min.
Codex limits — codex app-server JSON-RPC; the CLI handles auth.
$ / day and top projects — ccusage via
pnpm dlx , plus ~/.claude/projects/ JSONLs for project names and
per-day sparklines.
On trmnl.com create a Private Plugin with strategy Webhook ; copy
its URL.
cp config.json.example config.json , paste the URL ( config.json is
gitignored — the UUID lets anyone push to your screen).
Paste template.liquid into the plugin's Markup editor.
Test: uv run --no-project push_usage.py --dry-run , then without the flag.
Install the schedule (edit paths in the plist to match your machine):
cp com.pmigdal.trmnl-usage.plist ~ /Library/LaunchAgents/
launchctl bootstrap gui/ $( id -u ) ~ /Library/LaunchAgents/com.pmigdal.trmnl-usage.plist
Logs: /tmp/trmnl-usage.log . Fetch failures also render as a warning strip
on the device and exit non-zero.
key
default
webhook_url
—
required (or TRMNL_WEBHOOK_URL env var)
cost_days
7
days in the cost chart and sparklines
top_projects
6
project rows
chart_max_px
230
tallest chart bar, px
Notes
The template targets TRMNL X's logical viewport (~936×702 @ 2× density);
tiles at ≥90% usage invert to black.
TRMNL webhook caps: 12 pushes/hour, 2 KB payload.
The screen goes stale while the Mac sleeps — see the "updated" stamp.
TRMNL X screen for Claude Code and Codex vibe coders - use as it is or fork to your taste
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
