---
source: "https://github.com/SGSI/claudeometer"
hn_url: "https://news.ycombinator.com/item?id=48822355"
title: "Show HN: Claudeometer – Share and pool Claude usage across a team (macOS, OSS)"
article_title: "GitHub - SGSI/claudeometer: Share & pool Claude usage across your team — a free macOS menu-bar app with live quota, burn-rate alerts, and end-to-end-encrypted borrowing. · GitHub"
author: "WableSanket"
captured_at: "2026-07-07T19:43:03Z"
capture_tool: "hn-digest"
hn_id: 48822355
score: 1
comments: 0
posted_at: "2026-07-07T19:17:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claudeometer – Share and pool Claude usage across a team (macOS, OSS)

- HN: [48822355](https://news.ycombinator.com/item?id=48822355)
- Source: [github.com](https://github.com/SGSI/claudeometer)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T19:17:49Z

## Translation

タイトル: HN を表示: Claudeometer – チーム全体でクロードの使用状況を共有およびプールする (macOS、OSS)
記事のタイトル: GitHub - SGSI/claudeometer: クロードの使用状況をチーム全体で共有してプールする — ライブ クォータ、バーンレート アラート、エンドツーエンドで暗号化された借用を備えた無料の macOS メニュー バー アプリ。 · GitHub
説明: クロードの使用状況をチーム全体で共有およびプールします。ライブ クォータ、バーンレート アラート、エンドツーエンドで暗号化された借用を備えた無料の macOS メニュー バー アプリです。 - SGSI/クラウデメーター

記事本文:
GitHub - SGSI/claudeometer: クロードの使用状況をチーム全体で共有およびプールします。ライブ クォータ、バーンレート アラート、エンドツーエンドで暗号化された借用を備えた無料の macOS メニュー バー アプリです。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。これをリロードしてください

ページに
SGSI
/
クラウデメーター
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
31 コミット 31 コミット ソース ソース Tests/ ClaudeUsageBarCoreTests Tests/ ClaudeUsageBarCoreTests アセット アセット ドキュメント ドキュメント リレー リレー スクリプト スクリプト .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス Package.swift Package.swift README.md README.md claude-usage-swiftbar.js claude-usage-swiftbar.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロードの使用状況をバッテリーのように表示するネイティブ macOS メニューバー アプリ — ライブ
クォータ、バーンレート、および段階的なアラートを備えているため、制限に驚かれることはありません。
⌁ 62% 🙂 ← メニューバーで、登るにつれて色が緑→赤に変化します
何をするのか
メニュー バーのライブ クォータ - 5 時間のローリング ウィンドウをパーセンテージで表示します。
気分の絵文字と、使用量が増えるにつれて色が緑から赤に徐々に変化します。
詳細なポップオーバー — 以下のメニューバー アイコンをクリックします。
燃焼率 ( +X%/hr ) と ETA がフルの 5 時間ウィンドウ ヒーロー、
7-day 、 Sonnet 、 Opus 、および OAuth アプリの割り当て、
24 時間ペースのスパークライン (ローカルに保存された履歴から構築)、
Claude Code · ホット セッション — 過去 5 時間で最もトークンを大量に使用したセッション、
追加使用費用 (有効な場合)。
段階的な通知 — 50% 、 75% 、 90% (パニック) 、および 100% でのアラート
5 時間ウィンドウのうち、ウィンドウごとに 1 回ずつ実行されます。
ローカル履歴 - 使用状況のスナップショットは最大 30 日間ローカルに保存されます
( ~/Library/Application Support/Claudeometer/history.json ) および自動的にプルーニングされます。
すべてがマシン上に残ります。アプリはテレメトリをどこにも送信しません。
醸造タップ SGSI/クラウデメーター
brew trust sgsi/claudeometer # one-time: Homebrew 6.0 以降ではサードパーティのタップを信頼する必要があります
醸造インストール

l --cask claudeometer
Homebrew はインストール中に macOS 隔離フラグをクリアするため、アプリは何も表示されずに開きます。
ゲートキーパーの警告。
最新の Claudeometer.dmg を次の場所からダウンロードします。
リリースページを開き、
Claudeometer を [アプリケーション] にドラッグします。
アプリはまだ署名/公証されていないため、生のダウンロードにより macOS Gatekeeper がトリガーされます。
最初の打ち上げ。これを回避するには、次のコマンドを 1 回実行します。
xattr -dr com.apple.quarantine /Applications/Claudeometer.app
…または、それを開いてみて、 [システム設定] → [プライバシーとセキュリティ] → [とにかく開く] に移動します。
macOS の初回起動時に、Claude Code-credentials キーチェーンへのアクセスを要求される場合があります。
項目 — 許可します。アプリにトークンの有効期限が切れていると表示された場合は、[ログイン] (•••
メニュー) を実行するか、ターミナルで claude を一度実行します。
Claudeometer は、Claude Code が macOS キーチェーンに保存する OAuth トークンを読み取ります。
Claude Code-credentials の下で、次を呼び出します。
https://api.anthropic.com/api/oauth/usage
https://api.anthropic.com/api/oauth/profile
トークンとクォータのデータがマシンから流出することはありません。
⚠️ これらの人類のエンドポイントは文書化されておらず、いつでも変更または破損する可能性があります。
Claudeometer は非公式ツールであり、Anthropic とは提携していません。
「ホット セッション」リストとペース履歴は、ローカルのクロード コード ログから取得されます。
~/.claude/projects 内。
Claudeometer は、チーム全体で Claude の使用状況を共有およびプールすることもできます: 作成または参加
チーム (パブリックまたはプライベート、名前とパスワード付き) では、全員のライブ使用状況を
チームごとのボード、およびチームメイトのクロード クォータを固定ウィンドウで借用できます。
アウト。借用の範囲は共有チームに限定され、貸与される資格情報はエンドツーエンドで行われます。
借り手に暗号化される - リレーは不透明な暗号文を仲介するだけで、仲介することは決してない
トークンが見えます。
チーム モードは、小規模な自己ホスト型リレー (relay/ — Go + SQLite サービス、を参照) と通信します。
リレー/PROTOCOL.md )。相対

URL は意図的にこのリポジトリに組み込まれていません。ポイント
次のいずれかを使用して独自のリレーでアプリを実行します。
環境変数: CLAUDEOMETER_RELAY_URL=http://your-relay:8080 、または
~/Library/Application Support/Claudeometer/relay-url にある 1 行のファイルに URL が含まれています。
どちらの設定も行わない場合、Claudeometer は個人の使用量メーターのままであり、チーム プロンプトもネットワークもありません。
自分でリレーを主催したくないですか?連絡してください。私のものを共有します — で連絡してください
LinkedIn または sanketwable123@gmail.com に電子メールを送信してください。ホストされたリレーの URL を送信します。
アプリをポイントします。
要件: macOS 13 以降および最新の Swift ツールチェーン (Xcode 16 以降)。
chmod +x スクリプト/build-app.sh
./scripts/build-app.sh
build/Claudeometer.app を開く
アプリバンドルは build/Claudeometer.app で作成されます。
クロードの使用状況をチーム全体で共有およびプールします。ライブ クォータ、バーンレート アラート、エンドツーエンドで暗号化された借用を備えた無料の macOS メニュー バー アプリです。
sgsi.github.io/claudeometer/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
6
クロードメーター 0.2.4
最新の
2026 年 7 月 2 日
+ 5 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Share & pool Claude usage across your team — a free macOS menu-bar app with live quota, burn-rate alerts, and end-to-end-encrypted borrowing. - SGSI/claudeometer

GitHub - SGSI/claudeometer: Share & pool Claude usage across your team — a free macOS menu-bar app with live quota, burn-rate alerts, and end-to-end-encrypted borrowing. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
SGSI
/
claudeometer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
31 Commits 31 Commits Sources Sources Tests/ ClaudeUsageBarCoreTests Tests/ ClaudeUsageBarCoreTests assets assets docs docs relay relay scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE Package.swift Package.swift README.md README.md claude-usage-swiftbar.js claude-usage-swiftbar.js View all files Repository files navigation
A native macOS menu-bar app that shows your Claude usage like a battery — live
quota, burn rate, and graduated alerts so you never get surprised by a limit.
⌁ 62% 🙂 ← in your menu bar, color shifts green → red as you climb
What it does
Live quota in the menu bar — the 5-hour rolling window as a percentage, with a
mood emoji and a color that gradually shifts from green to red as usage climbs.
Detailed popover — click the menu-bar icon for:
the 5-hour window hero with burn rate ( +X%/hr ) and an ETA to full,
7-day , Sonnet , Opus , and OAuth apps quotas,
a 24-hour pace sparkline (built from locally-stored history),
Claude Code · Hot sessions — your most token-heavy sessions in the last 5h,
extra-usage spend (when enabled).
Graduated notifications — alerts at 50% , 75% , 90% (panic) , and 100%
of the 5-hour window, each firing once per window.
Local history — usage snapshots are stored locally for up to 30 days
( ~/Library/Application Support/Claudeometer/history.json ) and pruned automatically.
Everything stays on your machine. The app sends no telemetry anywhere.
brew tap SGSI/claudeometer
brew trust sgsi/claudeometer # one-time: Homebrew 6.0+ requires trusting third-party taps
brew install --cask claudeometer
Homebrew clears the macOS quarantine flag during install, so the app opens with no
Gatekeeper warning .
Download the latest Claudeometer.dmg from the
releases page , open it, and
drag Claudeometer to Applications .
The app is not yet signed/notarized, so a raw download triggers macOS Gatekeeper on
first launch. To get past it, run once:
xattr -dr com.apple.quarantine /Applications/Claudeometer.app
…or try to open it, then go to System Settings → Privacy & Security → Open Anyway .
On first launch macOS may ask for access to the Claude Code-credentials Keychain
item — allow it. If the app says the token is expired, click Login (in the •••
menu) or run claude in a terminal once.
Claudeometer reads the OAuth token that Claude Code stores in the macOS Keychain
under Claude Code-credentials , then calls:
https://api.anthropic.com/api/oauth/usage
https://api.anthropic.com/api/oauth/profile
The token and quota data never leave your machine.
⚠️ These Anthropic endpoints are undocumented and may change or break at any time.
Claudeometer is an unofficial tool and is not affiliated with Anthropic.
The "Hot sessions" list and pace history are derived from your local Claude Code logs
in ~/.claude/projects .
Claudeometer can also share and pool Claude usage across your teams : create or join
teams (public or private, with a name and password), see everyone's live usage on a
per-team board, and borrow a teammate's Claude quota for a fixed window when you're
out. Borrowing is scoped to teams you share, and a lent credential is end-to-end
encrypted to the borrower — the relay only ever brokers opaque ciphertext and never
sees a token.
Team mode talks to a small self-hosted relay ( relay/ — a Go + SQLite service; see
relay/PROTOCOL.md ). The relay URL is deliberately not baked into this repo. Point
the app at your own relay with either:
an environment variable: CLAUDEOMETER_RELAY_URL=http://your-relay:8080 , or
a one-line file at ~/Library/Application Support/Claudeometer/relay-url containing the URL.
With neither set, Claudeometer stays a personal usage meter — no team prompts, no network.
Don't want to host your own relay? Reach out and I'll share mine — connect with me on
LinkedIn or email sanketwable123@gmail.com , and I'll send you a hosted relay URL to
point the app at.
Requirements: macOS 13+ and a recent Swift toolchain (Xcode 16+).
chmod +x scripts/build-app.sh
./scripts/build-app.sh
open build/Claudeometer.app
The app bundle is produced at build/Claudeometer.app .
Share & pool Claude usage across your team — a free macOS menu-bar app with live quota, burn-rate alerts, and end-to-end-encrypted borrowing.
sgsi.github.io/claudeometer/
Topics
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
6
Claudeometer 0.2.4
Latest
Jul 2, 2026
+ 5 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
