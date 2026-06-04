---
source: "https://muhammadraza.me/2026/gguf-vs-mlx-decision-guide/"
hn_url: "https://news.ycombinator.com/item?id=48401571"
title: "GGUF vs. MLX when to choose what when running Local LLMs"
article_title: "GGUF vs MLX: A Decision Guide, Not Another Benchmark | Muhammad"
author: "mr_o47"
captured_at: "2026-06-04T18:54:01Z"
capture_tool: "hn-digest"
hn_id: 48401571
score: 2
comments: 0
posted_at: "2026-06-04T17:11:55Z"
tags:
  - hacker-news
  - translated
---

# GGUF vs. MLX when to choose what when running Local LLMs

- HN: [48401571](https://news.ycombinator.com/item?id=48401571)
- Source: [muhammadraza.me](https://muhammadraza.me/2026/gguf-vs-mlx-decision-guide/)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T17:11:55Z

## Translation

タイトル: GGUF と MLX ローカル LLM を実行するときに何を選択するか
記事のタイトル: GGUF 対 MLX: 別のベンチマークではなく、意思決定ガイド |ムハンマド
説明: Mac 上でローカル LLM を実行すると、GGUF ファイルと同じモデルの MLX ビルドのどちらかを選択するように求められ続けます。実際にどのように決定するかは次のとおりです。

記事本文:
GGUF 対 MLX: 別のベンチマークではなく、意思決定ガイド |ムハンマド・ムハンマド
GGUF 対 MLX: 別のベンチマークではなく、意思決定ガイド
数週間ごとに、誰かが同じモデルの GGUF ビルドと MLX ビルドをダウンロードして両方を実行し、1 秒あたりのトークン数カウンターのスクリーンショットを撮り、一方の形式が勝っていることの証拠として投稿します。返信は途中で分かれました。スレッドの半分は MLX が明らかに速いと言っており、残りの半分はテストが不正だったと言っています。
どちらも正しいのですが、それが問題なのです。画面上の番号は実際のものであり、実際に待っている番号でもありません。そして、選択すべきフォーマットは、とにかくその数字に関係するものではありませんでした。
私は、自分のマシン上で、またローカル推論を行うクライアントのために、この決定を何度も繰り返してきたので、比較表に誰も入れていない部分を書き留めておきたいと思います。GGUF と MLX の決定は 5 つの質問で決まり、そのうちの 1 つだけが速度に関するものです。
実際にどちらを選択しているか
GGUF は、llama.cpp プロジェクトのファイル形式です。 1 つのファイルには量子化された重み、トークナイザー、チャット テンプレート、メタデータが保持されており、それを読み込むことができるランタイムによってモデルが実行されます。これには、llama.cpp 自体、Ollama、LM Studio、KoboldCPP、およびその他のいくつかが含まれます。 CPU、NVIDIA、AMD、Apple Metal、さらに忍耐強くいれば Raspberry Pi など、基本的にあらゆるもので動作します。このフォーマットの最大のポイントは移植性です。
MLX はファイル形式ではありません。これは Apple の配列フレームワークであり、Apple Silicon 専用に構築された PyTorch にほぼ相当します。 MLX モデルは、セーフテンソル ファイルと、ランタイムが直接読み取る構成のディレクトリです。 mlx_lm.convert を使用して、1 つのコマンドでモデルの変換と量子化を行います。問題は名前にあります。MLX は Apple Silicon 上で動作し、他の場所では動作しません。
その前にクリアしておきたいことが 1 つあります

これは比較の半分に表示され、時代遅れであるため、さらに詳しく説明します。GGUF は賢明な混合精度の量子化を行うのに対し、MLX はフラットな均一 4 ビットに固執していると言われています。前半は本当です。後半はそうではありません。 Apple は、MLX を使用した大規模な言語モデルの実行に関する WWDC25 セッションで、埋め込み層と出力層を 6 ビットに保ち、モデルの残りの部分を 4 ビットに保つトリックなど、層ごとの混合精度について説明しました。 MLXならそれができる。 Hugging Face の周りに浮遊する MLX ビルドのほとんどは問題にならないため、実際には、GGUF の混合精度と均一な MLX クオントを比較することがよくあります。他の人の品質ベンチマークを読むときに知っておく価値があります。
画面上の数字はあなたに嘘をついています
それはあなたが見つけるであろうベンチマークのほとんどを汚染するので、簡単に迂回してください。
テキストのストリーミング中にランタイムが出力する 1 秒あたりのトークン数の数値は、モデルが新しいトークンを発行する速度であるデコード速度を測定します。これには、モデルが何かを言う前にプロンプ​​トを読むのに費やす時間、プレフィルは含まれません。あまり重要ではない短いプロンプトによるおしゃべりなやりとりの場合。ツールの出力、ファイルのチャンク、およびシステム プロンプトを毎ターン詰め込むエージェントの場合、待機内容のほとんどは事前入力であり、ストリーミング カウンターはそれを認識しません。
r/LocalLLaMA で話題になったベンチマークの記事があります。そこでは、作者の UI が MLX では GGUF の 1 秒あたりのトークンがほぼ 2 倍であることを誇らしげに報告しており、その後、実際の実時間では、実際のタスクのほとんどで GGUF が最初に終了しました。同じマシン、同じモデル。カウンターは間違っていませんでした。重要な質問とは別の質問に答えただけです。
この投稿の残りの部分では、このことを頭の中に入れておいてください。以下で、あるフォーマットが「速い」と言う場合、それは壁時計のことを意味します。

k は実際のワークロード上の数値であり、トークンのストリーミング中にスクロールして通過する数値ではありません。
実際にそれを決める5つの質問
1. RAM に対してモデルの大きさはどれくらいですか?
これは多くの議論を静かに解決する問題です。トークンの生成は、コンピューティングではなくメモリ帯域幅によって制限されます。 1 つのトークンを発行するには、GPU がメモリからモデル全体を読み取る必要があります。約 273 GB/秒の帯域幅を持つ M4 Pro では、約 17 GB の重さの 4 ビット 27B モデルは、実行するソフトウェアに関係なく、1 秒あたり 16 トークン近くに達します。 MLX はハードウェアが許可する速度よりも速くバイトをフェッチすることはできません。また、llama.cpp も同様です。
したがって、ユニファイド メモリの大部分を占める大規模なモデルの場合、形式は速度にはほとんど関係ありません。彼らは両方とも同じ壁にぶつかりました。興味深い違いは、およそ 8 ～ 14B 未満の小規模なモデルに現れ、モデルは快適にフィットし、ボトルネックは帯域幅からフレームワークのオーバーヘッドに移行します。ここでは、MLX のよりタイトな Apple 固有のカーネルが先行しており、多くの場合、シングル ユーザー デコードでは 15 ～ 40 パーセントの範囲で、フレームワークの効率を最も重視する非常に小規模なモデルではさらに広い範囲でパフォーマンスが向上します。
小型モデル、キビキビと動きたい: MLX には本物が備わっています。かろうじて適合する大きなモデル: スピードは重要なので、他の 4 つの質問を選択してください。
2. これを Mac 以外の場所で実行する必要が生じることはありますか?
同じアーティファクトを Linux ボックス、クラウド GPU、またはチームメイトの Apple 以外のマシンで実行する必要がある可能性がある場合は、GGUF が必要になります。同じファイルがそれらすべての間で移動します。 MLX は Apple Silicon を離れることはありません。 MLX を唯一のビルドとして出荷し、CUDA フォールバックが必要な場合は、プレッシャーの下で再量子化することになります。
これは他のほとんどすべてをオーバーライドします。移植性はパフォーマンス機能ではありませんが、それがなくなると最も恋​​しくなる機能です。
3. あなたの仕事は何ですか

実際の負荷はどのようになりますか？
「どのモデル」ではなく、トラフィックの形状です。具体的には、インプットとアウトプットの比率です。
モデルに多くのフィードを与え、少しだけ要求するワークロード (分類、短い応答を持つツール呼び出しエージェント、大量に注入されたコンテキストを持つ RAG) は、GGUF に傾いています。 llama.cpp には、より実戦テストされたプロンプト キャッシングと FlashAttendant があり、MLX のプレフィックス キャッシングは、歴史的に、特に新しいハイブリッド アテンション モデルにおいて、この 2 つのうちの信頼性が低くなっています。プレフィルが壁時計を支配すると、その成熟度が勝ちます。
短いプロンプトを受け取り、大量の結果を生成するワークロード (要約、長文チャット、ブレーンストーミング) は、MLX に傾いています。モデルがプレフィルを過ぎてトークンをストリーミングするだけになると、MLX のデコードの利点はさらに増大し、応答が長ければ長いほど、その効果は大きくなります。
コンテキストのサイズと応答の長さの両方に依存するクロスオーバー ポイントがあります。プロンプトが小さい場合、MLX は、より高速なデコードが低速なプリフィルを補うまでに数百の出力トークンを必要とします。コンテキストのトークンが数千ある場合、さらに数百が必要になります。エージェントの応答が 150 トークンであり、そのコンテキストが増加し続ける場合、あなたはそのクロスオーバーの間違った側に住んでおり、GGUF の方がより良いコールです。
4. トレーニングしたいですか、それともただ走るだけですか?
GGUF は推論フォーマットです。ダウンロードして実行する、それが関係です。微調整したい場合は、セーフテンソルに変換し直し、GPU を見つけて作業を行い、再度前方変換します。
MLX は完全なフレームワークです。 LoRA または QLoRA を Mac 上で直接微調整したり、アダプターをマージしたり、小規模なドラフト モデルを使用して投機的デコードを実行したりすることが、すべてネイティブに実行できます。ローカルに移行する理由の一部がモデルを提供するだけでなく実際にモデルを適応させることである場合、MLX が Apple Silicon の唯一の本格的なオプションであり、この質問だけで全体の妥当性を決定できます。

g.
5. エコシステムと正確なフィット感をどの程度気にしますか?
ここには GGUF の 2 つの実用的な利点があります。まず、対象範囲: すべてのオープン モデルは、不明瞭なモデルも含め、リリースから数時間以内に GGUF ビルドを取得します。 MLX のカバレッジは人気のあるモデルには適していますが、その他のモデルには遅れます。 2 番目に、粒度です。 GGUF では、クオンツ レベル、Q4_K_M、Q5_K_M、Q6_K、I クオントなどの長いラダーが提供されるため、ちょうど 16 GB で作業する場合、通常は適合するクオントを見つけることができます。 MLX ビルドはほとんどが 4 ビットと 8 ビットで公開されているため、必要な品質にはほんの少し小さすぎる 4 ビットや、適合しない 8 ビットが得られることがあります。
MLX 側の利点: Apple は llama.cpp が追いつく前に MLX でメタル抽象化を出荷するため、新しい Apple ハードウェア機能のサポートを最初に取得する傾向があります。
5 つの質問を順番に並べると、ほとんどの決定は約 10 秒以内に下されます。
Apple Silicon 以外のものを今すぐに実行する必要がありますか、それとも後で実行する必要がありますか? → ググフ 。ここで停止します。移植性が優先されます。
Appleシリコンを使い続けます。デバイス上で微調整またはトレーニングしたいですか? →MLX。
推論のみ。ワークロードは出力が少なく、事前入力が多い (エージェント、RAG、分類) ですか? → ググフ 。
長い出力、インタラクティブ、シングルユーザー、体感できる遅延? →MLX。
狭い RAM に適合する正確なクォントが必要ですか、それともリリースされたばかりのモデルや無名なモデルを実行する必要がありますか? → ググフ 。
まだ決まっていませんか？ → ググフ 。これは保守的なデフォルトです。スループットが制約になる場合は、それを出荷し、後で MLX ビルドの A/B を行います。
簡単に言うと、GGUF は迷ったときに選ぶものです。後悔しにくいものだからです。 MLX は、ハードウェアを所有し、シングル ユーザーで実行し、長時間の出力やオンデバイス トレーニングでのスループットなどの特別な理由がある場合に選択するものです。
選択したら、クオンツレベルを選択します
フォーマットはデの半分です

シシオン。ビット幅は残りの半分であり、デフォルトは適切ですが、常に正しいとは限りません。
GGUF の場合は Q4_K_M、MLX の場合は 4 ビットから開始します。 Q4_K_M がコミュニティのデフォルトになっているのには理由があります。ほとんどのテンソルを 4 ビットに保ち、品質に敏感なテンソル、つまり一部のレイヤーでのアテンション値の重みとフィードフォワード下方投影を 6 ビットに引き上げます。これにより、小さなサイズのコストでフラット 4 ビット クオントよりも優れた品質が維持されます。報告されている MMLU 上の FP16 に対する品質低下はモデルに依存しますが、小さいです。大きなモデルではポイントよりかなり下で、8B 未満ではポイントなどに向かって徐々に上昇し、均一な 4 ビット MLX ビルドではさらにもう少し大きくなります。 30B を超えるモデルでは、そのギャップはノイズになります。 8B 未満のもの、特に注意の精度が重要なコーディング タスクでは、それが目に見えてわかります。GGUF Q4_K_M を使い続けるか、MLX 6 ビットに移行するかです。これにより、約 30% 大きなファイルのギャップが縮まります。
RAM が本当に不足している場合は、重要度マトリックスを備えた GGUF の I-quants が、低ビット幅でのバイトあたりの品質のチャンピオンとなります。コストとしては、CPU でのデコードが遅くなるため、速度を追求する場合よりも、限られたメモリにモデルを圧迫する場合に意味があります。
形式に関係なく、ルールが 1 つあります。自分のタスクの品質を測定せずに、およそ 3 ビットを下回らないようにしてください。集約されたベンチマークは、そこで実際に何が表示されるかを予測しなくなります。
結果をひっくり返す 2 つの罠
M1とM2のbf16トラップ。多くの MLX ビルドは bf16 として出荷されており、M1 および M2 では、そのデータ型は fp16 が行うアクセラレーション パスを取得しません。プレフィル中、これらの重みは加速されずに実行され、入力トークンごとにペナルティが乗算されます。これが、一部の「MLX が遅い」レポートが古いハードウェアから発生する理由の 1 つです。修正するには、 --dtype float16 を使用して 1 分間再変換します。オンの場合

M1 または M2 と MLX が遅く感じられる場合は、フォーマットのせいにする前にこれを確認してください。
実際の変数はキャッシュです。私がこれまでに見たランタイム間の最大の変動は、GGUF 対 MLX に関するものではなく、プロンプトと KV キャッシュがそのランタイム上のそのモデルで実際に機能するかどうかに関するものでした。会話全体を毎ターン再処理するランタイムは、形式に関係なく、プレフィックスをキャッシュするランタイムに負けます。コミットする前に、実際のコンテキストの長さを使用してキャッシュをテストしてください。ストリーミング カウンターはキャッシュによって修正される部分を決して測定しないため、ストリーミング カウンターがそれについて通知することを信頼しないでください。
1 行バージョンが必要な場合: GGUF が保守的なデフォルトであり、不確かな場合、移植性が必要な場合、または特定のクオントが必要な場合は、GGUF を使用する必要があります。 Apple Silicon に縛られている場合、長い出力でシングルユーザーの対話型ワークロードを実行している場合、またはすでに所有しているマシンで微調整したい場合は、MLX を使用してください。
また、ラップトップではなくチーム用にこれを選択する場合は、それをアーキテクチャ上の決定として扱ってください。標準化する形式によって、スタックが存続する限りモデル カバレッジ、フォールバック オプション、サービス設定が決まり、事後的にフリートを再量子化するのは、避けられる一週間の作業です。

[切り捨てられた]

## Original Extract

If you run local LLMs on a Mac, you keep getting asked to choose between a GGUF file and an MLX build of the same model. Here is how to actually decide, in f...

GGUF vs MLX: A Decision Guide, Not Another Benchmark | Muhammad Muhammad
GGUF vs MLX: A Decision Guide, Not Another Benchmark
Every few weeks someone downloads the GGUF build and the MLX build of the same model, runs both, screenshots the tokens-per-second counter, and posts it as proof that one format wins. The replies split down the middle. Half the thread says MLX is obviously faster, the other half says the test was rigged.
They are both right, which is the problem. The number on the screen is real and it is also not the number you actually wait for. And the format you should pick was never really about that number anyway.
I have gone through this decision enough times now, on my own machine and for clients standing up local inference, that I want to write down the part nobody puts in the comparison tables: GGUF versus MLX is a five-question decision, and only one of those questions is about speed.
What you are actually choosing between
GGUF is the file format from the llama.cpp project. One file holds the quantized weights, the tokenizer, the chat template, and the metadata, and any runtime that can load it will run the model. That includes llama.cpp itself, Ollama, LM Studio, KoboldCPP, and a handful of others. It runs on basically everything: CPU, NVIDIA, AMD, Apple Metal, even a Raspberry Pi if you are patient. Portability is the whole point of the format.
MLX is not a file format. It is Apple’s array framework, the rough equivalent of PyTorch built specifically for Apple Silicon. An MLX model is a directory of safetensors files plus a config that the runtime reads directly. You convert and quantize a model in one command with mlx_lm.convert . The catch is in the name: MLX runs on Apple Silicon and nowhere else.
One thing worth clearing up before we go further, because it shows up in half the comparisons and it is out of date: people say GGUF does clever mixed-precision quantization while MLX is stuck on flat uniform 4-bit. The first half is true. The second half is not. Apple walked through per-layer mixed precision in their WWDC25 session on running large language models with MLX, including the trick of keeping the embedding and output layers at 6-bit while the rest of the model sits at 4-bit. MLX can do it. It is just that most of the MLX builds floating around Hugging Face do not bother, so in practice you often are comparing GGUF’s mixed precision against a uniform MLX quant. Worth knowing when you read someone else’s quality benchmark.
The number on the screen is lying to you
Quick detour, because it poisons most of the benchmarks you will find.
The tokens-per-second figure your runtime prints while text is streaming measures decode speed, the rate at which the model emits new tokens. It does not include prefill, the time the model spends reading your prompt before it says anything. For a chatty exchange with a short prompt that does not matter much. For an agent that stuffs tool output, a chunk of a file, and a system prompt into every turn, prefill is most of what you wait for, and the streaming counter never sees it.
There is a benchmark writeup that made the rounds on r/LocalLLaMA where the author’s UI proudly reported nearly twice the tokens per second on MLX as on GGUF, and then the actual wall-clock time had GGUF finishing first on most of the real tasks. Same machine, same model. The counter was not wrong. It was just answering a different question than the one that mattered.
Keep that in your head for the whole rest of this post. When I say one format is “faster” below, I mean wall-clock on a real workload, not the number that scrolls past while tokens stream.
Five questions that actually decide it
1. How big is the model relative to your RAM?
This is the question that quietly settles a lot of arguments. Token generation is bounded by memory bandwidth, not compute. To emit one token the GPU has to read the entire model out of memory. On an M4 Pro with roughly 273 GB/s of bandwidth, a 4-bit 27B model weighing about 17 GB caps out near 16 tokens per second no matter what software you run. MLX cannot fetch bytes faster than the hardware allows, and neither can llama.cpp.
So for large models, the ones that fill most of your unified memory, the format barely matters for speed. They both hit the same wall. The interesting differences show up on smaller models, under roughly 8 to 14B, where the model fits comfortably and the bottleneck shifts from bandwidth to framework overhead. That is where MLX’s tighter, Apple-specific kernels pull ahead, often in the 15 to 40 percent range on single-user decode, and wider still on very small models that lean hardest on framework efficiency.
Small model, want it snappy: MLX has something real to offer. Big model that barely fits: pick on the other four questions, because speed is a wash.
2. Will this ever need to run somewhere other than a Mac?
If there is any chance the same artifact has to run on a Linux box, a cloud GPU, or a teammate’s non-Apple machine, you want GGUF. The same file moves between all of them. MLX does not leave Apple Silicon, full stop. If you ship MLX as your only build and then need a CUDA fallback, you are re-quantizing under pressure.
This one overrides almost everything else. Portability is not a performance feature, but it is the feature you miss most when it is gone.
3. What does your workload actually look like?
Not “what model,” but the shape of the traffic. Specifically the ratio of input to output.
Workloads that feed the model a lot and ask for a little (classification, tool-calling agents with short replies, RAG with a big injected context) lean toward GGUF. llama.cpp has more battle-tested prompt caching and FlashAttention, and MLX’s prefix caching has historically been the less reliable of the two, especially on newer hybrid-attention models. When prefill dominates the wall clock, that maturity wins.
Workloads that take a short prompt and generate a lot (summaries, long-form chat, brainstorming) lean toward MLX. Once the model is past prefill and just streaming tokens, MLX’s decode advantage compounds, and the longer the reply the more it pays off.
There is a crossover point that depends on both context size and reply length. With a small prompt, MLX needs a couple hundred output tokens before its faster decode makes up for slower prefill. With a few thousand tokens of context, it needs several hundred more. If your agent’s replies are 150 tokens and its context keeps growing, you are living on the wrong side of that crossover, and GGUF is the better call.
4. Do you want to train, or just run?
GGUF is an inference format. You download it, you run it, that is the relationship. If you want to fine-tune, you convert back to safetensors, find a GPU, do the work, and convert forward again.
MLX is a full framework. You can fine-tune with LoRA or QLoRA directly on the Mac, merge adapters, and run speculative decoding with a small draft model, all natively. If part of your reason for going local is to actually adapt models and not just serve them, MLX is the only serious option on Apple Silicon, and this question alone can decide the whole thing.
5. How much do you care about ecosystem and exact fit?
Two practical edges for GGUF here. First, coverage: every open model gets GGUF builds within hours of release, including the obscure ones. MLX coverage is good for popular models and lags for everything else. Second, granularity. GGUF gives you a long ladder of quant levels, Q4_K_M, Q5_K_M, Q6_K, the I-quants, and so on, so when you have exactly 16 GB to work with you can usually find a quant that fits. MLX builds are mostly published at 4-bit and 8-bit, so you sometimes get a 4-bit that is a hair too small for the quality you want and an 8-bit that will not fit.
The edge on MLX’s side: it tends to get support for new Apple hardware features first, because Apple ships the metal abstraction in MLX before llama.cpp catches up.
Put the five questions in order and most decisions fall out in about ten seconds.
Need to run on anything other than Apple Silicon, now or later? → GGUF . Stop here, portability wins.
Staying on Apple Silicon. Do you want to fine-tune or train on-device? → MLX .
Inference only. Is your workload short-output and prefill-heavy (agents, RAG, classification)? → GGUF .
Long outputs, interactive, single user, latency you can feel? → MLX .
Need a precise quant to fit tight RAM, or running a just-released or obscure model? → GGUF .
Still undecided? → GGUF . It is the conservative default. Ship it, and A/B an MLX build later if throughput becomes the constraint.
The short version: GGUF is what you pick when you are not sure, because it is the one that is hard to regret. MLX is what you pick when you own the hardware, run single-user, and have a specific reason, throughput on long outputs or on-device training, to want it.
Once you have picked, pick a quant level
The format is half the decision. The bit width is the other half, and the defaults are good but not always right.
Start at Q4_K_M for GGUF or 4-bit for MLX. Q4_K_M is the community default for a reason. It keeps most tensors at 4-bit, then bumps the quality-sensitive ones to 6-bit: the attention value weights and the feed-forward down-projection, on a portion of the layers. That holds quality better than a flat 4-bit quant at a small size cost. The reported quality loss against FP16 on MMLU is model-dependent but small: well under a point on a big model, creeping up toward a point or so on something under 8B, and a little more again for a uniform 4-bit MLX build. On a 30B-plus model that gap is noise. On something under 8B, especially on coding tasks where attention precision matters, it is visible, and you have two outs: stay on GGUF Q4_K_M, or move to MLX 6-bit, which closes the gap for roughly a 30 percent larger file.
If RAM is genuinely tight, GGUF’s I-quants with an importance matrix are the quality-per-byte champions at low bit widths. The cost is slower decode on CPU, so they make more sense when you are squeezing a model onto limited memory than when you are chasing speed.
One rule regardless of format: do not drop below roughly 3-bit without measuring quality on your own task. The aggregate benchmarks stop predicting what you will actually see down there.
Two traps that will flip your results
The bf16 trap on M1 and M2. A lot of MLX builds ship as bf16, and on the M1 and M2 that data type does not get the accelerated path that fp16 does. During prefill those weights run un-accelerated and the penalty multiplies across every input token, which is part of why some “MLX is slow” reports come from older hardware. The fix is a one-minute reconvert with --dtype float16 . If you are on an M1 or M2 and MLX feels sluggish, check this before you blame the format.
Caching is the real variable. The biggest swings I have seen between runtimes were not about GGUF versus MLX at all, they were about whether prompt and KV caching actually worked for that model on that runtime. A runtime that reprocesses the full conversation every turn will lose to one that caches the prefix, regardless of format. Test caching with your real context lengths before you commit, and do not trust the streaming counter to tell you about it, because it never measures the part that caching fixes.
If you want the one-line version: GGUF is the conservative default, and you should reach for it whenever you are uncertain, need portability, or want a specific quant. Reach for MLX when you are locked to Apple Silicon, run single-user interactive workloads with long outputs, or want to fine-tune on the machine you already own.
And if you are choosing this for a team rather than a laptop, treat it as the architecture decision it is. The format you standardize on shapes your model coverage, your fallback options, and your serving setup for as long as the stack lives, and re-quantizing a fleet after the fact is the kind of avoidable week of work I keep getting hir

[truncated]
