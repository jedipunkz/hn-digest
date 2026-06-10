---
source: "https://www.vanta.com/resources/vanta-ai-quality-evaluation-maturity-model"
hn_url: "https://news.ycombinator.com/item?id=48478761"
title: "The Vanta AI Quality Eval Maturity Model"
article_title: "The Vanta AI Quality Eval Maturity Model | Vanta"
author: "hamelj"
captured_at: "2026-06-10T19:03:05Z"
capture_tool: "hn-digest"
hn_id: 48478761
score: 3
comments: 0
posted_at: "2026-06-10T16:28:14Z"
tags:
  - hacker-news
  - translated
---

# The Vanta AI Quality Eval Maturity Model

- HN: [48478761](https://news.ycombinator.com/item?id=48478761)
- Source: [www.vanta.com](https://www.vanta.com/resources/vanta-ai-quality-evaluation-maturity-model)
- Score: 3
- Comments: 0
- Posted: 2026-06-10T16:28:14Z

## Translation

タイトル: Vanta AI 品質評価成熟度モデル
記事のタイトル: Vanta AI 品質評価成熟度モデル |ヴァンタ
説明: 未加工の LLM と本番品質の AI の間には、大きな隔たりがあります。これを埋めるために Vanta が構築した 5 次元のフレームワークは次のとおりです。

記事本文:
Vanta AI 品質評価成熟度モデル |ヴァンタ
プラットフォーム
製品
プラットフォーム
コンプライアンス 自動化により、迅速かつ簡単にコンプライアンスを実現します。継続的な GRC GRC への最新の方法に参加してください。担当者とアクセス ユーザーのアクセスと権限を簡単に制御します。リスク管理 リスクを積極的に管理して、より賢明な意思決定を推進します。サードパーティのリスク管理 ベンダーのオンボーディングとセキュリティ レビューを 1 か所で管理します。アンケートの自動化 セキュリティアンケートの回答を自動化します。トラスト センター コンプライアンス ステータスとドキュメントを紹介します。監査の合理化 監査の準備と証拠の収集を自動化します。顧客のコミットメント あらゆる顧客のコミットメントを一元管理し、追跡し、行動します。
Vanta AI AI を使用してコンプライアンスを自動化し、洞察を明らかにします。 Agentic Trust Platform 単一の統合プラットフォームから信頼を構築し、証明します。統合 400 以上のツールからデータを自動的に取得します。 Vanta API カスタム統合とワークフローを構築します。 NEW RELEASE 最新情報を見る
ヴァンタが届ける 詳細はこちら
製品のコンプライアンス 自動化により、迅速かつ簡単にコンプライアンスを取得します。担当者とアクセス ユーザーのアクセスと権限を簡単に制御します。リスク管理 リスクを積極的に管理して、より賢明な意思決定を推進します。サードパーティのリスク管理 ベンダーのオンボーディングとセキュリティ レビューを 1 か所で管理します。アンケートの自動化 セキュリティアンケートの回答を自動化します。トラスト センター コンプライアンス ステータスとドキュメントを紹介します。監査の合理化 監査の準備と証拠の収集を自動化します。顧客のコミットメント あらゆる顧客のコミットメントを一元管理し、追跡し、行動します。
Vanta AI AI を使用してコンプライアンスを自動化し、洞察を明らかにします。プラットフォーム インタラクティブなデモを見る
Agentic Trust Platform 単一の統合プラットフォームから信頼を構築し、証明します。統合 [integrations_count] ツールからデータを自動的に取得します。

Vanta API カスタム統合とワークフローを構築します。ソリューション
サイズ
産業
フレームワーク
パートナーを探す
スタートアップ企業はコンプライアンスを自動化して構築を継続できます。中規模市場 セキュリティおよびコンプライアンス プログラムを次のように拡張します。
[切り捨てられた]
中小企業向けのセキュリティ ソリューションを加速
‍
タゴールは中小企業に戦略的なサービスを提供しています。
拡張可能なパートナーシップ
‍
タゴール氏は、確立された製品、専任のサポート チーム、迅速なリリース速度を備えたマネージド コンプライアンス パートナーを見つけることを優先しました。
競合他社よりも目立つ
‍
タゴールとヴァンタのパートナーシップは、戦略的焦点を強化し、顧客価値を深め、競争市場での差別化を生み出します。
このブログは Trustcraft シリーズの一部であり、AI を使用した構築に対する Vanta のアプローチを詳しく掘り下げています。 Trustcraft の定義方法について詳しくは、このシリーズの最初のブログをお読みください。
ChatGPT と Claude で何ができるかを見てきました。基盤モデルをほぼあらゆるものに接続できる MCP、CLI、および API について聞いたことがあるでしょう。 Vanta の AI 機能を見ると、当然の疑問が生じるかもしれません。汎用 LLM を自社のデータ ソースに接続して、それで終わりにしてはどうでしょうか?
それは合理的な本能です。しかし、コンプライアンス、信頼、セキュリティのワークフローにわたって何千もの顧客にサービスを提供する AI システムを構築した後、私たちは外からは明らかではないことを学びました。それは、未加工の LLM 統合と本番品質の AI 製品との間には非常に大きなギャップがあり、賭け金が上がるにつれてそのギャップは拡大するということです。
この投稿は、そのギャップに何が住んでいるのかについてです。
「LLM を接続するだけ」という幻想
ファウンデーションモデルは非常に優れた機能を備えています。しかし、能力と信頼性は同じものではありません。コンプライアンスとセキュリティにおいて、「良好」とは、管理要件を正しく解釈し、証拠に基づいて評価することを意味します。

常に規制上のエッジケースに対処し、監査を危険にさらす可能性のある詳細を決して幻覚させません。生の LLM 統合により、これらのそれぞれについてコイン投げが可能になります。 Vanta の AI は、大規模かつ繰り返し適切に実行できるように設計されています。
違いはモデルではなく、プロンプト、コンテキストの取得、メモリ、ドメイン固有のスキャフォールディング、および汎用モデルをコンプライアンス対応ツールに変えるシステム設計など、モデルの周囲のすべてです。 Vanta の AI を使用すると、これらの各レイヤーに投入された深く集中した作業から恩恵を受けることができます。
Vanta を使用する場合、敏速なエンジニアである必要はなく、独自の検索システムを設計する必要も、コントロール、証拠、フレームワークに関するコンテキストをつなぎ合わせる必要もありません。それはちょうどうまくいきます。 ChatGPT または Claude を自分で API に接続する場合、そのすべてを構築するか、より悪い回答を受け入れるかは自分の責任です。
すべての Vanta AI 機能を同じ厳格な品質基準に保つフレームワークを満たします
本番品質の AI 製品を確実に生産するために、私たちはただ何かを出荷するだけではありません。しかし、AI ポートフォリオが拡大するにつれて、私たちは問題に気づきました。AI チームは、品質と評価における「十分な品質」がどのようなものかについて共通の理解を持っていませんでした。チームは優れた機能を構築していましたが、それらの評価には一貫性がなく、異なる標準、異なるツール、異なる厳密さのレベルがありました。
そこで私たちは、AI 品質評価成熟度モデルと呼ぶ、厳格で多次元の品質システムを構築しました。このシステムは、Vanta のすべての AI を活用した機能が、時間の経過とともにどのように開発、測定、改善されるかを管理しています。
このモデルは、5 つの重要な側面にわたって AI システムを評価し、それぞれがチームに求める理想的な状態を表しています。
可観測性 : すべての AI インタラクション - 入力にわたる完全なトレース カバレッジ

、出力、推論ステップ、およびメタデータ - リアルタイムの監視と、動作が変動した場合の自動アラートを備えています。
厳選された評価データセット : GRC の対象分野の専門家によって厳選され、バージョン管理され、積極的に保守されているデータセット。アドホックなテスト セットではなく、現実世界の複雑さを反映して進化するコレクションです。
調整された評価者 : 対象分野の専門家の専門知識と Vanta の深い信頼とコンプライアンス領域の知識に基づいて調整された、検証済みの LLM-as-a-judge システムを含む正式な評価パイプライン。これにより、品質評価が一貫しており、単にモデルが正しいと考えているものではなく、本番環境で実際に重要なことと一致していることが保証されます。
体系的な実験 : 影響を測定し、次のステップを決定するための明確な基準を備えた、あらゆる AI 変更に対する構造化された反復可能な実験と分析のサイクル。
統合されたフィードバック ループ : 明示的および暗黙的なユーザー フィードバックが自動的に取得され、タグ付けされ、評価データセットにリンクされて、実際の顧客エクスペリエンスが改善を促進する継続的なサイクルが作成されます。
評価成熟度モデルによって私たちの仕事はどう変わったか
シンプルな評価システムを使用して、赤 (基礎/対応型)、黄 (体系的/発展途上)、緑 (高度/積極的) の 5 つの各側面ですべての AI チームをスコアリングします。
私たちが最初にこの評価を実行したとき、状況は謙虚なものでした。ほとんどのチームは、アドホックなデータセット、一貫性のない評価プロセス、手動のフィードバック ループに依存しており、赤と黄色のレベルにありました。 AI を使用して構築しているほとんどの組織が今日抱えているのと同じギャップが私たちにもありました。
しかし、まさにそれこそがこのモデルが存在する理由です。すべての AI チームが同じインフラストラクチャを活用できるように評価ツールを標準化し、評価のベスト プラクティスについてチームを教育し、AI プラットフォーム チームを ea と緊密に提携させました。

ch 製品チームがレベルアップを支援します。
その結果、かつては赤と黄色に染まっていたチームが、さまざまな面で緑に到達しました。
成熟度モデル全体で「オールグリーン」に到達することが最終目標ではないことに注意してください。本当の最終目標は、顧客が信頼できる高品質の AI 機能を作成することです。完璧なテスト カバレッジが優れたソフトウェア製品であることを保証しないのと同じように、完璧な成熟度スコアは優れた AI であることを保証しません。成熟度モデルは、そこに到達するための実践です。これは、測定、評価、継続的改善の規律であり、API に接続するだけでなく、専用の投資が必要です。
直観に反するかもしれないことは次のとおりです。AI の品質への投資は、私たちのスピードを遅らせるものではありません。それは私たちを速くします。深い可観測性、厳選されたデータセット、および調整された評価者を確立した後は、変更を出荷して、それが状況を改善したか、それとも破壊したかを数週間ではなく数時間以内に知ることができます。お客様が発見する前に、私たちは回帰を発見します。
顧客データに基づいたトレーニングは行っていません
これは繰り返して言う価値があります。Vanta は顧客データに基づいて AI モデルをトレーニングしません。
私たちが行ったすべての品質向上 (調整したすべての評価器、厳選したすべてのデータセット、閉じたすべてのフィードバック ループ) は、モデルのトレーニングに顧客データを使用することなく達成されました。当社の品質は、エンジニアリングの規律、何千もの顧客との協働によって構築された深い専門知識、そして絶え間ない評価の実践によってもたらされます。
これが重要なのは、他のアプローチに必要なプライバシーのトレードオフなしに、コンプライアンスと信頼のワークフローを深く理解する AI システムのメリットを顧客が得られるからです。
実際にはどうなるか
当社の AI 品質成熟度モデルにより、チームは迅速に反復し、Vanta の AI 機能全体で目に見える改善を実現できます。ここにあります

いくつかの例:
AI を活用したコントロールと証拠のマッピング機能は、お客様が期待する品質基準を満たしていませんでした。チームは修正を推測するのではなく、成熟度モデルの戦略に従いました。広範な制御記述にわたる数千の証拠提案からなるオフラインの包括的な評価データセットを構築し、ベースライン実験を実行したところ、精度が 35% 未満であることが判明しました (再現率が高く、モデルの網が広すぎました)。
適切な評価設定が整備されていたため、チームは自信を持って迅速に反復処理を行うことができました。次の改善ラウンドでは精度が 78% になり、カバレッジを維持しながら精度が 2 倍以上になりました。
Vanta AI は、新しく追加されたコントロールに適切な証拠マッピングを提案します。
当社のトラスト センター チャットボットは、顧客が自身の信頼データに基づいて AI が生成した回答を使用してセキュリティに関するアンケートに回答するのに役立ちます。これらの回答が正確で、関連性があり、ブランドに合ったものであることを確認するために、チームは成熟度モデルの各次元にわたる堅牢な評価システムを構築しました。
3 つの品質ディメンションが定義され、応答ごとに Webhook 経由でトリガーされるオンライン評価として実装されました。
関連性 (回答はユーザーの質問に対応しているか?)
忠実さ (RAG パイプラインから取得したデータに基づいていますか?)
スタイリング (簡潔で、整った形式で、調和していますか?)
スタイリング評価者に関する評価者の調整演習により、評価者と SME 間の迅速な合意が改善され、最大 65 ～ 70% から最大 85 ～ 90% に改善されました。これは、当社の自動品質チェックが対象分野の専門家の意見とほぼ一致することを意味します。
フィードバックに基づく継続的な改善の自動化により、SME レビューのスコアの低い回答が明らかになり、フィードバック パターンに基づいて評価者の調整データセットが改善されます。これにより、自己強化ループが作成されます。

システムは時間の経過とともに賢くなり続けます。
UI での親指の評価からの人間のフィードバックは評価トレースに直接組み込まれ、実際のユーザー満足度のシグナルを継続的に収集します。
その結果、トラスト センター チャットボットは、成熟度モデルの 5 つの側面のうち 3 つ (注釈とヒューマン フィードバック、可観測性、評価者) で黄色から緑色に移行し、エージェントのトラスト センター エクスペリエンスに向けて進化する中で、次の飛躍に向けた位置にありました。
自分自身に問うべき本当の質問
本当の質問は、「LLM はこれに答えられるか?」ということではありません。それは、「プロンプトやコンテキストに関係なく、LLM はこれを確実に、正確に、安全に実行できるか?」ということです。
DIY 統合により、プロンプトに応答できるモデルが得られます。 Vanta を使用すると、綿密なプロンプト エンジニアリング、コンテキスト設計、およびドメインの専門知識によって形成された AI が得られるため、自分でエンジニアリングしなくても適切な答えが得られます。反復するたびに改善され続ける高品質のシステムが得られます。コンプライアンス体制全体にわたって継続的な監視と可観測性が得られます。また、長期にわたってあなた (および監査人) の自信を保つ記録システムを手に入れることができます。
私たちはLLMと競合しているわけではありません。私たちはそれらの上に成り立っています。私たちが付加する価値は、品質層、ドメインの専門知識、

[切り捨てられた]

## Original Extract

The gap between a raw LLM and production-quality AI is enormous. Here's the five-dimension framework Vanta built to close it.

The Vanta AI Quality Eval Maturity Model | Vanta
Platform
Products
Platform
Compliance Get compliant quickly and painlessly with automation. Continuous GRC Join the modern way to GRC. Personnel and Access Easily control user access and permissions. Risk Management Proactively manage risk to drive smarter decisions. Third Party Risk Management Manage vendor onboarding and security reviews in one place. Questionnaire Automation Automate security questionnaire responses. Trust Center Showcase your compliance status and documentation. Streamlined audits Automate audit prep and evidence collection. Customer Commitments Centralize, track and act on every customer commitment.
Vanta AI Automate compliance and uncover insights with AI. Agentic Trust Platform Build and prove trust from a single, unified platform. Integrations Automatically pull data from 400+ tools. Vanta API Build custom integrations and workflows. NEW RELEASE See what's new from
Vanta Delivers Learn more
PRODUCTS Compliance Get compliant quickly and painlessly with automation. Personnel and Access Easily control user access and permissions. Risk Management Proactively manage risk to drive smarter decisions. Third Party Risk Management Manage vendor onboarding and security reviews in one place. Questionnaire Automation Automate security questionnaire responses. Trust Center Showcase your compliance status and documentation. Streamlined audits Automate audit prep and evidence collection. Customer Commitments Centralize, track and act on every customer commitment.
Vanta AI Automate compliance and uncover insights with AI. PLATFORM See an interactive demo
Agentic Trust Platform Build and prove trust from a single, unified platform. Integrations Automatically pull data from [integrations_count] tools. Vanta API Build custom integrations and workflows. Solutions
Size
Industry
Frameworks
Find a partner
Startups Automate compliance so you can keep building. Mid-market Expand your security and compliance program as
[truncated]
Accelerating security solutions for small businesses
‍
Tagore offers strategic services to small businesses.
A partnership that can scale
‍
Tagore prioritized finding a managed compliance partner with an established product, dedicated support team, and rapid release rate.
Standing out from competitors
‍
Tagore's partnership with Vanta enhances its strategic focus and deepens client value, creating differentiation in a competitive market.
This blog is part of our Trustcraft series, in which we dig into Vanta’s approach to building with AI. Read the first blog in this series to learn more about how we define Trustcraft.
You've seen what ChatGPT and Claude can do. You've heard about MCPs, CLIs, and APIs that let you wire a foundation model into just about anything. So when you look at Vanta's AI features, a fair question might come up: Why not just connect a general-purpose LLM to your own company’s data sources and call it a day?
It's a reasonable instinct. But after building AI systems that serve thousands of customers across compliance, trust, and security workflows, we've learned something that isn't obvious from the outside: The gap between a raw LLM integration and a production-quality AI product is enormous, and it widens as the stakes go up.
This post is about what lives in that gap.
The ‘just connect an LLM’ illusion
Foundation models are extraordinarily capable. But capability and reliability are not the same thing. In compliance and security, “good” means correctly interpreting control requirements, evaluating evidence accurately, handling regulatory edge cases, and never hallucinating details that could put your audit at risk. A raw LLM integration gives you a coin flip on each of these. Vanta's AI is engineered to get them right, repeatedly, at scale.
The difference isn't the model, but everything around the model—the prompts, the context retrieval, the memory, the domain-specific scaffolding, and the system design that turns a general-purpose model into a compliance-ready tool. When you use Vanta's AI, you benefit from deep, focused work put into every one of those layers.
When you use Vanta, you don't need to be a prompt engineer, you don't need to design your own retrieval system, and you don't need to stitch together context about your controls, evidence, and frameworks. It just works. When you connect ChatGPT or Claude to an API yourself, you're responsible for building all of that, or accepting worse answers.
Meet the framework that holds every Vanta AI feature to the same rigorous quality bar
To ensure we’re producing production-quality AI products, we don’t just ship anything. However, as our AI portfolio grew, we noticed a problem: Our AI teams didn't have a shared understanding of what "good enough" looked like for quality and evaluation. Teams were building great features but evaluating them inconsistently—different standards, different tools, different levels of rigor.
So we built a rigorous, multi-dimensional quality system we call our AI Quality Eval Maturity Model, which now governs how every AI-powered capability at Vanta is developed, measured, and improved over time.
The model evaluates our AI systems across five critical dimensions, each representing the ideal state we hold our teams to:
Observability : Full trace coverage across every AI interaction—inputs, outputs, reasoning steps, and metadata—with real-time monitoring and automated alerting when behavior drifts.
Curated evaluation datasets : Versioned, actively maintained datasets curated by our GRC subject matter experts. Not ad-hoc test sets, but evolving collections that reflect real-world complexity.
Calibrated evaluators : Formal evaluation pipelines, including validated LLM-as-a-judge systems, calibrated against the expertise of our subject matter experts and Vanta's deep trust and compliance domain knowledge. This ensures our quality assessments are consistent and aligned with what actually matters in production, not just what a model thinks is correct.
Systematic experimentation : A structured, repeatable experiment-and-analysis cycle for every AI change, with clear criteria for measuring impact and determining next steps.
Integrated feedback loops : Explicit and implicit user feedback are automatically captured, tagged, and linked back to evaluation datasets, creating a continuous cycle where real customer experiences drive improvement.
How our work has changed with the Eval Maturity Model
We score every AI team across each of the five dimensions using a simple rating system: Red (Foundational/Reactive), Yellow (Systematic/Developing), and Green (Advanced/Proactive).
When we first ran this assessment, the picture was humbling. Most teams were deep in red and yellow, relying on ad-hoc datasets, inconsistent evaluation processes, and manual feedback loops. We had the same gaps that most organizations building with AI have today.
But that's exactly why the model exists. We standardized our evaluation tooling so every AI team could leverage the same infrastructure, educated teams on evaluation best practices, and had our AI platform team partner closely with each product team to help them level up.
The result: Teams that were once deep in red and yellow have reached green across many dimensions.
It’s worth noting that reaching "all green" across the maturity model isn't the end goal. The real end goal is to produce high-quality AI features that customers can trust. In the same way that perfect test coverage doesn't guarantee a great software product, a perfect maturity score doesn't guarantee great AI. The maturity model is the practice that gets us there. It’s a discipline of measurement, evaluation, and continuous improvement that requires dedicated investment, not just plugging into an API.
Here's what might be counterintuitive: Investing in AI quality doesn't slow us down. It makes us faster. After establishing deep observability, curated datasets, and calibrated evaluators, we can ship a change and know within hours, not weeks, whether it improved things or broke them. We catch regressions before customers do.
We don't train on customer data
This is worth restating: Vanta does not train AI models on customer data.
Every quality improvement we've made—every evaluator we've calibrated, every dataset we've curated, every feedback loop we've closed—has been achieved without using customer data for model training. Our quality comes from engineering discipline, deep domain expertise built from working with thousands of customers, and a relentless evaluation practice.
This matters because customers get the benefits of an AI system that deeply understands compliance and trust workflows, without the privacy tradeoff that other approaches might require.
What this looks like in practice
Our AI Quality Maturity Model enables our teams to iterate quickly and deliver measurable improvements across Vanta's AI features. Here are a few examples:
Our AI-powered control-to-evidence mapping feature wasn't meeting the quality bar our customers deserved. Rather than guessing at fixes, the team followed the maturity model playbook: They built an offline, comprehensive evaluation dataset of thousands of evidence suggestions across an extensive range of control descriptions, then ran a baseline experiment that revealed precision was below 35% (with high recall—the model was casting too wide a net).
With a proper evaluation setup in place, the team was able to iterate rapidly and with confidence. The next round of improvements took precision to 78%, more than doubling accuracy while maintaining coverage.
Vanta AI will suggest the right evidence mapping to the newly added control.
Our Trust Center chatbot helps customers respond to security questionnaires using AI-generated answers grounded in their own trust data. To make sure those answers stay accurate, relevant, and on-brand, the team built out a robust evaluation system across the maturity model dimensions:
Three quality dimensions were defined and instrumented as online evaluators triggered via webhooks on every response:
Relevance (does the answer address the user's query?)
Faithfulness (is it grounded in retrieved data from our RAG pipeline?)
Styling (is it concise, well-formatted, and on-tone?)
An evaluator alignment exercise on the styling evaluator refined the prompt and improved evaluator-to-SME agreement from ~65–70% to ~85–90%, meaning our automated quality checks now closely match what our subject matter experts would say.
Continuous improvement automations based on feedback surface low-scoring responses for SME review, and improve evaluator alignment datasets based on feedback patterns. This creates a self-reinforcing loop where the system keeps getting smarter over time.
Human feedback from thumbs up/down ratings in the UI is wired directly into our evaluation traces, capturing real user satisfaction signals continuously.
The result: The Trust Center chatbot moved from yellow to green across three of the five maturity model dimensions—Annotations & Human Feedback, Observability, and Evaluators—and is now positioned for the next leap forward as we evolve toward an agentic Trust Center experience.
The real question to ask yourself
The real question isn't, "Can an LLM answer this?" It's "Can an LLM do this reliably, accurately, and safely, no matter the prompt or context?"
With a DIY integration, you get a model that can respond to a prompt. With Vanta, you get AI that's been shaped by deep prompt engineering, context design, and domain expertise, so you get good answers without having to engineer them yourself. You get a quality system that keeps getting better with every iteration. You get continuous monitoring and observability across your compliance posture. And you get a system of record that keeps you (and your auditors) confident over time.
We're not competing with LLMs. We're built on top of them. The value we add is the quality layer, the domain expertise, an

[truncated]
