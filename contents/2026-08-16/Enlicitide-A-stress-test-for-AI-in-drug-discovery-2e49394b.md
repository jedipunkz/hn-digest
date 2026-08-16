---
source: "https://www.empirical.health/blog/macrocyclic-peptides/"
hn_url: "https://news.ycombinator.com/item?id=49320793"
title: "Enlicitide: A stress test for AI in drug discovery"
article_title: "Enlicitide: a stress test for AI in drug discovery. | Empirical Health"
author: "brandonb"
captured_at: "2026-08-16T15:11:54Z"
capture_tool: "hn-digest"
hn_id: 49320793
score: 1
comments: 0
posted_at: "2026-08-16T15:07:05Z"
tags:
  - hacker-news
  - translated
---

# Enlicitide: A stress test for AI in drug discovery

- HN: [49320793](https://news.ycombinator.com/item?id=49320793)
- Source: [www.empirical.health](https://www.empirical.health/blog/macrocyclic-peptides/)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T15:07:05Z

## Translation

タイトル: Enlicitide: 創薬における AI のストレス テスト
記事のタイトル: Enlicitide: 創薬における AI のストレス テスト。 |経験的健康
説明: メルク社は、PCSK9 をブロックする可能性のある低分子について十数件の検査を実施しましたが、何も見つかりませんでした。最終的に効果を発揮した薬はリング状のペプチドで、13年の歳月を要した。

記事本文:
Enlicitide: 創薬における AI のストレス テスト。 |経験的健康
190 ドルのプログラムで ApoB、Lp(a)、および 100 以上のその他のバイオマーカーをテスト
App Store Google Play Web App はじめる はじめる Enlicitide: 創薬における AI のストレス テスト。
ブランドン・バリンジャー · 2026年7月26日
従来、医薬品は低分子か生物製剤のいずれかでした。アスピリンやアトルバスタチンのような小分子は錠剤として摂取できます。生物製剤は汎用性が高い（タンパク質表面全体をカバーできるほど大きい）ものの、注射する必要があります。
中間カテゴリーがあったらどうなるでしょうか？今はあります。大環状ペプチドは、環構造に折り畳まれた大きなペプチド (GLP-1 など) です。初の経口PCSK9阻害剤であるエンリシチド（商品名リプフェンドラ）が先週、FDAによって承認された。これは、LDL コレステロールを約 60% 低下させる PCSK9 阻害剤の丸剤であり、丸剤の安定性で生物学的製剤を正確に標的にします。
Enlicitide (Lipfendra) は AI なしで発見されました (プログラムは 13 年前に開始されました)。この発見には新しい化学が関与しているため、創薬における AI が実際にどこに役立つか (または役に立たないか) を理解するための興味深い現実世界のストレス テストとなります。この記事では、そもそも PCSK9 がなぜ難しい標的なのか、in vitro スクリーニングによってどのようにして 1 兆のペプチドから最初のヒットに絞り込まれたのか、そして最初のヒットを使用可能な薬剤に変える際の難しい部分について説明します。各セクションでは、AI がどの程度役立つかについて説明します。
そもそもなぜPCSK9が難しいターゲットなのか
エンシチドは、肝臓の LDL 受容体を破壊するタンパク質である PCSK9 をブロックすることによって機能します。これらの受容体は血液からコレステロールを引き出すため、PCSK9をブロックするとLDLが低下します。 PCSK9 を錠剤の標的とするのは特に困難な 3 つの副次的問題があります。

対生物学的製剤:
明らかに目標はフラットです。 PCSK9 が LDL 受容体をグリップする表面は、自然のポケットから 20 オングストローム以上離れたところにあります。
PCSK9 の酵素活性は非常に優れています。 PCSK9 はプロテアーゼであるため、その活性部位を妨害するのは明らかです。しかし、McNuttらは2007年に触媒作用で死んだバージョンを構築し、それが依然としてLDL受容体を破壊することを発見した。
その自然の深くて形の良い空洞は、すでに独自のプロドメインによって塞がれています。
メルクは2019年、この標的に対する従来の錠剤を見つけることができなかったことを説明する論文を発表した。対照的に、PCSK9 を標的にすることは「抗体は巨大であるため、抗体にとっては最適です」（エンリシチドを開発したチームの臨床ディレクターであるダグラス・ジョンズ氏の説明によると）。
以下は、enlicitide の最終的に設計された構造です。
エンリチサイド図。 8 つのアミノ酸のうち 6 つは自然界には存在しません。リングは構造を安定させるために使用されます。出典: 米国ペプチド協会、リップフェンドラ処方情報。
これは 8 つのアミノ酸残基 (そのうち 6 つは人工) と、分子の安定化に役立つ (結合親和性を助ける) 2 つの閉環で構成されています。どのようにしてこのエンドポイントに到達したのでしょうか?この構造がどのように導き出されたのかを見ていきます。物語は実際には4つの段階で展開します。
Enlicitide は 4 つの主要な段階で 13 年かかりました。 AI があればどこが役立つでしょうか?
構想から FDA の承認までには 13 年の開発期間が必要でした。これらは 4 つの段階に分けられます: 1. 平面に結合する最初の分子を見つける、2. 分子が薬として使用できるように結合力を高める、3. 腸壁を通過させる、4. 大規模に製造する。
enlicitide の開発の 4 つの段階と、AI がどの程度役立つかについての大まかな評価。
ステージ 1: mRNA ディスプレイがどのように開始分子を見つけたか
2013 年にメルク社

新興企業である Ra Pharmaceuticals から mRNA ディスプレイ技術のライセンスを取得しました。 Ra は、ノーベル賞受賞者ジャック・ショスタクの業績を中心に設立されたケンブリッジの企業です。条件は前払いで450万ドル、マイルストーンで最大5,600万ドルというものだった。ピーク時売上高が数十億ドルに達すると予想される医薬品としては、これはこれまでに購入した中で最も良い選択肢の一つだ。 UCBは最終的に2020年にRaを21億ドルで買収した。
mRNA ディスプレイは in vitro 技術です。それは本質的に、細胞ではなく試験管内で指向性進化を実行します。指向性進化には 2 つのことが必要です。変異体の母集団と、勝者の指示を読み取る方法です。自然はこのために細胞とそのすべての機構を使用します。 mRNAディスプレイは、各ペプチドをそれをコードするmRNAに結び付ける「ひも」（ピューロマイシン）を備えた試験管内で行われます。これは、洗浄ステップを経て生き残った結合剤が逆転写、増幅、変異し、再度実行できることを意味します。
mRNA ディスプレイでは、mRNA リードを使用して試験管内で指向性進化を実行し、各ステップで生存者を増幅します。
AI ならこれを高速化できるでしょうか?はい、後のどの段階よりもそうです。 De novo マクロサイクル設計は現在、真に機能しています。 David Baker の研究室は 2025 年 6 月に RF ペプチドを発表しました。彼らは 77 のデザインを合成し、39 をテストし、12 の結合剤を得ました。その中には、著者らが「かなり平坦で標的とするのが難しい」細菌タンパク質に対する 9.4 nM の結合剤も含まれます。 Latent Labs は、標準ベンチマークに対して 91% ～ 100% のヒット率を報告しています。これで、ステージ 1 をラップトップで実行できるようになりました。
RFpeptides は、ランダムな原子から開始して、ターゲットに適合するリング (t = 50/50 ～ t = 1) にそれらの原子をノイズ除去する AI モデルです。出典: Nature Chemical Biology 2025、CC BY 4.0。
実際的な制限の 1 つは、RF ペプチドが 20 個の天然アミノ酸のみを使用することです。 Enlicitide には 6 つの非天然残基が必要でした。 mRNAディスプレイはそれを処理します

ああ。明らかに、Unnatural Products のような AI ファーストのショップは、依然としてディスプレイによるスクリーニングから開始し、その後、結果として得られるヒット商品を最適化するために計算を適用しています。これまでのところ、AI によって設計された大環状ペプチドは臨床試験に入っていません。
ステージ 2: 190,000 倍の強力な結合
mRNA スクリーニングから得られた最初のヒットは良いスタートでしたが、それ自体では効果的な薬にはなりませんでした。 PCSK9 との結合は弱く、全血中での半減期はわずか 1 時間でした。最適化後、承認された薬剤は 190,000 倍強く結合し、半減期が 8 時間に近づきました。
最初のヒットと最適化された薬剤との間のギャップ。出典: ACS Med Chem Lett 2022 および Circulation 2023 。
その効力の変化は、PCSK9 との新しい接触点を見つけたことによってもたらされたものではありません。化学者らは 2 番目と 3 番目の環を追加し、分子のぐらつきの自由を奪い、効力を約 500 倍向上させました。 （フロッピーチェーンは、結合するために 1 つの形状に「凍結」する必要があり、結合エネルギーからその代償を支払いますが、リングはより硬いです。）ウプサラ大学の医薬品化学者、ヤン キールバーグ氏は、これを簡単に言います。「フロッピーな紐の代わりに、ドーナツにします。」
残りの作業は、苦労して勝ち取った小さな勝利の積み重ねでした。たとえば、2 つのプロテアーゼ切断により剛性が向上しました (これにより効力が向上します)。 5-フルオロトリプトファン (非天然アミノ酸残基) を使用してそのフッ素を PCSK9 の浅い凹みに貼り付けると、効力がさらに高まりました。さらに一連の変更により酸化が 98% から 2.9% に減少し、薬の血中での持続時間が長くなりました。
最初のヒトデータは、米国心臓協会の会議で発表されました。 1 回の投与で、血液中の遊離 PCSK9 の 93% 以上が除去されました。当時、研究者の一人は次のように述べています。

ee PCSK9 を基本的にゼロのレベルに引き上げます。これが抗体の働きです。そして、私たちは何かを持っていることに気づきました。」
AI ならこれを高速化できるでしょうか?期待よりも少ないです。タンパク質間標的に対する Latent-X の親和性は 5 ～ 72 マイクロモル、RF ペプチドは 6 ナノモル、エンシチドは 5 ピコモルです。つまり、私たちが到達すべき目標にはまだ 3 桁も達していません。
ステージ 3: エンシチドを吸収させる
次の課題は吸収です。すべてのリプフェンドラ錠剤には、腸細胞の密着結合を緩めるカプリン酸ナトリウムが含まれています。したがって、薬剤は腸細胞を「通過」せず、腸細胞の間を圧迫します。吸収のためのもう 1 つの手段は、まったくの効力です。5 ピコモルの親和性では、20 mg 用量の 1% を吸収するだけで十分です。
これが、リプフェンドラを空腹時に水、ブラックコーヒー、または普通のお茶と一緒に服用する必要があり、食べる前に少なくとも30分待つ必要がある理由の1つです。食物がエンハンサーの邪魔をします。 （ただし、このルールは永続的ではない可能性があります。David Baker のグループは、エンハンサーをまったく使用しないマウスで経口バイオアベイラビリティが 40% に達する大環状化合物を設計しました。）
AI ならこれを高速化できるでしょうか?そうではありません。アストラゼネカの分子 AI グループは 2026 年に、生成モデルがトレーニング データの範囲外にある場合、透過性予測は「ランダムな出力に低下する」と報告しました。そして、ここで機能した細胞傍吸収の「トリック」を扱うモデルはありません。
ステージ 4: 製造の規模拡大
8 個のアミノ酸 (そのうち 6 個は非天然) からなる環を大量に作るのは簡単ではありません。最初の合成ルートは 63 のステップを要し、クロマトグラフィーに頼っていました。単一の 100 kg バッチを生産するには、170 トンを超える中間体が必要になります (これにより、node_module さえも効率的に見えます)。 2026 年 5 月、メルクのプロセス化学者らは次の論文を発表しました。

13 の人工酵素を使用したサイエンスのルートで、ステップ数が半分に減ります。
AI ならこれを高速化できるでしょうか?はい、部分的にはそうです。メルクのルートは、実行するために進化したわけではない反応を触媒するように設計された 13 の酵素に依存しており、ML に基づくタンパク質工学はその仕事に本当に適しています。
AI が実際に創薬に役立つ場所
大環状ペプチドは、錠剤の形状で生物学的製剤の機能を備えた、本当に新しくて興味深いカテゴリーの医薬品のように思えます。
AI が役に立ったでしょうか?プロセスの両端で最も強くなり、プロセスの途中で最も弱くなります。ステージ1は圧勝です。試験管の中に 1 兆人ものメンバーが入ったライブラリを必要としていたものが、今ではラップトップ上で数十のデザインを必要としています。ステージ 4 は部分的な勝利です。新しい製造ルートは 13 の人工酵素に依存しており、ML に基づくタンパク質工学がその作業では日常的なものであるためです。
ステージ 2 と 3 がボトルネックです。ここは 13 年間の大部分が費やされた場所でもあります。ナノモルの結合剤をピコモルの薬剤に変えることは、公開されているデノボ手法が到達するレベルをまだ約 3 桁超えており、透過性の予測はトレーニング データの外では破綻します。どちらの差も小さいわけではなく、明らかに縮まっているわけではありません。
AIは分子を設計できる。それを麻薬に変えるのはまだ遅いです。
こちらも参照してください: 初の経口 PCSK9 阻害剤である Lipfendra、 Lipfendra 対 Repatha、および PCSK9 阻害剤 vs スタチン。
30 日間の心臓健康ガイドを無料で入手
心臓の健康を最適化するための証拠に基づいた手順。
心臓病によって死亡する人は、すべてのがんを合わせたよりも多くなります。
それをあなたにさせないでください。
今すぐ 2,200 の検査会場の 1 つに立ち寄って、より良い心臓への旅を始めましょう
健康。
メディケアの対象となる心臓血管ケア
ニューヨーク州ニューヨークの❤️で作られています · © 2026 Empirical Health

## Original Extract

Merck ran more than a dozen screens for a small molecule that could block PCSK9 and found nothing. The drug that finally worked is a ring-shaped peptide, and it took 13 years.

Enlicitide: a stress test for AI in drug discovery. | Empirical Health
Test ApoB, Lp(a), and 100+ other biomarkers for $190 Programs
App Store Google Play Web App Get started Get started Enlicitide: a stress test for AI in drug discovery.
Brandon Ballinger · Jul 26, 2026
Traditionally, medicines were either small molecules or biologics. Small molecules like aspirin or atorvastatin can be taken as pills. Biologics are highly versatile (they’re large enough to cover an entire protein surface) but must be injected.
What if there were a middle category? Now there is. Macrocyclic peptides are large peptides (like GLP-1s) folded into a ring structure. The first oral PCSK9 inhibitor, enlicitide (brand name Lipfendra), was approved by the FDA last week. It’s a pill form of PCSK9 inhibitor that lowers LDL cholesterol by about 60%, giving the precise targeting of a biologic with the stability of a pill.
Enlicitide (Lipfendra) was discovered without AI (the program began 13 years ago). The discovery involves some novel chemistry, and so it’s an interesting real-world stress test to understand where AI in drug discovery would actually help (or not). In this article, I’ll cover what makes PCSK9 a difficult target in the first place, how in vitro screening narrowed down one trillion peptides to the initial hit, and the difficult parts of turning that initial hit into a usable drug. Within each section, we’ll discuss how much AI would have helped.
Why PCSK9 is a difficult target in the first place
Enlicitide works by blocking PCSK9, a protein that destroys the liver’s LDL receptors. Those receptors are what pull cholesterol out of your blood, so blocking PCSK9 lowers LDL. There are three sub-problems that make PCSK9 an especially difficult target for a pill vs a biologic:
The obvious target is flat . The surface where PCSK9 grips the LDL receptor sits more than 20 angstroms from the natural pocket.
PCSK9’s enzyme activity is a red herring. PCSK9 is a protease, so the obvious move is to jam its active site. But McNutt and colleagues built a catalytically dead version in 2007, and found it still destroyed LDL receptors.
Its natural deep, well-shaped cavity is already plugged by its own prodomain.
In 2019, Merck published a paper describing its own failure to find a conventional pill for this target. In contrast, targeting PCSK9 is “ great for antibodies because they’re huge ” (as explained by Douglas Johns, a clinical director on the team that developed enlicitide).
Below is the final designed structure of enlicitide:
Enlicitide diagram. 6 of the 8 amino acids don’t occur in nature. Rings are used to stabilize the structure. Source: American Peptide Society , Lipfendra prescribing information .
It consists of 8 amino acid residues (6 of which are artificial), and two ring closures that help stabilize the molecule (helps with binding affinity). How did we reach this endpoint? We’ll go through the process of how this structure was derived. The story really unfolds in four stages.
Enlicitide took 13 years over four main stages. Where would AI have helped?
Going from conception to FDA approval required 13 years of development. Those were divided into four stages: 1. find an initial molecule that binds a flat surface, 2. increase binding potency so that molecule becomes usable as a drug, 3. get it across the gut wall, and 4. manufacture it at scale.
Four stages of enlicitide’s development, along with a rough rating of how much AI would have helped.
Stage 1: how mRNA display found the starting molecule
In 2013, Merck licensed an mRNA display technology from a startup, Ra Pharmaceuticals. Ra is a Cambridge company built around Nobel laureate Jack Szostak’s work. The terms were $4.5 million up front and up to $56 million in milestones, which for a drug now forecast at multibillion-dollar peak sales is one of the better options anyone has bought. UCB ultimately acquired Ra for $2.1 billion in 2020.
mRNA display is an in vitro technology. It essentially runs directed evolution in a test tube rather than a cell. Directed evolution needs two things: a population of variants, and a way to read the winners’ instructions back out. Nature uses cells and all their machinery for this. mRNA display does it in a test tube with a “leash” (puromycin) that links each peptide to the mRNA encoding it. This means the binders that survive a wash step can be reverse transcribed, amplified, mutated, and run again.
mRNA display runs directed evolution in a test tube, using an mRNA leash to amplify survivors at each step.
Would AI have sped this up? More than any later stage, yes. De novo macrocycle design genuinely works now. David Baker’s lab published RFpeptides in June 2025: they synthesized 77 designs, tested 39, and got 12 binders, including one at 9.4 nM against a bacterial protein the authors call “considerably flatter and difficult to target.” Latent Labs reports hit rates of 91% to 100% against standard benchmarks. So you can now run stage 1 on a laptop.
RFpeptides is an AI model that starts from random atoms and denoises them into a ring that fits the target (t = 50/50 to t = 1). Source: Nature Chemical Biology 2025 , CC BY 4.0.
One practical limitation is that RFpeptides only uses the 20 natural amino acids. Enlicitide needed six non-natural residues; mRNA display handles that today. Tellingly, the AI-first shops like Unnatural Products still start by screening with display, then apply computation to optimize the resulting hits. As yet, no AI-designed macrocyclic peptide has entered a clinical trial.
Stage 2: 190,000x tighter binding
The initial hit that came out of the mRNA screen was a good start, but it wouldn’t have been an effective drug on its own. It bound PCSK9 only weakly, and its half life in whole blood was only an hour. After optimization, the approved drug binds 190,000x more tightly and has a half-life closer to 8 hours:
The gap between the initial hit and the optimized drug. Source: ACS Med Chem Lett 2022 and Circulation 2023 .
That potency change didn’t come from finding new points of contact with PCSK9. The chemists added a second and third ring, which improved potency roughly 500-fold by taking away the molecule’s freedom to wobble. (A floppy chain has to “freeze” into one shape to bind, and pays for that out of its binding energy, whereas a ring is more rigid.) Jan Kihlberg, a medicinal chemist at Uppsala University, puts it simply : “Instead of being a floppy piece of string, you make it into a donut.”
The rest of the work was a slog of hard-won small victories. For example, two protease cuts improved rigidity (which improves potency). Using 5-fluorotryptophan (a non-natural amino acid residue) to stick its fluorine into a shallow dent on PCSK9 increased potency further. And then a series of changes reduced oxidation from 98% to 2.9%, making the drug last longer in blood.
The first human data was presented at an American Heart Association meeting. A single dose wiped out more than 93% of the free PCSK9 in people’s blood. At the time one of the researchers said : “That’s when I got goosebumps. We reduced free PCSK9 to levels that were basically to zero, which is what the antibodies do. And then we knew we had something.”
Would AI have sped this up? Less than you’d hope. Latent-X’s affinities against protein-protein targets are 5-72 micromolar, RFpeptides’ is 6 nanomolar, and enlicitide is 5 picomolar. So we’re still three orders of magnitude short of where we’d need to be.
Stage 3: making enlicitide get absorbed
The next challenge is absorption. Every Lipfendra tablet contains sodium caprate, which loosens the tight junctions in your intestinal cells. So the drug doesn’t “cross” through gut cells, it squeezes between them. The other lever for absorption is just sheer potency: at 5 picomolar affinity, absorbing 1% of a 20 mg dose is plenty.
This is partly why Lipfendra has to be taken on an empty stomach with water, black coffee, or plain tea, and requires you to wait at least 30 minutes before eating. Food interferes with the enhancer. (The rules may not be permanent, though: David Baker’s group has designed macrocycles that reach 40% oral bioavailability in mice with no enhancer at all.)
Would AI have sped this up? Not really. AstraZeneca’s Molecular AI group reported in 2026 that when generative models wander outside their training data, permeability predictions “reduce to random outputs.” And no model handles the paracellular absorption “trick” which worked here.
Stage 4: scaling manufacturing
A ring with eight amino acids, six of which are non-natural, isn’t easy to make in bulk. The first synthetic route took 63 steps and leaned on chromatography. Producing a single 100 kg batch would have required more than 170 metric tons of intermediates (this makes even node_modules look efficient). In May 2026, Merck’s process chemists published a route in Science using 13 engineered enzymes, which cut the number of steps in half.
Would AI have sped this up? Yes, in part. Merck’s route depends on 13 enzymes engineered to catalyze reactions they didn’t evolve to run, and ML-guided protein engineering is genuinely good for that job.
Where AI actually helps in drug discovery
Macrocyclic peptides do genuinely seem like a new and interesting category of medicine, with the capabilities of a biologic in the form factor of a pill.
Would AI have helped? It’s strongest at the two ends and weakest in the middle of the process. Stage 1 is the clear win. What took a trillion-member library in a test tube now takes a few dozen designs on a laptop. Stage 4 is a partial win, because the new manufacturing route depends on 13 engineered enzymes, and ML-guided protein engineering is routine for that job.
Stages 2 and 3 are the bottleneck. They are also where most of the 13 years went. Turning a nanomolar binder into a picomolar drug is still about three orders of magnitude past what published de novo methods reach, and permeability prediction falls apart outside its training data. Neither is a small gap, and neither is obviously closing.
AI can design the molecule. Turning it into a drug is still the slow part.
See also: Lipfendra, the first oral PCSK9 inhibitor , Lipfendra vs Repatha , and PCSK9 inhibitors vs statins .
Get your free 30-day heart health guide
Evidence-based steps to optimize your heart health.
Heart disease kills more people than all cancers combined.
Don't let it be you.
Stop by one of 2,200 testing sites today and start your journey to better heart
health.
Medicare-covered cardiovascular care
Made with ❤️ in New York, NY · © 2026 Empirical Health
