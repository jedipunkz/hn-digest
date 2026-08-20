---
source: "https://www.qawolf.com/blog/code-factories-without-quality"
hn_url: "https://news.ycombinator.com/item?id=49375253"
title: "Code Factories Without Quality: The AI Development Blind Spot"
article_title: "Code Factories Without Quality: The AI Development Blind Spot | QA Wolf"
image: "https://cdn.prod.website-files.com/6260298eca091b3621cf1890/6a86a0b515ce9b0d9bdb6060_og_blog_post%20(2).png"
author: "BinRoo"
captured_at: "2026-08-20T15:24:43Z"
capture_tool: "hn-digest"
hn_id: 49375253
score: 2
comments: 0
posted_at: "2026-08-20T14:39:10Z"
tags:
  - hacker-news
  - translated
---

# Code Factories Without Quality: The AI Development Blind Spot

- HN: [49375253](https://news.ycombinator.com/item?id=49375253)
- Source: [www.qawolf.com](https://www.qawolf.com/blog/code-factories-without-quality)
- Score: 2
- Comments: 0
- Posted: 2026-08-20T14:39:10Z

## Translation

タイトル: 品質のないコードファクトリー: AI 開発の盲点
記事のタイトル: 品質のないコード ファクトリ: AI 開発の盲点 | QAウルフ
説明: AI エージェントは、チームがテストするよりも速くコードを生成します。コードファクトリーが本番環境で中断される理由と、自律的な E2E テストがどのようにギャップを埋めるのか。

記事本文:
品質のないコード ファクトリ: AI 開発の盲点 | QAウルフ
プラットフォーム マッピング AI 数分でアプリ全体のアウトラインを自動的に作成 自動化 AI プロンプトを決定論的なものに変える Playwright & Appium インフラの実行 Web テストとモバイル テストを並行して実行 サービス顧客 QA Wolf が Salesloft を年間 75 万ドル以上節約 メトロノームは毎日 4 回のリリースを欠かさず実行 QA Wolf が Meow Wolf のモバイル テストに魔法をかける
約 6 ワードの長さの Quippy の見出し 長さ約 6 ワードの Quippy の見出し 約 6 ワードの長さ その他の顧客事例 リソース 変更ログ ソリューション ドキュメント ブログ ウェビナー ガイド コミュニティ 価格 サービス利用規約 プライバシー ポリシー 採用情報 お問い合わせ QA Wolf は、ソフトウェア チームが QA を完全に担当することで、より優れたソフトウェアをより迅速に出荷できるように支援するハイブリッド プラットフォームおよびサービスです。
品質のないコード ファクトリ: AI 開発の盲点
重要なポイント - コードに AI を掛け合わせます。検証は行われませんでした。そのギャップこそが事件の生きどころなのです。
- 報道は検証ではありません。回線カバレッジが 100% であっても、壊れたユーザー フローが送信されます。 E2E は、カバレッジが見逃しているものをキャッチします。
- 検証はパイプラインでなければなりません。コードのようにスケールし、デプロイごとに実行し、独自のフレークを修正する必要があります。
ザピエル。ヌーバンク。ゴールドマン・サックス。優れたエンジニアリング チームとリスク管理チームを擁する 3 社が現在、人間の介入を最小限に抑えながらコーディング タスクを AI エージェントに渡しています。
コード ファクトリのワークフローは一見単純です。開発者がタスクを引き継ぎ、エージェントが計画、コード化、レビュー、テスト、デバッグを行い、エンドツーエンドで出荷します。
Zapier は 800 以上の AI エージェントを社内に導入し、組織全体で 89% の AI を導入しました。 Nubank は、コア ETL を移行するという、社内で最も重要なプロジェクトの 1 つに自律型 AI エージェントを導入しました。一方、ゴールドマン・サックスは円周率です

自律的なソフトウェア エンジニアを独自のコードベースに配置します。
コード ファクトリの魅力は明らかです。それは速度です。大規模なコード生成。 10倍速以上。プロジェクト全体が人間の手がキーボードに触れることなく出荷されました。
しかし、どのコード ファクトリも、同じステップであるテストをひっそりと省略したり、投資を怠ったりしています。
エージェントは、誰が検証するよりも早く実稼働コードを生成しています。発電能力が10倍に増加。検証能力はありませんでした。そのため、コードはテストされず、または十分にテストされていない状態で出荷されるため、コードファクトリーは速度だけでなく、あらゆるコストをかけて速度を重視し、品質も重視されます。
そして、本番環境で問題が発生した場合 (実際に問題が発生するでしょう)、次世代をより良くするためのフィードバック ループは存在しません。システム全体は、より多くのスロップをより速く出荷できるように設計されています。
暗黙の前提: AI コードは本番環境に対応している
すべてのコード ファクトリは、誰も大声で言わない前提に基づいて構築されています。つまり、生成されたコードは本番環境で使用できるコードであるということです。
そうではありません。しかし、QA が方程式の中で顕著に表れることはほとんどないため、この仮定は残ります。
代わりに、QA は下流のチェックボックス、または最小限の範囲で解決できるものとして扱われます。コード ファクトリでは、速度が中心的な価値であり、QA は摩擦であり、より迅速な出荷を妨げるものです。
理由は複雑ですが、遅いと見なされているため、優先順位が低くされたり、合格しても実際にはあまり検証されない浅いテスト範囲で自動化されてしまいます。
そして、チームはこのことを知っていても、他に選択肢がないと考えて、壊れる可能性が高いとわかっているコードを出荷することがよくあります。速いスロップは、わずかに遅い品質よりも優れていると見なされています。
本当の問題は、検証の速度がまだ生成の速度と一致していないことです。 AI コード レビュー ツールは検証の一面を支援します。彼らはいくつかのバグを発見し、改善を提案し、フラグを立てています

g 明らかな問題。しかし、彼らは依然としてバグを通過させています。
今必要なのはテスト速度の向上です。
コードファクトリーでのアンダーテストのコスト
コード ファクトリは、80% のカバレッジで出荷し、それを成功と呼ぶかもしれません。しかし、そのようなカバレッジは通常、ユーザー フローが実際に機能するかどうかではなく、実行された行を測定します。ライン カバレッジを 100% にしても、負荷時にエラーが発生したり、エラーが間違って処理されたり、データが不整合な状態のままになったりするコードを出荷することができます。
コード ファクトリは、不正なコードが本番環境に現れる速度を速めます。
そしてその結果は残酷なものだ。コード生成の規模が拡大するにつれて、発生率は上昇しています。 AI 生成コードを出荷しているチームでは、変更の失敗率が 30% 増加し、プル リクエストごとのインシデントが 23.5% 増加しています。
適切な検証を行わずにより多くのコードをより速く出荷すると、本番環境でのバグが増加し、顧客に直面する障害が増加し、消火活動が増加することを意味します。
PR 段階で見つかったバグの修正には数分かかりました。ステージング中にバグが見つかると、数時間のコストがかかります。本番環境でバグが見つかると、数日のコストがかかり、顧客の怒り、インシデント対応、依存システム全体にわたる連鎖的な障害が発生する可能性があります。導入後の修正にかかるコストは、出荷前に修正を検出するコストに比べれば微々たるものです。
適切なテストにより、出荷の遅れを生じさせずに迅速に出荷できます
本当の問題は、コード ファクトリが検証速度ではなく生成速度に最適化されていることです。 QA は、ベロシティを実現するものと見なされるべきにもかかわらず、ベロシティを阻害するものとして位置づけられることがあります。なぜなら、適切な自動テストカバレッジが整備されていれば、コードファクトリーは実際に信頼して迅速に出荷できるものになるからです。
問題は、人間によるテストの作成は直線的であるのに対し、エージェントの生成は直線的ではないため、10 倍検証を行うことができないことです。また、コードを継承するテストが行​​われるため、エージェントに自分の宿題を採点させることもできません。

盲点は誤った自信を自動化するだけです。
これらの障害を実際に検出するテスト、つまり実際のユーザー フローをエンドツーエンドでカバーするテストは、歴史的に構築が最も遅く、維持が最も困難です。それが罠です。コード ファクトリが実際に必要とする検証は、生成速度が 3 つすべての問題をさらに悪化させるまさにその瞬間に、構築するのが最も難しく、実行が最も遅く、維持するのに最もコストがかかる正確な検証です。
検証も自律的なパイプラインになる必要があります。全員の速度を低下させるものとして下流側に取り付けられるのではなく、同じ速度で並行して実行される必要があります。
しかし、それをうまくやるには、いくつかの譲れない要素が必要です。
従業員数ではなく、世代ごとに規模を拡大する必要があります。テストの作成は、QA エンジニアを何人雇用できるかによって決まります。それがペースを維持できる唯一のレートであるため、コード出荷のレートで拡張する必要があります。
ユーザーが実際に製品をどのように使用するかをテストする必要があります。行数ではなく、実際のフロー全体にわたるエンドツーエンドのカバレッジ。コードが実行されるだけでなく、顧客フローが適切に機能することを確認する必要があります。
コーディングエージェントから独立
これは、それを生成したコーディング エージェントから独立している必要があります。生成エージェントとは別に構築された検証は、コードの盲点を継承しません。エコーではなく、実質的なセカンドオピニオンです。しかし、それだけではなく、コーディング エージェントは、優れたテストの作成方法を理解するように構築されているわけではありません。ただし、私たちを信用しないでください。過去数か月間、コーディング エージェントを使用してテスト カバレッジを自動化するのに苦労している人々の Reddit の数十のスレッドを信頼してください。
それはそれ自身を維持しなければなりません。フローが変化するとテストが自動的に更新されるため、フレークはゼロ付近に留まり、チームはシグナルを信頼し続けます。自己維持型の保険は、接触しても生き残る唯一の種類です。

速度0倍。
これらをまとめると、QA はコード ファクトリー内での摩擦ではなくなります。発送滞りなく早く発送させて頂く物となります。
検証も最終的には生成のスピードで進めなければなりません。
QA Wolf がコード ファクトリ QA を解決する方法
これは記事の中で少し突っ込んだ部分ですが、実際にはほとんどのチームが直面している問題の解決策であるため、ご容赦ください。
私たちのプラットフォームは、チェックしているコードから独立して、生成速度で実行される自律的な QA パイプラインを作成します。アプリをマップし、テストを作成し、デプロイごとに実行し、コードが変動してもアプリを存続させます。これは、テストのライフサイクル全体を対象としたエンドツーエンドのプラットフォームです。
マッピング エージェントは、コーディング エージェントの想定ではなく、アプリケーションを独自にナビゲートし、その機能と実際のユーザー ワークフローの詳細なマップを構築します。ユーザーの役割を切り替え、Web、iOS、Android を切り替えて、複数のユーザーと複数のプラットフォームにまたがるフローを捕捉します。これは、AI によって生成されたコードが静かに中断される正確なマルチステップ、クロスステートの動作です。エージェントは人間が単独で行うよりも 32 倍高速です。
実際の決定論的なテストを作成します
オートメーション エージェントは、マップされたワークフローを決定論的な Playwright (Web) および Appium (モバイル) コードに変換します。これらは、API、シード データベースを呼び出し、外部依存関係を模擬し、機能フラグを反転して実際の条件下でアプリを実行するテストです。テストは実行のたびにスクリプトを再解釈する LLM ではなく決定論的なコードであるため、失敗は実際に何かが壊れたことを意味し、その失敗は再現可能です。また、スイートは製品コードを生成したものとは独立して生成されるため、エコーではなく、本物のセカンドオピニオンとなります。自動化エージェントは QA エンジニアの能力を向上させます 10

-20 倍。これは、テスト作成が最終的に、雇用可能なレートではなく、出荷されたレート コードでスケールされる方法です。
デプロイごとに完全に並行して実行されます
実行に何時間もかかるカバレッジは、コード ファクトリにとって役に立ちません。 QA Wolf は完全なスイートを 100% 並列実行で実行し、展開時に即座に開始されます。 CI パイプラインに直接接続できるため、常にマージされるパイプラインに合わせて検証を行うことができます。
スイート自体が維持されるため、スイートは 10 倍の速度に耐えることができます。
これは、ほとんどのテスト作業を無駄にする部分です。自律エージェントは一度だけコードを生成するわけではありません。彼らはそれをリファクタリングし、最適化し、書き直すのですが、それらの変更はすべてスイートを破壊する恐れがあります。 QA Wolf の Automation AI は、タイミングの問題、実行時エラー、レンダリングされていないコンポーネントなどのフレークのほぼ 100% に対処しますが、セレクター修復ツールが検出するのは約 20% です。テストが失敗すると、テストは失敗を再現し、原因を診断し、コードを書き直し、修正を検証します。一方、マッピング エージェントは新しい機能を監視し、個別の 1 回限りのテストや既存のテストを複製するのではなく、既存のフローにそれらを組み込みます。
QA Wolf のライフサイクルをまとめると、コード ファクトリに欠けていたものがわかります。つまり、生成速度でマップ、書き込み、実行、維持が行われることを検証するため、開発者は出荷を継続し、QA がボトルネックにならなくなります。
そうすることで、配送の遅れを生じさせずに迅速に配送することができます。
実際に稼働する工場の構築
QA がコード ファクトリのボトルネックになることはできません。それは基礎的なものでなければなりません。機能するコード ファクトリは、テストをパイプラインの最上位層として扱い、テストの速度を生成の速度に一致させることに重点を置いています。
リアルタイム検証により、運用前に障害が検出されます。自律エージェントによって生成されたコードはすぐにテストされます。

徹底的に、そして継続的に。悪いコードは決して出荷されません。工場にはゲートがあり、そのゲートが機能します。
自律的なコード生成は、大規模に検証できる場合にのみ大規模に機能します。 QA Wolf の自動化は、世代とともにテストが拡張されることを意味します。発生率を 10 倍にすることなく、10 倍のコードを出荷できます。

## Original Extract

AI agents generate code faster than teams can test it. Why code factories break in production, and how autonomous E2E testing closes the gap.

Code Factories Without Quality: The AI Development Blind Spot | QA Wolf
Platform Mapping AI Autonomously outline your entire app in minutes Automation AI Turn prompts into deterministic Playwright & Appium Run Infra Run web and mobile tests in parallel Service Customers QA Wolf saves Salesloft $750K+ per year Metronome ticks four daily releases without missing a beat QA Wolf puts the magic in Meow Wolf’s mobile testing
Quippy headline about six words long Quippy headline about six words long Quippy headline about six words long More Customer stories Resources Changelog Solutions Docs Blog Webinars Guides Community Pricing Terms of Service Privacy Policy Careers Contact QA Wolf is a hybrid platform & service that helps software teams ship better software faster by taking QA completely off their plate.
Code Factories Without Quality: The AI Development Blind Spot
Key Takeaways - Code multiplied with AI. Verification didn't. That gap is where the incidents live.
- Coverage isn't verification. 100% line coverage still ships broken user flows. E2E catches what coverage misses.
- Verification has to be a pipeline. It must scale like code, runs on every deploy, fixes its own flakes.
Zapier. Nubank. Goldman Sachs. Three companies with great engineering and risk management teams – and all three are now handing coding tasks to AI agents with minimal human intervention.
The workflow of code factories is deceptively simple: a developer hands off a task and the agent plans, codes, reviews, tests, debugs, and ships it end-to-end.
Zapier deployed 800+ AI agents internally, with 89% AI adoption across the entire organization. Nubank pointed autonomous AI agents at one of the most critical projects in the company, migrating its core ETL. Goldman Sachs, meanwhile, is piloting autonomous software engineers on its own codebase.
The appeal of code factories is obvious: velocity. Code generation at scale. 10x speed or more. Entire projects shipped without a human hand touching the keyboard.
But every code factory is quietly skipping or underinvesting in the same step: testing.
Agents are generating production code faster than anyone can verify it. Generation capacity increased 10x. Verification capacity didn’t. So, the code ships untested or undertested making code factories not just about velocity, but about velocity at all costs, quality be damned.
And when it breaks in production (it will), there's no feedback loop to make the next generation any better. The whole system is designed to ship more slop, faster.
The unspoken assumption: AI code is production-ready
Every code factory is built on an assumption that nobody says out loud: generated code is production-ready code.
It's not. But the assumption persists because QA is rarely prominently in the equation.
Instead, QA is being treated as a downstream checkbox or something you can solve with minimal coverage. In a code factory, where velocity is the core value, QA is the friction—the thing preventing you from shipping even faster.
The reasons are complex but because it’s seen as slow, it gets deprioritized or automated away with shallow test coverage that passes but doesn't truly verify much.
And teams often know this but still ship code they know is likely to break because they feel they have no other choice. Fast slop is now seen as better than slightly slower quality.
The real problem is that verification velocity hasn't yet matched generation velocity. AI code review tools are helping with one side of verification. They're catching some bugs, suggesting improvements, flagging obvious issues. But they're still letting bugs through.
What's needed now is an increase in testing velocity.
The cost of undertesting in code factories
A code factory might say they ship with 80% coverage and call that success. But coverage like that typically measures lines executed, not whether user flows actually work. You can have 100% line coverage and still ship code that fails under load, handles errors wrong, or leaves data in an inconsistent state.
The code factory speeds up the rate at which bad code hits production.
And the results are brutal. Incident rates are climbing as code generation scales. Teams shipping AI-generated code are seeing a 30% increase in change failure rates and a 23.5% increase in incidents per pull request.
More code shipped faster without proper verification means more bugs in production, more customer-facing failures, more firefighting.
A bug caught at the PR stage cost minutes to fix. A bug caught in staging costs hours. A bug caught in production can cost days, angry customers, incident responses, and cascading failures across dependent systems. The cost of fixes post-deployment dwarfs the cost of catching them before they ship.
Good testing allows you to ship fast without shipping slop
The real problem is that code factories are optimized for generation speed, not verification speed. QA is positioned as sometimes a blocker to velocity, when it should be seen as the enabler of it. Because if you have good automated test coverage in place, code factories become something you can actually trust to ship fast.
The problem is that you can't hire your way to 10x verification since human test creation is linear, whereas agent generation isn't. You also can't let agents grade their own homework because then tests that inherit the code's blind spots just automate your false confidence.
The tests that actually catch these failures – end-to-end coverage of real user flows – are historically the slowest to build and most painful to maintain. That's the trap: the verification your code factory actually needs is the exact verification that's hardest to build, slowest to run, and most expensive to keep alive, at the precise moment your generation velocity makes all three problems worse.
Verification has to become an autonomous pipeline, too. It needs to have the same velocity and run in parallel, not bolted on downstream as the thing that slows everyone down.
But to do that well requires a few non-negotiables.
It has to scale like generation, not like headcount. Test creation can't depend on how many QA engineers you can hire. It has to expand at the rate code ships because that's the only rate that keeps pace.
It has to test the way users actually use the product. End-to-end coverage across real flows, not line counts. You need to make sure the customer flows work properly, not just that the code executes.
Independent from the coding agent
It has to be independent from the coding agent that generated it. Verification built separately from the generation agent doesn't inherit the code's blind spots. It's an actual second opinion instead of an echo. But more than just that, a coding agent isn’t built to understand how to create good tests. Don’t trust us though. Trust the dozens of threads on Reddit of people struggling to automate test coverage using a coding agent in the last few months.
It has to maintain itself. When a flow changes, the tests update automatically, so flake stays near zero and the team keeps trusting the signal. Self-maintaining coverage is the only kind that survives contact with 10x velocity.
Put those together and QA stops being the friction in the code factory. It becomes the thing that lets you ship fast without shipping slop.
Verification has to finally move at the speed of generation.
How QA Wolf solves code factory QA
This is the part of the article where we do a bit of a plug but bear with us because it’s actually a solution to the problems most teams are facing.
Our platform creates an autonomous QA pipeline that runs at the speed of generation, independent from the code it's checking. It maps your app, writes the tests, runs them on every deploy, and keeps them alive as the code churns. It’s an end-to-end platform for the entire testing lifecycle.
The Mapping Agent independently navigates your application and builds a detailed map of its features and real user workflows, instead of the coding agent’s assumptions about it. It switches user roles and toggles between web, iOS, and Android to catch flows that span multiple users and multiple platforms—the exact multi-step, cross-state behavior where AI-generated code quietly breaks. The agent is 32x faster than a human doing it alone.
It writes real, deterministic tests
The Automation Agent turns those mapped workflows into deterministic Playwright (web) and Appium (mobile) code. These are tests that call APIs, seed databases, mock external dependencies, and flip feature flags to exercise the app under real conditions. Because the tests are deterministic code rather than an LLM re-interpreting a script each run, a failure means something actually broke, and the failure is reproducible. And because the suite is generated independently of whatever produced your production code, it's a genuine second opinion instead of an echo. The Automation Agent increases a QA engineer's capacity 10-20x, which is how test creation finally scales at the rate code ships instead of at the rate you can hire.
It runs on every deploy, in full parallel
Coverage that takes hours to run doesn't help a code factory. QA Wolf runs the full suite with 100% parallel execution, kicked off instantly on deploy. It can be wired directly into your CI pipeline—so verification keeps pace with a pipeline that's merging constantly.
It maintains itself so the suite survives 10x velocity
This is the piece that kills most testing efforts. Autonomous agents don't generate code once; they refactor, optimize, and rewrite it, and every one of those changes threatens to break the suite. QA Wolf's Automation AI addresses close to 100% of flakes including timing issues, runtime errors, un-rendered components, versus the roughly 20% that selector-repair tools catch. When a test does break, it reproduces the failure, diagnoses the cause, rewrites the code, and validates the fix. Meanwhile, the Mapping Agent watches for new features and folds them into existing flows rather than dumping in isolated one-off tests or duplicating tests you already have.
Put the QA Wolf lifecycle together and you get the thing the code factory was missing: verification that maps, writes, runs, and maintains itself at generation speed, so your developers keep shipping and QA stops being the bottleneck.
That's how you ship fast without shipping slop.
Building factories that actually work
QA can’t be a bottleneck in code factories. It has to be foundational. Code factories that work treat testing as a first-class layer in the pipeline and focus on matching testing velocity to generation velocity.
Real-time verification catches failures before production. Code generated by autonomous agents gets tested immediately, comprehensively, and continuously. Bad code never ships. The factory has gates and those gates work.
Autonomous code generation only works at scale if you can verify at scale. QA Wolf's automation means tests scale with generation. You can ship 10x the code without 10x the incident rate.
