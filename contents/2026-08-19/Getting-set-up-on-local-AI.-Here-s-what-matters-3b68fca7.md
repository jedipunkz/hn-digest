---
source: "https://twitter.com/draginol/status/2089338779324543047"
hn_url: "https://news.ycombinator.com/item?id=49361834"
title: "Getting set up on local AI. Here's what matters"
article_title: "Brad Wardell on X: \"Here's how you get set up with a local AI on your PC or Mac. https://t.co/7Cn1GgTNIm\" / X"
image: "https://pbs.twimg.com/media/HP7SH9KXAAAgu9I.jpg"
author: "draginol"
captured_at: "2026-08-19T14:23:22Z"
capture_tool: "hn-digest"
hn_id: 49361834
score: 2
comments: 0
posted_at: "2026-08-19T14:07:17Z"
tags:
  - hacker-news
  - translated
---

# Getting set up on local AI. Here's what matters

- HN: [49361834](https://news.ycombinator.com/item?id=49361834)
- Source: [twitter.com](https://twitter.com/draginol/status/2089338779324543047)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T14:07:17Z

## Translation

タイトル: ローカル AI のセットアップ。ここが重要です
記事タイトル: Brad Wardell on X: 「PC または Mac でローカル AI をセットアップする方法は次のとおりです。https://t.co/7Cn1GgTNIm」 / X
説明: ローカル AI 入門書

記事本文:
Brad Wardell 氏、X について: 「PC または Mac でローカル AI をセットアップする方法は次のとおりです。https://t.co/7Cn1GgTNIm」 / X Post
Brad Wardell @draginol PC または Mac でローカル AI をセットアップする方法は次のとおりです。ローカル AI 入門書
ローカル AI については詳しく説明したので、その必要はありません。ローカルの PC または Mac 上で動作する AI を使用する方法です。雲はありません。
私のフィードはローカル AI のものでいっぱいです。この件については、「オオカミ少年」のような雰囲気が漂っています。なぜなら、10 万ドルのハードウェアで動作していたことが後になってわかるという主張を耳にするからです。そのために、私は何十年も Claude Max、Codex Pro、Grok Super Heavy Build を購読できました。
私が本当に知りたかったのは、ハードウェアで何ができるのかということです。本当の制限は何で、実際に何ができるのか。そして今、あなたはその旅から恩恵を受けることができます。そしてはい、私は全角ダッシュをよく使います。それらは私のものです。 AIはそれらを主張することはできません。
AI パラメーター: 楽しみと利益のための 2B、3B、9B、27B
「Llama 8B」または「Gemma 27B」というモデルを見ると、その数字はそのパラメーター数、つまりニューラル ネットワーク内で学習された重みの数 (10 億単位) です。パラメーターは、大まかに言うと、物事を認識し、それについて推論するモデルの能力です。
~1B 以下: 壮大な幻覚を伴うオートコンプリート。分類、段落の要約、簡単な書式設定タスクに適しています。
2B–4B: 実に会話的です。指示に従い、一般的な質問に答え、慎重な指示に従って簡単なツールの使用を行うことができます。これは電話と NPU のスイート スポットです。
7B–9B: 主力クラス。適切な一般知識があり、複数のステップの指示に従うことができ、簡単なエージェント作業 (ツールの呼び出し、結果の読み取り、応答) を処理できます。
13B–30B: ローカル モデルが「本物の」 AI のように感じられ始める場所。判断力が向上し、幻覚が減れば、自分の間違いから立ち直ることができる

タスクの途中。
70B+: フロンティアモデルの領域に近づいていますが、本格的なハードウェアが必要です。ラップトップではなく、128 GB のユニファイド メモリまたはマルチ GPU リグを備えた Mac Studio を考えてください。
パラメータ数は賢さだけでなく速度にも影響します。モデルが生成するすべてのトークンでは、基本的にすべてのパラメーターをメモリから読み取る必要があります。 27B モデルは 3B モデルよりも賢いだけでなく、ワードあたりの作業量も約 9 倍多くなります。そのことを念頭に置いてください。リアルタイムで何ができるのか、何をスケジュールする必要があるのか​​を考え始めると、それは大きな問題になります。
モデルは 16 ビット精度でトレーニングされます。パラメータあたり 16 ビットの 8B モデルは、重みだけで 16GB です。
グラフィック形式のようなものだと考えてください。 16 ビット精度はビットマップに似ています。もう、誰も .BMP を電子メールで送信しません。
量子化はニューラル ネットワークの非可逆圧縮です。各重みを 16 ビットではなく 8、5、または 4 ビットで保存します。つまり、BMP を取得して JPEG に変換するようなものです。 8 ビット量子化では、元の非圧縮バージョンと本質的に同じになります。 4 ビットでは、いくらか失われていますが、おそらくベンチマーク以外では問題になるほどで​​はありません。それを下回るとかなり悪化し始めます。一般に 4 ビットがスイート スポットです。
アルファベットスープ: Q4_0、Q4_K_M、Q8_0、IQ4_NL
量子化にもさまざまな方法があります。グラフィック形式が大量にあるのと同じように、これらのモデルを量子化する方法も多数あり、それぞれに独自のトレードオフがあります。
私の経験では、NVIDIA が好むもの、および Qualcomm などの NPU が好むであろう MLX (Mac ハードウェア用) に焦点を当てる傾向があります。
Q8_0: 8ビット。ほぼロスレスで、サイズは 4 ビットの 2 倍です。サイズが問題にならない小型モデルにのみ価値があります (Q8 の 0.6B モデルは 1GB 未満なので、なぜでしょうか)。
Q4_K_M: 最新の 4 ビット「K-quant」。ビットの賢い割り当て (重要なレイヤーの精度が向上)。

通常はギガバイトあたり最高の品質であり、CPU と GPU ではデフォルトの推奨事項になります。
Q4_0: オリジナルの最も単純な 4 ビット形式。 Q4_K_M よりも品質はわずかに劣りますが、そのプレーンなブロック レイアウトは、専用のハードウェア パスが構築されているものです。 ARM CPU および NPU では、高速カーネルは Q4_0 のみを話すため、Q4_0 は Q4_K_M よりも大幅に高速になることがよくあります。
IQ4_NL: 非線形コードブックを使用する新しい 4 ビット形式。品質は素晴らしいですが、ハードウェアのサポートが不安定です。
フォーマットはハードウェアに一致する必要があります。私たち自身のテストでは、同じ 4B モデルが CPU 上の Q4_K_M として 37 トークン/秒で実行されましたが、NPU にルーティングされると 14 トークンに落ちました。これは、NPU パスが K-quant を処理できず、低速パスに戻ったためです。同じモデル、同じコンピュータ、同じ「4 ビット」でも、2.5 倍の違いがあります。これは、人々が誤ってローカル AI セットアップをサンドバッグにしてしまう最も一般的な方法です。これが、依然として頭の痛い問題であり、Clairvoyance のようなプログラムの人気が爆発的に高まっている理由です。彼らは、このナンセンスに対処するだけです。
TOPS、NPU、GPU: 思考の速度
チップ ベンダーは、TOPS (1 秒あたり数兆回の演算) を宣伝しています。 Copilot+ 認定ラップトップの NPU は 40 ～ 80 TOPS を主張します。 RTX 4090 は 600 以上を実現します。45-TOPS NPU は 4090 の速度の何分の 1 かで AI を実行すると考えるかもしれません。
モデルが応答するときは、大きく異なる 2 つのフェーズがあります。
事前入力: プロンプトを読みます。これは並行して実行できるため、ここでは GPU と NPU がうまく機能します。
デコード: 答えを書きます。トークンは一度に 1 つずつ発行され、トークンごとにメモリを通じてモデル全体を再度ストリーミングする必要があります。 TOPS はほとんど無関係です。メモリ帯域幅がすべてです。
したがって、ラップトップの NPU は、記事全体のテキストを 0.5 秒で読んで理解できます。ただし、実際にコメントするにはメモリ帯域幅によって制限されており、1 分ほどかかる場合があります。
RAM: システム RAM、GPU VRAM、および u

統合された記憶
モデルは全体が高速メモリ内のどこかに存在する必要があります。
ディスクリート GPU (VRAM): これまでで最速のオプションですが、VRAM が不足しています。 12GB カードは 9B に快適に収まりますが、27B には全く入りません。システム RAM へのレイヤーのスピルは機能しますが、スピルされたすべてのレイヤーはシステム RAM の速度で実行されます。
システム RAM (CPU): 豊富で安価 (32GB は正常に実行できるものであれば何でも入ります) ですが、遅いです。以下を参照してください。
ユニファイド メモリ (Apple Silicon、Snapdragon X、AMD Strix Halo): CPU、GPU、および NPU は 1 つのプールを共有します。 128GB Mac の優れた点は、その帯域幅自体はそれほど顕著ではありませんが、128GB のすべてが適切な帯域幅でモデルで利用できることです。
サイズ設定の経験則: モデル ファイル サイズ + 作業コンテキストの 20 ～ 30%。 5GB Q4 モデルには、約 7GB の空き容量が必要です。
メモリ帯域幅: 無視されがちなボトルネック
メモリ帯域幅は、他のどの仕様よりもローカル AI デコード速度を正確に予測します。
一般的なラップトップ DDR5: ~60 ～ 90 GB/秒
Copilot+ 認定ラップトップ (LPDDR5X、共有): ~120 ～ 152 GB/秒
Snapdragon X2 Elite Extreme (LPDDR5X、共有): ~228 GB/秒
Nvidia DGX Spark (LPDDR5X、統合): ~273 GB/秒
Radeon RX 9070 XT (GDDR6): ~645 GB/秒
RTX 4090 (GDDR6X): ~1,008 GB/秒
RTX PRO 6000 ブラックウェル (GDDR7): ~1,792 GB/秒
これらの速度によって、タスクをリアルタイムで実行するか、スケジュールして実行するかの違いが生じます。
私が最初に気づいたことの 1 つは、私が選んだラップトップでは 27B モデルを問題なく実行できるということです。ただやるのが遅いだけです。しかし、AI を使用して行う必要がある作業のほとんどはリアルタイムではありません。私が Clairvoyance で最初に使用した機能はスケジュール設定でした。非常に多くのダッシュボード、クラッシュ レポート、センチメント レポート、販売データが入ってくるので、朝見るために一晩実行するだけです。以前はクロードにそれを持っていましたが、そのためにはトークンで月に100ドルかかりました。しかし、27B モデルならそれが可能です

私がスケジュールを立てている限り、費用はかかりません。
気晴らし: ハードウェアに興味を持つ
RTX 5090 は現在のコンシューマの王様です。ほぼ 1.8 TB/秒ということは、5 GB モデルが数百トークン/秒でデコードすることを意味し、その 32 GB の VRAM は 4 ビット 27 B に余裕を持って適合します。あなたの目標が「高速なローカル AI、お金は関係ありません」なら、これが答えです。この記事の執筆時点では、RTX 5090 の小売価格は 4 兆 6000 億ドル強です。
RTX PRO 6000 Blackwell は 5090 のワークステーションの兄弟です。バスは同じ ~1.8 TB/秒ですが、96 GB の VRAM を備えており、フル スピードで 1 枚のカードに 4 ビット 70B (またはギリギリの 8 ビット 70B) を保持するのに十分です。これは、2 台の 4090 では解決できない問題 (下記参照)、大きなモデルと 1 つのメモリ プールでの大きな帯域幅を、ほぼまともな中古車の価格で解決します。私たちの使用例は、これらの 1 つをラックに置き、複数のユーザーがその上で 27B モデルをリアルタイムで実行することです。
Nvidia DGX Spark は予想よりも遅いです。これは、卓上用の AI スーパーコンピューターとして販売されており、その処理能力に優れています。 128 GB のユニファイド メモリは、コンシューマ GPU が対応できないモデルに適合します。ただし、273 GB/秒の帯域幅は M5 Max の半分未満、4090 の 4 分の 1 です。大きなモデルも問題なく実行できます。どのモデルも高速に実行されません。速度ではなく容量を購入しているのです。これらがマイクロセンターに在庫されているのには理由があります。代わりに Snapdragon X2 Elite Extreme または Mac M5 Max を購入するでしょう。
AMD の Radeon RX 9070 XT は、約 645 GB/秒であらゆるラップトップのデコードを上回り、数分の 1 の価格で Spark を 2 倍以上実現します。モデルが 16 GB に収まる限り、これはローカル AI で最も価値のあるプレイの 1 つです。 4 ビット 9B は飛びますが、4 ビット 27B はわずかにきしみます。ここでの主な問題は、搭載される RAM の量です。 AMD の友人たちに向かって拳を振る私を想像してみてください。最大16GB？なぜ？なぜこんなことをしたのですか？
2 つの 4090 では 2,000 GB/秒の Mac は実現できません

ひね。モデルを 2 つのカードに分割する通常の方法は、レイヤーごとにネットワークの半分をそれぞれに分割し、トークンは引き続きレイヤーを順番に通過するため、各カードは半分の時間アイドル状態になります。 2 番目のカードで実際に購入するのは容量です。48 GB の VRAM は、およそシングル 4090 速度の 4 ビット 70B に十分です。 (Tensor 並列セットアップでは速度をいくらか取り戻すことができますが、それはチェックボックスではなく、サーバー ソフトウェアの領域です。) 帯域幅を「追加」するために 2 枚の中間層カードを使用するのは、よくある誤解であり、コストがかかります。しかし、1000 GB/秒は冗談ではありません。 48GB の RAM は、27B モデルを問題なく処理します。
裏計算: デコード速度 ≈ 帯域幅 ÷ モデル サイズ。 152 GB/秒のバス上の 5 GB モデルでは、TOPS の数に関係なく、最高で約 30 トークン/秒になります。これは、各トークンの生成が 5 GB すべての読み取りを意味するためです。私たちの測定結果はまさにこの線上にあります。CPU、NPU、またはその両方を使用しているかどうかに関係なく、Copilot+ 認定ラップトップで 4 ビット デコードの 9B モデルは 21 ～ 23 トークン/秒でした。
これは、大きなモデルを NPU にルーティングすることは無意味であるか、より悪いことであることを、ストップウォッチを使って苦労して学んだ理由でもあります。 NPU は CPU と同じメモリ バスを共有するため、これ以上高速にデコードすることはできません。また、専用の高速メモリは小さいため、およそ 3B パラメータを超えるものは、NPU が実際に高速である場所には適合しません。私たちの作業ルールは単純になりました。~3B を超えるモデルは CPU 上で実行されます。 NPU は小型モデルおよびプレフィル用です。 NPU の小型モデル: 素晴らしい (3B が完全なタスクを 5.8 秒で実行する)。 NPU で 9B: 最高でも CPU と同じ速度ですが、場合によってはそれより遅くなります。
小さいモデルは、少し速いだけではなく、比例して高速になります。同じマシン上で、パラメータの半分、1 秒あたりのトークンの 2 倍。これで、実際に実行すべき内容がわかります。
適切なタスクに適切なモデル
「どのモデルを使えばいいの？」あなたが理解するまで、それは間違った質問です

「何のために？」と答えました。機能はサイズに応じて均等に拡張されません。これらは段階的に到着します。
チャットと Q&A: 驚くほど奥深くまで機能します。 3Bモデルを探してください。それがあなたのスイートスポットです。 4B に移行するとすぐに、Copilot+ のメモリ帯域幅の仕様が最大になります。モデルがサポートしているかどうかを考えてオフにします。スピードの恩恵をすべて失い、3B モデルは決してスマートであるとは考えられません。
ツールの使用 (モデルは関数を呼び出します: 検索、ファイルを開く、クエリの実行): 9B。これは「このメールを編集」レベルです。
エージェント作業 (マルチステップ: 検索、結果の読み取り、決定、再度の行動): 現時点では 27B が最適なスポットだと思います。しかし、13B モデルはこれの多くを実行でき、M5 Max レベルのマシンでリアルタイムで実行できます。
判断（どちらが優れているのか？この主張は支持されているのか？私は間違いを犯したのか？）：最後に現れるもの。私たちの調査では、27B は、クリーンなツールまたは自己修正ツールの使用により、すべての構成で 3 対 3 の成績を収めた唯一のモデルでした。自身のエラーに気づき、修正しました。それは能力であり、圧縮されるものではありません。ただし、これは新しい 27B モデルにのみ適用されます。
補足: 考えることが常に良いアイデアであるとは限りません
推論モデル、つまり答える前に目に見えるスクラッチパッドで熟考するモデルは、質問を購入する明白な方法のように見えます。

[切り捨てられた]

## Original Extract

The Local AI Primer

Brad Wardell on X: "Here's how you get set up with a local AI on your PC or Mac. https://t.co/7Cn1GgTNIm" / X Post
Brad Wardell @draginol Here's how you get set up with a local AI on your PC or Mac. The Local AI Primer
I've taken the deep dive into local AI so you don't have to. This is how to use AI that works on your local PC or Mac. No cloud.
My feed is filled with local AI stuff. There's a lot of "boy who cried wolf" feeling around it because I hear claims and then only later you find out they were running on $100k of hardware. For that I could subscribe to Claude Max, Codex Pro and Grok Super Heavy Build for decades.
What I really wanted to know is what hardware can do what. What were the real limitations and what they could really do. And so I dug in. And now, you can benefit from that journey. And yes, I use em-dashes — a lot. They're mine. AI can't claim them.
AI Parameters: 2B, 3B, 9B, 27B for fun and profit
When you see a model called "Llama 8B" or "Gemma 27B," the number is its parameter count, the number of learned weights inside the neural network, in billions. Parameters are, roughly, the model's capacity to know things and reason about them.
~1B and under: Autocomplete with hallucinations of grandeur. Fine for classification, summarizing a paragraph, simple formatting tasks.
2B–4B: Genuinely conversational. Can follow instructions, answer general questions, and with careful prompting, do simple tool use. This is the sweet spot for phones and NPUs.
7B–9B: The workhorse class. Decent general knowledge, can follow multi-step instructions, handles light agentic work (call a tool, read the result, respond).
13B–30B: Where local models start feeling like "real" AI. Better judgment, fewer hallucinations, can recover from their own mistakes mid-task.
70B+: Approaching frontier-model territory, but you need serious hardware. Think a Mac Studio with 128GB of unified memory or a multi-GPU rig, not a laptop.
Parameter count also affects speed, not just smarts. Every single token the model generates requires reading essentially all of the parameters from memory. A 27B model isn't just smarter than a 3B model, it's also about 9x more work per word. Keep that in mind; it becomes a big deal when we start thinking about what you can do in real-time vs. what you should schedule out.
Models are trained in 16-bit precision. An 8B model at 16 bits per parameter is 16GB just for the weights.
Think of it like graphics formats. 16-bit precision is like a bitmap. No one, anymore, sends a .BMP in email.
Quantization is lossy compression for neural networks: store each weight in 8, 5, or 4 bits instead of 16. So it's like taking that BMP and turning it into a JPEG. At 8-bit quantization it's essentially the same as the original uncompressed version. At 4-bit it's lost some but probably not enough to matter other than in benchmarks. Below that and it starts to get pretty bad. 4-bit is generally the sweet spot.
The alphabet soup: Q4_0, Q4_K_M, Q8_0, IQ4_NL
Quantization also has many ways of doing it. Just like there's a ton of graphic formats there's a ton of ways of quantizing these models, each with their own trade-offs.
In my experience, I tend to focus on MLX (for Mac hardware), the ones that NVIDIA likes, and the ones that NPUs from Qualcomm and others will like.
Q8_0: 8-bit. Nearly lossless, twice the size of 4-bit. Worth it only for small models where the size doesn't hurt (a 0.6B model at Q8 is under 1GB, so why not).
Q4_K_M: the modern 4-bit "K-quant." Smarter allocation of bits (important layers get more precision). Usually the best quality per gigabyte, and the default recommendation on CPUs and GPUs.
Q4_0: the original, simplest 4-bit format. Slightly worse quality than Q4_K_M, but its plain block layout is what specialized hardware paths are built for. On ARM CPUs and NPUs, Q4_0 is often dramatically faster than Q4_K_M because the fast kernels only speak Q4_0.
IQ4_NL: a newer 4-bit format using a non-linear codebook. Great quality, but hardware support is spottier.
The format has to match the hardware. In our own testing, the same 4B model ran at 37 tokens/second as Q4_K_M on CPU but collapsed to 14 when routed to the NPU, because the NPU path couldn't handle K-quants and fell back to a slow path. Same model, same computer, same "4-bit," 2.5x difference. This is the single most common way people accidentally sandbag their local AI setup. This is why it's still a headache and programs like Clairvoyance are exploding in popularity — they just take care of this nonsense.
TOPS, NPUs, GPUs: the speed of thinking
Chip vendors advertise TOPS, trillions of operations per second. A Copilot+ certified laptop's NPU claims 40–80 TOPS; an RTX 4090 delivers over 600. You'd think a 45-TOPS NPU runs AI at some meaningful fraction of a 4090's speed.
There are two very different phases when a model responds to you:
Prefill: reading your prompt. This can be done in parallel so GPUs and NPUs do great here.
Decode: writing the answer. Tokens come out one at a time, and each one requires streaming the entire model through memory again. TOPS are nearly irrelevant; memory bandwidth is everything.
So a laptop NPU can read and understand a full article of text in half-a-second. But actually commenting on it is limited by the memory bandwidth which might take a minute.
RAM: system RAM, GPU VRAM, and unified memory
The model has to live somewhere, whole, in fast memory:
Discrete GPU (VRAM): Fastest option by far, but VRAM is scarce. A 12GB card fits a 9B comfortably or a 27B not at all. Spilling layers to system RAM works but every spilled layer runs at system-RAM speed.
System RAM (CPU): Plentiful and cheap (32GB fits anything you'd sanely run) but slow, see below.
Unified memory (Apple Silicon, Snapdragon X, AMD Strix Halo): CPU, GPU, and NPU share one pool. The great trick of a 128GB Mac is that all 128GB of it is available to the model at decent bandwidth, even though that bandwidth isn't remarkable on its own.
Sizing rule of thumb: model file size + 20–30% for the working context. A 5GB Q4 model wants roughly 7GB free.
Memory bandwidth: the often-ignored bottleneck
Memory bandwidth predicts local AI decode speed better than any other spec:
Typical laptop DDR5: ~60–90 GB/s
Copilot+ certified laptop (LPDDR5X, shared): ~120–152 GB/s
Snapdragon X2 Elite Extreme (LPDDR5X, shared): ~228 GB/s
Nvidia DGX Spark (LPDDR5X, unified): ~273 GB/s
Radeon RX 9070 XT (GDDR6): ~645 GB/s
RTX 4090 (GDDR6X): ~1,008 GB/s
RTX PRO 6000 Blackwell (GDDR7): ~1,792 GB/s
These speeds make the difference between whether you should be doing a task in real time or be scheduling it.
One of the first things I realized is that my laptop of choice can run a 27B model just fine. It's just slow at doing it. But most of the work I need to do with AI is not real-time. The very first feature that I used in Clairvoyance was the scheduling. I have so many dashboards, crash reports, sentiment reports, sales data coming in that I just have it run overnight for me to look at in the morning. I used to have that on Claude and that stuff was costing me $100 a month in tokens. But a 27B model can do it exactly as well and at no cost as long as I schedule it.
Diversion: Geek out on Hardware
The RTX 5090 is the current consumer king. Nearly 1.8 TB/s means a 5GB model decodes at hundreds of tokens/second, and its 32GB of VRAM fits a 4-bit 27B with room to spare. If your goal is "fast local AI, money is no object," this is the answer. As of this writing, an RTX 5090 currently retails for a little over $4.6 trillion dollars.
The RTX PRO 6000 Blackwell is the 5090's workstation sibling: same ~1.8 TB/s bus, but 96GB of VRAM, enough to hold a 4-bit 70B (or an 8-bit 70B, barely) on a single card at full speed. It solves the problem two 4090s can't (see below), big model and big bandwidth in one memory pool, for roughly the price of a decent used car. The use case for us would be you'd park one of these on the rack and have several users on it running a 27B model in real time.
The Nvidia DGX Spark is slower than you'd expect. It's marketed as an AI supercomputer for your desk, and for capacity it delivers. 128GB of unified memory fits models no consumer GPU can touch. But its 273 GB/s bandwidth is less than half an M5 Max, a quarter of a 4090. It runs big models acceptably; it does not run any model fast. You're buying capacity, not speed. There's a reason these are in stock at Microcenter. I'd just buy a Snapdragon X2 Elite Extreme or Mac M5 Max instead.
AMD's Radeon RX 9070 XT at ~645 GB/s out-decodes every laptop and more than doubles the Spark at a fraction of the price. It's one of the best value plays in local AI, provided the model fits in its 16GB. A 4-bit 9B flies, a 4-bit 27B just squeaks in. The main issue here is the amount of RAM they put on it. Envision me shaking my fist at my friends at AMD. 16GB max? Why? Why did you do this?
Two 4090s do not make a 2,000 GB/s machine. The usual way to split a model across two cards is by layers, half the network on each, and a token still passes through the layers in sequence, so each card sits idle half the time. What you actually buy with the second card is capacity: 48GB of VRAM, enough for a 4-bit 70B, at roughly single-4090 speed. (Tensor-parallel setups can claw back some speed, but that's server-software territory, not a checkbox.) Two mid-tier cards to "add up" bandwidth is a common and expensive misunderstanding. But 1000 GB/sec is no joke. 48GB of RAM handles that 27B model just fine.
The back-of-envelope math: decode speed ≈ bandwidth ÷ model size. A 5GB model on a 152 GB/s bus tops out around 30 tokens/second no matter how many TOPS you have, because generating each token means reading all 5GB. Our measurements land right on this line: 9B models at 4-bit decode at 21–23 tokens/second on a Copilot+ certified laptop whether we use the CPU, the NPU, or both.
This is also why we learned, the hard way, with a stopwatch, that routing big models to an NPU is pointless or worse. The NPU shares the same memory bus as the CPU, so it can't decode any faster, and its dedicated fast memory is tiny, so anything beyond roughly 3B parameters doesn't fit where the NPU is actually fast. Our working rule is now simply: models over ~3B run on the CPU; the NPU is for small models and for prefill. Small model on NPU: brilliant (a 3B doing a full task in 5.8 seconds). 9B on NPU: same speed as CPU at best, sometimes slower.
Smaller models are proportionally faster, not just a little faster. Half the parameters, twice the tokens per second, on the same machine. Which brings us to what you should actually run.
The right model for the right task
"Which model should I use?" is the wrong question until you've answered "for what?" Capabilities don't scale evenly with size. They arrive in tiers:
Chat and Q&A: Works surprisingly far down. Look for a 3B model. That's your sweet spot for this. As soon as you go to 4B you max out the Copilot+ spec for memory bandwidth. TURN OFF thinking if the model supports it. You lose all the benefit of speed and a 3B model will never think itself into being smart.
Tool use (the model calls functions: search, open a file, run a query): 9B. This is the "Edit this email" level.
Agentic work (multi-step: search, read the result, decide, act again): Right now I would say 27B is the sweet spot. But a 13B model can do a lot of this and run real-time on an M5 Max level machine.
Judgment (which of these is better? is this claim supported? did I make a mistake?): The last thing to emerge. In our sweeps, the 27B was the only model that went 3-for-3 on every configuration with clean or self-correcting tool use. It noticed its own errors and fixed them. That is a capability, and it doesn't compress. But this is only on newer 27B models.
Side Note: Thinking is not always a good idea
Reasoning models, the ones that deliberate in a visible scratchpad before answering, look like the obvious way to buy qu

[truncated]
