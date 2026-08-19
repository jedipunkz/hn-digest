---
source: "https://austinhenley.com/blog/dearresearchers.html"
hn_url: "https://news.ycombinator.com/item?id=49368150"
title: "Dear researchers: Is AI all you've got?"
article_title: "Dear researchers: Is AI all you've got? - Austin Z. Henley"
image: "https://austinhenley.com/blog/images/dearresearchers.png"
author: "jermaustin1"
captured_at: "2026-08-19T23:14:26Z"
capture_tool: "hn-digest"
hn_id: 49368150
score: 1
comments: 0
posted_at: "2026-08-19T22:47:04Z"
tags:
  - hacker-news
  - translated
---

# Dear researchers: Is AI all you've got?

- HN: [49368150](https://news.ycombinator.com/item?id=49368150)
- Source: [austinhenley.com](https://austinhenley.com/blog/dearresearchers.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T22:47:04Z

## Translation

タイトル: 研究者の皆様: あなたが持っているのは AI だけですか?
記事のタイトル: 研究者の皆様: あなたが持っているのは AI だけですか? - オースティン・Z・ヘンリー
説明: AI への過度の執着のせいで、私たちは次の大きなイノベーションを見逃していませんか?

記事本文:
研究者の皆様へ: あなたが持っているのは AI だけですか?
これは、Journal of Systems and Software (JSS) の研究者の皆様へのコラムに今後掲載される私の記事の草稿です。私はこのコラムの共同編集者であり、技術移転における実際的な課題について研究者に向けた記事を寄稿するよう業界の専門家を募集しています。興味があればメールしてください！
ソフトウェア エンジニアリングの研究者が AI に集中するために放棄した重要な問題は何ですか?
1900 年代初頭、バクテリオファージ療法は、特定の細菌を標的にする可能性があるため、医学および生物学の研究で人気のトピックでした 1 。しかし、1928 年のペニシリンの発見後、研究コミュニティは抗生物質に注目を移し、1950 年代までにバクテリオファージの研究はほとんど放棄されました 2 。抗生物質耐性がより一般的になり、ファージ療法の高度に標的を絞った性質がより魅力的になったため、2000 年代に再び関心が高まりました 3 。
同様の現象がソフトウェア エンジニアリング コミュニティでも起こっていますか?研究者は AI に関するインデックスを過剰に作成していますか? ICSE '25 では、研究分野の論文の 3 分の 1 と産業分野の論文の 3 分の 2 が AI に関係していました。対照的に、ソフトウェア エンジニアリングにおける 1992 年から 2016 年までの研究傾向を調査したところ、最も人気のあるトピックのトップ 10 に AI や ML は含まれていませんでした 4 。さらに、AI カンファレンスは急激に成長しており、AAAI は 2022 年の 9,000 件の申請から、2026 年のカンファレンスには 29,000 件の申請が予定されています。
私の懸念は、AI に大きな価値がない (私は学界と産業界の両方で AI に取り組んできました) とか、AI がこの分野にとって大きなイノベーションではないということではなく、むしろ他のトピックも重要であるということです。私たちが見逃す可能性があるイノベーションは何でしょうか?今後もどのような問題が解決されないでしょうか? AI を超えた日常的な問題点は何ですか?
ソフトウェアエンジンとして

高度に分散されたシステムを構築する場合、アーキテクチャと実装の決定がシステムの復元力において重要な役割を果たします。 2025 年には、Google Cloud 、AWS 、Azure 、Cloudflare でいくつかの重大な停止が発生し、再び Cloudflare が発生しました。インターネット全体がダウンしたように感じました。これらの障害は、Spotify、Discord、ChatGPT、Zoom、Venmo、Reddit、Amazon、LinkedIn、Shopify、Fortnite、Square、銀行、航空会社、さらには実店舗の小売店にまで影響を及ぼしました。
全ウェブサイトの約 20% が Cloudflare を使用しており、AWS、Azure、GCP が世界のクラウド インフラストラクチャの 62% を占めていることを考えると、これは大きな懸念事項です。 AWS の停止では、わずか 15 時間で企業に 5 億 8,100 万ドル以上の損失が生じたと推定され、Cloudflare の停止の 1 つでは 4 時間以内で 3 億ドル以上の損失が発生したと推定されています。実際、私は、これは地球規模の大惨事が起こるのを待っていると主張しており、警告の兆候は繰り返し見られていますが、研究者たちがあたかも命がかかっているかのようにこの問題の解決に群がっているようには見えません。 AI はソフトウェアの信頼性を劇的に高めましたか?
業界での私の経験から言えば、最近では買収を行った新興企業で私たちが直面した問題は根本的に人間の問題でした。たとえば、ビジネス パートナー、顧客、弁護士に、当社のソフトウェアは当社の言うとおりの動作をし、大きな混乱は生じないことを説得することは、当社がかなりの時間を費やした大きな課題でした。 AI により、より迅速に拡張できるようになりましたが、他の人にシステムを信頼してもらうのには役立ちませんでした。
2023 年後半の ChatGPT リリースに向けた ICSE と FSE の基調講演を見渡すと、ソフトウェアの安全性、信頼性、ランタイム監視、テスト、業界への影響、ソフトウェア エンジニアリングの教育など、かつて研究者が最優先に考えていたトピックを思い出すことができます。

これらの講演のいくつかは、FSE 2022 でのソフトウェアの安全性と信頼性に関する Marsha Chechik の基調講演など、私たちが経験している世界的なクラウド障害に関連しています。残念ながら、彼女の行動喚起は十分ではなかったようです。
しかし、もしかしたら私たちが思いもよらないイノベーションが世の中にあるかもしれません。クレイトン・クリステンセンは、独創的な著書『イノベーションのジレンマ 5』の中で、分野が成熟するにつれて、既存のイノベーションを最適化して維持する一方で、斬新で破壊的なイノベーションを見逃してしまうことが多いと主張しています。同氏は続けて、混乱は他者が小さすぎる、未熟すぎる、または直交しすぎるとして無視するような場所から始まることが多く、AIがあらゆる研究課題に対するデフォルトの答えになると、まさに無視される危険がある領域であると述べた。
なぜ研究者が AI の流行に飛びつくのかを簡単に説明すると、インセンティブにあります。 AI 関連の提案を求める資金提供機関、論文募集に AI に関する複数のトピックを追加する会議、AI 研究者に多数の教員募集を掲載する大学​​など、あらゆる人を AI に引き込む信じられないほど強い力が存在しています。研究者にとってそれを避けるのは非常に困難です。実際、私が AI に関連して共著した論文は、私の他のどの論文よりもはるかに高い割合で引用されています。私の共編者のオラフ・ジンマーマンが観察しているように、学問的行動は内発的動機よりも、学問の評価方法を決定するインセンティブ構造や指標によって形作られることが多い 6 。言い換えれば、「インセンティブ、インセンティブ、そしてインセンティブ」です。
AI への執着について懸念を表明したのは私が初めてではありません。実際、他の何人かが AI について議論しています。

AI の潜在的なマイナス面。ジョンソン氏とメンジーズ氏は、「AI の過剰な誇大宣伝」は危険であり、「そのような発言に対して結集するのはソフトウェア専門家の倫理的義務」であると述べた 7 。教室での AI の使用に関する初期の研究では、AI が学習を妨害する可能性があるという証拠が示されています 8、9、10 。この状況は、AI によってソフトウェア エンジニア、特にキャリアの初期段階にあるエンジニアの雇用市場が混乱していることによってさらに複雑になっています。
事実上すべての研究者が AI に焦点を当てていることを考えると、今が他のトピックで進歩を遂げるチャンスです。ディープ ニューラル ネットワークに関する基礎的な研究で 2018 年チューリング賞を受賞したヤン ルカン氏でさえ、LLM は今後も拡張されず、「LLM では機能しない」という自身の信念に関して、「LLM は基本的に出口、気晴らし、行き止まりである」と述べています。 。同氏はまた、「現在、彼らはどこへ行っても部屋の空気を吸い取っており、基本的に他のことに使えるリソースはない。そのため、次の革命に向けて、我々は一歩下がって、現在のアプローチに何が欠けているのかを把握する必要がある」とも警告した。
おそらく、AI の過剰なインデックス作成に関するこの懸念を解決するためのインスピレーションとして機械学習を使用できるかもしれません。機械学習アルゴリズムには、極大値を回避するメカニズムがあります。シミュレーテッドアニーリングと確率的勾配降下法はノイズを導入し、進化的アルゴリズムはランダムな開始点と突然変異によるバリエーションを作成し、最適化手法はランダムな摂動と再スタートを使用して探索空間の新しい領域を探索します。
研究者はこれらの概念を次の研究テーマの決定に適用できますか?研究者たちが極大値に陥っているのではないかと心配しています。さらに悪いことに、彼らは AI の時流に参加する (または乗り続ける) ことを目的として、研究トピックの探求における「ランダム性」を意図的に減らしています。
うーん

誰もが AI を研究している一方で、私たちの社会を静かに運営しているソフトウェア システムは、珍しいことでもなく、驚くべきことでもなく失敗し続けています。世界的なクラウドが繰り返しダウンし、銀行、航空会社、小売業者、政府に影響を与える可能性がある場合、回復力と信頼の問題は解決されません。残念な点は、これらの問題が難しいということではなく、むしろ時代遅れに感じられるようになっていることです。
2021年の著書『Think Again 11』の中で、組織心理学者のアダム・グラントは、「科学者のように考えるということは、単に広い心で反応するだけではありません。それは、積極的に心を開くことを意味します。私たちが正しくなければならない理由ではなく、間違っている可能性がある理由を探し、学んだことに基づいて自分の意見を修正することが必要です。」と主張しました。この科学的アプローチによる場合でも、ランダムな摂動による場合でも、極大値に囚われることはありません。
ソフトウェアとシステムの研究者の皆さん、あなたが持っているのは AI だけですか?
この記事を改善するために何度もフィードバックを提供してくださった Olaf Zimmermann 氏、そして私と『Dear Researchers』の共同編集者となってくれたことに特に感謝します。
研究コミュニティと共有したい意見がある場合は、研究者の皆様への記事を書くことを検討してください。私に連絡してください。
サモンド、GP、フィネラン、P.C. (2015)。ファージの 1 世紀: 過去、現在、未来。 Nature Reviews Microbiology、13(12)、777-786。
Wittebole, X.、De Roock, S.、Opal, S.M. (2014)。細菌性病原体の治療における抗生物質の代替としてのバクテリオファージ療法の歴史的概要。毒性、5(1)、226-235。 https://doi.org/10.4161/viru.25991
ゴルディロ・アルタミラノ、F.L.、バール、J.J. (2019)ポスト抗生物質時代のファージ療法。臨床微生物学総説、32(2)、10-1128。
マシュー、G.、アグラワル、A.、メンジーズ、T. (2018)。トレンドを見つける

ソフトウェアの研究。ソフトウェアエンジニアリングに関するIEEEトランザクション、49(4)、1397-1410。
クリステンセン、C.M. (2015)。革新者のジレンマ: 新しいテクノロジーが偉大な企業の失敗を引き起こすとき。ハーバード・ビジネス・レビュー・プレス。
ジマーマン、O. (2025)。研究と実践のギャップを克服する: 根本原因分析と、ソフトウェア アーキテクチャおよび分散システムにおける実際に関連するトピック。システムとソフトウェアジャーナル、230。
ジョンソン、B.、メンジーズ、T. (2024)。 Ai の過剰な誇大宣伝: 危険な脅威 (およびその修正方法) 。 IEEE ソフトウェア、41(6)、131-138。
Kazemitabaar, M.、Williams, J.、Drosos, I.、Grossman, T.、Henley, A. Z.、Negreanu, C.、Sarkar, A. (2024 年 10 月)。インタラクティブなタスク分解による AI 支援データ分析におけるステアリングと検証の改善。ユーザー インターフェイス ソフトウェアとテクノロジに関する第 37 回 ACM 年次シンポジウムの議事録 (1-19 ページ)。
Bastani, H.、Bastani, O.、Sungu, A.、Ge, H.、Kabakcı, Ö.、Mariman, R. (2024)。生成型 AI は学習に害を及ぼす可能性があります。ウォートンスクールの研究論文。
Prather, J.、Reeves, B. N.、Leinonen, J.、MacNeil, S.、Randrianasolo, A. S.、Becker, B. A.、... & Briggs, B. (2024 年 8 月)。広がる格差: 初心者プログラマーにとっての生成 AI の利点と害。国際コンピューティング教育研究に関する 2024 年 ACM 会議議事録、第 1 巻 (pp. 469-486)。
グラント、A. (2023)。もう一度考えてみましょう。自分が知らないことを知る力。ペンギン。

## Original Extract

Are we missing the next big innovation because of the over-fixation on AI?

Dear researchers: Is AI all you've got?
This is a draft of my upcoming article in the Dear Researchers column in the Journal of Systems and Software (JSS). I'm a co-editor of the column, and we invite industry practitioners to contribute articles addressed to researchers about practical challenges in technology transfer. Email me if you're interested!
What important problems have software engineering researchers abandoned so that they could focus on AI instead?
In the early 1900s, bacteriophage therapy was a popular topic in medical and biological research for its potential to target specific bacteria 1 . However, after the discovery of penicillin in 1928, the research community shifted its attention to antibiotics, mostly abandoning bacteriophage research by the 1950s 2 . Interest surged again in the 2000s as antibiotic resistance grew more common and the highly targeted nature of phage therapy became more appealing 3 .
Is a similar phenomenon happening in the software engineering community? Are researchers over-indexing on AI? At ICSE '25 , one-third of research-track papers and two-thirds of industry-track papers involved AI. In contrast, a study of research trends from 1992 to 2016 in software engineering did not even include AI or ML in the top-10 most popular topics 4 . Furthermore, AI conferences have grown exponentially, with AAAI going from 9000 submissions in 2022 to 29,000 submissions for the 2026 conference.
My concern is not that AI doesn't have considerable value (I have worked on AI in both academia and industry) or that it isn't a huge innovation for the field, but rather that other topics are important too. What innovations might we miss? What problems will continue to go unsolved? What daily pain points exist beyond AI?
As software engineers build highly distributed systems, architectural and implementation decisions play a critical role in system resilience. In 2025, we saw several significant outages from Google Cloud , AWS , Azure , Cloudflare , and Cloudflare again . It felt like the entire internet was down! These outages impacted Spotify, Discord, ChatGPT, Zoom, Venmo, Reddit, Amazon, LinkedIn, Shopify, Fortnite, Square, banks, airlines, and even brick-and-mortar retailers.
Given that approximately 20% of all websites use Cloudflare and that AWS, Azure, and GCP account for 62% of the global cloud infrastructure, this is of huge concern. The AWS outage was estimated to have cost businesses upwards of $581 million dollars in just 15 hours and one of the Cloudflare outages to have cost $300 million dollars in less than 4 hours. In fact, I argue that it is a global catastrophe waiting to happen and we have seen repeated warning signs, yet I do not see researchers flocking to solve the problem as if lives depend on it. Has AI dramatically increased the reliability of software?
From my experience in industry, most recently at a startup that went through an acquisition, the problems we faced were fundamentally human problems. For example, convincing business partners, customers, and lawyers that our software does what we say that it does and that it won't have major disruptions was a huge challenge that we spent considerable time on. AI allowed us to scale faster, but it did not help us convince others to trust our system.
Looking over the keynote talks from ICSE and FSE leading up to the release of ChatGPT in late 2023, we can remind ourselves of the topics that were once top of mind for researchers: software safety, reliability, runtime monitoring, testing, industry impact, software engineering education, research rigor and reproducibility, socio-technical coordination, ethics and privacy, environmental impacts, etc. Several of these talks are relevant to the global cloud outages we have been experiencing, such as Marsha Chechik's keynote at FSE 2022 on the safety and reliability of software. Unfortunately, it seems that her call to action was not enough.
But perhaps there are innovations out there that we aren't even thinking about. Clayton Christensen, in his seminal book, Innovator's Dilemma 5 , argues that as fields mature, they often optimize and sustain existing innovations while also overlooking novel, disruptive innovations. He goes on to state that disruption often begins in places that others dismiss as too small, too immature, or too orthogonal, which are the exact areas that risk being neglected when AI becomes the default answer to every research question.
A straightforward explanation for why researchers would jump on the AI bandwagon lies in incentives. There has been an incredibly strong force pulling everyone into AI, including funding agencies asking for AI-related proposals, conferences adding multiple topics in AI to their calls for papers, and universities posting numerous faculty openings for AI researchers. It makes it incredibly difficult for a researcher to avoid. In fact, the papers I co-author involving AI are being cited at a far, far higher rate than any of my other work. As my co-editor, Olaf Zimmermann, observes, academic behavior is often shaped less by intrinsic motivation than by the incentive structures and metrics that govern how academics are evaluated 6 . In other words, "incentives, incentives, and incentives".
I'm certainly not the first to voice concerns about the fixation on AI, in fact, several others have argued about the potential downsides of AI. Johnson and Menzies stated that "AI overhype" is dangerous and that it is "the ethical duty of software professionals to rally against such remarks" 7 . Early studies of AI use in the classroom have shown evidence that it may disrupt learning 8, 9, 10 . This is further complicated by the disruption in the job market for software engineers, especially those that are early in their careers, caused by AI.
Given that virtually every researcher is focused on AI, now is the opportunity to make progress on any other topic. Even Yann LeCun , the 2018 Turing Award winner for his foundational work on deep neural networks, said that "an LLM is basically an off-ramp, a distraction, a dead end" , in regards to his belief that LLMs will not continue to scale, and "don't work on LLMs." . He also warned , "right now, they are sucking the air out of the room anywhere they go, and so there's basically no resources for anything else. And so for the next revolution, we need to take a step back and figure out what's missing from the current approaches" .
Perhaps we can use machine learning as inspiration to solve this concern of over-indexing on AI. Machine learning algorithms have mechanisms for escaping local maxima. Simulated annealing and stochastic gradient descent introduce noise, evolutionary algorithms create variation with random starting points and mutations, and optimization methods use random perturbations and restarts to explore new areas of a search space.
Can researchers apply these concepts to how they decide their next research topics? I worry that researchers are stuck in a local maximum. Worse yet, they are intentionally reducing the "randomness" in their exploration of research topics with the goal of joining (or staying on) the AI bandwagon.
While everyone is researching AI, the software systems that quietly run our society continue to fail in ways that are neither rare nor surprising. If the global cloud can go down repeatedly, affecting banks, airlines, retailers, and governments, then resilience and trust are not solved problems. The unfortunate part is not that these problems are hard, but rather that they increasingly feel unfashionable.
In the 2021 book Think Again 11 , organizational psychologist Adam Grant claimed, "Thinking like a scientist involves more than just reacting with an open mind. It means being actively open-minded. It requires searching for reasons why we might be wrong, not for reasons why we must be right, and revising our views based on what we learn." Whether it is through this scientific approach or through random perturbations, we can't get stuck in local maxima.
So software and systems researchers, is AI all you've got?
Special thanks to Olaf Zimmermann for providing multiple rounds of feedback to improve this article, and for being the co-editor of Dear Researchers with me.
If you have opinions that you want to share with the research community, consider writing an article for Dear Researchers . Reach out to me .
Salmond, G. P., & Fineran, P. C. (2015). A century of the phage: past, present and future . Nature Reviews Microbiology, 13(12), 777-786.
Wittebole, X., De Roock, S., & Opal, S. M. (2014). A historical overview of bacteriophage therapy as an alternative to antibiotics for the treatment of bacterial pathogens . Virulence, 5(1), 226-235. https://doi.org/10.4161/viru.25991
Gordillo Altamirano, F. L., & Barr, J. J. (2019). Phage therapy in the postantibiotic era . Clinical microbiology reviews, 32(2), 10-1128.
Mathew, G., Agrawal, A., & Menzies, T. (2018). Finding trends in software research . IEEE Transactions on Software Engineering, 49(4), 1397-1410.
Christensen, C. M. (2015). The innovator's dilemma: when new technologies cause great firms to fail . Harvard Business Review Press.
Zimmermann, O. (2025). Overcoming the research-practice gap: Root cause analysis and topics of practical relevance in software architecture and distributed systems . Journal of Systems and Software, 230.
Johnson, B., & Menzies, T. (2024). Ai over-hype: A dangerous threat (and how to fix it) . IEEE Software, 41(6), 131-138.
Kazemitabaar, M., Williams, J., Drosos, I., Grossman, T., Henley, A. Z., Negreanu, C., & Sarkar, A. (2024, October). Improving steering and verification in AI-assisted data analysis with interactive task decomposition . In Proceedings of the 37th Annual ACM Symposium on User Interface Software and Technology (pp. 1-19).
Bastani, H., Bastani, O., Sungu, A., Ge, H., Kabakcı, Ö., & Mariman, R. (2024). Generative AI can harm learning. The Wharton School Research Paper .
Prather, J., Reeves, B. N., Leinonen, J., MacNeil, S., Randrianasolo, A. S., Becker, B. A., ... & Briggs, B. (2024, August). The widening gap: The benefits and harms of generative ai for novice programmers . In Proceedings of the 2024 ACM Conference on International Computing Education Research-Volume 1 (pp. 469-486).
Grant, A. (2023). Think again: The power of knowing what you don't know . Penguin.
