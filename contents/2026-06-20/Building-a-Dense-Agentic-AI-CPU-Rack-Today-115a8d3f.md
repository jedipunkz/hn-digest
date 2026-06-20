---
source: "https://www.servethehome.com/building-a-dense-agentic-ai-cpu-rack-amd-dell-today/"
hn_url: "https://news.ycombinator.com/item?id=48608400"
title: "Building a Dense Agentic AI CPU Rack Today"
article_title: "Building a Dense Agentic AI CPU Rack Today - ServeTheHome"
author: "ksec"
captured_at: "2026-06-20T12:43:25Z"
capture_tool: "hn-digest"
hn_id: 48608400
score: 1
comments: 1
posted_at: "2026-06-20T11:20:00Z"
tags:
  - hacker-news
  - translated
---

# Building a Dense Agentic AI CPU Rack Today

- HN: [48608400](https://news.ycombinator.com/item?id=48608400)
- Source: [www.servethehome.com](https://www.servethehome.com/building-a-dense-agentic-ai-cpu-rack-amd-dell-today/)
- Score: 1
- Comments: 1
- Posted: 2026-06-20T11:20:00Z

## Translation

タイトル: 高密度エージェント AI CPU ラックを今構築する
記事のタイトル: 高密度エージェント AI CPU ラックを今構築する - ServeTheHome
説明: AI エージェントを実行する側だけでなく、従来のワークロードも実行する側の両方でエージェント AI サーバーの CPU 需要を促進しているものについて説明します。

記事本文:
フェイスブック
リンクトイン
RSS
TikTok
×
ユーチューブ
フォーラム
ワークステーション
ワークステーションプロセッサ
TrueNAS / FreeNAS NAS サーバーの上位ハードウェア コンポーネント
pfSense アプライアンスの上位ハードウェア コンポーネント
napp-it および Solarish NAS サーバーの上位ハードウェア コンポーネント
Windows Server 2016 Essentials ハードウェアのおすすめ
DIY WordPress ホスティング サーバー ハードウェア ガイド
ストレージの信頼性
レイドカリキュレーター
RAID 信頼性計算ツール |単純な MTTDL モデル
編集および著作権ポリシー
ワークステーション
ワークステーションプロセッサ
TrueNAS / FreeNAS NAS サーバーの上位ハードウェア コンポーネント
pfSense アプライアンスの上位ハードウェア コンポーネント
napp-it および Solarish NAS サーバーの上位ハードウェア コンポーネント
Windows Server 2016 Essentials ハードウェアのおすすめ
DIY WordPress ホスティング サーバー ハードウェア ガイド
高密度エージェント AI CPU ラックを今すぐ構築
サーバー CPU は犬小屋から非常に重要なインフラストラクチャになりました。その理由はエージェント AI です。これは、私が数か月間組織と話し合ってきたトピックの 1 つであり、より広範な議論の材料を提供できるのではないかと考えました。現在、オンラインでの議論の多くは単に新しいクラスのワークロードとしてエージェントを実行することについてのものですが、それには十分な理由があります。これは、コンピューティング インフラストラクチャに対する純新規の需要です。それでも、それはせいぜい方程式の一部に過ぎません。 2026 年 6 月 3 日、Cloudflare CEO の Matthew Prince は、AI ボットのトラフィックがインターネット上の人間のトラフィックを凌駕していると述べました。この傾向は現実ではないふりをすることもできますが、この傾向はサーバーを実行するすべての人に影響を及ぼし、エージェント プラットフォームが日常のワークフローの一部になるにつれて事態は悪化する一方です。サーバーの CPU が加熱しているのには理由があり、今先を行く企業が大きなアドバンテージを得るでしょう。私たちは 10 年半以上サーバー CPU を扱ってきたので、私はそうすべきだと考えました。

d 人々が使用できるより広範なフレームワークを提供します。
これについてはビデオがあります。ここでは AMD EPYC サーバーと Dell サーバーを使用します。 AMD が CPU を送りました。 Dell Tech World への旅行費はデルが負担してくれました。これは後援されていると言わざるを得ません。それでも、STH サブスタックを読めば、なぜ来年 AMD EPYC について多くのことが語られ、変更されるのかは明らかです。
データセンターでは、CPU があらゆる場所にあります。これらは GPU と並んでデータを処理し、追加のメモリ プールをそれらのアクセラレータに接続します。これらは、ストレージ ノード、コントロール プレーン、Kubernetes ワークロード、ネットワーク スイッチ、さらには一部のネットワーク アダプターを実行します。クラスターを構築するときは常に、CPU が共通の分母となります。
Agentic AI は、これらの CPU の使用方法を変えます。 OpenClaw、Hermes、および類似のエージェント フレームワークなどのプラットフォームは、GPU では実行されません。これらのエージェント フレームワークは CPU 上で実行され、何かが起こったときにいつでも対応できるように、アクティブで応答性を維持する必要があります。 OpenClaw では、現在次のコマンドを実行するだけで簡単にセットアップできます。
カール -fsSL https://openclaw.ai/install.sh |バッシュ
openclaw オンボード --install-daemon
その後、出発して準備完了です。
そこから、セキュリティ層を追加し、アクセスを管理します。企業導入の場合、ほとんどのオンライン ガイドでは、パーソナル アシスタントとして OpenClaw をセットアップする方法が説明されています。企業に導入する場合は、請負業者を雇うのと同じように考えてください。すべてのアクセス権限を開くのではなく、データとサービスへの制限されたアクセスを許可します。 「Hermes または OpenClaw にすべてにアクセスできるようにする」というのが、私が YOLO AI エージェントと呼んでいるものです。
LLM 推論は通常、GPU への API を介して実行されます。 CPU 側のインフラストラクチャはその他すべてを処理します。決定的で再現可能な結果が必要な場合は、生の LLM 出力に依存してコマンドを直接実行するのではなく、LLM に CPU 上でタスクを決定的に実行するスクリプトを生成するように依頼してください。
ここ

は実践的な例です。 sudo アクセス権を持つパスワード認証された SSH ユーザーを使用して VM をセットアップし、それを LLM にフィードします。そのワークフローを通じて 100 回の呼び出しを行うと、GPT-5.5、Qwen3.5-397B、Fable-5 (再度アクセスできた場合)、Opus-4.8 などの強力なモデルでさえ、引用符の欠落または不正な形式により、かなりの数の呼び出しでエラーが発生します。私たちはこれについて実験を行いました。以前は、最初の呼び出しの約 40% が失敗していました。新しいモデルでは、ワークフローの 25% が、単純なフォーマット エラーによる破損した SSH コマンドを修正するためにループしていることが依然として確認されています。これらのエラーは、エージェントがユーザーの介入なしに修正している場合でもトークンを焼き付けます。
その特定のアクセス パス用のツールを構築すると、問題の多くが解決されます。 LLM は呼び出しをフォーマットしますが、SSH コマンドの共通部分はツールによって決定的に処理されます。その呼び出しは CPU 上で実行され、エージェントは順調に進みます。エージェントと仕事をしたことのあるほとんどの人は、これに慣れています。新規ユーザーは通常、何も知りません。エージェントが間違いに気づいてすぐに問題を解決したとしても、中間にいる多くの人々はこれを見て、トークンの使用量と時間の点でどれだけのコストがかかるか理解していない可能性があります。
ツールを呼び出してホストするにはさまざまな方法があり、多くの人が小規模なベアメタル マシンまたはクラウド VPS インスタンスにエージェント AI をデプロイしています。私たちは現在、短期間のサンドボックスが構築され、コマンドを発行し、取り壊されるのを目にしています。エージェントのパフォーマンスを決定する重要な要素は、LLM 自体とは何の関係もありません。
これが重要な理由、そしてデータ センターで CPU が重要視されている理由は、コマンドを実行および発行するエージェントが、人間のオペレーターからマシンのオペレーターに移行しつつあるまったく新しいワークロードであるためです。最終的には物事を決定的な方向に進めたいので、

それは、CPU にさらにプッシュすることです。ボットと人間のトラフィックに対する Cloudflare レーダーについて話すとき、それについて考える 1 つの方法は、これらのボットの実行が CPU 馬力を消費するということです。それこそが、業界の人々が現在注目していることなのです。
STH の読者であれば、これが方程式の一部にすぎないことをご存知でしょう。これらのボットは Cloudflare のエンドポイントや他のエンドポイントにアクセスしています。 Web サーバーは負荷を処理する必要があります。データベースやその他のアプリケーションは、これらのセッションをサポートする必要があります。誰もがエージェントを実行する場所に集中している一方で、より多くの CPU コンピューティングを使用して、AI エージェントのクエリに対して決定的な応答を提供できます。
エージェントがリクエストを行うと、多くの場合、フロントエンド アプリケーションがヒットします。たとえば、バックエンドで Oracle や SQL Server を実行している場合、これらのアプリケーションにはライセンス料金がかかる場合があります。業界ではまだこのことについて話し始めていませんが、最終的には、レガシー アプリケーションを使用したエージェント AI の時代に向けてライセンスを最適化する企業が現れるでしょう。データベースはサーバー上で実行され、ストレージ バックエンドを備えています。すべてがネットワーク経由でサービスされます。これらのサービスを提供するすべてのデバイスには CPU が搭載されています。これが、エージェント AI が市場で非常に多くの CPU 需要を生み出している理由の 1 つです。
私は通常、この点を説明するために上のスライドのバージョンを使用します。今日のインフラストラクチャは、80 億人をわずかに超える人口に合わせて最適化されています。新しいインフラストラクチャには、リクエストを行う人々とおそらくはるかに多くのエージェントがいます。そのため、アプリケーションは人間からマシンへのトラフィックだけでなく、マシンからマシンへのトラフィックに対しても最適化する必要があります。
2010 年代に私が PwC にいたとき、医療機器会社の注文から請求までのプロセスに 30 人のチームで取り組んでいました。彼らは死体実験室を持っていましたが、率直に言って、私が見るのは少し奇妙でした。もうひとつ覚えているのは、

その日、私たちは親切な女性に会いました。ファックス機から注文を受け取り（HIPAA に感謝します）、それを自分のデスクまで送り、会社の注文システムに入力してくれました。数四半期前のプロジェクトで、私は大規模なストレージ プロバイダーの価格設定、割引、取引管理のワークフローを徹底的に見直すチームを率い、非常に時間がかかる手動プロセスから脱却できるようにしました。このプロセスでは、速度のせいで直接販売とチャネル販売の両方が失われていました。私たちはワークフローの大部分を自動化し、それが完了するまでに、同社は見積書を作成し、取引価格を決定し、スピードだけで競合他社よりも取引を勝ち取るようになっていました。同社のチャネル パートナーは、取引ごとの価格設定が記載された承認済みの見積をほぼ瞬時に取得できますが、競合他社は依然として手動の承認を必要としていました。
エージェントAIについて考えるとき、私はいつもそれを思い出します。チャネル パートナーや顧客と同様に、エージェントにもタイムアウト期間が設けられます。彼らは見積もりを得るまで何日も待つつもりはありません。インフラストラクチャは進化する必要があり、そのためにエージェント AI 時代に必要なサーバー CPU が求められます。次にそれについて見ていきましょう。
HPE Discover 2026 での AMD EPYC Venice によるラックあたり 81920 コア
HPE Discover 2026 基調講演の内容
対数演算を備えた Tensordyne Napier AI プロセッサを発表
カッパコーン
2026 年 6 月 19 日午後 5 時 36 分
金曜日の読書にとても役立ちます。私は MSP を運営していますが、在庫システムの負荷の 18% が私たちがセットアップした専用の openclaw サーバーから来ている顧客がいます。まだ epyc 7003 を使用していますが、trad アプリにも重要な負荷がかかっているのは事実です。
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
STH ニュースレターに登録してください!
このサイトはスパムを低減するために Akismet を使っています。コメントデータがどのように処理されるかをご覧ください。
STH の最高の機能を毎週受信箱に配信します。厳選したものを厳選していきます

STH から毎週最高の投稿を選んで直接お届けします。
オプトインすると、ニュースレターの送信に同意したことになります。当社ではサードパーティのサービスを使用してサブスクリプションを管理しているため、いつでもサブスクリプションを解除できます。

## Original Extract

We go into what is driving agentic AI server CPU demand on both sides, running the AI agents, but also the legacy workloads

Facebook
Linkedin
RSS
TikTok
X
Youtube
Forums
Workstation
Workstation Processors
Top Hardware Components for TrueNAS / FreeNAS NAS Servers
Top Hardware Components for pfSense Appliances
Top Hardware Components for napp-it and Solarish NAS Servers
Top Picks for Windows Server 2016 Essentials Hardware
The DIY WordPress Hosting Server Hardware Guide
Storage Reliability
Raid Calculator
RAID Reliability Calculator | Simple MTTDL Model
Editorial and Copyright Policies
Workstation
Workstation Processors
Top Hardware Components for TrueNAS / FreeNAS NAS Servers
Top Hardware Components for pfSense Appliances
Top Hardware Components for napp-it and Solarish NAS Servers
Top Picks for Windows Server 2016 Essentials Hardware
The DIY WordPress Hosting Server Hardware Guide
Building a Dense Agentic AI CPU Rack Today
Server CPUs have gone from the doghouse to becoming ultra-important pieces of infrastructure, and agentic AI is the reason. This is one of those topics that I have been talking about with organizations for months, and I thought I might just put a broader discussion piece. Right now, much of the online discussion is simply on running agents as a new class of workload, and for good reason. It is net new demand on the compute infrastructure. Still, that is at best a part of the equation. On June 3, 2026, Cloudflare CEO Matthew Prince said that AI bot traffic has eclipsed human traffic on the Internet. You can pretend that trend is not real, but it impacts everyone who runs servers, and it is only going to get worse as agentic platforms become part of everyday workflows. Server CPUs are heating up for a reason, and the companies that get ahead now will have a real advantage. Since we have been doing server CPUs for a decade and a half-plus, I figured that I should give folks a broader framework to use.
We have a video for this one. We are going to use AMD EPYC and Dell servers here. AMD sent the CPUs. Dell paid for my travel to Dell Tech World. We have to say this is sponsored. Still, if you read the STH Substack , it is pretty clear why we will be talking a lot about AMD EPYC in the next year and change.
In the data center, CPUs are everywhere. They sit alongside GPUs to process data and attach extra memory pools to those accelerators. They run storage nodes, control planes, Kubernetes workloads, network switches, and even some network adapters. Whenever you build a cluster, CPUs are the common denominator.
Agentic AI changes how those CPUs get used. Platforms like OpenClaw, Hermes, and similar agent frameworks do not run on GPUs. These agent frameworks run on CPUs, and they need to stay alive and responsive so they can react whenever something happens. OpenClaw makes that straightforward to set up currently with just:
curl -fsSL https://openclaw.ai/install.sh | bash
openclaw onboard --install-daemon
Then you are off and ready to go.
From there, you add security layers and manage access. For company deployments, most guides online cover setting up OpenClaw as a personal assistant. When you deploy at a company, think about it more like bringing on a contractor. Grant constrained access to data and services, not open all access privileges. The “let Hermes or OpenClaw have access to everything” is what I have been calling a YOLO AI agent.
LLM inference usually runs over APIs to GPUs. CPU-side infrastructure handles everything else. If you want deterministic, reproducible results, ask the LLM to generate scripts that run tasks deterministically on the CPU rather than relying on raw LLM output to execute commands directly.
Here is a practical example. Set up a VM with a password-authenticated SSH user that has sudo access, then feed that to an LLM. Make 100 calls through that workflow, and even powerful models like GPT-5.5, Qwen3.5-397B, Fable-5 (if we get access again), and Opus-4.8 will error out on a significant number of calls due to missing or malformed quotations. We ran experiments on this. It used to be around 40 percent of initial calls that failed. With newer models, we still see 25 percent of workflows looping to fix broken SSH commands from simple formatting errors. Those errors burn tokens even when the agent is fixing them without user intervention.
Building a tool for that specific access path fixes much of the problem. The LLM formats the call, while the common parts of the SSH command get handled deterministically by the tool. That call runs on the CPU, and the agent stays on track. Most folks who have worked with agents are accustomed to this. New users generally have no idea. Many folks in the middle may see it and not realize how much it costs in terms of token usage and time, even if the agent realizes its mistake and quickly fixes the issue.
There are different ways to call and host tools, and many people have deployed agentic AI on smaller bare metal machines or cloud VPS instances. We are now seeing short-lived sandboxes being built, issuing commands, and torn down. The key factors that determine how well your agents perform have nothing to do with the LLM itself.
The reason this is important, and why CPUs are becoming a big deal in the data center, is that agents running and issuing commands is a net new workload that is moving from humans being the operators to machines being the operators. Since you want to eventually push things to a deterministic path, the way you do that is to push more to CPUs. When we talk about the Cloudflare Radar for Bot vs. Human traffic , one way to think about it is that running those bots consumes CPU horsepower. That is really what folks in the industry are focused on now.
If you are an STH reader, you probably know that this is only part of the equation. Those bots are hitting Cloudflare’s endpoints and other endpoints. Web servers have to handle the load. Databases and other applications have to support those sessions. While everyone is focused on where to run agents, more CPU compute can be used to provide deterministic responses to the AI agent queries.
When agents make requests, they often hit front-end applications. Those applications may have license fees if you are running Oracle or SQL Server, for example, on the backend. The industry has not started talking about it yet, but eventually, there will be firms that go out and optimize licensing for the era of agentic AI with legacy applications. Databases run on servers and have storage back-ends. Everything is serviced over the network. All of the devices providing these services have CPUs. That is one of the reasons agentic AI is creating so much CPU demand in the market.
I usually use a version of the slide above to illustrate this point. Today’s infrastructure is optimized for a population of just over 8B people. The new infrastructure has people and likely significantly more agents making requests. As a result, applications need to be optimized for machine-to-machine traffic, not just human-to-machine traffic.
When I was at PwC in the 2010s, I had a team of 30 people working on the order-to-invoice process for a medical devices company. They had a cadaver lab, which was frankly a bit freaky for me to see. The other thing I remember was the day we saw a kind woman who would take orders from the fax machine (thank HIPAA), walk them to her desk, and type them into the company’s ordering system. A project a few quarters prior, I led a team overhauling a large storage provider’s pricing, discounting, and deal management workflow so that they could go from a very slow manual process, which was losing both direct and channel sales simply due to speed. We automated a huge part of the workflow, and by the time we were done, the company was turning quotes, deal pricing, and winning deals over its competitor just based on speed. Its channel partners could get an approved quote with per-deal pricing almost instantly, whereas its competitor still required manual approvals.
I always think of that when I think of agentic AI. Agents will have timeout windows, just like channel partners and customers did. They are not going to wait days for quotes. The infrastructure must evolve, which brings us to the kind of server CPU you need for the agentic AI era. Let us get to that next.
81920 Cores Per Rack with AMD EPYC Venice at HPE Discover 2026
HPE Discover 2026 Keynote Coverage
Tensordyne Napier AI Processor Announced with Logarithmic Math
kappakorn
June 19, 2026 At 5:36 pm
Very useful Friday read. I run a MSP and we’ve got a customer where we’ve seen 18% of our inventory systems’ load is now coming from a dedicated openclaw server we’ve setup. They’re still on epyc 7003 but it’s true there’s important load on the trad apps as well.
Save my name, email, and website in this browser for the next time I comment.
Sign me up for the STH newsletter!
This site uses Akismet to reduce spam. Learn how your comment data is processed.
Get the best of STH delivered weekly to your inbox. We are going to curate a selection of the best posts from STH each week and deliver them directly to you.
By opting-in you agree to have us send you our newsletter. We are using a third party service to manage subscriptions so you can unsubscribe at any time.
