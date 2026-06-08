---
source: "https://github.com/kenm47/dbmachine"
hn_url: "https://news.ycombinator.com/item?id=48446233"
title: "Show HN: dbmachine – the back end and front end are Claude, just the plumbing&DB"
article_title: "GitHub - kenm47/dbmachine: Turn a Postgres database into a self-describing, agent-operable application backend. · GitHub"
author: "hank2000"
captured_at: "2026-06-08T16:30:11Z"
capture_tool: "hn-digest"
hn_id: 48446233
score: 1
comments: 1
posted_at: "2026-06-08T14:56:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: dbmachine – the back end and front end are Claude, just the plumbing&DB

- HN: [48446233](https://news.ycombinator.com/item?id=48446233)
- Source: [github.com](https://github.com/kenm47/dbmachine)
- Score: 1
- Comments: 1
- Posted: 2026-06-08T14:56:35Z

## Translation

タイトル: Show HN: dbmachine – バックエンドとフロントエンドはクロード、配管と DB だけです
記事のタイトル: GitHub - kenm47/dbmachine: Postgres データベースを自己記述型のエージェント操作可能なアプリケーション バックエンドに変える。 · GitHub
説明: Postgres データベースを、自己記述型のエージェント操作可能なアプリケーション バックエンドに変えます。 - kenm47/dbmachine

記事本文:
GitHub - kenm47/dbmachine: Postgres データベースを自己記述型のエージェント操作可能なアプリケーション バックエンドに変えます。 · GitHub
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
ケンム47
/
データベースマシン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーションオプション

イオン
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github/ workflows .github/ workflows docs docs サンプル サンプル src/ dbmachine src/ dbmachine テスト テスト .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SPEC.md SPEC.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Postgres データベースを自己記述型のエージェント操作可能なアプリケーション バックエンドに変えます。
ほとんどのビジネス ソフトウェアは、データベース + シン バックエンド + フロントエンドで構成されています。とき
フロントエンドはチャット エージェント (Claude Code など)、欠けている部分は標準です、
Postgres とエージェント間の安全な契約。 dbmachine はそのコントラクトです:
コーディング エージェントが YAML 仕様でアプリを宣言し、dbmachine が実際のアプリにコンパイルします。
Postgres バックエンド (スキーマ、制約、RLS、型付き操作、監査ログ)、および
エージェントが読んだ指示を生成して、1 つの自己文書を実行します。
CLI。
保証はプロンプトではなくデータベース内に存在します。エージェントは腐敗できない
データが間違っている場合でも、DB は不正な状態を拒否します。
🐘 本物の Postgres 、完全にローカル (Docker)、完全にオープン ソース。スーパーベースはありません、いいえ
サードパーティのアカウント。
🤖 エージェントネイティブ : 随所に豊富な説明、詳細な内省、
コマンドごとに機械可読な JSON を出力する CLI。
🔒 構造的安全性 : 必須フィールド、列挙型、一意性、チェックと
除外制約 (ダブルブッキング禁止など)、行レベルのセキュリティ、
完全な監査ログ — すべて仕様に基づいてコンパイルされています。
🧱 クリーンなコア: CLI は構造化リターン Python 上のシン アダプターです
したがって、将来の MCP サーバーは最大 100 行のドロップインになります。
ステータス: v1。単一の信頼できるオペレーター、CLI のみ。 HTTP API、MCPサーバー、および
マルチテナントのエンドユーザー認証は意図的に延期されます ( SPEC.md を参照)。
🤖 ブ

AI エージェントを使用してアプリを作成しますか?に向けてください
docs/AGENT_GUIDE.md 。その単一の自己完結型
この文書では、ゼロコンテキストのエージェントに仕様の作成方法と CRM の立ち上げ方法を教えています
または予約アプリのエンドツーエンド - 仕様リファレンス、完全な CLI、カスタム
操作、移行、エラー コード、およびコピー＆ペースト可能なレシピ。
dbmachine は uv によって管理されるプロジェクトローカルの依存関係です。
グローバルツールではありません。 uvx で一度ブートストラップすると、すべてが
プロジェクト独自の環境 - したがって、カスタム操作で自由にストライプをインポートできます /
パンダをインポートします (uv 追加 ...)。
uvx dbmachine init my_crm # プロジェクトをスキャフォールディングする (ブートストラップのみ)
cd my_crm
uv sync # プロジェクト環境をセットアップする
dbmachine up # ローカル Postgres を起動します (空きポートを自動検出します)
dbmachine merge # 仕様をコンパイル → スキーマを作成
dbmachine Inspection # 何を取得したかを確認します
dbmachine up は空きホスト ポートを自動検出します (従来の 5432 衝突を回避します)
Homebrew Postgres を使用して)、それを .dbmachine/config.json に永続化します。
「クロード、CRM を構築してください」 - ウォークスルー
エージェントは app.dbm.yaml を作成し、完全に CLI を通じてアプリを操作します。
バンドルされた CRM の例 (examples/crm) を使用した、正確な契約は次のとおりです。
# 1. コンパイルと移行
データベースマシンの移行
# {"ok": true, "data": {"applied": true, "fresh": true, ...}}
# 2. スキーマを学ぶ (必要なものだけを引き出す)
データベースマシン検査
dbmachine は取引を説明します
# 3. データを作成します — DB にヒットする前に検証します
dbmachine create company --json ' {"name": "Globex"} '
# {"ok": true, "data": {"id": "…", "name": "Globex", ...}}
dbmachine create deal --json ' {"title": "Big", "amount": 5000, "stage": "proposal", "company_id": "…"} '
# 4. 型指定されたカスタム操作を実行します (Python、トランザクション、副作用スタブあり)
dbmachine do win_deal --json ' {"deal_id": "…"} '
# {"ok": true, "d

ata": {"結果": {"ステージ": "勝利", ...}}}
#5. ビューによる読み取り専用分析
dbmachine query --sql ' SELECT * FROM Pipeline_by_stage '
# 6. ヒューリスティックマッピング + 重複排除による一括インポート (CSV/JSON/Excel)
dbmachine import contacts.csv --entity contact --dedup email
#7. 完全な変更履歴
dbmachine Audit --entity deal
すべてのコマンドは、必要に応じて --json 入力を受け取り、
{"ok": true|false, ...} 標準出力のエンベロープ - エージェントが解析できる信頼性の高いもの。
宣言型でバージョン管理された信頼できる情報源。予約内容より抜粋
例として、構造的なダブルブッキング防止の保証を示します。
エンティティ:
予定:
説明 : 一定期間にわたる 1 人の顧客に対する 1 つのリソースの予約。
フィールド:
starts_at : { タイプ: タイムスタンプ、必須: true }
end_at : { タイプ: タイムスタンプ、必須: true }
ステータス : { タイプ: 列挙型、列挙型: アポイントメントステータス、デフォルト: 保留中 }
関係 :
顧客 : { 所属: 顧客、必須: true }
resource : { 所属先: リソース、必須: true }
チェック:
- { 名前: 開始後終了、式: "終了後 > 開始" }
除外事項:
- 名前 : no_double_booking
付き:
- { フィールド: resource_id、演算子: "= }
- { expr: "tstzrange(starts_at, ends_at)", op: "&&" }
ここで: " ステータス <> 'キャンセル済み' "
その除外ブロックは Postgres GiST 除外制約にコンパイルされます。
エージェントは、どんなに試みても、重複する予約を作成できません。
dbmachine スキーマを実行して仕様の JSON スキーマ (エディター ツール) を取得します。
完全なアプリについては、example/booking および example/crm を参照してください。
コマンド
目的
初期化[ディレクトリ]
新しいプロジェクトの足場を築きます。
上/下/ステータス
ローカル Postgres を管理します (自動ポート)。
コンパイルする
仕様を検証 → schema.sql + データディクショナリをレンダリングします。
移行 [--ドライラン] [--確認]
仕様とライブ DB を比較し、適用します (Alembic ベース。破壊的な変更には --confirm が必要です)。
ドキュメント

AGENTS.md + データ ディクショナリを再生成します。
検査 [エンティティ] / 説明 <エンティティ>
粒度の細かい内省。
作成/取得/一覧表示/更新/削除
エンティティごとの自動 CRUD。
do <op> --json
型指定されたカスタム操作を実行します。
クエリ --sql
読み取り専用 SQL / ビュー → JSON。
import <ファイル> --entity
マッピング + 重複除去を使用して CSV/JSON/Excel を取り込みます。
監査
変更ログを表示します。
仕組み
app.dbm.yaml ──► コンパイラ ──► Postgres DDL (テーブル、FK、チェック、除外、
(仕様) │ enum 型、RLS、updated_at + 監査トリガー)
§──► SQLAlchemy MetaData ──► Alembic の移行 (ドライラン + 確認)
§──► Pydantic モデル ──► 入力検証
└──► AGENTS.md + データ辞書 ──► エージェントはこれらを読み取ります
コンパイラーは、宣言的な diff エンジン (テーブル/列) を手動でロールしません。
変更は Alembic を使用してライブ DB と比較され、レビュー ループ内で維持されます。
移行 --dry-run 。 dbmachine 管理オブジェクト (監査/アクション テーブル、トリガー、
ビュー、RLS) は、移行ごとに冪等に適用されます。
v1 脅威モデルは、独自のエージェントを独自に実行する信頼できるオペレーターです。
機械。 RLS は、オペレーターとエージェントの役割を構造的に正確にコンパイルします
将来的にホストされるマルチテナント API との上位互換性があります。
ローカルで認証されたエージェントに対するセキュリティ境界 (常に
エスカレートします）。実際の保証は、DB 制約、監査ログ、および
バージョン管理における信頼できる情報源としての仕様。
uv venv && uv pip install -e " .[dev] "
uv pip install pgserver # 統合スイート用の埋め込み Postgres
pytest # ユニット + 統合; Postgres テストは利用できない場合は自動スキップします
btree_gist 拡張機能 (除外制約の不変条件) を必要とするテスト
ストック postgres:16 に対して CI で実行します。埋め込みテスト Postgres は contrib を省略します
そしてそれらのテストはローカルでスキップされます。
もしそうなら

エージェントは仕様を作成できますが、最初に保証を変更するのを妨げるものは何でしょうか?
スキーマ変更経路は、信頼ベースであるだけでなく、構造的にエージェントからアクセスできません。
仕様 ( app.dbm.yaml ) はデータベースではなくバージョン管理に存在します。移行コマンド (仕様変更をライブ スキーマに変換する唯一のもの) はオペレーター専用であり、エージェント ロールには公開されません。そのため、エージェントが YAML ファイルを変更したとしても、人間が dbmachine merge を実行するまでデータベースには何も変化しません。
さらに重要なのは、重要な制約 (CHECK、UNIQUE、FK、GiST 除外制約) がアプリケーション層ではなく Postgres 内に存在することです。エージェントは、SELECT/INSERT/UPDATE 権限を持つ制限付き dbm_agent データベース ロールを介して接続しますが、DDL 権限はありません。物理的に ALTER TABLE を発行したり、制約を削除したりすることはできません。エージェントが何を試行しても、データベースはそれを拒否します。
この分離は、オペレータとエージェントを異なる信頼層として区別し、以下によって強制されます。
移行はオペレーターのみです - エージェントは仕様ファイルを編集できますが、移行を実行してライブ データベースに変更を適用できるのはオペレーターのみです
データベース ロールの権限 - エージェント ロールには DDL 権限がありません。
Postgres レベルの構造制約 - GiST 除外制約をオーバーライドするプロンプト エンジニアリングはありません
保証はデータベース内に存在します。エージェントはそれらを変更するアクセス権をまったく持っていないため、最初にそれらを破壊することはできません。
スキーマ変更を行うための UX は何ですか?エージェントが PR を送信してから、オペレーターが移行を実行しますか?
基本的にはそのとおりですが、予想よりも 1 ステップ簡単です。
エージェントは、単なるファイルである app.dbm.yaml を直接編集できます。実行できないのは、 dbmachine merge の実​​行です。これは、仕様変更をライブ データベースに適用するコマンドです。それが門です。
エージェントは仕様を編集します (新しいフィールド、エンティティ、制約を追加します)

int など)
エージェントはオペレーターに次のように伝えます。「仕様を更新しました。dbmachine 移行を実行してください。」
オペレーターは dbmachine merge --dry-run を実行してプレビューし、実際に実行します
個別の「移行の作成」ステップはありません。Alembic は自動的にライブ データベースと仕様を比較し、SQL を生成します。破壊的な変更 (列の削除、テーブル名の変更) には、明示的な --confirm フラグが必要です。これにより、暗黙的にデータ損失が発生することはありません。
PR ベースのワークフローは、あらゆる本番環境のセットアップにとって自然なパターンです。エージェントが PR として仕様変更を提案し、オペレーターがレビューしてマージし、その後 mitigate を実行します。ただし、これはプロセスの規則であり、dbmachine が強制するものではありません。ローカル/信頼されたセットアップでは、エージェントはその場でファイルを編集することができます。
Postgres データベースを自己記述型のエージェント操作可能なアプリケーション バックエンドに変えます。
Apache-2.0ライセンス
貢献する
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

Turn a Postgres database into a self-describing, agent-operable application backend. - kenm47/dbmachine

GitHub - kenm47/dbmachine: Turn a Postgres database into a self-describing, agent-operable application backend. · GitHub
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
kenm47
/
dbmachine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github/ workflows .github/ workflows docs docs examples examples src/ dbmachine src/ dbmachine tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SPEC.md SPEC.md pyproject.toml pyproject.toml View all files Repository files navigation
Turn a Postgres database into a self-describing, agent-operable application backend.
Most business software is a database + a thin backend + a frontend. When the
frontend is a chat agent (like Claude Code), the missing piece is a standard,
safe contract between Postgres and the agent . dbmachine is that contract: a
coding agent declares an app in a YAML spec, dbmachine compiles it into a real
Postgres backend (schema, constraints, RLS, typed operations, audit log), and
generates instructions the agent reads to drive it through one self-documenting
CLI.
The guarantees live in the database, not the prompt. The agent cannot corrupt
your data even when it's wrong — the DB rejects illegal states.
🐘 Real Postgres , fully local (Docker), fully open source. No Supabase, no
third-party accounts.
🤖 Agent-native : rich descriptions everywhere, granular introspection, and
a CLI that emits machine-readable JSON for every command.
🔒 Structural safety : required fields, enums, uniqueness, check &
exclusion constraints (e.g. no double-booking), row-level security, and a
full audit log — all compiled from the spec.
🧱 Clean core : the CLI is a thin adapter over a structured-return Python
core, so a future MCP server is a ~100-line drop-in.
Status: v1. Single trusted operator, CLI-only. HTTP API, MCP server, and
multi-tenant end-user auth are deliberately deferred (see SPEC.md ).
🤖 Building an app with an AI agent? Point it at
docs/AGENT_GUIDE.md . That single, self-contained
document teaches a zero-context agent how to author the spec and stand up a CRM
or booking app end to end — the spec reference, the full CLI, custom
operations, migrations, error codes, and copy-pasteable recipes.
dbmachine is a project-local dependency managed by uv ,
not a global tool. You bootstrap once with uvx , then everything runs in the
project's own environment — so custom operations can freely import stripe /
import pandas ( uv add ... ).
uvx dbmachine init my_crm # scaffold a project (bootstrap only)
cd my_crm
uv sync # set up the project environment
dbmachine up # start local Postgres (auto-detects a free port)
dbmachine migrate # compile the spec → create the schema
dbmachine inspect # see what you've got
dbmachine up auto-detects a free host port (dodging the classic 5432 collision
with a Homebrew Postgres) and persists it to .dbmachine/config.json .
"Claude, build me a CRM" — the walkthrough
The agent authors app.dbm.yaml , then operates the app entirely through the CLI.
Here is the exact contract, using the bundled CRM example ( examples/crm ):
# 1. Compile + migrate
dbmachine migrate
# {"ok": true, "data": {"applied": true, "fresh": true, ...}}
# 2. Learn the schema (pull only what you need)
dbmachine inspect
dbmachine describe deal
# 3. Create data — validated before it hits the DB
dbmachine create company --json ' {"name": "Globex"} '
# {"ok": true, "data": {"id": "…", "name": "Globex", ...}}
dbmachine create deal --json ' {"title": "Big", "amount": 5000, "stage": "proposal", "company_id": "…"} '
# 4. Run a typed custom operation (Python, transactional, with side-effect stubs)
dbmachine do win_deal --json ' {"deal_id": "…"} '
# {"ok": true, "data": {"result": {"stage": "won", ...}}}
# 5. Read-only analytics via views
dbmachine query --sql ' SELECT * FROM pipeline_by_stage '
# 6. Bulk import (CSV/JSON/Excel) with heuristic mapping + dedup
dbmachine import contacts.csv --entity contact --dedup email
# 7. Full change history
dbmachine audit --entity deal
Every command takes --json input where relevant and emits a
{"ok": true|false, ...} envelope on stdout — reliable for an agent to parse.
A declarative, version-controlled source of truth. Excerpt from the bookings
example, showing the structural no-double-booking guarantee:
entities :
appointment :
description : A booking of one resource for one customer over a time window.
fields :
starts_at : { type: timestamptz, required: true }
ends_at : { type: timestamptz, required: true }
status : { type: enum, enum: appointment_status, default: pending }
relationships :
customer : { belongs_to: customer, required: true }
resource : { belongs_to: resource, required: true }
checks :
- { name: ends_after_starts, expr: "ends_at > starts_at" }
exclusions :
- name : no_double_booking
with :
- { field: resource_id, op: "=" }
- { expr: "tstzrange(starts_at, ends_at)", op: "&&" }
where : " status <> 'cancelled' "
That exclusions block compiles to a Postgres GiST exclusion constraint — the
agent cannot create overlapping bookings, no matter what it tries.
Run dbmachine schema to get the JSON Schema for the spec (editor tooling), and
see examples/bookings and examples/crm for complete apps.
Command
Purpose
init [dir]
Scaffold a new project.
up / down / status
Manage local Postgres (auto-port).
compile
Validate spec → render schema.sql + data dictionary.
migrate [--dry-run] [--confirm]
Diff spec vs live DB and apply (Alembic-backed; destructive changes need --confirm ).
docs
Regenerate AGENTS.md + data dictionary.
inspect [entity] / describe <entity>
Granular introspection.
create / get / list / update / delete
Auto-CRUD per entity.
do <op> --json
Run a typed custom operation.
query --sql
Read-only SQL / views → JSON.
import <file> --entity
Ingest CSV/JSON/Excel with mapping + dedup.
audit
Show the change log.
How it works
app.dbm.yaml ──► Compiler ──► Postgres DDL (tables, FKs, checks, exclusions,
(spec) │ enum types, RLS, updated_at + audit triggers)
├──► SQLAlchemy MetaData ──► Alembic migrations (dry-run + confirm)
├──► Pydantic models ──► input validation
└──► AGENTS.md + data dictionary ──► the agent reads these
The compiler does not hand-roll a declarative diff engine — table/column
changes are diffed against the live DB with Alembic and kept in a review loop via
migrate --dry-run . dbmachine-managed objects (audit/actions tables, triggers,
views, RLS) are applied idempotently on each migrate.
The v1 threat model is a trusted operator running their own agent on their own
machine . RLS compiles operator vs agent roles for structural correctness
and forward-compatibility with a future hosted multi-tenant API — it is not a
security boundary against a locally-credentialed agent (which could always
escalate). The real guarantees are the DB constraints, the audit log, and
spec-as-source-of-truth in version control.
uv venv && uv pip install -e " .[dev] "
uv pip install pgserver # embedded Postgres for the integration suite
pytest # unit + integration; Postgres tests auto-skip if unavailable
Tests requiring the btree_gist extension (the exclusion-constraint invariant)
run in CI against stock postgres:16 ; the embedded test Postgres omits contrib
and those tests skip locally.
If the agent can author the spec, what stops it from changing the guarantees first?
The schema change pathway is structurally inaccessible to the agent — not just trust-based.
The spec ( app.dbm.yaml ) lives in version control, not in the database. The migrate command — the only thing that turns spec changes into live schema — is operator-only and not exposed to the agent role. So even if an agent modified the YAML file, nothing changes in the database until a human runs dbmachine migrate .
More importantly, the constraints that matter (CHECK, UNIQUE, FK, GiST exclusion constraints) live inside Postgres, not in the application layer. The agent connects via a restricted dbm_agent database role with SELECT/INSERT/UPDATE grants but no DDL permissions. It physically cannot issue ALTER TABLE or drop a constraint — the database rejects it regardless of what the agent tries.
The separation is operator vs. agent as distinct trust tiers, enforced by:
migrate is operator-only — the agent can edit the spec file, but only the operator can run migrate to apply changes to the live database
Database role permissions — the agent role has no DDL rights
Postgres-level structural constraints — no prompt engineering overrides a GiST exclusion constraint
The guarantees live in the database. The agent can't subvert them first because it never had access to change them at all.
What's the UX for making a schema change? Does the agent submit a PR, then the operator runs migrate?
Essentially yes, but one step simpler than you might expect.
The agent can edit app.dbm.yaml directly — it's just a file. What it can't do is run dbmachine migrate , which is the command that applies spec changes to the live database. That's the gate.
Agent edits the spec (adds the new field, entity, constraint, etc.)
Agent tells the operator: "I've updated the spec — please run dbmachine migrate "
Operator runs dbmachine migrate --dry-run to preview, then runs it for real
There's no separate "create the migration" step — Alembic automatically diffs the spec against the live database and generates the SQL. Destructive changes (dropping a column, renaming a table) require an explicit --confirm flag so nothing data-loss happens silently.
A PR-based workflow is the natural pattern for any production setup: agent proposes the spec change as a PR, operator reviews and merges, then runs migrate . But that's a process convention, not something dbmachine enforces — in a local/trusted setup the agent can just edit the file in place.
Turn a Postgres database into a self-describing, agent-operable application backend.
Apache-2.0 license
Contributing
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
