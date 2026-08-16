---
source: "https://daringfireball.net/2026/08/anthropics_watermark_text_adulteration_in_claude_is_a_perversion_of_writing"
hn_url: "https://news.ycombinator.com/item?id=49324087"
title: "Anthropic's 'Watermark' Text Adulteration in Claude Is a Perversion of Writing"
article_title: "Daring Fireball: Anthropic’s ‘Watermark’ Text Adulteration in Claude Is a Perversion of Writing"
author: "ropbear"
captured_at: "2026-08-16T22:11:03Z"
capture_tool: "hn-digest"
hn_id: 49324087
score: 7
comments: 2
posted_at: "2026-08-16T21:53:43Z"
tags:
  - hacker-news
  - translated
---

# Anthropic's 'Watermark' Text Adulteration in Claude Is a Perversion of Writing

- HN: [49324087](https://news.ycombinator.com/item?id=49324087)
- Source: [daringfireball.net](https://daringfireball.net/2026/08/anthropics_watermark_text_adulteration_in_claude_is_a_perversion_of_writing)
- Score: 7
- Comments: 2
- Posted: 2026-08-16T21:53:43Z

## Translation

タイトル: アントロピックのクロードにおける「透かし」テキストの改ざんは執筆の倒錯である
記事のタイトル: Daring Fireball: Anthropic の「ウォーターマーク」クロードにおけるテキストの改ざんはライティングの倒錯である
説明: 出所を示唆するためにテキスト内に隠された手がかりを埋め込む目的で、ツールが明瞭さ、一貫性、意味、品質などを犠牲にすることは容認できません。 「私」のニーズ以外のものはすべて「私のため」のテキストの生成に考慮されるべきだという考えは、明らかに不快です。

記事本文:
Drata の Agentic Trust Management プラットフォームで GRC をより迅速に管理
アントロピックのクロードにおける「透かし」テキストの改ざんは執筆の倒錯である
私が今週、この EU 規制に準拠するために、世界中のすべてのクロード モデルが、テキストを含む生成するすべてのものに間もなく「透かし」を開始するという Anthropic の発表について書いたとき、発表のタイトルが不条理かつ侮辱的に「クロードが AI 生成コンテンツにどのようにマークするか」という事実にもかかわらず、Anthropic はそれがどのように機能するかについての曖昧な説明さえ提供しなかったため、これがどのように機能するのか推測するしかありませんでした。
私の最初の推測は、テキスト内の非表示の非印刷 Unicode 文字を隠すのではないかということでした。ただ吐き出すだけ。彼らがやろうとしているのはそれではないことが判明した。彼らがやろうとしていることは、ステガノグラフィーの形式を適用することです。そこでは、推論時の単語 (または他のトークン出力) の選択により、後でおそらく確率的に検出できる指紋が残ります。
私が最初に「目に見えないキャラクター」と推測したのは、意味論的な単語選択テクニックを考えていなかったからではなく、彼らが何をするのかについての説明において、人間論をその言葉のままに受け取った私が愚かだったからです。元のサポート文書では次のように主張されています。
サポートされているクロード モデルがテキストを生成すると、
目に見えない透かしをテキスト自体に直接挿入します。そうはなりません
見ても、意味、品質、読みやすさは変わりません
クロードの返答。
彼らは「知覚できない」、「意味、品質、読みやすさを変えない」と言います。彼らの言葉。ほとんど知覚できないわけではありません。意味、品質、読みやすさは少しも変わりません。それは私にとって当然のことでした。なぜなら、それが私が個人的に使用するツールに望んでいること、いや、要求していることだからです。

出所を示唆するためにテキスト内に隠された手がかりを埋め込む目的で、ツールが明瞭さ、一貫性、意味、品質などを犠牲にすることは容認できません。それが私が望んでいることですし、これからも要求することです。そして、Anthropic の (オリジナルの) サポート文書では、それが彼らのシステムによって可能になると明確に主張しています。それが本当だとすると、テキスト内の目に見えない文字を隠す以外に何が残っているのかを見ることができませんでした。
私の間違いは、Anthropic 社のシステムが、モデルが生成するテキストのセマンティクスを改ざんしたり破損したりすることはないと信じていたことでした。実際、それがまさに彼らが計画していることなのです。クロードが AI 生成コンテンツをどのようにマークするか (またはマークする予定か) まったく説明していない「クロードが AI 生成コンテンツをどのようにマークするか」というタイトルの文書の 1 語を信じたかどうか、私の頭を調べてもらう必要があります。
実際にどのように機能するか
昨日、元の記事「AI 生成コンテンツをクロードがどのようにマークするか」 (その機能についてまったく説明していない記事) とはまったく別の Web サイトで、Anthropic は「クロードのテキスト ウォーターマークの仕組み」を公開しました。この記事では、実際にどのように機能するかを素人でもアクセスできる言葉で説明しています。以下で、Anthropic の非常に婉曲的でやや誤解を招く新しい説明に戻ります。
このトピックに関しては多くの研究があり、その一部を以下にリンクしています。しかし、この技術の背後にある一般的な考え方を最もよく説明しているのは、James Padolsey によるインタラクティブなエッセイ「How AI Text Watermarking Works」です。素晴らしく説得力のある読み物であり、インタラクティブな要素が主要なコンセプトを見事に説明しています。 A+の作品。これに少しでも興味があるなら、ぜひパドルシーの作品を読んで、遊んでみてください。
しかし、これが素人向けの私の要約です

やあ。コインを N 回投げて結果に注目すると、そのコインが公平か偏っているかをある程度の確実性で判断できます。 LLM は、一般的には非決定的です。同じモデルについて同じ質問をすると、少なくともわずかに異なる答えが得られることがよくあります。もしかしたら同じ意味かもしれませんが、言い方が違うかもしれません。次のトークンを生成するための各決定点で、モデルは選択を行います。これらの意味論的な透かし技術を使用すると、「緑」と「赤」と呼ばれる単語リストに基づいて、一部のトークンに対して異なる選択が行われます。それぞれの決定時点で、赤のリストよりも緑のリストから単語を選択する可能性が少し高くなります。だからといって、彼らがレッドリストから単語を決して選ばないというわけではありません。ただ、異物混入のマーキング技術が導入されていない場合よりもその可能性が低くなるというだけです。 (曲がった 51-49 コインでも、平均 100 回中 49 回は「間違った」面が表れるのと同じです。)
単語または単語フレーズは、各「次のトークン」生成ポイントで、オンザフライで決定的に緑と赤のリストに分類されます。したがって、特定の単語がグリーン リストに含まれる場合もあれば、レッド リストに含まれる場合もあります。秘密鍵を持っている人は、各トークン生成ポイントで単語がどのリストに含まれるかを決定できます (これにより透かしが検出されます)。秘密鍵を持たない人はできません。これは、クロードが好む言葉や避ける言葉のリストが決して存在しないことを意味します。
コイン投げでは、N が高くなるほど、つまり投げる回数が増えるほど、コインが公平であるか、偏っているか確信できます。このセマンティック透かしも同様です。テキスト内の単語が多いほど、テキストが特定の AI モデルによって生成されたかどうかの分析がより正確になります。コイン投げの回数が少なすぎると、

コインの公平性に関してはまったく自信がありません。単語 (またはトークン) が少なすぎると、テキスト文字列が AI によって生成されたものであるかどうかを確信する方法がありません。
特定の透かし入れシステムの兆候を調べるためのテキスト文字列が与えられた場合、緑としてタグ付けされた単語が予想よりも多く、赤としてタグ付けされた単語が少ない場合、そのテキストは、特定の秘密キー透かし入れシステムを適用する AI システムによって生成されたか、単に変更されたものとして、ある程度の自信を持ってフラグを付けることができます。決定の信頼度は、テキスト文字列のサイズと、グリーン リストとレッド リストの単語に与えられるランダムな重みに基づいて、明らかに大きく異なります。ただし、テキストがクロードによって生成されたと思われるかどうかを判断できるのは Anthropic だけであり、Anthropic が検出できるのはクロードによって適用されたウォーターマークのみです。各実装は LLM プロバイダーのみが保持する秘密鍵を前提としているため、Claude は、たとえば Gemini によって生成された隠し透かし信号を検出できません。また、Gemini も Claude によって作成された隠し透かし信号を検出できません。
技術的前提に対する異議
これに関する私の根本的な問題の 1 つは、まったく同じ意味をもつ同義語が 2 つもないということです。 「彼はチャンスに飛びついた」と「彼は機会に飛びついた」は、同じ一般的な感情を表現する非常に似た文ですが、同じではありません。文章を書くときに選ぶ正確な言葉が重要です。私は、私が使用する LLM には、あらゆる意思決定時点で最高かつ最も正確な言葉を選択してもらいたいと考えています。私が受け入れている明らかな制約は、時間と計算です。トークンごとに一定のコストで、推論を迅速に実行するという制約の中で、最良の言葉が必要です。この制約は人間の書き込みと一致します。きっとできるよ

書く時間を長くすることで、より良いコラムを書くことができます。私は、すべての単語や句読点の選択にどれだけ注意を払う必要があるかを意識して書きます。直感がそうすべきだと思うときは、特定の段落、文、さらには個々の単語の選択にもっと時間をかけます。
言い換えれば、これらは必要なトレードオフです。スピード、コスト、品質といった要素はすべて私にとって興味のあるものです。理想的には、ゼロコストで、瞬間的な生成速度で、完璧な書き込みがしたいと考えています。そういったことはどれも不可能です。計算は無料ではありません (そして、主要なモデルを使用したクラウドベースの LLM 推論は実際には高価です)。推論は瞬時には行われません。そして、自然であれ人工であれ、優れた文章は完璧に近づくことしかできません。
私のニーズ以外のものもテキストの生成に考慮されるべきだという考えは、明らかに不快です。
これは、人が自分の自然な作品として宣伝する目的で作成したテキストに限った話ではありません。これは、手書きの作品の LLM 校正の話ではありません。 Anthropic 社は、すべての新しいクロード モデルは、生成する 200 トークン (約 150 ワード) を超えるテキストのあらゆる部分を改ざんする予定であると述べています。これには、ユーザーが読むために提示するすべてのものが含まれます。そのため、ユーザーとクロードの間のプライベートな会話であっても、ユーザー以外の誰も読むことはありませんが、クロードは、明瞭さと精度を最大化するのではなく、統計的に予測可能な方法でその出力にマークを付けるという名目で単語の選択を開始します。
今日のいわゆるフロンティアモデルでさえ、すでに明らかに明快さを欠いています。クロード、ChatGPT、Grok、他。彼らはほとんどの人間よりも「優れた作家」であり、平均的な人間よりも優れた散文を生み出します。でも、そんなことはない。ほとんどの人はひどい作家です。 「普通の人」はかなり愚かで、全人類の半分は

私たちはそれよりも愚かです。そして、賢くて面白い人でも、残念な作家はたくさんいます。したがって、LLM は素晴らしいものですが、ハードルは低いです。これらのモデルから得られる最高の文章は、私が楽しみのために読むものよりもひどいものです。そして今、Anthropic は、私に何の利益にもならない目的で、意図的に状況を悪化させると言っているのでしょうか？たとえ少しだけ悪くても？
EU規制に対する異議
反対意見といえば、これらすべての動機となっている関連する EU の規制である「AI 生成コンテンツの透明性に関する実践規範」は、官僚的な乳母国家の空想上のナンセンスです。今週のペイウォール Stratechery アップデートからの Ben Thompson の要約は次のとおりです。
この規制は、200 トークンを超えるテキストに適用されます。
プロバイダーは、ユーザーが次のことを利用規約で義務付ける必要があります。
透かしを削除しないでください。
ソリューションは、「典型的な」問題を回避するという点で堅牢である必要があります。
スクリーンショット、スキャン、OCR などの「処理ソリューション」
コピーアンドペースト、翻訳など。
文字通りに解釈すると、準拠した LLM 利用規約では、単語の選択が目印となるため、この規制に準拠したモデルからの出力をユーザーが言い換えることを禁止する必要があります。しかし、不条理で非現実的で魔女狩りを煽る規制を全世界に課そうとしているのは欧州連合ではない。それはAnthropicに当てはまります。
特に text に関してこれに従うと、正直なユーザーにとって問題が生じるだけです。 AI が生成したテキストを自分の文章として偽装しようとする不正なユーザー (学生、従業員、誰でも) は、非準拠の AI 言い換えツールを使用して検出を回避するだけです。
James Padolsey — これらのスキームがどのように機能するかをインタラクティブに視覚的に説明し、上にリンクしました — は、「Anthropic’s W」というタイトルの投稿でこれを説明しています

eak ウォーターマークは弱い法則を緩和する 」 (これは、今日初めに独立した投稿でリンクしたものです):
この法律を導いたのと同じ考えが、
計算機はその誕生時に出力を備えていました
工芸品を通して自分自身を明らかにしました。ありがたいことに、支払われた金額は、
脳は、人間によって作られるものと何ら変わりなく扱われます。
計算機。スペルチェッカーも同様です。援助をするため
ツールが十分な機能を構築できるようになって初めて疑う。
文全体が原則的な境界ではありません。それは道徳的プレミアムです
難易度そのものに設定されています。
それにもかかわらず、Anthropic はモデルレベルのブランケットを選択しました
法律の最低限度よりも広範囲に見える実装
要件。それは便利なコンプライアンスエンジニアリングかもしれませんが、
法律が明示的に保持しようとした区別を破棄する。の
その結果、無害であることを示すのに十分な広さの信号が得られます。
補助的な使用ですが、動機のある人によって取り外されるのに十分なほど壊れやすい
大幅な再構成を通じて人物を表現します。集中する危険性がある
一般利用者や補助利用者に対する疑惑は依然として最も弱いままである
意図的な欺瞞に対して。
Padolsey は、「AI 風味のテキストを貼り付けると、同じコンテンツを単純な散文として取得する」ことができる、非常にシンプルな Web アプリである Declaude の作成者です。デクロードの本来の目的は、テキストから甘ったるいクロードの性格の悪臭を取り除くことです (それがクロードによって作成されたものであるか、他の LLM によって作成されたものであるかにかかわらず)

[切り捨てられた]

## Original Extract

It’s unacceptable for a tool to sacrifice an iota of clarity, coherence, meaning, quality, etc. for the purpose of embedding hidden clues within the text to suggest its provenance. The idea that anything other than *my* needs should factor into the generation of text *for me* is patently offensive.

Manage GRC Faster with Drata’s Agentic Trust Management Platform
Anthropic’s ‘Watermark’ Text Adulteration in Claude Is a Perversion of Writing
When I wrote this week about Anthropic’s announcement that all Claude models, worldwide, would soon begin “watermarking” everything they generate, including text, to comply with this EU regulation , we were left to speculate how this was going to work, because Anthropic offered not even a vague description of how it would work — despite the fact that the title of the announcement was, absurdly and insultingly, “ How Claude Marks AI-Generated Content ”.
My initial speculation was that maybe they’d hide invisible non-printing Unicode characters in the text. Just spitballing. Turns out that’s not what they’re going to do. What they’re going to do is apply a form of steganography, where the choice of words (or other token output) at inference time will leave fingerprints that can later, maybe, be detected probabilistically.
I initially guessed “invisible characters” not because I didn’t think of the semantic word-choice technique, but because I was a fool who took Anthropic at its word in their description of what they would do. Their original support document claims:
When a supported Claude model generates text, it weaves an
imperceptible watermark directly into the text itself. You won’t
see it, and it doesn’t change the meaning, quality, or readability
of Claude’s response.
They say “imperceptible” and “doesn’t change the meaning, quality, or readability”. Their words. Not almost imperceptible. Not slightly changes the meaning, quality, or readability. That made sense to me, because that’s absolutely what I want — nay, demand — from any tools I use personally. It’s unacceptable for a tool to sacrifice an iota of clarity, coherence, meaning, quality, etc. for the purpose of embedding hidden clues within the text to suggest its provenance. That’s what I would and will demand. And Anthropic’s (original) support document unambiguously claims that’s what their system will enable. So if that were true, I couldn’t see what was left other than hiding invisible characters within the text.
My error was believing Anthropic that their system wouldn’t adulterate and corrupt the semantics of the text their models generate. That is in fact exactly what they plan to do. I should have my head examined for believing a single word of a document titled “How Claude Marks AI-Generated Content” that doesn’t explain, at all, how Claude marks (or will mark) AI-generated content.
How It’s Actually Going to Work
Yesterday, on an entirely different website than the original “How Claude marks AI-generated content” article (the one that didn’t explain anything at all about it works), Anthropic published “ How Claude’s Text Watermark Works ”, which does actually explain in layman-accessible terms how it’s going to work. I will return to Anthropic’s new highly euphemistic and slightly misleading description below.
There’s a bunch of research on this topic, some of which I have also linked to below. But the very best description of the general idea behind the technique is an interactive essay by James Padolsey, “ How AI Text Watermarking Works” . It’s a wonderfully cogent read, and the interactive elements splendidly illustrate the main concepts. A+ work. If you have any interest in this at all, I dare say you must read — and play with — Padolsey’s piece.
But here’s my stab at a layman’s high-level summary. If you toss a coin N times and note the results, you can determine with a degree of certainty whether the coin is fair or biased. LLMs are, in their popular incarnations, non-deterministic. Ask the same question of the same model and you often get at least slightly different answers. Maybe the same meaning, but different phrasing. At each decision point for generating the next token, the model makes a choice. With these semantic watermarking techniques, they make different choices for some tokens based on word lists that could be called “green” and “red”. At each decision point, they’re a little more likely to pick a word from the green list than the red list. That doesn’t mean they never choose words from the red list. Just that they’re less likely to than they would if the adulterated marking technique weren’t in place. (Same way that a crooked 51-49 coin will still land “wrong” side up 49 times out of 100 on average.)
Words or word phrases are sorted into the green and red lists deterministically on the fly, at each “next token” generation point. So sometimes a specific word will be on the green list, and other times it will be on the red list. Someone with the secret key can determine which list a word will be on at each token generation point (which is how the watermarking is detected); those without the secret key cannot. This means there will never be a list of words that Claude prefers or eschews.
With coin flipping, the higher N is — the more times you flip — the more confident you can be that the coin is fair or biased. So too with this semantic watermarking. The more words in the text, the more accurate the analysis will be that the text was generated by a specific AI model or not. With too few coin flips, you can’t achieve any confidence at all regarding a coin’s fairness. With too few words (or tokens), there’s no way to achieve any confidence whether a string of text was AI-generated or not.
Given a string of text to examine for signs of a specific watermarking system, if there are more words tagged as green and fewer tagged as red than would otherwise be expected, the text can be flagged — with some degree of confidence — as having been generated, or merely modified, by the AI system that applies the specific secret-key watermarking system. The amount of confidence in the determination will obviously vary, significantly, based on the size of the text string and randomized weights given to words on the green and red lists. But only Anthropic will be able to determine if text was seemingly generated by Claude, and Anthropic will only be able to detect the watermarks that are applied by Claude. Claude can’t detect the hidden watermark signals generated by, say, Gemini, and Gemini can’t detect the hidden watermark signals created by Claude, because each implementation is predicated on secret keys held only by the LLM provider.
Objections to the Technical Premise
One of my fundamental problem with this is that no two synonyms carry the exact same meaning. “ He leaped at the chance ” and “ He jumped at the opportunity ” are very similar sentences expressing the same general sentiment, but they are not the same. The exact words we choose when writing matter. I want any LLM I use to choose the very best, most precise words at every single decision point. An obvious constraint that I accept is time and computation. Within the constraint of executing inference quickly, and at a certain cost per token, I want the best words. This constraint matches human writing. I could surely write a better column by taking longer to write it. I write with a sense of how much care I should put into every word and punctuation choice I make. I take more time with certain paragraphs, sentences, or even individual word choices when my gut feeling says I should.
In other words, these are necessary trade-offs. These factors are all in my interest: speed, cost, quality. Ideally I would like perfect writing, at instantaneous generation speed, at zero cost. None of those things are possible. Computation is not free of charge (and cloud-based LLM inference with leading models is actually expensive). Inference is not instantaneous. And great writing, whether natural or artificial, can only approach perfection.
The idea that anything other than my needs should factor into the generation of text for me is patently offensive.
This isn’t just about text one might generate with the intention of passing it off as their own natural work. This isn’t even about LLM proofreading of work written by hand. Anthropic is saying that all new Claude models are going to adulterate every single bit of text longer than 200 tokens (~150 words) they generate, including everything it presents to its users to read. So even in a private conversation between a user and Claude, which will never be read by anyone other than the user, Claude will begin making word choices in the name of marking its output in statistically predictable ways rather than maximizing clarity and precision.
Even today’s so-called frontier models are already decidedly lacking in lucidity . Claude, ChatGPT, Grok, et al. are “better writers” than most humans and produce better prose than the median human. But: no shit. Most people are terrible writers. The “average person” is pretty stupid and half of all people are stupider than that . And there are many smart, interesting people who are miserable writers. So as impressive as LLMs are, the bar is low. The best writing I see come out of these models is worse than anything I would choose to read for pleasure. And now Anthropic is saying they’re going to make it worse, on purpose, for purposes that do not benefit me in any way? Even if only slightly worse?
Objections to the EU Regulation
Speaking of objections, the relevant EU regulation motivating all of this, “ Code of Practice on Transparency of AI-Generated Content ”, is red-tape nanny-state pipe-dream nonsense. Here’s Ben Thompson’s summary from a paywalled Stratechery update this week :
The regulation applies to text longer than 200 tokens.
The provider must mandate in their terms-of-service that users
not remove the watermarking.
The solution should be robust in terms of evading “typical
processing solutions” like screen shots, scanning and OCR,
copy-and-pasting, translations, etc.
Taken literally, compliant LLM terms of service must forbid users from rephrasing the output from models that comply with this regulation, because the word choices are the marks. But it’s not the European Union that is trying to impose their absurd, impractical, witch-hunt-fueling regulation on the entire world. That falls on Anthropic.
Complying with this, particularly with regard to text , is only going to create problems for honest users. Dishonest users attempting to pass off AI-generated text as their own writing (students, employees, whoever) will simply circumvent detection through non-compliant AI paraphrasing tools.
James Padolsey — whose interactive visual explanation of how these schemes work I linked to above — explains this in a post titled “ Anthropic’s Weak Watermarks Appease a Weak Law ” (which, if it rings a bell, I linked to in a standalone post earlier today):
The same thought that led to this law could have applied to
calculators at the time of their inception, had their outputs
revealed themselves through artefacts. Thankfully, a sum borne of
the brain is treated no differently from one produced by a
calculator. Likewise with spellcheckers. To make assistance
suspect only once the tool becomes capable enough to compose a
whole sentence is not a principled boundary. It is a moral premium
placed on difficulty itself.
Anthropic has nevertheless chosen a blanket, model-level
implementation that appears broader than the law’s minimum
requirement. That may be convenient compliance engineering, but it
discards distinctions the law expressly attempted to preserve. The
result is a signal broad enough to implicate harmless and
assistive use, yet fragile enough to be removed by a motivated
person through substantial recomposition. It risks concentrating
suspicion on ordinary and assistive users while remaining weakest
against deliberate deception.
Padolsey is the creator of Declaude , a delightfully simple web app that allows you to “Paste in AI-flavored text and get the same content back as plain prose”. Declaude’s original purpose is cleaning the saccharine Claude personality stink from text (whether it was created by Claude or any other LLM

[truncated]
