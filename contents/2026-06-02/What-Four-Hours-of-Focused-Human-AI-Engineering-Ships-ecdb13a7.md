---
source: "https://canonical.agency/essays/what-four-hours-of-human-ai-engineering-actually-ships"
hn_url: "https://news.ycombinator.com/item?id=48359256"
title: "What Four Hours of Focused Human-AI Engineering Ships"
article_title: "What Four Hours of Focused Human-AI Engineering Actually Ships · Canonical"
author: "nickstinemates"
captured_at: "2026-06-02T00:36:26Z"
capture_tool: "hn-digest"
hn_id: 48359256
score: 2
comments: 0
posted_at: "2026-06-01T16:43:31Z"
tags:
  - hacker-news
  - translated
---

# What Four Hours of Focused Human-AI Engineering Ships

- HN: [48359256](https://news.ycombinator.com/item?id=48359256)
- Source: [canonical.agency](https://canonical.agency/essays/what-four-hours-of-human-ai-engineering-actually-ships)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T16:43:31Z

## Translation

タイトル: 4 時間の集中的な人間 AI エンジニアリングがもたらすもの
記事のタイトル: 4 時間の集中的な人間 AI エンジニアリングが実際に出荷するもの · Canonical
説明: 5 月 30 日に、Swamp 用の Databricks 統合パックの v0.1 を公開しました。同じ夕方の終わりまでに、v0.13 になりました。 15 のモデル、すべてのリリースで 100% A、すべてのモデルがエンドツーエンドのスモーク テスト済み。出荷されたもの、学んだこと、そして Databricks を他のものと一緒に使用することが重要である理由。

記事本文:
4 時間の集中的な人間 AI エンジニアリングが実際にもたらすもの · Canonical Canonical 独立系 AI コンサルタントの能力
← エッセイ 4 時間の集中的な人間 AI エンジニアリングが実際にもたらすもの
5 月 30 日に、Swamp 用の Databricks 統合パックの v0.1 を公開しました。同じ夕方の終わりまでに、v0.13 になりました。 15 のモデル、すべてのリリースで 100% A、すべてのモデルがエンドツーエンドのスモーク テスト済み。出荷されたもの、学んだこと、そして Databricks を他のものと一緒に使用することが重要である理由。
請求前の領収書。モデルは15人。 13 枚のリリース。ある晩。
5 月 30 日に、System Initiative の新しい宣言型自動化フレームワークである Swamp 用の Databricks 統合パックの v0.1 を公開しました。同じ夜の終わりには v0.13 になり、ジョブ、DLT パイプライン、SQL ウェアハウス、完全な Unity カタログ ツリー、ワークスペース シークレット、権限、DBSQL クエリ、Git Repos をカバーする 15 のモデルが含まれていました。すべてのリリースは、プラットフォームの品質ルーブリックで 100/A のスコアを獲得しました。すべてのモデルは、実際の Databricks ワークスペースに対してエンドツーエンドのスモーク テストが行​​われました。あるリリースでは実際のバグが出荷されました。 2 つのリリース後に修正を出荷しました。クロードと二人で、約 4 時間の集中作業を要しました。
この投稿では、パックの内容、スタック内の他のものと一緒に Databricks を使用することが重要である理由、およびその過程でエンジニアリングの速度について学んだことについて説明します。
15 個の Swamp モデル。 swamp.club/extensions/@mfbaig35r/databricks に公開され、ソースは github.com/mfbaig35r/swamp-databricks にあります。これらは Databricks データ エンジニアリングの面をカバーしています。
コンピューティングとオーケストレーション: ジョブ (ジョブ API 2.2)、DLT パイプライン、ステートメント実行を備えた SQL ウェアハウス
ワークスペース ストレージ: ノートブックとワークスペース ファイル (2 つの異なる Databricks オブジェクト タイプ、2 つの異なるモデル、異なる upl)

OAD セマンティクス)
Unity カタログ : カタログ、スキーマ、テーブル、ボリューム
アクセス制御: ワークスペース権限 (ジョブ、パイプライン、ウェアハウス、ノートブック、リポジトリ、クエリ、ダッシュボード、アラートなど)、UC 権限 (カタログ、スキーマ、テーブル、ボリューム、関数、外部の場所、ストレージ資格情報の完全付与モデル)
ワークスペース シークレット: Swamp 独自のボールトとは異なるスコープとキー。シークレット値は Swamp のデータ レイヤーに永続化されません。
DBSQL : ジョブモデルの sql_task.query.query_id フィールドとペアになる保存されたクエリ
Git Repos : プル、ブランチ切り替え、スパース チェックアウト、およびフル ライフサイクルを備えたワークスペース接続リポジトリ
各モデルは、実際のワークスペース ユーザーが触れるであろう表面を公開します。ジョブ モデルは、 for_each_task ( Zod z.lazy 経由で再帰的)、 dbt_task 、spark_python_task を含む完全なタスク タイプのスキーマ検証を備えた Jobs API 2.2 をカバーしています。 uc_permissions モデルは、変更スタイルの PATCH セマンティクスの追加/削除を使用して、15 の UC のセキュリティ保護可能なタイプをすべてカバーします。 sql_warehouse モデルには、ウェアハウスのライフサイクルと、ワークフローからクエリを実行するための SQL ステートメント実行 API が含まれています。ライフサイクル管理モデルは、ワークスペース優先リコンシリエーションによる冪等の create_or_update セマンティクスを公開し、削除してから再作成するおよび帯域外のワークスペース削除を正しく処理します。
すべてのリリースは、Swamp の品質ルーブリックで 100/A を獲得しました。すべてのモデルは、パックが次のリリースに移行される前に、実際の Databricks Free ワークスペースに対してエンドツーエンドでスモーク テストが行​​われました。 「スモーク」とは、実際の API パスを意味します。ワークスペースにアップロードされたノートブック、API 経由で作成されたジョブ、トリガーされた実行、終了状態になるまで待機された実行、検証された結果です。型チェックではありません。この規律は、v0.8 の create_or_update で真のセマンティック バグを発見し、v0.11 で修正を提供しました。詳細については以下をご覧ください。
1 つのリファレンス ワークフローがモードに同梱されています

ls: Met Museum API 取り込みパイプライン (examples/api-ingest/met-museum/ )。エンドツーエンドで検証され、資格情報なしで Databricks Free で実行可能。外部 HTTP API を Bronze + Silver Delta テーブルにプルするためのユニバーサル パターンを実装します ( mapInPandas を介したレート制限された Spark ファンアウト、スキーマ ドリフト耐性のための raw-JSON Bronze、SQL を介した型付き Silver)。 Stripe、GitHub、Salesforce、またはその他の API 用にフォークするには 45 ～ 90 分の作業が必要です。ノートブックの上部にある構成ブロックを変更し、Silver SQL を変更します。滞在間のすべて。
配布は行われない部分です。これを書いている時点では、パックにはプルがありません。今日は配布初日です。エンジニアリング成果物は完成しました。詳しくは最後に。
Databricks にはすでに、それ自体を管理する 2 つの方法があります。インフラ ライフサイクル用の Terraform プロバイダーと、Databricks 内パイプライン定義用の Databricks アセット バンドル (DAB) です。どちらも自分の仕事が得意です。どちらも Databricks の外部のものを使用して構成することはなく、ノートブック内で実際に取り込み、変換、トレーニング作業を行う Python の作成にも役立ちません。
このパックが構築されている明白ではない答えは、汎用の取り込みランタイムを同梱しないことです。エージェントに特定のケースのオンデマンドでノートブック コードを作成させ、その後、エージェントが作成したノートブックを Swamp の宣言型グラフを通じて調整します。
パック内のメトロポリタン美術館の例がその証拠です。私は 200 行のブロンズ インジェスト ノートを手書きしたわけではありません。私はクロードにタスクを説明し (Met のオープン カタログをデルタ テーブルにプルし、レート制限を尊重し、エラーをキャプチャしました)、クロードがノートブックを書き、私がそれをレビューし、Databricks Free でスモーク テストし、1 つのバグ (サーバーレス RDD 制限) を修正して、コミットしました。次に、パックの Notebook.upload ステップがコミットされたノートブックをワークスペースにアップロードし、job.run ステップがそのコミットされたノートブックをワークスペースにアップロードします。

それを実行します。エージェントは IDE です。 swamp-databricks はランタイムです。コミットされたノートブックがアルゴリズムです。
それは2つの境界を一度に崩壊させます。
まず、「1 回限りの統合の構築」と「運用統合の管理」の間です。どちらも同じツール (作成者用の Claude Code、オーケストレーション用の swamp-databricks) を介して処理されます。新しいベンダー API やスキーマの変更は、1 回のフレームワーク リリースではなく、1 回のエージェントで完了します。
2 つ目は、Databricks とその他すべての間です。ワークフローは宣言型 Swamp グラフであるため、エージェントが作成した Databricks ノートブックを実行する同じワークフローで、1 つのファイル内の S3、Postgres、Cloudflare、GitHub、Slack、またはその他の Swamp 拡張機能を、スケジュールに従って冪等な再実行で操作することもできます。ノートブックは不可欠なコアです。ワークフローは宣言的な外側のシェルです。エージェントは、命令型コアの作成と書き換えを安価にするものです。 Met ワークフローの最小バージョン:
手順:
- ID : スキーマ
モデル: "@mfbaig35r/databricks/uc_schema"
メソッド: create_or_update
引数: {名前:met_museum、カタログ名:ワークスペース}
- ID : ノートブック
モデル：「@mfbaig35r/databricks/notebook」
方法：アップロード
引数 : { パス : /Shared/met-museum-bronze 、コンテンツ : | ... }
- ID : トリガー
モデル：「@mfbaig35r/databricks/job」
メソッド: 実行
引数 : { job_ref :met-museum-bronze }
- ID : お待ちください
モデル：「@mfbaig35r/databricks/job」
メソッド: wait_run
引数 : { run_id : $ {{steps.trigger.outputs.run_id }} }
- ID : 許可
モデル: "@mfbaig35r/databricks/uc_permissions"
方法：アップデート
引数 : { 変更内容 : [ ... ] }
経済的な主張は具体的です。以前は汎用フレームワークのエンジニアリングに数週間かかっていた統合パックが、タスク固有の生成とレビューに数時間かかるようになります。 Met の例では、「良いデモ API とは何か」から「実行する」までに約 30 分かかりました。

Databricks Free と Bronze および Silver テーブルが設定されています。」 カーソル ページネーションと OAuth を使用した Stripe の取り込みは、プロンプトからパイプラインの実行まで 30 ～ 45 分でほぼ同じになります。
これは 1 つの条件下でのみ機能します。それは、スモーク テストとコード レビューを負荷がかかるものとして扱うエンジニアです。それがなければ、即時の変動によって生産が変動してしまいます。これにより、1 年前には実現できなかった統合速度が得られます。
Swamp は、Adam Jacob が Chef の頃から構築している System Initiative の新しい宣言型自動化フレームワークです。小さくて新しいです。私がこのパックのランタイムとしてこれを選択したのは、次の 3 つの具体的な理由からです。
構成単位としての宣言型モデル グラフ。モデルは相互に構成され、また他のドメインのモデルとも構成されます。ワークフローは、モデル上の型指定された DAG です。この構造特性により、「Databricks + その他すべて」を 1 つのファイルで表現できるようになります。これは、後から必須ツールに追加できる機能ではありません。それはランタイムの形状です。
CLI でのファーストクラスのエージェントの統合。 swamp repo init は、プロジェクト ルールを含む CLAUDE.md を書き込みます。拡張機能の作成スキルは自動的にロードされます。このフレームワークは、AI エージェントが自動化の作成と操作の両方を行うことを想定して設計されており、ツールもそれを反映しています。クロードと一緒にこのパックで作業することは、Terraform や Pulumi ではできなかったような点で、スムーズでした。
ツール内の緊密なフィードバック ループ。 swamp 拡張機能 fmt フォーマットと lint。公開されている 14 要素のルーブリックに対するスワンプ エクステンションの品質スコア。沼地モデルのメソッドを実行すると、エージェントが使用できる構造化された出力が返されます。レジストリには、ヤンクとアンヤンク、リポジトリで検証された状態、および CalVer バージョン管理が組み込まれています。このツールチェーンにより、リリースごとのスモーク テストの規律が野心的なものではなく実用的なものになりました。ループはしょー

変更のたびに実行しても速度が低下しない程度に十分です。
正直な賭け金のライン: 現在、Swamp には小さなコミュニティがあります。採用されるかどうかは証明されていません。構造上の選択は正しいと思いますので、エンジニアリングの時間を賭けに費やすつもりです。
AIと真剣にペアリングすると何が変わるのか
これは私にとって最も重要なセクションです。テーゼ: AI と真剣に連携したときに変わるのは、「AI がコードを書く」ことではありません。それは、エンジニアリングの判断をループに取り入れれば、以前は数週間かかっていた統合パックの構築コストが数時間に下がるということです。
このセッション全体で誰が何をしたかの具体的な内訳:
これを Vaporware ではないものにする規律:
すべてのリリースは、次のリリースが開始される前にエンドツーエンドでスモーク検証を受けました。型チェックや予行演習ではありません。ノートブックがアップロードされ、ジョブが作成され、実行がトリガーされ、実行がターミナルまで待機され、結果が検証されました。
バグは発見され、紙に隠蔽されるのではなく、修正が送られてきました。 v0.8 から v0.11 へのトゥームストーンのバグ修正は、最もクリーンな例です。セッション中に壊れたセマンティックに気づき、それが本物であると判断し、専用リリースする価値があると判断し、修正を出荷し、修正を検証し、間違いを繰り返さないようにメモリ ファイルを更新しました。
セッションは、git 履歴、リリース ノート、レジストリのバージョン ログに保存されます。誰でも監査することができます。 13 のリリース ノート自体がセッション ログです。何が出荷され、何がスモーク テストされ、どのようなバグが捕捉され、何がまだギャップがあるのか​​がわかります。
私自身の仕事について学んだこと:
私がまだ正しい下書きを決めている間に、誰か (何か) が最初の草稿を入力することで、より速く考えることができます。ボトルネックは入力から決定に移ります。これは、エンジニアリングの仕事の感じ方に大きな変化をもたらします。
判断の判断が価値です。 「どのタスク タイプを最初にスキーマ検証するか。」 「その墓石は本物かどうか」

「v0.2 を公開でヤンクするかどうか」「ワークスペース URL リークをスクラブするために強制プッシュするかどうか」これらはどれもプロンプトではありません。これらはすべて私が行った呼び出しであり、トレードオフに関するクロードからの情報に基づく入力が必要です。
これにより、より慎重な判断が得られるのではなく、より慎重な判断が得られます。摩擦が減ると、「ただ出荷したい」という誘惑が強くなります。リリースごとにスモークテストを行うという規律のおかげで、コンパイルはできるが機能しないコードの壁になるのを防ぐことができました。
これが、Canonical Agency でのエンジニアリングの仕事についての私の考え方です。私たちはプロンプトを販売していません。私たちは、あらゆる段階でのレシートの規律を守り、実行速度を高める判断を売りにしています。
引き出す価値のある 3 つの技術的教訓
1. 墓石のバグ。冪等性は見た目よりも難しいです。
v0.8 では、7 つのモデルに create_or_update メソッドが追加されました。セマンティクスの調整: args.name を持つリソースが Swamp のデータ層に存在する場合は、更新パスを選択します。それ以外の場合は作成します。シンプルで付加的な、真の自動化に必要なもの。
バグ: delete を呼び出した後も、Swamp の readResource は、 null ではなく、tombstone フラグを含む以前のデータを返します。私のコードはトゥームストーンを真実として扱い、存在しなくなったワークスペースリソースに対して PATCH パスを取得し、404 されました。

[切り捨てられた]

## Original Extract

On May 30th I published v0.1 of a Databricks integration pack for Swamp. By the end of the same evening it was at v0.13. Fifteen models, 100% A on every release, every model end-to-end smoke-tested. What shipped, what I learned, and why it matters if you have Databricks alongside anything else.

What Four Hours of Focused Human-AI Engineering Actually Ships · Canonical Canonical Independent AI consultancy Capabilities
← Essays What Four Hours of Focused Human-AI Engineering Actually Ships
On May 30th I published v0.1 of a Databricks integration pack for Swamp. By the end of the same evening it was at v0.13. Fifteen models, 100% A on every release, every model end-to-end smoke-tested. What shipped, what I learned, and why it matters if you have Databricks alongside anything else.
Receipts before claims. Fifteen models. Thirteen releases. One evening.
On May 30th I published v0.1 of a Databricks integration pack for Swamp , the new declarative automation framework from System Initiative. By the end of the same evening it was at v0.13, with fifteen models covering jobs, DLT pipelines, SQL warehouses, the full Unity Catalog tree, workspace secrets, permissions, DBSQL queries, and Git Repos. Every release scored 100/A on the platform's quality rubric. Every model was end-to-end smoke-tested against a real Databricks workspace. One release shipped a real bug; two releases later I shipped the fix. The whole thing took about four hours of focused work, paired with Claude.
This post is about what's in the pack, why it matters if you have Databricks alongside anything else in your stack, and what I learned about engineering velocity along the way.
Fifteen Swamp models, published to swamp.club/extensions/@mfbaig35r/databricks , with source at github.com/mfbaig35r/swamp-databricks . They cover the Databricks data engineering surface:
Compute and orchestration : jobs (Jobs API 2.2), DLT pipelines, SQL warehouses with statement execution
Workspace storage : notebooks and workspace files (two distinct Databricks object types, two distinct models, distinct upload semantics)
Unity Catalog : catalogs, schemas, tables, volumes
Access control : workspace permissions (jobs, pipelines, warehouses, notebooks, repos, queries, dashboards, alerts...), UC permissions (full grant model for catalogs, schemas, tables, volumes, functions, external locations, storage credentials)
Workspace secrets : scopes and keys, distinct from Swamp's own vault, with secret values never persisted in Swamp's data layer
DBSQL : saved queries that pair with the job model's sql_task.query.query_id field
Git Repos : workspace-attached repos with pull, branch switching, sparse checkout, and full lifecycle
Each model exposes the surface a real workspace user would touch. The job model covers Jobs API 2.2 with full task-type schema validation including for_each_task (recursive via Zod z.lazy ), dbt_task , and spark_python_task . The uc_permissions model covers all 15 UC securable types with the changes-style add/remove PATCH semantics. The sql_warehouse model includes warehouse lifecycle and the SQL Statement Execution API for running queries from a workflow. Lifecycle-managing models expose idempotent create_or_update semantics with workspace-first reconciliation, which correctly handles delete-then-recreate and out-of-band workspace deletes.
Every release scored 100/A on Swamp's quality rubric. Every model was smoke-tested end-to-end against a real Databricks Free workspace before the pack moved to the next release. "Smoke" means the actual API path: notebook uploaded to the workspace, job created via the API, run triggered, run waited to terminal state, result validated. Not type-check. That discipline caught a genuine semantic bug in v0.8's create_or_update that I shipped a fix for in v0.11. More on that below.
One reference workflow ships alongside the models: a Met Museum API ingestion pipeline at examples/api-ingest/met-museum/ , end-to-end validated, runnable on Databricks Free without credentials. It implements a universal pattern for pulling external HTTP APIs into Bronze + Silver Delta tables (rate-limited Spark fan-out via mapInPandas , raw-JSON Bronze for schema-drift immunity, typed Silver via SQL). Forking it for Stripe, GitHub, Salesforce, or any other API is a 45-90 minute exercise: change the config block at the top of the notebook, change the Silver SQL. Everything between stays.
Distribution is the part that isn't done. The pack has zero pulls as of writing. It's day one for distribution; the engineering artifact is finished. More on that at the end.
Databricks already has two ways to manage itself: the Terraform provider for infra lifecycle, and Databricks Asset Bundles (DAB) for in-Databricks pipeline definition. Both are good at what they do. Neither composes with anything outside Databricks, and neither helps you write the Python that actually does the ingest, transform, or training work inside a notebook.
The non-obvious answer the pack is built on: don't ship a generic ingest runtime. Let the agent write the notebook code on demand for the specific case, then orchestrate the agent-written notebook through Swamp's declarative graph.
The Met Museum example in the pack is the proof. I didn't hand-write the 200-line bronze ingest notebook. I described the task to Claude (pull the Met's open catalog into a Delta table, respect their rate limit, capture errors), Claude wrote the notebook, I reviewed it, smoke-tested it on Databricks Free, fixed one bug (the serverless RDD restriction), committed it. The pack's notebook.upload step then uploads that committed notebook to the workspace, and the job.run step executes it. The agent is the IDE; swamp-databricks is the runtime; the committed notebook is the algorithm.
That collapses two boundaries at once.
First , between "build a one-off integration" and "manage a production integration." Both flow through the same tools (Claude Code for authorship, swamp-databricks for orchestration). A new vendor API or a schema change is one agent turn away, not one framework release away.
Second , between Databricks and everything else. Because the workflow is a declarative Swamp graph, the same workflow that runs an agent-written Databricks notebook can also touch S3, Postgres, Cloudflare, GitHub, Slack, or any other Swamp extension in one file, with idempotent reruns, on a schedule. The notebook is the imperative core; the workflow is the declarative outer shell; the agent is what makes the imperative core cheap to write and rewrite. A minimal version of the Met workflow:
steps :
- id : schema
model : "@mfbaig35r/databricks/uc_schema"
method : create_or_update
arguments : { name : met_museum , catalog_name : workspace }
- id : notebook
model : "@mfbaig35r/databricks/notebook"
method : upload
arguments : { path : /Shared/met-museum-bronze , content : | ... }
- id : trigger
model : "@mfbaig35r/databricks/job"
method : run
arguments : { job_ref : met-museum-bronze }
- id : wait
model : "@mfbaig35r/databricks/job"
method : wait_run
arguments : { run_id : $ {{ steps.trigger.outputs.run_id }} }
- id : grants
model : "@mfbaig35r/databricks/uc_permissions"
method : update
arguments : { changes : [ ... ] }
The economic claim is concrete: integration packs that used to take weeks of generic-framework engineering drop to hours of task-specific generation plus review. The Met example took about 30 minutes from "what's a good demo API" to "running on Databricks Free with a Bronze and Silver table populated." A Stripe ingest with cursor pagination and OAuth would be roughly the same: 30-45 minutes from prompt to running pipeline.
This only works under one condition: an engineer who treats smoke tests and code review as load-bearing. Without that, you get production drift driven by prompt variation. With it, you get integration velocity that wasn't available a year ago.
Swamp is a new declarative automation framework from System Initiative, the company Adam Jacob is building since Chef. It is small and new. I picked it as the runtime for this pack for three concrete reasons:
Declarative model graph as the unit of composition. Models compose with each other and with models from other domains. Workflows are typed DAGs over models. That structural property is what makes "Databricks + everything else" expressible in one file. It's not a feature you can bolt onto an imperative tool later; it's the shape of the runtime.
First-class agent integration in the CLI. swamp repo init writes a CLAUDE.md with project rules. The extension authoring skills load automatically. The framework was designed with the assumption that AI agents would both author and operate automation, and the tooling reflects that. Working alongside Claude on this pack was friction-free in a way it wouldn't have been on Terraform or Pulumi.
Tight feedback loop in the tooling. swamp extension fmt formats and lints. swamp extension quality scores against a published 14-factor rubric. swamp model method run returns structured output usable by an agent. The registry has yank and unyank, repository-verified state, and CalVer versioning baked in. That toolchain is what made the smoke-test-every-release discipline practical instead of aspirational. The loop is short enough that running it on every change doesn't slow you down.
Honest size-of-bet line: Swamp has a small community today. Whether it finds adoption is unproven. I think the structural choices are right, and I'm willing to spend engineering time on the bet.
What changes when you pair seriously with an AI
This is the section that matters most to me. The thesis: what changes when you pair seriously with an AI is not "the AI writes code." It's that the cost of building integration packs that used to take weeks drops to hours, if you bring engineering judgment to the loop.
Concrete breakdown of who did what across this session:
The discipline that makes this not vaporware:
Every release got smoke-validated end-to-end before the next one started. Not a type-check, not a dry run. Notebook uploaded, job created, run triggered, run waited to terminal, result validated.
Bugs got caught and shipped fixes for, not papered over. The v0.8 to v0.11 tombstone bug fix is the cleanest example: I noticed the broken semantic mid-session, decided it was real, decided it was worth a dedicated release, shipped the fix, validated the fix, updated the memory file so I won't repeat the mistake.
The session is preserved in the git history, the release notes, and the registry's version log. Anyone can audit it. The 13 release notes are themselves a session log: what shipped, what was smoke-tested, what bugs were caught, what's still gap.
What I learned about my own work:
I think faster with someone (something) typing the first draft while I'm still deciding what the right draft is. The bottleneck moves from typing to deciding. That's a real shift in what engineering work feels like.
The judgment calls are the value. "Which task types to schema-validate first." "Whether the tombstone is a real bug or a tolerable quirk." "Whether to yank v0.2 publicly." "Whether to force-push to scrub a workspace URL leak." None of these are prompts. All of them are calls I made, with informed input from Claude on the tradeoffs.
This rewards more careful judgment, not less. The temptation to "just ship" gets stronger when the friction drops. The discipline of smoke-testing every release is what kept this from becoming a wall of code that compiles but doesn't work.
This is how I think about engineering work at Canonical Agency. We are not selling prompts. We are selling judgment that compounds with execution velocity, with the discipline of receipts at every step.
Three technical lessons worth pulling out
1. The tombstone bug. Idempotency is harder than it looks.
v0.8 added create_or_update methods to seven models. Reconcile semantics: if a resource with args.name exists in Swamp's data layer, take the update path; otherwise create. Simple, additive, what real automation needs.
Bug: after calling delete , Swamp's readResource still returns the prior data with a tombstone flag, not null . My code treated tombstone-as-truthy, took the PATCH path against a workspace resource that no longer existed, and 404'd.

[truncated]
