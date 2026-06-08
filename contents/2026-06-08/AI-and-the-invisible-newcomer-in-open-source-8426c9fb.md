---
source: "https://blog.stdlib.io/ai-and-the-invisible-newcomer-in-open-source/"
hn_url: "https://news.ycombinator.com/item?id=48446543"
title: "AI and the invisible newcomer in open source"
article_title: "What We're No Longer Seeing: AI and the Invisible Newcomer in Open Source"
author: "kgryte"
captured_at: "2026-06-08T16:28:52Z"
capture_tool: "hn-digest"
hn_id: 48446543
score: 2
comments: 0
posted_at: "2026-06-08T15:20:57Z"
tags:
  - hacker-news
  - translated
---

# AI and the invisible newcomer in open source

- HN: [48446543](https://news.ycombinator.com/item?id=48446543)
- Source: [blog.stdlib.io](https://blog.stdlib.io/ai-and-the-invisible-newcomer-in-open-source/)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T15:20:57Z

## Translation

タイトル: AI とオープンソースにおける目に見えない新参者
記事のタイトル: 私たちがもう見なくなったもの: AI とオープンソースの目に見えない新参者
説明: オープンソース コミュニティが新規参入者を確認し、歓迎するために常に依存してきた目に見える摩擦を AI がどのように吸収しているか。

記事本文:
私たちがもう見なくなったもの: AI とオープンソースにおける目に見えない新参者
オープンソース コミュニティが新規参入者を確認し、歓迎するために常に依存してきた目に見える摩擦を、AI がどのように吸収しているか。
昨年の冬、私たちのチームは NSF の POSE (オープンソース エコシステムを実現するための道) プログラムの一環として I-Corps トレーニングを受けました。 I-Corps は米国科学財団の顧客発見トレーニングであり、POSE はそれをオープンソース プロジェクトに適応させています。私たちにとって、これは 7 週間と、エコシステム内およびその周辺の関係者との 100 回以上のインタビューを意味しました。 (ここでの「us」は stdlib です。これは、数値計算と科学計算に重点を置いた、JavaScript および Node.js のオープンソース標準ライブラリです。)
これらのエコシステム発見インタビューには、売り込みをしないという厳格なルールが 1 つありました。 stdlib が何をするのか説明することも、決定を擁護することも、販売することもありません。
それが私たちがやったことです。カリキュラムでは、フレームワーク、キャンバス、面接の規律など、オープンソースのエコシステムがどのように維持されるべきかを学びました。インタビューは、誰もシラバスに載せていないことを私たちに教えてくれました。メンテナー、貢献者、ユーザーと話をすると、同じテーマが表面化し続けました。それは、人々がプロジェクトを見つけ、支援を受け、貢献する方法が AI によって変化しているということです。そして、私たちがそれを聞いたのはインタビューだけではありませんでした。カンファレンスや廊下での会話、ブログ投稿や LinkedIn のスレッドで、開発者との関係やコミュニティ活動を生業とする仲間たちが、それぞれの角度から同じ変化について説明していました。
私たちはエコシステムについて学びに行きました。私たちが見ている間に、私たちの周りの生態系が変化していました。
この投稿は私たちが見ていたものについてのものです。それは私たちがそうでなかったことについてでもあります。
長い間、オープンソース コミュニティは特定の種類の目に見える摩擦に依存してきました。人々は行き詰まってしまい、ショックを受ける

スタック オーバーフロー、GitHub の問題、メーリング リスト、フォーラム スレッドで問題が解決しました。公の場で助けを求めるという最初の行為自体が参加の一形態でした。それはあなたが他の人のコミュニティの端を見つける方法であり、他の人があなたを見つける方法でもありました。
チャールズ・ヴォーグルの『コミュニティの芸術』には古いモデルがあります。これは、どのコミュニティにもその周囲に同心円があるという考えです。人々は訪問者から始まります。彼らは足場を見つけます。彼らは、何らかの形で、私がここで何をしているのかと尋ねます。外側のリングから内側への旅、つまり訪問者から会員、そして長老への道は、自動的には進みません。それには目に見える経路が必要です。それには喜んで指導する人々が必要です。そして重要なことに、その道は他の人に知られている必要があります。
このモデルは Vogl やオープンソースに固有のものではありません。下の図は、ほとんどの読者が肌を持たないコミュニティ、つまり 1990 年代初頭の娯楽目的のスキューバ ダイバーに関する研究から得たものです。エドゥアール・ラガシュは、正当な周辺参加（ジャン・レーヴとエティエンヌ・ヴェンゲルが作った造語）を基礎にして、同じ幾何学模様を地図に描きました。中心には昔ながらの人々、最も外側のリングにはカジュアルな観光客、内側に引っ張るラベル付きの軌道が描かれています。人々はコミュニティの端でリスクの低いことをすることでコミュニティのメンバーになります。それは、周囲に自分たちを可視化する構造があるために存在するものです。 90 年代初頭のスキューバ ダイビングは一例にすぎませんが、オープンソース、スポーツ、ファンダム、名前を挙げられるあらゆるコミュニティに同じ力学と形が存在します。細部は異なりますが、ジオメトリは同じです。
構造的に言えば、オープンソースで機能するのは摩擦、つまり目に見える摩擦でした。誰かが行き詰まって現れ、双方がお互いを見ることができました。その最初の公開質問は一種のイニシエーションであり、またそれは活動への合法的な周辺参加でもありました。つまり、リスクが低く、目に見える形で参加する方法でした。

ああ、端です。この力学は、数十年にわたるオープンソース コミュニティ構築の実践を通じて織り込まれており、特に、新規参入者がオープン プロジェクトの貢献者になるのを助けるための 10 の簡単なルールの下にあります。公開質問によって新参者が目立つようになり、すでに内部にいる誰かが回答する機会が生まれました。
また、コミュニティが初心者にも見えるようになりました。ああ、ここにあるよ。人がいる。入る道があるよ。
しばらくの間、私たちが耳にし続けたこの変化は、発見というよりは雰囲気であり、チャート上でそれを示す前に会話の中で感じることができたものでした。しかし、それは目に見える形で現れ続けました。今年の FOSDEM で、Grafana で開発者関係を率いる David Allen 氏は全員に自分の数字を示し、一部のメンテナーはまだ経験していないが、多くのメンテナーは経験するであろう瞬間を説明しました。
彼は、四半期ごとに 8 ～ 18 パーセントのコミュニティの成長をほぼ自動的に期待することに慣れてきました。私は眠りにつくことができ、そうすればコミュニティは成長するだろうと彼は言いました。その後、4分の1では27％減少した。彼らはデータをチェックし、再度実行しましたが、それは測定エラーではありませんでした。全体的なユーザーの増加はまだ続いていたにもかかわらず、コミュニティに参加する人の数は減少していました。その仕組みも不思議なものではありませんでした。コミュニティ スペースへのオーガニック トラフィックは約 30% 減少しました。これは、LLM への質問と回答の行動が原因であると考えられます。
そして、それは 1 つのプロジェクトだけの話ではありません。 PNAS Nexus の 2024 年の研究では、影響を分離する明確な方法が見つかりました。英語での Stack Overflow アクティビティ (ChatGPT が利用可能な場合) と、ロシア語および中国語の対応するアクティビティ (ChatGPT が利用できない場合)、および数学フォーラム (LLM が弱い場合) と比較するというものです。 ChatGPT のリリースから 6 か月以内に、英語での活動はこれらのコントロールと比較して 25% 減少しました。

またはそれを下限と呼びます。減少は初心者や重複した質問に限定されませんでした。経験豊富なユーザーも経験の浅いユーザーも同様に投稿が少なくなります。
そしてそれは 2024 年のことでした。それ以来何が起こったのかを知るのに自然実験は必要ありません。LLM 支援開発は目新しいものからデフォルトになり、以前は公の場で尋ねられていた質問の多くは非公開で答えられるようになりました。
アレンと共同プレゼンターのアマンダ・ワグナーは、AI が原因ではないという私が何度も読み返す枠組みを提案しました。それは加速度的に、すでに脆弱だった基盤を露出させます。
ここで私が主張したい議論がさらに鮮明になります。
表面損失は信号です。以前は GitHub の問題やフォーラムのスレッドとして表面化していた質問が、現在は非公開で回答されるようになりました。それは現実であり、測定可能であり、何もないわけではありません。しかし、それは実際に何が起こっているかを過小評価しています。
AI が最初の摩擦を処理すると、新規参入者のブロックが解除されます。それらは機能的価値を獲得します。彼らは進歩します。彼らは問題を解決します。しかし、彼らは何のつながりもなく、誰にも見られず、始める瞬間もなく、それを理解します。これは本質的に悪いことではありません。ただ違うんです。しかし、私たちが常に「バグ」として対処してきた摩擦、つまり私たちが滑らかにしたいと思っていたものは、それがなくなるまで私たちが理解できない働きもしていました。
コミュニティとの関わりを深めるものとは何か、答えを得ることが旅の始まりに変わるものについて考えてみましょう。彼らはそれを選択し、その中で進歩し、そこにいる人々とのつながりを感じます。 AI は、驚くべき効率で選択と進歩を提供しますが、接続を複製するわけではありません。
そして、十分に語られていない部分は次のとおりです。
新規参入者は、自分がコミュニティの端にいることに気づいていません。彼らは AI 仲介者を通じてあなたのソフトウェアを使用しており、彼らは自分自身のことを考えていない可能性があります

プロジェクトとはまったく関係がなく、ツールとのみ関係しているのと同じです。彼らはあなたがコミュニティとして存在していることを知りません。彼らはあなたが依存関係として存在していることだけを知っています。彼らは、自分たちがすでにエコシステムの一部であることを知りません。
そして、あなたは彼らがそこにいることを知りません。
Vogl の内輪モデルは両方向の可視性に依存します。人は目に見える場合にのみ内側に誘うことができます。道があることを知っている場合にのみ、彼らは旅を始めることができます。プロジェクトとの最初の対話が AI レイヤーを通じて行われるとき、どちらの条件も確実に満たされません。
これは、潜伏者に関する古い問題とは別の問題です。潜伏者は静かでしたが、存在していました。彼らはあなたの空間にいたのです。あなたに見えなくても、彼らにはあなたのことが見えていました。 AI アシスタントを通じてあなたのライブラリを見つけた観​​察者は、あなたのスペースをまったく訪問しない可能性があります。
誤解のないように言っておきますが、誰もが内向きに進みたいわけではありませんし、それは問題ありません。フォーグルはこれについて直接書いています。どのコミュニティのほとんどのメンバーも外環に留まり、健全なコミュニティは彼らのためのスペースを作ります。ユーザーはプロジェクトに対して何の義務も負いません。しかし、オープンソースのエコシステムは誰もが維持できるわけではありません。それらは内向きに旅をする少数の人たちによって支えられています。そしてそれが、外側のリングが見えなくなることによる本当の代償なのです。目に見えないユーザーがすべて、失われた貢献者であるというわけではありません。それは、誰がそこにいるのかがもはや分からないということです。つまり、誰を招待すべきかが分からなくなっているのです。
つまり、実際に失われているのは、サポート トラフィックやページ ビュー、その他の表面的な指標ではありません。それは成熟の初期の瞬間であり、新しいメンバーをマークし、到着を知らせ、帰属のプロセスを開始する活動です。かつては一般質問もそのような瞬間の一つであり、誰もそのような意味で質問することはほとんどありませんでした。それは誰かの姿を現し、ドアを開け、反応するための条件を作り出しました。
そして私たちは聞き続けました

g 問題のまったく異なる分野に取り組んでいる人々によるこれのバージョン。アビゲイル・カブノック・メイズ氏は、FOSDEM での指導について語り、「寄付の量は増加しているが、信号対雑音比は実際に低下している」とはっきりと述べました。コミュニティは、コミット、PR、ダウンロードなどの表面的な指標によって生きているように見えますが、その下では帰属意識の基盤が空洞化しています。
それから数ヶ月が経っても、この状況は少しも沈静化しなかった。メンテナたちは経験を共有しています。ダニエル・ステンバーグ氏は、AI 時代がどのようなものであるかをカール内部から記録しており、フランク・ナイホフ氏は最近、Home Assistant のバージョンについて次のように述べています。「コントリビュータは増加しているが、メンテナはまだ検証のボトルネックになっている」ため、誰がレビューするよりも早くコントリビュートが到着します。これに応じて、エコシステム全体のプロジェクトが貢献ポリシーを再考しています。詳細は異なります。形は同じです。
そして、その根底にある不快な認識は次のとおりです。私たちは、必要であることさえ知らなかった摩擦に依存していたのです。オープンソース コミュニティを構築し維持するためのハンドブック (この問題を真剣に受け止めた多くの人々によって数十年にわたって蓄積されたもの) では、新規参入者が独自に注目を集め続けることが想定されていました。その仮定を書き留める必要がなかったので、誰もその仮定を書き留めませんでした。それが失敗しつつある今、戦略を変更する必要があります。
これは設計上の問題であり、危機ではありません
何かが侵食されると、本能はそれを嘆いたり、修復しようとしたりするものです。より有益な答えは、「以前は副産物として発生していたものを、意図的に構築する必要があるものは何ですか?」と尋ねることです。
あるトラックではプログラムから、もう一方ではインタビューや同僚から、私たちが学んでいたことはすべて、同じいくつかの答えに収束し続けました。
これを経て形になったコミュニティが、

ツールや賢いオンボーディング フロー、特定のプラットフォームではなく、人間関係と物語を重視しています。コミュニティの下にあるテクノロジー スタックは置き換え可能です。置き換えられないのは、そこに参加している人々がお互いに、そしてプロジェクトと本当の関係を持っているかどうか、そしてプロジェクトに誰かが入り込めるようなそれ自体のストーリーがあるかどうかです。それは柔らかいですね。それは難しい部分です。
その一部が、アレンとワグナーが努力の証拠と呼ぶものです。表面レベルのエンゲージメントを生み出すコストが下がると、もっともらしいと思われる問題、まとめられたPR、情報に通じているように聞こえるコメントなど、誰かが実際にここにいて、実際に注意を払い、実際にコミットしているという合図は、どこか別の場所から来なければなりません。将来の貢献者に投資を実証してもらうことの重要性は、それ以上に重要になっています。門番としての動きではない。本当の参加を再び可視化する方法として。プロジェクトの条件に従って何かをするよう人々に求めるコミュニティ (電話に出てくる、小さな責任を引き受ける、誰も書きたがらないドキュメントを書く) は貴重ではありません。彼らは、古い摩擦が提供していた可読性を再構築しています。
もう 1 つの変化は、異なる部屋の異なる人々によってほぼ同じ言葉で名前が付けられているのを聞いたものですが、抽出からのものです。

[切り捨てられた]

## Original Extract

How AI is absorbing the visible friction that open-source communities have always relied on to see—and welcome—newcomers.

What We're No Longer Seeing: AI and the Invisible Newcomer in Open Source
How AI is absorbing the visible friction that open-source communities have always relied on to see—and welcome—newcomers.
Last winter, our team went through I-Corps training as part of the NSF's POSE—Pathways to Enable Open-Source Ecosystems—program. I-Corps is the National Science Foundation's customer-discovery training, and POSE adapts it for open-source projects. For us, that meant seven weeks and more than a hundred interviews with stakeholders in and around our ecosystem. (The "us" here is stdlib , an open-source standard library for JavaScript and Node.js, with an emphasis on numerical and scientific computing.)
These ecosystem-discovery interviews had one hard and fast rule: you don't pitch. No explaining what stdlib does, no defending decisions, no selling.
So that's what we did. The curriculum taught us how open-source ecosystems are supposed to be sustained—the frameworks, the canvases, the interview discipline. The interviews taught us something nobody had put on the syllabus. Talking with maintainers, contributors, and users, the same theme kept surfacing: AI was changing how people found projects, got help, and contributed. And we weren't just hearing it in the interviews. At conferences and in hallway conversations, in blog posts and LinkedIn threads, peers who do developer relations and community work for a living were describing the same shift from their own angles.
We went to learn about our ecosystem. While we were looking, the ecosystem was changing around us.
This post is about what we were seeing. It's also about what we weren't.
For a long time, open-source communities have depended on a specific kind of visible friction. People got stuck and they showed up—on Stack Overflow, in GitHub issues, on mailing lists, in forum threads. That first act of asking for help in public was itself a form of participation. It was the way you found the edge of someone else's community, and the way they found you.
There's an old model from Charles Vogl's The Art of Community —the idea that any community has concentric rings around it. People start as visitors. They find their footing. They ask, in one way or another, what am I doing here? The journey from the outer ring inward—visitor to member to elder—isn't automatic. It requires visible pathways. It requires people willing to guide. And, crucially, it requires that the path be known to others.
That model isn't unique to Vogl, or to open source. The figure below comes from a study of a community most readers have no skin in: recreational scuba divers in the early 1990s. Edouard Lagache, building on legitimate peripheral participation (a term coined by Jean Lave and Étienne Wenger), mapped the same geometry—old-timers at the center, casual sightseers at the outermost ring, a labeled trajectory pulling inward. People become members of a community by doing low-risk things at its edges—things that exist because there's a structure around them that makes them visible. Early 90s scuba diving is just an example, but the same dynamic and shape exists in open source, in sports, in fandoms, in any community you can name. The details differ, but the geometry is the same.
What made it work in open source, structurally speaking, was friction—visible friction. Someone got stuck, showed up, and both sides could see each other. That first public question was a form of initiation, and it was also legitimate peripheral participation in action: a low-risk, visible way of taking part from the edge. That dynamic is woven through decades of open-source community-building practice—it's underneath Ten simple rules for helping newcomers become contributors to open projects , among others. The public question made the newcomer visible, and it created the opportunity for someone already inside to respond.
It also made the community visible to the newcomer. Oh—there's a there here. There are people. There's a way in.
For a while, the shift we kept hearing about was a vibe more than a finding—something you could feel in the conversations before you could point to it on a chart. But it kept turning out to be tangible. At FOSDEM this year, David Allen, who leads developer relations at Grafana, showed everyone his numbers , describing a moment some maintainers haven't had yet but many will.
He'd gotten comfortable expecting community growth of 8 to 18 percent a quarter, almost automatically— I could fall asleep, he said, and the community would grow. Then one quarter the number dropped 27 percent. They checked the data, ran it again, and it wasn't a measurement error: the number of people entering the community had collapsed, even though overall user growth was still going up. The mechanism wasn't mysterious, either. Organic traffic to community spaces was down roughly 30 percent, traceable to question-and-answer behavior moving to LLMs.
And it's not a one-project story. A 2024 study in PNAS Nexus found a clean way to isolate the effect: compare Stack Overflow activity in English (where ChatGPT was available) with Russian- and Chinese-language counterparts (where it wasn't), and with mathematics forums (where LLMs are weaker). Within six months of ChatGPT's release, English-language activity had dropped 25 percent relative to those controls—and the authors call that a lower bound. The decline wasn't limited to beginners or to duplicate questions. Experienced and inexperienced users alike posted less.
And that was 2024. You don't need a natural experiment to know what's happened since: LLM-assisted development has gone from novelty to default, and more of the questions that used to be asked in public are answered in private.
Allen and his co-presenter, Amanda Wagner, offered a framing I keep coming back to: AI isn't the cause of this. It's an accelerant, exposing foundations that were already fragile.
Here's where the argument I want to make sharpens.
The surface loss is signal. Questions that used to surface as GitHub issues or forum threads are now answered privately. That's real, and it's measurable, and it's not nothing. But it understates what's actually happening.
When AI handles the initial friction, the newcomer gets unblocked. They get functional value. They make progress; they solve the problem. But they get it without connection—without anyone seeing them, without a moment of initiation. This isn't inherently bad. It's just different. However, the friction that we've always addressed as a "bug"—the thing we wanted to smooth out—was also doing work we didn't understand until it was gone.
Think about what deepens anyone's engagement with a community—what turns getting an answer into the start of a journey: they choose it, they're making progress in it, and they feel some connection to the people in it. AI delivers choice and progress with remarkable efficiency, but it doesn't replicate connection.
And here's the part that doesn't get said enough:
The newcomer doesn't know they're at the edge of your community. They're using your software through an AI intermediary, and they may not think of themselves as being in relationship with a project at all—only with a tool. They don't know you exist as a community. They only know you exist as a dependency. They have no idea that they're already part of your ecosystem.
And you don't know they're there.
Vogl's inner-rings model depends on visibility in both directions. People can only be invited inward if they can be seen. They can only begin the journey if they know there's a path. When the first interaction with a project happens through an AI layer, neither condition is reliably met.
This is a different problem than the old problem of lurkers. Lurkers were quiet, but they were present. They were in your spaces. They could see you, even when you couldn't see them. The observer who finds your library through an AI assistant may never visit your spaces at all.
To be clear: not everyone wants to move inward, and that's fine. Vogl writes about this directly —most members of any community stay in the outer rings, and a healthy community makes room for them. Users don't owe a project anything. But open-source ecosystems aren't sustained by everyone; they're sustained by the few who do make the journey inward. And that's the real cost of the outer rings going invisible. It's not that every unseen user is a lost contributor. It's that we no longer know who's out there—which means we no longer know whom to invite.
What's actually being lost, then, is not support traffic or page views or the other surface metrics. It's the early moments of maturation—the activities that mark a new member, signal arrival, and begin the process of belonging. The public question used to be one of those moments, almost without anyone meaning it that way. It made someone visible, opened a door, created the conditions for response.
And we kept hearing versions of this from people working entirely different corners of the problem. Abigail Cabunoc Mayes, talking about mentorship at FOSDEM , put it plainly: "the volume of contributions is going up, but the signal-to-noise ratio is really going down." A community can look alive by the surface metrics—commits, PRs, downloads—while the substrate of belonging is hollowing out underneath.
The months since haven't quieted any of this. Maintainers are sharing their experiences—Daniel Stenberg has been chronicling what the AI era looks like from inside curl , and Franck Nijhof recently described Home Assistant's version of it : contributions arriving faster than anyone can review them, because "contributors are being amplified, but maintainers are still the verification bottleneck." Projects across the ecosystem are rethinking their contribution policies in response. The particulars differ; the shape is the same.
And here's the uncomfortable realization underneath all of it: we were relying on a friction we didn't even know we needed . The playbook for building and sustaining open-source communities—accumulated over decades by a lot of people who took the question seriously—assumed newcomers would keep becoming visible on their own. Nobody wrote that assumption down, because nobody had to. Now that it's failing, the playbook has to change.
This is a design problem, not a crisis
The instinct, when something erodes, is to mourn it or try to restore it. The more useful response is to ask: what has to be built intentionally that used to happen as a byproduct?
Everything we were learning—from the program on one track, from our interviews and our peers on the other—kept converging on the same few answers.
The communities that come through this in shape will be the ones oriented around relationships and narrative —not tooling, not clever onboarding flows, not any one platform. The technology stack underneath a community is replaceable. What isn't replaceable is whether the people in it are in real relationship with each other and with the project, and whether the project has a story about itself that someone can find their way into. That sounds soft. It is the hard part.
A piece of that is what Allen and Wagner call proof of effort. When the cost of producing surface-level engagement falls—an issue that reads as plausible, a PR that compiles, a comment that sounds informed—the signal that someone is actually here, actually paying attention, actually committed, has to come from somewhere else. Getting prospective contributors to demonstrate investment becomes more important, not less. Not as a gatekeeping move. As a way of making real participation visible again. Communities that ask people to do something on the project's terms —show up to a call, take a small responsibility, write the doc nobody wants to write—are not being precious. They are reconstructing the legibility the old friction used to provide.
The other shift—one we heard named, in almost the same words, by different people in different rooms—is from extractiv

[truncated]
