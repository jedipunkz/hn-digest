---
source: "https://systems.seas.harvard.edu/blog/burstiness-is-all-you-need/"
hn_url: "https://news.ycombinator.com/item?id=49126971"
title: "Bursty arrivals speed up LLM inference"
article_title: "When Bursty Traffic Makes LLM Inference Faster - Harvard Systems Group"
author: "ak33ra"
captured_at: "2026-07-31T19:08:27Z"
capture_tool: "hn-digest"
hn_id: 49126971
score: 2
comments: 1
posted_at: "2026-07-31T18:35:24Z"
tags:
  - hacker-news
  - translated
---

# Bursty arrivals speed up LLM inference

- HN: [49126971](https://news.ycombinator.com/item?id=49126971)
- Source: [systems.seas.harvard.edu](https://systems.seas.harvard.edu/blog/burstiness-is-all-you-need/)
- Score: 2
- Comments: 1
- Posted: 2026-07-31T18:35:24Z

## Translation

タイトル: バースト到着により LLM 推論が高速化
記事のタイトル: バースト トラフィックにより LLM 推論が高速化されるとき - ハーバード システム グループ
説明: バースト到着は通常、パフォーマンスの低下を引き起こし、運用システムでは頭痛の種とみなされます。このブログでは、バースト性が実際にパフォーマンスを向上させる現象を調査し、どのような条件がこの効果を生み出すかを明らかにします。

記事本文:
ハーバード システム グループの人々 セミナー 教育ブログ github バースト トラフィックが LLM 推論を高速化するとき
急増するワークロードは何十年もの間、コンピューター科学者に頭痛の種を与えてきましたが、最新の LLM サービスも例外ではありません。大きなバーストは瞬間的な負荷とレイテンシのスパイクを引き起こし、それらに対処する方法を探求する文献は数多くあります。
直観的にバースト性が問題となる理由は、平均到着率が同じであるにもかかわらず、リクエストが非常に近くに到着するか、または非常に離れて到着する可能性が高いためです。サーバーにとっては、短い時間枠内でフラッディングが発生し、その後にデッド期間が続いているように見えます。その結果、パフォーマンスは低下しますが、より通常のワークロードと比較して、長期実行時の使用率は同じになります。
そのため、定期的な vLLM ベンチマークを実行していたところ、到着のバースト性が高まるにつれて、遅延が増加するのではなく減少していることがわかり、これは非常に驚きでした。これは、標準的な直感に反するだけでなく、実稼働推論クラスターでリクエストのルーティングをどのように行うべきかに影響を与える可能性があります。
この短いブログでは、これらの結果に信頼性を与えるデータを調査しますが、プロジェクトをスライスしたパン以来最高のものから、スライスしたパン以来 2 番目に優れたものへと導く混乱を明らかにします。
結果: バースト性によりレイテンシが短縮される
このプロットでは、到着間隔時間の変動係数が増加する (よりバースト的になる) につれて、出力トークンあたりの時間の中央値 (TPOT) とエンドツーエンドの待ち時間 (E2EL) が減少することがわかります。これは、ShareGPT や BurstGPT などの合成ワークロードと現実的なワークロードの両方で観察されました。最初のトークンまでの時間 (TTFT) は構成に大きく依存しますが、到着の「クラスター化」により悪化することが予想されます。
図 1: 中央値/p90 TPOT および E2EL は、CV = 1 (ポアソン過程) と の間で約 50% 減少します。

BurstGPT では CV = 3.16 (非常にバーストな到着)。 TTFT は非単調に改善します。平均 TPOT (図示せず) は 70 ミリ秒から 30 ミリ秒まで 57% 向上します。バンドは、1,000 リクエストにおける 6 つのレプリケーションにわたる標準偏差を表します。
実際、バースト性は一般に分布の大部分に役立ちますが、場合によってはテールにわずかなダメージを与えます (図 2)。
図 2: BurstGPT での TPOT 推定 CDF は、バースト的な到着により分布全体が改善されることを示しています。
付録には、複数のモデル、GPU、データセットにわたるはるかに多くの結果が含まれており、この動作が一般化するという確信が高まりました。
大まかに言えば、バースト性が高いということは、到着間のギャップが非常に小さい値または非常に大きい値に重み付けされることを意味します。その後、大量の到着を取得し、その後、エンジンが中断することなく大量のデコードを実行する静かな期間が続く様子を想像できます。したがって、デコードのみを行う反復がさらに多くなることが予想されますが、より重要なのは、デコード トークンが大きなプレフィルから分離されることです。
この仮説は、図 3 に示す傾向に従うデータによって裏付けられています。つまり、バースト性が高くなると、デコードのみの反復の一部であるトークンの割合が増加します。
図 3: デコードのみの反復における出力トークンの割合は、バースト性が高い場合に 16.8 パーセント ポイント増加します。 (Qwen3.5-9B | GH200 | BurstGPT)
さらに、図 4 では、バースト性が高くなると、混合バッチが大きくなり、プレフィルがより多くなることがわかるため、プレフィルが一緒に圧縮されることが確認されています。
図 4: バースト性が高いと到着が集中し、より大規模でプレフィルが優勢な混合バッチが生成されます。 (Qwen3.5-9B | GH200 | BurstGPT)
図 1 ～ 4 では Qwen3.5-9B を使用しているため、その効果は GDN に関連しており、後述する Flash アテンションのバージョン管理には関連付けられていないことに注意してください。
プレフィルとデコードを混合すると効果がなくなるということは広く受け入れられています。

ses の干渉により、反復が遅くなります。これが PD disagg の背後にある 1 つの動機です。したがって、より排他的なデコード反復を作成すると、プリフィルとデコードの干渉が減り、平均デコード トークンが安くなり、エンジンが高速化することになります。
これらのいくつかの数値を使用すると、上記で示した 57% の TPOT 削減を大まかに回復することもできます。
低バースト性では、生成されたトークンの 71.2% で平均 26.05 ミリ秒の小規模なデコードのみの反復が発生し、28.8% で平均 156.98 ミリ秒の大規模な混合反復が発生しました。これにより、予想される GPU 転送間隔は 63.73 ミリ秒になります。バースト性が高い場合、対応する混合物は 14.36 ミリ秒で 88%、116.02 ミリ秒で 12% となり、26.59 ミリ秒になります。これは 58.3% の減少であり、観察された平均 TPOT 改善の 57% に近い値です。
さらに、3084 トークンの混合バッチが 2074 トークンの混合バッチよりも高速であることは直感に反するように思えるかもしれませんが、含まれるデコード トークンが少なく、プリフィル トークンよりも高価であることに注意してください。これは、図 A1 (付録) の線形フィットに示されています。
実際、その効果は 2 つあります。 TPOT の利点は次のとおりです。
デコード トークンを大規模な混合反復からより小さなデコード専用反復に移動します。
混合反復の構成を変更し、プレフィル トークンが安価になるため、平均して高速化します。
素晴らしい！つまり、全体として、バースト性の向上 → プリフィルとデコードの干渉の減少 → レイテンシの低下…ですよね?
研究を締めくくるために、カーネル レベルでプロファイリングを行い、この干渉の影響を検証することにしました。一般に、プレフィルはコンピューティングに依存し、デコードはメモリに依存するため、これら 2 つのワークロードによりリソースの競合が増加し、各トークンのコストが高くなる、と主張されます。言うまでもなく、プレフィルは一般にバッチ内のトークンの数を急増させ、反復を遅くして TPOT を不安定にします。
大丈夫。

この現象は、nsys (Nsight Systems) を数回実行することですぐに確認でき、干渉が確認されました。構成に応じて、混合バッチでは、純粋なデコード バッチと比較して、デコード トークンあたりのコストが最大 4 倍高くなりました (図 5)。
図 5: 各デコード トークンは、混合バッチでは反復ごとに最大 2 倍高価になります (Llama-3.1-8B、vLLM、A100)。注意点だけを考えても、混合バッチでは各デコード トークンが 4 倍高価になります。
結果は標準的な予想と一致しているため、ここで停止するのも不合理ではありません。ただし、これには非常に興奮しましたが、2 ～ 4 倍という差は額面通りに受け入れるには大きすぎます。
デコード専用バッチは CUDA グラフで実行され、起動のオーバーヘッドが除去されますが、それによるメリットはごくわずかです。プロファイリング結果を詳しく見ると、アテンション カーネルの名前と引数が混合バッチとデコード バッチで異なっていました。テンプレートのおかげで予想外ではありませんが、次の 2 つの点が際立っていました。
Gated DeltaNet (GDN) モデル (Qwen3.5-9B など) では、排他的デコード バッチは、混合バッチに使用されるチャンク デルタ ルールではなく、高速パス融合リカレント デルタ ルールを使用します。
Grouped Query Attendance (GQA) モデル (Llama-3.1-8B など) では、排他的デコード バッチは引数 Packed=true で実行され、混合デコード バッチは引数 Packed=false で実行されます。
デコード専用バッチではリソースの競合が少ないため、各トークンが本質的に安価であるというよりは、どちらの場合でも「干渉」は主に、デコード バッチの高速パスに関係するカーネル最適化のアーティファクトです。これは強く疑われていますが、GDN では十分に説明されておらず、GQA では確認されています。
まず、Gated DeltaNet のケースです。
簡単に言うと、GDN は、RNN と同様に、過去のトークンを固定サイズの状態に圧縮する線形アテンション アーキテクチャです。私たちの場合、私たちが気にしているのは、デコードがスーパーであるという事実です。

単純な直接状態更新: 出力されたトークンを見て、メモリを更新します。プレフィルも同様に簡単です。各トークンを順番に調べてメモリ状態を更新します。ただし、これには並列性がなく、大量の GPU リソースが無駄になるため、現在のアテンション カーネルはチャンキングと呼ばれるアプローチを使用してプロンプトのセクションを並列処理します。
この記事の執筆時点では、vLLM は混合バッチをチャンク カーネル経由でルーティングします。そこでは、ディメンションを均質化するためにデコード シーケンスも 64 個のトークンにパディングされ、冗長な作業が発生します。デコード専用バッチは、すべてのシーケンスで処理するトークンが 1 つだけであるという事実を利用する高速パスを通じてルーティングされます。
この体制では、混合バッチでは各デコード トークンのコストが人為的に高くなるのは当然です。簡単な修正は、デコード トークンとプレフィル トークンを分離して個別のカーネルを起動することで、起動のオーバーヘッドがわずかに発生しますが、冗長なデコード作業を削除することのようです。これを自分で試してみたところ、非常にわずかなメリットしか得られませんでしたが、これまでに観察された結果を説明するものではないため、GDN カーネルにどのような最適化が欠けているのかは現時点ではわかりません。ただし、デコード トークンあたりの秒数という点では、デコード専用バッチが混合バッチよりも本質的に効率的である可能性は低いと思います。
「標準」トランスフォーマーでは、各アテンション ヘッドが独自の K と V を計算し、大きなメモリ フットプリントを作成します。 GQA アーキテクチャでは、複数のアテンション ヘッドが少数の「KV ヘッド」を共有します。おそらく、このメモリ使用量の削減は、各転送パスで読み取るメモリが減るはずだと考えているかもしれませんが、それは正しいです。 GQA のアテンション カーネルは、共有するアテンション ヘッドに対して HBM から各 KV ヘッドを 1 回だけ読み取り、共有メモリに保存するように最適化できます。
これは紹介者です

「パッキング」として編集され、バッチが混合か排他かに関係なく実行できます。
ただし、混合バッチ パスでは、FA2 がこの最適化を実行しないため、デコード トークンごとのアテンション コストの観点から、Llama-3.1-8B で 4 倍の速度低下が観察されます。デコードは一般的にメモリに依存しており、Llama-3.1-8B には 32 個のアテンション ヘッドと 8 個の KV ヘッドがあり、メモリ トラフィックが 4 分の 1 に削減されるため、これは当然のことです。 FA3 と FA4 は両方ともまさにその問題に対処し、混合バッチと排他的バッチの両方でパッキングを実行します。
このアーティファクトは、混合バッチ中にメモリ帯域幅が圧迫されることを観察した、いくつかの出版された論文のプロファイリング結果と一致しています。通常、彼らは追加されたプレフィル トークンが原因でこれを無視しますが、KV ヘッドを冗長にロードすると、この観察はさらに悪化します。偶然にも、その構成では Ampere GPU で GQA モデルが使用されており、vLLM のデフォルトは FA2 です。また、それをテストした論文では、FA3 がデフォルトである Hopper GPU では競合の影響がほとんど消えます…うーん…
とにかく、さらにいくつかの測定と H200 へのアクセスにより、混合バッチにおける「干渉」の大部分が最適化されていない FA2 アテンション カーネルによるものであることが確認されました。これは、FA2 と FA3 を使用して H200 上の同じワークロードをベンチマークすることによってテストされました (図 6)。
図 6: バースティネスはフラッシュ アテンション 2 を使用すると TPOT を改善しますが、同じハードウェアとワークロードでフラッシュ アテンション 3 を使用すると傾向が逆転します。
FA3 を使用する場合、バースト性が高くても排他的デコード反復の割合は増加しますが、これによる TPOT への影響は最小限でした。バースト性の向上によって生じた TTFT ペナルティが利益を上回り、E2EL が悪化しました。
バースト性が実際に役立つのはどのような場合ですか?
専用バッチと混合バッチの効率の差が大きい場合、バースト性により大幅な改善がもたらされる可能性があります。

TPOT とエンドツーエンドの遅延の要素。これが当てはまるいくつかの体制は、FA2 を使用する場合の GQA モデルと線形注意モデルです。
エンジニアリングの観点から見ると、これらの結果は、このような状況でルーティングと負荷分散を実行する方法に影響を与えます。たとえば、「スティッキー」ラウンドロビンと JSQ が標準バージョンよりも大幅に優れていることを示しました (付録)。他の政策も恩恵を受ける可能性はあります。
また、LLM 推論関連のベンチマークを実行する場合、ハードウェアとアテンション バックエンドがベースラインの速度を超える影響を与えることも意味します。初めに、カーネルのバージョン、モデル アーキテクチャ、GPU の世代によってランタイムの動作が変わる可能性があるという事実を喜んでごまかしました。さまざまなモデルと、手に入る限り多くの GPU を試してみましたが、これらの実行にはたまたま相関関係がありました。 「モデル」と「ハードウェア」をそれぞれ 1 つの変数として扱うことは珍しいことではありませんが、その選択により、後で考慮する必要があるかもしれない多くの仮定が組み込まれます。
TL;DR、モデル x GPU スイープは思ったほど堅牢ではない可能性があることがわかりました。
要約すると、この楽しい調査は測定規律のレッスンも兼ねました。全体を通して、私はヤン教授とアレックスから、完全に説明できるようになるまですべての発見を疑うようにアドバイスされました。

[切り捨てられた]

## Original Extract

Bursty arrivals usually cause a drop in performance and are seen as headaches in production systems. In this blog, we investigate a phenomenon where burstiness actually improves performance, uncovering what conditions produce this effect.

Harvard Systems Group People Seminars Teaching Blog github When Bursty Traffic Makes LLM Inference Faster
Bursty workloads have been giving computer scientists headaches for decades, and modern LLM serving is no different. Big bursts create instantaneous load and latency spikes, and there’s a whole host of literature exploring how to deal with them.
The reason burstiness hurts, intuitively, is because we have the same mean arrival rate but requests are more likely to arrive very close together, or very far apart. To the server, it looks like it’s getting flooded within a short window, followed by a dead period. The result is worse performance but the same long run utilization compared to a more regular workload.
So, when I was doing some routine vLLM benchmarking and found that as the arrivals got more bursty, latency decreased rather than increased, it was pretty surprising! Beyond contradicting standard intuition, this would have implications for how request routing should be done in production inference clusters.
In this short blog, we’ll explore the data that gives confidence to those results, but surfaces a confound that takes the project from the best thing since sliced bread to… the second best thing since sliced bread.
Results: burstiness lowers latency
In this plot, we see that as the coefficient of variation of the interarrival times increases (more bursty), the median time per output token (TPOT) and end-to-end latency (E2EL) decrease. This was observed on both synthetic and realistic workloads, like ShareGPT and BurstGPT. Time to first token (TTFT) is much more config dependent, though we’d expect it to worsen due to the “clustering” of arrivals.
Figure 1: Median/p90 TPOT and E2EL decrease ~50% between CV = 1 (Poisson process) and CV = 3.16 (very bursty arrivals) on BurstGPT. TTFT improves non-monotonically. Mean TPOT (not shown) improves 57% from 70 ms to 30 ms. Bands represent standard deviation across 6 replications at 1k requests.
In fact, burstiness generally helps most of the distribution, sometimes slightly hurting the tail (Fig. 2).
Figure 2: TPOT estimated CDF on BurstGPT shows that burstier arrivals improve the whole distribution.
The appendix includes far more results, spanning multiple models, GPUs, and datasets, increasing our confidence that this behavior generalizes.
Loosely, higher burstiness means the gaps between arrivals are more weighted to very small or very large values. We can then picture getting a clump of arrivals, followed by a quiet period where the engine performs a bunch of decode without interruption. It follows that we should expect more iterations which are exclusively decode, but more importantly, that decode tokens get separated from large prefills.
This hypothesis is supported by the data, which follows the trend shown in Fig. 3: higher burstiness increases the fraction of tokens that are part of decode-only iterations.
Figure 3: The fraction of output tokens in a decode-only iteration increases by 16.8 percentage points at high burstiness. (Qwen3.5-9B | GH200 | BurstGPT)
Additionally, Fig. 4 confirms that prefills get compressed together, since we see that at high burstiness, mixed batches get larger and more prefill-heavy.
Figure 4: Higher burstiness concentrates arrivals, leading to larger and prefill-dominant mixed batches. (Qwen3.5-9B | GH200 | BurstGPT)
Note that figures 1 through 4 use Qwen3.5-9B, so the effects there are GDN-related and not tied to the Flash Attention versioning discussed later.
It’s widely accepted that mixing prefill and decode causes interference, hence slower iterations. This is one motivator behind PD disagg. It follows that if we create more exclusive decode iterations, we reduce the prefill-decode interference, making the average decode token cheaper and speeding up the engine.
Using these few numbers, we can also loosely recover the 57% TPOT reduction presented above.
At low burstiness, 71.2% of generated tokens experience small decode-only iterations averaging 26.05 ms, while 28.8% experience large mixed iterations averaging 156.98 ms. This gives an expected GPU-forward interval of 63.73 ms. At high burstiness, the corresponding mixture is 88% at 14.36 ms and 12% at 116.02 ms, giving 26.59 ms—a 58.3% reduction, close to the observed 57% mean TPOT improvement.
Additionally, while it may seem counter-intuitive that the 3084 token mixed batch is faster than the 2074 token one, note that it contains fewer decode tokens, which are more expensive than prefill tokens. This is shown in the linear fit of Fig. A1 (appendix).
In fact, the effect is twofold. We get TPOT benefits from:
Moving decode tokens out of large mixed iterations into smaller decode-only iterations, and
Changing the composition of mixed iterations, making them faster on average due to prefill tokens being cheaper.
Great! So overall, more burstiness → less prefill-decode interference → lower latency… Right?
To round out the study, we decided to profile at the kernel level and validate the effects of this interference. Commonly, people claim that because prefill is compute-bound and decode is memory-bound, these two workloads create more resource contention, making each token more expensive. Not to mention, prefill generally spikes the number of tokens in the batch, making the iteration slower and destabilizing TPOT.
Fine. The phenomenon was quick to verify with a few nsys (Nsight Systems) runs, and the interference was confirmed: the cost per decode token was up to 4x higher in mixed batches compared to pure decode batches (Fig. 5), depending on the config.
Figure 5: Each decode token is ~2x more expensive in a mixed batch, per iteration (Llama-3.1-8B, vLLM, A100). In attention alone, each decode token is 4x more expensive in a mixed batch.
It wouldn’t be unreasonable to stop here, since the results line up with standard expectations. However, while I was super excited about this, 2-4x is too large a difference to accept at face value.
Decode-only batches do run with CUDA Graphs, removing launch overhead, but the gains from that are negligible. If we look closer at the profiling results, the attention kernel’s name and arguments differed between mixed and decode batches. Not unexpected, due to templating, but two things stood out:
In Gated DeltaNet (GDN) models (like Qwen3.5-9B), exclusive decode batches use a fast path fused recurrent delta rule, rather than the chunked delta rule used for mixed batches.
In Grouped Query Attention (GQA) models (like Llama-3.1-8B), exclusive decode batches run with the argument packed=true , while mixed ones run with packed=false .
Rather than each token being inherently cheaper in a decode-only batch due to less resource contention, the “interference” in both cases is largely a kernel optimization artifact involving fast paths for decode batches. This is strongly suspected but not fully explained for GDN and confirmed for GQA.
First, the Gated DeltaNet case.
To be brief, GDN is a linear attention architecture which compresses past tokens into a fixed-size state, similar to RNNs. In our case, what we care about is the fact that decode is a super straightforward direct state update: look at the outputted token and update our memory. Prefill could be similarly easy: look at each token sequentially and update the memory state. However, this has no parallelism and wastes a lot of GPU resources, so current attention kernels use an approach called chunking to process sections of the prompt in parallel.
At the time of writing, vLLM routes mixed batches through the chunked kernel, where even decode sequences are padded to 64 tokens to homogenize the dimensions, creating redundant work. Decode-only batches are routed through a fast path which takes advantage of the fact that all sequences have only 1 token to process.
Under this regime, it makes sense that each decode token is artificially more expensive in a mixed batch. It seems like an easy fix would be to separate the decode and prefill tokens and launch separate kernels, inducing slight launch overhead but removing the redundant decode work. Trying this myself yielded extremely modest benefits, and doesn’t explain the results we’ve observed so far, so I’m currently not sure what optimizations the GDN kernel is missing. However, I find it unlikely that a decode-only batch is inherently more efficient than a mixed batch in terms of sec per decode token .
In a “standard” transformer, each attention head computes its own K and V, creating a large memory footprint. In the GQA architecture, we have multiple attention heads share a smaller number of “KV heads”. You’re now probably thinking this reduced memory footprint should mean less memory to read on each forward pass, and you’re correct! GQA’s attention kernel can be optimized to read each KV head from HBM only once for the attention heads which share it, storing it in shared memory.
This is referred to as “packing”, and can be done regardless of whether a batch is mixed or exclusive.
But, in the mixed batch path, FA2 does not perform this optimization, leading to the observed 4x slowdown on Llama-3.1-8B, in terms of attention cost per decode token. This should make sense since decode is generally memory-bound, and Llama-3.1-8B has 32 attention heads and 8 KV heads, leading to 4x reduced memory traffic. FA3 and FA4 both address that very issue and perform packing in both mixed and exclusive batches.
This artifact is consistent with the profiling results of several published papers, which observe strained memory bandwidth during mixed batches. They generally brush this off as being caused by added prefill tokens, but redundantly loading KV heads would exacerbate this observation. Coincidentally, their configurations use a GQA model on an Ampere GPU, where vLLM defaults to FA2. Also, in papers that test it, the contention effects mostly disappear on a Hopper GPU, where FA3 is the default… hmm…
Anyway, some more measurements and access to an H200 confirmed that a majority of the “interference” in mixed batches was due to the underoptimized FA2 attention kernel, tested by benchmarking the same workload on H200 with FA2 and FA3 (Fig. 6).
Figure 6: Burstiness improves TPOT using flash attention 2, but the trend reverses with flash attention 3 on the same hardware and workload.
When using FA3, higher burstiness still increased the fraction of exclusive decode iterations, but this had a minimal effect on TPOT. Any gains were outweighed by the TTFT penalties incurred from higher burstiness, worsening E2EL.
When does burstiness actually help?
When the difference in efficiency between exclusive and mixed batches is high, burstiness can produce significant improvements in TPOT and end-to-end latency. Some regimes where this holds are GQA models when using FA2, as well as linear attention models.
From an engineering standpoint, these results have implications for how one performs routing and load balancing in these situations. For example, we showed that “sticky” round robin and JSQ strongly outperform the standard versions (appendix). It’s possible that other policies would benefit.
It also implies that when doing LLM inference-related benchmarking, the hardware and attention backend have effects beyond how fast your baseline is. In the beginning, I happily glossed over the fact that kernel versions, model architectures, and GPU generations can change runtime behavior. Though I did experiment with different models and as many GPUs as I could get my hands on, these runs happened to be correlated. Treating “model” and “hardware” each as a single variable isn’t uncommon, but that choice bakes in a bunch of assumptions you might need to account for later.
TL;DR, it turns out your model x GPU sweep might not be as robust as you thought.
To summarize, this fun investigation doubled as a lesson in measurement discipline. Throughout, I was advised by Prof. Yang and Alex to question every finding until I could fully explain

[truncated]
