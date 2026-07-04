---
source: "https://isaaclyman.com/blog/posts/code-review/"
hn_url: "https://news.ycombinator.com/item?id=48781256"
title: "Coding without AI: a revolutionary new way to work"
article_title: "Myth vs. Fact: why is code review so hard? • Minimum Viable Blog"
author: "encyclopedism"
captured_at: "2026-07-04T00:01:43Z"
capture_tool: "hn-digest"
hn_id: 48781256
score: 1
comments: 0
posted_at: "2026-07-03T23:34:28Z"
tags:
  - hacker-news
  - translated
---

# Coding without AI: a revolutionary new way to work

- HN: [48781256](https://news.ycombinator.com/item?id=48781256)
- Source: [isaaclyman.com](https://isaaclyman.com/blog/posts/code-review/)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T23:34:28Z

## Translation

タイトル: AI を使用しないコーディング: 革新的な新しい働き方
記事のタイトル: 神話と事実: コード レビューはなぜ難しいのか? • 実行可能な最小限のブログ
説明: 多くのチームは、コード レビューを先延ばしにしたり、電話で問い合わせたり、PR ごとに際限なくやり取りをしたりしています。その理由について話しましょう。
誤解 #1 誤解: コード レビューは退屈でプル リクエストが長すぎるため、難しいです。
事実: コードレビューはプログラムの認知作業と重複するため困難です
[切り捨てられた]

記事本文:
ソフトウェアの構築、意見、そして自分自身
通説と事実: コードレビューはなぜ難しいのでしょうか?
投稿日
2026 年 3 月 23 日
8 分で読めます
多くのチームはコード レビューを先延ばしにしたり、電話で問い合わせたり、PR ごとに際限なくやり取りをしたりしています。その理由について話しましょう。
誤解: コードレビューは退屈でプルリクエストが長すぎるため、難しいです。
事実: コードレビューは、中間ステップをすべて行わずにプログラミングの認知作業を繰り返すため、困難です。
エンジニアはコードのレビューを嫌う傾向があります。それを単に仕事の不快な部分として受け入れるのではなく、その理由を尋ねるべきです。
最もよくある誤解から始めましょう。それは PR の長さの問題ではありません。一部のチームは、短い PR の標準を確立することでレビュー時間が短縮され、精神的疲労が軽減されることを発見しました。私はこれを心から支持しますが、これらの利点をもたらすのは PR の長さだけではありません。それは、エンジニアが短い PR を提出するよう奨励されると、各 PR を整理して文書化するための作業を増やすことで補うという事実です。
理想的な形式では約 1000 行になるコード変更を行う必要があるとします。チームが PR を 200 SLOC 以下にするという基準を確立している場合、少なくとも 5 つの PR が必要になります。また、それらの PR を個々の単位として意味を持たせたい場合は、より多くのストーリーテリングを行う必要があります。
この PR は、新しい列でデータベースを更新します。これがその目的です。
この PR は DTO クラスを更新し、これらの列のデータをネットワーク経由で送信できるようにします。
この PR はサービスをリファクタリングして、新しいデータを計算に組み込みます。
小さな PR を使用すると、コード作成者が自分自身をより適切に、段階的に、より細かく説明できるようになるため、レビューが容易になります。その説明がなければ、コードレビュー担当者はすべての推論とコンテキストを再構築しようとしなければなりません。

でPRに入りました。この抽象的な推測ゲームは、実際にコードを書くのに必要な精神エネルギーよりも簡単に多くの精神的エネルギーを消費する可能性があります。私たちはコードレビューを簡単なチェックボックスのタスクだと考えていますが、多くのチームがそれに取り組む方法は、コードレビューが最も精神的に負担のかかる作業です。
コードを作成するときは、ほとんどの時間を調査に費やします。要件を明確にし、コードベースの関連部分を調査し、ドキュメントを読み、さまざまなアプローチを試し、最終的に気に入ったソリューションに到達します。完了すると、送信したコードは N 次元コンテキストの 2 次元スナップショットになります。非可逆圧縮です。スナップショットをモデルに射影することはできますが、他の手段で情報が提供されない限り、大きなギャップが生じることになります。
これらのギャップを狭めるには、いくつかの方法があります。
大きな PR を小さな PR に分割します (おそらく機能ブランチをターゲットにします)。
どのような問題を解決しているのか、どのように解決しているのか、なぜそのアプローチを選んだのかを事前に文書化して共有してください。
自身の PR 全体にコメントを残して、あまり目立たない変更に関する思考プロセスを説明してください。
ペア プログラミングまたはモブ プログラミングを使用して、構築時にコンテキストを共有します。
インタラクティブなリベース (または選択したソース管理ツールでの同等の機能) を使用してコミットを再配置し、PR を小さな単位でレビューできるようにします。
私たちはお互いにコードレビューを容易にすることができますし、そうすべきです。 PR を短くすることは、それを達成するための 1 つの方法にすぎません。
誤解: コード レビューの主な目的は、不正なコードがアプリケーションに侵入するのをブロックすることです。
事実: コードレビューの主な目的は、チームの調整を行うことです。
誤解しないでください、お互いの間違いに注意することが重要です。セキュリティの脆弱性や重大なバグは、コード レビューで発見されることがよくありますし、発見されるべきですが、早期に発見する方が良いです。詳細については、次の記事を参照してください。

ちょっと。
ただし、コード レビューの主な目的、つまり常に最大の価値をもたらす目的は調整です。コードレビューは、チームが以下を確認するためのメカニズムです。
私たちはどのような変更が加えられているかを理解しています。
私たちはこのような変更が行われる理由を理解しています。
私たちはこのコードの保守に取り組むことができます。
ここには同意については何も書かれていないことに注意してください。合意と調整は別のものです。チームメンバーは問題の解決方法について意見が異なるかもしれませんが、問題の理解を共有し、さまざまなアプローチの長所と短所を共有していれば、不完全な解決策に取り組んで、とにかく前進することができます。それが調整です。
コードレビューの第二の目的は指導です。チームメンバーはこれを使用して、テクニックを教えたり、アイデアを共有したり、アプリケーションの直接のコンテキストとより広範なテクノロジー (プログラミング言語など) の両方について互いの知識を深めることができます。コード レビューのコメントで最も優れているものの 1 つは、「この技術/機能/クラスを使用すると、コードを数行削減できると思います: [リンク]」です。コードレビューがこのようなことが起こる唯一の場所であってはなりませんし、主要な場所でさえありませんが、軽視されるべきでもありません。PR コメントから学んだことは、デフォルトで関連性があり、実践的で、すぐに実行可能なものであるため、心に残る傾向があります。
この考え方により、コード レビューが簡素化されます。レビュー担当者の最も重要な仕事は、「内容」と「理由」をすべて理解していることを確認することです。これは必ずしもコードのすべての行を読む必要がなく、高レベルで行うことができます。彼らの 2 番目の仕事は、お互いを高め、教え合うことであり、これは多くの場合、より低い抽象レベルで行われますが、それでもセミコロンやオフバイワン エラーのレベルよりも上です。彼らの最後の仕事は間違いを監視することです。これは彼らにとって最も困難ですが最も重要ではない仕事であり、彼らはそれを他の人たちと共有します。

その他数名 (品質エンジニア、プロダクトオーナー、UX デザイナー) とツール (リンター、コンパイラー、単体テストと統合テスト、脆弱性スキャナーなど)。
誤解: コード レビューはソフトウェア品質の究極の門番です。
事実: コードレビューは形式的なときに最も効果を発揮します。それ自体では、コードベースを保護するというタスクには常に不十分です。
プルリクエストが驚くべきことはほとんどありません。すぐに成功したことや舞台裏での改善をすべて発表する必要があるわけではありませんが、ほとんどの場合、チームは今後何が起こるかを事前に知っておく必要があります。誰かがコードを提出する時点では、不明な点は実装の詳細だけになっているはずで、一般的なアプローチはすでに調整されています。
しかし、チームが今後何が起こるかを知っていても、提出者がコンテキストを提供して作業を文書化したとしても、コード作成者とレビュー担当者の間には依然として本質的な知識のギャップが存在します。彼らのうちの 1 人がコードを書きました。もう一人はただそれを見ているだけです。そのギャップは大きいか小さいかもしれませんが、完全に消すことはできないと私は主張します。
そのギャップを埋める力は信頼です。信頼できない人とソフトウェアで共同作業することはできません。それは盲目的な信頼ではありません。私たちはお互いの間違いに注意し、お互いの誤解に異議を唱え、信頼だけに頼ることはありません。しかし、日、PR、またはアプリケーションによっては、それが依然として重要なものになる可能性があります。悪意のあるチームメイトがプロジェクトの過程で問題のあるコードを忍び込む機会は無数にありますが、コード レビュー プロセスはそれを完全に防ぐことはできませんし、そうすべきではありません。あたかも作成者が妨害者であるかのようにあらゆるコードレビューに取り組むことは、敵対的であることは言うまでもなく、非常に非効率的です。
どの程度の信頼を満足できるかはコードレビューア次第ですが、前述したように、

さらに、彼らのレビューは、いくつかの（人的および自動化された）安全手段の 1 つである必要があります。そして、理想的な状況であっても、「信頼ギャップ」を最適化することはできません。 it is the optimization.結局のところ、コード作成者は誠実に行動する必要があります。しっかりとした高品質のコードを書くことに専念する必要があり、そうしないとコードベースの標準が時間の経過とともに変動することは避けられません。世界最高のコードレビューアでも、世界最低のプログラマーには敵いません。
では、コードレビューが究極の門番ではないとしたら、何でしょうか?それは、解決しようとしている問題によって異なります。
あまりにも多くのバグが本番環境に導入されている場合、正しい対応は QA にさらに投資することです。機能が繰り返し仕様を満たさない場合、製品所有者はエンジニアリングとの連携を強化する必要があります。セキュリティの脆弱性が頻繁に見落とされる場合、チームはさらに脅威のモデリングを行うか、侵入テスターを雇う必要があります。
Code review is valuable.誰もそうではないとは言っていません。しかし、それは大きな機械の小さな歯車に過ぎません。
通説: コードレビューは正しく行えば多くの問題を解決します。
事実: コードレビューはあらゆる有意義な結果を補足するものです。それ自体では問題を解決できるほど強力なツールではありません。
この神話は、前の神話と多くの点で同じ内容をカバーしていますが、強調する価値はあります。
この頭の体操を試してみてください。コード レビューが違法になったら、自分はどのように仕事を変えるだろうかと自分自身 (またはチーム) に問いかけてください。 Maybe you would:
変更をより慎重かつ協力的に計画します。
定期的にリファクタリングの時間を確保するなど、技術的負債にもっと対処してください。
開発プロセス中、QA および製品とより緊密に連携します。
チームの技術基準をより明確に文書化します。
ジュニアチームメンバーの指導とトレーニングにもっと積極的に取り組んでください。
さらに自動テストを作成して統合する

より優れた静的解析ツール。
思いついたことは何でも、それを実行する必要があります。フルストップ。コードレビューよりも良い結果が得られます。
コードレビューの約束はその多用途性です。優れたコードレビューアは、膨大な数のカテゴリの間違いを見つけることができます。しかし、それは非効率でもあります。これらのカテゴリのほとんどについては、より効果的でそれほど疲れない監視方法があります。その方法を見つけるのは技術的なリーダーの仕事であり、その仕事が終わることはありません。
コードレビューは確かに安全策ですが、他の何百もの安全策の失敗モードでもあります。注意を払えば、プロセスのどこが不足しているかを知ることができます。
コードレビューはしばしば骨の折れる仕事ですが、全員が自分の役割を果たせば、その必要はありません。コード作成者が十分なコンテキストとドキュメントを提供し、レビュー担当者が優先順位を理解し、チームが効果的なガードレールに投資すれば、その必要はありません。チームのプロセスが改善されるにつれて、コードレビューはより速く、より簡単になります。
カテゴリ: 開発
タグ: ソフトウェア構築
前へ
前の投稿: AI を使用しないコーディング: 革新的な新しい働き方
次の投稿: 生成 AI の 5 つの法則 Next
© 2026 アイザック・ライマン著。このブログには、AI によって生成されたコンテンツは一切掲載されていません。

## Original Extract

A lot of teams procrastinate code review, phone it in, or get stuck in endless back-and-forth on every PR. Let’s talk about why.
Myth #1 Myth: Code review is difficult because it’s boring and pull requests are too long.
Fact: Code review is difficult because it duplicates the cognitive work of progr
[truncated]

Software construction, opinion, and self
Myth vs. Fact: why is code review so hard?
Posted on
Mar 23, 2026
8 mins read
A lot of teams procrastinate code review, phone it in, or get stuck in endless back-and-forth on every PR. Let’s talk about why.
Myth: Code review is difficult because it’s boring and pull requests are too long.
Fact: Code review is difficult because it duplicates the cognitive work of programming without all the intermediate steps.
Engineers tend to dislike reviewing code. Instead of just accepting it as an unpleasant part of the job, we should ask why.
Let’s start with the most common misconception: it’s not about PR length. Some teams have discovered that establishing a norm of smaller PRs leads to faster review times and less mental fatigue, and I wholeheartedly support this, but it’s not the PR length alone that brings these benefits. It’s the fact that when engineers are encouraged to submit shorter PRs, they compensate by doing more work to organize and document each PR.
Suppose you need to make a code change that, in its ideal form, comes in around 1000 lines. If your team has established a norm of PRs being 200 SLOC or less, you’ll need at least five PRs, and if you want those PRs to make sense as individual units, you’ll have to do a lot more storytelling:
This PR updates the database with new columns; here’s what they’re for.
This PR updates DTO classes so the data in those columns can be sent over the wire.
This PR refactors a service to incorporate the new data into its calculations.
Small PRs are easier to review because they make code authors explain themselves better—step by step, at a finer grain. Without that explanation, the code reviewer has to try to to reconstruct all the reasoning and context that went into the PR. This abstract guessing game can easily take more mental energy than it took to actually write the code. We think of code review as a quick checkbox task, but the way many teams approach it, it’s the most mentally taxing thing they do.
When you code, you spend most of your time researching. You clarify requirements, explore the relevant parts of the codebase, read documentation, try out different approaches, and finally land on a solution you like. When you’re finished, the code you submit is a two-dimensional snapshot of N-dimensional context. It’s lossy compression. The snapshot can be projected back into a model, but unless information is provided by other means, there will be significant gaps.
There are several ways to narrow those gaps:
Split up large PRs into smaller ones (perhaps targeting a feature branch).
Document and share ahead of time what problem you’re solving, how you’re solving it, and why you chose that approach.
Leave comments throughout your own PRs, explaining your thought process around less obvious changes.
Use pair programming or mob programming to share context as you build.
Use interactive rebasing (or the equivalent in your source control tool of choice) to rearrange commits and make PRs reviewable in small chunks.
We can make code review easier for each other, and we should. Shorter PRs are just one way to accomplish that.
Myth: The primary purpose of code review is to block bad code from getting into the application.
Fact: The primary purpose of code review is to align the team.
Don’t get me wrong, it’s important to look out for each other’s mistakes. Security vulnerabilities and critical bugs are often caught in code review, and they should be, though it’s better to catch them earlier—more on that in a moment.
Code review’s primary purpose, though, the purpose that consistently brings the most value, is alignment. Code review is a mechanism by which a team affirms:
We understand what changes are being made.
We understand why these changes are being made.
We can commit to maintaining this code.
Note that there’s nothing here about agreement. Agreement and alignment are different things. Team members may disagree about how to solve a problem, but if they share an understanding of the problem, along with the pros and cons of different approaches, they can commit to an imperfect solution and move forward anyway. That’s alignment.
A secondary purpose of code review is mentorship. Team members can use it to teach techniques, share ideas, and deepen each other’s knowledge of both the immediate context of the application and broader technologies (e.g. programming languages). One of the best kinds of code review comments is, “I think we could shave off several lines of code by using this technique/feature/class: [link].” Code review shouldn’t be the only place this happens, or even the main place, but it shouldn’t be discounted either—the things you learn from PR comments tend to stick with you, since by default they’re relevant, hands-on, and immediately actionable.
This mindset simplifies code review. The reviewer’s most important job is to make sure they understand all the whats and whys , which can be done at a high level, without necessarily reading every line of code. Their second job is to lift and teach each other, which often happens at a lower level of abstraction, but still above the level of semicolons and off-by-one errors. Their final job is to watch for mistakes: this is their most difficult but least important job, and they share it with several other people (quality engineers, product owners, UX designers) and tools (linters, compilers, unit and integration tests, vulnerability scanners, etc.).
Myth: Code review is the ultimate gatekeeper of software quality.
Fact: Code review is at its best when it’s a formality. On its own, it will always be insufficient to the task of protecting a codebase.
Pull requests should rarely be a surprise. Not every quick win or behind-the-scenes improvement needs to be announced, but most of the time, the team should know well in advance what’s coming. By the time someone’s submitting code, the only remaining unknowns should be implementation details, the general approach having already been aligned on.
But even when the team knows what’s coming, and even when the submitter provides context and documents their work, there’s still an essential knowledge gap between the code author and the reviewer. One of them wrote the code; the other is just looking at it. That gap may be larger or smaller, but I would argue it can’t be erased completely.
The force bridging that gap is trust . You can’t collaborate on software with people you don’t trust. It’s not blind trust—we watch out for each other’s mistakes, we challenge each other’s misconceptions, we never rely on trust alone—but depending on the day, the PR, or the application, it can still be substantial. A malicious teammate will have myriad opportunities to sneak in problematic code over the course of a project, and no code review process can completely prevent that, nor should it try to. Approaching every code review as if the author might be a saboteur is incredibly inefficient, not to mention hostile.
It’s up to the code reviewer how much trust they’re comfortable with, but as I mentioned earlier, their review should be one of several (human and automated) safeguards. And even in ideal circumstances, the “trust gap” can’t be optimized out; it is the optimization. At the end of the day, the code author has to be acting in good faith. They have to be committed to writing solid, high-quality code, or the codebase’s standards will inevitably drift over time. The world’s best code reviewer is no match for the world’s worst programmer.
So if code review isn’t the ultimate gatekeeper, what is? That depends on the problem you’re trying to solve.
If too many bugs are making it to production, the right response is to invest more in QA. If features repeatedly don’t satisfy specifications, the product owner needs to tighten their collaboration with engineering. If security vulnerabilities are frequently overlooked, the team needs to do more threat modeling or hire a pentester.
Code review is valuable. No one’s saying it isn’t. But it’s a small cog in a big machine.
Myth: Code review solves a lot of problems when you do it right.
Fact: Code review is supplemental to every meaningful outcome. It’s not a powerful enough tool to solve problems on its own.
This myth covers a lot of the same ground as the previous one, but it’s worth emphasizing.
Try this mental exercise: ask yourself (or your team) how you would work differently if code review were outlawed. Maybe you would:
Plan your changes more carefully and collaboratively.
Do more to address technical debt, like regularly setting aside time to refactor.
Collaborate more closely with QA and Product during the development process.
Document the team’s technical standards more clearly.
Be more proactive about teaching and training junior team members.
Write more automated tests and integrate better static analysis tools.
Whatever you came up with, you should do it. Full stop. You’ll get better results than you get from code review.
The promise of code review is its versatility. A good code reviewer can catch mistakes in a huge number of categories. But it’s also inefficient: for most of those categories, there’s a more effective, less exhausting way to keep watch. Finding that way is the work of technical leadership, and the work is never done.
Code review is a safeguard, yes, but it’s also the failure mode of a hundred other safeguards. If we pay attention, it can tell us where our processes fall short.
Code review is often exhausting, but it doesn’t have to be if everyone does their part: if code authors provide ample context and documentation, reviewers know their priorities, and the team invests in effective guardrails. As the team’s processes improve, code review gets faster and easier.
Categories: development
Tags: software construction
Previous
Previous post: Coding without AI: a revolutionary new way to work
Next post: Five laws of generative AI Next
© 2026 by Isaac Lyman. No AI-generated content has ever appeared on this blog.
