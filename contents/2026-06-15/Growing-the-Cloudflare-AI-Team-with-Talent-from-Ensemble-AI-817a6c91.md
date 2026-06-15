---
source: "https://blog.cloudflare.com/ensemble-ai-talent-joins-cloudflare/"
hn_url: "https://news.ycombinator.com/item?id=48543012"
title: "Growing the Cloudflare AI Team with Talent from Ensemble AI"
article_title: "Growing the Cloudflare AI team with talent from Ensemble AI"
author: "jgrahamc"
captured_at: "2026-06-15T16:42:38Z"
capture_tool: "hn-digest"
hn_id: 48543012
score: 2
comments: 0
posted_at: "2026-06-15T15:45:34Z"
tags:
  - hacker-news
  - translated
---

# Growing the Cloudflare AI Team with Talent from Ensemble AI

- HN: [48543012](https://news.ycombinator.com/item?id=48543012)
- Source: [blog.cloudflare.com](https://blog.cloudflare.com/ensemble-ai-talent-joins-cloudflare/)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T15:45:34Z

## Translation

タイトル: Ensemble AI の人材による Cloudflare AI チームの成長
記事のタイトル: Ensemble AI の人材を活用して Cloudflare AI チームを成長させる
説明: Cloudflareは、Ensemble AIのチームメンバーを加えて、機械学習インフラストラクチャと効率に焦点を当ててAIへの投資を強化しています。

記事本文:
無料で始めましょう |営業担当者へのお問い合わせ Cloudflare ブログ
新しい投稿の通知を受け取るには購読してください:
Ensemble AI の人材を活用して Cloudflare AI チームを成長させる
本日、Ensemble AI のチームの主要メンバーが Cloudflare に参加し、AI インフラストラクチャでの作業を加速し、開発者が強力な AI モデルを効率的に大規模に実行できるようになることをお知らせできることを嬉しく思います。
2023 年にサンフランシスコで設立された Ensemble AI は、ここ数年、AI における最も重要な課題の 1 つである、品質を犠牲にすることなく、大規模なモデルをより高速に、より小型に、よりコスト効率よく提供できるようにすることに注力してきました。チームは、大規模な言語モデルとマルチモーダル アーキテクチャのメモリ、計算、展開のオーバーヘッドを削減するように設計された、モデル圧縮と効率的な推論に対する新しいアプローチを開発しました。
AI が開発者によるアプリケーション構築の中核部分になるにつれ、推論の経済性がこれまで以上に重要になっています。モデルは大型化しています。ワークロードはより動的になっています。そして顧客は、AI が世界中に分散され、高速で、信頼性が高く、手頃な価格でどこでも利用できることをますます期待しています。 Ensemble AI チームを Cloudflare に導入することで、それを可能にする能力が強化されます。
アンサンブルの専門知識を取り入れる
Ensemble AI のチームは、最新の AI モデル内の構造を維持しながら、実行コストを削減することに重点を置いています。 Ensemble は、モデルの効率を単なる量子化やハードウェアの問題として扱うのではなく、ニューラル ネットワークをアーキテクチャ レベルでよりコンパクトかつ効率的にできる新しいモデルのビルディング ブロックを探索しました。
この作業の中核となる部分は NdLinear です。これは、フラットではなく多次元アクティベーションで直接動作する、トランスフォーマー モデルの標準線形レイヤーのドロップイン置換です。

構造を遠ざけます。これにより、モデルはパラメーター数と計算を削減しながら、ヘッド、チャネル、空間次元、その他の構造化された表現などの意味のある軸を保持できるようになります。 Ensemble は、大規模モデルの微調整に必要なトレーニング可能なパラメーターを削減するように設計された効率的な適応方法である NdLinear-LoRA も開発しました。
これらのアプローチは、量子化やベクトル量子化などの他の効率化手法を補完します。これらは共に、開発者が大幅に低いメモリ、コンピューティング、およびコストの要件で有能な AI モデルを実行できる未来を示しています。
AI 推論の効率化
Cloudflare Workers AI により、開発者は Cloudflare のグローバル ネットワーク上でサーバーレス GPU を利用した推論にアクセスできるようになります。開発者がより多くの AI ネイティブ アプリケーションを構築するにつれて、モデルを効率的に提供する機能がプラットフォームの重要な部分になります。
推論コストは、AI アプリケーションの拡張における最大の障壁の 1 つです。モデル サイズ、メモリ フットプリント、スループット、GPU 使用率が向上するたびに、開発者にとって AI がより利用しやすくなり、顧客にとってはより経済的になります。 AI ワークロードが単純なテキスト生成を超えて、エージェント、マルチモーダル モデル、パーソナライゼーション、微調整、検索、強化学習にまで拡張されるため、これは特に重要です。
私たちは、Workers AI の高速化、柔軟性、コスト効率の向上に必要なコア機械学習機能への投資を強化しています。これは、推論エンジン Infire 、 Unweight などのテンソル圧縮技術、および超大規模言語モデルを実行するためのプラットフォームを含む、モデルの効率を向上させるための既存の作業の上に構築されています。チームは、大規模な言語モデルやその他の高度な AI アーキテクチャを提供する経済性の向上に重点を置きます。

モデルの効率、GPU 使用率、スケーラブルな展開について。
次世代の AI ワークロードの構築
AIインフラは新たなフェーズに入りつつある。開発者はモデルにアクセスするだけで済みます。ユーザーの近くでモデルを確実に、手頃な価格で実行できるインフラストラクチャが必要です。コストや運用の複雑さに妨げられることなく、さまざまなモデル サイズ、微調整アプローチ、展開パターンを実験できる必要があります。
Cloudflare は、この問題の解決を支援する独自の立場にあります。当社のグローバル ネットワーク、開発者プラットフォーム、サーバーレス アーキテクチャは、AI をアプリケーションがすでに実行されている場所に近づけるための基盤を提供します。 Workers AI Machine Learning Engineering チームは、そのエクスペリエンスの下にある効率層の改善を支援します。
Cloudflare のグローバル インフラストラクチャと Ensemble のモデル圧縮および効率的なアーキテクチャの取り組みを組み合わせることで、開発者が低コスト、より優れたパフォーマンス、より少ない運用オーバーヘッドで AI アプリケーションをデプロイできるプラットフォームの構築を継続できます。
私たちは協力して、AI をより効率的に、アクセスしやすく、どこの開発者にとっても役立つものにするために必要なインフラストラクチャの構築を続けていきます。私たちの目標はシンプルです。Cloudflare プラットフォーム全体で推論の経済性を向上させながら、開発者が世界規模で強力な AI ワークロードを実行できるように支援することです。私たちの使命に参加したい場合は、採用ページをチェックしてください。
最先端のサイバーモデルからの防御: 顧客ゼロとしての Cloudflare のアーキテクチャ
Project Glasswing に関する投稿の中で、パッチの速度よりも脆弱性に関するアーキテクチャの方が重要であると主張しました。ここでは、そのアーキテクチャがどのようなものであるか、それが防御する脅威、そしてCloudflareのカスタマーゼロとして私たち自身がそれをどのように実行するかについて説明します。 ...
AI の請求書 i

制御不能です。 Cloudflare なら今すぐ修正できます。
AI ゲートウェイには、複数の AI プロバイダーにわたるトークン請求の暴走を防ぐためのリアルタイムの支出制限が追加されました。 Cloudflare Accessと統合することで、企業はアイデンティティ主導の予算とポリシーを使用できるようになります。 ...
VoidZeroがCloudflareに参加します
Vite、Vitest、Rolldown、Oxc、Vite+ を支えるチームである VoidZero が Cloudflare に加わります。 Vite はオープンソースであり、ベンダーに依存せず、すべての人のために構築されています。
...
Cloudflareのデータプラットフォームとその上にAIエージェントを構築した方法
ここでは、Cloudflare の統合分析プラットフォームである Town Lake と、その上で実行される内部 AI エージェントである Skipper をどのように構築したかを説明します。 ...

## Original Extract

Cloudflare is deepening our investment in AI with the addition of team members from Ensemble AI, focusing on machine learning infrastructure and efficiency.

Get Started Free | Contact Sales The Cloudflare Blog
Subscribe to receive notifications of new posts:
Growing the Cloudflare AI team with talent from Ensemble AI
Today, we’re excited to share that key members of the team at Ensemble AI are joining Cloudflare to help accelerate our work in AI infrastructure and make it easier for developers to run powerful AI models efficiently at scale.
Ensemble AI, founded in 2023 in San Francisco, has spent the last few years focused on one of the most important challenges in AI: making large models faster, smaller, and more cost-effective to serve, without sacrificing quality. The team has developed new approaches to model compression and efficient inference that are designed to reduce the memory, compute, and deployment overhead of large language models and multimodal architectures.
As AI becomes a core part of how developers build applications, the economics of inference matter more than ever. Models are getting larger; workloads are becoming more dynamic. And customers increasingly expect AI to be available everywhere: globally distributed, fast, reliable, and affordable. Bringing the Ensemble AI team into Cloudflare strengthens our ability to make that possible.
Incorporating Ensemble’s expertise
The team at Ensemble AI has focused on preserving the structure inside modern AI models while reducing the cost of running them. Instead of treating model efficiency as only a quantization or hardware problem, Ensemble has explored new model building blocks that can make neural networks more compact and efficient at the architectural level.
A core part of this work is NdLinear , a drop-in replacement for standard linear layers in transformer models that operates directly on multidimensional activations rather than flattening structure away. This enables models to preserve meaningful axes, such as heads, channels, spatial dimensions, or other structured representations, while reducing parameter count and compute. Ensemble has also developed NdLinear-LoRA, an efficient adaptation method designed to reduce the trainable parameters required for fine-tuning large models.
These approaches complement other efficiency techniques, including quantization and vector quantization. Together, they point toward a future where developers can run capable AI models with substantially lower memory, compute, and cost requirements.
Making AI inference more efficient
Cloudflare Workers AI gives developers access to serverless GPU-powered inference on Cloudflare’s global network. As developers build more AI-native applications, the ability to serve models efficiently becomes a critical part of the platform.
Inference cost is one of the biggest barriers to scaling AI applications. Every improvement in model size, memory footprint, throughput, and GPU utilization can make AI more accessible to developers and more economical for customers. This is especially important as AI workloads expand beyond simple text generation into agents, multimodal models, personalization, fine-tuning, retrieval, and reinforcement learning.
We are deepening our investment in the core machine learning capabilities needed to make Workers AI faster, more flexible, and more cost-efficient. This builds on top of our existing work on improving model efficiency, including our inference engine Infire , tensor compression techniques like Unweight , and our platform for running extra large language models . The team will focus on improving the economics of serving large language models and other advanced AI architectures, with an emphasis on model efficiency, GPU utilization, and scalable deployment.
Building for the next generation of AI workloads
AI infrastructure is entering a new phase. Developers no longer need only access to models; they need infrastructure that can run models reliably, affordably, and close to users. They need the ability to experiment with different model sizes, fine-tuning approaches, and deployment patterns without being blocked by cost or operational complexity.
Cloudflare is uniquely positioned to help solve this. Our global network, developer platform, and serverless architecture give us the foundation to bring AI closer to where applications already run. The Workers AI Machine Learning Engineering team will help us improve the efficiency layer underneath that experience.
By combining Cloudflare’s global infrastructure with Ensemble’s work in model compression and efficient architectures, we can continue building a platform where developers can deploy AI applications with lower cost, better performance, and less operational overhead.
Together, we will continue building the infrastructure needed to make AI more efficient, accessible, and useful for developers everywhere. Our goal is simple: help developers run powerful AI workloads at global scale while improving the economics of inference across the Cloudflare platform. If you want to join us in our mission, check out our careers page .
Defend against frontier cyber models: Cloudflare's architecture as customer zero
In our post about Project Glasswing, we made the argument that the architecture around a vulnerability matters more than the speed of the patch. Here we walk through what that architecture looks like, the threats it defends against, and how we run it ourselves as Cloudflare's customer zero. ...
Your AI bill is out of control. Cloudflare can fix it now.
AI Gateway now features real-time spend limits to prevent runaway token bills across multiple AI providers. By integrating with Cloudflare Access, companies can use identity-driven budgets and policies. ...
VoidZero is joining Cloudflare
VoidZero, the team behind Vite, Vitest, Rolldown, Oxc, and Vite+, is joining Cloudflare. Vite stays open source, vendor-agnostic, and built for everyone.
...
How we built Cloudflare's data platform and an AI agent on top of it
Here’s how we built Town Lake, Cloudflare's unified analytics platform, alongside Skipper, an internal AI agent running on top of it. ...
