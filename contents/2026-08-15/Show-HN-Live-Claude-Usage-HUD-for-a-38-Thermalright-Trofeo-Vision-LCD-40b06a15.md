---
source: "https://github.com/christensen143/claude-trofeo-hud"
hn_url: "https://news.ycombinator.com/item?id=49314594"
title: "Show HN: Live Claude Usage HUD for a $38 Thermalright Trofeo Vision LCD"
article_title: "GitHub - christensen143/claude-trofeo-hud: Live Claude usage HUD for a $38 Thermalright Trofeo Vision LCD — session/weekly limits, token burn, and hypothetical API cost streamed from macOS over USB. · GitHub"
author: "christensen143"
captured_at: "2026-08-15T22:10:42Z"
capture_tool: "hn-digest"
hn_id: 49314594
score: 2
comments: 0
posted_at: "2026-08-15T21:42:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Live Claude Usage HUD for a $38 Thermalright Trofeo Vision LCD

- HN: [49314594](https://news.ycombinator.com/item?id=49314594)
- Source: [github.com](https://github.com/christensen143/claude-trofeo-hud)
- Score: 2
- Comments: 0
- Posted: 2026-08-15T21:42:39Z

## Translation

タイトル: HN を表示: 38 ドルの Thermalright Trofeo Vision LCD のライブ クロード使用状況 HUD
記事のタイトル: GitHub - christensen143/claude-trofeo-hud: 38 ドルの Thermalright Trofeo Vision LCD の Live Claude 使用状況 HUD — セッション/週ごとの制限、トークンの書き込み、USB 経由で macOS からストリーミングされる仮想 API コスト。 · GitHub
説明: 38 ドルの Thermalright Trofeo Vision LCD の Live Claude 使用状況 HUD — セッション/週ごとの制限、トークンの書き込み、USB 経由で macOS からストリーミングされた仮想 API コスト。 - christensen143/クロード-トロフェオ-hud

記事本文:
GitHub - christensen143/claude-trofeo-hud: 38 ドルの Thermalright Trofeo Vision LCD の Live Claude 使用状況 HUD — セッション/週ごとの制限、トークンバーン、USB 経由で macOS からストリーミングされた仮想 API コスト。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
クリステンセン143
/
クロード・トロフェオ・ハド
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット claude_trofeo_hud claude_trofeo_hud docs docs スパイク スパイク src/ claud

e_trofeo_hud src/ claude_trofeo_hud テスト テスト .gitignore .gitignore .python-version .python-version PLANNING.md PLANNING.md README.md README.md TASKS.md TASKS.md config.toml config.toml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Thermalright Trofeo Vision 6.86 インチでのクロードのライブ使用状況を表示するデスク HUD
LCD (1280×480、USB-C、約 38 ドル)、macOS から駆動。 r/ClaudeAI からインスピレーションを得た
「$38 クロード LCD テーブル ディスプレイ」の投稿。
表示される内容: Pro/Max セッション + リセット カウントダウン付きの週ごとの制限バー
(Anthropic の使用エンドポイントから)、今日のトークンと仮想 API コスト
(ccusage 経由)、ライブセッション
(プロジェクト、モデル、バーン レート)、時計、および時間ごとのトークン スパークライン。
macOS、Python 3.12+、uv 、ノード (npx ccusage 用)
brew install hidapi (hidapi Python パッケージの背後にある C ライブラリ)
クロード コードがインストールされ、ログインされています (HUD はローカル ログとそのログを読み取ります)
キーチェーンからの OAuth トークン — 読み取り専用、マシンからは何も残りません
api.anthropic.com への使用状況のクエリを除く)
UV同期
uv run python -m claude_trofeo_hudreview # モックレイアウトを out/preview.png にレンダリングします
uv run python -m claude_trofeo_hud run # LCD 上のライブ HUD (Ctrl-C 停止)
uv run python -m claude_trofeo_hud install-agent # launchd 経由でログイン時に開始
最初の実行時に、macOS は「Claude Code-credentials」へのキーチェーン アクセスを要求します。
— デーモンを無人で実行できるように、[常に許可] を選択します。
uninstall-agent は、launchd エージェントを停止して削除します。構成は次の場所に存在します
config.toml (fps、JPEG 品質、夜間薄暗く/オフ時間)。ログは次の場所に移動します
~/ライブラリ/ログ/claude-trofeo-hud/ 。
パネルはモニターではなく、USB HID デバイス (VID:PID 0416:5302) です。
リバースエンジニアリングされたプロトコルを介して JPEG フレームを受け入れます。私たちはデバイスを使用します
Thermalright-trcc-linux のクラス
HidApiTransport (IOHIDM) を使用

anager)、CLI をバイパスします — trcc のデフォルト
libusb を介したトランスポート ルートは、macOS が HID デバイスに対してブロックします。の
アイドル時にはファームウェアがブランクになるため、HUD は継続的にストリーミングします (デフォルトは 2 fps)。
プロトコルに関する完全なメモについては、PLANNING.md を参照してください。
「アクセスが拒否されました (権限が不十分です)」 — 何かが開いています
hidapi ではなく libusb 経由のデバイス。ではなく、CLI を実行していることを確認してください。
trccに直接アクセスしてください。
パネルにはブート ロゴ/ブランクが表示されます - フレームが到着していません。チェックする
~/ライブラリ/ログ/claude-trofeo-hud/hud.log 。取り外し/再接続が処理されます
バックオフで自動的に。
空のコスト/トークン — npx ccusage は最初にターミナルで動作する必要があります。の
launchd エージェントはインストール時にノード パスをその plist にベイクするため、
ノードのアップグレード後に install-agent を再実行します。
制限が古い - キーチェーンへのアクセスが許可されていないか、ログアウトされています
クロード・コード。
38 ドルの Thermalright Trofeo Vision LCD の Live Claude 使用状況 HUD — セッション/週ごとの制限、トークンの書き込み、USB 経由で macOS からストリーミングされる仮想 API コスト。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Live Claude usage HUD for a $38 Thermalright Trofeo Vision LCD — session/weekly limits, token burn, and hypothetical API cost streamed from macOS over USB. - christensen143/claude-trofeo-hud

GitHub - christensen143/claude-trofeo-hud: Live Claude usage HUD for a $38 Thermalright Trofeo Vision LCD — session/weekly limits, token burn, and hypothetical API cost streamed from macOS over USB. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
christensen143
/
claude-trofeo-hud
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits claude_trofeo_hud claude_trofeo_hud docs docs spike spike src/ claude_trofeo_hud src/ claude_trofeo_hud tests tests .gitignore .gitignore .python-version .python-version PLANNING.md PLANNING.md README.md README.md TASKS.md TASKS.md config.toml config.toml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A desk HUD that shows live Claude usage on a Thermalright Trofeo Vision 6.86"
LCD (1280×480, USB-C, ~$38), driven from macOS. Inspired by the r/ClaudeAI
"$38 Claude LCD Table Display" post.
What it shows: Pro/Max session + weekly limit bars with reset countdowns
(from Anthropic's usage endpoint), today's tokens and hypothetical API cost
(via ccusage ), the live session
(project, model, burn rate), a clock, and an hourly token sparkline.
macOS, Python 3.12+, uv , Node (for npx ccusage )
brew install hidapi (C library behind the hidapi Python package)
Claude Code installed and logged in (the HUD reads its local logs and its
OAuth token from the Keychain — read-only, nothing leaves your machine
except the usage query to api.anthropic.com)
uv sync
uv run python -m claude_trofeo_hud preview # render mock layout to out/preview.png
uv run python -m claude_trofeo_hud run # live HUD on the LCD (Ctrl-C stops)
uv run python -m claude_trofeo_hud install-agent # start at login via launchd
On the first run , macOS asks for Keychain access to "Claude Code-credentials"
— choose Always Allow so the daemon can run unattended.
uninstall-agent stops and removes the launchd agent. Config lives in
config.toml (fps, JPEG quality, night dim/off hours). Logs go to
~/Library/Logs/claude-trofeo-hud/ .
The panel is not a monitor — it's a USB HID device (VID:PID 0416:5302 ) that
accepts JPEG frames over a reverse-engineered protocol. We use the device
classes from thermalright-trcc-linux
with its HidApiTransport (IOHIDManager), bypassing its CLI — trcc's default
transport routes through libusb, which macOS blocks for HID devices. The
firmware blanks when idle, so the HUD streams continuously (default 2 fps).
See PLANNING.md for the full protocol notes.
"Access denied (insufficient permissions)" — something is opening the
device via libusb instead of hidapi; make sure you're running our CLI, not
trcc directly.
Panel shows boot logo / blanks — no frames arriving; check
~/Library/Logs/claude-trofeo-hud/hud.log . Unplug/replug is handled
automatically with backoff.
Empty cost/tokens — npx ccusage must work in a terminal first; the
launchd agent bakes the node path into its plist at install time, so
re-run install-agent after Node upgrades.
Limits stale — Keychain access not granted, or you're logged out of
Claude Code.
Live Claude usage HUD for a $38 Thermalright Trofeo Vision LCD — session/weekly limits, token burn, and hypothetical API cost streamed from macOS over USB.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
