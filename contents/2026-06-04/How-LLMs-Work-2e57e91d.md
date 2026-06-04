---
source: "https://practical-leaders.com/articles/how-llms-work"
hn_url: "https://news.ycombinator.com/item?id=48403751"
title: "How LLMs Work"
article_title: "How LLMs work | Practical Leaders"
author: "ivorc"
captured_at: "2026-06-04T21:37:30Z"
capture_tool: "hn-digest"
hn_id: 48403751
score: 2
comments: 2
posted_at: "2026-06-04T19:52:02Z"
tags:
  - hacker-news
  - translated
---

# How LLMs Work

- HN: [48403751](https://news.ycombinator.com/item?id=48403751)
- Source: [practical-leaders.com](https://practical-leaders.com/articles/how-llms-work)
- Score: 2
- Comments: 2
- Posted: 2026-06-04T19:52:02Z

## Translation

タイトル: LLM の仕組み
記事のタイトル: LLM の仕組み |実践的リーダー
説明: 大規模言語モデル (LLM) の内部を深く掘り下げます。 LLM がどのように機能するかを理解することは、LLM の能力と限界の両方を理解するのに役立ちます。

記事本文:
すべての記事 LLM の仕組み AI が製品デリバリーをどのように変えるかについての予測 位置付けの重要性 責任 (3/3): 診断および実装ガイド 責任 (2/3): 善良な人々が悪いシステムに閉じ込められるとき 責任 (1/3): 指揮統制を超えて すべての記事 LLM の仕組み
大規模言語モデル (LLM) の内部を深く掘り下げます。 LLM がどのように機能するかを理解することは、LLM の能力と限界の両方を理解するのに役立ちます。
アイヴァー・コールドウェル – 2026 年 5 月 30 日 – 66 分で読む
#ai #エンジニアリング
多くの人にとって、AI は Claude、ChatGPT、Gemini などのチャットボットと同義です。これらは大規模言語モデル (LLM) によって強化されており、それらがどのように機能するかを理解することは、その能力とその限界を理解するのに役立ちます。ここでは、主要な構成要素とそれらがどのように組み合わされるかについて説明します。
LLM がシーケンス内の次のトークンを予測することによって機能することは多くの人が知っていますが、経験に基づいてこの説明を理解するのは困難です。彼らの出力の豊かさは、賢いデザインと巨大なスケールの産物であり、この記事では、その作成に何が必要なのかを明らかにします。
この記事では、LLM がどのように機能するかを根本から説明します。その優れた能力を考慮すると、アーキテクチャは多くのソフトウェア システムと比較して比較的単純です。その威力は、複雑な設計ではなく、最大規模のモデルに 1 兆を超えるパラメーターを備えたスケールによってもたらされます。
しかし、これは依然として大きなトピックであり、この記事もそれに応じて長くなります。目的は、AI や機械学習の背景がない一般の技術者でも理解できるようにすることです。
この記事は、重要なポイントの要約から始まります。パート 1 と 2 では、概念的な理解を構築することに重点を置いています。パート 1 では推論、つまりモデルが出力を生成する方法について説明します。パート 2 では、トレーニング、つまり重みが値を取得する方法について説明します。パート3はかなり進みます

現実世界の実用性、最適化、エッジケースをさらに深く掘り下げます。必要に応じて、パート 3 をスキップして、要約と結論に直接進むこともできます。付録では、本文で簡潔にまとめられた数学的モデル、ニューラル ネットワーク、行列乗算などのトピックの背景を提供します。
かなり単純ですが、おそらく馴染みのない数学をいくつか含めました。これらのセクションを読み飛ばしても問題なく理解できます。
これは私が学んだことを再現する試みですが、私は専門家ではありません。興味深いことに、私が 25 年前に取得したニューラル ネットワークの修士号は、今でも多少は役に立ちますが、新しいこともたくさんあります。
この記事では、次の点について詳しく説明します。
LLM は一度に 1 トークンずつ出力を生成します。各トークンは、先読みせずに前のコンテキストを調べることによって生成されます。
流暢さ、知識、明らかな推論はすべて、モデルが学習したパターンから現れます。膨大な量のテキストをトレーニングすることで、モデルは言語の構造、言語にエンコードされた知識、人間が書いたテキストに現れる推論のパターンを学習できます。
事前トレーニングでは、膨大なコーパス (テキスト本体) にわたる次のトークンを予測するようにモデルを学習します。トレーニング後は、指示に従い、有害な要求を断り、ツールを使用するなど、アシスタントとしての役割を果たします。
LLM は会話間の記憶を持たず、ツールが与えられない限り世界にアクセスできません。
アテンションのメカニズムは、モデルがコンテキストをどのように利用するかということです。各層で、すべてのトークンが先行するトークンから関連情報を取得します。注意には因果関係があります。各トークンは過去のみを参照するため、左から右への生成が可能になります。
フィードフォワード層はパターン検索メモリのように機能し、モデルが処理している内容を認識し、トレーニング中にエンコードされた関連知識を提供します。
これら 2 つの操作を積み重ねる

何十ものレイヤーを重ねることで、豊かな出力が生まれます。 「理解」が起こる単一の場所はありません。それはネットワーク全体に分散されます。
制限はアーキテクチャから直接生じます。モデルは間違っている自信に満ちたテキストを生成する可能性があり、算術などの正確な記号操作を必要とするタスクに苦労し、トレーニング データの内容に制限されます。出力は確率的であり、決定的ではありません。
モデルを大きくすると知識は増えますが、推論能力への影響は少なくなります。広範なポストトレーニングと拡張推論（「思考モード」）は、推論力を高めるためにこれまでに発見された最も効果的な方法です。
LLM には固有の制限があり、他のメカニズムと組み合わせることによってのみ修正可能です。
LLM は、少数の繰り返しコンポーネントから構築された数学モデルです。それぞれについて詳しく説明する前に、それらがどのように接続されているかを確認するのに役立ちます。 (数学モデルの説明については、付録 1 を参照してください。)
動作中、LLM は一連のトークンを入力として受け取り、次にどのトークンが来るべきかについての最良の推測を出力します。これは一度に 1 つのトークンごとに行われます。新しいトークンがそれぞれコンテキストに追加され、プロセスが繰り返され、単語ごとに流暢なテキストが生成されます。このプロセスは 推論 と呼ばれます。
このパートとパート 2 では、概念的な理解を構築するための有用な出発点を提供する純粋な変圧器モデルと呼ばれるものに焦点を当てます。パート 3 では、代替モデル アーキテクチャについて触れます。
これらのトランスフォーマー モデルは内部的に、生成されたトークンごとに 1 回以下の手順に従って入力を処理します。それぞれについては後続のセクションで説明します。
入力内の各トークンはトークン ID に変換され、次に埋め込みベクトル (埋め込みベクトルの意味論的表現である数値のリスト) に変換されます。

トークン。
この埋め込みベクトルは、各トークンの残差ストリームの初期値として使用され、個別に学習された重みを使用して、構造が同一の数十の層を通過します。各レイヤーには 2 つのサブレイヤーがあります。
アテンションサブレイヤーは、各トークンがシーケンス内の先行するトークンから情報を取得できるようにします。
フィードフォワード サブレイヤは、小さなニューラル ネットワークを通じて各トークンを独立して処理し、関連情報を追加します。
各サブレイヤーで、出力が残りのストリームに追加され、ストリームが強化されます。
最終層の後、残差ストリームの最後の位置を使用して次のトークンが予測されます。
トレーニング中に学習されるパラメータは次のとおりです。
テキストを埋め込みに変換する変換。
アテンションサブレイヤーとフィードフォワードサブレイヤー内の重み。各レイヤーごとに個別に学習されます。
最後のレイヤーの最終残差ストリームを次のトークンの予測に変換する変換。
トレーニングが完了すると、ウェイトは固定されます。モデルは個々の会話から学習しません。チャット セッションで何も言われても、基礎となるパラメーターは更新されません。
LLM はテキストを 1 文字ずつ、または単語ごとに処理しません。代わりに、テキストはトークン、つまりおおよその音節から単語のサイズのチャンクに分割されます。
最も一般的なアプローチは、バイト ペア エンコード (BPE) です。これは、個々の文字から開始し、目標語彙サイズ (現在のフロンティア モデルでは最大 200,000 トークン) に達するまで、最も頻繁に使用される隣接するペアを繰り返しマージします。一般的な単語は単一のトークンとして扱われます。まれな単語は断片に分割されます。たとえば、単語「tokenisation」は [ token , isa , tion ] に分割される可能性がありますが、「cat」は単一のトークンになります。数字、句読点、空白もトークン化され、トークンには先頭にスペースが含まれることがよくあります。

単語の境界を知らせるため。マージ ルール自体はコーパスから頻度によって学習されますが、これはモデルのトレーニング前に 1 回だけ発生し、その後トークナイザーはフリーズされます。分割テキストへの適用は決定論的であり、トークナイザーは LLM 自体から分離されていると見なされます。
モデルは文字や単語ではなく、一連のトークンを認識します。異常なスペル、造語、または非常に専門的なテキストはトークン化が不十分であり、モデルがトレーニング中にあまり認識していない多くのフラグメントが生成される可能性があります。トークン化の段階で算術演算が失敗する可能性があります。「1234」は 1 つのトークンである可能性がありますが、「12345」は別の方法で分割される可能性があり、モデルは個々の数字に直接アクセスできません。
LLM のコンテキスト ウィンドウは、入力と出力を含め、一度に処理できるトークンの最大数です。おおよその目安として、100 トークンは約 75 個の英単語に相当します。現在のフロンティア モデルは最大 2,000,000 のトークン ウィンドウをサポートしています。これは、ブック全体またはコードベース全体を 1 つのコンテキストに適合できることを意味します。語彙 (モデルが認識しているトークンの完全なリスト) はトレーニング時に固定されます。各トークンには一意の整数 ID が割り当てられ、これらの ID が次のステージに流れます。
トークンを埋め込みベクトルに変換する
トークン ID 自体は単なる整数です。モデルには、トークンが語彙内のどのエントリに対応するかだけでなく、トークンが何を意味するかを捉える、より豊富な表現が必要です。
この表現は埋め込みベクトルと呼ばれます。その要素は一緒になってトークンの意味をエンコードしますが、単一の要素が人間が判読可能な 1 つのプロパティにきちんと対応することはありません。トークンが埋め込みに変換される方法はトレーニング中に学習され、各モデルに固有です。同様のコンテキストに出現するトークンは、同様の埋め込みを持ちます。これらは、通常 7,000 ～ 16,00 の高次元空間内のベクトルとして表現できます。

0次元。意味は個々の軸ではなく、この空間を通る方向によって伝えられ、モデルには寸法よりもはるかに明確な特徴が詰め込まれています。
埋め込み空間の幾何学形状には意味があります。「キング」と「クイーン」のベクトルは、「自転車」のベクトルよりも互いに近いです。これにより、関連する単語や概念全体でモデルを一般化できるようになります。個々の次元は、ほとんどの場合、単独では解釈できません。空間内の方向が何を意味するかは、訓練によって完全に固定されます。研究者は埋め込みを「調査」してこの構造の一部を復元できますが、実際には不透明度は問題ではありません。これらの埋め込みはモデルの内部機構の一部にすぎず、直接読み取る必要があるものではありません。
実際には、この変換は埋め込みテーブル (語彙内のトークンごとに 1 行を持つ大きな行列) を使用して行われます。各行は、そのトークンの埋め込みを表す数値のベクトルです。トークンの埋め込みを検索するには、このテーブルから関連する行を選択するだけです。
この段階では、埋め込みには位置に関する情報が含まれていません。 「the cat sit」と「sat the cat」では、各トークンに対して同一の埋め込みが生成されますが、これは語順が重要なタスクでは問題となります。これを解決するために、各ベクトルに位置情報を付加します。詳細については、「トークンの位置のエンコード」を参照してください。
結果は、入力トークンごとに 1 つずつベクトルのシーケンスとなり、それぞれがトークンの意味と位置をエンコードします。このシーケンスは、モデルの最初の層に入ります。
入力がモデルを通過すると、各トークンはその埋め込みとして開始されるベクトルを運び、コンテキストと解釈を追加して各層で徐々に洗練されます。このベクトルは残差ストリームと呼ばれます。
各トークンには独自の残差ストリームがあります

。多くのレイヤーにわたって、各トークンのベクトルは、そのトークンが何であるかの表現から、次のトークンが何であるべきかの予測へと徐々にシフトしていきます。推論時には、最終層の最後の残差ストリーム ベクトルのみが次のトークンの予測に使用されます。ただし、後の層の注意は、最後の層だけでなく、前の層がすべての位置に対して生成したものに依存するため、すべてのトークンはすべての層で処理されます。
各サブレイヤーの前に、残差ストリームが正規化されます。つまり、別のオーディオ トラックのボリュームを調整して他のトラックをかき消してしまわないように、その値が予測可能な範囲内に収まるように再スケーリングされます。これにより、値が数十のレイヤーを通過する際に、値が一貫したスケールに保たれます。これがないと、トレーニング中に値が指数関数的に拡大または縮小する可能性があり、モデルを確実に最適化することができなくなります。
アテンションは、各トークンがシーケンス内で先行するトークンをクエリし、関連情報を取得できるようにするメカニズムです。これにより、モデルは参照を解決し (「犬はお腹が空いたから猫を追いかけた」では、「それ」は犬を指します)、長距離の依存関係を追跡し、各トークンの文脈に応じた理解を構築することができます。
各トークンは、シーケンス内でその前にあるトークンのみを対象とします。これは因果的（またはマスクされた）注意と呼ばれます。それは何ですか

[切り捨てられた]

## Original Extract

A deep dive into the internals of Large Language Models (LLMs). Understanding how LLMs work helps us understand both their power and their limitations.

All Articles How LLMs work Predictions for how AI will change product delivery The importance of positioning Accountability (3/3): Diagnostic and implementation guide Accountability (2/3): When good people get trapped in bad systems Accountability (1/3): Beyond command and control All Articles How LLMs work
A deep dive into the internals of Large Language Models (LLMs). Understanding how LLMs work helps us understand both their power and their limitations.
Ivor Caldwell • 30 May 2026 • 66 min read
#ai #engineering
For many, AI is synonymous with chatbots like Claude, ChatGPT, or Gemini. These are powered by Large Language Models (LLMs), and understanding how they work helps us appreciate their power and their limitations. I’ll describe here the key building blocks and how they come together.
Many know that LLMs work by predicting the next token in a sequence, but this explanation is hard to square with experience. The richness of their output is a product of clever design and immense scale, and in this article we’ll uncover what goes into creating it.
This article explains how LLMs work from the ground up. Given their impressive abilities, the architecture is relatively straightforward compared to many software systems. Their power comes from scale, not complex design, with more than a trillion parameters in the largest models.
But this is still a big topic and this article is correspondingly long. The aim is for it to be understandable by a general technical audience with no background in AI or machine learning.
The article opens with a summary of the key points. Parts 1 and 2 focus on building conceptual understanding. Part 1 covers inference: how a model generates output. Part 2 covers training: how the weights get their values. Part 3 goes much deeper into real-world practicalities, optimisations, and edge cases. You can skip Part 3 and go straight to the recap and conclusion if you prefer. The appendices provide background on topics kept brief in the main text: mathematical models , neural networks , and matrix multiplication .
I have included some maths that is fairly simple but perhaps unfamiliar. You can safely skip over these sections and still gain a good understanding.
This is my attempt to play back what I have learned, but I’m not a specialist. Interestingly, the MSc in Neural Networks I got 25 years ago is still somewhat relevant, but much is also new.
This article expands on the following points:
LLMs generate output one token at a time. Each token is produced by looking at previous context, with no lookahead.
The fluency, knowledge, and apparent reasoning all emerge from patterns the model has learned. Training on huge bodies of text lets the model learn the structure of language, the knowledge encoded in it, and patterns of reasoning that appear in human-written text.
Pre-training teaches the model to predict the next token across a vast corpus (body of text). Post-training shapes it into an assistant: following instructions, declining harmful requests, and using tools.
LLMs have no memory between conversations and no access to the world unless given tools.
The attention mechanism is how the model makes use of context: at each layer, every token pulls relevant information from preceding tokens. Attention is causal: each token only ever looks backwards, which is what makes left-to-right generation possible.
Feed-forward layers act like a pattern-retrieval memory: recognising what the model is processing and contributing relevant knowledge encoded during training.
Stacking these two operations across dozens of layers is what produces the richness of the output. There is no single place where “understanding” happens; it is distributed across the whole network.
Limitations follow directly from the architecture: the model can produce confident-sounding text that is wrong, struggles with tasks that require exact symbol manipulation such as arithmetic, and is bounded by what was in its training data. Output is probabilistic, not deterministic.
Making the model bigger increases its knowledge but has less effect on its reasoning ability. Extensive post-training and extended reasoning (“thinking mode”) are the most effective ways discovered so far to boost reasoning.
Some limitations of LLMs are inherent and will only be fixable by combining them with other mechanisms.
LLMs are mathematical models built from a small number of repeated components. Before diving into each one, it helps to see how they connect. (See Appendix 1 for an explanation of what a mathematical model is.)
In operation, an LLM takes as input a sequence of tokens , and outputs its best guess at which token should come next. This happens one token at a time: each new token is appended to the context and the process repeats, producing fluent text word by word. This process is called inference .
In this part and Part 2, we’ll focus on what are called pure transformer models as they give a useful starting point for building a conceptual understanding. Part 3 touches on alternative model architectures.
Internally, these transformer models process input by following the steps below once for each token generated. Each is explained in subsequent sections.
Each token in the input is converted to a token ID and then to an embedding vector : a list of numbers that are a semantic representation of the token.
This embedding vector is used as the initial value of the residual stream for each token, which is passed through a few dozen layers that are identical in structure, with separately learned weights. Each layer has two sublayers:
an attention sublayer lets each token draw information from preceding tokens in the sequence;
a feed-forward sublayer processes each token independently through a small neural network to add relevant information.
At each sublayer, the output is added to the residual stream, enriching it.
After the final layer, the last position in the residual stream is used to predict the next token.
The parameters that get learned during training are:
the transformations that convert text to embeddings;
the weights inside the attention and feed-forward sublayers, learned separately for each layer;
the transformation that converts the last layer’s final residual stream into a prediction of the next token.
The weights are fixed once training is complete. The model does not learn from individual conversations: nothing said in a chat session updates the underlying parameters.
LLMs do not process text character by character or word by word. Instead, text is split into tokens : chunks that are roughly syllable- to word-sized.
The most common approach is byte pair encoding (BPE), which starts with individual characters and repeatedly merges the most frequent adjacent pairs until a target vocabulary size is reached (up to 200,000 tokens for current frontier models). Common words end up as a single token; rarer words get split into pieces. For example, the word “tokenisation” might be split into [ token , isa , tion ], while “cat” would be a single token. Numbers, punctuation, and whitespace are also tokenised, and tokens often include a leading space to signal word boundaries. The merge rules themselves are learned from a corpus by frequency, but this happens once, before model training, and the tokeniser is then frozen. Applying it to split text is deterministic, and the tokeniser is considered separate from the LLM itself.
The model sees a sequence of tokens, not characters or words. Unusual spellings, made-up words, or heavily technical text may tokenise poorly, producing many fragments that the model has seen less of during training. Arithmetic can go wrong at the tokenisation stage: “1234” may be one token, but “12345” might split differently, and the model has no direct access to individual digits.
An LLM’s context window is the maximum number of tokens it can process at once, including inputs and outputs. As a rough guide, 100 tokens corresponds to around 75 English words. Current frontier models support up to 2,000,000 token windows, meaning entire books or codebases can fit in a single context. The vocabulary (the full list of tokens the model knows) is fixed at training time. Each token is assigned a unique integer ID, and it is these IDs that flow into the next stage.
Converting tokens to embedding vectors
A token ID on its own is just an integer. The model needs a richer representation that captures what the token means, not just which entry in the vocabulary it corresponds to.
This representation is called an embedding vector. Its elements together encode the token’s meaning, but no single element maps neatly to one human-readable property. The way tokens are converted to embeddings is learned during training and is unique to each model. Tokens that appear in similar contexts end up with similar embeddings. These can be pictured as vectors in a high-dimensional space, typically 7,000–16,000 dimensions. Meaning is carried by directions through this space rather than by individual axes, and the model packs in far more distinct features than it has dimensions.
The geometry of the embedding space carries meaning: vectors for “king” and “queen” are closer to each other than either is to “bicycle.” This is what allows the model to generalise across related words and concepts. Individual dimensions are mostly not interpretable on their own; what a direction in the space means is fixed entirely by training. Researchers can “probe” embeddings to recover some of this structure, but the opacity doesn’t matter in practice: these embeddings are just part of the model’s internal machinery, not something we ever need to read directly.
In practice, this conversion is done with an embedding table : a large matrix with one row per token in the vocabulary. Each row is the vector of numbers that represents that token’s embedding. Looking up a token’s embedding is simply selecting the relevant row from this table.
At this stage, the embeddings contain no information about position. “the cat sat” and “sat the cat” would produce identical embeddings for each token, which is a problem for any task where word order matters. To solve this, positional information is added to each vector. See Encoding token position for more detail.
The result is a sequence of vectors, one per input token, each encoding the token’s meaning and position. This sequence is what enters the first layer of the model.
As the input passes through the model, each token carries a vector that starts as its embedding and is progressively refined at every layer, adding context and interpretation. This vector is called the residual stream .
Each token has its own residual stream. Across many layers, each token’s vector is gradually shifted from a representation of what that token is toward a prediction of what the next token should be. At inference time, only the last residual stream vector in the final layer is used to predict the next token. But all tokens are processed at every layer because later layers’ attention depends on what earlier layers produced for every position, not just the last one.
Before each sublayer, the residual stream is normalised: rescaled so its values sit within a predictable range, like adjusting the volume on different audio tracks so none drowns out the others. This keeps the values at a consistent scale as they pass through dozens of layers. Without it, values can grow or shrink exponentially during training, making the model impossible to optimise reliably.
Attention is the mechanism that lets each token query preceding tokens in the sequence and pull in relevant information. This is what allows the model to resolve references (in “the dog chased the cat because it was hungry,” “it” refers to the dog), track long-range dependencies, and build up a contextualised understanding of each token.
Each token only attends to tokens that came before it in the sequence. This is called causal (or masked) attention. It is wha

[truncated]
