---
source: "https://www.technologyreview.com/2026/07/13/1140343/what-anthropics-latest-ai-discovery-does-and-doesnt-show/"
hn_url: "https://news.ycombinator.com/item?id=48903165"
title: "What Anthropic's latest AI discovery does–and doesn't–show"
article_title: "What Anthropic’s latest AI discovery does—and doesn’t—show | MIT Technology Review"
author: "joozio"
captured_at: "2026-07-14T07:03:58Z"
capture_tool: "hn-digest"
hn_id: 48903165
score: 1
comments: 0
posted_at: "2026-07-14T07:01:34Z"
tags:
  - hacker-news
  - translated
---

# What Anthropic's latest AI discovery does–and doesn't–show

- HN: [48903165](https://news.ycombinator.com/item?id=48903165)
- Source: [www.technologyreview.com](https://www.technologyreview.com/2026/07/13/1140343/what-anthropics-latest-ai-discovery-does-and-doesnt-show/)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T07:01:34Z

## Translation

タイトル: Anthropic の最新の AI 発見が示すものと示さないもの
記事のタイトル: Anthropic の最新の AI 発見が示すものと示さないもの | MITテクノロジーレビュー
説明: 同社は、モデルがどのように答えに到達するかについて新たな扉を発見したと述べています。上級編集者のウィル・ダグラス・ヘブンにそれについて話を聞いた。

記事本文:
コンテンツへスキップ MIT テクノロジー レビュー
MIT テクノロジー レビューの特集
人工知能 Anthropic の最新の AI 発見で何が分かるのか、何が分からないのか
同社は、モデルがどのように答えに到達するかについて新たな扉を発見したと述べている。上級編集者のウィル・ダグラス・ヘブンにそれについて話を聞いた。
このストーリーは元々、AI に関する週刊ニュースレター『The Algorithm』に掲載されました。このような記事を受信トレイで受け取るには、まずここでサインアップしてください。
Anthropic は現在世界で最も価値のある AI 企業であり、評価額は 1 兆ドル近くに達していますが、奇妙で目まぐるしい研究を発表することで評判です。たとえば、AI モデルが痛みを感じることができるかどうかを調査しており、ユーザーがモデルを「悪用」していると疑われる場合にはチャットボットの会話を遮断することもあります。
Anthropic が他の AI 企業よりも多くの時間と資金を費やしているニッチ分野の 1 つは、機械的解釈可能性と呼ばれます。これは、AI モデルの複雑な数学的内部を調べて、特定の出力が得られ、別の出力が得られない理由を知ることを意味します。それは複雑なことです。結果に寄与する可能性のあるデータ ポイントが何百万もあり、それらを調べていくのは役に立つというよりも、言葉のサラダのように見える場合があります。これも物議を醸しています。心理学や神経科学から借用した用語を使用して AI モデルを説明すると、AI の動作が私たちが判断するよりも洗練されているように見えることがあります。
だからこそ、Anthropic が先週、モデルが答えを推論する際の「内なる思考」に新たな窓を発見したと発表したとき、私が話をしなければならない同僚がいたのです。上級編集者の Will Douglas Heaven は、コンピューター サイエンスの博士号を取得していることに加えて、AI モデルがどのように機能するかについて、私たちが言えることを深く掘り下げることに多くの時間を費やしてきました。私は彼と何を学ぶべきかについて話しました

Anthropic の新しい (そして予想どおり風変わりな) 研究。
Anthropic はここで正確に何を学んだのでしょうか?
Anthropic はここ数年、大規模言語モデル (LLM) がどのように機能するかを理解しようと努めてきました。これに注目しているのは Anthropic 社だけではありませんが、同社は他の企業よりもこれを中核的使命の一部として取り組んでいると思います。 Anthropic の CEO、Dario Amodei 氏は、LLM がどのように機能するかをもっと学ばない限り、LLM を完全に制御することはできないだろうと述べています。
したがって、この新しい研究はまさにその文脈に沿ったものです。 LLM 内部の奇妙なメカニズムをこれまで以上に深く掘り下げています。 Anthropic が学んだことは、LLM の内部には、出力には現れないが問題を解決する方法に影響を与えると思われる単語で満たされた空間 (Anthropic では J スペースと呼んでいます) があるということでした。これらすべては、Anthropic がモデルのクロードを調査するための新しい技術を開発するまで隠されていたため、これは本物の発見です。
これらの単語は、LLM が特定のタスクでどこまで到達したかを追跡する場合もあれば、認識の閃きのように見える場合もあり (たとえば、LLM にタンパク質配列の文字だけを与えると「タンパク質」がポップアップする場合があります)、場合によってはモデルの意思決定に関する一種の内部コメントを表す場合もあります。私のお気に入りの例では、「パニック」という単語が現れたときに、クロードはコーディング テストでカンニングをすることにしました。
Anthropic はまた、LLM がこの空間内の単語を記述し、操作できることも発見しました。それで、どういうわけか彼らはそれを利用しているようです。
少し後戻りしましょう。私は大規模な言語モデルをそれほど単純だとは思いませんが、魔法でもありません。単語間の関係を学習する数学がたくさんありますよね?では、何が起こっているかを知るために LLM を「覗く」のはなぜそれほど難しいのでしょうか?
そう、それらは魔法ではないのです！私たちが完全に解放していないという事実は、

それらが神話作りに関与していることを理解してください。そして、ここで Anthropic が傾いている物語全体、つまり、彼らはこの本当に謎に満ちたテクノロジーを構築しましたが、それを理解するのも彼らなので心配しないでくださいという物語全体が、同社の雰囲気と非常に一致していることは注目に値します。 [Anthropic が、自社の新しいモデルはコーディング能力が非常に優れているため、世界的なサイバーセキュリティのリスクを引き起こすと警告したが、その直後に米国政府によって閉鎖されたことをご覧ください。]
はい、LLM は単なる数学です。それでも、それは非常に複雑な数学です。今日の LLM は数千億の数値で構成されているだけでなく、LLM を実行すると、何百万もの計算がカスケードで行われます。私は昨年、中程度の LLM を紙に印刷すれば、サンフランシスコほどの規模の都市をカバーできると書きました。
特定の時点で LLM の特定の部分を強調表示する専門ツールがなければ、こうした数学を理解することは不可能です。どこを見るべきか、どのように見るべきかを知る必要があります。そして、これらのツールを構築するには、そもそもその複雑な数学を理解する必要があります。
あなたは、生物の脳を研究するような方法で LLM を研究するこの概念について、他の場所で書きました。 LLM の仕組みについて話すときに「脳のような」用語を使用するのは公平でしょうか?
私はそのような言葉を使うのは好きではありません。 LLM は頭脳ではありません。このような話は誤解を招きます。LLM が実際よりも人間に近いことができるかのように示唆したり、LLM がどのように動作するかについて、私たちがすべきではないような仮定を立てたりする可能性があるからです。擬人化というもの全体は、このテクノロジーとは何か、そしてそれが将来どうなるのかについての多くの強力なイデオロギー的立場とも結びついています。
しかし同時に、これらが何を意味するのかについて話すための適切な代替語彙が私たちには不足しています。

モデルたちがやってます。人々が「考える」「理解する」「脳のような」などの言葉に手を伸ばす理由は理解できます。これらは便利な略語です。
Anthropic は、LLM 内で発見されたこの新しい空間を、一部の神経科学者が私たちの脳が意識的な思考を追跡するために使用していると考えている空間と比較しています。私はその比較をどの程度真剣に受け止めるべきなのか同社に尋ねたところ、同社は声明で次のように述べた。「これらの類推を描くことは、J 空間について多くの非自明な実験予測を行うことができ、それが真実であることが判明したため、実験を計画する際に役に立ちました。同時に、J 空間 (および一般的な言語モデル) と人間の脳の間にはいくつかの重要な違いがあることに注意することが重要です。そのため、完全な一致があると主張するつもりはありません。」
J スペースのこの新しい概念を使用して解決できる可能性がある AI の問題は何ですか?
Anthropic 氏は、J スペースを監視することで、モデルがやってはいけないことをしているのを発見できる可能性があると述べています。モデルの出力には現れない単語がこのスペースに現れるため、モデルが偏った応答を返しているときや不正行為のメリットとデメリットを比較検討しているときなど、他の方法では気づかなかったかもしれないその動作についての情報を得ることができます。
少なくともそれが理論です。この結果は、それ自体が役に立つというよりも、この技術全体を理解するためのもう 1 つのステップとして考えたほうがよいと思います。
新しい研究に関するウィルの全文をご覧ください。
あるスタートアップは、LLM を妨げているボトルネックを突破したと主張する Will Douglas Heaven
中国が世界初の侵襲的ブレインコンピューターチップを承認、次はこれだ You Xiaoying
AI 雇用ヒステリーの現実確認 David Rotman
アントロピックのC

クロードとのオードは、コーディングの未来を誇示しました - 好むと好まざるにかかわらず、ウィル・ダグラス・ヘブン
あるスタートアップは、LLM を妨げているボトルネックを突破したと主張しています
Subquadratic は、新しいモデルに関する詳細を共有しました。しかし、まだ懐疑的な人もいます。
ウィル・ダグラス・ヘブンのアーカイブページ
AIの雇用ヒステリーに関する現実の確認
人工知能が労働市場に及ぼす影響について、数字は実際に何を物語っているのでしょうか?その答えはあなたを驚かせるかもしれません。
Anthropic の Code with Claude は、好むと好まざるにかかわらず、コーディングの未来を誇示しました
Claude Code のようなツールが改良されるにつれて、コーディング タスクを喜んでツールに任せる開発者が増えています。ソフトウェアの構築方法は永久に変わりました。
ウィル・ダグラス・ヘブンのアーカイブページ
Google I/O は、AI 主導の科学の道がどのように変化しているかを示しました
2 年前、ある AI ツールが Google DeepMind にノーベル賞を受賞しました。研究者たちは現在、新たな目標に向かって進んでいます。
最新のアップデートを次から入手してください
MITテクノロジーレビュー
特別オファー、トップニュース、
今後のイベントなど。
プライバシー ポリシー メールを送信していただきありがとうございます。
何か問題があったようです。
設定を保存できません。
このページを更新して更新してみてください
さらに時間がかかります。このメッセージが引き続き表示される場合は、
までご連絡ください
customer-service@technologyreview.com に受信を希望するニュースレターのリストを添えて送信してください。
レガシーの最新版
MIT テクノロジーレビューによる広告掲載
リンクトインは新しいウィンドウで開きます
インスタグラムが新しいウィンドウで開きます
フェイスブックは新しいウィンドウで開きます

## Original Extract

The company says it has found a new window into how its models arrive at answers. We spoke with senior editor Will Douglas Heaven about it.

Skip to Content MIT Technology Review Featured
MIT Technology Review Featured
Artificial intelligence What Anthropic’s latest AI discovery does—and doesn’t—show
The company says it has found a new window into how its models arrive at answers. We spoke with senior editor Will Douglas Heaven about it.
This story originally appeared in The Algorithm, our weekly newsletter on AI. To get stories like this in your inbox first, sign up here .
Anthropic—currently the world’s most valuable AI company, with a nearly $1 trillion valuation—has a reputation for publishing strange and heady research. It’s looking into whether AI models can feel pain , for example, and will sometimes cut off chatbot conversations if it suspects users are “abusing” the model.
One niche that Anthropic spends more time and money on than other AI companies is called mechanistic interpretability, which means looking inside the complex math of an AI model to learn why it comes up with one particular output and not another. It’s complicated stuff; there are millions of data points that might contribute to any result, and wading through them can look more like word salad than anything useful. It’s also controversial. Describing AI models with terms borrowed from psychology and neuroscience can make their behavior seem more sophisticated than we might otherwise judge it to be.
That’s why, when Anthropic announced last week that it had found a new window into its models’ “internal thoughts” as they reason through answers, there was one colleague I had to talk to. Senior editor Will Douglas Heaven, aside from having a PhD in computer science, has spent a lot of time digging into what we can say about how AI models work. I spoke with him about what we should take from Anthropic’s new (and predictably quirky) research.
What did Anthropic learn here, exactly?
Anthropic has been trying to understand how large language models (LLMs) work for a few years now. Anthropic isn’t the only one looking at this, but I think the company has made it part of its core mission more than most. Anthropic’s CEO, Dario Amodei, has said we won’t be able to control LLMs fully unless we learn more about how they work.
So this new research is very much in that context. It goes deeper into the weird mechanisms inside LLMs than ever before. What Anthropic learned was that LLMs have a space inside them—which Anthropic calls the J-space—filled with words that don’t appear in their output but that seem to influence the way they puzzle through problems. All this was hidden until Anthropic developed a new technique to probe its model Claude, so it’s a genuine discovery.
Sometimes these words keep track of where the LLM has got to in a particular task, sometimes they look more like flashes of recognition (for example, “protein” might pop up when you give an LLM only the letters of a protein sequence), and sometimes they represent a kind of internal commentary on the model’s decision-making. In my favorite example, Claude decided to cheat on a coding test when the word “panic” appeared.
Anthropic also found that LLMs are able to describe and manipulate the words in this space. So somehow they seem to be making use of it.
Let’s step back for a second. I don’t think of large language models as simple , but they’re also not magic. There’s a bunch of math that learns relationships between words, right? So why is it so hard to “peer” into an LLM to know what’s going on?
Yeah, they’re not magic! I think the fact we don’t fully understand them plays into the mythmaking. And it’s worth noting that the whole narrative that Anthropic is leaning into here—that they’ve built this really mysterious technology, but don’t worry, because they’re also the ones to figure it out—very much fits with the company’s vibe. [See how Anthropic warned that its new models were so good at coding they posed a global cybersecurity risk, only for the US government to shut them down shortly thereafter.]
So yes: LLMs are just math. And yet it’s vastly complex math. Not only are today’s LLMs made out of hundreds of billions of numbers, but running them triggers a cascade of millions and millions of calculations. I wrote last year that if you printed out even a medium-size LLM on pieces of paper, it would cover a city the size of San Francisco .
It’s impossible to make sense of any of that math without specialist tools that highlight specific parts of an LLM at specific times. You need to know where to look and how to look. And building those tools requires understanding something of that complex math in the first place.
You’ve written elsewhere about this concept of studying LLMs the way one might study an organism’s brain. Is it fair to use “brain-like” terms when talking about how an LLM works?
I don’t love using those kinds of terms. LLMs are not brains. Talking like this is misleading because it can suggest that LLMs are capable of more human-like things than they are or that we can make assumptions about how they might behave that we shouldn’t. The whole anthropomorphization thing is also tied up with a bunch of strong ideological positions about what this technology is and what it’s going to be .
But at the same time, we lack a good alternative vocabulary for talking about what these models are doing. I can understand why people reach for words like “think” and “understand” and “brain-like”—they’re convenient shorthand.
Anthropic compares this new space it found inside LLMs to the space that some neuroscientists think our brains use to keep track of conscious thoughts. I asked the company how seriously we should take that comparison and it said in a statement: “Drawing these analogies was helpful to us in designing our experiments, as they allowed us to make many non-obvious experimental predictions about the J-space that turned out to be true. At the same time, it’s important to note that there are some important differences between the J-space (and language models in general) and the human brain, so we don’t mean to claim there’s a perfect correspondence.”
What’s a problem in AI that this new concept of the J-space might be used to solve?
Anthropic has said that monitoring the J-space could be a way to catch models doing something they shouldn’t. Because words pop up in this space that don’t appear in a model’s output, they can tell you things about its behavior that you might not have noticed otherwise—such as when it is giving biased responses or when it is weighing the pros and cons of cheating.
That’s the theory, at least. I think it’s better to think of this result as one more step on the path to understanding this technology overall than as something that will be useful by itself.
Read more in Will’s full story about the new research .
A startup claims it broke through a bottleneck that’s holding back LLMs Will Douglas Heaven
China has approved the world’s first invasive brain-computer chip—here’s what’s next You Xiaoying
A reality check on the AI jobs hysteria David Rotman
Anthropic’s Code with Claude showed off coding’s future—whether you like it or not Will Douglas Heaven
A startup claims it broke through a bottleneck that’s holding back LLMs
Subquadratic has now shared more details about its new model. But some are still skeptical.
Will Douglas Heaven archive page
A reality check on the AI jobs hysteria
What do the numbers really say about the impact of artificial intelligence on the labor market? The answer might surprise you.
Anthropic’s Code with Claude showed off coding’s future—whether you like it or not
As tools like Claude Code get better, more and more developers are happy to hand off coding tasks to them. The way software gets built has changed for good.
Will Douglas Heaven archive page
Google I/O showed how the path for AI-driven science is shifting
Two years ago, a n AI tool won Google DeepMind a Nobel. Researchers are now climbing toward a new goal.
Get the latest updates from
MIT Technology Review
Discover special offers, top stories,
upcoming events, and more.
Privacy Policy Thank you for submitting your email!
It looks like something went wrong.
We’re having trouble saving your preferences.
Try refreshing this page and updating them one
more time. If you continue to get this message,
reach out to us at
customer-service@technologyreview.com with a list of newsletters you’d like to receive.
The latest iteration of a legacy
Advertise with MIT Technology Review
linkedin opens in a new window
instagram opens in a new window
facebook opens in a new window
