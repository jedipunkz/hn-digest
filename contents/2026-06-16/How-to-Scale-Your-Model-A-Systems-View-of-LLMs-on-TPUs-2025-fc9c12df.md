---
source: "https://jax-ml.github.io/scaling-book/"
hn_url: "https://news.ycombinator.com/item?id=48554799"
title: "How to Scale Your Model – A Systems View of LLMs on TPUs (2025)"
article_title: "How To Scale Your Model"
author: "frenchmajesty"
captured_at: "2026-06-16T13:56:20Z"
capture_tool: "hn-digest"
hn_id: 48554799
score: 1
comments: 0
posted_at: "2026-06-16T13:14:26Z"
tags:
  - hacker-news
  - translated
---

# How to Scale Your Model – A Systems View of LLMs on TPUs (2025)

- HN: [48554799](https://news.ycombinator.com/item?id=48554799)
- Source: [jax-ml.github.io](https://jax-ml.github.io/scaling-book/)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T13:14:26Z

## Translation

タイトル: モデルをスケールする方法 – TPU 上の LLM のシステム ビュー (2025)
記事のタイトル: モデルをスケールする方法
説明: LLM のトレーニングは錬金術のように感じることがよくありますが、モデルのパフォーマンスを理解して最適化することは錬金術のように感じられません。

記事本文:
モデルをスケールする方法 ナビゲーションを切り替え
セクション パート 0. はじめに パート 1. ルーフラインの概要 パート 2. TPU のすべて パート 3. シャード Matmuls パート 4. トランスフォーマー パート 5. トレーニング パート 6. LLaMA のトレーニング パート 7. 推論 パート 8. LLaMA の提供 パート 9. プロファイリング パート 10. JAX のすべて パート 11. 結論 パート 12. GPU
TPU 上の LLM のシステム ビュー (パート 0: 概要 | パート 1: ルーフライン)
LLM のトレーニングは錬金術のように感じることがよくありますが、モデルのパフォーマンスを理解して最適化することは必ずしも必要ではありません。この本の目的は、TPU (および GPU) がどのように動作し、相互に通信するのか、LLM が実際のハードウェアでどのように実行されるのか、大規模なスケールで効率的に実行できるようにトレーニングと推論中にモデルを並列化する方法など、スケーリング言語モデルの科学をわかりやすく説明することです。 「この LLM のトレーニングにどれくらいのコストがかかるのか」、「このモデルを自分で提供するにはどれくらいのメモリが必要か」、「AllGather とは何なのか」などと疑問に思ったことがある場合は、これが役立つことを願っています。
深層学習の多くは依然として一種の黒魔術に帰着しますが、モデルのパフォーマンスを最適化する必要はありません。それは大規模な場合であってもです。比較的単純な原則は、単一のアクセラレータの処理から数万のアクセラレータの処理に至るまで、あらゆる場所に適用され、それらを理解することで多くの便利なことが可能になります。
モデルの各部分が理論上の最適値にどの程度近づいているかを概算します。
さまざまなスケール (複数のデバイス間で計算を分割する方法) でのさまざまな並列処理スキームについて、情報に基づいた選択を行います。
大規模な Transformer モデルのトレーニングと実行に必要なコストと時間を見積もります。
特定のハードウェア アフォーダンスを活用するアルゴリズムを設計します。
現在のアルゴリズムのパフォーマンスを制限するものを明確に理解した上でハードウェアを設計します。
予想される背景: 行きます

LLM と Transformer アーキテクチャについては基本的に理解していると仮定しますが、必ずしもそれらが大規模にどのように動作するかは理解していません。 LLM トレーニングの基本を理解しており、理想的には JAX についての基本的な知識を持っている必要があります。有益な背景資料としては、Transformer アーキテクチャに関するこのブログ投稿や、Transformer のオリジナルの論文などが挙げられます。また、同時に読んだり将来読んだりするのに役立つこのリストもチェックしてください。
目標とフィードバック: 最終的には、特定のハードウェア プラットフォーム上で Transformer モデルに最適な並列処理スキームと、トレーニングと推論にかかるおおよその時間を快適に見積もれるようになっているはずです。そうでない場合は、メールでご連絡いただくか、コメントを残してください。これをより明確にする方法を知りたいと思っています。
NVIDIA GPU に関する新しいセクション 12 もぜひ読んでみてください。
3 ～ 4 年前であれば、ほとんどの ML 研究者はこの本の内容を理解する必要はなかったと思います。しかし、今日では「小規模」モデルでさえハードウェアの限界に非常に近い状態で動作するため、新しい研究を行うには大規模な効率を考慮する必要があります。歴史的に、ML 研究は、システムの革新とソフトウェアの改善の間のチクタク サイクルのようなものに従ってきました。 Alex Krizhevsky は、CNN を高速化するためにひどい CUDA コードを書かなければなりませんでしたが、数年以内に、Theano や TensorFlow のようなライブラリにより、その必要はなくなりました。おそらくここでもそれが起こり、この本のすべてが数年後には抽象化されるでしょう。しかし、スケーリングの法則により、私たちのモデルは常にハードウェアの最前線に押し上げられており、おそらく近い将来、最先端の研究を行うことは、モデルを大規模なハードウェア トポロジに効率的に拡張する方法の理解と密接に結びついていくと思われます。ルーフリンに 20% のコストがかかる場合、ベンチマークで 20% の勝利を収めても意味がありません。

効率性。有望なモデル アーキテクチャは、大規模に効率的に実行できないか、実行するための労力を誰も投入しないため、定期的に失敗します。
「モデル スケーリング」の目標は、スループットの比例的かつ線形的な増加を達成しながら、トレーニングまたは推論に使用するチップの数を増やすことができるようにすることです。これは「強力なスケーリング」として知られています。通常、チップを追加すると (「並列処理」)、計算時間が短縮されますが、チップ間の通信が追加されるというコストもかかります。通信に計算よりも時間がかかると、「通信の制約」が生じ、強力に拡張できなくなります。計算時間が短縮されると、通常、単一チップのレベルでボトルネックにも直面します。ピカピカの新しい TPU または GPU は 1 秒あたり 500 兆回の演算を実行すると評価されているかもしれませんが、メモリ内でパラメータを移動するのに行き詰った場合、注意しないとその 10 分の 1 も簡単に実行してしまう可能性があります。チップごとの計算、メモリ帯域幅、総メモリの相互作用は、スケーリングのストーリーにとって重要です。ハードウェアを十分に理解してこれらのボトルネックがどこで発生するかを予測できれば、ボトルネックを回避するようにモデルを設計または再構成できます。ハードウェア設計者は、コストを最小限に抑えながら、アルゴリズムに必要なだけのコンピューティング、帯域幅、メモリを提供するハードウェアを構築するという、逆の問題に直面しています。この「共同設計」問題がどれほどストレスフルであるかは想像できるでしょう。最初のチップが実際に利用可能になるとき、多くの場合 2 ～ 3 年後、アルゴリズムがどのようになるかに賭けなければなりません。 TPU のストーリーは、このゲームで大成功を収めています。行列乗算は、他のほとんどのメモリよりもはるかに多くのメモリ バイトあたりの FLOP (バイトあたり N FLOP)、および初期の TPU とその

ystolic アレイ アーキテクチャは、GPU が構築された時点よりもはるかに優れたパフォーマンス / $ を達成しました。 TPU は ML ワークロード向けに設計されており、Tensor コアを搭載した GPU もこのニッチ市場を満たすために急速に変化しています。しかし、もしニューラル ネットワークが普及していなかったら、あるいは TPU (本質的に低消費電力) よりも根本的な方法で変化していたら、どれほどのコストがかかったであろうか、想像できるでしょう。
[切り捨てられた]
本書の目的は、TPU (および GPU) ハードウェアがどのように動作するか、および現在のハードウェアで良好なパフォーマンスを発揮するために Transformer アーキテクチャがどのように進化したかを説明することです。これが、新しいアーキテクチャを設計する研究者と、現世代の LLM を高速に実行することに取り組むエンジニアの両方にとって役立つことを願っています。
本書の全体構成は次のとおりです。
セクション 1 では、ルーフライン分析と、拡張能力 (通信、計算、メモリ) を制限する要因について説明します。セクション 2 とセクション 3 では、TPU が個々のチップとして、および帯域幅と遅延が制限されたチップ間リンクを備えた相互接続システムとして (非常に重要) としてどのように機能するかについて詳しく説明します。次のような質問に答えます。
特定のサイズの行列の乗算にはどれくらい時間がかかりますか?どの時点でコンピューティング、メモリ、または通信帯域幅によって制限されますか?
TPU はどのように接続されてトレーニング クラスターを形成するのでしょうか?システムの各部分にはどれくらいの帯域幅がありますか?
複数の TPU にわたって配列を収集、分散、または再分散するにはどのくらい時間がかかりますか?
デバイス間で異なる分布をもつ行列を効率的に乗算するにはどうすればよいでしょうか?
図: TPU が要素単位の積を実行する方法を示すセクション 2 の図。配列のサイズとさまざまなリンクの帯域幅に応じて、コンピューティング バウンド (ハードウェアのコンピューティング能力をすべて使用する) またはメモリ バウンド (ボトルネック) が発生する可能性があります。

d メモリロードによる）。 5 年前、ML には ConvNet、LSTM、MLP、Transformer などのカラフルなアーキテクチャがありましたが、現在ではほとんどが Transformer だけになっています。私たちは、Transformer アーキテクチャのすべての部分を理解することは価値があると強く信じています。つまり、各行列の正確なサイズ、正規化が行われる場所、パラメーターと FLOP の数、浮動小数点 OP の数、基本的に必要な加算と乗算の合計数です。多くのソースでは FLOP を「1 秒あたりの操作数」の意味として捉えていますが、私たちはそれを明示的に示すために FLOP/秒 を使用します。各部にあります。セクション 4 では、この「トランスフォーマーの計算」を注意深く検討し、トレーニングと推論の両方でパラメーターと FLOP をカウントする方法を示します。これにより、モデルが使用するメモリの量、計算や通信に費やす時間、フィードフォワード ブロックに関連して注意が重要になる時期がわかります。
図: 各行列乗算 (matmul) が円内のドットとして示されている標準の Transformer レイヤー。すべてのパラメータ (ノルムを除く) は紫色で表示されます。セクション 4 では、この図を詳しく説明します。セクション 5: トレーニングとセクション 7: 推論はこの本の核心であり、そこでは基本的な問題について説明します。つまり、ある程度のサイズとチップ数のモデルが与えられた場合、「強力なスケーリング」体制に留まるようにモデルを並列化するにはどうすればよいでしょうか?これは単純な質問ですが、その答えは驚くほど複雑です。大まかに言うと、モデルを複数のチップに分割するために使用される 4 つの主要な並列処理手法 ( data 、 tensor 、 Pipeline 、および Expert ) と、メモリ要件を削減するためのその他の多数の手法 ( rematerialization 、 オプティマイザ/モデル シャーディング (別名 ZeRO) 、 ホスト オフロード 、 勾配累積 ) があります。ここではこれらの多くについて説明します。
これらのセクションが終わるまでに、あなたが次のレベルに達していることを願っています。

新しいアーキテクチャや設定のためにそれらの中から自分で選択することができます。セクション 6 とセクション 8 は、これらの概念を人気のあるオープンソース モデルである LLaMA 3 に適用する実践的なチュートリアルです。
最後に、セクション 9 とセクション 10 では、これらのアイデアの一部を JAX で実装する方法と、問題が発生したときにコードをプロファイリングしてデバッグする方法について説明します。セクション 12 は、GPU についても詳しく説明する新しいセクションです。
私たちは全体を通じて、あなたが自分で取り組むことができる問題を提供するよう努めています。すべてのセクションを読むか、順番に読むかというプレッシャーを感じる必要はありません。そしてフィードバックを残してください。現時点ではこれは草案であり、今後も修正が加えられていく予定です。ありがとう！
この本のアイデアの多くを導き出した James Bradbury と Blake Hechtman に感謝したいと思います。
さっそく、セクション 1 で TPU ルーフラインについて説明します。
このシリーズはおそらく必要以上に長くなりますが、それでも気が進まないことを願っています。最初の 3 章は準備編であり、内容をすでによく知っている場合は読み飛ばしていただいても構いません。ただし、後ほど使用する表記法について説明します。最後の 3 つの部分は、実際のモデルの操作方法を説明しているため、最も実用的です。
第 1 章: ルーフライン解析の簡単な紹介。アルゴリズムは、計算、通信、メモリという 3 つの要素によって制限されます。これらを使用して、アルゴリズムの実行速度を概算できます。
第 2 章: TPU についての考え方。 TPU はどのように機能するのでしょうか?それは、トレーニングして提供できるモデルにどのような影響を与えるでしょうか?
第 3 章: シャード行列とその乗算方法。ここでは、モデルのシャーディングとマルチ TPU 並列処理を、私たちのお気に入りの操作である (シャーディングされた) 行列乗算によって説明します。
第 4 章: 知っておくべき変圧器の数学のすべて。 Transformer は前方パスと後方パスで何 FLOP を使用しますか?パラメータの数を計算できますか? T

KV キャッシュのサイズは?ここではこの計算を行っていきます。
第 5 章: トレーニング用にトランスフォーマーを並列化する方法。 FSDP。メガトロンのシャーディング。パイプラインの並列処理。ある程度のチップ数がある場合、特定のサイズのモデルを特定のバッチ サイズでできるだけ効率的にトレーニングするにはどうすればよいでしょうか?
第 6 章: TPU での LLaMA 3 のトレーニング。 TPU で LLaMA 3 をトレーニングするにはどうすればよいでしょうか?どのくらい時間がかかりますか?費用はいくらくらいかかりますか？
第 7 章: トランスフォーマー推論のすべて。モデルをトレーニングしたら、それを提供する必要があります。推論では、レイテンシーという新たな考慮事項が追加され、メモリの状況が変化します。細分化された配信の仕組みと、KV キャッシュについての考え方について説明します。
第 8 章: TPU での LLaMA 3 の提供。 TPU v5e で LLaMA 3 を提供するにはいくらかかりますか?レイテンシとスループットのトレードオフは何ですか?
第 9 章: TPU コードをプロファイリングする方法。実際の LLM は、上記の理論ほど単純ではありません。ここでは、JAX + XLA スタックと、JAX/TensorBoard プロファイラーを使用して実際の問題をデバッグおよび修正する方法について説明します。
第 10 章: JAX での TPU のプログラミング。 JAX は計算を並列化するための魔法の API を多数提供していますが、それらの使用方法を知る必要があります。楽しい例とうまくいった問題。
パート 4: 結論とボーナス コンテンツ
第 11 章: 結論とさらなる読み物。終わりの思いとf

[切り捨てられた]

## Original Extract

Training LLMs often feels like alchemy, but understanding and optimizing the performance of your models doesn

How To Scale Your Model Toggle navigation
Sections Part 0. Introduction Part 1. Intro to Rooflines Part 2. All About TPUs Part 3. Sharded Matmuls Part 4. Transformers Part 5. Training Part 6. Training LLaMA Part 7. Inference Part 8. Serving LLaMA Part 9. Profiling Part 10. All About JAX Part 11. Conclusions Part 12. GPUs
A Systems View of LLMs on TPUs (Part 0: Intro | Part 1: Rooflines )
Training LLMs often feels like alchemy, but understanding and optimizing the performance of your models doesn't have to. This book aims to demystify the science of scaling language models: how TPUs (and GPUs) work and how they communicate with each other, how LLMs run on real hardware, and how to parallelize your models during training and inference so they run efficiently at massive scale. If you've ever wondered "how expensive should this LLM be to train" or "how much memory do I need to serve this model myself" or "what's an AllGather", we hope this will be useful to you.
Much of deep learning still boils down to a kind of black magic, but optimizing the performance of your models doesn’t have to — even at huge scale! Relatively simple principles apply everywhere — from dealing with a single accelerator to tens of thousands — and understanding them lets you do many useful things:
Ballpark how close parts of your model are to their theoretical optimum.
Make informed choices about different parallelism schemes at different scales (how you split the computation across multiple devices).
Estimate the cost and time required to train and run large Transformer models.
Design algorithms that take advantage of specific hardware affordances .
Design hardware driven by an explicit understanding of what limits current algorithm performance.
Expected background: We’re going to assume you have a basic understanding of LLMs and the Transformer architecture but not necessarily how they operate at scale. You should know the basics of LLM training and ideally have some basic familiarity with JAX. Some useful background reading might include this blog post on the Transformer architecture and the original Transformer paper . Also check out this list for more useful concurrent and future reading.
Goals & Feedback: By the end, you should feel comfortable estimating the best parallelism scheme for a Transformer model on a given hardware platform, and roughly how long training and inference should take. If you don’t, email us or leave a comment! We’d love to know how we could make this clearer.
You might also enjoy reading the new Section 12 on NVIDIA GPUs!
Three or four years ago, I don’t think most ML researchers would have needed to understand any of the content in this book. But today even “small” models run so close to hardware limits that doing novel research requires you to think about efficiency at scale. Historically, ML research has followed something of a tick-tock cycle between systems innovations and software improvements. Alex Krizhevsky had to write unholy CUDA code to make CNNs fast but within a couple of years, libraries like Theano and TensorFlow meant you didn't have to. Maybe that will happen here too and everything in this book will be abstracted away in a few years. But scaling laws have pushed our models perpetually to the very frontier of our hardware, and it seems likely that, for the foreseeable future, doing cutting-edge research will be inextricably tied to an understanding of how to efficiently scale models to large hardware topologies. A 20% win on benchmarks is irrelevant if it comes at a 20% cost to roofline efficiency. Promising model architectures routinely fail either because they can’t run efficiently at scale or because no one puts in the work to make them do so.
The goal of “model scaling” is to be able to increase the number of chips used for training or inference while achieving a proportional, linear increase in throughput. This is known as “ strong scaling ”. Although adding additional chips (“parallelism”) usually decreases the computation time, it also comes at the cost of added communication between chips. When communication takes longer than computation we become “communication bound” and cannot scale strongly. As your computation time decreases, you also typically face bottlenecks at the level of a single chip. Your shiny new TPU or GPU may be rated to perform 500 trillion operations per second, but if you aren't careful it can just as easily do a tenth of that if it's bogged down moving parameters around in memory. The interplay of per-chip computation, memory bandwidth, and total memory is critical to the scaling story. If we understand our hardware well enough to anticipate where these bottlenecks will arise, we can design or reconfigure our models to avoid them. Hardware designers face the inverse problem: building hardware that provides just enough compute, bandwidth, and memory for our algorithms while minimizing cost. You can imagine how stressful this "co-design" problem is: you have to bet on what algorithms will look like when the first chips actually become available, often 2 to 3 years down the road. The story of the TPU is a resounding success in this game. Matrix multiplication is a unique algorithm in the sense that it uses far more FLOPs per byte of memory than almost any other (N FLOPs per byte), and early TPUs and their systolic array architecture achieved far better perf / $ than GPUs did at the time they were built. TPUs were designed for ML workloads, and GPUs with their Tensor Cores are rapidly changing to fill this niche as well. But you can imagine how costly it would have been if neural networks had not taken off, or had changed in some fundamental way that TPUs (which are inherently less f
[truncated]
Our goal in this book is to explain how TPU (and GPU) hardware works and how the Transformer architecture has evolved to perform well on current hardware. We hope this will be useful both for researchers designing new architectures and for engineers working to make the current generation of LLMs run fast.
The overall structure of this book is as follows:
Section 1 explains roofline analysis and what factors can limit our ability to scale (communication, computation, and memory). Section 2 and Section 3 talk in detail about how TPUs work, both as individual chips and — of critical importance — as an interconnected system with inter-chip links of limited bandwidth and latency. We’ll answer questions like:
How long should a matrix multiply of a certain size take? At what point is it bound by compute or by memory or communication bandwidth?
How are TPUs wired together to form training clusters? How much bandwidth does each part of the system have?
How long does it take to gather, scatter, or re-distribute arrays across multiple TPUs?
How do we efficiently multiply matrices that are distributed differently across devices?
Figure: a diagram from Section 2 showing how a TPU performs an elementwise product. Depending on the size of our arrays and the bandwidth of various links, we can find ourselves compute-bound (using the full hardware compute capacity) or memory-bound (bottlenecked by memory loading). Five years ago ML had a colorful landscape of architectures — ConvNets, LSTMs, MLPs, Transformers — but now we mostly just have the Transformer . We strongly believe it’s worth understanding every piece of the Transformer architecture: the exact sizes of every matrix, where normalization occurs, how many parameters and FLOPs FLoating point OPs, basically the total number of adds and multiplies required. While many sources take FLOPs to mean "operations per second", we use FLOPs/s to indicate that explicitly. are in each part. Section 4 goes through this “Transformer math” carefully, showing how to count the parameters and FLOPs for both training and inference. This tells us how much memory our model will use, how much time we’ll spend on compute or comms, and when attention will become important relative to the feed-forward blocks.
Figure: a standard Transformer layer with each matrix multiplication (matmul) shown as a dot inside a circle. All parameters (excluding norms) are shown in purple. Section 4 walks through this diagram in more detail. Section 5: Training and Section 7: Inference are the core of this book, where we discuss the fundamental question: given a model of some size and some number of chips, how do I parallelize my model to stay in the “strong scaling” regime? This is a simple question with a surprisingly complicated answer. At a high level, there are four primary parallelism techniques used to split models over multiple chips ( data , tensor , pipeline , and expert ), and a number of other techniques to reduce the memory requirements ( rematerialization , optimizer/model sharding (aka ZeRO) , host offload , gradient accumulation ). We discuss many of these here.
We hope by the end of these sections you should be able to choose among them yourself for new architectures or settings. Section 6 and Section 8 are practical tutorials that apply these concepts to LLaMA 3, a popular open-source model.
Finally, Section 9 and Section 10 look at how to implement some of these ideas in JAX and how to profile and debug your code when things go wrong. Section 12 is a new section that dives into GPUs as well.
Throughout we try to give you problems to work for yourself. Please feel no pressure to read all the sections or read them in order. And please leave feedback. For the time being, this is a draft and will continue to be revised. Thank you!
We’d like to acknowledge James Bradbury and Blake Hechtman who derived many of the ideas in this book.
Without further ado, here is Section 1 about TPU rooflines.
This series is probably longer than it needs to be, but we hope that won’t deter you. The first three chapters are preliminaries and can be skipped if you’re already familiar with the material, although they introduce notation used later. The final three parts might be the most practically useful, since they explain how to work with real models.
Chapter 1: A Brief Intro to Roofline Analysis . Algorithms are bounded by three things: compute, communication, and memory. We can use these to approximate how fast our algorithms will run.
Chapter 2: How to Think About TPUs . How do TPUs work? How does that affect what models we can train and serve?
Chapter 3: Sharded Matrices and How to Multiply Them . Here we explain model sharding and multi-TPU parallelism by way of our favorite operation: (sharded) matrix multiplications.
Chapter 4: All the Transformer Math You Need to Know . How many FLOPs does a Transformer use in its forward and backward pass? Can you calculate the number of parameters? The size of its KV caches? We work through this math here.
Chapter 5: How to Parallelize a Transformer for Training . FSDP. Megatron sharding. Pipeline parallelism. Given some number of chips, how do I train a model of a given size with a given batch size as efficiently as possible?
Chapter 6: Training LLaMA 3 on TPUs . How would we train LLaMA 3 on TPUs? How long would it take? How much would it cost?
Chapter 7: All About Transformer Inference . Once we’ve trained a model, we have to serve it. Inference adds a new consideration — latency — and changes up the memory landscape. We’ll talk about how disaggregated serving works and how to think about KV caches.
Chapter 8: Serving LLaMA 3 on TPUs . How much would it cost to serve LLaMA 3 on TPU v5e? What are the latency/throughput tradeoffs?
Chapter 9: How to Profile TPU Code . Real LLMs are never as simple as the theory above. Here we explain the JAX + XLA stack and how to use the JAX/TensorBoard profiler to debug and fix real issues.
Chapter 10: Programming TPUs in JAX . JAX provides a bunch of magical APIs for parallelizing computation, but you need to know how to use them. Fun examples and worked problems.
Part 4: Conclusions and Bonus Content
Chapter 11: Conclusions and Further Reading . Closing thoughts and f

[truncated]
