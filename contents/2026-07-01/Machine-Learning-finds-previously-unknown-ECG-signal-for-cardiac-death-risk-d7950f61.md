---
source: "https://www.scientificamerican.com/article/ai-finds-hidden-ecg-signal-that-predicts-sudden-cardiac-death-risk/"
hn_url: "https://news.ycombinator.com/item?id=48745917"
title: "Machine Learning finds previously unknown ECG signal for cardiac death risk"
article_title: "AI finds hidden ECG signal that predicts sudden cardiac death risk | Scientific American"
author: "ck2"
captured_at: "2026-07-01T12:54:10Z"
capture_tool: "hn-digest"
hn_id: 48745917
score: 1
comments: 0
posted_at: "2026-07-01T12:53:36Z"
tags:
  - hacker-news
  - translated
---

# Machine Learning finds previously unknown ECG signal for cardiac death risk

- HN: [48745917](https://news.ycombinator.com/item?id=48745917)
- Source: [www.scientificamerican.com](https://www.scientificamerican.com/article/ai-finds-hidden-ecg-signal-that-predicts-sudden-cardiac-death-risk/)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T12:53:36Z

## Translation

タイトル: 機械学習により、心臓死亡リスクに関するこれまで知られていなかった ECG 信号が発見される
記事タイトル: AI が心臓突然死のリスクを予測する隠れた ECG 信号を発見 |サイエンティフィック・アメリカン
説明: 新しいモデルは、定期的な心電図検査から心臓突然死のリスクが高い人に警告を発し、心臓の電気活動の危険信号を明らかにします

記事本文:
メイン コンテンツにスキップ Scientific American 2026 年 6 月 30 日
AI が 100 年前の ECG テクノロジーに隠された心臓リスク信号を発見
新しいモデルは、定期的な心電図検査で心臓突然死のリスクが高い人に警告を発し、心臓の電気活動の危険信号を明らかにします
Jacek Krywko著、Eric Sullivan編集
ECG は、心臓の電気活動を山と谷の波形に変換します。新しい研究では、AI モデルのトレーニングとテストに 440,000 件を超えるそのような録音が使用されました。
植え込み型除細動器が数十年にわたって多くの致死性不整脈を阻止できてきたにもかかわらず、米国では心臓突然死により毎年30万人以上が亡くなっています。今日の主な問題は、心停止を止める装置ではありません。それは誰がそれを必要としているかを把握することです。新しい Nature 研究 では、カリフォルニア大学バークレー校の准教授 Ziad Obermeyer 率いるチームが、10 秒間の心電図からその質問に答えるようにニューラル ネットワークを訓練しました。次に、2 番目のニューラル ネットワークをトレーニングして、最初のニューラル ネットワークが何を重視しているかを明らかにしました。
2 つのモデルのセットアップは、医療における AI のより大きな野心を示しています。それは、人間の専門家が自分たちで見て確認できる新しい手がかりを機械に表面化させることです。オーバーマイヤー氏のチームは、最初のネットワークを使用してリスクを予測し、2 番目のネットワークを使用して、その予測を通常の ECG 上の目に見える特徴に変換し、心臓専門医が認識できるようにしました。
誰が除細動器を受けるべきかを判断するために、心臓専門医は現在、左心室が拍動ごとにどれだけの血液を送り出すかを示す超音波測定値、つまり左心室駆出率（LVEF）として知られる測定値に頼っています。オーバーマイヤー氏は、それは完璧には程遠いと指摘しています。 「心停止で突然死亡する人の多くは、これまでに超音波検査を受けたことがないか、あるいは受けたことがあり、

「結果は正常でした。同時に、そのテストの強度で埋め込まれた除細動器のほとんどは、最終的に点火することはありませんでした。「リスクが高そうに見えた人が、結局はそれほどハイリスクではなかったことがよくあります」とオーバーマイヤー氏は言います。この問題を回避するために、彼のチームはより良いリスクマーカーを探し始めました。
科学ジャーナリズムの支援について
この記事を気に入っていただけた場合は、 を購読して、受賞歴のあるジャーナリズムをサポートすることを検討してください。サブスクリプションを購入することで、今日の世界を形作る発見やアイデアに関する影響力のあるストーリーを将来にわたって確実に伝えることに貢献することになります。
心電図 (ECG) は心臓の電気活動を測定するもので、比較的安価でほぼ汎用的です。しかし、心臓専門医は何十年にもわたって ECG 波形を研究してきたにもかかわらず、心停止の高いリスクを確実に示すパターンを発見したことがありませんでした。彼のチームは人間による検査では見逃していたパターンを見つけるためにディープラーニングに目を向けました。チームが選択したアルゴリズムは、64 層の残差ニューラル ネットワーク (ResNet) でした。 「これは誰もが使用する主力モデルのようなものです。 「何も興味深いものはありません。興味深いのは、そこから学習されたデータです。」とオーバーマイヤー氏は言います。
ネットワークに情報を提供するために、オーバーマイヤー氏のグループは、スウェーデンの約 18 万人の患者から採取した 44 万件以上の ECG を国の死亡証明書と照合して、この種のデータセットとしては初の人口規模のデータセットの 1 つを構築しました。スウェーデンのデータに基づいてトレーニングされた、一般的な ResNet は、患者の約 2.2 パーセントに相当する高リスク群にフラグを立てました。チームが米国と台湾からの別のデータセットでモデルをテストしたときも信号は持続し、これがスウェーデンの人口や心電図装置の異常ではないことを示唆しました。その小さなグループ内での心臓突然死の年間発生率は 7% に達しました。

標準的な超音波検査でフラグが立てられた患者の割合は 4.6 パーセントです。さらに、アルゴリズムが抽出した患者の 86% 以上は、従来の LVEF マーカーによってフラグが立てられていませんでした。従来の方法では、このような患者の多くは除細動器なしで帰宅させられていたでしょう。
「これが機能していることを確認した後、このモデルが高リスクの人々の ECG 波形で何を確認しているのかを理解したいと思いました」と Obermeyer 氏は言います。顕著性マップのような標準的な AI 解釈ツールは、ニューラル ネットが波形のどの部分を最も重み付けしたかを強調表示できますが、そこで止まります。人間の心臓専門医は、ECG トレース上で異常な何かを発見すると、その異常な波をスケッチすることができます。ニューラル ネットワークはデフォルトではできません。そこで、オーバーマイヤー氏と彼の同僚は、まさにそれを行うための生成 AI モデルを構築しました。 「その仕事は、最初のモデルにとって高リスクに見える ECG 波形を生成することでした」と Obermeyer 氏は言います。
元のネットワークと組み合わせ、そのリスク スコアに基づいて、生成モデルは実際の低リスク患者の ECG を段階的に再加工し、同じトレースの高リスク バージョンにスムーズに変形させました。モデルが重視した特徴の多くは、心臓専門医にはすでに馴染みのあるものでした。
しかし、これまで医学文献には記載されていなかった特徴が 1 つあります。それは、aVL と呼ばれる 1 つの ECG リードの微妙な不明瞭さであり、心臓の電気信号が筋肉内を移動する際に断片化していることを示唆しています。
ジョンズ・ホプキンス大学の生物医学工学者、チャンシン・ライ氏は、『ネイチャー』誌に付随する分析を執筆したが、この研究には関与していないが、これがこの研究が際立っている理由だと述べている。 「ECG は 100 年以上前から存在しており、この種のデータは何世代にもわたる心臓専門医によって慎重に評価されてきました」と彼は言います。 「私たちは芸術から新しい知識を抽出しました

人工知能モデル。」
高リスク患者の一部については、チームは心臓磁気共鳴画像法（MRI）スキャンも行った。これらのスキャンでは、生成モデルが生成した合成波形に適合する方法で、心臓の電気信号に干渉する可能性がある不整脈に関連する、わずかなびまん性の線維症や瘢痕が示されました。オーバーマイヤー教授は、線維症との関連性は予備的なものであり、生検ではまだ確認されていないと警告している。
この発見は興味深いものではあるが、治療の指針となるまでには至っていない。 「これは重要な研究分野です」と、シーダーズ・サイナイ医療センターの心停止予防センターを所長するスミート・S・チューグ氏は言うが、この研究には関与していない。 「しかし、患者ケアの観点から見ると、このような発見を利用して一次予防植込み型除細動器の候補を特定する前に、さらに多くの研究を行う必要があります」と彼は付け加えた。
それでもオーバーマイヤー氏は、このアプローチは追求する価値があると考えている。 「MRIのような非常に高度な画像技術はいくつかありますが、これらは費用と不便さのため、集団のスクリーニングには現実的ではありません」とオーバーマイヤー氏は言う。 ECG はスペクトルの対極に位置すると彼は主張します。 Apple Watch やスマートフォンに接続する簡単なデバイスを使用して、ほぼどこでも記録できます。研究チームは、このモデルが医療グレードの心電図でトレーニングされており、消費者向けデバイスからの低品質の信号ではわずかにパフォーマンスが劣ることを認めていますが、わずかながらマイナーと説明しています。
「心電図のリスクが高いといわれたからといって、外出して除細動器の植え込みを受けることはお勧めしません」とオーバーマイヤー氏は言う。 「これの良いところは、AI をまったく信じる必要がないことです。従来のリスク マーカーを行う場合と同様に、追加のテストを対象とするために使用できます。

」
Jacek Krywko は、宇宙探査とコンピューター サイエンスをカバーするフリー ライターです。
科学のために立ち上がる時が来た
この記事が気に入っていただけましたら、サポートをお願いいたします。 Scientific American は 180 年にわたって科学と産業の擁護者としての役割を果たしてきましたが、今はその 2 世紀の歴史の中で最も重要な瞬間かもしれません。
私は 12 歳の頃からサイエンティフィック アメリカンを購読しており、それが私の世界に対する見方を形作るのに役立ちました。 SciAm はいつも私を教育し、楽しませてくれて、私たちの広大で美しい宇宙に対する畏敬の念を引き起こします。あなたにとってもそうなることを願っています。
Scientific American を購読すると、私たちの報道が有意義な研究と発見に集中するようになります。米国中の研究所を脅かす決定について報告するリソースがあること。そして、科学そのものの価値が認識されないことが多い現在、私たちは新進の科学者と現役の科学者の両方をサポートします。
その代わりに、重要なニュース、魅力的なポッドキャスト、素晴らしいインフォグラフィック、見逃せないニュースレター、必見のビデオ、やりがいのあるゲーム、科学界最高の執筆とレポートを入手できます。誰かに定期購入をプレゼントすることもできます。
私たちが立ち上がり、なぜ科学が重要なのかを示すことがこれほど重要な時期はありません。その使命において私たちをサポートしていただければ幸いです。
David M. Ewalt、サイエンティフィック・アメリカン編集長
Scientific American を購読して、今日の世界を形作る最もエキサイティングな発見、革新、アイデアを学び、共有してください。
Cookie の使用/データを販売しないでください
Scientific American は Springer Nature の一部であり、Springer Nature は数千の科学出版物を所有または商業関係を持っています (その多くは www.springernature.com/us でご覧いただけます)。 Scientific American は編集の独立という厳格なポリシーを維持しています

科学の発展を読者に報告するという目的で。
© 2026 サイエンティフィック アメリカン、スプリンガー ネイチャー アメリカ部門
すべての権利を留保します。

## Original Extract

A new model flags people at high risk of sudden cardiac death from a routine ECG—and reveals a warning sign in the heart’s electrical activity

Skip to main content Scientific American June 30, 2026
AI found a hidden heart-risk signal in 100-year-old ECG technology
A new model flags people at high risk of sudden cardiac death from a routine ECG—and reveals a warning sign in the heart’s electrical activity
By Jacek Krywko edited by Eric Sullivan
ECGs translate the heart’s electrical activity into a waveform of peaks and valleys. The new study used more than 440,000 such recordings to train and test an AI model.
Sudden cardiac death kills more than 300,000 people in the U.S. each year , even though implantable defibrillators have been able to stop many lethal arrhythmias for decades. The main issue today isn’t the device that stops a cardiac arrest; it is figuring out who needs one. In a new Nature study , a team led by Ziad Obermeyer , an associate professor at the University of California, Berkeley, trained a neural network to answer that question from a 10-second electrocardiogram. Then they trained a second neural network to reveal what the first was keying on.
The two-model setup points to a larger ambition for AI in medicine : getting a machine to surface a fresh clue that human experts can then see and check for themselves. Obermeyer’s team used the first network to predict risk and the second to translate that prediction into a visible feature on an ordinary ECG, one a cardiologist could learn to spot.
To decide who should get a defibrillator, cardiologists currently lean on an ultrasound measurement of how much blood the left ventricle pumps with each beat—a measure known as left ventricular ejection fraction , or LVEF. Obermeyer points out that it is far from perfect. “A lot of people who suddenly die of cardiac arrest either never had the ultrasound before or they had it and the results were normal,” he says. At the same time, most defibrillators implanted on the strength of that test never end up firing. “Often a person who looked high risk turned out not to be so high risk after all,” Obermeyer says. To get around the problem, his team went looking for a better risk marker .
On supporting science journalism
If you're enjoying this article, consider supporting our award-winning journalism by subscribing . By purchasing a subscription you are helping to ensure the future of impactful stories about the discoveries and ideas shaping our world today.
Electrocardiograms, or ECGs , measure the heart’s electrical activity and are cheap and nearly universal by comparison. Yet despite decades of studying ECG waveforms, cardiologists had never found a pattern that reliably flagged a high risk of cardiac arrest. His team turned to deep learning to find the pattern that human inspection had missed. The algorithm the team picked was a 64-layer residual neural network, or ResNet. “It’s kind of a workhorse model everyone uses. There’s nothing interesting about it,” Obermeyer says. “What is interesting is the data it’s learned from.”
To feed the network, Obermeyer’s group assembled one of the first population-scale datasets of its kind, with more than 440,000 ECGs from roughly 180,000 patients in Sweden, matched to national death certificates. Trained on the Swedish data, the otherwise generic ResNet flagged a high-risk group amounting to about 2.2 percent of patients. The signal held up when the team tested the model on separate datasets from the U.S. and Taiwan, suggesting this wasn’t a quirk of Sweden’s population or ECG equipment. Within that small group, the annual rate of sudden cardiac death reached 7 percent—well above the 4.6 percent rate among patients flagged by the standard ultrasound test. What’s more, more than 86 percent of the patients the algorithm singled out were not flagged by the traditional LVEF marker. By the traditional measure, many patients like these would have been sent home without a defibrillator.
“After we established this thing is working, we wanted to understand what this model is seeing in the ECG waveforms of high-risk people,” Obermeyer says. Standard AI interpretability tools like saliency maps can highlight which parts of a waveform a neural net weighted most heavily, but they stop there. A human cardiologist who spots something unusual on an ECG trace can sketch the anomalous wave. A neural network, by default, cannot. So, Obermeyer and his colleagues built a generative AI model to do just that. “Its job was to produce ECG waveforms that looked high-risk to the first model,” Obermeyer says.
Paired with the original network and guided by its risk score, the generative model reworked a real low-risk patient’s ECG step by step, morphing it smoothly into a high-risk version of the same trace. Many of the features the model keyed on were already familiar to cardiologists.
One feature, though, had never been described in the medical literature: a subtle slurring in one ECG lead called aVL, suggesting that the heart’s electrical signal was fragmenting as it moved through muscle.
Changxin Lai , a biomedical engineer at Johns Hopkins University who wrote an accompanying analysis in Nature and was not involved in the study, says this is why the work stands out. “The ECG has been around for more than 100 years, and this kind of data has been carefully evaluated by generations of cardiologists,” he says. “We extracted new knowledge from an artificial intelligence model.”
For some of the high-risk patients, the team also had cardiac magnetic resonance imaging, or MRI, scans. Those scans showed subtle, diffuse fibrosis, scarring associated with arrhythmias that can interfere with the heart’s electrical signals in a way that fits the synthetic waveforms the generative model produced. Obermeyer cautions that the fibrosis link is preliminary and has yet to be confirmed with biopsies.
The finding, while intriguing, is not ready to guide treatment. “This is an important area of research,” says Sumeet S. Chugh , who directs the Center for Cardiac Arrest Prevention at Cedars-Sinai Medical Center and was not involved in the study. “But from a patient care perspective there is much more research to be done before we will be using such findings to… identify candidates for the primary prevention implantable defibrillator,” he adds.
Even so, Obermeyer thinks the approach is worth pursuing. “There are some very fancy imaging techniques like MRI, but these things are not feasible for screening populations because of their expense and inconvenience,” Obermeyer says. ECGs, he argues, sit at the opposite end of the spectrum; they can be recorded nearly anywhere, with an Apple Watch or a simple device that connects to a smartphone. The team acknowledges that the model was trained on medical-grade ECGs and performs slightly worse on the lower-quality signals from consumer devices, though by a margin they describe as minor.
“I wouldn’t suggest going out and getting a defibrillator implanted just because we say your ECG is high risk,” Obermeyer says. “What’s nice about this is you don’t have to believe the AI at all. You can just use it to target additional testing like doing traditional risk markers.”
Jacek Krywko is a freelance writer who covers space exploration and computer science.
It’s Time to Stand Up for Science
If you enjoyed this article, I’d like to ask for your support. Scientific American has served as an advocate for science and industry for 180 years, and right now may be the most critical moment in that two-century history.
I’ve been a Scientific American subscriber since I was 12 years old, and it helped shape the way I look at the world. SciAm always educates and delights me, and inspires a sense of awe for our vast, beautiful universe. I hope it does that for you, too.
If you subscribe to Scientific American , you help ensure that our coverage is centered on meaningful research and discovery; that we have the resources to report on the decisions that threaten labs across the U.S.; and that we support both budding and working scientists at a time when the value of science itself too often goes unrecognized.
In return, you get essential news, captivating podcasts , brilliant infographics, can't-miss newsletters , must-watch videos, challenging games , and the science world's best writing and reporting. You can even gift someone a subscription .
There has never been a more important time for us to stand up and show why science matters. I hope you’ll support us in that mission.
David M. Ewalt, Editor in Chief, Scientific American
Subscribe to Scientific American to learn and share the most exciting discoveries, innovations and ideas shaping our world today.
Use of cookies/Do not sell my data
Scientific American is part of Springer Nature, which owns or has commercial relations with thousands of scientific publications (many of them can be found at www.springernature.com/us). Scientific American maintains a strict policy of editorial independence in reporting developments in science to our readers.
© 2026 SCIENTIFIC AMERICAN, A DIVISION OF SPRINGER NATURE AMERICA, INC.
ALL RIGHTS RESERVED.
