---
source: "https://www.microsoft.com/en-us/research/blog/orchard-an-open-framework-for-scalable-agentic-ai/"
hn_url: "https://news.ycombinator.com/item?id=49158402"
title: "Orchard: An open framework for scalable agentic AI"
article_title: "Orchard: An open framework for scalable agentic AI - Microsoft Research"
author: "andsoitis"
captured_at: "2026-08-03T17:47:46Z"
capture_tool: "hn-digest"
hn_id: 49158402
score: 4
comments: 0
posted_at: "2026-08-03T16:57:55Z"
tags:
  - hacker-news
  - translated
---

# Orchard: An open framework for scalable agentic AI

- HN: [49158402](https://news.ycombinator.com/item?id=49158402)
- Source: [www.microsoft.com](https://www.microsoft.com/en-us/research/blog/orchard-an-open-framework-for-scalable-agentic-ai/)
- Score: 4
- Comments: 0
- Posted: 2026-08-03T16:57:55Z

## Translation

タイトル: Orchard: スケーラブルなエージェント AI のオープン フレームワーク
記事のタイトル: Orchard: スケーラブルなエージェント AI のオープン フレームワーク - Microsoft Research
説明: Orchard は、タスク タイプ全体で AI エージェントをトレーニングおよび評価するための、研究コミュニティのためのオープンソース フレームワークです。研究者が同じインフラストラクチャを再利用できるようにすることで、複雑さを軽減しながら、より小規模なモデルから強力なパフォーマンスをサポートします。

記事本文:
メインコンテンツにスキップ
研究
出版物
コードとデータ
人々
マイクロソフトリサーチブログ
人工知能
オーディオと音響
コンピュータビジョン
グラフィックスとマルチメディア
人間とコンピュータの相互作用
人間の言語技術
検索と情報取得
データプラットフォームと分析
ハードウェアとデバイス
プログラミング言語とソフトウェアエンジニアリング
量子コンピューティング
セキュリティ、プライバシー、暗号化
システムとネットワーク
アルゴリズム
数学
エコロジーと環境
経済学
医療、健康、ゲノミクス
社会科学
新興市場向けのテクノロジー
学術プログラム
イベント・学会
マイクロソフトリサーチフォーラム
テクノロジー ポッドキャストの裏側
マイクロソフトリサーチブログ
マイクロソフトリサーチフォーラム
マイクロソフトリサーチのポッドキャスト
マイクロソフトリサーチについて
キャリアとインターンシップ
人々
名誉プログラム
ニュースと受賞歴
Microsoft Research ニュースレター
アフリカ
科学のための AI
AI フロンティア
アジア太平洋地域
ケンブリッジ
健康の未来
インド
モントリオール
ニューイングランド
ニューヨーク市
レドモンド
応用科学
複合現実と AI - ケンブリッジ
複合現実と AI - チューリッヒ
登録: 研究フォーラム
マイクロソフトのセキュリティ
アズール
ダイナミクス 365
マイクロソフト 365
マイクロソフトチーム
Windows 365
マイクロソフトAI
アズールスペース
複合現実
マイクロソフト ホロレンズ
マイクロソフト ビバ
量子コンピューティング
持続可能性
教育
自動車
金融サービス
政府
ヘルスケア
製造業
小売
パートナーを探す
パートナーになる
パートナーネットワーク
マイクロソフト マーケットプレイス
ソフトウェア会社
ブログ
マイクロソフトの広告
開発者センター
ドキュメント
イベント
ライセンス
Microsoft Learn
マイクロソフトリサーチ
サイトマップを見る
ブログホームに戻る
マイクロソフトリサーチブログ
Orchard: スケーラブルなエージェント AI のオープン フレームワーク
によって
彭宝林
、
主任研究マネージャー
ヤオ・ウェンリン
、
主任研究員
呉乾輝
、
主任研究員
ハオ・チェン
、
主任研究員
建峰G

あお
、
テクニカルフェロー兼執行役員
Orchard は、スケーラブルでコスト効率の高いエージェント AI 研究のためのオープンソース フレームワークで、タスク ドメイン全体でエージェントをトレーニングおよび評価するための再利用可能な環境サービスである Orchard Env を中心に構築されています。
同じ Orchard インフラストラクチャは、ソフトウェア エンジニアリング、Web ナビゲーション、およびパーソナル アシスタント エージェントをサポートし、Codex、OpenClaw、ZeroClaw などの実際の導入ハーネス内でこれらを直接トレーニングできるため、研究者は環境、データ パイプライン、評価ワークフローをタスク間で再利用できます。
Orchard-SWE、Orchard-GUI、および Orchard-Claw は、比較的小さな無重みモデルが現実世界の複雑なタスクで強力な結果を達成できることを示しています。たとえば、Orchard-SWE は、わずか約 30 億のアクティブ パラメーターを使用して、SWE ベンチ検証済みで 69.7% に達し、値モデルの再ランキングでは 73.0% に達し、10 倍を超えるモデルを使用するフロンティア システムに近づいています。
このプロジェクトは、モデルとワークフローに加えて、より広範な研究コミュニティによるオープン エージェント システムの構築と研究を支援することを目的としたトレーニング データと評価方法をリリースします。人工知能は、静的な質問応答を超えて、複雑な複数ステップの環境にわたって計画、推論、行動できる自律エージェントへと急速に移行しています。これらのシステムは、複雑なコードベースのバグを修正し、ユーザーに代わって Web を操作し、カレンダーや電子メールを含むワークフローを管理できます。
エージェント AI の機能について興奮が高まっている一方で、研究コミュニティは永続的なボトルネックに直面しています。最先端のエージェント システムを構築するには、多くの場合、カスタム サンドボックス、クローズド トレーニング パイプライン、ほとんどの研究者や実践者がアクセスしたり再現したりできない独自のデータセットなど、独自のインフラストラクチャが必要になります。
このギャップに対処するために、Orch を導入します。

ard (新しいタブで開きます) 、スケーラブルなエージェント モデリング用のオープンソース フレームワーク。 Orchard の中心には、軽量の Kubernetes 環境である Orchard Env があり、トレーニング データの収集から強化学習のロールアウトと評価まで、大規模なエージェントの実行と構築のための再利用可能な分離コンポーネントを提供します。
多くの既存のフレームワークとは異なり、Orchard Env は、さまざまなエージェント システムとタスク タイプを変更せずにサポートするように設計されています。同じサービスは、ドメイン全体でソフトウェア エンジニアリング エージェント、Web ブラウジング エージェント、およびパーソナル アシスタント エージェントをサポートできます。
このアプローチを実証するために、Orchard-SWE、Orchard-GUI、Orchard-Claw という 3 つのドメイン固有のトレーニング レシピをリリースします。 （新しいタブで開きます） 学習データとその構築に用いた評価手法も公開しています。
AI のテストと評価: 科学と産業から学んだこと
Microsoft が AI ガバナンスの柱として評価とテストを進めるために他の分野からどのように学んでいるかをご覧ください。
タスクの種類に応じて拡張できる環境レイヤー
Orchard の背後にある中心的な考え方は、ランタイム環境は特定のトレーニング フレームワーク内に組み込まれたインフラストラクチャではなく、スタンドアロンで再利用可能なサービスであるべきだということです。 Orchard Env の Kubernetes 基盤により、何千もの分離されたコンポーネントを並行して作成、管理、削除できます。
このシステムは、コーディング、Web ブラウジング、ツールの使用などのタスクにわたって機能するように設計されています。また、データの蒸留や強化学習のロールアウトなどのトレーニングと評価プロセスの段階とともに、さまざまなエージェント システムにわたって機能するように設計されています。
この柔軟性により、Orchard は研究規模で実用的になります。チームは、基盤となるインフラストラクチャを再構築することなく、新しいベンチマーク、エージェント システム、トレーニング アルゴリズムを導入できます。

ゼロからやり直します。
Orchard を使用すると、あらゆるハーネス内でエージェントをトレーニングすることもできます。現在の最も有能なエージェントがベア モデルとして実行されることはほとんどありません。これらは、Claude Code、Codex、OpenClaw などの高度なハーネスを通じて動作し、マルチターン推論、ツールの使用、外部システムへの接続を管理します。オープン トレーニング ツールは通常、こうしたステートフルなマルチプロセス ハーネスを処理できないため、研究者は簡略化された代役でトレーニングしてから実際の環境に導入する必要があり、不一致が生じます。 Orchard はこのギャップを埋めます。各ロールアウトが独自のコンテナ内で実行されている間に、軽量プロキシがハーネス独自のモデル呼び出しをトレーニング データとして記録するため、エージェントは、OpenClaw、Codex、ZeroClaw などの展開されるハーネス内で、また複数のハーネスにわたって直接エンドツーエンドでトレーニングできます。
Orchard-SWE: オープンソース ソフトウェア エンジニアリング エージェントの進歩
ソフトウェア エンジニアリングは、自律エージェントにとって最も要求の厳しい設定の 1 つです。実際のコードベースに対する複数段階の推論、ツールの使用、間違いから回復する能力が必要です。 Orchard-SWE は、このドメインのトレーニング ワークフローです。これは、ソフトウェア エンジニアリング タスクを自律的に解決するように設計された Mini-SWE-Agent フレームワークを使用して構築され、現実世界のコードベースをナビゲート、診断、修復するモデルの能力をテストする、広く使用されている SWE ベンチ検証済みベンチマークで評価されます。
システムをトレーニングするために、GitHub の広範な問題をカバーする 2 つの高度なオープンウェイト モデル (MiniMax-M2.5 および Qwen3.5-397B) から 107,000 のエージェント インタラクションを抽出しました。トレーニング プロセスでは、クレジット割り当ての監視下での微調整が使用されます。エージェントが問題を完全に解決できなかった試行を破棄するのではなく、システムはそれらの部分的な試行のうち生産的な部分から学習し、AM の機能を拡張します。

モデルで利用できる有用なトレーニング データが不足しています。
次に強化学習が行われますが、そのフィードバックはまばらです。エージェントは通常、最終パッチが隠れたテストに合格したか不合格だったかのみを学習します。これらのまれな成功シグナルを最大限に活用するように設計されたバランス型適応ロールアウトから始め、次に、より充実したガイダンスを実現する 2 つの「高密度報酬」テクニックを追加します。オンポリシー蒸留では、強力な教師モデルがエージェントの決定を段階的にスコア付けします。もう 1 つは、最終テストが合格したかどうかに関係なく、AI 審査員が健全な問題解決プロセス (バグを再現するテストの作成、修正の検証、既存の動作が引き続き機能するかどうかの確認) に報酬を与えるプロセス報酬モデルです。
最後に、過去の展開に基づいて価値モデルをトレーニングし、候補ソリューションを再ランク付けします。強化学習は、通常は破棄される多くの練習軌道を生成します。その代わりに、これまでの 20 回の実験からの軌跡によって、高品質の解決策を認識するコンパクトな 40 億パラメータ値モデルがトレーニングされ、問題解決時にいくつかの回答候補にスコアが付けられ、最良の回答が選択されます。これらの手法を組み合わせると、Orchard-SWE は、SWE ベンチ検証済みのベースラインの 61.4% から、Balanced Adaptive Rollout で 69.1%、そして同等の規模 (約 30 億のアクティブ パラメータ) のオープンソース モデルの中で最新鋭である高密度報酬手法で 69.7% に達し、値モデルの再ランキングにより 73% に上昇し、図 1 に示すように 10 倍以上のフロンティア システムに近づいています。
Orchard-GUI: 現実世界の Web タスク用の軽量ブラウザ エージェント
Web ナビゲーションにはさまざまな課題があります。エージェントは、視覚的なレイアウトを解釈し、動的なインターフェイスと対話し、自然言語のみで記述された無制限のタスクを完了する必要があります。
Orchard-GUI は 4-bi をトレーニングします

比較的少量の監視、つまり 2,200 のオープンエンド トレーニング タスクを組み合わせた 400 の抽出されたデモンストレーションを使用して、ブラウザ エージェントとして 100 パラメータの視覚言語モデルを構築します。このトレーニング データが限られているにもかかわらず、結果のモデルは、いくつかの Web ナビゲーション ベンチマークで優れた結果を達成しています。図 1 に示すように、WebVoyager で 74.1%、Online-Mind2Web で 67.0%、DeepShop で 64.0%、平均 68.4% という結果が得られました。
これらの結果により、Orchard-GUI は、大規模な独自モデルとの競争力を維持しながら、これまでで最も強力なオープンソース Web エージェントにランクされます。この結果は、適切なトレーニング アプローチと環境があれば、小規模なオープン モデルが現実世界の Web タスクで適切に実行できることも示唆しています。
Orchard-Claw: 日常の生産性を高めるパーソナル アシスタント エージェント
最も影響力のあるエージェント アプリケーションの多くには、電子メールの読み取りと下書き、カレンダーの管理、情報の検索、ツール間での調整など、日常の生産性タスクが含まれます。 Orchard-Claw は、わずか 200 の合成タスクについてエージェントをトレーニングすることで、パーソナル アシスタントのタスクに重点を置いています。現実的な生産性ワークフローをカバーするベンチマークである Claw-Eval で評価すると、最大 3 回の試行でタスクの 59.6% を正常に完了できます。より強力な ZeroClaw エージェント システムと組み合わせると、この割合は 73.9% に増加します。
Orchard は実際の導入ハーネス内でエージェントを直接トレーニングできるため、Orchard-Claw は、単一の単純化されたループではなく、ReACT、ZeroClaw、OpenClaw、Codex を含む複数のハーネスにわたってトレーニングされます。これらの実際のハーネス内でのトレーニングにより、エージェントの信頼性が大幅に向上します。たとえば、Codex ハーネスでは、成功率が未トレーニング モデルの 18.6% から Orchard トレーニング後の 51.5% に上昇します。
影響と今後の道のり
オーチャードの結果は、より広範な

ポイント: 環境層が重要です。 Orchard は、基盤となるインフラストラクチャをオープン、軽量、再利用可能にすることで、エージェント AI 研究のコストを削減します。チームは、分離されたカスタム環境を一から構築したり、独自のクラウド サービスに依存したりする必要がなくなりました。同じ Orchard Env を使用して、毎回システムを再構築することなく、トレーニング データの生成、強化学習ロールアウトの実行、最終モデルの評価を行うことができます。
将来的には、トレーニング エクスペリエンスの再利用が、エージェントの累積学習に向けた有望な方向性になると考えています。トレーニング実行の終了後に軌跡を破棄するのではなく、それらを永続的な資産として扱います。たとえば、再利用可能な価値モデルに抽出します。これにより、エージェントの経験が時間の経過とともに蓄積され、新しい世代のエージェントが最初から始めるのではなく、以前のエージェントが獲得した知識を継承して拡張できるようになります。
Orchard-GUI によって実証されたデータ効率は、手動で作成された大量のトレーニング データを必要とせずに、大規模な Web エージェントをトレーニングできることを示唆しています。環境サービス、トレーニング パイプライン、トレーニング データセットを含む完全な Orchard スタックをリリースすることで、より広範な研究コミュニティがより有能なオープン エージェントを構築できるよう支援したいと考えています。

[切り捨てられた]

## Original Extract

Orchard is an open-source framework for the research community to train and evaluate AI agents across task types. It reduces complexity while supporting strong performance from smaller models by enabling researchers to reuse the same infrastructure:

Skip to main content
Research
Publications
Code & data
People
Microsoft Research blog
Artificial intelligence
Audio & acoustics
Computer vision
Graphics & multimedia
Human-computer interaction
Human language technologies
Search & information retrieval
Data platforms and analytics
Hardware & devices
Programming languages & software engineering
Quantum computing
Security, privacy & cryptography
Systems & networking
Algorithms
Mathematics
Ecology & environment
Economics
Medical, health & genomics
Social sciences
Technology for emerging markets
Academic programs
Events & academic conferences
Microsoft Research Forum
Behind the Tech podcast
Microsoft Research blog
Microsoft Research Forum
Microsoft Research podcast
About Microsoft Research
Careers & internships
People
Emeritus program
News & awards
Microsoft Research newsletter
Africa
AI for Science
AI Frontiers
Asia-Pacific
Cambridge
Health Futures
India
Montreal
New England
New York City
Redmond
Applied Sciences
Mixed Reality & AI - Cambridge
Mixed Reality & AI - Zurich
Register: Research Forum
Microsoft Security
Azure
Dynamics 365
Microsoft 365
Microsoft Teams
Windows 365
Microsoft AI
Azure Space
Mixed reality
Microsoft HoloLens
Microsoft Viva
Quantum computing
Sustainability
Education
Automotive
Financial services
Government
Healthcare
Manufacturing
Retail
Find a partner
Become a partner
Partner Network
Microsoft Marketplace
Software companies
Blog
Microsoft Advertising
Developer Center
Documentation
Events
Licensing
Microsoft Learn
Microsoft Research
View Sitemap
Return to Blog Home
Microsoft Research Blog
Orchard: An open framework for scalable agentic AI
By
Baolin Peng
,
Principal Research Manager
Wenlin Yao
,
Principle Researcher
Qianhui Wu
,
Senior Researcher
Hao Cheng
,
Principal Researcher
Jianfeng Gao
,
Technical Fellow & Corporate Vice President
Orchard is an open-source framework for scalable and cost-effective agentic AI research, built around Orchard Env, a reusable environment service for training and evaluating agents across task domains.
The same Orchard infrastructure supports software-engineering, web-navigation, and personal-assistant agents, and can train them directly inside real deployment harnesses such as Codex, OpenClaw, and ZeroClaw—letting researchers reuse environments, data pipelines, and evaluation workflows across tasks.
Orchard-SWE, Orchard-GUI, and Orchard-Claw demonstrate that relatively small open-weight models can achieve strong results on complex real-world tasks. For example, Orchard-SWE reaches 69.7% on SWE-bench Verified—73.0% with value-model reranking—using only about 3 billion active parameters, approaching frontier systems using more than 10 times larger models.
Alongside the models and workflows, the project releases training data and evaluation methods intended to help the broader research community build and study open agentic systems. Artificial intelligence is rapidly moving beyond static question-answering toward autonomous agents that can plan, reason, and act across complex, multistep environments. These systems can fix bugs in complex codebases, navigate the web on a user’s behalf, and manage workflows involving calendars and email.
While there is excitement around agentic AI’s capabilities, the research community faces a persistent bottleneck. Building state-of-the-art agentic systems often requires proprietary infrastructure, including custom sandboxes, closed training pipelines, and proprietary datasets that most researchers and practitioners cannot access or reproduce.
To address this gap, we introduce Orchard (opens in new tab) , an open-source framework for scalable agentic modeling. At the center of Orchard is Orchard Env, a lightweight, Kubernetes environment that provides reusable isolated components for running and building agents at scale—from collecting training data to reinforcement learning rollouts and evaluation.
Unlike many existing frameworks, Orchard Env is designed to support different agent systems and task types without modification. The same service can support software-engineering agents, web-browsing agents, and personal-assistant agents across domains.
To demonstrate this approach, we are releasing three domain-specific training recipes— Orchard-SWE, Orchard-GUI, and Orchard-Claw. (opens in new tab) We are also releasing the training data and evaluation methods used to build them.
AI Testing and Evaluation: Learnings from Science and Industry
Discover how Microsoft is learning from other domains to advance evaluation and testing as a pillar of AI governance.
Environment layer that scales across types of tasks
The central idea behind Orchard is that the runtime environment should be a standalone, reusable service rather than infrastructure embedded inside a specific training framework. Orchard Env’s Kubernetes foundation enables it to create, manage, and remove thousands of isolated components in parallel.
The system is designed to work across tasks like coding, web browsing, using tools. It is also designed to work across different agent systems, along with stages of the training and evaluation process, including data distillation and reinforcement learning rollouts.
This flexibility makes Orchard practical at a research scale. Teams can introduce new benchmarks, agent systems, or training algorithms without rebuilding the underlying infrastructure from scratch.
Orchard also makes it possible to train agents inside any harness. Today’s most capable agents rarely run as a bare model. They operate through sophisticated harnesses—such as Claude Code, Codex, and OpenClaw—that manage multi-turn reasoning, tool use, and connections to external systems. Open training tools usually cannot handle these stateful, multi-process harnesses, forcing researchers to train on a simplified stand-in and then deploy in the real setting, which creates a mismatch. Orchard closes this gap: a lightweight proxy records the harness’s own model calls as training data while each rollout runs in its own container, so an agent can be trained end-to-end directly in the harness that it will be deployed with—OpenClaw, Codex, ZeroClaw, or others—and across several harnesses.
Orchard-SWE: Advancing open-source software engineering agents
Software engineering is one of the most demanding settings for autonomous agents. It requires multi-step reasoning over real codebases, tool use, and the ability to recover from mistakes. Orchard-SWE is our training workflow for this domain. It is built using the Mini-SWE-Agent framework, designed to autonomously solve software engineering tasks, and evaluated on the widely used SWE-bench Verified benchmark, which tests a model’s ability to navigate, diagnose, and repair real-world codebases.
To train the system, we distilled 107,000 agent interactions from two advanced open-weight models (MiniMax-M2.5 and Qwen3.5-397B) covering a broad range of GitHub Issues. The training process uses credit-assignment supervised fine-tuning: rather than discarding attempts where the agent failed to fully resolve an issue, the system learns from the productive portions of those partial attempts, expanding the amount of useful training data available to the model.
Reinforcement learning comes next, but its feedback is sparse—an agent usually learns only whether its final patch passed or failed the hidden tests. We start with Balanced Adaptive Rollout, designed to make the most of these infrequent success signals, and then add two “dense reward” techniques for richer guidance: on-policy distillation, in which a stronger teacher model scores the agent’s decisions step by step, and a process reward model, in which an AI judge rewards sound problem-solving process—writing tests that reproduce the bug, verifying the fix, and checking that existing behavior still works—independent of whether the final tests passed.
Finally, we train a value model on past rollouts to rerank candidate solutions. Reinforcement learning generates many practice trajectories that are normally discarded; instead, trajectories from 20 prior experiments train a compact 4-billion-parameter value model that recognizes high-quality solutions, and at problem-solving time it scores several candidate answers and picks the best one. Together, these techniques take Orchard-SWE from a 61.4% baseline on SWE-bench Verified to 69.1% with Balanced Adaptive Rollout and 69.7% with the dense-reward techniques—a new state of the art among open-source models of comparable size (roughly 3 billion active parameters)—rising to 73% with value-model reranking, approaching frontier systems more than 10 times larger, as shown in Figure 1.
Orchard-GUI: A lightweight browser agent for real-world web tasks
Web navigation presents a different set of challenges. Agents must interpret visual layouts, interact with dynamic interfaces, and complete open-ended tasks described only in natural language.
Orchard-GUI trains a 4-billion-parameter vision-language model as a browser agent using a relatively small amount of supervision: 400 distilled demonstrations combined with 2,200 open-ended training tasks. Despite this limited training data, the resulting model achieves strong results across several web-navigation benchmarks: 74.1% on WebVoyager, 67.0% on Online-Mind2Web, and 64.0% on DeepShop, for an average of 68.4%, as shown in Figure 1.
These results place Orchard-GUI among the strongest open-source web agents to date while remaining competitive with larger proprietary models. The results also suggest that with the right training approach and environment, small open models can perform well on real-world web tasks.
Orchard-Claw: Personal assistant agents for everyday productivity
Many of the most impactful agentic applications involve everyday productivity tasks, including reading and drafting emails, managing calendars, searching for information, and coordinating across tools. Orchard-Claw focuses on personal-assistant tasks by training an agent on just 200 synthetic tasks. Evaluated on Claw-Eval, a benchmark covering realistic productivity workflows, it successfully completes 59.6% of tasks when given up to three attempts. That increases to 73.9% when paired with the stronger ZeroClaw agent system.
Because Orchard can train agents directly inside real deployment harnesses, Orchard-Claw is trained across several of them—including ReACT, ZeroClaw, OpenClaw, and Codex—rather than a single simplified loop. Training inside these real harnesses substantially improves the agent’s reliability; under the Codex harness, for example, its success rate rises from 18.6% for the untrained model to 51.5% after Orchard training.
Implications and the road ahead
Orchard’s results reinforce a broader point: the environment layer matters. By making the underlying infrastructure open, lightweight, and reusable, Orchard lowers the cost of agentic AI research. Teams no longer need to build custom isolated environments from scratch or depend on proprietary cloud services. The same Orchard Env can be used to generate training data, run reinforcement learning rollouts, and evaluate final models without rebuilding the system each time.
Looking ahead, we see reusing training experience as a promising direction toward cumulative agent learning. Instead of discarding trajectories once a training run finishes, we treat them as persistent assets—for example, distilling them into reusable value models. This enables agentic experience to accumulate over time, allowing each new generation of agents to inherit and extend the knowledge acquired by previous ones, rather than starting from scratch.
The data efficiency demonstrated by Orchard-GUI suggests that larger-scale web agents could be trained without requiring large amounts of manually created training data. By releasing the complete Orchard stack, including the environment service, training pipelines, and training datasets, we hope to help the broader research community build more capable open agents more

[truncated]
