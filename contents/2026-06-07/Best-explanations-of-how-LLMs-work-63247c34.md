---
source: "https://vorushin.github.io/blog/best-explanations-llms"
hn_url: "https://news.ycombinator.com/item?id=48435356"
title: "Best explanations of how LLMs work"
article_title: "Best explanations of how LLMs work | Roman Vorushin"
author: "burakabo"
captured_at: "2026-06-07T15:34:41Z"
capture_tool: "hn-digest"
hn_id: 48435356
score: 2
comments: 0
posted_at: "2026-06-07T14:45:48Z"
tags:
  - hacker-news
  - translated
---

# Best explanations of how LLMs work

- HN: [48435356](https://news.ycombinator.com/item?id=48435356)
- Source: [vorushin.github.io](https://vorushin.github.io/blog/best-explanations-llms)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T14:45:48Z

## Translation

タイトル: LLM の仕組みについての最良の説明
記事のタイトル: LLM の仕組みについての最良の説明 |ロマン・ヴォルシン
説明: 私は、LLM がどのように訓練され、機能するかについての一連の最良の説明 1 を維持しています。ここでの「LLM」とは、2026 年に価値を生み出すフロンティア モデルを指す広義の用語です。LLM には単なる言語モデルよりも多くのコンポーネントがありますが、強力な言語モデルが必要な核となります。普遍的な説明者 考察：人間は
[切り捨てられた]

記事本文:
ロマン・ヴォルシン
RSS
×
GitHub
リンクトイン
再開
LLM の仕組みについての最良の説明
私は、LLM がどのように訓練され、機能するかについての一連の最良の説明 1 を維持しています。
ここでの「LLM」とは、2026 年に価値を生み出すフロンティア モデルを指す広義の用語です。LLM には単なる言語モデルよりも多くのコンポーネントがありますが、強力な言語モデルが必要な核となります。
考察：人間は普遍的な説明者である（ドイツ語） - 彼らは創造性を使って世界を説明します。
推測: LLM は非生物学的基質における普遍的な説明子です。
これは、LLM は確率論的なオウムである、つまり模倣するだけで、新しい知識を理解したり創造したりすることはないという主張の反対です。
人間の複雑なミームを理解できるように進化しました。
類人猿は単純なミームしか理解できず、生まれつきの語彙が限られています。
対照的に、複雑な人間ミーム - キャンプファイヤーの作り方: 豊富な語彙が必要であり、学習者は教師が説明しようとしている内容について仮説を繰り返し生成し改善する必要があります。
たまたま普遍的なリーチを持っていました。
現実についてより良い説明を作成し、それらを批判し、テストし、さらにより良い説明を作成します。
ルネッサンス以来の科学の進歩を推進します。
当初、創造性は、不明確なユーザー プロンプトの背後にある意図 (喜び) を理解するために不可欠でした。
現在、それは普遍的なリーチを持っているようです - モデルは複雑なエージェントループを何時間も駆動し、指定された目標に向かって反復することができます。
半導体、トークン、2026 年の IPO の需要を後押しします。
Chris Olah の枠組み 2 の精神に基づく生物学的な比喩:
ニューラル ネットワークは成長するものであり、プログラムされたり構築されたりするものではありません (ただし、ここでのプログラミングに最も近いのは微分可能な関数型プログラミングです 3 )。
ニューラル ネットワーク アーキテクチャは、回路が成長する足場のようなものです。
損失目標はネットワークを導く光です

望ましい結果に向けて成長すること。
結果として得られるネットワークは生物学的実体、つまり私たちが研究できる生物に似ています。
損失: 次のトークンの分布を予測します (最も可能性の高い単一のトークンだけでなく)。このモデルには多くのステップ (レイヤー) と大きな潜在的なスクラッチパッド (トークンの埋め込み + 残留ストリーム) があります。
すべての入力データを記憶するだけでは十分ではありません。通常、トレーニング セットはモデル パラメーターの数よりもはるかに大きくなります。
モデルはデータ内の再帰的な自己類似性を発見し、トレーニング データが表す現実についてのより深い真実を学習します。
私たちがプログラムすることなく、モデルは暗黙的に、これらのより深い真実を学習するのに最適なアルゴリズムを発明します (微分可能関数型プログラミングとソフトウェア 2.0 を参照)。
スクラッチパッド、電卓、Web 検索など、レイヤー間の潜在的な空間で考えることを超えて、外部ツールを使用する練習を提供します。
ロング RL は、長期的なエージェント タスクに不可欠な自己批判とエラー修正のための効率的な戦略を学ぶ機会です。
事前トレーニングよりも計算効率が大幅に低くなりますが、これは最適な RL 軌道でのトレーニング中に修正することができます。
一貫したペルソナ 4 を偽装する LLM は、長期的なタスクに優れているようです。ペルソナのトレーニングがなければ、LLM は人間に似たさまざまなエンティティの間を切り替えることができます。クリエイティブな文章を書くのには適していますが、物事を成し遂げるのにはあまり役に立ちません。
LLM は、(量子ではなく) 古典的な計算のみを使用することで優れた成果を上げています。人間の脳も古典的な計算によって強化されているようです。
人間の脳は非常にエネルギー効率が良いと考えられています。これは、継続学習、スパース性、モジュール性のための優れたソリューションがあれば、LLM も適度に効率的であるという強力な証拠です。
データの移動 (外部 -> チップ) は、計算よりも難しくなります 5 。
ほとんどのフロリダ

LLM の os は行列の乗算です。 GPU/TPU チップの大部分は、入出力行列とそれらを乗算するシストリック アレイを保持します。
最近のアルゴリズムの改善の多くは、スパース計算またはスムーズな最適化の考えに基づいています。
もう 1 つの重要なベクトルは、精度の低下です。ウェイトとアクティベーションを、表現するために多くのビットを必要としない適切な領域に保持することで、エネルギーとダイ領域を大幅に節約します。
汎用人工知能に対する完璧なソリューションについての非常に強い信念は、特に盲目的に保持されている場合、助けとなるよりもハンデを与える可能性があります。それにもかかわらず、役立つアプローチと役に立たないアプローチを追跡し、より多くのデータとコンピューティングに関して以前はうまくいかなかったアプローチを再検討することが重要です。
シンプルなアイデアが最も拡張性が高いようです。しかし、それらに到達するには、多くの場合、最初に複数のアドホック/狭い改善を試し、後でより一般的なスレッドを確認する必要があります。
David Deutsch、『無限の始まり: 世界を変える説明』優れた説明 - 世界を正確に説明しながら、変更するのが難しい大胆で創造的な推測。 ↩
特定の比喩は、このインタビューの最後の 1 時間のものです。 ↩
Chris Olah、ニューラル ネットワーク、型、および関数型プログラミング (2015 年): 「まったく新しい種類のプログラミング、微分可能な関数型プログラミングのようなものだと感じます。」ヤン・ルカンは後に「微分可能プログラミング」というより広い用語を普及させました。 ↩
Sam Marks、Jack Lindsey、Christopher Olah ペルソナ選択モデル: AI アシスタントが人間のように振る舞う理由 ↩
TPU の詳細 ↩ の「TPU とエネルギー効率 (TPUv4)」を参照してください。

## Original Extract

I maintain a set of best explanations1 of how LLMs train and work. “LLMs” here is a broad term for the frontier models that create value in 2026. They have more components than just a language model, but powerful language models are their necessary core. Universal explainers Considering: humans are
[truncated]

Roman Vorushin
RSS
X
GitHub
LinkedIn
Resume
Best explanations of how LLMs work
I maintain a set of best explanations 1 of how LLMs train and work.
“LLMs” here is a broad term for the frontier models that create value in 2026. They have more components than just a language model, but powerful language models are their necessary core.
Considering: humans are universal explainers (Deutsch) - they use creativity to explain the world.
Conjecture: LLMs are universal explainers on a non-biological substrate.
That’s the opposite of the claim that LLMs are stochastic parrots - that they only imitate, without understanding or creating new knowledge.
Evolved to understand complex human memes.
Apes understand only simple memes, limited to a small inborn vocabulary.
A complex human meme, by contrast - how to build a campfire: it needs a rich vocabulary and requires the learner to iteratively generate and improve hypotheses about what the teacher is trying to explain.
Happened to have universal reach.
Create better explanations of reality, critique them, test them, create better ones still.
Powers scientific progress since the Renaissance.
Initially, creativity was crucial for understanding the intent behind underspecified user prompts (hepfulness).
Now it seems to have universal reach - models can drive complex agentic loops for hours, iterating toward specified goals.
Powers demand for semiconductors, tokens, and 2026 IPOs.
A biological metaphor, in the spirit of Chris Olah’s framing 2 :
Neural networks are grown, not programmed or built (though the closest thing to programming here is differentiable functional programming 3 ).
A neural-network architecture is like a scaffold upon which circuits grow.
The loss objective is the light that guides the network’s growth toward the desired outcome.
The resulting network is akin to a biological entity - an organism that we can study.
Loss: predict the distribution of next tokens (not just the single most-probable token). The model has many steps (layers) and a large latent scratchpad (token embeddings + residual streams).
Memorizing all the input data isn’t enough - the training set is usually much larger than the number of model parameters.
The model discovers recursive self-similarities in the data, learning the deeper truths about the reality the training data describes.
Implicitly - without us programming it - the model invents whatever algorithm works best for learning these deeper truths (see differentiable functional programming and Software 2.0 ).
Provides practice using external tools, beyond thinking in the latent space between layers: scratchpad, calculator, web search, and so on.
Long RL is a chance to learn efficient strategies for self-critique and error correction - critical for long-horizon agentic tasks.
Much less compute-efficient than pre-training, but this can be fixed by mid-training on the best RL trajectories.
LLM impersonating a coherent persona 4 seems to be better at long horizon tasks. Without the persona training the LLM can switch between a wide range of human-like entities. Good for creative writing, not really helpful for getting things done.
LLMs are doing great by using only classical (as opposed to quantum) computation. Human brains seem to be powered by the classical computation as well.
Human brains are deemed quite energy efficient - strong evidence that with the good solutions for continuous learning, sparsity, modularity the LLMs can be reasonably efficient as well.
Data movement (outside -> chip) scales harder that computation 5 .
Most flos in LLMs are matrix multiplications; the large part of GPUs/TPUs chips holds the input/output matrices and systolic arrays that multiply them.
Many recent algorithmic improvements build on the idea of sparse computation or smooth optimization.
Another important vector: lowering precision. Keep weights and activations in nice regions that don’t require many bits to represent - save a lot of energy and die area.
Very strong beliefs about the perfect solution to artificial general intelligence may handicap more than help, especially when blindly held. Nonetheless, it’s important to keep track of approaches that help, approaches that don’t help, and revisit the previously non-working approaches on more data and compute.
Simple ideas seem to scale the best. But getting to them often requires to trying multiple ad-hoc/narrow improvements first, and only later seeing a more general thread.
David Deutsch, The Beginning of Infinity: Explanations That Transform the World . Good explanations - bold, creative conjectures that are hard to vary while still precisely accounting for the world. ↩
The particular set of metaphors is from the last hour of this interview . ↩
Chris Olah, Neural Networks, Types, and Functional Programming (2015): “It feels like a new kind of programming altogether, a kind of differentiable functional programming.” Yann LeCun later popularized the broader term “differentiable programming.” ↩
Sam Marks, Jack Lindsey, Christopher Olah The Persona Selection Model: Why AI Assistants might Behave like Humans ↩
See “TPUs and Energy Efficiency (TPUv4)” in TPU Deep Dive ↩
