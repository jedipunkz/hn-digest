---
source: "https://www.servethehome.com/tensordyne-napier-ai-processor-announced-with-logarithmic-math/"
hn_url: "https://news.ycombinator.com/item?id=48552631"
title: "Tensordyne Napier AI Processor Announced with Logarithmic Math"
article_title: "Tensordyne Napier AI Processor Announced with Logarithmic Math - ServeTheHome"
author: "lumpa"
captured_at: "2026-06-16T10:41:36Z"
capture_tool: "hn-digest"
hn_id: 48552631
score: 1
comments: 0
posted_at: "2026-06-16T09:18:46Z"
tags:
  - hacker-news
  - translated
---

# Tensordyne Napier AI Processor Announced with Logarithmic Math

- HN: [48552631](https://news.ycombinator.com/item?id=48552631)
- Source: [www.servethehome.com](https://www.servethehome.com/tensordyne-napier-ai-processor-announced-with-logarithmic-math/)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T09:18:46Z

## Translation

タイトル: 対数演算を備えた Tensordyne Napier AI プロセッサを発表
記事のタイトル: Tensordyne Napier AI プロセッサーが対数演算を搭載して発表 - ServeTheHome
説明: Tensordyne Napier は、AI 推論の背後にある数学と独自の 72 アクセラレータ アーキテクチャを対象とした新しい AI アクセラレータです

記事本文:
フェイスブック
リンクトイン
RSS
TikTok
×
ユーチューブ
フォーラム
ワークステーション
ワークステーションプロセッサ
TrueNAS / FreeNAS NAS サーバーの上位ハードウェア コンポーネント
pfSense アプライアンスの上位ハードウェア コンポーネント
napp-it および Solarish NAS サーバーの上位ハードウェア コンポーネント
Windows Server 2016 Essentials ハードウェアのおすすめ
DIY WordPress ホスティング サーバー ハードウェア ガイド
ストレージの信頼性
レイドカリキュレーター
RAID 信頼性計算ツール |単純な MTTDL モデル
編集および著作権ポリシー
ワークステーション
ワークステーションプロセッサ
TrueNAS / FreeNAS NAS サーバーの上位ハードウェア コンポーネント
pfSense アプライアンスの上位ハードウェア コンポーネント
napp-it および Solarish NAS サーバーの上位ハードウェア コンポーネント
Windows Server 2016 Essentials ハードウェアのおすすめ
DIY WordPress ホスティング サーバー ハードウェア ガイド
対数演算を備えた Tensordyne Napier AI プロセッサを発表
Tensordyne は、3nm AI プロセッサーおよび独自の対数数学を中心に構築されたラックスケール推論プラットフォームである Napier を発表しました。興味深いのは、別の AI チップ新興企業が混雑した市場に参入するというだけではなく、アクセラレータの計算を変更することで乗算器領域を削減し、オンチップ SRAM を増加し、ラックレベルの推論経済性を改善できるという同社の主張です。今のところ、Napier はまだテープアウトされたチップであり、2027 年のシステムロードマップにあるため、大きな問題は、パフォーマンスとソフトウェアの主張が実際の展開に耐えられるかどうかです。
Tensordyne Napier AI プロセッサを発表
Tensordyne は、AI 推論の速度とコストの両方を攻撃する方法として Napier を位置づけています。同社によれば、従来の行列乗算リソースのみを中心に構築するのではなく、対数数学的アプローチにより乗算演算が加算に変換されるという。加算器は乗算器よりも小さく、一般に消費電力が低いため、この約束はより有用です。

l メモリ用のシリコン領域とより優れたシステムバランス。
そのために、チップだけではなくクラスターアーキテクチャを搭載したエコシステムを発表している。
今日の AI インフラストラクチャに関する議論の多くは、もはやピーク アクセラレータの TOPS または FLOPS だけを問題にしているわけではないため、これは重要です。ロングコンテキスト推論、エージェントワークフロー、専門家混合モデルは、メモリ、相互接続、デコードスループット、ラック電力、冷却によって制約を受ける可能性があります。 Tensordyne の主張は、よりバランスの取れたチップとラックの設計により、現在のハイエンドの代替品よりもラックあたりのトークンとメガワットあたりのトークンの数を増やすことができるというものです。
Tensordyne は、自社の TDN72 ラックを、2 兆パラメータの GPT MoE モデルのより大きなマルチラック構成と比較しています。この比較において同社は、120kWのTDN72ラック1つでユーザーあたり1秒あたり1,300トークンに達できるが、NVIDIAとGroqでは9つのラックと1.5MWが必要で、AWSとCerebrasでは14つのラックと800kWが必要であるとしている。これらの比較は注目を集めますが、Napierは現時点で製品を発表しています。
完全な TDN72 システムは、72 ノード、合計 68 ペタフロップスのコンピューティング、および 42 TB の HBM を中心に設計されています。 Tensordyne によれば、その容量は最大 10 兆から 20 兆のパラメータを持つモデルを対象としており、メモリの設置面積と専門的なルーティングがシステムレベルの大きな課題となります。インターコネクト、メモリ、電源、または冷却インフラストラクチャが制限要因となる場合、アクセラレータを追加するだけでは役に立たないため、ラックスケール設計が重要になるのはここでもあります。
Napier自体は、1,380億個のトランジスタを備えた3nm TSMCチップです。 Tensordyne には、ダイあたり 2.1 ペタフロップスのコンピューティング、1.33 GHz アクセラレータ コア、1.5 GHz CPU、256 MB の SRAM、および 144 GB の HBM3E が記載されています。より重要な主張の 1 つは、Napier には NVIDIA Blackwell の 5 倍の SRAM があるということです。それが役に立つならw

orkloads の場合、追加の SRAM は、より多くのデータをコンピューティング ファブリックの近くに保持し、システム内でデータを移動するペナルティを軽減するのに役立ちます。
対数数学の概念はアーキテクチャ上のフックです。 Tensordyne によれば、乗算器のフットプリントを削減すると、SRAM の余地がさらに広がり、シストリック アレイとベクトル プロセッサがスループットを処理できるようになります。これは、AI アクセラレータの問題を組み立てる方法であり、より高密度の行列演算ユニットを単に数える方法とは異なります。同時に、数値的アプローチの変更は精度、ソフトウェア、モデル移植に影響を与える可能性があるため、サードパーティのワークロード テストが最も必要な部分でもあります。
トレイ レベルでは、Tensordyne は 9 個の Napier チップを 1.3TB の HBM3E、8TB のストレージ、Intel Xeon ホスト CPU、およびデュアル 200GbE を備えた 1RU AI コンピューティング トレイにパッケージ化しています。 4 つのトレイで TDN72 ポッドが構成され、4 つのポッドが標準の 52RU ラックに収まります。実用上重要な点は、Tensordyne が空冷システムをターゲットにしていることです。大規模AIには液体冷却が使用されますが、Tensordyneは空冷システムをターゲットにしています。また興味深いのは、フロントエンドが 2x 200GbE であるということは、Intel Xeon ホスト CPU が x16 リンクあたり 800Gbps で駆動できる PCIe Gen6 ではないことを示しているように見えることです。
スケールアップ接続も設計の主要な部分です。 Tensordyne はその相互接続を TDN Link と呼び、72 チップのシステム全体で 1TB/s の帯域幅でマイクロ秒未満のチップ間遅延を提供できると述べています。専門家が混在するワークロードやエージェント型 AI ワークロードの場合、インターコネクトはアクセラレータと同じくらい重要になる可能性があります。これは、エキスパートのルーティング、アクティベーションの移動、多くのユーザーへの供給の維持により、遅延と帯域幅の制限が生じる可能性があるためです。 NVL72 スパインの代わりに、これは従来のシャーシ スイッチ ネットワーキング ソリューションに似ています。
トポロジの柔軟性は同じ相互接続の一部です

オリー。 Tensordyne 氏は、どのチップもワークロードに合わせてグループ化できるため、ソフトウェア スタックがそれを透過的にできれば、フェイルオーバーやモデルの配置に役立つだろうと述べています。これは大規模な展開には有益な主張ですが、運用の詳細が重要な領域でもあります。顧客がメリットを感じる前に、クラスター スケジューラ、モデル サービング レイヤー、障害処理、および可観測性が十分に機能する必要があります。
ソフトウェアは、最終的には発売の中で最も難しい部分になるかもしれません。 Tensordyne は、SDK を備えた Hugging Face でホストされるモデル ハブ、PyTorch および Triton で定義されたモデルの直接コンパイル、および tensordyne.nn と呼ばれるカスタム Python eDSL について話しています。 NVIDIA の CUDA エコシステムは、フレームワーク、カーネル、プロファイリング ツール、展開パターン、開発者の習慣の巨大な基盤です。新しい AI アクセラレータは、ソフトウェア パスが顧客に試してもらえるほど簡単であると感じさせるものでなければなりません。
ここではパートナーも重要です。 Tensordyne は、シャーシおよびインフラストラクチャ コンポーネントに関して HPE および Juniper と協力していると述べており、これにより同社が単なるチップ開発者ではなくシステム ベンダーとしてより信頼できるものになるはずです。 Broadcom 経由の TSMC による 3nm テープアウトは意味のあるマイルストーンですが、ラックスケール AI システムには、サプライ チェーン、プラットフォーム検証、フィールド サポート、および新しいアーキテクチャにワークロードを賭ける意欲のある顧客が必要です。
もう一つの課題はタイミングです。 Tensordyne によると、ベータ プログラムは 2027 年第 1 四半期に計画されており、システムの出荷は 2027 年第 2 四半期末までに予定されています。その時までに、NVIDIA、AMD、ハイパースケールの社内シリコンへの取り組み、Cerebras、Groq、その他の AI インフラストラクチャのオプションが再び移行することになるでしょう。ネイピア氏は、主張されている効率が実際のモデル提供、実際のソフトウェア スタック、および実際の顧客の運用においても維持されることを示す必要があります。
Tensordyne Napier は、AI アクセラレータの発表の中で最も興味深いものの 1 つです。

NVIDIA とは異なるスケールにするだけでなく、計算を変更します。 NVIDIA と同様のフォーム ファクターを持つアクセラレータを構築し、自社の方が安価であると主張することは、他の企業が成功を見てきた方法とは異なる傾向があるため、この数学的変化は興味深いものです。 3nm テープアウト、1,380 億個のトランジスタという数字、大規模な SRAM の要求、42TB HBM ラック構成、および空冷 TDN72 システムはすべて、これを注目する価値のあるものにしています。
それでも、魅力的なローンチと AI プラットフォームの成功との間には大きな隔たりがあります。ラックあたりのパフォーマンスとメガワットあたりのパフォーマンスは、まさに目標とする適切な指標です。 Tensordyne のテクノロジーが機能し、2027 年に実現できれば、Napier は推論インフラストラクチャの注目すべき代替となる可能性があります。おそらく、数十億ドル規模の取引が見られるようになるでしょう。それまでは、これは証明すべきことがまだたくさんある野心的なアーキテクチャなので、注目してみるのは興味深いでしょう。
人間は寓話5と神話5へのアクセスを停止する
ASRock 産業用 NUC BOX-358H ミニ PC レビュー
Dell Pro Max 16 Plus レビュー よりモバイルな NVIDIA RTX Pro 5000 Blackwell システム
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
STH ニュースレターに登録してください!
このサイトはスパムを低減するために Akismet を使っています。コメントデータがどのように処理されるかをご覧ください。
STH の最高の機能を毎週受信箱に配信します。 STH から毎週最高の投稿を厳選し、直接お届けします。
オプトインすると、ニュースレターの送信に同意したことになります。当社ではサードパーティのサービスを使用してサブスクリプションを管理しているため、いつでもサブスクリプションを解除できます。

## Original Extract

Tensordyne Napier is a new AI accelerator targeting the math behind AI inference and its own 72 accelerator architecture

Facebook
Linkedin
RSS
TikTok
X
Youtube
Forums
Workstation
Workstation Processors
Top Hardware Components for TrueNAS / FreeNAS NAS Servers
Top Hardware Components for pfSense Appliances
Top Hardware Components for napp-it and Solarish NAS Servers
Top Picks for Windows Server 2016 Essentials Hardware
The DIY WordPress Hosting Server Hardware Guide
Storage Reliability
Raid Calculator
RAID Reliability Calculator | Simple MTTDL Model
Editorial and Copyright Policies
Workstation
Workstation Processors
Top Hardware Components for TrueNAS / FreeNAS NAS Servers
Top Hardware Components for pfSense Appliances
Top Hardware Components for napp-it and Solarish NAS Servers
Top Picks for Windows Server 2016 Essentials Hardware
The DIY WordPress Hosting Server Hardware Guide
Tensordyne Napier AI Processor Announced with Logarithmic Math
Tensordyne announced Napier, a 3nm AI processor and rack-scale inference platform built around proprietary logarithmic mathematics. The interesting part is not just another AI chip startup entering a crowded market, but the company’s claim that changing the math in the accelerator can reduce multiplier area, increase on-chip SRAM, and improve rack-level inference economics. For now, Napier is still a taped-out chip and 2027 system roadmap, so the big question is whether the performance and software claims survive contact with real deployments.
Tensordyne Napier AI Processor Announced
Tensordyne is positioning Napier as a way to attack both the speed and the cost of AI inference. Instead of building only around more conventional matrix-multiply resources, the company says its logarithmic math approach turns multiplication operations into additions. Adders are smaller and generally lower-power than multipliers, so the promise is more useful silicon area for memory and better system balance.
To that end, it is announcing an ecosystem to not just have a chip, but a cluster architecture.
That matters because a lot of today’s AI infrastructure discussion is no longer just about peak accelerator TOPS or FLOPS. Long-context inference, agentic workflows, and mixture-of-experts models can become constrained by memory, interconnect, decode throughput, rack power, and cooling. Tensordyne’s argument is that a more balanced chip and rack design can deliver more tokens per rack and more tokens per megawatt than current high-end alternatives.
Tensordyne compares its TDN72 rack against larger multi-rack configurations for two-trillion-parameter GPT MoE models. In that comparison, the company says one 120kW TDN72 rack can reach 1,300 tokens per second per user, while NVIDIA and Groq require nine racks and 1.5MW, and AWS plus Cerebras require fourteen racks and 800kW. Those comparisons are attention-grabbing, but Napier is announcing product at this point.
A full TDN72 system is designed around 72 nodes, 68 petaflops of total compute, and 42TB of HBM. Tensordyne says its capacity is aimed at models with up to 10 trillion to 20 trillion parameters, where the memory footprint and expert routing become major system-level challenges. This is also where rack-scale design matters, since simply adding accelerators does not help if the interconnect, memory, power, or cooling infrastructure becomes the limiting factor.
Napier itself is a 3nm TSMC chip with 138 billion transistors. Tensordyne lists 2.1 petaflops of compute per die, a 1.33GHz accelerator core, a 1.5GHz CPU, 256MB of SRAM, and 144GB of HBM3E. One of the more important claims is that Napier has five times the SRAM of NVIDIA Blackwell. If that holds up in useful workloads, the extra SRAM could help keep more data close to the compute fabric and reduce the penalty of moving data around the system.
The logarithmic math concept is the architectural hook. Tensordyne says reducing the multiplier footprint leaves more room for SRAM, while a systolic array and vector processor handle throughput. That is a different way to frame the AI accelerator problem than simply counting more dense matrix math units. At the same time, it is also the part of the story that most needs third-party workload testing, since changing numerical approaches can have accuracy, software, and model-porting implications.
At the tray level, Tensordyne is packaging nine Napier chips into a 1RU AI Compute Tray with 1.3TB of HBM3E, 8TB of storage, Intel Xeon host CPUs, and dual 200GbE. Four trays make a TDN72 pod, and four pods fit in a standard 52RU rack. An important practical point is that Tensordyne is targeting an air-cooled system. Liquid cooling is used for large-scale AI, but Tensordyne is targeting an air-cooled system. Also interesting is that the front-end as 2x 200GbE seems to indicate that the Intel Xeon host CPUs will not be PCIe Gen6 where you can drive 800Gbps per x16 link.
Scale-up connectivity is another major part of the design. Tensordyne calls its interconnect TDN Link and says it can provide sub-microsecond chip-to-chip latency with 1TB/s of bandwidth across the 72-chip system. For mixture-of-experts and agentic AI workloads, the interconnect can matter as much as the accelerator because routing experts, moving activations, and keeping many users fed can expose latency and bandwidth limits. Instead of the NVL72 spine, this looks more like a traditional chassis switch networking solution.
Topology flexibility is part of that same interconnect story. Tensordyne says any chips can be grouped for a workload, which would help with failover and model placement if the software stack can make that transparent. That is a useful claim for large deployments, but it is also an area where operational details matter. Cluster schedulers, model serving layers, failure handling, and observability need to work well before customers feel the benefit.
Software may end up being the harder part of the launch. Tensordyne is talking about a Hugging Face-hosted model hub with its SDK, direct compilation for PyTorch and Triton-defined models, and a custom Python eDSL called tensordyne.nn. NVIDIA’s CUDA ecosystem is a huge base of frameworks, kernels, profiling tools, deployment patterns, and developer habits. Any new AI accelerator has to make the software path feel easy enough that customers will try it.
Partners also matter here. Tensordyne says it is working with HPE and Juniper for chassis and infrastructure components, which should help the company look more credible as a systems vendor rather than only a chip developer. A 3nm tape-out through TSMC via Broadcom is a meaningful milestone, but rack-scale AI systems require a supply chain, platform validation, field support, and customers willing to bet workloads on a new architecture.
Timing is the other challenge. Tensordyne says beta programs are planned for Q1 2027, with system shipments expected by the end of Q2 2027. By then, NVIDIA, AMD, hyperscale internal silicon efforts, Cerebras, Groq, and other AI infrastructure options will have moved again. Napier needs to show that the claimed efficiency holds up in real model serving, real software stacks, and real customer operations.
Tensordyne Napier is one of the more interesting AI accelerator announcements because it is trying to change the math, not just scale differently from NVIDIA. Building an accelerator that has a similar form factor as NVIDIA and saying that you are cheaper tends not to be the way others have seen success, so the math change is interesting. The 3nm tape-out, 138 billion transistor figure, large SRAM claim, 42TB HBM rack configuration, and air-cooled TDN72 system all make this worth watching.
Still, the gap between a compelling launch and a successful AI platform is large. Performance per rack and performance per megawatt are exactly the right metrics to target. If Tensordyne’s technology works and can deliver in 2027, Napier could be a notable alternative for inference infrastructure. Perhaps we will start seeing deals on a multi-billion-dollar scale. Until then, this is an ambitious architecture with a lot still left to prove, so it will be interesting to watch.
Anthropic Halts Access to Fable 5 and Mythos 5
ASRock Industrial NUC BOX-358H Mini PC Review
Dell Pro Max 16 Plus Review A More Mobile NVIDIA RTX Pro 5000 Blackwell System
Save my name, email, and website in this browser for the next time I comment.
Sign me up for the STH newsletter!
This site uses Akismet to reduce spam. Learn how your comment data is processed.
Get the best of STH delivered weekly to your inbox. We are going to curate a selection of the best posts from STH each week and deliver them directly to you.
By opting-in you agree to have us send you our newsletter. We are using a third party service to manage subscriptions so you can unsubscribe at any time.
