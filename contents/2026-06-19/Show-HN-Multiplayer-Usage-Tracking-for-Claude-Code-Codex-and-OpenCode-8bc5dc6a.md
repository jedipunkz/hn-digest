---
source: "https://github.com/useautumn/summer"
hn_url: "https://news.ycombinator.com/item?id=48601140"
title: "Show HN: Multiplayer Usage Tracking for Claude Code, Codex and OpenCode"
article_title: "GitHub - useautumn/summer: Track AI usage across your organization · GitHub"
author: "johnyeocx"
captured_at: "2026-06-19T18:42:59Z"
capture_tool: "hn-digest"
hn_id: 48601140
score: 1
comments: 0
posted_at: "2026-06-19T17:51:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Multiplayer Usage Tracking for Claude Code, Codex and OpenCode

- HN: [48601140](https://news.ycombinator.com/item?id=48601140)
- Source: [github.com](https://github.com/useautumn/summer)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T17:51:03Z

## Translation

タイトル: HN を表示: クロード コード、コーデックス、および OpenCode のマルチプレイヤー使用状況追跡
記事のタイトル: GitHub - 秋/夏の使用: 組織全体で AI の使用状況を追跡する · GitHub
説明: 組織全体で AI の使用状況を追跡します。 GitHub でアカウントを作成して、秋/夏の開発に貢献してください。
HN テキスト: HN さん、私はこのオープン ソース ツールを構築しました。これにより、あなたとあなたのチームは、Claude Code / Codex / OpenCode 全体の使用状況を追跡できます。 1. ホスティングは関係ありません。使用量は https://useautumn.com/ によって保存および管理されます。 2. あなたとチーム メンバーは、ローカル ダッシュボードを起動して、使用量の優れたグラフを表示できます。 3. 必要に応じて、過去の使用量をバックフィルすることができます。 これを構築したのは、ccusage (素晴らしいツールです) を頻繁に実行して、チームの各メンバーが蓄積した AI の使用量を確認し、比較して楽しんでいるからです。このツールを使用すると、すべてが集中化され、モデルやハーネスごとにグループ化/フィルタリングし、統一された方法で異なるメンバー間の使用状況を比較および追跡できます。 tldr は、チーム間でこれらのエージェントの使用状況を比較する楽しい方法であり、5 分未満でセットアップできます。楽しんでください！

記事本文:
GitHub - 秋/夏の使用: 組織全体で AI の使用状況を追跡 · GitHub
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
秋を使う
/
夏
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード

その他のアクションメニューを開く フォルダーとファイル
7 コミット 7 コミット ダッシュ ダッシュ docs docs src src test test .gitignore .gitignore README.md README.md SPEC.md SPEC.md bun.lock bun.lock bunfig.toml bunfig.toml package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Summer は、AI コーディングの使用と支出のためのローカルのオープンソース ツールであり、によって構築されました。
秋。開発者ごとに、Claude Code、Codex、および
opencode を実行すると、各エンジニアはどのモデルでどれくらいの量を使用しており、その価値はいくらでしょうか?
夏にはホスティングは必要ありません。バックエンドとして秋を使用し、
ローカル ダッシュボードを使用 — Autumn はすべての使用イベントを保存し、トークンの価格を設定します (経由)
Models.dev ) を作成し、チーム全体の使用状況を集計します。
注: すでに Autumn 組織がある場合は、Summer で使用する新しい組織を作成します。
bun -g install @useautumn/Summer # Summer をインストール
夏はセットアップを開始し、バックグラウンドで使用状況を追跡します
サマーダッシュ # 使用状況ダッシュボードを開く
Summer Start は、初回セットアップに必要なすべてを実行します。
OAuth 経由で Autumn で認証します — ログイン (またはサインアップ) し、組織を設定します。
履歴を埋め戻すことを提案します。
ローカル デーモンを起動して、Claude Code + Codex の使用状況を収集し、Autumn に送信します。
Summer Dash はローカル UI を提供します (ブラウザーで開きます)。使用状況グラフを表示できます。
ハーネス/モデル/ユーザー/請求モードごとにグループ化し、任意のプロパティでフィルタリングし、検索します
開発者ごとの使用状況を確認し、生のイベントを検査します。
注: Summer は自動起動サービス (launchd/systemd) をインストールするので、再起動しても存続します。
プレーン バックグラウンド プロセスの場合は --no-service を渡します。
夏期開始後に使用量を送信するには、Claude Code を再起動する必要があります。新しいセッションを開始するか、既存のセッションの場合は再起動して /resume します。 (Codex/opencode は再起動する必要はありません。)
夏には、秋の組織内の全員の使用量がロールアップされます。
開ける

秋の組織設定 →
メンバー 。
彼らは秋に招待を受け入れ、夏に自分たちで走り始め、その後クロード・コードを再開します
(上記の注を参照してください。そうでない場合は、バックフィルのみが表示され、実際のクロードの使用法は表示されません)。
それだけです。それらの使用状況は、ダッシュボードにあなたの使用状況と並んで表示されます。
コマンド
何をするのか
夏の始まり
(必要な場合) をセットアップし、追跡を開始します。
サマーダッシュ
使用状況ダッシュボード (別名: ダッシュボード ) を開きます。
夏の埋め戻し
過去のクロード コード + コーデックスの使用状況 (過去のもの) をインポートします。
夏のレポート
ターミナルの使用状況ロールアップ。
夏のステータス
認証 + ローカル状態。
サマーストップ
ハーネス設定を停止して復元します (自動起動も削除します)。
サマーサービスのインストール/アンインストール/ステータス
ブート時の自動起動を管理します。
夏のログイン/ログアウト
秋の認証を管理します。
サポートされているハーネス
組織全体で AI の使用状況を追跡する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Track AI usage across your organization. Contribute to useautumn/summer development by creating an account on GitHub.

Hey HN, I built this open source tool which lets you and your team track usage across Claude Code / Codex / OpenCode. 1. There's no hosting involved. Usage is stored and managed by https://useautumn.com/ 2. You and your team members can spin up a local dashboard to view nice charts of your usage 3. You can backfill your historical usage if you wish to I built this because we often run ccusage (which is an amazing tool) to check how much AI spend each of us in our team have accumulated, and have fun comparing it. With this tool, everything is centralized, you can group by / filter models / harnesses, and compare & track usage between different members in a unified way. tldr, it's a fun way to compare usage of these agents between your team, and can be set up in <5 minutes. Have fun with it!

GitHub - useautumn/summer: Track AI usage across your organization · GitHub
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
useautumn
/
summer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits dash dash docs docs src src test test .gitignore .gitignore README.md README.md SPEC.md SPEC.md bun.lock bun.lock bunfig.toml bunfig.toml package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Summer is a local, open-source tool for AI-coding usage and spend , built by
Autumn . Per developer, across Claude Code , Codex , and
opencode , it answers: how much is each engineer using, on what models, and what's it worth?
Summer needs zero hosting . It uses Autumn as its backend and comes
with a local dashboard — Autumn stores every usage event, prices tokens (via
Models.dev ), and aggregates usage across your whole team.
Note: If you already have an Autumn org, create a new org to use with Summer.
bun -g install @useautumn/summer # install Summer
summer start # set up, then track usage in the background
summer dash # open the usage dashboard
summer start does everything first-time setup needs:
Authenticates with Autumn via OAuth — logs you in (or signs you up) and sets up your org.
Offers to backfill your history.
Starts a local daemon to collect Claude Code + Codex usage and send it to Autumn.
summer dash serves a local UI (and opens it in your browser): a usage chart you can
group by harness / model / user / billing mode, filter by any property, search
per-developer usage, and inspect the raw events .
Note: Summer installs an autostart service (launchd/systemd) so it survives reboots.
Pass --no-service for a plain background process.
Claude Code requires a restart to send usage after summer start . Either start new sessions, or for existing ones, restart and /resume . (Codex/opencode don't need a restart.)
Summer rolls up usage across everyone in your Autumn org.
Open your Autumn org settings →
Members .
They accept the invite in Autumn, run summer start themselves, then restart Claude Code
(see the note above — otherwise only their backfill shows up, not live Claude usage).
That's it — their usage shows up alongside yours in the dashboard.
Command
What it does
summer start
Set up (if needed) and start tracking.
summer dash
Open the usage dashboard (alias: dashboard ).
summer backfill
Import historical Claude Code + Codex usage (backdated).
summer report
Usage rollup in the terminal.
summer status
Auth + local state.
summer stop
Stop and restore harness settings (also removes autostart).
summer service install / uninstall / status
Manage on-boot autostart.
summer login / logout
Manage Autumn auth.
Harnesses Supported
Track AI usage across your organization
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
