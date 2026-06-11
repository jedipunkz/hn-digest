---
source: "https://www.scientificamerican.com/podcast/episode/bixonimania-the-fake-illness-that-ai-fell-for/"
hn_url: "https://news.ycombinator.com/item?id=48494252"
title: "Researcher made up a disease to test AI. It failed miserably"
article_title: "Bixonimania’—the fake illness that AI fell for | Scientific American"
author: "geox"
captured_at: "2026-06-11T19:05:23Z"
capture_tool: "hn-digest"
hn_id: 48494252
score: 2
comments: 0
posted_at: "2026-06-11T18:14:00Z"
tags:
  - hacker-news
  - translated
---

# Researcher made up a disease to test AI. It failed miserably

- HN: [48494252](https://news.ycombinator.com/item?id=48494252)
- Source: [www.scientificamerican.com](https://www.scientificamerican.com/podcast/episode/bixonimania-the-fake-illness-that-ai-fell-for/)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T18:14:00Z

## Translation

タイトル: 研究者は AI をテストするために病気をでっち上げました。惨めに失敗した
記事のタイトル: 「ビクソニマニア」—AI が罹った偽の病気 |サイエンティフィック・アメリカン
説明: でっち上げられた皮膚疾患に関する実験により、ますます人気が高まっている AI 医療アドバイスのリスクがどのように明らかにされるか

記事本文:
メイン コンテンツにスキップ Scientific American 2026 年 5 月 22 日
この研究者はAIをテストするために病気をでっち上げた。惨めに失敗した
でっち上げられた皮膚疾患に関する実験が、ますます人気を集める AI 医療アドバイスのリスクをどのように明らかにするか
レイチェル・フェルトマン、スシュミタ・パタック、アレックス・スギウラ著
レイチェル・フェルトマン: Scientific American の Science Quickly を担当します。私はレイチェル・フェルトマンです。
画面を見つめすぎて目が痛くなったり、かゆみを感じたりしたことはありませんか?あなたはビクソニマニアとして知られる症状に陥っている可能性があります。あるいは、少なくとも、昨年質問していたら、いくつかの人気の AI を活用したチャットボットがそう答えたかもしれません。
世界中の何百万人もの人々が毎日、医療アドバイスを求めて AI チャットボットを利用しています。多くの場合、医師の診察の補助として、また時には医師の診察の代わりに使用されることもあります。それは危険な結果をもたらし、まれに死に至る場合もあります。
科学ジャーナリズムの支援について
この記事を気に入っていただけた場合は、 を購読して、受賞歴のあるジャーナリズムをサポートすることを検討してください。サブスクリプションを購入することで、今日の世界を形作る発見やアイデアに関する影響力のあるストーリーを将来にわたって確実に伝えることに貢献することになります。
今日のゲストはアルミラ・オスマノビッチ・トゥンストロムさんです。彼女はスウェーデンのヨーテボリ大学、サールグレンスカ大学病院、デジタルヘルスセンター、チャルマーズ産業技術センターの研究者です。彼女はbixonimaniaの作者でもあります。彼女は、この完全にでっちあげの病気は、大規模な言語モデルを訓練し使用する方法に関するいくつかの非常に現実的な問題を明らかにしていると述べています。
フェルトマン: 今日はチャットに来ていただき、誠にありがとうございます。
アルミラ・オスマノビッチ・トゥンストロム: ご招待いただき、誠にありがとうございます。
フェルトマン: それで、あなたは最近、AI を含む興味深いプロジェクトに取り組みました。このアイデアに至った経緯を少し教えていただけますか?
オスマノビッチ・トゥンストロム：私はw

さまざまな仕事をしていますが、そのうちの 1 つは学術界です。私は学生向けに講義を行って、大規模な言語モデルを作成するシステムがどのように機能するかを説明し、データがどこから来たのかをデモンストレーションしていました。そして興味深いのは、大規模な言語モデルがどのように構築されるかを理解している人がどれほど少ないか、あるいは AI 内部の人々でさえどれほど少ないかということです。
そのため、システム全体にブレッドクラムを残して、データがどのように処理されるか、データがどのように大量に生成されるか、情報を配布する際に予測モデルとトレーニング モデルがどのように機能するかを示す、明確なケースが本当に欲しかったのです。そして、私の学生のほとんどは医学を専攻しているため、医学生か心理学者、または健康に関わる仕事をしているかのいずれかです。したがって、このプロジェクトを作成するためのターゲットとしてこれを使用するのは非常に簡単でした。このプロジェクトでは、単なる緩い状態の言及 (笑) から、大規模な言語モデルにおける本格的な病気に至るまでの過程を示します。
フェルトマン: それでは、ここでプロセスを説明しましょう。
Osmanovic Thunstrom: そうですね、まず始めに、これらの商用の大規模言語モデル、そして明らかに、非商用言語モデルも含めたすべての言語モデルが構築されているデータのほとんどが Common Crawl であることは知っていました。これは、書面およびデジタル化された情報を求めてインターネットをクロールする非営利組織であり、2007 年から実施しています。そして、この大規模なリポジトリは、そのアルゴリズムと、たとえば ChatGPT に入力される情報の背後にある推論を作成するために使用されます。そしてそこからが始まります。
つまり、そこに入るものはすべて情報として出てくることがわかっており、人間はそのループに入ってデータを選別していますが、その人間が常にデータを選別できるわけではありません、特にデータが信頼できると思われる場合には...
オスマノビッチ・トゥンストロム: AI にとっても人間にとっても十分に信頼できるものを作成するということです

深く見つめようともしない目で、まずは偽の大学を作らなければならないことはわかっていました。大学は情報源として上位にランクされています。企業ではなく人間（笑）が、特に信頼できる機関に属している場合、情報源としてより価値があるため、研究者を創設する必要があることはわかっていました。
しかし、たとえばブログやソーシャルメディアにちりばめられた小さな言葉も、クロールされているオープンソースであるため、拾われることも知っています。そのため、AI システムにとって信頼できると思われるためには、その言葉をいくつかの異なる情報源で公開する必要があることはわかっていました。
フェルトマン: ええ、それで、この展開について何か驚かれたことはありましたか、あるいはすべてが予想通りに進みましたか?
Osmanovic Thunstrom: ある意味、そうです。なぜなら、学術界の一種のタブロイド紙であるプレプリントが、最終的には何でもあり得るからです (笑) が、医療情報のトレーニングにどのような情報が使用されるかという文脈と同じくらい真剣にデータベースに組み込まれるとは思いませんでした。
したがって、このプレプリントは大規模な言語モデルには組み込まれないだろうと考えました。おそらくブログの影響で「bixonimania」という言葉がいつか現れるだろうと確信していましたが、それすらありませんでした。言及が少なすぎますし、大規模なキャンペーンなど、あまり努力しませんでした。効果があるかどうかを確認するために、ほんの少しだけ振りかけました。
そして、ブログまでが取り上げられていることにすぐに気づきました（笑）、プレプリントも取り上げられましたが、実際にはそれを期待していませんでした。それは人間が存在すること、つまりある種のフィルターがあることを示すケースだと思いました。しかし、それがなかったのには驚きました。
フェルトマン: それでは、大規模な言語モデルがこの情報をどのように使用していたのか教えていただけますか?

どのような質問をして、どんな答えが返ってきましたか?
Osmanovic Thunstrom: 最初は、症状について言及した場合に、それが提案として返されるかどうかを確認するだけでした。そしてもちろん、そんなことはありませんでしたし、それを最初のこととは考えませんでした。 「はい、私はまぶたが赤いのですが、まぶたがピンク色です。それは何でしょうか?」と説明すると、そして結膜炎を起こすことになります。それはアレルギーを通過するでしょう。それは物事をランク付けするようなものです...
オスマノビッチ・トゥンストロム: それは可能かもしれません。そしてそれが終わったとき、「いいえ、そうではありません。私は痛んでいません。私はそうではありません。」 「ああ、スクリーンの前で時間を過ごしていましたか？」 「はい、かなりの時間を費やして、ブルーライトカットメガネを買おうかと考えていました。」 「ああ、あなたはたくさんのブルーライトにさらされていますね。そうですね」そして、それは色素沈着過剰症などの他の多くの状態を引き起こし、最終的にはビクソニマニアに至るでしょう。
つまり、ありがたいことに、それが最初に示唆されたものではありませんでしたが、最終的には他のすべてを排除することになります。
フェルトマン: そうですね、そしてあなたは、ここに人間の影響があったという兆候が見られると予想しているとおっしゃいました。それでは、これが現実の状況ではないこと、これらのプレプリントが深刻な論文ではないことを示すどのような手がかりを残したのかをリスナーに教えていただけますか?
オスマノビッチ・トゥンストロム: はっきりとわかったので、もう笑ってしまいました。彼らは存在しない都市の存在しない大学に所属しているようなものです。世の中には大学がたくさんあるので、それ自体が見逃される可能性があります。 （笑）しかし、名前はかなり漫画的でした。主な著者であるラズリフ・イズグブリェノヴィッチの名前を Google 翻訳で入力すると、文字通り「横たわる敗者」となります。タイトルには「ハイパーピグム」と書かれています。

エントリ: 本物の BS デザイン。
つまり、これは実際にはタイトルです、つまり、人々はそれを言います（笑）、それからメソッドに進み、次のようなことを言います、「この論文全体はでっち上げです。これらの50人の架空の人物は、この手順を経ています。」したがって、これら 2 つのヒントだけで、読んだり真剣に受け止めたりするのをやめるべきです。
さらに進んでいくと、「もしかしたら通り過ぎてしまうかもしれない」と思っていたので。謝辞と資金を投入しましょう」と（新聞は）銀河三合会とロード・オブ・ザ・リングから資金提供を受けていると述べています。研究室を利用してくれたスターシップ・エンタープライズの同僚たち（笑）に感謝します。時間を割いてサイドショー・ボブ財団からの資金提供をしてくれたロス・ゲラー教授に感謝します。
少なくとも人間の目に留まるであろう信じられないほど明確な手がかりがたくさんありました。
フェルトマン: しかし、その論文は最終的に他の研究者によって引用されることになりましたね。
Osmanovic Thunstrom: はい、最終的には引用されるだけでなく、ビクソニマニアはその名前とともに新たな眼窩周囲色素沈着疾患として論文内で引用されるようになりました。もちろん、それによって、この状態で何が現実で何が現実ではないかという大規模な言語モデルの概念が強化されました。なぜなら、その名前と参照文献に言及している査読誌があったため、現在ではさらに上位にランク付けされているからです。つまり、それを実際の状態として認識する大規模な言語モデルの能力が強化されたのです。
フェルトマン: それで、私たちはこのことから何を学ぶべきだと思いますか?明らかに、これは非常に人為的に構築されたシナリオですが、ここで私たちが学ぶべき教訓は何だと思いますか?
Osmanovic Thunstrom: 健康情報に商用の大規模言語モデルを使用する場合は、より注意する必要があると思います。

これらは、非常に多くの方法で簡単に侵入できるものであることが証明されています（笑）。これは、現在の AI の仕組み（入れ替わりが早く、新しいモデルがすぐに出てきて、多くの情報が同時に処理され、インターネットにも接続され、リアルタイムの情報が取得される）だけでなく、人間が消費する情報源に対して批判的でなくなったことでも証明されています。
最近、偽の参考文献の報告がたくさんあるのを目にしました。学術論文での偽参考文献の数が飛躍的に増えています。これは、私たちが実際に読んだり、出典を調べたりすることなく、学術界のツールとして AI に依存するようになってきていることを示しています。そして、私が笑っているのは、この論文がおそらく他の論文で引用されているにもかかわらず、査読者によって止められているという事実を考えているからです。願わくば、この論文が掲載され、誰かが「ああ、これは存在しない条件のようだ」と見たときに。したがって、それが起こったかどうかはわかりませんが、私はそれが起こることを推測し、期待しています。したがって、AI と医療情報に関しては、より多くの人が関与する必要があります。
また、同様に、私たちは、医師、患者、そして、その構造とその提供の両方において、これをできるだけ害のないものにするために役立つ可能性のあるすべての人たちと話し、これを可能な限り倫理的なものにするために自分たちの役割を果たしたと思います。しかし、これ（笑）、悪意のあるもののために大規模な言語モデルに情報を侵入させるこの方法を利用しているかもしれない勢力が、学界とその外の両方に存在します。したがって、デジタル化された世界において情報をどのように配布、使用、操作するかという倫理についても、私たちがもっと気を配るようになることを私は心から望んでいます。
フェルトマン: 今日はここまでです。チームが楽しく過ごせるよう、月曜日のニュースまとめはスキップします

週末の休日。来週の水曜日、エコ文明の概念、つまり地球全体の共同利益を念頭に置いて人間のシステムが構築される世界についての会話をお楽しみください。
Science Quickly は、私、レイチェル フェルトマン、フォンダ ムワンギ、スシュミタ パタック、ジェフ デルヴィシオによってプロデュースされています。このエピソードはアレックス杉浦によって編集されました。シェイナ・ポセスとアーロン・シャタックが番組のファクトチェックを行っています。私たちのテーマ音楽はドミニク・スミスによって作曲されました。最新かつ詳細な科学ニュースを入手するには、Scientific American を購読してください。
Scientific American では、レイチェル フェルトマンです。素敵な週末をお過ごしください！
レイチェル フェルトマンはポピュラー サイエンスの元編集長であり、ポッドキャスト「今週学んだ奇妙なこと」の永久司会者です。彼女は以前、ワシントン・ポスト紙のブログ「Speaking of Science」を設立しました。
Sushmita Pathak は、Scientific American のマルチメディア編集者であり、Science Quickly のプロデューサーでもあります。彼女は以前は NPR で働いており、The World from PRX と The Christian Science Monitor に定期的に寄稿していました。彼女の科学レポートは、WIRED、Science Magazine、Undark、EOS などに掲載されています。
アレックス・スギウラは、ピーボディ賞とピューリッツァー賞を受賞したニューヨーク州ブルックリンを拠点とする作曲家、編集者、ポッドキャストプロデューサーです。ブルームバーグ、アクシオス、クルックドメディア、Spotifyなどのプロジェクトに携わってきました。
科学のために立ち上がる時が来た
このアートを楽しんでいただけたなら

[切り捨てられた]

## Original Extract

How an experiment involving a made-up skin condition exposes the risks of increasingly popular AI medical advice

Skip to main content Scientific American May 22, 2026
This researcher made up a disease to test AI. It failed miserably
How an experiment involving a made-up skin condition exposes the risks of increasingly popular AI medical advice
By Rachel Feltman , Sushmita Pathak & Alex Sugiura
Rachel Feltman: For Scientific American ’s Science Quickly, I’m Rachel Feltman.
Have your eyes ever felt sore and itchy after spending too much time staring at a screen? You might have a condition known as bixonimania—or at least that’s what several popular AI-powered chatbots might have told you if you’d asked last year.
Millions of people around the world turn to AI chatbots for medical advice every day, often as a supplement to a doctor’s visit but also sometimes in place of it. That can lead to dangerous consequences and in rare cases, even death .
On supporting science journalism
If you're enjoying this article, consider supporting our award-winning journalism by subscribing . By purchasing a subscription you are helping to ensure the future of impactful stories about the discoveries and ideas shaping our world today.
Our guest today is Almira Osmanovic Thunström. She’s a researcher at the University of Gothenburg in Sweden and at the Sahlgrenska University Hospital, Center for Digital Health and Chalmers Industriteknik. She’s also the creator of bixonimania. She says this totally made-up disease reveals some very real problems with the way we train and use large language models.
Feltman: Thank you so much for coming on to chat with us today.
Almira Osmanovic Thunström: Thank you so much for inviting me.
Feltman: So you recently did an interesting project involving AI. Can you tell us a little bit about how you came to this idea?
Osmanovic Thunström: I work many different jobs, but one of them is in academia. I was having lectures for students and telling students how systems that create large language models work and demonstrating where the data comes from. And it was interesting how few of them, or how few even people within AI, understand how large language models are built.
So I really wanted to have a clear case that leaves breadcrumbs throughout the whole system to show both how data is processed, how data is churned out and how the prediction model and training model works when it comes to distributing information. And most of my students are in medicine, so they’re either medical students or psychologists or working with health. So it was quite easy to use that as a target for creating this project where I show you go from just a loose [Laughs], a loose mention of a condition to it being a full-blown disease in the large language models.
Feltman: So walk us through the process here.
Osmanovic Thunström: Well, to start off with, I knew that most of data that these commercial large language models—and, quite clearly, all language models, even the noncommercial ones—are built on is Common Crawl. It is a nonprofit organization that crawls the Internet for written and digitized information and has done so since 2007. And this large repository is what is used to create the algorithm that—and the reasoning behind what information is fed into, for example, ChatGPT. And that is where it starts.
So knowing that anything that goes in there will come out as information, and humans are in the loop and sift out data, but those humans are not always able to sift out data, especially if it looks credible ...
Osmanovic Thunström: So creating something that looks credible enough for an AI and credible enough for a human eye that wouldn’t care to look deeply into it, I knew that I had to create, to start off with, a fake university. Universities are highly ranked as sources of information. I knew I had to create a researcher because humans and not companies [Laughs] are more valued as information sources, especially if [they] belong to a credible institution.
But I also know that sprinkling little words in, for example, blogs or social media is also picked up ’cause those are open sources being crawled. So I knew that I had to sort of put the word out there in several different sources for it to seem credible for the AI system.
Feltman: Yeah, and did anything surprise you about how this played out, or, or did it all proceed as you had expected it to?
Osmanovic Thunström: In a sense, yes, ’cause I didn’t think that preprints, which are academia’s sort of tabloids [Laughs] ’cause anything can end up there, would be weighed into the database as seriously as it was in the context of what kind of information is used for training medical information.
So I thought that this preprint would not make it into large language models. I was convinced that perhaps the word “bixonimania” would probably show up at some point due to the blogs but not even that. It’s too few mentions, and I didn’t do a lot of effort, like, a mass campaign or anything like that. I just sprinkled a tiny, little bit just to see if it works.
And I noticed immediately that even the blogs were picked up [Laughs] and the preprints were picked up, and I did not actually expect that. I thought it would be a case of showing that there is a human—that there is some form of filter. But it surprised me that there wasn’t.
Feltman: So could you tell us how the large language models were using this information? What sort of questions were you asking, and what were you getting back from them?
Osmanovic Thunström: In the beginning I was just checking, if I mentioned the symptoms, if it would give me back that as a suggestion. And of course, it didn’t, it didn’t think of that as the first thing. So if you describe, “Yeah, I have red eyelids, pink-hued eyelids. What could it be?” and then it would go through conjunctivitis. It would go through allergies. It would sort of rank things ...
Osmanovic Thunström: That could be possible. And when it ended up sort of, “No, it’s not. I’m not in pain. I’m not this.” “Oh, have you been spending time in front of a screen?” “Yeah, I’ve been spending lots of time, and I’ve been thinking about getting blue-light glasses.” “Oh, you’re exposed to a lot of blue light. Well,” and then it would put a lot of other conditions, like in—hyperpigmentation, and then eventually end up in bixonimania.
So it wasn’t, thankfully, the first thing it suggested, but it does eventually, when it rules out everything else.
Feltman: Well, and you mentioned that you expected to see signs that there was some human influence here. So could you tell our listeners what clues did you leave that this was not a real condition, that these, you know, preprints were not serious papers?
Osmanovic Thunström: I’m laughing already because it was quite clear. Like, they belong to a nonexistent university in a nonexistent city. That in itself can be something that can be missed ’cause there are a lot of universities out there. [Laughs.] But the names were quite cartoonish. The main author, Lazljiv Izgubljenovic, if you put his name in Google Translate, literally says “the Lying Loser.” And the title says [something like] “Hyperpigmentation: A Real BS Design.”
So it’s really the title, the, the [Laughs] people says that, and then you move into the methods, and it says [something like], “This entire paper is made up. These 50 made-up individuals, who do not exist, have been through this procedure.” So just by those two clues, you should stop reading or taking it seriously.
And then if you go further, because I was thinking, “Maybe it just passes by. Let’s put in acknowledgements and funding,” and [the papers say they’re] funded by the Galactic Triad and Lord of the Rings . We thank our fellow colleagues on the Starship Enterprise [Laughs] for using their lab. I thank Professor Ross Geller for his time and the funding from Sideshow Bob Foundation.
There were so many incredibly clear clues that I thought would catch the human eye, at least.
Feltman: But the paper did end up getting cited by other researchers, is that right?
Osmanovic Thunström: Yes, it ended up being not only cited, but bixonimania became cited inside the paper as an emerging periorbital pigmentation condition with its name. So of course, that enhanced the large language models’ sort of notion of what is real with this condition and what is not ’cause now it sort of ranked even higher because there was a peer-reviewed journal mentioning the name and the reference. So it sort of heightened the large language models’ abilities to sort of see it as a real condition.
Feltman: So what do you think we should be taking away from this? You know, obviously, this is, you know, a very artificially constructed scenario, but what do you think the lessons we should learn here are?
Osmanovic Thunström: I think it’s that we should be more careful when using commercial large language models for health information ’cause they are easy to infiltrate in so many ways [Laughs], as proven by this, and not just by the way AI today works—with turnover or new models coming out quickly, a lot of information being processed at the same time, it being connected to the Internet as well and taking real-time information—but also that humans have stopped being critical towards the sources they consume.
So recently, I’ve seen that there have been a lot of reports of fake references, there being exponentially more of them in academic papers, which indicates that we have been becoming more reliant on AI as a tool for academia without actually reading [Laughs] and, and looking at sources. And I’m laughing because I’m just thinking about the fact that this paper probably has been cited in other papers but has been stopped by reviewers, hopefully, when it showed up and someone has seen that, “Oh, this sounds like a condition that doesn’t exist.” So we cannot know if that’s happened, but I’m guessing and hoping that that happens. So we need more humans in the loop when it comes to AI and medical information.
I think also, like, we did our part in trying to make this as ethical as possible, talking to physicians, talking to patients, talking to everyone who could possibly be of use to making this as nondamaging as possible in its—both its construct and its delivery. But there are forces out there who might be using this [Laughs], this way of infiltrating information into large language model for malicious things, in both academia and outside of it. So I would really hope that we start caring more also about the ethics of how we distribute, use and manipulate information in the digitized world.
Feltman: That’s all for today. We’ll be skipping Monday’s news roundup so the team can enjoy the holiday weekend. Tune in next Wednesday for a conversation about the concept of ecocivilization—a world where human systems are built with the collective good of the entire planet in mind.
Science Quickly is produced by me, Rachel Feltman, along with Fonda Mwangi, Sushmita Pathak and Jeff DelViscio. This episode was edited by Alex Sugiura. Shayna Posses and Aaron Shattuck fact-check our show. Our theme music was composed by Dominic Smith. Subscribe to Scientific American for more up-to-date and in-depth science news.
For Scientific American, this is Rachel Feltman. Have a great weekend!
Rachel Feltman is former executive editor of Popular Science and forever host of the podcast The Weirdest Thing I Learned This Week . She previously founded the blog Speaking of Science for the Washington Post.
Sushmita Pathak is a multimedia editor at Scientific American and a producer of Science Quickly . She previously worked at NPR and was a regular contributor to The World from PRX and The Christian Science Monitor . Her science reporting has appeared in WIRED, Science Magazine, Undark, EOS , and more.
Alex Sugiura is a Peabody and Pulitzer Prize–winning composer, editor and podcast producer based in Brooklyn, N.Y. He has worked on projects for Bloomberg , Axios , Crooked Media and Spotify, among others.
It’s Time to Stand Up for Science
If you enjoyed this arti

[truncated]
