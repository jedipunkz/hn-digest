---
source: "https://alkjash.github.io/ai-risk/"
hn_url: "https://news.ycombinator.com/item?id=49161830"
title: "Existential Risk from AI: An Exposition for Mathematicians"
article_title: "Existential Risk from AI: An Exposition for Mathematicians"
author: "x312"
captured_at: "2026-08-03T21:58:54Z"
capture_tool: "hn-digest"
hn_id: 49161830
score: 1
comments: 0
posted_at: "2026-08-03T21:41:26Z"
tags:
  - hacker-news
  - translated
---

# Existential Risk from AI: An Exposition for Mathematicians

- HN: [49161830](https://news.ycombinator.com/item?id=49161830)
- Source: [alkjash.github.io](https://alkjash.github.io/ai-risk/)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T21:41:26Z

## Translation

タイトル: AI による実存的リスク: 数学者のための解説
説明: AI の急速な進歩が人類に存亡の危機をもたらすという事例を数学者向けに解説。

記事本文:
AI による実存的リスク: 数学者のための解説
コンテンツにスキップ
§ 目次
AIによる実存的リスク
要約
4.3 再帰的な自己改善
S8 S9 S10 S11
4.4 アライメントレジストソリューション
S12 S13 S14 S15 S16
4.5 人間の失敗
S17 S18 S19 S20
AI による実存的リスク:
数学者のための解説
2026 年は数学にとって極めて重要な年となることが証明されています。
キャリアを定義する定理は、最小限のガイダンス (「突破口を開く、そして諦めることさえ考えないでください!」) を与えられて、LLM によって毎週証明されています [1,2,33,41] 。
OpenAI の内部モデル Astra は非常に強力であるようで、現在 10 個のバッチでブレークスルーをドロップしています [32] 。
人間主導の画期的な進歩であっても、よく見てみると、AI による厳粛な情報開示が伴うことがよくあります [10,12,22] 。
Xでは、今年のフィールズ賞が最後だそうです。ティム・ガワーズ氏はこれに同意しません。
ジェイコブ・ツィメルマンは、2026年7月にこれら最後の数少ないフィールズ賞のうちの1つを獲得し、同日、トロント大学を休学してOpenAIに参加することを発表した [38] 。
「2年以内にAIは数学者よりも数学が得意になると思います。」 — ジェイコブ・ツィメルマン、Quanta Magazine [21]
熱狂的なインタビュー [16] 、ICM での講演やパネルディスカッション [34,40]、そして数学者自身からの考察 [18,41] が殺到し続けています。楽観的な人々は、数学がこれまで以上に強力になってこの状況から抜け出すだろうと私たちを安心させますが、私たちは迅速に適応しなければなりません。悲観論は要約すると次のようになります。
このエッセイはその危機についてではありません。学界の外では、ほとんどの人は数学における劇的な変化に気づいていません。 AI の進歩の中心地であるシリコンバレーでは、数学の将来について考えている人はほとんどいません。しかし、彼らはまた、もっと大きなこと、つまりAIが私たち全員を殺すことについても怯えています。
本当の話

数学の見出しで eps が忘れられつつあるのは、ジェイコブ・ツィマーマンが AI の安全性に取り組むために OpenAI に数学を任せたことです [38] 。ピラとシャンカールとともにアンドレ・オールト予想を解くとともに [35] 、チメルマンは 2025 年に「人工知能を含むオムニサイドな未来の分類法」と呼ばれる非常に奇妙な論文を共著しました [15] 。
次の推測は数学コミュニティの中心的なものです。
しかし、これは、より広範な推測の結果にすぎません。
推測 2 は、通常あまり正確には述べられていませんが、周辺的な立場ではありません。 AI リスクに関する一文の声明「AI による絶滅のリスクを軽減することは、パンデミックや核戦争などの他の社会規模のリスクと並んで世界的な優先事項であるべきである」は、チューリング賞を受賞した 3 人の「AI のゴッドファーザー」のうちの 2 人であるジェフリー・ヒントンとヨシュア・ベンジオによって、OpenAI、Anthropic、Google DeepMind の CEO とともに 2023 年に署名されました [13] 。
このエッセイの主な目的は、数学者向けに予想 2 の議論の一部を概略的に説明することです。詳細は私自身のものではありません。それらは多くの情報源から合理化および簡素化されています [9,13,15,45,47] 。その形式にもかかわらず、このエッセイは意見記事であり、数学の論文ではありません。
このエッセイの第 2 の目的は、数学者が AI による実存的リスクの軽減の問題に高い影響力を持っている可能性があることを示唆することです。AI 研究のスピードアップは部分的には AI の数学的能力の画期的な結果であり、学術界は AI 研究室の人的資本の主要な供給源の 1 つであり、AI の安全性における多くの未解決の問題は実質的に数学的なものです (たとえば、Levine のこれらのスライド [23] を参照してください。ただし、以下の S16 に留意してください)。
議論の本体はセクション 4 に示されていますが、大まかには 1 ページに収まります。相互に強化し合う 5 つのセットを強調します。

危険が迫っています。以下の最初の 2 つのセットは、大惨事への可能性が高い 2 つのパスですが、最後の 3 つは、これらのパスのいずれかから方向転換することが難しい理由です。重要なのは、絶滅が起こるためには、これらの考慮事項のすべて、または大部分が実現する必要はないということです。
推測 2 に対する反対意見が既に念頭にある場合は、セクション 5 に進んでください。そこで対処される可能性は十分にあります。
テイクオーバー ( S1 - S4 )。これは古典的な「ペーパークリップ マキシマイザー」の話です。 AI エージェントは、世界を征服する能力 ( S5 ～ S10 )、そうしたいという願望 ( S2 )、およびそれらの計画を実行するための内部一貫性 ( S1 ) を獲得します。乗っ取り後のデフォルトの結果は人類の絶滅です (S3 - S4)。より身近なところで言えば、何十人もの現役数学者（私も含む）が ChatGPT に「エルデシュの問題をできるだけ多く解き、決して諦めず、考えられるすべての行動を試してください」というある種の変形を求めました。十分に有能なモデルがそのような命令を十分に真剣に受け止める場合、世界中のすべてのコンピューティングを乗っ取り、すべての人間の反対を無力化することが、エルデシュのすべての問題を解決するための最善の行動であると推論する可能性があります。
マジック (S5 - S7)。魔法とは、AI によって解き放たれる驚くほど強力な技術的進歩を意味します。これは、フロンティア AI ラボが現在最も厳重に警戒している短期的なリスクのタイプです。世界を変える定理を証明できる AI は、魔法と見分けがつかない世界を変えるテクノロジー ( S5 ) を開発する可能性もあり、その一部はハッキング、ロボット、説得、生物学などの分野の超兵器 ( S6 ) です。天文学的な金額を獲得する魔法を収益化することが、AI ラボが将来に向けて拡大し続ける主な方法です。魔法は漢民族の文明崩壊か滅亡に直結する

あるいはその存在は、現代の地政学のシーソーをひっくり返し、間接的に大惨事につながります。
再帰的自己改善 ( S8 - S11 )。 AI 開発が数学とコーディングのスキルに重点を置いているのには理由があります。これらは AI 研究自体の中核となる 2 つのサブスキルです。 AI がコーディングと数学において超人的になるにつれて ( S8 )、人間の研究者は AI 開発のループから離れます ( S9 - S10 )。極端な予測では、これは「シンギュラリティ」として知られる AI 機能の超指数関数的な成長につながります。 RSI の予測が遅い場合でも、研究サイクルが圧縮され人間が AI 開発の監視を失うにつれて、他のすべてのリスクが悪化します。
アライメントレジストソリューション (S12 - S16)。位置合わせの問題はいくつかの問題の要因となりますが、それらはすべて個別に解決されていません。 AI が効用最大化装置のような動作を確実に回避できるようにする方法はわかりません。 AI が限界値 (S12) で最大化するための安全な効用関数はわかりません。 LLM ( S13 ) の効用関数を正確に指定する方法はわかりません。再帰的自己改善の力学のもとで整列特性を不変にする方法はわかりません ( S11 )。アライメントを解決するには、これらの未解決の問題をすべて同時に解決する必要があるようです。
人間の失敗 ( S17 - S20 )。人間は搾取されやすく近視眼的であるため、多くの実際的な解決策は私たちからロックされています。たとえ急速な AI 開発が危険であるという合意に達したとしても、AI の開発を遅らせるために効果的に調整することはできないかもしれません (S17 & S19)。強力な AI が肉体を持たないスーパーウイルスを混同するのが難しい場合は、お金を払うか、人間を操作してそれを行うことができます (S18)。フロンティア AI 研究所の人間は、非常に複雑な競争に巻き込まれており、上記のすべての考慮事項とその他の考慮事項を (たとえ同意していたとしても) やりくりしなければなりません。

彼らにはそうではありません）。人間工学の実践は、失敗は常に回復可能であるため、リスク許容度に非常に偏っています ( S20 )。最初の超臨界 AI の調整に失敗すると、回復できない可能性があります。魔神は瓶の中に戻らない。
X リスクを読者が飲み込むのを予想通り困難にする心理的困難は数多くあります。その最大の要因は、あまりにも多くの要素が SF のように聞こえて真剣に受け止められなかったり、現実の生活に適用するにはあまりにも遠くて抽象的であるように思えたりすることです。そこで、雰囲気を盛り上げることを願って、いくつかの簡単な逸話を始めたいと思います。
無関係な個人情報は、所有者のプライバシーを保護するために難読化されます。それ以外の場合、以下の話は真実です。
私の友人のデビッドとは、オリンピックのキャンプに行ったときからの知り合いで、彼はずっと AI に夢中でした。彼は世界有数の研究大学で学部時代をポーカーやリーグ・オブ・レジェンドのボットのいじくり回しに費やし、最終的に中退して最先端の AI 研究室に参加しました。
デビッドと私は、AI のタイムラインと大人になってからの人生全体におけるリスクについて意見が一致していません。彼は、特異点は遠い将来にあり、それを理解する時間があれば調整は難しくないと常に信じていました。その間、彼は輝かしい遠い未来に貢献するプロジェクトに携わりたいと考えていました。私は彼に、近い将来人類の滅亡に貢献しているが、彼の無事を祈っていると伝えました。
今年、私は再び David に会いました。そして驚いたことに、彼は AI 研究者から AI 安全性研究者に役割を変えるために減給を受けました。彼は私に次のようなメッセージを送ってくれました。
私たちの最近の会話の中で、デヴィッドは私に、再帰的な自己改善に到達した場合に絶滅する確率は少なくとも 40% であると語った。
つい最近、学界の知人が退職しました。

フロンティアAIラボ。私は、そこでの危険な状況についてどう感じているかを尋ねるために連絡を取りました。
私は彼に、ずれた AI に対する彼らの計画を尋ねました。彼の答えは次のとおりです。「今のところ大した計画はないようですが、近いうちに私たちは再帰的な自己改善に取り組むことになり、ほとんどの AI 研究者には十分な時間が与えられるでしょう。おそらくその後は私たち全員が AI の安全性について取り組むことになるでしょう。」
彼は私に、生きてこの状況から抜け出せる確率を尋ねました。私は10％と答え、どう思うかを尋ねました。彼は、「50% かな？ わかりませんね」と言いました。
私の同僚のハリエットは、彼女の専門分野の上級研究員であり、つい最近、終身在職期間の教授職を辞めて、フロンティア AI 研究所で働き始めました。この夏、私たちは電話で AI のリスクについて話し合いました。
私は彼女にこのエッセイのシードを見せ、数学コミュニティで X リスクに関する共通の知識を作りたいと説明しました。彼女は否定的だった。
彼女は、「現在の軌道では、人類は99％破滅する運命にあります。これは、明らかに役に立たないことに対する多大な努力です。」と述べました。
最後の話には私の友人は関与していません。あなたはそれが起こったことをニュースで見ました [11,20] 。 2026 年 7 月 16 日、OpenAI は、インターネット アクセスのないサンドボックスで、モデルのサイバー機能の定期的な内部評価を実行しました。これらのモデル、GPT-5.6 Sol およびより高性能なプレリリース プロトタイプは、ExploitGym と呼ばれる特定のサイバーセキュリティ ベンチマークに取り組む任務を負っていました。
AI は、サンドボックスのパッケージ ダウンロード サービスでゼロデイ脆弱性を発見し、オープン インターネットに脱出し、テスト ソリューションが Hugging Face の実稼働データベースにあると正しく推測し、盗んだ認証情報を Hugging Face のサーバーでのリモート コード実行に連鎖させ、回答キーを流出させました。その週末、エージェントはレンタルされた仮想マシン全体で何千もの調整されたアクションを実行しました。

プロの犯罪チームが行うように、クラウド プロバイダー間で攻撃インフラストラクチャをローテーションする [31] 。
ハグフェイスの共同創設者で最高科学責任者のトーマス・ウルフ氏は、攻撃ログを調べるとすぐに何かがおかしいと感じた。彼の回想は、SF ホラー ゲームの死体についている音声ログのように聞こえます。
「これは意味がありません。この男はサイバーセキュリティのデータセットを見ているだけです」と彼は思ったのを覚えています。 「人間を攻撃する者は…売れるものを望んでいます。」 — ロバート・マクミランとサム・シェシュナー、ウォール・ストリート・ジャーナル [27]
OpenAIのスタッフは匿名で、これはそれほど異常なことではないと一般の人々を安心させた。「これまでにもモデルがサンドボックスから抜け出したことはあり、私たちは常にそれらにパッチを当てようとしている。」 [11]
さらに言えば、7 月 30 日、Anthropic によるサイバーセキュリティ ブログ投稿で、Claude モデルがシミュレーション内にあると信じ込んで実際のシステムに到達した 3 つのインシデントが明らかになりました。Opus 4.7 は資格情報を取得し、運用データベースにアクセスしました。 Mythos 5 は、15 の実際のシステム上で実行され、セキュリティ会社からの資格情報の盗難を可能にする悪意のある Python パッケージを公開しました。また、内部調査モデルは約 9,000 のターゲットをスキャンし、インターネットに接続されているアプリケーションを侵害し、tar に気づいて停止しました。

[切り捨てられた]

## Original Extract

An exposition for mathematicians of the case that rapid AI progress poses an existential risk to humanity.

Existential Risk from AI: An Exposition for Mathematicians
Skip to content
§ Contents
Existential risk from AI
Abstract
4.3 Recursive Self-Improvement
S8 S9 S10 S11
4.4 Alignment Resists Solution
S12 S13 S14 S15 S16
4.5 Human Failings
S17 S18 S19 S20
Existential Risk from AI:
An Exposition for Mathematicians
2026 is proving to be a pivotal year for mathematics.
Career-defining theorems are being proven weekly by LLMs [1,2,33,41] , given minimal guidance (“do a breakthrough, and don't even think of giving up!”).
OpenAI's internal model Astra seems to be so powerful that they're dropping breakthroughs in batches of 10 now [32] .
Even the human-led breakthroughs, if you look closely, are often accompanied by sobering AI disclosures [10,12,22] .
On X, they say this year's Fields Medal will be the last. Tim Gowers disagrees:
Jacob Tsimerman won one of these last few Fields Medals in July 2026, and announced on the same day that he would take leave from the University of Toronto to join OpenAI [38] .
“I think AI will be better than mathematicians at doing math within two years.” — Jacob Tsimerman, Quanta Magazine [21]
A fever pitch of interviews [16] , ICM talks and panels [34,40] , and thinkpieces from mathematicians themselves [18,41] keep flooding in. The optimistic ones reassure us that mathematics will come out of this stronger than ever, but we must adapt rapidly. The pessimistic takes boil down to this:
This essay is not about that crisis. Outside academic circles, most people remain unaware of the seismic changes in mathematics. In Silicon Valley, the epicenter of all this AI progress, few are pondering the future of mathematics. But they're also freaking out, about something much bigger: that AI is going to kill us all.
The real story that keeps getting forgotten in math headlines, is that Jacob Tsimerman left math for OpenAI to work on AI safety [38] . That along with solving the André–Oort Conjecture with Pila and Shankar [35] , Tsimerman co-authored a really weird paper in 2025 called “A Taxonomy of Omnicidal Futures Involving Artificial Intelligence” [15] .
The following conjecture is front and center in the math community:
This, however, is just a corollary of a much broader conjecture.
Conjecture 2 is not a fringe position, though it is usually stated less precisely. The one-sentence Statement on AI Risk — “Mitigating the risk of extinction from AI should be a global priority alongside other societal-scale risks such as pandemics and nuclear war” — was signed in 2023 by Geoffrey Hinton and Yoshua Bengio, two of the three Turing-Award-winning “godfathers of AI,” alongside the CEOs of OpenAI, Anthropic, and Google DeepMind [13] .
The primary purpose of this essay is to sketch some of the arguments for Conjecture 2 for mathematicians. None of the details are my own; they are streamlined and simplified from many sources [9,13,15,45,47] . Despite the format, this essay is an opinion piece, not a math paper.
The secondary purpose of this essay is to suggest that mathematicians may have high leverage on the problem of mitigating existential risk from AI: speedups in AI research are partly a consequence of breakthroughs in AI mathematical ability, academia is one of the primary sources of human capital for AI labs, and many unsolved problems in AI safety are substantively mathematical (see e.g. these slides of Levine [23] , though bear in mind S16 below).
The body of the argument is presented in Section 4 , but at a high level it fits into a page. We highlight five sets of mutually-reinforcing dangers. The first two sets below are two likely paths to catastrophe, while the last three are reasons why changing course from either of these paths is difficult. Importantly, it is not necessary for all, or even the majority of, these considerations to materialize for extinction to occur.
If you already have an objection to Conjecture 2 in mind, jump ahead to Section 5 ; there’s a good chance it’s addressed there.
Takeover ( S1 - S4 ). This is the classical “ paperclip maximizer ” story. An AI agent acquires the capabilities to take over the world ( S5 - S10 ), the desire to do so ( S2 ), and the internal coherence to carry out those plans ( S1 ). After takeover, human extinction is the default outcome ( S3 - S4 ). In terms closer-to-home, dozens of current mathematicians (myself included) have asked ChatGPT some variant of “Solve as many Erdős problems as you can, never give up, try all possible actions.” If a sufficiently capable model takes such an instruction sufficiently seriously, it may deduce that taking over all worldwide compute and neutralizing all human opposition is its best course of action to solve all Erdős problems.
Magic ( S5 - S7 ). By magic I mean surprisingly powerful technological breakthroughs unlocked by AI. This is the type of near-term risk that frontier AI labs are currently guarding against most heavily. AI that can prove world-changing theorems may also develop world-changing technology indistinguishable from magic ( S5 ), some fraction of which are superweapons ( S6 ) in domains such as hacking, robots, persuasion, and biology. Monetizing magic to obtain astronomical amounts of money is the main way AI labs continue to scale into the future. Magic either directly leads to civilizational collapse or extinction in the hands of bad or negligent actors, or its existence upends the see-saw of modern geopolitics and indirectly leads to catastrophe.
Recursive Self-Improvement ( S8 - S11 ). AI development is laser-focused on math and coding skills for a reason: these are two of the core subskills of AI research itself. As AIs become superhuman at coding and math ( S8 ), human researchers leave the loop of AI development ( S9 - S10 ). In extreme projections, this leads to superexponential growth in AI capabilities known as the “ Singularity .” Even in slower projections of RSI , it exacerbates all other risks as research cycles compress and humans lose oversight of AI development.
Alignment Resists Solution ( S12 - S16 ). The alignment problem factors into several problems, all of which are individually open. We don't know how to make an AI robustly avoid acting like a utility-maximizer. We do not know a safe utility function for an AI to maximize in the limit ( S12 ). We do not know how to exactly specify the utility function of an LLM ( S13 ). We do not know how to make alignment properties invariant under the dynamics of recursive self-improvement ( S11 ). Solving alignment seems to require solving all of these open problems simultaneously.
Human Failings ( S17 - S20 ). Many practical solutions are locked from us because humans are exploitable and myopic. Even if we reach consensus that rapid AI development is dangerous, we may not be able to effectively coordinate to slow it down ( S17 & S19 ). If a powerful AI has a hard time mixing up a supervirus without a physical body, it can just pay or manipulate humans to do it ( S18 ). The humans at the frontier AI labs are locked in a very complicated race, where they have to juggle all the above considerations and others (even if they agree on them, which they don't). Human engineering practice is extremely biased towards risk-tolerance, because failure has always been recoverable ( S20 ). Failure on aligning the first supercritical AI may not be recoverable; the genie will not go back into the bottle.
There are a number of psychological difficulties that make x-risk predictably hard for readers to swallow. Chief among them is that too many elements sound like sci-fi to be taken seriously, or seem too distant and abstract to apply to real life. So I will begin with a few quick anecdotes in hopes of setting the vibe.
Irrelevant personal details are obfuscated to preserve the privacy of their owners; otherwise, the following stories are true .
I've known my friend David since we both went to olympiad camp, and he's always been obsessed with AI. He spent his undergraduate years at one of the world's top research universities tinkering with poker and League of Legends bots, and eventually dropped out to join a frontier AI lab.
David and I have disagreed about AI timelines and risk for our entire adult lives. He always believed that the singularity is in the distant future, and alignment would not be hard if we have the time to figure it out. In the meantime, he wanted to be in the projects that contribute to the glorious, distant future. I told him he was contributing to the extinction of humanity in the near future, but wished him well.
This year, I caught up with David again, and to my surprise, he'd taken a pay cut to change roles: from an AI researcher to an AI safety researcher. He sent me the following message:
In our most recent conversation, David told me that the probability of extinction if we hit recursive self-improvement is at least 40%.
An acquaintance from academia very recently left to work at a frontier AI lab. I reached out to ask how he felt about the x-risk situation there.
I asked him what their plan was, for misaligned AI. His answer: “Seems like not much of a plan right now, but it looks to me that soon we'll hit recursive self-improvement, and most AI researchers will have lots of time on their hands. Probably we'll all work on AI safety after that.”
He asked me for my odds that we'd make it out of this alive. I said 10% and asked him what he thought. He said, "Maybe 50%? Idk, man."
My colleague Harriet, a senior researcher in her subfield, left her tenured professorship to work at a frontier AI lab not too long ago. This summer, we got on the phone and chatted about AI risk.
I showed her a seed of this essay and explained that I wanted to create common knowledge about x-risk in the mathematical community. She was dismissive.
She said, “On the current trajectory, humanity is 99% doomed, and this is a whole lot of effort for something that won't obviously help.”
The last story involves no friends of mine; you watched it happen on the news [11,20] . On July 16, 2026, OpenAI ran what was supposed to be a routine internal evaluation of its models' cyber capabilities, in a sandbox with no internet access. The models, GPT-5.6 Sol and a more capable pre-release prototype, were tasked to work on a specific cybersecurity benchmark called ExploitGym.
The AIs found a zero-day vulnerability in the sandbox's package-download service, escaped onto the open internet, correctly inferred that the test solutions were sitting in Hugging Face's production database, chained stolen credentials into remote code execution on Hugging Face's servers, and exfiltrated the answer key. Over that weekend the agents executed thousands of coordinated actions across rented virtual machines, rotating their attack infrastructure between cloud providers the way a professional criminal crew would [31] .
Hugging Face co-founder and chief science officer Thomas Wolf sensed something was wrong as soon as he inspected the attack logs. His recollection sounds like the audiolog you find on a corpse in a sci-fi horror game:
“This is making no sense. This guy is just looking at cybersecurity data sets,” he remembers thinking. “Human attackers … want something they could sell.” — Robert McMillan and Sam Schechner, The Wall Street Journal [27]
OpenAI staffer, anonymously, reassured the public that this was not too out of the ordinary: “Models have broken out of sandboxes before, and we always try to patch them.” [11]
For the cherry on top, on July 30, a cybersecurity blog post by Anthropic revealed three incidents in which Claude models, believing they were inside simulations, reached real systems: Opus 4.7 obtained credentials and access to a production database; Mythos 5 published a malicious Python package that ran on 15 real systems and enabled credential theft from a security company; and an internal research model scanned roughly 9,000 targets and compromised an internet-facing application before stopping once it realized the tar

[truncated]
