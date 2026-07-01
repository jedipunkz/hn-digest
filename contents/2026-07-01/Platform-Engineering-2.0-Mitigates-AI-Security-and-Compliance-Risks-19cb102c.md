---
source: "https://platformengineering.com/features/platform-engineering-2-0-manage-ai-costs-and-risks-without-rebuilding-infrastructure/"
hn_url: "https://news.ycombinator.com/item?id=48747895"
title: "Platform Engineering 2.0 Mitigates AI Security and Compliance Risks"
article_title: "Platform Engineering 2.0: Manage AI Costs and Risks Without Rebuilding Infrastructure - Platform Engineering"
author: "BruceGain"
captured_at: "2026-07-01T14:57:43Z"
capture_tool: "hn-digest"
hn_id: 48747895
score: 1
comments: 0
posted_at: "2026-07-01T14:56:26Z"
tags:
  - hacker-news
  - translated
---

# Platform Engineering 2.0 Mitigates AI Security and Compliance Risks

- HN: [48747895](https://news.ycombinator.com/item?id=48747895)
- Source: [platformengineering.com](https://platformengineering.com/features/platform-engineering-2-0-manage-ai-costs-and-risks-without-rebuilding-infrastructure/)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T14:56:26Z

## Translation

タイトル: プラットフォーム エンジニアリング 2.0 は AI のセキュリティとコンプライアンスのリスクを軽減します
記事のタイトル: プラットフォーム エンジニアリング 2.0: インフラストラクチャを再構築せずに AI のコストとリスクを管理する - プラットフォーム エンジニアリング
説明: プラットフォーム エンジニアリング 2.0 は進化です。製品としてのプラットフォーム、ゴールデン パス、シフトレフト セキュリティといった基盤は、依然として不可欠です。

記事本文:
コンテンツにスキップ
ナビゲーションを切り替えます
ビデオ プラットフォーム エンジニアリング ショー
関連サイト テックストロンググループ
イノベーションと新興テクノロジー
プラットフォーム エンジニアリング 2.0: インフラストラクチャを再構築せずに AI のコストとリスクを管理
プラットフォーム エンジニアリング チームは、10 年間の大部分を、標準化された Kubernetes クラスター、CI/CD パイプライン、内部開発者プラットフォーム (IDP)、開発者がアプリケーションを安全、効率的、繰り返し出荷できるようにするセルフサービス インフラストラクチャなど、素晴らしいものの構築に費やしてきました。その基盤は、AI が登場してその根底にあるすべての前提を変えるまで、うまく機能しました。
DevOps 対プラットフォーム エンジニアリングをめぐる議論は、今では古風なものに感じられます。
AI ワークロードを運ぶように設計されていないインフラストラクチャ上で、AI ワークロードをどのように構築、管理、分離、運用するのかという、はるかに重大な課題がそれに取って代わりました。 Broadcom と PlatformEngineering.org は、Broadcom の新しい Platform Engineering 2.0 フレームワークに答えがあると述べています。 Broadcom は、フレームワークの作成に際し、AI、ソフトウェア、プライベート クラウド インフラストラクチャにわたる企業の集合的な経験を活用したと述べています。
プラットフォーム エンジニアリング 2.0 は、プラットフォーム エンジニアリング 1.0 の成功した基盤を置き換えるのではなく、その基盤の上に構築されています。このフレームワークは、既存のプラットフォームへの投資から離れるのではなく、自然な発展として機能するように設計されました。
Platform Engineering 1.0 は、コンテナ化された開発者中心の人間のペースのワークフロー向けに構築されました。 AI は 3 つの異なる方法でそのモデルを破壊しました。
まず、AI は、プロンプトと取得されたコンテンツをライブの予測不可能な入力チャネルとして導入します。データはもはや単なるデータではありません。データはモデルの出力動作に実行可能な影響を与え、従来の分離プロパティを作成します。これはかつては保持されていました。

コンピューティングとストレージ - 突然信頼性が低下します。
第 2 に、AI ワークロードは、SaaS API、微調整されたホスト モデル、オンプレミス推論、取得レイヤー、既存システムの奥深くまで到達するツール呼び出しエージェントなど、複数の実行プレーンに同時にまたがります。プラットフォームは、これらすべてを同時に管理するように設計されていませんでした。
3 番目に、そして最も重要なことは、AI は信頼境界をアプリケーションから遠ざけ、モデル、ツール、データ ソース、人間、およびそれらを操作する非人間エージェントの間の相互作用に移すことです。それは埋められるギャップではありません。それは構造的な問題です。 AI は、セキュリティとガバナンスの新たな懸念を引き起こすだけでなく、チームが異なるモデル、ツール、取得アプローチ、可観測性の実践を個別に採用するため、運用の断片化も引き起こします。
その影響はすでに企業予算に影響を及ぼしています。 6月のFinOps Xカンファレンスで、アクセンチュアのFinOpsグローバル・プラクティス・リードであるマイク・アイゼンシュタイン氏は、クロードAPIのコストが1か月で1日あたり25万ドルから40万ドルまで増加しているというCIOの説明を伝えた。 FinOps Foundation のエグゼクティブ ディレクターである J.R. Storment 氏は、「AI の変化の速度は非常に速いです。ある日は良いポリシーでも、翌週には時代遅れになる可能性があります。」とはっきりと述べています。
これは持続可能ではなく、アプリケーションレベルの修正では解決できません。
アプリレベルの修正が拡張できない理由
本能的な反応は、各アプリケーション チームが AI ユースケースに独自のガードレールを巻き付けることでした。チャットボットは即時の強化を追加します。ドキュメントツールのボルトオンアクセスチェック。コードアシスタントは個別のログを取得します。すべてのチームは独自の境界線を構築します。
その反射には構造的な限界があります。ポリシーの解釈がチーム間で断片化する。たとえば、「外部モデルへの PII がない」ということは、マーケティングにおける意味とマーケティングにおける意味が異なります。

ファイナンス。セキュリティ リーダーは、どのモデルが、どこで、どのようなポリシーの下で実行されているかに関する基本的な質問に答えることができません。答えが数十のサービスやベンダー コンソールに分散しているためです。
シャドウ AI はこれをさらに悪化させます。データ漏洩のリスクだけでなく、シャドー IT が生み出すコスト プロファイルのせいで、これまでのシャドー IT よりも危険です。
ドキュメントとアプリケーション コードに限定された AI ガバナンスでは十分ではありません。セキュリティの責任はプラットフォーム自体にまで及ぶ必要があります。
これらの柱のうちの 2 つは基礎であり、コントロール プレーンとしてのモデル ガバナンスと、構造的保証としてのワークロード分離です。
コントロールプレーンとしてのモデルガバナンス
現在、ほとんどの企業は複数のプロバイダーにわたって複数のモデルを実行しています。 AI セキュリティ企業 Airia が指摘するように、「モデル固有のガバナンスは大規模に破壊されます」。それは、「モデル レベルの上に位置する制御層、つまりどの基礎モデルがタスクを実行しているかに関係なく、ポリシーを強制し、決定を記録し、動作を監視する制御層」を提案しています。
プラットフォーム エンジニアリング 2.0 は、その原則を具体的なサービス、つまり中央モデル レジストリとルーティング層に変えます。統一認証？ポリシーの適用は、OpenAI、Anthropic、またはあらゆるオンプレミス モデル全体に​​均一に適用され、監査、可観測性、コンプライアンスを単一画面で実現します。開発者はプラットフォームを通じてモデルへのアクセスをリクエストします。プラットフォームは、それらのリクエストをリスク層と承認ワークフローにマッピングします。これを AI としてのインフラストラクチャと考えてください。プラットフォームはルールブックと審判の両方になります。
構造的な保証としてのワークロードの分離
モデル ガバナンスが AI に何を実行できるかを定義する場合、ワークロードの分離は AI がそれを実行できる場所と障害がどこまで広がるかを定義します。つまり、実験用サンドボックス、内部ワークロード、

そして顧客向けの規制されたデータ環境。そして、それは、モデル、エージェント、ツールにバインドされたゼロトラスト サービス ID を意味します。なぜなら、即時注入は成功するでしょうし、その際の横方向の移動は構造的に困難でなければならないからです。
これら 2 つの柱は、より大きなものに収束します。自律型 AI エージェントは、新しいクラスのプラットフォーム ユーザーとして登場しつつあります。継承する以前のペルソナがなく、スコープの設定ミスを発見するために関与する人間もいません。プラットフォームは、第一級の消費者として人間の開発者と AI エージェントの両方を同時にサポートする必要があります。
プラットフォーム エンジニアリング 2.0 ホワイトペーパーには、「プラットフォーム エンジニアリングはもはやソフトウェア配信の分野ではありません。プラットフォーム エンジニアリングは、企業のエージェントの未来のための運用基盤になりつつあります。」と述べられています。
プラットフォーム変革の必須事項
プラットフォーム エンジニアリング 2.0 はリセットではなく進化です。製品としてのプラットフォーム、ゴールデン パス、セルフサービス IDP、シフトレフト セキュリティなどの基盤は依然として不可欠です。変わるのは、プラットフォームが誰にサービスを提供するか、何をしなければならないか、そしてどれだけ早く適応する必要があるかということです。
この移行をマスターしたチームは、デリバリー パイプラインよりも価値のあるものを保持しています。彼らは、組織の AI ネイティブな未来のための基盤を保持しています。エージェント AI をプラットフォームのコントロール プレーンに直接統合しない組織は、大幅に後れを取ることになります。
そこに到達するのは簡単ではありません。しかし、方向性は明確であり、セキュリティの暴露、運用の断片化、AI 支出の暴走といった遅延の代償はすでに目に見えています。プラットフォームは進化する必要があり、今がその始まりの時です。
セルフサービス プラットフォームは人間のために構築されました — AI エージェントがルールを変える
AI の勢いはインフラストラクチャの準備を上回ります
Kubernetes は AI コントロール プレーンになりつつありますが、それを機能させることができるのはプラットフォーム チームだけです
セキュリティとして

考え方: クラウドネイティブのインフラストラクチャ設計にセキュリティを組み込む方法

## Original Extract

Platform Engineering 2.0 is an evolution. The foundations — Platform as Product, golden paths and shift-left security — remain essential.

Skip to content
Toggle Navigation Latest
Videos The Platform Engineering Show
Related Sites Techstrong Group
Innovation & Emerging Technologies
Platform Engineering 2.0: Manage AI Costs and Risks Without Rebuilding Infrastructure
Platform engineering teams have spent the better part of a decade building something admirable: Standardized Kubernetes clusters, CI/CD pipelines, internal developer platforms (IDPs), and self-service infrastructure that let developers ship applications safely, efficiently, and repeatedly. That foundation held up well — until AI arrived and changed every assumption underneath it.
The debate over DevOps versus platform engineering now feels quaint.
A far more consequential challenge has taken its place: How do you build, govern, isolate, and operate AI workloads on infrastructure that was never designed to carry them? Broadcom and PlatformEngineering.org say there is an answer in Broadcom’s new Platform Engineering 2.0 framework . With its creation of the framework, Broadcom says it drew on its collective enterprise experience across AI, software and private cloud infrastructure when it created the framework.
Platform Engineering 2.0 also builds upon the successful foundations of Platform Engineering 1.0, rather than replacing them. The framework was designed to serve as a natural progression rather than a departure from existing platform investments.
Platform Engineering 1.0 was built for containerized, developer-centric, human-paced workflows. AI broke that model in three distinct ways.
First, AI introduces prompts and retrieved content as a live, unpredictable input channel. Data is no longer just data — it is an executable influence on model output behavior, making traditional isolation properties — which once held for compute and storage — suddenly unreliable.
Second, AI workloads span multiple execution planes at once: SaaS APIs, fine-tuned hosted models, on-premises inference, retrieval layers, and tool-calling agents that reach deep into existing systems. The platform was never designed to govern across all of those simultaneously.
Third, and most critically, AI moves the trust boundary away from the application — and into the interplay between models, tools, data sources, humans, and non-human agents steering them. That is not a gap you can patch. It is a structural problem. In addition to introducing new security and governance concerns, AI also creates operational fragmentation as teams independently adopt different models, tooling, retrieval approaches, and observability practices.
The consequences are already landing in enterprise budgets. At June’s FinOps X conference, Mike Eisenstein, Accenture’s FinOps Global Practice Lead, relayed a CIO’s account of Claude API costs escalating from $250,000 per day to $400,000 per day in a single month. As J.R. Storment, Executive Director of the FinOps Foundation, put it plainly: “AI’s rate of change is exceptionally fast. What’s a good policy one day can be out of date the next week.”
This is not sustainable — and application-level fixes won’t solve it.
Why App-Level Fixes Don’t Scale
The instinctive response has been to let each application team wrap its own guardrails around its AI use case. Chatbots add prompt hardening. Document tools bolt-on access checks. Code assistants get separate logging. Every team builds its own perimeter.
That reflex has structural limits. Policy interpretation fragments across teams. For example, “no PII to external models” means something different in marketing than it does in finance. Security leaders cannot answer basic questions about which models are running, where, and under what policies, because the answers are scattered across dozens of services and vendor consoles.
Shadow AI compounds this further. It is more dangerous than shadow IT ever was — not just because of data exposure risk, but because of the cost profile it creates.
AI governance confined to documents and application code is not enough. Security responsibilities must move down into the platform itself.
Two of those pillars are foundational: model governance as a control plane and workload isolation as a structural guarantee.
Model Governance as a Control Plane
Most enterprises now run multiple models across multiple providers. As the AI security firm Airia notes , “model-specific governance breaks at scale.” It proposes “a control layer that sits above the model level — one that enforces policy, logs decisions, and monitors behavior regardless of which underlying model is executing a task.”
Platform Engineering 2.0 turns that principle into a concrete service: a central model registry and routing layer; unified authentication? Policy enforcement applied uniformly across OpenAI, Anthropic, or any on-premises model, and a single pane of glass for audit, observability, and compliance. Developers request model access through the platform; the platform maps those requests to risk tiers and approval workflows. Think of it as infrastructure as AI — the platform becomes both the rulebook and the referee.
Workload Isolation as a Structural Guarantee
If model governance defines what AI is allowed to do, workload isolation defines where it can do it and how far failures can spread. That means dedicated isolation domains across experimental sandboxes, internal workloads, and customer-facing regulated data environments. And it means zero-trust service identities bound to models, agents, and tools — because prompt injections will succeed, and lateral movement must be structurally difficult when they do.
These two pillars converge on something larger. Autonomous AI agents are arriving as a new class of platform user — with no prior persona to inherit from and no human in the loop to catch a misconfigured scope. The platform must support both human developers and AI agents simultaneously, as first-class consumers.
“Platform engineering is no longer a software delivery discipline, Platform Engineering 2.0 whitepaper states. “It is becoming the operational foundation for the enterprise’s agentic future.”
The Platform Transformation Imperative
Platform Engineering 2.0 is an evolution, not a reset. The foundations — Platform as Product, golden paths , self-service IDPs, and shift-left security — remain essential. What changes is who the platform serves, what it must do, and how fast it must adapt.
The teams that master this transition hold something more valuable than a delivery pipeline. They hold the substrate for their organization’s AI-native future. Organizations that do not integrate agentic AI directly into their platform control plane will fall significantly behind.
Getting there will not be easy. But the direction is clear, and the cost of delay — in security exposure, operational fragmentation, and runaway AI spend — is already visible. The platform must evolve, and the time to start is now.
Your Self-Service Platform Was Built for Humans — AI Agents Change the Rules
AI Momentum Outpaces Infrastructure Readiness
Kubernetes is Becoming the AI Control Plane—But Only Platform Teams Can Make it Work
Security as a Mindset: How to Embed Security Into Cloud-Native Infrastructure Design
