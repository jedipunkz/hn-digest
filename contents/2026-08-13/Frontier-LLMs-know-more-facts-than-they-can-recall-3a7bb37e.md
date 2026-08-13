---
source: "https://research.google/blog/empty-shelves-or-lost-keys-recall-is-the-bottleneck-for-parametric-factuality/"
hn_url: "https://news.ycombinator.com/item?id=49288011"
title: "Frontier LLMs know more facts than they can recall"
article_title: "Empty shelves or lost keys? Recall is the bottleneck for parametric factuality"
author: "MarcoDewey"
captured_at: "2026-08-13T16:46:06Z"
capture_tool: "hn-digest"
hn_id: 49288011
score: 3
comments: 0
posted_at: "2026-08-13T16:03:12Z"
tags:
  - hacker-news
  - translated
---

# Frontier LLMs know more facts than they can recall

- HN: [49288011](https://news.ycombinator.com/item?id=49288011)
- Source: [research.google](https://research.google/blog/empty-shelves-or-lost-keys-recall-is-the-bottleneck-for-parametric-factuality/)
- Score: 3
- Comments: 0
- Posted: 2026-08-13T16:03:12Z

## Translation

タイトル: Frontier LLM は思い出せないほど多くの事実を知っています
記事タイトル: 棚が空になったり、鍵が紛失したり?パラメトリックな事実性のボトルネックは想起です

記事本文:
棚が空になったり、鍵が紛失したりしませんか?パラメトリックな事実性のボトルネックは想起です
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
棚が空になったり、鍵が紛失したりしませんか?パラメトリックな事実性のボトルネックは想起です
Nitay Calderon 氏と Gal Yona 氏、Google Research の研究員
LLM が事実を誤解するとき、それは彼らがそれを学習していないためでしょうか、それともすでにエンコードした内容を思い出せないためでしょうか?私たちの知識プロファイリング フレームワークは後者を明らかにします。フロンティア LLM はほぼすべての事実をエンコードしますが、その多くを思い出すのに苦労します。
大規模言語モデル (LLM) の信頼性を高めるには、事実性が不可欠です。モデルが事実に関する質問に不正確に回答した場合、それはその事実がエンコードされていないためでしょうか、それともファクトはエンコードされているもののアクセスできないためでしょうか?標準的な精度指標は、たとえ非常に異なる制限や非常に異なる介入を示唆しているとしても、これらのケースをまとめてまとめます。エンコーディングの失敗には、モデル サイズのスケーリングやデータ カバレッジの拡大が必要ですが、リコールの失敗は、LLM がすでにエンコードしたものをより有効に活用するのに役立つトレーニング後および推論時のメソッドを示している可能性もあります。
「Empty Shelves or Lost Keys? Recall Is the Bottleneck for Parametric Factuality」では、エンコーディングとリコールの両方を測定する行動フレームワークである知識プロファイリングを紹介し、それを使用してフロンティア LLM (Gemini3 など) における事実性の根底にあるボトルネックを調査します。

および GPT-5 )。次に、フロンティア LLM の多くの事実上のエラーは、空のシェルフ (エンコードの失敗) ではなく、キーの紛失 (リコールの失敗) としてよりよく理解されることを示します。
類推により、符号化は事実のパラメトリック表現を表し、想起は符号化された事実を外部からの手がかりなしで取得することを表し、認識は正しい事実が選択肢の中から提示された場合にそれを識別することを表します。この分析をサポートするために、WikiProfile を導入します。これは、Wikipedia から派生した 2,150 の事実のベンチマークであり、それぞれがエンコード、想起、認識を調査する 10 の質問とペアになっています。
中心的なアイデア: ナレッジプロファイリング
ナレッジプロファイリングは、分析の単位を個々の質問から事実に移します。モデルが特定の質問に正しく答えたかどうかを尋ねる代わりに、事実はどのような状態にあるのかというより広範な質問をします。各事実を 5 つの知識プロファイルのいずれかに分類します。(1) エンコードの失敗、(2) 想起の失敗、(3) 直接の想起、(4) 思考による想起、および (5) エンコードなしの推論。これらのプロファイルは、質問レベルの精度だけよりも有益な診断を提供します。
分類は、ファクトがエンコードされているかどうか、およびそのファクトがどの程度アクセス可能であるかに基づいて行われます。想起できない、直接想起できる、または思考することによってのみ想起できる (最終的な答えの前に中間計算を誘発し、思考の連鎖を促す LLM や思考に最適化された LLM を含む)。
事実を特徴付ける 5 つの知識プロファイル。
これを 3 つの行動概念で運用します。
エンコーディング: モデルは、トレーニング前のようなコンテキストでファクトを正しく再現できる場合、ファクトをエンコードします。私たちの設定では、命題補完とコンテキスト質問を使用してこれを測定します。これにより、事前トレーニング中に事実が自然に現れるコンテキストと同様のコンテキストにモデルが配置されます（

答えを明らかにする)、それにより、ファクトがエンコードされているかどうかを明らかにするためにモデルを準備します。
知識: モデルは、直接質問と逆質問の両方を含む、さまざまな表現にわたって意味的に同等の質問に正しく答えることができる場合、事実を知っています (たとえば、*A が B* の場合、直接質問は「B とは何ですか?」と尋ねますが、逆質問は「A とは何ですか?」と尋ねます)。
再現: モデルは、エンコードされた事実を知っている場合、その事実を再現します。何も考えずにその事実を思い出す場合、これを直接想起と呼びます。エンコードされていない事実を知っている場合、これをエンコードなしの推論と呼びます。これは、思考が有効で、モデルが他のエンコードされた事実に依存し、マルチホップ推論または知識に基づいた推測を実行する場合にのみ発生します。
上 : トレーニング前データの主要なソースである Wikipedia から事実を抽出します。左: LLM に元のコンテキスト内でファクトを再現するよう促すことで、エンコードを測定します。右: 私たちは、考えながら、または考えずに、さまざまな言い回しや関係の方向性について質問することによって知識を測定します。
ナレッジプロファイリングを運用可能にするために、自然に発生する事実の事実性を測定するように設計されたベンチマークである WikiProfile を構築しました。 WikiProfile は、プロンプト付き LLM、思考を備えた Gemini-2.5-Pro を活用した完全に自動化されたパイプラインを使用して構築されています。プロンプトは、保持された小さなサブセットに対する手動の最適化を通じて開発されました。事実を特定することによって、ウィキペディアのページから事実の候補を抽出します。これは、順序付けされたエンティティのペア (主語と目的語) を含む命題であり、主語が文書内で最初に出現する命題です。各ファクトは 10 個のタスクとペアになっています。2 つはエンコード用、4 つは知識評価用、4 つは認識用の多肢選択バリアントです。
生成、改良、修正という 3 段階のプロセスを通じて、直接的な質問と逆の質問を生成します。

そしてフィルタリングにより、各質問が明確で、具体的で、最小限であり、固有の回答があることを確認します。すべての質問は、検索エンジンに基づいてフィルタリングされます。複数の回答が返された場合や説明が必要な場合は無視します。この自動フィルタリングと最終的な手動検証ステップを経た後、ベンチマークには 2,150 のファクトが含まれます。
プロンプト LLM に基づく完全に自動化されたパイプライン。左 (紫): 事実の抽出と提案完了タスクの構築。中央 (赤と青): 生成、洗練、フィルタリングによる直接質問と逆質問の構築。右 (緑): 直接/逆ペアに基づいた残りの質問 (自然な表現、文脈、多肢選択バージョン) の作成。
13 個の LLM を評価します。各モデルは、考えながら、または考えずに評価されます。モデル、ファクト、タスクごとに 8 つの応答をサンプリングします。回答は、指示された LLM 自動評価者によって自動的に採点され (詳細については論文を参照)、約 450 万件の回答が生成されます。
主な結果: エンコードではなくリコールがボトルネックになる
最先端の LLM (Gemini-2.5-Pro、Gemini-3-Pro、および Flash、GPT-5) では、事実のエンコーディングは飽和に近づいていますが、再現率は飽和していません。 Gemini-3-Pro と GPT-5 の場合、事実の 95 ～ 98% がエンコードされていますが、これらのモデルはまだ事実の 26 ～ 34% を直接呼び出すことができません。たとえ考えたとしても、事実の 11 ～ 12% については失敗します。これは、フロンティア モデルでは、事実の誤りが知識の不在からではなく、保存されているが確実にアクセスできない知識から生じることが増えていることを意味します。つまり、ボトルネックは知識の獲得から知識の活用へと移りつつあるのです。
スケーリングはこのイメージを強化します。 Gemma 3 ファミリでは、モデルが大きいほどエンコードの失敗ははるかに少なくなりますが、リコールの失敗は依然としてかなり多く、より規模が大きくなります。

残りのエラーのシェア。スケーリングにより、モデルがアクセスできる内容が改善されるよりも、モデルが保存する内容が改善されます。
13 個の LLM にわたる 5 つのプロファイルの分布 (パーセンテージ)。黒い線は潜在的な知識を示します。示されているように、エンコードの失敗は規模が大きくなるにつれて急激に減少しますが、リコールの失敗はフロンティア モデルでも依然として発生します。
私たちの結果は、想起が事実を学習した条件と密接に結びついていることを示唆しています。クエリがトレーニング時のコンテキスト、表現、またはファクトに遭遇した順序から逸脱すると、思い出すのが難しくなります。これが組織的に発生する 2 つのケースを取り上げます。
まれな事実は暗号化されているが、思い出すのは難しい
これまでの研究では、LLM がロングテール (まれな) 事実に苦戦していることが示されており、多くの場合、これをモデルの能力の問題として捉えています。私たちの結果は、補完的な状況を示唆しています。人気の低いファクトと人気の高いファクトを比較すると、まれなファクトが人気のファクトに近い割合でエンコードされていることがわかります。エンコードにおけるギャップは比較的わずかです。ただし、再現率のギャップはさらに大きくなります。これはロングテール問題を再構成したものです。モデルのパラメーターには、多くのまれな事実が含まれていません。それらは存在しますが、アクセスするのは困難です。ボトルネックは知識の獲得から活用へと移りました。
エンコード率と直接再現率の観点から、2 つの人気層 (下位 20% と上位 20%) を比較します。 Δは層間のギャップを示します。示されているように、エンコードに関しては狭いですが、リコールに関しては広いです。すべての LLM の結果については、論文を参照してください。
逆質問は検証可能だが思い出すのが難しい
また、逆転の呪い、つまり LLM が「A は B である」ことは知っているが、「B とは何ですか?」には答えられない場合についても再訪します。一見すると、これは LLM に双方向の知識が欠けていることを示唆している可能性があります。しかし、私たちの結果は、この見方が改良されていることを示唆しています。オープンエンド型生成では

逆質問は一貫して直接質問より難しいです。ただし、多肢選択式の検証 (つまり、認識) では、逆質問は直接質問ほど難しくはなく、多くの場合、より簡単です。この解離が重要です。モデルが、注意をそらすものの間で正解が提示された場合にはそれを認識できるが、逆クエリではそれを生成できない場合、問題は単に双方向の知識が欠落しているということではありません。むしろ、ファクトはエンコードされているように見え、認識可能ですらありますが、クエリの方向がトレーニング中にファクトに遭遇した方法から逸脱すると、思い出すのが困難になります。逆転の呪いはリコール問題です。
検証 (多肢選択) と生成 (クローズドブック) という 2 つのタスクにわたって直接質問と逆質問を比較します。 Δは、ダイレクト設定とリバース設定の間のギャップを示します。 LLM は検証では逆質問を効果的に処理しますが、生成では困難を伴います。
回復メカニズムとして考える
次に、他の方法ではアクセスできない知識の回復を可能にするものは何かという問題に移ります。この目的を達成するために、私たちはこの役割を果たすための思考の可能性を検討します。考えることで、直接の想起が最も弱い場所での想起が最も強く改善されます。この利益は、まれな事実や逆質問の場合に特に顕著であり、人気の差と方向性の差の両方が狭まります。
私たちは、思考が想起（コード化された事実を知ること）に及ぼす影響を調べます。左: 2 つのファクト人気層 (下位 20% と上位 20%) を比較します。右: 直接質問と逆質問を比較します。人気または方向性のギャップは、Δ (思考なし) と ΔT (思考あり) で示されます。
示されているように、考えることでギャップが狭まります ( ΔT < Δ)。
より具体的には、思考に最適化されたモデルでは、コード化されているものの直接的には知られていない事実のおよそ 40 ～ 65% が思考により復元されます。逆に

ただし、エンコードされていない事実についてはあまり役に立ちません。このパターンは、思考が主に想起促進メカニズムとして機能することを示唆しています。これは、主に複雑な複数ステップの推論を通じて答えを導き出すのではなく、モデルが既にエンコードされた事実にアクセスするのに役立ちます。とはいえ、考えることは自由ではありません。これには計算コストがかかり、モデルがいつそれを呼び出すべきかを正確に決定する方法は依然として不明です。
事実が暗号化されているか (赤)、否か (黄色) に応じて、思考することで知られるようになる、知られていない事実の割合を報告します。思考によって最適化された LLM では、エンコードされたファクトの 40 ～ 65% が回復されますが、エンコードされていないファクトは 5 ～ 15% のみです。
ナレッジプロファイリングにより、LLM の事実に基づく動作を正確に診断できます。この方法論をウィキペディアの事実に適用した結果は、フロンティア LLM における事実誤認についての考え方の変化を示唆しています。エンコードがすでに飽和状態に近づいている場合、(モデルのサイズまたはデータの) スケーリングによる事実のさらなる向上は少なくなる可能性があります。コード化されているものの直接的には知られていない事実のかなりの部分を思考が回復できることを示すことで、事実性の次の改善は、より良い知識の獲得だけでなく、既知の知識のより良い利用からもたらされる可能性があります。

[切り捨てられた]

## Original Extract

Empty shelves or lost keys? Recall is the bottleneck for parametric factuality
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
Empty shelves or lost keys? Recall is the bottleneck for parametric factuality
Nitay Calderon and Gal Yona, Research Scientists, Google Research
When LLMs get facts wrong, is it because they never learned them or because they can't recall what they’ve already encoded? Our knowledge profiling framework reveals the latter: frontier LLMs encode nearly all facts, yet struggle to recall many of them.
Factuality is essential for making Large Language Models (LLMs) reliable. When a model answers a factual question incorrectly, is it because the fact was never encoded, or because the fact is encoded but not accessible? Standard accuracy metrics collapse these cases together, even though they suggest very different limitations and very different interventions. Encoding failures call for scaling model size or expanding data coverage, while recall failures might also point to post-training and inference-time methods that help LLMs better utilize what they already encode.
In “ Empty Shelves or Lost Keys? Recall Is the Bottleneck for Parametric Factuality ”, we introduce knowledge profiling , a behavioral framework that measures both encoding and recall , and use it to examine the underlying bottlenecks of factuality in frontier LLMs (such as Gemini3 and GPT-5 ). We then show that many factual errors in frontier LLMs are better understood as lost keys (recall failures) , not empty shelves (encoding failures) .
By analogy, we use encoding to denote parametric representation of facts, recall to denote retrieving encoded facts without external cues, and recognition to denote identifying the correct fact when it is presented among alternatives. To support this analysis, we introduce WikiProfile , a benchmark of 2,150 Wikipedia-derived facts, each paired with ten questions that probe encoding, recall, and recognition.
The core idea: Knowledge profiling
Knowledge profiling shifts the unit of analysis from individual questions to facts. Instead of asking whether a model answered a specific question correctly, we ask a broader question: what is the state of the fact? We classify each fact into one of five knowledge profiles: (1) encoding failure, (2) recall failure, (3) direct recall, (4) recall with thinking, and (5) inference without encoding. These profiles provide a more informative diagnosis than question-level accuracy alone.
The classification is based on whether the fact is encoded and how accessible it is: Cannot be recalled, can be directly recalled, or can be recalled only with thinking (eliciting intermediate computations before the final answer, including chain-of-thought prompting and thinking-optimized LLMs ).
Five knowledge profiles that characterize facts.
We operationalize this with three behavioral notions:
Encoding: A model encodes a fact if it can correctly reproduce it in a pre-training-like context. In our setup, we measure this using proposition completion and contextual questioning, which place the model in contexts similar to those in which the fact would naturally appear during pre-training (without revealing the answer), thereby priming the model to expose whether the fact is encoded.
Knowledge: A model knows a fact if it can correctly answer semantically equivalent questions about it across different phrasings, including both direct and reverse questions (e.g., if *A is B*, a direct question asks "What is B?", while a reverse question asks "What is A?").
Recall: A model recalls a fact if it knows an encoded fact. If it recalls the fact without thinking, we refer to this as direct recall. If it knows a fact that is not encoded, we refer to this as inference without encoding. This occurs only when thinking is enabled and the model relies on other encoded facts and performs multi-hop reasoning or educated guesses.
Top : We extract facts from Wikipedia, a predominant source of pre-training data. Left : We measure encoding by prompting the LLM to reproduce facts within their original context. Right : We measure knowledge by asking questions across varied phrasings and relational directions, with and without thinking.
To operationalize knowledge profiling, we constructed WikiProfile , a benchmark designed to measure factuality on naturally occurring facts. WikiProfile is constructed using a fully automated pipeline powered by a prompted LLM, Gemini-2.5-Pro with thinking. Prompts were developed through manual optimization on a small held-out subset. We extract candidate facts from Wikipedia pages by identifying facts: a proposition involving an ordered pair of entities (subject and object), where the subject appears first in the document. Each fact is paired with 10 tasks: two for encoding, four for knowledge evaluation, and four multiple-choice variants for recognition.
We generate direct and reverse questions through a three-step process of generation, refinement, and filtering, ensuring that each question is unambiguous, specific, minimal, and has a unique answer. All questions undergo filtering grounded in a search engine. We discard cases where multiple answers are returned or clarification is needed. After this automated filtering and a final manual validation step, the benchmark contains 2,150 facts.
A fully automated pipeline based on prompted LLMs. Left (purple): Fact extraction and construction of the proposition completion task. Center (red and blue): Construction of direct and reverse questions via generation, refinement, and filtering. Right (green): Creation of remaining questions (natural phrasing, contextual, and multiple-choice versions) based on the direct/reverse pairs.
We evaluate 13 LLMs. Each model is evaluated both with and without thinking. For each model, fact, and task, we sample eight responses. Responses are graded automatically by prompted LLM autoraters (more details in the paper ), producing approximately 4.5 million responses.
Main result: Recall, not encoding, is the bottleneck
Across the frontier LLMs (Gemini-2.5-Pro, Gemini-3-Pro and Flash, GPT-5), factual encoding is close to saturation, but recall is not. For Gemini-3-Pro and GPT-5, 95–98% of facts are encoded, yet these models still fail to directly recall 26–34% of facts. Even with thinking, they still fail on 11–12% of facts . This means that in frontier models, factual errors increasingly come not from absent knowledge, but from knowledge that is stored and not reliably accessible. In other words, the bottleneck is shifting from knowledge acquisition to knowledge utilization.
Scaling reinforces this picture. In the Gemma 3 family, larger models show far fewer encoding failures, but recall failures remain substantial and become a larger share of the remaining errors. Scaling improves what the model stores more than it improves what the model can access.
Distribution of the five profiles across 13 LLMs (percentages). The black line marks potential knowledge. As shown, encoding failures decrease sharply with scale, while recall failures persist even in frontier models.
Our results suggest that recall is tightly coupled to the conditions under which a fact was learned. When the query diverges from the training-time context, phrasing, or ordering in which the fact was encountered, recall becomes harder. We highlight two cases where this happens systematically.
Rare facts are encoded, but hard to recall
Prior work has shown that LLMs struggle with long-tail (rare) facts, often framing this as a problem of model capacity. Our results suggest a complementary picture. When we compare low-popularity and high-popularity facts, we find that rare facts are encoded at rates close to popular facts. The gap in encoding is relatively modest; however, the gap in recall is larger. This reframes the long-tail problem: Many rare facts are not absent from the model's parameters. They are present, but difficult to access. The bottleneck has shifted from knowledge acquisition to utilization.
We compare two popularity tiers (bottom 20% vs. top 20%) in terms of encoding rates and direct recall rates. The Δ indicates the gap between tiers. As shown, it is narrow for encoding but wide for recall. See our paper for the results of all LLMs.
Reverse questions are verifiable, but hard to recall
We also revisit the reversal curse : when LLMs know "A is B" but can't answer "What is B?". At first glance, this could suggest that LLMs lack bidirectional knowledge. But our results suggest a refinement of this view. In open-ended generation (i.e., recall), reverse questions are consistently harder than direct questions. In multiple-choice verification (i.e., recognition), however, reverse questions are no harder than direct ones, and are often easier. This dissociation matters. If a model can recognize the correct answer when it is presented among distractors, but cannot generate it in a reverse query, then the issue is not simply that the bidirectional knowledge is missing. Rather, the fact appears to be encoded, and even recognizable, but difficult to recall when the query direction departs from how the fact was encountered during training. The reversal curse is a recall problem.
We compare direct and reverse questions across two tasks: verification (multiple-choice) and generation (closed-book). The Δ denotes the gap between the direct and reverse settings. LLMs handle reverse questions effectively in verification but struggle in generation.
Thinking as a recovery mechanism
We now turn to the question of what enables the recovery of otherwise inaccessible knowledge. To this end, we examine the potential of thinking to fill this role. Thinking improves recall most strongly exactly where direct recall is weakest. The gains are especially pronounced for rare facts and reverse questions, narrowing both the popularity gap and the directionality gap.
We examine the impact of thinking on recall (knowing encoded facts). Left : we compare two fact popularity tiers (bottom 20% vs. top 20%). Right : we compare direct and reverse questions.The popularity or directional gaps are denoted by Δ (no thinking) and ΔT (with thinking).
As shown, thinking narrows the gaps ( ΔT < Δ).
More specifically, in thinking-optimized models, thinking recovers roughly 40–65% of encoded-but-not-directly-known facts. By contrast, it helps much less on facts that are not encoded. This pattern suggests that thinking primarily acts as a recall-facilitation mechanism: it helps the model access facts it already encoded, rather than mainly deriving answers through complex multi-step reasoning. That said, thinking is not free. It carries a computational cost, and it remains unclear how to determine exactly when a model should invoke it.
We report the percentage of not-known facts that become known with thinking, conditioned on whether the fact is encoded (red) or not (yellow). Thinking recovers 40–65% of encoded facts in thinking-optimized LLMs, but only 5–15% of non-encoded facts.
Knowledge profiling enables us to precisely diagnose factual behavior in LLMs. Applying this methodology to Wikipedia facts, our results suggest a shift in the way we should think about factual errors in frontier LLMs. If encoding is already near saturation, then further gains in factuality may come less from scaling (of model size or data). By showing that thinking can recover a substantial fraction of encoded-but-not-directly-known facts, the next improvements in factuality may come not only from better knowledge acquisition, but from better utilization of knowled

[truncated]
