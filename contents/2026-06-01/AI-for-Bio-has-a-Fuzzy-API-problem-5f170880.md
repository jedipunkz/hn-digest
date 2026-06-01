---
source: "https://ankitg.me/blog/2026/05/04/fuzzy_api.html"
hn_url: "https://news.ycombinator.com/item?id=48344220"
title: "AI for Bio has a Fuzzy API problem"
article_title: "AI for Bio has a Fuzzy API problem"
author: "sebg"
captured_at: "2026-06-01T00:34:07Z"
capture_tool: "hn-digest"
hn_id: 48344220
score: 2
comments: 0
posted_at: "2026-05-31T09:21:58Z"
tags:
  - hacker-news
  - translated
---

# AI for Bio has a Fuzzy API problem

- HN: [48344220](https://news.ycombinator.com/item?id=48344220)
- Source: [ankitg.me](https://ankitg.me/blog/2026/05/04/fuzzy_api.html)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T09:21:58Z

## Translation

タイトル: AI for Bio にはファジー API の問題がある
説明: 創薬は

記事本文:
AI for Bio にはファジー API の問題があります
「バイオのためのAI」が再び熱くなっています。現在の興奮を考慮して、機械学習のアプリケーション ドメインとして生物学が実際に独特に難しい理由について少し共有したいと思いました。その理由は、明らかに生物学が複雑であるというだけではありません。 ML は、複雑な多くのことに適しています。より深い理由は、創薬には、他の場所で最新の ML を非常に強力にしたようなクリーンなフィードバック ループやクリーンなインターフェイスが存在しないことです。
ソフトウェアでは、API をクリーンアップすることに慣れています。あるチームがバックエンド サービスを構築してエンドポイントを公開し、別のチームがその上に構築することができます。インターフェースは型付きです。オブジェクトはコントラクトを満たしているか、満たしていないかのどちらかです。何かが壊れた場合、通常はその失敗をバグとして追跡し、コードを修正し、テストを再実行して、再度出荷することができます。これは、10 億ドル規模の企業が定期的に 1 つのインターフェイス (データベース用の Supabase、検索用の Exa、GPU コンピューティング用の NVIDIA など) を満たすように構築されている場合に当てはまります。
創薬についても同じように想像したくなります。
target = target_discovery (病気)
薬物 = 薬物デザイン (ターゲット)
医学 = 臨床試験 (薬剤)
ターゲット発見によりターゲットが提供されます。ドラッグデザインは分子を与えます。臨床試験によって、効果があるかどうかが分かります。 1 社が各インターフェースを満たします。
残念ながら、生物学はクリーンな API を公開していません。ターゲット検出の出力は実際にはターゲットではありません。これは、ある生物学的プロセスをある方向で、ある組織で、ある患者サブセットで、ある疾患段階で調節すると、許容できない毒性を伴わずに有用な臨床効果を生み出すだろう、という確率論的な仮説です。ドラッグデザインの成果物は実際には薬ではありません。これは介入の候補であり、その価値はタールが以下であるかどうかによって決まります。

仮説が正しかったかどうか、モダリティが適切かどうか、分子が適切な組織に到達するかどうか、十分な選択性があるかどうか、安全マージンが許容できるかどうか、製造可能かどうか、防御可能な知財の地位があるかどうか、未知の競争環境が実現するかどうか、実行可能な臨床戦略に適合するかどうか。臨床試験の成果は単なる「治療法」ではありません。これは、患者の選択、エンドポイントの選択、投与計画、施設での実施、統計的検出力、標準治療、および規制の解釈を通じてフィルタリングされた結果です。
したがって、API はあいまいです。分子が間違った組織に到達するまで、ターゲットは検証されたように見えます。分子は、病気の生物学が間違っていたために失敗するまでは素晴らしく見えます。動物モデルは、人間の病気が有意に異なるまでは説得力があるように見えます。たとえその薬がより狭い患者集団に効果があったとしても、試験は否定的に見える可能性があります。下流または上流のステージは、他のステージの特定の仮定をエンコードします。私にとって、それが核心的な問題です。バイオ用 AI にはあいまいな API の問題があります。ソフトウェアでは、優れた API が複雑さを隠します。生物学では、隠れた複雑さによってプログラムが誤って停止してしまうことがあります。このエッセイでは、その曖昧さがどこに現れるか、つまり標的発見、創薬、臨床開発、そしてそれが機械学習を使用してこの分野に革命を起こす際の課題と機会の両方をもたらす場所について説明します。
第一段階の創薬には、体内で起こるある種の有害なプロセスを阻止する介入を設計することが含まれます。実際、現代の治療法開発は通常次のようになります。
特定の患者集団の疾患を引き起こす原因生物学を特定します。
その生物学を調節する化学的または生物学的介入を設計します。
テスト

モデルシステム（細胞、動物など）に介入して、それがヒトに対して安全であり、おそらく有効であるという証拠を構築します。
臨床試験を実施して、人間に効果があるかどうかを確認します。
そのきれいなリストには多くのことが隠されています。介入自体は、モダリティと呼ばれるさまざまな形を取ることができます。これらには、特定の分子をブロックする小分子、特定の分子に毒素を送達する抗体、タンパク質の設計をブロックする RNA の断片、悪いプロセスを永久に削除する生体内編集、異物を認識するように免疫系を再プログラムする方法などが含まれます。その後、臨床開発は通常 3 つのフレーズを経て進行し、人体における安全性と有効性のさまざまな側面をテストします。
この投稿では、プロセス全体を 3 つの段階に単純化します。
ターゲットの発見 : どの生物学を調整する必要がありますか?
前臨床デザインとトランスレーションデザイン：どの介入がそれを調節できるか?この介入は、人体で試すことを正当化するほど、モデルシステムでは安全で効果的であるように見えますか?
臨床開発 : それは実際に対象となる人々を助けますか?
機械学習は 3 つの段階すべてで重要になります。ただし、重要な ML の種類と問題の難易度は、段階によって大きく異なります。
私が標的の発見について話すとき、本質的に私たちが言いたいのは、「標的とする価値のある何かの仮説を見つけるために生物学の研究を行う」ということです。歴史的に、多くの標的発見研究は学術界で行われてきました。 Broad Institute の DepMap のような大規模な取り組みは、がん細胞のどのような混乱ががん細胞の成長に影響を与えるかを特定するのに役立ちました。研究者らはそれを利用して細胞増殖を促進する潜在的な遺伝子を特定し、それらの遺伝子をブロックすれば抗がん剤が可能になるという仮説を立てた。これらは人体実験ではないことに注意してください。彼らはいた

通常は体外で増殖させた癌細胞株、または場合によっては患者由来の細胞株で行われます。
ゲノムワイド関連研究 (GWAS) などの他の取り組みでは、英国バイオバンクのような人口規模の配列決定の取り組みから収集された実際のヒトデータを使用しようとしています。アイデアは非常にシンプルです。十分な人数があれば、病気のリスクに関連する遺伝的変異を特定できます。 GWAS がヒットすると、ゲノムの領域が病気に関連していることがわかる場合があります。どの遺伝子が重要か、どの細胞型で、どの経路を介して、どの患者サブセットで、あるいはどのように治療介入すべきかを自動的に教えてくれるわけではありません。これもファジー API の問題です。人間の遺伝学によって返される「ターゲット」は、多くの場合、きれいな治療対象ではありません。これは物理実験によってさらに検証する必要がある手がかりです。そのため、歴史的には、ターゲット発見の中核となるディープラーニングはあまりありませんでした。遺伝子制御などのために、統計遺伝学手法、ネットワーク手法、深層学習ツールが存在します。しかし、私たちが実際に気にかけている質問に答えることができる統一モデルはありません。
「細胞型 y の患者 x に介入 z を行ったら、何が起こるでしょうか?」。
この問題の枠組みにより、この分野ではいわゆる「仮想セル」または「仮想人間」モデルが新たに注目されています。アイデアは大まかに言うと、細胞や人の状態に関する大量のデータを収集し、さまざまな方法でそれらに摂動を加え、トランスクリプトミクス、プロテオミクス、イメージング、細胞増殖、機能的表現型、臨床データなどの多くの読み取り値を測定し、それらの間の関係を学習するモデルをトレーニングするというものです。 LLM からのスケーリング則は、リソースの割り当てを考える方法の枠組みを提供し、初期の取り組みのいくつかは、これらのデータにも同様のスケーリング則があることを示唆しています。 Tahoe、Noetik、Re などの企業

cursion は皆、これの何らかのバージョンを追求しています。これは良い方向です。
これも非常に難しいことが分かります。まず、既存のデータはまばらで品質が低いです。数百万の遺伝子座、多くの考えられる細胞型、多くの病状、多くの介入があり、患者数が数十万人、あるいは数百万人しかいない場合でも、次元の呪いという点では依然として問題が残ります。バイオバンク データは、人間に自然に発生する変動を捉え、それを病気の転帰と相関させるため、強力です。これは、効果が自然に発生する変動を反映する薬を見つけるのに最適です。逆に言えば、このデータは薬物による恣意的な混乱の結果を捉えることはできません。人間を動揺させ、様子を見て何が起こるかを見る、倫理的に許容される方法はそれほど多くありません。
そこで私たちはモデルシステムに頼ります。細胞、オルガノイド、マウス、ラット、イヌ、ヒト以外の霊長類などを使用します。しかし、細胞株には免疫系、血管系、肝臓、マイクロバイオーム、または人体の完全な構成が備わっていません。動物には体がありますが、人間には体がありません。人工疾患モデルは有用ですが、私たちが実際に関心を持っているヒトの疾患プロセスを捉えていない可能性があります。モデル種で前臨床的に効果がある薬剤はヒトの臨床試験では定期的に失敗するため、これが真実であることがわかっています。
これにより、モデル システムと人間の間にシミュレーションと現実のギャップが生じます。さらに悪いことに、このシミュレーション データで機械学習モデルをトレーニングしているため、シミュレーションのモデルを構築し、それを現実に移行させようとしています。言い換えれば、sim^2 と実際のギャップがあり、優れた物理エンジンがありません。
だからこそ、ターゲットの発見は非常に難しいのです。より優れたモデルが必要なだけではありません。私たちはより良いデータ、より良い摂動システム、より良い因果推論、より良い人間関連の分析、そして賭けを必要としています。

予測と実験的検証の間のループを閉じるための 3 つの方法。誤解しないでください。企業がこれらの問題を解決し、ターゲット発見のための新しい ML モデルを開発するために真剣に取り組んでいることを見て興奮しています。また、歴史的には、標的を発見することが、その標的に対する医薬品を自社で製造する以外に効果的に収益化できるかどうかは明らかではありませんでしたが、Cartography Bio のような企業はそうではないことを証明しつつあります。
それでも、難しい部分は残ります。「ターゲット」オブジェクトがあいまいであるということです。モデルは、遺伝子、経路、細胞状態、またはバイオマーカーを指定できます。しかし、臨床的に関連する対象は、因果関係があり、方向性があり、状況に応じた介入仮説です。
これはプロセスの中で最も広く議論される段階です。ここには明らかに大きな期待があります。ボルツ、チャイ、ナブラなどのモデルは、構造の予測、バインダーの生成、タンパク質相互作用の推論に使用できます。 Dyno は機械学習を使用して AAV カプシドを設計しており、他の企業は生成モデルを使用して遺伝子医薬品のペイロードを設計しています。これらの中には、自然言語 LLM を使用しているものもあれば、コンピューター ビジョンと LLM で機能したことを学び、生物特有のユースケースをサポートする新しいモデル アーキテクチャを設計しているものもあります (Nabla、Chai、Dyno などがすべてこれに該当します)。それは素晴らしいことです。
ここで問題が発生します。創薬に対する一般的な ML パーソンの視点では、この段階が次のようなクリーンな API として認識されます。その実装は次のようになります。
def Drug_design_process (ターゲット、既知の阻害剤):
new_drug = なし
while too_similar ( new_drug 、 known_inhibitors ):
new_drug = design_and_optimize (ターゲット)
新しい薬を返す
素朴な見方では、while ループを適切に実行できれば、勝利することができます。 Reverie を始めたとき、確かにそうやってうまくいくと思いました。コンベにも合います

コアコンピテンシーに焦点を当てるための全国的なスタートアップの知恵。あなたが構造生物学に興味のある機械学習エンジニアであれば、最高の医薬品設計エンジンを構築するだけです。他の人にターゲットを見つけてもらいましょう。他の人にトライアルを実行してもらいます。
残念ながら、これは生物学の実際の仕組みではありません。有名なターゲット、つまり、優れた論文、信頼できるヒト遺伝学、または明白な臨床的根拠を持つターゲットを意味しますが、多くの場合、多くの競合他社が存在します。知っているものもあれば、知らないものもあります。 1 位でない場合は、有意義に改善する必要があります。しかし、通常、競合他社の完全な有効性データ、安全性データ、製剤の詳細、さらにはその構造さえも、ずっと後になるまでアクセスできません。それに加えて、あなたがアメリカの企業であれば、中国に迅速かつ効果的に行動する競合他社が存在する可能性が高く、その競合他社は開発コストの数分の一で自社の医薬品を製薬パートナーに提供できるでしょう。また、動物のデータをあなたよりも早く/安く生成でき、米国で行う場合の数分の一の費用でフェーズ 1 を実行できます。
したがって、混雑したターゲットを避けて、新しい生物学を追求するかもしれません。それは正しい行動かもしれません。おそらく、新しいターゲットを特定するためにターゲット発見プラットフォーム会社または独自のデータ プラットフォームを使用している可能性があります。

[切り捨てられた]

## Original Extract

Drug discovery doesn

AI for Bio has a Fuzzy API problem
“AI for bio” is getting hot again. Given the excitement in the current moment, I thought I’d share a bit about what actually makes biology uniquely hard as an application domain for machine learning. The reason is not simply that biology is complicated, though it obviously is. ML is good at many things that are complicated. The deeper reason is that drug discovery does not have the kind of clean feedback loops and clean interfaces that made modern ML so powerful elsewhere.
In software, we are used to clean APIs. One team can build a backend service, expose an endpoint, and another team can build on top of it. The interface is typed. The object either satisfies the contract or it does not. If something breaks, you can usually trace the failure to a bug, fix the code, rerun the test, and ship again. This is so much the case that billion dollar companies are regularly built satisfying exactly one interface (e.g. Supabase for databases, Exa for search, NVIDIA for GPU compute).
It is tempting to imagine drug discovery the same way:
target = target_discovery ( disease )
drug = drug_design ( target )
medicine = clinical_trial ( drug )
Target discovery gives you a target. Drug design gives you a molecule. Clinical trials tell you whether it works. One company satisfies each interface.
Unfortunately biology does not expose clean APIs. The output of target discovery is not really a target. It is a probabilistic hypothesis that modulating some biological process, in some direction, in some tissue, in some patient subset, at some disease stage, will produce a useful clinical effect without unacceptable toxicity. The output of drug design is not really a drug. It is a candidate intervention whose value depends on whether the target hypothesis was right, whether the modality is appropriate, whether the molecule reaches the right tissue, whether it has enough selectivity, whether the safety margin is acceptable, whether it can be manufactured, whether it has a defensible IP position, whether the unknown competitive landscape materializes, and whether it fits a viable clinical strategy. The output of a clinical trial is not simply a “cure”. It is an outcome filtered through patient selection, endpoint choice, dosing regimen, site execution, statistical power, standard of care, and regulatory interpretation.
So the API is fuzzy. A target can look validated until the molecule hits it in the wrong tissue. A molecule can look great until it fails because the disease biology was wrong. An animal model can look convincing until the human disease is meaningfully different. A trial can look negative even though the drug might have worked in a narrower patient population. The downstream or upstream stages encode specific assumptions in other stages. To me, that is the core problem: AI for bio has a fuzzy API problem. In software, good APIs hide complexity. In biology, the hidden complexity inadvertently kills programs. This essay is about where that fuzziness shows up: target discovery, drug design, and clinical development, and where that poses both challenges and opportunities to use machine learning to revolutionize the field.
The first order, drug discovery involves designing an intervention that stops some sort of deleterious process occurring inside the body. In practice, modern therapeutic development usually looks something like this:
Determine the causal biology driving a disease in a particular patient population.
Design a chemical or biological intervention that modulates that biology.
Test that intervention in model systems (cell, animals, etc) to build evidence that it is safe and plausibly effective in humans.
Run clinical trials to determine whether it works in humans.
That clean list hides a lot. The intervention itself can take many forms, called modalities. They can involve small molecules that block a particular molecule, antibodies that deliver toxins to certain molecules, pieces of RNA that block protein design, in-vivo edits that permanently delete bad processes, methods for reprogramming the immune system to recognize a foreign body, and several others. Clinical development then generally proceeds through three phrases, testing various aspects of safety and efficacy in humans.
For the purposes of this post, I’ll simplify the whole process into three stages:
Target discovery : What biology should we modulate?
Preclinical Design and Translation design : What intervention can modulate it? Does the intervention look safe and effective enough in model systems to justify trying it in humans?
Clinical development : Does it actually help the intended human population?
Machine learning can matter in all three stages. But the type of ML that matters, and the difficulty of the problem, varies a lot by stage.
When I talk about target discovery, in essence what we’re saying is “do biology research to find a hypothesis of something worth targeting”. Historically, a lot of target discovery research happened in academia. Large scale efforts like the Broad Institute’s DepMap helped identify what perturbations in cancer cells affected their growth. Researchers used that to identify potential genes that drove cell growth, and hypothesized that blocking those genes would enable anti cancer drugs. Note that these were not experiments on people. They were usually on cancer cell lines that had been grown ex-vivo, or occasionally on patient derived lines.
Other efforts like genome wide association studies (GWAS) attempt to use real human data collected from population scale sequencing efforts like the UK Biobank. The idea is pretty simple — with enough people, you can identify genetic variants associated with disease risk. A GWAS hit might tell you that a region of the genome is associated with a disease. It does not automatically tell you which gene matters, in which cell type, through which pathway, in which patient subset, or how you should intervene therapeutically. This is the fuzzy API problem again. The “target” returned by human genetics is often not a clean therapeutic object. It is a clue that needs to be further validated by physical experiments. So historically, there has not been much deep learning at the core of target discovery. There have been statistical genetics methods, network methods, and deep learning tools for things like gene regulation. But not a unified model that could answer the question we actually care about:
“if I intervene on patient x in cell type y with intervention z, what will happen?”.
That problem framing is making the field newly excited about so-called “virtual cell” or “virtual human” models. The idea is roughly: collect lots of data about cell/person states, perturb them in various ways, and measure many readouts — transcriptomics, proteomics, imaging, cell growth, functional phenotypes, clinical data — and train models that learn the relationships among them. Scaling laws from LLMs give a framework for how to think about allocating resources, and some early efforts suggest there are similar scaling laws in these data too. Companies like Tahoe, Noetik, and Recursion are all pursuing some version of this. This is a good direction.
This also turns out to be extremely hard. First, the existing data is sparse and low quality. If you have millions of genetic loci, many possible cell types, many disease states, many interventions, and only hundreds of thousands or even millions of patients, you are still in trouble in terms of the curse of dimensionality. Biobank data is powerful because it captures naturally occurring variation in humans and correlates that to disease outcomes. This is great for finding drugs whose effect mirrors naturally occurring variation. On the flipside, this data cannot capture the outcome of arbitrary drug perturbations. There are not many ethically acceptable ways to perturb a human, wait, and see what happens.
So we fall back on model systems: we use cells, organoids, mice, rats, dogs, non-human primates, and others. But, a cell line does not have an immune system, a vascular system, a liver, a microbiome, or the full context of a human body. An animal has a body, but not a human body. An engineered disease model can be useful, but it may not capture the human disease process we actually care about. We know this is true because drugs that work in model species preclinically regularly fail in human clinical trials.
This creates a sim-to-real gap between the model system and the human. Worse, since we are training machine learning models on this sim data, we are building models of the simulation and trying to get them to transfer to reality. In other words, we have sim^2-to-real gap and no great physics engine.
That is why target discovery is so hard. It is not just that we need better models. We need better data, better perturbation systems, better causal inference, better human-relevant assays, and better ways to close the loop between prediction and experimental validation. Don’t get me wrong. I’m excited to see that companies are making serious attempts to solve these problems and develop new ML models for target discovery. Also, historically it wasn’t clear that discovering targets could be effectively monetized other than by making drugs yourself against those targets, but companies like Cartography Bio are proving otherwise.
Still, the hard part remains: the “target” object is fuzzy. A model can nominate a gene, pathway, cell state, or biomarker. But the clinically relevant object is a causal, directional, context-specific intervention hypothesis.
This is the stage of the process that is most widely discussed. There is obviously enormous promise here. Models like Boltz, Chai, and Nabla can be used to predict structures, generate binders, and reason about protein interactions. Dyno is using machine learning to design AAV capsids, and other companies are designing payloads for genetic medicines with generative models. Some of these are using natural language LLMs while others are taking the lessons for what worked in computer vision and LLMs and designing new model architectures to support bio-specific use-cases (Nabla, Chai, Dyno, etc are all this). That’s great.
Now here is the problem: the common ML person view of drug discovery sees this stage as a clean API whose implementation looks something like this:
def drug_design_process ( target , known_inhibitors ):
new_drug = None
while too_similar ( new_drug , known_inhibitors ):
new_drug = design_and_optimize ( target )
return new_drug
The naive view is that as long as you can run the while loop well, you win. I certainly thought that’s how it worked when we started Reverie. It also matches the conventional startup wisdom to focus on your core competency. If you’re a machine learning engineer interested in structural biology, just build the best drug design engine. Let someone else find the target. Let someone else run the trial.
Unfortunately, this is not how biology works in practice. For any well-known target — meaning a target with a good paper, credible human genetics, or obvious clinical rationale — there will often be many competitors. Some you know about, some you do not. If you are not first, you need to be meaningfully better. But you usually do not have access to your competitors’ full efficacy data, safety data, formulation details, or usually even their structure until much later. On top of that, if you’re an American company, there is a high probability you have a competitor in China moving quickly and effectively, who can offer up their drug to a pharma partner for a fraction of what your development costs. They can also generate animal data faster/cheaper than you can, and run a Phase 1 for a fraction of the price of doing it in the US.
So maybe you avoid crowded targets and go after novel biology. That can be the right move. Perhaps you use a target discovery platform company or a proprietary data platform to identify new target

[truncated]
