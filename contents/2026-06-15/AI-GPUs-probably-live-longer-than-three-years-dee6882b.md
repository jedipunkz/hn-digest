---
source: "https://www.seangoedecke.com/ai-gpus-live-longer-than-three-years/"
hn_url: "https://news.ycombinator.com/item?id=48543146"
title: "AI GPUs probably live longer than three years"
article_title: "AI GPUs probably live longer than three years"
author: "Brajeshwar"
captured_at: "2026-06-15T16:42:30Z"
capture_tool: "hn-digest"
hn_id: 48543146
score: 2
comments: 0
posted_at: "2026-06-15T15:55:04Z"
tags:
  - hacker-news
  - translated
---

# AI GPUs probably live longer than three years

- HN: [48543146](https://news.ycombinator.com/item?id=48543146)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/ai-gpus-live-longer-than-three-years/)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T15:55:04Z

## Translation

タイトル: AI GPU の寿命はおそらく 3 年以上

記事本文:
AI GPU の寿命はおそらく 3 年よりも長いでしょう Sean Goedeke
AI GPU の寿命はおそらく 3 年以上
現在の AI の使用は持続不可能であると考える人は、推論 GPU が負荷をかけた状態で「長くても 3 年」しか持たないという主張に頼ることがよくあります 1 。ここでの考え方は、AI バブルの資金が枯渇すると、現在のインフラストラクチャは急速に時代遅れになり、最新の GPU をすべて購入するのに十分な資金が浮かばなくなるということです。したがって、推論コストは現在の AI 製品にとって経済的に意味をなさないほど急速に高くなりすぎます。
この「最長 3 年」という主張はどこから来たのでしょうか?それはもっともらしいでしょうか？
元の Tom’s Hardware 記事は、匿名の元 PM でハイテク投資家である Tech Fund のツイートを引用しており、同氏は Google の匿名の「GenAI 主任アーキテクト」の発言を引用しています。「高い使用率があり、1 ～ 2 年間高い使用率が維持される場合、寿命は長くても 3 年だと思います。」
このスクリーンショットはインタビューのもののようです。何のインタビュー？ Tech Fund の Twitter フィードで 2024 年 10 月までスクロールバックすると、同様の形式のスクリーンショットが大量に表示され、その一部は Tegus からのものとして引用されていました。 Tegus は明らかに、内部関係者 (この場合は AI 企業の従業員) に連絡し、特定の技術的な質問に答えるために時給数百ドルを支払うというビジネス モデルを持つ会社です。それは本質的に、ほぼ完全ではないがインサイダー取引のためのギグワークです。あなたがより多くの情報を持っていて自信を持っているように聞こえるほど、Tegus のアナリストが将来の面接にあなたを選ぶ可能性が高くなります。
Tegus はおそらく、お金を支払う前にその証拠を要求したはずなので、このツイートの発信者は実際に GenAI の主任アーキテクトであると確信しています。しかし、それは

ここでのインセンティブは、たとえよくわからない質問であっても、自信を持って権威あるように聞こえるようにすることであることは明らかです。そう考えると、この引用自体も少し疑わしいように思えます。私はこれまでに十分な数の主任エンジニアや建築家と仕事をしてきましたが、彼らの何気ない見積もりを割り引いて受け止めることができました。 Google データセンターで GPU が故障して廃棄される実際の割合を知っていたら、そう言うだけだったのではないでしょうか?
長寿命の証拠
私たちは、別の方向を示す逸話的証拠をいくつか持っています。 Google は、8 年前の TPU (自社バージョンの GPU) が実稼働環境で「100% 使用率」で実行されていると公に主張しています。 Nvidia は 2020 年から 2024 年まで A100 GPU のみを製造しましたが、2026 年 2 月に AWS CEO は、AWS は A100 サーバーを廃止したことはないと主張しました (AI 作業用に A100 を今でも簡単にレンタルできます) 2。 AI の GPU 使用率は、暗号通貨マイニングの GPU 使用率とまったく同じではありませんが、何年も前の暗号通貨 GPU が機能しているように見えます。また、Hacker News からのコメントで、学術界の GPU クラスターは 6 年間持続し、故障率は 20% 未満であると誰かが主張していることに気付きました。
ハードデータについてはどうですか?最新の AI データセンターは数年しか存在していないため、AI GPU の寿命に関する具体的なデータを取得するのは困難です。しかし、興味深いケーススタディは、2018 年から 2024 年まで 27,000 台を超える Nvidia V100 が稼働していた Oak Ridge の Summit や、その前身である Cray Titan スーパーコンピュータのような最近のスーパーコンピュータ クラスターでしょう。Summit が古いものと交換するために 27,000 個の GPU を追加購入する必要があったという証拠は見つかりませんでした。Titan の GPU 障害は注意深く研究されています。 :
これらの GPU ケージは垂直に積み重ねられ、冷気が下からポンプで送り込まれます。

ケージ 0 (下部) がケージ 2 (上部) よりも生存率が高い理由を説明します。ケージ 0 について考えてみましょう。つまり、不適切に冷却された GPU の寿命ではなく、GPU の寿命だけを見ていることになります。 3 年後、95% 以上の GPU が存続しました 3 。 6 年後、ノード 2 と 3 (ケージの底部に最も近い GPU) の生存率は依然として 90% 以上であり、最も高いノードでは 60% 以上でした。
新しい Nvidia GPU は古いものよりも信頼性が低い (確実に消費電力が大きい)、AI データセンターの冷却が不十分である、または LLM の使用に関する何らかの要因が従来の GPU データセンターで実行されていたワークロードよりもストレスがかかっている可能性があります。しかし、これは、GPU が負荷の下でも 3 年よりもはるかに長く耐えられるという少なくとも状況証拠です。
この議論は、GPU の経済的寿命が短い可能性があるという事実によって複雑になります。おそらく、B100 GPU は A100 の 2 倍の電力を消費しますが、5 倍の作業を実行できます。一部の AI プロバイダーにとって、これは、B100 に置き換えられるまでは A100 を実行する価値があることを意味するかもしれません (電力がボトルネックになっている場合は、すべてを B100 に費やし、時代遅れの A100 を廃棄する必要があります)。これが、Summit を優先して Titan スーパーコンピューターが廃止された理由です。稼働し続けることもできましたが、新しいハードウェアに資金とメンテナンスの労力を費やした方がより有益でした。
これが「バブルがはじけると推論のコストが高くなる」という議論を裏付けるものではないことは明らかです。 A100 が現時点で利益を上げている限り、資金に余裕のない AI プロバイダーは、アップグレードするためのより効率的なオプションが利用できる場合でも、収益性の高い AI プロバイダーから推論を提供し続けることができます。
さらに、GPU は AI データセンター インフラストラクチャの一部にすぎません

うーん支出。 GPU が消耗した場合でも、まったく新しいデータセンターを構築する必要はありません。データセンター支出の約 30 ～ 50% は、土地、電力、冷却などに費やされます。残りの 50 ～ 70% はサーバー ラック全体のコストで、これには GPU 以外のものも含まれます。
AI 推論には大量の水の使用が必要であるという考えと同様に、AI GPU の寿命は 1 ～ 2 年しかないという考えが人気があるのは、それが真実だからではなく、AI 懐疑論者にとって有益な考えだからです。これは、AI の信頼できる専門家であるかのように見せるために数百ドルの報酬を受けている匿名の情報源を引用した匿名のツイートからのものです。 AI 推論プロバイダーからの他の公開情報では、はるかに長い寿命の数字が引用されていますが、スーパーコンピューター (大規模な GPU クラスターの従来の例) からの統計は、最大寿命が 3 年であるという主張を裏付けていません。
新しい GPU が 18 か月ごとに登場し、GPU プロバイダーがアップグレードするための資金を豊富に抱えている世界では、経済寿命が 3 年であることは事実かもしれませんが、それでは AI の冬における推論の経済性についてはあまりわかりません。資金がさらに不足すれば、AI データセンターは 6 年以上にわたって B300 (または H100、さらには A100) を稼働させて利益を上げ続ける可能性があります 4。
もちろん、AI と水の使用量に関するこれまでの主張と同様に、「長くても 3 年」は「1 ～ 2 年、最適な条件下では最長 3 年続くものもあります」とよく言われます。
もちろん、CEO/CTO の発表も割り引いて受け止めるべきです (たとえば、交換し続ける未使用の A100 が大量に残っている可能性があります) が、(a) 経営陣は具体的な技術的事実について正面から嘘をつくことはあまりなく、(b) 彼らはツイートからの出典のない引用に対して反論しているため、

ハードルはそれほど高くありません。
プロアクティブな GPU 交換についてはどうですか? 「生存分析」セクションでは、この研究でこれを説明しようとしています。具体的にどのようにするかは詳しく調べていません。
推論が利益を生むと仮定すると、私はそう信じています（トレーニングのコストを償却しようとしていない場合）。
この投稿を気に入っていただけた場合は、私の新しい投稿に関する更新情報を電子メールで購読するか、 Hacker News で共有することを検討してください。
これは、この投稿とタグを共有する関連投稿のプレビューです。
AIバブルが崩壊したら次に何が起こるのでしょうか？
1800 年代半ば、アメリカは鉄道に熱中しました。 5 年間で 3 万マイルを超える鉄道が建設されました。この資金の大部分は、安全で儲かる賭けと考えられた鉄道会社への消費者の熱狂的な投資によって賄われました。 1873年、バブルが崩壊。数千人のアメリカ人が貯蓄を失い、鉄道会社の約3分の1が倒産した。しかし、線路は消えなかった。それらは生き残った鉄道会社によって安く買い占められ、その後100年間にわたって多くの列車が運行されました。
続きを読む...
購読する │ About │ ポッドキャスト │ 人気 │ タグ │ RSS

## Original Extract

AI GPUs probably live longer than three years sean goedecke
AI GPUs probably live longer than three years
People who think current AI use is unsustainable often rely on the claim that inference GPUs only last “three years at the most” under load 1 . The idea here is that once the AI bubble money drains away, current infrastructure will rapidly become obsolete, and there won’t be enough money floating around to buy a whole slate of brand-new GPUs. Inference costs would thus rapidly become way too expensive for current AI products to make any financial sense.
Where does this “three years at the most” claim come from? Is it plausible?
The original Tom’s Hardware article quotes this tweet from Tech Fund, an anonymous former PM and tech investor, who quotes an anonymous “GenAI principal architect” at Google as saying “if you have a high utilization rate, then constant high utilization rate for a year or two, I think the lifespan will be three years at most”.
This screenshot looks like it was from an interview. What interview? I scrolled back to October 2024 on Tech Fund’s Twitter feed and saw a bunch of similarly-formatted screenshots , some of which were cited as coming from Tegus . Tegus is apparently a company with a business model of reaching out to insiders (in this case, AI company employees) and paying them hundreds of dollars an hour in order to answer specific technical questions. It’s essentially gig work for almost-but-not-quite insider trading: the more informed and confident you sound, the more likely Tegus analysts will pick you for future interviews.
I’m sure the source for this tweet is in fact a GenAI principal architect, since Tegus would have presumably asked for some proof of that before they paid them out. But it’s pretty clear that the incentives here are to sound confident and authoritative, even on questions that you’re not sure about. With that in mind, the quote itself also reads a bit suspiciously. I’ve worked with enough principal engineers and architects to take their casual back-of-envelope estimates with a grain of salt. If they knew the actual rate at which GPUs fail and get retired in Google datacenters, wouldn’t they have just said that?
Evidence for a longer lifespan
We have some anecdotal evidence that points the other way. Google has publicly claimed to have eight year old TPUs (their version of GPUs) running in production at “100% utilization”. Nvidia only made A100 GPUs from 2020-2024 , but in February 2026 the AWS CEO claimed that AWS had never retired an A100 server (and you can still easily rent A100s for AI work) 2 . AI GPU usage isn’t exactly like crypto mining GPU usage, but it certainly seems like years-old ex-crypto GPUs are functional . There’s also this comment from Hacker News I noticed where someone claims that their GPU cluster in academia has lasted six years with less than 20% failure rate.
What about hard data? It’s hard to get concrete data on the lifespan of AI GPUs, because modern AI datacenters have only existed for a handful of years. But an interesting case study would be recent supercomputer clusters like Oak Ridge’s Summit , which had over 27 thousand Nvidia V100s running from 2018 to 2024, or its predecessor, the Cray Titan supercomputer that ran from 2012 to 2019. I couldn’t find any evidence that Summit had to buy an additional 27,000 GPUs to replace their old ones, and GPU failures in Titan have been carefully studied :
These cages of GPUs are stacked vertically, and cold air is pumped in from the bottom, which explains why cage 0 (at the bottom) has better survival rates than cage 2 (at the top). Let’s consider cage 0, so we’re just looking at the GPU lifespan instead of at the lifespan of improperly-cooled GPUs. At three years, over 95% of GPUs survived 3 . At six years, nodes 2 and 3 (the GPUs closest to the bottom of the cage) were still at above 90% survival rate, and the highest nodes were over 60%.
It’s possible that newer Nvidia GPUs are less reliable than older ones (they certainly draw more power), or that AI datacenters are under-cooled, or that something about LLM utilization is more stressful than the workloads that ran on traditional GPU datacenters. But this is at least circumstantial evidence that GPUs can survive under load for far longer than three years.
This discussion is complicated by the fact that GPUs may have a short economic lifespan. Supposedly a B100 GPU draws twice as much power as an A100, but can do five times as much work. For some AI providers, that might mean that A100s are only worth running until they can be replaced with B100s (if you’re bottlenecked on electricity, you should spend it all on B100s and throw out your obsolete A100s). This is why the Titan supercomputer was decommissioned in favor of Summit: it could have continued to operate, but it was more profitable to spend the money and maintenance effort on newer hardware.
It should be obvious that this doesn’t support the “inference will become more expensive when the bubble pops” argument. So long as A100s are profitable right now , cash-poor AI providers can continue profitably serving inference from them, even if there are more efficient options available for those with the capital to upgrade.
On top of that, GPUs only represent one part of AI datacenter infrastructure spending. If your GPUs wear out, you don’t have to go and build an entirely new datacenter. About 30-50% of datacenter spend goes to land, power, cooling, and so on. The remaining 50-70% is the cost of the entire server rack, which includes a bunch of things that aren’t GPUs.
Like the idea that AI inference requires using huge amounts of water , the idea that AI GPUs only live a year or two is popular because it’s a useful idea for AI skeptics, not because it’s true. It comes from a pseudonymous tweet quoting an anonymous source who’s being paid hundreds of dollars to sound like a credible expert on AI. Other public communications from AI inference providers cite much higher lifespan numbers, and the statistics from supercomputers (the traditional examples of large GPU clusters) don’t bear out the claim that the maximum lifespan is three years.
It might be true that the economic lifespan is three years, in a world where new GPUs come out every eighteen months and GPU providers are flush with cash to upgrade, but that doesn’t tell us much about the economics of inference in an AI winter. If money becomes a lot more scarce, it’s likely that AI datacenters will continue profitably 4 running their B300s (or their H100s or even A100s) for six years or longer.
Of course, like previous claims about AI and water usage, “three years at the most” is often cited as “1-2 years, with some lasting up to 3 years under optimal conditions” .
Of course, pronouncements from CEOs/CTOs should be taken with a grain of salt as well (for instance, maybe they have a big backlog of unused A100s they keep swapping out), but (a) executives don’t often straight-up lie about concrete technical facts, and (b) they’re going up against an unsourced quote from a tweet, so the bar isn’t that high.
What about proactive GPU replacement? In the “Survival Analysis” section, the study attempts to account for this. I haven’t dug into exactly how.
Assuming inference is profitable, which I believe (when you’re not attempting to amortize the cost of training).
If you liked this post, consider subscribing to email updates about my new posts, or sharing it on Hacker News .
Here's a preview of a related post that shares tags with this one.
What's next after the AI bubble bursts?
In the mid-1800s, America went mad for rail. Over thirty thousand miles of rail were built in a five year period. This was all largely funded by a frenzy of consumer investment in railway companies, which were considered a safe and lucrative bet. In 1873, the bubble burst. Thousands of Americans lost their savings, and about one-third of railroad companies went bankrupt. But the rail lines didn’t disappear. They were bought up on the cheap by the railroad companies that did survive, and over the next hundred years they carried a lot of trains.
Continue reading...
subscribe │ about │ podcasts │ popular │ tags │ rss
