---
source: "https://blog.joellehman.com/identifying-life-changing-books-with-llms.html"
hn_url: "https://news.ycombinator.com/item?id=48520231"
title: "Identifying Life-Changing Books with LLMs"
article_title: "Identifying Life-Changing Books with LLMs"
author: "andsoitis"
captured_at: "2026-06-13T21:33:09Z"
capture_tool: "hn-digest"
hn_id: 48520231
score: 4
comments: 0
posted_at: "2026-06-13T18:48:14Z"
tags:
  - hacker-news
  - translated
---

# Identifying Life-Changing Books with LLMs

- HN: [48520231](https://news.ycombinator.com/item?id=48520231)
- Source: [blog.joellehman.com](https://blog.joellehman.com/identifying-life-changing-books-with-llms.html)
- Score: 4
- Comments: 0
- Posted: 2026-06-13T18:48:14Z

## Translation

タイトル: LLM を使用して人生を変える本を特定する

記事本文:
古いブラウザを使用しています。エクスペリエンスを向上させるには、ブラウザをアップグレードするか、Google Chrome Frame をアクティブにしてください。
LLM を使用して人生を変える本を特定する
TLDR: 言語モデルは何百万もの書評を分析し、読者が「人生を変えた」と言う可能性が最も高い本を特定できます。最も人生を変えた300冊の生成されたリストは、こちらでご覧いただけます。
メインイベント: 人生を変える本のテーブル
人生を変える本は最も読まれた本や最高の評価を得た本ではない
人生を変える本と最も読まれた本の関係
余談: カルビンとホッブスは最高です
付録: 人生を変える本リストのガイド付きツアー
パート I (1-25): 自己啓発 / スピリチュアル / 宗教
第 2 フェーズ (25-150): さらなる多様性
最終段階(150-300): 自伝、ヤングアダルト、小説
付録: 技術的な詳細
LLM 埋め込み
人生を変えるスコアの計算
カテゴリと説明の自動生成
これは間違いなく、私が今まで読んだ中で最も美しく、人生を変えるものです。読んでみてください。今すぐ。
-Tiny Beautiful Thingsのレビュー
驚くべきことに、20 ドルで、あなたの人生の流れを変えるアーカイブされた知恵を購入できます。新しい人生哲学を提供したり、感情的な経験を打ち破ったり、世界の新しい見方を提供したりすることができます。しかし、私たちが本と出会うのは無計画です。たとえば、5 つ星のレビューをすべて同じものとしてカウントするアルゴリズム (それが命の恩人であるか、単にページをめくるだけであるかにかかわらず) によって行われます。
このプロジェクトは、あなたの人生を変える可能性のある本をデータに基づいて特定することを目的としています。重要なデータは、数百万件の書籍レビューを含む GoodReads のスクレイピングから得られます (2017 年まで。このデータセットの作成については Mengting Wan に感謝します)。
テキストレビューには、読者の経験（読者がどのように感じたか、人生にどのような影響を与えたか）に関する詳細が豊富に含まれています。言語モデル (LLM) を使用すると、これらのレビューを大規模に分析して調査することができます。

定量的な「人生の変化」。このプロジェクトでは、24,000 冊を超える書籍の 600 万件以上のレビューから、「人生を変える」書籍の上位 300 冊を特定します。
アプローチ: レビューを埋め込む (文章をベクトルに変換する) ために (LLM) を適用し、「人生を変える」レビュー (読者がその本で人生が変わったと言う) が最も頻繁にある本を確認します。
埋め込みモデル (OpenAI のモデルを使用しました) は文章を数値のリストに変更し、私がトレーニングした分類器はそれらの数値を、その文章が人生を変える感情 (「この本は私の人生を変えました」など) を表しているかどうかにマッピングします。その結果、指定された書籍のすべてのテキスト レビューを通過し、それらのレビューの何パーセントが「人生を変える」ものであったかを取得できるパイプラインが作成されます。次に、すべての本についてそれを行い、どの本が最も高い割合を占めているか (最も人生を変える本) を確認します。
技術的な詳細については、この投稿の最後に記載されています。
メインイベント: 人生を変える本のテーブル
以下は、300 冊の本自体を「人生を変える」スコアの順に並べた表です。興味があれば、（人生を変えるほとんどの著者と同様に）詳細な分析が表の後に続きます。
(列をクリックして) 並べ替えたり、(著者など) 列の非表示を解除したり、(検索バーに入力して) 検索したり、(ジャンル フィルターから) 興味のあるジャンルを選択したりできます。
ヒント: ジャンル フィルターをクリックして、健康、経済、生産性、哲学に関する最も人生を変える小説や本を探します。
TLDR: 結果は理にかなっています。多くの自己啓発本、インスピレーションを与える人物の自伝、実存的および哲学的なテーマを扱った小説。
ご想像のとおり、人生を変える本は、人生の主な葛藤を中心にしています。最大のカテゴリは、精神性、心理学、哲学、人間関係、健康、大きな問題に取り組む小説、インの自伝などです。

海賊版のフィギュア。
全体として、自助が主なテーマであり、特にリストの上位付近にあります。たとえば、最高得点の本は 1976 年の自己啓発本『Your Erroneous Zones』で、NYT のベストセラー リストに 64 週間掲載されていました (ただし、私はこれから読む予定になるまでその名前を聞いたことがありませんでした)。
さらなる検証は、多くの書籍のレビューを抜き打ちチェックすることで得られました。これにより、分類器がかなり信頼できることが明らかになり、結果が確実であると確信するのに十分でした。セクションの最初に置いたレビューは、これらのスポットチェックから得たものです (そして、すべてが分類子によって「人生を変える」と評価されています)。
最後に、私は個人的に、ここに挙げた本の多くが人生に影響を与えるものであると感じました。順不同で、私のお気に入りは次のとおりです。139 位の『Siddhartha』、6 位の『Loving What Is』、207 位のハリル・ジブランの『The Prophet』、43 位の『Man's Search for Meaning』、86 位のペマ・チョドロン著『When Things Fall Apart』、262 位の『The Unbearable Lightness of Being』、および 140 位のシェリル・ストレイドの『Tiny Beautiful Things』です。
このリストを見ると、178 位の『Jonathan Living Seagull』、89 位のベル・フックの『All About Love』、C.S. ルイスの『Screwtape Letters』などの本を読みたいという私の欲求も高まりました。
人生を変える本。他に何をすればいいのか分からないほど落ち込んでいたときにこれを読みました。そして驚いたことに、この本がとても役に立ちました。
- ミッチ・アルボムによる『モリーとの火曜日』のレビュー
3 冊以上の本をリストに載せることができた著者は 4 人だけでした。そして、公式に最も人生を変えた作家（だと私は思っている）ミッチ・アルボム、4作品目『モリーとの火曜日』、『少しだけ信仰を』、『天国で出会う5人』、そして『タイム・キーパー』でおめでとうございます。ミッチ、ジョン・グリーン、ブレン・ブラウンにとって、アイン・ランドは奇妙な仲間に見えるが、結果的にはそういうことだった。
そして、リスト内の 2 冊の本と同点となった次点の方々にもおめでとうございます (落ちた人に大声で叫びます)

バークレー在住のマイケル・ポーラン氏）：
人生を変える本は最も読まれた本や最高の評価を得た本ではない
残念ながらスケールは 5 までしかありません。これは 6 です。人生が変わる（私にとって）
読者への影響という点では、すべての 5 つ星の本が同じように作成されているわけではありません。
「人生を変える」レビューを誘発する可能性が最も高い本と、読者が 1 つ星から 5 つ星のスケールで最高評価を付けた本の間には、ほとんど重複がありません。 (ノイズを減らすために) 総合的な星評価が 1000 以上の書籍のみを含めたことに注意してください。
人生を変える本トップ 300 のうち、平均 1 ～ 5 つ星の評価が最も高い上位 300 冊に入っているのは、次の 1 冊だけです。ただし、GoodReads での評価は 2017 年以降下がっています (これらの統計が計算されたとき、おそらく最近爆発して、より多くの視聴者がそれにさらされているためです)。
人生を変える本と最も読まれた本の関係
人生を変えるような本の最初の 100 冊には、最も読まれた本との重複はほとんどありません (GoodReads での評価の数によって測定されます)。最も読まれた 300 冊の本には、『非常に効果的な人の 7 つの習慣』と『モリーとの火曜日』だけが含まれています。
人生を変えるリストには、学校で読書として割り当てられているものも含め、多くの影響力のある古典が含まれているため、最も読まれているものと人生を変えるものの間にはより多くの重複があります。 『アルジャーノンに花束を』、『ウォールフラワーであることの特典』、『すばらしい新世界』、『1984』などの本。
余談: カルビンとホッブスは最高です
GoodReads で最も平均的な星評価が高い本は The Complete Calvin and Hobbes で、アマゾンでも驚くほど高い評価を得ています (95% が 5 つ星のレビュー)。私はカルビンとホッブスが大好きだったので、これは嬉しい発見でした。
合理的な仮説は、熱心なファンだけがこの製品を購入する可能性が高いということです。

カルビンとホッブスのコミック本が揃っているため、（C&H がすでに優れていることに加えて）高評価を自動的に選択します。この仮説を裏付けるものとして、『カルビンとホッブズ』の次にトップスターと評価される本はワン・ダイレクションに関するものであり、その次は『王の道』というファンタジー本の続編である。
ワン・ダイレクションのファンだけがワン・ダイレクションに関する本を購入する可能性が高いという考えです（自主選択）。同様に、『王の道』の最初の本はすでに非常に高い評価を受けていますが、続けて 2 冊目を読んだ場合は、最初の本を楽しんだ可能性が非常に高く、否定的なレビューはさらに減ります (再び自己選択)。
私は本が大好きです。そして、本が私たちの生活を変える可能性を秘めていることが大好きです。皆さんがこれらの本を読むきっかけになれば幸いです (アフィリエイトリンクを通じて小銭が得られるからというだけではありません)。これは、より大きなプロジェクト (いくつかのトランスフォーマーのトレーニングを含む) の始まりであり、うまくいけば、最終的には、(このリストに示されている乱暴な組み合わせではなく、多くが感動しないであろう) あなたの人生を変える、よりパーソナライズされたおすすめの本を提供することになるでしょう。
twitter または substack で私をフォローして最新情報を入手するか、質問やコメント (またはさらなる分析のための提案 - 他にどのようなカテゴリの書籍が登場したら興味深いですか?) をお気軽にメールしてください。
もう 1 つレビューして終わります。
私がこの本を見つけて読んだのは神の介入でした。私は急いでトレーラーに荷物を満載して家を出たところでした。私は悪い場所にいた。仕事を失いました。私の結婚生活は大惨事でした。私は12歳の息子と一緒に両親と一緒に住むことになりました。私はとても傷つき、よくトイレに行って泣いていました。壊れた。
それを見た母は、家から出ようと半額ブックスに行こうと提案しました。本を買うお金がなかった。本当に何もなかった

何でも読みたいという欲求。店で棚を見ていたら、この本の背表紙を見つけました。 「なんてひどいことだろう」と、開ける前に思いました。しかし、ストアの最初のセクションを読んだとき、気分が良くなりました。
クレジットカードで購入しました。今、あの困難な時期をどうやって乗り越えたかと聞かれたら、正直に「この本が役に立った」と答えることができます。その日私は本を読みました。それから何度も読み返しました。
誰が言ったのか知りませんが、本当です。自分が非常に暗い場所にいることに気づいたときは、蝶が繭の中で行うように、この時間を利用して自分自身を作り直してください。そして、外に出たとき、あなたは何か違うもの、より良いものになるでしょう。この本は大きな転機となっただけでなく、命を救いました。これを100点の太字で書けるならそうします。
-Pema Chodron著『When Things Fall Apart』のレビュー
このプロジェクトなしでは存在できなかったデータセットを作成してくれた Mengtian Wan に感謝します。有益な議論をしてくれた Eric Frank と Ivan Vendrov に感謝します。また、この投稿の最初の草稿について洞察力に富んだコメントとフィードバックをくれた Rosanne Liu に感謝します。また、人生を変える本やGoodReadsの秘教について、私の話を熱心に聞いてくれた忍耐強い友人たちにも感謝します。
付録: 人生を変える本リストのガイド付きツアー
ここでは、リストがどのように進行するかについての一般的な雰囲気を示します。興味深いことに、読者が人生が変わったと感じる可能性が最も高いのはスピリチュアルな本のようです（リストの最初の部分をほぼ独占しています）。私自身にとって、瞑想は人生を変えるものでしたが、最初は本を通じて瞑想に行き着きました。
それから一般的な自己啓発本もたくさんあります。しかし、トップ 300 リストの中ほどから終わりにかけては、小説、ヤングアダルト作品、自伝がさらに多くなります。
パート I (1-25): 自己啓発 / スピリチュアル / 宗教
最初の 2 冊の最高得点の本は読んでいなかった

以前に聞いたことがある。 『Your Erroneous Zones』は 1976 年の心理的自己啓発本で、『Psycho-Cyber​​netics』は 1960 年の別の心理的自己啓発本です。モルモン書は 4 位に登場し、キリスト教をテーマにしたいくつかの本も初期に登場しています。
多くの瞑想/スピリチュアリティの本: 6 位の『Loving What Is Is』は、単純な質問ベースの瞑想形式を提唱していると見なされます。 15位の『Be Here Now』は、ハーバード大学教授からサイケデリックの第一人者となったラム・ダスによる、東洋のスピリチュアリティを西洋に紹介する重要な本だった。そして第19位は、より穏やかな仏教瞑想の道を歩んだタラ・ブラッハによる『Radical Acceptance』だ。
その他のトップスコアの本は、3 位の「非暴力コミュニケーション」、7 位の「The Road Les Travelled」、10 位のトニー・ロビンスの本など、私のような自己啓発マニアには馴染みのあるものかもしれません。
第 2 フェーズ (25-150): さらなる多様性
25冊になるまでは、ほとんどがスピリチュアル/心理学的なものです。次に、最初の健康本『Eat to Live』が出版され、2003 年に出版されました。この本には、食事制限のプランが記載されています (ほぼすべての健康本は食品に関連しています)。
そのすぐ後に、イシュマエルと呼ばれる最初の小説が出版されます (私は読んで面白かったのですが、詳細は思い出せません)。
最初の自伝は 63 歳頃で、聖人のような若い C について書かれています。

[切り捨てられた]

## Original Extract

You are using an outdated browser. Please upgrade your browser or activate Google Chrome Frame to improve your experience.
Identifying Life-Changing Books with LLMs
TLDR: Language models can analyze millions of book reviews, to identify books most likely for readers to say it "changed their life." See the generated list of the 300 most life-changing books here .
Main Event: Table of Life-Changing Books
Life-Changing Books Are Not the Most-Read or Top-Rated
Relationship between Life-changing and Most-Read Books
Aside: Calvin and Hobbes is the Best
Appendix: Guided Tour of the Life-Changing Books List
Part I (1-25): Self-help / Spiritual / Religion
Second phase (25-150): More diversity
Last phase(150-300): Autobiographies, young adult, and novels
Appendix: Technical Details
LLM Embeddings
Calculating Life-Changing Score
Autogeneration of categories and descriptions
This is hands down the most beautiful and life-changing thing I've ever read. READ IT. Right now.
-review of Tiny Beautiful Things
Amazingly, $20 can buy archived wisdom that alters the course of your life: providing a new life-philosophy, shattering emotional experiences, or a new way of seeing the world. But we encounter books haphazardly -- for example, through algorithms that count all five-star reviews as equal (whether they are life-savers or just page-turners).
This project aims at a data-driven way to identify books that can change your life. Key data comes from a scrape of GoodReads which includes millions of written reviews of books (up through 2017; thanks to Mengting Wan for creating this dataset).
Text reviews are rich with details about readers' experiences (how it made them feel; how it impacted their life). With language models (LLMs) we can analyze these reviews at scale to explore "life-changingness" quantitatively. From over six million reviews of over 24,000 books, this project identifies the top 300 "life-changing" ones.
The approach: apply (LLMs) to embed reviews (changing sentences into vectors), to see what books most frequently have "life-changing" reviews (where the reader says the book changed their life).
The embedding model (I used one from OpenAI) changes sentences into lists of numbers, and a classifier I trained maps those numbers to whether the sentence represents a life-changing sentiment (like "This book changed my life."). The result is a pipeline where you can pass through all text reviews of a gven book, and get out what percentage of those reviews were 'life-changing.' Then you can do that for all of the books, and see which ones have the highest percentage (the most life-changing books).
More technical details are at the post's end .
Main Event: Table of Life-Changing Books
Here's the table of the 300 books itself, in order of their "life-changing"ness scores. More analysis (like most life-changing authors ) follows the table, if you're interested.
You can sort (by clicking on columns), un-hide columns (like author), search (by entering into the search bar), or select a genre of interest (from the genre filter).
Tip: click on the genre filter to look for the most life-changing novels or books on health, finances, productivity, or philosophy.
TLDR: The results make sense; lots of self-help books, autobiographies of inspirational figures, and novels dealing with existential and philosophical themes.
As you might expect, the life-changing books center around the main struggles of life. The biggest categories are spirituality, psychology, philosophy, relationships, health, novels tackling big issues, and autobiographies of inspirational figures.
Overall, self-help is the dominant theme, especially near the top of the list. For example, the highest scoring book was Your Erroneous Zones , a self-help book from 1976 that was on the NYT bestseller list for 64 weeks (although I'd never heard of it before I'm planning on reading it).
Further validation came from spot-checking reviews for many of the books. This revealed that the classifier was pretty reliable, enough to convince me that the results are solid. The reviews I've placed at the beginning of sections came from these spot-checks (and all are rated as "life-changing" by the classifier).
Finally, I've personally found many of the listed books to be life-impacting. In no particular order, some of my favs are: Siddhartha at #139, Loving What Is at #6, Khalil Gibran's The Prophet at #207, Man's Search for Meaning at #43, When Things Fall Apart by Pema Chodron at #86, The Unbearable Lightness of Being at #262, and Tiny Beautiful Things by Cheryl Strayed at #140.
The list has also reinforced my desire to read books like Jonathan Living Seagull at #178, bell hook's All About Love at #89, and C.S. Lewis' Screwtape Letters .
A life-changing book. I read this when I was so down that I didn't know what else to do. And to my relief, this book helped me a lot.
-review of Tuesdays with Morrie by Mitch Albom
Only four authors managed to have three or more books on the list. And congrats to Mitch Albom, officially the most life-changing author (I suppose), with 4: Tuesdays with Morrie , Have a Little Faith , The Five People You Meet in Heaven , and The Time Keeper . Ayn Rand seems strange company for Mitch, John Green and Brene Brown, but that's how it turned out.
And congrats also to the runners-up that tied with two books on the list (shout out to fellow Berkeley resident Michael Pollan):
Life-Changing Books Are Not the Most-Read or Top-Rated
too bad the scale only goes to 5. This is a six. Life changing (for me)
Not all five-star books are created equal in terms of their impact on readers.
There is little overlap between the books most likely to induce a "life-changing" review with those that readers rate highest on a 1-5 star scale. Note that I only included books with at least 1000 overall star-ratings (to reduce noise).
Only one of the top-300 life-changing books also are among the top-300 books with the highest average 1-5 star rating: The Body Keeps the Score: Brain, Mind, and Body in the Healing of Trauma ; however, its rating on GoodReads has fallen since 2017 (when these stats were calculated, perhaps because it's recently blown up and a wider audience is being exposed to it).
Relationship between Life-changing and Most-Read Books
In the first 100 life-changing books, there is little overlap with the most-read books (as measured by how many ratings they have on GoodReads). Only The 7 Habits of Highly Effective People and Tuesdays with Morrie are among the 300 most-read books.
As the life-changing list goes on, it includes many impactful classics, including many that are assigned as reading in school, so there is more overlap between most-read and life-changing: E.g. books like Flowers for Algernon , The Perks of Being a Wallflower , Brave New World , and 1984 .
Aside: Calvin and Hobbes is the Best
The book with the highest average star-rating on GoodReads is The Complete Calvin and Hobbes , which also has astoundingly high ratings on amazon (95% 5-star reviews). I loved Calvin and Hobbes growing up, so this was a delightful discovery.
A reasonable hypothesis is that only devoted fans are likely to buy the complete set of Calvin and Hobbes comic books, so it self-selects for high ratings (in addition to C&H already being great). Supporting this hypothesis, the next top-star-rated book after Calvin and Hobbes is about the band One Direction, and next is the sequel to a fantasy book called The Way of Kings .
The idea is that only fans of One Direction are likely to buy books about them (self-selection). Similarly, while the first book in The Way of Kings is already very high-rated, if you go on to read the second book, it's very likely you enjoyed the first one, further reducing any negative reviews (self-selection again).
I love books -- and I love that books have the possibility to change our lives. I hope you end up being inspired to read one of these books (and not only because I'll get some spare change through affiliate links). This is the beginning of a larger project (involving training some transformers), that hopefully will end up giving you more personalized recommendations of books that change your life (rather than the wild mix presented in this list, of which many will not move you).
Follow me on twitter or substack to get updates, or feel free to email me with questions or comments (or suggestions for further analysis -- what other categories of books might be interesting to surface?).
I'll end with one more review:
It was divine intervention that I found & read this book. I had just hurriedly packed a trailer full of stuff & moved out of my house. I was in a bad place. I lost my job. My marriage was a huge disaster. I had to move in with my parents along with my son, 12. I was so wrecked, I often went into the bathroom to cry. Broken.
Seeing this, my mom suggested we go to Half Price Books to get out of the house. I had no money to buy a book. I really had no desire to read anything. At the store, I browsing thru the shelves, I saw this book spine. "What a load of crap," I thought before I opened it. But when I read the first section in the store, I felt better.
I bought it with a credit card. If you asked me now how I got thru that difficult time, I can honestly answer, "this book was instrumental." I read the book that day. Then I reread it over & over.
I don't know who said it, but it's true, when you find yourself in a very dark place, use this time to reshape yourself like a butterfly does in its cocoon. And when you come out, you will be something different, something better. This book was not only a HUGE turning point, it was life saving. If I could write this out this in 100-point bold type, I would.
-review of When Things Fall Apart by Pema Chodron
Thanks to Mengtian Wan for creating the dataset without which this project couldn't exist. Thanks to Eric Frank and Ivan Vendrov for helpful discussion, and to Rosanne Liu for insightful comments and feedback on an initial draft of this post. Also much thanks to my patient friends who listened to me drone on about life-changing books and the esoterica of GoodReads.
Appendix: Guided Tour of the Life-Changing Books List
Here I give the general flavor of how the list progresses; interestingly, it seems like spiritual books are the ones readers are most likely to feel changed their life (and semi-dominate the first part of the list). For myself, meditation has been life-changing, and I did come to it first through a book.
Then there is a lot of general self-help books. But towards the middle and end of the top 300 list, there are many more novels, young-adult works, and autobiographies.
Part I (1-25): Self-help / Spiritual / Religion
The first two top-scoring books, I hadn't heard of before; Your Erroneous Zones -- which is a psychological self-help book from 1976, and Psycho-Cybernetics , another psychological self-help book from 1960. The Book of Mormon makes an appearance at #4, and several Christian-themed books show up early as well.
Many meditation/spirituality books: Loving What Is at #6, can be seen as advocating for a simple question-based form of meditation; Be Here Now at #15 was a key book introducing Eastern spirituality to the West, by Harvard-professor-turned-psychedelic-guru Ram Dass. And at #19 is Radical Acceptance by Tara Brach, with her more-gentle Buddhist meditative path.
Other top-scoring books might be familiar to self-help junkies like myself, like Non-Violent Communication at #3, The Road Less Travelled at #7, and a Tony Robbins book at #10.
Second phase (25-150): More diversity
Until 25 books in, most are spiritual/psychological; then comes the first health book Eat to Live , a 2003 book with a restrictive diet plan (nearly all the health books relate to food).
The first novel comes shortly thereafter (one I read, and enjoyed, but can't recall the details of), called Ishmael .
The first autobiography hits around 63, about a saintly young C

[truncated]
