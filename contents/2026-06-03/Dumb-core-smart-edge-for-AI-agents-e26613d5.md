---
source: "https://arizenai.com/dumb-core-smart-edge/"
hn_url: "https://news.ycombinator.com/item?id=48369293"
title: "Dumb core, smart edge for AI agents"
article_title: "Dumb Core, Smart Edge: Agentic Design"
author: "arizen"
captured_at: "2026-06-03T00:48:15Z"
capture_tool: "hn-digest"
hn_id: 48369293
score: 2
comments: 0
posted_at: "2026-06-02T12:22:34Z"
tags:
  - hacker-news
  - translated
---

# Dumb core, smart edge for AI agents

- HN: [48369293](https://news.ycombinator.com/item?id=48369293)
- Source: [arizenai.com](https://arizenai.com/dumb-core-smart-edge/)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T12:22:34Z

## Translation

タイトル: ダムコア、AI エージェント向けのスマートエッジ
記事のタイトル: ダムコア、スマートエッジ: エージェントティックデザイン
説明: AI エージェント オーケストレーターを小型、型指定、テスト可能、置換可能に保ちながら、ドメイン インテリジェンスを専門エッジに移行するための実用的なアーキテクチャ パターン。

記事本文:
サインイン
購読する
イーゴリ・ボブリアコフ
で
原則
—
2026 年 3 月 10 日
ダムコア、スマートエッジ: エージェントティックデザイン
私が実稼働で失敗するのを見てきたエージェント システムの多くは同じ形状でした。つまり、インテリジェンスが中央に集中しており、テスト、置き換え、推論が困難でした。オーケストレーターは、リクエストのルーティング、ドメインの知識の保持、メモリの管理、ツールの選択、出力の形成など、多くのことを一度に実行しすぎていました。何かが壊れたとき、それを切り分けることはできません。要件が変更された場合、手動で更新することはできません。全体が動かなければなりませんでした。
これはモデルの品質の問題ではありません。これらのシステムを構築しているチームは、劣った LLM を使用していませんでした。彼らは、明確な反原理を伴って、名前を挙げるのに十分なほど繰り返し発生するアーキテクチャ上のエラーに遭遇していました。
原理は次のとおりです: ダムコア、スマートエッジ。オーケストレーター —
シーケンスが機能する中央ノード - ほぼステートレスで、制御フローのみをエンコードする必要があります。
インテリジェンスは周辺部、つまり独自のドメインを所有する専門家ノードに属します。
境界を考慮し、単独で推論、テスト、置き換えることができます。これはそうではありません
単なる文体の好みです。システムに複数のネットワークがある場合に便利なトポロジです。
独自に進化する必要があるドメイン、ツール、または責任。
ルーティング、ドメイン ロジック、メモリ、ツールの選択を 1 つのプロンプトで処理する「スマート コア」オーケストレーターは、運用エージェント システムでよく見られるアーキテクチャ障害です。
ダム コア、スマート エッジの原則: オーケストレーターは状態遷移をエンコードします。スペシャリストは、明示的な境界内でドメイン インテリジェンスを所有します。
経済的な利点は、より狭いコンテキストから生まれます。専門家は、完全なシステム プロンプトやすべてのツール スキーマの代わりに、自分のタスクに必要なスライスを受け取ります。
交換可能性テスト: 交換できるはずです。

オーケストレーターや隣接ノードを変更することなく、専門家による実装が可能です。
すべてのスペシャリスト エッジは、型付きの Pydantic コントラクトを発行する必要があります。オーケストレーターは、生のテキストではなく、型を消費します。
LangGraph の条件付きエッジは、このパターンに直接マッピングされます。型付きの状態はルーティング メタデータを保持し、ノードはインテリジェンスを保持します。
レイテンシは実際のトレードオフです。テスト容易性、置換可能性、並列実行の向上と比較して評価してください。
なぜ中核となるインテリジェンスが機能しないのか
知性を一元化したいという本能は理解できます。強力なモデルが 1 つあります。システムプロンプトを 1 つ作成します。すべてのツールを与えて、物事を理解させます。デモではこれで動作します。このモデルは懸念事項を同時にやりくりするのに十分な能力を備えており、幸せな道筋はきれいに見えます。
プロダクションでは、4 つの異なる方法でこれを打破します。
1つ目：絡み合い。オーケストレーターがルーティング ロジックとドメイン ナレッジの両方をエンコードする場合、一方を変更するにはもう一方に触れる必要があります。新しいツールはプロンプトを書き直すことを意味します。新しいドメイン ルールは、ルーティングを再テストすることを意味します。このシステムは、アーキテクチャ上の境界のない、長くて脆弱な自然言語命令の文字列だけの結合を開発します。
2 番目: テスト不可能性と不透明性。一度に 5 つのことを実行する LLM プロンプトの単体テストを作成することはできません。エンドツーエンドの評価を実行して、新たな動作が安定しているかどうかを観察することしかできません。エージェントが不正な動作をした場合、その障害がルーティング、ドメイン推論、ツール選択、または出力フォーマットのいずれにあったのかを知る必要があります。スマート コアはこれらすべてを統合します。本番環境まで回帰は目に見えませんが、本番環境では懐中電灯を使ってブラック ボックスをデバッグすることになります。
3 番目: コストの増大。スマートコア エージェントは、コンテキストのすべてのトークン (完全な履歴、すべてのツール スキーマ、すべてのドメイン ルール) を同じパスを通じて常に送信します。

ライステップ。複数ステップのワークフローでは、これにより、現在のステップで必要のないコンテキストが繰り返し発生します。ダムコアは専門家にルートします。専門家は必要なものだけを受け取ります。モデル ルーティングの経済規模 (数百万コール) では、経済は急速に悪化します。
構造的周辺の原理: エージェント システムでは
複数のドメインまたはツールの境界、
自律的な判断能力は最大限に配分されるべきである。
グラフのエッジを処理し、コアは状態遷移のみを担当します。
ルーティング契約。
これは新しいアイデアではありません。エージェント グラフ内に現れる古いシステム エンジニアリングです。 Unix パイプが機能するのは、各プログラムが 1 つのことを実行し、シンプルで安定したインターフェイスを通じて通信するためです。マイクロサービスは、機能するときは同じ理由で機能します。インターネットのコア プロトコルは意図的にシンプルになっています。インテリジェンスはエンドポイントに存在します。 Dumb Core、Smart Edge は、これと同じロジックをエージェント グラフに適用します。
オーケストレーターのジョブの範囲は正確に定められています。つまり、型指定された入力を受け取り、どのスペシャリストがそれを処理するかを決定し、ディスパッチし、結果を待ち、型指定された出力を出力します。ドメイン知識はエンコードされません。永続的なメモリは保持しません。内容については判断を下しません。これはステート マシンであり、ステート マシンとして読み取れる必要があります。
対照的に、スペシャリストはドメイン境界内で自律的です。 ResearchSpecialist は、情報源を照会し、信頼性を評価し、結果を総合する方法を知っています。 CodeReviewSpecialist は、コードベースの規則、チェックすべき障害モード、およびその出力のフォーマット方法を知っています。それぞれを完全に独立してプロンプト、微調整、交換、または評価することができます。オーケストレーターは何が変更されたかを知る必要はありません。
「ダムコア」は正確な用語であり、軽蔑語ではありません。オーケストレーターは引き続き CLAS に LLM を使用できます。

SIFY_INTENT ステップ — 受信リクエストを型指定されたインテントに分類することは、言語理解の正当な使用法です。してはいけないのは、そのインテントの内容について推論したり、ドメイン ヒューリスティックを適用したり、ドメイン内のエッジ ケースの処理方法について判断を下したりすることです。
有用なテスト: オーケストレーターのシステム プロンプトからドメイン固有の語彙をすべて削除し、汎用のプレースホルダーに置き換えた場合、ルーティング ロジックは引き続き機能しますか? 「はい」の場合、コアは適切にダムです。ルーティングが「請求に異議のある返金リクエスト」の意味を理解することに依存している場合、ドメインの知識がコアに漏洩しており、結合の問題が発生しています。
オーケストレーターのシステム プロンプトは、交通管制官のマニュアルのように読む必要があります。つまり、貨物についての意見はなく、フロー、優先順位、障害の処理に関するルールです。専門家のプロンプトは専門家向けのブリーフのように、深く、意見がまとまっており、範囲が狭いものである必要があります。
これはメモリ アーキテクチャにも影響します。オーケストレーターがアクセスできる共有グローバル メモリは臭いです。各専門家は、そのタスクの期間中、自分の作業記憶を所有する必要があります。永続的な記憶 (ユーザーの好み、以前の会話の概要、学んだ事実) は、オーケストレーターのコンテキストに事前にロードされるのではなく、専門家がオンデマンドで取得する必要があります。オーケストレーターは、メモリ ダンプではなく、セッション ID を渡します。
Dumb Core、Smart Edge の実​​際的な証明は、私が置き換え可能性テストと呼んでいるものです。オーケストレーターや隣接するスペシャリストを変更することなく、スペシャリスト ノードを交換できる必要があります。つまり、プロンプトベースの実装を決定論的アルゴリズム、より小さいモデル、または別のスペシャリスト実装に置き換えることができます。爆発範囲は専門家の境界に限定されるべきです。
スワップでノード外の変更が必要な場合は、

置き換えられ、システムが隠蔽されました
カップリング。最も一般的なソース: オーケストレーターが解析しているか、
入力された契約書を消費するのではなく、スペシャリストの出力の内部構造を利用します。
修正方法は常に同じです。エッジで Pydantic スキーマを定義します。
発行時に検証し、オーケストレーターに生のテキストではなく型を使用させます。
ここでコーザルエンジニアリングが登場します。ノードを交換し、システムの動作に対する個別の影響を観察できれば、システムの因果関係を把握できたことになります。同じオーケストレーター、同じ隣接するスペシャリスト、異なるスペシャリスト B など、制御された実験を実行できます。出力品質の差はスペシャリスト B のみに起因します。これは、直感や祈りではなく科学的にエージェント システムを改善する方法です。
LangGraph の実装パターンは直接的です。オーケストレーターを、型付き状態を持つグラフ (ルーティング メタデータ、セッション ID、結果スロットのみを運ぶ TypedDict または Pydantic モデル) として定義します。グラフ内の各ノードは、範囲が狭い入力を受け取り、型指定された出力を返す専門関数です。エッジは制御フローをエンコードします。ノードはインテリジェンスをエンコードします。
以下は、LangGraph v0.4 の最小限のダムコア オーケストレーターです。
import リテラルの入力から
pydanticインポートBaseModelから
langgraph.graph から StateGraph をインポート、END
# --- 型付き状態: ルーティング メタデータのみ、ドメイン知識なし ---
クラス OrchestratorState(BaseModel):
ユーザー入力: str
インテント: str = ""
スペシャリスト結果: str = ""
セッションID: str = ""
# --- 意図分類子 (コアが許可する 1 つの LLM 呼び出し) ---
async def destroy_intent(state: OrchestratorState) -> dict:
# 高速分類子モデルまたは決定論的分類子
意図 = await llm_classify(state.user_input) # 戻り値: "リサーチ" | 「タラ

e_review" | "不明"
return {"インテント": インテント}
# --- ルーティングエッジ: 型付き状態に対する純粋な関数 ---
def Route_by_intent(state: OrchestratorState) -> リテラル["research", "code_review", "fallback"]:
マッピング = {"リサーチ": "リサーチ", "コードレビュー": "コードレビュー"}
戻り値のマッピング.get(state.intent, "フォールバック")
# --- スペシャリスト ノード (それぞれがドメイン境界を所有) ---
async def Research_specialist(state: OrchestratorState) -> dict:
result = await run_research_agent(state.user_input, state.session_id)
return {"specialist_result": result.model_dump_json()}
async def code_review_specialist(state: OrchestratorState) -> dict:
result = await run_code_review_agent(state.user_input, state.session_id)
return {"specialist_result": result.model_dump_json()}
# --- グラフアセンブリ ---
グラフ = StateGraph(OrchestratorState)
graph.add_node("分類"、classify_intent)
graph.add_node("リサーチ", Research_specialist)
graph.add_node("コードレビュー", コードレビュー_スペシャリスト)
graph.add_node("fallback", lambda s: {"specialist_result": "それについてはお手伝いできません。"})
graph.set_entry_point("分類")
graph.add_conditional_edges("分類"、route_by_intent)
graph.add_edge("調査", END)
graph.add_edge("コードレビュー", END)
graph.add_edge("フォールバック", END)
app = chart.compile() オーケストレーターに含まれていないものに注目してください。ドメイン語彙、ツール スキーマ、メモリ取得、出力フォーマット ロジックなどはありません。 Route_by_intent 関数は純粋なマッピングです。を追加した場合
[切り捨てられた]
各スペシャリストは、型付き出力スキーマを使用して契約を強制します。
pydantic import BaseModel、Field から
クラス SpecialistResult(BaseModel):
"""スペシャリストとオーケストレーターとの間の契約。"""
答え: str = フィールド(..., min_length=1)
信頼度: float = フィールド(..., ge=0.0, le=1.0)
ソース: list[str] = フィールド(default_factory

=リスト)
ニーズエスカレーション: bool = False
クラス ResearchResult(SpecialistResult):
"""リサーチスペシャリストの契約を延長しました。"""
key_findings: list[str] = フィールド(..., min_length=1)
矛盾: list[str] = フィールド(default_factory=list)
# スペシャリストの内部:
async def run_research_agent(クエリ: str, session_id: str) -> ResearchResult:
raw = await llm_research(query, session_id)
return ResearchResult.model_validate(raw) # Pydantic はエミッションを検証します
プロのヒント: 導入する前に必ず語彙除去テストを実行してください。
オーケストレーターのシステム プロンプトからドメイン固有の単語をすべて削除し、置き換えます。
プレースホルダートークンを使用します。評価セット上でルーティングがまだ正しく機能する場合、コアは
きれいな。壊れた場合は、ドメインの知識が漏洩したことになります。適切なデータにリファクタリングされます。
出荷前に専門家が対応します。プロンプトの変更ごとに CI でこのテストを実行します。
CLASSIFY_INTENT ノードは生のユーザー入力を受け取り、型指定された IntentSchema を返します。それ以上のものはありません。質問への答えは始まりません。
DISPATCH エッジはインテント タイプを読み取り、適切なスペシャリストにルーティングします。これは LangGraph の条件付きエッジであり、型指定された状態に対する純粋な関数です。
各スペシャリスト ノードは、元の入力、そのドメイン コンテキスト、および ses など、必要なものだけを受け取ります。

[切り捨てられた]

## Original Extract

A practical architecture pattern for keeping AI agent orchestrators small, typed, testable, and replaceable while moving domain intelligence to specialist edges.

Sign in
Subscribe
By Igor Bobriakov
in
principle
—
10 Mar 2026
Dumb Core, Smart Edge: Agentic Design
Many agentic systems I've watched fail in production had the same shape: intelligence concentrated at the center, where it was hard to test, replace, or reason about. The orchestrator was doing too much — routing requests, holding domain knowledge, managing memory, selecting tools, and shaping outputs, all at once. When something broke, you couldn't isolate it. When requirements changed, you couldn't surgically update it. The whole thing had to move.
This is not a model quality problem. The teams building these systems weren't using inferior LLMs. They were running into an architectural error recurring enough to be useful to name, with a clear counter-principle.
The principle is this: Dumb Core, Smart Edge. The orchestrator — the
central node that sequences work — should be nearly stateless, encoding only control flow.
The intelligence belongs at the periphery, in specialist nodes that own their domain
boundary and can be reasoned about, tested, and replaced in isolation. This is not
just a stylistic preference. It is a useful topology when the system has multiple
domains, tools, or responsibilities that need to evolve independently.
A "smart core" orchestrator that handles routing, domain logic, memory, and tool selection in one prompt is a common architecture failure in production agentic systems.
The Dumb Core, Smart Edge principle: orchestrators encode state transitions; specialists own domain intelligence within explicit boundaries.
The economic advantage comes from narrower context: specialists receive the slice needed for their task instead of the full system prompt and every tool schema.
The replaceability test: you should be able to swap any specialist implementation without modifying the orchestrator or adjacent nodes.
Every specialist edge should emit a typed Pydantic contract — the orchestrator consumes the type, never raw text.
LangGraph conditional edges map directly to this pattern: typed state carries routing metadata, nodes carry intelligence.
Latency is a real tradeoff. Measure it against the gains in testability, replaceability, and parallel execution.
Why Intelligence at the Core Fails
The instinct to centralize intelligence is understandable. You have one powerful model. You write one system prompt. You give it all the tools and let it figure things out. In a demo, this works. The model is capable enough to juggle concerns simultaneously, and the happy path looks clean.
Production breaks this in four distinct ways.
First: entanglement. When the orchestrator encodes both routing logic and domain knowledge, changing one requires touching the other. A new tool means rewriting the prompt. A new domain rule means re-testing the routing. The system develops coupling that has no architectural boundary — just a long, fragile string of natural language instructions.
Second: untestability and opacity. You cannot write a unit test for an LLM prompt that does five things at once. You can only run end-to-end evaluations and observe whether the emergent behavior is stable. When an agent misbehaves, you need to know whether the failure was in routing, domain reasoning, tool selection, or output formatting. A smart core conflates all of these — regressions are invisible until production, and then you're debugging a black box with a flashlight.
Third: cost amplification. A smart-core agent sends every token of context — full history, all tool schemas, all domain rules — through the same path, for every step. In a multi-step workflow, this repeatedly pays for context the current step may not need. A dumb core routes to specialists. The specialists receive only what they need. At model-routing economics scale — millions of calls — the economics compound fast.
The Principle of Structural Periphery: In agent systems with
multiple domains or tool boundaries,
the capacity for autonomous judgment should be maximally distributed toward the
edges of the graph, leaving the core responsible only for state transitions and
routing contracts.
This is not a new idea; it is old systems engineering showing up inside agent graphs. Unix pipes work because each program does one thing and communicates through a simple, stable interface. Microservices work — when they work — for the same reason. The internet's core protocols are deliberately simple; the intelligence lives in the endpoints. Dumb Core, Smart Edge applies this same logic to the agent graph.
The orchestrator's job is precisely scoped: receive a typed input, determine which specialist handles it, dispatch, await a result, and emit a typed output. It encodes no domain knowledge. It holds no persistent memory. It makes no judgment calls about content. It is a state machine — and it should be readable as one.
The specialists, by contrast, are autonomous within their domain boundary. A ResearchSpecialist knows how to query sources, evaluate credibility, and synthesize findings. A CodeReviewSpecialist knows the codebase conventions, the failure modes to check, and how to format its output. Each one can be prompted, fine-tuned, swapped, or evaluated entirely independently. The orchestrator never needs to know what changed.
"Dumb core" is a precise term, not a pejorative. The orchestrator can still use an LLM for the CLASSIFY_INTENT step — classifying an incoming request into a typed intent is a legitimate use of language understanding. What it must not do is reason about the content of that intent, apply domain heuristics, or make judgment calls about how to handle edge cases within a domain.
A useful test: if you removed all domain-specific vocabulary from the orchestrator's system prompt and replaced it with a generic placeholder, would the routing logic still work? If yes, the core is appropriately dumb. If the routing depends on understanding what "a refund request with a disputed charge" means, domain knowledge has leaked into the core and you have a coupling problem.
The orchestrator's system prompt should read like a traffic controller's manual — rules about flow, priority, and failure handling, with no opinion about the cargo. The specialists' prompts should read like expert briefs — deep, opinionated, and narrow.
This also governs memory architecture. Shared, global memory accessible to the orchestrator is a smell. Each specialist should own its working memory for the duration of its task. Persistent memory — user preferences, prior conversation summaries, learned facts — should be retrieved by specialists on demand, not pre-loaded into the orchestrator's context. The orchestrator passes a session identifier, not a memory dump.
The practical proof of Dumb Core, Smart Edge is what I call the replaceability test: you should be able to swap any specialist node — replacing a prompt-based implementation with a deterministic algorithm, a smaller model, or a different specialist implementation — without modifying the orchestrator or any adjacent specialist. The blast radius should stay localized to that specialist boundary.
If a swap requires changes outside the node being replaced, the system has hidden
coupling. The most common source: the orchestrator is parsing or depending on the
internal structure of a specialist's output, rather than consuming a typed contract.
The fix is always the same — define a Pydantic schema at the edge,
validate on emission, and let the orchestrator consume the type, never the raw text.
This is where Causal Engineering enters the picture. When you can replace a node and observe the isolated effect on system behavior, you have a causal handle on the system. You can run controlled experiments: same orchestrator, same adjacent specialists, different specialist B. The delta in output quality is attributable to specialist B alone. This is how you improve an agentic system scientifically rather than through intuition and prayer.
The implementation pattern in LangGraph is direct. Define your orchestrator as a graph with typed state — a TypedDict or Pydantic model that carries only routing metadata, session identifiers, and result slots. Each node in the graph is a specialist function that receives a narrowly scoped input and returns a typed output. The edges encode the control flow. The nodes encode the intelligence.
Here is a minimal dumb-core orchestrator in LangGraph v0.4 :
from typing import Literal
from pydantic import BaseModel
from langgraph.graph import StateGraph, END
# --- Typed state: routing metadata only, no domain knowledge ---
class OrchestratorState(BaseModel):
user_input: str
intent: str = ""
specialist_result: str = ""
session_id: str = ""
# --- Intent classifier (the one LLM call the core is allowed) ---
async def classify_intent(state: OrchestratorState) -> dict:
# Fast classifier model or deterministic classifier
intent = await llm_classify(state.user_input) # returns: "research" | "code_review" | "unknown"
return {"intent": intent}
# --- Routing edge: pure function over typed state ---
def route_by_intent(state: OrchestratorState) -> Literal["research", "code_review", "fallback"]:
mapping = {"research": "research", "code_review": "code_review"}
return mapping.get(state.intent, "fallback")
# --- Specialist nodes (each owns its domain boundary) ---
async def research_specialist(state: OrchestratorState) -> dict:
result = await run_research_agent(state.user_input, state.session_id)
return {"specialist_result": result.model_dump_json()}
async def code_review_specialist(state: OrchestratorState) -> dict:
result = await run_code_review_agent(state.user_input, state.session_id)
return {"specialist_result": result.model_dump_json()}
# --- Graph assembly ---
graph = StateGraph(OrchestratorState)
graph.add_node("classify", classify_intent)
graph.add_node("research", research_specialist)
graph.add_node("code_review", code_review_specialist)
graph.add_node("fallback", lambda s: {"specialist_result": "I can't help with that."})
graph.set_entry_point("classify")
graph.add_conditional_edges("classify", route_by_intent)
graph.add_edge("research", END)
graph.add_edge("code_review", END)
graph.add_edge("fallback", END)
app = graph.compile() Notice what the orchestrator does not contain: no domain vocabulary, no tool schemas, no memory retrieval, no output formatting logic. The route_by_intent function is a pure mapping. If you added a
[truncated]
Each specialist enforces its contract with a typed output schema:
from pydantic import BaseModel, Field
class SpecialistResult(BaseModel):
"""Contract between any specialist and the orchestrator."""
answer: str = Field(..., min_length=1)
confidence: float = Field(..., ge=0.0, le=1.0)
sources: list[str] = Field(default_factory=list)
needs_escalation: bool = False
class ResearchResult(SpecialistResult):
"""Extended contract for research specialist."""
key_findings: list[str] = Field(..., min_length=1)
contradictions: list[str] = Field(default_factory=list)
# Inside the specialist:
async def run_research_agent(query: str, session_id: str) -> ResearchResult:
raw = await llm_research(query, session_id)
return ResearchResult.model_validate(raw) # Pydantic validates on emission
Pro Tip: Use the vocabulary-removal test before every deployment.
Strip all domain-specific words from your orchestrator's system prompt and replace them
with placeholder tokens. If the routing still works correctly on your eval set, the core is
clean. If it breaks, domain knowledge has leaked in — refactor it into the appropriate
specialist before shipping. We run this test in CI on every prompt change.
The CLASSIFY_INTENT node receives raw user input and returns a typed IntentSchema — nothing more. It does not begin answering the question.
The DISPATCH edge reads the intent type and routes to the correct specialist. This is a conditional edge in LangGraph — a pure function over the typed state.
Each specialist node receives only what it needs: the original input, its domain context, and a ses

[truncated]
