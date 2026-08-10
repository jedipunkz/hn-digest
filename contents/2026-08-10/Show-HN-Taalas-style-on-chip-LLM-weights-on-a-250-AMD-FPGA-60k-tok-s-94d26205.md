---
source: "https://www.mikeayles.com/blog/on-chip-llm-kv260/"
hn_url: "https://news.ycombinator.com/item?id=49242475"
title: "Show HN: Taalas-style on-chip LLM weights on a $250 AMD FPGA (60k tok/s)"
article_title: "Taalas-Style On-Chip Weights on a $250 FPGA: a Language Model at 60k tok/s | Michael Ayles"
author: "mikeayles"
captured_at: "2026-08-10T12:45:18Z"
capture_tool: "hn-digest"
hn_id: 49242475
score: 4
comments: 0
posted_at: "2026-08-10T11:52:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Taalas-style on-chip LLM weights on a $250 AMD FPGA (60k tok/s)

- HN: [49242475](https://news.ycombinator.com/item?id=49242475)
- Source: [www.mikeayles.com](https://www.mikeayles.com/blog/on-chip-llm-kv260/)
- Score: 4
- Comments: 0
- Posted: 2026-08-10T11:52:07Z

## Translation

タイトル: Show HN: 250 ドルの AMD FPGA 上の Taalas スタイルのオンチップ LLM 重み (60,000 トーク/秒)
記事のタイトル: 250 ドルの FPGA 上の Taalas スタイルのオンチップ重み: 60,000 tok/s の言語モデル |マイケル・アイルズ
説明: Xilinx Kria KV260 のオンチップ メモリ内で完全に動作する 3.16M パラメータの INT4 トランスです。トークン ループ内のゼロ DRAM、ファブリック上で 59,965 トークン/秒、ビット正確。ライブでチャットしましょう。

記事本文:
Michael Ayles ブログ RSS プロジェクト About Aug 10, 2026 · 250 ドルの FPGA 上の FPGA Taalas スタイルのオンチップ重み: 60k tok/s の言語モデル
Xilinx Kria KV260のオンチップ メモリ内で完全に動作する316万パラメータのINT4トランス。トークン ループ内のゼロ DRAM、ファブリック上で 59,965 トークン/秒、ビット正確。ライブでチャットしましょう。
私は Taalas による chatjimmy.ai のデモに非常に感銘を受け、FPGA の再構成可能なロジックのファブリック内に何を詰め込めるか見てみたいと思いました。 「PL」と同じです。モデル全体がここで実行される場所。 250 ドルの FPGA フィールド プログラマブル ゲート アレイ: 固定 CPU 上でソフトウェアを実行するのではなく、カスタム デジタル回路に配線する再構成可能なロジックが満載のチップです。 。 (4 GB が利用可能) DDR を使用しないことにより、CPU とファブリックの両方で共有されるオフチップ DRAM (ボードのメイン メモリ)、最大 20 GB/秒。この共有コントローラは「帯域幅の壁」です。そして、すべてを URAM UltraRAM に保存します。これは大きくて幅の広いオンチップ SRAM (約 18 Mb、64 ブロック) です。常駐の INT4 重みイメージを保持します。重要なのは、真のデュアルポートであるため、2 つのコホートの「スプリット ブレイン」が可能になります。 / BRAM ブロック RAM: 小型で柔軟なオンチップ SRAM ブロック (合計約 5 Mb)。アクティベーション、スクラッチ、および KV キャッシュをここに保持します。 , 私は100,000 tok/sを目標に設定しました。
見出し、ビット正確でシリコン上で測定: ファブリック上で 1 秒あたり 59,965 トークン。このボード独自の Arm コアの同じモデルは 11 tok/s を管理します。私のラップトップの RTX 3050 Ti は 719 を管理します。
Taalas は Llama 3.1 8B をチップに組み込むことに成功し、使用可能なチャットを行うことができました。私が扱うメモリは約 3 MB でした。つまり、適合できる最も強力なモデルは、TinyStories 約 210 万の非常に単純な童話の合成コーパスであり、小さなモデルでも一貫した英語を学習できるように設計されています。トレーニングデータはこちら。 -クラス 1: 3.16M

パラメータ、INT4 で約 1.5 MB。当初、私は各単語を辞書のルートに縮小するという見出し語化 (「saying」から「say」、「words」から「word」) にすることで、さらに縮小したいと考えていました。 Keviniser はコーパスを内容単語まで見出し化し、同じストーリーの文字数を約 30% 減らします。トレーニング コーパスを作成して、「トリックを実行できる人がほとんどいないのに、なぜ多くの言葉を使うのか」とケビン マローンのように聞こえますが、これによって実際にモデルが小さくなったわけではなく、最大 30% 高速化されただけだったので、私は自分自身に戻り、ハードウェアが両方のモデル間でフリップフロップできるようにしました。
このウィジェットは、通常の HTTP のリクエスト後にハングアップするのとは異なり、ブラウザとサーバー間の永続的な双方向接続である実際の WebSocket です。チャット ストリームのキャラクターをライブにするもの。ボードへの接続。あなたの言葉は、Cloudflare トンネルを通って、サービングボックス、Kria、ファブリックに入り、そして戻ってきます。ステータス ドットが緑色の場合は、ウェールズの単一の FPGA と通信していることを意味します。これはアシスタントではなく、ストーリージェネレーターです。それは質問を理解していません。 「昔々」と言うだけで終わります。
期待を裏切らないことを祈りますが、約 1.5 MB のウェイトでできることは限られています。
これはすべて、K26 SOM System-on-Module (KV260 ボードを中心に構築されているプラ​​グイン コンピューティング モジュール) を搭載した AMD KV260 開発ボード上で実行されています。 。これは、Zynq シリーズの最新 SoC の 1 つ (具体的には、Zynq UltraScale+ ZU5CG/ZU5EV、キャッチーですね?) で、OG 7000 の後継製品です。 10 年前、私は開発ボードがどれだけ欲しかったでしょうか。
私は別のサイド プロジェクトである確定的ビジョン パイプラインのために KV260 を購入しました。このボードは「ビジョン AI スターター キット」として販売されていますが、Vitis AMD/Xilinx のボード用アプリケーション開発スタックです。 KV260 の「ビジョン AI」デモは、KV260 上で実行されます。そのほとんどは、KV260 ではなく Arm コア上で実行されます。

ファブリック。物体検出は、KV260 (~1.33 GHz) 上のクアッドコア A53 CPU で実行されます。ファブリックが比較されるベースライン、および CPU がループ内にあるときのオーケストレーター。 A53 はハードウェア マットムルを持たない弱いコアです。 Vitis ライブラリは、ファブリックをプッシュするものではなく、基本的に Linux 上の OpenCV であることが判明しました。 CPU 上で確率的 AI を実行したい場合は、クアッドコア A53 は絶対に選択しないでしょう。
決定論的深度エンジンをファブリック上の Verilog で実行することに成功しましたが、これはロボット プロジェクトであるため、安全上の理由からハードウェアのダウンタイムが長くなります。そこで、空き時間にどこからでも作業できるように、予備の SD カードに新しい Linux イメージを置き、SSH をセットアップしました (すぐに制御不能になりました。別のリモート マシンが JTAG アクセスできるようになり、PSU は Zigbee リレー上にあるため、リモートで強制的に再起動できます)。
1 つの事実がすべてを動かします。一度に 1 つのトークンの生成はメモリに制限され、計算に制限されません。次のトークンを生成するには、モデル内のすべての重みを 1 回読み取ります。算数は安くて、読むのはコストがかかります。
KV260 では、A53 とファブリックは 1 つの DDR を共有します。オフチップ DRAM (ボードのメイン メモリ)、最大 20 GB/秒、CPU とファブリックの両方で共有されます。この共有コントローラは「帯域幅の壁」です。コントローラーの速度は約 20 GB/秒です。モデルが DDR にある場合、ファブリックと CPU は同じストローを使用し、ファブリックからは何も得られません。 DDR または CPU over AXI へのラウンドトリップ CPU をファブリックに接続するオンチップ バス プロトコル ファミリ。あなたを殺します。唯一の回避策は、帯域幅が数百 GB/秒のオンチップ メモリにすべてが収まるほど十分に小さいモデルです。これはターラス氏、グロク氏、セレブラス氏の洞察です。記憶の壁はエンタープライズです。

emy とオンチップの重みが逃げ道です。彼らはオンチップの予算を拡大するために数億ドルを費やしています。 KV260 では約 3 MB が得られます。
DDR 常駐: エスケープなし Arm A53 PL ファブリック DDR ~20 GB/s、共有 1 つのコントローラー、1 つのストロー ファブリックでは何も買わない: ~11 tok/s オンチップ: 唯一のエスケープ PL ファブリック URAM + BRAM 数百 GB/s ～ TB/s モデルは ~3 MB Firehose に収まる INT4、1.5 MB 常駐: 最大 59,965 tok/s がオンチップに十分に小さいことがすべての秘訣 シングルストリームデコードはメモリ帯域幅の制限を受けますが、計算の制限はありません。ウェイトが DDR に存在する場合、Arm コアとファブリックは 1 つの ~20 GB/s コントローラーを共有し、ファブリックは CPU よりも高速ではありません。勝利する唯一の方法は、帯域幅が数百 GB/秒から TB/秒のオンチップ SRAM にすべて収まるほどモデルを小さくすることです。このプロジェクトの他のすべては、その一文から導かれます。
3 MB には、重み、アクティベーション、および KV キャッシュを保持する必要があります。これでは、スマート モデルにとって十分な余地はありません。文をつなぎ合わせることができるモデルにとっては、かろうじて十分なスペースです。そこで、大手企業がほとんど引くことができない 2 番目のレバーは、問題がなくなるまでモデルを縮小することです。
モデルを十分に小さくする (ここで Kevin が登場します)
3.16M パラメータの INT4 4 ビット整数 (16 レベル)。重みは INT4 として保存されるため、モデルはオンチップに適合するほど十分に小さくなります。トランスは約 1.5 MB に収まります。これで予算は満たされましたが、依然としてすべてのバイトが重要であり、速度を高めるための最後の手段は「モデルをより愚かにする」ことであることが判明し続けました。 （あるいは私はそう思っていました）
そのため、トレーニング コーパス (TinyStories) は、英語から内容単語までを取り除くツールを通じて実行されます。ケビンに触発されて、「少数の単語で効果があるのに、なぜ多くの単語を言って時間を無駄にするのか」は、「少数の単語で効果があるのに、なぜ多くの単語を言って時間を無駄にするのか」になります。

『The Office』のマローンのコミュニケーション哲学、そしてその結果、モデルも彼のように話します。コーパス全体にわたって、圧縮には 3 億 7,170 万語が 2 億 6,050 万語 (測定値の約 70%) までかかります。
私は圧縮コーパス モデルが小さくなるだろうと予想していました。結果はまったく同じサイズでした。後から考えるとそれは明らかです。パラメータ数はコーパスではなくアーキテクチャによって固定されています。圧縮によって行われるのは、出力分布を縮小することです。同じストーリーを語るのに必要な文字数が最大 30% 少ないため、トークンあたりのレートが同じであっても、実質的な速度が向上します。
正直なバージョン: 見出し語化されたコーパスは約 1.5 倍購入します。オンチップ対 DDR の方が圧倒的に有利です。ケビンはチリの付け合わせであり、食事ではありません。
私の目標は100,000トーク/秒でした。私はそこまでは到達できませんでしたが、徹底的に頑張りました。最終的には、アテンション ウィンドウを T=1 に減らしてラウンド トリップや競合状態をできる限り排除することで、ユーザビリティを犠牲にすることさえできました。
レコードのビルドは、1 つのウェイト パスを共有する 16 個の並列ストリームであり、ファブリック内で完全にシーケンスされ、CPU は何も触れません。 200 MHz で 59,965.5 tok/s を測定し、16 ストリーム中 16 ストリームが整数基準に対してビット正確で、3 回の実行が 3 回実行されました。告白: その 16 個のストリームは何も覚えていません。それぞれが 1 つのトークンのアテンション ウィンドウでデコードされます。その上に構築されたチャットは、忠実な文字を発し、「彼、彼、彼」と階段から落ちます。速い、意味がない、それらは同じ性質を一歩踏み出しすぎたものです。
したがって、デプロイされたチャットは、別の正直なビルドになります。 1 つのストリーム、完全にトレーニングされたコンテキスト ウィンドウ、完全な再計算に対してビット正確な KV キャッシュ。これは、カウントされたサイクルで 19,242 tok/s のファブリックを実行し、ライブで測定すると約 21,300 tok/s であり、この数値は、負荷スイープでの同時接続数が 1 ～ 2,000 まで横ばいでした。

0 番目のエラー (ピークは 21,479 回観測されました)。それはあなたの最後の数回のターンを記憶しており、それは上のウィジェットにあるものです。どちらの数字も記録の見出しを借用することはできません。
11 tok/s から 60k までのパスははしごで、各段がシリコン上で測定されました。
そして、その上限は想定されたものではなく証明されたものです。 3 番目の MAC Multiply-ACCumulate のパッキングを偽るスクリプトがあります。つまり、1 乗算 1 加算演算の行列乗算が構築されます。ここでの計算の基本単位。 DSP48E2 へ DSP48E2: FPGA の専用ハードウェア乗算器ブロック (このチップ上では 1248)。それぞれに 2 つの INT4 x INT8 の積和演算をパックできます。 120 万を超えるランダム化された製品があり、N=32 にはチップにはない 2,048 個の DSP が必要です。このシリコン上のこのアーキテクチャの実際の制限は 62k ～ 78k です。 100,000 という数字は、始める前から適当に推測したもので、面白い目標値になるだろうとだけ思っていました。 60,000 を達成することは私にとって十分近い数字であり、使用可能な 20,000 を超える数字は、私が自分で設定したもう 1 つの目標である Taalas の数字を上回っています。
オンチップ トリックは、モデルがオンチップに適合する場合にのみ成功し、クロスオーバーは約 630 万のパラメーターになります。それを過ぎると、DDR に流れ出て、再び壁にぶつかります。長いコンテキストは、同じ方法で KV キャッシュを流出させます。これはおもちゃの模型のような構造上のテクニックであり、意図的に出力が悪く、良いものではありません。ユーモアのセンスのある測定器です。
スケールの場合: Taalas は、モデルごとにテープアウトされた ASIC 上のトランジスタに重みをエッチングします。 Cerebras は 44 GB の SRAM をウェーハ上に保持します。 Groq はチップあたり 230 MB を保持し、数百ものチップをギャングします。 9 桁の予算から素敵なディナーの価格まで、同じ一文で洞察が得られます。
以上がすべて物語です。エンジニアリングについてはサブ記事に記載されており、その中のすべての数値は引用される前にシリコン上でビット正確です。
チップ内部には、ワイドW

ord GEMV のトリック、デュアル ポート スプリット ブレイン、非線形ブリック、および 16 ストリームが上限であることの証明。
質問なしのサンプリング、1 つの未測定のホスト ループがすべての応答の 58% を消​​費する様子、およびトークンあたり 193 の読み取りを 1 つのシード書き込みに集約するガンベルマックス ID について説明します。
見知らぬ人への FPGA の提供、トンネル、サービス ボックス、そしてダッシュボードが意図的にボードから離れた場所にある理由。
戦争の話とその方法、iverilog の所在、シリコンがタイミング レポートを 1.3 ～ 1.76 倍上回る理由、そしてとにかく私が勝ってやめた最適化。
すべて手作業で LLM で書かれた Verilog であり、HLS 高位合成はありません。C/C++ をハードウェアにコンパイルするツールなので、Verilog を手で書く必要はありません。このプロジェクトには何もありません。すべてのブロックは私とクロード コードによって直接書かれた RTL であり、C からコンパイルされたことはありません。 Claude Code は確かな成果を上げました。結局のところ、これはサイド プロジェクト中のサイド プロジェクトですが、FPGA を限界まで追い込むことは、TypeScript で CRUD アプリを書くことほど快適ではありません。おそらく残高は次のようになります。
アーキテクト: 私 80%、クロード コード 20%
実装: 私 20%、クロード コード 80%
そして、HDL ハードウェア記述言語 (Verilog、VHDL)、つまり実行命令ではなく回路を記述するコードを書くのはそれほど快適ではないと言いながらも。このプロジェクトは

[切り捨てられた]

## Original Extract

A 3.16M-parameter INT4 transformer running entirely in the on-chip memory of a Xilinx Kria KV260. Zero DRAM in the token loop, 59,965 tok/s on the fabric, bit-exact. Chat with it live.

Michael Ayles Blog RSS Projects About Aug 10, 2026 · FPGA Taalas-Style On-Chip Weights on a $250 FPGA: a Language Model at 60k tok/s
A 3.16M-parameter INT4 transformer running entirely in the on-chip memory of a Xilinx Kria KV260. Zero DRAM in the token loop, 59,965 tok/s on the fabric, bit-exact. Chat with it live.
I was so impressed by the chatjimmy.ai demo by Taalas, I wanted to see what I could squeeze inside the fabric The FPGA's reconfigurable logic. Same thing as 'PL'. Where the whole model runs here. of a $250 FPGA Field-Programmable Gate Array: a chip full of reconfigurable logic you wire into a custom digital circuit, instead of running software on a fixed CPU. . By not using the (4 GB available) DDR The off-chip DRAM (the board's main memory), ~20 GB/s, shared by both the CPU and the fabric. This shared controller is 'the bandwidth wall'. , and keeping everything in URAM UltraRAM: the big, wide on-chip SRAM (~18 Mb, 64 blocks). Holds the resident INT4 weight image. Crucially, it is true dual-ported, which enables the two-cohort 'split-brain'. / BRAM Block RAM: small, flexible on-chip SRAM blocks (~5 Mb total). Holds activations, scratch, and the KV cache here. , I set myself a target of 100,000 tok/s.
The headline, bit-exact and measured on silicon: 59,965 tokens per second on the fabric. The same model on this board’s own Arm cores manages 11 tok/s. My laptop’s RTX 3050 Ti manages 719.
Taalas managed to bake Llama 3.1 8B into their chip, and have a usable chat. I had ~3 MB of memory to work with, which means the most powerful model we can fit is a TinyStories A synthetic corpus of about 2.1 million very simple children's stories, designed so that even tiny models can learn coherent English from it. The training data here. -class one: 3.16M parameters, ~1.5 MB at INT4. Initially, I wanted to shrink it further by lemmatising Reduce each word to its dictionary root ('saying' to 'say', 'words' to 'word'). The Keviniser lemmatises the corpus down to content words, about 30% fewer characters for the same story. the training corpus and making it sound like Kevin Malone, ‘why use many word when few do trick’, but since this didn’t actually make the model any smaller, it just speeds it up by ~30%, I went back on myself and just allowed the hardware to flip-flop between both models.
This widget is a real WebSocket A persistent two-way connection between browser and server, unlike ordinary HTTP's request-then-hang-up. What lets the chat stream characters live. connection to the board. Your words go through a Cloudflare tunnel, to a serving box, to the Kria, into the fabric, and back. If the status dot is green, you are talking to a single FPGA in Wales. It’s a story generator, not an assistant. It doesn’t understand questions. Give it “once upon a time” and it finishes it.
I hope you didn’t get your hopes up, there’s only so much you can do with ~1.5 MB of weights.
This is all running on my AMD KV260 dev board, featuring the K26 SOM System-on-Module: the plug-in compute module that the KV260 board is built around. . This is one of the latest SoCs in the Zynq line (specifically the Zynq UltraScale+ ZU5CG/ZU5EV, catchy eh?), the successor to the OG 7000s. God, how I wanted a dev board for one of those 10 years ago.
I bought the KV260 for a different side project, a deterministic vision pipeline. The board is sold as a “vision AI starter kit,” but the Vitis AMD/Xilinx's application development stack for their boards. The KV260's 'vision AI' demos run on it, mostly on the Arm cores rather than the fabric. object detection runs on the quad core A53 The quad-core Arm Cortex-A53 CPU on the KV260 (~1.33 GHz). The baseline the fabric is compared against, and the orchestrator when the CPU is in the loop. s, and the A53 is a weak core with no hardware matmul. The Vitis libraries turned out to be basically OpenCV on Linux rather than anything that pushes the fabric. If I wanted probabilistic AI running on a CPU, I sure as heck wouldn’t choose a quad core A53.
I did manage to get my deterministic depth engine running in Verilog on the fabric, but since it’s a robotics project, the hardware has a lot of downtime for safety reasons. So I put a fresh Linux image on a spare SD card and set up SSH so I can work on this from anywhere in what little free time I have (it very quickly spiralled out of control: another remote machine now has JTAG access to it, and the PSU is on a Zigbee relay so I can force reboots remotely).
One fact drives everything. Generating one token at a time is memory bound, not compute bound . To produce the next token you read every weight in the model once. The arithmetic is cheap, the reading is the cost.
On the KV260, the A53s and the fabric share one DDR The off-chip DRAM (the board's main memory), ~20 GB/s, shared by both the CPU and the fabric. This shared controller is 'the bandwidth wall'. controller at roughly 20 GB/s. If the model lives in DDR, the fabric and the CPU drink through the same straw and the fabric buys you nothing. The round trip to DDR or CPU over AXI The on-chip bus protocol family that connects the CPU to the fabric. kills you. The only escape is a model small enough that all of it fits in on-chip memory, where bandwidth is hundreds of GB/s. This is Taalas’ insight, and Groq’s, and Cerebras’: the memory wall is the enemy and on-chip weights are the escape. They spend hundreds of millions of dollars enlarging the on-chip budget. The KV260 gives you about 3 MB.
DDR-resident: no escape Arm A53 PL fabric DDR ~20 GB/s, shared one controller, one straw fabric buys nothing: ~11 tok/s on-chip: the only escape PL fabric URAM + BRAM hundreds of GB/s to TB/s model fits in ~3 MB firehose INT4, 1.5 MB resident: up to 59,965 tok/s being small enough to live on-chip is the entire trick Single-stream decode is memory-bandwidth bound, not compute bound. If the weights live in DDR, the Arm cores and the fabric share one ~20 GB/s controller and the fabric is no faster than the CPU. The only way to win is to make the model small enough that all of it fits in on-chip SRAM, where bandwidth is hundreds of GB/s to TB/s. Everything else in this project follows from that one sentence.
3 MB has to hold the weights, the activations, and the KV cache. That is not enough room for a smart model. It is barely enough room for a model that can string a sentence together. So the second lever, the one the big players mostly can’t pull: shrink the model until the problem disappears.
Making the model small enough (this is where Kevin comes in)
A 3.16M-parameter INT4 4-bit integers (16 levels). The weights are stored as INT4, which is what makes the model small enough to fit on-chip. transformer fits in ~1.5 MB. That’s the budget met, but every byte still counts, and the last lever toward speed kept turning out to be “make the model dumber.” (or so I thought)
So the training corpus (TinyStories) is run through a tool that strips English to its content words. “Why waste time saying a lot of words when a few words do the trick” becomes “why waste time say lot word when few word do trick,” inspired by Kevin Malone’s communication philosophy from The Office, and yes, the model consequently talks like him. Across the full corpus the compression takes 371.7M words down to 260.5M, about 70%, measured.
I expected the compressed-corpus model to come out smaller. It came out exactly the same size, and in hindsight that’s obvious: the parameter count is fixed by the architecture, not the corpus. What the compression does is shrink the output distribution. The same story takes ~30% fewer characters to tell, so effective speed goes up even though the per-token rate is identical.
Honest version: the lemmatised corpus buys about 1.5x. The order-of-magnitude win is on-chip versus DDR. Kevin is the garnish on the chilli, not the meal.
My goal was 100,000 tok/s. I didn’t get there, but I gave it a bloody good go, eventually even compromising the usability by reducing the attention window to T=1 to remove as many round trips and race conditions as possible.
The record build is 16 parallel streams sharing one weight pass, entirely sequenced in the fabric, CPU touching nothing. It measures 59,965.5 tok/s at 200 MHz , 16 of 16 streams bit-exact against the integer reference, three runs of three. The confession: those sixteen streams remember nothing . Each decodes with an attention window of one token. A chat built on it emits one faithful character and falls down the stairs, “he he he he he.” It is fast and it is meaningless, and those are the same property taken one step too far.
So the deployed chat is a different, honest build. One stream, full trained context window, KV caching bit-exact to a full recompute. That one runs 19,242 tok/s of fabric by counted cycles, and ~21,300 tok/s measured live , a number that held flat from 1 to 2,000 concurrent connections in a load sweep with zero errors (peak observed 21,479). It remembers your last couple of turns, and it’s the one in the widget above. Neither number gets to borrow the record’s headline.
The path from 11 tok/s to 60k was a ladder, every rung measured on silicon:
And the ceiling is proven, not assumed. There’s a script that falsifies packing a third MAC Multiply-ACcumulate: the one-multiply-one-add operation matrix multiplies are built from. The fundamental unit of compute here. into a DSP48E2 DSP48E2: the FPGA's dedicated hardware multiplier blocks (1248 on this chip). Each can pack two INT4-by-INT8 multiply-accumulates. over 1.2M randomised products, and N=32 needs 2,048 DSPs the chip doesn’t have. The real limit of this architecture on this silicon is 62k to 78k. 100k was a finger-in-the-air guess before I even started, I just thought it would be a funny target number. Hitting 60k is close enough for me, and >20k usable beat the Taalas number, which was another goal I set myself.
The on-chip trick only wins while the model fits on-chip, and the crossover is around 6.3M parameters. Past that you spill to DDR and you’re back at the wall. Long context spills the KV cache the same way. This is a toy-model technique by construction, and the output is bad on purpose, which does not make it good. It’s a measurement instrument with a sense of humour.
For scale: Taalas etches weights into transistors on a taped-out ASIC per model. Cerebras keeps 44 GB of SRAM on a wafer. Groq keeps 230 MB per chip and gangs hundreds of them. Same one-sentence insight at budgets from nine figures down to the price of a nice dinner.
Everything above is the story. The engineering is in the sub-articles, and every number in them is bit-exact on silicon before it’s quoted:
Inside the chip , the wide-word GEMV trick, the dual-port split-brain, the non-linear bricks, and the proof that 16 streams is a hard ceiling.
Sampling without asking , how one unmeasured host loop ate 58% of every reply, and the Gumbel-max identity that collapsed 193 reads per token into one seed write.
Serving an FPGA to strangers , the tunnel, the serving box, and why the dashboard deliberately lives off the board.
War stories and the method , where iverilog lies, why silicon beats the timing report by 1.3 to 1.76x, and the optimisation I won and stopped anyway.
It’s all hand and LLM written Verilog, no HLS High-Level Synthesis: tools that compile C/C++ into hardware so you don't write Verilog by hand. This project has none; every block is RTL written directly, by me and by Claude Code, never compiled from C. . Claude Code did a solid amount, it’s a side project on a side project after all, though pushing an FPGA to its limit is definitely not as comfortable for it as writing a CRUD app in TypeScript. I would probably put the balance at:
Architecting: me 80%, Claude Code 20%
Implementing: me 20%, Claude Code 80%
And despite saying it’s not that comfortable writing HDL Hardware Description Language (Verilog, VHDL): code that describes circuits rather than instructions to execute. This project is

[truncated]
