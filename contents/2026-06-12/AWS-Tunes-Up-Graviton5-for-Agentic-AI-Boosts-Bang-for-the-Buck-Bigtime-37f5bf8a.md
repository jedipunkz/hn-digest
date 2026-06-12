---
source: "https://www.nextplatform.com/compute/2026/06/11/aws-tunes-up-graviton5-for-agentic-ai-boosts-bang-for-the-buck-bigtime/5254538"
hn_url: "https://news.ycombinator.com/item?id=48502864"
title: "AWS Tunes Up Graviton5 for Agentic AI, Boosts Bang for the Buck Bigtime"
article_title: "AWS Tunes Up Graviton5 For Agentic AI, Boosts Bang For The Buck Bigtime"
author: "rbanffy"
captured_at: "2026-06-12T13:18:12Z"
capture_tool: "hn-digest"
hn_id: 48502864
score: 2
comments: 0
posted_at: "2026-06-12T11:43:08Z"
tags:
  - hacker-news
  - translated
---

# AWS Tunes Up Graviton5 for Agentic AI, Boosts Bang for the Buck Bigtime

- HN: [48502864](https://news.ycombinator.com/item?id=48502864)
- Source: [www.nextplatform.com](https://www.nextplatform.com/compute/2026/06/11/aws-tunes-up-graviton5-for-agentic-ai-boosts-bang-for-the-buck-bigtime/5254538)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T11:43:08Z

## Translation

タイトル: AWS が Agentic AI 用に Graviton5 を調整し、大きな利益をもたらす
記事のタイトル: AWS が Agentic AI 向けに Graviton5 を調整し、大きな利益をもたらす
説明: 12 月に遡ると、アマゾン ウェブ サービスのアンナプルナ ラボチップ部門はそのプレビューを披露しました...

記事本文:
メインコンテンツへジャンプ
検索
その他のトピック
あらゆるセクションの最新ニュースをすべてお届けします
AWS が Agentic AI 向けに Graviton5 をチューンアップし、大幅なコスト削減を実現
12 月に遡ると、アマゾン ウェブ サービスのアンナプルナ ラボチップ部門は Graviton5 Arm サーバー CPU のプレビューを披露し、このチップがどのようなものであるかについていくつかのヒントを得ました。今週、Graviton5 は新しい M9g および M9gd インスタンスで出荷され、AWS はいくつかの空白を埋めて Graviton5 に関する詳細を提供してくれました。
いきなりですが、AWS が 7 か月前の re:Invent カンファレンス中に示したブロック図は正確ではありませんでした。 96 ペアの「Poseidon」Neoverse V3 コアを備えたモノリシック ダイが示されました。結局のところ、Graviton5 チップは 4 つの CPU ブロックで構成されており、それぞれが独自の 48 個の V3 コアと、関連するメモリおよび I/O コントローラーを備えています。これは、AWS が Arm Holdings が作成し、自社の AGI CPU で使用している Poseidon Compute Subsystem チップのブロックを取り出し、64 コアから 48 コアに削減し、Arm ダイツーダイ相互接続を使用して、192 個の V3 コア、12 個の DDR5 メモリ コントローラー、および 96 レーンがあり、CXL をサポートしていると思われる 8 個の PCI-Express 6.0 コントローラーを備えた Graviton5 ソケットを作成したようです。 3.0 メモリ拡張プロトコル。
この後者の部分は、AWS が手頃な価格で Graviton5 ソケットに搭載できる以上のメモリ容量を必要とするインメモリ データベースやワークロードにとって重要になります。 (DIMM に容量を追加すると、メモリ スティックのコストが増加します。つまり、線形ではなく指数関数的に増加します。そのため、容量のニーズを満たす最も細いメモリを選択し、その容量に対して最大のメモリ帯域幅を得るために常にすべてのメモリ スロットを埋めることになります。)
ここで 4 つの個別の Graviton5 チップレットを確認できます。
4 つの D2D 相互接続ラインがあります

4 つのチップレットを仮想プロセッサに組み込むと、それらのリンクは 420 GB/秒で実行されます。これらの相互接続は多くのエネルギーを消費しますが、4 つのチップレットを使用することで、台湾積体電路製造会社のレチクルの限界に達するモノリシック設計よりも小さいチップの歩留まりがはるかに高くなるため、各チップレットのコストが低くなります。このチップレットあたりのコストの低下は、Graviton4 チップで使用されている 4 ナノメートル プロセスから、より高価ではあるがトランジスタ密度が高く電力効率の高い 3 ナノメートル プロセスに切り替えることで軽減されます。
私の推定では、96 コアの Graviton4 には 730 億個のトランジスタがあり、AWS は初めて 2 ソケットの NUMA マシンを作成して、2 年後の 2023 年 11 月に Graviton4 が発表されたときの Graviton5 と同様の単一ノードのパフォーマンスを実現しました。
Graviton4 は基本的に Graviton3 と同じアーキテクチャを備えており、中央のモノリシック コア チップレットが個別のメモリとそれにリンクされた I/O コントローラ チップで囲まれています。 Graviton3 には 64 個の「Zeus」 V1 コアがあり、Graviton4 には 96 個の「Demeter」 V2 コアがありました。
上の表からわかるように、L1 キャッシュと L2 キャッシュはコア数に比例して増加していますが、AWS はコア数よりも早く L3 キャッシュを増加させています。 Graviton5 にはコアあたり 2 MB の L2 キャッシュがあり、192 コア全体で合計 384 MB の L2 キャッシュがありますが、コアあたりの L3 キャッシュは Graviton5 の 2 倍の 384 MB です。
L3 キャッシュの増加、V3 コアと DDR5 メモリのクロック速度の高速化、および Graviton5 ソケット内の 4 つのチップレット間の D2d 相互接続により、Graviton5 のワット数が大幅に増加します。重量は約 650 ワットだと思います。これは、ワットあたりのパフォーマンスが Graviton4 の半分であることを意味します。ただし、ソケットごとに 2.4 倍のパフォーマンスが得られます。

Graviton5 と、NUMA 共有メモリ構成の 2 つのチップを備えた Graviton4 ノードを比較しても、単一の Graviton5 は、単一チップのみで 25 より多くの raw スループットを持っています。
データベースとエージェント AI をサポートするための応答性の高いシステムの必要性を考慮すると、これはすべて公平なトレードオフであり、低発熱よりも低遅延が必要となります。
以下は、ローカル フラッシュ ストレージを持たない Graviton5 を使用する M9g インスタンスが、同様にノードにローカル フラッシュ ストレージを持たない Graviton4 に基づく以前の R8g および X8g と比較してどのようにスタックされるかを示しています。
インスタンスの価格とパフォーマンスを、インスタンスの 1 年間のオンデマンド価格と比較すると、次のようになります。
M9g および M9gd インスタンス (後者にはローカル フラッシュが追加されます) と X8g インスタンスのインスタンス節約プランの価格は、何らかの理由で公に発表されていませんが、コンピューティング インスタンスの価格設定があり、これは値下げにはやや積極的ではありませんが、データセンター内やリージョン全体で他のインスタンスに移行するという点ではより柔軟です。オンデマンド料金は上に示したとおりです。
すぐに思い浮かぶのは、AWS が M9g インスタンスの料金を X8g インスタンスの約半額に設定していることです。M9g インスタンスのメモリ容量は 4 分の 1 または半分であるため、これは当然のことです。 (各ファミリ内のインスタンスによって異なります。)これは、DRAM とフラッシュ メモリの不足による直接的な影響です。 AWS は、ローカル ストレージを備えたバリアントにおいても、フラッシュ容量をそれほど寛大にする必要はありません。コンピューティングは安価ですが、メモリはそうではありません。
以下は、初期の Graviton2 および Graviton3 インスタンスの価格/パフォーマンスの表です。これらは、非常に優れた機能を備えた Graviton4 および Graviton5 インスタンスと比較すると、それほど印象的ではありませんでした。 M9g インスタンスは、31.9 ～ 33.6 パーセント優れたパフォーマンスを提供します。

同じ vCPU 数を持つ最も同等の R8g インスタンスよりもコストが高く、ローカル フラッシュを備えた M9gd インスタンスは、R8gd インスタンスよりも 22.1 ～ 32 パーセント優れた価格/パフォーマンスを実現します。高メモリ X8g インスタンスは非常に高価です。このインスタンスを選択するには、本当に大きなメモリ フットプリントが必要です。
Graviton5 をベースにした新しい M9g および M9gd インスタンスは、Graviton4 CPU を使用した以前の R8g および R8gd インスタンスと比較して、メモリが少し少ないと思われるかもしれません。 Graviton5 インスタンスの容量は Graviton4 インスタンスの 4 分の 1 から半分ですが、帯域幅に敏感なワークロードの場合、この容量はそれほど重要ではない可能性があります。私たちは、そう遠くない将来、より多くのメモリを追加する、より大容量のメモリを搭載した X9g インスタンスが登場すると強く疑っていますが、最近の DRAM のコストを考慮すると、その追加メモリには多額の割増料金を支払うことが予想されます。さらに、AWS は、DDR5 フォームファクターで利用可能な最速の 8.8 GHz の DRAM を使用しています。これらの Graviton5 インスタンスは、帯域幅で不足している可能性のある容量を補います。
この追加のメモリ コストの一部は、CXL 3.0 メモリ エクステンダによって軽減される可能性があり、AWS がこの機能を提供するために Graviton5 ラックに共有メモリ アプライアンスを作成したとしても驚かないでしょう。 PCI-Express 6.0 スロットを備えたノードでは CXL 3.0 メモリ エクステンダが使用されるとは思いませんが、それは、DRAM を拡張して共有するためのラックスケール メモリ アプライアンスよりも退屈だからです。
AWS が Agentic AI 向けに Graviton5 をチューンアップし、大幅なコスト削減を実現
ゲノミクスサンプルあたりのコストは?シーケンス試行あたりのコストを試す
D-Wave はゲート モデル量子の野望のためにデュアル レールに乗ります
AIチップ・シェパーズ、ブロードコムとマーベルが金羊毛の皮を剥いだ
チップ容量の制約により AI 支出が抑制される

成長
Intel Xeon 6+ によるサーバー統合のためのパフォーマンスの強化
HPE、エンタープライズ、ソブリン、ネオクラウドで最初の GenAI の波に乗ります
シスコ、AI エージェントの同僚、フロンティア モデルの脅威の世界に備える
AI
Nvidia が AI データセンターの支配力を外側に拡大
デル、急成長する AI サーバーの量で利益を上げる
AI スタックの中心となるデータとストレージ
GPUとRAMは不足しているが、AIの本当のボトルネックは電気技師にある
アンクル・サム、量子企業に20億ドル以上を授与、しかしその取り分を望んでいる
800 ボルトのデータセンター電源の充電はそれほど速くない
HPC
Oak Ridge が量子、古典的 HPC、AI システム スタックの織り込みを開始
計算する
AI インフラストラクチャがオンプレミスに移行するにつれて、デルはハードウェアを大量に増強
接続する
シスコ、マーチャント シリコンとオプティクスで AI 顧客を獲得
IPO の完了により、Cerbras は再び AI の限界を押し上げることができる
HPE、VMユーザーにライフラインを提供し、クラウドスタックでコンテナとVM管理を統合
OpenAI、マイクロソフト、そしてその友人たちが、より優れた、よりスケーラブルなイーサネットを構築
コンピューティングとメモリの価格高騰により IT 支出が大幅に増加
AI システムが冷却を保つ唯一の方法が空気である場合もあります
Arista は AI スケールアウトネットワークに乗り、スケールアクロスに移行し、スケールアップを待つ
Compute Engine を作成できれば、Compute Engine を販売できます
クリーブランドクリニック、量子中心のスーパーコンピューティングで巨大タンパク質をシミュレーション
計算する
Broadcom が CPU および XPU メーカーのコンピューティングの垂直化を支援
雲
Microsoft、AI インフラストラクチャを 2 年間で倍増させることを約束
雲
Google はフルスタック AI プレーヤーであり、好調です
雲
AWS は Google やおそらく Microsoft と同様に OEM となる
ハイエンド コンピューティングを詳しくカバー
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
レジスター
開発クラス
ブロックとファイル
クッキーポリシー
プライベート

y ポリシー
利用規約
私の個人情報を販売しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Back in December, the Annapurna Labs chip division of Amazon Web Services showed off a preview of it ...

Jump to main content
Search
More topics
All the latest news, from all sections
AWS Tunes Up Graviton5 For Agentic AI, Boosts Bang For The Buck Bigtime
Back in December, the Annapurna Labs chip division of Amazon Web Services showed off a preview of its Graviton5 Arm server CPU , and we got some hints about what this chip might look like. This week, the Graviton5 is shipping in new M9g and M9gd instances, and AWS has given us some more details about the Graviton5, filling in some blanks.
Right off the bat, the block diagram that AWS showed during the re:Invent conference seven months ago was not accurate. It showed a monolithic die with 96 pairs of “Poseidon” Neoverse V3 cores. As it turns out, the Graviton5 chip is comprised of four CPU blocks, each with their own 48 V3 cores and the associated memory and I/O controllers. This looks like AWS picked up a block of the Poseidon Compute Subsystem chip that Arm Holdings created and is using in its own AGI CPU and cut it back from 64 cores to 48 cores and used the Arm die-to-die interconnect to make a Graviton5 socket with 192 V3 cores, a dozen DDR5 memory controllers, and eight PCI-Express 6.0 controllers that I think have 96 lanes and that support the CXL 3.0 memory extension protocol.
This latter bit will be important for in-memory databases or workloads that need more memory capacity than AWS can affordably put on a Graviton5 socket. (Fatter memory sticks cost increasingly more as you add capacity to the DIMM – it scales exponentially, not linearly. So you go with the skinniest memory that meets your capacity needs and you always fill all the memory slots to get maximum memory bandwidth against that capacity.)
You can see the four individual Graviton5 chiplets here:
There are four D2D interconnects linking the four chiplets into a virtual processor, and those links run at 420 GB/sec. These interconnects burn a lot of energy, but by using four chiplets, the cost of each chiplet is lower because the yield is much higher for a smaller chip then it would be for a monolithic design pushing up against reticle limits at Taiwan Semiconductor Manufacturing Co. This lower per chiplet cost is mitigated by a switch from 4 nanometer processes used with the Graviton4 chip to the much more expensive but transistor dense and more power efficient 3 nanometer process.
I estimate that the Graviton4 with 96 cores had 73 billion transistors, and for the first time AWS created a two-socket NUMA machine to get single-node performance akin to the Graviton5 that was two years into the future when Graviton4 was revealed in November 2023.
The Graviton4 had essentially the same architecture as the Graviton3, with a central, monolithic core chiplet surrounded by separate memory and I/O controller chips linked to it. Graviton3 had 64 “Zeus” V1 cores and Graviton4 had 96 “Demeter” V2 cores.
As you can see from the table above, the L1 cache and L2 cache have risen linearly with core count, but AWS has increased the L3 cache faster than the core count. The Graviton5 has 2 MB of L2 cache per core, for a total of 384 MB of L2 cache across those 192 cores, but it has 384 MB of L3 cache per core, twice that of the Graviton5.
Between the increased L3 cache, the faster clock speeds on the V3 cores and the DDR5 memory, and the D2d interconnects between the four chiplets in the Graviton5 socket radically increase the wattage of the Graviton5. I think it weighs in at around 650 watts, which means the performance per watt is half that of the Graviton4. But you get 2.4X more performance per socket with Graviton5, and even comparing the Graviton4 node with two chips in a NUMA shared memory configuration, a single Graviton5 has 25 more raw throughout with only a single chip.
This is all fair tradeoffs given the need for responsive systems to support databases and agentic AI, which need low latency more than they need low heat.
Here is how the M9g instances using Graviton5, which do not have local flash storage, stack up against their R8g and X8g predecessors based on Graviton4, which also do not have local flash in the node:
And here is how the instance pricing and performance stacks up against on demand pricing for the instance for a year:
The instance savings plan pricing for the M9g and M9gd instances (the latter adds local flash) and the X8g instances have not been announced publicly for some reason, but there is compute instance pricing, which is a little bit less aggressive on the price cuts but which is more flexible in terms of converting to other instances within datacenters and across regions. On demand pricing is shown above.
What pops out immediately is that AWS is charging about half the price for M9g instances as it does for the X8g instances, and that stands to reason since the M9g instances have a quarter or a half of the memory capacity. (It varies by instance within each family.) This is the immediate impact of the DRAM and flash memory crunch. AWS has to be less generous with flash capacity as well in the variants that have local storage. Compute is cheap, memory is not.
Here is a price/performance table across the early Graviton2 and Graviton3 instances, which were not all that impressive by comparison to the Graviton4 and Graviton5 instances, which do pack a wallop. The M9g instances deliver between 31.9 percent and 33.6 percent better bang for the buck than the most equivalent R8g instances with the same vCPU counts, and the M9gd instances with local flash deliver between 22.1 percent and 32 percent better price/performance than the R8gd instances. The high memory X8g instances are very pricey. You have to really need that large memory footprint to pick this instance.
You might be thinking that the new M9g and M9gd instances based on Graviton5 are a little skimpy on the memory compared to their R8g and R8gd predecessors using the Graviton4 CPU. The capacity on the Graviton5 instances is a quarter to a half that of the Graviton4 instances, but for bandwidth sensitive workloads this capacity may not matter as much. We strongly suspect that there will be heavier-memoried X9g instances in the not too distant future that add more memory, but expect to pay a hefty premium for that extra memory given the cost of DRAM these days. Moreover, AWS is using the fastest DRAM available in the DDR5 form factor, at 8.8 GHz speeds. These Graviton5 instances make up in bandwidth what they might be lacking in capacity.
Some of this extra memory cost might be mitigated by CXL 3.0 memory extenders, and I would not be surprised if AWS has created a shared memory appliance in its Graviton5 racks to deliver this function. I do not think CXL 3.0 memory extenders will be used in the node with PCI-Express 6.0 slots, but that is just because that is more boring than a rackscale memory appliance for extending and sharing DRAM.
AWS Tunes Up Graviton5 For Agentic AI, Boosts Bang For The Buck Bigtime
Cost Per Genomics Sample? Try Cost Per Sequencing Attempt
D-Wave Riding The Dual-Rail For Its Gate-Model Quantum Ambitions
AI Chip Shepherds Broadcom And Marvell Have Skinned The Golden Fleece
Chip Capacity Constraints Put A Governor On AI Spending Growth
Enhanced Performance For Server Consolidation With Intel Xeon 6+
HPE Catches Its First GenAI Wave With Enterprises, Sovereigns, And Neoclouds
Cisco Preps For A World Of AI Agent Coworkers, Frontier Model Threats
AI
Nvidia Extends Its Grip On The AI Datacenter Outwards
Dell Makes The Profits Up In Volume For Booming AI Servers
Data And Storage At The Center Of The AI Stack
GPUs And RAM Are In Short Supply, But The Real Bottleneck For AI Is Electricians
Uncle Sam Awards $2 Billion-Plus To Quantum Companies, But Wants A Cut
Not So Fast On That Charge For 800 Volt Datacenter Power
HPC
Oak Ridge Starts Weaving Together A Quantum, Classical HPC, And AI System Stack
compute
Dell Bulks Up Hardware As AI Infrastructure Shifts To On-Premises
connect
Cisco Wins Over AI Customers With Merchant Silicon And Optics
With Its IPO Done, Cerebras Can Get Back To Pushing The AI Envelope
HPE Throws VM Users A Lifeline, Unifying Containers And VM Management In Cloud Stack
OpenAI, Microsoft And Friends Build A Better, More Scalable Ethernet
Compute And Memory Price Hikes Drive IT Spending Way Higher
Sometimes, Air Is The Only Way For AI Systems To Keep Their Cool
Arista Rides AI Scale Out Networks, Moves Into Scale Across, And Awaits Scale Up
If You Can Make A Compute Engine, You Can Sell A Compute Engine
Cleveland Clinic Simulates Large Proteins With Quantum-Centric Supercomputing
compute
Broadcom Helps CPU And XPU Makers Go Vertical With Compute
cloud
Microsoft Committed To Doubling AI Infrastructure In Two Years
cloud
Google Is A Full Stack AI Player, And Is Playing Well
cloud
AWS Will Be An OEM, Just Like Google And Maybe Microsoft
In-depth coverage of high end computing
Contact us
Advertise with us
Who we are
Newsletter
The Register
DevClass
Blocks and Files
Cookies Policy
Privacy Policy
Ts & Cs
Do not sell my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
