---
source: "https://block.xyz/inside/block-rolls-out-builderbot-a-new-suite-of-ai-native-tools-that-changes-the-way-we-ship"
hn_url: "https://news.ycombinator.com/item?id=48618973"
title: "Block rolls out Builderbot, a new suite of AI-native tools"
article_title: "Block - Block rolls out Builderbot, a new suite of AI-native tools that changes the way we ship"
author: "msolujic"
captured_at: "2026-06-21T14:45:48Z"
capture_tool: "hn-digest"
hn_id: 48618973
score: 1
comments: 1
posted_at: "2026-06-21T13:48:34Z"
tags:
  - hacker-news
  - translated
---

# Block rolls out Builderbot, a new suite of AI-native tools

- HN: [48618973](https://news.ycombinator.com/item?id=48618973)
- Source: [block.xyz](https://block.xyz/inside/block-rolls-out-builderbot-a-new-suite-of-ai-native-tools-that-changes-the-way-we-ship)
- Score: 1
- Comments: 1
- Posted: 2026-06-21T13:48:34Z

## Translation

タイトル: Block が AI ネイティブ ツールの新しいスイートである Builderbot を公開
記事のタイトル: Block - Block は、出荷方法を変える AI ネイティブ ツールの新しいスイートである Builderbot を展開します
説明: AI は Block の構築と出荷方法にネイティブに組み込まれています。過去 2 年間、私たちは

記事本文:
Block - Block は、出荷方法を変える AI ネイティブ ツールの新しいスイートである Builderbot を展開します
ブロックロゴニュース
2026 年 6 月 17 日 Block は、出荷方法を変える新しい AI ネイティブ ツール スイートである Builderbot を公開します
AI は Block の構築と出荷方法にネイティブに組み込まれています。過去 2 年間、私たちは AI をエンジニアの働き方の基盤にすることに投資してきました。これには、AI エージェント フレームワークである Goose のオープンソース化、Anthropic との Model Context Protocol (MCP) の共同開発、すべてのエンジニアの日常ワークフローに AI を組み込む内部ツールの構築などが含まれます。現在、Block のエンジニアの 100% が業務で定期的に AI を使用しています。
それでも、私たちは同じ天井にぶつかり続けました。ほとんどのコーディング ツールは 1 つのリポジトリ内で適切に機能しますが、数億行のコード、数百のサービス、および Block のような企業が実際に構築する方法の完全な複雑さにわたって動作できるものはありませんでした。そこで、大規模な実装の解決に役立つ新しいツール Builderbot を構築しました。
Builderbot は、コードベース全体で複数の AI エージェントを調整するオーケストレーション レイヤーです。これは Slack 内で動作します。誰でも @builderbot に必要なものの短い説明をタグ付けすることができ、バグ修正、サービス間の移行、新機能など、スレッド内ですぐに動作するようになります。複数のチーム メンバーがリアルタイムで共同作業を行い、研究、計画、実装を観察しながら方向性を定めることができます。コンテキストの切り替えはなく、会話が開発環境となります。
これがコーディングアシスタントと異なるのはスコープです。 Builderbot は、Block のコードベース、すべてのサービス、すべての API、すべての規約の完全なコンテキストを理解し、会社全体のあらゆるコードに貢献できます。 Cash App に取り組んでいるエンジニアは、それを使用して、自分が利用している Square サービスに変更を加えることができます。

システムはそのサービスがどのように機能するかをすでに知っているため、決して触れられることはありません。 Linear と Jira から直接チケットを取得し、ブランチを作成し、コードを作成し、プル リクエストを開き、CI を監視し、フィードバックに基づいて反復します。人間が最も価値を付加するところに人間が介入します。
また、ソース コードとシステム構成のみで動作します。顧客データ、支払い情報、個人を特定できる情報へのアクセスや処理は行いません。
私たちの構築方法にとってそれが何を意味するか
Builderbot は 1 日あたり 200,000 件を超えるオペレーションを実行し、1 週間あたり約 1,500 件のプル リクエストをマージします。これは、Block 全体のすべての運用コード変更の約 15% に相当します。以前は数か月かかっていた作業が、今では数日で完了します。
「Builderbot について考える最良の方法は、AI コーディング ツールとエンジニアリングが実際に大規模に機能する方法の間に欠けている層として考えることです」と、Block の AI 機能責任者の Brad Axen 氏は述べています。 「オーケストレーション、コンテキスト、環境を処理するため、当社のエンジニアは解決する価値のある問題に集中できます。Square 側では、販売者が何か月も待ち続けていた機能のリストを取得し、当社のエンジニアが数日で出荷しました。Builderbot が足場と反復作業を処理し、当社のエンジニアが製品を形づくる決定を下しました。つまり、アイデアがバックログから数か月ではなく数日で数百万の顧客の前で実現できることを意味します。」
Builderbot は、Google が開発し、Linux Foundation 傘下の Agentic AI Foundation (AAIF) に貢献したオープン ソース エージェント フレームワークである Goose 上に構築されています。グースの構築で直面した統合の課題は、現在 AI エージェントをツールやデータ ソースに接続するための業界標準となっている MCP での Anthropic とのコラボレーションにインスピレーションを与えました。
Builderbot は必要だったので作成しました。私たちが解決している問題は Block に固有のものではありません。つまり、集団全体で AI エージェントを調整するということです。

コードベースを強化し、迅速に品質を維持し、人間が足場ではなく判断と好みに集中できるようにします。私たちがその構築方法を共有しているのは、AI 支援コーディングから AI ネイティブ エンジニアリングへの移行が、現在テクノロジー業界で起こっている最も重要な議論の 1 つであると考えており、オープンに貢献したいと考えているからです。
検索 AI チャット あなたは、生成 AI を利用した自動チャットボットに接続されています。続けることで、
あなたは、Block のプライバシーに関する通知がこのツールの使用に適用されることを認識し、こことこちらに記載されている Block の AI 利用規約に同意するものとします。

## Original Extract

AI is native to how Block builds and ships. Over the past two years, we

Block - Block rolls out Builderbot, a new suite of AI-native tools that changes the way we ship
Block logo News
June 17, 2026 Block rolls out Builderbot, a new suite of AI-native tools that changes the way we ship
AI is native to how Block builds and ships. Over the past two years, we've invested in making AI foundational to how our engineers work: open sourcing goose , our AI agent framework, co-developing the Model Context Protocol (MCP) with Anthropic, and building internal tools that bring AI into every engineer's daily workflow. Today, 100% of Block's engineers regularly use AI in their work.
Still, we kept hitting the same ceiling. Most coding tools work great in a single repo, but none of them could operate across hundreds of millions of lines of code, hundreds of services, and the full complexity of how a company like Block actually builds. So we built a new tool, Builderbot, to help us solve for implementation at scale.
Builderbot is an orchestration layer that coordinates multiple AI agents across our entire codebase. It works inside Slack: anyone can tag @builderbot with a short description of what they need, and it gets to work right there in the thread, whether it’s a bug fix, a migration across services, or a new feature. Multiple team members can collaborate with it in real time, watching it research, plan, and implement while they steer the direction. There's no context switching, the conversation is the development environment.
What makes this different from a coding assistant is scope. Builderbot understands the full context of Block's codebase, every service, every API, every convention, and can contribute to any piece of code at the entire company. An engineer working on Cash App can use it to make a change in a Square service they've never touched, because the system already knows how that service works. It picks up tickets directly from Linear and Jira, creates the branch, writes the code, opens the pull request, watches CI, and iterates based on feedback. Humans step in where humans add the most value.
It also operates on source code and system configuration only; it does not access or process customer data, payment information, or personally identifiable information.
What it means for how we build
Builderbot executes over 200,000 operations per day and merges approximately 1,500 pull requests per week, about 15% of all production code changes across Block. What used to take months now takes days.
"The best way to think about Builderbot is as the missing layer between AI coding tools and how engineering actually works at scale," said Brad Axen, Head of AI Capabilities at Block. "It handles the orchestration, the context, the environment, so our engineers can focus on the problems worth solving. On the Square side, we took a list of features sellers had been waiting on for months and our engineers shipped them in days. Builderbot handled the scaffolding and the repetitive work, and our engineers made the decisions that shaped the product. It means an idea can go from backlog to live in front of millions of customers in days instead of months."
Builderbot is built on goose, an open source agent framework we developed and contributed to the Agentic AI Foundation (AAIF) under the Linux Foundation. The integration challenges we hit building goose inspired our collaboration with Anthropic on MCP, now an industry standard for connecting AI agents to tools and data sources.
We created Builderbot because we needed it. The problems we're solving aren't unique to Block: orchestrating AI agents across a massive codebase, maintaining quality at speed, keeping humans focused on judgment and taste rather than scaffolding. We're sharing how we built it because we think the shift from AI-assisted coding to AI-native engineering is one of the most important conversations happening in technology right now, and we want to contribute openly.
Search AI Chat You are being connected to our automated chatbot which utilizes generative AI. By continuing,
you recognize that Block’s Privacy Notice applies to your use of this tool and you agree to Block’s AI terms of use located here and here .
