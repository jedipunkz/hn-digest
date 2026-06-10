---
source: "https://shopify.engineering/quick"
hn_url: "https://news.ycombinator.com/item?id=48481517"
title: "Quick: An internal hosting platform for the AI era"
article_title: "Quick: An internal hosting platform for the AI era (2026) - Shopify"
author: "okneil"
captured_at: "2026-06-10T21:45:35Z"
capture_tool: "hn-digest"
hn_id: 48481517
score: 2
comments: 0
posted_at: "2026-06-10T19:38:22Z"
tags:
  - hacker-news
  - translated
---

# Quick: An internal hosting platform for the AI era

- HN: [48481517](https://news.ycombinator.com/item?id=48481517)
- Source: [shopify.engineering](https://shopify.engineering/quick)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T19:38:22Z

## Translation

タイトル: クイック: AI 時代の社内ホスティング プラットフォーム
記事タイトル: クイック: AI 時代の内部ホスティング プラットフォーム (2026) - Shopify
説明: Quick を使用すると、Shopify の誰でも数秒でサイトを出荷できます。それは私たちが構築し共有する方法の文化を変えました。

記事本文:
コンテンツにスキップ ソリューション 開始 ビジネスを開始します。ブランドを構築する
ウェブサイトを作成します。オンラインストア編集者
ストアをカスタマイズします。ストアのテーマ
ビジネス アプリを検索します。 Shopify アプリストア
自分のサイトのドメインを所有します。ドメインとホスティング
無料のビジネス ツールを探索してください。ビジネスを運営するためのツール
販売 製品を販売します。オンラインまたは対面で販売する
顧客をチェックしてください。ワールドクラスのチェックアウト
オンラインで販売します。オンラインでビジネスを成長させる
さまざまなチャネルで販売します。何百万人もの買い物客にリーチし、売上を伸ばす
世界中に販売します。海外販売
卸売＆直接販売。企業間 (B2B)
マーケット あなたのビジネスをマーケティングします。顧客にリーチして維持する
ソーシャル全体の市場。ソーシャルメディアの統合
顧客とのチャット。 Shopify 受信箱
顧客を育成する。 Shopify メッセージング
聴衆を知る。顧客の洞察を得る
管理 ビジネスを管理します。売上、注文、分析を追跡する
パフォーマンスを測定します。分析とレポート作成
在庫と注文を管理します。在庫と注文の管理
ビジネスを自動化します。 Shopify フロー
Shopify 開発者。 Shopify の強力な API を使用して構築する
プラス。成長するデジタル ブランドのためのコマース ソリューション
すべての製品。 Shopify のすべての製品と機能を詳しく見る
リソース ヘルプとサポート ヘルプとサポート 。 24時間年中無休のサポートを受ける
ビジネスコース。実績のある専門家から学ぶ
人気のトピック Shopify とは何ですか? 。当社のコマースプラットフォームの仕組み
必須ツール ビジネス名ジェネレーター。
新機能変更ログ。最近の更新情報のソース
ニュースルーム 。すべての企業ニュースとプレスリリース
ブログ | Culture Quick: AI 時代の社内ホスティング プラットフォーム
Quick を使用すると、Shopify の誰でも数秒でサイトを出荷できます。それは私たちが構築し共有する方法の文化を変えました。
Shopify には、いじくり回すのが好きな人々が集まります。ここでは建築がボトルネックになったことは一度もありません。人は常に何かを作っています: p

ロトタイプ、ダッシュボード、チーム用の小さなツール。大変だったのは、それらを他の人の前で披露することでした。
これに対する私たちの解決策は Quick でした。これは、HTML とアセットのフォルダーをドロップして、Shopify 従業員のみが閲覧できる安全な URL を取得する内部プラットフォームです。フレームワーク、パイプラインのデプロイ、構成ファイルは必要ありません。フォルダーをアップロードするだけでサイトが稼働します。データベース、AI、ファイル ストレージ、WebSocket が必要な場合は、API を呼び出すだけで利用できます。
私たちは 2025 年 7 月に Quick を立ち上げましたが、そのタイミングは完璧でした。 AI は、私たちエンジニアだけでなく、あらゆる分野の人々がプロンプトから機能する Web サイトを生成できるまでに十分に進歩したばかりでした。クイックは彼らにそれを置く場所を与えました。 AI がそれを構築した理由ではありませんが、AI が普及した理由の大きな部分を占めています。
現在、Quick は Shopify 全体で 50,000 を超えるサイトをホストしています。企業の半数以上が少なくとも 1 つを作成しています。チームが毎日依存しているダッシュボードから、誰も求めていないのに誰もがプレイするマルチプレイヤー登山ゲームまで、あらゆるものがあります。
HTML ファイルをどこかに配置して提供する最も簡単な方法を見つけるという考えから、すぐに始めました。そして、すべての「サイト」が Google Cloud Storage バケット内の単なるアセットのフォルダーであることよりも単純なことはあるでしょうか?
これらにサービスを提供するために、mysite.quick.shopify.io が mysite フォルダーに直接マップされるように、ワイルドカード構成を備えた軽量の NGINX サーバーを前面に配置しました。しかし、NGINX にはバケットのクエリについて何も知られたくありませんでした。ここで gcsfuse が登場しました。gcsfuse はバケットをローカル ファイルシステムの一部であるかのようにマウントするため、NGINX はローカル ファイルを提供しているだけであると認識します。
認証も同様に簡単です。サーバー全体が Identity-Aware Proxy (IAP) の背後に配置されているため、すべてのリクエストはサイトに到達する前にすでに認証された Shopify 従業員になります。
「クイック デプロイ」コマンドは何もありません

gcloud の rsync の小さなラッパー以上のものです。ローカル ディレクトリからファイルを取得し、バケットに直接プッシュします。古き良き時代の FTP のような感じです。
この構成では、Quick はすでに多くのユースケースに対応できます。しかし、プロトタイプやサイトにバックエンドの機能が少しだけ必要な場合はどうなるでしょうか?おそらくそれは、少量のデータを保存するため、ファイルをアップロードするため、または AI を呼び出すためです。突然、データベースとインフラストラクチャを起動しなければならなくなりましたが、多くの場合、それはやりすぎのように感じられます。
そのとき私たちは、「Quick 上のすべてのサイトがアクセスでき、基本的なバックエンド サービスを提供できる単一のサーバーがあればどうなるだろうか?」と考えました。クライアント側に優れた API を追加すれば、準備は完了です。
Quick サイトにブログ投稿を保存したいですか?
const post = Quick.db.collection('posts');
const created = awaitposts.create({
タイトル: 'Hello Quick DB',
ステータス: 'ドラフト'、
created_at: new Date().toISOString()
});
何か変更があったときにリアルタイムで更新する必要がありますか?
const unsubscribe =posts.subscribe({
onCreate: (doc) => console.log('New:', doc),
onUpdate: (doc) => console.log('Updated:', doc),
onDelete: (id) => console.log('削除済み:', id),
});
オリジナルの Firebase は私たちにとって大きなインスピレーションでした。これにより、インターネットに公開され、読み取り/書き込みだけが可能なキーバリュー ストアを簡単に起動できます。スキーマや移行などは必要ありませんでした。そして何よりも、そこに保存したものはすべて、接続されているすべてのクライアント間で魔法のように同期されます。
私たちは当初、各 Quick サイトに独自の sqlite データベースを持たせるというアイデアを試してみましたが、gcsfuse では期待したほどうまく機能しませんでした。単一の CloudSQL データベースをその前に nodejs サーバーを置いて起動する方がはるかに簡単でした。
次に追加した機能は AI で、どの Quick サイトもクライアント側で実行できるようになりました。

API キーを提供する必要なく、選択した LLM を呼び出すことができます。 API キーはサーバーに保存され、Shopify AI プロキシに渡されます。
// ブラウザから直接 LLM 呼び出しを行う
const res = await Quick.ai.chat([{ role: 'user', content: 'タスクの要約' }]);
// 画像生成やその他のフロンティア モデルも呼び出すことができます
それから私たちは、サイトに必要なもののリストを検討し続けました。ファイルのアップロード、サイトが Big Query からデータを取り込めるようにするデータ ウェアハウスのサポート、共同アプリの構築を可能にする Websocket のサポートを追加しました。
サイトは IAP の背後にあるため、あらゆるリクエストに必要なすべてのユーザー情報があり、クライアントに提供できます。そうすれば、サイトは誰がそれを使用しているかを即座に知ることができ、それは名前、役職、チーム、Slack ハンドルなどを含む優れた小さな Identity API の一部になります。
最後までに、次のコア機能グループにたどり着きました。
これらのいくつかの構成要素だけで、インターネット全体を再現できるように感じられるのは驚くべきことです
これらのサービスのいずれかを認証なしで公共のインターネットに公開すると、運用チームは睡眠不足になります。ただし、Quick サイトは Shopify の信頼できるウォール内でのみアクセスできるため、構成不要のクライアント側 API を提供できるという贅沢が得られます。すべてのキーはサーバーに保存されます。
すべての Quick API に関するドキュメントはありますが、実際に読んだ人がいるかどうかはわかりません。それは、Quick には、エージェントが使用するために必要なすべてのスキルが最初から備わっているからです。 `quick init` と入力してお気に入りのエージェントを起動するだけで、レースに出発できます。
「チームがリアルタイムでランチスポットに投票できるサイトを作ってください。」
「全員参加中にライブ投票を実施できるサイトを作ってください。」
1 分以内にサイトを立ち上げて稼働させることができます

全員が共有して使用できるようになります。
「サイト上の全員のマルチプレイヤー カーソルを表示できますか?」
WebSocket API のおかげで、それが可能になります。
Quick は 2025 年 7 月に社内でリリースされました。すでに HTML アーティファクトのバイブコーディングを行っていた私たちは、最終的にそれらをホストするための安全な場所を確保しました。アプリのデプロイメントの複雑さに苦労している人には、簡単な道が与えられました。そして他の人たちにとって、何が重大なことなのかはまだわかっていませんでした。
最初はジオシティーズと同じように感じました。人々は個人のホームページを作成していました。
純粋なノスタルジーから、誰かがウェブリングや、ゲストブックを備えた誕生日サイトまで作成しました。
最近では、誰かがオープンなウェブ上にオープンなゲストブックやコメント フィールドを設置すると、必ずハッキングされるか、スパムで埋め尽くされます。しかし、Shopify の信頼バブルでは、それは心配ありません。私たちは 2000 年代初頭のウェブの喜びを、何の欠点もなく追体験していました。大きな違いは、AI のおかげで誰でも HTML を作成できるようになったということです。
そして 2025 年 12 月に事態は本格的に動き出しました。
現在、プロトタイプ、ダッシュボード、開発ツール、プレゼンテーションに至るまで、あらゆるものに使用されています。
デザイナーはチームが使用するカスタム ツールを作成するだけです。 Artifact と同様、作業を共有するための内部ツールです。
それに手を伸ばすのが習慣になっています。管理画面でテーマがどのように表示されるかについてのアイデアを共有したいですか?画像を共有するのではなく、クイック サイトを共有してください。
最新の Remix ランディング ページのインタラクティブな背景をカスタマイズするツールが必要ですか?クイックサイトを構築します。
Google Meet は昨年小規模な障害が発生しましたが、15 分以内に誰かが通信に WebRTC を使用する代わりのデバイスをバイブコーディングしました。
それは私たちが構築し共有する方法の文化を完全に変えました。そして、私たちは主に「深刻な」ユースケースにそれを使用していますが、根底には人々を奇妙なものをいじくり回したくなるようなGeocitiesの雰囲気がまだ残っています。

d 楽しいこと。
特に WebSocket API を使用するとマルチプレイヤー機能を簡単に追加できるため、多くのゲームが作成されています。最近のゲームジャムでは 140 以上のゲームが提出されました。
インターネット上の公開ゲームにリーダーボードを追加したい場合、最初に考えられるのは、ハッカーを防ぐ方法です。もう一度言いますが、Quick が社内に存在することで、それらすべてが取り除かれます。 「リーダーボードを追加すべきか」は、「たぶん」から「はい！」までかかります。それはクリエイティブな楽園です。
Quick を構築するときに私たちが予想していなかった点の 1 つは、ユーザーがどのようにして他のサイト内にサイトを埋め込み始めるかということです。これは単なる Web であるため、あるサイトが別のサイトからコードを直接インポートできます。そこで私たちは、人々が共有 JS ライブラリを公開し始め、さらにそのライブラリ用のランディング ページを作成し始めていることに気づきました。
Quick は独自の内部エコシステムへと成長しています。 Figma スタイルのコメントをサイトに追加したり、音声、分析、実績などを追加したりするためのライブラリが見つかります。
Quick の哲学: シンプルに保ち、制約を受け入れる
アクセス許可をどのように実装したか、あるいは人々がサイトをどのように管理しているか疑問に思われるかもしれません。それは何もないことがわかりました。すべての Quick サイトはすべての従業員に公開されています。 「サイト所有者」という概念すらありません。サイトを更新したいですか?新しいファイルで上書きします。サブドメインを引き継ぎたいですか?上書きしてください！
何かが内部ツールになると、オープン Web の複雑さがまったくなくなるのは驚くべきことです。
毎日、新しい機能のリクエストが届きます。カスタム バックエンドを入手できますか? cronジョブを追加できますか?私たちはノーと言うのが本当に上手になりましたが、新機能を数分でバイブコーディングできる AI 環境では特に難しいことです。しかし、重要なのは制約です。
小規模で固定された機能セットにより、Quick の使用と保守が簡単になります。

そしてそれが人々の創造性を低下させるのではなく、より創造的にするのです。誰かが機能リクエスト X を持って私たちに来たとき、多くの場合、私たちはそれがすでに可能であることを彼らに示すことができます。彼らは違うアプローチをすればいいだけで、既存の作品で何ができるかに少し驚きます。
現在、50,000 を超える Quick サイトが作成されています。 Shopify 従業員の 50% 以上が少なくとも 1 つのサイトを作成したことがあります。これらはすべて、月額 200 ドルの実行料金がかかる単一の VM 上で実行されます。
その大部分はクライアント側であるため、サーバーはアセットの提供と API リクエストの処理のみを担当します。規模に関しては、従業員の数を把握しているため、何らかの形で 1,000 倍以上の人々が突然プラットフォームを使い始めるという危険はありません。
だからといって、途中で問題がなかったわけではありません。ループ上でバッチ処理ジョブを実行し、大量のデータをデータベースに保存しようとしている人を見つけることがあります。そのため、私たちはいくつかのレート制限を実装することになりました。
また、私たちは時間の経過とともにノードから Go を使用するように移行してきました。これはメモリ管理と並列処理に非常に役立ちます。
私たちが Quick を始めたのは、Shopify で何かを共有することは、構築することよりも難しいからです。修正は、ほとんど滑稽なほど単純であることが判明しました。ファイルのフォルダー、URL、そしてすべてが私であることに伴う信頼です。

[切り捨てられた]

## Original Extract

Quick lets anyone at Shopify ship a site in seconds. It has changed the culture of how we build and share.

Skip to Content Solutions Start Start your business . Build your brand
Create your website . Online store editor
Customize your store . Store themes
Find business apps . Shopify app store
Own your site domain . Domains & hosting
Explore free business tools . Tools to run your business
Sell Sell your products . Sell online or in person
Check out customers . World-class checkout
Sell online . Grow your business online
Sell across channels . Reach millions of shoppers and boost sales
Sell globally . International sales
Sell wholesale & direct . Business-to-business (B2B)
Market Market your business . Reach & retain customers
Market across social . Social media integrations
Chat with customers . Shopify Inbox
Nurture customers . Shopify Messaging
Know your audience . Gain customer insights
Manage Manage your business . Track sales, orders & analytics
Measure your performance . Analytics and Reporting
Manage your stock & orders . Inventory & order management
Automate your business . Shopify Flow
Shopify Developers . Build with Shopify's powerful APIs
Plus . A commerce solution for growing digital brands
All Products . Explore all Shopify products & features
Resources Help and support Help and support . Get 24/7 support
Business courses . Learn from proven experts
Popular topics What is Shopify? . How our commerce platform works
Essential tools Business name generator .
What’s new Changelog . Your source for recent updates
Newsroom . All company news and press releases
blog | Culture Quick: An internal hosting platform for the AI era
Quick lets anyone at Shopify ship a site in seconds. It has changed the culture of how we build and share.
Shopify attracts people who love to tinker. Building has never been the bottleneck here. People are always making things: prototypes, dashboards, little tools for their teams. The hard part was getting those things in front of anyone else.
Our solution to this was Quick: an internal platform where you drop in a folder of HTML and assets and get back a secure URL that only Shopify employees can see. No frameworks, deploy pipelines, or config files. You just upload a folder and your site is live. If you need a database, AI, file storage, or websockets, those are just an API call away.
We launched Quick in July 2025, and the timing turned out to be perfect. AI had just gotten good enough that people across every discipline, not just us engineers, could generate a working website from a prompt. Quick gave them somewhere to put it. AI wasn't why we built it, but it's a big part of why it took off.
Today, Quick hosts more than 50,000 sites across Shopify. Over half the company has created at least one. Everything from dashboards that teams rely on daily to a multiplayer mountain-climbing game that nobody asked for but everybody plays.
Quick started with the idea of finding the simplest way to put HTML files somewhere and serve them. And what's simpler than every "site" being just a folder of assets in a Google Cloud Storage bucket?
To serve them, we put a lightweight NGINX server in front, with a wildcard config so that mysite.quick.shopify.io maps straight to the mysite folder. But we didn't want NGINX to know anything about querying buckets. This is where gcsfuse came in: it mounts the bucket as if it were part of the local filesystem, so NGINX thinks it's just serving local files.
Authentication is just as simple. The whole server sits behind Identity-Aware Proxy (IAP), so every request is already a verified Shopify employee before it ever reaches a site.
The `quick deploy` command is nothing more than a small wrapper around gcloud’s rsync. It grabs the files from your local directory and pushes it up directly to the bucket. Feels like the good old days of FTP.
In this configuration, Quick could already accommodate many use cases. But what if your prototype or site needed just a little bit of functionality from a backend? Maybe it’s to store a bit of data, or upload a file, or call out to AI. Now suddenly you’re left having to spin up databases and infrastructure, which feels like overkill in many cases.
That’s when we thought to ourselves, “what if we had a single server that all sites on Quick could access and could provide basic backend services?” Give it a nice little client-side API and you’re all set.
Want your Quick site to save a blog post?
const posts = quick.db.collection('posts');
const created = await posts.create({
title: 'Hello Quick DB',
status: 'draft',
created_at: new Date().toISOString()
});
Need realtime updates when something changes?
const unsubscribe = posts.subscribe({
onCreate: (doc) => console.log('New:', doc),
onUpdate: (doc) => console.log('Updated:', doc),
onDelete: (id) => console.log('Deleted:', id),
});
The original Firebase was a big inspiration for us. It let you easily spin up a key value store exposed to the internet that you could just read / write to. You didn’t need schemas, or migrations, or anything. And best of all, anything you saved in it magically synced across all connected clients.
We initially played around with the idea of each Quick site having its own sqlite database, but that didn't play as well as we hoped with gcsfuse. It was much easier to spin up a single CloudSQL database with a nodejs server in front of it.
AI was the next feature we added, so that any Quick site could make client-side calls to the LLM of their choosing without needing to provide any API keys. The API keys are stored on the server and passed along to our Shopify AI proxy.
// Make LLM calls right from your browser
const res = await quick.ai.chat([{ role: 'user', content: 'Summarize my tasks' }]);
// Can also call image gen and any other frontier models
Then we kept going through the list of things that sites might need. We added file uploads, data warehouse support so sites could pull in data from Big Query, and Websocket support to make it possible to build collaborative apps.
Since the sites are behind IAP, we have all the user information needed for any request, which we can provide to the client. That way sites can instantly know who is using it, and that becomes part of a nice little Identity API with things like name, title, team, Slack handle, etc.
By the end of it we landed on this core group of features:
It's amazing how just with these few building blocks it feels like we can recreate the entire Internet
Exposing any of these services to the public internet without authentication would make an ops team lose sleep. But since Quick sites are only accessible within the trusted walls of Shopify, we have the luxury of being able to provide a zero-config client side API. All keys are stored on the server.
While there are docs for all the Quick APIs, I’m not sure if anyone has ever actually read them. That’s because Quick comes out of the box with all the skills your agent needs to use them. All you have to do is type `quick init` , launch your favourite agent, and you’re off to the races.
“Make me a site where my team can vote on lunch spots in real time.”
"Make me a site where we can run a live poll during the all-hands."
In less than a minute you have a site up and running, ready to be shared and used by all.
"Can I see multiplayer cursors for everyone on the site?”
Thanks to the websocket API, you sure can!
Quick launched internally in July 2025. Those of us who were already vibe-coding HTML artifacts finally had somewhere secure to host them. Those struggling with app deployment complexity now had an easy path. And for everyone else, they didn't see what the big deal was….yet.
For the first bit, it felt just like Geocities. People were making personal homepages.
Out of pure nostalgia, someone made a webring and even a birthday site with a guestbook.
These days, if someone puts an open guestbook or comment field on the open web, it’s sure to get hacked or filled with spam. But in the Shopify trust bubble, that’s not a concern. We were reliving the joys of the early 2000’s web without any of the downsides. The big difference is that anyone could now author HTML thanks to AI.
Then in December 2025 things really popped off.
It’s now being used for everything from prototypes, dashboards, dev tools, and presentations.
Designers will just create custom tools for their teams to use. Like Artifact, an internal tool for sharing work.
It’s become second nature to reach for it. Want to share an idea for how themes are shown in the admin? Don’t share an image, share a Quick site.
Need a tool for customizing the interactive background on the latest Remix landing page? Build a Quick site.
Google Meet had a small outage last year, and within 15 minutes someone vibecoded a replacement that used WebRTC for communication.
It’s completely changed the culture of how we build and share. And while we are using it mostly for “serious” use cases, there’s still that underlying Geocities feel that invites people to tinker with weird and delightful things .
There’s a good amount of games that people are building, especially since the websocket API makes it so easy to add multiplayer functionality. Over 140 games were submitted at a recent game jam.
If you wanted to add a leaderboard to a public game on the internet, your first thought would be how to prevent hackers. Once again, Quick being internal removes all that. It takes “should we add a leaderboard” from a “maybe” to a “hell yes!” It’s a creative paradise.
One thing we didn’t anticipate when building Quick is how people would start embedding sites within other sites. Because it's just the web, one site can import code straight from another. So we noticed people started publishing shared JS libraries, and even making landing pages for them.
Quick is growing into its own internal ecosystem. You can find libraries for adding Figma-style comments to your site, for adding voice, analytics, achievements, and so much more.
The Quick philosophy: Keep it simple + embrace the constraints
You might be wondering how we implemented permissions, or how people manage their sites. Turns out there’s none of that. All Quick sites are open to all employees. There’s not even a concept of a “site owner.” Want to update your site? Overwrite it with new files. Want to take over a subdomain? Overwrite it!
It’s amazing how the complexities of the open web just disappear when something is an internal tool.
Every day we get new feature requests. Can we get custom backends? Can you add cron jobs? We've gotten really good at saying no, which is especially hard in an AI climate where you can vibecode a new feature in minutes. But the constraints are the whole point.
A small, fixed set of capabilities is what keeps Quick simple to use and maintain, and it's what makes people more creative, not less. When someone comes to us with feature request X, more often than not we can show them it's already possible. They just have to approach it differently, and they leave a little surprised at what the existing pieces can do.
As of now over 50,000 Quick sites have been created. More than 50% of Shopify employees have created at least one site. And all of this is running on a single VM that costs $200 a month to run.
Since so much of it is client side, the server is only in charge of serving assets and processing API requests. As for scale, we know how many employees we have so there’s never a danger of somehow 1,000x more people starting to use the platform all of a sudden.
That’s not to say there haven’t been hiccups along the way. We’ll sometimes find someone trying to do a batch process job on a loop and store massive amounts of data in the database. That’s led us to implement some rate limiting practices.
We’ve also over time migrated away from node to using Go, which helps a lot with memory management and parallelism.
We started Quick because sharing things at Shopify was harder than building them. The fix turned out to be almost comically simple: a folder of files, a URL, and the trust that comes with everything being i

[truncated]
