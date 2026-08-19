---
source: "https://www.adaptyvbio.com/blog/anthropic-1"
hn_url: "https://news.ycombinator.com/item?id=49354929"
title: "Testing Claude-designed proteins in the wet lab"
article_title: "Case study: Benchmarking Claude’s protein designs in the wet lab | Adaptyv Bio"
image: "https://adaptyvbio.com/api/notion-image?pageId=3bb5ca69-e7be-8175-905f-c8a3473ef5dc"
author: "julian_englert"
captured_at: "2026-08-19T00:39:06Z"
capture_tool: "hn-digest"
hn_id: 49354929
score: 1
comments: 1
posted_at: "2026-08-19T00:25:34Z"
tags:
  - hacker-news
  - translated
---

# Testing Claude-designed proteins in the wet lab

- HN: [49354929](https://news.ycombinator.com/item?id=49354929)
- Source: [www.adaptyvbio.com](https://www.adaptyvbio.com/blog/anthropic-1)
- Score: 1
- Comments: 1
- Posted: 2026-08-19T00:25:34Z

## Translation

タイトル: クロードが設計したタンパク質をウェットラボでテストする
記事のタイトル: ケーススタディ: ウェット ラボでのクロードのタンパク質設計のベンチマーク |アダプティフ バイオ
説明: Anthropic チームは、タンパク質設計において最新のクロード モデルのベンチマークを行いたいと考えていました。彼らは、Claude Science で一連のモデルを実行して 16 の標的に対するタンパク質バインダーを設計し、実験検証のためにそれらを私たちの研究室に送りました。結果は、クロードモデルがプロットにおいてエキスパートレベルのスキルを示していることを示しています。
[切り捨てられた]

記事本文:
ケーススタディ: ウェットラボでのクロードのタンパク質設計のベンチマーク | Adaptyv Bio Services Binding Expression Thermostability Blog Team Careers API Proteinbase Docs ミーティングを予約する タンパク質をテストする サービス ブログ Team Careers API Proteinbase Docs ミーティングを予約する タンパク質をテストする Foundry にサインイン Adaptyv Bio Biopole Life Science Campus, Route de la Corniche 5, 1066 Epalinges, Lausanne, Switzerland
← ブログに戻る ケーススタディ: ウェットラボでのクロードのタンパク質設計のベンチマーク
- Anthropic チームは、Claude Science の Claude を使用して 16 のターゲットに対するタンパク質バインダーを設計し、それを私たちの研究室に送りました。
- 当社は、デジタルタンパク質配列から DNA、発現タンパク質、そして測定と品質管理に至るまで、すべての結合実験を自動ウェットラボで実行しました。
- 1,320 のデザインのうち 354 がターゲットを拘束し、クロードの平均命中率は 26.8% となりました。単一のターゲットのみがバインダーを生成しませんでした。
- タンパク質設計コンペティションの結果と比較すると、クロードの設計は成功率が著しく高く、6 つ中 5 つを獲得し、より強固な結合剤が得られたでしょう。
- この記事の最後では、薬剤科学、タンパク質設計、そしてすべての病気を治すための長い道のりについて、私たちがどのように考えているかについてさらに詳しく説明します。
生物工学の約束は、病気の検出に役立つ診断法、癌を破壊できる抗体、ウイルスから守る新しいワクチンなど、新しい生物学的ツールを設計することです。これらのツールを構築するには、最終的に AI は静的なデータセットだけでなく、物理世界と対話する必要があります。
生物学における AI の次の段階は、エージェントティック サイエンス ループです。これは、科学的知識を推論し、特殊なツールを使用して数千の分子を設計し、それらの設計を自動化されたウェットラボに送信し、得られたデータから学習できる AI システムです。

。
次の 3 つのレイヤーが連携する必要があります。
科学的目標を解釈し、キャンペーンを計画し、ワークフローを調整するジェネラリスト モデル。
タンパク質構造予測やタンパク質設計モデルなどのタスクを実行するための専門モデルとツール。
デジタル設計を物理世界からの測定値に変換する自動化されたウェット ラボ。
この投稿では、Anthropic が自動ウェット ラボでテストした新しいタンパク質バインダーの設計において、Claude Mythos Preview と Opus 4.8 をどのようにベンチマークしたのかを紹介します。全体として、クロード モデルはタンパク質設計において専門家レベルのスキルを示し、多くのタスクにおいて人間の専門家と同等またはそれを上回っていました。 Anthropic チームと協力して、私たちはクロードが設計した実際のタンパク質配列と実験データを、オープンタンパク質データ プラットフォームである Proteinbase に公開しています。
AI が設計したタンパク質を現実世界でテストする方法
Adaptyv では、タンパク質設計を実験データに変換するための、自動化された AI ネイティブのウェット ラボを構築しました。
白衣を着た科学者が小さなチューブに入ったサンプルを 1 つずつピペッティングする代わりに、私たちは同じ実験を何倍も速く、より安価に実行する自動化されたワークセルを構築しました。
ヒトタンパク質の設計者とエージェントは、当社の Web プラットフォームと API を介してラボにアクセスし、さまざまなアッセイでのウェットラボ検証用のデジタルタンパク質配列を送信できます。
当社のプラットフォームは、タンパク質配列 (アミノ酸の文字列) を自動的に処理し、特定のタンパク質を作成する方法に関する生物学的指示をコード化する特別な DNA 配列に変換します。このデジタル DNA 配列は、研究室で DNA 構成要素を 1 つずつ組み立てることによって、実際の物理的な DNA 分子に変換されます。次に、無細胞タンパク質合成では、細胞が DNA を読み取ってタンパク質 (リボソーム、酵素、アミノ酸、エネルギー供給源) を構築するために使用する機構を利用し、それを実行します。

周囲にセルがありません。これは、信じられないほど少量の液体を高速かつ高スループットでピペッティングできる自動ロボットを使用して行われます。
この時点で、私たちはデジタル AI が設計した配列から研究室で物理的なタンパク質を作成しました。しかし、そのタンパク質が実際に設計通りにうまく機能するかどうかはまだテストされていません。タンパク質は生命全体の分子機構であり、私たちが食べた食べ物を消化し、有害な分子を分解して生存に必要な分子を作り、細胞内でエネルギーを生成し、DNAを切断して編集するなど、さまざまな働きをします。
ここでは、結合剤、つまり 1 つの標的に付着し、他には付着しないタンパク質をテストしています。ほとんどの抗体抗がん剤は結合剤です。妊娠検査薬の試薬やほとんどの診断薬の捕捉分子も同様です。タンパク質結合剤のターゲットへの「付着強度」は、特殊な機器を使用して測定でき、いわゆる K_D 値が得られます。 K_D が低いほど、より優れた結合剤、より高い親和性結合剤を意味し、新しい治療法を開発する際にしばしば望ましい特性となります。 Adaptyv では、これらの特殊な機器を独自のソフトウェアで実行し、すべての生データを処理して、治療、診断、研究用途に関連する幅広い標的タンパク質に対するクリーンな K_D 値を取得します。
この単純な入出力インターフェイスの背後には、複雑な実験プロセスがあり、多くの試薬、機器、プロトコル、測定を調整し、検証する必要があります。 Adaptyv では、その複雑さを自動化された品質管理されたワークフローにパッケージ化し、生物学的な質問に再現性をもって答え、定義された種類のデータを大規模に生成します。
Anthropic のチームは、これまでの公開コンテスト、ハッカソン、BenchBB から 16 の既存のタンパク質ターゲットをベンチに選びました。

de novo バインダー設計用のさまざまなクロード バージョンはすべて Claude Science を使用しています。彼らは全員、同様のプロンプトを受け取り、同じ公開ツールにアクセスできました。
次に、これらすべてのターゲットに対する設計が匿名化された方法で当社のウェット ラボに送信されました。どのクロードモデルがどのタンパク質を設計したかについては情報がありませんでした。当社では、アフィニティー特性評価アッセイを実行し、5 つの標的濃度を使用し、二重測定で SPR 上の結合強度を測定し、結果が堅牢で正確で、当社の品質管理基準に一致していることを確認しました。
デザインの 95% が表現されました。これは 3 年前なら印象的な見出しになったはずで、近年のタンパク質デザイン用 AI ツールの急速な進歩を示しています。この数値は、何百人もの専門タンパク質設計者を擁する当社の EGFR 競合他社の最高の発現率と一致し、RBX1 などの他の課題を上回りました。これらのうち、すべてのデザインのうち 354 (1,320) がターゲットを制限し、全体のヒット率は 26.8% であり、ターゲットごとのヒット率は非常に大きく異なります。 1 つのターゲット (MBP) では結合剤が生成されませんでした。全体として、クロードは、15 の標的のうち 14 に結合するタンパク質を設計しましたが、1 つの標的 (GDF-8 成熟) は、おそらく標的が凝集し、それ自体の他の分子に結合しているため、セットアップでの測定品質が低かったために除外されました。
競合他社と比較すると、特に上記のプロットのように Anthropic が提出した各ターゲットの単一ランを見た場合、Claude はすべてのヒット率を上回っています。公平な比較のために、de novo ミニバインダーのみを含むように各コンテストの結果をサブセット化しました。 Claude は TREM2 で 80% のヒット率を達成し、競合他社で報告した 38.3% を大きく上回りました。15-PGDH などのよりトリッキーなターゲットでも 3-3 以上の成功率を示しています。

Proteinbase で観察されたものよりも 2 倍高くなります。
最高のクロード バインダーは、コンテストの優勝者 6 社中 5 社と比較して、より優れた結合親和性も示します。これは、最良の 15-PGDH バインダーを大幅に上回り、1.7 uM から 33.4 nM となり、RBX1 についても同様に (競合製品の 25.7 nM から 3.9 nM) でした。驚いたことに、ニパウイルスコンテストで最高のバインダーに勝つことができませんでした。 GDF-8、RBX1、および Nipah については、Claude は可能なエピトープを広範囲に探索したようですが、TREM2 の場合は単一の明確に定義されたエピトープに収束しています。全体として、クロードは、タンパク質設計コンペティションに参加していたら、多作なタンパク質設計者になっていたでしょう。
これらすべての印象的な結果を見ると、オープンに利用可能な設計ツールの調整に関しては、クロードは専門のタンパク質設計者に匹敵するか、あるいはそれを超えているようです。私たちは、これがどこに向かうのか、そして薬剤科学とタンパク質設計の未来がどのようなものになるのかを見ることに間違いなく興奮しています。私たちの考えの一部を以下で共有します。
薬剤科学と実際の病気の治療
では、AI はすべての病気を解決できるのでしょうか?いや、まだです。
このケーススタディは、クロードがタンパク質設計ツールの調整において少なくとも専門家レベルであることを示しています。タンパク質設計ツールは使いにくいので、これは素晴らしいニュースです。 AI が登場する前は、タンパク質構造予測ツールをセットアップするだけでも、不透明な conda エラーをデバッグするのに何時間もかかることがありました。
ここで、新しい癌治療法に取り組んでいる研究者が、2 つの変異した受容体変異体を区別できるアッセイを必要としていると想像してください。彼らは細胞サンプルの配列決定結果をクロードに渡し、バインダーの設計に協力してもらうことができます。クロードは論文や研究データベースを精査し、ゲノム内のどこに受容体がコードされているかを見つけ、2 つの配列を比較し、違いをマッピングします。私

ボルツのようなタンパク質折り畳みモデルを使用して受容体変異体の 3D 構造を生成し、次に BindCraft のようなタンパク質設計モデルを使用して、一方の変異体には結合するがもう一方の変異体には結合しないバインダーを生成します。何千もの計算設計を生成し、それらをスコアリングして、テストする最も有望な候補を特定します。その後、クラウド ラボ API のおかげで、クロードは最良の候補者をラボに提出し、数週間で実験的にテストしてもらうことができました。結果が得られたら、クロードはデータを分析して、最も強力なバインダーが実際にはあまりにも乱雑であり、両方のバリアントを結合することを発見することができます。最初のラウンドでの情報を受けて、次のラウンドで交差反応性を低減し、より優れたタンパク質を生産するための別の設計キャンペーンを開始する可能性があります。これらは蛍光マーカーに結合し、がん細胞に結合してマークを付けることで、がん細胞と健康な細胞を区別することができます。
ほんの数年前には、このプロセスだけで研究室全体の 1 年分の作業が必要でした。現在、科学に興味のある人なら誰でも、数週間以内にラップトップからこのキャンペーンをエンドツーエンドで実行でき、AI トークン、クラウド コンピューティング、ウェット ラボ クレジットを数千ドルで入手できます。
すべての病気を治すための道を最初に構築する必要がある
もちろん、ここでテストしたタンパク質は本当の治療薬ではありません。彼らは、プロセスの最初のステップ、つまりバインダーとして機能できることを実証しただけを完了しました。それでも、この研究は AI を使った実際の治療薬の開発への道を示しています。
薬を作ることは山を登ることに似ていると想像してください。人類はすでに多くの薬を作っているので、明らかにいくつかの山に登ることができました。しかし、頂上への道は危険な狭い道であり、登頂には長い年月と数十億ドル（そして多くのバイオ技術者の命）がかかります。
AI for dの目標

ラグの発見により、この狭い山道が高速道路に変わり、より簡単かつ安価に頂上に到達できるようになり、現在よりも 100 倍多くの治療薬を開発できるようになります。同様に、LLM が導入される前は、コードの作成は高度な専門知識を必要とする技術でしたが、現在ではほとんどが自動化されており、誰でもソフトウェアを生成できるようになりました。
実際、創薬の場合、これは次のことを意味します。
分子設計とバイオインフォマティクスの作業を自動化して、より優れた薬剤候補を作成する
関連する生物学的疑問に答える優れた実験結果を構築する
実験ワークフローを自動化してスループットを向上させ、大規模なコストを削減する
ウェットラボ側では、道路はおおよそ次のようになります。
表現: 新たに設計したタンパク質を確実に合成できるでしょうか?
バインダーの設計: 意図した標的に結合するタンパク質を作ることはできるでしょうか?
治療用フォーマットと開発可能性: 候補は業界で使用されているフォーマットで作成でき、安定性を維持し、確実に製造できるか?
細胞ベースの機能: 結合は生細胞システムで意図した効果を生み出しますか?
オルガノイドとより現実的なモデル: 人間の生物学をよりよく捉えたモデルで機能するでしょうか?
トランスレーショナルおよび生体内証拠: 最終的な環境において安全で効果的か?

[切り捨てられた]

## Original Extract

The Anthropic team wanted to benchmark the newest Claude models at protein design. They ran a suite of models in Claude Science to design protein binders against 16 targets and then sent them to our lab for experimental validation. The results show that Claude models show expert level skills at prot
[truncated]

Case study: Benchmarking Claude’s protein designs in the wet lab | Adaptyv Bio Services Binding Expression Thermostability Blog Team Careers API Proteinbase Docs Book a meeting Test your proteins Services Blog Team Careers API Proteinbase Docs Book a meeting Test your proteins Sign in to Foundry Adaptyv Bio Biopole Life Science Campus, Route de la Corniche 5, 1066 Epalinges, Lausanne, Switzerland
← Back to Blog Case study: Benchmarking Claude’s protein designs in the wet lab
- The Anthropic team used Claude in Claude Science to design protein binders against 16 targets and sent them to our lab.
- We ran all the binding experiments in our automated wet lab, going from digital protein sequence, to DNA, expressed proteins, then measurement and quality control.
- Out of 1,320 designs, 354 bound their target, giving Claude an average hit rate of 26.8%. Only a single target yielded 0 binders.
- Compared to the results of our Protein Design Competitions, Claude’s designs had noticeably higher success rates and would have won 5 out 6, yielding tighter binders.
- At the end of the post we expand more on how we think about the future of agentic science, protein design and the long all road to curing all diseases.
The promise of bioengineering is to design new biological tools: diagnostics that help us detect diseases, antibodies that can destroy cancers, new vaccines to protect against viruses. To build those tools, AI will ultimately need to interact with the physical world and not just with a static dataset .
The next phase of AI in biology is an agentic science loop: an AI system that can reason over scientific knowledge, use specialised tools to design thousands of molecules, send those designs into an automated wet-lab, and learn from the resulting data.
Three layers need to work together:
Generalist models to interpret a scientific goal, plan a campaign, and orchestrate workflows.
Specialist models and tools to perform tasks such as protein structure prediction and protein design models.
Automated wet labs to turn digital designs into measurements from the physical world.
In this post, we showcase how Anthropic benchmarked Claude Mythos Preview and Opus 4.8 at designing novel protein binders that we tested in our automated wet lab. Overall, Claude models showed expert level skill at protein design, matching or exceeding human experts on many tasks. Together with the Anthropic team, we’re releasing the actual protein sequences Claude designed as well as the experimental data on Proteinbase , the open protein data platform .
How to test AI-designed proteins in the real world
At Adaptyv, we have built an automated, AI-native wet lab for turning protein designs into experimental data.
Instead of scientists in lab coats pipetting samples in tiny tubes one by one, we have built automated workcells that run the same experiments many times faster and cheaper.
Human protein designers and agents can access the lab via our web platform and API to send digital protein sequences for wet lab validation on different assays.
Our platform automatically processes the protein sequences (a string of amino acids) and converts it into a special DNA sequence that encodes the biological instructions for how to create the specific protein. This digital DNA sequence is then turned into an actual physical DNA molecule in the lab by assembling the DNA building blocks one by one . Next, cell-free protein synthesis takes the machinery a cell uses to read DNA and build proteins (ribosomes, enzymes, amino acids, an energy supply) and runs it with no cell around it . This is done using automated robots able to pipette incredibly small amounts of liquid really fast and at high-throughput.
At this point we have now made physical proteins in the lab from the digital AI-designed sequence. But we haven’t yet tested if the protein actually performs well at what it was designed to do. Proteins are the molecular machinery of all of life and can do many different things: digesting the food we eat, breaking down harmful molecules and making ones that we need to survive, generating energy in our cells, cutting and editing DNA, and a million other things.
Here, we’re testing binders: proteins that should stick to one target and nothing else. Most antibody cancer drugs are binders. So are the reagents in a pregnancy test and the capture molecules in most diagnostics. The “sticking strength” of the protein binder to its target can be measured with special instruments, giving us a so called K_D value. A lower K_D means a better binder, higher affinity binder, an often desirable property when developing new therapeutics. At Adaptyv, we run these specialized instruments with our own software to process all raw data to obtain clean K_D values against a wide range of target proteins that are relevant for therapeutic, diagnostic, and research applications .
Behind this simple input-output interface is a complex experimental process: many reagents, instruments, protocols and measurements must be coordinated and verified. At Adaptyv we package that complexity into an automated, quality-controlled workflow that answers a biological question reproducibly and generates a defined type of data at scale.
The team at Anthropic chose 16 existing protein targets from our previous public competitions , hackathons, and BenchBB to benchmark different Claude versions for de novo binder design, all using Claude Science . They were all prompted similarly and had access to the same publicly available tools.
Then, designs against all these targets were sent to our wet lab, in an anonymized way. We did not have any information on which Claude model designed which protein. We ran our Affinity Characterization assay, to measure the binding strength on SPR using 5 target concentrations and with duplicate measurement, ensuring the results are robust, accurate, and matching our quality control standards.
95% of the designs expressed, which three years ago would have been an impressive headline, showing the rapid progress of AI tools for protein design in recent years. This number matched the best expression rates of our EGFR competition which had hundreds of expert protein designers, and surpassed other challenges such as the RBX1 one. Out of these, 354 of all designs (1,320) bound their target, an overall hit rate of 26.8%, and the per-target hit-rates vary quite widely. One target (MBP) yielded no binders. Overall, Claude designed proteins that bound 14 out of 15 targets, with 1 target (GDF-8 mature) being excluded due to low-quality measurements in our setup, likely due to the target aggregating, meaning it was binding to other molecules of itself.
When compared to our competitions, Claude surpasses all their hit rates, especially when looking at every single run for each target Anthropic submitted as in the plot above. For a fair comparison, we have subsetted each competition’s results to only include de novo minibinders. Claude achieved an 80% hit rate on TREM2, greatly improving over the 38.3% we reported in our competition , and even on trickier targets such as 15-PGDH, it has a success rate more than 3-fold higher than observed on Proteinbase .
The best Claude binders also show greater binding affinities compared to 5 out of 6 competition winners. It greatly surpassed the best 15-PGDH binder, going from 1.7 uM to 33.4 nM, and similarly for RBX1 (from 25.7 nM in the competition to 3.9 nM). Surprisingly, it was unable to beat the best binder on the Nipah Virus Competition. For GDF-8, RBX1, and Nipah, Claude seems to have explored a wide range of possible epitopes, while it converges on single well-defined epitopes in case of TREM2. Overall, Claude would have been a prolific protein designer if it were to participate in our Protein Design Competitions.
Looking at all these impressive results, Claude seems to have matched or even surpassed expert protein designers when it comes to orchestrating openly available design tools. We are definitely excited to see where this is heading towards and what the future of agentic science and protein design might look like. We’re sharing some of our thoughts down below.
Agentic science and curing actual diseases
So can AI now solve all diseases? Well no, not yet.
This case study shows that Claude is at least expert-level at orchestrating protein design tools. That’s great news, since protein design tools are hard to use. Before AI, even setting up a protein structure prediction tool could take hours of debugging opaque conda errors.
Now, imagine a researcher working on a new cancer therapeutic who needs an assay that can distinguish between two mutated receptor variants. They can give Claude the sequencing results from their cell samples and ask it to help design binders. Claude would sift through papers and research databases, find where in the genome the receptors are encoded, compare the two sequences, and map the differences. It would use protein folding models like Boltz to generate 3D structures of the receptor variants, then use protein design models like BindCraft to generate binders that attach to one variant but not the other. It would generate thousands of computational designs and score them to identify the most promising candidates to test. Then, thanks to our cloud lab API , Claude could submit the best candidates to the lab and have them tested experimentally in a couple of weeks. Once the results are in, Claude could analyze the data and find that the strongest binders are actually too promiscuous and bind both variants. Informed by the first round, it could then launch another design campaign to reduce that cross-reactivity and produce better proteins in the next round. Those can be linked to a fluorescent marker and then bind and mark the cancer cells to differentiate them from healthy cells.
Just a few years ago, this process alone would’ve required a year’s worth of a whole lab’s work. Today, anyone who is scientifically curious can run this campaign end-to-end from their laptop in a few weeks and for a couple thousand $ in AI tokens, cloud compute and wet lab credits.
The road to curing all disease needs to be built first
Of course, those proteins that we tested here are not real therapeutics. They completed only the first step of the process: demonstrating that they can function as binders. Still, this study shows a path towards making actual therapeutics with AI.
Imagine making a drug is like climbing a mountain. We have clearly been able to climb some mountains, as humanity has made many drugs already. But the way to the top is a dangerous narrow path and climbing it takes many years and costs billions of dollars (and the lives of many biotechs).
The goal of AI for drug discovery is turning this narrow mountain path into a highway, making it easier and cheaper to get to the top so that we can develop 100x more therapeutics than we have right now. Similarly, w riting code was a more of a high-expertise craft before LLMs, now it’s mostly automated and it has made generating software accessible for anyone.
In practice, for drug discovery, that means:
Automating molecular design and bioinformatics work to make better drug candidates
Building good experimental readouts that answer relevant biological questions
Automating those experimental workflows to increase throughput and lower costs at scale
On the wet-lab side, the road then looks roughly like this:
Expression: Can we reliably synthesise newly designed proteins?
Binder design: Can we make proteins that bind their intended target?
Therapeutic formats and developability: Can a candidate be made in the formats industry uses, remain stable and be manufactured reliably?
Cell-based function: Does binding produce the intended effect in a living-cell system?
Organoids and more realistic models: Does it work in a model that better captures human biology?
Translational and in-vivo evidence: Is it safe and effective in the settings that ultimatel

[truncated]
