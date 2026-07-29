---
source: "https://blog.cryptographyengineering.com/2026/07/29/some-notes-about-anthropics-new-results/"
hn_url: "https://news.ycombinator.com/item?id=49099804"
title: "Some thoughts about Anthropic's new cryptanalysis results"
article_title: "Some thoughts about Anthropic’s new cryptanalysis results – A Few Thoughts on Cryptographic Engineering"
author: "supermatou"
captured_at: "2026-07-29T17:05:01Z"
capture_tool: "hn-digest"
hn_id: 49099804
score: 2
comments: 0
posted_at: "2026-07-29T16:42:20Z"
tags:
  - hacker-news
  - translated
---

# Some thoughts about Anthropic's new cryptanalysis results

- HN: [49099804](https://news.ycombinator.com/item?id=49099804)
- Source: [blog.cryptographyengineering.com](https://blog.cryptographyengineering.com/2026/07/29/some-notes-about-anthropics-new-results/)
- Score: 2
- Comments: 0
- Posted: 2026-07-29T16:42:20Z

## Translation

タイトル: Anthropic の新しい暗号解析結果についてのいくつかの考え
記事のタイトル: Anthropic の新しい暗号解析結果についてのいくつかの考え – 暗号エンジニアリングに関するいくつかの考え
説明: 昨日、Anthropic は 2 つの新しい暗号解析結果を公開しました。どちらも、(まだ) 未リリースの高度なモデルである Claude Mythos の出力です。これらの結果の 1 つ目は HAWK と呼ばれる署名スキームを攻撃し、2 つ目は削減ラウンド AES に対する改良された攻撃です。 Anthropic もブログ pos を公開しました
[切り捨てられた]

記事本文:
Anthropic の新しい暗号解析結果についてのいくつかの考え – 暗号工学に関するいくつかの考え
コンテンツにスキップ
ホーム
メニュー
Anthropic の新しい暗号解析結果に関するいくつかの考え
私の学術ウェブサイト
ブルースカイ
マストドン
ツイッター
人気の投稿
便利な暗号リソース
ビットコインチップジャー
クリプトパルの課題
応用暗号研究: 委員会
暗号工学ジャーナル
(このブログとは関係ありません)
検索:
人気の投稿とページ
Anthropic の新しい暗号解析結果に関するいくつかの考え
暗号化された推論について話しましょう
Apple 様: 今すぐ「消えるメッセージ」を iMessage に追加してください
Siri の将来、または: プライベート推論が十分プライベートではない理由
Telegram は本当に暗号化されたメッセージング アプリですか?
ゼロ知識証明: 図解による入門書
匿名認証情報: 図解による入門書
認証された暗号化モードを選択する方法
昨日、Anthropic は 2 つの新しい暗号解析結果を公開しました。どちらも、(まだ) 未リリースの高度なモデルである Claude Mythos の出力です。これらの結果の 1 つ目は HAWK と呼ばれる署名スキームを攻撃し、2 つ目は削減ラウンド AES に対する改良された攻撃です。 Anthropic は、これらの結果を生み出した研究プロセスを説明するブログ投稿もリリースしました。オンラインで何人かの人が、これは何を意味するのかと私に尋ねました。すべての答えがあるかどうかはわかりませんが、現在の理解について少し書いても問題ないと思いました。これらは私の考えにすぎず、他の人（問題の 2 つの分野の専門家を含む）はおそらく異なるでしょう。そのため、それらをありのままに受け止めてください。
2 つの新しい結果は 2 つの非常に異なる領域をカバーしており、全体的に品質が非常に異なります。世界についての広範な意見や、すべての暗号通貨を売却すべきかどうかについて話す前に、少し時間を取って話しましょう。

物質。
鷹。 1 つ目は、非標準の署名スキーム HAWK に対する新しい鍵回復アルゴリズムです。 HAWK は、モジュール格子同型問題 (モジュール LIP) に基づいて提案されたポスト量子安全署名スキームです。クロードが書いた結果自体の簡単な要約については、ここを参照してください。この結果について知っておくべきことが 5 つあります。
HAWK は導入されたアルゴリズムや標準に採用されたアルゴリズムではなく、提案されたアルゴリズムです。これは標準化されている Falcon 署名スキームに関連していますが、攻撃はその設定 (別の難しい問題に基づいています) には移行しません。
ただし、HAWK は、将来の標準として評価される段階ではやや進んでいませんでした。
この攻撃は、SF の意味で「実際に配備された」HAWK を破壊するものではありません。結果として生じる攻撃は依然として指数関数的な時間ですが、アルゴリズム内のセキュリティの「ビット」数はおよそ半分になります。つまり、理論的にはキーのサイズを 2 倍にすることで修正できる可能性があります。欠点は、これによりスキームの効率が低下することであり、HAWK は代替手段よりも効率的であるということで完全に正当化されているため、スキームの存在を正当化することがはるかに困難になります。
この攻撃により、作成者がこの目的のために提供した、HAWK の弱体化された「チャレンジ インスタンス」に対して、実測時間で数時間で実行される実際のコードが生成されました。このインスタンスは、実際のデプロイメント用に提案されたパラメータを使用していませんが、暗号解読の弱点を十分に示しています。
特に懸念されるのは（そして AI にとって特に機が熟しているのは）、この攻撃が根本的に新しい数学を発明するものではないということです。これは、あちこちに転がっていてよく知られている一連のツールを拡張するだけで、良い結果が得られます。
この最後の部分は重要です。クロードに意見を聞いてみたが、そうではなかった

「これが本当に興味深いのは、そして率直に言って、この分野にとっては少し恥ずかしいことですが、材料がどれも珍しいものではないということです。」 TL;DR は、誰かが私たちの既知のツールをすべて適用して、より徹底的な仕事をしたということです。一言で言えば、AI を攻撃する類のことは素晴らしいということです。
AES。 2 番目の結果は、ラウンド数を削減した AES に対する新たな攻撃です。ほとんどの人は「AES への攻撃」と聞いてパニックになるため、この結果は最初はもっと刺激的に聞こえます。しかし、これはあまり面白くない結果でもあります。
このブログを読んでいるほとんどの人は、AES があらゆる場所で使用されている標準のブロック暗号であることを知っているでしょう。これは 2001 年以来の標準であり、導入されたバージョンはこれまでのところ、NSA によって実行されたかなりの量の非公開テストを含む、これまでに課せられたすべての重要な問題に耐えてきました。完全な暗号を攻撃するのは非常に難しいため、暗号解読者は暗号の弱体化されたバージョン、または「ラウンド数を減らした」バージョンに対して作業を行うのが標準です。完全な AES 暗号は、キーのサイズに応じて 10、12、または 14 ラウンドで実行されます。新しい Anthropic の結果は、暗号の弱い 7 ラウンド変種を攻撃します。
重要なことに、7 ラウンド AES に対する攻撃は新しいものではなく、いくつかの攻撃が存在しています。実際、この新しい Anthropic の結果は、2013 年以降の以前の研究に比べて、一定の要素によるわずかな改善です。これらの攻撃が実際に AES を「破る」ことからどの程度離れているかを理解していただくために、見出しの結果に注目したいと思います。新しい攻撃には 2 89 の暗号操作が必要であり、さらに悪いことに、この作業は、実際の暗号化プログラムが秘密鍵の下で選択された平文の 2 105 の暗号化を生成するよう何らかの方法で納得させた場合にのみ可能です。これらはどちらも現実の世界ではほとんど実用的ではありません。そして、新しい結果が得られる一方で、

以前の結果に比べてこの攻撃の速度がわずかに向上しましたが、この結果の速度向上がどの程度「現実的」であるかさえ明らかではありません。実際の攻撃には 2 89 の操作が必要で、実際には「実行」できないため、私たちが行っているのは机上の分析であり、すべての詳細が実際に解決された場合に実際の実行時間の改善が得られるかどうかはわかりません。
これによって結果が悪くなるわけではありません。実際、テクニックの観点から見ても興味深いものです。しかし、これは私たちの知識のごくわずかな増加であり、HAWK の取り組みのような実際的な新しい攻撃ではありません。つまり、TL;DR: ここには大幅に新しい数学的結果はありません。しかしそれでも、実際の暗号解読の進歩は科学者を興奮させるものです。そして確かに、HAWK の結果は非常に意味のあるものです。なぜなら、そのスキームには標準化の本当のチャンスがあったのに、現在は (おそらく) そうならないはずだからです。
では、私たちがどのようにしてここにたどり着いたのか、そしてそれが何を意味するのかについて話しましょう。
Anthropic はどのようにしてこのような結果を得たのでしょうか?
Anthropic の投稿には彼らが何をしたか詳しく書かれており、正直言って、ちょっと面白いです。いいえ、Anthropic のチームは、新しい結果を見つけるために AI を注意深く調整する大規模なドメイン専門家ではありませんでした。彼らはただ何らかの結果を得るように指示し、結果が見つかるまで砥石に鼻を縛り付けたようです。私を疑う人のために、彼らが使用したプロンプトの例をいくつか示します (投稿から引用)。
はい、AI はかなり良くなってきています。つまり、既存の暗号解析結果を理解し、それらを合成して実際の新しい攻撃を作成し、さらにはそれを拡張することもできるようになりました。彼らは明らかに人間による詳細な介入なしでこれを実行できるようです。これはまだ超インテリジェントな暗号解析ではありません。しかし、それはかなり印象的です。
検証可能性がボトルネックになっている
研究者として、私はモデルたちと多くの時間を費やし、対話してきました。

さまざまなアイデア。過去数か月の間にさえ、明らかに良くなっていると言っても誰も驚かないと思います。私には Mythos もありませんし、使える 10 万ドルもありませんが、少なくとも 1 つの新しい高度な未リリース モデルをクエリすることができました。また、数年間興味を持っていた質問に対していくつかの驚くべき新しい「結果」も得られました。
ここで本当の問題が生じます。モデルが明らかに新しい結果を吐き出したからといって、その結果が実際であるとは限りません。モデルが実際の結果を生成するのが得意であっても、実際のように見えるが誤解を招く結果を生成することの方がはるかに優れています。これは非常にイライラする可能性があり、多くの場合、人間の注意がこれまで以上に必要であることを意味します。
このルールには例外があります。HAWK のような「完全な」攻撃の場合、攻撃は (スキームの弱いバージョンに対して) 数時間で実行され、検証は非常に簡単です。コードを送信するだけで、キーが回復され、実際に選択されたメッセージに署名されることを誰でも確認できるようになります。 AES の結果のような、より巧妙な高速化攻撃の場合、有効性をチェックするのはそれほど簡単ではありません。ここでのアプローチはより具体的です。ここでは、形式的に検証可能なリーン証明が役立ちますが、(これらの証明が簡単に作成できる場合でも)、そのような証明は定理ステートメントをどのように定式化したかに依然として非常に敏感であり、人間の専門家によるチェックが必要になることがよくあります。
おそらく、最近の興味深い数学的結果の多くにこの傾向があることに気づくでしょう。それらには、よく理解されている定理の機械チェック可能な証明が含まれているか、(ヤコビアン予想のように) 計算できる単純な反例を見つけることが含まれています。あるいは、大勢の専門家が多くの時間を費やして結果を検討し、最終的にはその結果に納得しました。このソムの必要性

人間が作業をチェックすることは、私たちの進歩を遅らせることになります。壊滅的ではない暗号解読の例については、おそらくこれがしばらくの間続くことになるでしょう。
現実世界への影響は何でしょうか?
この質問に対する答えは、暗号の利用者について話しているのか、科学者について話しているのか、それとも人類全体について話しているのかによって大きく異なります。これらを一度に見てみましょう。
暗号化のユーザーにとって、良いニュースが 2 つと、混合ニュースがいくつかあります。 1 つ目は、対称暗号が非常に厄介で堅牢であるということです。トラクターを流砂地帯に引きずり出し、セメントの下に埋める農夫を想像してみてください。これが対称暗号設計のようなものです。素早く簡単に適用できる構造を考え出すように意図的に設計されていますが、非常に厄介で、もつれを解くのが困難です。多くの新しい生のインテリジェンス時間を追加しても、これが魔法のように改善されることはおそらくありません。 AI は最終的にはまったく新しい技術を使用してこれらの問題に対して真の進歩を遂げることができるかもしれませんが、今のところ、そのために必要となる真に画期的な直観力を実証できていません。たとえそうするとしても、暗号自体を改良して解読をより困難にすることができると期待する十分な理由があります。
公開鍵暗号化はさらに厄介です。公開鍵暗号には、一方向の高速計算は許可するが、もう一方の方向への計算は許可しないが、一方の当事者がプロセスを逆転できる便利な「落とし戸」を備えた数学的オブジェクトが必要です。私たち人間は、これを可能にする非常に都合よく構造化された数学的オブジェクトをほんの一握りしか思いつきませんでした。それらには、(EC)DLP 問題、RSA、格子問題、およびコーディング理論の領域からのさまざまな問題など、推測される「難しい問題」が含まれています。私たちが全力を尽くしてきた一方で、

これらの問題の分析に専念する人間が十分ではないため (RSA のような古い問題でさえも!)、これ以上優れた攻撃は存在しないと確信できるほどです。これは、コードベースの暗号や格子ベースの暗号などの新しい分野に特に当てはまります。
つまり、AI が真の進歩を遂げるための肥沃な土壌がたくさんあるということです。
そうは言っても、私は良いニュースがあると言いました、そしてそれは本気でした。現在、私たちは EC ベースの暗号化と RSA に基づく従来の公開鍵アルゴリズムから、新たな問題に基づく新しいポスト量子アルゴリズムに移行する歴史的な移行の真っ只中にいます。これが、HAWK のような非常に多くの標準が検討されている理由です。大規模な新しい公開暗号解析機能が稼働するのに最適な時期があったとしたら、私たちはまさにその時期にいます。したがって、AI が私たちのすべての困難な問題を完全に弱体化させることに成功しない限り (または、私たちが Impagliazzo の Minicrypt の中に住んでいる場合)、AI が暗号解読を得意とするのにこれ以上良い時期はありません。最良の場合、その結果、特定した問題に対して真の自信が得られ、暗号解析の文献がより堅牢になります。うまくいけば。
科学者にとって、これは素晴らしい時期でもあります。あなたには今、一緒にいて楽しいプラスチック製の友達がいて、最も難しい問題について話し合うことができます。同時に、それはまだ十分にスマートではありません

[切り捨てられた]

## Original Extract

Yesterday Anthropic published two new cryptanalysis results, both outputs of Claude Mythos, their (still) unreleased advanced model. The first of these results attacks a signature scheme called HAWK, while the second is an improved attack against reduced-round AES. Anthropic also released a blog pos
[truncated]

Some thoughts about Anthropic’s new cryptanalysis results – A Few Thoughts on Cryptographic Engineering
Skip to content
Home
Menu
Some thoughts about Anthropic’s new cryptanalysis results
My academic website
BlueSky
Mastodon
Twitter
Top Posts
Useful crypto resources
Bitcoin tipjar
Cryptopals challenges
Applied Cryptography Research: A Board
Journal of Cryptographic Engineering
(not related to this blog)
Search for:
Top Posts & Pages
Some thoughts about Anthropic's new cryptanalysis results
Let's talk about encrypted reasoning
Dear Apple: add "Disappearing Messages" to iMessage right now
The future of Siri, or: why private inference isn’t private enough
Is Telegram really an encrypted messaging app?
Zero Knowledge Proofs: An illustrated primer
Anonymous credentials: an illustrated primer
How to choose an Authenticated Encryption mode
Yesterday Anthropic published two new cryptanalysis results , both outputs of Claude Mythos, their (still) unreleased advanced model. The first of these results attacks a signature scheme called HAWK, while the second is an improved attack against reduced-round AES. Anthropic also released a blog post describing the research process that produced these results. A few people online have asked me what this all means. While I’m not sure I have all the answers, I figured it wouldn’t hurt to write a bit about my current understanding. These are only my thoughts and other folks will probably differ (including domain experts in the two areas at issue) so take them for what they are.
The two new results cover two very different areas, and are overall just very different in quality. Before we get to broad statements about the world, and whether you should sell all your cryptocurrency, let’s take a minute to talk about the substance.
Hawk. The first is a new key recovery algorithm against the non-standard signature scheme HAWK . HAWK is a proposed post-quantum-safe signature scheme that’s based on the module Lattice Isomorphism Problem (module-LIP). For a brief Claude-written summary of the result itself, see here . There are five things you need to know about this result:
HAWK is not a deployed or standards-adopted algorithm, it’s a proposed algorithm . It is related to the Falcon signature scheme , which is being standardized, but the attack does not transfer to that setting (which is based on a different hard problem.)
However, HAWK was somewhat far along in the process of being evaluated for a future standard.
The attack does not break “real deployed” HAWK in the sci-fi sense. The resulting attack is still exponential time, but roughly halves the number of “bits” of security in the algorithm. That means it could theoretically be fixed by doubling key sizes. The downside is that this makes the scheme less efficient, and, since HAWK is entirely justified by being more efficient than alternatives, that makes the existence of the scheme much harder to justify.
The attack produced real code that runs in a few hours of wall-clock time against a weakened “challenge instance” of HAWK that the authors provided for this purpose . While this instance doesn’t use the parameters that were proposed for real deployment, it does demonstrate the cryptanalytic weakness well enough.
What’s particularly concerning (and so especially ripe for AI) is that the attack does not invent fundamentally new mathematics. It simply extends a bunch of tools that were lying around and well-known, and gets a good result.
This last part is important. I asked Claude for its thoughts, and it doesn’t mince words: “what makes this genuinely interesting — and, frankly, a little embarrassing for the field — is that none of the ingredients are exotic.” The TL;DR is that someone just did a much more thorough job applying all of our known tools. In short: the sort of things that attack AIs are wonderful at.
AES. The second result is a new attack on reduced-round AES. This result initially sounds more exciting, since most people hear “attack on AES” and panic. However, this is also the result that’s much, much less interesting.
Most folks reading this blog will know that AES is a standard block cipher that’s used just about everywhere. It’s been a standard since 2001, and the deployed version has so far withstood everything significant that’s been thrown at it: that includes a substantial amount of non-public testing performed by the NSA . Since attacking full ciphers is very difficult, it’s standard for cryptanalysts to do their work against weakened, or “reduced-round” versions of a cipher. The full AES cipher runs for either 10, 12 or 14 rounds depending on key size. The new Anthropic result attacks a weaker 7-round variant of the cipher.
Critically, attacks against 7-round AES are not new: there have been several of these. In fact, this new Anthropic result is a modest constant-factor improvement on previous work from back in 2013 . To give you a sense of how far these attacks are from really “breaking” AES, I’d note the headline results: the new attack requires 2 89 cipher operations and, even worse, this work is only possible after you’ve somehow convinced a real encryptor to produce 2 105 encryptions of chosen plaintexts under their secret key! Neither of these things is remotely practical in the real world. And while the new result modestly speeds up this attack over the previous result, it’s not even clear how “real” the speedup in this result is: since the actual attack requires 2 89 operations and can’t really be “run”, what we have is an on-paper analysis that may or may not yield an actual runtime improvement if all details are actually worked out.
This does not make the result bad! In fact it’s still interesting from a techniques point of view. But it is very much a small increment in our knowledge, not a practical new attack like the HAWK work. So TL;DR: no wildly new mathematical results here. But still, real cryptanalytic progress of the sort that make scientists excited. And certainly the HAWK result is very meaningful, since that scheme had a real chance at standardization and is now (very likely) not going to be.
Now let’s talk about how we got here, and what it all means.
How did Anthropic get these results?
The Anthropic post is detailed about what they did, and honestly, it’s kind of hilarious. No, the team at Anthropic was not a large set of domain experts that carefully tuned their AI to find novel results. They appear to have just told it to get some results and then strapped its nose to the grindstone until it found some. If you doubt me, here are some examples of the prompts they used (cited from their post):
So yes, the AIs are getting pretty good. In short: they are now capable of understanding existing cryptanalysis results, synthesizing them into real new attacks, and even extending them. They can apparently do this without detailed human intervention. This isn’t yet super-intelligent cryptanalysis. but it’s pretty damn impressive.
Verifiability is now the bottleneck
As a researcher I’ve also been spending a lot of time with models, talking through various ideas. I don’t think I will surprise anyone when I say that they’re obviously getting better, even over the course of the past few months. While I don’t have Mythos and $100k to spend, I have been able to query at least one new advanced unreleased model, and I also have received some surprising new “results” to questions that I’ve been interested in for a few years.
Which brings me to the real problem: just because a model spits out an apparent new result, this does not mean the result is real . Even if models are good at producing real results, they’re much better at producing results that look real but are misleading. This can be enormously frustrating, and often means that human attention is more necessary than ever.
There are exceptions to this rule: for “full” attacks like HAWK, where the attack runs in a few hours (against a weaker version of the scheme), verification is extremely easy. You can just send over the code and let anyone check that it recovers keys and signs real chosen messages. For more subtle speedup attacks like the AES result, checking validity is not so easy. Here the approach is more specific: formally-verifiable Lean proofs can help here, but (even where these proofs are easy to make), such proofs are still highly sensitive to how you’ve formulated the theorem statement , and that often requires human experts to check.
You’ll probably notice that many of the exciting recent mathematical results have had this flavor: they either include a machine-checkable proof of a well-understood theorem, or (like the Jacobian conjecture) they involve finding a simple counterexample you can compute on. Alternatively, a bunch of experts spent a lot of time reviewing the result and were eventually convinced by it. This need for some humans to check the work is going to slow down our progress. For non-devastating examples of cryptanalysis, this is probably where we’re going to be for a while.
What are the implications for the real world?
The answer to this question really depends on whether you’re talking about the consumers of cryptography, scientists, or humanity at large. Let’s take these one at a time.
For users of cryptography: there are two pieces of good news and some mixed news. The first is that our symmetric ciphers are very messy and robust. Imagine a farmer who drags a tractor out into a patch of quicksand, and then buries it under cement. That’s what symmetric cipher design is like; it’s deliberately designed to come up with structures that are quick and easy to apply, but very messy and hard to untangle. The addition of many new raw intelligence-hours probably aren’t going to magically improve this. AIs may be able to eventually make real progress against these problems using entirely new techniques, but so far they’re not demonstrating the truly groundbreaking intuition that would be required to do so. And even if they do : there’s a good reason to hope that they’ll be able to improve the ciphers themselves to make them much harder to break.
Public-key cryptography is messier. Public-key crypto requires a mathematical object that admits fast calculations in one direction, but not the other, and yet has a convenient “trapdoor” that lets one party reverse the process. We humans have come up with only a handful of very conveniently-structured mathematical objects to enable this: they involve conjectured “hard problems” like the (EC)DLP problem, RSA, lattice problems, and various problems from the domain of coding theory. While we’ve given these our all, there simply have not been enough human beings dedicated to analyzing these problems (even the older ones like RSA!), such that we can be absolutely certain there are no more good attacks out there. This is particularly true for the novel areas like code-based crypto and even lattice-based cryptography.
That means there’s a lot of fertile ground for AIs to make real progress.
With that said, I said there was good news , and I meant it. Right now we’re in the midst of a historic transition from traditional public-key algorithms based on EC-based cryptography and RSA, moving over to new post-quantum algorithms based on novel problems. This is why there are so many standards like HAWK being considered. If there was ever a perfect time for a massive new public cryptanalysis capability to come on line, we’re in it. So unless AIs succeed in undermining all of our hard problems altogether (or we live in Impagliazzo’s Minicrypt ) then this could not be a better time for AI to get good at cryptanalysis. In the best case, the result is that we gain real confidence in the problems we’ve identified, and the cryptanalysis literature gets a lot more robust. Hopefully.
For scientists : this is also a wonderful time. You now have a plastic pal who’s fun to be with, and you can talk over your hardest problems. At the same time it’s not yet smart enough that

[truncated]
