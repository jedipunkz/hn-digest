---
source: "https://github.com/aalksii/creditwatcher"
hn_url: "https://news.ycombinator.com/item?id=48744006"
title: "Open-source Claude/Codex/Cursor limits tracker for Mac"
article_title: "GitHub - aalksii/creditwatcher: Safe read-only Codex, Claude Code, and Cursor usage CLI · GitHub"
author: "aalksii"
captured_at: "2026-07-01T09:05:52Z"
capture_tool: "hn-digest"
hn_id: 48744006
score: 1
comments: 0
posted_at: "2026-07-01T08:55:15Z"
tags:
  - hacker-news
  - translated
---

# Open-source Claude/Codex/Cursor limits tracker for Mac

- HN: [48744006](https://news.ycombinator.com/item?id=48744006)
- Source: [github.com](https://github.com/aalksii/creditwatcher)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T08:55:15Z

## Translation

タイトル: Mac 用オープンソース クロード/コーデックス/カーソル制限トラッカー
記事のタイトル: GitHub - aalksii/creditwatcher: 安全な読み取り専用コーデックス、クロード コード、およびカーソルの使用 CLI · GitHub
説明: 安全な読み取り専用 Codex、Claude Code、およびカーソル使用 CLI - aalksii/creditwatcher

記事本文:
GitHub - aalksii/creditwatcher: 安全な読み取り専用コーデックス、クロード コード、およびカーソルの使用 CLI · GitHub
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
アールクシー
/
クレジットウォッチャー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラン

ches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
27 コミット 27 コミット .github .github macos macos メディア メディア スクリプト scripts src src .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Codex 、 Claude Code 、および Cursor サブスクリプションの使用制限をローカルで、読み取り専用で、テレメトリなしで監視するための macOS メニュー バー アプリと CLI。
Stats (軽量 macOS メニュー バー ユーティリティ) のデザイン精神からインスピレーションを受けています。
macOS 用 CreditWatcher をダウンロード
事前に構築されたインストーラーは、GitHub Releases を通じて公開されます。
3 つのプロバイダー - Codex (OpenAI)、Claude Code (Anthropic)、1 か所のカーソル
macOS メニュー バー アプリ - ネイティブ Swift、実行時に Node.js は不要
CLI ダッシュボード — 色分けされた進行状況バーを備えた豊富なターミナル ビュー
読み取り専用の使用状況チェック — プロキシ推論や Web UI のスクレイピングは行いません
ローカル認証情報 — 公式ツールから既存のログイン情報を読み取ります。トークンはマシン上に残ります
共有キャッシュ — CLI およびメニュー バー アプリの共有 ~/.creditwatcher/ クォータ キャッシュ
60 秒のクールダウン - オンデマンドの更新により、プロバイダー API のハンマリングを回避します
オプションのローカル Web UI — CreditWatcher は 127.0.0.1 でのみ機能します
CreditWatcher はローカルファーストになるように設計されています。
テレメトリなし - 分析、クラッシュ レポーター、サードパーティ サーバーなし
トークンのログ記録なし — アクセス/リフレッシュ トークンは通常の操作では出力されません。
メニュー バー アプリにキーチェーンはありません — クロード認証は JSON ファイルと環境変数のみを使用します。アプリから macOS キーチェーン プロンプトが表示されない
オプションの CLI インポート - Creditwatcher ログイン クロードは、ターミナルでキーチェーンを 1 回読み取り、資格情報を ~/.creditwatcher/claude-auth.js にコピーできます。

に
サンドボックスフリーの macOS アプリ — ローカル認証情報ストアを読み取り、プロバイダー API を呼び出すために必要です。 SECURITY.mdを参照
macOS アプリ: macOS 14+ (Sonoma)、Xcode 15+
プロバイダーのログイン: 最初に公式ツールでサインインします ( codex login 、 claude 、 Cursor.app)
git clone https://github.com/aalksii/creditwatcher.git
CDクレジットウォッチャー
npmインストール
npm ビルドを実行する
npm link # オプション — `creditwatcher` を PATH にインストールします
リンクしない場合:
npm run ダッシュボード # リッチ ターミナル ダッシュボード (すべてのプロバイダー)
npm run status # プロバイダーごとの詳細なテキスト出力
npm run quote # JSON 出力 (メニュー バー統合で使用)
macOS メニュー バー アプリを構築する
macos/CreditWatcher.xcodeproj を開く
# 製品 → 実行 (⌘R)
またはコマンドラインから:
xcodebuild -project macos/CreditWatcher.xcodeproj -scheme CreditWatcher -configuration デバッグ ビルド
ビルドされたアプリは、 Xcode DerivedData または xcodebuild を使用する場合の build/ の下にあります。
ローカルの Drag-to-Applications インストーラーの場合:
npm run macos:dmg
DMG は dist/macos/CreditWatcher-<version>.dmg に書き込まれます。
一般に配布する場合は、Apple Developer ID で署名し、公証してください。
SIGN_IDENTITY= " 開発者 ID アプリケーション: あなたの名前 (TEAMID) " \
NOTARY_PROFILE= " 公証ツール プロファイル " \
npm run macos:dmg
xcrun notarytool store-credentials を使用して notarytool プロファイルを一度作成します。
コマンド
説明
クレジットウォッチャーダッシュボード
豊富なターミナル ダッシュボード — すべてのプロバイダー
クレジットウォッチャー ダッシュボード --force
60秒間の使用クールダウンをスキップする
クレジットウォッチャーのステータス [コーデックス|クロード|カーソル]
プロバイダーごとの詳細な使用方法
クレジットウォッチャークォータ --json
機械可読なクォータ JSON
Creditwatcher ログイン [コーデックス|クロード|カーソル]
インポートまたはOAuthログインヘルパー
クレジットウォッチャーサーブ
オプションのローカル Web UI (http://127.0.0.1:9477)
例:
クレジットウォッチャーダッシュボード
クレジットウォッチャーステータスクロード
クレジットウォッチャークォータ --json
macOS メニュー バー アプリ
ラウン

Xcode から CreditWatcher を ch (⌘R) するか、ビルドされた .app を開きます
メニュー バーにゲージ アイコンが表示されます (Dock アイコンは表示されません)。
アイコンをクリックすると、コーデックス、クロード、カーソルの使用カードが表示されたポップオーバーが表示されます
歯車ボタンを使用して、プロバイダーの表示/非表示、カードの並べ替え、ツールごとのサインインまたはサインアウトを行います。
アイコンの色合いは最悪の使用状況を反映します: システムのデフォルト <70%、黄色 70 ～ 90%、赤色 >90%
更新: 開くときに自動更新します。 [更新] ボタンは 60 秒のクールダウンをバイパスします。バックグラウンドは 60 秒ごとに更新されます。
終了: ポップオーバーで「終了」をクリックするか、メニューバーのアイコンを右クリックして「CreditWatcherの終了」を選択します。
ログイン時に起動します: システム設定 → 一般 → ログイン項目 → CreditWatcher を追加します。
CLI ボタン​​: Creditwatcher ダッシュボード --verbose でターミナルを開きます (オプション - ノード CLI が必要です)。
まずは公式ツールでサインインしてください。 CreditWatcher は既存の認証情報を読み取ります。置き換えません。
コーデックスログイン
クレジットウォッチャーステータスコーデックス
認証順序:
~/.codex/auth.json (公式 Codex CLI)
~/.creditwatcher/auth.json (creditwatcher ログイン コーデックス経由)
claude # 必要に応じてクロードコードでサインインします
Creditwatcher ログイン クロード # ~/.creditwatcher/claude-auth.json にインポートします
次に、メニュー バー アプリで [更新] をクリックします。
自動使用状況チェックの認証順序 (認証失敗時には最も新しいトークンが優先されます):
CLAUDE_CODE_OAUTH_TOKEN 環境変数
~/.creditwatcher/claude-auth.json (creditwatcher ログイン claude 経由でコピーをインポート)
メニュー バー アプリは上記のソースのみを読み取ります。認証情報をキーチェーンに保存するクロード コードのインストール (認証情報ファイルなし) の場合は、ターミナルで Creditwatcher login claude を実行します。CLI はキーチェーンを一度読み取って、アプリが使用できるファイルのコピーを保存します。
推奨: Cursor.app 経由でサインインし、次の操作を行います。
クレジットウォッチャーステータスカーソル
認証順序:
CURSOR_SESSION_TOKEN 環境変数
カーソル IDE SQLite 状態 DB — state.vscdb キー カーソル

認証/アクセストークン
~/.creditwatcher/cursor-auth.json (creditwatcher ログイン カーソル経由でコピーをインポート)
CreditWatcher.app (Swift) クレジットウォッチャー CLI (Node.js)
§── CodexProvider §── src/codex/
§── クロードプロバイダ §── src/claude/
━── CursorProvider ━━ src/cursor/
│ │
━━━ 共有キャッシュ ────────┘
~/.creditwatcher/
ネイティブ アプリは、 URLSession 経由でプロバイダー API を直接呼び出します。実行時に Node.js は必要ありません。
変数
デフォルト
説明
CODEX_HOME
~/.コーデックス
Codex CLI 認証ディレクトリ
CLAUDE_CONFIG_DIR
~/.クロード
クロード認証情報ディレクトリ
CURSOR_SESSION_TOKEN
—
カーソルセッションクッキーの値
CURSOR_STATE_DB
プラットフォームのデフォルト
Cursor state.vscdb へのパスをオーバーライドします
CREDITWATCHER_OAUTH_PORT
1455年
OAuth コールバック ポート (Codex ログイン)
免責事項
CreditWatcher は非公式であり、OpenAI、Anthropic、または Cursor によって承認されていません。
使用法 API は予告なく変更される場合があります (特に Cursor の非公式エンドポイント)
コンシューマ OAuth を使用するサードパーティ ツールはプロバイダの利用規約に違反する可能性があります
Anthropic はサードパーティ ツールでのコンシューマ OAuth を制限しています - Claude の統合は自己責任で使用してください
このツールは読み取り専用の使用状況チェックを実行します。プロキシ推論やトークンの共有には使用しないでください。
自己責任で使用してください。作者はプロバイダーによるアカウントのアクションについて責任を負いません
安全な使用パターンの詳細については、この README の「安全性とサービス利用規約」セクションを参照してください。
ビルド済み macOS リリース (署名付き .app / Homebrew Cask)
グローバル CLI インストール用の npm 公開
読み取り専用使用エンドポイントのみ ( GET /wham/usage 、 /api/oauth/usage 、 /api/usage-summary )
必要に応じて、公式トークンエンドポイント経由で期限切れの OAuth トークンを更新します
プロバイダーごとに 60 秒のクールダウンを伴うオンデマンド チェック
プロバイダー API への直接呼び出し - サードパーティのリレーなし

ChatGPT または Claude Web UI をスクレイピング
公式プロバイダー API を除く任意のサーバーにトークンを送信します
60 秒の更新間隔を超えるバックグラウンド ポーリング
アクセス/リフレッシュトークンのログまたは印刷
パターン
なぜ危険なのか
公式クライアントハーネスのなりすまし
サブスクリプション クォータをサードパーティ エージェント経由でルーティングする
資格情報の漏洩
OAuth トークンをサードパーティサーバーに送信する
推論プロキシ
他のユーザー/ツールのサブスクリプション トークンの使用
積極的なポーリング
使用エンドポイントの攻撃
トークンの共有
マシン間でのリフレッシュ トークンの配布
貢献する
COTRIBUTING.md を参照してください。セキュリティの問題は SECURITY.md 経由で報告してください。
MIT — 著作権 (c) 2026 アレクセイ・アルテミエフ
UI のインスピレーション: excelban による統計
公式 Codex CLI パターンに合わせた Codex OAuth フロー
安全な読み取り専用の Codex、Claude Code、およびカーソルの使用 CLI
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
クレジットウォッチャー v0.1.0
最新の
2026 年 6 月 30 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Safe read-only Codex, Claude Code, and Cursor usage CLI - aalksii/creditwatcher

GitHub - aalksii/creditwatcher: Safe read-only Codex, Claude Code, and Cursor usage CLI · GitHub
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
aalksii
/
creditwatcher
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits .github .github macos macos media media scripts scripts src src .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A macOS menu bar app and CLI to monitor Codex , Claude Code , and Cursor subscription usage limits — locally, read-only, with no telemetry.
Inspired by the design spirit of Stats (lightweight macOS menu bar utility).
Download CreditWatcher for macOS
Prebuilt installers are published through GitHub Releases .
Three providers — Codex (OpenAI), Claude Code (Anthropic), Cursor in one place
macOS menu bar app — native Swift, no Node.js required at runtime
CLI dashboard — rich terminal view with color-coded progress bars
Read-only usage checks — never proxies inference or scrapes web UIs
Local credentials — reads existing logins from official tools; tokens stay on your machine
Shared cache — CLI and menu bar app share ~/.creditwatcher/ quota cache
60-second cooldown — on-demand refresh, avoids hammering provider APIs
Optional local web UI — creditwatcher serve on 127.0.0.1 only
CreditWatcher is designed to be local-first:
No telemetry — no analytics, crash reporters, or third-party servers
No token logging — access/refresh tokens are not printed in normal operation
No Keychain in the menu bar app — Claude auth uses JSON files and env vars only; no macOS Keychain prompts from the app
Optional CLI import — creditwatcher login claude may read Keychain once in Terminal to copy credentials into ~/.creditwatcher/claude-auth.json
Sandbox-free macOS app — required to read local credential stores and call provider APIs; see SECURITY.md
macOS app: macOS 14+ (Sonoma), Xcode 15+
Provider logins: sign in with official tools first ( codex login , claude , Cursor.app)
git clone https://github.com/aalksii/creditwatcher.git
cd creditwatcher
npm install
npm run build
npm link # optional — install `creditwatcher` on your PATH
Without linking:
npm run dashboard # rich terminal dashboard (all providers)
npm run status # detailed text output per provider
npm run quota # JSON output (used by menu bar integration)
Build macOS menu bar app
open macos/CreditWatcher.xcodeproj
# Product → Run (⌘R)
Or from the command line:
xcodebuild -project macos/CreditWatcher.xcodeproj -scheme CreditWatcher -configuration Debug build
The built app is under Xcode DerivedData or build/ when using xcodebuild .
For a local drag-to-Applications installer:
npm run macos:dmg
The DMG is written to dist/macos/CreditWatcher-<version>.dmg .
For public distribution, sign and notarize it with an Apple Developer ID:
SIGN_IDENTITY= " Developer ID Application: Your Name (TEAMID) " \
NOTARY_PROFILE= " notarytool-profile " \
npm run macos:dmg
Create the notarytool profile once with xcrun notarytool store-credentials .
Command
Description
creditwatcher dashboard
Rich terminal dashboard — all providers
creditwatcher dashboard --force
Skip the 60-second usage cooldown
creditwatcher status [codex|claude|cursor]
Detailed usage per provider
creditwatcher quota --json
Machine-readable quota JSON
creditwatcher login [codex|claude|cursor]
Import or OAuth login helpers
creditwatcher serve
Optional local web UI at http://127.0.0.1:9477
Example:
creditwatcher dashboard
creditwatcher status claude
creditwatcher quota --json
macOS menu bar app
Launch CreditWatcher from Xcode (⌘R) or open the built .app
A gauge icon appears in the menu bar (no Dock icon)
Click the icon for a popover with Codex, Claude, and Cursor usage cards
Use the gear button to show/hide providers, reorder cards, and sign in or out per tool
Icon tint reflects worst-case usage: system default <70%, yellow 70–90%, red >90%
Refresh: auto-refresh on open; Refresh button bypasses the 60s cooldown. Background refresh every 60 seconds.
Quit: click Quit in the popover or right-click the menu bar icon and choose Quit CreditWatcher .
Launch at login: System Settings → General → Login Items → add CreditWatcher.
CLI button: opens Terminal with creditwatcher dashboard --verbose (optional — requires the Node CLI).
Sign in with the official tools first. CreditWatcher reads existing credentials — it does not replace them.
codex login
creditwatcher status codex
Auth order:
~/.codex/auth.json (official Codex CLI)
~/.creditwatcher/auth.json (via creditwatcher login codex )
claude # sign in with Claude Code if needed
creditwatcher login claude # import into ~/.creditwatcher/claude-auth.json
Then click Refresh in the menu bar app.
Auth order for automatic usage checks (freshest token wins on auth failure):
CLAUDE_CODE_OAUTH_TOKEN environment variable
~/.creditwatcher/claude-auth.json (import copy via creditwatcher login claude )
The menu bar app reads only the sources above. For Claude Code installs that store credentials in Keychain (no credentials file), run creditwatcher login claude in Terminal — the CLI may read Keychain once there and save a file copy the app can use.
Recommended: sign in via Cursor.app, then:
creditwatcher status cursor
Auth order:
CURSOR_SESSION_TOKEN environment variable
Cursor IDE SQLite state DB — state.vscdb key cursorAuth/accessToken
~/.creditwatcher/cursor-auth.json (import copy via creditwatcher login cursor )
CreditWatcher.app (Swift) creditwatcher CLI (Node.js)
├── CodexProvider ├── src/codex/
├── ClaudeProvider ├── src/claude/
└── CursorProvider └── src/cursor/
│ │
└──────── shared cache ──────────────┘
~/.creditwatcher/
Native app calls provider APIs directly via URLSession . No Node.js required at runtime.
Variable
Default
Description
CODEX_HOME
~/.codex
Codex CLI auth directory
CLAUDE_CONFIG_DIR
~/.claude
Claude credentials directory
CURSOR_SESSION_TOKEN
—
Cursor session cookie value
CURSOR_STATE_DB
platform default
Override path to Cursor state.vscdb
CREDITWATCHER_OAUTH_PORT
1455
OAuth callback port (Codex login)
Disclaimer
CreditWatcher is unofficial and not endorsed by OpenAI, Anthropic, or Cursor.
Usage APIs may change without notice (especially Cursor's unofficial endpoints)
Third-party tools using consumer OAuth may violate provider Terms of Service
Anthropic has restricted consumer OAuth in third-party tools — use Claude integration at your own risk
This tool performs read-only usage checks — do not use it to proxy inference or share tokens
Use at your own risk — the authors are not responsible for account actions by providers
See the Safety & Terms of Service section in this README for details on safe usage patterns.
Pre-built macOS release (signed .app / Homebrew cask)
npm publish for global CLI install
Read-only usage endpoints only ( GET /wham/usage , /api/oauth/usage , /api/usage-summary )
Refreshes expired OAuth tokens via official token endpoints when needed
On-demand checks with a 60-second cooldown per provider
Direct calls to provider APIs — no third-party relay
Scrape ChatGPT or Claude web UIs
Send tokens to any server except the official provider APIs
Background polling beyond the 60s refresh interval
Log or print access/refresh tokens
Pattern
Why it's risky
Spoofing official client harness
Routing subscription quota through third-party agents
Credential exfiltration
Sending OAuth tokens to third-party servers
Inference proxying
Using subscription tokens for other users/tools
Aggressive polling
Hammering usage endpoints
Token sharing
Distributing refresh tokens across machines
Contributing
See CONTRIBUTING.md . Please report security issues via SECURITY.md .
MIT — Copyright (c) 2026 Aleksei Artemiev
UI inspiration: Stats by exelban
Codex OAuth flow aligned with the official Codex CLI patterns
Safe read-only Codex, Claude Code, and Cursor usage CLI
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
CreditWatcher v0.1.0
Latest
Jun 30, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
