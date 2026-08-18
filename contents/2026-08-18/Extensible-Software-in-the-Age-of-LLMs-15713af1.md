---
source: "https://jeremymorrell.dev/blog/extensible-software-in-the-age-of-llms/"
hn_url: "https://news.ycombinator.com/item?id=49345304"
title: "Extensible Software in the Age of LLMs"
article_title: "Extensible Software in the age of LLMs | Jeremy Morrell"
image: "https://jeremymorrell.dev/_astro/omg-meme-social.BBc_Z2VQ.png"
author: "masterj"
captured_at: "2026-08-18T13:35:41Z"
capture_tool: "hn-digest"
hn_id: 49345304
score: 1
comments: 0
posted_at: "2026-08-18T13:25:03Z"
tags:
  - hacker-news
  - translated
---

# Extensible Software in the Age of LLMs

- HN: [49345304](https://news.ycombinator.com/item?id=49345304)
- Source: [jeremymorrell.dev](https://jeremymorrell.dev/blog/extensible-software-in-the-age-of-llms/)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T13:25:03Z

## Translation

タイトル: LLM 時代の拡張可能なソフトウェア
記事のタイトル: LLM 時代の拡張可能なソフトウェア |ジェレミー・モレル
説明: ソリッド コア + 機能ベースのサンドボックス + LLM = スーパーパワーを持つユーザー

記事本文:
--> LLM 時代の拡張可能なソフトウェア | Jeremy Morrell Jeremy Morrell ブログの RSS ブログに関するメモ
2026 年 8 月 18 日 LLM 時代の拡張可能なソフトウェア
私たちが現在やり取りしている Web ソフトウェアのほとんどは静的です。開発者は時間と注意力が限られているため、最大のユーザー グループにサービスを提供する機能の構築に集中します。需要曲線の頂点は既存のソフトウェアによって十分に対応されていますが、ユーザーごとに異なる満たされていないニーズのロングテールが存在します。
たとえ開発者があらゆる機能を盛り込むことに信じられないほどの意欲を持っていたとしても、ユーザー インターフェイスは使い物にならなくなる前に複雑になりすぎる可能性があります。機能が追加されるたびに、他のすべてのユーザーにとって製品が複雑になります。その機能の市場が小さい場合、それを必要としないすべてのユーザーにとって製品の品質が低下する可能性があります。
このような背景から、LLM 支援コーディングの台頭は、このロングテールに該当するものを必要とするすべての人にとって真に力を与えてくれました。
ソフトウェアはすっかり…ふにゃふにゃになってしまった #
LLM は、Software for One の構築において非常に優れていることがすぐにわかります。個人用アプリは、エンタープライズ ソフトウェアの複雑さと説明責任をすべて回避し、1 人のワークフローに合わせてカスタマイズできます。
Y Combinator の Pete Koomen 氏は、彼らが小規模ソフトウェアと呼んでいるものにチャンスがあると考えています。彼らは何かに気づいていると思います。
エージェントを使用すると、自分自身またはチーム用の個人用ツールを簡単に構築できます。しかし、そのソフトウェアの展開、保護、共有は、ソフトウェアを作成することよりもはるかに複雑です。
小規模なソフトウェア用に構築されたクラウドは、その複雑さを取り除き、Google ドキュメントと同じくらい簡単に特注ツールを同僚と共有できるようにする可能性があります。
Pi は、私が LLM ネイティブ ソフトウェアと考え始めているものの良い例です。

確立されたコアですが、要求するだけでほぼ無限に拡張でき、ユーザーは自分のカスタマイズを他の人と共有できます。過去 1 年で、ユーザーは突然、コードを読み上げて存在させる能力を獲得しました。既存のソフトウェアのほとんどはこれを利用できません。パイはそれに身を乗り出します。
この自己拡張パターンに従うソフトウェアがさらに増え始めるのではないかと思います。ただし、プラグイン可能なソフトウェアの既存の例のほとんどは、AI エージェント、開発者 IDE、ビデオ ゲーム用の MOD、Blender アドオン、CAD 拡張機能など、ローカル ソフトウェアです。これらは、参入障壁が高い専門的なツールである傾向があります。
Web は、世界で最も成功したソフトウェア配布システムです。それは残すべきではありません。
私の仮説は、Web 上に拡張可能なソフトウェアの新たなチャンスがあるということです。 LLM は拡張機能の作成コストを大幅に削減し、最新のサンドボックス プリミティブは展開コストを削減し、適切なセキュリティ境界を提供します。私たちはアプリを堅牢で責任のあるコアとして構築することができ、LLM に不足している部分を補わせることで、ユーザーがアプリをさまざまな方向に安全に拡張できるようにします。ユーザーにスーパーパワーを与えることができます。
開示: 私は現在 Cloudflare で働いており、そこで Kenton Varda の著作に頻繁に触れることで、私の考え方の多くが形作られました。最後近くで、ダイナミック ワーカーがこのモデルに特に適していることを主張しますが、最初にいくつかの代替案について説明します。
現在、多くの Web システムは Webhook に依存して、ユーザーがアプリの変更に反応できるようにしています。これは機能しますが、完全に別個のサービスを構築して運用し、さらに配信上の問題が発生した場合に対処するという、拡張のハードルが非常に高くなります。
レコードの更新にフックして、独自のロジックをスライドできるようにしたいと考えています。 「このタグをレコードに添付すると、関数が実行されます

「毎日の cron でこのアクションを実行してください」。
実のところ、そんなことはまったく考えたくないんです。後で読むアプリに次のように伝えたいです。
4000 ワードを超えるお気に入りの記事はすべて、<選択した電子書籍リーダー> に送信してください。
arxiv の <my専門> で毎週公開される新しい論文を探し、それが私の研究にどのように関連しているかについての独自の要約を先頭に追加し、<tag> でタグ付けします。
デフォルトのアルゴリズムでは、 <site I read頻繁に> が完全に文字化けします。いくつかの例を取り出して、それ用のカスタム パーサーを作成します。
そして、ロボットがコードの愚かな部分を押し出し、それらをいくつかの拡張ポイントにフックして、それを実現します。また、同じ機能を必要とする他の人とも自分が作ったものを共有できるはずです。 1
LLM ネイティブの拡張アプローチを見てみたい領域がさらにいくつかあります。
さて、これは明らかです。 pi 、 deepseek 、 opencode はすべてこの分野で実験を行っています。
Pi は、コアに新しいアイデアをすべて追加するのではなく、ツール、コマンド、イベント、UI に安定したフックを提供するため、リクエストを小さな TypeScript 拡張機能に変換し、適切な場所にリロードできます。これらの拡張機能は共有可能なパッケージにバンドルでき、ハーネス自体を肥大化させることなくエコシステムがアイデアのロングテールを吸収できるようになります。
しかし、これらの視聴者は、少なくとも現在存在する限り、かなり少数です。ローカル マシン上でカスタム ソフトウェアを快適に実行できる必要があります。企業環境では、組織は、誰も見たことのない、あるいは今後も見ることのないソフトウェアを実行することに慣れていなければなりません。自分で Pi をサンドボックス化しない限り、Pi 拡張機能は Pi 自体と同じ権限で実行されます。
ソフトウェア エンジニアはその方法を見つけるでしょうが、会計士、医師、弁護士、その他の何千もの職業にも、より優れたツールが必要です。彼らには安全なエージェントが必要です

自社のドメインや独自のワークフローに合わせて、確実かつ簡単にカスタマイズできます。
エージェントを利用する人を増やしたいとしても、それは彼らをソフトウェア開発者にするという意味ではありません。それは、ソフトウェアを彼らのニーズに合わせて作ることを意味します。
すべての企業は最終的に大量のデータを保有することになります。従業員は、データを表示し、クエリを実行し、調査し、他のシステム内の他のデータと関連付け、 <問題 x> に直面している顧客を見つけ、解約しそうな顧客を見つけ、その他にも百万もの作業を行う必要があります。
多くの企業が、AI 愛好家の従業員が独自のツールをバイブコーディングできるようにする実験を行っており、おそらくそれを PaaS にデプロイできます。これは方向的には正しいですが、下流で多くの問題を引き起こします。こうしたアプリを何百、何千も持ったら、どうやって維持するのでしょうか?彼らはどうやって必要なデータにアクセスするのでしょうか?必要なデータのみにアクセスするにはどうすればよいでしょうか?このソフトウェアの動作を監査するにはどうすればよいでしょうか?アクセス トークンに依存している場合、そのスコープは何でしょうか?誰がそれらを回転させますか?顧客情報がサードパーティにログアウトされていないことを確認するにはどうすればよいですか? GDPR に違反していないことを確認するにはどうすればよいでしょうか?
あるいは、実際の企業が自ら心配する必要があるその他のコンプライアンスやセキュリティに関する事項も数多くあります。
漏洩する可能性のある認証トークンがない場所にコードをデプロイする場所を与えたらどうなるでしょうか?データ アクセスは、すべてのコンプライアンス ボックスがチェックされていることを確認できる社内プラットフォーム チームによって処理されるのはどこですか?独自の自動化やカスタム ビューを安全に構築するためのスペースをユーザーに提供します。 2
スポイラー: これは基本的に Cloudflare OS です。
私はキャリアの多くを、扱いにくいサポート チケットの処理に費やしてきました。必然的に、ダッシュボードを調べたり、ログを検索したり、何百万もの異なる場所からデータを取得したりすることになります。ティックを開いたユーザーのデータを表示する拡張機能を作成しましょう

私の特定のシステムからサポート インターフェイスにアクセスします。私が調査する前に、最初の調査を行うためにエージェントを解雇できるように、フックを提供してください。 「特定のクォータ X をリセットする」など、実行する必要がある一般的なタスクがある場合は、それを実行できるボタンをビューに追加します。
それから、これらを私のチームにも共有して、みんなで助け合えるようにしましょう。
多くの可観測性ツールは同じ機能セット、つまり上部の小さな棒グラフでログを検索する方法に集中しています。個々のトレースを表示するためのトレース ウォーターフォール ビュー。カスタマイズ可能なメトリクス ダッシュボード。たぶんサービスマップ。いくつかの企業は、特にエージェントの台頭を受けて、新しい視覚化を実験しています。
由緒あるトレース ウォーターフォール図は、主に待ち時間と成功率を重視する、リクエスト/レスポンスとして形成されたシステムに非常に役立ちます。私たちの多くは、もう少し…ステートフル…またはダイナミックなシステムを使用していることに気づきました。最新のアプリでは、非決定的なエージェントや耐久性のあるワークフロー エンジンが実行されており、1 つのアクションに数時間または数日かかる場合があります。トレース スパンは信頼性を高めるための優れた情報源ですが、独自のビジュアライゼーション (または他の人のビジュアライゼーションをインストール) を試してみましょう。 3
見た目の美しさを超えて、私自身のロジックを注入してみましょう。
取り込み時のデータの任意の変換
アラームによって独自のスクリプトを開始させる: 決定論的なコードまたは独自のエージェント
最もリスクが高いときに独自のコードを実行するためのオプションを提供します: デプロイまたは機能フラグのロールアウト
ログに特別な MyResourceID がある場合は、それを別のプラットフォーム上のそのリソースに直接アクセスするリンクに変換します。
すべてのソフトウェアはおそらく次のようになっているはずです
opencode2 で行ったアーキテクチャ上の変更は、ほぼすべてが内部プラグインであることです。
そのうち 68 個が組み込みエージェントをカバーしています。

統合、設定の読み込みなど
これは、あらゆる動作を無効にできることを意味し、プラグイン API も適切にドッグフードします。
Web 上の拡張可能なソフトウェアは…難しい #
すべて簡単に聞こえるようにしました。それはそのようなものではありません。
私は、毎日使用するツールとしても、ソフトウェアとしても、Obsidian の大ファンです。
基本的なマークダウン エディタのように見えますが、数回クリックするだけで、カンバン ボードでタスクを追跡したり、メモをデータベースに変換したりするなど、ほぼ何でもできるように拡張できます。セマンティック検索のためにすべてのメモをベクトル データベースに保存したいですか?頑張れ ！さらに詳しく言えば、基礎となる Web UI プリミティブは簡単にハッキング可能です。
ただし、その能力には代償が伴います。Obsidian の拡張モデルでは、インストールするすべてのプラグインを信頼する必要があります。プラグインは基本的に何でもできます。 Obsidian は、自動化および手動レビュー、およびプラグイン作成者を検証することで、このセキュリティ課題と闘います。
メモ アプリの場合、これはおそらく適切なトレードオフです。リスクが低く、コミュニティが比較的小さいため、これが機能します。しかし、顧客記録、金融取引、プライベートメッセージなど、他人のデータを保持するソフトウェアに同じレベルの拡張性を求めた瞬間、このモデルは崩壊します。拡張性と Web サービスは常に課題です。
任意のコードを実行すると、セキュリティと悪用の問題が山積します。不完全なリスト:
ユーザーのコード内のエラーや無限ループによってサービスが停止されることはあってはならない
キーにアクセスできるため、顧客内線番号はキーをサードパーティに転送できます
同様に、機密データを公開する場合は、データが持ち出されないようにしてください
このシステムがサービス拒否攻撃に悪用されないようにする
ユーザーが誤ってサービス拒否を実行できないようにしてください。
Spectre 攻撃から保護する
人々が fr を使用できるようになれば

ええ、あなたの10セントで暗号通貨をマイニングするために計算します、彼らはそうします
しかし、確かに誰かがこれをやったでしょうか？ #
これを実現不可能だと断じる前に、Web 上のこの種の拡張性が巨大な規模で機能している明確な例があります。Salesforce です。
そう、それがセールスフォースです。彼らは 2007 年からそれを行っています。 (参考までに、AWS S3 と EC2 は 2006 年に開始されました。)
ほとんどの技術者に Salesforce とは何かと尋ねると、真っ白な顔をされるか、「CRM ではないのですか?」というような反応が返ってくるでしょう。ただし、Salesforce を大規模なマルチテナントのプログラム可能なプラットフォームと表現する方が正確です。初期のクラウド時代では、これは非常に困難でした。コンテナはなく、人々はこの奇妙な Java のようなカスタム言語 Apex を書くことを強制されました。
しかし、サーバーレスの台頭により、プラットフォームはより親しみのあるものに見え始めています。いくつかの例を考えてみましょう。
カスタム エンドポイントを公開する必要がある場合は、数行のコードで公開できます。プラットフォームは、ルーティング、認証、テナントの分離、実行を処理します。導入する Web サーバーはありません。目を細めて見ると、これが現代のサーバーレスの先駆けであることがわかります。
@RestResource (urlMapping = '/customer-health' )
共有クラス CustomerHealthApi を持つグローバル {
@HttpGet
グローバル静的アカウント getCustomer () {
文字列アカウントID =
RestContext.request.params。

[切り捨てられた]

## Original Extract

Solid core + capability-based sandboxes + LLMs = Users with superpowers

--> Extensible Software in the age of LLMs | Jeremy Morrell Jeremy Morrell blog notes about rss Blog
Aug 18, 2026 Extensible Software in the age of LLMs
Most of the web software we interact with today is static. The developers have a limited amount of time and attention, and focus on building the features that serve the largest group of users. The top of the demand curve is well-served by existing software, but there is a long-tail of unmet needs that’s different for every user.
Even if the developers were incredibly motivated to shove in every feature, user interfaces can only become so complex before they become unusable. Every additional feature added complicates the product for every other user. If the market for that feature is small, it can actively make the product worse for every user who doesn’t need it.
With this context the rise of LLM-assisted coding has been genuinely empowering for anyone who needed something that fell into this long tail.
Software has gotten all… squishy #
It’s become readily apparent that LLMs are really quite excellent at building Software for One . Personal apps that side-step all of the complexity and accountability of enterprise software and are custom fit for a single person’s workflow.
Pete Koomen at Y Combinator thinks there is an opportunity for what they are calling Small Software . I think they are onto something.
Agents make it easy to build personal tools for yourself or your team. But deploying, securing, and sharing that software is still far more complicated than creating it.
A cloud built for small software could remove that complexity and make bespoke tools as easy to share with a colleague as a Google Doc.
Pi is a good example of what I’m starting to think of as LLM-native software : a battle-tested core, but almost endlessly extensible just by asking, where users are able to share their customizations with others. In the past year your users have suddenly acquired the ability to speak code into existence. Most existing software can’t leverage this. Pi leans into it.
I suspect we’re going to start seeing more software following this self-extension pattern. However most of our existing examples of pluggable software are local software: AI agents, developer IDEs, mods for video games, Blender add-ons, CAD extensions. These tend to be professional tools with a high barrier to entry.
The web is the most successful software distribution system in the world. It shouldn’t be left behind.
My hypothesis is that there is a new opportunity for Extensible Software on the web . LLMs radically lower the cost of authoring extensions, and modern sandbox primitives lower the deployment cost and provide good security boundaries. We can build our app as a solid, accountable core, and allow users to safely extend it in many directions by having LLMs fill in the missing pieces. We can give our users super powers.
Disclosure: I currently work at Cloudflare, where high levels of exposure to Kenton Varda ’s writing have shaped much of my thinking here. Near the end, I’ll make the case that Dynamic Workers are a particularly good fit for this model, but I’ll cover several alternatives first.
A lot of web systems today rely on webhooks to allow the user to react to changes in the app. This ~kind of works, but it sets a really high bar for extension: building and operating a completely separate service plus dealing with whatever delivery issues arise.
I want to be able to hook into record updates and slide in my own logic. “When I attach this tag to a record, run my function”. “Do this action for me on a daily cron”.
Actually, I don’t want to have to think about that at all. I want to tell my read-it-later app:
Please send every article I fave longer than 4000 words to my <ereader of choice>
Look for new papers published on arxiv in <my specialty> each week, add your own summary of how it relates to my work at the top, and tag it with <tag>
The default algorithm completely garbles <site I read frequently> . Pull a few examples and make a custom parser for it.
And then a robot will extrude the silly bits of code, hook them into some extensions points, and make that happen. I should also be able to share what I’ve made with anyone else who might also want the same feature. 1
Here are some more areas where I’d love to see an LLM-native extension approach.
Okay, this is the obvious one. pi , deepseek , and opencode , are all experimenting in this space.
Rather than adding every new idea to its core, Pi provides stable hooks for tools, commands, events, and UI, so it can turn a request into a small TypeScript extension and reload it in place. Those extensions can then be bundled into packages that can be shared, letting the ecosystem absorb the long tail of ideas without bloating the harness itself.
However the audience of these, at least as they exist now, is fairly small. You have to be comfortable running custom software on your local machine. In corporate environments the organization has to be comfortable with you running software that no one has ever, or will ever, look at. Unless you sandbox Pi yourself, Pi extensions run with the same permissions as Pi itself.
Software engineers will find a way, but accountants, doctors, lawyers, and thousands of other professions deserve better tools too. They need agents that can be safely and easily tailored to their domain and their own workflows.
If we’re going to get more people using agents, that doesn’t mean making them software developers. It means making the software fit their needs.
All companies end up with tons of data. Employees need to view it, query it, investigate it, correlate it with this other data in this other system, find customers experiencing <problem x> , find customers about to churn, and a million more things.
A lot of companies are experimenting with allowing AI-enthusiast employees to vibe code their own tooling, maybe deploy it to a PaaS. This is directionally correct, but creates a bunch of downstream problems. Once you have hundreds or thousands of these apps, how do you maintain them? How do they get access to the data that they need? How do they get access to only the data that they need ? How can we audit what this software is doing? If we’re relying on access tokens, what are their scopes? Who rotates them? How do we make sure that we’re not logging out customer information to a third-party? How do we make sure we’re not violating GDPR?
Or a million other compliance and security things that real businesses need to worry themselves about.
What if we gave them a place to deploy code where there are no auth tokens that can leak? Where data access is handled by an internal platform team that can ensure all of the compliance boxes are checked? Give them the space to build their own automations or custom views, but safely. 2
Spoiler: This is basically Cloudflare OS .
I’ve spent a lot of my career handling tricky support tickets. Inevitably I end up digging through dashboards, searching logs, pulling data from a million different places. Let me create extensions that surface data for the user that opened the ticket from my particular system into the support interface. Give me hooks so I can kick off agents to do the first round of investigation for me, before I even look at it. If there are common tasks that I need to do like “reset specific quota X” let me add a button to my view that can do that.
Then also let me share these with my team so we can all help each other.
A lot of Observability tooling has converged towards the same feature set: a way to search your logs with the little bar graph on top. A trace waterfall view for viewing individual traces. Customizable metrics dashboards. Maybe a service map. A few are experimenting with new visualizations , especially with the rise of agents .
The venerable trace waterfall diagram is very useful for systems that are shaped as request / response, where you mainly care about latency and success rate. A lot of us are finding ourselves with systems that are a bit more… stateful… or dynamic. Modern apps are running non-deterministic agents or durable workflow engines where a single action might take hours or days. Trace spans are a great source-of-truth to build upon, but let me experiment with my own visualizations (or install someone else’s). 3
Beyond pretty things I can look at, let me inject my own logic:
arbitrary transforms for data on ingestion
have alarms kick off my own scripts: deterministic code or my own agent
give me options to run my own code at times of highest risk: deploys or feature flag rollouts
if I have a special MyResourceID in my logs, let me turn that into a link that goes straight to that resource on another platform
All software should probably look like this
an architectural change we made in opencode2 is nearly everything is an internal plugin
there's 68 of them that cover our built in agents, integrations, config loading, etc
this means you can disable any behavior and we also properly dogfood our plugin apis
Extensible software on the web is… harder #
I just made all of that sound easy. It’s nothing of the sort.
I’m a big fan of Obsidian, both as a tool I use every day and as a piece of software.
It seems like a basic markdown editor, but with a few clicks you can extend it to do just about anything: track your tasks in a kanban board or turn your notes into a database . Want to shove all your notes into a vector database for semantic search? Go for it ! And if you want to go further, the underlying web UI primitives are easily hackable.
However that power comes with a cost: Obsidian’s extension model requires you to trust every plugin you install. A plugin can basically do anything . Obsidian fights this security challenge with automated and manual review and by verifying plugin authors.
For a notes app this is likely the right tradeoff. It works because the stakes are low and the community is relatively small. But this model falls apart the moment you want the same level of extensibility in software holding other people’s data: customer records, financial transactions, private messages. Extensibility and web services have always been a challenge.
Executing arbitrary code is rife with security and abuse challenges. An incomplete list:
Errors or infinite loops in the user’s code should never take down your service
With access to keys, customer extensions can forward them to a third party
Likewise if you expose sensitive data, make sure it can’t be exfiltrated
Make sure this system can’t be abused to do a Denial of Service attack
Make sure the user can’t accidentally Denial of Service you
Protect against Spectre attacks
If people can use free compute to mine crypto on your dime, they will
But surely someone has done this? #
Before we write this off as infeasible, there is a clear example where this kind of extensibility on the web has worked at immense scale: Salesforce.
Yes, that Salesforce. And they’ve been doing it since 2007 . (As a point of reference, AWS S3 and EC2 were launched in 2006 .)
Ask most technologists what Salesforce is and you’ll either get a blank stare or maybe something to the effect of “Aren’t they a CRM?”. However it’s more accurate to describe Salesforce as a massive multi-tenant programmable platform. In the nascent cloud era this cut against the grain: no containers, and forcing people into writing this weird, custom Java-like language, Apex .
However with the rise of serverless, the platform starts to look a lot more familiar. Consider some examples:
If I need to expose a custom endpoint, I can do so with a few lines of code. The platform handles routing, authentication, tenant isolation, execution. There is no webserver to deploy. Squint and you can see it as a precursor to modern serverless .
@RestResource (urlMapping = '/customer-health' )
global with sharing class CustomerHealthApi {
@HttpGet
global static Account getCustomer () {
String accountId =
RestContext.request.params.

[truncated]
