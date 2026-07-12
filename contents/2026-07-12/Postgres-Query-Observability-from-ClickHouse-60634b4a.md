---
source: "https://github.com/ClickHouse/pg_stat_ch"
hn_url: "https://news.ycombinator.com/item?id=48882210"
title: "Postgres Query Observability from ClickHouse"
article_title: "GitHub - ClickHouse/pg_stat_ch · GitHub"
author: "saisrirampur"
captured_at: "2026-07-12T16:46:48Z"
capture_tool: "hn-digest"
hn_id: 48882210
score: 1
comments: 0
posted_at: "2026-07-12T16:08:32Z"
tags:
  - hacker-news
  - translated
---

# Postgres Query Observability from ClickHouse

- HN: [48882210](https://news.ycombinator.com/item?id=48882210)
- Source: [github.com](https://github.com/ClickHouse/pg_stat_ch)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T16:08:32Z

## Translation

タイトル: ClickHouse による Postgres クエリの可観測性
記事タイトル: GitHub - ClickHouse/pg_stat_ch · GitHub
説明: GitHub でアカウントを作成して、ClickHouse/pg_stat_ch の開発に貢献します。

記事本文:
GitHub - ClickHouse/pg_stat_ch · GitHub
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
読み込み中にエラーが発生しました。このページをリロードしてください。
クリックハウス
/
pg_stat_ch
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
114 コミット 114 コミット .claude/ スキル .claude/ スキル .github .github cmake cmake docker docker docs docs イメージ イメージには include スキーマ/ マイグレーション スキーマ/ マイグレーション スクリプト スクリプト スペック スペック sql sql src src t t テスト/回帰テスト/回帰 サードパーティ サードパーティ トリプレット トリプレット.clang-format .clang-format .clang-tidy .clang-tidy .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules CLAUDE.md CLAUDE.md CMakeLists.txt CMakeLists.txt CMakePresets.json CMakePresets.json INSTALL.md INSTALL.md LICENSE.md LICENSE.md META.json META.json Makefile Makefile README.md README.md misse.toml misse.toml pg_stat_ch.control pg_stat_ch.control vcpkg-configuration.json vcpkg-configuration.json vcpkg.json vcpkg.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
pg_stat_ch: ClickHouse への PostgreSQL クエリ テレメトリ エクスポーター
クエリごとの実行テレメトリをキャプチャし、リアルタイムで ClickHouse にエクスポートする PostgreSQL 拡張機能。 PostgreSQL で統計を集計する pg_stat_statements とは異なり、pg_stat_ch は生のイベントを ClickHouse にエクスポートし、ClickHouse の強力な分析エンジンを介して集計が行われます。
スキーマがプリロードされたローカル PostgreSQL + ClickHouse を実行します。
./scripts/quickstart.sh アップ
./scripts/quickstart.sh チェック
停止:
./scripts/quickstart.sh ダウン
エンドポイントとスタックの詳細については、docker/quickstart/README.md を参照してください。
推奨される PostgreSQL 設定
pg_stat_ch は、PostgreSQL で実行されるすべてのクエリの詳細なテレメトリをキャプチャし、単一のデータ パイプラインを介して ClickHouse にエクスポートします。
PostgreSQL フック (フォアグラウンド) → 共有メモリ キュー → バックグラウンド ワーカー → ClickHouse
主要な設計原則:
クエリ p でのネットワーク I/O がゼロ

ath - イベントは共有メモリのキューに入れられ、同期的に送信されません
集計ではなく生のイベント - すべての集計 (p50/p95/p99、上位クエリ、エラー) は ClickHouse で発生します
境界付きメモリ - オーバーフロー カウンタを備えた固定サイズのリング バッファ。ドロップされたイベントはクエリをブロックしません
最小限のオーバーヘッド - キャプチャされたステートメントあたり約 5μs の p99 オーバーヘッド
PostgreSQL フックは完全なインストルメンテーションでクエリの開始/終了をキャプチャします
共有メモリ リング バッファはイベントを保存します (MPSC: マルチプロデューサー、シングルコンシューマー)
バックグラウンド ワーカーがバッチをデキューし、ClickHouse にエクスポートします
ClickHouse は生のイベントを events_raw テーブルに保存します。ビュー/MV は集計を提供します
完全なクエリ テレメトリ: タイミング、行数、バッファ使用量、WAL 使用量、CPU 時間
すべてのステートメント タイプ: DML (SELECT/INSERT/UPDATE/DELETE/MERGE)、DDL、ユーティリティ ステートメント
エラー キャプチャ: Emit_log_hook による SQLSTATE コードとエラー レベル
JIT インストルメンテーション (PG15+): 関数数、生成/インライン化/最適化/発行時間
パラレル ワーカーの統計 (PG18+): 計画されたワーカーと開始されたワーカー
クライアントコンテキスト: アプリケーション名、クライアントIPアドレス
クエリテキスト: 切り詰めてキャプチャ (最大 2KB)
グレースフル デグラデーション : キュー オーバーフローにより、カウンターを使用してイベントがドロップされます。 ClickHouse が利用できない場合でも PostgreSQL はブロックされません
PostgreSQL 16、17、および 18 は完全にサポートされています。機能マトリックスとバージョン固有のフィールドについては、docs/version-compatibility.md を参照してください。
CMake 3.16+、C++17 コンパイラ (GCC 9+、Clang 10+)
PostgreSQL 16 以降の開発ヘッダー
ミス (推奨) または手動での PostgreSQL インストール
clickhouse-c はサブモジュールとして販売されており、静的にリンクされています。
git サブモジュール update --init --recursive
# misse の使用（推奨）
mise run build # デバッグビルド (特定のバージョンには build:16/17/18 を使用)
miss run build:release # リリースビルド
miss run install # 拡張機能をインストールする
# または手動で
cmake -B build -G Ninj

-DPG_CONFIG=/path/to/pg_config
cmake --build ビルド && cmake --install ビルド
インストール
shared_preload_libraries = ' pg_stat_ch '
track_io_timing = on # キャプチャされたイベントの I/O タイミング列を有効にします
# ClickHouse 接続 (セットアップに合わせて変更)
pg_stat_ch.clickhouse_host = ' localhost '
pg_stat_ch.clickhouse_port = 9000
pg_stat_ch.clickhouse_database = ' pg_stat_ch '
# TLS (運用環境に推奨)
pg_stat_ch.clickhouse_use_tls = オン
pg_stat_ch.clickhouse_skip_tls_verify = オフ
2.PostgreSQLを再起動します
# サービスマネージャー (systemd、brew services、Docker など) を使用して PostgreSQL を再起動します。
3. ClickHouse でスキーマを設定する
クイックスタート パス: ./scripts/quickstart.sh を実行すると、スキーマのセットアップが自動的に処理されます。
手動/既存の ClickHouse パス:
クリックハウスクライアント < docker/init/00-schema.sql
events_raw とマテリアライズド ビューを作成します。完全なスキーマの詳細については、docs/clickhouse.md を参照してください。
拡張機能 pg_stat_ch を作成します。
5. 確認する
SELECT pg_stat_ch_version();
SELECT * FROM pg_stat_ch_stats();
構成
パラメータ
種類
デフォルト
リロード
説明
pg_stat_ch.enabled
ブール
に
ため息
テレメトリ収集の有効化/無効化
pg_stat_ch.clickhouse_host
文字列
ローカルホスト
再起動
ClickHouseサーバーのホスト名
pg_stat_ch.clickhouse_port
整数
9000
再起動
ClickHouse ネイティブ プロトコル ポート
pg_stat_ch.clickhouse_user
文字列
デフォルト
再起動
クリックハウスのユーザー名
pg_stat_ch.clickhouse_password
文字列
「」
再起動
クリックハウスのパスワード
pg_stat_ch.clickhouse_database
文字列
pg_stat_ch
再起動
ClickHouse データベース名
pg_stat_ch.queue_capacity
整数
65536
再起動
リング バッファ サイズ (2 の累乗である必要があります)
pg_stat_ch.string_area_size
整数
64MB
再起動
可変長クエリ文字列とエラー文字列の DSA サイズ
pg_stat_ch.flush_interval_ms
整数
1000
ため息
エクスポートバッチ間隔 (ミリ秒)
pg_stat_ch.batch_max
整数
10000
ため息

ClickHouse 挿入あたりの最大イベント数
pg_stat_ch.clickhouse_use_tls
ブール
オフ
再起動
ClickHouse 接続の TLS を有効にする
pg_stat_ch.clickhouse_skip_tls_verify
ブール
オフ
再起動
TLS 証明書の検証をスキップする (安全ではない)
pg_stat_ch.log_min_elevel
列挙型
警告
スーパーユーザー
キャプチャする最小エラー レベル (debug5..panic)
ClickHouse のエラー レベルとその数値の完全なリストについては、「エラー レベルの値」を参照してください。
機能
説明
pg_stat_ch_version()
拡張機能のバージョン文字列を返します
pg_stat_ch_stats()
キューとエクスポーターの統計 (列の詳細)
pg_stat_ch_reset()
すべてのキューカウンターをゼロにリセットします
pg_stat_ch_flush()
キューに入れられたイベントの ClickHouse への即時フラッシュをトリガーします
クエリの例
生の実行ごとのイベントにより、パーセンタイル、時系列、アプリごとのドリルダウンが可能になります。
-- 特定のアプリケーションの最も遅いクエリ (パーセンタイル付き)
SELECT クエリ ID、
count () AS 呼び出し、
分位数( 0 . 95 )(duration_us) / 1000 AS p95_ms、
分位数( 0 . 99 )(duration_us) / 1000 AS p99_ms
pg_stat_ch から。イベント生
WHERE アプリ = ' myapp '
かつ ts_start > now() - 間隔 1 時間
GROUP BY クエリ ID
p99_ms DESC で注文
リミット10;
マテリアライズド ビューの定義とその他のクエリの例については、docs/clickhouse.md を参照してください。
miss run test:all # すべてのテストを実行する
miss run test:regress # SQL 回帰テストのみ
./scripts/run-tests.sh 18 all # 特定の PG バージョン
./scripts/run-tests.sh ../postgres/install_tap Tap # ローカル PG ビルドを使用した TAP テスト
テストの種類、TAP テストのセットアップ、テスト ファイルの完全なリストについては、docs/testing.md を参照してください。
一般的な問題: 拡張機能がロードされない (shared_preload_libraries を確認)、イベントが表示されない (pg_stat_ch_stats() でエラーを確認)、キューの使用量が多いかイベントがドロップされた (pg_stat_ch.queue_capacity 、 pg_stat_ch.flush_interval_ms 、 pg_stat_ch.batch_max を調整)。
詳細については docs/troubleshooting.md を参照してください。

詳細な解決策。
このプロジェクトは、Apache License バージョン 2.0 に基づいてライセンスされています。ライセンスの全文については、LICENSE.md を参照してください。
pg-stat-ch.clickhouse.com/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
4
フォーク
レポートリポジトリ
リリース
20
v0.3.11
最新の
2026 年 6 月 13 日
+ 19 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to ClickHouse/pg_stat_ch development by creating an account on GitHub.

GitHub - ClickHouse/pg_stat_ch · GitHub
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
ClickHouse
/
pg_stat_ch
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
114 Commits 114 Commits .claude/ skills .claude/ skills .github .github cmake cmake docker docker docs docs images images include include schema/ migrations schema/ migrations scripts scripts specs specs sql sql src src t t test/ regression test/ regression third_party third_party triplets triplets .clang-format .clang-format .clang-tidy .clang-tidy .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules CLAUDE.md CLAUDE.md CMakeLists.txt CMakeLists.txt CMakePresets.json CMakePresets.json INSTALL.md INSTALL.md LICENSE.md LICENSE.md META.json META.json Makefile Makefile README.md README.md mise.toml mise.toml pg_stat_ch.control pg_stat_ch.control vcpkg-configuration.json vcpkg-configuration.json vcpkg.json vcpkg.json View all files Repository files navigation
pg_stat_ch: PostgreSQL Query Telemetry Exporter to ClickHouse
A PostgreSQL extension that captures per-query execution telemetry and exports it to ClickHouse in real-time. Unlike pg_stat_statements which aggregates statistics in PostgreSQL, pg_stat_ch exports raw events to ClickHouse where aggregation happens via ClickHouse's powerful analytical engine.
Run local PostgreSQL + ClickHouse with schema preloaded:
./scripts/quickstart.sh up
./scripts/quickstart.sh check
Stop:
./scripts/quickstart.sh down
See docker/quickstart/README.md for endpoints and stack details.
Recommended PostgreSQL Settings
pg_stat_ch captures detailed telemetry for every query executed in PostgreSQL and exports it to ClickHouse via a single data pipeline:
PostgreSQL Hooks (foreground) → Shared Memory Queue → Background Worker → ClickHouse
Key design principles:
Zero network I/O on query path - Events are queued in shared memory, not sent synchronously
Raw events, not aggregates - All aggregation (p50/p95/p99, top queries, errors) happens in ClickHouse
Bounded memory - Fixed-size ring buffer with overflow counters; dropped events don't block queries
Minimal overhead - ~5μs p99 overhead per captured statement
PostgreSQL Hooks capture query start/end with full instrumentation
Shared Memory Ring Buffer stores events (MPSC: multi-producer, single-consumer)
Background Worker dequeues batches and exports to ClickHouse
ClickHouse stores raw events in events_raw table; views/MVs provide aggregates
Full Query Telemetry : Timing, row counts, buffer usage, WAL usage, CPU time
All Statement Types : DML (SELECT/INSERT/UPDATE/DELETE/MERGE), DDL, utility statements
Error Capture : SQLSTATE codes and error levels via emit_log_hook
JIT Instrumentation (PG15+): Function count, generation/inlining/optimization/emission time
Parallel Worker Stats (PG18+): Planned vs launched workers
Client Context : Application name, client IP address
Query Text : Captured with truncation (2KB max)
Graceful Degradation : Queue overflow drops events with counters; ClickHouse unavailability doesn't block PostgreSQL
PostgreSQL 16, 17, and 18 are fully supported. See docs/version-compatibility.md for the feature matrix and version-specific fields.
CMake 3.16+, C++17 compiler (GCC 9+, Clang 10+)
PostgreSQL 16+ development headers
mise (recommended) or manual PostgreSQL installation
clickhouse-c is vendored as a submodule and statically linked.
git submodule update --init --recursive
# Using mise (recommended)
mise run build # Debug build (use build:16/17/18 for specific versions)
mise run build:release # Release build
mise run install # Install the extension
# Or manually
cmake -B build -G Ninja -DPG_CONFIG=/path/to/pg_config
cmake --build build && cmake --install build
Installation
shared_preload_libraries = ' pg_stat_ch '
track_io_timing = on # Enables I/O timing columns for captured events
# ClickHouse connection (change for your setup)
pg_stat_ch.clickhouse_host = ' localhost '
pg_stat_ch.clickhouse_port = 9000
pg_stat_ch.clickhouse_database = ' pg_stat_ch '
# TLS (recommended for production)
pg_stat_ch.clickhouse_use_tls = on
pg_stat_ch.clickhouse_skip_tls_verify = off
2. Restart PostgreSQL
# Restart PostgreSQL using your service manager (systemd, brew services, Docker, etc.)
3. Set Up Schema on ClickHouse
Quickstart path: run ./scripts/quickstart.sh up and schema setup is handled automatically.
Manual / existing ClickHouse paths:
clickhouse-client < docker/init/00-schema.sql
Creates events_raw plus materialized views. See docs/clickhouse.md for full schema details.
CREATE EXTENSION pg_stat_ch;
5. Verify
SELECT pg_stat_ch_version();
SELECT * FROM pg_stat_ch_stats();
Configuration
Parameter
Type
Default
Reload
Description
pg_stat_ch.enabled
bool
on
SIGHUP
Enable/disable telemetry collection
pg_stat_ch.clickhouse_host
string
localhost
Restart
ClickHouse server hostname
pg_stat_ch.clickhouse_port
int
9000
Restart
ClickHouse native protocol port
pg_stat_ch.clickhouse_user
string
default
Restart
ClickHouse username
pg_stat_ch.clickhouse_password
string
""
Restart
ClickHouse password
pg_stat_ch.clickhouse_database
string
pg_stat_ch
Restart
ClickHouse database name
pg_stat_ch.queue_capacity
int
65536
Restart
Ring buffer size (must be power of 2)
pg_stat_ch.string_area_size
int
64MB
Restart
DSA size for variable-length query and error strings
pg_stat_ch.flush_interval_ms
int
1000
SIGHUP
Export batch interval in milliseconds
pg_stat_ch.batch_max
int
10000
SIGHUP
Maximum events per ClickHouse insert
pg_stat_ch.clickhouse_use_tls
bool
off
Restart
Enable TLS for ClickHouse connections
pg_stat_ch.clickhouse_skip_tls_verify
bool
off
Restart
Skip TLS certificate verification (insecure)
pg_stat_ch.log_min_elevel
enum
warning
Superuser
Minimum error level to capture (debug5..panic)
See Error Level Values for the complete list of error levels and their numeric values in ClickHouse.
Function
Description
pg_stat_ch_version()
Returns extension version string
pg_stat_ch_stats()
Queue and exporter statistics ( column details )
pg_stat_ch_reset()
Reset all queue counters to zero
pg_stat_ch_flush()
Trigger immediate flush of queued events to ClickHouse
Example Queries
Raw per-execution events enable percentiles, time-series, and per-app drill-downs:
-- Slowest queries for a specific application, with percentiles
SELECT query_id,
count () AS calls,
quantile( 0 . 95 )(duration_us) / 1000 AS p95_ms,
quantile( 0 . 99 )(duration_us) / 1000 AS p99_ms
FROM pg_stat_ch . events_raw
WHERE app = ' myapp '
AND ts_start > now() - INTERVAL 1 HOUR
GROUP BY query_id
ORDER BY p99_ms DESC
LIMIT 10 ;
See docs/clickhouse.md for materialized view definitions and more query examples.
mise run test:all # Run all tests
mise run test:regress # SQL regression tests only
./scripts/run-tests.sh 18 all # Specific PG version
./scripts/run-tests.sh ../postgres/install_tap tap # TAP tests with local PG build
See docs/testing.md for test types, TAP test setup, and a full listing of test files.
Common issues: extension not loading (check shared_preload_libraries ), events not appearing (check pg_stat_ch_stats() for errors), high queue usage or dropped events (tune pg_stat_ch.queue_capacity , pg_stat_ch.flush_interval_ms , pg_stat_ch.batch_max ).
See docs/troubleshooting.md for detailed solutions.
This project is licensed under the Apache License, Version 2.0. See LICENSE.md for the full license text.
pg-stat-ch.clickhouse.com/
Topics
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
4
forks
Report repository
Releases
20
v0.3.11
Latest
Jun 13, 2026
+ 19 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
