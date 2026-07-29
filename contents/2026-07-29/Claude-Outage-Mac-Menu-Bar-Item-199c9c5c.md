---
source: "https://github.com/titojankowski/claude-status"
hn_url: "https://news.ycombinator.com/item?id=49103101"
title: "Claude Outage – Mac Menu Bar Item"
article_title: "GitHub - titojankowski/claude-status: Tiny macOS menu bar app showing status.claude.com as a green/yellow/red dot · GitHub"
author: "tito"
captured_at: "2026-07-29T21:49:48Z"
capture_tool: "hn-digest"
hn_id: 49103101
score: 3
comments: 1
posted_at: "2026-07-29T21:07:58Z"
tags:
  - hacker-news
  - translated
---

# Claude Outage – Mac Menu Bar Item

- HN: [49103101](https://news.ycombinator.com/item?id=49103101)
- Source: [github.com](https://github.com/titojankowski/claude-status)
- Score: 3
- Comments: 1
- Posted: 2026-07-29T21:07:58Z

## Translation

タイトル: クロードの停止 - Mac メニュー バー項目
記事のタイトル: GitHub - titojankowski/claude-status: status.claude.com を緑/黄/赤の点として表示する小さな macOS メニュー バー アプリ · GitHub
説明: status.claude.com を緑/黄/赤の点として表示する小さな macOS メニュー バー アプリ - titojankowski/claude-status

記事本文:
GitHub - titojankowski/claude-status: status.claude.com を緑/黄/赤の点として表示する小さな macOS メニュー バー アプリ · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ティトヤンコウスキー
/
クロードステータス
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .gitignore .gitignore ライセンス ライセンス README.md README.md build.sh build.sh main.swift main.swift すべてのファイルを表示 リポジトリ ファイルのナビゲーション
status.claude.com の現在のステータスを色付きのドットで表示する小さな macOS メニュー バー アプリ。
◯ 中空リング — すべてのシステムが動作中。テンプレート画像として描画されるので、AppKit
メニュー バーに合わせて色合いを変えます。ライト モードでは黒、ダーク モードでは白になります。
🟡 黄色で塗りつぶされた — 軽微な問題またはメンテナンス
🔴 赤で塗りつぶされる — 大規模または重大な障害
⚫️ 灰色で塗りつぶされた — ステータスチェック自体が失敗しました (オフラインなど)
トラブルのみが塗りつぶされたドットなので、健全なステータスは視覚的に静かなままです。
ドットをクリックすると内訳が表示されます。全体的な説明、各コンポーネントと独自のステータス、
アクティブなインシデント (クリックするとステータス ページが開きます)、および最後のチェックの時刻。
60 秒ごとにポーリングし、Mac がスリープから復帰するとすぐに再チェックします。
./build.sh # ClaudeStatus.app を生成します
ClaudeStatus.app を開く
単一ファイルの Swift、依存関係なし。 macOS 13 以降と Xcode コマンド ライン ツールが必要です。
中空リングでも多すぎる場合は、設定→すべてのシステムが動作しているときに非表示にします
ステータスが緑色の間はメニュー バーから完全に削除されるため、アプリのみが表示されます
実際に何かが間違っているとき。
灰色の「チェックに失敗しました」ステータスが引き続き表示されます。沈黙を健康と誤解してはいけません。
非表示のアプリにはクリックするメニューがないため、元に戻す方法は 2 つあります。
アプリを再度起動します (Finder でダブルクリックします)。ドットは 15 で再表示されます
数秒間メニューを開くことができ、その後再び非表示になります。
コマンドラインから:
ClaudeStatus.app/Contents/MacOS/ClaudeStatus --hide-when-green # 設定を表示
クロードステータス.app/C

ontents/MacOS/ClaudeStatus --hide-when-green off # 元に戻す
実行中のアプリは、60 秒以内の次のポーリングでこれを取得します。
アプリのメニューの「設定」→「ログイン」から開始します。 macOS がリクエストを
承認されると、アプリはシステム設定のログイン項目ペインを開きます。
コマンドラインから切り替えることもできます。
ClaudeStatus.app/Contents/MacOS/ClaudeStatus --login-status
ClaudeStatus.app/Contents/MacOS/ClaudeStatus --ログイン
ClaudeStatus.app/Contents/MacOS/ClaudeStatus --ログインオフ
仕組み
標準の Statuspage 概要エンドポイントを読み取ります。
https://status.claude.com/api/v2/summary.json
メニュー バーのドット マップの status.indicator (なし / マイナー / メジャー / クリティカル /
メンテナンス）色に合わせて。メニューの各行は、そのコンポーネント自体のステータスをマップします。
( 運用可能 / パフォーマンス低下 / 部分的停止 / 主要な停止 /
under_maintenance )同様に。
パブリック ドメイン — Unlicense 。著作権はありません。クローンを作成し、フォークし、
変更、出荷、販売、帰属表示は必要ありません。 「ライセンス」を参照してください。
status.claude.com を緑/黄/赤の点として表示する小さな macOS メニュー バー アプリ
Readme ライセンス解除アクティビティのスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Tiny macOS menu bar app showing status.claude.com as a green/yellow/red dot - titojankowski/claude-status

GitHub - titojankowski/claude-status: Tiny macOS menu bar app showing status.claude.com as a green/yellow/red dot · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
titojankowski
/
claude-status
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .gitignore .gitignore LICENSE LICENSE README.md README.md build.sh build.sh main.swift main.swift View all files Repository files navigation
A tiny macOS menu bar app that shows the current status of status.claude.com as a colored dot.
◯ hollow ring — all systems operational. Drawn as a template image, so AppKit
tints it to match the menu bar: black in light mode, white in dark mode.
🟡 filled yellow — minor issue or maintenance
🔴 filled red — major or critical outage
⚫️ filled gray — the status check itself failed (offline, etc.)
Only trouble is a filled dot, so a healthy status stays visually quiet.
Click the dot for a breakdown: overall description, every component with its own status,
any active incidents (click to open the status page), and the time of the last check.
Polls every 60 seconds, plus an immediate re-check when the Mac wakes from sleep.
./build.sh # produces ClaudeStatus.app
open ClaudeStatus.app
Single-file Swift, no dependencies. Requires macOS 13+ and the Xcode command line tools.
If even the hollow ring is too much, Settings → Hide When All Systems Operational
removes it from the menu bar entirely while status is green, so the app only appears
when something is actually wrong.
A gray "check failed" status still shows — silence should never be mistaken for health.
Since a hidden app has no menu to click, there are two ways back:
Launch the app again (double-click it in Finder). The dot reappears for 15
seconds so you can open the menu, then hides itself again.
From the command line:
ClaudeStatus.app/Contents/MacOS/ClaudeStatus --hide-when-green # show setting
ClaudeStatus.app/Contents/MacOS/ClaudeStatus --hide-when-green off # bring it back
A running app picks this up on its next poll, within 60 seconds.
Settings → Start at Login in the app's menu. If macOS parks the request behind an
approval, the app opens the Login Items pane in System Settings for you.
You can also toggle it from the command line:
ClaudeStatus.app/Contents/MacOS/ClaudeStatus --login-status
ClaudeStatus.app/Contents/MacOS/ClaudeStatus --login on
ClaudeStatus.app/Contents/MacOS/ClaudeStatus --login off
How it works
Reads the standard Statuspage summary endpoint:
https://status.claude.com/api/v2/summary.json
The menu bar dot maps status.indicator ( none / minor / major / critical /
maintenance ) to a color; each row in the menu maps that component's own status
( operational / degraded_performance / partial_outage / major_outage /
under_maintenance ) the same way.
Public domain — the Unlicense . No copyright. Clone it, fork it,
modify it, ship it, sell it, no attribution needed. See LICENSE .
Tiny macOS menu bar app showing status.claude.com as a green/yellow/red dot
Readme Unlicense Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
