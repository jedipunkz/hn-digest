---
source: "https://vettedconsumer.com/gguf-vs-gptq-vs-awq-the-plain-english-guide-to-llm-quantization-and-which-one-to-pick/"
hn_url: "https://news.ycombinator.com/item?id=48425712"
title: "GGUF vs. GPTQ vs. AWQ: The Plain-English Guide to LLM Quantization"
article_title: "GGUF vs GPTQ vs AWQ: The Plain-English Guide to LLM Quantization (and Which One to Pick)"
author: "ermantrout"
captured_at: "2026-06-06T15:31:40Z"
capture_tool: "hn-digest"
hn_id: 48425712
score: 1
comments: 0
posted_at: "2026-06-06T14:58:11Z"
tags:
  - hacker-news
  - translated
---

# GGUF vs. GPTQ vs. AWQ: The Plain-English Guide to LLM Quantization

- HN: [48425712](https://news.ycombinator.com/item?id=48425712)
- Source: [vettedconsumer.com](https://vettedconsumer.com/gguf-vs-gptq-vs-awq-the-plain-english-guide-to-llm-quantization-and-which-one-to-pick/)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T14:58:11Z

## Translation

タイトル: GGUF vs. GPTQ vs. AWQ: LLM 量子化のわかりやすい英語ガイド
記事のタイトル: GGUF vs GPTQ vs AWQ: LLM 量子化のわかりやすい英語ガイド (およびどれを選択するか)
説明: GGUF、GPTQ、AWQ、Q4_K_M、NF4 — 量子化アルファベット スープ。所有している VRAM により大きなモデルを収めたいだけの人のために説明されています。各フォーマットとは何か、実際の VRAM の計算、および使用する決定テーブル。

記事本文:
ホーム
統合メモリ
AI用GPU
ミニ PC とサーバー
エッジAI
購読する
ブログ
暗い
ソフトウェアとツール
GGUF vs GPTQ vs AWQ: LLM 量子化のわかりやすい英語ガイド (およびどれを選択するか)
GGUF、GPTQ、AWQ、Q4_K_M、NF4 — 量子化アルファベット スープ。所有している VRAM にさらに大きなモデルを収めたいだけの人のために説明されています。各フォーマットとは何か、実際の VRAM の計算、および使用する決定テーブル。
トーマス・ニューカーク
2026 年 6 月 6 日
7 分で読めます
自分のマシンで大規模な言語モデルを実行しようとして時間を費やしたことがあるなら、誰もが経験するのと同じ壁にぶつかったことがあるでしょう。モデルは巨大であるのに、VRAM はそうではありません。ネイティブ 16 ビット精度の 700 億パラメータ モデルには、重みを保持するためだけで約 140 GB のメモリが必要です。それを持っている人はほとんどいません。量子化はギャップを埋める秘訣であり、専門用語の雪崩が始まる場所でもあります。ググフ。 GPTQ。 AWQ。 Q4_K_M. NF4。 EXL2。
このガイドは、量子化が実際に行うこと、各主要フォーマットの実際の目的、正直なトレードオフ、および 30 秒で使用できるデシジョン テーブルなど、私たちが開始したときに存在していればよかったと願っていたバージョンです。手を振る必要はありません。すでに新聞を読んでいると仮定する必要はありません。
モデルの「重み」は単なる数値の巨大な山です。デフォルトでは、それぞれは 16 ビット精度 (FP16 または BF16)、つまり重みあたり 2 バイトで保存されます。量子化では、より少ないビット (8、5、4、場合によっては 2) を使用して同じ数値を保存します。重みあたりのビット数が少ないと、ある程度の精度が犠牲になりますが、ファイルが小さくなり、メモリが少なくなります。
メモリの計算は驚くほど簡単です。パラメータ数に重みあたりのバイト数を掛けます。
FP16 (16 ビット): 2 バイト/重量 → 70B モデルには最大 140 GB が必要
8 ビット: ～ 1 バイト/重量 → ～ 70 ～ 75 GB
4 ビット: ~0.5 バイト/重量 → ~40 GB
16 ビットから 4 ビットへのその 1 つのジャンプにより、「データ CE が必要」になります。

「GPU 間」を「48 GB カードまたはユニファイド メモリ ボックスで実行」に変更します。驚くべき点、そして量子化が随所に存在する理由は、よくできた 4 ビット モデルの品質が驚くほどオリジナルに近いということです。劣化は現実のものですが、ごくわずかで、ほとんどの用途では目に見えません。70B モデルに実際に必要な VRAM の量については、関連記事で正確な数値を詳しく説明します。
GGUF — ほとんどの人が使用すべきもの
GGUF は、llama.cpp (およびその上に構築されたすべてのもの: Ollama、LM Studio、Jan、KoboldCpp) で使用されるファイル形式です。これは、古い GGML 形式の後継です。 Hugging Face からモデルをダウンロードし、ファイル名が .gguf で終わる場合、これが得られます。
GGUF のスーパーパワーは柔軟性です。 CPU、GPU、またはその両方の組み合わせで実行されます。必要に応じて多くのレイヤーを GPU にオフロードし、残りを CPU に処理させることができます。そのため、これが Mac ユーザー (金属経由の Apple Silicon) や、モデルが VRAM に適合しない人にとってのデフォルトとなっています。また、広範囲にわたる量子化レベルである「k-quants」も同梱されており、不可解なサフィックスはそこから来ています。
Q8_0 — ~8.5 ビット/重量、本質的にロスレス。記憶力があり、妥協を一切望まない場合に使用してください。
Q6_K — ~6.6 bpw、完全な精度とほとんど区別がつきません。
Q5_K_M — 〜5.7 bpw、高品質の中間点。
Q4_K_M — ~4.8 bpw。これはコミュニティのデフォルトであり、スイート スポットです。元のサイズの約 3 分の 1 に対して、約 1% の混乱が発生します。
Q3_K_M — ~3.9 bpw。著しく劣化していますが、メモリが不足している場合には使用できます。
Q2_K — ~3.4 bpw。 「とにかくロードしたいだけ」層。品質が大幅に低下します。それは最後の手段として扱ってください。
文字の接尾辞は重要です。_M (中) と _S (小) は、サイズと引き換えに品質が少し劣ります。新しい I-quants (IQ4_XS、IQ3_M など) は、もう少し多くの品質を絞り出します。

重要度マトリックス キャリブレーションを使用してバイトごとに y を計算しますが、一部のハードウェアでは推論がわずかに遅くなります。
Mac を使用している場合、CPU と GPU を混合している場合、最も幅広いモデル選択が必要な場合、または単に抵抗が最も少ないパスが必要な場合は、GGUF を使用してください。ほとんどの読者にとって、正直な答えは「ここから始めてください」です。
GPTQ — GPU ネイティブの 4 ビット標準
GPTQ は、Frantar らによって導入されたトレーニング後の量子化手法です。 2022 年 (arXiv:2210.17323、後に ICLR 2023 で発表)。すべての重みを単純に丸めるのではなく、近似 2 次 (ヘッセ行列) 情報を使用して、導入された誤差を補正しながら、一度に 1 列ずつ重みを量子化します。これは、巨大なモデルであっても数 GPU 時間で実行されるワンショット プロセスです。
実際的な点: GPTQ は重みのみ、GPU のみを使用し、推論が高速です。モデル全体が VRAM に収まり、GPU ランタイムを通じてモデルを提供している場合に輝きます。これは、長い間 Hugging Face で主流の 4 ビット形式であり、サービング スタックによって広くサポートされています。この弱点は、GGUF のように CPU に適切にスピルされないことです。これはオールイン VRAM フォーマットです。
モデルが GPU メモリに完全に収まり、GPU サービス用に十分にサポートされている成熟した 4 ビット形式が必要な場合は、GPTQ を使用します。
AWQ — 精度を重視するチャレンジャー
AWQ (Activation-aware Weight Quantization) は、Lin らによるものです。 ( arXiv:2306.00978 、MLSys 2024 最優秀論文)。その洞察は賢明です。すべての重みが同じように重要であるわけではありません。 「顕著な」ウェイト チャネルのごく一部 (~1%) (ウェイト チャネルそのものではなく、そこを流れるアクティベーションを調べることによって識別されます) が、モデルの品質において大きなシェアを占めています。 AWQ は、量子化する前にチャネルをスケーリングすることでこれらのチャネルを保護するため、重要な部分は 4 ビット圧縮後もほぼそのまま残ります。
実際には、AWQ は多くの場合、または b と一致します。

同じビット幅での精度で GPTQ を使用し、vLLM のような高スループットのサービング エンジンでは第一級の地位を占めます。 GPTQ と同様に、GPU 指向であり、モデルが VRAM 内に存在することを前提としています。
AWQ は、GPU (特に vLLM 経由) でサービスを提供していて、4 ビットで得られる最高の品質が必要な場合に使用します。
bitsandbytes / NF4 — QLoRA 論文 (Dettmers et al.、arXiv:2305.14314) の 4 ビット NormalFloat 形式。これは、コンシューマ GPU で量子化モデルを微調整するための頼りになるもので、オンザフライで量子化します。トレーニングに最適です。通常、純粋な推論を求める場合は最初に選択しません。
EXL2 (ExLlamaV2) — 可変ビットレートの GPU 専用フォーマット (たとえば 4.65 bpw をターゲットにできます)。 NVIDIA カードでは非常に高速でメモリ効率が高くなります。単一の GPU で 1 秒あたりのトークン数を最適化する人々に愛されています。
どのビットレベルを選択する必要がありますか?
これは実際に日常的に重要な質問です。数え切れないほどのコミュニティテストを通じて維持されてきたコンセンサス:
Q4_K_M がデフォルトであるのには理由があります。これは、ほとんどの人にとってギガバイトあたりの品質が最高であり、完全な精度に比べて複雑さが約 1% 増加し、サイズが最大 65% 削減されます。
メモリに余裕がある場合は、Q5_K_M または Q6_K にステップアップしてから、より低い量のより大きなモデルに手を伸ばしてください。通常、Q4 の 70B は Q2 の 70B を上回ります。
しかし、通常、中程度のクオンツではより大きなモデルが、高クオンツではより小さなモデルよりも優れています。通常、Q4_K_M の 70B は Q8 の 13B を上回ります。迷った場合は、第 4 四半期あたりまで、小規模モデルの上位クオントではなく、より大規模なモデルの下位クオンツを選択してください。 Q3 以下では、計算が逆転します。
他に選択肢がない限り、Q2 は避けてください。品質の崖は、約 3 ビット以下で急速に急勾配になります。
これらの選択肢が 70B モデルの実メモリ内でどのように表示されるかを次に示します (重みのみ。コンテキスト/KV キャッシュ用のヘッドルームを追加します)。
低ビット量子化が想起より推論に悪影響を与える理由
ここ

これは、ビットレベル テーブルが隠しているニュアンスです。量子化によってすべてのタスクが均等に劣化するわけではありません。ダメージは均一ではありません。その理由を理解することで、実際のワークロードに対して間違ったクオントを選択することがなくなります。
機械的には、量子化によりすべての重みに少量の丸めノイズが追加されます。事実を思い出す、多肢選択式の質問に答える、文章を完成させるという単一ステップのタスクの場合、通常、そのノイズはモデルの信頼性によって圧倒されます。正しい答えが勝ちます。これが、アグレッシブなクオンツが、HellaSwag のような短いベンチマークではほとんどロスレスに見えることが多い理由です。信号は存続します。
複数ステップの推論 (ロジックのチェーン、コンパイルが必要なコード、正確にフォーマットする必要があるツールの呼び出し、多くの詳細を一度に追跡する必要があるロングコンテキストのタスク) の場合、同じ小さなノイズが複合します。ステップ 1 での小さなエラーによりステップ 2 が変更され、さらにステップ 3 が変更され、最後にはチェーンがコースから外れてしまいます。モデルは何も「忘れて」いませんでした。長いパスにわたって丸め誤差が蓄積されました。長いコンテキストは問題をさらに悪化させます。なぜなら、注意をまっすぐに保つためにより多くの競合する詳細があり、低ビットの重みによって区別が曖昧になるからです。
これは所有者の報告と一致します。 r/LocalLLaMA に関する繰り返しのコンセンサスは、チャットには 4 ビットが適していますが、正確な複数ステップの作業を信頼する前に、もう一度考えてみる価値があるということです。
「チャットには問題ありませんが、実際のデータ作業には使用しないでください。」
— u/autonomousdev_、Q8 が抽出タスクで Q4 が見逃した 2 つのエッジ ケースをキャッチした後
「アグレッシブなクオンツ (<Q6) は、4K コンテキストをよく使用するベンチマークでは問題なく動作します。実際には、30K 以上を処理する必要がある場合、それらは崩壊します。」
— u/ClearApartment2627
カジュアルなチャット、ブレーンストーミング、要約、簡単な Q&A: Q4_K_M を完全な精度と区別するのは非常に困難です。 VRAMを保存します。
コーディング、エージェント

ol の使用、構造化抽出、長い文書: 適合する場合は Q5_K_M または Q6_K にステップアップします。余分なビットにより、量子化ノイズが最も大きなダメージを与える正確な場所の信頼性が高まります。
KV キャッシュにも注意してください。メモリを節約するためのコンテキスト キャッシュの量子化は、重みの量子化とは別の操作です。そして、長いコンテキストでは、重みの量子化よりも害が大きいことがよくあります。モデルが約 16 ～ 32,000 トークンを超えてバラバラになった場合は、重みのせいにする前に、KV キャッシュの量子化を疑ってください。
これらはいずれも、「Q4_K_M がデフォルトである」と矛盾しません。これを洗練させます。Q4 が一般的な使用のデフォルトであり、意図的なステップアップは、すべてのステップが到達する必要があるワークロードに対する安価な保険になります。
情報源と調査方法
この説明者は、元のメソッドの論文である GPTQ ( Frantar et al., 2022 )、AWQ ( Lin et al., 2023 )、および QLoRA/NF4 ( Dettmers et al., 2023 ) を、llama.cpp ドキュメントと GGUF k-quant ビット/ウェイト数値の量子化テーブルと併せて合成します。ビットレベルのガイダンスは、r/LocalLLaMA コミュニティとメンテナーのベンチマークからの広範で耐久性のあるコンセンサスを反映しており、私たちの側での直接のテストではありません。 VRAM の数値は、わかりやすくするために四捨五入された重みのみの推定値です。実際の使用法では、KV キャッシュとコンテキスト ウィンドウ用のメモリが追加されますが、これについては別途説明します。
70B モデルをローカルで実行するには実際にどれくらいの VRAM が必要ですか?
Ollama vs LM Studio vs llama.cpp: 実際にどのランタイムを使用する必要がありますか?
70B モデルをローカルで実行する最も安価な方法
精査された消費者向けニュースレターを入手する
レビュー、購入アドバイス、フィールドノート。毎月お届けします。
地元の LLM ハードウェアに関する正直なアドバイス — レビューとその所有者を横断的にチェックしました。

## Original Extract

GGUF, GPTQ, AWQ, Q4_K_M, NF4 — the quantization alphabet soup, explained for people who just want to fit a bigger model in the VRAM they have. What each format is, the real VRAM math, and a decision table for which to use.

Home
Unified Memory
GPUs for AI
Mini PCs & Servers
Edge AI
Subscribe
Blog
Dark
Software & Tools
GGUF vs GPTQ vs AWQ: The Plain-English Guide to LLM Quantization (and Which One to Pick)
GGUF, GPTQ, AWQ, Q4_K_M, NF4 — the quantization alphabet soup, explained for people who just want to fit a bigger model in the VRAM they have. What each format is, the real VRAM math, and a decision table for which to use.
Thomas Newkirk
June 6, 2026
7 min read
If you have spent any time trying to run a large language model on your own machine, you have hit the same wall everyone does: the model is enormous and your VRAM is not. A 70-billion-parameter model in its native 16-bit precision wants about 140 GB of memory just to hold the weights. Almost nobody has that. Quantization is the trick that closes the gap — and it is also where the jargon avalanche begins. GGUF. GPTQ. AWQ. Q4_K_M. NF4. EXL2.
This guide is the version we wish existed when we started: what quantization actually does, what each of the major formats is really for, the honest trade-offs, and a decision table you can use in thirty seconds. No hand-waving, no assuming you already read the papers.
A model's "weights" are just a giant pile of numbers. By default each one is stored at 16-bit precision (FP16 or BF16) — two bytes per weight. Quantization stores those same numbers using fewer bits : 8, 5, 4, sometimes as low as 2. Fewer bits per weight means a smaller file and less memory, at the cost of some precision.
The memory math is refreshingly simple. Multiply the parameter count by the bytes per weight:
FP16 (16-bit): 2 bytes/weight → a 70B model needs ~140 GB
8-bit: ~1 byte/weight → ~70–75 GB
4-bit: ~0.5 byte/weight → ~40 GB
That single jump from 16-bit to 4-bit is what turns "needs a data-center GPU" into "runs on a 48 GB card, or a unified-memory box." The surprising part — and the reason quantization is everywhere — is that a well-done 4-bit model is shockingly close in quality to the original. The degradation is real but small, and for most use it is invisible. We break the exact numbers down in our companion piece on how much VRAM you actually need for a 70B model .
GGUF — the one most people should use
GGUF is the file format used by llama.cpp (and everything built on it: Ollama, LM Studio, Jan, KoboldCpp). It is the successor to the older GGML format. If you download a model from Hugging Face and the filename ends in .gguf , this is what you have.
GGUF's superpower is flexibility . It runs on CPU, GPU, or a mix of both — you can offload as many layers to your GPU as fit and let the CPU handle the rest. That is why it is the default for Mac users (Apple Silicon via Metal) and for anyone whose model does not quite fit in VRAM. It also ships in a huge range of quantization levels, the "k-quants," which is where the cryptic suffixes come from:
Q8_0 — ~8.5 bits/weight, essentially lossless. Use it when you have the memory and want zero compromise.
Q6_K — ~6.6 bpw, near-indistinguishable from full precision.
Q5_K_M — ~5.7 bpw, a high-quality middle ground.
Q4_K_M — ~4.8 bpw. This is the community default and the sweet spot: about a 1% perplexity hit for roughly a third of the original size.
Q3_K_M — ~3.9 bpw. Noticeably more degraded, but usable when memory is tight.
Q2_K — ~3.4 bpw. The "I just want it to load at all" tier. Quality drops meaningfully; treat it as a last resort.
The letter suffix matters: _M (medium) and _S (small) trade a sliver of quality for size. The newer I-quants (IQ4_XS, IQ3_M, etc.) squeeze out a bit more quality per byte using importance-matrix calibration, at the cost of slightly slower inference on some hardware.
Use GGUF if: you are on a Mac, you are mixing CPU and GPU, you want the widest model selection, or you simply want the path of least resistance. For most readers, the honest answer is "start here."
GPTQ — the GPU-native 4-bit standard
GPTQ is a post-training quantization method introduced by Frantar et al. in 2022 ( arXiv:2210.17323 , later presented at ICLR 2023). Rather than naively rounding every weight, it uses approximate second-order (Hessian) information to quantize weights one column at a time while compensating for the error introduced — a one-shot process that runs in a few GPU-hours even for huge models.
The practical point: GPTQ is weight-only, GPU-only, and fast at inference . It shines when the whole model fits in VRAM and you are serving it through a GPU runtime. It was the dominant 4-bit format on Hugging Face for a long time and is widely supported by serving stacks. Its weakness is that it does not gracefully spill to CPU the way GGUF does — it is an all-in-VRAM format.
Use GPTQ if: your model fits entirely in GPU memory and you want a mature, well-supported 4-bit format for GPU serving.
AWQ — the accuracy-focused challenger
AWQ (Activation-aware Weight Quantization) comes from Lin et al. ( arXiv:2306.00978 , MLSys 2024 best paper). Its insight is clever: not all weights matter equally. A small fraction (~1%) of "salient" weight channels — identified by looking at the activations flowing through them, not the weights themselves — carry an outsized share of the model's quality. AWQ protects those channels by scaling them before quantizing, so the important parts survive 4-bit compression nearly intact.
In practice AWQ often matches or beats GPTQ on accuracy at the same bit width , and it is a first-class citizen in high-throughput serving engines like vLLM . Like GPTQ, it is GPU-oriented and assumes the model lives in VRAM.
Use AWQ if: you are serving on a GPU (especially via vLLM) and want the best quality you can get at 4-bit.
bitsandbytes / NF4 — the 4-bit NormalFloat format from the QLoRA paper (Dettmers et al., arXiv:2305.14314 ). It is the go-to for fine-tuning quantized models on consumer GPUs, and it quantizes on the fly. Great for training; usually not your first pick for pure inference.
EXL2 (ExLlamaV2) — a GPU-only format with variable bitrate (you can target, say, 4.65 bpw). Extremely fast and memory-efficient on NVIDIA cards; beloved by people optimizing for tokens-per-second on a single GPU.
Which bit level should you pick?
This is the question that actually matters day to day. The consensus that has held up across countless community tests:
Q4_K_M is the default for a reason. It is the best quality-per-gigabyte for most people — roughly a 1% perplexity increase over full precision while cutting size by ~65%.
If you have spare memory, step up to Q5_K_M or Q6_K before you reach for a larger model at a lower quant. A 70B at Q4 generally beats a 70B at Q2.
But a bigger model at moderate quant usually beats a smaller model at high quant. A 70B at Q4_K_M will typically outperform a 13B at Q8. When in doubt, go bigger-model-lower-quant rather than smaller-model-higher-quant — down to about Q4. Below Q3, the math flips.
Avoid Q2 unless you have no other choice. The quality cliff steepens fast below ~3 bits.
Here is what those choices look like in real memory for a 70B model (weights only; add headroom for context/KV cache):
Why low-bit quantization hurts reasoning more than recall
Here is the nuance the bit-level table hides: quantization does not degrade all tasks equally. The damage is uneven — and understanding why will keep you from picking the wrong quant for your actual workload.
Mechanically, quantization adds a small amount of rounding noise to every weight. For a single-step task — recalling a fact, answering a multiple-choice question, finishing a sentence — that noise is usually swamped by the model's confidence; the right answer still wins. This is why aggressive quants often look almost lossless on short benchmarks like HellaSwag. The signal survives.
For multi-step reasoning — chains of logic, code that has to compile, tool calls that must be formatted exactly, long-context tasks where attention has to track many details at once — that same small noise compounds . A tiny error in step one shifts step two, which shifts step three, and by the end the chain has drifted off course. The model did not "forget" anything; it accumulated rounding error across a long path. Long context makes it worse, because attention has more competing details to keep straight and low-bit weights blur the distinctions.
This matches what owners report. The recurring consensus on r/LocalLLaMA is that 4-bit is fine for chat but worth a second thought before you trust it with precise, multi-step work:
"Fine for chatting but don't use it for actual data work."
— u/autonomousdev_, after Q8 caught two edge cases in an extraction task that Q4 missed
"Aggressive quants (<Q6) do fine on benchmarks that often use 4k contexts. They fall apart in real life, when you need to handle 30k or more."
— u/ClearApartment2627
Casual chat, brainstorming, summarizing, simple Q&A: Q4_K_M is genuinely hard to tell apart from full precision. Save the VRAM.
Coding, agentic tool use, structured extraction, long documents: step up to Q5_K_M or Q6_K if it fits. The extra bits buy reliability exactly where quantization noise does the most damage.
Watch your KV cache, too. Quantizing the context cache to save memory is a separate knob from weight quantization — and at long context it often hurts more than the weight quant does. If a model falls apart past ~16–32k tokens, suspect KV-cache quantization before you blame the weights.
None of this contradicts "Q4_K_M is the default." It refines it: Q4 is the default for general use, and a deliberate step up is cheap insurance for the workloads where every step has to land.
Sources & how we researched this
This explainer synthesizes the original method papers — GPTQ ( Frantar et al., 2022 ), AWQ ( Lin et al., 2023 ), and QLoRA/NF4 ( Dettmers et al., 2023 ) — alongside the llama.cpp documentation and quantization tables for the GGUF k-quant bits-per-weight figures. The bit-level guidance reflects the broad, durable consensus from the r/LocalLLaMA community and maintainer benchmarks, not first-hand testing on our part. VRAM figures are weights-only estimates rounded for clarity; real-world usage adds memory for the KV cache and context window, which we cover separately.
How much VRAM do you actually need to run a 70B model locally?
Ollama vs LM Studio vs llama.cpp: which runtime should you actually use?
The cheapest way to run a 70B model locally
Get the Vetted Consumer newsletter
Reviews, buying advice, and field notes. Delivered monthly.
Honest local-LLM hardware advice — cross-checked across reviews and the people who own it.
