---
source: "https://coroot.com/blog/hard-part-of-ai-root-cause-analysis-is-no-longer-the-model/"
hn_url: "https://news.ycombinator.com/item?id=48751050"
title: "The hard part of AI root cause analysis is no longer the model"
article_title: "The hard part of AI root cause analysis is no longer the model | Coroot Blog"
author: "nikolay_sivko"
captured_at: "2026-07-01T18:37:15Z"
capture_tool: "hn-digest"
hn_id: 48751050
score: 1
comments: 0
posted_at: "2026-07-01T18:16:53Z"
tags:
  - hacker-news
  - translated
---

# The hard part of AI root cause analysis is no longer the model

- HN: [48751050](https://news.ycombinator.com/item?id=48751050)
- Source: [coroot.com](https://coroot.com/blog/hard-part-of-ai-root-cause-analysis-is-no-longer-the-model/)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T18:16:53Z

## Translation

タイトル: AI 根本原因分析の難しい部分は、もはやモデルではない
記事タイトル: AI 根本原因分析の難しい部分は、もはやモデルではない |コルートのブログ
説明: 11 個の LLM を通じて 1 つの実際のインシデントを実行しました。推論は簡単であることが判明しました。本当の課題は、モデルに適切なコンテキストを供給するハーネスです。

記事本文:
AI 根本原因分析の難しい部分は、もはやモデルではありません |コルートのブログ
メイン コンテンツにスキップ 製品 製品概要 AI 根本原因分析 エージェント対応の可観測性 Coroot Enterprise Edition エディションの比較 ソリューション Kubernetes の可観測性 PostgreSQL の可観測性 GitOps の可観測性 eBPF の可観測性 OpenTelemetry 継続的プロファイリング 価格設定 リソース ドキュメント ブログ AI RCA ベンチマーク ポッドキャストの概念と用語 Coroot Slack コミュニティ デモ デモの予約 ライブ デモ カスタマー ポータル トグル メニュー ← すべての投稿 エンジニアリング ハードAI 根本原因分析の一部はモデルではなくなりました
数週間ごとに誰かが、根本原因分析は現在解決された問題であると教えてくれます。テレメトリを LLM にパイプすると、何が問題なのかを教えてくれます。そんなに簡単だったらいいのに。これを何年も続けてきた私は、「AI で RCA ができるだろうか？」と考えています。これは間違った質問です。LLM を使用して RCA を実行することは実際には 2 つの別個の作業であり、それぞれの答えが異なるからです。それらは完全に異なる方法で壊れるので、分解する価値があります。
1 つは推論です。モデルはその前のデータを取得して点を結びつけることができるでしょうか?サービスが遅くなります。 CPU が不足している、ノードの CPU が最大になっている、そのノード上の近隣ノードが CPU をすべて食べているという 3 つの事実が同時にテーブルにあります。理由を考えてそれらを 1 つの物語、騒々しい隣人に結び付けるモデル。より弱いものは、無関係な 3 つの「問題」を報告するか、最も大きな症状を取り上げて、それを原因と呼びます。
もう 1 つはハーネス、つまりモデルの周囲にあるすべてのものです。どのようなデータをどのような形式でその前に配置するか。通常、これはツール呼び出しを意味し、何を取得するか、いつ停止するかをモデルに決定させます。ここでは多くの間違いが発生しますが、モデルが推論できるかどうかについては何も問題がありません。適切なデータが得られなかっただけです。
人々はいつもこの 2 つを混同します。モデルは悪い答えを出しますが、

そしてLLMにはRCAはできないと誰もが言います。しかし通常、モデルは必要なデータを取得できませんでした。論理的に考えることができなかったわけではなく、決して公平な判断ができなかったのです。そして、この 2 つを区別するまでは、どちらが本当の問題であるかわかりません。
写真からハーネスを取り出します #
それで、私たちは意図的にそうしました。 Coroot の AI RCA を使用すると、モデル ツールを渡して調査に送る必要はありません。代わりに、決定論的パイプラインが重労働を行い、信号を相関させて結果に変換します。モデルは、生のテレメトリではなく、1 つの焦点を絞ったコンテキストでこれらの結果を取得します。ツールもエージェント ループもありません。答えを見つけるために必要なものはすべてすでにそこにあります。
それはすべてを推論に要約します。モデルに完全なコンテキストが含まれていても根本原因が見落とされている場合、他に責められる人はいません。ハーネスやデータの欠落ではありません。ただのモデルです。そして、それがついに正確に測定できるようになりました。
それでは実験です。コンテキストにすでに答えが含まれている実際の障害を取り上げ、その同じダンプを多数のモデルに渡し、どのモデルが実際の根本原因を抽出できるかを確認します。取ってくる必要もなければ、何を見るかを決める必要もありません。ただの推理。そして、それは思っているよりも難しいのです。たとえ答えがデータの中にあったとしても、モデルを誤った結論に導く可能性のある罠がそこには存在します。
私は 1 つのシナリオを選択しました。それは、カタログ サービスとその Postgres データベース db-main の間のネットワーク遅延です。クエリが遅くなり、タイムアウトが広がり、フロントエンドが 502 の処理を​​開始します。しかし、実際にはデータベースやサービスには何も問題はありません。原因は、クラスター内で実行されている Chaos Mesh NetworkChaos 実験で、カタログ ↔ db-main パスに遅延をもたらし、それが Kubernetes イベントに直接現れます。したがって、修正するには実験を削除することと、同じくらい重要なこととして、スケジュールを削除することです。

そうすればすぐに元に戻ります。
問題が広がっていることがわかります。フロントエンドにはエラーと遅延が表示されますが、依存関係チェーンを追跡すると、実際の信号が存在するカタログ、TCP ネットワーク、および db-main への接続遅延につながります。
このマップは、根本原因ではないもの、つまり kafka のレイテンシ、カタログと db-main の CPU、データベースのストレージにもフラグを立てていることに注意してください。分散システムでは、1 つの問題があらゆる場所のメトリクスに影響し、その一部は間違った方向を示します。データベースを取得します。カタログと db-main の間のラウンドトリップ時間が増加すると、クライアントのクエリ応答が遅くなり始めました。ただし、Postgres は、受信した最初のバイトからクライアントに送り返す最後のバイトまでのクエリを計測するため、ネットワーク遅延はクエリ時間の一部としてカウントされます。 pg_stat_statements を読むと、まったく同じクエリでデータベースが突然遅くなったことがわかります。
そうではありませんでした。延長時間は Postgres 内ではなくネットワーク上にありました。単純な読み取りはデータベースのせいにして先に進みますが、これがまさに罠です。前処理によってこのような信号が表面化するのはバグではありません。次のステップでは、下流の影響から真の原因を伝えることがモデルの仕事であり、それには真の推論が必要です。モデルがデータにだまされないように十分に推論できる限り、ここでデータが増えても問題はありません。
それらの罠の他に、本当の証拠はすべてそこにあります。 Coroot は伝播経路を追跡し、速度低下を追跡する db-main へのネットワーク RTT を発見し、問題が発生した瞬間にカオス実験が開始されたことを示す Kubernetes イベントにフラグを立てました。答えはプロンプトの中にあり、何も隠されておらず、ツール呼び出しの背後にも何もありません。
次に、すべてのモデルに対してまったく同じプロンプトを再利用し、それぞれに同じ 3 つのことを尋ねました。根本原因は何か、原因と結果の連鎖は何か、そして、

は即時修正です。
得点はなく、合格か不合格だけです。これを正直な数字に変換することはできませんでした (半分正しい説明には何ポイントの価値がありますか?)。そのため、バーは単純です。カオス実験を固定し、フロントエンドへのリップルを説明し、実験とそのスケジュールの両方を削除することを示していますか?
ここがバーです。クロード作品 4.8 はまさにそれです。
この概要では、フロントエンド エラー、catalog の遅い Postgres クエリ、NetworkChaos の挿入遅延に戻るという正しい順序でチェーンが実行され、相関する RTT と TCP のスパイクが証拠として挙げられています。
証拠は当てはまります。クエリが遅いと RTT と接続時間が上昇し、 gorm.Query スパンが遅くなり、実験の開始と回復がいつ行われたかを示す Kubernetes イベントに加えて、 db-main の CPU プロファイルがクリーンであるため、これはデータベースの問題ではないという注意事項があります。
そして、この修正は実際に役に立ちます。すぐに実験を削除して接続を回復し、その後、スケジュールを削除して、再度実行できないようにします。この 2 番目のステップは見逃しやすいものですが、まさに午前 3 時に知りたいことです。
では、他に誰がクリアしたか見てみましょう。
11 個のモデル、同じコンテキスト、それぞれ 1 ショット。入力は ~9,800 トークン、出力は ~1,000 であるため、コスト列はインシデントごとの実際の価格です。
結果キー: ✅ 正しい根本原因と正確な完全な修正。 🟡 根本原因は見つかりましたが、修正が不完全だったか、フォーマット手順が無視されました。 ❌ 根本原因を見逃していると、修正は実際には機能しません。
フィールドを次の 3 つの方法で分割すると便利です。
クローズドフロンティアモデル：Opus 4.8、GPT-5.5、Gemini 3.1 Pro。その大きさは分かりませんが、ほぼ間違いなく巨大です。
クラウドで実行できるような大きなオープンウェイト モデル: DeepSeek V4、GLM-5.2、MiniMax M3、Nemotron 3 Ultra。
実際にセルフホストできる小型のオープンウェイト モデル: Gemma 4 31B と Qwens。
トップに驚くことはありません:

3 つのフロンティア モデルはすべてき​​れいに合格し、大手のオープン モデルはほぼ追いつきました。
驚きはスモールエンドにありました。自分のハードウェアで実際に実行できるモデルの中で、サイズはほとんど問題にならず、際立ったものは群の中で最も小さいものでした。Gemma 4 31B は、根本原因をまったく見つけることができた唯一の自己ホスト型モデルでした。より大きな Qwen3.6 35B と Qwen3 Coder Next は両方とも失敗しました。また、大規模化するという保証もありませんでした。Nemotron 3 Ultra は 5,500 億のパラメーターで、中間層にしか到達しませんでした。自分で走れるものが欲しいなら、ジェマが最適でした。
ただし、誰が合格し、誰が不合格だったのかをよく見てみると、その分かれ目は 2 種類の失敗に帰着します。そのうちの 1 つだけが実際にモデルに関するものです。
❌ のケースは実際の推論の失敗です。 Qwen3 Coder Next はクエリが遅いことを確認し、データベースのせいだとしましたが、これは明らかに間違った答えでしたが、カオス イベントはコンテキスト内に留まっていたのです。 Qwen3.6 は実験を見つけましたが、間違ったリソースを削除し、実行されないコマンドを返しました。どちらも必要なものはすべて揃っていましたが、まだ欠けていました。それを修正できるハーネスはありません。
🟡 の場合はより期待が高まります。彼らは全員、推論を正しく理解していましたが、パッケージ化でつまずいただけでした。生のウィジェット ID が証拠に漏洩し、コード ブロック内で野良バッシュが発生し、カオス オブジェクトは削除されましたが、それを再作成するスケジュールは削除されませんでした。これらはフォーマットと指示に従う伝票であり、モデルではなくハーネス側に固定する種類のものです。これらのモデルについては、その推論は適切でした。
表には大きなスプレッドが示されているので、コストについて一言。一見したところ、モデル間のギャップは非常に大きいように見えます。ただし、LLM はここでは生のテレメトリを分析していません。その作業は呼び出される前にすでに完了しています。約 10,000 トークンの 1 回の短い呼び出しで、小さな既製のコンテキストとそれに関する理由を読み取るだけです。そのサイズでは、最も経験値が高くても

エンシブ フロンティア モデルの料金は 1 件あたり数セントなので、最上位モデルに手を伸ばすのは完全に合理的であるように見えます。
1 つのインシデントはベンチマークではありませんが、私たちが見続けているものと一致しています。 AI RCA の推論部分は基本的に解決されています。フロンティア モデルはそれを成功させ、自己ホストできる小さなモデルでも、まだそれほど予測可能ではありませんが、まともな仕事をします。コンテキストがわかれば、根本原因を見つけることは、実際には未解決の問題ではなくなります。
難しいのはハーネスです。テレメトリは急速に成長し続けており、生データをすべて LLM に渡して分類すると、すぐに速度が遅くなり、コストも高くなります。したがって、本当の仕事は、より賢いモデルを見つけることではありません。モデルを呼び出す前に、モデル用の適切でコンパクトなコンテキストが準備されます。
コード変更なしで、フルスタックの可観測性を数分で実現します。 AI による根本原因分析による eBPF を活用したモニタリング。
コルートの概念と用語
© 2026 コルート株式会社、全著作権所有

## Original Extract

We ran one real incident through eleven LLMs. The reasoning turned out to be the easy part; the real challenge is the harness that feeds the model the right context.

The hard part of AI root cause analysis is no longer the model | Coroot Blog
Skip to main content Product Product Overview AI Root Cause Analysis Agentic-ready Observability Coroot Enterprise Edition Compare Editions Solutions Kubernetes Observability PostgreSQL Observability GitOps Observability eBPF Observability OpenTelemetry Continuous Profiling Pricing Resources Documentation Blog AI RCA Benchmark Podcast Concepts and Terminology Coroot Slack Community Demo Book a Demo Live Demo Customer Portal Toggle menu ← All posts Engineering The hard part of AI root cause analysis is no longer the model
Every few weeks someone tells me root cause analysis is a solved problem now: pipe your telemetry into an LLM, let it tell you what broke. I wish it were that easy. After years on this, I think "can AI do RCA?" is the wrong question, because doing RCA with an LLM is really two separate jobs, and the answer is different for each. They break in completely different ways, so it's worth pulling them apart.
One is reasoning: can the model take the data in front of it and connect the dots? A service slows down. Three facts are on the table at once: it's starved of CPU, the node's CPU is maxed, and a neighbor on that node is eating all of it. A model that reasons ties them into one story, a noisy neighbor. A weaker one reports three unrelated "issues", or grabs the loudest symptom and calls it the cause.
The other is the harness: everything around the model. What data you put in front of it, in what shape. Usually it means tool-calling, letting the model decide what to fetch and when to stop. Plenty goes wrong here, and none of it is about whether the model could reason. It just never got the right data.
People mix these two up all the time. A model gives a bad answer, and everyone says LLMs can't do RCA. But usually the model just never got the data it needed. It's not that it couldn't reason, it never had a fair shot. And until you separate the two, you can't tell which one is the real problem.
Take the harness out of the picture #
So we did, on purpose. With Coroot's AI RCA , we don't hand the model tools and send it off to investigate. Instead, a deterministic pipeline does the heavy lifting: it correlates the signals and turns them into findings. The model gets those findings in one focused context, not the raw telemetry. No tools, no agent loop. Everything it needs to find the answer is already there.
That boils the whole thing down to reasoning. If the model has the full context and still misses the root cause, there's no one else to blame. Not the harness, not missing data. Just the model. And that's finally something you can measure cleanly.
So here's the experiment. Take a real failure where the context already holds the answer, hand that same dump to a bunch of models, and see which ones can distill it into the actual root cause. No fetching, no deciding what to look at. Just reasoning. And it's harder than it sounds: even with the answer sitting in the data, there are traps in there that can walk a model straight into the wrong conclusion.
I picked one scenario: a network delay between the catalog service and its Postgres database, db-main . The queries slow down, timeouts spread, and front-end starts serving 502s. But nothing is actually wrong with the database or the service. The culprit is a Chaos Mesh NetworkChaos experiment running in the cluster, injecting delay on the catalog ↔ db-main path, and it shows up right in the Kubernetes events. So the fix is to delete the experiment, and just as importantly, the schedule that would spin it right back up.
You can see the problem fan out: front-end shows errors and latency, but tracing the dependency chain leads through catalog , where the real signal lives, the TCP network and connection latency to db-main .
Notice the map also flags things that aren't the root cause: latency on kafka , CPU on catalog and db-main , storage on the database. In a distributed system one problem bleeds into metrics everywhere, and some of them point the wrong way. Take the database. When the round-trip time between catalog and db-main went up, the client started getting its query responses slower. But Postgres times a query from the first byte it receives to the last byte it sends back to the client, so that network delay gets counted as part of the query time. Read pg_stat_statements and it looks like the database suddenly got slower at the exact same queries.
It didn't. The extra time was on the wire, not inside Postgres. A naive read blames the database and moves on, which is exactly the trap. And it isn't a bug that the pre-processing surfaces a signal like this. Telling a real cause from its downstream effects is the model's job in the next step, and it takes real reasoning. More data here isn't a problem, as long as the model can reason well enough not to be fooled by it.
Alongside those traps, the real evidence is all there. Coroot traced the propagation path, found the network RTT to db-main tracking the slowdown, and flagged the Kubernetes event showing the chaos experiment started right when things broke. The answer is right there in the prompt, nothing hidden, nothing behind a tool call.
I then reused that exact same prompt against every model, asking each the same three things: what's the root cause, what's the cause-and-effect chain, and what's the immediate fix.
No score, just pass or fail. I couldn't turn this into an honest number (how many points is a half-right explanation worth?), so the bar is simple: did it pin the chaos experiment, explain the ripple out to front-end , and point at deleting both the experiment and its schedule?
Here's the bar. Claude Opus 4.8 nails it.
The summary runs the chain in the right order: front-end errors, back through catalog 's slow Postgres queries, back to the NetworkChaos injecting delay, with the correlated RTT and TCP spikes cited as proof.
The evidence holds up: RTT and connection time climbing with the slow queries, the slow gorm.Query spans, the Kubernetes events marking when the experiment started and recovered, plus a note that db-main 's CPU profile is clean, so this isn't a database problem.
And the fix is actually useful: delete the experiment now to get connectivity back, then delete its schedule so it can't fire again. That second step is the easy one to miss, and exactly what you'd want to know at 3am.
Now let's see who else clears it.
Eleven models, same context, one shot each. Input is ~9,800 tokens, output ~1,000, so the cost column is the real per-incident price.
Result key: ✅ correct root cause and a correct, complete fix. 🟡 found the root cause, but the fix was incomplete or it ignored the formatting instructions. ❌ missed the root cause, or the fix wouldn't actually work.
It helps to split the field three ways:
Closed frontier models : Opus 4.8, GPT-5.5, Gemini 3.1 Pro. We don't know their size, but almost certainly huge.
Big open-weight models , the kind you'd still run in the cloud: DeepSeek V4, GLM-5.2, MiniMax M3, Nemotron 3 Ultra.
Small open-weight models you could actually self-host: Gemma 4 31B and the Qwens.
No surprises at the top: all three frontier models passed cleanly, and the big open models mostly kept up.
The surprise was at the small end. Among the models you could realistically run on your own hardware, size barely mattered, and the standout was the smallest of the bunch: Gemma 4 31B was the only self-hostable model to find the root cause at all. The bigger Qwen3.6 35B and Qwen3 Coder Next both missed it. And going huge was no guarantee either: Nemotron 3 Ultra, at 550 billion parameters, only made the middle tier. If you want something you can run yourself, Gemma was the one to beat.
Look closer at who passed and who didn't, though, and the split comes down to two kinds of failure, only one of which is really about the model.
The ❌ cases are real reasoning failures. Qwen3 Coder Next saw slow queries and blamed the database, the obvious wrong answer, while the chaos event sat right there in the context. Qwen3.6 found the experiment but then deleted the wrong resource, handing back a command that wouldn't run. Both had everything they needed and still missed, and no harness can fix that.
The 🟡 cases are more hopeful. They all got the reasoning right and only stumbled on packaging: raw widget IDs leaking into the evidence, a stray bash in a code block, deleting the chaos object but not the schedule that recreates it. Those are formatting and instruction-following slips, the kind you fix on the harness side, not the model. For these models, the reasoning was fine.
A word on cost, since the table shows a big spread. At first glance the gap between models looks huge. But the LLM isn't analyzing raw telemetry here, that work is already done before it's called. It just reads a small, ready-made context and reasons over it, one short call of around 10k tokens. At that size even the most expensive frontier model runs a few cents per incident, which makes reaching for a top model look perfectly reasonable.
One incident isn't a benchmark, but it lines up with what we keep seeing. The reasoning part of AI RCA is basically solved: frontier models nail it, and even small models you can self-host do a decent job, though they're still not as predictable. Finding the root cause, once the context is there, isn't really the open problem anymore.
The hard part is the harness. Telemetry keeps growing fast, and if you hand all that raw data to an LLM to sort through, it gets slow and expensive in a hurry. So the real work isn't finding a smarter model. It's preparing the right, compact context for the model before you call it.
Get full-stack observability in minutes with zero code changes. eBPF-powered monitoring with AI-guided root cause analysis.
Coroot Concepts and Terminology
© 2026 Coroot Inc., All Rights Reserved
