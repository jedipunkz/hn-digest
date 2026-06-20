---
source: "https://github.com/jjviscomi/bqemulator"
hn_url: "https://news.ycombinator.com/item?id=48610052"
title: "New GCP Big Query Emulator"
article_title: "GitHub - jjviscomi/bqemulator: Local emulator for Google BigQuery. DuckDB-backed, SQLGlot-powered. Drop-in replacement for the real service in dev, CI, and offline replicas. · GitHub"
author: "jjviscomi"
captured_at: "2026-06-20T18:37:11Z"
capture_tool: "hn-digest"
hn_id: 48610052
score: 3
comments: 1
posted_at: "2026-06-20T15:42:15Z"
tags:
  - hacker-news
  - translated
---

# New GCP Big Query Emulator

- HN: [48610052](https://news.ycombinator.com/item?id=48610052)
- Source: [github.com](https://github.com/jjviscomi/bqemulator)
- Score: 3
- Comments: 1
- Posted: 2026-06-20T15:42:15Z

## Translation

タイトル: 新しい GCP Big Query エミュレータ
記事タイトル: GitHub - jjviscomi/bqemulator: Google BigQuery のローカル エミュレーター。 DuckDB を利用し、SQLGlot を利用しています。開発、CI、オフライン レプリカの実際のサービスをドロップインで置き換えます。 · GitHub
説明: Google BigQuery のローカル エミュレータ。 DuckDB を利用し、SQLGlot を利用しています。開発、CI、オフライン レプリカの実際のサービスをドロップインで置き換えます。 - jjviscomi/bqemulator

記事本文:
GitHub - jjviscomi/bqemulator: Google BigQuery のローカル エミュレーター。 DuckDB を利用し、SQLGlot を利用しています。開発、CI、オフライン レプリカの実際のサービスをドロップインで置き換えます。 · GitHub
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
ジェビスコミ
/
bqエミュレータ
公共
通知
あなたはきっとsでしょう

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
148 コミット 148 コミット .github .github docs docs fuzz fuzz scripts scripts src/ bqemulator src/ bqemulator テスト テスト .coderabbit.yaml .coderabbit.yaml .commitlintrc.yaml .commitlintrc.yaml .gitignore .gitignore .jscpd.json .jscpd.json .lychee.toml .lychee.toml .pip-audit-ignore .pip-audit-ignore .pre-commit-config.yaml .pre-commit-config.yaml .tool-versions .tool-versions .vulture_whitelist.py .vulture_whitelist.py AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md _typos.toml _typos.toml codecov.yml codecov.yml docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml osv-scanner.toml osv-scanner.toml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Google BigQuery 用のローカルのドロップイン エミュレータ。
DuckDB を利用し、SQLGlot を利用し、実際のサービスに対してテストされています。公式の Google Cloud クライアント ライブラリを指定して、ラップトップまたは CI で BigQuery コードを実行します。実際のプロジェクト、請求、ネットワークは必要ありません。
ドキュメント
· クイックスタート
· 例
· 互換性マトリックス
· 変更履歴
実際の BigQuery に対してコードをテストするのは遅く（ネットワーク + サービスのレイテンシ）、費用がかかり（すべてのクエリが課金される）、危険が伴います（共有環境ではロールバックがありません）。代替手段 (モック、偽物、共有サンドボックス) は、追いかけるのをやめた瞬間に本物のサービスから離れてしまいます。
bqemulator は、BigQuery の実際のワイヤ プロトコル（REST +

gRPC) を実際の分析 SQL エンジン (DuckDB) に戻し、ルールベースの ADR ベースのトランスレーター (SQLGlot) を使用して GoogleSQL → DuckDB SQL を変換します。公式の google-cloud-bigquery 、 @google-cloud/bigquery 、cloud.google.com/go/bigquery 、 com.google.cloud:google-cloud-bigquery 、および bq CLI クライアントはすべて、変更されることなく動作します。エンドポイントのみが異なります。
一時的な CI フィクスチャ — pytest プラグインは、ランダムなポートでインプロセス エミュレータを開始します。 pip install bqemulator[testing] は必要な配線のすべてです。
長時間実行されるローカル開発サーバー — bqemulator start --data-dir ~/bqemu は実行中も状態を保持します。公式の bq CLI、dbt、Airflow、PySpark、Beam、Scio で動作します。
実際のプロジェクトのオフライン レプリカ — bqemulator import --from-project <id> は、実際の BigQuery からローカル データ ディレクトリにスキーマ (およびオプションでデータ) をクローンします。
🟢 完全な REST + gRPC API パリティ — データセット、テーブル、ジョブ、TableData、ルーチン、行アクセス ポリシー、承認されたビュー、およびモデル CRUD メタデータ。ストレージ読み取り API (Arrow および Avro)。ストレージ書き込み API (4 つのストリーム タイプすべて - DEFAULT 、 COMMITTED 、 PENDING 、 BUFFERED - proto および Arrow row 形式の両方)。
⚡ 本物の SQL — GoogleSQL は 92 の SQLGlot ルール + 24 のリライターを介して DuckDB SQL に変換されました。日付/時刻、文字列、配列、構造体、範囲、地理、JSON、近似集計、統計、正規表現、常用時間、およびビット操作をカバーします。
🧠 goccy/bigquery-emulator にはない機能 - JavaScript UDF (mini-racer を介した組み込み V8)、手続き型スクリプト (DECLARE / BEGIN…END / IF / LOOP / EXCEPTION / BEGIN TRANSACTION)、タイムトラベル (FOR SYSTEM_TIME AS OF)、テーブル スナップショット、テーブル クローン、リフレッシュ ディスパッチを伴うマテリアライズド ビュー、GEOGRAPHY (DuckDB 空間 + S2 ヘルパーを介した平面)、RANGE、INTERVAL、承認されたビュー、行アクセス ポリシー、INFORMA

TION_SCHEMA 。
🔌 5 クライアントの e2e マトリックス — すべてのリリースは、公式の Python、Node.js、Go、Java BigQuery クライアント ライブラリと、ライブ Docker コンテナ内の Google の bq CLI に対して実行されます。
🧪 7 層テスト ピラミッド — ユニット + プロパティ + 統合 + 適合性 + e2e + perf + カオス、さらに突然変異 / ファズ / 差分兄弟。結合されたカバレッジは、回線 + 分岐の 90% 以上でゲートされます。
📐 適合コーパス — 実際の BigQuery に対して記録された 1,288 のフィクスチャ。エミュレータと実際のサービスの間の漂流は、テストの失敗として表面化します。文書化された相違は ADR 参照とともに固定されます。
🐍 ネイティブ pytest プラグイン — pip install bqemulator は pytest プラグインを登録します。 bqemu_server フィクスチャは、ランダムな空きポートで一時的なインプロセス エミュレータを起動し、 BIGQUERY_EMULATOR_HOST を設定します。 conftest.py の配線は必要ありません。
🐳 マルチ アーキテクチャ コンテナ — ghcr.io/jjviscomi/bqemulator は、GitHub OIDC を介した cosign キーレス署名を使用して、 linux/amd64 + linux/arm64 用にビルドされます。
🔭 本番グレードの可観測性 — structlog JSON ログ、OpenTelemetry トレース (構成可能な OTLP エクスポータ)、Prometheus メトリクス エンドポイント。
pip インストール bqemulator
オプションの追加物:
pip install " bqemulator[testing] " # pytest、仮説、testcontainers、bigquery クライアント
pip install " bqemulator[udf-js] " # JavaScript UDF サポート (組み込み V8)
pip install " bqemulator[orc] " # ロードジョブの ORC 形式
pip install " bqemulator[compression] " # ロード/抽出ジョブ用の zstd + snappy
pip install " bqemulator[import] " # bqemulator import --from-project
pip install " bqemulator[all] " # すべてのランタイム エクストラ (テスト用エクストラはなし)
ドッカー:
docker run --rm -p 9050:9050 -p 9060:9060 ghcr.io/jjviscomi/bqemulator:latest
pip と公開されたイメージはどちらも同じエミュレータをバンドルしています。このイメージは、9050 で REST を公開し、9060 で gRPC を公開します。

障害 — 変更するには構成リファレンスを参照してください。
Windows ユーザー: WSL2 バックエンド (Docker Desktop 4.x 以降のデフォルト) を使用して Windows 用 Docker Desktop をインストールします。公開された Linux イメージは、Windows 固有の構成を必要とせずに、WSL2 でネイティブに実行されます。イメージのネイティブ Windows コンテナーのバリアントは、v1.0 の範囲外であることが明示されています。その根拠については、docs/reference/out-of-scope.md#native-windows-containers を参照してください。
OSをインポートする
グーグルから。クラウドインポートbigquery
# BIGQUERY_EMULATOR_HOST を設定します (すべての Google Cloud ライブラリによって選択されます)
# または api_endpoint をクライアントに明示的に渡します。どちらも機能します。
OS 。環境 [ "BIGQUERY_EMULATOR_HOST" ] = "ローカルホスト:9050"
クライアント = bigquery 。クライアント ( project = "my-test-project" )
クライアント。 create_dataset ( "売上" )
クライアント。 create_table (
ビッグクエリ。テーブル (
"販売.注文" 、
スキーマ = [
ビッグクエリ。スキーマフィールド ( "id" 、 "INT64" )、
ビッグクエリ。 SchemaField ( "amount" 、 "NUMERIC" )、
ビッグクエリ。スキーマフィールド ( "placed_at" 、 "TIMESTAMP" )、
]、
)
)
クライアント。 insert_rows_json (
"販売.注文" 、
[{ "id" : 1 , "amount" : "12.50" , "placed_at" : "2026-05-21T00:00:00Z" }],
)
client の行の場合。クエリ (「SELECT COUNT(*) AS n FROM sales.orders」)。結果():
print ( row .n ) #1
pytest
bqemulator には、pytest11 エントリ ポイント経由で pytest プラグインが同梱されています。パッケージをインストールするだけで必要な作業がすべて完了します。conftest.py は空のままです。
グーグルから。クラウドインポートbigquery
def test_orders_table(bqemu_client:bigquery.Client) -> なし:
bqemu_client 。 create_dataset ( "売上" )
# ... あなたのテスト ...
bqemu_server フィクスチャはセッション スコープです (テスト セッションごとに 1 つのエミュレータ)。 bqemu_client フィクスチャは関数スコープであり、事前構成された bigquery.Client を返します。完全な Flask アプリについては、pytest フィクスチャ ガイドと python/pytest-integration の例を参照してください。

h 統合テスト。
const { BigQuery } = require ( '@google-cloud/bigquery' ) ;
const bq = 新しい BigQuery ( {
projectId : 'my-test-project' 、
apiEndpoint : 'http://localhost:9050' 、
token : 'dummy' , // エミュレータは任意のトークンを受け入れます
} ) ;
バーベキューを待ちます。 createDataset ( 'sales' ) ;
Node.js クイックスタートと nodejs/nestjs-app の例を参照してください。
クライアント 、 _ := bigquery 。新しいクライアント (
ctx 、 "私のテストプロジェクト" 、
オプション。 WithEndpoint ( "http://localhost:9050" )、
オプション。認証なし ()、
)
Go クイックスタートと go/beam-pipeline の例を参照してください。
BigQuery bq = BigQueryOptions 。 newBuilder()
。 setProjectId ( "私のテストプロジェクト" )
。 setHost ( "http://localhost:9050" )
。 setCredentials ( NoCredentials .getInstance ())
。建てる （）
。 getService();
Java クイックスタートと java/spring-boot の例を参照してください。
bq --api=http://localhost:9050 \
--project_id=私のテストプロジェクト \
クエリ --use_legacy_sql=false ' SELECT 1 AS n '
bq CLI ガイドと bq-cli-quickstart の例を参照してください。
サービス:
bqエミュレータ :
画像：ghcr.io/jjviscomi/bqemulator:最新
ポート: ["9050:9050"、"9060:9060"]
ヘルスチェック:
テスト: ["CMD"、"curl"、"-sf"、"http://localhost:9050/healthz"]
間隔：2秒
再試行: 30
アプリ:
ビルド: 。
環境：
BIGQUERY_EMULATOR_HOST : bqemulator:9050
依存関係 :
bqemulator : { 条件: service_healthy }
アプリ + エミュレーター + Prometheus + Grafana の docker-compose/full-stack の例を参照してください。
bqemulator は v1.3.0 です - 安定版の 3 番目のマイナーです
ライン。 SemVer 適用: 重大な変更は MAJOR でのみ出荷されます。
非推奨の存続期間はマイナー 2 か月または 6 か月以上です。互換性マトリックスは、CI 実行ごとに適合性コーパスから自動生成されます。適合カバレッジ マトリックスは、表面項目ごとにサポートを分類します。
適合コーパスの深さ (適合カバレッジ マトリックスには、ライブの自動生成されたデータが含まれます)

d 内訳）：
さらに、ADR 0022 によって適合性コーパスから除外され、代わりにユニット/プロパティ/統合層で実行される 10 個の非決定的項目 ( RAND 、 CURRENT_* 、 SESSION_USER 、 GENERATE_UUID 、 TABLESAMPLE 、 FOR SYSTEM_TIME AS OF <expression> ) testing/conformance/ の下の 1,288 フィクスチャ適合コーパス (1,213 SQL + 49 HTTP + 26 gRPC)。
当社は延期なしの原則に従っており、機能は完成した状態で出荷されるか、文書化された根拠とともに除外されます。 「v1.1のTODO」はありません。スコープ境界は docs/reference/out-of-scope.md にカタログ化されています。
完全なドキュメントは jjviscomi.github.io/bqemulator にあります。主要なエントリーポイント:
はじめに — 最初の 10 分間。
言語ごとのクイックスタート — Python、Node.js、Go、Java、pytest、docker-compose、Testcontainers。
ガイド — データのロード、クエリ、ストリーミング挿入、ストレージ API、UDF、スクリプト、パーティショニング、タイム トラベル、マテリアライズド ビュー、行アクセス ポリシー、dbt、Airflow、Spark、bq CLI、可観測性など。
リファレンス - 構成、CLI、REST カバレッジ、SQL 関数マッピング、互換性マトリックス、適合カバレッジ マトリックス、範囲外のカタログ、トラブルシューティング。
アーキテクチャ — ヘキサゴナル アーキテクチャ、ストレージ モード

[切り捨てられた]

## Original Extract

Local emulator for Google BigQuery. DuckDB-backed, SQLGlot-powered. Drop-in replacement for the real service in dev, CI, and offline replicas. - jjviscomi/bqemulator

GitHub - jjviscomi/bqemulator: Local emulator for Google BigQuery. DuckDB-backed, SQLGlot-powered. Drop-in replacement for the real service in dev, CI, and offline replicas. · GitHub
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
jjviscomi
/
bqemulator
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
148 Commits 148 Commits .github .github docs docs fuzz fuzz scripts scripts src/ bqemulator src/ bqemulator tests tests .coderabbit.yaml .coderabbit.yaml .commitlintrc.yaml .commitlintrc.yaml .gitignore .gitignore .jscpd.json .jscpd.json .lychee.toml .lychee.toml .pip-audit-ignore .pip-audit-ignore .pre-commit-config.yaml .pre-commit-config.yaml .tool-versions .tool-versions .vulture_whitelist.py .vulture_whitelist.py AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md _typos.toml _typos.toml codecov.yml codecov.yml docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml osv-scanner.toml osv-scanner.toml pyproject.toml pyproject.toml View all files Repository files navigation
A local, drop-in emulator for Google BigQuery.
DuckDB-backed, SQLGlot-powered, and tested against the real service. Point the official Google Cloud client libraries at it and run your BigQuery code on your laptop or in CI — no real project, no billing, no network.
Documentation
· Quickstart
· Examples
· Compatibility matrix
· Changelog
Testing code against real BigQuery is slow (network + service latency), expensive (every query is billable), and dangerous (no rollback in shared environments). The alternatives — mocks, fakes, and shared sandboxes — drift from the real service the moment you stop chasing them.
bqemulator is a process you can run locally that speaks BigQuery's actual wire protocol (REST + gRPC), backs onto a real analytical SQL engine (DuckDB), and translates GoogleSQL → DuckDB SQL with a rule-based, ADR-grounded translator (SQLGlot). The official google-cloud-bigquery , @google-cloud/bigquery , cloud.google.com/go/bigquery , com.google.cloud:google-cloud-bigquery , and bq CLI clients all work against it unchanged — only the endpoint differs.
Ephemeral CI fixture — pytest plugin starts an in-process emulator on a random port; pip install bqemulator[testing] is all the wiring you need.
Long-running local dev server — bqemulator start --data-dir ~/bqemu persists state across runs; works with the official bq CLI, dbt, Airflow, PySpark, Beam, Scio.
Offline replica of a real project — bqemulator import --from-project <id> clones schema (and optionally data) from real BigQuery into a local data directory.
🟢 Full REST + gRPC API parity — Datasets, Tables, Jobs, TableData, Routines, Row Access Policies, Authorized Views, plus Models CRUD metadata. Storage Read API (Arrow and Avro). Storage Write API (all four stream types — DEFAULT , COMMITTED , PENDING , BUFFERED — with both proto and Arrow row formats).
⚡ Real SQL — GoogleSQL translated to DuckDB SQL via 92 SQLGlot rules + 24 rewriters; covers date/time, string, array, struct, range, geography, JSON, approximate-aggregate, statistical, regex, civil-time, and bit operations.
🧠 Features goccy/bigquery-emulator doesn't have — JavaScript UDFs (embedded V8 via mini-racer ), procedural scripting ( DECLARE / BEGIN…END / IF / LOOP / EXCEPTION / BEGIN TRANSACTION ), time travel ( FOR SYSTEM_TIME AS OF ), table snapshots, table clones, materialized views with refresh dispatch, GEOGRAPHY (planar via DuckDB-spatial + S2 helpers), RANGE, INTERVAL, authorized views, row-access policies, INFORMATION_SCHEMA .
🔌 Five-client e2e matrix — every release is exercised against the official Python, Node.js, Go, and Java BigQuery client libraries plus Google's bq CLI in a live Docker container.
🧪 7-tier test pyramid — unit + property + integration + conformance + e2e + perf + chaos, plus mutation / fuzz / differential siblings. Combined coverage is gated at ≥90% line + branch.
📐 Conformance corpus — 1,288 fixtures recorded against real BigQuery. Drift between the emulator and the real service surfaces as a failing test; documented divergences are pinned with ADR references.
🐍 Native pytest plugin — pip install bqemulator registers a pytest plugin; the bqemu_server fixture starts an ephemeral in-process emulator on random free ports and sets BIGQUERY_EMULATOR_HOST . No conftest.py wiring required.
🐳 Multi-arch container — ghcr.io/jjviscomi/bqemulator builds for linux/amd64 + linux/arm64 , with cosign keyless signatures via GitHub OIDC.
🔭 Production-grade observability — structlog JSON logs, OpenTelemetry tracing (configurable OTLP exporter), Prometheus metrics endpoint.
pip install bqemulator
Optional extras:
pip install " bqemulator[testing] " # pytest, hypothesis, testcontainers, bigquery client
pip install " bqemulator[udf-js] " # JavaScript UDF support (embedded V8)
pip install " bqemulator[orc] " # ORC format for load jobs
pip install " bqemulator[compression] " # zstd + snappy for load/extract jobs
pip install " bqemulator[import] " # bqemulator import --from-project
pip install " bqemulator[all] " # all runtime extras (no testing extras)
Docker:
docker run --rm -p 9050:9050 -p 9060:9060 ghcr.io/jjviscomi/bqemulator:latest
Both pip and the published image bundle the same emulator. The image exposes REST on 9050 and gRPC on 9060 by default — see configuration reference to change them.
Windows users: install Docker Desktop for Windows with the WSL2 backend (default since Docker Desktop 4.x); the published Linux image runs natively under WSL2 with no Windows-specific configuration. Native Windows-container variants of the image are explicitly out of scope for v1.0 — see docs/reference/out-of-scope.md#native-windows-containers for the rationale.
import os
from google . cloud import bigquery
# Either set BIGQUERY_EMULATOR_HOST (picked up by every Google Cloud library)
# or pass api_endpoint explicitly to the Client. Both work.
os . environ [ "BIGQUERY_EMULATOR_HOST" ] = "localhost:9050"
client = bigquery . Client ( project = "my-test-project" )
client . create_dataset ( "sales" )
client . create_table (
bigquery . Table (
"sales.orders" ,
schema = [
bigquery . SchemaField ( "id" , "INT64" ),
bigquery . SchemaField ( "amount" , "NUMERIC" ),
bigquery . SchemaField ( "placed_at" , "TIMESTAMP" ),
],
)
)
client . insert_rows_json (
"sales.orders" ,
[{ "id" : 1 , "amount" : "12.50" , "placed_at" : "2026-05-21T00:00:00Z" }],
)
for row in client . query ( "SELECT COUNT(*) AS n FROM sales.orders" ). result ():
print ( row . n ) # 1
pytest
bqemulator ships a pytest plugin via the pytest11 entry point. Installing the package is all the wiring you need — your conftest.py stays empty.
from google . cloud import bigquery
def test_orders_table ( bqemu_client : bigquery . Client ) -> None :
bqemu_client . create_dataset ( "sales" )
# ... your test ...
The bqemu_server fixture is session-scoped (one emulator per test session); the bqemu_client fixture is function-scoped and returns a pre-configured bigquery.Client . See the pytest fixture guide and the python/pytest-integration example for a complete Flask app with integration tests.
const { BigQuery } = require ( '@google-cloud/bigquery' ) ;
const bq = new BigQuery ( {
projectId : 'my-test-project' ,
apiEndpoint : 'http://localhost:9050' ,
token : 'dummy' , // emulator accepts any token
} ) ;
await bq . createDataset ( 'sales' ) ;
See the Node.js quickstart and the nodejs/nestjs-app example.
client , _ := bigquery . NewClient (
ctx , "my-test-project" ,
option . WithEndpoint ( "http://localhost:9050" ),
option . WithoutAuthentication (),
)
See the Go quickstart and the go/beam-pipeline example.
BigQuery bq = BigQueryOptions . newBuilder ()
. setProjectId ( "my-test-project" )
. setHost ( "http://localhost:9050" )
. setCredentials ( NoCredentials . getInstance ())
. build ()
. getService ();
See the Java quickstart and the java/spring-boot example.
bq --api=http://localhost:9050 \
--project_id=my-test-project \
query --use_legacy_sql=false ' SELECT 1 AS n '
See the bq CLI guide and the bq-cli-quickstart example.
services :
bqemulator :
image : ghcr.io/jjviscomi/bqemulator:latest
ports : ["9050:9050", "9060:9060"]
healthcheck :
test : ["CMD", "curl", "-sf", "http://localhost:9050/healthz"]
interval : 2s
retries : 30
app :
build : .
environment :
BIGQUERY_EMULATOR_HOST : bqemulator:9050
depends_on :
bqemulator : { condition: service_healthy }
See the docker-compose/full-stack example for app + emulator + Prometheus + Grafana.
bqemulator is at v1.3.0 — third minor on the production-stable
line. SemVer applies: breaking changes ship only in MAJOR,
deprecations live ≥2 MINOR or 6 months. The compatibility matrix is auto-generated from the conformance corpus on every CI run; the conformance coverage matrix breaks down support by surface item.
Conformance corpus depth (the conformance coverage matrix carries the live, auto-generated breakdown):
Plus 10 non-deterministic items ( RAND , CURRENT_* , SESSION_USER , GENERATE_UUID , TABLESAMPLE , FOR SYSTEM_TIME AS OF <expression> ) that are excluded from the conformance corpus by ADR 0022 and exercised in unit / property / integration tiers instead — bringing the full inventory to 419 surface items across 20 categories , backed by a 1,288-fixture conformance corpus (1,213 SQL + 49 HTTP + 26 gRPC) under tests/conformance/ .
We follow a no-deferral principle : features either ship complete or are excluded with documented rationale. There is no "TODO for v1.1." Scope boundaries are catalogued in docs/reference/out-of-scope.md .
The full documentation lives at jjviscomi.github.io/bqemulator . Key entry points:
Getting started — your first ten minutes.
Per-language quickstarts — Python · Node.js · Go · Java · pytest · docker-compose · Testcontainers.
Guides — loading data, querying, streaming inserts, Storage API, UDFs, scripting, partitioning, time travel, materialized views, row access policies, dbt, Airflow, Spark, the bq CLI, observability, and more.
Reference — configuration, CLI, REST coverage, SQL function mapping, compatibility matrix, conformance coverage matrix, out-of-scope catalogue, troubleshooting.
Architecture — hexagonal architecture, storage mode

[truncated]
