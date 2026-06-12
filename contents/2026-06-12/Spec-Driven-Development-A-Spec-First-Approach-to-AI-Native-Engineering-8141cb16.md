---
source: "https://developer.microsoft.com/blog/spec-driven-development-ai-native-engineering"
hn_url: "https://news.ycombinator.com/item?id=48503856"
title: "Spec-Driven Development: A Spec-First Approach to AI-Native Engineering"
article_title: "Spec-Driven Development: A Spec-First Approach to AI-Native Engineering - Microsoft for Developers"
author: "sambcui"
captured_at: "2026-06-12T16:08:58Z"
capture_tool: "hn-digest"
hn_id: 48503856
score: 2
comments: 0
posted_at: "2026-06-12T13:30:00Z"
tags:
  - hacker-news
  - translated
---

# Spec-Driven Development: A Spec-First Approach to AI-Native Engineering

- HN: [48503856](https://news.ycombinator.com/item?id=48503856)
- Source: [developer.microsoft.com](https://developer.microsoft.com/blog/spec-driven-development-ai-native-engineering)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T13:30:00Z

## Translation

タイトル: スペック駆動開発: AI ネイティブ エンジニアリングへのスペックファースト アプローチ
記事のタイトル: 仕様駆動開発: AI ネイティブ エンジニアリングへの仕様優先アプローチ - 開発者向け Microsoft
説明: AI によりソフトウェアの配信は高速化されましたが、速度だけではより良い結果が保証されるわけではありません。チームが AI ネイティブ開発を採用する際の本当の課題は、

記事本文:
メインコンテンツにスキップ
マイクロソフト
開発者ブログ
開発者ブログ
開発者ブログ
ホーム
開発者
開発者向けのマイクロソフト
Microsoft Entra ID 開発者
プラットフォームとツール
ビジュアルスタジオ
スペック駆動開発: AI ネイティブ エンジニアリングへのスペックファーストのアプローチ
スペック駆動開発: AI ネイティブ エンジニアリングへのスペックファーストのアプローチ
AI によりソフトウェアの配信は高速化されましたが、速度だけでより良い結果が保証されるわけではありません。チームが AI ネイティブ開発を採用する際の本当の課題は、最終結果が元の意図を反映したままになるように、要件、設計、実装、検証を調整し続けることです。スペック駆動開発 (SDD) は、構造化されたスペックを人間と AI の両方にとって真実の共有ソースにすることで、この問題に対処します。最初にプロンプ​​トを出し、後で調整するのではなく、チームは最初に調整し、AI が明確な仕様に基づいて実行を加速できるようにします。
なぜ AI 支援開発は依然として失敗するのか
チームは、機能するソフトウェアを出荷することがよくありますが、依然として本来の意図を逸脱しています。問題はコードの品質だけではありません。これは、アイデアが利害関係者のニーズから要件、アーキテクチャ、実装、検証へと移行するときに意味が失われることです。
翻訳損失は通常、次の 4 つの場所で発生します。
利害関係者の製品要件に対するニーズ
アーキテクチャとデザインに対する要件
実装から検証、リリースまで
意図を保持する共有アーティファクトがなければ、すべてのハンドオフは解釈のステップになります。 AI はこれらのステップを加速できますが、決して解決されなかった曖昧さを修正することはできません。
プロンプトファーストのワークフローでは不十分な理由
プロンプトファーストのワークフローは単純なタスクにはうまく機能しますが、範囲と複雑さが増すにつれて困難になることがよくあります。
要件、制約、およびエッジケースがプロンプト内にのみ存在する場合、チームは信頼できる信頼できる情報源がなくても、迅速な出力を得ることができます。それはアーキテクチャのドリフトにつながります、タラ

人やツール間で前提が異なる場合、逸脱、一貫性のない実装、厳しいレビュー、およびやり直しが発生します。
仕様優先のワークフローはそのダイナミックさを変えます。 AI に散在するプロンプトから意図を推測するよう依頼するのではなく、チームは明示的に意図を定義し、AI を使用してそれに対して実行します。その結果、調整が改善され、より迅速な配送が可能になります。
スペック駆動開発とは何ですか?
仕様駆動開発 (SDD) は仕様優先のアプローチです。チームは共通のガードレール、要件、制約、受け入れ基準、エッジケースを事前に定義し、AI を使用してその共有コンテキストからコード、テスト、サポートするアーティファクトを生成します。
実際には、スペックはライフサイクル全体にわたる結合組織となります。ビジネスの意図をアーキテクチャ、実装、テスト、検証に結び付けることで、AI によって生成された出力が同じコンテキストに基づいたままになります。
チームが SDD を採用するのは、実装前の明確さが向上し、AI が作業するための強力な基盤が得られるからです。最大のメリットは実際的なものです。
要件が早期に明確になるため、曖昧さややり直しが少なくなります。
信頼できる情報源を共有することで、製品、エンジニアリング、テスト全体の連携が向上します。
AI が構造化されたコンテキストに基づいて生成できるため、実装が迅速化されます。
検証が元の意図に結び付けられるため、配信がより予測可能になります。
実際には、SDD はチームがどこに労力を注ぐかによって変化します。意図の明確化と事前の計画により多くの時間が費やされ、下流の手戻り作業に費やされる時間は減ります。
プロダクト マネージャーはシナリオと制約の定義を支援し、アーキテクトは計画モデルを形成し、エンジニアは AI を使用して実装を加速し、受け入れ基準が最初から明示されているため移行を早期にテストします。その結果、分断された成果物ではなく、共有された意図を中心としたワークフローが実現します。
SDD の背後にある考え方のみが重要です

チームがそれらを日々のエンジニアリング作業に一貫して適用できるかどうか。ここでツールキットが重要になります。ツールキットは、共有された意図、明示的な仕様、早期検証の原則を、チームが導入して拡張できる実用的なワークフローに変換します。
GitHub Spec Kit は、チームが SDD を実践するのに役立つツールキットです。 Microsoft が作成したオープンソース ツールで、要件を計画、実装タスク、検証ステップに変換するための構造化されたワークフローを提供し、GitHub Copilot などの AI コーディング ツールと連携して機能します。
github/spec-kit の詳細については、リンクを参照してください: 💫 仕様駆動開発の開始に役立つツールキット 。
私たちは昨年、ブログ記事「Diving Into Spec-Driven Development With GitHub Spec Kit – Microsoft for Developers」で Spec Kit を紹介しました。 Spec Kit の使用方法を簡単に学習するために参照してください。
GitHub 仕様キット エンジニアリングのライフサイクル
ライフサイクルはシンプルです。意図を定義し、曖昧さを取り除き、制約を設けて計画し、AI を使用して実装し、仕様に照らして検証します。
憲法 - 原則、基準、ガードレールを定義します。
指定 – 要件、シナリオ、および受け入れ基準をキャプチャします。
明確にする – 曖昧さ、依存関係、特殊なケースを解決します。
計画 – 意図をアーキテクチャ、フロー、制約に変換します。
タスク – 作業を実装可能な単位に分割します。
実装 – AI を使用してコードとテストを生成および改良します。
検証 – 出力が仕様と一致していることを確認します。
各ステップが次のステップを強化するため、ワークフローがより予測可能になり、チーム間での拡張が容易になります。
チームやユースケース全体で、いくつかの実践的な教訓が際立ちました。
調整はチームの習慣であり、単なるツールの選択ではありません。
優れた仕様は、構造だけでなく、意図、制約、および受け入れ基準を捉えています。
企画が大きすぎる

d 実装品質への影響。
通常、早期に明確性を高めると、総納期が短縮されます。
すべての変更にライフサイクル全体が必要なわけではないため、適切な規模で導入する必要があります。
結果の品質は仕様の品質と密接に関係しています。要約すると、スペック品質 = 出力品質です。
例: 繰り返されるオンボーディングをスケーラブルなパターンに変える
ブラウンフィールド プロジェクトの 1 つでは、チームは、新しいアセット タイプのオンボーディングは同じ基本フローに従っているものの、UI、API、テストの変更が繰り返し必要であることを認識しました。
パラメータ化された仕様で再利用可能なパターンをキャプチャし、各新しい資産の差異のみを文書化することで、構成主導のモデルに移行し、オンボーディング時間を 2 ～ 3 週間から数日に短縮しました。
例: 複雑なマルチサービス プラットフォームの調整
大規模なグリーンフィールド プロジェクトの 1 つでは、SDD は、参加者、施設、セキュリティ、ベンダー、物流、コンプライアンスにわたる数千の可動部分にわたるグローバルに分散されたプラットフォームを構築する前に、チームが共通の語彙を中心に PM、アーキテクト、エンジニアを調整するのに役立ちました。
構成、仕様、計画を真実の情報源として扱うことで、チームはサービス間の一貫性を向上させ、アーキテクチャ上の制約を明確にし、実装がコンポーネント間でスケールされるにつれて実行のチャーンを削減しました。
例: プロトタイプから実用的な製品への移行をより迅速に行う
別のブラウンフィールド プロジェクトでは、チームは SDD を使用して、React と TypeScript のプロトタイプから、DRI、プロビジョニング、ポリシー用の複数のエージェントに加え、接続状態の監視と管理ダッシュボードを備えた実用的な製品に移行しました。
構造化されたワークフローにより、AI が生成した出力を目に見える UI の動作と照らし合わせて検証することが容易になり、チームはカスタム プロンプトと品質ゲート スクリプトの導入を強化して、プロセスをより繰り返し実行できるようにしました。

貢献者間で共有できます。
チームが SDD を始める方法
チームは SDD ライフサイクル全体を一度に採用する必要はありません。多くの場合、小規模で集中的なパイロットは、実際に何が機能するかを学ぶための最良の方法です。以下は、導入のためのシンプルな 4 ステップのプレイブックです。
パイロット – 調整の問題が明らかな 1 つの機能またはワークフローから開始します。
形式化 – シナリオ、制約、および受け入れ基準を取り込んだ軽量仕様を作成します。
反復 – AI を使用して、共有コンテキストから実装成果物を生成します。
調整とスケール – 仕様に照らして出力をレビューし、学習しながらワークフローを調整します。
最初はプロセスを軽量にしておきます。仕様を生きた成果物として扱い、早期に過剰に仕様を指定することを避け、明確な価値を追加する場合にのみワークフローを拡張します。
ソフトウェア エンジニアリングは、AI 支援タスクから AI ネイティブ ワークフローに移行しています。この変化が続くにつれ、制限要因はコードをいかに速く生成できるかということではなくなりました。ライフサイクル全体にわたって、意図をいかに明確に捉え、共有し、検証できるかが重要です。
SDD は、チームに翻訳ロスを削減し、共有された意図に沿って調整し、ライフサイクル全体で AI をより予測可能に使用するための実用的な方法を提供し、より迅速な配信、より良い品質、より強力な調整を可能にします。
Apoorv Gupta は、Microsoft の主任ソフトウェア エンジニアです。
ルディ・ラルノ -->
ルディ・ラルノ
-->
2026 年 6 月 11 日
0
-->
Rudi Larno によるコメントを折りたたむ
-->
Rudi Larno によるコメントへのリンクをコピーします
UBB によるコストの爆発的な増加を考慮すると、SDD は依然として効率的な方法なのでしょうか?現在の一般的なコストと PRU のコスト モデルに関するデータはありますか?
エージェント拡張機能は実際に機能していますか?
あなたのエージェントは 2020 年からのプロジェクトの足場を築きました
教育者のトレーニングと能力開発
学生と保護者向けの特典
AI マーケットプレイス アプリのサポート
プライバシーに関する選択
消費者健康プライバシー

y
サイトマップ

## Original Extract

AI has made software delivery faster, but speed alone does not guarantee better outcomes. As teams adopt AI-native development, the real challenge is

Skip to main content
Microsoft
Dev Blogs
Dev Blogs
Dev Blogs
Home
Developer
Microsoft for Developers
Microsoft Entra Identity Developer
Platform and Tools
Visual Studio
Spec-Driven Development: A Spec-First Approach to AI-Native Engineering
Spec-Driven Development: A Spec-First Approach to AI-Native Engineering
AI has made software delivery faster, but speed alone does not guarantee better outcomes. As teams adopt AI-native development, the real challenge is keeping requirements, design, implementation, and validation aligned so the final result still reflects the original intent. Spec-Driven Development (SDD) addresses this by making structured specs the shared source of truth for both humans and AI. Instead of prompting first and aligning later, teams align first and let AI accelerate execution from a clear spec.
Why AI-assisted development still breaks down
Teams often ship software that works but still misses the original intent. The problem is not just code quality. It is the loss of meaning as ideas move from stakeholder needs to requirements, architecture, implementation, and validation.
Translation loss usually appears in four places:
Stakeholder needs to product requirements
Requirements to architecture and design
Implementation to validation and release
Without a shared artifact that preserves intent, every handoff becomes an interpretation step. AI can accelerate those steps, but it cannot correct ambiguity that was never resolved.
Why prompt-first workflows are not enough
Prompt-first workflows can work well for simple tasks, but they often struggle as scope and complexity increase.
When requirements, constraints, and edge cases live only in prompts, teams get fast output without a durable source of truth. That leads to architectural drift, code drift, inconsistent implementations, harder reviews, and rework when assumptions differ across people or tools.
A spec-first workflow changes that dynamic. Instead of asking AI to infer intent from scattered prompts, teams define intent explicitly and use AI to execute against it. The result is faster delivery with better alignment.
What is Spec-Driven Development?
Spec-Driven Development (SDD) is a spec-first approach. Teams define common guardrails, requirements, constraints, acceptance criteria, and edge cases up front, then use AI to generate code, tests, and supporting artifacts from that shared context.
In practice, the spec becomes the connective tissue across the lifecycle. It links business intent to architecture, implementation, tests, and validation so that AI-generated output stays grounded in the same context.
Teams adopt SDD because it improves clarity before implementation and gives AI a stronger foundation to work from. The biggest benefits are practical:
Less ambiguity and rework because requirements are clarified earlier.
Better alignment across product, engineering, and test through a shared source of truth.
Faster implementation because AI can generate against structured context.
More predictable delivery because validation is tied back to the original intent.
In practice, SDD changes where teams invest effort. More time goes into clarifying intent and planning up front, and less time is lost to downstream rework.
Product managers help define scenarios and constraints, architects shape the planning model, engineers use AI to accelerate implementation, and test shifts earlier because acceptance criteria are explicit from the start. The result is a workflow centred on shared intent rather than disconnected artifacts.
The ideas behind SDD matter only if teams can apply them consistently in day-to-day engineering work. That is where the toolkit becomes important: it translates the principles of shared intent, explicit specs, and early validation into a practical workflow teams can adopt and scale.
GitHub Spec Kit is the toolkit that helps teams put SDD into practice. An open-source tools, created by Microsoft, it provides a structured workflow for turning requirements into plans, implementation tasks, and validation steps that work well with AI coding tools such as GitHub Copilot.
Refer to the link for more details on github/spec-kit: 💫 Toolkit to help you get started with Spec-Driven Development .
We introduced Spec Kit last year in the blog post Diving Into Spec-Driven Development With GitHub Spec Kit – Microsoft for Developers . Refer to it for a quick start on learning how to use Spec Kit.
GitHub Spec Kit Engineering lifecycle
The lifecycle is simple: define intent, remove ambiguity, plan with constraints, implement with AI, and validate against the spec.
Constitution – Define principles, standards, and guardrails.
Specify – Capture requirements, scenarios, and acceptance criteria.
Clarify – Resolve ambiguity, dependencies, and edge cases.
Plan – Translate intent into architecture, flows, and constraints.
Tasks – Break the work into implementation-ready units.
Implement – Use AI to generate and refine code and tests.
Validate – Verify that the output matches the spec.
Each step reinforces the next, which makes the workflow more predictable and easier to scale across teams.
Across teams and use cases, a few practical lessons stood out:
Alignment is a team habit, not just a tooling choice.
Good specs capture intent, constraints, and acceptance criteria, not just structure.
Planning has an outsized impact on implementation quality.
More clarity early usually reduces total delivery time.
Not every change needs the full lifecycle, so adoption should be right-sized.
The quality of the outcome is closely tied to the quality of the spec. In summary, Spec quality = output quality.
Example: turning repeated onboarding into a scalable pattern
In one of the brownfield projects, the team recognized that onboarding new asset types followed the same basic flow but repeatedly required UI, API, and test changes.
By capturing the reusable pattern in parameterized specs and documenting only the deviations for each new asset, they shifted to a configuration-driven model that reduced onboarding time from 2–3 weeks to a few days .
Example: coordinating a complex multi-service platform
In one of the big greenfield projects, SDD helped the team align PMs, architects, and engineers around a shared vocabulary before building a globally distributed platform spanning thousands of moving parts across attendees, facilities, security, vendors, logistics, and compliance.
By treating the constitution, specs, and plans as the source of truth, the team improved cross-service consistency, made architectural constraints explicit, and reduced execution churn as implementation scaled across components.
Example: moving from prototype to working product faster
In another brownfield project, the team used SDD to move from a React and TypeScript prototype to a working product with multiple agents for DRI, Provisioning, and Policy, plus connectivity health monitoring and admin dashboards.
The structured workflow made it easier to validate AI-generated output against visible UI behaviour, and the team strengthened adoption with custom prompts and quality-gate scripts that made the process more repeatable across contributors.
How teams can get started with SDD
Teams do not need to adopt the full SDD lifecycle at once; a small, focused pilot is often the best way to learn what works in practice. Below is a simple 4-step playbook for adoption.
Pilot – Start with one feature or workflow where alignment problems are visible.
Formalize – Create a lightweight spec that captures scenarios, constraints, and acceptance criteria.
Iterate – Use AI to generate implementation artifacts from that shared context.
Refine and scale – Review the output against the spec and refine the workflow as you learn.
Keep the process lightweight at first. Treat specs as living artifacts, avoid over-specifying too early, and expand the workflow only where it adds clear value.
Software engineering is moving from AI-assisted tasks toward AI-native workflows. As that shift continues, the limiting factor is no longer how quickly code can be generated. It is how clearly intent can be captured, shared, and validated across the lifecycle.
SDD gives teams a practical way to reduce translation loss, align around shared intent, and use AI more predictably across the lifecycle—enabling faster delivery, better quality, and stronger alignment.
Apoorv Gupta is a Principal Software Engineer at Microsoft.
Rudi Larno -->
Rudi Larno
-->
June 11, 2026
0
-->
Collapse comment by Rudi Larno
-->
Copy link to comment by Rudi Larno
Given the explosions in cost because of UBB, is SDD still an efficient method? Is there some data available on the typical cost today vs the PRU cost model?
Is your agent extension actually working?
Your agent just scaffolded a project from 2020
Educator training and development
Deals for students and parents
Support for AI marketplace apps
Your Privacy Choices
Consumer Health Privacy
Sitemap
