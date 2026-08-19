---
source: "https://www.anthropic.com/research/Claude-accelerates-protein-design"
hn_url: "https://news.ycombinator.com/item?id=49356105"
title: "How Claude is accelerating protein design and analytical chemistry"
article_title: "How Claude is accelerating protein design and analytical chemistry \\ Anthropic"
image: "https://cdn.sanity.io/images/4zrzovbb/website/e3758f1bc27af0786f4249cc1ab194fc2c6cce63-3840x2160.png"
author: "starshadowx2"
captured_at: "2026-08-19T03:39:00Z"
capture_tool: "hn-digest"
hn_id: 49356105
score: 1
comments: 0
posted_at: "2026-08-19T03:00:42Z"
tags:
  - hacker-news
  - translated
---

# How Claude is accelerating protein design and analytical chemistry

- HN: [49356105](https://news.ycombinator.com/item?id=49356105)
- Source: [www.anthropic.com](https://www.anthropic.com/research/Claude-accelerates-protein-design)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T03:00:42Z

## Translation

タイトル: クロードはタンパク質設計と分析化学をどのように加速しているか
記事のタイトル: クロードはタンパク質設計と分析化学をどのように加速しているか \ Anthropic
説明: この投稿では、生命科学者が研究のペースを速めるためにクロードがどのように役立つかを示す 2 つの結果を共有します。最初の実験では、タンパク質結合剤をゼロから設計するクロードの能力をテストしました。これは、専門家がこれまで数週間または数か月を費やしてきたタンパク質ベースの医薬品を作成するための重要なステップです。
[切り捨てられた]

記事本文:
クロードがタンパク質設計と分析化学をどのように加速しているか \ Anthropic メインコンテンツにスキップ フッターにスキップ 研究
科学 クロードがタンパク質設計と分析化学をどのように加速させているか
概要: この投稿では、生命科学者が研究のペースを速めるためにクロードがどのように役立つかを示す 2 つの結果を紹介します。最初の実験では、タンパク質結合剤をゼロから設計するクロードの能力をテストしました。これは医薬品設計プロセスの初期段階を代表する重要なタスクであり、これまで専門家が標的ごとに数週間から数か月を要してきたタスクです。 Claude (Mythos Preview および Opus 4.8) は 15 の標的に対するタンパク質バインダーを設計し、そのうち 14 に対して成功しました。セットアップに応じて、個々のデザインの 22% ～ 35% が正常にバインドされました。これに対し、今日のタンパク質設計キャンペーンでは通常 10 ～ 15% です。その最も強力な設計のいくつかは、以前に公開された最良の結果よりも数倍しっかりと結合しました。 2 番目の例では、Claude が化学分析を加速できるかどうかを評価しました。一般的に入手可能なモデルである Claude Opus 5 には、NMR および LC-MS データ (化学者が扱う化合物の正体と純度を評価できるデータ) が与えられました。契約ラボの生ファイルと 2 文のプロンプトのみを提供されたクロードは、水素数と純度に関するラボ独自の分析 (96.4% 対 96.33%) と一致する完成結果を 23 分と 19 分で返しました。これらの例は、複雑な科学的タスクを進めるために現在必要とされている時間と計算の専門知識をクロードがどのように削減できるかを示しています。
AI を活用した発見のペースはここ数カ月で加速しています。これらの発見の大部分は、検証が比較的早い分野で行われています。たとえば数学では、エージェントが働き始めています。

未解決の問題を解決する方法: 何十年も続いたエルデシュの問題は月に数回のペースで減少しており、私たちは最近、クロードがリーマン ゼータ関数の長年にわたる下限をどのように改善したかを共有しました。
AI モデルは、ライフサイエンスなど、結果の検証がより複雑でコストがかかる実験分野でも進歩を加速し始めています。この投稿では、クロードの科学的能力に関する 2 つの実験の結果を共有します。まず、タンパク質設計キャンペーンにおけるクロードのパフォーマンスに関する調査結果を提示します。これは、クロードが人間の主要な専門家と同様に (またはそれ以上に) さまざまな標的に対するタンパク質バインダーを設計できることを示しています。次に、Claude Opus 5 が分析化学タスクでどのように実行されたかを共有し、一般アクセス モデルが研究の日常的で時間のかかる側面をどのようにサポートできるかを示します。
以下に説明するタンパク質設計および分析化学のタスクは、医薬品開発プロセスの初期段階の一部を構成する作業の代表的なものです。これらの段階の加速は、医薬品開発をエンドツーエンドでスピードアップするという私たちのより大規模な取り組みの一部であり、その多くの側面は中核となる科学的能力の向上よりも、政策や運用上のボトルネックに関係しています。
今日共有する結果は、Mythos モデルと Opus モデルを組み合わせて得られたものです。現在、生命科学の研究タスクは当社の最も有能なモデルでブロックされていますが、当社の最優先事項の 1 つは科学者向けのアクセス プログラムを開始することであり、これについては間もなく詳細を共有できる予定です。それまでのところ、Opus 5 は依然として最も高性能な一般入手可能なモデルです。
Claude Mythos 5 を発表したとき、私たちはこのモデルにアクセスするために実験していることを共有しました。

医薬品設計プロセスの一部を学びます。この研究の継続的な一環として、私たちは複数のタンパク質標的に対するミニバインダーを設計するクロードの能力を調査してきました。ミニバインダーは、標的タンパク質にしっかりと結合するように設計された小さなタンパク質です。結合とは、現代の医薬品の大部分がどのように機能するかです。つまり、標的に結合して、標的に何かを阻害、活性化、または送達します。新しいバインダーの設計 (de novo 設計と呼ばれる) には、これまでタンパク質エンジニアがターゲットごとに数か月の計算、最​​適化、スクリーニングを要してきました。
近年、タンパク質を設計し、結合する可能性が最も高いものをランク付けできる機械学習モデルにより、タンパク質の設計プロセスが大幅に迅速化されました。しかし、これらのモデルは依然として、計算の専門家による数日 (場合によっては数週間) の骨の折れる調整を必要とします。そして、クロードのような一般的な推論モデルは、専門家と非専門家の両方がコンピュータでタンパク質をより効率的に設計するのに役立ちますが、ウェットラボ（科学者が化学物質、薬物、その他の生物学的物質を物理的に試験する場所）でそのデータを検証するにはまだ数週間かかります。
私たちは現在、これらの実験の最初の実験である、Claude Opus 4.8 と Mythos Preview を使用した 15 のターゲットに対するマルチアームタンパク質設計キャンペーンのウェット ラボ データを受け取りました。当社の外部評価者である Adaptyv Bio と Twist Bioscience は、クロードの設計を研究室で独自に作成およびテストし、私たちが設計した 15 のターゲットのうち、クロードがそのうち 14 のターゲットに対するバインダーの設計に成功したことを発見しました。これらには、少なくとも 6 つの標的に対する高親和性バインダー 1、および少なくとも 4 つの標的に対して報告されている最良の親和性と同等またはそれを超えるバインダーが含まれます。親和性は、タンパク質がその標的にどれだけ強く結合するかを示す尺度です。治療効果を達成するには、一般に高親和性結合剤が必要です

なぜなら、それらはより低用量で薬の効果を高め、副作用のリスクと製造コストを削減するからです。
Mythos Preview と Opus 4.8 は、48 時間のセッションですべてのターゲットに対して同時にデザインを行った場合、全体のヒット率 (実際にバインダーとなるデザインの数) をそれぞれ 26.7% と 22.6% に達しました。今日のタンパク質設計キャンペーンでは 10 ～ 15% が一般的です。 2
複数のターゲットに対して設計するクロードの能力を評価した後、特にタンパク質エンジニアが一般的に採用するアプローチをよりよく表していることを考慮すると、一度に 1 つのターゲットに焦点を当てることでパフォーマンスが向上するかどうかを理解したいと考えました。実際、複数の 24 時間セッションを使用して各ターゲットに対して個別に設計した場合、Mythos Preview は全体のヒット率 35.1% を達成することがわかりました。
このキャンペーンは、最初のプロンプト (リンク) でクロードに提供した情報を超えて、人間の関与を最小限に抑えて実行されました 3。私たちは、このアプローチが専門のタンパク質設計者の手に渡れば、特に中間結果についてクロードに積極的な指導とフィードバックを与えた場合には、さらに強力な結果が得られると期待しています。
私たちは、Adaptyv Bio の BenchBB のすべてを含む、タンパク質設計ベンチマークで一般的に使用される複数のターゲット 4 を選択することからタンパク質設計キャンペーンを開始しました。これらのターゲットは広範囲に調査されているため、結果を公開されているヒット率およびアフィニティと比較できます。また、Adaptyv Bio の最新のコンペティションから 2 つの新しいターゲット、15-PGDH と GDF-8 を選択し、クロードがトレーニング データやオンライン検索に事前に記録された成功を利用することなく、ターゲットに合わせて設計できるようにしました (すべてのターゲットについて、クロードにその設計がオリジナルであることを確認して確認するよう要求しました)。
次にクロードにデザインを依頼しました。

Claude Science では、これらの標的に対するタンパク質結合剤を研究しています。このために、私たちは 2 つのアプローチを採用しました。 1 つ目はマルチターゲット モードで、クロードは単一のクロード サイエンス セッションですべてのターゲットに対して同時に設計を行いました。 2 つ目は単一ターゲット モードで、各セッションが 1 つのターゲットに対処し、すべてのターゲットのセッションが並行して実行されました。
Opus 4.8 と Mythos Preview をマルチターゲット モードで実行し、特殊なタンパク質設計と折りたたみモデルを実行するために 48 時間の実行時間と最大 12,500 NVIDIA H100 時間のコンピューティングを実行しました。また、Mythos Preview をシングルターゲット モードで実行し、ウォール タイムは 24 時間、ターゲットごとに最大 2,500 NVIDIA H100 時間のコンピューティングを実行しました。 5
典型的なタンパク質設計キャンペーン中に利用可能なリソースをエミュレートするために、Claude に次のものを与えました。
広範なタンパク質設計により、エージェントのコンテキストにも含まれるプロンプト 6 が作成されました。
インターネットおよびタンパク質設計に関する論文などのリソースのコーパスへのアクセス。
Google Drive、Slack、Gmail、BioRxiv 用のコネクタ。
特殊なタンパク質設計とフォールディング モデルを実行するための GPU へのアクセス。
割り当てられた時間内でのトークンとサブエージェントの予算に制限はなく、高速モードが有効になっています。
クロードにプロンプ​​トを与えた後、モデルが自律的に実行されるようにしました。キャンペーンを開始した後、私たちは追加の科学的、技術的、または運用上のガイダンスを提供しませんでした。
私たちが関与したのは、アクセス承認 (ネットワーク アクセス要求など) を付与することと、セッションが実行されていることを確認するためにインフラストラクチャを監視することだけでした。クロードは、人間のオペレーターが数週間かかるバインダーの設計に必要なすべての作業を実施しました。各タンパク質ターゲットのどこを対象に設計するかを選択しました。いくつかの構造設計、配列設計、および共折り畳みモデルを調整することにより、候補構造および配列を生成しました (mo

タンパク質の構造とそれが結合するものを 1 回のパスで予測するデル)。インシリコ最適化の複数サイクルを通じて設計を実行しました。そして、発現し、可溶性を維持し、結合する新規で多様な候補をコンピューターによってスクリーニングします。
15 のターゲットごとに、私たちはクロードに 30 のタンパク質バインダーを設計するよう依頼しました。クロードは、公開されている専門家によるタンパク質設計と、この分野ですでに使用されている共折り畳みモデルを操作することでこれを行いました。その後、クロードの設計は検証のために Adaptyv Bio と Twist Bioscience に送られました。
ターゲットに対するクロードのパフォーマンス
この取り組みの終了までに、合計 1,320 のデザインを使用して、15 ターゲット中 14 に対して 354 のバインダーを作成しました。これは、公的に利用可能な de novo タンパク質設計の総コーパスへの重要な貢献を表しています。たとえば、2 つの最大のコレクション、proteinbase.com と Overath らによってキュレーションされたコレクションです。 、40 のターゲットに対する 5,700 のデザインから約 770 のバインダーで構成されます。以下に、Claude の機能を強調する 3 つの例と、その限界を示す 1 つの例を紹介します。詳細については、技術レポート (リンク) をご覧ください。
Adaptyv Bio がコンテストを開催したターゲットについて、クロードはヒット率と親和性の両方において上位参加者のレベル以上のパフォーマンスを示していることがわかりました。 RBX1 (特定の制御タンパク質の標的破壊を促進する小さなタンパク質) に対して、シングルターゲット モードの Mythos Preview では 40% のヒット率を達成しましたが、参加者のヒット率は 3.7% でした。その最高位のデザインは、応募された 245 デザインの中で、受賞デザインを上回る高親和性バインダーでした。
興味深いことに、Mythos Preview ではなく Opus 4.8 は、複数の専門家グループが苦労してきたターゲットである TNFα に成功しています。 TNFαはシグナル伝達タンパク質です。

炎症を引き起こすのは免疫系によって緩和され、それを阻止することが、ヒュミラを含むこれまでに作られた最も影響力のあるいくつかの薬の治療の基礎となっています。これは多量体構造のため、設計が難しい標的であり、2 つのタンパク質によって形成される溝の結合部位を標的にする必要があります。 Mythos Preview は失敗に終わりましたが、Opus 4.8 では、複数のバインダーを設計しました。その中には、種を超えて機能し、ヒト、カニクイザル、マウスの TNFα に結合するものも含まれており、これは動物研究を行う上で重要です。 Opus 4.8 がこのターゲットで成功し、Mythos Preview が成功しなかった理由はわかりません。モデルの機能を評価するときは、総合的に評価します。タンパク質設計に固有の複雑さを考慮すると、全体的に能力が低いモデルでも、一般に能力が高いモデルよりも優れたパフォーマンスを発揮できる特定の領域が存在することは驚くべきことではありません。
コンピューターによって設計されたバインダーのほとんどは、コイルからなるタンパク質の二次構造であるαヘリックスの束です。伸長したアミノ酸鎖が横に並んでいる必要があるβシートは、設計が難しく、誤った折り畳みや凝集（タンパク質分子が分離して適切に折り畳まれた状態を保つのではなく、互いにくっついている場合）が起こりやすくなります。クロード・デ

[切り捨てられた]

## Original Extract

In this post, we share two results that show how Claude can help life scientists increase the pace of their research. In the first, we tested Claude’s ability to design protein binders from scratch, a key step in creating protein-based drugs that has historically taken a specialist weeks or months p
[truncated]

How Claude is accelerating protein design and analytical chemistry \ Anthropic Skip to main content Skip to footer Research
Science How Claude is accelerating protein design and analytical chemistry
Summary: In this post, we share two results that show how Claude can help life scientists increase the pace of their research. In the first, we tested Claude’s ability to design protein binders from scratch, a key task representative of the early parts of the drug design process and one that has historically taken a specialist weeks or months per target. Claude (Mythos Preview and Opus 4.8) designed protein binders against 15 targets, and succeeded against 14 of them. Between 22% and 35% of its individual designs bound successfully, depending on the setup, compared to the 10-15% that is typical in protein design campaigns today. Some of its strongest designs bound several times more tightly than the best previously published result. In the second example, we evaluated whether Claude can accelerate chemical analysis. Claude Opus 5, a generally available model, was given NMR and LC-MS data (the data that allows chemists to assess the identity and purity of the compounds they work with). Provided with only a contract lab’s raw files and a two-sentence prompt, Claude returned finished results in 23 and 19 minutes, matching the lab’s own analysis on hydrogen counts and purity (96.4% versus 96.33%). These examples demonstrate how Claude can reduce the time and computational expertise currently required to make progress on complex scientific tasks.
The pace of AI-enabled discoveries has quickened over the past few months. The bulk of these discoveries have been in areas where verification is relatively fast. In mathematics, for example, agents have begun to work their way through unsolved problems: Erdős problems that have stood for decades are falling at a rate of several a month, and we recently shared how Claude improved on a longstanding lower bound on the Riemann zeta function .
AI models are also beginning to hasten progress in experimental fields where verifying the results is more complex and expensive, such as in the life sciences. In this post, we share the results of two experiments into Claude’s scientific capabilities. First, we present findings from our investigation into Claude’s performance on a protein design campaign, showing that Claude can design protein binders against a variety of targets as well as (or even better than) leading human experts. Second, we share how Claude Opus 5 performed on an analytical chemistry task, demonstrating how general-access models can support the routine and time-intensive aspects of research.
The protein design and analytical chemistry tasks described below are representative of the work that makes up some parts of the early stages of the drug development process. Accelerating these phases is one component of our much larger effort to speed up drug development end-to-end, many aspects of which have more to do with policy and operational bottlenecks than with improvements in core scientific capabilities.
The results that we’re sharing today were obtained with a combination of our Mythos and Opus models. While life science research tasks are currently blocked in our most capable model, one of our highest priorities is to launch an access program for scientists, and we expect to share more on this soon. In the meantime, Opus 5 remains our most capable generally available model.
When we announced Claude Mythos 5 , we shared that we were experimenting with the model to accelerate parts of the drug design process. As an ongoing part of this work, we have been investigating Claude’s ability to design minibinders for multiple protein targets. A minibinder is a small protein designed to latch tightly onto a target protein. Binding is how a large proportion of modern medicines work: they attach to a target and inhibit, activate, or deliver something to it. Designing a new binder (known as de novo design) has historically taken protein engineers months of computation, optimization, and screening per target.
In recent years, machine-learning models that can design proteins and rank which are most likely to bind have greatly expedited the protein design process. But these models still generally require days (and often weeks) of laborious orchestration by computational experts. And although general reasoning models like Claude can help both experts and non-experts more efficiently design proteins computationally, validating that data in a wet lab (where scientists physically test chemicals, drugs, and other biological substances) still takes weeks.
We have now received wet lab data back for the first of these experiments, a multi-arm protein design campaign against 15 targets using Claude Opus 4.8 and Mythos Preview. Our external evaluators, Adaptyv Bio and Twist Bioscience , independently produced and tested Claude’s designs in the lab, finding that of the 15 targets we designed against, Claude successfully designed binders against 14 of them. These include high-affinity binders 1 against at least six targets, and binders matching or exceeding the best reported affinity against at least four targets. Affinity is a measure of how strongly a protein binds to its target; high-affinity binders are generally needed to achieve a therapeutic effect because they make the drug effective at lower doses, reducing the risk of side effects and the cost to manufacture them.
Mythos Preview and Opus 4.8 achieve overall hit rates—how many of the designs are, in fact, binders—of 26.7% and 22.6%, respectively, when designing against all targets simultaneously in a 48-hour session. 10 to 15% is typical in protein design campaigns today. 2
After assessing Claude’s ability to design against multiple targets, we wanted to understand whether having it focus on a single target at a time would improve its performance, especially given that this better represents the approach typically taken by a protein engineer. Indeed, we found that Mythos Preview achieves an overall hit rate of 35.1% when designing against each target separately using multiple 24-hour sessions.
This campaign was carried out with minimal human involvement 3 beyond the information we provided Claude in our initial prompt ( link ). We expect that in the hands of expert protein designers this approach would yield even stronger results, especially if they give Claude active guidance and feedback on intermediate results.
We began our protein design campaign by selecting multiple targets 4 that are commonly used in protein design benchmarks, including all of Adaptyv Bio’s BenchBB . Because these targets have been studied extensively, we can compare our results against published hit rates and affinities. We also chose two novel targets, 15-PGDH and GDF-8 , from Adaptyv Bio’s most recent competitions to ensure Claude was able to design against targets without drawing upon pre-recorded successes in its training data or from online search (for all targets, we required Claude to check for and ensure that its designs were original).
We then prompted Claude to design protein binders against these targets in Claude Science . For this, we took two approaches. The first was a multi-target mode, where Claude designed against all targets simultaneously in a single Claude Science session. The second was a single-target mode, in which each session addressed one target and sessions for all targets ran in parallel.
We ran Opus 4.8 and Mythos Preview in multi-target mode with 48 hours of wall time and up to 12,500 NVIDIA H100 hours of compute for running specialized protein design and folding models. We also ran Mythos Preview in single-target mode with 24 hours of wall time and up to 2,500 NVIDIA H100 hours of compute for each target. 5
To emulate the resources available during a typical protein design campaign, we gave Claude the following:
An extensive protein design prompt 6 that was also included in the agent context;
Access to the internet and a corpus of resources, such as papers, on protein design;
Connectors for Google Drive, Slack, Gmail, and BioRxiv;
Access to GPUs for running specialized protein design and folding models;
No limits on token and sub-agent budget within the allotted time, and fast mode enabled.
After giving Claude the prompt, we left the model to execute autonomously. We provided no additional scientific, technical, or operational guidance after we initiated the campaigns.
Our only involvement was granting access approvals (such as network access requests) and monitoring the infrastructure to ensure the sessions were running. Claude conducted all of the work that goes into designing a binder, which can take a human operator weeks. It chose where on each protein target to design against; generated candidate structures and sequences by orchestrating several structure design, sequence design, and co-folding models (models that predict the structure of a protein, together with whatever it binds, in a single pass); ran the designs through multiple cycles of in silico optimization; and computationally screened for novel, diverse candidates that would express, stay soluble, and bind.
For each of the 15 targets, we asked Claude to design 30 protein binders. Claude did this by operating publicly available specialist protein design and co-folding models that the field already uses. Claude’s designs were then sent to Adaptyv Bio and Twist Bioscience to validate.
Claude’s performance on the targets
By the end of this effort, we produced 354 binders against 14 of 15 targets using a total of 1,320 designs. This represents a significant contribution to the total corpus of publicly available de novo protein designs; for example, the two largest collections, proteinbase.com and the collection curated by Overath et al. , consist of approximately 770 binders out of 5,700 designs against 40 targets. Below, we share three examples highlighting Claude’s capabilities, and one showing its limitations. You can find more detail in our technical report ( link ).
We found that for the targets Adaptyv Bio has run competitions for, Claude performs at or beyond the level of the top participants on both hit rate and affinity. Against RBX1 (a small protein that drives the targeted destruction of specific regulatory proteins), Mythos Preview in single-target mode achieved a 40% hit rate, compared to a 3.7% hit rate among participants. Its top-ranked design was a high-affinity binder that outperformed the winning design, which was among 245 designs entered.
Interestingly, Opus 4.8, and not Mythos Preview, succeeds on TNFα, a target multiple expert groups have struggled with. TNFα is a signaling protein released by the immune system to trigger inflammation, and blocking it is the therapeutic basis for some of the most impactful drugs ever made, including Humira. It’s a challenging target to design against because of its multimeric structure, which requires targeting a binding site in the groove formed by two proteins. Although Mythos Preview was unsuccessful, Opus 4.8 designed multiple binders, including some that worked across species, binding human, cynomolgus monkey, and mouse TNFα, which is important for conducting animal studies. We’re not sure why Opus 4.8 was successful on this target and Mythos Preview was not. When we assess our models capabilities, we do so holistically. Given the inherent complexity of protein design, it’s unsurprising that there would be specific areas where an overall less capable model could still outperform one that was generally more capable.
Most computationally designed binders are bundles of α-helices, a protein secondary structure consisting of coils. β-sheets, in which extended strands of amino acids must line up side by side, are harder to design and more prone to misfolding and aggregation (when protein molecules stick to each other instead of staying separate and properly folded). Claude des

[truncated]
