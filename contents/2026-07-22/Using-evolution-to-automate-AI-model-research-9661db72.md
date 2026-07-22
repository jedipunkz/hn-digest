---
source: "https://imbue.com/blog/2026-07-20-imbue-catalyst-nanochat"
hn_url: "https://news.ycombinator.com/item?id=49010350"
title: "Using evolution to automate AI model research"
article_title: "Automating AI model research with evolution - Imbue"
author: "danielmewes"
captured_at: "2026-07-22T18:03:59Z"
capture_tool: "hn-digest"
hn_id: 49010350
score: 4
comments: 0
posted_at: "2026-07-22T17:31:05Z"
tags:
  - hacker-news
  - translated
---

# Using evolution to automate AI model research

- HN: [49010350](https://news.ycombinator.com/item?id=49010350)
- Source: [imbue.com](https://imbue.com/blog/2026-07-20-imbue-catalyst-nanochat)
- Score: 4
- Comments: 0
- Posted: 2026-07-22T17:31:05Z

## Translation

タイトル: 進化を利用して AI モデル研究を自動化する
記事のタイトル: 進化による AI モデル研究の自動化 - Imcue

記事本文:
進化による AI モデル研究の自動化 - Imbue
/ ブログ 製品 バウンサー
記事 / 研究 進化による AI モデル研究の自動化
解釈を進化させて行き詰まりから脱出する
実験効率と多様性のトレードオフ
付録: 再生産の詳細
計算研究と科学的発見のための AI ツールである Imbue Catalyst をオープンソース化しています
当社の進化ベースのオプティマイザーは、通常の AutoResearch エージェントよりも nanochat LLM のパフォーマンスを 3 倍向上させます
最近のモデルとハーネスの改良により、「自己改善」、つまり自らの基盤を繰り返し改善する AI システムの実用的なユースケースが可能になりました。
Darwin Godel Machine などのアプローチはシステムのハーネスの改善に焦点を当てていますが、他のアプローチはトレーニング レシピ、アーキテクチャの選択、モデル自体の背後にあるアルゴリズムを最適化しています (Andrej Karpathy の AutoResearch 、Recursive の「自動 AI 研究に向けた第一歩」)。これらの結果の多くは小さな「おもちゃ」モデルでのものですが、次世代の LLM が今日の AI システムの助けを借りてすでに構築されていることが明らかになりつつあります (例: Anthropic「AI が自ら構築するとき」)。
今日、私たちは Imbue Catalyst を発表します。 Catalyst は、進化にインスピレーションを得た手法を使用して、次の研究タスクを支援する研究ツールです。
特定のメトリクスに関してコード、アルゴリズム、またはモデルを最適化します。
特定のプログラム検証基準を満たすソリューションを見つける
計算で再現可能な現象の説明を発見する
計算研究分野における理論のレビュー、定式化、編集を支援する
私たちは、AI 研究はオープンで共同作業として行うのが最も効果的であると信じています。したがって、AGPL-3.0 ライセンスに基づいて Imbue Catalyst をリリースします。楽しみにしています

それによって何が発見されるのか！
この投稿では、Catalyst を使用して、ゼロからトレーニングされた小さなトランスフォーマー言語モデルである nanochat を最適化し、通常のコーディング エージェントが横ばいになった後も Catalyst が改善点を発見し続けることを可能にするメカニズムについて詳しく説明します。
Andrej Karpathy は最初に AutoResearch セットアップを提案しました。このセットアップでは、コーディング エージェントが小さなトランスフォーマー言語モデル ( nanochat ) のトレーニング セットアップを改善するように促されます。これを行うには、実験を繰り返し実行し、git を使用してその進行状況を追跡することが求められます。
セットアップは、単一の H100 GPU でちょうど 5 分間実行されるトレーニング スクリプトから始まります。エージェントは、5 分のトレーニング時間予算を超えない限り、検証データセットでの結果として得られる BPB (ビット/バイト) メトリクスを最小限に抑えるために、モデル アーキテクチャ、トレーニング レシピ、ハイパーパラメーターを自由に変更できます。
Nanochat トレーニングの最適化は、Catalyst で検証可能な目標と呼ばれるものの一例です。この文脈での検証可能とは、プログラムを使用して、候補ソリューションが目前の目標をどの程度達成しているかを測定および/または検証できることを意味します。この場合、検証者は結果のモデルの BPB メトリックを測定します。
Imbue Catalyst の検証可能なゴール ソルバー「Catalyst (Evolution)」を以下の基準点と比較します。
プレーンなオートリサーチ。 Weco AI によって報告された「3 回の実行の平均」データを示します。
Recursive SuperIntelligence が最近報告した結果。 2 つの注意点があります: Recursive は実行の途中で別の GPU (B200) に切り替えました。オリジナルの H100 GPU で実行された実行の最初の部分のみが含まれています。
Recursive はブログ投稿で実験の数を報告しませんでしたが、累積実時間を報告しました。彼らの投稿からは、彼らが

システムは複数の実験を並行して実行し、各実験の一般的な期間はどれくらいでしたか。したがって、結果の X 軸のスケーリングは推定されたものであり、システムの実際の実験効率に対応していない可能性があります。
アブレーションの目的で、Catalyst ソルバーの単純化された「線形」構成とも比較します。この構成では、進化ベースのメカニズムは無効になります。他のすべての要素 (ベース LLM、ハードウェア、プロンプトなど) は、同一の比較を可能にするために、主な結果と同一に保たれます。
* 実験の数 (x 軸) は Recursive の実行に対して推定されたものです
Catalyst の進化ベースのソルバーは、オリジナルの AutoResearch と Catalyst 独自の線形アブレーション ベースラインの両方を大幅に上回ります。
測定された実行では、Imbue Catalyst は val_bpb スコア 0.9361 を達成しました。特に、340 回の実験後に実行を停止したとき、val_bpb はまだ頭打ちになっておらず、十分な時間と計算予算があればさらなる改善が可能である可能性があることを示唆しています。
私たちのソリューションは、Recursive が報告した H100 の実行によって達成された最高の結果も上回っていましたが、H100 でさらに長く実行し続けていたら、ソリューションがさらに改善されていたかどうかは不明です。また、彼らの実行における実験の数 (x 軸) は不明であり、私たちが知っている限りでは、彼らのソリューションは、特定の val_bpb に到達する点で私たちのソリューションよりも効率的である可能性があることを再度強調したいと思います。
私たちの進化ベースのソルバーは、どのようにして単一のエージェントよりもこれほど大きな差で優れたパフォーマンスを発揮するのでしょうか?
私たちの直線走行に関与したエージェントの痕跡から、発煙筒が見つかります。あるエージェントのコメントは次のとおりです。
「プログラムの状態 (すべての機構軸が閉じている[…]) を考慮すると、最高値の次のステップは […] バイト同一のレコード対応スタンドアロン シード-97 db128 の再ロールです。」

記録グラフ。これは […] 唯一の記録を進めるアクションです。 」
言い換えれば、エージェントは、すべての最適化アイデアが使い果たされたと結論付ける状態になりました。その時点で、ランダムなシードを変えて同じ設定を繰り返し「再ロール」することに頼っていました。この状態に達すると、「ズームアウト」して、問題を解決するための他の角度が利用可能かどうかを再検討することはできないようです。
私たちの調査では、この現象が非常に頻繁に観察されています。今日のフロンティア LLM は、以前の結論を修正して「振り出しに戻る」ことに苦労することがあります。むしろ、彼らは、データの別の説明を検討することを妨げる「仮説の崩壊」を伴う、トンネル視野に似た現象に苦しんでいます。
エージェントが永続メモリにアクセスできる場合、ハーネスの場合のように、CoT 推論トレースと全体的なエージェント コンテキストがターンの間にクリアされる場合でも、この動作が観察されることがあります。
この問題は、現在のモデルの別の弱点によってさらに悪化します。モデルは、実験結果を解釈する際に確証バイアスや認識エラーを示すことがよくあります。人間によるレビューでプロットが仮説をまったくサポートしていないか、実験設定の欠陥を示す明確な指標が含まれていることが明らかになった場合でも（特定の結論を導くにはトレーニング ステップの数が不十分であるなど）、モデルが実験の結果プロットが以前の仮説を決定的に支持していると結論付けるのを私たちは見てきました。
解釈を進化させて行き詰まりから脱出する
これらの問題に対処するために、Imbue Catalyst はダーウィンの進化論からインスピレーションを得たアイデアを実装しています。以前の研究では、LLM ベースの進化をコードの最適化にどのように使用できるかを調査しました。 Catalyst はこれらのアイデアに基づいて構築されていますが、

2 つの機能強化:
これにより、手書きのミューテーション関数やスコアリング関数が不要になります。最適化問題は自然言語として指定でき、Catalyst はコーディング エージェントを使用して仕様を適切な検証ツールに変換します。
実験の共有と組み合わせて、「解釈ストランド」の概念を追加します。この組み合わせにより、進化プロセスの実験効率が大幅に向上し、データ収集に高価な実験や長時間にわたる実験が必要な問題に進化を適用できるようになります。
解釈ストランドは、観察された実験と文献検索に対するエージェントの現在の解釈を表します。さらに、各要素には、エージェントの現在の研究ロードマップ、つまり研究タスクをさらに進める方法に関する未解決の質問、賭け、アイデアのリストが要約されています。
Imbue Catalyst では、解釈鎖が進化の主な対象です。私たちは、任意の時点で 1 つだけではなく、解釈鎖の母集団全体を維持します。
最初は、すべて同じ初期適合度スコアを持つ少数の空のストランドから開始します。
この開始点から、一連の反復を進めていきます。反復では、(1) データ収集と (2) 統合とスコアリングの 2 つのモードが交互に行われます。通常、ほとんどの反復はデータ収集の反復です。ナノチャットの場合、9:1 の比率を使用します。つまり、10 回ごとの反復は統合とスコアリングのタイプになります。
データ収集の反復は上に示したように機能します。まず、母集団からストランドをサンプリングします。サンプリングは、Darwinian Evolver フレームワークのアプローチを使用して、各ストランドの現在のフィットネス スコアによって重み付けされます。次に、別々のエージェントに、サンプリングされた鎖ごとに 1 つずつ、次の実験を提案するよう依頼します。提案は以下に基づいてランク付けされます。

予想される研究への影響を考慮し、現在のイテレーションで実行するトップを選択します。
実験が完了すると、その結果は、最初にサンプリングされたストランドのスーパーセットに提示され、母集団からの追加の適応度加重サンプリングによって拡張されます。別々のエージェントは、新しい実験結果を考慮して各鎖の解釈を更新するように求められます。
数回の繰り返しごとに、追加の知識統合とスコアリングのステップを実行します。
知識の統合: ストランドのサブセットは、最近の観察を「理論」、つまりこれまでに確認されたすべてのデータの現在の理解を統合したスナップショットに統合するように求められます。このプロセスの一環として、エージェントはそれぞれの研究ロードマップを改訂することも求められます。
分岐: 一部の統合ターゲットについては、エージェントに対し、現在の理論の 2 番目の代替継続を分岐するよう特に要求します。エージェントは、データの別の解釈を検討し、明確で、しばしばより大胆な研究ロードマップを考案することが奨励されます。この分岐した理論は新しい解釈の始まりを形成し、それが母集団に追加されます。
フィットネススコアリング: 新たに統合および/または分岐したすべてのストランドが 2 つの次元に沿って評価されます。各ストランドのエージェントは、その時点での最善の理解に基づいて、当面の研究課題に対する解決策の候補を提案するように求められます。それらのソリューションは評価されます。それらを実行し、val_bpb のパフォーマンスを観察します。別の審査エージェントも、報酬ハッキングの兆候やユーザーの調査仕様の順守に問題がないかどうかをチェックします。直観的には、データを「より正確に」理解した解釈鎖は、より良い結果を生み出すことができると期待しています。

逆に、よりパフォーマンスの高い解候補は、より正確で完全な、したがってより望ましい解釈鎖を示す代理として採用されます。
さらに、エージェントに各分野の研究ロードマップをレビューし、それぞれの新規性をランク付けするよう依頼します。
ストランドの全体的な適合度スコアは、そのソリューション候補のパフォーマンスとその研究ロードマップの新規性を組み合わせたものです。
実験効率と多様性のトレードオフ
私たちのアプローチでは、次の 3 つの要素を通じてストランド間に多様性を導入します。
解釈鎖の分岐は上記で説明しました。これは最も直接的で、通常は多様性の最も強力な推進力ですが、統合とスコアリングの反復中にのみ発生します。
実験の可視性サンプリング: 研究プロセスの初期段階、母集団が小さいとき、各結果はすべての解釈鎖によって観察されます。ただし、母集団が増加するにつれて、私たちのプロセスは、解釈鎖のランダムなサブセットのみに所定の結果を提示し始めます。これは、時間の経過とともに、各ストランドで、重複するがわずかに異なる実験のサブセットが観察されることを意味します。
LLM トークン サンプリングの確率性: 最後になりましたが、他のすべてが同じであっても、基礎となる LLM のランダム トークン サンプリングが原因で、時間の経過とともにストランド間の変動が発生します。
の

[切り捨てられた]

## Original Extract

Automating AI model research with evolution - Imbue
/ blog Products Bouncer
Article / Research Automating AI model research with evolution
Escaping dead ends by evolving interpretations
The experiment efficiency vs. diversity tradeoff
Appendix: Reproduction details
We’re open-sourcing Imbue Catalyst , an AI tool for computational research and scientific discovery
Our evolution-based optimizer improves nanochat LLM performance 3x further than a regular AutoResearch agent
Recent model and harness improvements have enabled practical use cases of “self-improvement”: AI systems that iteratively improve their own foundations.
Approaches such as Darwin Godel Machine focus on improving a system’s harness, while others optimize the training recipes, architecture choices and algorithms behind the models themselves ( Andrej Karpathy’s AutoResearch , Recursive’s “First Steps Toward Automated AI Research” ). Even though many of these results have been on small “toy” models, it is becoming clear that the next generation of LLMs is already being built with the help of today’s AI systems (e.g. Anthropic “When AI builds itself” ).
Today, we’re announcing Imbue Catalyst . Catalyst is a research tool that uses evolution-inspired methods to help with the following research tasks:
Optimize a piece of code, algorithm, or model with respect to a given metric
Find solutions that satisfy certain programmatic verification criteria
Discover explanations for computationally reproducible phenomena
Assist with reviewing, formalizing, and editing theories in computational research fields
We believe that AI research works best as an open, collaborative effort. Hence, we’re releasing Imbue Catalyst under an AGPL-3.0 license. We’re looking forward to seeing what you’ll discover with it!
In this post, we’ll use Catalyst to optimize nanochat , a small transformer language model trained from scratch, and dive into the mechanisms that allow Catalyst to continue discovering improvements well after regular coding agents level off.
Andrej Karpathy first proposed the AutoResearch setup, in which a coding agent is prompted to improve the training setup for a small transformer language model ( nanochat ). It is asked to do so by iteratively running experiments, and using git to keep track of its progress.
The setup starts out with a training script that runs for exactly 5 minutes on a single H100 GPU. The agent is free to change the model architecture, training recipe and hyperparameters in order to minimize the resulting BPB (bits per byte) metric on a validation dataset, as long as it doesn’t exceed the 5-minute training time budget.
Nanochat training optimization is an example of what we call a verifiable goal in Catalyst. Verifiable in this context means that you can use a program to measure and/or verify how well a candidate solution achieves the goal at hand. In this case, the verifier is measuring the BPB metric of the resulting model.
We compare Imbue Catalyst’s verifiable goal solver, ”Catalyst (Evolution)”, against the following reference points:
Plain AutoResearch. We show the “mean of 3 runs” data reported by Weco AI .
Recursive Superintelligence’s recently reported results . With two caveats: Recursive switched to a different GPU (B200) part-way through their run. We only include the first part of their run that ran on the original H100 GPU.
Recursive did not report the number of experiments in their blog post, but rather reported cumulative wall-clock time. It is unclear from their post whether their system executed multiple experiments in parallel, and what the typical duration of each experiment was. Thus, our x-axis scaling of their results is estimated and might not correspond to the true experiment efficiency of their system.
For ablation purposes, we also compare against a simplified “linear” configuration of Catalyst’s solver. In this configuration, the evolution-based mechanisms are disabled. All other factors (base LLM, hardware, prompts, etc.) are kept identical to our main result to allow for an apples-to-apples comparison.
* Number of experiments (x-axis) is estimated for Recursive’s run
Catalyst’s evolution-based solver significantly outperforms both the original AutoResearch and Catalyst’s own linear ablation baseline.
In the measured run, Imbue Catalyst achieved a val_bpb score of 0.9361. Notably, when we stopped the run after 340 experiments, val_bpb had not plateaued yet, suggesting that further improvements could be possible given enough time and compute budget.
Our solution also surpassed the best result achieved by Recursive’s reported H100 run, but it is uncertain whether their solution would have continued improving if they had kept running on the H100 for longer. We would also like to re-emphasize that the number of experiments in their run (x-axis) is unknown, and, for all we know, their solution could be more efficient than ours in reaching a given val_bpb.
How does our evolution-based solver outperform a single agent by such a big margin?
A smoking gun can be found in the traces from the agents involved in our linear run. Here is one agent’s comment:
“ Given the program state — all mechanistic axes closed[…] — the highest-value next step is […] a record-capable standalone seed-97 db128 re-roll of the byte-identical record graph. This is […] the only record-advancing action. ”
In other words: The agent got into a state where it concluded that all optimization ideas had been exhausted. At that point, it resorted to “re-rolling” the same setup repeatedly with varying random seeds. Once it reached this state, it seemed unable to “zoom back out” and reconsider whether any other angles for attacking the problem were available.
In our research, we have observed this phenomenon quite frequently: Today’s frontier LLMs can struggle with revising earlier conclusions and “going back to the drawing board.” Rather, they suffer from something akin to tunnel vision, accompanied by a “hypothesis collapse” that prevents them from considering alternative explanations of the data.
When agents have access to persistent memory, this behavior can be observed even when CoT reasoning traces and the overall agent context are cleared in between turns, as is the case for our harness.
The problem is further exacerbated by another weakness in current models: Models often show confirmation bias and perception errors when interpreting experiment results. We’ve seen models conclude that an experiment’s result plot was conclusively in support of their prior hypothesis, even when a human review made clear that the plot either didn’t support the hypothesis at all, or contained clear indicators pointing to a faulty experiment setup (e.g. insufficient number of training steps to allow for a particular conclusion).
Escaping dead ends by evolving interpretations
To combat these issues, Imbue Catalyst implements ideas inspired by Darwinian evolution. In our previous work, we’ve investigated how LLM-based evolution can be used for code optimization . Catalyst builds on these ideas, but makes two enhancements:
It removes the need for hand-written mutation and scoring functions. The optimization problem can be specified as natural language, and Catalyst uses coding agents to translate the specification into an appropriate verifier.
We add the concept of “interpretation strands”, paired with experiment sharing. This combination allows us to significantly increase experiment efficiency in the evolution process, making it possible to apply evolution to problems that require expensive and/or long-running experiments to gather data.
An interpretation strand represents an agent’s current interpretation of its observed experiments and literature searches. Each strand additionally summarizes the agent’s current research roadmap - a list of open questions, bets and ideas for how the research task can be advanced further.
In Imbue Catalyst, the interpretation strand is the primary object of evolution: We maintain not just one, but an entire population of interpretation strands at any given point in time.
Initially, we start with a small number of empty strands, all with the same initial fitness scores.
From this starting point, we proceed through a sequence of iterations. Iterations alternate between two modes: (1) data gathering and (2) integration & scoring . Typically, most iterations are data gathering iterations. For nanochat, we use a ratio of 9:1, i.e. every 10th iteration is of the integration & scoring type.
Data gathering iterations work as depicted above: First, we sample strands from the population. Sampling is weighted by each strand’s current fitness score, using the approach from our Darwinian Evolver framework . We then ask separate agents to each propose a single next experiment, one per sampled strand. The proposals are ranked based on their expected research impact, and we select the top for execution in the current iteration.
Once the experiments have been completed, their results get presented to a superset of the initially sampled strands, extended by an additional fitness-weighted sampling from the population. Separate agents are asked to update each strand’s interpretations given the new experiment results.
Every few iterations, we perform the additional knowledge integration & scoring steps:
Knowledge integration: A subset of strands is asked to integrate their recent observations back into a “theory” - a consolidated snapshot of their current understanding of all data seen so far. As part of this process, the agents are also asked to revise their respective research roadmaps.
Branching: For some of the integration targets, we specifically request the agents to branch off a second alternative continuation of their current theory. The agents are encouraged to consider alternative interpretations of the data and to come up with a distinct, often bolder, research roadmap. This branched theory forms the beginning of a new interpretation strand, which gets added back to the population.
Fitness scoring: All newly integrated and/or branched strands get evaluated along two dimensions: An agent for each strand is asked to propose a solution candidate for the research problem at hand, using its best understanding to that point. Those solutions are evaluated, e.g. by executing them and observing their val_bpb performance. A separate review agent also checks for any indications of reward hacking or adherence issues with the user’s research specifications. Intuitively, we expect that an interpretation strand that has a “more correct” understanding of the data will also be able to produce a better solution candidate, and conversely, a better performing solution candidate is taken as a proxy to indicate a more accurate, complete and hence more desirable interpretation strand.
Additionally, we ask an agent to review the research roadmaps of each strand and rank the novelty of each.
The overall fitness score for a strand is a combination of the performance of its solution candidate and the novelty of its research roadmap.
The experiment efficiency vs. diversity tradeoff
Our approach introduces diversity between strands through three factors:
The interpretation strand branching described above. This is the most direct, and typically strongest driver of diversity, but it only happens during integration & scoring iterations.
Experiment visibility sampling: Early in the research process, when the population is small, each result is observed by every interpretation strand. However, as the population grows, our process begins presenting a given result to only a random subset of interpretation strands. This means that over time, each strand will have observed an overlapping, but slightly different subset of experiments.
LLM token sampling stochasticity: Last but not least, even with everything else being equal, variations between strands would arise over time simply due to the random token sampling in the underlying LLMs.
The

[truncated]
