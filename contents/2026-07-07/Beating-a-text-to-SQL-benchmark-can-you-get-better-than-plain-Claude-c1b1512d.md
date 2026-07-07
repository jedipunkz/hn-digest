---
source: "https://motley.ai/blog-posts/bird-interact-benchmark-top-score/"
hn_url: "https://news.ycombinator.com/item?id=48820308"
title: "Beating a text-to-SQL benchmark: can you get better than plain Claude?"
article_title: "Topping the BIRD-INTERACT benchmark with SLayer | Motley"
author: "yannranchere"
captured_at: "2026-07-07T17:06:33Z"
capture_tool: "hn-digest"
hn_id: 48820308
score: 2
comments: 0
posted_at: "2026-07-07T16:44:58Z"
tags:
  - hacker-news
  - translated
---

# Beating a text-to-SQL benchmark: can you get better than plain Claude?

- HN: [48820308](https://news.ycombinator.com/item?id=48820308)
- Source: [motley.ai](https://motley.ai/blog-posts/bird-interact-benchmark-top-score/)
- Score: 2
- Comments: 0
- Posted: 2026-07-07T16:44:58Z

## Translation

タイトル: text-to-SQL ベンチマークの達成: プレーンなクロードよりも優れたものを得ることができますか?
記事のタイトル: SLayer で BIRD-INTERACT ベンチマークを超える |モトリー
説明: Claude SDK と SLayer を使用して、エージェントは BIRD-INTERACT (a-interact、mini-interact) で 75.3% の合格率を達成しました。これに対し、公式最高の 36.33% は…

記事本文:
SLayer で BIRD-INTERACT ベンチマークを上回る | Motley SLayer v0.9 はライブです。AI エージェント用に構築されたオープンソースのセマンティック レイヤーです。 GitHub で表示 → Product SLayer エージェント用のオープンソース セマンティック レイヤー Motley Cloud マネージド プラットフォーム: ホスティング、アクセス、レポート
ソリューション ビルダー向け アプリのデータ上に管理された MCP サーバーを出荷する データ チーム向け 組織全体に管理された分析エージェントを提供する
星 115 個 デモを予約する Cloud Product SLayer に移動する
BIRD-INTERACT ベンチマークを超える: セマンティック レイヤーを追加することで、単純なクロードよりもはるかに優れたものを得ることができますか?
製品をより良いものにする最善の方法は、それをさまざまなケースで使用することです。 Motley では、エージェント用のオープンソースのセマンティック レイヤーである SLayer を構築する際に、すでに顧客や設計パートナーから多くのことを学んできました。次にその多様性を探すのに自然な場所は、text-to-SQL ベンチマークでした。
Claude SDK と SLayer (私たちが開発したオープンソースのセマンティック レイヤー) を使用すると、エージェントは BIRD-INTERACT (a-interact、mini-interact) で 75.3% の合格率を達成しました。これに対し、同等の Lite リーダーボードでは公式最高の 36.33% でした。
最大の要因はセマンティック レイヤーではなく、エージェント ハーネスでした。 Claude SDK 単独では、生の SQL エージェントが 71.7% に達しましたが、PydanticAI ベースのセットアップでは、はるかにわずかな改善しか得られませんでした。
ベンチマークのゴールドアンサーの大部分が間違っているため、注釈エージェントを構築し、元の回答と注釈付きの回答の両方に対する結果をレポートします。
SLayer は依然としてトップに優位性を追加しました。元のゴールド回答に対して 75.3% 対 71.7%、修正された (「注釈付き」) 回答に対して 83.7% 対 79.0% でした。
この実行では、SLayer の取り込みが完全に自律的に実行されたため、人間が管理する設定では数値がさらに高くなる可能性があります。
SLayer に慣れていない場合は、これが何をするのかを説明します。

エージェントとデータベースの間で、エージェントがクエリを実行する定義が保持されます。テーブルとフィールドのセマンティック モデル (手書きまたはスキーマから自動取り込み)、「純収益」などのメトリクス定義、およびメモリとして添付されたフリーテキストのナレッジです。エージェントは、生の SQL を記述する代わりに、SLayer が SQL にコンパイルする構造化 JSON を通じてクエリを実行します。 text-to-SQL ベンチマークはその仕事の大部分を実行し、ナレッジ ベースが添付された不慣れなデータベースを介してビジネス上の質問に答えます。そのため、SLayer のストレス テストを行う自然な方法でした。
そこで、私たちは 1 つを実行しました。 Claude SDK と SLayer を使用すると、エージェントは BIRD-INTERACT (a-interact、mini-interact) で 75.3% の合格率を達成しました。これに対し、同等の Lite リーダーボードでは公式最高が 36.33% でした 1 。ブーム。見出し番号。現実はもっと微妙で、そこに至るプロセスはもっと興味深いものです。
この投稿はシリーズの最初のものです。ここでは、私たちがどのようにしてそこに到達したのか、そしてそのランニングが私たちに何を教えてくれたのかについて説明します。後の投稿では、個々のレッスンについてさらに詳しく説明します。
Text-to-SQL ベンチマークは、魅力的で厄介な領域です。 BIRD ファミリだけを見ても、約 20 の異なるサブベンチマークに相当します。
まず、オリジナルの BIRD とその完全なサブセットと開発サブセットがあります。次に、Base-Full、Base-Lite、Large-Full、および Large-Lite バリアント (および時折 SQLite バリアントが投入される) を備えた LiveSQLBench があります。次に、Full、Lite、mini-interact の各バリアントを備えた BIRD-INTERACT があり、それぞれに c-interact と a-interact のフレーバーがあり、エージェントの予算に制約があるかどうかに応じて個別のリーダーボードが表示されます。そして、独自のサブベンチマークを備えた BIRD-CRITIC があります。
家族構成を並べるとこんな感じです。
これは、コードを 1 行書く前に、20 近くの異なるリーダーボードを作成することになります。
条件下で SLlayer を実行したいと考えていました

実際の使用方法に近いもので、エージェントとして Claude SDK、LLM として Opus を意味します。そのため、1 回の完全なベンチマークの実行コストは、小規模なベンチマークでは約 1000 ドルとなり、大規模なベンチマークではさらに高くなります。すべてのフレーバーにわたって、必然的に発生する問題を修正するために実行し、再実行するという選択肢はありませんでした。まず最初に 1 つに決める必要がありました。
私たちはBIRD-INTERACT、a-interactフレーバーを選択しました。 BIRD-INTERACT の特徴は、単発の質問ではなく、実際の作業環境をシミュレートしようとすることです。各データベースは、階層的なナレッジ ベース、メタデータ ファイル、および関数駆動型の user_sim シミュレーターと結合されているため、エージェントは人間が介入することなく、説明を求め、ナレッジを取得し、実行エラーから独自に回復できます。 a-interact フレーバーは、これを最も進化させたものです。これは、ユーザー シミュレーターにいつ質問するか、いつデータベースを探索するか、最終回答を送信するまでに何ターンかかるかをモデルが自ら決定する、オープンエンドのエージェント設定です。これは、エージェントが実際に現実世界でどのように動作するかに最も近いフレーバーであるため、これを選択しました。
mini-interact を選択したのは、Postgres ではなく SQLite に基づいているため、実行に手間がかからず、完全バージョンよりも小さいためです。
他の選択肢と比較して、これら 2 つの選択肢が勝った理由は次のとおりです。
ミニインタラクトの公式リーダーボード結果はまだありません。 Lite バージョン用のものもあり、同様のタスクを使用しますが、SQLite ではなく Postgres データベースを使用します。 a-interact での最良の結果は、ベンチマーク作成者が提供する単純なエージェント ラッパーを備えた Opus-4.6 を使用した場合の 36.33% です。これは、SQLite と Postgres の注意点を前に記した上で比較する数値です。
次の課題は、g のかなりのシェアが

ベンチマークが提供する古い答えは単純に間違っています。これは、たとえばここやここなどで何度か指摘されています。
これは本当に問題です。 SLayer とエージェントのハーネスの両方を改善するには、エージェントが実際に間違っていたケースと、エージェントの回答は正しかったが、間違った参照回答に対して評価された偽陰性のケースを区別する必要があります。
-- タスク: 「各登録の精度比と信頼レベルを表示します。」
-- 登録 ID、プロジェクト ID、精度、エラー値、計算された比率 (RAR)、
-- そしてそれが意味する信頼度。」
-- ベンチマークのゴールドアンサー (信頼レベル分岐):
ケース
rar > 1 の場合。 5 と sr 。 refmark LIKE '%Target%' THEN 'High Confidence'
rar が 1 の間の場合。 0 と 1 。 5 その後「中程度の信頼度」
ELSE 「信頼性が低い」
END AS 信頼レベル
-- 間違っている理由: KB #44 によると、「高信頼性」では、LogMethod に「ターゲット」を含める必要があります。
-- ゴールドは代わりに、数値コードのみを保持する refmark と一致します。
-- ('40', '31', '25') なので、高信頼性ブランチは起動せず、0 行を返します。
-- ログメソッド (「ターゲットベース」、「ハイブリッド」、「機能ベース」) でのマッチングにより、意図した結果が得られます。
-- 50 個の高信頼行。
そこで私たちが最初に構築したのは、アノテーション エージェントでした。タスクがエージェントに提供したすべてのデータを考慮して、各タスクを検討し、次の 3 つのことを決定しました。
金の答えが唯一の正解かどうか、
それがそのタスクとまったく互換性があるかどうか、
そして、タスクステートメントを考慮すると、同様に有効な他の回答があるかどうか。
これらのアノテーション自体は確率的なものであるため、成功率を 2 つの方法で報告します。元のゴールドアンサーによって暗示されるものと、アノテーションによって暗示されるものです。注釈付きの数値も福音のような真実ではありませんが、数値よりもはるかに優れた評価シグナルを与えます。

元の答えのみに依存します。
エージェント ハーネスの効果と SLayer 使用の効果を区別するために、すべてのエージェント バリアントについて、できる限り類似した 2 つのバージョンを構築しました。 1 つは、SLayer を使用してデータベースと関連するナレッジ ベースをイントロスペクトし、回答を送信しました。もう 1 つは、生の SQL とベンチマーク リポジトリに付属のツールを使用しました。
SLayer バージョンでは、標準の SLayer スキーマ取り込みを使用してモデルを事前に設定し、元のスキーマ内の JSON 列の文書化されたコンポーネントごとにフィールドを決定的に作成しました。これらのドキュメント自体は JSON 構造化されているため、取り込みは簡単でした。
// 元のスキーマ: 文書化された 1 つの JSON 列、returns.return_details
// (ベンチマークは JSON の各リーフを構造化テキストとして文書化します)
{
"column_meaning" : "返品理由、承認、配送メタデータをグループ化する JSONB 列。" 、
"フィールドの意味" : {
「詐欺」: {
"risk_level" : "VARCHAR(50)。リターンの不正リスク レベル。NULL は評価されていないことを意味します。"
}
}
}
# そこから決定論的に作成された SLayer フィールド (文書化された JSON パスごとに 1 つ)
- 名前 : return_details__fraud__risk_level
SQL : JSON_EXTRACT(return_details, '$.fraud.risk_level')
タイプ: テキスト
description : 「返品の不正リスク レベル。NULL は不正リスク レベルが評価されていないことを意味します。」
メタ :
派生元 :
json_col : return_details
パス: [不正行為、リスクレベル]
ここで 1 つ明確にしておきたいのは、ベンチマークの場合、この取り込み全体が完全に自律的かつ決定論的に実行され、人間や LLM が関与することはありませんでした。これは、比較をクリーンに保ち、結果を再現可能にするために意図的に行われたものです。実際の使用では、これは SLayer の本来の動作とは異なります。この段階で知識ベース項目をエンコードするための人的入力は、大幅に追加的になります。したがって、数字を扱います

自動セットアップによって得られる内容を以下に示します。
また、関連する概念を文書化し、相互に参照できるナレッジ ベースのアイテムである生のテキストを SLayer メモリとして取り込みました。知識ベースの項目が相互参照する方法は、記憶構造と完全に一致していました。
エージェント ハーネスについては、2 つのセットアップを試しました。
カスタム サブエージェント構造を備えた PydanticAI は、必要なすべての概念に対して探索および定義サブエージェントを生成します。
Claude SDK と適切なツールセット。
エージェントには全体的に Opus-4.7 を使用し、 user_sim には Sonnet を使用しました。
ベンチマークが SLayer に教えてくれたこと
ベンチマークには非常に多様なクエリのセットが含まれているため、SLayer バージョンに対してベンチマークを実行することは、SLayer 自体にとって素晴らしいトレーニングになりました。
これにより、いくつかの特殊なバグが発生しました。 DSL ではまだまったく表現できないいくつかの種類のクエリが指摘されていたため、それを追加しました。そして、エージェントが行った、ほぼ正しいとは言えないツール呼び出しのコレクションが得られました。これを使用して、クエリ解析中に自動修正するパターンのセットを拡大しました。
私たちは SLayer 用ツールの設計において独自のアプローチを採用しています。つまり、エージェントによる潜在的な失敗を考慮し、意図が明らかな場合は自己修正します。詳細については、以前の投稿「MCP ツールはどのくらい寛容ですか?」をご覧ください。
# エージェントが書いた内容 (SQL スタイルの集計 -- 実行中に 265 回出現):
カウント(*)
合計(収益)
# SLayer が自動修正する内容 (ネイティブのコロン構文):
*:カウント
収益:合計
これはスコア以上に重要な部分です。これほど多様なベンチマークでは、厳選されたテスト スイートでは決して得られないギャップが表面化します。
これは、SLayer が現在の主要モデルの結果を 2 倍以上改善していることを誇ることができる点です。公式のベンチマーク最高解像度の 36.33% に対して、SLayer では 75.3% です。

究極。
しかし、より正直な結論は、エージェント ハーネスの選択が合格率に大きな影響を与えたということです。
私たちは、PydanticAI (ツールを使用した基本的なループ。ツールにはタスク固有のサブエージェントを再帰的に生成する機能が含まれていました) を使用して構築されたカスタム エージェントから始めました。ただし、claude_sdk を使用すると合格率がはるかに高かったため (比較に使用した 53 タスク セットの場合、PydanticAI ベースの場合は 43%、claude_sdk の場合は 64%)、全体的に claude_sdk に切り替えました。残念ながら、コスト上の理由から、タスクセット全体にわたる 2 つの包括的な比較は行われませんでした。
Claude SDK エージェントでは話が異なり、全体的なパフォーマンスが大幅に向上しました。そこでは、SLayer を使用するエージェントのパフォーマンスが、生の SQL ツインよりわずかに優れていました。
したがって、SLayer を使用した Claude SDK は、元のゴールド アンサーに対して 75.3%、注釈付きのアンサーに対して 83.7% の合格率に達しました。
名誉のために言っておきますが、BIRD-INTERACT は曖昧さを生み出します。これがインタラクティブ設定のポイントであり、エージェントがユーザー シミュレーターに質問して意図を特定します。ただし、その曖昧さはタスクごとに 1 回、単一の標準的な答えに対して解決されます。ビジネス上の質問は異なります。同じ質問が毎月、表現を変えて繰り返されます

[切り捨てられた]

## Original Extract

With the Claude SDK and SLayer, an agent reached a 75.3% pass rate on BIRD-INTERACT (a-interact, mini-interact), against an official best of 36.33% on…

Topping the BIRD-INTERACT benchmark with SLayer | Motley SLayer v0.9 is live: the open-source semantic layer built for AI agents. View on GitHub → Product SLayer Open-source semantic layer for agents Motley Cloud The managed platform: hosting, access, reporting
Solutions For builders Ship a governed MCP server over your app's data For data teams Give your whole org a governed analytics agent
115 stars Book a demo Go to Cloud Product SLayer
Beating the BIRD-INTERACT benchmark: can you get much better than plain Claude by adding a semantic layer?
The best way to make a product good is to use it in a wide variety of cases. At Motley, we had already learned a lot from our customers and design partners while building SLayer , the open-source semantic layer for agents. The next natural place to look for that variety was text-to-SQL benchmarks.
With the Claude SDK and SLayer (the open-source semantic layer we developed), an agent reached a 75.3% pass rate on BIRD-INTERACT (a-interact, mini-interact), against the official best of 36.33% on the comparable Lite leaderboard.
The biggest lever was not the semantic layer, but the agent harness . The Claude SDK alone took a raw-SQL agent to 71.7%, while the PydanticAI-based setup yielded a much more modest improvement.
A substantial share of the benchmark’s gold answers are wrong, so we built an annotation agent and report results against both the original and the annotated answers.
SLayer still added an edge on top: 75.3% vs 71.7% against the original gold answers, and 83.7% vs 79.0% against the corrected (“annotated”) ones.
The SLayer ingestion ran fully autonomously for this run, so the numbers would probably be higher in a human-curated setup.
If you’re unfamiliar with SLayer, here’s what it does: sitting between the agent and the database, it holds the definitions the agent queries against: semantic models of your tables and fields (hand-written or auto-ingested from the schema), metric definitions like “net revenue”, and free-text knowledge attached as memories. Instead of writing raw SQL, the agent queries through structured JSON that SLayer compiles into SQL. A text-to-SQL benchmark exercises a large part of that job, answering business questions over an unfamiliar database with a knowledge base attached, so it was a natural way to stress-test SLayer.
So we ran one. With the Claude SDK and SLayer, an agent reached a 75.3% pass rate on BIRD-INTERACT (a-interact, mini-interact), compared to an official best of 36.33% on the comparable Lite leaderboard 1 . Boom. Headline number. The reality is more subtle, and the process to get there more interesting.
This post is the first in a series: here we cover how we got there and what the run taught us. Later posts go deeper on the individual lessons.
Text-to-SQL benchmarks are a fascinating, messy domain. Even if you just look at the BIRD family, it amounts to almost 20 different sub-benchmarks.
First there is the original BIRD, with its full and dev subsets. Then there is LiveSQLBench, with Base-Full, Base-Lite, Large-Full, and Large-Lite variants (and the occasional SQLite variant thrown in). Then there is BIRD-INTERACT with Full, Lite, and mini-interact variants, and each of those comes in c-interact and a-interact flavors, with separate leaderboards depending on whether the agent was budget-constrained. And then there is BIRD-CRITIC, with its own sub-benchmarks.
Laid out, the family looks like this:
That is close to 20 distinct leaderboards before you have written a line of code.
We wanted to run SLayer in conditions close to how it would actually be used, which meant the Claude SDK for the agent and Opus as the LLM. That made a single full benchmark run cost on the order of $1000 for the smaller benchmarks, more for the larger ones. Running, and re-running to correct the issues that inevitably arose, across all the flavors was not an option. We had to settle on one to start with.
We chose BIRD-INTERACT, a-interact flavor . What sets BIRD-INTERACT apart is that it tries to simulate a real working environment rather than a one-shot question. Each database is coupled with a hierarchical knowledge base, metadata files, and a function-driven user_sim simulator, so the agent can solicit clarifications, retrieve knowledge, and recover from execution errors on its own, with no human in the loop. The a-interact flavor takes this furthest: it is an open-ended agentic setting where the model decides for itself when to question the user simulator, when to explore the database, and how many turns to take before submitting a final answer. That is the closest of all the flavors to how agents actually operate in the real world, which is why we picked it.
We chose mini-interact because it is based on SQLite rather than Postgres, so it was less hassle to run, and it is smaller than the full version.
Put against the alternatives, here is why those two choices won:
There are no official leaderboard results for mini-interact yet. There are for the Lite version, which uses similar tasks but on Postgres databases rather than SQLite. The best result on a-interact there is 36.33% , using Opus-4.6 with a simple agent wrapper provided by the benchmark authors. That is the number we compare against, with the SQLite-versus-Postgres caveat noted up front.
The next challenge is that a substantial share of the gold answers the benchmark provides are simply wrong. This has been pointed out several times, for example here and here .
This is a real problem. To improve both SLayer and the agent harness, we need to distinguish cases where the agent actually went wrong from false negatives, where the agent’s answer was correct but evaluated against a wrong reference answer.
-- Task: "Show the accuracy ratio and confidence levels for each registration:
-- registration ID, project ID, accuracy, error values, the calculated ratio (RAR),
-- and the confidence level it translates to."
-- Benchmark gold answer (the confidence-level branch):
CASE
WHEN rar > 1 . 5 AND sr . refmark LIKE '%Target%' THEN 'High Confidence'
WHEN rar BETWEEN 1 . 0 AND 1 . 5 THEN 'Medium Confidence'
ELSE 'Low Confidence'
END AS confidence_level
-- Why it is wrong: per KB #44, 'High Confidence' requires LogMethod to contain 'Target'.
-- The gold matches against refmark instead, which only ever holds numeric codes
-- ('40', '31', '25'), so the High-Confidence branch never fires and returns 0 rows.
-- Matching on logmethod ('Target-based', 'Hybrid', 'Feature-based') gives the intended
-- 50 High-Confidence rows.
So the first thing we built was an annotation agent . It went through each task and decided three things, given all the data the task made available to the agent:
whether the gold answer was the single correct answer,
whether it was compatible with the task at all,
and whether there were other answers that would be equally valid given the task statement.
These annotations are themselves probabilistic, so we report success rates two ways: implied by the original gold answers, and implied by the annotations. The annotated numbers are not gospel truth either, but they give a much better evaluation signal than relying on the original answers alone.
To separate the effect of the agent harness from the effect of using SLayer, we built two versions of every agent variant, as similar to each other as possible. One used SLayer to introspect the database and the associated knowledge base and to submit the answer. The other used raw SQL and the tools that shipped with the benchmark repo.
For the SLayer version, we used the standard SLayer schema ingestion to pre-populate the models, then deterministically created a field for each documented component of a JSON column in the original schema. The documentation for those was itself structured JSON, so the ingestion was trivial.
// Original schema: one documented JSON column, returns.return_details
// (the benchmark documents each leaf of the JSON as structured text)
{
"column_meaning" : "JSONB column grouping return reason, authorization, and shipping metadata." ,
"fields_meaning" : {
"fraud" : {
"risk_level" : "VARCHAR(50). Fraud risk level for the return. NULL means not assessed."
}
}
}
# SLayer field created deterministically from it (one per documented JSON path)
- name : return_details__fraud__risk_level
sql : JSON_EXTRACT(return_details, '$.fraud.risk_level')
type : TEXT
description : "Fraud risk level for the return. NULL means no fraud risk level assessed."
meta :
derived_from :
json_col : return_details
path : [ fraud , risk_level ]
Worth being clear about one thing here: for the benchmark, this entire ingestion ran fully autonomously and deterministically, with no human or LLM in the loop. That was deliberate, to keep the comparison clean and the result reproducible. In real use it is not how SLayer is meant to work. Human input at this stage to encode the knowledge base items would be massively additive. So treat the numbers below as a floor for what an automated setup gets you .
We also ingested the knowledge base items, raw pieces of text that document relevant concepts and can refer to each other, as SLayer memories . The way the knowledge base items cross-referenced each other matched the memories structure perfectly.
For the agent harness, we tried two setups:
PydanticAI with a custom sub-agent structure, spawning exploration and definition sub-agents for every concept needed.
Claude SDK with the appropriate set of tools.
We used Opus-4.7 throughout for the agent, and Sonnet for the user_sim .
What the benchmark taught SLayer
Because the benchmark contains such a varied set of queries, running it against the SLayer version was a great workout for SLayer itself.
It triggered a couple of edge-case bugs. It pointed out some types of queries the DSL could not yet represent at all, which we then added. And it gave us a collection of almost-but-not-quite-correct tool calls the agent made, which we used to enlarge our set of patterns that we auto-correct for during query parsing.
We are taking a unique approach in designing tools for SLayer: we account for potential failures by the agent and self-correct them when the intent is obvious. You can learn more in our earlier post, How forgiving are your MCP tools?
# What the agent wrote (SQL-style aggregation -- appeared 265 times across the run):
count(*)
sum(revenue)
# What SLayer now auto-corrects it to (its native colon syntax):
*:count
revenue:sum
This is the part that matters beyond the score. A benchmark with this much variety surfaces gaps a curated test suite never would.
This is where we could boast of SLayer improving on the current leading model result by more than 2X: 75.3% with SLayer against the official 36.33% top-of-benchmark result.
But the more honest takeaway is that the choice of agent harness made the bigger impact on the pass rate .
We started with custom agents built using PydanticAI (basic loops with tools, where the tools included spawning task-specific sub-agents recursively). However, using claude_sdk yielded much higher pass rates (43% for PydanticAI-based vs 64% for claude_sdk, for a 53-task set we were using for that comparison), so we switched to claude_sdk throughout. Unfortunately, a comprehensive comparison of the two across the whole task set was not done for cost reasons.
The Claude SDK agents were a different story, with a significant improvement in performance overall. There, the SLayer-using agent performed marginally better than its raw-SQL twin.
So the Claude SDK with SLayer reached a 75.3% pass rate against the original gold answers, and 83.7% against the annotated ones.
To its credit, BIRD-INTERACT does build ambiguity in: that is the point of the interactive setting, where the agent questions a user simulator to pin down intent. But it resolves that ambiguity once, per task, against a single canonical answer. Business questions are different. The same question comes back month after month, phrased differently

[truncated]
