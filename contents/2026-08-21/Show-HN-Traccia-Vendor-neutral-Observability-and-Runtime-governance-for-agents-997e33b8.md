---
source: "https://traccia.ai/"
hn_url: "https://news.ycombinator.com/item?id=49391995"
title: "Show HN: Traccia-Vendor neutral,Observability and Runtime governance for agents"
article_title: "Traccia | Agent Observability, Evaluation, Governance & Policy Enforcement"
image: "https://traccia.ai/opengraph.png"
author: "vijaypoudel"
captured_at: "2026-08-21T18:22:25Z"
capture_tool: "hn-digest"
hn_id: 49391995
score: 1
comments: 0
posted_at: "2026-08-21T18:20:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Traccia-Vendor neutral,Observability and Runtime governance for agents

- HN: [49391995](https://news.ycombinator.com/item?id=49391995)
- Source: [traccia.ai](https://traccia.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T18:20:16Z

## Translation

タイトル: HN を表示: Traccia ベンダー中立、エージェントの可観測性とランタイム ガバナンス
記事タイトル: トラッチャ |エージェントの可観測性、評価、ガバナンス、およびポリシーの施行
説明: ランタイム ポリシーの適用、プロンプト レジストリと評価、コスト帰属、および EU AI 法の証拠を備えた OpenTelemetry ネイティブのトレース。 LangChain、CrewAI、OpenAI Agents SDK などに対応する 1 つの SDK。
HN テキスト: AI アプリケーションはエージェントとなり、自律的な決定を下し始めています。エージェントや llms 呼び出しの動作をトレースしたり観察したりできるツールやプラットフォームが多数あります。彼らはその仕事においては優れていますが、AI エージェントの時代には追跡と観察可能性だけでは十分ではありません。エージェントのアクションを管理し、最終的に監査するための実行時ポリシーを観察、評価、作成するのに役立つソリューションが必要です。私たちはこの問題を解決するためにTracciaを開発しました。良い点は、これらすべてを数行のコードを記述するだけで実現できることです。 Traccia には、grafana、tempo、jaeger などの既存の可観測性ツールと連携できるオープンソースの SDK があります。可観測性以上のものが必要な場合に備えて、Traccia はエージェントを評価、制御、監査するためのプラットフォームを提供します。プラットフォームは使いやすいです。製品のドキュメントは非常に充実しています。また、クラウド ベンダーやフレームワークに依存しません。 Traccia は、バンガロールに本拠を置くインドの新興企業によって建設されています。 3 か月の無料トライアルを実施しているため、何の制約も課されることなくお試しいただけます。私たちは即興で改善することにオープンですので、コメントやフィードバックをお寄せください。

記事本文:
トラッチャ |エージェントの可観測性、評価、ガバナンスおよびポリシーの適用 Traccia について 価格設定 ガバナンス セキュリティ ブログ 投資 ドキュメント ログイン はじめる トグル メニュー OpenAI エージェント SDK の特集 AI エージェントの可観測性、評価、ガバナンスおよびポリシーの適用
無料で始めましょう ★ GitHub でスターを獲得 OpenTelemetry ネイティブ Apache 2.0 SOC 2 (進行中) LangChain CrewAI OpenAI エージェント SDK スクロール 01 / 06 現在、いくつの AI エージェントが実行されていますか?
どのエージェントが LLM コストを最も多く生み出しましたか?
エージェントが顧客データを漏洩したことがありますか?
制限された LLM モデルを使用したエージェントはいますか?
このプロンプトは本番環境よりも優れていますか?
1 回の init 呼び出し — すべてのエージェントをトレース、コスト、管理します
LangSmith、カスタム Grafana ダッシュボード、スプレッドシート間の切り替えを停止します。 Traccia は、フレームワークに関係なく、すべての AI エージェントに対して単一の画面を提供します。
リアルタイムのヘルス スコア エラー、遅延、スループットを即座に可視化して、エージェントのパフォーマンスを監視します。コストの帰属 LLM 支出を促進しているエージェントとタスクを正確に把握します。ポリシー違反アラート エージェントがガバナンス ルールまたは支出制限に違反した場合に即座に通知します。レジストリと評価バージョンのプロンプトを表示し、データセットとスコアラーで採点し、実験の証拠で昇格します。コンプライアンス証拠 EU AI 法および HIPAA コントロールの監査対応パックをガバナンス ハブからエクスポートします。 app.traccia.ai システム健全ダッシュボード
5% 傾向 エラー (24 時間) 142 12% 対前ポリシー違反 23 0% 横ばい傾向 トークン支出 $1,250.45 8% 対前エージェント
エージェント名 ステータス 所有者 環境 実行コスト Web 研究者 HEALTHY J Avam... プロダクション 248 $220.00 Zendesk Copilot DEGRADED J アバター... ステージング 27 $12.00 Stripe Risk Evaluator CRITICAL MA Kivin dev 18 $100.00 データサイエンティスト GPT HEALTHY J Avam... プロダクション 30 $13.00 Shopify Support Bot健康的

I アバター ステージング 20 $7.00 すべてのエージェントを表示 コスト スナップショット - 過去 24 時間 $1,250.45 上位支出者の Web 研究者 $1,250.45 Zendesk Copilot $12.00 トップ モデル GPT-4 $7.00 Claude 3.5 Sonnet $0.00 システム動作
AI エージェントを観察、評価、管理、制御するために必要なものすべて。
LangChain、CrewAI、OpenAI Agents SDK 全体のすべてのエージェントにとって唯一の信頼できる情報源。バージョン、健全性、環境、所有権を追跡します。コード変更は一切ありません。
すべての LLM コール、ツールの使用、エージェントの決定に対する OpenTelemetry ネイティブの完全な可視性。 LangChain、CrewAI、OpenAI Agents SDK、AutoGen、LlamaIndex をサポートします。
ポリシーの監視と施行
明示的なアノテーション、プロバイダーネイティブのシグナル (OpenAI、Anthropic、Google)、ヒューリスティックにわたる 3 層のガードレール検出。ハード ブロックは、事後ではなく、実行中にエージェントを停止します。
2,500 以上のモデルにわたる正確なコスト。サンプリングとは無関係に、スパンエンドでローカルに計算されます。トークンとコストの合計は、トレース サンプリングが 10% であっても 100% の精度を保ちます。
PII と機密データの検出
PII の漏洩 (患者名、認証情報、ユーザー データ) がログに到達する前に、エージェントの追跡にフラグを立てます。違反に注釈を付け、ハードブロックを強制し、自動的に準拠を維持します。
不変バージョンと保護された実稼働ラベルを持つ名前付きプロンプト。文言とモデルをレイテンシ、トークン、コストと並べて比較し、SDK からライブ バージョンをロードして、すべてのトレースが出荷されたものにリンクするようにします。
データセット、スコアラー、実験
テストケースを管理し、組み込みチェック、LLM-as-judge、またはカスタムスコアラーを使用して採点し、不変の実験レポートを保存します。ベースラインと候補を比較し、証拠を添付して宣伝します。
記事にマッピングされたすぐに使える EU AI 法の証拠 — Art. 12 ガバナンスイベントの範囲、アート。 14 ヒューマンレビュー、アート。 SDKの50discovery()。 HIPAA コンサルテーションに参加する

PHI 対応エージェントの在庫、セーフガードドラフト、ラベル付き輸出品のトロール。ワンクリックで証拠パックをエクスポートします。
美術。 12 ガバナンス イベントの期間と保存期間14 ヒューマン レビュー レビュー キューを統合したアート。 50 SDK 経由の開示 Disclosure() HIPAA PHI 在庫とラベル付き輸出を制御 Evidence Pack の輸出対応 OpenTelemetry-Native One SDK。すべてのフレームワーク。
LangChain、CrewAI、OpenAI Agents SDK、AutoGen、および LlamaIndex のドロップイン トレース。
あらゆる LLM と連携: OpenAI と Anthropic には自動トークンとコスト追跡が含まれます。 HTTP ベースのプロバイダーは OpenTelemetry 経由で追跡され、より詳細なカバレッジが継続的に追加されます。
監視するだけでなく強制する唯一のツール
事後に警告する人もいます。 Traccia のポリシー エンジンは、実行中にエージェントをハードブロックし、暴走コスト、無限ループ、PII 漏洩を本番稼働前に阻止します。
ランタイムオーバーヘッドゼロでのガードレール検出
パッシブ スパン プロセッサは、終了時にすべてのトレースを検査します。明示的、プロバイダーネイティブ、ヒューリスティックの 3 つの検出層で、エージェント コードを変更せず、遅延の追加もありません。
10% のサンプリングで嘘のない合計コスト
サンプリングにより微量量は削減されますが、コストはそれに比例しません。 Traccia は、すべての LLM コールの OTEL メトリクスを個別に発行します。サンプル レートに関係なく、トークンとコストの合計は 100% 正確です。
1 つのレジストリ、1 つの SDK にロックされない
LangSmith は、LangChain を使用する場合に最適に機能します。 Traccia は OpenTelemetry ネイティブであり、現在すべてのフレームワークで動作し、次の四半期に採用されるすべてのフレームワークで動作します。午後、 pip install tr​​accia で切り替えます。
企業が信頼するセキュリティ
大規模な可観測性、監査可能性、ガバナンスを必要とするチーム向けに構築されています。
SOC 2 Type II 認証取得中。企業調達向けに設計されたセキュリティ管理。
オプションのパターンベースのマスキング

init(redact_pii=True) 経由で鉱石をエクスポートします。 ML PII 検出が計画されています。
今日のきめ細かなロールベースのアクセス制御。エンタープライズ ロードマップ上の OIDC/SAML SSO。
今すぐTracciaを使って構築を始めましょう。
AI エージェントのトレースと監視を開始するパスを選択してください
トレースを独自の可観測性スタックにエクスポートします。 Jaeger、Grafana Tempo、Zipkin、SigNoz、および OpenTelemetry 互換のバックエンドで動作します。
エージェントファーストのガバナンス、コスト帰属、ポリシーのガードレール、チームのコラボレーションを備えた完全な Traccia ダッシュボードにアクセスします。

## Original Extract

OpenTelemetry-native tracing with runtime policy enforcement, prompt registry & evals, cost attribution, and EU AI Act evidence. One SDK for LangChain, CrewAI, OpenAI Agents SDK & more.

AI applications are becoming agents, which has started to take autonomous decisions. There are plenty of tools and platform available to trace, and observe what an agent or llms calls does. They are good in what they do, but tracing and observability isnt enough for AI agents era. We need a solution that can help you observe, evaluate, create run time policies to govern and finally audit the actions of the agent. We built Traccia to solve this problem. The good part, all of these can be achieved by just writing few lines of code. Traccia has an open-sourced sdk that can work with your existing observability tool like grafana, tempo, jaeger, etc. In case you need more than just observability, Traccia provides the platform to evaluate, control and audit the agents. The platform is easy to use. The product's documentation is quite extensive. It is also cloud vendor and framework agnostic. Traccia is being built by an Indian start up ,based out of Bengaluru. We are running a 3 months free trials so that you can explore without any strings attached. We are open to improvise and get better so please drop your comments and feedbacks.

Traccia | Agent Observability, Evaluation, Governance & Policy Enforcement About Traccia Pricing Governance Security Blog Invest Docs Login Get Started Toggle menu Featured in OpenAI Agents SDK AI Agent Observability, Evaluation, Governance & Policy Enforcement
Get Started For Free ★ Star on GitHub OpenTelemetry-native Apache 2.0 SOC 2 (in progress) LangChain CrewAI OpenAI Agents SDK Scroll 01 / 06 How many AI agents are running right now?
Which agent generated the most LLM cost?
Did any agent expose customer data?
Did any agent use a restricted LLM model?
Is this prompt better than production?
one init call — trace, cost, and govern every agent
Stop switching between LangSmith, custom Grafana dashboards, and spreadsheets. Traccia gives you a single pane of glass for all your AI agents, regardless of framework.
Real-time Health Scores Monitor agent performance with instant visibility into errors, latency, and throughput. Cost Attribution Know exactly which agents and tasks are driving your LLM spend. Policy Violation Alerts Instant notifications when agents breach governance rules or spending limits. Prompt Registry & Evals Version prompts, grade with datasets and scorers, and promote with experiment evidence. Compliance Evidence Export audit-ready packs for EU AI Act and HIPAA Controls from the Governance Hub. app.traccia.ai System Healthy Dashboard
5% trend Errors (24h) 142 12% vs prev Policy Violations 23 0% flat trend Token Spend $1,250.45 8% vs prev Agents
Agent Name Status Owner Env Executions Cost Web Researcher HEALTHY J Avam... production 248 $220.00 Zendesk Copilot DEGRADED J Avatar... staging 27 $12.00 Stripe Risk Evaluator CRITICAL MA Kivin dev 18 $100.00 Data Scientist GPT HEALTHY J Avam... production 30 $13.00 Shopify Support Bot HEALTHY I Avatar staging 20 $7.00 View all agents Cost Snapshot - Last 24 hours $1,250.45 Top Spenders Web Researcher $1,250.45 Zendesk Copilot $12.00 Top Models GPT-4 $7.00 Claude 3.5 Sonnet $0.00 System Behavior
Everything you need to observe, evaluate, govern and control your AI agents .
One source of truth for all agents across LangChain, CrewAI, and OpenAI Agents SDK. Track version, health, environment, and ownership — zero code changes.
Full OpenTelemetry-native visibility into every LLM call, tool use, and agent decision. Supports LangChain, CrewAI, OpenAI Agents SDK, AutoGen, and LlamaIndex.
Policy Monitoring & Enforcement
3-tier guardrail detection across explicit annotations, provider-native signals (OpenAI, Anthropic, Google), and heuristics. Hard blocks stop agents mid-execution — not after the fact.
Accurate cost across 2,500+ models. Computed locally at span-end, independent of sampling. Token and cost totals stay 100% accurate — even at 10% trace sampling.
PII & Sensitive Data Detection
Flags PII exposure — patient names, credentials, user data — in agent traces before they reach your logs. Annotate violations, enforce hard blocks, stay compliant automatically.
Named prompts with immutable versions and a protected production label. Compare wording and models side by side with latency, tokens, and cost, then load the live version from the SDK so every trace links back to what shipped.
Datasets, Scorers & Experiments
Curate test cases, grade with built-in checks, LLM-as-judge, or custom scorers, and save immutable experiment reports. Compare baseline vs candidate, then promote with evidence attached.
Article-mapped EU AI Act evidence out of the box — Art. 12 GovernanceEvent spans, Art. 14 human review, Art. 50 disclosure() in the SDK. Opt into HIPAA Controls for PHI-capable agent inventory, safeguard drafts, and labeled exports. Export evidence packs in one click.
Art. 12 Governance Event spans & retention Art. 14 Human Review Review queue integrated Art. 50 Disclosure disclosure() via SDK HIPAA Controls PHI inventory & labeled exports Evidence Pack export-ready OpenTelemetry-Native One SDK. Every Framework.
Drop-in tracing for LangChain, CrewAI, OpenAI Agents SDK, AutoGen, and LlamaIndex .
Works with any LLM: OpenAI & Anthropic include automatic token & cost tracking; any HTTP-based provider is traced via OpenTelemetry, with deeper coverage added continuously.
The only tool that enforces, not just observes
Others alert you after the fact. Traccia’s policy engine hard-blocks agents mid-execution — stopping runaway costs, infinite loops, and PII leaks before they hit production.
Guardrail detection with zero runtime overhead
A passive span processor inspects every trace as it ends. Three detection tiers — explicit, provider-native, heuristic — with no changes to your agent code and zero added latency.
Cost totals that don’t lie at 10% sampling
Sampling cuts trace volume but your costs don’t scale with it. Traccia emits OTEL metrics for every LLM call independently — token and cost totals stay 100% accurate regardless of sample rate.
One registry, not locked to one SDK
LangSmith works best if you use LangChain. Traccia is OpenTelemetry-native — works across every framework today and every one you adopt next quarter. Switch in one afternoon with pip install traccia .
Security that enterprises trust
Built for teams that need observability, auditability, and governance at scale.
SOC 2 Type II certification in progress. Security controls designed for enterprise procurement.
Optional pattern-based masking before export via init(redact_pii=True) . ML PII detection planned.
Granular role-based access control today. OIDC/SAML SSO on the enterprise roadmap.
Start building with Traccia today.
Choose your path to start tracing and monitoring your AI agents
Export traces to your own observability stack. Works with Jaeger, Grafana Tempo, Zipkin, SigNoz, and any OpenTelemetry-compatible backend.
Access the complete Traccia dashboard with agent-first governance, cost attribution, policy guardrails, and team collaboration.
