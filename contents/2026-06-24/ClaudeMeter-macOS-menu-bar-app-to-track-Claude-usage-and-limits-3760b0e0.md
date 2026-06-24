---
source: "https://github.com/yotake/claude-meter"
hn_url: "https://news.ycombinator.com/item?id=48653089"
title: "ClaudeMeter – macOS menu bar app to track Claude usage and limits"
article_title: "GitHub - yotake/claude-meter: macOS menu bar app showing Claude usage & limits (session/weekly, Codex rate limits, API spend). Multi-account, JP/EN. · GitHub"
author: "chibeeMan"
captured_at: "2026-06-24T00:27:34Z"
capture_tool: "hn-digest"
hn_id: 48653089
score: 2
comments: 0
posted_at: "2026-06-23T23:28:59Z"
tags:
  - hacker-news
  - translated
---

# ClaudeMeter – macOS menu bar app to track Claude usage and limits

- HN: [48653089](https://news.ycombinator.com/item?id=48653089)
- Source: [github.com](https://github.com/yotake/claude-meter)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T23:28:59Z

## Translation

タイトル: ClaudeMeter – クロードの使用量と制限を追跡する macOS メニュー バー アプリ
記事タイトル: GitHub - yotake/claude-meter: クロードの使用量と制限 (セッション/週、Codex レート制限、API 使用量) を表示する macOS メニュー バー アプリ。マルチアカウント、JP/EN。 · GitHub
説明: クロードの使用量と制限 (セッション/週、Codex レート制限、API 使用量) を表示する macOS メニュー バー アプリ。マルチアカウント、JP/EN。 - yotake/クロードメーター

記事本文:
GitHub - yotake/claude-meter: クロードの使用量と制限 (セッション/週、Codex レート制限、API 使用量) を表示する macOS メニュー バー アプリ。マルチアカウント、JP/EN。 · GitHub
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
よたけ
/
クロードメーター
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット .github .github パッケージング パッケージング ソース/ ClaudeMeter ソース/ ClaudeMeter アセット アセット .gitignore .gitignore ライセンス ライセンス Package.swift Package.swift README.ja.md README.ja.md README.md README.md build.sh build.sh release.sh release.sh すべてのファイルを表示 リポジトリ ファイルナビゲーション
⬇️ クロードメーター (.dmg) をダウンロード
日本語 README はこちら / 日本語 README
クロードの使用状況をリアルタイムで表示する軽量の macOS メニュー バー アプリ。
クロードメーターは無料です。時間を節約できる場合は、開発のスポンサーになることができます ❤️ 。
UI は macOS の言語 (日本語/英語) に従います。
claude.ai → [設定] → [使用状況] と同じ番号がメニュー バーに表示されます。
現在のセッション - リセットまでの時間を含む 5 時間のローリング ウィンドウの使用率
週間制限 — 全モデル / Sonnet のみ / Opus のみ、リセット日あり
Codex レート制限 (オプション) — ローカル Codex CLI ログから読み取ります
API 支出 (オプション) — クロード管理者キーによる今月の支出
複数のアカウント - 複数のサブスクリプション/キーを一度に追跡
メニュー バーには、選択したアカウントのセッション % (例: 21% ) が色付きで表示されます。
燃焼率の予測。アイコンをクリックすると完全な内訳が表示されます。
https://api.anthropic.com/api/oauth/usage を 5 分ごとにポーリングします。
サブスクリプションの使用量 (5 時間のセッション + 7 日間の制限)。そのエンドポイントには、
user:profile スコープの OAuth トークン。
ClaudeMeter は macOS キーチェーンに直接アクセスしません (アドホック署名された
アプリがキーチェーンを読み取ると、ウイルス対策の警告がトリガーされる可能性があります)。代わりに、
クロード CLI がすでに管理している OAuth トークンがポップオーバーに一度追加されます。
貼り付けられたトークンは次の場所に保存されます。
~/ライブラリ/アプリケーション サポート/ClaudeMeter/credentials.json

(モード0600)および
アプリはリフレッシュ トークンを使用して自動更新するため、通常は更新トークンを使用することはありません。
もう一度貼り付けます。
注: 最初の自動リフレッシュではリフレッシュ トークンがローテーションされるため、クロード
CLI では、次回もログインするように求められる場合があります。その後、ClaudeMeter と CLI
トークンを独立して管理します。
使用エンドポイントは、積極的なポーリングを抑制します。 429 では、アプリは待機します。
サーバーの再試行後 (+60 秒)、自動的に回復します。更新をスパム送信しないでください。
Claude Max / Pro サブスクリプション (使用エンドポイントに到達するために必要)
(ソースからのビルドのみ) Swift コマンド ライン ツール — xcode-select --install
ClaudeMeter-<バージョン>.dmg を次からダウンロードします。
をリリースして開き、
ClaudeMeter.app をアプリケーションにドラッグします。
リリースされたビルドは現在アドホック署名されている (公証されていない) ため、
ゲートキーパーが最初の起動をブロックします。次のいずれかを使用して許可します。
ClaudeMeter.appを右クリック→開く→ダイアログで開く
システム設定 → プライバシーとセキュリティ → とにかく開く
ターミナルから隔離属性を削除します。
xattr -d com.apple.quarantine /Applications/ClaudeMeter.app
メニュー バー アイコンをクリックし、トークンを貼り付けます (下記のセットアップを参照)。
すでに claude CLI を使用している場合は、そのトークンをクリップボードにコピーします。
security find-generic-password -s " Claude Code-credentials " -w | security find-generic-password -s " Claude Code-credentials " -w | pbcopy
セキュリティは、Apple が署名した組み込み CLI です。キーチェーン プロンプトが表示された場合は、
許可します (ClaudeMeter 自体がキーチェーンに触れることはありません)。次に、
ポップオーバーをクリックしてフィールドに貼り付け、「保存」を押します。それは示します
「認証済み (自動更新)」となり、それ以降はトークンが更新されます。
API の使用量を追跡するには、代わりに Claude Admin キー ( sk-ant-admin… ) を貼り付けます。
ログが保存されていない場合は、設定でカスタム Codex セッション フォルダーを設定します。
デフォルトの ~/.codex/sessions (CODEX_HOME 環境変数も考慮されます)。
./build.sh
ClaudeMeter.app を開く
ログイン時に開始
システム設定 →

「一般」→「ログイン項目」→「ClaudeMeter.app」を追加します。
release.sh は、ユニバーサル (Apple Silicon + Intel) DMG を dist/ に構築します。
./release.sh # アドホック署名済み DMG (ゲートキーパーがユーザーに警告)
ゲートキーパーの警告なしで公証された DMG を作成するには (Apple の許可が必要です)
開発者プログラム メンバーシップ)、署名 ID と公証人の資格情報を設定します
そして再実行します — 正確な環境変数については release.sh のヘッダーを参照してください。
SIGN_ID= " 開発者 ID アプリケーション: あなたの名前 (TEAMID) " \
NOTARY_PROFILE= " クロードメーター " \
./release.sh
シークレットはリポジトリに書き込まれません。認証情報は環境から読み取られます。
/ 実行時のみのログイン キーチェーン。
ClaudeMeter は無料のオープンソースです。お役に立てば、ちょっとしたヒントを
GitHub スポンサー ❤️ は、
Apple Developer メンバーシップを保持し、それを維持します。
プライベート API ( /api/oauth/usage ) を使用し、ローカル CLI トークン/ログを読み取ります。
そのため、Anthropic/OpenAI の変更によってそれが壊れる可能性があります。
トークンの有効期限が切れてエラーが表示された場合は、新しいトークンを取得します (claude を 1 回実行し、
それからセキュリティ… | pbcopy を再度)、Update token を使用して再貼り付けします。
Claude の使用量と制限 (セッション/週、Codex レート制限、API 使用量) を表示する macOS メニュー バー アプリ。マルチアカウント、JP/EN。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
クロードメーター 1.1
最新の
2026 年 6 月 18 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

macOS menu bar app showing Claude usage & limits (session/weekly, Codex rate limits, API spend). Multi-account, JP/EN. - yotake/claude-meter

GitHub - yotake/claude-meter: macOS menu bar app showing Claude usage & limits (session/weekly, Codex rate limits, API spend). Multi-account, JP/EN. · GitHub
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
yotake
/
claude-meter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .github .github Packaging Packaging Sources/ ClaudeMeter Sources/ ClaudeMeter assets assets .gitignore .gitignore LICENSE LICENSE Package.swift Package.swift README.ja.md README.ja.md README.md README.md build.sh build.sh release.sh release.sh View all files Repository files navigation
⬇️ Download ClaudeMeter (.dmg)
日本語 README はこちら / Japanese README
A lightweight macOS menu bar app that shows your Claude usage in real time.
ClaudeMeter is free. If it saves you time, you can sponsor development ❤️ .
The UI follows your macOS language (Japanese / English).
The same numbers as claude.ai → Settings → Usage , right in your menu bar:
Current session — 5-hour rolling window utilization, with time until reset
Weekly limits — all models / Sonnet only / Opus only, with reset day
Codex rate limits (optional) — read from your local Codex CLI logs
API spend (optional) — this month's spend via a Claude Admin key
Multiple accounts — track several subscriptions/keys at once
The menu bar shows the selected account's session % (e.g. 21% ), tinted by a
burn-rate forecast. Click the icon for the full breakdown.
Polls https://api.anthropic.com/api/oauth/usage every 5 minutes for
subscription usage (5-hour session + 7-day limits). That endpoint requires an
OAuth token with the user:profile scope .
ClaudeMeter does not access the macOS Keychain directly (an ad-hoc–signed
app reading the Keychain can trigger antivirus warnings). Instead, you paste the
OAuth token that the claude CLI already manages, once, into the popover.
The pasted token is stored at
~/Library/Application Support/ClaudeMeter/credentials.json (mode 0600 ) and
the app auto-refreshes it using the refresh token, so you normally never
paste again.
Note: the first auto-refresh rotates the refresh token, so the claude
CLI may ask you to log in again next time. After that, ClaudeMeter and the CLI
manage tokens independently.
The usage endpoint throttles aggressive polling. On a 429 the app waits the
server's Retry-After (+60s) and recovers automatically — don't spam refresh.
A Claude Max / Pro subscription (required to reach the usage endpoint)
(build from source only) Swift Command Line Tools — xcode-select --install
Download ClaudeMeter-<version>.dmg from
Releases , open it, and
drag ClaudeMeter.app to Applications.
The released build is currently ad-hoc signed (not notarized) , so
Gatekeeper blocks the first launch. Allow it with any of:
Right-click ClaudeMeter.app → Open → Open in the dialog
System Settings → Privacy & Security → Open Anyway
Remove the quarantine attribute from Terminal:
xattr -d com.apple.quarantine /Applications/ClaudeMeter.app
Click the menu bar icon and paste your token (see Setup below).
If you already use the claude CLI, copy its token to the clipboard:
security find-generic-password -s " Claude Code-credentials " -w | pbcopy
security is Apple's signed, built-in CLI. If a Keychain prompt appears, click
Allow (ClaudeMeter itself never touches the Keychain). Then open the
popover, paste into the field, and press Save . It shows
"Authenticated (auto-refresh)" and refreshes the token for you from then on.
Paste a Claude Admin key ( sk-ant-admin… ) instead to track API spend.
Set a custom Codex sessions folder in Settings if your logs aren't in the
default ~/.codex/sessions (it also honors the CODEX_HOME env var).
./build.sh
open ClaudeMeter.app
Start at login
System Settings → General → Login Items → add ClaudeMeter.app .
release.sh builds a universal (Apple Silicon + Intel) DMG into dist/ :
./release.sh # ad-hoc signed DMG (Gatekeeper warns users)
To produce a notarized DMG with no Gatekeeper warning (needs an Apple
Developer Program membership), set the signing identity and notary credentials
and re-run — see the header of release.sh for the exact env vars:
SIGN_ID= " Developer ID Application: Your Name (TEAMID) " \
NOTARY_PROFILE= " claude-meter " \
./release.sh
No secrets are written to the repo — credentials are read from the environment
/ login keychain at run time only.
ClaudeMeter is free and open source. If it's useful to you, a small tip via
GitHub Sponsors ❤️ helps cover the
Apple Developer membership and keeps it maintained.
It uses a private API ( /api/oauth/usage ) and reads local CLI tokens/logs,
so an Anthropic/OpenAI change can break it.
If a token expires and you see an error, grab a fresh one (run claude once,
then security … | pbcopy again) and re-paste via Update token .
macOS menu bar app showing Claude usage & limits (session/weekly, Codex rate limits, API spend). Multi-account, JP/EN.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
ClaudeMeter 1.1
Latest
Jun 18, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
