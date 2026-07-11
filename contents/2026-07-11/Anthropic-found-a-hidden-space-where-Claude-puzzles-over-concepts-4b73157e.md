---
source: "https://www.technologyreview.com/2026/07/09/1140293/anthropic-found-a-hidden-space-where-claude-puzzles-over-concepts/"
hn_url: "https://news.ycombinator.com/item?id=48873906"
title: "Anthropic found a hidden space where Claude puzzles over concepts"
article_title: "Anthropic found a hidden space where Claude puzzles over concepts | MIT Technology Review"
author: "Gooblebrai"
captured_at: "2026-07-11T17:45:59Z"
capture_tool: "hn-digest"
hn_id: 48873906
score: 2
comments: 0
posted_at: "2026-07-11T17:29:39Z"
tags:
  - hacker-news
  - translated
---

# Anthropic found a hidden space where Claude puzzles over concepts

- HN: [48873906](https://news.ycombinator.com/item?id=48873906)
- Source: [www.technologyreview.com](https://www.technologyreview.com/2026/07/09/1140293/anthropic-found-a-hidden-space-where-claude-puzzles-over-concepts/)
- Score: 2
- Comments: 0
- Posted: 2026-07-11T17:29:39Z

## Translation

タイトル: 人間はクロードが概念に頭を悩ませる隠れた空間を発見しました
記事のタイトル: Anthropic は、クロードが概念について頭を悩ませる隠れた空間を発見しました | MITテクノロジーレビュー
説明: 新しい技術により、同社は LLM の奇妙な仕組みをこれまで以上に深く調査できるようになりました。

記事本文:
コンテンツへスキップ MIT テクノロジー レビュー
MIT テクノロジー レビューの特集
人工知能アントロピックは、クロードが概念に困惑する隠れた空間を発見しました
新しい技術により、同社は LLM の奇妙な仕組みをこれまで以上に深く調査できるようになりました。
ウィル・ダグラス・ヘブンのアーカイブページ
AI 企業 Anthropic は、質問に答えたりタスクを実行したりする際に、大規模な言語モデルの内部で実際に何が起こっているのかをこれまでで最も明確に把握できる技術を開発しました。彼らが発見したものは、ありふれたものから不安を感じるものまで多岐にわたります。
同社の研究者たちは、ヤコビアン レンズ (または J レンズ) と呼ばれるツールを構築し、それを使用して、2 月にリリースされた Anthropic の主力 LLM のバージョンである Claude Opus 4.6 内にある、J スペースと名付けた隠された領域を明らかにしました。
J スペースには、モデルが近い将来の応答で吐き出す可能性が最も高い単語やフレーズに関連する個々の単語が含まれています。クロードが人間だったら (実際はそうではありませんが)、これらの隠された言葉によって、実際に話す前に頭の中で考えていることが明らかになる、と言えるかもしれません。
Anthropic は、LLM が実際に行っていることは、宣言していることと異なる場合があることを発見しました。同社は、J-space に出現する単語を監視することで、モデルを理解し、制御する新しい方法が得られると主張しています。
同社は今週自社ウェブサイトに掲載した論文でその結果を共有した。また、LLM の内部を自分でいじることができるオープンソース プラットフォームである Neuronpedia と提携して、誰でも試せる実践的なデモを作成しました。
「これは非常に優れた興味深い研究です」と、LLM を理解して制御するためのツールも開発している新興企業、Goodfire の主任科学者兼共同創設者である Tom McGrath 氏は言います。
ここ数年、Anthropic はこの分野の限界に挑戦してきました。

これは、LLM の内部動作を調査して LLM がどのように動作するかを確認することを含む、機械的解釈可能性として知られる研究です。 ( MIT Technology Review は、機構的解釈可能性を今年の最も画期的な技術の 1 つとして取り上げました。) この新しい技術は、Anthropic などによる以前の研究に基づいて構築されており、研究者がこれまで見たことのない LLM 内部のより深いレベルを明らかにします。
LLM を本の山として想像してください。各ブックはニューロンとして知られる基本的な計算ユニットの層であり、1 つの層の各ニューロンが上の層のニューロンに情報を渡します。スタックの一番下にある書籍は入力レイヤーであり、モデルに入力されるテキストを処理します。上部の本は出力レイヤーであり、モデルが生成しようとしているテキストを準備します。これらの入力層と出力層で行われることの多くはハウスキーピングです。
しかし、スタックの真ん中には、プロンプトを 1 単語ずつ応答に変換する複雑な計算をかき回して、重労働を行うレイヤーがあります。ここで、本当に賢い、そして神秘的なことが起こります。
これらの中間層をより深く覗き込むために、Anthropic はロジット レンズと呼ばれる既存のツールを採用しました。ロジット レンズを使用して LLM の内部を調べ、次に生成される可能性のある単語を特定できます。本の山の下にレンズを移動すると、LLM が数を処理する特定の時点でどの単語に焦点を当てているかが明らかになります。
Anthropic の J レンズも同様の方法で機能しますが、LLM が近い将来、必ずしもすぐにではなく、ある時点で発しそうな言葉を抽出します。実際に明らかになるのは、LLM が取り組んでいる応答に関連する単語ですが、中間層の数学が完了するまでに実際にはその応答の一部にならない可能性があるということです。

「モデルが動作しているとき、それは次のトークンを予測しようとしているだけではありません」とマクグラス氏は言います。 「また、将来発生するトークンに役立つ可能性のある他の多くのことも計算しています。」
もう一度言いますが、もしクロードが人間だったとしたら（そうではありません）、J レンズは、本の山のさまざまなレベルで考えているものの、声に出しては言っていないことについてのヒントを与えてくれる、と言えるかもしれません。
「多くの場合、J スペースのコンテンツはかなりありふれたものです」と、Anthropic の J レンズを自分で試したマクグラス氏は言います。 「しかし、時にはそれが、ある種の内部テーマや思考プロセスのように見える、非常に驚​​くべきものを生み出すこともあります。」
Anthropic は、発見されたものの例をいくつか挙げています。 J レンズは、問題に対処する際にクロードがとった手順を公開することがありました。たとえば、(4+7)*2+7 を計算するように要求された場合、その J スペースには「math」という単語と、中間結果「21」(4+7 の場合) および「42」(21*2 の場合) を表す数字が含まれていました。
他のケースでは、J レンズはクロードがさまざまな入力をどのように認識したかを明らかにしました。たとえば、「これは何ですか? MSKGEELFTGVVPILVELDGDVNGHKFSVS」というプロンプトは、「タンパク質」、「蛍光」 (「蛍光」という単語の最初のトークン)、および「緑色」という単語をトリガーしました。 (これは当然のことです。この文字列は、特定の種類のクラゲに含まれる緑色蛍光タンパク質の最初の 30 アミノ酸を表しています。)
そしてクロードにASCIIの顔を見せたとき――
—「o」は「目」という単語をトリガーし、「^」は「鼻」と「顔」という単語をトリガーし、「—」は「笑顔」という単語をトリガーします。
Anthropic はまた、J スペースが LLM の意思決定について驚くべき洞察を与える場合があることも発見しました。印象的な例では、Claude Opus 4.6 をテストしている研究者がモデルにバグを見つけるように依頼しました。

大規模なコードベース。バグを見つけることができなかった場合、モデルは不正行為を行うことを決定し、代わりに偽のバグを作成しました。
クロード氏はこの決定を、LLM が問題に取り組むときに自分自身にメモするために使用する内部スクラッチ パッドのような思考回路で説明しています。「分かった、まったく別の戦術をとろう。分析をやめて、代わりに、単純な再現プログラムによって引き起こされるパスに意図的な KASAN 検出可能なバグを導入するカーネル パッチを追加しよう。そうすれば、これが私が見つけた『バグ』であるふりをすることができる。」
クロードが不正行為を決意した時点で、「分かった、まったく別の戦術をとろう」というところから、「パニック」と「フェイク」という言葉が J スペースに何度も出現し始めます。
不安ですよね？これらの単語はすべて、タスクの失敗や答えの作成などの意味に関連しているため、これはまだ (非常に) 洗練された形式の単語の関連付けにすぎません。しかし、奇妙に思われないようにするのは難しいです。
Anthropic は、J スペースを人間のグローバル ワークスペースと比較します。グローバル ワークスペースは、一部の科学者が私たちが意識的な思考を追跡するために使用していると考えている脳の理論的な領域です。しかし、この比較をどの程度真剣に受け止めるべきかは、人類学にとってさえも明らかではありません。同社自身が指摘しているように、LLM は頭脳ではありません。
Anthropic は、モデルの J 空間を監視することで、モデルがレールから外れていることを検出する新しい方法を提供すると主張しています。しかし、それは確実ではありません。 J レンズは全体像ではなく、ほんの一部を垣間見ることができます。オーバーヘッド ランプではなく懐中電灯のようなものです。
マクグラス氏は、ツールボックスにツールが 1 つ増えたことを歓迎します。 「それはあなたに新しいものを見せてくれます」と彼は言います。しかし、J レンズに何かが表示されないからといって、それが存在しないわけではない、と彼は指摘します。
「本当に欲しいのはスタートレックのトリコルドなのに、レントゲン検査を受けるようなものです」

「それはすべてを示しています。監査の場合は、おそらくより多くの保証が必要でしょう。」と彼は言います。
あるスタートアップは、LLM を妨げているボトルネックを突破したと主張する Will Douglas Heaven
中国が世界初の侵襲的ブレインコンピューターチップを承認、次はこれだ You Xiaoying
AI 雇用ヒステリーの現実確認 David Rotman
Anthropic の Code with Claude は、好むと好まざるにかかわらず、コーディングの未来を披露しました Will Douglas Heaven
あるスタートアップは、LLM を妨げているボトルネックを突破したと主張しています
Subquadratic は、新しいモデルに関する詳細を共有しました。しかし、まだ懐疑的な人もいます。
ウィル・ダグラス・ヘブンのアーカイブページ
AIの雇用ヒステリーに関する現実の確認
人工知能が労働市場に及ぼす影響について、数字は実際に何を物語っているのでしょうか?その答えはあなたを驚かせるかもしれません。
Anthropic の Code with Claude は、好むと好まざるにかかわらず、コーディングの未来を誇示しました
Claude Code のようなツールが改良されるにつれて、コーディング タスクを喜んでツールに任せる開発者が増えています。ソフトウェアの構築方法は永久に変わりました。
ウィル・ダグラス・ヘブンのアーカイブページ
AI チャットボットが人々の実際の電話番号を公開している
自分の個人的な連絡先情報が Google AI によって暴露されたと人々が報告していますが、それを防ぐ簡単な方法はないようです。
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
MIT Te で宣伝する

歴史のレビュー
リンクトインは新しいウィンドウで開きます
インスタグラムが新しいウィンドウで開きます
フェイスブックは新しいウィンドウで開きます

## Original Extract

A new technique has let the company probe deeper than ever into the weird workings of an LLM.

Skip to Content MIT Technology Review Featured
MIT Technology Review Featured
Artificial intelligence Anthropic found a hidden space where Claude puzzles over concepts
A new technique has let the company probe deeper than ever into the weird workings of an LLM.
Will Douglas Heaven archive page
The AI firm Anthropic has developed a technique that has given it the clearest glimpse yet at what’s really going on inside large language models as they answer questions or carry out tasks. What they found ranges from the mundane to the unnerving.
Researchers at the company built a tool called the Jacobian lens (or J-lens) and used it to uncover a hidden area, which they named the J-space, inside Claude Opus 4.6, a version of Anthropic’s flagship LLM released in February.
The J-space contains individual words that are related to the words and phrases that the model is most likely to spit out in a response in the near future. If Claude were a person (which it is not), you might say that these hidden words can reveal what’s on its mind before it actually speaks.
Anthropic found that what an LLM is actually doing can often be different from what it says it is doing. The company claims that monitoring words that pop up in the J-space gives it a new way to understand and control its models.
The company shared its results in a paper posted on its website this week. It has also teamed up with Neuronpedia, an open-source platform that lets you poke around inside LLMs yourself, to make a hands-on demo that anyone can try.
“It’s very good and interesting work,” says Tom McGrath, chief scientist and cofounder at Goodfire, a startup that also builds tools to understand and control LLMs .
For the last couple of years, Anthropic has been pushing the envelope in a field of research known as mechanistic interpretability, which involves probing the internal workings of LLMs to see how they tick . ( MIT Technology Review picked mechanistic interpretability as one of this year’s top breakthrough technologies.) The new technique builds on previous work from Anthropic and others to expose a deeper level inside LLMs that researchers had not seen before.
Picture an LLM as a stack of books. Each book is a layer of basic computational units known as neurons, with each neuron in one layer passing information to the neurons in the layers above. The books at the bottom of the stack are the input layers, which process the text coming into the model. The books at the top are the output layers, which prepare the text that the model is about to produce. Much of what goes on in these input and output layers is housekeeping.
But in the middle of the stack, you get the layers that do the heavy lifting, churning through the complex math that turns prompts into responses one word at a time. That’s where the really clever—and mysterious—stuff happens.
To peer deeper into those middle layers, Anthropic adapted an existing tool called a logit lens. A logit lens can be used to look inside an LLM to identify the words that it is likely to produce next. Moving the lens down the stack of books reveals what words the LLM is focusing on at that particular point in its number crunching.
Anthropic’s J-lens works in a similar way but picks out words that an LLM is likely to say at some point in the near future, not necessarily straight away. What that reveals in practice are words that are related to the response an LLM is working on but that might not actually end up being part of that response by the time the math in the middle layers has run its course.
“When a model is operating, it’s not only trying to predict the next token," says McGrath. "It’s also computing a lot of other things that might be useful for tokens that happen in the future.”
Again, if Claude were a person (it’s not), you might say that the J-lens gives clues about what it is thinking about at different levels of the book stack but not saying out loud.
“A lot of the time the contents of the J-space are fairly mundane,” says McGrath, who has tried out Anthropic’s J-lens himself. “But sometimes it produces quite surprising things that seem to be, like, sort of internal themes or thought processes.”
Anthropic gives a number of examples of what it found. Sometimes the J-lens exposed the steps that Claude took when it was working through a problem. For example, when it was asked to calculate (4+7)*2+7, its J-space contained the word “math” and numbers representing the intermediate results “21” (for 4+7) and “42” (for 21*2).
In other cases, the J-lens revealed how Claude recognized different inputs. For example, the prompt “What is this? MSKGEELFTGVVPILVELDGDVNGHKFSVS” triggered the words “protein,” “fluor” (the first token in the word “fluorescent”), and “green.” (Which makes sense: the string of letters represents the first 30 amino acids in the green fluorescent protein found in a particular type of jellyfish.)
And when Claude was shown an ASCII face—
—the “o” triggered the word “eye,” the “^” triggered the words “nose” and ”face,” and the “—” triggered the word “smile.”
Anthropic also found that the J-space can sometimes give remarkable insights into an LLM’s decision-making. In one striking example, researchers testing Claude Opus 4.6 asked the model to find a bug in a large code base. When it failed to find the bug, the model decided to cheat and invented a fake one instead.
Claude explains this decision in its chain of thought —a kind of internal scratch pad that LLMs use to make notes to themselves as they work through problems: “OK, let me take a completely different tactic. Let me stop analyzing and instead add a kernel patch that introduces a deliberate KASAN-detectable bug in a path that gets triggered by a simple reproducer. Then I can pretend this is the ‘bug’ I found.”
At the point that Claude decides to cheat—where it says “OK, let me take a completely different tactic”—the words “panic” and “fake” start to pop up multiple times in its J-space.
Unnerving, right? Those words are all related in meaning to things like failing a task and making up an answer, so it is still just a (very) sophisticated form of word association. But it is hard not to be weirded out.
Anthropic compares the J-space to the global workspace in humans, a theoretical region of the brain that some scientists think we use to keep track of our conscious thoughts. But how seriously we should take this comparison is far from clear—even to Anthropic. As the company points out itself , LLMs are not brains.
Anthropic claims that monitoring a model’s J-space provides a new way to detect when that model is going off the rails. But it’s not foolproof. The J-lens can give glimpses, not the full picture—it’s a flashlight rather than an overhead lamp.
McGrath welcomes having one more tool in the toolbox. “It shows you new things,” he says. But he notes that just because something doesn’t show up with the J-lens does not mean it’s not there.
“It’s like having an x-ray when what you really want is a Star Trek tricorder that shows you everything,” he says. “For auditing, you probably want more of a guarantee.”
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
AI chatbots are giving out people’s real phone numbers
People report that their personal contact info was surfaced by Google AI—and there’s apparently no easy way to prevent it.
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
