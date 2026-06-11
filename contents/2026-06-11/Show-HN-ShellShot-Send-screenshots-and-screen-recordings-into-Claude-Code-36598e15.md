---
source: "https://github.com/APIANT/shellshot"
hn_url: "https://news.ycombinator.com/item?id=48489832"
title: "Show HN: ShellShot – Send screenshots and screen recordings into Claude Code"
article_title: "GitHub - APIANT/shellshot: A macOS menubar app that sends screenshots and short recordings directly into a running Claude Code iTerm2 terminal session. Also supports sending screenshots from an iPad/iPhone on the same WiFi network. · GitHub"
author: "rcyeager"
captured_at: "2026-06-11T13:27:08Z"
capture_tool: "hn-digest"
hn_id: 48489832
score: 1
comments: 1
posted_at: "2026-06-11T13:05:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ShellShot – Send screenshots and screen recordings into Claude Code

- HN: [48489832](https://news.ycombinator.com/item?id=48489832)
- Source: [github.com](https://github.com/APIANT/shellshot)
- Score: 1
- Comments: 1
- Posted: 2026-06-11T13:05:57Z

## Translation

タイトル: Show HN: ShellShot – スクリーンショットと画面録画をクロード コードに送信
記事のタイトル: GitHub - APIANT/shellshot: 実行中の Claude Code iTerm2 ターミナル セッションにスクリーンショットと短い記録を直接送信する macOS メニューバー アプリ。同じ WiFi ネットワーク上の iPad/iPhone からのスクリーンショットの送信もサポートします。 · GitHub
説明: スクリーンショットと短い録音を実行中のクロード コード iTerm2 ターミナル セッションに直接送信する macOS メニューバー アプリ。同じ WiFi ネットワーク上の iPad/iPhone からのスクリーンショットの送信もサポートします。 - APIANT/シェルショット

記事本文:
GitHub - APIANT/shellshot: 実行中の Claude Code iTerm2 ターミナル セッションにスクリーンショットと短い録音を直接送信する macOS メニューバー アプリ。同じ WiFi ネットワーク上の iPad/iPhone からのスクリーンショットの送信もサポートします。 · GitHub
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

アピアント
/
シェルショット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット アプリ アプリ メディア メディア プロトタイプ プロトタイプ ツール ツール .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
実行中の Claude Code iTerm2 ターミナル セッションにスクリーンショットと短い録音を直接送信する macOS メニューバー アプリ。同じ WiFi ネットワーク上の iPad/iPhone からのスクリーンショットの送信もサポートします。
⬇ Apple Silicon (v0.1.0) をダウンロードするか、ソースからビルドします。
⌥⌘C — エリアキャプチャ、プレビュー、矢印による注釈付け、送信
⌥⌘R — 領域をフレーム シーケンスとして記録します (知覚的に重複排除、最大 20 フレーム)
すべてのライブ クロード コード セッションにわたるセッション ピッカー。フォーカスされたものを事前に選択します
「クリップボードにコピー」ボタンは、注釈付きキャプチャを送信する代わりにフル解像度でコピーします。
iPad/iPhone 共有 - 自動生成されたショートカットを介して、同じ Wi-Fi 上の iPad または iPhone から Mac セッションにスクリーンショットを送信します。
クロードに送信される画像は、トークンの使用量を適切に保つために縮小されます。
ログイン時に起動します。 7日後の自動プルーンをキャプチャします
macOS + iTerm2 のみをサポートします。 tmux、Kitty、および WezTerm アダプターが可能です。
macOS 13 以降および iTerm2 で [設定] → [一般] → [Magic] → [Python API を有効にする] が必要です。
最新のリリースを取得し、解凍して、ShellShot.app を /Applications に移動します。ビルドはアドホック署名されている (公証されていない) ため、Gatekeeper の隔離を一度クリアします。
xattr -dr com.apple.quarantine /Applications/ShellShot.app
ソースからビルドする
Python 3 が必要です (ビルド時のみ。アプリには自己完結型のサイドカーが埋め込まれます)。
git clone https://github.com/APIANT/shellshot ~ /shellshot
cd ~ /shellshot
python3 -m venv pr

ototype/.venv
プロトタイプ/.venv/bin/pip iterm2 pyinstaller をインストールします
./app/build-app.sh
/アプリケーション/ShellShot.appを開きます
macOS では、画面録画のプロンプトが 1 回、iTerm2 自動化のプロンプトが 1 回表示されます。インストールされたアプリは自己完結型です。
iPad または iPhone (Mac と同じ Wi-Fi 上) からスクリーンショットを Claude Code セッションに送信します。
ShellShot メニューで iPad 共有を有効にし、iPad ショートカットの設定…を選択します。Mac のアドレスとトークンが焼き付けられた QR コードが表示されます。
デバイスで、カメラを QR コードに向け、バナーをタップし、 [ショートカットを追加] をタップします。
使用するには: スクリーンショットを撮り、[共有] をタップし、[ShellShot] を選択します。どのクロード セッションに送信するかを選択し (フォーカスされているセッションには ● マークが付いています)、オプションのメッセージを追加すると、Mac に投稿されます。
Mac はこのために、トークンで保護された小さな HTTP リスナー (ポート 8472) を実行します。デフォルトではオフになっており、LAN のみで使用されます。画像はローカル キャプチャと同じ方法で縮小されます。
実行中の Claude Code iTerm2 ターミナル セッションにスクリーンショットと短い録音を直接送信する macOS メニューバー アプリ。同じ WiFi ネットワーク上の iPad/iPhone からのスクリーンショットの送信もサポートします。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
シェルショット v0.1.0 (Apple シリコン)
最新の
2026 年 6 月 10 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A macOS menubar app that sends screenshots and short recordings directly into a running Claude Code iTerm2 terminal session. Also supports sending screenshots from an iPad/iPhone on the same WiFi network. - APIANT/shellshot

GitHub - APIANT/shellshot: A macOS menubar app that sends screenshots and short recordings directly into a running Claude Code iTerm2 terminal session. Also supports sending screenshots from an iPad/iPhone on the same WiFi network. · GitHub
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
APIANT
/
shellshot
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits app app media media prototype prototype tools tools .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
A macOS menubar app that sends screenshots and short recordings directly into a running Claude Code iTerm2 terminal session. Also supports sending screenshots from an iPad/iPhone on the same WiFi network.
⬇ Download for Apple Silicon (v0.1.0) — or build from source .
⌥⌘C — area capture, preview, annotate with arrows, send
⌥⌘R — record a region as a frame sequence (perceptually deduped, max 20 frames)
Session picker across all live Claude Code sessions; preselects the focused one
"Copy to Clipboard" button copies the annotated capture at full resolution instead of sending
iPad/iPhone sharing — send a screenshot from an iPad or iPhone on the same Wi-Fi into a Mac session via an auto-generated Shortcut
Images sent to Claude are downscaled to keep token usage reasonable
Launch at login; captures auto-prune after 7 days
Supports macOS + iTerm2 only. tmux, Kitty, and WezTerm adapters are possible.
Requires macOS 13+ and iTerm2 with Settings → General → Magic → Enable Python API .
Grab the latest release , unzip, and move ShellShot.app to /Applications . The build is ad-hoc signed (not notarized), so clear Gatekeeper's quarantine once:
xattr -dr com.apple.quarantine /Applications/ShellShot.app
Build from source
Requires Python 3 (build time only; the app embeds a self-contained sidecar).
git clone https://github.com/APIANT/shellshot ~ /shellshot
cd ~ /shellshot
python3 -m venv prototype/.venv
prototype/.venv/bin/pip install iterm2 pyinstaller
./app/build-app.sh
open /Applications/ShellShot.app
macOS will prompt once for Screen Recording and once for iTerm2 automation. The installed app is self-contained.
Send a screenshot from an iPad or iPhone (on the same Wi-Fi as the Mac) into a Claude Code session:
In the ShellShot menu, enable iPad Sharing , then choose Set Up iPad Shortcut… — a QR code appears with the Mac's address and a token baked in.
On the device, point the Camera at the QR code, tap the banner, and Add Shortcut .
To use it: take a screenshot, tap Share , and pick ShellShot . Choose which Claude session to send to (the focused one is marked ●), add an optional message, and it posts to the Mac.
The Mac runs a small token-guarded HTTP listener (port 8472) for this; it's off by default and LAN-only. Images are downscaled the same way as local captures.
A macOS menubar app that sends screenshots and short recordings directly into a running Claude Code iTerm2 terminal session. Also supports sending screenshots from an iPad/iPhone on the same WiFi network.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
ShellShot v0.1.0 (Apple Silicon)
Latest
Jun 10, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
