---
source: "https://9to5mac.com/2026/06/08/craig-federighi-details-apples-collaboration-with-google-for-siri-ai-in-ios-27/"
hn_url: "https://news.ycombinator.com/item?id=48463064"
title: "Craig Federighi Details Apple's Collaboration with Google for Siri AI in iOS 27"
article_title: "Craig Federighi details Apple’s collaboration with Google for Siri AI in iOS 27 - 9to5Mac"
author: "tosh"
captured_at: "2026-06-09T18:53:08Z"
capture_tool: "hn-digest"
hn_id: 48463064
score: 1
comments: 0
posted_at: "2026-06-09T16:15:37Z"
tags:
  - hacker-news
  - translated
---

# Craig Federighi Details Apple's Collaboration with Google for Siri AI in iOS 27

- HN: [48463064](https://news.ycombinator.com/item?id=48463064)
- Source: [9to5mac.com](https://9to5mac.com/2026/06/08/craig-federighi-details-apples-collaboration-with-google-for-siri-ai-in-ios-27/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T16:15:37Z

## Translation

タイトル: クレイグ・フェデリギ、iOS 27 の Siri AI における Apple と Google のコラボレーションについて詳しく説明
記事タイトル: Craig Federighi が iOS 27 の Siri AI における Apple と Google の協力について詳しく説明 - 9to5Mac
説明: クレイグ・フェデリギ率いる Apple の Siri チームは、今日の午後、報道陣と WWDC 後の基調講演を開催しました...

記事本文:
メインコンテンツにスキップ
メインメニューを切り替えます
9to5Mac ロゴ 9to5Mac ホームページに移動します
サイトを切り替える
9to5おもちゃ
ダークモードを切り替える
検索:
送信する
検索フォームを切り替え
フォーラム
クレイグ・フェデリギ氏、iOS 27のSiri AIに関するAppleとGoogleの協力について詳しく語る
17件のコメント
クレイグ・フェデリギ率いる Apple の Siri チームは、今日の午後、WWDC 後の基調講演の技術トークを報道関係者と開催し、iOS 27 と新しい Siri AI について話しました。
講演の中で、フェデリギ氏はAppleとGoogleの協力についてさらに詳しく語った。
Federighi 氏には、Amar Subramanya (AI 担当副社長)、Mike Rockwell (Siri リード)、Sebastien Marineau-Mes (ソフトウェア副社長) が加わりました。
Googleとのコラボレーションについて、フェデリギ氏は次のように説明した。
もちろん、私たちのアプリとして Gemini アプリはありません。実際、そのクライアント コードはどれも、iOS での実行方法には含まれていません。これらのモデルについては、Google が顧客にデプロイするモデルは一切使用せず、Google が顧客にモデルをデプロイするインフラストラクチャや手段も使用しません。そして、ナレッジベースに関しては、もちろん Google 検索などをシステムの基盤として使用しません。それが明らかであることを願っています。 Google アシスタントの使用量はまったくありません。
それでは、私たちが実際に何を使用しているか、または私たちのシステムがどのように構築されているかについて話しましょう。
もちろん、それはアシスタントの経験から始まります。そして、今日先ほどご覧になったように、このアシスタント エクスペリエンスは、システム、iOS、iPadOS、macOS に深く統合されています。 iPhone では、ダイナミック アイランドから液体ガラスの中にアシスタントがどのように現れるのか、非常に美しいと思います。サイド ボタンから、または Siri に名前で話しかけることでアシスタントを呼び出す方法を見ました。しかしそれだけではなく、システム内のあらゆる場所に統合されています。したがって、書き込みツールを使用して書いている場合でも、w をクリックすると、

コンテキスト メニューでは、これらすべてがシステム エクスペリエンスに深く統合されています。
さて、そこにSiriアプリが接続されています。 Siri アプリは、以前に開始した会話に戻ったり、以前に行っていたことを確認したり、会話を延長したり、新しい会話を開始したりするのに最適な方法です。しかし、このアプリはクラウド内のモデルにアクセスするだけではありません。これは、Apple Intelligence の強力なシステム ソフトウェア上に構築されています。
これには、システム全体のプライバシー アーキテクチャの鍵となる System Orchestrator が含まれます。これは、アプリ内のアクションへのアクセスを提供するアプリ ツールボックス、リクエストの実行に役立つ個人コンテンツにアクセスするための Spotlight セマンティック インデックス、さらにはリクエストを行っているときに見ているものを理解するための画面上のコンテキストなどに対するリクエストを調整するものです。
これは、一連の強力なオンデバイス モデルに基づいて構築されています。これらは、音声の理解から、あなたに話しかけてくる音声の合成、環境と画面上のコンテキストの視覚的な理解、そこに関連するものがあるかどうかの理解、画面上にある可能性のあるテキストの理解、および他の一連のモデルに至るまで、あらゆることを処理します。
また、一部のリクエストについては、モデルは Siri リクエストを完全にデバイス上でローカルに処理できます。しかし、システム オーケストレーターは、それがより高度な質問であることに気づき、さらに優れたインテリジェンスを利用したいと考えることがあります。これは、プライベート クラウド コンピューティングで実行されているモデルに接続することによって行われます。
プライベート クラウド コンピューティングの目標は、iPhone と同じプライバシーの約束をクラウドに拡張し、リクエストが完全にプライベートになるようにすることです。それらは決して保存されず、誰もアクセスすることはできません。

Apple を含め、それらはリクエストの一部としてのみ処理され、何もアクセスできません。これらのプロパティはすべて、アーキテクチャ的にシステムの奥深くに組み込まれているだけでなく、サードパーティの研究者が継続的に検証できるものでもあります。
現在、その展開モデルには、AFM Cloud および AFM Cloud Pro モデルから AFM Fusion モデルおよびイメージ モデルに至る、Apple Foundation モデルのファミリー (第 3 世代) が含まれています。これらのモデルは Google とのコラボレーションの成果であり、それについては今後さらに詳しく説明します。ただし、これらはデプロイメント アーキテクチャ上で実行できるように設計されています。これらは、Apple Intelligence エクスペリエンスのために特別に設計されたモデルです。これらは、今日の基調講演でご覧になったすべての原動力となっています。
最後に、時事問題や世界の知識のその他の要素に関連するリクエストを行うと、それらの応答は Apple の World Knowledge Service にアクセスすることによって根拠付けられます。これは私たちが長年にわたって構築してきたものであり、お客様のリクエストを満たすための優れた情報源を提供します。
したがって、このシステムは、前に説明した Apple Intelligence エクスペリエンスの全範囲をサポートするものです。それが私たちの議論の根拠になることを願っています。そして私たちが今やりたいのは、これらすべての作業に貢献したリーダーシップチームのメンバーを引き出すことです。
6 人が「いいね！」しました
興味深いコメントで、Apple が Google モデルとサーバーを使用する際のユーザーのプライバシー保護について正確にどのように取り組んでいるかを具体化する素晴らしい仕事をしています。数か月にわたる憶測を経て、詳細がわかってうれしいです。
私たちは、Google と提携した第 3 世代の Apple Foundation Models (AFM) に非常に興奮しています。私たちは、オンデバイスからクラウドまでにわたるモデル ファミリを構築しました。さて、各モデルについて説明する前に、

このファミリーの見出しは、この世代のすべてのモデルが前世代と比較して品質と機能の両方で大幅に向上していると言えます。
各モデルについて説明していきますが、まずオンデバイス モデルから始めます。まず最初に、AFM コアがあります。これは、現在デバイスに搭載されている次世代のオンデバイス モデルです。緻密なアーキテクチャに従っています。
そして次は、AFM Core Advanced です。
これは、これまでに実行したオンデバイス モデルとは異なります。スパースなアーキテクチャを使用しており、ネイティブにマルチモーダルです。その結果、このモデルの機能は大幅に向上し、今朝聞いたいくつかの機能 (招待や表現力豊かな音声など) が可能になり、このモデルの結果としてデバイス上ですべて完全に動作するようになりました。
サーバー モデルに移ります。これらはすべて、プライベート クラウド コンピューティングから提供されます。まず、AFM クラウドがあります。これはサーバーのハードワークモデルです。基本的に、レイテンシーとサービスコストに関して最適化されています。
そして次は、AFM Cloud Image です。これは私たちの次世代の画像生成および編集モデルであり、ご存知のとおり、今朝聞いた空間リフレーミングなどを含む、数多くの驚くべき体験を可能にします。
そして、先ほど説明した 4 つのモデル、FM、Core、Core Advanced Cloud、Cloud Image はすべて、Apple Silicon 用のカスタム ビルドであり、独自のデータを使用してトレーニングされ、Gemini フロンティア モデルの外部を使用して洗練されています。
最後に、エージェント ツールの使用や複雑な推論など、最も要求の厳しいタスクのいくつかに対応するために、AFM Cloud Pro が用意されています。これは、Gemini フロンティア モデルと同様の品質を備えた最も高性能なモデルです。
そして、このモデルを実稼働環境に導入するために、Google と Nvidia の両方と協力してプライベート クラウド コンピューティング インフラストラクチャを拡張しています。

Apple の比類のないプライバシー保証を維持しながら、Google のクラウド内の NVIDIA GPU に構造を組み込むことはできますよね?
したがって、このモデルファミリー全体にわたって、私たちの目標は、すべてのユーザーリクエストを、最小のレイテンシーで最高の応答を提供するモデルに一致させることです。そのため、私たちは共に、この次世代モデルと、新しい Siri AI エクスペリエンスや OS 全体にわたるすべての驚くべきインテリジェント エクスペリエンスなど、その上に構築できる素晴らしい機能に非常に興奮しています。
ワイヤレス CarPlay をあらゆる車に導入
デビッド・ポーグ著「アップル：最初の50年」
AirPods Pro 3: 199 ドル (通常 199 ドル)
Beats 編み込み USB-C 充電ケーブル
Chance をフォローしてください: Threads、Bluesky、Instagram、Mastodon。
FTC: 当社は収入を得る自動アフィリエイト リンクを使用しています。もっと。
iOS 27 の CarPlay の新機能はすべてここにあります
Apple Watch を強化する watchOS 27 の主な機能
ここから公式 macOS 27 Golden Gate 壁紙をダウンロードしてください
iOS 27のApple Walletの新機能はすべてここにあります
その他の Apple ニュースについては、YouTube の 9to5Mac をご覧ください。
展開する
閉じる
コメント
展開する
閉じる
コメント
ガイド
iOS は Apple のモバイル オペレーティング システムです。
Chance は 9to5Mac の編集長で、サイト全体の運営を監督しています。また、9to5Mac Daily および 9to5Mac Happy Hour ポッドキャストのホストも務めています。
ヒント、質問、タイプミスは、chance@9to5mac.com に送信できます。

## Original Extract

Apple’s Siri team, led by Craig Federighi, held a post-WWDC keynote tech talk with members of the press this afternoon...

Skip to main content
Toggle main menu
9to5Mac Logo Go to the 9to5Mac home page
Switch site
9to5Toys
Toggle dark mode
Search for:
Submit
Toggle search form
Forums
Craig Federighi details Apple’s collaboration with Google for Siri AI in iOS 27
17 Comments
Apple’s Siri team, led by Craig Federighi, held a post-WWDC keynote tech talk with members of the press this afternoon to talk through iOS 27 and the new Siri AI.
During the talk, Federighi shared more details about Apple’s collaboration with Google.
Federighi was joined by Amar Subramanya (vice president of AI), Mike Rockwell (Siri lead), and Sebastien Marineau-Mes (software VP).
On the Google collaboration, Federighi explained:
Of course, we don’t have the Gemini app as our app. In fact, none of that client code is part of how we run on iOS. For these models, we use none of the models that Google deploys to their customers, nor do we use the infrastructure and means by which they deploy models to their customers. And then, when it comes to the knowledge base, we of course don’t use Google Search or anything like that as the foundation of our system. So I hope that’s clear. The amount of the Google Assistant we use is none.
So let’s talk about what we do use, or how our system is built.
It starts, of course, with our Assistant experience. And as you saw earlier today, this Assistant experience is deeply integrated into the system, into iOS, into iPadOS, into macOS. You saw on the iPhone how the assistant emerges, I think really beautifully, in Liquid Glass out of the Dynamic Island, how you can summon it from the side button or by speaking to Siri by name. But more than that, it’s integrated across all sorts of places in the system. So whether you’re writing with Writing Tools, clicking with the context menu, all of this is deeply integrated into the system experience.
Now, plugged into that is the Siri app. The Siri app is a great way to get back to a conversation that you previously started, to either look at what you’ve previously been doing, maybe extend that conversation, or start a new one. But this app isn’t just reaching out to some model in the cloud. It’s built on top of powerful system software in Apple Intelligence.
This includes the System Orchestrator, which is key to the privacy architecture of our entire system. It’s what coordinates requests against things like the App Toolbox that provides access to actions within your apps, the Spotlight Semantic Index to access personal content to help fulfill your request, and even things like on-screen context to understand what you might be looking at at the moment you’re making a request.
This, in turn, is built on a set of powerful on-device models. These handle everything from understanding speech to synthesizing the voice that speaks back to you, to understanding visually the environment and the on-screen context, understanding if there’s something relevant there, understanding text that might be on the screen, as well as a whole set of other models.
And for some requests, models are capable of processing your Siri requests entirely locally on the device. But sometimes the System Orchestrator realizes that it’s a more sophisticated question, and then it wants to draw on greater intelligence. It does that by contacting our models running in Private Cloud Compute.
The goal of Private Cloud Compute is to extend the same privacy promise of the iPhone up into the cloud, such that your requests are completely private to you. They’re never stored, they’re never accessible to anyone, including Apple, they’re only processed as part of the request, and nothing can access them. All of those properties are not only built architecturally deep into the system, but are also something that third-party researchers can continuously verify.
Now, in that deployment model, we have a family—our third generation—of Apple Foundation Models, from our AFM Cloud and AFM Cloud Pro models to our AFM Fusion model and image model. These are the models that are the product of our collaboration with Google, and you’ll hear more about that when we continue. But those are architected to run on our deployment architecture. These are models designed specifically for our Apple Intelligence experiences. They’re what powered everything you saw in the keynote presentation earlier today.
Finally, when you make a request involving current events or other elements of world knowledge, those responses are grounded by accessing Apple’s World Knowledge Service. This is something that we’ve built over many years and provides a great source of information to satisfy your request.
So this system is what supports the full range of Apple Intelligence experiences that you saw earlier. I hope that grounds our discussion as well. And what we’d like to do now is bring out members of the leadership team who helped work on all of this.
Liked by 6 people
Interesting comments, and does a nice job of fleshing out how exactly Apple is going about protecting user privacy when using Google models and servers. After months of speculation, it's nice to get the details.
We’re super excited about our 3rd generation of Apple Foundation Models, or AFM, in partnership with Google. We’ve built a family of models spanning on-device to the cloud. Now, before I walk through each of the models in the family, the headline, I would say, for this generation is every model is a significantly both in quality and capability compared to our previous generation.
Just walking through each of the models, we’d start with our on-device models. So 1st off, we have AFM Core. This is the next generation of our on-device model that we’re shipping on devices today. It follows the dense architecture.
And then next up, we have AFM Core Advanced.
This is unlike any on-device model we’ve run before. It uses a sparse architecture, it’s natively multimodal. And as a result of that, huge leap in the capabilities of this model, enabling some of the features you heard about this morning, like invitation and expressive voices, all working completely on device as a result of this model.
Moving on to our server models, all of which are served out of our Private Cloud Compute. First off, we have AFM Cloud. This is our server work hard model. It’s basically optimized for latency and serving cost.
And then next up, we have AFM Cloud Image. This is our next generation of image generation and editing model, and, you know, enables a number of amazing experiences, including things like spatial reframing, that you also heard about this morning.
And all of these four models that we just talked about, FM, Core, Core Advanced Cloud, and Cloud Image, all of these are custom builds for Apple Silicon, trained using proprietary data, and refined using outwards from Gemini frontier models.
Now, finally, for some of the most demanding tasks, like agentic tool use, complex reasoning, we have AFM Cloud Pro. This is our most capable model with quality similar to Gemini frontier models.
And to bring this model to production, we work with both Google and Nvidia to extend our Private Cloud Compute infrastructure to NVIDIA GPUs in Google’s cloud, while maintaining Apple’s unmatched privacy guarantees, right?
So, across all of this family of models, our goal is to match every user request to the model which provides the best response at the lowest latency. And so together we’re super excited about this next generation of models and the amazing features it enables us to build on top of them, including the new Siri AI experience and all of the amazing intelligent experiences across the OS.
Bring wireless CarPlay to any car
“Apple: The First 50 Years” by David Pogue
AirPods Pro 3: $199 (Reg. $199)
Beats Woven USB-C Charging Cables
Follow Chance : Threads , Bluesky , Instagram , and Mastodon .
FTC: We use income earning auto affiliate links. More.
Here’s everything new for CarPlay in iOS 27
Top watchOS 27 features that will enhance your Apple Watch
Download the official macOS 27 Golden Gate wallpapers here
Here’s everything new for Apple Wallet in iOS 27
Check out 9to5Mac on YouTube for more Apple news:
Expand
Close
comments
Expand
Close
comments
Guides
iOS is Apple's mobile operating system that runs…
Chance is the editor-in-chief of 9to5Mac, overseeing the entire site’s operations. He also hosts the 9to5Mac Daily and 9to5Mac Happy Hour podcasts.
You can send tips, questions, and typos to chance@9to5mac.com.
