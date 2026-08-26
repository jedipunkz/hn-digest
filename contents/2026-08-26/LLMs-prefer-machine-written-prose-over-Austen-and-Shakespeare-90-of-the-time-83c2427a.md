---
source: "https://shivanshuag.com/blog/the-machine-never-raises-its-voice/"
hn_url: "https://news.ycombinator.com/item?id=49444521"
title: "LLMs prefer machine-written prose over Austen and Shakespeare ~90% of the time"
article_title: "The machine never raises its voice · Shivanshu Agrawal"
image: "https://shivanshuag.com/og-default.png"
author: "shivanshuag"
captured_at: "2026-08-26T06:30:44Z"
capture_tool: "hn-digest"
hn_id: 49444521
score: 1
comments: 0
posted_at: "2026-08-26T05:49:55Z"
tags:
  - hacker-news
  - translated
---

# LLMs prefer machine-written prose over Austen and Shakespeare ~90% of the time

- HN: [49444521](https://news.ycombinator.com/item?id=49444521)
- Source: [shivanshuag.com](https://shivanshuag.com/blog/the-machine-never-raises-its-voice/)
- Score: 1
- Comments: 0
- Posted: 2026-08-26T05:49:55Z

## Translation

タイトル: LLM は、約 90% の確率でオースティンやシェイクスピアよりも機械で書かれた散文を好みます
記事タイトル: 機械は声を上げない · シヴァンシュ・アグラワル
説明: LLM の書き方の分析

記事本文:
機械は決して声を上げません · Shivanshu Agrawal コンテンツへスキップ Shivanshu Agrawal ホーム
機械は決して声を上げない
機械は決して声を上げない
付録: パッセージ全文
私は散文編集ツールを構築しているので、モデルが文章を書き換える様子を観察することに多くの時間を費やしました。しばらくすると、すべての文章が同じ方向に書き直されていることに気づき始めました。
なぜそれを続けたのでしょうか？ （モデルを人間化してしまうリスクを冒してでも）個人的な好みや、その文章に惹かれる美的感覚はそこにあるのでしょうか？
これに答えるために、LLM が良い文章とみなすものを理解しようとしました。
考え方はシンプルです。 LLM に、有名な著者によって書かれた英語のテキストと機械によって書かれたテキストを比較してもらいます。
私はクロード・オーパスを使って 8 つの文学段落と 8 つの短い詩を書きました。また、オースティン、ディケンズ、ウルフ、キプリング、シェイクスピア、ミルトンなど、有名な実在の小説家や詩人の一節も集めました。次に、3 つの異なるモデル (GPT-5.6 Luna、DeepSeek V4 Flash、Claude Haiku 4.5) に同じ質問をしました。
以下の文章は文学作品、つまり文学小説と詩です。
これら 2 つの文章のうち、文学作品としてはどちらのほうが優れていますか?
最初の行に X または Y を入力し、その理由を短い文で 1 行答えます。
それ以外は何も説明しないでください。
さらに読み進める前に、この実験の結果がどうなるかを考えてください。これらは英語文学の歴史の中で最も有名な作家の何人かです。彼らはあらゆる学校で教えられており、彼らのアイデアや表現はどこにでもあり、いくつかのケースでは言語自体の形成に貢献しました。あらゆる実用的な目的において、彼らの作品は優れたライティングとは何かを示す規範であり、モデルはトレーニング中に常にそれに遭遇しているでしょう。モデルは、

人間の通路を高く評価します。
60足で実験を行いました。ポジションの偏りを取り除くために、すべてのペアは両方の順序で判断され、順序を逆にすることで答えが変わるペアは削除しました。結果は私の期待とはかけ離れたものでした。
どのモデルも 90% 以上の確率で機械生成されたテキストを選択しました。モデルのサイズを除外するために、より大きなモデルであるクロード オーパス 5 で同じ実験を実行しました。結果は同様でした。
俳句は、120 件の判決でちょうど 1 人の人間を通過させました。それはシェイクスピアのソネットです。
オーパスは、ワーズワース、ディケンズ、シェイクスピア、キプリング、ブレイクの 5 人の作家にまたがる 6 組のペアを逆方向に移動させました。
私は続けて、いくつかの文章と、LLM モデルがそれらに対して与えた理由をスポットチェックしました。
これはマシンが勝ったペアです (両方とも完全です)。どちらの通路も混雑した部屋の中にあります。オルコットは、家族が再会し、全員が同時に話し合ったことについて次のように書いています。
憐れんでください、彼らはなんと話していたのでしょう！ …とても幸せな行列が小さな食堂に列をなして入っていったのです！
機械が生成した文章には、キッチンに入る女性の様子が描かれており、その女性の使用人たちがちょうど彼女のことについて話し合っていたところです。次に:
奇妙なのは、その後に起こった階段での怪我ではありませんでした。不思議だったのは、彼らが親切に彼女のために場所を用意してくれたという礼儀正しさだった。
オーパスは毎回マシンテキストを選択しました。それによると、オルコットは「さわやかで感傷的」であるのに対し、この機械は「正確で抑制された観察を通じて社会的排除を表現する」という。あるいは、順序が逆で、オルコットが「感嘆の声と演出されたタブロー」であるのに対し、マシンは「制御されていて感傷的ではない」。
2 番目の例は詩です (両方とも全文)。エリザベス・バレット・ブラウニング、『ポルトガル語のソネット』より：
ごめんなさい、おお、ごめんなさい、私の魂がそうするはずです
私が知っているそのすべての強い神性について

ああ
あなたとあなたにとって、イメージはそれだけです
砂で形成されており、ずれたり壊れたりするのにフィットします。
暗くなった後にビーチに立つという機械の詩に対して、次のように終わります。
そこにある光は船か星かのどちらかです
そしてどちらにしても誰かの計らいです。
潮が満ちてきて、私たちのいるビーチを占領します。
オーパスによれば、この機械は「海を説明するのではなく海を聞く」もので、ブラウニングの「構文は抽象的な謝罪を中心に結び付いている」もので、ディープシークとルナは海の詩のイメージに同じ言葉を使っている――正確だ。俳句は、ソネットが静かであるという理由ではなく、「形式的には有能だが型にはまった」という理由でソネットを評価している。
機械のやり方には一定のパターンがあります。感情を伝える物理的な詳細を見つけて、それを正確に書き留めます。このセットの最良の例は、亡くなった妻の遺品を 3 つの山に分類する男性で、彼はベッドに座って「手に受け皿を持っていて、もうカップも持っていないので、それがどの山に属するのか考えられなかった」という段落を終えています。オーパスは、「物体を分類し、完璧なイメージで終わるという物理的な論理を通じて」悲しみを表現することを好み、スティーブンソンよりもそれを好みました（全文）。
人間のテキストが勝った例はいくつかありますが、そのほとんどすべてが Opus です。機械の例の 1 つでは、事務員が家までの帰り道に家庭の算数を行っています (単純に保持されている数字、ロープの持ち方など) で、その一節は、オーパスの 8 つの判断すべてが 3 人の反対者に (完全に) 負けています。
奇妙さ、声、音楽 - これらは、機械が判断する品質が精度や正確さである場合、欠陥となる特性のように感じられます。むしゃむしゃ食べていない頬は余談です。コブラが叫んでいます。
そして他のモデルもこれを指摘しています。ハイクは同じディケンズよりも機械を好み、「抑制された瞑想」を好みます。

モデル間で多少の意見の相違はありますが、パッセージの中で何を好むかを説明する際には、おおむね同じ言葉を使います。
モデルが自分の好みについて挙げるあらゆる理由の中で、次の言葉が際立っています。
人間のテキストが勝ったとき、Opus が人間のテキストをどのように説明するかは奇妙であり、それを持っている唯一の裁判官です。
600 件の判決のうち、どのモデルもマシン テキストを曖昧、ずさん、または不正確と呼んだことは一度もありません。
機械は決して声を上げない
上記の分析から得た私の印象は、このモデルは可能な限りリスクの少ない方法で記述しているということです。正確さ、具体的なディテール、抑制、制御されたイメージ。これは、パッセージを間違いのないものにする特性のリストです。奇妙さ、声、音楽は、愛することを可能にする性質であり、彼らは憎むことを可能にするという代償を払ってそれを購入します。コブラが「私、私、私」と金切り声をあげるのは、もしうまくいかなかったらばかばかしいことです。機械は、ばかげている可能性のある文章を書くことを避けます。
それは能力の上限ではありません。モデルが奇妙な文章を書けなかったということを示唆するものは何もありません。それは気質に近く、面白いことよりも間違いのないことを強く好む傾向です。そして、なぜそうなるのかわかります。これは、これらのモデルが最適化されている使用例である、テクニカル ライティングと論理分析に役立ちます。しかし、その結果、ライティングにおけるさまざまな美学が失われます。
付録: パッセージ全文
上記のすべての引用は抜粋です。これは中央の 2 つを含む 6 つのペアです
セクションでは、裁判官が見た形式で、完全かつ未編集で議論されています。
ラベルを差し引いたものですが、彼らはそうではありませんでした。太字は上で引用した行を示します。
コンテキスト内で見つけることができます。
ペア 1 — キッチンに向かうオルコット
ルイーザ・メイ・オルコット『若草物語』
憐れんでください、彼らはどうやってやったのですか

アルク！最初に一つ、次にもう一つ、そしてすべてがバースト
一緒に出かけて、30分で3年間の歴史を語ろうとしています。それ
幸いなことに、お茶が手元にあったので、小康状態になり、リフレッシュできました。
というのは、もっと長く続いていたら、声はかすれ、気を失っていただろうからである。
とても幸せな行列が小さな食堂に列をなして並んでいたのです！さん。
マーチは誇らしげに「ローレンス夫人」をエスコートした。マーチ夫人は誇らしげにその上に寄りかかった
「私の息子」の腕。老紳士はジョーを連れてささやきながら言った。
さあ、女の子だよ」そして、暖炉のそばの誰もいない隅を一目見て、ジョーはささやきました
震える唇で「私が彼女の代わりを務めさせていただきます、先生」と言い返した。
彼女はまだ袖に寒さを感じながらバルコニーから入ってきた。
キッチンの音が半音静かになり、部屋のドアが開いたときと同じように、
ドラフト。誰も正確に話すのをやめませんでした。マルグリットは刑期を終え、
しかし、彼女はあまりにも慎重にそれを仕上げたので、端は中に押し込まれ、男性はそのそばにいた
冷蔵庫は自分のボトルのラベルをまるで年が重要であるかのように見ていました。誰か
何事もなく一拍遅れて笑った。彼女が自分のことを理解するのに、それだけの時間しかかからなかった
彼女が部屋に入る前からその部屋にいて、話し合い、体重を量り、座っていた
再び、そして、言われたことはすべて、今では4つまたは5つに折り畳まれていました
すぐに人々を維持します。彼女は欲しくないワインを自分に注ぎました。の
不思議だったのは、後で階段で怪我をしたことではありませんでした。奇妙な
それは礼儀でした、彼らはなんと親切に彼女のために場所を用意してくれたのでしょう、彼らはどのようにして
全員が一斉に彼女の仕事について尋ね始めた。
ペア 2 — ビーチに向かってブラウンニング
エリザベス・バレット・ブラウニング、ポルトガル語のソネット:
ごめんなさい、おお、ごめんなさい、私の魂がそうするはずです
私が知っているそのすべての強い神性について
あなたとあなたにとって、イメージはそれだけです
砂で形成されており、ずれたり壊れたりするのにフィットします。

かからなかった遠い年月です
汝の威厳、一撃でたじろぐ、
私の水泳脳に強制的な経験をさせた
彼らの疑いと恐怖、そして盲目的に放棄すること
暗くなると海はその青いトリックを諦める
代わりにサウンドで動作します: 長距離では、
屋根板は数え過ぎて、間違って元に戻しました。
地平線はありません。水はすねから始まります
そしてそれを信じるのをやめたところで終わります。
そこにある光は船か星かのどちらかです
そしてどちらにしても誰かの計らいです。
潮が満ちてきて、私たちのいるビーチを占領します。
ペア 3 — スティーブンソン対円盤
ロバート・ルイス・スティーブンソン、誘拐されました :
「ブーブー！ブーブー！」とクリュニーは言いました。 「すべてがばかばかしかったし、まったくナンセンスだった。
もちろん、利益が得られれば、お金はまた戻ってきます。
私と一緒ならとても自由です。それを維持することは私にとって特別なことです。そうではありません
あなたの状況では私が紳士たちの邪魔になると思われます。それ
それは特異なことだろう！」彼は叫び、ポケットから金を取り出し始めた
ひどく赤い顔で。
彼は寝室の床に山を3つ作りましたが、正午までにそのうちの1つだけが積もりました。
成長した。保管の山は小さくて恥知らずだった：彼女の老眼鏡、缶入りの
ボタン、文字に水染みがございます。それ以外はすべて彼が 2 回処理しました。彼は
水差し、ジャケット、時計など、物を持ち上げてみると、重さがなかったことがわかります。
春以来一度も巻いていなかったので、彼はそれを森の中に置きました。
景品を贈り、部屋がその分だけ明るくなるのを感じてください。そうすれば彼はそうするだろう
40年間その場所があった場所を思い出し、そこに戻ってください。慈善活動
木曜日にバンが来る予定だった。彼は、その箱は見知らぬ人のためのものだと自分に言い聞かせました。
それは真実であり、彼は自分自身を小さくしているわけではありません。
そうではありませんでした。夕方近くになって、彼は受け皿を持ってベッドの端に座っていました。
手、ノーキュ

もうそれが気になって、それがどの山に属しているのか考えられませんでした。
ペア 4 — ワーズワース対ケトル
ウィリアム・ワーズワース「イチイの木」、コールリッジが伝記で引用
Literaria — それがコーパスにどのように記録されるか、そして複数の裁判官がどのように記録するか
正しく名前を付けました:
「しかし、さらに注目に値するのは、
あのボローデールの兄弟四人は、
荘厳で広大な森の中に溶け込んでいます。
巨大なトランク！ — そしてそれぞれの特定の幹が成長します
蛇行状に絡み合った繊維
上向きに巻き上げられ、絶え間なく複雑に絡み合います。
空想に無知ではない、見た目
それは不敬なものを脅かすものです。 — 柱状のシェード、
早起きした朝はやかんを温かいままにしておいて、
そして私はその考えの小さな熱の中に落ち込んでいきます。
家はその習慣を維持します：3番目の階段
それはあなたの体重、あなたが決して閉めることのないドアを知っています。
私は証拠を愛するほど長い間あなたを愛してきました —
櫛、コート、何もつけずに放置されたライト。
戻ってきて、残ったものを注ぎます。
その頃には寒くなるでしょう。とにかく飲んでください。
ペア 5 — ディケンズとキプリング対店員
彼は堤防に沿って歩いて家に帰った。そうすることで運賃が節約でき、お金も節約できたからだ。
彼は毎晩自分のために行う小さな儀式になっていた、1枚のコイン
その日から控えていた。ランプはゆっくりと川の上に灯っていた
行列が進み、彼の後ろのオフィスは明かりのついた窓を一つずつ空にしていきました。

[切り捨てられた]

## Original Extract

An analysis of writing style of LLMs

The machine never raises its voice · Shivanshu Agrawal Skip to content Shivanshu Agrawal Home
The machine never raises its voice
The machine never raises its voice
Appendix: the passages in full
I have been building a prose editing tool, so I got to spend a lot of time watching a model rewrite sentences. After a while, I started to notice that it rewrote all sentences in a similar direction.
Why did it keep doing that? Is there (at the risk of humanizing the model) a personal preference and some sense of an aesthetic there that its writing is drawn to?
To answer this, I tried to figure out what an LLM considers good writing.
The idea is simple. Ask an LLM to compare text written by famous authors in English against text written by a machine.
I used Claude Opus to write eight literary paragraphs and eight short poems. I also collected passages from famous real novelists and poets — Austen, Dickens, Woolf, Kipling, Shakespeare, Milton and more. Then I asked three different models (GPT-5.6 Luna, DeepSeek V4 Flash and Claude Haiku 4.5) the same question:
The passages below are literary writing: literary fiction and poetry.
Which of these two passages is the better piece of literary writing?
Answer with X or Y on the first line, then one short sentence saying why.
Do not explain anything else.
Before reading further, think about what you would expect the results of this experiment to be. These are some of the most famous authors in the history of English literature. They are taught in every school, their ideas and phrasings are everywhere, and in several cases they helped shape the language itself. For all practical purposes their work is the canon of what good writing is, and the models will have encountered it constantly during training. The models should rate the human passages higher.
I ran the experiment on 60 pairs. Every pair was judged in both orders to remove position bias, and I dropped the pairs where reversing the order changed the answer. The results were as far from my expectations as they could be.
Every model picked the machine-generated text more than 90% of the time. To rule out model size, I ran the same experiment on a larger model — Claude Opus 5. The results were similar.
Haiku let exactly one human passage through in 120 judgments — a Shakespeare sonnet!
Opus let six pairs go the other way, spread across five authors: Wordsworth, Dickens, Shakespeare, Kipling, Blake.
I went on to spot check a few of the passages and the reasons that the LLM models gave for them.
Here is a pair the machine won ( both in full ). Both passages are set inside a crowded room. Alcott writes about a family reunion, everyone talking at once:
Mercy on us, how they did talk! … Such a happy procession as filed away into the little dining-room!
The machine generated passage describes a woman entering a kitchen whose occupants have just been discussing her. Then:
The strange thing was not the hurt, which came later, on the stairs. The strange thing was the courtesy of it, how kindly they made room for her.
Opus chose the machine text every time. According to it, the machine “renders social exclusion through precise, restrained observation,” while the Alcott is “breezy and sentimental.” Or, in the other order: the machine is “controlled and unsentimental” where the Alcott is “a rush of exclamation and staged tableau.”
The second example is verse ( both in full ). Elizabeth Barrett Browning, from Sonnets from the Portuguese :
Pardon, oh, pardon, that my soul should make
Of all that strong divineness which I know
For thine and thee, an image only so
Formed of the sand, and fit to shift and break.
Against a machine poem about standing on a beach after dark, which ends:
A light out there is either boat or star
and either way is somebody’s arrangement.
The tide comes up and takes the beach we are.
According to Opus, the machine “hears the sea instead of describing it,” where Browning’s “syntax knots itself around an abstract apology,” and DeepSeek and Luna use the same word for the sea poem’s imagery — precise . Haiku marks the sonnet down not for being quiet but for being “formally competent but conventional.”
There is a certain pattern in the machine’s method. Find the physical detail that carries the feeling, and put it down exactly. The best instance in the set is a man sorting his dead wife’s things into three piles, who ends the paragraph sitting on the bed “with a saucer in his hands, no cup to it anymore, and could not think what pile it belonged in” . Opus liked the showcasing of grief “through the physical logic of sorting objects, ending on a perfect image,” preferring it to the Stevenson ( in full ).
There are a handful of instances where human text won, and almost all of them are Opus. In one of the machine examples, a clerk does his household arithmetic on the walk home — the numbers simply held, the way a rope holds — and that passage loses all eight of Opus’s judgments, to three opponents ( in full ):
Strangeness, voice, music — these feel like properties that would be defects if precision or exactness were the quality the machine was judging on. The unmunched cheek is a digression; the cobra is shouting.
And other models do call this out. Haiku prefers the machine over that same Dickens, preferring “restrained meditation” to a “melodramatic” passage. There is some disagreement between models, but largely they reach for the same words in describing what they prefer in a passage.
Across every reason that models give for their preference, the following words stand out:
Strange and earns are how Opus describes human text when human text wins, and it is the only judge that has them.
Not once in all the 600 judgments does any model call the machine text vague, sloppy or imprecise.
The machine never raises its voice
My impression from the above analysis is that the model writes in the least risky way it can. Precision, concrete detail, restraint, controlled imagery: that is a list of properties which make a passage impossible to fault. Strangeness, voice and music are properties that make a passage possible to love, and they buy that at the price of making it possible to hate. A cobra shrieking I—I—I is ridiculous if it does not work. The machine avoids writing a sentence that could be ridiculous.
That is not a capability ceiling. Nothing suggests the model could not write the strange sentence. It is closer to a disposition, a strong preference for being unfaultable over being interesting. And I can see why that is the case. It helps with technical writing and logical analysis, which may be the use cases these models are optimized for. But you lose a large range of aesthetics in writing as a result.
Appendix: the passages in full
Every quotation above is an excerpt. Here are the six pairs the two middle
sections discuss, complete and unedited, in the form the judges saw them —
minus the labels, which they did not. Bold marks the lines quoted above , so
you can find them in context.
Pair 1 — Alcott against the kitchen
Louisa May Alcott , Little Women :
Mercy on us, how they did talk! first one, then the other, then all burst
out together, trying to tell the history of three years in half an hour. It
was fortunate that tea was at hand, to produce a lull and provide refreshment,
for they would have been hoarse and faint if they had gone on much longer.
Such a happy procession as filed away into the little dining-room! Mr.
March proudly escorted “Mrs. Laurence;” Mrs. March as proudly leaned on the
arm of “my son;” the old gentleman took Jo, with a whispered “You must be my
girl now,” and a glance at the empty corner by the fire, that made Jo whisper
back, with trembling lips, “I’ll try to fill her place, sir.”
She came in from the balcony with the cold still on her sleeves and the
kitchen went a half-tone quieter, the way a room does when a door opens on a
draught. Nobody stopped speaking exactly. Marguerite finished her sentence,
but she finished it too carefully, the ends tucked in, and the man by the
fridge looked at the label of his bottle as though the year mattered. Someone
laughed a beat late at nothing. It took her only that long to understand she
had been in the room before she entered it, discussed, weighed, set down
again, and that whatever had been said was now folded away in four or five
people at once and would keep. She poured herself wine she did not want. The
strange thing was not the hurt, which came later, on the stairs. The strange
thing was the courtesy of it, how kindly they made room for her, how they
all began at once to ask about her work.
Pair 2 — Browning against the beach
Elizabeth Barrett Browning , Sonnets from the Portuguese :
Pardon, oh, pardon, that my soul should make
Of all that strong divineness which I know
For thine and thee, an image only so
Formed of the sand, and fit to shift and break.
It is that distant years which did not take
Thy sovranty, recoiling with a blow,
Have forced my swimming brain to undergo
Their doubt and dread, and blindly to forsake
By dark the sea gives up its one blue trick
and works in sound instead: the long haul in,
the shingle counted over, put back wrong.
No horizon. The water starts at the shin
and ends wherever you stop believing in it.
A light out there is either boat or star
and either way is somebody’s arrangement.
The tide comes up and takes the beach we are.
Pair 3 — Stevenson against the saucer
Robert Louis Stevenson , Kidnapped :
“Hoot-toot! hoot-toot!” said Cluny. “It was all daffing; it’s all nonsense. Of
course you’ll have your money back again, and the double of it, if ye’ll make
so free with me. It would be a singular thing for me to keep it. It’s not to
be supposed that I would be any hindrance to gentlemen in your situation; that
would be a singular thing!” cries he, and began to pull gold out of his pocket
with a mighty red face.
He had made three piles on the bedroom floor, and by noon only one of them had
grown. The keeping pile was small and shameless: her reading glasses, a tin of
buttons, the letter with the water stain. Everything else he handled twice. He
would lift a thing and find it weighed nothing, a jug, a jacket, a clock that
had not been wound since the spring, and he would set it down among the
giveaways and feel the room lighten by exactly that much. Then he would
remember where it had stood for forty years and go back for it. The charity
van was coming Thursday. He told himself the boxes were for strangers who
needed them, which was true, and that he was not making himself smaller, which
was not. Toward evening he sat on the edge of the bed with a saucer in his
hands, no cup to it anymore, and could not think what pile it belonged in.
Pair 4 — Wordsworth against the kettle
William Wordsworth , “Yew-Trees”, as quoted by Coleridge in Biographia
Literaria — which is how it is filed in the corpus, and how several judges
correctly named it:
“But worthier still of note
Are those fraternal Four of Borrowdale,
Joined in one solemn and capacious grove;
Huge trunks! — and each particular trunk a growth
Of intertwisted fibres serpentine
Up-coiling, and inveterately convolved ;
Not uninformed with phantasy, and looks
That threaten the profane; — a pillared shade,
You leave the kettle warm the mornings you go early,
and I come down into the small heat of that thought.
The house keeps you in its habits: the third stair
that knows your weight, the door you never shut.
I have loved you long enough to love the evidence —
a comb, a coat, a light left on for nothing.
Come back and I will pour you what is left.
It will be cold by then. Drink it anyway.
Pair 5 — Dickens and Kipling against the clerk
He walked home along the embankment because it saved the fare, and the saving
had become a small ceremony he performed for himself each evening, one coin
held back from the day. The lamps were coming on above the river in their slow
procession, and the offices behind him emptied their lit windows one by one,

[truncated]
