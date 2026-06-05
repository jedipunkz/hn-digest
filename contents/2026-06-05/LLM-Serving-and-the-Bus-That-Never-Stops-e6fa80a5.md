---
source: "https://joker666.github.io/blog/2026-06-02-llm-serving-in-flight-batching"
hn_url: "https://news.ycombinator.com/item?id=48414773"
title: "LLM Serving and the Bus That Never Stops"
article_title: "LLM Serving and the Bus That Never Stops | Rafi Hasan"
author: "joker666"
captured_at: "2026-06-05T18:46:41Z"
capture_tool: "hn-digest"
hn_id: 48414773
score: 1
comments: 0
posted_at: "2026-06-05T16:27:20Z"
tags:
  - hacker-news
  - translated
---

# LLM Serving and the Bus That Never Stops

- HN: [48414773](https://news.ycombinator.com/item?id=48414773)
- Source: [joker666.github.io](https://joker666.github.io/blog/2026-06-02-llm-serving-in-flight-batching)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T16:27:20Z

## Translation

タイトル: LLM サービスと止まらないバス
記事のタイトル: LLM サービスと止まらないバス |ラフィ・ハサン
説明: 実行中のバッチ処理は、LLM サービスが GPU シートを無駄にしないようにするためのトリックです。

記事本文:
LLM サービスと止まらないバス |ラフィ・ハサン ラフィ・ハサンのブログ
検索 ⌘ K ← BACK April 2, 2026 • 8 min read LLM サービスと止まらないバス
実行中のバッチ処理は、LLM サービスが GPU シートを無駄にしないようにするためのトリックです。
私は、機械学習モデルに対するリクエストのバッチ処理は解決された問題だと考えていました。私はモデルをホストし、リクエストに応えてきました。バッチ処理により、リクエストの処理が高速化されます。たとえば、バッチ処理を行わず、一度に 1 つの小さなリクエストを送信すると、GPU は 1 人の乗客を乗せた巨大なバスのように動作します。移動することはできますが、経済状況はひどいです。これがバッチ処理が存在する理由です。しかし、LLM はバッチ処理を奇妙にします。
従来の Web バックエンドや標準のコンピューター ビジョン パイプラインでは、これは簡単です。リクエストをキューに入れ、バッチ サイズが 4 または 8 に達するまで待機し、それらを GPU に叩き込み、結果を返します。標準エンジニアリング。モデル全体を 1 回移動します。
入力 -> モデル -> 出力
画像分類器がこれを行います。埋め込みモデルがこれを行います。データを渡すとモデルが実行され、結果が得られます。しかし、LLM サービスはこのメンタルモデルを完全に打ち破ります。 LLM の生成は反復的です。一度に 1 つのトークンを生成します。
プロンプト -> 事前入力
トークン1 -> デコード
トークン2 -> デコード
トークン 3 -> デコード
...
これは、LLM の提供が単に「モデルを 1 回実行する」だけではないことを意味します。これは、すべてのトークンを繰り返すスケジュールの問題です。 LLM リクエストを従来の Web リクエストと同様に扱うと、GPU 効率が急激に低下し、レイテンシーが急増し、クラウド料金が跳ね上がります。
LLM サービングはループです。すべてのトークンは、GPU を無駄にするか、GPU を埋めるチャンスとなります。
プレフィル、デコード、および KV キャッシュ
各リクエストには 2 つのフェーズがあります。最初のフェーズは prefill で、モデルはプロンプトを読み取り、内部アテンション状態を構築します。 2 番目のフェーズは decode で、autoregre を使用します。

ssive デコーディングにより、一度に 1 トークンずつテキストを生成し、生成された各トークンをモデルにフィードバックして次のトークンを予測します。
デコード ステップは終了条件が満たされるまで繰り返し実行されるため、リクエストの継続時間は大幅に異なります。すべてのステップでプロンプト履歴が再計算されるのを避けるために、サーバーは GPU メモリ内に KV キャッシュを維持します。スケジューラの目標は、この有限のキャッシュ メモリを使い果たさずに、GPU をトークン生成でビジー状態に保つことです。
3 つのリクエストが同時に届くと想像してください。
A -> 8 つの出力トークンが必要です
B -> 2 つの出力トークンが必要
C -> 6 つの出力トークンが必要
静的バッチ処理では、それらの乗客を同じバスに乗せ、新しい乗客を乗せる前にバスが全行程を終了するようにします。
ステップ 1: A B C
ステップ 2: A B C # B が完了しました
ステップ 3: A_C
ステップ 4: A_C
ステップ 5: A_C
ステップ 6: A _ _ # C が完了しました
ステップ 7: A _ _
ステップ 8: A _ _ # A が完了しました
B と C は早く終わっても、席の割り当てはできません。 GPU は実行を続けますが、メモリを浪費し、空のパディング トークンを計算します。それが無駄なのです。
固定ツアーバス vs. ダイナミック市内交通バス
私は、静的バッチ処理と機内バッチ処理の違いを、事前予約されたツアー バスと公共の都市交通バスの違いのように考えるのが好きです。
固定ツアー バス (静的バッチ処理): ツアー バスは、設定された乗客リストを持って駅を出発します。乗客が乗り場 2 で早めに降りる場合でも、残りの旅行中は座席を空席のままにしておく必要があります。バスは路上で新たな乗客を乗せることはできません。代わりに、新しいグループを積み込む前に、ツアー全体を完了してステーションに戻る必要があります。
ダイナミックシティトランジットバス（機内バッチング）：連続循環するトランジットバス。乗客が目的地に到着して降車するとすぐに、バスは次の停留所で一時停止し、空いた座席に新しい乗客を乗せます。

そしてすぐに旅を続けます。
LLM サービスでは、バスがアクティブなバッチです。シートは単なる「バッチスロット」ではありません。 GPU メモリと KV キャッシュの容量を表します。オフになるということは、リクエストが終了条件に達したことを意味します。ボーディングとは、新しいリクエストがアクティブな生成ループに参加するのに十分なメモリ バジェットを持っていることを意味します。
移動しながら乗客を変えるバス
各シートは、GPU メモリと KV キャッシュによってサポートされるバッチ スロットです。各ティックは 1 世代の反復です。
重要な部分は、バッチが固定されたリクエストのグループではなくなったことです。これは動的で流動的な変数です。
実行中のバッチ処理は、連続バッチ処理または反復レベルのバッチ処理とも呼ばれます。では、vLLM や TensorRT-LLM などのエンジンは実際にこの「動的バス」をコードでどのように実装するのでしょうか?これらは、スケジューリング境界を要求レベルから反復レベルにシフトします。
生成を繰り返すたびに、スケジューラは次のことを尋ねます。
どのリクエストがまだアクティブですか?
どのような新しいリクエストが待っていますか?
KV キャッシュ領域は十分ですか?
レイテンシをあまり損なわずに新しい作業を追加できるでしょうか?
次のバッチを実行する前にリクエストのバッチが完全に終了するのを待つのではなく、実行エンジンはトランスフォーマー モデルの単一のフォワード パスを実行し (すべてのアクティブなリクエストに対して 1 つのトークンを正確に生成します)、マイクロ秒間停止してキューを調べ、次のトークンに向けてバッチを再構築します。
もちろん、実稼働システムはさらに複雑です。これらは、優先順位、タイムアウト、チャンク プレフィル、投機的デコード、複数の GPU、テンソル並列処理、および公平性を扱います。しかし、このループはアイデアの形です。
スループットを向上させる単純な方法は、バッチを大きくすることです。それは機能しなくなるまで機能します。バッチを大きくするとスループットが向上しますが、最初のトークンが届くまでにユーザーの待ち時間が長くなる可能性もあります。チャットボットでは最初のトークンが非常に重要です

。ユーザーは時間の経過とともに長い応答のストリーミングを許容できますが、何かが表示されるまで待機しすぎると壊れたように感じます。したがって、LLM サービングには、単純な推論とは少し異なる目的があります。
最初のトークンまでの時間を破壊することなく
長いリクエストが短いリクエストをブロックしないようにする
実行中のバッチ処理は、反復ごとに有用な作業を含むアクティブなバッチの密度を維持することで、これらの制約を解決します。バッチを動的にすることで、静的バッチ処理のアイドル スロットとパディングを回避します。 LLM 自体が注目を集めていますが、それを経済的に実行可能にするのはスケジューラです。
TensorRT-LLM KV キャッシュと注意
NVIDIA Triton 動的バッチング ガイド
Orca: 生成深層学習モデルのインフライトバッチ処理のための分散サービスシステム

## Original Extract

In-flight batching is the trick that keeps LLM serving from wasting GPU seats.

LLM Serving and the Bus That Never Stops | Rafi Hasan Rafi Hasan Blog
Search ⌘ K ← BACK April 2, 2026 • 8 min read LLM Serving and the Bus That Never Stops
In-flight batching is the trick that keeps LLM serving from wasting GPU seats.
I used to think batching requests for a machine learning model was a solved problem. I have hosted models and served requests. Batching speeds up serving requests. If you don't batch, for example, send one tiny request at a time, the GPU behaves like a giant bus carrying one passenger. It can move, but the economics are terrible. This is why batching exists. But LLMs make batching weird.
In a traditional web backend or a standard computer vision pipeline, it's straightforward. You put requests into a queue, wait until you hit a batch size of 4 or 8, slam them into the GPU, and return the results. Standard engineering. A single trip through the model.
input -> model -> output
An image classifier does this. An embedding model does this. You pass in data, the model runs, and you get the result. But LLM serving completely breaks this mental model. LLM generation is iterative. It generates one token at a time.
prompt -> prefill
token 1 -> decode
token 2 -> decode
token 3 -> decode
...
This means serving an LLM is not just "run the model once." It is a scheduling problem that repeats every token. If you treat LLM requests like traditional web requests, your GPU efficiency plummets, your latency spikes, and your cloud bill skyrockets.
LLM serving is a loop. Every token is another chance to waste the GPU or fill it.
Prefill, Decode, and the KV Cache
Each request has two phases. The first phase is prefill , where the model reads the prompt and builds the internal attention state. The second phase is decode , which uses autoregressive decoding to generate text one token at a time, feeding each generated token back into the model to predict the next.
Because decode steps run repeatedly until an end condition is met, requests vary wildly in duration. To avoid recomputing the prompt history at every step, the server maintains a KV cache in GPU memory. The scheduler's goal is to keep the GPU busy with token generation without running out of this finite cache memory.
Imagine three requests arrive together.
A -> needs 8 output tokens
B -> needs 2 output tokens
C -> needs 6 output tokens
Static batching puts them on the same bus and makes the bus finish the whole trip before taking new passengers.
Step 1: A B C
Step 2: A B C # B is done
Step 3: A _ C
Step 4: A _ C
Step 5: A _ C
Step 6: A _ _ # C is done
Step 7: A _ _
Step 8: A _ _ # A is done
Even though B and C finished early, their seats cannot be reassigned. The GPU keeps running, but it wastes memory and computes empty padding tokens. That is the waste.
The Fixed Tour Bus vs. The Dynamic City Transit Bus
I like thinking about the difference between static and in-flight batching as the difference between a pre-booked tour bus and a public city transit bus:
The Fixed Tour Bus (Static Batching): A tour bus leaves the station with a set passenger list. Even if a passenger decides to get off early at stop 2, their seat must remain empty for the rest of the trip. The bus cannot pick up new passengers on the road. Instead, it must complete the entire tour and return to the station before loading a new group.
The Dynamic City Transit Bus (In-Flight Batching): A transit bus that runs a continuous loop. As soon as a passenger reaches their destination and steps off, the bus pauses briefly at the next stop, lets a new passenger board to fill the empty seat, and immediately continues its journey.
In LLM serving, the bus is the active batch. A seat is not just a "batch slot." It represents GPU memory and KV cache capacity. Getting off means a request has hit an end condition. Boarding means a new request has enough memory budget to join the active generation loop.
The bus that changes passengers while moving
Each seat is a batch slot backed by GPU memory and KV cache. Each tick is one generation iteration.
The important part is that the batch is no longer a fixed group of requests. It is a dynamic, fluid variable.
In-flight batching is also called continuous batching or iteration-level batching. So how do engines like vLLM or TensorRT-LLM actually implement this "dynamic bus" in code? They shift the scheduling boundary from the request level to the iteration level.
At every generation iteration, the scheduler asks:
Which requests are still active?
Which new requests are waiting?
Is there enough KV cache space?
Can we add new work without hurting latency too much?
Instead of waiting for a batch of requests to finish entirely before running the next batch, the execution engine runs a single forward pass of the transformer model (generating exactly one token for all active requests), pauses for a microsecond to look at the queue, and rebuilds the batch for the very next token.
Of course production systems are more complicated. They deal with priorities, timeouts, chunked prefill, speculative decoding, multiple GPUs, tensor parallelism, and fairness. But this loop is the shape of the idea.
The naive way to improve throughput is to make the batch larger. That works until it does not. Bigger batches can increase throughput, but they can also make users wait longer before the first token. In a chatbot, the first token matters a lot. A user can tolerate a long answer streaming over time, but waiting too long before anything appears feels broken. So LLM serving has a slightly different objective than plain inference:
without destroying time to first token
without letting long requests block short ones
In-flight batching resolves these constraints by keeping the active batch dense with useful work on every iteration. By making the batch dynamic, we avoid the idle slots and padding of static batching. The LLM itself gets all the attention, but the scheduler is what makes it economically viable.
TensorRT-LLM KV cache and attention
NVIDIA Triton dynamic batching guide
Orca: A Distributed Serving System for In-Flight Batching of Generative Deep Learning Models
