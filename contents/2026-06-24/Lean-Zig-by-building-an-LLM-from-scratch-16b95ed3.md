---
source: "https://hamanlp.org/small-kernels.html"
hn_url: "https://news.ycombinator.com/item?id=48659158"
title: "Lean Zig by building an LLM from scratch"
article_title: "Small Kernels — Learn Zig and LLMs by Building Both From Scratch"
author: "boodleboodle"
captured_at: "2026-06-24T13:43:43Z"
capture_tool: "hn-digest"
hn_id: 48659158
score: 1
comments: 1
posted_at: "2026-06-24T13:08:26Z"
tags:
  - hacker-news
  - translated
---

# Lean Zig by building an LLM from scratch

- HN: [48659158](https://news.ycombinator.com/item?id=48659158)
- Source: [hamanlp.org](https://hamanlp.org/small-kernels.html)
- Score: 1
- Comments: 1
- Posted: 2026-06-24T13:08:26Z

## Translation

タイトル: LLM をゼロから構築することによる Lean Zig
記事のタイトル: 小さなカーネル — Zig と LLM の両方を最初から構築して学ぶ

記事本文:
2 値、タイプ、制御フロー
10 アクティベーション、ソフトマックス、ノルム
13 クロスエントロピーと分類
パート 4 · ゼロからの言語
33 エンコーダー - デコーダー & クロスアテンション
単一のドット積から GPT-2、最新の Gemma、Whisper への入り口に至るまで、すべての部分を自分で作成することで、Zig プログラミング言語と大規模な言語モデルが実際にどのように機能するかを学びます。
私はもともと、Zig と LLM を独学し、MLE の面接に備えるために、この本をバイブコーディングしました。皆さんと同じように、私もクローディズムには不快感を感じているので、当初の意図は出版する前にすべてを手書きで書き直すことでした。
残念ながら採用されてしまいました。そして、私の新しい雇用主はサイドプロジェクトの公開に対してより厳格なので、新しい仕事を始める前にこれをそのまま公開しています。それでも非常に役立つので、この資料を最後まで読んで、最終的には Zig で人気のある LLM および ASR モデルの推論エンジンを構築することをお勧めします。
エラーや改善のアイデアを見つけた場合は、お気軽にご連絡ください。
この本が役立つと思われた場合は、Zig に寄付してください。
最初にお読みください: 2 つのスキル、1 つの旅
Zig と呼ばれるプログラミング言語と、大規模な言語モデルが実際にどのように機能するかという 2 つのことを同時に学習することになります。そして、これらを習得する唯一の確実な方法は、何もないところから小さな部分をすべて構築することです。 2 つの数値リストを乗算する関数から始めます。最後に、GPT-2 からテキストを生成し、次に最新の Gemma モデルからテキストを生成する独自のコードで、音声認識への明確な道筋を示します。難しい部分を代わりにやってくれるフレームワークはありません。最後のページまでに、変圧器を動かすコードを入力しているので、変圧器を流れるすべての数値を理解できるようになります。
この本は、Zig の使用経験や実際のシステム プログラミングの経験がないことを前提としています。何らかの言語で小さなコードを書いたことがある場合、

ループ、関数、if だけで十分です。パート 1 では、Zig を 1 行目から学びます。まったくプログラミングをしたことがない場合でも、このまま進めていただけます。パート 1 ではゆっくりと進み、すべての例を実行してください。数学も同じです。高校の代数 (乗算、合計、ちょっとした表記法) に慣れる必要があり、内積や微分などのアイデアが浮かんだときに喜んでそれに取り組む姿勢が必要です。シンボルを 1 つ増やすのではなく、必要なだけ正確に計算を構築します。
なぜなら、LLM はその下に巨大な数値の山とその上にいくつかのループがあるからです。Zig を使用すると、これらのループを何も隠さずに見ることができます。いつ一時停止するかを決定するガベージ コレクターも、隠れた割り当ても、密かにライブラリを呼び出すオペレーターもありません。 (a, b) |x, y| と書くと合計 += x * y; CPU が実行する実際の演算を見ていることになります。この透明性は、迅速なスクリプト作成にとっては呪いであり、モデルが実際にどのように実行されるかを学ぶための贈り物でもあります。おまけに、本当に役立つ低水準言語を知ることができるでしょう。
この本は単一の階段として続きます。各ステップは小さく、各ステップが実行され、各ステップが次のステップで使用されます。私たちはこのステップをカーネルと呼んでいます。この言葉には、私たちが頼りにする 2 つの意味があります。システム プログラマにとって、「カーネル」は 1 つの数値ジョブ (行列の乗算、ソフトマックス) を実行する緊密なループです。他の人にとって、それは物事の核、種を意味します。どちらもここにあります。LLM を一度に 1 つの小さな数値カーネルずつ成長させ、各カーネルが次のアイデアの種となります。
コードを貼り付けるのではなく、入力します。すべての値はあなたの指とデバッガにあります。すべてのコード ブロックを、コピーするスニペットとしてではなく、再実装するものとして扱います。 [コピー] ボタンは、入力をスキップするためではなく、自分のバージョンと私のバージョンを比較するために存在します。
すべてを実行します。すべてのカーネルの後に wa があります

y は、テスト、印刷された数字、生成された単語など、機能することを確認します。実行していないブロックを決して読み込まないでください。機能する「はず」のモデルと、実際に機能するのを見たモデルは、異なる種類の知識です。
オラクルを保管してください。パート 3 以降では、既知の正しい参照 (NumPy または PyTorch を使用した Python の数行) に対して数値をチェックします。あなたの出力がオラクルと一致しない場合、そうでないことが証明されるまでは、あなたは間違っていることになります。付録 D は、これらを構築するためのレシピブックです。
各ヘッダーのボタンを使用してセクションに完了のマークを付けます。進行状況とテーマはこのファイルのローカル ストレージに保存されるため、タブを閉じて来週から始めることができます。
ゆっくりでも大丈夫です。これは何ヶ月もかけて夜をかけて登ることができる階段です。すべての部分が実行されることで終わるため、満足のいくチェックポイントから遠く離れていることはありません。
すべてのコードは、この本が書かれた時点で最新の行である Zig 0.16 をターゲットとしています。 Zig は 1.0 より前であり、リリース間で変更されます。特に標準ライブラリと build.zig 形式が作り直されています。スニペットがコンパイルされない場合、最初に疑われるのはバージョンの不一致です。 zig version を実行します。新しいものを使用している場合は、空のフォルダーで zig init を実行し、生成されたファイルを私たちのファイルと比較します。付録 A には、ドリフトする可能性が最も高いスポットが記載されています。
10部。それぞれの行は、走って感じることができるマイルストーンです。
これは無関係な 3 つのプロジェクトではありません。 GPT-2 は最も単純な実際のトランスフォーマーであり、テキストを読み取り、次のトークンを予測するデコーダーです。 Gemma は、すべての部分が最新化された同じデコーダーです。一度に 1 つずつ部品を交換し、2018 年モデルが 2025 年モデルになるのを観察します。 Whisper は、同じトランスに後半部分 (音声を読み取るエンコーダーとクロスアテンション ブリッジ) が取り付けられているため、デコーダーは聞きながらテキストを予測します。デコーダを深く学び、他の 2 つはほとんど理解できません。

再配線します。そのため、カリキュラムはこのように順序付けされています。各モデルは、最後のモデルに 1 つの新しいアイデアを追加します。
Zig で作成したプログラム — small-llm -m gpt2.safetensors -p "OnceUpon a time" — は、実際の重みをロードし、流暢な継続テキストを出力します。ここでは、任意の段階をポイントして、すべての数値が何をしているのか説明できます。そこから、現代​​の Gemma モデルとスピーチへの明確な入り口が見えてきます。魔法ではなく理解です。それが成果物です。
Zig をまったく使用しない状態から、実際のプログラムの作成、構築、テストに至るまでの 6 章。私たちは言語だけを学びます - 数学はパート 2 から始まります - しかし、すべての例は、今後の数値コードに傾いています。最後には、本全体がその中に収まるプロジェクトの骨組みが完成します。
インストールするものは 2 つあります: Zig コンパイラーと Zig を理解するエディターです。次に、プログラムを作成、実行、分析します。これにより、この短い章が終わるまでにツールが邪魔にならず、アイデアに集中できるようになります。
公式ダウンロード ページからマシン用のビルドをダウンロードします。0.16 行を選択します。 Zig は、単一の自己完結型フォルダーとして出荷されます。システムの意味で「インストール」するものは何もありません。安定した場所で解凍し、そのフォルダーを PATH に置きます。
Windows では、解凍してシステム設定からフォルダーを PATH に追加するか、最初はフルパスで実行します。今のところ Homebrew/apt は避けてください。パッケージ化された Zig は、動きの速いリリースに遅れることが多く、初心者のコードがコンパイルできない最も一般的な理由はバージョンのずれです。
その 1 つのフォルダーには、コンパイラー、標準ライブラリ全体 (読み取り可能な Zig ソースとして開くことをお勧めします)、および C/C++ コンパイラーも含まれています。Zig は C プロジェクトをビルドできます。個別のランタイムをインストールしたり、グローバルに構成したりする必要はありません。 Zig を「アンインストール」するには、フォルダーを削除します。このシンプルな

icity はこの言語のブランドに非常に適しています。
任意のエディタを使用しますが、オートコンプリート、定義へのジャンプ、およびインライン エラーのために ZLS (Zig Language Server) をインストールすると、学習が劇的に速くなります。公式の「Zig」拡張機能を備えた VS Code が最も簡単に始めることができます。 Neovim、Helix などもうまく機能します。拡張機能で zig バイナリを指定し、一致する ZLS を取得させます。
ファイル hello.zig を作成し、次のように正確に入力します。
直接実行します。プロトタイピング中に別のコンパイル手順は必要ありません。
これらの行は何千回も書くことになるため、ここにあるすべてのトークンは理解する価値があります。
const std = @import("std"); — @import は組み込みです (すべての組み込みは @ で始まります)。モジュールを取得し、それを定数 std (標準ライブラリ) にバインドします。 const は「この名前は決して変更されない」という意味です。
pub fn main() void — プログラムのエントリ ポイントである main という名前の関数を宣言します。 pub により、このファイルの外でそれが見えるようになります (ランタイムがそれを認識する必要があります)。 void は戻り値の型であり、何も返しません。
std.debug.print(fmt, args) — フォーマットされた文字列を出力します。最初の引数はフォーマット文字列です。 2 番目は、それを埋めるための引数です。
.{} — これは空の匿名構造体リテラルです。 Zig は print の引数を構造体として渡します (ここでは引数がないので空です)。 .{ ... } が常に表示されます。今のところは、それを「小さな価値観の束」として読んでください。
フォーマット文字列内では、 {} スタイルのプレースホルダーが args バンドルから順番に埋められます。一般的なもの: {s} は文字列/バイトスライス、{d} は 10 進数の整数または浮動小数点数、{d:.4} は小数点 4 桁の浮動小数点、{any} は「できる限りフォーマットする」(デバッグ中の配列に最適)、および {c} は文字としての 1 バイトです。 \n は改行です。プレースホルダーの数が引数の数と一致しない場合、Zig はコンパイル時に通知します。
を証明しましょう

プレースホルダーは機能します。数字とループ — これから先のすべての原料:
実行すると、カウント、合計 2.300 、および配列が出力されるはずです。正確な型についてはまだ心配する必要はありません — []const u8 、 f32 、 [_]f32 — 第 2 章と第 3 章でそれらを解凍します。現時点での利点は、何でも印刷できる、つまり何でも見ることができるということです。これが、この本のすべてのカーネルをデバッグする方法です。
std.debug.print は、標準出力ではなくエラー ストリームに書き込みます。これは、本書全体で「この番号を表示する」すべての作業に使用するツールです。現在の Zig で実際のプログラム出力を stdout に書き込むには、バッファ付きライターを経由します。これは、パート 6 で実際に生成されたテキストを出力するときに適切に設定します。それまでは、std.debug.print だけで十分なので、邪魔になることはありません。
重みの最大値も出力するように、numbers.zig を変更します。ループ内で更新する var と組み込みの @max(a, b) が必要です。次に、配列の値を変更し、コードが適応することを確認します。小さいですが、これはドット積 (第 7 章) が構築される配列パターンのループです。
zig バージョンは 0.16 リリースを出力し、zig run hello.zig が挨拶し、文字列、整数、浮動小数点数、および配列を出力できます。道具は邪魔にならない。さて言語です。
Zig は静的に型指定されており、非常に明示的です。すべての値の型はコンパイル時に認識され、整数のサイズは詳しく説明され、数値型間の変換はユーザーが要求するものであり、背後で行われることはありません。最後のルールは、初心者を午後つまずかせますが、その後、他の言語では見逃してしまう機能になります。私たちはそれを正面から迎えます。
単純な int はありません。あなたはまさにあなたが言いたいことを言います：u8は符号なし8ビット整数（0〜255）、i32は符号付き32ビット整数、u64は大きな符号なし整数です。 usesize は unsigned int です

たとえば、メモリ アドレスの幅です。これはあらゆる長さとインデックスの型であるため、常に使用することになります。実数については、本書のほぼすべての場所で f32 (32 ビット浮動小数点) を使用します。これは、ニューラル ネットワークの数学が実行されるためです。
const は、名前を決して変更されない値にバインドします。 var は再割り当てできる変数です。 Zig は const に誘導します。 var を宣言してそれを変更しない場合、コンパイラーはエラーを起こし、それを const にするように指示します。これは煩わしいように聞こえますが、実際には静かなスーパーパワーです。コード内の var は「これが変更される」という信号であるため、ループとアキュムレータが一目で目立ちます。
Zig は数値型を黙って混合しません。 i32 を usize に追加することはできません。また、整数を別の整数で除算して浮動小数点を取得することも、当然のことながらできません。よく利用する 2 つの組み込み機能:
@floatFromInt(x) — 整数を浮動小数点に変換します。
@intFromFloat(x) — 浮動小数点を整数に変換します (切り捨て)。
これらは「結果配置」です。Zig は、値の行き先からターゲットの型を判断します。そのため、それらを @as(T, ...) で囲んで、その型を明示的に指定することがよくあります。典型的なケースは、合計をカウントで割って平均を計算することです。
これは初心者が最もよく遭遇するエラーです。これで修正方法がわかりました。整数が浮動小数点である必要がある (またはその逆)。

[切り捨てられた]

## Original Extract

2 Values, types & control flow
10 Activations, softmax & norms
13 Cross-entropy & classification
Part 4 · Language from scratch
33 Encoder–decoder & cross-attn
Learn the Zig programming language and how large language models actually work — by writing every piece yourself, from a single dot product to GPT-2, a modern Gemma, and the doorway to Whisper.
I originally vibe-coded this book to teach myself Zig and LLMs, and prepare for MLE interviews. Like everyone I find the Claudisms off-putting, so my original intention was to rewrite everything by hand before publishing.
Unfortunately, I was hired. And my new employer is more strict about publishing side-projects, so I am publishing this as-is before I start my new job. It is still very helpful, so I recommend you go through the material, eventually building inference engines for popular LLMs and ASR models in Zig.
If you find any errors or ideas for improvement, feel free to reach out .
If you find this book helpful, please donate to Zig !
Read me first: two skills, one journey
You're going to learn two things at once — a programming language called Zig , and how a large language model actually works — and you're going to learn them the only way that really sticks: by building, from nothing, every small piece. We start with a function that multiplies two lists of numbers. We end with your own code generating text from GPT-2, then from a modern Gemma model, with a clear path into speech recognition. No framework does the hard part for you. By the last page you will understand every number that flows through a transformer, because you will have typed the code that moves it.
This book assumes no prior Zig and no real systems-programming background. If you've written a little code in any language — a loop, a function, an if — you have enough. Part 1 teaches Zig from the first line. If you've never programmed at all, you can still follow along; just go slower through Part 1 and run every example. The math is the same story: you need comfort with high-school algebra (multiplying, summing, a little notation) and a willingness to meet ideas like the dot product and the derivative when they arrive. We build the math up exactly as much as we need and not one symbol more.
Because an LLM is, underneath, a giant pile of numbers and a handful of loops over them — and Zig lets you see those loops with nothing hidden. There's no garbage collector deciding when to pause, no hidden allocations, no operator that secretly calls a library. When you write for (a, b) |x, y| sum += x * y; you are looking at the actual arithmetic the CPU will do. That transparency is a curse for quick scripting and a gift for learning how a model really runs. As a bonus you'll come out knowing a genuinely useful low-level language.
The book runs as a single staircase. Each step is small, each step runs , and each step is used by the next. We call the steps kernels — a word with two meanings we'll lean on. To a systems programmer a "kernel" is a tight loop that does one numerical job (a matrix multiply, a softmax). To everyone else it means the core, the seed of a thing. Both are right here: we grow an LLM one small numerical kernel at a time, and each kernel is the seed of the next idea.
Type the code, don't paste it. The entire value is in your fingers and your debugger. Treat every code block as something to re-implement, not a snippet to copy. The Copy button exists for when you're comparing your version against mine, not for skipping the typing.
Run everything. After every kernel there's a way to see it work — a test, a printed number, a generated word. Never read past a block you haven't run. A model that "should" work and a model you've watched work are different kinds of knowledge.
Keep an oracle. From Part 3 on, we check our numbers against a known-correct reference — a few lines of Python with NumPy or PyTorch. When your output disagrees with the oracle, you are wrong until proven otherwise. Appendix D is the recipe book for building these.
Mark sections complete with the button in each header. Your progress and theme are saved in this file's local storage, so you can close the tab and pick up next week.
It's okay to be slow. This is a staircase you can climb over months of evenings. Every part ends in something that runs, so you're never far from a satisfying checkpoint.
All code targets Zig 0.16 , the line current as this book was written. Zig is pre-1.0 and changes between releases — the standard library and the build.zig format in particular get reworked. If a snippet doesn't compile, the very first suspect is a version mismatch: run zig version , and if you're on something newer, run zig init in an empty folder and compare its generated files against ours. Appendix A notes the spots most likely to drift.
Ten parts. Each row is a milestone you can run and feel.
This isn't three unrelated projects. GPT-2 is the simplest real transformer — a decoder that reads text and predicts the next token. Gemma is the same decoder with every part modernized; you'll swap pieces one at a time and watch a 2018 model become a 2025 one. Whisper is that same transformer with a second half bolted on — an encoder that reads audio and a cross-attention bridge — so the decoder predicts text while listening . Learn the decoder deeply and the other two are mostly rewiring. That's why the curriculum is ordered this way: each model adds exactly one new idea to the last.
A program you wrote in Zig — small-llm -m gpt2.safetensors -p "Once upon a time" — that loads real weights and prints fluent continuation text, where you can point at any stage and explain what every number is doing. From there, a clear on-ramp to a modern Gemma model and to speech. Understanding, not magic. That's the deliverable.
Six chapters that take you from no Zig at all to writing, building, and testing real programs. We learn only the language — the math starts in Part 2 — but every example leans toward the numerical code to come. By the end you'll have the project skeleton the whole book grows inside.
Two things to install: the Zig compiler and an editor that understands Zig. Then we write, run, and dissect a program, so that by the end of this short chapter the tooling is out of the way and you can focus on ideas.
Download the build for your machine from the official downloads page — pick the 0.16 line. Zig ships as a single self-contained folder; there's nothing to "install" in the system sense. Unpack it somewhere stable and put that folder on your PATH .
On Windows, unzip it and add the folder to your PATH through System Settings, or just run it by full path while you're starting out. Avoid Homebrew/apt for now — packaged Zig often lags the fast-moving releases, and version drift is the single most common reason a beginner's code won't compile.
That one folder contains the compiler, the entire standard library (as readable Zig source you're encouraged to open), and a C/C++ compiler too — Zig can build C projects. There is no separate runtime to install and nothing global to configure. To "uninstall" Zig you delete the folder. This simplicity is very on-brand for the language.
Use any editor, but install ZLS (the Zig Language Server) for autocomplete, jump-to-definition, and inline errors — it makes learning dramatically faster. VS Code with the official "Zig" extension is the easiest start; Neovim, Helix, and others work well too. Point the extension at your zig binary and let it fetch the matching ZLS.
Create a file hello.zig and type this exactly:
Run it directly — no separate compile step needed while prototyping:
Every token here is worth understanding, because you'll write these lines a thousand times.
const std = @import("std"); — @import is a builtin (all builtins start with @ ); it pulls in a module and binds it to the constant std , the standard library. const means "this name never changes."
pub fn main() void — declares a function named main , the program's entry point. pub makes it visible outside this file (the runtime needs to see it). void is the return type: it returns nothing.
std.debug.print(fmt, args) — prints a formatted string. The first argument is a format string ; the second is the arguments to fill it in.
.{} — this is an empty anonymous struct literal . Zig passes print's arguments as a struct (here, no arguments, so it's empty). You'll see .{ ... } constantly; for now read it as "a little bundle of values."
Inside the format string, {} -style placeholders are filled from the args bundle in order. The common ones: {s} for a string/byte-slice, {d} for an integer or float in decimal, {d:.4} for a float with 4 decimals, {any} for "format this however you can" (great for arrays while debugging), and {c} for a single byte as a character. The \n is a newline. If the number of placeholders doesn't match the number of args, Zig tells you at compile time .
Let's prove the placeholders work. Numbers and a loop — the raw material of everything ahead:
Run it: you should see the count, the sum 2.300 , and the array printed out. Don't worry about the exact types yet — []const u8 , f32 , [_]f32 — Chapters 2 and 3 unpack them. Right now the win is: you can print anything, which means you can see anything, which is how we'll debug every kernel in this book.
std.debug.print writes to the error stream, not standard output, and it's the tool we use for all the "let me see this number" work throughout the book. Writing real program output to stdout in current Zig goes through a buffered writer, which we'll set up properly in Part 6 when we actually emit generated text. Until then, std.debug.print is all you need and it never gets in your way.
Modify numbers.zig to also print the largest value in weights . You'll need a var that you update inside the loop and the builtin @max(a, b) . Then change the array's values and confirm your code adapts. Small, but it's the loop-over-an-array pattern that the dot product (Chapter 7) is built from.
zig version prints a 0.16 release, zig run hello.zig greets you, and you can print strings, integers, floats, and arrays. The tools are out of the way. Now the language.
Zig is statically typed and refreshingly explicit: the type of every value is known at compile time, integer sizes are spelled out, and conversions between number types are something you ask for, never something that happens behind your back. That last rule trips up newcomers for an afternoon and then becomes a feature you'll miss in other languages. We'll meet it head-on.
There is no plain int . You say exactly what you mean: u8 is an unsigned 8-bit integer (0–255), i32 a signed 32-bit integer, u64 a big unsigned one. usize is an unsigned integer the width of a memory address — it's the type of every length and index, so you'll use it constantly. For real numbers we use f32 (32-bit float) almost everywhere in this book, because that's what neural network math runs in.
const binds a name to a value that never changes; var is a variable you can reassign. Zig nudges you toward const — if you declare a var and never mutate it, the compiler errors and tells you to make it const . This sounds annoying and is actually a quiet superpower: a var in the code is a signal that "this changes," so loops and accumulators stand out at a glance.
Zig will not silently mix number types. You cannot add an i32 to a usize , and you cannot divide an integer by another and get a float, without saying so. The two builtins you'll reach for most:
@floatFromInt(x) — turn an integer into a float.
@intFromFloat(x) — turn a float into an integer (truncating).
These are "result-located": Zig figures out the target type from where the value is going, which is why you often wrap them in @as(T, ...) to name that type explicitly. The classic case is computing an average — dividing a sum by a count:
This is the error you'll hit most as a beginner, and now you know the fix: something is an integer that needs to be a float (or vice-

[truncated]
