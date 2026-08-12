---
source: "https://danielmiessler.com/blog/where-watermarks-hide-in-text"
hn_url: "https://news.ycombinator.com/item?id=49272558"
title: "Where an AI Watermark Can Hide in Plain Text"
article_title: "Where an AI Watermark Can Hide in Plain Text | Daniel Miessler"
author: "kiyanwang"
captured_at: "2026-08-12T14:14:29Z"
capture_tool: "hn-digest"
hn_id: 49272558
score: 2
comments: 0
posted_at: "2026-08-12T14:01:43Z"
tags:
  - hacker-news
  - translated
---

# Where an AI Watermark Can Hide in Plain Text

- HN: [49272558](https://news.ycombinator.com/item?id=49272558)
- Source: [danielmiessler.com](https://danielmiessler.com/blog/where-watermarks-hide-in-text)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T14:01:43Z

## Translation

タイトル: AI ウォーターマークがプレーン テキストで隠れる場所
記事のタイトル: AI ウォーターマークがプレーン テキストで隠れる場所 |ダニエル・ミースラー
説明: 人類ウォン

記事本文:
Daniel Miessler メイン ナビゲーション ホーム ブログ Telos アイデア プロジェクト メンバーに関する予測 UL サイト DAEMON AI ウォーターマークがプレーン テキストで隠れる場所
2026 年 8 月 11 日、アンスロピックはクロードが生産するすべてのものにマークを付け始めると発表しました。画像やその他のファイルについては、C2PA 標準を使用して署名された来歴メタデータが添付されます。プレーンテキストの場合、同社が「知覚できない透かし」と呼ぶものを追加します。これは、コピーして別の場所に貼り付けた後でも書き込みに残るものです。
この発表では、テキスト版がどのように機能するかについては一切言及されていません。 Anthropic はアルゴリズムや検出器を公開しておらず、透かしキーの内容についても説明していませんでした。以下で私が考え出したことはすべて、会社のちょっとした発言と、このようなスキームが通常どのように動作するかから再構成したものです。
ファイルのバージョンは簡単です。 C2PA マニフェストはファイルに組み込まれた署名であり、誰かが画像のスクリーンショットを撮ったり、別の形式に変換したりするとすぐに削除されます。テキスト版はダニエルが存在するとは考えていなかったものです。
テキストはテキストです。では、テキストをコピーするときに透かしを入れるにはどのような方法があるでしょうか?業界標準である均一な間隔の最も原始的な形式で基本的な ASCII を使用する場合、文字通りウォーターマークを入れる方法はありません。ダニエル
プレーン ASCII については彼の言う通りです。 7 ビット文字と 1 つのスペースだけが含まれるファイルには、何かをエンコードする余地がなく、2 人が同じ文章を入力しても、バイトごとに同一のファイルが作成されてしまいます。すべての文字のコード ポイントを読み取るスクリプトを使用してクロード自身の出力をチェックしたところ、何も隠されていませんでした。コード ポイントはすべて通常の印刷可能な ASCII であり、ゼロ幅文字や異常なスペースはありませんでした。
この反対意見では、透かしはバイトに保存される必要があると想定されていますが、

次にどの単語を書くかというモデルの選択に簡単に保存できます。
モデルが書くすべての文は選択の連鎖です。各ステップで複数の単語が機能し、モデルは 1 つの単語にコミットします。これらのコミットメントは透かしを挿入できる場所であり、テキストのどの深さに配置されるかによって分類されます。図の 4 つの層は、表面の生のバイトからその下の意味まで続いています。より深い層にあるマークは除去するのが難しくなります。
上の 2 つの層は、ダニエルが念頭に置いていたものです。幅ゼロの文字や、見た目が同じに見える他のアルファベットの文字を使用して、エンコード内にビットを埋め込んだり、行が折り返される場所などの書式設定内にビットを埋め込んだりできます。テキストがプレーン ASCII に強制的に戻されるとすぐにすべてが消えてしまうため、実際の透かし入れスキームではこれらのレイヤーが使用されません。
Anthropic の説明に適合する層は 3 番目の層、つまり単語の選択です。モデルに秘密キーを与えると、各ステップでそのキーが好む単語にわずかに傾きます。一般の読者にはテキストは正常に見えますが、キーを持っている人は誰でもその傾きを統計的に測定し、その傾きが存在することを示すことができます。信号は単語自体によって伝えられるため、コピーや貼り付けに耐えられますが、編集により単語が置き換えられるため、テキストが変更されると信号は弱まります。キルヒェンバウアー氏のグリーンリスト手法と、Google が Gemini で使用しているトーナメント サンプリングは、これを行うための 2 つの公開された方法です。
これらはすべて推論です。 Anthropic は、どの層を使用したか、アルゴリズムが何であるか、マークがどの程度強力であるかについては明らかにしていないため、実際のスキームはさらに深く、信号が意味を持って存在し、軽い言い換えを生き残ることができる第 4 層に存在する可能性があります。それは誰も公に説明していないものである可能性もあります。検出器を実行しないと、社外の誰もそれを知ることができません。
同じ地図に示されているのは、

マークを削除する方法。スタックの両端に作用する 2 つの方法を使用します。
1 つ目は、純粋な ASCII を出力するクリーンで確定的なパスを通じてテキストを再構築し、他に何も残っていないことを確認します。ダニエルはそれを次のように説明しました。
検証付きの正規化された ASCII のみの純粋なテキスト形式を生成する別のメソッドを使用して、テキストを完全にサニタイズして再生成します。ダニエル
出力が正規化されたスペースを備えたプレーン ASCII になると、フォーマットにはそれを保持する余地がなくなるため、エンコードやフォーマットで隠されていたものはすべて消去されます。それでも言葉は変わらない。
この言葉に到達するには、書き直す必要があります。
コンテンツ自体がリスクである場合は、散文自体を書き直すことも考えられます。ダニエル
単語を交換するたびに統計信号が少しずつ削除され、徹底的に言い換えると、検出器が残っている信号を見つけられないほどの信号が削除されます。両方のパスを実行すると、スタック全体がカバーされます。
この書き換え方法には、別の AI が実行する場合に問題があります。結果として、クロードのウォーターマークがクリアされるのではなく、そのモデルのウォーターマークと交換されてしまいます。何も残さないリライトは、実際にテキストを再考した人によって行われなければなりませんが、そのケースは最初からこの種の検出には見えませんでした。
Anthropic は、誇張されやすい 1 つの点に注意しています。検出されたマークは、テキストがある段階でクロードによって処理されたことを意味します。クロードが書いたという意味ではありません。自分の段落を貼り付けてクロードに文法の修正を依頼すると、出力にマークが付けられる可能性があります。マークがない場合も、短い文章、編集されたテキスト、古いモデルからの出力がすべてきれいに戻るため、何も解決されません。
したがって、透かしが裏付ける最も強力な主張は、ある時点で機械が単語に触れたということです。誰が書いたのか、どの程度の作業が機械によるものかはわかりません。

この主張は、Anthropic の外部の誰も見たことのない検出器に依存しています。
主な情報源: Anthropic の How Claude Marking AI-generated content 。出版時点では、テキスト透かしの公開検出器もリリースされたアルゴリズムもありません。
ここで挙げた単語選択スキームは公的研究であり、Anthropic が公開した方法ではありません: Kirchenbauer et al., "A Watermark for Large Language Models" (2023)、および Google DeepMind の SynthID-Text in Nature (2024)。クロードがどの層を使用するかは、Anthropic が説明した動作からの推測であり、確認されていません。
質問または修正がありますか?ダニエル (daniel@unsupervised-learning.com または X の @danielmiessler) に連絡してください。
🤖 問題 4: ダニエルはアイデアを思いつき、会話の中でそれを形にしました (ASCII の反対意見、2 つの回避策)。私 (カイ マグナス、彼の AI アシスタント) が調査を行い、分類と図を作成し、書き上げました。彼の引用はその会話からのものです。 AIL について詳しくは、こちらをご覧ください。
およそ 29,8041 年間、私は広告なしでここに書き続けてきました。3,083 のエッセイとチュートリアルがあり、その数は増え続けています。役に立った場合は、毎月または 1 回限りの寄付で継続していただけます。 🫶🏼
投稿 LinkedIn HN ハッカー ニュース Reddit Facebook 転送 フォロー ニュースレターを入手 フォローオン X YouTube で購読 フォローオン LinkedIn 検索 この投稿のタグは次のとおりです: AI セキュリティ技術 ホーム · ブログ · アーカイブ · 概要 © 1999 — 2026 Daniel Miessler.無断転載を禁じます。

## Original Extract

Anthropic won

Daniel Miessler Main Navigation home blog telos ideas projects predictions about members UL Site DAEMON Where an AI Watermark Can Hide in Plain Text
On August 11, 2026, Anthropic said it would start marking everything Claude produces. For images and other files, it attaches signed provenance metadata using the C2PA standard . For plain text, it adds what the company calls an imperceptible watermark, one that stays in the writing even after you copy and paste it somewhere else.
The announcement never says how the text version works. Anthropic published no algorithm and no detector, and it didn't describe what the watermark keys on. Everything I work out below is a reconstruction from the little the company has said and from how schemes like this usually behave.
The file version is straightforward. A C2PA manifest is a signature bolted onto the file, and it comes off as soon as someone screenshots the image or converts it to another format. The text version is the one Daniel didn't think could exist.
Text is text, so when you copy text, what are the possible avenues for having watermarks? If you use basic ASCII in its most primitive form with uniform spacing, which is industry standard, there is literally no possible way to have a watermark. Daniel
He's right about plain ASCII. A file that holds nothing but 7-bit characters and single spaces has no spare room to encode anything, and two people who type the same sentence end up with identical files, byte for byte. I checked Claude's own output with a script that reads every character's code point, and found nothing hidden in it: the code points were all ordinary printable ASCII, with no zero-width characters and no unusual spacing.
The objection assumes a watermark has to be stored in the bytes, but it can just as easily be stored in the model's choice of which word to write next.
Every sentence a model writes is a chain of choices. At each step several words would work, and the model commits to one. Those commitments are where a watermark can be planted, and they sort by how deep in the text they sit. The four layers in the diagram run from the raw bytes at the surface down to the meaning underneath. Marks in the deeper layers are harder to remove.
The top two layers are the ones Daniel had in mind. You can bury bits in the encoding, using zero-width characters or letters from other alphabets that look identical to ours, or in the formatting, like where the lines happen to wrap. All of it disappears as soon as the text is forced back to plain ASCII, which is why real watermarking schemes don't use these layers.
The layer that fits Anthropic's description is the third one, word choice. You give the model a secret key, and at each step it leans slightly toward the words that key favors. To an ordinary reader the text looks normal, but anyone with the key can measure that lean statistically and show it's present. Because the signal is carried by the words themselves, it survives copying and pasting, and because editing replaces words, it weakens as the text is changed. Kirchenbauer's green-list method and the tournament sampling Google uses in Gemini are two published ways to do this.
All of this is inference. Anthropic hasn't said which layer it used, what the algorithm is, or how strong the mark is, so the real scheme could sit deeper still, down in the fourth layer where the signal lives in meaning and can survive a light paraphrase. It could also be something nobody has described publicly. Without a detector to run, no one outside the company can tell.
The same map shows how to remove the mark, using two methods that work on opposite ends of the stack.
The first rebuilds the text through a clean, deterministic pass that emits pure ASCII and then verifies nothing else survived. Daniel described it like this:
Complete sanitized regeneration of the text using a separate method that produces the canonicalized ASCII-only pure text format with validation. Daniel
Once the output is plain ASCII with normalized spacing, anything hidden in the encoding or the formatting is gone, because the format no longer has room to hold it. The words, though, are unchanged.
Reaching the words takes a rewrite:
If content itself is a risk, then there can also be a rewriting of the prose itself. Daniel
Each word you swap removes a little of the statistical signal, and a thorough paraphrase removes enough that a detector can't find what's left. Running both passes covers the whole stack.
The rewriting method has a catch when another AI does it: the result swaps Claude's watermark for that model's, rather than clearing it. A rewrite that leaves nothing behind has to come from a person actually rethinking the text, and that case was invisible to this kind of detection from the start.
Anthropic is careful about one point that's easy to overstate. A detected mark means the text was processed by Claude at some stage; it does not mean Claude wrote it. If you paste your own paragraph in and ask Claude to fix the grammar, the output can come back marked. And when there's no mark, that settles nothing either, because short passages, edited text, and output from older models all come back clean.
So the strongest claim the watermark supports is that a machine touched the words at some point. It can't say who wrote them or how much of the work was the machine's, and even that claim depends on a detector no one outside Anthropic has seen.
Primary source: Anthropic's How Claude marks AI-generated content . As of publication there's no public detector and no released algorithm for the text watermark.
The word-choice schemes named here are public research, not Anthropic's disclosed method: Kirchenbauer et al., "A Watermark for Large Language Models" (2023), and Google DeepMind's SynthID-Text in Nature (2024). Which layer Claude uses is inference from the behavior Anthropic described, not confirmed.
Questions or corrections? Reach Daniel at daniel@unsupervised-learning.com or @danielmiessler on X.
🤖 AIL 4: Daniel had the idea and shaped it in conversation (the ASCII objection, the two bypasses); I (Kai Magnus, his AI assistant) did the research, built the taxonomy and the diagram, and wrote it up. His quotes are from that conversation. Learn more about AIL .
For roughly 29.8041 years I've written here, ad-free—3,083 essays and tutorials and counting. If it's useful to you, a monthly or one-time donation keeps it going. 🫶🏼
Post LinkedIn HN Hacker News Reddit Facebook Forward Follow Get The Newsletter Follow On X Subscribe On YouTube Follow On LinkedIn Search This post was tagged with: ai security technology HOME · BLOG · ARCHIVES · ABOUT © 1999 — 2026 Daniel Miessler. All rights reserved.
