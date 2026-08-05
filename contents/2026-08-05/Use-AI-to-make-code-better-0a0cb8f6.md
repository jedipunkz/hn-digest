---
source: "https://joshmock.com/post/2026-08-use-ai-to-make-code-better/"
hn_url: "https://news.ycombinator.com/item?id=49183895"
title: "Use AI to make code better"
article_title: "Use AI to make code better // Josh Mock"
author: "JoshMock"
captured_at: "2026-08-05T15:08:35Z"
capture_tool: "hn-digest"
hn_id: 49183895
score: 1
comments: 0
posted_at: "2026-08-05T14:56:42Z"
tags:
  - hacker-news
  - translated
---

# Use AI to make code better

- HN: [49183895](https://news.ycombinator.com/item?id=49183895)
- Source: [joshmock.com](https://joshmock.com/post/2026-08-use-ai-to-make-code-better/)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T14:56:42Z

## Translation

タイトル: AI を使用してコードを改善する
記事のタイトル: AI を使用してコードを改善する // Josh Mock
説明: さらに多くのコードを書くことは間違った目標です。

記事本文:
コードを書いて生計を立てており、自分が上司ではない場合は、望むと望まざるにかかわらず、ワークフローのどこかで AI を利用したコーディング ツールを使用している可能性が高くなります。また、組織内の誰かが、反対の証拠がたくさんあるにもかかわらず、「ボットがコードを書く」ということは「プログラマーがより多くの仕事をより速くこなす」ことを意味すると考えている可能性も十分にあります。
確かにエージェントはほとんどの開発者よりも速くコードを書くことができますが、ほとんどの場合、それはちょっとひどいコードです。過去 25 年間、くだらないコードをたくさん書いてきた人間として、私にはその判断を下す資格があると感じています。くだらないコードは開発者の速度を低下させ、ソフトウェアの速度も低下させます。見つけるのが難しいバグや、解明するには多大な労力を必要とするパフォーマンスの低下が発生します。私たちのほとんどは、それが良いアイデアであることはめったにないことを知っているにもかかわらず、クソコードは多くの全面的な書き換えを引き起こしました。
クソコードが配信スケジュールやパフォーマンスを遅らせるのであれば、なぜ私たちは自信を持ってクソコードを書くエージェントをあらゆることに投げ込むのでしょうか。彼らは、人間がすでに引き起こしているのと同じ問題を、わずか 20 倍の速度で引き起こします。ここ数年、私たちは AI が間違った問題に対して間違った方法で適用されたために、ソフトウェアとサービス (および政治的演説やブログ投稿など) の品質が着実に低下しているのを目の当たりにしてきました。すでに作成した問題にそれらを投げて、興味深い作業は自分たちで保存する必要があります。スピードよりも品質を最適化する必要があります。それが長期的にはさらに前進することになるからです。スピードより持久力。カメ対ウサギ。 「遅いことはスムーズであり、スムーズなことは速いことです。」などなど
技術的負債は手抜きの副作用です。変更をより速く提供するために、コードを最適な状態に保つために時間がかかりすぎるため、理想的ではないトレードオフが行われます。アイデアは、「

プロジェクトを長期的に健全に保つためには、最終的には借金を返済しなければならないことを承知しており、このことをより早く完了させるために、未来の自分たちから「融資」を受けています。
技術的負債は通常、適切な変更が面倒であるか、検索/置換やコード修正するには少し複雑すぎるか、すぐに解決するには多すぎるファイルやプロジェクトにまたがる場合、またはリグレッションを避けるために欠落しているテスト カバレッジを追加する必要がある場合に発生します。退屈な作業をしても疲れ果てず、依存関係を簡単に追跡できるツールを使用でき、単体テストを書くのが得意な脳は誰だと思いますか?エージェントの皆さん！
そのため、新たな技術的負債を導入する代わりに、未解決の技術的負債を残したまま、関心のある問題に取り組むことができるようになりました。次に、エージェントを退屈な作業に投入し、債務が存在しなくなるまで、または合意された停止点に達するまで、債務を反復処理させることができます。
既存のテクノロジー負債もそれほど変わりません。どのような負債がエージェントによって解決できるかを特定し、それらを解放してください。上記の負債のリストがまだ手元にない場合は、上級エンジニアを集めて 1 時間かけてリストを作成してください。私は以前にこの演習を行ったことがあります。それはワイルドで謙虚な経験です。コーディング エージェントはおそらくそれらすべてを削除することはできませんが、面倒なものや、解明するのに面倒な量のヤクの削り作業が必要で、自動テストとコード レビューによって検証できるものであれば、解決できる可能性があります。
テストカバレッジおよびその他の決定的なガードレール
私はコード変更に常に単体テストを含めてきましたが、100% のテスト カバレッジを目指すのは時間の有効活用ではないと常々感じていました。 100% が必要になることはめったにないと今でも思っていますが、LLM で生成されたコードがプロジェクトに追加される場合には、優れた単体テストのような決定論的なガードレールがこれまで以上に不可欠です。
Red-Green TDD は、エージェントが適切な作業と簡潔なソリューションに集中し続けるための優れた方法です。素敵な

この方法で作業することの副作用として、最新の変更には常にテスト カバレッジが含まれるため、時間の経過とともにカバレッジが向上します。また、新しいテストを追加するとき、またはいくつかの良いカルマのために一連の変更の最後にテストを追加するとき、私が取り組んでいるコード領域のテストを改善する簡単な機会を特定するようエージェントに促す習慣もついています。
他にも多くの決定論的なガードレールが存在しており、同様に投資する価値があります。
CodeQL などによる静的分析
これらをまだ使用していない場合は、それらを追加し、それらによって明らかになった問題の多くをすぐに解決する作業が、これまでになく簡単になりました。ガードレールを使用すると、コードの品質が永続的に向上します。
これらのファジングや静的解析の結果は、初めて実行したときに身の毛もよだつような結果になるかもしれません。幸いなことに、彼らが提起する問題の多くは修正するのが難しくなく、多くの場合、複雑というよりも退屈です。エージェントはこの点で優れています!
私はかつて、サードパーティのレッド チームが Web アプリに文字通り数百もの XSS 脆弱性を発見したプロジェクトに取り組んでいました。使用されたテンプレート ライブラリの 1 つは、デフォルトでは文字列を HTML エンコードしませんでした。エージェントは、必要なテスト カバレッジを追加し、私たちが発見したすべてを修正するのにかかる時間の数分の一で、それらをすべてクリーンアップすることができたでしょう。
私は、明確な実装計画のないリポジトリにエージェントを任せるのは好きではありませんが、Andrej Karpathy の自動調査ツール (または、私の場合はそれにインスピレーションを得た Pi プラグイン) は例外です。コード内で改善したい測定可能な指標があるものの、そこに到達するための明確な指示がまだない場合です。
最近、仕事で作成していた CLI ツールで自動リサーチを一晩実行しました。その結果、事前に選択したコマンド セットの実行速度が 1 桁向上しました。一日中掃除をしなければならなかった

結果を改善し、いくつかの醜い微細な最適化を取り除き、変更を一貫した PR に変えましたが、同じ作業を手動で実行すると、はるかに時間がかかるでしょう。
「うまくいっていれば何をしても構わない」
ソフトウェア プロジェクトのすべての部分が他の部分と同じように重要であるわけではありません。何かが機能するだけで十分な場合がありますが、それが機能する限り、それがどのように機能するかを知る必要はありません。
最近、リリースの行き詰まりに陥って、パッケージの特定のバージョンを公開するには release-please が必要ですが、バージョンを間違った番号に自動的にバンプする PR が開き続けました。 release-please がどのように機能するかは気にしませんし、1 時間もドキュメントをざっと読みたくありませんでした。エージェントは 5 分かけて、git と GitHub CLI コマンドを使用して問題を解決する使い捨て Bash スクリプトを作成しました。明らかな問題がないかざっと確認し、実行して機能することを確認し、削除しました。 10 分と 2 時間。
別の例では、Jujuku を使用して、ブランチ上の大量のコミットを並べ替えて、より一貫性のある読みやすい履歴にまとめたいと考えていました。繰り返しますが、私は使い捨ての Bash スクリプトを要求し、それを確認し、実行して、削除しました。
Pi エージェント ハーネスの作成者である Mario Zechner 氏も、PragProg ポッドキャストで、Pi のさまざまな部分がどのように構築されたかについて同様のことを述べています。
Pi には手抜きがたくさんありますが、重要なコードであるとわかっている部分では、手抜きを避けるようにしています。たとえば、現在のセッションを取得して、GitHub などでホストできる HTML ファイルを吐き出す HTML エクスポート機能があります。私はその関数のコードを一行も見ていません。壊れていても、外したときに見た目が良ければ気にしません。ただし、エージェント ループ自体や拡張機能の読み込みメカニズムなども必要です。そしてそれは重要です。
冗談です。 AI の落書きを自分の責任で公開するのはやめましょう

名前！人間性の本質を欠いた言葉を自分の手柄にしないでください。スロップを公開する必要がある場合は、あなたに直接関連付けられないように、ペンネームまたはチーム名を使用してください。
ここでコードについても学ぶべき教訓があります。書かれた言葉には人間の本質が込められていますが、コードの場合は少し異なります。コードはそれを実行するコンピューターを対象としているため、書き言葉のような表現力を持たない厳密性があります。これにより、AI にシェアを書き込むことを正当化することが少し簡単になります。
しかし、構造と構成、読みやすさと保守のしやすさ、システム全体のアーキテクチャ、ユーザー インターフェイスなど、さまざまな点で私たちの人間性が依然としてコードの中に残っています。結局のところ、私たちが作成するソフトウェアの向こう側のどこかに人間がいます。私たちの監視下で生成を許可するコードについては、このことを念頭に置く必要があります。
性交の速度を落とすことについての考え
ジョン・レーガー「実行可能な神託」について語る
コード分析用の形式的推論エンジンを LLM に提供する
複数のことが同時に真実になることがある
メンテナンスコストを削減する AI が必要です
すべての単語 (エムダッシュを含む) は LLM を使用せずに私によって書かれました。
このページは、Cookie などを介して個人を特定できる情報を収集しないことにより、お客様のプライバシーを尊重しています。
特に明記されていない限り、ここにあるすべての資料は Josh Mock によって書かれており、次のライセンスが付与されています。

## Original Extract

Writing more code is the wrong goal.

If you write code for a living and you aren't your own boss, there's a solid chance by now that you are using AI-powered coding tools somewhere in your workflow, whether you want to or not. There's also a decent chance that someone in your org thinks that "bots write code" means "coders get more done, faster" despite plenty of evidence to the contrary.
Agents can indeed write code faster than most devs, but most of the time it's kinda shitty code. As a person who has written lots of shitty code over the last 25 years, I feel pretty qualified to make that judgment. Shitty code slows developers down, and slows software down too. It introduces bugs that are hard to find and performance regressions that take serious effort to unravel. Shitty code has inspired a lot of total rewrites, despite most of us knowing it's rarely a good idea .
If shitty code slows down delivery schedules or performance, why are we confidently throwing agents that write shitty code at everything. They create the same problems humans already do, just at 20x speed. Over the last couple years, we've witnessed the quality of software and services (and political speeches, and blog posts, and...) in steady decline thanks to AI being applied in the wrong ways on the wrong problems. We should be throwing them at problems we've already created and save the interesting work for ourselves. We should be optimizing for quality over speed, because that ends up getting us further in the long run. Endurance over speed. Tortoise vs. hare. "Slow is smooth and smooth is fast." etc. etc. etc.
Technical debt is side effect of cutting corners: in order to deliver a change faster, some non-ideal tradeoff is made to keep code in an optimal state because it would take too long. The idea is that we "take out a loan" from our future selves to get this thing done sooner, knowing that we will have to eventually pay down the debt to keep the project healthy in the long term.
Tech debt usually arises because the correct change is tedious, or a bit too complex for a search/replace or codemod, or crosses too many files or projects to resolve quickly, or involves adding missing test coverage to avoid regressions. Guess whose brains don't burn out when given tedium, can use tools that easily keep track of dependencies, and are great at writing unit tests? Agents!
So now, instead of introducing new tech debt, we can work on the problem we care about, leaving bits of unresolved tech debt in our wake. Then we can throw agent at the tedium and have it iterate on the debt until it doesn't exist any more, or reaches some agreed-upon stopping point.
Existing tech debt isn't much different: identify what debts can be resolved by an agent, and let them loose. If you don't have a list of said debts handy already, gather your senior engineers for an hour and catalog them. I've done this exercise before; it's a wild and humbling experience. Coding agents probably won't be able to remove all of them but anything tedious, or that takes some annoying amount of yak-shaving to unravel, and can be verified by automated testing and a code review, may be solvable.
Test coverage, and other deterministic guardrails
I've always included unit tests in my code changes but always felt that aiming for 100% test coverage isn't the best use of time. I still think 100% is rarely necessary, but deterministic guardrails like good unit tests are more essential than ever when LLM-generated code is added to a project.
Red-green TDD is a great way to keep an agent focused on the right work and a concise solution. A nice side effect of working this way is that your latest change always has test coverage, and thus coverage improves over time. I've also gotten into the habit of prompting an agent to identify any simple opportunities to improve tests on the area of code I'm working on, either while adding the new test, or tacking it on at the end of a set of changes just for some good karma.
Many other deterministic guardrails exist and are equally worth the investment:
Static analysis, via CodeQL or similar
If you aren't using these yet, the effort to add them, and then immediately resolve many of the issues they uncover, has never been easier. The quality of your code will permanently improve if the guardrails are used.
All those fuzzing and static analysis results might scare the hair off your head the first time you run them. Fortunately, many of the problems they raise are not hard to fix, and are often more tedious than complex. Agents are great at this!
I once worked on a project where a third-party red team found literally hundreds of XSS vulns in our web app. One of the templating libraries used did not HTML-encode strings by default. An agent could have added needed test coverage then cleaned them all up in a fraction of the time it took us to fix everything they found.
I don't prefer letting an agent loose on a repo with no clear implementation plan, but Andrej Karpathy's autoresearch tool —or, in my case the Pi plugin inspired by it —is an exception, when you have a measurable metric you want to improve in your code, but no clear direction on how to get there yet.
I recently ran autoresearch overnight on a CLI tool I've been building at work, and its results improved the runtime speed of a preselected set of commands by an order of magnitude . I had to spend a day cleaning up the results, stripping out some ugly micro-optimizations and turning the changes into a coherent PR, but the same effort done manually would have taken far longer.
"I don't care what you do as long as it works"
Not every part of a software project is as critical as every other. Sometimes you just need something to work, and you don't need to know how it works as long as it works.
I recently got stuck in release limbo, needing release-please to publish an specific version of a package, but it kept opening a PR that auto-bumped the version to the wrong number. I don't care how release-please works, and I didn't want to skim the docs for an hour. An agent took five minutes to write a throwaway Bash script that used git and GitHub CLI commands to put things right. I skimmed it for obvious problems, ran it, verified it worked, and deleted it. 10 minutes vs. 2 hours.
On another instance, I wanted to use Jujutsu to reorder and squash a bunch of commits on a branch into a more coherent, readable history. Again, I asked for a throwaway Bash script, reviewed it, ran it, deleted it.
Mario Zechner, creator of the Pi agent harness, said something similar on the PragProg podcast about how different parts of Pi were built:
There's a lot of slop in Pi, but I try to avoid it in the bits and pieces where I know that it's important code. Like, we have an HTML export functionality where it takes the current session and just spits out an HTML file that you can then host on GitHub and whatever. I have not looked at a single line of code for that function. I don't care if it's broken, if it looks right when it comes out. But then there's the agent loop itself or the extension loading mechanism and all of that stuff. And that's important.
Just kidding. Stop publishing AI slop under your own name! Don't take credit for words that lack the very essence of humanity. If you must publish slop, use a pseudonym or the name of your team so it isn't associated with you directly.
There's a lesson to be learned here with code, too: the written word carries our human essence, but it's a little different with code. Code targets the computer running it, so there's a strictness to it that isn't expressive in the way written languages are. It makes it a bit easier to justify letting AI write its share.
However , our humanity is still in the code in so many ways: the structure and organization, its readability and maintainability, the overall architecture of the system, the user interface... At the end of the day, there's a human somewhere at the other end of the software we write. We have to keep this in mind with whatever code we allow to be generated under our watch.
Thoughts on slowing the fuck down
John Regehr on "executable oracles"
Giving LLMs a Formal Reasoning Engine for Code Analysis
Multiple things can be true at the same time
You need AI that reduces maintenance costs
All words—including emdashes—written by me without the use of LLMs.
This page respects your privacy by not collecting any personally identifiable information, via cookies or otherwise.
All material here is written by Josh Mock and has the following license, unless otherwise noted:
