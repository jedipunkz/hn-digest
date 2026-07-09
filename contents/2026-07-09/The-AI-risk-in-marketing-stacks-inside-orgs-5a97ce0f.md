---
source: "https://singulr.ai/blog/the-hidden-ai-risk-in-your-marketing-stack-shadow-ai-automation-pipelines-and-the-hidden-data-leaks"
hn_url: "https://news.ycombinator.com/item?id=48843431"
title: "The AI risk in marketing stacks inside orgs"
article_title: "Hidden AI Risks in Your Marketing Stack: Shadow AI, Pipelines, and Data Leaks"
author: "shivc"
captured_at: "2026-07-09T10:46:46Z"
capture_tool: "hn-digest"
hn_id: 48843431
score: 1
comments: 0
posted_at: "2026-07-09T10:01:39Z"
tags:
  - hacker-news
  - translated
---

# The AI risk in marketing stacks inside orgs

- HN: [48843431](https://news.ycombinator.com/item?id=48843431)
- Source: [singulr.ai](https://singulr.ai/blog/the-hidden-ai-risk-in-your-marketing-stack-shadow-ai-automation-pipelines-and-the-hidden-data-leaks)
- Score: 1
- Comments: 0
- Posted: 2026-07-09T10:01:39Z

## Translation

タイトル: 組織内のマーケティング スタックにおける AI リスク
記事のタイトル: マーケティング スタックに隠れた AI リスク: シャドウ AI、パイプライン、データ漏洩
説明: マーケティング チームは、企業内で最も早く AI を導入しているチームの 1 つですが、最も管理が進んでいないチームの 1 つです。シャドウ AI から安全でない自動化パイプラインまで、機密データがどこですり抜けているか、そして実行時の強制が実際にどのように行われているかを発見します。

記事本文:
マーケティング スタックに隠れた AI リスク: シャドウ AI、パイプライン、データ漏洩
Agent Pulse™ の紹介: Agentic エンタープライズ向けのランタイム ガバナンス。
続きを読む 新しいホワイトペーパー: AI 時代のデータ リスクの再考 続きを読む 製品 Singulr Agent Pulse Singulr AI コントロール プレーン AI ガバナンス プラットフォーム ソリューション プライバシー、リスク、コンプライアンス チーム CISO およびセキュリティ チーム CIO および IT チーム リソース ブログ 資料 FAQ バイヤーズ ガイド
業界での評価
用語集
ホワイトペーパー パートナーについて デモを入手 → 製品 Singulr エージェント Pulse Singulr AI コントロール プレーン AI ガバナンス プラットフォーム ソリューション プライバシー、リスク、コンプライアンス チーム CISO およびセキュリティ チーム CIO および IT チーム リソース ブログ 資料に関するよくある質問 バイヤーズ ガイド
業界での評価
用語集
パートナーに関するホワイトペーパー デモを入手 →
ホーム / ブログ 2026 年 4 月 9 日 5 分で読む マーケティング スタックに隠れた AI リスク: シャドウ AI、自動化パイプライン、および隠れたデータ漏洩
現在、マーケティング チームは AI と自動化に取り組んでいます。成長マーケティング担当者は、CRM から意図の高いアカウントのリストをエクスポートし、それを ChatGPT に貼り付けて、パーソナライズされたアウトリーチの草案を作成する場合があります。キャンペーン リーダーは、パフォーマンス スプレッドシートをアップロードしたり、通話記録を AI ツールにアップロードして、何が機能しているのか、何が改善の必要があるのか​​を分析することがあります。
別のマーケティング担当者は、Clay または n8n オートメーションを構築して、リードをリアルタイムで強化してルーティングするかもしれません。これらのワークフローは効率的で、数時間の作業が数分に短縮されます。実際、一部の調査では、マーケターの約 88% がすでに日常の業務で AI ツールに依存していることが示されています。
しかし、このスピードとイノベーションには隠れたコストが伴います。機密の顧客データ、財務数値、または内部文書がサードパーティの AI または自動パイプラインにプッシュされるたびに、組織は攻撃対象領域を拡大します。 contr に存在するデータ

CRM や社内データベースなどのすべてのシステムは、IT 部門の完全な制御が及ばないネットワーク間を移動するようになりました。この効率性により、まったく新しい AI セキュリティとガバナンスのリスクが生じます。
マーケティングの新たな役割と高まるリスク
マーケティングはかつてのものではありません。もはやクリエイティブなキャンペーンやブランド メッセージだけではありません。今日のマーケティング部門は、顧客の PII (名前、電子メール、企業情報) から行動分析 (使用パターン、意図シグナル) に至るまで、さらにはヘルスケアや PHI などの業界でも、大量の機密データを処理します。そのデータは現在、AI システムとカスタム ワークフローを通じて大規模に処理されています。
一方、マーケティング チームは迅速な実験で成功します。低摩擦の新しいツール (ChatGPT、Claude、Notion AI、Zapier、n8n など) が採用されています。 「この AI ツールをすぐに試してみよう」というシンプルな考え方が、迅速な結果をもたらします。しかしそれは同時に、貴重なデータがこれまで以上に速く、より多くの人の手を介して移動することを意味します。
その結果、マーケティングは多くの組織において最もデータにさらされる部門の 1 つになりました。チームは「迅速に行動して学習する」ことに強い意欲を持っており、プロセスをスピードと引き換えにすることがよくあります。
人工知能インデックス レポート 2025 によると、マーケティングと営業が AI から最も恩恵を受けており、71% がコスト削減または収益の増加を報告しています。
AI はマーケティング業務に深く組み込まれており、それは当然のことです。しかし、統合や新しいワークフローが増えるたびに、意図しないデータ漏洩の可能性が高まります。
セキュリティ専門家は、AIの急速な導入とガバナンスの間にあるこのギャップは危険な盲点であると警告しています。
たとえば、IBM の 2025 年の侵害レポートでは、AI 関連の侵害を経験した企業の 97% が適切なアクセス制御を導入していなかったことがわかりました。マーケティング ワークフローは、その性質上、正式なポリシーの作成を上回ることがよくあります。
多くのチームは顧客リストをコピーアンドペーストします。

IT 部門の誰もが知る前に、公開ツールにアクセスしたり、新しい AI 機能を実験したりできます。
AI がマーケティングでデータを公開する方法
これを具体的にするために、AI 主導のマーケティングがデータ漏洩につながる具体的な方法をいくつか紹介します。
データの直接公開: マーケティング担当者が分析のためにスプレッドシート、顧客リスト、または CRM 抽出を AI ツールにアップロードするのは一般的です。保護策がなければ、これらの入力内の機密情報が漏洩する可能性があります。
実際には、営業電話の記録やキャンペーン結果を ChatGPT で要約するなど、一見無害に見える行為でも、機密コンテンツがクラウド モデルに送信される可能性があります。そのデータは保存される可能性があり、企業ポリシーが設定されていない場合や、「デフォルトでトレーニングがオン」のままの個人用 ChatGPT アカウントが使用されている場合には、将来のモデル出力をトレーニングするために使用することもできます。
規制された業界 (ヘルスケアや金融サービスなど) では、PHI や財務の詳細が 1 回でも公開されるだけで、重大なコンプライアンス違反につながる可能性があります。
シャドウ AI の使用: 最大のリスク倍増は、IT 部門の視野外で使用されている膨大な数の AI ツールです。この「シャドウ AI」は、個人アカウント、新しいアプリ、または独自に構築したワークフローを通じて、チームがセキュリティや IT に通知せずにツールを導入したときに発生します。 IBM の報告によると、組織の約 5 分の 1 が、こうした未承認の AI 活動による侵害をすでに経験しています。さらに悪いことに、シャドウ AI に関連するインシデントでは、より多くの PII が漏洩する傾向があります。IBM は、シャドウ AI 関連の侵害の 65% が顧客データを漏洩したのに対し、平均的な侵害では 53% であることを発見しました。
実際には、シャドウ AI は、暗号化されていない Wi-Fi で個人の AI アカウントを使用するマーケターから、自分でセットアップしたローコード プラットフォームでワークフローを構築するチーム全体まで、あらゆるものを意味します。企業がそれが起こっていることを認識していなければ、それを確保することはできません。
安全でない自動化PI

pelines: 現代のマーケティングでは、ツールを接続してワークフローを自動化するために、統合プラットフォーム (Zapier、n8n、Clay など) に依存することがよくあります。これらは効率を高める強力なツールですが、安全でない方法でデータを移動する可能性があります。たとえば、Zapier 統合のスコープが不十分な場合、CRM からデータを取得し、広すぎる権限を使用して複数のステップを介してデータを送信する可能性があります。
エンドポイントがロックダウンされていない場合、または API キーの範囲が厳密に定められていない場合、攻撃者 (または誤って設定されたプロセス) によって、意図したよりも多くのデータが抜き取られる可能性があります。
たとえば、複数の API を呼び出すリード エンリッチメント パイプライン: これらのエンドポイントのいずれかが適切な認証なしでインターネットに公開されると、ハッカーが悪用する可能性があります。統合が追加されるたびに、チェーン内に保護する必要がある新しいリンクが作成されます。
自動化は生産性の向上に優れていますが、最初からセキュリティを考慮しない限り、データが気づかれずにすり抜けてしまう経路も増えてしまいます。
AI エージェントとナレッジ ベース: 企業がより高度な AI 機能を構築するにつれて、マーケティング担当者は内部システムに直接接続する AI 「エージェント」を導入する可能性があります。
たとえば、Notion ナレッジ ベースから読み取る LLM、営業電話の録音に関連付けられたカスタム チャットボット、さらには CRM やコラボレーション ツールに埋め込まれた LLM などです。これらのエージェントに広範なアクセスが与えられている場合、意図せずにプッシュすべきではないデータをプッシュしてしまう可能性があります。
たとえば、すべての顧客サービス チケットから取得することが許可されている AI アシスタントは、間違った方法で指示された場合、個人情報を漏洩する可能性があります。 AI エージェントのアクセスは慎重かつ厳密に範囲を限定する必要があり、厳密な読み取りおよび書き込み権限を持つ特定のドキュメントまたはデータ セグメントのみが開かれる必要があります。それがなければ、自律エージェントがいつでもデータを公開する可能性があります。
迅速なインジェクションとデータ漏洩: ついに AI による新しい攻撃テクノロジーが導入されました

ニケス。即時注射も脅威です。
この好例は、チットポトレのウェブサイトのチャットボットを人々がコーディングやその他の作業を手伝うために使用した最近の事件です。チャットボットが与えるべき応答の種類についてのガードレールが設置されていなかったため、これは技術的にはプロンプト インジェクションではありませんが、これは、誰でもプロンプトをいじることができ、十分にうまくやれば、これらの公開 AI ツールを取得して、業務上共有していないデータを共有できるという私のポイントを強調しているだけです。
これらの問題のほとんどは、悪意なく発生します。ほとんどのチームは、単純に反復処理を高速化したいと考えています。彼らはデータを漏洩しようとしているわけではありません。彼らはガードレールを設置するのに時間がかかっていないだけです。
AIは敵ではない、加速するものである
明確にしておきたいのは、解決策はマーケティング チームでの AI 使用にブレーキをかけることではないということです。その船は出航しました。マーケティング担当者は AI を使用する必要があります。それは彼らの働き方を根本的に変えます。
AI 主導のワークフローは、より迅速な市場調査、より深い顧客セグメント化、大規模なパーソナライズされたメッセージングなど、ビジネスに大きなメリットをもたらします。
AI は、マーケティング チームがより賢く、生産性を高めるのに役立ちます。これにより、迅速なアイデア生成が可能になり、反復的なタスクが自動化され、手動分析では何時間もかかるような洞察が得られます。営業チームはより適切な準備を整え、キャンペーンのターゲティングはより厳しくなり、マーケティングはほぼリアルタイムで市場の変化に対応できるようになります。これらは真の競争上の利点です。
重要なのは、AI が機能するということです。問題はテクノロジーとしての AI ではなく、その導入方法が管理されていないことにあります。目標は、マーケティングを AI で安全に迅速に進めることです。言い換えれば、目的は制御された加速です。
安全な AI ワークフローの構築: マーケティング担当者のためのガードレール
責任ある AI が使用するものはどのようなものですか

練習のように？強力なツールを扱うのと同じように、ガードレール、監視、ベスト プラクティスを使用します。
安全で機敏なマーケティング AI プログラムの重要な要素は次のとおりです。
エンタープライズ グレードの AI アカウント: すべての AI 作業は、個人アカウントではなく、会社が管理するツールを通じて実行する必要があります。たとえば、マーケティング担当者は、個人ログインではなく、SSO ログインを使用して ChatGPT のエンタープライズ エディションを使用する必要があります。
これにより、データ処理ポリシーを一元的に適用できるようになります。承認されたエンタープライズ ツールは、入力をトレーニングしないように構成でき、組織に関連付けられたログを生成します。対照的に、個人アカウントには監視や制御がありません。許可されていないアクセスを排除します。仕事のための個人的な AI ログインはありません。
データ規律: すべてのデータが等しいわけではありません。 AI ツールに何を入力できるか、何ができないかを明確に定義します。個人を特定できる情報、機密の財務数値、または健康データは、明示的に暗号化または匿名化しない限り、立ち入りを禁止する必要があります。 Singulr などのツールを使用して、そのようなプロンプトを自動的に検出し、編集/ブロックします。たとえば、AI へのプロンプトは、ネットワークを離れる前に顧客名やクレジット カード番号を自動的にスクラブする必要があります。システムが違反のフラグを立てた場合、リクエストをブロックするか、少なくともセキュリティに警告する必要があります。実際には、これはマーケターが AI UI に入力した内容と、自動化パイプラインを通じて送信された内容をスキャンすることを意味します。 Singulr のようなソリューションでは、実行時により詳細なポリシーを適用することもできます。これらの制御により、マーケティング担当者が誤って危険なことを試みたとしても、ツールによってそれが阻止されます。
ID とアクセス制御: 最小特権の原則をあらゆる場所に適用します。マーケティング担当者と彼らが導入するボットは、必要なデータとシステムにのみアクセスできるようにする必要があります。人間のユーザーの場合、これは次のような標準的な IAM プラクティスを意味します。

ロールベースのアクセス許可、シングル サインオン、MFA、定期的なアクセス レビューなど。しかし、これは AI の統合にも当てはまります。 AI エージェントに CRM への読み取りアクセスを許可する場合は、CRM を読み取り専用にします。 「書き込みアクション」を許可する場合は、まずそれをサンドボックス環境に置きます。たとえば、AI 主導の電子メールや広告キャンペーンは、送信前に人間の承認が必要であることを義務付けます。
安全な自動化設計: マーケティング オートメーションを他のコードやシステムと同様に扱います。すべてのエンドポイントを検証し、強力な認証 (OAuth スコープ、署名された Webhook など) を使用し、絶対に必要な場合を除き、生の機密データを決して渡しません。たとえば、n8n ワークフローが顧客名を取得する場合、完全なレコードを送信する代わりにデータをハッシュまたは切り詰めることができるかどうかを検討してください。内部データベースとインターネットの間の橋渡しをする自動化に常に注目してください。データ フローを監査してログに記録します。防御的に自動化を構築します。各ステップを構築するとき、チームは「これは露出が多すぎませんか?」と尋ねる必要があります。
AI エージェントのガバナンス: 多くのマーケティング チームは定期的に AI エージェントを実験しています。これらのエージェントには設計によりガバナンスが備わっている必要があります。これは、クエリできる内部システムを明示的に制限することを意味します。たとえば、エージェントが電子メール キャンペーン ツールや CRM に接続している場合、ポリシーでどのデータ フィールドを読み取ることができるかを正確に定義します。

[切り捨てられた]

## Original Extract

Marketing teams are among the fastest AI adopters in the enterprise and among the least governed. From shadow AI to unsecured automation pipelines, discover where sensitive data is slipping through and what runtime enforcement looks like in practice.

Hidden AI Risks in Your Marketing Stack: Shadow AI, Pipelines, and Data Leaks
Introducing Agent Pulse™ : Runtime Governance for the Agentic Enterprise.
Read More New whitepaper: Rethinking Data Risk for the AI Era Read More Product Singulr Agent Pulse Singulr AI Control Plane AI Governance Platform Solutions Privacy, Risk, and Compliance Teams CISO and Security Teams CIO and IT Teams Resources Blogs Collateral FAQs Buyer’s Guide
Industry Recognition
Glossary
Whitepapers About Partner Get Demo → Product Singulr Agent Pulse Singulr AI Control Plane AI Governance Platform Solutions Privacy, Risk, and Compliance Teams CISO and Security Teams CIO and IT Teams Resources Blogs Collateral FAQs Buyer’s Guide
Industry Recognition
Glossary
Whitepapers About Partner Get Demo →
Home / Blog April 9, 2026 5 Min Read The Hidden AI Risk in Your Marketing Stack: Shadow AI, Automation Pipelines, and the Hidden Data Leaks
Marketing teams today are piling on AI and automation. A growth marketer might export a list of high-intent accounts from the CRM and paste it into ChatGPT to draft personalized outreach. A campaign lead might upload performance spreadsheets or call transcripts into an AI tool to analyze what’s working and what needs improvement.
Another marketer might build a Clay or n8n automation to enrich and route leads in real time. These workflows are efficient, turning hours of work into minutes. In fact, some surveys show roughly 88% of marketers already rely on AI tools in their day-to-day roles.
But all of this speed and innovation comes with a hidden cost. Every time sensitive customer data, financial numbers, or internal docs get pushed into a third-party AI or automated pipeline, the organization expands its attack surface . Data that lived in controlled systems like CRMs and internal databases now moves across networks outside IT’s full control. This efficiency introduces a brand-new AI security and governance risk.
Marketing’s New Role and Rising Risk
Marketing isn’t what it used to be. It’s no longer just creative campaigns and brand messaging. Today’s marketing function handles massive amounts of sensitive data , from customer PII (names, emails, firmographics) to behavioral analytics (usage patterns, intent signals), and in industries like healthcare, even PHI. That data is now being processed through AI systems and custom workflows at scale .
Meanwhile, marketing teams thrive on rapid experimentation. New tools (ChatGPT, Claude, Notion AI, Zapier, n8n, etc.) are adopted with low friction. A simple “Let me try this AI tool quickly” mindset drives fast results. But it also means valuable data moves faster and through more hands than ever.
The result: marketing has become one of the most data-exposed functions in many organizations. Teams have a high appetite for “move fast and learn” and often trade processes for speed.
The Artificial Intelligence Index Report 2025 shows that marketing and sales benefit most from AI, with 71% reporting cost reduction or revenue gains.
AI is deeply embedded in marketing’s operations and rightly so. But with every integration and every new workflow, the chance of unintended data leaks grows.
Security experts warn that this gap between rapid AI adoption and governance is a dangerous blind spot.
For example, IBM’s 2025 breach report found that 97% of companies experiencing an AI-related breach had no proper access controls in place. Marketing workflows, by their very nature, often outpace the creation of formal policies.
Many teams copy-paste customer lists into public tools or experiment with new AI features before anyone on IT even knows.
How AI Can Expose Data in Marketing
To make this concrete, here are some specific ways AI-driven marketing can lead to data exposure:
Direct data exposure: It’s common for marketers to upload spreadsheets, customer lists, or CRM extracts into AI tools for analysis. Without safeguards, any sensitive information in those inputs can leak.
In practice, even a seemingly harmless act like summarizing sales call transcripts or campaign results in ChatGPT can result in confidential content being sent to a cloud model. That data could be stored, or even used to train future model outputs if enterprise policies aren’t in place or if a personal ChatGPT account was used, which keeps “training ON by default”.
In regulated industries (such as healthcare or financial services), even a single exposure of PHI or financial details can result in severe compliance violations.
Shadow AI usage: The biggest risk multiplier is the sheer number of AI tools being used outside of IT’s view. This “shadow AI” happens when teams adopt tools without informing security or IT - whether through personal accounts, new apps, or self-built workflows. IBM reports that about one in five organizations has already experienced a breach coming from these unsanctioned AI activities. Worse, incidents involving shadow AI tend to leak more PII: IBM found 65% of shadow-AI-related breaches exposed customer data, vs. 53% in an average breach.
In practice, shadow AI could mean anything from a marketer using a personal AI account on unencrypted Wi-Fi to an entire team building workflows on a low-code platform they set up themselves. If the company doesn’t know it’s happening, it can’t secure it.
Unsecured automation pipelines: Modern marketing often relies on integration platforms (Zapier, n8n, Clay, etc.) to connect tools and automate workflows. These are powerful efficiency boosters, but they can move data in unsafe ways. For example, a poorly scoped Zapier integration might pull data from the CRM and send it through multiple steps with overly broad permissions.
If endpoints aren’t locked down or API keys aren’t tightly scoped, an attacker (or a misconfigured process) could exfiltrate more data than intended.
For example, lead enrichment pipelines that call multiple APIs: if any of those endpoints are exposed to the internet without proper auth, a hacker could exploit them. Each additional integration creates a new link in the chain that needs to be secured.
Automation is great for productivity, but it also creates more pathways for data to slip through unnoticed unless security is factored in right from the start.
AI agents and knowledge bases: As enterprises build more advanced AI capabilities, marketers might deploy AI “agents” that connect directly to internal systems.
For example, an LLM reading from a Notion knowledge base, a custom chatbot tied to sales call recordings, or even LLMs embedded in CRM or collaboration tools. If these agents are given broad access, they can unintentionally push data they shouldn’t.
For example, an AI assistant that’s allowed to pull from all customer service tickets might leak private details if prompted the wrong way. Any AI agent’s access should be carefully and strictly scoped, and only specific documents or data segments, with strict read and write permissions, should be opened up. Without that, an autonomous agent could expose your data at any time.
Prompt injection and data exfiltration: Finally, AI introduces new attack techniques. Prompt injection is a menace as well.
A great example of this is the recent incident in which people used Chitpotle’s website chatbot to help them with coding and anything in between. Although that isn’t technically prompt injection since there were no guardrails in place around the type of responses that the chatbot should give, it simply highlights my point that anyone can play around with the prompts, and if they do so well enough, they can get these public-facing AI tools to share data they have no business sharing!
Most of these issues arise without any malicious intent. Most teams simply want to iterate faster. They’re not trying to leak data. They just haven’t taken the time to put guardrails in place.
AI Is Not the Enemy, It’s the Accelerator
It’s important to be clear: the solution isn’t to slam the brakes on AI use in marketing teams. That ship has sailed. Marketers should be using AI; it fundamentally transforms how they work.
AI-driven workflows offer significant business benefits such as faster market research, deeper customer segmentation, and personalized messaging at scale.
AI can help marketing teams be smarter and more productive. It enables fast idea generation, automates repetitive tasks, and surfaces insights that would take hours of manual analysis. Sales teams get better prep, campaigns get tighter targeting, and marketing can react to market shifts in near-real time. Those are real competitive advantages.
The point is that AI works . The problem isn’t AI as a technology but the uncontrolled way it’s usually adopted. The goal should be to let marketing move fast with AI safely . In other words, the objective is controlled acceleration .
Building Safe AI Workflows: Guardrails for Marketers
What does responsible AI use look like in practice? The same way we treat any powerful tool, with guardrails, oversight, and best practices.
Here are the key elements of a secure and agile marketing AI program:
Enterprise-grade AI accounts: Require all AI work to happen through company-controlled tools, not personal accounts. For example, marketers should use the enterprise edition of ChatGPT with SSO login, rather than a personal login.
This ensures data handling policies can be enforced centrally. The approved enterprise tools can be configured not to train on your inputs, and they produce logs tied to your organization. By contrast, a personal account offers no oversight or control. Eliminate unsanctioned access: no personal AI logins for work .
Data discipline: Not all data is equal. Clearly define what can and cannot be fed into AI tools. Personally Identifiable Information, sensitive financial figures, or health data should be off-limits unless explicitly encrypted or anonymized. Use tools like Singulr to automatically detect and redact/block such prompts. For example, a prompt to an AI should automatically scrub customer names or credit card numbers before ever leaving your network. If your system flags a violation, it should either block the request or at least alert security. In practice, this means scanning what marketers input into AI UIs and what’s sent through automation pipelines. Solutions like Singulr can also enforce more fine-grained policies at runtime. These controls ensure that even if a marketer accidentally attempts something risky, the tool prevents it.
Identity and access control: Apply the principle of least privilege everywhere. Marketers and the bots they deploy should have access only to the data and systems they need. For human users, it means standard IAM practices such as role-based permissions, single sign-on, MFA, and regular access reviews. But it also applies to AI integrations. If you grant an AI agent read access to the CRM, make it read-only . If you do allow ‘write actions’, put it in a sandbox environment first. For example, require that AI-driven emails or ad campaigns must have human approval before being sent.
Secure automation design: Treat marketing automations like any other code or system. Validate all endpoints, use strong authentication (OAuth scopes, signed webhooks, etc.), and never pass raw sensitive data unless absolutely needed. For example, if an n8n workflow fetches customer names, consider whether it can hash or truncate the data instead of sending full records. Keep an eye on any automation that bridges between an internal database and the internet. Audit and log the data flows. Build your automations defensively. When building each step, the team should ask, “Is this exposing too much?”
Governance for AI agents: Many marketing teams regularly experiment with AI agents. These agents should come with governance by design. That means explicitly limiting what internal systems they can query. For example, if an agent is connected to your email campaign tool and CRM, define in policy exactly which data fields it may read or

[truncated]
