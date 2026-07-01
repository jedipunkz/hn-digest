---
source: "https://www.stigg.io/blog-posts/announcing-stigg-2-0-the-usage-runtime-for-ai-products"
hn_url: "https://news.ycombinator.com/item?id=48746573"
title: "The Usage Runtime for AI Products"
article_title: "Announcing Stigg 2.0 - The Usage Runtime for AI Products"
author: "gemanor"
captured_at: "2026-07-01T13:47:30Z"
capture_tool: "hn-digest"
hn_id: 48746573
score: 1
comments: 0
posted_at: "2026-07-01T13:37:53Z"
tags:
  - hacker-news
  - translated
---

# The Usage Runtime for AI Products

- HN: [48746573](https://news.ycombinator.com/item?id=48746573)
- Source: [www.stigg.io](https://www.stigg.io/blog-posts/announcing-stigg-2-0-the-usage-runtime-for-ai-products)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T13:37:53Z

## Translation

タイトル: AI 製品の使用状況ランタイム
記事のタイトル: Stigg 2.0 の発表 - AI 製品の使用ランタイム
説明: Stigg 2.0 の発表 - AI 製品の使用ランタイム。銀行台帳のように構築されたクレジット、支出管理、リアルタイムの計測、モジュール式の BYOC 導入。

記事本文:
Stigg 2.0 の発表 - AI 製品の使用ランタイム
DOCS 顧客の価格設定ブログ ログイン 人間と話す AI を使用したインストール ブログ / Stigg の発表 Stigg 2.0 の発表
すべての AI リクエストは支出の決定となります。ミリ秒単位で作成します。
コピー https://www.stigg.io/blog-posts/payment-stigg-2-0-the-usage-runtime-for-ai-products https://www.stigg.io/blog-posts/payment-stigg-2-0-the-usage-runtime-for-ai-products 4 年前、私たちはソフトウェア会社の価格設定と権利の管理方法が根本的に壊れていると信じて Stigg の構築を開始しました。エンジニアたちは、製品を出荷する代わりに、何ヶ月もかけて自社製の請求ロジックを構築していました。すべての価格変更はデプロイメントでした。すべての企業取引はカスタム統合でした。
私たちはその問題に関して正しかった。しかし、世界は誰もが予想していたよりも早く変化しました。そしてその変化により、私たちが当初十分に理解していなかったことが明らかになりました。それは、資格が単なる請求の便宜ではないということです。これらは、AI 経済において最も重要なインフラストラクチャの抽象化になりつつあります。
本日、当社は、当社の歴史の中で最も重要なリリースである Stigg 2.0 を発表します。私たちが構築したものについて説明する前に、それを構築した理由を説明したいと思います。なぜなら、「なぜ」がすべての物語だからです。
AI 分野で最も賢い企業は社内で AI を構築しています。それはあなたを怖がらせるはずです。
ほとんど誰も公に話していないトレンドが今起きています。世界で最も技術的に洗練された企業、つまり OpenAI、Anthropic、フロンティア ラボ、AI 製品がどのようなものであるかを定義する企業は、独自の請求およびアクセス制御インフラストラクチャをゼロから構築しています。
最大手の AI フロンティアの金融エンジニアリング責任者は、私たちに直接こう言いました。「私たちが本当に必要としていたのは、リアルではないにしても、リアルタイムに近いものでした」

彼女は、彼らが市場にあるすべてのサードパーティの計測および課金プラットフォームを評価したと述べました。それらのプラットフォームは、同期的なアクセスの決定を行うことができませんでした。ほとんどは、異なる時代に構築されており、月の使用量を集計し、最後に請求書を送信します。単一の API 呼び出しに何ドルもかかる場合、エージェントがミリ秒単位でサブエージェントを生成する場合、およびバッチ ジョブの実行前に詐欺ユーザーが数千のクレジットを使い果たす可能性がある場合には、そのモデルは機能しません。
2026 年 2 月、OpenAI は「Beyond Rate Limits」を公開しました。これは、AI 企業が自社の請求内部についてリリースした文書の中で、最も構造的に明らかになった文章です。重要な概念は「決定ウォーターフォール」です。システムは、「このリクエストは許可されますか?」と尋ねる代わりに、「どこからどこまで許可されるか?」と尋ねます。すべてのリクエストは、レート制限を同期的にチェックし、クレジットを検証し、1 つの最終的な結果を返す単一の評価パスを通過します。貸方、借方は非同期で決済されます。レート制限、無料枠、クレジット、プロモーション、エンタープライズ資格はすべて、同じ意思決定スタック内のレイヤーです。
私たちはその投稿を読み、自分たちのアーキテクチャを認識しました。
しかし、これを読んでいるすべてのエンジニアリング チームが懸念すべき点は、OpenAI にはこれを構築する余裕があるということです。彼らには何百人ものエンジニアがおり、専任チームを置くのに見合った収益があります。ほとんどの企業はそうではありません。そして、「週末にクレジットカウンターを構築できる」という宣伝文句は、ソフトウェアにおける最も危険な嘘の 1 つです。はい、週末でカウンターを構築できます。 2 か月後、エッジ ケースの 30% を処理するシステムが完成します。残りの 70% は、請求に関する紛争、収益漏洩、企業顧客の怒り、および午前 3 時のページとして表示されます。
自社で構築するというトレンドは、企業が構築したいから起こっているわけではありません。それは起こっているから

既存の課金プラットフォームはそれらに失敗しました。既存企業のソリューションは、サブスクリプションが毎月更新され、使用量が項目に集約され、請求書が決定的な瞬間となる世界向けに設計されています。 AI では、決定的瞬間は API 呼び出しです。リクエストです。推論。エージェントのアクション。そのリクエストを続行するかどうかをリアルタイムで決定できない場合は、詐欺、予算超過、または信頼を損なう顧客エクスペリエンスのいずれかで、すでに負けています。
AI では、決定的瞬間は API 呼び出しです。リクエストです。推論。エージェントのアクション。そのリクエストを続行するかどうかをリアルタイムで決定できない場合は、詐欺、予算超過、または信頼を損なう顧客エクスペリエンスのいずれかで、すでに負けています。
Stigg 2.0 は、社内ビルドのパスを不要にするために存在します。 OpenAI が製品として内部的に構築したアーキテクチャ。あらゆる AI 企業に。
権利がすべてを変える抽象化である理由
「請求」が販売されたものの記録システムである場合、エンタイトルメントは履行されたものの記録システムです。この違いは、エンタープライズ ソフトウェアにおいて最も重要なものになろうとしています。その理由は次のとおりです。
AI 企業は、コマースが追いつくよりも早く機能を出荷します。 2025 年だけでも、500 社にわたる 1,800 件の価格変更を追跡しました。 OpenAI、Anthropic、Cursor、Figma、Clay、Monday.com、HubSpot、Notion、Slack - それらは 1 つずつクレジット、トークン、成果ベースの価格設定に移行していきました。シートが死にかかっています。しかし、AI 製品の進化のスピードは、もはや価格設定を導入することができないことを意味します。チームが新しいモデル、新しいエージェント機能、または新機能を毎週出荷するとき、エンタイトルメント層 (各顧客が製品に対して何を行うことが許可されているか) も同様に迅速に移行する必要があります。そのロジックをアプリケーションにハードコーディングすることは、あらゆるPRを行うための選択です。

課金プロジェクトを立ち上げます。プログラム可能で外部化されたエンタイトルメント レイヤーにより、数週間かかるエンジニアリング プロジェクトの価格変更が、数分で完了する構成変更に即座に変換されます。
エージェントを使用するには、ミリ秒単位の遅延の強制が必要です。人間がボタンをクリックすると、その権限を確認するために数百ミリ秒かかります。 AI エージェントが 50 個の API 呼び出しを並行して実行するサブエージェントを生成し、それぞれが共有クレジット プールからトークンを消費する場合、ミリ秒の時間があり、答えが正しい必要があります。最終的には一貫性がありません。 「後で仲直りしよう」ではありません。正解です。高価なことが起こる前に、今すぐに。なぜなら、エージェントの時代においては、「後で調整します」というのは、「コストを負担するか、顧客を驚かせるか」ということの言い換えにすぎないからです。どちらも戦略ではありません。これが、私たちがリクエスト パスでのミリ秒の強制にこだわる理由です。それはパフォーマンスのフレックスではありません。これは、エージェントがユーザーである場合に機能する唯一のアーキテクチャです。
企業に販売するには、ガバナンスと制御が必要です。会話のたびにそれが聞こえてきます。ある大企業は、実際の制限を動的に強制するシステムがなかったため、顧客との関係を壊さないようにするためだけに、顧客が実際に購入した金額よりも高い権利制限を設定する必要があると語った。 Cursor 社の企業顧客には、組織全体の AI クレジット割り当てを 1 日で使い果たしたユーザーがいました。エンタープライズバイヤーは「優れたダッシュボードはありますか?」と尋ねているわけではありません。彼らは、「チームが独自の予算を設定できるか? AI コストを部門に帰属させることができるか? 推論で 50,000 ドルを使い果たす前に暴走エージェントを制限できるか?」と考えています。これらの質問はすべて、請求に関する質問ではなく、権利に関する質問です。そして、これらはあると便利なものではなく、取引の要件となることが増えています。入力

ライズ社はガバナンス管理がなければ、6桁のAI契約には署名しないだろう。当社の顧客の製品リーダーの 1 人は次のように述べています。「今日、ますます多くの企業が、購入したものが実際にどのように使用されるかを自己管理する機能を求めています。」
SOX コンプライアンスでは、何が販売されたかだけでなく、何が履行されたかについての監査可能性が求められます。これはまだ誰も話していませんが、急速に実現されつつあります。 SOX では、企業が財務報告の管理を実証することが求められています。収益が使用量ベースの消費に依存している場合、「契約に記載されている内容」と「顧客が実際に受け取った内容」とのギャップが監査の対象となります。別のチームが私たちを評価したとき、問題は請求書に関するものではなく、資格の変更追跡に関するものでした。「レブレックの観点から見ると、収益を適切に認識するには、顧客の機能セットがいつ変更されたかを知る必要があります。」従来の請求システムは、請求内容を追跡します。資格システムは、何が配信されたかを追跡します。 AI 企業が成熟し、上場し、監査の精査に直面するにつれて、エンタイトルメント層は、契約、製品エクスペリエンス、収益ラインを結び付ける来歴チェーンになります。何が履行されたかを証明できなければ、収益を認識することはできません。収益に影響を与える資格には、SOC 1 Type 2 が要求するのと同じレベルの監査可能性が必要です。これがまさに Stigg がその認定を保持している理由です。
これら 4 つの力、つまり既存の請求プラットフォームの失敗、AI 製品の進化の速度、エージェント アーキテクチャのリアルタイムの要求、企業バイヤーのガバナンスとコンプライアンスの要件が現在収束しつつあります。それらはすべて同じ結論を示しています。エンタイトルメント層はもはや請求システムの機能ではありません。これは独立した基礎的なインフラストラクチャです。そしてそれを構築する必要があります

あるいは、私たちが後に残そうとしている世界ではなく、私たちが入っている世界です。
権利レイヤーは、請求システムの機能ではなくなりました。これは独立した基礎的なインフラストラクチャです。そして、それは、私たちが去ろうとする世界ではなく、私たちが参入しようとしている世界のために構築される必要があります。
Stigg 2.0 は、AI 製品用のランタイムです。課金プラットフォームではありません。計測ツールではありません。アプリケーションと請求スタックの間に位置するリアルタイムの強制およびガバナンス層で、すべての顧客、ユーザー、チーム、エージェントが実行しようとする瞬間に、各顧客、ユーザー、チーム、エージェントが何を実行できるかを決定します。
1 つの API 呼び出し。リアルタイムの決定は 1 つだけです。決定的な答えが 1 つあります。ミリ秒単位。毎回。
私たちは強制します。他の人は報告するだけです。
既存企業は、請求サイクルが終了した後に何が起こったかを教えてくれます。 Stigg は、リクエストが完了する前に何が起こってもよいかを決定します。それがダッシュボードとランタイムの違いです。
本日発送するものはすべてこちらです。
1. クレジット エンジン - ゼロから再構築
現在、あらゆる AI 製品はクレジットで実行されます。しかし、クレジットは一見難しいです。ウォレット、元帳、バーンダウン、ロールオーバー、チャージ、有効期限、優先順位に基づく消費、複数通貨プール、プロモーション補助金、定期的な割り当てなど、エッジケースは急速に増加します。
私たちはクレジット インフラストラクチャ全体を再構築し、現在はまさにこの種の作業のために設計された金融取引データベースとして構築しました。結果:
リアルタイム残高。最終的には一貫性がありません。 cron ジョブでは更新されません。ユーザーがクレジットを使用すると、API 応答が返される前に残高が更新されます。
当座貸越ゼロ。ウォレットレベルでのハードおよびソフト強制。ユーザーのクレジットが 12 残っており、アクションのコストが 15 の場合、リクエストは拒否され、遡って請求されることはありません。
優先順位に基づいた消費。プロモーション クレジットが最初に燃えます。有料クレジットは最後に燃えます。経験値

有効なクレジットは、有効期限のないクレジットよりも先に消費されます。ルールを定義するのはあなたです。それらをアトミックに実行します。
リソースごとのウォレットの分離。サブスクリプションごと、製品ごと、チームごとに個別のクレジット プール。リソース間の流出はありません。ある顧客の AI アシスタントが別の製品の予算を使い果たすことはありません。
ASC 606 準拠の台帳。すべての信用取引は完全な出所とともに記録されます。時間、イベント タイプ、アクター、またはディメンションでフィルターします。財務チームは、監査証跡を作成しなくても、必要な監査証跡を取得できます。
スタンドアロンの助成金。サブスクリプションを必要とせずに顧客にクレジットを付与します。プロモーション クレジット、パートナー割り当て、手動チャージ - すべて API 経由で、すべて同じ台帳の整合性を保ちます。
予約＆決済。消費率が事前に不明な、長期稼働エージェントのクレジット収益化を解決します。
クレジットの可視性。アラート、最も低い粒度（任意の次元）による使用状況の内訳。
2. ガバナンス - 実際に適用される AI 使用の予算と割り当て
これは私たちが最も誇りに思っている機能セットであり、他には存在しません。
今日 AI 製品を購入する企業顧客は、誰も公には語らない問題を抱えています。それは、1 人のパワー ユーザーが組織全体の AI 割り当てを消費してしまう可能性があるということです。

[切り捨てられた]

## Original Extract

Announcing Stigg 2.0 - The usage runtime for AI products. Credits built like a bank ledger, spend controls, real-time metering and modular BYOC deployment.

Announcing Stigg 2.0 - The Usage Runtime for AI Products
DOCS CUSTOMERS PRICING BLOG Login Talk to a Human Install with AI Blog / Stigg Announcements Announcing Stigg 2.0
Every AI request is a spend decision. Make it in milliseconds.
Copied https://www.stigg.io/blog-posts/announcing-stigg-2-0-the-usage-runtime-for-ai-products https://www.stigg.io/blog-posts/announcing-stigg-2-0-the-usage-runtime-for-ai-products Four years ago, we started building Stigg because we believed the way software companies manage pricing and entitlements was fundamentally broken. Engineers were burning months building homegrown billing logic instead of shipping product. Every pricing change was a deployment. Every enterprise deal was a custom integration.
We were right about the problem. But the world changed faster than anyone expected - and the change revealed something we didn't fully appreciate at the start: entitlements aren't just a billing convenience. They are becoming the most critical infrastructure abstraction in the AI economy.
Today, we're announcing Stigg 2.0 - the most significant release in our company's history. Before I walk through what we built, I want to explain why we built it. Because the "why" is the whole story.
The Smartest Companies in AI Are Building In-House. That Should Terrify You.
There's a trend happening right now that almost nobody is talking about publicly. The most technically sophisticated companies in the world - OpenAI, Anthropic, the frontier labs, the companies defining what AI products look like - are building their own billing and access control infrastructure from scratch.
The Head of Financial Engineering from one of the largest AI Frontier told us directly: "What we really needed was something that was close to real time, if not real time, that could tell us - do you have credits or not?" She said they evaluated every third-party metering and billing platform on the market. None of them could make synchronous access decisions. Most were built for a different era - aggregate usage over the month, send an invoice at the end. That model doesn't work when a single API call can cost dollars, when agents spawn sub-agents in milliseconds, and when a fraudulent user can burn through thousands of credits before your batch job runs.
In February 2026, OpenAI published "Beyond Rate Limits" - the most architecturally revealing piece of writing any AI company has released about its billing internals. The key concept is a "decision waterfall": instead of asking "is this request allowed?", their system asks "how much is allowed, and from where?" Every request passes through a single evaluation path that synchronously checks rate limits, verifies credits, and returns one definitive outcome. Credit debits settle asynchronously. Rate limits, free tiers, credits, promotions, and enterprise entitlements are all layers in the same decision stack.
We read that post and recognized our own architecture.
But here's what should concern every engineering team reading this: OpenAI can afford to build this. They have hundreds of engineers and the revenue to justify a dedicated team. Most companies don't. And the "we can build a credit counter in a weekend" pitch is one of the most dangerous lies in software. Yes, you can build a counter in a weekend. Two months later, you'll have a system that handles 30% of the edge cases - and the other 70% will show up as billing disputes, revenue leakage, angry enterprise customers, and 3am pages.
The build-in-house trend isn't happening because companies want to build. It's happening because existing billing platforms failed them. Incumbents solutions were designed for a world where subscriptions renew monthly, usage gets aggregated into line items, and the invoice is the moment of truth. In AI, the moment of truth is the API call. The request. The inference. The agent action. If you can't decide in real time whether that request should proceed, you've already lost - either to fraud, to budget overruns, or to a customer experience that breaks trust.
In AI, the moment of truth is the API call. The request. The inference. The agent action. If you can't decide in real time whether that request should proceed, you've already lost - either to fraud, to budget overruns, or to a customer experience that breaks trust.
Stigg 2.0 exists to make the build-in-house path unnecessary. The architecture OpenAI built internally - as a product. For every AI company.
Why Entitlements Are the Abstraction That Changes Everything
If "billing" is the system of record for what was sold, entitlements are the system of record for what was fulfilled. That distinction is about to become the most important one in enterprise software. Here's why.
AI companies ship features faster than commerce can keep up. In 2025 alone, we tracked 1,800 pricing changes across 500 companies. OpenAI, Anthropic, Cursor, Figma, Clay, Monday.com, HubSpot, Notion, Slack - one by one, they moved to credits, tokens, and outcome-based pricing. The seat is dying. But the speed at which AI products evolve means pricing can't be a deployment anymore. When your team ships a new model, a new agent capability, or a new feature every week, the entitlements layer - what each customer is allowed to do with your product - has to move just as fast. Hardcoding that logic into your application is a choice to make every product launch a billing project. An entitlements layer that's programmable, externalized, and instant turns pricing changes from multi-week engineering projects into configuration changes that take minutes.
Agentic usage demands milliseconds-latency enforcement. When a human clicks a button, you have hundreds of milliseconds to check their permissions. When an AI agent spawns a sub-agent that makes 50 API calls in parallel, each consuming tokens from a shared credit pool, you have milliseconds - and the answer has to be correct. Not eventually consistent. Not "we'll reconcile later." Correct, right now, before the expensive thing happens. Because in the age of agents, "we'll reconcile it later" is just another way of saying "we'll eat the cost or surprise the customer." Neither is a strategy. This is why we obsess over millisecond enforcement in the request path. It's not a performance flex. It's the only architecture that works when agents are the users.
Selling to the Enterprise demands governance and controls. We hear it in every conversation. A large Enterprise told us they had to set entitlement limits higher than what customers actually purchased - just to avoid breaking the customer relationship - because they had no system to enforce the real limits dynamically. Cursor had a user at an enterprise customer who burned down the entire organization's AI credit allocation in a single day. Enterprise buyers aren't asking "do you have a nice dashboard?" They're asking: "Can my teams set their own budgets? Can I attribute AI costs to departments? Can I cap a runaway agent before it drains $50,000 in inference?" Every one of those questions is an entitlements question, not a billing question. And increasingly, these aren't nice-to-haves - they're deal requirements. Enterprises won't sign six-figure AI contracts without governance controls. As one of our customers' product leads put it: "More and more enterprises today are looking for the ability to self-govern how what they purchased is actually being used."
SOX compliance demands auditability on what was fulfilled, not just what was sold. This is the one nobody is talking about yet, but it's coming fast. SOX requires that companies demonstrate controls over financial reporting - and when your revenue depends on usage-based consumption, the gap between "what the contract says" and "what the customer actually received" becomes an audit surface. When another team evaluated us, the question wasn't about invoices - it was about entitlement change tracking: "From a revrec perspective, we need to know when customers' feature sets were changed to recognize revenue properly." Traditional billing systems track what was billed. Entitlements systems track what was delivered. As AI companies mature, go public, and face audit scrutiny, the entitlements layer becomes the provenance chain that connects the contract to the product experience to the revenue line. If you can't prove what was fulfilled, you can't recognize the revenue. Revenue-impacting entitlements need the same level of auditability that SOC 1 Type 2 demands - which is exactly why Stigg holds that certification.
These four forces - the failure of existing billing platforms, the speed of AI product evolution, the real-time demands of agentic architectures, and the governance and compliance requirements of enterprise buyers - are converging right now. They all point to the same conclusion: the entitlements layer is no longer a feature of your billing system. It's a separate, foundational piece of infrastructure. And it needs to be built for the world we're entering, not the world we're leaving behind.
The entitlements layer is no longer a feature of your billing system. It's a separate, foundational piece of infrastructure. And it needs to be built for the world we're entering, not the world we're leaving behind.
Stigg 2.0 is the usage runtime for AI products. Not a billing platform. Not a metering tool. The real-time enforcement and governance layer that sits between your application and your billing stack, deciding what every customer, user, team, and agent is allowed to do - at the moment they try to do it.
One API call. One real-time decision. One definitive answer. In milliseconds. Every time.
We enforce. Others just report.
Incumbents tell you what happened after the billing cycle closes. Stigg decides what's allowed to happen before the request completes. That's the difference between a dashboard and a runtime.
Here's everything we're shipping today.
1. Credits Engine - Rebuilt from the Ground Up
Every AI product runs on credits now. But credits are deceptively hard. Wallets, ledgers, burn-downs, roll-overs, top-ups, expirations, priority-based consumption, multi-currency pools, promotional grants, recurring allocations - the edge cases multiply fast.
We rebuilt our entire credits infrastructure, now as a financial transactions database designed for exactly this kind of work. The result:
Real-time balances. Not eventually consistent. Not updated on a cron job. When your user spends a credit, their balance updates before the API response returns.
Zero overdraft. Hard and Soft enforcement at the wallet level. If a user has 12 credits left and the action costs 15, the request is denied - not billed retroactively.
Priority-based consumption. Promotional credits burn first. Paid credits burn last. Expiring credits burn before non-expiring ones. You define the rules. We execute them atomically.
Wallet-per-resource isolation. Separate credit pools per subscription, per product, per team. No cross-resource bleed. One customer's AI assistant doesn't drain another product's budget.
ASC 606-compliant ledger. Every credit transaction is recorded with full provenance. Filter by time, event type, actor, or dimension. Your finance team gets the audit trail they need without building it.
Standalone grants. Give credits to a customer without requiring a subscription. Promotional credits, partner allocations, manual top-ups - all through the API, all with the same ledger integrity.
Reserve & settle. Solving Credit monetization for long-running agents, where the consumption rate is unknown in advance.
Credits visibility. Alerting, usage breakdown by the lowest granularity (any dimension).
2. Governance - AI Usage Budgets and Allocations That Actually Enforce
This is the feature set we're most proud of, and the one that doesn't exist anywhere else.
Enterprise customers buying AI products today have a problem nobody talks about openly: a single power user can consume an entire organization's AI allocation i

[truncated]
