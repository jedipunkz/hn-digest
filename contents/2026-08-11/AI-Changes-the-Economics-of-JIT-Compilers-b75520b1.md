---
source: "https://malisper.me/how-ai-changes-the-economics-of-jit-compilers/"
hn_url: "https://news.ycombinator.com/item?id=49261233"
title: "AI Changes the Economics of JIT Compilers"
article_title: "How AI Changes the Economics of JIT Compilers - malisper.me"
author: "poly2it"
captured_at: "2026-08-11T17:49:52Z"
capture_tool: "hn-digest"
hn_id: 49261233
score: 1
comments: 0
posted_at: "2026-08-11T17:03:51Z"
tags:
  - hacker-news
  - translated
---

# AI Changes the Economics of JIT Compilers

- HN: [49261233](https://news.ycombinator.com/item?id=49261233)
- Source: [malisper.me](https://malisper.me/how-ai-changes-the-economics-of-jit-compilers/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T17:03:51Z

## Translation

タイトル: AI が JIT コンパイラーの経済性を変える
記事のタイトル: AI が JIT コンパイラーの経済性をどのように変えるか - malisper.me
説明: Rust でコピーとパッチの JIT コンパイラを構築し、ARM64 マシン コードを生成し、インタプリタと手書きの実装に対してベンチマークを実行します。

記事本文:
AI が JIT コンパイラーの経済性をどのように変えるか - malisper.me
プライマリーメニュー
malisper.me
Postgres 投稿の目次
電子メール アドレスを入力してこのブログを購読し、新しい投稿の通知を電子メールで受け取ります。
ホーム » pgrust » AI は JIT コンパイラーの経済性をどのように変えるか
AI が JIT コンパイラーの経済性をどのように変えるか
歴史的に、JIT コンパイルは黒魔術でした。高速な JIT コンパイラを作成するには、アセンブリの作成方法を知る必要があります。その好例: 現在、独自の JIT コンパイラーを備えた運用対応データベースは存在しません。これらはすべて、LLVM を使用するか、C/C++ コードを生成します。これらのオプションはどちらもコンパイル時間が長いため、適用性が制限されます。 AI の使用により、アセンブリを直接ターゲットにすることで、コンパイル時間の短い JIT コンパイラーをこれまでより簡単に作成できるようになりました。これは、新しいデータベースが古いデータベースを改善できる可能性のある領域の 1 つでもあります。 pgrust を構築するとき、私は最初、JIT コンパイラーを実装するのは非常に難しいだろうと思いました。最終的には、AI 支援のおかげで予想よりもはるかに簡単であることがわかり、それが pgrust が非常に高速である理由の 1 つとなっています。この投稿では、独自の JIT コンパイラーを構築する方法を説明します。例として JIT コンパイルを使用する単純な正規表現エンジンを構築します。
JIT コンパイルは、実行時または「ジャスト イン タイム」でコンパイルされたコードを生成する手法です。正しく実行すると、パフォーマンスが大幅に向上し、多くの場合 2 ～ 5 倍、場合によってはそれ以上の効果が得られます。 JIT コンパイルの主な使用例は、プログラムの動作を大幅に変更する情報を実行時に取得する場合です。これは、プログラミング言語インタプリタで特に一般的です。実行時に実行するコードを受け取ります。 JIT コンパイラはドメインでも役立ちます

データの解析など、プログラミング言語を超えた機能。場合によっては、実行時まで解析しているデータのスキーマがわからないことがありますが、JIT はそれを解決するのに役立ちます。
まずは、おもちゃの正規表現エンジンを実装してみましょう。物事を簡単にするために、リテラル文字列と繰り返し (つまり、正規表現 *) という 2 つの機能のみをサポートします。また、パーサーをスキップし、正規表現をすでに解析された Rust 構造として表します。これは、次のような文字列をサポートできることを意味します。
ただし、交代や後ろ向きなどはありません。
コードではこれは非常に簡単です。リテラル文字列ノード、繰り返しノード、および 2 つのノードの組み合わせである連結ノードの 3 種類のノードがあります。最終的には次のようになります。
列挙型ノード {
リテラル(&'静的文字列)、
Concatenation(Box<Node>, Box<Node>),
繰り返し(ボックス<ノード>)、
}
fn リテラル(テキスト: &'static str) -> ノード {
ノード::リテラル(テキスト)
}
fn concatenation(左: ノード、右: ノード) -> ノード {
Node::Concatenation(Box::new(左), Box::new(右))
}
fn 繰り返し(本体: ノード) -> ノード {
Node::Repetition(Box::new(body))
}
正規表現エンジンのインタープリタを作成するのも簡単です。
fn match_node(node: &Node, input: &[u8], pos: usesize, next: &dyn Fn(usize) -> bool) -> bool {
マッチノード {
Node::Literal(text) => {
リテラル = text.as_bytes(); とします。
input[pos..].starts_with(literal) && next(pos + literal.len())
}
Node::Concatenation(left, right) => {
match_node(left, input, pos, &|left_end| {
match_node(右、入力、左端、次)
})
}
Node::Repetition(body) => {
match_node(body, input, pos, &|body_end| {
match_node(ノード、入力、ボディエンド、次)
}) ||次(位)
}
}
}
fn interp_match(正規表現: &Node, 入力: &str) -> bool {
let bytes = input.as_bytes();
match_node(regex, bytes, 0, &|pos| pos == bytes.len

())
}
この正規表現エンジンは非常にシンプルです。コードは 20 行未満ですが、パフォーマンスの点でどのように機能するかを見てみましょう。比較のために、コードを正規表現専用に実装された手書きコードと比較します。この例では、正規表現 b(an)* を使用します。手書きのコードは次のようになります。
fn 手書き_b_an_star(入力: &str) -> bool {
let bytes = input.as_bytes();
mut pos = 0 とします。
if pos == bytes.len() ||バイト[位置] != b'b' {
false を返します。
}
pos += 1;
while pos < bytes.len() {
if bytes[pos] != b'a' {
false を返します。
}
pos += 1;
if pos == bytes.len() ||バイト[位置] != b'n' {
false を返します。
}
pos += 1;
}
本当の
}
(このコードを最適化して大幅に高速化する方法はありますが、ここでは良い比較として役立ちます)
これら 2 つに対していくつかの例をベンチマークすると、手書きバージョンはインタプリタより 10 ～ 20 倍高速であることがわかりました。明らかに改善の余地がたくさんあります。
ここで、JIT コンパイルを使用して、手書きバージョンと同様に機能する一般的な正規表現エンジンを取得する方法を見てみましょう。
コードを JIT コンパイルするには 2 つの手順があります。まず、実行するコードのアセンブリを生成します。コードを取得したら、アセンブリ コードを関数にパッケージ化し、他のコードと同様にプログラム内で呼び出すことができます。
アセンブリを生成するには、コピーとパッチと呼ばれるアプローチの変形を使用します。考え方としては、JIT コンパイルしたいさまざまな操作に対応する一連のテンプレートをアセンブリ内に用意するということです。このテンプレートを「ステンシル」と呼びます。操作を JIT コンパイルする場合は、関連するステンシルを取得し、操作の詳細に基づいて小さな調整を加えます。実際のステンシルを塗りつぶすのと非常によく似ています。これをいくつか組み合わせることで、

塗りつぶされたステンシルを使用すると、実行時に手書きバージョンと同様のパフォーマンスを持つプログラムを構築できます。
私たちがたどるパスは次のとおりです。まず、b(an)* 用に生成する ARM64 コードを見ていきます。次に、繰り返される命令シーケンスを再利用可能なステンシルに変換し、正規表現 AST からこれらのステンシルを埋めて結合するエミッタを作成します。最後に、生成された命令を実行可能メモリにコピーして、Rust が通常の関数のように呼び出せるようにします。
これがどのように機能するかを説明するには、生成されたコードから始めて、JIT コンパイラー自体にまで遡って作業するのが最も簡単です。ここでも、正規表現「b(an)*」を使用しています。設計上の決定事項をいくつか説明すると、次のようになります。
バックトラッキングにはスタックを使用します。スタックは、正規表現で行き詰まりになった場合に進むべき状態を追跡します。
照合対象の文字列は null バイトで終わります。つまり、文字列の末尾に達すると、文字比較は自動的に失敗します。これは、どの時点でも長さの比較を行う必要がないことを意味します
プログラムの状態については、次のレジスタを使用します。
x0 – 文字列内の現在位置と戻り値
x1 – バックトラッキングに使用されるスタックの最上位
x2 – バックトラックに使用されるスタックの一番下 (これはスタックが空かどうかを判断するために必要です)
x9 – 一時変数として使用されます
プログラムへの入力として、次のものが渡されます。
x0 – 文字列の先頭へのポインタ
x1 – スタックに使用する場所へのポインタ
ここまでは終わったので、生成されたアセンブリをパーツごとに見ていきましょう。これは特に ARM64 を搭載した macOS で発生します。まず、プログラムを初期化するプロローグがあります。スタックの最上位と最下位を値に設定してスタックを初期化するだけです。

e が渡されました:
0: aa0103e2 mov x2、x1
次に、文字 b をチェックするコードがあります。 b ではない文字が見つかった場合は、フォールバック ロジックを処理するコード ブロックにジャンプします。それ以外の場合は、文字列内の位置を進めます。
; CHAR「b」
4: 39400009 ldrb w9、[x0] ;負荷電流入力バイト
8: 7101893f cmp w9、#0x62 ;それは「b」ですか？
c: 54000281 b.ne 0x5c ;いいえ -> フォールバックブロック
10: 91000400 x0、x0、#1 を追加します。はい -> 事前入力
次に、繰り返し (an)* があります。繰り返しの場合はバックトラッキングを行う必要があります。ここでバックトラックすると、ループの最後にすぐにジャンプすることになります。つまり、ループ後の命令のアドレスとスタック上の文字列内の位置の両方を保存する必要があります。
14: d2800989 movz x9、#0x004c ;ビルド再開アドレス
18: f2a00009 movk x9、#0x0000、lsl #16 ; = 0x1_0000_004c
1c: f2c00029 movk x9、#0x0001、lsl #32 ; (ループ出口)
20: f2e00009 movk x9、#0x0000、lsl #48 ;
24: a8810029 stp x9、x0、[x1]、#16 ; (exit, pos) をスタックにプッシュします
これで、繰り返しの本体を実行できるようになります。これにより、文字「a」と「n」がチェックされ、それらが見つかった場合は、繰り返しの先頭に戻りますが、新しい文字列の位置に戻ります。
; CHAR「あ」
28: 39400009 ldrb w9、[x0]
2c: 7101853f cmp w9、#0x61 ; 「あ」？
30: 54000161 b.ne 0x5c ;いいえ -> フォールバックブロック
34: 91000400 x0、x0、#1 を追加
;チャー「ん」
38: 39400009 ldrb w9、[x0]
3c: 7101b93f cmp w9、#0x6e ; 「ん」？
40: 540000e1 b.ne 0x5c ;いいえ -> フォールバックブロック
44: 91000400 x0、x0、#1 を追加
; JMP
48: 17fffff3 b 0x14 ;ループの先頭に戻る
これでループを越えました。これは、一度後戻りすると、後戻りがジャンプする場所です。繰り返しが完了すると、正規表現の終わりになります。ここでしなければならないのは、文字列の最後にいるかどうかを確認することだけです。もし私たちが終わりにいるなら

文字列の場合、成功した場合は 1 を返します。一致しない場合は、正規表現の一致に失敗したことを意味するため、フォールバックを行うために失敗ロジックを実行する必要があります。
4c: 39400009 ldrb w9、[x0]
50: 35000069 cbnz w9、0x5c ; NUL ではない -> フォールバック ブロック
54: d2800020 mov x0、#1 ;成功
58: d65f03c0 レット
そして最後に、フォールバック ロジックがあります。これにより、スタックが空かどうかがチェックされます。空の場合は 0 を返します。空でない場合は、フォールバック アドレスとフォールバック文字列の位置の両方をスタックからポップし、フォールバック アドレスにジャンプします。
5c: eb02003f cmp x1、x2 ;枠は残っていますか？
60: 54000060 b.eq 0x6c ;いいえ -> 諦めます
64: a9ff0029 ldp x9、x0、[x1、#-16]! ;ポップ (履歴書、履歴書)
68: d61f0120 br x9 ;そこにジャンプしてください
6c: d2800000 mov x0、#0 ;一致しません
70: d65f03c0 レット
ステンシルの構築
コンパイルされたコードを確認する機会が得られたので、コピーとパッチのコンパイラーがどのように動作するかを理解し始める必要があります。共通の命令セットがありますが、それらの間にはわずかな違いがあります。これらの関数ブロックごとに、それぞれのコードを生成する関数を作成できます。各関数はコードを変更するために使用する値を受け取ります。たとえば、stencil_char の引数の 1 つは、比較する正規表現内の文字になります。その文字をマシンコードに直接挿入します。
プロローグは単なるコードのブロックであるため、単純です。
const PROLOGUE_WORDS: 使用量 = 1;
fn stencil_prologue() -> [u32;プロローグ_ワード] {
[0xAA0103E2] // 移動 x2、x1
}
文字比較の場合、比較する文字とフォールバック ロジックのジャンプ先を挿入する必要があります。
const CHAR_WORDS: 使用量 = 4;
fn stencil_char(byte: u8, stencil_pos: 使用、fail_pos: 使用) -> [u32; CHAR_WORDS] {
[
0x39400009, // ldrb w9, [x0]
0x7100013F | ((u32 としてのバイト) << 10), // cmp w9, #byte
0

x54000001 | cond_branch_offset(stencil_pos + 2,fail_pos), // b.ne 失敗
0x91000400, // x0、x0、#1を追加します
】
}
繰り返しの場合、スタックにプッシュするループの開始と、終了へのジャンプがあります。
const SPLIT_WORDS: 使用量 = 5;
fn stencil_split(resume_addr: u64) -> [u32; SPLIT_WORDS] {
[
0xD2800009 | addr_bits(resume_addr, 0), // movz x9, #addr[0..16]
0xF2A00009 | addr_bits(resume_addr, 1), // movk x9, #addr[16..32], lsl 16
0xF2C00009 | addr_bits(resume_addr, 2), // movk x9, #addr[32..48], lsl 32
0xF2E00009 | addr_bits(resume_addr, 3), // movk x9, #addr[48..64], lsl 48
0xA8810029, // stp x9, x0, [x1], #16
】
}
const JMP_WORDS: 使用量 = 1;
fn stencil_jmp(stencil_pos: 使用サイズ, target_pos: 使用サイズ) -> [u32; JMP_WORDS] {
[0x14000000 | Branch_offset(stencil_pos, target_pos)] // b ターゲット
}
そして、非常にきれいな match ブロックと failed ブロックがあります。
const MATCH_WORDS: 使用量 = 4;
fn stencil_match(stencil_pos: 使用サイズ、fail_pos: 使用サイズ) -> [u32; MATCH_WORDS] {
[
0x39400009, // ldrb w9, [x0]
0x35000009 | cond_branch_offset(stencil_pos + 1,fail_pos), // cbnz w9, 失敗
0xD2800020, // mov x0, #1
0xD65F03C0, // リセット
】
}
const FAIL_WORDS: 使用量 = 6;
fn stencil_fail() -> [u32; FAIL_WORDS] {
[
0xEB02003F, // cmp x1, x2
0x54000060, // b.eq +3 (以下の mov へ)
0xA9FF0029, // ldp x9, x0, [x1, #-16]!
0xD61F

[切り捨てられた]

## Original Extract

Build a copy-and-patch JIT compiler in Rust, generate ARM64 machine code, and benchmark it against an interpreter and handwritten implementation.

How AI Changes the Economics of JIT Compilers - malisper.me
Primary Menu
malisper.me
Table of Contents for Postgres Posts
Enter your email address to subscribe to this blog and receive notifications of new posts by email.
Home » pgrust » How AI Changes the Economics of JIT Compilers
How AI Changes the Economics of JIT Compilers
Historically, JIT compilation was a black art. To write a fast JIT compiler, you would need to know how to write assembly. Case in point: there is no production-ready database today that has its own JIT compiler. They all either use LLVM or generate C/C++ code. Both of these options suffer from high compile times, which limits their applicability. Now, with the use of AI, it’s easier than ever to write a JIT compiler with fast compile times by directly targeting assembly. This is also one area of opportunity for new databases to improve on old ones. When building pgrust , I initially thought it would be really hard to implement a JIT compiler. In the end, I found it much easier than I expected due to AI assistance and it ends up being part of the reason why pgrust is so fast. In this post, I’ll walk you through how you can build your own JIT compiler. We’ll build a simple regular expression engine that uses JIT compilation as an example.
JIT compilation is the practice of generating compiled code at runtime or “Just In Time”. When done right, it can result in big performance wins, often on the order of 2-5x and sometimes even more. The main use case for JIT compilation is when there’s information you gain at runtime that drastically alters the behavior of your program. This is particularly common with programming language interpreters; they receive the code to execute at runtime. JIT compilers are also useful in domains beyond programming languages, such as parsing data. Sometimes you don’t know the schema of the data you’re parsing until runtime, and a JIT can help with that.
To kick things off, let’s implement a toy regular expression engine. To keep things simple, we’ll support only two features: literal strings and repetition (i.e. the regex *). We’ll also skip the parser and represent the regular expression as already parsed Rust structures. This means we’ll be able to support strings such as:
but no alternation or lookbehind or anything like that.
In code this is pretty simple. We’ll have 3 types of Nodes: a literal string node, a repetition node, and a concatenation node, which is the combination of two nodes. This ends up looking like this:
enum Node {
Literal(&'static str),
Concatenation(Box<Node>, Box<Node>),
Repetition(Box<Node>),
}
fn literal(text: &'static str) -> Node {
Node::Literal(text)
}
fn concatenation(left: Node, right: Node) -> Node {
Node::Concatenation(Box::new(left), Box::new(right))
}
fn repetition(body: Node) -> Node {
Node::Repetition(Box::new(body))
}
Writing an interpreter for our regular expression engine is also straightforward:
fn match_node(node: &Node, input: &[u8], pos: usize, next: &dyn Fn(usize) -> bool) -> bool {
match node {
Node::Literal(text) => {
let literal = text.as_bytes();
input[pos..].starts_with(literal) && next(pos + literal.len())
}
Node::Concatenation(left, right) => {
match_node(left, input, pos, &|left_end| {
match_node(right, input, left_end, next)
})
}
Node::Repetition(body) => {
match_node(body, input, pos, &|body_end| {
match_node(node, input, body_end, next)
}) || next(pos)
}
}
}
fn interp_match(regex: &Node, input: &str) -> bool {
let bytes = input.as_bytes();
match_node(regex, bytes, 0, &|pos| pos == bytes.len())
}
Now this regular expression engine is pretty simple. It’s under 20 lines of code, but let’s see how it does in terms of performance. For comparison, we’ll compare the code against handwritten code implemented specifically for the regex. For our example we’ll use the regex b(an)*. The handwritten code ends up looking like:
fn handwritten_b_an_star(input: &str) -> bool {
let bytes = input.as_bytes();
let mut pos = 0;
if pos == bytes.len() || bytes[pos] != b'b' {
return false;
}
pos += 1;
while pos < bytes.len() {
if bytes[pos] != b'a' {
return false;
}
pos += 1;
if pos == bytes.len() || bytes[pos] != b'n' {
return false;
}
pos += 1;
}
true
}
(There are ways you could optimize this code and make it much faster, but for our purposes it serves as a good comparison)
When I benchmark a couple of examples against these two, I get that the handwritten version is 10-20x faster than the interpreter. Clearly a lot of room for improvement.
Now let’s take a look at how we can use JIT compilation to get a general regular expression engine that performs as well as the handwritten version.
There are two steps to JIT compile code. First you generate the assembly for the code you want to run. Once you have the code, you then package the assembly code into a function that you can call like any other code into your program.
To generate the assembly, we will use a variant of an approach called copy-and-patch. The idea is that we have a series of templates in assembly for the different operations we want to JIT compile. These templates are called “stencils”. When we want to JIT compile an operation, we take the associated stencil and make small tweaks based on the specifics of the operation. Very similar to filling in a real stencil. By stringing together several of these filled stencils, we can construct a program at runtime that has similar performance to the handwritten version.
Here’s the path we’ll take: first we’ll look at the ARM64 code we want to generate for b(an)*. Then we’ll turn repeated instruction sequences into reusable stencils, write an emitter that fills and combines those stencils from the regex AST, and finally copy the generated instructions into executable memory so Rust can call them like a normal function.
To walk you through how this works, it’s easiest to start with the generated code and work backwards to the JIT compiler itself. Again, we’re working with the regex “b(an)*”. To lay out some design decisions:
We’ll use a stack for backtracking. The stack will keep track of the state we should go to if we hit a dead end in the regex
The string we are matching with will end in a null byte. That means any of our character comparisons will automatically fail if we hit the end of the string. This means we don’t have to do any length comparisons at any point
For the state of our program we will use the following registers:
x0 – current position in string and return value
x1 – top of stack used for backtracking
x2 – bottom of stack used for backtracking (this is needed to determine if the stack is empty)
x9 – used as a temporary variable
For the inputs into our program, we will be passed:
x0 – a pointer to the start of the string
x1 – a pointer to the location we will use for our stack
Now that we’ve taken care of that, let’s walk through the generated assembly part by part. This is specifically on macOS with ARM64. First up, we have the prologue, which initializes the program. All it does is initialize the stack by setting the top of the stack and the bottom of the stack to the value passed in:
0: aa0103e2 mov x2, x1
Next up, we have the code that checks for the character b. If it sees a character that’s not b, we jump to a block of code that handles fallback logic. Otherwise, we advance our position in the string:
; CHAR 'b'
4: 39400009 ldrb w9, [x0] ; load current input byte
8: 7101893f cmp w9, #0x62 ; is it 'b'?
c: 54000281 b.ne 0x5c ; no -> fallback block
10: 91000400 add x0, x0, #1 ; yes -> advance input
Next up, we have the repetition (an)*. For the repetition, we need to do the backtracking. If we backtrack here, that means we jump immediately to the end of the loop. That means we need to store both the address of the instruction after the loop and our position in the string on the stack.
14: d2800989 movz x9, #0x004c ; build resume address
18: f2a00009 movk x9, #0x0000, lsl #16 ; = 0x1_0000_004c
1c: f2c00029 movk x9, #0x0001, lsl #32 ; (the loop exit)
20: f2e00009 movk x9, #0x0000, lsl #48 ;
24: a8810029 stp x9, x0, [x1], #16 ; push (exit, pos) onto stack
With that in place, we can now execute the body of the repetition. This will check for the characters ‘a’ and ‘n’ and, if it sees them, go back to the top of the repetition, but at a new string location.
; CHAR 'a'
28: 39400009 ldrb w9, [x0]
2c: 7101853f cmp w9, #0x61 ; 'a'?
30: 54000161 b.ne 0x5c ; no -> fallback block
34: 91000400 add x0, x0, #1
; CHAR 'n'
38: 39400009 ldrb w9, [x0]
3c: 7101b93f cmp w9, #0x6e ; 'n'?
40: 540000e1 b.ne 0x5c ; no -> fallback block
44: 91000400 add x0, x0, #1
; JMP
48: 17fffff3 b 0x14 ; back to top of loop
Now we’re past the loop. This is where the backtracking will jump once we backtrack. Once we finish the repetition, we’re at the end of the regex. All we have to do now is check if we’re at the end of the string. If we are at the end of the string, we return 1 for success. If we are not, that means the regex failed to match, and we need to run the fail logic to do a fallback.
4c: 39400009 ldrb w9, [x0]
50: 35000069 cbnz w9, 0x5c ; not at NUL -> fallback block
54: d2800020 mov x0, #1 ; success
58: d65f03c0 ret
And then finally, we have the fallback logic. This checks if the stack is empty. If it is, we return 0. If it’s not empty, we pop both the fallback address and the fallback string position off the stack, and then jump to the fallback address.
5c: eb02003f cmp x1, x2 ; any frames left?
60: 54000060 b.eq 0x6c ; no -> give up
64: a9ff0029 ldp x9, x0, [x1, #-16]! ; pop (resume, pos)
68: d61f0120 br x9 ; jump there
6c: d2800000 mov x0, #0 ; no match
70: d65f03c0 ret
Building the Stencils
Now that you’ve had the chance to see the compiled code, you should start to get a sense of how the copy-and-patch compiler would work. We have common sets of instructions with only minor differences between them. For each of these blocks of functions, we can write a function to generate the respective code. Each function will take in values to use to modify the code. For example, one of the arguments to stencil_char will be the char in the regex to compare against. We’ll insert that char directly into the machine code.
The prologue is straightforward since it’s just a block of code:
const PROLOGUE_WORDS: usize = 1;
fn stencil_prologue() -> [u32; PROLOGUE_WORDS] {
[0xAA0103E2] // mov x2, x1
}
For character comparison, we need to insert the character we’re comparing against and where to jump for the fallback logic:
const CHAR_WORDS: usize = 4;
fn stencil_char(byte: u8, stencil_pos: usize, fail_pos: usize) -> [u32; CHAR_WORDS] {
[
0x39400009, // ldrb w9, [x0]
0x7100013F | ((byte as u32) << 10), // cmp w9, #byte
0x54000001 | cond_branch_offset(stencil_pos + 2, fail_pos), // b.ne fail
0x91000400, // add x0, x0, #1
]
}
For the repetition, we have the start of the loop that pushes onto the stack and the jump onto the end:
const SPLIT_WORDS: usize = 5;
fn stencil_split(resume_addr: u64) -> [u32; SPLIT_WORDS] {
[
0xD2800009 | addr_bits(resume_addr, 0), // movz x9, #addr[0..16]
0xF2A00009 | addr_bits(resume_addr, 1), // movk x9, #addr[16..32], lsl 16
0xF2C00009 | addr_bits(resume_addr, 2), // movk x9, #addr[32..48], lsl 32
0xF2E00009 | addr_bits(resume_addr, 3), // movk x9, #addr[48..64], lsl 48
0xA8810029, // stp x9, x0, [x1], #16
]
}
const JMP_WORDS: usize = 1;
fn stencil_jmp(stencil_pos: usize, target_pos: usize) -> [u32; JMP_WORDS] {
[0x14000000 | branch_offset(stencil_pos, target_pos)] // b target
}
And then we have the match and fail blocks which are pretty clean:
const MATCH_WORDS: usize = 4;
fn stencil_match(stencil_pos: usize, fail_pos: usize) -> [u32; MATCH_WORDS] {
[
0x39400009, // ldrb w9, [x0]
0x35000009 | cond_branch_offset(stencil_pos + 1, fail_pos), // cbnz w9, fail
0xD2800020, // mov x0, #1
0xD65F03C0, // ret
]
}
const FAIL_WORDS: usize = 6;
fn stencil_fail() -> [u32; FAIL_WORDS] {
[
0xEB02003F, // cmp x1, x2
0x54000060, // b.eq +3 (to the mov below)
0xA9FF0029, // ldp x9, x0, [x1, #-16]!
0xD61F

[truncated]
