---
source: "https://arpitbhayani.me/blogs/how-llms-work/"
hn_url: "https://news.ycombinator.com/item?id=48343242"
title: "How LLMs Work"
article_title: "How LLMs Really Work"
author: "dharaniES"
captured_at: "2026-06-01T00:35:11Z"
capture_tool: "hn-digest"
hn_id: 48343242
score: 4
comments: 0
posted_at: "2026-05-31T05:18:31Z"
tags:
  - hacker-news
  - translated
---

# How LLMs Work

- HN: [48343242](https://news.ycombinator.com/item?id=48343242)
- Source: [arpitbhayani.me](https://arpitbhayani.me/blogs/how-llms-work/)
- Score: 4
- Comments: 0
- Posted: 2026-05-31T05:18:31Z

## Translation

タイトル: LLM の仕組み
記事のタイトル: LLM が実際にどのように機能するか
説明: ChatGPT、Gemini、または Claude を使用したことがある場合は、これらのシステムが何を行うかについてすでに直感を形成しているでしょう。何かを入力すると、一貫性があり、知識が豊富で、時には不気味な人間味を感じるテキストが返されます。しかし、その下にある機械は、ほとんどの人よりも単純であると同時に奇妙でもあります。
[切り捨てられた]

記事本文:
LLM の実際の仕組み
アルピット・バヤニ
ブログ
最初の原則から
--> 紙棚 本棚
セッション
--> コースでは LLM が実際にどのように機能するかを説明します
エンジニアリング、データベース、システム。常に構築しています。
ChatGPT 、 Gemini 、または Claude を使用したことがある場合は、これらのシステムが何を行うかについてすでに直感を形成しているでしょう。何かを入力すると、一貫性があり、知識が豊富で、時には不気味な人間味を感じるテキストが返されます。しかし、その下にある機械は、ほとんどの人が予想しているよりも単純であると同時に奇妙でもあります。
この記事では、その仕組みを解き明かし、言語モデルが機械レベルで何をしているのか、つまり言語モデルがそのような出力を生成する理由、同一の入力が異なる実行で異なる出力を生成する理由、そして「創造性ダイヤル」を超えて実際に「温度」が何を意味するのかを説明します。
大規模言語モデル (LLM) は、最も基本的なレベルで、一連のトークンを入力として受け取り、次のトークンがどうなるかについての語彙全体にわたる確率分布を出力する関数です。これがコア操作の完全な説明です。他のすべて - 明らかな推論、会話能力、コード生成 - は、膨大な量のトレーニング データにわたって、この 1 つのことを巨大な規模で実行することによって生まれます。
具体的には、モデルに「素早い茶色のキツネ」のトークンを与えると想像してください。このモデルは「ジャンプ」という言葉を生成しません。これは確率の表を作成します。「ジャンプ」の確率は 42%、「座る」の確率は 12%、「跳躍」の確率は 8%、100,000 語の語彙内の他のすべてのトークンは、残りの確率質量のゼロ以外のスライスを取得します。次に、モデルはその分布からサンプリングして次のトークンを選択します。そのトークンがシーケンスに追加され、停止条件が発生するまでプロセス全体が繰り返されます。

に達します。
これは 自己回帰生成 と呼ばれます。生成された各トークンは、次の予測の入力の一部になります。モデルは常に同じ質問をします。「これまで見てきたことを考慮すると、次に来る可能性が最も高いトークンは何ですか?」
モデルは、大量のテキスト コーパス (基本的にインターネット、書籍、コード、学術論文で書かれたものの大部分) をトレーニングすることによって、これらの確率分布を生成する方法を学習します。トレーニング中に、モデルは一連のトークンを確認し、次のトークンを予測しようとします。
それが間違っている場合、エラー信号はネットワークを介して逆方向に流れ（逆伝播を介して）、何十億もの内部パラメーター (モデルの「重み」) を正しい予測の可能性を高める方向にわずかに微調整します。
こうした何兆回もの更新を経て、モデルの重みは注目すべきもの、つまり言語がどのように機能するかを示す圧縮された統計モデルをエンコードします。 「エッフェル塔は～にあります」の後には「パリ」が続くことが非常に多く、Python の関数定義は「def」で始まり、「To be or not to」で始まる文はほぼ確実に「be」で続くことがわかります。
重要なのは、モデルには個々のトレーニング例の記憶がないことです。内部化された統計パターンがあります。これが、新しい入力に一般化できる理由です。保存された文を取得するのではなく、学習された分布からサンプリングします。
ロジッツ、ソフトマックス、そして確率が重要な理由
モデルはこれらのクリーンな確率を生成する前に、ロジットと呼ばれる生のスコア (語彙内のトークンごとに 1 つの実数) を生成します。これらのロジットは、ニューラル ネットワークの最後の線形層の生の出力です。
ロジットを確率分布に変換するために、モデルはソフトマックス関数を適用します。
P (トークン i ) = e ロジ

t i ∑ j e logit j P(\text{token}_i) = \frac{e^{\text{logit}_i}}{\sum_j e^{\text{logit}_j}} P ( token i ) = ∑ j e logit j e logit i
Softmax は 2 つのことを行います。まず、各ロジットをべき乗して差を増幅します。つまり、2 倍の大きさのロジットの可能性が指数関数的に高くなります。次に、すべての確率の合計が 1 になるようにすべてを正規化します。その結果、語彙全体にわたる有効な確率分布が得られます。
これを実際に確認するには、モデルが「素早い茶色のキツネ」の次の単語を予測していると想像してください。 4 つの単語からなる小さな語彙の生のロジットを生成します。
これは、サンプリング前にモデルが実際に渡す番号です。温度、top-k、核サンプリングのドラマ全体が、トークンが抽出される前のこの分布の操作においてここで起こります。
温度は、プロンプトで最も誤解されているパラメータです。これは一般に「創造性」または「ランダム性」と表現されますが、技術的には正しいのですが、それがどのように機能するのかが正確にはわかりにくくなっています。それを正確に理解することで、意図的に使用できるようになります。
温度は、ソフトマックスが適用される前のロジットを除算するスカラーです。
P ( トークン i ) = e logit i / T ∑ j e logit j / T P(\text{token}_i) = \frac{e^{\text{logit}_i / T}}{\sum_j e^{\text{logit}_j / T}} P (トークン i ) = ∑ j e logit j / T e logit i / T
T = 1.0 T = 1.0 T = 1.0 の場合、何も変わりません。確率は、前の例で生のソフトマックスが生成したものとまったく同じです。
T T T を調整することで、分布を「尖らせる」か「平坦化」することができます。
T < 1.0 T < 1.0 T < 1.0 (例: T = 0.5 T = 0.5 T = 0.5 ) の場合、分数で割るとロジットが拡大されます。温度が下がると、すでに大きな上位トークン間の差がさらに大きくなります。モデルのベコ

これはほぼ決定的であり、最も可能性の高い単一のトークンを圧倒的に選択します。
T > 1.0 T > 1.0 T > 1.0 (例: T = 2.0 T = 2.0 T = 2.0 ) の場合、大きな数で割るとロジットが平坦になります。火力を上げると、確率質量がより均一に広がります。これまでありそうになかったトークンがもっともらしい候補になり、モデルはより驚くべき継続をサンプリングすることになります。
これには見落とされがちな実際的な意味があります。つまり、温度によってモデルが認識していることやモデルの推論方法が変化することはありません。確率分布のどの領域からサンプリングするかが変わります。低温では、モデルの最も信頼性の高い予測を利用することになります。高温では、分布の末尾を探索することになります。これには、有効ではあるが異常な継続や、支離滅裂な継続が含まれます。
T = 0.0 T = 0.0 T = 0.0 ～ 0.3 0.3 0.3 : ほぼ決定的な出力、コード生成、事実に基づく Q&A、構造化データの抽出に適しています
T = 0.7 T = 0.7 T = 0.7 ～ 1.0 1.0 1.0 : バランスが取れており、チャット、要約、汎用用途に適しています
T = 1.2 T = 1.2 T = 1.2 ～ 2.0 2.0 2.0 : 多様性が高く、ブレーンストーミング、クリエイティブな執筆、珍しい表現の探索に適しています。ただし、ハイエンドでは出力の信頼性がますます低くなります。
出力が本質的に確率的である理由
言語モデルに「2 + 2 2 + 2 2 + 2 とは何ですか?」と尋ねると、温度に関係なく、毎回 " 4 4 4 " が返されます。これは、確率の質量がそのトークンに非常に集中しているため、高温のサンプリングでも他のものはほとんど選択されないためです。ただし、複数の継続が考えられるプロンプトの場合、モデルの出力は確率分布から抽出されます。同じプロンプトを 100 回実行すると、100 個のわずかに異なる、場合によっては実質的に異なる出力が得られます。
これはバグではありません。それは

モデルがどのようにトレーニングされたかに直接影響されます。トレーニング データには膨大なバリエーションが含まれており、さまざまな人々が同じアイデアを何千もの異なる方法で表現しています。モデルはこの変動を学習しました。電子メールを書いたり文書を要約したりするように依頼する場合、さまざまな表現が合理的であり、モデルはそれを反映しています。
出力の確率的な性質は、経験豊富なエンジニアが苦労して学ぶいくつかの実際的な結果をもたらします。
同じプロンプトを使用した場合でも、モデルが常に同じ構造を出力に生成するとは限りません。厳密な出力解析では変動に対処する必要があります。
モデルは、同一の入力であっても、別々の呼び出し間で矛盾する可能性があります。一貫性が必要な場合は、温度 0 0 0 を使用するか、検証ロジックを実装します。
「間違った答えが得られた」と「確実に間違った答えが得られた」は、まったく異なる失敗モードです。プロンプトが機能すると結論付ける前に、必ず複数の実行にわたってテストしてください。
初心者がつまずく点の 1 つは、モデルは認知的な意味で何を言うかを「決定」しないことです。内なる独白や、返答を書く前に概要を説明する計画ステップはありません。各トークンは一度に 1 つずつ左から右に生成され、コミット後に以前のトークンを修正することはできません。
これが、最終的な答えを与える前にモデルに段階的に推論を求める「思考連鎖プロンプト」が実際に複雑なタスクの精度を向上させる理由です。中間推論トークンを生成することにより、モデルはその推論に基づいて後のトークンを条件付けします。スクラッチ スペースは実際の機能的なものです。出力に「ステップバイステップで考えさせます」と書き込むと、正確性が向上する形で後続のトークンの分布が実際に変更されます。それは劇場版ではありません。
また、モデルが「幻覚」を起こす理由も説明されています。

te ” - 自信があるように聞こえるが虚偽のテキストを生成します。文脈上特定の詳細 (作成者名、統計、URL) を期待するプロンプトが与えられると、モデルは学習した分布からもっともらしい響きの継続をサンプリングします。その分布は実際のテキストに基づいて構築されていますが、事実の正確さのためにインデックス付けされていません。もっともらしいトークンは真のトークンと同じではありません。
「モデルは知っている」が実際に何を意味するか
エンジニアが言語モデルが何かを「知っている」と言うとき、トレーニング コーパスには、その情報がコンテキスト内で出現する多くの例が含まれており、モデルの重みがそれを表現する継続に向けた強い事前確率をエンコードすることを意味します。モデルには事実のデータベースがありません。これには、数千億のトークンにわたる共起統計の圧縮された非可逆エンコードが含まれています。
これは実際には重要です。モデルは、トレーニングで何度も見たことについて自信があり、一貫性があります。滅多に登場しないものや表現に一貫性がないものについては信頼性が低いです。トークン予測機構は「私はこれを学んだ」と「私はもっともらしいものにパターンマッチングしている」を区別しないため、トレーニング データでは過小評価されているドメインの詳細を自信を持って補います。
これを理解すると、プロンプトを適切に設計するのに役立ちます。一般的な知識に基づいたタスクの場合、モデルは強力な加速剤となります。特に特定の数字、引用、最近の出来事など、正確な事実の想起が必要なタスクでは、モデルを検証が必要な出発点として扱う必要があります。
言語モデルは、大規模なコーパスから保持されたトークンを予測する際の誤差を最小限に抑えることによってトレーニングされた次のトークン予測マシンです。各ステップでの語彙にわたる確率分布と温度制御を出力します。

サンプリング前にその分布がどの程度鋭いピークに達しているかがわかります。
モデルは人間の言語の自然なバリエーションを学習しているため、出力は確率的になります。モデルを検索エンジンや知識ベースとして扱うのではなく、これを理解することが、運用システムで LLM を効果的に使用するための基礎となります。
これが役に立ち、興味深いと思われた場合は、
RSS フィードを購読すると、新しいフィードを公開したときに通知が届きます。
Razorpay のプリンシパル エンジニア II、GCP Memorystore および Dataproc の元スタッフエンジニア、DiceDB の作成者、元 Amazon Fast Data、元 Engg ディレクター。 SREと
Unacademy のデータ エンジニアリング。私は自分自身を通じてエンジニアリングへの好奇心を刺激します
綿密なエンジニアリングビデオ
YouTube
そして私のコース
145,000 人のエンジニアが読んだ Arpit のニュースレター
現実世界のシステム設計、分散システム、または
非常に賢いアルゴリズムを深く掘り下げてみましょう。
LinkedIn で購読する Substack で購読する
このウェブサイトに掲載されているコースは、
リログディープテック株式会社株式会社
203, Sagar Apartment, Camp Road, Mangiral Plot, アムラヴァティ, マハーラーシュトラ州, 444602
GSTIN: 27AALCR5165R1ZF
YouTube (200k) Twitter (100k) LinkedIn (250k) GitHub (6k)
© アルピット・バヤニ、2025

## Original Extract

If you have used ChatGPT, Gemini, or Claude, you have already formed an intuition about what these systems do. You type something in, and text comes back that feels coherent, knowledgeable, and sometimes eerily human. But the machinery underneath is simultaneously simpler and stranger than most peop
[truncated]

How LLMs Really Work
Arpit Bhayani
Blogs
From the First Principles
--> Papershelf Bookshelf
Sessions
--> Courses Talks How LLMs Really Work
engineering, databases, and systems. always building.
If you have used ChatGPT , Gemini , or Claude , you have already formed an intuition about what these systems do. You type something in, and text comes back that feels coherent, knowledgeable, and sometimes eerily human. But the machinery underneath is simultaneously simpler and stranger than most people expect.
This article tears open that machinery and explains what a language model is doing at a mechanical level - why it produces the outputs it does, why identical inputs produce different outputs on different runs, and what “temperature” actually means beyond “a creativity dial.”
A large language model (LLM) is, at its most fundamental level, a function that takes a sequence of tokens as input and outputs a probability distribution over its entire vocabulary for what the next token should be. That is the complete description of the core operation. Everything else - the apparent reasoning, the conversational ability, the code generation - emerges from doing this one thing at enormous scale, across an enormous amount of training data.
Concretely, imagine you feed the model the tokens for “The quick brown fox”. The model does not produce the word “jumps”. It produces a table of probabilities: “jumps” might have a 42% chance, “sat” a 12% chance, “leaped” an 8% chance, and every other token in a 100,000-word vocabulary gets some non-zero slice of the remaining probability mass. The model then samples from that distribution to pick the next token. That token gets appended to the sequence, and the whole process repeats until a stop condition is reached.
This is called autoregressive generation . Each token generated becomes part of the input for the next prediction. The model is always asking the same question: “given everything I have seen so far, what token is most likely to come next?”
The model learns to produce these probability distributions by training on a massive corpus of text - essentially a large fraction of the written internet, books, code, and academic papers. During training, the model sees a sequence of tokens and tries to predict the next one.
When it is wrong, the error signal flows backward through the network (via backpropagation ), nudging billions of internal parameters - the model’s “ weights ” - very slightly in the direction that would have made the correct prediction more probable.
After trillions of these updates, the model’s weights encode something remarkable: a compressed statistical model of how language works. It learns that “The Eiffel Tower is located in” is very frequently followed by “Paris,” that Python function definitions start with “def,” and that a sentence starting “To be or not to” almost certainly continues with “be.”
Crucially, the model does not have a memory of individual training examples. It has internalized statistical patterns. This is why it can generalise to novel inputs - it is not retrieving stored sentences, it is sampling from learned distributions.
Logits, Softmax, and Why Probabilities Matter
Before the model produces those clean probabilities, it produces raw scores called logits - one real number per token in the vocabulary. These logits are the raw output of the final linear layer in the neural network.
To convert logits to a probability distribution, the model applies the softmax function :
P ( token i ) = e logit i ∑ j e logit j P(\text{token}_i) = \frac{e^{\text{logit}_i}}{\sum_j e^{\text{logit}_j}} P ( token i ​ ) = ∑ j ​ e logit j ​ e logit i ​ ​
Softmax does two things. First, it exponentiates each logit, which amplifies differences: a logit that is twice as large becomes exponentially more probable. Second, it normalizes everything so that all probabilities sum to 1. The result is a valid probability distribution over the entire vocabulary.
To see this in action, imagine the model is predicting the next word after “The quick brown fox”. It generates raw logits for a tiny vocabulary of four words:
This is the number the model actually hands you before sampling. The entire drama of temperature, top-k, and nucleus sampling happens here, in the manipulation of this distribution before a token is drawn from it.
Temperature is the most misunderstood parameter in prompting. It is commonly described as “creativity” or “randomness,” which is technically correct but obscures exactly how it works. Understanding it precisely lets you use it deliberately.
Temperature is a scalar that divides the logits before the softmax is applied:
P ( token i ) = e logit i / T ∑ j e logit j / T P(\text{token}_i) = \frac{e^{\text{logit}_i / T}}{\sum_j e^{\text{logit}_j / T}} P ( token i ​ ) = ∑ j ​ e logit j ​ / T e logit i ​ / T ​
When T = 1.0 T = 1.0 T = 1.0 , nothing changes. The probabilities are exactly what the raw softmax produced in our previous example.
By adjusting T T T , we can either “sharpen” or “flatten” the distribution:
When T < 1.0 T < 1.0 T < 1.0 (e.g., T = 0.5 T = 0.5 T = 0.5 ), dividing by a fraction magnifies the logits. By cooling the temperature, the already-large difference between the top tokens becomes enormous. The model becomes nearly deterministic, overwhelmingly picking the single most likely token.
When T > 1.0 T > 1.0 T > 1.0 (e.g., T = 2.0 T = 2.0 T = 2.0 ), dividing by a large number flattens the logits. By turning up the heat, the probability mass is spread more evenly. Previously unlikely tokens become plausible candidates, meaning the model will sample more surprising continuations.
This has a practical implication that is easy to miss: temperature does not change what the model knows or how it reasons. It changes which region of the probability distribution you sample from. At low temperature you are exploiting the model’s most confident predictions. At high temperature you are exploring the tail of the distribution, which contains valid but unusual continuations - as well as incoherent ones.
T = 0.0 T = 0.0 T = 0.0 to 0.3 0.3 0.3 : near-deterministic output, good for code generation, factual Q&A, structured data extraction
T = 0.7 T = 0.7 T = 0.7 to 1.0 1.0 1.0 : balanced, good for chat, summarisation, general-purpose use
T = 1.2 T = 1.2 T = 1.2 to 2.0 2.0 2.0 : high diversity, good for brainstorming, creative writing, exploring unusual phrasings - but outputs become increasingly unreliable at the high end
Why Outputs are Inherently Probabilistic
If you ask a language model “What is 2 + 2 2 + 2 2 + 2 ?” you will get " 4 4 4 " back every time regardless of temperature, because the probability mass is so concentrated on that token that even high-temperature sampling almost never picks anything else. But for any prompt where multiple continuations are plausible, the model’s outputs are drawn from a probability distribution. Run the same prompt a hundred times and you will get a hundred slightly different outputs, sometimes substantively different.
This is not a bug. It is a direct consequence of how the model was trained. The training data contains enormous variation: different people express the same idea in thousands of different ways. The model has learned this variation. When you ask it to write an email or summarise a document, many different phrasings are reasonable, and the model reflects that.
The probabilistic nature of outputs has several practical consequences that experienced engineers learn the hard way:
You cannot assume the model will always produce the same structure in its output, even with the same prompt. Strict output parsing must handle variation.
The model can contradict itself across separate calls even with identical input. For anything requiring consistency, either use temperature 0 0 0 or implement validation logic.
“It gave me a wrong answer” and “it gives wrong answers reliably” are very different failure modes. Always test across multiple runs before concluding a prompt works.
One thing that trips up newcomers: the model does not “decide” what to say in any cognitive sense. There is no inner monologue, no planning step where it outlines a response before writing it. Each token is generated one at a time, left to right, with no ability to revise earlier tokens once they are committed.
This is why “ chain-of-thought prompting” - asking the model to reason step by step before giving a final answer - actually improves accuracy on complex tasks. By generating intermediate reasoning tokens, the model conditions later tokens on that reasoning. The scratch space is real and functional: writing “let me think step by step” into the output genuinely changes the distribution over subsequent tokens in a way that improves correctness. It is not theatrical.
It also explains why the model can “ hallucinate ” - generate confident-sounding but false text. Given a prompt that contextually expects a specific detail (an author name, a statistic, a URL), the model samples a plausible-sounding continuation from its learned distribution. That distribution was built on real text, but it was not indexed for factual accuracy. A plausible token is not the same as a true one.
What “the model knows” Actually Means
When engineers say a language model “knows” something, they mean the training corpus contained many examples where that piece of information appeared in context, causing the model’s weights to encode a strong prior toward continuations that express it. The model does not have a database of facts. It has a compressed, lossy encoding of co-occurrence statistics across hundreds of billions of tokens.
This matters in practice. The model is confident and coherent about things it has seen many times in training. It is unreliable about things that appeared rarely or were expressed inconsistently. It will confidently make up details in domains that are underrepresented in its training data, because the token-prediction machinery does not distinguish between “I learned this” and “I am pattern-matching to something plausible.”
Understanding this helps you design prompts appropriately. For tasks grounded in common knowledge, the model is a powerful accelerant. For tasks requiring precise factual recall, especially of specific numbers, citations, or recent events, the model needs to be treated as a starting point that requires verification.
A language model is a next-token prediction machine trained by minimizing the error on predicting held-out tokens from a large corpus. It outputs a probability distribution over its vocabulary at each step, and temperature controls how sharply peaked that distribution is before sampling.
Outputs are probabilistic because the model has learned the natural variation in human language. Understanding this - rather than treating the model as a search engine or a knowledge base - is the foundation for using LLMs effectively in production systems.
If you find this helpful and interesting,
subscribe to my RSS feed and get notified the moment I publish a new one.
Principal Engineer II at Razorpay, Ex-staff engg at GCP Memorystore & Dataproc, Creator of DiceDB , ex-Amazon Fast Data, ex-Director of Engg. SRE and
Data Engineering at Unacademy. I spark engineering curiosity through my
no-fluff engineering videos on
YouTube
and my courses
Arpit's Newsletter read by 145,000 engineers
Weekly essays on real-world system design, distributed systems, or a
deep dive into some super-clever algorithm.
Subscribe on LinkedIn Subscribe on Substack
The courses listed on this website are offered by
Relog Deeptech Pvt. Ltd.
203, Sagar Apartment, Camp Road, Mangilal Plot, Amravati, Maharashtra, 444602
GSTIN: 27AALCR5165R1ZF
YouTube (200k) Twitter (100k) LinkedIn (250k) GitHub (6k)
© Arpit Bhayani, 2025
