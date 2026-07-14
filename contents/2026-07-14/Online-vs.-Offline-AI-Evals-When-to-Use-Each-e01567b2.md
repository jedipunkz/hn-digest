---
source: "https://www.inngest.com/blog/online-vs-offline-ai-evals-when-to-use-each"
hn_url: "https://news.ycombinator.com/item?id=48913338"
title: "Online vs. Offline AI Evals: When to Use Each"
article_title: "Online vs Offline AI Evals: When to Use Each - Inngest Blog"
author: "aldersondev"
captured_at: "2026-07-14T22:47:21Z"
capture_tool: "hn-digest"
hn_id: 48913338
score: 6
comments: 1
posted_at: "2026-07-14T21:47:42Z"
tags:
  - hacker-news
  - translated
---

# Online vs. Offline AI Evals: When to Use Each

- HN: [48913338](https://news.ycombinator.com/item?id=48913338)
- Source: [www.inngest.com](https://www.inngest.com/blog/online-vs-offline-ai-evals-when-to-use-each)
- Score: 6
- Comments: 1
- Posted: 2026-07-14T21:47:42Z

## Translation

タイトル: オンライン AI とオフライン AI の評価: それぞれをいつ使用するか
記事のタイトル: オンラインとオフラインの AI 評価: それぞれをいつ使用するか - Ingest Blog
説明: オフライン評価とオンライン評価の仕組み、実行コスト、実稼働環境の AI エージェントにそれぞれをいつ使用するか。

記事本文:
オンラインとオフラインの AI 評価: それぞれをいつ使用するか - インジェスト ブログ プラットフォーム
6.9K オープンソース 6.9K サインイン 無料で始める ブログ記事 オンラインとオフラインの AI 評価: それぞれをいつ使用するか
ミッチェル・アルダーソン • 2026/7/14 • 9 分で読む
AI エージェントは物事を破壊することに優れています。それらは非決定的です。同じ入力が異なる出力を生成する可能性があり、中には悪い結果を伴うものもあります。
評価は、AI システムのパフォーマンスを監視および測定する方法ですが、いつ、どのようにエージェントをテストして採点するのでしょうか?このガイドでは、エージェントが実稼働環境で想定どおりに動作することを確認するための 2 つのパターンを示します。
簡単に言えば、オフライン評価とオンライン評価では、いつ実行されるか、何を測定するかが異なります。
オフライン評価では、出荷前に固定データセットに対してエージェントのスコアが付けられます。AI の単体テストは、事前に選択した入力に対して CI で実行されます。
オンライン評価では、実稼働環境で発生する各インタラクションを、実際のトラフィックと実際の結果に対してスコア付けします。既知の一連のケースに対する事前展開ゲートとしてオフラインを使用します。オンラインを使用して、実際のユーザーがエージェントにアクセスしたときにエージェントが実際にどのように動作するかを確認します。ほとんどのチームは、リリース前に回帰を捕捉するためにオフラインで、オフラインでは近似することしかできない真実を測定するためにオンラインで、両方を実行しています。
良好な評価に共通する要素
私たちが「AI評価」と呼ぶものは、いくつかの部分から構成されています。これらを組み合わせることで、エージェントのパフォーマンスを完全に把握できます。
これは、発生時のデータのライブ ストリーム、または静的ファイルの場合があります。いずれにせよ、入力は最初の変数であり、すべての eval が実行される共通の真実の情報源です。
同じ入力を与えた場合、GPT と Claude のどちらがより良い結果を生み出しましたか?プロンプト A とプロンプト B はどうなるでしょうか?実験では、複数の理論を一度にテストしたり、理論をスキップしてデプロイされたエージェントだけをテストしたりできます。
得点はすべてがひとつになる場所です。変数の組み合わせ (in

put、モデル、プロンプト)、何をどのように測定したいかに基づいて、毎回同じ方法でそれらを測定するルーブリックを作成できます。
LLM-as-Judge: スコアリング関数は、エージェントの出力と評価プロンプトを使用して 2 番目の LLM を呼び出すことができます。非常に微妙で、より決定的なものに当てはめるのが難しい評価に使用します。
アルゴリズム: コードのみのスコアリング関数。トークンコスト、ワード数、出力生成速度などについてエージェントを測定します。
信号ベース: これは、人間のフィードバックを待つスコアリング関数、またはシステム内の他の場所からのイベントである可能性があります。たとえば、注文が正常に作成されたということは、ツール呼び出しが成功したことを示します。
スコアリングのルーブリックを定義することは、優れた評価システムの最も重要な部分です。エージェントの全体像を把握するには、いくつかの側面を捉えることを目指します。
これらのピースが組み合わさって 2 つのパターンになります。それぞれを詳しく説明する前に、次のように比較してください。
オフライン評価: ポイントインタイムの安心感
オフライン評価の重要な定義は静的データセットです。静的な入力、スコアラー、および分割テスト実験が与えられると、定義された評価が得られます。
導入前にエージェントへの変更をテストできるため、人気があります。単体テストを考えてみましょう。ただし AI の場合です。静的評価はローカルで実行することも、既知の入力に対する CI パイプラインの一部として実行することもできます。
欠点: テストの範囲は本質的に制限されています。静的データセット内の少数の例は、運用環境でエージェントが実際にヒットするものを示す弱いシグナルであり、テストを構築したときに思いつかなかったケースを見逃しがちです。
2番目のコストがかかります。誰かがその黄金のデータセットを構築し、それを代表的なものに保ち、製品が変更されるたびに再キュレーションする必要があります。
維持管理をスキップすると、セットは古くなります。グリーン スイートは、エージェントが数か月間示さなかった動作の認証を開始します。オフライン評価も可能

事前に選択した入力に対して単一ショットではないため、含めると考えたケースについては通知されますが、含めなかったケースについては何も表示されません。
オンライン評価: ライブデータを活用したスコアリング
オンライン評価では、異なるアプローチが採用されます。つまり、発生した各エージェントの対話をライブ データに基づいてスコア付けして評価します。
これには 2 つの大きな利点があります。まず、合成テストではなく、実際のデータに基づいてエージェントをスコアリングします。第 2 に、スコアの量が多くなり、信号がより正確になることを意味します。
これは、Ingest 内で eval を実行することが効果を発揮する場所であり、その理由は構造的なものです。 Ingest はすでに実行されており、エージェントを永続的に調整します。すべてのステップ、再試行、結果は実行時に保持されます。したがって、評価に必要なデータはすでに存在します。スコアを付けるということは、最初のシステムを監視するために 2 番目のシステムを立ち上げるのではなく、プラットフォームがすでに保持しているランを読み取ることを意味します。耐荷重部分である耐久性の高い機能はすでに一般提供されており、実証済みです。 eval は、新しい問題を指摘する同じ機械です。オフライン セットアップではできないことが 2 つあります。
ライブデータに対して分割テスト実験を実行できます。 GPT 対クロードの例では、一方のモデルが他方のモデルよりも優れたパフォーマンスを発揮すると、トラフィックを勝者にルーティングするのは簡単になります。
const gptTraffic = 50 ; const クロードトラフィック = 50 ; const {結果:brief、variant、experimentRef} = 待機グループ。 Experiment ( "research-agent-model" , {variants : { "gpt-5.5" : () => step . run ( "call-llm-synthesize-gpt-5.5" , () => synthesizeBrief ({ model : "gpt-5.5" ,evidence }) ), "claude-opus-4.8" : () => step . run ( "call-llm-synthesize-claude-opus-4.8" , () => synthesizeBrief ({ model : "claude-opus-4.8" ,evidence }) ), }, select : Experiment .weighted ({ "gpt-5.5" : gptTraffic , "claude-opus-4.8" : ClaudeTraffic , }), } );
実行層で eval を実行すると、製品シグナルを使用することもできます。

s スコア: 実際の結果であり、その代理ではありません。
コード エージェントが実行された後、ユーザーは回答を受け入れましたか? PR は統合されましたか?多くの場合、その結果はリクエスト時に到着しないため、遅延スコアリングが登場します。評価は実際の結果が到着するまで永続的に待機し、時間がかかる場合は数日後、それに対してスコアを付けます。
各結果はスコアリング機能に送信される単なるイベントであり、より従来のスコアラーと簡単に組み合わせることができます。以下の waitForEvent ステップは、プロセスを開いたままにせずに、最大 1 週間一時停止します。
import const ResearchSavedOutcomeScorer = createScorer ( inngest , { id : "research-saved-outcome-scorer" }, async ({event , step }) => { const Saved = await step .waitForEvent ( "wait-for-brief-save" , {event : "research/brief.saved" 、 if : "async.data.researchRunId == events.data.researchRunId" 、タイムアウト : "7d" 、}); return { 名前 : "research_saved" 、値 : 1 : 0 } );
オンライン評価の実行にはどのくらいの費用がかかりますか?
オンライン評価は、固定セットではなくライブ トラフィックをスコアリングするため無料ではありませんが、コスト プロファイルは見た目とは異なり、通常は重要な部分は小さくなります。
評価データは実行の副産物です。実行する 2 番目のプラットフォームも、外部ツールにエクスポートするトレースも、維持するためのスケジュールされたオフライン ジョブもありません。これにより、可動部品が減り、耐荷重システムが 2 つではなく 1 つになり、オフライン設定の実際の複雑さのコストが実際にかかるところになります。
正直な反論: オンライン スコアリングは実際のトラフィックに対して計算を実行し、デプロイ前のゲートを放棄します。ユーザーが見る前に新しいリリースが何をするかではなく、すでに何が起こったかを示します。そのためには、最初に変更を健全性チェックするためのオフラインまたは合成セットが必要です。 2 つのパターンは競合するものではありません。さまざまなリスクをカバーします。
両パターンのAI評価

イオンは、エージェントの品質テストのさまざまな部分に役立ちます。
エージェントをテストする静的な「ゴールデン」データセットがあります。
エージェントに変更をコミットする前に、展開前の健全性チェックを実行したいと考えています。
テストしたい合成ベンチマークはありますが、ライブ データはまだありません。
エージェントのライブパフォーマンスをリアルタイムで評価したいと考えています。
運用トラフィックに対してエージェントへの変更を分割テストまたはカナリアデプロイしたいと考えています。
ユーザーフィードバックやプラットフォームイベントなど、エージェントの出力に対するスコアリングメトリクスとして使用できるライブ製品シグナルがあります。
通常、eval は別個の製品です。つまり、エージェントを実行するシステムの横で実行される、トレースのエクスポート先となる別のプラットフォームです。インジェストではそうではありません。すでにエージェントを永続的に実行しているため、評価は同じシステム内で行われ、別のシステムを購入して接続する必要はありません。
Ingest でネイティブ AI 評価を備えた耐久性のあるエージェントを起動し、独自の運用トラフィックにスコアラーを向けます。測定できないものは改善できません。
オンライン AI 評価とオフライン AI 評価の違いは何ですか?
オフライン評価では、CI で実行される単体テストと同様に、デプロイ前に固定データセットに対してエージェントのスコアが付けられます。オンライン評価では、本番環境で発生する実際のインタラクションをスコア化します。
オフラインでは、リリース前に既知のリグレッションを検出します。 online は、実際のユーザーがエージェントにアクセスしたときにエージェントが実際にどのように動作するかを測定します。
オンライン評価とオフライン評価を併用できますか?
はい、ほとんどのチームはそうすべきです。オフラインは、複製できるケースの展開前のゲートとして機能します。オンラインでは、完全には予測できない現実が測定されます。これらは異なるリスクをカバーするため、両方を実行するとリリース前のチェックとリリース後のシグナルが得られます。
オンライン評価はオフラインよりも高価ですか?
本質的にはそうではありません。より大きなコスト レバーはスコアリング方法であり、評価が実行される場所ではありません。アルゴリズム

マイクと信号ベースのスコアラーはどちらにしても安価です。 LLM-as-judge は高価なものです。エージェントの呼び出しに加えて 2 番目のモデル呼び出しが追加されるため、オンラインでもオフラインでも、そのパスのトークン請求額がおよそ 2 倍になる可能性があります。バッチ処理とサンプリングにより元の状態に戻ります。どのスコアラーを選択するかが、法案が実行される場所よりもはるかに大きな影響を及ぼします。
2 番目のモデルがルーブリックまたはプロンプトに対してエージェントの出力を評価するスコアリング方法。これは、決定論的なコードとして表現するのが難しい微妙な判断に役立ち、データセットに対してオフラインで実行することも、ライブ出力に対してオンラインで実行することもできます。
遅延スコアリングでは、エージェントが応答した時点ではなく、実際の結果が判明した時点でインタラクションを評価します。マージされた PR や出荷された注文など、関心のある結果が数日後にしか解決されない場合、評価はそのイベントを待機して、そのイベントに対してスコアを付けます。
良好な評価に共通する要素
オフライン評価: ポイントインタイムの安心感
オンライン評価: ライブデータを活用したスコアリング
オンライン評価の実行にはどのくらいの費用がかかりますか?
2 番目のミドルウェアを追加すると、typescript タイプが壊れました
スタッキング ミドルウェアがどのようにして TypeScript 推論を静かに崩壊させたのか、そしてそれを修正するには何が必要だったのか。
エージェント評価の紹介: 実際の結果に基づいてエージェントをスコアリングします
実験、スコアリング、延期 - エージェントがすでに実行されている実行レイヤーでの結果ベースの評価。
スコアリングの導入: あらゆるランに判定を付加します
関数を書くのと同じくらい簡単な結果ベースのスコアリング。エクスポートするトレースも、個別のダッシュボードも、ID の調整も必要ありません。
数分でプロジェクトに Ingest を追加できます。無料で始められ、クレジットカードは必要ありません。
インジェストキューと従来のキュー
インジェストは SOC 2 Type II に準拠しています。
© 2026 Inngest Inc. 無断複写・転載を禁じます。
© 2026 Inngest Inc. 無断複写・転載を禁じます。

## Original Extract

How offline and online evals work, what they cost to run, and when to use each for AI agents in production.

Online vs Offline AI Evals: When to Use Each - Inngest Blog Platform
6.9K Open Source 6.9K Sign in Start free Blog Article Online vs Offline AI Evals: When to Use Each
Mitchell Alderson • 7/14/2026 • 9 min read
AI agents excel at breaking things. They're non-deterministic: the same input can produce different outputs, some with bad consequences.
Evaluations are how we monitor and measure the performance of AI systems, but how and when do you test and score your agent? In this guide we'll show the two patterns for checking that your agent works the way you think it does in production.
The short version: offline and online evals differ in when they run and what they measure.
An offline eval scores your agent against a fixed dataset before you ship: a unit test for AI, run in CI on inputs you choose in advance.
An online eval scores each interaction as it happens in production, against real traffic and real outcomes. Use offline as a pre-deploy gate on a known set of cases; use online to see how the agent actually behaves once real users hit it. Most teams run both, offline to catch regressions before release, online to measure the truth that offline can only approximate.
The Common Components of a Good Eval
What we call "AI evaluation" is made up of a few parts. Together they give a complete view of an agent's performance.
This can be a live stream of data as it happens, or a static file. Either way, the input is the first variable: the common source of truth every eval runs against.
Given the same input, did GPT or Claude produce the better result? What about prompt A vs. B? Experiments let you test multiple theories at once, or skip them and just test the deployed agent.
Scoring is where it all comes together. Given the combination of variables (the input, the model, the prompt), we can create a rubric that measures them the same way every time, based on what and how we care to measure:
LLM-as-Judge: A scoring function can call a second LLM with the agent output and an evaluation prompt. Use it for evaluations that are highly nuanced and hard to fit into something more deterministic.
Algorithmic: A code-only scoring function; measure the agent on token cost, number of words, speed of output generation, etc.
Signal-based: This could be a scoring function that waits for human feedback, or an event from somewhere else in the system; an order successfully created indicates a successful tool call, for example.
Defining your scoring rubrics is the most critical piece of a good eval system. Aim to capture a few dimensions to get the fullest view of your agent.
Those pieces combine into two patterns. Here's how they compare before we dig into each:
Offline Evals: Point-in-Time Peace of Mind
The key definition of an offline eval is a static dataset . Given static input, scorer, and any split-test experiment, we get a defined evaluation.
It's popular because you can test a change to the agent before deployment. Think unit tests, but for AI. You can run a static evaluation locally or as part of a CI pipeline on a known input.
The downside: the test's scope is inherently limited. A handful of examples in a static dataset is a weak signal for what the agent will actually hit in production, and it's easy to miss cases you didn't think of when you built the test.
There's a second cost. Someone has to build that golden dataset, keep it representative, and re-curate it every time the product changes.
Skip the upkeep and the set goes stale: a green suite starts certifying behavior the agent hasn't shown in months. Offline evals also run single-shot, on the inputs you chose in advance, so they tell you about the cases you thought to include and nothing about the ones you didn't.
Online Evals: Scoring Powered by Live Data
Online evals take a different approach: score and evaluate each agent interaction as it happens, on live data.
This has two big advantages. First, you're scoring the agent on real data, not synthesized tests. Second, the sheer volume of scores is higher, which means a more accurate signal.
This is where running evals inside Inngest pays off, and the reason is structural. Inngest already runs and orchestrates your agent durably: every step, retry, and outcome is persisted as it executes. So the data an eval needs is already there. Scoring it means reading runs the platform is already keeping, not standing up a second system to watch the first. The load-bearing part, durable execution, is already GA and proven; evals are that same machinery pointed at a new problem. Two things follow that offline setups can't do.
You can run split-test experiments against live data. In the GPT-vs-Claude example, once one model outperforms the other, it's trivial to route traffic to the winner.
const gptTraffic = 50 ; const ClaudeTraffic = 50 ; const { result : brief , variant , experimentRef } = await group . experiment ( "research-agent-model" , { variants : { "gpt-5.5" : () => step . run ( "call-llm-synthesize-gpt-5.5" , () => synthesizeBrief ({ model : "gpt-5.5" , evidence }) ), "claude-opus-4.8" : () => step . run ( "call-llm-synthesize-claude-opus-4.8" , () => synthesizeBrief ({ model : "claude-opus-4.8" , evidence }) ), }, select : experiment . weighted ({ "gpt-5.5" : gptTraffic , "claude-opus-4.8" : ClaudeTraffic , }), } );
Running evals in the execution layer also lets you use product signals as scores: the real outcome, not a proxy for it.
After the code agent ran, did the user accept the answer? Did the PR get merged? Often that outcome doesn't arrive at request time, which is where deferred scoring comes in: the eval durably waits for the real result to land, days later if that's how long it takes, then scores against it.
Each outcome is just an event sent to the scoring function, easily combined with more traditional scorers. The waitForEvent step below pauses for up to a week without holding a process open.
export const researchSavedOutcomeScorer = createScorer ( inngest , { id : "research-saved-outcome-scorer" }, async ({ event , step }) => { const saved = await step . waitForEvent ( "wait-for-brief-save" , { event : "research/brief.saved" , if : "async.data.researchRunId == event.data.researchRunId" , timeout : "7d" , }); return { name : "research_saved" , value : saved ? 1 : 0 }; } );
What Does It Cost to Run Online Evals?
Online evals aren't free, since you're scoring live traffic instead of a fixed set, but the cost profile is different from how it looks, and usually smaller where it counts.
The eval data is a byproduct of execution: no second platform to run, no traces to export to an external tool, no scheduled offline job to keep alive. That's fewer moving parts, one load-bearing system instead of two, which is where the real complexity cost of an offline setup actually sits.
The honest counterpoint: online scoring does run compute against real traffic, and it gives up the pre-deploy gate. It tells you what already happened, not what a new release will do before users see it. For that you still want some offline or synthetic set to sanity-check a change first. The two patterns aren't rivals; they cover different risks.
Both patterns of AI evaluation are useful for different parts of your agent quality testing.
You have a static "golden" dataset that you want to test your agent against.
You want to run a pre-deployment sanity check before committing changes to your agent.
You have a synthetic benchmark you want to test against, but no live data yet.
You want to evaluate your agent's live performance in real time.
You want to split-test or canary-deploy a change to your agent against production traffic.
You have live product signals you can use as scoring metrics against your agent's output, like user feedback or platform events.
Usually, evals are a separate product: another platform you export traces to, running beside the system that executes your agent. On Inngest they aren't. Because it already runs your agent durably, evaluation happens in the same system, not a second one to buy and wire up.
Spin up a durable agent with native AI evals on Inngest and point a scorer at your own production traffic. You can't improve what you can't measure!
What is the difference between online and offline AI evals?
Offline evals score an agent against a fixed dataset before you deploy, like a unit test run in CI. Online evals score real interactions as they happen in production.
Offline catches known regressions before release; online measures how the agent actually behaves once real users hit it.
Can you use online and offline evals together?
Yes, and most teams should. Offline acts as a pre-deploy gate on the cases you can replicate; online measures the reality you can't fully predict. They cover different risks, so running both gives you a check before release and a signal after it.
Are online evals more expensive than offline?
Not inherently: the bigger cost lever is your scoring method, not where the eval runs. Algorithmic and signal-based scorers are cheap either way. LLM-as-judge is the expensive one: it adds a second model call on top of the agent's, which can roughly double the token bill on that path, online or offline. Batching and sampling bring it back down. Which scorer you pick moves the bill far more than where it runs.
A scoring method where a second model evaluates the agent's output against a rubric or prompt. It's useful for nuanced judgments that are hard to express as deterministic code, and it can run offline on a dataset or online against live output.
Deferred scoring evaluates an interaction once its real outcome is known, rather than at the moment the agent responds. If the outcome you care about, like a merged PR or a shipped order, only resolves days later, the eval waits for that event and scores against it.
The Common Components of a Good Eval
Offline Evals: Point-in-Time Peace of Mind
Online Evals: Scoring Powered by Live Data
What Does It Cost to Run Online Evals?
Adding a second middleware broke our typescript types
How stacking middleware quietly collapsed our TypeScript inference—and what it took to fix it.
Introducing Agent Evals: Score your agents on real outcomes
Experiments, Scoring, and Defer—outcome-based evals in the execution layer your agents already run on.
Introducing Scoring: Attach a verdict to any run
Outcome-based scoring as easy as writing a function—no traces to export, no separate dashboards, no ID reconciliation.
Add Inngest to your project in minutes. Free to start, no credit card required.
Inngest vs. Traditional Queues
Inngest is SOC 2 Type II compliant.
© 2026 Inngest Inc. All rights reserved.
© 2026 Inngest Inc. All rights reserved.
