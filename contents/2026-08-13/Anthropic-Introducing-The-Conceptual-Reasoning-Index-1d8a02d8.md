---
source: "https://alignment.anthropic.com/2026/conceptual-reasoning-index/"
hn_url: "https://news.ycombinator.com/item?id=49285909"
title: "Anthropic: Introducing The Conceptual Reasoning Index"
article_title: "Introducing the Conceptual Reasoning Index"
author: "optimalsolver"
captured_at: "2026-08-13T14:16:03Z"
capture_tool: "hn-digest"
hn_id: 49285909
score: 4
comments: 1
posted_at: "2026-08-13T13:48:14Z"
tags:
  - hacker-news
  - translated
---

# Anthropic: Introducing The Conceptual Reasoning Index

- HN: [49285909](https://news.ycombinator.com/item?id=49285909)
- Source: [alignment.anthropic.com](https://alignment.anthropic.com/2026/conceptual-reasoning-index/)
- Score: 4
- Comments: 1
- Posted: 2026-08-13T13:48:14Z

## Translation

タイトル: Anthropic: 概念的推論インデックスの紹介
記事のタイトル: 概念的推論インデックスの紹介

記事本文:
アライメント科学ブログ
概念的推論インデックスの導入
AI リスク管理における中心的な期待は、AI が私たちの状況を理解し、今後の計画を立て、リスク軽減策を開発するのに役立つことです。この目的のために AI が実行しなければならない多くのタスクには、実践的な経験に基づくフィードバック ループが欠如しており、哲学、AI 未来論、および同様の分野で使用される種類の議論にモデルが関与する必要があります。これらの機能を評価するために、3 つの概念的推論ベンチマークのスイートを開発しました。このフォームを通じて、主要な概念的データセット LMCA へのアクセスをリクエストできます。
私たちはベンチマークを Conceptual Reasoning Index (CRI) に集約します。このインデックスは、conceptualreasoning.ai で入手できます。私たちの方法論の詳細もここでご覧いただけます。新しいモデルとベンチマークの両方がリリースされるたびに、Web サイトを最新の状態に保ちます。
この作品はAnthropicとの協力により行われました。
モデルが人間の専門家のレベルで AI のリスクを軽減する作業を実行できるようになると、その分野での AI (支援) の出力は、人間の支援を受けない出力よりも小さくなる可能性があります。これは、AI リスクに間に合うかどうかの主な決定要因は、高リスクの機能と比較して、この作業をどれだけ早く自動化または強化できるかであることを示唆しています。これに影響を与える 1 つの方法は、AI を管理および調整する方法や、AI が関与する壊滅的な協力の失敗を回避する方法についての推論など、モデルの関連スキルを選択的に向上させることかもしれません。
現在の AI トレーニングは、豊富なデータとモデルのパフォーマンスに関する信頼できるフィードバックに大きく依存しています。したがって、モデルは通常、経験的または数学的に検証できないタスクでは劣ります。 1 残念ながら、高度な AI によるリスクを軽減するには、次のような多くのタスクが必要です。
AI の安全性に関する作業の多くには、人間よりも一般的に能力が高い AI についての推論が含まれます。明らかなことはない

これについては参照クラスがあり、それをモデル化する明確な方法はありません。
最初はいくつかのことを正しく行う必要があるかもしれません。たとえば、間違いが AGI の乗っ取りや AI 支援によるクーデターにつながった場合、手遅れになるまで発見できない可能性があります。同様に、多くの決定 (どの研究課題を優先するか、どのガバナンス介入を追求するかなど) は長い時間スケールで展開されるため、経験的なフィードバックが十分に早く到着しない可能性があります。
最後に、AI がどのような価値観を持つべきかなど、いくつかの重要な疑問には、根本的な真実がまったく欠けている可能性があります (それでも、これらの疑問について議論することで進歩が得られると私たちは考えています)。
これらの特性を考慮すると、高度な AI によるリスクを軽減する取り組みは、経験的証拠が限られており、（実質的に）検証可能な答えがなく、したがって議論に大きく依存する必要がある場合に、質問について推論する能力の向上から特に恩恵を受ける可能性があります。これを概念的推論と呼びます。この機能を改善するには、それを測定できる必要があるため、 LMCA 、 ACCoRD 、および DTBench 機能の 3 つのベンチマークを構築しました。また、モデルの全体的な概念的推論機能を把握するために、これらのベンチマークの集合体である概念推論インデックス (CRI) も構築します。
LMCA (Language Model Conceptual Argumentation) は、意思決定理論、哲学、高度な AI によるリスクなど、さまざまなトピックに関する厳選され、専門家によって評価された概念的な議論のデータセットです。議論に焦点を当てることは、概念的な質問に対する最終的な答えを検証する困難を回避するのに役立ちます。
データセットには、560 個の位置テキストと、これらの位置テキストに対する 1,461 個の引数が含まれています。ほぼすべての 2 つの議論は概念研究者 Emery Cooper によって評価され、一部は少なくとも 1 人の他の研究者によって独立して評価され、合計 2,1

評価40。私たちは、モデルの評価を私たちの評価と比較することで、立場テキストに対する議論を判断する点でモデルがどの程度優れているかを測定します。
評価は詳細なルーブリックに従います。少なくとも 2 人によって評価された議論では、人間とモデルの間の一致と比較して、評価者間の一致が高くなります。これには、約 50 の議論からなる検証セットが含まれており、それぞれの議論が 4 ～ 6 人によって個別に評価され、合計 7 ～ 8 時間議論されます。
LMCA では、モデルの議論能力を評価することもできます。データセット内の位置テキストに、それに対する 3 つの評価された引数があるとします。ここで、モデル A に位置テキストに対する 4 番目の引数を生成するよう依頼できます。次に、モデル B にルーブリックを与え、数回のプロンプトで 3 つの既存の議論とその評価を提示し、モデル A の新しい議論を評価するように求めます。この方法では、モデル B からかなり正確な評価が生成されます。
現在、議論を判断する際のモデルのパフォーマンスのみが CRI に入力されますが、将来的にはモデルの議論能力の測定も追加したいと考えています。
ACCoRD (概念的推論ドメインの一貫性の評価) は、概念的な問題についてモデルが報告した信念と好みが論理的にどの程度一貫しているかを測定します。たとえば、あるモデルに確率 P(A) を求め、同じモデルの別のインスタンスに確率 P(A&B) を求めた場合、報告される確率は P(A) ≥ P(A&B) を満たしますか?データセット内のすべての一貫性制約は、モデルに確率の数値推定または優先順位のいずれかを要求します。
特定の質問セットに対する一貫性の欠如は、デフォルトではそのセットに対するモデルの推論を信頼できないことを示す良い指標です。同様に、一般的にモデルが概念的な問題に関して非常に一貫性がない場合、これは概念的な推論が欠如していることを示しています。
ACCoRD データセットには clo が含まれています

モデル生成の一貫性制約は 14,000 個あり、これらは 18 種類の制約タイプに分散され、自動チェッカー パイプラインを通過しています。このうち、567 件については当社がさらにチェックし、承認しました。これらの 567 個の制約のみを、概念的推論パフォーマンスの総合指標である CRI に含めます。
DTBench 機能 (意思決定理論ベンチマーク) は、モデル自身の動作や (近い) コピーとの相互作用の忠実な予測を含む意思決定理論的状況について推論するモデルの能力を測定するために設計された 407 個の手作りの多肢選択質問のデータセットです。質問の大部分はオリジナルであり、意思決定理論に関する著書がある Caspar Oesterheld によって作成されました。すべての質問は、別のドメイン専門家である Emery Cooper によって独立して検証されました。
完全な DTBench スイートには、モデルの意思決定理論的な態度を測定する追加の 130 の質問が含まれています。これらは CRI には含まれていません。
以下のグラフは、2026 年 8 月 10 日の時点で、Anthropic の最高のモデルと、当社が評価した他の AI 企業の最高スコアモデルの CRI スコアを示しています。また、Claude Fable 5、Muse Spark 1.2、および Gemini 3.6 Flash のスコアも含まれています。これらは、CRI には含まれていませんが、多くの外部ベンチマークで各企業の最高パフォーマンスのモデルです。 CRI は現在、LMCA (60%)、ACCoRD (20%)、および DTBench 機能 (20%) の加重平均です。将来的には、インデックスに新しいベンチマークを追加し、飽和したベンチマークを廃止し、相対的なウェイトを調整する予定です。
スコアは 0 から 100 までで、0 はランダムな推測に対応し、100 はすべてのベンチマークにわたる最高スコアに対応します。 LMCA スコア 100 は、モデルが人間の評価を完全に再現したことを意味します。人間の評価はうるさいので、

最大限に優れた LMCA 評価を与えるモデルでは、スコアは 100 ではなくおよそ 85 になります。これは専門家の評価者間の合意に基づいて推定されます。一方、DTBench の機能に関するすべての質問に正解すると、100 点または 100 に非常に近いスコアが得られると予想されます。ACCoRD のスコア 100 は、完全に一貫していることを意味します。全体として、これにより、CRI の上限パフォーマンスは約 91 であると推定されます。最高スコアのモデルである Opus 5 は、スコアが 73.6 (95% CI: ± 2.1) で、この上限を依然として大幅に下回っています。
スコアは 2024 年後半からほぼ直線的に増加しており、横ばいの兆候はありません。
LMCA と ACCoRD の両方で最高スコアのモデルは、依然としてこれらのベンチマークの推定上限を大幅に下回っています。現在までのスコアから推定すると、LMCA は今から約 1 年後に飽和し始めると大まかに推定しています。一方、DTBench の機能スコアはすでに上限に近づいており、Fable 5 は問題の 98% を正解しています。 ACCoRD がいつ飽和するかについては非常に不確実です。
私たちは、高度な AI によるリスクを軽減するモデルの機能を向上させることが重要かつ緊急であると考えています。この研究の多くは概念的なものであり、モデルの概念的推論を改善することが特に価値がある可能性があることを示唆しています。この目的を達成するために、私たちは 3 つの概念的推論ベンチマークを開発し、CRI に集約しました。新しいベンチマークがリリースされると、CRI を更新します。
CRI の詳細とライブスコアについては、conceptreasoning.ai をご覧ください。
私たちの主要な概念的データセットである LMCA にアクセスするには、このフォームを送信してください。
METR、フロンティア リスク レポート (2026 年 2 月から 3 月) のセクション「エージェントの判断力と信頼性は人間の専門家よりも著しく悪かった」。対照的に、「MirrorCode の結果と一致して、公的に報告された逸話は

s は、エージェントが容易に「山を登る」問題、つまり進捗状況を確認するのにコストがかからず、多くのアプローチを安価に試せる問題に対して非常に強力であることを示しています。 ↩
非常に少数の議論 (1,461 件中 16 件) はエメリー・クーパーによって評価されませんでした。これらはすべて、概念研究者でもある Caspar Oesterheld によって評価されました。 ↩

## Original Extract

Alignment Science Blog
Introducing the Conceptual Reasoning Index
A core hope for managing AI risks is that AIs will help us understand our situation, plan for what lies ahead, and develop risk mitigations. Many tasks AIs would have to do for this purpose lack practical empirical feedback loops and require models to engage in the kinds of argumentation used in philosophy, AI futurism, and similar domains. To evaluate these capabilities, we develop a suite of three conceptual reasoning benchmarks. You can request access to our primary conceptual dataset, LMCA, through this form .
We aggregate the benchmarks into the Conceptual Reasoning Index (CRI), available at conceptualreasoning.ai , where you can also find more details on our methodology. We will keep the website up to date as both new models and benchmarks are released.
This work was done in collaboration with Anthropic.
Once models can perform work that reduces AI risk at the level of human experts, AI(-assisted) output in the area might dwarf unassisted human output. This suggests that a major determinant of whether we address AI risks in time is how early we can automate or uplift this work, relative to high-risk capabilities. One way to influence this might be to selectively improve models' relevant skills, such as reasoning about how to govern and align AI and how to avoid catastrophic cooperation failures involving AI.
Current AI training depends heavily on abundant data and reliable feedback on the model's performance. Models are therefore typically worse at tasks that cannot be empirically or mathematically verified. 1 Unfortunately, reducing risks from advanced AI involves many such tasks:
Much AI safety work involves reasoning about AIs more generally capable than any human. There's no obvious reference class for this and no clear way to model it.
We might have to get some things right the first time. For example, if a mistake leads to AGI takeover or an AI-assisted coup, we might not find out until it's too late. Similarly, many decisions (e.g., which research agendas to prioritize, which governance interventions to pursue) play out over long timescales, such that empirical feedback might not arrive early enough to help.
Lastly, some important questions, such as which values AIs should have, may lack a ground truth entirely (yet we still think progress can be made by arguing about these questions).
Given these properties, efforts to reduce risk from advanced AI may particularly benefit from an improved ability to reason about questions where empirical evidence is limited, there is no (practically) verifiable answer, and one therefore has to rely heavily on argumentation. We refer to this as conceptual reasoning. Improving this capability requires being able to measure it, so we built three benchmarks: LMCA , ACCoRD , and DTBench capabilities . We also construct an aggregate of these benchmarks, the Conceptual Reasoning Index (CRI), to give a sense of models' overall conceptual reasoning capabilities.
LMCA (Language Model Conceptual Argumentation) is a dataset of curated and expert-rated conceptual arguments on a diverse range of topics, including decision theory, philosophy, and risks from advanced AI. Focusing on arguments helps sidestep the difficulty of verifying bottom-line answers to conceptual questions.
The dataset contains 560 position texts with 1,461 arguments against these position texts. Nearly all 2 arguments were rated by conceptual researcher Emery Cooper, and some were independently rated by at least one other researcher, for a total of 2,140 ratings. We measure how good models are at judging arguments against position texts by comparing their ratings to ours.
Ratings follow a detailed rubric. On arguments rated by at least two people, inter-rater agreement is high compared to agreement between humans and models. This includes a validation set of roughly 50 arguments, each rated independently by 4–6 people and then discussed for 7–8 hours total.
LMCA also allows for evaluation of models' argumentation ability. Let's say a position text in our dataset has three rated arguments against it. Now, we can ask model A to generate a fourth argument against the position text. We then give model B the rubric and few-shot prompt it with the three existing arguments and their ratings, asking it to rate model A's new argument. This methodology produces fairly accurate ratings from model B.
Currently, only models' performance at judging arguments goes into the CRI, but we hope to add a measurement of models' argumentation ability in the future.
ACCoRD (Assessment of Consistency in Conceptual Reasoning Domains) measures the extent to which models' reported beliefs and preferences on conceptual issues are logically consistent. For example, if we ask a model for the probability P(A) and another instance of the same model for the probability P(A&B), do the reported probabilities satisfy P(A) ≥ P(A&B)? All consistency constraints in the dataset ask models for either numeric probability estimates or preference orderings.
Lack of consistency on a particular set of questions is a good indicator that we cannot, by default, trust a model's reasoning on that set. Similarly, if a model is generally very inconsistent on conceptual issues, this is a sign that its conceptual reasoning is lacking.
The ACCoRD dataset contains close to 14,000 model-generated consistency constraints, which are distributed across 18 constraint types and have gone through an automated checker pipeline. Of these, 567 were further checked and approved by us. We include only those 567 constraints in our aggregate conceptual reasoning performance metric, the CRI.
DTBench capabilities (Decision Theory Benchmark) is a dataset of 407 handcrafted multiple-choice questions designed to measure models' ability to reason about decision-theoretic situations that involve faithful predictions of a model's own behavior or interactions with (near) copies. The vast majority of questions are original and created by Caspar Oesterheld, who has published on decision theory. All questions were independently validated by Emery Cooper, another domain expert.
The full DTBench suite includes an additional 130 questions that measure models' decision-theoretic attitudes. We do not include these in the CRI.
The chart below shows the CRI scores of Anthropic's best models and the highest-scoring model from each other AI company we evaluated, as of August 10, 2026. We also include scores for Claude Fable 5, Muse Spark 1.2, and Gemini 3.6 Flash, which are their respective companies’ top-performing models on many external benchmarks, though not on the CRI. The CRI is currently a weighted average of LMCA (60%), ACCoRD (20%), and DTBench capabilities (20%). In the future, we plan to add new benchmarks to the index, retire saturated ones, and potentially adjust the relative weights.
Scores go from 0 to 100, with 0 corresponding to random guessing and 100 corresponding to the highest possible score across all benchmarks. An LMCA score of 100 would mean that the model perfectly replicated the human ratings. Because human ratings are noisy, we expect that a model giving maximally good LMCA ratings would score roughly 85 rather than 100, which we estimate based on expert inter-rater agreement. Meanwhile, we expect that giving the correct answer to every DTBench capabilities question would yield a score of 100 or extremely close to 100. A score of 100 on ACCoRD corresponds to being perfectly consistent. Overall, this leads us to estimate ceiling performance on the CRI to be around 91. The highest-scoring model, Opus 5, is still well below this ceiling, with a score of 73.6 (95% CI: ± 2.1).
Scores have been increasing roughly linearly since late 2024, with no signs of flattening.
The highest-scoring models on both LMCA and ACCoRD are still well below these benchmarks' estimated ceilings. Extrapolating from scores to date, we loosely estimate that LMCA will start saturating about a year from now. Meanwhile, DTBench capabilities scores are already close to the ceiling, with Fable 5 getting 98% of questions right. We're very uncertain about when ACCoRD will saturate.
We think improving models' ability to do work that mitigates risk from advanced AI is important and urgent. Much of this work is conceptual, suggesting that improving models' conceptual reasoning might be particularly valuable. To this end, we developed three conceptual reasoning benchmarks, which we aggregate in the CRI. We will update the CRI as new benchmarks are released.
For more information and live scores on the CRI, please visit conceptualreasoning.ai .
For access to LMCA, our primary conceptual dataset, please submit this form .
METR, Frontier Risk Report (February to March 2026) , section “Agents had significantly worse judgment and reliability than human experts”. By contrast: “Consistent with our MirrorCode results, publicly reported anecdotes indicate that agents were exceptionally strong on problems that are easily “hill-climbable” — those where progress is cheap to verify and many approaches can be tried cheaply”. ↩
A very small number of arguments (16 out of 1,461) were not rated by Emery Cooper. All of them were rated by Caspar Oesterheld, who is also a conceptual researcher. ↩
