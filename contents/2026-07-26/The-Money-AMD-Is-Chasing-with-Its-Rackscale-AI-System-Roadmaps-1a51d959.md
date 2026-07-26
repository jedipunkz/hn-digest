---
source: "https://www.nextplatform.com/compute/2026/07/24/the-money-amd-is-chasing-with-its-rackscale-ai-system-roadmaps/5278510"
hn_url: "https://news.ycombinator.com/item?id=49057722"
title: "The Money AMD Is Chasing with Its Rackscale AI System Roadmaps"
article_title: "The Money AMD Is Chasing With Its Rackscale AI System Roadmaps"
author: "rbanffy"
captured_at: "2026-07-26T12:58:22Z"
capture_tool: "hn-digest"
hn_id: 49057722
score: 1
comments: 0
posted_at: "2026-07-26T12:57:17Z"
tags:
  - hacker-news
  - translated
---

# The Money AMD Is Chasing with Its Rackscale AI System Roadmaps

- HN: [49057722](https://news.ycombinator.com/item?id=49057722)
- Source: [www.nextplatform.com](https://www.nextplatform.com/compute/2026/07/24/the-money-amd-is-chasing-with-its-rackscale-ai-system-roadmaps/5278510)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T12:57:17Z

## Translation

タイトル: AMD がラックスケール AI システムのロードマップで追い求めている資金
記事のタイトル: AMD がラックスケール AI システムのロードマップで追いかけている資金
説明: 今週、AMD はシリコンバレーで現在毎年恒例となっている Advancing AI イベントを開催しました。

記事本文:
メインコンテンツへジャンプ
検索
その他のトピック
あらゆるセクションの最新ニュースをすべてお届けします
AMDがラックスケールAIシステムのロードマップで追い求めている資金
今週、AMD は現在毎年恒例となっている Advancing AI イベントをシリコンバレーで開催し、資金、ロードマップ、製品の詳細について多くの話題が交わされました。
AMDの最高経営責任者リサ・スー氏は間違いなくチップメーカーをある種の自滅的な傾向から救っただけでなく、AI業界の巨人であるNvidiaやかつてのCPUライバルであるIntelに代わる絶対的に信頼できる代替品にするためにハードウェアとソフトウェアのラインナップを拡大、深化させたように、我々もこのイベントの報道を、対象となる市場全体とそれを推進する力に関する最新情報から始めるつもりだ。
今後数日間で、私たちは「Altair」MI400アーキテクチャ、特にAMDの時流に乗ってきた多くのOEMおよびODMパートナーを通じて今年後半にAMDの「Helios」ラックスケールサーバー設計でデビューするMI455Xバリアントの詳細を明らかにする予定です。また、今年後半に予定されている次期「Venice」Epyc 9006 プロセッサに関する新しい詳細も見ていきます。次に、AMD がチャンスを公平に獲得するのに役立つ、AMD が公開したロードマップを見ていきます。
それでは、早速、お金の話をしましょう。
スー氏は冒頭の基調講演で、2020年以降、AIトレーニングのワークロードによるコンピューティングニーズは毎年5倍に増加しており、この勢いが弱まる兆しはないと述べた。これにより、規模の大小を問わず、AI モデルの構築者に一定量のコンピューティングが必要になります。しかし、スー氏によれば、大きな推進力は、比較的少数の企業での AI トレーニングから、GenA をレンタルしている企業での AI 推論への避けられない待望の移行です。

API を介してモデル化するか、AI システムを購入して、ライセンスまたはオープンソース モデルを使用して独自の推論を実行します。
後者の変化は長い間予想されており、GenAI が登場する前の数年前の予測では、最終的には推論によってトレーニングとしてのコンピューティング (したがって収益) がおそらく 3 倍か 4 倍になることが示されていました。私たちはまだそこまで到達していませんが、明らかに進んでいます。Su 氏によると、昨年 AI アクセラレーター市場では推論とトレーニングがほぼ 50 対 50 に分かれていると宣言しましたが、今年は推論が 60 パーセント、トレーニングが 40 パーセントになると予想しています。 2024 年時点では約 40 ～ 60 でした。
私たち全員が知りたい疑問は、推論が 2:1、3:1、または 4:1、さらには 10:1 になるのはいつになるかということです。それが実現すると、AI 推論、特に従来の機械学習ではなく生成型 AI 推論が、企業だけでなく、ハイパースケーラー、クラウド ビルダー、AI モデル ビルダーにとっても実稼働ワークロードとなることを意味します。 (ちなみに、Su が提示したグラフの一部を分解し、それらのグラフの基礎となるデータを再構成すると、AI トレーニングと推論の間、さらには CPU プラスでの GPU/XPU トレーニングとエージェント間の収益率を導き出すことができます。
Su 氏が披露した世界中のトークン消費曲線は、Exponential View の State of the AI Economy レポートから引用したもので、2024 年 1 月から今年 2 月までの GenAI モデルによって消費または生成されたトークンの数の月ごとの増加を示しています。
Exponential View は、2 月に月間 35,000 兆という驚異的なトークンが消化または吐き出されたと推定しており、この指数関数的な曲線を考慮すると、8 日後に 7 月が終わる時点で、月間 50,000 兆トークンに十分に達するはずです。すべての AI モデル ビルダーと

GenAI ブームに自社だけでなく自社もサービスを提供しているハイパースケーラーやクラウド ビルダーは、この曲線が上昇し続けることを必死に望んでおり、そのため、今年構築している AI システム容量の数千億ドルが正当化されます。曲線が成長し続ける唯一の方法は、マシンを増やすことですが、可能性としては、フィールド・オブ・ドリームスの逆が常にあります。つまり、あなたがそれを構築しても、それらが来ないということです。
これまでのところ、AI 処理に対する需要が供給を上回っており、そのため、CPU、GPU、DRAM メモリ、フラッシュ、スイッチ ASIC、光学コンポーネントなど、すべての主要コンポーネントの価格がすべて同様の曲線に沿って上昇しています。
AMDのトップエグゼクティブとしてのスー氏の仕事は、このような指数関数的な変化と、HPC全般、特にGenAIが世界に与える影響について、将来について楽観的になることである。AMDが追い求め、Advancing AIイベントで明らかにした、対応可能な市場の総量はこれを反映している。
「今日の AI の状況を考えると、私たちが目にする最大の変化は、何が可能になるかについて話し合わなくなったことです」と、AMD の上層部による数時間にわたるプレゼンテーションの最後に Su 氏は述べました。 「私たちは、AIがあらゆる業界や私生活のあらゆる部分に現実的かつ重要な影響を与える可能性を実際に目の当たりにしています。私はハイパフォーマンスコンピューティングが世界を信じられないほどより良い場所にできると信じてテクノロジー業界でキャリア全体を費やしてきましたが、今ほどそう信じていることはありません。AMDでは、私たちが注力しているのはテクノロジー、ロードマップ、パートナーシップの構築です。そして、3万人を超える当社のエンジニアにとって、今日ほどエキサイティングな瞬間はありません。これはまさに次のフェーズです。」 AI: これは、あらゆるツールとあらゆる機能を備えた AI に、実際に成果をもたらす機会を提供する場所です。

現実世界に意味のある影響を与えます。そして、皆さんと一緒に私たちのエコシステムを構築できることにこれ以上興奮することはないと言えます。」
Su と同様に、私も AI が大きな影響を与えると信じています。実際、すでに大きな影響を及ぼしています。その影響のベクトルの大きさについては私も同意します。しかし、それらのベクトルの方向性と、世界の文化や経済における AI 勢力の不均一な分布に関して、私はほぼ確実に意見が異なります。二次的、三次的影響が重要になり、おそらく一次的影響よりも明らかな影響が重要になります。
それらを概説し、文化への影響を予測することは、The Next Platform の仕事ではありませんが、断言させていただきますが、私たちは多くの皆さんと同じように、このことについて日夜考えています。従来の ModSim HPC ワークロードや、市場に登場した初期の機械学習に対して行ったのと同様です。ここでの私たちの仕事は、新しいテクノロジーを調査し、それが IT 市場に与える影響を確認または予測することです。それ以上でもそれ以下でもありません。
Su 氏が昨日明らかにした、改訂されたデータセンター AI アクセラレータ TAM は次のとおりです。
この予測は 2025 年 6 月の予測より 2 年長く、AI トレーニングと AI 推論を組み合わせることで、2025 年から 2030 年までの年間平均成長率が 45% 以上増加し、1 兆 4000 億ドル以上に達すると予想されています。これらは AMD 独自の内部予測であり、Su はありがたいことに AI トレーニング用の TAM と AI 推論の比率を教えてくれました。データに関してはエンドポイントしかありませんが、ピクセルをカウントして、それらのエンドポイント間の 2026 年から 2029 年のデータを導き出すことができます。
AMD によると、データセンターの CPU TAM がどのように機能するかは次のとおりです。これは、多くのエージェント AI ワークロード、つまりモデルがコードを作成したり他の種類のタスクを実行したり、GPU または XPU にフックされるときに大量の Python コードを実行したりするサンドボックスであるため重要です。

推論エンジン – GPU や XPU だけでなく、CPU 上でも実行されることになります。ちょっと見てみましょう:
ありがたいことに、Su 氏はもう一度、世界中のすべてのベンダーで予想されるデータセンター CPU の収益源を 3 つのバケットに分割しました。バック オフィス アプリケーションやデータベース、Web インフラストラクチャやさまざまな種類のデータ分析を駆動する汎用 CPU。 AI クラスター ノードのホスト プロセッサとして使用される CPU。そして、これらの新しいエージェント AI クラスターは、大量の Python コードで満たされたサンドボックスを実行します。
これは非常に便利な内訳です​​。ありがとう、リサ。
これらのエージェント AI サーバー クラスターは、他の 2 種類のマシンとはわずかに異なるパフォーマンスと容量の要件を備えており、IT 支出のサーバー CPU セグメントの活性化と、これまで見たことのない成長と収益レベルを推進することになります。 TAM のこうした黄金の期待が実際にうまくいくのであれば。
Su 氏は、私たちが活用できるもう 1 つのグラフを提供してくれました。これは、AMD がデータセンター、クライアント デバイス、組み込みデバイス全体で追いかけているコンピューティングの対象市場の合計です。
AMD によると、これらのカテゴリー全体のコンピューティング TAM の合計は、2025 年から 2030 年の間に約 40% の年平均成長率で増加し、2 兆ドルを超える見込みです。そして、ご覧のとおり、データセンターがそのドーナツに占める割合は昨年よりもさらに大きくなるでしょう。
OK、これで、Su が提示したデータを再構成し、いくつかのギャップを埋めて、分析を行うことができます。それが私の脳にとっては楽しいことなのです。 。 。 。覚えておいてください。これはコンピューティング エンジン チップからの収益に対する TAM であり、DRAM、フラッシュ、インターコネクト、スクリーン、キーボード、およびシャーシが追加された完全なシステムまたはデバイスに対する TAM ではありません。
ここにすべてをまとめて配置するかわいい小さなテーブルがあります。

それらのギャップ内の lls は、通常どおり太字の赤い斜体で表示されます。
年間複合成長率について少し誤解を招きやすいのは、2 つのエンドポイント間の直線の平均成長率を計算しているだけであるということです。その間何が起こるかは実際にはわかりません。しかし、6 年間すべての実際のデータセンター AI アクセラレータとデータセンター CPU の支出を考えると、より広範なデータをより適切に埋めることができます。これは、より広範なデータがこれら 2 つのカテゴリに大きく影響され、これら 2 つのカテゴリを合計したものよりも大きくなければならないためです。
私は、世界のコンピューティング TAM は 2026 年と 2027 年にその 40 パーセントの CAGR を上回るペースで成長し、2028 年にはそれに向けて減速し、2029 年と 2030 年には急激に減速すると考えています。データセンターの成長は今後 2 年間でさらに大きくなり、また鈍化するでしょう。
実際の AMD データからわかるように、このパターンは多少の変更を加えてデータセンター AI トレーニング アクセラレータ TAM でも表現されますが、推論を目的とした AI アクセラレータがデータセンターで最も成長する分野になるでしょう。
この表を見て、いくつかのことがすぐに思い浮かびました。まず、データセンターのコンピューティングの総量は、2025 年から 2030 年の 6 年間で 6 兆 4,500 億ドルと膨大です。クライアント デバイスのコンピューティングがデータセンターのコンピューティングよりもはるかに大きかった時代もありましたが、それは忘れてください。データセンターは、この数年間で組み込みおよびクライアントよりも合計で 8.3 倍大きくなり、2030 年だけでも 10 倍になるでしょう。 AI PCなどでもその差はますます大きくなっています。
もう 1 つ注目すべき点は、GenAI ワークロードによって CPU 支出が復活したとしても、AI アクセラレータへの支出が CPU への支出をはるかに超えることです。物理デバイスの比率は長期的には 1:1 になる可能性がありますが、データセンターの AI GPU と XPU に費やされる費用はそれを超えるでしょう

AMD の分析によると、データセンターの CPU への支出が増加しています。データセンターの CPU に対するデータセンターの GPU および XPU の比率は、2025 年には 7.9 倍で、2026 年には 6.8 倍になると予想されていますが、AMD が提供する TAM の数値では、今後 2 年間は約 6.4 倍に落ち着き、6 年間の平均は 6.5 倍になります。
その他の興味深い点について: エージェント AI CPU TAM は、2026 年には汎用 CPU TAM とほぼ同じ規模になる予定ですが、昨年はゼロでした。 2030 年までに、そのエージェント AI CPU セグメントは、データセンター内の汎用 CPU セグメントの 5.2 倍の大きさになる予定です。クレイジーですよね？
エージェント AI CPU と適切に呼べるものは何もないかもしれません。機能の組み合わせが異なるだけで、特別な機能はありません。しかし、AMD の考えが正しければ、人々は間違いなくこれらのエージェント サンドボックス用に大量の CPU を購入するでしょう。
ここからが興味深いところです。 GPU、XPU、CPU にわたる AI 推論から AI トレーニングを抽出したい場合、Su のデータを使用するとそれが可能になります。 AI トレーニング用の GPU と XPU をトレーニング システム内の AI ホスト CPU のシェアに追加すると (CPU の比率が GPU の比率と同じであると仮定)、データセンター コンピューティングのトレーニング側に 6 年間で累積支出額は 1 兆 4,200 億ドルになります。そして、あなたがそれを取った場合、

[切り捨てられた]

## Original Extract

This week, AMD held what is now an annual Advancing AI event in Silicon Valley, and there is much ta ...

Jump to main content
Search
More topics
All the latest news, from all sections
The Money AMD Is Chasing With Its Rackscale AI System Roadmaps
This week, AMD held what is now an annual Advancing AI event in Silicon Valley, and there is much talk about money, roadmaps, and product deep dives.
Like AMD chief executive officer Lisa Su, who has without a doubt not only saved the chip maker from some self-destructive tendencies but has broadened and deepened its lineup of hard and soft wares to make it an absolutely credible alternative to AI industry juggernaut Nvidia and former CPU rival Intel, we will start off our coverage of the event with an update on total addressable markets and the forces that drive them.
In the coming days, we will uncover the details of the “Altair” MI400 architecture, specifically for the MI455X variant that makes its debut in AMD’s “Helios” rackscale server designs later this year through the many OEM and ODM partners that have all jumped on the AMD bandwagon. We will also take a look at new details about the forthcoming “Venice” Epyc 9006 processors also expected later this year. And then we will take a look at the roadmaps that AMD gave some sneak peeks of, which will help them capture AMD’s fair share of the opportunities.
So, without further ado, let’s talk money.
In her opening keynote address, Su said that AI training workloads are increasing their compute needs by a factor of 5X every year since 2020, and that there is no signs of this abating. That drives a certain amount of compute for the AI model builders, large and small. But the big driver, according to Su, is the inevitable and much-anticipated shift from AI training at a relative handful of companies to AI inference with companies either renting GenAI models through APIs or buying AI systems to run their own inference with licensed or open source models.
That latter shift has been long anticipated, and the projections from several years ago, before GenAI hit, showed that inference would eventually drive maybe 3X or 4X the compute (and therefore revenues) as training. We are not quite there yet, but we are apparently on our way, according to Su, who proclaimed at last year in the AI accelerator market, inference and training were split about 50-50, but this year she expects for it to look more like 60 percent for inference and 40 percent for training. It was about 40-60 back in 2024:
The question we all want to know is when will it be 2:1, 3:1, or 4:1 for inference, or even 10:1. When that happens, that will mean that AI inference – and specially generative AI inference rather than traditional machine learning – really is a production workload for the enterprises as well as hyperscalers, cloud builders, and AI model builders. (By the way, if you tear apart some of the charts Su presented and reconstitute the data underlying those charts, you can derive a revenue ratio between AI training and inference, and even between GPU/XPU training and agentic on CPU plus
The curve for token consumption worldwide that Su showed off came from the State of the AI Economy report from Exponential View, and it shows the monthly increases in the number of tokens that are consumed or generated by GenAI models from January 2024 through February of this year:
Exponential View estimated that an astounding 35 quadrillion tokens per month were chewed up or spat out in February, and given this exponential curve it should be well into 50 quadrillion tokens per month here as we finish July in eight days. All of the AI model builders and the hyperscalers and cloud builders that are serving them as well as themselves in the GenAI boom are hoping like hell that this curve keeps rising, therefore justifying the hundreds of billions of dollars in AI system capacity they are building out this year. The only way the curve can keep growing is to have more machines, but there is always the inverse of the Field of Dreams as a possibility: You build it, and they don’t come.
So far, demand for AI processing is exceeding supply, and that is why pricing for all key components –CPUs, GPUs, DRAM memory, flash, switch ASICs, optical components – are all rising along a similar curve.
Su’s job as the top executive at AMD is to be optimistic about the future in terms of such exponentials and about the affect that HPC in general and GenAI in particular will have on the world, and the total addressable markets that AMD is chasing and divulging at the Advancing AI event reflect this.
“When I think about where AI is today, the biggest change that we see is we are no longer talking about what might be possible,” Su said at the end if several hours of presentations by the top brass at AMD. “We are actually seeing how AI can have real and significant impact across every industry and every part of our personal lives. I have spent my entire career in tech believing that high performance computing can make the world an incredibly better place, and I have never ever believed that more than I do today. At AMD, what we are focused on is building the technology, the roadmaps, and the partnerships. And there has never been a more exciting moment than today for our 30,000 plus engineers. This is really the next phase of AI. This is where we give AI the opportunity, with all the tools and all the capabilities, to really bring meaningful real-world impact. And I can tell you, I could not be more excited to build all of that together with you, our ecosystem.”
Like Su, I believe that AI will have a huge impact – has already had a huge impact, in fact. I even agree on the size of the vectors of that impact. But I almost certainly differ in the direction of those vectors and the non-uniform distribution of AI forces across the cultures and economies of the world. Secondary and tertiary impacts will matter, and perhaps more than primary ones that are more obvious.
Outlining those and predicting impacts on cultures is not the job of The Next Platform , but let me assure you, we think about this day and night as many of you do. As we did for traditional ModSim HPC workloads or early machine learning as they emerged in the market. Our job here is to examine new technologies and to ascertain or predict their effects on the IT market. No more, no less.
Here is the revised datacenter AI accelerator TAM that Su divulged yesterday:
It goes out two more years than the June 2025 forecast, and the expectation is that AI training and AI inference together will drive more than a 45 percent compound annual growth rate between 2025 and 2030 to hit more than $1.4 trillion dollars. These are AMD’s own internal projections, and Su thankfully gave us the ratio of the TAM for AI training split from AI inference. We only have the endpoints in terms of data, but we can count pixels to derive the data for 2026 through 2029 in between those endpoints.
Here is how the datacenter CPU TAM will play out according to AMD, which is important because a lot of agentic AI workloads – sandboxes where models are creating code or performing other kinds of tasks, running a lot of Python code as they are hooked to GPU or XPU inference engines – are going to be running on CPUs, not just GPUs or XPUs. Take a gander:
Once again, Su thankfully broke the expected datacenter CPU revenue streams across all vendors around the world down into three buckets: The general purpose CPUs that drive back office applications and databases as well as Web infrastructure and various kinds of data analytics; the CPUs that are used as host processors for AI cluster nodes; and these new agentic AI clusters running those sandboxes full of churning Python code.
This is a very handy breakdown, and thank you, Lisa.
These agentic AI server clusters, which have slightly different performance and capacity requirements from the other two types of machines, are going to drive a revitalization in the server CPU segment of IT spending, and growth and revenue levels that we have never, ever seen before. Provided these golden TAM expectations actually pan out.
Su gave us one more chart we can work with, which is the total addressable market for compute that AMD is chasing across the datacenter, client devices, and embedded devices:
The total compute TAM across these categories, according to AMD, is going to grow at around a 40 percent compound annual growth rate between 2025 and 2030 to reach more than $2 trillion, and as you can see, the datacenter is going to be an even larger share of that donut than it was last year.
OK, with all of that, I can reconstitute the data Su presented, fill in some gaps, and then do some analysis. Which is what passes for fun for my brain. . . . And remember: This is the TAM for the revenue from compute engine chips, not for complete systems or devices with their DRAM, flash, interconnect, screens, keyboard, and chasses added in.
Here is a pretty little table that brings it all together and fills in those gaps, which are shown in bold red italics as usual:
The thing about compound annual growth rates that are a little misleading is that they just calculate the straight line average growth between two endpoints. You can’t really see what happens in between. But given the actual datacenter AI accelerator and datacenter CPU spending for all six years allows us to better fill in the broader data because that broader data is affected by these two categories greatly and must be larger than these two categories added together.
I think that the worldwide compute TAM is going to grow faster than that 40 percent CAGR in 2026 and 2027, and slow down towards it in 2028, and then slow radically in 2029 and 2030. Datacenter growth is going to be even higher in the next two years, and will also peter out.
This pattern, with a few wiggles, will also be expressed in the datacenter AI training accelerator TAM, as you can see from the actual AMD data, but AI accelerators aimed at inference are going to be the highest growth area in the datacenter.
Some things immediately jump out to me in this table. First, the total amount of datacenter compute is just enormous, at $6.45 trillion across the six years from 2025 through 2030. There was a time when client device compute was much larger than datacenter compute, but forget that. Datacenter is 8.3X larger in aggregate across those years than embedded and client, and will be 10X in 2030 alone. The gap is getting bigger, even with AI PCs and such.
The other thing to note is how even with a resurgence in CPU spending driven by GenAI workloads, spending on AI accelerators is going to far exceed that on CPUs. The ratio of physical devices might be 1:1 in the long run, but the money spent on datacenter AI GPUs and XPUs is going to exceed spending on datacenter CPUs, according to AMD’s analysis. The ratio of datacenter GPUs and XPUs to datacenter CPUs was 7.9X in 2025, and is expected to be 6.8X in 2026, but it settles down to around 6.4X for the next couple of years and averages 6.5X over the six years in the TAM figures supplied by AMD.
On other interesting thing: The agentic AI CPU TAM will be almost as large as the general purpose CPU TAM in 2026, and was zero last year. By 2030, that agentic AI CPU segment will be 5.2X larger than that general purpose CPU segment in the datacenter. Crazy, isn’t it?
There may be nothing that can properly be called agentic AI CPUs – they just have a different mix of features, but no special features – but people are sure going to buy a lot of CPUs for these agentic sandboxes if AMD is right.
Now here is where it gets interesting. If you want to try to extract AI training from AI inference across GPUs, XPUs, and CPUs, Su’s data lets you do that. If you take the AI training GPUs and XPUs and add it to the share of AI hosts CPUs that are in training systems (assuming the ratio for CPUs is the same as for GPUs), you get a cumulative spending of $1.42 trillion for the six years for the training side of datacenter compute. And if you take the

[truncated]
