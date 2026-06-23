---
source: "https://simoncozens.github.io/state-of-ai-font-generation/"
hn_url: "https://news.ycombinator.com/item?id=48644306"
title: "The State of AI Font Generation"
article_title: "The State of AI Font Generation | Font and Text Technology"
author: "gsky"
captured_at: "2026-06-23T13:52:15Z"
capture_tool: "hn-digest"
hn_id: 48644306
score: 1
comments: 0
posted_at: "2026-06-23T12:59:13Z"
tags:
  - hacker-news
  - translated
---

# The State of AI Font Generation

- HN: [48644306](https://news.ycombinator.com/item?id=48644306)
- Source: [simoncozens.github.io](https://simoncozens.github.io/state-of-ai-font-generation/)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T12:59:13Z

## Translation

タイトル: AI フォント生成の現状
記事のタイトル: AI フォント生成の現状 |フォントとテキストのテクノロジー
説明: 2 つの免責事項: まず、この記事は完全に人間が作成したものです。執筆には AI は関与していません。次に、私は今年、グリフ生成モデルのトレーニングに多くの時間を費やしてきました。私がこのことについて率直に申し上げたいのは、特定の人たちにとって、これは私の信用を完全に失墜させることになると分かっているからです。
[切り捨てられた]

記事本文:
AI フォント生成の現状 |フォントとテキストのテクノロジー
simoncozens.github.io
AI フォント生成の現状
2 つの免責事項: まず、この記事は完全に人間が作成したものです。執筆には AI は関与していません。次に、私は今年、グリフ生成モデルのトレーニングに多くの時間を費やしてきました。私がこのことについて率直に申し上げたいのは、特定の人たちにとって、これが私の信用を完全に落とすことになることを知っているからです。そして、後で知って裏切られたと感じてほしくないのです。しかし、それなら、AI フォント生成の経験のない人が書いた AI フォント生成に関する記事を読みたいですか?
最近、機械学習モデルが新しいフォントを生成したり、(グリフセットやデザイン空間を拡張することで) 既存のフォントを「埋める」機能について多くの議論が行われてきましたが、その多くは Typedrawers のこのスレッドによって引き起こされました。そのスレッドの一部として、Dave Crossland は次のように書きました。
インパラリが最初だったようだ。エリックが2位になりました。 3番目があることを願っています。
ここでは、フォント生成の倫理的または法的側面、つまりこれをすべきかすべきでないかについて話すつもりはありません。これらは重要な質問ですが、ここでは取り上げません。しかし、それを読んだときの私の反応は、「エリックは2番目ではなかった。彼は22番目でもなかった。ただ、他の全員が今目の前にいないだけだ」というものでした。
したがって、この記事の目的は、フォント生成研究の発展の一部を調査し、そこにあるもの、それがどのように機能するか、そして将来何が起こると私が期待しているかについての明確な視点を提示することです。なぜなら、今後さらに多くのことが起こるからです。私たちが望んでいるかどうかに関係なく、人々はこの問題に取り組むことになるでしょう。したがって、あなたの本能がそれを歓迎したいのか、批判したいのか（あるいはもちろん両方）、それを認識し、それを遵守するようにしましょう。

我慢してください。
しかし、それを自分自身を理解する機会としても利用しましょう。調査中に、いくつかの要因といくつかのギャップに気づきました。おそらく、技術的な内容ではなく、これらのギャップについて話すことが、この投稿の本当のポイントになると思います。私が話しているのは、書体デザイナーと ML エンジニアとの間のギャップ、そして西洋の書体デザイナーと中国/日本/韓国の書体デザイナーとの間のギャップについてです。
必須の 2 つの文化に関するリファレンス
私が気づいた特徴の 1 つは、ML 研究者とタイプデザイナーの間で、分野を超えたコラボレーションが明らかに欠如していることです。 ML 研究者とタイプ エンジニアの間のコラボレーションが欠如していることに確かに気づきました。いくつかのモデルのコードを調べていると、フォントからグリフのアウトラインとラスタライズがどのように抽出されているかを読むと身震いします。 type-world の誰もこれにまったく関与していないことは明らかです。しかし、より一般的に言えば、品質管理レベルにも欠陥があることがわかります。 ML エンジニアがタイポグラフィーの成功と考えるものは、書体デザイナーが公の場で共有する勇気のないものです。私はこのスレッドだけでなく、Neural Axis variations のようなことも考えています。これは本当に興味深いアイデアで、おそらくうまく実装できますが、明らかに出版前にタイプ デザイナーによるレビューが必要でした。
これは「完全に機能するバリアブルフォント」ではありません。
私は人々の仕事を嘲笑するためにこれを言っているのではありません。重要なのは、自分が何を知らないのかが分からないということです。その理由について私には仮説があります。書体デザイナーとのコラボレーションに関しては、認識される必要性について明らかな疑問があると思います。
ML 研究者はコンピューター グラフィックスを扱うことに非常に慣れています。それは彼らが常に行っていることであり、本当にうまくやっています。そして、私としては

TypeDrawers という大きなスレッドで、「文字のデザインは非常に欺瞞的です。なぜなら、一連の認識可能な文字の形を描くのは簡単だからです」と主張し、そのせいでフィードバックが必要であることに気づきにくくなります。 ML モデルがアミノ酸の長い文字列を出力する場合、それが意味があるかどうかを生物学者に教えてもらう必要があることがわかります。このような状況では、分野を超えたコラボレーションが明らかに必要であるため、そうしないとは考えられません。
しかし、あなたがコンピューター グラフィックスの作業に慣れている研究者で、その仕事が文字「a」を描くことである場合は、「a」がどのようなものであるかを誰もが知っています。それには専門家の意見は必要ありません。 （ネタバレ注意：実際にそうしていることが判明しました！）
そして同時に…これをどう表現すればいいでしょうか？たとえ ML 研究者が AI 支援によるフォント生成に協力するためにタイポグラフィー コミュニティに手を差し伸べたとしても、必ずしも温かく心から歓迎されるとは限りません。なんだかデリケートな話題ですね。タイプの人は、倫理や合法性について厄介な質問をする癖があります。それに加えて、自分や同僚を失業させる責任を負う人は誰も望んでいません。それはわかります。
つまり、ここには矛盾があります。AI が生成したフォント デザインは、効果的かつ高品質なものとなるためには、タイポグラフィー コミュニティとの深い関わりからのみ生まれます。そして現時点で、タイポグラフィー コミュニティは実際には、効果的で高品質な AI 生成のフォント デザインを望んでいません。
これは活版印刷コミュニティにとって良いニュースだと思いますね?研究者が悪いモデルを作成するのは、タイポグラファーを探していないか、タイポグラファーが一緒に仕事をしてくれないためです。そして、タイプデザイナーは悪いモデルの出力を見て、自分たちの仕事が安全であると安心します。きちんとした閉ループ。
ただし、

帽子はとても西洋的な視点です。
フォント デザイン用の AI ツールに対するデザイナーからの主な反発の 1 つは、AI ツールが単に「面白いことをする」だけではないことを確認する必要があるということです。デザイナーは絵を描くのが好きでレターフォームを描くので、たとえまともなレターフォームであっても、レターフォームを描画する AI ツールは、「退屈な部分」を解決するのに役立たず、逆効果です。文字デザインの退屈な部分は、スペース、カーニング、グリフセットの拡張と言語サポート、デザインスペースの拡張など、デザイナーごとに異なります。それについては私たちも同意できません。しかし、明確に一致していると思われることの 1 つは、実際の形状を描くのが楽しい部分であり、それを奪う AI モデルは望まないということです。
ただし、私たち全員が同じ目標を持っているわけではありません。描くキャラクターが 100 人、さらには 1,000 人いる場合、自分のデザインがそれぞれのキャラクターにどのように反映されるかを考えるのは、楽しくて興味深いパズルです。
描かなければならないキャラクターが 10 万人もいると、それはとても楽しくて興味深いものではなくなり、非常に苦痛なものになってしまいます。中国語のフォント デザインはまったく別のやかんであり、デザイナー チームは 1 つのフォントを作成するのに長年の作業を必要とします。もちろん、それを軽減するのに役立つ可変コンポーネントなどのいくつかの自動化はありますが、それでも...描画は中国語活字デザインの「楽しい部分」ではありません。
deepvecfont-v2 コードベースを調べているときに、すべてのモデル トレーニング データが利用可能になっていることに気付きました。データダンプ全体とすべてのフォントをダウンロードできます。これはちょっと珍しいですね。フォントベースの ML プロジェクトがトレーニング データを利用可能にしているのはあまり見かけません。その主な理由は、研究者が集めたフォントの網に「偶然」どれだけ多くの商用フォントが巻き込まれたのかを知ることができるからです。

インターネットから削り取られた。しかし、これは違いました。 deepvecfont-v2 プロジェクトのトレーニング データ リリースの一部として、商用中国語フォントの巨大なキャッシュが公開されました。これはどのようにして可能でしょうか?これも合法なのでしょうか？
さて、リポジトリの下部にある小さなメモには次のように書かれています。
すべての中国語フォントは Founder から収集されたものであり、Founder の許可なしにフォントを商用目的で使用することはできないことに注意してください。
これをより明確に説明すると、西側のファウンドリは「AI モデルのトレーニングに当社のフォントを使用することはできない」と明言している一方で、中国の主要活字ファウンドリの 1 つである Founder Type は、AI ベースのフォント生成を研究するチームと提携し、「インターネット上の誰にでも当社の中国語フォント カタログ全体の無料ライセンスを提供します。ただし、AI モデルのトレーニングに使用する場合に限ります」と述べています。
それは…西洋ではおそらく起こらないだろう。状況が異なれば、目的も異なります。中国の活字ファウンドリは、「退屈な部分」を解決するのに役立つ可能性があるため、AI フォントの研究を歓迎しています。彼らは、「少数ショットのフォント生成」の問題、つまり、私が特定のスタイルで 50 個のグリフを描画し、コンピューターが残りを埋めるだけというアイデアに、重大な商業的関心を持っています。西洋の鋳物工場は、全体的にこの熱意を共有していません。 (これは、昨年のATypIコペンハーゲンで顕著に表れた力学でした。)
これが意味するのは、AI フォント生成に関して行われている興味深い取り組みの多くが中国と日本で行われており、ファウンドリのサポートとともに行われているということです。ラテン語フォント生成の現状は、まあ、かなり貧弱ですが、中国語の AI フォント生成はすでに本番環境に対応しています。
これはダメですよね？プロダクション対応の中国語フォントを生成できれば、ラテン文字のデザインを求められることになるでしょう。

の上。まあ、たぶん。しかし、そうではないかもしれません。周りを見回してみれば、現時点では実稼働対応のラテン語フォント生成モデルは存在せず、存在するモデルは非常にひどいものであることがわかるでしょう。これはなぜでしょうか?何が起こっているのか？後でいくつかのアイデアを。
革命は TypeDrawers スレッドではない
とりあえず、最後にフォント生成の現状を見てみましょう。私たちは、ML エンジニアが書体デザイナーと協力して作業していないことを確認しました。これが意味するのは、タイポグラフィ メディアにおけるフォント生成の進歩について知ることはできないということです。 ML コミュニティの人々がどこで互いにコミュニケーションをとっているのか、会議資料やプレプリント、より具体的には Arxiv を探し始める必要があります。 Arxiv は、科学者が進行中の研究論文を共有する場所であり、特に「コンピューター ビジョンとパターン認識」(cs.CV) サブセクションは、ML っぽいグラフィックっぽいものがすべて最終的に行き着く場所です。フォント生成の研究で実際に何が起こっているのかを知りたい場合は、そこに行く必要があります。今日 arxiv で「font」を検索すると、328 件の論文がヒットします。すべてがフォント生成に関するものではありませんが、多くはフォント生成に関するものです。予想通り、そのほとんどは中国、日本、韓国出身者だ。それらの多くは、予想どおり、「少数ショットのフォント生成」の問題に関するものです。50 の参考文献から 50,000 文字のフォントを描画してください。
フィリップ・パルディアは、2024 年初頭までの最先端プロジェクトのかなり網羅的なリストをまとめました。しかし、この分野はそれ以来大幅に発展しました。それらを簡単に紹介し、最新情報をまとめてみましょう。以下のセクションは代表的なものであり、すべてを網羅するものではありません。本当にこれに興味がある場合は、Arxiv で自分で検索するのが最適です…
敵対的生成ネットワーク
初期のフォント生成作業

(2018-2023) は、DALL-E (初版) のような画像ジェネレーターの基盤となる同じテクノロジーである敵対的生成ネットワークに基づいて構築される傾向がありました。 GAN は、「偽造者」と「探偵」の 2 つのモデルで構成されます。偽造者は画像を作成しようとし、探偵はそれが機械で生成されたものかどうかを検出しようとします。偽造者が探偵が本物だと思う画像を作成できるまで、この作業を繰り返します。その時点で作業は完了です。
フォント生成には、条件付き GAN が使用されました。単に画像を生成しようとするのではなく、GAN に特定のグリフ (コンテンツ) と特定のスタイルを与え、そのスタイルに一致するグリフ画像を「偽造」しようとします。フォント GAN モデルの命名は信じられないほど想像力に欠けていました。彼らは行きました：
DG-Font (2021): OG というわけではありませんが、最も影響力のあるものの 1 つです。
LF-Font (2021): 「スタイル」をコンポーネント レベルの機能の階層的なセットとして捉えました。
MX-Font (2021): 複数のコンポーネント レベルの「専門家」が、グリフのさまざまな部分を作成する方法を学習します。 (現在は MX-Font++ がありますが、まだ見ていません。)
CG-GAN (2022): 別の命名スキームを試すための少なくともいくつかのポイント。
FS-Font (2022): フォント スタイルを定義するローカルの詳細 (端末など) をさらに特定しようとします。
CF-Font (2023): スタイルとコンテンツをトレーニング用の 1 つの事前条件に融合します。
DA-Font (2025): コーナーと「弾性メッシュ」を改善して構造を改善するさまざまな方法を試みました。

[切り捨てられた]

## Original Extract

Two disclaimers: First, this article was entirely human generated; no AI was involved in the writing. Second, I have spent a lot of time this year attempting to train a glyph generation model. I want to be up front about this because I know that for certain people, this will discredit me entirely, a
[truncated]

The State of AI Font Generation | Font and Text Technology
simoncozens.github.io
The State of AI Font Generation
Two disclaimers: First, this article was entirely human generated; no AI was involved in the writing. Second, I have spent a lot of time this year attempting to train a glyph generation model. I want to be up front about this because I know that for certain people, this will discredit me entirely, and I don’t want you to find out afterwards and feel betrayed. But then, would you rather read an article about AI font generation written by someone who didn’t have any experience of it?
There’s been a lot of discussion recently about the ability of machine learning models to generate new fonts, or “fill in” existing fonts (by expanding the glyphset or design space), much of it prompted by this thread on Typedrawers . As part of that thread, Dave Crossland wrote:
Impallari appears to have been the first. Eric’s become second. I hope there’ll be a third.
I’m not going to talk here about the ethical or legal side of font generation, about whether you should or shouldn’t do this. These are important questions, but not for here; but my reaction when I read that was “Eric wasn’t the second; he wasn’t even the twenty-second - it just that all the others aren’t in your face right now.”
So the point of this article is to try to survey some of the development in font generation research, to present a clear-eyed view of what’s out there, how it works, and what I expect coming down the line. Because more will be coming down the line. People are going to work on this problem, whether we want them to or not. So whether your instinct is to welcome it or critique it (or of course both), let’s make sure we’re aware of it and we understand it.
But let’s also use it as an opportunity to understand ourselves as well. During my investigations, I noticed a couple of factors, a couple of gaps, and I think perhaps talking about these gaps is going to be the real takeaway from this post, rather than the technical stuff. I’m talking about the gaps between type designers and ML engineers, and also the gaps between Western type designers and Chinese/Japanese/Korean type designers.
The obligatory Two Cultures reference
One feature I’ve noticed is a very obvious lack of collaboration across disciplines, between ML researchers and type designers. I’ve certainly noticed a lack of collaboration between ML researchers and type engineers - as I look into the code for some of the models, I shudder to read how glyph outlines and rasterizations are being extracted from fonts; it’s clear that nobody from type-world has had any input into this at all. But more generally, I can see that there is also a lack at the quality control level; what ML engineers consider a typographic success would be something that a type designer would never dare share in public. I’m not just thinking of this thread , but also stuff like Neural Axis Variations - a genuinely interesting idea which probably could be well implemented, but clearly needed a type designer to review it before publication.
This is not “a fully functional variable font”.
I don’t say this to mock people’s work. The whole point is that you don’t know what you don’t know . I have a theory about why this is: when it comes to collaboration with type designers, I think there is an obvious question of perceived need.
ML researchers are really used to working with computer graphics. It’s something that they do all the time and do really well. And, as I argued in the big TypeDrawers thread, “type design is very deceptive because it’s easy to draw a series of recognisable letter forms” and that makes it hard for you to realise that you need feedback. If your ML model outputs a long string of amino acids, you know that you’re going to need a biologist to tell you whether it makes sense or not. Collaboration across disciplines is so obviously necessary in such a situation that you wouldn’t think not to.
But if you’re a researcher used to working with computer graphics and your task is to draw the letter ‘a’… well, everyone knows what an ‘a’ looks like. You don’t need professional input for that. (Spoiler alert: turns out you do!)
And at the same time… how can I put this delicately? Even if ML researchers did reach out to the typographic community in order to collaborate on AI-assisted font generation, they may not necessarily receive a warm and hearty welcome. It’s kind of a touchy subject. Type people have a habit of asking awkward questions about ethics and legality. And on top of that, nobody wants to be the one responsible for putting themselves and their peers out of work. You know, I get that.
So there is a paradox here: that AI generated font design can only come out of deep engagement with the typographic community if it’s going to be effective and high quality, and right now the typographic community doesn’t actually want effective and high quality AI generated font design.
Which I guess is good news for the typographic community, right? Researchers produce bad models because they either don’t seek out typographers or typographers won’t work with them, and type designers see the output of the bad models and feel reassured that their jobs are safe. A neat closed loop.
Except, that’s such a Western perspective.
One of the main pushbacks from designers about AI tooling for font design is that it has to make sure that it’s not just “doing the fun bit”. Designers draw letterforms because they like drawing, and so AI tools which draw letterforms - even decent letterforms - but don’t help with the “boring bit” are counterproductive. The boring bits of type design are different for each designer: spacing, kerning, glyphset expansion and language support, designspace expansion. We can disagree about that. But one thing that does seem to be a clear consensus is that drawing the actual shapes is the fun bit, and there’s no desire for an AI model which takes that away.
Except, not all of us have the same objectives. When you have a hundred or even a thousand characters to draw, working out how your design inhabits each of those characters is a fun and interesting puzzle to solve.
When you have a hundred thousand characters to draw, it stops being quite so fun and interesting and turns into a royal pain in the arse. Chinese font design is a different kettle of fish altogether, requiring teams of designers many years of work to produce a single font. Of course, we have some automations such as variable components which can help bring that down, but even so… drawing is not the “fun bit” of Chinese type design.
While I was looking at the deepvecfont-v2 codebase, I noticed that all of the model training data was made available. You can download the whole data dump, all the fonts . This is kind of unusual. You don’t often see font-based ML projects making their training data available, largely because it allows you to see quite how many commercial fonts were, uh, “accidentally” caught up in the dragnet of fonts that the researchers scraped off the Internet. But this one was different. A huge cache of commercial Chinese fonts were made publicly available as part of the training data release for the deepvecfont-v2 project. How is this possible? Is this even legal?
Well, a little note at the bottom of the repository says:
Please note that all the Chinese fonts are collected from Founder, and the fonts CANNOT be used for any commercial uses without permission from Founder.
To spell that out more explicitly: While Western foundries are explicitly saying “you can’t use our fonts to train AI models”, Founder Type, one of the main Chinese type foundries, have partnered with a team researching AI-based font generation and said “We are giving anyone on the Internet a free license to our entire Chinese font catalogue but only if you use it to train an AI model.”
That… probably wouldn’t happen in the West. Different situations have different objectives. Chinese type foundries welcome AI font research because it can help them with their “boring bit”. They have a serious commercial interest in the problem of “few-shot font generation”: the idea that I draw fifty glyphs in a particular style, and the computer just fills in the rest. Western foundries do not, on the whole, share this enthusiasm. (This was a dynamic which was very much in evidence at ATypI Copenhagen last year.)
What this means is that a lot of the interesting work going on in terms of AI font generation is happening in China and Japan, and is happening with foundry support. And while the current state of Latin font generation is, well, pretty poor, AI font generation in Chinese is already production-ready.
This is bad, right? If they can generate production-ready Chinese fonts, then they’re coming for Latin type design soon. Well, maybe. But maybe not. Look around you and you will discover that there aren’t production-ready Latin font generation models right now, and those which do exist are pretty woeful. Why is this? What is going on? Some ideas later.
The revolution will not be a TypeDrawers thread
For now, let’s finally take a look at the current state of font generation. We’ve established that ML engineers are not working in collaboration with type designers, and what that means is that you aren’t going to find out about font generation advances in the typography media. You’re going to need to start looking where people in the ML community communicates with one another: conference papers and pre-prints, and more specifically Arxiv . Arxiv is where scientists share in-progress research papers, and in particular, the “Computer Vision and Pattern Recognition” (cs.CV) subsection is where all the ML-y graphic-y stuff ends up. If you want to find out what’s actually going on in font generation research, that’s where you will need to go. A search for “font” on arxiv today turns up 328 papers. Not all of them are about font generation, but a lot of them are. Most of those, as predicted, are from China, Japan and Korea. Many of those are, predictably, about the “few-shot font generation” problem: draw me a 50,000 character font from 50 references.
Filip Paldia put together quite an exhaustive list of state-of-the-art projects up to early 2024. But the field has developed considerably since then. Let’s take a quick tour of them, and bring things up to date. The following sections will be representative , not exhaustive . If you’re really into this, you’re best placed doing your own searches on Arxiv…
Generative Adversarial Networks
Early font generation work (2018-2023) tended to be built on Generative Adversarial Networks, the same technology underlying image generators like DALL-E (first edition). GANs are comprised of two models, the “forger” and the “detective”; the forger tries to create an image and the detective tries to detect whether or not it’s machine-generated. They keep iterating on this task until the forger manages to create an image which the detective thinks is real, at which point it’s good to go.
For font generation, conditional GANs were used - we don’t just try to produce any image, but we give the GAN a particular glyph (content) and a particular style, and try to “forge” a glyph image matching that style. Font GAN model naming was incredibly unimaginative. They went:
DG-Font (2021): Not quite the OG, but one of the most influential.
LF-Font (2021): Saw “style” as a hierarchical set of component level features.
MX-Font (2021): Multiple component-level “experts” learn to produce different parts of a glyph. (There’s now an MX-Font++ but I haven’t looked at it yet.)
CG-GAN (2022): A few points at least for trying a different naming scheme.
FS-Font (2022): Tries to further identify the local details (terminals etc.) which define font style.
CF-Font (2023): Fuses style and content into a single prior for training.
DA-Font (2025): Tried various ways to improve corners and an “elastic mesh” to improve structure.

[truncated]
