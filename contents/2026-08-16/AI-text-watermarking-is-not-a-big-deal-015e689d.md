---
source: "https://www.seangoedecke.com/ai-text-watermarking-is-not-a-big-deal/"
hn_url: "https://news.ycombinator.com/item?id=49317892"
title: "AI text watermarking is not a big deal"
article_title: "AI text watermarking is not a big deal"
author: "meetpateltech"
captured_at: "2026-08-16T08:17:27Z"
capture_tool: "hn-digest"
hn_id: 49317892
score: 1
comments: 0
posted_at: "2026-08-16T08:00:38Z"
tags:
  - hacker-news
  - translated
---

# AI text watermarking is not a big deal

- HN: [49317892](https://news.ycombinator.com/item?id=49317892)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/ai-text-watermarking-is-not-a-big-deal/)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T08:00:38Z

## Translation

タイトル: AI テキスト透かしは大した問題ではない

記事本文:
AI テキスト透かしは大した問題ではない Sean Goedeke
AI テキスト透かしは大した問題ではない
クロード モデルの出力に隠し透かしを含める予定であるという Anthropic の最近の発表について、人々はかなり不満を抱いています。これは人類モデルからの大量流出につながるでしょうか?ウォーターマークの導入はユーザーにとって有意義な変化となるでしょうか?
いいえ、AI テキスト透かしは大したことではありません。テキストが悪化するわけでもなく、AI 出力が実際に検出可能になるわけでもなく、ユーザーのプライバシーを侵害するわけでもなく、関係なく、2027 年までに誰もがそうすることになるでしょう。
透かし入りのテキストは低品質ではありません
透かしの入ったテキストと透かしの入っていないテキストの間には、品質に大きな違いはありません。これについてはこちらで詳しく書きましたが、Google の SynthID-Text と Meta の TextSeal という 2 つの一般的な方法は、ユーザーにとって完全に透過的です。これらは、擬似ランダム ロジット サンプラーを別の擬似ランダム ロジット サンプラーに置き換えることによって機能します。
友達とコイン投げでギャンブルをしていて、コインを投げる代わりに次のようにすることにしたとします。
午前0時からの現在時刻を秒単位で確認します
ブリタニカ百科事典でその数の単語を数えてください
到達した単語の文字数が偶数か奇数かを数えます 1
それでもギャンブルとしては十分ランダムですよね？しかし、透かしのように、各「コイン投げ」の正確な時間を記録している限り、理論的にはその手法が使用されたことを特定することができます。テキストの透かしも同様に機能します。つまり、事後的に検出できる「ランダム性」の方法が選択されます。透かし入りのモデルは、透かしなしのモデルよりも機能が劣ることはありません。
モデルが何かを引用している場合、または数学の専門家に答えを与えている場合はどうなるでしょうか。

それとも、出力がほぼ事前に決定されている何か他のことをしているのでしょうか?そこにウォーターマークを強制すると出力が悪化するのではありませんか?それはそうなるでしょう。だからこそ、どの AI 研究所もそれをやろうとしないのです。テキスト透かしアプローチは、ロジット サンプラーの既存のランダム性を置き換えるだけです。モデルが常に同じトークンを選択する場合、基本的にランダム性を考慮する必要がないため、それらのトークンには検出可能な透かしは存在しません。
これはすべて、以前は最高のトークンを取得していたのに、今はウォーターマークを満たす低品質のトークンを取得しているのではないかという不安から来ていると思います。たとえば、Anthropic の発表では、「曇り」と「灰色」の間の決定などの選択肢に透かしが表示されることが示唆されています。予想どおり、多くの人が、このような決定は優れた文章を書くために非常に重要であり、これらの言葉が同じであると考えるのは文盲の技術者だけである、と言いました。
これは、Anthropic の立場と透かしの仕組みについての誤解です。具体的には、透かしのないモデルは「曇り」を選択し、透かしのあるモデルは「灰色」を選択することを示唆しているため、これは誤解です。これはやり方ではありません! Claude Fable が特定のコンテキスト (たとえば、80% から 20%) で「灰色」よりも「曇り」を好む場合、透かし入りモデルと透かしなしモデルの両方で 20% の確率で「灰色」が得られます。創造性を促進するために、モデルにはすでに健全な量のランダム性が含まれています。テキストの透かしは、事後的に検出可能なランダムな選択を行う方法を導入しただけです。
AI 出力にはすでに「透かし」が入っています
AI 透かしを心配しなくてもよいもう 1 つの大きな理由は、AI テキスト コンテンツには常に効果的に透かしが入っていることです。最も慎重な地域

AI の出力を読んでいる人は、いつ AI の出力を読んでいるのかを知ることができます。これは、言語モデルが言語の特定の習慣 (エムダッシュ、修辞的反対、パンチの効いた一文、「クローデーゼ」など) に引き寄せられる傾向があるためです。実際、AI と人間の書き込みを確実に区別する分類子モデルをトレーニングすることが可能です。
私の知る限り、透かし入れに対する反発の一部は、AI 推論を自分の作品であるかのように装うために AI 推論を購入し、透かし入れによって透かし入れが困難になるのではないかと心配する人々から来ているようです。これらの人々にとって、透かし入りの発表は、「おい、君を賢く見せる代わりに、公の場で君を AI ユーザーだと決め付けて、バカに見せてやろう」という人類の言葉に似ている。
しかし、もちろんこれは常にそうでした！現在、AI の出力を自分のものとして偽装して逃れようとしている人は、透かしによって捕まることはありません。ほとんどの場合、この記事を読んだ人にとって、何が起こっているのかはすでに痛いほど明らかです。 「ハウス スタイル」を避けている洗練された AI ユーザーにとって、Anthropic の透かし検出器に自分の内容を貼り付ける不審な読者は、すでにそれを Pangram に貼り付けている可能性があります。
Pangram 2 のようなツールでは、出力が AI によって生成される可能性の推定値しか得られません。透かしを入れたほうが確実な確認になるのではないでしょうか?あまり。 SynthID によって選択されたトークンは理論的には人間によって選択された可能性があるため、テキストの透かしも確率的です。 Anthropic の透かしページは、ソースから直接提供されているため、パングラムよりも信頼できると考えられると思いますが、場合によっては、パングラムが実際に透かしよりも AI によって生成されたテキストの識別に優れている可能性があることも不可能ではありません。
AI テキスト透かしはプライバシーの侵害ではありません
私も見ました

ウォーターマークは機密コンテンツを出力にエンコードするか、何らかの方法で出力に個人情報のタグを付けるという理論が浮上しています。 AI ラボがデータを出力にエンコードするためにウォーターマークを使用しているとは思いません。テキストの透かしは難しいです。先ほども言いましたが、モデルが同じ単語でしか応答できない場合は実行できません。非常に短い応答では機能せず、長い応答でも確率的なフィンガープリントしか提供できません。そして、それは情報の 1 つのビット 3 をエンコードしているのです。
長いメッセージをウォーターマークにエンコードすることが技術的に不可能だと言っているわけではありません。それが機能する方法を説明した論文もあります。しかし、どの研究室でもそれを実行しているわけはありません 4 。そんなにあなたとあなたの回答を結びつけたいなら、彼らは生成したすべての模範的な回答を秘密裏に保存するだけでしょう。
AIテキスト透かしは避けられない
個々の AI ラボが透かしを入れることに対してあまり怒らないもう 1 つの理由は、今年はすべての AI ラボがテキストの透かしを行う予定であるということです。それは単なる人類的なものではありません。代替案は、EU AI 法を理由に EU でのビジネスを完全に停止することです。それは現在600億ドルの市場です。私は弁護士ではありませんが、AI ラボが EU の回答に透かしを入れるだけのようなことを合法的に行うことができるかどうかは、まったく不明瞭に思えます。完全に異なる claude-eu.ai サービスがあることを除けば、この法の平文は、そのサービスが特に EU 国民に出力するコンテンツだけでなく、EU 内で提供されるあらゆるサービスに適用されるように思えます。
もし人々が本当に透かしを嫌うのであれば、一部の研究室は完全に別個の EU サービスを立ち上げるか、EU AI 法を積極的に解釈して法廷闘争の行方を見守るかもしれない。反水に最大限の慈善活動をしようとすると、

歴史主義を批判する私は、次のような解釈を採用します。人々がウォーターマークはプライバシーの侵害であり、出力を悪化させるなどと言っているのは、それを信じているからではなく、完全に別のインターフェイスの背後で EU AI 規制をファイアウォールするよう AI 研究所に圧力をかけようとしているからです。この場合、テキストの透かしは大したことではないので、おそらく問題ではありませんが、米国の消費者が将来のより積極的な規制を懸念し、できるだけ早く明確な線を引きたいと考えているのはわかります。
興味深いことに、これは非常にわずかに有利である可能性があります。
私はパングラムのスポンサーではありません。私の理解では、Pangram は現時点で群を抜いて最高の AI 検出ツールです (競合他社の多くが怪しげで、有料の「AI 検出回避」サービスを促進するために存在していることも理由の 1 つです)。
技術的には、これは「ゼロビット ウォーターマーク」と呼ばれます。これは、ウォーターマーク自体から (単にウォーターマークの存在だけから) Yes または No の値を復元することができないためです。
AI ラボがユーザーごとの秘密キーを使用してテキストに透かしを入れ、キーを反復処理するだけで誰が何を生成したかを把握できると提案している人を見かけました。どのような規模でもこれをどのように実行できるのかわかりません。透かし検出はモデル推論よりも安価ですが、それでも (a) 信じられないほど計算量が多く、(b) おそらく偽陽性率が十分に高いため、数億のユーザーに対する実行は複数の人物と一致するでしょう。
この投稿を気に入っていただけた場合は、私の新しい投稿に関する更新情報を電子メールで購読するか、 Hacker News で共有することを検討してください。
これは、この投稿とタグを共有する関連投稿のプレビューです。
テキスト AI 透かしは常に簡単に削除できます
欧州連合 AI 法は、今から 1 か月後の 2026 年 8 月に施行されます。の 1 つ

最大の新たな要件は第 50 条で、すべての AI 出力が「人工的に生成されたものとして検出可能」であることを要求しています。言い換えれば、LLM プロバイダーが EU でビジネスを行いたい場合は、出力にウォーターマーク (AI コンテンツを識別するために使用できる隠された署名) を適用する必要があります。
続きを読む...
購読する │ About │ ポッドキャスト │ 人気 │ タグ │ RSS

## Original Extract

AI text watermarking is not a big deal sean goedecke
AI text watermarking is not a big deal
People are pretty unhappy about Anthropic’s recent announcement that they’re planning to include a hidden watermark in Claude model outputs. Will this lead to a mass exodus from Anthropic models? Will the introduction of watermarking be a meaningful change for users?
No. AI text watermarking is not a big deal. It doesn’t make the text worse, it doesn’t make AI outputs more detectable in practice, it doesn’t violate user privacy, and everyone’s going to be doing it by 2027 regardless.
Watermarked text is not lower-quality
There is no meaningful difference in quality between watermarked and unwatermarked text. I wrote about this more here , but the two popular ways to do it — Google’s SynthID-Text and Meta’s TextSeal — are completely transparent to the user. They work by replacing the pseudo-random logit sampler with a different pseudo-random logit sampler.
Suppose you were gambling on coin flips with your friends, and instead of flipping a coin you decided to do this:
Check the current time since midnight in seconds
Count that many words forward in the Encyclopaedia Britannica
Count whether the word you land on has an even or odd number of letters 1
That would still be random enough to gamble with, right? But, like a watermark, you could theoretically go back and identify that that method was used, so long as you recorded the exact time of each “coin flip”. Text watermarking works the same way: it chooses a method of “randomness” that can be detected after-the-fact. Watermarked models will not be any less capable than unwatermarked models.
What about cases where the model is quoting something, or giving you the answer to a mathematical problem, or doing something else where the output is largely pre-determined? Wouldn’t enforcing a watermark there make the output worse? It would, which is why none of the AI labs are going to do that. Text watermarking approaches only replace the existing randomness in the logit sampler: in any case where the model is always going to pick the same tokens, there’s basically no randomness to play with, so there won’t be a detectable watermark in those tokens.
I think all this comes from a worry that you were previously getting the best token, but now you’re getting a lower-quality token that satisfies the watermark. For instance, Anthropic’s announcement suggested that the watermarking is visible in choices like the decision between “overcast” and “grey”. Many people have predictably come out to say that decisions like these are really important to good writing, and that only an illiterate tech bro could think these words are identical.
This is a misunderstanding of Anthropic’s position and of how watermarking works. Specifically, it’s a misunderstanding because it suggests that the unwatermarked model would choose “overcast” while the watermarked one would choose “grey”. This is not how it works! If Claude Fable prefers “overcast” to “grey” in a particular context (say, 80% to 20%), you’ll get “grey” 20% of the time from both the watermarked and unwatermarked model. Models already include a healthy amount of randomness in order to promote creativity. Text watermarking just introduces a way to make those random choices that’s detectable after the fact.
AI outputs are already “watermarked”
The other big reason to not worry about AI watermarking is that AI text content has always effectively been watermarked . Most careful readers can tell when they’re reading AI outputs , because language models tend to gravitate towards certain habits of language : em-dashes, rhetorical opposition, punchy one-liners, “claudese”, and so on. In fact, it’s possible to train classifier models that reliably distinguish AI from human writing.
From what I can tell, some of the backlash to watermarking comes from people who buy AI inference in order to pass it off as their own work, and who worry that watermarking will make it harder for them to do that. For these people, the watermarking announcement is akin to Anthropic saying “hey, instead of making you seem smart, we’re going to publicly brand you as AI users and make you seem dumb”.
But of course this has always been the case! Nobody who is currently getting away with passing off AI outputs as their own will be caught by watermarking. For the majority of cases, it’s already painfully clear what’s happening for anyone who reads the slop . For sophisticated AI users who are avoiding the “house style”, any suspicious readers who would paste their stuff into Anthropic’s watermark detector could already have been pasting it into Pangram .
Tools like Pangram 2 only give you an estimate of the chance that output is AI-generated. Wouldn’t a watermark be a more solid confirmation? Not really. Text watermarks are probabilistic too, because any token chosen by SynthID could theoretically have been chosen by a human. I suppose the Anthropic watermark page could be considered more trustworthy than Pangram, because it comes right from the source, but it’s not impossible that in some cases Pangram might actually be better at identifying AI-generated text than the watermarking too.
AI text watermarking is not a violation of privacy
I’ve also seen theories floating around that watermarking encodes secret content into your outputs, or somehow tags outputs with your personal information. I don’t think AI labs are using watermarks to encode data into your outputs. Text watermarking is hard : like I just said, you can’t do it when the model can only respond with the same words, it doesn’t work for very short responses, and even on long responses it can only provide a probabilistic fingerprint. And that’s encoding one single bit 3 of information!
I’m not saying that encoding longer messages into a watermark is technically impossible — there are papers describing ways it might work — but there’s no way any of the labs are doing it 4 . If they wanted to associate you with your responses that badly, they’d just secretly store every model response they generated.
AI text watermarking is inevitable
Another reason to not get too angry at any individual AI lab for watermarking is that every single AI lab is going to do text watermarking this year . It won’t just be Anthropic. The alternative is to completely stop doing business in the EU, because of the EU AI Act . That’s currently a sixty-billion-dollar market. I am not a lawyer, but to me it seems genuinely unclear whether an AI lab could even legally do something like only watermarking EU responses: short of having an entirely different claude-eu.ai service, the plain text of the Act seems like it applies to any service offered in the EU , not just the content that service outputs to EU citizens specifically.
If people really hate watermarking enough, some labs might stand up a completely separate EU service, or make an aggressive interpretation of the EU AI Act and see how the legal battle goes. When I try to be maximally charitable to anti-watermarking histrionics, I adopt an interpretation like this: people are saying that watermarking is an invasion of privacy and makes outputs worse and so on not because they believe it, but because they’re trying to pressure AI labs to firewall EU AI regulations behind a completely separate interface. In this case, it probably doesn’t matter — text watermarking is not a big deal — but I can see an American consumer being worried about more aggressive future regulation, and wanting to draw a firm line in the sand as early as possible.
Interestingly, this might be very slightly even-favored.
I am not sponsored by Pangram. As I understand it, Pangram is by far the best AI-detection tool right now (in part because many of its competitors are shady and exist to promote paid “AI-detection-evasion” services).
Technically, this is called a “zero-bit watermark”, because you can’t recover a yes-or-no value from the watermark itself (merely from the presence of a watermark).
I saw someone suggesting that an AI lab could use per-user secret keys to watermark text, and then simply iterate over the keys to figure out who generated what. I just don’t see how you could do this at any scale: watermark detection is cheaper than model inference, but it’s still (a) computationally intensive enough to be implausible, and (b) probably has a high enough false-positive rate that any run against hundreds of millions of users would match multiple people.
If you liked this post, consider subscribing to email updates about my new posts, or sharing it on Hacker News .
Here's a preview of a related post that shares tags with this one.
Text AI watermarks will always be trivial to remove
The European Union AI Act will begin to be enforceable in August 2026, one month from now. One of the biggest new requirements is Article 50 , which requires all AI outputs to be “detectable as artificially generated”. In other words, if LLM providers want to do business in the EU, they will have to apply a watermark to their outputs: some hidden signature that can be used to identify AI content.
Continue reading...
subscribe │ about │ podcasts │ popular │ tags │ rss
