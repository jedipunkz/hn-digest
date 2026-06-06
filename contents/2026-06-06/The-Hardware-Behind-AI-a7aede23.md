---
source: "https://www.pathtostaff.com/p/unpacking-ai-the-hardware-behind"
hn_url: "https://news.ycombinator.com/item?id=48425317"
title: "The Hardware Behind AI"
article_title: "Unpacking AI: The Hardware Behind AI - by Sidwyn Koh"
author: "sidwyn"
captured_at: "2026-06-06T15:32:12Z"
capture_tool: "hn-digest"
hn_id: 48425317
score: 4
comments: 1
posted_at: "2026-06-06T14:11:42Z"
tags:
  - hacker-news
  - translated
---

# The Hardware Behind AI

- HN: [48425317](https://news.ycombinator.com/item?id=48425317)
- Source: [www.pathtostaff.com](https://www.pathtostaff.com/p/unpacking-ai-the-hardware-behind)
- Score: 4
- Comments: 1
- Posted: 2026-06-06T14:11:42Z

## Translation

タイトル: AI の背後にあるハードウェア
記事のタイトル: AI の開梱: AI の背後にあるハードウェア - シドウィン・コー著
説明: AI を基礎から学ぶ新しいシリーズ

記事本文:
AI を解き明かす: AI の背後にあるハードウェア - シドウィン・コー著
スタッフエンジニアへの道
AI を解き明かす: AI の背後にあるハードウェア
AIを基礎から学ぶ新シリーズ
Sidwyn Koh 2026 年 6 月 6 日 1 2 シェア スタッフへの道へようこそ。私は最近、個人的な理由で Meta を辞めました (解雇されたわけではありません!) ので、書く時間がずっと増えました。これは、AI についてできる限り多くのことを学び、学んだことを抽出して皆さんと共有することを意味します。
私はエンジニアですが、AI がどのように機能するかについて詳しく調べたことはありません。私は毎日それを使っていますが、テクノロジーからは遠く離れていると感じています。数週間前、私はついに AI を徹底的に理解することにしました。内部構造を理解することで、はるかに使いやすくなりました。
そして、それらの学びを皆さんと共有したいと思います。
この新しいシリーズは Unpacking AI と呼ばれます。
この詳細な説明は 5 部構成のシリーズです。
AI の背後にあるハードウェア。トランジスタ、半導体、製造業者。大手企業 (TSMC、Nvidia、ASML) について学びましょう。メモリ計算のボトルネック。そして、あなたがいつも疑問に思っていたすべての頭字語 (TPU、ASIC、FPGA、CUDA など)
データとモデルのアーキテクチャ。モデルが何でできているかを学びましょう。すべての始まりとなった論文 (「Attending is All You Need」) を取り上げ、さらに変圧器と拡散モデルについてもお話します。そしてもちろん、これらのモデル用にトレーニング データがどのように準備されるか (ソースは何ですか? データはどのように除染され、フィルター処理されますか?) についても説明します。
トレーニング。モデルを教える仕組み。事前トレーニングはどのように機能しますか?これには何が含まれますか (バックプロパゲーション、オプティマイザー、損失関数)?高価なトレーニング (最大数億ドル) を開始する前に、どのようなスケーリングの法則を理解しておく必要がありますか?
トレーニング後の調整と調整。モデルを教えた後、どのようにモデルをガイドすればよいでしょうか?安全性をどのように適用するか?どのようにベンチマークを行うか

モデルが改良されたことをご存知ですか?モデルのパフォーマンスをどのように評価するのでしょうか?
推論、サービス、エージェント。これは、AI ユーザーとして最も身近なトピックであるため、最も馴染みのあるトピックかもしれません。モデルはどのようにトークンを出力し、その結果をユーザー (SSE) に提供しますか?システムはどのようにして公平かつ高速に保たれるのでしょうか?どのようなツールが利用可能ですか (MCP、RAG、ツールの使用)、エージェントはどのように機能しますか?
このシリーズを読み進めるうちに、AI についてさらに学ぶにつれてシラバスも変わると思います。いくつかの頭字語が含まれるセクションについては、セクションの先頭にその定義もリストします。コメントでの質問もお待ちしております！これはこのシリーズを改善するのに役立ちます。
簡単なメモ: 私は AI を使用するのが大好きですが、この記事と今後の記事は手書きで書かれ、AI による簡単な編集が行われます。結局のところ、私たちは皆、AI の文章を読むのにうんざりしています。ただし、一部の画像はAIによって生成されます。
トランジスタとその重要性
人工知能を理解するには、まずその核心に迫る必要があります。
AI は GPU チップを使用して実行されます。これらのチップは、EUV マシンを使用して製造されたトランジスタから作られています。
トランジスタから始めて、これらをそれぞれ分解してみましょう。
トランジスタは、電気の流れを制御する半導体デバイスです。 1 つの端子で小さな電気信号を使用して、はるかに大きな電流を制御します。 (1) 信号をブーストするか、(2) 電流を流すことができるかどうかを決定します。言い換えれば、アンプまたはスイッチとして機能します。
半導体、最も一般的にはシリコンは、特定の条件下でのみ電気を伝導する材料です。不純物を添加することで導電率を変えることができます。
誰がこれらのチップを設計しているのでしょうか?チップの設計に関しては、Nvidia と AMD が最大のプレーヤーです。チップ設計に関しては、Google、Amazon、そして今では Meta にもそれほど遅れはありません。しかし、これらの会社が運営しているのは、

ファブレスデザイナーとして。つまり、彼らはアーキテクチャのみを構築し、物理的な生産はファウンドリに委託しており、構築と維持には最大 200 億ドルの費用がかかる可能性があります。
では、これらのトランジスタは誰が作っているのでしょうか？ TSMC（台湾積体電路製造会社）はそうしています。ただし、極紫外線 (EUV) 装置と呼ばれる特殊な装置と、チップ上に印刷するリソグラフィーと呼ばれるプロセスが必要です。サムスンやインテルなどのファウンドリは他にもありますが、現在世界のファウンドリ収益の 70% を握っている TSMC ほど先進的ではありません。
これらのEUVマシンは誰が製造しているのでしょうか?これらは現在、EUV 装置業界で独占的な足場を持つオランダの企業である ASML (Advanced Semiconductor Materials Lithography) によってのみ製造されています。中国は急速に追いつきつつあるが、依然として約5年遅れている。
ASML 研究所の写真 (出典) 興味深い事実: 現在、ASML に対する主要な競合他社は存在しません。彼らが今日の段階に到達するまでに 30 年かかりました。彼らは数千のサプライヤーを統合して、毎秒 50,000 個の液滴を発射する発電機を構築しました。彼らは、これらのEUV光源を製造する大手企業Cymerも所有しています。これらの光源は非常に短い (13.5nm と短い) ため、自然光源は存在しません。
さて、トランジスタとは何か、そしてそれがどのように作られるのかがわかりました。これで、GPU (グラフィックス プロセッシング ユニット) について理解できるようになりました。 1 つの GPU には数十億個のトランジスタが含まれており、これらはさまざまな製造技術で製造されたダイ (シリコンの生のブロック) 上に詰め込まれています。
ここでは、これらのさまざまなタイプの製造技術について説明します。
これは役に立ちましたか?購読して、Unpacking AI シリーズの残りの部分を見逃さないようにしてください。
ダイシュリンク: 50 年で 10,000nm から 2nm へ
die shriの歴史を簡単に振り返ってみましょう

nks 。
なぜダイを小さくする必要があるのでしょうか?フィーチャが小さいほど、平方ミリメートルあたりのトランジスタ数が多くなります。これにより、TSMC のようなメーカーは、各チップにより多くの機能を詰め込み、より高性能なチップを構築できるようになります。より多くのコア、より多くのキャッシュ、より多くのテンソル ユニット。ただし、注意すべき点は、構築するフィーチャを小さくすると、ウェーハあたりの欠陥が増加することです。そのため、必ずしも生産コストが安くなるとは限りません。
これらのダイとダイの縮小の歴史は何ですか?複雑ですが、数段落に分けて説明していきます。最初のマイクロプロセッサである Intel 4004 は 1971 年に 10µm のプロセス線幅で発売されました。これは、ゲート長（ドレイン電極とソース電極間の距離）が10μmであることを意味します。つまり、トランジスタがオンになるたびに、電子はこの 10,000nm を通過する必要があるということになります。ゲートが短いほどスピードが速くなります。
Intel 4004。写真は Wikipedia から引用。当時 (1971 年)、このチップは日本の電卓会社、ビジコム向けに設計されました。しかし、インテルは、これが大衆市場にとってはるかに有益であると認識すると、ビジコムからマーケティングおよびテクノロジーの権利を買い戻しました。
数年前、ゴードン・ムーアは、皆さんもムーアの法則として聞いたことのある観察を発表しました。チップ上のトランジスタの数は、低コストでおよそ 2 年ごとに 2 倍になり続けるというものです。
次の数十年間で、600nm → 250nm → 180nm → 130nm → 45nm と進みました。 2000年代初頭、メーカーは壁にぶつかりました。次のジャンプをしてゲートの長さを短くする方法はありませんでした。しかし、Burn-Jeng Lin という名前の TSMC エンジニアは、レンズとウェハの間に水を加えるという画期的な発見をしました。これは2003年から2004年にかけてのASMLによる大きな賭けであったが、ASMLは当時ニコンやキヤノンに次ぐ小規模なヨーロッパの挑戦者であった。彼らは全力を尽くして勝利を収めました。
ニコ

n 社とキヤノンは 157nm ドライリソグラフィーの開発に固執しましたが、その時にはもう手遅れでした。 ASML は非常に有利なスタートを切りました。キヤノンは本質的に最先端のリソグラフィーから撤退し、ニコンは最終的に液浸ツールを開発したものの、リードを取り戻すことはなく、その後EUVから完全に撤退した。
現在、iPhone は 3nm チップで動作しています。現在の GPU はすべて、TSMC の 5/4/3 nm バリアントを使用しています。現在の開発状況は 2nm で、2026 年後半から 2027 年にかけて 1.6nm (TSMC の「A16」) に達するという目標があり、1.4nm はさらに遅く、本当の 1nm になるのは 10 年代の後半になると予想されています。
残念ながら、悲しい事実の 1 つは、これらの数値はもはやゲート長を意味するものではないということです。それらはマーケティングに多く使用されます。実際に改善されるのは、MTr/mm² (1 平方ミリメートルあたり数百万個のトランジスタ) で測定されるトランジスタ密度ですが、これさえもファウンドリ全体で一貫して測定されているわけではありません。
50 年にわたるダイの縮小の歴史 CPU から GPU への移行
そもそも GPU がどのようにして有名になったのかについて少し話しましょう。すべては CPU から始まりました。
CPU は 1971 年に最初のマイクロプロセッサが登場してから使用されています。ただし、1990 年代に Quake のようなゲームが導入されたとき、グラフィックスをレンダリングしようとするとかなりの遅れが生じました。激しいゲーム (DotA など) をプレイするたびに、自分のコンピューターが停止してしまうのを覚えています。
グラフィック アクセラレータ カードがその答えでした。いくつかの高度なコアを用意する代わりに、何千ものダム コアを並行して実行することになります。個々のコアは非常に弱いですが、GPU で組み合わせるとスループットは巨大になります。 4K 画像をレンダリングするには 8M ピクセルの色を個別に計算する必要があるため、これらはゲームに最適でした。 NVIDIA はこれらのカードを最初に構築したわけではありません (3dfx と呼ばれる他の企業と ATI が構築しました) が、GeForce 2 を組み合わせて GPU という用語を造語しました。

1999年には56歳。
2006 年に遡ると、NVIDIA の CEO であるジェンセン ファンは大きな賭けをしました。ムーアの法則は減速しているということです。シングルスレッド CPU のパフォーマンスは長期的には最適化されていませんでした。彼は、グラフィックス カード上に科学計算用のプログラミング プラットフォームを構築したいと考えていました。この賭けは、スーパーコンピューターへのアクセスを望む科学者をターゲットにしていました。当時、これらのスーパーコンピューターは、政府の研究所と少数の企業だけが所有していた数百万ドルのマシンでした。
この賭けは CUDA (Compute Unified Device Architecture) と呼ばれていました。これにより、CPU は並列計算タスクを CPU から GPU にオフロードできるようになりました。これが結局彼らの堀となった。エコシステム (PyTorch、TensorFlow、後の章で説明します) は、最終的に CUDA ファーストで構築されました。シリコンとネットワークもこのアーキテクチャに特化しました。
GPU を活用する AI の最初のヒントは、2012 年に生まれました。トロント大学の 3 人の研究者、Alex Krizhevsky、Ilya Sutskever、Geoffrey Hinton は、 AlexNet と呼ばれるニューラル ネットワークを提出しました。 (これらの名前を知り始めるかもしれません!) これは、Alex の寝室にある 2 つの NVIDIA GTX 580 ゲーム GPU でトレーニングされました。
これにより、GPU がディープ ニューラル ネットワークをトレーニングできることが初めて証明されました。これらのニューラル ネットワークは主に行列の乗算に基づいており、誰かの寝室で実現できました。同じニューラル ネットワークを CPU でトレーニングした場合、何世紀もかかるでしょう。
CPU (左) と GPU (右) を比較すると、次のようになります。
GPU の構造。 GPU にはさらに多くのコアがあることがわかります。 CUDA プログラミング ガイドから引用。これは役に立ちましたか?購読して、Unpacking AI シリーズの残りの部分を見逃さないようにしてください。
次に GPU を見てみましょう。
かなりの警告: ここからは非常に技術的な話になります。私はそうします

それを壊すために最善を尽くしてください。
1 つ目は、Blackwell GPU (2024 年第 4 四半期に発売) の全体図です。これは最新のチップ アーキテクチャではありません。Rubin R100 は最近発表され、数か月以内に出荷される予定です。残念ながら、R100 に関する適切なインフォグラフィックが見つからなかったようです。もし見つけたらお知らせください。
それでも、Blackwell Ultras がエンドツーエンドでどのように動作するかがよくわかるので、Blackwell Ultras を調べてみましょう。
Blackwell の概要からの画像のクリーンアップされたバージョン。ここ Blackwell には、 NV-HBI と呼ばれるカスタム相互接続で溶接された 2 つのダイがあります。インターコネクトは基本的に、2 つのものを接続する物理的なワイヤーです。この場合、NV-HBI は、10 TB/秒の電力を供給する超低遅延の独自のダイツーダイ相互接続です。
さて、それぞれのダイを都市として考えてみましょう。各都市には次のものが含まれます。
4 つのグラフィックス処理クラスター (GPC)
GPC は都市内の地区です。これは、番号を伝達するために地元のインフラストラクチャを共有する工場のグループが集まる地区と考えてください。 GPC 内には 20 個の SM があり、これはダイごとに 80 個の SM があり、両方のダイで合計 160 個の SM があることを意味します。すごい計算能力ですね！
ギガスレッド エンジン + MIG コントロール
これは都市レベルのディスパッチャです。その仕事は単純です。 CPU から (PCIe Gen 6 経由で) 作業を受け取り、ファームします。

[切り捨てられた]

## Original Extract

A new series on learning AI, from the ground up

Unpacking AI: The Hardware Behind AI - by Sidwyn Koh
Path to Staff Engineer
Subscribe Sign in Unpacking AI: The Hardware Behind AI
A new series on learning AI, from the ground up
Sidwyn Koh Jun 06, 2026 1 2 Share Welcome back to Path to Staff. I recently left Meta for personal reasons ( was not laid off! ), and have found much more time to write. This means learning as much as I can about AI, and distilling what I’ve learned and sharing with you all.
I’m an engineer, but never really dug into how AI works. I use it every day, yet I feel so far away from the tech. A few weeks ago, I finally dove really deep to understand AI from the bottom up. Understanding its internals has helped make me far better at using it.
And I want to share those learnings with you.
This new series is called Unpacking AI .
This deep dive is a five-part series:
The Hardware Behind AI. Transistors, semiconductors, and fabricators. Learn about the big players (TSMC, Nvidia, ASML). The memory-compute bottleneck. And all the acronyms you always wondered about (TPU, ASIC, FPGA, CUDA, etc.)
Data & Model Architecture. Learn about what models are made of. We’ll cover the paper that started it all (”Attention is All You Need”), plus talk about transformers and diffusion models. And of course, we’ll cover how training data is prepared for these models (what sources? how is the data decontaminated and filtered?)
Training . The mechanics of teaching a model. How does pretraining work? What goes into it (backpropagation, optimizers, loss functions)? What scaling laws should we understand before we kick off an expensive training run (up to hundreds of millions $)?
Post-Training & Alignment. How does one guide a model once it’s been taught? How do we apply safety? How do we benchmark and know the model got better? How do we evaluate a model’s performance?
Inference, Serving and Agents. This might be the most familiar topic, since it’s closest to you as an AI user. How does a model output its token and serve the result to you (SSE)? How do systems stay fair and fast? What tools are available (MCP, RAG, tool use) and how do agents work?
Over the course of this series, I expect the syllabus to change as I learn more about AI. For sections with several acronyms, I also list their definition at the top of the section. I also welcome questions in the comments! This will help me improve this series.
Quick note: As much as I love using AI, this article and future ones will be written by hand, with light editing by AI. After all, we’re all too sick and tired of reading AI slop. Some images, however, will be generated by AI.
Transistors and Their Importance
To understand Artificial Intelligence, we first have to go to the core of it.
AI runs off GPU chips. These chips are made from transistors which are manufactured using EUV machines.
Let’s break each of these down, starting with a transistor.
A transistor is a semiconductor device that controls the flow of electricity. It uses a small electrical signal at one terminal to control a much larger current. It either (1) boosts a signal, or (2) decides whether a current can pass. In other words, it acts as either an amplifier or switch.
A semiconductor , most commonly silicon, is a material that conducts electricity only under certain conditions. Its conductivity can be modified by adding impurities.
Who designs these chips? Nvidia and AMD are the biggest players when it comes to designing chips. Not far behind are Google, Amazon and now Meta when it comes to chip design. However, these companies operate as “fabless” designers. That means they only build the architecture, while outsourcing the physical production to foundries, which can cost up to $20B to build and maintain.
Who then makes these transistors? TSMC (Taiwan Semiconductor Manufacturing Company) does. But they require special machines, called extreme ultraviolet (EUV) machines, and a process called lithography, which is the act of printing on chips. There are other foundries like Samsung and Intel, but they are not as advanced as TSMC, which currently holds 70% of the global foundry revenue.
Who makes these EUV machines? These are currently only being manufactured by ASML (Advanced Semiconductor Materials Lithography), a Dutch company that has a monopoly foothold in the EUV machine industry. China is fast catching up, but is still roughly 5 years behind.
A picture of ASML laboratory ( source ) Fun fact: there are no major competitors to ASML today. It took them 30 years to reach the stage they’re in today. They’ve integrated thousands of suppliers together to build a generator that fires 50,000 droplets per second. They also own the major company Cymer that makes these EUV sources. These light sources are so short (as short as 13.5nm) that there is no natural source for it.
OK, now we know what a transistor is and how it’s made. Now we can understand GPUs (graphics processing units). A single GPU contains billions of transistors, which are packed onto a die (a raw block of silicon) manufactured with various fabrication technologies.
We’ll now get into these different types of fabrication technologies.
Found this helpful? Subscribe to make sure you don’t miss the rest of the Unpacking AI series.
Die Shrinks: Going from 10,000nm -> 2nm in 5 decades
Let’s take a quick history detour of die shrinks .
Why do we need a die to be smaller? Smaller features mean more transistors per square millimeter. This allows manufacturers like TSMC to pack more into each chip and build more capable chips. More cores, more cache and more tensor units. However, one point to note is that by building features smaller, this leads to more defects per wafer. As such, it isn’t necessarily cheaper to produce.
What’s the history of these dies and die shrinks? It’s complicated, but I’ll try to explain in a couple of paragraphs. The first microprocessor, the Intel 4004 , was launched in 1971, at a 10µm process line width. This means that the gate length (distance between drain and source electrodes) was at 10µm. Which in turn means that electrons have to travel across this 10,000nm whenever a transistor switches on. A shorter gate = faster speed.
The Intel 4004. Picture taken from Wikipedia . At that point in time (1971), this chip was designed for a Japanese calculator company, Busicom. However, once Intel realized that this was much more useful for the mass market, Intel repurchased the marketing and technology rights from Busicom.
A few years earlier, Gordon Moore made the observation you’ve heard of as Moore’s Law : that the number of transistors on a chip would keep doubling at low cost roughly every two years.
Over the next few decades, we went from 600nm → 250nm → 180nm → 130nm → 45nm . In the early 2000s, manufacturers hit a wall. There was no way to take the next jump and shorten gate length. However, a TSMC engineer named Burn-Jeng Lin had a breakthrough: adding water between the lens and the wafer. This was a huge bet by ASML in 2003-04, which was at that time a smaller European challenger behind Nikon and Canon. They went all in on immersion and won.
Nikon and Canon stuck to their guns on 157nm dry lithography, but by then it was too late. ASML had a huge head start. Canon essentially exited leading-edge lithography, and while Nikon did eventually build immersion tools, it never recovered its lead, and later sat out EUV entirely.
Today, your iPhones run on 3nm chips. GPUs today all use TSMC’s 5/4/3 nm variants. The state of development is currently at 2nm, and there’s targets to hit 1.6nm (TSMC’s “A16”) around late 2026–2027, with 1.4nm later and true 1nm not expected until the back half of the decade.
Unfortunately, one sad fact is that these numbers no longer mean gate lengths. They’re used more for marketing. The thing that actually improves is transistor density, measured in MTr/mm² (millions of transistors per square millimeter), and even that isn’t measured consistently across foundries.
Five decades of shrinking dies The Shift from CPU to GPU
Let’s talk a bit about how GPUs got famous in the first place. It all started with the CPU .
CPUs have been in place since 1971, since the first microprocessor. However, when games like Quake were introduced in the 1990s, they lagged pretty badly while trying to render graphics. I remember my own computer grinding to a halt whenever I played intense games (DotA, anyone?).
Graphics accelerator cards were the answer. Instead of having a few sophisticated cores, you’d have thousands of dumb cores running in parallel. Each individual core is super weak, but when combined together in a GPU, the throughput is gigantic. These were great for games, since rendering a 4K image meant computing colors of 8M pixels independently. NVIDIA wasn’t the first to build these cards (other companies called 3dfx and ATI did), but they did coin the term GPU with the GeForce 256 in 1999.
Fast forward to 2006, Jensen Huang, CEO of NVIDIA made a huge bet. That Moore’s Law is slowing. Single-threaded CPU performance was not optimized for the long run. He wanted to build a programming platform for scientific computing on graphics cards. This bet was targeted at scientists who wanted access to supercomputers. At that point of time, these supercomputers were multi-million dollar machines only owned by government labs and a few corporations.
This bet was called CUDA (Compute Unified Device Architecture). This allowed the CPU to offload parallelized computing tasks from the CPU to the GPU. This ended up being their moat. The ecosystem (PyTorch, TensorFlow, which we will cover in later chapters) ended up being built CUDA-first. Silicon and networking was also specialized around this architecture.
The first hint of AI leveraging GPUs came about in 2012. Three University of Toronto researchers, Alex Krizhevsky, Ilya Sutskever and Geoffrey Hinton submitted a neural network called AlexNet . (You might start to know these names!) This was trained on two NVIDIA GTX 580 gaming GPUs in Alex’s bedroom.
This proved that GPUs were feasible to train deep neural networks for the first time. These neural networks, which were mostly based on matrix multiplication, were able to be achieved in someone’s bedroom. If the same neural net were to be trained on a CPU, it would have taken centuries.
If we were to compare a CPU (left) to a GPU (right):
How a GPU is structured. You can see that there are many more cores in a GPU! Taken from the CUDA programming guide . Found this helpful? Subscribe to make sure you don’t miss the rest of the Unpacking AI series.
Now let’s take a look at a GPU.
Fair warning: it starts to get very technical from here on out. I will try my best to break it down.
The first is an overall view of the Blackwell GPUs (launched Q4 2024). This is not the latest chip architecture: Rubin R100 was recently announced and plans to ship in a few months. I could not seem to find a good infographic on R100, unfortunately, so let me know if you do.
Nevertheless, let’s examine the Blackwell Ultras, since this does give us a good sense of how it works end-to-end.
Cleaned up version of an image from Blackwell’s overview . Here in Blackwell, we have 2 dies that have been welded together with a custom interconnect called NV-HBI . An interconnect is basically a physical wire linking two things together. In this case, NV-HBI is an ultra-low latency, proprietary die-to-die interconnect that powers 10 TB/s.
Now, let’s think of each die as a city . Each city contains:
4 Graphics Processing Clusters (GPCs)
A GPC is a district within the city. Think of it as a district housing a group of factories that share some local infrastructure to pass numbers around. Within a GPC, there are 20 SMs, which means that there are 80 SMs per die, and 160 SMs total across both dies. That’s a lot of compute power!
GigaThread Engine + MIG Control
This is the city-level dispatcher. Its job is simple. It receives work from the CPU (through the PCIe Gen 6) and farm

[truncated]
