---
source: "https://www.builder.io/blog/developers-drowning-in-ai-prs"
hn_url: "https://news.ycombinator.com/item?id=48408569"
title: "I Didn't Become a Developer to Review AI Slop"
article_title: "I Didn't Become a Developer to Review AI Slop"
author: "ankitg12"
captured_at: "2026-06-05T07:38:00Z"
capture_tool: "hn-digest"
hn_id: 48408569
score: 4
comments: 0
posted_at: "2026-06-05T06:03:49Z"
tags:
  - hacker-news
  - translated
---

# I Didn't Become a Developer to Review AI Slop

- HN: [48408569](https://news.ycombinator.com/item?id=48408569)
- Source: [www.builder.io](https://www.builder.io/blog/developers-drowning-in-ai-prs)
- Score: 4
- Comments: 0
- Posted: 2026-06-05T06:03:49Z

## Translation

タイトル: AI のスロップをレビューするために開発者になったわけではありません
説明: AI により PR は安価になりますが、信頼は依然として高価です。開発者が終わりのない AI のスロップのレビューに行き詰っている理由と、自動化されたコード レビューがボトルネックをどのように解決できるか。

記事本文:
AI のスロップをレビューするために開発者になったわけではありません メイン コンテンツにスキップ Frete がフロントエンドのビルド時間を 70% 削減した方法をご覧ください
最適な AI ツールとは何ですか? AI の現状に関する調査に参加する
本番環境に対応した Web アプリと UI を生成する
ページとヘッドレスコンテンツを生成、反復、最適化する
AI と設計システムを使用してコーディングを高速化する
コードで設計し、自信を持って引き継ぎます
ソフトウェア開発ライフサイクル全体を加速するプロトタイプを出荷
開発者に依存せずにエクスペリエンスを生成して公開する
最新の洞察、チュートリアル、お知らせ
ウェビナー、ケーススタディ、ガイドのライブラリ
今後のイベントと録画されたセッション
Builder を使用した実際の成功事例
ディスカッションに参加したり、質問したりする
Builder パートナーを調べてチームとつながる
スタックと統合し、ツールを接続します
最新の機能と改善点
新規または既存のアプリのコードに雰囲気を与える方法を学ぶ
Builder を使用してサイトのコンテンツを公開する
Figma デザインを実際のコードに変換する
Frete がどのようにフロントエンドのビルド時間を 70% 削減したかをご覧ください
最適な AI ツールとは何ですか? AI の現状に関する調査に参加する
‹ ブログに戻る AI 私は AI の失敗をレビューするために開発者になったわけではありません
でも最近は、まさにそれが仕事だと感じています。
私の PR キューは、技術的にコンパイルできる作業でいっぱいです。この要約はもっともらしく聞こえます。いくつかのテストもあるかもしれません。次に、diff を開いたときに実際の作業が始まります。
変更によって何が行われることになっていたのでしょうか?実際にフローを実行した人はいますか?このヘルパーが 6 回複製されているのはなぜですか?これは実際にバグを修正しているのでしょうか、それとも AI がただグルグル回って完了したと判断しただけなのでしょうか?
AI のおかげで、私のチーム (そしてあなたのチーム) の誰もがコードを簡単に作成できるようになりましたが、そのコードが信頼できるものになったわけではありません。
Stack Overflow の 2025 年の開発者調査では、AI ツールに対する最も一般的な不満は、出力が「ほぼ正しいが完全ではない」ことであることがわかりました。

Sonar の 2026 State of Code レポートによると、開発者の 96% が AI によって生成されたコードを完全には信頼しておらず、38% が AI 生成コードのレビューには人間が書いたコードのレビューよりも多くの労力がかかると述べています。
それは、AI コードは問題ないように見えますが、何がうまく機能しているかを確認するには実際に深く掘り下げる必要があるからです。はっきり言って、悪いコードは拒否するのがはるかに簡単です。
イライラしてる。もしかしたらあなたもそうかもしれません。これを掘り下げて一緒に解決しましょう。
PR は...今ではあまりにも安すぎる
AI エージェントは、バグが本物であることに誰かが同意する前に、Jira チケットからブランチを作成したり、Slack スレッドからパッチを作成したり、バグ レポートから完全な PR を作成したりすることができます。正直、とても素晴らしい世界です。
しかし重要なのは、これらのツールを使用しているのは開発者だけではないということです。 PM は、3 つのスプリントにわたって説明しようとしてきた機能のプロトタイプを作成しますが、そのほとんどは曖昧で役に立たない手を振るだけです。デザイナーは UX フローを微調整し、優先順位が下がり続けるレイアウトを修正します。マーケティング担当者はランディング ページとフォームを更新します。 (継続的に。) サポートは、彼らが最もよく知っている顧客の問題点にパッチを当てます。
そしてそれはすべて勝利です。小さな修正は、コードのその部分をたまたま知っているエンジニアを待ってバックログ地獄に放置すべきではありません。製品知識は、より早く実用的なソフトウェアに変わるはずです。
しかし、PR を開くのが簡単になるにつれて、開発者は PR をレビューする義務が増えます。また、PR は存在するだけで価値があるわけではありません。信頼できる場合にのみ価値があります。
そして信頼は依然として非常に高価です
AI はコードを書くのが非常に得意です。最近のハッカソンでは、GPT 5.5 で 10,000 行の動作コードを約 45 分でスピンアップさせました。アプリはほとんど動作しました。確かに、UI は悪夢でしたが、核となる機能は存在していました。
しかし、コードを書くことと、信頼できるスケーラブルなコードを書くことは別のことです。モデルは diff を生成し、それを説明し、さらにいくつかのプログラムを実行することもできます。

ハッピーパステスト。しかし、実際に重要なことについては、依然として誰かが責任を負っています。
このコードは実際に記載された問題を解決しましたか?
作者はこのシステムを本当に理解していましたか? それとも、これが後々の技術的負債を生み出しているのでしょうか?
差分が必要以上に大きくなっていませんか？ (ほぼ間違いなく。)
この修正により、コード内の他のフローが静かに中断されますが、これは 1 人のユーザーが試しただけであれば明らかです。
UI は実際のブラウザーで実際のユーザーに対して実際に機能しますか?
この修正はデモ後も存続しますか?
これは根本的な問題の解決策なのでしょうか、それとも単なる絆創膏なのでしょうか?
このセキュリティのトレードオフは許容できますか?
これらは構文に関する質問ではありません。それらは信頼に関する質問です。そして今、それらはすべて開発者のあなたと私に降りかかります。 @richiemcilroy は、先日バイラルにツイートされたビデオでそれをうまく表現しました。
数字も同じことを物語っています。 LinearB の 2026 年のベンチマークでは、AI PR は人間が作成した PR に比べて 4.6 倍長く審査を待ち、拒否されることがはるかに多いことがわかりました。経験豊富なオープンソース開発者を対象とした METR の調査では、2025 年初頭の AI ツールにより実際に開発者の速度が 19% 遅くなっていることが判明しました。その理由の 1 つは、実際の作業には入力だけではなく、スタイル、テスト、ドキュメント、レビューが含まれるためです。
AIが役に立たないと言っているわけではありません。そして、ツールは日々改良され続けています。しかし、ソフトウェアの実際の仕事は、単にコードをファイルに入力することではありません。何を変更する必要があり、何を変更すべきではないか、また、問題を技術的に修正する表面レベルのパッチが実際に今後 6 か月間チームに悩まされる時期を知ることです。
そこに注意を向ける必要があります。 PR がすでにキューに入った後で基本的なものを手動で再発見するのではなく、センスとコンテキストが必要なものを比較検討する必要があります。
今開発中は気持ち悪いです
AI ツールによってすべての人の生産性が向上しているとはいえ、ボトルネックになるのは恐ろしいことです。他の全員が達成することができます

突然コードがオープンになるため、これまで以上に多くのことを行っています。
開発者であるあなたは、レビューの負債として誇大宣伝を経験するだけです。あなたは構築していません。レビューしてるんですね。あなたはシステムを設計しているのではありません。あなたはその端を守っているのです。難しい問題を直接解決しているわけではありません。エージェントやチームメイトがやろうとしていたことをリバースエンジニアリングして、その差分を保存しておいても安全かどうかに午後を賭けているのです。
AI が楽しい部分を担当します。あなたはロボットになれるのです。
それはあなたが役に立たないという意味ではありません。むしろ、今はあなたの判断がはるかに重要です。
しかし、ワークフローはあなたの判断力をひどく消耗させます。システム内で最も不足しているリソース、つまり経験豊富なエンジニアリングの注意力を使い、謎の差分、肥大化したパッチ、欠落しているコンテキスト、正しく見えるだけの生成されたコードをターゲットにしています。
だから、うん、退屈だよ。はい、イライラします。誰かが「今では誰もがコードを配布できるようになりました」と言うと、あなたや私は「今では誰もが私たちのために作品を作成できるようになりました」と聞きます。
リポジトリをロックダウンすると間違った問題が解決される
それで、どうすればいいでしょうか？そうですね、当然の反応は、リポジトリをロックすることでしょう。開発者のみ。
それはわかります。午前 2 時に製品がダウンしたときにページングを受けるのはあなたです。コードを守ることはエリート主義ではありません。あなたには記憶があるだけです。
しかし、アクセスを制限すると間違った問題が解決されます。
部門を超えた PR が必ずしも悪いわけではありません。実際、多くの点で、これらはまさに私たちが長年望んでいたものです。エンジニアのカレンダーを待つことなく、製品知識が小さな修正に変わります。
しかし問題は、誰もが PR を開くことができるようになったにもかかわらず、PR の摂取自体が進化していないことです。チームは依然として PR を開発者間の引き継ぎのように扱います。これが差分で、ここに説明があります。頑張ってください。著者が同じローカルコンテキスト、同じテスト習慣を持つ別のエンジニアだった場合、これはうまく機能しました。

そして、査読者が何を必要としているのかについても同じ直感を持っていました。
しかし、その想定は今や崩れ去ります。非開発者が不注意だからではありません。実際、デザイナー、PM、マーケティング担当者、サポート チームは全員、問題に最も近い立場にあるため、最適なユーザー コンテキストを持っています。しかし、彼らはおそらく、開発者としてリスクを評価するために何が必要なのかを知りません。そして、AI が実際の実装を生成したとき、PR を開いた人でさえ、何が変更されたのか全容を把握できない可能性があります。
ミステリーの差分は、共同作業を行うための合理的な方法ではありません。では、PR との仕事のやり方をどのように変えるのでしょうか?
PR の証拠の基準を高める
開発者は、生成された PR または機能横断的な PR を開いて、最初からリバース エンジニアリングする必要はありません。すべての PR は領収書を持って現れる必要があります。
意味のある変更の概要。
影響を受けるフローに関するブラウザベースの QA。
スクリーンショット、リプレイ、またはその他の行動の証拠。
何か障害が発生した場合のコンソールとネットワークのログ。
既知のリスク、スキップされたケース、未解決の質問。
同じブランチの問題を修正するためのパス。
しかし、それが問題なのです。私たちは、PM、デザイナー、マーケター、サポートに直接貢献してもらいたいと言っていますが、レビューする前に彼らが上級エンジニアのように行動することを期待しています。
PM は、厳密な diff をスコープする方法を知る必要はありません。設計者はネットワーク トレースを読み取るべきではありません。マーケティング担当者は QA を行うべきではありません。サポートは、修正を提案するためだけに完璧なテスト計画を作成すべきではありません。
参入障壁は低く保つ必要があります。レビューバーを上げる必要があります。
ギャップを埋めるための解釈層が存在する場合、これらは矛盾しません。私たちはすでに素晴らしい AI を持っているのに、エンジニアが時間を無駄にする前に、なぜそれを利用して品質をレビューし、意図を解釈しないのでしょうか?
寄稿者は、何が問題なのか、なぜそれが重要なのか、何が良いのかなど、製品のコンテキストを伝えることができます。そして、彼らは AI エージェントと協力して情報を送信することもできます。

d そもそもPR。次に、レビュー ツールチェーンはその実装を開発者が信頼できるものに変換する必要があります。
ツールチェーンは、差分スコープを維持し、実際の変更を要約し、チェックを実行し、ブラウザで製品を開き、フローをクリックし、スクリーンショットをキャプチャし、コンソールエラーを表示し、テストしなかったものにフラグを立てる必要があります。これにより、貢献者はリリース エンジニアに転向することなく、同じブランチの問題を修正できるようになります。
また、ボタンが機能しないことを最初に発見するのは開発者でなくても済むはずです。
では、レビューの自動化とは実際にはどのようなものなのでしょうか?
誰もがこの問題に気づき始めています。そして、PR レビューの自動化が最良の答えであると思われます。そうは言っても、既存の PR レビュー ツールの多くは、その機能がかなり表面的なものであり、コードがコンテキスト内で意味をなすかどうかを確認する別の AI エージェントとして機能するだけであることがわかりました。
実際に必要なのは、ブラウザーでコードを実行し、実際のエッジケースをテストして障害モードを特定するエージェントです。十分な CI 接着剤を使用すれば、間違いなく自分で組み立てることができます。または、棚から入手することもできます。
それが私たちの (ビルダーの) 品質レビュー エージェントのポイントです。実際のブラウザでアプリを開き、影響を受けるフローをたどり、リプレイ リンク、コンソール エラー、ネットワーク トレース、および変更に関連付けられた特定の調査結果を含む、クリックした内容、何が起こったのか、何が失敗したかの証拠を返します。
したがって、謎の差分として始まる何百もの PR をレビューする代わりに、製品固有のレビュー パケットを取得できるようになります。
影響を受けるフロー、リプレイ、スクリーンショット。
コンソールとネットワークは実行からの信号を送信します。
特定の失敗は変更に結びついています。
リスク、スキップされたケース、および残りの判断コール。
結局のところ、目的は開発者をレビューから外すことではありません。 Qを上げるにはまだそこにいる必要があります

私たちだけが知っている方法でコードの性質を確認します。しかし、目標は、マシンが文句を言わずに処理できる準備作業に開発者の注意を向けないようにすることです。
見て。 AI により、誰でも非常に簡単にコードを配布できるようになりました。できなかったのは、魔法のようにそのコードを信頼できるものにすることでした。そしてそれは、開発者たちが今、そのすべてを見直す必要があり、最も負担を感じていることを意味します。
全員をリポジトリから締め出すことは解決策ではありません。すべての広報担当者に、探偵ごっこで午後を無駄にするのではなく、実際に判断力を発揮できるという十分な背景と証拠を提示してもらうだけです。
現在のエージェント ツールを使用すると、これは自分で組み立てることができる信頼レイヤーであり、別の会社から入手することもできます。それに対する私たちの見解は、 Builder QR Agent です。
いずれにせよ、会社の人的マージキューに入る前に、その苦痛を優先することが最善かもしれません。
リアルタイムのコラボレーション、並列エージェント、ビジュアル編集。現在、チーム全体が Al を使用して実際のコードを出荷しています。
AI 4 分 POGR が 3 万ドルと 1 年間の UI 作業をビルダーで削減した方法 執筆者: Amy Cross 2026 年 6 月 4 日 AI 8 分 エージェントティック開発プラットフォームを実装する前に尋ねるべき 5 つの質問 執筆者: Amy Cross 2026 年 5 月 28 日 AI 7 分 AI 製品ラダー (およびほとんどのアプリが標準である理由)

[切り捨てられた]

## Original Extract

AI makes PRs cheap, but trust is still expensive. Why developers are stuck reviewing endless AI slop and how automated code review can fix the bottleneck.

I Didn't Become a Developer to Review AI Slop Skip to main content See how Frete cut frontend build time by 70%
What are best AI tools? Take the State of AI survey
Generate production-ready web apps and UIs
Generate, iterate, and optimize pages and headless content
Code faster with AI and your design system
Design with code, handoff with confidence
Ship prototypes that accelerate the entire software development lifecycle
Generate and publish experiences without developer dependency
Latest insights, tutorials, and announcements
Our library of webinars, case studies and guides
Upcoming events and recorded sessions
Real-world success stories with Builder
Join the discussion, ask questions
Explore Builder partners and connect with a team
Integrate with your stack and connect your tools
Latest features and improvements
Learn to vibe code in new or existing apps
Use Builder to publish content for your site
Convert Figma designs into real code
See how Frete cut frontend build time by 70%
What are best AI tools? Take the State of AI survey
‹ Back to blog AI I Didn't Become a Developer to Review AI Slop
But lately, that's exactly what the job feels like.
My PR queue fills with work that, yes, technically compiles. The summary sounds plausible. It might even have some tests. Then when I open the diff, the real work starts.
What was the change supposed to do? Did anyone actually run the flow? Why is this helper duplicated six times? Is this actually fixing a bug, or did the AI just run around in circles and call it done?
AI made it effortless for anyone on my team (and yours) to create code, but it didn't make that code trustworthy.
Stack Overflow's 2025 Developer Survey found the most common frustration with AI tools is output that's "almost right, but not quite." Sonar's 2026 State of Code report found that 96% of developers don't fully trust AI-generated code, and 38% say reviewing it takes more effort than reviewing human-written code.
That's because AI code looks fine, but you have to really dig in to see what it's doing well. Straight up bad code is much easier to reject.
I'm annoyed. Maybe you are, too. Let's dig into this and solve it together.
A PR is... almost too cheap now
AI agents can spin up branches from Jira tickets, patches from Slack threads, or even full PRs from a bug report before anyone even agrees that the bug is real. It's honestly a pretty awesome world.
But the thing is, developers aren't the only ones using these tools. PMs will prototype the feature they've been trying to explain for three sprints, mostly with vague, unhelpful hand waves. Designers will tweak UX flows and fix layouts that keep getting deprioritized. Marketers will update landing pages and forms. (Constantly.) Support will patch the customer pain points they know best.
And all that is a win. Small fixes shouldn't sit in backlog hell waiting for an engineer who happens to know that part of the code. Product knowledge should be turning into working software faster.
But the easier it gets to open a PR, the more developers are obligated to review them. And PRs aren't valuable just because they exist . They're only valuable when they can be trusted.
And trust is still really expensive
AI is really good at writing code. For a recent hackathon, I had GPT 5.5 spin up 10,000 lines of working code in about 45 minutes. The app mostly worked. Sure, the UI was a nightmare, but the core functionality was there.
But writing code and writing trustworthy, scaleable code are two different things . A model can generate a diff, explain it, and even run some happy-path tests. But someone's still accountable to the stuff that actually matters:
Did this code actually fix the stated problem?
Did the author really understand the system, or is this creating tech debt for later?
Is the diff bigger than it needs to be? (Almost definitely.)
Does this fix silently break some other flow in the code, that would be obvious if a single user just tried it out?
Does the UI actually work for real users in real browsers?
Will this fix survive past a demo?
Is this actually a fix to the root problem, or just a bandaid?
Is this security tradeoff acceptable?
These aren't syntax questions. They're trust questions. And right now, they all land on you and me, the developers. @richiemcilroy put it well in a viral tweeted video the other day:
The numbers tell the same story. LinearB's 2026 benchmarks found AI PRs sit waiting 4.6x longer for review and get rejected way more than human-written ones. METR's study of experienced open-source developers found early-2025 AI tools actually made devs 19% slower, partly because real work includes style, tests, docs, and review—not just typing.
That's not saying AI is useless. And the tools really do keep getting better everyday. But the real work of software was never just typing code into files. It's knowing what should change, what shouldn't, and when a surface-level patch that technically fixes the problem is actually going to haunt your team for the next six months.
That's where your attention needs to go. You should be weighing the stuff that needs taste and context, not manually rediscovering the basics after the PR is already in your queue.
Developing feels bad right now
Even though AI tools are making everyone more productive, being the bottleneck feels terrible . Everyone else gets to accomplish more than they've ever done before, because suddenly code is open to them.
You as a developer just experience the hype as incoming review debt. You aren't building. You're reviewing. You aren't designing the system; you're policing its edges. You aren't solving the hard problem directly; you're reverse-engineering what an agent or teammate was trying to do, then betting your afternoon on whether the diff is safe to keep.
The AI gets to do the fun part. You get to be a robot.
That doesn't mean you're useless. If anything, your judgment matters way more now.
But the workflow is spending your judgment terribly. It's taking the scarcest resource in the system— experienced engineering attention —and aiming it at mystery diffs, bloated patches, missing context, and generated code that only looks correct.
So yeah, it's boring. Yeah, it's frustrating. When someone says "now everyone can ship code," what you and I hear is "now everyone can create work for us."
Locking down the repo solves the wrong problem
So, what do we do? Well, the obvious reaction would be to lock up the repo. Devs only.
And I get that. You're the one who gets paged at 2am when prod goes down. Being protective of the code isn't elitism. You just have a memory.
But limiting access solves the wrong problem.
Cross-functional PRs aren't automatically bad. In fact, in many ways, they're exactly what we've wanted for years: product knowledge turning into small fixes without waiting on an engineer's calendar.
But the problem is that, even though everyone can now open PRs, PR intake itself hasn't evolved. Teams still treat a PR like a dev-to-dev handoff: here's the diff, here's the description, good luck. That worked great when the author was another engineer with the same local context, the same testing habits, and the same gut sense of what reviewers needed.
But that assumption falls apart now. Not because non-devs are careless. In fact, designers, PMs, marketers, support teams—they all have the best user context since they're closer to the problem. But they probably don't know what you need as a dev to evaluate risk. And when AI generated the actual implementation, even the person opening the PR might not know the full scope of what changed.
Mystery diffs aren't a reasonable way to collaborate. So, how do you change the way you work with PRs?
Raise the bar for evidence on PRs
No dev should open a generated or cross-functional PR and have to reverse-engineer it from scratch. Every PR needs to show up with receipts:
A summary of meaningful changes.
Browser-based QA on the affected flow.
Screenshots, replay, or other behavioral proof.
Console and network logs when something is failing.
Known risks, skipped cases, and open questions.
A path to fix issues on the same branch.
But that’s the problem. We say we want PMs, designers, marketers, and support to directly contribute , but then we expect them to act like senior engineers before we'll even review it.
A PM shouldn't need to know how to scope a tight diff. A designer shouldn't read network traces. A marketer shouldn't be QA. Support shouldn't write a perfect test plan just to propose a fix.
The entry bar needs to stay low. The review bar needs to go up.
Those aren't in conflict if there's an interpretation layer to bridge the gap. We already have amazing AI, so why aren't we using it, per PR, to review the quality and interpret intent before engineers waste their time?
The contributor can bring the product context: what hurts, why it matters, what good looks like. And they can be the ones who work with an AI agent to send the PR in the first place. Then, a review toolchain should translate that implementation into something a dev can trust.
The toolchain should keep diffs scoped, summarize real changes, run checks, open the product in a browser, click the flow, capture screenshots, surface console errors, and flag what it didn't test. It should let the contributor fix issues on the same branch without turning them into a release engineer.
And it should spare the developer from being the first person to discover the button doesn't work.
So, what does review automation look like in practice?
Everyone's starting to wake up to this problem. And PR review automation seems to be the best answer. That said, I've found that a lot of the existing PR review tools are pretty surface-level in what they do, mostly just acting as another AI agent to see if the code makes sense in context.
What you actually want is an agent that runs the code in the browser and tests real edge cases to spot failure modes. You can definitely piece it together yourself with enough CI glue. Or, you can get it off the shelf.
That's the point of our (Builder's) Quality Review Agent . It opens your app in real browsers, walks the affected flow, and returns evidence of what it clicked, what happened, and what failed, complete with replay links, console errors, network traces, and specific findings tied back to the change.
So now, instead of reviewing hundreds of PRs that start as mystery diffs, you get a product-specific review packet:
The affected flow, replay, and screenshots.
The console and network signals from the run.
The specific failures tied back to the change.
The risks, skipped cases, and remaining judgment calls.
After all, the goal isn't to remove developers from review. We still need to be there to raise the quality of the code in ways only we know how. But the goal is to stop burning developer attention on prep work that machines can handle without complaint.
Look. AI made it dead simple for anyone to ship code. What it didn't do was magically make that code trustworthy. And that means devs are feeling the burden the most right now, having to review all that slop.
Locking everyone out of the repo isn't the answer. We just need every PR to show up with enough context and proof that we can actually use our brains for judgment instead of wasting afternoons playing detective.
With today's agentic tools, that's a trust layer you can either try to assemble yourself, or you can get it from another company. Our take on it is the Builder QR Agent .
Regardless, it might be best to prioritize that pain before you turn into your company's human merge queue.
Real-time collaboration, parallel agents, and visual editing. The whole team ships real code with Al now.
AI 4 MIN How POGR Cut $30K and a Year of UI Work with Builder WRITTEN BY Amy Cross June 4, 2026 AI 8 MIN 5 Questions to Ask Before Implementing an Agentic Development Platform WRITTEN BY Amy Cross May 28, 2026 AI 7 MIN The AI Product Ladder (and why most apps are st

[truncated]
