---
source: "https://venturebeat.com/data/57-of-enterprises-traced-a-wrong-ai-answer-to-missing-business-context-credible-bets-portable-open-source-semantic-code-beats-proprietary-metadata"
hn_url: "https://news.ycombinator.com/item?id=49108687"
title: "Malloy semantic modeling lang gets 'Credible' backing to build out AI context"
article_title: "57% of enterprises traced a wrong AI answer to missing business context — Credible bets portable, open-source semantic code beats proprietary metadata | VentureBeat"
author: "richij"
captured_at: "2026-07-30T12:24:13Z"
capture_tool: "hn-digest"
hn_id: 49108687
score: 1
comments: 0
posted_at: "2026-07-30T11:46:05Z"
tags:
  - hacker-news
  - translated
---

# Malloy semantic modeling lang gets 'Credible' backing to build out AI context

- HN: [49108687](https://news.ycombinator.com/item?id=49108687)
- Source: [venturebeat.com](https://venturebeat.com/data/57-of-enterprises-traced-a-wrong-ai-answer-to-missing-business-context-credible-bets-portable-open-source-semantic-code-beats-proprietary-metadata)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T11:46:05Z

## Translation

タイトル: Malloy セマンティック モデリング lang が AI コンテキストを構築するために「信頼できる」裏付けを取得
記事タイトル: 企業の 57% が、ビジネス コンテキストの欠落に対する AI の誤った回答を追跡 — 信頼できる賭けは、ポータブルなオープンソースのセマンティック コードが独自のメタデータを上回る |ベンチャービート
説明: 101 社の企業を対象とした VB Pulse の調査では、AI エージェントのエラーのほとんどがコンテキストの欠落に起因していることがわかりました。 Credible は、移植可能なオープンソースの代替品を構築するために 1,000 万ドルを調達しました。

記事本文:
企業の 57% が、ビジネス コンテキストの欠落に対する AI の誤った回答を追跡しました — ポータブルなオープンソースのセマンティック コードが独自のメタデータを上回ると確信しています | VentureBeat オーケストレーション
企業の 57% が、ビジネス コンテキストの欠落に対する AI の誤った回答を追跡 — 移植可能なオープンソースのセマンティック コードが独自のメタデータを上回るという信頼できる賭け
クレジット: FLUX-2-Pro を使用して VentureBeat によって生成された画像
企業は、以前はデータベースだけでなくビジネスを理解している人が必要だった質問に答えるために AI エージェントを信頼するようになっています。この信頼は、エージェントに正確かつ一貫したコンテキストが与えられていることを前提としていますが、市場の多くはまだそれを保証できません。
このギャップを埋めるために Credible Data が構築されています。そのプラットフォームは、オープンソースのセマンティック モデリング言語である Malloy 上で実行され、顧客のセマンティック モデルを独自のメタデータではなく移植可能なコードとして保存します。同社は月曜日、Fivetranの共同創設者であるTaylor Brown氏やG2の共同創設者Godard Abel氏を含むエンジェル投資家とともに、Gradient、SignalFire、K5 Globalからの1,000万ドルのシードラウンドを明らかにした。
問題の大きさが数字に表れます。 6月に101社の適格企業を対象に実施したVBパルスの調査では、57％が自信を持って間違ったエージェントの回答を過去6カ月間のビジネスコンテキストの欠落または一貫性の欠如に遡り、31％がそのようなことが複数回あったと回答した。現在実稼働環境で管理されたコンテキスト レイヤーを実行しているのは 25% のみで、34% がまだ構築中で、41% はまだ開始していません。
CEO 兼創設者の Kyle Nesbit 氏は、Looker 買収後の Google 勤務中に、Malloy の作成者で Looker の創設者である Lloyd Tabb 氏と出会い、Tabb 氏が自身のキャリアを決定するプロジェクトと表現する言語の構築を開始しました。ネスビット氏は、最終的に Go を辞める原因となった問題について説明した

を見て、Credible を開始します。
「これは人々が30年間抱えてきた悩みと同じ、管理されたデータ分析の欠如だ」とネスビット氏はVentureBeatの独占インタビューで語った。 「AI の登場により、同じ問題が発生しますが、混乱と苦痛は桁違いに大きくなります。」
Credible が構築するものはすべて Malloy 上で実行されます。
Malloy はオープンソースのセマンティック モデリング言語ですが、Looker 独自の LookML を含むほとんどの競合製品とは異なり、完全なクエリ言語でもあります。ネスビット氏は、組み合わせこそが他社との違いだと主張する。
「Malloy は SQL でできることなら何でもできます」と Nesbit 氏は説明しました。 「これは完全なセマンティック モデル言語とクエリ言語を組み合わせたもので、SQL で表現できるものはすべて Malloy で表現できるように設計されています。」
Malloy は、最初から MIT ライセンスを取得したオープンソース言語として Google で構築されました。その後、この基金は新しく設立された非営利団体マロイ財団に引き渡され、ネスビットはクレディブルが最初の企業スポンサーになると述べた。その結果、セマンティック レイヤーが誕生し、そのガバナンスは Credible 自身の商業的利益の範囲外にあり、顧客が構築する言語を単一のベンダーが制御できないように構造化されています。
Credible のプラットフォームは、同社がその基盤の上に構築したもので、3 つの部分に分かれています。
データコンテキスト。これはセマンティック モデルそのものであり、人間とエージェントの両方が生データを使用できるようにする定義、関係、タイプです。
ワークフローのコンテキスト。 Nesbit はこれらのスキルを、エージェントがデータを取得した後に実際にどのように使用するかについての指示と呼んでいます。 Credible は、デフォルトのスキルセットをオープンソース化し、顧客と協力して独自のワークフローに合わせて拡張しています。
携帯性。これは、モデリング言語とクエリ言語の両方としての Malloy の二重の役割を通じて実行されます。同社は、Malloy パブリッシャーと呼ばれるオープンソース サーバーを構築して、Malloy モデルをコンパイルして提供し、顧客は軽量のサーバーを実行できます。

データ モデルを展開する前に、そのバージョンをコーディング エージェントとともにローカルで構築してテストします。
「そのモデルとデータ アプリをローカルで開発したら、それは実際には単なるパッケージ、アーティファクトであり、Credible にプッシュしてエンタープライズ グレードのコントロールで提供したり、バンドルして自分で提供したりできます。」
初期の顧客からの構築か購入かの相談
アドテク企業 VideoAmp の製品管理担当シニア VP である Peter Nummerdor 氏は、Credible が選択肢として存在する前にこの問題に遭遇しました。 2023 年、VideoAmp は AWS Bedrock 上に独自の自然言語分析チャットボットを構築しました。チームはそれを市場に出すことはありませんでした。データ品質の問題によりボットの回答の信頼性が低くなり、プロジェクト終了後に再利用可能なものが何も生み出されずに費やされた労力が幻覚を引き起こすのを防ぐために、プロンプトとコンテキストのエンジニアリングが必要でした。
この経験が、このオプションが登場したときに VideoAmp が Credible を評価する方法を形作りました。
「以前に学んだ教訓を踏まえ、中核事業に集中するためには、ここを買うべきだと本当に感じました。」
VideoAmp は Credible を使用して、顧客向けデータのセマンティック モデルを構築します。 Credible は RAW ストレージとコンテキスト ウィンドウの間に位置するため、自然言語クエリで適切なメトリクスとディメンションが見つけられます。また、ラベル付きデータを VideoAmp のエージェントに返す取得レイヤーも強化します。 Nummerdor 氏は、インフラストラクチャの作業が見た目以上に重要である理由について率直に語りました。
「見栄えの良いものなら何でもすぐに構築できます。Joe Schmoe の競合他社は、Claude にデータの CSV を提供して、見栄えの良いダッシュボードを構築できます。」
コンテキストの競争市場
Credible は、AI コンテキストの課題に取り組むさまざまなアプローチの増え続けるリストに加わります。
Snowflake は最近、Snowflake とその管理者によって定義されたコンテキストを提供する Horizo​​n Context および Cortex Sense 製品を開始しました。

マーズ。 Amazon には、手動キュレーションではなくエージェントの使用方法から学習するナレッジ グラフである AWS Context サービスがあります。 Pinecone 、 Couchbase 、 Redis 、さらには Oracle を含むデータベース ベンダーはすべて、一般に基盤となるデータベースと緊密に結合されたコンテキスト レイヤーを備えています。
Credible は、オープンソースの Malloy テクノロジーを基盤として差別化を図っており、Looker によるデータのポータビリティとビジネス インテリジェンス領域の専門知識を約束します。
しかし、そのオープンさは、どの組織やハイパースケーラーでもオープンソース テクノロジーを使用して構築できるような特に深い競争堀を持たないため、同社を多少脆弱にする可能性があります。
「私たちは、私たちのような組織に集中しているとは思えない分散システムと機械学習の経験を持っています」とネスビット氏は語った。 「しかし、私たちはこの AI 革命全体の初期段階にいます。テクノロジーの曲線が平坦になった後、最先端を維持することが私たちを際立たせることになります。」
これが企業にとって何を意味するか
セマンティック コンテキスト レイヤーは、チームが理解して管理する必要がある運用インフラストラクチャ データの一部になりつつあります。 Credible のアプローチは、このカテゴリを評価するすべてのチームが現在直接直面しなければならない質問、つまり、データの意味を定義するレイヤー上でベンダーにどの程度の制御を渡すかという質問に対する 1 つの答えです。
所有権は現在、調達に関する実際の質問です。長年にわたり、セマンティック レイヤーは企業が実行するあらゆるデータ ツールの内部に存在し、プラットフォーム スイッチ上でセマンティック レイヤーがどうなるのかを尋ねる人はほとんどいませんでした。 AI エージェントは、ダッシュボードの代わりにそのレイヤーが意思決定を行うようになったため、賭け金を変更します。そのため、ポータビリティは、仮定するのではなくベンダーに直接尋ねる価値のある問題になります。
このカテゴリはまだ十分に確立されていないため、安全な答えはまだありません。チームが建築家を選択する

私たちは今、市場がどのアプローチが勝つかを告げる前に選択をしています。
洗練されたデモと運用グレードのコンテキスト レイヤーは、精査されるまでは同一に見えます。誰でも LLM をスプレッドシートに接続して、完成したように見えるものを作成できます。実際にこの 2 つを分けるものは、クエリが複雑になり、エッジケースでサポート チケットが生成され始めたときにのみ現れます。これは、VideoAmp が Credible を検討する前に、費用のかかる方法で学んだ正確な教訓です。
「データスタックが重要になるのは、ゴムが道路に接触し、品質を実際に分析するときだけです」とヌマードール氏は語った。
私の個人情報を販売または共有しないでください
私の機密個人情報の使用を制限する
© 2026 ベンチャービート。無断転載を禁じます。

## Original Extract

A VB Pulse survey of 101 enterprises found most AI agent errors trace to missing context. Credible raised $10M to build a portable, open-source alternative.

57% of enterprises traced a wrong AI answer to missing business context — Credible bets portable, open-source semantic code beats proprietary metadata | VentureBeat Orchestration
Newsletters 57% of enterprises traced a wrong AI answer to missing business context — Credible bets portable, open-source semantic code beats proprietary metadata
Credit: Image generated by VentureBeat with FLUX-2-Pro
Enterprises are increasingly trusting AI agents to answer questions that used to require a person who understood the business, not just the database. That trust assumes the agent has been given context that's accurate and consistent, and much of the market still can't guarantee that it has.
That gap is what Credible Data is built to close. Its platform runs on Malloy, an open-source, semantic modeling language, and stores a customer's semantic model as portable code rather than proprietary metadata. The company disclosed a $10 million seed round on Monday from Gradient, SignalFire and K5 Global, along with angel investors including Fivetran co-founder Taylor Brown and G2 co-founder Godard Abel.
The scale of the problem shows up in the numbers. In a VB Pulse survey of 101 qualified enterprises conducted in June, 57% traced a confidently wrong agent answer back to missing or inconsistent business context in the past six months, and 31% said it happened more than once. Only 25% run a governed context layer in production today, while 34% are still building one and 41% haven't started.
CEO and founder Kyle Nesbit met Malloy creator and Looker founder Lloyd Tabb while working at Google following its acquisition of Looker, as Tabb began building the language he has described as his career-defining project. Nesbit described the problem that eventually led him to leave Google and start Credible.
"It's the same pain point people have had for 30 years, the lack of governed data analysis," Nesbit told VentureBeat in an exclusive interview. "Now with AI, it's the same problem, but orders of magnitude more chaos and pain."
Everything Credible builds runs on Malloy.
Malloy is an open-source semantic modeling language, but unlike most competing products, including Looker's own LookML, it is also a full query language. Nesbit argues that combination is what sets it apart.
"Malloy can do anything SQL can do," Nesbit explained. "It's a full semantic model language plus query language, designed so anything you can express in SQL you can express in Malloy."
Malloy was built at Google as an open-source language, MIT-licensed from the start. It has since been handed to the newly formed Malloy Foundation, a nonprofit Nesbit said Credible will be the first corporate sponsor of. The result is a semantic layer whose governance sits outside Credible's own commercial interests, structured so no single vendor controls the language customers build on.
Credible's platform is what the company has built on top of that foundation, broken into three parts.
Data context. This is the semantic model itself, the definitions, relationships and types that make raw data usable by both people and agents.
Workflow context. Nesbit calls these skills, the instructions for how an agent should actually use that data once it has it. Credible has open-sourced its default skill set and works with customers to extend it for their own workflows.
Portability. This runs through Malloy's dual role as both a modeling language and a query language. The company built an open-source server called the Malloy publisher to compile and serve Malloy models, and customers can run a lightweight version of it locally alongside a coding agent to build and test data models before deploying them.
"Once you've developed that model and data app locally, it's really just a package, an artifact, that you can push to Credible to serve with enterprise-grade controls, or bundle and serve yourself."
A build-versus-buy call from an early customer
Peter Nummerdor, senior VP of product management at ad tech company VideoAmp, ran into this problem before Credible existed as an option. In 2023, VideoAmp built its own natural-language analytics chatbot on AWS Bedrock. The team never brought it to market. Data quality issues made the bot's answers unreliable, and the prompt and context engineering needed to keep it from hallucinating consumed effort that produced nothing reusable once the project ended.
That experience shaped how VideoAmp evaluated Credible when the option came up.
"We really felt that this was somewhere to buy, after the lessons learned earlier, to let us focus on our core business."
VideoAmp uses Credible to build semantic models of its customer-facing data. Credible sits between raw storage and the context window, so natural language queries find the right metrics and dimensions. It also powers a retrieval layer that returns labeled data into VideoAmp's agents. Nummerdor was direct about why that infrastructure work matters more than it looks like it should.
"You can build anything quickly that mostly looks good. Joe Schmoe competitor can give Claude a csv of data and build a pretty looking dashboard."
The competitive market for context
Credible joins a growing list of different approaches to tackle the challenge of AI context.
Snowflake recently launched its Horizon Context and Cortex Sense offerings, which provide context defined by Snowflake and its customers. Amazon has its AWS Context service , a knowledge graph that learns from how agents use it rather than from manual curation. Database vendors including Pinecone , Couchbase , Redis and even Oracle all have context layers that are generally tightly coupled with the underlying database.
Credible differentiates with its open-source Malloy technology at the foundation, which offers the promise of data portability and expertise in the business intelligence space from Looker.
That openness however also might leave the company somewhat vulnerable, as it doesn't have a particularly deep competitive moat as any organization or hyperscaler could build with the open-source technology.
"We have distributed systems and machine learning experience you don't see concentrated in an organization like ours," Nesbit said. "But we're at the early innings of this whole AI revolution, and staying on the cutting edge is what's going to set us apart once the technology curve flattens out."
What this means for enterprises
The semantic context layer is becoming a piece of production infrastructure data teams have to understand and manage. Credible's approach is one answer to a question every team evaluating this category now has to face directly: how much control to hand a vendor over the layer that defines what their data means.
Ownership is now a live procurement question. For years, the semantic layer sat inside whatever data tool a company happened to run and few people asked what happened to it on a platform switch. AI agents change the stakes, since that layer now drives decisions instead of dashboards, which makes portability a question worth asking a vendor directly rather than assuming.
The category is still unsettled enough that no answer is safe yet. Teams choosing an architecture now are choosing before the market has told them which approach wins.
A polished demo and a production-grade context layer look identical until scrutiny hits. Anyone can wire an LLM to a spreadsheet and produce something that looks finished. What actually separates the two only shows up once queries get complex and edge cases start generating support tickets, the exact lesson VideoAmp learned the expensive way before it looked at Credible.
"It is only when the rubber meets the road and you really analyze quality that your data stack matters," Nummerdor said.
Do Not Sell or Share My Personal Information
Limit the Use Of My Sensitive Personal Information
© 2026 VentureBeat. All rights reserved.
