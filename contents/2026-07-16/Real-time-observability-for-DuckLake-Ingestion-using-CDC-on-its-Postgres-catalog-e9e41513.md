---
source: "https://github.com/el10savio/duckWatch"
hn_url: "https://news.ycombinator.com/item?id=48933734"
title: "Real-time observability for DuckLake Ingestion using CDC on its Postgres catalog"
article_title: "GitHub - el10savio/duckWatch: Real-time observability for DuckLake Ingestion levaraging CDC on the attached Postgres catalog using Debezium, a Go collector, Kafka and Clickhouse. · GitHub"
author: "ugabuga"
captured_at: "2026-07-16T13:21:05Z"
capture_tool: "hn-digest"
hn_id: 48933734
score: 1
comments: 0
posted_at: "2026-07-16T12:47:36Z"
tags:
  - hacker-news
  - translated
---

# Real-time observability for DuckLake Ingestion using CDC on its Postgres catalog

- HN: [48933734](https://news.ycombinator.com/item?id=48933734)
- Source: [github.com](https://github.com/el10savio/duckWatch)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T12:47:36Z

## Translation

タイトル: Postgres カタログで CDC を使用した DuckLake インジェストのリアルタイム可観測性
記事のタイトル: GitHub - el10savio/duckWatch: Debezium、Go コレクター、Kafka、Clickhouse を使用した、添付された Postgres カタログ上の CDC を活用した DuckLake Ingestion のリアルタイム監視。 · GitHub
説明: Debezium、Go コレクター、Kafka、および Clickhouse を使用して、添付された Postgres カタログ上で CDC を利用する DuckLake Ingestion のリアルタイム監視。 - el10savio/ダックウォッチ

記事本文:
GitHub - el10savio/duckWatch: Debezium、Go コレクター、Kafka、Clickhouse を使用した、添付された Postgres カタログ上の CDC を活用した DuckLake Ingestion のリアルタイム監視。 · GitHub
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
エル10サビオ
/
アヒルウォッチ
公共
通知
あなたはしなければなりません

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .github/ workflows .github/ workflows clickhouse clickhouse cmd cmd config/ debezium config/ debeziumdeploy/ supersetdeploy/ superset docs docs scripts scripts src src .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum setup.sh setup.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Debezium、Go コレクター、Kafka、および Clickhouse を使用して、接続された Postgres カタログ データベース上の CDC を活用する DuckLake のリアルタイム監視。
ファイルベースのカタログ (Iceberg、Delta) には、コミットされた状態 (スナップショットとファイル) のみが記録されます。
生き残った統計。 DuckLake のカタログは実行中のデータベースであるため、さらに
ファイル カタログが構造的に生成できないシグナルのクラス全体を公開します。
ポーリング頻度 (実行中のトランザクション、ライブロック競合、コミットなど)
共有の ducklake_snapshot 上で競合するオプティミスティック同時実行ライターの競合/再試行率。 CDC とスケジュールされた Go ポーラーを利用してこれらのイベントを取り込み、Kafka 経由で ClickHouse にデータを取り込んで分析を取得します。
Docker (Compose v2) および Go 1.25 が必要です
セットアップ # 書き込みパス + 可観測性プレーン、有線エンドツーエンド
スーパーセットアップ # ダッシュボードを http://localhost:8088 (admin/admin) に作成します
makeload RPS=10 WRITERS=4 # ロードを再実行します。ライターが増える -> 争い/衝突が増える
分解を行う # コネクタとボリュームを削除する
スーパーセット
Debezium、Go コレクター、Kafka、および Clickhouse を使用して、添付された Postgres カタログ上の CDC を利用する DuckLake Ingestion のリアルタイムの可観測性。
読み込み中にエラーが発生しました。

このページをリロードしてください。
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

Real-time observability for DuckLake Ingestion levaraging CDC on the attached Postgres catalog using Debezium, a Go collector, Kafka and Clickhouse. - el10savio/duckWatch

GitHub - el10savio/duckWatch: Real-time observability for DuckLake Ingestion levaraging CDC on the attached Postgres catalog using Debezium, a Go collector, Kafka and Clickhouse. · GitHub
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
el10savio
/
duckWatch
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github/ workflows .github/ workflows clickhouse clickhouse cmd cmd config/ debezium config/ debezium deploy/ superset deploy/ superset docs docs scripts scripts src src .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum setup.sh setup.sh View all files Repository files navigation
Real-time observability for a DuckLake levaraging CDC on the attached Postgres catalog database using Debezium, a Go collector, Kafka and Clickhouse.
A file-based catalog (Iceberg, Delta) records only committed state: the snapshots and file
statistics that survived. DuckLake's catalog is a running database, so it additionally
exposes an entire class of signal that a file catalog structurally cannot produce at any
polling frequency, like in-flight transactions, live lock contention, and the commit
conflict/retry rate of optimistic-concurrency writers racing on the shared ducklake_snapshot . We leaverage CDC and a scheduled Go poller to ingest these events and populate ClickHouse via Kafka to get analytics on them.
Requires Docker (Compose v2) and Go 1.25
make setup # write path + observability plane, wired end to end
make superset-up # dashboard at http://localhost:8088 (admin/admin)
make load RPS=10 WRITERS=4 # re-run load; more writers -> more contention/conflict
make teardown # drop the connector and volumes
Superset
Real-time observability for DuckLake Ingestion levaraging CDC on the attached Postgres catalog using Debezium, a Go collector, Kafka and Clickhouse.
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
