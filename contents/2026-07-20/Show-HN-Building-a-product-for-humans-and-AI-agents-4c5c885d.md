---
source: "https://slobodan.me/posts/introducing-competitor-tracker/"
hn_url: "https://news.ycombinator.com/item?id=48979190"
title: "Show HN: Building a product for humans and AI agents"
article_title: "Building Competitor Tracker - a product origin story in the age of AI :: Slobodan Stojanović — CTO and co-founder of Vacation Tracker, AWS Serverless Hero and co-author of Serverless Applications with Node.js book"
author: "slobodan_"
captured_at: "2026-07-20T15:02:27Z"
capture_tool: "hn-digest"
hn_id: 48979190
score: 1
comments: 0
posted_at: "2026-07-20T14:15:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Building a product for humans and AI agents

- HN: [48979190](https://news.ycombinator.com/item?id=48979190)
- Source: [slobodan.me](https://slobodan.me/posts/introducing-competitor-tracker/)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T14:15:59Z

## Translation

タイトル: Show HN: 人間と AI エージェントのための製品の構築
記事のタイトル: 競合他社トラッカーの構築 - AI 時代の製品誕生の物語 :: Slobodan Stojanović — CTO 兼 Vacation Tracker の共同創設者、AWS Serverless Hero、著書『Serverless Applications with Node.js』の共著者
説明: 急激な変化は時々恐ろしいものです。 AI を扱うのは楽しいですが、すべての変化についていくのは不可能なので、ストレスがかかることもあります。朝Twitterを開くと、昨日のモデルよりもはるかに優れた新モデルを見逃してしまい、軽いパニック発作を引き起こす可能性があります、とAnthropicが発表
[切り捨てられた]

記事本文:
競合他社トラッカーの構築 - AI 時代の製品誕生の物語 :: Slobodan Stojanović — CTO 兼 Vacation Tracker の共同創設者、AWS サーバーレス ヒーロー、著書『Serverless Applications with Node.js』の共著者
>
スロボダン
について
競合他社トラッカーの構築 - AI 時代の製品誕生の物語
急激な変化は時に怖いものです。 AI を扱うのは楽しいですが、すべての変化についていくのは不可能なので、ストレスがかかることもあります。朝、Twitter を開くと、昨日のモデルよりもはるかに優れた新しいモデルを見逃したため、軽いパニック発作が起こる可能性があります。Anthropic は、神話上のモデルを使用するための新しい期限、製品を一発で構築する新しいプロンプト、ツール、ベンチマークなどを発表しました。すべてが非常に速く進んでいるように見え、自分がバブルの中にいることを忘れてしまいます。明日はじけるバブルではなく、アーリーアダプターのバブルだ。
この毎日の更新の洪水はほとんどノイズをもたらすことが多いですが、AI と LLM の効果、利点、欠点は現実のものです。彼らはすでに多くのことに影響を与えています。製品、特にソフトウェアの構築方法の変更も含まれます。これらの変化の当然の結果の 1 つは、おそらく今が独自の製品を構築するのに最適な時期であるということです。最も楽な時期ではありません。それは存在しません。製品の構築はハード モードでゲームをプレイすることです。
私の共同創設者である Lav と私が最近構築した製品の話をさせてください。複数のエージェントがマーケティング サイトのコードと新しいページを作成して公開している間に、長文のテキストを手書きするのは少し奇妙ですが、すべてのストーリーが AI によって語られるわけではありませんし、語られるべきでもありません。
AI は製品構築のボトルネックを変更しましたか?
何年も前に素晴らしい本『ザ・ゴール』が私たちに教えてくれたように、すべての組織と何かを構築するには、主な制限要因が 1 つあります。ブーイング

k は制約理論 (TOC) を導入し、主な制限要因 (ボトルネック) が許す限り速く進むことができることを説明しています。
製品を構築することは常に困難であり、多くのことが実際のボトルネックとなってきました。アイデアから始まり、投資 (製品の構築に費やした時間と資金)、好みと構成 (何をどのように構築するか)、実際の製品の構築 (ソフトウェア開発、設計、テストなど)、そして市場投入の全体の動き (製品を誰に販売するか、実際にその製品にお金を払う人をどのように見つけるか) に至るまでです。しかし、各組織にとってボトルネックとなるのは、これらのうちの 1 つだけです。たとえば、一部の組織では、開発者、デザイナー、またはテスターを追加したり、バックログから項目を削除したり、製品やプロジェクトの管理方法を変更したりして、機能をより迅速に出荷して運用パイプラインを高速化し、運用ローカルのボトルネックの制限を緩和することができます。しかし、たとえそうであっても、市場開拓が依然としてボトルネックとなる可能性があります。あるいは、毎週わずか数時間しか時間がなく、他の人を雇う予算もなく、それ以上の時間を確保する方法もないかもしれません。
実際のボトルネックは組織ごと、また製品ごとに異なります。しかし、最も一般的なものを選択する必要があるとすれば、それはおそらく市場投入でしょう。製品を構築するのは難しいですが、それを販売するのは、幸運なマーケティングと販売のマスターでない限り、非常に困難です。
では、AI によってボトルネックが変化したのでしょうか?
答えはおそらく「はい」ですが、ボトルネックによっては、それが良い知らせになることもあれば、悪い知らせになることもあります。説明しましょう！
実際のボトルネックが開発時間の不足だった場合、朗報です。 AI は、コードの作成、アプリの設計、マーケティング コピーの作成、さらにはアプリのテストを支援します。
私のお気に入りの法則の 1 つは、現在 30 年になるトグの通勤法則です。

ith テスラーの複雑さ保存の法則: 「通勤時間は固定されています。変化できるのは距離だけです。」現在の言葉で言えば、コードを書いて機能をより速く出荷できるようになれば、以前と同じ量のコードや機能を使って終わりになることはなくなります。代わりに、より多くの機能やコードをリリースするか、他のものを構築することに同じ時間を費やすことになります。
それは悪い知らせではありません。もっと自分でできることがあります。しかし悪いニュースは、他の人もそうなる可能性があるということです。そのため、適切な潜在顧客にリーチし、すべてのノイズや実際の競合他社を排除するのは難しいため、市場開拓戦略はさらに大きなボトルネックになります。
わかりましたが、製品の構築は速くなりましたね? Twitter のあの人が Claude Fable 5 を使ってアプリケーション全体をワンショットで実行したとしたら、誰でもそれができるでしょう?いや、そうではありません。素晴らしいデモアプリと製品の間には大きな違いがあります。しかし、最初に私たちのバックストーリーをお話しさせてください。
前回の記事「従業員、AI、および AI 従業員」で述べたように、3 年以上前、私の共同創設者である Lav と私は AI 共同創設者 (CofounderGPT と呼んでいました) を作ろうとしましたが、とても楽しかったですが、失敗しました。正直に言うと、それはもう 100% 真実ではありません。過去 6 か月間、CofounderGPT という AI エージェントが協力してきましたが、それが新製品の構築に大きな役割を果たしましたが、それについては後ほど説明します。それまでの間、興味がある場合は、https://cofoundergpt.ai で CofounderGPT の週次ログをチェックしてください。しかし、「原点」の話に戻りましょう。
すべては私たち自身の AI 実験から始まりました。しかし、当時は AI を使って面白いものを作るのに費やす時間が最大のネックだったので、人を雇って「研究開発部門」を立ち上げました。 Tw のクールなデモのいくつかを再構築できたので、それは楽しかったです。

それを試したり (そして、そのほとんどはデモや手品としてのみクールでした)、さまざまなモデルを試し、いくつかのモデルを微調整することさえできました。
最後に、私たちは CofounderGPT に取り組み、AI を備えた小規模な製品の構築を開始しました。特に LLM のような新興市場では、製品構築を実質的にアウトソーシングすることができないため、どちらも失敗しました。また、私たちは忍耐力が足りませんでした。私たちはより良いモデルを望んでいたのですが、その時点ではすでに手遅れだと考えていました。
私たちは今後も AI を活用した製品を作りたいと考えています。しかし、それを絞り込む必要がありました。 CofounderGPT 全体を構築する代わりに、1 つの小さなコンポーネントを構築したらどうなるでしょうか?私たちの問題を解決し、時間を節約し、より良い洞察を与えてくれる何か。多くのアイデアがありましたが、その 1 つは、私たちが抱えていた厄介な問題に焦点を当てていました。誰かが直接および間接的な競合他社の Web サイトをチェックし、最も重要な変更を提示し、巨大で常に古いスプレッドシートを管理する必要がありました。
AI にとってなんと素晴らしい仕事でしょう。そしてそれは簡単なはずです。有名な最後の言葉。
他の AI 実験の中でも特に、複数の人員でこれに数か月取り組んだ後、結果は残念なものでした。このツールは機能しましたが、ノイズを拾いすぎていました。文字通り、ノイズを除去することが、ツールが完璧に実行する必要がある主な仕事です。結局、「研究開発部門」は閉鎖され、ほぼ断念しました。
幸運なことに、誰かが私たちに、AI のみに焦点を当てている若くて賢い人々のチームを紹介してくれました。彼らはこのツールを構築できると言い、実際には 2 か月以内に納品しました。私たちはその中核機能に焦点を当て、不要なものはすべて削除することにしました。しかし、2025 年 9 月には実用的なバージョンが完成しました。
アプリの重要な部分は、毎週月曜日に受信するメールでした。これは、あなたの会社のすべての重要な変更について知らせる、カスタマイズされた「ニュースレター」です。

過去 7 日間に参加した参加者。さらに、競合他社を追加すると、インターネット ウェイバック マシンから自社の Web サイトの古いバージョンを取得しようとするため、1 週間以上待たずにプレビューを見ることができるというすばらしい機能も追加されました。私たちは Vacation Tracker の競合他社を設定し、数人の友人を招待してテストしてもらいました。その後、私たちはそのことをほとんど忘れていました。
時間が経つにつれて、優先事項が新製品をリストの下に押し下げました。しかし、私たちは月曜日のメールを読み続けました。最終的に、Lav と私はなんとか時間を確保し、いくつかの新しい製品を構築し、AI だけを使用して製品を構築および保守できるというアイデアを再テストすることにしました。私たちは遊び場を必要としていたのですが、Competitor Tracker は完璧な遊び場でした。
思ったよりも時間がかかりました。最初のプロンプトから、ブランドの声とスタイルに従い、インデックスに登録された Web サイトを備えた製品が完全に機能するまで 2 か月かかりました。私たちを遅らせた主な問題は、時間、数多くの予期せぬ障害 (クロール、さまざまなエラー、タイムアウト、ノイズなど)、そして AWS ほど馴染みのないプラットフォームである Cloudflare の使用です。これらの選択肢については、以下で詳しく説明します。
競合他社トラッカーの紹介
ついに新製品、Competitor Tracker をご紹介できることになりました。
Competitor Tracker は、競合他社の追跡に役立つアプリケーションです。当然のことですが、私たちは名前に関してあまり創造的ではありません。実際、少なくとも私たちの場合、最適な名前を​​見つけることは MVP の一部ではありません。必要に応じて、後でそれを理解することができます。
しかし、Competitor Tracker (ブランド上では Competitor Tracker & Co.) は、競合他社を追跡する単なるアプリケーションではありません。 AI によって世界が変化したため、私たちは人間にも AI にも同様に優しい製品を構築しました。
まず、すべてが API です (詳細なスコープベースの権限を備えています)。 OAuth2 または API キーを使用して API を使用できますが、ab

絶対にすべて API を使用します。だからといって製品が AI 対応になるわけではありませんが、すべてが API と詳細なドキュメントから始まるのであれば、正しい方向に進んでいることになります。
次に、Web アプリと MCP はその AI を使用して、すべての機能を人間とエージェントに公開します。 Web ダッシュボードから実行できることはすべて、エージェントは MCP 経由で実行できます。例外はありません。
Competitor Tracker は毎週月曜日にダイジェストメールを送信します。ただし、エージェントは Webhook を取得するため、エージェントに受信トレイへのアクセスを許可する必要はありません。同じデータ、同じフィルター、同じ時間。
最後に、Web サイトは人間とエージェントのために構築されます。エージェントにはマークダウンを、人間には HTML を提供します。試してみたい場合は、各ページの下部にマトリックスに入るスイッチがあります。申し訳ありませんが、「AI エージェント」ビューを確認するのに役立ちます。
人間とエージェントのほかに、私たちには犬もいます。C.T. ラッキーは私たちの新しいノワール プラットフォームを通してあなたを案内し、月曜日に「書類」を送ってくれる私たちの主任刑事です。 AI犬ですが、それでも良い犬です。
それでは、Competitor Tracker & Co. をご紹介します。この AI 代理店のエージェントのパックは、競合他社の公開ページを 1 週間追跡し、毎週月曜日に 1 つの短いランク付けされた文書をファイルします。これにより、チームは別のアラート フィードを子守する代わりに、重要な動きを読み取ることができます - https://competitortracker.io 。
C.T. Lucky と Competitor Tracker & Co. に会いたい場合は、「Competitor Tracker & Co とは」にアクセスしてください。ページで、エージェントのパックが何ができるかを確認するか、製品デモ (人間とエージェント用) をチェックして試してみてください (セットアップには数分かかります。最初の 25 コインは家にあります)。
昨日、素晴らしいツイートを目にしました: https://x.com/krishdotdev/status/2077723982028210381 。
これは、これまで AI を使用して製品を構築してきた私の経験とほぼ完全に一致しています。アイデアが必要です

息を呑むようなデモには数時間かかります。残りの 10% を完了するには数か月かかります。もちろん、これは本当の 10% ではなく、すべてをほぼ瞬時に実行できると思わせる幻想です。
私はさらに一歩進んで、デモには数時間、コードの 200% を実装するには数週間、そして AI の残骸をすべてクリーンアップしてその一部を動作する機能に置き換えるには数か月かかると言います。
画像に基づく: https://x.com/krishdotdev/status/2077723982028210381 。
Competitor Tracker でも同様のことが起こりました。しかし、それは旅の楽しい部分ではありませんでした。
個人的に楽しかったのは、すべてをセットアップすることです。
そしてすべてはブランドの決定から始まりました。いくつかのスタイルを試してみましたが、フィルム ノワールのスタイルが一番合っているように思えました。同時に、クロード・コード社にプロトタイプのソースコードに基づいた初期ドキュメントの準備を依頼しました。
次に、Lav が Hermes エージェントと OpenClaw エージェントをリポジトリに接続することがわかっているため、製品リポジトリとマーケティング Web サイトのリポジトリを分離することにしました。誤解しないでください。私は Hermes と OpenClaw の両方が好きですが、完全な権限セットを持つ AI エージェントを信頼しません。
次のステップはプラットフォームを持ち上げることでした。私もよく知っているので、AWS を選択するのは当然でした。しかし、明らかな選択は面白くありません。プロセスをより楽しくするために

[切り捨てられた]

## Original Extract

Rapid changes are sometimes scary. While working with AI is fun, it can also cause stress because it’s impossible to keep up with all the changes. Opening Twitter in the morning can lead to a mini panic attack because you missed new models that are way better than yesterday’s ones, Anthropic announc
[truncated]

Building Competitor Tracker - a product origin story in the age of AI :: Slobodan Stojanović — CTO and co-founder of Vacation Tracker, AWS Serverless Hero and co-author of Serverless Applications with Node.js book
>
slobodan
About
Building Competitor Tracker - a product origin story in the age of AI
Rapid changes are sometimes scary. While working with AI is fun, it can also cause stress because it’s impossible to keep up with all the changes. Opening Twitter in the morning can lead to a mini panic attack because you missed new models that are way better than yesterday’s ones, Anthropic announced new deadlines for using their mythical models, new prompts that build products in one shot, tools, benchmarks, etc. Everything seems to be moving so fast that you forget you are in a bubble. Not the one that will burst tomorrow, but a bubble of early adopters.
While this flood of daily updates often brings mostly noise, the effects, benefits, and downsides of AI and LLMs are real. They have already affected many things. Including changing the way we build products, especially software. One of the natural consequences of these changes is that now is probably the best time to build your own product. Not the easiest time. That does not exist. Building products is playing the game on hard mode.
Let me tell you a story of a product that my co-founder, Lav, and I recently built. It’s a bit weird to write a long-form text by hand while multiple agents write and publish the code and new pages for our marketing site, but not all stories can and should be told by AI.
Did AI shift the bottleneck for building products?
As the amazing book The Goal told us many, many years ago: there’s one main limiting factor for all organizations and for building anything. The book introduces the Theory of Constraints (TOC) , which explains that we can move as fast as our main limiting factor ( the bottleneck ) allows.
Building products has always been hard, and many things have been real bottlenecks. Starting from the idea, then your investment (time and money spent building the product), taste and organization (what to build and how), actual product building (software development, design, testing, etc.), to the whole go-to-market motion (who do you sell the product to and how do you find people that would actually pay for it). But only one of these is the bottleneck for each organization. E.g., some organizations can add developers, designers, or testers, remove items from the backlog, or change how they manage products and projects to ship features faster and speed up the production pipeline, and make their production local bottleneck less restrictive. However, even with that, go-to-market may still be a bottleneck for them. Or you might have limited time to just a few hours each week, and no budget to hire anyone else, and no way to free up more time.
The actual bottleneck varies from organization to organization and from product to product. But, if we need to pick the most common one, it’s probably go-to-market. Building products is hard, but selling them - unless you are one of the lucky marketing and sales masters, it’s extremely hard.
So, did AI shift the bottleneck?
The answer is probably yes, but depending on your bottleneck, that could be good or bad news. Let me explain!
If your actual bottleneck was a lack of time for development, I have good news for you! AI can help you write code, design the app, write marketing copy, and even test the app.
One of my favorite laws is the now 30-year-old Tog’s Law of Commuting, which pairs with Tesler’s Law of Conservation of Complexity: “The time of a commute is fixed. Only the distance is variable.” To put it in current terms: if we can write code and ship features faster, we won’t stop with the same amount of code/features as before and go fishing. Instead, we’ll spend the same time either shipping more features or code, or building other things.
That’s not bad news - you can do more by yourself! But the bad news is that others can too. So, the go-to-market strategy becomes an even bigger bottleneck, as it’s hard to reach the right potential customers and cut through all the noise and actual competitors.
Ok, but building products is faster now, right? If that guy on Twitter one-shotted the whole application using Claude Fable 5, everyone could do that, right? Well, not really. There’s a big difference between an amazing demo app and a product. But let me tell you our backstory first.
As mentioned in my previous article, “Employees, AI, and AI employees,” more than 3 years ago, my cofounder, Lav, and I tried to create an AI cofounder (CofounderGPT, as we called it), had a lot of fun, and failed. Well, to be honest, that’s not 100% true anymore, as we have had an AI agent called CofounderGPT working with us for the last 6 months, and it played a big role in building our new product, but we’ll talk about that later. In the meantime, if you are curious, check https://cofoundergpt.ai for CofounderGPT’s weekly log. But let’s get back to the “origin” story.
Everything started with our own AI experiments. However, because the time we could spend building fun stuff with AI was our main bottleneck at the time, we hired someone and started an “R&D department.” That was fun, as we managed to rebuild some of the cool demos from Twitter (and most of them were cool only as demos and magic tricks), try out different models, and even fine-tune a few.
Finally, we started working on CofounderGPT and building a small product with AI. Both things failed because you can’t really outsource product building, especially not in an emerging market like LLMs. Also, we were not patient enough. We wanted better models, and, at that time, thought that we were already late.
We still want to build a product with AI. But we needed to narrow it down. What if, instead of building the whole CofounderGPT, we built one small component? Something that would solve our problem, save our time, and give us better insights. There were many ideas, but one was focusing on an annoying problem that we had - someone needed to check our direct and indirect competitors’ websites, present the most important changes, and maintain that big and always outdated spreadsheet.
What an amazing job for an AI! And it should be easy. Famous last words.
After months of working on this, among other AI experiments, and with multiple people who worked on it, the results were disappointing. The tool worked, but it was picking up too much noise, and filtering out the noise is literally the main job it needs to do perfectly. Eventually, we closed the “R&D department,” and almost abandoned the idea.
Luckily, someone introduced us to a team of young, smart people who focused solely on AI. They said they could build this tool, and they actually delivered in less than 2 months. We decided to focus on its core function and remove everything nonessential. But we had a working version in September 2025.
The key part of the app was an email that you receive every Monday. It’s a tailored “newsletter” that tells you about all the key changes your competitors made over the past 7 days. They even added a brilliant feature - when you add competitors, they try to get an older version of their website from the Internet Wayback Machine so you can see a preview without waiting a week or longer. We set up Vacation Tracker’s competitors and invited a few friends to test it out. Then we almost forgot about it.
As time passed, priorities pushed the new product down the list. But we continued reading these Monday emails. Eventually, Lav and I managed to free up some time and decided to build some new products and retest the idea that we can build and maintain a product with AI alone. We needed a playground, and Competitor Tracker was a perfect one.
It took longer than I expected. Two months from the first prompt to a fully working product with a website that follows the brand voice and style and is indexed. The main problems that slowed us down are time, numerous unexpected roadblocks (crawling, various errors, timeouts, noise, etc.), and the use of Cloudflare, a platform I’m less familiar with than AWS. More on these choices below.
Introducing Competitor Tracker
We can finally introduce our new product: Competitor Tracker!
Competitor Tracker is an application that helps you - track your competitors. We are not very creative with names, obviously. Actually, figuring out the best name is not part of the MVP, at least not in our case. We can figure that out later if needed.
But Competitor Tracker, or, to be on-brand, Competitor Tracker & Co., is more than just another application that tracks competitors. The world changed with AI, so we built this product to be equally human and AI-friendly.
First, everything is an API (with detailed scope-based permissions). You can use the API with OAuth2 or an API key, but you can do absolutely everything using the API. That does not make the product AI-friendly, but if everything starts with an API and detailed documentation, then you are going in the right direction.
Then, the web app and MCP use that AI to expose all the functionality to humans and agents. Everything you can do from the web dashboard, your agent can do via MCP. No exceptions.
Competitor Tracker sends you a digest email every Monday. But agents get a webhook, so you do not need to give them access to your inbox. Same data, same filters, same time.
Finally, the website is built for humans and agents. We serve markdown for agents and HTML for humans. If you want to test it out, there’s a switch at the bottom of every page that helps you enter the Matrix. Pardon, it helps you see the “AI agent” view.
Besides humans and agents, we also have a dog - C. T. Lucky is our lead detective who guides you through our new noir platform and sends you Monday “dossiers.” It’s an AI dog, but still a good one.
So, it’s time - meet Competitor Tracker & Co. : the AI agency whose pack of agents tails your competitors’ public pages all week and files one short, ranked dossier every Monday, so your team reads the moves that matter instead of babysitting another alert feed - https://competitortracker.io .
If you want to meet C. T. Lucky and Competitor Tracker & Co., head to the “What is Competitor Tracker & Co.” page, see what our pack of agents can do , or check out the product demo (for humans and agents) and try it out (it takes a few minutes to set it up, and the first 25 coins are on the house).
I saw an amazing tweet yesterday: https://x.com/krishdotdev/status/2077723982028210381 .
It’s almost perfectly aligned with my experience building products with AI so far. You need an idea and a few hours for a breathtaking demo. Then you need months to finish the last 10%. Of course, that’s not a real 10%, it’s an illusion that tricks you into thinking you can do everything almost instantly.
I would go a step further and say that you need a few hours for a demo, a few weeks to implement 200% of the code, and then months to clean up all the AI slop and replace a small part of it with working features.
Image based on: https://x.com/krishdotdev/status/2077723982028210381 .
Something similar happened to Competitor Tracker. But that wasn’t the fun part of the journey.
Personally, the fun part was setting everything up.
And everything started with a brand decision. I tried out a few styles, and the film noir one looked like the best fit. At the same time, I asked Claude Code to prepare the initial set of documents based on the prototype source code.
Then, because I know that Lav will connect his Hermes and OpenClaw agents to the repo, I decided to separate the product and marketing website repositories. Don’t get me wrong, I like both Hermes and OpenClaw, but I do not trust any AI agent with the full set of permissions.
The next step was picking up the platform. AWS was an obvious choice, as I know it well enough. But obvious choices are not fun. To make the process more fun

[truncated]
