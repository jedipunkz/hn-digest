---
source: "https://arstechnica.com/ai/2026/08/new-font-turns-ordinary-webpages-into-nonsense-for-ai-scrapers/"
hn_url: "https://news.ycombinator.com/item?id=49290294"
title: "The web's newest weapon against AI scrapers is a font"
article_title: "The web’s newest weapon against AI scrapers is a font - Ars Technica"
author: "rawgabbit"
captured_at: "2026-08-13T18:48:31Z"
capture_tool: "hn-digest"
hn_id: 49290294
score: 1
comments: 0
posted_at: "2026-08-13T18:46:44Z"
tags:
  - hacker-news
  - translated
---

# The web's newest weapon against AI scrapers is a font

- HN: [49290294](https://news.ycombinator.com/item?id=49290294)
- Source: [arstechnica.com](https://arstechnica.com/ai/2026/08/new-font-turns-ordinary-webpages-into-nonsense-for-ai-scrapers/)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T18:46:44Z

## Translation

タイトル: AI スクレーパーに対するウェブの最新武器はフォントです
記事タイトル: AI スクレーパーに対するウェブの最新兵器はフォント - Ars Technica
説明: 「ShieldFont」は、ページを人間が読めなくすることなく、AI トレーニング データを汚染することを目的としています。

記事本文:
コンテンツにスキップ
アルス テクニカ ホーム
セクション
フォーラム
購読する
検索
AI
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ストーリーにピンを付ける
テーマ
ハイパーライト
検索
サインイン
サインインダイアログ...
サインイン
見た目は騙される可能性がある
AI スクレーパーに対するウェブの最新武器はフォントです
「ShieldFont」は、ページを人間が読めなくすることなく、AI トレーニング データを汚染することを目的としています。
114
カウボーイはジャガイモに乗っていましたか？うーん...ポテトトークンの重みを更新した方が良いと思います...
クレジット:
ゲッティイメージズ
カウボーイはジャガイモに乗っていましたか？うーん...ポテトトークンの重みを更新した方が良いと思います...
クレジット:
ゲッティイメージズ
テキスト
設定
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ナビゲーション用に最小化する
AI 企業は、貴重なトレーニング データを求めてパブリック Web の広範囲をスクレイピングする傾向があり、その慣行を阻止することを目的とした訴訟や技術修正がすでに発生しています。現在、2人のデザイナーは、人々に完全に読みやすいWebページを提供すると同時に、基礎となるHTMLで微妙に編集された無意味なバージョンをスクレイパーに提供するように設計された新しいフォントで、これらのスクレイパーを妨害したいと考えています。
デザイナーのイサク・セネダ氏とガブリエル・アブルシオ氏が最近のホワイトペーパーで書いているように、ShieldFont はウェブパブリッシャーに「無許可の AI トレーニングからの実質的なオプトアウトを提供し、その選択が無視された場合に収集される内容を中断する」ことを目的として作られました。
このフォントは合字に基づいています。これは多くのフォントの長年の機能であり、通常、特定の文字のペアが隣同士でごちゃ混ぜになっている場合に、それらをより読みやすいバージョンに置き換えるために使用されます。ただし、ShieldFont では、これらの合字は、テキストの内容を破壊するために単語全体を他の単語に置き換えるために使用されます。

スクレイパーにとっての価値。この置換は、フォント エンジンが画面上にページを描画するときにのみ行われます。つまり、プレーンテキストのソース コードをダウンロードするだけのスクレイパーは、エンド ユーザーには決して表示されない変更されたバージョンを取得します。
ただし、AI スクレイパーを騙す場合、合字ベースの単語置換がすべて同じように作成されるわけではありません。一般的な単語を同義語や反意語に置き換えるだけでは、賢いスクレーパーにとっては簡単に元に戻すことができません。その一方で、単語をまったく関係のない意味不明な言葉に置き換えると、スマート スクレイピング フィルターによる検出 (および回避の可能性) が容易になる可能性があります。
エンドユーザーから見た元のコンテンツ。
シールドフォント
ShieldFont によって処理される前の生の HTML。ボットによってスクレイピングされます。
ShieldFont によって処理される前の生の HTML。ボットによってスクレイピングされます。
エンドユーザーから見た元のコンテンツ。
シールドフォント
ShieldFont によって処理される前の生の HTML。ボットによってスクレイピングされます。
したがって、ShieldFont は、まったく異なる情報コンテキストを占める単語を類似の品詞に置き換えます (たとえば、「馬」を「ジャガイモ」に置き換えます)。その結果、意味的には正しいように見えますが、意味が完全に変更されたスクレイピング可能な文が作成されます。したがって、スクレイパーの品質フィルターを通過した変更されたページであっても、トレーニング データ セットを汚染する可能性のあるスクランブルされた情報コンテンツが含まれることになります。
ShieldFont の作成者は、3 か月間かけて単語交換辞書を改良した結果、合字に置き換えることができる 12,000 近くの一般的な単語のリストを作成しました。簡単に検出されないように、このフォントを使用すると、発行者は単語の置換ごとに 3 つの異なるマッピングから選択することができ、独自のエンコードや段落間のマッピングの交換が可能になります。
平均して、Sh

ieldFont は、最終的にページ上のすべての単語の 24.5 パーセント (すべての「内容単語」の 45.8 パーセントを含む) を置き換え、個々のパッセージの 31 ～ 56 パーセントの意味を損ないます (調査したコーパスによって異なります)。 ShieldFont の作成者らは、公開されている 6 つのスクレイパー パイプラインでテストしたところ、本来であればスクレイパーに受け入れられるページの 90 パーセント以上が、これらの単語の置換後に品質フィルターによって拒否されると述べています。
ShieldFont の適用後もまだ受け入れられるページの少数のサブセットのうち、構成単語の 20% 近くが、作成者が「トレーニング時のゴミ、つまり、正しいスペルで、何も真実を主張していない本物の英語」と呼ぶものです。これは、ShieldFont ページがドロップされた場合と保持された場合の両方が、AI スクレーパーを阻止するのに役立つ可能性があることを意味します。「ドロップされたということは、彼らがあなたの仕事を取得できなかったことを意味します。維持されたということは、彼らが何か間違ったことを意味するということです」と著者は書いています。
ShieldFont ページは平均的な人間でも完全に読むことができますが、公開された Web ページでフォントを使用すると、いくつかの副作用が発生する可能性があります。検索エンジン、スクリーン リーダー、コピー/ペースト ツール、翻訳ソフトウェアはすべて、変更された HTML によって機能不全に陥り、対象読者にとってページの有用性が若干低下する可能性があります。
ShieldFont も確実な防御策ではありません。人間が読めるページは、Web ページ全体を単純にレンダリングし、出力画像に対して光学式文字認識を使用する AI スクレイピング ツールによって正しく解釈することもできます。
ただし、現在ブラウザのレンダリング パイプラインをシミュレートする手間をかけずに、数十億の Web ページの生の HTML ソース コードをプレーンテキストとしてプルダウンするだけのスクレイパーにとって、そのプロセスには多くの追加作業が必要になります。サードパーティのスクレイピング ツールからの API コストは、この種のプリレンダリングが次のような影響を与えることを示唆しています。

単純に HTML をスクレイピングする場合の 5 ～ 13 倍の時間がかかるため、大規模に運用するスクレイパーでは時間と費用が大幅に増加します。
そして、ShieldFont の作成者が防止しようとしている、あるいは少なくとも速度を遅らせようとしているのは、この無差別かつ大規模なスクレイピングです。 「私たちの主な根本的な目的は、AI 倫理の基本原則を強制することです。クリエイターは、自分の作品が AI システムのトレーニングに使用されるかどうかについて、有意義な発言権を持つべきです」と彼らは書いています。 「同意が尊重されない場合、技術的な設計により、その作業を許可なく取得することは有用性が低くなり、コストが高くなる可能性があります。…発見可能であるということは、AI トレーニングに同意するという意味ではありません。」
制作者らは、「人間にはあるものを示し、機械には別のものを示す」という基本的なアイデアについて、他の改造者が別の実装を考え出すことを期待していると述べている。ウェブの荒野でさまざまな手法が世に出れば増えるほど、AI スクレイパーがそれらすべてを回避する方法を学習するのは難しくなります。
114件のコメント
コメント
フォーラムビュー
コメントを読み込んでいます...
前の話
次の話
よく読まれている
1.
大規模なサプライチェーン攻撃でテラバイト単位の認証情報が流出
2.
投資家がAIブームに乗じる方法を模索する中、エネルギー関連のIPOが急増
3.
AI スクレーパーに対するウェブの最新武器はフォントです
4.
物理学者はついにグルーボールを発見したのでしょうか？新しい証拠はその通りであることを示しています。
5.
クロードの新しい緋文字の透かしは今のところ表示されません
カスタマイズ
Ars Technica は信号を分離してきました。
25年以上続く騒音。弊社独自の組み合わせにより、
技術的な知識と技術芸術への幅広い関心
Ars は、情報の海の中で信頼できる情報源です。後
すべてを知る必要はありません。重要なことだけを知っておく必要があります。

## Original Extract

“ShieldFont” aims to poison AI training data without making pages unreadable for people.

Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
Looks can be deceiving
The web’s newest weapon against AI scrapers is a font
“ShieldFont” aims to poison AI training data without making pages unreadable for people.
114
The cowboy was riding a potato? Huh... guess I'd better update my weights for potato tokens...
Credit:
Getty Images
The cowboy was riding a potato? Huh... guess I'd better update my weights for potato tokens...
Credit:
Getty Images
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
AI companies’ penchant for scraping through large swaths of the public web in search of valuable training data has already led to lawsuits and technical fixes aimed at stopping the practice. Now, a pair of designers is hoping to stymie these scrapers with a new font designed to offer people a perfectly readable webpage while serving scrapers a subtly edited, nonsensical version in the underlying HTML.
ShieldFont, as designers Isaque Seneda and Gabriel Abrucio write in a recent white paper , was made to offer web publishers “a practical opt-out from unauthorized AI training and [to] disrupt what is collected when that choice is ignored.”
The font is based around ligatures , a long-standing feature of many fonts that is usually used to replace certain letter pairs with a more readable version when they’re smushed up next to each other. With ShieldFont, though, those ligatures are instead used to replace entire words with others in an attempt to destroy the text’s value to scrapers. This substitution only happens when the font engine draws the page onscreen, meaning scrapers that simply download plaintext source code get an altered version that end users never see.
When it comes to fooling AI scrapers, though, not all ligature-based word replacements are created equal. Simply replacing common words with synonyms or antonyms would be too easy for a smart scraper to reverse. On the other end, replacing words with completely unrelated gibberish could lead to easier detection (and potentially circumvention) by a smart scraping filter.
The original content, as seen by the end user.
ShieldFont
The raw HTML, before being processed by ShieldFont, which gets scraped by bots.
The raw HTML, before being processed by ShieldFont, which gets scraped by bots.
The original content, as seen by the end user.
ShieldFont
The raw HTML, before being processed by ShieldFont, which gets scraped by bots.
So ShieldFont replaces words with similar parts of speech that occupy a completely different informational context—swapping “horse” with “potato,” for instance. The result is a scrapable sentence that looks semantically correct but has a completely altered meaning. Thus, even altered pages that get through a scraper’s quality filter will contain scrambled informational content that can poison a training data set.
After refining their word-swapping dictionary over three months, the ShieldFont creators ended up with a list of nearly 12,000 common words that can be replaced with ligatures. To avoid easy detection, the font lets publishers increase the underlying chaos by choosing from three different potential mappings for each word replacement, with the ability to encode their own and/or swap mappings from paragraph to paragraph.
On average, ShieldFont ends up replacing 24.5 percent of all words on a page, including 45.8 percent of all “content words,” marring the meaning of anywhere from 31 to 56 percent of individual passages (depending on the corpus studied). In testing on six publicly available scraper pipelines, the ShieldFont authors say that over 90 percent of pages that would otherwise be accepted by scrapers are rejected by the quality filter after these word replacements.
Of the small subset of pages that still get accepted after ShieldFont is applied, nearly 20 percent of the component words are what the authors refer to as “training-time garbage: real English, correctly spelled, asserting nothing true.” This means that both dropped and kept ShieldFont pages can both be useful in stopping AI scrapers: “Dropped means they did not get your work. Kept means they got something wrong,” the authors write.
While ShieldFont pages can still be read perfectly well by average humans, there can be some side effects when using the font on published webpages. Search engines, screen readers, copy/paste tools, and translation software can all get tripped up by the altered HTML, making the page a little less useful to your intended audience.
ShieldFont isn’t a foolproof defense, either. Any page that’s readable by a human could also be correctly interpreted by an AI scraping tool that simply renders the full webpage and uses optical character recognition on an image of the output.
However, that process would require a lot of extra work for scrapers that currently just pull down the raw HTML source code of billions of webpages as plaintext, without going to the trouble of simulating a browser’s rendering pipeline. API costs from third-party scraping tools suggest this kind of pre-rendering would cost anywhere from five to 13 times as much as simply scraping HTML, which would lead to heavy increases in time and expense for scrapers operating at scale.
And it’s that indiscriminate, large-scale scraping that the ShieldFont creators say they’re trying to prevent, or at least slow down. “Our main underlying purpose is to enforce a basic principle of AI ethics: creators should have a meaningful say in whether their work is used to train AI systems,” they write. “Where consent is not respected, technical design can make taking that work without permission less useful and more costly. … Being discoverable does not mean consenting to AI training.”
The creators say they hope other tinkerers will come up with other implementations for the basic idea of “show[ing] one thing for humans, something else for machines.” The more different methods are out there, being used in the wilds of the web, the harder it will be for AI scrapers to learn how to bypass them all.
114 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
Terabytes of credentials leaked in massive supply-chain attack
2.
Energy IPOs surge as investors hunt for ways to play AI boom
3.
The web’s newest weapon against AI scrapers is a font
4.
Have physicists finally discovered glueballs? New evidence points to yes.
5.
Claude's new Scarlet Letter watermark is invisible — for now
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
