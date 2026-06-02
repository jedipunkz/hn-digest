---
source: "https://blog.jetbrains.com/ai/2026/06/mellum2-goes-open-source-a-fast-model-for-ai-workflows/"
hn_url: "https://news.ycombinator.com/item?id=48356880"
title: "Mellum2 Goes Open Source: A Fast Model for AI Workflows"
article_title: "Mellum2 Goes Open Source: A Fast Model for AI Workflows | The JetBrains AI Blog"
author: "chillax"
captured_at: "2026-06-02T00:40:00Z"
capture_tool: "hn-digest"
hn_id: 48356880
score: 7
comments: 0
posted_at: "2026-06-01T13:52:30Z"
tags:
  - hacker-news
  - translated
---

# Mellum2 Goes Open Source: A Fast Model for AI Workflows

- HN: [48356880](https://news.ycombinator.com/item?id=48356880)
- Source: [blog.jetbrains.com](https://blog.jetbrains.com/ai/2026/06/mellum2-goes-open-source-a-fast-model-for-ai-workflows/)
- Score: 7
- Comments: 0
- Posted: 2026-06-01T13:52:30Z

## Translation

タイトル: Mellum2 のオープンソース化: AI ワークフローの高速モデル
記事のタイトル: Mellum2 のオープンソース化: AI ワークフローの高速モデル | JetBrains AI ブログ
説明: Mellum2 は、ゼロからトレーニングされ、実際の展開向けに設計されており、ソフトウェア エンジニアリング システムでのルーティング、Q&A、サブエージェント、およびプライベート AI の使用のために構築されています。本日、Mellum2 をオープンソース化します

記事本文:
プラグインとサービス
ビッグデータツール
.NET と Visual Studio
.NETツール
教育と研究
ジェットブレインズアカデミー
多くの JetBrains 製品内の AI を活用した機能でツールを強化します
Mellum2 のオープンソース化: AI ワークフローの高速モデル
この投稿を他の言語で読んでください:
Mellum2 は、ゼロからトレーニングされ、実際の展開向けに設計されており、ソフトウェア エンジニアリング システムでルーティング、Q&A、サブエージェント、およびプライベート AI を使用できるように構築されています。
現在、私たちは、実稼働 AI の最も困難な部分であるレイテンシ、スループット、コストを解決するために設計された 12B モデルである Mellum2 をオープンソース化しています。スクラッチから構築され、Apache 2.0 ライセンスに基づいてリリースされた Mellum2 は、インフラストラクチャに高性能でコスト効率の高い代替手段を提供します。
Mellum はコード補完から始まりました。今では、自然言語とコードの両方を処理できるように進化させました。これは、最新の AI ワークフロー全体でルーティング、要約、中間推論のステップを強化できる多用途ツールになりました。
実験、微調整、または大規模な展開を希望する場合でも、Mellum2 は独自のシステムで実行する準備ができています。
Mellum2 は、そのアーキテクチャと集中的で効率重視の設計を通じて、生産規模のシステムのボトルネックを解決するように設計されています。
Mixture-of-Experts (MoE) 設計: このモデルは合計 12B のパラメーターを備えていますが、MoE 設計を使用しているため、トークンごとにアクティブになるパラメーターは 2.5B のみです。これにより、リアルタイム ワークロードに対する高スループット、低遅延の推論が可能になりながら、コンピューティング コストが削減されます。
特化した焦点: 多くの最新モデルとは異なり、Mellum2 はマルチモーダルではありません。これは、特に自然言語とコード データに基づいてトレーニングされています。この特殊化により、モデルは無駄のない高速性を維持しながら、ソフトウェア エンジニアリング環境で優れた性能を発揮します。
技術レポートでは、コード生成全体にわたるモデルのパフォーマンスについて詳しく説明しています。

イオン、科学、数学、推論のベンチマーク。 Mellum2 は、他の同様のサイズのモデルと比較して競争力があり、推論時間を半分以下に短縮します。これは、実稼働グレードの導入にとって決定的な利点です。
AI ワークロードのルーティングとオーケストレーション: Mellum2 を使用して受信プロンプトを分析し、各タスクに適切なモデルまたはツールの選択を支援します。
低レイテンシの RAG パイプラインを構築します。関連するコンテキストを取得し、Mellum2 を使用してそれを要約し、即座に応答を生成します。
複雑なワークフローで高速サブエージェントを強化: エージェント パイプラインをコンテキストの収集、計画、検証などのステップに分割します。単一の大規模モデルに依存するのではなく、高速で特殊なタスクには Mellum2 を使用してください。
プライベートなローカル AI デプロイメントを有効にする: Mellum2 をローカルで実行するか自己ホストして、コードとデータを完全に制御下に置きます。
「焦点モデル」の哲学: 焦点を絞ったモデルの拡張性が向上する理由
AI システムがより複雑になるにつれて、パフォーマンスのボトルネックは、生の機能から、大規模なレイテンシ、スループット、コストへと移行します。すべてのタスクに最大のモデルが必要なわけではありません。最新の AI システムの多くのステップは反復的で、遅延の影響を受けやすく、頻度が高くなります。これらのステップは、効率的にルーティング、ホスト、制御できる高速で信頼性の高いモデルの恩恵を受けます。
JetBrains では、未来は単一のモデルではなく、調整されたシステムに属すると信じています。フロンティア モデルは引き続き限界を押し広げていくでしょうが、実用的な AI 製品にはフォーカル モデル、つまり高頻度のタスクを効率的に処理する高速で特殊なコンポーネントも必要です。
それが、次世代の AI ソフトウェア ツールにおける Mellum2 の役割です。
ソフトウェア エンジニアリング用の AI システムを構築している場合 (IDE 内、RAG パイプライン、エージェント ワークフローの一部として、または完全に独自のインフラストラクチャ上で) をぜひお試しください。
オープンソース

そうすることでより良いツールが作られるのです。
このフォームを送信することにより、JetBrains s.r.o. が次のことに同意します。 (「JetBrains」) は、商業通信を含むニュースレターを私に送信し、この目的で私の個人データを処理するために、私の名前、電子メール アドレス、および位置データを使用することがあります。 JetBrains が、JetBrains プライバシー ポリシーに従って、この目的のためにサードパーティのサービスを使用して当該データを処理する場合があることに同意します。私は、自分のプロフィールでいつでもこの同意を取り消すことができることを理解しています。さらに、各メールには購読解除リンクが含まれています。
プロジェクト内の最近のアクティビティや非自明なコードを理解するのに積極的に役立つ 2 つの実験的な AI 機能、要約とインサイトを紹介します。
Cursor は、エージェント クライアント プロトコルを通じて JetBrains IDE 内の AI エージェントとして利用できるようになりました。
JetBrains コンソールを導入します。これは、チーム全体で AI の使用状況とコストを管理、観察、制御するための新機能を含む、組織向けに強化された AI 管理と分析を提供します。
ネイティブ ACP × Koog 統合を使用してカスタム Koog エージェントを IDE に取り込む方法に関するステップバイステンガイド

## Original Extract

Trained from scratch and designed for practical deployment, Mellum2 is built for routing, Q&A, sub-agents, and private AI use in software engineering systems. Today, we’re open-sourcing Mellum2

Plugins & Services
Big Data Tools
.NET & Visual Studio
.NET Tools
Education & Research
JetBrains Academy
Supercharge your tools with AI-powered features inside many JetBrains products
Mellum2 Goes Open Source: A Fast Model for AI Workflows
Read this post in other languages:
Trained from scratch and designed for practical deployment, Mellum2 is built for routing, Q&A, sub-agents, and private AI use in software engineering systems.
Today, we’re open-sourcing Mellum2, a 12B model engineered to solve the hardest parts of production AI: latency, throughput, and cost. Built from scratch and released under the Apache 2.0 license, Mellum2 offers a high-performance, cost-efficient alternative for your infrastructure.
Mellum began with code completion ; now we’ve evolved it to handle both natural language and code. It is now a versatile tool ready to power routing, summarization, and intermediate reasoning steps across your modern AI workflows.
Whether you want to experiment, fine-tune, or deploy at scale, Mellum2 is ready to run in your own systems.
Mellum2 is engineered to solve the bottlenecks of production-scale systems through its architecture and focused, efficiency-driven design.
Mixture-of-Experts (MoE) design: The model features 12B total parameters, but because it uses a MoE design, only 2.5B parameters are active per token. This reduces compute costs while enabling high-throughput, low-latency inference for real-time workloads.
Specialized focus: Unlike many modern models, Mellum2 is not multimodal. It is trained specifically on natural language and code data. This specialization ensures the model excels in software engineering environments while remaining lean and fast.
In our technical report , we detail our model’s performance across code generation, science, math, and reasoning benchmarks. Mellum2 is competitive with other similar-sized models while cutting inference time to less than half – a definitive advantage for production-grade deployments.
Route and orchestrate AI workloads: Use Mellum2 to analyze incoming prompts and help select the right model or tool for each task.
Build low-latency RAG pipelines: Retrieve relevant context, use Mellum2 to summarize it, and generate responses instantly.
Power fast sub-agents in complex workflows: Break down agent pipelines into steps like context gathering, planning, and validation. Use Mellum2 for fast, specialized tasks instead of relying on a single large model.
Enable private, local AI deployment: Run Mellum2 locally or self-host it to keep code and data fully under your control.
The “focal model” philosophy: Why focused models scale better
As AI systems become more complex, performance bottlenecks shift from raw capability to latency, throughput, and cost at scale. Not every task requires the largest model. Many steps in modern AI systems are repetitive, latency-sensitive, and high-frequency. These steps benefit from a fast and reliable model that can be efficiently routed, hosted, and controlled.
At JetBrains, we believe the future belongs to coordinated systems, not single models. Frontier models will continue to push the limits, but practical AI products also require focal models: fast, specialized components that handle high-frequency tasks efficiently.
That’s the role we see for Mellum2 in the next generation of AI software tooling.
If you’re building AI systems for software engineering – whether inside an IDE, in a RAG pipeline, as part of an agent workflow, or fully on your own infrastructure – we’d love for you to try Mellum2.
Open source is how better tools get made.
By submitting this form, I agree that JetBrains s.r.o. ("JetBrains") may use my name, email address, and location data to send me newsletters, including commercial communications, and to process my personal data for this purpose. I agree that JetBrains may process said data using third-party services for this purpose in accordance with the JetBrains Privacy Policy . I understand that I can revoke this consent at any time in my profile . In addition, an unsubscribe link is included in each email.
Introducing recap and insights, two experimental AI features that proactively help you understand recent activity and non-obvious code in your project.
Cursor is now available as an AI agent inside JetBrains IDEs through the Agent Client Protocol.
We’re introducing the JetBrains Console, which provides enhanced AI management and analytics for organizations, including new capabilities to manage, observe, and control AI usage and costs across teams.
Step by sten guide on how to bring your custom Koog agent into IDE with native ACP × Koog integration
