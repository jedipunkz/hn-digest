---
source: "https://www.redhat.com/en/blog/operationalizing-agentic-ai-day-0-2-blueprint-enterprise-infrastructure"
hn_url: "https://news.ycombinator.com/item?id=49348421"
title: "Operationalizing agentic AI: The Day 0-2 blueprint for enterprise infrastructure"
article_title: "Operationalizing agentic AI: The Day 0-2 blueprint for enterprise infrastructure"
image: "https://www.redhat.com/themes/custom/rhdc/img/red-hat-social-share.jpg"
author: "mooreds"
captured_at: "2026-08-18T17:19:13Z"
capture_tool: "hn-digest"
hn_id: 49348421
score: 1
comments: 0
posted_at: "2026-08-18T16:43:21Z"
tags:
  - hacker-news
  - translated
---

# Operationalizing agentic AI: The Day 0-2 blueprint for enterprise infrastructure

- HN: [49348421](https://news.ycombinator.com/item?id=49348421)
- Source: [www.redhat.com](https://www.redhat.com/en/blog/operationalizing-agentic-ai-day-0-2-blueprint-enterprise-infrastructure)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T16:43:21Z

## Translation

タイトル: エージェント AI の運用化: エンタープライズ インフラストラクチャの Day 0-2 ブループリント
説明: Red Hat AI のプラットフォームが、Bring Your Own Agent (BYOA) をサポートし、コードを変更せずにエージェント フレームワークを運用できるようにする方法を学びます。

記事本文:
メニュー Red Hat を探索する
AI
概要
ハイブリッドクラウド
プラットフォームソリューション
人工知能 AI モデルとアプリを構築、展開、監視します。
Linux の標準化 オペレーティング環境全体で一貫性を確保します。
アプリケーション開発 アプリケーションの構築、展開、管理の方法を簡素化します。
自動化 自動化を拡張し、技術、チーム、環境を統合します。
仮想化 仮想化およびコンテナ化されたワークロードの運用を最新化します。
デジタル主権 重要なインフラを制御し、保護します。
セキュリティをコーディングし、セキュリティに重点を置いたソフトウェアを構築、展開、監視します。
エッジ コンピューティング エッジ テクノロジーを使用してワークロードをソースの近くに展開します。
Red Hat® Hybrid Cloud Console で、自分のペースでクラウド製品とソリューションを使用する方法を学びましょう。
Red Hat AI アイコン 人工知能、Red Hat Enterprise Linux AI、Red Hat OpenShift AI、RHEL AI、機械学習 38 38 2025-03-12T19:43:40.963Z image/svg+xml Red Hat AI アイコン 人工知能、Red Hat Enterprise Linux AI、Red Hat OpenShift AI、RHEL AI、機械学習 アイコンなし2025-03-12T19:39:59.817Z テクノロジー アイコン 標準 Red Hat AI ハイブリッド クラウド全体で AI ソリューションを開発およびデプロイします。
Red Hat Enterprise Linux アイコン RHEL、Linux プラットフォーム、CentOS 2024-03-01T15:26:42.958Z 保留中 TRA3b65dd25-844d-49bb-93c1-30f5b34684f1 アイコン 2024-03-01T15:26:42.958Z 真の保留中2024-03-21T00:40:29.326Z rhcc-audience:internal いいえ テクノロジー アイコン DER3b65dd25-844d-49bb-93c1-30f5b34684f1 標準 はい rhcc-product:red-hat-enterprise-linux テクノロジー アイコン image/svg+xml 2024-05-10T14:11:29.114Z Red Hat Enterprise Linux icon RHEL、Linux プラットフォーム、CentOS Activate Activate 2024-05-10T14:11:29.836Z workflow-process-service Activate workflow-process-service false 2024-05-10T14:11:29.836Z workflow-process-service 2024-05-10T14:11:29.836Z

テクノロジーアイコンを使用して、Red Hat 製品とコンポーネントを表します。境界図形からアイコンを削除しないでください。 Red Hat Enterprise Linux 柔軟なオペレーティング システムでハイブリッド クラウド イノベーションをサポートします。
Red Hat OpenShift アイコン クラウド、コンテナー、Kubernetes 2024-03-01T15:26:53.684Z 保留中 TRA9ec76aa9-ef09-4c49-8816-01dd13970ca7 アイコン 2024-03-01T15:26:53.684Z 真の保留中2024-03-21T00:39:44.126Z rhcc-audience:internal いいえ テクノロジーアイコン DER9ec76aa9-ef09-4c49-8816-01dd13970ca7 標準 はい rhcc-product:red-hat-openshift rhcc-product:red-hat-openshift-on-ibm-cloud rhcc-product:microsoft-azure-red-hat-openshift rhcc-product:red-hat-openshift-service-on-aws rhcc-product:red-hat-openshift-container-platform rhcc-product:red-hat-openshift-platform-plus テクノロジーアイコン画像/svg+xml 2024-05-10T14:18:23.703Z Red Hat OpenShift アイコン クラウド、コンテナー、Kubernetes Activate Activate 2024-05-10T14:18:25.221Z workflow-process-service Activate workflow-process-service false 2024-05-10T14:18:25.221Z workflow-process-service 2024-05-10T14:18:25.221Z テクノロジー アイコンを使用して、 Red Hat 製品とコンポーネントを表します。境界図形からアイコンを削除しないでください。 Red Hat OpenShift アプリを大規模に構築、モダナイズ、デプロイします。
Red Hat Ansible Automation Platform アイコン 管理、エッジ 2024-03-01T15:26:35.068Z 保留中 TRA759b57c4-760b-45a0-a939-821f47181964 アイコン 2024-03-01T15:26:35.068Z 真保留中2024-03-21T00:39:55.923Z rhcc-audience:internal いいえ テクノロジー アイコン DER759b57c4-760b-45a0-a939-821f47181964 標準 はい rhcc-product:red-hat-ansible-automation-platform テクノロジー アイコン image/svg+xml 2024-05-10T14:04:00.014Z Red Hat Ansible Automation Platform アイコン 管理、エッジ アクティブ化 アクティブ化 2024-05-10T14:04:01.784Z workflow-process-service アクティブ化 workflow-process-service false 2024-05-10T1

4:04:01.784Z workflow-process-service 2024-05-10T14:04:01.784Z テクノロジー アイコンを使用して Red Hat 製品とコンポーネントを表します。境界図形からアイコンを削除しないでください。 Red Hat Ansible Automation Platform 企業全体の自動化を実装します。
Red Hat OpenShift 仮想化エンジン
主要なクラウドプロバイダーとの統合
トレーニング
トレーニングと認定
Red Hat 認定システム管理者試験
Red Hat システム管理 I
Red Hat Learning サブスクリプションのトライアル (無料)
Red Hat 認定エンジニア試験
Red Hat 認定 OpenShift 管理者試験
クラウドネイティブのアプリケーションとサービスの構築、配信、管理に役立つリソースとツールを見つけてください。
信頼できるパートナーを活用したソリューションを構築する
Red Hat® Ecosystem Catalog の専門家とテクノロジーの協力コミュニティからソリューションを見つけてください。
閲覧中に好みのリソースをお勧めします。とりあえずこれらの提案を試してみてください。
ログイン Red Hat アカウントでさらに入手
一部のサービスでは購読が必要な場合があります。
口座番号: [[口座番号]]
Red Hat Ansible 自動化プラットフォーム
ポッドキャスト
Chris Wright と技術的に話す
その他のブログ
Red Hat 開発者ブログ
エージェント AI の運用化: エンタープライズ インフラストラクチャの Day 0-2 ブループリント
あなたのエージェントは機能します。推論し、ツールを呼び出し、デモで誰もが感銘を受ける回答を返します。ただし、そのノートブックと運用環境のデプロイメントの間には、モデルやフレームワークとは何の関係もないギャップがあります。一夜にして単一の AI エージェント導入に 3 つの障害が発生しました。43 枚の重複チケット、間違ったアカウントへの 4,000 ドルの請求、そして会社が遵守しなければならなかった 280 ドルの返品につながる幻覚的な返金ポリシーです。エージェントは LangChain 上で実行されました。ステージングでは完璧に機能しました。すべての失敗はインフラストラクチャの失敗であり、インテリジェンスの失敗ではありませんでした。
やったよ

CHECK チームは、ID を配線し、追跡を強化し、オーダーメイドのガバナンスを構築するなど、数か月をかけて手作業でそのギャップを埋めています。諦めて、実験から運用までのガイド付きパスを提供するハイパースケーラー プラットフォームに移行する人もいます。彼らは道を得るのです。彼らはまた、ロックインを取得します。 Bring Your Own Agent (BYOA) は Red Hat AI のアプローチです。プラットフォームは、コードを変更することなく、あらゆるエージェント フレームワークに本番インフラストラクチャを提供します。問題は、エージェントに運用インフラストラクチャが必要かどうかではなく、チームがそこに到達する方法です。
答えは 3 段階の旅です。
0 日目: リスクのない実験
1 日目: 自動的に導入されたインフラストラクチャを使用してデプロイする
2 日目: 可観測性とガバナンスを継続的に実行しながら運用する
この記事では、プラットフォームが機能することを証明するために、Red Hat 独自の本番エージェントを使用するチームの各フェーズを追っています。最も信頼できる証拠は、Red Hat がこのプラットフォームに独自の本番エージェントをデプロイしていることだからです。
0 日目: コミットする前に機能するかどうかを確認する
0 日目は、「これは私たちのユースケースで機能するでしょうか?」という質問に答えます。目標は、エンジニアリング リソースを運用化に投入する前に、それに答えることです。
Gen AI スタジオ (デプロイ前にモデル、エージェント、モデル コンテキスト プロトコル (MCP) サーバー、ガードレールを実験するための Red Hat AI 内の管理されたサンドボックス) は、Day 0 環境を提供します。チームは、vLLM (OpenAI 互換エンドポイントを提供する自己ホスト型推論エンジン) 上で実行されるさまざまなオープンウェイト モデルをエージェントに指示したり、MCP サーバーに接続してエージェントがアクセスできるツールを確認したり、NVIDIA NeMo ガードレールを適用して安全動作をテストしたりできます。これらはすべて、実験が実稼働データや実稼働インフラストラクチャに影響を及ぼさない管理された空間内で実行できます。私はそれを実験台だと考えています。本物の試薬、本物の装置があり、生産ラインを汚染するリスクはありません。
というチームにとっては、

どのフレームワークが必要かがわかっている場合は、スターター キットが 0 日目の開始点となります。 LangGraph、CrewAI、LlamaIndex、Langflow、Google ADK などの事前設定されたテンプレートは、すでに接続された Red Hat AI プラットフォーム統合とともに出荷されます。エージェントは、チームが統合コードを作成しなくても、推論のために vLLM に接続し、MCP ゲートウェイ (現在技術プレビュー段階にある Envoy ベースのプロキシで、プロンプト コンテンツではなくトークン クレームによってツール アクセスを管理します) に接続します。これらは「Hello World」デモではありません。これらには、認証、MCP 接続、トレースの初期化などのプラットフォーム統合パターンが含まれており、そうでなければチームが最初から構築するのに数週間を費やすことになります。
意思決定者にとって、リスクが最も低いのは 0 日目です。 gen AI Studio での実験では、運用に影響を与えずに管理されたリソースを使用します。 1 日目以前の投資は、1 チーム、1 サンドボックス、および管理されたコストです。開発者にとって、スターター キットのコードは実稼働環境で実行されているのと同じコードです。実験と展開の境界で何も捨てられることはありません。
1 日目: 書き換えを行わずにデプロイする
0 日目の実験のインフラストラクチャが引き継がれます。 1 日目は、エージェントがサンドボックスから本番環境に移行し、プラットフォームが最も困難な部分を引き継ぐ場所です。
開発者は、エージェントを Open Container Initiative (OCI) コンテナ (クラウドネイティブ インフラストラクチャ全体で使用される標準コンテナ形式、Docker イメージと同じ形式) としてパッケージ化し、AgentCard (エージェントの機能とツール要件の機械可読な記述) を作成して、それを Red Hat OpenShift AI にデプロイします。そこから、OpenShell (NVIDIA と協力して開発されたオープン ソース エージェント ランタイム) が運用化を自動的に処理し、ID、ポリシーの適用、セキュリティ監視を含む保護されたサンドボックスにエージェントをラップします。

開発者の元のエージェント コードを変更する必要はありません。
OpenShell はエージェントをラップし、開発者が作成したことのない運用インフラストラクチャを提供します。
Secure Production Identity Framework forEveryone (SPIFFE) / SPIFFE Runtime Engine (SPIRE) (各エージェントに偽造や共有が不可能な検証可能なデジタル ID を与えるための標準) による暗号化 ID。そのため、すべてのアクションは検証された ID に遡ります。 OpenShell スーパーバイザは、SPIFFE JSON Web トークン SPIFFE Verifiable Identity Document (JWT-SVID) トークンを発行し、エージェントの SPIFFE ID を使用した OAuth2 クライアント アサーション フローをサポートします。
MLflow (エージェントの実行を追跡するためのオープンソース プラットフォーム) および OpenTelemetry (OTEL) (可観測性データを収集するための標準) を介した分散トレース - したがって、すべての推論ステップ、ツール呼び出し、およびモデルの呼び出しが記録されます。 OpenShell サンドボックスは、構成可能なコレクター エンドポイントを介して OTEL トレースをエクスポートします。
MCP ゲートウェイを介したツール ガバナンス - そのため、ツールへのアクセスは、プロンプト コンテンツではなくトークン クレームによって制御されます。
Kata コンテナによるサンドボックス実行 (各エージェントに独自のカーネルを与えるハードウェア分離されたコンテナの実行) と OpenShell のアプリケーション レベルのサンドボックス (Landlock によるカーネル レベルのファイル システムの分離、seccomp による syscall フィルタリング、ネットワーク名前空間の分離、およびインプロセスの Open Policy Agent (OPA) ポリシーの強制) の 2 つの独立した防御層で、それぞれがもう一方に障害が発生しても保持されます。
開発者はこれについては何も書きません。 OpenShell は、開発者の問題ではなく、プラットフォームの問題としてこれを提供します。エージェントが未加工の資格情報を参照することはありません。OpenShell スーパーバイザ プロキシが API 呼び出しをインターセプトし、ネットワーク境界で資格情報を挿入します。
私がこの変化に何度も戻ってくるのは、これがこのシリーズの中で最も重要な理由だからです。ハイパースケーラー プラットフォームでは、開発者は

rite agents to the platform's security model. On Red Hat AI, developers write agents.プラットフォームはエージェントにセキュリティ モデルをもたらします。チームが 0 日目に書いたコードは、本番環境で実行されているコードと同じです。リファクタリングはありません。プラットフォーム固有のソフトウェア開発キット (SDK) のインポートはありません。
意思決定者にとって、OpenShell を通じて展開されるすべてのエージェントには、ID、トレース、ガバナンスが最小要件として含まれており、個々のチームが実装するかどうかのオプションのアドオンとしてではありません。セキュリティ体制は、フリート内のすべてのエージェントにわたって一貫しています。開発者にとって、チームが恐れているフレームワークの移行は決して起こりません。
Day 2: Operate what you deployed
Deployment isn't the finish line. 2 日目は最初の展開後に始まり、終わりません。ここは、大規模にエージェントを実行する実際の仕事が存在する場所であり、午前 6 時のインシデントが危機になる前に発見されたであろう場所です。
自問してください: 現在、展開されているエージェントに関するこれらの質問に答えることができますか?
過去 24 時間以内にどのエージェントがどのツールを呼び出しましたか?
各エージェントが推論に費やした金額は、増加傾向にあるのでしょうか?
前回のモデル更新以降、出力品質は変わりましたか?
エージェントが一晩で悪い結果を出した場合、何が起こったのかを正確に再現できますか

[切り捨てられた]

## Original Extract

Learn how Red Hat AI's platform supports bring your own agent (BYOA), providing operationalization for agent frameworks without code changes.

Menu Explore Red Hat
AI
Overview
Hybrid cloud
Platform solutions
Artificial intelligence Build, deploy, and monitor AI models and apps.
Linux standardization Get consistency across operating environments.
Application development Simplify the way you build, deploy, and manage apps.
Automation Scale automation and unite tech, teams, and environments.
Virtualization Modernize operations for virtualized and containerized workloads.
Digital sovereignty Control and protect critical infrastructure.
Security Code, build, deploy, and monitor security-focused software.
Edge computing Deploy workloads closer to the source with edge technology.
Learn how to use our cloud products and solutions at your own pace in the Red Hat® Hybrid Cloud Console.
Red Hat AI icon artificial intelligence, Red Hat Enterprise Linux AI, Red Hat OpenShift AI, RHEL AI, machine learning 38 38 2025-03-12T19:43:40.963Z image/svg+xml Red Hat AI icon artificial intelligence, Red Hat Enterprise Linux AI, Red Hat OpenShift AI, RHEL AI, machine learning Icon no 2025-03-12T19:39:59.817Z Technology icon Standard Red Hat AI Develop and deploy AI solutions across the hybrid cloud.
Red Hat Enterprise Linux icon RHEL, Linux platforms, CentOS 2024-03-01T15:26:42.958Z pending TRA3b65dd25-844d-49bb-93c1-30f5b34684f1 Icon 2024-03-01T15:26:42.958Z true pending 2024-03-21T00:40:29.326Z rhcc-audience:internal no Technology icon DER3b65dd25-844d-49bb-93c1-30f5b34684f1 Standard yes rhcc-product:red-hat-enterprise-linux Technology icon image/svg+xml 2024-05-10T14:11:29.114Z Red Hat Enterprise Linux icon RHEL, Linux platforms, CentOS Activate Activate 2024-05-10T14:11:29.836Z workflow-process-service Activate workflow-process-service false 2024-05-10T14:11:29.836Z workflow-process-service 2024-05-10T14:11:29.836Z Use technology icons to represent Red Hat products and components. Do not remove the icon from the bounding shape. Red Hat Enterprise Linux Support hybrid cloud innovation on a flexible operating system.
Red Hat OpenShift icon Cloud, Containers, Kubernetes 2024-03-01T15:26:53.684Z pending TRA9ec76aa9-ef09-4c49-8816-01dd13970ca7 Icon 2024-03-01T15:26:53.684Z true pending 2024-03-21T00:39:44.126Z rhcc-audience:internal no Technology icon DER9ec76aa9-ef09-4c49-8816-01dd13970ca7 Standard yes rhcc-product:red-hat-openshift rhcc-product:red-hat-openshift-on-ibm-cloud rhcc-product:microsoft-azure-red-hat-openshift rhcc-product:red-hat-openshift-service-on-aws rhcc-product:red-hat-openshift-container-platform rhcc-product:red-hat-openshift-platform-plus Technology icon image/svg+xml 2024-05-10T14:18:23.703Z Red Hat OpenShift icon Cloud, Containers, Kubernetes Activate Activate 2024-05-10T14:18:25.221Z workflow-process-service Activate workflow-process-service false 2024-05-10T14:18:25.221Z workflow-process-service 2024-05-10T14:18:25.221Z Use technology icons to represent Red Hat products and components. Do not remove the icon from the bounding shape. Red Hat OpenShift Build, modernize, and deploy apps at scale.
Red Hat Ansible Automation Platform icon Management, edge 2024-03-01T15:26:35.068Z pending TRA759b57c4-760b-45a0-a939-821f47181964 Icon 2024-03-01T15:26:35.068Z true pending 2024-03-21T00:39:55.923Z rhcc-audience:internal no Technology icon DER759b57c4-760b-45a0-a939-821f47181964 Standard yes rhcc-product:red-hat-ansible-automation-platform Technology icon image/svg+xml 2024-05-10T14:04:00.014Z Red Hat Ansible Automation Platform icon Management, edge Activate Activate 2024-05-10T14:04:01.784Z workflow-process-service Activate workflow-process-service false 2024-05-10T14:04:01.784Z workflow-process-service 2024-05-10T14:04:01.784Z Use technology icons to represent Red Hat products and components. Do not remove the icon from the bounding shape. Red Hat Ansible Automation Platform Implement enterprise-wide automation.
Red Hat OpenShift Virtualization Engine
Integrate with major cloud providers
Training
Training & certification
Red Hat Certified System Administrator exam
Red Hat System Administration I
Red Hat Learning Subscription trial (No cost)
Red Hat Certified Engineer exam
Red Hat Certified OpenShift Administrator exam
Discover resources and tools to help you build, deliver, and manage cloud-native applications and services.
Build solutions powered by trusted partners
Find solutions from our collaborative community of experts and technologies in the Red Hat® Ecosystem Catalog.
We'll recommend resources you may like as you browse. Try these suggestions for now.
Log in Get more with a Red Hat account
A subscription may be required for some services.
Account number: [[account_number]]
Red Hat Ansible Automation Platform
Podcasts
Technically Speaking with Chris Wright
More blogs
Red Hat Developer blog
Operationalizing agentic AI: The Day 0-2 blueprint for enterprise infrastructure
Your agent works. It reasons, calls tools, and returns answers in the demo that impress everyone. But between that notebook and a production deployment sits a gap having nothing to do with your model or your framework. Three failures hit a single AI agent deployment overnight—43 duplicate tickets, $4,000 charged to the wrong account, and a hallucinated refund policy leading to a $280 return the company had to honor. The agent ran on LangChain. It worked perfectly in staging. Every failure was an infrastructure failure, not an intelligence failure.
I've watched teams spend months closing that gap by hand—wiring identity, bolting on tracing, building bespoke governance. Some give up and move to a hyperscaler platform offering a guided path from experiment to production. They get the path. They also get the lock-in. Bring your own agent (BYOA) is Red Hat AI's approach: the platform provides production infrastructure for any agent framework without code changes. The question isn't whether your agent needs production infrastructure, it's how your team gets there.
The answer is a 3-phase journey.
Day 0: Experiment without risk
Day 1: Deploy with infrastructure injected automatically
Day 2: Operate with observability and governance running continuously
This article follows a team through each phase using Red Hat's own production agents as proof the platform works, because the most credible evidence is that Red Hat deploys its own production agents on this platform.
Day 0: Find out if it works before you commit
Day 0 answers the question: will this work for our use case? The goal is to answer it before committing engineering resources to operationalization.
Gen AI studio (a governed sandbox within Red Hat AI for experimenting with models, agents, Model Context Protocol (MCP) servers, and guardrails before deployment) provides the Day 0 environment. A team can point an agent at different open-weight models running on vLLM (a self-hosted inference engine providing OpenAI-compatible endpoints), connect MCP servers to see what tools the agent can reach, and apply NVIDIA NeMo Guardrails to test safety behavior—all in a governed space where experiments don't touch production data or production infrastructure. I think of it as the lab bench: real reagents, real equipment, no risk of contaminating the production line.
For teams that know which framework they want, starter kits provide the Day 0 starting point. Preconfigured templates for LangGraph, CrewAI, LlamaIndex, Langflow, Google ADK, and more ship with Red Hat AI platform integration already wired. The agent connects to vLLM for inference and to MCP Gateway (an Envoy-based proxy, currently in tech preview, governing tool access via token claims rather than prompt content) without the team writing integration code. These aren't "hello world" demos. They include the platform integration patterns—authentication, MCP connection, tracing initialization—teams would otherwise spend weeks building from scratch.
For the decision maker, Day 0 is where the risk is lowest. Experiments in gen AI studio use governed resources with no production impact. The investment before Day 1 is 1 team, 1 sandbox, and controlled cost. For the developer, the starter kit code is the same code running in production. Nothing gets thrown away at the boundary between experimentation and deployment.
Day 1: Deploy without rewriting
The infrastructure from Day 0 experimentation carries forward. Day 1 is where the agent moves from sandbox to production—and where the platform takes over the hardest part.
A developer packages their agent as an Open Container Initiative (OCI) container (the standard container format used across cloud-native infrastructure—the same format as Docker images), writes an AgentCard (a machine-readable description of the agent's capabilities and tool requirements), and deploys it to Red Hat OpenShift AI . From there, OpenShell (an open source agent runtime developed in collaboration with NVIDIA) will handle operationalization automatically—wrapping the agent in a protected sandbox including identity, policy enforcement, and security observability, all without requiring changes to the developer's original agent code.
OpenShell wraps the agent and provides the production infrastructure the developer never wrote:
Cryptographic identity via Secure Production Identity Framework for Everyone (SPIFFE) / SPIFFE Runtime Engine (SPIRE) (a standard for giving each agent a verifiable digital identity that can't be faked or shared)—so every action traces back to a verified identity. The OpenShell supervisor issues SPIFFE JSON Web Token SPIFFE Verifiable Identity Document (JWT-SVID) tokens and supports OAuth2 client assertion flows using the agent's SPIFFE identity.
Distributed tracing via MLflow (an open source platform for tracking agent execution) and OpenTelemetry (OTEL) (a standard for collecting observability data)—so every reasoning step, tool call, and model invocation is recorded. OpenShell sandboxes export OTEL traces via configurable collector endpoints.
Tool governance via MCP Gateway—so tool access is controlled by token claims, not prompt content.
Sandboxed execution via Kata Containers (hardware-isolated container execution giving each agent its own kernel) and OpenShell's application-level sandboxing (kernel-level filesystem isolation via Landlock, syscall filtering via seccomp, network namespace isolation, and in-process Open Policy Agent (OPA) policy enforcement)—2 independent layers of defense, each holding even if the other fails.
The developer will write none of this. OpenShell provides it as a platform concern, not a developer concern. The agent never sees raw credentials—the OpenShell supervisor proxy intercepts API calls and injects credentials at the network boundary.
I keep coming back to this shift because it's the most consequential in the series: on hyperscaler platforms, developers write agents to the platform's security model. On Red Hat AI, developers write agents. The platform brings the security model to the agent. The code a team wrote on Day 0 is the same code running in production. No refactoring. No platform-specific software development kit (SDK) imports.
For the decision maker, every agent deployed through OpenShell will have identity, tracing, and governance as minimum requirements—not as optional add-ons individual teams may or may not implement. The security posture will be consistent across every agent in the fleet. For the developer, the framework migration that teams dread never happens.
Day 2: Operate what you deployed
Deployment isn't the finish line. Day 2 begins after the first deployment and doesn't end. This is where the real work of running agents at scale lives—and where the 6 AM incident would have been caught before it became a crisis.
Ask yourself: Can you answer these questions about your deployed agents right now?
Which agents called which tools in the last 24 hours?
What did each agent spend on inference—and is that number trending up?
Has output quality changed since the last model update?
If an agent produced a bad result overnight, can you reconstruct exactly what happ

[truncated]
