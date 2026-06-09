---
source: "https://squeezlabs.github.io/handcrank/"
hn_url: "https://news.ycombinator.com/item?id=48463753"
title: "CrankGPT is an offline and off-the-grid AI box"
article_title: "CrankGPT — fully offline, human-powered local AI | CrankGPT"
author: "bobbiechen"
captured_at: "2026-06-09T18:52:30Z"
capture_tool: "hn-digest"
hn_id: 48463753
score: 1
comments: 0
posted_at: "2026-06-09T16:54:53Z"
tags:
  - hacker-news
  - translated
---

# CrankGPT is an offline and off-the-grid AI box

- HN: [48463753](https://news.ycombinator.com/item?id=48463753)
- Source: [squeezlabs.github.io](https://squeezlabs.github.io/handcrank/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T16:54:53Z

## Translation

タイトル: CrankGPT はオフラインのオフグリッド AI ボックスです
記事のタイトル: CrankGPT — 完全オフラインの人力ローカル AI |クランクGPT
説明: 手回し式のシングルボード コンピューター上で実行されるローカルの音声アシスタント。

記事本文:
CrankGPT — 完全オフラインの人力ローカル AI |クランクGPT
メインコンテンツにスキップ
リンク
メニュー
展開する
(外部リンク)
コピー
コピーされました
クランクGPT
動機
ハードウェア
シングルボードコンピュータ
すべてをまとめると
起動時間
CrankGPT は、完全にオフラインのオフグリッド AI ボックスです。
私たちの現在のデモは音声アシスタントのバリエーションであり、クランクを回して何かを言い、応答を受け取りますが、同じセットアップを使用して画像 (小さい) を生成し、詩を作成し (悪い)、コードを作成しました。バッテリーもクラウドもありません。手回しクランク、小さなコンピューター、そしてローカルで実行される音声モデルと言語モデルの小さなスタックだけです。電子機器が乾燥した状態で適切な温度に保たれていれば、これが 1000 年後も機能しないという理由はありません。
ハードウェア プロジェクトに取り組んだことのある人にはよく知られていると思いますが、概念実証の構築には約 1 週間かかり、カーネルの最適化、ボードの改訂、コードのリファクタリング、CAD の調整に何ヶ月もかかり、思い描いたとおりに動作するものに到達しました。この記事では、ハードウェア、ローカル音声エージェント スタック、およびこれほど小さなデバイスで会話をリアルに感じるために必要なエンジニアリングなど、私たちがどのように構築したかを説明します。
何かが「スマート」であるためには、現在、コンセントとデータセンターが想定されています。 CrankGPT は小さな引数であり、どちらも真である必要はありません。
ローカルモデルはプライベートモデルです。なぜ必要のないものを手放すのでしょうか？
私たちの周りの人々が、小さなモデルでも巨大なモデルと同じように達成できるタスクに、数分の1のコストとエネルギーでキロワットと何千ものトークンを投げているのを見るのは、ヨーロッパの小型実用車の感覚を不快にさせました。
誰もが物事を大きくするのに忙しい。私たちは、物事を小さくする機会がたくさんあると考えました。
8GB RAM と冷却ファン HAT を搭載した純正の Raspberry Pi 5 を使用しました。

同じ価格でよりパフォーマンスの高い SBD はありますが (以下で説明するように、より高速な DDR5 RAM を備えた Orange Pi は LLM 推論にさらに適しています)、Pi のアクセシビリティとソフトウェア エコシステムに勝るものはありません。 Pi は、音声認識、言語モデル、およびテキスト読み上げを CPU (アクセラレータなし) 上でローカルに実行します。
KEYESTUDIO ReSpeaker 2-Mic Pi HAT を使用しました。これは、音声アシスタント専用に設計された Pi 用のオールインワン オーディオ I/O ソリューションです。これには、ステレオ MEMS マイク アレ​​イとさまざまなオーディオ出力が含まれています (WM8960 コーデックを備えた古いバージョンを使用しました)。 Pi の GPIO ヘッダーに直接設置されており、エンクロージャ内であっても、適切な遠距離マイク性能を備えています。
私たちは、緊急 USB 充電用に市販されている、安価な既製の切り替え可能な電圧 20W 手回し発電機を選択しました。 Pi は通常約 1.5A を消費しますが、ハードに動作しているとき (CPU で推論を実行しているときなど)、電流要件が大幅に増加する可能性があり、その結果、発電機の電圧が Pi に必要な 4.8V を下回ったり、瞬間的な 5A スパイクの場合には、発電機の内部過電流保護が作動して電圧出力が完全に遮断され、Pi が電圧低下を引き起こすことがあります。
完全な推論スタックが起動したときに Pi が安定した電圧を確実に受け取るようにするため (そしてクランカーに少し休息を与えるため)、ジェネレーターの出力を平滑化し、短期間 (約 20 秒) の電力貯蔵庫として機能するカスタム コンデンサー ボードを構築しました。
クランクを通してその負荷曲線を感じることができます。LLM 推論と音声合成が同時に実行されると、クランクを回すのがはるかに難しくなります。
クランキング中は一秒一秒が大切で、Raspian が起動するまでの 1 分ほどが永遠のように感じられます。 DietPi は、高速ブート TI を優先する、最小限の機能をそぎ落とした Debian ベースのイメージです。

すぐに利用できる多くのデフォルト サービスよりも優れています。これにより起動時間が大幅に短縮され、不要な無線サービス (Bluetooth、Wi-Fi など) をオフにすることでさらに短縮され、Linux の起動から使用可能なユーザースペースまで約 3 秒で完了しました。
私たちは、RPI クラスのボード用に最適化された独自のエッジ音声エージェントを作成しました。これを既存のフレームワーク (Pipecat など) の上に構築するのではなく、最初から構築する動機: システムをエンドツーエンドで理解し、依存関係をできるだけ少なくしたかったのです。パイプラインは明らかなもので、CPU のレイテンシが最小限になるように各ステージが調整されています。
自動音声認識 (ASR) + 音声アクティビティ検出 (VAD)
Moonshine ASR は、CPU ベースの ASR としては断然最速のオプションであることが判明しました。 Whisper ベース サイズのモデルや NVIDIA の FastConformer と比較すると、騒がしい環境 (騒音の多いクランクのシナリオに関連) やアクセントのある音声では堅牢性がわずかに劣ります。ただし、リアルタイム音声エージェントという目標を考慮して、低遅延になるように最適化しました。エンドポイントには Silero VAD を使用します。
LLM は llama.cpp 上で実行されます。私たちが推奨するモデルは、小型の Liquid AI LFM2 バリアント (例: 350M または 1.2B) と、1B 形式の Gemma 3 です。
llama.cpp (pp512 と tg128 を備えた llama-bench、それぞれ 4 スレッド) を使用して測定された Raspberry Pi 5 のパフォーマンス:
トークンの生成は自己回帰デコードにおける最大のボトルネックであり、(生のコンピューティングではなく) メモリ帯域幅によって最も制約されます。これは、Raspberry Pi 5 (DDR4 RAM) と Orange Pi 5 Pro (DDR5 RAM) のプリフィルと生成レートを比較するとはっきりとわかります。
Orange Pi 5 Pro の生成レートは 29 ～ 58% 高くなります。これは主に DDR5 のメモリ帯域幅が大幅に高いためです。
ほとんどの大規模な LLM は、エッジ最適化されているものであっても、どちらのプラットフォームでも速度が遅すぎて、リアルタイムで役に立ちません。

音声エージェント。 1 桁のトークン生成レート (例: 7.8 トークン/秒の Qwen 3.5 2B) では、リアルタイムに近いと感じる会話には遅延が長すぎる応答時間につながります。
自然なサウンドで CPU で実行可能な音声モデルのリストは増えていますが、そのほとんどは単純に Raspberry Pi 上でリアルタイムで実行できません。 Kokoro 、 KittenML 、 PocketTTS 、および Piper は、低リソースのエッジ推論の候補となる可能性があります。 Piper はレイテンシーと生成速度で大差で勝利します。具体的には、Raspberry Pi 5 では、Piper は 20 単語のテスト発話を約 0.5 秒で合成しますが、Kokoro は 9 倍近く遅いです。 PocketTTS はストリーミングをサポートしているため、最初のバイトまでの時間が大幅に短縮されますが、Raspberry Pi ではリアルタイム係数 (RTF) が依然として 1.0 を超えており、音の途切れが発生します。
Piper のヘッドルームにより、実際の会話におけるストリーミング LLM 出力に追いつくことができます。他の人はそれができません。
LLM の出力を文ごとに Piper にストリーミングします。再生中の一時停止を避けるために、文の最大長に制限を設け、最初の文についてはより積極的に制限します。これにより、モデルに全体的な短い回答を事前に約束させることなく、できるだけ早く音声が開始されます。ユーザーは最初の単語をすぐに聞き、再生が追いつくまでモデルはバックグラウンドで生成を続けます。
すべてのコンポーネントは ONNX ランタイムで実行されます。 PyTorch の依存関係 (技術的には必要ではないものの、一部のコンポーネントに残っている) は、RAM を節約し、起動時間を短縮するために削除されました。
クランクを開始してから CrankGPT と会話するまでにかかる時間は約 30 秒です。起動時間には次のものが含まれます。
~10 ～ 15 秒 — 完全なファームウェア シーケンスによる Pi 5 コールド ブート
~3s — ユーザー空間への Linux ブート (DietPi)
~10 ～ 15 秒 — 音声エージェントの起動 (Python インポート、ロードイン)

gモデル重量)
BOOT_DELAY=0、スプラッシュの無効化、未使用のブート ソースの削除、利用可能な最速の SD カードなど、明らかな最適化をすべて行ったとしても、Pi 5 の Linux 前の段階では、依然として約 10 ～ 15 秒かかります。 Pi 4 とは異なり、Pi 5 はカーネルをロードする前に、より PC に近いファームウェア シーケンス (PMIC ランプ、RP1 初期化、EEPROM ブートローダーを介した PCIe/USB 列挙) を実行します。そのため、その下限をユーザーランドから突破するのは困難です。そして残念ながら、Pi 5 にはスリープ モード/DRAM 保存機能がないため、電圧が最小要件を下回るたびにゼロから開始する必要があります。
音声エージェントの起動時に最も遅い部分は、最初の実行時の Python インポートです。私たちは明らかな修正を試みましたが、どれも意味のある解決策はありませんでした。バイトコードのプリコンパイル (compileall) は何も行わず、Python はすでに .pyc ファイルを自動的にキャッシュしています。遅延インポートでは、せいぜい数百ミリ秒短縮されます。コールド スタート時間の大部分はコードに含まれていません。大規模な共有ライブラリ (特に ONNX ランタイム) の dlopen 処理と、Python がインポート グラフを実行する際の SD カードからの数百回の小さなランダム読み取りに時間がかかっています。最初の呼び出しではページ キャッシュは定義上コールドであるため、起動後にページ キャッシュをウォームアップしてもほんのわずかしか役に立ちませんでした。
NVMe は最も驚くべき行き詰まりでした。ランダム読み取りの高速化は明らかな勝利だったはずですが、Pi 5 では、EEPROM ブートローダーが起動する前に PCIe を列挙して NVMe コントローラーのファームウェアをロードする必要があり、Linux 前の段階に約 10 秒追加されます。実行時に得たものは、起動時に失われ、その後、一部が失われます。私たちのユースケース (セッションごとにコールドスタート) では、SD カードの方がエンドツーエンドで高速になりました。
起動時間をさらに短縮するには、Python レイヤーを削除し、エージェント グルーを C (または Rust) バージョンに置き換えると、さらに 5 秒ほどスターを節約できる可能性があります。

タップ時間。これは読者の演習として残しておきます。
私たちの目標は、ローカル音声エージェントのデモンストレーションでよく見られる数秒の遅延を発生させずに、スムーズなリアルタイム会話を可能にする、完全にオフラインのオフグリッド音声エージェントを構築することでした。上で説明したように、LLM の選択 (およびそれぞれのトークン生成レート) は、ユーザーが認識する最初のバイトまでの時間 (TTFB) の主な要因です。以下は、典型的な会話からの測定値であり、すべてのターンの TTFB を平均しています。
CrankGPT の消費電力は実際に実行される AI 推論の量に依存します。 Pi は電圧にこだわりがあり、電圧は 4.8V から 5.3V の間である必要があるため、興味深い変数は電流です。最大負荷時に最大 5A の短時間の電流スパイクが観察されました。一般的なシナリオをいくつか示します。
現時点では、最も高度な AI ワークロードを Raspberry Pi 上で実行することは現実的ではありませんが、私たちの研究は、大量の電力を消費せずにエッジ上でローカルに実行できる未開発の AI アプリケーションが存在することを示唆しています。そして、モデルが小型化され、より効率的になるにつれて（自己回帰デコーディングから脱却する可能性があり）、「エッジ」は高価な最新モデルの iPhone から、はるかに小型で安価なデバイスに移行することになります。

## Original Extract

A local, voice assistant running on a hand-crank-powered single board computer.

CrankGPT — fully offline, human-powered local AI | CrankGPT
Skip to main content
Link
Menu
Expand
(external link)
Copy
Copied
CrankGPT
Motivations
Hardware
Single Board Computer
Putting it all together
Startup Time
CrankGPT is a fully offline and off-the-grid AI box.
Our current demos are variations on voice assistants—turn the crank, say something, get a response—but we’ve generated images (small), made poetry (bad), and written code using the same setup. There’s no battery or cloud. Just a hand crank, a little computer, and a small stack of speech and language models running locally. Provided the electronics are kept dry and at a reasonable temperature, there’s no reason this thing won’t still work in a thousand years.
As will be familiar to anyone who has ever undertaken a hardware project, it took about a week to build a proof of concept and many months of kernel optimizations, board revisions, code refactors, and CAD tweaks to get to a thing that works as we envisioned. This article walks through how we built it: the hardware, the local voice agent stack, and the engineering required to make a conversation feel real on a device this small .
For something to have “smarts” currently assumes a wall socket and a data center. CrankGPT is a small argument that neither has to be true.
Local models are private models. Why give away what we don’t have to?
It offended our European small-practical-car sensibilities to see people around us throwing kilowatts and thousands of tokens at tasks small models could accomplish just as well as huge ones, for a fraction of the cost and energy.
Everyone is busy making things bigger. We figured opportunities abound to make things smaller.
We used a stock Raspberry Pi 5 with 8GB RAM and a cooling fan HAT. There are better performing SBDs for the same price (an Orange Pi with its faster DDR5 RAM is an even better fit for LLM inference as we’ll discuss below), but it’s hard to beat the Pi’s accessibility and software ecosystem. The Pi runs speech recognition, a language model, and text-to-speech locally on CPU (no accelerators).
We used the KEYESTUDIO ReSpeaker 2-Mic Pi HAT : an all-in-one audio I/O solution for Pi designed specifically for voice assistants. It includes a stereo MEMS mic array and various audio outputs (we used the older version with the WM8960 codec). It sits directly on the Pi’s GPIO headers and has decent far-field mic performance, even within an enclosure.
We chose a cheap off-the-shelf switchable voltage 20W hand-crank generator marketed for emergency USB charging. The Pi normally draws around 1.5A, but when it’s working hard (as it does when doing inference on the CPU), its current requirements can increase substantially, causing the generator voltage to sag below the Pi’s required 4.8V or even, in the case of a momentary 5A spike, to trigger the generator’s internal overcurrent protection and shut off the voltage output entirely, causing the Pi to brown out.
To ensure the Pi sees a steady voltage when the full inference stack kicks in (and to afford crankers a little rest), we built a custom capacitor board to smooth out the generator’s output and act as a short-term (~20 second) power reservoir.
You can feel that load curve through the crank: when LLM inference and speech synthesis run together, the crank gets a lot harder to turn.
When you’re cranking, every second counts—the minute or so it takes Raspian to boot up feels like an eternity. DietPi is a minimalistic, stripped-down Debian-based image that prioritizes fast boot time over lots of immediately available default services. It shrank our startup time substantially, and turning off unneeded radio services (Bluetooth, Wi-Fi, etc.) reduced it even further: from Linux boot to a usable userspace in around 3 seconds.
We wrote our own edge voice agent optimized for RPI-class boards. Our motivation for building this from scratch rather than on top of existing frameworks (like e.g. Pipecat): we wanted to understand the system end to end and have as few dependencies as possible. The pipeline is the obvious one, with every stage tuned for minimal latency on CPU:
Automatic Speech Recogntion (ASR) + Voice Activity Detection (VAD)
Moonshine ASR turned out to be by far the fastest option for CPU-based ASR. It’s slightly less robust in noisy environments (relevant in our scenario with a noisy crank) or on accented speech compared to Whisper-base-sized models or NVIDIA’s FastConformer. But we optimized for low latency given our goal of a real-time voice agent. For endpointing, we use Silero VAD .
The LLM runs on llama.cpp . Our preferred models are small Liquid AI LFM2 variants (e.g. 350M or 1.2B), along with Gemma 3 in its 1B form.
Raspberry Pi 5 performance measured using llama.cpp ( llama-bench with pp512 and tg128, 4 threads each):
Token generation is the biggest bottleneck in autoregressive decoding and is most constrained by memory bandwidth (not raw compute). This is clearly visible when comparing the prefill and generation rates on a Raspberry Pi 5 (DDR4 RAM) versus an Orange Pi 5 Pro (DDR5 RAM):
Generation rates on the Orange Pi 5 Pro are 29-58% higher, mainly due to the significantly higher memory bandwidth of DDR5.
Most larger LLMs—even those marketed as edge-optimized —are way too slow on either platform to be useful in a real-time voice agent. Single-digit token generation rates (e.g. Qwen 3.5 2B at 7.8 tok/sec) lead to response times with latency too high for conversation that feels anywhere close to real time.
There’s a growing list of natural-sounding, CPU-runnable voice models, but most simply don’t run in real time on a Raspberry Pi. Kokoro , KittenML , PocketTTS and Piper are the likely contenders for low-resource edge inference. Piper wins by a large margin on latency and generation speed. Concretely , on a Raspberry Pi 5, Piper synthesizes our 20-word test utterance in about half a second, while Kokoro is nearly 9× slower. PocketTTS does support streaming, which significantly reduces time-to-first-byte, but its real-time factor (RTF) is still above 1.0 on a Raspberry Pi causing audible stuttering .
Piper’s headroom is what lets it keep up with streaming LLM output in a real conversation. The others just can’t.
We stream the LLM’s output sentence-by-sentence into Piper. To avoid pauses during playback, we cap the maximum sentence length, and we cap it more aggressively for the first sentence. That gets speech started as fast as possible without forcing the model to pre-commit to a short answer overall. The user hears the first words quickly, and the model keeps generating in the background while playback catches up.
All components run on ONNX Runtime. PyTorch dependencies (lingering in some components while not technically required) were removed to save RAM and improve startup time.
It takes about 30 seconds from the moment you start cranking to the moment you’re having a conversation with CrankGPT. Startup time includes:
~10–15s — Pi 5 cold boot through full firmware sequence
~3s — Linux boot to userspace (DietPi)
~10-15s — Voice Agent startup (python imports, loading model weights)
Even with all the obvious optimizations—BOOT_DELAY=0, splash disabled, unused boot sources removed, fastest available SD card—the Pi 5’s pre-Linux stage still costs us ~10–15 seconds. Unlike the Pi 4, the Pi 5 runs a much more PC-like firmware sequence (PMIC ramp, RP1 init, PCIe/USB enumeration via the EEPROM bootloader) before it ever loads a kernel, and that floor is hard to break through from userland. And unfortunately, Pi 5 doesn’t have a sleep mode/DRAM preservation, so every time the voltage drops below its minimum requirements, you have to start from zero.
During voice agent startup, the slowest part are Python imports on first run. We tried the obvious fixes and none of them helped meaningfully. Precompiling bytecode ( compileall ) was a no-op—Python already caches .pyc files automatically. Lazy imports trimmed a few hundred milliseconds at best; the bulk of cold-start time isn’t in our code, it’s in dlopen -ing large shared libraries (ONNX Runtime in particular) and in hundreds of small random reads off the SD card as Python walks the import graph. Warming the page cache after boot helped only marginally because for the first invocation the page cache is cold by definition.
NVMe was the most surprising dead end. Faster random reads should have been a clear win, but on the Pi 5 the EEPROM bootloader has to enumerate PCIe and load the NVMe controller’s firmware before it can boot, adding roughly 10 seconds to the pre-Linux stage. What we gained at runtime, we lost at boot, and then some. For our use case—cold start every session—an SD card ended up being faster end to end.
To reduce startup time further, dropping the Python layer and replacing the agent glue with a C (or Rust) version could probably save another ~5s of startup time. We leave this as an exercise to the reader.
Our goal was to build a fully offline and off-the-grid voice agent that allows for smooth, real-time conversations without the multi-second latencies often seen in demonstrations of local voice agents. As described above, LLM selection (and its respective token generation rate) is the primary driver of the time-to-first-byte (TTFB) the user perceives. Below are some measurements from a typical conversation, averaging TTFB across all turns:
CrankGPT’s power draw really depends on the amount of AI inference running. The Pi is picky about its voltage, which needs to be between 4.8V and 5.3V, so the interesting variable is current. We’ve observed brief current spikes of up to 5A under maximum load. Here are a few common scenarios:
While it’s currently impractical to run most sophisticated AI workloads on a Raspberry Pi, our work suggests that there exist a whole class of unexplored AI applications that can run locally on the edge without consuming huge amounts of power. And as models get smaller and more efficient (potentially moving away from autoregressive decoding), the “edge” will migrate from your expensive latest-model iPhone to much smaller and cheaper devices.
