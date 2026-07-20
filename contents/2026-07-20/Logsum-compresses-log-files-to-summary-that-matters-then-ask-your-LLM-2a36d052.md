---
source: "https://github.com/flouthoc/logsum"
hn_url: "https://news.ycombinator.com/item?id=48979370"
title: "Logsum compresses log files to summary that matters, then ask your LLM"
article_title: "GitHub - flouthoc/logsum: Compress large log files to summary that actually matter — then ask your LLM. · GitHub"
author: "fl_git"
captured_at: "2026-07-20T15:02:11Z"
capture_tool: "hn-digest"
hn_id: 48979370
score: 1
comments: 0
posted_at: "2026-07-20T14:29:03Z"
tags:
  - hacker-news
  - translated
---

# Logsum compresses log files to summary that matters, then ask your LLM

- HN: [48979370](https://news.ycombinator.com/item?id=48979370)
- Source: [github.com](https://github.com/flouthoc/logsum)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T14:29:03Z

## Translation

タイトル: Logsum はログ ファイルを重要な要約に圧縮します。LLM に問い合わせてください。
記事のタイトル: GitHub - flouthoc/logsum: 大きなログ ファイルを実際に重要な要約に圧縮します。その後、LLM に問い合わせてください。 · GitHub
説明: 大きなログ ファイルを実際に重要な要約に圧縮します。その後、LLM に問い合わせてください。 - flouthoc/logsum

記事本文:
GitHub - flouthoc/logsum: 大きなログ ファイルを実際に重要な要約に圧縮してから、LLM に問い合わせてください。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
フルホク
/
ログサム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コ

デ
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット アセット アセット src src テスト テスト .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
生のアプリケーション ログを、LLM 消費用に設計されたトークンに最適化された Markdown 形式に圧縮する、軽量のローカル ファースト CLI ツールです。
生のログには、冗長なタイムスタンプ、定型スタック トレース、および何千もの重複行が含まれています。これらを LLM に直接渡すと、コンテキスト ウィンドウが無駄になり、API コストが増加し、診断の精度が低下します。
logsum は、ノイズを取り除き、繰り返し発生するイベントを分解し、残りの異常を構造化された分析マトリックスにパッケージ化して、LLM が即座に低コストで解析できる決定論的エンジンです。
ログサム [オプション] [ファイル]
引数:
[FILE] ログファイルへのパス (省略した場合は stdin を読み取ります)
オプション:
-o, --output <FILE> 出力をファイルに書き込みます (デフォルト: stdout)
-f, --format <FMT> 強制フォーマット: "json" または "text" (デフォルト: 自動検出)
-v、--verbose 処理統計を標準エラー出力に表示します。
-h, --help ヘルプを印刷する
例
# ファイルから
ログサム /var/log/app.log
# stdin から (別のコマンドからのパイプ)
kubectl は私のポッドをログに記録します。ログサム
# JSON 形式を強制して出力を保存する
logsum -f json -o summary.md app.log
# 詳細モード (標準エラー出力の統計)
logsum -v パイプライン.log > summary.md
仕組み
入力 (標準入力/ファイル)
→ 形式の検出 (JSON とプレーン テキスト、自動検出)
→ 解析（タイムスタンプ、レベル、コンポーネント、メッセージの抽出）
→ 重複排除 (フィンガープリント + 繰り返しのエントリの折りたたみ)
→ 圧縮 (エラー/警告 + コンテキストを保持、ノイズをドロップ)
→ レンダリング (マークダウン分析マトリックス)
処理パイプライン
フォーマット検出 — 最初の空でない 10 行をサンプリングします。 70% を超える場合は、有効な JSON オブジェクトとして解析され、JSON m が使用されます。

頌歌;それ以外の場合はプレーンテキストモード。
解析 — 各行から構造化フィールドを抽出します。
JSON モード: タイムスタンプ、レベル/重大度、メッセージ/メッセージ、コンポーネント/ロガーフィールドを読み取ります。
プレーン テキスト モード: 一般的な形式 (YYYY-MM-DD HH:MM:SS LEVEL [コンポーネント] メッセージ、syslog など) の正規表現ベースの抽出
キーワード推論: 非構造化ログ (CI 出力) の場合、 FAILED 、 Traceback 、 Exception 、 PASSED などのキーワードからレベルを推論します。
重複排除 — 可変部分 (UUID、16 進数 ID、タイムスタンプ、IP) を削除してメッセージを正規化し、フィンガープリントを作成します。連続する同一のエントリを折りたたみ、3 回以上見つかった行をグローバルに重複排除します。
圧縮 — ログ構造に基づく 2 つの戦略:
構造化ログ: すべての警告/エラー エントリ + 3 行のコンテキスト ウィンドウ、および時間範囲の最初/最後のエントリを保持します。
非構造化ログ (CI/パイプライン出力): JSON 構成のバルクを削除しながら、セクション ヘッダー、ステータス行、エラーとコンテキスト、および意味のある短い行を保持します。
レンダリング — 分析マトリックスを出力します (下記を参照)。
出力形式は、トークンの使用を最小限に抑えながら、LLM にとって最大限に役立つように設計されています。各セクションの内容は次のとおりです。
### ログプロセッサ分析マトリックス
* トークン効率を高めるために、生のログは [ X ] % で圧縮されます。 *
#### 1. メタデータの概要
#### 2. ログ ストリーム スキーマ
#### 3. 異常のハイライト
セクションの内訳
LLM がスコープを理解するための簡単な統計:
期間: ログの時間範囲 (開始→終了、デルタ付き)
ボリューム : 元の行数 → 圧縮されたエントリ数
コンポーネント : ログ内で検出されたモジュール/サービス
トークン効率のために短縮キーを使用した、圧縮された JSONL 形式の圧縮ログ エントリ:
セクション 3 — 異常のハイライト
エラーと警告のみを重点的に表示し、繰り返し回数も表示します。これにより、LLM は問題の原因に直接ジャンプできるようになります。

ストリーム全体をスキャンしています。
出力例 (構造化されたアプリケーション ログ)
入力: ヘルスチェックスパムと接続プール障害を含む 26 行のアプリケーションログ。
### ログプロセッサ分析マトリックス
* トークン効率を高めるために、生のログは 61.5% 圧縮されました。 *
#### 1. メタデータの概要
* ** 期間: ** 10:00:00 ～ 10:00:10 (10 秒)
* ** ボリューム: ** 26 行が 10 エントリに減りました。
* ** コンポーネント: ** main、db、http
#### 2. ログ ストリーム スキーマ
``` jsonl
{ "t" : " 0s " 、 "lvl" : " INFO " 、 "comp" : " main " 、 "msg" : " アプリケーション起動中 " }
{ "t" : " 0s " 、 "lvl" : " INFO " 、 "comp" : " db " 、 "msg" : " 接続プールが 10 個の接続で初期化されました " }
{ "t" : " 1s " 、 "lvl" : " INFO " 、 "comp" : " http " 、 "msg" : " サーバーは 0.0.0.0:8080 をリッスンしています " }
{ "t" : " 5s " 、 "lvl" : " 情報 " 、 "comp" : " http " 、 "msg" : " GET /health 200 1ms " 、 "n" : 10 }
{ "t" : " 10s " 、 "lvl" : " WARN " 、 "comp" : " db " 、 "msg" : " 接続プールが不足しています (残り 2 つ) " }
{ "t" : " 10s " 、 "lvl" : " WARN " 、 "comp" : " db " 、 "msg" : " 接続プールが不足しています (残り 1 つ) " }
{ "t" : " 10s " 、 "lvl" : " エラー " 、 "comp" : " db " 、 "msg" : " 接続プールが枯渇しました - 5000 ミリ秒後のクエリ タイムアウト " }
{ "t" : " 10s " 、 "lvl" : " エラー " 、 "comp" : " http " 、 "msg" : " POST /api/users 503 5001ms " }
{ "t" : " 10s " 、 "lvl" : " 情報 " 、 "comp" : " db " 、 "msg" : " 接続を再利用しようとしています " }
{ "t" : " 10s " 、 "lvl" : " 情報 " 、 "comp" : " db " 、 "msg" : " 3 つのアイドル接続を回収しました " }
3. 異常のハイライト
[警告] [db] 接続プールが不足しています (残り 2 つ)
[警告] [db] 接続プールが不足しています (残り 1 つ)
[エラー] [db] 接続プールが枯渇しました - 5000 ミリ秒後のクエリ タイムアウト
[エラー] [http] POST /api/users 503 5001ms
主な所見:
- 同一のヘルスチェックライン 10 行 → 単一

`"n":10` を含むエントリ
- すべてのエラーと警告は周囲のコンテキストとともに保存されます
- 開始からの相対タイムスタンプ (絶対タイムスタンプではなくトークンを保存)
サポートされているログ形式
フォーマット
検出
例
構造化されたプレーンテキスト
自動
2024-01-15 10:23:45 エラー [db] 接続に失敗しました
JSON行
自動
{"タイムスタンプ":"..."、"レベル":"エラー"、"メッセージ":"..."}
シスログ
自動
1 月 15 日 10:23:45 ホスト sshd[1234]: パスワードに失敗しました
CI/パイプライン出力
自動
セクションヘッダーのある非構造化テキスト
開発
# テストを実行する
貨物試験
# ビルドリリースバイナリ
カーゴビルド --release
# 詳細な出力で実行する
貨物の実行 -- -v your.log
ライセンス
大きなログ ファイルを実際に重要な要約に圧縮してから、LLM に問い合わせてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Compress large log files to summary that actually matter — then ask your LLM. - flouthoc/logsum

GitHub - flouthoc/logsum: Compress large log files to summary that actually matter — then ask your LLM. · GitHub
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
flouthoc
/
logsum
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits assets assets src src tests tests .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml README.md README.md View all files Repository files navigation
A lightweight, local-first CLI tool that compresses raw application logs into a token-optimized Markdown format designed for LLM consumption.
Raw logs are filled with redundant timestamps, boilerplate stack traces, and thousands of duplicate lines. Passing them directly to LLMs wastes context window, increases API costs, and dilutes diagnostic accuracy.
logsum is a deterministic engine that strips noise, collapses repeating events, and packages the remaining anomalies into a structured Analysis Matrix that LLMs can parse instantly and cheaply.
logsum [OPTIONS] [FILE]
Arguments:
[FILE] Path to log file (reads stdin if omitted)
Options:
-o, --output <FILE> Write output to file (default: stdout)
-f, --format <FMT> Force format: "json" or "text" (default: auto-detect)
-v, --verbose Show processing stats on stderr
-h, --help Print help
Examples
# From a file
logsum /var/log/app.log
# From stdin (pipe from another command)
kubectl logs my-pod | logsum
# Force JSON format and save output
logsum -f json -o summary.md app.log
# Verbose mode (stats on stderr)
logsum -v pipeline.log > summary.md
How It Works
Input (stdin/file)
→ Format Detection (JSON vs plain text, auto-detected)
→ Parsing (extract timestamp, level, component, message)
→ Deduplication (fingerprint + collapse repeated entries)
→ Compression (keep errors/warnings + context, drop noise)
→ Render (Markdown Analysis Matrix)
Processing Pipeline
Format Detection — Samples the first 10 non-empty lines. If >70% parse as valid JSON objects, uses JSON mode; otherwise plain text mode.
Parsing — Extracts structured fields from each line:
JSON mode : reads timestamp , level / severity , msg / message , component / logger fields
Plain text mode : regex-based extraction for common formats ( YYYY-MM-DD HH:MM:SS LEVEL [component] message , syslog, etc.)
Keyword inference : for unstructured logs (CI output), infers levels from keywords like FAILED , Traceback , Exception , PASSED
Deduplication — Normalizes messages by stripping variable parts (UUIDs, hex IDs, timestamps, IPs) to create fingerprints. Collapses consecutive identical entries and globally deduplicates lines seen more than 3 times.
Compression — Two strategies based on log structure:
Structured logs : keeps all Warn/Error entries + 3-line context window, plus first/last entries for time range
Unstructured logs (CI/pipeline output): keeps section headers, status lines, errors + context, and meaningful short lines while dropping JSON config bulk
Rendering — Outputs the Analysis Matrix (see below).
The output format is designed to be maximally useful to LLMs while minimizing token usage. Here's what each section contains:
### LOG PROCESSOR ANALYSIS MATRIX
* Raw logs compressed by [ X ] % for token efficiency. *
#### 1. Metadata Summary
#### 2. Log Stream Schema
#### 3. Anomaly Highlights
Section Breakdown
Quick stats for the LLM to understand scope:
Duration : time range of the log (start → end, with delta)
Volume : original line count → compressed entry count
Components : detected modules/services in the log
The compressed log entries in minified JSONL format with abbreviated keys for token efficiency:
Section 3 — Anomaly Highlights
A focused view of just the errors and warnings, with repeat counts. This lets LLMs jump straight to what went wrong without scanning the full stream.
Example Output (Structured Application Log)
Input: 26-line application log with health check spam and a connection pool failure.
### LOG PROCESSOR ANALYSIS MATRIX
* Raw logs compressed by 61.5% for token efficiency. *
#### 1. Metadata Summary
* ** Duration: ** 10:00:00 to 10:00:10 (10s)
* ** Volume: ** 26 lines reduced to 10 entries.
* ** Components: ** main, db, http
#### 2. Log Stream Schema
``` jsonl
{ "t" : " 0s " , "lvl" : " INFO " , "comp" : " main " , "msg" : " Application starting up " }
{ "t" : " 0s " , "lvl" : " INFO " , "comp" : " db " , "msg" : " Connection pool initialized with 10 connections " }
{ "t" : " 1s " , "lvl" : " INFO " , "comp" : " http " , "msg" : " Server listening on 0.0.0.0:8080 " }
{ "t" : " 5s " , "lvl" : " INFO " , "comp" : " http " , "msg" : " GET /health 200 1ms " , "n" : 10 }
{ "t" : " 10s " , "lvl" : " WARN " , "comp" : " db " , "msg" : " Connection pool running low (2 remaining) " }
{ "t" : " 10s " , "lvl" : " WARN " , "comp" : " db " , "msg" : " Connection pool running low (1 remaining) " }
{ "t" : " 10s " , "lvl" : " ERROR " , "comp" : " db " , "msg" : " Connection pool exhausted - query timeout after 5000ms " }
{ "t" : " 10s " , "lvl" : " ERROR " , "comp" : " http " , "msg" : " POST /api/users 503 5001ms " }
{ "t" : " 10s " , "lvl" : " INFO " , "comp" : " db " , "msg" : " Attempting to reclaim connections " }
{ "t" : " 10s " , "lvl" : " INFO " , "comp" : " db " , "msg" : " Reclaimed 3 idle connections " }
3. Anomaly Highlights
[WARN] [db] Connection pool running low (2 remaining)
[WARN] [db] Connection pool running low (1 remaining)
[ERROR] [db] Connection pool exhausted - query timeout after 5000ms
[ERROR] [http] POST /api/users 503 5001ms
Key observations:
- 10 identical health check lines → single entry with `"n":10`
- All errors and warnings preserved with surrounding context
- Timestamps relative to start (saves tokens vs absolute timestamps)
Supported Log Formats
Format
Detection
Example
Structured plain text
Auto
2024-01-15 10:23:45 ERROR [db] Connection failed
JSON lines
Auto
{"timestamp":"...","level":"ERROR","msg":"..."}
Syslog
Auto
Jan 15 10:23:45 host sshd[1234]: Failed password
CI/Pipeline output
Auto
Unstructured text with section headers
Development
# Run tests
cargo test
# Build release binary
cargo build --release
# Run with verbose output
cargo run -- -v your.log
License
Compress large log files to summary that actually matter — then ask your LLM.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
