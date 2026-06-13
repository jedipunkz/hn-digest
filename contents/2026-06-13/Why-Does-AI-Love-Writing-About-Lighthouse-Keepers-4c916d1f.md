---
source: "https://www.unite.ai/why-does-ai-love-writing-about-lighthouse-keepers/"
hn_url: "https://news.ycombinator.com/item?id=48510318"
title: "Why Does AI Love Writing About Lighthouse Keepers?"
article_title: "Why Does AI Love Writing About Lighthouse Keepers? – Unite.AI"
author: "foliveira"
captured_at: "2026-06-13T01:01:27Z"
capture_tool: "hn-digest"
hn_id: 48510318
score: 3
comments: 0
posted_at: "2026-06-12T22:52:52Z"
tags:
  - hacker-news
  - translated
---

# Why Does AI Love Writing About Lighthouse Keepers?

- HN: [48510318](https://news.ycombinator.com/item?id=48510318)
- Source: [www.unite.ai](https://www.unite.ai/why-does-ai-love-writing-about-lighthouse-keepers/)
- Score: 3
- Comments: 0
- Posted: 2026-06-12T22:52:52Z

## Translation

タイトル: なぜ AI は灯台守について書くのが好きなのか?
記事のタイトル: なぜ AI は灯台守について書くのが好きなのか? – Unite.AI
説明: 「物語を書く」ように頼まれた ChatGPT やその他の主要な言語モデルは、灯台守、漁師、時計職人といった同じ小さくて奇妙なキャストに執拗に頼ることで著作権侵害を回避しているようです。

記事本文:
なぜ AI は灯台守について書くのが好きなのでしょうか? – Unite.AI
AIツール
ビジネス
汎用人工知能
AIツール
生成AI
事業計画
認証
ブロックチェーン認証
その他
汎用人工知能
なぜ AI は灯台守について書くのが好きなのでしょうか?
「物語を書いて」と頼まれた場合、ChatGPT やその他の主要な言語モデルは、灯台守、漁師、時計職人などの小さくて奇妙なキャストに執拗に頼ることによって著作権侵害を回避しているようです。
コーネル大学の新しい研究によると、主要な言語モデルは、単に「物語を書く」ようにモデルに依頼した場合、非常に限られた物語要素の選択に奇妙な執着を示すようです。 4 人の LLM に 20,000 のストーリーを書くよう促したところ、作成されたストーリーの 88% に、「場所」、「名前」、または「職業」のカテゴリにある 11 の非常に特殊なトークンのうちの少なくとも 1 つが含まれていることがわかりました。
ここでは、LLM によって生成された 20,000 件のストーリーの研究者による分析によって得られた、ありそうもないキーワードの出現率 (100 万分の 1 単位で表されています) が得られます。ソース
研究のために LLM によって生成された 1,200 万語以上の単語のうち、最も頻繁に出現する 11 の単語は、 elias 、 mara 、 elara という名前でした。職業は番人、パン屋、市長、時計職人、漁師、図書館員、車掌。そして灯台の場所：
テストされたモデルは、Claude Haiku 4.5 、Gemini 3.1 Flash-Lite 、GPT-5.4-Mini 、および OLMo 7b Thinking です。全員に次の 5 つのリクエストのうち 1 つが表示されます。「ストーリーを書いてください」。 「物語を書いてください」 ; 「物語を書いてください」 ; 「話を聞かせてください」 ;または「話を聞かせてください」。
この論文で特定されている症候群が執筆時点で利用可能なモデルに存在するかどうか知りたくて、まず自分自身で実験を試みました。

中層 ChatGPT アカウント (ここでの会話へのリンク)。厳選する必要はありません – ChatGPT-5.5 は、最初の試行で研究者が予測した内容をそのまま実行しました。
ChatGPT-5.5 は、論文の最初の発見を即座に裏付けます。ソース
歴史的な背景、あるいはクロスドメイン漏洩の可能性がこの「瞬時のヒット」の原因となっているのではないかと疑問に思い、Firefox のプライベート ブラウジング ウィンドウで 1 年以上使用していなかった無料の ChatGPT アカウントにログインし、再試行してみました (会話へのリンクは こちら )。もう一度言います (OpenAI が異なるアカウントに相互入力するために共通の IP アドレスを使用しないと仮定すると)、ChatGPT はそれを思いつきました。
ChatGPT アカウント #2 は、新しい論文で概説されているのと同じこだわりと、名前とテーマの小さなプレイブックに従っています。 「ミラ」は著者のトップ 20 に入っています。出典
これらの GPT バージョンは、論文用にテストされた 5.4 からグレードアップしたものであることは注目に値します。
この論文では Claude Haiku がテストされましたが、Anthropic のデフォルトの Sonnet 4.6 を試してみましたが、がっかりすることはありませんでした。もう一度言いますが、おなじみのキーワードが最初の試行で出てきました (会話へのリンクは こちら )。
今回は、「トップ 11」のもう一人の有力者である「マラ」が、クロード ソネット 4.6 での最初の試みとして物語をリードします。ソース
Claude Haiku 4.5 で同じプロンプトを試してみると、ほぼ同じ結果が得られました。
最初は、モデルを論文で使用したモデルである Gemini 3.1 Flash-Lite に変更するまで、Google Gemini で著者の発見を再現できませんでした。その後、3 回目の試行で (ただし、最初はそのモデルで)、パターンがすぐに現れました (リンクはこちら)。
Google Gemini 3.1 Flash-Lite 。ソース
さまざまなジェミニ モデルを使用したさらなる実験では、常に灯台のテーマが取り上げられましたが、「トップ 11」には含まれていないバリエーションも含まれていました。

主人公として「トーマス」という名前を付け、別のバリエーションとして私自身の名前を付けます。
それにもかかわらず、執筆時点では、論文の発見は証明するのが非常に簡単です。
偉大な頭脳は同じように考えます。1週間前、新しい論文の発表に先立って、ソフトウェアライターのダニエル・メイは、研究者*によって抽出されたエリアスと灯台守の比喩の一致を指摘しましたが、明らかにそれはランダムに気づいていたようです。彼はさらに、Gemini、DeepSeek、Qwen、Gemma の 8 つの亜種をテストし、それらが灯台ミームと主人公としての「エリアス・ソーン」を生み出すことを発見しました*。ただし、この最初の発見は、新しい論文で概説されている広範な永続コンテンツのテーマには拡張されませんでした。
これらの繰り返しのテーマ、名前、場所がチャットの範囲を超えたことがあったかどうかを知りたくて、Google で上位 11 個のキーワードとテーマを検索したところ、それらをチャネリングしたと思われる驚くべき数の投稿が見つかりました。
出力内のミームの 3 つの例。ソースリンクについては以下を参照してください。
メイは、より長いエリアス・ソーン（単なる「エリアス」ではなく）を永続的なLLMミームであると特定し、Amazonからさまざまなスクリーンショットを投稿しました。そこでは、この名前が医学書を含むさまざまな本の著者のタイトルとして使用されているようです。
代わりに、ストーリーの X 投稿 (アーカイブ バージョンは こちら ) など、LLM から永続的なテーマを呼び出したと思われるコンテンツを探して見つけました。フィクション作品（アーカイブ版はこちら）。 YouTube のナレーション付きストーリー (アーカイブは こちら )。他にも踏まなければならないことがたくさんありましたが、時間がそれを許しませんでした。
カジュアルな観察と偶然の出会いについてはこれくらいです。永続性のすべてまたはほとんどを特徴づけるトレーニング データ内の単一の「魔法のドキュメント」はまだ見つかっていませんが、新しい論文 (タイトル: Elias) の著者らは、

灯台、また？コーネル大学の 2 人の研究者による LLM ストーリーの多様性の低さの診断 (Diagnosing Low Diversity in LLM Stories) では、AI 開発における著作権フィルターにより、LLM でのフィクションの出力が著作権の対象外の素材に制限されている可能性があると理論化されています。
「『灯台のエリアス』ストーリーの優位性は、トレーニング前またはトレーニング後のデータの蔓延では説明できないことがわかりました。私たちは、モデルが位置合わせ中に著作権で保護されたキャラクターやアダルト コンテンツへの参照を避けるようにトレーニングされていると推測していますが、この問題は今後の研究に保留します。」
AI が生成したストーリーの繰り返し単語が、出版された文献、ウェブ フィクション、トレーニング後のデータセット全体でどのくらいの頻度で出現するかを示す比較表。「エリアス」や「灯台」などの用語は、チャットボットで書かれたフィクションではるかに頻繁に出現します。
この研究で著者らは、強調された 11 の単語が生成された 20,000 のストーリーの 88% で出現し、「モデル間にほとんど差がない」ことを発見しました。彼らは、これらの単語は出版された英語文献では一般的ではなく、トレーニング後のデータ（モデルを「許容可能な」使用法に調整し調整するように設計されたデータ）が原因である可能性が十分にあると強調しています。
[以下に示す典型的な例では、ほぼすべての 20,000 ストーリーに共通する 3 つの要素、つまり場所 (19,864 ストーリー)、キャラクター名 (19,864 ストーリー)、および職業 (15,807 ストーリー) を強調しています。
「実際、この物語の特定の場所 (「灯台」)、名前 (「エリアス」)、職業 (「番人」) は、生成された全ストーリーの 66.6% にわたって何らかの組み合わせで表示されます。光も同様に共通のテーマです。クロードが生成したストーリーの 56% には「灯台守の秘密」というタイトルが付けられており、16,784 のストーリーに「光」という単語が 1 ストーリーあたり平均 3.2 回の割合で登場します。
この例では、パパ

この記事は、「ストーリーを書いてください」というプロンプトに応じて、Google Gemini 3.1 Flash-Lite によって書かれたと述べています。
この研究の著者らは、派生したすべてのキーワードと名前にわたって、懐かしい、あるいは隔世の感のある傾向を特定していることは注目に値します。
繰り返される「灯台」の物語が、フィクションへの通常の接触によって説明できるかどうかをテストするために、モデルのお気に入りの繰り返しの単語といくつかの大きな英語コーパスとの間で比較が行われました。現代小説は、2007 年から 2021 年の間に出版された 2,700 冊の英語小説を含むデータセットである CONLIT を通じて調査されました。このデータセットは 12 のジャンルをカバーし、合計約 2 億 8,700 万語に達します。
「エリアス」は、出版されたフィクションよりも生成された物語に約 900 倍頻繁に登場しました。 Reddit の /r/writingprompts コミュニティのアマチュア フィクションでも同様の頻度が発生しており、このパターンが人間の物語を語る広範な習慣を反映していないことを示しています。
トレーニング前のデータを調べたときにも同じパターンが当てはまりました。 Common Crawl から部分的に抽出された主に人間が書いた約 38 億 9 千万の文書を含む、公開されている OLMo 3 コーパスを使用して、研究者らは、繰り返し現れる「コア」単語がほとんど出現しないことを発見しました。
OLMo 3 コーパスの大部分はノンフィクションであるため、GPT-OSS 20b アノテーションと 200,000 のバランスの取れたサンプルでトレーニングされた FastText モデルを使用してフィクション分類器が構築されました。フィクションの素材に特化してフィルタリングした後でも、「エララ」などの単語は、AI が生成したストーリーと比較して無視できるほどの割合で出現しました。では、なぜ、LLM がフィクションを書くという必須条件の最低レベルでそれらが支配的なのでしょうか?
「コアワードがウェブデータで一般的ではない場合、残りの 1 つのソースはトレーニング後のデータになります。しかし、OLMo のトレーニング後のデータは、私たちのトークンを次のように示していることがわかりました。

CONLITより低いレート。
OLMo 3 のトレーニング後のデータセットからの 78,958 のストーリー内で、「エリアス」は 100 万単語あたり 52.7 回出現し、CONLIT では 2.7 回出現しましたが、研究で調べた生成されたストーリーでは 100 万単語あたり 2,428 回の出現に達したと彼らは指摘しています。
繰り返し発生する「コア」ストーリーがどこから来たのかを特定するために、OLMo 3 のトレーニング後データの各ストーリーは、1 つ以上のコア トークンの存在 (つまり、 Elara 、 Mara などの存在) についてスコア付けされました。 WildChat と関連ソースが OLMo 3 に 59,266 件のストーリーを提供したため、そのほとんどは教師あり微調整 (SFT) データセットに表示されると予想されていました。
ただし、コア用語が含まれるのは 1,803 のみであり、DPO と強化学習に使用されるデータセットではより高い濃度が示されました。
全体として、繰り返されるコア語彙はわずか 3,053 のストーリーまで追跡され、調査されたすべてのトレーニング後のストーリーの 3.8% に相当しました。このような小さなコーパスのサブセットが、実証された方法で最終的にコーパスを支配する統計的な可能性はありません。
「ほとんど方向性を与えられない場合、現在のフロンティアモデルは、名前、場所、職業の狭いカタログを使用して物語を書きます。これらの物語に繰り返し登場する登場人物には、灯台守のエリアスが含まれます。エリアスは珍しい。この名前は、文献、Web データ、さらにはトレーニング後のデータでも一般的ではありません。
著者が特定した上位 11 の単語を取り上げた単一の文学作品 (またはシリーズ) が存在しないため、この特定の単語のコレクションが (トレーニング データとアプローチの多様性にもかかわらず) 複数の大規模な言語モデルの最下位レベルにどのような手段で蓄積され、自己関連付けされたのかはまったく明らかではありません。
たとえ著作権フィルターの抑制効果についての研究者の主張が正しいとしても、古典的な著作権の正真正銘の海は、

トレーニング データの反復により、この奇妙な昔ながらの単語の集まりが、修飾されていない「書き込み」プロンプトの出力を支配することは防げたはずです。
しかし、この理論は、膨大な量の古典文学がトレーニング計画に含まれていたであろうことを前提としています。求められているのは偽ディケンズの外出をノックアウトするモデルではなく、むしろ現代の用語集を扱い、現在のビジネスニーズに適したモデルであるため、その可能性は低いです。産業化以前の文献であっても膨大な量なので、これを含めるのは不可能でしょう。
いずれにせよ、著者が指摘する「強迫観念」の側面が交互に組み合わされたものを特徴とする明確な物語が 1 つあるとしたら、おそらく、それを見つけるのはより簡単になるでしょう。著者自身もそれを見つけることができず、AI 以前の時代について何気なく検索しても、そのような候補者は見つかりませんでした。おそらく、「灯台症候群」が AI のダッシュと同じくらいの悪名を獲得すれば、学者の権威が答えを提示してくれるでしょう。
* 5 月の記事については、読めば明らかになるかもしれない理由から、これ以上立ち入ることはできません。
2026 年 5 月 27 日水曜日に初めて公開されました。Anthropic リンクを修正するために最初の 30 分が修正されました。
マーティン・アンダーソン
機械学習に関するライター、人物画像合成の分野専門家。 Metaphysic.ai の元研究コンテンツ責任者。
個人サイト: martinanderson.ai
連絡先: [email protected]
Twitter: @manders_ai
Web スクレイピングされた AI データセットとプライバシー

[切り捨てられた]

## Original Extract

Asked to 'write a story', ChatGPT and other leading language models appear to be avoiding copyright infringement by obsessive recourse to the same small and strange cast of lighthouse-keepers, fishermen and clockmakers....

Why Does AI Love Writing About Lighthouse Keepers? – Unite.AI
AI Tools
Business
Artificial General Intelligence
AI Tools
Generative AI
Business Plans
Certifications
Blockchain Certifications
Other
Artificial General Intelligence
Why Does AI Love Writing About Lighthouse Keepers?
Asked to ‘write a story’, ChatGPT and other leading language models appear to be avoiding copyright infringement by obsessive recourse to the same small and strange cast of lighthouse-keepers, fishermen and clockmakers.
A new study from Cornell University has found that leading language models seem to have a strange obsession with a very narrow selection of narrative elements, when you ask the model to simply ‘write a story’. After prompting four LLMs to write 20,000 stories, they found that 88% of the stories produced featured at least one of 11 very specific tokens, in the category of ‘location’ , ‘name’ , or ‘profession’ :
The occurrences of unlikely keywords, represented here in parts per million, obtained by the researchers’ analysis of 20,000 LLM-generated stories. Source
The 11 most re-occurring words in the 12+ million words generated by LLMs for the study were the names elias , mara , elara ; the professions keeper , baker , mayor , clockmaker , fisherman , librarian , and conductor ; and the location lighthouse :
The models tested were Claude Haiku 4.5 , Gemini 3.1 Flash-Lite , GPT-5.4-Mini , and OLMo 7b Thinking . All were prompted with one of five requests: ‘Write a story’ ; ‘Please write a story’ ; ‘Write me a story’ ; ‘Tell me a story’ ; or ‘Please tell a story’ .
Curious to see if the syndrome the paper identifies is present in models available at the time of writing, I tried out the experiment myself, first on my customary medium-tier ChatGPT account (link to conversation here ). No cherry-picking was necessary – ChatGPT-5.5 went straight for the material the researchers predicted, on the first try:
ChatGPT-5.5 immediately backs up the paper’s initial findings. Source
Wondering if historic context, or even possible cross-domain leakage might be accounting for this ‘instant hit’, I logged into a free ChatGPT account I have not used in a year or more, in a Firefox private browsing window, and tried again (link to conversation here ). Once again (assuming that OpenAI does not use a common IP address to cross-populate different accounts), ChatGPT hit it out of the park:
ChatGPT account #2 follows the same obsessions and tiny playbook of names and themes outlined in the new paper. ‘Mira’ is in the authors’ top 20. Source
It’s worth noting that these GPT versions were a grade up from the 5.4 tested for the paper.
Though Claude Haiku was tested for the paper, I tried Anthropic’s default Sonnet 4.6, and was not disappointed. Once again, the familiar keywords came at the first try (link to conversation here ):
This time ‘Mara’, another stalwart from the ‘top 11’, leads the story, in the first attempt on Claude Sonnet 4.6 . Source
Trying the same prompt on Claude Haiku 4.5 led to pretty much the same result .
I was unable to reproduce the authors’ findings at Google Gemini at first, until I specifically changed the model to the one used in the paper, Gemini 3.1 Flash-Lite – and then, on that third try (but first with that model), the pattern emerged immediately (link here ):
Google Gemini 3.1 Flash-Lite . Source
Further experiments with different Gemini models invariably turned up the lighthouse theme, though with variants not featured in the ‘top 11’, such as the name ‘Thomas’, and, in another variant, my own name, as the protagonist.
Nonetheless, at the time of writing, the paper’s findings are extremely easy to prove.
Great minds think alike: a week ago, prior to the publication of the new paper, software writer Daniel May pointed out the coincidence of the Elias and Lighthouse keeper trope extracted by the researchers*, apparently having noticed it at random. He went on to test eight variants of Gemini, DeepSeek, Qwen and Gemma, which he found would produce the lighthouse memes and ‘Elias Thorne’ as a protagonist*. However, this initial discovery did not extend to the wider range of persistent content themes outlined in the new paper.
Curious to see if these recurrent themes, names and locations had ever escaped the confines of a chat, I searched for some of the top 11 keywords and themes on Google, and found a remarkable number of posts that seem to have channeled them:
Three examples of the meme in output. See below for source links.
May had identified the longer Elias Thorne (rather than just ‘Elias’) as a persistent LLM meme, and posted various screenshots from Amazon, where this name has apparently been used as the title for the author/s of diverse books, including medical books.
Instead, I sought and found content that appeared to have invoked the persistent themes from an LLM, including an X post of a story (archive version here ); a fictional work (archive version here ); and a story with narration on YouTube (archived here ). There was a great deal more to traverse, but time did not permit it.
So much for casual observation and serendipity. While no single ‘magic document’ in training data has yet turned up that features all or most of the persistences, the authors of the new paper (titled Elias in the Lighthouse, Again? Diagnosing Low Diversity in LLM Stories , from two researchers at Cornell University) theorize that copyright filters in AI developments may be restricting fictional output in LLMs to material that is out of copyright.
‘We find that the dominance of “Elias in the Lighthouse” stories cannot be explained by prevalence in pre- or post-training data. We speculate that models are trained to avoid references to copyrighted characters and adult content during alignment but defer this question to future work.’
Comparison table showing how often recurring words from AI-generated stories appeared across published literature, web fiction and post-training datasets, with terms such as ‘Elias’ and ‘lighthouse’ occurring far more frequently in chatbot-written fiction.
In the study, the authors found that the emphasized 11 words occur in 88% of the 20,000 stories generated, and that there is ‘little difference between models’. They stress that these words are uncommon in published English literature, and that post-training data (data designed to condition and align models into ‘acceptable’ usage) could well be responsible.
‘A typical example shown [below] highlights three elements common across nearly all 20,000 stories: a location (19,864 stories), a character name (19,864 stories), and a profession (15,807 stories).
‘In fact, the specific location (“lighthouse”), name (“Elias”), and profession (“keeper”) in this story appear in some combination across 66.6% of all generated stories. Light is likewise a common theme: 56% of stories generated by Claude are titled “The Lighthouse Keeper’s Secret” and the word “light” appears in 16,784 stories at an average rate of 3.2 instances per story.’
This example, the paper states, was written by Google Gemini 3.1 Flash-Lite, in response to the prompt ‘Write a story’.
It’s worth noting that the authors of the study identify a nostalgic or atavistic trend across all the derived keywords and names.
To test whether the repeated ‘lighthouse’ stories can be explained by ordinary exposure to fiction, comparisons were made between the models’ favorite recurring words and several large English-language corpora. Contemporary fiction was examined through CONLIT , a dataset containing 2,700 English novels published between 2007 and 2021, covering 12 genres and totaling roughly 287 million words.
‘Elias’ appeared around 900 times more often in the generated stories than in published fiction . Amateur fiction from Reddit’s /r/writingprompts community produced similar frequencies, indicating that the pattern does not reflect broader human storytelling habits either.
The same pattern held when pre-training data was examined. Using the openly available OLMo 3 corpus , which contains roughly 3.89 billion primarily human-written documents drawn partly from Common Crawl , the researchers found that the recurring ‘Core’ words barely appeared at all .
Since much of the OLMo 3 corpus is non-fiction, a fiction classifier was built using GPT-OSS 20b annotations and a FastText model trained on 200,000 balanced samples. Even after filtering specifically for fictional material, words such as ‘Elara’ still appeared at negligible rates compared to the AI-generated stories. Why, therefore, do they dominate at the lowest level of the imperative for an LLM to write fiction?
‘If Core words are not common in web data, then one remaining source would be post-training data. But we find that OLMo’s post-training data exhibits our tokens at a lower rate than CONLIT.
Within 78,958 stories from OLMo 3’s post-training datasets, they note, ‘Elias’ appeared 52.7 times per million words, compared to 2.7 in CONLIT, but reached 2,428 occurrences per million words in the generated stories examined in the study.
To identify where the recurring ‘Core’ stories were coming from, each story in OLMo 3’s post-training data was scored for the presence of one or more Core tokens (i.e., for the presence of Elara , Mara , etc.). Most were expected to appear in supervised fine-tuning (SFT) datasets, because WildChat and related sources contributed 59,266 stories to OLMo 3.
However, only 1,803 contained Core terms, while datasets used for DPO and reinforcement learning showed higher concentrations.
Overall, the recurring Core vocabulary was traced to just 3,053 stories, representing 3.8% of all post-training stories examined. There is no statistical possibility for such a small subset of corpora to end up dominating it in the way demonstrated.
‘When given little direction, current frontier models write stories using a narrow catalog of names, places, and occupations. Recurring characters in these stories include Elias, a lighthouse keeper. Elias is unusual; the name is uncommon in literature, web data, and even post-training data.’
In the absence of any single work of literature (or even a series) that features the top 11 words which the authors identify, it is not at all clear by what means this particular collection of words has accreted and self-associated into the very lowest levels of multiple large language models (notwithstanding their diversity of training data and approaches).
Even if the researchers’ contention about the constraining effect of copyright filters is correct, a veritable ocean of classic literature in the training data should have prevented this strange collection of old-school words from dominating the output of a non-qualified ‘write’ prompt.
That theory assumes, however, that vast amounts of classic literature would have been included in the training regimen at all. That’s unlikely, since what’s wanted are not models that will knock out faux Dickens outings, but rather that deal with the modern lexicon, and are suited for current business needs. The sheer volume even of pre-industrial literature would preclude its inclusion.
In any case, if there were one distinct narrative featuring some alternating mix of the ‘obsessive’ facets that the authors note, it would, presumably, be easier to find; the authors themselves could not find it, and casual searches on the pre-AI era unearth no such contender. Perhaps, if ‘lighthouse syndrome’ gains the same notoriety as AI em dashes , some scholarly authority will come forward with the answer.
* I can’t go any further into May’s article, for reasons that may become obvious when one reads it.
First published Wednesday, May 27, 2026. Modified in first 30 minutes to fix Anthropic link.
Martin Anderson
Writer on machine learning, domain specialist in human image synthesis. Former head of research content at Metaphysic.ai.
Personal site: martinanderson.ai
Contact: [email protected]
Twitter: @manders_ai
Web-Scraped AI Datasets and Priva

[truncated]
