---
source: "https://internals-for-interns.com/posts/go-ai-inference-what-is-inference/"
hn_url: "https://news.ycombinator.com/item?id=48977228"
title: "Understanding Go AI Inference: What Is Inference?"
article_title: "What is Inference? | Internals for Interns"
author: "valyala"
captured_at: "2026-07-20T12:09:20Z"
capture_tool: "hn-digest"
hn_id: 48977228
score: 1
comments: 0
posted_at: "2026-07-20T11:19:53Z"
tags:
  - hacker-news
  - translated
---

# Understanding Go AI Inference: What Is Inference?

- HN: [48977228](https://news.ycombinator.com/item?id=48977228)
- Source: [internals-for-interns.com](https://internals-for-interns.com/posts/go-ai-inference-what-is-inference/)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T11:19:53Z

## Translation

タイトル: 囲碁 AI 推論を理解する: 推論とは?
記事タイトル: 推論とは何ですか? |インターン向け社内向け
説明: 新しいシリーズへようこそ!今日のほとんどの開発者にとって、大規模な言語モデルの使用は 1 つのことを意味します。それは、他人のコンピューターへの HTTP 呼び出しです。 API にプロンプ​​トを送信すると、トークンが戻ってきます。その間のすべては誰かの魔法です。
しかし、私がもっと興味深いと思ったのは、次のとおりです。
[切り捨てられた]

記事本文:
インターン向け社内向け
私について
シリーズ
本
エス
Go AI 推論を理解する: 推論とは何ですか?
X（ツイッター）
リンクトイン
フェイスブック
リンクをコピー
コピーしました！ヘスス・エスピノ著
•
2026 年 7 月 20 日
•
15 分で読めます #Go
#AI
#推論
#LLM 新しいシリーズへようこそ!今日のほとんどの開発者にとって、大規模な言語モデルの使用は 1 つのことを意味します。それは、他人のコンピューターへの HTTP 呼び出しです。 API にプロンプ​​トを送信すると、トークンが戻ってきます。その間のすべては誰かの魔法です。
しかし、私がもっと興味深いと思うのは、これらのモデルをローカルで、独自のプロセス内で、独自のハードウェア上で実行できることです。また、Go 開発者であれば、Go から直接実行できます。現在、これを本当に快適にしている 2 つのプロジェクトがあります。 Yzma では、Go が llama.cpp ライブラリを直接呼び出すことができます (cgo は使用せず、このシリーズの後半で独自の記事が提供されます)。もう 1 つは、高レベルの OpenAI-API 感覚の SDK とモデル サーバーを Yzma 上に構築する Kronk です。そして、両方の下には、通常のハードウェア上で LLM を実行できるようにする C/C++ 推論エンジンである llama.cpp があります。
このシリーズでは、そのスタック全体を内部から理解していきます。つまり、llama.cpp が「モデルを実行する」ときに実際に何をするのか、Yzma がどのように Go からそれを呼び出すことができるのか、そして Kronk がそれらすべてを実際に運用環境で必要なエンジンにどのように変換するのかを説明します。
この記事の目的は、ニューラル ネットワーク自体の内部で何が起こるかではなく、モデルの仕組み (モデル ファイルがどのように保存、ロード、実行されるか) を理解することです。これらのレイヤー内の計算はそれ自体が独立したトピックなので、ここでは扱いません。
llama.cpp の詳細は、2026 年 7 月時点のマスター (コミット ad8d8219 ) に対するものです。llama.cpp は高速に動作し、安定版リリースを行わないため、バージョンの代わりにコミットを固定しています。そしてこれは意図的に

概念ツアー: C++ コードの詳細には立ち入りません。また、Go レイヤーには完全に関与しません。Yzma と Kronk が独自の記事を入手します。
範囲を設定したので、詳しく見ていきましょう。
この言葉を分かりやすく理解することから始めましょう。言語モデルには 2 つのまったく異なる性質があります。
最初の人生はトレーニングです。モデルにテラバイトのテキストを表示し、最適化プロセスによって数十億の数値 (重み (パラメーター とも呼ばれます)) が徐々に調整され、モデルが 1 つの狭いタスク (テキストのシーケンスが与えられて、次に何が来るかを予測する) がうまくできるようになるまで続きます。トレーニングには数か月かかり、データセンターにも時間がかかります。このシリーズではトレーニングについてはまったく話していません。
第 2 の人生は推論です。トレーニングは終了し、ウェイトは凍結され、あとはそれらを使用するだけです。最良のメンタル モデルは単純な純粋な関数です。これは 2 つの入力 (一連のテキストと数十億の固定重み) を受け取り、1 つの出力 (次のトークンの予測) を返します。 nextToken(input,weights) のようなもの。それでおしまい。それがすべての秘密です。学習は行われず、何も更新されません。重みは読み取り専用の引数であるため、モデル ファイルは 100 万回目の呼び出しでも最初の呼び出しとまったく同じように賢くなります。
これらの重みが実際にどのくらいのメモリを占有するのか、量子化によってどのように重みが縮小されるのか、モデルをローカルで実行するために必要な VRAM を通常重みが占める理由について興味がある場合は、Kronk の VRAM 計算機に関する記事がおすすめです。
したがって、「モデル」をダウンロードするとき、ダウンロードするものは、基本的に、凍結された数値 (重み) とその使用方法の説明の巨大な山です。モデルを「実行する」ということは、入力とそれらの数値の間で非常に大量の演算を行うことを意味します。不審に手を振っているように聞こえるかもしれませんが、心配しないでください。ウェイトについてはわかりやすく説明します。

記事の後半で、実際のモデル ファイルを開くときに説明します。今のところ、押さえておく必要があるのは、モデルとはすでに誰かが計算した膨大な数値の山であるということだけです。
トークンについてはすでに何度か言及しましたが、トークンとは正確には何でしょうか?見てみましょう。
演算を行う前に問題が 1 つあります。それは、モデルはテキストを読み取れないということです。これらはトークン (固定語彙からの整数 ID) を使用して機能します。各 ID はテキストの塊 (単語、単語の一部、句読点) に対応します。
「フランスの首都は」というプロンプトは、[791, 6864, 315, 9822, 374] のように、5 つまたは 6 つのトークン ID になる可能性があります (正確な ID はモデルの語彙に完全に依存します)。この時点から、モデルは二度と文字を見ることはありません。下流の処理はすべてこれらの整数に関して行われます。
したがって、モデルは整数のリストを受け取ります。それは実際に彼らにどのような影響を与えるのでしょうか？
本質を取り除いた計算の形を次に示します。前もって重要な点が 1 つあります。モデルは入力を単語ごとに別々のラウンドで処理しません。トークンのシーケンス全体が 1 回のパスで一緒にネットワークを通過し、1 つの予測が生成されます。この図を使って説明してみましょう。
これを左から右に読んで、各トークン ID をエンベディング (モデルが使用できる形式でトークンを表す数値のベクトル) に変換することから始めます (エンベディングをよりよく理解したい場合は、Weaviate に優れた入門書があります)。次に、これらすべてのベクトルがまとめてモデルに入力され、大量の計算が実行され、最終的に各トークンが次に出現する可能性を示すスコアのリスト (語彙内のトークンごとに 1 つ) が生成されます。そしてそのリストから「パリ」を 1 つ選びます。
その最終的な選択にはサンプリングという名前があり、私たちは協力します

実際にどのように機能するかは後で説明します。しかし、最初に重要なことに注目してください。このすべての作業で生成されたトークンは 1 つだけです。次の単語を予測するだけのマシンから文全体を取得するにはどうすればよいでしょうか?
もう一度やることで。そしてまた。これが自己回帰の部分であり、LLM の会話を活発にするものです。この図でそれがどのように機能するかを見てみましょう。
矢印に従ってください。 「フランスの首都は」というシーケンスをモデルに入力し、1 回のパスを実行して、トークン「パリ」を選択します。単一の予測を完全な答えに変えるトリックは次のとおりです。停止する代わりに、シーケンスに「パリ」を追加し、今度は「フランスの首都はパリ」でモデルを再度実行します。次のトークンは「.」である可能性があり、それも追加し、モデルが特別な世代終了トークン (「終わった」と言う言い方) を選択するか、長さの制限に達して自分たちで停止するまで、1 周につき 1 つの新しいトークンを追加します。
あなたがこれまで読んだすべての LLM 回答のすべての単語は、このループによって一度に 1 つのトークンが生成されました。生成は 1 つの大きな計算ではなく、for ループ内での同じ次のトークンの予測です。
理論についてはこれくらいです。そのループが実際に実行されるのを見る前に、すべては私たちが検討してきた 1 つのこと、つまり凍結されたウェイトにかかっています。それらはどこから来たのでしょうか?また、「モデル」は実際にはディスク上ではどのように見えるのでしょうか?箱を開ける時が来ました。
GGUF: モデル全体を 1 つのファイルにまとめたもの
これらの重みをわかりやすく説明しましょう。凍結された数値の山は、形のないヒープではありません。テンソルと呼ばれる名前付き配列に編成されています (テンソルは単なる n 次元配列であり、より多くの次元に一般化された行列です)。たとえば、ニューラル ネットワークの単一層は通常、複数のテンソルで構成されます。 GGUF (GGML Universal File) は、テンソルのバッグを保存する方法です

ディスク上にあり、その設計全体は 1 つの目標、つまり単一の自己記述ファイルに従っています。 1 つの .gguf には、重み、アーキテクチャ、ハイパーパラメータ、トークナイザー全体、さらにはチャット テンプレートなど、モデルの実行に必要なすべてが保持されており、サイドカー ファイルをダウンロードする必要はありません。 1 つのファイルをコピーし、モデルを実行します。 (非常に大きなモデルは、実用性を考慮していくつかの .gguf シャードに分割されることがありますが、原則は当てはまります。)
では、これらのファイルの 1 つは、実際にはどのような中身になっているのでしょうか?この形式は、ggml の GGUF ヘッダーのすぐ上部に文書化されています。tensor ライブラリ llama.cpp が構築されており、ファイル形式自体はそこに存在します ( ggml/include/gguf.h )。ただし、生の仕様を貼り付けるのではなく、描画してみましょう。
上から下まで読んでみましょう。先頭には小さなヘッダー、つまりマジック バイトの GGUF (プログラムが「はい、これは本当に GGUF ファイルである」を確認できる)、バージョン番号、および予想されるテンソルの数とメタデータ エントリの数を読者に伝える 2 つのカウントがあります。その後のすべては 3 つのセクションに分かれており、このセクションの残りの部分は実際には、メタデータ、テンソル記述子、およびデータ BLOB の 3 つのボックスの単なるガイド付きツアーです。一つずつ見ていきましょう。
メタデータ: それ自体を説明するモデル
それが図のボックス 1 です。メタデータは、型指定されたキーと値のペアのフラット リストであり、値は単純なスカラー、文字列、同種の配列です。キーは名前空間付きの文字列であり、いくつかの名前空間 (図で確認できるもの) があることに注意してください。
general.* — アイデンティティ: general.architecture は、これがどの種類のモデルであるかを示します ( llama 、 qwen2 、 gemma など)。この 1 つのキーによって、llama.cpp が推論機構をセットアップするためにどのグラフ構築コードを使用するかが決まります。
<arch>.* — アーキテクチャ名 ll が接頭辞として付くハイパーパラメータ

ama.embedding_length (トークン ベクトルのサイズ)、llama.block_count (レイヤーの数)、llama.context_length (ウィンドウに収まるトークンの数) など…
tokenizer.* — トークナイザー (テキストをトークンに変換する部分) 全体がモデル ファイル内に存在します。 tokenizer.ggml.tokens は文字列配列としての完全なボキャブラリであり、トークナイザーの構成パラメータや、チャットの会話をフラットなプロンプト文字列に変換するテンプレートである tokenizer.chat_template も含まれます。
llama.cpp がファイルを開くと、最初に general.architecture が読み取られ、他のすべてがそこから続きます。
それがモデルの正体です。では、実際の数字がどこにあるのか見てみましょう。
BLOB: 重みをメモリに取り込む
次に、図のボックス 2 と 3 について説明します。これらはペアとして機能します。ボックス 2 はテンソル記述子です。テンソルごとに 1 行にその名前、形状、タイプが含まれますが、データはありません。代わりに、各行には offset が含まれており、それが図の赤い矢印です。オフセットは、そのテンソルのバイトが実際にボックス 3 (データ blob) 内に存在する場所を指します。
興味深いのは、これらのバイトがどのようにしてメモリに取り込まれるかということです。ローダー ( l​​lama_model_loader 、 src/llama-model-loader.cpp ) は最初にヘッダー、メタデータ、テンソル記述子のみを解析します。テンソル データはまったく読み取られません。何がどこに行くのかを知るにはこれで十分です。一部のレイヤーは CPU に残り、他のレイヤーは GPU に移動し、各宛先が独自の方法でバイトを処理します。 CPU テンソルの場合、 llama.cpp はデフォルトでギガバイトの重みを新たに割り当てられたバッファに read() しません。 mmap を使用してファイルをプロセスのアドレス空間にマップし、各テンソルを直接mapping_address + tensor_offset にポイントします。 GPU テンソルの場合、そのトリックは機能しません。VRAM はまったく別のメモリです。そのため、バイトは実際にコピーされます。

編集: ロード時にファイルから GPU バッファに 1 回アップロードされます。
これで、ファイルのテンソルがメモリ内の適切な場所に配置されました。何かを実行する時間です。
KV キャッシュ: 生成が高速な理由
すでに持っているもの、つまりモデルのデータ、つまり決して変更されない部分から始めましょう。これをロードすると、llama_model が渡されます。これは、メモリ内にあるファイル自体です。重み、トークナイザー、その他すべての読み取り専用のものです。何も変更されないため、共有することができます。ロードされた 1 つのモデル、多くの会話、およびそれらの数ギガバイトの重みは 1 回だけメモリに保存されます。
ここで動的データが必要になります。これが llama_context の目的です。これは 1 つの会話を表し、すべての変化する状態が存在する場所、つまり計算のための作業スペース、そして何よりも KV キャッシュです。モデルを読み込む場合と比較して、コンテキストの作成は低コストであるため、会話ごとにコンテキストを 1 つ起動します。
では、この KV キャッシュとは何でしょうか?ループの各ステップでモデルに非常に長いシーケンスを渡しますが、パスごとに最初から再処理するのは非常に無駄です。これを回避するのが KV キャッシュの仕事です。
重要なのは、モデルがトークンの処理中に、その後変更されることのないいくつかの中間値を計算することです。したがって、llama.cpp はそれらを再計算する代わりに、KV キャッシュに保存して再利用します。新しいトークンはそれぞれ独自の作業を実行し、残りをメモリから読み取るだけで済みます。そのため、テキストの長さに関係なく、生成されたすべてのトークンのコストはほぼ同じになります。

[切り捨てられた]

## Original Extract

Welcome to a new series! For most developers today, using a large language model means one thing: an HTTP call to somebody else’s computer. You send a prompt to an API, tokens come back, and everything in between is somebody else’s magic.
But here’s what I find much more interesting: you can run the
[truncated]

Internals for Interns es
About Me
Series
Books
es
Understanding Go AI Inference: What is Inference?
X (Twitter)
LinkedIn
Facebook
Copy link
Copied! By Jesús Espino
•
July 20, 2026
•
15 min read #Go
#AI
#Inference
#LLM Welcome to a new series! For most developers today, using a large language model means one thing: an HTTP call to somebody else’s computer. You send a prompt to an API, tokens come back, and everything in between is somebody else’s magic.
But here’s what I find much more interesting: you can run these models locally , inside your own process, on your own hardware — and if you’re a Go developer, you can do it directly from Go. Two projects make this genuinely pleasant today: Yzma , which lets Go call the llama.cpp libraries directly (without cgo — that gets its own article later in the series), and Kronk , which builds a high-level, OpenAI-API-feeling SDK and model server on top of Yzma. And underneath both of them sits llama.cpp , the C/C++ inference engine that made running LLMs on ordinary hardware practical.
In this series we’re going to understand that whole stack from the inside: what llama.cpp actually does when it “runs a model,” how Yzma manages to call it from Go, and how Kronk turns all of that into an engine you’d actually want in production.
The goal of this article is to understand the mechanics around the model — how a model file is stored, loaded, and run — not what happens inside the neural network itself. The math inside those layers is a whole topic of its own, and it’s out of scope here.
The llama.cpp details are against master as of July 2026 (commit ad8d8219 ) — llama.cpp moves fast and doesn’t do stable releases, so I’m pinning a commit instead of a version. And this is deliberately a conceptual tour: we’re not going to get into the C++ code details. We’ll also stay out of the Go layers entirely — Yzma and Kronk get their own articles.
With the scope set, let’s dive in.
Let’s start by demystifying the word. A language model has two very different lives.
The first life is training : you show the model terabytes of text, and an optimization process gradually adjusts billions of numbers — the weights (also called parameters ) — until the model gets good at one narrow task: given a sequence of text, predict what comes next. Training takes months and data centers. We’re not talking about training in this series, at all.
The second life is inference : training is over, the weights are frozen , and now you just… use them. The best mental model is a plain, pure function: it takes two inputs — your sequence of text and those billions of frozen weights — and returns one output, a prediction for the next token. Something like nextToken(input, weights) . That’s it. That’s the whole secret. No learning happens, nothing is updated — the weights are read-only arguments, so a model file is exactly as smart the millionth time you call it as the first.
If you’re curious about how much memory those weights actually take up — how quantization shrinks them, and why they usually dominate the VRAM you need to run a model locally — Kronk’s VRAM calculator write-up is a nice read.
So when you download a “model,” what you’re downloading is essentially a giant pile of frozen numbers — those weights — plus a description of how to use them. And “running” the model means doing a very large amount of arithmetic between your input and those numbers. If that sounds suspiciously hand-wavy, don’t worry: we’ll demystify the weights later in the article, when we crack open an actual model file. For now, all you need to hold on to is that a model is a huge pile of numbers that somebody else already calculated for you.
We’ve already mentioned tokens a couple of times now — but what exactly is a token? Let’s see.
There’s one catch before any arithmetic can happen: models don’t read text. They work with tokens — integer IDs from a fixed vocabulary, where each ID corresponds to a chunk of text (a word, a piece of a word, a punctuation mark).
Our prompt “The capital of France is " might become something like five or six token IDs — say [791, 6864, 315, 9822, 374] (the exact IDs depend entirely on the model’s vocabulary). From this point on, the model never sees letters again; everything downstream is done in terms of these integers.
So the model receives a list of integers. What does it actually do with them?
Here’s the shape of the computation, stripped to its essence. One important thing up front: the model doesn’t process the input word by word in separate rounds — the whole sequence of tokens goes through the network together, in a single pass, and out comes a single prediction. Let’s walk through it with this diagram:
Reading it left to right, we start by converting each token ID into an embedding — a vector of numbers that represents the token in a form the model can work with (if you’d like to understand embeddings better, Weaviate has a nice primer ). Then all of those vectors are fed into the model together, which runs a ton of calculations and ends up producing a list of scores — one per token in the vocabulary — saying how likely each token is to come next. And from that list, we pick one: " Paris”.
That final pick has a name — sampling — and we’ll come back to how it really works later. But notice something crucial first: all this work produced just one token. How do we get a whole sentence out of a machine that only predicts the next word?
By doing it again. And again. This is the autoregressive part, and it’s what makes an LLM conversation tick. Let’s see how it works with this diagram:
Follow the arrows. We feed our sequence — “The capital of France is " — into the model, it does one pass, and we pick a token: " Paris”. Here’s the trick that turns a single prediction into a whole answer: instead of stopping, we append " Paris" to the sequence and run the model again , now on “The capital of France is Paris”. The next token might be “.”, we append that too, and around we go — one new token per lap — until the model picks a special end-of-generation token (its way of saying “I’m done”), or we hit a length limit and stop it ourselves.
Every word of every LLM answer you’ve ever read was produced one token at a time by this loop. Generation isn’t one big computation — it’s the same next-token prediction, in a for loop.
So much for the theory. Before we can watch that loop actually run, everything hinges on one thing we’ve been hand-waving about: those frozen weights. Where do they come from, and what does a “model” actually look like on disk? Time to open the box.
GGUF: The Whole Model in One File
Time to demystify those weights. That pile of frozen numbers isn’t a shapeless heap: it’s organized into named arrays called tensors (a tensor is just an n-dimensional array — a matrix generalized to more dimensions) — for example, a single layer of the neural network is typically made up of several tensors. GGUF (GGML Universal File) is how that bag of tensors is stored on disk, and its whole design follows from one goal: a single, self-describing file . One .gguf holds everything needed to run the model — weights, architecture, hyperparameters, the entire tokenizer, even the chat template — with no sidecar files to download. Copy one file, run the model. (The very largest models sometimes get split into a few .gguf shards for practicality, but the principle holds.)
So what does one of these files actually look like inside? The format is documented right at the top of the GGUF header in ggml — the tensor library llama.cpp is built on, and where the file format itself lives ( ggml/include/gguf.h ). But rather than paste the raw spec, let me draw it for you:
Let’s read it top to bottom. Right at the start sits a tiny header — the magic bytes GGUF (so a program can check “yes, this really is a GGUF file”), a version number, and two counts telling the reader how many tensors and how many metadata entries to expect. Everything after that falls into three sections, and the rest of this section is really just a guided tour of those three boxes: metadata , tensor descriptors , and the data blob . Let’s take them one at a time.
The metadata: a model that describes itself
That’s box 1 in the diagram. The metadata is a flat list of typed key-value pairs — the values are simple scalars, strings, and homogeneous arrays. The keys are namespaced strings, and there are a few namespaces (the ones you can spot in the diagram) I want to call your attention to:
general.* — identity: general.architecture says which kind of model this is ( llama , qwen2 , gemma , …). This one key decides which graph-building code llama.cpp will use to set up the inference machinery.
<arch>.* — hyperparameters, prefixed by the architecture name: llama.embedding_length (the size of those token vectors), llama.block_count (how many layers), llama.context_length (how many tokens fit in the window), etcetera…
tokenizer.* — the entire tokenizer (the piece responsible for turning our text into tokens) lives inside the model file . tokenizer.ggml.tokens is the full vocabulary as a string array, alongside the tokenizer’s configuration parameters and even tokenizer.chat_template — the template that turns a chat conversation into a flat prompt string.
When llama.cpp opens a file, it reads general.architecture first, and everything else follows from there.
That’s the what of the model. Now let’s see where the actual numbers live.
The blob: getting the weights into memory
Now for boxes 2 and 3 in the diagram, which work as a pair. Box 2 is the tensor descriptors : one row per tensor with its name, shape, and type — but no data . Instead, each row carries an offset , and that’s the red arrow in the picture: the offset points at where that tensor’s bytes actually live inside box 3, the data blob .
The interesting part is how those bytes get into memory. The loader ( llama_model_loader , src/llama-model-loader.cpp ) first parses only the header, metadata, and tensor descriptors — no tensor data is read at all. That’s enough to know what goes where: some layers stay on the CPU, others go to the GPU, and each destination handles its bytes its own way. For CPU tensors, llama.cpp by default doesn’t read() gigabytes of weights into freshly allocated buffers — it maps the file into the process’s address space with mmap and points each tensor directly at mapping_address + tensor_offset . For GPU tensors that trick doesn’t work — VRAM is a different memory altogether — so their bytes really do get copied: uploaded from the file into GPU buffers, once, at load time.
So now we have the file’s tensors sitting in memory, each in its right place. Time to run something.
The KV Cache: Why Generation Stays Fast
Let’s start with what we already have: our model’s data, the part that never changes. Loading it hands us the llama_model — the file itself, in memory: the weights, the tokenizer, all the read-only stuff. Because nothing about it changes, it can be shared — one loaded model, many conversations, and those multi-gigabyte weights sit in memory only once.
Now we need the dynamic data, and that’s what the llama_context is for: it represents a single conversation, and it’s where all the changing state lives — the working space for the calculations and, above all, the KV cache . Compared to loading a model, creating a context is cheap, so you spin up one per conversation.
So what is this KV cache? At every step of the loop we hand the model an ever-longer sequence, and reprocessing it from scratch on every pass would be hugely wasteful. Avoiding that is the job of the KV cache .
The trick is that while processing a token, the model computes some intermediate values for it that never change afterwards. So instead of recomputing them, llama.cpp stores them in the KV cache and reuses them. Each new token only has to do its own work and read the rest from memory — so every generated token costs roughly the same, no matter how long the text has al

[truncated]
