---
source: "https://fwdnow.io/en/blog/finops_ai_kosten_optimieren/"
hn_url: "https://news.ycombinator.com/item?id=48539815"
title: "We Get AI Costs Under Control"
article_title: "FinOps for AI: How We Get AI Costs Under Control - forwardnow GmbH"
author: "ayoisaiah"
captured_at: "2026-06-15T14:17:22Z"
capture_tool: "hn-digest"
hn_id: 48539815
score: 2
comments: 0
posted_at: "2026-06-15T11:36:02Z"
tags:
  - hacker-news
  - translated
---

# We Get AI Costs Under Control

- HN: [48539815](https://news.ycombinator.com/item?id=48539815)
- Source: [fwdnow.io](https://fwdnow.io/en/blog/finops_ai_kosten_optimieren/)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T11:36:02Z

## Translation

タイトル: AI コストを制御できるようになりました
記事タイトル: AI 向け FinOps: AI コストを管理する方法 - forwardnow GmbH

記事本文:
AI 向け FinOps: AI コストを管理する方法 - forwardnow GmbH
今すぐ進む前に
この Web サイトの機能を向上させるために、Formspree、Google Analytics、Hotjar などのサードパーティ Cookie とスクリプトを使用したいと考えています。
承認する
拒否する
詳細情報
サービス
AI のための FinOps: AI コストを管理する方法
トークンが新しいコスト単位になりつつある理由と、透明性、明確な制限、権限のあるチームがトークンをどのように管理するか
業界全体に波紋を呼んだ一例があります。ある AI コンサルタントは報道機関 Axios に対し、従業員の AI ライセンスに使用制限が設定されていなかったため、顧客の 1 つが 1 か月で約 5 億ドルを使い果たしたと語った。この事件は極端に聞こえるが、その規模では例外だ。根本的なパターンはそうではありません。小規模ではありますが、多くの組織が現在この問題を経験しています。システム内でアラームが鳴ることもなく、どのサービスやどのユーザーが増加の原因となったかを知る方法もないにもかかわらず、請求額が月に数百ユーロから数千ユーロに跳ね上がります。
クラウドの世界の FinOps に精通している人なら誰でも、すぐに問題に気づきます。コストを左右するのは個々の高価なトークンではなく、可視性の欠如と明確な境界の欠如です。 FinOps はまさにここで登場します。FinOps は、財務ガバナンスと運用の透明性を結びつけ、事業部門、IT、財務が共有データに基づいて意思決定を行えるようにします。同じロジックが AI にも当てはまります。コストの単位が変更されただけです。
朗報です。この透明性に関するオープンスタンダードが存在します。ログ、メトリクス、トレースが従来のモニタリングで確立されたのと同じように、OpenTelemetry GenAI セマンティック規約は、ベンダー中立的な環境で AI の消費を捉えるための統一された語彙を提供します。

はい、それを個々の情報源に帰属させます。この標準は、帰属からアーキテクチャ、継続的な制御に至るまで、以降のセクションを貫く共通のスレッドです。
つまり、AI 向け FinOps は 5 つのレバーに基づいています。1 つはトークン アトリビューションによる透明性であり、すべてのモデル呼び出しをユーザー、チーム、機能に割り当てることを意味します。 2 つ目は、すべての AI トラフィックの中央制御ポイントとしての AI プロキシです。第三に、制御不能なコストを防ぐ明確な制限とガードレール。 4 番目は、適切なモデルの選択、無駄のない呼び出し、キャッシュによる継続的な最適化です。 5 つ目は、コストを理解し、コストを負担する権限を与えられたチームです。次のセクションでは、これらのレバーを 1 つずつ説明します。
トークンはお金です: AI は従来のソフトウェアのように請求を行いません
クラシック ソフトウェアは、ほとんどの場合、ライセンスごとまたはシートごとに請求されます。コストは予測可能であり、2 つの請求期間の間で変化することはほとんどなく、調達は明確に定義された購入プロセスを通じて実行されます。 AIは異なる動作をします。ここでコストの単位はトークンであり、入力の長さ、応答の長さ、選択したモデルに応じて、呼び出しごとにコストが新たに発生します。
ここではクラウドとの類似点が決定的ですが、多くの人がそれを過小評価しています。購入の決定は組織内で民主化されています。クラウドがインフラストラクチャの調達を購買からエンジニアリングの手に移したのと同じように、AI は支出権限をさらに細かく分散します。中央機関がコストを決定することはなくなりました。代わりに、すべての開発者は、すべてのプロンプト、すべてのモデル選択、および起動するすべてのエージェントで実際の支出を引き起こします。したがって、コスト関連の決定が行われる頻度は、従来のソフトウェアよりも桁違いに高くなります。
この効果は、エージェント ワークフロー、つまり AI システムで特に顕著です。

複数のステップのタスクを自律的に処理するアイテム。このようなプロセスはループ内で動作し、コンテキストを繰り返し運び、中間ステップを生成するため、単一のモデル呼び出しの倍数を消費します。したがって、単一の不注意なループや無制限のバックグラウンド ジョブにより、短期間に多大なコストが発生する可能性があります。クラウドの構成ミスは数日にわたって影響を与えることがよくありますが、エージェントが制御不能になると数分以内にコストが増大する可能性があります。
これにより、中心的な質問が変わります。もはや AI を使用すべきかどうかではなく、AI の消費をどのように可視化し、個々のソースに帰属させ、必要に応じて制限できるかが重要です。
クラウドタグ付けからトークンタグ付けへ
意思決定者にとって最も重要な考え方は、AI のコスト管理はまったく新しい問題ではないということです。これは、より細かい粒度とはるかに高速な FinOps です。クラウドでのタグ付け戦略をすでに確立している人は、すでに決定的な考え方を持っています。新しい原価単位に転送するだけで済みます。
クラウド FinOps では、コスト センター、チーム、環境、プロジェクトなどのタグを介してすべてのリソースを帰属させます。 AI にもまったく同じ規律が必要ですが、ユーザー、チーム、機能、ワークフローなど、すべてのモデル呼び出しにタグが付けられるようになりました。通話の時点でこの属性がないと、プロバイダーの後で集計された請求を内訳することができなくなります。異常は説明できず、個々の機能のビジネス価値を計算することもできません。
ここでは 3 つの類似点が特に洞察力に富んでいます。まず、どちらの世界でも現実的な最大の問題はカバレッジです。どのクラウド プロジェクトでも、リソースの大部分は最初はタグ付けされておらず、最終的には未割り当てのバケット内に置かれます。 AI に関しても、同じ疑問が生じます。すべての通話には本当に ID が含まれているのでしょうか?第二に、どちらの場合も、レバーは t での執行にあります。

彼は情報源だ。クラウドでは、これはタグのないリソースがポリシーに違反していることを意味します。 AI では、帰属のない電話はそもそも通されないことを意味します。第三に、本当の戦略的利点は、同じ分類法が両方の世界で実行されるときに現れます。そうすることで初めて、インフラストラクチャと AI を合わせた機能の合計コストがいくらになるかを答えることが可能になります。これはまさに、FinOps オープン コストおよび使用法仕様 (略して FOCUS) が追求する共通点です。これにより、AI 消費データがますます均一な形式になり、Apptio などの既存のツールとの自然な接続が形成されます。
個々の部分を詳しく説明する前に、これらの構成要素がどのように組み合わされるかを次のアーキテクチャの概要で示します。
図: AI プロキシの構成図 — 左側にコンシューマー、中央に ID、使用状況測定、制限、ルーティング、テレメトリ エクスポート用のモジュールを備えたプロキシ、右側にオンプレミスを含むプロバイダー、そしてその下に分析システムとしての Dash0
可視性第一: AI テレメトリの標準
最適化する前に、測定する必要があります。冒頭で述べたように、OpenTelemetry GenAI Semantic Conventions により、このためのオープン標準が確立されました。具体的には、使用されるモデル、プロバイダー、操作の種類、入出力トークンの消費など、AI テレメトリの統一語彙を定義します。組織にとってのメリットは非常に大きいです。この標準に基づいて計測を行うと、単一のプロバイダーに縛られず、分析対象の場所に自由にデータを送信できるようになります。
コストを個々のユーザーに帰属させるには、独自のビジネス属性 (ユーザー ID、チーム、機能、またはコスト センター) をこれらの標準化フィールドに追加します。それはエグザックにあります

これらの属性を使用すると、後で集計が行われます。これは、クラシック ロギングのトレース ID または相関 ID と同じ考え方で、エントリをリクエストまたはビジネス プロセスに明確に割り当てることができます。ここでのみ、識別子はトラブルシューティングではなく、経済的な帰属として機能します。
実際には、これは、標準化されたフィールドと独自の属性属性を保持するスパンですべてのモデル呼び出しをラップするように見えます。
opentelemetry インポート トレースから
トレーサー = トレース。 get_tracer( "ai-プロキシ" )
トレーサー付き。 start_as_current_span( "chat" ) をスパンとして:
# 標準化された GenAI 属性 (OpenTelemetry セマンティック規則)
スパンもset_attribute( "gen_ai.operation.name" , "チャット" )
スパンもset_attribute( "gen_ai.provider.name" , "anthropic" )
スパンもset_attribute( "gen_ai.request.model" , "claude-sonnet-4-6" )
# 独自のアトリビューション属性 — あらゆるコストの内訳の基礎
スパンもset_attribute( "enduser.id" , "k.herings" )
スパンもset_attribute( "チーム名" , "顧客サポート" )
スパンもset_attribute( "feature.name" , "support-rag" )
スパンもset_attribute( "コストセンター" , "CC-4711" )
応答 = クライアント 。メッセージ。作成( ... )
# 通話後の消費を記録する
使用法 = 応答。使用法
スパンもset_attribute( "gen_ai.usage.input_tokens" , 使用法 . input_tokens)
スパンもset_attribute( "gen_ai.usage.output_tokens" , 使用法 .output_tokens)
gen_ai.* フィールドはオープンスタンダードに従っているため、互換性のあるどの分析システムでも同一です。 enduser.id や Team.name などのフィールドはビジネス レベルの追加であり、後で請求書をユーザー、チーム、または機能ごとに分類できます。重要なのは、この属性は通話時に設定されるということです。これは、後でプロバイダーの合計請求書から属性を再構築することができないためです。
データ広報からの重要な注意事項

保護の観点: 純粋なコスト監視では、プロンプトコンテンツを保存する必要はありません。メタデータは、トークン数、コスト、モデル、レイテンシ、および属性属性で十分です。特に、データ保護とデータの保存に関する高い要件がある組織にとって、この分離は決定的です。
中央制御ポイント: AI プロキシ
可視性と制御を統合するための最も実用的なアーキテクチャは、すべての AI トラフィックの中央通過点であり、多くの場合、AI ゲートウェイまたは AI プロキシと呼ばれます。各アプリケーションがプロバイダーと直接通信するのではなく、トラフィックはこの単一ポイントを介して実行されます。開発環境、チャット インターフェイス、エージェント パイプライン、社内サービスなどのアプリケーションとツールは、実際のプロバイダー キーを受け取りませんが、ユーザー、チーム、またはコスト センターに内部的にマッピングされた仮想キーを受け取ります。したがって、個々の開発者が特別なことをする必要がなく、すべての呼び出しが自動的に帰属されます。前に示した構成要素のビューは、この 1 つのコンポーネントで測定と施行がどのように連携するかを明らかにしています。
決定的な点は、測定と執行が同じ構成要素内にあるということです。プロキシは、消費量、モデル、コスト、レイテンシをキャプチャし、制限を適用し、単純なタスクをより小規模なモデルにルーティングし、オープン スタンダードに従ってテレメトリを分析システムにエクスポートします。これにより、透明性は月末の下流分析ではなく、すべての問い合わせの一部となりました。歓迎すべき副作用は、このパスにより、これまで制御されていなかった直接使用、いわゆるシャドウ AI も管理された環境に持ち込まれることです。
図: Dash0 の AI プロキシの健全性 — 異常検出の基礎となるサービスごとのスループット、レイテンシー、エラー率
可視化だけではコストの爆発を防ぐことはできません。それは th です

これは、まず第一に、賢明な境界を定義できるための前提条件です。実際には、明確な成熟パターンが現れています。チームは最初にリクエスト数の制限を導入し、最初のサプライズ請求の後にトークン消費の制限を追加し、2 番目以降は期間およびチームごとのハード予算制限を導入します。
これらの制限は、いくつかのレベルで連携して機能します。リクエスト数の制限によりインフラストラクチャが保護されます。トークンは計算量とコストに直接相関するため、トークンベースの制限により実際の消費量が決まります。予算制限により、最終的に、バッチ処理やエージェント ループによる予期せぬ負荷のスパイクが発生して、耐えられない請求が発生することを防ぎます。ゲートウェイでは、これを仮想キーごとに宣言できます。
# 仮想キーごとの制限 (例: カスタマー サポート チーム)
キー：サポートラグ
制限:
リクエストごとの分 : 120 # インフラストラクチャを保護します
tokens_per_minut : 200000 # 実際の消費量を調整します
予算：
期間：毎月
Soft_limit_eur : 5000 # ソフトしきい値 → アラート
hard_limit_eur : 6000 # ハードリミット → 通話は拒否されます
サーキットブレーカー :
cost_velocity_eur_per_min : 20 # 数分以内に暴走エージェントを停止します
特に効果的なのは、早期に警告するソフトしきい値とハード安全メカニズムの組み合わせです。

[切り捨てられた]

## Original Extract

FinOps for AI: How We Get AI Costs Under Control - forwardnow GmbH
Before you continue to forwardnow
We would like to use third-party cookies and scripts to improve the functionality of this website, including Formspree, Google Analytics, and Hotjar.
Approve
Deny
More info
Services
FinOps for AI: How We Get AI Costs Under Control
Why tokens are becoming the new unit of cost, and how transparency, clear limits, and empowered teams keep it under control
One example caused a stir across the industry: an AI consultant told the news outlet Axios that one of their clients burned through roughly half a billion dollars in a single month because no usage limits had been set on the employees’ AI licenses. The case sounds extreme, and at that scale it is the exception. The underlying pattern is not. On a smaller scale, many organizations are experiencing it right now: a bill jumps from a few hundred to several thousand euros a month, without any alarm going off in the system or any way to tell which service or which user caused the increase.
Anyone familiar with FinOps from the cloud world recognizes the problem immediately. It is not the individual expensive token that drives the cost, but the lack of visibility and the absence of clear boundaries. This is exactly where FinOps comes in: it connects financial governance with operational transparency, so that business units, IT, and finance can make decisions on a shared data basis. The same logic applies to AI. Only the unit of cost has changed.
The good news up front: there is now an open standard for this transparency. Just as logs, metrics, and traces became established in classic monitoring, the OpenTelemetry GenAI Semantic Conventions provide a uniform vocabulary to capture AI consumption in a vendor-neutral way and attribute it to individual sources. This standard is the common thread running through the sections that follow, from attribution through architecture to ongoing controlling.
In short, FinOps for AI rests on five levers: first, transparency through token attribution, meaning the assignment of every model call to user, team, and feature; second, an AI proxy as a central control point for all AI traffic; third, clear limits and guardrails that prevent uncontrolled costs; fourth, continuous optimization through the right model choice, lean calls, and caching; and fifth, empowered teams that understand and own their costs. The following sections walk through these levers one by one.
Tokens are money: AI does not bill like classic software
Classic software is mostly billed per license or per seat. Costs are predictable and rarely change between two billing periods, and procurement runs through a clearly defined purchasing process. AI behaves differently. Here the unit of cost is the token, and costs arise anew with every single call, depending on the length of the input, the length of the response, and the chosen model.
A parallel to the cloud is decisive here, and many underestimate it: the purchasing decision is democratized within the organization. Just as the cloud shifted the procurement of infrastructure out of purchasing and into the hands of engineering, AI distributes spending authority even more finely. No longer does a central body decide on costs; instead, every developer triggers real spending with every prompt, every model choice, and every agent they start. The frequency at which cost-relevant decisions are made is therefore orders of magnitude higher than with classic software.
The effect is especially pronounced with agentic workflows, that is, AI systems that handle multi-step tasks autonomously. Such processes consume a multiple of a single model call because they work in loops, repeatedly carry context along, and generate intermediate steps. A single careless loop or an unbounded background job can thus cause significant costs in a short time. Whereas a cloud misconfiguration often unfolds its effect over days, an agent running out of control can become costly within minutes.
This shifts the central question. It is no longer whether AI should be used, but how its consumption can be made visible, attributed to individual sources, and limited when necessary.
From cloud tagging to token tagging
The most important idea for decision-makers is this: AI cost control is not an entirely new problem. It is FinOps with a finer granularity and a much higher velocity. Anyone who has already established a tagging strategy in the cloud already holds the decisive mindset. It only needs to be transferred to the new unit of cost.
In cloud FinOps, we attribute every resource via tags such as cost center, team, environment, or project. AI needs exactly the same discipline, only now the tags hang on every model call: user, team, feature, or workflow. Without this attribution at the moment of the call, the provider’s later aggregated bill can no longer be broken down. Anomalies cannot be explained, and the business value of an individual feature cannot be calculated.
Three parallels are particularly insightful here. First, the biggest practical problem in both worlds is coverage. In every cloud project, significant portions of resources are initially untagged and end up in an unallocated bucket. With AI the same question arises: does every single call really carry an identity? Second, in both cases the lever lies in enforcement at the source. In the cloud this means an untagged resource violates a policy. With AI it means a call without attribution is not let through in the first place. Third, the real strategic advantage emerges when the same taxonomy runs across both worlds. Then, for the first time, it becomes possible to answer what a feature costs in total, infrastructure and AI together. This is exactly the common denominator pursued by the FinOps Open Cost and Usage Specification, FOCUS for short, which increasingly brings AI consumption data into a uniform format and thus forms the natural connection to existing tools such as Apptio.
How these building blocks fit together is shown by the following architecture overview, before we look at the individual parts in more detail.
Figure: Building-block view of the AI proxy — consumers on the left, the proxy in the center with its modules for identity, usage metering, limits, routing, and telemetry export, the providers including on-prem on the right, and Dash0 below as the analysis system
Visibility first: a standard for AI telemetry
Before you can optimize, you have to measure. As mentioned at the outset, an open standard has established itself for this with the OpenTelemetry GenAI Semantic Conventions. Concretely, they define a uniform vocabulary for AI telemetry, for example for the model used, the provider, the type of operation, and the consumption of input and output tokens. The advantage for organizations is considerable. Once you instrument against this standard, you are not tied to a single provider but can freely send the data wherever it is to be analyzed.
To attribute costs to individual users, your own business attributes are added to these standardized fields, that is, user identifier, team, feature, or cost center. It is on exactly these attributes that aggregation later happens. It is the same idea as the trace ID or correlation ID in classic logging, with which an entry can be unambiguously assigned to a request or business process. Only here the identifier does not serve troubleshooting but economic attribution.
In practice, this looks like wrapping every model call in a span that carries the standardized fields and your own attribution attributes:
from opentelemetry import trace
tracer = trace . get_tracer( "ai-proxy" )
with tracer . start_as_current_span( "chat" ) as span:
# Standardized GenAI attributes (OpenTelemetry Semantic Conventions)
span . set_attribute( "gen_ai.operation.name" , "chat" )
span . set_attribute( "gen_ai.provider.name" , "anthropic" )
span . set_attribute( "gen_ai.request.model" , "claude-sonnet-4-6" )
# Your own attribution attributes — the basis of every cost breakdown
span . set_attribute( "enduser.id" , "k.herings" )
span . set_attribute( "team.name" , "customer-support" )
span . set_attribute( "feature.name" , "support-rag" )
span . set_attribute( "cost_center" , "CC-4711" )
response = client . messages . create( ... )
# Record consumption after the call
usage = response . usage
span . set_attribute( "gen_ai.usage.input_tokens" , usage . input_tokens)
span . set_attribute( "gen_ai.usage.output_tokens" , usage . output_tokens)
The gen_ai.* fields follow the open standard and are therefore identical across any compatible analysis system. Fields such as enduser.id or team.name are the business-level addition along which the bill can later be broken down by user, team, or feature. What matters is that this attribution is set at the time of the call, because it cannot be reconstructed afterward from the provider’s aggregated bill.
An important note from a data protection perspective: pure cost monitoring does not require storing prompt content. The metadata is sufficient, that is, token counts, costs, model, latency, and the attribution attributes. Especially for organizations with high requirements for data protection and data residency, this separation is decisive.
A central control point: the AI proxy
The most practical architecture for bringing visibility and control together is a central passage point for all AI traffic, often referred to as an AI gateway or AI proxy. Instead of each application talking directly to the providers, the traffic runs through this single point. Applications and tools such as development environments, chat interfaces, agentic pipelines, or in-house services do not receive real provider keys, but virtual keys that are mapped internally to user, team, or cost center. Every call is thus automatically attributed, without the individual developer having to do anything extra. The building-block view shown earlier makes clear how measurement and enforcement come together in this one component.
The decisive point is that measurement and enforcement sit in the same building block. The proxy captures consumption, model, cost, and latency, enforces limits, can route simple tasks to smaller models, and exports the telemetry to an analysis system following the open standard. With this, transparency is no longer a downstream analysis at the end of the month but part of every single call. A welcome side effect is that this path also brings previously uncontrolled direct usage, so-called shadow AI, into a governed environment.
Figure: Health of the AI proxy in Dash0 — throughput, latency, and error rate per service as the basis for anomaly detection
Visibility alone does not prevent a cost explosion. It is the prerequisite for being able to define sensible boundaries in the first place. In practice, a clear maturity pattern has emerged. Teams first introduce limits on the number of requests, add limits on token consumption after the first surprise bill, and introduce a hard budget limit per period and team after the second.
These limits work together on several levels. Request-count limits protect the infrastructure. Token-based limits steer the actual consumption, since tokens correlate directly with compute effort and cost. Budget limits finally prevent unexpected load spikes from batch processing or agent loops from leading to untenable bills. At the gateway, this can be declared per virtual key:
# Limits per virtual key (example: Customer Support team)
key : support-rag
limits :
requests_per_minute : 120 # protects the infrastructure
tokens_per_minute : 200000 # steers the actual consumption
budget :
period : monthly
soft_limit_eur : 5000 # soft threshold → alert
hard_limit_eur : 6000 # hard limit → calls are rejected
circuit_breaker :
cost_velocity_eur_per_min : 20 # stops runaway agents within minutes
Particularly effective is the combination of soft thresholds that warn early and a hard safety mechanism

[truncated]
