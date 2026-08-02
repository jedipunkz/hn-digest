---
source: "https://buildtounderstand.org/explorations/what-is-the-smallest-ai-platform-that-could-possibly-work/"
hn_url: "https://news.ycombinator.com/item?id=49148769"
title: "What Is the Smallest (AI) Platform That Could Possibly Work?"
article_title: "What Is the Smallest (AI) Platform That Could Possibly Work? | Build to Understand"
author: "svenmalvik"
captured_at: "2026-08-02T22:45:59Z"
capture_tool: "hn-digest"
hn_id: 49148769
score: 1
comments: 0
posted_at: "2026-08-02T21:53:55Z"
tags:
  - hacker-news
  - translated
---

# What Is the Smallest (AI) Platform That Could Possibly Work?

- HN: [49148769](https://news.ycombinator.com/item?id=49148769)
- Source: [buildtounderstand.org](https://buildtounderstand.org/explorations/what-is-the-smallest-ai-platform-that-could-possibly-work/)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T21:53:55Z

## Translation

タイトル: 機能する可能性のある最小の (AI) プラットフォームは何ですか?
記事のタイトル: 機能する可能性のある最小の (AI) プラットフォームとは何ですか? |理解するために構築する
説明: 繰り返し作業が AI プラットフォームを正当化するのはどのような場合ですか?私は、所有権を高めたり、チームの方向転換の自由を制限したりすることなく、その作業を排除できる最小のプラットフォームを模索しています。

記事本文:
エンジニアリングの自由に関する探求、原則、システム。
システム
機能する可能性のある最小の (AI) プラットフォームは何ですか?
Vipps で最も一般的なリクエストの 1 つは、LLM へのアクセスです。エンジニアに API キーを与えるのは簡単です。即時アクセスの問題を解決します。ただし、どのチームがコストを作成したかを明らかにしたり、コンプライアンスがデータの行き先を理解したり、インシデント対応にシステムを観察する方法を提供したりすることはありません。
追加の作業は常に後から、場合によっては主な作業の数週間後に行われます。それでは、定期的な作業が複数のチームで使用されるプラットフォームを正当化するのはどのような場合でしょうか?そして、もしそれができるとしたら、機能する可能性のある最小のプラットフォームは何でしょうか?
今回はAIエージェントを使用しています。エージェントとは、LLM、スキル、および内部または外部ツールとの統合で構築されたソフトウェアを意味します。
プラットフォームはどのような問題を解決しますか?
組織は、エンジニアと非エンジニアが生産性を高めたり、新しい製品を作成したりするために便利なエージェントを構築することを望んでいます。より多くの人がそのようなエージェントをどのような形式であれ構築するにつれて、同じニーズや要求が現れます。開発者にはモデルへのアクセスと評価ツールが必要です。財務部門は、どのチームがどのくらいのコストを作成したかを知りたいと考えています。コンプライアンスでは、データ フローと監査証拠の制御が必要です。インシデント対応には監視できるシステムが必要ですが、リストは続きます。
これらの要求だけでは、AI プラットフォームを正当化することはできません。モデルへのアクセスを要求しているチームにアクセスの問題があります。法案について尋ねている財務部門は会計上の問題を抱えています。
実験に AI が使用されているという理由だけで、実験のコレクションがプラットフォームの問題になるわけではありません。
チームが同じアクセス制御、それを作成したチームにコストを結び付ける方法、評価、可観測性、または監査証拠を繰り返し必要とするときに、プラットフォームに問題があることがわかり始めました。で

その点、すべてのユースケースを個別に解決すると、より多くの作業、一貫性のない制御、または不必要なリスクが発生します。
企業は多くの実験を行っているが、プラットフォームに問題がない場合があります。別の企業では、本番環境に機密データを扱うエージェントが 2 つしかなく、すでに厳格な管理が必要な場合もあります。違いは、同じ問題が引き続き発生するかどうか、無視された場合に何が起こるか、複数のチームで一度解決することで、発生する作業よりも多くの作業を削減できるかどうかです。
その場合でも、解決策は既存のプラットフォームを拡張することかもしれません。すべてのチームが同じ AI 固有のコンポーネントを再構築しないと、既存のシステムが繰り返し発生する問題を解決できない場合にのみ、専用の AI プラットフォームを検討します。
繰り返しの作業は、常に同じように見えるとは限りません。ここで言及するのに十分なほど、私は 2 つの例を頻繁に見てきました。
アナリストとプロダクト マネージャーは、Claude Code または Codex を使用して何か (おそらくダッシュボード) を構築します。それは彼らのマシンで動作します。他の人が利用できるようにするということは、企業のランタイム環境にデプロイすることを意味します。
Vipps では、サービスは内部開発者ポータルを通じて利用できる契約に従います。これを使用するには依然としてエンジニアリングの経験が必要であるため、プラットフォーム チームは手動によるヘルプを繰り返し提供します。
シグナル: 新しいビルダーが作成されるたびに、同じセットアップと翻訳作業が戻ります。
エンジニアは AI を活用した Slack アプリを構築し、LLM への API キーを要求します。リクエスト自体にはほとんど時間がかかりません。
しかし、キーを配布しても、どのチームが各コストを作成したのか、データの管理、評価、有用な可観測性を提供したのかはわかりません。
シグナル: 簡単にアクセスできると、同じように答えのない運用上の疑問が生じます。
最初の例では、展開サポートを繰り返します。 2 つ目は、他の関係者が必要とする制御なしでアクセスを繰り返します。どちらもプラットフォームが解決できる可能性がある問題です。
私たちの AI プレイグラウンドは、最初の問題の一部を解決します。 A (非

)エンジニアは、社内開発者ポータルにアクセスし、エージェントに名前を付け、テスト環境にデプロイされた動作エージェントを含むリポジトリを受け取ります。 (非) エンジニアも、コーディング エージェントの開始プロンプトを受け取ります。価値は、テンプレートによって削除される繰り返し設定、テンプレートが使用する会社で許可されているテクノロジ、およびスキルとして追加されるベスト プラクティスです。
複数のチームが同じサービスを使用する必要があるのはどのような場合ですか?
チームはまだ学習中であるため、作業を繰り返す場合があります。集中化が早すぎると、実験が誰も望んでいない、必要としないサービスになってしまう可能性があります。共有サービスの構築を検討するのは、繰り返しの作業によって有用なワークフローが妨げられる場合、または目に見えるコスト、リスク、インシデント、または監査作業が発生する場合のみです。
複数のチーム向けの共有サービスを構築する前に、ワークフローが必要かどうか、または繰り返しの作業を削除できるかどうかも検討する必要があります。明らかなプラットフォームの問題は、所有権が不明瞭であること、トレーニングが欠如していること、または単純に削除する必要があるワークフローである可能性があります。
次のテストは私にとって重要です:
順序が重要です。まずはドキュメントから始めます。それが十分でない場合は、規則を確立してから、テンプレート、ライブラリ、または CLI を検討してください。マネージド サービスは、単純な介入では繰り返しの作業を取り除くことができない場合にのみ構築してください。
サービスは、複数のチームがそれを選択すると、プラットフォームのように機能し始めます。その理由は、単独で問題を解決するよりも簡単で安全だからです。
最小のプラットフォームとは何ですか?どこで終了する必要がありますか?
プラットフォームのサイズは、その機能、サービス、または構成行を数えることによって簡単に測定できます。その基準からすると、バージョン管理された構成ファイルは、機能やゲートウェイなどのサービスよりも小さいように見えます。
有用な最小の AI プラットフォームの考えられる定義は次のとおりです。
有用な最小の AI プラットフォームとは、総作業量が最小のシステムです

k は、定義されたグループが 1 つの貴重なワークフローを完了できるようにし、最低限必要な制御を満たし、繰り返し作業を削除するよりも変更または削除の方が安価なままです。
これにより、「最小の」4 つのテストが得られます。
それは測定された問題を解決し、対象のユーザーがジョブを完了できるようにするか?
データ、権限、可逆性、および起こり得る影響に必要な最小限の制御が提供されていますか?
中央およびローカルの統合、サポート、監査、インシデント、移行、および例外作業が軽減されますか?
許容可能なコストで、その契約、状態、証拠、ユーザーを変更または廃止することができますか?
その意味での「最小」とは、特定のワークフローの 4 つのテストすべてに合格する最も単純なソリューションです。
私が正当化できる最小のプラットフォーム
定義と上記の 4 つのテストに基づくと、考えられる開始点はエージェント コントラクトです。これは、エージェントを記述し、エージェントが使用するシステムを接続/記述するコードが格納されたバージョン管理された構成ファイルです。
私が思いつく最小のプラットフォームは、バージョン管理されたこのエージェント契約を中心に構築された狭いサービスです。契約自体はプラットフォーム単体ではなく、契約とサービスです。このサービスには、CI での検証、運用事実を所有するシステムへの接続、開発者ポータルでのビューなどが含まれます。
シンプルな内部の読み取り専用エージェントを使用して、初期スコープを定義します。その契約では、ワークフローに責任を持たせ、監視可能にするために必要な情報のみが宣言されています。
エージェントのアイデンティティ、目的、所有者、ライフサイクル
リスククラスと使用される可能性のあるデータの種類
承認されたモデルアクセスとアプリケーションランタイムへの参照
実行をコストおよびトレースに結び付ける識別子
インシデントの連絡先とエージェントを無効にする手順
これは、あらゆる種類の代理店の一般的な契約よりも小規模です

。ワークフローで必要な場合は、影響の大きいアクション、ツールの承認、受け入れられた例外、および保持された証拠を追加できます。初期スコープの一部である必要はありません。
エージェント：
氏名：和解説明員
所有者 : 支払い
目的 : 決済偏差の説明
ライフサイクル : 実験的
リスク : 内部読み取り専用
allowed_data : 内部
model_access : 承認済み EU プロバイダー
ランタイム : アプリケーションプラットフォーム
cost_id :支払い-決済
トレースID : 決済説明者
評価: evals/settlement.yaml
Incident_contact : オンコール支払い
無効化: runbook/disable-agent.md
この契約では、1 つの内部の読み取り専用ワークフローを説明責任と監視可能にするために必要なものだけを宣言します。
契約は、別のシステムがすでに知っている事実の 2 番目の情報源になってはなりません。安定した情報を宣言し、それが作成された運用上の事実を指摘する必要があります。
開発者ポータルは、これらのファクトをすべてリポジトリにコピーしなくても、1 つのビューに結合できます。契約には識別子と要件を含めることができます。元のシステムは引き続き現在の状態を担当します。
エンジニアは、同じリポジトリ内でエージェントとそのコントラクトを作成または変更します。
CI は、宣言された ID、所有権、リスク、データ クラス、運用上の参照、インシデントの連絡先、および無効化の指示を検証します。
配信、評価、可観測性、および請求システムは、運用上の事実を生成し、所有し続けます。
開発者ポータルは、宣言、現在の事実、およびソースへのリンクを 1 つのビューにまとめます。
エージェントは既存のアプリケーション ランタイム上で引き続き実行されます。 AI プラットフォームには新しいランタイムは導入されません。
プラットフォームは既存の情報を結び付けます。それを生成するシステムから所有権が奪われることはありません。
最小値はどのように変化する必要がありますか

リスク？
公開データを使用し、外部アクションを行わない実験またはテスト エージェントには、資金を移動したり実稼働インフラストラクチャを変更したりするエージェントと同じ制御は必要ありません。
初期スコープの内部の読み取り専用エージェントには、データ クラス、承認されたプロバイダーとリージョン、評価参照、各実行をそのトレースとコストに接続する方法、およびインシデントの連絡先が必要です。アクションを実行するエージェントには、ツールの権限、各実行の ID、承認、監査ログ、およびアクションを無効化または元に戻す方法も必要になる場合があります。起こり得る損害が大きくなり、アクションを元に戻すのが難しくなるほど、ワークフローにはより多くの制御が必要になります。
コントロールを定義して強制するのは誰ですか?
複数のチームが使用するコントロールには、次の 4 つの異なる責任があります。
プロバイダーからの直接アクセスが利用できない場合、ゲートウェイは許可されたモデルを強制できます。導入システムでは、エージェントをリリースする前に評価証拠が必要になる場合があります。ダウンストリーム サービスは、ID と権限を使用してビジネス アクションを承認する必要があります。
契約はこれらの決定を結びつけますが、執行に代わるものではありません。エージェントが過剰なアクセス権を持って資格情報を使用するのを止めることはできません。
プラットフォームはどこで終わるべきでしょうか?
この契約により、プラットフォームに可能性が与えられます。
AI プラットフォームは共通部分を定義し、組織がすでに使用しているシステムに接続できます。ワークフローに依存する決定は、ワークフローを理解しているチームが行う必要があります。
ワークフローでは LLM を使用するため、AI ゲートウェイや共有評価サービスは必要ありません。測定された繰り返し作業が削除される場合、または契約や既存のシステムでは提供できない必要な制御が提供される場合には、追加します。
チームが依然として同じサポートを必要とする場合、その情報が信頼できなくなる場合、またはすべてのシステムにカスタム アダプターが必要になる場合、契約はもはや十分ではありません。 T

次に追加されるのは、ジェネレーター、AI ゲートウェイ、認可サービス、評価サービス、または専用ランタイムです。
プラットフォームはどのようにして自由を維持できるのでしょうか?
便利なプラットフォームにより、チームはより簡単に作業できるようになります。その方法が唯一の現実的な方法または選択肢である場合、問題が発生します。
チームは、モデルへのアクセス、デプロイメント、可観測性、サポート、およびコストを特定する方法をプラットフォームに依存している場合があります。チームが他のものを選択したときにこれらを失った場合は、すでに構築されているものを再構築する必要があります。それは私には愚かに聞こえます。
自由には、前進できる有効な道が必要です。
チームは改善できるはずですが、プラットフォームから離れることもできます。標準プラットフォームのみが通常のサポートと監査の承認を受けている場合、別の実装は現実的な選択肢ではありません。
チームはどのようにしてプラットフォームを離れることができるのでしょうか?
辞めても仕事が生まれる。プロンプト、ツール スキーマ、評価、構成、状態、ログ、および保持されている証拠は、移動または再作成する必要がある場合があります。古いアクセスを取り消す必要があります。チームは移行中に両方のオプションを実行する必要がある場合があります。
プラットフォームは、この作業を可視化し、サポートし、可能にすることで自由を維持します。実用的な代替パスには次のものが必要です。
既存のアイデンティティ、展開、可観測性、インシデント対応、およびコスト システムへのアクセス
を示す方法

[切り捨てられた]

## Original Extract

When does repeated work justify an AI platform? I explore the smallest platform that could remove that work without creating more ownership or limiting teams’ freedom to change direction.

Explorations, principles, and systems about engineering freedom.
Systems
What Is the Smallest (AI) Platform That Could Possibly Work?
At Vipps, one of the most common requests I see is access to LLMs. Giving an engineer an API key is easy. It solves the immediate access problem. However, it doesn’t show finance which team created the cost, help compliance understand where the data goes, or give incident response a way to observe the system.
There is always additional work that often comes later, sometimes weeks after the main work. So, when does recurring work justify a platform used by several teams? And if it does, what is the smallest platform that could possibly work?
I am using AI agents as the current case. By agents, I mean software built with LLMs, skills, and integrations with internal or external tools.
What Problem Would a Platform Solve?
Organizations want engineers and non-engineers to build useful agents to become either more productive or to create new products. As more people build such agents in whatever form, the same needs and demands appear. Developers need model access and evaluation tools. Finance wants to know which team created what cost. Compliance wants control over data flows and audit evidence. Incident response wants a system it can observe, and the list goes on.
None of these requests alone justifies an AI platform. A team asking for model access has an access problem. The finance department asking about a bill has an accounting problem.
A collection of experiments doesn’t become a platform problem just because the experiments use AI.
I start seeing a platform problem when teams repeatedly need the same access controls, a way to connect costs to the teams that created them, evaluation, observability, or audit evidence. At that point, solving every use case separately creates more work, inconsistent controls, or unnecessary risk.
A company may have many experiments and no platform problem. Another company may have only two agents in production that handle sensitive data and already need strict controls. The difference is whether the same problem keeps appearing, what happens when it is ignored, and whether solving it once for several teams would remove more work than it creates.
Even then, the answer may be to extend an existing platform. I would consider a dedicated AI platform only when existing systems can’t solve the repeated problem without every team rebuilding the same AI-specific components.
Repeating work doesn’t always look the same. I have seen two examples often enough to be mentioned here.
Analysts and product managers use Claude Code or Codex to build something—perhaps a dashboard. It works on their machine. Making it available to others means deploying it to the company runtime environment.
At Vipps, services follow a contract available through the internal developer portal. Using it still requires engineering experience, so the platform team repeatedly provides manual help.
Signal: the same setup and translation work returns with every new builder.
Engineers build AI-powered Slack apps and ask for an API key to an LLM. The request itself takes little time.
But handing out a key doesn't show which team created each cost or provide data controls, evaluation, or useful observability.
Signal: easy access creates the same unanswered operational questions.
The first example repeats deployment support. The second repeats access without the controls other stakeholders need. Both are problems a platform might solve.
Our AI Playground removes part of the first problem. A (non-)engineer visits our internal developer portal, gives an agent a name, and receives a repository with a working agent deployed to the test environment. The (non-)engineer also receives a starter prompt for a coding agent. The value is the repeated setup that the template removes, the company’s allowed technologies it uses, and the best practices that are added as skills.
When Should Several Teams Use the Same Service?
Teams may repeat work because they are still learning. Centralizing it too early can turn experiments into a service nobody wants and needs. I would consider building a shared service only when repeated work blocks a useful workflow or creates visible cost, risk, incidents, or audit work.
Before I build a shared service for several teams, I should also ask whether the workflow is needed or whether the repeated work can be removed. The apparent platform problem may be unclear ownership, missing training, or a workflow that should simply be removed.
My next test is important for me:
The order matters. Start with documentation. If that isn’t enough, establish a convention, then consider a template, a library, or a CLI. Build a managed service only when the simpler interventions cannot remove the repeated work.
A service starts to act like a platform when several teams choose it because it is easier and safer than solving the problem on their own.
What Is the Smallest Platform, and Where Should It End?
It is simple to measure the size of a platform by counting its features, services, or lines of configuration. By that measure, a versioned configuration file looks smaller than a feature, or a service like a gateway.
Here is a possible definition for a smallest useful AI platform:
The smallest useful AI platform is the system that creates the least total work, lets a defined group complete one valuable workflow, meets the minimum required controls, and remains cheaper to change or remove than the repeated work it removes.
This gives “smallest” four tests:
Does it solve the measured problem and let the intended user complete the job?
Does it provide the minimum controls required by the data, authority, reversibility, and possible impact?
Does it reduce central and local integration, support, audit, incident, migration, and exception work?
Can its contracts, state, evidence, and users be changed or retired at an acceptable cost?
“Smallest” in that sense is the simplest solution that passes all four tests for a particular workflow.
The Smallest Platform I Can Justify
Based on the definition and the four tests from above, a possible starting point is an agent contract: a versioned configuration file stored with the code that describes the agent and connects/describes the systems it uses.
The smallest platform I can then can come up with is a narrow service built around a versioned this agent contract. The contract alone is not the platform alone but the contract and the service. The service includes validation in CI, connections to the systems that own operational facts, and a view in the developer portal, etc.
I use a simple, internal, and read-only agent to define the initial scope. Its contract declares only the information needed to make that workflow accountable and observable:
the agent’s identity, purpose, owner, and lifecycle
its risk class and the kind of data it may use
references to approved model access and the application runtime
identifiers that connect its runs to cost and traces
an incident contact and instructions for disabling the agent
This is smaller than a general contract for every kind of agent. High-impact actions, tool approvals, accepted exceptions, and retained evidence can be added when a workflow requires them. They don’t need to be part of the initial scope.
agent :
name : settlement-explainer
owner : payments
purpose : explain settlement deviations
lifecycle : experimental
risk : internal-read-only
allowed_data : internal
model_access : approved-eu-provider
runtime : application-platform
cost_id : payments-settlement
trace_id : settlement-explainer
evaluation : evals/settlement.yaml
incident_contact : payments-on-call
disable : runbooks/disable-agent.md
The contract declares only what is needed to make one internal, read-only workflow accountable and observable.
The contract shouldn’t become a second source for facts that another system already knows. It should declare stable information and point to operational facts where they are produced.
A developer portal can combine these facts into one view without copying all of them into the repository. The contract can contain identifiers and requirements. The original systems remain responsible for the current state.
An engineer creates or changes the agent and its contract in the same repository.
CI validates the declared identity, ownership, risk, data class, operational references, incident contact, and disable instructions.
Delivery, evaluation, observability, and billing systems continue to produce and own their operational facts.
The developer portal brings the declarations, current facts, and links to their sources into one view.
The agent continues to run on the existing application runtime; the AI platform does not introduce a new runtime.
The platform connects information that already exists. It doesn’t take ownership away from the systems that produce it.
How Should the Minimum Change With Risk?
An experiment or test agent using public data and taking no external action doesn’t need the same controls as an agent that moves money or changes production infrastructure.
The internal, read-only agent in the initial scope needs a data class, approved provider and region, an evaluation reference, a way to connect each run to its traces and cost, and an incident contact. An agent that takes action may also need tool permissions, identity for each run, approval, an audit log, and a way to disable or reverse the action. The greater the possible harm, and the harder the action is to reverse, the more controls the workflow needs.
Who Defines and Enforces a Control?
A control used by several teams has four different responsibilities:
A gateway can enforce allowed model if direct provider access is unavailable. A deployment system can require evaluation evidence before releasing an agent. A downstream service must authorize a business action using the identity and authority.
The contract connects these decisions but doesn’t replace enforcement. It can’t stop an agent from using a credential with too much access.
Where Should the Platform End?
The contract gives the platform a possible beginning.
The AI platform can define common parts and connect them to systems the organization already uses. Decisions that depend on the workflow should stay with the team that understands it.
An AI gateway or shared evaluation service doesn’t become necessary because the workflow uses an LLM. I would add one when it removes measured repeated work or provides a required control that the contract and existing systems can’t provide.
The contract is no longer enough when teams still need the same help, its information becomes unreliable, or every system needs a custom adapter. The next addition could then be a generator, AI gateway, authorization service, evaluation service, or dedicated runtime.
How Can the Platform Preserve Freedom?
A useful platform gives teams an easier way to work. It becomes a problem when that way is the only practical way or option.
A team may depend on the platform for model access, deployment, observability, support, and a way to identify its costs. If the team loses these when it chooses something else, it must rebuild what has been build already. That sounds stupid to me.
Freedom needs a usable path forward.
Teams should be able to improve but also leave the platform. If only the standard platform receives normal support and audit recognition, another implementation isn’t a real option.
How Can Teams Leave the Platform?
Leaving still creates work. Prompts, tool schemas, evaluations, configuration, state, logs, and retained evidence may need to move or be recreated. Old access needs to be revoked. The team may need to run both options during migration.
The platform preserves freedom by making this work visible, supported, and possible. A practical alternative path needs:
access to the existing identity, deployment, observability, incident response, and cost systems
a way to show t

[truncated]
