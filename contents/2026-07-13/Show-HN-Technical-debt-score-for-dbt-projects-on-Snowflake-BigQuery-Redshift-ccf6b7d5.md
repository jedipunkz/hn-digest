---
source: "https://github.com/obrienciaran/dbt-debt"
hn_url: "https://news.ycombinator.com/item?id=48895569"
title: "Show HN: Technical-debt score for dbt projects on Snowflake/BigQuery/Redshift"
article_title: "GitHub - obrienciaran/dbt-debt: A technical-debt scorecard for dbt projects on Snowflake, BigQuery, and Redshift. Finds unused models, columns, orphaned tables, and undeclared sources from query history. Reports only, never modifies your data. · GitHub"
author: "kiki_kuuki"
captured_at: "2026-07-13T17:07:24Z"
capture_tool: "hn-digest"
hn_id: 48895569
score: 1
comments: 0
posted_at: "2026-07-13T17:01:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Technical-debt score for dbt projects on Snowflake/BigQuery/Redshift

- HN: [48895569](https://news.ycombinator.com/item?id=48895569)
- Source: [github.com](https://github.com/obrienciaran/dbt-debt)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T17:01:18Z

## Translation

タイトル: HN を表示: Snowflake/BigQuery/Redshift 上の dbt プロジェクトの技術的負債スコア
記事のタイトル: GitHub - obrienciaran/dbt-debt: Snowflake、BigQuery、Redshift 上の dbt プロジェクトの技術的負債スコアカード。未使用のモデル、列、孤立したテーブル、および宣言されていないソースをクエリ履歴から検索します。レポートのみを提供し、データを変更することはありません。 · GitHub
説明: Snowflake、BigQuery、Redshift 上の dbt プロジェクトの技術的負債スコアカード。未使用のモデル、列、孤立したテーブル、および宣言されていないソースをクエリ履歴から検索します。レポートのみを提供し、データを変更することはありません。 - obrienciaran/dbt-debt

記事本文:
GitHub - obrienciaran/dbt-debt: Snowflake、BigQuery、Redshift 上の dbt プロジェクトの技術的負債スコアカード。未使用のモデル、列、孤立したテーブル、および宣言されていないソースをクエリ履歴から検索します。レポートのみを提供し、データを変更することはありません。 · GitHub
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
{

{ メッセージ }}
オブリエンシアラン
/
dbt-借金
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
47 コミット 47 コミット dbt_debt dbt_debt テスト テスト .gitignore .gitignore DESIGN.md DESIGN.md LICENSE LICENSE README.md README.md USAGE.md USAGE.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
dbt-debt は、BigQuery、Snowflake、または Redshift 上の dbt プロジェクト内のデッド ウェイトを見つけます。
もう誰も使用していないモデルと列、ほとんど使用されていないもの、削除しても安全なものはどれですか。
どのテーブルがウェアハウス内に存在し、その背後に dbt モデルが存在しないのか。
これは、dbt プロジェクトをウェアハウスがすでに知っている 2 つのことと比較することで機能します。
実行されたすべてのクエリのログ
実際に存在するテーブルのリスト
アカウントを作成する必要も、ログインする必要もありません。 dbt run がマシン上で動作する場合、dbt-debt scan も動作します。
適切なウェアハウス権限を持っている場合に限ります。
👉 dbt-debt のみがレポートします。何も編集したり削除したりすることはありません。
dbt_debt_example.mp4
モデル:
✓ 213 がアクティブ
✗ 未使用 17 個 (種子 2 個を含む)
~ 5 ほとんど使用されない (最大 5 クエリ。「未使用」にはカウントされない)
? 1 件は新しすぎて判断できません (最近初めて確認されました。「未使用」にはカウントされません)
列:
✓ 4382 がアクティブ
✗ 3 未使用
出典:
✗ 1 宣言されたソースはプロジェクトに何も読み込まれません
！ 1 つのソースが古い (30 日以上新しいデータがない)
ドキュメントのドリフト:
！文書化された 2 つの列がテーブルに存在しなくなりました
孤児:
✗ dbt モデルのない管理データセット内の 4 つのテーブル
！ 2 つのソースが見つかりましたが、マニフェストで宣言されていません
対象範囲:
- テスト: 230 モデル中 121 モデルに少なくとも 1 つのテストがあります (53%)
- モデルドキュメント: 230 モデル中 88 モデル (38%) に説明があります。
- コラム ドキュメント: 4385 コラム中 1930 コラムに説明があります (44%、カタログ コラム)

s)
潜在的な節約効果:
- 3列取り外し可能
- 2つのテストを取り外し可能
！ 1 つの露出は未使用のモデル (おそらく死んでいる) によってのみ供給されます
- レガシー_kpi_ダッシュボード
！ 1 件の露出が影響を受けました (削除する前に確認してください)
- 毎週の収益ダッシュボード
！ 1 セマンティック層コンシューマーは未使用のモデルを読み取ります (これらのモデルが削除されると壊れます)。
- total_revenue (メトリック) — Legacy_revenue に基づいて構築 (未使用)
- 5.0 GB の再利用可能なストレージ
未使用列 3 つのうち上位 3 つ（テーブルのバイト数でランク付けされています。BigQuery には列ごとのサイズがありません）:
1. dim_customer.old_marketing_score
2. fct_orders.legacy_discount_code
3. mart_sales.temp_margin_calc
📊 数字の意味
モデルは .sql ファイルの 1 つです。列は、モデルが構築するテーブル内の 1 つのフィールドです。
ルックバック ウィンドウは、ウェアハウスのクエリ ログをどれだけ遡って読み取るかです。デフォルトでは 180 日です。
これは、BigQuery が保持する最も多くのデータでもあります。 Snowflake は 1 年保持されるため、--lookback-days は増加する可能性があります
そこの365まで。 Redshift の SYS クエリ履歴ビューははるかに少なく保存されます (AWS は正確なビューを残します)
保持期間は明記されていない。古い STL ビューは 7 日間保持されるため、Redshift では有効期間が長くなります。
アカウントが実際に保持する履歴はどれだけ多くても、「未使用」とは「次の期間で未使用」を意味します。
その歴史」。
アクティブ/未使用モデル。ウィンドウ内で何もクエリされなかった場合、モデルは未使用です。
そこから何かが構築されることを問うものは何もありませんでした。それ以外はすべてアクティブです。シードとスナップショットは、
同じ方法でチェックし、 (seed) / (snapshot) のタグを付けました。
使用頻度の低いモデル。クエリは実行されますが、ウィンドウ全体で最大 5 回です ( --rare-threshold ;
0 はバンドをオフにします)。これらは依然として使用済みとしてカウントされ、リムーバブルまたは再利用可能なものにフィードされることはありません。
数字。それぞれのクエリ数、最後にクエリされた日付、サイズ、およびそれらのバイト数がリストされます。
スキャンされたクエリ、最も多くのクエリが最初にスキャンされます。わずかなクエリ数で大量のバイトがスキャンされると、コストが高くなります
そしてほとんど使用されていない、st

非推奨となる最も長いケース。
判断するには新しすぎる。 7 日以内にクエリ ログで最初に確認されたモデル
( --min-age-days ) はクエリされる機会が十分にないため、代わりに個別にリストされています。
未使用品といいます。ガードは、dbt が構築するもの (モデル、シード、スナップショット) のみをカバーします。の
「dbt にはこのテーブルのレコードがありません」がその日から true であるため、孤児はその年齢に関係なく報告されます。
1つ。 Snowflake では、初見の日付がまったくないモデルも同じ方法で保存されます (「
最初に見た日付、おそらく新しいテーブル")、日付の背後にあるメタデータがあるため
( ACCOUNT_USAGE.TABLES ) には約 90 分の遅れがあります。
アクティブ/未使用の列。クエリがそれを読み取らず、何も読み込まれない場合、その列は未使用です。
それをもとに作られたコラム。クエリがどの列を読み取ったかを確認するために、dbt-debt はすべての列の SQL テキストを解析します。
ログ内のクエリ。一部の SQL は解析に失敗して除外されるため、レポートには、SQL のどれだけが
読み取れるクエリ テキスト (「クエリ テキストの 96% に基づく列の判定」)、判定はそれに基づいて行われます。
解析されていない残りの部分には、スキャンで確認されなかった列の読み取りが含まれている可能性があります。モデル
判定はクエリ ログ自体のメタデータから得られ、解析には決して依存しません。
削除できる列。未使用の列はプロジェクト内の何も依存していません (データがありません)
テスト、契約なし）。 「未使用」は次のようなものであるため、これらは提案であり、自動削除ではありません。
クエリ ログ。すべてを確認できるわけではありません (「使用量としてカウントされるもの」を参照)。未使用の列
まだ依存関係がある場合は別途リストされ、ここではカウントされません。
削除できるテスト。未使用のモデルまたは列に接続されたデータ テスト。モデルを削除する
または列とそのテストがそれに伴います。
露出が影響を受けます。エクスポージャは、下流の消費者 (ダッシュボード、レポート) です。
チームは dbt プロジェクトに書き込みました。未使用のモデルにフィードするモデルには「影響を受ける」というフラグが立てられるため、
チェ

取り外す前に確認してください。モデルがすべてアクティブなエクスポージャーはリストされません。
おそらく死んでいる可能性がある露出。露出で読み取られるすべてのモデルが未使用の場合、ダッシュボードは
それ自体がおそらく死んでいる。これらは退職候補者とは別に名前でリストされている。
上記の影響を受けたエクスポージャには、まだ少なくとも 1 つのアクティブなモデルが背後にあります。
影響を受けるセマンティック層の消費者。セマンティックモデル、メトリクス、または
保存されたクエリ (メトリクスのチェーンを介した場合でも) も同様にレビュー用のフラグが立てられ、列が追加されます。
セマンティック モデル名は、削除可能としてカウントされることはありません。各消費者は理由とともにリストされます
それは現れます。未使用のモデルに直接構築されるか、リストされている別のモデルから読み取られます。
順に影響を受ける消費者 (「経由」として表示)。 「via」行に沿って全体を歩きます
保存されたクエリから未使用のモデルまでチェーンします。
背後に dbt モデルがないテーブル (孤立)。データセット dbt 内の実際のテーブルまたはビュー
に組み込まれますが、dbt レコードはなく、通常は名前変更または削除されたモデルから残ります。または
手作りで作られています。 dbt が構築するデータセットのみが検索されます。 dbt が読み取るだけのデータセット
from (宣言されたソースが存在し、dbt の外部のローダーによって埋められる場所) は決して検索されません。
そこでは、テーブルのレコードがない dbt が正常であり、生の入力テーブルにはすべてフラグが立てられます。
それぞれが、ユーザーがそれに対して実行した直接クエリ (カウント、最終日付、
スキャンされたバイト数);クエリされた孤児はまだ使用中であり、削除するのは危険であるため、それらが最初に行われます。
モデルが読み取ったテーブルですが、dbt はそれについて知らされませんでした。モデルは決して読まないテーブルを読み取ります
宣言した。それを source() として追加します。モデルの SQL を読み取ることで検出されるため、追加の処理は必要ありません
倉庫の許可。
ソースは宣言されていますが、読み込まれていません。逆のケースでは、sources.yml エントリにモデルがなく、露出があり、
またはセマンティック層の消費

えー参考資料。それぞれのファイル パスと直接クエリがリストされます。
人々が生のテーブル (カウント、最終日付、スキャンされたバイト数) に対して実行したため、死者を見分けることができます。
チームが dbt の外部で読み取るテーブルからの宣言 (エントリの削除) (モデリングを考慮する)
それ）。ソースにアタッチされたデータ テストは読み取りではなく宣言であるため、
ここに登場するソース。このリストはレビューのみを目的としており、未使用モデルの数値を変更することはありません。
古いソース。 30 日以上新しいデータがない宣言された情報源
( --stale-source-days ; 0 はチェックをオフにします)。これは通常、
DBが停止しました。日付はウェアハウスのメタデータから取得されます。 BigQuery には読み取りアクセスが必要です
ソース データセット (それがないと警告が表示されてスキップされます)。 Snowflake には追加の助成金は必要ありません。オン
Snowflake では、DDL の変更 (テーブルのコメントでも) によって日付も移動するため、テーブルが古くなると、
データよりも新鮮に見えることもあります。 Redshift は最終変更されたメタデータをまったく公開しません。
したがって、メモを使用してチェックはスキップされます。
ドキュメントのドリフト。モデルの YAML で宣言された列が、
構築されたテーブル (catalog.json ごと) は削除すべき古いドキュメントです。 dbt ドキュメント生成を再実行する
まずカタログが古い場合。
カバレッジ。少なくとも 1 つのテストを持つモデルの数、テストの数をカバーする 3 つの衛生文
説明があり、列がいくつあるか。列の数値は、実際の列を数えます。
存在する場合はcatalog.json、それ以外の場合はYAMLで宣言されたもの。
パーティショニングやクラスタリングを行わない大きなテーブル（BigQuery のみ）。テーブルまたはインクリメンタル
Partition_by もcluster_by も使用せずに構築された 1 GB 以上のモデルは、からフル スキャンを取得します。
それに関わるすべてのクエリ。違反者 (最大 20 人) が、保存されたサイズと
ウィンドウ全体でスキャンされたユーザー クエリのバイト数。最も多いものが最初にスキャンされるため、

一番上のエントリは
最も節約できるパーティション分割の修正。 BigQuery は Snowflake マイクロパーティションとしてのみ
自動的にソートされ、Redshift 自体がソートと分散を管理します。
メンテナンスが必要な大きなテーブル (Redshift のみ)。 1 GB 以上のテーブル
SVV_TABLE_INFO 行には、未ソートの大きな領域 (20% 以上、VACUUM が必要)、古いプランナーが表示されます
統計 (stats_off 10+、ANALYZE が必要)、または重いスライス スキュー (4x+、分析が必要)
配布キーのレビュー)。保存されたサイズとユーザー クエリがスキャンされたバイト数がリストされています。
最初にスキャンされました。自動バキュームと分析では通常、このリストは空のままであり、空のリストは
健康な状態。 BigQuery と Snowflake はストレージ レイアウト自体を維持します。
上位の未使用モデル/列。まずは最大の勝利。未使用のテーブル全体がストレージを示します
あなたは取り戻すでしょう。 Snowflake と Redshift では、サイズはウェアハウスからライブで取得されます (
dbt ドキュメントの生成が必要です)、Snowflake にはタイムトラベルとフェールセーフのコピーが含まれています
まだ請求されています。列のサイズを変更することはできません (列ごとにウェアハウス レポートのストレージがありません)。
テーブルのサイズによってランク付けされます。
まだ PyPI に対応していないため、このリポジトリのコピーからインストールします (Python 3.10 以降が必要)。
git clone <このリポジトリ>
cd dbt-借金
pip インストール 。
(Snowflake および Redshift インストール エクストラの場合、接続

[切り捨てられた]

## Original Extract

A technical-debt scorecard for dbt projects on Snowflake, BigQuery, and Redshift. Finds unused models, columns, orphaned tables, and undeclared sources from query history. Reports only, never modifies your data. - obrienciaran/dbt-debt

GitHub - obrienciaran/dbt-debt: A technical-debt scorecard for dbt projects on Snowflake, BigQuery, and Redshift. Finds unused models, columns, orphaned tables, and undeclared sources from query history. Reports only, never modifies your data. · GitHub
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
obrienciaran
/
dbt-debt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
47 Commits 47 Commits dbt_debt dbt_debt tests tests .gitignore .gitignore DESIGN.md DESIGN.md LICENSE LICENSE README.md README.md USAGE.md USAGE.md pyproject.toml pyproject.toml View all files Repository files navigation
dbt-debt finds the dead weight in a dbt project on BigQuery, Snowflake, or Redshift.
Which models and columns nobody uses anymore, which are barely used, which are safe to remove, and
which tables exist in your warehouse with no dbt model behind them.
It works by comparing your dbt project against two things the warehouse already knows:
a log of every query that has run
a list of the tables that actually exist
There's no account to make and nothing to log into. If dbt run works on your machine, dbt-debt scan works too,
provided you have the right warehouse permissions.
👉 dbt-debt only reports. It never edits or deletes anything.
dbt_debt_example.mp4
Models:
✓ 213 active
✗ 17 unused (incl. 2 seeds)
~ 5 rarely used (at most 5 queries; not counted in 'unused')
? 1 too new to judge (first seen recently; not counted in 'unused')
Columns:
✓ 4382 active
✗ 3 unused
Sources:
✗ 1 declared source nothing in the project reads
! 1 source stale (no new data in 30+ days)
Docs drift:
! 2 documented columns no longer exist in the table
Orphans:
✗ 4 tables in managed datasets with no dbt model
! 2 sources found but not declared in the manifest
Coverage:
- tests: 121 of 230 models have at least one test (53%)
- model docs: 88 of 230 models have a description (38%)
- column docs: 1930 of 4385 columns have a description (44%, catalog columns)
Potential savings:
- 3 columns removable
- 2 tests removable
! 1 exposure fed only by unused models (likely dead)
- legacy_kpi_dashboard
! 1 exposure affected (review before removing)
- weekly_revenue_dashboard
! 1 semantic-layer consumer reads unused models (it would break if those models are removed):
- total_revenue (metric) — built on legacy_revenue (unused)
- 5.0 GB reclaimable storage
Top 3 of 3 unused columns (ranked by table bytes; BigQuery has no per-column sizes):
1. dim_customer.old_marketing_score
2. fct_orders.legacy_discount_code
3. mart_sales.temp_margin_calc
📊 What the numbers mean
A model is one of your .sql files. A column is one field in the table that model builds.
The lookback window is how far back we read the warehouse's query log: 180 days by default,
which is also the most BigQuery keeps. Snowflake keeps a year, so --lookback-days can go up
to 365 there. Redshift's SYS query-history views keep much less (AWS leaves the exact
retention unstated; the older STL views keep seven days), so on Redshift the effective window
is however much history the account actually retains, so "unused" there means "unused within
that history".
active / unused models. A model is unused if, in the window, nothing queried it and
nothing queried anything built from it. Everything else is active . Seeds and snapshots are
checked the same way and tagged (seed) / (snapshot) .
rarely used models. Queried, but at most 5 times in the whole window ( --rare-threshold ;
0 turns the band off). These still count as used and never feed the removable or reclaimable
figures. Each is listed with its query count, last-queried date, size, and the bytes those few
queries scanned, most scanned first. A big bytes scanned on a tiny query count means expensive
and barely used, the strongest case for deprecating.
too new to judge. A model first seen in the query log fewer than 7 days ago
( --min-age-days ) hasn't had a fair chance to be queried, so it's listed separately instead of
being called unused. The guard only covers what dbt builds (models, seeds, snapshots); an
orphan is reported whatever its age, because "dbt has no record of this table" is true from day
one. On Snowflake, a model with no first-seen date at all is set aside the same way ("missing a
first-seen date, likely a new table"), because the metadata behind the date
( ACCOUNT_USAGE.TABLES ) lags about 90 minutes.
active / unused columns. A column is unused if no query read it and nothing read a
column built from it. To see which columns a query read, dbt-debt parses the SQL text of every
query in the log. Some SQL fails to parse and is left out, so the report says how much of the
query text it could read ("column verdicts based on 96% of query text"), the verdicts rest on
that share, and the unparsed remainder could contain column reads the scan did not see. Model
verdicts come from the query log's own metadata and never depend on parsing.
columns you could remove. Unused columns nothing in the project still depends on (no data
test, no contract). These are suggestions, not an automatic delete, since "unused" comes from
the query log, which can't see everything (see What counts as usage ). An unused column that
still has a dependency is listed separately and not counted here.
tests you could remove. Data tests attached to an unused model or column. Remove the model
or column and its tests go with it.
exposures affected. An exposure is a downstream consumer (a dashboard, a report) your
team has written into the dbt project. An unused model feeding one is flagged "affected" so you
check before removing it. Exposures whose models are all active aren't listed.
exposures that are likely dead. When every model an exposure reads is unused, the dashboard
itself is probably dead. These are listed by name as candidates to retire, separately from the
affected exposures above, which still have at least one active model behind them.
semantic-layer consumers affected. An unused model feeding a semantic model, metric, or
saved query (even through a chain of metrics) is flagged for review the same way, and a column
a semantic model names is never counted as removable. Each consumer is listed with the reason
it appears. Either it is built directly on an unused model, or it reads from another listed
consumer that is affected in turn (shown as "via"). Following the "via" lines walks the whole
chain from the saved query down to the unused model.
tables with no dbt model behind them (orphans). A real table or view in a dataset dbt
builds into, but with no dbt record, usually left over from a renamed or deleted model, or
made by hand. Only the datasets dbt builds into are searched. The datasets dbt merely reads
from (where declared sources live, filled by loaders outside dbt) are never searched, because
dbt having no record of a table is normal there and every raw input table would be flagged.
Each is listed with any direct queries people ran against it (count, last date,
bytes scanned); a queried orphan is still in use and dangerous to drop, so those come first.
tables your models read but dbt was never told about. A model reads a table you never
declared; add it as a source() . Found by reading the model's SQL, so it needs no extra
warehouse permission.
sources declared but never read. The reverse case, a sources.yml entry no model, exposure,
or semantic-layer consumer references. Each is listed with its file path and any direct queries
people ran against the raw table (count, last date, bytes scanned), so you can tell a dead
declaration (delete the entry) from a table your team reads outside dbt (consider modelling
it). A data test attached to the source is a declaration, not a read, so it does not stop the
source appearing here. This list is for review only and never changes the unused-model figures.
stale sources. A declared source with no new data for more than 30 days
( --stale-source-days ; 0 turns the check off), which usually means the loader upstream of
dbt has stopped. The date comes from warehouse metadata. BigQuery needs read access to the
source datasets (skipped with a warning without it); Snowflake needs no extra grant. On
Snowflake the date also moves on DDL changes (even a table comment), so a stale table can
occasionally look fresher than its data. Redshift exposes no last-modified metadata at all,
so the check is skipped there with a note.
documentation drift. A column declared in a model's YAML that no longer exists in the
built table (per catalog.json ) is stale documentation to delete. Rerun dbt docs generate
first if the catalog is old.
coverage. Three hygiene sentences covering how many models have at least one test, how many
have a description, and how many columns do. The column figure counts the real columns from
catalog.json when present, else the ones declared in YAML.
large tables without partitioning or clustering (BigQuery only). A table or incremental
model of 1 GB or more built with neither partition_by nor cluster_by gets a full scan from
every query that touches it. The offenders (up to 20) are listed with stored size and the
bytes user queries scanned over the window, most scanned first, so the top entry is the
partitioning fix that saves the most. BigQuery only as Snowflake micro-partitions
automatically, and Redshift manages sort and distribution itself.
large tables needing maintenance (Redshift only). A table of 1 GB or more whose
SVV_TABLE_INFO row shows a big unsorted region (20%+, needs VACUUM), stale planner
statistics ( stats_off 10+, needs ANALYZE), or heavy slice skew (4x+, needs a
distribution-key review). Listed with stored size and the bytes user queries scanned, most
scanned first. Automatic vacuum and analyze usually keep this list empty, and an empty list is
the healthy state. BigQuery and Snowflake maintain storage layout themselves.
top unused models / columns. Biggest win first. A whole unused table shows the storage
you'd reclaim; on Snowflake and Redshift the sizes come live from the warehouse (no
dbt docs generate needed), and Snowflake's include the time-travel and fail-safe copies
still billed for it. Columns can't be sized (no warehouse reports storage per column), so
they rank by their table's size.
Not on PyPI yet, so install from a copy of this repo (needs Python 3.10+):
git clone <this repo>
cd dbt-debt
pip install .
(For the Snowflake and Redshift install extras, connection s

[truncated]
