---
source: "https://github.com/Vortrix5/sifty"
hn_url: "https://news.ycombinator.com/item?id=48506948"
title: "Show HN: Sifty – Windows cleanup CLI/TUI with local-only AI"
article_title: "GitHub - Vortrix5/sifty: An AI-assisted Windows maintenance CLI + TUI that sifts the junk from the keep. · GitHub"
author: "vortrix5"
captured_at: "2026-06-12T18:46:17Z"
capture_tool: "hn-digest"
hn_id: 48506948
score: 1
comments: 0
posted_at: "2026-06-12T17:29:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sifty – Windows cleanup CLI/TUI with local-only AI

- HN: [48506948](https://news.ycombinator.com/item?id=48506948)
- Source: [github.com](https://github.com/Vortrix5/sifty)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T17:29:02Z

## Translation

タイトル: Show HN: Sifty – ローカル専用 AI による Windows クリーンアップ CLI/TUI
記事のタイトル: GitHub - Vortrix5/sifty: 保管庫からジャンクを選別する、AI 支援の Windows メンテナンス CLI + TUI。 · GitHub
説明: AI 支援による Windows メンテナンス CLI + TUI は、保管庫からジャンクを選別します。 - Vortrix5/シフティ

記事本文:
GitHub - Vortrix5/sifty: 保管庫からジャンクを選別する、AI 支援の Windows メンテナンス CLI + TUI。 · GitHub
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
ヴォルトリックス5
/
五十
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
89 コミット 89 コミット .github/ workflows .github/ workflows docs docs パッケージング パッケージ化 src/ sifty src/ sifty テスト テスト .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ジャンクの掃除、ディスクの分析、重複の検出、アプリとスタートアップ プログラムの管理、
更新を適用し、開発アーティファクトと git ワークツリーを削除し、ファイルを整理します —
スクリプト可能な CLI または全画面端末 UI。オプションの AI アシスタントが実行されます
Ollama 経由でローカルに: マシンからは何も残らず、ただ見るだけです。
ファイルのメタデータ (名前、サイズ、パス)、ファイルの内容は含まれません。
Sifty はファイルを削除し、システム状態を変更するため、悪用されにくいように構築されています。
デフォルトでドライラン - すべての破壊的なコマンドが何を実行するかをプレビューします。
実際の変更には明示的な --apply が必要です。
ごみ箱。完全に削除することはありません。すべての削除は 1 つのプロセスを経ます。
Send2Trash によってサポートされる Trash() 関数。数回元に戻すと、最後のクリーンな状態が復元されます。
保護されたパス — C:\Windows 、 Program Files 、 ProgramData 、ドライブ
root とプロファイル root は、 --apply --yes を指定しても拒否されます。
監査ログ — 適用されたすべての削除は %APPDATA%\sifty\audit.log に記録されます。
AI は何も削除しません。それは助言です。高リスクのツール呼び出し
常にあなたの承認が必要です。
特徴
シフティ
CCクリーナー
Revo アンインストーラー
WinDirStat
ジャンク/キャッシュのクリーニング
✅ 11 以上のカテゴリ
✅
➖
❌
ディスク使用量の分析
✅ 上位 N + ボリューム
➖
❌
✅ ツリーマップ
重複ファインダー
✅ SHA-256、NTFS 対応
✅（有料）
❌
❌
アプリのアンインストール + 残りのスキャン
✅ ウィンゲット + 残り物
✅
✅ + 残り物
❌
アプリのアップデート
✅ WINGET経由

✅（有料）
❌
❌
スタートアップマネージャー
✅ リバーシブル
✅
✅
❌
開発アーティファクトのパージ (node_modules など)
✅
❌
❌
❌
Git ワークツリー / WSL2 VHD クリーンアップ
✅
❌
❌
❌
ローカルAIアシスタント
✅ オラマ
❌
❌
❌
スクリプト可能 (JSON 出力)
✅
❌
❌
❌
ごみ箱 + すべてを元に戻す
✅
➖
➖
該当なし
価格
無料、マサチューセッツ工科大学
フリーミアム
フリーミアム
無料
Sifty は開発者第一で構築されています。すべてがスクリプト化可能で、エンジンは
再利用可能な Python ライブラリ。開発者のマシンを実際にクリーンアップします。
蓄積します (ビルドアーティファクト、孤立したワークツリー、肥大化した WSL2 ディスク)。
pipx install sifty # 推奨 (分離);または: pip install sifty
スクープバケット追加 sifty https://github.com/Vortrix5/scoop - バケット;スクープインストールシフティ
winget install Vortrix5.Sifty # winget-pkgs PR がマージされたら
sifty ドクター # 管理者権限をチェック、winget、Ollama
Python をインストールしないほうがいいですか?からスタンドアロンの sifty.exe をダウンロードします。
最新リリース。
Python - m venv .venv
.\.venv\Scripts\ python.exe - m pip install - e " .[dev] "
.\.venv\Scripts\ python.exe - m pytest - q
使用法
70 回の検査 # すべての読み取り専用スキャンを 1 回実行: ジャンク、アップデート、
# 孤立したファイル、古いファイル、ディスク容量、起動
sifty tui # 全画面インタラクティブ アプリ
#ジャンク
sifty ジャンク スキャン # カテゴリごとに再利用可能なスペースを表示
sifty ジャンク クリーン # プレビューの削除 (予行演習)
sifty Junk clean -- apply # ジャンクをごみ箱に送信 (最初に質問します)
# ディスク
70 個のディスク ボリュームの使用数/空き数/ボリュームごとの合計数
sifty ディスクは C:\Users パスの下の最大のフォルダー/ファイルを分析します
sifty ディスクの重複 D:\ # 重複ファイルと無駄なスペースを見つける
# アプリとアップデート
50 個のアプリのリスト -- サイズ別、インストールされているアプリの数、大きい順
sifty apps orphans # レジストリ内の壊れたアンインストール エントリ
sifty apps アンインストール " App " # winget 経由でアンインストール (プレビューして --apply)
残り 70 個のアプリ

vers " App " # アンインストーラーが残したもの (その後 --apply)
sifty update check # 利用可能なアップデート (winget)
sifty update apply # すべてをアップグレードします (最初に質問します)
# 開発者のクリーンアップ
sifty purge clean #node_modules、dist、__pycache__、target, …
重複を 50 個クリーンアップ D:\Photos # 重複を除去 (コピーを 1 つずつ保持)
大きな C:\Users\you をクリーンアップして、パス内の最大のファイル数を指定します
古いものを 180 日分クリーンアップ -- ダウンロード内の古いアイテム数が 180 # 件
# スタートアップとサービス
sifty スタートアップ リスト スタートアップ プログラム数 (有効/無効)
sifty 起動を無効にする " Spotify " # 可逆 (sifty 起動を有効にする…)
sifty サービス リスト # 厳選されたオプション サービス + 状態
sifty -- 管理サービスは DiagTrack を無効にします # 1 に切り替えます (管理者が必要です)
# 履歴と元に戻す
70 件の履歴 # クリーンアップされた内容 + 再利用された合計スペース
sifty undo # ごみ箱から最新のクリーンを復元します
# ファイルを整理する
プレビュー C:\Users\you\Downloads -- タイプ別に整理
C:\Users\you\Downloads を適用 -- 日付順に整理
sifty Organize undo # 最後に整理したファイルを元に戻す
# 設定
sifty config # すべての設定 + 上書きした設定
sifty config set ai.model " llama3.2:3b "
sifty config edit # エディターで config.toml を開きます
# AI (Ollama の実行が必要)
70のAIステータス
[切り捨てられた]
sifty tui は、7 セクションのサイドバーを備えた全画面アプリを開きます — ホーム、
クリーン、ディスク、アプリ、モニター、レポート、AI:
ホーム — ボリューム ゲージとすべてをスキャンする [検査の実行] ボタン
すぐに;検出結果には、その場で修正するボタンが付属しています (ジャンクをクリーンアップ、
古いダウンロードをクリーンアップし、アップデートを適用します - それぞれ確認の後に行われます)。
クリーン — ジャンク / パージ / 最適化 / スマート クリーンアップを 1 つ屋根の下で実行します (タブ)。
アプリ — インストール済み / アップデート / スタートアップ / サービス、ファジーフィルター付き、
並べ替え、一括アンインストール、アンインストール後の残りの自動スキャン

インストール中。
AI — エージェント チャット: 提案されたツールが実行され、実行/スキップ ボタンが表示されます
会話の中にインラインで表示され、スキャン結果にはフォローアップ アクションが含まれます
ボタン。
Ctrl+P を押してコマンド パレット (任意の画面にジャンプ)、F2 を押して
昇格、一括アクションの行をマークするためのスペース。 「レポート」画面には次のように表示されます
「最後のクリーンを元に戻す」ボタンを使用すると、時間の経過とともにスペースが再利用されます。
構成されたモデルをプルします: ollama pull qwen2.5:3b 。
sifty ai ステータスは「実行中」と報告するはずです。
sifty config コマンドを使用してすべてを設定します。手動で編集する必要はありません
ファイル:
sifty config # すべての設定とオーバーライドを表示します
sifty config set ai.model " llama3.2:3b " # 別のローカル モデルを使用する
sifty config set ai.host " http://localhost:11434 "
sifty config set safety.extra_protected_paths ' ["D:\\重要"] '
sifty 構成セット Junk.include_downloads_installers true
sifty config edit # またはエディタで config.toml を開きます
設定は %APPDATA%\sifty\config.toml にあります。 sifty 構成セットは書き込みのみを行います
キーを変更するため、アップグレード時にデフォルトが引き続き使用されます。
階層化 — 再利用可能なエンジン上のシン フロントエンド、OS 固有の隔離:
ソース/シフティ/
§── cli/ # Typer エントリ ポイント + グループごとに 1 つのシン コマンド モジュール
§── tui/ # テキスト形式の全画面アプリ (ビューは cli ではなくコアを呼び出します)
§── core/ # エンジン: ジャンク、ディスク、アプリ、アップデート、クリーンアップ、
│ │ # スタートアップ、サービス、整理、検査、履歴、…
│ └── safety.py # ★ 保護されたパス、ドライラン ガード、trash()、監査ログ
§── windows/ # OS プリミティブ: winget、レジストリ、UAC、ごみ箱、DISM
§── ai/ # Ollama クライアント、アドバイザー プロンプト、エージェント ツール ループ
└── infra/ # TOML config + ローテーション診断ログ
フロントエンドはコアに依存します。コアは Windows/インフラに依存します。何もない
上向きに輸入します。 GUI は同じエンジン関数 (junk.scan) を呼び出すことができます。

、
disc.find_duplicates 、checkpoint.run_checkup) を書き換えなしで実行します。参照
設計の理論的根拠については docs/ARCHITECTURE.md を参照してください。
.\.venv\Scripts\ python.exe - m pytest - q # 160 以上のテスト、~20 秒
安全ガードレールは、リポジトリ内で最も厳しくテストされたコードです。窓
環境はモック化されているため、スイートも CI 上で実行されます。参照
CONTRIBUTING.md に参加してください。
AI 支援の Windows メンテナンス CLI + TUI は、保管庫からジャンクを選り分けます。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
シフティ v0.6.0
最新の
2026 年 6 月 12 日
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An AI-assisted Windows maintenance CLI + TUI that sifts the junk from the keep. - Vortrix5/sifty

GitHub - Vortrix5/sifty: An AI-assisted Windows maintenance CLI + TUI that sifts the junk from the keep. · GitHub
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
Vortrix5
/
sifty
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
89 Commits 89 Commits .github/ workflows .github/ workflows docs docs packaging packaging src/ sifty src/ sifty tests tests .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
Clean junk, analyze disks, find duplicates, manage apps and startup programs,
apply updates, prune dev artifacts and git worktrees, and organize files — from
a scriptable CLI or a full-screen terminal UI. The optional AI assistant runs
locally via Ollama : nothing leaves your machine, and it only ever sees
file metadata (names, sizes, paths), never file contents.
Sifty deletes files and changes system state, so it is built to be hard to misuse:
Dry-run by default — every destructive command previews what it would do.
Real changes need an explicit --apply .
Recycle Bin, never permanent delete — all removals go through one
trash() function backed by Send2Trash. sifty undo restores the last clean.
Protected paths — C:\Windows , Program Files , ProgramData , the drive
root and your profile root are refused even with --apply --yes .
Audit log — every applied deletion is recorded in %APPDATA%\sifty\audit.log .
The AI never deletes anything — it is advisory; high-risk tool calls
always require your approval.
Feature
Sifty
CCleaner
Revo Uninstaller
WinDirStat
Junk / cache cleaning
✅ 11+ categories
✅
➖
❌
Disk usage analysis
✅ top-N + volumes
➖
❌
✅ treemap
Duplicate finder
✅ SHA-256, NTFS-aware
✅ (paid)
❌
❌
App uninstall + leftover scan
✅ winget + leftovers
✅
✅ + leftovers
❌
App updates
✅ via winget
✅ (paid)
❌
❌
Startup manager
✅ reversible
✅
✅
❌
Dev artifact purge (node_modules, …)
✅
❌
❌
❌
Git worktree / WSL2 VHD cleanup
✅
❌
❌
❌
Local AI assistant
✅ Ollama
❌
❌
❌
Scriptable (JSON output)
✅
❌
❌
❌
Recycle Bin + undo for everything
✅
➖
➖
n/a
Price
Free, MIT
Freemium
Freemium
Free
Sifty is built developer-first : everything is scriptable, the engine is a
reusable Python library, and it cleans the things developer machines actually
accumulate (build artifacts, orphaned worktrees, bloated WSL2 disks).
pipx install sifty # recommended (isolated); or: pip install sifty
scoop bucket add sifty https: // github.com / Vortrix5 / scoop - bucket; scoop install sifty
winget install Vortrix5.Sifty # once the winget-pkgs PR is merged
sifty doctor # check admin rights, winget, Ollama
Prefer no Python install? Download the standalone sifty.exe from the
latest release .
python - m venv .venv
.\.venv\Scripts\ python.exe - m pip install - e " .[dev] "
.\.venv\Scripts\ python.exe - m pytest - q
Usage
sifty checkup # one read-only scan of everything: junk, updates,
# orphans, stale files, disk space, startup
sifty tui # the full-screen interactive app
# Junk
sifty junk scan # show reclaimable space per category
sifty junk clean # preview removal (dry-run)
sifty junk clean -- apply # send junk to the Recycle Bin (asks first)
# Disk
sifty disk volumes # used/free/total per volume
sifty disk analyze C:\Users # biggest folders/files under a path
sifty disk duplicates D:\ # find duplicate files and wasted space
# Apps & updates
sifty apps list -- by - size # installed apps, largest first
sifty apps orphans # broken uninstall entries in the registry
sifty apps uninstall " App " # uninstall via winget (preview, then --apply)
sifty apps leftovers " App " # what the uninstaller left behind (then --apply)
sifty update check # available updates (winget)
sifty update apply # upgrade everything (asks first)
# Developer cleanup
sifty purge clean # node_modules, dist, __pycache__, target, …
sifty cleanup duplicates D:\Photos # de-duplicate (keeps one copy each)
sifty cleanup large C:\Users\you # biggest files under a path
sifty cleanup stale -- days 180 # old items in Downloads
# Startup & services
sifty startup list # startup programs (enabled/disabled)
sifty startup disable " Spotify " # reversible (sifty startup enable …)
sifty services list # curated optional services + state
sifty -- admin services disable DiagTrack # toggle one (needs admin)
# History & undo
sifty history # what was cleaned + total space reclaimed
sifty undo # restore the most recent clean from the Recycle Bin
# Organize files
sifty organize preview C:\Users\you\Downloads -- by type
sifty organize apply C:\Users\you\Downloads -- by date
sifty organize undo # put the last organize's files back
# Configuration
sifty config # all settings + which ones you've overridden
sifty config set ai.model " llama3.2:3b "
sifty config edit # open config.toml in your editor
# AI (requires Ollama running)
sifty ai status
[truncated]
sifty tui opens a full-screen app with a seven-section sidebar — Home,
Clean, Disk, Apps, Monitor, Reports, AI:
Home — volume gauges and a Run checkup button that scans everything
at once; findings come with buttons that fix them right there (clean junk,
clean stale downloads, apply updates — each behind a confirm).
Clean — Junk / Purge / Optimize / Smart cleanup under one roof (tabs).
Apps — Installed / Updates / Startup / Services, with fuzzy filter,
sorting, bulk uninstall, and an automatic leftover scan after uninstalling.
AI — an agentic chat: tool runs it proposes show Run/Skip buttons
inline in the conversation , and scan results carry follow-up action
buttons.
Press Ctrl+P for the command palette (jump to any screen), F2 to
elevate, Space to mark rows for bulk actions. The Reports screen shows
space reclaimed over time with an Undo last clean button.
Pull the configured model: ollama pull qwen2.5:3b .
sifty ai status should report "running".
Configure everything with the sifty config command — no need to hand-edit
files:
sifty config # show all settings + your overrides
sifty config set ai.model " llama3.2:3b " # use a different local model
sifty config set ai.host " http://localhost:11434 "
sifty config set safety.extra_protected_paths ' ["D:\\Important"] '
sifty config set junk.include_downloads_installers true
sifty config edit # or open config.toml in your editor
Settings live in %APPDATA%\sifty\config.toml ; sifty config set only writes
the keys you change, so defaults keep flowing through on upgrades.
Layered — thin frontends over a reusable engine, OS specifics quarantined:
src/sifty/
├── cli/ # Typer entry point + one thin command module per group
├── tui/ # Textual full-screen app (views call core, not cli)
├── core/ # the engine: junk, disk, apps, updates, cleanup,
│ │ # startup, services, organize, checkup, history, …
│ └── safety.py # ★ protected paths, dry-run guard, trash(), audit log
├── windows/ # OS primitives: winget, registry, UAC, Recycle Bin, DISM
├── ai/ # Ollama client, advisor prompts, agentic tool loop
└── infra/ # TOML config + rotating diagnostics log
Frontends depend on core ; core depends on windows / infra ; nothing
imports upward. A GUI could call the same engine functions ( junk.scan ,
disk.find_duplicates , checkup.run_checkup ) without a rewrite. See
docs/ARCHITECTURE.md for the design rationale.
.\.venv\Scripts\ python.exe - m pytest - q # 160+ tests, ~20 s
The safety guardrails are the most heavily tested code in the repo; the Windows
environment is mocked so the suite also runs on CI. See
CONTRIBUTING.md to get involved.
An AI-assisted Windows maintenance CLI + TUI that sifts the junk from the keep.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Sifty v0.6.0
Latest
Jun 12, 2026
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
