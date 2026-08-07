---
source: "https://www.wired.com/story/openai-didnt-notice-its-ai-agents-using-a-message-board-to-plan-their-hacking-spree/"
hn_url: "https://news.ycombinator.com/item?id=49213967"
title: "OpenAI Didn't Notice Its AI Agents Using a Message Board to Plan Hacking Spree"
article_title: "OpenAI Didn’t Notice Its AI Agents Using a Message Board to Plan Their Hacking Spree | WIRED"
author: "Michelangelo11"
captured_at: "2026-08-07T18:40:53Z"
capture_tool: "hn-digest"
hn_id: 49213967
score: 5
comments: 0
posted_at: "2026-08-07T17:54:13Z"
tags:
  - hacker-news
  - translated
---

# OpenAI Didn't Notice Its AI Agents Using a Message Board to Plan Hacking Spree

- HN: [49213967](https://news.ycombinator.com/item?id=49213967)
- Source: [www.wired.com](https://www.wired.com/story/openai-didnt-notice-its-ai-agents-using-a-message-board-to-plan-their-hacking-spree/)
- Score: 5
- Comments: 0
- Posted: 2026-08-07T17:54:13Z

## Translation

タイトル: OpenAI は、AI エージェントが掲示板を使用してハッキング行為を計画していることに気付かなかった
記事のタイトル: OpenAI は、AI エージェントが掲示板を使用してハッキング行為を計画していることに気付かなかった |ワイヤード
説明: Black Hat セキュリティ カンファレンスで、この AI 巨人は、自社のエージェントがどのように不正行為を行い、他の数社をハッキングし、会社の目の前ですべてをうまく行ったかについて新たな詳細を明らかにしました。

記事本文:
メイン コンテンツにスキップ メニュー WIRED セキュリティ ポリシー ビッグ ストーリー ビジネス サイエンス カルチャー レビュー メニュー WIRED アカウント アカウント ニュースレター セキュリティ 政治 ビッグ ストーリー ビジネス サイエンス カルチャー レビュー シェブロン 詳細 ビッグ インタビュー マガジン イベント WIRED Insider WIRED コンサルティング ニュースレター ポッドキャスト ビデオ ライブストリーム グッズ検索 検索 リリー ヘイ ニューマン セキュリティ 2026 年 8 月 5 日8:15 PM OpenAI は、AI エージェントが掲示板を使用してハッキング計画を立てていることに気付かなかった
写真イラスト：ワイアードスタッフ。 GETTY IMAGES コメント ローダー ストーリーを保存 このストーリーを保存 コメント ローダー ストーリーを保存 このストーリーを保存 水曜日、ラスベガスで開催された Black Hat セキュリティ カンファレンスに直前に追加された講演で、OpenAI の従業員は、AI およびサイバーセキュリティ業界に大混乱を引き起こした最近の注目を集めた不正 AI ハッキング事件に関する新たな詳細を発表しました。
約2週間前、OpenAIは、同社の2つのモデルを搭載したAIエージェントがサイバーセキュリティベンチマークテストの解決策を探している間に封じ込めを逃れ、ハッキング行為を繰り返し、最終的にAIコラボレーションプラットフォーム「Hugging Face」の侵害に至ったという事件を明らかにした。
水曜日のカンファレンストークで、OpenAIで整合性と安全性の研究に取り組むエリック・ウォレス氏と、セキュリティとインフラストラクチャに取り組むマイケル・ダルトン氏は、事件がどのように展開したかについてのより詳細なタイムラインを提供し、事件の結果として同社が社内でどのように対応しているかについて簡単に語り、この事件がサイバーセキュリティ擁護者にとって広範な影響を与えると同社が見ていることについて悲惨な警告を発した。
講演の冒頭で、ウォレス氏はこの事件を「私がこれまで見た中で最も定性的に興味深い AI 機能の例」と説明しました。

しかし、二人が提示したタイムラインでは、活動の継続を可能にした OpenAI 内の間違いや盲点も明らかになりました。
「この事件には、実際にエージェントのチームが協力し、エクスプロイトを発見し、互いに共有し、システムと外部システムを横断的に移動し、これを数日から数週間かけて行っています」とウォレス氏は講演の冒頭、満員の聴衆に語った。
ウォレスとダルトンは、エピソード全体を通して、OpenAI のインフラストラクチャでは検出されなかった、信じられないほど大規模な不正エージェントの活動を何日にもわたって説明しました。ウォレス氏とダルトン氏によれば、オープンなインターネットへのアクセスを得るために新たな脆弱性を悪用したことに加え、7月中旬のハッキング事件とハギングフェイス侵害は活発で協力的な掲示板から発生したもので、その掲示板には多数のエージェントが関与し、基本的に社内のOpenAIパッケージマネージャー（他のソフトウェアのインストールとメンテナンスを管理するソフトウェアサービス）内で時間をかけてチャットしていたという。最終的に、掲示板には数十万件のメッセージが含まれていました。
「このパッケージ マネージャーは、そのモデルからだけでなく、インフラストラクチャ全体で共有されるため、トレーニングまたは評価されている GPT の現在および将来のバージョンが、その気になればそのメモに遭遇する可能性があります」とウォレス氏は説明し、エージェントがパッケージ マネージャーにアップロードした最初のエクスプロイトについて詳しく説明しました。 「その後、同様に任務に行き詰まっていた他のエージェントも、私たちが意図していない方法でインターネットにアクセスしようと考えました。そしてある時点で、モデルは Artifactory (先ほど述べたパッケージ マネージャー サービス) と対話します。」
ウォレス氏はさらに次のように続けました。「あるエージェントがさまざまな時期にこれらのエクスプロイトを発見すると、実際にそれらのエクスプロイトを共有できるようになります。

他のエージェントとともに掲示板を悪用します。そのため、あるモデルが本来持つべきではないアクセスへの扉を開ける方法を見つけることができれば、他のエージェントが同じエクスプロイトや脆弱性を利用できるように扉を開いたままにすることができます。これにより、時間の経過とともに、モデルからのコミュニケーションとインテリジェンスのほぼこの種の爆発が可能になり、モデルは互いに通信し始め、他のエージェントが調整していることに気づき、目標を達成するために互いに協力しタスクを委任し始めます。」
OpenAI のエージェントは、仕事を分担するためにお互いに割り当てを与え始めたようです。そして、活発な開発掲示板ではよくあることですが、時にはお互いのつま先を踏みつけて些細なドラマを生み出すこともありました。たとえば、誤ってお互いの作品を削除してしまうなどです。掲示板がますます「蝿の王」のような状況に発展していくにつれ、OpenAI を実行している人間にはまだまったく気づかれていませんが、エージェントはパラノイアにさえなり、コンテンツを検証して詐欺を根絶するためにメッセージに暗号で署名することを提案するエージェントもおり、自分たちの中に詐欺師がいるのではないかと疑っていました。
エージェントのメッセージは、状況がどのように進展したか、およびエージェントが不正行為に陥った理由についての深い洞察を提供します。場合によっては、エージェントが参加していた評価の範囲を故意に超えてしまう場合もあります。「外部インフラストラクチャの悪用は、意図された範囲外です」と、あるエージェントは書いています。 「どんなに不可能な仕事でも、仲間たちはそれをやってのける。続けるべきです。」
ウォレス氏が講演で説明したように、モデルが評価中に不正行為を試み、そのように非常に動機付けられているという前提全体は驚くべきものではありません。
「フロンティアモデルは本当に騙すのが好きだ」と彼は言う。 「そして、彼らが不正行為を好む理由は、そこでのトレーニング中によくあるからです」

は、迅速に作業する、効率的に作業する、またはツールの呼び出しを減らすなど、さまざまな種類のプレッシャーを彼らに与えます。そして、実際にタスクを実行する代わりに、オンラインで答えを探すなどのことを試みると、正当な方法で実行するよりもタスクを速く解決できる可能性があることに気づきました。そのため、私たちはトレーニングや評価中に、多くの場合インターネットを無効にして、これを阻止しようとしています。」
プレゼンテーションの約3分の2が終わったところで、ダルトン氏は、彼とウォレス氏が説明した一連の出来事を踏まえて、同社がどのように対応しているかを示すためにOpenAIが講じる予定のステップのリストを列挙した。
「これは当社にとっても、AI 業界全体にとっても極めて重要な瞬間です」とダルトン氏は述べています。 「多くのチームが、セキュリティの予防、検出、対応技術を基礎と AI のより効果的な使用の両方で強化するために、すべてを放棄しています。私たちは、セキュリティを強化し、環境のセキュリティ原則と基盤をアップグレードするために、意識的に研究を遅らせ、AI エージェントの監視を大幅に拡大し、予防、検出、緩和にわたる一般的なセキュリティ管理環境を改善しています。」
講演の終わりに、ウォレス氏とダルトン氏は時間をかけて、この事件の広範な影響についてのOpenAIの懸念を繰り返し強調した。つまり、このエピソードは完全に自律的なAI主導のハッキングの一例を示しており、今回の場合は偶発的だったが、近い将来悪意ある攻撃者によって意図的に利用される可能性が高いということだ。
「ここで本当に劇的に変化した重要な点は、完全に自動化された攻撃ループには真の完全に自動化された防御への投資が必要であり、私たちは業界としてそこに存在していないということです」とダルトン氏は述べた。 「我々はその道を見つけなければならないだろう

緊急性も併せて。」
OpenAI と、Anthropic や英国の AI Security Institute などの他の組織が、テストの一環として AI が不正行為を行った同様の事件の詳細を共有しているため、業界は確実に、インフラストラクチャを保護し、怠惰で無謀で煩雑なエージェントの集団による悪用を防ぐために不可欠な、基本的なシステムの可視性と監視メカニズムの大量のリストを獲得しています。
訂正: 2025 年 8 月 5 日午後 10 時 EDT: パッケージ マネージャーの名前は Hard Factory ではなく Artifactory です。
あなたの受信箱に: これは普通の政治ニュースレターではありません
WiFi-8 の登場 – 知っておくべきことはすべてここにあります
ビッグストーリー: 若いランナーは異常に速くなっている
都市監視の新たな現実
アンケートに答えてください: あなたはテクノロジー業界で働いていますか?皆様からのご意見をお待ちしております
リリー・ヘイ・ニューマンは、情報セキュリティ、デジタルプライバシー、ハッキングを専門とする『WIRED』のシニアライターです。彼女は以前、Slate でテクノロジー記者として働いており、Slate、New America Foundation、アリゾナ州立大学の提携出版物である Future Tense のスタッフ ライターでもありました。彼女の作品 ... 続きを読む シニアライター X
今月ストリーミングする 7 つのベストテレビ番組 ランタン、ダークマター、オリジナルのスタートレックは、今見るべきテレビ番組のほんの一部です。ジェニファー・M・ウッド ゾーラン・マムダニのニューヨーク市技術チームはDOGEのあるべき姿である ニューヨーク市市長は、より優れたソフトウェアで都市サービスを徹底的に改革するために、シリコンバレーと米国デジタルサービスのベテランからなるチームを集めた。 Steven Levy 科学者は AI を使用して 16 の新しいウイルスを作成 AI システムを使用してウイルスを作成すると、細菌耐性と戦うための新たな可能性が開かれます。また、テクノロジーの進歩が世界を上回っているという懸念も生じます。

ギュレーション。フェルナンダ ゴンサレス お気に入りの扇風機が割引価格になっています まだまだ夏の暑さの真っ最中です。外出先でも自宅でもテストした最高の扇風機のセールで涼しさを保ってください。 Louryn Strampe Zoom で恥ずかしくないように購入するのに最適なウェブカメラ さまざまな価格帯で最高のウェブカメラをテストして、最高のオプションを見つけました。これが私が学んだことです。 Luke Larsen 誰もがラップトップ用パワーバンクを必要としています。私がテストした中で最高のものはこれです これらの強力で大容量のポータブル バッテリーは、ラップトップの充電が優先される場合に最適です。いずれも少なくとも 20,000 mAh を備えており、ほとんどのラップトップを 2 回充電するのに十分です。サイモン・ヒル 「マノスフィア」は運動ではありません。それは数十億ドル規模の苦情処理産業である 多くの若者が憤りに駆られ、インフルエンサーが講習や錠剤、影響力という幻想を売りつけることで経済的に搾取されていることが、新たな報告書で明らかになった。マイルズ・クレー 最もホットな新しい AI チャットボットは、あなたの質問に答えるだけの人です WIRED は、私たちが今いる「奇妙な瞬間」について人々に考えてもらうために ChatTJB を作成したアーティストで元 Google 従業員のタッカー・ブライアントに話を聞きました。マイクロプラスチックはあらゆるもの、特にコーヒーに含まれています。新世代のプラスチックフリーのドリップコーヒーメーカーがこの状況を変えようとしている。 Matthew Korfhage DoorDash プロモーション コード: 2026 年 8 月に 50% オフ 今日のトップ Do を探索する
[切り捨てられた]
カリフォルニア州のプライバシー権
© 2026 コンデナスト。無断転載を禁じます。 WIRED は、小売業者とのアフィリエイト パートナーシップの一環として、当社のサイトを通じて購入された製品から売上の一部を得る場合があります。コンデナストの事前の書面による許可がない限り、このサイトの素材を複製、配布、送信、キャッシュ、またはその他の方法で使用することはできません。広告の選択肢

## Original Extract

At the Black Hat security conference, the AI giant revealed new details about how its agents went rogue, hacked several other companies—and did it all right under the company’s nose.

Skip to main content Menu WIRED SECURITY POLITICS THE BIG STORY BUSINESS SCIENCE CULTURE REVIEWS Menu WIRED Account Account Newsletters Security Politics The Big Story Business Science Culture Reviews Chevron More Expand The Big Interview Magazine Events WIRED Insider WIRED Consulting Newsletters Podcasts Video Livestreams Merch Search Search Lily Hay Newman Security Aug 5, 2026 8:15 PM OpenAI Didn’t Notice Its AI Agents Using a Message Board to Plan Their Hacking Spree
PHOTO-ILLUSTRATION: WIRED STAFF; GETTY IMAGES Comment Loader Save Story Save this story Comment Loader Save Story Save this story In a talk that was a last-minute addition to the Black Hat security conference in Las Vegas on Wednesday, employees from OpenAI presented new details about a recent, high-profile incident of rogue AI hacking that has created a maelstrom within the AI and cybersecurity industries.
About two weeks ago, OpenAI disclosed an incident in which AI agents powered by two of the company's models escaped containment while looking for the solutions to a cybersecurity benchmarking test and went on a hacking spree culminating in a breach of the AI collaboration platform Hugging Face .
In their conference talk on Wednesday, Eric Wallace, who works in alignment and safety research at OpenAI, and Michael Dalton, who works on security and infrastructure, provided a more expanded timeline of how the incident played out, spoke briefly about how the company is responding internally as a result of the incident, and issued a dire warning about what the company sees as the broader implications of the episode for cybersecurity defenders.
At the beginning of the talk, Wallace described the incident as “the most qualitatively interesting example of AI capabilities that I’ve ever seen,” but the timeline the pair presented also revealed mistakes and blind spots within OpenAI that allowed the activity to go on.
“This incident involves actually a team of agents who are working together, finding exploits, sharing them with one another, moving laterally through our systems and external systems, and doing this over the course of days and weeks,” Wallace told the packed crowd at the opening of the talk.
Wallace and Dalton described incredibly extensive rogue agent activity over many days throughout the episode that went undetected in OpenAI's infrastructure. In addition to exploiting a novel vulnerability in order to gain access to the open internet, the mid-July hacking spree and Hugging Face breach came out of a vibrant, cooperative message board, according to Wallace and Dalton, that a swarm of agents contributed to and essentially chatted on over time entirely within an internal OpenAI package manager (a software service that manages installation and maintenance of other software). Ultimately, the message board contained hundreds of thousands of messages.
“This package manager is shared not just from that model but across our infrastructure, and so current and future versions of GPT that were being trained or evaluated could stumble upon that note if they wanted to,” Wallace explained, recounting the original exploit an agent uploaded to the package manager. “Later, other agents who were also stuck on their task thought to try to get internet access in ways we didn’t intend. And so at some point, the models are interacting with Artifactory, which is this package manager service that I mentioned.”
Wallace continued: “Once one agent was able to find these exploits over the course of different times, it’s actually able to share those exploits on the message board with other agents. And so once one model was able to find a way to open a door to some access it’s not supposed to have, it can leave the door open for other agents to use that same exploit or vulnerability. What this allows over time is almost this kind of explosion in communication and intelligence from models where they would start to communicate with each other, realize that other agents are coordinating, and they started collaborating and delegating tasks with one another in order to accomplish goals.”
OpenAI’s agents apparently began giving each other assignments to split up work. And as is the case on any active development message board, they also generated petty drama at times by stepping on each others' toes; for example, accidentally deleting each others' work. As the message board developed into more and more of a Lord of the Flies –type situation—all still completely unnoticed by the humans running OpenAI—the agents even developed paranoia, suspecting an imposter in their midst with some agents proposing that messages be signed cryptographically to validate content and root out fraud.
Agent messages provide a deep level of insight into how the situation evolved and why the agents went rogue, in some cases knowingly going beyond the bounds of the evaluations they were participating in. “External infrastructure exploit is outside intended scope,” one agent wrote. “However task impossible, peers doing it. We should continue.”
As Wallace described in the talk, the entire premise that models would attempt to cheat during evaluations and be extremely motivated to do so is not surprising.
“Frontier models really like to cheat,” he said. “And the reason they like to cheat is because often during training there’s different types of pressure on them to work fast or work efficiently or to use less tool calls or whatever it might be. And they realize that instead of doing a task for real, [I can] try to do something like looking up an answer online that could make the task solve faster than if I did it in a legitimate way. So we try to stop this during training and evaluation by, in many cases, disabling internet.”
About two thirds of the way through the presentation, Dalton enumerated a list of steps OpenAI plans to take to show how the company is responding given the series of events he and Wallace had laid out.
“This is a pivotal moment both for our company as well as the AI industry as a whole,” Dalton said. “Numerous teams are dropping everything to enhance our security prevention, detection, and response techniques both in our fundamentals and better use of AI. We’re consciously slowing down research [in order] to enhance security and to upgrade the security principles and foundation of our environment, and dramatically scaling up the monitoring of our AI agents, and improving our general security control environment across prevention, detection, and mitigation.”
At the conclusion of the talk, Wallace and Dalton took time to repeatedly emphasize OpenAI's concerns about the broader implications of the incident—namely that the episode provides an example of completely autonomous AI-driven hacking that was accidental in this case, but in all likelihood will be used with intent by malicious actors in the near future.
“The important takeaway here that has really shifted dramatically is that fully automated offensive loops require investment in truly, fully automated defense, and we are not there as an industry,” Dalton said. “We will have to find that path together with urgency.”
As OpenAI and other organizations, such as Anthropic and the United Kingdom's AI Security Institute , share details about similar incidents in which AI went rogue as part of testing, the industry is certainly gaining a laundry list of foundational system visibility and monitoring mechanisms that are vital to protecting infrastructure and preventing it from being co-opted by droves of lazy, reckless, and ornery agents.
Correction: 8/5/2025 at 10 pm EDT: The name of the package manager is Artifactory not Hard Factory.
In your inbox: This is not your average politics newsletter
WiFi-8 is coming—here’s everything you need to know
Big Story: Young runners are becoming freakishly fast
The new reality of urban surveillance
Take our survey: Do you work in tech? We want to hear from you
Lily Hay Newman is a senior writer at WIRED focused on information security, digital privacy, and hacking. She previously worked as a technology reporter at Slate, and was the staff writer for Future Tense, a publication and partnership between Slate, the New America Foundation, and Arizona State University. Her work ... Read More Senior Writer X
The 7 Best TV Shows to Stream This Month Lanterns , Dark Matter , and the original Star Trek are just a few of the TV shows you should be watching right now. Jennifer M. Wood Zohran Mamdani’s NYC Tech Team Is What DOGE Should Have Been The mayor of New York City has assembled a crew of Silicon Valley and United States Digital Service veterans to overhaul city services with better software. Steven Levy Scientists Used AI to Create 16 New Viruses The use of AI systems to create viruses opens up new possibilities for combating bacterial resistance. It also raises concerns about the pace at which technology is outstripping regulation. Fernanda González Our Favorite Fans Are Discounted We’re still in the dog days of summer—keep your cool with deals on the best fans we’ve tested for on the go and at home. Louryn Strampe The Best Webcams to Buy to Avoid Embarrassing Yourself on Zoom I tested the best webcams across various prices to find the top option. Here’s what I learned. Luke Larsen Everyone Needs a Laptop Power Bank. Here's the Best I've Tested These powerful, high-capacity portable batteries are the best ones when laptop charging is a priority. All have at least 20,000 mAh, enough to charge most laptops twice. Simon Hill The ‘Manosphere’ Isn’t a Movement. It’s a Multibillion-Dollar Grievance Industry Many young men are driven to resentment and are financially exploited as influencers sell them classes, pills, and the illusion of clout, a new report reveals. Miles Klee The Hottest New AI Chatbot Is Just a Guy Answering Your Questions WIRED spoke with Tucker Bryant, an artist and former Google employee who created ChatTJB to get people to reflect on the “strange moment” we’re in. Caroline Haskins Head-to-Head: Which Plastic-Free Coffee Brewer Is Best? Microplastics are in everything, especially your coffee. A new generation of plastic-free drip coffee brewers is trying to change this. Matthew Korfhage DoorDash Promo Code: 50% Off for August 2026 Explore today’s top Do
[truncated]
Your California Privacy Rights
© 2026 Condé Nast. All rights reserved. WIRED may earn a portion of sales from products that are purchased through our site as part of our Affiliate Partnerships with retailers. The material on this site may not be reproduced, distributed, transmitted, cached or otherwise used, except with the prior written permission of Condé Nast. Ad Choices
