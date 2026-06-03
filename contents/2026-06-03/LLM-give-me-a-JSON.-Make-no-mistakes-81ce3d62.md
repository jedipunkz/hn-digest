---
source: "https://nobodywho.ooo/posts/llm-give-me-a-json/"
hn_url: "https://news.ycombinator.com/item?id=48369379"
title: "LLM, give me a JSON. Make no mistakes"
article_title: "LLM, give me a JSON. Make no mistakes. - NobodyWho"
author: "marek-hradil"
captured_at: "2026-06-03T00:47:55Z"
capture_tool: "hn-digest"
hn_id: 48369379
score: 8
comments: 0
posted_at: "2026-06-02T12:31:28Z"
tags:
  - hacker-news
  - translated
---

# LLM, give me a JSON. Make no mistakes

- HN: [48369379](https://news.ycombinator.com/item?id=48369379)
- Source: [nobodywho.ooo](https://nobodywho.ooo/posts/llm-give-me-a-json/)
- Score: 8
- Comments: 0
- Posted: 2026-06-02T12:31:28Z

## Translation

タイトル: LLM、JSON をください。間違いを犯さないでください
記事のタイトル: LLM、JSON をください。間違いを犯さないでください。 - 誰でもない
説明: では、LLM 出力を JSON にするにはどうすればよいでしょうか?ボンネットの下では何が起こっているのでしょうか?そして、それを信頼性が高く、高速にするにはどうすればよいでしょうか?制約付きサンプリングについて詳しく説明します。

記事本文:
LLM さん、JSON をください。間違いを犯さないでください。 - 誰でもない
誰でもない
について
アプリ
ブログ
ドキュメント
GitHub
ギットハブ
誰でもない
について
アプリ
ブログ
ドキュメント
GitHub
ギットハブ
LLM さん、JSON をください。間違いを犯さないでください。
では、LLM 出力を JSON にするにはどうすればよいでしょうか?ボンネットの下では何が起こっているのでしょうか?そして、それを信頼性が高く、高速にするにはどうすればよいでしょうか?
想像してみてください。ついにアプリケーション用に LLM 推論をセットアップすることができ、アプリケーションが応答できるようになりました。
そして、それは非常に多くのことを行うことができます！
しかし、これらのユースケースのほとんどでは、テキスト「だけ」を取得することは非常に制限されます。実際、チャットボット以外のほとんどのユースケースを機能させるには、
JSON などのより構造化された情報が必要になります。したがって、プロンプトに次のように追加するだけです。
出力は必ず JSON 形式で提供してください。間違いを犯さないでください。
JSON 出力がどんどん長くなるにつれて、どういうわけかスーパースマート モデルが時々失敗するようになります。オブジェクトキーを正しく取得できないことは別として、
最後のキーと値の末尾に追加の が追加され、パーサーにエラーが発生します。もっと良い方法はないのかと疑問に思うかもしれません。
LLM がどの形式を生成するかを正確に制御できることは、非常に価値があり、技術的に非常に興味深いものです。
そこで、「間違いを犯さない」をどのようにして乗り越えるか、そして推論エンジンがどのようにしてそれを確実かつ高速に実行するのかを深く掘り下げてみましょう。
注 : JSON スキーマと GBNF に慣れている場合は、「文法の処理」セクションに進んでください。
最初に思い浮かぶ解決策は、メッセージ レベルで何らかの再試行戦略を採用することです。基本的に:
True の場合:
答え = llm (プロンプト)
if is_json (答え):
休憩
これは機能します。これについて私が言える唯一の良い点は、LLM を完全なブラックボックスとして扱うことができるということです。
一部のライブラリでは実行可能です (実際、これが LangChain の機能だと私は信じています)。ネガティブな点については、

たくさんあります:
「運が悪い」か、より小さなモデルを使用すると、目的の出力に到達する前に、非常に長い時間または永久にループする可能性があります。
たとえすべてが間違っているわけではないとしても、メッセージ全体を破棄することで、膨大な数のトークンを無駄にしていることになります。
JSON を取得できるようにするには、特定のパーサーを構築またはダウンロードする必要がありますが、これはあまり拡張性がありません。
したがって、その必要がない場合は、これを行わないでください。ただし、LLM を詳しく調べると、次のことがわかります。
もう少し原則的なアプローチ。
観察できることは 2 つあります。
まず、LLM はトークンごとに出力トークンを生成します。通常、何かが間違っていることを確認するために、答え全体を生成する必要はありません。
モデルが正しい形式ではない最初のエラーを起こした場合、すぐに再試行できます。この方法では、メッセージ全体を削除するわけではありません。
ただし最後のトークンだけ:
答え = ""
True の場合:
トークン = llm 。 next_token (プロンプト + 回答)
トークン == "<eos>" の場合:
休憩
if is_partial_json (答え):
答え += トークン
第二に、回答トークンは突然現れるわけではありません。テキストが与えられると、LLM は次のトークンの確率分布を生成します。
分布からただちにサンプリングするのではなく、すべてのトークンの確率を設定することから始めることができます。
このようにして、正しいトークンのみをサンプリングする (したがって出力する) ことが保証されます。
数値だけが必要な場合は、次のようなことができます。
専門的には、このプロセスは「マスキング」と呼ばれます。対処すべきもう 1 つの詳細は、0 にしたくないトークンの確率を縮小することです。
分布プロパティが壊れてしまいます (確率の合計が 1 になることが必要です)。したがって、実際には、解決策は基礎となるものを設定することです。
-inf にログインすると、不要なトークンの確率が 0 になります。

他のトークンをわずかに押し上げます。
擬似コードは次のようになります。
答え = ""
True の場合:
トークン = llm 。 next_token (
プロンプト + 回答、
マスク = possible_next_json_tokens (プロンプト + 回答)
)
トークン == "<eos>" の場合:
休憩
したがって、まだループしているにもかかわらず、破棄は行われていません。運が悪いはずはなく、コンピューティングを無駄にすることはありません。
難しい部分は、一般に「マスク」と呼ばれる、次に考えられるトークンを指定する方法と、それを迅速に行う方法です。
したがって、LLM は私たちを待っていません。
JSONとは一体何なのでしょうか？マスクを構築するには、このトークンにトークンごとに応答できる必要があります。
テキスト形式を正確に指定できる優れた方法の 1 つは正規表現です。
残念ながら、名前が示すように、正規表現は通常の言語を指定するために作成されていますが、JSON はそうではありません。
正規表現機能 (の基本セット) では、たとえば、開かれた { が対応する } によって閉じられることを保証することはできません。
このユースケースにより適した方法は、JSON スキーマを使用することです。
JSON スキーマを知らない場合は、JSON スキーマは本質的にメタ言語です。
JSON の上に、期待する JSON 形式を指定します。このようにして、たとえば次のように言えます。
{ "タイプ" : "オブジェクト" }
任意の JSON オブジェクトを取得します。または、より具体的なものが必要な場合は、次のようにします。
{
"タイプ" : "オブジェクト" ,
"プロパティ" : {
"名前" : "文字列" ,
"年齢" : "整数"
}
}
これはより具体的な仕様であり、一部の推論エンジン/API は JSON スキーマを直接受け入れます。 (例: Claude API 、OpenAI API 、vLLM など)
一方で、私たちは実際に、私たちが望むものの「下位レベル」の仕様に向かって進んだわけではありません。 JSON スキーマは依然として非常に抽象的です。
すべての形式 (正規表現、JSON スキーマ、文法) において、LLM に何を伝えるかが依然として重要であることに注意してください。
それはあなたが答えに期待していることです。 LLM は、そうなることを知りません。

制約があるので、期待していることを伝えないと、
トークンの配布は出力が制約されていないかのように残り、パフォーマンスに悪影響を与える可能性があります。
より低レベルで一般的な仕様は、文法、特にコンテキストフリー文法 (CFG) を渡すことによって実現できます。
文法は、どのような文字列を生成できるかを正確に記述したものです。文法は通常、プログラミング言語理論 (パーサー)、言語学、理論的なコンピューター サイエンスの一部などの分野で利用されます。
推論エンジンでは通常、正規表現で指定されているかどうかに関係なく、すべての制約された世代の処理の基礎となります。
JSON スキーマまたは「手動」。また、Python 構文など、広く使用されている構造化形式のほとんどもカバーしています。
このような文法の例は次のとおりです。
INT ::= "-"? [0-9]+
これは、符号付き整数を生成するための INT ルールを提供します。
これらのルールは、互いに「呼び出し」て、より大きな構造を形成する方法で構成できます。
ROOT ::= "{ \"名前\":" STRING ",\"年齢\":" INT "}"
文字列 ::= "\"" [^"]* "\""
INT ::= "-"? [0-9]+
これは、前に説明した JSON スキーマの非常に単純化された文法です。 STRING ルールは、「,」のないシーケンスを消費します。
ROOT は単にエントリポイントとして機能し、すべてをまとめます。文法を書き留めるこの正確な形式は GBNF として知られています
llama.cpp だけでなく他のエンジンの制約付き生成の基礎にもなります。
さらに詳しく読みたい場合は、ここから始めることをお勧めします。
文法を整備したので、「ユーザー」側での作業が完了しました。ユーザーからの一般的で正確な仕様が得られました。
彼らが実際に望んでいることについて。ただし、これらはトークンのマスクではありません。ここで、推論エンジンがどのように機能するのかをさらに詳しく見ていきます。
文法をマスクに変換します (高速)。
文法が int にどのように変換されるかを理解するには

マスクについては、少し理論を理解する必要があります。
まず簡単に始めましょう。単純な整数ルールと文字列全体を前もって使用し、それが文法に一致するかどうかを判断するだけです。
INT ルールから始めます。
INT ::= "-"? [0-9]+
糖衣構文を削除します。
INT ::= "-" [0-9]+ | [0-9]+
+ と * を再帰に置き換えます。
INT ::= "-" 桁 |数字
数字 ::= [0-9] | [0-9] 桁
これにより、式が単純化され、作業が容易になります。実際に何が起こったのか少し迷っている場合は、
正規表現変換をいくつか行っただけで、これは以前と同じです。
ここで、マスクを作成するには、スタックを備えたマシン (適切な用語では「プッシュダウン オートマトン」、PDA) を使用します。
スタックが何なのかわからない場合は、一方の端から要素を追加したりポップしたりできる配列のように考えることができます。
私たちの場合、これらは文字 (実際にはバイト) とルールの名前 (INT、ROOT、STRING など) になります。
作業を少し簡単にするために、もう一度、より単純なケースに焦点を当てます。まだトークンから文字列を生成しません。
代わりに、最初に文字列全体が与えられ、そのような文字列が存在するかどうかをチェックするだけのタスクが課されます。
現在の文法を使用して生成しても、そうでなくても生成できます。次のアプローチは通常、トップダウン解析と呼ばれます。
繰り返しますが、70 年代からここにいます。
次の 4 つの簡単なルールを続けていきます。
まず、ルート ルールをスタックに追加します。
先頭にルールがある場合は、それをポップし、その定義の 1 つに置き換えます。
入力と一致する文字が先頭にある場合は、両方を消費します。そうしないとスタックして拒否されます。
スタックと入力の両方が空の場合は、 を受け入れます。
これがどのように行われるかを示すために、INT ルールを見て -1234 と入力します。まず、スタックは空なので、(1) に進みます。
として

p は規則なので、(2) に従います。どの定義を選択すればよいかをどのようにして知ることができるのでしょうか?
ここでは、マシンが単に知っていると仮定しましょう。
先頭が入力と一致する文字なのでルール(3)を適用します。
それでは、何も新しいことは起こりません。先ほどのルール (2、3、2、3) に従うだけです。
空のスタックと空の入力に到達するまで。よくやった！ルール (4) を適用して受け入れることができます。
このようにして、文字列が指定された文法に準拠していることがわかります。 .677 などの非整数入力を解析している場合、
この状態に到達します。
どちらも文字 (または文字範囲) であり、一致しないため、入力は強制的に拒否されます。
このようなアプローチをどのように実装するかについては理解できたと思います。
最後の 2 つの問題は次のとおりです。
文字列全体を消費できますが、LLM はトークンを生成します。
マシンはどの定義を選択すべきかをどのようにして判断するのでしょうか?
これらの問題をどのように正確に解決できるかは、使用する特定の文法ライブラリによって異なります。
以下では、llama.cpp がこれをどのように行うかについて説明します。
文字列への適応は、見た目よりも簡単になりました。ここで行うことは、各トークンに対してトップダウン解析を実行するだけです。
前のアルゴリズムの唯一の例外は、入力全体とスタック全体を消費する必要がないことです。
途中で行き詰まらないようにするだけで十分です。結果はトークンのみを含むディストリビューションになります。
文法に準拠していること。
ただし、小さな注意点が 1 つあります。次の分布からのサンプリングを想像してください。
INT ルールを使用してスタックを再度初期化し、解析を実行すると、最終的には -42 と - だけになります。
次に、トークンに関連付けられたそれぞれのスタックを選択してサンプリングします。 -42 をサンプリングしたとしましょう。
続行して次のトークンを取得するには、INT ルールでスタックをリセットするだけではなく、次のことを覚えておく必要があります。
文法上の位置とコル

応答スタック。
あなたはまだ、私たちが扱っているのはこの神託機械であり、何らかの方法で情報を知っていると主張するかもしれません。
ルール [0-9] DIGITS または DIGITS が正しいパスである場合。あなたは正しいでしょう。繰り返しますが、簡単な解決策を提案できます。
決断をしなければならないときは、両方の方向に進んでください。
具体的には、DIGITS | のようなルールを見つけたときを意味します。 [0-9] DIGITS あなたが持っている現在のスタックをフォークするだけです
1 つは DIGITS を使用し、もう 1 つは [0-9] DIGITS を使用します。トークンが通過するかどうかを確認する
次に、「受け入れるスタックはありますか?」という質問に答えます。すべてのトークンの長さは有限であるため、気にする必要はありません
無限再帰について (常に DIGITS ブランチを展開するスタックがある)。
それで終わりです！概念的には、これが llama.cpp の仕組みです。
llama.cpp ソリューションはすべてを実行するため、遅いソリューションの側にあることに留意してください。
推論時。最終的には O(vocab_size * active_stacks) のようなコストが発生します。
トークンごと、および文法を通るさまざまなルートごとにスタックを保持しているためです。
現在の語彙サイズ (100k+) と長い文法では、信じられないほど時間がかかる可能性があります。
確かに、XGrammar や

[切り捨てられた]

## Original Extract

So how exactly do you make your LLM output a JSON? What happens under the hood? And how do you make it reliable and fast? Diving into constrained sampling.

LLM, give me a JSON. Make no mistakes. - NobodyWho
NobodyWho
About
Apps
Blog
Docs
GitHub
Github
NobodyWho
About
Apps
Blog
Documentation
GitHub
Github
LLM, give me a JSON. Make no mistakes.
So how exactly do you make your LLM output a JSON? What happens under the hood? And how do you make it reliable and fast?
Imagine, you have finally managed to set up the LLM inference for your application, and now it is even able to respond to you.
And it can do so much stuff!
But for most of these use cases, getting "just" text back is very limiting. In fact, in order to make most of the non-chatbot use cases work,
you would need more structured info like JSON. So you just append to the prompt:
Remember to give me the output in JSON format. Make no mistakes.
As the JSON output gets longer and longer, somehow your super smart model fails from time to time. Apart from not getting the object keys right,
it appends the additional , at the end of the last key-value, which makes the parser complain. You might ask, is there a better way?
Being able to control what format exactly does your LLM produce is super valuable and technically super interesting.
Let us thus take a deep dive into how you go past "make no mistakes" and how the inference engines do it reliably and fast.
Note : If you feel familiar with JSON schemas and GBNF, just skip into the section "Processing Grammars".
The first solution that comes to mind is just to employ some retry strategy at the message level. Essentially:
while True :
answer = llm ( prompt )
if is_json ( answer ) :
break
This works. The only positive thing I have to say about it is that you can treat the LLM as a complete blackbox, which might
be viable for some libraries (actually I believe this is what LangChain does). For the negatives, there are plenty:
by being "unlucky" or employing smaller models, you can be looping for a very long time or forever, before reaching the desired output
you're wasting an enormous number of tokens, by discarding whole messages, even though they might not be all wrong
to be able to get a JSON, you need to construct or download a specific parser, which is not very extensible
So if you don't have to, just don't do this please. However, by looking more closely into the LLM, you can have a little
bit more principled approach.
There are two observations, which we can make.
First, LLMs generate outputs token by token. Usually you don't have to generate the whole answer to see that something is wrong.
We can retry right away when the model makes the first error which is not in the right format. This way, we are not deleting the whole message,
but just the last token:
answer = ""
while True :
token = llm . next_token ( prompt + answer )
if token == "<eos>" :
break
if is_partial_json ( answer ) :
answer += token
Secondly, the answer tokens don't appear out of the blue. Given a text, LLMs produce a probability distribution on the next token.
Instead of simply sampling from the distribution immediately, we can start by setting the probability of all tokens
leading to incorrect output to 0. This way, we are guaranteed to only sample (and thus output) a correct token.
If we wanted just a number, we could do something like:
Technically, this process is called "masking". One more detail to address is that just shrinking the probability of the tokens we don't want to 0
would break the distribution property (we want the probabilities to sum to 1). In reality the solution is therefore to set the underlying
logits to -inf , which will result in turning the unwanted tokens' probabilities to 0, but slightly bumping the other tokens up.
The pseudocode then could look like this:
answer = ""
while True :
token = llm . next_token (
prompt + answer ,
mask = possible_next_json_tokens ( prompt + answer )
)
if token == "<eos>" :
break
So even though we are still looping, there is no discarding going on - we can't be unlucky and we are not wasting compute.
The hard part is now how to specify the possible next tokens, generally called "mask", and how to do that quickly,
so the LLM is not waiting for us.
What exactly is JSON? To construct the mask, we have to be able to answer this token by token.
One of the great ways that we are able to precisely specify some text format is regexes.
Unfortunately, as the name suggests, regexes are made for specifying regular languages, which JSON is not.
With the (basic set of) regex features, you won't be able to for example guarantee that any opened { will also be closed by a corresponding } .
A more fitting way for this use case is employing JSON schemas .
If you don't know JSON schemas, they are essentially a metalanguage
on top of JSON, to specify what JSON format you expect. This way, we can say for example:
{ "type" : "object" }
to get any JSON object. Or, if you want something more specific:
{
"type" : "object" ,
"properties" : {
"name" : "string" ,
"age" : "integer"
}
}
This is a more concrete specification and some of the inference engines/APIs accept JSON schemas directly! (e.g. Claude API , OpenAI API , vLLM , etc.)
On the other hand, we did not really move towards a "lower level" specification of what we want; JSON schemas are still quite abstract.
Note that with all of the formats (regexes, JSON schemas, grammars) it is still important to tell the LLM what
it is that you expect in the answer. The LLM does not know it is being constrained, so if you do not tell it what you expect,
the token distributions will stay as if the output was not constrained, possibly hurting performance.
More low-level and general specification can be achieved by passing in grammars, specifically Context-Free Grammars (CFGs).
Grammars are precise descriptions on what strings can be generated - they are typically utilized in fields like programming language theory (parsers), linguistics, and some parts of theoretical computer science.
In the inference engines they typically underlie processing all of the constrained generation, whether it is specified by regexes,
JSON schemas or "manually". They also cover most of the wide-spread structured formats used, such as Python syntax .
Example of such a grammar could be:
INT ::= "-"? [0-9]+
which provides an INT rule for generating any signed integer.
These rules can then be composed in a way where they "call" each other, to form larger structures:
ROOT ::= "{ \"name\":" STRING ",\"age\":" INT "}"
STRING ::= "\"" [^"]* "\""
INT ::= "-"? [0-9]+
which is a very simplified grammar for the JSON schema we discussed earlier. The STRING rule consumes any sequences without " ,
and ROOT just acts as an entrypoint and assembles everything together. This exact format of writing down the grammars is known as GBNF
and underlies the constrained generation in llama.cpp, but also other engines.
If you want to read more about it, I recommend starting here .
With grammars in place, we have done the work on the "user's" side - we have a general and exact specification from the user
of what they actually want. Still, these are not the token masks. Now, we will dive deeper into how the inference engines
convert the grammars into masks (fast).
To understand how the grammars then translate into masks, we will need a bit of theory under our belts.
Let us start simple: with a simple integer rule and a whole string upfront, for which we just decide if it corresponds to the grammar.
Starting with the INT rule:
INT ::= "-"? [0-9]+
We remove the syntactic sugar:
INT ::= "-" [0-9]+ | [0-9]+
and replace the + and * with recursion:
INT ::= "-" DIGITS | DIGITS
DIGITS ::= [0-9] | [0-9] DIGITS
This gets us simplified expressions, which are easier to work with. If you feel a bit lost on what actually happened,
we just made a few regex transformations, which are equivalent to what we had before.
Now, to produce the masks, we will use a machine with a stack (or in the right terms "pushdown automaton", PDA).
If you don't know what a stack is, you can think about it like an array, where from one end, we are allowed to append and pop elements.
In our case, these will be characters (in practice bytes) and names of the rules (INT, ROOT, STRING, ...).
To make things a little bit easier, we will again focus on a simpler case - we won't generate strings from tokens yet.
Instead, we will be given a whole string at the start, and only be tasked with checking if such a string
can be generated with the current grammar or not. The following approach is usually called top-down parsing
and again, has been here since the 70s .
We will continue with the following four simple rules:
Start with appending the root rule to the stack.
If there is a rule at the top, pop it and replace it with one of its definitions.
If there is a character at the top, which matches the input, consume both. Otherwise get stuck and reject .
If both stack and input are empty, accept .
To illustrate how this would go, look at the INT rule and input -1234 . First, the stack is empty, so we proceed with (1).
As the top is a rule, we follow with (2). How do we know which definition to choose from?
For now, let's pretend the machine just knows :
Applying rule (3), as the top is a character which matches the input.
Then, nothing new happens. We just go by the previous rules (2, 3, 2, 3).
Until arriving at the empty stack and empty input. Well done! We can apply rule (4) and accept.
This way we know the string conforms to the grammar specified. If we were parsing non-integer input, like .677 ,
we would arrive at this state:
As both are characters (or character ranges) and don't match, we are forced to reject the input.
Hopefully, now you have an idea of how you would implement such an approach.
The last two problems we have are:
We can consume whole strings, but LLMs produce tokens.
How does the machine just know which definition to choose?
How exactly you solve these problems then depends on the particular grammar library you use.
In the following, I will describe how llama.cpp does this.
Adapting to strings is now easier than it looks. What we will do is simply run the top-down parsing for each token.
The only exception to the previous algorithm is that we don't need to consume the whole input and the whole stack.
What simply suffices is that we don't get stuck on the way. The result should be a distribution, containing only tokens
conforming to the grammar.
There is however one small caveat. Imagine sampling from the following distribution:
Initializing the stack again with the INT rule, and running the parsing, we end up just with -42 and - .
Now we sample, choosing the respective stack associated with the token. Let's say we sampled -42 .
To continue and get the next token, we can't just reset the stack with the INT rule, but have to remember
the position in the grammar and the corresponding stack.
You might still argue that we are dealing with this oracle machine, which somehow knows
if the rule [0-9] DIGITS or DIGITS is the right path. You would be right. Again, we can propose a simple solution.
Whenever there is a point where you have to make a decision, just go both ways.
Concretely, that means that when finding a rule like DIGITS | [0-9] DIGITS you just fork the current stack you have
and in one go with DIGITS , whereas the other one will have [0-9] DIGITS . Checking if a token goes through
is then about answering "Is there a stack which accepts?". Since all tokens are of finite length, you don't have to care
about infinite recursion (having a stack that would always expand DIGITS branch).
And that's it! Conceptually, this is how llama.cpp works!
Bear in mind that the llama.cpp solution is more on the side of slow solutions, as it does everything
at inference time. In the end, you incur something like O(vocab_size * active_stacks) cost,
as you are keeping stacks for each token and for each different route through the grammar.
With current vocab sizes (100k+) and longer grammars, that can be just incredibly time-consuming.
Certainly, better solutions like XGrammar or

[truncated]
