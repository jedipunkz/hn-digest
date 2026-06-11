---
source: "https://artificialanalysis.ai/agents/coding-agents"
hn_url: "https://news.ycombinator.com/item?id=48489034"
title: "Claude Opus is more performant on OpenCode than Claude Code"
article_title: "AI Coding Agent Benchmarks & Leaderboard | Artificial Analysis"
author: "log101"
captured_at: "2026-06-11T13:29:30Z"
capture_tool: "hn-digest"
hn_id: 48489034
score: 4
comments: 0
posted_at: "2026-06-11T11:43:54Z"
tags:
  - hacker-news
  - translated
---

# Claude Opus is more performant on OpenCode than Claude Code

- HN: [48489034](https://news.ycombinator.com/item?id=48489034)
- Source: [artificialanalysis.ai](https://artificialanalysis.ai/agents/coding-agents)
- Score: 4
- Comments: 0
- Posted: 2026-06-11T11:43:54Z

## Translation

タイトル: Claude Opus は OpenCode 上で Claude Code よりもパフォーマンスが高い
記事のタイトル: AI コーディング エージェントのベンチマークとリーダーボード |人工分析
説明: ソフトウェア エンジニアリング タスクにおけるコーディング エージェントの実際のパフォーマンス (コスト、トークン使用量、実行時間など) を測定します。エージェント、モデル、実行設定間でパフォーマンスがどのように変化するかを比較します。

記事本文:
AI コーディング エージェントのベンチマークとリーダーボード |人工分析 人工分析 人工分析モデル
K 人工分析コーディング エージェント ベンチマーク
ソフトウェア エンジニアリング タスクにおけるコーディング エージェントの実際のパフォーマンス (コスト、トークン使用量、実行時間など) を測定します。エージェント、モデル、実行設定間でパフォーマンスがどのように変化するかを比較します。
言語モデルを比較するには、モデルのベンチマークを参照してください。
ベンチマーク 機能の概要 人工分析 コーディング エージェント インデックス
3 つのベンチマークの複合指数:
SWE-Bench-Pro-Hard-AA コード生成、150 タスク By Scale AI
Terminal-Bench v2 エージェント ターミナルの使用、84 のタスク Laude Institute 著
SWE-Atlas-QnA 技術 Q&A、スケール AI 別の 124 タスク
インデックスは、各ベンチマークの 3 回の実行にわたる平均 pass@1 を表します。方法論を見る
コーディング エージェント 一般業務 チャットボット プレゼンテーション OCR データ分析 カスタマー サポート ハイライト
パフォーマンス ハーネスの比較 トークン使用コスト 実行時間 パフォーマンス
人工分析コーディング エージェント インデックス全体のパフォーマンス。
人工分析コーディング エージェント インデックス
モデルエージェントによる色分け 21 モデル中 14 このメトリクスの意味
Artificial Analysiscoding Agent Index は、SWE-Bench-Pro-Hard-AA、ターミナルベンチ v2、および SWE-Atlas-QnA から構築された複合スコアです。
これは簡単な比較に役立ちますが、評価ごとの内訳と併せて読む必要があります。同様のインデックス値を持つ 2 つのエージェントでも、リポジトリ タスク、ターミナル ワークフロー、およびルーブリック ベースの評価にわたって異なる強みを持つ可能性があります。
Claude Opus 4.7 のハーネス別人工分析コーディング エージェント インデックス。
ハーネスの比較: 人工分析コーディング エージェント インデックス
このチャートは、Claude Opus 4.7 での基礎となるモデル定数を保持し、Cursor、Clau を含むさまざまなコーディング エージェント ハーネス間でのパフォーマンスを比較しています。

コード、およびオープンコード。
Artificial Analysis コーディング エージェント インデックス全体のトークン消費量 (総使用量、トークンの組み合わせ、効率、ベンチマークごとの内訳など)。
21 モデル中 14 モデル 出力 キャッシュ 入力 入力 プロンプト キャッシュのヒット率はプロバイダーのルーティングによって大きく異なる可能性があり、実効コストが大幅に変化する可能性があります。入力トークン
プロンプト キャッシュから提供されなかったプロンプト、指示、ツール コンテキスト、タスク コンテキストなど、モデルに送信されたキャッシュされていない入力トークン。
毎回完全に新しい入力として処理されるのではなく、テレメトリが利用可能なときにプロバイダー プロンプト キャッシュを通じて請求されるプロンプト トークンを再利用します。
一部のプロバイダーは、繰り返されるリクエストを異なるバックエンド レプリカ間でルーティングします。プロンプト キャッシュ状態がこれらのレプリカ間で一貫して共有されていない場合、ベンチマーク タスク フローがそれ以外は同一であっても、モデルが受け取るキャッシュ ヒットが少なくなる可能性があります。
キャッシュの再利用率を高めるためにカスタム リレー ヘッダーやプロバイダー固有のアフィニティ制御を追加することはありません。これは、ベンチマークが一般的なユーザー セットアップをあまり代表しなくなるためです。その結果、報告されるコストは、最適化されたベストケースのキャッシュ シナリオではなく、構成されたプロバイダー パスを通じて観察されたキャッシュの動作を反映しています。
タスク中に目に見える応答でモデルによって返されたトークン。
人工分析コーディング エージェント インデックスと総トークンの比較
モデルエージェントによる色分け 21 モデル中 14 Anthropic OpenAI Cursor Z.ai Moonshot AI DeepSeek Google このチャートの見方
各点はコーディング エージェントのバリアントを表します。右に行くほどベンチマークのパフォーマンスが高く、左に行くほどトークン使用量が低いことを意味します。左上に向かうエージェントは、一定のパフォーマンス レベルで使用するトークンの量が少なくなります。
キャッシュ書き込み価格を含む、現在のトークンごとの API 価格に基づく Artificial Analysis コーディング エージェント インデックス全体のコスト

可能な場合はキャッシュ割引も可能です。多くのユーザーは、トークンごとの支払いではなく、サブスクリプション プランを通じてコーディング エージェント ハーネスにアクセスします。
モデルエージェントによるカラー 21 モデル中 14 コストを測定する内容
このグラフは、SWE-Bench-Pro-Hard-AA、ターミナルベンチ v2、および SWE-Atlas-QnA を組み合わせた Artificial Analysiscoding Agent Index 全体のタスクあたりのペイパートークン API コストの平均を示しています。
該当する場合、そのコスト モデルには、すべてのプロンプト トークンを同じ非キャッシュ入力レートで請求されるかのように扱うのではなく、標準入力価格、割引されたキャッシュ入力価格、個別のキャッシュ書き込み料金、および出力価格が含まれます。
これは、消費者向けプランの価格やシステムを運用環境に展開するための運用コスト全体ではなく、トークンごとに支払う API コストを示すことを目的としています。インフラストラクチャ、エンジニアリング、および監督のコストは、この指標の焦点では​​ありません。
人工分析コーディング エージェント インデックスとタスクあたりのコスト
モデルエージェントによる色分け 21 モデル中 14 Anthropic OpenAI Cursor Z.ai Moonshot AI DeepSeek Google このチャートの見方
各点はコーディング エージェントのバリアントを表します。右に行くほどベンチマークのパフォーマンスが高く、グラフの下に行くほどタスクあたりの平均コストが低いことを意味します。最も効率的なエージェントは右下に位置し、低コストでより強力な結果をもたらします。
Artificial Analysis コーディング エージェント インデックス全体のアクティブなエージェント ランタイム。
モデルエージェントごとに色分け 21 モデル中 14 実行時間の測定内容
このグラフでは、エージェント ウォール タイム、つまりエージェント プロセスが各タスクでアクティブに実行されていた時間を使用します。
これには、環境の起動、ベリファイアやジャッジの時間、その他のハーネスのオーバーヘッドは含まれていないため、エージェント自体が動作していた時間をより明確に比較できます。
人工分析コーディング エージェント インデックスと実行時間
モデルエージェントによるカラー 21 モデル中 14 Anthropic OpenAI Cursor Z.ai

Moonshot AI DeepSeek Google このチャートの見方
各点はコーディング エージェントのバリアントを表します。右に行くほどベンチマークのパフォーマンスが高く、グラフの下に行くほどタスクごとの平均エージェント実行時間が短いことを意味します。右下に向かうエージェントは、エージェントのアクティブ時間が短くなり、より強力な結果をもたらします。
人工分析コーディング エージェント インデックスとは何ですか?
人工分析コーディング エージェント インデックスは、このページの公開ベンチマーク スイート全体におけるコーディング エージェントのパフォーマンスの複合スコアです。現在、SWE-Bench-Pro-Hard-AA、ターミナル-ベンチ v2、および SWE-Atlas-QnA を組み合わせて、実装、ターミナル ワークフロー、およびリポジトリ理解のパフォーマンスを単一のヘッドライン メトリクスでキャプチャします。
現在インデックスに含まれているベンチマークはどれですか?
現在の公開インデックスには、SWE-Bench-Pro-Hard-AA、ターミナルベンチ v2、および SWE-Atlas-QnA が含まれています。これらのベンチマークは、同じタスク形式を繰り返すのではなく、コーディング エージェント ワークフローのさまざまな部分に重点を置くため、結合されます。
これらのベンチマークは実際にどのような種類のタスクをテストしているのでしょうか?
パブリック ベンチマーク スイートは、いくつかのソフトウェア エンジニアリング タスク スタイルを組み合わせています。一部のタスクは、コードベースの読み取り、アーキテクチャまたは動作の理解、技術的な正しい答えの生成に焦点を当てた Q&A およびリポジトリ理解タスクです。一部は、コードの変更を必要とする実装およびバグ修正タスクであり、古典的な「パッチを適用する」フレームワークに近いものです。一部は、エージェントがシェル駆動環境をナビゲートできるか、ツールを正しく実行できるか、複数ステップのコマンドライン ワークフローを完了できるかどうかをテストするターミナル ワークフロー タスクです。また、このスイートは、バイナリ結果とルーブリック スコアの部分評価結果を効果的に組み合わせています。これは、エージェントが難しいタスクを完全に解決しなくても、有益な進歩を示すことができるため重要です。
Q&A形式でどのように行うか

タスクは SWE-Bench-Pro-Hard-AA スタイルのタスクとは異なりますか?
Q&A スタイルのタスクでは、リポジトリの理解、コードの読み取り、動作のトレース、および正しい技術的な説明の作成に重点が置かれます。 SWE-Bench-Pro-Hard-AA スタイルのタスクは、実用的な変更の出荷に近づいています。エージェントはタスクを理解し、リポジトリを移動し、ファイルを正しく編集し、実行制約の下で評価者またはテストベースの結果を満たす必要があります。これらは関連する機能ですが、同一ではありません。エージェントはリポジトリの推論には優れていても、信頼性の高いパッチの実行には弱い、またはその逆の場合もあります。これが、複合インデックスをベンチマークごとのチャートと並行して解釈する必要がある理由の 1 つです。
エージェントは各ベンチマークでどのようにスコア付けされますか?
ベンチマーク ページでは、平均 pass@1 を使用してコンポーネント スコアが報告されます。これは評価者がタスクに割り当てたスコアであり、ベンチマークに応じて 2 点評価または部分評価のいずれかになります。合格した実行は、自動的に解決されたタスクと同じになるわけではありません。実行は正常に完了しても、スコアが 0 になる場合があります。現在の方法論では、タスクが合格して肯定的なスコアを受け取った場合にのみ、タスクが解決されたとみなされます。これは、SWE-Atlas-QnA などのルーブリック スコア付きタスクの場合に特に重要です。部分的な評価によって、厳密な合否判定基準では失われる有用な進捗状況を把握できる場合があります。
全体的なインデックスの重み付けはどのように行われますか?
インデックスは、現在のパブリック スイートを構成する構成済みのベンチマーク コンポーネントから計算されます。現在の Artificial Analysis コーディング エージェント インデックスの場合、公開手法は利用可能なコンポーネントのベンチマーク スコア全体の単純平均であり、現在の構成では、含まれるベンチマーク コンポーネントをその複合の等しいコンポーネントとして扱います。ベンチマーク手法はカバレッジが向上するにつれて進化する可能性があるため、比較可能性が最も重要です。

時代を超えた絶対スコアとしてではなく、公開されたベンチマーク スイートとその現在のコンポーネント セット内で表現されます。
実行時間とはどういう意味ですか?
このページの実行時間は、生のモデルのレイテンシーだけではなく、タスクごとの実時間タスクの実行時間を指します。これは、エージェント ワークフロー全体を実行するためにユーザーが直面する時間コストを反映することを目的としています。これには、推論、ツール呼び出しの発行、ファイルの読み取りと書き込み、シェル ステップの実行、モデルの応答の待機に費やされる時間が含まれます。したがって、エージェントの基礎となるモデルは高速であっても、ワークフローが長くなったり、ツールの負荷が高くなったりすると、全体としては遅くなる可能性があります。
トークンの使用は何を意味し、なぜ重要なのでしょうか?
トークン使用量は、ベンチマーク スイート全体で観察されたタスクごとの平均トークン消費量です。このページでは、入力トークン、キャッシュトークン、出力トークンに分けて説明します。入力トークンは、プロンプト、指示、ツール コンテキスト、タスク コンテキストなど、モデルに送信されるトークンです。キャッシュ トークンは、プロバイダーがテレメトリを公開するときにプロンプ​​ト キャッシュを通じて再利用されるプロンプト トークンです。出力トークンは、モデルの応答内で生成されるトークンです。トークンの使用量が重要なのは、トークンの使用量によって多くの場合コストが発生し、エージェントが作業を完了するために消費するコンテキストの量も示されるためです。ただし、トークンの効率とコストは同一ではありません。これは、プロバイダーによってトークン カテゴリの価格設定が異なり、キャッシュによって請求額が大幅に変わる可能性があるためです。
私のユースケースでは、インデックスが高いエージェントの方がさらに悪い可能性があるのはなぜですか?
インデックス スコアが高いほど、含まれるベンチマーク ミックス全体でパフォーマンスが向上していることを意味しますが、エージェントがすべてのワークフローに最適であることを意味するわけではありません。このインデックスは、ベンチマーク品質全体のバランスを示すものであり、特定のレイテンシー、コスト、ツール、タスク タイプの優先順位を直接測定するものではありません。実際の選択は、ワークフローがリポジトリ Q&A、パッチ適用、ターミナのどれに似ているかによって決まります。

l 実行、および IDE 統合、モデルの可用性、信頼性などの実際的な制約に基づきます。
これらのタスクはどの程度現実的ですか?各エージェントにはどのような設定が使用されましたか?
これらのベンチマークは、リポジトリ、ツール、複数ステップのワークフロー、評価者ベースの結果にわたるコーディング エージェントのパフォーマンスを測定します。このページの結果には、一般的な製品名だけでなく、評価された特定のエージェント バリアントが反映されています。モデルの選択、設定、実行構成によって結果が大きく変わる可能性があります。そのため、結果では単一のエージェント ファミリが複数のバリアントで表示される場合があります。ベンチマークの実行、タスク レベルのスコアリング、および方法論の詳細については、コーディング エージェントのベンチマーク方法論のページを参照してください。コーディング エージェントのベンチマーク手法を表示する
新しい記事に関する通知を受け取る
X LinkedIn YouTube Rednote Discord © 2026 人工分析

## Original Extract

We measure real-world performance of coding agents on software engineering tasks, including cost, token usage, and execution time. We compare how performance changes across agents, models, and execution settings.

AI Coding Agent Benchmarks & Leaderboard | Artificial Analysis Artificial Analysis Artificial Analysis Models
K Artificial Analysis Coding Agent Benchmarks
We measure real-world performance of coding agents on software engineering tasks, including cost, token usage, and execution time. We compare how performance changes across agents, models, and execution settings.
To compare language models see our model benchmarks .
Benchmarks Features Overview Artificial Analysis Coding Agent Index
Composite index of 3 benchmarks:
SWE-Bench-Pro-Hard-AA Code generation, 150 tasks By Scale AI
Terminal-Bench v2 Agentic terminal use, 84 tasks By Laude Institute
SWE-Atlas-QnA Technical Q&A, 124 tasks By Scale AI
Index represents the average pass@1 across 3 runs of each benchmark. View methodology
Coding Agents General Work Chatbots Presentations OCR Data Analysis Customer Support Highlights
Performance Harness Comparison Token Usage Cost Execution Time Performance
Performance across the Artificial Analysis Coding Agent Index.
Artificial Analysis Coding Agent Index
Color by Model Agent 14 of 21 models What this metric means
The Artificial Analysis Coding Agent Index is a composite score built from SWE-Bench-Pro-Hard-AA, Terminal-Bench v2, and SWE-Atlas-QnA.
It is useful for quick comparison, but it should be read alongside the per-eval breakdowns. Two agents with similar index values can still have different strengths across repository tasks, terminal workflows, and rubric-based evaluations.
Artificial Analysis Coding Agent Index by harness for Claude Opus 4.7.
Harness Comparison: Artificial Analysis Coding Agent Index
This chart holds the underlying model constant at Claude Opus 4.7 and compares how it performs across different coding-agent harnesses, including Cursor, Claude Code, and OpenCode.
Token consumption across the Artificial Analysis Coding Agent Index, including total usage, token mix, efficiency, and per-benchmark breakdowns.
14 of 21 models Output Cached Input Input Prompt cache hit rates can vary significantly by provider routing, which can materially change effective cost. Input tokens
Non-cached input tokens sent to the model, including prompts, instructions, tool context, and task context that were not served from prompt cache.
Reused prompt tokens billed through provider prompt caching when that telemetry is available, rather than being processed as a fully fresh input each time.
Some providers route repeated requests across different backend replicas. When prompt cache state is not shared consistently across those replicas, a model may receive fewer cache hits even when the benchmark task flow is otherwise identical.
We do not add custom relay headers or provider-specific affinity controls to force higher cache reuse, because that would make the benchmark less representative of a typical user setup. As a result, reported costs reflect the cache behavior observed through the configured provider path, not an optimized best-case cache scenario.
Tokens returned by the model in its visible response during the task.
Artificial Analysis Coding Agent Index vs. Total Tokens
Color by Model Agent 14 of 21 models Anthropic OpenAI Cursor Z.ai Moonshot AI DeepSeek Google How to read this chart
Each point represents a coding-agent variant. Farther right means higher benchmark performance, while lower token usage appears farther left. Agents toward the upper-left use fewer tokens for a given level of performance.
Cost across the Artificial Analysis Coding Agent Index based on current per-token API pricing, including cache write pricing and cache discounts where available. Many users will access coding agent harnesses through subscription plan offerings rather than pay-per-token.
Color by Model Agent 14 of 21 models What cost is measuring
This chart shows the mean pay-per-token API cost per task across the Artificial Analysis Coding Agent Index, which combines SWE-Bench-Pro-Hard-AA, Terminal-Bench v2, and SWE-Atlas-QnA.
Where applicable, that cost model includes standard input pricing, discounted cached-input pricing, separate cache-write charges, and output pricing rather than treating all prompt tokens as if they were billed at the same uncached input rate.
It is intended to show pay-per-token API cost, not consumer plan pricing or the full operational cost of deploying the system in production. Infrastructure, engineering, and supervision costs are not the focus of this metric.
Artificial Analysis Coding Agent Index vs. Cost per Task
Color by Model Agent 14 of 21 models Anthropic OpenAI Cursor Z.ai Moonshot AI DeepSeek Google How to read this chart
Each point represents a coding-agent variant. Farther right means higher benchmark performance, while lower on the chart means lower mean cost per task. The most efficient agents sit toward the lower-right: stronger results at lower cost.
Active agent runtime across the Artificial Analysis Coding Agent Index.
Color by Model Agent 14 of 21 models What execution time is measuring
This chart uses agent wall time: how long the agent process was actively running on each task.
It does not include environment startup, verifier or judge time, or other harness overhead, so it is a cleaner comparison of how long the agent itself was working.
Artificial Analysis Coding Agent Index vs. Execution Time
Color by Model Agent 14 of 21 models Anthropic OpenAI Cursor Z.ai Moonshot AI DeepSeek Google How to read this chart
Each point represents a coding-agent variant. Farther right means higher benchmark performance, while lower on the chart means shorter mean agent runtime per task. Agents toward the lower-right deliver stronger results in less active agent time.
What is the Artificial Analysis Coding Agent Index?
The Artificial Analysis Coding Agent Index is our composite score for coding-agent performance across the public benchmark suite on this page. It currently combines SWE-Bench-Pro-Hard-AA, Terminal-Bench v2, and SWE-Atlas-QnA to capture implementation, terminal workflow, and repository-understanding performance in a single headline metric.
Which benchmarks are included in the index right now?
The current public index includes SWE-Bench-Pro-Hard-AA, Terminal-Bench v2, and SWE-Atlas-QnA. These benchmarks are combined because they stress different parts of the coding-agent workflow rather than repeating the same task format.
What kinds of tasks are these benchmarks actually testing?
The public benchmark suite mixes several software engineering task styles. Some tasks are Q&A and repository-understanding tasks that focus on reading a codebase, understanding architecture or behavior, and producing a correct technical answer. Some are implementation and bug-fix tasks that require code changes and are closer to the classic make-a-patch-that-works framing. Some are terminal workflow tasks that test whether the agent can navigate a shell-driven environment, execute tools correctly, and complete a multi-step command-line workflow. The suite also mixes effectively binary outcomes with rubric-scored partial-credit outcomes, which matters because an agent can show useful progress on a difficult task without fully solving it.
How do Q&A-style tasks differ from SWE-Bench-Pro-Hard-AA-style tasks?
Q&A-style tasks emphasize repository understanding, code reading, tracing behavior, and producing a correct technical explanation. SWE-Bench-Pro-Hard-AA-style tasks are closer to shipping a working change: the agent has to understand the task, navigate the repository, edit files correctly, and satisfy an evaluator or test-based outcome under execution constraints. Those are related capabilities, but they are not identical. An agent can be strong at repository reasoning and still be weaker at reliable patch execution, or vice versa, which is one reason the composite index should be interpreted alongside the per-benchmark chart.
How are agents scored on each benchmark?
The benchmark page reports component scores using average pass@1. This is the evaluator-assigned score for a task, and depending on the benchmark it can be either binary or partial credit. A passed run is not automatically the same thing as a solved task: a run can complete cleanly and still receive a zero score. In the current methodology, a task is counted as solved only when it passed and received a positive score. This matters especially for rubric-scored tasks such as SWE-Atlas-QnA, where partial credit can capture useful progress that would be lost in a strict pass-fail metric.
How is the overall index weighted?
The index is computed from the configured benchmark components that make up the current public suite. For the current Artificial Analysis Coding Agent Index, the public methodology is a simple average across the available component benchmark scores, and the current configuration treats the included benchmark components as equal components of that composite. Benchmark methodology can evolve as coverage improves, so comparability is best interpreted within the published benchmark suite and its current component set rather than as a timeless absolute score.
What does execution time mean?
Execution time on this page refers to wall-clock task runtime per task, not just raw model latency. It is meant to reflect the user-facing time cost of running the whole agent workflow. That includes time spent reasoning, issuing tool calls, reading and writing files, executing shell steps, and waiting on model responses. So an agent can have a fast underlying model and still be slower overall if its workflow is longer or more tool-heavy.
What does token usage mean, and why does it matter?
Token usage is the average observed token consumption per task across the benchmark suite. On this page we break it out into input, cache, and output tokens. Input tokens are the tokens sent into the model, including prompts, instructions, tool context, and task context. Cache tokens are prompt tokens reused through prompt caching when the provider exposes that telemetry. Output tokens are tokens generated by the model in its response. Token usage matters because it often drives cost and can also indicate how much context an agent consumes to get work done, but token efficiency and cost are not identical because providers price token categories differently and caching can materially change the bill.
Why can a higher-index agent still be worse for my use case?
A higher index score means stronger performance across the included benchmark mix, but it does not mean the agent is best for every workflow. The index is a balance across benchmark quality, not a direct measure of your specific latency, cost, tooling, or task-type priorities. Real-world choice still depends on whether your workflow looks more like repository Q&A, patching, or terminal execution, and on practical constraints such as IDE integration, model availability, and reliability.
How realistic are these tasks, and what setup was used for each agent?
These benchmarks measure coding-agent performance across repositories, tools, multi-step workflows, and evaluator-based outcomes. Results on this page reflect specific evaluated agent variants, not just generic product names: model choice, settings, and execution configuration can materially change outcomes, which is why a single agent family may appear in multiple variants in the results. For more background on benchmark runs, task-level scoring, and methodology, see the coding-agents benchmarking methodology page. View the coding-agents benchmarking methodology
Get notified about new articles
X LinkedIn YouTube Rednote Discord © 2026 Artificial Analysis
