---
source: "https://www.jay.ai/blog/llms-are-not-a-black-box"
hn_url: "https://news.ycombinator.com/item?id=48377631"
title: "LLMs are not the black box you were promised"
article_title: "LLMs are not the Black Box you were promised · Jay Hack"
author: "_jayhack_"
captured_at: "2026-06-03T00:34:26Z"
capture_tool: "hn-digest"
hn_id: 48377631
score: 12
comments: 1
posted_at: "2026-06-02T23:27:30Z"
tags:
  - hacker-news
  - translated
---

# LLMs are not the black box you were promised

- HN: [48377631](https://news.ycombinator.com/item?id=48377631)
- Source: [www.jay.ai](https://www.jay.ai/blog/llms-are-not-a-black-box)
- Score: 12
- Comments: 1
- Posted: 2026-06-02T23:27:30Z

## Translation

タイトル: LLM は約束されたブラックボックスではありません
記事のタイトル: LLM は約束されたブラック ボックスではありません · Jay Hack
説明: メカニズムの解釈可能性は大幅に進歩しました。 Anthropic の「大規模言語モデルの生物学」を巡るツアー。

記事本文:
LLM は約束されたブラック ボックスではありません · Jay Hack すべての書き込み LLM は約束されたブラック ボックスではありません
機械的な解釈可能性は大きく進歩しました。 Anthropic の「大規模言語モデルの生物学」を巡るツアー。
LLM は、約束された「ブラック ボックス」ではありません。
機械的な解釈可能性 (ニューラル ネットワークを覗いてその内部動作をリバース エンジニアリングする) は、大きな進歩を遂げました。 Anthropic の『On the Biology of a Large Language Model』(2025) は、その取り組みにおける画期的な作品です。以下は、彼らの進捗状況とそれに関連する考えの概要です。
LLM は実際に何を「考えている」のでしょうか?
LLM が「考えている」ことをどのように理解できるでしょうか?そうすることは明らかに非常に価値があり、モデルの動作を制御したり、危険な意図を検出したりできる可能性があります。
しかし、重ね合わせがあるため、個々のニューロンの活性化を単に観察するよりもはるかに困難です。単一のニューロンは多くの無関係な概念に関与しており、特定の概念は多くのニューロンに塗りつぶされています。 1 つの単位だけで意味を読み取ることはできません。創造性を発揮する必要があります。
1 つのアプローチ: 2 番目のモデルをトレーニングして個別の概念を特定し、それらの概念がフォワード パスの過程でどのように相互作用するかを監視します。
Anthropic の回路トレース技術は、「置換」モデルをトレーニングして、基本モデルの MLP 層の出力をまばらに再作成します。これにより、基本モデルのアクティベーションが一連のまばらな特徴に効果的に分解されます。そして、これらの特徴が、「テキサス」や「オリンピック」など、人間が容易に識別できる高レベルの概念に対応していることがわかります。
人間が解釈できるこれらの特徴を取得したら、順方向パス中にそれらがどのように相互作用するかを追跡することで、それらを因果関係のあるクラスターにグループ化し、計算の配線図を構築できます。
モデルは本当に理性を持っています

複数のステップで
これを実際に実行すると、モデルが中間概念を介して真の複数ステップの推論を行うのを見ることができます。このモデルは、詩を計画するときに、将来の韻の候補を「先取りして考える」こともできます。
「ダラスを含む州の首都はどこですか」と質問すると、次の順序で観察できます。
ダラス機能がアクティブになり、
これによりテキサスの特徴が点灯します。
するとオースティンが明るくなります。
これが高レベルの概念間の意味論的関係を追跡していることはかなり明らかであるように思われます。そしてその際、一部の哲学者が「高等推論」と表現するものに似た、一種の擬似記号推論を実行しています。
この現象は言語モデルだけに当てはまるわけではありません。 AlphaZero のような MCTS ベースのシステムも、人間が認識する概念に収束します。
DeepMind (2022) は、AlphaZero が「チェック」や「駒を固定する」などの人間のチェスの概念に沿った中間表現を、人間のチェスの知識が提供されることなく完全に単独で学習したことを示しました。
より良い理解 → より良いアルゴリズム
モデルの暗黙的な推論を分析することは、より優れた学習アルゴリズムを設計するのに役立ちます。
例: Claude 3.5 Haiku は、人間の暗算にきれいに対応しない小さな整数の加算アルゴリズムを学習しました。問題を複数の並列経路に分割し、正確な 1 桁と一緒に大まかな大きさを計算し、記憶された「ルックアップ テーブル」機能に基づいてそれらを再結合します。
自然な疑問は次のとおりです。これを特定して、モデルをより良いアルゴリズムに「導く」ことができるでしょうか?
モデルには「潜在意識」がある
モデル自体には、回路トレースによって明らかにされる根底にある思考プロセスに対するメタ認知的な洞察が必ずしも含まれているわけではないことは注目に値します。 2 つの数値をどのように加算するかを説明してもらいます。

きちんとした人間流の手順を説明しますが、実際に実行されたアルゴリズムではありません。
良くも悪くも、モデルにはある程度の潜在意識が備わっています。そしてそれこそが、私たちが覗き込むことができるものなのです。
メカニズムの解釈可能性は、スコアボードに主要な W が含まれる魅力的で急速に発展している分野です。
10 年前に機械学習の教授が言ったこととは異なり、これはある意味、これまで私たちがモデルから抽出した中で最も多くの洞察をもたらします。そして、その影響は、モデルの不正動作の特定、ステアリング、さらにはより良い学習アルゴリズムの設計にとっても重要です。
元のスレッドについては、 X の投稿を参照してください。研究の詳細については、Anthropic の論文をお読みください。

## Original Extract

Mechanistic interpretability has made major strides. A tour through Anthropic's "On the Biology of a Large Language Model."

LLMs are not the Black Box you were promised · Jay Hack All writing LLMs are not the Black Box you were promised
Mechanistic interpretability has made major strides. A tour through Anthropic's "On the Biology of a Large Language Model."
LLMs are not the "black box" you were promised.
Mechanistic interpretability — peering into a neural network to reverse engineer its inner workings — has made major strides. Anthropic's On the Biology of a Large Language Model (2025) is a landmark in that effort. What follows is a summary of their progress and some related thoughts.
What is an LLM actually "thinking"?
How can we understand what an LLM is "thinking"? It's clearly very valuable to do so — it could enable steering model behavior, detecting dangerous intent, and more.
But it's much harder than simply observing individual neuron activations, because of superposition : a single neuron participates in many unrelated concepts, and any given concept is smeared across many neurons. You can't just read meaning off one unit. You need to get creative.
One approach: train a second model to identify discrete concepts, then monitor how those concepts interact over the course of a forward pass.
Anthropic's circuit tracing technique trains a "replacement" model to sparsely recreate the outputs of the base model's MLP layers. This effectively decomposes the base model's activations into a set of sparse features — and it turns out these features correspond to high-level concepts that humans can readily identify, like "Texas" or "the Olympics."
Once you have these human-interpretable features, you can group them into causally-linked clusters by tracing how they interact during the forward pass — building up a wiring diagram of the computation.
Models really do reason in multiple steps
When you run this in practice, you can watch models engage in genuine multi-step reasoning via intermediary concepts. The model will even "think ahead" to future rhyme candidates when planning a poem.
Ask it "what is the capital of the state containing Dallas" and you can observe, in order:
the Dallas feature goes active,
which causes the Texas feature to light up,
which then causes Austin to light up.
It seems fairly clear that this is tracing semantic relationships between high-level concepts — and in doing so, performing a kind of pseudo-symbolic inference, similar to what some philosophers would describe as "higher reasoning."
This phenomenon doesn't only apply to language models. MCTS-based systems like AlphaZero also converge on concepts that humans recognize.
DeepMind (2022) showed that AlphaZero learned intermediary representations aligning with human chess concepts such as "in check" and "pinning a piece" — entirely on its own, with no human chess knowledge supplied.
Better understanding → better algorithms
Breaking down a model's implicit reasoning can help us design better learning algorithms.
For example: Claude 3.5 Haiku learned an algorithm for small-integer addition that does not cleanly map to human mental math. It splits the problem into multiple parallel pathways — computing a rough magnitude alongside the precise ones-digit — and recombines them, leaning on memorized "lookup table" features.
The natural question follows: can we identify this, then "guide" the model toward a better algorithm?
The model has a "subconscious"
It's worth noting that the model itself does not necessarily have metacognitive insight into the underlying thinking process uncovered by circuit tracing. Ask it to explain how it added two numbers and it will narrate a tidy, human-style procedure — which is not the algorithm it actually ran.
For better or worse, the model has some level of subconscious. And that's precisely what lets us peer in.
Mechanistic interpretability is a fascinating, fast-developing line of work with major Ws on the scoreboard.
Contrary to what your ML professor may have told you a decade ago, in some ways this is now the most insight we've ever extracted from a model. And the implications are significant — for identifying model misbehavior, for steering, and even for designing better learning algorithms.
For the original thread, see the post on X . For the full research, read Anthropic's paper .
