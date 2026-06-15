---
source: "https://aws.amazon.com/blogs/machine-learning/how-frontier-teams-are-reinventing-ai-native-development/"
hn_url: "https://news.ycombinator.com/item?id=48539489"
title: "Frontier teams are reinventing AI-native development"
article_title: "How frontier teams are reinventing AI-native development | Artificial Intelligence"
author: "kator"
captured_at: "2026-06-15T11:10:23Z"
capture_tool: "hn-digest"
hn_id: 48539489
score: 1
comments: 0
posted_at: "2026-06-15T10:59:34Z"
tags:
  - hacker-news
  - translated
---

# Frontier teams are reinventing AI-native development

- HN: [48539489](https://news.ycombinator.com/item?id=48539489)
- Source: [aws.amazon.com](https://aws.amazon.com/blogs/machine-learning/how-frontier-teams-are-reinventing-ai-native-development/)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T10:59:34Z

## Translation

タイトル: フロンティア チームは AI ネイティブ開発を再発明しています
記事のタイトル: フロンティア チームが AI ネイティブ開発をどのように再発明しているか |人工知能
説明: フロンティア チームは、コーディングを高速化するために AI を使用しているだけではありません。彼らはソフトウェアの構築方法を再設計しています。その結果、生産性が 4.5 倍、場合によっては 10 倍以上向上します。

記事本文:
フロンティア チームが AI ネイティブ開発をどのように再発明しているか
フロンティア チームは、AI を使用してコーディングを高速化するだけではありません。彼らはソフトウェアの構築方法を再設計しています。その結果、生産性が 4.5 倍、場合によっては 10 倍以上向上します。
エンジニアは6人。七十六日。プロジェクトは 30 人の開発者が 12 ～ 18 か月かけて実施し、1 四半期以内に完了しました。それは仮説ではありません。これは、Amazon Bedrock チームが AI をコーディングのショートカットとして扱うのをやめ、AI を仕事の仕組みの基礎として扱い始めたときに起こったことです。チームは、過去 10 年間よりも多くの製品コードを 5 か月間で出荷しました。
このようなチームと他のチームとの差は急速に広がっています。 AI コーディング エージェントは、ソフトウェアが作成される速度を根本的に変えましたが、ソフトウェアが顧客に届く速度は根本的に変えませんでした。コミットが急増し、CI/CD パイプラインはかつてないほど混雑しています。しかし、本番環境に出荷される機能は同じペースを保っていません。ボトルネックは、エージェントが出力を生成する能力ではありません。それは、エージェントが適切な意思決定を行うために必要な知識にアクセスできることと、その現実に基づいて作業を再構築するチームの意欲です。
これを理解したチームを私たちは「フロンティアチーム」と呼んでいます。彼らはエリート研究室に限定されません。これらは業界や企業規模を超えて存在しており、AI 導入をツールの展開ではなくエンジニアリング投資として扱うという共通の規律を共有しています。どのエンジニアリング チームもフロンティア チームになれる可能性があります。そこに行く方法をご案内します。
Amazon における AI ネイティブ開発への 3 つの道
AI ネイティブのソフトウェア開発では、AI をソフトウェア構築の基礎として扱い、人間の専門家によって指揮されるエージェントの能力がますます高まります。チームがエージェントをどのように指示するかによって、結果が決まります。 Amazon では、AI 開発の主な推進力として、

開発者が文書化、調整、運用などのコーディング以外のタスクに費やす時間を削減し、技術的負債を解消し、数千の小規模な「ピザ 2 枚」の開発者チーム全体でのコーディングの不一致を最小限に抑えることです。私たちは何百ものエンジニアリング チームにわたって実験を行っており、少なくとも 3 つの道筋を特定しました。1 つは課題に取り組む専門家によるパスファインダー イニシアチブ、明確に定義された計画に基づいて実行するための構造化されたスプリント、そして既存のアプローチと AI に適応したワークフローの間でチームを半分に分ける現場実験です。パスの構造は異なりますが、同じ洞察に収束します。
パスファインダーの取り組みは対照的な実験でした。 6 人の上級エンジニアは、Amazon Bedrock 推論エンジンを再構築するという 1 つの任務を受け取りました。このプロジェクトは、当初 30 人の開発者が 12 ～ 18 か月かかると見積もられていました。チームは人員を増やすのではなく、最初の数週間を AI を中心としたワークフローの再設計、個別のタスクから目標主導の結果への移行、複数のエージェントの並行実行、AI が勤務時間外に独立して動作するシステムのセットアップに費やしました。プロジェクトは 76 日で完了しました。正規化されたコミット速度 (リポジトリの複雑さとチームのサイズを調整した、1 週間あたりの開発者あたりのコミット数) で測定した場合、個々の開発者の生産性は約 20 倍増加しました。コミット数は週 2 件から 40 件に増加しました。本番環境にデプロイされたラインで測定したところ、チームは過去 10 年間のプロジェクトよりも多くの高品質コードを 5 か月間で出荷しました。
構造化されたスプリントでは異なるアプローチが取られました。 Prime Video Financial Systems チームは、パスファインダー モデルに触発された 10 日間の実験を実施しました。 6 人のエンジニア、1 つの部屋、コンテキストの切り替えなし、オンコール業務なし、他のプロジェクトなし、会議は限られています。シニアエンジニアの支出

3 週間前に、複雑さを詳細な要件を含む範囲の広いタスクに分割します。チームは、複雑な機能作業には仕様駆動型開発を使用し、要件がすでに明確なタスクには直接エージェント支援開発を使用しました。 10 日間で、ベースラインの 96 に対して 556 のコミットが生成され、90 週間のプロジェクト見積もりが 24 週間に短縮されました。これは、ほぼ 6 倍のスループットと 4 倍の加速に相当します。彼らは、AI による利益の要因として、判断力の低い作業の高速化 (1.5 倍)、コンテキストの切り替えなしで判断力の高い作業への集中力 (1.5 倍)、およびエージェントが取得したドメインの専門知識への即時アクセス (1.5 倍) という 3 つの要因が重なったことによると考えています。どれか 1 つの要素を削除すると、ゲインは崩壊します。チームは現在、ドメインの知識をカプセル化した詳細な製品仕様と、集中時間を解放する自律エージェントを使用して、通常の運用におけるこれら 3 つの要素を最適化することを検討しています。
現場実験 では、50 以上のチームを調査し、新しいツールと新しい実践の両方を導入した 25 チームが、単に既存のワークフローに AI を追加したチームを上回りました。 Amazon ストアは、Kiro と専用の AI ツールを使用して、特別な条件や厳選されたエンジニアを使用せずに、通常のバックログに取り組む典型的な開発チームで構造化されたパイロットを実行しました。生産性の向上の中央値は 4.5 倍で、一部のチームは正規化されたデプロイ速度 (スプリントごとにデプロイされた機能、過去のベースラインに対して正規化されたもの) が 10 倍以上向上しました。 Perfect Order Experience では、機能が 2 週間ではなく午後に出荷されるようになりました。 WW Grocery では、設計ドキュメントの作成を 5 日から数時間に短縮しました。
道は違えど、教訓は同じ。ツールだけではなく、ワークフローも重要です。
フロンティアチームになるための 5 つのステップ
3つの道すべてにおいて、最高のパフォーマンスを発揮するお茶

ms は、共通のロジックを持つ 5 つのプラクティスを共有します。エージェントにとってコンテキストに対する障壁を減らし、エージェントが独立して実行できる作業の表面積を増やします。
これは、フロンティアチームが以前の習慣から逸脱する場所です。個々のコード生成の速度に合わせて最適化された歴史的なアプローチ。フロンティア チームは、これとは異なる点、つまり実稼働対応の正確なソフトウェアが顧客に届く速度を最適化します。この違いが、以下のすべての実践を推進します。
エージェントのコンテキストに投資します。最も先進的なチームは、エージェント ステアリング ファイルやチームの規則、コーディング標準、テスト、コードベース ナビゲーションに関するガイダンスを通じて、エージェントがプロジェクトや知識を利用しやすくすることに多額の投資を行っています。 Bedrock インフラストラクチャ チームは、すべてのコードとドキュメントをモノリポジトリに配置し、AI エージェントが生成したインライン コメントを永続メモリとして扱いながら保持しました。このステップを省略したチームは、エージェントがなぜ同じ間違いを繰り返すのか疑問に思います。
速度を上げるには速度を落としてください。上記の練習には時間がかかり、チームには忍耐力が求められます。成績の良いチームはいずれも、モデルを学習するにつれて最初は作業が遅くなったと報告しました。彼らは、部門を超えた専門知識をエージェント向けに再利用可能なステアリング ドキュメントにエンコードし、LLM が推論できるようにリポジトリを再構築し、コメントを追加して AI 利用のためにコード分割を再構築しました。その学習曲線を突き進み、期待される成果を定義したチームは、初めて加速を経験しました。ワークフローを変更せずにすぐに利益が得られると期待していたチームは失望しました。最初の 2 週間は遅く感じると予想してください。数週間後は劇的に速く感じることが予想されます。 2 週目でやめたチームは、それがさらに悪化することはありません。
エージェントに子守りをするのではなく、食事を与えます。フロンティアチームは安定を維持

明確な結果を伴う明確な範囲のタスクのバックログ。複数のエージェントを並行して実行し、出力を非同期でレビューします。ビルダーは、エージェントがタスクを完了するのを積極的に待っていないときでも、主要な機能が短期間で完了し、作業が進んでいると報告しています。ある主任エンジニアは、エンジニアがコード レビュー、運用サポート、会議の間を移動している間にエージェントが作業したため、わずか「数時間の連続時間」で完全な変更を出荷しました。
コードを記述する前に意図を明示します。構造化された仕様書、詳細な要件文書、または広範囲にわたるタスクの分解を通じて、フロンティア チームはエージェントがコードの生成を開始する前に、「完了」がどのようなものであるかについて明確なコンテキストを確実に把握できるようにします。このアプローチを使用している一部のチームは、コードの 1 ～ 2% のみを手書きしていると報告していますが、以前よりも 1 週間当たりのコミット数が大幅に増加しています。
「シフトテストは終了です。」フロンティア チームは、エージェントがすべての統合テストをローカルで実行し、コードがパイプラインに到達する前に自己修正できるようにツールを構築します。 Prime Video チームは、問題を早期に発見する自動ガードレール、コンポーネント テスト、パフォーマンス テスト、フォーマッタに投資しました。コード レビューでは、コード スタイルや命名規則ではなく、インターフェイス定義やアーキテクチャ上の決定に焦点が移されました。
テクノロジーリーダーが今日できること
すべてのチームがこのような結果を達成できるわけではありません。コンテキスト構築フェーズをスキップしたり、AI をドロップイン代替品として扱ったり、働き方を再構築せずにすぐに利益を期待したりするチームは、常にパフォーマンスを下回っています。業界全体の開発者が AI コーディング ツールを採用しています。そのすべてが生産量の増加を実感しているわけではありません。彼らは間違ったツールを使用しているわけではありません。彼らは間違ったワークフロー内で正しいツールを使用しています。
働き方を変えることで

AI は最高のパフォーマンスを発揮します。
3 つの要素が相乗して成果をもたらします。判断力の低い作業を処理する AI x 判断力の高い作業への中断のない集中力 x 専門知識への即時アクセスです。
実際の出発点は、広範な展開ではありません。意図的なパイロットです。実稼働コードを作成する前に、エージェント コンテキスト (ステアリング ファイル、仕様テンプレート、モノリポジトリ) の構築に最初の数週間を費やすことをいとわない小規模なチームから始めます。チームにワークフローを再構築する権限を与えます。開発者の満足度スコアとともに、コミット速度、デプロイ頻度、解決までの時間を測定します。次に、学んだことを活用して、組織の残りの部分向けのハンドブックを構築します。
4.5 倍から 10 倍以上の生産性向上を達成しているチームは、単により優れたテクノロジーを採用しているだけではありません。彼らはそれを別の方法で扱う方法を見つけ出しました。この決定は、現在、すべてのエンジニアリング組織が利用できます。もちろん、コードのコミット速度は話の一部にすぎません。私たちは、リリース管理、運用、セキュリティ運用の合理化であっても、EOL アップグレードやソフトウェア エンジニアリングに伴う無数の未分化タスクへの取り組みであっても、ソフトウェア開発ライフサイクルのあらゆる側面を支援したいと考えています。次回のブログでは、これらにどのようにアプローチするかを説明しますので、ご期待ください。
フロンティアチームについて詳しく知る >
AI ネイティブ開発の詳細については、AWS Summit New York City をご覧ください。
Swami Sivasubramanian は、アマゾン ウェブ サービス (AWS) の Agentic AI 担当副社長です。 AWS では、Swami は、Amazon DynamoDB、Amazon SageMaker、Amazon Bedrock、Amazon Q などの主要な AI サービスの開発と成長を主導してきました。彼のチームの使命は、顧客やパートナーが自信を持ってエージェント AI を使用して革新し、信頼性の高いエージェントを構築するために必要な規模、柔軟性、価値を提供することです。

強力で効率的であるだけでなく、信頼性と責任感も備えています。スワミはまた、2022 年 5 月から 2025 年 5 月まで、国家人工知能諮問委員会の委員を務めました。この諮問委員会は、国家 AI イニシアチブに関連するトピックについて米国大統領と国家 AI イニシアチブ事務局に助言する任務を負っていました。
英語
トップに戻る
Amazon は、マイノリティ / 女性 / 障害者 / 退役軍人 / 性同一性 / 性的指向 / 年齢の機会均等雇用主です。
×
フェイスブック
リンクイン
インスタグラム
けいれん
ユーチューブ
ポッドキャスト
電子メール
プライバシー

## Original Extract

Frontier teams are not just using AI to code faster. They’re redesigning how software gets built. The result is 4.5x productivity gains, in some cases more than 10x.

How frontier teams are reinventing AI-native development
Frontier teams are not just using AI to code faster. They’re redesigning how software gets built. The result is 4.5x productivity gains, in some cases more than 10x.
Six engineers. Seventy-six days. A project scoped for 30 developers over 12 to 18 months, delivered within a quarter. That is not hypothetical. It’s what happened when an Amazon Bedrock team stopped treating AI as a coding shortcut and started treating it as the foundation of how they work. The team shipped more production code in five months than in the previous ten years.
The gap between teams like this and everyone else is widening fast. AI coding agents have fundamentally changed the rate at which software gets written, but not the rate at which it reaches customers. Commits are surging, and CI/CD pipelines are busier than ever. Yet, features shipped to production have not kept the same pace. The bottleneck is not the agent’s ability to generate output. It is the agent’s access to the knowledge it needs to make good decisions, and the team’s willingness to restructure work around that reality.
We call the teams that have figured this out “frontier teams.” They are not confined to elite labs. They exist across industries and company sizes, and they share a common discipline: they treat AI adoption as an engineering investment, not a tool rollout. Any engineering team can become a frontier team; we can show you how to get there.
Three paths to AI-native development at Amazon
AI-native software development treats AI as the foundation of how software is built, with increasingly capable agents directed by human experts. How teams direct those agents determines outcomes. At Amazon, the primary drivers for AI in development were to reduce the time developers spent on non-coding tasks such as documentation, coordination, and operations, retire technical debt, and minimize coding inconsistencies across thousands of small “two-pizza” teams of developers. We have been experimenting across hundreds of engineering teams and have identified at least three paths: a pathfinder initiative with experts tackling a challenge, a structured sprint to execute on a well-defined plan, and an in-situ experiment splitting teams in half between existing approaches and AI-adapted workflows. The paths differ in structure but converge on the same insight.
The pathfinder initiative was a controlled experiment. Six senior engineers received a single mandate: rebuild the Amazon Bedrock inference engine, a project originally estimated at 30 developers working 12 to 18 months. Rather than adding headcount, the team spent its first weeks redesigning workflows around AI, shifting from discrete tasks to goal-driven outcomes, running multiple agents in parallel, and setting up systems for AI to work independently during off-hours. The project was delivered in 76 days. Individual developer productivity increased approximately 20x as measured by normalized commit velocity (the number of commits per developer per week, adjusted for repository complexity and team size). Commits went from 2 per week to 40. The team shipped more high-quality code in five months than it did on projects over the previous ten years, as measured by lines deployed to production.
The structured sprint took a different approach. The Prime Video Financial Systems team ran a 10-day experiment inspired by the pathfinder model. Six engineers, one room, zero context switching, no on-call duties, no other projects, limited meetings. A senior engineer spent three weeks beforehand breaking complexity into well-scoped tasks with detailed requirements. The team used spec-driven development for complex feature work and direct agent-assisted development for tasks where requirements were already clear. Over 10 days, they produced 556 commits against a baseline of 96 and reduced a 90-week project estimate to 24 weeks. That translates to nearly 6x throughput and 4x acceleration. They attributed the AI-enabled gain to three factors multiplying together: acceleration of low-judgment work (1.5x), higher focus on high-judgment work with no context-switching (1.5x), and instant access to agent-captured domain expertise (1.5x). Remove any one factor and the gains collapse. The team is now looking to optimize these three factors in normal operations using detailed product specs that encapsulate domain knowledge and autonomous agents that free up focus time.
In the in-situ experiment , of the 50-plus teams studied, the 25 teams that implemented both new tools and new practices outperformed those that simply added AI to existing workflows. Amazon Stores ran structured pilots with typical development teams working against their regular backlogs, using Kiro and purpose-built AI tools with no special conditions and no handpicked engineers. The median productivity gain was 4.5x, with some teams reaching more than 10x improvement in normalized deployment velocity (features deployed per sprint, normalized against historical baselines). Perfect Order Experience now ships features in an afternoon instead of two weeks. WW Grocery cut design document creation from five days to a few hours.
Different paths, same lesson. The workflow matters, not just the tool.
Five steps to becoming a frontier team
Across all three paths, the highest-performing teams share five practices with a common logic. Reduce the barriers to context for the agent and increase the surface area of work it can do independently.
This is where frontier teams diverge from prior habits. The historical approach optimized for the speed of individual code generation. Frontier teams optimize for something different: the rate at which correct, production-ready software reaches customers. That distinction drives every practice below.
Invest in agent context. The most advanced teams invest heavily in making projects and knowledge easier for agents to consume through agent steering files and guidance on team conventions, coding standards, testing, and codebase navigation. The Bedrock infrastructure team placed all code and documentation into a monorepo and kept the inline commentary that AI agents generated, treating it as persistent memory. Teams that skip this step wonder why their agents keep making the same mistakes.
Slow down to speed up. The above-mentioned practice takes time and requires teams to be patient. Every high-performing team reported that things initially slowed down as they learned the models. They encoded cross-functional expertise into reusable steering docs for agents, restructured repositories so LLMs could reason over them, and added comments and re-architected code splits for AI consumption. The teams that pushed through that learning curve and defined the expected outcomes first experienced compounding acceleration. The teams that expected immediate gains without changing their workflows were disappointed. Expect the first two weeks to feel slower. Expect the weeks after to feel dramatically faster. The teams that quit in week two never see the compounding.
Feed agents instead of babysitting them. Frontier teams maintain a steady backlog of well-scoped tasks with clear outcomes, running multiple agents in parallel and reviewing output asynchronously. Builders report finishing major features in short bursts, with work advancing even when they are not actively waiting for the agent to complete a task. One principal engineer shipped a complete change with only ‘a couple of hours of contiguous time’ because the agent worked while the engineer moved between code reviews, operational support, and meetings.
Make intent explicit before code gets written. Whether through structured specifications, detailed requirements documents, or well-scoped task decomposition, frontier teams ensure agents have clear context about what ‘done’ looks like before they start generating code. Some teams using this approach report handwriting only 1–2% of their code while pushing significantly more commits per person per week than before.
“Shift testing left.” Frontier teams build tooling so agents can run all integration tests locally and self-correct before code ever reaches the pipeline. The Prime Video team invested in automated guardrails, component tests, performance tests, and formatters that caught issues early. Code reviews shifted focus to interface definitions and architectural decisions rather than code style and naming conventions.
What technology leaders can do today
Not every team achieves these results. Teams that skip the context-building phase, treat AI as a drop-in replacement, or expect immediate gains without restructuring how they work consistently underperform. Developers across the industry have adopted AI coding tools. Not all of them are seeing production gains. They are not using the wrong tools. They’re using the right tools inside the wrong workflows.
Change how you work to make AI work at its best.
Three factors multiply to deliver outcomes: AI handling low-judgment work x uninterrupted focus on high-judgment work x instant access to domain expertise.
The practical starting point is not a broad rollout. It is a deliberate pilot. Start with a small team willing to spend the first weeks building agent context (steering files, spec templates, monorepos) before writing production code. Give the team a mandate to restructure workflows. Measure commit velocity, deployment frequency, and time-to-resolution, along with developer satisfaction scores. Then use what they learn to build the playbook for the rest of the organization.
The teams achieving 4.5x to more than 10x productivity gains have not just adopted better technology. They’ve figured out how to work differently with it. That decision is available to every engineering organization today. Of course, code commit velocity is only part of the story. We want to help with all aspects of the software development lifecycle, whether that is streamlining release management, operations, and security operations, or tackling EOL upgrades and the countless undifferentiated tasks that come with software engineering. Stay tuned for the next blog, where I will walk through how we are approaching these.
Learn more about frontier teams >
Tune in to AWS Summit New York City for more on AI-native development.
Swami Sivasubramanian is Vice President for Agentic AI at Amazon Web Services (AWS). At AWS, Swami has led the development and growth of leading AI services like Amazon DynamoDB, Amazon SageMaker, Amazon Bedrock, and Amazon Q. His team’s mission is to provide the scale, flexibility, and value that customers and partners require to innovate using agentic AI with confidence and build agents that are not only powerful and efficient, but also trustworthy and responsible. Swami also served from May 2022 through May 2025 as a member of the National Artificial Intelligence Advisory Committee, which was tasked with advising the President of the United States and the National AI Initiative Office on topics related to the National AI Initiative.
English
Back to top
Amazon is an Equal Opportunity Employer: Minority / Women / Disability / Veteran / Gender Identity / Sexual Orientation / Age.
x
facebook
linkedin
instagram
twitch
youtube
podcasts
email
Privacy
