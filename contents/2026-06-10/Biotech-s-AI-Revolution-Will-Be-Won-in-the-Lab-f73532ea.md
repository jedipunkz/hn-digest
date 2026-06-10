---
source: "https://www.tamiola.com/blog/biotechs-ai-revolution-will-be-won-in-the-lab"
hn_url: "https://news.ycombinator.com/item?id=48475398"
title: "Biotech's AI Revolution Will Be Won in the Lab"
article_title: "Biotech’s AI Revolution Will Be Won in the Lab — Kamil Tamiola"
author: "ktamiola"
captured_at: "2026-06-10T13:18:21Z"
capture_tool: "hn-digest"
hn_id: 48475398
score: 1
comments: 0
posted_at: "2026-06-10T12:38:37Z"
tags:
  - hacker-news
  - translated
---

# Biotech's AI Revolution Will Be Won in the Lab

- HN: [48475398](https://news.ycombinator.com/item?id=48475398)
- Source: [www.tamiola.com](https://www.tamiola.com/blog/biotechs-ai-revolution-will-be-won-in-the-lab)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T12:38:37Z

## Translation

Title: Biotech's AI Revolution Will Be Won in the Lab
Article title: Biotech’s AI Revolution Will Be Won in the Lab — Kamil Tamiola
Description: AI drug discovery is only as strong as the experimental data behind it. At Peptone, we are investing in independent data generation.アウトソーシングされたデータセットと再現性の課題がモデルのパフォーマンスを左右する可能性がある分野では、モデルとラボのループを所有することが戦略的な利点になりつつあります。

記事本文:
0
Skip to Content
Kamil Tamiola
ホーム
研究
写真
ブログ
Open Menu
Close Menu
Kamil Tamiola
ホーム
研究
写真
ブログ
Open Menu
Close Menu
ホーム
研究
写真
ブログ
バイオテクノロジーの AI 革命は研究室で勝ち取られる
The next generation of AI for drug discovery will not be the biggest model trained on the most data. It will be the most specialized model trained on the right data — data we understand because we made it.
それは単純な信頼行為から始まります
Imagine you are in a lab, and the experiment you are about to run depends on a simple act of trust. You open a supplier's catalogue, choose an antibody, and look at a validation image called a Western blot.非専門家にとって、ウェスタンブロットは基本的に、抗体が適切なタンパク質を認識するかどうかを示す実験室テストです。適切な暗いバンドが適切な場所に現れた場合、科学者はそれを抗体が機能している証拠とみなします。
ここで、誰かがズームインしてコントラストを調整し、それらのバンドの一部がコピー、反転、ペイント、または別の製品ページ間で再利用されている可能性があることを発見したと想像してください。 Suddenly the problem is no longer one questionable image. It is a much larger question: how much of the scientific supply chain are we trusting without truly verifying?
The bands labeled 1 through 4 are all identical to one another after a vertical flip, a horizontal flip or a 180 degree rotation. Source: How much of Thermo Fisher’s antibody data has been manipulated?リース・リチャードソン著、2026年5月28日。
最近の Nature 記事は、100 を超える Thermo Fisher Scientific 抗体のカタログ エントリに、抗体の品質と性能を実証することを目的とした画像を含む、操作されたと思われる画像が含まれていると報告しました (Nature)。画像の変更は自動化されない

つまり、根底にある製品に欠陥があることを意味しますが、このエピソードは、信頼できる商用データ ストリームであっても隠れた不確実性を伴う可能性があることをはっきりと思い出させます。
What makes the Thermo Fisher case troubling is that the reported problems were not limited to a single figure or one isolated product page.リース・リチャードソン氏の分析では、オンラインの一次抗体カタログで「操作の兆候のある100枚以上の画像」が特定された。反転や回転後に同一に見えるウェスタンブロットのバンド、コントラスト調整後の顕著なブラシストロークのような編集、背景ノイズの反復ブロック、背景テクスチャの突然の不連続性などだ。ある繰り返しのパターンでは、同じ背景を共有する検証のしみが数十ページに示されており、この記事の執筆時点で 50 件の事例が文書化されています。
In academia, a pattern of apparent image manipulation on this scale would likely be career-ending.
バンドの重複、ブロットの特徴の回転、多くの人物の背景の繰り返しの疑惑に直面した主任研究者は、機関による調査、撤回、助成金の精査、そして場合によっては研究室の信頼性の崩壊を予想するでしょう。したがって、産業界にとって不快な質問は次のとおりです。何千もの科学者が依存している商用検証データに同様の問題が現れる可能性がある場合、試薬の選択、モデルトレーニングパイプライン、および治療プログラムがそのデータに依存しているすべての小規模なバイオテクノロジー、CRO、プラットフォーム企業、および製薬組織にとって、それは何を意味するのでしょうか?
私たちは静かに共存することを学んだ危機
これはどれも新しいものではありません。 Nature 's b roader coverage shows that reproducibility is a structural issue, not a one-off scandal. In its 2016 survey, two-thirds of responding researchers viewed current levels of reproducibility as a major problem , wi

出版への圧力、選択的な報告、貧弱な統計、そして気難しいプロトコルがすべて寄与しています ( Nature )。
この問題は前臨床生物学において特に深刻です。 Nature Reviews Drug Discoveryは、バイエルが引き受けた前臨床学術プロジェクトの約25％しか再現できなかったと報告したが、アムジェンはがん論文からの発見を再現しようとした場合の成功率が11％であると報告した（Nature Reviews Drug Discovery）。これらの数字は、非常に現実的な結果をもたらすため、創薬用の計算プラットフォームを構築するすべての企業にとって重要であるはずです。
製薬会社は、モデルが洗練されているように見えるからといって、単純にモデルの出力を信頼するわけではありません。
大手製薬会社が報告した前臨床研究の再現性の問題
大企業がすでに公表されている前臨床生物学の大部分を再現するのに苦労しているのであれば、予測、ターゲット、メカニズムを販売する AI ファーストの企業は、より厳しい努力を期待する必要があります。つまり、生データ、アッセイの詳細、試薬の来歴、複製、およびモデルが人工物ではなく実際の生物学を学習しているという証拠に対するより多くの要求が発生するはずです。実際には、データの再現性に対する信頼が弱いと、あらゆる提携サイクルが長くなります。
これが、外部データだけでは不十分な理由でもあります。 CRO は不可欠なパートナーですが、外部委託されたデータは、試薬履歴、プロトコルのドリフト、失敗した実験、オペレーターの影響、バッチアーチファクト、微妙なアッセイ動作など、生物学を解釈可能にする暗黙の実験コンテキストから切り離されて届くことがよくあります。多くの場合、これらの詳細によって、モデルが生物学を学習するかノイズを学習するかが決まります。弱く制御されたデータに基づいてトレーニングされたモデルは、隠れたアーティファクトを再現するのが非常に得意になりますが、コンピュータでは自信を持って、現実世界では間違っています。私が繰り返し戻ってくる教訓は、データの信頼は実験的に獲得する必要があるということです。

重要なフィードバック ループを自分で所有することで、じっくりと透過的に実行できます。
しかし、2 番目の、より巧妙な罠があり、今週発表された論文によって、私はこの罠にはっきりと焦点を当てることができました。
データが増えることと知識が増えることは同じではありません
より多くのデータ、より大きなモデル、より優れたパフォーマンスという、私たち全員が言語モデルと視覚モデルからインポートした慰めとなる物語があります。 Single-cell biology borrowed that playbook wholesale, scaling foundation models from corpora of one million cells to atlases of more than a hundred million, on the assumption that scale alone would unlock the same gains.
Alan DenAdel らによる Nature Methods の新しい研究では、その仮定が残酷なテストにさらされました (DenAdel et al., Nature Methods, 2026)。 Working from a 22.2-million-cell corpus, they pretrained 400 models across five architectures — from humble PCA and a variational autoencoder up to the Geneformer transformer — and ran 6,400 evaluation experiments . Jorge Bravo-Abad が要約したように、結果は厳粛なものです ( @bravo_abad )。
パフォーマンスはほぼすぐに飽和状態になってしまいました。 On cell-type classification, batch integration, and perturbation prediction, most models hit their ceiling at roughly 1% of the corpus — about 200,000 cells .それ以上に、何百万ものセルを追加しても、本質的には何も変わりません ( @bravo_abad )。多様性を高めても役に立ちませんでした。 Even spiking in genome-scale Perturb-seq data — feeding the models perturbed phenotypes rather than just healthy cells — failed to move the needle ( DenAdel et al. ).
まず、PCA やロジスティック回帰などの単純なベースラインは、トランスフォーマー ( @bravo_abad ) と一致するか、それを上回ることがよくありました。
第 2 に、最も強いモデルが勝ったのは、そのモデルが大きかったからではなく、そのトレーニング目標が実際の下流タスクと一致していたからです。
言い換えれば、何に基づいてトレーニングし、どのように組み立てるかということです。

質問は、どれだけ集めたかよりもはるかに重要です。
すでに再現性と格闘している分野において、大規模な細胞データセットが具体的なメリットをほとんど示さないときに、生物学的データの生成に台所のシンクを投入することは実際に良い考えなのでしょうか?
I do not think it is.事前トレーニング コーパスをスケーリングし続けたいという本能が、見返りのない膨大なコンピューティングを消費している可能性があります。本当の活用は、はるかに魅力的ではないところにあります。それは、高品質でタスクに関連したデータを厳選し、本当に答えようとしている質問にモデルを照合することです。
私たちがタンパク質を中心に会社を設立した理由
これらはまさに、「基礎モデル」というフレーズが誰かが大声で言うようになるずっと前から、長年にわたり私のチームと私を魅了してきた基礎的な構造生物学の疑問です。固定構造をもたないタンパク質は、どのようにして非常に特異的なことを行うのでしょうか?信号は実際にどこに存在するのでしょうか?そして、それをキャプチャするにはどのようなデータが必要ですか?
その魅力に惹かれて、私たちは Peptone を設立する旅に乗り出した理由です。Peptone はタンパク質に特化した会社であり、厳密な実験と深い専門知識がサポート機能ではなく、科学的発見への入り口となります。私たちは、最大限のデータを収集することを目指したわけではありません。私たちは、基礎となる物理学を理解している人々によって調査された、適切なターゲットに対して適切なデータを生成することに着手しました。
そして、最も困難な標的である本質的に無秩序なタンパク質に対するアプローチほどそのアプローチが必要なものはありません。
データだけでは不可能であることを専門家の物理学が明らかにしたこと
Nature Communications に掲載された私たちの最新の論文はこれを具体的にしています (Streit, Invernizzi, Bottaro, Tamiola & Lindorff-Larsen、2026)。私たちの細胞内のタンパク質残基の約 3 分の 1 は、本質的に無秩序です。それらは単一のきちんとした構造に折り畳まれるのではなく、幅広く変化する集合体を形成しています。彼らは

転写からシグナル伝達まであらゆるものを駆動しており、実験であれコンピューターであれ、特徴づけるのが難しいことで知られています。
https://www.nature.com/articles/s41467-026-73067-3
この研究では、マルチサーマルアンサンブルにおけるオンザフライ確率強化サンプリング (OPES) という強化されたサンプリング手法を導入しました。これにより、これらのシミュレーションで通常必要となる脆弱なパラメーター調整やレプリカ交換機構を必要とせずに、1 回のシミュレーションで 15 ～ 71 残基の無秩序なペプチドおよびタンパク質の構造ランドスケープを効率的に探索できるようになります (Nature Communications)。
The headline finding is the part that still gives me chills. We applied the method to ACTR, a 71-residue disordered transcriptional coactivator that folds only when it meets its partner. In its free state, ACTR looks like a floppy, structureless chain.しかし、私たちのシミュレーションは、実験だけでは見えない何かを明らかにしました。それは、複数のαヘリックスが協力して折り畳まれ、本物の三次接触を形成する、まれで低人口（〜3％）の一時的な構造状態のセットであり、タンパク質が何度も訪れたり放棄したりする、束の間の結合可能な形状です。
These hidden states are not a numerical artifact.それらは可逆的にサンプリングされ、わずか約 13 kJ/mol という適度な自由エネルギー障壁によって無秩序な基底状態から分離されており、アンサンブルの検証に使用した広範な NMR および SAXS データと決定的に一致しています (Nature Communications)。物理学と実験は一致します。そして、その見返りは、これらの一時的な、部分的に折りたたまれた立体構造が、まさに「治療不可能な」無秩序なタンパク質を突然アドレス可能にする種類の結合ポケットを抱えている可能性があるということです。
You do not find these hidden states by collecting more data. r を実行するとそれらを見つけることができます

自分が何を見ているのかを知っている人々によって解釈される物理学。
これが Peptone の背後にある理論全体です。当社は、オーダーメイドの物理現実シミュレーションを、深い専門知識と高度に専門化されたデータと組み合わせます。データは、棚から引っ張り出したり、一般的なアトラスから借用したものではなく、私たちが取り組んでいる特定の無秩序なターゲットに対して常に生成されます。当社の技術は、あらゆるサイズや複雑さの無秩序なタンパク質をモデル化し、従来の構造ベースの創薬が真っ直ぐ通り過ぎてしまう「目に見えない」ポケットを明らかにします。
Peptone Switzerland AG 研究所で開発された、本質的に無秩序なタンパク質の構造特性評価のための特注のロボット式超高速混合水素重水素交換対応質量分析システム。
これは、キッチンシンクのアプローチとはまったく逆です。より小型で、より特化したモデル。自分たちが作ったからこそわかるデータ。物理学は、経験的測定と人間の専門知識を融合させてうまく機能しました。
この専門家によるアプローチの具体的な成果は抽象的なものではありません。
医薬品開発が飛躍的に加速。モデルがアーティファクトではなく実際の生物物理学を学習すると、設計、測定、意思決定のすべてのサイクルがより緊密かつ高速になります。
まさに斬新な生物学的洞察。一時的な隠れ状態、

[切り捨てられた]

## Original Extract

AI drug discovery is only as strong as the experimental data behind it. At Peptone, we are investing in independent data generation. In a field where outsourced datasets and reproducibility challenges can shape model performance, owning the model-lab loop is becoming a strategic advantage.

0
Skip to Content
Kamil Tamiola
Home
Research
Photography
Blog
Open Menu
Close Menu
Kamil Tamiola
Home
Research
Photography
Blog
Open Menu
Close Menu
Home
Research
Photography
Blog
Biotech’s AI Revolution Will Be Won in the Lab
The next generation of AI for drug discovery will not be the biggest model trained on the most data. It will be the most specialized model trained on the right data — data we understand because we made it.
It starts with a simple act of trust
Imagine you are in a lab, and the experiment you are about to run depends on a simple act of trust. You open a supplier's catalogue, choose an antibody, and look at a validation image called a Western blot. For a non-specialist, a Western blot is basically a lab test that shows whether an antibody recognises the right protein: if the right dark band appears in the right place, scientists take it as evidence that the antibody works.
Now imagine someone zooms in, adjusts the contrast, and discovers that some of those bands may have been copied, flipped, painted over, or reused across different product pages. Suddenly the problem is no longer one questionable image. It is a much larger question: how much of the scientific supply chain are we trusting without truly verifying?
The bands labeled 1 through 4 are all identical to one another after a vertical flip, a horizontal flip or a 180 degree rotation. Source: How much of Thermo Fisher’s antibody data has been manipulated? by Reese Richardson May 28, 2026 .
A recent Nature article reported that catalogue entries for more than 100 Thermo Fisher Scientific antibodies contained images that appeared to have been manipulated — including images intended to demonstrate antibody quality and performance ( Nature ). Image alteration does not automatically mean the underlying products are defective, but the episode is a stark reminder that even trusted commercial data streams can carry hidden uncertainty.
What makes the Thermo Fisher case troubling is that the reported problems were not limited to a single figure or one isolated product page. Reese Richardson's analysis identified "more than 100 images bearing signs of manipulation" in the online primary-antibody catalogue: Western blot bands that appear identical after flipping and rotation, conspicuous brushstroke-like edits after contrast adjustment, repetitive blocks of background noise, and abrupt discontinuities in background texture. In one recurring pattern, dozens of pages showed verification blots sharing the same background — with 50 instances documented at the time of writing.
In academia, a pattern of apparent image manipulation on this scale would likely be career-ending.
A principal investigator facing allegations of duplicated bands, rotated blot features, and repeated backgrounds across many figures would expect institutional investigation, retractions, grant scrutiny, and possibly the collapse of their lab's credibility. So the uncomfortable question for industry is this: if comparable problems can appear in commercial verification data that thousands of scientists rely on, what does that mean for every small biotech, CRO, platform company, and pharma organisation whose reagent choices, model-training pipelines, and therapeutic programs may depend on that data?
A crisis we have quietly learned to live with
None of this is new. Nature 's b roader coverage shows that reproducibility is a structural issue, not a one-off scandal. In its 2016 survey, two-thirds of responding researchers viewed current levels of reproducibility as a major problem , with pressure to publish, selective reporting, poor statistics, and finicky protocols all contributing ( Nature ).
The issue is especially acute in preclinical biology. Nature Reviews Drug Discovery reported that Bayer could replicate only about 25% of the preclinical academic projects it took on, while Amgen reported an 11% success rate when trying to recreate findings from cancer papers ( Nature Reviews Drug Discovery ). Those numbers should matter t o every company building computational platforms for drug discovery, because they carry a very practical consequence:
Pharma will not simply trust the model output because the model looks sophisticated.
Pre-clinical research reproducibility issues reported by Big Pharma companies
If large companies have already struggled to reproduce a major fraction of published preclinical biology, any AI-first company selling predictions, targets, or mechanisms should expect tougher diligence — more requests for raw data, assay details, reagent provenance, replication, and evidence that the model is learning real biology rather than artifacts. In practice, weak trust in data reproducibility lengthens every partnering cycle.
This is also why external data alone is not enough. CROs are essential partners, but outsourced data often arrives detached from the tacit experimental context that makes biology interpretable: reagent history, protocol drift, failed runs, operator effects, batch artifacts, and subtle assay behaviours. Those details often decide whether a model learns biology or learns noise. A model trained on weakly controlled data can become very good at reproducing hidden artifacts — confident in silico , and wrong in the real world. The lesson I keep coming back to is that trust in data has to be earned experimentally , repeatedly and transparently, by owning the critical feedback loops yourself.
But there is a second, subtler trap — and a paper published this week brought it into sharp focus for me.
More data is not the same as more knowledge
There is a comforting story we have all imported from language and vision models: more data, bigger model, better performance. Single-cell biology borrowed that playbook wholesale, scaling foundation models from corpora of one million cells to atlases of more than a hundred million, on the assumption that scale alone would unlock the same gains.
A new study in Nature Methods by Alan DenAdel and colleagues put that assumption to a brutal test ( DenAdel et al., Nature Methods , 2026 ). Working from a 22.2-million-cell corpus, they pretrained 400 models across five architectures — from humble PCA and a variational autoencoder up to the Geneformer transformer — and ran 6,400 evaluation experiments . As Jorge Bravo-Abad summarised it, the result is sobering ( @bravo_abad ).
Performance saturated almost immediately. On cell-type classification, batch integration, and perturbation prediction, most models hit their ceiling at roughly 1% of the corpus — about 200,000 cells . Beyond that, adding millions more cells changed essentially nothing ( @bravo_abad ). More diversity did not help. Even spiking in genome-scale Perturb-seq data — feeding the models perturbed phenotypes rather than just healthy cells — failed to move the needle ( DenAdel et al. ).
First, simple baselines like PCA and logistic regression often matched or beat the transformers ( @bravo_abad ).
Second, the strongest model won not because it was bigger, but because its training objective was aligned with the actual downstream task .
In other words: what you train on and how you frame the question matter far more than how much you collect.
In a field already wrestling with reproducibility, is throwing the kitchen sink at biological data generation actually a good idea — when massive cellular datasets show so little tangible benefit?
I do not think it is. The instinct to keep scaling pretraining corpora may be burning enormous compute for no return. The real leverage sits somewhere far less glamorous: curating high-quality, task-relevant data, and matching the model to the question you are genuinely trying to answer.
Why we built a company around proteins
These are precisely the questions in fundamental structural biology that have fascinated my team and me for years — long before "foundation model" was a phrase anyone said out loud. How does a protein that has no fixed structure still do something exquisitely specific? Where does the signal actually live? And what kind of data do you need to capture it?
That fascination is why we embarked on the journey to build Peptone: a company devoted to proteins, where experimental rigour and deep expert knowledge are not support functions but the gate to scientific discovery. We did not set out to collect the most data. We set out to generate the right data, for the right targets, interrogated by people who understand the physics underneath.
And nowhere is that approach more necessary than with the hardest targets of all: intrinsically disordered proteins.
What expert physics reveals that data alone cannot
Our latest paper in Nature Communications makes this concrete (Streit, Invernizzi, Bottaro, Tamiola & Lindorff-Larsen, 2026). Roughly a third of all protein residues in our cells are intrinsically disordered — they do not fold into a single, tidy structure, but populate broad, shifting ensembles. They drive everything from transcription to signalling, and they are notoriously difficult to characterise, whether you reach for an experiment or a computer.
https://www.nature.com/articles/s41467-026-73067-3
In this work we introduced an enhanced-sampling approach — On-the-fly Probability Enhanced Sampl ing, or OPES, in a mu ltithermal ensemble — that lets a single simulation efficient ly explore the conformational landscape of disordered peptid es and proteins from 15 to 71 residues, without the brittle parameter tuning and replica-exchange machinery these simulations normally demand ( Nature Communications ).
The headline finding is the part that still gives me chills. We applied the method to ACTR, a 71-residue disordered transcriptional coactivator that folds only when it meets its partner. In its free state, ACTR looks like a floppy, structureless chain. But our simulations revealed something experiments alone cannot see: a rare, low-population (~3%) set of transiently structured states in which multiple α-helices fold cooperatively and form genuine tertiary contacts — a fleeting, binding-competent shape that the protein visits and abandons, over and over.
These hidden states are not a numerical artifact. They are reversibly sampled, separated from the disordered ground state by a modest free-energy barrier of only ~13 kJ/mol, and — critically — consistent with extensive NMR and SAXS data we used to validate the ensemble (Nature Communications). The physics and the experiments agree. And the payoff is that these transient, partially folded conformations may harbor exactly the kind of binding pockets that make an "undruggable" disordered protein suddenly addressable.
You do not find these hidden states by collecting more data. You find them by running the right physics, interpreted by people who know what they are looking at.
This is the whole thesis behind Peptone . We combine bespoke physical-reality simulations with deep expert knowledge and highly specialized data — data that is always generated for the specific disordered target we are working on , not pulled off a shelf or borrowed from a generic atlas. Our technology models disordered proteins of any size or complexity to reveal the "invisible" pockets that conventional structure-based drug discovery walks straight past.
A bespoke robotic and ultra-fast mixing hydrogen-deuterium exchange-enabled mass spectrometry system for structural characterisation of Intrinsically Disordered Proteins, developed in Peptone Switzerland AG laboratories.
It is the exact inverse of the kitchen-sink approach. Smaller, more specialized models. Data we understand because we made it. Physics done well, blended with empirical measurement and human expertise.
The tangible fruits of this expert approach are not abstract:
Exponentially accelerated drug development. When your model learns real biophysics instead of artifacts, every cycle of design, measurement, and decision gets tighter and faster.
Genuinely novel biological insight. Transient hidden states,

[truncated]
