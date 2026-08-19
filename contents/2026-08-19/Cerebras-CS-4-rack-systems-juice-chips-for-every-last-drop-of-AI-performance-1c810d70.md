---
source: "https://www.theregister.com/systems/2026/08/19/cerebras-cs-4-rack-systems-juice-chips-for-every-last-drop-of-ai-performance/5289286"
hn_url: "https://news.ycombinator.com/item?id=49354816"
title: "Cerebras CS-4 rack systems juice chips for every last drop of AI performance"
article_title: "Cerebras CS-4 rack systems juice chips for every last drop of AI performance"
image: "https://image.theregister.com/5289328.jpg?imageId=5289328&x=0&y=0&cropw=100&croph=100&panox=0&panoy=0&panow=100&panoh=100&width=1200&height=683"
author: "joebuckwilliams"
captured_at: "2026-08-19T00:39:26Z"
capture_tool: "hn-digest"
hn_id: 49354816
score: 2
comments: 0
posted_at: "2026-08-19T00:14:25Z"
tags:
  - hacker-news
  - translated
---

# Cerebras CS-4 rack systems juice chips for every last drop of AI performance

- HN: [49354816](https://news.ycombinator.com/item?id=49354816)
- Source: [www.theregister.com](https://www.theregister.com/systems/2026/08/19/cerebras-cs-4-rack-systems-juice-chips-for-every-last-drop-of-ai-performance/5289286)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T00:14:25Z

## Translation

タイトル: Cerebras CS-4 ラック システム、AI パフォーマンスの最後の一滴までジュース チップを提供
説明: 次世代システムは、ラックに 3 倍のチップを詰め込みながら、チップあたりのパフォーマンスを 2 倍にします

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
2026 年クラウド インフラストラクチャ月間
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
Cerebras CS-4 ラック システムは、AI パフォーマンスの最後の一滴までジュース チップを提供します
次世代システムは、ラックに 3 倍のチップを詰め込みながら、チップあたりのパフォーマンスを 2 倍にします
トビアス・マン
トビアス
マン
システムエディター
発行済み
2026年8月19日水 // 01:00 UTC
高速な AI 推論を求めている場合、メモリ帯域幅が克服すべきボトルネックになります。 Cerebras のディナープレートサイズの AI アクセラレータは、気の遠くなるような 21.6 ペタバイト/秒 (PB/s) のメモリ帯域幅を備えており、すでに Nvidia や AMD の最高の GPU よりも 1,000 倍高速でした。
チップの新参者である同社は、火曜日に次世代のウェーハスケールエンジン（WSE）とNexusラックシステムを発表した。 Cerebras は、ワットあたりのスループットを前世代の 10 倍に高めることで、そのリードを拡大することを目指しています。
Cerebras はいくつかの方法でこれを実現します。しかし、私たちが知ることができることから、主なレバーは、そのチップをあらゆるヘルツで絞り出すことから来ています。新しく発表された WSE-3T (ここでの「T」は「Turbo」の略) は、2 年前の WSE-3 の 2 倍のコンピューティング、メモリ ファブリック、および I/O 帯域幅を約束します。
しかし、以下のグラフを見ると、同じプロセス技術、ウェーハ面積サイズ、トランジスタ数、コア数、SRAM 容量を使用してこれを達成していることがわかります。それは、WSE-3T が新しいシリコンではないためです。その代わりに、Cerebras は既存のウェーハスケールエンジンをさらに強化しているだけだと語った。
今回の主な革新は電力供給に関連しているようで、明らかに非常に効率的であるため、チップに2倍の電力を送り込むことができ、「より高い動作周波数が可能になり、

d より高速なトークン生成。」
クロックはどのくらい高くなりますか?私たちの推定によると、Cerebras は現在、前世代の 1.4 GHz から 2.8 GHz でシリコンを実行していますが、これはかなりの成果でしょう。
いずれにせよ、各 WSE-3T は 250 ペタ FLOPS の AI コンピューティング、44 GB の SRAM (これはタイプミスではありません。実際にはそれだけ多くの SRAM が搭載されています)、43.2 PB/s のメモリ帯域幅、および 2.4 Tbps のオフダイ接続に優れています。
紙の上では、実際よりも印象的に聞こえます。 AMD と Nvidia の最新 GPU は、高密度 FP16 コンピューティングで 4 ～ 5 ペタ FLOPS、FP4 では 35 ～ 50 ペタ FLOPS を提供します。 Cerebras のヘッドラインのパフォーマンス数値はスパース性に大きく依存しており、一般的に LLM 推論には役に立ちません。
WSE-3 で見たのと同じ 10 倍のスパース性を仮定すると、WSE-3T の高密度 FP16 パフォーマンスは 25 ペタ FLOPS に近づくはずです。これは依然として印象的ですが、チップメーカーが信じさせているほど印象的ではありません。
また、WSE-3T のピーク メモリ帯域幅は純粋に理論上のものであると考えられます。 LLM 推論中、WSE-3 には SRAM を自力で飽和させるのに必要な計算能力が不足しており、Turbo バリアントが何か違うと信じる理由はありません。
ただし、今回 Cerebras は推論スタック全体を独自のアクセラレータで実行しようとしているわけではありません。代わりに、アマゾン ウェブ サービス (AWS) および AMD と提携して、推論パイプラインの計算集約型プロンプト処理ビットをそれぞれの Trainium XPU および Instinct GPU にオフロードしました。
少なくとも推論すると、Cerebras のチップは現在、Nvidia が Groq を使用しているのと同様に、主にデコード アクセラレータとして機能しています。Elon Musk の Grok モデル ファミリと混同しないでください - LPX ラック システムの LPU です。
Cerebras にとっての主な利点は、そのチップに大量の SRAM が搭載されていることです。つまり、2,000L 必要ではなく

Cerebras では、重みが保存される精度に応じて、数兆個のパラメータ モデルを実行する PU を数十個使用するだけで済みます。
奇妙なことに、Cerebras は、5 年前の WSE-2 の発売以来、大幅に増加していない SRAM 容量を増やすのではなく、今世代でパフォーマンスを 2 倍にすることを選択しました。
プリフィルが GPU によって処理される分離された推論環境では、Cerebras がコンピューティングよりも SRAM 容量を優先することが予想されていました。ただし、これらの細分化されたコンピューティング アーキテクチャが比較的新しい現象であることを考えると、Cerebras はすでに本番環境で方向転換するにはかなり進んでいた可能性があります。
これは、Turbo の命名規則を説明している可能性があります。 Cerebras が今後もこの道を進むつもりであれば、おそらく今後も登場する WSE-4 は、SRAM 容量を約 2 倍にしながら、FP16 ではわずかなパフォーマンス向上しか提供しないと予想されます。私たちの兄弟サイトである The Next Platform では、WSE-4 がどのようなものになるかについての予測をいくつか作成していますので、ご興味があればご覧ください。
Cerebras の最新世代のウェーハ スケール アクセラレータでは、同社がラックスケール コンピューティング アーキテクチャにも真剣に取り組んでいることがわかります。
Nvidia の NVL72 ラックや AMD の Helios ラックと同様に、Cerebras の CS-4 は、モノリシック システムから、コンピューティング、電力供給、およびケーブル配線を分離して導入、メンテナンス、アップグレードを容易にするモジュラー アーキテクチャへの飛躍を遂げています。
同社のチップは現在、同社が「バックパック」フォームファクターと呼ぶもの、つまりすべての制御電子機器を搭載した一種の自己完結型システムに収容されている。
各 CS-4 にはこれらのバックパックを最大 3 つまで装備でき、その名前が示すように、ラックの背面に差し込みます。一方、ラックの前面は獣に餌を与えるために使用される電源シェルフ専用です。ケーブル配線はおそらくシステムの中心を通っていると思われます。私たちは

おそらく OCP に準拠した設計ではないと推測しますが、チップが豚肉ボードのサイズである場合、そのようなことは決してありません。
システムあたりのアクセラレータが 3 倍になると、消費電力も当然のことながら増加します。 Cerebras は、ラックがどのくらいの電力を吸い戻すかについては明らかにしていないが、より効率的な電力供給により、チップに 2 倍のワットを送り込むことができると述べている。 WSE-3 は、ウェハー レベルで 15 kW、システム レベルで約 23 kW のホットチップでした。これは、CS-4 バックパックごとにおそらく約 46 kW、合計システム電力が 120 kW ～ 140 kW になることを意味します。
数年前であれば、水冷マシンであっても、これは途方もないパワーでした。現在、このマシンは、今年後半に AMD と Nvidia から発売される 240 ～ 250 kW のラック システムに比べて、かなり保守的なものに見えます。
人類が別の惑星の表面に触れてから 60 年
OracleとOpenAIのテキサス・スターゲート・データセンター拡張は計画通りに進んでいないと報じられている
アナリストらは、雇用統計の低迷をまだAIのせいにしてはいけないと言う
米国の州法によりオペレーティング システムに年齢チェックが組み込まれる
スイッチを廃止することで遅延をなくす
ラックあたり 132 GB の SRAM メモリを備えていても、適切なサイズのモデルを実行するには多数のラックが必要になるため、I/O が重要な考慮事項となります。
たまたま、新しいラックとそれらを駆動するシリコンが、この点で大幅にアップグレードされています。
各チップには、1.2 Tbps から 2.4 Tbps のチップ間帯域幅が装備されています。しかしおそらく、より重要な改善は遅延の改善であり、いくつかの巧妙なトリックのおかげで、遅延は 5 マイクロ秒からわずか 2 マイクロ秒に短縮されました。
それは、帯域幅を大量に消費するテンソルと専門的な並列処理に大きく依存してコンピューティング能力と効率性を高める GPU とは異なるためです。

多くの帯域幅を備えているため、Cerebras のチップはすでに非常に高速なので、パイプライン並列処理を使用しても問題なく使用できます。
パイプラインの並列処理は、マルチ アクセラレータの推論と同じくらい簡単です。モデルの重みは各アクセラレータに分散され、作業は 1 チップずつ順番に実行されるため、パイプライン並列処理と呼ばれます。インターコネクトの帯域幅はそれほど問題ではありませんが、パイプラインの並列処理は、接続の遅延が非常に低いことから恩恵を受けます。
Cerebras は非常に簡単な方法で相互接続の遅延を短縮できます。余分なスイッチをすべて取り除き、チップが相互に通信するようにするだけです。
これは、多くの AI チップ設計者が独自のラックスケール設計で行ってきたこととはかなり大きく異なります。たとえば、AWS はチップツーチップ メッシュを廃止し、Trainium3 アクセラレータのスイッチド ファブリックを採用しました。これについては、昨年末に詳しく調査しました。
Cerebras によれば、2D トーラスを使用しているとのことです。これは、端が反対側に回り込むグリッドと考えることができます。同社によれば、このトポロジは最大 50 兆パラメータのサイズのモデルをサポートできるが、我々の知る限りそのようなモデルは現時点では存在しない。
そして、存在するモデルの場合、そのパーツはかなりのパフォーマンスを発揮するようで、単一の CS-4 システム上の gpt-oss-120b ではユーザーあたり最大 4,400 tok/s の速度を達成します。これに対し、Artificial Analysis のベンチマーク専門家によると、現在最速の GPU ベースの推論サービスでは約 350 tok/s です。
Cerebras の新しいメッシュ トポロジは必須ではありません。スイッチド ファブリックを使用してチップを接続したい場合は、それを妨げるものは何もありません。基本的にすべての最新チップと同様に、RDMA over converged Ethernet (RoCE) もサポートしています。妥協策としては、レイテンシが少し高くなる可能性があります。
チップ新興企業は、最初の CS-4 ベースのシステムが稼働することを期待しています

今四半期後半に。 ®
システム
Cerebras CS-4 ラック システムは、AI パフォーマンスの最後の一滴までジュース チップを提供します
次世代システムは、ラックに 3 倍のチップを詰め込みながら、チップあたりのパフォーマンスを 2 倍にします
OpenAI はセキュリティを強化するため、一部のワークロードでオーバーヘッドが 20% 増加します
多段階の思考連鎖監視の拡張により、フロンティアモデルの作業コストが高くなる
プラットフォーム エンジニアリング 2.0: プラットフォームは別の時代に構築されました。 AIが暴露しただけ
パートナーのコンテンツ: プラットフォーム エンジニアリングが議論に勝ちました。今では、AI 時代に向けて急速に成長し、進化する必要があります。
期限切れのクレジットカードが研究者によって復活し、不正な支払いが可能に
有効期限チェックにギャップがあると、死んだプラスチックが再び購入される可能性があります
コラムニスト
嫌いになってもいい、AI はここに残る
良いニュースは？テクノロジー大手がすべてをコントロールするという最悪の事態は間もなく終わるかもしれない
メタアプリと Google モバイルアプリがユーザーデータを大量に消費: 研究
ザック帝国はアップルやマイクロソフトの3倍の食欲を誇っている一方、グーグルはリーダーボードを独占していると開発者調査が示した
AIとml
Google、墜落した航空会社スピリットのデータをオークションで購入、理由はAI
AIとML
Excel のコパイロット機能がごみ箱に向かう
デボプス
リポジトリのダウンロードのエラー率が 50% に達するため、GitHub に問題が発生
オフビート
レゴの超巨大ハッブルはもう少し輝くべきだ
aiとml
Anthropicは、テキスト透かしスキームは重要でない単語に依存していると述べています
システム
部品が合わなかったため、技術者はドライバーを取り出しました。するとマザーボードから何かが飛んできた
次世代システムは、ラックに 3 倍のチップを詰め込みながら、チップあたりのパフォーマンスを 2 倍にします
AI の推論エンジンをソーシャル エンジニアリングする方法
企業がモデルのオーケストレーションに苦戦する中、AI ゲートウェイは有望に見える
「あなたを夏の午後に例えてみませんか」というようなことを言うと、他の人も採用する可能性が高いと思われます
中国のAI

米国の研究所が防御を行う一方で、研究所は前進を続ける
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を持っているのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
FOSS は Microsoft の独占を 1 つ打ち破りました。 20年間の失敗を経て、次の失敗を打ち破る時が来た
一言
GNOME は Windows のように見えることができ、Flashback は拡張機能なしで実行できます
新しい「シンプルタスクバー」はオプションですが、よりシンプルで安定した方法があります
x86-32 での Debian の最終リリースに黙祷を捧げてください
新しい Debian バージョンが 13.6 および 12.15 の形で FOSSland に登場
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
リンの次のバージョン

[切り捨てられた]

## Original Extract

Next-gen systems double per-chip performance while cramming 3x as many into a rack

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Cerebras CS-4 rack systems juice chips for every last drop of AI performance
Next-gen systems double per-chip performance while cramming 3x as many into a rack
Tobias Mann
Tobias
Mann
SYSTEMS EDITOR
Published
wed 19 Aug 2026 // 01:00 UTC
If high-speed AI inference is what you’re after, memory bandwidth is the bottleneck to beat. At a mind-numbing 21.6 petabytes per second (PB/s) of memory bandwidth, Cerebras' dinner-plate-sized AI accelerators were already 1,000x faster than Nvidia's or AMD’s best GPUs.
The chip newcomer unveiled its next-gen Wafer Scale Engine (WSE) and Nexus rack systems on Tuesday. Cerebras aims to extend that lead by boosting throughput per watt tenfold over the previous generation.
Cerebras accomplishes this in a couple of ways. But, from what we can tell, the primary lever comes from squeezing its chips for every hertz they’ve got. The newly announced WSE-3T — the “T” here stands for “Turbo” — promises twice the compute, memory fabric, and I/O bandwidth of the now two-year-old WSE-3 .
Yet, if you look at the chart below, you’ll notice it accomplishes this using the same process tech, wafer area size, transistor count, core count, and SRAM capacity. That's because the WSE-3T isn't new silicon. Instead, Cerebras tells us it's just pushing its existing wafer scale engine harder.
The main innovation this time around seems to be related to power delivery, which is apparently so efficient that they’re able to push twice the power through the chip, which “enables higher operating frequencies and faster token generation.”
How much higher does it clock? By our estimate, Cerebras is now running the silicon at 2.8 GHz, up from 1.4 GHz last gen, which would be quite the accomplishment.
In any case, each WSE-3T boasts 250 petaFLOPS of AI compute, 44 GB of SRAM (that’s not a typo, there really is that much SRAM on there), good for 43.2 PB/s of memory bandwidth, and 2.4 Tbps of off-die connectivity.
On paper that sounds more impressive than it really is. AMD and Nvidia’s latest GPUs offer 4 to 5 petaFLOPS of dense FP16 compute or 35 to 50 petaFLOPS at FP4. Cerebras’ headline performance figure relies heavily on sparsity, which as a general rule doesn't benefit LLM inference.
Assuming the same 10x sparsity we saw with the WSE-3, the WSE-3T’s dense FP16 performance should be closer to 25 petaFLOPS, which is still impressive, just not as impressive as the chipmaker would have you believe.
We also suspect the WSE-3T’s peak memory bandwidth is purely theoretical. During LLM inference, the WSE-3 lacked the compute necessary to saturate its SRAM on its own, and we have no reason to believe the Turbo variant will be any different.
However, this time around Cerebras isn’t trying to run the entire inference stack on its own accelerators. Instead, it has partnered with Amazon Web Services (AWS) and AMD to offload the compute-intensive prompt processing bits of the inference pipeline onto their respective Trainium XPUs and Instinct GPUs.
At least for inference, Cerebras’ chips now function primarily as decode accelerators, similar to how Nvidia is using Groq — not to be confused with Elon Musk’s Grok family of models — LPUs in its LPX rack systems .
The major benefit for Cerebras is its chips have a whack ton of SRAM on board. So, instead of needing 2,000 LPUs to run a trillion-parameter model, Cerebras can get away with using a few dozen, depending on the precision at which the weights are stored.
Curiously, Cerebras opted to double performance this generation rather than boost SRAM capacity, which hasn’t increased meaningfully since the WSE-2 launched five years ago.
In a disaggregated inference environment where prefill is handled by GPUs, we’d have expected to see Cerebras prioritize SRAM capacity over compute. However, given that these disaggregated compute architectures are a relatively new phenomenon, it’s possible Cerebras was already too far along in production to pivot.
This likely explains the Turbo naming convention. If Cerebras plans to continue down this path, we expect the WSE-4, which is presumably still coming, to offer only modest performance gains at FP16 while roughly doubling SRAM capacity. Our sibling site The Next Platform has drawn up some predictions of what the WSE-4 might look like if you’re interested.
Cerebras' latest generation of wafer scale accelerators also sees the company get serious about rack-scale compute architectures.
Much like Nvidia’s NVL72 and AMD’s Helios racks, Cerebras’ CS-4 makes the leap from a monolithic system to a modular architecture that breaks out compute, power delivery, and cabling for easier deployment, maintenance, and upgrades.
The company’s chips are now housed in what it calls a “backpack” form factor, a sort of self-contained system with all the control electronics on board.
Each CS-4 can be equipped with up to three of these backpacks, which, as their name suggests, plug into the back of the rack, while the front of the rack is dedicated to the power shelves used to feed the beast. Cabling presumably runs down the center of the system. We're guessing it's probably not an OCP-compliant design, but when your chip is the size of a charcuterie board, it never was going to be.
With 3x the accelerators per system, power consumption has unsurprisingly increased. Cerebras hasn’t said how much power the racks will suck back, but it has said that its more efficient power delivery means it can push twice as many watts through the chip. The WSE-3 was already a hot chip at 15 kW at the wafer level and around 23 kW at the system level. This means we’re probably looking at around 46 kW for each CS-4 backpack and a total system power of between 120 kW and 140 kW.
A few years ago, that’d have been a monstrous amount of power, even for a liquid-cooled machine. Today, the machine looks positively conservative next to the 240 to 250 kW rack systems coming from AMD and Nvidia later this year.
60 years since humanity touched the surface of another planet
Oracle and OpenAI's Texas Stargate datacenter expansion reportedly on the skids
Don’t blame AI yet for poor jobs numbers, analysts say
US state laws push age checks into the operating system
Killing latency by ditching the switch
Even with 132 GB of SRAM memory per rack, you’re still going to need a lot of racks to run any reasonably sized model, which means I/O is a major consideration.
It just so happens that the new racks and silicon that power them have gotten some beefy upgrades in this respect.
Each chip is equipped with 2.4 Tbps of chip-to-chip bandwidth, up from 1.2 Tbps. But arguably the more important improvement is to latency, which thanks to some clever tricks has been cut from five microseconds down to just two.
That’s because unlike GPUs, which rely heavily on bandwidth-intensive tensor and expert parallelism to multiply their compute and effective memory bandwidth, Cerebras' chips are already so fast they can get away with using pipeline parallelism.
Pipeline parallelism is about as simple as multi-accelerator inference gets. Model weights are distributed across each accelerator and work is performed sequentially, one chip after another, hence the name pipeline parallelism. While interconnect bandwidth isn’t as much of an issue, pipeline parallelism does benefit from very low latency connections.
Cerebras is able to lower its interconnect latency in a pretty simple way: Get rid of all the extra switches and just have the chips talk to one another.
This is a pretty big departure from what a lot of AI chip designers have been doing with their own rack-scale designs. AWS for example ditched its chip-to-chip mesh in favor of a switched fabric in its Trainium3 accelerators, which we looked at in detail late last year.
Cerebras tells us it's using a 2D torus, which you can think of as a grid where the ends wrap around to the other side. The topology, the company says, can support models up to 50 trillion parameters in size, though no such model currently exists to our knowledge.
And for models that do exist, it appears the parts will be quite performant, achieving speeds of up to 4,400 tok/s per user in gpt-oss-120b on a single CS-4 system, compared to around 350 tok/s on the fastest GPU-based inference service today, according to the benchmarking gurus at Artificial Analysis.
Cerebras' new mesh topology isn’t mandatory. If you wanted to connect the chips using a switched fabric, there’s nothing stopping you. Like basically every modern chip, it supports RDMA over converged Ethernet (RoCE), too. The compromise is your latencies may be a bit higher.
The chip upstart expects the first CS-4-based systems to come online later this quarter. ®
SYSTEMS
Cerebras CS-4 rack systems juice chips for every last drop of AI performance
Next-gen systems double per-chip performance while cramming 3x as many into a rack
OpenAI's overhead will rise 20 percent for some workloads as it hardens security
Expanded multistage chain of thought monitoring makes frontier model work more expensive
Platform Engineering 2.0: your platform was built for a different era. AI just exposed it
PARTNER CONTENT: Platform engineering won the argument. Now it has to grow up fast and evolve for the AI era.
Expired credit cards revived by researchers to make unauthorized payments
Gaps in expiry checks could let dead plastic make purchases again
COLUMNISTS
Be a hater all you want, AI's here to stay
The good news? One of the worst bits, tech giants controlling it all, might soon be over
Meta and Google mobile apps gorge on user data: Study
Zuck's empire declares triple the appetite of Apple or Microsoft, while Google packs the leaderboard, dev survey shows
AI and ml
Google buys crashed airline Spirit’s data at auction, because AI
AI and Ml
Excel's Copilot function is headed for the Recycle Bin
DEVOPS
GitHub has Issues as repo downloads hit 50% error rate
OFFBEAT
Lego's supersized Hubble deserves a little more shine
ai and ml
Anthropic says text watermarking scheme relies on inconsequential words
systems
Part didn't fit so techie got out his screwdriver. Then something flew off the motherboard
Next-gen systems double per-chip performance while cramming 3x as many into a rack
How to social engineer an AI's reasoning engine
AI gateways look promising as companies struggle with model orchestration
'Shall I compare thee to a summer's afternoon' is the sort of thing this will make, and others look likely to adopt it
Chinese AI labs keep moving forward while US labs play defense
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
FOSS smashed one Microsoft monopoly. After 20 years of failure, it's time to smash another
Word up
GNOME can look like Windows – and Flashback can do it without extensions
New 'Simple-taskbar' is an option, but there's a simpler, stabler way
A moment of silence, please, for the final release of Debian on x86-32
New Debian versions hit FOSSland in the form of 13.6 and 12.15
Baddies caught exploiting extensions bugs with perfect 10 scores on vulnerable Joomla websites
Flaws in iCagenda, Balbooa Forms extensions can impact open source CMS that powers a million sites worldwide
Frame: A new X11 server – implemented directly in assembly
Joins yserver, Phoenix, and of course XLibre – and outlier Arcan
Cinnamon 6.8 will support Wayland – if you want it
Next version of Lin

[truncated]
