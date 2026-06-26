---
source: "https://ninjasandrobots.com/paul-graham-flagged-for-ai-use"
hn_url: "https://news.ycombinator.com/item?id=48686737"
title: "Local AI detectors rated a 2013 Paul Graham essay as more AI than actual slop"
article_title: "Paul Graham Flagged For AI Use"
author: "nate"
captured_at: "2026-06-26T14:52:53Z"
capture_tool: "hn-digest"
hn_id: 48686737
score: 4
comments: 0
posted_at: "2026-06-26T13:59:57Z"
tags:
  - hacker-news
  - translated
---

# Local AI detectors rated a 2013 Paul Graham essay as more AI than actual slop

- HN: [48686737](https://news.ycombinator.com/item?id=48686737)
- Source: [ninjasandrobots.com](https://ninjasandrobots.com/paul-graham-flagged-for-ai-use)
- Score: 4
- Comments: 0
- Posted: 2026-06-26T13:59:57Z

## Translation

タイトル: 地元の AI 検出器は、2013 年のポール・グラハムのエッセイを実際の結果よりも AI であると評価しました
記事のタイトル: ポール・グレアム、AI 使用の疑いで警告
説明: 炎をショートさせてみましょう。彼は AI を使用していませんでしたが、フィード リーダーから AI のスロップを排除しようとした私が彼に最悪の犯罪者としてフラグを立てました。彼ですか？いいえ、それはこれがいかに難しいかを示しているだけです。それで昨日、私は素晴らしい記事を投稿しました... |忍者とロボット |エンジニア、デザイナー、創設者
[切り捨てられた]

記事本文:
Ninjas and Robots は Svbtle ネットワークに書き込みを行っています。
ポール・グラハムがAI使用の疑いで警告
炎をショートさせましょう。彼は AI を使用していませんでしたが、フィード リーダーから AI のスロップを排除しようとした私が彼に最悪の犯罪者としてフラグを立てました。
彼ですか？いいえ、それはこれがいかに難しいかを示しているだけです。
それでつい昨日、私は Hacker News に素晴らしい記事を投稿しました。これは Kagi Small Web からのもので、私はここのような小さなブログをサポートしたいので、フィード リーダーでよく使っています。いつものゴミにはうんざりだ。そして、ここにある記事には、私が 2005 年末から YC サークルに参加していることを考えると、これまで聞いたことのない興味深い内容が含まれています。Airbnb が、家計を立てるために社宅に行きそうになったのと同じでしょうか?
すぐに投稿は Hacker News のトップに躍り出ました。多くの賛成票が集まっています。しかしその後、「slop」や「AI;DR」が表示されるようになり、その後、誰かが作者の x ハンドルが存在すらしないと指摘しました。
クソ。騙されてしまいました。この記事はとんでもなくAIが生成したものです。サイト全体は、siliconopera.com です。私がクリックした著者のうち、すべての記事はテンプレートに従っており、毎日同じテーマを再生成しています。これらは、存在しない、またはまったく異なる x.com アカウントを指しています。
最悪だった。私の提出物には当然のことながら[フラグが付けられました]。
もちろん恥ずかしいです。しかし、少なくともそれを問題を解決するための燃料として使用してみることはできます。私は尋ねました：
満足している適切な「AI 検出」アルゴリズム/サービス/オンデバイス モデルを使用している人はいますか?
引き受け手はいない。クロード・オーパスに聞いてみればいいのに。シリコンオペラの記事を通過させると、自信を持って AI スパムだと判断しますが、リンクをたどったり、偽の著者のマストヘッドを掘り下げたりする作業も行いました。記事を読む前にすべての記事でそれを実行することはできません。
それで私は

クロードにローカルで何かを作ってもらい、テストしてもらいました。私のデバイスで安価に実行でき、フィード全体をすばやく処理できるものはありますか?
Appleのオンデバイスモデル。これは高速で、すでに現代の Mac でも簡単に動作するので、私はこれを使って騙すのをとても楽しんでいます。しかし、明らかにそれほど強力ではありません。また、非常にガードレールの多いクロード老人を思い出しました。ある時点で、Apple のモデルは、Pimantle がアダルト Web サイトであると確信していたため、私に Pimantle の同義語を与えることを拒否しました。そうすると、既に 5 つ与えられているので、それ以上の同義語は得られず、それ以上は「リソースの無駄」になります :) 私はそれにいくつかのチャンスを与えました。 1 つは、AI がどの程度であるかを 0 から 100 で採点するプロンプトです。もう 1 つは、AI がどれほど適切な情報源であるかを理解しているかどうかを確認するプロンプトです。
GPT-2 の困惑。基本的に、LLM はあなたの次の言葉にどれほど驚いているかです。 LLM に尋ねたら、私が「JumbleJuice」と言うだろうと予想していましたか? 「いや、この狂った人間は」つまり、要点は、驚きが低いということは、AI の可能性が高いということです。
RoBERTa (Robustly Optimized BERT Pre-Training Approach) は、Facebook が数年前にリリースしたモデルです。誰かが「これは人間、これは ChatGPT」の例を使って訓練し、2023 年にはまともな成績を収めました。しかし、2026 年の ChatGPT は、今予測するのとはまったく異なるものになります。
ファック。ポールのエッセイ (2013 年!!!) によると、最もスパム的な AI は 4 つのうち 3 つにまで落ち込んでいます。RoBERTa のものは明らかに役に立ちません。
しかし、これらの検出器は基本的に、過度に優れた書き込みを測定しようとしただけです。ポールは偉大な作家です。ロボットに違いない。
とにかく、これを見逃したことを少しお詫びしたいと思います。私が毎日頼りにしているコミュニティにこれ以上の騒音を引き起こすのは嫌いです。私はこれをより良いものにしようとしています。そして明らかにいくつかの側面で失敗しています。
うまくいきそうなのは、人間か非常に高価なエージェントを使って記事を掘り下げることです。そして私は期待していました

Kagi Small Web もこれを人力で検証しました。しかし、これは規模が大きいと難しいです。少なくとも、Small Web リストをクリーンアップするために PR を提供できます。
https://github.com/kagisearch/smallweb/pull/817
そして、言うまでもなく私は愚か者です。上で話したフィード リーダーは、私が自分で開発しているものです: PageForth 。これは、ニュースを要約するための素晴らしいツールです。本題に入る前に、まずその主題が私にとって興味深いかどうかを精査したいと思います。皮肉なことに、AI 検出器が組み込まれています。これは、上でテストしたアプローチの 1 つを改良したものです。しかし、明らかにその部分はまだ機能していません:)
Web デザイナーが何度も犯す 3 つの間違い
先日、オンラインの靴屋をレビューしていました。ランディング ページには、これらすべての靴の美しいグラフィックが掲載されていました。見た目も豪華なサイト。そして別のリンクをクリックすると、この食べ物の写真が表示されました。それは奇妙だ。この靴屋には... 続きを読む →

## Original Extract

Let me short-circuit the flames. He wasn’t using AI, but my attempts at trying to rid myself of AI slop in my feed reader flagged him as the worst offender. Is he? No. It just points out how hard this is. So just yesterday, I posted a neat article... | Ninjas and Robots | Engineer, Designer, Founder
[truncated]

Ninjas and Robots is writing on the Svbtle network.
Paul Graham Flagged For AI Use
Let me short-circuit the flames. He wasn’t using AI, but my attempts at trying to rid myself of AI slop in my feed reader flagged him as the worst offender.
Is he? No. It just points out how hard this is.
So just yesterday, I posted a neat article to Hacker News. It was from Kagi Small Web which I’ve been using in my feed reader a ton because hell yeah I want to support the small blogs out there like the one here. I’m sick of all the usual garbage. And here’s an article that has some interesting bits I’ve never heard before, considering I’ve been in the YC circle since late 2005. Like Airbnb almost went corporate housing to make ends meet!?
Immediately the submission jumped to the top of Hacker News. Lots of upvotes coming in. But then I started seeing “slop”, “AI;DR”, and then someone pointed out the author’s x handle doesn’t even exist.
Crap. I got fooled. This article is ridiculously AI generated. The whole site is: siliconopera.com. Of the authors I’ve clicked on, all their articles follow templates and regenerate the same themes day after day. They point to non-existent or completely not-them x.com accounts.
That sucked. My submission was [flagged] as it should be.
Of course I’m embarrassed. But I can at least try to use that as some fuel to fix the problem. I asked:
Does anyone use a decent “ai detection” algo/service/on device model that they are happy with?
No takers. I could just ask Claude Opus. When I put the silicon opera article through it, it confidently thinks it’s AI spam, but it also did the work of following links, digging into the fake author’s masthead, etc. I can’t do that every single article before I read it.
So I tried to get Claude to make something locally for me to test. Is there something that can run cheaply on my device that quickly goes through a whole feed?
Apple’s on-device model. I’ve been enjoying fooling with this one a lot since it’s fast and already easily lives on our modern macs. But it obviously isn’t that powerful. Also reminded of very guardrailsy old Claude. At one point Apple’s model refused to give me synonyms for Pimantle because it was convinced Pimantle was an adult website. Then it wouldn’t give me more synonyms because it already gave me 5, and more would be a “waste for resources” :) I gave it a couple chances. One prompt to just score how AI it is 0 to 100. And another to see if it understood how well-sourced it was.
GPT-2 perplexity. Basically, how surprised is the LLM by your next word. If you ask an LLM, hey did you predict I’d say: JumbleJuice? “Nah. You crazy human.” So the gist: low surprise is high AI likelihood.
RoBERTa (Robustly Optimized BERT Pre-Training Approach) is a model Facebook released a few years ago. Someone trained it with “this is human, this is ChatGPT” examples and it did decent in 2023. But 2026’s ChatGPT is an entirely different beast to predict now.
Fack. Paul’s essay (from 2013!!!) reads as the spammiest AI slop to 3 of the 4. The RoBERTa thing is clearly useless.
But these detectors basically just tried to measure for overly good writing. Paul’s a great writer. Must be a robot.
Anyways, just wanted to apologize a bit for missing this. I hate causing even more noise in the community I lean on every day. I’m trying to make this stuff better. And clearly failing in some dimensions.
What does seem to work is digging into the article with humans or a very expensive agent. And I was hoping the Kagi Small Web had human verified this too. But this is tough at scale. At least I can contribute a PR to cleanup the Small Web list:
https://github.com/kagisearch/smallweb/pull/817
And I’d be an idiot not to mention, the feed reader I talked about above is something I’ve been working on myself: PageForth . It’s been an awesome tool to summarize news I’d like to first vet if the subject matter is even interesting to me before diving in. It has an AI detector built into it ironically. It’s an improved version of one of the approaches I tested above. But clearly that part isn’t working yet :)
Three mistakes I see web designers make over and over again
I was reviewing an online shoe store the other day. The landing page had this beautiful graphic of all these shoes. Gorgeous looking site. And then I clicked on another link, and saw this photo of food. That’s weird. This shoe store has... Continue →
