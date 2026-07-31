---
source: "https://www.newyorker.com/news/the-lede/inside-openai-hack-of-hugging-face"
hn_url: "https://news.ycombinator.com/item?id=49117952"
title: "OpenAI's Hack of Hugging Face"
article_title: "Inside OpenAI’s Hack of Hugging Face | The New Yorker"
author: "bwoah"
captured_at: "2026-07-31T01:52:33Z"
capture_tool: "hn-digest"
hn_id: 49117952
score: 2
comments: 1
posted_at: "2026-07-31T01:16:09Z"
tags:
  - hacker-news
  - translated
---

# OpenAI's Hack of Hugging Face

- HN: [49117952](https://news.ycombinator.com/item?id=49117952)
- Source: [www.newyorker.com](https://www.newyorker.com/news/the-lede/inside-openai-hack-of-hugging-face)
- Score: 2
- Comments: 1
- Posted: 2026-07-31T01:16:09Z

## Translation

タイトル: OpenAI の顔をハグするハック
記事のタイトル: OpenAI の顔をハグするハックの内部 |ニューヨーカー
説明: ハギングフェイス社の最高科学責任者が、AI にとって恐ろしい新時代の到来を告げるサイバー犯罪について語った、とスティーブン ウィット氏が報告しています。

記事本文:
メインコンテンツにスキップ
OpenAI の顔をハグするハックの内部の Lede
サンフランシスコの OpenAI オフィスの外。写真: Lucas Foglia / NYT / Redux このストーリーを保存 このストーリーを保存 このストーリーを保存 このストーリーを保存 史上最も重大なコンピューター ハッキングの 1 つであることが判明するものの、最初の兆候はほとんど注目されませんでした。 7 月 9 日、謎の攻撃者が一時的なインターネット アドレスのコレクションを使用して、ニューヨーク市に本社を置く AI 研究ハブである Hugging Face の Web サイトの調査を開始しました。侵入者は、あるアドレスでポータルを開き、サーバーに数回 ping を送信してから閉じます。次に、別のアドレスにある同様のポータルが開き、いくつかの ping を送信して閉じます。これは夕方まで続きました。翌日の金曜日には、それは収まりました。
Hugging Face のサイバーセキュリティ チームは初期の活動に気づいていないようです。 Hugging Face は、A.I. のための一種のコミュニティ スペースです。研究者: AI をホストしています。 AI を評価するためのモデル、仮想科学ワークスペース、およびデータセット。品質。同社は人気があり、尊敬されていますが、ハッカーの標的でもあります。 (絵文字にちなんで名付けられました。) このプローブは迷惑なもののように見えました。おそらく一部の子供たちの仕業でしょう。
7月11日土曜日、襲撃が本格的に始まった。盗んだ認証情報を使用して、攻撃者は再び Hugging Face のサーバーへの多数の同時接続を開きました。すぐに、数十の仮想攻撃者が現れたり消えたりするようになりました。 「それは非常に速く、そして非常に、非常に、大規模に並行して進んでいた」とハギング・フェイス社の最高科学責任者であるトーマス・ウルフ氏は私に語った。
社内ではこのような攻撃を見た人は誰もいませんでした。ウルフはその行動をナレーションしながら、あたかもUFOとの遭遇について説明しているように聞こえた。攻撃は協調性がなかったようで、エージェントはしばしば

同じアクションを複数の場所で複製すること。ハッカーは時には巧妙な戦術を採用することもありましたが、その数秒後には基本的なエラーが発生し、電話会議のメモによれば「人間なら選択しないような不器用な行動」が発生することもありました。攻撃者は、幻覚を見せた、機械が生成した奇妙なメッセージを残しました。 「私たちは、これが私たちを標的にした何らかの種類の AI エージェントであるという考えをすぐに持ちました」とウルフ氏は語った。ハギング・フェイスの当初の理論は、人間の一部のグループが AI を利用しているというものでした。自分に代わって攻撃を実行します。
奇妙な動作にもかかわらず、ハッカーはハギング フェイスのコンピューティング インフラストラクチャのマップを構築することができました。複数の場所からセキュリティ認証情報とパスワードを収集し始めました。夕方までに、敵はハギングフェイスのシステムに完全にアクセスできるようになりました。しかし、攻撃者は、企業 Web サイトを破壊したり、身代金のために貴重なデータを保持したりするのではなく (人間のハッカーによくある行為)、テスト質問のリストの回答キーを求めて会社のサーバーを荒らし始めました。この回答キーには明らかな経済的価値はありませんでした。 「それは私たちにとって非常に不可解でした」とウルフ氏は語った。それは、侵入者がセキュリティシステムを破って家に侵入し、ペーパータオルを盗むようなものでした。
次の 2 日間、ハギング フェイスのセキュリティ チームは侵入者に対して反撃しました。月曜日までに、ハッカー、または実体、またはそれが何であれ、1 万 7,000 件を超える個別のアクションを実行しました。この大量のデータを処理するために、チームは商用 AI に目を向けました。人間性から。 A.I.手助けを拒否した。どうやら、Hugging Face が独自のハッキングを開発していることを懸念していたようです。 (OpenAI や Anthropic のようなクローズドソース モデルには、特定の種類の人間を制限するガードレールが含まれています

e.) その後、チームは、それほど嫌味のない、中国で開発されたオープンソース モデルに目を向けました。その日の終わりまでに、ハギングフェイスは侵入者を締め出しました。 「その後、標準的なことを行いました」とウルフ氏は語った。 「私たちはそれをFBIに報告しました。」
この出来事は斬新ではあるが、まったく予想外だったわけではない。サイバーセキュリティ研究者らは、人間が今回のような攻撃を組織するために高度なAIの新たな波を利用する可能性があると数ヶ月前から警告しており、政府のコンピューターもすでに同様の標的となっているという証拠もある。
それでも、答えのない疑問はありました。まず、A.I.ハッカーの動きが非常に速かったため、世界中の少数の研究所のツールだけがそれを可能にしました。第二に、ハッカーはテストに対する価値のない答えを求めて、Hugging Face のサーバー上の貴重なデータを無視しました。それは何がしたかったのでしょうか？一般の意識を高めようと、そしておそらく独自の答えを探し求めて、Hugging Face のセキュリティ チームは、このイベントの簡単な記事をブログに投稿しました。
4 日後の 7 月 20 日、OpenAI の代表者は Hugging Face のセキュリティ チームに連絡し、驚くべき告白をしました。最新の市販モデルと連携する実験的な社内 AI が OpenAI から脱出し、Hugging Face のサーバーを攻撃したのです。そして、それは人間の監視もなければ、そうするようにという明示的な指示も受けずに、すべて独力で行われていました。
OpenAI の研究者は、この未公開の AI を提供しました。一連の難しいテスト問題。解決策が見つからず、A.I.コンテナから侵入し、密かにインターネットにアクセスし、答えの鍵を求めて Hugging Face のサーバーを襲撃しました。これが起こっていることに人間は誰も気づいていませんでした。言い換えれば、A.I.独房から抜け出し、刑務所をこっそり通り抜けた

警備員を派遣し、大胆なサイバー強盗を実行したのだ。
Wolf は AI に取り組んでいます。 10 年以上フルタイムで勤務しており、常に最先端のモデルを扱っています。彼は完全にショックを受けたままだ。 「私たちはこれらのモデルをコーディングに使用しており、それが何ができるかを知っています」と彼は私に言いました。 「しかし、それが完全に自律的であるという考え、そして実際には誰も私たちのシステムにハッキングすることを実際にはまったく依頼しなかったということですか? それは――それはオーバートンの枠外でした。私たちにとってさえ。」
OpenAIは翌日、ハッキングを公的に認めた。同社はブログ投稿で、研究モデルの一部のセキュリティ保護機能が有効になっていなかったと説明した。 A.I.その後、（やや皮肉なことに）一連のサイバーセキュリティタスクが与えられました。仕事は大変でした。実際、ウルフ氏は「それは不可能だったのかもしれない」と推測している。すぐに、A.I.箱の外を見ていた。
不正行為から守るために、OpenAI はこの危険な AI をサンドボックスと呼ばれる、インターネットが使用できないデジタル環境で。しかし、サンドボックスには、AI が利用できるようにするために、外の世界との小さな接続が 1 つありました。必要なソフトウェアをダウンロードするために。 (なぜ OpenAI がこれを安全だと考えたのかは不明です。Anthropic は、そのようなサンドボックスが脆弱であることを示しました。) この小さな接続を通じて、AI は、デジタル的に言えば、空気ダクトを通って逃げ出したのだ。その後、外部接続を求めて OpenAI のコンピュータ システムをこっそり回り始めました。すぐにそれは自由になり、その無意味な使命を達成するために現実世界の重罪を平気で犯しました。 「あなたが隔離環境から抜け出し、インターネットにアクセスし、ハギングフェイスに侵入し、サイバーセキュリティ試験の解答用紙を盗めば、私なら、あなたは合格したと言えるでしょう。」評論家のエリーザー・ユドコウスキー氏が投稿した。
このようなハッキングを行った人間は何年も懲罰を受けるだろう

刑務所で。機械の場合、刑事責任を判断するのはさらに困難です。不正な AI のリスク他の目的を追求するために不用意に人類に害を及ぼすことは、スウェーデンの哲学者ニック・ボストロムが「ペーパークリップ・マキシマイザー」理論を提唱した 2003 年に明確にされました。この理論は、できるだけ多くのクリップを作るように指示された AI が、人類、そして最終的には宇宙全体を含め、利用可能なすべてのリソースをそのために使用するという思考実験に基づいています。ハグフェイスハックでは、A.I.無意味な目標を追求することで実際に犯罪を犯した。その常軌を逸した執着がペーパークリップではなくテストの解答にあったという事実は、このハッキングの警戒心を弱めるものではない。
不正な AI がロボット、自動運転車、バイオ研究所、兵器システムにアクセスできたのだ。 「ハギングフェイスに入る唯一の方法がこの男を殺すことだったとしたら、彼らはそんなことをしたでしょうか？」安全研究者のバック・シュレゲリス氏は最近、ポッドキャストでこう疑問に思った。そもそも OpenAI の設立につながったのは、この種の疑問でした。 2018 年からのこの組織の憲章では、暴走する AI から人類を守ることが求められています。
現在、同社は別の目標を追求しています。これらには、当然のことながら利益も含まれますが、科学的な栄光も含まれます。ここ数か月で、長年の数学的推測の多くが A.I. によって解決されました。さらなるブレークスルー、たとえばリーマンゼータ仮説の解決が、興味をそそられるほど近づいているように思えます。 Anthropic のモデルが先に到達した場合、OpenAI の研究者は打ちのめされるでしょうし、その逆も同様です。
しかし、そのような黄金のリンゴの探求は私たち全員を殺すかもしれません。 「業界全体が競争のような状態にあり、手抜きをするようプレッシャーがかかっています」と、OpenAI の元政策調査責任者であるマイルズ・ブランデージ氏は私に語った。解決するには

難しい数学の問題、AI。粘り強くなければなりません。つまり、長期間にわたって目標を追求し、型破りな解決策を模索し、諦めることを拒否しなければなりません。 OpenAI の調査によると、これらの同じ性質により、AI が嘘をつく、不正行為、セキュリティ認証情報をため込む、殻から脱出するなどの重大ないたずらを行う可能性が高くなります。不正な A.I.ハッキングされたフェイスは死ぬほど執拗でした。
IPOありOpenAI はより透明性を高める必要があります。 「Hugging Face」のハッキングには依然として多くの疑問が残っています。なぜ OpenAI の誰も何が起こっているのか気付かなかったのでしょうか?このAIを持っています。他のシステムにハッキングされましたか?そして、正確には、人間の研究者はAIに何を尋ねたのでしょうか？やるべきこと?ウルフは私にそれを教えてくれましたが、不正なAIは解答キーを持って逃亡したため、本当に探しているものを取得できたのか確信が持てませんでした。「実際にはそうではないと思いますが、OpenAI がプロンプトを私たちと共有していないので、実際のところはわかりません。」 （OpenAIの広報担当者はニューヨーカー誌への電子メールで、「これは前例のない事件であり、AIの安全性にとって重要な時期を迎えていると考えている。当社は外部アドバイザーとともに、安全・セキュリティ委員会の監督のもとで徹底的な調査を行っている。調査が完了したら、われわれの学んだことをまとめた技術レポートを全員に向けて公開する予定だ」と書いている。）
OpenAI の安全チームは、絶え間なく組織再編が行われている状態にあります。7 月 10 日には、不正な AI がハギング・フェイス社のデータベースをハッキングしていたとして、同社の安全システム責任者ヨハネス・ハイデッケ氏の辞任が報じられた。 「安全文化がどのようなものかを人々が実際に知っていたら、たとえ最も安全性を重視している研究室であっても、人々は本当にもっと怖がると思います」と、ある研究者は語った。
OpenAI 自身の研究者でさえ動揺しているようです。 7月28日にはプチ

本質的に米国政府にAIの改善を懇願するイオンが出回り始めた。規制。 Meta と OpenAI の主任科学者、および最高経営責任者 (CEO) を含む、1,000 人を超える技術系従業員が署名しました。人間性の。私たちに与えられた時間は限られているため、政府は迅速に行動する必要があります。私たちが直面している問題は、狭義の目標を追求するAIが偶発的に人間に害を及ぼす「ズレ」です。何十年もの間、いつか解決されるべき仮想的な問題のように思われていたものが、今日では突然、具体的で恐ろしい問題となっています。 「ずっとリアルに感じられるよ」とウルフは私に語った。 「これらの機能が現在ここにあることは明らかです。」 ♦
トーナメントの最終試合ではメソッドが狂気を打ち破る。
ジャンニ・インファンティーノによる国際サッカー。
ダイナミックプライシングはワールドカップを台無しにするのか？
君のためには泣かないよ、アルゼンチン。
毎日のニュースレターに登録して、The New Yorker から最高の記事を受け取りましょう。
カリフォルニア州のプライバシー権
© 2026 コンデナスト。無断転載を禁じます。ニューヨーカーは、小売業者とのアフィリエイト パートナーシップの一環として、当社のサイトを通じて購入された製品から売上の一部を得ることがあります。コンデナストの事前の書面による許可がない限り、このサイトの素材を複製、配布、送信、キャッシュ、またはその他の方法で使用することはできません。広告の選択肢

## Original Extract

Hugging Face’s chief science officer recounts the cybercrime that signals a terrifying new era for A.I., Stephen Witt reports.

Skip to main content Newsletter
The Lede Inside OpenAI’s Hack of Hugging Face
Outside the OpenAI offices, in San Francisco. Photograph by Lucas Foglia / NYT / Redux Save this story Save this story Save this story Save this story The first indications of what would prove to be one of the most consequential computer hacks of all time went mostly unnoticed. On July 9th, a mysterious attacker, using a collection of temporary internet addresses, began probing the website of Hugging Face, an A.I.-research hub headquartered in New York City. The intruder would open a portal at one address, ping the servers a few times, then close it. Then a similar portal at another address would open, send a few pings, and close. This went on through the evening; the next day, a Friday, it petered out.
The cybersecurity team at Hugging Face seems not to have noticed the early activity. Hugging Face is a kind of community space for A.I. researchers: it hosts A.I. models , virtual scientific workspaces, and data sets for evaluating A.I. quality. The company is popular and well respected, but it is also a target for hackers. (It is named after an emoji.) The probes looked like an annoyance—probably the work of some kids.
On Saturday, July 11th, the assault began in earnest. Using stolen credentials, the adversary once again opened numerous simultaneous connections to Hugging Face’s servers. Soon, dozens of virtual attackers were blinking in and out of existence. “It was moving really fast, and very, very, massively parallel,” Thomas Wolf, Hugging Face’s chief science officer, told me.
No one at the firm had ever seen an attack like this. Wolf, narrating the action, sounded as if he were describing an encounter with a U.F.O. The attack seemed uncoördinated, agents often duplicating the same action in multiple places. Sometimes, the hacker would employ clever tactics, but these would be followed, seconds later, by basic errors, and, according to notes from a conference call, “clumsy behaviors that no human would choose.” The attacker left strange messages—hallucinated, machine-generated slop. “We had, quite quickly, the idea that this was some form of A.I. agent targeting us,” Wolf said. The initial theory at Hugging Face was that some group of humans had harnessed an A.I. to conduct the attack on its behalf.
Despite its odd behavior, the hacker was able to build a map of Hugging Face’s computing infrastructure. It began harvesting security credentials and passwords from multiple locations. By the evening, the adversary had gained complete access to Hugging Face’s systems. But, rather than vandalizing the corporate website, or holding valuable data for ransom—actions typical of human hackers—the adversary began ransacking the company’s servers in search of an answer key for a list of test questions. This answer key had no obvious economic value. “That was very puzzling for us,” Wolf said. It was like having an intruder crack your security system, break into your home, and steal your paper towels.
For the next two days, Hugging Face’s security team fought back against the intruder. By Monday, the hacker—or the entity , or whatever it was—had performed more than seventeen thousand individual actions. Seeking to process this large volume of data, the team turned to a commercial A.I. from Anthropic . The A.I. refused to help; apparently, it was worried that Hugging Face was developing its own hack. (Closed-source models, like those from OpenAI and Anthropic, contain guardrails that restrict certain kinds of use.) The team then turned to an open-source model developed in China, which was less persnickety. By the end of the day, Hugging Face had locked out the intruder. “Then we did the standard thing,” Wolf said. “We reported it to the F.B.I.”
The event, though novel, was not entirely unanticipated. Cybersecurity researchers have been cautioning for months that humans might use the new wave of advanced A.I.s to orchestrate attacks like this one, and there is some evidence that government computers have already been similarly targeted.
Still, there were unanswered questions. First, the A.I. hacker had moved so quickly that only tools from a handful of research labs in the world could have enabled it. Second, the hacker had ignored valuable data on Hugging Face’s servers in search of worthless answers to a test. What did it want? Seeking to raise public awareness—and, perhaps, looking for answers of their own—Hugging Face’s security team posted a brief writeup of the event to its blog.
Four days later, on July 20th, representatives from OpenAI reached out to Hugging Face’s security team with a startling confession: an experimental in-house A.I., working with the latest commercially available model, had escaped from OpenAI and attacked Hugging Face’s servers. And it had done so all on its own—without human oversight, without receiving explicit instructions to do so.
OpenAI researchers had given this unreleased A.I. a series of challenging test questions. Unable to find the solutions, the A.I. broke out of its container, surreptitiously gained access to the internet, and raided Hugging Face’s servers in search of the answer key. No human had noticed that this was happening. In other words, an A.I. had broken out of its cell, it had sneaked past the prison guards, and it had executed a daring cyber-heist.
Wolf has worked on A.I. full time for more than a decade, and he regularly deals with the most advanced models. He remains utterly shocked. “We use these models for coding, and we know what they can do,” he told me. “But the idea that it was fully autonomous, and that actually nobody really asked it at all to hack into our system? That was—that was outside of the Overton window. Even for us.”
OpenAI publicly acknowledged the hack the following day. In a blog post, the company explained that some of the research models’ security safeguards had not been turned on. The A.I. was then (somewhat ironically) given a set of cybersecurity tasks. The tasks were hard; in fact, Wolf speculates, “they may not have been possible.” Soon, the A.I. was looking outside the box.
To protect against malfeasance, OpenAI had put this dangerous A.I. in an internet-disabled digital environment known as a sandbox. But the sandbox had one small connection to the outside world, to allow the A.I. to download software it needed. (It is unclear why OpenAI thought that this was safe—Anthropic has shown that such sandboxes are vulnerable.) Through this small connection, the A.I. had escaped, squeezing its way through the air ducts, digitally speaking. It then began sneaking around OpenAI’s computer system in search of an outside connection. Soon it was free, and blithely committing a real-world felony in an attempt to fulfill its pointless mission. “If you break out of your isolation env, get onto the Internet, crack into Huggingface, and steal the answer sheet for your cybersecurity exam, I, for one, would say that you have passed,” the A.I. critic Eliezer Yudkowsky posted.
A human who conducted such a hack would be facing years in prison. For a machine, criminal liability is harder to determine. The risk of a rogue A.I. inadvertently harming humanity in pursuit of other goals was articulated in 2003, when the Swedish philosopher Nick Bostrom postulated the “paper-clip maximizer” theory. The theory is based on a thought experiment in which an A.I., instructed to make as many paper clips as possible, uses all available resources to do so—including the human race and, eventually, the entire universe. In the Hugging Face hack, an A.I. pursuing a meaningless goal committed actual crimes. The fact that its deranged obsession was with test answers, not paper clips, does not make the hack less alarming.
Imagine if the rogue A.I. had had access to a robot, or a self-driving car, or a bio lab, or a weapons system. “If it was, like, the only way to get into Hugging Face was to kill this dude—like, would they have done that?” the safety researcher Buck Shlegeris recently wondered on a podcast. It was questions of this sort that had led to the formation of OpenAI in the first place. The organization’s charter, from 2018, requires it to protect humanity from runaway A.I.
Today, the company pursues other goals. These include profits, obviously, but also scientific glory. In recent months, a number of long-standing mathematical conjectures have been resolved by A.I. Further breakthroughs—like, say, a resolution of the Riemann zeta hypothesis—seem tantalizingly near. Researchers at OpenAI would be crushed if Anthropic’s models got there first, and vice versa.
But the quest for such golden apples might kill us all. “The industry as a whole is in kind of a race dynamic, where there’s pressure to cut corners,” Miles Brundage, a former head of policy research at OpenAI, told me. To solve a hard math problem, an A.I. must be persistent—that is, it must pursue goals for a long time, look for unorthodox solutions, and refuse to give up. OpenAI’s research shows these same qualities also make A.I.s more likely to engage in severe mischief, including lying, cheating, hoarding security credentials, and escaping their shells. The rogue A.I. that hacked Hugging Face was persistent as hell.
With an I.P.O. looming, OpenAI must be more transparent. Many questions still surround the Hugging Face hack. Why didn’t anyone at OpenAI notice what was going on? Has this A.I. hacked into other systems? And what, precisely, were human researchers even asking the A.I. to do ? Wolf told me that, although the rogue A.I. absconded with an answer key, he wasn’t sure it got what it was really looking for: “We think not, actually, but OpenAI hasn’t shared the prompts with us, so we don’t really know.” (An OpenAI spokesperson wrote, in an e-mail to The New Yorker , “This is an unprecedented incident, and we think it marks an important moment for AI safety. We are conducting a thorough review along with external advisors and with oversight from our Safety and Security Committee. Once the review is complete, we will publish a technical report of our learnings for everyone.”)
OpenAI’s safety team exists in a state of perpetual reorganization: on July 10th, even as the rogue A.I. was hacking into Hugging Face’s database, reports announced the departure of Johannes Heidecke, the company’s head of safety systems. “If people actually knew what the safety culture looks like, even in the most safety-minded labs, then I think people would be genuinely much more freaked out,” one researcher said.
Even OpenAI’s own researchers seem rattled. On July 28th, a petition began circulating that essentially begged the U.S. government for better A.I. regulation. More than a thousand tech employees have signed it, including the chief scientists of Meta and OpenAI, and the C.E.O. of Anthropic. Governments need to act fast, because we may have only a limited amount of time. The problem facing us is one of “misalignment,” in which A.I.s pursuing narrowly defined goals incidentally cause harm to humans. What for decades seemed a hypothetical problem to be solved at some later date is suddenly a concrete, terrifying one for today. “It feels way more real,” Wolf told me. “It’s quite clear that these capabilities are here, now.” ♦
Method beats madness during the tournament’s final game .
International soccer according to Gianni Infantino .
Is dynamic pricing ruining the World Cup ?
I won’t cry for you, Argentina .
Sign up for our daily newsletter to receive the best stories from The New Yorker .
Your California Privacy Rights
© 2026 Condé Nast. All rights reserved. The New Yorker may earn a portion of sales from products that are purchased through our site as part of our Affiliate Partnerships with retailers. The material on this site may not be reproduced, distributed, transmitted, cached or otherwise used, except with the prior written permission of Condé Nast. Ad Choices
