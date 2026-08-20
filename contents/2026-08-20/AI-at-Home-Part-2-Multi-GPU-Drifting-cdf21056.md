---
source: "https://jdagostino.github.io/ai-pt2-multi-gpu-drifting/index.html"
hn_url: "https://news.ycombinator.com/item?id=49377155"
title: "AI at Home Part 2: Multi-GPU Drifting"
article_title: "AI At Home Part 2: Multi GPU Drifting"
image: "https://jdagostino.github.io/ai-pt2-multi-gpu-drifting/drifting_cover.jpeg"
author: "timmmmmmay"
captured_at: "2026-08-20T17:20:34Z"
capture_tool: "hn-digest"
hn_id: 49377155
score: 2
comments: 0
posted_at: "2026-08-20T16:53:55Z"
tags:
  - hacker-news
  - translated
---

# AI at Home Part 2: Multi-GPU Drifting

- HN: [49377155](https://news.ycombinator.com/item?id=49377155)
- Source: [jdagostino.github.io](https://jdagostino.github.io/ai-pt2-multi-gpu-drifting/index.html)
- Score: 2
- Comments: 0
- Posted: 2026-08-20T16:53:55Z

## Translation

タイトル: AI at Home パート 2: マルチ GPU ドリフト
記事タイトル: AI At Home パート 2: マルチ GPU ドリフト
説明: 電子廃棄物 GPU のボックスで AI モデルを高速に実行する

記事本文:
前の章では、AI 言語モデルを実行するために電子廃棄物グレードの GPU を使用してとんでもないホーム サーバーを構築しました。ここでは、この種のセットアップからある種の妥当なパフォーマンスを引き出すために何をする必要があったのかについて説明します。 (これには既存のコードとテクニックを使用し、llama.cpp 設定などをいじります。新しい ROCm カーネルの作成はこの章の範囲外です)。
ここで背景について説明します。変圧器モデルがどのように機能するかをすでに理解している場合は、セクション 2 に進んでください。マルチ GPU 並列処理がどのように機能するかをすでに知っていて、テストしている部分だけを知りたい場合は、セクション 3 に進んでください。
基本的に、現在どこでも使用されている AI テキスト生成ソフトウェアはすべて、「トランスフォーマー モデル」または「大規模言語モデル」のインスタンスです。設計と基本的なテクニックは、論文「アテンション・イズ・オール・ユー・ニード」で紹介されました。これは、おそらく過去のコンピューターサイエンスの分野で最も重要な論文です。20年も経ったかわかりませんが、この論文は、これらの点に関する限り、かなり読みやすいです。このブログ記事を読んでいるソフトウェア エンジニアは全員、すでに読んでいるはずです。 （そうですよね？）
とにかく。ここでは物事を少し単純化しすぎて、これらのことをくだらないハードウェアで高速に実行させようとしている人の観点に焦点を当て、テンソル計算について深く掘り下げたり、モデルのトレーニングについてはあまり話さないようにします。なぜなら、このブログ投稿はすでに非常に長すぎるためです。
LLM は「トークン」の観点から機能します。トークンは基本的に単語の断片です。個々の文字を入力および出力するのではなく、文字を一連の文字に分割し、モデルに処理させる方が効率的です。 (これが、初期の AI モデルが「文字 R は何回ですか」のような質問に正しく答えることができなかった理由です。

モデルが異なれば、言語のトークン化方法も異なります。これは、周波数エンコーディングのようなものと考えることができます。モデルが別のトークンを生成するたびに、実際には確率分布が生成され、その分布からランダムに 1 つがサンプリングされます。言語は、次に最も可能性の高いものを毎回正確に選択するよりも、その方がうまく機能するためです。
言語モデルは、複数の層に分割されたニューラル ネットワークです。入力層と出力層があり、その間には入力または出力と直接対話しない多数の層 (「隠れ層」) があります。入力層は入力プロンプト全体を受け取り、各層は前の層の出力に対して連続して計算を実行します。モデルが入力全体を一度に調べ、入力系列の異なる時点で異なるトークン間の関係を調べるものは、「アテンション メカニズム」と呼ばれます。 AI 言語モデルについて読んだことがあるなら、誰かが「これらの AI モデルは、マルコフ連鎖と同じように、単なる次の単語を生成するものである」と自信を持って主張しているのを聞いたことがあるでしょう。そしておそらく、これらの AI モデルがマルコフ連鎖とはまったく異なる出力を生成することに気づき、その人が一体どこで間違ったのか疑問に思ったことがあるでしょう。マルコフ チェーンにはアテンション メカニズムはなく、シリーズ内の前のトークンに基づいて新しいトークンを生成するだけです。
したがって、次のトークンを生成するために、モデルはトークン化されたプロンプト全体を読み取り、これを埋め込み行列に変換し (各トークンは、長さが各レイヤーの非表示次元であるベクトルに変換されます)、次に各レイヤーでアテンション計算を実行し (トークンごとに、直列のレイヤーごとに巨大な行列の乗算)、新しいトークンをサンプリングしてプロンプトに追加し、トークンに到達するまでこれをループで繰り返します。

停止する時期が来たことを示します。
私たちの目的にとって、これが意味するのは、コンピューターがトークンを生成するたびに、トークンを生成するために行列の乗算をすべて実行するために、既存のコンテキストとモデル内のすべての重みを読み取る必要があるということです。 Gemma4-31B モデルについては前の章で述べました。名前が示すように、モデルには 310 億の重みがあります (60 のレイヤーに分割されています)。次のトークンを計算するには、それらをすべて GPU にロードする必要があります。これらのロードには、実際の注意の計算よりもはるかに長い時間がかかります。トークンの生成は、（通常は）コンピューティングではなくメモリ帯域幅によって制限されます。これが、私が構築したサーバーに大量の VRAM が接続されたすべての GPU を搭載している理由です。モデルの重みと KV キャッシュは、GPU にすぐにアクセスできる VRAM 内にある必要があります。それに比べて、CPU に接続されているメモリははるかに遅いです。 (これは、業界全体がデータセンター GPU 用の高帯域幅メモリの生産を優先する方向に移行しているため、現在メモリ不足に陥っている理由でもあります)。
これは明らかにあまりうまく拡張できません。モデルのサイズが大きくなると、トークンの生成が大幅に遅くなります。この種のことに対する実際的な制限はすでに達成されています。これに応えて、誤解を招くような名前の「専門家の混合」モデル アーキテクチャが誕生しました 1 。ここでの考え方は、最初のいくつかの層が入力埋め込み行列全体を処理するために毎回使用され、その後、これがモデルの一部のサブセットにルーティングされるということです。中間層の場合、各入力トークンは重みのサブセットにルーティングされるため、各トークンに対してロードする必要があるのは重みの一部のセットだけです。これらのサブセットのそれぞれは「エキスパート」と呼ばれますが、私はこの枠組みが本当に嫌いです。なぜなら、これらのブランチの 1 つが Python について知っており、ブランチの 1 つがロケット工学について知っているという完全に誤った印象を与えるからです。

ines とそのうちの 1 人がドイツ語を話すことができるので、モデルの不要な部分をトリミングするだけで済みます。しかし、決してそうではありません。 「専門家」は基本的にランダムであるか、少なくとも予測不可能であり、ほとんどの MoE モデルでは、異なるトークンがほぼ均一な分布で異なる専門家にルーティングされます。したがって、たとえば、Deepseek V4 Flash には 2,840 億の重みがありますが、各トークンについてロードして処理するのは 130 億のみです。 (これの短縮形は「284B-A13B」で、13B のみが「アクティブ化」されます)。しかし、事前にどれが 130 億であるかはわかりません。また、生成されるトークンごとに変化するため、トークンの生成を高速に保つために、2,840 億の重みすべてをかなり高速なメモリに保存する必要があります。 MoE のおかげで、より多くの VRAM が必要になりますが、トークンの生成が高速になります。
大量の VRAM を使用する場合、それは個別の GPU に分割されます。私のサーバーには 4 台があり、それぞれ 32GB です。本格的なビジネス向けのものには、それぞれ 192GB 程度の GPU が 8 つあり、非常に大きなモデルを複数の計算ノードに分割する必要がありますが、原理は同じです。 GPU は自身のメモリを非常に速く読み取ることができますが、他の GPU からのメモリはそれほど速くは読み取れず、まったく別の計算ノード上のメモリはさらに遅くなります。これを分割するにはいくつかの方法があります。私のユースケースにとって実際に重要なのは、そのうちの 2 つだけです。ガレージにボックスがあり、実際のユーザーは私だけで、最もスマートなモデルを自分にとって適切な速度で動作させようとしている場合です。
レイヤー並列 - モデルを取得し、さまざまな GPU にさまざまなレイヤーを配置します。新しいトークンを生成するとき、GPU 1 でいくつかのレイヤーを実行し、その中間状態を GPU 2 に移動して次のレイヤーのセットを処理する、というように繰り返します。これは非常に簡単です、それは

すべてのモデルの重みが VRAM に組み込まれます。ただし、各レイヤーは前のレイヤーの出力に依存するため、各レイヤーは直列に処理されます。そのため、ここで期待できる理論上の最高速度は、セット全体と同じ量の VRAM を搭載した場合に 1 つの GPU で得られる速度と基本的に同じになります。 (実際にはもう少し遅くなります)。 AMD V620 のメモリ帯域幅は 512 GB/秒です。モデルを 4 つの GPU レイヤー間で並列に分割する場合、1 つのプロンプトに対して、GPU 間でデータを移動するオーバーヘッドを差し引いた 512 GB/s のメモリ帯域幅が得られます。
Tensor Parallel - モデルを取得し、各層を複数の GPU に分割します。各レイヤーについて、各 GPU は行列の乗算の一部を計算し、その結果を他の GPU に送信し、同期して最終結果を計算してから次のレイヤーに進みます。理論的には、これによりメモリ帯域幅を並列化できるはずです。512GB/s のカードが 4 枚あり、モデルをテンソル並列に分割する場合、各層を通過するときに 2TB/s のメモリ帯域幅が得られるはずです。もちろん、GPU 間でデータを移動するオーバーヘッドは差し引かれます。残念ながら、それはそれ以上です。レイヤー並列の場合のように GPU ごとにトークンごとに 1 回だけではなく、レイヤーごとに複数回データを移動します。これは、メモリ帯域幅の増加による速度の向上を上回る可能性があります (私のサーバーではそうなるでしょう)。
同時に複数のユーザーにサービスを提供するというコンテキストでより意味のある他の並列処理についても説明します。
データ並列は、複数の個別の GPU で同じモデルを実行し、受信クエリを並列実行するようにルーティングするだけです。将来的には複数のサブエージェントを同時に実行するためにこれを試してみるかもしれませんが、GPU は情報を処理できるため、非常に大規模になるまでは実際には必要ありません。

すでにバッチ内にあります。 (前述したように、推論パイプラインはメモリ帯域幅によって制限されているため、少なくとも再び計算に依存するようになるまでは、ある程度の追加の行列計算を「無料」でそこにスライドさせることができます)。
Expert Parallel は、すべての GPU で共有テンソルを保持し、共有されていない各「エキスパート」 (またはそのサブグループ) を異なる GPU に配置するため、複数のユーザーは、ほとんどの推論パイプラインでクエリを異なる GPU にルーティングできます。繰り返しになりますが、これは多くの同時ユーザーにとってほぼ意味があり、商用 LLM プロバイダーによって広く使用されています。ただし、それは私の場合にはあまり役に立ちません。
とにかく、私のサーバーはほとんどの場合、一度に 1 つのリクエストを処理します。ある時点で並列サブエージェントを試してみるつもりです。おそらく、これを同時に実行するゲストが数人いるかもしれませんが、実際には複数のユーザーがいるわけではないので、これらのケースについては最適化していません。
さて、私のサーバーについて話しましょう。
ここで使用している GPU アレイは、共有 PCI Express バスに接続された 4 枚の AMD Radeon Pro V620 カードです。 GPU 間のトラフィックは遅いと予想されますが、マザーボードが古い PCIe 3.0 標準を実行しており、ほとんどのカードが 16 ではなく 8 レーンを使用しているため、予想よりもさらに遅くなります。
jacob@daedalus:~$ sudo lspci -vvv
[...]
LnkCap: ポート #2、速度 16GT/s、幅 x16、ASPM L1、終了レイテンシー L1
lspci は 8 車線で LA の高速道路のように遅いことを示しています
したがって、テンソル並列はあまりうまく機能しないのではないかと思います。 AI ワークロード用に実際に設計および販売されているカードには、PCIe に加えて、ある種の高速かつ低遅延の GPU 間接続が備わっており、これらはすべてメーカー固有です。 Nvidia には NVLink、AMD には Infinity Fabric、Intel には Xe Link などがあります。これらのカードにはそのようなものはありません。アロウン

AMD がクラウド ゲーム スキーム用にこれらのカードを製造しているのと同じ時期に、AMD Instinct MI210 と呼ばれる AI に重点を置いたカードも製造していました。このカードには約 3 倍のメモリ帯域幅があり、Infinity Fabric 経由で最大 4 枚のカードを接続する特別なブリッジ コネクタを入手できます。また、これらのカードの 1 枚は、このボックス全体よりも (現在使用すると) 高価になるため、これはまったくスターターではありません。持っているものでなんとかしなければなりません。
私が持っているのは 4 枚のカードですが、それぞれのカードはかなり高速ですが、カード間の相互接続が不十分です。つまり、これをレイヤー並列で実行し、パイプラインを埋める方法を考える必要があります。そして、後で説明するように、レイヤーの並列化でもパフォーマンスに影響を与えます。
ここでの研究の目的のために、私は 2 つのモデルを使って遊んでみました。 Gemma4-31B と Deepseek V4 Flash を持っていました。
Gemma4 31B には 310 億のパラメーターがあり、4 ビット量子化での重みは約 18 GiB で、モデルはそれぞれ 32 GB の VRAM を備えたこれらの GPU の 1 つに快適に収まります。これは高密度のモデルです。すべてのトークンは、それらの重みの 18 GiB すべてを通過します。
Deepseek V4 Flash には 2,840 億のパラメータがあり、これは MoE モデルであり、トークンごとに 130 億のパラメータが有効になります。

[切り捨てられた]

## Original Extract

getting AI models to run fast on a box of e-waste GPUs

In the last chapter I built a ridiculous home server out of e-waste-grade GPUs to run AI language models. Here, I'll be talking about what I needed to do to squeeze some kind of reasonable performance out of this kind of setup. (This is going to be using existing code and techniques, messing around with llama.cpp settings and so on; writing new ROCm kernels is out of scope for this chapter).
I'm gonna go into some background here. If you already know how transformer models work, go ahead and jump to section 2 . If you already know how multi-GPU parallelism works and just want to get to the part where I'm testing things, jump to section 3 .
Okay, basically all of the AI text generation software that's currently in use everywhere are instances of the "transformer model" or "large language model". The design and basic technique was introduced in the paper Attention is All You Need , which is probably the most important paper in the field of computer science in the past, I dunno, twenty years? The paper is pretty readable as far as these things go. I'm sure every software engineer reading this blog post has already read it, right? (Right?)
Anyway. I'm going to over-simplify things a bit here and focus from the perspective of somebody who is trying to get these things to run fast on crappy hardware, and not go deep into the tensor math or talk much about model training, because this blog post is already going to be way, way too long.
The LLM works in terms of "tokens". A token is basically a word fragment; instead of inputting and outputting individual letters it's more efficient to chop these up into sequences of letters and have the model process those. (This is why early AI models were bad at correctly answering questions like "how many times is the letter R in the word raspberry?") Different models tokenize language differently, you can think of this as like a frequency encoding. Each time a model generates another token, it'll actually generate a probability distribution and then randomly sample one from that distribution, because language works better that way than picking the exact most likely next thing every time.
The language model is a neural network that's divided into layers. You have an input layer and an output layer and a bunch of layers in between that don't directly interact with the input or output ("hidden layers"). The input layer takes in the entire input prompt, and then each layer does math on the output of the previous layer in series. The thing where the model looks at the entire input at once, and looks at the relations between different tokens at different points in the input series, is called the "attention mechanism". If you've been reading about AI language models you probably have heard somebody confidently claim "these AI models are just next word generators, like a Markov chain is", and then you probably noticed that these AI models generate very different outputs than a Markov chain does, and wondered where exactly that guy went wrong. Well, a Markov chain doesn't have the attention mechanism, it just generates a new token based on the previous token in the series.
So, to generate the next token, the model reads in the entire tokenized prompt, turns this into an embedding matrix (each token gets turned into a vector where the length is the hidden dimension of each layer), then does the attention math on each layer (gigantic matrix multiplication for each token, for each layer in series), samples a new token, adds it to the prompt, and keeps doing this in a loop until it gets to a token that indicates that it's time to stop.
For our purposes, what this means is that every time the computer generates a token, it needs to read in the existing context, and also every weight in the model, in order to do all those matrix multiplications to generate the token. I'd mentioned the Gemma4-31B model in the last chapter; as the name implies, the model has 31 billion weights (divided into 60 layers). We've got to load all of them into the GPU to calculate the next token. Loading these takes a lot longer than the actual attention math does; token generation is (usually) limited by memory bandwidth rather than compute. This is why the server I built has all of those GPUs with lots of VRAM attached to them; the model weights and KV cache need to be in VRAM that can get to the GPU quickly. The memory attached to the CPU is by comparison a lot slower. (This is also why we are in a memory shortage right now as the entire industry shifts to prioritize producing high-bandwidth memory for data center GPUs).
This obviously isn't going to scale super well. As model sizes increase, token generation slows way down; practical limits on this kind of thing got hit already. In response, we have the misleadingly-named "Mixture of Experts" model architecture 1 . The idea here is that the first couple of layers are used every time to process the whole input embedding matrix, then this gets routed to some subset of the model. For the middle layers, each input token gets routed to a subset of the weights, so only some set of the weights will need to get loaded for each token. Each of these subsets is called an "expert", and I really hate this framing, because it gives you the completely false impression that like one of these branches knows about Python and one of them knows about rocket engines and one of them knows how to speak German and you can just trim the parts of the model you don't care about. This is absolutely not the case, though! The "experts" are basically random, or at least unpredictable, and for most MoE models, different tokens will get routed to different experts in a mostly uniform distribution. So, for example, Deepseek V4 Flash has 284 billion weights but for each token we will only load and process 13 billion. (The shorthand for this is "284B-A13B", only 13B get "activated"). But I don't know which 13 billion ahead of time, and it's going to change for each token that gets generated, so I still need to store all 284 billion weights in some pretty fast memory to keep token generation fast. The MoE thing makes token generation fast at the cost of needing more VRAM.
When you're using a lot of VRAM, it's going to be divided among separate GPUs. My server has four, with 32GB each; the serious business ones will have eight GPUs with like 192GB each, and then will have to split really big models up between multiple compute nodes, but the principles are the same. A GPU can read its own memory pretty quickly and memory from some other GPU not very quickly and memory on a whole different compute node will be slower still. There's several different ways to split this up. Only two of them really matter for the use case I have, where I've got a box in the garage and I'm the only real user and I'm trying to get the smartest model working at adequate speeds for myself.
Layer Parallel - We take the model and put different layers on the different GPUs. As we're generating a new token, we run through a few of the layers on GPU 1, then move that intermediate state over to GPU 2 and process the next set of layers, and so on. This is pretty simple, it gets all of our model weights into VRAM. Each layer is getting processed in series, though, because each layer depends on the output of the previous layer, and so the theoretical best speed we can expect here is basically the same as what one GPU would give us if it had as much VRAM as the whole set. (In practice it'll be a bit slower). An AMD V620 has 512 GB/s of memory bandwidth. If I'm splitting the model up between four GPUs layer parallel, for one prompt I'm gonna get... 512 GB/s of memory bandwidth, minus the overhead of moving data between the GPUs.
Tensor Parallel - We take the model and split each layer between multiple GPUs. For each layer, each GPU calculates a fraction of that matrix multiplication, then sends that result to some other GPU that will synchronize and calculate the final result before moving on to the next layer. In theory this should allow us to parallelize our memory bandwidth: I have four cards at 512GB/s, if I'm splitting the model up tensor parallel I should get 2TB/s of memory bandwidth as I move through each layer! Minus the overhead of moving data between the GPUs, of course. Which is, unfortunately, way more! We're moving data around multiple times for each layer, instead of just once per GPU per token as in layer parallel. This can (and will, on my server) outweigh the speed increase from the increased memory bandwidth.
You'll also see discussion of other parallelisms that make more sense in the context of serving multiple users at once:
Data Parallel is just running the same model on multiple separate GPUs and routing incoming queries to run in parallel. I may play around with that for multiple subagents running simultaneously in the future, but it's not actually necessary until you get to very large scale because a GPU can process inference in a batch already. (The inference pipeline is bound by memory bandwidth, as mentioned earlier, so you can slide some amount of extra matrix math in there for "free", at least until you try to do so much that it becomes compute-bound again).
Expert Parallel keeps shared tensors on every GPU and then puts each of the non-shared "experts" (or subgroups of them) onto different GPUs, so multiple users would get their queries routed to different GPUs for most of the inference pipeline. Again, this one mostly makes sense for lots of concurrent users, and is widely used by commercial LLM providers. It doesn't really help my case though.
Anyway, my server will mostly be serving me making one request at a time. I'll play around with parallel subagents at some point, maybe I'll have a couple guests hitting this thing at the same time, but I don't really have multiple users here so I'm not optimizing for these cases.
Okay, let's talk about my server.
The GPU array I have here is four AMD Radeon Pro V620 cards connected to a shared PCI Express bus. You'd expect the inter-GPU-traffic to be slow, and it's even slower than you'd expect, because the motherboard is running the older PCIe 3.0 standard and most of the cards are using 8 lanes instead of 16.
jacob@daedalus:~$ sudo lspci -vvv
[...]
LnkCap: Port #2, Speed 16GT/s, Width x16, ASPM L1, Exit Latency L1
lspci shows eight lanes and slow, like the LA freeways
So I suspect that tensor parallel isn't going to work very well. The cards that are actually designed and marketed for AI workloads have some kind of high speed, low latency inter-GPU connection in addition to PCIe, and these are all manufacturer-specific; Nvidia has NVLink, AMD has Infinity Fabric, Intel has Xe Link, etc. These cards don't have anything like that. Around the same time AMD was making these cards for their cloud gaming scheme, they were making an AI-focused card called the AMD Instinct MI210, it has about three times the memory bandwidth and you can get a special bridge connector that connects up to four of them via Infinity Fabric. Also, a single one of those cards would cost (used, today) more than this whole box, so that's all kind of a non-starter. I have to make do with what I have.
What I have are four cards that are respectably fast individually and have poor interconnects between them, which means we're going to have to do this in layer parallel, then figure out how to fill up the pipeline. And even layer parallel has a performance impact, as we'll see.
For the purposes of this study here I had two models I was playing with; I had Gemma4-31B, and Deepseek V4 Flash.
Gemma4 31B has 31 billion parameters, at a 4-bit quantization the weights are about 18 GiB and the model fits comfortably into a single one of these GPUs that each have 32GB of VRAM. It's a dense model; every token will run through all 18 GiB of those weights.
Deepseek V4 Flash has 284 billion parameters, it's a MoE model and 13 billion parameters will activate for each token.

[truncated]
