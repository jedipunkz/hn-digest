---
source: "https://power2026.ai/"
hn_url: "https://news.ycombinator.com/item?id=49116741"
title: "Power 2026 – Electricity Pricing in the Age of AI"
article_title: "Power 2026 - Electricity Pricing in the Age of AI"
author: "nsomani"
captured_at: "2026-07-30T23:10:55Z"
capture_tool: "hn-digest"
hn_id: 49116741
score: 1
comments: 0
posted_at: "2026-07-30T22:35:59Z"
tags:
  - hacker-news
  - translated
---

# Power 2026 – Electricity Pricing in the Age of AI

- HN: [49116741](https://news.ycombinator.com/item?id=49116741)
- Source: [power2026.ai](https://power2026.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T22:35:59Z

## Translation

タイトル: Power 2026 – AI 時代の電力料金
記事のタイトル: Power 2026 - AI 時代の電力料金
説明: Power 2026 は AI の入門書です

記事本文:
インデックス
01
序文: 動機
発電所のすべて
02
商品価格設定の基礎
03
発電所
04
発電所を建てるには
05
ケーススタディ: ホーマーシティ
06
高まる需要に応える
電力を取引する方法
07
米国の電力市場
08
ケーススタディ: アルバータ州
09
生産コストモデル
10
実際的な近似
11
電力取引の種類
計算中...
パワー2026
AI時代の電力料金設定
AI の需要は爆発的に増加しています。しかし、GPU とメモリのサプライ チェーンについては多くの議論が行われていますが、AI の能力を拡大するための本当の制約は電力です。
データセンターはすでに米国の電力消費量の約 5% を占めています [1] 。データセンターの電力需要は 2 年ごとに 2 倍になるため [2] 、理論上は 2030 年代半ばまでに需要が米国の総発電量を上回ることになります。
データセンターの構築には、許可、周囲のインフラストラクチャ、予想される電力価格を考慮した慎重な計画が必要です。多くの場合、変動する電力コストによって、データセンターが存続できるかどうかが決まります。
私は大手ヘッジファンドで電力とガスを担当していた元クオンツ研究員で、ここ数年のほとんどをデータセンターの構築（GPU の調達、石炭火力発電所との交渉、用地の特定）について創業者や投資家にアドバイスすることに費やしてきました。
この電力とデータセンターに関する入門書は、エネルギー市場にチャンスがあると感じ、最新情報を知りたいと考えている幅広い読者を対象としています。
この入門書を読み終わるまでに、次のことが理解できるようになります。
パート 1: 発電所の仕組み、データセンターの開発方法、企業の対応方法 (創業者/投資家向け)。
パート 2: 米国電力市場の価格設定方法 (トレーダー向け)。
権力に本当に興味があるのなら、全部読むことをお勧めします。タイムリーで興味深い内容だと思っていただければ幸いです。
ティナ・ヒーとサミュエル・スピッツに感謝します。

日記的なフィードバック。
商品価格設定の基礎
この章の内容
商品は、供給、需要、純輸出、保管量の変化を結び付けるシステムバランス方程式によって支配されます。
競争市場では、商品の価格は追加の需要単位に対応するための限界費用に等しくなります。
送電網の安定化には物理的な制約が必要なため、特に電力市場は独立系システムオペレーター (ISO) によって管理されることがよくあります。
権力について詳しく説明する前に、より広い文脈の中で権力がどのような位置にあるのかを理解する必要があります。主要な商品カテゴリには、エネルギー (電力、原油、天然ガスなど)、農業、金属の 3 つがあります。すべての商品には、他の資産クラスとは異なる共通の基本的な特性があります。
株式とコモディティは、主に 2 つの点で異なります。
まず、株式には価格をその基礎的価値まで崩壊させるような強制的な機能はありません。これに対して、商品契約には固定の決済日があります。決済日に誰かがやって来て、その取引可能な商品を物理的に売買します。トウモロコシが高すぎると、彼らはそれを買わないでしょう。価格は現実に崩壊しなければなりません。
第二に、株式が表す基礎となる企業には、契約、共通の顧客、サプライチェーンなど、相互にさまざまな関係があり、それらを解きほぐすのが難しい場合があります。コモディティは、より直接的な方法で相互に供給されます。地面に掘削された坑口からは、ある程度の原油とある程度の天然ガス（主にメタン）が得られます。 [3] 原油は世界中に簡単に輸送できるので優れています。その世界市場は、海外の価格がここ米国にいる私たちに影響を与えることを意味しており、そのため、アメリカ海峡の閉鎖についての懸念が広がっています。

ホルムズ。その原油には市場価格があり、製油所はそれを購入して留出物（ガソリン、ディーゼル、ジェット燃料などの石油製品）を製造します。それはガソリンスタンドで買うのと同じガソリンです。
井戸から石油がほとんど生産されない場合、私たちはそれをガス専用井戸と呼び、国内の天然ガスの大部分はそこから供給されます。天然ガスはそのままでは簡単に輸送できませんが、冷却して液化天然ガス（LNG）の形で輸送することは可能です。最近、海外の天然ガスの価格が米国の価格をはるかに上回っているため、私たちはできる限り輸出しています。その結果、海外の天然ガス価格は、短期的にはここ米国にそれほど影響を与えません [4] 。単純化した意味で、ここ米国にある天然ガスは、物理的に輸出できなかった供給源です [5] 。天然ガスは、いくつかの産業用途 (例: 化学製造)、住宅・商業用途 (例: 家の暖房)、および電力 [6] - この入門書の焦点です。
商品市場は、それぞれの場所で特定の方法で「解決」する必要があります。
供給 = 需要 + 純輸出 + 貯蔵量の変化 [7]
供給量は、地元の商人が生産する必要がある量です。需要とは、特定の場所でどれだけ消費されるかです。純輸出とは、ある地域から輸出された単位から輸入されたものを差し引いたものです。貯蔵量の変化は、貯蔵量から貯蔵量を引いたものです。この方程式は全世界のあらゆる場所で成り立つはずです。追加の制約があります。
ある地点から別の地点に移送できる供給量には限界があります（利用可能なトラックや船舶の数が限られている、または送電線がサポートできる送電量の限界が限られています。これについては後で詳しく説明します）。
投入した以上の量をストレージから引き出すことはできず、一度に最大量の供給しか保管できません。
上記の各変数

s は価格の関数です。価格は消費者が支払い、供給者が受け取る金額です。ある場所の価格が上昇すると、そこでより多くの供給を生産することが経済的となり（効率の悪い生産者でも損失を出さずに需要に応えることができるようになります）、その地域の需要は下がります。もちろん、純輸出も減少し（存在する場合）、貯蔵庫（存在する場合）からのより多くの供給を奨励することになります。システムが「バランス」するまで、価格はどこでも変わります。
常に、販売されるすべてのユニットは、「最初に」販売されたユニットであろうと「最後に」販売されたユニットであろうと、同じ価格で販売されます。サプライヤーが商品を製造するのに本当に安いのであれば、おそらくもう少し安く請求しようとするだろうと思うかもしれません。しかし、競争の激しい市場では、販売する価格を個人的に「選択」することはできません。これは、商品の価格設定がどのように機能するかについての重要な前提条件です。では、その市場価格は何によって決まるのでしょうか？
需要/保管/輸出を満たす最後の供給ユニットを生産するには、フルキャパシティに達していない最も安価な生産者を選択します。その供給単位を生産するのにかかる費用を「限界費用」といいます。たとえば、100 バレルの石油が必要な場合、生産者 A は 99 バレルを 1 バレルあたり 5 ドルでオファーし、生産者 B は次に安いオファーを 1 バレルあたり 6 ドルで提供するでしょう。この場合、限界費用は 6 ドルになります。
競争市場では、誰もが販売する市場価格は限界費用とまったく同じです。あなたが最後の供給単位を提供した生産者 (つまり、生産者 B) である場合、その最後の単位で得られる利益は 0 ドルであることに注意してください。価格は限界費用と同じです。この性質は「限界価格」と呼ばれます [8] 。それは電力価格設定に非常に関係します。
これが真実である理由の簡単な正当化: 限界生産者が販売を拒否したと仮定しましょう。

彼らの限界費用。彼らはより高く売りたいだけなのです。問題は市場の競争が激しいことだ。その試みの市場価格を引き下げようとしている別のサプライヤーがいます。これは、市場価格が、追加の供給単位を提供できる最も安価な生産者の限界費用と正確に等しくなるまで発生し続けます。価格が低ければ、誰でもそのユニットを供給するのは不経済になります。
電力は物理的な「バス」でグリッドに注入またはグリッドから取り出されます [9] 。これらのバスは多数あり、米国のあらゆる都市内に分散しています。それらはすべて伝送線を介して接続されています。
特定の場所 (物理的な電力バス、都市全体、州全体など) の電力条件で上記の方程式を書き直すと、次のようになります。
発電量 = 電力消費量 + 正味電力輸出量 + 貯蔵量の変化 + 損失
発電とは、その場所にある発電所によって生成される電力のことですが、これについては後ほど説明します。消費量とは、ある場所での電力の使用量です。
エクスポート/送信は物理的な電力線を介して行われます。力を入れると発熱します。その熱により、途中で少しパワーが失われることになります。金属も高温になると膨張します。力を加えれば押すほど、ラインは熱くなり、より広がります。送電線が広がりすぎると、送電線が低くなりすぎて木に火が着く可能性があります。そのため、各伝送線路には「定格」が設定されています。それが、そのラインを通過できる最大の力です。
この場合の貯蔵とは、バッテリー、揚水貯蔵、圧縮空気などを指します。前述したように、電力が送電線を介して伝送されるときに損失が発生します。
最後に、電力が他の商品と異なる点が 1 つあります。石油を1000バレル買えば、

それらの石油バレルは実際に物理的に配送されます。市場で 1 メガワット時 (MWh) の電力を購入したとしても、1000 個の GPU を壁に接続して 1 時間で 1 MW の消費を開始できるとは限りません。
電力市場には独立系システムオペレーター（ISO）と呼ばれる中心的な非営利機関があるからだ。彼らは、地元の電力会社（負荷供給事業体）からのすべての入札と地元の発電所からのすべてのオファーを受け取り、各場所の効率的な市場価格に加えて、予定されている発電量と流量を計算します。その後、1 日を通して、各場所で予想される消費量に基づいて、どのくらいの電力を生成する必要があるかを各発電機に通知します。 (一部の地域では ISO が存在しませんが、平衡化当局と呼ばれる組織が電力網の安定性を確保しています。)
ISO が必要な理由は、グリッドのバランスを保つためです。グリッドから大量の電力を引き出すと、その周波数が下がります。 [10] 電力を送り込む場合はその逆が当てはまります。しかし、これらの発電機は、正確な速度で回転して適切な周波数 (60 Hz) で電力を生成する 1 億ドル以上の機械のようなものです。この回転については後ほど詳しく説明しますが、今知っておく必要があるのは、送電網の周波数が正しい周波数と異なる場合、発電機に重大な損傷を引き起こす可能性があるということだけです。
極端な場合、誰かが実際に電力を消費しすぎている場合、バランス調整権限がその人を遮断します。時々、人々は送電網を台無しにしようとさえしていないのに、主要な発電機が停止した場合や、非常に寒い冬で皆が大量の電力を消費している場合など、需要を満たすのに十分な供給が存在しない場合があります。発電機を節約するために、私たちは送電網の一部を遮断する停電や停電に頼らざるを得ません。
世代的にそう見えるかもしれない

作成者にとって、あなたは ISO の日々の気まぐれに左右され、ISO はあなたに支払う金額を正確に教えてくれます。そうは言っても、物理的な発電所を所有している場合には、電力を取引したり、物理的な発電所をヘッジしたりするさまざまな方法があり、それについては後で説明します。ほとんどのデータセンター運営者は、予測可能な経済性を確保するために、何らかの方法でヘッジを行っています。
この章の内容
多くの場合、天然ガスまたは石炭発電機が電力価格を決定します。
天然ガス発電機のエンジニアリングは、起動コストと有効熱量によって把握されるコスト構造を暗示します。
先ほども述べたように、電力価格は電力を生産するための限界費用から生じます。しかし、そのコスト曲線はどこから来たのでしょうか?発電所の紹介です。
発電所は、「ユニット」と呼ばれることもある 1 つまたは複数の発電機で構成されます。すべての発電機にはある程度の容量があり、それが生成できる最大電力量です。ほとんどの単位ではメガワット (MW) で測定され、原子力発電所などではギガワットで測定されます。米国の送電網上の 1 MW を超えるほぼすべての発電所は、エネルギー情報局 (EIA) が毎年まとめている文書である EIA-860 [11] に記載されています。 EIA-860 が必ずしも正確ではない場合もありますが、始めるには十分です。できます

[切り捨てられた]

## Original Extract

Power 2026 is a primer on AI

INDEX
01
Preface: Motivation
All About Power Plants
02
Fundamentals of Commodities Pricing
03
The Power Plant
04
To Build A Power Plant
05
Case Study: Homer City
06
Meeting the Growing Demand
How To Trade Power
07
Power Markets in the US
08
Case Study: Alberta
09
The Production Cost Model
10
Practical Approximations
11
Types of Power Trades
CALCULATING...
Power 2026
Electricity Pricing in the Age of AI
AI demand is exploding. But while there's been plenty of discussion on the GPU and memory supply chains, the real constraint to expanding AI capacity is power.
Data centers already account for ~5% of US power consumption [1] . With data center power demand doubling every two years [2] , demand would in theory outpace total US power generation by the mid-2030s.
Building a data center requires thoughtful planning around permitting, surrounding infrastructure, and anticipated power prices. In many cases, the variable cost of power determines whether or not a data center is viable at all.
I'm a former quant researcher at a major hedge fund who covered power and gas , and I've spent much of the last couple of years advising founders and investors on data center buildouts (procuring GPUs, negotiating with coal plants, identifying sites).
This primer on power and data centers is for a broader audience who senses there's an opportunity in energy markets and wants to get up to speed.
By the end of this primer, you'll understand:
Part 1: how power plants work , how data centers are developed , and how companies might respond (for founders/investors).
Part 2: how to price the US power markets (for traders).
If you're really interested in power, I'd advise you to read the whole thing. Hopefully you find it timely and interesting.
Thanks to Tina He and Samuel Spitz for editorial feedback.
Fundamentals of Commodities Pricing
In This Chapter
Commodities are governed by a system balance equation that ties together supply, demand, net exports, and change in storage.
In a competitive market, the price of a commodity equals the marginal cost to serve an additional unit of demand.
Due to the physical constraints required for grid stabilization, the power market in particular is often managed by an independent system operator (ISO).
Before we dive into power, you have to understand where it sits in the broader context. There are three major commodity categories: energy (e.g. power, crude oil, natural gas), agriculture, and metals. All commodities have some basic properties in common that set them apart from other asset classes.
Equities and commodities are different from each other in two major ways.
First, a stock doesn't have any sort of forcing function collapsing the price to its fundamental value. A commodities contract, in comparison, has a fixed settlement date. On the settlement date, someone will come and physically buy or sell that tradeable commodity. If the corn is too expensive, they're not going to buy it. Prices must collapse to reality.
Second, the underlying companies that equities represent have various relationships with each other that are sometimes hard to untangle, such as contracts, shared customers, supply chains, and so forth. Commodities feed into each other in a much more straightforward way. A wellhead drilled into the ground results in some amount of crude oil and some amount of natural gas (mostly methane). [3] Crude oil is great because it can be shipped easily all over the world. Its global market means that prices abroad impact us here in the US, hence the widespread concern around the closure of the Strait of Hormuz. That crude oil has a market price, and refineries buy it to produce distillates: petroleum products like gasoline, diesel, and jet fuel. That's the same gasoline that you buy at the gas pump.
If a well produces almost no oil, we call it a dedicated gas well, and that's where the majority of domestic natural gas comes from. Natural gas cannot be shipped in its raw form so easily, but you can cool it and ship it in the form of liquefied natural gas (LNG). The prices for natural gas abroad have lately far exceeded the prices in the US, so we export as much as we can. As a result, natural gas prices abroad don't impact us here in the US as much in the short-term [4] . In a simplified sense, the natural gas we have here in the US is the supply that couldn't be physically exported [5] . Natural gas is consumed by some industrial use cases (e.g. chemical manufacturing), residential-commercial (e.g. heating your home), and power [6] - the focus of this primer.
The commodities market has to "solve" in a particular way at each location:
Supply = Demand + Net Exports + Change in Storage [7]
Supply is the amount that the local merchants need to produce. Demand is how much is consumed at a specific location. Net exports are the units exported from a region minus whatever was imported. Change in storage is the quantity stored minus the amount drawn from storage. This equation must hold true at every location in the entire world. There are additional constraints:
You can only transfer so much supply from one point to another (only so many trucks/ships available, or a power line can only support so much transmission - more on this later)
You can't draw more from storage than you put in, and you can only store a maximum amount of supply at a given time
Each of the above variables is a function of price. Price is the amount of money that consumers pay and suppliers receive. If price increases at a location, then it becomes economical to produce more supply there (less efficient producers can now serve the demand without losing money), and the local demand goes down. Of course, it'll also decrease net exports (if there were any at all) and incentivize more supply to come out of storage (if any exists). The price will change everywhere until the system "balances."
At any given time, all of the units that are sold, whether the "first" unit sold or the "last" one, are sold at the same price. You might think that if it's really cheap for a supplier to produce goods, maybe they'd try to charge a little less. But in a competitive market, you don't personally get to "pick" the price at which you sell. This is a critical assumption for how commodities pricing works. So what determines that market price?
To produce the last unit of supply to meet demand/storage/exports, we would pick the cheapest producer that is not at full capacity. The cost of producing that unit of supply is called the "marginal cost." For example, if we need 100 barrels of oil, maybe Producer A offers 99 barrels at $5/barrel, and Producer B has the next cheapest offer at $6/barrel. The marginal cost is then $6.
In a competitive market, the market price that everyone sells at is exactly equal to that marginal cost. Note that if you're the producer who provided the last unit of supply (i.e. Producer B), you make $0 of profit on that last unit; the price is equal to your marginal cost. This property is called "marginal pricing" [8] . It will be extremely relevant for power pricing.
A simple justification for why this is true: Let's pretend like that marginal producer refuses to sell at their marginal cost. They're only willing to sell for higher. The issue is that the market is competitive. There's another supplier who's going to undercut that attempted market price. This will keep happening until the market price is exactly equal to the marginal cost of the cheapest producer that can offer the additional unit of the supply. If the price were lower, then it would be uneconomical for anyone to supply that unit.
Power is injected into or withdrawn from the grid at physical "buses" [9] . These buses are numerous and are spread out within every city in the US. They're all connected via transmission lines.
To rewrite the equation above in electricity terms for a specific location, which could be a physical power bus, a whole city, or an entire state:
Generation = Power Consumption + Net Power Exports + Change in Storage + Losses
Generation is the electricity produced by power plants at that location, which we'll dive into later. Consumption is the usage of power at a location.
Exports/transmission occurs over a physical power line. When you push power through it, it heats up. That heat means it loses a little bit of power along the way. The metal also expands when it's hot. The more power you push through, the hotter the line gets, and the more it expands. If the power line expands too much, the power line could droop too low and light a tree on fire. As a result, each transmission line comes with a "rating." That's the maximum amount of power that can be pushed through that line.
Storage, in this case, refers to batteries, pumped storage, compressed air, and so forth. Losses, as mentioned, occur when power is transmitted over a power line.
There's one last thing that makes power different from the other commodities. If I buy 1000 barrels of oil, I can literally take physical delivery of those barrels of oil. If I buy 1 megawatt-hour (MWh) of power on the market, that doesn't necessarily mean I can just go and plug 1000 GPUs into the wall and start consuming 1 MW for an hour.
That's because the power market has a central nonprofit authority called the independent system operator (ISO). They take in all the bids from local utilities (load serving entities) and all the offers from local power plants, and they compute the efficient market price at each location, plus intended generation and flows. Then throughout the day, they tell each generator how much power they need to produce based on how much consumption they anticipate at each location. (In some regions, there is no ISO, but an entity called a balancing authority still ensures the grid stays stable.)
The reason we need the ISO is to keep the grid balanced. When you draw a ton of power from the grid, you lower its frequency. [10] The opposite holds true if you're pumping power in. But these generators are like $100M+ machines that are rotating at the exact speed to produce power at the right frequency (60 Hz). We'll talk more about that rotation later, but all you need to know right now is that if the frequency of the grid differs from the correct frequency, it can cause serious damage to the generators.
In extreme cases, if someone is really consuming too much power, the balancing authority will cut them off. Sometimes, people aren't even trying to screw up the grid, but there just isn't enough supply to meet demand, like if major generators go on outage, or if it's a really cold winter and everyone is consuming a lot of power. To spare the generators, we're forced to resort to brownouts or blackouts, where parts of the grid are cut off.
It might seem like as a generator, you're subject to the day-to-day whims of the ISO, who tells you exactly how much you're going to get paid. That said, there are a variety of ways to trade power and hedge a physical plant if you own one, which we'll get into. Most data center operators hedge in some way to ensure predictable economics.
In This Chapter
Natural gas or coal generators are often the units that set the power price.
The engineering of a natural gas generator implies its cost structure, captured by its startup costs and effective heat rate.
As I mentioned, the power price comes from the marginal cost to produce power. But where does that cost curve even come from? Introducing the power plant.
A power plant comprises one or more generators, sometimes called "units." Every generator has some capacity, which is the maximum amount of power it can produce. It's measured in megawatts (MW) for most units, or gigawatts for something like a nuclear plant. Nearly every power plant on the US grid greater than 1 MW can be found in the EIA-860 [11] , a document that the Energy Information Administration (EIA) puts together annually. Sometimes the EIA-860 isn't exactly right, but it's good enough to get started. You can jus

[truncated]
