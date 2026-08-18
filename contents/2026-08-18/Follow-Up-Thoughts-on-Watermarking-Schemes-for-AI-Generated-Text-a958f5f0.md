---
source: "https://daringfireball.net/2026/08/follow-up_thoughts_on_watermarking"
hn_url: "https://news.ycombinator.com/item?id=49339918"
title: "Follow-Up Thoughts on Watermarking Schemes for AI-Generated Text"
article_title: "Daring Fireball: Follow-Up Thoughts on Watermarking Schemes for AI-Generated Text"
image: "https://daringfireball.net/graphics/df-wide-card.png"
author: "tambourine_man"
captured_at: "2026-08-18T02:10:05Z"
capture_tool: "hn-digest"
hn_id: 49339918
score: 4
comments: 0
posted_at: "2026-08-18T01:16:14Z"
tags:
  - hacker-news
  - translated
---

# Follow-Up Thoughts on Watermarking Schemes for AI-Generated Text

- HN: [49339918](https://news.ycombinator.com/item?id=49339918)
- Source: [daringfireball.net](https://daringfireball.net/2026/08/follow-up_thoughts_on_watermarking)
- Score: 4
- Comments: 0
- Posted: 2026-08-18T01:16:14Z

## Translation

タイトル: AI 生成テキストの透かし入れスキームに関するフォローアップの考察
記事のタイトル: Daring Fireball: AI 生成テキストの透かし入れスキームに関するフォローアップの考察
説明: 私は、読んだ回答が説得力があり、明快で、正確で、ありがたいことに簡潔であることを望みます。そして理想的には、読んでいる耳に心地よい一貫したトーンを発することです。魔神は瓶の中に戻らない。

記事本文:
Drata の Agentic Trust Management プラットフォームで GRC をより迅速に管理
AI 生成テキストの透かしスキームに関するフォローアップの考察
今週末のステムワインダー「クロードにおけるアンスロピックの「ウォーターマーク」テキスト不正行為は、ライティングの倒錯である」のフォローアップの一部:
Hacker News やその他の場所にいる愚か者たちとは対照的に、人気のある LLM が各決定点で単に「最良の」トークン (単語) を選択するわけではないことを私は理解しています。直感に反しますが、常に最も確率の高いオプションを選択すると、望ましくない結果が生じます。したがって、モデルはある程度のランダム化を適用します。「温度」は、「より良い」(モデルによってランクが高い) 選択肢が選択される可能性が高くなるように適用される重み付けを表す用語です。
温度が 1 の場合、モデルは組み込みの確率分布を使用します。温度が 1 より大きい場合、この分布はより平坦になります。つまり、可能性の低い代替案が選択される確率は高く、可能性の高い代替案は選択される確率が低くなります。温度が 1 より低い場合、確率分布はより上位のオプションに偏ります。温度が 0 の場合、最高ランクのオプションが常に選択されます。温度が 0 の場合、一般に望ましくない結果が生じます。予測可能すぎて、行き詰まる可能性が高すぎます。カメラ センサーからの画像を過度に平滑化するのと同様、すべてのノイズを除去すると、たとえ個別に評価された「ノイズ」の各ビットが何らかの意味で間違っていたとしても、全体の結果が悪化します。
温度ベースのランダム性 (LLM 出力を非決定的にするもの) は、出力を改善するために設けられています。散文は、温度 0 (ランダム性なし) よりも温度 1 (重み付きランダム性あり) の方が明らかに優れています。一方、透かし入れスキームは予測を適用しています。

秘密キーを使用したランダム性は、出力の品質を向上させることとはまったく異なる目的であり、したがって、本質的に出力を少なくともわずかに悪化させると私は信じています。
テキスト用の LLM 透かしスキームの支持者は、このスキームは温度を変更せず、ランダム性のソースを変更するだけであるため、生成される散文の品質を必ずしも低下させるわけではないと主張しています。今日、Daniel Jalkut がこれについて良い記事を書きました。それが本当だといいのですが。それが真実である可能性はあると思います。それが真実である可能性は非常に低いと思います。それが本当なら、彼らはそれが真実であることを証明する例を示すと思います。また、Anthropic 自身も、プログラミング言語コードであるテキストに適切に透かしを入れることができないことを認めています。
同じ理由で、コード (非常に多くの場合、次のようにする必要があります)
正確 — 通常、他の形式よりも透かしが少ない
テキストの。
とはいえ、任意の選択ができる領域では、
コード内の特定の単語または用語の間、ウォーターマーク
コード内のコメントなどに使用できます。しかし、定義上、それは
生成される実際のコードへの影響は無視できます。
私は、優れた散文はプログラミング コードに似ていると考えています。言葉の選択、言い回し、口調、さらには句読点さえも正確である方が、不正確であるよりも常に優れています。違いは、ずさんなプログラミング コードは実行されないか、正しく実行されないことです。一方、人間の脳は、不正確な、あるいは雑な散文を解析して意味を理解することに熟達しています。
品質に悪影響が及ばない場合でも反対します
私は、これらの計画が散文の質をわずかにでも低下させることなく機能するとは信じていません。しかし、繰り返しになりますが、私は間違いであることが証明されることを受け入れます。しかし、たとえそのような透かし入れスキームが必ずしも品質を低下させるわけではないことを現時点では認めたとしても、

生成された散文の数 (1 イオタではありません) が、ユーザーの背後で密かに適用される場合、私は依然としてその使用に反対します。便利な透かしは、誰でもチェックできる透かしです。これらの SynthID の「ウォーターマーク」は、LLM プロバイダー (これまでのところ、Anthropic/Claude および Google/Gemini) が保持する秘密に完全に依存しています。私はそれが受け入れがたいと思っています。その理由は私がエッセイで明らかにしたかったからです。
テキストにこの透かしを入れることに賛成している人たちは、夢物語、幻想を売られているのです。私は、AI 嫌いの人々 (特に Bluesky でのコメントが多いですが、Threads や Hacker News もそうです) からのコメントに何十件も遭遇しました。彼らは、AI が生成したテキストに透かしを入れることに反対できるのは、AI が生成したテキストを二枚舌で自分の文章だと偽っている人だけだと確信しています。したがって、私が怒るのは、すぐに私にも問題が起きるからだと考えているのです。もちろんこれは真実ではありません。私は、自分の仕事の一文どころか、テキスト メッセージや電子メールを書くのにも AI を使用しません。
しかし、LLM は「スロップ」を生成するだけで、何も有益なものは何も生成しないと信じていると主張する非常に多くの人が、同じ LLM が信頼性の高い方法で出力に透かしを入れることができると 100% 確信しているように見えるのは、面白いと思います。これらの人々は、AI が生成したテキストを指摘できるようになりたいと切望しているため、彼らのおかげでそれができるようになるという Google と Anthropic の主張に引っ掛かっています。
時間の無駄なので、これについてあまり考えたくありませんが、これらの人々は、これらの必須の透かしや検出ツールの存在が何かをより良い方向に変えると具体的にどのように考えているのでしょうか?あなたがオフィスで働いていて、多くの同僚が AI を使って絵馬を書いているのではないかと疑っているとします。

ils およびその他の仕事関連のメッセージ。彼らのメッセージは長すぎ、内容が多すぎ、明快さに欠けています。これから何をするつもりですか？各メッセージをコピーして、Anthropic、Google、OpenAI のウォーターマーク検出機能に貼り付けますか?すべての LLM に対して単一の検出器が存在することはできません。そして、一致するということがわかったとしても、電子メールやブログ投稿、あるいは Slack メッセージが、たとえばクロードによって生成された可能性が非常に高いことがわかったとしても、どうするつもりですか?同僚のオフィスに行進して、捕まえたと伝えますか？
「バレる」ことが重要な状況にある人、たとえば学生は、透かしの入っていない LLM を使用したり、透かし入りのテキストを Decclaude のような言い換えツールを使って実行したりするでしょう。
たとえこれらの透かし入れスキームが約束どおりに機能したとしても、これによって実際的な利益がもたらされることはありません（そして、明確にしておきますが、私はどれも約束どおりに機能するとは思っていません）。 1
私のアドバイスは、何かが AI によって書かれたのか人間によって書かれたのかを気にしないことです。評価に値する唯一のことは、私たち人間の読者が自然に判断できるもの、つまりそれが良いか悪いか、ということです。良かったら読んでみてください。そうでない場合は、やめてください。 AI が生成したメッセージで受信箱が埋め尽くされる同僚に囲まれて耐えられない仕事に就いている場合は、新しい仕事に就くか、その仕事に耐えることを学びましょう。隠された秘密の透かし信号は、たとえ機能したとしても、物事を以前の状態に戻すことはできません。何かを読んで楽しんだ後、それが LLM によって生成されたものだとわかったとしても、悪く思う必要はありません。良いものを読んで楽しかったです。
今日、LinkedIn の投稿のほとんどが AI によって生成されているという内容の記事を読みました。プラットフォーム全体が AI のスロップで溢れかえっているということ。そうかもしれませんが、LinkedIn を見たことがないのでわかりません

常にゴミで満たされているからです。糞が人間の肛門から出たものであろうと、糞を生成するロボットから出たものであろうと、それがクソのような臭いであれば、それはクソです。
人間だけが本当に書ける議論
今日、Six Colors で「LLM は書いていない」と執筆している Dan Moren 氏:
LLM は自分が選んだ単語を気にしません。
何も気にすることができない。
同じではない 2 つのことについて、ジョンは正しく指摘しています。
「彼はチャンスに飛びついた」というフレーズの違いを理解する
そして「彼はその機会に飛びついた」。これらは確かに別個のフレーズであり、たとえ意味的には似ていても、それぞれがより類似したものである可能性があります。
特定の状況に適しています。または、別の言い方をすると、次のようになります。
それぞれのフレーズを使うと、何か違うことが分かります。
描写されている人物についてであっても、作家についてであっても。
しかし、LLM にはこれらのフレーズのどれが正しいのかわかりません
使用するフレーズ。モデルと重量に基づいて推測があり、
入力。しかし、これらのフレーズの究極の選択は私たちに何も伝えません
作家がいないので作家について。
モーレン氏のコメントは私の投稿に対する優れた反論ですが、たとえ哲学的なレベルであっても、私は根本的に同意しません。書かれた作品を、それを生み出した心への洞察を得るためだけに読んでいるのであれば、AI が生成したテキストの向こう側には心は存在しません。しかし、作品自体は存在します。モーレンに対する私の意見の相違は、彼の（素晴らしく要約的な）見出しで始まり、事実上終わります。あなたが何かを読むことができるなら、それは必然的に書かれたものだと私は言います。
繰り返しますが、これは哲学的です。 AIによって生成されたフォトリアルな画像が撮影されたのでしょうか？いいえ、そうではなかったと思います。写真とは、レンズを通した光を捕捉センサー上に集束させ、ある程度現実を捕捉する行為だと言えます。モーレンは書くこととはそういうものだと主張しているのだと思います。写真撮影なら

hy は現実から物理的なシーンをキャプチャし、ライティングは実際の心から思考をキャプチャします。 LLM によって生成された、読み取ることができるものは、単に書き込みとして適格ではない方法で生成されただけです。セマンティクス。私が気になるのは文章の記事だけです。モーレン氏は、LLM は書いているわけではないと主張します。そうだと私は言います。しかし、私たちは「書く」という言葉が何を意味するかについてのみ意見が異なっており、何が制作されるかについては意見が異なっていません。
「彼はチャンスに飛びついた」と「彼はチャンスに飛びついた」など、意味的には似ているが語調が異なるフレーズの違いを「気にする」ことについては、いいえ、もちろん LLM は「気にしません」。しかし、読者である私はとても気になります。私は 11 月に ChatGPT がユーザーに選択できる「パーソナリティ」を変更 (および名前変更) することについてコラムを書きました。これらのパーソナリティは、著しく異なるスタイルとトーンのテキストを生成します。私は ChatGPT を使用しているため、クエリに対する応答の調子やスタイルを非常に気にしています。それは、私がそれらを自分の文章として誤魔化すつもりがあるからではなく、私がそれらを読んでいる一人だからです。
モーレン氏はコラムの終わり近くでこう述べています。
結局のところ、こう尋ねる以外にうまく要約することはできません。
言葉の選択をそれほど気にするのに、なぜ AI を使用するのでしょうか。
テキストを生成しますか?
これにより、AI が生成したテキストが本当に悪化するのであれば、それは良いことです。あ
多くの人はすでに LLM が生み出すものを喜んで受け入れています
「十分に良い」と考えていますが、現実的に言えば、これはそうではないと思います
何もかも変えるだろう。しかし、それがより多くの人に影響を与えるとしたら、
彼らが与えられているパブラムに不満があり、代わりに向きを変える
自分でテキストを書いたり編集したりするのであれば、それは実際には
ポジティブな結果。もしかしたら、人間の作家がさらに少なくなるかもしれない
職を失う。
同情するが、そのような可能性があるという点には同意しかねる

LLM がより質の悪い散文を生成することは、正味の利益であると見なされます。私は LLM の出力を毎日読みます。 AI に（テキストで）質問するため、AI を使用してテキストを生成します。私は、読む答えが説得力があり、明快で、正確で、ありがたいことに簡潔であることを望みます。そして理想的には、読んでいる耳に心地よく、一貫したトーンを打ち出すことです。魔神は瓶の中に戻らない。
英語は最も優れた言語であるため、おそらくより指紋採取可能である
最後に、考えるべき興味深い点があります。英語は世界で最も表現力豊かな言語です。私の言葉を真に受けないでください。私が話せるのはスペイン語だけです（高校で 4 年間スペイン語を勉強したにもかかわらず）。 20 世紀の著名な作家、スペイン語を第一言語とする多言語を話すアルゼンチン人、ホルヘ ルイス ボルヘスの言葉を取り上げましょう。 1977年にはウィリアム・F・バックリー監督の『ファイアリング・ライン』にゲスト出演した。このインタビューは YouTube で見ることができます (そしてそうすべきです) が、ジョーダン M. ポスの関連部分の転写は次のとおりです。
ボルヘス: 私はほとんど英語で読書をしました。見つけました
英語はスペイン語よりもはるかに優れた言語です。
ボルヘス: そうですね、理由はたくさんあります。まず、英語は両方とも
ゲルマン語とラテン語。これら 2 つのレジスタ — アイデアは何でも構いません
言葉は 2 つあります。これらの言葉は正確な意味ではありません
同じです。たとえば、私が「堂々とした」と言ったとしても、それは正確にはそうではありません。
「王様」と言うのと同じこと。あるいは私が「兄弟的」と言ったら、それは違います
「兄弟として」と言っているのと同じです。あるいは「暗い」「曖昧な」。それら
を

[切り捨てられた]

## Original Extract

I want the answers that I read to be cogent, lucid, accurate, blessedly terse — and ideally to strike a consistent tone that is pleasant to my reading ear. The genie is not going back in the bottle.

Manage GRC Faster with Drata’s Agentic Trust Management Platform
Follow-Up Thoughts on Watermarking Schemes for AI-Generated Text
Some follow-up to this weekend’s stemwinder “ Anthropic’s ‘Watermark’ Text Adulteration in Claude Is a Perversion of Writing ”:
Contra a bunch of idiots at Hacker News and elsewhere, I understand that popular LLMs do not just pick the “best” token (word) at each decision point. Counterintuitively, always selecting the highest-probability option produces undesirable results. So the models apply some randomization, and “temperature” is the term for the weighting that’s applied so that the “better” (higher-ranked by the model) choices have a higher chance of being chosen.
With a temperature of 1, models use their built-in probability distribution. With a temperature greater than 1, this distribution gets flatter — less-likely alternatives get a higher probability of being selected, and more-likely alternatives lower. With a temperature lower than 1, the probability distribution leans more toward the higher-ranked options. And with a temperature of 0, the highest-ranked option is always chosen. A temperature of 0 generally produces undesirable results — too predictable, too likely to get stuck. Like over-smoothing an image from a camera sensor, eliminating all noise makes the overall result worse, even if each single bit of “noise”, evaluated in isolation, is in some sense wrong.
The temperature-based randomness — which is what makes LLM output non-deterministic — is in place to help make the output better . The prose is clearly better with a temperature of 1 (with weighted randomness) than at temperature 0 (with no randomness). The watermarking schemes, on the other hand, are applying predictable-with-the-secret-key randomness for an entirely different purpose than improving the quality of the output, and thus, I believe, inherently make the output at least slightly worse.
Advocates of LLM watermarking schemes for text argue that the schemes don’t necessarily lower the quality of the generated prose, because they don’t change the temperatures — they only change the source of the randomness. Daniel Jalkut wrote a good piece today about this . I hope that’s true. I believe it’s possible that it is true. I think it’s highly unlikely that it is true. If it were true I think they’d show examples proving that it’s true. Also, Anthropic itself admits that it can’t properly watermark text that is programming language code :
For the same reason, code — which in very many cases has to be
exact — has generally less watermarking than some other forms
of text.
Having said that, in areas where there is an arbitrary choice
between particular words or terms within the code, the watermark
can be used, such as comments within code. But by definition, it
will have a negligible effect on the actual code produced.
I hold that good prose is much more like programming code. Exactness in word choice, phrasing, tone, and even punctuation is always better than imprecision. The difference is that sloppy programming code doesn’t run, or doesn’t run correctly. The human brain, on the other hand, is adept at parsing and making sense out of inexact, even sloppy, prose.
I Object Even If Quality Isn’t Adversely Affected
I do not believe these schemes can work without degrading prose quality, if only slightly. Again, though, I am open to being proven wrong. But even if we concede for the moment that such watermarking schemes do not necessarily degrade the quality of generated prose — not one iota — I still object to their use when they are being applied secretly, behind users’ backs. A useful watermark would be one that anyone can check. These SynthID “watermarks” are entirely dependent upon secrets held by the LLM providers (so far, Anthropic/Claude and Google/Gemini). I find that unacceptable, for reasons I hopefully made clear in my essay .
The people in favor of this watermarking for text have been sold a pipe dream, a fantasy. I’ve encountered dozens of comments from angry AI haters (many of them on Bluesky in particular, but also Threads and Hacker News) who are convinced that the only people who could be against the watermarking of AI-generated text are those who are duplicitously passing off AI-generated text as their own writing — and thus that I must be upset only because the jig will soon be up for me too. This of course is not true. I don’t even use AI to write text messages or emails for me, let alone a single sentence of my work.
But I find it funny that so many people who claim to believe that LLMs only produce “slop” and never anything useful also seem 100 percent convinced that the same LLMs are capable of watermarking their output in reliable ways. These people so desperately want to be able to point a finger at AI-generated text that they’ve fallen hook, line, and sinker for the argument from Google and Anthropic that, thanks to them, they’ll be able to.
I don’t want to spend too much time thinking about this because it’s a waste of time, but how exactly do these people think the existence of these mandatory watermarks and detection tools will change anything for the better? Let’s say you work at an office and you suspect that numerous of your colleagues are using AI to write emails and other work-related messages. Their messages are too long, too prolific, and lack lucidity. What are you going to do now? Copy and paste each of their messages into the watermark detectors from Anthropic, Google, and OpenAI? There cannot exist a single detector for all LLMs. And even if you find out that it says it’s a match, that an email or blog post or Slack message was very likely generated by, say, Claude, what are you going to do? March into your colleague’s office and tell them you caught them?
Anyone in a situation where “getting caught” would matter — students, say — is going to use non-watermarking LLMs or run their watermarked text through paraphrasing tools like Declaude .
No practical good is going to come of this, even if these watermarking schemes work as promised (and to be clear, I don’t believe any of it is going to work as promised). 1
My advice is not to care whether anything was written by an AI or a human. The only thing worth evaluating is what we human readers are naturally good at determining: whether it is good or bad. If it’s good, read it. If it’s not, don’t. If you’ve got a job where you’re surrounded by colleagues filling your inbox with AI-generated messages that you can’t abide, get a new job or learn to live with it. Hidden secret watermarking signals — even if they work — aren’t going to make things go back to the way they used to be. If you read something and enjoy it, and subsequently find out it was generated by an LLM, don’t feel bad. You read something good that you enjoyed.
I read something earlier today that claimed most of the posts on LinkedIn are generated by AI. That the whole platform is just inundated with AI slop. Maybe it is, but I wouldn’t know, because I never look at LinkedIn because it’s always been filled with crap. If it smells like crap it’s crap, whether the turds came out of a human anus or a turd-generating robot.
The Argument That Only People Can Truly Write
Dan Moren, writing at Six Colors today, “ LLMs Aren’t Writing ”:
LLMs do not care about the words that they pick because they
cannot care about anything.
Speaking of two things that are not the same, John rightly points
out the difference between the phrases “he leaped at the chance”
and “he jumped at the opportunity”. Those are indeed distinct — if semantically similar — phrases, each of which might be more
apt in a particular situation; or, to put it in another fashion:
the use of each of those phrases tells us something different,
whether about the person being described or the writer.
But the LLM doesn’t know which of those phrases is the right
phrase to use. It has a guess, based on its models and weights and
inputs. But the ultimate choice of those phrases tells us nothing
about the writer because there is no writer.
Moren’s is a fine retort to my post, but I fundamentally disagree — albeit at a philosophical level. If you’re reading a written work only to gain insight into the mind that produced it, there is no mind on the other end of AI-generated text. But the work itself exists. My disagreement with Moren starts and effectively ends with his (wonderfully summative) headline. I say if you can read something, it was necessarily written.
Again, this is philosophical. Was a photorealistic image generated by AI photographed ? No, I would say it was not. Photography, I would say, is the act of focusing light through a lens onto a capturing sensor, capturing, to some extent, reality. I think Moren is arguing that writing is like that. If photography captures a physical scene from reality, writing captures thoughts from an actual mind. That something you can read that was produced by an LLM was merely generated in a way that doesn’t qualify as writing . Semantics. I just care about the article of text. Moren argues that LLMs are not writing; I say they are. But we’re disagreeing only over what the word writing means, not what is being produced.
As for “caring” about the difference between semantically similar but tonally different phrases, like “ he leaped at the chance ” versus “ he jumped at the opportunity ”, no, of course the LLM doesn’t “care”. But I, the reader, care very much. I wrote a column back in November on ChatGPT changing (and renaming) the “personalities” it allows users to choose from. These personalities generate text with strikingly different styles and tones. Because I use ChatGPT, I care very much about the tone and style of its responses to my queries. Not because I’m ever going to pass them off as my own writing, but because I’m the one who is reading them.
Moren, near the end of his column:
In the end, I can’t summarize it any better than to ask: if you
care so much about word choice, why are you using AI to
generate text ?
If this does truly make AI-generated text worse, well… good . A
lot of people are already willing to accept what an LLM churns out
as “good enough” and, if I’m being realistic, I don’t think this
will change anything. But if it does lead to more people being
dissatisfied with the pablum they’re being fed and turning instead
to writing and editing their own text, then that would actually be
a positive outcome. Maybe it’d even mean fewer human writers being
put out of jobs.
I sympathize, but I must disagree that it can possibly be seen as a net good for LLMs to produce worse prose. I read the output of LLMs every day. I use AI to generate text because I ask it questions (in text). I want the answers that I read to be cogent, lucid, accurate, blessedly terse — and ideally to strike a consistent tone that is pleasant to my reading ear. The genie is not going back in the bottle.
English Is the Finest Language, and Thus, Perhaps, More Fingerprintable
Lastly, here’s an interesting point to ponder. English is the most expressive language in the world. Don’t take my word for it — it’s the only language I speak (despite four years of Spanish in high school). Take the word of famed 20th century author Jorge Luis Borges, an Argentine polyglot whose first language was Spanish. In 1977 he was the guest on William F. Buckley’s “Firing Line”. You can (and should) watch the interview on YouTube , but here’s a transcript of the relevant portion from Jordan M. Poss :
Borges: I have done most of my reading in English. I find
English a far finer language than Spanish.
Borges: Well, many reasons. Firstly, English is both a
Germanic and a Latin language. Those two registers — for any idea
you take, you have two words. Those words will not mean exactly
the same. For example if I say “regal” that is not exactly the
same thing as saying “kingly.” Or if I say “fraternal” that is not
the same as saying “brotherly.” Or “dark” and “obscure.” Those
wo

[truncated]
