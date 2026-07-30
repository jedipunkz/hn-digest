---
source: "https://swellweb.github.io/reame/bytes/"
hn_url: "https://news.ycombinator.com/item?id=49113783"
title: "LLM re-reads the same text a thousand times. Here's what I measured"
article_title: "Your LLM re-reads the same text a thousand times — Reame"
author: "targetbridge"
captured_at: "2026-07-30T19:10:13Z"
capture_tool: "hn-digest"
hn_id: 49113783
score: 3
comments: 0
posted_at: "2026-07-30T18:28:32Z"
tags:
  - hacker-news
  - translated
---

# LLM re-reads the same text a thousand times. Here's what I measured

- HN: [49113783](https://news.ycombinator.com/item?id=49113783)
- Source: [swellweb.github.io](https://swellweb.github.io/reame/bytes/)
- Score: 3
- Comments: 0
- Posted: 2026-07-30T18:28:32Z

## Translation

タイトル: LLM は同じテキストを 1,000 回読み直します。私が測定したものはこれです
記事のタイトル: LLM は同じテキストを何千回も再読します — Reame
説明: 2 週間にわたる CPU 推論の最適化、そして成功したのは

記事本文:
LLM は同じテキストを何千回も再読します
CPU 推論を高速化するために 2 週間を費やしました。そのこと
最終的に機能したのはエンジンではありませんでした。
無料の Oracle ARM ボックスで測定 — 4 コア、月額 0 ユーロ。すべて
以下は再現可能です。スクリプトは最後にリンクされています。
私は、llama.cpp 上に構築された CPU ファーストの推論サーバーを保守しています。フォールバックとしてではなく
欠落している GPU の場合、ターゲットとして。私が気になるハードウェアは無料枠です。
2 ～ 4 つの ARM コア、アクセラレータなし、予算なし。
そこでのボトルネックはプレフィルです。モデルが単一のトークンを書き込む前に、
プロンプト全体を読み、読みが優先される任意のサイズの文書を読みます。
他のすべて。そこで、わかりやすい場所でのスピードを追求しました。
私は llama.cpp を上流のカーネル作業の 1 年分前倒ししました。
ARM の最適化が表示されます。この量子化とこの CPU では: 0% 。
量子化をもっと粗くしてみました。 Q4_0 はプレフィルで 37% 高速になり、
デコードが 19% 速くなり、抽出テストでは 20 件中 5 件の事実が削除されました。
拒否されました。
専門家の混合で事前入力中にアクティブな専門家の数を半分にしました
モデル。 44% 高速で、静かにキャッシュを破壊します: 4 で構築された KV キャッシュ
専門家とリードバックすると、14/20 のコントロールに対して 8 スコア 11/20 でした。ダメージ
出力だけでなく、キャッシュされた表現に焼き付けられます。あれは
知っておく価値があります。
同じページです。同じモデルです。同じ質問です。として書き換えられた
label: value 、1 行に 1 つのファクト、条件が付加されたまま
資格があるもの:
6.5 倍速く、より正確です。その 2 番目の部分は、
私が予想していなかったことが、これを書き留める価値がある理由です。少ない
トークンは情報が少なく、より悪い回答を意味する必要があります。そうではありませんでした。
モデルは正しい数字を見て間違って答えました
なぜ精度が高いのかを理解するには

上がったので、ソフトマックス後の注意を捨てましたが、
モデルは、4 つの同様の価格をリストしたページに関する質問に答えました (1 つずつ)。
サービスはすべて同じブロック内にあります。
1.9% 対 1.7% はコイントスです。モデルは数字を無視することはありませんでした -
それを質問に結び付けることができませんでした。散文は価格と
他のすべてと同じ文でサービスを提供します。ラベル: 価値が生み出すもの
バインディングはセマンティックではなく構造的であり、比率は 1.1:1 から
7:1。
実際的な結果: 小さなモデルが事実を誤った場合、「
より大きなモデル」は 1 つの答えですが、多くの場合、「ドキュメントを別の方法で書く」が重要です。
安くて、より良く機能します。
それからバイトがどこに行くのか探しに行きました
推測するのをやめて、モデルが生成されたデータごとに実際に何を読み取るかを数えてみました。
トークン。どこにも書かれていないことが 2 つ出てきました。両方とも
GPU では表示されません。
1. 出力ヘッドは帯域幅の 4 分の 1 です。
多くの小規模モデルでは、埋め込み行列は結合されています。つまり、埋め込み行列は
出力投影が行われ、生成されたトークンごとに完全に読み取られ、すべてのトークンがスコア付けされます。
151,936 の語彙エントリ。
トークンあたりの帯域幅の 3 分の 1 近くが、言語の単語のスコアリングに使用されます。
この展開では決して放出されません。私のドメイン - イタリア語、独自のコード、技術的
英語 — 151,936 エントリのうち 9,314 を使用します。
語彙を 32,000 にトリミングすると、構造的には安全であることが判明します。
正確に述べるべき理由。トークナイザーはバイトレベルであり、すべて 256
バイト文字は保持されるため、テキストが表現できなくなることはありません。最悪の場合は
トリミングされた単語には追加のトークンが必要です。そして埋め込みは行のあるQ6_Kです
整数の量子化ブロックにまたがるため、行全体が何もせずにドロップアウトします。
ブロックを分割することはありません: 生き残る重みはビットごとです
同一です。再量子化や数値損失はありません。
tで測定

フリーボックス、3 つの独立したラン、スプレッド 2% 未満:
+17.8% デコード (46.7 トークン/秒から 55.0 トークン/秒)。 20の質問
事実試験のスコアは前後で同じで、両方とも 13/13 の重要な事実でした。
保留されたテキストで測定されたコストは、同じドキュメントのトークンの数が 1.9% 増加します。
したがって、これは入力がすでに不足している場合に最も効果的です。
帯域幅だけによる理論上の上限は +30% でした。 +17.8%でした。ギャップ
これは体重測定ではない時間なので、測定した数値を公表したいと思います
計算されたものよりも。
2. 重量の再梱包にかかる費用は 4.2 GB 誰も言及していません
私のサーバーでは、5.37 GB モデルの場合、9,825 MB が常駐していました。漏れを想定しました。
それはそうではありません:
load_tensors: CPU_Mapped モデルのバッファ サイズ = 5068.51 MiB
load_tensors: CPU_REPACK モデルのバッファ サイズ = 4254.45 MiB
llama.cpp はロード時に量子化された重みを NEON レイアウトに再パックします
カーネルの読み取りが速くなります。これは賢明なエンジニアリング取引ですが、GPU ではそうではありません
おそらくそれが、メモリのコストがどこにも文書化されていない理由でしょう。
見つけます。
私は最初にソースを読んで、リパックでは 2D テンソルのみに触れていると結論付けました。
そうすれば環境省の専門家は助かっただろう。ログには別のことが書かれていました:
ffn_gate_exps と ffn_up_exps も再パックされます。
ffn_down_exps は q6_K から q8_0_4x4 に移行します —
6 ビットから 8 ビットへ。リパックは再編成するだけではありません。ところどころそれ
が拡大します。
-4.2 GB の RAM で -12.8% の速度。メモリを搭載したマシンでは、
スペア、それは悪い取引です。 8 GB ボックスではモデル間の違いになります
まったく実行されている場合と実行されていない場合、まさにこれが対象のハードウェアです。
事実試験は私のものです: 1 つの実際のイタリアのビジネスについて私が書いた 20 の質問
ページ、正規表現によって評価されます。 1 ページ、1 言語、1 ドメイン。それは最も弱いです
これはこの記事の一部であり、指摘されるよりもそう言いたいと思います。
正直さを保つのは、同じことです

試験は私自身の 2 つのアイデアを台無しにしました。
量子化の高速化とエキスパートの削減（両方とも上記）。そしてその試験では、
私がサーブする速い MoE は、より遅い密度の高い 3B に負けます。入力の再フォーマット
高速モデルが失いかけていた精度を回復しました。それは小さなモデルを作りませんでした
賢い。資料が提供されていないのに何でも尋ねると、物事を発明します — 私は出版しました
あのテーブルも。
小規模モデル用の公開敵対的事実抽出セットを知っている場合は、次の点を指摘します。
私はそれに取り組んでいます、そして私はそれを実行し、悪いものも含めて出てくるものをすべて公開します
結果。
ここでのすべての勝利は、実際に作業がどこに進むかを測定することから生まれました。
損失は仮定から生じました。カーネルバンプ、粗い量子化、減少した
専門家: すべての合理的な考えは、このハードウェアではすべて間違っています。テキストを書き換えると、
語彙のトリミング、再パックの切り替え：すべて退屈、すべて測定可能、すべて
重要な。
CPU 上で最も速い計算は、それが
決して読んだことのないトークン、決して得点したことのない語彙エントリー、または決して得点したことのないバイト
再梱包します。
ページ、20 問、採点者は次のとおりです。
ベンチ/
OpenAI 互換サーバーに対して実行します。を含む完全なベンチマーク
否定的な結果は、
ベンチマーク.md 。
サーバーはMITのReameです
ライセンスを取得した。ライブデモがあります
同じ無料の ARM ボックス上で実行されます。1 つの小さなインスタンスで、自動スケーリングはありません。
月額 0 ユーロの正直な容量を見ていると遅いです。

## Original Extract

Two weeks of CPU inference optimization, and the thing that won wasn

Your LLM re-reads the same text a thousand times
I spent two weeks trying to make CPU inference faster. The thing
that finally worked was not in the engine.
Measured on a free Oracle ARM box — 4 cores, €0/month. Everything
below is reproducible; the scripts are linked at the end.
I maintain a CPU-first inference server built on llama.cpp. Not as a fallback
for a missing GPU — as the target. The hardware I care about is the free tier:
two to four ARM cores, no accelerator, no budget.
The bottleneck there is prefill. Before a model writes a single token it has
to read your entire prompt, and on a document of any size that read dominates
everything else. So I went looking for speed in the obvious places.
I bumped llama.cpp forward by a year of upstream kernel work, expecting the
ARM optimizations to show up. On this quantization and this CPU: 0% .
I tried a coarser quantization. Q4_0 is 37% faster at prefill and
19% faster at decode — and it dropped 5 facts out of 20 on my extraction test.
Rejected.
I halved the number of active experts during prefill on a mixture-of-experts
model. 44% faster, and it silently corrupts the cache: a KV cache built with 4
experts and read back with 8 scores 11/20 against a 14/20 control. The damage
is baked into the cached representation, not just the output. That one was
worth knowing about.
The same page. The same model. The same question. Rewritten as
label: value , one fact per line, with conditions kept attached to
what they qualify:
Six and a half times faster, and more accurate. That second part is
what I didn't expect, and it's the reason this is worth writing down. Fewer
tokens should mean less information and worse answers. It didn't.
The model saw the right number and answered wrong
To understand why accuracy went up, I dumped the post-softmax attention while
the model answered a question about a page listing four similar prices — one per
service, all in the same block.
1.9% against 1.7% is a coin flip. The model was never blind to the number —
it just couldn't bind it to the question. Prose puts the price and the
service in the same sentence as everything else; label: value makes
the binding structural instead of semantic, and the ratio goes from 1.1:1 to
7:1.
The practical consequence: when a small model gets a fact wrong, "use a
bigger model" is one answer, but "write the document differently" is often
cheaper and works better.
Then I went looking for where the bytes go
Having stopped guessing, I counted what a model actually reads per generated
token. Two things came out that I have not seen written down anywhere, both
invisible on a GPU.
1. The output head is a quarter of your bandwidth
In many small models the embedding matrix is tied : it doubles as the
output projection and is read in full for every token generated, to score all
151,936 vocabulary entries.
Nearly a third of the bandwidth per token goes to scoring words in languages
this deployment will never emit. My domain — Italian, my own code, technical
English — uses 9,314 of those 151,936 entries.
Trimming the vocabulary to 32k turns out to be safe by construction, for two
reasons worth stating precisely. The tokenizer is byte-level and all 256
byte-characters are kept, so no text becomes unrepresentable — the worst case is
that a trimmed word costs an extra token. And the embedding is Q6_K with rows
spanning a whole number of quantization blocks, so whole rows drop out without
ever splitting a block: the surviving weights are bit-for-bit
identical . No requantization, no numerical loss.
Measured on the free box, three independent runs, under 2% spread:
+17.8% decode , from 46.7 to 55.0 tokens/second. The 20-question
fact exam scores the same before and after — 13/13 critical facts both times.
The cost, measured on held-out text, is 1.9% more tokens for the same document,
so this pays off most when the input is already short.
The theoretical ceiling from bandwidth alone was +30%. I got +17.8%. The gap
is time that isn't weight reading, and I'd rather publish the measured number
than the calculated one.
2. Weight repacking costs 4.2 GB nobody mentions
My server showed 9,825 MB resident for a 5.37 GB model. I assumed a leak.
It isn't:
load_tensors: CPU_Mapped model buffer size = 5068.51 MiB
load_tensors: CPU_REPACK model buffer size = 4254.45 MiB
llama.cpp repacks quantized weights at load time into a layout its NEON
kernels read faster. It's a sensible engineering trade — and on a GPU it doesn't
exist, which is probably why the memory cost isn't documented anywhere I could
find.
I had first read the source and concluded the repack only touched 2D tensors,
which would have spared the MoE experts. The logs said otherwise:
ffn_gate_exps and ffn_up_exps get repacked too, and
ffn_down_exps goes from q6_K to q8_0_4x4 —
from 6 bits to 8. The repack doesn't only reorganize; in places it
expands .
−4.2 GB of RAM for −12.8% speed. On a machine with memory to
spare that's a bad trade. On an 8 GB box it's the difference between the model
running at all and not running — which is exactly the hardware this is for.
The fact exam is mine: 20 questions I wrote over one real Italian business
page, graded by regex. One page, one language, one domain. It is the weakest
part of this write-up and I'd rather say so than have it pointed out.
What keeps it honest is that the same exam killed two of my own ideas — the
faster quantization and the reduced experts, both above. And on that exam the
fast MoE I serve is beaten by a slower dense 3B. Reformatting the input
recovered accuracy the fast model was losing; it did not make the small model
smart. Asked anything with no document supplied, it invents things — I published
that table too.
If you know a public adversarial fact-extraction set for small models, point
me at it and I'll run it and publish whatever comes out, including a bad
result.
Every win here came from measuring where the work actually goes, and every
loss came from assuming. The kernel bump, the coarser quantization, the reduced
experts: all reasonable ideas, all wrong on this hardware. The text rewrite,
the vocabulary trim, the repack switch: all boring, all measurable, all
significant.
On a CPU, the fastest computation is the one you don't do — whether that's a
token you never read, a vocabulary entry you never score, or a byte you never
repack.
The page, the 20 questions and the scorer are in
bench/
and run against any OpenAI-compatible server. Full benchmarks, including the
negative results, are in
BENCHMARKS.md .
The server is Reame , MIT
licensed. There's a live demo
running on the same free ARM box — one small instance, no autoscaling, so if
it's slow you're watching the honest capacity of €0/month.
