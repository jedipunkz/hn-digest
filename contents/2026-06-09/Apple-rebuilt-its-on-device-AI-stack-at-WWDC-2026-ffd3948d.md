---
source: "https://ziraph.com/blog/apple-on-device-ai-wwdc-2026"
hn_url: "https://news.ycombinator.com/item?id=48458183"
title: "Apple rebuilt its on-device AI stack at WWDC 2026"
article_title: "Apple rebuilt its on-device AI stack at WWDC 2026 - Ziraph blog"
author: "ABS"
captured_at: "2026-06-09T10:19:07Z"
capture_tool: "hn-digest"
hn_id: 48458183
score: 2
comments: 1
posted_at: "2026-06-09T08:16:05Z"
tags:
  - hacker-news
  - translated
---

# Apple rebuilt its on-device AI stack at WWDC 2026

- HN: [48458183](https://news.ycombinator.com/item?id=48458183)
- Source: [ziraph.com](https://ziraph.com/blog/apple-on-device-ai-wwdc-2026)
- Score: 2
- Comments: 1
- Posted: 2026-06-09T08:16:05Z

## Translation

タイトル: Apple、WWDC 2026 でオンデバイス AI スタックを再構築
記事タイトル: Apple、WWDC 2026 でオンデバイス AI スタックを再構築 - Ziraph ブログ
説明: WWDC 2026 には新しいチップはありませんが、Apple シリコン上で AI がどのように動作するかが静かに再構築されます。コア AI、新しいモデル形式、第 3 世代のファウンデーション モデル、Google Cloud の NVIDIA GPU で動作する主力クラウド モデルです。ロードマップ、大きな変更点、そして奇妙な発見。

記事本文:
Apple、WWDC 2026 でオンデバイス AI スタックを再構築 - Ziraph ブログ
ベータ版（ソフト）近日公開予定！お知らせリストを購読して、Web サイトとベータ版がオープンする瞬間を聞いてください ↗
Apple、WWDC 2026 でオンデバイス AI スタックを再構築
新しいチップはありません。しかし、機械学習が Apple シリコン上で実行される方法の静かな再構築、つまり新しい推論フレームワーク、新しいモデル形式、新世代のオンデバイス モデル、そして他社の GPU 上で実行される主力クラウド モデルです。興味深い部分は、基調講演のスライドには載っていない部分です。
· 2026 年 6 月 9 日 · 10 分で読めます
WWDC 2026 では新しいシリコンは登場しませんでした。代わりに、Apple シリコン上で AI がどのように実行されるかを構造的に再構築しました。
新世代のオンデバイス モデル、
そしてクラウドに対する姿勢が著しく異なります。
どれも見出しではありませんでした。見出しは消費者向けの機能でした。しかし、開発者ドキュメント、セッション コード、および 1 つの機械学習研究投稿を合計すると、基調講演よりも明確なロードマップが得られ、加えて、本当に奇妙な点がいくつかあります。
私はこのレイヤーを注意深く読み、そのプロファイラーを構築しています。そのため、目立った点は次のとおりです。大きな変更点、微妙な変化、そして信じる前に再確認する必要があった結果です。基本的なルールを 1 つ挙げておきます。以下の内容はすべて Apple 自身のドキュメント、WWDC セッション ページ、および研究投稿からのものであり、重要な場合には引用されています。 Apple の言葉ではなく、個人の開発者の主張やフォーラムの記事の内容については、私はそう言います。 Apple が何も言わない場合は、私もそう言います。
そして最大の注意点: 私はヨーロッパにいるので、夜通し見て、読んで、調べて過ごしました - 睡眠不足のために何か間違ったことをしたと思います。 :-)
大きな変化: Core AI がニューラル ネットワークの Core ML を置き換える
10 年間、Core ML は「iPhone 上でモデルを実行する」ための答えでした。

「WWDC 2026 で Apple は Core AI を導入しました。その枠組みは追加ではなく引き継ぎです。Core AI のドキュメントでは、古いケースが Core ML に送り返されています。」
「アプリでデシジョン ツリーや表形式特徴量エンジニアリングなど、ニューラル ネットワーク以外のモデル タイプを使用している場合は、Core ML を参照してください。」 - Apple、コア AI ドキュメント
そして、Core ML のドキュメントには新しいドキュメントが送信されます。
「アプリが最新のアーキテクチャと推論技術を使用して AI モデルを統合している場合は、Core AI を参照してください。」 - Apple、Core ML ドキュメント
一緒に読んでください。これは分割です。Core ML は、従来の非ニューラル機械学習 (デシジョン ツリー、表形式の機能) に絞り込まれますが、ニューラル ネットワークとトランスフォーマーは、Apple が製品自体の背後にあるエンジンと説明する Core AI に移行します。
「Core AI により、アプリは CPU、GPU、ニューラル エンジン全体で最新のモデル アーキテクチャと推論技術を使用できるようになります。」 - Apple、コア AI ドキュメント
微妙な特徴はツールにあります。 Apple の新しい Core AI デバッグ ゲージには 1 行の制限があります。
「このゲージは Core ML フレームワークをサポートしていません。」 - Apple、Core AI デバッグ ゲージのドキュメント
新しい計測器は古いフレームワークを考慮していないだけです。 Core ML は非推奨ではありません。その API はそのままであり、そこには実際の後方互換性の価値があります。しかし、重心とツールへの投資は移動しました。
新しいアーティファクト: .aimodel バンドル
Core AI は新しいオンディスク形式 .aimodel を同梱していますが、最初に奇妙なのは、これがファイルではないことです。ディレクトリです。 Apple のオープンな coreai-models リポジトリは、それを全体として 1 つとして扱います。Python エクスポーターはディレクトリのみの呼び出しで古いリポジトリを削除し、Swift ランタイムはそれを「 .aimodel ディレクトリ」として解決します。周囲のモデル バンドル内には、プレーンな JSON の metadata.json (スキーマ バージョン) が含まれています。

0.2 )、モデルの種類 (LLM、VLM、拡散、セグメンター)、トークナイザー、語彙サイズ、コンテキストの長さ、圧縮プリセット、およびどのファイルがモデルであるかを記録します。その JSON は文書化されており、解析可能です。ウェイト ペイロード自体 (正確なテンソルごとのビット幅を示す部分) は不透明なフレームワーク呼び出しによって記述されており、そのバイト レイアウトは私が見つけたどこにも公開されていません。したがって、形式はハーフオープンです。つまり、文書化されていない BLOB を読み取り可能なマニフェストで囲んだものです。
モデルは、新しい Python ツール - 圧縮用の Core AI Optimization ( coreai-opt 、 coremltools の後継)、および PyTorch から形式に直接エクスポートするための Core AI PyTorch Extensions ( coreai-torch ) で準備され、その後、オプションで xcrun coreai-build を使用して事前にコンパイルされ、アーキテクチャごとの .aimodelc アセットにコンパイルされます。圧縮メニューは GGUF の世界よりも幅広く、2、4、8 ビットの整数重み付けが可能です。 FP8 (E4M3) および FP4 (E2M1) を含む浮動小数点マイクロフォーマット。ブロックスケールのMXFP8。 1 から 8 ビットまでのパレット化。あるフォーラム読者 (HN、意見) は、Apple も w4a8 / w4a16 のようなアクティベーション量子化を推進していると指摘しました。 Apple のインストールベースを考えると、Apple が恩恵をもたらすフォーマットが、最終的には 100B 未満のモデルをすべての人に提供する方法を形作ることになる可能性があります。
ハードウェアは、Matmul が GPU に移動したことを伝えます
新しいチップはありませんが、WWDC 2026 では M5 と A19 GPU のストーリーが明確になり、今週の最も明確なハードウェアシグナルです。 Apple の M5/A19 技術講演より:
「ニューラル アクセラレータは、行列乗算専用に構築された M5 の専用ハードウェアです。ニューラル アクセラレータは、ALU、レイトレーシングなどの他の GPU パイプラインと並んで各シェーダ コアに組み込まれています。各シェーダ コアには独自のニューラル アクセラレータがあります。」 - Apple、「M5 および A19 GPU で機械学習ワークロードを高速化」
Apple の数字: 行列の乗算

アプリケーションは最大 4 ～ 8 倍高速化され、LLM は最初のトークンまでの時間が最大 4 倍高速化され (プレフィル、コンピューティング依存型)、トークン生成は最大 25% 高速化 (デコード、メモリ依存型) になります。そして、その下のフレームは、屋根のラインであるため、人々が認識する局所的な推論の 1 つです。現在、Apple のメタル パフォーマンス プリミティブ ガイドには次のように記載されています。
「演算強度が低い GEMM はメモリに制約されたワークロードであり、演算強度が高い GEMM は計算に制約されたワークロードであり、カーネル パフォーマンスのルーフライン モデルの基礎を形成します。」 - Apple、メタル パフォーマンス プリミティブ プログラミング ガイド
コンピューティングに依存したプリフィル、メモリに依存したデコード: プリフィルとデコードの分割は、単なるコミュニティのヒューリスティックではなく、Apple 独自の言語になりました。
2 つ目のテルは、スライドではなくコードに隠されています。 coreai-models ソースは、モデルの優先計算ユニットをそのグラフ構造から推測します。チャンク化された静的形状のグラフはニューラル エンジンを優先します。動的な形状のグラフでは GPU が優先されます。これは、Apple が何年もほのめかしてきた分岐を静かに形式化したものです。つまり、静的で古典的な形状の作業用のニューラル エンジンと、トランスフォーマ ママル用の GPU (現在は各シェーダ コア内にニューラル アクセラレータが搭載されています) です。強調する価値があります。これは、エクスポート時にエンコードされるモデルの優先ターゲットであり、特定の実行が実際にどこで実行されるかを保証するものではありません。
モデル: AFM 3 と帯域幅の壁
Apple はまた、第 3 世代の Foundation モデルも発表しました。デバイス上: 30 億パラメータの高密度モデル (「AFM 3 Core」) と 200 億パラメータの疎な専門家混合モデル (「AFM 3 Core Advanced」、ネイティブにマルチモーダルで、一度にわずか 10 ～ 40 億のパラメータを有効にし、最も性能の高い Apple シリコンにゲート制御されている)。
興味深いのはメモリのセクションで、Apple は制約を明確に述べています。

:
「完全なモデルはフラッシュ メモリ (NAND) に保存されています。」 ... 「NAND から DRAM への帯域幅が遅すぎるため、トークンごとに重みを交換できません。」 - Apple 機械学習研究、AFM 3
それは、Apple が、ローカル LLM ランナーが必ずぶつかる壁、つまり常駐を維持するには大きすぎるモデルの代金をトークンごとに移動したバイト数で支払うことについて説明していることです。彼らの答えは、常にアクティブな「共有エキスパート」と入力依存の「ルーティングされたエキスパート」の割合が高いエキスパートの混合です。常にオンの重みをメモリ内に保持し、ストリームをできるだけ少なくし、量子化を意識したトレーニングで残りを圧縮します。これは、Apple も物理学から免除されているわけではないことを思い出させます。研究投稿の中でそれについて異例に率直に述べているだけだ。
境界: オンデバイス、クラウド、不透明な中間
Apple の基盤モデルは現在、オンデバイスからクラウドまで幅広い範囲に広がっており、クラウドエンドは驚くべき形状をしています。 AFM 3 の投稿より:
「私たちは Google および NVIDIA と協力して、プライベート クラウド コンピューティングを Google Cloud の NVIDIA GPU に拡張しました。」 - Apple 機械学習研究、AFM 3
Apple のセキュリティ チームは次のように述べています。
「Google および NVIDIA と協力して、Google Cloud 上で新しい Apple Intelligence ワークロードを実行します。」 - Apple セキュリティ、プライベート クラウド コンピューティングの拡大
Apple の最も要求の厳しいモデルは、Google で構築された Google Cloud の NVIDIA GPU 上で実行されます。独自のシリコンを設計し、オンデバイスのプライバシーを販売する企業にとって、競合他社のクラウド内で競合他社のハードウェア上に存在する主力クラウド モデルは、今週の最も驚くべき情報です。
一番確認したかったのはスイッチです。リクエストがデバイス上で実行されるのと、プライベート クラウド コンピューティングに送信されるのはいつですか。どちらが発生したかを事後的に知ることができますか?それで探しに行きました。 Apple の API は、プライベート クラウド コンピューティング モデル オプション、専用の PrivateCloudComp など、明示的な選択肢を公開します。

意図的に採用するuteLanguageModelタイプ。 Core AI ドキュメント、Foundation Models ドキュメント、または Expanding-Private-Cloud-Compute セキュリティに関する投稿で私が見つけることができなかったのは、デバイス上のリクエストがいつ透過的にオフロードされるか、またはそのルーティングが開発者またはユーザーに表示されるかどうかに関する記述です。つまり、この発見の正直なバージョンは、スペクトルは本物であり、クラウドは Google と NVIDIA であり、トリガー メカニズムとその監査可能性は単に公的に指定されていないということです。沈黙をどうするかはあなた次第です。
開発者が確認できるもの: タイミング
Core AI には、スタンドアロンのデバッガー アプリ、Xcode デバッグ ゲージ、およびインストゥルメント テンプレートの 3 つのツールが同梱されています。これらのツールは実際に何かを測定しているため、何を測定するのかを正確に把握する価値があります。コア AI 機器、ドキュメントによると:
「CPU、GPU、ニューラル エンジン全体の実行タイミングをプロファイルします...どの計算ユニットがモデルを実行するかなど...トレースは、Core AI イベントとハードウェア アクティビティを関連付けます。」 - Apple、コア AI ドキュメント
レイテンシ、トークン数、およびモデルを実行したコンピューティング ユニット - Xcode 内で、独自のアプリの Core AI 呼び出しについて。エネルギー、メモリ帯域幅、熱状態は、Core AI プロファイリングのドキュメントのどこにも登場しません。これはツールが報告する内容についての記述であり、判断ではありませんが、デバイス上のパフォーマンスのどれだけがまさにこれら 3 つによって決定されるかを考えると、この差は注目に値します。
並行して、Apple はパワー ユーザー向けの自力ウェイト パスとして MLX への投資を続けました。 WWDC 2026 では、複数の Mac にわたる分散推論 (Thunderbolt 5 上の新しい JACCL バックエンド)、OpenAI 互換の mlx_lm.server 、およびそれを中心に構築された Agentic-on-Mac ストーリーが追加されました。明らかに、MLX セッションにはコア AI または基盤モデルへの線引きはなく、システムとの意図的な 2 つのトラックの姿勢がとられています。

Core AI 上の独自のモデルと、MLX 上のオープン コミュニティのモデルである Foundation Models。
(そのようなフラグが付けられた 1 つの色の詳細: 新しい fm コマンドライン ツールの開発者のスクリーンショット - プラットフォーム一般教書に示されていると彼らは言う - OpenAI 互換の fm サーブと、同じオンデバイス対プライベート クラウド コンピューティング モデルの分割を公開しています。これはドキュメントではなく、1 人の人物のキャプチャとして扱ってください。)
後ろに引くと、ロードマップが読みやすくなります。
オンデバイス AI は現在、第一級のプラットフォーム機能です。 Apple Intelligence を強化するのと同じ推論エンジンが、独自の形式、ツールチェーン、プロファイラーを備えた開発者フレームワークです。それは機能よりも大きな取り組みです。
スタックは統合される前に断片化しています。 Core ML、Core AI、MLX は現在共存しており、開発者らは数時間以内にそう発表しました。Core AI の発表に関するスレッドは、3 つのうちのどれを使用するのか、そしてその理由を尋ねる人々でいっぱいです。 Apple はフレームワークを説明するストーリーよりも早くフレームワークを出荷しました。
難しい問題は普遍的な問題です。 AFM 3 の NAND 帯域幅許可とプレフィル対デコードのルーフラインは、すべてのローカル推論プロジェクトが戦う同じ制約です。興味深いのは、Apple が問題を解決したということではありません。それは、Apple が他の人々と同じ用語でそれらを説明しているということです。
雲の境界はtの部分です

[切り捨てられた]

## Original Extract

No new chips at WWDC 2026 - but a quiet rebuild of how AI runs on Apple silicon: Core AI, a new model format, the third-gen Foundation Models, and a flagship cloud model running on NVIDIA GPUs in Google Cloud. The roadmap, the major changes, and the odd findings.

Apple rebuilt its on-device AI stack at WWDC 2026 - Ziraph blog
Beta version (soft) launch soon! Subscribe to the announcement list to hear the moment the website and beta open ↗
Apple rebuilt its on-device AI stack at WWDC 2026
No new chips. But a quiet rebuild of how machine learning runs on Apple silicon - a new inference framework, a new model format, a new generation of on-device models, and a flagship cloud model running on someone else's GPUs. The interesting parts are the ones that weren't on a keynote slide.
· June 9, 2026 · 10 min read
WWDC 2026 brought no new silicon. What it brought instead was a structural rebuild of how AI runs on Apple silicon:
a new generation of on-device models,
and a noticeably different posture toward the cloud.
None of it was the headline - the headline was the consumer features. But the developer documentation, the session code, and one machine-learning-research post add up to a clearer roadmap than the keynote did, plus a few details that are genuinely odd.
I read this layer closely - I'm building a profiler for it - so here is what stood out: the major changes, the subtle tells, and the findings I had to double-check before believing. One ground rule up front: everything below is from Apple's own documentation, WWDC session pages, and research posts, quoted where it matters. Where something is an individual developer's claim or a forum reading rather than Apple's word, I say so. Where Apple simply does not say, I say that too.
And the biggest caveat of all: I'm in Europe, so I spent the night watching, reading, and researching - I'm sure I got something wrong due to lack of sleep. :-)
The big change: Core AI replaces Core ML for neural networks
For a decade, Core ML was the answer to “run a model on an iPhone.” At WWDC 2026 Apple introduced Core AI , and the framing is a handover, not an addition. Core AI's documentation sends the old cases back to Core ML:
“If your app uses model types other than neural networks, such as decision trees or tabular feature engineering, see Core ML.” - Apple, Core AI documentation
And Core ML's documentation sends the new ones forward:
“If your app integrates AI models using the latest architectures and inference techniques, see Core AI.” - Apple, Core ML documentation
Read together, that is a split: Core ML narrows to classic, non-neural machine learning - decision trees, tabular features - while neural networks and transformers move to Core AI, which Apple describes as the engine behind the product itself:
“Core AI allows your app to use the latest model architectures and inference techniques across the CPU, GPU, and Neural Engine.” - Apple, Core AI documentation
The subtle tell is in the tooling. Apple's new Core AI debug gauge carries a one-line restriction:
“The gauge does not support the Core ML framework.” - Apple, Core AI debug gauge documentation
The new instrumentation simply does not look at the old framework. Core ML is not deprecated - its APIs are intact, and there is real backward-compatibility value in that - but the center of gravity, and the tooling investment, has moved.
A new artifact: the .aimodel bundle
Core AI ships a new on-disk format, .aimodel , and the first odd thing about it is that it is not a file. It is a directory. Apple's open coreai-models repository treats it as one throughout - the Python exporter deletes an old one with a directory-only call, and the Swift runtime resolves it as a “ .aimodel directory.” Inside the surrounding model bundle is a plain-JSON metadata.json (schema version 0.2 ) that records the model kind (LLM, VLM, diffusion, segmenter), the tokenizer, vocabulary size, context length, the compression preset, and which file is the model. That JSON is documented and parseable. The weight payload itself - the part that would tell you the exact per-tensor bit-widths - is written by an opaque framework call, and its byte layout is not published anywhere I could find. So the format is half-open: a readable manifest wrapped around an undocumented blob.
Models are prepared with new Python tooling - Core AI Optimization ( coreai-opt , the successor to coremltools) for compression, and Core AI PyTorch Extensions ( coreai-torch ) to export straight from PyTorch into the format - then optionally compiled ahead of time with xcrun coreai-build compile into per-architecture .aimodelc assets. The compression menu is wider than the GGUF world's: integer weights at 2, 4, and 8 bits; float micro-formats including FP8 (E4M3) and FP4 (E2M1); block-scaled MXFP8; and palettization from 1 to 8 bits. One forum reader ( HN , opinion) noted Apple is also pushing activation quantization like w4a8 / w4a16 ; given Apple's install base, the formats it blesses could end up shaping how sub-100B models ship to everyone.
The hardware tell: the matmul moved into the GPU
No new chip, but WWDC 2026 made the M5 and A19 GPU story explicit, and it is the clearest hardware signal of the week. From Apple's M5/A19 tech talk:
“Neural accelerators are dedicated hardware in M5 purpose built for matrix multiplication. They're built into each shader core right alongside the other GPU pipelines such as ALU, raytracing... Each shader core has its own neural accelerator.” - Apple, “Accelerate your machine learning workloads with the M5 and A19 GPUs”
Apple's numbers: matrix multiplications up to 4 to 8 times faster, LLM time-to-first-token up to four times faster (prefill, which is compute-bound), token generation up to 25% faster (decode, which is memory-bound). And the framing underneath is one local-inference people will recognize, because it is the roofline - now stated in Apple's own Metal Performance Primitives guide:
“GEMMs with low arithmetic intensity are memory bound workloads, and GEMMs with high arithmetic intensity are compute bound workloads, forming the basis of a roofline model for kernel performance.” - Apple, Metal Performance Primitives Programming Guide
Compute-bound prefill, memory-bound decode: that prefill-versus-decode split is now Apple's own language, not just a community heuristic.
A second tell is hiding in code rather than slides. The coreai-models source infers a model's preferred compute unit from its graph structure: chunked, static-shape graphs prefer the Neural Engine; dynamic-shape graphs prefer the GPU. That quietly formalizes the bifurcation Apple has been hinting at for years - the Neural Engine for static, classic-shaped work, and the GPU (now with a Neural Accelerator inside each shader core) for the transformer matmul. Worth stressing: that is the model's preferred target encoded at export time, not a guarantee of where any given run actually executes.
The model: AFM 3, and the bandwidth wall
Apple also introduced its third generation of Foundation Models . On device: a 3-billion-parameter dense model (“AFM 3 Core”) and a 20-billion-parameter sparse mixture-of-experts (“AFM 3 Core Advanced,” natively multimodal, activating just 1 to 4 billion parameters at a time, and gated to the most capable Apple silicon).
The interesting part is the memory section, where Apple states the constraint plainly:
“the full model is stored in flash memory (NAND)” ... “NAND-to-DRAM bandwidth is too slow to swap weights token by token.” - Apple Machine Learning Research, AFM 3
That is Apple describing the exact wall every local-LLM runner runs into: a model too big to keep resident pays for it in bytes moved per token. Their answer is a mixture-of-experts with a high share of always-active “shared experts” alongside input-dependent “routed experts” - keep the always-on weights in memory, stream as little as possible - with quantization-aware training compressing the rest. It is a reminder that Apple is not exempt from the physics; it is just unusually candid about it in a research post.
The boundary: on-device, cloud, and the opaque middle
Apple's foundation models now span a spectrum, from on-device to the cloud, and the cloud end has a surprising shape. From the AFM 3 post:
“we worked with Google and NVIDIA to extend Private Cloud Compute to NVIDIA GPUs in Google Cloud.” - Apple Machine Learning Research, AFM 3
And from Apple's security team:
“collaborating with Google and NVIDIA to run new Apple Intelligence workloads on Google Cloud.” - Apple Security, Expanding Private Cloud Compute
Apple's most demanding model runs on NVIDIA GPUs, in Google Cloud, built with Google. For a company that designs its own silicon and markets on-device privacy, the flagship cloud model living on a competitor's hardware in a competitor's cloud is the most surprising tell of the week.
The part I most wanted to confirm is the switch. When does a request run on the device versus go to Private Cloud Compute, and can you tell after the fact which happened? So I went looking. Apple's APIs expose explicit choices - a Private Cloud Compute model option, a dedicated PrivateCloudComputeLanguageModel type you deliberately adopt. What I could not find - in the Core AI docs, the Foundation Models docs, or the Expanding-Private-Cloud-Compute security post - is any statement of when an on-device request transparently offloads, or whether that routing is visible to the developer or the user. So the honest version of this finding: the spectrum is real, the cloud is Google plus NVIDIA, and the triggering mechanism and its auditability are simply not publicly specified. Make of the silence what you will.
What developers can see: timing
Core AI ships three tools - a standalone Debugger app, an Xcode debug gauge, and an Instruments template - and it is worth being precise about what they measure, because they do measure something real. The Core AI instrument, per the docs:
“profiles execution timing across the CPU, GPU, and Neural Engine ... such as which compute units run your model ... The trace correlates Core AI events with hardware activity.” - Apple, Core AI documentation
Latency, token counts, and which compute unit ran the model - inside Xcode, for your own app's Core AI calls. Energy, memory bandwidth, and thermal state do not appear anywhere in the Core AI profiling documentation. That is a statement about what the tooling reports, not a judgment - but it is a notable gap given how much of on-device performance is decided by exactly those three.
Running in parallel, Apple kept investing in MLX as the bring-your-own-weights path for power users. WWDC 2026 added distributed inference across multiple Macs (a new JACCL backend over Thunderbolt 5 ), an OpenAI-compatible mlx_lm.server , and an agentic-on-Mac story built around it. Tellingly, the MLX sessions draw no line back to Core AI or Foundation Models - a deliberate two-track posture: the system's own models on Core AI and Foundation Models, the open community's models on MLX.
(One color detail, flagged as such: a developer's screenshot of a new fm command-line tool - shown, they say, in the Platforms State of the Union - exposes an OpenAI-compatible fm serve and the same on-device-versus-Private-Cloud-Compute model split. Treat it as one person's capture, not documentation.)
Pull back, and the roadmap is legible.
On-device AI is now a first-class platform capability. The same inference engine that powers Apple Intelligence is a developer framework, with its own format, toolchain, and profiler. That is a bigger commitment than a feature.
The stack is fragmenting before it consolidates. Core ML, Core AI, and MLX now coexist, and developers said so within hours - the thread under the Core AI announcement is full of people asking which of the three to use and why. Apple shipped the frameworks faster than the story that explains them.
The hard problems are the universal ones. AFM 3's NAND-bandwidth admission and the prefill-versus-decode roofline are the same constraints every local-inference project fights. The interesting thing is not that Apple solved them; it is that Apple now describes them in the same terms the rest of us do.
The cloud boundary is the part t

[truncated]
