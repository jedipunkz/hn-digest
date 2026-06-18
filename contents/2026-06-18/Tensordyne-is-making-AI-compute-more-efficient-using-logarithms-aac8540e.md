---
source: "https://www.zach.be/p/tensordyne-is-making-ai-compute-more"
hn_url: "https://news.ycombinator.com/item?id=48578295"
title: "Tensordyne is making AI compute more efficient using logarithms"
article_title: "Tensordyne is making AI compute more efficient with new math."
author: "hasheddan"
captured_at: "2026-06-18T01:02:36Z"
capture_tool: "hn-digest"
hn_id: 48578295
score: 2
comments: 0
posted_at: "2026-06-17T23:09:37Z"
tags:
  - hacker-news
  - translated
---

# Tensordyne is making AI compute more efficient using logarithms

- HN: [48578295](https://news.ycombinator.com/item?id=48578295)
- Source: [www.zach.be](https://www.zach.be/p/tensordyne-is-making-ai-compute-more)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T23:09:37Z

## Translation

タイトル: Tensordyne は対数を使用して AI 計算をより効率的にしています
記事のタイトル: Tensordyne は、新しい数学により AI の計算をより効率的にしています。
説明: しかし、重要なのは計算効率なのでしょうか?

記事本文:
Tensordyne は、新しい数学を使用して AI の計算をより効率的にしています。
ザックの技術ブログ
Tensordyne は、対数を使用して AI の計算をより効率的にしています。
しかし、計算効率は重要なのでしょうか?
zach 2026年6月17日 5 シェア ほとんどの AI チップ企業は、メモリ階層を変更することで生成 AI ワークロードを高速化することに重点を置いています。 Groq は、Nvidia が依存している低速だがコスト効率の高い高帯域幅メモリ (HBM) ではなく、超高速だが高価な SRAM メモリを使用することで、高コストで超低遅延を実現しました。 Cerebras は、そのアイデアをさらに進めるためにウェハ スケールの処理を使用し、各ウェハで 44GB の SRAM を提供します。 d-Matrix はメモリ内での処理を利用して、メモリとの間のデータ移動を削減することでパフォーマンスを向上させます。 Taalas は読み取り専用メモリを使用しています。これは SRAM よりもさらに高速で安価ですが、変更することはできません。そのため、各チップは 1 つのモデルのみを実行できます。
Tensordyne は、AI チップに異なるアプローチをとっています。彼らは、対数に基づく異なる種類の数値体系を使用しているため、従来の浮動小数点演算を使用するチップよりも効率的に数値を乗算できます。しかし、より効率的な乗算器を実際の AI ワークロードでより優れたパフォーマンスに変換するのは困難です。今日のブログでは、Tensordyne の独自のテクノロジーと、市場における他のチップとの相対的な相対的な位置付けについて考察し、より効率的なコンピューティング エンジンは AI チップにとって実際に重要なのか?という質問に答えようとしています。
zach の技術ブログをお読みいただきありがとうございます。無料で購読して新しい投稿を受け取り、私の仕事をサポートしてください。
Tensordyne のアーキテクチャの中核となる技術的利点は、その対数形式にあります。値 A を単に A として保存するのではなく、実際には log 2 (A) を保存します。次に、 A に別の値 B を掛けたい場合、

対数積ルールを使用して、加算による乗算を実行できます。
log 2 (A x B) = log 2 (A) + log 2 (B)
このアイデアは新しいものではありませんが、商用 AI アクセラレータでこれを使用して成功した人はこれまでいません。大規模な行列の乗算を実行するには、多数の乗算器が必要ですが、同時に多くの加算器も必要になります。対数積ルールにより乗算は簡単になりますが、加算は非常に難しくなります。 Tensordyne は近似を使用してこれを回避します。
しかし、この近似だけでは十分ではありません。当初、Tensordyne は量子化を意識したトレーニングを使用して、この近似にもかかわらずモデルを実行していましたが、Tensordyne チップのためだけにモデルを再トレーニングするよう顧客を説得することは大きな課題になるでしょう。そこで彼らは、大規模なモデルに対して十分な精度の近似を行うための追加の独自のトリックを考案しました。
Tensordyne のソフトウェア スタックは、特定の数値形式で実行できるように PyTorch モデルを変換します。精度を大幅に低下させることなくこれを実行できるかどうかについては少し懐疑的ですが、たとえ実行できたとしても、この変換プロセスにより別の問題が発生します。最新の最適化された推論パイプラインは、PyTorch だけで書かれているわけではありません。 GPU とその固有の数値特性に合わせて最適化されたカスタム カーネルを備えています。顧客がそれをまったく新しい数値形式のチップに移植したい場合、最適化された GPU カーネルのパフォーマンス レベルに合わせるには多くの作業が必要になります。 Tensordyne は、AI エージェントが顧客のコードを Tensordyne のソフトウェアと数値形式に変換するのを支援すると主張しています…これは少し取り締まりのような答えのように思えます。基本的に、カーネルを記述するのは難しく、Tensordyne のチップの (非常に現実的な) 利点を享受するには、顧客が苦労してカーネルを Tensordyne の対数形式に移植する必要があるかもしれません。
このスペシャリ

また、zed 数値形式により、Tensordyne がトレーニング ワークロードをサポートすることが非常に困難になります。推論ワークロードは量子化などの数値変化に対して比較的寛容ですが、トレーニング ワークロードには数値安定性に関する大きな懸念があります。トレーニング パイプラインはハードウェアの癖に合わせて特別に設計されており、実行には数億ドルの費用がかかる場合があります。とにかく、主にネットワークとメモリ帯域幅によって支配されるトレーニングのようなワークロードのために、Tensordyne のチップ用の新しいトレーニング パイプラインを構築することは意味がありません。
メモリ帯域幅と言えば… Tensordyne が新しい番号システムを宣伝している一方で、ほとんどの AI チップは新しいメモリ階層を使用して画期的なパフォーマンスを達成することに重点を置いています。 Tensordyne は生成 AI ワークロードの間違ったボトルネックに焦点を当てているのでしょうか?
zach の技術ブログをお読みいただきありがとうございます。この投稿は公開されているので、自由に共有してください。
より効率的なコンピューティングは重要ですか?
ほとんどの AI チップがコンピューティングの仕組みを根本的に変えるのではなく、メモリ階層の革新に重点を置いているのには理由があります。多くの生成 AI ワークロードはメモリに依存しているからです。単一の GPU を使用して単一の LLM クエリを処理する場合、その GPU は、そのクエリを処理するために LLM 内のすべてのレイヤーのすべての重みをロードおよびアンロードする必要があります。 GPU の行列乗算の効率を大幅に向上させたとしても、メモリ帯域幅がボトルネックになります。
GPU はバッチ処理を通じてこの問題を回避します。一度に 1 つのクエリを処理するのではなく、複数のクエリのバッチを処理します。バッチに 32 個のクエリが含まれている場合、GPU はロードするレイヤーごとに 32 セットの行列乗算を実行します。これにより、メモリ帯域幅によるワークロードのボトルネックが軽減され、コンピューティング エンジンの効率が実際に重要になります。このc

「ルーフライン モデル」と呼ばれるものを使用してモデル化できます。バッチ サイズを増やすと、ワークロードの演算強度が増加します。つまり、転送されるメモリの 1 バイトあたりの演算量が増加します。演算強度が高くなるほど、コンピューティング エンジンの効率がより重要になります。演算強度が低いほど、メモリ帯域幅がより重要になります。
出典: Wikipedia ただし、バッチ処理には問題があり、LLM クエリが遅くなります。バッチ内の最初のクエリは、バッチ全体が処理されるまで待つ必要があります。 Cerebras や Groq など、バッチ サイズ 1 のオール SRAM アーキテクチャが非常に低いレイテンシーを実現するのはこのためです。全 SRAM アーキテクチャの支持者は、モデルが何時間も実行される可能性があるエージェント ワークロードにとってレイテンシーが非常に重要であると主張しています。トークンを 2 倍の速さで生成すると、エージェントのワークロードが 8 時間から 4 時間に短縮される可能性があります。これは意味のある違いです。
一方、バッチ処理により、ワークロードのコスト効率が向上します。 Nvidia が Groq の LPX チップを Vera Rubin NVL72 + Groq 3 LPX ハイブリッド クラスターに統合したとき、Groq のチップは高対話性、高コストのワークロードにのみ使用されました。ワットあたりのスループットの最大化 (コスト効率の代用) に重点を置いたワークロードの場合、バッチ処理を使用して動作する NVL72 ラックのみを使用しました。
出典: Nvidia ブログ Tensordyne のより効率的な乗算器がパフォーマンスに大きな違いをもたらすには、ルーフライン モデルの右側にあるコンピューティング制限領域で動作する必要があります。これはコスト重視のワークロードには最適な場所ですが、効率よりも速度に重点を置いた高コストのエージェント ワークロードでは、メモリの制約が多すぎて Tensordyne のハードウェアの恩恵を実際に受けられない可能性があります。ただし、対数計算 g には別の利点もあります。

ives: 物理的に小さい乗数。
zach の技術ブログをお読みいただきありがとうございます。無料で購読して新しい投稿を受け取り、私の仕事をサポートしてください。
コンピューティングが小さいほどメモリが増える
Tensordyne のコンピューティング エンジンがより効率的であれば、メモリに制約されたワークロードにとってはそれほど重要ではありません。しかし、コンピューティング エンジンが小さい場合は、メモリ階層に意味のある変更を加えることができます。乗算器を小さくすることでスペースを節約できるため、各ダイにより多くの SRAM メモリを搭載できるようになります。
出典: Tensordyne そこから、パフォーマンスの利点が明確になります。コンピューティング エンジンが小さくなるとオンチップ メモリが増加し、同じワークロードに必要なオフチップ メモリへのアクセス数が減り、チップの効率が向上します。
2024年にこれについてツイートしました！ただし、これにより Tensordyne のチップの効率がどれほど向上するかは不明です。確かに、コンピューティングに特化したバッチ処理されたワークロードで得られる利点ほど大きな利点はありません。実際の利点を見積もるには、チップ自体の技術仕様をさらに詳しく知る必要があります。
結局のところ、Tensordyne のチップは明らかなトレードオフをもたらします。モデルを推論のワークロードに制限する特殊な演算を備えた特殊なチップ上でモデルを実行する方法を見つけ出す必要があります。たとえ特別なコンパイラ、優れたサポート、AI エージェントの支援があったとしても、モデルを対数計算で適切に動作させるのは困難であり、パフォーマンスが低下する可能性があります。ただし、その代わりに、コンピューティングに依存した高レイテンシの推論ワークロードの効率が大幅に向上し、メモリに依存した低レイテンシの推論ワークロードの効率が若干向上します。市場が努力する価値があると考えるかどうか見てみましょう。いずれにせよ、新しい奇妙な数値形式が実際に世の中に浸透していくのを見るのは、とてもエキサイティングでクールなことだと思います。

市販のシリコン。
zach の技術ブログをお読みいただきありがとうございます。この投稿は公開されているので、自由に共有してください。
5 共有 この投稿に関するディスカッション コメント 再スタック トップ 最新のディスカッション 投稿はありません

## Original Extract

But is compute efficiency the thing that matters?

Tensordyne is making AI compute more efficient with new math.
zach's tech blog
Subscribe Sign in Tensordyne is making AI compute more efficient using logarithms.
But is compute efficiency the thing that matters?
zach Jun 17, 2026 5 Share Most AI chip companies are focusing on speeding up generative AI workloads by changing the memory hierarchy. Groq achieved ultra-low-latency at a high cost by using super-fast but expensive SRAM memory, rather than the slower but more cost-effective high-bandwidth memory (HBM) that Nvidia relies on. Cerebras is using wafer-scale processing to take that idea even further, with each wafer providing 44GB of SRAM. d-Matrix is leveraging processing-in-memory to improve performance by reducing data movement to and from memory. Taalas is using read-only memory, which is even faster and cheaper than SRAM, but can’t be changed -- so each chip can only run one model.
Tensordyne is approaching AI chips differently. They’re using a different kind of number system based on logarithms that allows them to multiply numbers more efficiently than chips using conventional floating point math. But translating more efficient multipliers into better performance on real AI workloads is hard. Today on the blog, we’re looking at Tensordyne’s unique technology, how it positions their chips relative to others in the market, and trying to answer the question: do more efficient compute engines actually matter for AI chips?
Thanks for reading zach's tech blog! Subscribe for free to receive new posts and support my work.
The core technical advantage of Tensordyne’s architecture comes from its logarithmic number format. Instead of storing a value A as just A , they actually store log 2 (A). Then, if they want to multiply A by another value, B , they can use the logarithm product rule to perform multiplication using addition:
log 2 (A x B) = log 2 (A) + log 2 (B)
This idea isn’t new, but nobody has successfully used it in a commercial AI accelerator before. Running large matrix multiplications requires a lot of multipliers, but it also requires a lot of adders. The log product rule makes multiplication easier, but it makes addition much harder. Tensordyne gets around this using an approximation :
But this approximation isn’t enough by itself. Originally, Tensordyne used quantization-aware training to get models running despite this approximation, but convincing customers to re-train models just for Tensordyne chips would be a huge challenge. So they figured out a set of additional proprietary tricks to make their approximation accurate enough for large models.
Tensordyne’s software stack converts PyTorch models to run with their specific number format. I’m a bit skeptical about whether they can do this without meaningful accuracy degradation, but even if they can, this conversion process raises another issue. Modern optimized inference pipelines aren’t just written in PyTorch; they have custom kernels optimized for GPUs and their specific numeric quirks . If a customer wants to port that to a chip with a whole new number format, it’ll take a lot of work to match the level of performance of optimized GPU kernels. Tensordyne claims that they’ll have AI agents help customers convert their code to Tensordyne’s software and numerical format … which seems like a bit of a cop-out answer. Fundamentally, writing kernels is hard, and to reap the (very real) benefits of Tensordyne’s chips, customers may need to bite the bullet and port kernels to Tensordyne’s logarithmic number format.
This specialized number format also makes it much, much harder for Tensordyne to support training workloads. Inference workloads are relatively forgiving to numerical changes like quantization, but training workloads have major numerical stability concerns . Training pipelines are specifically designed around hardware quirks, and can cost hundreds of millions of dollars to run. It just doesn’t make sense to build a new training pipeline for Tensordyne’s chips for a workload like training, which is primarily dominated by networking and memory bandwidth anyways.
Speaking of memory bandwidth… while Tensordyne is touting their new number system, most AI chips are focused on achieving breakthrough performance using new memory hierarchies. Is Tensordyne focusing on the wrong bottleneck for generative AI workloads?
Thanks for reading zach's tech blog! This post is public so feel free to share it.
Does more efficient compute matter?
There’s a reason why most AI chips have focused on innovating on the memory hierarchy, rather than fundamentally changing how the compute works: a lot of generative AI workloads are memory bound. If you want to use a single GPU to process a single LLM query, that GPU needs to load and unload all of the weights of every layer in the LLM to process that query. Even if you make the GPU much more efficient at matrix multiplication, you’re bottlenecked by the memory bandwidth.
GPUs get around this issue through batch processing. Instead of processing a single query at a time, they process a batch of multiple queries. If a batch has 32 queries in it, the GPU performs 32 sets of matrix multiplications for each layer it loads. This makes the workload less bottlenecked by memory bandwidth, and makes the efficiency of your compute engines actually matter. This can be modeled using something called the “ roofline model ”. As we increase the batch size, the arithmetic intensity of the workload increases -- in essence, it does more math per byte of memory transferred. The higher the arithmetic intensity, the more the efficiency of the compute engines matters. The lower the arithmetic intensity, the more memory bandwidth matters.
Source: Wikipedia However, there’s a problem with batching -- it slows LLM queries down. The first query in the batch needs to wait for the entire batch to be processed. This is why all-SRAM architectures with a batch size of 1, like Cerebras and Groq, offer extremely low latency. Advocates of all-SRAM architectures argue that latency is extremely important for agentic workloads, where models might be running for hours. If you generate tokens twice as fast, you might cut an agentic workload from 8 hours to 4 hours, which is a meaningful difference.
On the other hand, batching makes workloads more cost-effective. When Nvidia integrated Groq’s LPX chips into their Vera Rubin NVL72 + Groq 3 LPX hybrid clusters , they only used Groq’s chips for high-interactivity, high-cost workloads. For workloads focused on maximizing throughput-per-watt (which is a proxy for cost-efficiency), they just used NVL72 racks, operating using batched processing.
Source: Nvidia blog For Tensordyne’s more efficient multipliers to make a major performance difference, they need to be operating in the compute-bound regime on the right of the roofline model. This is a great place to be for cost-sensitive workloads, but high-cost agentic workloads that focus more on speed than efficiency may be too memory-bound to really benefit from Tensordyne’s hardware. However, there’s another advantage that logarithmic computation gives: physically smaller multipliers.
Thanks for reading zach's tech blog! Subscribe for free to receive new posts and support my work.
Smaller compute means more memory
If Tensordyne’s compute engine is more efficient, it’s not super important for memory-bound workloads. But if their compute engine is smaller, that allows them to make meaningful changes to the memory hierarchy, too. The space they’re saving with smaller multipliers allows them to fit more SRAM memory on each die.
Source: Tensordyne From there, the performance advantage becomes straightforward. Smaller compute engines result in more on-chip memory, which reduces the number of off-chip memory accesses needed for the same workload, and makes chips more efficient.
I tweeted about this in 2024! It’s unclear how much more efficient this makes Tensordyne’s chips, though. It’s certainly not as big of an advantage as the one they get for compute-bound batched workloads. To estimate the actual advantage, I’d need a lot more technical specifications of the chip itself.
Ultimately, Tensordyne’s chips are offering a clear tradeoff. You have to figure out how to get your models to run on their special chips with their special math, which will limit them to inference workloads. Getting models to work well with logarithmic math is likely going to be difficult and degrade performance, even if they have a special compiler, great support, and help from AI agents. But in return, you get much better efficiency on compute-bound, high-latency inference workloads and somewhat better efficiency on memory-bound, low-latency inference workloads. We’ll see if the market thinks it’s worth the effort. Either way, I think it’s super exciting and cool to see new, weird number formats actually make their way into commercially available silicon.
Thanks for reading zach's tech blog! This post is public so feel free to share it.
5 Share Discussion about this post Comments Restacks Top Latest Discussions No posts
