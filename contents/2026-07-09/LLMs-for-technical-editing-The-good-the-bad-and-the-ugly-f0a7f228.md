---
source: "https://techstackups.com/articles/llms-for-technical-editing-the-good-the-bad-and-the-ugly/"
hn_url: "https://news.ycombinator.com/item?id=48845274"
title: "LLMs for technical editing: The good, the bad, and the ugly"
article_title: "LLMs for Technical Editing: The Good, the Bad, and the Ugly | Tech Stackups"
author: "ritzaco"
captured_at: "2026-07-09T13:43:52Z"
capture_tool: "hn-digest"
hn_id: 48845274
score: 2
comments: 0
posted_at: "2026-07-09T13:10:48Z"
tags:
  - hacker-news
  - translated
---

# LLMs for technical editing: The good, the bad, and the ugly

- HN: [48845274](https://news.ycombinator.com/item?id=48845274)
- Source: [techstackups.com](https://techstackups.com/articles/llms-for-technical-editing-the-good-the-bad-and-the-ugly/)
- Score: 2
- Comments: 0
- Posted: 2026-07-09T13:10:48Z

## Translation

タイトル: 技術編集用の LLM: 良いもの、悪いもの、そして醜いもの
記事のタイトル: 技術編集のための LLM: 良いもの、悪いもの、そして醜いもの |技術の積み重ね
説明: LLM: テクニカル エディターのツール、それとも代替品?

記事本文:
メイン コンテンツにスキップ 技術スタックアップ ホーム トピック 比較 ガイド 記事 ニュース 技術編集用の AX LLM: 良い点、悪い点、そして醜い点
Opus 4.8 の存在と、Fable の全世界向けの限定的再リリースにより、ライターや編集者を AI に完全に置き換えることが可能だと考えているかもしれません。
確かにそれは可能ですが、人々に実際にあなたのコンテンツに関心を持ってもらい、あなたのブランドと結びついてもらいたいのであれば、それは最も非効率的で自己破壊的な決断になるでしょう。
そうは言っても、私には偏見があることを認めます。私は編集者であり、企業が AI にこだわるのは逆効果だと考えています。
それにもかかわらず、一部の人々は AI が編集者に取って代わることができると信じています。なぜそれが真実ではないのかを説明します。客観性を保つために、私は AI を使用して、エラーが含まれているすでに公開された記事を編集しました。
この実験が終わるまでに、Claude が 2 つの両極端の間のどこに位置するのかがわかるでしょう。それはエディターを完全に置き換えることができるのでしょうか、それとも単なる派手な (そして場合によっては不正確な) オートコンプリートなのでしょうか。
公開された AX 記事をシードしました
さまざまな重大度の 23 個のエラーがあります:
次に、Opus 4.8 と Fable の両方に、からのプロンプトを使用して、エラーのある記事を評価するように依頼しました。
私たちの編集プロンプトライブラリ。
注意: 私はクロードに、修正を提案する前にまずエラーを特定するように依頼しました。
(また、Opus 4.8 が AX 記事のクリーン バージョンの編集にどのように取り組んだかを知るために、この記事を読むことをお勧めします。)
良い点: 作品をまとめることができる
さて、あまり言いたくないのですが、クロードは構造的および論理的エラーにフラグを立てるのが本当に上手でした。見出しの不一致、内容の矛盾、または情報の欠落があった場合、Opus と Fable は両方とも、これらのエラーにフラグを立て、適切な修正を提案してくれました。
1.

枠組みの矛盾を捉えた
この記事では、発見 (エージェントはあなたのことを知っていますか?)、オンボーディング (サインアップできますか?)、および使用という 3 つのハードルを定義しています。最初の 2 つのラベルを交換するこの文をシードしました。
エージェントはハンドラーからの最小限の助けを借りてサインアップしますか (検出?) エージェントはあなたのことを知っていますか (オンボーディング)
これは、記事自体が 2 つのセクションで設定したフレームワーク (ディスカバリは「エージェントが解決策を提供することをエージェントに認識させる必要がある」) であり、オンボーディングはサインアップ フローであるというフレームワークと直接矛盾します。読者が読み終えることを意図した概念的なポイントが 1 つありますが、それは逆向きに述べられています。
これを理解するには、前の 2 つのセクションの定義を念頭に置き、後の文をそれらと照合する必要があります。単純なパターン マッチング ツールではこれを行うことはできません。
2. 自身のセクションに反論する見出しを見つけました
見出しを反転して「サインアップは通常、DX よりも AX に関するものです」と表示しました。本文のすぐ上に、サインアップ ページにアクセスするのは依然として人間であることを説明しています。
人間がサインアップを行う場合、それは開発者の経験がテストされることになります。
エージェントのものではありません。見出しには、サインアップは依然として AX よりも DX に関するものであると記載する必要があります。
3. 独自の例によって反論された主張を捉えた
この文の主語を入れ替えたので、例は主張の反対を証明しています。
人間は通常、エージェントよりもはるかに長い Web 検索を実行することに注意してください。人間なら Steel キャプチャのようなものを検索するでしょうが、Claude は Steel.dev がキャプチャ セッション設定 Python SDK 2025 を解決します。
人間のクエリはもっと​​短いです。エージェントの方が長いです。この例は次のことを証明しています
主張に反する。
4. 1 回のパスで 30 分間の整合性チェックを実行しました
どちらのモデルも、英国の「最適化」という 1 つのプロンプトで機械的な不一致について問題を一掃しました。

米国英語の記事では、数十の直線アポストロフィの中に 1 つの中折れアポストロフィがあり、全角ダッシュを使用する文書内のダッシュの間隔、「ヘッドレス」と「ヘッドレス」、および大文字の「CAPTCHA」と小文字の「キャプチャ」が使用されています。
クロードが最も苦労したのは、タイプミスや単純な文法上の間違いを見つけることであったと知ると驚くかもしれません。
1. どちらのモデルもエラーを直視し、間違ったものにフラグを立てました
この文では「it's」を「they're」に変更し、代名詞の同意を破りました。
Skyscanner にはよりシンプルな CAPTCHA がありますが、Steel の自動解決には含まれていないため、時間がかかりました。
どちらのモデルもタイプミスのフラグを立てましたが、どちらも代名詞の不一致を発見しませんでした。
248 行目: 「they're」では中巻きアポストロフィーが使用されています。ファイル内の 1 つおきの短縮形では、直線のアポストロフィが使用されます。
この文書は、1 つの中巻きアポストロフィ (248 番目の「they're」) を除いて、全体に直線の引用符/アポストロフィが使用されています (検証済み)。
2. 悪魔は細部に宿る（ではない）
Fable はこれらのほとんどを検出しましたが、Opus は行レベルでのいくつかのエラー、つまり「簡単な」エラーを見逃しました。
Opus のみ (7 月 12 日以降も利用できる 2 つのうちの唯一の 1 つ) のみを実行した場合、最終的に 7 つのエラーが公開されることになります。そして、彼らは両方とも最後の 2 つのエラーを見逃していたため、どちらのエラーでもクリーンな実行は得られません。
The Ugly: 記事を悪化させる自信に満ちたアドバイス
AI が生成した文章を「下手」と呼ぶのには理由があります。 LLM は限られたデータセットからしか抽出できません。つまり、データセットがどれほど大きくても、モデルは最終的に物事を表現するための「独自の」方法を使い果たすことになります。だからこそ、私たちは今、文章を見て、その執筆に AI が関与しているかどうかを判断することが非常に簡単になっています。コンセプト、構造、特に表現には、明らかに明白な兆候があります。
trも同じです

AI に文章の「編集」を依頼する場合に最適です。編集者の仕事は、単に間違いを特定するだけではなく、それらを修正し、より良い構造、スタイル、トーンに書き直すことでもあります。 AI の「編集者」が変更を加えると、著者の元の言葉やスタイルが、今日のあらゆるブランドのコンテンツで見られる一般的な企業向け LinkedIn のトーンに置き換えられます。 (そして、ブランド ボイスが AI によって生成されたスロップに置き換えられたと顧客が考える瞬間、ROI を言うよりも速く背景に消えていくことになります。)
したがって、クロードが提案した修正が作者の声も作者の意図も保持しないことは驚くべきことではありません。
1. 意図的な選択を欠陥と診断した
元の記事では、Skyscanner CAPTCHA と Google 一括検索という 2 つのクロード セッションを並行して実行し、それらのセッションを交互に挟んでナレーションを入れています。これは、著者がセッションを実行した順序がその順序であるためです。
Skyscanner 対 Google の並行する二重の物語は、1 つの CAPTCHA ケースに崩壊します。 [...] Skyscanner (より難しく、より興味深い失敗) を選択し、Google の打撃を断ち切ります。
その「修正」を適用すると、記事の最良の物語装置、そして偶然にも、最も実用的なアドバイスを含むセクションが削除されます。
2. 2 つのモデルはまったく逆のアドバイスをしました
同じ質問をしました。どの単一セクションを削除しても損失が最も少ないでしょうか? – 2 つのモデルは、相反する理由から、反対のセクションを選択しました。
あるモデルは、記事の中で最も価値のあるセクションを呼び出します。もう一方は削除すると言っています。誰かがそれを判断する必要があり、それが第 3 のモデルであることはできません。
これが、私が AI を栄光の魔法の 8 ボールと呼ぶ理由です。今日得られるものは、明日得られるものではないかもしれません。 AI を中心としたコンテンツ システムを構築しようとしている場合は、予測可能な条件に留意する必要があります。

真実性は当てにできるものではありません。
3. 存在しないエラーをでっち上げた
この記事では、Steel はエージェントの生の Web 検索結果には決して表示されないにもかかわらず、エージェントは「GEO を表示する方が SEO よりもはるかに優れている」として推奨したと述べています。
Opus はこれを論理エラーとしてフラグを立てました。
GEO は生成エンジンによって表面化されるものであり、言及しないことは GEO の失敗であり、勝利ではありません。
しかし、文は書かれているとおり正しいです: 検索結果に表示されない (悪い SEO)、とにかく LLM によって推奨される (良い GEO)。 Opus は主張を誤解し、存在しないエラーに対する修正を自信を持って規定しました。寓話も同じテキストで問題ありませんでした。
4. 構造上の誤りに対する間違った修正を規定した
「エージェントが実際にこれを使用できるかどうかを見てみましょう…」という移行段落を、クイックスタート セクションの上部から発見可能性セクションの中央に移動しました。
どちらのモデルも何かが間違っていることに気づきました。Fable は、クイックスタートが「トランジションなしで突然開く」ようになったと個別にフラグを立てました。しかし、どちらも 2 つの観察を結びつけませんでした。どちらも段落をカットすることを推奨しました。
どちらも、本来あるべき場所に戻すという正しい修正が何であるかを示唆しませんでした。
5. 書き換えにより、作者の声が完全に無効化されます
これまでのところ、モデルにフラグを立てるだけです。書き換えると何が起こるでしょうか?
これは、トリアージ プロンプト「読者はどこで読むのをやめてタブを閉じますか? その行を引用してください」に対するファブルの応答です。
このセクションの冒頭の一文は懐疑論者を説得するはずだが、逆に彼らを侮辱している。 「（間違っているが、何でも）」は、著者が気にしていないと説得することが最も必要な読者に正確に伝えます。
– そして、私が書き直しを求めたときに提案された修正:
エージェントと一緒に別の開発者として行動し、製品が各段階でどれだけうまくいくかを確認してください。
オーパスはmじゃなかった

括弧も削除するとさらに良いです。
現在、他の多くの企業がソフトウェアと対話する主な方法としてエージェント AI を使用しています。
嫌味な性格は消え、この作品を他の作品よりも読む価値のあるものにしている個性も消えました。著者は読者に、反AIバンドワゴンは間違っていると考えており、それに目を丸くしていることを知ってもらいたいと考えている。
他のLLM以外の誰も読まないような退屈な文章を書こうとするのであれば、この線を切るのは全く問題ありませんが、人間の読者を惹きつけて、息づかいのある聴衆とつながろうとするのであれば、全く役に立ちません。
どちらのモデルでも 2 つのエラーが見逃されました。代名詞の一致エラーと文の断片であり、どちらも平易な散文文法です。
どちらのモデルも、フレームワークの入れ替え、見出しの反転、自己反駁の主張、段落の移動など、シードされたロジックと構造のエラーをすべて検出しましたが、それらを修正するための提案が不完全または誤っていました。どちらのモデルも少なくとも 2 つのラインレベルのエラーを見逃していました。
私たちは人間をリソースとして評価する必要があるという不快な立場にいますが（企業にとっては何も目新しいことではありません）、この点では実際に人間は AI に勝っています。人間には有限のエネルギーと時間があり、生活賃金を支払う必要があるかもしれませんが、大手テクノロジー企業の気まぐれで上限が設定されたりキャンセルされたりする 1 日の使用量制限に制限されることもありません (参照: 最初に一般公開されてから数日後に削除された Fable、ChatGPT 5 アップデートをめぐる初期の騒動、LLM 企業が近い将来に無料枠を削除またはハンディキャップにするという一般的な警戒)。
いずれにせよ、ほとんどの人は 7 月 12 日以降、Fable (より正確には「エディター」) にアクセスできなくなるため、Opus 4.8 以下のモデルの制限内で作業することになります。
このため、人間の編集者を継続的に配置する方がはるかに賢明です。

彼は、疑いを持たない記事に対して一連の自動化された AI ワークフローを解き放つのではなく、指揮を執ります。 (少なくとも、公開するものすべてを AI の粗末なものにしたくないのであれば、それは賢明です。)
この実験の結果に基づいて、人間とクロードの推奨される役割分担は次のとおりです。
予算: どの記事でも 60 個のプロンプトを実行することはできません。
プロンプト ライブラリには約 60 のプロンプトがあり、私たちの手法ではその多くを新しいセッションで実行し、各セッションで記事全体を再読み込みしました。 Claude のサブスクリプションでは、5 時間のローリング ウィンドウと毎週の上限があり、Opus モデルは Sonnet の数倍の速さでクォータを消費しますが、これは高価なワークフローです。
AI の使用コストを最適化するための 3 つのルールをお勧めします。
1. ロジックと構造には小さなモデルを使用します。
Fable と Opus はロジックと構造のエラーに関して同じスコア (4/4) を示したため、モデルの品質は何も役に立ちませんでした。相違点はラインレベル (17/19 対 12/19) であり、それはより優れたモデルであっても単独では信頼できない層であるため、人間による校正によってバックストップが行われます。とにかく人間が再チェックしなければならないタイプミスを見つけるためにフロンティアモデルの料金を支払うのは、割り当ての最悪の使い方です。
2. 定期的な編集では、トリミングされたプロンプトのセットがバッチ化されて取得されます。通常の記事の場合: トリアージ プロンプト、1 つのロジック パス、1 つの構造パス、1 つの整合性スイープ (3 番目の

[切り捨てられた]

## Original Extract

LLMs: A technical editor's tool or replacement?

Skip to main content Tech Stackups Home Topics Comparisons Guides Articles News AX LLMs for Technical Editing: The Good, the Bad, and the Ugly
With the existence of Opus 4.8 and the limited re-release of Fable to the global public, you may be thinking that it’s possible to completely replace your writers and editors with AI.
It’s certainly possible, but it would also be the most inefficient, self-sabotaging decision you could make if you want people to actually care about your content and connect with your brand.
That said, I’ll admit that I have a bias – I'm an editor who finds the corporate obsession with AI counterproductive .
Nevertheless, some people are still convinced that AI can replace editors, and I'm going to show you why that isn't true. In the interests of remaining objective, I've used AI to edit an already published article seeded with errors.
By the end of this experiment, we'll be able to tell where Claude falls between two extremes: Can it replace editors entirely, or is it just fancy (and sometimes incorrect) autocomplete?
I seeded our published AX article
with 23 errors of varying severity:
Then I asked both Opus 4.8 and Fable to evaluate the error-ridden article, using prompts from
our editing prompt library .
NB: I asked Claude to identify the errors first before suggesting fixes.
(I'd also recommend reading this article to see how Opus 4.8 fared on editing the clean version of the AX article.)
The Good: Holding your piece together ​
Okay, as much as I hate to say it, Claude did really well with flagging structural and logical errors. If there was a mismatch in your headings, contradictions in content, or even missing information, both Opus and Fable were good about flagging these errors and suggesting appropriate fixes.
1. It caught a framework contradiction ​
The article defines three hurdles – discovery (does the agent know about you?), onboarding (can it sign up?), and usage. I seeded this sentence, which swaps the first two labels:
Does the agent sign up with minimal help from its handler (discovery?) Does it know about you (onboarding)?
This directly contradicts the framework the article itself set up two sections earlier, where discovery is "the agent should know that you provide a solution" and onboarding is the sign-up flow. The one conceptual takeaway a reader is meant to walk away with is stated backwards.
Catching this requires holding definitions from two sections earlier in mind and checking a later sentence against them. Simple pattern-matching tools wouldn't be able to do that.
2. It caught a heading arguing against its own section ​
I inverted a heading to read "Signing up is usually still more about AX than DX" – directly above body text explaining that it's still a human who visits your sign-up page.
If a human does the signing up, that's the developer's experience being tested,
not the agent's — the heading should say sign-up is still more about DX than AX.
3. It caught a claim refuted by its own example ​
I swapped the subjects in this sentence, so that the example now proves the opposite of the claim:
Note that humans generally do much longer web searches than agents. While a human would have searched for something like Steel captcha , Claude does Steel.dev solve captcha session config Python SDK 2025 .
The human query is shorter; the agent's is longer. The example proves the
opposite of the claim.
4. It did 30 minutes of consistency checking in one pass ​
Both models swept the piece for mechanical inconsistencies in a single prompt: a British "optimise" in a US-English article, one curly apostrophe among dozens of straight ones, spaced en dashes in a document that uses em dashes, "head-less" versus "headless", and lowercase "captcha" against uppercase "CAPTCHA".
It may surprise you to learn that where Claude struggled the most was with picking out typos and simple grammatical errors.
1. Both models looked straight at an error and flagged the wrong thing ​
I changed "it's" to "they're" in this sentence, breaking the pronoun agreement:
Skyscanner has a simpler CAPTCHA but they're not part of Steel's automatic solving, so it took longer.
Both models flagged a typographical error, but neither spotted the pronoun mismatch.
Line 248: "they're" uses a curly apostrophe; every other contraction in the file uses straight apostrophes.
The document is straight quotes/apostrophes throughout (verified) except one curly apostrophe — "they're" at 248.
2. The devil's (not) in the details ​
Fable caught most of these, but Opus missed several errors at the line level – the "easy" stuff:
If you ran only Opus – the only one of the two that's still going to be available after July 12 – you'd end up publishing seven errors. And they both missed the last two errors, so you wouldn't get a clean run with either one.
The Ugly: Confident advice that would make the article worse ​
There’s a reason we call AI-generated writing “slop.” LLMs can only draw from a limited dataset, which means that no matter how large that dataset is, the model will eventually run out of “unique” ways to say things. That’s why it’s so easy for us to now look at a piece of writing and gauge whether or not AI was involved in writing it – there are glaringly obvious signs in the concepts, structure, and especially the phrasing.
The same is true for when you ask AI to “edit” a piece of writing. An editor’s job isn’t merely to identify errors, but also to fix them and rewrite for better structure, style, and tone. Any changes an AI “editor” makes will replace the author’s original words and style with the generic corporate LinkedIn tone you see on every brand’s content nowadays. (And the second your customers think you’ve replaced your brand voice with AI-generated slop, you’re going to fade into the background faster than you can say ROI.)
It should then come as no surprise that Claude's proposed fixes would preserve neither authorial voice nor authorial intent.
1. It diagnosed a deliberate choice as a defect ​
The original article runs two Claude sessions in parallel – a Skyscanner CAPTCHA and a Google mass-search – and narrates them interleaved, because that's the sequence in which the author ran them.
The parallel Skyscanner-vs-Google dual narrative — collapse to one CAPTCHA case. [...] Pick Skyscanner (the harder, more interesting failure) and cut Google's blow-by-blow.
Applying that "fix" deletes the article's best narrative device – and, as it happens, the section containing its most actionable advice.
2. The two models gave flatly opposite advice ​
Asked the identical question – which single section could be deleted with the least loss? – the two models chose opposite sections, for contradictory reasons:
One model calls a section the most valuable in the article; the other says to delete it. Someone has to adjudicate that, and it can't be a third model.
This is why I call AI a glorified magic 8-ball: what you get today may not be what you get tomorrow. If you're trying to build a content system with AI as the fulcrum, you'll need to keep in mind that predictable quality is not something you can count on.
3. It invented an error that isn't there ​
The article observes that Steel never appears in the agent's raw web-search results, yet the agent recommended it anyway – "showing its GEO is much better than its SEO."
Opus flagged this as a logic error:
GEO is about being surfaced by generative engines — a no-mention is a GEO miss, not a win.
But the sentence is right as written: absent from search results (bad SEO), recommended by the LLM anyway (good GEO). Opus misread the claim, then confidently prescribed a fix for an error that doesn't exist. Fable, on the same text, saw no problem.
4. It prescribed the wrong fix for a structural error ​
I moved the transition paragraph "Now let's see if the agent can actually use this thing…" from the top of the Quickstart section into the middle of the discoverability section.
Both models noticed something was wrong – Fable even separately flagged that Quickstart now "opens abruptly with no transition." But neither connected the two observations. Both recommended cutting the paragraph.
Neither suggested what would have been the correct fix – putting it back where it belongs.
5. Its rewrites completely neuter the author's voice ​
So far I've only let the models flag. What happens when they rewrite?
Here's Fable's response to the triage prompt "Where would a reader stop reading and close the tab? Quote the line."
The opening line of the section is supposed to persuade skeptics — and it insults them instead. "(mistakenly, but whatever)" tells the exact reader who most needs convincing that the author can't be bothered
– and the fix it suggested when I asked for a rewrite:
Play-act as another developer with an agent and see how well your product fares at each stage.
Opus wasn't much better, also removing the parentheses:
Many others now use agentic AI as their primary way of interacting with software.
The snark is gone — and so is the personality that makes this piece worth reading over others. The author wants the reader to know that they think anti-AI-bandwagoners are mistaken, and that they're rolling their eyes about it.
Cutting this line is all well and good if you're going for boring writing that nobody besides other LLMs will ever read, but no use to you at all if you're trying to hook human readers and connect with a breathing audience.
Two errors were missed by both models – the pronoun agreement error and the sentence fragment, both plain-prose grammar.
Both models caught every seeded logic and structure error: the swapped framework, the inverted heading, the self-refuting claim, the displaced paragraph, but made incomplete or erroneous suggestions for fixing them. Both models missed at least two line-level errors.
We're in the uncomfortable position of needing to evaluate humans as a resource (nothing new for corporate), but they actually win against AI in this regard. Humans may have finite energy and time and need to be paid living wages, but they’re also not confined to a daily usage limit that’s capped or canceled at the whims of Big Tech (see: Fable being pulled days after it was first released to the public, the initial furore over the ChatGPT 5 update , and general wariness that LLM companies are going to remove or handicap their free tiers in the near future).
And in any case, most people are also going to lose access to Fable (the more accurate "editor") after July 12, so you'll be working within the limits of Opus 4.8 and smaller models.
For this reason, it's far more sensible to continue to have a human editor at the helm, rather than unleashing a series of automated AI workflows on your unsuspecting articles. (Or at least, it's sensible if you don't want everything you publish to be AI slop.)
Here's the recommended human-to-Claude duty separation, based on the results of this experiment:
The budget: you can't run 60 prompts on every article anyway ​
The prompt library has around 60 prompts, and our methodology ran many of them in fresh sessions, each one re-reading the whole article. On a Claude subscription — five-hour rolling windows plus weekly caps, with Opus models burning through quota several times faster than Sonnet — that's an expensive workflow.
We'd suggest three rules for optimising your AI use cost:
1. Use smaller models for logic and structure.
Fable and Opus scored identically — 4/4 — on logic and structure errors, so model quality bought nothing there. Where they diverged was line-level (17/19 versus 12/19), and that's the tier where even the better model can't be trusted alone, so a human proofread backstops it regardless. Paying frontier-model rates to catch typos a human must re-check anyway is the worst possible use of your quota.
2. Routine edits get a trimmed set of prompts, batched. For an ordinary article: use the triage prompts, one logic pass, one structure pass, one consistency sweep (the thre

[truncated]
