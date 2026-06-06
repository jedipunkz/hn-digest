---
source: "https://github.com/fabioconcina/claumon"
hn_url: "https://news.ycombinator.com/item?id=48423227"
title: "Claumon – Forecasting Claude Code usage limits with a Gamma process"
article_title: "GitHub - fabioconcina/claumon: Claude Code dashboard — minimal, fast, single binary · GitHub"
author: "fabioconcina"
captured_at: "2026-06-06T12:34:52Z"
capture_tool: "hn-digest"
hn_id: 48423227
score: 1
comments: 1
posted_at: "2026-06-06T09:52:57Z"
tags:
  - hacker-news
  - translated
---

# Claumon – Forecasting Claude Code usage limits with a Gamma process

- HN: [48423227](https://news.ycombinator.com/item?id=48423227)
- Source: [github.com](https://github.com/fabioconcina/claumon)
- Score: 1
- Comments: 1
- Posted: 2026-06-06T09:52:57Z

## Translation

タイトル: Claumon – ガンマ プロセスによるクロード コードの使用制限の予測
記事のタイトル: GitHub - fabioconcina/claumon: Claude Code ダッシュボード — 最小限、高速、単一バイナリ · GitHub
説明: Claude Code ダッシュボード — 最小限、高速、単一バイナリ - fabioconcina/claumon

記事本文:
GitHub - fabioconcina/claumon: クロード コード ダッシュボード — 最小限、高速、単一バイナリ · GitHub
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
ファビオコンチーナ
/
クラモン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店

s タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
85 コミット 85 コミット .github/ workflows .github/ workflows アセット アセット ドキュメント ドキュメント 内部 内部スクリプト スクリプト Web Web .gitignore .gitignore .goreleaser.yml .goreleaser.yml CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md bench.go bench.go benchtools_stub.go benchtools_stub.go Diagnostics.go Diagnostics.go go.mod go.mod go.sum go.sum main.go main.go priority.json priority.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code ダッシュボード - 実行して、トークンがどこに行くかを確認します。
単一のバイナリ、設定なし、ブラウザ タブ 1 つ。 macOS、Linux、Windows 上で動作します。
Claude Code のダッシュボードはたくさんあります。これは特別なことをしようとしているわけではありません。これは単なる私の見解であり、最小限、高速、セットアップが簡単であるという、私が重視している点に合わせて最適化されています。
単一のバイナリ、依存関係、ビルド手順、構成は不要です。それを実行し、ブラウザーのタブを開きます。
使用量予測、セッションごとのトークンの内訳、コスト見積もり、履歴傾向、会話履歴を含むレート制限ゲージ、および関係グラフ、ヘルス スコア、および失効アラートを備えたメモリ ブラウザーが提供されます。すべては SSE 経由でリアルタイムに更新されます。毎日の集計は SQLite に保存されるため、現在のセッションだけでなく、数週間にわたって使用状況を追跡できます。
ライブ ダッシュボード: レート制限ゲージ、トークンとコストの内訳、および 14 日間の傾向。
セッション予測: 80% 信頼区間でのリセット時の予測使用量。
醸造インストール ファビオコンシナ/クラウモン/クラウモン
クラモン
最新リリースからビルド済みバイナリをダウンロードします。
# macOS (アップルシリコン)
カール -Lo claumon https://github.com/fabioconcina/claumon/releases/latest/download/claumon-darwin-arm64
chmod +x クラウモン
./クラウモン
# macOS (インテル)
カール -ロー・クラウモン https://

github.com/fabioconcina/claumon/releases/latest/download/claumon-darwin-amd64
chmod +x クラウモン
./クラウモン
# Linux (x86_64)
カール -Lo claumon https://github.com/fabioconcina/claumon/releases/latest/download/claumon-linux-amd64
chmod +x クラウモン
./クラウモン
# Windows (x86_64)
# リリース ページから claumon-windows-amd64.exe をダウンロードします。
# Defender によってブロックされている場合: .exe を右クリック -> プロパティ -> ブロックを解除
# または次を実行します: Unblock-File claumon-windows-amd64.exe
ソースからビルドする:
github.com/fabioconcina/claumon@latest をインストールしてください
クラモン
ブラウザで http://localhost:3131 を開きます。
claumon は、 ~/.claude/.credentials.json (claude /login によって作成) から資格情報を読み取るか、VS Code 拡張機能によって使用される OS 資格情報ストア (macOS キーチェーン、Windows 資格情報マネージャー) にフォールバックします。認証情報が不足している場合でも、セッション追跡は機能します。 API 使用状況ゲージのみが利用できません。
ログイン時に claumon が自動的に実行されるため、ダッシュボードは常に利用可能です。まず、バイナリを永続的な場所に移動してから、サービスを登録します。
mkdir -p ~ /.local/bin
mv クラウモン ~ /.local/bin/
# まだ PATH に追加していない場合は追加します。export PATH="$HOME/.local/bin:$PATH"
Claumon サービスのインストール
~/Library/LaunchAgents/com.claumon.dashboard.plist に LaunchAgent を作成します。ログは ~/.claumon/claumon.log に保存されます。
mv クラウモン ~ /.local/bin/
Claumon サービスのインストール
~/.config/systemd/user/claumon.service に systemd ユーザー ユニットを作成します。ログは、journald (journalctl --user -u claumon) に保存されます。デフォルトでは、ログアウトすると systemd ユーザー サービスが停止します。ログアウト後も claumon を実行し続けるには、残留を有効にします:loginctl enable-linger $USER 。
# exeを永続的な場所に移動します
新しい項目 - ItemType ディレクトリ - 「 $ env: LOCALAPPDATA \Programs\claumon 」を強制します
Move-Item claumon-windows-amd64.exe " $ env: LOCALAPPDATA \Programs\claumon\claum

.exe上」
# PATH に追加 (現在のユーザー、セッション間で保持)
$p = [環境]::GetEnvironmentVariable(' パス ' , ' ユーザー ' )
[ 環境 ]::SetEnvironmentVariable( ' パス ' , " $p ; $ env: LOCALAPPDATA \Programs\claumon " , ' ユーザー ' )
# 端末を再起動してから、次のようにします。
Claumon サービスのインストール
Windows スタートアップ フォルダーに claumon を追加します。ログイン時に自動的に (非表示に) 起動します。
Claumon サービスのステータス # 実行中かどうかを確認する
claumon サービスの再起動 # 設定変更後の再起動
Claumon サービスのアンインストール # 停止して削除します
どのプラットフォームでも root や管理者は必要ありません。
クラウモンのアップデート
GitHub リリースで新しいバージョンを確認し、プラットフォームに適したバイナリをダウンロードして、現在のバイナリを置き換えます。バックグラウンド サービスがインストールされている場合、アップデート後に自動的に再起動されます。
バックグラウンド サービス (インストールされている場合) を削除し、バイナリを削除し、必要に応じてローカル データ (SQLite データベース、ログ、構成、価格設定キャッシュ) を削除します。
Claumon サービスのアンインストール
rm ~ /.local/bin/claumon
rm -rf ~ /.クラウモン
Windows (PowerShell):
Claumon サービスのアンインストール
Remove-Item (Get-Command claumon).Source
Remove-Item - 再帰 " $ env: USERPROFILE \.claumon "
追跡する内容
レート制限 (Claude API からライブ)
セッション使用量 - リセットカウントダウン付きの 5 時間のスライディング ウィンドウ使用量
毎週の使用量 - リセットカウントダウン付きの 7 日間のウィンドウ使用量
モデルごとの割り当て - Opus と Sonnet の個別の週次制限 (該当する場合)
追加使用量クレジット - 毎月の制限と使用量 (有効な場合)
ゲージは色分けされています: 緑 (<50%)、黄 (50 ～ 80%)、赤 (>80%)。
リセット時の予測使用率 - 80% の信頼できる間隔と、各レート制限ウィンドウのしきい値までの ETA なので、そこに到達する前にどこに着地するかを確認できます。
経験的ベイズ モデル。単調な (減少しない) 使用量にわたって、自分の過去のウィンドウから毎日再調整されます。

使用率がウィンドウ内でのみ増加することを考慮して間隔が設定されるように処理します。完全な仕様は、internal/forecast/MODEL.pdf (LaTeX ソース: MODEL.tex ) にあります。
Claumon ベンチによるアウトオブサンプルのベンチマーク - リーブワンアウトおよびテンポラルホールドアウトプロトコル、カバレッジとバイアスの内訳を含む適切なスコアリング (CRPS/ピンボール)。
トークンの使用法 (セッション ファイルから)
セッションごとの入力/出力/キャッシュ読み取り/キャッシュ書き込みトークン
現在のモデルの価格設定を使用した同等の API コストの推定
14 日間の傾向グラフを含む日次集計
アクティビティ ヒートマップ - 時間ごとのトークン アクティビティの 24 時間の内訳
アクティブなセッション テーブル - プロジェクト、モデル、トークン、コスト、メッセージ、最後のアクティビティ
今日/最近の切り替え - 今日のセッションと常時の最新 50 セッションを切り替えます
実行中のプロセスの検出 - どのセッションがアクティブに実行されているかを緑色の「アクティブ」バッジで表示します
セッション詳細ビュー - メッセージごとのトークン数とツールの使用状況を含む完全なメッセージ タイムライン
自動検出 - ~/.claude/projects/ で新しいセッションと更新されたセッションを監視します
ライブ プロセス テーブル - 実行中のすべてのクロード コード プロセスを PID、チャット タイトル、プロジェクト、タイプ、稼働時間とともに表示します。
[停止] ボタン - SIGINT を送信して、実行中のプロセスを正常に停止します (会話はディスク上に保存されます)。
自動更新 - SSE および定期的なポーリングによる更新
すべてのメモリ ファイル - CLAUDE.md、ルール、自動メモリ インデックス、プロジェクトごとのメモリ
コンテンツ、パス、プロジェクト名、フロントマター全体を検索します
カテゴリによるフィルター - CLAUDE.md、ルール、インデックス、メモリ
古さのインジケーター - 緑 (今日)、灰色 (<7 日)、黄色 (7 ～ 30 日)、赤 (>30 日)
失効アラート - 壊れた MEMORY.md リンク、孤立したファイル、インデックスの不一致
VS Code の統合 - クリックすると、vscode:// リンクを介してエディターでファイルが開きます
ファイルごとのグレーディング - 各メモリ ファイルには、鮮度、構造、特異性、および安全性に基づいて文字グレード (A ～ F) が付けられます。

接続性
改善提案 - スコアの低いファイルに対する実用的なヒント (フロントマターの追加、MEMORY.md からのリンクなど)
インタラクティブな視覚化 - ノードはメモリ ファイル、エッジは関係を示します
プロジェクト フィルター - 特定のプロジェクトに焦点を当てる
クリック可能な凡例 - ノード タイプのオン/オフを切り替えます
クリックして移動 - ノードをクリックしてメモリ ブラウザ内のファイルにジャンプします
キー
アクション
1 2 3
ダッシュボード、メモリー、グラフタブに切り替えます
/
メモリーサーチへジャンプ
Esc
セッション詳細パネルを閉じる
ライブアップデート
ダッシュボードはサーバー送信イベント (SSE) を使用します。ポーリングや手動更新はありません。変更は即座に反映されます。
使用状況ゲージは Claude API から 2 分ごとに更新されます (構成可能)
ディスク上の .jsonl セッション ファイルが変更されると、セッション テーブルが更新されます
メモリブラウザは、視覚的なパルスで変更されたファイルを強調表示します
ディスク上のメモリ ファイルが変更されると、グラフと古さが再レンダリングされます
上部バーのステータス ドットは接続状態を示します (緑 = 接続、赤 = 切断)。
オプション。 ~/.claumon/config.json を作成します。
{
「ポート」: 3131 、
"ポーリング間隔秒" : 120 、
"credentials_path" : " ~/.claude/.credentials.json " ,
"claude_dir" : " ~/.claude " ,
"db_path" : " ~/.claumon/usage.db " ,
「スタックしきい値分」: 10
}
すべてのフィールドはオプションです。デフォルトは上に示されています。 claumon は設定ファイルなしで動作します。
claumon は、使用状況のスナップショットと毎日の集計を ~/.claumon/usage.db の SQLite データベースに保存します。初回起動時に既存のセッション ファイルをすべてスキャンすることにより、履歴データが自動的にバックフィルされます。
データベースは、書き込み中の同時読み取りに WAL モードを使用します。メンテナンスは必要ありません。
claumon は 3 つのデータ ソースを組み合わせます。
Claude API ( /api/oauth/usage ) - レート制限の使用状況について定期的にポーリングされます。 claude /login からの OAuth 認証情報が必要です。
セッション ファイル ( ~/.claude/projects/*/*.jsonl ) - fsnot 経由で監視

リアルタイムのトークンカウント用にify。各 JSONL ファイルは、メッセージ レベルのトークンの使用法について解析されます。
メモリ ファイル ( ~/.claude/projects/*/memory/*.md 、 CLAUDE.md 、 rules) - フロントマターが検出および解析され、HTML としてレンダリングされ、相互参照と古さについて分析されます。
//go:embed を介してすべてが単一のバイナリに埋め込まれます。外部ファイル、Node.js、フロントエンドのビルド ステップはありません。
グラフTD
API["クロード OAuth API"] --> エンジン
JSONL["セッション JSONL ファイル"] --> エンジン
MEM["メモリ ファイル"] --> エンジン
PRICE["価格設定データ"] --> エンジン
ENGINE["ポーラー · ウォッチャー · パーサー"] --> DB["SQLite"]
エンジン --> キャッシュ["メモリ内キャッシュ"]
エンジン --> SSE["SSE ブローカー"]
DB --> HTTP["HTTP サーバー"]
キャッシュ --> HTTP
JSONL -->|"新しく解析されました"| HTTP
SSE --> HTTP
HTTP --> SPA["組み込み SPA"]
読み込み中
ライセンス
Claude Code ダッシュボード — 最小限、高速、単一バイナリ
claumon.fabioconcina.com
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
30
v0.13.1
最新の
2026 年 6 月 5 日
+ 29 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Code dashboard — minimal, fast, single binary - fabioconcina/claumon

GitHub - fabioconcina/claumon: Claude Code dashboard — minimal, fast, single binary · GitHub
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
fabioconcina
/
claumon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
85 Commits 85 Commits .github/ workflows .github/ workflows assets assets docs docs internal internal scripts scripts web web .gitignore .gitignore .goreleaser.yml .goreleaser.yml CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md bench.go bench.go benchtools_stub.go benchtools_stub.go diagnostics.go diagnostics.go go.mod go.mod go.sum go.sum main.go main.go pricing.json pricing.json View all files Repository files navigation
Claude Code dashboard - run it, see where your tokens go.
Single binary, zero config, one browser tab. Runs on macOS, Linux, and Windows.
There are plenty of Claude Code dashboards out there. This one isn't trying to be special - it's just my take on it, optimized for the things I care about: minimal, fast, and easy to set up.
Single binary, no dependencies, no build step, no config required. Run it and open a browser tab.
It gives you rate limit gauges with usage forecasts, per-session token breakdowns, cost estimates, historical trends, conversation history, and a memory browser with relationship graph, health scores, and staleness alerts. Everything updates in real time via SSE. Daily aggregates are stored in SQLite so you can track usage over weeks, not just the current session.
Live dashboard: rate-limit gauges, token and cost breakdowns, and 14-day trends.
Session forecast: projected usage at reset with an 80% confidence interval.
brew install fabioconcina/claumon/claumon
claumon
Download a prebuilt binary from the latest release :
# macOS (Apple Silicon)
curl -Lo claumon https://github.com/fabioconcina/claumon/releases/latest/download/claumon-darwin-arm64
chmod +x claumon
./claumon
# macOS (Intel)
curl -Lo claumon https://github.com/fabioconcina/claumon/releases/latest/download/claumon-darwin-amd64
chmod +x claumon
./claumon
# Linux (x86_64)
curl -Lo claumon https://github.com/fabioconcina/claumon/releases/latest/download/claumon-linux-amd64
chmod +x claumon
./claumon
# Windows (x86_64)
# Download claumon-windows-amd64.exe from the releases page
# If blocked by Defender: right-click the .exe -> Properties -> Unblock
# Or run: Unblock-File claumon-windows-amd64.exe
Build from source:
go install github.com/fabioconcina/claumon@latest
claumon
Open http://localhost:3131 in your browser.
claumon reads credentials from ~/.claude/.credentials.json (created by claude /login ), or falls back to the OS credential store (macOS Keychain, Windows Credential Manager) used by the VS Code extension. If credentials are missing, session tracking still works. Only the API usage gauges are unavailable.
Run claumon automatically on login so the dashboard is always available. First, move the binary to a permanent location, then register the service.
mkdir -p ~ /.local/bin
mv claumon ~ /.local/bin/
# Add to PATH if not already: export PATH="$HOME/.local/bin:$PATH"
claumon service install
Creates a LaunchAgent at ~/Library/LaunchAgents/com.claumon.dashboard.plist . Logs go to ~/.claumon/claumon.log .
mv claumon ~ /.local/bin/
claumon service install
Creates a systemd user unit at ~/.config/systemd/user/claumon.service . Logs go to journald ( journalctl --user -u claumon ). By default, systemd user services stop when you log out. To keep claumon running after logout, enable lingering: loginctl enable-linger $USER .
# Move the exe to a permanent location
New-Item - ItemType Directory - Force " $ env: LOCALAPPDATA \Programs\claumon "
Move-Item claumon-windows-amd64.exe " $ env: LOCALAPPDATA \Programs\claumon\claumon.exe "
# Add to PATH (current user, persists across sessions)
$p = [ Environment ]::GetEnvironmentVariable( ' Path ' , ' User ' )
[ Environment ]::SetEnvironmentVariable( ' Path ' , " $p ; $ env: LOCALAPPDATA \Programs\claumon " , ' User ' )
# Restart your terminal, then:
claumon service install
Adds claumon to the Windows Startup folder - it launches automatically (hidden) on login.
claumon service status # check if running
claumon service restart # restart after config changes
claumon service uninstall # stop and remove
No root or admin required on any platform.
claumon update
Checks GitHub releases for a newer version, downloads the right binary for your platform, and replaces the current one. If a background service is installed, it restarts automatically after the update.
Remove the background service (if installed), delete the binary, and optionally remove local data (SQLite database, logs, config, pricing cache).
claumon service uninstall
rm ~ /.local/bin/claumon
rm -rf ~ /.claumon
Windows (PowerShell):
claumon service uninstall
Remove-Item ( Get-Command claumon).Source
Remove-Item - Recurse " $ env: USERPROFILE \.claumon "
What it tracks
Rate limits (live from Claude API)
Session usage - 5-hour sliding window utilization with reset countdown
Weekly usage - 7-day window utilization with reset countdown
Per-model quotas - separate Opus and Sonnet weekly limits (when applicable)
Extra usage credits - monthly limit and spend (if enabled)
Gauges are color-coded: green (<50%), yellow (50–80%), red (>80%).
Projected utilization at reset - an 80% credible interval and ETA to threshold for each rate-limit window, so you can see where you'll land before you get there.
Empirical-Bayes model, refit daily from your own past windows, over a monotone (non-decreasing) usage process so the interval respects that utilization only grows within a window. Full spec in internal/forecast/MODEL.pdf (LaTeX source: MODEL.tex ).
Benchmarked out-of-sample with claumon bench - leave-one-out and temporal-holdout protocols, proper scoring (CRPS/pinball) with coverage and bias breakdowns.
Token usage (from session files)
Input / output / cache read / cache write tokens per session
Estimated equivalent API cost using current model pricing
Daily aggregates with 14-day trend charts
Activity heatmap - 24-hour breakdown of token activity by hour
Active sessions table - project, model, tokens, cost, messages, last activity
Today / Recent toggle - switch between today's sessions and the 50 most recent across all time
Running process detection - shows which sessions are actively running with green "active" badge
Session detail view - full message timeline with per-message token counts and tool usage
Automatic discovery - watches ~/.claude/projects/ for new and updated sessions
Live process table - shows all running Claude Code processes with PID, chat title, project, type, and uptime
Stop button - send SIGINT to gracefully stop any running process (conversation is preserved on disk)
Auto-refresh - updates via SSE and periodic polling
All memory files - CLAUDE.md, rules, auto-memory indexes, per-project memories
Search across content, path, project name, and frontmatter
Filter by category - CLAUDE.md, Rules, Index, Memory
Staleness indicators - green (today), gray (<7d), yellow (7–30d), red (>30d)
Staleness alerts - broken MEMORY.md links, orphaned files, index mismatches
VS Code integration - click to open files in your editor via vscode:// links
Per-file grading - each memory file gets a letter grade (A–F) based on freshness, structure, specificity, and connectedness
Improvement suggestions - actionable tips for low-scoring files (add frontmatter, link from MEMORY.md, etc.)
Interactive visualization - nodes are memory files, edges show relationships
Project filters - focus on specific projects
Clickable legend - toggle node types on/off
Click to navigate - click a node to jump to its file in the memory browser
Key
Action
1 2 3
Switch to Dashboard, Memories, Graph tab
/
Jump to memory search
Esc
Close session detail panel
Live updates
The dashboard uses Server-Sent Events (SSE) - no polling, no manual refresh. Changes appear instantly:
Usage gauges update every 2 minutes (configurable) from the Claude API
Session table updates when any .jsonl session file changes on disk
Memory browser highlights changed files with a visual pulse
Graph and staleness re-render when memory files change on disk
A status dot in the top bar shows connection state (green = connected, red = disconnected).
Optional. Create ~/.claumon/config.json :
{
"port" : 3131 ,
"poll_interval_seconds" : 120 ,
"credentials_path" : " ~/.claude/.credentials.json " ,
"claude_dir" : " ~/.claude " ,
"db_path" : " ~/.claumon/usage.db " ,
"stuck_threshold_minutes" : 10
}
All fields are optional. Defaults are shown above; claumon works without a config file.
claumon stores usage snapshots and daily aggregates in a SQLite database at ~/.claumon/usage.db . Historical data is backfilled automatically on first startup by scanning all existing session files.
The database uses WAL mode for concurrent reads during writes. No maintenance required.
claumon combines three data sources:
Claude API ( /api/oauth/usage ) - polled periodically for rate limit utilization. Requires OAuth credentials from claude /login .
Session files ( ~/.claude/projects/*/*.jsonl ) - watched via fsnotify for real-time token counting. Each JSONL file is parsed for message-level token usage.
Memory files ( ~/.claude/projects/*/memory/*.md , CLAUDE.md , rules) - discovered and parsed for frontmatter, rendered as HTML, analyzed for cross-references and staleness.
Everything is embedded in a single binary via //go:embed - no external files, no Node.js, no build step for the frontend.
graph TD
API["Claude OAuth API"] --> ENGINE
JSONL["Session JSONL Files"] --> ENGINE
MEM["Memory Files"] --> ENGINE
PRICE["Pricing Data"] --> ENGINE
ENGINE["Pollers · Watchers · Parsers"] --> DB["SQLite"]
ENGINE --> CACHE["In-Memory Cache"]
ENGINE --> SSE["SSE Broker"]
DB --> HTTP["HTTP Server"]
CACHE --> HTTP
JSONL -->|"parsed fresh"| HTTP
SSE --> HTTP
HTTP --> SPA["Embedded SPA"]
Loading
License
Claude Code dashboard — minimal, fast, single binary
claumon.fabioconcina.com
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
30
v0.13.1
Latest
Jun 5, 2026
+ 29 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
