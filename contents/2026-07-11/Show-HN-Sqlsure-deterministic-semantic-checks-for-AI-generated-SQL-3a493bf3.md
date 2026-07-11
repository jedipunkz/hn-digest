---
source: "https://github.com/sqlsure/sqlsure"
hn_url: "https://news.ycombinator.com/item?id=48875342"
title: "Show HN: Sqlsure – deterministic semantic checks for AI-generated SQL"
article_title: "GitHub - sqlsure/sqlsure: Semantic inspector for SQL — catches fan-out double-counting, additivity violations, wrong join keys, and policy breaches before the query runs. Found real bugs in the BIRD/Spider text-to-SQL benchmarks. · GitHub"
author: "tejusarora"
captured_at: "2026-07-11T20:49:37Z"
capture_tool: "hn-digest"
hn_id: 48875342
score: 2
comments: 0
posted_at: "2026-07-11T20:03:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sqlsure – deterministic semantic checks for AI-generated SQL

- HN: [48875342](https://news.ycombinator.com/item?id=48875342)
- Source: [github.com](https://github.com/sqlsure/sqlsure)
- Score: 2
- Comments: 0
- Posted: 2026-07-11T20:03:42Z

## Translation

タイトル: Show HN: Sqlsure – AI 生成 SQL の決定論的セマンティック チェック
記事のタイトル: GitHub - sqlsure/sqlsure: SQL のセマンティック インスペクター — クエリの実行前に、ファンアウトの二重カウント、加法性違反、間違った結合キー、およびポリシー違反を検出します。 BIRD/Spider text-to-SQL ベンチマークで実際のバグが見つかりました。 · GitHub
説明: SQL のセマンティック インスペクター — クエリの実行前に、ファンアウトの二重カウント、加算性違反、間違った結合キー、およびポリシー違反を検出します。 BIRD/Spider text-to-SQL ベンチマークで実際のバグが見つかりました。 - sqlsure/sqlsure

記事本文:
GitHub - sqlsure/sqlsure: SQL のセマンティック インスペクター — クエリの実行前に、ファンアウトの二重カウント、加算性違反、間違った結合キー、およびポリシー違反を検出します。 BIRD/Spider text-to-SQL ベンチマークで実際のバグが見つかりました。 · GitHub
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
シュルシュア
/
シュルシュア
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット .claude .claude .github/ workflows .github/ workflows ベンチマーク ベンチマーク ドキュメント ドキュメント統合 統合 sqlsure sqlsure テスト テスト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md check.py check.py model.example.json model.example.json pyproject.toml pyproject.toml要件.txt要件.txtサーバー.jsonサーバー.jsonすべてのファイルを表示リポジトリ ファイルのナビゲーション
AI が SQL を作成します。 sqlsure はそれが正しいことを確認します。
クエリは完全に有効で、エラーなしで実行され、数値を返すことができます。
それは黙って間違っています — 収益は結合によって二重にカウントされ、平均です
合計すると、患者識別情報が暴露されます。データベースはこれを捉えません。
リンターはこれを理解できません。独自の SQL をレビューしている LLM はこれを認識しません。
sqlsure は、クエリが実行される前に、0.1 ミリ秒以内に決定的に実行します。
約束ではなく証拠: 私たちは 2 人の黄金の答えをめぐって論争を繰り広げました
すべての Text-to-SQL モデルが評価されるベンチマーク。 2,568 の専門家が執筆
クエリ、45 のフラグ、誤報ゼロ - BIRD dev Gold Answer を含む
それは正確な値の 8 倍間違っていることが証明されています
バグ クラス sqlsure ターゲット、およびスキーマの欠陥
現在上流にファイルされています。
sqlsure チームがすでに宣言した事実に基づいて SQL を判断します — dbt unique
テストは粒状になり、関係テストは結合カーディナリティになり、一行になります
メタタグは、合計しても安全なものをマークします。新しい言語を学ぶ必要もなければ、モデルを学ぶ必要もありません
手作業でメンテナンスします。ルールは辞書検索であり、LLM 呼び出しではありません。同じ入力、
毎回同じ判決、

オフライン。
すべての拒否には機械による修正が含まれるため、AI エージェントは
自己修復: ドラフト→チェック→修正→チェック→実行。私たちのベンチマークでは、
修正をそのまま適用すると、10/10 回合格するクエリが生成されました。
pip インストール sqlsure
sqlsure import SemanticModel から、チェックしてください
Violations = check ( sql , model ) # [] は意味的に安全であることを意味します
または、30 秒のデモを複製して実行します。
python check.py # 5 つの間違ったクエリが拒否され、1 つは承認されました — 修正あり
python -m sqlsure.scan path/to/dbt-repo --report report.md # dbt リポジトリを監査します
3 つのドア、1 つのエンジン
1. CI ゲート — PR がダブルカウントした場合にマージをブロックします。
python -m sqlsure.cli --model model.json query.sql # 違反時に終了 1
2. MCP サーバー — AI エージェントは、以下を実行する前に検査に合格する必要があります。
クロード mcp add sqlsure -- python -m sqlsure.mcp_server --model /abs/path/model.json
ツールのリファレンスとエージェント ループ パターンについては、docs/MCP.md を参照してください。
3. ライブラリ — text-to-SQL 製品またはエージェント内に check() を埋め込む
フレームワーク。ドロップイン SemanticGate ラップ
Vanna/WrenAI スタイルのジェネレーター。ある
セマンティック評価メトリック スコア NL2SQL 出力
実行の正確さは盲目です。
ルール
重大度
キャッチ
ファンアウト
エラー
1 対多結合後の加算メジャーの SUM/COUNT
キャズム
エラー
2 つ以上のファンアウト結合が相互に乗算される
相加性
エラー
非加算メジャーの SUM (レート、平均)
SEMI_ADDITIVE
エラー
スナップショットのディメンション全体で合計された残高/国勢調査
JOIN_KEY
エラー
宣言された関係に一致しない列の結合
CROSS_JOIN
エラー
述語なしで結合する
WEIGHTED_AVG
警告
AVG はファンアウトによってサイレントに再重み付けされます
UNDECLARED_JOIN
警告
関係が宣言されていない結合 (検証不可能 ≠ 安全)
センシティブ_カラム
政策
クエリ出力で公開される PHI/PII 列
sqlsure が何かを検証できない場合、「検証できません」と表示されます。
いいよ。」正直

不確実性が特徴です。
決定的 — 同じ SQL + 同じルールブック = 常に同じ判定。
ルールは辞書検索であり、行ごとに監査可能です
オフライン — ネットワーク通話はゼロです。 SQL がマシンから離れることはありません
データ アクセスなし - クエリ テキストを解析します。データベースに接続しない
テレメトリなし - 何も収集されません ( SECURITY.md )
サプライ チェーン — リリースは PyPI Trusted Publishing 経由でのみ出荷されます
(OIDC) パブリック CI 実行によるタグ付きコミットから。 2つのランタイムdeps
dbt (現在動作します):manifest.json または schema.yml — テスト チーム
すでに書き込まれたセマンティクスが強制可能になり、設定は不要
プレーンな PK/FK 宣言 (今日でも機能 - ベンチマーク監査を強化)
ライブデータベース自体（現在も機能します）：セマンティックレイヤーがまったくありませんか？
sqlsure.introspect はカタログからルールブックを構築します — SQLite
PRAGMA または information_schema PK/FK (postgres/mysql)。内省する
BIRD 独自のデータベース ファイルは、データベースから失われた 2 つの外部キーを回復しました。
ベンチマークの公開スキーマ
( 鳥ベンチ/mini_dev#37 )
sqlsure から。イントロスペクト インポート model_from_sqlite
model = model_from_sqlite ( "app.db" ) # PK -> 粒子、FK -> エッジの結合
手書きの JSON — model.example.json
OSI および WrenAI MDL (動作ローダー
統合/ ): OSI
仕様の公開された例で実証されています。
WrenAI MDL は WrenAI 自身で実証されました
出荷されたサンプルマニフェスト — 主キー → 粒度、関係
joinType + 条件 → エッジの結合、キューブ メジャー → 加法性
Cube、Snowflake Semantic Views — ロードマップ上のアダプター。の
エンジンは 1 つの SemanticModel のみを認識します
16/16 ルール テスト、ペアの再現率 100% / 偽陽性 0%
ベンチマーク ( docs/METRICS.md )
実際の本番リポジトリ (Mattermost の倉庫、Fivetran パッケージ、
dbt のジャッフル ショップ) — docs/TEST-REPORTS.md
Spider + BIRD ゴールド クエリ — ノイズのない外部 au

上と同じ
docs/EVIDENCE.md — それがあなたに何をもたらすか
再実行可能な測定にリンクされたクレーム
docs/ARCHITECTURE.md — 物理的にどのように機能するか、
ELI5 → 神レベル、実際の中間出力あり
docs/FOR-DUMMIES.md — すべてのコンセプトをゼロから
docs/INTEGRATIONS.md — GitHub アクション、コミット前、
MCP、Snowflake UDF / Cortex Agent ツール、クエリ履歴監査
docs/MCP.md — MCP サーバーのドキュメント
COTRIBUTING.md — ルールとローダーの追加
mcp 名: io.github.sqlsure/sqlsure
SQL のセマンティック インスペクター — クエリの実行前に、ファンアウトの二重カウント、加算性違反、間違った結合キー、およびポリシー違反を検出します。 BIRD/Spider text-to-SQL ベンチマークで実際のバグが見つかりました。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
2
sqlsure v0.1.1
最新の
2026 年 7 月 3 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Semantic inspector for SQL — catches fan-out double-counting, additivity violations, wrong join keys, and policy breaches before the query runs. Found real bugs in the BIRD/Spider text-to-SQL benchmarks. - sqlsure/sqlsure

GitHub - sqlsure/sqlsure: Semantic inspector for SQL — catches fan-out double-counting, additivity violations, wrong join keys, and policy breaches before the query runs. Found real bugs in the BIRD/Spider text-to-SQL benchmarks. · GitHub
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
sqlsure
/
sqlsure
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits .claude .claude .github/ workflows .github/ workflows benchmarks benchmarks docs docs integrations integrations sqlsure sqlsure tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md check.py check.py model.example.json model.example.json pyproject.toml pyproject.toml requirements.txt requirements.txt server.json server.json View all files Repository files navigation
AI writes your SQL. sqlsure makes sure it's right.
A query can be perfectly valid, run without error, and return a number
that's silently wrong — revenue double-counted by a join, an average
summed, a patient identifier exposed. Databases don't catch this.
Linters don't catch this. LLMs reviewing their own SQL don't catch this.
sqlsure does — deterministically, in 0.1 ms, before the query runs.
Proof, not promises: we ran sqlsure over the gold answers of the two
benchmarks every text-to-SQL model is graded on. 2,568 expert-written
queries, 45 flags, zero false alarms — including a BIRD dev gold answer
that is provably wrong by 8× from the exact
bug class sqlsure targets, and a schema defect
now filed upstream .
sqlsure judges SQL against facts your team already declared — dbt unique
tests become grain, relationships tests become join cardinality, one-line
meta tags mark what's safe to sum. No new language to learn, no model to
maintain by hand. Rules are dictionary lookups, not LLM calls: same input,
same verdict, every time, offline.
Every rejection carries a machine-actionable fix , so AI agents
self-repair: draft → check → fix → check → execute. In our benchmark,
applying the fix verbatim produced a passing query 10/10 times.
pip install sqlsure
from sqlsure import SemanticModel , check
violations = check ( sql , model ) # [] means semantically safe
Or clone and run the 30-second demo:
python check.py # 5 wrong queries rejected, 1 approved — with fixes
python -m sqlsure.scan path/to/dbt-repo --report report.md # audit any dbt repo
Three doors, one engine
1. CI gate — blocks the merge when a PR double-counts:
python -m sqlsure.cli --model model.json query.sql # exit 1 on violations
2. MCP server — your AI agent must pass inspection before executing:
claude mcp add sqlsure -- python -m sqlsure.mcp_server --model /abs/path/model.json
See docs/MCP.md for tool reference and agent-loop patterns.
3. Library — embed check() inside any text-to-SQL product or agent
framework. A drop-in SemanticGate wraps
Vanna/WrenAI-style generators; a
semantic eval metric scores NL2SQL output
where execution-accuracy is blind.
Rule
Severity
Catches
FANOUT
error
SUM/COUNT of additive measure after one-to-many join
CHASM
error
two+ fan-out joins multiplying each other
ADDITIVITY
error
SUM of a non-additive measure (rates, averages)
SEMI_ADDITIVE
error
balances/censuses summed across their snapshot dimension
JOIN_KEY
error
join on columns matching no declared relationship
CROSS_JOIN
error
join with no predicate
WEIGHTED_AVG
warning
AVG silently re-weighted by fan-out
UNDECLARED_JOIN
warning
join with no declared relationship (unverifiable ≠ safe)
SENSITIVE_COLUMN
policy
PHI/PII column exposed in query output
When sqlsure can't verify something, it says "can't verify" — never "looks
fine." Honest uncertainty is a feature.
Deterministic — same SQL + same rulebook = same verdict, always;
rules are dictionary lookups, auditable line by line
Offline — zero network calls; your SQL never leaves your machine
No data access — parses query text ; never connects to a database
No telemetry — nothing collected, ever ( SECURITY.md )
Supply chain — releases ship exclusively via PyPI Trusted Publishing
(OIDC) from tagged commits with public CI runs; two runtime deps
dbt (works today): manifest.json or schema.yml — the tests teams
already wrote become enforceable semantics, zero config
Plain PK/FK declarations (works today — powered the benchmark audits)
The live database itself (works today): no semantic layer at all?
sqlsure.introspect builds the rulebook from the catalog — SQLite
PRAGMAs or information_schema PK/FK (postgres/mysql). Introspecting
BIRD's own database files recovered 2 foreign keys missing from the
benchmark's published schema
( bird-bench/mini_dev#37 )
from sqlsure . introspect import model_from_sqlite
model = model_from_sqlite ( "app.db" ) # PK -> grain, FK -> join edges
Hand-written JSON — model.example.json
OSI and WrenAI MDL (working loaders in
integrations/ ): OSI
demonstrated on the spec's published examples;
WrenAI MDL demonstrated on WrenAI's own
shipped example manifest — primaryKey → grain, relationship
joinType + condition → join edges, cube measures → additivity
Cube, Snowflake Semantic Views — adapters on the roadmap; the
engine only ever sees one SemanticModel
16/16 rule tests, 100% recall / 0% false positives on the paired
benchmark ( docs/METRICS.md )
Real production repos (Mattermost's warehouse, Fivetran packages,
dbt's jaffle shop) — docs/TEST-REPORTS.md
Spider + BIRD gold queries — the zero-noise external audit above
docs/EVIDENCE.md — what it does for you, every
claim linked to a rerunnable measurement
docs/ARCHITECTURE.md — how it physically works,
ELI5 → god level, with real intermediate outputs
docs/FOR-DUMMIES.md — every concept from zero
docs/INTEGRATIONS.md — GitHub Action, pre-commit,
MCP, Snowflake UDF / Cortex Agent tool, query-history audit
docs/MCP.md — MCP server documentation
CONTRIBUTING.md — adding rules and loaders
mcp-name: io.github.sqlsure/sqlsure
Semantic inspector for SQL — catches fan-out double-counting, additivity violations, wrong join keys, and policy breaches before the query runs. Found real bugs in the BIRD/Spider text-to-SQL benchmarks.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
2
sqlsure v0.1.1
Latest
Jul 3, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
