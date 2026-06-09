---
source: "https://entropicthoughts.com/llms-and-almost-good-code"
hn_url: "https://news.ycombinator.com/item?id=48459991"
title: "LLMs and Almost Good Code"
article_title: "LLMs and almost good code"
author: "ibobev"
captured_at: "2026-06-09T13:07:37Z"
capture_tool: "hn-digest"
hn_id: 48459991
score: 1
comments: 0
posted_at: "2026-06-09T12:04:39Z"
tags:
  - hacker-news
  - translated
---

# LLMs and Almost Good Code

- HN: [48459991](https://news.ycombinator.com/item?id=48459991)
- Source: [entropicthoughts.com](https://entropicthoughts.com/llms-and-almost-good-code)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T12:04:39Z

## Translation

タイトル: LLM とほぼ良好なコード
記事のタイトル: LLM とほぼ良好なコード

記事本文:
tl;dr : 私の新しい優先事項は、最上位の llm が簡単なタスクに取り組んでいることです。
おそらく必要より 10% 複雑なコードを生成します。私も思います
この複雑さは正しいコードから来ているので、私たちはあまりにも簡単にこの複雑さを受け入れてしまいます。
今ここで、差し迫った問題を解決しています。これは影響を与える可能性があります
長期的なメンテナンスに。
この発見の背景には、汚い配管工事をする必要があったことが挙げられます。
仕事のプロジェクト。これは既存のものをほとんど反映した単純な変更でした
機能性。私の経験では、これは llm に完全に適合するため、
フロンティア モデルを使用してコードを生成します。最終的に変更されたのは合計です
200 行強、ほとんどが追加です。
生成されたコードのうちこれから説明する部分は 24 行の関数です。
任意の（ユーザー指定の）文字列を安全な http ヘッダーに変換します
値。 1 1 混乱を避けるために、安全な値にエンコードする必要があります
間違いを防ぐだけでなく、http ヘッダー インジェクション攻撃を防ぐこともできます。
toHeaderValue :: テキスト -> テキスト
toHeaderValue raw =
させてください
attrChars = "!#$&+-.^_`|~"
PadHex t = if Text.length t < 2 then "0" <> t else t
パーセントエンコード c =
if (isAscii c && isAlphaNum c) || elem c attrChars then
Text.singleton c
それ以外の場合
Text.concat
[ "%" <> PadHex (Text.toUpper (Text.pack (showHex b "" )))
| b <- ByteString.unpack (encodeUtf8 (Text.singleton c))
】
rfc5987Encode = Text.concatMap パーセントエンコード
isPrintable c = c >= ' ' && c /= '\DEL'
replacePathSeparator c =
c == '/' || の場合c == '\\' その後
'_'
それ以外の場合
c
きれいになった =
Text.map replacePathSeparator (Text.filter isPrintable raw)
で
rfc5987エンコードがクリーン化されました
この関数をこのように見ると、明らかに少しやりすぎのように思えます
複雑ですが、これは 200 行の変更のうちのわずか 24 行であったことを思い出してください。私は
根底にある考え方が正しいことと、生成された

d テスト
私がカバーしたいエッジケースをすべてカバーしました。きれいなコードではありませんが、
しかし、それはテストによって正しいことが証明されています。
さらに重要なのは、それが非常にローカルであるということです。このコードに関して必要なことがあれば
交換する場合は、何も触れずに交換できます。見習いレベル
プログラマーはどこでも同じようにコードの品質を心配します。ずっとやりたかった
「心配しないでください、地元のことです」という記事を書いて、これらのことを伝えます
プログラマーは、コードが自己完結型である限り、コードの品質が悪くても問題ないと考えています。
小さな場所。 2 2 私がその記事を書かなかった理由は、
記事に必要なサンプルがたくさんあるのを待っています。しかし、私はそれを忘れ続けます
それらを収集するため、これまでに収集した例の総数はゼロです。
このコードを受け入れました。実装が機能する必要があり、このコードが必要でした
明らかに効いた。それはまさに今そこにありました。それは愚かなことだったでしょう
受け入れないでください！それを受け入れるのは簡単な選択だったし、確かに悪いことではなかった
決断。
しかし、嬉しい運命のいたずらで、このプロジェクトの CI パイプラインには
必須のステートメントテストカバレッジチェック 3 3 私はステートメントの大ファンではありません
しかし、それは私が何年も書くのを先延ばしにしてきた別の記事です。 、
そして、このコードのチェックは失敗しました。場所と理由を見つけられるかどうかを確認してください。
ヒントを与えます。これは、padHex 関数に関係しています。
0x0 ～ 0xff の範囲の 16 進値で、それより小さい場合はゼロ埋めされます。
0x10 。
PadHex に渡されたデータはすでに isPrintable フィルターを通過しています。
これにより、 0x20 より下のすべてのバイトが削除されます。したがって、padHex に渡される値はありません。
0x10 を下回っても、何も埋め込まれることはありません。それは常にノーオペです。
ステートメント カバレッジ チェックは、 PadHex のパディング ブランチに対して警告を出します。
テストなしで行使されます。それ

実際、テストでそれを実行することは不可能です。
一方で、percentEncode が常に次のように呼び出されると想定すべきではありません。
0x1f より大きい文字 (現時点でそれが当てはまっている場合でも)。
そのような仮定は、遠くでの不気味な行動に依存しています。
はこの関数に対してローカルであるため、これは避けたいと考えています。
一方で、取材レポートも正しい：何かがある
この構造全体については厄介です。
そこで私は介入して独自の実装を作成しました。終了した実装
アップ送料はこれに近かったです:
toHeaderValue :: テキスト -> テキスト
toHeaderValue =
させてください
tainPrintable = Text.filter ( \ c -> c >= ' ' && c /= '\DEL' )
replacePathSeparators = Text.replace "/" "_" 。 Text.replace "\\" "_"
-- URL エンコードは、正当な RFC5987 エンコードでもあります。
rfc5987エンコード = decodeUtf8 。 urlEncode True 。エンコードUtf8
で
rfc5987エンコード 。 replacePathSeparators 。印刷可能を保持
これは複雑さが 15 行短くなります。これは変化の約 8 % に相当します。
llm は不正なコードを生成しませんでした。 4 4 ある意味、コードの方が優れています。の
rfc 5987 エンコードは URL エンコードよりも緩いので、私の実装は
技術的にオーバーエンコードされます。 8% 複雑なコードが生成されました。
そうである必要がありました。今日ではそれは大惨事ではありませんが、出荷のプレッシャーがあるときは、
それは今すぐそこにあるので、それを受け入れるのは簡単です、そしてそれは解決します
問題。私は受け入れて、8% 複雑すぎるコードを出荷しようとしていました。
もっと深く調べて問題に気づいたのは偶然でした
それと一緒に。
この経験により、私には答えのない多くの疑問が残りました。
不必要に複雑ではあるが、他のすべての変更についてはどうでしょうか。
とにかく受け入れますか？
これが簡単なケースで、より複雑なケースで LLM を作成した場合はどうなるでしょうか。
タスクでは、c が生成されます。

8 % を超えるオードは複雑すぎます (20 % や 40 % など)。
それとも必要以上に 3 倍も複雑ですか?
不必要に複雑なコードを受け取ったときに、私たちは立ち止まるでしょうか?
それとも、今日は災害ではなく、すぐそこにあるので、受け入れるでしょうか。
今？
コードを出荷し続けると 1 ～ 2 年後に何が起こるか
常に必要以上に複雑になっていませんか?
一方で、これは私を心配させます。一方で、明らかなことは、
反論は、コード作成ロボットは十分な速さで上達するため、
数年後にこれが問題になったとき、彼らはそれに対処する方法を知るでしょう。
混乱を避けるために、安全な値にエンコードする必要があります。
間違いを防ぐだけでなく、http ヘッダー インジェクション攻撃を防ぐこともできます。
その記事を書かなかった理由は、私が
記事に必要なサンプルがたくさんあるのを待っています。しかし、私はそれを忘れ続けます
それらを収集するため、これまでに収集した例の総数はゼロです。
私はステートメントの大ファンではありません
しかし、それは私が何年も書くのを先延ばしにしてきた別の記事です。
ある意味、そのコードの方が優れています。の
rfc 5987 エンコードは URL エンコードよりも緩いので、私の実装は
技術的にオーバーエンコードされます。
これが気に入ってもっと欲しい場合は、そうすべきです
コーヒーを買ってきてください。
これは、170 を超える未処理のアイデアを記事にまとめるのに役立ちます。
私の素晴らしい妻に叫びます
誰のサポートがなかったら、私は決してしなかったでしょう
最初の文を過ぎてください。 ♥
S = k log W。
コメント?送信
私にメールをください。
新しい記事を購読することもできます。

## Original Extract

tl;dr : My new prior is that top-of-the-line llm s working on easy tasks
generate code that is maybe 10 % more complicated than necessary. I also think
we accept this complexity too easily, because it comes from code that is right
here , right now , solving an immediate problem. This may have consequences
for maintenance in the long term.
The background to this discovery was that I needed to do some crud plumbing in
a work project. It was a simple change that mostly mirrored existing
functionality. This is a perfect fit for llm s, in my experience, so I used a
frontier model to generate the code for it. The change ended up being a total of
just over 200 lines, mostly additions.
The part of the generated code we’ll talk about is a 24-line function that
converts an arbitrary (user-supplied) string to a safe http header
value. 1 1 Encoding it into a safe value is necessary to avoid confusing
mistakes, but also to prevent http header injection attacks.
toHeaderValue :: Text -> Text
toHeaderValue raw =
let
attrChars = "!#$&+-.^_`|~"
padHex t = if Text.length t < 2 then "0" <> t else t
percentEncode c =
if (isAscii c && isAlphaNum c) || elem c attrChars then
Text.singleton c
else
Text.concat
[ "%" <> padHex (Text.toUpper (Text.pack (showHex b "" )))
| b <- ByteString.unpack (encodeUtf8 (Text.singleton c))
]
rfc5987Encode = Text.concatMap percentEncode
isPrintable c = c >= ' ' && c /= '\DEL'
replacePathSeparator c =
if c == '/' || c == '\\' then
'_'
else
c
cleaned =
Text.map replacePathSeparator (Text.filter isPrintable raw)
in
rfc5987Encode cleaned
When looking at this function like this, it obviously seems a bit too
complicated, but recall that this was just 24 lines in a 200-line change. I
confirmed that the underlying idea was correct, and that the generated tests
covered all the edge cases I would want to see covered. It’s not pretty code,
but it is proven correct by tests.
More importantly, it is highly local. If anything about this code needs
replacing, it can be replaced without touching anything else. Apprentice-level
programmers worry equally about code quality everywhere; I’ve long wanted to
write an article called “ Don’t worry, it’s local ” where I tell these
programmers that bad code quality is fine, as long as it’s self-contained in a
small location. 2 2 The reason I haven’t written that article is that I’m
waiting on a bunch of examples the article would need. But I keep forgetting to
collect them, so the total number of examples I have collected so far are zero.
I accepted this code. I needed the implementation to work, and this code
obviously worked. It was right there , right now . It would have been silly to
not accept it! Accepting it was the easy choice, and certainly not a bad
decision.
However, in a pleasant twist of fate, the ci pipeline for this project has a
mandatory statement test coverage check 3 3 I’m not a huge fan of statement
coverage, but that’s another separate article I’ve put off writing for years. ,
and that check failed for this code. See if you can spot where and why.
I’ll give you a hint: it has to do with the padHex function, which takes a
hexadecimal value in the range 0x0 – 0xff and zero-pads it if it is less than
0x10 .
The data passed into padHex has already gone through the isPrintable filter,
which removes all bytes lower than 0x20 . Thus no value passed to padHex is
ever below 0x10 , and it never ends up padding anything! It is always a no-op.
The statement coverage check warns on the padding branch of padHex , because it
is exercised by no test. It is in fact impossible to exercise it in a test.
On the one hand, we shouldn’t assume percentEncode is always called with
characters greater than 0x1f , even if that happens to be true at the moment.
Such an assumption relies on spooky action at a distance, which – even if it
is local to this function – we want to avoid.
On the other hand, the coverage report is right too: there is something
awkward about this whole construction.
So I stepped in and wrote my own implementation. The implementation that ended
up shipping was closer to this:
toHeaderValue :: Text -> Text
toHeaderValue =
let
retainPrintable = Text.filter ( \ c -> c >= ' ' && c /= '\DEL' )
replacePathSeparators = Text.replace "/" "_" . Text.replace "\\" "_"
-- URL encoding is also legal RFC5987 encoding.
rfc5987Encode = decodeUtf8 . urlEncode True . encodeUtf8
in
rfc5987Encode . replacePathSeparators . retainPrintable
This is 15 lines of complexity shorter. That’s around 8 % of the change.
The llm did not generate bad code. 4 4 In some sense, its code is better. The
rfc 5987 encoding is more lax than url encoding, so my implementation
technically over-encodes. It just generated code that was 8 % more complex than
it needed to be. That’s not a disaster today, and when there’s pressure to ship,
it is easy to accept it because it is right there , right now , and it solves
the problem. I accepted and was about to ship code that was 8 % too complex.
It was only by chance I looked into it more deeply and realised the problems
with it.
This experience leaves me with a bunch of questions I don’t have answers to.
What about all the other changes that are also unnecessarily complex, but which
I accept anyway?
What if this was an easy case, and when we sic an llm on a more complicated
task, it generates code that is more than 8 % too complex, like 20 %, or 40 %,
or even 3× more complex than it needs to be?
Will we put our foots down when we get code that is so unnecessarily complex?
Or will we accept, because it’s not a disaster today, and it is right there ,
right now ?
What happens in a year or two, when we continue shipping code that’s
consistently more complex than it needs to be?
On the one hand, this worries me. On the other hand, the obvious
counter-argument is that code-writing robots improve fast enough that in two
years’ time when this becomes a problem, they will know how to deal with it.
Encoding it into a safe value is necessary to avoid confusing
mistakes, but also to prevent http header injection attacks.
The reason I haven’t written that article is that I’m
waiting on a bunch of examples the article would need. But I keep forgetting to
collect them, so the total number of examples I have collected so far are zero.
I’m not a huge fan of statement
coverage, but that’s another separate article I’ve put off writing for years.
In some sense, its code is better. The
rfc 5987 encoding is more lax than url encoding, so my implementation
technically over-encodes.
If you liked this and want more you should
buy me a coffee .
That helps me turn my 170+ ideas backlog into articles.
Shoutout to my amazing wife
without whose support I would never
make it past the first sentence. ♥
S = k log W.
Comments? Send
me an email .
You can also subscribe for new articles .
