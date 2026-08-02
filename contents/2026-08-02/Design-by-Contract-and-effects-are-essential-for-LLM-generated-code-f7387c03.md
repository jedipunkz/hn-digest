---
source: "https://gavinray97.github.io/blog/design-by-contract-and-effects-for-llms"
hn_url: "https://news.ycombinator.com/item?id=49148299"
title: "Design by Contract and effects are essential for LLM-generated code"
article_title: "Design by Contract and effects are essential for LLM-generated code"
author: "gavinray"
captured_at: "2026-08-02T21:44:51Z"
capture_tool: "hn-digest"
hn_id: 49148299
score: 1
comments: 0
posted_at: "2026-08-02T21:05:18Z"
tags:
  - hacker-news
  - translated
---

# Design by Contract and effects are essential for LLM-generated code

- HN: [49148299](https://news.ycombinator.com/item?id=49148299)
- Source: [gavinray97.github.io](https://gavinray97.github.io/blog/design-by-contract-and-effects-for-llms)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T21:05:18Z

## Translation

タイトル: LLM で生成されたコードには契約による設計とエフェクトが不可欠です
説明: LLM がより多くのコードを作成するにつれて、コントラクトとエフェクト システムは、動作の制約、副作用の公開、実装の検証に不可欠なツールになります。

記事本文:
LLM で生成されたコードには Design by Contract とエフェクトが不可欠です ブログ タグ プロジェクトについて ブログ タグ プロジェクトについて 公開日: 2026 年 8 月 2 日 日曜日 Design by Contract とエフェクトは LLM で生成されたコードに不可欠です
名前 ギャビン・レイ Twitter @GavinRayDev
この投稿では、ソフトウェア開発が人間によるコードから LLM によるコードに移行するにつれて、大幅に重要になったと私が考える 2 つの言語機能について説明したいと思います。
他の多くの人と同様に、私が昨年「書いて」（プロンプトを出して）出荷したコードの大部分は、私が作成したものではありません。 LLM に仕様を入力すると、LLM は実装を吐き出し、動作を検証するテストを作成するように要求します。テストを実行し、アプリを手動で実行して期待どおりに動作するかどうかを確認し、時間があればコードを確認します。
この種の開発ループは、「約束どおりに動作するコード」に計り知れない価値をもたらします。
これを促進するプログラミング言語の機能が 2 つあり、機械生成コードのこの新時代ではそれらの価値が 10 倍になると私は確信しています。
これらの最終的な結果は、次のような PR 内のセマンティック変更に関するコンパイラ生成レポートが得られる可能性です。
追加されたエフェクト: PaymentProcessor.process
+ net.connect
+ 再試行.非決定的
事後条件の弱体化: Ledger.append
- ledger.length == old(ledger.length) + 1 であることを保証します。
AI (不) 使用に関する免責事項: 散文のいかなる部分も機械によって生成されたものではありません。このブログには機械で書かれた散文はありません。それは大変失礼なことだと思います。
Design-by-Contract (DbC) は、メソッドおよび構造体/クラスの署名の一部として、前提条件、事後条件、および不変式 (常に真) を記述することを可能にする言語/構文機能です。
これは、アドホックなテストと正式な仕様の組み合わせの中間に位置します。多くの言語で

つまり、コントラクトには、パフォーマンス上の理由から、コンパイル時から実行時まで動作とチェックを切り替えるビルド時の「レベル」スイッチがあります。
百聞は一見に如かずなので、契約について説明し続けるのではなく、（できれば）わかりやすい例をいくつか挙げてみましょう。
構造体ユーザー {
電子メール: 文字列
email_verified : ブール値
}
この状態を妨げるものは何もありません。
ユーザー {
電子メール: "" 、
email_verified : true
}
しかし、契約による設計では、いくつかの単純な不変条件をエンコードして、それを排除できます。
構造体ユーザー {
電子メール: 文字列
email_verified : ブール値
不変の電子メール。 is_valid_email ()
不変の email_verified は ! を意味します。メールでご連絡ください。 is_empty ( )
}
電子メールを変更すると、他に何を変更する必要があるかを指定できます。
fn change_email ( user : & mut User , new_email : String )
new_email が必要です。 is_valid_email ()
ユーザーを保証します。電子メール == 新しい電子メール
ユーザーを保証します。 email_verified == false
事後条件が保証されていない場合、生成された実装は、email_verified を true に設定したままアドレスを更新する可能性があります。
プロパティを不変条件としてエンコードできるデータ構造は、この種の設計に特に適しています。
struct MinHeap < T : 順序付け > {
項目 : 配列 < T >
invariant forall i in 1 .. < items .長さ :
アイテム [ 親 ( i ) ] <= アイテム [ i ]
}
struct BTreeNode < K , V , const ORDER : use > {
キー : 配列 < K >
値: 配列 < V >
子: 配列 < NodeRef >
is_leaf : ブール値
不変キー 。長さ == 値 。長さ
不変キー 。長さ <= ORDER - 1
不変厳密_増加 (キー)
不変式 is_leaf
子供を意味します。長さ == 0
不変です！ is_leaf
子供を意味します。長さ == キー 。長さ + 1
}
効果
エフェクトは、メソッド内の明示的な機能を示す方法です。
通常、これらは「ファイル/ネットワーク IO、メモリ割り当て、状態変更」などの動作です。

メソッドの効果が明示的であることが必要な場合は、メソッドの署名を、その許可された動作の検証済み契約として使用できます。
一例として、次のメソッドがあるとします。
fn load_settings ( path : パス ) -> 設定
戻り値の型には、設定がどこから来たのか、操作で何が行われるのかについては何も示されていません。
効果を認識したシグネチャは次のことを行います。
fn load_settings ( path : パス ) -> 設定
パスが必要です。存在します ( )
パスが必要です。 is_file ( )
結果を保証します。有効です ( )
効果 {
fs 。読み取り (パス) 、
割り当てる
}
次に、これをメモリにすでに保持されている設定を単に解析する関数と比較してください。
fn parse_settings (内容:文字列) -> 設定
が必要です！コンテンツ 。 is_empty ( )
結果を保証します。有効です ( )
効果 {
割り当てる
}
これらの関数は同じ TYPE を返す可能性がありますが、OPERATIONAL 動作は大幅に異なります。その違いは通話現場で目に見えるはずです。
最新の内部メタデータを更新するキャッシュ取得
一見すると、以下のメソッドは読み取り専用であると思われるかもしれません。
fn get (cache : & mut Cache < K , V > , key : K ) -> オプション < V >
ただし、多くのキャッシュは使用状況の統計を保持しているため、get メソッドが変化している場合もあります。以下のシグネチャは、次のような動作について明示的に示しています。
fn get (cache : & mut Cache < K , V > , key : K ) -> オプション < V >
結果を保証します。 is_some ( ) == old ( キャッシュ . を含む ( キー ) )
キャッシュを確保します。エントリ == 古い (キャッシュ . エントリ)
効果 {
状態も書き込み (キャッシュの最新性)
}
キャッシュの内容は変更されませんが、エビクション メタデータは変更されます。
割り当ててはいけない関数
呼び出し元が提供するバッファーに値を書き込む/エンコードするメソッドがあるとします。
fn エンコード (
値: メッセージ、
宛先: & mut ByteBuffer
) -> 使用する
宛先が必要です。残り >= 値。エンコードされたサイズ ( )
結果 == va を保証します

ルー。エンコードされたサイズ ( )
目的地を確保します。位置 == 古い (目的地の位置) + 結果
効果 {
記憶に残る書き込み（宛先）
}
この実装は、提供されたバッファに書き込むことはできますが、一時的な配列または文字列を作成することはできません (alloc 効果がないため)。

## Original Extract

As LLMs write more code, contracts and effect systems become essential tools for constraining behavior, exposing side effects, and verifying implementations.

Design by Contract and effects are essential for LLM-generated code Blog Tags Projects About Blog Tags Projects About Published on Sunday, August 2, 2026 Design by Contract and effects are essential for LLM-generated code
Name Gavin Ray Twitter @GavinRayDev
In this post, I want to discuss two language features that I think have become substantially more important as software development shifts from human-authored to LLM-authored code.
Like many others, the majority of the code I have "written" (prompted) and shipped in the last year was not authored by me. You feed specifications to an LLM, it spits out an implementation, and you ask it to write tests to verify the behavior. You run the tests, manually poke the app to see if it behaves how you expect, and if you have the time you review the code.
This sort of development loop gives immense value to "code that verifiably does what it says on the tin."
There are two programming language features that facilitate this, and I'm convinced that their value is tenfold in this new era of machine-generated code:
The net result of these is the possibility to have compiler-generated reports of semantic changes in a PR, like:
Effects added : PaymentProcessor.process
+ net.connect
+ retry.nondeterministic
Postcondition weakened : Ledger.append
- ensures ledger.length == old(ledger.length) + 1
AI (Dis)use Disclaimer: No part of the prose was machine-generated. You will not find machine-written prose on this blog. I consider it deeply disrespectful.
Design-by-Contract (DbC) is language/syntax feature that allows writing preconditions, postconditions, and invariants (always-true) as part of the signature of methods and structs/classes.
It's somewhere between a mix of ad-hoc testing and formal specification. In many languages, contracts have build-time "level" switches which toggle the behavior and checks from compile-time to run-time, for performance reasons.
A picture is worth a thousand words, so rather than continue describing Contracts, let me give some (hopefully) self-explanatory examples.
struct User {
email : String
email_verified : Bool
}
Nothing prevents this state:
User {
email : "" ,
email_verified : true
}
But with Design-by-Contract, we can encode a few simple invariants to rule it out:
struct User {
email : String
email_verified : Bool
invariant email . is_valid_email ( )
invariant email_verified implies ! email . is_empty ( )
}
Changing the email can then specify what else must change:
fn change_email ( user : & mut User , new_email : String )
requires new_email . is_valid_email ( )
ensures user . email == new_email
ensures user . email_verified == false
Without the postcondition ensures , a generated implementation might update the address while leaving email_verified set to true .
Data structures whose properties can be encoded as invariant conditions are particularly well suited to this sort of design:
struct MinHeap < T : Ordered > {
items : Array < T >
invariant forall i in 1 .. < items . length :
items [ parent ( i ) ] <= items [ i ]
}
struct BTreeNode < K , V , const ORDER : usize > {
keys : Array < K >
values : Array < V >
children : Array < NodeRef >
is_leaf : bool
invariant keys . length == values . length
invariant keys . length <= ORDER - 1
invariant strictly_increasing ( keys )
invariant is_leaf
implies children . length == 0
invariant ! is_leaf
implies children . length == keys . length + 1
}
Effects
Effects are a way to denote explicit capabilities in methods.
Typically these are behaviors such as "file/network IO, memory allocation, state changes", etc. By requiring that method effects be explicit, you can use a method's signature as a verified contract of its permitted behaviors.
As one example, suppose we have a method:
fn load_settings ( path : Path ) -> Settings
The return type says nothing about where the settings come from or what the operation may do.
An effect-aware signature does:
fn load_settings ( path : Path ) -> Settings
requires path . exists ( )
requires path . is_file ( )
ensures result . is_valid ( )
effects {
fs . read ( path ) ,
alloc
}
Now compare that with a function that merely parses settings already held in memory:
fn parse_settings ( contents : String ) -> Settings
requires ! contents . is_empty ( )
ensures result . is_valid ( )
effects {
alloc
}
These functions may return the same TYPE, but they have meaningfully different OPERATIONAL behavior. That difference should be visible at the call site.
Cache get that updates recency internal metadata
At first glance, you might think that the below method is read-only:
fn get ( cache : & mut Cache < K , V > , key : K ) -> Option < V >
But many caches maintain usage statistics, and it may be the case that our get method is mutating. The below signature is explicit about behavior like this:
fn get ( cache : & mut Cache < K , V > , key : K ) -> Option < V >
ensures result . is_some ( ) == old ( cache . contains ( key ) )
ensures cache . entries == old ( cache . entries )
effects {
state . write ( cache . recency )
}
The cache contents do not change, but its eviction metadata does.
A function that must not allocate
Suppose that we have a method which writes/encodes some value into a caller-supplied buffer:
fn encode (
value : Message ,
destination : & mut ByteBuffer
) -> usize
requires destination . remaining >= value . encoded_size ( )
ensures result == value . encoded_size ( )
ensures destination . position == old ( destination . position ) + result
effects {
memory . write ( destination )
}
The implementation can write into the supplied buffer but cannot create a temporary array or string (because it does not possess the alloc effect)
