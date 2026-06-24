---
source: "https://buildaharness.com/"
hn_url: "https://news.ycombinator.com/item?id=48660207"
title: "Show HN: Build A Harness – open-source, modular harness layer for AI agents"
article_title: "Build A Harness — Visual Canvas for AI Agent Harnesses"
author: "philiparxist"
captured_at: "2026-06-24T14:56:58Z"
capture_tool: "hn-digest"
hn_id: 48660207
score: 1
comments: 0
posted_at: "2026-06-24T14:11:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Build A Harness – open-source, modular harness layer for AI agents

- HN: [48660207](https://news.ycombinator.com/item?id=48660207)
- Source: [buildaharness.com](https://buildaharness.com/)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T14:11:14Z

## Translation

タイトル: Show HN: Build A Harness – AI エージェント用のオープンソースのモジュラー ハーネス レイヤー
記事のタイトル: ハーネスの構築 — AI エージェント ハーネスのビジュアル キャンバス
説明: AI エージェント ハーネス用の Apache 2.0 ビジュアル キャンバス。一度描画すれば、LangGraph、CrewAI、Mastra、または MS Agent Framework で実行できます。ラングフューズの可観測性、HITL、REST/MCP/A2A。

記事本文:
ハーネスを構築する
ハーネス
ノードライブラリ
始めましょう
仕組み
フレームワーク
例
貢献する
よくある質問
ハーネスエンジニアリング
GitHub
オープンソース · ビジュアルキャンバス · APACHE 2.0
設計、テスト、導入
AI エージェントの活用。
エージェントが必要とするノードを選択します。グラフを描きます。任意のフレームワーク上で実行できます。
ワークフローはプロンプトをルーティングします。ハーネスは、エージェントが何を信じ、何を許可され、どのように回復するかを管理します。キャンバス上に構築し、任意のフレームワークにコンパイルし、Langfuse ですべての決定をトレースし、REST エンドポイント、MCP ツール、または A2A エージェントとしてデプロイします。
図: 完全なビルド A ハーネスと比較した単純なエージェント ループ (入力 → LLM 呼び出し → ツール呼び出し → 出力) — 呼び出し元の状態、ワールド モデル、推論、5 層の制御層、計画、実行、9 層の検証、リカバリ、メモリ、オプションの学習、および出力レビューアー パスを追加する 22 ノードにわたる 11 層のアーキテクチャ。
ハーネスの役割
ワークフローではできません。
ワークフローはプロンプトをノードからノードにルーティングします。ハーネスは、AI が何を信じるか、何が許可されるか、AI 自身の間違いをどのように発見するか、次回のために何を学ぶかを管理します。 3 つのノードまたは 11 つのノードを使用します。どちらでも同じ FlowSpec が実行されます。
world_model ノードを追加すると、エージェントは、単に LLM に質問して期待するのではなく、型付けされた信念を追跡し、矛盾を検出し、行動する前に 4 つの生成ソースからの仮説を評価します。
control_state ノードをドロップすると、エージェントはすべてのアクション (通常 → 警戒 → ブロック) を制御する 5 層のリゾルバーを取得します。診断上の健康ベクトルがそれを推進します。デッドロックを検出すると、デッドロックのエスカレーションが永久に阻止されます。
verify_gate は、すべてのアクションの前に 9 つのチェックを実行します。これを reviewer_pass と組み合わせて、戻る前に敵対的なレビューと契約の検証を行います。信頼しますが、検証し、実際にそれを実行します。
si の回復を追加

x 名前付き戦略、型指定された障害検出、およびローカルとグローバルの再計画。 exp_store をアタッチすると、エージェントは今後の実行で成功した分解を再利用します。
27 ノード タイプ · 14 実行 + 13 ハーネス · 4 フレームワーク アダプター · Langfuse 可観測性
必要なすべてのレイヤー。 1 つまたはすべてを使用します。
27 種類のノード — 14 種類の実行、13 種類のハーネス。エージェントが必要とするものを使用します。最小限のハーネスは 3 つのノードで、フルスタックは 11 層まで実行されます。どちらも有効な FlowSpec です。
図: 同じノード タイプが異なるハーネス グラフに構成されています。最小限のハーネス (入力 → llm_call → verify_gate → リカバリ → 出力) と完全な 11 層ハーネスです。どちらも、LangGraph、CrewAI、Mastra、または Microsoft Agent Framework にコンパイルされる有効な FlowSpec です。
world_model 、仮説セット、証拠ストア、制御状態、および 6 つの状態のタスクグラフ — 行動する前に考えるエージェントのインフラストラクチャ。
9 つのレイヤーを備えた verify_gate、実行前レビュー ゲート、 Recovery による 6 つの名前付き回復戦略、コンテキスト圧縮、クロスラン再利用のための exp_store、および敵対的 reviewer_pass 。
キャンバス上の 13 のハーネス ノード タイプ、ハーネス固有のスパンを使用した Langfuse トレース、4 つすべてのランタイムのフレームワーク アダプター、およびすべての組み合わせにわたるエンドツーエンドのテスト。
稼働中
2つのコマンド。
Build A Harness は完全に Docker 経由でマシン上で実行されます。クラウド アカウントもサインアップも必要ありません。リポジトリを複製し、2 つのコマンドを実行するだけで、数分でハーネスを描画できます。
# 1. シークレットを生成し、環境を構成する
./scripts/setup-env.sh
# 2. すべてのサービスを開始する
ドッカーの構成
何が始まるのか
ビジュアル ハーネス エディター。 27 のノード タイプ (ワールド モデル、制御状態、検証ゲート、リカバリを含む 14 個の実行ノードと 13 個のハーネス ノード) を使用して設計します。
FlowSpec を LangGraph、CrewAI、Mastra にコンパイルします

、または Microsoft Agent Framework を使用してハーネスを実行します。各ハーネスを REST エンドポイント、MCP ツール、および A2A エージェントとして同時に公開します。
ハーネスを意識した可観測性ダッシュボード。すべての実行は、ワールド モデルの遷移、制御状態の変更、回復戦略のアクティブ化など、4 つのフレームワークすべてにわたって自動的にトレースされます。
全部で9つのサービス。 scripts/setup-env.sh は、すべてのシークレットと構成を自動的に処理します。必要なのは、LLM API キーを指定することだけです (または、それをスキップして、代わりに無料のローカル モデルを使用します)。
OpenAI、Claude を使用するか、実行します。
ローカルでは完全に無料です。
Build A Harness は、すべての AI 呼び出しを LiteLLM (任意のモデル プロバイダーと連携するプロキシ) 経由でルーティングします。ワークフローでモデルを選択します。残りは自動的に処理されます。
OPENAI_API_KEY を追加し、任意のフローで gpt-4o または gpt-4o-mini を使用します。
ANTHROPIC_API_KEY を追加し、 claude-sonnet 、 claude-haiku 、または claude-opus を使用します。
misstral 、 qwen3 、または qwen2.5-coder をローカルで実行します。 API キーもコストも、マシンからデータが流出することもありません。
1 つの構成ファイルを編集して、セルフホスト型または微調整されたモデルを含む、OpenAI 互換のモデルまたはエンドポイントを追加します。
API キーなしで試してみませんか? Ollama をインストールし、モデル ( ollama pull mistral ) をプルし、 ./scripts/setup-ollama.sh を実行します。有料アカウントは必要なく、4 つのフレームワークすべてをエンドツーエンドでテストします。
ハーネスを一度引き出します。
任意のエンジンで実行してください。
Build A Harness は、すべてのハーネスを FlowSpec (所有するオープン JSON 形式) に保存します。キャンバスに描画すると、仕様が自動的に書き込まれます。同じファイルは、LangGraph、CrewAI、Mastra、または MAF にコンパイルされます。 verify_gate 、 world_model 、または Recovery などのハーネス ノードを追加すると、それらもコンパイルされます。書き換えやロックインは必要ありません。
{
"ワークフロー": "サポートトリアージ",
「ノード」: [
{
"タイプ": "llm_call",
"prompt": "重大度を分類"
}、
{
"タイプ": "条件",
"ルート": "こんにちは

gh_優先度"
}、
{
"タイプ": "ヒューマンレビュー"
}
]、
「テレメトリー」: {
"プロバイダー": "ラングヒューズ"
}
}
構築できるもの
最小限のハーネスは、入力→ llm_call → 出力だけです。必要に応じて、verify_gate とリカバリを追加します。すべての組み合わせが有効な FlowSpec です。
ガードレール、再試行ロジック、および HITL 承認チェックポイントは、後から追加するラッパーではなく、ファーストクラスのノード タイプです。
ハーネスは JSON ファイルです。チェックイン、共有、キャンバスへのロード、または CI パイプラインへのインポートを行います。
Langfuse は、セットアップを行わずに、4 つのフレームワークすべてのすべてのステップ (プロンプト、ツール呼び出し、ハーネス移行、障害、コスト) をトレースします。
チームで AI フレームワークを使用する
すでに知っています。
1 つおきのビジュアル キャンバスは 1 つのランタイムに関連付けられます。 Build A Harness は、ハーネス設計を FlowSpec を介した実行から分離するため、アーキテクチャを毎回再構築することなく、ランタイムの実験、移行、比較を行うことができます。
ステートフル エージェント オーケストレーションの実稼働標準。 LangGraph のグラフ モデルは FlowSpec にきれいにマップされます。条件付きエッジは書き換えなしでコンパイルされ、 world_model 、 control_state 、 verify_gate スロットなどのノードを変換レイヤーのない標準グラフ ノードとして利用します。
Postgres または Redis へのチェックポイント - 失敗後に任意のノードから再開
完全なトークンとグラフ更新のストリーミング
最適: RAG エージェント、複雑なマルチステップのステートフル フロー
役割ベースのマルチエージェント パターンに関する最大のコミュニティ。 Agent_role ノードは、CrewAI クルー、エージェント、およびタスクに直接マップされます。キャンバス デザインは、翻訳なしでクルーになります。ファンアウト用のParallel_fork、ベクトル取得用のmemory_read、およびネイティブ ベクトル ストアのサポートが含まれています。
ロールベースのエージェントはネイティブにコンパイルされます - ラッピングは必要ありません
top-k およびしきい値による完全なメモリ (ベクター) サポート
最適なチーム: 専門エージェント チーム、リサーチ チーム、並行レブ

なるほど
ファーストクラスのハーネスをサポートする唯一の TypeScript ネイティブ オーケストレーターです。完全な HITL の一時停止と再開、完全なストリーミング (トークンとグラフの更新)、およびネイティブ ベクター メモリはすべて同じ FlowSpec からのもので、後で切り替える場合は LangGraph または CrewAI でも実行されます。
フル HITL — 一時停止、状態の検査、キャンバスからの再開
フルストリーミングおよびネイティブベクターメモリ
最適: TypeScript / Node.js チーム、JS ネイティブ スタック
セマンティック カーネルと AutoGen を 1 つの SDK に統合し、エンタープライズ .NET、Python、Java チームを 1 つのアダプターでカバーします。 Agent_role と Agent_debate は AgentGroupChat にネイティブにマップされます。 HITL は、_HitlPause 例外パターンを介して実行されます。 Dapr または Orleans を介した耐久性のある状態のチェックポイント。
AgentGroupChat ネイティブ — ディベートと役割のパターンが直接コンパイルされます
Dapr または Orleans による永続的な実行
最適: エンタープライズ .NET、Java、または Python チーム
既製ハーネス 5 個
フォークして構築します。
このリポジトリには、実際に動作する 5 つのハーネスが同梱されています。それぞれがノード タイプの異なる構成を示しています。キャンバスで開き、実行し、ユースケースに合わせて変更します。
AI が何をしたかを正確に把握する —
そしてそれがなぜそうなったのか。
Langfuse は、Build A Harness に組み込まれているオープンソースの可観測性プラットフォームです。すべての実行は 4 つのランタイムすべてにわたって自動的にトレースされます。構成は必要ありません。 1 つのダッシュボードで、すべてのプロンプト、ツールの呼び出し、決定、失敗、回復アクションがカバーされます。
Build A Harness は、標準のトレース ライブラリでは生成されないハーネス固有のスパンで Langfuse を拡張します。
LangGraph と Mastra で同じ FlowSpec を実行し、トレースを並べて比較します。 Langfuse はローカルの localhost:3001 で実行されます。データはマシンから出ません。
ワンクリックでハーネスを公開します
API、Claude ツール、またはエージェント。
1 つの API 呼び出しで、REST エンドポイント、MCP という 3 つの方法でハーネスを同時に公開します。

ツールと A2A エージェント - したがって、システムが必要とするものはすべて、理にかなった方法で呼び出すことができます。
どのアプリでも、標準の HTTP POST を使用してフローをトリガーできます。特別な SDK は必要ありません。フル ハーネスを実行するには、curl コマンドまたはフェッチ呼び出しだけが必要です。
Model Context Protocol (MCP) を使用すると、AI アシスタントが推論中に外部ツールを呼び出すことができます。 MCP サーバーとして公開すると、Claude Desktop、Cursor、および任意の MCP クライアントが、追加のコードを使用せずに、11 層の制御アーキテクチャを含むフル ハーネスをプロンプトから直接呼び出すことができることを意味します。
エージェント間 (A2A) プロトコルは、AI エージェントの相互運用性のためのオープン スタンダードです。 A2A エージェントとして公開すると、Google ADK、OpenAI Agents SDK、Claude Agent SDK、または任意の A2A ランタイム上に構築された他のエージェントが、エージェント カードを介してハーネスを検出し、カスタム アダプターを使用せずにハーネスにタスクを委任できることを意味します。
@buildaharness/canvas npm パッケージを使用して、ビジュアル エディターを独自のアプリにドロップします。
MCP は、AI アシスタントがツールを呼び出し、推論中にコンテキストを読み取ることを可能にする Anthropic のオープン スタンダードです。 Build A Harness がハーネスを MCP サーバーとして公開すると、各ハーネスは型指定された入出力スキーマを持つ呼び出し可能なツールとして公開されます。 MCP クライアント (Claude Desktop、Cursor、Claude Agent SDK) は、追加のセットアップを行わずに MCP を検出して呼び出すことができます。
これが実際に意味すること: Claude Desktop のユーザーが「この契約を要約し、異常な条項にフラグを立てる」と入力すると、Claude はローカル ツールであるかのように、コンテンツ モデレーション ハーネス (HITL コントロール、 verify_gate 、および Langfuse トレースを備えたもの) を自動的に呼び出します。
A2A は、AI エージェントの相互運用性のための Google のオープン プロトコルです。 A2A エージェントは、エージェント カード (その機能、入力、出力の機械可読な説明) を発行します。他のエージェントは、そのフレームワークに関係なく、

内部アーキテクチャを知らなくても、カードを検出し、標準の HTTP インターフェイス経由でエージェントを呼び出します。
これが実際に意味すること: Google ADK 上に構築されたリサーチ エージェントは、Build A Harness フローに「並行リスク評価」サブタスクを委任し、どちらのエージェントもカスタム アダプターを必要とせずに構造化された結果を受け取ることができます。 Google ADK、OpenAI Agents SDK、Claude Agent SDK、およびあらゆる A2A ランタイムと互換性があります。
Build A Harness はパブリック アルファ版です。
それを使って構築してください。
キャンバス、フレームワーク アダプター、可観測性レイヤー、およびフル ハーネス ノード ライブラリは動作しており、すぐに使用できます。実際のハーネス、バグレポート、貢献によって、何が修正され、何が次に優先されるかが決まります。
パブリックアルファ — 現在および計画されている作業は公開されています
キャンバスとアダプターのレイヤーは安定していますが、アルファ版です。API は変更される可能性があり、Docker Compose の動作は異なる可能性があり、あまり一般的ではないノードの組み合わせにおけるエッジ ケースはまだ完全にはカバーされていません。フルハーネス推論および制御アーキテクチャが実装され、テストされています。実行して、壊して、必要なものをお知らせください。すべてのレポートは、何を優先するかを決定します。
1 つの Docker 構成で 9 つのサービスすべてが開始されます。おもちゃの例ではなく、実際のフロー、つまり実際のユースケースを指します。何かが壊れたり、静かにクラッシュしたり、間違った出力が生成されたりした場合は、FlowSpec JSON を使用してバグ レポートを開きます。

[切り捨てられた]

## Original Extract

Apache 2.0 visual canvas for AI agent harnesses. Draw once, run on LangGraph, CrewAI, Mastra, or MS Agent Framework. Langfuse observability, HITL, REST/MCP/A2A.

Build A Harness
Harness
Node library
Get started
How it works
Frameworks
Examples
Contribute
FAQ
Harness engineering
GitHub
OPEN SOURCE · VISUAL CANVAS · APACHE 2.0
Design, test, and deploy
AI agent harnesses.
Pick the nodes your agent needs. Draw the graph. Run on any framework.
A workflow routes prompts. A harness governs what your agent believes, what it's allowed to do, and how it recovers. Build one on a canvas, compile to any framework, trace every decision with Langfuse, and deploy as a REST endpoint, MCP tool, or A2A agent.
Diagram: a simple agent loop (input → LLM call → tool call → output) compared with a full Build A Harness — an 11-layer architecture across 22 nodes that adds caller state, a world model, reasoning, a 5-tier control layer, planning, execution, 9-layer verification, recovery, memory, optional learning, and an output reviewer pass.
A harness does what a
workflow can't.
A workflow routes prompts from node to node. A harness governs what the AI believes, what it's allowed to do, how it catches its own mistakes, and what it learns for next time. Use three nodes or eleven — the same FlowSpec runs either.
Add a world_model node and your agent tracks typed beliefs, detects contradictions, and evaluates hypotheses from four generation sources before acting — instead of just asking the LLM and hoping.
Drop in a control_state node and your agent gains a five-tier resolver that governs every action — NORMAL → CAUTIOUS → BLOCKED. Diagnostic health vectors drive it; deadlock detection stops it escalating forever.
A verify_gate runs nine checks before every action. Pair it with reviewer_pass for adversarial review and contract validation before every return. Trust, but verify — and actually enforce it.
Add recovery for six named strategies, typed failure detection, and local vs global replanning. Attach exp_store and your agent reuses successful decompositions across future runs.
27 node types · 14 execution + 13 harness · 4 framework adapters · Langfuse observability
Every layer you need. Use one or all.
27 node types — 14 execution, 13 harness. Use what your agent needs: a minimal harness might be three nodes, the full stack runs to eleven layers. Both are valid FlowSpec.
Diagram: the same node types compose into different harness graphs — a minimal harness (input → llm_call → verify_gate → recovery → output) and a full 11-layer harness — both valid FlowSpec that compile to LangGraph, CrewAI, Mastra, or Microsoft Agent Framework.
world_model , hypothesis_set , evidence_store , control_state , and a six-state task_graph — the infrastructure for an agent that thinks before it acts.
verify_gate with nine layers, pre-execution review gate, six named recovery strategies via recovery , context compression, exp_store for cross-run reuse, and adversarial reviewer_pass .
Thirteen harness node types on the canvas, Langfuse tracing with harness-specific spans, framework adapters for all four runtimes, and end-to-end tests across every combination.
Up and running in
two commands.
Build A Harness runs entirely on your machine via Docker. No cloud account, no sign-up — clone the repo, run two commands, and you're drawing harnesses in minutes.
# 1. Generate secrets and configure your environment
./scripts/setup-env.sh
# 2. Start all services
docker compose up
What starts up
The visual harness editor. Design with 27 node types — 14 execution nodes and 13 harness nodes including world model, control state, verify gate, and recovery.
Compiles FlowSpec to LangGraph, CrewAI, Mastra, or Microsoft Agent Framework and runs your harnesses. Publishes each harness simultaneously as a REST endpoint, MCP tool, and A2A agent.
Harness-aware observability dashboard. Every run traced automatically across all 4 frameworks — including world model transitions, control state changes, and recovery strategy activations.
Nine services in total. scripts/setup-env.sh handles all the secrets and configuration automatically — you only need to provide your LLM API key (or skip it and use a free local model instead).
Use OpenAI, Claude, or run
completely free locally.
Build A Harness routes all AI calls through LiteLLM — a proxy that works with any model provider. You pick the model in your workflow; the rest is handled automatically.
Add your OPENAI_API_KEY and use gpt-4o or gpt-4o-mini in any flow.
Add your ANTHROPIC_API_KEY and use claude-sonnet , claude-haiku , or claude-opus .
Run mistral , qwen3 , or qwen2.5-coder locally. No API key, no cost, no data leaving your machine.
Edit one config file to add any OpenAI-compatible model or endpoint — including self-hosted or fine-tuned models.
Want to try it without an API key? Install Ollama , pull a model ( ollama pull mistral ), and run ./scripts/setup-ollama.sh — it tests all four frameworks end-to-end with no paid account needed.
Draw your harness once.
Run it on any engine.
Build A Harness stores every harness in FlowSpec — an open JSON format you own. Draw on the canvas and the spec is written for you. The same file compiles to LangGraph, CrewAI, Mastra, or MAF. Add harness nodes like verify_gate , world_model , or recovery and they compile too — no rewriting, no lock-in.
{
"workflow": "support-triage",
"nodes": [
{
"type": "llm_call",
"prompt": "Classify severity"
},
{
"type": "condition",
"route": "high_priority"
},
{
"type": "human_review"
}
],
"telemetry": {
"provider": "langfuse"
}
}
What you can build
A minimal harness is just input → llm_call → output . Add verify_gate and recovery when you want them. Every combination is valid FlowSpec.
Guardrails, retry logic, and HITL approval checkpoints are first-class node types — not wrappers you bolt on afterwards.
Your harness is a JSON file. Check it in, share it, load it on any canvas, or import into your CI pipeline.
Langfuse traces every step across all four frameworks — prompts, tool calls, harness transitions, failures, and costs — without any setup.
Use the AI framework your team
already knows.
Every other visual canvas is tied to one runtime. Build A Harness separates harness design from execution via FlowSpec — so you can experiment, migrate, or compare runtimes without rebuilding your architecture each time.
The production standard for stateful agent orchestration. LangGraph's graph model maps cleanly to FlowSpec — conditional edges compile without rewriting, and harness nodes like world_model , control_state , and verify_gate slot in as standard graph nodes with no translation layer.
Checkpoint to Postgres or Redis — resume from any node after failure
Full token and graph-update streaming
Best fit: RAG agents, complex multi-step stateful flows
The largest community for role-based multi-agent patterns. agent_role nodes map directly to CrewAI Crew, Agent, and Task — your canvas design becomes a crew without translation. parallel_fork for fan-out, memory_read for vector retrieval, and native vector store support included.
Role-based agents compile natively — no wrapping needed
Full memory (vector) support with top-k and threshold
Best fit: specialist agent teams, research crews, parallel review
The only TypeScript-native orchestrator with first-class harness support. Full HITL pause and resume, full streaming (tokens and graph updates), and native vector memory — all from the same FlowSpec that also runs on LangGraph or CrewAI if you switch later.
Full HITL — pause, inspect state, resume from the canvas
Full streaming and native vector memory
Best fit: TypeScript / Node.js teams, JS-native stacks
The merger of Semantic Kernel and AutoGen in one SDK, covering enterprise .NET, Python, and Java teams in a single adapter. agent_role and agent_debate map natively to AgentGroupChat; HITL runs via the _HitlPause exception pattern; durable state checkpoints via Dapr or Orleans.
AgentGroupChat native — debate and role patterns compile directly
Durable execution via Dapr or Orleans
Best fit: enterprise .NET, Java, or Python teams
Five ready-made harnesses to
fork and build on.
The repo ships with five real, working harnesses. Each demonstrates a different composition of node types — open them in the canvas, run them, and modify them for your use case.
Know exactly what your AI did —
and why it did it.
Langfuse is the open-source observability platform built into Build A Harness. Every run is traced automatically across all four runtimes — no configuration required. One dashboard covers every prompt, tool call, decision, failure, and recovery action.
Build A Harness extends Langfuse with harness-specific spans that no standard tracing library produces:
Run the same FlowSpec on LangGraph and Mastra and compare traces side by side. Langfuse runs locally at localhost:3001 — no data leaves your machine.
One click to publish your harness as
an API, a Claude tool, or an agent.
One API call publishes your harness three ways at once — a REST endpoint, an MCP tool, and an A2A agent — so whatever system needs it can call it the way that makes sense.
Any app can trigger your flow with a standard HTTP POST. No special SDK needed — a curl command or a fetch call is all it takes to run the full harness.
The Model Context Protocol (MCP) lets AI assistants call external tools during inference. Publishing as an MCP server means Claude Desktop, Cursor, and any MCP client can invoke the full harness — including its 11-layer control architecture — directly from a prompt, with no extra code.
The Agent-to-Agent (A2A) protocol is an open standard for AI agent interoperability. Publishing as an A2A agent means other agents — built on Google ADK, OpenAI Agents SDK, Claude Agent SDK, or any A2A runtime — can discover your harness via its Agent Card and delegate tasks to it without a custom adapter.
Drop the visual editor into your own app with the @buildaharness/canvas npm package.
MCP is Anthropic's open standard that lets AI assistants call tools and read context during inference. When Build A Harness publishes a harness as an MCP server, each harness is exposed as a callable tool with a typed input and output schema. Any MCP client — Claude Desktop, Cursor, the Claude Agent SDK — can discover and invoke it with no additional setup.
What this means in practice: a user in Claude Desktop types "summarise this contract and flag unusual clauses" and Claude automatically calls your content-moderation harness — complete with HITL controls, verify_gate , and Langfuse tracing — as if it were a local tool.
A2A is Google's open protocol for AI agent interoperability. An A2A agent publishes an Agent Card — a machine-readable description of its capabilities, inputs, and outputs. Other agents, regardless of their framework, can discover the card and invoke the agent over a standard HTTP interface without needing to know its internal architecture.
What this means in practice: a research agent built on Google ADK can delegate a "parallel risk assessment" subtask to your Build A Harness flow and receive a structured result — with neither agent needing a custom adapter. Compatible with Google ADK, OpenAI Agents SDK, Claude Agent SDK, and any A2A runtime.
Build A Harness is in public alpha.
Build with it.
The canvas, framework adapters, observability layer, and the full harness node library are working and ready to use. Your real harnesses, bug reports, and contributions shape what gets fixed and what gets prioritised next.
Public alpha — current and planned work is in the open
The canvas-and-adapters layer is stable but alpha — APIs may shift, Docker Compose behaviour may vary, and edge cases in less-common node combinations aren't fully covered yet. The full harness reasoning and control architecture is implemented and tested. Run it, break it, and tell us what you need. Every report shapes what gets prioritised.
One docker compose up starts all nine services. Point it at a real flow — your actual use case, not a toy example. When something breaks, crashes silently, or produces wrong output, open a bug report with your FlowSpec JSON, the

[truncated]
