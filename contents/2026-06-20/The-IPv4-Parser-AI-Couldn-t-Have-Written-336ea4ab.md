---
source: "https://extractingcycles.com/blog/the-ipv4-parser-ai-couldnt-have-written/article/"
hn_url: "https://news.ycombinator.com/item?id=48608826"
title: "The IPv4 Parser AI Couldn't Have Written"
article_title: "Extracting Cycles"
author: "csno"
captured_at: "2026-06-20T15:38:42Z"
capture_tool: "hn-digest"
hn_id: 48608826
score: 6
comments: 0
posted_at: "2026-06-20T12:43:18Z"
tags:
  - hacker-news
  - translated
---

# The IPv4 Parser AI Couldn't Have Written

- HN: [48608826](https://news.ycombinator.com/item?id=48608826)
- Source: [extractingcycles.com](https://extractingcycles.com/blog/the-ipv4-parser-ai-couldnt-have-written/article/)
- Score: 6
- Comments: 0
- Posted: 2026-06-20T12:43:18Z

## Translation

タイトル: AI が書けなかった IPv4 パーサー
記事のタイトル: サイクルの抽出

記事本文:
キャッシュに優しい読み取り、ユーザーに敵対的な書き込み。
IPv4 パーサー AI は書けなかった
私は Daniel Lemire のブログを読んでいました。これはとてもお勧めです。具体的には: IP アドレスを迅速に解析します (SIMD マジックを使用せずに移植可能)。
タイトルを読んで、彼のこれまでの実績を考えると、私はある種の狂気の SWAR (SIMD Within A Register) 解析テクニックを使用するのではないかと身構えました。おそらく、私が 12 回読まないと理解できないような難解なビットいじりが組み合わさったものかもしれません。
驚いたことに、Lemire は主に AI によって書かれたリニア スキャン関数を紹介しました。
まだ引退すべきではないという彼の結論には同意しますが、高速なコードを生成するために AI に完全に依存することは避けるべきだと思います。
私の主張に入る前に、まず SWAR とちょっとしたトリックに対する飢えを満たしましょう。
このような問題に取り組むときに私が最初に行うことは、有効な入力の感覚をつかむことです。 IPv4 について考えるとき、すぐに 3 つの例が思い浮かびます。
192.168.1.1
255.255.255.255
0.0.0.0
話を簡単にするために、ドット付き 10 進数バージョンのみを扱います。
そこから、いくつかのことに気づき始めます。
各トリプレットの有効な入力範囲は [0, 255] です。
入力長は 7 ～ 15 文字の範囲です。
各文字列には 3 つのドット「.」が必要です。
入力文字列は常に 16 バイトに収まるため、入力文字列を 2 つの u64 に分割し、最初の u64 にバイト [0..8] をロードし ( head と呼ぶことにします)、2 番目の u64 にバイト [len - 8..len] をロードし ( tail と呼ぶことにします) ます。
残念ながら、 address.len == 7 では、2 つの問題が発生します。最初の読み取りでは、入力長を超えて 1 バイトにアクセスします (バイト 8 まで読み取ろうとするため)。もう 1 つは、ポインター計算を考慮して、2 回目の読み取りでは、入力文字列の開始前の 1 バイトにアクセスすることです。
幸いなことに、最初の問題は

これは許容できる癖であり、これを API コントラクトに文書化するだけで済みます。ほとんどの場合、文字列は null で終了するか、より大きなバッファーの一部になるためです。どちらの場合でも、安全な 8 バイトの読み取りが得られます。
2 番目の問題に関しては、入力の長さが 7 の場合に、オーバーフローしたり、文字列の開始前に読み取られたりしないように、減算を調整する必要があります。
そうすると、2 つの読み取り値は次のようになります。
var head : u64 = 未定義 ;
@memcpy(std.mem.asBytes(&head), アドレス [0..8]) ;
var tail : u64 = 未定義 ;
const len_is_7 = アドレス。レン == 7 ;
const begin = if (len_is_7) 0 else アドレス。レン - 8 ;
const 終了 = 開始 + 8 ;
@memcpy(std.mem.asBytes(&tail),アドレス[begin..end]);
tail <<= if ( len_is_7 ) 8 else 0 ; // それについては後で詳しく説明します。
注: まともなコンパイラは、これらの if の分岐を生成しません。
IP アドレス「192.168.1.1」で何が起こるかを見てみましょう。
低い = '192.168'。
高 = '.168.1.1'
low 変数は SWAR 整数解析の優れた候補のように見えます。
これは、Lemire 氏のブログでよく説明されているテクニックです: リンク 。基本的には、専用の SIMD レジスタではなく、通常の汎用レジスタ内で並列処理を実行します。
通常、「192」のような文字列を整数表現に解析する場合、各バイトを反復処理して、それが 0x30 ( '0' ) から 0x39 ( '9' ) の有効な ASCII 範囲内にあるかどうかを確認します。有効な文字ごとに、ASCII オフセットを減算し、その結果に位置の値 (10 の累乗) を乗算し、それを累計に加算します。反復が完了すると、合計が最終的な整数値を保持します。
この反復作業はすべて、次のスニペットを使用してブランチレスで実行できます。
注: Zig では、% が続く算術演算子は、オーバーフロー (+% 加算を使用した) を許可する明示的な方法です。

h オーバーフロー)。
// リトルエンディアン CPU では、メモリは右から左にレジスタにロードされます。
// メモリレイアウト: ['1', '9', '2'] -> 0x31, 0x39, 0x32
定数入力 = 0x32_39_31 ;
// 3 バイトすべてから同時に ASCII オフセット (0x30) を減算します。
// 単一のラッピング減算 (-%) を使用します。
const d = 入力 -% 0x30_30_30 ; // = 0x32_39_31 - 0x30_30_30 = 0x02_09_01
// 次に、これらの値に位置の値を乗算する必要があります。
// 1 * 100、9 * 10、2 * 1 -> * 0x_64_0A_01
定数値 = d *% 0x640A01 ;
結果は 3 番目のレーンにあり、 (value & 0x00FF0000) >> 16 で簡単に抽出できます。ただし、smarty-pants をプレイして、BMI2 拡張機能で導入された専用の x86_64 命令 pext を使用できます。
pext は入力とマスクの 2 つのオペランドを受け入れ、基本的にマスク ビットが 1 である入力ビットを右端のスポットにコピーします。
これを使用すると、抽出は単純に次のようになります。 pext(value, 0xFF0000) ;このようなあいまいな操作を行う理由は、後ほど明らかになるでしょう。
この関数は、「999」などの無効な 3 文字や数字以外の文字も含めてすべてを解析しようとするため、検証ステップが必要です…
通常、SWAR 検証は、処理された文字が 0x00 と 0x09 の間に適切に収まるかどうかを確認する範囲チェックです。各レーンの最上位ビット 0x80 を利用することで、ループを行わずにこれを実現できます。 0x09 のような有効な数字に 0x76 を追加すると、それは 0x7F になり、その最上位ビットは設定されないままになります。しかし、値が無効な 0x0A 以上に達した瞬間、 0x76 を追加すると値が 0x80 以上にプッシュされ、すぐに最上位ビットが 1 に反転されます。
落とし穴があります。入力レーンに最初から最上位ビットが設定されている場合 (非 ASCII 文字など)、この追加によりオーバーフローが発生し、ビットがラップアラウンドしてクリアされ、エラーが隠蔽される可能性があります。防ぐために

ここでは、ビットごとの OR を使用して、元の減算レジスタ d と加算結果を結合します。これにより、加算の前または後に最上位ビットが設定された場合、そのビットは設定されたままになります。最後に、0x808080 でマスキングすると、これらのビットが 3 バイトすべてで同時に分離されます。結果がゼロの場合、すべての文字は有効な数字です。
const err = (( d | ( d +% 0x767676 ) ) & 0x808080 ) == 0 ;
ただし、私たちのケースは、IPv4 の厳密な上限である 255 を強制する必要があるため、少し複雑になります。各桁の最大許容制限は、その左側の桁に完全に依存します。
これらのレーン間の依存関係の概略図は次のとおりです。
百人レーン:
└── 指定できる範囲: [0, 2]
テンズレーン:
§── Hundreds が 0 または 1 の場合、許容範囲: [0, 9]
└── 百の数が 2 の場合、許容範囲: [0, 5]
ワンズレーン:
§── 100 の位 == 2 かつ 10 の位 == 5 の場合、許容範囲: [0, 5]
└── それ以外の場合、許容範囲: [0, 9]
検証マスクは次のようになります。
0x76767D = いいえ 2
0x767A7D = 2 は存在しますが、5 は存在しません
0x7A7A7D = 2 と 5 の両方が存在します
これらのカスケード境界を分岐せずにマップするには、数字から直接正確な 0x04 ビット調整を計算することで、ベース マスク ( 0x76767D ) を動的にモーフィングできます。 「100」の位が "2" の場合、「10」レーンに 0x04 を追加して、その最大境界を "9" から "5" にシフトします (0x76 を 0x7A に変換します)。両方の「hundreds」が「2」で「tens」が「5」の場合、「ones」レーンに 0x04 を追加します。
const ベースマスク = 0x76767D ;
const get_25 = (d +% 0x7B7E ) & 0x8080 ;
// 100 のレーンと 10 のレーンを比較します。 200ビットが設定されている場合
// 50 ビットを設定したままにすることができます (オンの場合)
// 最終結果をシフトして、0x04 を正しい位置に配置します
const 調整 = (get_25 & ((get_25 << 8 ) | 0x80 ) ) << 3 ;
const err = ( d | ( d +% 0

x76767D +% 調整 ) ) & 0x808080 ;
同様に、ビット操作と私の個人的な精神的健康という安価なコストで、入力全体を検証することができます。
u64 レジスタがあるため、一度に 2 つの入力トリプレットを解析できます。プロセスは同じで、ドット文字のチェックを追加するのは非常に簡単です。最終結果は次のとおりです。
inline fn parseTwoTripletsLow (チャンク : u64 、 active_lanes : u64 ) ParseResult {
// 0x2E は「.」を表します。検証
const d = チャンク -% (0x2E3030302E303030 & active_lanes ) ; // 詳細については後ほど
const get_25 = (d +% 0x00007B7E00007B7E) & 0x0000808000008080 ;
const 調整 = (get_25 & ((get_25 << 8 ) | 0x0000008000000080 ) ) << 3 ;
const err = (d | (d +% 0x7F76767D7F76767D +% 調整) ) & 0x8080808080808080 ;
const コンパクト = d *% 0x640A01 ;
戻る 。 { 。 value = @intCast(pext(compact, 0x00FF000000FF0000)) , .エラー = エラー } ;
}
注: これで、pex が意味を持ちます。これにより、2 つの AND、2 つのシフト、および 1 つの OR を実行する必要がなくなります。
注: わずかに異なる定数を持つ別の関数を、high 変数の解析に使用できます。
次号が1マイル離れたところから来るのが見えました。私たちが解析ルーチンを開発したとき、たまたま「192.168」としてきちんと構造化されていた低位変数のみを考慮しました。 。しかし、変数が高いことが示すように、私たちの入力は大きく異なります。
高 = '.168.1.1'
考えられるすべての入力を parseTwoTriplets が期待するレイアウトに再フォーマットする機械が必要です。この大きな変数の場合、ターゲット レイアウトは次のようになります。
チャンク = ".1.1"
チャンクは本質的に高く、そのバイトは特定の位置に分散されています。それは私たちの古い友人のpextの逆です!幸いなことに、ハードウェアの神様はまさにそれを提供してくれました: pdep を満たすことです。
pdep は入力値とマスクを受け取ります。入力のビットを最下位から最上位まで読み取り、それぞれを

e をマスク内の次の設定ビットの位置に、同様に最下位から最上位の順に挿入します。言い換えれば、マスクは出力ビットがどこに到達するかを記述します。
この例では、 pdep(high >> 40, 0xFF0000FFFF0000FF) を呼び出します。 pdep は実際には右から左に読み取って (ええと、リトル エンディアン)、上記の ".1.1" レイアウトを生成するため、開始時に興味深い部分を取得するにはシフトが必要です。マスクは、どのバイト レーンにデータを入力し、どのバイト レーンをゼロ設定のままにするかをエンコードします。非アクティブなレーンにゼロ以外の値があると、レーン - 0x30 の減算により有効性チェックがトリップされる可能性があるため、ゼロ化されたレーンは重要です。
残りの問題は、特定の入力のドット構成に適した pdep マスクをどのように見つけるかということです。ハッシュマップしてみます!!
そのテーブルにインデックスを構築するには、入力内のドットを見つける必要があります。まず、入力を 0x2E2E2E2E2E2E2E2E と XOR 演算します (0x2E は '.' の ASCII コードです)。 XOR 恒等式 x ^ x == 0 のおかげで、ドットを含むレーンは 0x00 になります。次に、結果から 0x0101010101010101 を減算します。0x00 だったレーンは上のレーンから借用し、その位置に 0xFF を生成します。最後に、0x8080808080808080 との AND 演算により上位ビットのみが分離され、ドット位置のきれいなビットマスクが得られます。
次に、このドット位置マスクにマジック定数を乗算して、32 エントリのルックアップ テーブルの一意のスロットにマップします (すべての有効なドット構成をカバーするには 32 スロットだけが必要です)。マップ内で 2 つの有効なドット マスクが衝突しないことを保証する必要があるため、定数は慎重にブルート フォースで存在させる必要があります。:P
var Index_high = テール ^ 0x2E2E2E2E2E2E2E2E ;
インデックス高 -%= 0x0101010101010101 ;
インデックス高 &= 0x8080808080808080 ;
インデックス高 *%= 0x39EDDB77A53AC365 ; // 魔法の定数
インデックス高 >>= 59 ; // 有効なすべての CO には 32 スロットで十分です

構成
注: 無効な文字によって、有効に見えるインデックスが生成される場合があります。 parseTwoTriplets の既存のチェックでそれらのケースをキャッチできるようにするだけです。
インデックスを使用して、レーン レイアウトの pdep マスクと、重要なバイトを分散する前に右揃えするためのシフト量の両方を検索します。 pdep マスクはどのレーンがアクティブであるかをすでに正確に記述しているため、同時に parseTwoTriplets 減算マスクを使用して AND を実行し、チェックに引っかからないように非アクティブなレーンを無効にすることができます。
const high_pdep = lut_high 。 pdep [インデックス高] ;
const high_shift = lut_high 。シフト [index_high ] ;
const high = parseTwoTripletsHigh (pdep (tail >> @intCast (high_shift) , high_pdep ) , high_pdep ) ;
今何をすればいいでしょうか？
さて、私たちは今、両親が私を愛してくれているように、自分の仕事を愛し始める必要があります。条件付きで、そして私が完璧にパフォーマンスをした場合にのみです。
自分たちの仕事を信頼してはいけないので、本当に優れたテストスイートを構築する必要があります。これは投稿の範囲外ですが、これは本当に拷問テストであるはずです。
その理由は、たとえ最高のアイデアがあっても、エッジケースを見逃してしまうからです。たとえば、読み込みルーチンは 255.255.255 を 255.255.255.255 と同じように扱い、有効な入力とします:)
簡単な修正は、address.len - 3 == lut_high[high_index].active_lanes を確認することです。

[切り捨てられた]

## Original Extract

Cache-friendly reads, user-hostile writes.
The IPv4 Parser AI Couldn't Have Written
I was reading Daniel Lemire’s blog, which I highly recommend; specifically: Parsing IP addresses quickly (portably, without SIMD magic) .
Reading the title and given his track record, I braced myself for some sort of insane SWAR (SIMD Within A Register) parsing technique, maybe combined with obscure bit twiddling that I’m too stupid to understand without reading it 12 times, or something along those lines.
To my surprise, Lemire showcased a linear scan function mostly written by AI.
Although I agree with his conclusion that he shouldn’t retire just yet, I think we should avoid relying on AI to generate fast code altogether!!
Before diving into my claim, let’s firstly satisfy our hunger for SWAR and bit tricks!
The first thing I do when I approach such problems is to get a feel for the valid inputs. When I think about IPv4, three examples quickly come to mind:
192.168.1.1
255.255.255.255
0.0.0.0
To keep things simple, we are going to treat only the dotted decimal version!
From there, we start noticing some things:
The valid input range for each triplet is [0, 255].
The input length ranges from 7 to 15 characters.
Each string must have 3 dots, ‘.’
Since the input string will always fit into 16 bytes, we can smash it into 2 u64 , loading bytes [0..8] in the first u64 , which we will call head , and bytes [len - 8..len] in the second u64 , which we will refer to as tail .
Unfortunately, with an address.len == 7 , we will encounter two problems: the first read will access 1 byte over the input length (since it will try to read until byte 8), and the second read will access 1 byte before the start of our input string, given pointer math.
Luckily, the first problem is an acceptable quirk, and we can simply document it in the API contract, since almost always, strings will be null-terminated or part of a larger buffer, and in both cases, we get a safe 8-byte read!
As for the second problem, we have to adjust our subtraction such that, in case the input length is 7, we don’t overflow or read before the start of the string.
Doing so, the two reads become:
var head : u64 = undefined ;
@memcpy ( std . mem . asBytes ( & head ) , address [ 0 .. 8 ] ) ;
var tail : u64 = undefined ;
const len_is_7 = address . len == 7 ;
const begin = if ( len_is_7 ) 0 else address . len - 8 ;
const end = begin + 8 ;
@memcpy ( std . mem . asBytes ( & tail ) , address [ begin .. end ] ) ;
tail <<= if ( len_is_7 ) 8 else 0 ; // More on that one later!
Note: Every decent compiler will not emit branches for those ifs!
Let’s now have a look at what happens with IP address "192.168.1.1" :
low = '192.168.'
high = '.168.1.1'
The low variable looks like a great candidate for SWAR integer parsing!!
It’s a technique that Lemire discusses a lot on his blog: link . It’s basically doing parallel operations inside a regular general-purpose register instead of a dedicated SIMD one!
Typically, when parsing a string like "192" into its integer representation, we iterate through each byte to verify it falls within the valid ASCII range of 0x30 ( '0' ) to 0x39 ( '9' ). For each valid character, we subtract the ASCII offset, multiply the result by its positional value (its power of 10), and add it to a running total. Once the iteration is complete, the total holds the final integer value.
All of that iterative work can be done in a branchless manner with the following snippet:
Note: In Zig, math operators followed by % are an explicit way to allow overflow (+% addition with overflow).
// On a Little-Endian CPU, memory is loaded right-to-left into registers.
// Memory layout: ['1', '9', '2'] -> 0x31, 0x39, 0x32
const input = 0x32_39_31 ;
// Subtract the ASCII offset (0x30) from all 3 bytes simultaneously
// using a single wrapping subtraction (-%).
const d = input -% 0x30_30_30 ; // = 0x32_39_31 - 0x30_30_30 = 0x02_09_01
// Now we need to multiply those values by their positional value
// 1 * 100, 9 * 10, 2 * 1 -> * 0x_64_0A_01
const value = d *% 0x640A01 ;
Our result now resides in the third lane and could easily be extracted with (value & 0x00FF0000) >> 16 . However, we can play smarty-pants and use a purpose-built x86_64 instruction introduced with the BMI2 extension: pext .
pext will accept two operands, an input and a mask, and will basically copy the input bits where the mask bits are 1 to the rightmost spot.
Using this, our extraction simply becomes: pext(value, 0xFF0000) ; the reason to involve such an obscure operation will be clearer later!
Since this function will try to parse everything, invalid triplets like "999" or even non-digit characters, we need a validation step…
Usually, SWAR validation is simply a range check to see if our processed characters fall nicely between 0x00 and 0x09 . We can achieve this without looping by exploiting the most significant bit, 0x80 , in each lane. If we add 0x76 to a valid digit like 0x09 , it becomes 0x7F , leaving that highest bit unset. But the moment a value hits an invalid 0x0A or higher, adding 0x76 pushes it to 0x80 or above, immediately flipping the highest bit to 1 .
There is a catch: if an input lane already had its highest bit set from the start (like a non-ASCII character), this addition might cause an overflow that wraps around and clears the bit, hiding the error. To prevent this, we combine the original subtracted register d with the addition result using a bitwise OR. This ensures that if the highest bit was set either before or after the addition, it stays set. Finally, masking with 0x808080 isolates these bits across all three bytes simultaneously. If the result is zero, every single character is a valid digit.
const err = ( ( d | ( d +% 0x767676 ) ) & 0x808080 ) == 0 ;
Our case, however, is a bit more convoluted because we need to enforce the strict IPv4 upper bound of 255. The maximum allowed limit for each digit depends entirely on the digits to its left.
Here is the schematic breakdown of these cross-lane dependencies:
Hundreds Lane:
└── Allowed range: [0, 2]
Tens Lane:
├── If Hundreds is 0 or 1, Allowed range: [0, 9]
└── If Hundreds is 2, Allowed range: [0, 5]
Ones Lane:
├── If Hundreds == 2 AND Tens == 5, Allowed range: [0, 5]
└── Otherwise, Allowed range: [0, 9]
Our validation masks then are:
0x76767D = no 2
0x767A7D = 2 present but 5 is not
0x7A7A7D = 2 and 5 both present
To map these cascading boundaries without branching, we can dynamically morph the base mask ( 0x76767D ) by computing precise 0x04 bit adjustments directly from the digits. If the ‘hundreds’ digit is "2" , we add 0x04 to the ‘tens’ lane to shift its max bound from "9" to "5" (turning 0x76 into 0x7A ). If both ‘hundreds’ is "2" and ‘tens’ is "5" , we add 0x04 to the ‘ones’ lane.
const base_mask = 0x76767D ;
const get_25 = ( d +% 0x7B7E ) & 0x8080 ;
// Compare the hundreds lane with the tens lane; if the 200 bit is set
// then the 50 bit can be left set (if it is on)
// Shift the final result to get 0x04 in the right position
const adjust = ( get_25 & ( ( get_25 << 8 ) | 0x80 ) ) << 3 ;
const err = ( d | ( d +% 0x76767D +% adjust ) ) & 0x808080 ;
Like so, we have the whole input validated at the cheap cost of some bit manipulation and my personal mental well-being.
Since we have u64 registers, we could parse two input triplets at a time; the process is identical, and the check for the dot character is super simple to add. The final result is:
inline fn parseTwoTripletsLow ( chunk : u64 , active_lanes : u64 ) ParseResult {
// 0x2E is for the '.' validation
const d = chunk -% ( 0x2E3030302E303030 & active_lanes ) ; // More on that & later
const get_25 = ( d +% 0x00007B7E00007B7E ) & 0x0000808000008080 ;
const adjust = ( get_25 & ( ( get_25 << 8 ) | 0x0000008000000080 ) ) << 3 ;
const err = ( d | ( d +% 0x7F76767D7F76767D +% adjust ) ) & 0x8080808080808080 ;
const compact = d *% 0x640A01 ;
return . { . value = @intCast ( pext ( compact , 0x00FF000000FF0000 ) ) , . err = err } ;
}
Note: Now the pext makes sense; it avoids us doing 2 ANDs, 2 shifts, and 1 OR.
Note: Another function with slightly different constants can be used for parsing the high variable.
We saw the next issue coming from a mile away. When we developed the parsing routine, we only considered the low variable, which happened to be neatly structured as "192.168." . But our inputs vary wildly, as evidenced by the high variable:
high = '.168.1.1'
We need machinery to reformat every possible input into the layout that parseTwoTriplets expects. For this high variable, the target layout would be:
chunk = ". 1. 1"
The chunk is essentially high with its bytes scattered into specific positions; it’s the inverse of our old friend pext ! Fortunately, the hardware gods have provided exactly that: meet pdep .
pdep takes an input value and a mask. It reads the bits of the input from least to most significant and places each one into the position of the next set bit in the mask, also from least to most significant. In other words, the mask describes where the output bits should land.
In our case, we call pdep(high >> 40, 0xFF0000FFFF0000FF) . The shift is necessary to get the interesting part at the start, since pdep will actually read right-to-left (eheh Little-Endian) to produce the ". 1. 1" layout above. The mask encodes which byte lanes should be populated and which should remain zeroed out. The zeroed lanes are important because a non-zero value in an inactive lane would cause the lane - 0x30 subtraction to trip our validity checks.
The remaining problem is: how do we find the right pdep mask for a given input’s dot configuration? We hash-map it!!
To build the index into that table, we need to locate the dots in the input. We start by XORing the input against 0x2E2E2E2E2E2E2E2E (where 0x2E is the ASCII code for '.' ). Thanks to the XOR identity x ^ x == 0 , lanes that contain a dot become 0x00 . We then subtract 0x0101010101010101 from the result: any lane that was 0x00 will borrow from the lane above, producing 0xFF in that position. Finally, ANDing with 0x8080808080808080 isolates only those high bits, giving us a clean bitmask of the dot positions.
This dot-position mask is then multiplied by a magic constant to map it to a unique slot in a 32-entry lookup table (we only need 32 slots to cover all valid dot configurations). We need to guarantee that no two valid dot masks collide in the map, so a constant needs to be carefully brute-forced into existence :P
var index_high = tail ^ 0x2E2E2E2E2E2E2E2E ;
index_high -%= 0x0101010101010101 ;
index_high &= 0x8080808080808080 ;
index_high *%= 0x39EDDB77A53AC365 ; // Magic constant
index_high >>= 59 ; // 32 slots are enough for all valid configurations
Note: Invalid characters may occasionally produce a valid-looking index. We simply let parseTwoTriplets ’s existing checks catch those cases.
With the index in hand, we look up both the pdep mask for the lane layout and a shift amount to right-align the significant bytes before scattering them. Since the pdep masks already describe exactly which lanes are active, we can AND them with the parseTwoTriplets subtraction mask at the same time, disabling the inactive lanes so they don’t trip our checks.
const high_pdep = lut_high . pdep [ index_high ] ;
const high_shift = lut_high . shift [ index_high ] ;
const high = parseTwoTripletsHigh ( pdep ( tail >> @intCast ( high_shift ) , high_pdep ) , high_pdep ) ;
What do we do now?
Well, now we need to start loving our work the way my parents love me: conditionally, and only if I perform flawlessly.
We need to build a really really good test suite because we shouldn’t trust our work; this is outside the scope of the post, but it should really be a torture test!!
The reason is that even with the best ideas we will miss edge cases. For example, our loading routine treats 255.255.255 the same as 255.255.255.255 , making it a valid input :)
A simple fix is to check that address.len - 3 == lut_high[high_index].active_lanes

[truncated]
