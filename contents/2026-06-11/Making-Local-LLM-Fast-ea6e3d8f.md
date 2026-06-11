---
source: "https://bogdan.nimblex.net/programming/2026/06/10/making-local-llm-fast.html"
hn_url: "https://news.ycombinator.com/item?id=48489344"
title: "Making Local LLM Fast"
article_title: "Bogdan's Ramblings - Technology blog"
author: "bogdan_r"
captured_at: "2026-06-11T13:29:01Z"
capture_tool: "hn-digest"
hn_id: 48489344
score: 1
comments: 0
posted_at: "2026-06-11T12:20:07Z"
tags:
  - hacker-news
  - translated
---

# Making Local LLM Fast

- HN: [48489344](https://news.ycombinator.com/item?id=48489344)
- Source: [bogdan.nimblex.net](https://bogdan.nimblex.net/programming/2026/06/10/making-local-llm-fast.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T12:20:07Z

## Translation

タイトル: ローカル LLM を高速化する
記事のタイトル: Bogdan の放浪記 - テクノロジー ブログ
説明: ボグダン

記事本文:
ホーム
購読する
ローカル LLM を高速化する
私は Fono というツールを作成しました。これは、自分のマシン上で完全に実行できる、コンピュータ用の音声フロントエンドです。クラウドもアカウントもラップトップから音声も出ません。 2 つのホットキーで 3 つのジョブがあります。
F7 キーを押し続けると、入力していたウィンドウにテキストが入力され、クリーンアップされたテキストが表示されます。
F8 キーを押したままアシスタントに話しかけると、さまざまな考えに答えてくれます。また、単に返信するだけでなく、実際に何かを行うためのツール (画面上の内容を確認するなど) を呼び出すこともできます。
コンピューターに音声を与えます。コーディング エージェントや他のツールは、それを通じてあなたに話しかけることができます。
最初の 2 つのジョブは、ローカル言語モデルに依存できます。問題は、ローカル モデルは発言する前に考える必要があり、GPU のラック上で実行されるクラウド API とは異なり遅く感じられることです。私のラップトップでは、アシスタントの最初のターンから一言が返されるまでにほぼ 3 秒かかりました。会話の 6 ターン目までに、会話は 7 に近づきました。それは1000ミリ秒以内に死を意味します。完璧に機能している場合でも、全体が壊れているように感じます。
この投稿は、最初の単語を約 3 分の 1 秒で発声するというストーリーです。実際に何が時間のかかるのかを説明し、私が作ったバグ (これは良いものです) を示し、すべてを自分のハードウェアで再現するための正確なコマンドを提供します。私は「信じてください」ベンチマークはしません:)
プロンプトをローカル モデルに送信すると、作業は 2 つのまったく異なるフェーズに分割されます。
Prefill はプロンプトを読み取るモデルです。すべてのトークン (テキストのすべてのチャンク) がニューラル ネットワークを介してプッシュされるため、モデルは応答する前にコンテキストを「理解」します。
デコードは模範解答です。あなたはこれが起こるのを一字一句見ていきます。
ユーザーが感じる数字が最初のトークンまでの時間です。それが私とのギャップです

ホットキーを解除すると、最初の単語が表示されます。 (Fono ではもう少し複雑ですが、これでいきます。) このギャップは、ほぼ完全に事前に埋め込まれています。デコードは、回答がどのくらいの速さでストリーミングされるかを制御しますが、プリフィルは、何かが起こっているのではないかとそこに座って疑問に思う原因になります。
注意しないと、ローカルのチャット アシスタントが何度も事前入力料金を請求することになります。毎ターン、モデルはこれまでの会話全体、つまりシステムの指示、以前のすべての質問、以前のすべての回答を再読み取りします。新しい文を 1 つ追加するだけですべてが機能します。 6 ターン目では、1 ターン目から 5 ターン目までの料金が再度支払われます。
迂回: ストレス ハードウェアをプリフィルおよびデコードする方法
非常に興味深い内容なので、2 分間お話ししましょう。コンピューターのさまざまな部分での事前入力とデコードのボトルネック。
プレフィルはコンピューティングに依存します。すべてのプロンプトのトークンは、行列乗算の 1 つの大きな並列バッチで一緒にネットワークを通過できます。これにより、すべてのコアとすべての SIMD レーン (クロックごとに多数の乗算と加算を行う AVX/NEON 命令) がビジー状態に保たれます。より多くのコアと幅広いベクトル命令により、プリフィルが高速化されます。このため、CPU に適した llama.cpp のビルドが多少重要になります。
デコードはメモリ帯域幅の制限を受けます。 1 つのトークンを生成するには、モデルは基本的にすべての重みを RAM から読み取る必要があります。 4 ビットの 40 億パラメータのモデルは、およそ 2 ～ 3 GB です。 100 個のトークンを生成するということは、それらのギガバイトをメモリから 100 回ストリーミングすることを意味します。コアは、その時間のほとんどをコンピューティングではなくデータの待機に費やします。これがローカル LLM 速度の秘密です。デコードは通常、生の計算ではなく、モデルを RAM から移動できる速度によって制限されます。これは、量子化 (重みの縮小) によってデコードが高速化される理由、および高速メモリが重要である理由でもあります。
もちろん注意点もたくさんありますが、

これをスリムに保ちます:)
この投稿の残りの部分は 2 つの要点で構成されています。まず、レイテンシーはプレフィルによって支配されており、プレフィルのコストは「どれだけのプロンプトを読み取らなければならないか」であるため、勝利への方法は、プロンプトの読み取りを減らすことです。次に、毎ターン再読するプロンプトの部分は、まさに変更されない部分です。私たちはそれをまったく読むべきではありません。
コツ: すでに読んだものを読むのをやめる
モデルのプロンプトの「理解」は、事前入力後に破棄されません。これは、KV キャッシュ (キー/値キャッシュ、ここでは名前は関係ありません) と呼ばれるメモリのチャンク内に存在します。これは、これまでに読み取ったすべてのモデルの作業記憶であると考えてください。
ローカル モデルを使用可能にする洞察: このターンのプロンプトの開始が前のターンのプロンプトと同一であれば、その部分に対するモデルのワーキング メモリも同一です。したがって、再計算せずに再利用してください。その状態のスナップショットを作成し、次のターンでは最初から事前に入力するのではなく、スナップショットを復元します。復元は基本的にメモリのコピーです。 Fono では、サイズに関係なく 15 ～ 40 ミリ秒かかります。同じコンテンツをコールドで事前入力するには数秒かかる場合があります (上のアニメーションの右側のレーン)。それがゲーム全体だ。
ただし、鉄則が 1 つあります。それは私が足を撃った場所です。
再利用は前面からのみ機能します
キャッシュの再利用は、最初のトークンとバイトごとに同一であるプレフィックスに対してのみ機能します。プロンプトの早い段階で何かが変更された瞬間、その後のすべてのトークンを再計算する必要があります。スプレッドシートのようなものです。セル A1 を変更すると、以下のすべての数式が再計算されます。
したがって、物を入れる順序によって、どれだけ再利用できるかが完全に決まります。決して変更されないもの (システム プロンプト、ツールの説明) は先頭にあり、一度キャッシュして永久に再利用できる安定したプレフィックスを形成する必要があります。変化するもの

毎ターン (ユーザーが今言った新しいこと) は後ろに進みます。
いくつかの愚かなバグがありました。モデルはある時点で 2 回ロードされました。その他、もう少し複雑な場合もあります。一緒に笑えるように、簡単な方法を共有します。
Fono には、デフォルトのローカル モデルとして Google の Gemma が同梱されるようになりました。 Gemma のプロンプト形式には、一部のモデルのようなシステム命令用の専用スロットがないため、最初に実装したときは、システム プロンプトを現在のユーザー ターンに貼り付けました。
【会話履歴】
<start_of_turn>ユーザー
{大きなシステム全体のプロンプト + ツール}
ユーザーリクエスト: {今言ったこと}
接頭辞のルールを念頭に置いて読んでください。 1 ターン目では履歴がなく、システム プロンプトが前面に表示されるため問題ありません。しかし、2 ターン目になると、会話履歴がシステム プロンプトの前にスライドします。大きな不変のブロックは、位置 0 ではなくなりました。プレフィックスが壊れています。キャッシュがミスします。モデルはすべてを 1 ターンごとに再読み込みします。
私はスナップショットと復元を行うマシン全体を構築し、最初のターン以降は基本的に起動しないようにプロンプ​​トをレイアウトしました。
システム プロンプトを適切な場所に置きます: first 。
現在、プロンプトは追加専用です。各ターンは末尾に追加するだけであり、新しい行の前のすべてはバイトごとに安定したプレフィックスです。システム プロンプトは起動時に一度キャッシュされ、会話のたびに再利用されます。各ターンでモデルが実際に読み取るのは、あなたが話したばかりの新しい文だけです。復元されたスナップショットの上にある数十のトークン。
これを回帰テストでラップし、3 ターンの会話を実行し、各ターンのプロンプトが次のターンの正確な文字列プレフィックスであることを確認します。誰かがプロンプトの順序を変更し、追加専用プロパティを再度破ると、テストは大声で失敗します。私はこの間違いを二度繰り返さないと自分自身を信じていません。

その後、テンプレートがオフだったのでガベージをキャッシュしていましたが、それでもキャッシュは機能しました:)
2 つのパスとキャッシュの再利用方法
追加専用のレイアウトにより、両方のホットキーがプロンプトの前面に安定したベースを提供します。 Fono は起動時に一度それを事前に入力し、スナップショットを撮って固定するので、決して追い出されることはありません。 F7 (ディクテーション) の場合、ベースはクリーンアップ手順と個人辞書です。 F8 (アシスタント) の場合、システム プロンプトとツールの説明が表示されます。他のものはすべて上に積み重ねられます。
以下のボタンを操作して、その仕組みを直感的に感じてください。
F7 がファンアウトします。ターミナルのテキストはチャット ボックスのテキストとは読み方が異なるため、ディクテーションではどのアプリを使用しているかが考慮されます。したがって、各アプリは独自のスナップショットを取得し、すべて同じ固定ベースから分岐します。初めてエディターに書き込むと、Fano はベースを復元し、その上の小さなエディター コンテキストをデコードして、そのブランチを保存します。その後、エディターにディクテーションを入力するたびに、エディターのスナップショット全体が復元され、単語のみがデコードされます。実際に発言した内容はキャッシュされることはありません。再利用可能なプレフィックスのみが保持されます。
F8 はチェーンを構築します。アシスタントは会話なので、各ターンは前のターンに積み重ねられます。ターン 1 では、ピン留めされたベースが復元され、最初のリクエストと応答がデコードされ、結果がターン 1 のスナップショットとして保存されます。ターン 2 はターン 1 を復元し、新しい交換のみをデコードします。会話の前半部分が再読されることはありません。ターンのコストは、話し続けた時間ではなく、今言ったことを追跡します。会話は 5 分間アイドル状態になると期限切れになり、チェーンが固定されたベースに切り戻されます。
どちらの形状も同じ 2 つのアイデアを使用しています。高価な作業を 1 回だけ支払う固定ベース、および最長プレフィックスの復元: まだ一致する最も深いスナップショットから開始し、デルタのみをデコードします。
同じラップトップ、

ame モデル (Gemma 4 E2B、4 ビット量子化)、同じ 6 ターンの会話。まずは見出し。キャッシュが機能している場合と機能していない場合に、順番に最初の単語を取得するまでの時間。 「not」行は、まさにシステムインザテールのバグが生み出したものです。灰色の線はオラマです。これについては後ほど説明します。
また、増加するプレフィルを置き換えるチェックポイント サイズと復元時間を含むテーブルと同じデータ:
キャッシュされていない列は、会話が大きくなるにつれて増加します。各ターンでスナップショット (15 ～ 40 ミリ秒) が復元され、新しい文のみが読み取られるため、キャッシュされた列はフラットのままです。
効果を見逃すことがないようにするために、キャッシュされたプレフィックスのサイズも直接スイープしました。アシスタントがシステム プロンプトに大量のツール定義を持っていると想像してください。以下は、プレフィックスが小さなトークンから約 3,300 トークンまで増加するときの、コールド プレフィルとキャッシュされた復元の最初のトークンのレイテンシです。
最後の行を 2 回読み取ります。この最新の CPU では、3,251 トークンのシステム プロンプトのコールド プレフィルに 45 秒かかります。キャッシュされた復元では最初のトークンが 133 ミリ秒で提供され、コストは RAM 内にある 60 MB のスナップショットです。少ないメモリで多くのレイテンシを節約できます。
これにより、将来的にはローカル ツールの使用が可能になります。すでにロードマップに載っています:)
レイテンシとメモリを交換することは、メモリが使い果たされない場合にのみ有効です。スケールは見た目より小さいです。 KV キャッシュはトークンごとに約 18 KB で実行されるため、ユーザーとアシスタントの 1 回の交換で約 1 メガバイトが追加されます。大きなスナップショットは履歴ではなくプレフィックスから取得されます。ファット ツール定義ブロックは、それ自体で 60 MB です。
キャッシュを抑制するには次の 3 つのことが必要です。
剪定。会話が成長するにつれて、新しいチェックポイントには以前のチェックポイントが含まれるため、古いチェックポイントはその場で削除されます。短い会話は、ターンごとに 1 つではなく、1 つのライブ チェックポイントに配置されます。
ローリングウィンドウ。アシスタントのメモリは 12 ターンに制限されています

5 分間のアクティビティを開始します。ウィンドウがスライドし始めると、最も古いターンが前面から脱落し、古いチェックポイントはもう剪定できなくなり、下のハードバウンドの下に落ちます。 5分以上放置すると履歴が消去されるので、次のターンでベースだけが回復します。それが最も安価なケースです。
ハードバウンド。それ以外のすべて (他のアプリのディクテーション コンテキスト、1 回限りのウィンドウの取得、長い会話の残り) は、8 つのチェックポイントと合計 256 MB に制限されます。制限を超えると、最も最近使用されていないエントリが削除されます。固定された基地は決して立ち退きされません。
1 つの微妙な点で、これらすべての削除を安全に行うことができます。各スナップショットは、モデル状態の完全なスタンドアロン コピーであり、前のスナップショットへのリンクではありません。チェーンは、スナップショットを安価に構築する方法にすぎません。一度保存すると、他のスナップショットは参照されないため、1 つのスナップショットを削除しても別のスナップショットが破損することはありません。コストはある程度の冗長性 (ベース プレフィックスは、その上に構築されるすべてのスナップショット内で複製されます) であり、まさにそれがバイト バジェットとプルーニングが存在する理由です。
しかし、実際にオラマよりも速いのでしょうか?
チャートで嘘をつくのは簡単なので、ここは正直に言わなければなりません。
ollam は、ローカル モデルを実行する一般的な方法です。 Fono が使用しているのと同じエンジン (llama.cpp) 上に構築されています。まったく同じモデルファイルをollamにインポートして実行しました

[切り捨てられた]

## Original Extract

Bogdan

Home
Subscribe
Making local LLM fast
I built a tool called Fono . It’s a voice front-end for your computer that can run entirely on your own machine. No cloud, no account, no audio leaving the laptop. It has three jobs, on two hotkeys:
Hold F7 to dictate and cleaned-up text lands in whatever window you were typing in.
Hold F8 to talk to an assistant and it answers various thinkgs. It can also call tools (look at what’s on your screen, and more to come) to actually do things, not just reply.
Give your computer a voice. Coding agents or other tools can speak to you through it.
The first two jobs can lean on a local language model. The problem is that a local model has to think before it speaks, and that feels slow in a way a cloud API running on a rack of GPUs does not. On my laptop, the very first assistant turn was taking almost three seconds before a single word came back. By the sixth turn of a conversation it was closer to seven. That’s death by a thousand milliseconds. It makes the whole thing feel broken even when it’s working perfectly.
This post is the story of getting that first word out in about a third of a second instead. I will explain what actually costs the time, show you the bug I made (it’s a good one) and give you the exact commands to reproduce everything on your own hardware. I don’t do “trust me bro” benchmarks :)
When you send a prompt to a local model, the work splits into two very different phases.
Prefill is the model reading your prompt. Every token (every chunk of text) gets pushed through the neural network so the model “understands” the context before it responds.
Decode is the model answering. You watch this happen, word by word.
The number a user feels is the time to first token . That’s the gap between letting go of the hotkey and the first word appearing. (In Fono it’s a bit more complex, but we’ll go with this.) That gap is almost entirely prefill. Decode controls how fast the answer then streams out, but prefill is what makes you sit there wondering if anything is happening.
If you are not careful, a local chat assistant makes you pay prefill over and over. Every turn, the model re-reads the entire conversation so far: the system instructions, every previous question, every previous answer. All that work just to append one new sentence. Turn six pays for turns one through five all over again.
Detour: how prefill and decode stress hardware
Let’s spend 2 minutes on this because I find it very interesting. Prefill and decode bottleneck on different parts of your computer.
Prefill is compute-bound. All the prompt’s tokens can go through the network together, in one big parallel batch of matrix multiplications. That keeps every core and every SIMD lane (the AVX/NEON instructions that do many multiply-adds per clock) busy. More cores and wider vector instructions make prefill faster. This is why the right build of llama.cpp for your CPU matters a little bit.
Decode is memory-bandwidth-bound. To produce one token, the model has to read essentially all of its weights out of RAM. A 4-billion-parameter model at 4 bits is roughly 2-3 GB. Generating 100 tokens means streaming those gigabytes from memory a hundred times over. Your cores spend most of that time waiting for data, not computing. This is the dirty secret of local LLM speed. Decode is usually limited by how fast you can move the model out of RAM, not by raw math. It’s also why quantization (shrinking the weights) speeds decoding up and why fast memory matters.
Of course, there are many caveats here, but let’s keep this lean :)
Two takeaways set up the rest of the post. First, since the latency you feel is dominated by prefill, and prefill cost is “how much prompt do I have to read”, the way to win is to read less prompt . Second, the part of the prompt we re-read every turn is exactly the part that doesn’t change. We shouldn’t be reading it at all.
The trick: stop reading what you already read
The model’s “understanding” of the prompt isn’t thrown away after prefill. It lives in a chunk of memory called the KV cache (key/value cache, the name doesn’t matter here). Think of it as the model’s working memory of everything it has read so far.
The insight that makes local models usable: if the start of this turn’s prompt is identical to last turn’s, the model’s working memory for that part is identical too. So don’t recompute it, reuse it. Snapshot that state and next turn restore the snapshot instead of prefilling from scratch. Restoring is basically a memory copy. In Fono it takes 15-40 milliseconds regardless of size. Prefilling the same content cold can take seconds (the right lane in the animation above). That’s the whole game.
There is one ironclad rule though, and it’s where I shot myself in the foot.
Reuse only works from the front
Cache reuse only works for a prefix that is byte-for-byte identical from the very first token . The moment something changes early in the prompt, every token after it must be recomputed. It’s like a spreadsheet. Change cell A1 and every formula below recalculates.
So the order you put things in completely decides how much you get to reuse. The stuff that never changes (system prompt, tool descriptions) wants to be at the front , forming a stable prefix you cache once and reuse forever. The stuff that changes every turn (the new thing the user just said) goes at the back .
There were several dumb bugs. The model got loaded twice at some point. And other where a bit more complex. I’ll share an easy one so we can laugh together.
Fono now ships with Google’s Gemma as the default local model. Gemma’s prompt format has no dedicated slot for system instructions like some models have, so when I first implemented it I glued the system prompt onto the current user turn:
[ conversation history ]
<start_of_turn>user
{the entire big system prompt + tools}
User request: {what you just said}
Read that with the prefix rule in mind. On turn one it’s fine because there is no history and the system prompt sits at the front. But on turn two, the conversation history slides in front of the system prompt. The big unchanging block is no longer at position zero. The prefix is broken. The cache misses. The model re-reads everything, every single turn.
I built the entire snapshot-and-restore machine and then laid out the prompt so it could basically never fire after the first turn.
Put the system prompt where it belongs: first .
Now the prompt is append-only . Each turn only adds to the end, and everything before the new line is a byte-for-byte stable prefix. The system prompt is cached once, at startup, and reused on every turn of every conversation. Each turn, the only thing the model actually reads is the new sentence you just spoke. A couple of dozen tokens on top of a restored snapshot.
I wrapped this in regression tests that walk a three-turn conversation and assert each turn’s prompt is an exact string prefix of the next turn’s. If anyone reorders the prompt and breaks the append-only property again, the tests fail loudly. I don’t trust myself to not make this mistake twice.
Then I was caching garbage because the template was off but still, caching worked :)
The two paths, and how they reuse the cache
The append-only layout gives both hotkeys a stable base at the front of the prompt. Fono prefills it once at startup, snapshots it and pins it so it’s never evicted. For F7 (dictation) the base is your cleanup instructions plus a personal dictionary. For F8 (the assistant) it’s the system prompt plus the tool descriptions. Everything else stacks on top.
Play with the buttons below to get an intuitive feel how it works.
F7 fans out. Dictation cares which app you’re in because text for a terminal reads differently than text for a chat box. So each app gets its own snapshot and they all branch off the same pinned base. The first time you dictate into your editor, Fono restores the base, decodes the small editor context on top and saves that branch. Every dictation into the editor after that restores the editor snapshot whole and decodes nothing but your words. What you actually said is never cached. Only the reusable prefix is kept.
F8 builds a chain. The assistant is a conversation, so each turn is stacked on the previous one. Turn one restores the pinned base, decodes your first request and the reply, and saves the result as the turn-one snapshot. Turn two restores turn one and decodes only the new exchange. Nothing earlier in the conversation is ever re-read. The cost of a turn tracks what you just said, not how long you’ve been talking. Conversations expire after five idle minutes, which trims the chain back to the pinned base.
Both shapes use the same two ideas. A pinned base that pays the expensive work only once, and longest-prefix restore: start from the deepest snapshot that still matches and decode only the delta.
Same laptop, same model (Gemma 4 E2B, 4-bit quantized), same six-turn conversation. First the headline. Time to first word, turn by turn, with the cache working versus not. The “not” line is exactly what my system-in-the-tail bug produced. The gray line is ollama, which we get to in a bit.
And the same data as a table, with the checkpoint size and the restore time that replaces the growing prefill:
The uncached column climbs as the conversation grows. The cached column stays flat because each turn restores a snapshot (15-40 ms) and only reads your new sentence.
To make the effect impossible to miss, I also swept the size of the cached prefix directly. Imagine an assistant with a big pile of tool definitions in its system prompt. Here is the first-token latency as that prefix grows from tiny to about 3,300 tokens, cold prefill versus cached restore:
Read the last row twice. A cold prefill of a 3,251-token system prompt takes 45 seconds on this modern CPU. The cached restore serves the first token in 133 ms and the cost is a 60 MB snapshot sitting in RAM. A little memory for a lot of saved latency.
This is what makes local tool use feasible in the future. It’s already on the roadmap :)
Trading memory for latency is only a good trade if the memory can’t run away. The scale is smaller than it looks. The KV cache runs about 18 KB per token , so one user-and-assistant exchange adds about a megabyte. The big snapshots come from prefixes, not history. A fat tool-definition block is 60 MB on its own.
Three things keep the cache in check:
Pruning. While a conversation grows, each new checkpoint contains the previous one, so the old one is deleted on the spot. A short conversation sits at one live checkpoint, not one per turn.
A rolling window. Assistant memory is capped at 12 turns within a 5-minute window of activity. Once the window starts sliding, the oldest turn drops off the front, old checkpoints can’t be pruned anymore and they fall under the hard bound below. Go idle for more than 5 minutes and the history is wiped, so the next turn restores just the base. That’s the cheapest case there is.
A hard bound. Everything else (other apps’ dictation contexts, one-off window grabs, leftovers from long conversations) is capped at 8 checkpoints and 256 MB total . When a limit is crossed, the least-recently-used entry is evicted. The pinned bases are never evicted.
One subtlety makes all this deleting safe. Each snapshot is a complete, standalone copy of the model state, not a link to the previous one. The chain is only how a snapshot gets built cheaply. Once saved, it refers to nothing else, so removing one snapshot can never corrupt another. The cost is some redundancy (the base prefix is duplicated inside every snapshot built on top of it) and that is exactly why the byte budget and the pruning exist.
But is it actually faster than ollama?
Here is where I have to be honest, because it’s easy to lie with a chart.
ollama is the popular way to run local models. It’s built on the same engine (llama.cpp) that Fono uses. I imported the exact same model file into ollama and ran

[truncated]
