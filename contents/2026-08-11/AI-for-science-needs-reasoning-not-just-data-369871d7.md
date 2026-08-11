---
source: "https://www.technologyreview.com/2026/08/10/1141384/ai-agents-for-science/"
hn_url: "https://news.ycombinator.com/item?id=49256255"
title: "AI for science needs reasoning, not just data"
article_title: "AI for science needs reasoning, not just data | MIT Technology Review"
author: "joozio"
captured_at: "2026-08-11T11:38:11Z"
capture_tool: "hn-digest"
hn_id: 49256255
score: 2
comments: 0
posted_at: "2026-08-11T11:02:25Z"
tags:
  - hacker-news
  - translated
---

# AI for science needs reasoning, not just data

- HN: [49256255](https://news.ycombinator.com/item?id=49256255)
- Source: [www.technologyreview.com](https://www.technologyreview.com/2026/08/10/1141384/ai-agents-for-science/)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T11:02:25Z

## Translation

タイトル: 科学用 AI にはデータだけでなく推論が必要
記事のタイトル: 科学のための AI にはデータだけでなく推論が必要 | MITテクノロジーレビュー
説明: 数十年ごとに、誰かが科学が終焉を迎えたと発表します。 1903 年、尊敬される物理学者アルバート マイケルソンは「物理科学の事実はすべて発見された」と書きました。 1980 年代、スティーブン ホーキング博士は、理論物理学は今世紀末までに完成するかもしれないと予測しました。ウィ
[切り捨てられた]

記事本文:
コンテンツへスキップ MIT テクノロジー レビュー
MIT テクノロジー レビューの特集
Opinion 科学用 AI にはデータだけでなく推論も必要
人間の研究プロセスをモデル化できる AI エージェントは、科学における発見を加速します。
数十年ごとに、誰かが科学が終焉を迎えたと発表します。 1903 年、尊敬される物理学者アルバート マイケルソンは「物理科学の事実はすべて発見された」と書きました。 1980 年代、スティーブン ホーキング博士は、理論物理学は今世紀末までに完成するかもしれないと予測しました。人工知能の爆発的な到来により、その感情が再び漂い、今度はノーベル賞も伴います。
2024 年、Google DeepMind のデミス・ハサビスとジョン・ジャンパーは、実験的に測定された何千もの形状から学習してタンパク質の三次元構造を予測するニューラル ネットワーク AlphaFold により、化学分野でノーベル賞の一部を受賞しました。この悪魔的な問題は、半世紀にわたって組織的な攻撃に耐えてきました。 AlphaFold はそれをきっぱりと解決したかに見え、世界はそのアプローチの有望さに注目するようになりました。ハサビス氏と彼のチームは、AlphaFold を「AI が科学のすべてをデジタル化の速度に加速する方法のテンプレート」と呼んでいます。 DeepMind の成功に後押しされて、生物学、化学、材料発見のための基礎モデルを構築する新興企業の波が数十億ドルを集めました。 AlphaFold は、AI と十分なデータを組み合わせることで (たとえ関係する基礎的なメカニズムが理解できなかったとしても) 画期的な発見を生み出す可能性があることを示し、再び科学の残りの部分を通る道が私たちの前に提示されたように見えました。
確かに、AI は科学に驚異的な変化をもたらすでしょうが、AlphaFold やそれに類するものはそうではない可能性があることがますます明らかになってきています。

その変身に最適なテンプレートです。これは大きな成果ではあるが、AlphaFold のような製品が生み出された条件は稀であり、他の分野でその条件を満たすまでにかかる時間は数年ではなく数十年単位になるだろう。代わりに、科学の加速は別のアプローチ、AI エージェントのおかげで実現します。
AlphaFold の成功の主な条件は、DeepMind のチームがモデルをトレーニングできる、実験的に検証された約 170,000 個のタンパク質構造のデータ セットであるタンパク質データ バンクの存在でした。タンパク質データバンクの創設は単純ではありませんでした。構築には 53 年間にわたる国際的な科学協力と、最近の推定によると約 210 億ドル相当の実験作業が必要でした。この規模の取り組みは、資金調達が困難であることは悪名高いほどで、調整するのはほぼ不可能で、実行には膨大な時間がかかります。その結果、彼らは失敗することがよくありました。
しかし、必要な結束力とリソースがあり、関連データが商業的所有権によってアクセスできなくなることがない分野であっても、別の障壁、つまり比較可能なデータを生成することが科学的に不可能であることについてはあまり議論されていません。タンパク質構造の場合、重要な実験技術であるタンパク質結晶構造解析は、非常に再現性が高く信頼できるツールであり、25 を超えるノーベル賞がこの技術に依存しているほどです。しかし、ほとんどの実験科学では、結果が異なることがよくあります。細胞株はドリフトします。化学薬品には微量汚染物質が含まれています。実験室の湿度が変化します。生物学やほとんどの化学分野で最新のニューラル ネットワークをトレーニングできるほど一貫性があり、十分に正確で、十分に拡張性のある測定データセットを作成するには、新しい種類の測定と新しい標準化されたアプローチが必要になります。

いつでもすぐに準備できます。
もちろん、天気予報、ゲノミクスの多く、化学の非常に限られた分野など、これらの要件が満たされる分野はいくつかあります。まだ AlphaFold スタイルの画期的な進歩が見られていないとしても、近いうちにこれらの分野に現れるかもしれません。新興バイオテクノロジーに関する米国国家安全保障委員会が主張しているように、これらのデータセットの作成と調整に対する政府の支援が重要になります。しかし、科学における未解決の疑問のほとんどについては、少なくとも短期的には別の計画が必要になるだろう。幸いなことに、より静かで控えめなものが有望になり始めています。
科学者は常に不確実性の下で推論してきました。新しい薬剤標的の特定に取り組んでいる生物学者は、これまで完璧なデータセットを持っていませんでした。代わりに、彼らはドッキング計算と既知の構造を組み合わせ、分子動力学を考慮し、いくつかの結合アッセイを実行し、独自の判断を使用して、その特定の強みと失敗点に従って各方法を比較検討します。科学のスキルは単一のツールにあるわけではありません。多くのツールが生成するものを統合し、証拠が得られると結果を修正することです。これが、ほとんどの実用的な研究が実際に進む方法です。しかし、つい最近まで、それを実行できるソフトウェアはありませんでした。
エージェントはそれができるようになりました。簡単に言えば、エージェントは、デジタルまたは物理的なツールとそれらを使用する機能へのアクセスを与えられた AI 推論エンジンです。ここ数年、AI の根本的なアーキテクチャの変化により、大規模な言語モデルを活用したこれらのプログラムの急速な普及が可能になり、科学的に特化したデータセットの必要性が劇的に減少しました。科学にとって、この技術の進歩は根本的な変化を表しています。これにより、実際の研究の反復的で偶然性の高いプロセスを模倣できるデジタル ツールを作成できるようになりました。一方、A のようなツールは、

lphaFold は限られた質問に強力なアプローチを適用します。エージェントは本質的にジェネラリストです。これらは科学を行うための新しい方法を表すものではなく、人間の発見プロセスをデジタル的にモデル化したものです。
5 月に発表された Google の AI 共同科学者について考えてみましょう。研究者らはこの論文に1ページの概要と目標を与えた：薬剤耐性感染症の主要な要因である抗生物質耐性が細菌種間でどのように広がるかを解明する。システムはサブエージェントを起動しました。ある人は文献から仮説を立てました。別の人は査読者のようにそれらを区別しました。 3 番目はトーナメントを開催して、最強の候補者をランク付けしました。 4 番目の仮説は、勝利仮説を洗練させました。エージェントは、耐性遺伝子が細菌ウイルスに乗り、ウイルスを新しい宿主に運ぶことができるものを借用していると結論付けた。仮説は正しかったのです。インペリアル・カレッジ・ロンドンの研究者たちは、10 年を費やして、骨の折れるウェットラボ研究を通じて同じ結論に達しました。彼らの論文は、これまで共同研究者によって確認されていなかったが、まだ査読中だった。
Co-Scientist のようなエージェントはまだ新しいツールであり、科学プロセスの遍在的な部分になる前に克服すべき本当の課題があります。エージェントは依然として幻覚を起こしやすく、判断が一貫していません。また、自律的に実行できる時間を制限する記憶と入力の制約があります。しかし、これらの技術的な障壁は取り除かれ、それが進むにつれ、私たちは科学が行われる信頼性、一貫性、速度に対する科学的要因の複合的な影響に気づき始めるでしょう。
おそらく最も注目に値するのは、エージェントが科学の「再現性の危機」、つまり研究者がお互いの結果を再現できないという広範な問題に対して構造的な解決策を提供していることです。何十年もの間、科学界は研究者に生のデータを共有するよう懇願してきました。

実験プロセスを標準化するための正確なコード。しかし、研究者たちは、興味深い科学がすでに完成した後に行われるこの退屈な管理作業に長い間抵抗してきました。対照的に、エージェントはすべての動きを自動的に記録し、結果に至った方法の正確な記録を作成し、正確な複製を可能にします。
2 番目の結果は、科学的記憶の増幅です。研究者間の知識の伝達が曖昧なプロセスであることはよく知られています。それが何年にもわたる訓練と観察を経て行われなければ、大学院生は先任者たちが数十年にわたって保管してきた乱雑な研究ノートを熟読し、自分たちのプロトコルを成否する詳細を探すことになる。しかし、科学プロセスにおいてエージェントがますます大きな部分を占めるようになるにつれて、研究室の科学史全体が、組織の知識の中央の標準化されたリポジトリに記録されることになります。
しかし、エージェントの最も重要な影響はスピードです。どの分野でも、アイデアをテストする方が会議で議論するよりも時間がかからない場合、人々は議論をやめてテストを実行します。 1 時間で 1,000 件の論文を読み、500 個の分子を設計し、朝までに失敗したテストから学習できるエージェントがあれば、実験のコストが削減され、科学の進歩のペースが根本的に変わります。また、研究者は、これまで時間をかけて取り組むことのなかった大胆で奇妙な質問を自由に追求できるようになり、私たちがまだ想像していなかった科学の扉が開かれることになります。
AlphaFold テンプレートは確かに驚くべき発見の鍵となるでしょうが、それだけでは科学の終焉をもたらすことはできません。むしろ、エージェント AI への移行は、科学のあらゆる分野を一度に包み込むツールという、はるかに稀なレベルのブレークスルーを表しています。歴史的には、

微積分、統計的推論、分光法、コンピューターなど、そのような範囲が登場したのはほんの数回です。それぞれが、誰も定式化することを考えていなかった問題の世界を明らかにし、それらの問題が順番に、彼らの分野を新たに定義しました。エージェントに関しては、そのような変革がもう一つ起こりつつあります。
エリック・シュミットは、2001 年から 2011 年まで Google の CEO を務めました。2024 年に妻のウェンディとともに、科学技術における型破りな探求分野に資金を提供する慈善事業であるシュミット・サイエンシズを共同設立しました。
スハス・マヘシュ氏は、シュミット科学 AI センターで AI for Science の取り組みを率いています。彼は材料発見のための AI の専門家です。
エリック・シュミット事務所のアソシエイト兼サイエンスリーダーのマヤ・レビンによる追加研究。
エリック・シュミット & スハス・マヘシュ著
あるスタートアップは、LLM を妨げているボトルネックを突破したと主張する Will Douglas Heaven
根本的な欠陥により、LLM はウィル・ダグラス・ヘブンの攻撃に対して著しく脆弱になります
欧州の不妊治療団体ジェシカ・ハムゼロウ氏、精子提供者には制限が必要だと語る
内部受容: キャサリン・W・アイザックスの中で自分がどのように感じているかという隠された感覚
あるスタートアップは、LLM を妨げているボトルネックを突破したと主張しています
Subquadratic は、新しいモデルに関する詳細を共有しました。しかし、まだ懐疑的な人もいます。
ウィル・ダグラス・ヘブンのアーカイブページ
根本的な欠陥により、LLM は攻撃に対して著しく脆弱になります
これにより、航空機のナビゲーション システムを妨害する方法を教えるなど、してはいけないことを簡単に騙すことができます。
ウィル・ダグラス・ヘブンのアーカイブページ
アントロピックは、クロードが概念について頭を悩ませる隠れた空間を発見しました。
新しい技術により、同社は LLM の奇妙な仕組みをこれまで以上に深く調査できるようになりました。
ウィル・ダグラス・ヘブンのアーカイブページ
Claude Science は Anthropic の最新の主力製品です
同社は科学分野での AI の活用を強化しています。
ゲ

からの最新アップデート
MITテクノロジーレビュー
特別オファー、トップニュース、
今後のイベントなど。
プライバシー ポリシー メールを送信していただきありがとうございます。
何か問題があったようです。
設定を保存できません。
このページを更新して更新してみてください
さらに時間がかかります。このメッセージが引き続き表示される場合は、
までご連絡ください
customer-service@technologyreview.com に受信を希望するニュースレターのリストを添えて送信してください。
レガシーの最新版
MIT テクノロジーレビューによる広告掲載
リンクトインは新しいウィンドウで開きます
インスタグラムが新しいウィンドウで開きます
フェイスブックは新しいウィンドウで開きます

## Original Extract

Every few decades, someone announces that science has reached its end. In 1903, the revered physicist Albert Michelson wrote that the “facts of physical science have all been discovered.” In the 1980s, Stephen Hawking predicted that theoretical physics might be finished by the end of the century. Wi
[truncated]

Skip to Content MIT Technology Review Featured
MIT Technology Review Featured
Opinion AI for science needs reasoning, not just data
AI agents that can model the human process of research will accelerate discoveries in science.
Every few decades, someone announces that science has reached its end. In 1903, the revered physicist Albert Michelson wrote that the “facts of physical science have all been discovered.” In the 1980s, Stephen Hawking predicted that theoretical physics might be finished by the end of the century. With the explosive arrival of artificial intelligence, the feeling is in the air again—this time accompanied by a Nobel Prize.
In 2024, Demis Hassabis and John Jumper of Google DeepMind were awarded part of the Nobel in chemistry for their neural network AlphaFold, which predicts the three-dimensional structures of proteins by learning from thousands of experimentally measured shapes. This devilish problem had resisted systematic attacks for half a century; AlphaFold seemed to have solved it once and for all, and the world became fixated on the promise of its approach. Hassabis and his team called AlphaFold “the template for how AI can accelerate all of science to digital speed.” A wave of startups building foundation models for biology, chemistry, and materials discovery raised billions of dollars, buoyed by DeepMind’s success. AlphaFold had shown that the combination of AI and sufficient data could make groundbreaking discoveries (even if we did not understand the underlying mechanisms involved), and it seemed, once again, that a path through the rest of science was laid out before us.
To be sure, AI will bring extraordinary changes to science, but it has become increasingly clear that AlphaFold, and things like it, may not be the best template for that metamorphosis. Though it is a profound achievement, the conditions that produced the likes of AlphaFold are rare, and the time it will take to meet those conditions in other fields will be measured in decades, not years. Instead, the acceleration of science will come about thanks to another approach: AI agents.
The primary condition for AlphaFold’s success was the existence of the Protein Data Bank, a data set of roughly 170,000 experimentally validated protein structures on which DeepMind’s team could train its model. The creation of the Protein Data Bank was not simple: It took 53 years of international scientific cooperation and, by a recent estimate , roughly $21 billion worth of experimental work to assemble. Efforts of that scale are infamously difficult to fund, next to impossible to coordinate, and hugely time-consuming to execute; they have often been unsuccessful as a result.
But even in fields with the requisite cohesion and resources, and where the relevant data are not rendered inaccessible by commercial ownership, another barrier is too little discussed: the scientific impossibility of generating comparable data. In the case of protein structures, the key experimental technique—protein crystallography—is an unusually replicable and dependable tool, so much so that over 25 Nobel Prizes have relied on it. But in most of experimental science, results vary more often than not. Cell lines drift. Chemicals have trace contaminants. Lab humidity changes. The creation of measured datasets that will be consistent enough, accurate enough, precise enough, and scalable enough to train a modern neural network in biology or most of chemistry would require new kinds of measurement and new standardized approaches—none of which will be ready anytime soon.
Of course, there are a handful of fields where these requirements are met: weather forecasting, much of genomics, very limited areas of chemistry. These may see AlphaFold-style breakthroughs soon, if they haven’t already . Government support for the production and coordination of those datasets will be critical, as the US National Security Commission on Emerging Biotechnology has argued . But for most open questions in science, we will need a different plan, at least in the short term. Luckily, something quieter and more modest has begun to show promise.
Scientists have always reasoned under uncertainty. Biologists working to identify new drug targets have never had perfect datasets. Instead, they combine docking calculations and known structures, factor in molecular dynamics, run a handful of binding assays, and use their judgment to weigh each method according to its particular strengths and points of failure. The skill of science is not in any single tool; it is synthesizing what many tools produce, and revising the results as the evidence comes in. This is how most working research actually proceeds. But until very recently, no software could do it.
Agents now can. Simply put, an agent is an AI reasoning engine that has been given access to tools—digital or physical—and the capabilities to use them. Over the last few years, a fundamental architectural shift in AI has enabled the rapid proliferation of these programs, which are powered by large language models, dramatically reducing the need for scientifically specialized datasets. For science, this technological advancement represents a foundational change: it has allowed us to create digital tools that can mimic the iterative, highly contingent process of actual research. While tools like AlphaFold apply a powerful approach to a limited question, agents are inherently generalists. They do not represent a new way to do science—instead, they digitally model the human process of discovery.
Consider Google’s AI Co-Scientist , announced in May. Researchers gave it a one-page brief and a goal: Figure out how antibiotic resistance spreads between bacterial species, a key driver of drug-resistant infections. The system spun up sub-agents. One drafted hypotheses from the literature. Another picked them apart like a peer reviewer. A third ran tournaments to rank the strongest candidates. A fourth refined the winning hypothesis. The agent concluded that resistance genes were hitching rides on bacterial viruses, borrowing whichever virus could ferry them into a new host. The hypothesis was correct. Researchers at Imperial College London had spent a decade reaching the same conclusion through painstaking wet-lab work; their paper, previously unseen by Co-Scientist, was still in peer review.
Agents like Co-Scientist are still novel tools, and there are real challenges to overcome before they become a ubiquitous part of the scientific process: They are still liable to hallucinate, their judgment is not consistent, and they have memory and input constraints that limit the time they can run autonomously. But these technical barriers will fall away, and as they do we will begin to notice the compounding effects of scientific agents on the reliability, consistency, and velocity with which science is done.
Perhaps most notably, agents offer a structural fix for science’s “reproducibility crisis,” the widespread problem of researchers’ inability to replicate each other’s results. For decades, the scientific community has begged researchers to share their raw data and exact code in an effort to standardize experimental processes. But researchers have long resisted this tedious administrative work, which happens after the interesting science is already done. Agents, in contrast, automatically log every move they make, creating an exact record of the method that led to their results and allowing for precise replication.
A second consequence will be an amplification of scientific memory. The transfer of knowledge between researchers is a famously murky process; if it isn’t done over years of training and observation, graduate students are left to pore through the messy lab notebooks kept by decades of predecessors, looking for the details that will make or break their protocol. As agents become an increasingly large part of the scientific process, though, a lab’s entire scientific history will be recorded in a central, standardized repository of institutional knowledge.
But the most important impact of agents will be speed. In any field, when testing an idea takes less time than arguing about it in a meeting, people stop debating and just run the test. An agent that can read a thousand papers in an hour, design 500 molecules, and learn from its failed tests by morning will bring down the cost of experimentation and fundamentally change the pace at which science gets done. It will also give researchers the freedom to chase bold, strange questions they never would have risked their time on before, opening scientific doors we have yet to imagine.
While the AlphaFold template will certainly be key to incredible discoveries, it alone will not bring us to the end of science. Instead, the shift toward agentic AI represents a much rarer tier of breakthrough: a tool that envelops every field of science at once. Historically, tools of such scope have arrived just a handful of times: calculus, statistical inference, spectroscopy, the computer. Each revealed a world of problems no one had thought to formulate, and those problems, in turn, defined their fields anew. With agents, another such transformation is upon us.
Eric Schmidt was the CEO of Google from 2001 to 2011. In 2024, with his wife Wendy, he co-founded Schmidt Sciences, a philanthropic venture to fund unconventional areas of exploration in science & tech.
Suhas Mahesh leads AI for Science work at the AI Center of Schmidt Sciences. He is a specialist in AI for materials discovery.
Additional research by Maya Levin, associate and sciences lead, Office of Eric Schmidt.
by Eric Schmidt & Suhas Mahesh
A startup claims it broke through a bottleneck that’s holding back LLMs Will Douglas Heaven
A fundamental flaw leaves LLMs strikingly vulnerable to attack Will Douglas Heaven
Sperm donors need limits, says a European fertility group Jessica Hamzelou
Inside interoception: The hidden sense of how you feel inside Katherine W. Isaacs
A startup claims it broke through a bottleneck that’s holding back LLMs
Subquadratic has now shared more details about its new model. But some are still skeptical.
Will Douglas Heaven archive page
A fundamental flaw leaves LLMs strikingly vulnerable to attack
It makes it easy to trick them into doing things they shouldn’t, such as telling you how to sabotage an aircraft’s navigation system.
Will Douglas Heaven archive page
Anthropic found a hidden space where Claude puzzles over concepts
A new technique has let the company probe deeper than ever into the weird workings of an LLM.
Will Douglas Heaven archive page
Claude Science is Anthropic’s newest flagship product
The company is doubling down on AI for science.
Get the latest updates from
MIT Technology Review
Discover special offers, top stories,
upcoming events, and more.
Privacy Policy Thank you for submitting your email!
It looks like something went wrong.
We’re having trouble saving your preferences.
Try refreshing this page and updating them one
more time. If you continue to get this message,
reach out to us at
customer-service@technologyreview.com with a list of newsletters you’d like to receive.
The latest iteration of a legacy
Advertise with MIT Technology Review
linkedin opens in a new window
instagram opens in a new window
facebook opens in a new window
