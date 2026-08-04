---
source: "https://www.dynatrace.com/news/blog/how-observability-transforms-vibe-coding-into-ai-engineering/"
hn_url: "https://news.ycombinator.com/item?id=49168370"
title: "Observability Turns Vibe Coding into AI Engineering"
article_title: "How Observability Turns Vibe Coding into AI Engineering"
author: "speckx"
captured_at: "2026-08-04T13:51:51Z"
capture_tool: "hn-digest"
hn_id: 49168370
score: 2
comments: 0
posted_at: "2026-08-04T13:06:09Z"
tags:
  - hacker-news
  - translated
---

# Observability Turns Vibe Coding into AI Engineering

- HN: [49168370](https://news.ycombinator.com/item?id=49168370)
- Source: [www.dynatrace.com](https://www.dynatrace.com/news/blog/how-observability-transforms-vibe-coding-into-ai-engineering/)
- Score: 2
- Comments: 0
- Posted: 2026-08-04T13:06:09Z

## Translation

タイトル: 可観測性がバイブコーディングを AI エンジニアリングに変える
記事のタイトル: 可観測性がバイブコーディングを AI エンジニアリングに変える方法
説明: 信頼性の高い AI エンジニアリングのための可観測性、計画、スキル、実行時フィードバックにより、バイブコーディングを超えてください。

記事本文:
可観測性がバイブコーディングを AI エンジニアリングに変える方法
Observability がバイブコーディングを AI エンジニアリングにどのように変換するか
1. 典型的な AI チームメイトの最初のコミット
2. 信頼性の高い AI 生成コードの 4 層
4. 落とし穴: 運用環境における予測不可能なバグ
5. 新しいソフトウェア開発ループ
6. Dynatrace を使用して信頼性の高い AI エンジニアリング ワークフローを構築する
迅速なエンジニアリング、構造化された計画、スキル、可観測性を活用した実行時フィードバック ループを組み合わせることにより、AI によって生成されたコードとコーディング エージェントを予測不可能な貢献者から信頼できるエンジニアリング作業に変換する方法を発見します。
AI アシスタント、エージェント、実稼働テレメトリの間のループを閉じることで、組織は規律ある AI エンジニアリング ワークフローをより適切にサポートできるようになります。
AI は現在、どのような作業を行っているかに関係なく、ソフトウェアの構築方法の一部となっています。一部の開発者は、IDE 内のコパイロットとインライン提案に頼っています。 AI を活用した機能やアプリケーションを構築している人もいます。また、人間の入力をほとんど行わずにサービスを足場にしてプル リクエストをオープンするエージェントを調整している人もいます。自律性のレベルは異なりますが、コード生成はこれまでよりも高速になっており、そのコードが実際のシステムでどのように動作するかを理解することが新たなボトルネックとなっています。
CodeRabbit の調査によると、AI が生成したプル リクエストには、人間が作成したプル リクエストよりも高い割合で特定の欠陥が含まれていることが報告されています。
バグは避けられません。重要なのは、それらを素早く特定し、そこから学ぶことです。従来のソフトウェア ライフサイクル、つまり計画、コーディング、デプロイ、テスト、観察は、歴史的には一方通行であり、人間が「観察」段階から「計画」段階に洞察を移送していました。適切なツールがあれば、AI はそのループ自体を移動できます。サイクルは次のようになります。
計画→コード化→デプロイ→テスト→観察→そしてそれらの観察結果を内部にフィードバックします。

○次の計画。
このループを有効にするには、エージェントはツール (展開、テスト、可観測性用) にアクセスする必要があります。最新のコーディング エージェントはシェル環境で良好に動作します。これが、エージェントが外部システムにアクセスするための標準としてカスタム CLI が登場している理由の 1 つです。ライブ可観測性データの場合、エージェントは dtctl を使用します。これは、Dynatrace プラットフォームにアクセスするためのオープンソース CLI であり、kubectl からインスピレーションを得たものですが、エージェントは AWS CLI を使用してデプロイメントを実行できます。
可観測性は、コード生成からスケーラブルなソフトウェア エンジニアリングの実践に移行するための重要な要因となります。
典型的な AI チームメイトの最初のコミット
「ユーザーのバッチの注文を取得する関数を作成してください」という単純な質問を考えてみましょう。コードがインライン コパイロットからのものであっても自律コーディング エージェントからのものであっても、結果は多くの場合次のようになります。
async def get_user_orders(user_ids: list[str]) -> list[dict]:
"""ユーザーのバッチの注文を取得します。"""
httpx.AsyncClient() をクライアントとして非同期:
タスク = [
client.get(f"https://api.internal/orders?user_id={uid}")
user_ids の uid の場合
】
応答 = await asyncio.gather(*タスク)
注文 = []
応答内の応答の場合:
response.status_code == 200の場合:
order.extend(response.json())
返品注文
表面的には、これは良い仕事です。構文エラーや論理エラーはなく、メソッドの目的が文書化されています。しかし、実稼働環境でサービスを実行したことがある人には、ギャップがあることは明らかです。
これは、組織における「本番品質」が何を意味するのかについて知らされていない AI のデフォルトの動作です。数十のアーキテクチャ上の決定、HTTP クライアントの選択、エラー処理戦略、接続管理、再試行ロジック、タイムアウトがサイレントに行われ、内容や理由を通知することなくそれらが出荷されました。
信頼性の高い AI 生成コードの 4 層
プロンプトだけで成功することはほとんどありません

o 実稼働品質のコード。より信頼性の高いパターンは、AI が何かを書き込む前に一時停止させ、AI が何を構築しているのか、どのシステムに触れるのか、どこで失敗する可能性があるのか​​、実行後に結果をどのように観察するかなどのアプローチを説明することです。このステップは複雑である必要はありませんが、コード生成前に推論を強制する必要があります。
これは、軽量のスキルや再利用可能な指示が役立つ場所でもあります。すべてのプロンプトで同じ期待を再度表明する代わりに、チームはエージェントと副操縦士に、エラー処理、ロギング、トレース、レート制限、展開規則などに関するコンパクトな標準セットを提供できます。計画を立てることで、エージェントに考える機会が与えられます。スキルによって、その考え方が組織のソフトウェアの構築および運用方法と一致するようになります。
これら 2 つの実践を組み合わせることで、推測からデザインへと移行します。すべての実行時の問題を検出できるわけではありませんが、最初のバージョンがアーキテクチャの前提条件、運用標準、運用対応コードの定義を反映している可能性が大幅に向上します。
これを防ぐには、エージェントの上にガードレールを重ねる必要があります。
迅速なエンジニアリング: 明確かつ直接的にし、コンテキストを追加し、例を使用し、XML タグを使用し、出力形式を定義します。
計画: コードを記述する前に AI に「考えさせ」ます (オプションで専用の計画モードを使用)。
スキル: AI が知っていることと、AI がどのように機能するかを、指示、知識、スクリプトを通じて拡張します。
ツール: MCP サーバーまたは CLI (可観測性プラットフォーム、データベース、API) を介して AI に外部システムへのアクセスを提供します。
各レイヤーは、その下のレイヤーの値を合成します。迅速なエンジニアリングだけでは脆弱です。行動する前に計画と AI の理由を追加します。スキルを追加すると、自分の基準に基づいた推論が可能になります。ツールを追加すると、AI がその動作を現実に照らして検証できるようになります。これらのレイヤーは、ユーザーに合わせて拡張できます。

エージェントのフリートを運用するための 1 つの提案を受け入れます。
可観測性は、AI がデフォルトで何も設定しない領域であるため、スキルとしてエンコードするのに最も価値のある領域の 1 つです。優れた可観測性スキルは、以下に関する指針を提供します。
インストルメンテーションの規則: ログ形式と重大度レベル、何がどこに記録されるか、エラーがどのように記録されるか。これは、チームまたは組織の基準に基づいて独自に作成する必要があるスキルです。
OpenTelemetry パターン: 自動インスツルメンテーションと手動スパン、使用する各言語のセットアップ、クエリ時に自動的に関連付けられるように 3 つのシグナル (メトリクス、トレース、ログ) をリンクする方法。
計画を立てて適切なスキルをロードすると、同じ「注文のフェッチ」リクエストによって劇的に異なるコードが生成されます。現在、エージェントはセマフォを使用してダウンストリームのレート制限を尊重し、2 つのスパン層 (フェッチごとに 1 つ、バッチ用に 1 つ) でインストルメンテーション標準に従い、トレース相関を含む構造化ログを出力し、バッチ全体の接続プーリングに単一の非同期 HTTP クライアントを使用します。実際に運用できるサービスです。
落とし穴: 本番環境における予測不可能なバグ
私たちの関数ははるかに良く見えますが、それを出荷すると、可観測性データで次のことが確認され始めます。
ダウンストリーム サービスは 50 件のオーダーでページネーションされます。この関数は最初のページのみを返し、残りのページは黙って削除されます。テストやレビューはそれを捕らえませんでした。それをカバーできるスキルはありませんでした。
これは、コードベースの外、およびチームの文書化された知識の外に存在するバグのクラスです。これはシステムの実行時の動作にのみ存在します。しかし、AI がランタイム情報と照らし合わせてコードを検証できたらどうなるでしょうか?
新しいソフトウェア開発ループ
ここで新しいループが始まります: 計画 → コード → デプロイ → テスト → 観察 → 観察者にフィードを与える

次の計画に戻りましょう。
エージェントと副操縦士が実行環境でそのコードがどのように動作するかを確認できるようになると、上記のページネーションのバグが将来の反復の指針となるフィードバック信号となる可能性があります。 AI がループのどの程度を独自に実行するかは、ユーザー次第です。 AI 支援の開発者は、テレメトリを使用して次のプロンプトを手動で通知する場合があります。エージェント主導のアプローチでは、ループ全体が無人で実行されるように接続される可能性があります。原則はどちらの場合でも当てはまります。
実際には、このループを構築するということは、それぞれがサイクルの 1 つの段階に範囲を限定した、小規模な CLI、ツール、およびスキルのセットをエージェントに提供することを意味します。 Dynatrace 環境では、テレメトリ フローを迅速に取得する dtwiz と、コードの実行後に Dynatrace API を操作する dtctl を含めることができます。
コードを書く: Claude Code などのコーディング エージェントは、プロジェクトのスキルをロードして、計画に従って機能を実装します。
デプロイ:deploy.md スキルとdeploy.sh スクリプトは、選択したコードとしてのインフラストラクチャ ツールを介して、コードの変更をクラウド プロバイダーにプッシュします。新しいサービスまたは環境の場合、セットアップ ステップで dtwiz を使用してシステムを分析し、適切な可観測性パスを推奨し、最初に手動でコレクタ構成を構築しなくてもテレメトリ フローを取得できるようにすることができます。
テスト: 開発環境またはステージング環境で実行している場合は、合成ロードを実行するか、ロード テスト環境をトリガーします。
観察: dtctl.md スキルは、Dynatrace API を操作するための CLI である dtctl と組み合わせて、デプロイされたコードが出力したばかりのトレース、メトリクス、ログ、イベント、またはエンティティ データを取得します。
計画: Analysis.md スキルと planing.md スキルは、次の反復の前にテレメトリをスキャンして問題がないか改善を計画するようエージェントに指示します。
重要な詳細: 自動化されたワークフローの場合、このフィードバック パイプラインはスキップされるオプションのステップではなく、フックとして実行される必要があります。フックを使用すると、すべてのコード変更を強制的に通過させることができます

「完了」とみなされる前に、deploy-test-monitor-analyze を実行します。人間が常に情報を把握している場合でも、ゲートを習慣化することで、バイブ コーディング セッションがエンジニアリングに変わります。
Dynatrace で信頼性の高い AI エンジニアリング ワークフローを構築する
ループ内での可観測性を備えた AI 支援開発についてさらに詳しく調べることに興味がありますか?以下のリソースを確認してください。
Skills.sh でコミュニティ スキルを参照してインストールします。
Dynatrace AI スキル バンドルを追加します: npx skill add dynatrace/dynatrace-for-ai
エージェントと副操縦士をリアルタイムの可観測性データに接続するための Dynatrace MCP の詳細をご覧ください
実際の GenAI トレースをスコアリングし、評価結果を Dynatrace AI Observability に送り返すためのオープンソース CLI である dt-evals を使用して、LLM とエージェントの品質を評価します。
タグ:
AIエンジニアリング , AI支援開発 , クロードコード , コーディングエージェント , MCP , 可観測性 , オープンテレメトリ , スキル
フロリアン・メア
Florian は、Dynatrace のプロダクト アーキテクトであり、ハイパースケーラーの可観測性に重点を置いています。彼はソフトウェア エンジニアおよびアーキテクトとして 10 年以上の経験があり、最近では分散システム、クラウド アーキテクチャ、およびエージェント システムに取り組んでいます。エンジニアリングを超えて、フロリアンは情熱的な登山家であり、ヨーロッパ中の最も高い山のいくつかに登っています。
自分が何をやっているのか分かっているなら、それはバイブコーディングですか?
Dynatrace for AI: AI コーディング エージェントに Dynatrace の使用方法を教える
新しいディスカッションを開始するか、Q&A フォーラムで助けを求めてください。

## Original Extract

Move beyond vibe coding with observability, planning, skills, and runtime feedback for reliable AI engineering.

How Observability Turns Vibe Coding into AI Engineering
How Observability transforms vibe coding into AI engineering
1. A typical AI teammate’s first commit
2. The four layers of a reliable AI-generated code
4. The catch: Unpredictable bugs in production
5. The new software development loop
6. Building reliable AI engineering workflows with Dynatrace
Discover how to transform AI-generated code and Coding Agents from unpredictable contributors into reliable engineering work by combining prompt engineering, structured planning, skills, and a runtime feedback loop powered by observability.
By closing the loop between AI assistants, agents, and production telemetry, organizations can better support a disciplined AI engineering workflow.
AI is now part of how software gets built, no matter how you work. Some developers lean on copilots and inline suggestions inside the IDE. Some are building AI-powered features and applications. Others are orchestrating agents that scaffold services and open pull requests with little human input. The level of autonomy differs, but code generation gets faster than ever, and understanding how that code behaves in a real system is the new bottleneck.
According to research from CodeRabbit , AI-generated pull requests were reported to contain higher rates of certain defects than human-written pull requests.
Bugs are inevitable. The key is to identify them quickly and learn from them. The traditional software lifecycle, plan, code, deploy, test, observe, has historically been a one-way street, with humans transporting insights from the “observe” stage back to “plan.” With the right tools, AI can travel that loop itself. The cycle becomes:
Plan → Code → Deploy → Test → Observe → and feed those observations back into the next plan.
To enable this loop, Agents need access to tools (for deployment, testing and observability). Modern coding agents are performing well in shell environments, which is one of the reasons why custom CLIs are emerging as a standard for Agents to access external systems. For live observability data, the Agent will use dtctl , which is an open-source CLI for accessing the Dynatrace platform, inspired by kubectl while the Agent can run deployments using the AWS CLI.
Observability can be a key enabler for moving from code generation toward scalable software engineering practices.
A typical AI teammate’s first commit
Consider a simple ask: “Write a function to fetch orders for a batch of users.” Whether the code comes from an inline copilot or autonomous coding agent, the result often looks like this:
async def get_user_orders(user_ids: list[str]) -> list[dict]:
"""Fetch orders for a batch of users."""
async with httpx.AsyncClient() as client:
tasks = [
client.get(f"https://api.internal/orders?user_id={uid}")
for uid in user_ids
]
responses = await asyncio.gather(*tasks)
orders = []
for response in responses:
if response.status_code == 200:
orders.extend(response.json())
return orders
On the surface, this is good work. There are no syntax or logic errors, and the method’s purpose is documented. But the gaps are obvious to anyone who has run a service in production:
This is the default behavior of AI that hasn’t been told what “production quality” means at your organization. It made dozens of architectural decisions, HTTP client choice, error handling strategy, connection management, retry logic, timeouts – silently, and shipped them without telling you what or why.
The four layers of a reliable AI-generated code
Prompting alone rarely gets you to production-quality code. A more reliable pattern is to make AI pause and explain its approach before it writes anything: what it is building, which systems it will touch, where it might fail, and how the result should be observed once it runs. That step doesn’t need to be elaborate, but it should force the reasoning before code generation.
This is also where lightweight skills or reusable instructions help. Instead of restating the same expectations in every prompt, teams can give agents and copilots a compact set of standards for things like error handling, logging, tracing, rate limits, and deployment conventions. Planning gives the agent a chance to think; skills make that thinking consistent with how your organization builds and operates software.
Together, those two practices move away from guessing and toward design. They won’t catch every runtime issue, but they dramatically improve the odds that the first version reflects your architectural assumptions, operational standards, and definition of production-ready code.
Preventing this requires layering guardrails on top of the agent:
Prompt engineering: Be clear and direct, add context, use examples, use XML tags, and define output format.
Planning: Make the AI “think” before writing code, optionally in a dedicated plan mode.
Skills: Extend what AI knows and how it works through instructions, knowledge, and scripts.
Tools: Give AI the access to external systems through MCP servers or CLIs (observability platforms, databases, APIs).
Each layer compounds the value of the one below it. Prompt engineering alone is fragile. Add planning, and AI reasons before acting. Add skills, and reasoning is grounded in your standards. Add tools, and AI can verify its work against reality. These layers scale with you, from accepting a single suggestion to running a fleet of agents.
Observability is one of the highest-value domains to encode as a skill, because it’s the area where AI defaults to nothing at all. A good observability skill provides guidance on:
Instrumentation conventions: Logging format and severity levels, what gets logged where, how errors are recorded; This is a skill you should create on your own based on team or organizational standards.
OpenTelemetry patterns: Auto-instrumentation versus manual spans, setup for each language you use, and how to link the three signals (metrics, traces, and logs) so they’re automatically correlated at query time.
With planning and the right skills loaded, the same “fetch orders” request produces dramatically different code. Now the agent respects downstream rate limits with a semaphore, follows your instrumentation standard with two layers of spans (one per fetch, one for the batch), emits structured logs with trace correlation, and uses a single async HTTP client for connection pooling across the batch. That’s a service you can actually operate.
The catch: Unpredictable bugs in production
Our function looks much better, but when you ship it you start seeing this in your observability data:
The downstream service paginates at 50 orders. The function returns the first page only and silently drops the rest. No test or review caught it. No skill covered it.
This is the class of bug that lives outside your codebase and outside your team’s documented knowledge. It only exists in the runtime behavior of the system. But what if the AI could verify its code against runtime information?
The new software development loop
This is where the new loop comes into play: Plan → Code → Deploy → Test → Observe → and feed those observations back into the next plan.
When agents and copilots can see how its code behaves in a running environment, the pagination bug above may become a feedback signal that helps guide future iterations. How much of the loop the AI runs on it own is up to you. An AI-assisted developer might use telemetry to inform the next prompt by hand. An agent-led approach might wire the whole loop to run unattended. The principle holds either way.
In practice, building this loop means giving the agent a small set of CLIs, tools, and skills, each scoped to one stage of the cycle. In a Dynatrace environment, that can include dtwiz to get telemetry flowing quickly and dtctl to work with Dynatrace APIs once the code is running:
Write code: A coding agent, such as Claude Code, with your project skills loaded, implements features per the plan.
Deploy: A deploy.md skill and deploy.sh script push code changes to your cloud provider via the infrastructure-as-code tool of your choice. For a new service or environment, a setup step can use dtwiz to analyze the system, recommend the right observability path, and help get telemetry flowing without hand-building collector configuration first.
Test: Run synthetic load or trigger load test environments if running in a development or staging environment.
Observe: A dtctl.md skill paired with dtctl, the CLI for working with Dynatrace APIs, retrieves the traces, metrics, logs, events, or entity data the deployed code just emitted.
Plan: An analysis.md and planning.md skill instructs the agent to scan the telemetry for issues and plan improvements before the next iteration.
The critical detail: for automated workflows, this feedback pipeline should run as a hook, not an optional step that gets skipped. Hooks let you enforce that every code change passes through deploy-test-monitor-analyze before it’s considered “done.” Even when a humans stays in the loop, making the gate a habit is what turns a vibe coding session into engineering.
Building reliable AI engineering workflows with Dynatrace
Interested in going deeper on AI-assisted development with observability in the loop? Check out these resources:
Browse and install community skills at skills.sh
Add the Dynatrace AI skills bundle: npx skills add dynatrace/dynatrace-for-ai
Learn more about Dynatrace MCP for connecting agents and copilots to real-time observability data
Evaluate LLM and agent quality with dt-evals, an open source CLI for scoring real GenAI traces and sending evaluation results back to Dynatrace AI Observability.
Tags:
AI engineering , AI-assisted development , Claude Code , Coding Agents , MCP , observability , opentelemetry , skills
Florian Mair
Florian is a Product Architect at Dynatrace, focusing on Hyperscaler Observability. He has 10+ years of experience as software engineer and architect, most recently working on distributed systems, cloud architecture and Agentic Systems. Beyond engineering, Florian is a passionate mountaineer and has climbed some of the highest mountains across Europe.
Is it vibe coding if you know what you’re doing?
Dynatrace for AI: Teach your AI coding agent how to use Dynatrace
Start a new discussion or ask for help in our Q&A forum.
