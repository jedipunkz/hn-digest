---
source: "https://cholling.com/posts/quallms/"
hn_url: "https://news.ycombinator.com/item?id=49249037"
title: "I'm not anti-AI, but I have QUALLMS"
article_title: "I'm not anti-AI, but I have QUALLMS - Chip Hollingsworth"
author: "CagedCoder"
captured_at: "2026-08-10T20:31:30Z"
capture_tool: "hn-digest"
hn_id: 49249037
score: 3
comments: 0
posted_at: "2026-08-10T20:11:14Z"
tags:
  - hacker-news
  - translated
---

# I'm not anti-AI, but I have QUALLMS

- HN: [49249037](https://news.ycombinator.com/item?id=49249037)
- Source: [cholling.com](https://cholling.com/posts/quallms/)
- Score: 3
- Comments: 0
- Posted: 2026-08-10T20:11:14Z

## Translation

タイトル: 私は反AIではありませんが、QUALLMSを持っています
記事のタイトル: 私は反 AI ではありませんが、QUALLMS を持っています - チップ・ホリングスワース
説明: チップ・ホリングスワースのポートフォリオ

記事本文:
ホーム
ブログ
私は反AIではありませんが、QUALLMSを持っています
私は、現在の大規模言語モデルの誇大宣伝に対して、恥ずかしがらずに否定的な意見を表明してきました。しかし、私は自分自身を「反AI」と表現するのにはためらいがあります。それは、私が「場合によってはLLMでも大丈夫」とか「これらの問題を解決できればLLMでも大丈夫」といった、ある種の超微妙な立場を持っているからではありません。実際、私はこの誇大宣伝の根拠となっている仮定そのものに断固として反対しています。つまり、単語の共起頻度の統計モデルは、質問に答えたり、意味のあるテキストやコードを生成したりするための多目的ツールの優れた基盤であるというものです。作成者の明示的な許可を得て取得したテキストでローカル モデルをトレーニングし、再生可能エネルギーのみを燃料とするマシンで実行することですが、私はそれが良いアイデアだとは今でも思っていません。商用 LLM が同意と著作権を無視し、法外な量のリソースを消費するという事実は、根本的に悪いことをさらに悪化させるだけです。
では、なぜ私は自分自身を「反AI」と呼ばないのでしょうか？それは 2 つのあいまいさのためです。1 つは「AI」の定義、もう 1 つは「反 AI」の意味です。だからこそ、私は私の立場を正確に説明するために別の命名法を提案しているのです。
「AI」がほとんど意味のない用語である理由
「人工知能」という用語は、通常、1956 年の人工知能に関するダートマス夏季研究プロジェクトの頃に誕生したと考えられています。当時でさえ、この用語が何を指すのかは正確には明確ではありませんでした。それは、情報理論、コンピューターサイエンス、その他の分野における多くの半関連研究を表す一種の包括的な用語として選ばれました。大まかに言えば、人間の思考の効果をコンピューターで模倣する試みを指すようです。私は言う、「ヒューマの影響を模倣する」

なぜなら、コンピュータが知識表現、言語処理、視覚認識、その他の「AI」の取り組みを処理するプロセスが、人間の脳内で起こっていることと何ら変わらないと関係者は幻想を抱いていなかったからだ。人間の脳が、私たちが聞いたすべての文を解析するためにLispインタプリタを実行しているとは、まったく誰も考えていなかった。だからこそ、彼らは「人工」という言葉を使ったのだ：これは、物そのものではなく、物を模倣したものであることを強調するために。（信念）脳は実際にデジタル コンピューター内で行われているのと同様の何らかの計算を行っており、実際に人間の脳を模倣するコンピューター システムを設計できるということは、認知科学という関連はあるものの別の分野に属します。)
当初、「人工知能」と呼ばれるもののほとんどはルールベースのシステムであり、他の種類のコンピューター プログラムと本質的には変わりませんでした。主な違いは、コンピューティングは伝統的に、発射体の軌道の計算や代数方程式の解法など、アルゴリズム形式で表現するのが比較的簡単なタスクに使用されてきたのに対し、人工知能は、自然言語テキストの情報内容の表現、それについての推論、自然言語応答の作成など、アルゴリズムのあまり明らかではない領域を征服しようとしていたという点です。しかし時間の経過とともに、自動化された統計的または準統計的手法を使用して、一連の入力データから一般化する「機械学習」が主流になりました。これには、食べ物の写真を「ホットドッグ」または「ホットドッグではない」という大まかなカテゴリに分類したり、Web サイトの顧客が最近購入したいくつかの商品から次に何を買いたいかを予測したりすることが含まれる場合があります。これは非常に異なるアプローチを表しました

「古き良き人工知能」 (GOFAI) として知られるようになったものよりも、タスクを実行するために必要な手順を考え出してコンピューターにそれを実行するように指示するのではなく、大量のデータをプログラムに投げてコンピューターに何をすべきかを判断させるだけです。少なくとも、それは擬人化的に説明されることが多いです。実際、これはより高い抽象レベルでの同じルールベースのアプローチです。入力を出力に変換するためのルールではなく、トレーニング データを入力から出力に変換するためのルールに変換する高次のルールです。しかし、コンピューターが物事のやり方を「学習する」と言うのははるかに簡単で、残念なことに、機械学習システムが実際よりも人間の脳に近いことをしていると多くの人が思い込んでしまいます。
そのため現在では、「人工知能」は、シンボルを操作するための昔ながらのアルゴリズムと、新しく巧妙な統計的トリックの両方を指すために使用されています。 2 つのまったく異なるテクノロジーが同じ用語で説明されていますが、唯一の共通点は、人々が行うことをエミュレートしたいという漠然とした願望です。時間が経つにつれ、コンピューティング ハードウェアが安価になり、より強力になるにつれて、「ディープ ラーニング」が注目されるようになりました。これは実際には単なる多層パーセプトロンであり、1950 年代に遡る機械学習手法のもう少し高度なバージョンです。今になって初めて、それらをさらに大きくし、より大きなデータセットでトレーニングできるようになりました。ほぼ同時期に、インターネットが本格的に普及し、恐ろしい量のデータ、特に自然言語データを収集することがこれまで以上に簡単になりました。ディープラーニングは、初期の形式の機械学習と根本的に異なる種類のものではありませんでしたが、関与する規模により、以前は私が理解していなかった多くのこと (画像認識、自動翻訳) を実行できるようになりました。

thods はまあまあでした。そのため、多くの人々がディープラーニングをそれ自体として捉え始め、拡大を続ける「AI」の傘下にさらに 1 つのカテゴリーを作りました。
ごく最近、人々は、ある種の深層学習アーキテクチャと膨大な量のテキストを使用すれば、数百の単語に先行する次の単語を予測するようにトレーニングできることに気づきました。そして彼らは、人間が会話の半分を提供し、単語予測器が残りの半分を提供する、対話をシミュレートするフレームワークをこれに巻き付けました。そして、人間味のある会話を生み出すのが非常に上手で、ベンチャーキャピタルの資金が非常に潤沢だったので、企業はそれを市場のあらゆるソフトウェア製品に注ぎ込み始めました。しかし、彼らはそれを「大規模言語モデル」、「ネクストトークン予測子」、さらには「チャットボット」とは呼びませんでした。彼らはそれを「AI」と呼んだだけです。
その結果、現在、誰かが「AI」と言うとき、その人はさまざまな時点ですべて「AI」と呼ばれている、いくつかの広範なソフトウェア カテゴリのうちの 1 つのうち、1 つの特定の形式の 1 つの特定のアプリケーションを指していることになります。そして、この AI のサブサブサブ分野は、その環境への影響、非倫理的に収集されたトレーニング データ、雇用主が AI を人間の労働者の代わりにできると自慢し続けるという事実、そして AI が売りにしていることのほとんどで実際にはあまり得意ではないという事実から、まったく人気がありません。ですから、多くの人が「AIは嫌いだ！」と言っています。
この用語の固有の曖昧さにより、「AI」という用語を自社の限られた製品とすぐに結び付けてきた同じテクノロジー企業の CEO が完璧な復活を可能にします。「AI が嫌いですか? でも、AI が成し遂げた素晴らしいことをすべて見てください! あなたの携帯電話がどのように認識するかを見てください。」

画像に写っている人を見つけてタグ付けしてください!リアルタイムであなたの写真を漫画の猫に置き換えることができる様子を見てください。以前は検出されなかったがん細胞を医師がどのようにして特定できるのかを見てみましょう。 AI が嫌いなら、これらすべても嫌いに違いありません。」そしてもちろん、これらのものはいずれも LLM を使用しません。テクノロジー企業があらゆるものにチャットボットを導入し始める前から、チャットボットはあちこちに存在していました。それらは、いかなる形であっても、大多数の「AI 嫌い」が「AI」と言うときに話しているものではありませんし、テクノロジー仲間たちはそれを知っています。反対意見を表明しにくくするために言葉をねじ曲げることは、権威主義者が長年好んで使ってきた手法だからだ。ダブルプラスダメ。
まあ、私はそれには落ちません。可能な限り、私は「人工知能」や「AI」という用語の使用を一切避け、代わりに自分が話していることを正確に言うようにしています (LLM、ロジスティック回帰、サポート ベクター マシン、エキスパート システム、Lisp など)。
「反AI」もあまり意味がない
そのため、「AI」という用語の曖昧さにより、LLM シラーは、実際には反対していないものに反対しているとして相手を攻撃することができます。しかし、単に「AI」を「大規模な言語モデル」を意味するものとして再定義したらどうなるでしょうか?では、私たちは自分たちを「反AI」と呼ぶことができるでしょうか？
まあ、いいえ、LLMブースターがその用語を宇宙に植え付けたので。 PauseAI のような、反 AI のように見える組織も存在しますが、大手 AI 企業の CEO らの支援を受けています。そして、彼らの苦情を詳しく見てみると、環境への懸念、知的財産、スキルの低下、またはLLMについて人々が実際に抱いているその他の懸念についてはほとんど言及されていません。代わりに、「チャットボットは非常に賢く、指数関数的に進化している」という主張があります。

もっと賢い！彼らが世界を乗っ取り、私たちをクリップに変えようとするのは時間の問題です。したがって、私たちはすべての AI 開発を一時停止し、避けられず決して止めることのできない AI 開発が正しい方法で行われるように、厳格な規制と世界的な管理機関を設立する必要があります。」
この議論には 2 つの目的があります。 1 つは、LLM が非常にインテリジェントであり、ますます賢くなっていることを前提としており、人々がそれを受け入れて、「これらは本当に賢いのか?」という問題を乗り越えることを期待しているということです。 「これらの信じられないほどスマートなマシンをどうすればよいでしょうか?」という質問に真っ直ぐに答えます。しかし、この前提は明らかに間違っています。知能の正確な定義を特定するのは難しいですが、知能の最も広範な基準さえ満たさないものもあります。大規模な言語モデルは学習しません。これらは、パーセプトロン ノード間の重みがトレーニング データに基づいて変更される 1 回限りの「トレーニング フェーズ」を受けます。このフェーズが終了し、ネットワークが本番環境に展開される前に、重みは琥珀色に化石化し、二度と変更されることはありません。多層パーセプトロンにより多くの「情報」を取り込む唯一の方法は、より多くのトレーニング データを使用して新しいネットワークをトレーニングし、それを置き換えることです。確かに、チャットボットはユーザーの言ったことを覚えているように見えるかもしれませんが、それは単なる応対です。チャットボットにクエリを実行するたびに、これまでの会話の履歴全体 (または少なくともコンテキスト ウィンドウ内に収まる履歴全体) がネットワークにフィードバックされます。 LLM は実際には、非常に巨大な 2 進数 (エンコードされた会話) を受け取り、より小さい 2 進数 (予測された次のトークンに対応する) を返す単なる関数です。それには内なる生命がありません。プロンプトの間では何も考えず、実際に変更も行いません

全然。それが行うものをすべて「知性」と呼ぶことは、知性の概念そのものを侮辱することになります。したがって、当然のことながら、「これは知的ですか?」と尋ねます。それはLLMシラーが決してあなたに望んでいないことです。
「AI の一時停止」議論のもう 1 つの目標は、より平凡なものです。それは、規制による捕捉です。今日の大手 AI 企業が世界の政府を説得して、AI 企業が作成を支援する AI 規制を制定し、AI 企業が選択を支援できる政府機関を設立することができれば、潜在的な競合他社の参入コストを引き上げながら、同時に自分たちがやりたいことを確実に実行できるようになります。
実際には、3 番目の目標、または少なくとも嬉しい予期せぬ結果があるかもしれません。チャットボットが超強力になって世界を征服できるという考えは明らかにばかばかしいため、ほとんどの人はそれを即座に無視するでしょうし、それは当然のことです。 LLMシラーが物語をコントロールして、「反AI」が「スカイネットを信じているが、クリップに奇妙なこだわりを持っている」と同義になるようになれば、どんな反対派もばかげているように聞こえるようになるだろう。
正直に言うと、私には悩みがありました
「LLM への反対」と「ランダムフォレスト分類器への反対」を区別し、LLM に関する正当な懸念を天文学的な反スカイネットの誇大広告から遠ざけるために、私は「反 AI」に代わるものを考え出しました。私はそれを「QUALLMS: 全ての大規模言語モデルを歌うのはやめよう」と呼んでいます。次のような利点があります。
漠然とした「人工」ではなく、特に LLM をターゲットにしています。

[切り捨てられた]

## Original Extract

Portfolio of Chip Hollingsworth

Home
Blog
I'm not anti-AI, but I have QUALLMS
I haven’t been shy about voicing my negative opinions of the current large language model hype. But I’m hesitant to describe myself as “anti-AI”. And it’s not because I have some sort of super-nuanced position like “LLMs are fine in some cases” or “I’d be OK with LLMs if only we could solve these problems”. I’m in fact adamantly opposed to the very assumptions on which the hype is based: that a statistical model of word co-occurrence frequency is a good foundation for a multi-purpose tool for answering questions or generating meaningful text or code. Train a local model on text acquired with the creators’ express permission and run it on a machine fueled entirely by renewable energy, and I still don’t think it’s a good idea. The fact that commercial LLMs disregard consent and copyright and consume ungodly amounts of resources only makes a fundamentally bad thing even worse .
So why don’t I call myself “anti-AI”? Well, because of two ambiguities: one in the definition of “AI”, and another in what it means to be “anti-AI”. And that’s why I’m proposing an alternate nomenclature to explain exactly what my position is.
Why “AI” is a mostly meaningless term
The term “artificial intelligence” is usually believed to have come into existence around the time of the Dartmouth Summer Research Project on Artificial Intelligence in 1956. Even then, it wasn’t exactly clear what the term referred to; it was chosen as a sort of blanket term for a lot of semi-related research in information theory, computer science, and other disciplines. Broadly speaking, it seemed to refer to attempts to mimic the effects of human thought with computers. I say “mimic the effects of human thought” rather than “mimic human thought,” because the people involved were under no illusion that the processes by which computers handled knowledge representation, language processing, visual recognition, or other “AI” endeavors was anything like what was happening in the human brain. Absolutely nobody thought that human brains had Lisp interpreters running in them to parse all the sentences we hear. That’s why they used the word “artificial”: to stress that this was an imitation of a thing, not the thing itself. (The belief that brains really do some form of computation similar to what goes on inside a digital computer, and that we can design computer systems that actually mimic the human brain, belongs to the related but separate discipline of cognitive science .)
In the beginning, most things that were called “artificial intelligence” were rule-based systems, not essentially different from other types of computer programs. The main difference was that while computing had traditionally been used for tasks that were relatively simple to express in algorithmic form, such as calculating projectile trajectories or solving algebraic equations, artificial intelligence sought to conquer less obviously algorithmic territory, such as representing the information content of natural-language text, reasoning about it, and forming a natural-language reply. But over time, “machine learning” became a thing: using automated statistical or quasi-statistical methods to generalize from a set of input data. This might involve sorting pictures of food into the broad categories “hot dog” or “not hot dog,” or predicting what a website customer might want to buy next given the last several things they’ve bought. This represented a very different approach than what came to be known as “good old-fashioned artificial intelligence” (GOFAI): instead of working out the steps needed to perform a task and telling the computer to do them, you just threw a bunch of data at a program and let the computer figure out what to do. At least, that’s how it’s often anthropomorphically described. In fact, it’s just the same rule-based approach at a higher level of abstraction: instead of rules for turning inputs into outputs, it’s higher-order rules for turning training data into rules for turning inputs into outputs. But it’s a lot easier to say the computer “learns” how to do things, which unfortunately convinces a lot of people that machine learning systems are doing something a lot more similar to the human brain than they really are.
So now “artificial intelligence” was being used to refer to both old-school algorithms for manipulating symbols, and new-fangled statistical trickery. Two very different technologies described with the same terminology, their only commonality being a vague desire to emulate what people do. Over time, as computing hardware got cheaper and more powerful, “deep learning” became a thing. This was really just a multi-layer perceptron, a slightly more advanced version of a machine learning method dating back to the 1950s. Only now, we could make them much bigger, and train them on much bigger datasets. At around the same time, the Internet was really taking off, making it easier than ever to scoop up terrifying amounts of data, especially natural-language data. Even though deep learning wasn’t a fundamentally different kind of thing than earlier forms of machine learning, the scale involved allowed us to do a lot of things (image recognition, automatic translation) that earlier methods were only sort of OK at. So a lot of people started viewing deep learning as its own thing, making yet one more category under the ever-expanding umbrella of “AI”.
Most recently, people realized that if you take a certain kind of deep learning architecture, and an enormous amount of text, you can train it to predict what the next word will be given several hundred preceding words. And they wrapped a framework around this that would simulate dialogue, with a human providing one half of the conversation and the word-predictor providing the other. And it was so good at producing human-sounding dialogue, and the venture capital money was so plentiful, that companies started shoving it into every software product on the market. But they didn’t call it “large language models” or “next-token predictors” or even “chatbots”. They just called it “AI”.
The result is that now, when somebody says “AI,” there’s about a 99% chance they’re referring to one specific application of one particular form of one of several broad categories of software that have all, at various points in time, been called “AI”. And this one particular sub-sub-sub-discipline of AI is quite unpopular, for its environmental impact, its unethically sourced training data, the fact that employers keep bragging that it will let them replace human workers, and the fact that it isn’t actually very good at most of the things it’s sold to do. So lots and lots of people are saying “I hate AI!”
The inherent ambiguity of the term then allows the same tech CEOs who have been so quick to associate the term “AI” with their narrow selection of products the perfect comeback: “You hate AI? But look at all these wonderful things AI has done! Look at how your phone can recognize people in your images and tag them! Look at how it can replace your picture with a cartoon cat in real-time! Look at how doctors can identify cancer cells that would previously have gone undetected! If you don’t like AI, you must hate all these things too!” And, of course, none of these things use LLMs in any way. They were all around before tech companies started shoving chatbots into everything. They are not, in any way, shape, or form, what the vast majority of “AI haters” are talking about when they say “AI,” and the tech bros know it. They just don’t care, because twisting language to make it more difficult to express dissent has long been a favorite trick of authoritarians. Double plus ungood .
Well, I’m not falling for it. Wherever possible, I’m trying to avoid using the term “artificial intelligence” or “AI” at all, and instead say exactly what I’m talking about (LLMs, logistic regression, support vector machines, expert systems, Lisp, whatever.)
“Anti-AI” isn’t very meaningful either
So the ambiguity of the term “AI” lets LLM-shillers attack their opponents by accusing them of opposing things they’re not actually opposed to. But what if we simply redefined “AI” to mean “large language models”? Could we then call ourselves “anti-AI”?
Well, no, because LLM boosters have astroturfed that term. There are organizations like PauseAI out there that might seem to be anti-AI, but have the support of some of the biggest AI company CEOs. And when you look closer at their complaints, there’s scarcely any mention of environmental concerns, intellectual property, de-skilling, or other concerns that people actually have about LLMs. Instead, the argument is that “chatbots are super-duper smart and getting exponentially smarter! It’s only a matter of time before they try to take over the world and turn us into paperclips ! So we have to pause all AI development and set up strict regulations and global governing bodies to make sure AI development, which is inevitable and can never be stopped, is done the right way!”
This argument has two goals. One, it takes it as given that LLMs are extremely intelligent and getting ever smarter, with the hopes that people will just accept it, and move past the question of “are these things even smart?” straight to “what do we do about these incredibly smart machines?” The premise, however, is demonstrably false. While the precise definition of intelligence is tough to pin down, some things just fail to meet even the broadest criteria for intelligence. Large language models do not learn. They undergo a one-time “training phase,” during which the weights between their perceptron nodes are modified based on training data. Once that phase is over, and before the network is ever deployed to production, the weights are fossilized in amber and will never change again. The only way to get more “information” into a multi-layer perceptron is to train a new network with more training data to replace it. Sure, chatbots might seem to remember things you tell them, but that’s just a parlor trick: every time you query a chatbot, the entire history of the conversation so far (or at least the entire history that will fit inside its context window) is fed back to the network. An LLM is really just a function that takes a very huge binary number (the encoded conversation) and returns a much smaller binary number (corresponding to the predicted next token). It has no inner life. It doesn’t think about anything in between prompts, or in fact change at all. To call anything it does “intelligence” insults the very concept of intelligence. So of course, asking “is this thing intelligent?” is something the LLM-shillers never, ever want you to do.
The other goal of the “Pause AI” argument is more prosaic: regulatory capture. If today’s big AI companies can convince the world’s governments to enact AI regulations that the AI companies help write, and form a government body that AI companies get to help pick, they can ensure that they get to do whatever they want to do, while simultaneously raising entry costs for potential competitors.
Actually, there might be a third goal, or at least a happy unintended consequence: the idea that chatbots can become super-powerful and take over the world is so patently ridiculous that most people would dismiss it out of hand, and rightfully so. If LLM-shillers can take control over the narrative so that “anti-AI” comes to be synonymous with “believes in Skynet but with an odd paperclip fixation,” they can make any opposition sound ridiculous.
I must admit, I had my QUALLMS
In order to distinguish “opposition to LLMs” from “opposition to random-forest classifiers,” and to distance the legitimate concerns about LLMs from the astroturfed anti-Skynet hype, I’ve come up with an alternative to “anti-AI”: I call it QUALLMS: Q uit U sing A ll L arge L anguage M odel s . It has the following advantages:
It specifically targets LLMs, rather than vague “artificial

[truncated]
