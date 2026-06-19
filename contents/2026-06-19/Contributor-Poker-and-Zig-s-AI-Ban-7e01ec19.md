---
source: "https://kristoff.it/blog/contributor-poker-and-ai/"
hn_url: "https://news.ycombinator.com/item?id=48595826"
title: "Contributor Poker and Zig's AI Ban"
article_title: "Contributor Poker and Zig's AI Ban | Loris Cro's Blog"
author: "birdculture"
captured_at: "2026-06-19T07:59:20Z"
capture_tool: "hn-digest"
hn_id: 48595826
score: 2
comments: 0
posted_at: "2026-06-19T07:30:47Z"
tags:
  - hacker-news
  - translated
---

# Contributor Poker and Zig's AI Ban

- HN: [48595826](https://news.ycombinator.com/item?id=48595826)
- Source: [kristoff.it](https://kristoff.it/blog/contributor-poker-and-ai/)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T07:30:47Z

## Translation

タイトル: 寄稿者ポーカーと Zig の AI 禁止
記事のタイトル: 寄稿者 ポーカーと Zig の AI 禁止 |ロリス・クロのブログ
説明: モンティ ホールをプレイするときは、常にドアを変更する必要があります。

記事本文:
寄稿者 Poker と Zig の AI 禁止
2026 年 4 月 29 日
•
6
最小読み取り時間
ロリス・クロ
Zig Software Foundation での在職中、私はソフトウェアについて多くの興味深いことを学ぶ機会がありました。今日私が共有したいことは、貢献者を惹きつけるほど大きなオープンソース プロジェクトにとって重要な理解事項です。
オープンソース開発には微妙なメリットとデメリットがあり、そこから得られる良い点を活用して、たくさんある悪い点に対処する必要があることを補うかどうかはあなた次第です。
何よりもまず、オープンソースは多くのビジネス モデルと互換性がなく、価値のあるものを無料で提供しなければならないという考えに基づいています。もちろん、ものを無料で配布する代わりに、主にコードの貢献という形で、ものを無料で入手することもできます。
残念ながら、これらの貢献（今後は PR と呼びます）は請求額を支払わないだけでなく、それ自体が余分な労力と摩擦の原因となることがよくあります。実際、PR 作成者と協力してコードをマージ可能な状態にするよりも、メンテナが特定の変更を直接実装するほうが手間が少なく済むことはよくあります。
これまで述べてきたことを踏まえると、オープンソース開発はかなりくだらない取​​引のように思えますが、それでも私は、オープンソース開発には、ほとんどの代替開発モデルよりもはるかに安価に、高品質の製品を生み出す可能性があると強く信じています。必要なのは、オープンソース ゲームに慣れることだけです。このテーマ全般についてはすでに書きましたが、今日は 1 つの特定の側面、コントリビューター ポーカーに焦点を当てます。
オープンソース プロジェクトが成功すると、最終的には処理能力を超える PR が得られるようになります。これまで述べてきたことを考慮すると、ACを停止するのは理にかなっています

仕事の ROI を最大化するために不完全な PR を受け入れることもありますが、Zig プロジェクトではそれは行いません。その代わりに、私たちは、たとえ彼らがそこに到達するまでに何らかの手助けを必要とするとしても、新しい寄稿者が自分の作品を投稿できるよう最善を尽くしています。私たちがこれを行うのは、それが「正しい」ことだからだけではなく、賢明な行いだからでもあります。
オープンソース プロジェクトへの貢献は反復的なゲームであり、貢献者がプロジェクトにもたらす価値の大部分は、後の反復にあります。言い換えれば、あなたは最初に新しい貢献者を受け入れるためにある程度のエネルギーを投資し（つまり、賭けをし）、後でその貢献者がより​​信頼され、多作になるにつれてその関係があなたに返済し始めることを期待します。
私がこれを「コントリビューター ポーカー」と呼ぶ理由は、実際のカード ゲームについてよく言われるように、「カードをプレイするのではなく、人をプレイする」からです。コントリビューター ポーカーでは、最初の PR の内容ではなく、コントリビューターに賭けます。
このダイナミクスを明確に理解することで、Zig プロジェクトは時間の経過とともに莫大な価値を生み出すことができました。コンパイラ ツールチェーンをゼロから構築することは、貢献者からの多大な支援がなければカバーすることが不可能な広大な範囲に及びます。
Ryan Liptak などの貢献者のおかげで、Zig は Windows リソース スクリプト ( .rc ) ファイルをコンパイルできるため、Zig ユーザーは Windows 上で実行可能アイコン (およびその他) を設定する贅沢を享受できるようになりました。
もう 1 つの注目すべき例は、Frank Denis による貢献です。彼が std.crypto で行った仕事に金銭的価値を与えるにはどこから始めればよいのかさえわかりません。
Zig の初期の頃は、すべての新しいコントリビューターに投資することが可能でしたが、今ではプロジェクトは成長し、受信する PR の量が、コントリビューターのエネルギー コアの量をはるかに超えるまでに成長しました。

コントリビューター ポーカーをプレイすることを提案します。
これは実際には、優れた PR が長期間レビューされずに放置されている例があり、貴重な貢献者が Zig への貢献への関心を失う可能性があることを意味します。
これは私たちが年次財務報告書で述べたことであり、さらに重要なことに、これは私たちが積極的に認識している問題であり、将来的には解決（または少なくとも緩和）したいと考えています。
残念ながら、これは本質的に取り組むのが難しい問題であるだけでなく、AI によって事態はさらに悪化しています。
Zig プロジェクトが AI の貢献を禁止する理由については多くの憶測がありましたが、コントリビューター ポーカーの重要性を理解した今では、なぜ私たちがそうするのかが簡単に理解できます。
影響力のある作品を提供できるようにするには、貢献者はコードベースと問題空間に精通している必要があり、たまたま CI に合格したランダムなソリューションを提出するのではなく、最適なアプローチを追求するために PR によって導入されたすべての変更を熟考しているという点でコア チームから信頼される必要があります。
さらに、信頼性を高めるプロセスの一環として、貢献者は、コードがマージされた後もしばらくの間、提出したコードに対して責任を負うことが期待されます。完璧な人はいませんし、事後的に問題が見つかることもあります。軌道修正方法を決定するためのフォローアップのディスカッションは、特定の問題領域について重要な洞察を構築したエンジニアとの関係を繰り返すことの価値を示すもう 1 つの例です。
これは重要です。なぜなら、Zig のユーザーは、できる限り優れた言語とツールチェーンを提供してくれる Zig Software Foundation に賭けているからです。
残念ながら、LLM ベースの貢献の現実は、バックグラウンドの増加により、私たちにとってほとんどマイナスです。

幻覚に満ちた価値のないドライブバイ PR (CI に合格するどころか、コンパイルすらできない) によるノイズから、非常識な 10,000 行の長さの初回 PR まで。その間に、表面的には問題ないように見える PR も多数受け取りました。その中には、LLM を使用していないと明示的に主張しているものもありましたが、その後の議論で、著者がこっそりと LLM に相談し、その間違いだらけの返信を私たちに吐き戻していることがすぐに明らかになりました。
誤解のないように言っておきますが、ここで重要なのは、これが AI のすべてであると信じているということではありません。そうではありません。これは明らかにツールの誤用ですが、私たちのプロジェクトにおける LLM ベースの貢献の圧倒的多数がそうであったものでもあります。
したがって、理論的には、LLM を利用する有効な寄稿者になることができますが、寄稿者ポーカーの観点から見ると、このリスク要因を提示しない他の寄稿者が膨大に存在する中で、LLM ユーザーに賭けるのは単純に非合理的です。
寄付が LLM からのものであるかどうかを知ることは不可能であると発言した人々は、このポリシーの要点を完全に見逃しており、明らかに寄付者ポーカーについて認識していません。
私たちにとって、貢献者がシステム思考を改善し、他の有能で信頼できる多作なエンジニアと交流できる魅力的なエコシステムを提供できることは、私たちのビジネス モデルの重要な側面です。
以前にも述べたように、Zig はプロジェクトを取り巻く技術的、社会的、ビジネス管理上の問題を考えることに多大な労力を費やしているため、その重量資金クラスを大きく上回る成果を上げることができます。
コントリビューター ポーカーは私たちの戦略の重要な部分であり、ゲームを効果的にプレイする能力を妨げるものに対して抵抗することがプロジェクトの最大の利益になります。そうは言っても、私たちは、

まだ多くの未解決の問題があり、より多くの洞察が得られ次第、ポリシーを調整する予定です。
私たちの成功を支援したい場合は、Zig Software Foundation への毎月少額の寄付をご検討ください。

## Original Extract

You should always change doors when playing Monty Hall

Contributor Poker and Zig's AI Ban
April 29, 2026
•
6
min read • by
Loris Cro
During my tenure at the Zig Software Foundation I’m having the opportunity to learn many interesting things about software. The one I want to share today is a key piece of understanding for any open source project big enough to attract contributors.
Open source development comes with a nuanced set of pros and cons, and it’s up to you to leverage the good that comes from it in order to compensate for having to deal with the bad, of which there’s plenty.
First and foremost, open source is incompatible with many business models and it’s based on the idea that you have to give away something of value for free. Of course in exchange for giving stuff out for free, you also get stuff for free, mostly in the form of code contributions.
Unfortunately not only do those contributions (let’s call them PRs from now on) not pay the bills, they themselves are often a source of extra labor and friction. In fact it’s pretty common that it would take a maintainer less effort to implement a given change directly, rather than work with a PR author to get their code to a mergeable state.
Based on what I mentioned so far, open source development seems a pretty shitty deal, and yet I firmly believe it has the potential to produce higher quality products and for significantly cheaper than most alternative development models ; you just have to get good at the open source game . I’ve already written on the subject in general, but today I’ll focus on one specific aspect: contributor poker.
In successful open source projects you eventually reach a point where you start getting more PRs than what you’re capable of processing. Given what I mentioned so far, it would make sense to stop accepting imperfect PRs in order to maximize ROI from your work, but that’s not what we do in the Zig project. Instead, we try our best to help new contributors to get their work in, even if they need some help getting there . We don’t do this just because it’s the “right” thing to do, but also because it’s the smart thing to do .
Contributing to an open source project is an iterated game and the majority of the value that a contributor can bring to a project lies in the later iterations. In other words, you initially invest some energy (i.e. place a bet) to onboard a new contributor, and you hope that later on that relationship starts paying you back as the contributor becomes more trusted and prolific.
The reason I call it “contributor poker” is because, just like people say about the actual card game, “you play the person, not the cards”. In contributor poker, you bet on the contributor, not on the contents of their first PR.
Having an explicit understanding of this dynamic has netted the Zig project a huge amount of value over time. Building a compiler toolchain from the ground up is a huge scope that would have been impossible to cover without significant help from contributors.
Thanks to contributors like Ryan Liptak Zig users can now enjoy the luxury of setting the executable icon (and more) on Windows because Zig can compile Windows resource script ( .rc ) files.
Another notable example are the contributions from Frank Denis . I wouldn’t even know where to begin to give a monetary value to the work he has done in std.crypto .
In the early days of Zig it was possible to invest on every new contributor, but now the project has grown to the point where the amount of incoming PRs far exceeds the amount of energy core contributors have at their disposal to play contributor poker.
In practice this means that there have been instances of good PRs that have gone un-reviewed for extended periods of time, potentially causing valuable contributors to lose interest in contributing to Zig.
This is something we have mentioned in our yearly financial reports , and more importantly it’s an issue we’re actively aware of, and that we hope to solve (or at least mitigate) in the future.
Unfortunately, not only is this an inherently hard problem to tackle, but AI has made things worse.
There has been a lot of speculation about why the Zig project bans AI contributions, but now that you understand the importance of contributor poker, it’s easy to see why we do it.
To be able to provide impactful work a contributor needs to be familiar with the codebase and the problem space, and they need to be trusted by the core team to have thought through all the changes introduced by their PRs in order to strive for an optimal approach, rather than just submitting a random solution that happens to pass CI.
Additionally, as part of the process of becoming more trusted, contributors are expected to be responsible for the code they submit for a while more after their code is merged. Nobody is perfect and sometimes issues are discovered after the fact. Follow up discussions to decide how to course-correct are another example of the value of having an iterated relationship with engineers that have built significant insight on a given problem space.
This is important because users of Zig have in turn bet on the Zig Software Foundation to provide them with a language and toolchain that strives to be as good as we can make it.
Unfortunately the reality of LLM-based contributions has been mostly negative for us, from an increase in background noise due to worthless drive-by PRs full of hallucinations (that wouldn’t even compile, let alone pass CI), to insane 10 thousand line long first time PRs. In-between we also received plenty of PRs that looked fine on the surface, some of which explicitly claimed to not have made use of LLMs, but where follow-up discussions immediately made it clear that the author was sneakily consulting an LLM and regurgitating its mistake-filled replies to us.
To be clear, the point here is not to say that we believe that this is all that AI is. We don’t. This is clearly a misuse of the tool, but it is also what the overwhelming majority of LLM-based contributions looked like for our project.
So while one could in theory be a valid contributor that makes use of LLMs, from the perspective of contributor poker it’s simply irrational for us to bet on LLM users while there’s a huge pool of other contributors that don’t present this risk factor .
The people who remarked on how it’s impossible to know if a contribution comes from an LLM or not have completely missed the point of this policy and are clearly unaware of contributor poker.
For us the ability to provide contributors with an engaging ecosystem where they can improve their systems thinking and interact with other competent, trusted and prolific engineers is a critical aspect of our business model.
As I’ve mentioned before, Zig is able to punch well above its weight funding class because we put huge effort in thinking about technical, social and business management issues that surround the project.
Contributor poker is a key part of our strategy and it’s in the project’s best interest to push back against anything that hinders our ability to play the game effectively. That being said, we are aware that there are still many unresolved issues and we plan to adjust our policy as we gain more insight.
If you want to help us succeed, consider a small monthly donation to the Zig Software Foundation .
