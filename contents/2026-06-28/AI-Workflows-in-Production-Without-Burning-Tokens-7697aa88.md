---
source: "https://unmeshed.io/blog/bringing-ai-workflow-into-production-without-burning-tokens"
hn_url: "https://news.ycombinator.com/item?id=48712618"
title: "AI Workflows in Production Without Burning Tokens"
article_title: "Bringing AI Workflows into Production without Burning Tokens - Unmeshed"
author: "jusonchan81"
captured_at: "2026-06-28T23:22:35Z"
capture_tool: "hn-digest"
hn_id: 48712618
score: 1
comments: 0
posted_at: "2026-06-28T22:55:43Z"
tags:
  - hacker-news
  - translated
---

# AI Workflows in Production Without Burning Tokens

- HN: [48712618](https://news.ycombinator.com/item?id=48712618)
- Source: [unmeshed.io](https://unmeshed.io/blog/bringing-ai-workflow-into-production-without-burning-tokens)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T22:55:43Z

## Translation

タイトル: トークンを消費しない本番環境での AI ワークフロー
記事のタイトル: トークンを消費せずに AI ワークフローを本番環境に導入 - アンメッシュ
説明: この記事では、トークン コストを制御しながら AI を本番環境に導入し、費用対効果の方程式を有利に保ち、実際のビジネス価値を実現する方法を検討します。

記事本文:
トークンを消費せずに AI ワークフローを本番環境に導入 - アンメッシュ
ワークフロー マルチステップ ワークフローの自動化と実行 API オーケストレーション ワークフローに基づく大規模な API エンドポイントの構築と実行 バッチ ジョブ処理 タスク間の依存関係を管理しながら、第 2 レベルの粒度でジョブを実行
統合 システムとの統合、100 以上のアプリへの組み込み接続 Agentic AI ワークフロー内で AI を活用 セキュリティ エンタープライズ グレードのセキュリティ 新機能
人間による参加 承認など人間の入力が必要なワークフローを構築 意思決定エンジン 組み込みのルール エンジンを使用してワークフロー内のロジックをスケールアップ ホストされた関数 Python、JS/TS、または Go でステップを作成し、個別のホスティングなしで実行 ソリューション
Unmeshed がチーム全体の自動化ニーズにどのように役立つか
開発者 ボイラープレートやインフラ オーバーヘッドなしでワークフローを実行 RevOps オーケストレーションで収益ワークフローを加速 マーケティング 接続されたフローでキャンペーン ROI を向上 IT 自動化優先のワークフローで信頼性の高い運用を実行 人事 オンボーディング、承認、人事業務を自動化 販売 自動化されたルーティングでより迅速に取引を成立させる カスタマー サポート 調整されたワークフローでより迅速にチケットを解決する ユースケース別
Conductor からの移行 Conductor のワークフローを最新のオーケストレーションに移行 API オーケストレーション REST + GraphQL をファンアウトして統一応答を返す Autosys から移行 従来のスケジューラのワークロードを低中断で移動 休暇ワークフロー ルーティング 承認とループバックで休暇リクエストをルーティング 不正検出の決定 リスク チェックをファンアウトし、リアルタイムで返品の決定を行う ローン申請処理 ルールと人間によるレビューによるローン承認のルーティング 顧客オンボーディングと KYC オンボーディングの自動化チェックと承認付き ソリューションに関するブログを読む
オーケストレーションが効果を促進する 5 つの方法

効率性とイノベーション オーケストレーションによる速度、信頼性、配信の向上 オーケストレーションによる不正検出ワークフローの変革 オーケストレーションされた意思決定による低遅延の不正チェックの設計 アンメッシュによる物流ソフトウェアの最適化 サプライ チェーン システムを回復力のある観察可能なフローに接続 商品取引 API オーケストレーション 市場データと取引 API を大規模にオーケストレーション CA Autosys からアンメッシュへの簡単な移行 ワークフロー優先の自動化による従来のスケジューラの最新化 リソース
Unmeshed を使用して効果的に学習、構築、運用する
ドキュメント セットアップ、ガイド、API リファレンスですぐに始められます 認証と認可 安全なオーケストレーションのための SSO、アクセス ポリシー、RBAC の構成 ブログ 製品アップデート、エンジニアリング ノート、ソリューションの詳細 アンメッシュ オートメーションの理由 プラットフォームの違い、アーキテクチャの利点、移行に関する洞察 アンメッシュとその他 一般的な自動化ツールおよびオーケストレーション ツールとの比較 ガイド 一般的な自動化パターンのステップバイステップ チュートリアル 概念 アンメッシュ サポートの背後にある中心的なアイデアとアーキテクチャ パターン
ヘルプセンター お問い合わせ 専門家を雇う プロフェッショナル サービス トレーニング ツールとユーティリティ
非メッシュ CLI CI パイプラインおよびワークフロー自動化操作用のコマンド ライン ユーティリティ SDK Java、Golang、TypeScript/JavaScript、Python 用の言語 SDK ブログに戻る トークンを焼くことなく AI ワークフローを本番環境に導入
AI (または LLM の機能) を本番環境に導入することは、今日のほとんどのエンジニアにとって中心的な指標または目標です。この記事では、コストと利益の方程式が利益の範囲内に収まり、ビジネスに価値が付加されるように、トークン コストを制御しながら AI を本番環境に導入する最善の方法について説明します。
10 分で読めます 2026 年 6 月 22 日 Agentic にしましょう!
押し込み

市場ではエージェント フローが使用されます。エージェントティックとは、リクエストまたはフローの処理方法をモデルに決定させ、コンテキストを解析して理解する能力を期待して、ユースケースに最適な結果をもたらすことを期待します。その考え方は、モデルが成熟し、より「インテリジェント」になるにつれて、結果の品質がより高くなり、人間がコード化した固定アルゴリズムを上回るというものです。
これを念頭に置くと、モデル呼び出しに完全に依存するユースケースが運用環境にプッシュされることがよくあります。
たとえば、エージェントは、入力の解析、データの検証、リクエストの分類、ポリシーの確認、適切な担当者へのルーティング、応答の草案作成など、すべてモデルを呼び出すことによってユースケースを実行できます。
現在、多くのエージェント フレームワークが存在するため、これを構築するのは非常に高速であることが多く、デモは通常、管理者にとって素晴らしく印象的です。
ただし、大規模なユースケースでこれを本番環境で起動すると、モデルプロバイダーから請求書が届いたときにショックを受けるでしょう。モデルが進化するにつれて、トークンのコストは減少するどころか増加しています。さて、そのユースケースは、実行コストに十分な価値を加えているでしょうか?
一貫性、遅延、セキュリティ、ガバナンスなどの問題についてはどうですか?そしてさらに大きな疑問は、なぜその決定が下されたのか本当に知っていますか？
今、市場ではトレンドの変化が起きていると思います。変化は、価値創造に対するトークンの支出を再確認することです。この変化は主に、AI を活用してすでに相当量のユースケースを出荷している企業の間で起こっています。いくつかのユースケースの導入に取り組んでいる陣営は、まだ実際に打撃を受けていないため、支出を気にしていません。しかし、予算が使い果たされ始めると、問題が発生するのはほぼ確実です。
あらゆることにAIを使用したいという本能は間違っていませんでした。しかし、それは常に意味があるのでしょうか？チーム

そして人々は、どのステップが実際に「インテリジェンス」を必要とするのか、どのステップが単にルールやロジックを必要とするのか、と疑問を持ち始めています。これは、トークンの支出だけでなく、レイテンシと一貫性についても答えを導き出します。一貫性とは、システムが何かを実行している理由を知っていることを意味します。
しかし、AI を使用していなければ、私たちは損をすることになるのではないでしょうか?それは今みんなが言うことではないですか？乗り込むか、それとも取り残されて、より現代的な競合企業に食い荒らされるか。
解決策は、AI を最大限に活用することですが、すべてにおいて盲目的に行うのではなく、最大限の価値を生み出す方法で行うことです。これを説明するには例が遅すぎると思います。
経費の承認 - これは非常に一般的であり、どの企業にも必要であり、通常は次の 2 つの方法で行われます。
人間が手動で各経費を確認し、公開されたポリシーに基づいて承認します。
一部の経費を自動的に承認し、他の経費を承認者にルーティングできる人事システムの一部のルール
企業の財務チームが、よりダイナミックになり、トレンドの変化に迅速に対応し、定められたルールに拘束されない経費管理を可能にすることで、従業員の生産性を最大化するために会社に利益をもたらすシステムを構築したいと考えているとします。
これを構築するよう依頼されたエンジニアリング チームは、これを簡単に行うことができます。財務チームにポリシーを Google ドキュメントか何かに書くように依頼し、それを公式ポリシーとして社内ポータルに公開し、このポリシーを読み取ってポリシーに基づいてすべてのリクエストを承認する AI エージェントを構築すると言うこともできます。財務チームはポリシーを随時更新できるようになり、開発者が関与することなく、各経費承認リクエストをポリシーに反映できるようになりました。クールですよね？
ユーザーが経費エージェントとのチャットを開始する
挨拶を交わす（もちろん、エージェントであるかどうかにかかわらず、私たち人間は常にそうします！）
上へ

領収書をロードし、費用を説明する
エージェントは領収書を解析し、金額と日付を検証します
リクエストに対して経費ポリシー全体を評価します
リクエストを決定し、ユーザーに通知します
承認された場合は、必要な経費の払い戻しを記録するよう人事システムに要求します。
これは、もし構築するとしたら非常に素晴らしいエージェント フローであり、おそらく財務チームと会社全体が喜んで使用すると思います。まず、ポリシーは（AI がインテリジェントであると仮定して）実際に適用され始め、財務チームは希望に応じて毎週月曜日にポリシーを変更できる柔軟性を備えています。 Win Win - CTO は、現在利用可能なモデルのインテリジェンスをどのように活用してビジネスに価値を付加したかを取締役会に提示できます。
ここでの大きな節約は、このような手段がなければ財務チームが費やさなければならない手動の承認時間を短縮できることです。あるいは、変更に応じてポリシーを変更し続けるために開発者に必要な時間と、最新の更新に従って経費が処理されない可能性がある間、そのためのリードタイムがさらに長くなります。
では、この価値はモデルプロバイダーから今後請求される可能性のある新しいトークン請求に見合う価値があるのでしょうか?
ここで何か違うことができるでしょうか?間違いなく使用する価値のあるモデルです。多くのシナリオで非常に効果的であることが疑いの余地なく証明されています。上記の例を変更すると、次のようになります。
ポリシーが更新されるたびに、次の手順が実行されます。
ポリシーから抽出された基本的なケースのルールのセットを作成します
人間が検証するためのテストシナリオを作成する
テストシナリオとルールセットを財務チームに送信して承認してもらいます
経費申請ごとに次の手順を実行します。
ユーザーにデータ入力フォームを提示する
ユーザーが非構造化フォーム経由で入力することを選択した場合は、モデルを実行して構造化フォームに変換します。
デプロイされたルールを実行する
一致する場合は、ルールに従って承認または拒否します
どれも一致しない場合

s
モデルを実行して承認を決定する
または、音量を下げるために人間にルーティングします
ユーザーに通知し、払い戻しのために人事システムを呼び出します。
まだそのモデルを使っています！ただし、重要な場合に限り、トークンのコストを 80 ～ 90% 削減できます。構造化入力を使用する人が増えれば増えるほど、モデルが入力を解析する必要さえなくなります。最終的にはモデルのインテリジェンスを得ることができますが、静的で決定的なルールが作成されるため、より一貫した方法でモデルが得られます。大量のリクエストは、財務チームによるポリシーの更新ごとに更新されたルールセットによって自動的に承認されるようになります。以前のフローと同じ柔軟性がありますが、トークンの使用量は大幅に削減されます。そして、私たちは AI を使用して、AI が優れた目的で設計された 1 つのこと、つまり、物事をコード化することを実行できるようになります。ループ内の開発者は、他のコード レビューと同様にコードをレビューし、エンジニアリング標準との整合性を確認することもできます。全体的に勝利、勝利、勝利。
いつモデルが必要なのか、いつ単純な論理的解釈が必要なのかを判断するのは非常に簡単です。ワークフローのすべてのステップで、次のことを決定します。
このステップでは、コンテキストの理解、言語の生成、または微妙な決定を行う必要があるのでしょうか、それともルールに従うだけでよいのでしょうか?
これは、建築から魅力の一部を取り除くため、有益な質問です。正直に聞いてみると、ほとんどの手順は最初に見たほど謎ではありません。
あなたがこれを読んでいるということで、私たちのプラットフォームである Unmeshed について少し共有したいと思います。 Unmeshed は、チームがモデル呼び出し、決定論的ルール、API 統合、人間による承認、可観測性を 1 か所で組み合わせた AI を活用したワークフローを構築するのに役立ちます。
Unmeshed では、エンジニアリング チームは次のことができます。
1 つのワークフローで API 呼び出し、ルール、デシジョン テーブル、人間の承認、AI ステップをモデル化します。
どのステップがモデルを呼び出すのか、そしてその理由を確認する
コストをスペックに帰属させる

具体的なワークフローと結果
エージェントの実行に予算、範囲、ツールの許可リストを設定する
高リスクまたは曖昧な決定について人間が常に把握できるようにする
時間の経過とともに、反復可能な決定をモデル呼び出しから決定論的ロジックに移行します。
デシジョン テーブルはルール エンジンであり、AI を使用するか手動で複雑なビジネス ルールを作成できます。ここでは、モデルを使用してルールを作成し、人間がルールを検証してから、ユースケースの実行ごとに個別のモデルを呼び出す場合と比較して、ほとんど無料で効果的にルールを実行できます。
チームが AI をビジネス プロセスに追加し、コスト、レイテンシ、結果がどこから来ているのかを考え始めている場合、Unmeshed はワークフローを 1 か所で設計、実行、観察、最適化する場所を提供します。
Unmeshed を使用すると、モデル呼び出し、決定論的ルール、デシジョンテーブル、および人間の承認を 1 つのワークフローに混在させることができるため、他のあらゆる場所で価値とロジックを追加する AI を使用できます。
AI ワークフローを導入すれば、機能を削減することなくトークン コストを削減できます。
トークンの効率とは何ですか? AI チームのための実践ガイド
トークン効率は、消費されたトークンの合計に対する有用な出力の比率です。このガイドでは、AI チームが最も多くの損失を被る場所と、それを回避する方法について詳しく説明します。
2026 年 6 月 24 日 再 13 分

[切り捨てられた]

## Original Extract

In this article, we explore how to bring AI into production while keeping token costs under control, ensuring the cost-benefit equation stays favorable and delivers real business value.

Bringing AI Workflows into Production without Burning Tokens - Unmeshed
Workflows Automate and run multi-step workflows API Orchestration Build and run high-scale API endpoints backed by workflows Batch Job Processing Run jobs at second-level granularity while managing dependencies between tasks Capabilities
Integrations Integrate with systems, built-in connections to over 100+ apps Agentic AI Leverage AI within your workflows Security Enterprise-grade security What's new
Human in the loop Build workflows that require human input, such as approvals Decision Engine Use our built-in rules engine to scale up logic in your workflows Hosted functions Write steps in Python, JS/TS, or Go and run them without separate hosting Solutions
How Unmeshed can help your automation needs across teams
Developers Run workflows without boilerplate or infra overhead RevOps Accelerate revenue workflows with orchestration Marketing Improve campaign ROI with connected flows IT Run reliable operations with automation-first workflows HR Automate onboarding, approvals, and HR operations Sales Close deals faster with automated routing Customer Support Resolve tickets faster with coordinated workflows By Use Case
Migrate from Conductor Migrate Conductor workflows to modern orchestration API Orchestration Fan out REST + GraphQL and return unified responses Migrate from Autosys Move legacy scheduler workloads with low disruption Leave Workflow Routing Route leave requests with approvals and loopbacks Fraud Detection Decisioning Fan out risk checks and return decisions in real time Loan Application Processing Route loan approvals with rules and human review Customer Onboarding and KYC Automate onboarding with checks and approvals Read our blogs on solutions
5 ways Orchestration Drives Efficiency and Innovation How orchestration improves speed, reliability, and delivery How Orchestration Transforms Fraud Detection Workflows Design low-latency fraud checks with orchestrated decisioning Optimizing Logistics Software with Unmeshed Connect supply-chain systems into resilient, observable flows Commodity Trading API Orchestration Orchestrate market data and trade APIs at scale Easy Migration from CA Autosys to Unmeshed Modernize legacy schedulers with workflow-first automation Resources
Learn, build, and operate effectively with Unmeshed
Documentation Start quickly with setup, guides, and API references Authentication and Authorization Configure SSO, access policies, and RBAC for secure orchestration Blogs Product updates, engineering notes, and solution deep dives Why Unmeshed Automations Platform differences, architecture advantages, and migration insights Unmeshed v/s Others Compare Unmeshed with common automation and orchestration tools Guides Step-by-step tutorials for common automation patterns Concepts Core ideas and architecture patterns behind Unmeshed Support
Help Center Contact Us Hire an Expert Professional Services Training Tools & Utilities
Unmeshed CLI Command line utilities for CI pipelines and workflow automation operations SDKs Language SDKs for Java, Golang, TypeScript/JavaScript, and Python Back to Blog Bringing AI Workflows into Production without Burning Tokens
Adopting AI (or the abilities of an LLM) into production is a core metric or goal for most engineers today. In this article we look at the best way to bring AI into production while keeping the token costs under control such that the cost vs benefit equation lands in the benefits bucket and adds value to the business.
10 min read June 22, 2026 Let’s make it Agentic!
The push in the market is to use agentic flows. Agentic is when you let a model decide how to process a request or a flow and expect its abilities to parse and understand context to result in the best possible outcome for the use case. The idea is that as models mature and become more “intelligent” the outcomes become more high quality and beats a human coded fixed algorithm.
With this in mind, oftentimes you’d see a use case pushed into production which relies on model calls completely.
For example, an agent may execute a use case by parsing input, validating data, classifying the request, checking policy, routing it to the right person, and drafting a response, all by calling a model.
It's often quite fast to build this, with the many agentic frameworks out there today and the demo is usually great and impressive to the management.
Launching this in production in a high volume use case will bring a shock though - when the bill arrives from the model providers. Token costs are increasing rather than decreasing as models evolve. Now is that use case adding sufficient value to the cost of running it?
What about questions such as consistency, latency, security and governance? And an even bigger question, do you really know why a decision was made?
I think there is a trend shift happening now in the market. The shift is to double check token spend to value creation. This shift is primarily among those who have already shipped a reasonable amount of use cases leveraging AI. The camp that is still working to deploy some use cases is not bothered by the spend yet as it hasn’t really hit them yet. But it's almost a certainty that once your budget starts to get eaten up, the question will come.
The instinct to use AI for everything was not wrong. But does it always make sense? Teams and people are starting to ask which of the steps actually needs “intelligence” and which ones just need some rules or logic? This leads to answers for not just token spend, but the latency and consistency as well. Consistency means you know the reasons why the system is doing something.
But if we are not using AI, are we not losing out? Isn’t that what everyone says now? Get onboard or be left behind to be eaten by more modern competing companies.
The solution is to maximize the use of AI, but in a way that it yields maximum value and not just blindly at everything. I think an example is overdue for explaining this.
Expense approvals - This is quite common and every company needs it and typically it's done by a couple methods:
A human manually reviews and approves each expense based on some published policy
Some rules in a HR system that can automatically approve some expenses and route others to approvers
Let’s say the finance team of a company wants to be more dynamic, rapidly respond to changing trends and create a system that can benefit the company to maximize employee productivity - by letting them manage expenses that are not bound by rules set in stone!
An engineering team asked to build this could simply do this - ask the finance team to write the policy in a Google Doc or something - which can then be published to the internal portal as the official policy and then say build an AI agent that reads this policy and approves every request based on the policy. Now the finance team can update the policy every now and then, and without any developer in the loop, the policy can reflect on each expense approval request - Et Voila! Cool right?
User initiate a chat with an expense agent
Exchange greetings (of course we humans always do that, agent or not!)
Upload a receipt, explain the expense
Agent parses the receipt, validates the amounts and dates
Evaluates the entire expense policy against the request
Decides on the request, informs the user
If approved, make a request to the HR system to note the required expense reimbursement
This is a pretty cool agentic flow if you were to build it and I think the finance team and the entire company is probably going to be thrilled to use it. First, policies will start being practically applied (assuming the AI is intelligent) and the finance team has the flexibility to change it every Monday if they want to. Win win - And the CTO can present to the board on how they leveraged the intelligence of the models available today to add value to the business.
The big savings here is the manual approval times that the finance team would have to spend on without something like this. Or even more is the developer time required to keep changing policies as they change and lead time for doing so while expenses may not be processed as per the latest update.
So is this value worth the new token bills that may now start coming from the model providers?
What could we do differently here? The models are worthy of use for sure. It's proven beyond doubt that it can be very effective in a lot of scenarios. A change up for the example above could be this:
On every policy update, steps:
Create a set of rules for basic cases extracted from the policy
Create test scenarios for the human to verify
Send the test scenarios and ruleset for finance team to approve
On every expense request, steps:
Present a data entry form for user
If user chooses to enter via an unstructured form, then run the model to convert it into structured
Run the rules deployed
if its matching, approve or deny as per rules
If none matches
Run the model to decide approval
Or route to human for the lower volume
Inform the user and invoke the HR system for reimbursement
Still using that model! But only where it matters and you can pretty much cut down the token costs by 80-90%. The more people who use structured input, there is not even a need for the model to parse the inputs. We still end up the model’s intelligence but in a more consistent way since it created static, deterministic rules. The large volume of requests will be now automatically approved by the updated rulesets for each policy update by the finance team. Same flexibility as the earlier flow, but with a lot lower token spend. And we get to use AI to do the one thing it's designed to do great - code things out! A developer in the loop can also review the code and ensure its consistent with the engineering standards just like any other code review. Win win win across the board.
It’s quite easy to decide when you need a model and when you need a simple logical interpretation. In every workflow step decide on this:
Does this step require understanding context, generating language, or making a nuanced decision, or does it just need to follow a rule?
It is a useful question because it removes some of the glamour from the architecture. Most steps, when you ask it honestly, are less mysterious than they first looked.
So since you are reading this - I’d love to share a bit about our platform - Unmeshed . Unmeshed helps teams build AI-powered workflows that combine model calls, deterministic rules, API integrations, human approvals, and observability in one place.
In Unmeshed, engineering teams can:
Model API calls, rules, decision tables , human approvals , and AI steps in one workflow
See which steps call models and why
Attribute cost to specific workflows and outcomes
Put budgets, scopes, and tool allow-lists around agentic execution
Keep humans in the loop for high-risk or ambiguous decisions
Move repeatable decisions from model calls into deterministic logic over time
The decision table is a rules engine and you can create complex business rules using AI or manually. This is where you could use models to create your rules, have humans validate them and then run them for effectively next to nothing compared to individual model calls for every use case execution.
If your team is adding AI into business processes and starting to ask where the cost, latency, and outcomes are coming from, Unmeshed gives you one place to design, run, observe, and optimize the workflow.
Unmeshed lets you mix model calls, deterministic rules, decision tables, and human approvals in one workflow - so you use AI where it adds value and logic everywhere else.
Bring your AI workflow and we can help you cut token costs without cutting capability.
What Is Token Efficiency? A Practical Guide for AI Teams
Token efficiency is the ratio of useful output to total tokens consumed. This guide breaks down where AI teams lose the most money and how to get ahead of it.
June 24, 2026 13 min re

[truncated]
