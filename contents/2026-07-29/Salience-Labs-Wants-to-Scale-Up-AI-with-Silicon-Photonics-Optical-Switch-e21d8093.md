---
source: "https://www.nextplatform.com/connect/2026/07/22/salience-labs-wants-to-scale-up-ai-with-silicon-photonics-optical-switch/5276643"
hn_url: "https://news.ycombinator.com/item?id=49095549"
title: "Salience Labs Wants to Scale Up AI with Silicon Photonics Optical Switch"
article_title: "Salience Labs Wants To Scale Up AI With Silicon Photonics Optical Switch"
author: "rbanffy"
captured_at: "2026-07-29T10:30:40Z"
capture_tool: "hn-digest"
hn_id: 49095549
score: 1
comments: 0
posted_at: "2026-07-29T10:23:25Z"
tags:
  - hacker-news
  - translated
---

# Salience Labs Wants to Scale Up AI with Silicon Photonics Optical Switch

- HN: [49095549](https://news.ycombinator.com/item?id=49095549)
- Source: [www.nextplatform.com](https://www.nextplatform.com/connect/2026/07/22/salience-labs-wants-to-scale-up-ai-with-silicon-photonics-optical-switch/5276643)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T10:23:25Z

## Translation

タイトル: Salience Labs はシリコン フォトニクス光スイッチを使用して AI をスケールアップしたいと考えています
記事のタイトル: Salience Labs はシリコン フォトニクス光スイッチで AI をスケールアップしたいと考えています
説明: 光回線スイッチは新しいものではなく、ネットワーク技術の研究室や企業で使用されています。

記事本文:
メインコンテンツへジャンプ
検索
その他のトピック
あらゆるセクションの最新ニュースをすべてお届けします
Salience Labs はシリコン フォトニクス光スイッチで AI をスケールアップしたいと考えています
光回線スイッチは新しいものではなく、過去 25 年間、ネットワーク技術研究室や通信業界の生産現場で使用されてきました。しかし、Salience Labsによると、ハイパフォーマンスコンピューティングクラスタのスケールアップネットワーク領域に移行する時期が来たという。同研究所は完全にシリコンフォトニクス技術に基づいた光回路スイッチを構築しており、皮肉なことに、この技術はフォトニクスコンピューティングプラットフォームを構築しようとする高度な研究から得られたものである。
2 種類の光スイッチが登場し、現在実稼働環境で使用されています。 1 つは微小電気機械システム (MEMS) ミラー技術に基づいており、小さな機械ミラーのアレイが回転して光ファイバー ケーブルのストランドを接続する回路を作成します。 Google は、最初の 3 世代の TPU AI アクセラレータ システムを直接配線接続しました (これまで Nvidia がノードスケールおよびラックスケール GPU 設計で行ってきたのと同じように) が、2015 年以降の過去 4 世代にわたって、Google は TPU クラスターのバックボーンとして「Apollo」 OCS の一部である「Palomar」 MEMS デバイスを使用してきました。 OCS を使用すると、クラスタをオンザフライで再構成できるほか、最大 9,216 個の TPU をコヒーレント メモリ クラスタ内で結合することもできます。
さまざまな光学技術に関して Nvidia と提携している Lumentum は、MEMS ベースの OCS ギアも販売しており、通信以外の分野や AI データセンターにも野心を抱いています。製造中のもう 1 つの OCS テクノロジは、液晶オンシリコン (LCoS) と呼ばれるもので、このアプローチは Coherent の OCS デバイスで使用されています。当然のことですが、Nvidia は次の企業とも提携しています。

さまざまな光学技術に対応するコヒレント。
Salience Labs はその特定のスイッチング メカニズムを明らかにしていませんが、OCS スタートアップとの会話から、回路の切り替えに MEMS 技術も LCoS 技術も使用していないことは確かにわかっています。このスタートアップは英国のオックスフォード大学とミュンスター大学からスピンアウトし、具体的にはオックスフォード大学の応用ナノマテリアル教授であるHarish Bhaskaran氏と、ミュンスター大学の実験物理学の教授であるWolfram Pernice氏が率いるチームの研究に基づいて設立されました。両者は Salience Labs の共同創設者であり、共同研究活動では相変化オプトエレクトロニクスに注目していました。
Salience Labs が、PH18DA 統合型 III-V レーザーおよび TPS45PH 低損失窒化シリコン導波路に関して Tower Semiconductor と提携しているという事実は、同社がシリコン フォトニクス フレームワーク内でスイッチングを行うためにある種の相変化技術を使用していることを示すもう 1 つの指標です。
それが何であれ、Salience Labsの共同創設者兼最高経営責任者であるVaysh Kewada氏は、同社が32ポートのOCSを展開し、AIネットワークをスケールアップするために64ポートと128ポートのOCSの提供を検討しているにもかかわらず、The Next Platformにも他の誰にも語っていません。しかし、それは、熱光学マッハツェンダー干渉計、ある種の電気光学素子、または光スイッチングの心臓部であると予想される相変化材料など、さまざまなものである可能性があります。
重要なのは、それがミラーや LCD ではないということです。これらは非常に遅く、Salience Labs が使用していると思われるオンチップ導波路スイッチングでは 300 マイクロ秒未満でポート間のリンクを再構成するのに数ミリ秒かかります。
Kewada 氏が強調しているのは、光回路スイッチングを根本から見直し、実際に機能する真のシリコン フォトニクス製品を手に入れる時期が来たということです。

XPU のフリート全体でメモリ ファブリックをスケールアップするだけでなく、Google が Apollo スイッチで行っているようなスケールアウト ネットワーク内のいくつかのアグリゲーション レイヤーもスケールアップします。
「私たちが OCS の開発を選択したのは、CPO、NPO、XPO のフロントで、非常に多くのプレーヤーがソリューションを開発し、チップからデータを光ファイバーにどのように取得するかという問題を解決しようとしていると考えたからです」と Kewada 氏は The Next Platform に語ります。
ちなみに、ケワダ氏はインペリアル・カレッジ・ロンドンで物理学の学士号と修士号を取得しており、オックスフォード・サイエンス・エンタープライズのインキュベーターに在籍する起業家でもありました。
「しかし、私たちにとって興味深かったのは、この傾向が確実に近づいているということでした」とケワダ氏は続けます。 「データセンター内でより多くの光接続が確立されると、スイッチング アーキテクチャはどのようになるかを考え、非常に単純な仮説に基づいて製品を開発しました。プラガブル、CPO、NPO、その他の形式を問わず、データセンターには光接続が増えるでしょう。そうなると、スイッチング層はより異質なアーキテクチャ、つまりある場合には電気パケット スイッチと他の場合には光回線スイッチの組み合わせに移行するでしょう。私たちの見解では、OCS は破壊の機が熟した市場でした。現在市場で入手可能な OCS は、基本的には 20 年前のテクノロジーに基づいています。」
Salience Labs OCS は 2 つのチップで構成されています。 1 つは完全に統合されたシリコン フォトニクス OCS で、2 つ目は増幅および信号調整チップです。この後者のチップは、銅線回路のリタイマーやリドライバーに似ており、帯域幅が広いため、ワイヤの長さがますます短くなり、禁止がますます高くなるとノイズの問題が発生します。

幅。
上の画像でカードの両面が見られるように、2 つのチップは PCB カード上に背中合わせに実装されています。この合成写真では、32 ポート テスト システムの前面パネルも見ることができます。
「OCS を使用すると、大量生産に到達するために統合チップ形式で製品を生産できるようになり、ポートあたりのコストが高く、製造性が向上し、統合形式であることのすべての利点が得られるようにしたいと考えています」と Kewada 氏は説明します。 「シリコンフォトニクスの主な欠点の1つは、常に損失の問題でした。オンチップに統合するとすぐに、ある程度の光損失が発生することになります。そこで、私たちは独自のコンポーネント設計である増幅を使用してアレイ上に製造することでその問題を解決しました。これにより、部品表のコスト目標を達成することができます。そしてそれが私たちのソリューションの核心です。この増幅チップは他に類を見ないものであり、スタックの核心です。私たちはまた、コアアーキテクチャについても多くの作業を行っています。」ポート数を継続的に拡張できるようにするためです。」
Salience Labs が 3 月に発表し、現在運用を開始している OCS には 32 ポートがありますが、アーキテクチャは 64 ポート、128 ポート、256 ポートまで拡張できます。 OCS からの出力は、PAM4 変調 (信号あたり 2 ビット) を使用した 100 Gb/秒のネイティブ ライン レートで、レーンあたりの実効帯域幅は 200 Gb/秒になります。
Salience Labs と話している潜在的な顧客が熱心に考えているのは、AI システムのスケールアップ ドメインのサイズを拡大することですが、現時点では Nvidia と AMD の両方の 72 GPU にとどまっており、スイッチの遅延が短く、再構成時間が比較的短いことを考慮すると、多数の OCS デバイスがここで役立つ可能性があります。一部の顧客は、GPU やその他の XPU のコヒーレント メモリ ドメインに対して 1 つのラックを超えて使用したいと考えており、また、コヒーレント メモリの代替を必要としている顧客もいます。

現在、ラックスケール マシンで利用可能なファブリックがあり、さらに他の企業は、GenAI のデコード フェーズで複数のマシンを結合するネットワークの最初の層に OCS を使用することを検討しています。これらのさまざまなオプションを、一連の AI システム全体にわたってオンザフライで再構成できれば素晴らしいでしょう。
Salience Labs が現在出荷している初期の OC-32M デバイスのポート間のホップは 10 ナノ秒未満、つまりメイン メモリの速度です。 Broadcom Tomahawk Ultra Ethernet ASIC 上のポート間のホップは約 250 ナノ秒です。これは、スイッチ全体のレイテンシが 25 分の 1 低いことになります。システム設計者なら誰でも、メモリとそれに関連するコンピューティングをより緊密に結合するために、その低いレイテンシを望むでしょう。他のイーサネット スイッチは、非常に優れたものでは 450 ナノ秒から 650 ナノ秒程度ですが、メモリ ファブリックのスケールアップには全く適さないものでは数ミリ秒になる場合もあります。
Salience Labs が競合他社 (少なくともこの表ではイーサネット スイッチと MEMS ベースの OCS デバイス) とどのように比較しているかを以下に示します。
注目すべき重要な点は、SiPho OCS スイッチ チップのリコンフィギュレーション時間が MEMS デバイスよりもはるかに短いことです (3 倍以上)。どちらのタイプの OCS でも、ポートあたりのエネルギー消費量は、比較的高速なイーサネット スイッチと比較して 8 分の 1 です。Tomahawk Ultra については、まだ十分な知識がありませんので、より具体的には言えません。
以下は、OC-32M スイッチ シャーシのより良いショットです。
そして、これはアンプチップとそこから出ているファイバーのズームショットです。
この特定のボックスには OCS ソケットが 1 つありますが、Salience Labs によれば、これらのモジュールを 1 つの 1U シャーシに最大 8 つパッケージ化できるそうです。
これは、Salience Labs がスケールアップ領域での OCS の採用に関して行っている、採用に向けた議論の要点です: Y

低レイテンシの推論を実現するために 2 種類の異なるコンピューティング エンジンを使用する必要はありません。 Salience Labs が、576 個の GPU を束ねた仮想マシン (この場合は Nvidia 独自の Megatron 2T モデル) のパレート曲線に沿ったパフォーマンスを OCS スイッチがどのように期待しているかを示します。
スケールアップ レイヤーでの OCS のレイテンシが低くなったことで、デュアル システム アーキテクチャに移行することなく、ユーザーあたりの 1 秒あたりのトークン スループットが 80% 向上します。これは主にメモリ ファブリック スイッチのレイテンシがはるかに低いためです。 OCS を使用することで、複数のラックにスケールアップしてスループットを向上させながら、GPU を使用し続けることができます。
Nvidia と AMD が耳を傾けてくれることを願っています。
AMDがお金を追いかけるために使用しているラックスケールAIシステムのロードマップ
AI ホストとサンドボックスがインテルのデータセンター CPU Cookie を保存
Nvidia、AI エージェントでチップ エンジニアリングを加速
AMDがラックスケールAIシステムのロードマップで追い求めている資金
米国エネルギー省、AI をターゲットとしたジェネシス ミッションに向けて 50 億ドルのスターター ガンを発射
Google の Axion CPU がクラウドにもたらすメリット
Salience Labs はシリコン フォトニクス光スイッチで AI をスケールアップしたいと考えています
GenAI のハードウェア投資はモデルとプラットフォームの収益を大きく上回っています
Microsoft、大規模な AI CPU および GPU クラスターのために AMD を活用
Google、量子誤り訂正に AI 強化学習を使用
マーベル、Teralynx T100 で基数、低遅延、帯域幅を実現
AI チップが TSMC の収益の約 3 分の 1 を占める
SRE から AI エージェントへ: 本番環境に触れる前に自分自身を証明する
量子古典 HPC データセンターにおける HPE とデルの願望
AI 対応データが真の利点である理由
QuiX Quantum が HPC データセンター向けのフォトニック アーキテクチャを披露
AI コンピューティングが進むにつれてイーサネット ネットワーキングも進む
もちろんメタプラットフォームはクラウドになるだろう
3 人の HPC 達人が尋ねる: まだ GPU は必要ですか?
光学Sc

ale Up ファブリックはアーキテクチャではなく製造によって制限されます
AMD、フラッシュ拡張メモリでサーバー DRAM を拡張
メモリ市場の好不況サイクルの終焉
中国の「LineShine」オール CPU、エクサフロップスクラスのスーパーコンピューターの詳細
HPE、セキュリティ、主権、マルチテナントのためのアップグレードされた HPC ハードウェア、ソフトウェアを提供
HPC スーパーコンピューティングでは AMD と Nvidia が互角です
企業
HPE、Agentic AI の波に乗ってデータセンターに再び参入
店
Everpure の AI 戦略はほぼ純粋に Nvidia に基づいています
計算する
サーバーブームにより価格上昇とチップ不足が両立
接続する
HPE のデータセンター ネットワーキングの全体像がより明確に焦点になる
ハイエンド コンピューティングを詳しくカバー
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
レジスター
開発クラス
ブロックとファイル
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を販売しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Optical circuit switches are not new and have been used in network technology research labs and in t ...

Jump to main content
Search
More topics
All the latest news, from all sections
Salience Labs Wants To Scale Up AI With Silicon Photonics Optical Switch
Optical circuit switches are not new and have been used in network technology research labs and in the telecom industry in production for the past two and a half decades. But the time has come for them to move into the scale up network domain in high performance computing clusters according to Salience Labs, which has built an optical circuit switch that is based entirely on silicon photonics technology and that, ironically enough, is derived from advanced research trying to create a photonics computing platform.
Two types of optical switches have emerged and are used in production today. One is based on micro-electrical-mechanical systems (MEMS) mirror technology, where arrays of tiny mechanical mirrors spin to create circuits linking strands of fiber optic cables. Google hardwired the first three generations of its TPU AI accelerator systems together directly (just like Nvidia has done with its nodescale and rackscale GPU designs thus far), but for the past four generations since 2015, Google has used its “Palomar” MEMS devices that are part of the “Apollo” OCS to be the backbone of its TPU clusters. The OCS allows for clusters to be reconfigured on the fly, and also allows up to 9,216 TPUs to be lashed together in a coherent memory cluster.
Lumentum, which has partnered with Nvidia for various optical technologies , also sells MEMS-based OCS gear and has aspirations outside of telecom and in the AI datacenter. The other OCS technology in production is called liquid crystal on silicon (LCoS), and this approach is used by Coherent in its OCS devices. Not surprisingly, Nvidia has also partnered with Coherent for various optical technologies .
Salience Labs has not divulged its particular switching mechanism, but we know for sure from talking to the OCS startup that it is using neither MEMS or LCoS techniques to flip through circuits. The startup spun out of Oxford University in the United Kingdom and the University of Münster, and specifically is founded on the work of teams led by Harish Bhaskaran , a professor of applied nanomaterials at Oxford, and Wolfram Pernice , a professor of experimental physics at Münster. Both are co-founders at Salience Labs and a joined research effort was looking at phase change optoelectronics.
The fact that Salience Labs has partnered with Tower Semiconductor for its PH18DA integrated III-V lasers and TPS45PH low-loss silicon nitride waveguides is another indicator that it is using some sort of phase change technology to do the switching inside of a silicon photonics framework.
Whatever it is, Vaysh Kewada , co-founder and chief executive officer at Salience Labs, is not telling The Next Platform or anyone else even as the company is rolling out a 32 port OCS and is looking at delivering OCSes with 64 and 128 ports for scale up AI networks. But it could be any number of things, including a thermo-optical Mach-Zehnder interferometer, an electro-optic element of some sort, or the expected a phase change material that is the heart of the optical switching.
What matters is that it is not mirrors or LCDs, because these are very slow, taking milliseconds to reconfigure a port to port link compared to under 300 microseconds with the on-chip waveguide switching Salience Labs is likely using.
What Kewada is emphatic about is that it is time to shake up optical circuit switching and get a true silicon photonics product that can actually do scale up memory fabrics across fleets of XPUs as well as maybe some of the aggregation layers out in the scale out networks like what Google is doing with its Apollo switches.
“We chose to develop OCS because our view is that there was quite a large number of players developing solutions on the CPO, NPO, XPO front, trying to solve the question of how you get data off of the chip and onto the optical fiber,” Kewada tells The Next Platform .
And by the way, Kewada has a bachelor’s and master’s degree in physics from Imperial College London and was also an entrepreneur in residence of the Oxford Science Enterprises incubator.
“But what was interesting to us was that this trend was coming down the line,” Kewada continues. “When there were more optical connections established in the datacenter, we thought about what the switching architecture would look like, and so we developed our products with a very simple hypothesis: There would be more optical connections in the datacenter, whether pluggable, CPO, NPO, whatever format, and when that happened, the switching layers would move towards more heterogeneous architectures, i.e. combinations of electrical packet switches in some cases and optical circuit switches in others. Our view is OCS was a market ripe for disruption, because the OCSes that are available today on the market are based off a technology that fundamentally is twenty years old if not older.”
The Salience Labs OCS is comprised of two chips. One is a fully integrated silicon photonics OCS, and the second is an amplification and signal conditioning chip. This latter chip is akin to retimers and redrivers for copper circuits where they bandwidth is high by the length of the wire keeps getting shorter because of noise issues at higher and higher bandwidths.
The two chips are mounted back to back on a PCB card, as you can see the two sides of that card in the image above. You can also see the front panel of the 32 port test system in this composite photo.
“With OCS, you want to be able to produce the product in an integrated chip format to reach volume production, giving good cost per port, good manufacturability, and all the advantages of being in an integrated format,” Kewada explains. “One of the key disadvantages of being in silicon photonics has always been the question of loss. As soon as you integrate on chip, you are going to incur some optical loss, and so we have solved that using an amplification that is our own component design and fabricating it on arrays, which let us meet the bill of materials cost goals. And that's the crux of our solution. This amplification chip is one of a kind, and the crux of the stack. We are also doing a lot of work on our core architecture to ensure that we can continue to scale port counts.”
The OCS that Salience Labs launched in March and is now ramping in production has 32 ports, but the architecture scales to 64 ports, 128 ports, and 256 ports. The output from the OCS is 100 Gb/sec native line rate with PAM4 modulation (with two bits per signal) to get that to 200 Gb/sec effective bandwidth per lane.
What the potential customers talking to Salience Labs seem to be keen on is expanding the size of the scale up domain for an AI system, which is stuck at 72 GPUs from both Nvidia and AMD right now, and a bunch of OCS devices might help here given the low latency of the switch and the relatively low reconfiguration time. Some customers want to go beyond one rack for that coherent memory domain for GPUs and other XPUs, others want an alternative to the coherent memory fabrics currently available for rackscale machines, and still others are looking at using OCS for the first layer of networks gluing together multiple machines for the decode phase of GenAI. It would be great if these different options could be reconfigured on the fly across a row of AI systems.
The port to port hop on the initial OC-32M device that Salience Labs is now shipping is under 10 nanoseconds – that is main memory speed. The port to port hop on the Broadcom Tomahawk Ultra Ethernet ASIC is around 250 nanoseconds . That is a factor of 25X lower latency across the switch, and any system architect would want that low latency to more tightly couple memories and their associated compute. Other Ethernet switches are on the order of 450 nanoseconds to 650 nanoseconds for very good ones, and can be milliseconds for ones that are really not appropriate for scale up memory fabrics at all.
Here is how Salience Labs says it stacks up to the competition, which is Ethernet switches and MEMS-based OCS devices at least in this table:
The key thing to notice is how that reconfiguration time is much lower for the SiPho OCS switch chip than it is for a MEMS device – more than a factor of more than 3X. The energy consumption per port is 8X lower for either type of OCS compared to a relatively fast Ethernet switch – we don’t know enough about the Tomahawk Ultra yet to be more specific.
Here is a better shot at the OC-32M switch chassis:
And here is a zoom shot on the amplifier chip and the fibers coming off of it:
This particular box has one OCS socket, but Salience Labs says it can package up to eight of these modules into a single 1U chassis.
Here is the cut and dry of the argument for adoption that Salience Labs is making for the adoption of OCS in the scale up domain: You don’t have to have two different kinds of compute engines to get low latency inference. Here is how Salience Labs expects for its OCS switchery to performance along a Pareto curve for a hypothetical machine with 576 GPUs lashed together, in this case with Nvidia’s own Megatron 2T model:
That lower latency of the OCS at the scale up layer improves token per second per user throughput by 80 percent without having to move to a dual-system architecture, mostly due to the much lower latency of the memory fabric switch. By going OCS, you can scale up across multiple racks and boost throughput and still stay on GPUs.
I hope Nvidia and AMD are listening.
The Rackscale AI System Roadmaps That AMD Is Using To Chase Money
AI Hosts And Sandboxes Save Intel’s Datacenter CPU Cookies
Nvidia Accelerates Chip Engineering With AI Agents
The Money AMD Is Chasing With Its Rackscale AI System Roadmaps
DoE Fires The $5 Billion Starter Gun For Its AI-Targeted Genesis Mission
How Google's Axion CPUs Benefit The Cloud
Salience Labs Wants To Scale Up AI With Silicon Photonics Optical Switch
GenAI Hardware Investments Are Way Ahead Of Model And Platform Revenues
Microsoft Taps AMD For At Scale AI CPU And GPU Clusters
Google Uses AI Reinforcement Learning For Quantum Error Correction
Marvell Brings Radix, Low Latency, And Bandwidth To Bear With Teralynx T100
AI Chips Drive Around A Third Of TSMC Revenues
SREs To AI Agents: Prove Yourself Before You Touch Production
The Aspirations Of HPE And Dell In The Quantum-Classical HPC Datacenter
Why AI-Ready Data Is The Real Advantage
QuiX Quantum Shows Off A Photonic Architecture For HPC Datacenters
As Goes AI Compute, So Goes Ethernet Networking
Of Course Meta Platforms Is Going To Be A Cloud
Three HPC Gurus Ask: Do We Still Need GPUs?
Optical Scale Up Fabrics Are Limited By Manufacturing, Not Architecture
AMD Stretches Server DRAM With Flash Extended Memory
The End Of Boom/Bust Cycles For The Memory Market
A Deep Dive On China’s “LineShine” All-CPU, Exaflops-Class Supercomputer
HPE Delivers Upgraded HPC Hardware, Software For Security, Sovereignty, And Multi-Tenancy
AMD And Nvidia Are Neck And Neck In HPC Supercomputing
enterprise
HPE Rides The Agentic AI Wave Back Into The Datacenter
store
Everpure’s AI Strategy Is Almost Purely Based On Nvidia
compute
The Server Boom Balances Price Increases Against Chip Shortages
connect
HPE’s Datacenter Networking Picture Comes Into Clearer Focus
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
