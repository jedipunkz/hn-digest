---
source: "https://research.google/blog/thinking-to-recall-how-reasoning-unlocks-parametric-knowledge-in-llms/"
hn_url: "https://news.ycombinator.com/item?id=48693524"
title: "Thinking to recall: How reasoning unlocks parametric knowledge in LLMs"
article_title: "Thinking to recall: How reasoning unlocks parametric knowledge in LLMs"
author: "krackers"
captured_at: "2026-06-27T00:30:58Z"
capture_tool: "hn-digest"
hn_id: 48693524
score: 3
comments: 0
posted_at: "2026-06-26T23:49:43Z"
tags:
  - hacker-news
  - translated
---

# Thinking to recall: How reasoning unlocks parametric knowledge in LLMs

- HN: [48693524](https://news.ycombinator.com/item?id=48693524)
- Source: [research.google](https://research.google/blog/thinking-to-recall-how-reasoning-unlocks-parametric-knowledge-in-llms/)
- Score: 3
- Comments: 0
- Posted: 2026-06-26T23:49:43Z

## Translation

タイトル: 思い出すための思考: 推論が LLM のパラメトリック知識をどのように解き放つか

記事本文:
思い出すための思考: 推論が LLM のパラメトリック知識をどのように解き放つか
メインコンテンツにスキップ
当社が注力している多くの分野をご覧ください
協力的なエコシステムの構築
発見を現実世界への影響に変える
当社が注力している多くの分野をご覧ください
協力的なエコシステムの構築
発見を現実世界への影響に変える
Google
研究
Google AI
当社のすべての AI について学ぶ
Googleディープマインド
AI のフロンティアを探索する
Google Labs
AI 実験を試してみる
研究
リソース
カンファレンスとイベント
キャリア
ブログ
について
検索
ホーム
思い出すための思考: 推論が LLM のパラメトリック知識をどのように解き放つか
Google Research、研究員、Zorik Gekhman 氏と Jonathan Herzig 氏
私たちは、複雑な段階的な解決策が必要ない場合でも、言語モデルが単純な事実を思い出すのに推論が役立つという直感に反する現象を研究します。我々は、この現象が 2 つのメカニズムによって引き起こされていることを示します。(1) 生成された推論トークンを使用して潜在計算を実行すること、および (2) 関連する事実を生成して正解を思い出すことを促進することです。
大規模言語モデル (LLM) で段階的な推論トレース (一般に思考連鎖 (CoT) として知られる) を生成できるようにすると、複雑なタスクのパフォーマンスが向上することが確立されています。モデルが難しい数式を解いたり、ソフトウェアを作成したり、マルチホップの事実に関する質問に答えたりする場合、問題を管理可能な論理ステップに分割することは非常に効果的です。
ただし、単純なシングルホップの事実に関する質問については、このアプローチの有用性は依然として不明瞭です。たとえば、「メアリー エングル ペニントンが全米発明家の殿堂入りを果たしたのは何年ですか?」のようなクエリを考えてみましょう。 LLM は、そのパラメトリック メモリ (重みに直接エンコードされた知識) にファクトが保存されているか、または保存されていません。複雑な算術や論理的推論は必要ありません。では、なぜ

推論トレースの助けになりますか？
「 思い出すための思考: 推論が LLM のパラメトリック知識を解き放つ方法」では、この現象を調査します。モデルが推論トレースを生成できるようにすると、他の方法では実質的に到達できない正しい答えが得られることを示します。実行する複雑な推論ステップがない場合に、推論支援がパラメトリック知識を呼び出す理由を理解するために、一連の仮説に基づいた制御実験を実施します。私たちの調査結果は、これを促進する 2 つの相補的なメカニズム、つまり計算上のバッファー効果と事実によるプライミングを明らかにしました。
知識の限界を探る
まず、pass@k メトリックを使用してパラメトリックな再現能力の境界を測定します。 pass@k は、モデルによって生成された 1 つの回答のみをチェックするのではなく、生成された複数の試行内に正しいファクトが存在するかどうかをチェックします。 pass@k は、モデルの出力分布における成功した推論パスの存在を正確なランキングにあまり依存せずに評価することで、現在のモデルのトップ 1 の動作のみを調べるのではなく、事実を再現する推論の可能性を推定するのに役立ちます。パラメトリック知識を制御しながら推論の影響を評価するために、推論を有効または無効 (オンまたはオフに切り替える) できる推論 LLM (R-LLM) に焦点を当て、これら 2 つのモード間の pass@k を比較します。私たちは Gemini-2.5 (Flash および Pro) および Qwen3-32B モデルに焦点を当て、2 つの困難なクローズドブック QA データセット、 SimpleQA Verified および EntityQuestions を使用します。
結果は驚くほど一貫しています。推論が有効な場合、モデルは推論がオフの場合には事実上回復不可能な答えを正常に呼び出すことができます。重要なのは、この改善は、モデルが複雑な質問を分解しているからだけではありません。これは、主に以下を含むデータセットに意図的に焦点を当てた結果です。

単純なシングルホップの質問。
2 つのクローズドブック QA データセットと 3 つの LLM に @𝑘 曲線を渡し、推論を有効にした場合 (ON) と推論を無効にした場合 (OFF) で同じモデルを比較します。
これらの結果から疑問が生じます。効果が段階的な推論によってもたらされない場合、どのような推論パターンによってモデルは正しい答えを取得できるのでしょうか?
メカニズム 1: 計算バッファー
私たちの最初の仮説は、生成の仕組みに焦点を当てています。私たちは、追加のトークンの生成が、追加のフォワード パスを提供することによって計算時間の延長として機能するという長年の仮説を採用し、R-LLM のパラメトリック知識想起の新しい設定でそれをテストします。具体的には、モデルは、生成される実際の意味論的コンテンツとは無関係に、潜在的な処理を実行するための計算バッファーとしてこれらの推論トークンを暗黙的に使用すると仮説を立てます。
これをテストするために、推論トレースからすべての意味のあるコンテンツを削除する実験を設計します。モデルの推論プロセスを傍受し、その生成されたトレースを意味のない文字列「Let me think」に置き換えます。これは、元の推論トレースの長さと一致するまで何度も繰り返されます。次に、このダミー テキストに基づいて最終的な答えをモデルに予測させます。
驚くべきことに、この無意味なトレースに基づいてモデルを条件付けすると、推論が完全にオフになったベースラインと比較して、正解を想起する能力が大幅に向上します。これは、モデルにより多くの計算滑走路を与えるだけで、モデルの内部状態が改善され、到達しにくい事実を取得できるという強力な証拠を提供します。
Gemini-2.5-Flash での計算バッファー効果。 ON ダミーは、元のトレースのトークン長と一致するように繰り返される、事実の内容を含まない短いシーケンスで思考トレースをオーバーライドします。
ただし、この compute-buf

フェル効果には限界があります。ダミー テキストを長くすると、最終的には利益が減少し、モデルの自然な推論トレースのパフォーマンスと完全に一致することはありません。これは、追加の計算は役立ちますが、思考の実際の内容が依然として重要であることを意味します。
ダミー推論トレースを条件付けする場合のトークン単位の入力長の関数としての推論の有効性。 ON ダミー X は、入力長が X トークンになるように繰り返される短いダミー シーケンスで推論トレースをオーバーライドします。推論有効性メトリック (Ω) は、すべての k 値にわたる pass@k ゲインを要約します。これを、推論 ON モードと OFF モード間の pass@𝑘 の加重平均相対差として定義します。
単純な事実に関する質問に対して生成された自然推論トレースを分析すると、共通のパターンに気づきます。モデルは論理的な証明を書き出しているわけではありません。彼らは関連する事実を明らかにしつつある。
人間の認知には、 拡散活性化 として知られる概念があり、特定の概念を処理すると、関連する概念が意味記憶内でプライミングされ、検索が容易になります。私たちは、言語モデルが同様の生成的自己検索メカニズムを示し、これを事実のプライミングと呼ぶと仮説を立てます。質問にトピック的に関連する事実を生成することにより、モデルは正しい答えの検索を容易にするコンテキスト ブリッジを構築します。
仮説をテストするために、モデルの推論トレースから具体的な事実だけを抽出し、厳密なフィルタリングを適用してフィラーテキスト、検索プラン、または最終的なターゲット回答の明示的な言及を取り除きます。次に、想起された事実の影響を分離し、想起された事実の短いリストに基づく条件付けが推論の利益のほとんどを回復し、推論がオフの場合でも役立つことを示します。
ジェミンに対する事実のプライミング効果

i-2.5-フラッシュ。まず、推論中に言及された事実を抽出します。 ON Facts は、この短いファクト リストでモデルの元の推論トレースをオーバーライドし、最終的な答えを再生成します。一方、OFF Facts は、プロンプトの一部として追加の入力コンテキストとして提供されたファクト リストを使用してモデル推論を無効にして実行します。
たとえば、ネパールの 10 代国王の名前を尋ねられた場合、推論モデルでは、最初に前の 9 人の国王がリストされる可能性があります。最初の 9 つを呼び出すことは意味論的なウォームアップとして機能し、ネットワークが 10 番目を正常に呼び出す準備をします。事実そのものが踏み台です。
「事実のプライミング」の動作の図。中間の事実の検索 (以前の 9 王のリスト) により、ネパールの 10 代国王を首尾よく思い出すためにモデルが準備されます。モデルは、推論を有効 (ON) にすると正解に成功しますが、推論を有効にしないと失敗します。また、推論中に想起された事実の短いリスト (ON 事実) のみを予測の条件とする場合にも成功します。
生成的自己検索は強力なメカニズムですが、根本的なリスクをもたらします。モデル自体がこれらの中間事実を生成するため、幻覚が見られる可能性があります。したがって、これらの推論段階のエラーが最終的な答えにどのような影響を与えるかを確認します。これを確認するために、検索可能な検証ツールを使用して大規模な監査パイプラインを構築し、数十万の推論トレースにわたって生成されたすべての中間ファクトの正しさを独立してチェックします。
監査により、明確なパターンが明らかになりました。推論トレースに幻覚的な中間事実が 1 つでも含まれている場合、モデルが正しい最終答えに到達する可能性は大幅に低くなります。これは、実際のプライミングメカニズムは効果的ではあるものの、脆弱である可能性があることを示唆しています。
推理痕跡に幻覚が含まれる場合の正答率

幻覚を含まないもの（クリーン）と比較したもの（幻覚あり）。
これらのメカニズムを理解することで、モデルの信頼性を向上させるための実践的な手段が得られます。事実によるプライミングは効果的であり、幻覚的な中間事実はパフォーマンスを低下させるため、両方の洞察を活用してモデルの精度を向上させることができます。
これらの洞察の可能性を評価するために、単一の質問に対して複数の推論軌跡を生成し、検証可能で幻覚のない事実を含むものだけを保持するテスト時の選択戦略を使用します。これらの軌道を優先すると、精度が大幅に向上します。実際には、この優先順位付けは、事実に裏付けられた中間ステップを奨励するプロセス報酬を介してトレーニング中に実装できます。
事実の再現性と事実の正しさに基づいたテスト時の選択基準の下で期待される精度。
私たちの調査結果は、言語モデルでの推論が、単なるタスクの分解や数学的論理よりもはるかに幅広い目的に役立つことを強調しています。これは、モデルの内部メモリを公開し、そのパラメトリック知識境界を拡張するための基本的なメカニズムとして機能します。これらの洞察は、将来の研究に刺激的な方向性をもたらします。事実に基づいて正確な推論トレースによってより良い答えが得られることがわかれば、トレーニング レシピをさらに最適化できることがわかります。事実に基づいてサポートされる中間ステップを特に奨励するプロセス報酬を利用することで、本質的に信頼性が高く、幻覚が起こりにくいモデルをトレーニングできる可能性があります。私たちは、研究コミュニティが推論、記憶、検索の交差点をどのように探求し続けるかを見ることを楽しみにしています。
この研究は、Zorik Gekhman、Roee Aharoni、Eran Ofek、Mor Geva、Roi Reichart、Jonathan Herzig によって実施されました。 Eyal Ben-David と Avinatan Hassidim に感謝します。

または、作品とその貴重な提案をレビューします。

## Original Extract

Thinking to recall: How reasoning unlocks parametric knowledge in LLMs
Skip to main content
Explore our many areas of focus
Building a collaborative ecosystem
Translating discovery into real-world impact
Explore our many areas of focus
Building a collaborative ecosystem
Translating discovery into real-world impact
Google
Research
Google AI
Learn about all our AI
Google DeepMind
Explore the frontier of AI
Google Labs
Try our AI experiments
Research
Resources
Conferences & events
Careers
Blog
About
Search
Home
Thinking to recall: How reasoning unlocks parametric knowledge in LLMs
Zorik Gekhman and Jonathan Herzig, Research Scientists, Google Research
We study the counterintuitive phenomenon where reasoning helps language models recall simple facts, even when no complex step-by-step solutions are required. We show that this phenomenon is driven by two mechanisms: (1) using generated reasoning tokens to perform latent computation, and (2) generating related facts to prime correct answer recall.
It is well-established that allowing large language models (LLMs) to generate step-by-step reasoning traces, commonly known as chain-of-thought (CoT), enhances performance on complex tasks. When a model solves difficult math equations, writes software, or answers multi-hop factual questions, breaking the problem down into manageable logical steps is highly effective.
However, the utility of this approach remains unclear for simple, single-hop factual questions. For instance, consider a query like: "What year was Mary Engle Pennington inducted into the National Inventors Hall of Fame?" An LLM either has the fact stored in its parametric memory (knowledge encoded directly into its weights) or it doesn't; no complex arithmetic or logical deduction is required. So why would a reasoning trace help?
In " Thinking to Recall: How Reasoning Unlocks Parametric Knowledge in LLMs ”, we investigate this phenomenon. We demonstrate that allowing a model to generate a reasoning trace unlocks correct answers that are otherwise effectively unreachable. To understand why reasoning aids parametric knowledge recall when there are no complex reasoning steps to execute, we conduct a series of hypothesis-driven controlled experiments. Our findings reveal two complementary mechanisms driving this: a computational buffer effect and factual priming.
Probing the knowledge boundary
We first measure the parametric recall capability boundary using the pass@k metric. Instead of only checking one model-generated answer, pass@k checks if the correct fact exists within multiple generated attempts. By evaluating the presence of successful reasoning paths in the model’s output distribution while being less sensitive to their exact ranking, pass@k helps us estimate the potential of reasoning for factual recall, rather than only looking at the current model’s top-1 behavior. To assess the impact of reasoning while controlling for parametric knowledge, we focus on reasoning LLMs (R-LLMs) where reasoning can be enabled or disabled (toggled on or off), and compare pass@k between these two modes. We focus on the Gemini-2.5 (Flash and Pro) and Qwen3-32B models, using two challenging closed-book QA datasets: SimpleQA Verified and EntityQuestions .
The results are surprisingly consistent. When reasoning is enabled, the models successfully recall answers that are virtually unrecoverable when reasoning is off. Importantly, this improvement isn't just because the model is decomposing complex questions. This results from our deliberate focus on datasets containing predominantly simple, single-hop questions.
Pass@𝑘 curves across two closed-book QA datasets and three LLMs, comparing the same models with reasoning enabled (ON) vs reasoning disabled (OFF).
These results raise the question: if the effect does not come from step-by-step reasoning, what reasoning patterns enable the model to retrieve the correct answer?
Mechanism 1: The computational buffer
Our first hypothesis focuses on the mechanics of generation. We take the long-standing hypothesis that generating extra tokens acts as extended computation time by providing additional forward passes, and test it in the new setting of parametric knowledge recall in R-LLMs. Specifically, we hypothesize that models implicitly use these reasoning tokens as a computational buffer to perform latent processing, independent of the actual semantic content being generated.
To test this, we design an experiment that removes all meaningful content from the reasoning trace . We intercept the model's reasoning process and replace its generated trace with a meaningless string "Let me think" , repeated over and over until it matches the length of the original reasoning trace. We then let the model predict the final answer conditioned on this dummy text.
Remarkably, conditioning the model on this meaningless trace substantially improves its ability to recall the correct answer compared to the baseline where reasoning is completely turned off. This provides strong evidence that simply giving the model more computational runway helps it refine its internal state and fetch hard-to-reach facts.
Computation buffer effect on Gemini-2.5-Flash. ON Dummy overrides the thinking trace with a short sequence without factual content that is repeated to match the token length of the original trace.
However, this compute-buffer effect has its limits. Pushing the dummy text to longer lengths eventually offers diminishing returns, and it never fully matches the performance of the model's natural reasoning traces. This means that while extra computation helps, the actual content of the thoughts still matters.
Reasoning effectiveness as a function of the input length in tokens when conditioning on dummy reasoning traces. ON Dummy X overrides the reasoning trace with a short dummy sequence which is repeated such that the input length will be X tokens. The reasoning effectiveness metric (Ω) summarizes the pass@k gains across all k values. We define it as a weighted average relative difference in pass@𝑘 between reasoning ON and OFF modes.
When we analyze the natural reasoning traces generated for simple factual questions, we notice a common pattern. The models aren't writing out logical proofs; they are surfacing related facts.
In human cognition, there is a concept known as spreading activation , where processing a specific concept primes related concepts in semantic memory, making them easier to retrieve. We hypothesize that language models exhibit a similar generative self-retrieval mechanism, which we call factual priming . By generating facts topically related to the question, the model builds a contextual bridge that facilitates the retrieval of the correct answer.
To test hypotheses, we extract just the concrete facts from the model’s reasoning traces, applying strict filtering to strip away any filler text, search plans, or explicit mentions of the final target answer. We then isolate the effect of the recalled facts, and show that conditioning on a short list of recalled facts recovers most of reasoning’s gains and helps even when reasoning is OFF.
Factual priming effect on Gemini-2.5-Flash. We first extract the facts mentioned during reasoning. ON Facts overrides the models’ original reasoning trace with this short fact list and regenerates the final answer, while OFF Facts runs the model reasoning disabled with the fact list provided as additional input context as part of the prompt.
For example, if asked for the name of the 10th King of Nepal, a reasoning model might first list the previous nine kings. Recalling those first nine acts as a semantic warm-up, priming the network to successfully recall the 10th. The facts themselves are the stepping stones.
An illustration of "factual priming" in action where intermediate factual retrieval (listing the previous nine Kings) primes the model to successfully recall the 10th King of Nepal. The model succeeds to answer correctly with reasoning enabled (ON) while failing without it. It also succeeds when the prediction is conditioned only on a short list of facts recalled during reasoning (ON Facts).
While generative self-retrieval is a powerful mechanism, it introduces a fundamental risk. Because the model generates these intermediate facts itself, they might be hallucinated. We thus check how these reasoning-stage errors impact the final answer. To find out, we build a large-scale auditing pipeline using a search-enabled verifier to independently check the correctness of every single intermediate fact generated across hundreds of thousands of reasoning traces.
The audit reveals a distinct pattern. If a reasoning trace contains even a single hallucinated intermediate fact, the model is significantly less likely to arrive at the correct final answer. This suggests that, while effective, the factual priming mechanism might be fragile.
Ratio of correct answers when reasoning traces contain hallucinations (hallucinated) compared to those that do not contain hallucinations (clean).
Understanding these mechanisms provides practical avenues for improving model reliability. Because factual priming is effective and hallucinated intermediate facts degrade performance, we can leverage both insights to improve model accuracy.
To evaluate the potential of these insights, we use a test-time selection strategy that generates multiple reasoning trajectories for a single question, retaining only those that contain verifiable, hallucination-free facts. Prioritizing these trajectories considerably improves accuracy. In practice, this prioritization could be implemented during training via process rewards that encourage factually supported intermediate steps.
Expected accuracy under test-time selection criteria based on factual recall and factual correctness.
Our findings highlight that reasoning in language models serves a much broader purpose than just task decomposition or mathematical logic. It acts as a fundamental mechanism for exposing a model's internal memory and expanding its parametric knowledge boundary. These insights open up exciting directions for future research. Knowing that factually accurate reasoning traces yield better answers suggests that training recipes can be further optimized. By utilizing process rewards that specifically encourage factually supported intermediate steps, we might be able to train models that are inherently more reliable and less prone to hallucination. We look forward to seeing how the research community continues to explore the intersections of reasoning, memory, and retrieval.
This research was conducted by Zorik Gekhman, Roee Aharoni, Eran Ofek, Mor Geva, Roi Reichart and Jonathan Herzig. We thank Eyal Ben-David and Avinatan Hassidim for reviewing the work and their valuable suggestions.
