---
source: "https://www.databricks.com/blog/evaluating-ai-agents-live-grounded-reasoning-cup"
hn_url: "https://news.ycombinator.com/item?id=49376261"
title: "Evaluating AI Agents Live at the Grounded Reasoning Cup"
article_title: "Evaluating AI Agents Live at the Grounded Reasoning Cup | Databricks Blog"
image: "https://www.databricks.com/sites/default/files/2026-08/2026-08-Blog-OfficeQA-Pro-V2-Evaluating-AI-Agents-Live-at-the-Grounded-Reasoning-Cup-OG-1200x628-2x.png"
author: "iwhalen"
captured_at: "2026-08-20T16:23:04Z"
capture_tool: "hn-digest"
hn_id: 49376261
score: 2
comments: 0
posted_at: "2026-08-20T15:46:09Z"
tags:
  - hacker-news
  - translated
---

# Evaluating AI Agents Live at the Grounded Reasoning Cup

- HN: [49376261](https://news.ycombinator.com/item?id=49376261)
- Source: [www.databricks.com](https://www.databricks.com/blog/evaluating-ai-agents-live-grounded-reasoning-cup)
- Score: 2
- Comments: 0
- Posted: 2026-08-20T15:46:09Z

## Translation

タイトル: Ground Reasoning Cup での AI エージェントのライブ評価
記事のタイトル: Ground Reasoning Cup での AI エージェントの評価 |データブリックのブログ
説明: 有力な学術チームが、新しくリリースされたエンタープライズ向けの根拠のある推論ベンチマークに基づいて AI エージェントを評価した、グラウンデッド リーズニング カップからの教訓。

記事本文:
Ground Reasoning Cup での AI エージェントのライブ評価 | Databricks ブログ メイン コンテンツにスキップ ログイン Databricks がアプリ開発者向けに発見する理由
パートナー パートナーの概要 Databricks パートナー エコシステムを探索する
パートナー スポットライト 注目のパートナーの発表
パートナー プログラム メリット、レベル、パートナーになる方法を確認する
AWS、Azure、GCP 上のクラウド プロバイダー Databricks
パートナーを探す ニーズに合った Databricks パートナーを見つけてください
パートナー ソリューション カスタム業界および移行ソリューションを見つける
製品 Databricks プラットフォーム プラットフォームの概要 データ、分析、AI のための統合プラットフォーム
データ エンジニアリング ETL とバッチおよびストリーミング データのオーケストレーション
ビジネス ユーザー向け AI Assistant Agentic コワーカー
データ ウェアハウス SQL 分析用のサーバーレス データ ウェアハウス
アプリケーション開発 安全なデータと AI アプリを迅速に構築
データ アプリと AI エージェント用のデータベース Postgres
人工知能 ML および GenAI アプリケーションを構築してデプロイする
ガバナンス すべてのデータ、分析、AI 資産に対する統合ガバナンス
ビジネス インテリジェンス 実世界のデータに対するインテリジェントな分析
AI 時代に向けて構築されたセキュリティ オープン エージェント SIEM
Databricks に組み込まれた顧客データ プラットフォーム Agentic CDP
共有 データ、分析、AI のためのオープンデータ共有
統合とデータ マーケットプレイス データ、分析、AI のオープン マーケットプレイス
IDE 統合 お気に入りの IDE で Lakehouse を構築
Partner Connect Databricks エコシステムを発見して統合する
価格設定 Databricks の価格設定 製品の価格設定、DBU などを調べる
コスト計算ツール あらゆるクラウド上のコンピューティング コストを見積もります
オープンソース オープンソース テクノロジー プラットフォームの背後にあるイノベーションについて詳しく知る
産業向けソリューション Databricks 電気通信
異業種ソリューション AI エージェント
移行と展開 データの移行
それで

ソリューション アクセラレータ アクセラレータを探索する 重要な結果に向けてより迅速に行動する
リソース 学習トレーニング ニーズに合わせたカリキュラムを見つける
Databricks アカデミー Databricks 学習プラットフォームにサインインする
認証取得による認知と差別化
無料版 プロフェッショナルなデータと AI ツールを無料で学習
University Alliance Databricks を教えたいですか?その方法をご覧ください。
ブログとポッドキャスト Databricks ブログ ニュース、製品発表などをご覧ください
AI ブログ AI の研究とエンジニアリングの取り組みを詳しく見る
Data Brew ポッドキャスト データについて話しましょう!
Champions of Data + AI ポッドキャスト イノベーションを推進するデータ リーダーからの洞察
安心と信頼 安心と信頼
Ground Reasoning Cup での AI エージェントのライブ評価
一流の学術チームが、実際の競争条件下で新しい企業文書コーパスに一般化する AI エージェントをどのように構築したか。
Databricks AI 研究チームによる
Grounded Reasoning Cup では、11 の学術チームに対し、OfficeQA Pro で開発されたエージェントを、約 120,000 ページの米国財務省文書から構築された新しくリリースされたベンチマークである OfficeQA Pro V2 に適用するよう挑戦しました。
結果は、一般化を仮定できないことを示しました。使い慣れたベンチマークに基づいて開発されたアプローチは、常に新しいコーパスに確実に移行するとは限らず、すぐに使用できるフロンティア エージェントの精度は平均 30% 未満でした。
スタンフォード大学の優勝チームは、再利用可能なスキルのライブラリ、ターゲットを絞った文書表現のフォールバック、適応型検証を組み合わせたエンドツーエンドのエージェント最適化戦略により、63.3% の精度を達成しました。
今年、Databricks は、複雑なエンタープライズ スタイルのドキュメント コレクションに対する AI エージェントの推論能力を評価する、この種初のライブ AI コンテストである第 1 回 Ground Reasoning Cup を主催しました。ライブ競争条件下で新しくリリースされたコーパスでエージェントをテストすることによって

つまり、グラウンデッド リーズニング カップは、AI 評価における最も難しい質問の 1 つである、ベンチマークでのパフォーマンスの向上が、同様の現実世界のタスクにどの程度一般化されるのかという質問に答えるのに役立つように設計されました。
このコンテストには、米国とカナダ全土から 11 のトップ学術チームが集まり、OpenAI、Anthropic、Google DeepMind などのフロンティア ラボからのリソースと指導を受けました。 2 か月にわたって、チームは経済的に価値のあるエンタープライズ ワークフローを反映するように設計された主力の根拠に基づいたベンチマークである OfficeQA に基づいてエージェントを開発し、最適化しました。競技会当日、彼らは、改善が一般化するかどうかをテストするために設計された、新しくリリースされた根拠のある推論ベンチマークである OfficeQA Pro V2 にこれらのシステムをリアルタイムで適用するという課題に挑戦しました。
スタンフォード大学は、63.3% の精度を達成したシステムで勝利し、平均的なチームを約 +22 ポイント上回り、平均的なフロンティア エージェントのオフライン ベースラインを約 +35 ポイント上回りました。トップチームは、文書の前処理、対象を絞った検索、並列エージェント、構造化されたツールの使用、検証を通じて大幅な成果を上げたことを実証しました。同時に、18.8% の質問がどのチームでも解決されず、企業に根拠のある推論にどれだけの余地が残っているかを浮き彫りにしました。
このブログ投稿では、コンテストを総括し、Grounded Reasoning Cup の優勝チームであるスタンフォード大学、マサチューセッツ大学アマースト校、エール大学からのエージェント最適化戦略と洞察について説明します。
一般に、次のことがわかります。
一般化には、代表的な、継続的な評価が必要です。 OfficeQA で開発された技術が、常に新しいベンチマークに確実に移行できるわけではありません。これは、OfficeQA Pro V2 のような既存のテスト セットを利用して、ソリューションを新しい例に確実に一般化することの重要性を強調しています。
銀

ent のパフォーマンスは、モデルだけでなくシステム全体に依存します。同じモデルを使用した最高得点チームと最低得点チームの平均差は 30.4 ポイントでした。解析、取得、ツールの使用、検証、並列処理、および運用インフラストラクチャはすべて、エージェントがエンドツーエンドの根拠に基づいた推論タスクを正常に完了できるかどうかに違いをもたらしました。
企業に根ざした推論は依然として解決には程遠いです。勝利したチームでさえ、ベンチマークの検索、解析、分析の要求の多くに苦労しており、継続的な研究と改善の余地が大きく残されていました。実務者には、この作業を継続するために、公開されている OfficeQA ベンチマーク スイートを使用することをお勧めします。
Ground Reasoning Cup の目標は、一流の学術チームを結集して、根拠のある推論への一般化可能なアプローチを開発することでした。これは、大規模な、多くの場合専有文書コレクションからの証拠を使用して複雑な質問に答える必要がある、企業環境における一般的なタスクです。
学術機関を代表する 2 ～ 4 人のチームが、OpenAI、Anthropic、または Google DeepMind の業界パートナーとペアになり、開発期間を通じてモデルへのアクセスと指導を提供しました。チームは、適切と思われるアプローチを使用してエージェントを構築するのに約 2 か月かかりました。ただし、エージェントを強化するためにパートナー ラボのモデル ファミリーのみを使用しなければならないという 1 つの制約がありました。この期間中、彼らは OfficeQA ベンチマークを使用して、同様の根拠のある推論タスクに一般化できると考えられる新しい手法を評価しました。
競技当日、各チームは、リリースされたばかりの新しいベンチマークにエージェントをリアルタイムで適用するという課題を課されました。コンテストは次のルールに従って行われました。
ベンチマーク リリース: 新しいコーパス (米国財務省

ry の受入および支出の報告書）は、競技会のわずか 36 時間前に公開されました。これにより、チームは、メソッドが新しいベンチマークにオーバーフィットする機会を制限しながら、データを処理してインデックスを作成する時間を確保できました。
設計上の制約: チームは、割り当てられた業界ラボ パートナーからのモデルを使用する限り、任意のエージェント フレームワーク、コーパス バージョン、取得戦略、ツール使用セットアップ、または人間参加型ワークフローを使用できます。
形式: コンテストは 15 分間の 6 つのラウンドで構成され、各ラウンドに 15 問が出題されます。イベントが進むにつれて、ラウンドは徐々に難しくなっていきました。
採点: チームは正解ごとに 1 ポイントを獲得しました。低遅延を奨励するために、与えられた質問に最初に正解した場合には、0.25 ポイントのスピード ボーナスも与えられました。最も難しい問題で構成された最終ラウンドでは、ポイントが 2 倍になりました。また、各チームはコンテスト期間中 3 回の再提出が許可されており、以前の解答を修正するために選択して適用することができました。
このコンテストで明らかになったことが 1 つあります。7 か月前に OfficeQA ベンチマークをリリースして以来、エンタープライズ スタイルのドキュメント コーパスに対する根拠のある推論は改善されましたが、まだ解決には程遠いということです。チームの平均スコアは約 41% でしたが、上位 3 チームの精度は 50% を超え、スタンフォード大学チームが 63.3% の精度で優勝しました。これらの結果は、トップチームの素晴らしい仕事と、探求すべき余地がまだ十分に残っていることを示しています。
それぞれが独自のアプローチを採用していましたが、上位 3 チームが構築したエージェント全体でいくつかのパターンが現れました。強力なシステムは、慎重な文書の前処理、対象を絞った検索、構造化されたツールの使用、および回答の検証ステップを組み合わせる傾向がありました。多くの場合、パフォーマンスは単一のモデル呼び出しには依存せず、モデル呼び出しに依存します。

周囲のシステム: 文書がどのように解析され、証拠がどのように取得され、中間計算がどのように実行され、提出前に回答がどのように検証されたか。これらは最もパフォーマンスの高いシステムに一般的に共通する特性でしたが、以下に説明するように、それぞれが異なる創造的な戦略も採用していました。
1位：スタンフォード大学
スタンフォード大学チームの勝利のアプローチは、一般的な根拠に基づいた推論の失敗モードを、Claude Opus 4.8 Claude Code エージェントが競争の質問から学び、適用できるように再利用可能な操作手順に変えることから生まれました。 Opus 4.8 や Fable 5 (利用可能な場合) を使用した実験も含め、公開されている OfficeQA ベンチマークの開発中に、チームは間違った回答をエージェントの正確な間違いにまで遡って繰り返し追跡し、そのパターンを表のローカリゼーション、回答の書式設定、一般的な財務用語の明確化などのスキルに変えました。チームはまた、解析されたコーパス テキストとマークダウン スタイルの文書表現をいつ検索するか、いつソースにフォールバックするかを決定するスキルも統合しました。解析されたテキストに完全なコンテキストが欠けている場合は PDF。スタンフォード大学は、競技会当日までに、エージェントに約 100 以上のスキルを備えたプレイブックを準備していました。彼らは、88 回の試行中 57 回の正解で、全チームをリードしました。
この入念な準備にも関わらず、スタンフォード大学は大会中に戦略を適応させる必要がありました。最初の 3 ラウンドでは、チームは別の Claude Code エージェントを検証者として使用して、中間値を再抽出し、データ リネージュと処理単位のスケーリングによる修正値の追跡などの一般的な障害モードをチェックし、不一致が特定された場合は計算をパッチしました。しかし、それらのラウンドでスピードボーナスを 2 つだけ獲得した後、スタンフォード大学は追加の検証を削除しました。

最後の3名に合格します。この変更により待ち時間が大幅に短縮され、チームは 14 の質問でスピード ボーナスを獲得し、復活を促進することができました。それにもかかわらず、スタンフォード大学が最終再提出を通じて回答を修正するために検証者をオンに戻し、最終的に勝利を確定させたとき、最終ラウンドでは検証者が決定的なものであることが判明しました。
UMass チームはスピードに賭けました。彼らは、Claude Opus 4.8 Fast をプライマリ モデルとして使用し、コーパスを前処理して、解析されたドキュメントの迅速な検索とフィルタリングを可能にするメタデータ カタログを作成しました。低遅延を維持しながら回答の品質を向上させるために、各質問に対して 3 人のエージェントを並行して実行し、その後、最終的な Opus 検証コールを実行して最良の回答を選択しました。
この戦略により、UMass は正解の平均提出時間が最速の 4 分となり、チーム平均の 8 分 30 秒の半分未満となりました。その結果、最初に正解したチームとして、彼らは 36 回のスピード ボーナス (それぞれ 0.25 ポイント相当) を獲得しました。これは、スタンフォード大学の 2 位合計 16 回の 2 倍以上でした。これらのボーナスのおかげで、ハーフタイムまでにスタンフォード大学に対して 10.25 ポイントのリードを築き、3.75 ポイントのアドバンテージを維持して最終ラウンドに臨みました。マサチューセッツ大学は試合終了の最後の56秒までリードを保った。

[切り捨てられた]

## Original Extract

Lessons from the Grounded Reasoning Cup, where leading academic teams evaluated AI agents live on a newly released enterprise grounded-reasoning benchmark.

Evaluating AI Agents Live at the Grounded Reasoning Cup | Databricks Blog Skip to main content Login Why Databricks Discover For App Developers
Partners Partner Overview Explore the Databricks partner ecosystem
Partner Spotlight Featured partner announcements
Partner Program Explore benefits, tiers and how to become a partner
Cloud Providers Databricks on AWS, Azure and GCP
Find a Partner Discover Databricks partners for your needs
Partner Solutions Find custom industry and migration solutions
Product Databricks Platform Platform Overview A unified platform for data, analytics and AI
Data Engineering ETL and orchestration for batch and streaming data
AI Assistant Agentic coworker for business users
Data Warehousing Serverless data warehouse for SQL analytics
Application Development Quickly build secure data and AI apps
Database Postgres for data apps and AI agents
Artificial Intelligence Build and deploy ML and GenAI applications
Governance Unified governance for all data, analytics and AI assets
Business Intelligence Intelligent analytics for real-world data
Security Open agentic SIEM built for the AI era
Customer Data Platform Agentic CDP embedded in Databricks
Sharing Open data sharing for data, analytics and AI
Integrations and Data Marketplace Open marketplace for data, analytics and AI
IDE Integrations Build on the Lakehouse in your favorite IDE
Partner Connect Discover and integrate with the Databricks ecosystem
Pricing Databricks Pricing Explore product pricing, DBUs and more
Cost Calculator Estimate your compute costs on any cloud
Open Source Open Source Technologies Learn more about the innovations behind the platform
Solutions Databricks for Industries Telecommunications
Cross Industry Solutions AI Agents
Migration & Deployment Data Migration
Solution Accelerators Explore Accelerators Move faster toward outcomes that matter
Resources Learning Training Discover curriculum tailored to your needs
Databricks Academy Sign in to the Databricks learning platform
Certification Gain recognition and differentiation
Free Edition Learn professional Data and AI tools for free
University Alliance Want to teach Databricks? See how.
Blog and Podcasts Databricks Blog Explore news, product announcements, and more
AI Blog Explore our AI research and engineering work
Data Brew Podcast Let’s talk data!
Champions of Data + AI Podcast Insights from data leaders powering innovation
Security and Trust Security and Trust
Evaluating AI Agents Live at the Grounded Reasoning Cup
How top academic teams built AI agents that generalized to a new enterprise document corpus under live competition conditions.
by Databricks AI Research Team
The Grounded Reasoning Cup challenged 11 academic teams to apply agents developed on OfficeQA Pro to OfficeQA Pro V2, a newly released benchmark built from approximately 120,000 pages of U.S. Treasury documents.
Results showed that generalization cannot be assumed. Approaches developed on a familiar benchmark did not always transfer reliably to a new corpus, and out-of-the-box frontier agents averaged less than 30% accuracy.
Stanford’s winning team achieved 63.3% accuracy through an end-to-end agent optimization strategy that combined a library of reusable skills, targeted document-representation fallbacks, and adaptive verification.
This year, Databricks hosted the inaugural Grounded Reasoning Cup, a first-of-its-kind live AI competition to evaluate AI agents’ ability to reason over complex, enterprise-style document collections. By testing agents on a newly released corpus under live competition conditions, the Grounded Reasoning Cup was designed to help answer one of the hardest questions in AI evaluation: how well do performance improvements on a benchmark generalize to similar, real-world tasks?
The competition brought together 11 top academic teams from across the U.S. and Canada, paired with resources and mentorship from frontier labs including OpenAI, Anthropic, and Google DeepMind. Over the course of two months, teams developed and optimized their agents on OfficeQA , our flagship grounded-reasoning benchmark designed to reflect economically valuable enterprise workflows. On competition day, they were challenged to apply those systems in real time to a newly released grounded-reasoning benchmark, OfficeQA Pro V2 , designed to test whether their improvements generalized.
Stanford won with a system that achieved 63.3% accuracy, beating out the average team by approximately +22 points, and the average frontier agent offline baseline by approximately +35 points. The top teams demonstrated substantial gains through document preprocessing, targeted retrieval, parallel agents, structured tool use, and verification. At the same time, 18.8% of questions went unsolved by every team, underscoring how much headroom remains in enterprise grounded reasoning.
In this blog post, we recap the competition and discuss agent optimization strategies and insights from the Grounded Reasoning Cup’s winning teams: Stanford, University of Massachusetts Amherst, and Yale.
In general, we find the following:
Generalization requires representative, held-out evaluations. Techniques developed on OfficeQA did not always transfer reliably to our new benchmark. This stresses the importance of utilizing held out test sets like OfficeQA Pro V2 to ensure solutions generalize to new examples.
Agent performance depends on the full system, not only the model. The average gap between the top scoring and lowest scoring teams using the same model was 30.4 points. Parsing, retrieval, tool use, verification, parallelism, and operational infrastructure all made the difference in whether agents could successfully complete end-to-end grounded reasoning tasks.
Enterprise grounded reasoning remains far from solved. Even winning teams struggled with many of the benchmark’s retrieval, parsing, and analytical demands, leaving substantial room for continued research and improvement. We encourage practitioners to use the publicly available OfficeQA benchmark suite to continue advancing this work.
The goal of the Grounded Reasoning Cup was to bring together top academic teams to develop generalizable approaches to grounded reasoning – a common task in enterprise settings that involves answering complex questions using evidence from large, often proprietary, document collections.
Teams of 2-4 people representing their academic institution were paired with an industry partner from OpenAI, Anthropic, or Google DeepMind, who provided access to their models and mentorship throughout the development period. Teams had approximately two months to build an agent using any approaches they saw fit, with the one constraint that they must use their partner lab’s model family exclusively to power their agent. During this period, they used the OfficeQA benchmark to evaluate new techniques they believed would generalize to similar grounded reasoning tasks.
On the day of the competition, teams were then tasked with applying their agents in real time on a new, freshly released benchmark. The competition was governed by the following rules:
Benchmark release: The new corpus (the U.S. Treasury’s Accounts of Receipts and Expenditures) was released just 36 hours before the competition. This gave teams time to process and index the data while limiting the opportunity for methods to overfit to the new benchmark.
Design Constraints: Teams could use any agent framework, corpus version(s), retrieval strategy, tool-use setup, or human-in-the-loop workflow, so long as they used models from their assigned industry lab partner.
Format: The competition consisted of six 15-minute rounds with 15 questions per round. Rounds became progressively harder over the course of the event.
Scoring: Teams received 1 point for each correct answer. To incentivize low latency, they were also awarded a 0.25-point speed bonus if they were the first to answer a given question correctly. Points were worth 2x points in the final round, which consisted of the most challenging questions. Each team was also allowed 3 resubmissions throughout the competition, which they could choose to apply to correct a previous answer.
The competition made one thing clear: grounded reasoning over enterprise-style document corpora has improved since we released the OfficeQA benchmark 7 months ago, but it is still far from solved. The average team score was ~41%, while the top three teams exceeded 50% accuracy with the Stanford team winning the competition with 63.3% accuracy. These results point towards impressive work from top teams, as well as plenty of remaining headroom to explore.
While each took a unique approach, several patterns emerged across the agents built by the top three teams. Strong systems tended to combine careful document preprocessing, targeted retrieval, structured tool use, and answer verification steps. In many cases, performance depended less on a single model call and more on the surrounding system: how documents were parsed, how evidence was retrieved, how intermediate calculations were performed, and how answers were verified before submission. While these were qualities the most performant systems generally had in common, they also each employed distinct and creative strategies, as described below.
1st Place: Stanford University
The Stanford team’s winning approach came from turning common grounded reasoning failure modes into reusable operating procedures for their Claude Opus 4.8 Claude Code agent to learn from and apply on the competition questions. During their development on the public OfficeQA benchmark, including experiments ablating with Opus 4.8 and even Fable 5 (while available), the team repeatedly traced wrong answers back to the agent’s exact misstep, and then turned those patterns into skills for table localization, answer formatting, clarifications on common financial wordings, etc. The team also integrated skills that would decide when to search across parsed corpus text and markdown-style document representations and when to fall back to the source PDFs if the parsed text lacked full context. By competition day, Stanford had prepared their agent with a playbook of ~100+ skills. They led all teams in accuracy in the 57 correct answers of 88 they attempted.
Despite this in-depth preparation, Stanford still had to adapt its strategy during the competition. In the first three rounds, the team used another Claude Code agent as a verifier to re-extract intermediate values, check common failure modes such as tracing revised values through data lineage and handling unit scaling, and patch calculations when it identified discrepancies. But after earning only two speed bonuses in those rounds, Stanford removed the extra verification pass for the final three. The change significantly reduced latency, helping the team earn speed bonuses on 14 questions and fuel its comeback. The verifier nevertheless proved decisive in the final round, when Stanford switched it back on to correct an answer through a final resubmission, ultimately securing their win.
The UMass team made a bet on speed. They used Claude Opus 4.8 Fast as their primary model and preprocessed the corpus to create a metadata catalog that enabled rapid search and filtering over parsed documents. To improve answer quality while preserving low latency, they ran three agents in parallel on each question, followed by a final Opus verification call to select the best answer.
This strategy gave UMass the fastest average submission time for correct answers: four minutes, less than half the team average of eight minutes and 30 seconds. As a result, they earned 36 speed bonuses, worth 0.25 points each, for being the first team to answer correctly, more than twice Stanford’s second-place total of 16. These bonuses helped them build a 10.25-point lead over Stanford by halftime and preserve a 3.75-point advantage entering the final round. UMass held that lead until the final 56 seconds of the com

[truncated]
