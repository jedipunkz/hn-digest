---
source: "https://www.categoryvc.com/writing/where-the-ai-hardware-market-is"
hn_url: "https://news.ycombinator.com/item?id=48341537"
title: "AI Hardware"
article_title: "On AI Hardware"
author: "gmays"
captured_at: "2026-05-31T00:24:01Z"
capture_tool: "hn-digest"
hn_id: 48341537
score: 1
comments: 0
posted_at: "2026-05-30T23:18:28Z"
tags:
  - hacker-news
  - translated
---

# AI Hardware

- HN: [48341537](https://news.ycombinator.com/item?id=48341537)
- Source: [www.categoryvc.com](https://www.categoryvc.com/writing/where-the-ai-hardware-market-is)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T23:18:28Z

## Translation

タイトル: AI ハードウェア
記事タイトル: AI ハードウェアについて
説明: 推論の物理学とその周りに形成されるインフラストラクチャ スタック。

記事本文:
AI ハードウェアについて チーム企業の視点 © 2026 カテゴリ Ventures SF ← AI ハードウェアに関するすべての視点
最新の GPU 内の演算ユニットは、LLM 推論の待機に多くの時間を費やします。
H100 は、精度とスパース性を考慮するかどうかに応じて、2 ～ 4 PFLOP/秒程度の膨大なテンソル スループットを備えています。ただし、自己回帰デコード中のボトルネックは通常、乗算ではありません。低度から中度のバッチ デコードでは、新しいトークンごとにシステムがアクティブなモデルの重みをストリーミングし、蓄積された状態に対処する必要がありますが、GPU の HBM 帯域幅 (H100 SXM で約 3.35 TB/秒) は有限です。ワークロードがメモリから取得したバイトごとに何百もの有用な操作を実行しない限り、Tensor コアを完全に占有したままにすることはできません。¹
¹ 正確な比率は精度と会計規則によって異なりますが、重要な点は安定しています。最新の GPU は、HBM から取得できる各バイトに対して数百ユニットの計算が利用可能です。自己回帰デコードはそのしきい値をはるかに下回ります。
そして、この問題は GPU の世代を超えて続いています。テンソルのピーク スループットは、デコード フェーズの外部メモリ帯域幅よりも速く増加しているため、動作強度のギャップは埋まっていません。どちらかというと幅が広がりました。
それが、現在の興味深い AI ハードウェア市場のほとんどの根底にある基本的な事実です。この分野の企業にとって有益な質問は、メモリ問題のどの部分を攻撃しているのか、そしてそれを解決することで NVIDIA との真っ向からの競争を回避できるのかということです。チップ内、サービングエンジン内、キャッシュ階層内、または物理パッケージとラック内など、その制約を攻撃する場所を中心に市場が組織され始めています。
Groq ↗ は、HBM をまったく使用せずにチップを構築し、それをオンチップ SRAM (直接構築されたメモリ) に置き換えました。

DRAM よりもはるかに高速ですが、密度がはるかに低いため、単位面積あたりの容量がはるかに少なくなります)。実行モデル全体は決定的でした。コンパイラは各サイクルを静的にスケジュールし、キャッシュや動的スケジューリングは行いません。大きなモデルを保持するにはさらに多くのチップが必要ですが、HBM を待つ必要はありません。 NVIDIA の Groq 契約は競争の枠組みを変えました。SRAM スタイルの推論はもはや単なるスタートアップの賭けではなく、NVIDIA 独自のプラットフォーム ロードマップの一部となっています。その結果、Groq 3 LPX 推論アクセラレータが、メイン GPU とともに NVIDIA の Vera Rubin ↗ プラットフォームに統合されました。
Cerebras ↗ は、44 GB のオンチップ SRAM と 21 PB/s の内部メモリ帯域幅を備え、H100 ダイの 50 倍を超える面積のシリコン ウェーハ全体にまたがる単一チップを構築しました。同じ理論、異なるメカニズム。彼らは2026年5月に株式を公開した。
MatX ↗ は、トランスフォーマー推論のアクセス パターン用に設計されたスクラッチパッド メモリ (ハードウェア管理のキャッシュではなく、ソフトウェア管理の高速メモリ) を中心に構築されています。 d-Matrix ↗ はメモリ内計算を実行し、演算をメモリ配列自体に移動します。
これらはどれも、データが存在する場所と演算が行われる場所との間のギャップを縮小または排除するための異なる試みです。
ハードウェアが不完全であっても、ソフトウェアがボトルネックを回避することができます。ここから、有用な裏返しのルーフライン ヒューリスティックが得られます。大まかなルーフライン ヒューリスティックでは、デコードのおおよそのバッチ サイズは 300 × ( N to t a l / N a c t i v e ) 300 \times (N_{total} / N_{active}) 300 × ( N to t a l / N a c t i v e ) のようになります。ここで、300 は現在のメモリ帯域幅に対するピーク コンピューティングのおおよその比率を反映しています。ハードウェア。高密度モデルの場合、約 300 トークンになります。 DeepSeek ↗ スタイルの MoE の場合、合計 700B のうち約 37B がパラメータになります

rs はトークンごとにアクティブになり、おおよそのバッチは 6,000 近くになります。実際には、実際の数値は、精度、KV プレッシャー、並列処理戦略、レイテンシの目標、およびサービススタックに応じて変化しますが、重要なのは方向です。それ以下では帯域幅が制限されます。それを超えると、コンピューティングはますます制限されます。大まかなハードウェア レベルでは、システムは数十ミリ秒程度のサイクルで実行され、これは HBM を介したストリーミングにかかる​​時間とほぼ同じであり、各サイクルでアクティブなシーケンスごとに 1 つの新しいトークンが生成されると考えることができます。
このことから生じるスケジューリングの問題は、レイテンシの目標を達成しながらスループットを最大化するために、リクエストをこれらのサイクルにどのように詰め込むかということです。通常、プレフィルは計算負荷が高くなります。デコードは通常、メモリ帯域幅に依存します。彼らはハードウェアとは異なるものを求めており、それらを効率的に組み合わせるのは困難です。
RadixArk ↗ と Inferact ↗ は両方とも最近これに着手しました。 RadixArk は SGLang ↗ チームであり、推論用の SGLang と RL トレーニング用の Miles ↗ 上に商用インフラストラクチャを構築しています。 Inferact は、vLLM で同じことを行っている vLLM ↗ チームです。両方のバークレー プロジェクトは広く展開されており、現在は企業として互いに競争しています。
なぜこれが興味深いのでしょうか?問題が複雑化するからです。新しいモデル アーキテクチャ、新しいハードウェア プラットフォームごとに最適化面が変化します。ハードウェアの異種混合が進み、ワークロードがより多様になるにつれて、問題はますます困難になるため、ルーフラインを意識したスケジューリングに関する深い専門知識を構築したチームは、その利点を蓄積し、成長していきます。それほど明白ではない問題は、この市場が 1 つの主要なエンジンを中心に統合されるのか、それとも内部フォーク、マネージド サービス、およびオープンソースのデフォルトにまたがって断片化されるのかということです。 RadixArk はマイルを通じて差別化を図っており、RL トレーニングを提供しています。

第二の表面積。しかし、真っ向からの推理は熾烈を極めるだろう。
メモリの問題は、体重の読み取りよりもさらに深刻です。推論中に、コンテキスト内の以前のすべてのトークンについて保存されたアテンション状態である KV キャッシュも読み取ります。重みは固定されています。 KV キャッシュは、コンテキストの長さとバッチ サイズに応じて直線的に増加します。
大規模な変圧器モデルの場合、すべてのレイヤー、KV ヘッド、ヘッドの寸法、精度を考慮すると、KV キャッシュはトークンごとに数百 KB になる可能性があります。約 100B のアクティブ パラメーターとトークンあたりの合計 KV フットプリントが約 500 KB のモデルの場合、KV キャッシュは約 200,000 トークンで 100 GB に達し、重量フットプリントと一致します。それ以下では、主に読み取られた重量に対して料金を支払います。それ以上では KV キャッシュが支配的となり、コストは直線的に上昇します。これは、ロングコンテキストの価格設定が 200,000 トークンなどのしきい値を超えて増加することがよくある技術的な理由として考えられますが、価格設定には製品のセグメント化と容量計画も反映されています。
Reiner Pope の講義 ↗ には、メモリ階層がこれにどのようにマッピングされるかについての素晴らしい観察があります。各層の「ドレイン時間」（容量を帯域幅で割ったもの）によって、その自然な保存期間が設定されます。 HBM は約 20 ミリ秒で排出されます。 DDR は数秒で消耗します。約1分以内に点滅します。 1時間ほどでディスクが回転します。アイドル状態の会話用に KV キャッシュを保持している場合は、HBM を占有するのではなく、この階層の下に移行する必要があります。
TensorMesh ↗ はこれを中心に構築されています。彼らのオープンソース プロジェクト LMCache ↗ は、GPU、CPU RAM、NVMe、S3 全体に KV キャッシュを保存するため、再利用可能な KV 状態は最初から再計算されるのではなく、引き戻されます。私の推測では、これは大規模なスタンドアロン プラットフォームのままであるよりも、サービス提供エンジンに吸収される可能性が高いですが、その間、LMCache はエンジンがまだ所有したくないと思われる作業を実行しています。

直接。
チップとサーバーのレベルでメモリとスケジューリングを十分に強化すると、制約はパッケージ、ボード、そして最終的にはラックにまで広がります。
CoWoS (TSMC の高度なパッケージング技術、GPU ダイをインターポーザー上の HBM スタックに物理的に接続するもの) は、真の供給ボトルネックとなっています。また、72 GPU システムから 500 以上の GPU スケールアップ ドメインへの移行は、単なるネットワークの問題ではありません。コネクタの密度、ケーブルの曲げ半径、電力供給、液体冷却、および HBM 接続の歩留まりに影響します。問題は機械工学と材料科学であり、これらはチップ企業が得意とする分野から実際の距離を生み出しています。
同様に考慮に値するのは、すべての高度なファブが依存するリソグラフィー マシンを作成する ASML です。これらは、どのチップ設計者とも競合することなく、すべてのチップ設計者を可能にします。パッケージングや相互接続のボトルネックを解決する企業は、ASML のような市場での立場に立つ可能性がありますが、独占的ではない可能性があります。ここでの価値は、EUV リソグラフィのように単一のシステムオブシステムに集中するのではなく、TSMC、基板ベンダー、光相互接続サプライヤー、ODM に分割される可能性があります。
市場はメモリの問題の山になりつつあります。チップは重みを計算に近づけようとします。サービス エンジンは、帯域幅に合わせてバッチ処理とスケジュールを試みます。 KV キャッシュ システムは、アテンション状態を HBM 内に残したままにするのではなく、階層を通じて移動しようとします。パッケージ化と相互接続により、ラックがより 1 台のマシンのように動作するように努めます。
TensorMesh が構築されているオープンソース プロジェクトである LMCache ↗ は、同社がスタンドアロンで販売しようとしているにもかかわらず、すでに主要なサービス エンジンとランタイムに取り込まれています。おそらくそれがテンプレです。便利なプリミティブが統合されると、スタンドアロンのラッパーは別のものを見つけなければなりません

であること。
そして、Etched ↗ はまったく異なる種類の賭けをしており、トランスのアーキテクチャが固定機能のシリコンに焼き付くほど十分に確立されていることに賭けています。それが正しい場合、すべてのトランジスタは正確に正しい動作をします。それが間違っている場合、汎用のアクセル装置よりも操作の余地がはるかに少なくなります。そして、Groq 契約により NVIDIA の Rubin プラットフォームに独自の SRAM ベースの推論アクセラレータが組み込まれたことで、Etched はもはやレガシー GPU と競合するだけではなくなりました。
主な対抗要因はアルゴリズムの進歩です。ハードウェアはゆっくりと変化しますが、サービス提供技術やモデル アーキテクチャは急速に変化します。投機的デコード、KV キャッシュ圧縮、スパース性、蒸留、およびネイティブの低ビット モデルはすべて、有用なトークンごとに必要なメモリ移動の量を削減します。また、これらの技術がきれいに拡張できれば、一部の特殊なハードウェア アプローチの根拠を弱めることができます。ただし、根底にある制約は削除されません。それらはその形を変えます。精度が低いとパラメータあたりのバイト数が減りますが、ネイティブの低ビット演算が必要になります。投機的デコードでは重みの読み取りが償却されますが、ドラフト モデルの受け入れ率に依存し、KV キャッシュの圧力はそのまま残ります。スパース性によりアクティブなパラメータは減少しますが、ルーティングと相互接続の複雑さは増加します。ターゲットは移動しているため、耐久性のあるハードウェア企業は、ボトルネックが変化しても有用であり続けるアーキテクチャを必要としています。
各層で耐久性のある企業が構築されると私は信じています。それほど明確ではないのは、これらの企業が独立性を維持しているかどうかです。 AI インフラストラクチャで最も有用なプリミティブは、サービス エンジン、クラウド、ラボ、または NVIDIA 自体によってすぐに吸収される傾向があります。難しいのは、ボトルネックが重要であることを証明することではなく、システムの他の場所に内部化できない制御ポイントを所有することです。

スタック。

## Original Extract

The physics of inference and the infrastructure stack forming around it.

On AI Hardware Team Companies Perspectives © 2026 Category Ventures SF ← All perspectives On AI Hardware
The arithmetic units inside a modern GPU spend much of LLM inference waiting.
An H100 has enormous tensor throughput, on the order of 2 to 4 PFLOP/s depending on precision and whether you count sparsity. But during autoregressive decode, the bottleneck is usually not multiplication. In low-to-moderate batch decode, each new token effectively requires the system to stream through the active model weights and attend over accumulated state, while the GPU’s HBM bandwidth (about 3.35 TB/s on an H100 SXM) is finite. Unless a workload performs hundreds of useful operations per byte fetched from memory, the tensor cores can’t stay fully occupied.¹
¹ The exact ratio depends on precision and accounting conventions, but the important point is stable: modern GPUs have hundreds of units of compute available for every byte they can pull from HBM. Autoregressive decode is far below that threshold.
And this problem has persisted across GPU generations. Peak tensor throughput has grown faster than external memory bandwidth for the decode phase, so the operational-intensity gap hasn’t closed. If anything it has widened.
That is the basic fact underneath most of the interesting AI hardware market right now. The useful question for any company in this space is which part of the memory problem they’re attacking, and whether solving it lets them avoid competing head-on with NVIDIA. The market is starting to organize around where you attack that constraint: inside the chip, inside the serving engine, inside the cache hierarchy, or inside the physical package and rack.
Groq ↗ built a chip with no HBM at all, replacing it with on-chip SRAM (memory built directly into the processor die, much faster than DRAM but also much less dense, so you get far less capacity per unit area). The whole execution model was deterministic. The compiler scheduled every cycle statically, no caches, no dynamic scheduling. You need many more chips to hold a large model, but you never wait for HBM. NVIDIA’s Groq deal changed the competitive frame: SRAM-style inference is no longer just a startup bet, it is now part of NVIDIA’s own platform roadmap. The result is the Groq 3 LPX inference accelerator, integrated into NVIDIA’s Vera Rubin ↗ platform alongside the main GPU.
Cerebras ↗ built a single chip that spans an entire silicon wafer, over 50x the area of an H100 die, with 44 GB of on-chip SRAM and 21 PB/s of internal memory bandwidth. Same thesis, different mechanism. They went public in May 2026.
MatX ↗ is building around scratchpad memories (software-managed fast memory, as opposed to hardware-managed caches) designed for the access patterns of transformer inference. d-Matrix ↗ is doing in-memory compute, moving the arithmetic into the memory array itself.
Every one of these is a different attempt to shrink or eliminate the gap between where data lives and where arithmetic happens.
Even with imperfect hardware, software can route around the bottleneck. A useful back-of-the-envelope roofline heuristic falls out of this. A crude roofline heuristic says the approximate batch size for decode scales like 300 × ( N t o t a l / N a c t i v e ) 300 \times (N_{total} / N_{active}) 300 × ( N t o t a l ​ / N a c t i v e ​ ) , where the 300 reflects the approximate ratio of peak compute to memory bandwidth on current hardware. For a dense model that’s about 300 tokens. For a DeepSeek ↗ -style MoE where roughly 37B of 700B total parameters are active per token, the approximate batch is closer to 6,000. In practice, the true number moves around with precision, KV pressure, parallelism strategy, latency targets, and the serving stack, but the direction is what matters. Below that point, you are bandwidth-limited. Above it, you are increasingly compute-limited. At a rough hardware level, you can think of the system as running in cycles on the order of tens of milliseconds, roughly the time it takes to stream through HBM, with each cycle producing one new token per active sequence.
The scheduling problem that falls out of this is how to pack requests into these cycles to maximize throughput while hitting latency targets. Prefill is usually compute-heavy. Decode is usually memory-bandwidth-sensitive. They want different things from the hardware, and mixing them efficiently is hard.
RadixArk ↗ and Inferact ↗ have both recently launched into this. RadixArk is the SGLang ↗ team, building commercial infrastructure on top of SGLang for inference and Miles ↗ for RL training. Inferact is the vLLM ↗ team doing the same with vLLM. Both Berkeley projects, both widely deployed, now racing each other as companies.
Why is this interesting? Because the problem compounds. Every new model architecture, every new hardware platform changes the optimization surface. A team that builds deep expertise in roofline-aware scheduling accumulates an advantage that grows, because the problem keeps getting harder as hardware gets more heterogeneous and workloads get more varied. The less obvious question is whether this market consolidates around one dominant engine, or fragments across internal forks, managed services, and open-source defaults. RadixArk has some differentiation through Miles, which gives them RL training as a second surface area. But the head-to-head on inference will be intense.
The memory problem goes deeper than weight reads. During inference, you also read the KV cache, the stored attention state for every previous token in the context. The weights are fixed. The KV cache grows linearly with context length and batch size.
For a large transformer model, the KV cache can easily be hundreds of KB per token once you account for all layers, KV heads, head dimensions, and precision. For a model with roughly 100B active parameters and a total KV footprint of around 500 KB per token, the KV cache reaches 100 GB at about 200,000 tokens, matching the weight footprint. Below that, you're mostly paying for the weight read. Above it, KV cache dominates and costs climb linearly. This is a plausible technical reason long-context pricing often steps up past thresholds like 200K tokens, though pricing also reflects product segmentation and capacity planning.
There's a nice observation from Reiner Pope's lecture ↗ about how the memory hierarchy maps onto this. The "drain time" of each tier (capacity divided by bandwidth) sets its natural retention horizon. HBM drains in about 20 milliseconds. DDR drains in seconds. Flash in about a minute. Spinning disk in about an hour. If you're holding KV cache for an idle conversation, it should migrate down this hierarchy rather than hogging HBM.
TensorMesh ↗ is building around this. Their open-source project LMCache ↗ stores KV caches across GPU, CPU RAM, NVMe, and S3, so reusable KV state gets pulled back in rather than recomputed from scratch. My guess is that this is more likely to be absorbed into serving engines than remain a large standalone platform, but in the meantime LMCache is doing work the engines don't yet seem to want to own directly.
Once you push memory and scheduling hard enough at the chip and server level, the constraint moves outward to the package, the board, and eventually the rack.
CoWoS (TSMC's advanced packaging technology, the thing that physically connects GPU dies to HBM stacks on an interposer) has been a genuine supply bottleneck. And the jump from 72-GPU systems to 500+ GPU scale-up domains is not just a networking problem. It runs into connector density, cable bend radius, power delivery, liquid cooling, and HBM attach yield. The problems are mechanical engineering and materials science, which creates real distance from what the chip companies are good at.
The parallel worth thinking about is ASML, which makes the lithography machines every advanced fab depends on. They enable all chip designers without competing with any of them. A company that cracks a packaging or interconnect bottleneck could be in an ASML-like market position, though likely less monopolistic: the value here may be split across TSMC, substrate vendors, optical interconnect suppliers, and ODMs rather than concentrating in a single system-of-systems the way EUV lithography has.
The market is becoming a stack of memory problems. Chips try to keep weights closer to compute. Serving engines try to batch and schedule around bandwidth. KV cache systems try to move attention state through a hierarchy instead of leaving it stranded in HBM. Packaging and interconnect try to make the rack behave more like one machine.
LMCache ↗ , the open-source project TensorMesh is built on, is already being pulled into the major serving engines and runtimes even as the company tries to sell it standalone. That's probably the template. The useful primitive gets integrated, and the standalone wrapper has to find something else to be.
And then there's Etched ↗ , which is making a different kind of bet entirely, wagering that the transformer architecture is settled enough to burn into fixed-function silicon. If it's right, every transistor does exactly the right thing. If it's wrong, there's much less room to maneuver than there is in a general-purpose accelerator. And now that NVIDIA's Rubin platform includes its own SRAM-based inference accelerator via the Groq deal, Etched isn't just competing against legacy GPUs anymore.
The major counterweight is algorithmic progress: hardware changes slowly, while serving techniques and model architectures can move quickly. Speculative decoding, KV-cache compression, sparsity, distillation, and native low-bit models all reduce the amount of memory movement required per useful token, and if those techniques scale cleanly, they weaken the case for some exotic hardware approaches. But they do not remove the underlying constraint. They change its shape. Lower precision reduces bytes per parameter, but creates demand for native low-bit arithmetic. Speculative decoding amortizes weight reads, but depends on draft-model acceptance rates and leaves KV-cache pressure intact. Sparsity reduces active parameters, but increases routing and interconnect complexity. The target is moving, which is why durable hardware companies need architectures that remain useful as the bottleneck shifts.
I believe durable companies will be built at each layer. What's less clear is whether those companies remain independent. The most useful primitives in AI infrastructure tend to get absorbed quickly by serving engines, clouds, labs, or NVIDIA itself. The hard part isn't proving that the bottleneck matters, but owning a control point that can't be internalized somewhere else in the stack.
