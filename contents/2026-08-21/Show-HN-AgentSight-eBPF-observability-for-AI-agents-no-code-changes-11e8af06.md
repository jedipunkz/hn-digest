---
source: "https://github.com/alibaba/anolisa/blob/main/docs/user-guide/en/agent-observability/agentsight.md"
hn_url: "https://news.ycombinator.com/item?id=49389493"
title: "Show HN: AgentSight – eBPF observability for AI agents, no code changes"
article_title: "anolisa/docs/user-guide/en/agent-observability/agentsight.md at main · alibaba/anolisa · GitHub"
image: "https://repository-images.githubusercontent.com/1195873793/e3f09d3a-3b99-4615-8948-44bccc219b45"
author: "forrestly"
captured_at: "2026-08-21T15:23:41Z"
capture_tool: "hn-digest"
hn_id: 49389493
score: 1
comments: 0
posted_at: "2026-08-21T15:21:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentSight – eBPF observability for AI agents, no code changes

- HN: [49389493](https://news.ycombinator.com/item?id=49389493)
- Source: [github.com](https://github.com/alibaba/anolisa/blob/main/docs/user-guide/en/agent-observability/agentsight.md)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T15:21:10Z

## Translation

タイトル: HN の表示: AgentSight – AI エージェントの eBPF 可観測性、コード変更なし
記事のタイトル: anolisa/docs/user-guide/en/agent-observability/agentsight.md (メイン) · alibaba/anolisa · GitHub
説明: ANOLISA (Agentic Nexus オペレーティング層およびインターフェイス システム アーキテクチャ) |ランタイム、セキュリティ、可観測性、トークンレス応答圧縮を備えたエージェントティック OS により、トークンの使用量とコストが削減されます。 - メインの anolisa/docs/user-guide/en/agent-observability/agentsight.md · alibaba/anolisa

記事本文:
anolisa/docs/user-guide/en/agent-observability/agentsight.md (メイン) · alibaba/anolisa · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
アリババ
/
アノリサ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 304 行 (206 loc) · 10.2 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
rawファイルをコピー rawファイルをダウンロード アウトライン編集してr

aw アクション AgentSight
AgentSight は、eBPF に基づいたゼロインストルメンテーション AI エージェント可観測性ツールです。エージェント コードを変更せずに、LLM API 呼び出し、トークン消費、およびプロセスの動作をカーネル レベルでキャプチャします。
AgentSight は、Linux 上で実行される AI エージェントにフルスタックの可観測性を提供します。
要件
最小値
OS
Linux
カーネル
>= 5.8 (BTF サポートが必要)
特権
root または CAP_BPF (eBPF プローブの場合)
ANOLISA生パッケージ
Linux x86_64、システムモード
macOS : macOS では、AgentSight は 2 つのコマンド、trace (ローカル JSONL セッション ファイルをスキャンするトラジェクトリ コレクター、eBPF なし) とserve (ダッシュボード ビューアー) を提供します。他のすべての eBPF 依存コマンドは Linux のみです。
ANOLISA CLI を使用して公開コンポーネントをインストールします。
# 推奨 (システム モードが必要 — eBPF には root が必要)
sudo anolisa エージェントサイトをインストールする
# 代替案 (Alinux、YUM リポジトリ設定が必要)
sudo yum インストール Agentsight
# ソースビルド (開発者のみ)
cd src/agentsight && make build-all
ソース ビルドには make build-all を使用します。これにより、ダッシュボード フロントエンド、メイン バイナリ、agentsight-enforcer が順番にビルドされます。 make build のみを実行するとエンフォーサがスキップされ、serv ではログ記録が継続され、AgentSight の強制は利用できなくなります。
通常の展開には systemd ユニットを使用します。 eBPF トレースを実行し、
ダッシュボードを統合し、必要な順序でエンフォーサの依存関係を開始します。
sudo systemctl Enable --now Agentsight.service
sudo systemctlステータスagentsight.service
サービスがアクティブになったら、http://localhost:7396 を開きます。有効にする
メインユニットは再起動後も AgentSight を利用可能な状態に保ちます。
バンドルされている systemd ランチャーは、ダッシュボードを 0.0.0.0 にバインドします。ポートを制限する
ホストをセキュリティにさらす前に、ファイアウォールまたはセキュリティ グループを使用して 7396 を実行してください。
信頼できないネットワーク。
サービスはプライベート umask を使用して root として実行され、データを以下に保存します。
/

var/log/sysak/.agentsight 。 CLI クエリとダッシュボード アクセスには sudo を使用する
このサービスが所有するデータを読み取るコマンド。
フォアグラウンドのトラブルシューティングでは、最初に systemd ユニットを停止します。
2 番目のトレーサーと競合します。次に、2 つの端末を使用して、両方のコマンドを次のように実行します。
根。両方を連続して入力すると、2 番目のコマンドには到達しません。
Agentsight トレースがフォアグラウンドに留まるため、
sudo systemctl stop Agentsight.service
# 第1ターミナル
sudo エージェントサイト トレース
# ターミナル 2: ダッシュボードを開始します
sudo エージェントサイト サーブ
# ブラウザで http://localhost:7396 を開きます
# ダッシュボードの URL とトークンを出力します。デスクトップ ユーザーとして URL を開きます
sudo Agentsight ダッシュボード --no-open
ローカルホストへのアクセスは認証不要です。リモート アクセスにはトークンが必要です。「 ダッシュボードのアクセスと認証 」を参照してください。
Agentsight トレース — eBPF トレースの開始
AI エージェント アクティビティのカーネル レベルのキャプチャを開始します。
sudo エージェントサイト トレース
root権限が必要です。 SSL/TLS トラフィック、プロセス イベント、およびファイル操作をキャプチャします。
フォアグラウンド トレーサーを開始する前に、sudo systemctl stop Agentsight.service を実行します。
Agentsight Serve — API とダッシュボードの開始
# デフォルト: 127.0.0.1:7396 にバインド
sudo エージェントサイト サーブ
# すべてのインターフェイスにバインド (リモート アクセス)
sudo エージェントサイト サーブ --ホスト 0.0.0.0 --ポート 7396
トレースを実行するのと同じユーザーとしてserveを実行すると、両方のコマンドで問題が解決されます。
同じデータディレクトリ。 0.0.0.0 にバインドすると、すべてのデバイスでダッシュボードが公開されます。
インターフェイス。そのフォームを使用する前にネットワーク アクセスを制限してください。
ダッシュボード トークン認証はデフォルトで有効になっています。
ローカルホスト アクセス (ループバック) は認証をバイパスします。 http://127.0.0.1:7396 を開くだけです。
リモート アクセスにはトークンが必要です。ブラウザ URL に ?token=<TOKEN> を追加するか、Authorization: Bearer <TOKEN> HTTP ヘッダーを設定します。
トークンは最初のサービス起動時に自動生成されます (64 16 進数文字)。

データベースの隣にある .dashboard_token ファイルに保存されます (デフォルトは /var/log/sysak/.agentsight/.dashboard_token )。再起動しても再利用されます。
sudo Agentsight Dashboard --no-open を実行してサービス所有のアクセス URL とトークンを出力し、デスクトップ ユーザーとしてその URL を開きます。
認証を無効にするには (信頼できる内部ネットワークでのみ推奨)、構成ファイルで次のように設定します。
{
"サーバー" : { "認証" : { "有効" : false } }
}
/etc/agentsight/config.json を編集した後、 sudo systemctl reload Agentsight.service を実行して変更を適用します。再起動は必要ありません。
GET /api/docs は完全な API ルート インベントリ (メソッド、パス、説明) を返すため、スクリプトと統合はエンドポイントを検出できます。不明な /api/ パスへのリクエストも、404 応答でそのパスを指します。
カール http://127.0.0.1:7396/api/docs
Agentsight ダッシュボード — ダッシュボードのアクセス情報を表示
ダッシュボード URL と認証トークンを表示し、ブラウザーを開こうとします。 ECS インスタンスでは、セキュリティ グループ構成ガイドも出力されます。
# root 所有のブラウザを開かずに URL とトークンを表示する
sudo Agentsight ダッシュボード --no-open
Agentsight の概要 — 統合された概要
セッションとトークンの使用状況、重大度ごとにグループ化された中断イベント、最近の時間枠のトークンレス節約量をロールアップします。これは、全体的な状態の状況を把握するための 1 つのコマンドです。
# 過去 24 時間 (デフォルト)
エージェントサイトの概要
# 過去 7 日間、JSON 出力
エージェントサイトの概要 --last 168 --json
データ ソースは独立して劣化します。欠落しているデータベースは、レポートの残りの部分に影響を与えることなく、ゼロを提供します。
Agentsight トークン — クエリ トークンの使用法
#今日の使い方
sudo エージェントサイトトークン
# 週ごとの比較
sudo Agentsight トークン --期間週 --比較
# JSON出力
sudo Agentsight トークン --json
Agentsight Audit — 監査イベントのクエリ
# 最近の出来事
エージェントサイト監査
# PID と t によるフィルター

そうそう
Agentsight 監査 --pid 12345 --type llm
# 要約統計量
Agentsight 監査 --概要
Agentsight Discover — エージェントのスキャン
# 実行中の AI エージェントを検出する
エージェントサイトの発見
# 既知のエージェント タイプをリストします。
Agentsight Discover --list-known
エージェントサイト中断 — セッション中断イベント
AI エージェントのセッション中断イベントをクエリおよび管理します。
# 中断イベントをリストします (デフォルト: 過去 24 時間)
Agentsight 中断リスト [--last < HOURS > ] [--type < TYPE > ] [--severity < LEVEL > ]
# タイプ別の統計
Agentsight の中断統計
# 重大度別にカウントする
Agentsight の中断回数
# ID で単一のイベントを取得する
Agentsight の中断 < ID > を取得
# セッション/会話のすべての中断イベントを一覧表示します
Agentsight 中断セッション < SESSION_ID >
Agentsight の中断会話 < CONVERSATION_ID >
# 解決済みとしてマークする
Agentsight の中断解決 < ID >
構成
設定ファイル: /etc/agentsight/config.json ( --config で上書き)。
重要 : ユーザー設定ファイルは、組み込みのデフォルト ルールを置き換えます (拡張ではありません)。必要なすべてのエージェント ルールが構成に含まれていることを確認してください。
特徴
JSONパス
デフォルト
説明
トークンの統計
features.token_stats
本当の
コアトークンアカウンティング
SQLiteストレージ
features.sqlite_storage.enabled
本当の
ローカル永続性
中断検出
機能.中断検出.有効
本当の
エラー/クラッシュ検出
監査
機能.監査
本当の
LLM コール監査
セッションマッピング
features.session_mapping.enabled
本当の
応答Id→セッションId
実行時間の制限
構成
デフォルト
説明
イベントチャネル容量
10,000
プローブイベント限定チャネル容量
pending_genai_max_count
1,000
session_id を待機している最大イベント数
max_connection_body_mb
8
単一の HTTP 接続本文のバッファ制限
リングバッファ_mb
32
eBPF リング バッファ サイズ (2 の累乗である必要があります)
エージェントフレームワークの統合
アジャン

tSight は、Copilot Shell に組み込みの会話スキルを提供します。ユーザーは、自然言語を介してトークンの使用状況と監査ログをクエリできます。
「今日私はどれくらいのトークンを使いましたか？」
「今日のLLM通話記録を見せて」
トークンの節約 (トークンレス統合)
AgentSight はトークンレスコンポーネントと統合して、ダッシュボードにトークン節約データを表示します。追加の構成は必要ありません。両方がインストールされている場合は、節約データが自動的に表示されます。
デフォルトの最大データベース サイズ: 200 MB。到達すると、自動クリーンアップがトリガーされます。
環境変数を使用してカスタマイズします。
エクスポート AGENTSIGHT_GENAI_DB_MAX_SIZE_MB=500
履歴をクリアする
rm -rf /var/log/sysak/.agentsight
# その後、AgentSight を再起動します
よくある質問
Q: OpenClaw のトークン データが表示されないのはなぜですか?
A: AgentSight は、openclaw-gateway デーモンを監視します。クライアントとゲートウェイの接続を確認します。 「ペアリングが必要です」エラーが表示された場合は、 openclaw devices Approve を実行します。
Q: トークンの節約ページに 0 が表示されるのはなぜですか?
A: 考えられる原因: (1) AK/SK 認証モードはまだサポートされていません。 (2) セッション ID 形式が非標準の UUID である。
Q: 累積節約額が 1 回のコールの差額を超えるのはなぜですか?
A: エージェントには、コンテキスト内の履歴メッセージが含まれます。節約はターンをまたいで蓄積されるため、累積節約はターンごとの差を超えます。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

ANOLISA (Agentic Nexus Operating Layer & Interface System Architecture) | Agentic OS with runtime, security, observability, and Tokenless response compression for lower token usage and cost. - anolisa/docs/user-guide/en/agent-observability/agentsight.md at main · alibaba/anolisa

anolisa/docs/user-guide/en/agent-observability/agentsight.md at main · alibaba/anolisa · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
alibaba
/
anolisa
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 304 lines (206 loc) · 10.2 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions AgentSight
AgentSight is a zero-instrumentation AI Agent observability tool based on eBPF. It captures LLM API calls, Token consumption, and process behavior at the kernel level without modifying Agent code.
AgentSight provides full-stack observability for AI Agents running on Linux:
Requirement
Minimum
OS
Linux
Kernel
>= 5.8 (BTF support required)
Privileges
root or CAP_BPF (for eBPF probes)
ANOLISA raw package
Linux x86_64, system mode
macOS : On macOS, AgentSight provides two commands — trace (trajectory collector that scans local JSONL session files, no eBPF) and serve (Dashboard viewer). All other eBPF-dependent commands are Linux-only.
Install the published component with the ANOLISA CLI:
# Recommended (system mode required — eBPF needs root)
sudo anolisa install agentsight
# Alternative (Alinux, requires YUM repo configuration)
sudo yum install agentsight
# Source build (developers only)
cd src/agentsight && make build-all
Use make build-all for source builds: it builds the Dashboard frontend, the main binary, and agentsight-enforcer in sequence. Running only make build skips the enforcer, and serve will keep logging AgentSight enforcement unavailable .
Use the systemd unit for a normal deployment. It runs eBPF tracing and the
Dashboard together and starts the enforcer dependency in the required order:
sudo systemctl enable --now agentsight.service
sudo systemctl status agentsight.service
Open http://localhost:7396 after the service becomes active. Enabling the
main unit also keeps AgentSight available after a reboot.
The bundled systemd launcher binds the Dashboard to 0.0.0.0 . Restrict port
7396 with a firewall or security group before exposing the host to an
untrusted network.
The service runs as root with a private umask and stores data under
/var/log/sysak/.agentsight . Use sudo for CLI queries and Dashboard access
commands that read this service-owned data.
For foreground troubleshooting, stop the systemd unit first so it does not
compete with a second tracer. Then use two terminals and run both commands as
root. The second command is not reached if both are entered sequentially
because agentsight trace stays in the foreground:
sudo systemctl stop agentsight.service
# Terminal 1
sudo agentsight trace
# Terminal 2: Start Dashboard
sudo agentsight serve
# Open http://localhost:7396 in browser
# Print the Dashboard URL and token; open the URL as your desktop user
sudo agentsight dashboard --no-open
Localhost access is authentication-free; remote access requires a token, see Dashboard Access & Authentication .
agentsight trace — Start eBPF Tracing
Starts kernel-level capture of AI Agent activity.
sudo agentsight trace
Requires root privileges. Captures SSL/TLS traffic, process events, and file operations.
Run sudo systemctl stop agentsight.service before starting a foreground tracer.
agentsight serve — Start API & Dashboard
# Default: bind to 127.0.0.1:7396
sudo agentsight serve
# Bind to all interfaces (remote access)
sudo agentsight serve --host 0.0.0.0 --port 7396
Run serve as the same user that runs trace so both commands resolve the
same data directory. Binding to 0.0.0.0 exposes the Dashboard on every
interface; restrict network access before using that form.
Dashboard token authentication is enabled by default:
Localhost access (loopback) bypasses authentication — just open http://127.0.0.1:7396 .
Remote access requires a token: append ?token=<TOKEN> to the browser URL, or set the Authorization: Bearer <TOKEN> HTTP header.
The token is auto-generated on the first serve startup (64 hex characters) and persisted to the .dashboard_token file next to the database (default /var/log/sysak/.agentsight/.dashboard_token ); it is reused across restarts.
Run sudo agentsight dashboard --no-open to print the service-owned access URL and token, then open the URL as your desktop user.
To disable authentication (only recommended on trusted internal networks), set in the config file:
{
"server" : { "auth" : { "enabled" : false } }
}
After editing /etc/agentsight/config.json , run sudo systemctl reload agentsight.service to apply the change — no restart needed.
GET /api/docs returns the full API route inventory (method, path, description) so scripts and integrations can discover endpoints; requests to unknown /api/ paths also point to it in the 404 response.
curl http://127.0.0.1:7396/api/docs
agentsight dashboard — Show Dashboard Access Info
Displays the Dashboard URL and auth token, then tries to open a browser. On ECS instances it also prints a security-group configuration guide.
# Show URL and token without opening a root-owned browser
sudo agentsight dashboard --no-open
agentsight summary — Unified Overview
Rolls up sessions and Token usage, interruption events grouped by severity, and Tokenless savings for a recent time window — one command for the overall health picture.
# Last 24 hours (default)
agentsight summary
# Last 7 days, JSON output
agentsight summary --last 168 --json
Data sources degrade independently: a missing database contributes zeros without affecting the rest of the report.
agentsight token — Query Token Usage
# Today's usage
sudo agentsight token
# Weekly comparison
sudo agentsight token --period week --compare
# JSON output
sudo agentsight token --json
agentsight audit — Query Audit Events
# Recent events
agentsight audit
# Filter by PID and type
agentsight audit --pid 12345 --type llm
# Summary statistics
agentsight audit --summary
agentsight discover — Scan for Agents
# Discover running AI Agents
agentsight discover
# List known Agent types
agentsight discover --list-known
agentsight interruption — Session Interruption Events
Query and manage AI Agent session interruption events.
# List interruption events (default: last 24h)
agentsight interruption list [--last < HOURS > ] [--type < TYPE > ] [--severity < LEVEL > ]
# Statistics by type
agentsight interruption stats
# Count by severity
agentsight interruption count
# Get a single event by ID
agentsight interruption get < ID >
# List all interruption events of a session / conversation
agentsight interruption session < SESSION_ID >
agentsight interruption conversation < CONVERSATION_ID >
# Mark as resolved
agentsight interruption resolve < ID >
Configuration
Configuration file: /etc/agentsight/config.json (override with --config ).
Important : User config files replace (not extend) the built-in default rules. Ensure your config includes all Agent rules you need.
Feature
JSON Path
Default
Description
Token stats
features.token_stats
true
Core Token accounting
SQLite storage
features.sqlite_storage.enabled
true
Local persistence
Interruption detection
features.interruption_detection.enabled
true
Error/crash detection
Audit
features.audit
true
LLM call audit
Session mapping
features.session_mapping.enabled
true
responseId→sessionId
Runtime Limits
Config
Default
Description
event_channel_capacity
10,000
Probe event bounded channel capacity
pending_genai_max_count
1,000
Max events awaiting session_id
max_connection_body_mb
8
Single HTTP connection body buffer limit
ring_buffer_mb
32
eBPF Ring Buffer size (must be power of 2)
Agent Framework Integration
AgentSight provides a built-in conversational skill for Copilot Shell. Users can query Token usage and audit logs via natural language:
"How much Token did I use today?"
"Show me today's LLM call records"
Token Savings (Tokenless Integration)
AgentSight integrates with the Tokenless component to display Token savings data in the Dashboard. No additional configuration needed — if both are installed, savings data appears automatically.
Default maximum database size: 200 MB. When reached, automatic cleanup triggers.
Customize via environment variable:
export AGENTSIGHT_GENAI_DB_MAX_SIZE_MB=500
Clear History
rm -rf /var/log/sysak/.agentsight
# Then restart AgentSight
FAQ
Q: Why can't I see Token data for OpenClaw?
A: AgentSight monitors the openclaw-gateway daemon. Check client-gateway connectivity. If you see "pairing required" errors, run openclaw devices approve .
Q: Why does the Token savings page show 0?
A: Possible causes: (1) The AK/SK authentication mode is not yet supported; (2) Session ID format is non-standard UUID.
Q: Why do cumulative savings exceed the single-call difference?
A: Agents include historical messages in context. Savings accumulate across turns, so cumulative savings exceed per-turn differences.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
