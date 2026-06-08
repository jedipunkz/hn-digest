---
source: "https://www.cocoawithlove.com/blog/macos-ml-frameworks.html"
hn_url: "https://news.ycombinator.com/item?id=48442089"
title: "Training an LLM in Swift, Part 2: macOS built-in frameworks"
article_title: "Training an LLM in Swift, Part 2: macOS built-in frameworks | Cocoa with Love"
author: "cl3m"
captured_at: "2026-06-08T07:51:13Z"
capture_tool: "hn-digest"
hn_id: 48442089
score: 2
comments: 0
posted_at: "2026-06-08T06:58:35Z"
tags:
  - hacker-news
  - translated
---

# Training an LLM in Swift, Part 2: macOS built-in frameworks

- HN: [48442089](https://news.ycombinator.com/item?id=48442089)
- Source: [www.cocoawithlove.com](https://www.cocoawithlove.com/blog/macos-ml-frameworks.html)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T06:58:35Z

## Translation

タイトル: Swift での LLM のトレーニング、パート 2: macOS 組み込みフレームワーク
記事のタイトル: Swift での LLM のトレーニング、パート 2: macOS 組み込みフレームワーク |愛をこめてココア
説明: macOS での Accelerate、BNNS、CoreML、および MPS の実装について説明します。

記事本文:
愛をこめてココア
メニュー
について
検索
アーカイブ
2026 年 6 月 8 日
マット・ギャラガー著
Swift での LLM のトレーニング、パート 2: macOS 組み込みフレームワーク
macOS での Accelerate、BNNS、CoreML、および MPS の実装について説明します。
この記事では、数値アルゴリズム用に macOS に組み込まれているフレームワークのいくつかを見ていきます。このシリーズのテーマに合わせて、ML モデルをトレーニングできるフレームワークを主に見ていきます。しかし、Accelerate (BLAS)、BNNS、CoreML、MPSGraph など、さまざまなアプローチがたくさんあり、本当の課題は、トレーニングに使用できるかどうかを判断することです。
前回と同様に、例でコードについて説明しますが、LLM の仕組みや用語については実際には説明しません。私がここに来たのは、モデルについてではなく、macOS フレームワークについて話すためです。かなり難解ですが、これは初心者向けのコースではありません。用語についてさらに詳しく知りたい場合は、Andrej Karpathy 氏の「Let's recruit GPT-2 (124M)」を視聴してみてください。そこで彼はすべてを説明しています。
最初に説明する機械学習用の macOS ライブラリは、Accelerate フレームワークです。 Accelerate は実際には、少数の小規模なライブラリのための包括的なフレームワークです。 Accelerate は macOS の重要なライブラリですが、直接使用せずにキャリア全体を費やすこともできるライブラリです。これは Mac OS X 10.2 Jaguar 以来、個別の vecLib、BLAS、LAPACK ライブラリとして存在していましたが、Mac OS X 10.3 Panther では Accelerate に統合されました。大きな猫の名前を覚えていますか？楽しい時間。
一般的な意味で、Accelerate には SIMD ベクトル命令用に最適化された再利用可能なアルゴリズムが含まれています。 Swift では、SIMD ベクトル化に Accelerate フレームワークは厳密には必要ありません (前の記事では、Relaxed.multiplyAdd と Swift の autovectoriz を使用しました)

優れた SIMD ベクトル化を実現するためには必要ですが)、自分のアセンブリを見つめたくない場合には Accelerate を使用するのが非常に役立ちます。
例として、私は最近、単純な PDF パーサー ライブラリにレンダリングを追加し、Accelerate の vImage に基づく次のコードを使用してイメージ マスクを適用しました。
マットにすれば{
ガード var matBuffer = 試してみますか? vImage_Buffer (幅:幅、高さ:高さ、bitsPerPixel:8) else { return nil }
defer { マットバッファ .無料 () }
vImageBufferFill_ARGB8888 (& matBuffer , [ 1 , matt . r , matt . g , matt . b ], vImage_Flags ( kvImageNoFlags ))
vImageAlphaBlend_ARGB8888 (& BaseBuffer 、& MatteBuffer 、& BaseBuffer 、vImage_Flags ( kvImageNoFlags ))
}
これは、以前に使用していた生のピクセル反復よりも約 5 倍高速になりました (デバッグ ビルドでは約 20 倍高速になりました)。
CGContext に描画することでこれを実行できると考えるかもしれません。それは正しいですが、内部で何を使用しているか推測できますか?同じ機能。ここで私がやっているのは、仲介者を排除して、自分自身にもう少し直接的なコントロールを与えることだけです。
前の記事の行列乗算のトピックに戻りますが、Accelerate は BLAS sgemm 実装を提供します。 BLAS は「Basic Linear Algebra Subprograms」の略で、sgemm は「Single precision GENeral Matrix Multiplication」の略です。行列の乗算を他の人に最適化してもらうのは良いことですが、Accelerate BLAS にはもう 1 つの重要な利点があります。それは、前の記事で必要とした醜いハックなしで Apple Silicon AMX ユニットにアクセスできるということです。
それがどのように機能するかを確認するために、前回の Swift の基本的な (単純な) 行列乗算カーネルを考えてみましょう。
static func matmul_forward ( out : inout [ Float ]、 inp : [ Float ]、weight : [ Float ]、bias : [ Float ]?、 B : Int 、 T : Int 、 C : Int 、 OC : I

nt){
0 の b の場合。 .< B {
0 の t の場合。 .< T {
bt = b * T + t とします
0 の o の場合。 .< OC {
var val = バイアス ?[ o ] ?? 0
for i in 0. .< C {
val += inp [ bt * C + i ] * 重み [ o * C + i ]
}
out [ bt * OC + o ] = val
}
}
}
}
BLAS で同じことを行うと次のようになります。
static func matmul_forward ( out : inout [ Float ], inp : [ Float ], Weight : [ Float ],bias : [ Float ]?, B : Int , T : Int , C : Int , OC : Int ) {
cblas_sgemm ( CblasColMajor 、 CblasTrans 、 CblasNoTrans 、 Int32 ( OC )、 Int32 ( B * T )、 Int32 ( C )、 1 、 重み 、 Int32 ( C )、 inp 、 Int32 ( C )、 0 、 & out 、 Int32 ( OC ))
ガード変数バイアス else { return }
出ます。 withUnsafeMutableBufferPointer { outBuffer in
Guard let outBase = outBuffer 。 BaseAddress else { return }
0 の bt の場合。 < ( B * T ) {
cblas_saxpy ( Int32 ( OC ), 1 , &bias , 1 , outBase .advanced ( by : bt * OC ), 1 )
}
}
}
matmul_forward 関数は、追加のバイアス ステップと融合しているだけで、一般的な sgemm 関数とほぼ同じです。
前の記事で必要とした他のすべての最適化を無視し、適用される「Basic Swift」実装の 9 か所で cblas_sgemm を使用するだけで、次の結果が得られます。
この 1 つの変更 (9 か所) により、「Basic Swift」の実装が 144 倍高速になりました。 AMX ユニットを直接使用するという私のハック的な試みよりも 1.2 倍も高速です。 Accelerate BLAS の実装は本当に印象的です。そしてそれはただのスピードだけではありません。実行時の CPU 使用率のグラフは次のとおりです。
AMX (左半分) および BLAS (右半分) モデルのインスツルメント CPU 使用率グラフ
Accelerate フレームワークは、高速かつ非常に低い CPU 使用率で行列乗算を管理します。これを上回る唯一のものは、GPU 上で実行される「Tiled Metal」実装です。
それはそれほど驚くべきことではありません。 GPU は実質的に高性能を目指して作られています

行列計算を行うため、私は非常に多くの GPU コアを搭載した M3 Max で実行しています。
CPU がその GPU スコアを上回ると考えるのは、かなりの狂人です。
BLAS は「Basic Linear Algebra Subprograms」の略であると述べたことを覚えていますか?同じ 1970 年代の命名スキームに基づいて、Accelerate は BNNS (「Basic Neural Network Subprograms」) も提供しています。キャッチー。頭の中で「バンズ」と発音します。
BNNS では数回の反復が行われました。以下のものがあります:
BNNSNNDArrayDescriptor のものを含む「クラシック BNNS API」（非推奨）
BNNS。名前空間 (これも非推奨なので、おそらく準クラシック API)
少なくとも 1 つの古い BNNS API に基づいた ML Compute フレームワーク (非推奨)
BNNSGraph ← これは良いもので、 BNNSTensor 、 BNNSGraph.Builder 、 BNNSGraph.Builder.Tensor 、および BNNSGraph.Context が含まれます
公正な警告: BNNS のドキュメントを検索するのは面倒です。 Xcode のドキュメントには多くの欠陥があり、非推奨ではない BNNS API を見つけようとすると、そのうちのいくつかに遭遇します。
これまでに示したニューラル ネットワークの実装はすべて必須でした。これは、それらが (コンパイラーによる小さな変更を無視して) 再解釈の可能性なしに最初から最後まで実行される一連の操作で構成されていることを意味します (コンピューターは最後まで処理がどこに行くのかわからないため)。
BNNSGraph は、名前が示すように、「グラフ」API (宣言型データフロー API の一種) です。これは、ユーザーが行程全体を宣言し、API が一連の命令を構築することを意味しますが、API は目的地から逆方向にトレースして何が必要かを確認することで、あらゆる種類の最適化を行うことができます。
BNNSGraph を使用する一般的なパターンは、Swift Array<Float> データを BNNSGraph.Builder.Tensor<Float> 値にロードし、解釈することです。

値をさまざまな形状として設定し、さまざまな数学関数を適用して、最終的に最後にさまざまな Array<Float> 値に書き込まれる出力としていくつかの値を宣言します。
これはこれまで私たちが行ってきたこととあまり変わりませんが、for ループを reshape (後続の関数が値を横切る方法を変更するため) と、これまで依存してきた単純な * 、 + 、 exp (または cblas_sgemm ) よりもわずかに大きな関数に置き換えます。
たとえば、BNNSGraph での matmul_forward の実装は次のとおりです。
static func matmul_forward (inp : Tensor 、weight : Tensor 、bias : Tensor 、B : Int 、T : Int 、C : Int 、OC : Int ) -> Tensor {
入力。変形 ( : [ B * T , C ])
。リニア (重み: 重み、バイアス: バイアス)
。変形 ( : [ B , T , OC ])
}
この場合、linear は基本的に転置行列による乗算です (行優先行列の BLAS 例の cblas_sgemm(CblasColMajor, CblasTrans, CblasNoTrans) と同じです)。ただし、この場合、BLAS 実装で厄介なポストステップであったものを含めるためのバイアス パラメーターも含まれています。
しかし、このケースで注目したいのは、両側の reshape コールです。注意すべき重要な点は、このようなグラフ API では、reshape や Gather などのトラバーサル演算子が必ずしも関数に変換されるわけではないということです。 reshape などの操作は、後続の操作に特定の順序で反復する必要がある可能性があることを伝えますが、それ以外の場合は、必ずしも作業を行う必要はありません。さらに、この関数の完全な出力が必要ない場合 (おそらく出力からマスクされているか、出力から削除されている可能性があります)、線形であっても命令型実装で実行されるすべてのことを実行できるわけではない可能性があります。
ここで宣言型 API を最適化できます。

ze: 操作をまとめて実行し、不必要な作業をスキップし、メモリをトラバースする最適な順序を見つけ出すことができます (そのため、前の記事の最後を埋めていたタイリングとスレッドがはるかに少なくなります)。
もう 1 つの注目すべき点は、バイアスを伴う重みの適用を実行するために、linear という名前の関数があることです。これはかなりニューラル ネットワーク固有のタスクです。 BNNSGraph は他の数値処理にも使用できますが、明らかにニューラル ネットワークの利用者に対応しています。
BNNSGraph のattention_forward ブロックは次のとおりです。
static functention_forward (inp : Tensor 、mask : Tensor 、validAttendant : BoolTensor 、B : Int 、T : Int 、C : Int 、NH : Int 、HS : Int 、scale : Float ) -> tentionForwardResult {
q = inp [0. .< B 、0. .< T 、0. .< 1 、0. .< NH 、0. .< HS ] とします。絞り（軸：2）。転置 (軸: [ 0 , 2 , 1 , 3 ])
k = inp [0. .< B 、0. .< T 、1. .< 2 、0. .< NH 、0. .< HS ] とします。絞り（軸：2）。転置 (軸: [ 0 , 2 , 3 , 1 ])
v = inp [0. .< B 、0. .< T 、2. .< 3 、0. .< NH 、0. .< HS ] とします。絞り（軸：2）。転置 (軸: [ 0 , 2 , 1 , 3 ])
preatt = q とします。マトムル（その他：k）・スケール
SavedPreatt = validAttend にします。 select ( preatt 、 0 を Float として)
att = ( preatt + マスク ) とします。ソフトマックス（軸：3）
出す = アット 。マトムル (その他 : v )
。転置 (軸: [ 0 , 2 , 1 , 3 ])
。変形 ( : [ B * T , C ])
return AttendantForwardResult (out : out 、 preatt : SavedPreatt 、 att : att )
}
私が使用してきたattention_forwardの以前の実装をコピーアンドペーストするつもりはありません。それらの長さは約70行です。単一コマンドとしてsqueeze、matmul、select、softmaxを発行できることにより、大幅な簡素化が実現します。
そして、そのパフォーマンスは、CPU からは想像もしていなかったものです。
それが

CPU (AMX を含む) はハイエンド GPU を上回ります (確かに、GPU アルゴリズムは粗く調整されているだけです)。はい、それは単なる CPU です。 Instruments の CoreML プロファイリング ツールでエンジンを実行すると、CPU、GPU、ANE の使用状況を観察し、このエンジンが CPU 単独での GPU アルゴリズムよりも 20% 高速にトレーニングし、推論では 3 倍高速にトレーニングしていることを確認できます。
ただし、LLM のトレーニングという特定のトピックに関しては、BNNSGraph にはいくつかの欠点があります。
それは実際にはトレーニングを目的としたものではありません
実際には LLM を対象としたものではありません
2 番目の点に最初に対処します。squeeze、matmul、select、softmax は優れていますが、LLM を対象としたフレームワークには、scaledDotProductAttend (attention_forward 関数を大幅に簡素化) が含まれるため、依然としてプリミティブから基本ブロックを構築しているように感じられます。
最初の点に対処するために、このようなグラフでは通常、唯一の出力がロジット (出力トークンの予測) であることが期待されるという事実を強調したいと思います。ただし、この場合はさらに多くのことを出力する必要がありました。すべてのレイヤーからすべての単一投影を出力する必要がありました。
ファイルの半分は、グラフに追加される次の出力の海によって占められています。
出力。追加 ( ln1 .out )
出力。追加 ( ln1 . 平均 )
出力。追加 ( ln1 . rstd )
出力。 append ( qkv . reshape ( to : [バッチサイズ , sequenceLength , 3 * C ]))
出力。追加する

[切り捨てられた]

## Original Extract

Covering the Accelerate, BNNS, CoreML and MPS implementations in macOS.

Cocoa with Love
Menu
About
Search
Archive
June 8, 2026
by Matt Gallagher
Training an LLM in Swift, Part 2: macOS built-in frameworks
Covering the Accelerate, BNNS, CoreML and MPS implementations in macOS.
In this article, I’m going to look at some of the frameworks that are built into macOS for numerical algorithms. Fitting with the theme of this series, I’ll mostly be looking at frameworks that can also train an ML model. But there’s a lot of different approaches – Accelerate (BLAS), BNNS, CoreML, MPSGraph – and the real challenge is knowing which one to use – if they are even usable for training.
As with last time, I’ll be talking about code in examples but I won’t really be explaining the mechanics of LLMs or the terminology. I’m really here to talk about the macOS frameworks, not the models. I know, it’s pretty cryptic but this isn’t a beginner’s course. If you want to learn more about the terminology, try watching Andrej Karpathy’s Let’s reproduce GPT-2 (124M) where he goes through everything.
The first macOS library for machine learning that I want to talk about is the Accelerate framework. Accelerate is really an umbrella framework for a handful of smaller libraries. Accelerate is a critical library in macOS but one that you can spend a whole career without directly using. It’s been around since Mac OS X 10.2 Jaguar as separate vecLib, BLAS and LAPACK libraries and then in Mac OS X 10.3 Panther was unified into Accelerate. Remember the big cat names? Fun times.
In a general sense, Accelerate contains reusable algorithms optimized for SIMD vector instructions. In Swift, we don’t strictly need the Accelerate framework for SIMD vectorization (in the previous article, I used Relaxed.multiplyAdd and Swift’s autovectorization to get excellent SIMD vectorization) but it’s still really helpful to use Accelerate when you don’t want to stare at your own assembly.
As an example, I recently added rendering to a simple PDF parser library and used the following code based on Accelerate’s vImage to apply an image mask:
if let matte {
guard var matteBuffer = try ? vImage_Buffer ( width : width , height : height , bitsPerPixel : 8 ) else { return nil }
defer { matteBuffer . free () }
vImageBufferFill_ARGB8888 (& matteBuffer , [ 1 , matte . r , matte . g , matte . b ], vImage_Flags ( kvImageNoFlags ))
vImageAlphaBlend_ARGB8888 (& baseBuffer , & matteBuffer , & baseBuffer , vImage_Flags ( kvImageNoFlags ))
}
which ended up being about 5 times faster than the raw pixel iteration that I was using before (and about 20 times faster in Debug builds).
You might think you could do this by drawing into a CGContext , and you’d be correct but guess what that uses internally? Same functions. All I’m doing here is cutting out the middleman and giving myself a little more direct control.
Getting back to the matrix multiplication topic from the previous article, Accelerate offers us its BLAS sgemm implementation. BLAS stands for “Basic Linear Algebra Subprograms” and sgemm stands for “Single precision GEneral Matrix Multiplication”. Having someone else optimize matrix multiplication is good but Accelerate BLAS offers another key advantage: it lets us access the Apple Silicon AMX unit without the ugly hacks that I needed in the previous article.
To see how it works, let’s consider the basic (naïve) matrix multiplication kernel in Swift from last time:
static func matmul_forward ( out : inout [ Float ], inp : [ Float ], weight : [ Float ], bias : [ Float ]?, B : Int , T : Int , C : Int , OC : Int ) {
for b in 0. .< B {
for t in 0. .< T {
let bt = b * T + t
for o in 0. .< OC {
var val = bias ?[ o ] ?? 0
for i in 0. .< C {
val += inp [ bt * C + i ] * weight [ o * C + i ]
}
out [ bt * OC + o ] = val
}
}
}
}
Doing the same thing with BLAS looks like this:
static func matmul_forward ( out : inout [ Float ], inp : [ Float ], weight : [ Float ], bias : [ Float ]?, B : Int , T : Int , C : Int , OC : Int ) {
cblas_sgemm ( CblasColMajor , CblasTrans , CblasNoTrans , Int32 ( OC ), Int32 ( B * T ), Int32 ( C ), 1 , weight , Int32 ( C ), inp , Int32 ( C ), 0 , & out , Int32 ( OC ))
guard var bias else { return }
out . withUnsafeMutableBufferPointer { outBuffer in
guard let outBase = outBuffer . baseAddress else { return }
for bt in 0. . < ( B * T ) {
cblas_saxpy ( Int32 ( OC ), 1 , & bias , 1 , outBase . advanced ( by : bt * OC ), 1 )
}
}
}
The matmul_forward function is almost exactly the same as a typical sgemm function, just fused with an additional bias step.
Ignoring all the other optimizations we needed in the previous article, just using cblas_sgemm in the 9 places in the “Basic Swift” implementation where it applies, gives us:
This one change (in 9 places) made the “Basic Swift” implementation 144 times faster. It’s even 1.2 times faster than my hacky attempt to use the AMX unit directly. The Accelerate BLAS implementation is really impressive. And it’s not just raw speed. Here’s a graph of the CPU usage running:
Instruments CPU usage graph for AMX (left half) and BLAS (right half) models
The Accelerate framework manages matrix multiplication that is both fast and very low CPU usage. The only thing that beats it is my “Tiled Metal” implementation running on the GPU.
That’s not really surprising. The GPU is practically made for high-performance matrix math and I’m running on an M3 Max which has quite a lot of GPU cores.
You’d have to be crazy to think the CPU would beat that GPU score.
Remember when I said that BLAS stands for “Basic Linear Algebra Subprograms”? Based on the same 1970s naming scheme, Accelerate also offers BNNS or “Basic Neural Network Subprograms”. Catchy. I’m going to pronounce it in my head as “buns”.
BNNS has had a few iterations. There are:
the “Classic BNNS API” which includes BNNSNDArrayDescriptor stuff (deprecated)
the BNNS. namespace (maybe the semi-Classic API because it’s also deprecated)
the ML Compute framework which was based on at least one of the older BNNS APIs (and guess what: deprecated)
BNNSGraph ← this is the good one and includes BNNSTensor , BNNSGraph.Builder , BNNSGraph.Builder.Tensor and BNNSGraph.Context
Fair warning: searching the documentation for BNNS is a mess. Xcode’s documentation has many flaws and trying to find the non-deprecated BNNS APIs runs into a few of them.
All of the neural network implementations I’ve shown so far have been imperative. This means that they’ve been made of a set of operations that are (ignoring minor changes made by the compiler) executed from start to finish, without possibility of reinterpretation (since the computer doesn’t know where it’s going until the end).
BNNSGraph is, as the name implies, a “graph” API (a kind of declarative dataflow API). This means that you declare the entire journey and the API builds a set of instructions, but the API is allowed to make all kinds of optimizations by tracing backwards from the destination to see what is required.
The general pattern for using BNNSGraph is to load your Swift Array<Float> data into BNNSGraph.Builder.Tensor<Float> values, interpret those values as different shapes, apply different mathematical functions, and then declare some values as the output to ultimately be written into different Array<Float> values at the end.
It’s really not very different from what we’ve been doing up to this point but it replaces for loops with reshape (to change how subsequent functions stride through the values) and with some slightly bigger functions than the simple * , + and exp (or even cblas_sgemm ) that we’ve relied upon to this point.
For example, here’s an implementation of matmul_forward in BNNSGraph :
static func matmul_forward ( inp : Tensor , weight : Tensor , bias : Tensor , B : Int , T : Int , C : Int , OC : Int ) -> Tensor {
inp . reshape ( to : [ B * T , C ])
. linear ( weight : weight , bias : bias )
. reshape ( to : [ B , T , OC ])
}
In this case linear is essentially a multiplication by a transposed matrix (same as cblas_sgemm(CblasColMajor, CblasTrans, CblasNoTrans) from the BLAS example was for our row-major matrix) but in this case, it also includes a bias parameter to include what was a messy post-step in our BLAS implementation.
But what I want to draw attention to in this case are the reshape calls on either side. The important point to note is that in a graph API like this, a reshape or a gather or some other traversal operator doesn’t necessarily get translated to a function. Operations like reshape tell subsequent operations that they might need to iterate in a particular order but otherwise, they don’t necessarily do any work. Furthermore, if the full output of this function isn’t needed (maybe it’s causally masked or otherwise removed from the output) then even linear might not do everything it would do in an imperative implementation.
This is where a declarative API can optimize: it can roll operations together, skip unnecessary work, and can work out the best order to traverse memory (so there’s much less of the tiling and threading that filled the end of the previous article).
Another point to notice is that we have a function named linear for performing the application of weights with a bias. That’s a pretty neural-network-specific task. While BNNSGraph could be used for other numerical processing, it’s very clearly catering to a neural network audience.
Here’s our attention_forward block in BNNSGraph :
static func attention_forward ( inp : Tensor , mask : Tensor , validAttention : BoolTensor , B : Int , T : Int , C : Int , NH : Int , HS : Int , scale : Float ) -> AttentionForwardResult {
let q = inp [ 0. .< B , 0. .< T , 0. .< 1 , 0. .< NH , 0. .< HS ]. squeeze ( axis : 2 ). transpose ( axes : [ 0 , 2 , 1 , 3 ])
let k = inp [ 0. .< B , 0. .< T , 1. .< 2 , 0. .< NH , 0. .< HS ]. squeeze ( axis : 2 ). transpose ( axes : [ 0 , 2 , 3 , 1 ])
let v = inp [ 0. .< B , 0. .< T , 2. .< 3 , 0. .< NH , 0. .< HS ]. squeeze ( axis : 2 ). transpose ( axes : [ 0 , 2 , 1 , 3 ])
let preatt = q . matmul ( other : k ) * scale
let savedPreatt = validAttention . select ( preatt , 0 as Float )
let att = ( preatt + mask ). softmax ( axis : 3 )
let out = att . matmul ( other : v )
. transpose ( axes : [ 0 , 2 , 1 , 3 ])
. reshape ( to : [ B * T , C ])
return AttentionForwardResult ( out : out , preatt : savedPreatt , att : att )
}
I’m not even going to copy-paste the previous implementations of attention_forward that I’ve been using: they’re around 70 lines long. The ability to issue squeeze , matmul , select and softmax as single commands is a huge simplicity gain.
And the performance is something I never expected to see from the CPU:
That’s the CPU (including AMX) outperforming a high-end GPU (admittedly the GPU algorithm was only coarsely tuned). And yes, it is just the CPU. If you run the engines in Instruments’ CoreML profiling tool, you can observe CPU, GPU and ANE usage and confirm that this engine is training 20% faster than a GPU algorithm on the CPU alone and 3 times faster for inference.
However, there are some downsides to BNNSGraph on the specific topic of training LLMs.
it’s really not intended for training
It’s not really intended for LLMs
Addressing the second point, first: squeeze , matmul , select and softmax are nice but LLM-targeted frameworks would have scaledDotProductAttention (dramatically simplifying our attention_forward function) so it still feels like we’re building fundamental blocks from primitives.
Addressing the first point, I’d like to highlight the fact that usually for a graph like this, you’d expect the only outputs to be the logits (the predictions for output tokens). However, in this case, I’ve had to output a lot more: I’ve had to output every single projection from every layer.
Half the file is taken up by a sea of these outputs appended to the graph:
outputs . append ( ln1 . out )
outputs . append ( ln1 . mean )
outputs . append ( ln1 . rstd )
outputs . append ( qkv . reshape ( to : [ batchSize , sequenceLength , 3 * C ]))
outputs . append

[truncated]
