---
source: "https://arpitbhayani.me/blogs/ai-topological-sort/"
hn_url: "https://news.ycombinator.com/item?id=48371047"
title: "AI Workflows Need Topological Sort"
article_title: "AI Workflows Need Topological Sort"
author: "random42"
captured_at: "2026-06-03T00:43:13Z"
capture_tool: "hn-digest"
hn_id: 48371047
score: 2
comments: 0
posted_at: "2026-06-02T14:51:50Z"
tags:
  - hacker-news
  - translated
---

# AI Workflows Need Topological Sort

- HN: [48371047](https://news.ycombinator.com/item?id=48371047)
- Source: [arpitbhayani.me](https://arpitbhayani.me/blogs/ai-topological-sort/)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T14:51:50Z

## Translation

タイトル: AI ワークフローにはトポロジカルなソートが必要
説明: すべての AI ワークフローには依存関係の問題があります。出力を生成するステップ、それらの出力を消費する他のステップ、およびプロデューサーが終了するまでコンシューマを実行できないという厳しい制約があります。順序を間違えると、古いデータを読み取ったり、コンテキストが欠落しているツールを呼び出したり、事前にエージェントをトリガーしたりすることになります。
[切り捨てられた]

記事本文:
AI ワークフローにはトポロジカルなソートが必要
アルピット・バヤニ
ブログ
最初の原則から
--> 紙棚 本棚
セッション
--> コースでの講演 AI ワークフローにはトポロジカルなソートが必要
エンジニアリング、データベース、システム。常に構築しています。
すべての AI ワークフローには依存関係の問題があります。出力を生成するステップ、それらの出力を消費する他のステップ、およびプロデューサーが終了するまでコンシューマを実行できないという厳しい制約があります。順序を間違えると、古いデータを読み取ったり、コンテキストが欠落しているツールを呼び出したり、入力の準備が整う前にエージェントをトリガーしたりすることになります。
有向非巡回グラフ (DAG) は、これに適したモデルです。トポロジカル ソートは、DAG を実行順序に変換します。これらは共に、応用 AI システム実行の基本を形成しており、ワークフローを設計、デバッグ、スケールする際には、第一原理レベルでそれらを理解することが重要です。
AI ワークフローにおける依存関係の問題
実際のドキュメント処理パイプラインを考えてみましょう。
ストレージから生のドキュメントを取得する
エンベディングをベクトルインデックスに保存する
インデックスに対して取得クエリを実行します。
取得したチャンクをプロンプト テンプレートに渡します
各ステップはその前のステップに依存します。ステップ 2 が完了する前にステップ 3 を実行すると、ダーティ テキストが埋め込まれてしまいます。ステップ 4 の前にステップ 5 を実行すると、不完全なインデックスをクエリすることになります。依存関係はオプションの制約ではなく、正確性の制約です。
単純な線形パイプラインでは、ステップ 1 から 8 を順番に実行するだけです。しかし、実際のワークフローは分岐します。いくつかのステップは互いに独立しています。一部のステップは並列サブタスクにファンアウトし、またファンインします。線形リストはすぐに分解されます。
ここで DAG が適切な抽象化になります。
DAG は、ワークフローをノード (タスク) と有向エッジ (依存関係) のセットとして表します。 A から B へのエッジは、「B が開始する前に A が完了する必要がある」ことを意味します。非周期性制約

循環依存関係がないことを意味します。タスクはそれ自体に推移的に依存できません。
DAG としてモデル化された分岐 RAG パイプラインを次に示します。
store_index が終了すると、keyword_filter とretrieve は互いに独立します。これらは並行して実行できます。 merge_context は両方が終了するまで開始できません。線形リストではこれを表現できません。 DAG は可能です。
DAG のトポロジー的順序付けは、すべてのエッジ u → v u \to v u → v に対して、ノード u u u が v v v の前に現れるような、すべてのノードのシーケンスです。生産者は常に消費者に先立ちます。カーンのアルゴリズムは、これを O ( V + E ) O(V + E) O ( V + E ) 時間で計算します。
上記のパイプラインに適用すると、fetch_documents が最初に実行され、call_llm が最後に実行され、keyword_filter とretrieve が merge_context の前に表示されますが、相互に順序付けの制約はありません。その自由こそが並列処理を可能にするのです。
順序を超えたトポロジカルソート
トポロジー順序の同じ「レベル」にあるノード間には依存関係がありません。これらは同時に実行できます。より賢いスケジューラーは、可能な限り早い開始時間に基づいてノードをグループ化します。
def 実行レベル (グラフ: dict[ str , list[ str ]]) -> リスト[list[ str ]]:
in_degree = defaultdict( int )
グラフ内のノードの場合:
グラフ[ノード]のdepの場合:
度[度] += 1
レベル = []
ready = [n はグラフ内の n の場合 if in_degree[n] == 0 ]
準備ができている間:
レベル.追加(準備完了)
next_ready = []
準備ができているノードの場合:
グラフ[ノード]のdepの場合:
度[度] -= 1
in_degree[dep] == 0 の場合:
next_ready.append(dep)
準備完了 = next_ready
リターンレベル
レベル内の各リストは、並行して実行できるタスクのバッチです。 Prefect や Airflow などのツールは、これを正確に計算して、エグゼキュータのスループットを最大化します。
実行前のサイクル検出
ユーザーまたは構成が循環依存関係を宣言した場合、 len(order) != len(graph) は sin の前にそれをキャッチします。

gleタスクが実行されます。これはあると便利なものではありません。サイクルはデッドロックを意味します。A は B を待ち、B は A を待ちます。何も進みません。実行時ではなく定義時に検出するかどうかは、明確なエラー メッセージが表示されるか、午前 2 時にパイプラインがハングするかの違いです。
LangGraph のようなマルチエージェント オーケストレーション フレームワークは、同じ理由で、エージェントの相互作用を DAG としてモデル化します。これは、互いの出力を生成および消費するエージェント間で正しい実行順序を強制するためです。
4 つのエージェントを使用したリサーチ ワークフローを考えてみましょう。
オーケストレーターは、critic_agent が終了するまで Writer_agent をディスパッチできず、また、summaryr_agent が終了するまで Critic_agent をディスパッチすることもできません。トポロジカル ソートでは、依存関係の宣言からこの順序が自動的に生成されます。オーケストレーターはシーケンス ロジックをハードコーディングする必要はありません。
search_agent と retrieval_agent は独立しています。これらは同時に実行できます。 summary_agent は両方を待機します。トポロジの並べ替えではこれが考慮されます。両方の検索エージェントが同じ実行レベルに表示され、susumizer_agent は、両方の残りの依存関係がゼロになった場合にのみ次のレベルに表示されます。
これはスケールします。 10 個のエージェントを追加し、条件付き分岐を追加し、ファンアウトを追加します。DAG モデルとトポロジカルな並べ替えが複雑さを処理します。ハードコードされた順次ディスパッチは行われません。
上流のタスクが失敗した場合や入力が変更された場合でも、パイプライン全体を再実行する必要はありません。影響を受けるノードからグラフを前方に移動し、そのサブグラフ内のノードのみを再計算します。そのサブグラフの外側にあるすべてのノードには、有効なキャッシュされた出力があります。
これはビルド システム ( Bazel 、 Buck ) の標準であり、AI パイプライン フレームワークでも一般的になってきています。 DAG 構造により、扱いやすくなります。明示的な依存関係エッジがなければ、どの下流ノードが影響を受けるかを知ることができません。
ファンインのボトルネックは現実のものです。 10 個の並列タスクがすべてあった場合

1 つのマージ ノードにフィードを送信すると、10 個のマージ ノードのうち最も遅いものが終了するまで、マージ ノードを開始できません。トポロジカルソートは順序を正確に示しますが、作業のバランスを自動的に調整しません。クリティカル パス (依存タスクの最長チェーン) をプロファイリングして、並列処理の向上が実際に制限されている場所を見つけます。
構成内のサイクルはユーザー エラーであり、フレームワークのバグではありません。ワークフロー定義時にそれらを捕捉し、特定された問題のあるサイクルで明らかなエラーを表示し、実行が開始される前にワークフローを拒否する検証を構築します。
DAG は契約です。ここにタスクがあり、その依存関係があり、循環待機がないことが保証されています。トポロジカル ソートは、そのコントラクトを実行可能な実行スケジュールに変換するメカニズムです。
AI ワークフローを設計するときは、最初に依存関係グラフを描画します。どのタスクが真に連続しており、どのタスクが独立しているかを特定します。そのグラフはあなたの仕様です。トポロジカルな順序付けはスケジューラの入力です。他のすべて - 並列処理、サイクル安全性、増分再計算 - は、すでに宣言した構造から外れます。
脚注: DAG は、エッジが実行順序の制約をエンコードする依存関係グラフとして AI ワークフローとマルチエージェント システムをモデル化します。トポロジカル ソートは、このグラフを O ( V + E ) O(V + E) O ( V + E ) 時間で有効なタスク スケジュールに変換し、実行が開始される前に循環依存関係を検出し、どのタスクが並列実行できるかを明らかにします。マルチエージェント オーケストレーションの場合、これは、ハードコーディングされたシーケンス ロジックが必要なく、入力の準備ができた場合にのみエージェントがディスパッチされることを意味します。
これが役に立ち、興味深いと思われた場合は、
RSS フィードを購読すると、新しいフィードを公開したときに通知が届きます。
Razorpay のプリンシパル エンジニア II - Agent Studio の構築、元スタッフ エンジニア

GCP Memorystore & Dataproc、DiceDB の作成者、元 Amazon Fast Data、元 Engg ディレクター。 SREと
Unacademy のデータ エンジニアリング。私は自分自身を通じてエンジニアリングへの好奇心を刺激します
綿密なエンジニアリングビデオ
YouTube
そして私のコース
145,000 人のエンジニアが読んだ Arpit のニュースレター
現実世界のシステム設計、分散システム、または
非常に賢いアルゴリズムを深く掘り下げてみましょう。
LinkedIn で購読する Substack で購読する
このウェブサイトに掲載されているコースは、
リログディープテック株式会社株式会社
203, Sagar Apartment, Camp Road, Mangiral Plot, アムラヴァティ, マハーラーシュトラ州, 444602
GSTIN: 27AALCR5165R1ZF
YouTube (200k) Twitter (100k) LinkedIn (250k) GitHub (6k)
© アルピット・バヤニ、2025

## Original Extract

Every AI workflow is a dependency problem. You have steps that produce outputs, other steps that consume those outputs, and a hard constraint: consumers cannot run before their producers finish. Get the order wrong and you read stale data, call a tool with missing context, or trigger an agent before
[truncated]

AI Workflows Need Topological Sort
Arpit Bhayani
Blogs
From the First Principles
--> Papershelf Bookshelf
Sessions
--> Courses Talks AI Workflows Need Topological Sort
engineering, databases, and systems. always building.
Every AI workflow is a dependency problem. You have steps that produce outputs, other steps that consume those outputs, and a hard constraint: consumers cannot run before their producers finish. Get the order wrong and you read stale data, call a tool with missing context, or trigger an agent before its inputs are ready.
Directed acyclic graphs (DAGs) are the right model for this. Topological sort turns a DAG into an execution order. Together they form a primitive in applied AI system execution, and understanding them at a first-principles level is important when you design, debug, and scale workflows.
The Dependency Problem in AI Workflows
Take a realistic document processing pipeline :
Fetch raw documents from storage
Store embeddings in a vector index
Run a retrieval query against the index
Pass retrieved chunks into a prompt template
Each step depends on the one before it. If you run step 3 before step 2 finishes, you embed dirty text. If you run step 5 before step 4, you query an incomplete index. The dependencies are not optional constraints - they are correctness constraints.
In a simple linear pipeline you can just run steps 1 through 8 in order. But real workflows branch. Some steps are independent of each other. Some steps fan out into parallel subtasks and fan back in. A linear list breaks down fast.
This is where a DAG becomes the right abstraction.
A DAG represents a workflow as a set of nodes (tasks) and directed edges (dependencies). An edge from A to B means “A must complete before B starts.” The acyclicity constraint means there are no circular dependencies - a task cannot transitively depend on itself.
Here is a branching RAG pipeline modeled as a DAG:
keyword_filter and retrieve are independent of each other once store_index finishes. They can run in parallel. merge_context cannot start until both finish. A linear list cannot express this. A DAG can.
A topological ordering of a DAG is a sequence of all nodes such that for every edge u → v u \to v u → v , node u u u appears before v v v . Producers always precede consumers. Kahn’s algorithm computes this in O ( V + E ) O(V + E) O ( V + E ) time.
Applied to the pipeline above, this produces an order where fetch_documents runs first, call_llm runs last, and keyword_filter and retrieve appear before merge_context but without any ordering constraint between each other. That freedom is what enables parallelism.
Topological Sort Beyond Ordering
Nodes at the same “level” of the topological order have no dependency between them. They can run concurrently. A smarter scheduler groups nodes by their earliest possible start time:
def execution_levels (graph: dict[ str , list[ str ]]) -> list[list[ str ]]:
in_degree = defaultdict( int )
for node in graph:
for dep in graph[node]:
in_degree[dep] += 1
levels = []
ready = [n for n in graph if in_degree[n] == 0 ]
while ready:
levels.append(ready)
next_ready = []
for node in ready:
for dep in graph[node]:
in_degree[dep] -= 1
if in_degree[dep] == 0 :
next_ready.append(dep)
ready = next_ready
return levels
Each list in levels is a batch of tasks that can execute in parallel. Tools like Prefect and Airflow compute exactly this to maximize executor throughput .
Cycle Detection Before Execution
If a user or a config declares a circular dependency, len(order) != len(graph) catches it before a single task runs. This is not a nice-to-have. A cycle means a deadlock : A waits for B, B waits for A, nothing makes progress. Detecting it at definition time rather than runtime is the difference between a clear error message and a hung pipeline at 2am.
Multi-agent orchestration frameworks like LangGraph model agent interactions as DAGs for the same reason: to enforce correct execution order across agents that produce and consume each other’s outputs.
Consider a research workflow with four agents:
The orchestrator cannot dispatch writer_agent until critic_agent finishes, and cannot dispatch critic_agent until summarizer_agent finishes. Topological sort produces this order automatically from the dependency declarations. The orchestrator does not need to hardcode sequencing logic.
search_agent and retrieval_agent are independent. They can run simultaneously. summarizer_agent waits on both. The topological sort respects this: both search agents appear in the same execution level, and summarizer_agent appears in the next level only after both have zero remaining dependencies.
This scales. Add ten agents, add conditional branches, add fan-outs - the DAG model and topological sort handle the complexity. Hardcoded sequential dispatch does not.
When an upstream task fails or an input changes, you do not need to re-run the entire pipeline. Traverse the graph forward from the affected node and recompute only the nodes in its subgraph. Every node outside that subgraph has valid cached output.
This is standard in build systems ( Bazel , Buck ) and is increasingly common in AI pipeline frameworks. The DAG structure is what makes it tractable. Without explicit dependency edges you cannot know which downstream nodes are affected.
Fan-in bottlenecks are real. If ten parallel tasks all feed into one merge node, the merge node cannot start until the slowest of the ten finishes. Topological sort tells you the order correctly but does not automatically balance work. Profile the critical path - the longest chain of dependent tasks - to find where parallelism gains are actually limited.
Cycles in configuration are a user error, not a framework bug. Build validation that catches them at workflow definition time, surfaces a clear error with the offending cycle identified, and rejects the workflow before any execution begins.
A DAG is a contract. It says: here are the tasks, here are their dependencies, and here is the guarantee that there are no circular waits. Topological sort is the mechanism that converts that contract into an actionable execution schedule.
When you design an AI workflow, draw the dependency graph first. Identify which tasks are truly sequential and which are independent. That graph is your specification. The topological ordering is your scheduler’s input. Everything else - parallelism, cycle safety, incremental recomputation - falls out of the structure you have already declared.
Footnote: DAGs model AI workflows and multi-agent systems as dependency graphs where edges encode execution order constraints. Topological sort converts this graph into a valid task schedule in O ( V + E ) O(V + E) O ( V + E ) time, detects circular dependencies before execution begins, and reveals which tasks can run in parallel. For multi-agent orchestration, this means agents are dispatched only when their inputs are ready, with no hardcoded sequencing logic required.
If you find this helpful and interesting,
subscribe to my RSS feed and get notified the moment I publish a new one.
Principal Engineer II at Razorpay - building Agent Studio, Ex-staff engg at GCP Memorystore & Dataproc, Creator of DiceDB , ex-Amazon Fast Data, ex-Director of Engg. SRE and
Data Engineering at Unacademy. I spark engineering curiosity through my
no-fluff engineering videos on
YouTube
and my courses
Arpit's Newsletter read by 145,000 engineers
Weekly essays on real-world system design, distributed systems, or a
deep dive into some super-clever algorithm.
Subscribe on LinkedIn Subscribe on Substack
The courses listed on this website are offered by
Relog Deeptech Pvt. Ltd.
203, Sagar Apartment, Camp Road, Mangilal Plot, Amravati, Maharashtra, 444602
GSTIN: 27AALCR5165R1ZF
YouTube (200k) Twitter (100k) LinkedIn (250k) GitHub (6k)
© Arpit Bhayani, 2025
