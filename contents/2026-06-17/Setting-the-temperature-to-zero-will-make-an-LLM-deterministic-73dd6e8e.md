---
source: "https://www.zansara.dev/posts/2026-03-24-temp-0-llm/"
hn_url: "https://news.ycombinator.com/item?id=48571300"
title: "Setting the temperature to zero will make an LLM deterministic?"
article_title: "Setting the temperature to zero will make an LLM deterministic? · Sara Zan"
author: "zansara"
captured_at: "2026-06-17T16:22:55Z"
capture_tool: "hn-digest"
hn_id: 48571300
score: 2
comments: 0
posted_at: "2026-06-17T14:47:17Z"
tags:
  - hacker-news
  - translated
---

# Setting the temperature to zero will make an LLM deterministic?

- HN: [48571300](https://news.ycombinator.com/item?id=48571300)
- Source: [www.zansara.dev](https://www.zansara.dev/posts/2026-03-24-temp-0-llm/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T14:47:17Z

## Translation

タイトル: 温度をゼロに設定すると、LLM は決定的になりますか?
記事のタイトル: 温度をゼロに設定すると、LLM は決定的になりますか?・サラ・ザン
説明: サラ・ザン

記事本文:
サラ・ザンのブログ
について
投稿
プロジェクト
出版物
会談
温度をゼロに設定すると、LLM が決定的になりますか?
これは、仕事中に受け取った質問に答える一連の短いブログ投稿のエピソード 8 です。彼らは、さまざまな生成 AI テクノロジーに関するよくある誤解や疑問について話し合います。シリーズ全体はここで見つけることができます: 実践的な質問 。
LLM の「温度」パラメータの一般的な説明の 1 つは、それが答えの「ランダム性」を表すというものです。
それはおおむね正しいです。温度は LLM の最終デコード ステップのパラメータであり、Transformer アーキテクチャ全体の中で設計によりランダム性が実際に組み込まれている唯一のパラメータです。この段階で、モデルが次のトークン候補のロジットを計算したら、それらの値をリストの実際のトークンにマッピングする必要があります。通常、LLM は、必ずしも単一の最適なトークンを選択する必要はなく、N 個の最適なトークンの中からランダムに選択できる場合に最高のパフォーマンスを発揮します。N のサイズは、多かれ少なかれ、温度パラメーターが表すものとなります。
したがって、温度を 0 に設定すると、LLM はランダムな選択をせずに、常に最適な次のトークンを選択する必要があります。つまり、入力が固定され、アーキテクチャ内のランダム性の唯一の原因が削除されている場合、出力は常に同一になるはずです...そうですよね?
しかし、実際にはそうではないことがよくあります。同じモデル、同じパラメータ、温度 0 を使用して同じプロンプトを 2 回実行すると、遅かれ早かれ出力は少し異なるものになります。通常はそれほど多くはありません。それはたった一言から始まるかもしれません。その後、文は少し違った方向に進み、最終的に残りの部分が消え去ります。
LLM が単なる数学関数であると仮定すると、実際に温度 = 0 によってデコードが決定的になるはずです。

各ステップで、モデルはロジットを出力し、argmax トークンを取得してコンテキストに追加し、これを繰り返します。問題は、実際の推論が超並列ハードウェア、通常は数学的に純粋な状態ではなく、可能な限り高速にしようとするサーバー上で浮動小数点演算を使用して実行されることです。
浮動小数点演算は、実数演算の近似にすぎません。特に、これは結合的ではありません。通常の数学では、(a + b) + c = a + (b + c) が常に成り立ちますが、浮動小数点数の場合、中間の各ステップが丸められるため、これら 2 つの式はわずかに異なる結果を生成する可能性があります。同じことが、ニューラル ネットワーク全体の行列の乗算、減算、累算にも当てはまります。演算の順序を変更すると、結果の最後の数ビットを変更できます。
通常、これらの違いはわずかで、多くの場合無関係ですが、この場合は影響を及ぼします。 2 つの候補となる次のトークンのロジットが非常に似ている場合、わずかな数値の違いによって順序が入れ替わる可能性があり、1 つのトークンが変更されると、次のデコード ステップが異なるプレフィックスで実行されるため、相違が増大します。サンプリング ルールは決定的ですが、ロジットを生成した計算が実行間で同一であるとは保証されません。
これは次のように考えることができます。サンプリング決定論はシステム決定論と同じではありません。
ただし、これは問題の一部にすぎません。同じデータを使用して GPU 上で同じ行列の乗算を繰り返し実行すると、常にビット単位で同一の結果が得られることに、すでに異議があるかもしれません。計算は浮動小数点演算で行われ、コンピュータの電源が入っている間、GPU 上で他のジョブが実行されているはずです。では、温度 = 0 での LLM サンプリングが決定的ではないのに、これらの計算はなぜ決定的なのでしょうか?
Thinkに関する最近の投稿で

ing Machines のブログ、Defeeting Nondeterminism in LLM Inference、Horace 彼はこの問題をさらに深く掘り下げています。浮動小数点演算が不完全であるというだけではありません。最新の推論システムでは、リクエストをまとめてバッチ処理する必要もあり、1 つのリクエストの結果は、それが実行されたバッチ コンテキストに依存する可能性があります。特定の正確なバッチの場合、順方向パスは決定的である可能性があります。しかし、ユーザーの観点から見ると、バッチ自体は実行ごとに安定していないため、システムは依然として非決定的です。プロンプトは同一である可能性がありますが、同時にバッチ処理される入力は同一ではありません。
これは、ローカル テストではプロンプトが安定しているように見えても、本番環境では不安定になる場合がある理由でもあります。モデルが突然よりクリエイティブになったのではなく、システム条件が変化したのです。温度=0 では、トークン選択ルールのみが決定的になります。推論システム全体が毎回まったく同じロジックを生成することは保証されません。
今日の LLM 推論の仕組み、特に大規模な仕組みでは、決定的な出力を保証できる条件を強制するためのオプションがあまり残されていません。ホスト型 API と自己ホスト型推論ではトレードオフのみが大きく異なります。
ランダム性を減らし、LLM 出力を再現可能にするために、固定シードの使用を推奨する人もいます。実際、一部のプロバイダーは固定シードを公開しています。たとえば、OpenAI はシード パラメータを文書化し、決定論的なサンプリングに最善の努力を払っていると述べていますが、決定論は保証されておらず、バックエンドの変更が依然として出力に影響を与える可能性があることを明示的に警告しています。 system_fingerprint フィールドは正確に存在するため、基礎となるサービス構成がいつ変更されたかを知ることができます。
固定シードの問題は、温度がゼロより高い場合に結果を再現するのに役立つことです。

、すでにゼロになっている場合ではありません。これは、固定シードがサンプリング ステップのランダム性を制御するためです。温度を 0 に設定することで、そのランダム性の原因がすでに除去されているため、最終的な結果は固定シードの有無にかかわらず同じになりますが、GPU やスタックの残りの部分から来る他の非決定性のソースは影響を受けません。
したがって、テスト、デモ、回帰チェックなど、ゼロ以外の温度での呼び出しで同じ結果を得ようとする場合には、固定シードを使用する価値があります。ただし、これらはサンプラーにのみ影響し、温度がゼロの場合には役に立たないことに留意する必要があります。
自己ホストする場合、ランダム性を大幅に減らす 1 つのオプションは、同時実行性を減らすか排除することです。
これが機能するのは、バッチ処理とスケジューリングが安定するという単純な理由からです。 vLLM の再現性ガイダンスには、デフォルトではそれ自体では再現性が保証されないと記載されています。オフライン モードでは、スケジュールを確定的にするためにマルチプロセッシングを無効にする必要があります。一方、オンライン モードでは、バッチ処理の影響を受けない出力が必要な場合は、バッチ不変性のサポートが必要です。また、vLLM はバッチ不変性を個別の機能として文書化しており、現在は特定のハードウェア サポートに依存していることを指摘しています。
これは、ニーズに応じて、いくつかの異なる構成を選択できることを意味します。
動的バッチ処理による共有オンライン サービング: 最速、安価、最も再現性が低い
孤立したワーカー / 同時ジョブがない: 速度が遅く、コストが高く、再現性が高い
特殊なバッチ不変のサービング パス: 再現性は向上しますが、ハードウェアと機能の制約があります
全体的なパターンとしては、スループットを最適化すればするほど、再現性が低下するということです。
キャッシュ自体は再現性の問題に正確に対処するわけではありませんが、多くのアプリケーションでは適切です。

同じ入力から同じ出力を生成する場合は、抽象化のレベルを設定します。ベンチマークや評価を実行している場合を除き、これは多くの場合、最も実行可能なオプションであるだけでなく、最も安価で、最もシンプルで、最速のオプションでもあります。
実際には、同じリクエストに対して同じ表示結果だけが必要な場合、最も信頼できる方法は、リクエストをまったく再生成しないことです。プロンプト、モデル ID、および関連パラメーターをキャッシュ キーに正規化し、最初の成功した応答を保存し、それを後続の同一のリクエストで提供します。もちろん、これによってモデルが決定的になるわけではありませんが、アプリケーションがインターフェイス境界で決定的になります。これは通常、アプリケーション ビルダーが必要とするものです。
キャッシュには、シードやスケジューラーのトリックに比べて非常に優れた利点もあります。推論スタック内の隠された実装の詳細に依存しません。
もちろん、キャッシュには限界があります。これはリクエストが繰り返される場合にのみ役立ちます。ツール呼び出し、タイムスタンプ、外部取得、または非表示のコンテキストによって、一見同一の 2 つのリクエストが実際には同一ではない場合には、扱いにくくなる可能性があります。それでも、これは通常、この問題に対する他のどの解決策よりもはるかに便利であり、ほとんどの実稼働システムにとって唯一実用的な解決策です。
LLM の非決定論に直面すると、それをバグのように扱い、それを排除しようとする反応がよくあります。ただし、LLM は、ある程度の非決定性が許容される場合にパフォーマンスが向上するため、LLM にはランダム性要素が組み込まれて設計されていることにも留意する必要があります。
それはわかります。アプリケーションのビジネス ロジックの中核に、これほど巨大でランダムなブラック ボックスを置くことを好む人はいません。しかし、出力からランダム性を取り除くことは、LLM の動作を管理する正しい方法ではありません。完全に決定的な出力が必要な場合は、LLM を使用してデータを設計することをお勧めします。

ecision ツリー (または必要に応じてより洗練されたモデル) を作成し、それをアプリケーションで使用します。
LLM 出力の処理は、むしろ検証の問題です。スキーマとバリデータを使用して、小さなテキストのずれによって下流のコードが破損しないようにします。スポットチェックの代わりに eval を使用します。一貫性が重要な場合、または数ドルを節約する必要がある場合はキャッシュします。言い換えれば、ランダム性をモデル自体から削除しようとするのではなく、システム境界でランダム性を処理します。

## Original Extract

Sara Zan

Sara Zan's Blog
About
Posts
Projects
Publications
Talks
Setting the temperature to zero will make an LLM deterministic?
This is episode 8 of a series of shorter blog posts answering questions I received during the course of my work. They discuss common misconceptions and doubts about various generative AI technologies. You can find the whole series here: Practical Questions .
One common explanation of the "temperature" parameter of LLMs is that it represents the "randomness" of the answer.
That's broadly correct. Temperature is a parameter of the LLM final decoding steps, and the only one in the whole Transformer architecture that truly incorporates some randomness by design. At this stage, once the model has calculated the logits of the next token candidates, it has to map those values to an actual token from a list. Normally, LLMs perform best when they’re allowed to pick not necessarily the single best token, but instead choose at random among the N best tokens: the size of N is, more or less, what the temperature parameter represents.
Therefore, when we set the temperature to 0, the LLM must always choose the best next token, without making random choices. So, if the input is fixed and we have removed the only source of randomness in the architecture, the outputs should always be identical... right?
And yet, in practice, they often are not. Run the same prompt twice, with the same model, the same parameters, and temperature 0, and sooner or later the output will be a bit different. Not by much, usually. It may start with just one word; then the sentence takes a slightly different spin, until eventually the rest of the completion drifts away.
If we pretend an LLM is just a mathematical function, temperature=0 should indeed make decoding deterministic. At each step, the model emits logits, we take the argmax token, append it to the context, and repeat. The problem is that real inference is performed with floating-point arithmetic on massively parallel hardware, usually on a server that is trying to be as fast as possible rather than mathematically pristine.
Floating-point arithmetic is only an approximation of real-number arithmetic. In particular, it is not associative : in ordinary math, (a + b) + c = a + (b + c) always holds, but with floating-point numbers those two expressions can produce slightly different results because each intermediate step is rounded. The same applies to matrix multiplications, reductions, and accumulations throughout a neural network. Change the order of operations, and you can change the last few bits of the result.
Usually, those differences are tiny and often irrelevant, but in this case they have an impact. If two candidate next tokens have very similar logits, a minute numerical difference can swap their order, and once one token changes, the next decoding step runs on a different prefix, so the divergence compounds. The sampling rule is deterministic, while the computation that produced the logits is not guaranteed to be identical across runs.
You can think of it this way: sampling determinism is not the same thing as system determinism .
However, this is only part of the problem. You may already be objecting that running the same matrix multiplication on a GPU with the same data repeatedly will always provide bitwise-identical results. The computations are done in floating-point arithmetic, and there are surely other jobs running on the GPU while your computer is on. So why are these calculations deterministic, while LLM sampling with temperature=0 is not?
In a recent post on Thinking Machines's blog, Defeating Nondeterminism in LLM Inference , Horace He's digs even deeper into the issue. It's not merely that floating-point arithmetic is imperfect. Modern inference systems also need to batch requests together, and the result for one request can depend on the batch context in which it was executed. For a given exact batch, the forward pass may be deterministic. But from the user's point of view, the system is still nondeterministic, because the batch itself is not stable from run to run. Your prompt may be identical, but the inputs that get batched together with yours are not.
This is also why a prompt can look stable in local testing and then become flaky in production: the model did not suddenly become more creative, it's the system conditions that changed. temperature=0 makes only the token selection rule deterministic. It does not guarantee that the entire inference system will produce exactly the same logits every time.
The way LLMs inference works today, especially at scale, doesn't leave us with many options to enforce the conditions that can guarantee deterministic outputs. There are only trade-offs, and they differ quite a lot between hosted APIs and self-hosted inference.
To reduce randomness and make LLM outputs reproducible, some people recommend using a fixed seed, and indeed some providers expose one. OpenAI, for example, documents a seed parameter and says it makes a best effort to sample deterministically, while explicitly warning that determinism is not guaranteed and that backend changes can still affect outputs. Their system_fingerprint field exists precisely so you can notice when the underlying serving configuration has changed.
The problem with fixed seeds is that they help reproduce results when the temperature is above zero, not when it's already zeroed out. That's because a fixed seed controls the randomness of the sampling step: by setting the temperature to zero, we are already removing that source of randomness, so the net result is identical with or without a fixed seed, while every other source of nondeterminism coming from the GPU and the rest of the stack is unaffected.
So fixed seeds are worth using when you are trying to get the same results for a call with non-zero temperature, such as for tests, demos, and regression checks. But you must keep in mind that they affect only the sampler, and they won't help you when temperature is zero.
If you self-host, one option to drastically reduce randomness is to reduce or eliminate concurrency.
This works for the simple reason that it stabilizes batching and scheduling. vLLM's reproducibility guidance says that by default it does not guarantee reproducibility on its own. In offline mode, you should disable multiprocessing to make scheduling deterministic, while in online mode, you need batch invariance support if you want outputs that are insensitive to batching. vLLM also documents batch invariance as a distinct feature and notes that it currently depends on specific hardware support.
This means that you can pick a few different configurations, depending on your needs:
shared online serving with dynamic batching: fastest, cheapest, least reproducible
isolated worker / no concurrent jobs: slower, more expensive, more reproducible
specialized batch-invariant serving paths: better reproducibility, but with hardware and feature constraints
The overall pattern is that the more you optimize for throughput, the more reproducibility suffers.
Caching doesn't exactly address the reproducibility issue per se, but in many applications it's the right level of abstraction if you want the same input to produce the same output. It's often not only the most viable option, but also the cheapest, simplest, and fastest, unless you're running a benchmark or an evaluation.
In practice, if you just need the same visible result for the same request, the most reliable method is not to regenerate it at all. Normalize the prompt, model ID, and relevant parameters into a cache key, store the first successful response, and serve that on subsequent identical requests. This does not make the model deterministic, of course, but it does make your application deterministic at the interface boundary, which is usually what application builders need.
Caching also has a very nice advantage over seeds and scheduler tricks: it does not depend on hidden implementation details inside the inference stack.
Of course, caching has limits. It only helps when requests repeat, and it can become awkward if tool calls, timestamps, external retrieval, or hidden context make two apparently identical requests not truly identical. Still, it is usually far more convenient than any other solution to this problem, and the only practical one for most production systems.
When faced with LLM nondeterminism, there's often the reaction to treat it like a bug and to try to eliminate that. However, you should also keep in mind that LLMs were designed with a randomness factor built-in for a reason: because they perform much better when they are allowed a slight degree of nondeterminism.
I get it: nobody likes having such a huge, random black box at the core of an application's business logic. But removing randomness from the outputs is not the right way to manage an LLM's behavior. If you need completely deterministic output, it is better to use the LLM to design a decision tree (or a more sophisticated model, if needed) and then use that in your application.
Handling LLM outputs is rather matter of validation. Use schemas and validators so small textual drift does not break downstream code. Use evals instead of spot-checking. Cache where consistency matters, or where you need to save a few bucks. In other words, handle the randomness at the system boundary rather than trying to remove it from the model itself.
