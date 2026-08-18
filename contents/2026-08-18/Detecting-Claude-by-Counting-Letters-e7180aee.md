---
source: "https://www.atomic14.com/2026/08/18/detecting-claude-with-letter-counts"
hn_url: "https://news.ycombinator.com/item?id=49344384"
title: "Detecting Claude by Counting Letters"
article_title: "Detecting Claude by counting letters | atomic14"
image: "https://www.atomic14.com/assets/article_images/2026-08-18/meatbag.webp"
author: "iamflimflam1"
captured_at: "2026-08-18T12:25:04Z"
capture_tool: "hn-digest"
hn_id: 49344384
score: 1
comments: 0
posted_at: "2026-08-18T11:58:26Z"
tags:
  - hacker-news
  - translated
---

# Detecting Claude by Counting Letters

- HN: [49344384](https://news.ycombinator.com/item?id=49344384)
- Source: [www.atomic14.com](https://www.atomic14.com/2026/08/18/detecting-claude-with-letter-counts)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T11:58:26Z

## Translation

タイトル: 文字を数えてクロードを見つける
記事のタイトル: 文字を数えてクロードを検出する |アトミック14
説明: Anthropic にはクロードの透かしが入っていますが、自分たちで実行できる検出器が欲しかったので、構築しました。単語を完全に無視し、4 文字の n グラムをカウントし、それらを線形 SVM に供給し、確実に Clau にフラグを立てます。

記事本文:
🌈 ESP32-S3 Rainbow: ZX Spectrum エミュレータ ボード!
Crowd Supply で入手 →
文字を数えてクロードを検出する
« クロードを責めないで - それは私です、私が問題です、それは私です...
私の仕事をサポートしてください: 気分が高揚している場合は、Patreon にお立ち寄りください。または、ko-fi 経由で 1 回限りの寄付をすることもできます
2026 年 8 月 2 日、Anthropic は新しいクロード モデルが生成するテキストに透かしを入れ始めました。サポート ページにはどの製品に透かしが入っているかが記載されています。
ウォーターマークは隠し文字やメタデータではありません。それは統計的なものです。秘密キーは、人間には明らかではない方法で生成された出力を微調整しますが、キーを持っていれば検出できます。これは、人類はそれを検出できますが、私たち一般人は検出できないことを意味します。
他のプロバイダーも同様のことを行っているか、近いうちに行う予定です。
独自の検出器を構築できますか?テキストの一部を見て、モデルがそれを書いたことがわかりますか?
物事を単純化して、1 つの狭いターゲット、つまりクロード オーパス 5 によって書かれたテキストを試してみましょう。主に、それが私が購読しているものだからです…
そして、ネタバレ注意ですが、それができることが分かりました。その Web バージョンをここにデプロイしました。 Opus 5 を使用して Claude でテキストを生成し、それが検出できるかどうかを確認します。
そして、はい、このブログ投稿をフィードすると、私が大幅に編集したにもかかわらず、クロードが書いたものとして表示されます。
昔、私は博士号を取得し、その一環として、誰が文書を書いたかを解明しようとする問題文書検査について調べました。シェイクスピアが戯曲を書いたのか、それとも別の誰かが書いたのか?連邦党員の誰がどの論文を書きましたか?誰がこのような脅迫状を送ったのでしょうか?
非常に効果的だと思われる、本当に素晴らしいシンプルなテクニックがあります。多くの場合、短い文字数を数えることで作者を特定できます。単語や文の構造などは無視します。run の出現だけを無視します。

fの文字。これらの実行は N グラムと呼ばれます。
n グラムは、シーケンスから取得された n 個の項目の単なる実行です。単語または文字から構築できます。ここでは文字を使用しています。
「素早い茶色のキツネ」を連れて行きましょう。スペース、句読点、大文字を取り除くと、quickbrownfox が完成します。 4 文字のウィンドウをそれに沿って 1 つずつ位置をスライドさせ、表示された内容を記録します。
theq hequ equi qui quic uick ickb ckbr kbro brow rown ownf wnfo nfox
これをパッセージ全体に対して行うと、カウントの表が作成されます。出現するランごとに、that は 40 回、ting は 22 回、quic は 1 回というように表示されます。通路を捨ててください。カウント表は、検出器がこれまでに確認したすべてです。
カウントは正規化されているため、理論的にはテキストの長さは関係ありません。これには明らかに限界があります。テキストが短すぎると、カウントは統計的に有効ではなくなり、何も予測できなくなります。
統計的分布はライターとして意識されるものではなく、難読化するのが非常に難しいという考え方です。それは無意識の物語を暴露します。
この種の問題の場合と同様、難しいのはデータの取得ですが、実際の「科学」の部分は簡単です。
私には人間らしい、人間らしい文章が必要でした。 「間違いなく」の最も安全な定義は、「ChatGPT が存在する前に書かれた」ということなので、私が使用したものはすべて 2021 年より前に公開されたものです。Claude は、次の 3 つの場所からそれをまとめるのを手伝ってくれました。
プロジェクト グーテンベルクの書籍、すべて 1929 年以前。
2004 年以前に書かれた Blog Authorship Corpus のブログ投稿。非営利の研究用途のみ。
2016 年から 2019 年の間に発行された CC-News のニュース記事。
その結果、26 人の著者による 519 個の文章が得られました。そのうち 180 個は書籍から、179 個はニュースから、160 個はブログからでした。 3 つの異なる種類の書き込みと、ビクトリア州の日付のスプレッド

n からほぼ現代まで。
AI テキストを取得する明白な方法は、クロードに何かを書くように依頼することです。
そこで私は人間のテキストを書き換えるように依頼しました。人間のパッセージごとに、私はクロード作品 5 に、意味、詳細のレベル、おおよその長さを維持しながら、同じパッセージを再度作成するよう依頼しました。これにより、一致するコンテンツのペアが得られます。ペア内では、主題が同一で、事実が同一で、長さが近いため、それらを分離する唯一のものは文章自体です。
文章のスタイルは 1 つではないので、3 つの異なる方法を尋ねました。
plain - 「次の文章を明瞭な自然な英語で書き直してください。」
direct - 「次の文章を平易で直接的な現代英語で書き直してください。」
正式 - 「次の一節を正式な正確な記録簿に書き直してください。」
これは、バッチ API を通じて生成された 1,557 回の書き換えになります。また、クロードに同じトピックについて新しい文章を一から書かせましたが、それはすべてのトレーニングから除外し、後でテストとして使用しました。
まず、精度の測定方法についてです。下の数字は ROC AUC です。人間のパッセージと AI のパッセージをランダムに 1 つずつ取得し、両方のスコアを獲得し、AI のパッセージの方が高いスコアを得たかどうかを尋ねます。 ROC AUC は、ランダム ペアの割合です。
0.5はコイントスです。 1.0 は、ペアを逆に取得しないことを意味します。特定のスコアが正しいかどうかではなく、順序を測定します。
2番目に、分割です。何かをトレーニングする前に、26 人の著者のうち 4 人を脇に置き、それらには触れませんでした。他のものはすべて相互検証によって測定され、テストに使用された作成者はトレーニングに使用された作成者ではありません。文章ごとではなく著者ごとに分割することが重要です。パッセージごとに分割すると、分類器は個々のライターを認識することを学習します。これは非常に得意ですが、それは私たちが望んでいることではありません。
次に、n グラムの長さを 1 から 8 までスイープしました。

g 他はすべて修正されました:
個々の文字を数えることができます。 28 個の特徴 (文字の頻度だけ) は 0.651 に達します。それは大したことではありませんが、非常に単純な機能セットからすれば、コイントスのようなものではありません。
精度は 4 まで急上昇し、5 でピークに達し、その後低下します。より長い実行では、データが不足してサポートできなくなるまで、より多くの情報が含まれます。これは、単語カウントを使用することも非常にうまく機能することを示唆しているかもしれませんが、おそらくより多くのデータが必要になるでしょうか?研究すべきことがある…
また、連続した文字ではなく 2 文字ごと、または 3 文字ごとに文字を取得する、間隔をあけた N グラムも試してみました。これはまったく機能しませんでした。精度が大幅に悪化しました。
結局5つではなく4つを選んだ
ピークは 5 です。しかし、私は 4 を使用することにしました。これにはいくつかの理由があります。
1 つ目の理由は、5 文字にすると、より多くの機能を犠牲にして、精度がわずかに向上するということです。
2 つ目は、長いテキスト部分のみを改善できることです。その理由は、データが非常に薄く拡散するためです。 4 グラムの可能性は 26⁴ = 457,000 個、5 グラムの可能性は 26⁵ = 1,190 万個あります。
私のコーパスから 2,259 文字の長さの典型的な一節を抜粋すると、次のようになります。
5 文字では、パッセージ内のランの 10 分の 9 が 1 回しか現れず、テキストが短くなるにつれて状況は悪化します。200 文字では、3 グラムの 90% と 5 グラムの 97% が 1 回しか現れません。
以下は、各長さで構築された同じ検出器です。クロードがプロンプトから書いた 320 個の文章に対して、公開ベンチマークで 320 個の人文ソース文書に対してテストされ、テキストはさまざまな長さに切り詰められています。
どの長さでも4グラムが勝つ。
完成したものは複雑ではありません。
テキストを受け取ります。文字以外の部分 (スペースも含む) をすべて削除し、残った部分を小文字にします。
4 グラムをすべて数えます。それぞれの通路は、

60,211 個の番号のリスト。
正規化すると、数値は比例し、通路の長さに依存しません。
それを線形サポート ベクター マシンにフィードすると、クロードの文章でどのカウントが上位にあるかが学習されます。
結果を 0 から 1 までの数値に変換します。大きいほどクロードに似ていることを意味します。
N-gram カウントと分類器は両方とも標準の scikit-learn コンポーネントです。
ステップ 1 のため、句読点、文の長さ、段落構造、大文字の使用、または個々の単語を確認できません。ステップ 3 のため、パッセージの長さがわかりません。文字の比率のみが参照されます。
最初に取っておいた 4 人の著者は、最後に 1 回採点されます。
2 つの数字は接近しており、それが私が最も見たかったものです。これは、分類器が特定の著者を認識することを密かに学習していなかったということを意味します。
次に、トレーニング データにまったく関係のないテキストを指定しました。まず、クロードが書いた 8 つの文章に対して、2006 年から 2022 年の間に書かれた私自身のブログの 175 件の投稿です。
2番目に、より難しいテストです。 RAID は、機械生成テキスト検出の公開ベンチマークであり、8 つの主題ドメインにわたる 11 のジェネレーターからの 600 万個のパッセージが含まれています。私はそのプロンプトを 320 個取得し、Claude Opus 5 に回答してもらい、そのプロンプトの元となった 320 個の人文書と照らし合わせて回答を採点しました。異なるコーパスと異なるタスク。これらは書き直されたのではなく、最初から書かれたためです。スコアは 0.889 です。
そして、私がまったく訓練を受けずに保持していた、4 人の未見の著者のトピックについてゼロから書いた新鮮な文章のスコアは 0.977 でした。
つまり、本文を全く理解せずに 4 文字の並びを数えることは、人間の文章と、同じ文章を読んだことのない著者によるクロードの書き直しを区別することになるのが約 93% であるということです。それがうまくいくとは予想していませんでした

良い。
実際に何を拾っているのでしょうか？
モデルは非常にシンプルなので、それが何をしているのかを確認できます (まあ、とにかくクロードにやらせることができます)。
ここでは、単純なリライトの各側の最も強力な機能を 3 文字で示しているため、読みやすいように短くなります。
言葉ではなく断片ですが、読むことはできます。クロード側: -ing 語尾、-ly 副詞 ( ely 、 tly 、 ctl は「完全に」、「直接」、「正確に」を意味します)、「that」、「though」、「through」、および Ive (アポストロフィを取り除いた短縮形「I've」)。
人間の側: 「どれ」、「違う」、「しかし」、「だろう」、「だった」、 -tion と -ation の名詞語尾 ( ion 、 ati 、 tio )、 「of the」 と 「of a」 ( fth 、 oft 、 ofa )、そして「私は」と「私は」を完全に書き出す Iam と Iwi 。
AI の各文章を書き換え元となった特定の人間の文章と比較すると、個々の文字を観察することはさらに興味深いものになります。ペア内の違いは、主題や時代ではなく、書き方の違いです。
クロードによれば、そのパターンは認識できるという。 i 、 c 、 f 、 s 、 m は、英語がラテン語とフランス語から借用した単語の文字です:「考慮」、「重要性」、「特定」。 g 、 h 、 w 、 k 、 y は、古いゲルマン語の中核の文字です。「考える」、「しかし」、「働く」、「知っている」、「高い」です。クロードは語彙を次から次へ移動させています。
単語レベルでチェックするとそれが確認されます。ラテン語の語尾を持つ単語は両方に含まれ、プレーンな書き換えでは平均単語の長さはほとんど変化しないため、変化は単語の長さではなく、クロードがどの単語を選択したかです。
「明瞭で自然な英語」「平易で直接的な現代英語」ということなのでしょう。
3 番目の命令は、「正式で正確なレジスタ」を要求するものです。
すべてのレット

えー、立場が入れ替わりました。単語レベルでは、これは微妙ではありません。ラテン語の語尾が上がり、平均単語の長さも長くなります。
私たちは検出器を 2 回作りました。 3 つの命令すべて (上記のバージョン) に 1 回、プレーンおよび直接の書き換えのみを 1 回実行し、正式なものは省略します。他はすべて同一です。
正式な書き換えを一度も見たことがない検出器は、それらを検出できません。スコアは 0.153 で、実際に人間が書いた文字よりも人間らしく見えると確実に評価されます。混乱することも、壊れることもありません。 「平易になった文章はAIである」と学習し、より堅苦しくなった文章を渡され、それに応じて答えた。
新鮮な書き込みは 0.881 から 0.977 になり、どちらのトレーニング セットにも含まれません。クロードがオリジナルをたどることなく、ゼロから書いた文章。 3 番目の書き方を追加すると、モデルに 3 番目の書き方を教えるだけでなく、これまでに見たことのない 4 番目の条件に一般化されました。
同じことがコーパスの外側でも起こり、RAID は 0.794 から 0.889 になります。
さらにスタイルを追加すると、検出器をさらに汎用的にするのに役立つでしょう。
残念ながら、これは、トレーニングされていないスタイルを使用して検出器を騙すことができることも示しています。
すべてはブラウザ内で実行され、テキストがマシンから離れることはありません。
ちょっと狂ったプロジェクト、有益/教育的なビデオ、そして一般的に興味深いもののコレクション。
周囲のプロジェクトの構築

[切り捨てられた]

## Original Extract

Anthropic now watermarks Claude, but I wanted a detector we can run ourselves—so I built one. It ignores words entirely, counts 4‑letter n‑grams, feeds them to a linear SVM, and reliably flags Clau...

🌈 ESP32-S3 Rainbow: ZX Spectrum Emulator Board!
Get it on Crowd Supply →
Detecting Claude by counting letters
« Don't blame Claude - It's me, I'm the problem, it's me...
HELP SUPPORT MY WORK: If you're feeling flush then please stop by Patreon Or you can make a one off donation via ko-fi
On 2 August 2026 Anthropic began watermarking the text that new Claude models produce - their support pages cover which products are marked .
The watermark is not a hidden character or a bit of metadata. It is statistical. A secret key nudges the generated output in a way that isn’t obvious to a person, but that can be detected if you have the key. This means that Anthropic can detect it but we common folk cannot.
Other providers are doing similar things , or will be soon.
Can we build our own detector? Given a piece of text, can you tell that a model wrote it?
Let’s simplify things and just try one narrow target: text written by Claude Opus 5. Mostly because that’s what I have a subscription to…
And - spoiler alert - it turns out you can. I’ve deployed a web version of it here . Generate some text in Claude using Opus 5 and see if it can detect it.
And yes, if you feed this blog post in, despite me editing it substantially - it’s comes out as written by Claude.
A long time ago I did a PhD and, as part of that, I looked into questioned document examinations - a field that tries to work out who wrote a document. Did Shakespeare write a play, or was it someone else? Which of the Federalists wrote which paper? Who sent these threatening letters?
There’s a really nice simple technique that seems to work very well. You can often identify an author by counting short runs of characters. Ignore the words, sentence structure etc… just the occurrences of runs of letters. These runs are called n-grams.
An n-gram is simply a run of n items taken from a sequence. You can build them from words or from characters. Here I’m using characters.
Take “the quick brown fox”. Strip the spaces, the punctuation and the capitals and you have thequickbrownfox . Slide a four-character window along it, one position at a time, and record what you see:
theq hequ equi quic uick ickb ckbr kbro brow rown ownf wnfo nfox
Do that to a whole passage and you end up with a table of counts: that appeared 40 times, ting appeared 22 times, quic appeared once, and so on for every run that turned up. Throw the passage away. The table of counts is all the detector ever sees.
The counts are normalised so, in theory, the length of the text doesn’t matter. Obviously there is a limit to this - if your text is too short then the counts are not statistically valid and you won’t be able to predict anything.
The idea is that the statistical distribution is not something you are aware of as a writer and it’s quite hard to obfuscate. It exposes unconscious tells.
As with any problem of this type, the hard part is getting data, the actual “science” part is easy.
I needed human writing that was definitely human. The safest definition of “definitely” is “written before ChatGPT existed”, so everything I used was published before 2021. Claude helped me pull it together from three places:
Books from Project Gutenberg , all pre-1929.
Blog posts from the Blog Authorship Corpus , written in 2004 or earlier. Non-commercial research use only.
News articles from CC-News , published between 2016 and 2019.
That gave me 519 passages by 26 authors: 180 from books, 179 from news, 160 from blogs. Three different kinds of writing, and a spread of dates from Victorian to nearly modern.
The obvious way to get AI text is to ask Claude to write something.
So I asked it to rewrite the human text. For every human passage I asked Claude Opus 5 to produce the same passage again, preserving the meaning, the level of detail and roughly the length. That gives me matched pairs of content. Within a pair the subject is identical, the facts are identical and the length is close, so the only thing left to separate them is the writing itself.
Writing style is not one thing, so I asked three different ways:
plain - “Rewrite the following passage in clear natural English.”
direct - “Rewrite the following passage in plain, direct modern English.”
formal - “Rewrite the following passage in a formal, precise register.”
That comes to 1,557 rewrites, generated through the batch API. I also had Claude write fresh passages from scratch on the same topics, which I kept out of all the training and used later as a test.
First, how the accuracy is measured. The number below is ROC AUC . Take one human passage and one AI passage at random, score both, and ask whether the AI one got the higher score. ROC AUC is the proportion of random pairs where it did.
0.5 is a coin toss. 1.0 means it never gets a pair the wrong way round. It measures ordering, not whether any particular score is correct.
Second, the split. Before training anything I set four of the 26 authors aside and did not touch them. Everything else was measured by cross-validation, where the authors used for testing are never the authors used for training. Splitting by author rather than by passage is important. Split by passage and the classifier learns to recognise individual writers, which it is very good at, but it’s not what we want!
Then I swept the n-gram length from 1 to 8, holding everything else fixed:
Counting individual letters works! 28 features, just the letter frequencies, reach 0.651. That is not much, but it is well clear of a coin toss, from a really simple feature set.
Accuracy climbs steeply to 4, peaks at 5, and then falls away. Longer runs carry more information, right up until there is too little data to support them. This may hint that using word counts would work quite well too - but maybe that would require more data? Something to research…
I also tried spaced n-grams, where you take every second or every third character instead of consecutive ones. This didn’t work at all - it made the accuracy much worse.
I ended up picking 4 instead of 5
The peak is at 5. But I decided to use 4, and there are a couple of reasons for this:
The first reason is that going to 5 letters gives a marginal improvement in accuracy at the cost of a lot more features.
The second is that it only gives an improvement on long passages of text. The reason is how thinly the data spreads. There are 26⁴ = 457,000 possible 4-grams and 26⁵ = 11.9 million possible 5-grams.
If you take a typical passage from my corpus, with a length of 2,259 letters:
At 5 characters, nearly nine in ten of the runs in a passage appear only once and it gets worse as the text gets shorter: at 200 letters, 90% of 3-grams and 97% of 5-grams appear only once.
Here is the same detector built at each length, tested on 320 passages Claude wrote from prompts in a public benchmark against their 320 human source documents, with the text truncated to various lengths:
The 4-grams win at every length.
The finished thing is not complicated:
Take the text. Strip everything that is not a letter (including any spaces), and lowercase what remains.
Count all the 4-grams. Each passage becomes a list of 60,211 numbers.
Normalise, so the numbers are proportions and independent of passage length.
Feed it to a linear support vector machine, which learns which counts run higher in Claude’s writing.
Convert the result to a number between 0 and 1, where higher means more like Claude.
The n-gram counting and the classifier are both stock scikit-learn components.
Because of step 1 it cannot see punctuation, sentence length, paragraph structure, capitalisation or any individual word. Because of step 3 it cannot see how long the passage is. It sees letter proportions and nothing else.
The four authors set aside at the beginning were scored once, at the end:
The two numbers are close, which is the thing I most wanted to see. It means the classifier was not quietly learning to recognise particular authors.
Then I pointed it at text with no connection to the training data at all. First, 175 posts from my own blog written between 2006 and 2022, against eight passages written by Claude:
Second, a harder test. RAID is a public benchmark for machine-generated text detection, containing six million passages from 11 generators across eight subject domains. I took 320 of its prompts, had Claude Opus 5 answer them, and scored those answers against the 320 human source documents the prompts came from. Different corpus, and a different task, because these were written from scratch rather than rewritten. It scores 0.889.
And the fresh passages I held out of training entirely, written from scratch on the four unseen authors’ topics, score 0.977.
So: counting four-letter sequences, with no understanding of the text whatsoever, separates human writing from Claude’s rewriting of that same writing about 93% of the time on authors it has never read. I did not expect it to work that well.
What is it actually picking up?
The model is simple enough that we can see what it’s doing (well, I can get Claude to do it anyway).
Here are the strongest features on each side for the plain rewrites, at 3 characters so they are short enough to read:
They are fragments, not words, but you can read them. On the Claude side: -ing endings, -ly adverbs ( ely , tly , ctl for “entirely”, “directly”, “exactly”), “that”, “though” and “through”, and Ive , the contraction “I’ve” with the apostrophe stripped out.
On the human side: “which”, “not”, “but”, “will”, “was”, the -tion and -ation noun endings ( ion , ati , tio ), “of the” and “of a” ( fth , oft , ofa ), and Iam and Iwi , which are “I am” and “I will” written out in full.
Looking at individual letters is even more interesting when we compare each AI passage against the specific human passage it was rewritten from. Any difference inside a pair is a difference in writing, not in subject or era.
According to Claude, the pattern is recognisable. i , c , f , s and m are the letters of words English borrowed from Latin and French: “consideration”, “significance”, “specific”. g , h , w , k and y are the letters of the older Germanic core: “think”, “though”, “work”, “know”, “high”. Claude is moving the vocabulary from one to the other.
Checking at the word level confirms it. Words with Latin endings drop in both, and in the plain rewrites the mean word length barely moves, so the shift is which words Claude chose rather than how long they were.
I guess this is what: “clear natural English”, “plain, direct modern English” means.
Now the third instruction, the one asking for “a formal, precise register”:
Every letter has swapped sides. At the word level it is not subtle: Latin endings go up and mean word length also goes up.
We built the detector twice. Once on all three instructions, which is the version described above, and once on the plain and direct rewrites only, leaving the formal ones out. Everything else is identical.
A detector that never saw the formal rewrites cannot detect them. It scores 0.153, which means it reliably rates them as more human-looking than the actual human writing beside them . It is not confused and it is not broken. It learned “text that got plainer is AI”, it was handed text that got more formal, and it answered accordingly.
Fresh writing goes from 0.881 to 0.977, and it is in neither training set. Passages Claude wrote from scratch, with no original to follow. Adding a third writing style did not just teach the model a third writing style - it generalised to a fourth condition it had never seen.
The same happens outside the corpus , where RAID goes from 0.794 to 0.889.
Adding more styles would probably help to make the detector even more general.
Unfortunately it also indicates that you can cheat the detector by using a style it has not bene trained on.
The whole thing runs in a browser , and the text never leaves your machine.
A collection of slightly mad projects, instructive/educational videos , and generally interesting stuff.
Building projects aroun

[truncated]
