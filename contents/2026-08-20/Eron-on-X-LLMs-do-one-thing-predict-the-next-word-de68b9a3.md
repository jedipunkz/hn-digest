---
source: "https://twitter.com/0xEronn/status/2089079333885198384"
hn_url: "https://news.ycombinator.com/item?id=49373160"
title: "Eron on X: \"LLMs do one thing: predict the next word"
article_title: "Eron on X: \"https://t.co/fQgIxVyHZJ\" / X"
image: "https://pbs.twimg.com/media/HP3Ej3SWMAAJB-F.jpg"
author: "bilsbie"
captured_at: "2026-08-20T12:26:42Z"
capture_tool: "hn-digest"
hn_id: 49373160
score: 2
comments: 0
posted_at: "2026-08-20T11:28:26Z"
tags:
  - hacker-news
  - translated
---

# Eron on X: "LLMs do one thing: predict the next word

- HN: [49373160](https://news.ycombinator.com/item?id=49373160)
- Source: [twitter.com](https://twitter.com/0xEronn/status/2089079333885198384)
- Score: 2
- Comments: 0
- Posted: 2026-08-20T11:28:26Z

## Translation

タイトル: Eron on X: 「LLM が行うことは 1 つあります。それは、次の単語を予測することです」
記事タイトル: エロン X: "https://t.co/fQgIxVyHZJ" / X
説明: LLM は次の単語を予測するという 1 つのことを行います。考えているようなものを作成する方法は次のとおりです

記事本文:
エロンが X に投稿: "https://t.co/fQgIxVyHZJ" / X ポスト
Eron @0xEronn LLM が行うことは 1 つあります。それは、次の単語を予測することです。考えているようなものを作成する方法は次のとおりです
ChatGPT、Claude、または Gemini にメッセージを送信するたびに、同じことが起こります。モデルは、あなたが書いたすべてを見て、次にどの単語が来るかという 1 つの決定を下します。その後、それが再び行われます。そしてまた。応答が完了するまで。
それでおしまい。それがすべてのトリックです。
そして、どういうわけか、理解できない量のテキスト上で何十億回も繰り返される、その非常に単純なタスクから、コードを記述し、量子物理学を説明し、詩を翻訳し、驚くほど人と話しているような会話を行うことができる何かが現れます。
実際に内部で何が起こっているかは次のとおりです。
すべてを大局的に捉える数字
GPT-3 は、人間が 1 日 24 時間、睡眠も休憩もせずに立ち止まらずにそのすべてを読もうとすると、2,600 年以上かかるという十分な量のテキストを学習しました。
それ以来、大規模なモデルはより多くのトレーニングを行っています。
モデルはあなたが学習した方法では言語を学習しなかったので、これは重要です。誰も座ってルールを書きませんでした。誰も文法や意味、文脈を説明しませんでした。モデルには、ほとんど理解できないほどの量のテキストが表示され、次に何が起こるかをより良く予測できるようになりなさいという 1 つの指示が与えられました。
言語、事実、推論、口調、性格について知っていることはすべて、頭の中に留めておくのが本当に難しい規模で適用された単一の最適化ターゲットから生まれました。
モデルの動作は完全に数値によって決まります。何千億も。これらはパラメータまたは重みと呼ばれ、モデルを別のモデルと区別する唯一のものです。
トレーニングの開始時に、すべてのパラメーターがランダムな値に設定されます。モデルは純粋なギベを生成します

上昇。それからトレーニングが始まります。
プロセスは次のように動作します。テキストの一部を取り出します。最後の単語を除くすべてをモデルにフィードします。モデルが次に何が起こるかを予測してみましょう。それを実際の最後の言葉と比較してください。モデルがどの程度間違っていたかを測定します。次に、これらの数千億のパラメータをすべて、モデルの誤りが少なくなる方向にわずかに調整します。
これを何兆もの例に対して実行します。
関係する数学は驚異的です。 1 秒あたり 10 億回の計算を実行でき、最大の言語モデルをトレーニングするために必要な計算を再現しようとすると、1 億年以上かかることになります。これは、膨大な数の操作を同時に実行する GPU と呼ばれる特殊なチップのおかげでのみ可能です。
トレーニングが完了すると、これらの数千億のパラメータは、言語がどのように機能するか、物事が何を意味するか、概念がどのように関係しているか、およびどのような状況でどのような応答が適切であるかについて、誰も完全に理解していない形式でエンコードします。
たった 1 つの予測がどのように会話になるのか
チャットボットにメッセージを送信すると、まさに次のようなことが起こります。
システムは会話の完全なコンテキストを取得し、AI の役割と目的の説明を追加して、そのすべてをモデルにフィードします。モデルはこのテキスト ブロック全体を調べて、最も可能性の高い次の単語を予測します。その単語がテキストに追加されます。その後、全体が再び実行されます。そしてまた。
各予測は確率的なものです。モデルは、次の可能性が最も高い単一の単語を毎回選択するわけではありません。ディストリビューションからサンプリングします。最も可能性の高い単語が最も頻繁に選択されますが、まれに、可能性の低い単語が選択されることもあります。同じ質問を 2 回すると異なる答えが得られ、応答が機械的ではなく自然に感じられるのはこのためです。

純粋な最尤予測では、堅苦しい繰り返しのテキストが生成されます。ディストリビューションからサンプリングすると、まるで人が書いたかのようなものが生成されます。
AI ツールで時々見られる温度設定は、この分布を制御します。温度が低いと、モデルがより決定的で焦点が絞られたものになります。温度が高いと、より多様性と創造性が高まり、行き過ぎると最終的には支離滅裂になります。
すべてを可能にしたアーキテクチャ
2017 年以前の言語モデルは、テキストを最初から最後まで単語ごとに処理していました。各単語は順番に処理されました。これにより、長いテキストに根本的な問題が発生しました。これは、文章の先頭からの情報が、最後に影響を与えるためにすべての中間ステップを通過する必要があるためです。
2017 年、Google のチームは「Attending Is All You Need」という論文を発表しました。彼らが導入したトランスフォーマーというアーキテクチャがすべてを変えました。
トランスフォーマーはテキストを順番に読み取りません。これらは入力全体を一度に取り込み、並行して処理します。すべての単語は、他のすべての単語に同時に直接関与できます。長い文章の始まりは、隣り合った単語と同じくらい直接的に終わりに影響を与えます。
重要なメカニズムに注意を喚起します。入力内のすべての単語は、数値の長いリストとして表されます。次に、アテンションにより、すべての単語が他のすべての単語を検査し、周囲のコンテキストに基づいてそれ自体の表現が更新されます。
「銀行」という言葉は、「川岸」というフレーズと「銀行口座」というフレーズでは意味が異なります。アテンションにより、モデルは周囲の単語を調べ、銀行の数値表現が実際にエンコードする内容を調整することで、これらの意味を区別できるようになります。
これは多くの層を通過します。各レイヤーはコンテキストに基づいて表現を調整します。最後の層までに、各単語の r が

表現は、次に何が起こるかを予測するのに役立つ可能性のあるあらゆる情報をエンコードします。これには、長さに関係なく、入力全体から引き出された情報も含まれます。
ネットワークは実際に何でできているのか
ニューラル ネットワークは、脳内でニューロンがどのように機能するかに大まかにインスピレーションを得た、接続された数学的演算のシステムです。
基本単位はニューロンです。 A neuron holds a number.その値は、前の層からそれに接続されているニューロンの値の加重和を取得し、バイアスと呼ばれる定数を追加し、その結果を出力を特定の範囲内に保つ関数に渡すことによって決定されます。
これらの接続の重みがパラメータです。トレーニング中に、ネットワークがそのタスクをより適切に実行できるように重みが調整されます。
画像認識の文脈でこれを具体的に見ることができます。手書きの数字を認識するようにトレーニングされたネットワークは、個々のピクセル値を処理することから始まります。最初の隠れ層はエッジの検出を学習する可能性があります。次のレイヤーでは、エッジを組み合わせて形状を作ります。最後のレイヤーでは、形状を組み合わせて数字の ID を作ります。誰もこれをプログラムしたわけではありません。それはトレーニングから生まれました。
言語モデルは、より複雑な構造を備えたより大規模なスケールで機能しますが、同じ基本原則が適用されます。モデルは、そのタスクに役立つ内部表現を学習します。それらの表現はデザインされたものではありません。 They emerge.
役に立つトレーニング
インターネット テキストでの事前トレーニングにより、モデルに言語を予測するよう教えられます。しかし、ランダムなインターネット テキストを予測することと、有用なアシスタントになることは同じではありません。
トレーニング前のデータのみでトレーニングされたモデルは、有害なコンテンツ、間違った情報、または技術的には流暢だがまったく役に立たない応答を生成するなど、与えられたパターンを継続するテキストを生成する可能性があります。
このチャットボットに対処するには、

したがって、人間のフィードバックからの強化学習と呼ばれる第 2 のトレーニング段階があります。人間の評価者はさまざまな応答を評価し、どれがより優れているかを示します。このフィードバックは、モデルのパラメーターをさらに調整して、人間が好む応答を生成する可能性を高めるために使用されます。
これが、モデルが特定のリクエストを拒否し、ただ流暢に話すだけでなく役立つように努め、一般にオートコンプリート システムよりもアシスタントのように動作する理由です。事前トレーニングにより、彼らは言語と世界についての知識を得ることができました。トレーニングの第 2 段階では、その知識をどのように応用するかを形作りました。
モデルが実際に学習したこと
これらのシステムがどのように機能するかについて、本当に奇妙な点があります。
手書きの数字を認識するようにトレーニングされたニューラル ネットワークが視覚化されたとき、各ニューロンがどのようなパターンに反応するように学習したかを見ると、結果は必ずしも期待どおりになるとは限りません。ネットワークは、エッジを明確に学習してから形成するのではなく、明確な解釈のないパターンを学習することがよくあります。それらは機能しますが、期待どおりには機能しません。
同じことが、はるかに大きな規模の言語モデルにも当てはまります。研究者は、モデルが推論のようなこと、類推のようなこと、コンテキストの理解のようなことを行うことを観察できます。しかし、これらの動作が基礎となる数学からどのように現れるのかは、これらのシステムを構築した人々さえも正確に理解していません。
モデルは推論するようにプログラムされていません。彼らはテキストを予測するように訓練されました。その推論は、それが本当であるとすれば、その大規模なトレーニングから明らかになりました。
これが、これらのシステムを優れたものにすると同時に、予測が困難な点で信頼性を低くするものでもあります。複雑な概念を明確に説明できるモデルは、何か完全に間違っていることを自信を持って述べている可能性もあります。それは何を学んだのか

流暢で役立つテキストのようです。その文章の内容が正確であるかどうかは別の問題です。
誰も完全に理解していないこと
大規模な言語モデルをランダムなラベルでトレーニングする場合、つまりトレーニング前にすべての正解をスクランブルした場合でも、モデルはトレーニング データに対して完全な精度を達成できます。それはただ暗記するだけです。それは一般化されません。
正しくラベル付けされたデータでトレーニングすると、モデルははるかに速く学習し、これまでに見たことのない例に一般化します。データ内の構造により、パターンを見つけやすくなります。
これは重要なことを教えてくれます。これらのモデルは単に暗記するだけではありません。彼らは構造を見つけつつあります。その構造が何に相当するのか、それが理解のようなものなのか、あるいは単にそれに似ているものなのかは、研究者たちが積極的に取り組んでいる問題ですが、まだ答えは出ていません。
明らかなことは、十分なデータに基づいて十分な規模で次の単語を予測すると、誰も明示的にプログラムしていないことを実行できるシステムが生成されるということです。コードを書きます。言語を翻訳します。書類を要約します。トレーニング データのどこにも現れないトピックに関する質問に答えます。
なぜこれが終わりではなく始まりなのか
2017 年のトランスフォーマー アーキテクチャにより、大規模なテキストの並列処理が可能になりました。最大のモデルをトレーニングするために必要な計算の規模は、手動で実行するには何百万年もかかる操作で測定されます。次の単語の予測だけから有用な機能が出現したことは、これらのシステムを構築した研究者さえも驚かせました。
これはどれも、機能する前には明らかではありませんでした。テキストを予測することで、ソフトウェアを作成したり、科学を説明したり、複数回にわたる一貫した会話を行うことができるものを生み出すことができるという考えは、誰も自信を持って予測したものではありませんでした。
言語に with が含まれているため機能しました

そこには世界に関する膨大な量の構造が含まれています。物理実験の説明で次にどの単語が来るかを予測するには、物理​​についてある程度理解する必要があります。コード内の次の単語を予測するには、コードが何をしているのかを理解する必要があります。モデルには物理学やプログラミングを直接教えられたことはありません。テキストをより正確に予測するために学習するために必要なものはすべて学習しました。
仕事がひとつ。次の単語を予測します。数兆のサンプルに対して数千億のパラメータを使用してそれを実行します。
考えていることが向こう側に落ちていくような気がします。その理由は誰にも完全にはわかりません。
/ これが役に立った場合は、フォローしてください。次のものが最初にここにドロップされます。
Carl Reed @CarlReed_on_X 8 月 18 日 記事をざっと読んだ...数時間以内に戻って最後まで読みます。時間をかける価値がある、いや、時間をかける価値があるとすでにわかっています。 1 834
何が起こっているかを確認して会話に参加してください
電話で続行 Apple で続行 Google で続行、またはユーザー名または電子メールでログイン 関係者

## Original Extract

LLMs do one thing: predict the next word. Here's how that creates something that feels like thinking

Eron on X: "https://t.co/fQgIxVyHZJ" / X Post
Eron @0xEronn LLMs do one thing: predict the next word. Here's how that creates something that feels like thinking
Every time you send a message to ChatGPT, Claude or Gemini the same thing happens. The model looks at everything you wrote and makes one decision: what word probably comes next. Then it does it again. And again. Until the response is complete.
That's it. That's the whole trick.
And somehow from that one absurdly simple task repeated billions of times on an incomprehensible amount of text something emerges that can write code, explain quantum physics, translate poetry and hold a conversation that feels remarkably like talking to a person.
Here's what's actually happening inside.
The number that puts everything in perspective
GPT-3 trained on enough text that if a human tried to read all of it without stopping - no sleep, no breaks, twenty four hours a day - it would take over 2,600 years.
Larger models since then trained on significantly more.
This matters because the model didn't learn language the way you learned it. Nobody sat down and wrote rules. Nobody explained grammar or meaning or context. The model was shown an almost incomprehensible volume of text and given one instruction: get better at predicting what comes next.
Everything it knows about language, facts, reasoning, tone and personality emerged from that single optimization target applied at a scale that is genuinely difficult to hold in your head.
The model's behavior is determined entirely by numbers. Hundreds of billions of them. These are called parameters or weights and they're the only thing that makes one model different from another.
At the start of training every parameter is set to random values. The model produces pure gibberish. Then training begins.
The process works like this. Take a piece of text. Feed all of it except the last word into the model. Look at what the model predicted would come next. Compare that to the actual last word. Measure how wrong the model was. Then adjust every one of those hundreds of billions of parameters slightly in the direction that would have made the model less wrong.
Do this for trillions of examples.
The math involved is staggering. If you could perform one billion calculations per second and you tried to replicate the computation required to train the largest language models it would take you over 100 million years. This is only possible because of specialized chips called GPUs that run enormous numbers of operations simultaneously.
After training is complete those hundreds of billions of parameters encode - in a form that nobody fully understands - something about how language works, what things mean, how concepts relate and what kinds of responses are appropriate in what kinds of situations.
How a single prediction becomes a conversation
When you send a message to a chatbot here is exactly what happens.
The system takes the full context of your conversation, adds a description of the AI's role and purpose, and feeds all of it into the model. The model looks at this entire block of text and predicts the most likely next word. That word gets added to the text. Then the whole thing runs again. And again.
Each prediction is probabilistic. The model doesn't pick the single most likely next word every time. It samples from a distribution. The most likely words get picked most often but occasionally a less likely word gets selected. This is why asking the same question twice gives you different answers and why responses feel natural rather than mechanical. Pure maximum-likelihood prediction produces stilted repetitive text. Sampling from the distribution produces something that reads like a person wrote it.
The temperature setting you sometimes see in AI tools controls this distribution. Low temperature makes the model more deterministic and focused. High temperature makes it more varied and creative and eventually incoherent if pushed too far.
The architecture that made everything possible
Before 2017 language models processed text word by word from start to finish. Each word was handled sequentially. This created a fundamental problem with long texts because information from the beginning of a passage had to travel through every intermediate step to influence the end.
In 2017 a team at Google published a paper called Attention Is All You Need. The architecture they introduced - the transformer - changed everything.
Transformers don't read text sequentially. They take in the entire input at once and process it in parallel. Every word can directly attend to every other word simultaneously. The beginning of a long passage influences the end as directly as words that appear right next to each other.
The key mechanism is called attention. Every word in the input gets represented as a long list of numbers. Then attention lets every word examine every other word and update its own representation based on the context around it.
The word bank means something different in the phrase river bank than in the phrase bank account. Attention allows the model to distinguish between these meanings by examining the surrounding words and adjusting what the number representation of bank actually encodes.
This runs through many layers. Each layer refines the representations based on context. By the final layer the goal is that each word's representation encodes whatever information might be useful for predicting what comes next - including information pulled from across the entire input, however long it is.
What the network is actually made of
A neural network is a system of connected mathematical operations loosely inspired by how neurons work in the brain.
The basic unit is a neuron. A neuron holds a number. Its value is determined by taking a weighted sum of the values of the neurons connected to it from the previous layer, adding a constant called a bias and passing the result through a function that keeps the output within a certain range.
The weights on those connections are the parameters. During training the weights get adjusted to make the network better at its task.
In the context of image recognition you can see this concretely. A network trained to recognize handwritten digits starts by processing individual pixel values. The first hidden layer might learn to detect edges. The next layer combines edges into shapes. The final layer combines shapes into digit identities. Nobody programmed any of this. It emerged from training.
Language models work at a much larger scale with much more complex structure but the same basic principle applies. The model learns internal representations that are useful for its task. Those representations aren't designed. They emerge.
The training that makes it useful
Pre-training on internet text teaches the model to predict language. But predicting random internet text is not the same as being a useful assistant.
A model trained only on pre-training data might produce text that continues whatever pattern it was given - including producing harmful content, wrong information or responses that are technically fluent but completely unhelpful.
To address this chatbots undergo a second training stage called reinforcement learning from human feedback. Human raters evaluate different responses and indicate which ones are better. This feedback gets used to further adjust the model's parameters to make it more likely to produce responses that humans prefer.
This is why models refuse certain requests, try to be helpful rather than just fluent, and generally behave more like assistants than autocomplete systems. The pre-training gave them knowledge of language and the world. The second training stage shaped how they apply that knowledge.
What the model actually learned
Here is something that is genuinely strange about how these systems work.
When a neural network trained to recognize handwritten digits is visualized - when you look at what patterns each neuron has learned to respond to - the result is not always what you'd expect. Instead of clearly learning edges and then shapes the network often learns patterns that don't have obvious interpretations. They work but not in the way you might have hoped.
The same is true for language models at a much larger scale. Researchers can observe that the models do something that looks like reasoning, something that looks like analogy, something that looks like understanding context. But exactly how those behaviors emerge from the underlying mathematics is not well understood even by the people who built these systems.
The models are not programmed to reason. They were trained to predict text. The reasoning - if that's what it is - emerged from that training at scale.
This is both what makes these systems impressive and what makes them unreliable in ways that are hard to predict. The model that can explain a complex concept clearly might also confidently state something completely false. It learned what fluent helpful text looks like. Whether the content of that text is accurate is a different question.
The thing nobody fully understands
If you train a large language model on random labels - if you scramble all the correct answers before training - the model can still achieve perfect accuracy on the training data. It just memorizes. It doesn't generalize.
If you train on correctly labeled data the model learns much faster and generalizes to examples it has never seen. The structure in the data makes the patterns easier to find.
This tells you something important. These models are not simply memorizing. They are finding structure. What that structure corresponds to - whether it's something like understanding or something that merely resembles it - is a question that researchers are actively working on and haven't answered.
What is clear is that next-word prediction at sufficient scale on sufficient data produces systems that can do things nobody explicitly programmed them to do. Write code. Translate languages. Summarize documents. Answer questions about topics that appear nowhere in their training data.
Why this is the beginning not the end
The transformer architecture from 2017 enabled parallel processing of text at scale. The scale of computation required to train the largest models is measured in operations that would take millions of years to perform by hand. The emergence of useful capabilities from next-word prediction alone surprised even the researchers who built these systems.
None of this was obvious before it worked. The idea that predicting text could produce something capable of writing software or explaining science or engaging in coherent multi-turn conversation was not something anyone predicted with confidence.
It worked because language contains within it an enormous amount of structure about the world. To predict what word comes next in a description of a physics experiment you need to understand something about physics. To predict the next word in a piece of code you need to understand something about what the code is doing. The model was never taught physics or programming directly. It learned whatever it needed to learn to get better at predicting text.
One job. Predict the next word. Do it with hundreds of billions of parameters on trillions of examples.
Something that feels like thinking falls out the other side. Nobody is entirely sure why.
/ If this was useful - follow, the next one drops here first.
Carl Reed @CarlReed_on_X Aug 18 Skimmed the article...will come back to it and read it through in a few hours. I can already tell it will be worth my time, WELL worth my time. 1 834
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
