---
source: "https://www.oreilly.com/radar/why-open-source-matters-for-ai/"
hn_url: "https://news.ycombinator.com/item?id=49301569"
title: "Why Open Source Matters for AI"
article_title: "Why Open Source Matters for AI – O’Reilly"
author: "BerislavLopac"
captured_at: "2026-08-14T17:45:26Z"
capture_tool: "hn-digest"
hn_id: 49301569
score: 3
comments: 0
posted_at: "2026-08-14T17:06:09Z"
tags:
  - hacker-news
  - translated
---

# Why Open Source Matters for AI

- HN: [49301569](https://news.ycombinator.com/item?id=49301569)
- Source: [www.oreilly.com](https://www.oreilly.com/radar/why-open-source-matters-for-ai/)
- Score: 3
- Comments: 0
- Posted: 2026-08-14T17:06:09Z

## Translation

タイトル: AI にとってオープンソースが重要な理由
記事のタイトル: AI にとってオープンソースが重要な理由 – オライリー
説明: 1995 年、メディアで話題になったのは、Netscape と Microsoft のどちらが Web を支配するかということでした。答えは、どちらでもないことが判明しました。Netscape と

記事本文:
メインコンテンツにスキップ
エンタープライズ向け
スキルを調べる クラウド コンピューティング Microsoft Azure
データ エンジニアリング データ ウェアハウス
ソフトウェア アーキテクチャ オブジェクト指向
ペネトレーションテスト/エシカルハッキング
ソフトスキル プロフェッショナルコミュニケーション
書籍、コース、イベントなどを検索します
計画
クラウド コンピューティング クラウド コンピューティング
データエンジニアリング データエンジニアリング
プログラミング言語 プログラミング言語
ソフトウェア アーキテクチャ ソフトウェア アーキテクチャ
ペネトレーションテスト/エシカルハッキング
書籍、コース、イベントなどを検索します
ダークモードを切り替える
AI と機械学習
ビジネス
データ
イノベーション
研究
セキュリティ オライリー学習プラットフォームをお試しください
オライリーの学習プラットフォームを使用すると、スキルを磨き、先を行くためのリソースとガイダンスが得られます。最大 14 日間無料でお試しいただけます。
オライリー プラットフォームのライブ オンライン イベントに参加して、テクノロジーを形作る専門家から学びましょう。
レーダートレンドのニュースレターを入手する
当社のプライバシーポリシーをお読みください。
O’Reilly Radar Trends to Watch ニュースレターをご購読いただきありがとうございます。
AI にとってオープンソースが重要な理由
ティム・オライリー著 2026 年 8 月 10 日 • 11 分で読めます
1995 年、メディアでの疑問は、Netscape と Microsoft のどちらが Web を支配するかということでした。答えは、どちらでもないことが判明しました。
Netscape と Microsoft はいずれも、Web サーバーとブラウザの市場を支配することを目指しており、接続の両端を制御する者が、Microsoft がパーソナル コンピュータで享受していた死の支配に匹敵するインターネット「プラットフォーム」を手に入れることになるだろうと推論しました。両社は、最も統合され、フル機能の Web サーバーを構築した企業が勝つという理論に基づいて、思いつくすべての機能を製品に直接組み込もうと競い合いました。
オープン ソースの Apache Web サーバーは、逆の賭けをしました。クリーンな拡張層を備えた Web サーバーのままなので、誰でも利用できます。

ld は、許可を求めたり、次のリリース サイクルを待たずに、何か新しいものをボルトで固定します。数年のうちに、Apache がダントツで最も人気のある Web サーバーとなり、Netscape のサーバーと Microsoft の Internet Information Server (IIS) は過去のものとなりました。人々は、合法的なプラットフォームとしての LAMP スタック (Linux、Apache、MySQL、および (Perl | Python | PHP)) について話し始めました。堀は機能ではなくモジュール性でした。そのスタックの主要な要素が他の要素が交換または拡張されても存続するという事実は、構成可能性と分散型イノベーションの力の証拠です。
2004 年にこのパターンについて書いたとき、私はそのパターンを「参加のアーキテクチャ」と呼びました。私は、当時のライセンス議論が無視していた不都合な事実を説明しようとしていたのです。私は System III の時代に Unix を使い始め、AT&T が独自のライセンスで Unix を提供していたにもかかわらず、共同プロジェクトとしてそれがどのように成功したかを見てきました。数年後、モノリシック アーキテクチャを備えた OpenOffice のような名目上はオープン ソース プロジェクトがあまりコミュニティを構築していないことに気づきました。オープンソースは単なるライセンスの問題ではなく、アーキテクチャの問題であることに気づきました。許可を求めずに作業を拡張できる標準インターフェイスを備えた小さなカーネルは、秘密のソースの重要な部分です。
このストーリーで Netscape と Microsoft を OpenAI と Anthropic に置き換えると、その反響がわかるかもしれません。
モデルの性格、デフォルト、および履歴は、少しの努力でそれらを確認し、編集できる場所に存在していました。そうしない人も増えています。先日 Drew Breunig が私に指摘したように、フロンティア モデルの新しいバージョンはそれぞれ、製品の動作を編集可能なレイヤーから重み自体に少しずつ移動させます。

研究室はそれを見ることができます、ましてやそれを変更することはできません。モデルは自分で構築するコンポーネントではなくなり、好みに合わせて調整できるようになり、レンタルするアプライアンスになり始めます。トレーニング後のトレーニングは重要ですが、それは「多様性を信頼性と引き換えにしている」ともドリュー氏は指摘します。それは多くの人にとって良い取引ですが、「本物の食品」の方が優れているとわかっているときに高度に加工された食品を提供するのと同じ種類の取引です。
オープンウェイトは単なる賭け金です
オープンソース AI に関する公の議論は、モデルの重み、その国家安全保障への影響、研究室が重みを公開するかどうか、またどのライセンスに基づいて公開するかということにあまりにも重点を置いているようです。しかし、これはオープンソースが実際に重要である理由のほんの一部しかカバーしていません。 Apache は、どちらのソースがより入手可能であるかをめぐって Netscape や Microsoft と競合することはありませんでした (そして Linux が Windows と競合することもありませんでした)。彼らはもっと重要なことをめぐって競争していた。 Red Hat の創設者であるボブ ヤングと彼のビジネス モデルについて話したときのことを覚えています。彼は「私たちが顧客に本当に売っているのはコントロールです」と言いました。オープンソースとは、アプリケーションが依存するプラットフォームが、1 つの企業からライセンスを取得した密封された箱ではなく、誰の許可も求めずに拡張してその上にビジネスを構築できるレイヤーであることを意味します。それはイノベーションの爆発を引き起こしました。これにより、Google や Amazon などの企業は、Microsoft の支配的なパラダイムから自由に成長することができました。
メインフレームから PC、インターネットに至るまで、コンピューティングのあらゆる波は同じサイクルを通過してきました。つまり、最初は分散型イノベーションであり、最終的に勝者が外堀を築くために徐々にそのサービスを閉鎖していくというものです。市場をオープンに保つのは、単一コンポーネントのライセンスではありません。より優れたコンポーネントが登場したときに、あるコンポーネントを別のコンポーネントに交換するのは非常に簡単です。
接続プロトコル

ピースはその写真の重要な部分です。 Unix ユーティリティは stdin と stdout を想定しており、シェルはそれらを接続する一種のハーネスとして機能するため、既存のツールとシームレスに動作する新しいツールを構築するのは簡単でした。そのアプローチの力の証拠は、シェルと Unix ユーティリティが、発明されてから 50 年以上経った今日、いかにエージェント ツールの共通語になっているかということです。 TCP/IP、HTTP、およびその他のインターネット プロトコルも、インターネットをオープンで構成可能に保つ上で同様の役割を果たしました。
幸いなことに、これまでのところ、AI におけるコンポーザブルなプロトコル中心のアーキテクチャがいくつかの成功を収めています。 Anthropic のモデル コンテキスト プロトコルは、その方向への破壊的な動きであり、各ペアリングのカスタム統合を行わずに、あらゆるアプリケーションがあらゆるツールやデータ ソースにアクセスできるようにするためのオープン スタンダードでした。他のオープン プロトコルと同様に、MCP は現在、Anthropic の外に Agentic AI Foundation (Linux Foundation のサブプロジェクト) に拠点を置いており、これにより独立性が少なくとも部分的に保証されています。
Isobel Moure、Ilan Strauss、そして私は、今年初めに『プロトコルとパワー』で、モデルがコモディティ化するにつれて、競争がコンテキストに合わせてスタックを引き上げていくという主張をしました。そのコンテキストにアクセスする手段を開くと、オープンウェイトかクローズドウェイトかに関係なく、市場が開かれます。これは、Apache が Web サーバーを Web アプリケーションから分離したのと同じように、コンテキストからハーネスへのモデルを分離するものです。
エージェント スキルもオープンソース AI の将来にとって重要な要素になる可能性がありますが、LAMP スタックの歴史が示すように、Perl や PHP と同じように脇に追いやられる可能性があります。それはそれでいいのです。構成可能性とは、より良いものが登場したとき、またはより多くの人がそれに同意したときに、簡単に切り替えることができることを意味します。
素敵なものもたくさんありますよw

ork は、Letta や Nous Research などのプレーヤーのポータブル メモリで行われています。 Goose や Pi のようなオープンソースのエージェント ハーネスも、人々に権力を取り戻す上で大きな役割を果たしています。特に Pi は変更可能になるように最適化されています。クロードやコーデックスのような「/exit」コマンドではなく、「/quit」コマンドをパイに与えるというマリオ・ゼヒナーの決定について語られた楽しい話があります。数え切れないほどの問題や PR が Pi のリポジトリに送信され、「/exit」を要求または実装していますが、Zechner 氏は頑固です。彼の反論は、Pi にそれをインストールに追加するよう依頼すればよいというものです。
しかし、オープンソース AI の規模と範囲という点では、上に挙げたプロジェクトは氷山の一角にすぎません。現在の AI のオープンソース ギャップ マップは、24,600 を超えるオープンソース AI プロジェクトをカバーしており、そのうちの 421 プロジェクトがオープン性、機能、導入に関して詳細にスコアリングされています。マップはスタックを 3 つの層に編成します。1) モデルと、データ セット、微調整ツール、 VLLM などの推論フレームワーク、および eval を含む関連要素。 2) ハーネスやパーソナルエージェントを含む製品およびUXレイヤー。 3) PyTorch などのコア ML フレームワーク、Ollama などのデプロイメント ツール、エッジ ハードウェアなどの基盤となるインフラストラクチャ。
現在の AI 自体は、昨年パリで開催された AI アクション サミットから生まれた官民パートナーシップです。この夏、同社は AI Potluck を発表しました。これを「完全にオープンソース コンポーネントから組み立てられた垂直統合型 AI 製品を構築する公開プロジェクト…特定の企業や国に所有されない独自の AI に代わる実行可能な代替品」と説明しています。これまでのところ、フランス政府、DeepMind や Salesforce などのテクノロジー企業、および主要な慈善活動からの 5 年間の 25 億ドルのコミットメントのうち、約 4 億ドルが支援されています。

Omidyar の AI Collaborative、マッカーサー財団、およびフォード財団を利用して。
この組織が Agentic AI Foundation などの組織とともに存在するという事実は、オープンソース AI への関心が高まっていることの証拠です。利害関係者の連合は、AI の主権、大手研究所の横暴な野心からの企業の独立、公益のためのテクノロジーへの関心など、その関心を推進する根本的な動機についても多くのことを語っています。
もう一つの要素があります。先日の会話でドリュー・ブルーニグが指摘しました。 1 つまたは 2 つの大きなクローズド モデルが AI を支配し、それらのモデルが望ましい性格、ビジネス目標、ガードレールを重み自体にますます固定化することの問題は、イノベーションの中心である多様性が減少することです。
ドリュー氏は、研究所がデフォルトの結果を出したものに満足するのではなく、奇妙にさせ、意図的にモデルを配布から除外するのが私たちの仕事だと語った。彼は、彼のチームがまさにその理由から、最近のプロジェクトで React を組み込まないことを選択した理由について説明しました。どのモデルもすでに React をよく知っているため、それを組み込むことは、真に独自のものではなく、他の全員が行っていたことの平均を出荷することを意味します。彼が GLM と Kimi を使い始めたのは、お金を節約するためではなく、より順応性があり、カスタム ハーネス内での方向性が優れているためです。そして彼は、モデルがアプライアンスになるのではなくインフラストラクチャーであり続けるために、オープンウェイトのエコシステムが正確に存続することを望んでいます。
それが参加のアーキテクチャの実際の目的です。モデル、ハーネス、アプリケーションを実際に分離する必要があります。そうすれば、何か奇妙なものを構築したい人でも、ラボのロードムなしでそれを行うことができます。

AP とガードレールが許可するかどうかを決定します。
「奇妙」というと、すべての開発者が望んでいないもののように聞こえるかもしれません。しかし、私たちが実際に話しているのは、非常に実践的なことです。上にリンクされている、多様性と信頼性のトレードオフに関する短いエッセイの中で、Drew Breunig 氏は次のように述べています。
研究室は、素人がモデルに怠惰なプロンプトを与えたときに「十分な」結果が得られる製品を出荷する必要があります。方向性がなければ、モデルはまともなものを返さなければなりません。 (Web サイトの場合は、ReAct と Tailwind で実装された Inter フォント、単色の枠線、グラデーションを備えたカードが使用されます)。 Anthropic は、このデフォルト出力を「分布収束」と名付けました。 CAIS では、@trq212 がうまく表現しています。「プロンプトに含まれていない場合は、配布されているものを入手していることになります」…。モデルの多様性が少ないほど、コーディング エージェントの信頼性は高まりますが、出力のモノカルチャーが促進されます。
O’Reilly AI Codecon の私の共同議長であるアディ・オスマニは、この記事の草稿を読んだ後、モデルの多様性を超えてこの点を捉えました。「私が一緒に働いている人はほとんど誰もウェイトをいじっていませんが、彼らはハーネスとその周りにあるもの、つまりスキル、サブエージェント、フック、コンテキスト ファイルなどをかなり頻繁に書き換えています。現在、そこに参加が起こっています。」アディは続けて、デフォルトのメモリを採用する代わりにスキルをフォークしたことに注目しました。

[切り捨てられた]

## Original Extract

In 1995, the question in the media was whether Netscape or Microsoft would control the web. The answer, it turned out, was neither.Both Netscape and

Skip to main content
For Enterprise
Explore Skills Cloud Computing Microsoft Azure
Data Engineering Data Warehouse
Software Architecture Object-Oriented
Penetration Testing / Ethical Hacking
Soft Skills Professional Communication
Search for books, courses, events, and more
Plans
Cloud Computing Cloud Computing
Data Engineering Data Engineering
Programming Languages Programming Languages
Software Architecture Software Architecture
Penetration Testing / Ethical Hacking
Search for books, courses, events, and more
Toggle dark mode
AI & ML
Business
Data
Innovation
Research
Security Try the O’Reilly learning platform
With the O’Reilly learning platform, you get the resources and guidance to keep your skills sharp and stay ahead. Try it free for up to 14 days.
Join a live online event on the O’Reilly platform to learn from the experts shaping tech.
Get the Radar Trends newsletter
Please read our privacy policy .
Thank you for subscribing to the O’Reilly Radar Trends to Watch newsletter.
Why Open Source Matters for AI
By Tim O’Reilly August 10, 2026 • 11 minute read
In 1995, the question in the media was whether Netscape or Microsoft would control the web. The answer, it turned out, was neither.
Both Netscape and Microsoft aimed to dominate the web server and browser market, reasoning that whoever controlled both ends of the connection would have an internet “platform” to rival the deathgrip that Microsoft had enjoyed on the personal computer. The two companies raced to build every feature they could think of directly into the product, on the theory that whoever built the most integrated and full featured web server would win.
The open source Apache web server took the opposite bet. It stayed a web server with a clean extension layer, so anyone could bolt something new onto it without asking permission or waiting for the next release cycle. Within a few years, Apache was far and away the most popular web server, and Netscape’s server and Microsoft’s Internet Information Server (IIS) were history. People started talking about the LAMP stack : Linux, Apache, MySQL, and (Perl | Python | PHP) as a legitimate platform. Modularity, not features, was the moat. The fact that major elements of that stack survive while others have been swapped out or extended is a testament to the power of composability and distributed innovation.
I called that pattern the architecture of participation when I wrote about it in 2004. I was trying to explain an inconvenient fact that the licensing debates of that era ignored. I had started working with Unix in the System III days, and saw how it had succeeded as a collaborative project even though AT&T offered Unix under a proprietary license. A few years later, I observed that nominally open source projects like OpenOffice with monolithic architectures never built much of a community. I realized that open source wasn’t just about licenses, but about architecture. A small kernel with standard interfaces that lets people extend your work without asking for permission is an important part of the secret sauce.
Swap out Netscape and Microsoft for OpenAI and Anthropic in this story, and perhaps you can see the echoes.
A model’s personality, its defaults, and its history used to live where you could, with a little effort, see them and edit them. Increasingly, they don’t. As Drew Breunig pointed out to me the other day , each new version of the frontier models moves a little more of the product’s behavior out of an editable layer and into the weights themselves, where nobody outside the lab can see it, let alone change it. The model stops being a component you build with and can adjust to your liking and starts being an appliance you rent. Post-training is important but Drew points out that it is also “ trading diversity for reliability .” That’s a good trade for many people, but it is the same kind of trade that gives us highly processed foods when we know that “ real food ” is better.
Open weights are just table stakes
The public debate about open source AI seems devoted far too much to model weights, their national security implications, and whether a lab releases weights and under what license. But that covers only a fraction of what actually makes open source matter. Apache was never competing with Netscape and Microsoft (and Linux was never competing with Windows) over whose source was more available. They were competing over something more important. I remember talking with Bob Young, the founder of Red Hat, about his business model, and he said “What we really sell to our customers is control.” Open source meant that the platform your application depended on was no longer a sealed box you licensed from one company but a layer you could extend and build a business on top of without asking anyone’s permission. It sparked an explosion of innovation. It enabled companies like Google and Amazon to grow up free from Microsoft’s dominant paradigm.
Every wave of computing, from mainframes to PCs to the internet, has run through the same cycle: distributed innovation at the start, with the eventual winner gradually closing down its offerings to build a moat. What keeps a market open isn’t the license on any single component. It’s how easy it is to swap out one component for another when a better one appears.
The protocols connecting the pieces are an important part of that picture. Unix utilities expected stdin and stdout, and the shell acted as a kind of harness to connect them, so it was easy to build a new tool that worked seamlessly with existing ones. A testament to the power of that approach is just how much the shell and Unix utilities are the lingua franca of agentic tooling today, more than 50 years after they were invented! TCP/IP, HTTP, and other internet protocols played a similar role in keeping the internet open and composable.
Fortunately, so far, we are seeing some wins for composable, protocol-centric architectures in AI. Anthropic’s Model Context Protocol was a disruptive move in that direction, an open standard for letting any application reach any tool or data source without a custom integration for each pairing. Along with other open protocols, MCP also now has a home outside of Anthropic at the Agentic AI Foundation (a subproject of the Linux Foundation), which is at least a partial guarantee of its independence.
Isobel Moure, Ilan Strauss, and I made the case earlier this year in Protocols and Power that as models commoditize, competition moves up the stack to context. Opening the means of accessing that context opens the market, regardless of whether open or closed weights sit underneath it. That’s an unbundling, model from harness from context, done the way Apache unbundled web server from web application.
Agentic skills may also be a critical element of the open source AI future, though as the history of the LAMP stack shows, they may fall by the wayside in the same way that Perl and PHP did. And that’s just fine. Composability means that it’s easy to switch to something better when it comes along, or when more people agree on it.
There’s also a lot of great work going on in portable memory from players like Letta , Nous Research , and others. Open source agentic harnesses like Goose and Pi are also a big part of giving power back to the people. Pi in particular is optimized to be modifiable. There’s a fun story told about Mario Zechner’s decision to give Pi a “/quit” command rather than an “/exit” command like Claude or Codex. Countless issues and PRs have been submitted to Pi’s repo, asking for or implementing “/exit”, but Zechner is stubborn. His retort is that you should just ask Pi to add it to your install.
But the projects I listed above are just the tip of the iceberg when it comes to the scale and scope of open source AI. Current AI’s Open Source Gap Map covers more than 24,600 open source AI projects!!, with 421 of them scored in depth across openness, capability, and adoption. The map organizes the stack into three layers: 1) models and associated elements including data sets, fine tuning tools, inference frameworks like VLLM , and evals; 2) the product and UX layer, including harnesses and personal agents; and 3) the infrastructure underneath, including core ML frameworks like PyTorch , deployment tools like Ollama , and edge hardware.
Current AI itself is a public-private partnership that came out of the AI Action Summit in Paris last year. This summer they announced AI Potluck , which they describe as “a public project to build a vertically integrated AI product assembled entirely from open source components… a viable alternative to proprietary AI that isn’t owned by any one company or country.” It is backed so far by roughly $400 million of a five-year, $2.5 billion commitment from the French government, tech companies including DeepMind and Salesforce, and major philanthropies including Omidyar’s AI Collaborative, the Macarthur Foundation, and the Ford Foundation.
The fact that this organization exists, along with others like the Agentic AI Foundation, is a testament to the rising tide of interest in open source AI. The coalition of interested parties also says a lot about the underlying motivations that are driving that interest: AI sovereignty, corporate independence from the overweening ambition of the major labs, and an interest in technology for the public good.
There’s another element, which Drew Breunig put his finger on in our conversation the other day. The problem with having one or two big closed models dominating AI, and having those models increasingly locking their desired personality, business goals, and guardrails into the weights themselves, is that they will reduce the diversity that is at the heart of innovation.
It’s our job, Drew said, to make it weird, to push a model deliberately out of distribution rather than to settle for whatever the labs have made the default outcome. He described how his team chose not to build in React for a recent project for exactly that reason: every model already knows React too well, so building in it means shipping the average of what everyone else was doing instead of something genuinely their own. He has started using GLM and Kimi not to save money but because they are more malleable and take direction better inside a custom harness. And he wants the open-weight ecosystem to survive precisely so that models stay infrastructure rather than becoming appliances.
That’s what an architecture of participation is actually for. We need real separation between the model, the harness, and the application, so that someone who wants to build something weird can still do it without a lab’s roadmap and guardrails deciding whether they’re allowed to.
“Weird” may make it sound like something that not all developers might want. But we’re really talking about something intensely practical. In his short essay on trading reliability for diversity, linked above, Drew Breunig put it this way:
Labs have to ship a product that delivers “good enough” results when a layperson gives a model a lazy prompt. Without direction, the model must return something decent. (If it’s a website it’ll use the Inter font, cards with a single colored border, gradients, implemented with ReAct and Tailwind). Anthropic named this default output “distribution convergent.” At CAIS, @trq212 put it well, roughly, “If it’s not in your prompt, you’re getting what’s in-distribution” …. Less diverse models make for more reliable coding agents, but they encourage a monoculture of output.
Addy Osmani, my co-chair of the O’Reilly AI Codecon , took this point beyond model diversity after reading a draft of this piece: “Almost nobody I work with is tinkering with weights, but they’re rewriting the harness and what sits around it pretty constantly—skills, subagents, hooks, context files etc etc. That’s where participation is currently happening.” Addy went on to note that forking a skill instead of adopting the default, memory

[truncated]
