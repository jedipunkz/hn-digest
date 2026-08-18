---
source: "https://maritime.sh"
hn_url: "https://news.ycombinator.com/item?id=49354052"
title: "Show HN: Maritime, a platform for running AI agents for $1 a month"
article_title: "Maritime | Deploy AI Agents for $1/month"
image: "https://maritime.sh/og.jpg"
author: "mariagorskikh1"
captured_at: "2026-08-18T23:13:17Z"
capture_tool: "hn-digest"
hn_id: 49354052
score: 3
comments: 0
posted_at: "2026-08-18T23:02:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Maritime, a platform for running AI agents for $1 a month

- HN: [49354052](https://news.ycombinator.com/item?id=49354052)
- Source: [maritime.sh](https://maritime.sh)
- Score: 3
- Comments: 0
- Posted: 2026-08-18T23:02:32Z

## Translation

タイトル: Show HN: Maritime、月額 1 ドルで AI エージェントを実行するプラットフォーム
記事タイトル: 海事 | AI エージェントを月額 1 ドルで導入
説明: AI エージェントをデプロイ、管理、拡張する最も簡単な方法。スリープ/スリープ解除マイクロ VM とフラット プランでは、エージェントあたり月額 1 ドルになります。コードをプッシュし、ライブ API エンドポイントを取得します。
HN テキスト: こんにちは、HN。私の名前はマリアです。Maritime の共同創設者です。私たちは、顧客のために何千もの分離された AI エージェントを実行する必要がある企業のためのインフラストラクチャを構築するために、MIT で Maritime を設立しました。 OpenClaw のようなエージェント、またはカスタム フレームワークを備えたパーソナル アシスタント エージェントをセットアップし、その個別のバージョンをすべての顧客/友人に提供したいと想像してください。各顧客は、永続的な状態、シークレット、トリガー、スリープ/ウェイク動作を備えた、分離された microVM で実行される独自のエージェントを必要とします。このようなスケーラブルで安全なインフラを構築するには数か月かかり、数十万の費用がかかります。また、インフラをうまくコード化することはまだできません。したがって、Maritime がそのインフラストラクチャを処理します。料金はエージェントあたり月額 1 ドルなので、100 個の独立したエージェントを実行すると月額 100 ドルかかります。 Maritime は、数千のエージェントを実行している企業向けに設計されていますが、今日から、開発者は 3 つのエージェントを永久に無料で実行できるようになります。テンプレートを使用して OpenClaw、Hermes、DeepSeek エージェントを起動し、3 つすべてを無料で維持できます。また、CLI または SDK を通じてカスタム エージェントをデプロイし、Maritime を戦略のインフラ プロバイダーとして使用することもできます。 https://maritime.sh で試すことができます。特に開発者のエクスペリエンスに関するフィードバックをお待ちしております

記事本文:
海事 | AI エージェントを月額 1 ドルで導入できます。 海事機能 価格設定 ドキュメント ブログ ログイン サインアップ 1 つのコマンドで OpenClaw、Hermes、または ZeroClaw をスピンアップします。スピンアップ 1 千のエージェントを 1 つと同じくらい簡単に実行できます。 100 個のエージェントを実行する
1つと同じくらい簡単に。
分離された microVM、簡単なスケーリング、フルマネージドのインフラストラクチャを使用して、顧客対応の数千の AI エージェントを Maritime 上で実行します。
デモを予約する 最初のエージェントを導入するためのカードはありません
これらの企業で海事業務に基づいて開発を行っている開発者
顧客ごとに 1 人のエージェント。
バックエンドからプロビジョニングされます。
すべてのサインアップには独自のエージェントが付きます。 1 回の API 呼び出しで専用のマイクロ VM が起動され、ID がタグ付けされ、独自のシークレットを保持し、ライフサイクル Webhook を送信します。
共有ワーカーはいません。別の顧客の VM には顧客データがありません。
npm では TypeScript、PyPI では Python、その下に REST があります。
エージェントに外部 ID をタグ付けして、データベースが信頼できる情報源となるようにします。
エージェントが展開、スリープ、ウェイク、または失敗した瞬間を把握します。
完全に非対話型であるため、独自のエージェントがエージェントをプロビジョニングできます。
共有ランタイムではありません。エージェントごとに 1 つのマイクロ VM。
1つのコマンド
実行中のエージェントに。
Maritime はフレームワークを検出し、分離されたマイクロ VM を構築し、エージェントを稼働させます。配線するインフラがありません。
1 つのエージェントまたは数百のエージェントを実行します。各 VM はその作業に合わせてサイズ設定され、アイドル状態ではスリープするため、待機中のサーバーに料金を支払う必要はありません。
実際の CLI、型指定された SDK、ダッシュボード、および出発点となるテンプレートのライブラリ。接着コードはありません。
OpenClaw、Hermes、ZeroClaw、または独自のスタックを実行します。すでに使用しているモデルやツールをご持参ください。
端末から実行中のエージェントに対して 1 つのコマンドを実行します。アイドル状態のときはスリープし、トラフィックが到着するとすぐにスリープ状態から復帰します。
孤立については顧客に伝えられる可能性があります。
他の人のプロセス内のスレッドではなく、独自のカーネルを持つエージェントごとのマイクロ VM。
認証情報は保存時に暗号化され、保存されている 1 つの VM にのみマウントされます。

まで伸びる。
標準フレームワーク、任意の VM への SSH、エクスポート可能な状態、E2B、LangGraph、およびその他 8 つのプラットフォームからのガイド付き移行。
設計された唯一のプラットフォーム
数千のエージェントまで拡張できます。
E2B Daytona Modal AWS EC2 Google Cloud DigitalOcean Railway Fly.io 月額料金 $1 /月 定額 $48-198/月 $48/月 $23-68/月 $8-15/月 $6-49/月 $6/月 $5-10/月 $6+/月 セットアップ時間 分 分 分 分 時間 時間 分 分 中 コールド スタートなし なし なし なし あり なし あり なしなし なし ステートフル コンテナ ○ ○ ○ × ○ × ○ × ○ スリープ / ウェイク 組み込み一時停止 API 自動停止 × × × × × × エージェントファースト設計 ○ ○ ○ × × × × × × ゼロから約 1 分でエージェントが実行されるまで。
メモリ、ファイル、秘密はそのまま残された場所にあります。
ストリームと音声がまっすぐに通過します。
発売日のスパイクは、気づく前に処理されます。
最初のエージェントから何百万もの販売者まで、1 つのプラットフォームで対応します。
エージェント 1 人を 1 ドルで発送します。すでに使用しているフレームワークやツールを使用して構築します。
› SSH でマシンに接続し、すべてのログを読み取ります
顧客にエージェントを販売します。これを構築し、数行のコードを追加すると、Maritime はサインアップするすべての顧客に対してコピーを実行します。
› 顧客ごとに 1 つの分離されたマイクロ VM
› 一晩で 1,000 件のサインアップが処理されました
スタートアップセットアップを参照してください。
Maritime をオンプレミスで実行します。すでに所有しているハードウェアにインストールして、数千の分離されたエージェントをホストします。
› データはインフラストラクチャ上に留まります
› 既存のサーバーでの大規模な密度
› ネットワーク内の自動スケーリングと分離
一律料金設定。何も計測されていません。
スターター $20 /月 エージェントが含まれます エージェントあたり 20 の追加エージェント / 月 $1.50 追加 RAM / GB / エージェント / 月あたり、2 GB を含み、最大 8 $3 追加 SSD 5 GB / エージェントあたり、5 GB を含み、最大 100 $2 サポート Discord コミュニティ 始める 成長 $100 /月 エージェントを含む

追加エージェント 100 エージェントあたり/月 $1.25 追加 RAM/GB/エージェントあたり、2 GB 含む、最大 8 $2.50 追加 SSD/5 GB/エージェントあたり、5 GB 含む、最大 100 $1.50 サポート Private Slack + 担当者 開始 スケール $500/月 エージェント含む エージェントあたり 500 追加エージェント/月 $1.00 追加 RAM あたりGB / エージェント / 月、2 GB 含む、最大 8 $2 5 GB / エージェント / 月ごとに 2 ドルの追加 SSD、5 GB 含む、最大 100 ドル サポート プライベート Slack + 担当者 はじめる エンタープライズ カスタム エージェントが含まれる エージェントごとにカスタム追加エージェント / 月 ボリューム GB / エージェント / 月ごとに追加 RAM、2 GB 含む、5 GB / エージェント / 月ごとに最大 8 個のカスタム追加 SSD、5 GB 含む、最大 100 カスタム サポート専用お問い合わせ 無料 $0 スターター $20 /月 拡張 $100 /月 スケール $500 /月 エンタープライズ カスタム エージェントを含む 3 20 100 500 エージェントあたりのカスタム追加エージェント / 月 – $1.50 $1.25 $1.00 ボリューム 追加 RAM / GB / エージェント / 月、2 GB を含む、最大 8 – $3 $2.50 $2 カスタム追加 SSD (5 GB / エージェントあたり / 月) 5 GB が含まれ、最大 100 – $2 $1.50 $1 カスタム サポート Discord コミュニティ Discord コミュニティ プライベート Slack + 担当者 プライベート Slack + 担当者 専用 無料で始めましょう 始めましょう 始めましょう 始めましょう お問い合わせください すべてのエージェントには 1 vCPU、2 GB RAM、および 5 GB SSD が搭載されており、スリープから約 1 秒で復帰します。
Maritime スタートアップ プログラムは、ホスティングと LLM クレジットの付与によってインフラストラクチャをカバーします。
最初のデプロイ
エージェントは1ドルで。
テンプレートから開始して約 1 分で出荷します。準備ができたらスケールし、実行した分だけ支払います。
海洋デモを予約する AI エージェントを月額 1 ドルで導入できます。
フラットプラン、メーターなし。
© 2026 Maritime AI, Inc. 全著作権所有。

## Original Extract

The simplest way to deploy, manage, and scale AI agents. Sleep/wake micro-VMs and flat plans that work out to $1 per agent per month. Push your code, get a live API endpoint.

Hi HN, my name is Maria, and I’m a co-founder of Maritime. We started Maritime at MIT to build infrastructure for companies that need to run thousands of isolated AI agents for their customers. Imagine you set up an agent like OpenClaw, or a personal assistant agent with a custom framework, and want to give a separate version of it to every customer/friend. Each customer needs their own agent running in an isolated microVM, with persistent state, secrets, triggers, and sleep/wake behavior. Building such scalable and secure infra will take you months and will cost hundreds of thousands. Also, you still can't vibe code the infra well. So Maritime handles that infrastructure for you. We charge $1 per agent per month, so running 100 isolated agents costs $100 a month. Maritime is designed for companies running thousands of agents, but starting today, any developer can run three agents for free forever. You can use our templates to spin up OpenClaw, Hermes, and the DeepSeek agent and keep all three for free. You can also deploy custom agents through our CLI or SDK and use Maritime as the infra provider for your stratup. You can try it at https://maritime.sh . We’d really appreciate your feedback, especially on the developer experience

Maritime | Deploy AI Agents for $1/month maritime Features Pricing Docs Blog Log in Sign up Spin up OpenClaw, Hermes, or ZeroClaw in one command. spin up Run a thousand agents as easily as one. Run 100 agents
as easily as one.
Run thousands of customer-facing AI agents on Maritime with isolated microVMs, effortless scaling, and fully managed infrastructure.
Book a demo No card to deploy your first agent
Developers building on Maritime work at these companies
One agent per customer.
Provisioned from your backend.
Every signup gets its own agent. One API call brings up a dedicated micro-VM: tagged with your IDs, holding its own secrets, sending you lifecycle webhooks.
No shared workers. No customer data in another customer's VM.
TypeScript on npm, Python on PyPI, REST underneath.
Tag agents with external IDs so your database stays the source of truth.
Know the moment any agent deploys, sleeps, wakes, or fails.
Fully non-interactive, so your own agents can provision agents.
Not a shared runtime. A micro-VM per agent.
One command
to a running agent.
Maritime detects your framework, builds an isolated micro-VM, and puts your agent live. No infra to wire up.
Run one agent or hundreds. Each VM is sized to its work and sleeps when idle, so you never pay for a server sitting warm.
A real CLI, a typed SDK, a dashboard, and a library of templates to start from. No glue code.
Run OpenClaw, Hermes, ZeroClaw, or your own stack. Bring whatever model and tools you already use.
One command from your terminal to a running agent. It sleeps when idle and wakes the moment traffic arrives.
Isolation your customers can be told about.
A micro-VM per agent with its own kernel, not threads in someone else's process.
Credentials encrypted at rest, mounted only into the one VM they belong to.
Standard frameworks, SSH into any VM, exportable state, guided migrations in from E2B, LangGraph, and eight other platforms.
The only platform designed
to scale to thousands of agents.
E2B Daytona Modal AWS EC2 Google Cloud DigitalOcean Railway Fly.io Monthly cost $1 /mo flat $48-198/mo $48/mo $23-68/mo $8-15/mo $6-49/mo $6/mo $5-10/mo $6+/mo Setup time Minutes Minutes Minutes Minutes Hours Hours Minutes Minutes Medium No cold starts None None None Yes None Yes None None None Stateful containers Yes Yes Yes No Yes No Yes No Yes Sleep / wake Built-in Pause API Auto-stop No No No No No No Agent-first design Yes Yes Yes No No No No No No From zero to a running agent in about a minute.
Memory, files, and secrets exactly where it left them.
Streams and voice pass straight through.
A launch-day spike is handled before you notice it.
One platform, from your first agent to the millions you sell.
Ship one agent for a dollar. Build it with whatever framework and tools you already use.
› SSH into the machine, read every log
Sell agents to your customers. Build one, add a few lines of code, and Maritime runs a copy for every customer who signs up.
› One isolated micro-VM per customer
› A thousand overnight sign-ups, handled
See the startup setup Enterprises
Run Maritime on-prem. Install it on hardware you already own and host thousands of isolated agents.
› Your data stays on your infrastructure
› Massive density on existing servers
› Auto-scaling and isolation inside your network
Flat pricing. Nothing metered.
Starter $20 /mo Agents included 20 Extra agent per agent / mo $1.50 Extra RAM per GB / agent / mo, 2 GB included, up to 8 $3 Extra SSD per 5 GB / agent / mo, 5 GB included, up to 100 $2 Support Discord community Get started Growth $100 /mo Agents included 100 Extra agent per agent / mo $1.25 Extra RAM per GB / agent / mo, 2 GB included, up to 8 $2.50 Extra SSD per 5 GB / agent / mo, 5 GB included, up to 100 $1.50 Support Private Slack + your rep Get started Scale $500 /mo Agents included 500 Extra agent per agent / mo $1.00 Extra RAM per GB / agent / mo, 2 GB included, up to 8 $2 Extra SSD per 5 GB / agent / mo, 5 GB included, up to 100 $1 Support Private Slack + your rep Get started Enterprise Custom Agents included Custom Extra agent per agent / mo Volume Extra RAM per GB / agent / mo, 2 GB included, up to 8 Custom Extra SSD per 5 GB / agent / mo, 5 GB included, up to 100 Custom Support Dedicated Talk to us Free $0 Starter $20 /mo Growth $100 /mo Scale $500 /mo Enterprise Custom Agents included 3 20 100 500 Custom Extra agent per agent / mo – $1.50 $1.25 $1.00 Volume Extra RAM per GB / agent / mo, 2 GB included, up to 8 – $3 $2.50 $2 Custom Extra SSD per 5 GB / agent / mo, 5 GB included, up to 100 – $2 $1.50 $1 Custom Support Discord community Discord community Private Slack + your rep Private Slack + your rep Dedicated Start free Get started Get started Get started Talk to us Every agent comes with 1 vCPU, 2 GB RAM, and 5 GB SSD, and wakes from sleep in about a second.
The Maritime startup program covers your infrastructure with hosting and LLM credit grants.
Deploy your first
agent for $1.
Start from a template and ship in about a minute. Scale when you're ready and only pay for what you run.
Book a demo maritime Deploy any AI agent for $1/month.
Flat plans, no meters.
© 2026 Maritime AI, Inc. All rights reserved.
