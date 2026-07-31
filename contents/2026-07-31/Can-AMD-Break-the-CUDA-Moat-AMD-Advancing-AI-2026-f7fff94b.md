---
source: "https://newsletter.semianalysis.com/p/can-amd-break-the-cuda-moat-amd-advancing"
hn_url: "https://news.ycombinator.com/item?id=49128814"
title: "Can AMD Break the CUDA Moat? AMD Advancing AI 2026"
article_title: "Can AMD break the CUDA Moat? AMD Advancing AI 2026"
author: "rbanffy"
captured_at: "2026-07-31T21:57:59Z"
capture_tool: "hn-digest"
hn_id: 49128814
score: 2
comments: 0
posted_at: "2026-07-31T21:39:05Z"
tags:
  - hacker-news
  - translated
---

# Can AMD Break the CUDA Moat? AMD Advancing AI 2026

- HN: [49128814](https://news.ycombinator.com/item?id=49128814)
- Source: [newsletter.semianalysis.com](https://newsletter.semianalysis.com/p/can-amd-break-the-cuda-moat-amd-advancing)
- Score: 2
- Comments: 0
- Posted: 2026-07-31T21:39:05Z

## Translation

タイトル: AMDはCUDAの堀を突破できるか? AMD の AI 2026 の進歩
記事タイトル: AMDはCUDAの堀を突破できるか? AMD の AI 2026 の進歩
説明: エージェント カーネルの生成、ソフトウェア品質の向上、不安定な内部開発クラスター、Helios MI455X プロダクション ランプ ヘル、ファイナンス エンジニアリングからの最大 105% 割引

記事本文:
AMDはCUDAの堀を打ち破ることができるでしょうか? AMD の AI 2026 の進歩
AMDはCUDAの堀を打ち破ることができるのか? AMD の AI 2026 の進歩
OpenAI、エージェントティック カーネル生成、ソフトウェア品質の向上、不安定な内部開発クラスタ、Helios MI455X プロダクション ランプ ヘルに対する最大 105% の株式リベート割引
184 4 13 シェア 私たちが最初の AMD ソフトウェア記事を公開したとき、AMD が AI アクセラレータにおいて Nvidia との差を縮める可能性は 0% と与えました。ソフトウェアは壊れており、進歩は刺激的ではありませんでしたが、私たちは数十人の AMD エンジニアがバグ レポートの優先順位を付け、何ヶ月にもわたってトップのバグ提出者でした。
6 か月後、AMD 2.0 の記事で、私たちは 0% の確率から、より意味のある成功の確率にアップグレードするという非合意の立場をとりました。私たちがこの意見を発表したのは、AMD に対する感情がどん底にあり、市場のほとんどが私たちが AMD に対してあまりにも楽観的すぎると考えていたときでした。
この見解は、AMD が委員会形式のリーダーシップに悩まされるのではなく、変化を生み出すことができるリーダーシップを持っているという私たちの観察に基づいています。リサはすぐに私たちとの電話に飛びつき、それ以来私たちの提案の多くを実行してくれました。私たちは、最終的にソフトウェアの重要性を認識し、目的地はまだ遠いにもかかわらず、危機感を持って正しい方向に進んでいる AMD の姿を目にしました。それ以来、信号はさらに強くなりました。
今年のAMDソフトウェアスタックに関する経験に基づいて、AMDが以下に概説する2つの主要なリスクを解決する限り、成功の可能性はゼロではないということから、現在では大きな成功の可能性があるとの見方を再度更新します。 AMDが市場シェアを獲得したからといって、Nvidiaの業績が悪くなるわけではないことを強調することが重要です。すべての人のパイは急速に拡大しており、Nvidia は今後も収益を大幅に成長させていきます

うえ。 AMDはソフトウェア面でNvidiaとの潜在的な競争相手となっており、ジェンセン氏がNvidiaの迅速な行動とリードの防衛を望むのであれば、官僚主義を削減し、最も単純なタスクであっても必要な社内関係者のさまざまな承認の階層を平坦化する必要があるだろう。
Anthropic は 2GW の AMD チップを導入すると公表しており、Lisa Su とチームはエージェント指向のエンジニアリング文化に傾倒しています。 Anthropic のコンピューティング責任者である Tom Brown 氏は、その好例として、週末に「/goal」を指定して Claude を使用し、AMD ハードウェア上の内部クロード推論スタックを起動した方法を説明しました。 AMD のコンパイラとそのカーネルのほとんどはオープンソースであるため、後で説明する 1 つの大きなリスクを除けば、エージェントの時代に適した位置にあると私たちは信じています。 3 か月前、当社のアクセラレータ モデルでは、Anthropic が AMD の顧客になると発表されました。私たちは 1 週間前にソーシャルメディアでもこのことを指摘しました。
2023年、Microsoftは信頼性の低いSamsung 2023 HBMメモリと低品質のソフトウェア品質を理由にMI300Xの後にAMDを廃止し、その後MI325XとMI355Xの両方をスキップしました。 Microsoft はその後態度を変え、MI455X Helios を導入すると発表しました。私たちは、OpenAI が Azure の MI455X ラックの主要なエンド カスタマーになると考えています。 Nvidia-Groq 契約とやや似た戦略で、AMD は超高速インタラクティブ推論のための PD disagg を行う Cerebras との契約を発表しています。
AMD が防御する必要がある 2 つの主要なリスクがあります。
サプライ チェーンのチェックとエンジニアリングの第一原理分析と理解により、AMD の最初の AI ラック スケール システムである Helios は、Rubin Oberon ラックが現在採用しているケーブルレス トレイ設計を使用していないことを考慮すると、現在ラックの生産がゆっくりと進んでいることがわかります。さらに、以来、

AMD は SerDes 設計が弱く、バックプレーンの最大 85% をリタイミングする必要があり、ラックごとに 550 以上の Broadcom イーサネット リタイマーが必要です。さらに、ラック生産増加の地獄の中でバックプレーンの信頼性の課題に直面しています。
社内のほとんどの AMD エンジニアの主な不満は、社内ソフトウェア開発チーム用の安定した GPU クラスターと、自動テスト CI 用の安定した GPU クラスターが引き続き不足していることです。これは、AMD の進歩の速度を妨げており、各 AI エージェントには GPU も必要であり、テスト ツールの使用ループも必要であるため、AMD が AI コーディング エージェントの潜在的な利点を活用することを妨げています。以下の推奨事項セクションで詳しく説明します。
AMD がこれらの課題を克服できれば、AMD は好成績を収めて市場シェアを獲得できる有利な立場にあると強く信じています。これは、AMDが巧妙な金融エンジニアリングを使用してMetaとOpenAIに105％近い株式リベート割引を与える最近のストックオプションベースの構造によっても促進されます。全額リベートにより、OpenAI/Meta が十分なコンピューティングを購入すると、AMD 株は最終レベルの 600 ドルに達します。 TCO あたりの Helios パフォーマンスは非常に優れているため、この構造と組み合わせると、100 万トークンあたりのコストである AMD のコストは実質的にマイナスになります。 AMD は事実上、Helios ラックとそれに加えて 5% の追加料金を、SF に拠点を置く OpenAI と呼ばれる非営利団体に無償で提供しています。
記事の最初の部分では、Helios アーキテクチャと、Hopper SM90 の ISA のクローンである MI455X (gfx1250) 命令セットについて詳しく説明します。 Helios のスケールアウトおよびスケールアップ ネットワーキング アーキテクチャについても説明します。記事の後半では、ソフトウェア スタックの詳細をさらに詳しく説明します。 3 番目のパートでは、所有者の経済学に焦点を当てます。

また、OpenAI/Anthropic/Meta の株式ベースのリベート割引の経済性と、この構造が総所有コスト (TCO) にどのような影響を与えるかについても分析します。
私たちは MI300X、MI325X、MI355X の ROCm ソフトウェアを毎日使用しており、四半期ごとに一貫してナンバー 1 のバグ レポーターでもあります。 AMD ソフトウェアの改善に全力で取り組み、バグ レポートを迅速にトリアージしてくれている数人の 10x エンジニアに声を大にして伝えたいと思います。 Honxia、Chun Fang、HaiShaw、Thomas Wang、Andy Luo、Seungrok、Bill He、Teresa Shan、Parth、Duyi Wang、Gilbert などに感謝します。 AMD の優秀な 10x エンジニアのほとんどは上海にいます。 AMD の MoRI 集団および UMBP KVCache オフロード チーム、AMD の分散アプリケーション フォワード デプロイメント エンジニアリング チーム、および第一原理ベースの推論エンジニアリングの実行方法を理解しているその他の AMD チームは、ほとんどが上海に拠点を置いています。 ROCm ソフトウェア スタックの最も重要な部分の多くは中国で構築されています。
Lisa Su、Mark Papermaster、Sharon Zhou、Anush Elangovan、Vamsi Boppana への推薦
ソフトウェア品質を向上させるための自動テストに関しては、過去 1 年間で順調な進歩が見られましたが、進歩のペースは必要とされるほど積極的ではなく、自動テスト CI および社内ソフトウェア開発チームに十分な GPU クラスターを提供することに関しては、依然として緊迫感が欠けています。いつも遅すぎるのです。
私たちが皆さん (Vamsi、Anush) と数か月ごとに会い、時にはリサとも会うたびに、CI がもっと良くなる可能性があることを常に強調し、確かに CI が正しい方向に進んでいる具体的な例を示すことはできますが、全体的な戦略を積極的に強化する必要があります。たとえば、Kubernetes Inferencing Pollara NIC CI はまだ存在します。

Nvidia の ConnectX Nightly CI と同等の 0% です。 Kubernetes は、世界中のほとんどの推論デプロイメントで使用されているレイヤーです。問題は、AMD エンジニアリングがそのサポートを追加したくないことではなく、むしろ内部 CI 能力への投資不足によって妨げられているということです。 Advancing AI 2026 までにこの面で同等に達する AMD の予定された ETA は、クラスターの問題により実現しませんでした。
vLLM 側では、今週 AMD クラスター インフラの安定性の問題により、vLLM ゲート自動テストの進捗が大幅に後退しました。 AMD の熱心なエンジニアは、バックストップされた一時クラスターに過度に依存している AMD の内部容量不足を考慮して、AMD のリーダーシップが AMD 社内の vLLM チームからクラスターを引き離し、別の場所にデプロイし始めるまで、Advancing AI 2026 までにゲーティングで CUDA と同等以上の 90% を達成するという目標を達成するために、過去数週間で vLLM ゲーティングに関して順調な進歩を遂げていました。
ゲート/ブロッキング テストは、テストに合格しない限り PR をマージできないため、テストが最高品質であることを意味します。これにより、バグがマージされるのを防ぎます。 AMD のリーダーシップは非技術者の注意をそらして非ゲート合格率を示すかもしれませんが、本当に重要な指標はゲート同等性とゲート通過率です。
私たちは、AMD のリーダーが社内の vLLM チームに安定したクラスターを提供することの優先順位を再設定し、将来このような問題を防ぐためにキャパシティ プランニングの考え方を更新し、AMD の社内の熱心な vLLM エンジニアがゲーティングで CUDA vLLM との 90% 以上の同等性を達成する作業に集中し、AMD の社内 SGLang チームと同じ速度で動作するために必要なツールを備えられるようにすることを願っています。
ほとんどの AMD エンジニアの主な不満は、単にランダムにクラスタを構築する必要がない、安定した CI クラスタの提供に向けて、リーダーシップが依然として視点を更新する必要があるということです。

細分化され、ある CSP から別の CSP に移動されます。
さらに、社内で開発に使用できる GPU が常に不足しています。単一ノードの集約推論の場合、内部で動作するのに十分な GPU があります。しかし、分散マルチノード推論の最適化 (wideEP と非集約 PD) の時代では、十分な GPU はありません。今月さらに 2,000 台の MI355X がオンラインになり、今年後半に 6,000 台の MI325X/MI355X がオンラインになるとしても、総容量は依然として十分ではなく、Nvidia が内部開発用に保有している安定した長期クラスターよりも容量が 1 桁以上少ないままです。この GPU ノードの不足は、エージェント コーディングの台頭によりさらに悪化している問題です。以前は、各人間のエンジニアが DI 推論ソフトウェア開発を行うには、いくつかのノードが必要でした。しかし、現在ではエージェント コーディングが使用されており、各エージェントはコードをテストするために GPU を必要とし、各人間は同時に数十のエージェントを実行でき、各エージェントは同時に数十のサブエージェントしか実行できません。したがって、 DynoSim のような分散推論シミュレーション ツールがなければ、内部の内部 GPU 容量がさらに大幅に不足することになります。
さらに、Blackwell (SM100) とよく似た ISA を使用する Rubin (SM107) とは異なり、MI455 (gfx1250) は MI355 (gfx950) とは完全に異なる ISA であり、コードパスとカーネルが完全に異なるため、gfx950 と gfx1250 の両方でテストする必要があります。これは、今後も能力を圧迫するもう 1 つの要因です。大胆な大きな目標は、Advancing AI 2026 までに MI455X オープンソース vLLM、SGLang の夜間自動 CI を実現することでしたが、このタイムラインは逃されました。新しいタイムラインは2026年10月まで延期されたと聞いていますが、彼らはまだそれを左シフトして元に戻そうとしています。

2026 年 8 月または 9 月のタイムラインに。
私たちは、社内ソフトウェア開発の GPU クラスターと CI クラスターの自動テストのこのゆっくりとした増加が、ソフトウェアの品質とパフォーマンスの向上のペースを遅らせる大きなリスクの 1 つであると考えており、AMD のリーダーシップが今後のキャパシティ プランニング戦略を再検討することを期待しています。
パート 1: AMD MI455 シリコンおよび Helios ラック
出典: AMD AMD はシリコン エンジニアリングで再びそれを行いました。 MI455X は、シリコン最前線の工場から生み出される最先端のチップです。 AMD は、MI455 のコンピューティング タイルと、N2 の初期採用者である Venice CPU の両方を搭載した 2nm データセンター シリコンを出荷した最初の企業です。競合するアクセラレータ プラットフォームはすべて N3 上にあります。
AMD は、パッケージ統合に関しても引き続きリードしています。 MI455 パッケージは、5.5 倍のレチクル サイズで出荷される最大の CoWoS-L モジュールです。 AMD は、TSMC の SoIC-X ハイブリッド ボンディングを採用している唯一の企業です。これにより、AMD は、x および y 次元だけでなく、z 次元でもシリコン フットプリントを拡張できます。これらすべてが合わさって、MI455 のパッケージ内のロジック シリコンの総量は 3,470mm2 となり、単一パッケージで出荷されるシリコンの量としてはこれまでで最も多くなります。
パッケージ レイアウトは MI355X から大きく取り入れられています。 8 N2'X

[切り捨てられた]

## Original Extract

Agentic Kernel Generation, Improvement in Software Quality, Unstable Internal Development Clusters, Helios MI455X Production Ramp Hell, Up to 105% Discounts from Finance Engineering

Can AMD break the CUDA Moat? AMD Advancing AI 2026
Subscribe Sign in Can AMD break the CUDA Moat? AMD Advancing AI 2026
Up to 105% Equity Rebate Discounts for OpenAI, Agentic Kernel Generation, Improvement in Software Quality, Unstable Internal Development Clusters, Helios MI455X Production Ramp Hell
184 4 13 Share When we published our first AMD software article , we gave AMD a 0% chance of closing the gap with Nvidia in AI accelerators. Software was broken, progress was unexciting, and we were the top bug submitter for many months with dozens of AMD engineers triaging our bug reports.
Six months later, in our AMD 2.0 article, we took the non-consensus position of upgrading from 0% chance to a much more meaningful chance at success . We published this opinion when sentiment towards AMD was sitting at rock bottom, and when most of the market thought we were were being far too optimistic towards AMD.
That view was based on our observations that AMD has the leadership that can create change rather than suffer from committee style leadership. Lisa quickly hopped on a calls with us and has since then implemented many of our suggestions . We saw an AMD that had finally recognized the importance of software and had a sense of urgency, moving in the right direction even if the destination was still far off. Since then, the signal has only gotten stronger.
Based on our experience on AMD software stack this year, we update our view again from non-zero percentage chance to now a great chance of success as long as AMD solves the two major risks we outline below. It is important to highlight that just because AMD gains market share, that doesn’t mean that Nvidia will do poorly. The pie is growing rapidly for everyone, and Nvidia will continue to massively grow revenue. AMD poses potential competition to Nvidia on the software front, and Jensen will need to cut bureaucracy and flatten the layers of different required internal stakeholder approvals required for even the simplest of tasks if he wants Nvidia to move faster and defend their lead.
Anthropic has publicly announced that they will deploy 2GW of AMD’s chips, Lisa Su and team have leaned into an agentic oriented engineering culture. Anthropic Head of Compute Tom Brown explained how he used Claude over the weekend with “/goal” to bring-up internal Claude inference stack on AMD hardware as a case in point. We believe that since AMD’s compiler and most of their kernels are open sourced, that they are better positioned for the agentic age besides one major risk that we will discuss later . Three months ago, our Accelerator model noted that Anthropic will be an AMD customer . We also noted this on our socials a week ago .
In 2023, Microsoft dropped AMD after the MI300X due to unreliable Samsung 2023 HBM memory and poor software quality, subsequently skipping both the MI325X and the MI355X. Microsoft has since made an about face and has announced that it will deploy MI455X Helios. We believe that OpenAI will be the main end customer for Azure’s MI455X racks. In a strategy somewhat similar to the Nvidia-Groq deal, AMD is announcing a deal with Cerebras to do PD disagg for ultra-fast interactivity inferencing.
There are two major risks that AMD needs to guard against:
Supply chain checks and engineering first principles analysis and understanding show that AMD’s first AI rack scale system, Helios, is going currently going through a slow rack production ramp given that it is not using a cableless tray design, which the Rubin Oberon rack is now adopting. Furthermore, since AMD has a weak SerDes design, up to 85% of its backplane needs to be retimed, requiring over over 550 Broadcom ethernet retimers per rack. Furthermore, it is running into backplane reliability challenges during rack production ramp hell.
The chief complaint from most AMD engineers internally is that there is a persistent lack of stable GPU clusters for internal software development teams and a lack of stable GPU clusters for automated testing CI . This is blocking AMD’s rate of progress and it is holding AMD back from harnessing the potential upside of AI coding Agents because each AI Agent requires GPUs as well and also requires a testing tool use loop. We will more explain below on our recommendation section.
If AMD can overcome these challenges, we strongly believe that AMD will be well positioned to do well and take market share. This will be helped along as well by the recent stock option based structure whereby AMD gives Meta and OpenAI close to a 105% equity rebate discount using some clever financial engineering. The full rebate triggers a AMD stock reaches the final level of $600 and once OpenAI/Meta buys enough compute. Helios performance per TCO is so great that AMD that cost per million tokens is practically negative cost when combined with this structure! AMD is practically giving away Helios racks and an 5% extra on top of that to an SF-based nonprofit called OpenAI.
In the first part of our article, we go through deep dive into the Helios architecture and the MI455X (gfx1250) instruction set, which is an clone of Hopper SM90’s ISA. We will also discuss the Helios scale-out and scale-up networking architecture. The second part of the article we will further into details of the software stack. For the third part, we will focus on the economics of owning and operating the Helios rack, and also break down the economies of OpenAI/Anthropic/Meta’s equity based rebate discounts and how this structure affects total cost of ownership (TCO).
We are a daily user of ROCm software on MI300X, MI325X, MI355X and are also the #1 bug reporter consistently every quarter! There are a few 10x engineers we want to shout out to that have been working 997 to improve AMD software and have been quickly triaging our bug reports. Many thanks to Hongxia, Chun Fang, HaiShaw, Thomas Wang, Andy Luo, Seungrok, Bill He, Teresa Shan, Parth, Duyi Wang, Gilbert, and many more. Most of AMD’s best 10x engineers are in Shanghai. AMD’s MoRI collective and UMBP KVCache offloading team, AMD’s disaggregated application forward deployed engineering team, and other AMD teams that understand how to do first principles-based inference engineering are all mostly based in Shanghai. A lot of the ROCm software stack’s most important pieces are built in China.
Recommendations to Lisa Su, Mark Papermaster, Sharon Zhou, Anush Elangovan and Vamsi Boppana
There has been good progress over the past year on automated testing to improve software quality, but the pace of progress has not been as aggressive as needed and there continues to lack of sense of urgency with respect to providing enough GPU clusters for automated testing CI and for internal software development teams. It is always too little too late.
Every time we meet with y’all (Vamsi, Anush) every couple of months and occasionally meet with Lisa, we always highlight that CI could be better and indeed we can point to specific examples where CI is heading in correct direction, but the overall strategy needs to be aggressively ramped up. For example, the Kubernetes Inferencing Pollara NIC CI still sits at 0% parity with Nvidia’s ConnectX Nightly CI. Kubernetes is the layer that most inference deployments across the world use. The issue is not with AMD engineering not wanting to add support for it, rather they are being blocked by lack of investment in internal CI capacity. The planned ETA of AMD reaching parity on this front by Advancing AI 2026 was missed due to cluster issues.
On the vLLM side, vLLM gating automated test progress has massively regressed due to AMD cluster infra stability issues this week. AMD’s hardcore engineers were making good progress on vLLM gating over the past couple of weeks to reach the goal of attaining at least 90% parity with CUDA on gating by Advancing AI 2026 until AMD leadership started pulling clusters away from AMD’s internal vLLM team to deploy elsewhere given AMD’s internal capacity crunch that overly relies on backstopped temporary clusters.
Gating/blocking tests mean the tests are of the highest quality because PRs cannot merge unless the test is passing. This prevents bugs from being merged. While AMD’s leadership may distract non-technical folks by showing their non-gating pass rate, gating parity and gating pass rate are the metrics that truly matter.
We hope that AMD’s leadership can re-prioritize giving their internal vLLM team stable clusters and update their philosophy on capacity planning to prevent issues like this in the future so that AMD’s internal hardcore vLLM engineers focus on the work of reaching 90%+ parity with CUDA vLLM on gating and have the tools needed to operate at the same velocity as the AMD internal SGLang team.
The chief complaint from most AMD engineers is that leadership still needs to update their viewpoint towards providing stable CI clusters that don’t just randomly need to be migrated and shifted from one CSP to another CSP.
Moreover, there is a constant lack of GPUs available internally for development. For single node aggregated inferencing, there are enough GPUs to go around internally. But in the age of Distributed Multi-node Inferencing optimizations (wideEP and disaggregated PD), there is nowhere near enough GPUs. Even with the additional 2,000 MI355Xs coming online this month and the 6,000 MI325X/MI355X coming online later this year, total capacity will still not be enough and remains more than an order of magnitude less capacity than the stable long term clusters Nvidia has for internal development. This lack of GPU nodes is a problem that has been getting even worse due to the rise of agentic coding. Previously, each human engineer required a couple of nodes to conduct DI inference software development. But now with agentic coding, each agent requires GPUs to test their code against and each human can have dozens of agents running at the same time and each agent in turn can only dozens of sub agents running at the same time too. Thus, without Distributed Inferencing simulation tools like DynoSim , there will be an even greater shortage of internal GPU capacity internally.
Furthermore, unlike Rubin (SM107), which uses a very similar ISA to Blackwell (SM100), MI455 (gfx1250) is a completely different ISA from MI355 (gfx950) and has completely different codepaths and kernels tha will require testing on both gfx950 and gfx1250. This is another factor that will continue to stress capacity. The big audacious goal was to have MI455X open-source vLLM, SGLang nightly automated CI by Advancing AI 2026, but this timeline was missed. We hear that the new timeline has been delayed untill October 2026 but they are still trying to left-shift it back to an August/September 2026 timeline.
We view this slow ramp of internal software development GPU clusters and automated testing of CI clusters as one of the major risks slowing the pace of improvement of software quality and performance and we hope that AMD leadership revisits their capacity planning strategy going forward.
Part 1: AMD MI455 Silicon & Helios Rack
Source: AMD AMD has done it again on silicon engineering. The MI455X is the most advanced chip that comes out of a fab on the silicon front. AMD is the first company to ship 2nm datacenter silicon with both the MI455’s compute tiles as well as the Venice CPU being early adopters of N2. All of the competing accelerator platforms are on N3.
AMD continues to lead with respect to on package integration as well. The MI455 package is the largest CoWoS-L module shipping at 5.5x reticle size. AMD is still the only company adopting TSMC’s SoIC-X hybrid bonding, which allows AMD to scale silicon footprint in the z dimension as well as the x and y dimensions. All this comes together to bring the MI455 to a total of 3,470mm2 of logic silicon in the package, by far the most amount of silicon that is being shipped in a single package.
The package layout borrows heavily from the MI355X. 8 N2 ‘X

[truncated]
