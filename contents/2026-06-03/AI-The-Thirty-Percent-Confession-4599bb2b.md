---
source: "https://www.cringely.com/2026/06/02/the-thirty-percent-confession/"
hn_url: "https://news.ycombinator.com/item?id=48374488"
title: "AI: The Thirty Percent Confession"
article_title: "The Thirty Percent Confession | I, Cringely"
author: "dxs"
captured_at: "2026-06-03T00:37:55Z"
capture_tool: "hn-digest"
hn_id: 48374488
score: 3
comments: 0
posted_at: "2026-06-02T18:48:21Z"
tags:
  - hacker-news
  - translated
---

# AI: The Thirty Percent Confession

- HN: [48374488](https://news.ycombinator.com/item?id=48374488)
- Source: [www.cringely.com](https://www.cringely.com/2026/06/02/the-thirty-percent-confession/)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T18:48:21Z

## Translation

タイトル: AI: 30パーセントの告白
記事のタイトル: 30 パーセントの告白 |私、シリンジリー

記事本文:
30パーセントの告白 |私、シリンジリー
ホーム
前回、AI 業界は支払う必要のない税金を支払っているとお話しました。私たちが大々的に「AI」と呼んでいるものの多くは、実際にはただ調べているだけであり、私たちはその調査を、これまでに製造された中で最も高価なシリコンについて行うことにしたのです。皆さんの何人かが、私がそれを誇張していると書いてきました。確かに、何千億ドルも燃やしている人々は私が知らないことを知っている、とあなたは言いました。
だから今週はあなたと議論しません。代わりに、エンタープライズ ソフトウェアの最大手企業の 1 社に議論してもらいましょう。なぜなら、その企業はすでに研究論文を発表しており、あまり深く読まないでほしいと思っているようだからです。
その会社はセールスフォースです。同じ Salesforce が、人間の労働力の横に置く、たゆまぬデジタル労働力である「エージェント」、「代理店企業」を販売しています。建物の一部がマーケティングを担当する一方で、別の部分 (物を販売するのではなく測定することを仕事とする人々である Salesforce AI Research) は、今日の最高の AI が、輝かしく地味なこと、つまり普通の会社の混乱に散らばっているときに適切な情報を見つけることをどの程度うまく実行できるかを調べるためのテストを構築しました。糸のたるみ。 GitHub。会議の記録。誰も正しくファイルしていない文書。すべての実際のビジネスが実際に実行されているもの。
彼らはこれを HERB (異機種混合エンタープライズ RAG ベンチマーク) と名付けましたが、安易に構築したわけではありません。これは人工的ですが、念入りに現実的な会社です。30 の製品に 530 人の従業員が在籍し、39,190 のドキュメント、メッセージ、トランスクリプト、プル リクエストを生成し、実際の姿をちりばめています。論文は arXiv にあります。データはHugging Faceにあります。誰でも私の算術をチェックできるので、私が喜んで構築する理由はまさにこれです

その上に列を追加します。
Salesforce が、お金で買える最高のエージェント検索システムを HERB で開発したとき、それは計画とツールの使用を備えた最上位モデルで、優れたものでした。彼らは 100 点中 32.96 点を獲得しました。(正確に言うと 33 点。見出し用に切り捨てました)。
3番目。間違いなく建物のどこかにあると思われる情報を見つけるテストです。 3 回のうち 2 回は、市場で最も先進的な AI が存在する答えを探し求め、間違った答え、または確信を持ってナンセンスを返してきました。
2 階上のマーケティング部門が自律型デジタル従業員を売り込んでおり、研究部門はデジタル従業員が約 3 分の 1 の確率で適切なファイルを見つけるという証拠を発表したところだからです。
しかし、楽譜のせいで夜も眠れなくなるわけではありません。その下に 2 つの発見があります。
1 つ目は、Salesforce の研究者が書き留めた診断です。ボトルネックは考え方ではなく、発見です。モデルはうまく推論できるかもしれませんが、単に推論するための適切な材料を取得できなかっただけです。証明はその単純さゆえに残酷だ。研究者たちがシステムの探索をやめ、その代わりにモデルに会社の文書を完全に渡したところ、最高のモデルはその惨めな 3 分の 1 から 76.55 に跳ね上がりました。同じモデルです。同じ質問です。変わったのは、証拠を見つけなければならないか、それとも証拠を渡されるかだけだった。
これを二度読んでください。これは今年エンタープライズ AI に関して発表された最も重要な文章ですが、ほとんど誰も気づかなかったからです。モデルは決して問題ではなかったのです。高価な部分、つまり誰もが今後 10 年かけてさらに購入するために GPU を食い荒らす巨大な脳を抵当に入れている部分は、完全に機能し、足を軽くたたきながら、安価で鈍く、地味な検索が行われるのを待っています。

正しい段落を持ってきてください。そして、検索層ではそれができません。
これがボールゲームの全体であり、前回お話ししたところと全く同じところに着地します。私は、エンタープライズ AI の 3 分の 2 は、実際には知性を着飾った検索であると主張しました。ここに Salesforce がいます。友好的な証言者ではありませんが、売り込み全体がその逆が真実であるかどうかに依存している企業です。この企業は、検索こそがまさに企業が崩壊するところであること、そして、モデルがすでに十分に優れていたため、より大きく、より賢く、より飢えたモデルがあなたを救ってくれるわけではないことを確認しています。
2 番目の発見は、私が最もひどいと思うもので、データセット自体の構造の中に隠れています。 HERB の 1,514 の質問のうち、回答があるのは 815 のみです。残りの 699 件 (ほぼ半数) は、設計上応答不可能です。 Salesforce は、シミュレーションされた企業のどこにも裏付けとなる証拠が存在しない、完全に合理的と思われる質問を意図的に何百も書き、AI が知らなかったと認めるかどうかを観察しました。
それが何を意味するのか考えてみましょう。 HERB は、システムが答えを見つけられるかどうかをテストするだけではありません。その半分近くは、システムが、それが存在しないことを認識しているかどうかをテストするものです。もっともらしい質問を与えられても、それを裏付ける事実がない場合に、適切に聞こえるものを作成する代わりに、「それは見つかりません」と答えるだけの背骨があるかどうかです。これは企業が AI に求める最も重要な動作であり、市場のほとんどのシステムが確実に備えている動作ではありません。その代わりに彼らが何をするのかについて、私たちは愛称さえ持っています。私たちはそれを幻覚と呼んでいます。あたかもそれが、テクノロジーを重要な仕事に使用できなくする正確なものではなく、魅力的な癖であるかのように呼んでいます。
したがって、2 つの調査結果をまとめてください。前者に対する業界の答えは「より大きな頭脳を購入する」ですが、データは役に立たないと言っています。そして、

後者を修正する頭脳は購入できません。なぜなら、自信を持って答えを生み出すことは知性の不足ではありません。それは、自身の知識の限界を知るために決して構築されていないアーキテクチャの特性だからです。
ここで詳しくは説明しませんが、答えは私にあります。
誰かがその仮定を拒否したとします。彼らが、回収は発電機にホッチキスで留められる配管ではなく、メインイベントであると決定したとします。実際のマシンはゼロから構築され、前回お話しした安価でクールで豊富なシリコンで動作し、高価な頭脳は、何かを見つけるのではなく真に生成しなければならないまれな瞬間に備えて保管されています。そして、同じシステムが、証拠によって裏付けられるものの境界を知るために、つまり 815 に答えるのと同じくらいすぐに 699 に対して「わかりません」と言えるように、根本から設計されたとします。
そして、誰かが HERB (Salesforce 独自の残酷で公開された容赦のないテスト) を受験し、実行したとします。
これだけは言っておきます。 3点目を奪えなかった。僕らは40点も取れなかった。私たちは、同じ公開ベンチマークで Salesforce の最高のシステムが到達できる上限を 2 倍以上に引き上げました。そして、ベンチマークの厳しい側が実際に要求していること、つまりいつ口を閉ざすべきかを知ることを尊重しながらそれを実現しました。私たちの数値は本物であり、誰でもダウンロードできる同じデータに対して測定されたものであり、これまでの 3 年間の大型化した GPU が明らかに果たせなかったことを実現しています。
そして、いいえ、尋ねる前に、それは 100 ではありません。100 であると言う人には私は警戒します。 HERB のほぼ半数にはまったく答えがないことを忘れないでください。このようなテストで満点を記録するシステムは、まだ知恵に達していません。ハッタリで引っ掛け問題を乗り越える方法を学びました。完璧は決して目標ではなかった

— 地球上で最高のモデルが、すべての文書を完全に手渡したにもかかわらず、まだ 70 年代半ばになってようやく世に出たことを思い出してください。ターゲットは異なります。証拠があるときは正しく、証拠がないときはそう主張するシステムです。そして、そのようなシステムは、その正直さに対してポイントで報われます。スコアカードでは「わからない」と「間違った」の区別がつかないからです。この数値と 100 の間の距離は、マシンの故障ではありません。その多くは、マシンが嘘をつかないことです。マシンと一緒に座っているとき、それがこのエクササイズの要点です。
それは質問を変えます。業界は長い間「どのモデルを?」と問い続けてきましたが、その下に「どのアーキテクチャ?」という事前の質問があったことを忘れていました。 HERB は、モデルの問題はほぼ解決され、ほぼ的外れであること、つまり次の 10 年は GPU 内部ではなく、検索レイヤーと正直レイヤーで決定されることを Salesforce が偶然認めたものです。
私が傍観者ではないことを知っておくべきです。私は、まさにこのコラムの異端に基づいて設立された小さな会社を共同設立しました。つまり、検索がメインイベントであり、それは安価なシリコンに属し、システムはいつ「分からない」と言えるかを知っているべきであるというものです。したがって、それに応じて私の熱意を検討してください。
しかし、私の利益相反が触れられないことに注目してください。 32.96 は Salesforce の番号であり、私の番号ではありません。取得がボトルネックであるという診断は、私ではなく Salesforce の研究者によって書かれました。ベンチマークのほぼ半分を回答不能にするという選択は、私のものではなく Salesforce のものでした。今年エンタープライズ AI について誰かが言った最も正直な言葉は、Salesforce の研究論文に埋もれた告白でした。皇帝の頭脳は素晴らしい、皇帝は自分のファイルを見つけることができず、約半分の時間は自分が何を知らないのかさえ分かっていません。
業界全体が聞いています

それで、もっと大きな脳を買いに行きました。これはコンピューティングの歴史の中で最も高価な間違いだと思います。その理由を説明します。
ランドール
2026 年 6 月 2 日午前 9 時 03 分 - 返信
arXiv 上の論文へのリンクがあると良いでしょう。
ジョエル
2026 年 6 月 2 日午前 10:23 - 返信
クイック Google で https://arxiv.org/pdf/2506.23139 にアクセスします。
AI：三成认罪 - 偏执的コード农
2026 年 6 月 2 日午後 4 時 31 分 - 返信
[…] 详情参考 […]
クリンジリーのかわいい女の子は誰ですか?アニーナと360度ファッション

## Original Extract

The Thirty Percent Confession | I, Cringely
Home
Last time I told you the AI industry is paying a tax it doesn’t have to pay — that a great deal of what we grandly call “AI” is really just looking things up, and we’ve chosen to do that looking-up on the most expensive silicon ever manufactured. A number of you wrote to say I was overstating it. Surely, you said, the people setting hundreds of billions of dollars on fire know something I don’t.
So this week I won’t argue with you. I’ll let one of the largest companies in enterprise software argue with you instead — because it already has, in a research paper it published itself and seems to have hoped you wouldn’t read too closely.
The company is Salesforce. The same Salesforce selling you “agents,” an “agentic enterprise,” a tireless digital workforce to set beside your human one. While one part of the building handled the marketing, another part — Salesforce AI Research, the people whose job is to measure things rather than sell them — built a test to find out how well today’s best AI can do something gloriously unglamorous: find the right piece of information when it’s scattered across the mess of a normal company. Slack threads. GitHub. Meeting transcripts. Documents nobody filed correctly. The stuff every real business actually runs on.
They named it HERB — the Heterogeneous Enterprise RAG Benchmark — and they didn’t build it on the cheap. It’s a synthetic but painstakingly realistic company: 530 employees across 30 products, generating 39,190 documents, messages, transcripts, and pull requests, strewn about the way they really would be. The paper is on arXiv. The data is on Hugging Face. Anyone can check my arithmetic, which is exactly why I’m happy to build a column on it.
When Salesforce turned the best agentic retrieval systems money can buy loose on HERB — top-tier models, the good stuff, with planning and tool use — they scored 32.96 out of 100. (Thirty-three, if we’re being precise; I rounded down for the headline.)
A third. On a test of finding information that is definitely, provably somewhere in the building. Two times out of three, the most advanced AI on the market went hunting for an answer that existed and came back with the wrong one — or with confident nonsense.
Sit with that, because two floors up the marketing department is selling you an autonomous digital employee, and the research department just published evidence that the digital employee finds the right file about a third of the time.
But the score isn’t the part that should keep you up at night. Two findings underneath it are.
The first is the diagnosis Salesforce’s own researchers wrote down: the bottleneck isn’t the thinking, it’s the finding. The models could reason fine — they simply couldn’t retrieve the right material to reason over. The proof is brutal in its simplicity. When the researchers stopped making the system hunt and instead handed the model the company’s documents outright, the best one leapt from that miserable third to 76.55. Same model. Same questions. The only thing that changed was whether it had to find the evidence or was handed it.
Read that twice, because it’s the most important sentence published in enterprise AI this year and almost nobody noticed: the model was never the problem. The expensive part — the giant, GPU-devouring brain everyone is mortgaging the next decade to buy more of — is sitting there perfectly capable, tapping its foot, waiting for the cheap, dull, unglamorous retrieval layer to bring it the right paragraph. And the retrieval layer can’t.
This is the whole ballgame, and it lands exactly where I left you last time. I claimed two-thirds of enterprise AI is really retrieval wearing intelligence as a costume. Here is Salesforce — not a friendly witness, but a company whose entire pitch depends on the opposite being true — confirming that retrieval is precisely where the enterprise falls apart, and that a bigger, smarter, hungrier model does not rescue you, because the model was already good enough.
The second finding is the one I find most damning, and it’s hiding in the dataset’s own structure. Of HERB’s 1,514 questions, only 815 have answers. The other 699 — nearly half — are unanswerable by design. Salesforce deliberately wrote hundreds of perfectly reasonable-sounding questions for which no supporting evidence exists anywhere in the simulated company, and then watched to see whether the AI would admit it didn’t know.
Think about what that means. HERB isn’t only a test of whether a system can find the answer. Nearly half of it is a test of whether the system knows when there isn’t one — whether, handed a plausible question and no facts to support it, it has the spine to say “I can’t find that” instead of manufacturing something that sounds right. That is the single most important behavior an enterprise needs from AI, and the one almost no system on the market reliably has. We even have a pet word for what they do instead. We call it hallucination, as though it were a charming quirk rather than the precise thing that makes the technology unusable for any job that matters.
So put the two findings together. The industry’s answer to the first is “buy a bigger brain,” which the data says won’t help. And there is no brain you can buy that fixes the second, because confidently inventing answers isn’t a shortage of intelligence — it’s a property of an architecture that was never built to know the edge of its own knowledge.
Which brings me to the part I’m not going to fully tell you here, though I have the answer.
Suppose someone refused the assumptions. Suppose they decided retrieval wasn’t plumbing to be stapled onto a generator but the main event — the actual machine, built from scratch, running on the cheap, cool, abundant silicon I told you about last time, with the expensive brain held in reserve for the rare moment something genuinely must be generated rather than found. And suppose that same system was designed, from the ground up, to know the boundary of what it can support with evidence — to say “I don’t know” on the 699 as readily as it answers the 815.
And suppose that someone took HERB — Salesforce’s own brutal, public, no-mercy test — and ran it.
I’ll tell you only this. We didn’t score a third. We didn’t score forty. We more than doubled the ceiling that Salesforce’s best systems could reach, on the identical public benchmark — and did it while honoring the thing the benchmark’s harder half actually demands: knowing when to keep their mouth shut. Our number is real, it was measured against the same data anyone can download, and it does what three years of ever-bigger GPUs have conspicuously failed to do.
And no — before you ask — it isn’t 100, and I’d be wary of anyone who told you it was. Remember that nearly half of HERB has no answer at all. A system that posts a perfect score on a test like that hasn’t reached wisdom; it’s learned to bluff its way past the trick questions. Perfection was never the target — recall that the best model on earth, handed every document outright, still only clawed its way into the mid-70s. The target is different: a system that’s right when the evidence is there and says so when it isn’t. And a system like that pays for its honesty in points, because a scorecard can’t tell the difference between “I don’t know” and “I got it wrong.” The distance between that number and 100 isn’t the machine failing. Much of it is the machine refusing to lie — which, when you sit with it, is the whole point of the exercise.
It changes the question. The industry has been asking “ which model? ” for so long it forgot there was a prior question underneath: which architecture? HERB is Salesforce’s accidental admission that the model question is largely settled and largely beside the point — that the next decade gets decided at the retrieval layer and the honesty layer, not inside the GPU.
You should know I’m not a bystander. I co-founded a small company built on exactly the heresy in this column — that retrieval is the main event, that it belongs on cheap silicon, and that a system ought to know when to say it doesn’t know. So weigh my enthusiasm accordingly.
But notice what my conflict of interest cannot touch. The 32.96 is Salesforce’s number, not mine. The diagnosis that retrieval is the bottleneck was written by Salesforce’s researchers, not me. The choice to make nearly half the benchmark unanswerable was Salesforce’s, not mine. The most honest thing anyone has said about enterprise AI this year was a confession buried in a Salesforce research paper: the emperor’s brain is magnificent, the emperor cannot find his own files, and about half the time he doesn’t even know what he doesn’t know.
The whole industry heard that and went out to buy a bigger brain. I think that’s the most expensive mistake in the history of computing, and I’m going to show you why.
Randall
June 2, 2026 at 9:03 am - Reply
A link to the paper on arXiv would be nice.
Joel
June 2, 2026 at 10:23 am - Reply
Quick Google leads to https://arxiv.org/pdf/2506.23139
AI：三成认罪 - 偏执的码农
June 2, 2026 at 4:31 pm - Reply
[…] 详情参考 […]
Who's the pretty girl with Cringely? Anina and 360 Fashion
