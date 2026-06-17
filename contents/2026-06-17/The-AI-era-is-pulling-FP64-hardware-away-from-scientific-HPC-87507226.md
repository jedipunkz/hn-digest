---
source: "https://fortran-lang.discourse.group/t/the-ai-era-is-pulling-fp64-hardware-away-from-scientific-hpc/10971"
hn_url: "https://news.ycombinator.com/item?id=48572420"
title: "The AI era is pulling FP64 hardware away from scientific HPC"
article_title: "The AI era is pulling FP64 hardware away from scientific HPC - Fortran Discourse"
author: "Bostonian"
captured_at: "2026-06-17T16:20:14Z"
capture_tool: "hn-digest"
hn_id: 48572420
score: 2
comments: 0
posted_at: "2026-06-17T16:05:18Z"
tags:
  - hacker-news
  - translated
---

# The AI era is pulling FP64 hardware away from scientific HPC

- HN: [48572420](https://news.ycombinator.com/item?id=48572420)
- Source: [fortran-lang.discourse.group](https://fortran-lang.discourse.group/t/the-ai-era-is-pulling-fp64-hardware-away-from-scientific-hpc/10971)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T16:05:18Z

## Translation

タイトル: AI 時代により、FP64 ハードウェアが科学的 HPC から引き離されています。
記事のタイトル: AI 時代は FP64 ハードウェアを科学的 HPC から引き離しています - Fortran Discourse
説明: こんにちは。
1 つ、おそらく 2 つ質問があります。この質問は、私が参加したハイ パフォーマンス コンピューティング (HPC) に関するウェビナー シリーズ (CNR を通じてイタリア側で企画されたイタリアとドイツの HPC ウェビナー) から生まれました。

記事本文:
= 40rem)" rel="スタイルシート" data-target="デスクトップ" />
= 40rem)" rel="スタイルシート" data-target="chat_desktop" />
= 40rem)" rel="スタイルシート" data-target="discourse-ai_desktop" />
= 40rem)" rel="スタイルシート" data-target="discourse-reactions_desktop" />
= 40rem)" rel="スタイルシート" data-target="poll_desktop" />
Fortran 談話
AI 時代により、FP64 ハードウェアが科学的 HPC から引き離されています
ザギ
2026 年 6 月 17 日、午前 9 時 31 分
1
皆さん、こんにちは。
1 つ、おそらく 2 つ質問があります。この質問は、私が参加したハイ パフォーマンス コンピューティング (HPC) に関するウェビナー シリーズ (CNR を通じてイタリア側で企画されたイタリアとドイツの HPC ウェビナー) から生まれました。そこで私はこの懸念を提起しましたが、私の印象では、他の講演者も同じ程度にそれを共有していなかったように感じました。部屋は私よりも楽観的に傾いていました。まさにそれが、私がこの話をより多くの聴衆に伝えたい理由です。私は間違っているかもしれないし、他の人がどこに着地するのかを聞きたいのです。
懸念されるのは精度です。ほとんどの科学 HPC は倍精度 (FP64) を必要とします。私の専門分野である数値流体力学では、何桁にもわたる物理スケールを解決しますが、それを (非常に高精度の手法を使用して) 正確に行うには、64 ビット浮動小数点が必要です。
AI コンピューティングにはこれは必要ありません。トレーニングと推論は 8 ビットでも適切に機能します (現在は 4 ビットでも)。したがって、2 つのワークロードには異なるハードウェアが必要です。AI には多くの低精度コアが必要ですが、科学には強力な FP64 機能が必要です。
問題は、そこにお金が集まるため、ベンダーが AI 市場に追随していることです。ベクトル FP64 (ピーク、密) を比較すると、最近の傾向は、フラットに保つか下げて、代わりに低精度の演算にトランジスタを費やすことです。
NVIDIA H100: 34 TFLOP/s ベクター FP64、または FP64 テンソルコア パスの場合は 67。新しい B200 は約 40 のベクター FP64 を実行します。ブラックウェル

Hopper が持っていた専用の FP64 テンソルコア パスを削除し、AI 用に約 20 PFLOP/s の FP4 を獲得しました。ルービンのロードマップはFP64をさらに短縮すると伝えられている。
AMD MI300X: 81.7 TFLOP/秒 FP64。新しい MI355X は 78.6 であり、前任者よりも低く、すべてのゲインは AI 推論の FP8/FP4 にあります。
Intel は専用 HPC GPU から撤退しました。現在の HPC シリコンである Aurora の Max シリーズ (Ponte Vecchio) には、スタンドアロンの後継製品がありません。 Intelは、2025年初めに製品としてのFalcon Shoresをキャンセルし、HPCおよびAIラインを1チップのJaguar Shoresに統合し、2026/2027年頃にリリースする予定だ。 Intel は、AI と HPC の両方にサービスを提供すると説明していますが、ピーク FLOPS ではなく総所有コストで競争すると述べており、FP64 の数値は公表していません。
消費者向けシリコンは方向性を最も明確にします。新しい Blackwell ラップトップ チップである NVIDIA の N1X は、AI 精度の数値 (NVFP4、約 1000 TOPS) のみを公開しており、FP64 についてはまったく引用していません。倍精度は単に設計目標ではありません。
したがって、3 つのベンダーすべての方向性は同じであるように見えます。新しいチップは AI 用に構築されており、途中で倍精度の優先順位が静かに下げられます。
強い逆流が 1 つあります。今年発売される AMD の MI430X は、意図的な HPC パーツです。 AMDは、FP64の性能が200 TFLOP/秒を超えていると主張しており、独自の推定では、Alice Recoqueのエクサスケール契約から約211TFLOPSを差し引いており、これはこれまでのGPUの中で最高となる一方、依然としてAI向けにFP4/FP8を搭載している。これは、米国の Discovery およびドイツの Herder と並んで、欧州の次期エクサスケール マシンである Alice Recoque に動力を供給します。したがって、現時点では専用の FP64 ラインがまだ存在します。
しかし、それは、市場全体が反対方向に動いているのに対して、1 つのベンダーの 1 つの製品ラインです。それは私が解決できないことです。一流の FP64 ハードウェア ラインが生き残るか、それともわずかなプレミアムに縮小するかです。

それ以外はすべて AI 向けに最適化されています。
あなたもこの懸念を共有しますか、それとも私がそれを誇張していると思いますか?
それを共有すれば、すでに解決策は見えていますか？
Fortran および HPC コミュニティの他の人々がこれについてどう考えているかを知りたいと思っています。
Sandia のスーパーコンピューター Spectra には NextSilicon の Maverick-2 もあります。
これは常に私の心の中にあります。しかし、私は nvidia または AMD で働く人々と話をしましたが、彼らは皆、FP64 が熱いジャガイモのように廃止されることはないと私に保証してくれました。カードは 2 つの行に分かれており、たとえば B300 は AI/推論手順用で、B200 は科学や HPC 関連のものになります。
真実なのは、FP64 のパフォーマンスの驚異的な向上は見られないということです。つまり、V100 FP64 は 7.9 TFLOP/s、A100 は 19 (??)、H200 は 34 程度で、B200 は、ご指摘のとおり、40 TFLOP/s のクリープよりも低いです。
ここで重要なのは、世に出ているほとんどのワークロードはこれらのピーク FLOP レートを使用しておらず、すべてが DGEMM であるわけではなく、これらを実際に実行できるアプリケーションの数は限られているということです。他のほとんどのアプリは、AI と HPC の両方のワークフローにとって重要なメモリ帯域幅によって制限されていると感じています。ただし、AMD は NVIDIA よりも FP64 サポートを支持しているようですが、Fortran の GPU サポートはまだ最高ではありませんが、大幅に改善されています。 Python スクリプトを介して omp ターゲットに変換できる同時実行コードがありますが、非常にうまく機能します。
確かに心配はしていますが、命を脅かすものではないと思います。重要なのは、ベンダーにFP64をサポートさせるためにFP64が必要なものを開発し続けることです。
あなたもこの懸念を共有しますか、それとも私がそれを誇張していると思いますか?
はい、共有しますが、誇張されているかどうかはわかりません…
昨年、私はモンペリエのシネスで開催されたセミナーに参加しました。

スーパー コンピューター Adastra についての素晴らしい話 https://mumps-solver.org/doc/workshop2025/Hautreux_Slides.pdf (フランス語)、議論の 1 つはまさにこの点でした。このセミナーはパリで開催された Nvidia GTC の数日後に開催され、nvidia はより低精度に取り組む計画について言及し、実際に AMD が FP64 HPC 負荷を擁護していることを示しました。
私の意見では、FP64 などを備えた HPC 専用の GPGPU は今後も存在します。30 年前やそれ以前の HPC ハードウェアが「ニッチ」であったのと同じように、それは一種の「ニッチ」ハードウェアになるでしょう。私たちは、HPC がコンシューマ ハードウェア (ハイエンドの CPU や GPU など、コンシューマ市場向けに設計されたもの、またはコンシューマ市場向けのバージョンから派生したもの) をほぼ使用できるようになる (比較的短い) 期間を生きてきましたが、もう終わりです。 HPC 専用のハードウェアは今後さらに高価になる可能性があります。
2026 年 6 月 17 日、午後 1 時 1 分
6
実際、私たちは現在、AI が GPU イノベーションの主要な「エンジン」である時代にいます。その結果、ハードウェアはそのタスクのために非常に特化されています。これは AI 研究にとっては素晴らしいことですが、従来の科学技術コンピューティング コミュニティが、特殊なハードウェアに割増料金を支払うか、ますます「AI ファースト」になる消費者グレードのカードの制限に対処しなければならない厄介な立場に置かれていることは間違いありません。
しかしその一方で、私が添付したリンクにある Nvidia の声明を見ると、将来は現在見られるものよりも少し悲劇的ではないかもしれません。これまで、私のような科学技術コンピューティングに携わる者は、自分の仕事に意図されていないリソースを頻繁に使用してきたと言わなければなりません。プレイステーションに Linux をインストールしました…今回も、間違って、ハードウェアが私たちにとっても役立つものに向かって進化することを願っています。
エヌビディア

HPC の改善をもたらす次世代 GPU FP64 のサポートを再確認
「Hopper」や「Blackwell」など、NVIDIA の最近の世代の GPU は、ハイ パフォーマンス コンピューティング (HPC) 指向の FP64 倍精度パフォーマンスが低下することで知られています。ただし、同社は本質的に HPC スペースを放棄しているわけではありません。
こんにちは、ありがとうございます。 NextSilicon さんの取り組みは知りませんでしたが、とても興味深いですね。
これは非常に役立ち、ベンダーから直接のフィードバックを得ることができるのは非常に励みになります。 Fortran の GPU サポートに関しては、これは私たちが毎日苦労して対処しなければならない別の問題です。別のトピックで詳しく説明できるかもしれません。
最近掲載された記事:
松岡聡、[2606.06510] 必要なのは FP8 だけ (パート 1): HPC の聖杯としてのハードウェア FP64 を暴く (6 月 13 日バージョン)
これは尾崎スキーム II に基づいています。
尾崎勝久、内野有紀、今村俊之、[2504.08009] 尾崎スキーム II: 整数モジュラー技術を使用した浮動小数点行列乗算の GEMM 指向エミュレーション
次のスキームを使用する GEMM エミュレーションがあります: GitHub - RIKEN-RCCS/GEMMul8: GEMMul8 (GEMMulate): GEMM エミュレーションと、Ozaki Scheme II に基づく INT8/FP8 マトリックス エンジンを使用した BLAS のような行列演算への拡張 · GitHub
このことから、高速科学 HPC には専用の FP64 ハードウェアは必要ないかもしれませんが、低精度のハードウェアを使用してエミュレートできることがわかります。
ありがとう、スライドは本当に興味深いです。 AMDでは信頼する必要があります
実際、これが私の懸念です。HPC が他のアプリケーションとハードウェアを共有するハイエンド CPU-GPU 時代 (ゲーム、仮想通貨マイニングなど、お金が存在する場所) を経て、HPC は古いコーナーに戻りつつあります。
このリンクを共有してくれてありがとう、私はそれを知りませんでした。皆さんは私よりも楽観的なので、考えを改めなければなりません。とにかく

、私は今でもNVIDIAの発表よりもAMDを信頼しています。
これらの参考資料を共有していただきありがとうございます。すぐに読みます。
私はすでに、低精度の表現と、精度を (少なくとも部分的に) 回復するための補正アルゴリズムを検討するために、いくつかのおもちゃのコードを試していますが、まだ証拠に基づいた強力な意見を持っていません。さらに、NVIDIA が主張する「エミュレーション モード」、特にそれが精度を確保し、パフォーマンスを維持するかどうかはわかりません (私にとっては両方の側面が重要です)。
私もこの懸念を共有しますが、それについて私にできることは何もないので、心配することのリストにはありません。私の中の楽観主義者は、いつか Nvidia、AMD、Intel が自社のコモディティ GPU で FP64 を有効にしてくれることを期待しています。私の中の現実主義者は、それをビジネスケースにするのは難しいので、おそらくそれが起こらないことを知っています。私はゲーマーではないので、SciVis を行うためにできるだけ多くのメモリを搭載した GPU が必要です。したがって、自分のニーズを満たすものが 300 ～ 400 ドルで手に入るのであれば、GPU に 1000 ドルを無駄にするつもりはありません。 1000 ドルのカードで FP64 サポートが得られれば状況は変わるかもしれません。起こらないだろうけど。
数値解析での FP32 の使用に関しては、Jack Dongarra のグループの最新の線形代数パッケージ (私の記憶が正しければ Magma) は、対話型リファインメントを使用して、マトリックス ソルバーに FP32 を使用したときに失われた精度の一部を回復していると思います。
これは、GPU での混合精度反復リファインメントの使用に関する Dongarra のグループの論文です。 MAGMA Web サイトには他にもいくつかの論文があります
松岡聡、[2606.06510] 必要なのは FP8 だけ (パート 1): HPC の聖杯としてのハードウェア FP64 を暴く (6 月 13 日バージョン)
参考資料をありがとうございます。ここで有名な引用が思い出されます。
畑を耕すとしたら、どちらを使いますか?二頭の強い牛

鶏は1024羽？
私はすでに、低精度の表現と精度を（少なくとも部分的に）回復するための補償アルゴリズムを探索するために、いくつかのおもちゃのコードで遊んでいます。
精度を下げた命令を使用して高精度の計算をシミュレートすることはいつでも可能ですが、それは複雑さの層が「1 つ増えただけ」です。しかし、これらの層はベンダー ライブラリ (LAPACK、FFT など) の一部であり、ユーザーに対して透過的であると想像できます。
Discourse によって提供されており、JavaScript を有効にして表示するのが最適です

## Original Extract

Hi all.
I have one, maybe two, questions for you. This question came out of a webinar series on High Performance Computing (HPC) I took part in (the Italy–Germany HPC webinars organised on the Italian side through CNR).&hellip;

= 40rem)" rel="stylesheet" data-target="desktop" />
= 40rem)" rel="stylesheet" data-target="chat_desktop" />
= 40rem)" rel="stylesheet" data-target="discourse-ai_desktop" />
= 40rem)" rel="stylesheet" data-target="discourse-reactions_desktop" />
= 40rem)" rel="stylesheet" data-target="poll_desktop" />
Fortran Discourse
The AI era is pulling FP64 hardware away from scientific HPC
szaghi
June 17, 2026, 9:31am
1
Hi all.
I have one, maybe two, questions for you. This question came out of a webinar series on High Performance Computing (HPC) I took part in (the Italy–Germany HPC webinars organised on the Italian side through CNR). I raised this concern there, and my impression was that the other speakers did not share it to the same degree. The room leaned more optimistic than I am. That is exactly why I want to put it to a wider audience: I may be wrong, and I would like to hear where others land.
The concern is precision. Most scientific HPC needs double precision (FP64). In computational fluid dynamics, which is my field, we resolve physical scales spanning many orders of magnitude, and to do that correctly (with very high-order accuracy methods), we need 64-bit floating point.
AI computing does not need this. Training and inference work well at 8-bit (now even at 4-bit). So, the two workloads require different hardware: AI needs many low-precision cores, while science requires strong FP64 capabilities.
The problem is that the vendors follow the AI market because that is where the money is. Comparing on vector FP64 (peak, dense), the recent trend is to hold it flat or lower it, and spend the transistors on low-precision math instead:
NVIDIA H100: 34 TFLOP/s vector FP64, or 67 with the FP64 tensor-core path. The newer B200 does about 40 vector FP64. Blackwell dropped the dedicated FP64 tensor-core path that Hopper had, and gained around 20 PFLOP/s of FP4 for AI. The Rubin roadmap reportedly cuts FP64 further.
AMD MI300X: 81.7 TFLOP/s FP64. The newer MI355X does 78.6, below its own predecessor , with the gains all in FP8/FP4 for AI inference.
Intel has stepped back from a dedicated HPC GPU. Its current HPC silicon, the Max-series (Ponte Vecchio) in Aurora, has no standalone successor. Intel cancelled Falcon Shores as a product in early 2025 and folded its HPC and AI lines into one chip, Jaguar Shores, due around 2026/2027. Intel describes it as serving both AI and HPC, but says it will compete on total cost of ownership rather than peak FLOPS, and has published no FP64 figure.
Consumer silicon makes the direction plainest. NVIDIA’s N1X, the new Blackwell laptop chip, publishes only AI-precision figures (NVFP4, around 1000 TOPS) and quotes no FP64 at all. Double precision is simply not a design goal there.
So across all three vendors the direction looks the same. The new chips are built for AI, and double precision gets quietly de-prioritized along the way.
There is one strong counter-current. AMD’s MI430X, coming this year, is a deliberate HPC part. AMD claims more than 200 TFLOP/s of FP64, and independent estimates back out around 211 from the Alice Recoque exascale contract, which would be the highest of any GPU so far, while it still carries FP4/FP8 for AI. It will power Alice Recoque, the next European exascale machine, alongside the US Discovery and Germany’s Herder. So a dedicated FP64 line still exists, for now.
But it is one product line, from one vendor, against a whole market moving the other way. That is what I cannot resolve: whether a first-class FP64 hardware line survives, or shrinks to a small premium niche while everything else is optimized for AI.
Do you share this concern, or do you think I am overstating it?
If you share it, do you already see a way out?
I would be glad to hear how others in the Fortran and HPC community are thinking about this.
There is also NextSilicon’s Maverick-2 in Sandia’s supercomputer Spectra .
This is in my mind constantly. However, I’ve talked to people that work at either nvidia or AMD and they’ve all assured me that FP64 is not going to be dropped like a hot potato. Cards are separating into two lines, the B300 for example is for AI/inference procedures while the B200 is going to be for science, HPC stuff.
The thing that is true is that we won’t see the crazy increase in FP64 performance, i.e. V100 FP64 was 7.9 TFLOP/s, A100 was 19 (??), H200 was 34 ish and the B200 is as you pointed out less than that creep at 40 TFLOP/s.
The thing here is that most of the workloads out there are not using the peak FLOP rate of these things, not everything is DGEMMs there are a limited amount of applications that can really run these. Most of other apps I feel are limited by memory bandwidth which is important to both AI and HPC workflows. AMD however seems to be championing the FP64 support more than NVIDIA, but GPU support for Fortran is still not the best but it is improving a LOT. I have my do concurrent code which I can translate via a python script to omp target and it works very nicely.
So yeah I am concerned but I don’t see it as a life threatening. The thing is to keep developing things that NEED FP64 to force vendors to support it.
Do you share this concern, or do you think I am overstating it?
Yes, I share it, can’t say if it is overstated or not …
Last year I attended a seminar in Montpellier at CINES and there was a very nice talk about their super computer Adastra https://mumps-solver.org/doc/workshop2025/Hautreux_Slides.pdf (in French), one of the discussions was precisely this point. The seminar happened a few days after the Nvidia GTC in Paris where nvidia mentioned their plans to go more into lower precision, they showed that indeed AMD is the one championing FP64 HPC loads.
IMO there will still be GPGPU dedicated to HPC, with FP64 etc… Just it will become a kind of “niche” hardware, just like the HPC hardware was “niche” 30 years ago and before. We have lived a (relatively short) window where HPC could almost use consumer hardware (i.e. high-end CPU’s and GPU’s, but which were designed for the consumer market, or which were derived from versions for the consumer market), but it’s over. The hardware dedicated to HPC will probably be more expensive from now.
June 17, 2026, 1:01pm
6
Indeed, we are currently in a period where AI is the primary “engine” of GPU innovation, and as a result, hardware is being hyper-specialized for that task. While this is fantastic for AI research, it undeniably leaves the traditional scientific computing community in an awkward position where they must either pay a premium for specialized hardware or deal with the limitations of consumer-grade cards that are increasingly “AI-first.”
On the other hand, however, looking at Nvidia’s statements in the link I attach, the future may be a little less tragic than what is currently in sight. It must be said that in the past, those who work in scientific computing like me have often used resources that weren’t intended for my work: we installed Linux on a PlayStation… I hope that this time too, by mistake, the hardware will evolve towards something that can be useful to us too.
NVIDIA Reaffirms Support for FP64, Next-Gen GPU to Bring HPC Improvements
NVIDIA's recent generations of GPUs, such as "Hopper" and "Blackwell," have been known for stalling their high-performance computing (HPC) oriented FP64 double-precision performance. However, the company is not inherently abandoning the HPC space and...
Hi, thank you very much. I did not know NextSilicon’s work, it is really interesting.
This helps a lot, having first-hand feedback from the vendor is quite encouraging. As for the GPU support for Fortran, this is all another problem that we have to daily struggle to deal with, maybe we can discuss in detail in another topic
This article appeared recently:
Satoshi Matsuoka, [2606.06510] FP8 is All You Need (Part 1): Debunking Hardware FP64 as the HPC Holy Grail (June 13th version)
It is based on the Ozaki Scheme II:
Katsuhisa Ozaki, Yuki Uchino, Toshiyuki Imamura, [2504.08009] Ozaki Scheme II: A GEMM-oriented emulation of floating-point matrix multiplication using an integer modular technique
There is an GEMM emulation which uses the scheme: GitHub - RIKEN-RCCS/GEMMul8: GEMMul8 (GEMMulate): GEMM emulation and its extension to BLAS-like matrix operations using INT8/FP8 matrix engines based on the Ozaki Scheme II · GitHub
It appears from this that dedicated FP64 hardware may not be required for fast scientific HPC but can be emulated using low precision hardware.
Thank you, the slides are really interesting. In AMD we have to trust
Indeed, this is my concern, after the high-end CPU-GPU era where HPC shares hardware with other applications (where money is, e.g., gaming, crypto-mining, etc), HPC is coming back into its old corner.
Thank you for sharing this link, I did not know it. All of you are more optimistic than me, so I have to change my mind. Anyhow, I still trust more in AMD than in NVIDIA announcement
Thank you for sharing these references. I’ll read them soon.
I am already playing with some toy code for exploring low-precision representation and compensative algorithms to recover (at least part) of the accuracy, but I do not yet have a strong, evidence-based opinion. Moreover, I do not know the “emulation mode” claimed by NVIDIA, in particular, whether it ensures precision and preserves performance (both aspects are important for me).
I share this concern also but there is not anything I can do about it so its not at my list of things to worry about. The optimist in me hopes that one day Nvidia, AMD, and Intel will enable FP64 in their commodity GPUS. The realist in me knows thats probably never going to happen because its hard to make a business case for it. I’m not a gamer and just need a GPU with as much memory as I can afford to do SciVis. Therefore, I’m not going to waste a $1000 on a GPU when I can get something that meets my needs for around $300 to $400. That might change if I could get FP64 support in the $1000 card. Not going to happen though.
As to using FP32 for numerical analysis, I think that Jack Dongarra’s group’s latest Linear Algebra packages (Magma if I remember correctly) uses interative refinement to recover some of the accuracy lost using FP32 for their matrix solvers.
Here is a paper from Dongarra’s group on using mixed precision iterative refinement on GPU’s. There are several other papers on the MAGMA web site
Satoshi Matsuoka, [2606.06510] FP8 is All You Need (Part 1): Debunking Hardware FP64 as the HPC Holy Grail (June 13th version)
Thanks for the references. A famous quote comes to mind here:
If you were plowing a field, which would you rather use? Two strong oxen or 1024 chickens?
I am already playing with some toy code for exploring low-precision representation and compensative algorithms to recover (at least part) of the accuracy
It’s always possible to simulate high precision computations with reduced precision instructions, it’s “just” one more layer of complexity. But one can imagine that these layers are part of the vendor libraries (for LAPACK, FFTs,…) and transparent to the user.
Powered by Discourse , best viewed with JavaScript enabled
