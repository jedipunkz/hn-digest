---
source: "https://yeokhengmeng.com/2025/04/llama2-llm-on-dos/"
hn_url: "https://news.ycombinator.com/item?id=48840660"
title: "Llama 2 LLM on DOS (2025)"
article_title: "Llama 2 LLM on DOS - YKM's Corner on the Web"
author: "userbinator"
captured_at: "2026-07-09T03:33:18Z"
capture_tool: "hn-digest"
hn_id: 48840660
score: 1
comments: 0
posted_at: "2026-07-09T03:31:33Z"
tags:
  - hacker-news
  - translated
---

# Llama 2 LLM on DOS (2025)

- HN: [48840660](https://news.ycombinator.com/item?id=48840660)
- Source: [yeokhengmeng.com](https://yeokhengmeng.com/2025/04/llama2-llm-on-dos/)
- Score: 1
- Comments: 0
- Posted: 2026-07-09T03:31:33Z

## Translation

タイトル: DOS 上の Llama 2 LLM (2025)
記事タイトル: DOS 上の Llama 2 LLM - Web 上の YKM のコーナー
説明: DOS を実行しているビンテージ PC 上でローカルのラージ言語モデル (LLM) を実行したいと考えたことはありますか?今ならそれが可能です！

記事本文:
メーカー、プログラマー、プライベート パイロット、レトロコンピューティング愛好家
DOS を実行しているビンテージ PC 上でローカルのラージ言語モデル (LLM) を実行したいと考えたことはありますか?今ならそれが可能です！
このビデオでは、2 台のビンテージ DOS マシン上で実行されている DOS Llama 2 LLM クライアントを示しています。
Thinkpad T42 (2004) および Toshiba Satellite 315CDT (1996) で実行。 CPUはそれぞれPentium M 735 1.7GhzとPentium MMX 200Mhzです。
Core i5-10310U 1.7Ghz を搭載した Thinkpad X13 Gen 1 (2020) 上の FreeDOS 1.4 で実行。
すべてがオープンソースであり、実行可能ファイルはここから入手できます: https://github.com/yeokm1/dosllam2
2 年前、私は DOS ChatGPT クライアントを作成しましたが、コミュニティ内の多くのレトロコンピューティング愛好家が他のビンテージ プラットフォーム用に同様のクライアントを作成しました。それらはすべて、リモート サービスへの依存という同様の問題を抱えていました。ローカルで実行する場合はどうでしょうか?
従来の通念では、LLM をローカルで実行するには、高性能仕様のコンピューター、特に VRAM を大量に搭載した GPU が必要であると考えられています。しかし、これは本当に本当なのでしょうか？
2023 年、有名な機械学習の専門家である Andrej Karpathy は、たった 1 つの C ファイルで Meta の Llama 2 モデルの推論を可能にするオープンソースの llama2.c プロジェクトをリリースしました。その後、他の開発者が彼のプロジェクトを移植し、Windows 98 や Windows 98 などのビンテージ プラットフォームを含む多くのプラットフォームで実行できるようにしました。
パワーブックG4。誰かがそれをマイクロコントローラー上で実行できるように移植したことさえあります。
そこで私は、DOS のようなさらに古いプラットフォームでも同じことができるのではないかと考えました。
デモ ビデオとスクリーンショットに基づいて、これを実行できたことは明らかなので、このブログ投稿でその方法を示します。
llama2.c は、Llama-2 アーキテクチャの fp32 モデルのみを推論するように設計された単一の C ファイルで記述されています。
このリポジトリは依然として効率を重視していますが、シンプルさ、読みやすさ、移植性を犠牲にすることはありません。
https://github.com/karpathy/llama2

.c
より小さなモデルのテストを容易にするために、Karpathy は、サイズが 260K、15M、42M、110M のみの TinyStories データセットでいくつかの小さなモデルをトレーニングしました。これは、リソース制約システムで基本レベルの LLM 機能を有効にするためです。
llama2.c は移植性を念頭に置いて書かれていますが、コードベースをビンテージ システムで動作させるにはまだいくつかの課題があります。
私が選択したコンパイラは Open Watcom v2 (OWC) です。 DOS ChatGPT クライアントにも使用しており、慣れていたのでこれを使用しました。
OWC は GUI IDE を提供しますが、プロジェクトを比較的軽量にして管理しやすくするために、私は Makefile を使用することを好みます。
これは、プロジェクトをコンパイルするために使用した OWC Makefile です。
ターゲット = ドスラム2
OBJS = dosllam2.obj
CFLAGS = -za99
LDFLAGS = システム dos32a 名 $( TARGET )
すべて: クリーン $( TARGET ) .exe
$( TARGET ) .obj: dosllam2.c
wcc386 $( CFLAGS ) dosllam2.c
$( TARGET ) .exe: $( OBJS )
wlink $( LDFLAGS ) ファイル $( OBJS )
コンパイラ ツールチェーンは、 wcc386 と wlink を使用します。 -za99 を使用して C99 サポートを有効にしました。有効にしないと C90 に戻ります。
dos32a フラグには特別な注意が払われます。ほとんどの DOS プログラムでは、通常、次の 2 つの典型的なモードでコンパイルできます。
最初の IBM PC で使用された Intel 8088 と互換性のある従来の 16 ビット アーキテクチャ
Intel 386 CPU 以降の 32 ビット DOS エクステンダーのサポート
llama2.c コードベースは少なくとも 32 ビットの最新システム向けに書かれているため、後者のオプションを選択する必要があります。最小の LLM であってもメモリ要件には、少なくとも 32 ビット システムも必要です。
DOS 時代に使用されていた一般的な 32 ビット DOS エクステンダーの 1 つは DOS/4G です。このエクステンダを使用すると、CPU を 32 ビット保護モードに切り替えることで、プログラムが 640KB の従来メモリおよび上位メモリ領域を超える拡張メモリにアクセスできるようになります。非常に一般的な使用例の 1 つは、ゴブを必要とするゲームです。

記憶の。エクステンダーを使用した有名な例としては、Doom や Descent があります。
OWC は、エクステンダーの選択肢として DOS/4G をサポートしているほか、2006 年にリリースされた Narech Koumar による別の新しいオープンソース DOS/32 もサポートしています。
このエクステンダーは、ランタイムとしてプログラムに埋め込むことも、個別に実行することもできます。いずれにせよ、ほとんどの DOS エクステンダーは上記のような起動バナーを表示します。
元の llama2.c コードベースは依然として、適度に最新の機能を備えた C コンパイラを前提としています。 OWC プロジェクトはまだ維持されていますが、コードをコンパイルするには少し移植作業が必要でした。
推論には浮動小数点演算を多用する必要がありますが、OWC はこれを十分にサポートしていません。
#define sqrtf(x) ((float)sqrt((double)(x)))
#define powf(x, y) ((float)pow((double)(x), (double)(y)))
#define cosf(x) ((float)cos((double)(x)))
#define sinf(x) ((float)sin((double)(x)))
#define expf(x) ((float)exp((double)(x)))
そのため、既存の double 関数を使用していくつかのマクロを定義する必要がありました。
最新のコンパイラは、ストレージ内のファイルをメモリにマップできる mmap または同様の関数をサポートしています。ファイルの内容はすべて一度にメモリに読み込まれるのではなく、遅延読み込みされます。つまり、必要なときに必要な部分だけが読み取られるため、特に大きなファイルの使用を高速化できます。
OWC はこれをサポートしていないため、LLM ファイル全体をメモリにロードするには、すべてのメモリ マッピング呼び出しを置き換える必要があります。これにより、初期ロード時間が増加するという影響があります。
私は独自の mmap バージョンを作成することを検討しました。ただし、LLM の使用例では、最終的にはほとんど (すべてではないにしても) 重みが使用されるため、ファイル全体のロードが延期されるだけです。これまでのところ、物事を単純にして、最初にファイル全体を一度にロードする方が簡単です。
推論速度を測定するには、プログラムでタイミング API を使用する必要があります。
長い時間ミリ秒 () {
セント

ルート時間仕様時間。
Clock_gettime (CLOCK_REALTIME, & 時間);
time.tv_sec * 1000 + time.tv_nsec / 1000000 を返します。
}
これらの API は利用できないため、他の API に変更する必要があります。
長い時間ミリ秒 () {
return ( クロック () * 1000L ) / CLOCKS_PER_SEC;
}
OWC は、プロセッサのティック時間をカウントする Clock() を提供します。秒は、内部マクロ CLOCKS_PER_SEC で除算することで取得できます。
DOS には 8.3 のファイル名制限があります。これは、ファイルの読み取りに使用される C コードにも影響します。
// 前へ
char * tokenizer_path = "tokenizer.bin" ;
//新規
char * tokenizer_path = "tokenize.bin" ;
元のプログラムは、8.3 ルールに準拠していないハードコーディングされたパスを含むトークナイザー ファイルをロードします。ファイル名を 1 文字短縮する解決策は簡単ですが、最初はなぜこのロードが私のビンテージ システムで機能しないのか疑問に思い、トラブルシューティングに時間がかかりました。
30 年前の PC から最新のシステムまでを使用してベンチマークをいくつか実行しました。
データは 1 秒あたりのトークン単位です。すべてのオペレーティング システムは MS-DOS 6.22 であり、メモリ マネージャーはロードされていません。
予想通り、システムが最新であればあるほど、推論速度は速くなります。
興味深い観察は、最新の Ryzen CPU と 128GB RAM を搭載した私のデスクトップでは、メモリ割り当てエラーにより、より大きなモデルをロードできないことです。おそらく、DOS エクステンダーまたは AMD CPU/マザーボードのプロテクト モード実装におけるある種のオーバーフロー/バグが考えられます。
Thinkpad T42 のパフォーマンスはいくつかの新しいラップトップを上回り、16 年新しい Thinkpad X13G1 と同等です。
この表にある私のビンテージ PC の詳細については、ここを参照してください。
私が行った移植作業、テストと文書化は週末の簡単な仕事です。努力のほとんどは、このプロジェクトの元の開発者である Andrej Karpathy に帰する必要があります。このプロジェクトなしでは、このプロジェクトは実現しなかったでしょう。

可能になった。
それでも、ビンテージシステムで今でも何ができるかには驚かされます。
Apple II に Swift を導入
Creative Midi Blaster MB-10 シンセサイザーのレビューと分解
新たに設立された NUS CEG 修士課程卒業生の感想
ヴィンテージ風のモダンな 2024 PC (Ryzen 5 7600 & GeForce 4060 Ti 搭載)
ハック 27
個人18
レトロコンピューティング 17
シンガポール 15
航空 14
分解 11
学校10
書評 7
製品レビュー 5
ハッカソン 4
政治 4
管理者3
修理-コピティアム3
キーボード2
クイズ2
ドヴォルザーク 1
フライトシム 1
ハードウェア1
予防接種1
© 2026
YKMのWebコーナー
。
テーマ: ヒューゴ・フューチャー・インパーフェクト・スリム
HTML5 UP ポート |ヒューゴが提供する

## Original Extract

Ever thought of running a local Large Language Model (LLM) on a vintage PC running DOS? Now you can!

Maker, Coder, Private Pilot, Retrocomputing Enthusiast
Ever thought of running a local Large Language Model (LLM) on a vintage PC running DOS? Now you can!
This video shows DOS Llama 2 LLM client running on 2 vintage DOS machines.
Running on Thinkpad T42 (2004) and Toshiba Satellite 315CDT (1996). Their CPUs are Pentium M 735 1.7 Ghz and Pentium MMX 200Mhz respectively.
Running on FreeDOS 1.4 on Thinkpad X13 Gen 1 (2020) with Core i5-10310U 1.7Ghz.
Everything is open sourced with executable available here: https://github.com/yeokm1/dosllam2
2 years ago I wrote a DOS ChatGPT client and many retrocomputing enthusiasts in the community wrote similar clients for other vintage platforms. They all had a similar issue, dependency on some remote service. What about running things locally?
Conventional wisdom states that running LLMs locally will require computers with high performance specifications especially GPUs with lots of VRAM. But is this actually true?
In 2023, famous machine learning expert Andrej Karpathy released his open-source llama2.c project which allows inferencing of Meta’s Llama 2 models with just a single C-file. Other developers then ported his project to run on many platforms including some vintage ones like Windows 98 and
Powerbook G4 . Someone even ported it to run on a microcontroller .
This then inspired me to think, could the same thing be done for an even older platform like DOS?
Based on the demo video and screenshots, it is obvious I managed to do this so this blog post will show how I did it.
llama2.c is written in a single C-file designed to inference only fp32 models of the Llama-2 architecture.
This repo still cares about efficiency, but not at the cost of simplicity, readability or portability.
https://github.com/karpathy/llama2.c
For ease of testing of smaller models, Karpathy trained several small models on the TinyStories dataset that only have sizes of 260K, 15M, 42M and 110M. This is to enable some basic level of LLM functionality on resource-constraint systems.
llama2.c although written for portability in mind, still has some challenges when it comes to making the codebase work for vintage systems.
The compiler I selected is Open Watcom v2 (OWC) . I used this as I was familiar with it having used it for my DOS ChatGPT client as well.
OWC provides a GUI IDE but I prefer to use a Makefile to keep the project relatively lightweight and easier to manage.
This is the OWC Makefile I used to compile my project:
TARGET = dosllam2
OBJS = dosllam2.obj
CFLAGS = -za99
LDFLAGS = SYSTEM dos32a NAME $( TARGET )
all: clean $( TARGET ) .exe
$( TARGET ) .obj: dosllam2.c
wcc386 $( CFLAGS ) dosllam2.c
$( TARGET ) .exe: $( OBJS )
wlink $( LDFLAGS ) FILE $( OBJS )
The compiler toolchain uses wcc386 and wlink . I used -za99 to enable C99 support otherwise it will revert to C90.
Special attention is given to the dos32a flag. For most DOS programs, you generally can compile for 2 typical modes:
Traditional 16-bit architecture compatible with the Intel 8088 used in the first IBM PC
32-bit DOS extender support for Intel 386 CPUs or newer
Since the llama2.c codebase is written for modern systems of at least 32-bit, I must go for the latter option. The memory requirement even for the tiniest of LLMs demands at least a 32-bit system too.
One common 32-bit DOS extender used in the DOS-era is DOS/4G . This extender allows programs to access extended memory above the 640KB conventional memory and Upper Memory Area by switching the CPU to 32-bit protected mode. One very common use case is that of games which required gobs of memory. Famous examples using extenders are Doom and Descent.
OWC supports DOS/4G as a choice of extender as well as another newer and open-source DOS/32 by Narech Koumar released in 2006.
This extender can be embedded into your program as a runtime or run separately. Either way, most DOS-extenders will show a startup banner like the above.
The original llama2.c codebase still assumes a reasonably modern-featured C compiler. Although the OWC project is still being maintained, there was a bit of porting effort required for me to get the code to compile.
The inferencing requires heavy use of floating point operations which OWC does not support well.
#define sqrtf(x) ((float)sqrt((double)(x)))
#define powf(x, y) ((float)pow((double)(x), (double)(y)))
#define cosf(x) ((float)cos((double)(x)))
#define sinf(x) ((float)sin((double)(x)))
#define expf(x) ((float)exp((double)(x)))
So I had to define some macros using existing double functions.
Modern compilers support the mmap or similar functions which allow files in storage to be mapped into memory. The contents of the file are not all read into memory at once but are lazy-loaded i.e only necessary portions are read when required helping to speed up the use of especially large files.
OWC does not support this thus I have to replace all memory-mapping calls to load the entire LLM file into memory. This has the effect of increasing initial load times.
I did consider writing my own version of mmap . However given the LLM use case where pretty much most (if not all) the weights will be used eventually, it is just postponing loading the entire file. So far easier to just keep things simple and load the entire file all at once initially.
To measure inference speed, the program requires use of timing APIs:
long time_in_ms () {
struct timespec time;
clock_gettime (CLOCK_REALTIME, & time);
return time.tv_sec * 1000 + time.tv_nsec / 1000000 ;
}
Those APIs are not available thus I have to change to something els.
long time_in_ms () {
return ( clock () * 1000L ) / CLOCKS_PER_SEC;
}
OWC provides the clock() which counts processor tick time. Seconds can be obtained by dividing by the internal macro CLOCKS_PER_SEC .
DOS has a 8.3 filename limit. This also impacts the C code used to read the files.
// Previous
char * tokenizer_path = "tokenizer.bin" ;
//New
char * tokenizer_path = "tokenize.bin" ;
The original program loads a tokenizer file containing a hard-coded path which does not adhere to the 8.3 rule. Although the solution to shorten the filename by one character is simple, it caused me some troubleshooting time initially as I wondered why this loading did not work on my vintage system.
I did some benchmarks using PCs of 3-decades ago all the way to modern systems.
Data is in tokens per second. All the operating systems are in MS-DOS 6.22 with no memory manager loaded.
As expected, the more modern the system, the faster the inference speed.
An interesting observation is that my desktop with a modern Ryzen CPU and 128GB RAM cannot load the larger models due to memory allocation error. Perhaps some kind of overflow/bug in the DOS extender or protected mode implementation of the AMD CPU/motherboard.
The performance of the Thinkpad T42 beats out some newer laptops and is similar to my Thinkpad X13G1 which is 16 years newer.
More details about my vintage PCs in that table can be found here .
The porting effort I made, testing and documenting this is a simple weekend job. Most of the efforts have to be attributed to the original developer of this project Andrej Karpathy without which this would not have been possible.
Still, I’m amazed what can still be accomplished with vintage systems.
Bringing Swift to the Apple II
Review and Teardown of Creative Midi Blaster MB-10 synthesizer
Reflections of a newly minted NUS CEG Master graduate
Modern 2024 PC with a vintage twist (ft. Ryzen 5 7600 & GeForce 4060 Ti)
hacks 27
personal 18
retrocomputing 17
singapore 15
aviation 14
teardowns 11
school 10
book-reviews 7
product-reviews 5
hackathon 4
politics 4
admin 3
repair-kopitiam 3
keyboards 2
quiz 2
dvorak 1
flight-sim 1
hardware 1
vaccination 1
© 2026
YKM's Corner on the Web
.
Theme: Hugo Future Imperfect Slim
A HTML5 UP port | Powered by Hugo
