---
source: "https://en.wikipedia.org/wiki/Bitter_lesson"
hn_url: "https://news.ycombinator.com/item?id=49338480"
title: "The bitter lesson is the observation in artificial intelligence"
article_title: "Bitter lesson - Wikipedia"
image: ""
author: "doener"
captured_at: "2026-08-17T23:14:17Z"
capture_tool: "hn-digest"
hn_id: 49338480
score: 2
comments: 0
posted_at: "2026-08-17T22:22:28Z"
tags:
  - hacker-news
  - translated
---

# The bitter lesson is the observation in artificial intelligence

- HN: [49338480](https://news.ycombinator.com/item?id=49338480)
- Source: [en.wikipedia.org](https://en.wikipedia.org/wiki/Bitter_lesson)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T22:22:28Z

## Translation

タイトル: 苦い教訓は人工知能の観察です
記事タイトル: 苦い教訓 - Wikipedia

記事本文:
コンテンツへジャンプ
メインメニュー
メインメニュー
サイドバーに移動
隠す
ナビゲーション
メインページ
人工知能 (AI) に関するシリーズの一部
主な目標
汎用人工知能
超知性
ディープフェイクポルノ
テイラー・スウィフトのディープフェイクポルノ論争
Google Gemini の画像生成に関する論争
一年で最も恐ろしい時期です
OpenAIからサム・アルトマン氏を削除
Voiceverse NFT盗作スキャンダル
苦い教訓は、長期的には、利用可能な計算能力に応じて拡張する一般的なアプローチの方が、時間の経過とともに低下する計算コストをよりうまく利用できるため、ドメイン固有の理解に基づくアプローチよりも優れたパフォーマンスを発揮する傾向があるという人工知能の観察です。この原理は、Richard Sutton による 2019 年のエッセイ [1] で提案および命名され、現在では広く受け入れられています。 [ 2 ] [ 3 ] [ 4 ] [ 5 ] [ 6 ] [ 7 ] [ 8 ]
サットンは、この教訓を説明するいくつかの例を挙げています。
ゲームプレイ中。チェスでは、世界チャンピオンを破った最初のコンピュータ対戦相手となったディープ ブルー システムは、最善手を検索するために大量の特殊なハードウェアを適用することでスケールアップした、比較的単純なアルファ - ベータ検索アルゴリズムに依存していました。これにより、チェスの独特な構造を利用したり、グランドマスターの知識を直接組み込んだりする以前の試みが打ち破られました。同様に、囲碁ゲームでも、人間のパフォーマンスを超えた AlphaGo アルゴリズムは、前世代の AI よりもゲーム自体の専門スキルに依存することがはるかに少なく、人間の専門知識を完全に排除し、自己対局のみによって訓練された AlphaGo Zero によってさらに上回りました。
音声認識。多数の音声サンプルを使用した汎用の隠れマルコフ モデルのトレーニングに基づくアプローチは、1970 年代の手作りのアプローチを常に上回っていました。

ディープラーニングもこの傾向を引き継いでいます。
コンピュータビジョン。人間の視覚システムに近似すると想定されていたアルゴリズム (明示的にエンコードされたエッジ検出や SIFT による高レベルの特徴の検出など) は、視覚認識の性質についての仮定がはるかに少ない畳み込みニューラル ネットワークのパフォーマンスよりも優れていました。
サットン氏は、これまで以上に複雑な人間の洞察を導入するよりも、ムーアの法則を活用できるシンプルでスケーラブルなソリューションを見つけることに時間を投資する方が良いと結論付け、これを「苦い教訓」と呼んでいます。彼はまた、効果的に拡張できることが証明されている 2 つの汎用手法、検索と学習を挙げています。この教訓が「苦い」と考えられているのは、多くの研究者が予想していたほど人間中心主義的ではなく、そのため彼らがそれを受け入れるのが遅かったためである。
このエッセイは 2019 年にサットン氏の Web サイト incompleteideas.net で公開され、Google Scholar によると数百件の正式な引用を受けています。これらの中には、原則の別の記述を提供するものもあります。たとえば、Google DeepMind の 2022 年の論文「A Generalist Agent」では、教訓が次のように要約されています。 [ 2 ]
歴史的に、より優れた汎用モデル
コンピューティングの活用は、最終的にはより専門化されたドメイン固有のアプローチを追い越す傾向にもあります。
この原理の別の表現は、Noam Shazeer との共著によるスイッチ変圧器に関する Google の論文に見られます。 [3]
豊富な計算予算、データセットのサイズ、パラメータ数に裏打ちされたシンプルなアーキテクチャは、より複雑なアルゴリズムを上回ります。
この原理は、人工知能に関する他の多くの著作でもさらに参照されています。たとえば、『ディープラーニングから合理的マシンまで』では、モラベックのパラドックスやニートとだらしない人の対比など、この分野で長年議論されてきた議論とのつながりが描かれています。

。 [9] 「人工知能の少ないエンジニアリング」の中で、著者らは「これまでの柔軟な手法は長期的には常に手作りのドメイン知識を上回ってきた」ことに同意しているが、「適切な（暗黙の）仮定がなければ一般化は不可能である」と述べている。 [5] より最近では、「脳の苦い教訓: 自己教師あり学習による音声デコーディングの拡張」でサットンの議論が続き、音声認識と脳データの分野では教訓が（2025年現在）完全に学習されていないと主張している。 [ 6 ]
他の研究では、この原則を適用し、新しい領域でそれを検証することが検討されています。たとえば、2022年の論文「Beyond the Imitation Game」では、この原則を大規模な言語モデルに適用し、「スケールだけで解決できそうな問題に研究リソースを投入することを避ける」ために「その機能と限界を理解することが非常に重要である」と結論付けている。 [7] 2024 年、「苦い教訓を学ぶ: 20 年間の CVPR 訴訟からの経験的証拠」では、コンピュータ ビジョンとパターン認識の分野からのさらなる証拠を検討し、この分野における過去 20 年間の経験が「CVPR 訴訟の強力な遵守を示している」と結論付けています。
「苦い教訓」の中核となる原則。 [ 4 ] 「アクター批評家における過大評価、過適合、可塑性: 強化学習の苦い教訓」の中で、著者らはアクター批評家アルゴリズムの一般化に着目し、「勾配ベースの学習の安定化を動機とする一般的な手法は、さまざまな環境にわたる RL 固有のアルゴリズムの改善よりも大幅に優れている」ことを発見し、これが苦い教訓と一致していると指摘しています。 [ 8 ]
↑ リッチ州サットン（2019年3月13日）。 「苦いレッスン」。 www.incompleteideas.net 。 2025 年 9 月 7 日に取得。
1 2 リード、スコット;ゾ

lna、コンラッド。パリゾット、エミリオ。他。 （2022年）。 「ゼネラリストエージェント」。機械学習研究に関するトランザクション (2834–8856)。 arXiv: 2205.06175 。 2025 年 9 月 7 日に取得。
1 2 フェドゥス、ウィリアム。ゾフ、バレット。ノーム・シャジーア（2022）。 「スイッチ トランスフォーマー: シンプルかつ効率的なスパース性による数兆パラメータ モデルへのスケーリング」。機械学習研究ジャーナル。 23 (120): 1–39 。 2025 年 9 月 14 日に取得。
1 2 ユセフィ、モジタバ;コリンズ、ジャック。 「苦い教訓を学ぶ: 20 年間の CVPR 訴訟からの経験的証拠」。 NLP for Science (NLP4Science) に関する第 1 回ワークショップの議事録。計算言語学協会。 175–187ページ。 2025 年 9 月 7 日に取得。
1 2 シンツ、ファビアン H.ピトコウ、Xaq;ライマー、ジェイコブ。他。 （2019年）。 「人工知能の少ないエンジニアリング」。ニューロン。 103（6）。エルゼビア: 967–979。土井: 10.1016/j.neuron.2019.08.034。 2025 年 9 月 13 日に取得。
1 2 ジャヤラス、ダルハン;ランダウ、ギラッド。ブレンダン・シリングフォード。ウールリッチ、マーク。パーカー・ジョーンズ、「オイウィ」（2025）。 「脳の苦い教訓: 自己教師あり学習による音声解読の拡張」機械学習に関する第 42 回国際会議。機械学習研究の議事録。 2025 年 9 月 13 日に取得。
1 2 スリバスタヴァ、アーロヒ。ラストギ、アビナフ。ラオ、アビシェク。アワル、アブ。アビド、アブバカール。他。 「イミテーション ゲームを超えて: 言語モデルの機能の定量化と推定」。学習表現に関する第 14 回国際会議。
1 2 ナウマン、ミハル。ボルトキェヴィチ、ミハウ;ミウォシュ、ピョートル。トシンスキー、トマシュ。オスタシェフスキ、マテウシュ。他。 （2024年）。 「俳優と批評家における過大評価、過適合、可塑性: 強化学習の苦い教訓」第 41 回機械学習国際会議の議事録。 P

機械学習研究の議事録。 2025 年 9 月 13 日に取得。
↑ バックナー、キャメロン J. (2023 年 12 月 11 日)。ディープラーニングから合理的マシンまで: 人工知能の未来について哲学の歴史が教えてくれること 。オックスフォード大学出版局。土井：10.1093/oso/9780197653302.001.0001 。 ISBN 9780197653302 。
「 https://en.wikipedia.org/w/index.php?title=Bitter_lesson&oldid=1343753531 」から取得
カテゴリ : 人工知能の哲学
短い説明のある記事
短い説明がウィキデータと一致します
2025 年 10 月からの mdy 日付を使用します
このページは、2026 年 3 月 16 日の 04:40 (UTC) に最後に編集されました。
ページは Parsoid でレンダリングされました。
テキストは、クリエイティブ コモンズ 表示 - 継承 4.0 ライセンスに基づいて利用できます。
追加の条件が適用される場合があります。このサイトを使用すると、利用規約とプライバシー ポリシーに同意したことになります。 Wikipedia® は、非営利団体である Wikimedia Foundation, Inc. の登録商標です。

## Original Extract

Jump to content
Main menu
Main menu
move to sidebar
hide
Navigation
Main page
Part of a series on Artificial intelligence (AI)
Major goals
Artificial general intelligence
superintelligence
Deepfake pornography
Taylor Swift deepfake pornography controversy
Google Gemini image generation controversy
It's the Most Terrible Time of the Year
Removal of Sam Altman from OpenAI
Voiceverse NFT plagiarism scandal
The bitter lesson is the observation in artificial intelligence that, in the long run, general approaches that scale with available computational power tend to outperform ones based on domain-specific understanding because they are better at taking advantage of the falling cost of computation over time. The principle was proposed and named in a 2019 essay by Richard Sutton [ 1 ] and is now widely accepted. [ 2 ] [ 3 ] [ 4 ] [ 5 ] [ 6 ] [ 7 ] [ 8 ]
Sutton gives several examples that illustrate the lesson:
Game playing . In chess , the Deep Blue system that became the first computer opponent to defeat a world champion relied on a relatively simple alpha–beta search algorithm that scaled up by applying large amounts of specialized hardware to search for the best move. This defeated previous attempts to exploit the unique structure of chess or to include grandmaster knowledge directly. Likewise in the game of Go , the AlphaGo algorithm that surpassed human performance relied much less on expert skill at the game itself than previous generations of AI, and was further surpassed by AlphaGo Zero , which removed human expertise completely and trained only by self-play .
Speech recognition . Approaches based on training a general-purpose hidden Markov model with large numbers of speech samples consistently outperformed the hand-crafted approaches of the 1970s, and deep learning has continued this trend.
Computer vision . Algorithms that were assumed to approximate the human visual system (such as explicitly encoded edge detection or detecting high-level features with SIFT ) were outperformed by convolutional neural networks that make far fewer assumptions about the nature of visual perception .
Sutton concludes that time is better invested in finding simple scalable solutions that can take advantage of Moore's law, rather than introducing ever-more-complex human insights, and calls this the "bitter lesson". He also cites two general-purpose techniques that have been shown to scale effectively: search and learning . The lesson is considered "bitter" because it is less anthropocentric than many researchers expected and so they have been slow to accept it.
The essay was published on Sutton's website incompleteideas.net in 2019, and has received hundreds of formal citations according to Google Scholar . Some of these provide alternative statements of the principle; for example, the 2022 paper "A Generalist Agent" from Google DeepMind summarized the lesson as: [ 2 ]
Historically, generic models that are better at
leveraging computation have also tended to overtake more specialized domain-specific approaches, eventually.
Another phrasing of the principle is seen in a Google paper on switch transformers coauthored by Noam Shazeer : [ 3 ]
Simple architectures—backed by a generous computational budget, data set size and parameter count—surpass more complicated algorithms.
The principle is further referenced in many other works on artificial intelligence. For example, From Deep Learning to Rational Machines draws a connection to long-standing debates in the field, such as Moravec's paradox and the contrast between neats and scruffies . [ 9 ] In "Engineering a Less Artificial Intelligence", the authors concur that "flexible methods so far have always outperformed handcrafted domain knowledge in the long run" although note that "[w]ithout the right (implicit) assumptions, generalization is impossible". [ 5 ] More recently, "The Brain's Bitter Lesson: Scaling Speech Decoding With Self-Supervised Learning" continues Sutton's argument, contending that (as of 2025) the lesson has not been fully learned in the fields of speech recognition and brain data . [ 6 ]
Other work has looked to apply the principle and validate it in new domains. For example, the 2022 paper "Beyond the Imitation Game" applies the principle to large language models to conclude that "it is vitally important that we understand their capabilities and limitations" to "avoid devoting research resources to problems that are likely to be solved by scale alone". [ 7 ] In 2024, "Learning the Bitter Lesson: Empirical Evidence from 20 Years of CVPR Proceedings" looked at further evidence from the field of computer vision and pattern recognition , and concludes that the previous twenty years of experience in the field shows "a strong adherence to
the core principles of the 'bitter lesson'". [ 4 ] In "Overestimation, Overfitting, and Plasticity in Actor-Critic: the Bitter Lesson of Reinforcement Learning", the authors look at generalization of actor-critic algorithms and find that "general methods that are motivated by stabilization of gradient-based learning significantly outperform RL -specific algorithmic improvements across a variety of environments" and note that this is consistent with the bitter lesson. [ 8 ]
↑ Sutton, Rich (March 13, 2019). "The Bitter Lesson" . www.incompleteideas.net . Retrieved September 7, 2025 .
1 2 Reed, Scott; Zolna, Konrad; Parisotto, Emilio; et al. (2022). "A Generalist Agent" . Transactions on Machine Learning Research ( 2834– 8856). arXiv : 2205.06175 . Retrieved September 7, 2025 .
1 2 Fedus, William; Zoph, Barret; Shazeer, Noam (2022). "Switch Transformers: Scaling to Trillion Parameter Models with Simple and Efficient Sparsity" . Journal of Machine Learning Research . 23 (120): 1– 39 . Retrieved September 14, 2025 .
1 2 Yousefi, Mojtaba; Collins, Jack. "Learning the Bitter Lesson: Empirical Evidence from 20 Years of CVPR Proceedings" . Proceedings of the 1st Workshop on NLP for Science (NLP4Science) . Association for Computational Linguistics. pp. 175– 187 . Retrieved September 7, 2025 .
1 2 Sinz, Fabian H.; Pitkow, Xaq; Reimer, Jacob; et al. (2019). "Engineering a Less Artificial Intelligence" . Neuron . 103 (6). Elsevier: 967– 979. doi : 10.1016/j.neuron.2019.08.034 . Retrieved September 13, 2025 .
1 2 Jayalath, Dulhan; Landau, Gilad; Shillingford, Brendan; Woolrich, Mark; Parker Jones, ʻŌiwi (2025). "The Brain's Bitter Lesson: Scaling Speech Decoding With Self-Supervised Learning" . Forty-second International Conference on Machine Learning . Proceedings of Machine Learning Research . Retrieved September 13, 2025 .
1 2 Srivastava, Aarohi; Rastogi, Abhinav; Rao, Abhishek; Awal, Abu; Abid, Abubakar; et al. "Beyond the Imitation Game: Quantifying and extrapolating the capabilities of language models" . The Fourteenth International Conference on Learning Representations .
1 2 Nauman, Michal; Bortkiewicz, Michał; Miłoś, Piotr; Trzciński, Tomasz; Ostaszewski, Mateusz; et al. (2024). "Overestimation, Overfitting, and Plasticity in Actor-Critic: the Bitter Lesson of Reinforcement Learning" . Proceedings of the 41st International Conference on Machine Learning . Proceedings of Machine Learning Research . Retrieved September 13, 2025 .
↑ Buckner, Cameron J. (December 11, 2023). From Deep Learning to Rational Machines: What the History of Philosophy Can Teach Us about the Future of Artificial Intelligence . Oxford University Press. doi : 10.1093/oso/9780197653302.001.0001 . ISBN 9780197653302 .
Retrieved from " https://en.wikipedia.org/w/index.php?title=Bitter_lesson&oldid=1343753531 "
Category : Philosophy of artificial intelligence
Articles with short description
Short description matches Wikidata
Use mdy dates from October 2025
This page was last edited on 16 March 2026, at 04:40 (UTC) .
Page was rendered with Parsoid .
Text is available under the Creative Commons Attribution-ShareAlike 4.0 License ;
additional terms may apply. By using this site, you agree to the Terms of Use and Privacy Policy . Wikipedia® is a registered trademark of the Wikimedia Foundation, Inc. , a non-profit organization.
