---
source: "https://jdagostino.github.io/ai-pt1-box-o-scraps/index.html"
hn_url: "https://news.ycombinator.com/item?id=49288293"
title: "We Have AI at Home Chapter 1: A Box of Scraps"
article_title: "AI At Home Part 1: A Box Of Scraps"
author: "timmmmmmay"
captured_at: "2026-08-13T16:45:20Z"
capture_tool: "hn-digest"
hn_id: 49288293
score: 2
comments: 0
posted_at: "2026-08-13T16:22:05Z"
tags:
  - hacker-news
  - translated
---

# We Have AI at Home Chapter 1: A Box of Scraps

- HN: [49288293](https://news.ycombinator.com/item?id=49288293)
- Source: [jdagostino.github.io](https://jdagostino.github.io/ai-pt1-box-o-scraps/index.html)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T16:22:05Z

## Translation

タイトル: 我が家にも AI がある 第 1 章: スクラップの箱
記事のタイトル: AI アットホーム パート 1: スクラップの箱
説明: データセンターはとてもクールなので、どの家にもデータセンターがあるはずです

記事本文:
ソフトウェア開発者にとって、今は興味深い時代です。私の意見では、トランスフォーマー大規模言語モデルは、この分野における久しぶりの本当に新しくて興味深い技術開発です。これに基づいて構築された AI コーディング エージェントが急速にこの職業の核となり、良くも悪くも手工具を使った木工作業から電動工具を使うようになったような感じです。
私は、テクノロジーを信頼できるのは、それを分解してガレージで組み立てることができる場合だけです。サービスに接続して、他人の制御下でコーディング エージェントを使用するということは、私を悩ませます。これがうまくいかない方法はたくさんあります。あなたはそれに依存し始め、そしてそれはあなたから奪われます。このテクノロジー全体はエキサイティングですが、私はどこかの大きな AI データセンターに接続したくありません。すべての家庭に小さな AI データセンターが欲しい!そして現在、コンピューターの部品が不足している状態です (あるいは、2010 年代後半はコンピューターの部品が余っていたのでしょうか?) ので、手頃な価格を保つためには、ゴミからそれを構築する必要があります。
ここで主に必要なのは、大量の GPU です。多くの並列行列計算を行うものには高速メモリが必要であり、それが GPU です。ただし、前に述べたように、AI ワークロードを実行するように設計されたものはすべて、あるいは AI に優れていることが知られているものであっても、現時点では非常に高価です。
2022 年に遡ると、多くの企業が「クラウド ゲーム」全体を普及させようとしていました。ビデオ ゲームはどこかのデータ センターで実行されるので、高価なゲーム PC やコンソールを購入する必要はありません。 AMD はワークステーション GPU の 1 つを採用し、それに追加の RAM を追加し、すべてのビデオ出力を削除しました。彼らはその結果生まれたカードを「V620」と呼びました。それはありませんでした

一般に販売されているため、聞いたことがないかもしれません。 AMD はこれらのものを過剰に制作し、その後、ラグに関する (今から思えば明白な) 問題のせいで、クラウド ゲーム全体がうまくいきませんでした。
これらのカードは実際には AI ワークロード向けに設計されておらず、AMD のソフトウェア サポートは悪名高いので、eBay の「中古サーバー ハードウェア」電子廃棄物再販業者から比較的安価で購入できます。私が入手したものは実際に使用されたことはないと思いますが、基本的には新品のように見えます。それぞれに 32GB のかなり高速な VRAM が搭載されています。確かに、それらは私が使用したい用途には適さないという評判がありますが、これを実際に機能させるのはどれほど難しいでしょうか?
ああ、ところで、彼らにはファンがいません。これらはサーバー カードであり、サーバーが非常に大音量の送風ファンでカードを冷却することを期待しています。それについてはすぐに説明します。
ここでは部品が不足しているので、GPU アレイを結び付けるために別の電子廃棄物を購入しました。マザーボードは 2017 年の X299 プラットフォームのもので、誰かの古いゲーム機からのもので、eBay で安く買えるほど古いものでした。 4 つの PCIe x16 スロットが適切な間隔であり、4 つの GPU を並べて取り付けることができます。ここで私が気にしたのはそこだけです。搭載されている CPU は Intel Core i9 10900X で、これは Intel が過去 20 年間に製造した中で最悪の CPU です。 2019年に発売されたとき、それはSkylakeチップの再再リフレッシュ版であり、当時Intelがプロセッサ設計またはリソグラフィノードサイズのいずれかにおいてAMDのRyzenチップに完全に追いつくことができなかったため、価格が高すぎて消費電力が高かった。わかりました。つまり、今は安いので、仕事はうまくいきます。
RAM と SSD は家中の他のコンピューターから回収されました。もちろん、これは不正行為です。今買わなければならないとしたら、うっとうしいほど高価になるでしょう。ただし、ここでの要件はあなたが思っているほど厳しくありません

考える;すべての作業は GPU の VRAM 上で行われます。モデルがディスクからロードされた後、これらはほとんどアイドル状態のままになります。まだサルベージパーツを持っていない場合は、ここでかなり安く買えるでしょう。
4 つの GPU と Intel の最も効率の悪いプロセッサに電力を供給しているので、大きな電源を確保しました。何らかの理由で、eBay では 1200W の代わりに 1600W の方が安かったのです。それほど多くの電力を常に消費することは期待していませんが、起動中にすべてが同時にオンになることに対処する必要があります。
新しいケースとファンに散財しました。そうそう、ファンの皆さん！これらのデータセンターの GPU には冷却ファンがなく、サーバー ファンの壁からの空気の流れを期待しています。これらのカードの 1 つを冷却したい場合は、オンラインで 3D 印刷可能なファン シュラウド モデルがたくさんありますが、4 つ並べて積み重ねると、どれも最適ではないように見えます。小型で超大音量の 40mm ファンが端に付いているか、横に突き出ていてカードを並べて置くことができないかのどちらかです。 4 枚のデュアル スロット カードを並べると幅は約 160 mm になります。そのため、私が本当に必要としているのは、カードのヒートシンクを吹き飛ばす 2 つの 80 mm サーバー ファンです。
そこで、80mm ファンを 2 枚のカードに取り付けるシュラウドをモデリングしてプリントしました。 1 カードの背面には、小さなネジで固定されている、ある種の金属製のケーブル ガイド フィンがありました。このネジ穴を使ってシュラウドを取り付けました。そこには電気コネクタ用の切り欠きと、背面にファンを取り付けるための 4 つの穴があります。私はこれをカーボンファイバー ASA から印刷しましたが、おそらく退屈な古い PLA でもうまく機能したでしょう。
私がこれらの 10,000 RPM ファンを購入したのは、十分な空気を確実に送りたかったためです。価格は低速のファンと同じでした。本当にうるさいです！マザーボードが GPU の温度に基づいて速度を制御するようにしたかったのですが、これは最初は機能しませんでした。

o 初期セットアップ時に耳保護具を着用しただけです。
最初はボックスが起動しませんでした。 4 枚のカードすべてが接続された状態でマザーボードを実際に起動するには、どの PCIe 設定を設定する必要があるかを理解するまで、earpro を着てガレージに立って BIOS 設定をいじって 1 時間ほど費やしたと思います。 (高度なオプション設定の奥深くにある 2 つの別個のメニューで、[サイズ変更可能な BAR] と [MMIO 高サイズ] を有効にする必要があります。2017 年には、このボードにこれほど多くの VRAM を接続するとは誰も考えていなかったでしょう)。この頃、ガレージにイーサネットがないことにも気づき、午後 11 時から乾式壁に穴を開け始めました。
ともかく！それらのいくつかの間違ったスタートの後、私はこれを起動し、そこに Ubuntu 24.04 のインストールを開始しました。耳栓をしていてもファン制御がない状態で、私のお気に入りの推論サーバーである優れた llama.cpp のビルドをコンパイルし、1 枚のカードに快適に収まる Gemma4 モデルをテストしました。これは、単一の GPU ワークステーションで遊んでいたときの私のお気に入りのローカル モデルで、非常にうまくテストできました。以前使用していた RTX 3090 ほど高速ではありませんが、まったく悪くありません。次に、このビルドの本当のターゲットである Deepseek V4 Flash を起動します。遅いよ！現時点では、それを高速化する方法はわかりません (詳細は次の章で説明します)。ただそれが適合するかどうかを確認したかっただけで、適合しました。
ファンの元に戻りましょう！繰り返しますが、これらがどれほどうるさいか強調してもしすぎることはありません。全開にすると家の壁越しに音が聞こえますが、これは全開にする必要はありません。当初の計画では、マザーボードの内蔵ファン制御を使用してファンを制御することでしたが、このボードは、異なるファンを異なる速度で制御することを拒否しました。どうやらこれはSupermicroのマザーボードに共通しているようです。そこでファンコントロールを作りました

えー、その代わりに。私は、関税が課される前の時代に Aliexpress から入手した、ノーブランドの Arduino Nano クローンを箱ごと持っています。私は 10 年以上前に大学のマイクロコントローラーのクラスで PWM モーターを実行するために書いたコードを掘り出し、クリーンアップしました。
ファン コントローラーはサーバーからパーセンテージを取得するだけです。温度を読み取り、ファン速度を適切に調整するためのスクリプトがサーバー上に必要です。これは AI サーバーなので、独自のファン制御スクリプトを作成させました。それはぴったりだと思えた。 Deepseek V4 Flash はこれに非常に優れており、軽量 AI モデルから適切なソフトウェアを取得するために必要な通常量のガイダンスと人間の対話が必要です。 2 スクリプトをテストできるように GPU を加熱する何かを考えるように指示すると、スクリプトが行列の乗算を行う PyTorch スクリプトを書き始めました。そこで私が「いいえ、これらのカード上で実行しています。このコンピュータ上の llama.cpp インスタンスに存在します。何かを言うだけでカードが読み込まれてしまいます。」と伝えると、一時的に存続の危機に陥りました。愛らしい！
コントローラー自体はプロトタイピング用のブレッドボード上にあり、サーバー ケースに詰め込んだところです。ケースに接触してショートしないように、絶縁テープを少し貼りました。
この時点で、ボックスはかなり安定しており、GPU をテストし、高速に実行する方法に本格的に取り組む準備が整いました。これについては次の章で説明します。
シュラウドの STL ファイルをダウンロードし、必要に応じて印刷します。 [戻る]
Arduino 用のコードとサーバー スクリプトはここから入手できます。 [戻る]

## Original Extract

data centers are so cool there should be one in every house

It's an interesting time to be a software dev; the transformer large language model is, in my opinion, the first really new and interesting technological development in the field in a long time. AI coding agents built on this have rapidly become core to working in this profession, and the feeling is kind of like going from doing woodworking with hand tools to using power tools, for better or for worse.
Me, I've only ever been able to trust a piece of technology if I can take it apart and put it together in my garage. The thing where you connect to a service and use a coding agent under somebody else's control bugs me; there's too many ways for this to go wrong. You'll start depending on it and then it'll get taken away from you. This whole technology is exciting, but I don't want to go and connect to some big AI data center somewhere; I want a little AI data center in every home! And we're in something of a computer parts shortage right now (or, maybe the late 2010s were a computer parts surplus?) so to keep things affordable, I'm going to have to build it out of garbage.
The main thing I need here is a bunch of GPUs. I need fast memory going into something that does lots of parallel matrix math and that's what a GPU is. Like I said earlier, though: everything that's designed to run any AI workload at all, or even anything that's known to be good at it, is wildly expensive right now.
Back in 2022, a bunch of companies were trying to make the whole "cloud gaming" thing take off. You know, where you run your video games in a data center somewhere so you don't need to buy an expensive gaming PC or console. AMD took one of their workstation GPUs, gave it some extra RAM and removed all of the video outputs. They called the resulting card the "V620". It was never sold to the public, so you might not have heard of it. AMD wildly overproduced these things and then the whole cloud gaming thing didn't work out because of (in retrospect obvious) problems with lag.
These cards aren't really designed for AI workloads and AMD's software support is notoriously bad, so these things are relatively cheap to buy from the "used server hardware" e-waste resellers on eBay. I don't think the ones I got had actually ever been used, they look basically new. They each have 32GB of pretty fast VRAM. Sure, they have a reputation for being bad at the thing I want to use them for, but how hard can it be to get this to actually work?
Oh, they don't have fans, btw. They're server cards, they expect the server to cool them with some extremely loud blower fan. We'll get to that in a minute.
We're in a parts shortage here so I bought some other e-waste to tie the GPU array together. The motherboard is from the X299 platform from 2017, it's from somebody's old gaming rig and was old enough to be cheap on eBay. It has four PCIe x16 slots at the correct spacing for me to stick four GPUs side-by-side, which is the only thing I cared about here. The CPU in it is the Intel Core i9 10900X, which is the worst CPU that Intel has produced in the last twenty years; when it launched in 2019 it was a re-re-re-refreshed version of the Skylake chips, overpriced and power-hungry because Intel couldn't quite catch up with AMD's Ryzen chips at the time, either on processor design or lithography node size. Fine with me, that means it's cheap now, and it'll get the job done!
RAM and SSD were salvaged from other computers around the house. This is cheating, of course; if I had to buy them now they'd be annoyingly expensive. But the requirements here are less severe than you'd think; all of the work will be getting done on the GPU's VRAM; after the model loads from disk these mostly stay idle. If I didn't have salvage parts already, I could go pretty cheap here.
I'm powering four GPUs and Intel's least efficient processor so I got a big power supply. For some reason it was cheaper on eBay to get a 1600W one instead of a 1200W one. I don't expect to be drawing that much power constantly, but I need it to handle everything turning on at once during startup.
Splurged on a new case and fans. Oh, right, the fans! Those data center GPUs do not have cooling fans, they expect airflow from a wall of server fans. There's a bunch of 3D-printable fan shroud models online if you want to cool one of these cards, but if you stack up four of them side-by-side these all seem suboptimal; they either use a tiny super-loud 40mm fan right at the end, or they stick out way to the side and you can't put a bunch of cards next to each other. Four dual-slot cards next to each other is a width of about 160mm, so what I really want is two 80mm server fans right there, blowing over the cards' heatsinks.
So I modeled and then printed a shroud that would attach an 80mm fan to two cards. 1 The cards had some kind of metal cable guide fin on the back, held on by little screws; I used those screw-holes to attach the shroud. There's a cutout in there for the electrical connectors and four holes to mount the fans to the back. I printed this out of carbon fiber ASA, but probably boring old PLA would have worked just fine.
I got these 10,000 RPM fans because I wanted to make sure I was moving enough air, and they were the same price as slower fans. They are really loud! I wanted the motherboard to control their speed based on the GPU temperature, and this didn't work at first, so I just wore ear protection during initial setup.
The box didn't want to boot at first. I probably spent an hour over here wearing earpro standing in the garage playing around with BIOS settings until I figured out which PCIe settings needed to get set in order for the motherboard to actually start up with all four cards connected. (you gotta enable Resizable BAR and MMIO High Size, in two separate menus deep in the advanced option settings, because nobody in 2017 thought you would plug this much VRAM into this board). Around this time I also realized that I didn't have Ethernet in the garage, so I started cutting holes in the drywall at 11:00 PM.
Anyway! After those few false starts, I got this thing to boot and started installing Ubuntu 24.04 on there. With earplugs in and still no fan control, I compiled a build of my favorite inference server, the excellent llama.cpp , and tested out the Gemma4 model, which fits comfortably into one card. This was my favorite local model when I was playing around on a single GPU workstation and it tested out pretty well; not quite as fast as it did on the RTX 3090 that I used to have, but not badly at all. Then I spin up Deepseek V4 Flash, which was the real target for this build. It's slow! At this point I have no idea how to make it fast (more on that next chapter), I just wanted to see if it would fit, and it did.
Back to the fans! Again I cannot stress enough how loud these are. At full blast you can hear them through the walls of the house, and this thing just does not require full blast. The initial plan was to control the fans using the motherboard's built-in fan control, but this board refuses to control different fans at different speeds. Apparently this is common for motherboards from Supermicro. So I built a fan controller instead; I have a whole box of off-brand Arduino Nano clones that I got from Aliexpress in the pre-tariff days; I dug up some code I wrote for a microcontrollers class in college over a decade ago to run a PWM motor and cleaned it up.
The fan controller just gets a percentage from the server; I need a script on the server to read temperatures and scale the fan speed appropriately. And it's an AI server, so I had it write its own fan control script; it seemed fitting. Deepseek V4 Flash is quite capable of this, with the usual amount of guidance and human interaction you need to get decent software out of a lightweight AI model. 2 We had a funny moment in there when I told it to think of something to heat up the GPUs so I could test the script, and it starts writing a PyTorch script to do matrix multiplication, and I tell it "no, you are running on these cards, you exist in a llama.cpp instance on this computer, just saying anything will load up the cards" and it had a brief existential crisis. Adorable!
The controller itself is on a prototyping breadboard that I just stuffed into the server case. I put a bit of electrical tape on it so it doesn't short against the case.
At this point the box is pretty stable, I've tested the GPUs out, and I'm ready to really dive into how to get it to run fast, which I'll talk about in the next chapter.
Download the shroud's STL file and print it if you need. [back]
Code for arduino and server script available here. [back]
