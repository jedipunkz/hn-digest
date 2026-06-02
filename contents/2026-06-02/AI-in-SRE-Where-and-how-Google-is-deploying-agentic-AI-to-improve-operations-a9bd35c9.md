---
source: "https://cloud.google.com/blog/products/devops-sre/how-google-sre-is-using-agentic-ai-to-improve-operations"
hn_url: "https://news.ycombinator.com/item?id=48360793"
title: "AI in SRE: Where and how Google is deploying agentic AI to improve operations"
article_title: "How Google SRE is using agentic AI to improve operations | Google Cloud Blog"
author: "geoffbp"
captured_at: "2026-06-02T00:34:42Z"
capture_tool: "hn-digest"
hn_id: 48360793
score: 6
comments: 0
posted_at: "2026-06-01T18:33:14Z"
tags:
  - hacker-news
  - translated
---

# AI in SRE: Where and how Google is deploying agentic AI to improve operations

- HN: [48360793](https://news.ycombinator.com/item?id=48360793)
- Source: [cloud.google.com](https://cloud.google.com/blog/products/devops-sre/how-google-sre-is-using-agentic-ai-to-improve-operations)
- Score: 6
- Comments: 0
- Posted: 2026-06-01T18:33:14Z

## Translation

タイトル: SRE における AI: Google は業務改善のためにエージェント AI をどこに、どのように導入しているか
記事のタイトル: Google SRE がエージェント AI を使用して業務を改善する方法 | Google クラウド ブログ
説明: SRE AI により、Google は AI とエージェント テクノロジーを完全に採用し、制御を維持しながら AI を戦力増強手段として活用することを計画しています。

記事本文:
Google SRE がエージェント AI を使用して業務を改善する方法 | Google Cloud ブログ コンテンツに移動 クラウド ブログ 営業担当者へのお問い合わせ 無料で始めましょう クラウド ブログ ソリューションとテクノロジー AI と機械学習
SRE における DevOps と SRE AI: Google は運用を改善するためにエージェント AI をどこに、どのように導入しているか
著名なソフトウェアエンジニア
優秀なサイト信頼性エンジニア
職場における AI への玄関口
20 年以上前の創業以来、Google はサイト信頼性エンジニアリング (SRE) を使用して、検索、Gmail、マップ、YouTube、Google Cloud などのサービスの信頼性と可用性を維持し、信頼性第一の考え方の原則と実践を遵守してきました。
しかし最近、AI の出現により、システムの複雑さは段階的に変化してきました。コンポーネント間の相互作用は、さまざまな要因によりさらに複雑になっています。
マイクロサービス アーキテクチャでは、システムはより広範囲の地理的場所と、ハードウェアの多様性がより優れたデータ センターに分散されます。
エンタープライズ クラウド製品は、信じられないほど複雑な製品セットを備えた広範な機能を提供します。
Google サービスは現在、より独自のビジネス要件と規制要件をカバーしているため、全体的なトポロジと分類がはるかに複雑で理解しにくくなっており、継続的なデプロイメント パイプラインから生じる絶え間ないシステム変更の流れによって課題がさらに増幅されています。
AI コード生成機能により、ソフトウェア開発者は桁違いに多くのコードを提供できるようになり、その結果、信頼性の問題が発生する機会が増加しました。
AI はある意味で SRE チームの作業をより困難なものにしていますが、同時に、本番運用を含むソフトウェア開発ライフサイクルを理解し、改善するための新しい方法も提供します。 Google SRE は AI と人工知能の完全導入に向けて進んでいます

tic テクノロジーは、制御を維持しながら力を増幅する手段として AI を活用します。私たちはこれを SRE AI と呼んでいます。
このトピックを考える際の考慮事項の概要については、この記事を読んでください。また、包括的なホワイトペーパー「AI in SRE Practice: Moving Beyond Automation at Google」を直接読んで、Google SRE が決定論的自動化からエージェント的 AI への移行をどのようにナビゲートしているかを詳しく調べることもできます。
SRE AI の機会の展望
SRE AI 戦略の定義を支援するために、私たちは機会のある分野のソフトウェア開発ライフサイクル (SDLC) 全体を検討しました。
上の図は、SRE が関与する各フェーズを示しており、SRE AI で改善できる可能性があります。
おそらく、エージェント AI から恩恵を受ける可能性がある最も明白な SRE 分野は、従来の SRE 分野の基礎である根本原因分析 (RCA) とも呼ばれる調査と緩和です。しかし、RCA が SRE AI のすべてではありません。 SRE AI に対する当社の計画は、RCA やトラブルシューティングをはるかに超えて、SDLC 全体に対応しています。私たちが取り組んでいる分野をいくつか紹介します。
SRE は、設計、立ち上げ、展開の各フェーズを通じて、信頼性がシステム設計の不可欠な部分であることを保証するために必要なポリシー、ツール、および手順に取り組んできました。エージェント的アプローチは、特にリスクの高いサービスや機能については、プロセスから人を排除することを必ずしも意味するものではありませんが、人によるレビューが必要になる前に多くの問題を検出して自動対処できるため、人が費やす時間が大幅に削減されます。
インシデント時に使用されるランブック (プレイブック) およびその他のドキュメントは、重要な運用成果物です。 Google SRE は、インシデント時の使用状況に基づいてハンドブックと本番ドキュメントを継続的に監視し、改善する AI エージェントを開発しました。 AI エージェントができること

インシデントから新しいプレイブックも生成します。
異常の検出と警告
SRE 実践の中核は、サービス レベル インジケーター (SLI) とサービス レベル目標 (SLO) を定義し、それらのアラートを構成することです。サービスのユースケースがかなり均一で、顧客の期待に沿った目標を定義できる場合、このアプローチは問題ない傾向があります。
ただし、Google Cloud の多くのプロダクトと同様に、さまざまな顧客のユースケースやワークロードをサポートするプロダクトの場合、さまざまなワークロードにわたって機能する静的なしきい値を定義するのが難しい場合があります。 Google SRE は AI を使用して、静的に事前定義されたしきい値ではなく、通常の動作における異常の検出に基づいてアラートを発行する、異常検出による従来のアプローチを強化しています。このアプローチは、エージェントに依存してシグナルを収集し、それをモデル (例: TimesFM ) に供給して異常検出を実行します。以前の顧客事例からの履歴シグナルは、AI エージェントが顧客指向の SLO を予測するのに役立ちます。さらに、AI ベースの異常検出は、サービス自体によって生成される信号を超えたソース (顧客のフィードバックなど) を参照できます。
このモデルでは、SRE AI エージェントが異常を検出すると、アラートをトリガーします。次に、SRE AI アラート エージェントは、アラートをグループ化し、前処理し、必要なコンテキストと情報でアラートを強化します。これらのアラートは自律型 AI アラート ハンドラーを通じて実行され、多数の問題に対処または軽減できます。このシステムの結果、問題解決が迅速化され、SRE が確認する必要があるアラートの数が大幅に削減される可能性があります。
このエージェントのエコシステムで重要なのは、データ エージェントが何をどのように評価しているのかを一貫して透明性を保ち、運用状態の望ましくない突然変異を防ぐための一貫した制御を行うことです。
グー内

gle SRE、インシデント管理、または IMAG は、明確な役割と責任、およびツールを備えた確立されたプロセスです。 SRE AI には、現在の IMAG プロセスの上にエージェント オーケストレーション レイヤーが含まれており、次のエージェントで構成されます。
インシデント中に使用されたコミュニケーション サーフェス (インシデント対応ツール、チャット スペース、ビデオ、追跡文書) を監視し、データを統合/要約してインシデント中のコミュニケーションと情報共有を改善します。
必要なコンテキストを含む引き継ぎドキュメントを作成することで、インシデントに参加している SRE 間の引き継ぎをサポートします。
インシデント事後分析の草案を自動的に作成し、その品質を向上させ、SRE の労力を削減し、関連情報が確実に含まれるようにします。
社内外のインシデントコミュニケーションを管理する
Google SRE チームは、インシデントを調査し、場合によっては自律的に問題を軽減するためのエージェントも作成しました。
仮説の形成と緩和手順の提案に進む前に、これらのエージェントは可観測性データ (ロギング、モータリング、トレース) に加え、システム トポロジ、分類、および依存関係データを使用してドメインと意図を確立します。これらのエージェントが使用する他のいくつかの構成要素は、プレイブックの操作と実行、アラートへのアクセス、異常検出の実行、およびインシデントの洞察の導出のためにチームが作成した個別のエージェントです。
SRE には、エンドツーエンドのシステムと効果的な緩和ソリューションの理解、過去のインシデントから学んだ経験と教訓、リスク管理を実行する能力が必要です。自律型 AI エージェントが実稼働環境を管理できるようにするには、同様のスキルが必要です。
共通のトポロジまたは分類システムは、エージェントにエンドツーエンド システム、および十分に文書化され説明された運用モデル コンテキスト プロトコル (MCP) ツールとスキーについて教えることができます。

エージェントは利用可能なツールについて教えることができますが、歴史的な問題とそれに関連するリスクについてエージェントに継続的に教える方法が必要です。この問題を解決するために、Google SRE チームは AI Insights を作成しました。このシステムは、既知のインシデントを継続的にレビューし、そこから有意義な情報を抽出し、エージェントがそれを利用してより適切な調査と軽減手順を推進できるようにします。 Gemini 埋め込みモデルとベクトル対応データベースがこのシステムを強化します。
システムのもう 1 つの部分はリスクに関する洞察です。 AI システムは、各インシデントに適切なリスク カテゴリをマークします。このカテゴリは、エージェントが緩和策を適用する前に使用することも、SRE が対処すべき重要領域を決定するために使用することもできます。
これらのエージェントを構築する前に、Google SRE はエージェントの導入に関するいくつかの高レベルの原則を定義しました。
すでに自動化に成功しているプロセスや運用、または従来の非 AI ベースのシステムで簡単に自動化できるプロセスや運用は、(ビジネス ニーズを満たしている限り) 置き換える必要はありません。
新しい AI ベースのシステムは、顧客に対する強い約束を守るために、既存および今後のポリシーと手順に準拠する必要があります。
SRE AI エージェントは、現在のシステムや人間と同様に、セキュリティ、安全性、プライバシーの要件を満たす必要があります。
SRE AI エージェントには強力な ID が必要です (エージェントにはロールと権限が割り当てられています)。
SRE AI エージェントは、高レベルの信頼性 SLO を提供し、明確に定義されたバックアップ オプション (自動または手動) を備えている必要があります。
SRE AI エージェントは、アクションを実行した理由と方法、およびどのオプションが検討され拒否されたかについて説明し、推論できなければなりません。言い換えれば、私たちはブラックボックスの自動化よりも透明性を重視します。
事業継続計画には、潜在的な AI 障害に対する緊急事態を含める必要があります。
AI ベースのシステムは継続的にアクセスする必要があります。

正しい意思決定を行うための生産データ。
AI システムは、検出と対応などのセキュリティ ツールを可能にする監査とレポートをサポートするだけでなく、品質フレームワークに照らして継続的に評価する必要があります。
さらに、SRE AI システムは次の少なくとも 1 つを達成することで、ユーザーと顧客にとって Google サービスをさらに向上させる必要があると規定しました。
エンジニアを面倒な繰り返し作業から解放します。
エンジニアの意思決定と実行の質とスピードを向上させるのを支援する
SRE が以前よりも適切に問題を防止、検出、軽減できるようになります。
サービスの信頼性の向上を促進する自律的なエージェントのフィードバック ループを有効にする
全体的な運用コストを削減する
実証済みのインフラストラクチャ上に構築
Google SRE AI は、実証済みの Google インフラストラクチャ上に構築されています。
Gemini : Google SRE AI の背後にある基本基盤モデル。 SRE チームは、Google 内部のデータと知識に基づいてカスタムで微調整された Gemini モデルにも大きく依存しています。
Gemini Enterprise Agent Platform (旧 Vertex AI) : ソリューション開発用の完全な AI スタック。
エージェント開発キット (ADK): 開発プラットフォーム。
MCP サーバー: 標準の Google API インフラストラクチャ上で実行され、これは外部顧客に MCP サポートを提供するために使用されるインフラストラクチャと同じです。
標準の内部可観測性インフラストラクチャ (モニタリング、ロギング、トレース)。
Google BigQuery および Google ベクトル データベースに組み込まれた AI および ML 機能。
これらのインフラストラクチャ コンポーネントを自律システムにグループ化します。 Google では、生産を管理するための自律システムを開発し、使用してきました。ただし、今日の AI ベースの自律システムは非常に強力であり、必ずしも決定的であるとは限りません。システムが実際にどれほど自律的であるかを理解するのに役立つように、

自律レベルを追跡する方法を模索しました。
さらに詳しく: ホワイトペーパーを読む
これらのイノベーションの背後にある技術アーキテクチャと厳格なガバナンス モデルを探求したいエンジニアやリーダーは、包括的なホワイトペーパー「SRE 実践における AI: Google における自動化を超える」をぜひお読みください。このホワイトペーパーでは、Google SRE が決定論的自動化からエージェント的 AI への移行をどのように進めているかについて詳しく説明しています。ここからホワイトペーパーをダウンロードしてください。
Gemini Cloud Assist: 要求する前から、お客様に合わせて機能するプロアクティブなクラウド運用
Michael Bachman著 • 5分で読めます
統合メンテナンス: Google Cloud 全体のメンテナンスを管理する新しい統合方法
エロル・ヴァレリウ・キオアスカ著 • 2 分で読めます
プラットフォーム使用率の罠 パート 1: 高いアクティビティが必ずしも高い価値を意味するとは限らない理由
ダレン・エヴァンス著 • 9 分で読めます
プラットフォーム使用率の罠パート 2: 意味のある監視メトリクスの選択
ダレン・エヴァンス著 • 9 分で読めます
言語 英語 ドイツ語 フランス語 한국어 日本語

## Original Extract

With SRE AI, Google plans to fully adopt AI and agentic technologies, leveraging AI as a force multiplier while also maintaining control.

How Google SRE is using agentic AI to improve operations | Google Cloud Blog Jump to Content Cloud Blog Contact sales Get started for free Cloud Blog Solutions & technology AI & Machine Learning
DevOps & SRE AI in SRE: Where and how Google is deploying agentic AI to improve operations
Distinguished Software Engineer
Distinguished Site Reliability Engineer
The front door to AI in the workplace
Since its inception over 20 years ago, Google has used Site Reliability Engineering (SRE) to keep services like Search, Gmail, Maps, YouTube and Google Cloud reliable and highly available, adhering to the principles and practices of the reliability-first mindset.
Recently though, the emergence of AI has driven multiple step-changes in system complexity. Interactions between components are now more complicated due to a variety of factors:
With microservice architectures, systems are distributed across wider geographical locations and data centers that have greater hardware diversity.
Enterprise cloud products offer an extensive array of capabilities with an incredibly complex set of products.
Google services now cover more unique business and regulatory requirements, making the overall topology and taxonomy much more complex and difficult to understand, a challenge amplified by the constant stream of system changes resulting from continuous deployment pipelines.
AI code generation capabilities have enabled software developers to deliver orders of magnitude more code, resulting in more opportunities to introduce reliability issues.
While AI is in some ways making the SRE team’s work more challenging, it also provides new ways to understand and improve software development lifecycles, including production operations. Google SRE is on the path to fully adopt AI and agentic technologies, leveraging AI as a force multiplier while also maintaining control . We call this SRE AI.
Read on for a summary of considerations when thinking about this topic, or you can dive straight into our comprehensive whitepaper, AI in SRE Practice: Moving Beyond Automation at Google , for an in-depth look at how Google SRE is navigating the transition from deterministic automation to agentic AI .
The SRE AI opportunity landscape
To help define our SRE AI strategy, we considered the overall software development lifecycle (SDLC) for areas of opportunity.
The above diagram shows each of the phases where SRE is involved, and that could be improved with SRE AI.
Perhaps the most obvious SRE area that could benefit from agentic AI is investigation and mitigation , sometimes referred to as root cause analysis (RCA), a cornerstone of the traditional SRE discipline. But RCA is by no means the whole SRE AI. Our plans for SRE AI go far beyond RCA and troubleshooting, and address the entire SDLC. Here are a few areas we are working on:
SRE has been working on the policies, tooling and procedures you need to ensure reliability is an integral part of system design through the design, launch, and deployment phases. An agentic approach does not necessarily imply removing people from the process, specifically for higher-risk services and features, but it does significantly reduce the time people need to spend, as a number of issues can be detected and auto-addressed before they need to be reviewed by a person.
Runbooks (playbooks) and other documentation to be used during incidents are important production artifacts. Google SRE has developed AI agents to continuously monitor and improve playbooks and production documentation based on their usage during incidents. AI agents can also generate new playbooks from incidents.
Anomaly detection and alerting
A core SRE practice is to define service level indicators (SLIs) and service level objectives (SLOs) , and to configure alerts for them. This approach tends to be ok if service use cases are fairly uniform, and if it is possible to define objectives that align to customers' expectations.
However, for products that support a range of customer use cases and workloads, like many in Google Cloud, it can be difficult to define a static threshold that works across a variety of workloads. With AI, Google SRE is augmenting our more traditional approaches with anomaly detection , with alerts based on detecting anomalies in regular behavior rather than statically predefined thresholds. This approach relies on agents to collect signals and feed them to a model (e.g., TimesFM ) to perform anomaly detection. Historical signals from prior customer cases help the AI agent to predict customer-oriented SLOs. Further, AI-based anomaly detection can consult sources beyond signals produced by service itself — for instance, customer feedback.
In this model, when the SRE AI agent detects an anomaly, it triggers an alert. Then, the SRE AI alerting agent groups, pre-processes, and enriches the alerts with the necessary context and information. These alerts in turn are run through autonomous AI alert handlers, which can address or mitigate a multitude of issues. The outcome of this system is faster issue resolution and a likely significant reduction in the number of alerts that SREs need to review.
What's key in this ecosystem of agents is to be consistently transparent about what the data agents are evaluating — and how — and having consistent controls to prevent unwanted mutations of production state.
Within Google SRE, incident management, or IMAG , is a well-established process with clear roles and responsibilities, as well as tooling. SRE AI includes an agentic orchestration layer on top of the current IMAG process, which consists of agents that:
Monitor the communication surfaces used during the incident (incident response tools, chat spaces, videos, tracking documents), and consolidate/summarize data to improve communication and information sharing during the incident
Support handoff between SREs participating in the incident, by creating handoff documents with necessary context
Automatically create drafts of incident postmortems, improving their quality, reducing SRE effort, and ensuring that relevant information is included
Manage internal and external incident communications
The Google SRE team has also created agents to investigate incidents, and in some cases to autonomously mitigate issues.
Before they can proceed to form hypotheses and propose mitigation steps, these agents use observability data (logging, motoring, tracing), as well as system topology, taxonomy, and dependency data to establish domain and intent. A few other building blocks that these agents use are distinct agents the team has created for navigating and executing playbooks, accessing alerting, performing anomaly detection, and deriving incident insights.
SRE requires an understanding of the end-to-end system and effective mitigation solutions, experience and lessons learned from past incidents, and the ability to perform risk management. Autonomous AI agents need similar skills to be able to manage production environments.
While a common topology or taxonomy system can teach agents about the end-to-end system, and well-documented and described production Model Context Protocol (MCP) tools and skills can teach them about available tooling, there needs to be a way to continuously teach agents about historical issues and their associated risks. To solve that problem, the Google SRE team created AI Insights, a system that continuously reviews known incidents and extracts meaningful information from them, then makes it available to agents to drive better investigations and mitigation steps. Gemini embedding models and vector-enabled databases power this system.
The other part of the system is risk insights. The AI system marks each incident with appropriate risk categories that can be used both by agents before applying mitigations, and by SREs to determine critical areas to address.
Before building out these agents, Google SRE defined a few high level principles for their adoption:
Processes and operations that are already successfully automated, or that can be easily automated with classic non-AI based systems, do not need to be replaced (as long as they meet business needs).
Any new AI-based system must comply with existing and upcoming policies and procedures to keep the strong promises we have to our customers.
An SRE AI agent needs to meet security, safety, and privacy requirements the same way as current systems and humans.
SRE AI agents must have a strong identity (agents have roles and permissions assigned).
SRE AI agents need to provide a high level of reliability SLOs and have well-defined backup options (automated or manual).
SRE AI agents must be able to explain and reason about why and how they performed an action, as well as what options were considered and rejected. In other words, we favor transparency over black-box automation.
Business continuity plans must include contingencies for potential AI failures.
AI-based systems need continuous access to production data to make correct decisions.
AI systems need to be continuously evaluated against a quality framework, as well as to support auditing and reporting to enable security tooling like detection and response.
In addition, we stipulated that SRE AI systems should make Google services even better for users and customers by accomplishing at least one of the following:
Relieve engineers from laborious and repetitive operations
Help engineers improve the quality and speed of decision making and execution
Allow SREs to better prevent, detect, and/or mitigate problems than they could address before
Enable autonomous agentic feedback loops that drive toward service reliability improvements
Reduce overall operational costs
Built on proven infrastructure
Google SRE AI is built on proven Google infrastructure:
Gemini : The base foundational model behind Google SRE AI. The SRE team also depends heavily on custom fine-tuned Gemini models based on internal Google data and knowledge.
Gemini Enterprise Agent Platform (formerly Vertex AI) : A full AI stack for developing solutions.
Agent Development Kit ( ADK): The development platform.
MCP servers: Running on top of standard Google API infrastructure, this is the same infrastructure used to provide external customers with MCP support .
Standard internal observability infrastructure (monitoring, logging, tracing).
AI and ML capabilities built into Google BigQuery , and Google vector databases .
We group these infrastructure components together into autonomous systems. At Google, we’ve been developing and using autonomous systems to manage production for a long time. However, today’s AI-based autonomous systems are very powerful and not always deterministic. To help us understand how autonomous the systems truly are, we developed a way to track autonomous levels.
Dive deeper: Read the white paper
For engineers and leaders looking to explore the technical architecture and rigorous governance models behind these innovations, we invite you to read our comprehensive whitepaper, “AI in SRE Practice: Moving Beyond Automation at Google,” which provides an in-depth look at how Google SRE is navigating the transition from deterministic automation to agentic AI. Download the whitepaper here .
Gemini Cloud Assist: Proactive cloud operations that work for you, even before you ask
By Michael Bachman • 5-minute read
Unified Maintenance: A new, unified way to manage maintenance across Google Cloud
By Erol-Valeriu Chioasca • 2-minute read
The platform usage trap part 1: Why high activity doesn’t necessarily mean high value
By Darren Evans • 9-minute read
The platform usage trap part 2: Choosing meaningful monitoring metrics
By Darren Evans • 9-minute read
Language ‪English‬ ‪Deutsch‬ ‪Français‬ ‪한국어‬ ‪日本語‬
