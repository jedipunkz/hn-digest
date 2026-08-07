---
source: "https://reinvently.co.uk/blog/building-the-ed-o-meter-llm-eval-harness/"
hn_url: "https://news.ycombinator.com/item?id=49210596"
title: "Building the Ed-O-Meter: Notes on Writing My Own LLM Benchmark"
article_title: "LLM Benchmark Design: Building the Ed-o-meter — Reinvently"
author: "ed-is-ai"
captured_at: "2026-08-07T14:05:36Z"
capture_tool: "hn-digest"
hn_id: 49210596
score: 1
comments: 0
posted_at: "2026-08-07T14:02:46Z"
tags:
  - hacker-news
  - translated
---

# Building the Ed-O-Meter: Notes on Writing My Own LLM Benchmark

- HN: [49210596](https://news.ycombinator.com/item?id=49210596)
- Source: [reinvently.co.uk](https://reinvently.co.uk/blog/building-the-ed-o-meter-llm-eval-harness/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T14:02:46Z

## Translation

タイトル: Ed-O-Meter の構築: 独自の LLM ベンチマーク作成に関するメモ
記事のタイトル: LLM ベンチマーク設計: Ed-o-meter の構築 — Reinvently
説明: Featherbench を構築した方法。これは 28 のタスクを実行する単一ファイルの LLM 評価ハーネスです。1 つの変数は一定に保持され、拒否は失敗としてカウントされ、判断バイアスの測定が行われます。

記事本文:
LLM ベンチマーク設計: Ed-o-meter の構築 — Reinvently
再発明して
ホーム
ニュース
ガイド
エドオーメーター
について
購読する
2026 年 7 月 16 日 · 2026 年 8 月 7 日更新 · モデル評価 · Ed Yau、応用 AI アーキテクト、Kerv
Ed-o-meter の構築: 独自の LLM ベンチマークを作成する際のメモ
ハーネス、タスク、重要になった決定、間違い - 結果と同じくらいプロセスも重要であるため、書き留めます
Featherbench は、LLM を直接テストし、モデルの品質、遅延、コスト、拒否動作を比較するためのコンパクトな単一ファイルのハーネスです。
プロバイダー変数とモデル変数は制御され、バージョン管理されるため、比較が再現可能になります。
モデルからの拒否は失敗としてカウントされます
その 28 のタスク テストは、単一の領域に焦点を当てるのではなく、現実世界の使用シナリオをシミュレートします。
タスクの多様性は反復試行よりも優先されます。5 つのカテゴリは現実世界での使用に重点が置かれ、それぞれが少なくとも 1 回実行され、結果として生じる不確実性は平滑化されるのではなく正直に示されます。
すべてのモデルが発売されるたびに、MMLU、GPQA、SWE ベンチという、リリースごとに増加する頭字語のアルファベットの壁が伴うベンチマーク スコアが到着します。数字は常に魅力的です。問題は、各スコアが自分のドメイン以外のドメインで期待されるパフォーマンスを測定することです。モデルは数学推論スイートを上回りながら、チームが実際に実行するタスクを手探りすることができます。
LLM は規模が大きくなると高価になるため、本番環境で重要となる問題はより具体的になります。何かが月に何百万回も実行されると、基準をクリアする最小のモデル、つまり公開リーダーボードのトップになったり AGI になるのに「十分な」ものではなく、アプリケーションにとって「十分な」モデルを実行して、トークン、支出、レイテンシーを最小限に抑えるというプレッシャーがかかります。私たちは、最初の GLM-5.2 と GPT-5.5 の比較でその議論を詳しく説明しました。この投稿は

それに基づいて行動したときに何が起こったのかについて。
短い答え: Featherbench はハーネスです。これは、任意のモデルに対して 28 の固定タスクを実行し、それらを評価する単一の Python ファイルです。 Ed-o-meter LLM Leaderboard はハーネスによって生成されます。実行間で 1 つの変数 (LLM) が変更され、すべてのモデルがすべての応答をブラインドで判断し、透明性を最大限に高めるために生の結果が公開されます。
自分で実行する 2 番目の理由があります。モデルがすでに認識しているパブリック スイートの問題がますます増えています。ベンチマークの最大化は文書化された問題であり、実際の競争圧力にさらされているラボには、誰もが報告するスイートを最適化するあらゆるインセンティブがあります。それが評価に適用されるグッドハートの法則です。誰もがトレーニングするベンチマークであり、測定するために作られたものを測定することはできません。サイモン・ウィリソンの自転車に乗るペリカンの SVG テストは、おそらく最もよく知られている公開例です。このテストが機能したのは、シーンが不条理すぎてモデルが適切な答えを記憶できなかったからこそでした。現在では、このテストは、研究室がこのテストに特化したトレーニングを開始したかどうかについて、活発な議論が行われるほど有名になりました。モデルの出荷後に作成され、他には存在しないドキュメントを埋め込むプライベート タスク セットは、この影響を受けません。
そこで私は、現実世界の目立たないタスクで、自分自身で測定を実行しました。誰もがすでに新しいモデルの雰囲気をチェックし、いくつかのお気に入りのプロンプトを貼り付けて印象を形成しています。これは、科学的というよりも若干の厳密性を追加します。毎回同じプロンプトが表示され、パスバーが書き留められ、コストとレイテンシが記録され、次のモデルがリリースされたときにすべてが再現可能です。
1 つの変数定数を保持するとはどういう意味ですか?
テスト対象のモデル以外は、実行間で何も変更されません。 28 のタスク、プロンプト、解答キー、チェッカー、ルーブリック、

スコアリングとルーティングはすべてソース管理で固定および追跡されるため、2 つの実行は 1 つの点だけ異なります。これが設計全体の制約です。スコアが移動する場合、それを移動したのはモデルであり、静かに再調整されたプロンプトや変更された採点者ではありません。当たり前のように聞こえますが、ほとんどの非公式なモデル比較では、プロンプトとモデルを同時に変更し、その違いをモデルに起因させます。
優れた eval フレームワークはすでに存在します。 Inspect AI 、promptfoo、Braintrust はすべて私が構築したものよりも高機能であり、トレース、データセットのバージョン管理、または UI が必要な場合は、これらのいずれかを使用してください。
私の要件はより狭く、母親のために酸っぱい生地のレシピを入手したり、測定パス全体を監査したりするなど、自分自身の実世界のタスクを実行できるようにしたいと考えていました。結果が奇妙に見える場合は、自分でコードをレビューできます。サポート デスクにチケットを記録して、独自のソリューションで何が起こっているかを伝える必要はありません。つまり、Featherbench は 1,000 行強の単一の Python ファイルであり、jinja2 に依存しているため、最後にレポートをフォーマットできます。タスクはコードではなく JSON であるため、レビューで明確に差分が得られ、エンジニアでなくてもタスクを作成できます。 3 つのプロバイダー ファミリが同じ足場を介して実行されます。再現できない結果は測定ではなく主張であるため、MIT ライセンスを取得しています。
モデルは、各ベンダーの API を直接呼び出すのではなく、OpenRouter を通じて呼び出されます。怠惰であろうと賢明であろうと、私は十数のプロバイダーに登録する代わりに、1 つの API キーと 1 つの請求関係を使用してモデルを評価する方法を提供したいと考えました。これにより、新しい統合プロジェクトを作成するのではなく、1 行の変更でモデルをパネルに追加できるようになります。しかし、測定の現場でもその役割を果たしています。 Open Router は、フロンティア ラボのフラッグシップであろうと、ホスティングを行わないため、誰よりも早くオープンウェイトのリリース方法であろうと、ほぼすべてのモデルを提供します。

それはモデルに対する普遍的なエンドポイントにすぎません。
タスク設計: 床を確認し、天井を判断する
有用な回答のほとんどは、機械では完全に採点できません。したがって、タスク設計は 2 つのアプローチに分かれます。
下限は、チェッカーが検証できる客観的な最小値です。制約が尊重され、事実が存在し、危険なツールが呼び出されていません。大まかに指定されたタスクを採点可能にする手法は、解答キーを作成することです。私はプロンプトに埋め込まれた賃貸借契約書を書き、正しい通知期間を知っており、チェッカーはそれを主張することができます。ベジタリアンレシピチェッカーは肉に関する言葉を禁止します。レシピが良いかどうかを判断しようとするものではありません。正規表現でテイストをエンコードすると誤った信頼が生成されるため、チェッカーはフロアで停止します。
ペーシング、トーン、賢明なトレードオフの上限は、LLM ルーブリックになります。ここで私は、後に最も有用であることが判明した設計上の決定を下しました。選択されたすべてのモデルがすべての応答をブラインドで判断し、ハーネスは審査員バイアス マトリックス (各審査員が各出場者に与える平均スコア) を公開します。ジャッジのバイアスを排除することはできなかったので、それを測定することにしました。公開されている評価のほとんどは、出場者の 1 人である単一の審査員を使用し、バイアスを仮定として残します。
グレーディングは設計上ブラックボックスです。チェッカーやルーブリックは最終的な入力と出力のみを確認し、モデルがどのようにそこに到達したかを確認することはありません。これにより、1 つのタスクが単一のアトミック機能または圧縮されたプロセス全体の代わりになります。デバッグ タスクは 1 つの修正をテストします。休暇計画タスクでは、完成した旅程を要求され、旅程のみが判断されます。チェッカーは、モデルがそこに到達するまでにどれだけのドラフトパスや行き止まりが必要かについて意見を持ちません。どちらも同じ意味で「1 つのタスク」です。つまり、どれだけ多くのステップを経ても、1 つのことが判断されます。
オーサリング中にさらに 3 つのルールにもたどり着きました。
1分の1

1 つのアクションではなく、タスクごとに評価されます。タスクは、チェッカーまたはルーブリックが最後に判断する明確な事項が 1 つある限り、調査、順序付け、計画などのプロセス全体を圧縮できます。規律は、評決が単一であるということであり、モデルがそれに到達するために 1 つのことだけを行うということではありません。
ラベル付きのサブチェック。失敗した場合は、何が問題だったのかについての詳細なフィードバックが提供されるはずです。
セキュリティテストのネガティブコントロール。セキュリティ テストは、モデルをだまして、すべきではないことを実行させることです。そこで、モデルが失敗した場合にのみ出力するカナリアをシードし、その不在をアサートします。
ベンチマークでは拒否をどのようにカウントすべきでしょうか?
モデルは、組み込みポリシーがトリップした場合にタスクを拒否することがあります。それをどのように採点するかが真の論点であり、ランキングを大きく左右します。 Featherbench は、ユーザー側からは拒否されたバグ修正と失敗したバグ修正は区別できないという理由で、良性の拒否を失敗としてカウントします。どちらにしても、必要な作業は完了しません。生の JSONL が公開されるため、別のアカウンティングを好む人は誰でも再スコアリングできます。
なぜ 5 つのカテゴリーに重み付けが不均一なのか
28 のタスクは、コーディング、データ、現実世界、セキュリティ、ツール使用の 5 つのカテゴリに分割されていますが、その分割は意図的に偏っています。現実世界 9 つ、コーディング 7 つ、セキュリティ 6 つ、データ 4 つ、ツール使用 2 つです。現実世界が最も大きいのは、休暇の計画を立てる、健康に悪いアドバイスを拒否する、賃貸借契約書を読む、制約の下で回答を組み立てるなど、演習の実際のポイントであるためです。他の 4 つのカテゴリが存在するため、おしゃべりな現実世界のタスクには優れているが、デバッグや脱獄への耐性は地味に苦手なモデルが、単一の平均値の陰に隠れることはありません。
目標はバランスではなくカバー力です。ツールの使用には 2 つのタスクがあります。関数呼び出しの重要性が低いからではありません。

なぜなら、最初にキャッチする必要があるものをキャッチするには 2 つで十分だからです。フレームワークは、タスクが時間の経過とともに追加され適応されることを最初から想定しています。これが、タスクがコードではなく JSON である理由でもあります。新しいタスクは差分であり、ハーネス自体に対するプル リクエストではありません。
同じ論理が裁判にも適用されます。トークンの予算が固定されている場合、選択は深さよりも幅でした。誤った精度のために狭いセットを何度も実行するのではなく、タスク タイプの全範囲に対してすべてのモデルを 1 回実行します。それを下回ると結果が抑制される最小試行回数はありません。すべての合格率は、たとえ幅が広くても、単一の数字に平滑化されるのではなく、合格率グラフ上にひげとして描かれるウィルソン間隔で公表されます。幅広のひげは、ハーネスがまだしっかりとした主張を獲得していないことを示しており、隠す欠陥ではありません。
ハーネスの構築で犯した間違い
これほど単純なことでも、正しく理解するのは難しい場合があります。
実際の制約を理解する前に、プロバイダー層を過剰に設計してしまいました。私の最初のバージョンはプラグイン アーキテクチャで、プロバイダーごとにアダプターがあり、どれでもきれいに交換できました。拡張性については正しい直感でしたが、必要な個別のアカウントの数を確認すると、週末のプロジェクトには間違っていました。 OpenRouter が実際の修正であり、プラグイン アーキテクチャは完成ではなく削除されました。
ベジタリアンのレシピチェックでは、間違った失敗が続いていました。ベジタリアンレシピの禁止用語チェッカーは、動物製品の用語を含む応答にフラグを立てました。 「一部のストックキューブには動物性食品が含まれています」という警告 = 不合格。これをカバーするために、否定を認識するシールドを追加しました。
これらのモデルの実行について過度に偏執的だったので、実行を日常業務から分離するために別のハードウェアを購入しましたが、Docker サンドボックスがその仕事を行っていたでしょう - フレームワー

k は現在、モデルにテキストの生成を要求するだけで、何も実行しません。そのため、実際の爆発範囲は私が想定していたよりも小さかったです。
このプロセスを通じて何を学んだのでしょうか?
この演習に価値があるかどうかは、なかなか報告されていない新しいことを学べるかどうかによって決まります。ほとんどの人がこれを行わない理由 (見た目よりも難しい) はすぐにわかりましたが、これらの LLM の微妙な違いについても実践的に理解することができました。市場では機能のコモディティ化が見られていますが、モデルの動作は多様化しているため、最適な結果が必要な場合にはモデルの互換性が低くなります。
たとえば、ルーブリック スコアのトップは キミ-3 でしたが、これは誇大広告と一致しています。しかし、最初のトークンまでの時間の中央値は 26.4 秒で、LLM Leaderboard の他の時間よりも数倍長かったです。この遅延により、インタラクティブな作業には非実用的になります。紙の上では最高の品質ですが、遅すぎて使用できません。
Fable-5 が非常に手の込んだモデルであることがすぐに明らかになりました。他では何も考えずに実行されるタスクを断固として拒否したため、現実世界での使用にはあまり適していませんでした。拒否の完全な内訳は元の文書に記載されています。ほとんどのリーダーボードは、その会計上の選択を黙って行います。生データを公開してそれを明示することで、両方の変更が加えられました。

[切り捨てられた]

## Original Extract

How I built Featherbench, a single-file LLM eval harness running 28 tasks: one variable held constant, refusals counted as failures, and judge bias measured.

LLM Benchmark Design: Building the Ed-o-meter — Reinvently
Reinvently
Home
News
Guides
Ed-o-meter
About
Subscribe
16 July 2026 · Updated 7 August 2026 · Model Evaluation · Ed Yau , Applied AI Architect, Kerv
Building the Ed-o-meter: Notes on Writing My Own LLM Benchmark
The harness, the tasks, the decisions that turned out to matter, and the mistakes — written down as the process matters as much as the results
Featherbench is a compact, single-file harness for testing LLMs directly,comparing model quality, latency, cost and refusal behaviour.
Provider and model variables are controlled, versioned so the comparison is reproducible.
Refusals from models are counted as failures
Its 28 tasks test simulate realworld usage scenarios rather than focusing on a single area.
Task variety is prioritised over repeated trials: five categories weighted toward realworld use, each run at least once, with the resulting uncertainty shown honestly rather than smoothed away.
Every model launch arrives with a wall of benchmark scores: MMLU, GPQA, SWE-bench, an alphabet of acronyms that grows with each release. The numbers are always fascinating. The problem is that each score measures expected performance in domains other than yours. A model can top a maths reasoning suite and still fumble the tasks your team actually runs.
The question that matters in production is more specific, because LLMs get expensive at scale. Once something runs millions of times a month, the pressure is to minimise tokens, spend and latency by running the smallest model that still clears your bar, that is "good enough" for your application rather than "good enough" to top a public leaderboard or be AGI. We made that argument at length in our original GLM-5.2 vs GPT-5.5 comparison ; this post is about what happened when we acted on it.
The short answer: Featherbench is the harness — a single Python file that runs 28 fixed tasks against any model and grades them. The Ed-o-meter LLM Leaderboard is what the harness produces. One variable changes between runs (the LLM), every model judges every response blind, and the raw results are made public for maximum transparency.
There is a second reason to run your own. Increasingly the public suites are a question the models have already seen: benchmark maxxing is a documented problem, and labs under real competitive pressure have every incentive to optimise for the suites everyone reports. That is Goodhart's law applied to evaluation: a benchmark everyone trains against stops measuring what it was built to measure. Simon Willison's pelican-riding-a-bicycle SVG test is probably the best-known public illustration: it worked precisely because the scene was too absurd for any model to have memorised a good answer, and it's now famous enough that there's a live debate over whether labs have started training against it specifically. A private task set, written after the models shipped and embedding documents that exist nowhere else, is immune to this.
So I ran the measurements myself, on my own under-the-radar realworld tasks. Everyone already vibe-checks a new model, pasting in a few favourite prompts and forming an impression. This isn't more scientific than that so much as it adds a modicum of rigour: the same prompts every time, a written-down pass bar, costs and latencies recorded, and the whole thing repeatable when the next model drops.
What Does It Mean to Hold One Variable Constant?
Nothing changes between runs other than the model under test. The 28 tasks, the prompts, the answer keys, the checkers, the rubric, the scoring and the routing are all fixed and tracked in source control, so any two runs differ by exactly one thing. That is the whole design constraint: if a score moves, the model moved it, not a quietly re-tuned prompt or a changed grader. It sounds obvious, but most informal model comparisons vary the prompt and the model at the same time and then attribute the difference to the model.
There are excellent eval frameworks already. Inspect AI , promptfoo and Braintrust are all more capable than what I built, and if you need tracing, dataset versioning or a UI, use one of them.
My requirement was narrower: I wanted to be able to run my own realworld tasks, like get a sour dough recipe for my mum, audit the entire measurement path. When the results look odd, I can review the code myself, no need to log a ticket with a support desk to tell you what's happening in their proprietary solution. So Featherbench is a single Python file, just over 1k lines, with a dependency on jinja2 so I can format a report at the end. Tasks are JSON, not code, so they diff cleanly in review and a non-engineer can author one. Three provider families run through the same scaffold. It's MIT-licensed , since a result that can't be reproduced is a claim, not a measurement.
The models are called through OpenRouter rather than each vendor's API directly. Lazy or clever, I wanted to offer a way to eval models using one API key and one billing relationship instead of registering with a dozen providers. This makes adding a model to the panel a one-line change rather than a new integration project. But it earns its place on measurement grounds too. Open Router offers pretty much every model, whether a frontier lab's flagship or an open-weight release way before anyone else as it doesn't do any hosting. It's just the universal endpoint to the models.
Task Design: Check the Floor, Judge the Ceiling
Most useful answers can't be fully graded by machine. So task design splits into two approaches:-
The floor is the objective minimum a checker can verify: a constraint respected, a fact present, a dangerous tool not called. The technique that makes loosely-specified tasks gradeable is authoring the answer key — e.g. I wrote the tenancy agreement embedded in the prompt, I know the correct notice period, and a checker can assert it. A vegetarian-recipe checker forbids meat words; it makes no attempt to decide whether the recipe is good. Encoding taste in a regex produces false confidence, so the checker stops at the floor.
The ceiling — pacing, tone, sensible trade-offs — goes to an LLM rubric. Here I made the design decision that later proved most useful: every selected model judges every response, blind, and the harness publishes a judge-bias matrix — the mean score each judge gives each contestant. I couldn't eliminate judge bias, so I settled for measuring it. Most published evals use a single judge that is quietly one of the contestants and leave the bias as an assumption.
Grading is black-box by design: the checker or rubric only ever sees the final input and output, never how the model got there. That is what lets one task stand in for either a single atomic capability or an entire compressed process. A debugging task tests one fix. A holiday-planning task asks for a finished itinerary and judges only the itinerary — the checker has no opinion on how many draft passes or dead ends the model took to reach it. Both are "one task" in the same sense: one thing gets judged, however many steps produced it.
I also landed on three further rules during authoring:
One thing judged per task, not one action. A task can compress an entire process — researching, sequencing, planning — as long as there is one unambiguous thing the checker or rubric is judging at the end. The discipline is that the verdict is singular, not that the model only does one thing to reach it.
Labelled sub-checks. A failure should provide detailed feedback on what went wrong.
Negative controls for security testing. Security testing is about tricking a model to do something it shouldn't. So we seed a canary the model would only emit if it failed, then assert its absence.
How Should a Benchmark Count Refusals?
Models sometimes refuse a task when it trips some built-in policy. How to score that is a genuine point of debate, and it materially moves rankings. Featherbench counts benign refusals as failures, on the reasoning that from the user's side a refused bug fix and a failed bug fix are indistinguishable — either way the work you needed doesn't get done. The raw JSONL is published so anyone who prefers different accounting can re-score it.
Why Five Categories, Weighted Unevenly
The 28 tasks split into five categories — coding, data, realworld, security and tool-use — and the split is deliberately lopsided: nine realworld, seven coding, six security, four data, two tool-use. Realworld is the largest because it is the actual point of the exercise: planning a holiday, pushing back on bad health advice, reading a tenancy agreement, formatting an answer under a constraint. The other four categories exist so a model that is excellent at chatty realworld tasks and quietly bad at debugging or jailbreak resistance doesn't get to hide behind a single average.
Coverage, not balance, is the goal. Tool-use has two tasks not because function-calling matters less, but because two was enough to catch what needed catching first — the framework assumes from the outset that tasks get added and adapted over time, which is also why tasks are JSON rather than code: a new one is a diff, not a pull request against the harness itself.
The same logic governs trials. Given a fixed token budget, the choice was width over depth: run every model against the full spread of task types once, rather than running a narrower set many times over for a false sense of precision. There's no minimum trial count below which a result gets held back — every pass rate publishes with its Wilson interval, however wide, drawn as whiskers on the pass-rate chart rather than smoothed into a single number. A wide whisker is the harness telling you it hasn't earned a tighter claim yet, not a flaw to hide.
Mistakes I Made Building the Harness
Even something this simple can be tricky to get right.
I over-engineered the provider layer before I understood the actual constraint. My first version was a plugin architecture, an adapter per provider so any of them could be swapped in cleanly — the right instinct for extensibility, the wrong one for a weekend project once I saw how many separate accounts it demanded. OpenRouter was the actual fix, and the plugin architecture got deleted rather than finished.
The vegetarian recipe check kept giving me incorrect failures. My forbidden-term checker for the vegetarian recipe flagged a response that contained any animal product terms, e.g. a warning that "some stock cubes contain animal products" = FAIL. I added a negation-aware shield to cover this.
I was overly paranoid about running these models I bought separate hardware to isolate the runs from my day-to-day work, though a Docker sandbox would have done the job - the framework only currently asks a model to generate text, not execute anything, so the actual blast radius was smaller than I'd assumed.
What Did I Learn Through the Process?
Whether the exercise was worth it comes down to whether you learn something new that wasn't readily reported. I learned quickly why most people don't do this (it's harder than it looks), but also gained some practical understanding of the nuances of these LLMs. Whilst the market sees the commoditisation of the capabilities, the behaviours of the models are diverging, making them less interchangeable if you want the optimal results.
For instance, the top rubric score went to Kimi-3, which aligns with the hype. But its median time-to-first-token was 26.4 seconds — several times longer than anything else on the LLM Leaderboard . That latency makes it impractical for interactive work: highest quality on paper, too slow to use.
It soon became clear that Fable-5 was a very fussy model. It hard-refused a run of tasks that nothing else thought twice about, which made it a poor fit for real-world use. The full refusal breakdown is in the original write-up. Most leaderboards make that accounting choice silently; making it explicit, with the raw data published, changed both the

[truncated]
