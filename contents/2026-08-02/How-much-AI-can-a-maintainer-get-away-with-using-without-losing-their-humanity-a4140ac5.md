---
source: "https://www.jvt.me/posts/2026/08/02/ai-maintainer/"
hn_url: "https://news.ycombinator.com/item?id=49148708"
title: "How much AI can a maintainer get away with using without losing their humanity?"
article_title: "How much AI can a maintainer get away with using without losing their humanity? · Jamie Tanna | Software Engineer"
author: "zdw"
captured_at: "2026-08-02T22:46:14Z"
capture_tool: "hn-digest"
hn_id: 49148708
score: 2
comments: 0
posted_at: "2026-08-02T21:47:45Z"
tags:
  - hacker-news
  - translated
---

# How much AI can a maintainer get away with using without losing their humanity?

- HN: [49148708](https://news.ycombinator.com/item?id=49148708)
- Source: [www.jvt.me](https://www.jvt.me/posts/2026/08/02/ai-maintainer/)
- Score: 2
- Comments: 0
- Posted: 2026-08-02T21:47:45Z

## Translation

タイトル: メンテナは人間性を失わずにどれだけ AI を使用できるでしょうか?
記事のタイトル: メンテナは人間性を失わずに、どれだけ AI を使用できるでしょうか? · ジェイミー・タンナ |ソフトウェアエンジニア
説明: AI によって生成されたコントリビュートの猛攻撃について、そしてメンテナーが「火で消火」できる余地があるかどうかを考えます。

記事本文:
メンテナは人間性を失わずにどれだけ AI を使用できるでしょうか? · ジェイミー・タンナ |ソフトウェア エンジニア ジェイミー タンナ |ソフトウェアエンジニア
私が維持しているオープンソース プロジェクト
メンテナは人間性を失わずにどれだけ AI を使用できるでしょうか?
なぜプレッシャーを感じているのでしょうか？
私は最近、オープンソースのメンテナンスが最近少しずつ大変になってきているように感じられていることについて考えています。主に AI エージェントで可能な貢献量が増加しているためです。また、それがメンテナーであることにどのような影響を与えるかについて考えています。
私の文字通りの仕事が (とりわけ) オープンソース プロジェクトのメンテナーであること、そしてそれが他の多くの人にはない特権であることは非常に幸運です。とはいえ、毎月増加しているように見える仕事の量が元に戻るわけではありません。
Mike McQuaid は最近、オープンソースがどのように楽しくなければならないかについて雄弁な投稿を書きました。これは当然のことながら共感を呼びました。私たちの仕事は持続可能で楽しいものである必要があります。
おそらく私の ADHD のせいで、常に燃え尽き症候群の限界に達していると感じている私は、やるべきことが多すぎて残念ながら、自分の精神的健康のためにプロジェクトから一歩退かなければならないという状況に陥ったことが何度かあり、少し壁にぶつかりました。
これは、oapi-codegen とそれをより持続可能なものにしようとしている方法に特に当てはまります。 oapi-codegen は、難しい API サーフェスを備えた広く使用されているプロジェクトです。任意の OpenAPI 仕様を提供していただければ、複雑または単純な仕様を、一般的に使いやすい Go の形式に変換します。 Go のそれほど優れた型システムではないのと同様に、ユーザーの使用方法に従うことは難しい場合があります。
最近では、Renovate でこの苦痛を感じており、コントリビュートと機能リクエストが大幅に増加しています。

at は、AI エージェントの支援による貢献の使用が主な原因で、私たちのチームの規模に実際には対応していません。
(しかし同時に、採用とエンゲージメントの増加が見られ、コードベースのパターンが AI エージェントが貢献を支援するのに非常にうまく機能していることは非常に素晴らしいことです。)
これは少し前から気になっていたことで、プロジェクトリーダーとして、コミュニティからの貴重な貢献も考慮しながら、プロジェクトとして物事をより持続可能なものにするために何ができるかを検討しています。
この投稿には答えがありません。これは、私が自分の気持ちをじっくり考えるための方法であり、他のメンテナーの状況を聞いて、AI への貢献の増加に対処するために AI に依存する必要があるかどうか、あるいは自由に使える他の実践方法があるかどうかを聞くための行動喚起です。
なぜプレッシャーを感じているのでしょうか？
これは、特にここ数年のオープンソース プロジェクトの運営に慣れている人なら驚くような答えではありません。しかし、明確にしておきたいのは、バグ レポート、機能リクエスト、レビューするプル リクエスト、(潜在的な) セキュリティ脆弱性など、プロジェクトへの貢献が大幅に増加しているためです。
9 月に私がプロジェクト リーダーとして Renovate プロジェクトに参加したとき、オープンなプル リクエストは約 100 件あり、私はできる限りその数を 100 未満に抑えるよう努めていました。今日の時点で、約 340 頭となっていますが、これは過去数か月間かなり安定した数字であり、ヒドラの首を切り落とすのと同じ速度で増加しています。
(この記事を書いている時点で、私にはオープンな PR が 72 個あることに気づいた人もいるかもしれません。調整を完了する必要があるため、これらの多くは下書きであり、これが私たちが所有する PR の全体の数にかなりの部分を追加していることは承知しています)
以前のように

先ほども言いましたが、コアメンテナチームはわずか 3 人です。 Rahul や Sergei のような素晴らしい仕事をする協力者がいますが、機能とレビューの最終決定は私たち 3 人のメンテナーです。私たち 3 人のうちプロジェクトに 100% フルタイムで取り組んでいる人は 0 人もいないことを考えると、このチームは小さなチームであり、プロジェクトとして非常に広く使用されています。
この寄付金の増加を視野に入れるために、過去数年間のいくつかの統計を見てみましょう。
私は、私たちの感情が事実と一致するかどうかをある程度理解するために、「メンテナ ダッシュボード」からデータを取得するように Claude Opus 4.8 を設定しました。これにはプル リクエストからの情報も含まれています。このデータはすべて、2024 年、2025 年、2026 年の 1 月から 7 月までに分析されました。
また、GitHub Discussions がユーザーとの重要なインターフェースであることを考慮して、GitHub Discussions のやり取りがどのように変化したかについても調査したいと思いました。全体的に Renovate の採用が増加していますが、ディスカッション フォーラムに寄せられる種類の質問に答えるために AI を使用する人が増えているようです。これは良い面と悪い面があり、AI モデルがユーザーが必要とするサポートを提供できるのは良いことですが、プロジェクトとのやり取りが減ると、ドキュメントを改善したり粗雑な部分を滑らかにしたりするために、一般的な問題に関する情報が少なくなります。
また、2024 年と 2025 年は PR への貢献額がかなり一貫していましたが、2026 年には同じ期間に最大 60% 多くの貢献があったこともわかります。
データを見ると、1 回限りの貢献者と、プロジェクトにより深く関与し続ける人々をどのくらいの頻度で見かけるかについて、より詳細な内訳も確認できます。
最後に、プロジェクトへの貢献がより一般的に大きくなったこともわかります。PR のサイズも拡大していることがわかります。
これらはすべて、需要が増加していることを示しています

これらの PR をレビューするメンテナの作業に貢献します。私たちは、レビューの回数や一般的な問題の修正が必要になる回数を減らすために、リンターとドキュメントを追加するためにできる限りのことを行っていますが、変更にレビューを適用する必要がある貴重な作業がまだたくさんあります。
貢献度の増加とクローン技術の不足により、メンテナー チームは当面 3 人にとどまることを意味します。ここで疑問が生じます。どうすれば仕事を持続可能にできるでしょうか?すべてに対応するには AI ツールに頼る必要があるでしょうか?
私は、ほとんどの場合、自分のコミュニケーション スキルとユーザーへの共感に誇りを持っています。それをすべて AI の支配者にアウトソーシングし始めるのであれば、人間が会話に参加することに何の意味があるのでしょうか?
私は、問題や PR コメントに返信するために完全に AI エージェントを使用しているメンテナーとして見えるいくつかのプロジェクトと対話しましたが、それは非人間的で少し残念に感じられます。彼らがそうする理由はよくわかりますが、それでも気分が良いものではありません。
AI が生成したメッセージを受信する側になるのが楽しくないのに、なぜユーザーにそんなことをするのでしょうか?たとえエージェントを私にとってより自然な話し方になるよう説得しようとしても、それは私の言葉や考えではありません。そして私は一般的に、「あなたがそれを書くのに抵抗がなければ、なぜそれを読む必要があるでしょうか？」という考えを持っています。
私の共同メンテナである Marcin は最近、oapi-codegen で私たちをより良い状態に導くために非常に多くの素晴らしい仕事をし、膨大な量の圧倒的なバックログをなんとかクリアしました。それは、彼が Anthropic の Claude に頼っていたからこそ可能になったもので、彼がそれを実現できたことを本当にうれしく思います。しかし同時に、それは「ボット」が作業を行っていると感じた一部のユーザーにとって、いくつかの驚くべきシナリオにもつながりました。
AI の使用をコーデックに反映させることを目指していますが、

PR/問題などに関するメッセージややり取りについては、AI 経由であることを誰かが知っていればさらに改善されるかどうかはまだわかりません。人間ではなく AI エージェントが私に返信していると知って気分が良くなったわけではありません。
しかし、ユーザーが AI エージェントを使用して返信しているように見える場合、時間をかけて返信する価値があるのでしょうか、それともエージェントにも返信してもらう必要があるのでしょうか?
キーストロークが非常に多く残っていることを考えると、私は自分自身について感謝している側面、つまりユーザーに対する人間的共感、そしてユーザーに対する私への人間的共感を失わずに、自分の時間をどのように最適化したいかを考えています。人間としてイライラしたり、就寝前の数分間に返信したりすることがありますが、AI モデルを通じて無害化するのではなく、それも重要です。
そして最も重要なことは、世界中のオープンソース管理者の皆さん、このすべてについてあなたが何を考えているのか、そしてどのように対処しているのかを教えてください。
あなたは AI を使用してプロジェクトや精神的健康を維持しているメンテナーですか?何が効果的ですか?何をお勧めしませんか?
自動コード レビュー ツールや、「問題を解決できる可能性がある方法」を処理するためのツールなど、便利だと感じているツールはありますか?
自分の言葉を使用するのと比べて、作成/実際にユーザーに応答するという点でどのくらい使用していますか?
人々が AI を使用して対話している場合 (あなたへの返信など)、あなたは AI を使用して彼らに返信することに頼っていますか?つまり、 「AIでAIと戦う」？
この投稿のパーマリンクは https://www.jvt.me/posts/2026/08/02/ai-maintainer/ で、概要は次のとおりです。
AI によって生成されたコントリビュートの猛攻撃について、そしてメンテナーが「火で消火」できる余地があるかどうかを考えます。
この投稿の正規 URL は次のとおりです。
https://www.jvt.me/posts/2026/08/02/ai-maintainer/
。
ジェイミー・タンナ著
日, 02 8月 2026 20:39:42+

01:00 、最終更新日は 2026 年 8 月 2 日日曜日 20:39:42+01:00 です。
この記事のコンテンツは、クリエイティブ コモンズ表示、非営利継承 4.0 インターナショナルの条件に基づいて共有され、コードは Apache License 2.0 に基づいて共有されます。
🤖 このブログ投稿のコンテンツ (散文またはコード スニペット) には、次の LLM から派生したコードが含まれています。
この内容は役に立ちましたか?あなたが何週間も追い求めてきた解決困難な問題は解決しましたか?それとも、毎日再利用できる新しいことを教えてくれましたか?
このようなコンテンツを作成し続けることができるように、私をサポートすることを検討してください。
この投稿は 記事 の下に保管されました。
以下に、このページが WebMention を使用して行ったインタラクションを示します。
この投稿に返信を書きましたか? URLを教えてください:
WebMention 機能を備えた Web サイトを設定していませんか?コメントパレードを使用できます。
興味がありそうなその他のリンク:
このブログをフォローする (つまり、RSS フィード経由)
このサイトは IndieWeb Webring の一部です 🕸💍
連絡を取りたいですか? hi@jamietanna.co.uk までメールを送ってください。すべてに返信するよう努めていますが、数週間以内に返信がない場合は、遠慮なく小言を言ってください。
私を見つけることができるすべての場所のリストをチェックすることもできます。

## Original Extract

Thinking about the onslaught of AI-generated contributions, and if there is space for maintainers to "fight fire with fire".

How much AI can a maintainer get away with using without losing their humanity? · Jamie Tanna | Software Engineer Jamie Tanna | Software Engineer
Open Source Projects I Maintain
How much AI can a maintainer get away with using without losing their humanity?
Why am I feeling the pressure?
I've been recently thinking about how Open Source maintenance seems to be becoming a little more a slog recently - largely with the increased amount of contributions possible with AI agents - and how it impacts being a maintainer.
I'm very fortunate that my literal job is to be an Open Source project maintainer (among other things) and that it's a privilege not many others have. That being said, it still doesn't mean that undoes the amount of work that seems to be increasing every month.
Mike McQuaid recently wrote an eloquent post about how Open Source needs to be fun , which naturally struck a chord - the work that we do needs to be sustainable and enjoyable.
As someone who feels like they're consistently pushing the edges of burnout - possibly due to my ADHD - I've hit a bit of a wall several times, where the overwhelm of so many things to do unfortunately leads to needing to (often silently) take a step back from the project for my own mental wellbeing.
This is moreso true with oapi-codegen and how we're trying to make it more sustainable . oapi-codegen is a widely used project with a difficult API surface: give me your arbitrary OpenAPI specs, and we'll convert the complex or straightforward spec to a form of Go that is generally nice to use. Being guided by our users' usage can be difficult at time, as well as Go's not-that-great type system.
More recently, we've been feeling this pain with Renovate , and a significant uptick in contributions and feature requests that aren't really scaling with the team size we have, largely due to the usage of contributions with help from AI agents.
(But at the same time, it's pretty great that we are seeing an increase in adoption and engagement, and the patterns in our codebase work quite well for AI agents to help with contributions!)
This has been something on my mind for a bit - as the project lead, I'm looking at what we can do to make things more sustainable for us as a project, while also taking into account valuable contributions from the community.
This post doesn't come with an answer - it's more of a way for me to think through my feelings, and a call to action from other maintainers to hear how they're doing, and whether they're having to rely on AI to be able to deal with the increase in AI contributions, or if they have any other practices at their disposal.
Why am I feeling the pressure?
This isn't really an answer that should surprise anyone familiar with running an Open Source project - especially in the last few years. But to make it clear, it's because there's a significant increase in contributions to the project, be they bug reports, feature requests, Pull Requests to review, and (potential) security vulnerabilities.
When I joined the Renovate project as the project lead in September, we had roughly 100 open Pull Requests, and I'd been trying to keep that number below 100 as long as I could. As of today, we have ~340, which has been a fairly consistent number over the last few months, increasing at a rate consistent with cutting off the head of a hydra.
(Some folks may notice that as of writing, I have 72 of those open PRs - a number of these are drafts as I need to finish tweaks on them, and I'm aware that adds a good chunk to the overall number of PRs we have)
As previously mentioned , the core maintainer team is only 3 people . Although we have some collaborators like Rahul and Sergei who do great work, the final decisions on functionality and reviews are the three of us maintainers - which is a tiny team considering that zero of the three of us are working 100% full-time on the project, and we're so widely used as a project.
To put this increase in contributions into perspective, let's look at some stats over the last few years.
I set Claude Opus 4.8 to take data from my "maintainer dashboard" - which also includes information from Pull Requests - in a way to get some understanding of whether our feelings matched the facts. This data has all been analysed between January-July across 2024, 2025, and 2026.
Something I also wanted to investigate is how GitHub Discussions interactions have changed, given it is our key interface with users . Although we're generally seeing increased adoption of Renovate, it appears like more folks are using AI to answer the sorts of questions they'd come to our Discussion forums. This is good and bad - it's good that AI models are able to provide the support users need, but reduced interactions with the project means that we have less information about common issues, to improve documentation or smooth out rough edges.
We can also see that although 2024 and 2025 were fairly consistent in their incoming PR contributions, 2026 has had ~60% more contributions during the same period:
Looking at the data, we can also see more of a breakdown of how often we see one-off contributors vs people who stay more engaged with the project:
Finally, we can also see that contributions to the project now are more generally larger - we see that our PRs are also growing in size:
All of these show that there's an increased demand for maintainer work to review these PRs - we're doing what we can to add linters and documentation to reduce how many rounds of reviews or common issues need fixing, but there's still a lot of valuable work that reviews need to apply to the changes.
With the increase in contribution, and the lack of cloning technology, this means that the maintainer team is remaining at 3 for the time being. Which begs the question - how do we keep our work sustainable? Do we need to lean on AI tooling to be able to keep up with it all?
I'd say that I - most of the time - pride myself on my communication skills and empathy for my users . If I'm starting to outsource all that to our AI overlords, what's the point of a human being part of the conversation?
I've interacted with some projects that do appear as maintainers completely using an AI agent to reply to your issue/PR comments and it feels impersonal and a little disappointing. I absolutely get why they are doing that, but it still doesn't feel great.
If I don't enjoy being on the receiving end of AI-generated messages, why would I do that to my users? Even if I tried to coax the agent into a more natural speaking style for me, that still isn't my words and thoughts - and I'm generally of the opinion of "if you couldn't be arsed to write that, why should I read it?"
My co-maintainer Marcin recently has done so much great work on oapi-codegen to get us into a better place, and has managed to clear a huge amount of the overwhelming backlog. That was only possible because of his leaning on Anthropic's Claude, and I'm really glad he was able to do so - but at the same time, that also led to some surprising scenarios for some of our users who felt it was "a bot" doing the work.
Although I aim to attribute my usage of AI in my codec changes and interactions on PRs/Issues/etc, I'm still not sure that someone knowing it's via AI would make it any better - I don't really feel better knowing an AI agent is replying to me instead of a person.
However, if it reads like a user is replying to me with an AI agent, is it worth me spending my time to reply to them, or should I also have an agent replying to them?
Given we only have so many keystrokes left , I'm thinking about how I want to try and optimise my time, while also not losing an aspect I appreciate about myself - my human empathy for my users, and theirs with mine. Sometimes as humans we get frustrated, or reply in the few minutes before bed, and that's important too, instead of sanitising it through an AI model.
And most importantly, to Open Source maintainers everywhere, please let me know what you're thinking about this all and how you're handling it.
Are you a maintainer who's using AI to help maintain your project and/or your mental health? What's working for you? What don't you recommend?
Are there tools like automated code review tools or tools for handling "here's how you may be able to solve your issue" tools you're finding useful?
How much are you using in terms of authoring/actually responding to users vs using your own words?
If folks are using AI to interact (including in terms of their replies to you), are you leaning on using AI to reply to them? I.e. "fight AI with AI"?
This post's permalink is https://www.jvt.me/posts/2026/08/02/ai-maintainer/ and has the following summary:
Thinking about the onslaught of AI-generated contributions, and if there is space for maintainers to "fight fire with fire".
The canonical URL for this post is
https://www.jvt.me/posts/2026/08/02/ai-maintainer/
.
Written by Jamie Tanna
on Sun, 02 Aug 2026 20:39:42+01:00 , and last updated on Sun, 02 Aug 2026 20:39:42+01:00 .
Content for this article is shared under the terms of the Creative Commons Attribution Non Commercial Share Alike 4.0 International , and code is shared under the Apache License 2.0 .
🤖 Content in this blog post (prose or code snippets) includes code derived from the following LLMs:
Has this content helped you? Did it solve that difficult-to-resolve issue you've been chasing for weeks? Or has it taught you something new you'll be able to re-use daily?
Please consider supporting me so I can continue to create content like this!
This post was filed under articles .
Below you can find the interactions that this page has had using WebMention .
Have you written a response to this post? Let me know the URL:
Do you not have a website set up with WebMention capabilities? You can use Comment Parade .
Other links that may be of interest:
Follow This Blog (i.e. via RSS feed)
This site is part of an IndieWeb Webring 🕸💍
Want to get in touch? Drop me an email at hi@jamietanna.co.uk. I try to get back to everything, if I don't reply in a couple of weeks, feel free to give me a nudge!
You can also check out a list of all the places /elsewhere you can find me.
