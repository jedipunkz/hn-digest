---
source: "https://aws.amazon.com/blogs/machine-learning/evaluating-ai-agents-a-production-blueprint-with-strands-and-agentcore/"
hn_url: "https://news.ycombinator.com/item?id=49027991"
title: "Evaluating AI Agents: A Production Blueprint with Strands and AgentCore"
article_title: "Evaluating AI Agents: A production blueprint with Strands and AgentCore | Artificial Intelligence"
author: "Imbiss"
captured_at: "2026-07-23T21:58:25Z"
capture_tool: "hn-digest"
hn_id: 49027991
score: 1
comments: 0
posted_at: "2026-07-23T20:59:57Z"
tags:
  - hacker-news
  - translated
---

# Evaluating AI Agents: A Production Blueprint with Strands and AgentCore

- HN: [49027991](https://news.ycombinator.com/item?id=49027991)
- Source: [aws.amazon.com](https://aws.amazon.com/blogs/machine-learning/evaluating-ai-agents-a-production-blueprint-with-strands-and-agentcore/)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T20:59:57Z

## Translation

タイトル: AI エージェントの評価: Strands と AgentCore を使用した運用ブループリント
記事のタイトル: AI エージェントの評価: Strands と AgentCore を使用した運用ブループリント |人工知能
説明: Motorway と AWS は協力して、エンドツーエンドの評価パイプラインを構築しました。これにより、誤った結果が 8 クエリに 1 から 50 に 1 に減少し、問題検出時間が数時間から数分に短縮されました。このパイプラインは、Strands Agents SDK と、デプロイ用のフルマネージド サービスである Amazon Bedrock AgentCore を組み合わせたものです。
[切り捨てられた]

記事本文:
データ移行
オンプレミスのデータベースはエージェント AI 用に構築されたものではありません。アマゾンRDSは
独立系ソフトウェアベンダー
ソフトウェアとテクノロジーのための AI: 構築、拡張、収益化
中小企業向けの AI
お客様のビジネス向けに設計されたソリューション。 AWS パートナーによって提供される
人工知能 (AI)
AWS で AI の実験から実稼働までを加速する
AIエージェント
AgentCore: エージェントを構築、接続、最適化するための 1 つのプラットフォーム
ニュースとお知らせ
AWS ブログ
AWSについて
AWS は世界で最も包括的なクラウドであり、組織がイノベーションを加速し、コストを削減し、より効率的に拡張できるようにします。
これらの注目のサービスのいずれかを開始するか、すべてを参照してください
トランスフォーム
エージェント AI で技術的負債を排除し、レガシー システムとコードを最新化します
オーロラ
PostgreSQL、MySQL、DSQL 用のサーバーレス リレーショナル データベース サービス
アマゾンの岩盤
生成 AI アプリケーションとエージェントを構築するためのエンドツーエンドのプラットフォーム
Amazon Connectの顧客
顧客とのやり取り全体で優れたエクスペリエンスを提供する AI ネイティブ ソリューション
EC2
事実上あらゆるワークロードに対応する安全でサイズ変更可能なコンピューティング容量
ノヴァ
最先端のインテリジェンスと最高のコストパフォーマンスを実現するファウンデーションモデル
オープンサーチサービス
リアルタイム検索、可観測性、セキュリティ分析
S3
AI、分析、アーカイブ用の事実上無制限の安全なオブジェクト ストレージ
新機能
新しい機能と最新のテクノロジーを探索する
お客様の成功事例
世界中の顧客がクラウドをどのように加速しているかをご覧ください
業界別
AWS クラウド ソリューションで業界を変革します。実績のあるアーキテクチャ、コンプライアンスガイド、お客様のセクターに合わせた AWS 製品を使用したお客様の成功事例にアクセスしてください。あなたの業界向けの AWS ソリューションを検討してください。
航空宇宙および衛星
顧客の衛星構築、宇宙飛行、打ち上げを支援するクラウド ソリューション

運用と宇宙探査の再考
自動車
クラウド ソリューションでよりスマートな車両を構築し、モビリティを変革する
教育
教育、学習、生徒の参加を促進し、学習成果を向上させ、IT 運用を最新化するのに役立つソリューション
エネルギーと公共事業
従来の業務を刷新し、革新的な再生可能エネルギー ビジネス モデルの開発を加速する
金融サービス
銀行業務、資本市場、保険、決済にわたる革新的で安全なソリューションを開発する
ゲーム
AAA タイトルからインディー スタジオまで、あらゆるジャンルとプラットフォームのゲーム開発を可能にします
政府
政府機関の最新化、義務への対応、コスト削減、ミッション成果の達成を支援するために設計されたソリューション
ヘルスケアとライフサイエンス
医療データの管理とセキュリティによりイノベーションを加速し、患者ケアを改善します
産業用
製造、自動車、エネルギー、電力と公共事業、輸送と物流にわたる顧客向けのサービスとソリューション
製造業
生産を最適化し、市場投入までの時間を短縮します
メディアとエンターテイメント
あらゆるクラウドの中で最も特化した機能とパートナー ソリューションでメディアとエンターテイメントを変革します
非営利
非営利団体や NGO がミッションへの影響力と寄付者との関係を最大化できるよう支援するクラウド ソリューション
小売および消費財
クラウド ソリューションでブランドが市場の差別化と成長を推進できるようにする
半導体
クラウドテクノロジーを使用した次世代チップソリューションの設計と提供
スポーツ
革新的なファン、放送、アスリートのエクスペリエンスを促進する
持続可能性
組織が AWS を使用して持続可能性の目標を達成できるように支援するツールとリソース
電気通信
クラウドベースの通信ソリューションでイノベーションを加速し、自信を持って拡張し、俊敏性を高めます
旅行とホスピタリティ
旅行とホスピタリティを向上させる

顧客体験と業務効率を向上させるソリューション
新機能
新しい機能を探索する
[切り捨てられた]
AI エージェントの評価: Strands と AgentCore を使用した実稼働ブループリント
この投稿は、Motorway と AWS プロトタイピングおよび AI カスタマー エンジニアリング (PACE) チームとの共同執筆です。
英国に本拠を置くオンライン自動車市場であるモーターウェイでは、毎日オークションを開催しており、最大 8,000 のディーラーが最大 2,500 台の車両に入札しています。 Motorway は、AWS プロトタイピングおよび AI カスタマー エンジニアリング (PACE) と協力して、ディーラーによる車両の検索方法を変革する AI を活用したディーラー在庫検索エージェントを構築し、何時間にもわたる手動フィルタリングを自然言語クエリに置き換えました。
エージェントは自信に満ちたような返答をしますが、実際のお金がかかっている状態でそれが確実に機能することをどうやって証明するのでしょうか?
ツールの選択ミスは間違った検索結果を引き起こし、ディーラーの信頼を損ないます。
セマンティック検索の誤解により、無関係な結果が返されます。 「5 年前までのガソリン車、ハイブリッド車、電気自動車」のようなクエリでは、エージェントが複数の制約を正しく解析する必要があります。
マルチターン会話におけるコンテキストのドリフトにより、ディーラーの調整が失われます。
出力が非決定的であるため、単回試行テストの信頼性が低くなります。
Motorway と AWS は協力して、エンドツーエンドの評価パイプラインを構築しました。これにより、誤った結果が 8 クエリに 1 から 50 に 1 に減少し、問題検出時間が数時間から数分に短縮されました。
このパイプラインは、Strands Agents SDK と、AI エージェントを大規模にデプロイおよび運用するためのフルマネージド サービスである Amazon Bedrock AgentCore を組み合わせています。この投稿では、独自のエージェント用にこのパイプラインを構築する方法を学びます。
structs-agents-evals (Strands Agents のオープンソース評価ライブラリ) を使用したビルド時テストと、Amazon Bedrock AgentCore Evaluation を使用した本番モニタリングに及ぶ 2 フェーズの評価戦略。
3層構造のフレームワーク

ツールの使用法、推論、出力品質を評価します。
メトリクスがしきい値を下回った場合にリリースをブロックする品質ゲートを備えた 5 段階のデプロイメント パイプライン。
コンパニオン リポジトリは、独自のエージェントに適応できる展開可能なブループリントを提供します。ブループリントでは AWS のサービスを使用していますが、中心となる原則は、本番環境に対応した AI エージェントにとって不可欠であり、システムに依存しない要件です。これらの原則には、3 層の評価フレームワークと一貫性のための pass^k メトリックの使用が含まれます。
この投稿を進めるには、次の前提条件を満たしている必要があります。
Amazon Bedrock 、 AWS Lambda 、 Amazon Simple Storage Service (Amazon S3) 、 Amazon DynamoDB 、 Amazon EventBridge 、 Amazon CloudWatch 、および Amazon Simple Notice Service (Amazon SNS) のアクセス許可を持つ AWS アカウント。
AWS Cloud Development Kit (AWS CDK) v2 がインストールされ、ブートストラップされています。
Amazon Bedrock モデル アクセスを通じて、Anthropic Claude モデルと Amazon Titan モデルにアクセスします。
Python、AWS CDK、エージェントの概念 (ツール呼び出し、複数ターンの会話) に精通していること。
完了までの時間: 初期展開に 30 ～ 45 分、ドメインに合わせてカスタマイズするには 2 ～ 3 時間。
推定コスト: サンプル評価スイートの実行には、Amazon Bedrock 推論料金として約 5 ～ 10 ドルかかります。生産監視のコストはサンプリング レートによって異なります。
セキュリティ上の注意: コンパニオン リポジトリは、最小特権の AWS Identity and Access Management (IAM) ロールを実装し、API キーを AWS Systems Manager パラメータ ストア (環境変数ではなく) に保存し、型指定されたパラメータを使用してインジェクション攻撃の防止に役立ちます。詳細については、リポジトリの README を参照してください。
事例：ディーラー在庫検索エージェント
Motorway は、Strands Agents SDK と Amazon Bedrock AgentCore に基づいてディーラー在庫検索エージェントを構築しました。 T

このエージェントは、89 を超える車両属性にわたる構造化フィルタリングと、LanceDB および Amazon Titan Text Embeddings V2 を利用したベクトル類似性検索を組み合わせた 8 つのツールを公開しています。
ディーラーは通常、CSV と厳格なフィルターを使用してリストを閲覧するのに何時間も費やしていました。会話型 AI エージェントを導入することで、ディーラーはエージェントに「ディーラーの近くで 25,000 未満のディーゼル SUV を探して」または「家族向けのスポーティでオートマチックの車を探して」と会話できるようになります。
ピーク時には約 1,500 人の同時ユーザーが存在するため、エージェントの動作を適切にすることは必須ではありません。ツールの選択エラーやセマンティック検索の誤解は、ユーザーの信頼に直接影響します。図 1 は、エンドツーエンドのリクエスト フローを示しています。ディーラーは Web インターフェイスを通じて自然言語クエリを送信し、Amazon Bedrock AgentCore Runtime にルーティングします。ランタイムは、Amazon Bedrock モデル (推論には Claude、埋め込みには Amazon Titan) を使用しながら、8 つの異なるツールへの呼び出しを調整します。ツールの応答はランタイムを介して逆流し、ディーラー向けの最終結果を生成します。
エージェントの評価が異なる理由
大規模言語モデル (LLM) の評価は、一貫性、事実の正確さ、応答の関連性といったテキスト生成の品質に焦点を当てています。エージェントの評価では、根本的に異なるものを評価します。次のように考えてください。LLM 評価はエンジンの性能を検査します。エージェントの評価では、渋滞中、雨天、または後部座席に乗客がいっぱいの状態で、車全体がどのように運転するかを評価します。
従来の LLM 指標では、高速道路の担当者が「グレード 1 スズキ モデル」の適切な検索ツールに電話をかけたかどうかはわかりません。エージェントが正しいフィルター パラメーターを LanceDB に渡したかどうかは明らかにされません。そして、ディーラーが前のターンの結果を精錬したときに正しい反応が得られるかどうかも見逃してしまいます。
「フォルクスワーゲン ゴルフ 7 ～ 12 年前」のような正確なクエリは、

これは問題ありませんが、セマンティック検索レイヤーが適切に評価されていない場合、「古い VW を探しています」のような口語的な表現は失敗する可能性があります。
structs-agents-evals を使用してデプロイメント前に問題を検出する
ブループリントは、GenAIOps ライフサイクルに対応する 2 つのフェーズにわたる評価を実装します。ビルド時の評価はデプロイ前に問題を検出し、実稼働評価は総合テストで見逃したものを検出します。次の図 (図 2) は、ツールの使用法、推論、出力品質の各レイヤーが展開前にどのように通過する必要があるかを示しています。このフレームワークは、次の 3 つの層にわたってエージェントを評価します。
レイヤ 1 (ツールの使用法) は、95% を超えるしきい値で正しいツールの選択とパラメータの受け渡しを検証します。
レイヤ 2 (推論) は、85% を超えるしきい値で論理的な意思決定を評価します。
レイヤ 3 (出力品質) は、90% を超えるしきい値で応答の有用性と精度を測定します。導入を続行する前に、3 つの層すべてを通過する必要があります。
開発、継続的インテグレーション、継続的デプロイ (CI/CD) 中に、パイプラインは structs-agents-evals フレームワークを使用して、デプロイ前に問題を検出します。出力の検証、軌道評価、マルチターン会話シミュレーション、および自動実験生成を提供します。それぞれは、Strands Agents SDK 上に構築されたエージェントとネイティブに動作するように設計されています。フレームワークは 3 つのプリミティブを提供します。
実験 : エージェントに対して実行されるテスト ケースのコレクション。
ケース : 入力クエリ、期待される出力、および期待されるツール軌跡。
評価者 : スコアリング ロジック (決定論的または LLM ベース)。
テストをレイヤーに構造化します。レイヤ 1 は、ツール選択の精度を高めるために、決定論的なコードベースのグレーダーを実行します。レイヤ 2 とレイヤ 3 は、推論と出力品質のために、LLM-as-judge エバリュエーター (LLM を使用してエージェントの出力をスコアリングする) を使用します。
クスト

m エバリュエーターのサブクラスは、ドメイン固有の問題を処理します。高速道路代理店の場合、これらはデータの鮮度、ディーラーの調査、安全ガードレールをカバーします。エージェントには独自のドメイン制約があります。
評価フレームワークでは 3 つの採点タイプが使用され、それぞれが異なる評価ニーズに適しています。
実際には、エージェントが作成したものを評価する方が、エージェントが通過したパスを評価するよりも多くの問題を発見できます。エージェントがどのツールを最初に呼び出したかではなく、ユーザーが適切な結果を得たかどうかが重要です。
3層の評価フレームワーク
ビルド時の評価は 3 つの異なるレイヤーにわたって実行され、それぞれに特定の合格/不合格しきい値が設定されます。
レイヤ 1: ツールの使用状況 (> 95 パーセントのしきい値)。エージェントは正しいパラメータを使用して適切なツールを呼び出しましたか?
「7,000 ポンドから 20,000 ポンドのディーゼル車」は、型付きフィルターを備えた search_vehicles を使用する必要があります
(燃料タイプ=ディーゼル、最小価格=7000、最大価格=20000)。
「走行距離の少ない最新のハッチバック」は、セマンティック埋め込みと構造化フィルターを組み合わせて、 hybrid_search をトリガーする必要があります。
これは決定論的に測定します。ToolSelectionGrader はどのツールが呼び出されたかを確認し、TrajectoryOrderGrader は呼び出しシーケンスを検証します。
レイヤ 2: 推論 (> 85 パーセントのしきい値)。意思決定プロセスは論理的でしたか? st の HelpfulEvaluator と TrajectoryEvaluator

[切り捨てられた]

## Original Extract

Together, Motorway and AWS built an end-to-end evaluation pipeline that reduced incorrect results from 1 in 8 queries to 1 in 50 and cut issue detection time from few hours to few minutes. The pipeline combines the Strands Agents SDK with Amazon Bedrock AgentCore, a fully managed service for deployi
[truncated]

Data Migration
On-prem databases weren't built for agentic AI. Amazon RDS is
Independent Software Vendors
AI for software & tech: build, scale, and monetize
AI for Small Businesses
Solutions designed for your business. Delivered by AWS Partners
Artificial Intelligence (AI)
Accelerate AI from experimentation to production with AWS
AI agents
AgentCore: One platform to build, connect and optimize agents
News & announcements
AWS blog
About AWS
AWS is the world's most comprehensive cloud, enabling organizations to accelerate innovation, reduce costs, and scale more efficiently
Get started with one of these featured services or browse all
Transform
Eliminate tech debt with agentic AI to modernize legacy systems and code
Aurora
Serverless relational database service for PostgreSQL, MySQL, and DSQL
Amazon Bedrock
The end-to-end platform for building generative AI applications and agents
Amazon Connect Customer
AI-native solution for delivering exceptional experiences across customer interactions
EC2
Secure and resizable compute capacity for virtually any workload
Nova
Foundation models delivering frontier intelligence and top price performance
OpenSearch Service
Real-time search, observability, and security analytics
S3
Virtually unlimited secure object storage for AI, analytics, and archives
What’s new
Explore new capabilities and the latest technologies
Customer success stories
Learn how customers around the world accelerate their cloud
By industry
Transform your industry with AWS cloud solutions. Access proven architectures, compliance guides, and success stories of customers using AWS products tailored to your sector. Explore AWS solutions for your industry.
Aerospace and Satellite
Cloud solutions to help customers build satellites, conduct space and launch operations, and reimagine space exploration
Automotive
Build smarter vehicles and transform mobility with cloud solutions
Education
Solutions to help facilitate teaching, learning, student engagement, better learning outcomes and modernize IT operations
Energy and Utilities
Revamp legacy operations and accelerate the development of innovative renewable energy business models
Financial Services
Develop innovative and secure solutions across banking, capital markets, insurance, and payments
Games
Enable game development for every genre and platform, from AAA titles to indie studios
Government
Solutions designed to help government agencies modernize, meet mandates, reduce costs, and deliver mission outcomes
Healthcare and Life Sciences
Accelerate innovation and improve patient care with healthcare data management and security
Industrial
Services and solutions for customers across Manufacturing, Automotive, Energy, Power & Utilities, Transportation & Logistics
Manufacturing
Optimize production and speed time-to-market
Media and Entertainment
Transform media & entertainment with the most purpose-built capabilities and partner solutions of any cloud
Nonprofit
Cloud solutions to help nonprofits and NGOs maximize their mission impact and donor relationships
Retail and Consumer Goods
Empower brands to drive market differentiation and growth with cloud solutions
Semiconductor
Design and deliver next-generation chip solutions with cloud technology
Sports
Fuel innovative fan, broadcast, and athlete experiences
Sustainability
Tools and resources to help organizations achieve their sustainability goals with AWS
Telecommunications
Accelerate innovation, scale with confidence, and add agility with cloud-based telecom solutions
Travel and Hospitality
Elevate travel and hospitality with solutions that boost customer experiences and operational efficiency
What's new
Explore new capabil
[truncated]
Evaluating AI Agents: A production blueprint with Strands and AgentCore
This post was co-written with Motorway and the AWS Prototyping and AI Customer Engineering (PACE) team.
Motorway, a UK-based online car marketplace, runs a daily auction where up to 8,000 dealers bid on up to 2,500 vehicles. Motorway worked with AWS Prototyping and AI Customer Engineering (PACE) to build an AI-powered dealer stock search agent that transforms how dealers find vehicles, replacing hours of manual filtering with natural language queries.
The agent gives a confident-sounding response, but how do you prove it works reliably with real money on the line?
Tool selection errors cause wrong search results, eroding dealer trust.
Semantic search misinterpretations return irrelevant results. A query like “Petrol, Hybrid and electric cars up to 5 years old” requires the agent to correctly parse multiple constraints.
Context drift in multi-turn conversations loses dealer refinements.
Non-deterministic outputs make single-trial testing unreliable.
Together, Motorway and AWS built an end-to-end evaluation pipeline that reduced incorrect results from 1 in 8 queries to 1 in 50 and cut issue detection time from few hours to few minutes.
The pipeline combines the Strands Agents SDK with Amazon Bedrock AgentCore, a fully managed service for deploying and operating AI agents at scale. In this post, you will learn how to build this pipeline for your own agents:
A two-phase evaluation strategy spanning build-time testing with strands-agents-evals (the open source evaluation library for Strands Agents) and production monitoring with Amazon Bedrock AgentCore Evaluations .
A three-layer framework for assessing tool usage, reasoning, and output quality.
A five-stage deployment pipeline with quality gates that block releases when metrics fall below thresholds.
A companion repository provides a deployable blueprint that you can adapt for your own agents. Although the blueprint uses AWS services, the core principles are essential and system-agnostic requirements for any production-ready AI agent. These principles include the three-layer evaluation framework and the use of the pass^k metric for consistency.
You must have the following prerequisites to follow along with this post.
An AWS account with permissions for Amazon Bedrock , AWS Lambda , Amazon Simple Storage Service (Amazon S3) , Amazon DynamoDB , Amazon EventBridge , Amazon CloudWatch , and Amazon Simple Notification Service (Amazon SNS) .
AWS Cloud Development Kit (AWS CDK) v2 installed and bootstrapped.
Access to Anthropic Claude models and Amazon Titan models through Amazon Bedrock model access .
Familiarity with Python, AWS CDK, and agent concepts (tool calling, multi-turn conversations).
Time to complete: 30–45 minutes for initial deployment and 2–3 hours to customize for your domain.
Estimated costs: Running the sample evaluation suite costs approximately $5–10 in Amazon Bedrock inference charges. Production monitoring costs vary based on sampling rate.
Security note: The companion repository implements least-privilege AWS Identity and Access Management (IAM) roles, stores API keys in AWS Systems Manager Parameter Store (not environment variables), and uses typed parameters to help prevent injection attacks. See the repository README for details.
The worked example: A dealer stock search agent
Motorway built the dealer stock search agent on the Strands Agents SDK and Amazon Bedrock AgentCore . The agent exposes eight tools that combine structured filtering across over 89 vehicle attributes with vector similarity search powered by LanceDB and Amazon Titan Text Embeddings V2 .
Dealers typically spent hours browsing through listings using CSVs and rigid filters. By introducing a conversational AI agent, dealers can now talk to the agent: “Find me diesel SUVs under 25k near my dealership” or “something sporty and automatic for a family.”
With around 1,500 concurrent users during peak hours, getting agent behavior right isn’t optional. A tool selection error or semantic search misinterpretation directly impacts user trust. Figure 1 shows the end-to-end request flow: dealers submit natural language queries through a web interface, which routes to Amazon Bedrock AgentCore Runtime. The runtime orchestrates calls to eight different tools while using Amazon Bedrock models (Claude for reasoning, Amazon Titan for embeddings). Tool responses flow back through the runtime to generate the final dealer-facing results.
Why agent evaluation is different
Large language model (LLM) evaluation focuses on text generation quality: coherence, factual accuracy, response relevance. Agent evaluation assesses something fundamentally different. Think of it this way: LLM evaluation examines engine performance. Agent evaluation assesses how the whole car drives in traffic, in rain, or with a backseat full of passengers.
Traditional LLM metrics don’t tell you whether the Motorway agent called the right search tool for “Grade 1 Suzuki models.” They don’t reveal if the agent passed correct filter parameters to LanceDB. And they miss whether a dealer refining results from a previous turn gets the right response.
A precise query like “Volkswagen Golf 7-12 years old” might work perfectly, but a colloquial variant like “I’m looking for an older VW” could fail if the semantic search layer isn’t properly evaluated.
Catch issues before deployment with strands-agents-evals
The blueprint implements evaluation across two phases that map to the GenAIOps lifecycle . Build-time evaluation catches issues before deployment, and production evaluation catches what synthetic tests miss. The following diagram (Figure 2) shows how the tool usage, reasoning, and output quality layers must pass before deployment. The framework evaluates agents across three layers:
Layer 1 (Tool Usage) validates correct tool selection and parameter passing with a greater than 95 percent threshold.
Layer 2 (Reasoning) assesses logical decision-making with a greater than 85 percent threshold.
Layer 3 (Output Quality) measures response helpfulness and accuracy with a greater than 90 percent threshold. All three layers must pass before deployment proceeds.
During development and continuous integration and continuous deployment (CI/CD), the pipeline uses the strands-agents-evals framework to catch issues before deployment. It provides output validation, trajectory evaluation, multi-turn conversation simulation, and automated experiment generation. Each is designed to work natively with agents built on the Strands Agents SDK. The framework provides three primitives:
Experiment : A collection of test cases run against the agent.
Case : Input query, expected output, and expected tool trajectory.
Evaluator : Scoring logic (deterministic or LLM-based).
Structure your tests into layers. Layer 1 runs deterministic code-based graders for tool selection accuracy. Layers 2 and 3 use LLM-as-judge evaluators (using an LLM to score agent outputs) for reasoning and output quality.
Custom Evaluator subclasses handle domain-specific concerns. For the Motorway agent, these cover data freshness, dealer scoping, and safety guardrails. Your agent will have its own domain constraints.
The evaluation framework uses three grader types, each suited to different evaluation needs.
In practice, grading what the agent produced catches more issues than grading the path it took. You care that the user got relevant results, not which tool the agent called first.
Three-layer assessment framework
Build-time evaluation operates across three distinct layers, each with specific pass/fail thresholds.
Layer 1: Tool Usage (> 95 percent threshold). Did the agent call the right tools with correct parameters?
“Diesel vehicles from £7,000 to £20,000” should use search_vehicles with typed filters
( fuel_type=diesel , min_price=7000 , max_price=20000 ).
“Modern hatchback with low mileage” should trigger hybrid_search , combining semantic embeddings with structured filters.
You measure this deterministically: ToolSelectionGrader checks which tools were called, TrajectoryOrderGrader verifies the call sequence.
Layer 2: Reasoning (> 85 percent threshold). Was the decision-making process logical? The HelpfulnessEvaluator and TrajectoryEvaluator from st

[truncated]
