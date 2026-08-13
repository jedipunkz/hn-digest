---
source: "https://sdtimes.com/programming-languages/elixir-clojure-or-python-for-llm-agents-our-experience-with-all-three/"
hn_url: "https://news.ycombinator.com/item?id=49289358"
title: "Modeling One LLM Agent Three Ways: Python, Clojure, Elixir"
article_title: "Elixir, Clojure, or Python for LLM Agents? Our Experience with All Three - SD Times"
author: "ViktoriiaYarosh"
captured_at: "2026-08-13T17:50:07Z"
capture_tool: "hn-digest"
hn_id: 49289358
score: 1
comments: 0
posted_at: "2026-08-13T17:38:59Z"
tags:
  - hacker-news
  - translated
---

# Modeling One LLM Agent Three Ways: Python, Clojure, Elixir

- HN: [49289358](https://news.ycombinator.com/item?id=49289358)
- Source: [sdtimes.com](https://sdtimes.com/programming-languages/elixir-clojure-or-python-for-llm-agents-our-experience-with-all-three/)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T17:38:59Z

## Translation

タイトル: 1 つの LLM エージェントのモデリング 3 つの方法: Python、Clojure、Elixir
記事のタイトル: LLM エージェントには Elixir、Clojure、または Python? 3 つすべてについての私たちの経験 - SD Times
説明: ほとんどのエージェント ツールは Python ファーストです。 LangChain、AutoGen、CrewAI、および LangGraph はすべて Python をターゲットとしています。 Python が 2 番目に人気のあるプログラミング言語であることを考えると、現在のエコシステムは、すでにそれを使用しているチームにとってはうまく機能する可能性があります。それでも、JVM インフラストラクチャまたは Erlang/OTP システムを実行している組織
[切り捨てられた]

記事本文:
LLM エージェントには Elixir、Clojure、または Python? 3 つすべてについての私たちの経験 - SD Times
コンテンツにスキップ
SDタイムズ
データ品質スポンサー ショーケース: Melissa
LLM エージェントには Elixir、Clojure、または Python? 3 つすべてに関する私たちの経験
ほとんどのエージェント ツールは Python ファーストです。 LangChain、AutoGen、CrewAI、および LangGraph はすべて Python をターゲットとしています。 Python が 2 番目に人気のあるプログラミング言語であることを考えると、現在のエコシステムは、すでにそれを使用しているチームにとってはうまく機能する可能性があります。それでも、JVM インフラストラクチャまたは Erlang/OTP システムを実行している組織は、エージェントを Python に移行するか、すでに運用しているランタイムで構築するかという問題に直面しています。
関数型プログラミングのアンバサダーとして、私たちは選択した言語である Elixir と Clojure のエージェント システムを試してきました。この記事は、これまでの取り組みを部分的に要約したもので、それらを Python と比較し、それぞれが運用エージェント システムの特定の要件をどのように処理するかを検証します。
しかし、必要な人のためにトリビアから始めましょう。 LLM エージェントは、言語モデルと関数を呼び出す機能を組み合わせます。多くの場合 ReAct (Reasoning and Acting) と呼ばれるコア ループは次のように機能します。LLM は会話と利用可能なツールを調べ、ツールを呼び出すか応答するかを決定し、ツールを呼び出した場合は結果が会話にフィードバックされます。ループは、エージェントが最終応答を生成するか、ステップ制限に達するまで継続します。
Anthropic は、ワークフロー (事前定義されたコード パスを通じて調整される LLM) とエージェント (独自のプロセスとツールの使用を動的に指示する LLM) を区別します。どちらも同じ基本ループに従います。違いは、LLM がシーケンスをどの程度制御するかです。
言語によって異なるのは、ツール、状態、ループ自体をどのように表現するかです。
シンプルな分析エージェントを com として使用します。

パリソンポイント。たとえば、毎週のユーザーに関する統計を返すためにデータベースにクエリを実行し、必要に応じてグラフを生成します。
Python は、エージェントを起動するためのすぐに使えるフレームワークを提供します。 Python エージェントも最初から作成しますが、これらを省略することはできません。
「」パイソン
langchain_openai から ChatOpenAI をインポート
langchain.agents からインポートの初期化_agent、ツール
def run_sql(クエリ: str):
...
llm = ChatOpenAI(model="gpt-4.1-mini")
ツール = [
ツール(name="run_sql", func=run_sql,
description="分析データベースで SQL クエリを実行します。")
】
エージェント = 初期化_エージェント(
tools=ツール、llm=llm、
エージェント="ゼロショット反応説明"、冗長=True、
)
result = Agent.run("先週のアクティブ ユーザーの数は何人ですか?")
「」
エージェント ループは「initialize_agent」内で実行されます。状態とトレースには、フレームワーク API を通じてアクセスします。ツールは「Tool」クラスのインスタンスです。
「」パイソン
ツール = {
"run_sql": {"run": run_sql},
"render_chart": {"run": render_chart},
}
def run_agent(質問: str) -> dict:
状態 = {
"会話": [{"役割": "ユーザー", "コンテンツ": 質問}],
「トレース」: [],
}
決定 = call_llm(状態["会話"], ツール)
if Decision["type"] == "tool_call":
ツール名 = 決定["ツール"]
params = 決定["params"]
result = TOOLS[ツール名]["実行"](params)
state["会話"].append({
"役割": "ツール"、"名前": ツール名、
"content": repr({"params": params, "result": result}),
})
state["トレース"].append({
"ステップ": 1、"ツール": ツール名、
"params": パラメータ、"result": 結果
})
状態を返す
「」
ツールは辞書です。状態は辞書です。制御フローが表示されます。このバージョンは、以下の Clojure バージョンと同じ方法でテストできます。ここでのトレードオフは、Python の可変データ構造は、ツール関数が参照を通じて「状態」を変更でき、その変更が tr に表示されないことを意味します。

エース。そのような行動は言語の欠陥であると主張する人もいます。私たちはそれが管理すべき財産であると信じています。
Clojure は、エージェントを不変マップ上のデータ変換として表します。
「」クロージャー
(def run-sql-tool
{:名前「run_sql」
:description "分析データベースで SQL クエリを実行します"
:params [:map [:クエリ文字列?]]
:run (fn [{:keys [クエリ]}]
(db/run-sql クエリ))})
(デフォルトツール
{"run_sql" ラン SQL ツール
"render_chart" レンダーチャートツール})
「」
ツールは地図です。パラメーター スキーマは、クラスやデコレーターではなくデータ構造としてスキーマを定義する Malli を使用します。これは、スキーマをプログラムで生成、シリアル化、変換できることを意味します。これは、LLM API が期待する JSON 形式に変換するときに役立ちます。
「」クロージャー
(defn run-agent-once [状態設定]
(let [決定 (llm/call-llm-with-tools
(:モデル構成) (:API キー構成)
ツール/ツール (:会話状態))]
(case (:型決定)
:メッセージ
{:state (追加メッセージ状態「アシスタント」(:内容決定))
：終わった？本当}
:ツールコール
(let [{:keys [tool params]} 決定
tool-def (ツール/ツールツールの取得)
params' (tools/validate-params tool-def params)
結果 ((:run ツール定義) パラメータ')]
{:state (append-tool-result 状態ツール パラメーターの結果)
：終わった？ false}))))
(defn run-agent [ユーザー質問設定]
(loop [state (初期状態のユーザー質問)
ステップ0]
(let [{:keys [状態完了?]} (エージェントを 1 回実行する状態の設定)]
(if (または完了? (>= ステップ (:max-steps config 8)))
状態
(再発状態 (ステップを含む))))))
「」
各反復は状態を取得し、新しい状態を返します。古い状態は変わりません。これは、2 つの状態を比較して、特定の反復によって何が変更されたかを確認できることを意味します。完全な状態を EDN にシリアル化して保存し、後で実行を再実行できます。開発中に、REPL を使用すると、キャプチャされた状態で「run-agent-once」を呼び出し、手動でステップ実行できます。
「」c

ロジュール
(最も巧妙なエージェントがトレースを生成します)
(let [state (core/run-agent "アクティブ ユーザーは何人ですか?" config)]
(is (= 1 (カウント (:トレース状態))))
(is (= "run_sql" (-> state :trace first :tool))))
「」
関数を呼び出して、返されたマップに対してアサートします。スタブ LLM により、動作が決定的になります。モックするフレームワーク内部がないため、モック ライブラリは必要ありません。
Elixir は、アクター モデルを使用して各エージェントをプロセスとしてモデル化します。プロセスは軽量 (キロバイトのメモリ) で、メッセージ パッシングを通じて通信し、障害回復のために監視されます。
「」エリクサー
defmodule AnalyticsAgent が行うこと
GenServerを使用する
def start_link(opts) を実行します
GenServer.start_link(__MODULE__, opts)
終わり
def init(opts) を実行します
{:わかりました、%{
会話: []、
トレース: []、
ツール: %{
"run_sql" =>&Tools.run_sql/1,
"render_chart" =>&Tools.render_chart/1
}
}}
終わり
def handle_call({:ask, question}, _from, state) do
state = update_in(state.conversation, &[%{role: "user", content: question} | &1])
{結果、新しい状態} = 実行ループ(状態、最大ステップ数: 8)
{:返信、結果、新しい状態}
終わり
defp run_loop(state, opts) do
case LLM.call_with_tools(state.conversation, state.tools) do
{:メッセージ、コンテンツ} ->
{コンテンツ, append_message(状態, "アシスタント", コンテンツ)}
{:tool_call、ツール、パラメータ} ->
結果 = state.tools[ツール].(params)
new_state = append_tool_result(状態、ツール、パラメータ、結果)
run_loop(new_state, opts)
終わり
終わり
終わり
「」
メッセージパッシングモデルは、標準エージェントのワークフローパターンに直接マッピングされます。プロンプトチェーンは、メッセージを転送するプロセスです。ルーティングは、特化されたエージェント プロセスにディスパッチする分類子プロセスです。オーケストレーター プロセスはワーカー プロセスを生成し、管理します。複数のエージェント プロセスはデフォルトで同時に実行されます。これが Elixir のコア機能であるためです。
「」エリクサー
defmodule AgentSupervisor が実行する
スーパーバイザーを使用する
def 初期化(_

オプト) する
子供 = [
{Analytics エージェント、名前: :analytics}、
{CodeGenAgent、名前: :codegen}、
{レビューエージェント、名前: :review}
】
Supervisor.init(子供、戦略: :one_for_one)
終わり
終わり
「」
1 つのエージェント プロセスがクラッシュした場合（不正な LLM 出力、API タイムアウト、または不正なツール結果が原因）、スーパーバイザはそのプロセスを再起動します。他のエージェント プロセスは影響を受けません。このようにして、Erlang/OTP は 1980 年代以来プロセス障害を処理してきました。このアプローチは、変更せずに LLM エージェントに適用されます。
Python は「asyncio」、スレッド化、またはマルチプロセッシングを使用します。 GIL は、CPU に依存した並列処理を制限します。 I/O バウンドのエージェント作業 (ほとんどの LLM API 呼び出しがこれに該当します) では、「asyncio」 が適切に機能します。 CPU に制約のある作業や多数の同時エージェントの場合は、Ray や Celery などの外部ツールが一般的です。
Clojure には同時実行プリミティブ (atoms、refs、agent、core.async) があり、JVM スレッドで実行されます。複数のエージェントを同時に実行するには、これらのプリミティブを明示的に使用する必要がありますが、十分にサポートされています。
Elixir は、プリエンプティブ スケジューリングを使用して BEAM VM 上で軽量プロセスを実行します。 1 台のマシンで、すべての CPU コアに分散された数百万のプロセスを実行できます。エージェントを同時に実行するには、特別な設定は必要ありません。プロセスを開始するだけです。
Python の状態はデフォルトでは変更可能です。フレームワークベースのエージェントでは、通常、状態はクラス インスタンスの内部にあります。プレーンコード エージェントでは、状態は辞書内にあり、参照を使用してどこからでも変更できます。トレーサビリティは、ログ記録の規律に依存します。
Clojure の状態は不変です。エージェントの反復ごとに、前の状態マップを変更せずに新しい状態マップが生成されます。状態は、比較、シリアル化、保存、再生することができます。 REPL を使用すると、開発中に中間状態を直接検査できます。
Elixir プロセスには孤立した状態があり、各プロセスは他のプロセスとは異なる独自の状態を維持します。

直接アクセスすることはできません。これにより、エージェント間での偶発的な状態破損が防止されます。検査は `:sys.get_state/1` および `:observer` を通じて利用できますが、モデルはデータ中心ではなくプロセス中心です。
Python には Try/Except が用意されています。再試行ロジックとサーキット ブレーカーは、手動またはライブラリ経由で実装されます。エージェント フレームワークは、失敗の処理方法が異なります。再試行メカニズムを備えているものもあれば、開発者に任せているものもあります。
Clojure は JVM 例外処理を継承します。監視パターンはライブラリを使用して構築できますが、言語とランタイムはそれらをネイティブに提供しません。
Elixir には、コア ランタイム機能として監視ツリーがあります。スーパーバイザはプロセスを監視し、構成可能な戦略に従ってプロセスを再起動します。このアプローチは何十年にもわたって Erlang/OTP システムの標準であり、エージェント プロセスに直接適用されます。
Python では、マシン間でエージェントを配布するための外部インフラストラクチャ (Kubernetes、Celery、Ray) が必要です。調整プロトコルは別途追加する必要があります。
Clojure は JVM クラスタリング ソリューションを使用できます。 Agent-o-Rama ライブラリは、Rama 上での分散エージェントの実行を提供します。ディストリビューションは言語には組み込まれていませんが、JVM エコシステムを通じて利用できます。
Elixir は Erlang のクラスタリングを継承します。プロセス間のメッセージの受け渡しは、プロセスが同じマシン上にあっても、異なるマシン上にあっても、同じように機能します。エージェント コードを変更せずに、1 台のマシンで開発し、クラスターに拡張できます。
Python には最大の AI エコシステムがあります。すべての主要な LLM プロバイダーには Python SDK が同梱されています。エージェント フレームワーク、埋め込みライブラリ、ベクター ストアの統合、評価ツールはすべて Python ファーストです。特定の統合が必要な場合は、おそらく Python ですでに利用可能です。
Clojure には、AI 固有のライブラリのための小規模なエコシステムがあります。 OpenAI および Anthropic API クライアントが存在します。 T

JVM は Java ライブラリへのアクセスを提供します。多くの統合では、ラッパー コードを作成します。
Elixir には、数値計算用の Nx、モデル推論用の Bumblebee、構造化出力用の Instructor など、新たな AI エコシステムがあります。 LLM API 統合は存在しますが、Python ほど包括的ではありません。
Python のテストはアプローチによって異なります。プレーンコード エージェント (ツールは辞書として、状態は辞書として) は、他の Python コードと同じ方法でテストします。フレームワークベースのエージェントでは、多くの場合、テストをフレームワークの実装に結合するフレームワーク内部のモックが必要になります。
Clojure のテストは、データ指向の設計を直接受け継いでいます。関数を呼び出して、返されたマップを確認します。スタブ LLM をスワップインし、エージェントを実行し、トレースでアサートします。特別なテスト インフラストラクチャは必要ありません。
Elixir テストでは、プロセスベースの分離を備えた ExUnit を使用します。個々のエージェントのテストは簡単です。同時エージェント間の対話をテストするには、非同期メッセージ パッシングを処理するための追加のセットアップが必要です。
エージェントには、呼び出すことができる関数と操作するデータ型に関する構造化された情報が必要です。
Elixir はドキュメントを第一級の言語機能として扱います。 `@doc`、`@moduledoc`、および `@spec` アノテーションは標準ワークフローの一部です。これらは型信号を提供します

[切り捨てられた]

## Original Extract

Most agent tooling is Python-first. LangChain, AutoGen, CrewAI, and LangGraph all target Python. Given that Python is the second-most-popular programming language, the current ecosystem might work well for teams already using it. Still, organizations running JVM infrastructure or Erlang/OTP systems
[truncated]

Elixir, Clojure, or Python for LLM Agents? Our Experience with All Three - SD Times
Skip to content
SD Times
Data Quality Sponsor Showcase: Melissa
Elixir, Clojure, or Python for LLM Agents? Our Experience with All Three
Most agent tooling is Python-first. LangChain, AutoGen, CrewAI, and LangGraph all target Python. Given that Python is the second-most-popular programming language , the current ecosystem might work well for teams already using it. Still, organizations running JVM infrastructure or Erlang/OTP systems face the question of whether to move agents to Python or build them in the runtime they already operate.
As ambassadors of functional programming, we have been toying with the agentic systems in our languages of choice, Elixir and Clojure. This article, which partially summarizes our previous endeavors, compares them with Python and examines how each handles the specific requirements of production agent systems.
But let’s start with trivia for those who need it. An LLM agent combines a language model with the ability to call functions. The core loop — often called ReAct (Reasoning and Acting) — works like this: the LLM examines the conversation and available tools, decides whether to call a tool or respond, and if it calls a tool, the result gets fed back into the conversation. The loop continues until the agent produces a final answer or hits a step limit.
Anthropic distinguishes workflows (LLMs orchestrated through predefined code paths) and agents (LLMs that dynamically direct their own processes and tool usage). Both follow the same basic loop. The difference is in how much the LLM controls the sequencing.
What varies across languages is how you represent tools, state and the loop itself.
We’ll use a simple analytic agent as the comparison point. It will query the database to, let’s say, return statistics on weekly users, optionally generating charts if requested.
Python provides us with ready frameworks for spinning up agents. We cannot omit them, though we will also write Python agents from scratch.
```python
from langchain_openai import ChatOpenAI
from langchain.agents import initialize_agent, Tool
def run_sql(query: str):
...
llm = ChatOpenAI(model="gpt-4.1-mini")
tools = [
Tool(name="run_sql", func=run_sql,
description="Run an SQL query on the analytics db.")
]
agent = initialize_agent(
tools=tools, llm=llm,
agent="zero-shot-react-description", verbose=True,
)
result = agent.run("How many active users did we have last week?")
```
The agent loop runs inside `initialize_agent`. State and trace are accessed through framework APIs. Tools are `Tool` class instances.
```python
TOOLS = {
"run_sql": {"run": run_sql},
"render_chart": {"run": render_chart},
}
def run_agent(question: str) -> dict:
state = {
"conversation": [{"role": "user", "content": question}],
"trace": [],
}
decision = call_llm(state["conversation"], TOOLS)
if decision["type"] == "tool_call":
tool_name = decision["tool"]
params = decision["params"]
result = TOOLS[tool_name]["run"](params)
state["conversation"].append({
"role": "tool", "name": tool_name,
"content": repr({"params": params, "result": result}),
})
state["trace"].append({
"step": 1, "tool": tool_name,
"params": params, "result": result
})
return state
```
Tools are dictionaries. State is a dictionary. The control flow is visible. This version is testable in the same way as the Clojure version below. The trade-off here is that Python’s mutable data structures mean that a tool function can modify `state` through a reference without that modification showing up in the trace. Some would argue that such behaviour is a language flaw; we believe that it is a property to manage.
Clojure represents the agent as data transformations on immutable maps.
```clojure
(def run-sql-tool
{:name "run_sql"
:description "Run an SQL query on the analytics db"
:params [:map [:query string?]]
:run (fn [{:keys [query]}]
(db/run-sql query))})
(def tools
{"run_sql" run-sql-tool
"render_chart" render-chart-tool})
```
Tools are maps. Parameter schemas use Malli, which defines schemas as data structures rather than classes or decorators. It means schemas can be programmatically generated, serialized and transformed, which is useful when converting to the JSON format that LLM APIs expect.
```clojure
(defn run-agent-once [state config]
(let [decision (llm/call-llm-with-tools
(:model config) (:api-key config)
tools/tools (:conversation state))]
(case (:type decision)
:message
{:state (append-message state "assistant" (:content decision))
:done? true}
:tool-call
(let [{:keys [tool params]} decision
tool-def (get tools/tools tool)
params' (tools/validate-params tool-def params)
result ((:run tool-def) params')]
{:state (append-tool-result state tool params' result)
:done? false}))))
(defn run-agent [user-question config]
(loop [state (initial-state user-question)
steps 0]
(let [{:keys [state done?]} (run-agent-once state config)]
(if (or done? (>= steps (:max-steps config 8)))
state
(recur state (inc steps))))))
```
Each iteration takes a state and returns a new state. The old state is unchanged. It means you can diff two states to see what a specific iteration changed. You can serialize the full state to EDN, save it and replay execution later. During development, the REPL lets you call `run-agent-once` with a captured state and step through execution manually.
```clojure
(deftest agent-produces-trace
(let [state (core/run-agent "How many active users?" config)]
(is (= 1 (count (:trace state))))
(is (= "run_sql" (-> state :trace first :tool)))))
```
You call a function and assert on the returned map. The stub LLM makes behavior deterministic. No mocking libraries are needed because there are no framework internals to mock.
Elixir models each agent as a process using the Actor Model. Processes are lightweight (kilobytes of memory), communicate through message passing, and are supervised for fault recovery.
```elixir
defmodule AnalyticsAgent do
use GenServer
def start_link(opts) do
GenServer.start_link(__MODULE__, opts)
end
def init(opts) do
{:ok, %{
conversation: [],
trace: [],
tools: %{
"run_sql" => &Tools.run_sql/1,
"render_chart" => &Tools.render_chart/1
}
}}
end
def handle_call({:ask, question}, _from, state) do
state = update_in(state.conversation, &[%{role: "user", content: question} | &1])
{result, new_state} = run_loop(state, max_steps: 8)
{:reply, result, new_state}
end
defp run_loop(state, opts) do
case LLM.call_with_tools(state.conversation, state.tools) do
{:message, content} ->
{content, append_message(state, "assistant", content)}
{:tool_call, tool, params} ->
result = state.tools[tool].(params)
new_state = append_tool_result(state, tool, params, result)
run_loop(new_state, opts)
end
end
end
```
The message-passing model maps directly to standard agent workflow patterns. Prompt chaining is processes passing messages forward. Routing is a classifier process dispatching to specialized agent processes. An orchestrator process spawns and manages worker processes. Multiple agent processes run concurrently by default because that’s the core Elixir’s offer.
```elixir
defmodule AgentSupervisor do
use Supervisor
def init(_opts) do
children = [
{AnalyticsAgent, name: :analytics},
{CodeGenAgent, name: :codegen},
{ReviewAgent, name: :review}
]
Supervisor.init(children, strategy: :one_for_one)
end
end
```
If one agent process crashes (due to bad LLM output, API timeout, or malformed tool result), the supervisor restarts it. The other agent processes are unaffected. In this way, Erlang/OTP has handled process failures since the 1980s; this approach applies to LLM agents without modification.
Python uses `asyncio`, threading, or multiprocessing. The GIL limits CPU-bound parallelism. For I/O-bound agent work (which most LLM API calls are), `asyncio` works adequately. For CPU-bound work or large numbers of concurrent agents, external tools like Ray or Celery are common.
Clojure has concurrency primitives (atoms, refs, agents, core.async) and runs on JVM threads. Running multiple agents concurrently requires explicit use of these primitives but is well-supported.
Elixir runs lightweight processes on the BEAM VM with preemptive scheduling. A single machine can run millions of processes distributed across all CPU cores. Running agents concurrently requires no special setup; you just start processes.
Python state is mutable by default. In framework-based agents, state is typically internal to class instances. In plain-code agents, the state is in dictionaries that can be mutated from anywhere with a reference. Traceability depends on logging discipline.
Clojure state is immutable. Each agent iteration produces a new state map without modifying the previous one. States can be diffed, serialized, stored, and replayed. The REPL allows direct inspection of any intermediate state during development.
Elixir processes have an isolated state — each process maintains its own state that other processes cannot directly access. It prevents accidental state corruption across agents. Inspection is available through `:sys.get_state/1` and `:observer`, but the model is process-centric rather than data-centric.
Python provides try/except. Retry logic and circuit breakers are implemented manually or via libraries. Agent frameworks vary in how they handle failures — some have retry mechanisms, others leave it to the developer.
Clojure inherits JVM exception handling. Supervision patterns can be built using libraries, but the language and runtime don’t provide them natively.
Elixir has supervision trees as a core runtime feature. Supervisors monitor processes and restart them according to configurable strategies. This approach has been the standard in Erlang/OTP systems for decades and applies directly to agent processes.
Python requires external infrastructure (Kubernetes, Celery, Ray) for distributing agents across machines. Coordination protocols must be added separately.
Clojure can use JVM clustering solutions. The Agent-o-Rama library provides distributed agent execution on Rama. Distribution isn’t built into the language but is available through the JVM ecosystem.
Elixir inherits Erlang’s clustering. Message passing between processes works the same way whether processes are on the same machine or different machines. You can develop on one machine and scale to a cluster without changing the agent code.
Python has the largest AI ecosystem. Every major LLM provider ships a Python SDK. Agent frameworks, embedding libraries, vector store integrations, and evaluation tools are all Python-first. If you need a specific integration, it’s probably already available in Python.
Clojure has a smaller ecosystem for AI-specific libraries. OpenAI and Anthropic API clients exist. The JVM gives access to Java libraries. For many integrations, you’ll write wrapper code.
Elixir has an emerging AI ecosystem—Nx for numerical computing, Bumblebee for model inference, Instructor for structured outputs. LLM API integrations exist but are less comprehensive than Python’s.
Python testing depends on the approach. Plain-code agents (tools as dictionaries, state as dictionaries) test the same way as any other Python code. Framework-based agents often require mocking framework internals, which couples tests to the framework’s implementation.
Clojure testing follows directly from the data-oriented design. Call the function and check the returned map. Swap in a stub LLM, run the agent, assert on the trace, and no special test infrastructure.
Elixir testing uses ExUnit with process-based isolation. Testing individual agents is straightforward. Testing interactions between concurrent agents requires more setup to handle asynchronous message passing.
Agents need structured information about the functions they can call and the data types they work with.
Elixir treats documentation as a first-class language feature. `@doc`, `@moduledoc`, and `@spec` annotations are part of the standard workflow. These provide type signa

[truncated]
