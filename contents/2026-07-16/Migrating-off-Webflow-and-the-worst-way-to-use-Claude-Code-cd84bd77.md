---
source: "https://www.sisyphus.bar/blog/migrating-off-webflow"
hn_url: "https://news.ycombinator.com/item?id=48934161"
title: "Migrating off Webflow and the worst way to use Claude Code"
article_title: "Migrating off Webflow and the worst way to use Claude Code — Sisyphus Bar and Grill"
author: "itunpredictable"
captured_at: "2026-07-16T13:20:29Z"
capture_tool: "hn-digest"
hn_id: 48934161
score: 1
comments: 0
posted_at: "2026-07-16T13:19:42Z"
tags:
  - hacker-news
  - translated
---

# Migrating off Webflow and the worst way to use Claude Code

- HN: [48934161](https://news.ycombinator.com/item?id=48934161)
- Source: [www.sisyphus.bar](https://www.sisyphus.bar/blog/migrating-off-webflow)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T13:19:42Z

## Translation

タイトル: Webflow からの移行とクロード コードの最悪の使用方法
記事のタイトル: Webflow からの移行とクロード コードの最悪の使用方法 — Sisyphus Bar and Grill
説明: テクノロジー、ベンチャーキャピタル、料理、執筆、その他の岩を丘に転がす方法についてのジャスティンのメモ。

記事本文:
Webflow からの移行とクロード コードの最悪の使用方法 — Sisyphus Bar and Grill Home Writing About Louis-Jean-François Lagrenée, La Mélancolie, 1785 Webflow からの移行とクロード コードの最悪の使用方法
Webflow がいかにひどいかを詩的に表現することもできます。 CMS での下書きの仕組みからサポート チケットの流れに至るまで、あらゆる細かい点がひどく不快なほど悪いため、考えられる結論は 1 つだけです。「彼らはこれを意図的に行っている」ということです。これほど仕事が下手なチームはないでしょう。この製品が、何か役に立つものを作ろうという彼らの真の努力の不幸で偶然の結果であるということは、まったく考えられません。
いいえ、その代わりに、鋭い歯と血のように赤い目をした邪悪な地下室の住人たちがスクリーンの周りに集まり、狂乱の騒ぎを起こしている様子を想像するのが好きです。彼らは、くすくすと笑いながら、ユーザーを敵に回すための、ますます賢くてイライラする方法をブレインストーミングします。ユーザーの CMS 編集を保存するかどうかについては、ちょっとした問題にしてください。ハイライトされたテキストだけでなく、すべてのテキストを H2 に切り替えます。はい…AI チャットボットと会話した後、ユーザーにサポート リクエストを 2 回入力してもらいます。そして、グランド クレッシェンドでは、72 時間の停止をしましょう !!!!
確かに彼らは優秀なエンジニアです。彼らのゲームに関する知識は非常に完全かつ高度であるため、物事を間違って構築する方法を正確に知っています。そして彼らは、（Webflow Agency Industrial Complex のおかげで）何百万もの投獄されたユーザーに引き起こすであろう日常の小さな苦痛を楽しみながら、喜んでそうします。
ジョン・ミルトンは『失楽園』で次のように述べています。
天国で奉仕するよりも、地獄で君臨するほうが良い。
とにかく、Webflow がいかにひどいかについて何年も不平を言ってきましたが、ついに Amplify Web サイトを Webflow からカスタム セットアップに移行する時期が来たと感じました。
新しい

スタックと移行のロジスティクス
これは、Amplify 用に構築した内部アプリの単純な Web スタックです。
クリーンな CSS、Tailwind が嫌い、LLM が私に CSS の使用を要求するのが嫌い
特別なことは何もありませんが、繰り返しになりますが、私は単なる地の塩です。
Webflow ではこれが簡単に実現できないことを知っておくべきでした。単に「コードをエクスポート」して (ハッ! ハハ!)、それを新しいプラットフォーム用にモデルでリファクタリングするほど簡単ではないことを知っておくべきでした。
もちろん、コードをエクスポートするには、Webflow の有料プランが必要です。もちろん。
それは何だと思いますか？有料プランはありますか？すでに Webflow に年間 1,000 ドル以上を費やしていますか?申し訳ありませんが、私たちは 6 か月ごとにプランを変更しており、あなたは曾祖父の SKU を利用していることになりますが、その SKU はもう存在せず、弊社のエンジニアの 1 人が 6 か月かけて書いたコードでは「有料」としてカウントされません。
残念ながら、私たちはお客様の忠実な複数年サブスクリプションに報いる代わりに、お客様の生活を可能な限り困難にする必要があります。大丈夫だといいのですが。
Webflow を嫌うことについて知っておくべき重要なことは、あなたは決して一人ではないということです。私の場合、優秀な紳士が、私が困ったときに役立つソフトウェア Exflow を作成してくれました。したがって、クロードを支援するために、少なくとも自分のサイトの部分的なシミュラクラムをエクスポートすることができました。
ここでの作業はそれほど魅力的なものではありませんでした。ルートごとに作業を進め、Claude Code (Sonnet 4.6、私は倹約家です) に第一原則、つまり既存の公開済み Web サイトから根気よくページを再構築してもらいました。
クロードがそのページを最初に通過したときは、オリジナルと漠然と似ていますが、明らかにいくつかの幻覚が含まれており、特にフォントのサイズや太さ、さらには境界線の半径などに関して幻覚が見られました。
私のプロンプトのほとんどは次のようになります。
HTMLとCSをダウンロードする

このページを編集し、ローカル サイトに再構築します。スタイル、形状、境界線、フォントを正確に一致させます。 Playwright 統合を使用して作業を再確認します。
そういえば…Claude Code Playwright の統合はあまりうまく機能しませんでした。みんながループが必要だと言い続けたので、クロードに「ループ」を与えようとしたのですが、
Playwright の使用とスクリーンショットにはそれぞれ 3 分以上 (?) かかりました。
頻繁に「ループ」にはまってしまいます (結局のところ、ループから抜け出すにはどうすればよいのでしょうか)
…最終的な結果は、私が手動で修正点を指摘するよりも優れたものではありませんでした
したがって、私は数日かけてページごとに次のようなプロンプトを作成しました。
pタグのフォントサイズは1.2remである必要があります
ボタンの境界線はまだ黒です。修正は機能しませんでした
「Amplify について」CTA を 50 ピクセル下に移動します
時間が経つにつれて、モデルは実際に大幅に改善されました。サイト全体に、コピーして貼り付けるだけで使用できる優れたコンポーネントのライブラリがあったからです。しかし、これでも完璧に近いものは得られませんでした。 「このボタンを他のボタンとまったく同じように見せる」などのことを継続的に言う必要があり、そこで私は基本的に第一原理からデザイン システムを発見しました。
フロントエンドに LLM を使用するのはアンチパターンで最悪です
要点を理解してもらうためにこれらの詳細を説明しましたが、要点は、これは非常に悪い時期だったということです。魔法のような瞬間も、笑顔も、印象的な成果も、歓喜の感情さえもありませんでした（少しの歓喜さえありませんでした）。実際のところ、それは信じられないほど煩わしい数日間で、集中力が分散し、フォローが不十分でイライラし、そもそもなぜこれをやろうと思ったのか、Webflow 刑務所に留まるべきだったかもしれない、スープとパンがある、結局のところ、それは本当に悪いことなのでしょうか？ (はい)
ここでこのブログ投稿の要点に戻ります。 2026年は再

このようなプロジェクトに LLM を使用するのは時期が悪いです。もしあなたにセンスがあり、フロントエンド作業に LLM を使用したい場合は、必然的に、両者がミットの側面でお互いを軽くかすめ続けるモデルとのボクシングの試合にはまり込むことになりますが、決して離れることもできません。今日の LLM は、プロジェクトの最後の 20% を理解していません。そして、フロントエンドは視覚的で細部まですべてが重要な性質を持っているため、それに LLM を使用するのはアンチパターンであり、最悪です。
実際、途中からまた CSS を手動で書き始めたところです。プロンプトを作成し、モデルに考えさせ、成功率 50% でモデルがゴールにシュートを打つのを待つのにかかる時間は、昔ながらの方法でフロントエンド エンジニアのコスプレをするのにかかる時間よりも長かったのです。
たぶん私はただのバカです。おそらく、まともなエンジニアならこのような方法で LLM を使用することはないだろうし、LLM スクールでこれを教えてくれなかっただけで、もちろんこれはひどいアイデアでした。おそらく私は、フロントエンドの小さな詳細に気づくという不幸な能力に呪われているだけかもしれませんが、他の愚か者は、境界線の半径がソースサイトとまったく同じでなくても大丈夫でしょう。おそらく未来は、誰もこうした詳細を気にすることすらなくなる時代になるでしょう…世の中に存在するクロードスロップの量を見れば、この結論に至るのは当然かもしれません。
ただし、わかっていることが 1 つあります。次回 Web サイトを構築するときは、おそらくフロントエンドの詳細に LLM を使用しないでしょう。これらのおかげで、スキャフォールディング、バックエンド作業、基本的なルートとページの設定にかかる時間を大幅に節約できます。しかし、彼らは詳細を理解していません。そして、私は、「必要なものを入力」し続けても結果が得られないよりも、CSS ルールを手動で書いて ~~1 回目~~ 3 回目で正しいものにしたいと思っています。
追伸サイトはここにあります。いろいろ調べて、問題が見つかった場合はお知らせください。

## Original Extract

Notes from Justin on technology, venture capital, cooking, writing, and other ways of rolling the rock up the hill.

Migrating off Webflow and the worst way to use Claude Code — Sisyphus Bar and Grill Home Writing About Louis-Jean-François Lagrenée, La Mélancolie, 1785 Migrating off Webflow and the worst way to use Claude Code
I could wax poetic about how bad Webflow is. Every small detail, from the way drafts in their CMS work to their support ticket flow, is so egregiously and offensively bad that there’s only one possible conclusion: they are doing this on purpose. No team could be this bad at their jobs. It’s simply not possible that this product is the unfortunate and accidental result of their genuine efforts to build something useful.
No, instead I like to imagine a group of evil basement dwellers with sharp teeth and blood red eyes, gathered around a screen and working themselves into an orgiastic frenzy. As they cackle and roar, they brainstorm increasingly clever and frustrating ways to antagonize their users. Make it a crapshoot as to if we save the user’s CMS edits…no wait, make ALL of the text switch to an H2, not just the highlighted text. Yes…now make the user type their support request TWICE after talking to our AI chatbot! AND FOR THE GRAND CRESCENDO, LET’S HAVE 72 HOUR OUTAGES!!!!!
They are cracked engineers to be sure; their knowledge of the game is so complete, so advanced, that they know exactly how to build things wrong . And they do so with delight, reveling in the small daily pains they’ll cause their millions of imprisoned users (thanks to the Webflow Agency Industrial Complex).
As John Milton said in Paradise Lost:
Better to reign in Hell than serve in Heaven.
Anyway, after many years of complaining about how bad Webflow is, it finally felt like the time to migrate the Amplify website off of it and onto a custom setup.
The new stack and migration logistics
Here is my simple web stack for the internal apps I’ve been building for Amplify:
Clean CSS, I hate Tailwind and hate that LLMs want me to use it
Nothing fancy, but then again, I’m just a salt of the earth guy .
I should have known that Webflow wasn’t going to make this easy. I should have known that it wouldn’t be as simple as just “exporting the code” (hah! haha!) and having a model refactor it for a new platform.
Of course, to export your code, you must have a paid plan on Webflow. Of course.
What’s that, you say? You have a paid plan? You’re already spending more than $1K/year on Webflow? Sorry broh, but we change our plans every 6 months and you’re on a great-great-great grandfathered SKU that doesn’t exist anymore and doesn’t count as “paid” in the code that one of our engineers spent 6 months writing for this.
Instead of rewarding your loyal multi-year subscription, we are unfortunately going to have to make your life as hard as we can. Hope that’s alright.
The important thing to know about hating Webflow is that you’re never alone. In my case a fine gentleman created just the piece of software to help me in my time of need: Exflow . So I was at least able to export some partial simulacrum of my site to aid my Claudes.
The work here was not very glamorous, it involved going route by route and having Claude Code (Sonnet 4.6, I’m frugal) patiently rebuild the page from first principles i.e. the existing published website.
Claude’s first pass at the page would look like, vaguely similar to the original but clearly involved several hallucinations, in particular around things like font sizes and weights and then border radii as well.
Most of my prompts looked something like this:
Download the HTML and CSS of this page and rebuild it on our local site. Match the styles, shapes, borders, and fonts exactly. Double check your work using the Playwright integration.
Speaking of which…the Claude Code Playwright integration didn’t work very well. Everyone kept telling me that we need loops, LOOPS!, and so I tried to give Claude a “loop” but:
Each Playwright use and screenshot took like 3+ minutes (?)
It would get stuck in the “loop” (after all how do you get out of a loop) often
…the end result wasn’t any better than me manually pointing out fixes
Ergo, I spent a few days going page by page writing prompts like:
p tag font size should be 1.2rem
the border on the button is still black, your fix didn’t work
move the “About Amplify” CTA 50px lower down
Over time, the model actually got a lot better: because it had a library of good components across the site to just copy and paste from. But even this didn’t yield anything close to perfect. I needed to continually say things like “make this button look exactly like the rest of the buttons” – in which I essentially discovered a design system from first principles.
Using LLMs for frontend is an antipattern and it sucks
I bore you with these details to get a point across, and that point is that this was a very bad time . At no point was there a magic moment, a smile, an impressive output, or even a feeling of glee (not even a bit of glee). Actually, it was an incredibly annoying few days of scattered attention, frustration at poor prompt following, and consistently asking myself why I had decided to do this in the first place, maybe I should just have stayed in the Webflow prison, they have soup and bread, after all is it really that bad? (yes)
Which brings me to the point of this blog post. 2026 is a really bad time to be using LLMs for projects like this. If you have good taste and want to use an LLM for frontend work, you will inevitably be stuck in a boxing match with a model where you both keep lightly grazing each other with the sides of your mitts but also you can never leave. LLMs today do not understand the last 20% of any project. And because of the visual and all-about-the-details nature of frontend, using LLMs for it is an antipattern and it sucks.
In fact, halfway through I just started manually writing CSS again. The amount of time it took to write a prompt, have the model think, and then wait for it to make a shot on goal with a 50% success rate was more time than it took to just go cosplay as a frontend engineer the old fashioned way.
Maybe I’m just an idiot. Maybe no engineer worth their salt would ever use an LLM this way, and they just didn’t teach me this in LLM school and of course this was a terrible idea. Maybe I’m just cursed with this unfortunate ability to notice tiny frontend details, while other idiots would be fine if the border radius isn’t exactly the same as the source site. Maybe the future is one where nobody even cares about these details anymore…the amount of Claudeslop out there could reasonably lead you to this conclusion.
There’s one thing I do know, though. Next time I build a website, I probably won’t use LLMs for frontend details. They save me a lot of time with scaffolding, any backend work, and setting up basic routes and pages; but they just don’t get the details. And I’d rather write CSS rules by hand and get them right the ~~1st~~ 3rd time than continually “type what I want” and not get it.
P.S. the site is live here . Please poke around and let me know if you find any issues.
