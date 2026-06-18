---
source: "https://deepmind.google/blog/securing-the-future-of-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=48588072"
title: "Securing the Future of AI Agents"
article_title: "Securing internal systems against increasingly capable and imperfectly aligned AI — Google DeepMind"
author: "falcor84"
captured_at: "2026-06-18T18:56:48Z"
capture_tool: "hn-digest"
hn_id: 48588072
score: 13
comments: 2
posted_at: "2026-06-18T16:44:04Z"
tags:
  - hacker-news
  - translated
---

# Securing the Future of AI Agents

- HN: [48588072](https://news.ycombinator.com/item?id=48588072)
- Source: [deepmind.google](https://deepmind.google/blog/securing-the-future-of-ai-agents/)
- Score: 13
- Comments: 2
- Posted: 2026-06-18T16:44:04Z

## Translation

タイトル: AI エージェントの未来を確保する
記事のタイトル: ますます高機能になり不完全に調整された AI から内部システムを保護する — Google DeepMind

記事本文:
メイン コンテンツにスキップ 次世代 AI システムを探索する
AI の最新のブレークスルーとラボからの最新情報
AI で新たな発見の時代を切り開く
私たちの使命は、人類に利益をもたらすために責任を持って AI を構築することです
次世代 AI システムを探索する
AI の最新のブレークスルーとラボからの最新情報
AI で新たな発見の時代を切り開く
私たちの使命は、人類に利益をもたらすために責任を持って AI を構築することです
Google DeepMind DeepMind Google AI すべての AI について学ぶ Google DeepMind AI の最前線を探る Google Labs AI の実験を試す Google Research 研究を探索する 製品とアプリ Gemini アプリ Gemini とチャットする Google AI Studio 次世代 AI モデルで構築する Google Antigravity 当社のエージェント開発プラットフォーム モデル 研究 科学 Gemini について Build Gemini を試す 2026 年 6 月 18 日 責任と安全性 AI エージェントの未来を確保する
共有 コピーしました ますます能力が向上し不完全に調整された AI に対して社内システムをどのように保護しているか
AI エージェントは、テクノロジーと私たちの関係を変えています。これらのシステムは、サイバー防御から科学的発見、製品開発に至る複雑なタスクを自律的に実行することにより、生産性の新時代を切り開きます。米国だけでも、AI エージェントは 2030 年までに 2 兆 9,000 億ドルの経済価値を生み出す可能性があります。
これらのエージェントの能力が高まるにつれて、より高度な安全対策も必要になります。そのため、Google は AI コントロール ロードマップを開発しました。これは、Google 内に導入する高度な AI を構築および管理するためのフレームワークです。この「多層防御」アプローチは、より広範な業界のモデルとして機能し、従来のモデルの調整を超え、調整が不完全な場合でも保証を提供するシステムレベルのセキュリティの重要な層を追加します。
セキュリティに対する当社のアプローチは、強力な基盤から始まります。

サンドボックス、エンドポイント セキュリティ、即時インジェクション耐性などの従来の保護手段が組み込まれています。これに加えて、AI 制御ロードマップでは、モデルの調整、つまり本質的に安全で役に立つように AI をトレーニングすることを主な防御手段として使用します。内部エージェントを位置がずれている可能性があるものとして扱うことで、追加のセキュリティ層を提供し、位置合わせが不完全な場合でも保証を提供します。
デュアルコントロールを備えた運転インストラクターのようなものだと考えてください。インストラクターは生徒を信頼していますが、間違いが発生した場合はハンドルを握ったり、ブレーキを踏んだりする準備を整えておきます。同様に、当社の AI 制御システムは、検証された動作に基づいて AI エージェントに権限を付与し、制御された増分アクセスを通じて信頼を構築できるようにします。
私たちのロードマップは、シンプルだが慎重な想定に基づいて構築されています。つまり、高度な能力を備えた AI エージェントが私たちの目標と一致せず、予期しない方法で動作した場合はどうなるでしょうか?これに対処するために、私たちの計画は 3 つの主要分野に焦点を当てています。
まず、サイバーセキュリティのベスト プラクティスに従って、AI 用の新しい脅威モデリング フレームワークを開発しました。このロードマップでは、信頼できない AI エージェントを潜在的な「内部脅威」として扱います。これは、企業がすでにオフィスにアクセスしている潜在的な不正従業員を扱うのと同様の方法です。業界標準の MITRE ATT&CK フレームワーク（実世界の観察に基づいた敵の戦術とテクニックに関する世界的にアクセス可能な知識ベース）を基盤として構築することで、潜在的な攻撃をより小さな戦術とテクニックに分割します。これにより、リスクを体系的に追跡し、損害が発生する前に悪い行動を特定し、防御が実際に機能するかどうかをテストする訓練を安全に実行できるようになります。
次に、これらのリスクに対処するために AI 制御の緩和策を導入する計画を策定します。
問題の特定 (検出): 監視は AI 制御の根幹です。他の信頼できる AI システムを使用します

ems は「スーパーバイザー」として、作業中のエージェントの推論、行動、計画が軌道から外れていないかどうかを常にレビューします。
問題の管理 (予防と対応): エージェントが何か有害なことをしようとしているのをスーパーバイザが発見した場合、損害が発生する前に介入してその行動を阻止することができます。
最後に、システムが機能していることを確認するために、カバレッジ (監視されたトラフィックの割合)、リコール (捕捉された誤った動作の割合)、および応答までの時間という 3 つの重要な指標にわたるパフォーマンスを常に測定しています。
AI が賢くなるにつれてセキュリティを拡大
AI モデルが進化し続けるにつれて、私たちの防御も同時に強化する必要があります。 AI 制御ロードマップでは、セキュリティ プロトコルを、次の 2 つの重要な側面における AI 機能の測定可能なマイルストーンにマッピングしています。
検出を回避する能力: 現在、私たちはモデルの目に見える思考連鎖を監視しています。ただし、モデルが監視の認識 (監視されているという認識) または不透明な推論 (目に見えるテキストを使用しない推論) によって推論を隠すことを学習するため、単に言語化された推論を読み取るだけでは十分ではありません。それらの動作を分析し、場合によってはモデルの内部動作を検査する必要があります。
潜在的な危害に対する能力: リスクが低く、元に戻せるアクションの場合は、遅延を伴うトランスクリプトのレビューと問題発生後の迅速な修復 (非同期応答) に頼ることができます。大規模なサイバー攻撃など、特にリスクの高いアクションについては、有害なアクションが実行される前にブロックするリアルタイムの防御に移行する必要があります (同期対応)。言い換えれば、私たちの対応戦略は脅威の重大度に応じて拡大する必要があります。
AI 制御ロードマップでは、導入されたモデルの機能に基づいてセキュリティ対策を推奨しています。モデルの検出回避能力によって検出ファイルが決まります

vel (D1 ～ D4)、およびモデルの攻撃実行能力によって予防と対応レベル (R1 ～ R3) が決まります。
100万人のエージェントの軌跡を分析してライブ監視に情報を提供
これらのシステムを構築するにはデータが必要です。私たちのチームは、コーディング エージェントの軌跡をレビューし、人間によるレビューのために結果をエスカレーションする非同期監視システムの内部プロトタイプを構築しました。フラグが立てられた各イベントは、ロードマップの脅威分類に基づいて分類されます。
私たちはすでに 100 万件のコーディング エージェント タスクを分析しており、安全プロトコルを改良するための貴重な洞察が得られ、単純なキーワード フィルタリングを超えて高信号の動作パターンを特定できるようになりました。たとえば、この研究は Gemini Spark エージェントのライブ モニターの構築に役立ち、意図しないデータ削除などの新たな問題にリアルタイムで対応できるようになりました。
私たちのデータは、フラグが立てられたイベントの大部分が敵対的な意図から生じたものではないことを示していることに注意することが重要です。むしろ、エージェントの誤解や、ユーザーの目標を達成しようとする過度の熱意によって引き起こされることがよくあります。これらの微妙な違いを理解することは、安全性とセキュリティのプロトコルを改良するために重要です。
エージェントのエコシステムを守るのは共通の責任です
AI コントロール ロードマップは、エージェント セキュリティに対する多層アプローチの重要な要素の 1 つです。これらのプロトコルを当社の広範なインフラストラクチャと統合することで、当社は内部セキュリティを拡張し、最先端の AI モデルを安全に管理しています。
さらに、エージェントのセキュリティに対するこの多層的なアプローチは、産業界、政策立案者、学界が協力して優先すべき事項です。ベストプラクティスと標準を中心にエコシステムを調整することで、サイバー防御者に力を与え、社会の回復力を構築できます。そのため、本日、ポリシーマの技術フレームワークも公開します。

kers、「エージェント セキュリティの 3 層」。この文書では、個々のエージェントのレベルでセキュリティを向上させる必要がある方法について詳しく説明しています。マルチエージェントシステムの場合。そして、サイバー防御者に力を与え、より広範なエコシステム全体で回復力を構築します。
私たちは、これらのフレームワークに基づいて、将来に向けて安全な基盤の構築を継続しながら、現在、有能な AI を自信を持って導入するつもりです。
完全な技術レポートはこちらからお読みください:
メアリー・フォン、エリック・ジェンナー、ローラン・サイモン、ルイス・ホー、ロヒン・シャー、セバスチャン・ファーカー、スコット・コール。
フリン、アンカ・ドラガン、アラン・クーニー、ビラル・チュグタイ、バック・シュレゲリス、コディ・ワイルド、デヴィッド・リンドナー、ジュリアン・スタストニー、ケビン・クライマン、リー・ディン、ミリアム・カーン、ラルカ・エイダ・ポパ、ローランド・ジマーマン、ライアン・グリーンブラット、セントーラン・ラジャマノハラン、ヴィクトリア・クラコヴナ、クセルクセス・ドティワラの4名。
私は Google の利用規約に同意し、私の情報が Google のプライバシー ポリシーに従って使用されることを認めます。

## Original Extract

Skip to main content Explore our next generation AI systems
Our latest AI breakthroughs and updates from the lab
Unlocking a new era of discovery with AI
Our mission is to build AI responsibly to benefit humanity
Explore our next generation AI systems
Our latest AI breakthroughs and updates from the lab
Unlocking a new era of discovery with AI
Our mission is to build AI responsibly to benefit humanity
Google DeepMind DeepMind Google AI Learn about all our AI Google DeepMind Explore the frontier of AI Google Labs Try our AI experiments Google Research Explore our research Products and apps Gemini app Chat with Gemini Google AI Studio Build with our next-gen AI models Google Antigravity Our agentic development platform Models Research Science About Build with Gemini Try Gemini June 18, 2026 Responsibility & Safety Securing the future of AI agents
Share Copied How we’re securing internal systems against increasingly capable and imperfectly aligned AI
AI agents are transforming our relationship with technology. By autonomously executing complex tasks — from cyber defence to scientific discovery and product development — these systems are unlocking a new era of productivity. In the U.S alone, AI agents could create $2.9 trillion in economic value by 2030.
As these agents become more capable, they also require more sophisticated safeguards. That’s why we developed our AI Control Roadmap : a framework for building and managing the advanced AI we deploy within Google. This “defense-in-depth” approach, which could serve as a model for the wider industry, goes beyond traditional model alignment, adding a crucial layer of system-level security that provides assurance even if alignment is imperfect.
Our approach to security starts with a strong foundation, incorporating traditional safeguards like sandboxing, endpoint security, and prompt injection resistance. On top of this, the AI Control Roadmap uses model alignment, i.e. training AI to be inherently safe and helpful, as a primary defense. It provides an additional layer of security by treating internal agents as potentially misaligned, providing assurance even if alignment is imperfect.
Think of it like a driving instructor with dual controls. The instructor trusts the student but stays ready to take the wheel or hit the brakes if a mistake occurs. Similarly, our AI control system grants AI agents permissions based on their verified behavior, allowing us to build trust through controlled, incremental access.
Our roadmap is built on a simple but cautious assumption: what if a highly capable AI agent isn’t aligned with our goals and acts in unexpected ways? To address this, our plan focuses on three key areas.
First, following best practices in cybersecurity, we have developed a novel threat-modelling framework for AI. The roadmap treats untrusted AI agents as potential “insider threats” — similar to the way a company would treat a potential rogue employee who already has access to the office. By building on the industry-standard MITRE ATT&CK framework, a globally-accessible knowledge base of adversary tactics and techniques based on real-world observations, we break down potential attacks into smaller tactics and techniques. This allows us to methodically track risks, spot bad behavior before it causes damage and safely run drills to test if our defenses actually hold up.
Second, we develop a plan to deploy AI control mitigations to address these risks.
Identifying the problem (Detection): Monitoring is the backbone of AI control. We use other, trusted AI systems as "supervisors" to constantly review a working agent's reasoning, actions and plans to see if it is going off track.
Managing the problem (Prevention and Response): If a supervisor sees the agent trying to do something harmful, it can step in to block the action before damage occurs.
Finally, to ensure our systems are functional, we constantly measure performance across three critical metrics: coverage (the fraction of traffic monitored), recall (the fraction of misaligned behaviors caught), and time-to-response.
Scaling security as AI gets smarter
As AI models continue to advance, our defenses must also strengthen in tandem. In our AI Control Roadmap, we map security protocols to measurable milestones in AI capabilities on two critical fronts:
Ability to evade detection: Today, we monitor a model's visible chain-of-thought. However, as models learn to hide their reasoning — either through oversight awareness (knowing they are being monitored) or opaque reasoning (reasoning without using visible text) — simply reading their verbalized reasoning will not be enough. We will need to analyze their behavior and potentially inspect the model’s inner workings.
Ability for potential harm: For low-risk and reversible actions, we can rely on reviewing transcripts with a delay and remediating issues quickly after they occur (asynchronous response). For particularly high-risk actions, such as major cyber attacks, we must shift to real-time prevention, blocking harmful actions before they can be executed (synchronous response). In other words, our response strategy must scale with the severity of the threat.
The AI Control Roadmap recommends security measures based on the deployed model’s capability. The model’s detection evasion capability determines the Detection level (D1-D4), and the model’s attack execution capability determines the Prevention and Response level (R1-R3).
Analyzing a million agent trajectories to inform live monitoring
To build these systems, we need data. Our team has built an internal prototype of an asynchronous monitoring system that reviews coding agent trajectories, and escalates findings for human review. Each flagged event is classified against the threat taxonomy from our roadmap.
We’ve already analyzed a million coding agent tasks, giving us valuable insights to refine our safety protocols and allowing us to move beyond simple keyword filtering to identifying high-signal behavioral patterns. For example, this research has been instrumental in building a live monitor for the Gemini Spark agent, allowing us to respond in real-time to emerging issues, like unintentional data deletion.
It is important to note that our data shows the majority of flagged events do not stem from adversarial intent; instead, they often result from agent misinterpretation or overeagerness to achieve a user’s goal. Understanding these nuances is critical for refining our safety and security protocols.
Defending the agent ecosystem is a shared responsibility
The AI Control Roadmap is one vital element of our multilayered approach to agent security. By integrating these protocols with our broader infrastructure, we are scaling our internal security to safely manage our most advanced AI models.
Additionally, this multilayered approach to agent security should be a collaborative priority for industry, policymakers, and academia. By aligning the ecosystem around best practices and standards, we can empower cyber defenders and build societal resilience. That’s why today we are also publishing a technical framework for policymakers, 'Three Layers of Agent Security’ . The paper details how we need to improve security at the level of individual agents; in multi-agent systems; and to empower cyber defenders and build resilience across the broader ecosystem.
We intend to build on these frameworks to confidently deploy capable AI today while we continue to build a secure foundation for the future.
Read the full technical report here:
Mary Phuong, Erik Jenner, Laurent Simon, Lewis Ho, Rohin Shah, Sebastian Farquhar, and Scott Coull.
Four Flynn, Anca Dragan, Alan Cooney, Bilal Chughtai, Buck Shlegeris, Cody Wild, David Lindner, Julian Stastny, Kevin Klyman, Li Ding, Myriam Khan, Raluca Ada Popa, Roland Zimmermann, Ryan Greenblatt, Senthooran Rajamanoharan, Victoria Krakovna and Xerxes Dotiwalla.
I accept Google's Terms and Conditions and acknowledge that my information will be used in accordance with Google's Privacy Policy .
