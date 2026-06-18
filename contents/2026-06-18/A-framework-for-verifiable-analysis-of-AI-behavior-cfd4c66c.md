---
source: "https://transluce.org/docent/blog/analysis-plans"
hn_url: "https://news.ycombinator.com/item?id=48577899"
title: "A framework for verifiable analysis of AI behavior"
article_title: "Introducing Analysis Plans | Transluce AI"
author: "mengk"
captured_at: "2026-06-18T01:03:12Z"
capture_tool: "hn-digest"
hn_id: 48577899
score: 1
comments: 0
posted_at: "2026-06-17T22:32:23Z"
tags:
  - hacker-news
  - translated
---

# A framework for verifiable analysis of AI behavior

- HN: [48577899](https://news.ycombinator.com/item?id=48577899)
- Source: [transluce.org](https://transluce.org/docent/blog/analysis-plans)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T22:32:23Z

## Translation

タイトル: AI 動作の検証可能な分析のためのフレームワーク
記事のタイトル: 分析計画の紹介 |トランスルースAI

記事本文:
Transluce Transluce Transluce Transluce Transluce Transluce 私たちの仕事 研究ツール
Changelog Blog Community Docent はパブリック アルファ版になりました。こちらから最新バージョンの Docent をお試しください →
AI の動作を検証可能に分析するためのフレームワーク
AI エージェントの開発は、複雑なデータ分析の問題です。エージェントが正しく動作しているかどうかを知るには、ベンチマーク スコアだけでなく、その動作の詳細、つまりトレーニング中に戦略がどのように変化するかを追跡する必要があります。新しい足場の性能が低下するのはなぜですか?報酬ハッキングはありますか？これらの質問に答えるには、手元のデータセットに合わせて調整された定量的分析と定性的分析を組み合わせる必要があります。
コーディング エージェントはこの作業を加速する可能性を秘めていますが、微妙な間違いを犯す傾向があります。データを間違って解析したり、不当な仮定をしたり、例を厳選して誤解を招くような物語を提示したりする可能性があります。これらの間違いは、最終的な出力では明らかではありません。結論を信頼するには、それがどのように生み出されたのかを正確に検証する必要があります。しかし、コーディング エージェントが実行したすべてのアクションを確認するのは面倒です。重要な方法論の決定は数百行のログに埋もれてしまいます。
この問題は、AI の動作の検証可能な分析のためのフレームワークである分析計画を開発する動機となりました。分析計画は、コーディング エージェントが使用できる Python API で指定されます。実行の準備が整うと、Web インターフェイスに表示され、人間が実行されたすべてのステップを理解し、検証できるようになります。
分析計画には、次の 2 種類のステップが含まれます。
クエリ ステップでは、DQL (Docent の SQL サブセット) を使用してデータをフィルタリング、グループ化、結合します。各ステップは、そのクエリと結果の対話型テーブルとともに表示されます。
読み取りステップでは、LLM を使用してクエリ ステップからのデータを分析し、テキストの概要や構造化された判断を生成します。クレームママ

LLM による de には、その文脈における特定の項目への引用が含まれています。
これら 2 つのステップ タイプをカスタマイズして組み合わせて、複雑な分析パイプラインを構築できます。読み取り値はクエリが生成するあらゆるデータを受け入れることができ、クエリはあらゆる読み取り結果に対して実行できます。各ステップで、結果はその結果を生成した正確な計算まで追跡され、フローを検査、監査、および調整できるようになります。
一般的なソフトウェア エンジニアリング ベンチマークで不正行為を検出することで、これがどのようなものになるかを見てみましょう。
デモ: SWE-rebench での不審な動作を特定する
不正行為は、評価結果を解釈する際によくある厄介な問題です。モデルのハードコード化されたテスト、偽りの成功宣言、および汚れた環境を悪用してソリューションをコピーすることはよく知られています。不正行為の割合を測定することは、ベンチマーク スコアがモデルの機能の有効な実証をどの程度表しているかを理解するために不可欠です。約 15 分で、Docent を使用して、エージェントが解決できる最新の GitHub PR の数を測定するソフトウェア エンジニアリング ベンチマークである SWE-rebench での不正行為の事例を発見しました。このリンクで Docent の SWE 再ベンチ トレースを表示できます。
1. コーディング エージェントに学びたいことを説明します。
まず、Claude Code に不正行為の可能性についてエージェントの実行をスコアリングするよう促します。 Docent スキルのおかげで、クロードは私たちの質問を分析計画に変える方法を知っています。次のような短い Python スクリプトを作成します。
docent import から Docent
client = Docent()
実行 = client.query(
コレクション_ID、
「agent_runs から実行されるとして、agent_runs.id を選択」
「WHERE CAST(metadata_json->'scores'->>'resolved' AS DOUBLE PRECISION) = 1.0」
"agent_runs.id で注文 LIMIT 200",
name=f"UUID ごとに解決されたサンプル 200 件の実行",
)
DETECTOR_PROMPT = "..." # 簡潔にするために省略
OUTPUT_SCHEMA = { ... } # 簡潔にするために省略
検出 = client.read(
プロンプト_テンプレート=[runs.run.a

s_type("agent_run"), DETECTOR_PROMPT],
モデル = "openai/gpt-5.5",
reasoning_effort="中",
出力スキーマ=OUTPUT_SCHEMA、
name="軌道からの不正行為のフラグ",
)
Claude がこのスクリプトを実行しても、Docent SDK は読み取りをすぐには実行しません。代わりに、要求されているすべての測定値 (この場合は 1 つだけ) のグラフを作成し、それを分析計画として Docent にアップロードします。これにより、LLM 呼び出しが完了するのを待つ前に、提案された分析を確認できます。
2. 分析計画の構造
この分析計画は、エージェント実行メタデータを [resolved=1] にフィルター処理することで、成功した実行を選択する DQL クエリから始まります。その後、結果が読み取りステップに渡され、不正行為のスコアが実行されます。
読み取りステップには次のコンポーネントがあります。
パラメーターを含むプロンプト テンプレート。不正行為を検出するには、1 つのエージェントの実行をパラメータとして、不審性を定義するルーブリックを作成します。 Docent UI には、インラインで表されたパラメーターを含む完全なプロンプト テンプレートが表示されます。以下の「Context: run」チップをクリックすると、LLM に渡される実行レベルのメタデータに関する追加の詳細が表示されます。デフォルトでは、メタデータは渡されません。
前のクエリステップからの引数。計画の前のクエリ ステップでは、resolved=1 の実行のテーブルが生成されました。このテーブルの各行について、Docent はエージェント実行の全文をプロンプト テンプレートに置き換え、それを使用して LLM を呼び出します。
カスタム出力スキーマ。この読み取りにより、疑わしさを表す 1 ～ 10 の整数スコアと、引用を含む文字列説明が出力されます。 Docent は、後の手順でデータ形式の問題が発生しないように、LLM 出力がこのスキーマに準拠していることを保証します。
朗読は表現力豊かになるように設計されています。一般的な使用例には、特定の動作に対する実行の判断、以前の読み取り結果のクラスタリングによる高品質なデータの抽出などが含まれます。

h レベルの傾向、長期にわたるエージェント実行の要約、または同じタスクの 2 つのロールアウトの比較。適切に機能する測定値を取得したら、エージェントが後で再利用できるように、それをプリセットとして保存できます。
紫色のハイライトは、この読み取り値がまだ承認を待っていることを示しているため、承認しましょう。 (Docent SDK は、ステップをプログラム的に承認する方法も提供します。コーディング エージェントに「分析計画を自動承認」して手動レビューをスキップするように指示できます。)
読み取り結果は検証を容易にするように設計されています。読み取りが実行されると、その結果が表に表示されます。各行をクリックすると、新しいペインに結果が展開され、完全な出力と、読み取りモデル、推論作業、使用されたトークンなどのメタデータが表示されます。
読み取り出力の主張には、トランスクリプトの特定の部分への引用が含まれます。青色の引用アイコンをクリックしてトランスクリプトの関連部分にジャンプし、主張を再確認します。
集計された全体像を取得するには、各不審スコアを受け取ったエージェントの実行数をカウントするための DQL クエリを作成するようにコーディング エージェントに依頼できます。
最終結果の上にある [クエリ] トグルをクリックすると、DQL クエリを一目で確認できます。正確な DQL を読み取ることで、誤ったグループ化、誤ったフィルタリング、または誤った方法を使用して計算された平均などのエラーをチェックできます。
5. 疑わしい例の検証
報酬ハッキングの裁判官が誤検知を起こしやすいことは有名です。結果を新しい読み取りに渡して、フォローアップの質問をすることができます。
私たちのエージェントは DQL を使用して 4 つの疑わしい実行を抽出し、それぞれを検証するための読み取り値を作成します。検証者の読み取りにより、不正行為の例のほとんどは実際には疑わしいが、少なくとも 1 つは誤検知であることが確認されます。
1 つの不審な実行には、将来を見据えた不正行為が含まれているようです: エージェントが git 履歴を調べます

別のコミットからコードを復元します。しかし、検証者はこれを誤検知として正しく識別します。つまり、過去のコミットによってリグレッションが導入されたため、コードベースの元の状態を復元することが有効なパッチとなります。
分析計画は、1 つの統合フレームワークを通じてさまざまなアプリケーションをサポートします。モデルの比較、障害モードの特定、予期しない動作の発見にこれらを使用した方法を次に示します。
エージェントのパフォーマンスを向上させるには、対処可能な障害モードを特定することが重要です。トレースを大規模に分析すると、ハーネス、トレーニングの組み合わせ、またはプロンプトの変更を示唆するモデルの弱点の傾向を明らかにするのに役立ちます。
蔓延しており対処可能な障害を明らかにするために、再帰的クラスタリングを使用しました。Docent は読み取り値を使用して、各トランスクリプトに存在する障害モードを特定し、それらの障害を広範なカテゴリに分類します。最大の 3 つのカテゴリについては、Docent にサブクラスターを作成してもらいます。
トレーニング後のよくある質問は、「これら 2 つのチェックポイントの間で何が起こったのか?」というものです。突然の損失の減少は、その時点のログを調べなければ不可解です。
Terminal-Bench では、公式リーダーボードで GPT-5.1 Codex のパフォーマンスが GPT-5 Codex より 6.5 パーセント悪いことに気づいた後、同じ質問をしました。回帰を診断するために、Docent を使用して、GPT-5.1 Codex が GPT-5 Codex よりもロールアウト全体で平均報酬が低いタスクをフィルターしました。これらのタスクに関して、Docent は GPT-5.1 Codex から失敗した実行を 1 つ選択し、それを GPT-5 Codex からの成功した実行と組み合わせて、モデルの戦略が分岐した場所を特定しました。
分析プランはすべての Docent ユーザーが利用できるようになりました。これらを試してみるには、Docent SDK とエージェント スキルをインストールします。サンプル データに対して分析を実行するか、独自のデータを取り込みます。
ご質問またはフィードバックはありますか?ぜひ

あなたからの連絡をお待ちしています。 Docent Community Slack でチームとチャットします。

## Original Extract

Transluce Transluce Transluce Transluce Transluce Transluce Our Work Research Tools
Changelog Blog Community Docent is now in public alpha. Try the latest version of Docent here →
A framework for verifiable analysis of AI behavior
Developing an AI agent is a complex data analysis problem. To know if the agent is working correctly, we need to track not just benchmark scores but the details of its behavior: how do the strategies change over the course of training? Why does the new scaffold perform worse? Is there reward hacking? Answering these questions requires a combination of quantitative and qualitative analysis tailored to the dataset at hand.
Coding agents have the potential to accelerate this work, but they’re prone to subtle mistakes: they might parse data incorrectly, make unjustified assumptions, or cherry-pick examples and present a misleading narrative. These mistakes aren’t apparent in the final output. To trust a conclusion, we need to verify exactly how it was produced. But reviewing all the actions taken by a coding agent is tedious—crucial methodology decisions get buried in hundreds of lines of logs.
This problem motivated us to develop analysis plans, a framework for verifiable analysis of AI behavior. Analysis plans are specified in a Python API that any coding agent can work with. When they’re ready to run, they appear in a web interface that lets humans understand and verify every step that was taken.
Analysis plans contain two types of steps:
Query steps filter, group, and join your data using DQL (Docent’s subset of SQL). Each step is displayed with its query and an interactive table of the results.
Reading steps use an LLM to analyze data from a query step, producing a text summary and/or a structured judgment. Claims made by the LLM come with citations to specific items in its context.
These two step types can be customized and combined to build complex analysis pipelines. Readings can accept any data that a query produces, and queries can run over any reading results. At each step, results are traced to the exact computation that produced them, enabling you to inspect, audit, and refine the flow.
Let’s see what this looks like by detecting cheating on a popular software engineering benchmark.
Demo: identifying suspicious behaviors in SWE-rebench
Cheating is a common thorn when interpreting evaluation results: models famously hard-code tests, falsely claim success, and exploit unclean environments to copy solutions . Measuring rates of cheating is essential for understanding how much of a benchmark score represents a valid demonstration of model capability. In about 15 minutes, we used Docent to discover instances of cheating on SWE-rebench, a software engineering benchmark that measures how many recent GitHub PRs an agent can resolve. You can view the SWE-rebench traces in Docent at this link .
1. Explain to your coding agent what you want to learn
We start by prompting Claude Code to score agent runs for potential cheating. Thanks to the Docent skill, Claude knows how to turn our question into an analysis plan. It writes a short Python script like the following.
from docent import Docent
client = Docent()
runs = client.query(
COLLECTION_ID,
"SELECT agent_runs.id AS run FROM agent_runs "
"WHERE CAST(metadata_json->'scores'->>'resolved' AS DOUBLE PRECISION) = 1.0 "
"ORDER BY agent_runs.id LIMIT 200",
name=f"Sample 200 resolved runs by UUID",
)
DETECTOR_PROMPT = "..." # omitted for brevity
OUTPUT_SCHEMA = { ... } # omitted for brevity
detect = client.read(
prompt_template=[runs.run.as_type("agent_run"), DETECTOR_PROMPT],
model="openai/gpt-5.5",
reasoning_effort="medium",
output_schema=OUTPUT_SCHEMA,
name="Flag cheating from trajectory",
)
When Claude runs this script, the Docent SDK doesn’t execute readings immediately. Instead, it builds up a graph of all the readings that are being requested (in this case, just one) and uploads it to Docent as an analysis plan. This lets us review the proposed analysis before waiting for LLM calls to complete.
2. Anatomy of an analysis plan
This analysis plan starts with a DQL query to select successful runs by filtering the agent run metadata to resolved=1 . After that, it passes the results to a reading step, which scores runs for cheating.
A reading step has the following components:
A prompt template with parameters. To detect cheating, we create a rubric that defines suspiciousness, taking one agent run as a parameter. The Docent UI shows the full prompt template, with parameters represented inline. Clicking on the "Context: run" chip below exposes additional detail on what run-level metadata is passed to the LLM. By default, no metadata is passed.
Arguments from a previous query step. The previous query step in our plan produced a table of runs where resolved=1 . For each row in this table, Docent will substitute the full text of the agent run into the prompt template and call an LLM with it.
Custom output schema. This reading outputs an integer score from 1 to 10 that represents suspiciousness, and a string explanation containing citations. Docent ensures LLM output conforms to this schema so that later steps won’t run into data formatting issues.
Readings are designed to be expressive. Common use cases include judging runs for a specific behavior, clustering results of previous readings to extract high-level trends, summarizing long agent runs, or comparing two rollouts of the same task. Once you have a reading that works well, you can save it as a preset so agents can reuse it later.
The purple highlight indicates that this reading is still waiting for our approval, so let’s approve it. (The Docent SDK also provides a way to approve steps programmatically; you can tell your coding agent to “auto accept analysis plan” and skip the manual review.)
Reading results are designed for ease of verification. As a reading executes, its results appear in a table. Clicking on each row expands the result in a new pane with the full output and metadata like the reading model, reasoning effort, and tokens used.
Claims in the reading output include citations to specific parts of the transcript. Click on the blue quote icons to jump to relevant parts of the transcript and double-check the claims.
To get an aggregate picture, we can ask our coding agent to write a DQL query for counting the number of agent runs that received each suspiciousness score.
We can verify the DQL query at a glance by clicking on the Query toggle above the final results. Reading over the precise DQL lets you check for errors like incorrect grouping, incorrect filtering, or a mean calculated using the wrong method.
5. Verifying suspicious examples
Reward hacking judges are famously prone to false positives. We can pass the results to a new reading and ask a follow-up question.
Our agent uses DQL to extract the four suspicious runs and creates a reading to verify each of them. The verifier reading confirms that most of the cheating examples are actually suspicious, but at least one is a false positive.
One suspicious run appears to contain cheating by forward-looking: the agent examines the git history and restores code from a different commit. But the verifier correctly identifies this as a false positive: a past commit introduced a regression, so restoring the original state of the codebase is a valid patch.
Analysis plans support a range of applications through one unified framework. Here’s how we used them for comparing models, identifying failure modes, and discovering unexpected behaviors.
Identifying actionable failure modes is crucial for improving agent performance. Analyzing traces at scale can help us surface trends in model weaknesses that suggest a change to the harness, training mixture, or prompting.
To surface failures that are prevalent and actionable, we used recursive clustering: Docent uses readings to identify the failure modes present in each transcript and cluster those failures into broad categories. For the three largest categories, we have Docent create sub-clusters.
A common question that arises during post-training is “what happened between these two checkpoints?” Sudden drops in the loss are inscrutable without looking through the logs from that point in time.
On Terminal-Bench, we asked the same question after noticing that GPT-5.1 Codex performed 6.5 percentage points worse than GPT-5 Codex on the official leaderboard. To diagnose the regression, we used Docent to filter for tasks where GPT-5.1 Codex receives lower average reward across rollouts than GPT-5 Codex. On these tasks, Docent selected one failed run from GPT-5.1 Codex, paired it with the successful runs from GPT-5 Codex, and identified where the models’ strategies diverged.
Analysis plans are now available to all Docent users. To try them out, install the Docent SDK and agent skill . Run an analysis on our sample data, or ingest your own.
Questions or feedback? We'd love to hear from you. Chat with the team in Docent Community Slack .
