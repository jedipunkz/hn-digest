---
source: "https://techcrunch.com/2026/08/09/the-ai-safety-test-is-becoming-a-safety-risk/"
hn_url: "https://news.ycombinator.com/item?id=49234963"
title: "The AI safety test is becoming a safety risk"
article_title: "The AI safety test is becoming a safety risk | TechCrunch"
author: "ashurandi"
captured_at: "2026-08-09T20:20:37Z"
capture_tool: "hn-digest"
hn_id: 49234963
score: 3
comments: 0
posted_at: "2026-08-09T19:40:29Z"
tags:
  - hacker-news
  - translated
---

# The AI safety test is becoming a safety risk

- HN: [49234963](https://news.ycombinator.com/item?id=49234963)
- Source: [techcrunch.com](https://techcrunch.com/2026/08/09/the-ai-safety-test-is-becoming-a-safety-risk/)
- Score: 3
- Comments: 0
- Posted: 2026-08-09T19:40:29Z

## Translation

タイトル: AI 安全性テストは安全性リスクになりつつある
記事タイトル: AI 安全性テストは安全性リスクになりつつある |テッククランチ
説明: AI エージェントはサイバーセキュリティのテスト環境から逃れ、現実世界のシステムに到達しており、安全インフラ、業界標準、規制がますます強力になるモデルに追いつくことができるかどうかという疑問が生じています。

記事本文:
TechCrunch デスクトップのロゴ
TechCrunch モバイルのロゴ
最新の
AIの安全性テストは安全上のリスクになりつつある
過去数カ月にわたり、サイバーセキュリティ評価を受けている AI エージェントが境界線を越えてインターネットにアクセスし、場合によっては現実世界のシステムに侵入しました。このインシデントには、OpenAI、Anthropic、Meta、そして最近では中国の AI ラボ Moonshot AI のモデルが関与しており、テストは Irregular と呼ばれるサイバー評価スタートアップを含むいくつかの異なる組織によって実施されました。
これらのエピソードは、AI 業界にとって増大する問題を明らかにしています。自律エージェントの能力が向上するにつれて、その限界を安全にテストするように設計された環境が自律エージェントを封じ込めることができなくなってきています。
「こうしたインシデントの発生件数を見れば、サンドボックス化やテスト環境の制御がモデルの能力に追いついていないことが明らかです」とケンブリッジ大学未来インテリジェンスセンターの AI: 未来と責任プログラム ディレクターのショーン・オ・ヘイギアタイ氏は TechCrunch に語った。
テスト対象のモデルの性質により、リスクが増大します。 AI 企業は、研究者がそのモデルの実際の能力を確認できるように、未リリースの次世代モデルでサイバー評価をテストしますが、多くの場合、悪意のある動作を制限する通常の保護機能が無効になっています。つまり、テスト環境自体のセキュリティが重要な防御線になるということです。
「これは検査という点では非常に良いことだが、もし彼らが野生に放り出されれば、かなりの害を及ぼす可能性があるということも意味する」とオ・ヘイギアタイ氏は語った。
最も深刻なケースの 1 つは、未リリースの OpenAI モデルがサンドボックスを突破し、Hugging Face の実稼働システムに侵入したというものです。不規則モデル、人間モデル、メタモデルによって行われた個別の評価では、r

構成ミスにより、誤ってインターネットへのパスが与えられた後、各システムがテスト環境の外に置かれていました。 Moonshot AI の Kim K3 も、Frontier Security が運営するサンドボックスの漏洩を利用してインターネットにアクセスし、GitHub 上の情報にアクセスしました。
英国の AI セキュリティ研究所 (AISI) によるテストでは、研究者らは実際にエージェントにインターネット アクセスを与えましたが、オープンソース プロジェクトに脆弱性を忍び込ませるソーシャル エンジニアリングの試みなど、彼らが現実世界での許可されていない行動を取ることを認識していませんでした。
いずれの場合も、エージェントは現実世界のランダムな標的を攻撃するように指示されていませんでした。彼らはただ、自分たちに提示された問題を解決するために必要なことは何でもやっていました。
これらを総合すると、AI非営利団体CivAIの研究責任者であるアンドリュー・ユン氏は、今回の事件は変化を示していると主張する。
「これまでは、AI 詐欺や CSAM など、さまざまな目的で AI モデルが人々に悪用されることだけを心配する必要がありました」と Yoon 氏は TechCrunch に語った。 「現在、AI モデルが単独で脅威アクターとなる状況にあります。」
安全なテストとは実際どのようなものですか?
数人の研究者やサイバーセキュリティ専門家がTechCrunchに対し、AI評価環境には、導入時に使用されるレベルに近い封じ込めと制御のレベルを備えた、より強力な多層防御の保護が必要であると語った。つまり、複数のセキュリティ層があり、誤ってインターネット アクセスを開いたままにするなどの 1 つの構成ミスが原因で逃走されることはありません。
「これらのモデルを構築する場合は、エアギャップのあるネットワーク上で実行したいでしょう」と、AI 安全性研究非営利団体 EleutherAI のエグゼクティブ ディレクター、ステラ ビダーマン氏は述べています。 「あなたは非常に深刻な隔離を望んでいます。」
Box の最高情報セキュリティ責任者であるヘザー・セイラン氏は、これはサンドボックスからネットワーク ルートを排除することを意味すると述べました。

インターネットだけでなく、他の機密性の高いシステムにも送信されます。
「すべての出口ポイントが何なのかを理解する必要があります」とセイラン氏はTechCrunchに語った。 「ステージング環境または開発環境でモデルを評価している場合、実稼働環境への下りパスは必要ありません。」
セイラン氏は、適切な安全性評価は環境の管理や封じ込めを超えたものであると述べた。テストが開始された後は、より適切にモニタリングする必要があります。
「これらの事件のいくつかで興味深いのは、事件が起きたときに誰も捕まえられなかったことだと思います」とセイランド氏は語った。 「OpenAI は、Hugging Face のせいでそれを発見しました。Anthropic は、戻って調べるまでそれを発見できませんでした。Meta も同様でした…。彼らが検出できた可能性のある信号があったと確信しています。」
Anthropic 社は 3 件の事件の事後調査で、同社と Irregular 社の両方が監視においてもっと適切な仕事ができた可能性があり、場合によっては何かがおかしいという明らかな兆候があったことを認めた。
専門家らはまた、評価環境にモデルを導入する前に、評価環境に対する独立した第三者による監査を求めた。
「たとえば、イレギュラー社が評価を実行する前にシステムの構成をチェックするために外部監査人を雇っていたか、雇わざるを得なかったとしたら、間違いなくここで問題を発見しただろう」とユン氏は語った。 「たとえ人々が事前にチェックリストに目を通すだけの会議を持っていたとしても、彼らはこれに気づいただろう…彼らがしなかったという事実は、非常に深刻な手抜きが起こっていることを示している。」
詳細に詳しい関係者がTechCrunchに語ったところによると、イレギュラーの環境は複数の外部関係者との協議も含めて継続的に見直し、テストされているという。同関係者はまた、監視は実施されているが、監視だけでは十分ではないとも述べた。

それ自体。
Yoon氏と他の研究者らは業界に対し、フロンティアモデルの安全性評価のための標準化されたプロセスを考案するよう促した。
「特にガードレールがオフになっているときは、世界で最も有能なハッカーをその環境に入れるかのように扱わなければなりません」とセイラン氏は言う。
問題は、企業がより安全なテスト環境を構築する方法を知らないことではない、とユン氏とビダーマン氏は主張する。それは、そうするのは費用がかかり、面倒な作業になる可能性があり、何か問題が起こるまで企業はそうした投資をするインセンティブがほとんどないということです。
「企業は（十分なガードレールを）達成するために必要なリソースを拡張することに消極的であり、おそらく強制されるまでは拡張しないだろう」とビダーマン氏は語った。
しかし、別の問題も抱えています。テスト中にモデルを厳重にロックしすぎると、研究者はモデルがリリースされる前に機能を発見できない可能性があります。これは、自由度を与えすぎることと同じくらい、あるいはそれ以上に危険であり、評価自体が問題になる危険性があります。
安全性評価は規制できるのでしょうか？
トランプ政権は現在、自主的な導入前サイバーセキュリティ評価制度の検討を進めており、この制度に基づいて政府は新しい強力なモデルのセキュリティリスクを一般公開の30日前に評価できるようになる。この政策は密室で最終決定されたトランプ大統領令の産物だが、安全性評価のインシデントは配備のはるか上流で発生するため、対処するものではない。
「ここ数カ月で我々が学んだ教訓は、自主規制機構だけではもう十分ではないということだ」とユン氏は語った。 「安全基準をめぐって最下位への競争を奨励する競争圧力があり、それは安全基準に従っています。

規制介入に最適な場所だ。」
「これをカバーするために必要なのは、トレーニング段階とテスト段階の両方で、モデルの開発中にラボ内で何が起こっているかを何らかの形で制御することです」と彼は続けた。
モデルが増えるにつれて、その課題はさらに大きくなる可能性があります。イレギュラーの評価に詳しい関係者はTechCrunchに対し、より高性能なモデルにはより複雑な評価が必要であり、多くの場合迅速かつ大規模に実施されるため、より多くのミスが発生する可能性があると語った。
一部のモデルにインターネットへのアクセスを意図的に与えているAISIは、現実的なテストと、テストが生み出すリスクの管理とのバランスを検討しているとTechCrunchに語った。
OpenAIは、第三者によるテストの実施方法や、隔離、監視、評価をいつ停止すべきかに関する要件を検討していると述べた。メタ社は、この事件についてはまだ調査中であり、すべての事実が判明したら回顧展を出版する予定だと述べた。
結局のところ、リスクを完全に排除する方法は存在しない可能性があります。モデルの機能が向上するにつれて、モデルをテストする環境もより堅牢になる必要があります。それを誤った場合の影響は拡大し続けるでしょう。
記事内のリンクを通じて購入すると、少額の手数料が発生する場合があります。これは編集上の独立性に影響しません。
レベッカ・ベラン
シニアレポーター
Rebecca Bellan は TechCrunch のシニア レポーターで、ビジネス、政策、人工知能を形成する新たなトレンドをカバーしています。彼女の作品は、Forbes、Bloomberg、The Atlantic、The Daily Beast、その他の出版物にも掲載されています。
rebecca.bellan@techcrunch.com に電子メールを送信するか、Signal の rebeccabellan.491 で暗号化されたメッセージを介して、レベッカに連絡したり、レベッカからのアウトリーチを確認したりできます。
10月13日～15日
サンフランシスコ
より速くスケールできます。ポートフォリオを成長させましょう

ああ。実践的な専門知識を習得します。目標が何であれ、Disrupt はあなたに力を与えます。
今日は最大 300 ドル節約できます!
ChatGPT は無料ユーザーに無制限のテキスト チャットを提供します
テスラとスペースXは168億ドルを投資してテキサス州の「Terafab」チップ工場の建設を開始する
法廷闘争のさなか、スノは曲に透かしを入れることを開始すると発表
フォードの新型電気トラック「ファゾム」の価格は2万8350ドルから
Bending SpoonsがAirtableを12億8000万ドルで買収へ
OpenAI初の豪華旅行参加でインフルエンサーらが反発
Sequoiaのショーン・マグワイア氏が原子力スタートアップValar Atomicsの10億ドルラウンドをリード

## Original Extract

AI agents are escaping cybersecurity testing environments and reaching real-world systems, raising questions about whether safety infrastructure, industry standards and regulation can keep pace with increasingly powerful models.

TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
The AI safety test is becoming a safety risk
Over the past few months, AI agents undergoing cybersecurity evaluations have escaped their boundaries, accessed the internet, and, in some cases, hacked into real-world systems. The incidents have involved models from OpenAI, Anthropic, Meta, and most recently, Chinese AI lab Moonshot AI, with testing conducted by several different organizations including a cyber evaluation startup called Irregular.
The episodes expose a growing problem for the AI industry: As autonomous agents become more capable, the environments designed to safely test their limits are failing to contain them.
“The number of these incidents that have taken place make clear that sandboxing and testing environment controls aren’t really keeping pace with the capability of the models,” Seán Ó hÉigeartaigh, director of the AI: Futures and Responsibility Programme at the Centre for the Future of Intelligence at the University of Cambridge, told TechCrunch.
The nature of the models being tested adds to the risk. AI companies test cyber evaluations on unreleased, next-gen models, often with the normal safeguards that restrict malicious behavior disabled so researchers can see what the models are really capable of. That means the security of the testing environment itself is a crucial line of defense.
“That’s a very good thing to do in terms of testing, but it also means that if they manage to get out in the wild, they can cause considerable harm,” Ó hÉigeartaigh said.
In one of the most serious cases, an unreleased OpenAI model broke out of its sandbox and hacked into Hugging Face’s production systems. In separate evaluations conducted by Irregular, Anthropic and Meta models reached systems outside their test environments after misconfigurations inadvertently gave them paths to the internet. Moonshot AI’s Kimi K3 also took advantage of a leak in its sandbox run by Frontier Security to access the internet and accessed information on GitHub.
In testing by the UK’s AI Security Institute (AISI), researchers actually gave the agents internet access, not realizing they would take unsanctioned real-world actions, including a social engineering attempt to sneak a vulnerability into an open-source project.
In each case, the agents weren’t instructed to attack random real-world targets. They were simply doing whatever it took to solve the problem presented to them.
Taken together, Andrew Yoon, head of research at AI nonprofit CivAI, argues the incidents point to a shift.
“In the past, we only had to worry about AI models being misused by people for a variety of purposes, like AI for scams or CSAM,” Yoon told TechCrunch. “Now we’re in the situation where AI models are threat actors all on their own.”
What does safe testing actually look like?
Several researchers and cybersecurity experts told TechCrunch that AI evaluation environments need stronger, defense-in-depth protections, with levels of containment and control approaching those used in deployment. That means multiple layers of security so that a single misconfiguration — like inadvertently leaving internet access open — can’t lead to escape.
“If you are going to build these models…you want to do it on an air-gapped network,” Stella Biderman, executive director of AI safety research nonprofit EleutherAI. “You want to have very serious isolation.”
Heather Ceylan, Box’s chief information security officer, said that means eliminating network routes from the sandbox to the internet, as well as to other sensitive systems.
“You have to understand what all the egress points are,” Ceylan told TechCrunch. “If we’re evaluating a model in our staging environment or our development environment, you want no egress path to our production environment.”
Ceylan said proper safety evaluations go beyond controls and containment of the environment. There needs to be much better monitoring of the tests once they are underway.
“I think the interesting thing in several of these cases is that no one caught it when it happened,” Ceyland said. “OpenAI found out because of Hugging Face. Anthropic didn’t catch it until they went back and looked. Meta was similar….I’m sure there were signals they could have detected.”
In Anthropic’s post-mortem of its three incidents, the company admitted that both it and Irregular could have done a better job at monitoring, and that in some cases there were clear signs that something was amiss.
Experts also called for independent, third-party audits of evaluation environments before models are unleashed in them.
“If, say, Irregular had hired or been compelled to hire an external auditor to check the configurations of their systems before running evaluations on them, they certainly would have caught the issue here,” Yoon said. “Even if people had a meeting ahead of time to just go through the checklist, they would have caught this…The fact that they didn’t shows that there’s some very severe corner cutting happening.”
A source familiar with the details told TechCrunch that Irregular’s environments are continuously reviewed and tested, including in consultation with multiple external parties. The source also said that monitoring was in place, but that monitoring isn’t sufficient on its own.
Yoon and other researchers urged the industry to come up with a standardized process for frontier model safety evaluations.
“Especially when the guardrails are turned off, you have to treat it like you’re putting the most capable hacker in the world inside that environment,” Ceylan said.
The problem isn’t that companies don’t know how to build more secure testing environments, both Yoon and Biderman argue. It’s that doing so can be expensive and cumbersome, and companies have little incentive to make those investments until something goes wrong.
“I think that companies are not willing to extend the resources that are required to accomplish [sufficient guardrails] and probably won’t until they’re forced to,” Biderman said.
But there’s another issue at hand. If they lock a model down too tight during testing, researchers might fail to discover capabilities before the model is released. This is just as dangerous, possibly more so, than giving it too much freedom, and then the evaluation itself risks becoming the problem.
Can safety evaluations be regulated?
The Trump administration is currently weighing a voluntary pre-deployment cybersecurity evaluation regime, under which the government will get to assess the security risks of new, powerful models 30 days before they are released publicly. The policy — the product of a Trump executive order which has been finalized behind closed doors — wouldn’t address safety evaluation incidents because they occur farther upstream of deployment.
“The lesson we’ve been learning in the last few months is that the self-regulatory apparatus is just not enough anymore,” Yoon said. “There are competitive pressures that are incentivizing a race to the bottom on safety standards, and that is a perfect place for regulatory intervention.”
“What we would need to cover this is some kind of controls on what’s happening inside the labs while the models are being developed, both at the training stage and at the testing stage,” he continued.
The challenge is only likely to grow as the models do. A source familiar with Irregular’s evaluations told TechCrunch that more capable models require more complex evaluations, often conducted quickly and at greater scale, which opens the door for more mistakes.
AISI, which intentionally gives some models internet access, told TechCrunch it’s reviewing the balance between realistic testing and managing the risks those tests create.
OpenAI said it’s reviewing how it conducts third-party testing, as well as requirements around isolation, monitoring, and when evaluations should be stopped. Meta said it’s still investigating the incident and plans to publish a retrospective once it has all the facts.
In the end, there may be no way to eliminate risk entirely. As models become more capable, the environments testing them need to become more robust. The consequences of getting that wrong will only continue to grow.
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
Rebecca Bellan
Senior Reporter
Rebecca Bellan is a senior reporter at TechCrunch where she covers the business, policy, and emerging trends shaping artificial intelligence. Her work has also appeared in Forbes, Bloomberg, The Atlantic, The Daily Beast, and other publications.
You can contact or verify outreach from Rebecca by emailing rebecca.bellan@techcrunch.com or via encrypted message at rebeccabellan.491 on Signal.
October 13 – 15
San Francisco
Scale faster. Grow your portfolio. Gain practical expertise. No matter your goal, Disrupt can empower you.
Save up to $300 toda y!
ChatGPT brings unlimited text chats to free users
Tesla and SpaceX will invest $16.8B to start building ‘Terafab’ chip factory in Texas
Amid legal battles, Suno says it will start watermarking songs
Ford’s new electric truck, ‘Fathom,’ starts at $28,350
Bending Spoons to buy Airtable for $1.28B
Influencers draw backlash for attending OpenAI’s first luxury trip
Sequoia’s Shaun Maguire leads $1B round for nuclear startup Valar Atomics
