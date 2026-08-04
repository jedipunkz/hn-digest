---
source: "https://ardor.cloud/blog/ardor-getblock-agentic-operations"
hn_url: "https://news.ycombinator.com/item?id=49164831"
title: "We gave an AI agent rootless VPN access to 1k live servers"
article_title: "How Ardor runs inside GetBlock’s production stack | Ardor — The Fastest Way to Build Agentic Software"
author: "Daniiar9"
captured_at: "2026-08-04T06:24:59Z"
capture_tool: "hn-digest"
hn_id: 49164831
score: 2
comments: 0
posted_at: "2026-08-04T06:08:27Z"
tags:
  - hacker-news
  - translated
---

# We gave an AI agent rootless VPN access to 1k live servers

- HN: [49164831](https://news.ycombinator.com/item?id=49164831)
- Source: [ardor.cloud](https://ardor.cloud/blog/ardor-getblock-agentic-operations)
- Score: 2
- Comments: 0
- Posted: 2026-08-04T06:08:27Z

## Translation

タイトル: AI エージェントに 1,000 台のライブ サーバーへのルートレス VPN アクセスを許可しました
記事のタイトル: GetBlock のプロダクション スタック内で Ardor がどのように実行されるか | Ardor — エージェント ソフトウェアを構築する最速の方法
説明: GetBlock が Ardor VPN にどのようにアクセスしてライブ運用システムを運用し、内部ツールを構築し、運用を自動化し、月額 30,000 ドル以上の節約を実現したかをご覧ください。

記事本文:
Ardor が GetBlock の運用スタック内でどのように実行されるか
ほとんどの AI ツールはコード生成で止まります。 Ardor はライブ本番システム内で動作します。このケーススタディでは、GetBlock がどのようにして Ardor VPN へのアクセス、運用認証情報、ツールの構築、インフラストラクチャの運用、隠れた運用上の問題の表面化、ブロックチェーン インフラストラクチャ企業内での目に見えるコスト削減を実現するための自律性を提供したのかを探ります。
GetBlock は、通常の運用モデルを破壊する規模でブロックチェーン インフラストラクチャを実行します。私たちは 130 以上のブロックチェーン (このセグメントで最大の名簿)、3 つのデータセンターにまたがる 1,000 台を超える物理サーバー、そして毎月 350 ～ 400 億のリクエストについて話しています。 Fortune Crypto Top 100 企業のうち 5 社がこれらの企業に依存しています。
その規模では、インフラストラクチャの管理は簡単ではなく、途方もない作業になります。ほとんどの Web3 チームにとって、社内でノードを実行することが悪夢である理由はまさにこれです。サポートされているすべてのチェーンにわたるノード フリートのバランスを常に調整し、それぞれに独自の稼働時間計算を適用します。クライアントの負荷パターンごとに異なる動作をする接続プール層。ベア メタル パークは複数のリージョンに広がり、それぞれに独自のコストとメンテナンス プロファイルがあります。
60,000 人のアクティブな顧客にサービスを提供する場合、24 時間の稼働時間は目標ではありません。それはベースラインです。エンジニアリング上の義務は明確です。それは、クライアントの稼働時間と応答品質が他のすべてに優先することです。それは無慈悲な優先順位付けを意味します。エンジニアがフォーチュン 100 のクライアントのピーク負荷時にノードがドロップしないようにすることに集中すると、内部コストの最適化は当然バックログの最後に落ちてしまいます。
すべてを視界に入れておくことがゲーム全体です。現在利用可能な AI ツールのほとんどは、この現実から安全な距離を保って動作します。彼らは IDE でコードを書き、プロトタイプを足場にして、最悪の場合の失敗が更新であるサンドボックス内で実行します。

GetBlock は、AI エージェントを本物の中で動作させたらどうなるかを調べることにしました。
つまり、GetBlock は Ardor を自社のインフラストラクチャに導入したり、内部統合を事前に構築したりしませんでした。彼らは単に、VPN アクセスと読み取り専用資格情報という、信頼できる従業員が受け取るものと同じ運用アクセスを Ardor に与えただけです。
内部に入ると、Ardor は GetBlock のシステムを調査し、インフラストラクチャがどのように編成されているかを学び、ブロックチェーン ノードの運用、内部ダッシュボード、サポート ワークフロー、ETL パイプライン、データベース、その他の実運用サービスなど、実際の運用作業を完了するために必要な内部ツールとサービスの構築を開始しました。
これは実稼働環境でのエージェント操作の様子です。エージェントは既存の操作環境に入り、システムについての独自の理解を構築し、必要なツールを作成し、見つけたものに基づいて動作します。
GetBlock が VPN の背後に Ardor を置いた理由
GetBlock の CEO、Vasily Rudomanov 氏は、自分が何を求めているかについて明確でした。それは、上級エンジニアが操作するのと同じ方法でライブ システムを操作できるエージェントです。彼は、実際の認証情報を保持し、実際の運用データにアクセスし、必要なときにアクションを実行できるようにしたいと考えていました。市場にあるほとんどの AI ツールではそれができません。彼らは星系から遠すぎてオペレーターとして役に立たなかった。
それはかなりの信頼のハードルです。 AI システムにライブ認証情報を与え、実稼働 VPN の背後でアクセスできるようにすることは、セキュリティ チームを会議に参加させる一種の決定です。実稼働のブロックチェーン インフラストラクチャ環境全体でこれを実行すると、会議が長くなるような決断が必要になります。
Vasily は Ardor を GetBlock のインフラストラクチャに導入したり、そのためのカスタム統合を構築したりしませんでした。彼は、信頼できるエンジニアが初日から受け取ることになるものを Ardor に提供しました。それは、VPN アクセス、読み取り専用の実稼働資格情報、およびアクセス権です。

○ライブ環境。
そこから、Ardor は GetBlock のシステムをマッピングし、その仕組みを学習し、実際の運用作業を完了するために必要な内部ツールとサービスを構築しました。これには、ブロックチェーン ノードの操作、内部ダッシュボード、サポート ワークフロー、ETL パイプライン、メタベース、その他の実稼働サービスが含まれます。 Ardor は、事前に構築された統合に依存するのではなく、新しい問題に直面したときに必要なものを作成しました。
「Ardor は良い意味で飽くなき好奇心を持っています。Ardor に 1 つの質問をすると、知らなければならなかった 3 つのことが浮かび上がってきます。」
— GetBlock CEO、ヴァシリー・ルドマノフ氏
決定はレバレッジに関するものでした。 GetBlock の運用面は幅広く、ブロックチェーン ノードの運用、内部分析、ダッシュボード、サポート ワークフロー、マーケティング データ パイプラインなどです。これらすべての層にわたって有意義に働くことができた人材は、会社が手薄に配置する余裕のなかった一握りの上級エンジニアでした。
Ardor にこれらのエンジニアと同じ運用環境へのアクセスを提供することで、GetBlock は単に個々のタスクを自動化するだけではありませんでした。これらにより、エージェントは環境について独自の理解を構築し、必要なツールを作成し、企業の運用面全体でコンテキストを保持し、見つけたものに基づいて行動できるようになりました。それが彼らが探していた種類の影響力でした。
GetBlock のプロダクション スタック内で Ardor が日々行っていること
Ardor は毎日、部門を超えたシニア エンジニアが扱うのと同じ一連のシステムで動作しています。ただし、継続的に、すべてを一度に実行する場合は除きます。
ブロックチェーン ノードの運用において、Ardor は Solana WSS 接続の問題をデバッグし、コンピューティング ユニット外イベントを監視し、フェイルオーバー支出を検査し、フリート全体の成功率を比較します。
内部分析は別のレイヤーです。エージェントはメタベース内で直接動作します

、アクセスをリセットし、数値がチームの期待と一致しない場合にチャーン ロジックを調査し、誰も手動で系統を追跡することなく、データマートと生データの不一致を CEO に説明します。
データ パイプラインも日常の一部です。 5 月 6 日から 11 日までのイベント期間に Mixpanel のデータが欠落していたとき、Ardor はデータをバックフィルし、ギャップが再発しないように 1 時間ごとの ETL をデプロイしました。後続のセッションでは、Google 広告のランディング ページのアトリビューション フローが更新されました。
サポートとチケットの操作は同じワークスペースにあります。 Ardor は Jira と Zendesk をリンクし、サポート チケットをカウントして分類し、ユーザーを製品、プロトコル、地域にマッピングします。
Ardor は動作するソフトウェアも出荷しています。 5 月中旬、エージェントは SitemapHQ をエンドツーエンドで構築してデプロイしました。これは、Prisma を備えた Next.js アプリケーション、パブリック ビュー、管理パネル、インポート/エクスポート機能、検索、マーメイド モード、および完全なデプロイ フローです。要件からライブ デプロイメントまで、同じ VPN の背後で行われます。
それだけではありませんでした。コンテンツ チームのエンジニアのバックグラウンドはゼロですが、チームの悩みはよく理解している人が、複雑なブランド ガイドラインとパーソナライズされた意見を厳密に遵守しながら、11 の異なるチャネルにわたるコンテンツを計画および生成するためのツールを求めました。 Ardor もそれを一から構築しました。フロントエンド、バックエンド、データベース、完全な K8s デプロイメントなど、セットアップ全体を処理し、平易な英語のリクエストから実際の運用サービスまでを処理しました。
パターンはこれらすべてで一貫しています。エージェントは、質問が関係するあらゆるシステムにわたって動作し、見つけたものを表面化します。
このテクノロジーは注目を集めています。同様に注目に値するのは、ビジネス開発、マーケティング、販売など、端末を決して開かない人々に対して Ardor が行っていることです。
これらのワークロードに対しては、微妙だが根本的な処理を実行します。

つまり、ハードルが下がります。
Ardor は GetBlock が実行するすべての手段に接続し、サイトにアクセスした人、ファネル内をどのように移動するか、何を購入するか、どこから来たのか、何が戻ってくるのかを追跡します。このプラットフォームはすべてに組み込まれており、エンジニアリングの知識のない人にとっては目や耳になります。 Docs セクションがクライアントのオンボーディングに機能しているかどうか、また改善するには何を変更すればよいかを知りたいとします。 Ardor を開いて尋ねます。
これはまったく新しいレベルの抽象化です。
Python や SQL はもう必要ありません。 Arango、Metabase、Mixpanel、Google Analytics 4、または Ahrefs について知る必要はなくなり、ましてやそれらをクエリする方法は言うまでもありません。どの番号がここに存在し、どの番号がそこに存在し、誰がそれらの間のパイプラインを維持しているかなど、頭の中で地図を保持する必要はもうありません。
Ardor を使用すると、これらのシステムが存在することを知る必要はありません。
このプラットフォームは、マーケティング、分析、ビジネス分析データの内部に十分深く組み込まれているため、マーケティングジェネラリスト、コンテンツ開発者、テクニカルライター、グロース PM など、一般のオペレーターは平易な英語でプラットフォームに話しかけ、必要なものをすべて引き出すことができます。
ここから組織図が曲がり始めます。ほとんどの企業ではデータには所有者がおり、その所有者は通常エンジニアリング担当者です。他の全員がリクエストを提出し、ダッシュボードを待ちます。マーケティングは製品に数字を尋ね、製品は分析に尋ね、分析は倉庫に尋ねます。
熱意はその行列を崩壊させます。 1 つのシステムが、オンボーディングに関するマーケティング担当者の質問とレイテンシに関するエンジニアの質問に同じくらい簡単に答えられるようになると、製品ツールとマーケティング ツールの区別はもはや成り立ちません。
単一の操作レイヤーがあり、誰もが同じ方法でそれに対して話します。
誰も想像できなかった 6 桁の最適化
ag の最も明確な図

本番環境でのエンティックな運用手段は、誰も Ardor に特に何かを依頼するわけでもなく、単に「当社のインフラストラクチャを探索して、コストを最適化できるものを検討してください」という広範で自由なプロンプトを提供するだけだったときに生まれました。
ネットワークが 1,000 台以上のサーバーにまたがり、毎月 400 億件のリクエストを処理する場合、オーバープロビジョニングは生き残るための戦術です。容量を使い果たすことは絶対に避けたいです。しかし、ネットワークが進化し、クライアントの負荷が変化すると、当然、ハードウェアの一部が立ち往生します。上級エンジニアは、トラフィックゼロのノードを見つけるために 1,000 台のサーバーを手動で監査することに日々を費やす必要はありません。彼らは顧客のために接続を維持するのに忙しいのです。
チーム メンバーがエージェントに、GetBlock のインフラストラクチャに関する日常的な質問をしました。インフラストラクチャを運用している人であれば、週に何十回も尋ねるような質問です。
そのたった 1 つの緩い指示を与えられて、アーダーは仕事に取り掛かりました。世界中のフリート全体でテレメトリを取得し、請求データを実際の RPC トラフィックと相互参照しました。
Ardor は、GetBlock のベアメタル サーバーのインベントリを作成するだけで、不均一な最適化面を特定しました。
財務チームは、誰も手動で作成する必要のない廃止リストを受け取り、年間コストを 36 万ドル削減しました。
仕組みは複雑ではありません。エージェントはアクセス権とコンテキストを持っています。当面の質問に答えても終わりではありません。大規模なクライアントが稼働時間に依存している場合、人間が時間を無駄にすべきではない徹底的なクリーニングと最適化を実行します。目に見える操作面を観察し続け、意味のないものを表面化します。
これがカテゴリにとって何を意味するか
エンジニアリング チームが現在評価しているほとんどの AI 開発ツールは、コードの生成またはアプリの生成の 2 つのいずれかを実行します。どちらもそれぞれの分野で役に立ちます。どちらも、運用システムが実際に存在する層では動作しません。
次のカテゴリー

エージェントをコーディングした後のウェーブである y は、エージェント操作です。エージェントは実際の認証情報を使用してセキュリティ境界の背後で実行され、ライブ システムに対して動作します。可能な場合は自己修復し、できない場合は問題を表面化し、破壊的な行為を行う前に人間の承認を求めます。
GetBlock がその証拠です。ライブ ブロックチェーン インフラストラクチャ、実際の認証情報、実際の運用データ、エージェントが尋ねられずに取った行動に基づいて追跡された実際のコスト削減。これは、オペレーション層が実際に実行されているときの様子です。
仕事はコードをより速く書くことではありません。仕事は、チームが忙しいために気づかないことに気づくことです。
AI を活用した研究の本当のボトルネックはインテリジェンスではありません
AI エージェントの保護: モデルが壊れる前に信頼が壊れる理由
リポジトリ インテリジェンス: コード生成が簡単だった理由
「95% AI 失敗」という統計は本当です。解釈はそうではありません。
2026 年の AI 計算: 開発チームが無視できない 5 つのトレンド
AI を活用した研究の本当のボトルネックはインテリジェンスではありません
AI エージェントの保護: モデルが壊れる前に信頼が壊れる理由
リポジトリ インテリジェンス: コード生成が簡単だった理由
AI を活用した研究の本当のボトルネックはインテリジェンスではありません
AI エージェントの保護: Mo の前に信頼が崩れる理由

[切り捨てられた]

## Original Extract

See how GetBlock gave Ardor VPN access to operate live production systems, build internal tools, automate operations, and uncover $30K+/month in savings.

How Ardor runs inside GetBlock’s production stack
Most AI tools stop at code generation. Ardor operates inside live production systems. This case study explores how GetBlock gave Ardor VPN access, production credentials, and the autonomy to build tooling, operate infrastructure, surface hidden operational issues, and deliver measurable cost savings inside a blockchain infrastructure company.
GetBlock runs blockchain infrastructure at a scale that breaks normal operational models. We're talking 130+ blockchains (the biggest roster in the segment), over 1,000 physical servers spread across three data centers, and 35-40 billion requests monthly. Five of the Fortune Crypto Top 100 companies rely on them.
At that scale, managing infrastructure is a non-trivial, monumental task. It’s exactly why running nodes in-house is a nightmare for most Web3 teams. You are constantly balancing a node fleet spanning every supported chain, each with its own uptime math. A connection pool layer that behaves differently under every client's load pattern. A bare metal park spread across regions, each with its own cost and maintenance profile.
When you serve 60,000 active customers, twenty-four-hour uptime isn't a goal; it's a baseline. The engineering mandate is clear: client uptime and response quality over absolutely everything else. Which means ruthless prioritization. When engineers are focused on ensuring a node doesn't drop during peak load for a Fortune 100 client, internal cost optimization naturally falls to the bottom of the backlog.
Keeping all of it in view is the entire game. Most AI tools available today operate at a safe distance from this reality. They write code in IDEs, scaffold prototypes, and run inside sandboxes where the worst-case failure is a refresh.
GetBlock decided to find out what would happen if they let an AI agent operate inside the real thing.
In short: GetBlock did not deploy Ardor into its infrastructure or pre-build internal integrations for it. They simply gave Ardor the same production access a trusted employee would receive: VPN access and read-only credentials.
Once inside, Ardor explored GetBlock's systems, learned how the infrastructure was organized, and began building the internal tools and services it needed to complete real operational work, including blockchain node operations, internal dashboards, support workflows, ETL pipelines, databases and other production services.
This is what agentic operations looks like in production: an agent that enters an existing operational environment, builds its own understanding of the system, creates the tooling it needs, and acts on what it finds.
Why GetBlock put Ardor behind their VPN
Vasily Rudomanov, CEO at GetBlock, was clear about what he was looking for: an agent that could operate live systems the way a senior engineer operates them. He wanted it to hold real credentials, access real production data, and take action when action was warranted. Most AI tools on the market can’t do that. They lived too far from the system to be useful as operators.
That's a substantial trust bar. Giving any AI system live credentials and access behind a production VPN is the kind of decision that gets a security team into a meeting. Doing it across a production blockchain infrastructure environment is the kind of decision that gets you a longer meeting.
Vasily didn't deploy Ardor into GetBlock's infrastructure or build custom integrations for it. He gave Ardor what a trusted engineer would receive on day one: VPN access, read-only production credentials, and access to the live environment.
From there, Ardor mapped GetBlock's systems, learned how they worked, and built the internal tools and services it needed to complete real operational work. That included blockchain node operations, internal dashboards, support workflows, ETL pipelines, Metabase, and other production services. Rather than relying on pre-built integrations, Ardor created what it needed as it encountered new problems.
"Ardor is insatiably curious in the best way. You ask it one question, and it surfaces three things you didn't know you needed to know."
— Vasily Rudomanov, CEO, GetBlock
The decision was about leverage. GetBlock's operational surface is wide: blockchain node operations, internal analytics, dashboards, support workflows, marketing data pipelines, and more. The people who could meaningfully work across all of those layers were the same handful of senior engineers the company couldn't afford to spread thin.
By giving Ardor access to the same operational environment as those engineers, GetBlock wasn't just automating individual tasks. They were enabling an agent to build its own understanding of the environment, create the tooling it needed, hold context across the company's operational surface, and act on what it found. That was the kind of leverage they were looking for.
What Ardor does inside GetBlock's production stack day to day
Day to day, Ardor operates across the same set of systems a cross-functional senior engineer would touch. Except continuously, and across all of them at once.
On blockchain node operations, Ardor debugs Solana WSS connection issues, monitors out-of-compute-unit events, inspects failover spending, compares success rates across the fleet, and more.
Internal analytics is another layer. The agent works directly inside Metabase, resets access, investigates churn logic when the numbers don't match what the team expects, and explains datamart-versus-raw-data discrepancies to the CEO without anyone manually tracing the lineage.
Data pipelines are part of the day-to-day too. When the May 6 to 11 event window was missing data in Mixpanel, Ardor backfilled it and then deployed an hourly ETL to keep the gap from recurring. Subsequent sessions updated Google Ads landing-page attribution flows.
Support and ticket operations sit in the same workspace. Ardor links Jira and Zendesk, counts and classifies support tickets, and maps users to products, protocols, and geographies.
Ardor also ships working software. In mid-May, the agent built and deployed SitemapHQ end-to-end: a Next.js application with Prisma, a public view, an admin panel, import/export functionality, search, Mermaid mode, and the full deploy flow. From requirements to live deployment, behind the same VPN.
It didn't stop there. Someone from the content team with zero engineering background - but a perfect understanding of their team's pain - asked for a tool to plan and generate content across 11 different channels, strictly adhering to complex brand guidelines and a personalized voice. Ardor built that from scratch too. It handled the entire setup: frontend, backend, database, and a full K8s deployment, taking it from a plain-English request to a live production service.
The pattern is consistent across every one of these. The agent operates across whatever systems the question touches, and surfaces what it finds.
The technology has had its share of attention. What deserves equal attention is what Ardor does for the people who never open a terminal: business development, marketing, sales.
For those workloads it does one subtle, but radical thing: it lowers the bar.
Ardor connects to every instrument GetBlock runs to track what happens on its own site: who visits, how they move through the funnel, what they buy, where they came from, what brings them back. Wired into all of it, the platform becomes eyes and ears for anyone without an engineering background. Suppose you want to know whether the Docs section is working for client onboarding, and what you might change to improve it. You open Ardor and you ask.
This is an entirely new level of abstraction.
You no longer need Python or SQL. You no longer need to know what Arango, Metabase, Mixpanel, Google Analytics 4, or Ahrefs are, let alone how to query them. You no longer need to hold the map in your head: which numbers live here, which live there, who maintains the pipeline between them.
With Ardor you do not need to know that any of those systems exist.
The platform sits deep enough inside the marketing, analytics, and business-analysis data that an ordinary operator, whether a marketing generalist, a content developer, a technical writer, or a growth PM, can talk to it in plain English and pull whatever they need.
This is where the org chart starts to bend. In most companies data has an owner, and the owner is usually engineering. Everyone else files a request and waits for a dashboard. Marketing asks product for numbers, product asks analytics, analytics asks the warehouse.
Ardor collapses that queue. When one system answers a marketer's question about onboarding and an engineer's question about latency with equal ease, the distinction between a product tool and a marketing tool no longer holds.
There is a single operational layer, and everyone speaks to it the same way.
A six-digit optimization nobody could imagine
The clearest illustration of what agentic operations means in production came when nobody was asking Ardor to do anything in particular - just a broad, open-ended prompt: "Explore our infrastructure and see what could be cost-optimized."
When a network spans 1,000+ servers and handles 40 billion requests a month, over-provisioning is a survival tactic. You never want to run out of capacity. But as networks evolve and client loads shift, some of that hardware naturally gets stranded. Your senior engineers shouldn't be spending their days manually auditing 1,000 servers to find zero-traffic nodes—they are busy keeping the connection up for your customers.
A team member asked the agent a routine question about GetBlock's infrastructure. The kind anyone running infrastructure asks dozens of times a week.
Given that single loose instruction, Ardor went to work. It pulled telemetry across the global fleet and cross-referenced billing data against actual RPC traffic.
By simply taking inventory of GetBlock’s bare metal servers, Ardor identified an uneven optimization surface.
The finance team received a decommissioning list that nobody had to manually build, cutting $360,000 in yearly costs.
The mechanic isn't complicated. The agent has access and context. It doesn't stop when the immediate question is answered. It does the deep-cleaning and optimization that humans shouldn't be wasting their time on when massive clients are depending on their uptime. It keeps looking at the operational surface it can see, and surfaces what doesn't make sense.
What this means for the category
Most AI dev tools your engineering team is evaluating right now do one of two things: generate code or generate apps. Both are useful in their domain. Neither operates at the layer where production systems actually live.
The next category, the wave after coding agents, is agentic operations. The agent runs behind your security perimeter with real credentials, operating against live systems. It self-heals where it can, surfaces issues where it can't, and asks for human approval before any destructive action.
GetBlock is the proof point. Live blockchain infrastructure, real credentials, real production data, real cost savings traced to behavior the agent took without being asked. This is what the operations layer looks like when it's actually running.
The job isn't to write code faster. The job is to notice what your team can't notice, because your team is busy.
The Real Bottleneck in AI-Powered Research Isn’t Intelligence
Securing AI Agents: Why Trust Breaks Before Models Do
Repository Intelligence: Why Code Generation Was the Easy Part
The “95% AI Failure” Stat Is Real. The Interpretation Isn’t.
2026’s AI Reckoning: 5 Trends Dev Teams Can’t Afford to Ignore
The Real Bottleneck in AI-Powered Research Isn’t Intelligence
Securing AI Agents: Why Trust Breaks Before Models Do
Repository Intelligence: Why Code Generation Was the Easy Part
The Real Bottleneck in AI-Powered Research Isn’t Intelligence
Securing AI Agents: Why Trust Breaks Before Mo

[truncated]
