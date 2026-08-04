---
source: "https://www.union.ai/blog-post/flyte-2-is-generally-available-the-durable-open-source-ai-runtime"
hn_url: "https://news.ycombinator.com/item?id=49170846"
title: "Flyte 2 is GA: durable distributed AI workflows using regular Python"
article_title: "Flyte 2 Is Generally Available: The Durable, Open-Source AI Runtime | Union.ai"
author: "kumare3"
captured_at: "2026-08-04T16:07:20Z"
capture_tool: "hn-digest"
hn_id: 49170846
score: 1
comments: 0
posted_at: "2026-08-04T16:05:52Z"
tags:
  - hacker-news
  - translated
---

# Flyte 2 is GA: durable distributed AI workflows using regular Python

- HN: [49170846](https://news.ycombinator.com/item?id=49170846)
- Source: [www.union.ai](https://www.union.ai/blog-post/flyte-2-is-generally-available-the-durable-open-source-ai-runtime)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T16:05:52Z

## Translation

タイトル: Flyte 2 が GA です: 通常の Python を使用した耐久性のある分散 AI ワークフロー
記事のタイトル: Flyte 2 が一般提供開始: 耐久性のあるオープンソース AI ランタイム |ユニオンアイ
説明: ロジックとインフラ全体にわたる耐久性。再現可能、タイプセーフ、生産規模。

記事本文:
Flyte 2 が一般提供開始: 耐久性のあるオープンソース AI ランタイム | Flyte 2 が一般提供開始ユニオンアイ
アナリストによると、Union.ai は 9.8 倍の ROI を達成しています。
ユニオンアイ
ユニオンアイ
プラットフォーム
プラットフォームの概要 AI をこれまでより迅速に構築、出荷、拡張できます。
オープンソース AI ランタイムと本番 AI ランタイムを比較します。
Union.ai は、データが独自のセキュリティ境界を離れることがないように構築されています。
ケーススタディ 実際のチームによって報告された実際の結果
Union.ai チームの最新情報を読む
バーチャル ワークショップやトレーニングなどにご参加ください
サポート階層の詳細を確認する
検証済みの Union.ai 統合を参照する
Flyte 2 が一般提供開始: 耐久性のあるオープンソース AI ランタイム
例 H3 例 H4 例 H5 例 H6
https://www.union.ai/blog-post/flyte-2-is-generally-available-the-durable-open-source-ai-runtime
現在、Flyte 2 は一般提供されています。より高速で、よりスケーラブルで、(個人的には保証しますが) 使っていて本当に楽しいです。これはオープンソースのままであり、Linux Foundation の一部です。その内容を説明する前に、Flyte 2 は増分リリースではないため、なぜこれが存在するのかを説明したいと思います。これは、私たちが 10 年近く改良してきたシステムを根本から再構築したもので、Flyte が生まれた世界とは似ても似つかない世界に向けて再構築されました。
Flyte のこの新たな進化により、Flyte は AI および ML チームにとってオーケストレーターよりも価値のあるもの、つまり耐久性のある AI ランタイムに変わりました。オーケストレーターがシステムの上に置いて実行したり忘れたりするのに対し、ランタイムはコードそのものだけでなく、コードの下にあるインフラストラクチャを所有します。耐久性のある実行システムがロジック障害から回復できる場合、耐久性のあるランタイムは、インフラストラクチャの所有権を使用してコンピューティング障害 (OOM など) から回復することもでき、エンドツーエンドのランタイムの耐久性を実現します。これが、Flyte 2 を耐久性のある AI ランタイムと呼ぶときの意味です。

オーケストレーターであり、以下のすべてを実行するスレッドです。
FlyteはLyftから始まりました。私たちは、データの準備、モデルのトレーニング、大規模なシミュレーションの実行、モデルのバックテストを行うためにこれを構築しました。これは、ETA や価格設定などの製品の背後にある地味で重要な機構です。そのとき私たちが認識したことは今でも当てはまります。ML 製品の構築には、ユーザー向けソフトウェアの信頼性、各ワークロードに合わせて再構築するインフラストラクチャのダイナミズム、そして継続的に実験する機敏性が必要です。当時は 3 つすべてを提供するシステムがなかったため、1 つを構築しました。
私たちがそれを構築した世界を考えてみましょう。Python は単一スレッドのスクリプト言語として広く無視され、タイプヒントはまれで、async/await はほとんど存在しませんでした。タスクごとにコンテナをスピンアップすることはタブーと考えられていました。CUDA ドライバの競合と依存関係の管理により、コンテナ化された ML が困難になり、ルートレス ビルドは後回しになりました。 Kubernetes はほとんど機能しませんでした。私たちは v1.10 で実行しましたが、その機能の半分は実験的なものであり、頼りになるマネージド製品はありませんでした。 Apache Arrow はまだ始まったばかりでした。数百台のマシンからなるクラスターを実行するのは大変でした。そのような背景に対して、私たちは、コンテナーのネイティブ実行、強力な型指定、デフォルトでの再現性など、当時は逆張りだと感じていましたが、今では当然のように感じられる賭けをしました。私たちは、2018 年に Kubernetes に賭ける前に、まず代替案を検討しました。Flyte v0.1 は Airflow 上に構築され、v0.2 は AWS Step Functions 上に構築されました。(私たちのストーリーの詳細については、昨年の発表をお読みください)
私たちは、ユーザーが自分自身を傷つけることなく、スケーラブルなワークフローを調整できるように Flyte 1 を設計しました。静的ワークフローは、実行前にコンパイルおよび型チェックが行われるため、正確性はシステム固有の特性でした。私たちはこれをワークフロー オペレーティング システムとして考えました。複雑さを抽象化し、ハードウェアに合わせて拡張し、テナントを安全に分離する必要があります。

y を実行し、毎回結果を再現します。この規律こそが、Flyte がミッション クリティカルなワークロードで信頼を得た理由であり、今日の Flyte 2 を耐久性のあるものにしているのと同じ規律です。
しかし、規模と安全性を優先するにはコストがかかりました。Flyte は学習するのが難しすぎました。プロミス、動的なワークフロー、独自の条件文: ユーザーが最初の実際のパイプラインを作成する前に、概念的に重い抽象化を内部化する必要がありました。確かに YAML よりも優れていますが、ハードルは低いです。そしてユーザーは、敬意を表して、容赦なく私たちをダイナミズムに向けて推し続けてくれました。
そうしたユーザーが誰であるかという点では、私たちは非常に幸運でした。私たちは、LinkedIn、Spotify、Stripe、Tesla、Expedia、Toyota、Mercedes-Benz、Amazon、NVIDIA、Apple などの世界有数のエンジニアリング組織や、OpenAI、Runway、Luma Labs、Mistral、Wayve、Applied Intuition などの新世代の AI ネイティブ企業と協力することになりました。
彼らと一緒に働いて、私たちは学びました。私たちは、基礎的な部分 (データ、モデル、コンピューティング、確実に調整されたもの) がそのまま残されたまま、新しいワークロードが出現するのを観察しました。そして私たちは、AI の世界には、ユーザーのインフラストラクチャ内に組み込まれた 1 つの共通のオープンな実行層が必要であるとこれまで以上に確信するようになりました。
2024 年までに、エコシステムはまさに私たちが賭けていた方向に成熟しました。 Async Python には大きな勢いがあり、無料のスレッド Python が目前に迫っています。コンテナー ビルド、CUDA、コンテナー化された GPU ワークロードは許容できるだけでなく、期待されていました。 Kubernetes はこれまでよりも安定しており、管理対象クラスターは数千ノードまで拡張できました。そして、新しいタイプのワークロードが出現しました。これは、完全に動的な制御フローと完全に動的なインフラストラクチャ、つまり実行時に形状が決定されるプログラム、場合によっては人間によって、そしてますますモデルによって決定されるプログラムです。
ある AI ラボは次のように言いました。

本質的に動的でパスに依存するプロセスの固定パスをハードコーディングすることはできません。」
そこで、2025 年の初めに、私たちは難しい決断を下しました。私たちはまったく新しい Flyte をゼロから構築することに着手しました。フライト2。
私たちは次の 3 つの目標を掲げました。
学習を簡単にし、可能な限り存在しないようにします。 Python を書く (または生成する) ことができれば、Flyte を書くことができます。 DSL も約束もワークフロー言語もメンタルモデル税もありません。タスクは単にタスクを呼び出します。ループ、条件、Try/Exception: ユーザーがロジックを作成し、Flyte がそれを調整します。 Flyte は目に見えないと感じられるはずです。
コード レベルだけでなく、インフラストラクチャ レベルでも耐久性のある基盤を構築します。現在構築されているシステムは、エージェントによって作成、起動、デバッグ、さらには修復が行われており、長時間実行されている AI ワークロードを実際に破壊する原因のほとんどはバグではなく、インフラストラクチャの障害です。つまり、OOM のクラッシュ、プリエンプトされたノード、実行中に消える GPU などです。 Flyte 2 には、同じインフラストラクチャの前提条件に対してコードを再試行するだけでなく、これらの障害から回復する能力が必要でした。
リアルタイムとバッチを統合します。何かが長期間存続するサービスとして実行されるか、スケジュールされたパイプラインとして実行されるかは、たまたま選択したインフラストラクチャの偶然ではなく、問題の特性である必要があります。 1 つのプログラミング モデル、1 つのパッケージング モデル、両方のモダリティ。
初のオープンソース AI ランタイム
その結果が、オープンソース初の耐久性のある AI ランタイムである Flyte 2 です。
これが実際に何を意味するかは次のとおりです。これは完全な Flyte 2 プログラムです。
クリップボードにコピーされました!
非同期をインポートする
輸入フライト
env = flyte.TaskEnvironment(
名前 = "hello_world",
image=flyte.Image.from_debian_base(python_version=(3, 12)),
)
@env.task
def 計算(x: int) -> int:
x * 2 + 5 を返します
@env.task
async def main(numbers: list[int]) -> float:
results = await asyncio.gather(*[calculate.aio(n) for n in数値])
レット

urn sum(結果) / len(結果)
__name__ == "__main__"の場合:
flyte.init()
run = flyte.run(main,numbers=list(range(10)))
print(実行.結果)
純粋なパイソン。 DSLはありません。ワークフロー オブジェクトはありません。タスクはタスクを呼び出します。ファンアウトは単なる asyncio.gather です。そして、この正確なプログラムはラップトップ上、 devbox 上、またはクラスター全体で実行され、各タスクは独自のコンテナー内で、すべての入出力は型指定され、バージョン管理され、耐久性があり、すべての実行は停止した場所から回復可能です。プログラムは他の実行層に渡されません。 Flyte は実行層です。それがランタイムの意味です。
Flyte 2 がどのようにしてコードとインフラストラクチャの両方の障害に対してワークフローを耐久性のあるものにするかを見てみましょう。
Flyte 2 はインフラ対応であるため、平均的な耐久性のある実行プラットフォームとは根本的に異なります。 Flyte 2 は、実行時にコンピューティング リソースに影響を与えることにより、インフラストラクチャによって引き起こされる障害など、他の人では不可能な世界全体の障害を解決できます。
耐久性は、非決定論的な LLM 駆動のプロセスにも決定性をもたらすリプレイ ログとトレースから得られます。つまり、クラッシュし、回復し、中断したところから正確に続行されます。この部分、クラッシュ後のロジックの再生は、耐久性のある実行エンジンにとって非常に重要です。
重要でないのは、クラッシュがコードではなくインフラストラクチャに起因する場合に何が起こるかです。 OOM キル、プリエンプトされたスポット インスタンス、トレーニング実行中に消滅する GPU: Flyte 2 は障害モードを認識し、再生するコード パスだけでなく、再試行時に要求するリソースを変更します。それが、バグに対して耐久性のあるシステムと、バグが実行されるインフラストラクチャに対して耐久性のあるシステムとの違いです。 AI と ML のワークロードは常に第 2 の種類に該当します。
すべての実行はログに記録され、単一の UI から再現できるため、失敗しても検査して再現できます。

、盲目的に再実行して、うまくいくことを期待するものではありません。中核ではまだ言語に依存しませんが、Python での使用は非常に簡単で、パフォーマンスが重要な場合はホット パスを Rust で書き直しました。エンジンは Flyte 1 よりも何倍も高速で、データの移動は s5cmd などの専用ツールに対して独自の性能を発揮し、単一のプロセスで数千の同時タスクを実行できます。すべてが瞬時に感じられます。
エージェント ネイティブが実際に何を意味するかは次のとおりです。エージェントのワークフローは文字通り while ループになる可能性があります。
クリップボードにコピーされました!
gpu = flyte.TaskEnvironment(name="worker", resource=flyte.Resources(gpu=1))
@flyte.trace
async def plan(history: list[str]) -> str:
... # LLM 呼び出しがトレースされるため、クラッシュはゼロからではなくここで再開されます
@gpu.task
async def act(ステップ: str) -> str:
... # 重労働、このステップ専用にプロビジョニングされた GPU 上で
@env.task
非同期デフォルトエージェント(目標: str) -> リスト[str]:
歴史 = [目標]
while (ステップ := await plan(history)) != "完了":
History.append(動作(ステップ)を待つ)
返品履歴
動的な制御フロー、耐久性のある非決定的なステップ、ループの途中で実現してその後消える GPU : エージェント フレームワークは必要なく、すでに使用しているフレームワークは内部で問題なく実行されます。
ループ自体以外にも、Flyte 2 にはオーケストレーション サンドボックスが同梱されているため、LLM で生成されたパイプラインはデフォルトで分離され、ネットワークがブロックされて安全に実行されます。これは、モデルがワークフローを作成する場合、ランタイムが構築によってそれを安全にする必要があるためです。人間参加型の機能をファーストクラスの構造として出荷するため、ワークフローは人の判断のために一時停止し、到着した瞬間に再開できます。 MCP を話すため、コーディング エージェントまたは IDE が Flyte 2 を直接操作できます。つまり、実行の起動、状態の検査、障害の修正が可能です。
別個のサービススタックを必要としないリアルタイム + バッチ推論
チームが別の推論ステーションをつなぎ合わせているのを私たちは何度も見てきたからです。

オーケストレーターの隣に、同じモデルに推論とサービスを含めました。長時間実行されるアプリも単なる環境の 1 つです。
クリップボードにコピーされました!
fastapi から FastAPI をインポート
輸入フライト
flyte.app.extras から FastAPIAppEnvironment をインポート
app = FastAPI()
serving = FastAPIAppEnvironment(name="my-model", app=app)
@app.get("/predict")
async def detect(x: float) -> dict:
return {"result": await core.aio(x)} # パイプラインが使用するのと同じタスクを呼び出します
同じプログラミング モデル、同じパッケージング、同じプラットフォーム。コードがスケジュールされたパイプライン、エージェント ループ、またはリアルタイム エンドポイントとして実行されるかどうか。モダリティは問題の選択であり、インフラストラクチャではありません。 1 セットのプリミティブを使用してエージェントをトレーニング、評価、提供、調整します。
データは堀です。そのままにしておいてください。
Flyte の設計にはもう 1 つの信念が深く組み込まれており、AI 時代にはこれまで以上に重要になります。
この世界では、データ、コンテキスト、プロセスが堀です。勝利したチームは、インテリジェンスを消費するだけでなく、データに基づいて調整されたモデル、プロセスによって形成されたエージェント、その動作方法をエンコードしたシステムなど、独自のシステムを構築します。独自のインテリジェンスを構築するには、その基盤となる基盤を所有する必要があります。

[切り捨てられた]

## Original Extract

Durability across logic and infra. Reproducible, type safe, production scale.

Flyte 2 Is Generally Available: The Durable, Open-Source AI Runtime | Union.ai
Union.ai achieves 9.8x ROI according to analysts.
Union.ai
Union.ai
Platform
Platform Overview Build, ship, and scale AI faster than ever before.
Compare our open-source and production AI runtimes.
Union.ai is built so your data never leaves your own security perimeter.
Case Studies Real results reported by real teams
Read the latest from the Union.ai team
Join us for virtual workshops, trainings, and more
Explore details about support tiers
Browse the verified Union.ai integrations
Flyte 2 Is Generally Available: The Durable, Open-Source AI Runtime
Example H3 Example H4 Example H5 Example H6
https://www.union.ai/blog-post/flyte-2-is-generally-available-the-durable-open-source-ai-runtime
Today, Flyte 2 is generally available. It is faster, more scalable, and (my personal guarantee) genuinely fun to use. It remains open source and part of the Linux Foundation. Before I tell you what's in it, I want to tell you why it exists, because Flyte 2 is not an incremental release. It is a ground up rebuild of a system we have been refining for nearly a decade, rebuilt for a world that looks nothing like the one Flyte was born into.
This new evolution of Flyte has turned it into something more valuable than an orchestrator for AI and ML teams: a durable AI runtime . Where an orchestrator sits on top of systems to execute and forget, a runtime owns the infrastructure underneath your code, not just the code itself. Where a durable execution system can recover from logic failures, a durable runtime can use its infrastructure ownership to recover from compute failures (e.g., OOM), too, for end-to-end runtime durability. That is what we mean when we call Flyte 2 a durable AI runtime instead of an orchestrator, and it is the thread running through everything below.
Flyte began at Lyft. We built it to prepare data, train models, run large scale simulations, and back test models: the unglamorous, essential machinery behind products like ETA and pricing. What we realized then still holds today: building ML products requires the reliability of user facing software, the dynamism of infrastructure that reshapes itself around each workload, and the agility to experiment constantly. No system at the time offered all three, so we built one.
Consider the world we built it in. Python was widely dismissed as a single threaded scripting language, type hints were rare, async/await barely existed. Spinning up a container per task was considered taboo: CUDA driver conflicts and dependency management made containerized ML painful, and rootless builds were an afterthought. Kubernetes barely worked; we ran on v1.10 with half its features experimental and no managed offerings to lean on. Apache Arrow had just gotten started. Running a cluster of a few hundred machines was nuts. Against that backdrop, we made bets, container native execution, strong typing, reproducibility by default, that felt contrarian then and feel obvious now. We even lived through the alternatives first: Flyte v0.1 was built on Airflow, v0.2 on AWS Step Functions, before we bet on Kubernetes in 2018. (For more about our story, read last year’s announcement )
We designed Flyte 1 to make sure users could orchestrate scalable workflows without shooting themselves in the foot. Static workflows, compiled and type checked before they ran, so correctness was an inherent property of the system. We thought of it as a workflow operating system: it should abstract complexity, scale with hardware, isolate tenants safely, and reproduce results every single time. That discipline is why Flyte earned trust for mission critical workloads, and it's the same discipline that makes Flyte 2 durable today.
But there was a cost to prioritizing scale and safety: Flyte became too hard to learn. Promises, dynamic workflows, our own conditionals: conceptually heavy abstractions users had to internalize before writing their first real pipeline. Better than YAML, sure, but that's a low bar. And our users kept pushing us, respectfully, relentlessly, toward dynamism.
We have been extraordinarily lucky in who those users are. We got to work with some of the best engineering organizations in the world: LinkedIn, Spotify, Stripe, Tesla, Expedia, Toyota, Mercedes-Benz, Amazon, NVIDIA, Apple, and a new generation of AI native companies like OpenAI, Runway, Luma Labs, Mistral, Wayve, Applied Intuition, and many more.
Working alongside them, we learned. We watched new workloads emerge while the foundational pieces (data, models, compute, coordinated reliably) stayed intact. And we became more convinced than ever that the AI world needs one common, open layer of execution embedded within users’ infrastructure.
By 2024, the ecosystem had matured in exactly the directions we'd bet on. Async Python had serious momentum, with free threaded Python on the horizon. Container builds, CUDA, and containerized GPU workloads were not just acceptable but expected. Kubernetes was more stable than it had ever been, with managed clusters scaling to thousands of nodes. And a new type of workload was emerging: one that needed durability across a fully dynamic control flow and fully dynamic infrastructure, programs whose shape is decided at runtime, sometimes by a human, increasingly by a model.
As one AI lab put it to us: “You can't hardcode a fixed path for an inherently dynamic, path-dependent process.”
So in early 2025, we made the hard call. We set out to build a whole new Flyte, from the ground up. Flyte 2.
We held ourselves to three goals:
Make learning easy , as close to nonexistent as possible. If you can write (or generate) Python, you can write Flyte. No DSL, no promises, no workflow language, no mental model tax. Tasks simply call tasks. Loops, conditionals, try/except: you write the logic, Flyte orchestrates it. Flyte should feel invisible.
Build a foundation that is durable at the infrastructure level, not just the code level . The systems being built today are authored, invoked, debugged, and even repaired by agents, and most of what actually breaks a long running AI workload isn't a bug, it's an infrastructure failure: an OOM crash, a preempted node, a GPU that disappears mid-run. We needed Flyte 2 to have the power to recover from these failures, not just retry code against the same infrastructure assumptions.
Unify real time and batch. Whether something runs as a long-lived service or a scheduled pipeline should be a property of the problem, not an accident of the infrastructure you happened to choose. One programming model, one packaging model, both modalities.
The first open-source AI runtime
The result is Flyte 2: the first durable AI runtime for open source.
Here's what that means in practice. This is a complete Flyte 2 program:
Copied to clipboard!
import asyncio
import flyte
env = flyte.TaskEnvironment(
name="hello_world",
image=flyte.Image.from_debian_base(python_version=(3, 12)),
)
@env.task
def calculate(x: int) -> int:
return x * 2 + 5
@env.task
async def main(numbers: list[int]) -> float:
results = await asyncio.gather(*[calculate.aio(n) for n in numbers])
return sum(results) / len(results)
if __name__ == "__main__":
flyte.init()
run = flyte.run(main, numbers=list(range(10)))
print(run.result)
Pure Python. No DSL. No workflow objects. Tasks call tasks; fanout is just asyncio.gather. And this exact program runs on your laptop, on a devbox , or across a cluster, each task in its own container, every input and output typed, versioned, and durable, every run recoverable from where it stopped. Your program doesn't get handed off to some other execution layer; Flyte is the execution layer. That's what runtime means.
Let’s explore how Flyte 2 makes workflows durable from both code and infrastructure failures.
Flyte 2 is fundamentally different from your average durable execution platform because it is infra-aware. By letting you affect your compute resources at runtime, Flyte 2 can solve an entire world of failures that others can’t – the ones caused by infrastructure.
Durability comes from a replay log and traces that bring determinism even to non-deterministic, LLM-driven processes: crash, recover, and continue exactly where you left off. That part, replaying logic after a crash, is table stakes for any durable execution engine.
What isn't table stakes is what happens when the crash comes from the infrastructure, not the code. An OOM kill, a preempted spot instance, a GPU that vanishes mid-training run: Flyte 2 recognizes the failure mode an changes the resources it asks for on the retry , not just the code path it replays. That's the difference between a system that's durable against your bugs and one that's durable against the infrastructure your bugs run on. AI and ML workloads hit the second kind constantly.
Every execution is logged and reproducible from a single UI, so a failure is something you can inspect and replay, not something you re-run blind and hope goes better. It is still language agnostic at its core, but extremely simple to use in Python, and where performance mattered, we rewrote the hot paths in Rust: the engine is many times faster than Flyte 1, data movement holds its own against dedicated tools like s5cmd, and a single process can drive thousands of concurrent tasks. Everything feels instant.
Here's what agent native means in practice: your agent's workflow can literally be a while loop.
Copied to clipboard!
gpu = flyte.TaskEnvironment(name="worker", resources=flyte.Resources(gpu=1))
@flyte.trace
async def plan(history: list[str]) -> str:
... # your LLM call, traced, so a crash resumes here, not from zero
@gpu.task
async def act(step: str) -> str:
... # heavy lifting, on a GPU provisioned just for this step
@env.task
async def agent(goal: str) -> list[str]:
history = [goal]
while (step := await plan(history)) != "done":
history.append(await act(step))
return history
Dynamic control flow, durable non-deterministic steps, and a GPU that materializes mid-loop and disappears after: no agent framework required, and any framework you already use runs happily inside.
Beyond the loop itself, Flyte 2 ships with orchestration sandboxes so LLM-generated pipelines run safely, isolated and network blocked by default, because if models are going to write workflows, the runtime must make that safe by construction. It ships human-in-the-loop as a first class construct, so a workflow can pause for a person's judgment and resume the moment it arrives. It speaks MCP, so your coding agent or IDE can operate Flyte 2 directly: launch runs, inspect state, fix failures.
Realtime + batch inference without a separate serving stack
Because we kept seeing teams stitch a separate inference stack next to their orchestrator, we included inference and serving in the same model. A long running app is just another environment:
Copied to clipboard!
from fastapi import FastAPI
import flyte
from flyte.app.extras import FastAPIAppEnvironment
app = FastAPI()
serving = FastAPIAppEnvironment(name="my-model", app=app)
@app.get("/predict")
async def predict(x: float) -> dict:
return {"result": await score.aio(x)} # calls the same task your pipeline uses
Same programming model, same packaging, same platform, whether the code runs as a scheduled pipeline, an agent loop, or a real time endpoint. The modality is a choice of the problem, not the infrastructure. Train, evaluate, serve, and orchestrate agents with one set of primitives.
Your data is your moat. Keep it.
There's one more belief embedded deep in Flyte's design, and it matters more in the AI era than ever before.
In this world, your data, your context, and your processes are your moat. The teams that win won't just consume intelligence, they'll build their own: models tuned on their data, agents shaped by their processes, systems that encode how they operate. Building proprietary intelligence means you need to own the foundations it stands on, and you nee

[truncated]
