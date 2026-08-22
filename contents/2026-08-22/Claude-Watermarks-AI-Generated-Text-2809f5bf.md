---
source: "https://magazine.sebastianraschka.com/p/claude-watermarking"
hn_url: "https://news.ycombinator.com/item?id=49400307"
title: "Claude Watermarks AI-Generated Text"
article_title: "How Claude Watermarks AI-Generated Text"
image: "https://substackcdn.com/image/fetch/$s_!2E7U!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Ffb73bfb5-0291-4e5e-b425-6c3be054276e_1914x1076.png"
author: "vismit2000"
captured_at: "2026-08-22T15:11:34Z"
capture_tool: "hn-digest"
hn_id: 49400307
score: 1
comments: 0
posted_at: "2026-08-22T14:45:00Z"
tags:
  - hacker-news
  - translated
---

# Claude Watermarks AI-Generated Text

- HN: [49400307](https://news.ycombinator.com/item?id=49400307)
- Source: [magazine.sebastianraschka.com](https://magazine.sebastianraschka.com/p/claude-watermarking)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T14:45:00Z

## Translation

タイトル: クロードの透かし AI 生成テキスト
記事のタイトル: クロードが AI 生成テキストに透かしを入れる方法
説明: トークンのサンプリング、透かしの検出、削除に関する 48 分間のビデオ ウォークスルー

記事本文:
クロードが AI 生成テキストに透かしを入れる方法
クロードが AI によって生成されたテキストに透かしを入れる方法
トークンのサンプリング、透かしの検出、削除に関する 48 分間のビデオ ウォークスルー
Sebastian Raschka、PhD 2026 年 8 月 22 日 61 1 5 シェア 最近、Claude の新しい透かしプロセスと実装に関する Substack ノートを投稿しました。これは非常に人気のあるトピックであり、非常に活発な議論を引き起こしたため、その仕組みを説明するときにもう少し詳しく説明するのは興味深いかもしれないと思いました。
通常のテキスト記事の代わりに、（通常の記事とは少し趣向を変えて）このテーマに関するちょっとした講義を録画しました。ということで、以下はビデオとトランスクリプトです。
当初は 10 枚のスライドを作成し、10 分の短いビデオを録画する予定でした。ただし、まとめているときに重要な詳細をあちこちに追加したため、スライドが 50 枚を超え、録音が 48 分になりました。
ただし、これでうまく説明できると思います。楽しく見てください!
YouTube プレーヤーを使用したい場合は、YouTube バージョンもあります
そして、ここにスライドへのリンクがあります
注: 以下のトランスクリプトは、読みやすくするために若干編集および整理されていますが、上記のビデオ講義の全体的な順序と流れは維持されています。
クロードのテキスト透かしの仕組み
スライド 2/52、タイムスタンプ 0:00 皆さん、こんにちは。そこで、数日前、Anthropic は、Claude モデルのテキスト出力に透かしを入れると発表しました。次に、それがどのように機能するかを簡単に説明するソーシャルメディア投稿をしました。はい、これは非常に人気のある投稿でした。つまり、透かしそのものが人気だったわけではなく、その背後にある説明やメカニズムが人気だったのだと思います。次に、この投稿には図が 1 つしかなく、多くの質問や議論があったため、これを少し拡張してより詳細に説明する価値があるかもしれません。
それで、もう少し作ってみようと思いました

e数字。実は当初は 10 枚ほどのスライドを作成して説明する予定でした。最終的には 50 枚のスライドになりましたが、この透かし技術がどのように機能するか、透かし自体がどのように失敗したり削除されたりする可能性があるかなどが、これでよく説明できれば幸いです。最近、多くの人が LLM を使用しており、LLM によって生成される可能性のあるインターネット上のテキストを大量に消費しているため、これは興味深いトピックかもしれないと思います。
そして今度はこの透かしが入ることになるのですが、透かしによってテキストが悪化するのではないかという不安、あるいはこの透かしの実際の利点は何でしょうか?それで、それは何を意味するのでしょうか？ウォーターマークとは何かをもう少し理解できれば、それは大いに役立つと思います。そうすれば、それが良いことなのか悪いことなのか、メリットとデメリットなどについて自分で判断できるようになります。
したがって、ここでの私の目的は、基本的なメカニズムがどのように機能し、このタイプの透かし、テキスト透かしをどのように実装するかを説明することです。
スライド 2/52、タイムスタンプ 1:41 これは、物事をゼロから理解することが実際に非常に役立つ理由を説明する優れた例でもあります。この透かし技術は、従来のモデルや LLM が一般に内部でどのように機能するかを説明するのに適した方法でもあります。そうです、あなたは私が物事を一からやるのが好きなことを知っているかもしれません。たとえば、私には『Build a Large Language Model From Scratch』、『Build a Reasoning Model From Scratch』という本があります。ゼロからラベルを付けた記事がいくつかあります。
つまり、私にとって、「スクラッチにはコーディングが含まれることがよくあります。したがって、これはコーディングとは関係ありませんが、何かがどのように実装されているかを理解するのに本当に役立つため、スクラッチからコーディングすることは実際には非常に便利なテクニックです。そして、そこから私たちの理解、図、概念を導き出すことができます。なぜなら、実際に物事を実装しない場合、もし存在する場合は」

コードがないので、本当に曖昧な場合があります。そしてもちろん、私も気づいたように、誰もがゼロからコーディングしているわけではありません。
昔のように、何かをゼロからコーディングすることだけが私たちにできるすべてでした。つまり、コーディングしているのは人間だけでした。現在、コーディングは LLM によって実行できます。ただし、コードには多くの情報が含まれているため、コードを読むことが役に立たなくなったというわけではありません。この場合、この透かしを使用して LLM を最初からコーディングすることに時間を費やすと、このサンプリングが内部でどのように実装されているかがよくわかります。関連するコード スニペットがまだいくつかあります。
そして、それは実際に、「ああ、この位置に透かしが適用されていて、これが何らかの影響を与える」ということを理解するのに役立ちます。したがって、人々がゼロからコーディングすることは、少なくとも常にはなくなったとしても、たとえば、教育目的で何かを深く理解するために、また研究目的でそれを大量の抽象化レイヤーに隠されない透過的な方法で操作するために、何かをゼロから構築できることは依然として有用だと思います。
しかし、それはさておき、このスライドデッキでは実際に最初から作成したコーディング資料から多くの図を使用したため、これは単なる偶然の素晴らしい関係だと思います。
スライド 3/52、タイムスタンプ 3:58
それで、数日前 (これは 8 月 14 日です) に、「クロードのテキスト透かしの仕組み」という記事があり、ここにも記事がありました。これは画面録画のようなものなので、スライドにすべてを含めることができますが、詳細はたくさんあります。実際には数回更新されているので、私がこれを読んだ当初はもっと短かったです。それでも、それは非常に概念的なものだと私は思います。このような概要がありますが、つまり、そこには人物が一人もいません。
それで、それは一種の聖です

彼らが何をしようとしているのか理解するのは難しいです。そのため、彼らはなぜそれを行うのかについては多くのことを説明しますが、その方法については説明しません。彼らはどこかにある 1 つの論文にリンクしていますが、これも非常に技術的です。したがって、一歩下がって最初から始めて、この透かし技術を使用して彼らがここで何を実装しようとしているのかを理解するのが理にかなっていると思います。
ちなみに、透かしを入れる動機は、誰かが「ああ、このテキストはクロード Opus 4.8 モデルによって生成されたものだ」と言えるようなテキストを投稿したかどうかを識別して、「この透かしが含まれているので、このテキストは AI によって生成されたものだ」と判断できるようにすることです。そして、この透かしはユーザーには見えないため、それをデコードしてテキストに透かしがあるかどうかを確認できるのはユーザーだけです。
なぜ彼らだけがそれができるのでしょうか？これについては後ほど説明します (長すぎるビデオにならないことを願っています) が、一度に 1 つずつ説明します。
スライド 4/52、タイムスタンプ 5:35 そこで、LLM でテキスト生成がどのように機能するかを説明するための簡単な前置きから始めたいと思いました。これに基づいて、電子透かしがどのように機能するかをより簡単に理解できるため、これは実際には、それに加えて大規模で高価なものではないことを理解できるからです。これは実際には、通常のテキスト生成プロセス内の小さな微調整のようなものだと思います。
スライド 5/52、タイムスタンプ 6:01
たとえば、ChatGPT のようなものを使用しているときに、ドイツの首都は ChatGPT または他の LLM であると質問したとします。これは、たとえば「ベルリン」と答える例と同じです。この場合、「Berlin」とピリオドのような 2 つのトークンが生成されます。ただし、話を簡単にするために、トークンが 1 つ生成されていると仮定します。したがって、次のトークンは「ベルリン」トークンです。このトークンは内部でどのように生成されるのでしょうか?
下で何が起こっているのか

ドイツの首都のようにここに何かを入力すると、「ベルリン」のようなトークンが返されるのですか?舞台裏で実際に何が起こっているのでしょうか？
スライド 6/52、タイムスタンプ 6:41 そこで、次のいくつかのスライドでは、この次のトークンが生成されるときに内部で何が起こるかについて簡単に説明したいと思います。
スライド 7/52、タイムスタンプ 6:50 そこで、もう一度プロンプトがドイツの首都であると仮定します。ここでの最初のステップは、これをトークン ID に変換することです。したがって、トークン化してトークン ID に変換することは、最初の主要な手順の 1 つです。ここは外です。 LLM の内部にはありません。それはLLMの範囲外です。したがって、テキストをトークン ID に変換するだけです。これは、埋め込みレイヤーが使用できる形式にすぎません。
スライド 8/52、タイムスタンプ 7:22 そして、これは LLM を通過します。そして、LLM は次のトークンのスコア分布を提供します。
スライド 9/52、タイムスタンプ 7:31 繰り返しになりますが、これは LLM が内部でどのように機能するかを簡単に概要を示したものです。したがって、LLM 機構自体については取り上げません。これについては、他の From Scratch LLM のビデオや書籍で何度も話しました。重要な部分は、次のトークン (たとえば、「ベルリン」) を生成するときに、この時点でスコアの分布が得られるということです。これが LLM によって生成された出力です。
この場合、ロジット値を調べています。したがって、これらは、スコアの範囲のような、マイナス無限大からプラス無限大までの単なるスコアです。これは、約 -8 または -9 から 20 までの範囲の例です。これらを確率分布に変換できますが、技術的には、サンプリング方法によっては厳密に必要というわけではありません。ただし、ロジット値は生のスコアと考えることができます。
そして、生のスコアは語彙全体に及びます。
スライド 10/52、タイムスタンプ 8:39 つまり、LLM が認識できるすべての単語を意味します

生成する可能性があります。ここで、語彙インデックスでは、ある値 (インデックス位置 19,846) が最も高いスコアを獲得します。それで配布を広げました。このプロンプトを LLM 経由で実行すると、さらに極端な結果が表示されることもあります。つまり、すべてが非常に非常にゼロに近いということです。そして、「ベルリン」はおそらくそれをはるかに上回るでしょう。
でも、ここにあるピークのようなものをいくつかお見せするために、もう少し面白く見えるように、ズームインしてみました。で分布を少し広げます。ここで、「ベルリン」が最高スコアです。このような非常に具体的なプロンプトがある場合、それが最も可能性が高い、またはもっともらしい次のトークンであると考えることができるからです。つまり、その他の都市は、LLM が誤って推測する可能性があるハンブルクやミュンヘンのような都市である可能性があります。
しかし、今日では、LLM はここで「ベルリン」が正しい答えであるとかなり確信しているはずです。ここには語彙索引も表示されています。つまり、語彙全体を網羅しているようなものです。現在、LLM は出力として 250,000 個ほどのトークンを生成できます。このスライドではスペースが非常に多いため、ここでは 19,800 から 19,900 までを切り捨てています。もし私が 250,000 語という非常に現実的な語彙を持っていたとしたら、すべてが非常に狭くなり、この分布ではほとんど何も伝えることも見ることもできなくなるでしょう。
したがって、これは教育目的で省略されているだけです。重要な点は、通常のテキスト生成ではこのスコア分布が得られるということです。ここで、最高スコアを確認します。
スライド 11/52、タイムスタンプ 10:33 これがどのように選択されるかについては後で詳しく説明します。したがって、必ずしも正確に最高のスコアである必要はありませんが、簡単にするために、ここでは最高のスコアを取得していると仮定します。この場合、それは 19,846 です。
スライド 12/52、タイムスタンプ 10:52 そして、このスコアはデトになります

ケンナイズされ、「ベルリン」が戻ってきます。これがこのスライドのプロセスです。入力プロンプトからトークン ID への変換とトークン化、それを LLM に渡し、このスコア分布を取得し、次のトークンを取得してテキストに変換し直します。
スライド 13/52、タイムスタンプ 11:12 そして、このテキストが入力に追加されます。したがって、複数の出力トークンを必要とする質問がある場合は、回答が完了するまでこのループを続けます。これは通常、LLM がテキスト終了トークンを生成することを意味します (例: ここ)。わかりやすくするために、1 つのトークンを生成する反復を 1 回だけ示します。しかし、そうですね、先ほども言いましたが、次のラウンドのために修正された入力を LLM にフィードバックするというような感じで継続します。
さて、ここで実際に次のトークンをサンプリングするにはどうすればよいでしょうか?
スライド 14/52、タイムスタンプ 11:44 私は簡単に言った、そうですね、技術的には最も高いもの、つまり最も高いスコアを持つものを選択するだけです。これは貪欲デコードと呼ばれます。それも一つの方法です。しかし、ほとんどの LLM は、使用する場合と同様に、常に最高のものを選択する貪欲なデコードを行いません。他のプロンプトで質問すると、トレーニング データが記憶されてしまうため、常に最高のスコアを取得することは望ましくない可能性があります。
常に同じような応答が返ってきます。それで私たちはACします

[切り捨てられた]

## Original Extract

A 48-minute video walkthrough of token sampling, watermark detection, and removal

How Claude Watermarks AI-Generated Text
Subscribe Sign in How Claude Watermarks AI-Generated Text
A 48-minute video walkthrough of token sampling, watermark detection, and removal
Sebastian Raschka, PhD Aug 22, 2026 61 1 5 Share I recently posted a Substack note about Claude’s new watermarking process and implementation. Since it’s such a popular topic and sparked such a lively discussion, I thought it might be interesting to go into a bit more detail when explaining how it works.
Instead of the usual text article, I recorded a little lecture on the topic (to change it up a bit from my usual articles). So, below is the video along with a transcript.
Originally, I planned to make 10 slides and record a short 10-min video. However, while putting it together, I added some crucial details here and there, resulting in >50 slides and a 48 min recording.
I hope that this now explains it well, though! Happy watching!
I also have a YouTube version if you prefer using the YouTube player
And here is a link to the slides
Note: The transcript below is slightly edited and cleaned up for readability but preserves the overall order and flow of the video lecture above.
How Claude’s Text Watermarking Works
Slide 2 of 52, time stamp 0:00 Hi everyone. So, a few days ago, Anthropic announced that they will watermark the text outputs of their Claude models. I then did a social media post briefly explaining how that works. And yeah, this was quite the popular post. So not the watermarking itself was popular, but I guess the explanation or the mechanism behind it. Then, it might be worthwhile expanding this a bit to explain it in more detail, because this post only had one figure, and there were a lot of questions and discussions.
So, I thought, well, let’s make a few more figures. I actually originally planned to do like 10 slides and walk you through it. It ended up being 50 slides, but I hope this really explains how this watermarking technique works well, how watermarking itself can fail or be removed, and so forth. So I think it might be an interesting topic because a lot of people use LLMs these days and also consume a lot of text on the Internet that might be generated by LLMs.
And now there’s going to be this watermarking, and there’s this, I guess, fear of watermarking making text worse, or what’s actually the benefit of this watermarking? And so what does it mean? And I think if we understand a bit better what watermarking is, that goes a long way, and then we can make up our own minds about whether that’s a good thing or not, and so forth, like the pros and cons.
So, my goal here is really to explain how the underlying mechanism works and how they are going to implement this type of watermarking, text watermarking.
Slide 2 of 52, time stamp 1:41 It’s also a great example to illustrate why understanding things from scratch is actually quite useful. This watermarking technique is also a nice way to explain how conventional models or LLMs in general work under the hood. So yeah, you may know I like doing things from scratch. Like, I have my books: Build a Large Language Model From Scratch, Build a Reasoning Model From Scratch. I have some articles labeled from scratch.
So, for me, “ from scratch often includes coding. So this one will not be coding-related, but coding from scratch is actually a very, very useful technique because it really helps you understand how something is implemented. And then from that we can derive our understanding, figures, concepts, because if we don’t really implement things, if there’s no code, it’s really sometimes ambiguous. And of course, you know, as I realized, not everyone is coding from scratch anymore.
Like back in the day, coding something from scratch was all we had. I mean, there were only humans coding. Nowadays, coding can be done by LLMs. However, that doesn’t mean reading code is no longer useful, because it carries a lot of information. So in this case here with this watermarking, spending some time coding an LLM from scratch really makes you realize how this sampling inside is implemented. We still have some relevant code snippets.
And then that really, in turn, helps us understand, oh, the watermarking is applied at this position, and this has so-and-so consequences and so forth. So I think even though people may not be coding from scratch, at least not all the time anymore, it is still useful being able to, let’s say, build something from scratch for educational purposes to understand something deeply and then also for research purposes to manipulate this in a transparent way that is not hidden away in tons of layers of abstraction.
But that aside, I think it’s just a coincidental nice relationship here because, for this slide deck, I actually used a lot of figures from my from-scratch coding materials.
Slide 3 of 52, time stamp 3:58
So a few days ago (this is August 14), there was this article, How Claude’s Text Watermark Works, and there was this article here; it’s just like a screen recording, so it can have everything in the slides, but there’s plenty of detail. They updated it actually a couple of times, so originally when I read this, it was a way shorter. Still, it is very, I guess, conceptual; there’s like this overview, and there’s, I mean, there’s not a single figure in there.
And so it’s kind of still hard to understand what they’re trying to do. So they explain a lot about why they’re going to do it, but they don’t explain how. They’re linking to one paper somewhere there, which is very technical also. So I do think it makes sense maybe to take a step back and start at the beginning to kind of understand what they’re trying to implement here with this watermarking technique.
And so the motivation, by the way, of watermarking is for them to identify if someone posts some text that they can say, oh, this text was generated by our Claude Opus 4.8 model, for example, so that they have a way to tell, OK, this text is AI-generated because it carries this watermark. And this watermark is invisible to users, so only they can decode it and find out whether the text has their watermark.
Why can only they do it? We will get to that later in this (hopefully not too long a video), but one thing at a time.
Slide 4 of 52, time stamp 5:35 So I wanted to start with a brief prelude to explain how text generation works in LLMs, because based on that we can then more easily understand how the watermarking works and that this is actually not a huge, expensive thing on top of it. It’s really just like a minor, I guess, tweak inside the regular text generation process.
Slide 5 of 52, time stamp 6:01
So when we are using something like ChatGPT, for example, let’s say I ask the question, the capital of Germany is, and yeah, ChatGPT or other LLMs, so this is just like an example would, for example, answer “Berlin”. So here, in this case, it’s generating two tokens, like “Berlin” and the period. But for simplicity, let’s assume it’s generating one token. So the next token is the “Berlin” token. How is this token generated internally?
What is happening under the hood when we type something here like the capital of Germany is and receive a token like “Berlin” back? What is actually going on there behind the scenes?
Slide 6 of 52, time stamp 6:41 So in the next couple of slides, I want to briefly talk about what happens under the hood when this next token is generated.
Slide 7 of 52, time stamp 6:50 So assume again that our prompt is the capital of Germany is. And the first step here is to convert this into token IDs. So tokenizing it and converting it into token IDs is one of the main steps at the beginning. This is outside. It’s not inside the LLM; it’s outside of the LLM. So we are simply converting the text into token IDs. It’s just a format that embedding layers can work with.
Slide 8 of 52, time stamp 7:22 And then this passes through the LLM. And the LLM gives us a score distribution for the next token.
Slide 9 of 52, time stamp 7:31 So again, this is just like a brief overview of how LLMs work internally. So I’m not covering the LLM machinery itself. I talked about it many times in my other From Scratch LLMs videos and books. The important part is that when we generate the next token (for example, “Berlin”), we have, at this point, a distribution of scores. So this is the output produced by the LLM.
Here in this case, we’re looking at logit values. So these are just scores from minus infinity to plus infinity, like a range of scores. Here’s an example, ranging from about -8 or -9 to 20. We could convert these into a probability distribution, but technically, it’s not strictly necessary depending on how we sample. But so you can think of the logit values as the raw scores.
And the raw scores go over the entire vocabulary.
Slide 10 of 52, time stamp 8:39 That means every possible word that the LLM could generate. Now here, in the vocabulary index, a certain value (index position 19,846) receives the highest score. So I spread out the distribution. If you would run this prompt through an LLM, you would even see something more extreme: that everything is, like, very, very, very close to zero. And “Berlin” would probably be much, much higher even.
But just to show you a few, you know, like peaks here so it looks a bit more interesting, I kind of zoomed in; in and spread out the distribution a bit. Now here, “Berlin” is the highest score because you can think of it as the most, I guess, probable or plausible next token if I have a very specific prompt like this. So the other ones, I mean, it could be something like Hamburg or Munich that the LLM might guess incorrectly.
But nowadays an LLM should be fairly certain that “Berlin” is the correct answer here. You are also seeing here the vocabulary index. So that’s like over the whole vocabulary. Nowadays, LLMs have like 250,000 possible tokens as output. I’m truncating it here from 19,800 to 19,900 because there’s just so much space here on this slide. If I would have a very realistic vocabulary of 250,000 words, everything would be so narrow that we would barely even be able to tell or see anything on this distribution.
So this is just truncated for educational purposes. The important point is that in regular text generation, we get this score distribution. Now, what we do is look at the highest score.
Slide 11 of 52, time stamp 10:33 I will get into more detail later on how this is selected. So it’s not necessarily precisely the highest one, but for simplicity, assume we are taking the highest score here. And in this case, it’s 19,846.
Slide 12 of 52, time stamp 10:52 And this score is then detokenized, and we get “Berlin” back. So that is the process here on this slide: from an input prompt to conversion into token IDs and tokenization, passing it to the LLM, getting this score distribution, getting the next token, and converting it back into text.
Slide 13 of 52, time stamp 11:12 And then this text is appended to the input. So if we have a question that requires multiple output tokens, we keep going in this loop until the answer is complete. That usually means that the LLM generates an end-of-text token, for example, here. For simplicity, I’m showing you only one iteration where it generates one token. But yeah, as I said, it would kind of continue like that, where we are feeding back the modified input to the LLM for the next round.
Now, how do we actually sample this next token here?
Slide 14 of 52, time stamp 11:44 I briefly said, well, we could just technically select the highest one, the one with the highest score. This is called greedy decoding. That’s one way to do it. But most LLMs, like if you use them, they don’t do greedy decoding where they always pick the highest one. Because if you ask it on some other prompt, it might not be what we want to always have the highest score, because then it would memorize the training data.
It would always kind of give the same response and so forth. So we ac

[truncated]
