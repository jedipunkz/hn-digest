---
source: "https://commandline.microsoft.com/agent-control-specification-runtime-governance/"
hn_url: "https://news.ycombinator.com/item?id=48374497"
title: "Agent Control Specification: Portable Runtime Governance for AI Agents"
article_title: "Agent Control Specification: Portable runtime governance for AI Agents"
author: "thm"
captured_at: "2026-06-03T00:37:45Z"
capture_tool: "hn-digest"
hn_id: 48374497
score: 5
comments: 0
posted_at: "2026-06-02T18:49:00Z"
tags:
  - hacker-news
  - translated
---

# Agent Control Specification: Portable Runtime Governance for AI Agents

- HN: [48374497](https://news.ycombinator.com/item?id=48374497)
- Source: [commandline.microsoft.com](https://commandline.microsoft.com/agent-control-specification-runtime-governance/)
- Score: 5
- Comments: 0
- Posted: 2026-06-02T18:49:00Z

## Translation

タイトル: エージェント制御仕様: AI エージェントのためのポータブル ランタイム ガバナンス
記事のタイトル: エージェント制御仕様: AI エージェントのポータブル ランタイム ガバナンス
説明: ACS は、フレームワーク、ランタイム、またはポリシー エンジンに依存せず、エージェントのライフサイクル全体にわたってランタイム ガバナンスがどのように適用されるかを定義する、オープンでベンダー中立の標準です。

記事本文:
メインコンテンツにスキップ
メインコンテンツにスキップ
ソース
シグナルブログ
マイクロソフトの公式ブログ
マイクロソフトの問題点について
アジア
カナダ
ヨーロッパ、中東、アフリカ
ラテンアメリカ
私たちの規範
コネクシオネス
今日の新着情報
AI
イノベーション
デジタルトランスフォーメーション
持続可能性
セキュリティ
仕事と生活
ダイバーシティとインクルージョン
ロック解除済み
マイクロソフト 365
アズール
副操縦士
窓
表面
XBOX
お得情報
中小企業
サポート
Windows アプリ
展望
OneDrive
マイクロソフトチーム
OneNote
マイクロソフトエッジ
Skype から Teams への移行
コンピュータ
XBOX を購入する
付属品
VRと複合現実
認定再生品
下取りで現金化
XBOX ゲームパス アルティメット
PC ゲームパス
XBOX ゲーム
PCゲーム
マイクロソフト AI
マイクロソフトのセキュリティ
ダイナミクス 365
ビジネス向け Microsoft 365
マイクロソフト パワー プラットフォーム
Windows 365
中小企業
デジタル主権
アズール
マイクロソフト開発者
Microsoft Learn
AI マーケットプレイス アプリのサポート
マイクロソフト技術コミュニティ
マイクロソフト マーケットプレイス
ソフトウェア会社
ビジュアルスタジオ
マイクロソフト リワード
無料ダウンロードとセキュリティ
教育
ギフトカード
ライセンス
ロック解除されたストーリー
サイトマップを見る
ビルダーによる、ビルダーのための。
エージェント制御仕様の紹介: AI エージェントのポータブル ランタイム ガバナンス
ACS は、フレームワーク、ランタイム、またはポリシー エンジンに依存せず、エージェントのライフサイクル全体にわたってランタイム ガバナンスがどのように適用されるかを定義する、オープンでベンダー中立の標準です。
AI は新たな段階に入りつつあり、エージェントはそれを実際の経済効果に変えるメカニズムです。
私たちは、テキストやコードを生成するモデルを超えて、データを取得し、ツールを呼び出し、ワークフローを実行し、環境全体で意思決定を行うことによって機能するシステムに移行しています。 LangChain、AutoGen、CrewAI、Microsoft Agent スタックなどのフレームワークにより、エンドツーエンドで推論して操作できるエージェントの構築が簡単になりました。ソフトウェアが動作を開始するとき

人々に代わって、自動化の対象領域はあらゆるワークフロー、あらゆるシステム、あらゆる業界にわたって拡大します。
この変化により、新たな問題が生じます。エージェントが自律性を獲得するにつれて、問題はもはやエージェントに何ができるかだけではなく、何を許可されるべきなのか、誰がその境界を定義するのかということになります。最近の業界の取り組みでは、従来のシステムには存在しない障害モード、つまりツールの誤用や、エージェントのワークフロー全体で発生する複数ステップの障害にわたる意図しないアクションが浮き彫りになっています。同時に、規制当局の期待も加速しており、ハイリスク AI システムと説明責任に関する新たな要件がすでに課せられています。
現在、ガバナンスは主に開発層に存在しています。個々のチームは、プロンプト、アプリケーション コード、またはフレームワーク固有のフックにルールを埋め込みます。これらの制御は断片化されており、一貫性がなく、各エージェントの構築方法と密接に結びついています。セキュリティ チームやコンプライアンス チームがエージェント間でポリシーを定義、強制、監査できる標準的なメカニズムはありません。
既存のプレイブックが壊れる理由
従来のセキュリティ モデルは、固定されたアクターと固定されたスコープを前提としています。エージェントの場合、ある瞬間には安全かもしれない同じ資格情報が、次の瞬間には危険なものになります。会議の概要を投稿するのに適した Slack トークンは、エージェントが機密マークが付けられた文書を読んだり、グループに外部ユーザーが含まれたりした瞬間に危険になります。従来のアクセス制御にはこれに対する語彙がありません。「この資格情報でこのリソースへの呼び出しが許可されていますか?」に答えることしかできず、「このエージェントがこの会話で触れたすべてのことを考慮すると、この呼び出しはまだ安全ですか?」には答えられません。
お客様は、コードベース全体で分類子、検証、カスタム チェックをつなぎ合わせることで、これらのギャップを埋めていると語っていました。どのチームもこれのいくつかのバージョンを構築していました。私たちは構成されています

同じパターンが出現するのを見たことがありません。
多くの場合、システム プロンプトは防御の最前線となります。彼らはエージェントに何をすべきか、何をすべきではないかを伝えます。便利ではありますが、強制力はありません。プロンプト指示は、ユーザー入力、取得されたコンテンツ、ツールの結果、および攻撃者が制御する可能性のあるテキストと同じストリーム内に存在します。
エージェント内のカスタム ロジックにより、確定的チェックに対するより強力な保証を提供できます。しかし、これらのルールは通常、アプリケーション コードに埋め込まれています。これらは監査が難しく、再利用が難しく、チームがフレームワークを変更するときに移動するのが困難です。
入力および出力分類子は、ジェイルブレイク、プロンプト インジェクション、有害性、機密コンテンツなどのリスクを検出するのに役立ちます。しかし、分類子は多くの場合、孤立したテキストのみを認識します。どのツールが実行されようとしているのか、どのようなデータラベルが付けられているのか、エージェントがどのようなコンテキストを蓄積しているのかを自動的に知ることはできません。
フレームワーク固有のガードレール (OpenAI Agents SDK 入出力ガードレール、セマンティック カーネル フィルター、LangChain コールバック ハンドラー、Anthropic ツール使用コールバック) は適切な形状を取得しますが、途中で止まります。これらは各フレームワークが公開するネイティブ拡張ポイントです。つまり、組織が使用するすべてのフレームワークに対して同じポリシーを書き直す必要があり、セキュリティ チームにはコントロールを作成または監査するための単一の場所がありません。
汎用ポリシー エンジン (OPA/Rego、Cedar) は、構造化された入力に対して承認の質問に決定論的に答えます。彼らは成熟していて表現力が豊かで、広く理解されています。しかし、エージェント ループのモデルがなく、ライフサイクルのどの時点でエージェントに問い合わせるべきか、どの状態を収集すべきか、ホスト ランタイムでの判定を強制する方法についての概念もありません。
これらすべてのアプローチにおいて、ギャップは同じです。施行は分散しており、フレームワーク固有であり、エージェントのライフサイクルのより広範なコンテキストから切り離されています。実際の結果は次のようになります。

エージェントのセキュリティは、最終的にシステム プロンプト、フレームワーク コールバック、アプリケーション コード、コンテンツ分類子、ポリシー エンジンに分散しており、ポリシーをどのように評価して適用するかを記述した単一の契約はありません。
エージェント制御仕様 (ACS) の導入
ACS は、AI エージェントのランタイム ガバナンス層のオープン仕様およびリファレンス実装です。これは Microsoft のエージェント ガバナンス ツールキット (AGT) 内の新しいモジュールであり、開発者が AI エージェントを管理および統治する方法を拡張します。その核となるアーティファクトは、エージェント フレームワーク、ランタイム、またはルール自体を作成するポリシー エンジンから独立して、エージェント ライフサイクル全体にわたってポリシーがどこで、いつ、どのように評価され、適用されるかを定義する移植可能なマニフェストです。
ACS は、Rego などのポリシー言語をエージェント コンテキストで使用できるようにする欠落しているレイヤー、つまり標準化されたフック、入力、および施行コントラクトを提供します。ポリシー評価に関するオーケストレーションを管理します。
ライフサイクルインターセプト: エージェントループ内でチェックが行われる場所
正規の入力整形: どのような構造化コンテキストがポリシー エンジンに渡されるか
証拠の収集: 分類子、DLP サービス、裁判官、またはエンドポイントが事実をどのように提供するか
情報フローのチェック: ラベルとツールのクリアランスがどのように適用されるか
評決の正規化: ポリシーの結果がどのようにして標準的な ACS 決定となるか
最終的な適用: ホストによる許可、警告、拒否、またはエスカレーションの判定と編集効果がどのように適用されるか
フェールクローズ処理: ポリシー、証拠、または判定の処理が失敗した場合に何が起こるか
ACS: オープン標準の傍受ポリシー
ACS は、エージェントの実行時コンテキストに対してポリシーを評価できる 8 つのインターセプト ポイントを定義します。各ポイントは現在のスナップショットに対してポリシーを評価し、ポリシーは AC を許可、警告、拒否、またはエスカレーションできます。

ション。
Agent_startup : エージェントの実行を開始する前に構成と環境を評価します。
入力 : モデルが認識する前にユーザー入力を検査します。
pre_model_call : モデルに送信される完全なコンテキストを検査します。
post_model_call : ランタイムがモデルに作用する前にモデルの応答を検査します。
pre_tool_call : 実行前にツール名とパラメータを検査します。
post_tool_call : モデル コンテキストに再度入る前にツールの出力を検査します。
出力 : エージェントを離れる前に最終応答を検査します。
Agent_shutdown : ロギングと監査のためのセッション終了条件を評価します。
ACS への各呼び出しは独立しています。ホスト ランタイムは現在のスナップショットを渡し、ACS は正規入力を形成し、設定された証拠プロバイダーを実行し、ポリシー エンジンを呼び出して、正規化された判定を返します。以前のツール呼び出し、蓄積された機密性、承認状態、ユーザー履歴など、ポリシーがセッションについて知る必要があるものはすべて、ホストが渡すスナップショットに保存されます。
各介入ポイントで、ACS は現在のエージェント コンテキストを標準ポリシー入力に変換します。この入力は、エージェント ランタイムとポリシー エンジンの間のブリッジとなります。
ポリシー入力には、いくつかの重要な部分が含まれます。
Intervention_point は、エージェントのライフサイクルのどこでチェックが行われるかをポリシー エンジンに伝えます。たとえば、pre_tool_call は、ACS がツールの実行前にツール呼び出しを評価していることを意味します。
policy_target は、評価される特定のものです。ツール呼び出しでは、これはツールの引数になる場合があります。出力チェックでは、それが最終応答となる可能性があります。
スナップショットは、ホスト ランタイムによって提供されるより広範なコンテキストです。これには、アクター、役割、会話状態、以前のツール呼び出し、データ機密性、承認ステータス、またはポリシーに必要なその他のものが含まれます。
注釈には、ポリシーの実行前に収集された証拠 (ポリシーの実行結果など) が含まれています。

分類器、DLP システム、LLM ジャッジ、または外部サービス。
介入ポイントにツールが含まれる場合、ツールには、ツール名、クリアランス レベル、セキュリティ ラベルなどのツール メタデータが含まれます。
実用的な例: 2 つの SDK にわたる同じマニフェスト
外部受信者にメッセージを送信してはならない電子メール エージェントを考えてみましょう。 1 つのマニフェストは 1 つの Rego ポリシーを pre_tool_call 介入ポイントにバインドし、ポリシーが理由を示すツールを宣言します。
Rego ポリシーは、投影されたツールの引数を読み取り、外部受信者を拒否します。
Python ホストは、AgentControl.from_path を使用してマニフェストをロードし、ツール前のスナップショットを評価します。
どちらの SDK も同じネイティブ コアをロードし、同じ Rego バンドルを評価します。クロス SDK 適合フィクスチャは、.NET SDK と Rust SDK が同じスナップショットに対して同一の判定を返すことを主張するため、コントロールは、Python サービスからノード サイドカーに移動するとき、またはローカル スクリプトからホストされたランタイムに移動するときにエージェントに従います。
リポジトリは github.com/microsoft/agent-governance-toolkit にあります。
このドキュメントには、一般的なフレームワークと言語のクイックスタート ガイドが含まれています。このプロジェクトは MIT ライセンスを取得しており、オープンに開発されています。スペックこそが真実の源です。 SDK の動作から逸脱するものはバグです。問題、RFC、アダプターの貢献は歓迎されます。
エージェントのフレームワークが変わります。ポリシーエンジンは進化し​​ます。ガバナンスの要件は増加します。執行契約は毎回書き直す必要はありません。
Agent Framework Toolkitとの関係
ACS はコントロール層であり、エージェント フレームワークではありません。ループの調整、ツールの選択、メモリの管理は行いません。これらの責任はホストと、エージェントが構築されているフレームワーク (Microsoft Agent Framework など) に属し、そのオブジェクトは ACS SDK 適応を通じて直接保護できます。

えーっ。 ACS は、各フレームワークがすでに公開しているモデレーション ポイントに接続し、どのフレームワークも独自に提供しない上位の部分、つまり正規の入力形状、証拠パイプライン、正規化された評決、およびフェールクローズされた執行コントラクトを提供します。その結果、チームはポリシー サーフェスを書き換えることなくエージェント フレームワークを選択または変更できるようになり、セキュリティ チームはその下のランタイムに関係なく、コントロールを 1 か所で作成、バージョン管理、監査できるようになります。
エージェント ガバナンス ツールキットとの関係
Agent Governance Toolkit (AGT) は、運用エージェントのポリシー適用、ID、サンドボックス化、監査をバンドルする Microsoft 署名付きランタイムです。 AGT の次のバージョンでは、ポリシー言語として ACS が採用されるため、既存の AGT ユーザーは、AGT の ID、サンドボックス化、および監査の保証を維持しながら、8 つの介入ポイント、正規のポリシー入力、Rego ベースの意思決定、ACS が提供するフレームワーク アダプターを利用できるようになります。
>
「エージェント エコシステムには、ツール プロトコルやモデル インターフェイスのオープン スタンダードが必要であるのと同じように、ガードレールのオープン スタンダードが必要です。CrewAI は常に、誰でも拡張できる OSS コアである YAML で宣言されたオープン プリミティブ、エージェント、およびタスクに依存しており、ガードレールも同じパターンに従う必要があります。宣言

[切り捨てられた]

## Original Extract

ACS is an open, vendor-neutral standard that defines how runtime governance is applied across the agent lifecycle, independent of framework, runtime, or policy engine.

Skip to main content
Skip to main content
Source
Signal blog
Official Microsoft Blog
Microsoft On The Issues
Asia
Canada
Europe, Middle East and Africa
Latin America
The Code of Us
Conexiones
What's new today
AI
Innovation
Digital Transformation
Sustainability
Security
Work & Life
Diversity & Inclusion
Unlocked
Microsoft 365
Azure
Copilot
Windows
Surface
XBOX
Deals
Small Business
Support
Windows Apps
Outlook
OneDrive
Microsoft Teams
OneNote
Microsoft Edge
Moving from Skype to Teams
Computers
Shop XBOX
Accessories
VR & mixed reality
Certified Refurbished
Trade-in for cash
XBOX Game Pass Ultimate
PC Game Pass
XBOX games
PC games
Microsoft AI
Microsoft Security
Dynamics 365
Microsoft 365 for business
Microsoft Power Platform
Windows 365
Small Business
Digital Sovereignty
Azure
Microsoft Developer
Microsoft Learn
Support for AI marketplace apps
Microsoft Tech Community
Microsoft Marketplace
Software companies
Visual Studio
Microsoft Rewards
Free downloads & security
Education
Gift cards
Licensing
Unlocked stories
View Sitemap
By builders, for builders.
Introducing Agent Control Specification: Portable runtime governance for AI Agents
ACS is an open, vendor-neutral standard that defines how runtime governance is applied across the agent lifecycle, independent of framework, runtime, or policy engine.
AI is entering a new phase, and agents are the mechanism that will turn it into real economic impact.
We’re moving beyond models that generate text or code, and into systems that act by retrieving data, calling tools, executing workflows, and making decisions across environments. Frameworks like LangChain, AutoGen, CrewAI, and the Microsoft Agent stack have made it straightforward to build agents that can reason and operate end-to-end. When software starts acting on behalf of people, the surface area of automation expands across every workflow, every system, and every industry.
That shift introduces a new problem: As agents gain autonomy, the question is no longer just what they can do, but what they should be allowed to do and who defines those boundaries. Recent industry work highlights failure modes that don’t exist in traditional systems: tool misuse and unintended actions across multi-step failures that emerge across agent workflows. At the same time, regulatory expectations are accelerating, with new requirements around high-risk AI systems and accountability already coming into play.
Today, governance largely lives in the development layer. Individual teams embed rules in prompts, application code, or framework-specific hooks. These controls are fragmented, inconsistent, and tightly coupled to how each agent is built. There is no standard mechanism by which a security or compliance team can define, enforce, and audit policy across agents.
Why the existing playbook breaks
Traditional security models assume a fixed actor and a fixed scope. With agents, the same credential that may be safe in one moment becomes risky in the next. A Slack token that’s fine for posting a meeting summary becomes dangerous the moment the agent has read a document marked confidential or included external users in the group. Traditional access control has no vocabulary for this: It can only answer “is this credential allowed to call this resource?”—not “given everything this agent has touched in this conversation, is this call still safe?”
Customers told us they were patching these gaps by stitching together classifiers, validations, and custom checks throughout their codebases. Every team had built some versions of this. We consistently saw the same patterns emerge.
System prompts are often the first line of defense. They tell the agent what it should or should not do. They’re useful, but they aren’t enforceable. A prompt instruction lives in the same stream as user input, retrieved content, tool results, and potentially attacker-controlled text.
Custom logic inside the agent can provide stronger guarantees for deterministic checks. But those rules are usually buried in application code. They’re hard to audit, hard to reuse, and hard to move when the team changes frameworks.
Input and output classifiers help detect risks like jailbreaks, prompt injection, toxicity, or sensitive content. But classifiers often only see isolated text. They don’t automatically know which tool is about to run, what data labels are attached, or what context the agent has accumulated.
Framework-specific guardrails (OpenAI Agents SDK input/output guardrails, Semantic Kernel filters, LangChain callback handlers, Anthropic tool-use callbacks) get the shape right but stop short. They’re the native extension points each framework exposes, which means the same policy has to be rewritten for every framework an organization uses, and a security team has no single place to author or audit controls.
General-purpose policy engines (OPA/Rego, Cedar) answer authorization questions deterministically on structured inputs. They’re mature, expressive, and widely understood. But they have no model of an agent loop, no notion of when in a lifecycle to consult them, what state to collect, or how to enforce the verdict in the host runtime.
Across all these approaches, the gap is the same: Enforcement is scattered, framework-specific, and disconnected from the broader context of the agent’s lifecycle. The result, in practice, is that agent security ends up scattered across system prompts, framework callbacks, application code, content classifiers, and policy engines, with no single contract that describes how policies should be evaluated and enforced.
Introducing Agent Control Specification (ACS)
ACS is an open specification and reference implementation for the runtime governance layer of AI agents. It’s a new module within Microsoft’s Agent Governance Toolkit (AGT), extending how developers manage and govern AI agents. Its core artifact is a portable manifest that defines where, when, and how policies are evaluated and enforced across the full agent lifecycle, independent of the agent framework, the runtime, or the policy engine that authors the rules themselves.
ACS provides the missing layer that makes policy languages like Rego usable in the agent context: the standardized hooks, inputs, and enforcement contract. It owns the orchestration around policy evaluation:
Lifecycle interception: where checks happen in the agent loop
Canonical input shaping: what structured context is passed to the policy engine
Evidence collection: how classifiers, DLP services, judges, or endpoints contribute facts
Information flow checks: how labels and tool clearances are enforced
Verdict normalization: how policy results become standard ACS decisions
Final enforcement: how allow, warn, deny, or escalate verdicts, plus redaction effects, are applied by the host
Fail-closed handling: what happens when policy, evidence, or verdict processing fails
ACS: Open standard interception policies
ACS defines eight interception points where policies can be evaluated against the agent’s runtime context. Each point evaluates a policy against the current snapshot, and the policy can allow, warn, deny, or escalate the action.
agent_startup : evaluate configuration and environment before the agent begins running
Input : inspect user input before the model sees it
pre_model_call : inspect the full context being sent to the model
post_model_call : inspect the model’s response before the runtime acts on it
pre_tool_call : inspect tool name and parameters before execution
post_tool_call : inspect tool output before it re-enters the model context
output : inspect the final response before it leaves the agent
agent_shutdown : evaluate end-of-session conditions for logging and audit
Each call to ACS stands on its own: The host runtime passes the current snapshot, and ACS shapes the canonical input, runs configured evidence providers, invokes the policy engine, and returns a normalized verdict. Anything the policy needs to know about the session, including prior tool calls, accumulated sensitivity, approval state, user history, lives in the snapshot the host passes in.
At each intervention point, ACS turns the current agent context into a standard policy input. This input is the bridge between the agent runtime and the policy engine.
A policy input includes a few key pieces:
intervention_point tells the policy engine where in the agent lifecycle the check is happening. For example, pre_tool_call means ACS is evaluating a tool call before the tool runs.
policy_target is the specific thing being evaluated. In a tool call, this might be the tool arguments. In an output check, it might be the final response.
snapshot is the broader context provided by the host runtime. This can include the actor, roles, conversation state, prior tool calls, data sensitivity, approval status, or anything else the policy may need.
annotations contains evidence collected before the policy runs, such as results from classifiers, DLP systems, LLM judges, or external services.
tool includes tool metadata, like the tool name, clearance level, and security labels, when the intervention point involves a tool.
A worked example: The same manifest across two SDKs
Consider an email agent that must not send messages to external recipients. A single manifest binds one Rego policy to the pre_tool_call intervention point and declares the tool the policy reasons about:
The Rego policy reads the projected tool arguments and denies external recipients:
A Python host loads the manifest with AgentControl.from_path and evaluates the pre-tool snapshot:
Both SDKs load the same native core and evaluate the same Rego bundle. Cross-SDK conformance fixtures assert that the .NET and Rust SDKs return identical verdicts for the same snapshots, so the controls follow the agent when it moves from a Python service to a Node sidecar, or from a local script to a hosted runtime.
The repository is at github.com/microsoft/agent-governance-toolkit .
The documentation includes a quickstart guide for common frameworks and languages. The project is MIT-licensed and developed in the open. The spec is the source of truth. SDK behavior that diverges from it is a bug. Issues, RFCs, and adapter contributions are welcome.
Agent frameworks will change. Policy engines will evolve. Governance requirements will increase. The enforcement contract shouldn’t have to be rewritten each time.
Relationship with Agent Framework Toolkit
ACS is a controls layer, not an agent framework. It does not orchestrate the loop, choose tools, or manage memory. Those responsibilities belong to the host and to the framework the agent is built on, including the Microsoft Agent Framework , whose objects can be guarded directly through the ACS SDK adapters. ACS plugs into the moderation points each framework already exposes and supplies the pieces above them that no framework provides on its own: the canonical input shape, the evidence pipeline, the normalized verdict, and the fail-closed enforcement contract. The effect is that a team can pick or change agent frameworks without rewriting its policy surface, and a security team gets one place to author, version, and audit controls regardless of the runtime underneath.
Relationship with Agent Governance Toolkit
Agent Governance Toolkit (AGT) is the Microsoft-signed runtime that bundles policy enforcement, identity, sandboxing, and audit for production agents. The next version of AGT adopts ACS as its policy language, so existing AGT users gain the eight intervention points, the canonical policy input, Rego-based decisioning, and the framework adapters that ACS provides, while keeping AGT’s identity, sandboxing, and audit guarantees.
>
“The agent ecosystem needs an open standard for guardrails the same way it needs open standards for tool protocols and model interfaces. CrewAI has always leaned on open primitives, agents, and tasks declared in YAML, an OSS core anyone can extend, and guardrails should follow the same pattern: declarat

[truncated]
