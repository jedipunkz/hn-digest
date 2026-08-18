---
source: "https://nolanfrausto.com/posts/ai-slop-spiral/"
hn_url: "https://news.ycombinator.com/item?id=49348311"
title: "The AI Slop Spiral"
article_title: "The AI Slop Spiral — On Resonance"
image: "https://nolanfrausto.com/posts/ai-slop-spiral/opengraph-image.png"
author: "frausto"
captured_at: "2026-08-18T17:19:24Z"
capture_tool: "hn-digest"
hn_id: 49348311
score: 4
comments: 1
posted_at: "2026-08-18T16:37:38Z"
tags:
  - hacker-news
  - translated
---

# The AI Slop Spiral

- HN: [49348311](https://news.ycombinator.com/item?id=49348311)
- Source: [nolanfrausto.com](https://nolanfrausto.com/posts/ai-slop-spiral/)
- Score: 4
- Comments: 1
- Posted: 2026-08-18T16:37:38Z

## Translation

タイトル: AI スロップスパイラル
記事のタイトル: AI スロップ スパイラル — 共鳴について
説明: AI は信じられないほど強力なツールですが、考えているように見えるものを生成しながら、考えないようにすることも簡単にします。その結果、出力は速くなり、引き継ぎは悪化し、組織プロセスは徐々に悪化していきます。

記事本文:
AI スロップ スパイラル — On Resonance On Resonance 執筆について 購読 すべての執筆 ノーラン フラウスト 2026 年 8 月 18 日 13 分で読む AI スロップ スパイラル
AI is an incredibly powerful tool, but it also makes it easy not to think while producing something that looks very much like thinking.その結果、出力は速くなり、引き継ぎは悪化し、組織プロセスは徐々に悪化していきます。
I open the document. My computer screen glares brightly as the sun sets outside my home office windows. A dying lamp hums faintly in the background, interrupted only by the click of the mouse wheel as I scroll.
少なくともこれはわずか 12 ページです。
It is the seventh document I have received today. A steady stream of PRDs, technical specifications, incident reports, and QA testing plans.それぞれが印象的に見えます。 They are thorough, professionally formatted, and full of tables, edge cases, diagrams, and implementation details.
At a surface level, they all seem smart and comprehensive, and it gives the confidence that the organizational engine continues to churn.作業はプロセスを進め続けます。
最新のPRDを引き上げます。 It is long, however, and I have a growing pile to get through, not to mention all the other tasks on my plate. I put it into AI and ask for a summary, risks, and priorities.返答を受け取ったら、時間をかけて頭の中でプロジェクトを大まかに概念化し、いくつかの質問をして、いくつかの仮定を検証します。 AI は、要件が明確で重要な詳細が説明されていることを保証するため、下流の誰かが特定の実装の詳細を検証する必要があることをいくつかのメモで承認します。
The PRD is asynchronously handed off to engineering and design.デザインでは AI を使用してアセットとモックアップを構築します。 Engineering uses

AI to turn it into a technical plan. QA は AI を通じてそれを実行し、合格基準を生成します。
これまでよりも早く、プル リクエストが受信され、コードが完成します。 Copilot はすでにそれをレビューし、エラー処理、名前付け、および起こり得るエッジケースに関する一連のメモを残しました。数人のコーディング エージェントとやり取りを行った後、コードは AI と別の実際の人間のエンジニアの承認を得て、テストに移ります。
QA は AI を使用して、AI が生成した PRD から AI が生成した合格基準に照らして実装を検証し、上層部の意識的なメモの AI ストリームから AI が抽出した設計を支援します。 PM はサインオフします。デザインのサインオフ。 QA はサインオフします。全員が本物の人間です。
The feature goes to production in record time.
Several events are missing. Important data is not being populated.複数のアカウントを作成できる重複レコードのバグがあり、たとえそれがなくなっていたとしても、この機能は実際には構想時に主要な関係者が想定していたものではなかったことがわかりました。
次の 2 週間は、トリアージ、レベル設定、何が起こったのかを解明することに費やされます。
会議が予定されています。 Requirements are clarified.所有権については議論がある。テストが追加されます。チームは元の文書と資産を遡って、特定されていないものもあれば、技術的には特定されているものの異なる解釈がされているものもあること、いくつかの重要な前提が誰かの頭の中にのみ存在していたか、まったく存在していなかったことを発見しました。
Eventually, we ship the feature again.しかし今では、本来よりもはるかに長い時間がかかっています。 Multiple stakeholders are frustrated.チームはプロセスとお互いに対する信頼をいくらか失いました。
それでも、最初から、個々のステップはより速く進んでいるように見えました。
AI has changed the way we work.
これは来ません。それはもう幸せだ

ネッド。
それは非常に強力なツールです。これにより、人々はこれまでよりも迅速に物事を記述、調査、分析、設計、構築できるようになります。ほんの数年前にはばかげているように思えたであろう量の作品を、一人の人間が生み出すことができます。
しかし、AI のおかげで、考えないようにすることもずっと楽になりました。さらに悪いことに、思考しているように見えるものを生成しながら、思考しないことが可能になってしまいました。
12 ページの PRD は厳格さを示しています。対応する技術仕様には、専門用語と理解がにじみ出る詳細が満載されています。洗練されたモックアップのセットを見ると、ユーザー エクスペリエンスが慎重に考慮されているように見えます。 40 の合格基準は、包括的でよく考えられた範囲のように見えます。数十のコメントといくつかの後続のコミットを含むプル リクエストは、慎重なレビューのように見えます。
私たちは現在、何が構築されているかを誰も理解できない世界に住んでいます。さらに言えば、これはどの時点においても、チームが行われている作業とその影響範囲について共通の理解を持っていなかったということを意味するものではありません。
AI 製品要件ドキュメントが作成される → AI モックアップと資産が作成される → AI 技術仕様が作成される → AI 受け入れ基準が作成される → AI 生成コードが作成される → AI コード レビューが作成される → AI テストこれはチーム間およびチーム内で行われる電話ゲームであり、あらゆる段階で AI が参加します。コンテキストは圧縮され、変換され、拡張され、そして再び変換され、その過程で共通の理解が失われます。
私はこれを「AI スロップ スパイラル」と呼んでいます。
本当の危険は、AI が悪い仕事を生み出すことではありません。公平を期すために言うと、人間は常に悪い仕事を生み出してきました。危険なのは、AI が完成したように見える仕事を生み出し、それがどのような意味であっても、それを完成させるのに十分な強力な仕事を生み出すことです。
漠然とした考えを洗練された文書に変えます。

思考は完全に形成されました。ギャップを埋め、矛盾を滑らかにし、実際には誰も議論しなかった詳細を自信を持って提供します。これらの詳細は合理的である可能性さえありますが、合理的というのは合意されたものと同じではありません。
出力が洗練されているように見えるほど、それに挑戦する人は少なくなります。そして、管理する AI エージェントの数が増えるほど、綿密に監視する時間が減ります。乱雑な 2 ページの文書が質問を招きます。見出し、マトリックス、図、および「エッジケースと故障モード」というセクションを含む 12 ページの文書は、質問がすでに答えられているかのような印象を与えます。
だから人々はそれを流し読みするのです。流し読みします。最近はそれが多すぎて、そうでなかったら一日中ドキュメントとコードのレビュー以外何もしていないでしょう。次に、その文書が完全であるかどうかを AI に尋ねると、文書を文書として評価するのが得意な AI は、完全であると答えます。
これにより、組織全体にある種の人為的な自信が生まれます。 PRD は詳細に記述されているため、PM はエンジニアリングが要件を理解していると想定します。エンジニアリングは、PRD が承認されたため、要件が解決されたと想定しています。 QA は、期待される動作が PRD から取得されたものであるため、その動作が正しいことを前提としています。リーダーシップは、生産量が増加し、チケットの処理が速くなっているため、プロセスが機能していると想定します。
私たちは皆、AI によって生産性が 100% 向上するはずなので、結果を示す必要があります。人間同士のやりとりは、煩雑で遅く見え始めます。
局所的には誰もが正しい。時にはミーティングも行うこともあります。全員がドキュメントを読み進め、それが理にかなっていることに同意し、前進する準備ができていると言います。 AIが良いと言ったからです。誰もが自分が作成したドキュメントをポイントできます。
むしろプロジェクトは失敗し、組織は

誰もすべてを理解していない完璧な記録を作成しました。
私たちは皆、このことにある程度満足しています。
あらゆるテクノロジーの誇大宣伝サイクルには恐怖が存在しますが、特に AI に関してはその恐怖が強烈です。チャンスを逃すのではないかという恐怖、競合他社がより速く動いているのではないかという恐怖、そしてバカのように会議をしている間に別の会社が魔法のような新しい働き方を発見してしまうのではないかという恐怖があります。
したがって、方向は上から下になります。
「AIをもっと活用しましょう。」 「トークンを最大化します。」 「AIファースト」 「コードの 99 パーセントは AI によって書かれるべきです。」などなど。
正しい実装がどのようなものかを実際には誰も知らないため、これらの指示は通常あいまいです。そうすれば、その指示はさらに役立つでしょう。どの作業が AI から恩恵を受けるのか、どの作業が人間の直接的な判断を必要とするのか、成功とはどのようなものなのか、どのような失敗モードを測定する必要があるのか​​を特定します。
むしろ、AIの活用自体が目的になってしまいます。
これは、これらのシステムを機能させるという日常の現実の中で生きる必要のない AI 企業、コンサルタント、投資家、経営陣によるマーケティングによって促進されています。
彼らはデモを見ます。彼らはベンチマークの改善を確認しています。彼らは、ある人が 15 分でアプリケーションを構築するのを目にしました。 AI の使用が、存在するすべての可能性を解き放つ黄金の鍵であるとあらゆることが叫ばれていますが、当然、誰もが AI を使用する必要があります。
その後、ダウンタイムが増加します。手戻りが増える。コストが増加します。プロジェクトにはどういうわけか時間がかかります。人々はシステムがどのように機能するのかを見失います。すべての生産性指標が改善しているように見えるのに、なぜ生産性が悪く感じるのか誰も説明できません。
私たちはアウトプットを最大化しており、アウトプットは進歩と同じであると想定しています。
AI は、私たちがその幻想を維持するのに非常に優れています。
そのすべて

これは実際には AI の問題ではなく、AI が非常に得意とする人間の問題です。人間はどこまでも怠惰な生き物です。
私たちは創造的で、知的で、順応性があり、信じられないほど優れています。しかし、複雑な問題を深く理解するか、深く理解しているように見える何かを生成するかの選択を考慮すると、多くの人はしばしば 2 番目の選択肢を選択します。そして AI は、私たちが綿密であるように見せながら怠け者にさせることで、その選択を信じられないほど簡単にします。
さらに良いことに、怠惰は会社によって義務付けられていることがよくあります。私たちは大変な努力を怠っているわけではありません。私たちは「AIの活用」を行っています。私たちは AI 利用の最前線に立つ必要があると絶えず言われてきました。これは双方にとって有利です。
機能が動作しないことを除いて。または、正常に動作しますが、静かにデータを破損します。または、2 日で起動し、安定するまでに翌月かかります。あるいは 6 か月後、生成されたシステムを安全に変更できるほど十分に理解している人は誰もおらず、かつては単純なソリューションだったものが飛躍的に複雑になってしまいました。
問題は、AI がコードやドキュメントを書いたことではありません。問題は、その人が理解できないことに対して責任を負ったことです。
最近、私は人々に、送られてくる書類を順番に見てもらうように頼み始めました。それが PRD であれ、コードであれ、テスト計画であれ、その他のものであれ。私との会議に参加して、非同期で説明してください。チームとして話し合ってください。共通の理解が暗黙的ではなく明示的であることを確認し、少なくともプロセスのその部分に AI が関与していないことを確認してください。そしてそのたびに、私たちは何か重要なものを見つけます。
著者がセクションを説明できない場合があります。場合によっては、明らかにすべてを読んでいなかったり、書かれているすべての部分を本当に理解していなかった場合もあります。場合によっては、文書に同意できない内容が記載されていることがあります。たまには二人で

まったく異なる行動を想像しながら、同じ言葉を使っているのです。
このことだけを聞くと、私が雲に向かって物事が以前と違うと叫んでいる老人のように聞こえるかもしれないことはわかっています。はっきりさせておきますが、私は反AIではありません。まったく逆です。私はAI推進派です。好例: このブログ投稿の編集に AI を使用しました。そこで重要な単語は edit です。最初のバージョンは自分で書いて、後からAIにクリーンアップしてもらって、その後また編集しました。執筆においてそれを最大限に活用する方法をまだ考え中ですが、この特定の投稿を書くときに取り組むのは有意義な練習だと思いました。
AI は世界を変えるテクノロジーです。これは非常に強力なツールであり、私は常にそれを使用しています。私のチームにも使ってもらいたいです。私は、これがほぼすべての技術的およびクリエイティブなワークフローの通常の一部になることを期待しています。
現状でも、私の大げさな終末論的な例は別として、それは全体として非常に大きなプラスでした。しかし、よく言われるように、大いなる力には大いなる責任が伴います。これほど変化の可能性があるものには、必ず調整期間があり、変化自体によって新たな問題が生じることになります。
しかし、それは魔法の杖ではありません。チームがあらゆるプロジェクトで使用できる AI フォワード共有コンテキスト システムを構築する場合でも、理論的には

[切り捨てられた]

## Original Extract

AI is an incredibly powerful tool, but it also makes it easy not to think while producing something that looks very much like thinking. The result is faster output, worse handoffs, and an organizational process that slowly spirals into slop.

The AI Slop Spiral — On Resonance On Resonance Writing About Subscribe All writing Nolan Frausto August 18, 2026 13 min read The AI Slop Spiral
AI is an incredibly powerful tool, but it also makes it easy not to think while producing something that looks very much like thinking. The result is faster output, worse handoffs, and an organizational process that slowly spirals into slop.
I open the document. My computer screen glares brightly as the sun sets outside my home office windows. A dying lamp hums faintly in the background, interrupted only by the click of the mouse wheel as I scroll.
At least this one is only twelve pages.
It is the seventh document I have received today. A steady stream of PRDs, technical specifications, incident reports, and QA testing plans. Each one looks impressive. They are thorough, professionally formatted, and full of tables, edge cases, diagrams, and implementation details.
At a surface level, they all seem smart and comprehensive, and it gives the confidence that the organizational engine continues to churn. Work continues moving through the process.
I pull up the latest PRD. It is long, however, and I have a growing pile to get through, not to mention all the other tasks on my plate. I put it into AI and ask for a summary, risks, and priorities. Once I get the response, I then take the time to roughly conceptualize the project in my head and ask a few questions and validate several assumptions. AI assures me that the requirements are clear and the important specifics have been accounted for, so I sign-off with a few notes that someone downstream should validate certain implementation details.
The PRD is asynchronously handed off to engineering and design. Design uses AI to build out assets and mockups. Engineering uses AI to turn it into a technical plan. QA runs it through AI to generate acceptance criteria.
Faster than ever, a pull request comes in, the code is done. Copilot has already reviewed it and left a series of notes about error handling, naming, and possible edge cases. After some back and forth with a few coding agents, the code gets the approval of AI and another real human engineer and it moves to testing.
QA uses AI to help validate the implementation against the acceptance criteria that AI generated from the AI-generated PRD and designs that AI distilled from someone higher-up’s AI stream of conscious notes. The PM signs off. Design signs off. QA signs off. All are real humans.
The feature goes to production in record time.
Several events are missing. Important data is not being populated. There is a duplicate-record bug that can allow someone to create multiple accounts, and even if it had gone out, we learn that the feature was not actually what was envisioned by the key stakeholders at conception.
The next two weeks are spent triaging, level-setting, and trying to figure out what happened.
Meetings are scheduled. Requirements are clarified. Ownership is debated. Tests are added. The team goes back through the original documents and assets, and discovers that some things were never specified, other things were technically specified but interpreted differently, and a few critical assumptions existed only in someone’s head or not at all.
Eventually, we ship the feature again. But now it has taken much longer than it should have. Multiple stakeholders are frustrated. The team has lost some trust in the process and in each other.
Still, from the outset, every individual step appeared to move faster.
AI has changed the way we work.
This is not coming. It already happened.
It is an enormously powerful tool. It allows people to write, research, analyze, design, and build things faster than ever. A single person can produce an amount of work that would have seemed absurd only a few years ago.
But AI has also made it much easier not to think. Worse, it has made it possible not to think while producing something that looks very much like thinking .
A twelve-page PRD gives the appearance of rigor. The corresponding technical specification is filled with jargon and detail that exudes understanding. A set of polished mockups makes it look like the user experience has been carefully considered. Forty acceptance criteria look like comprehensive and well-thought coverage. A pull request with dozens of comments and several follow-on commits looks like a careful review.
We live in a world now where none of those things mean that anybody understands what is being built. More to the point, none of this means that at any point did the team have a shared understanding of the work being done and its impact area.
An AI product requirement document begets → AI mockups and assets begets → AI technical specification begets → AI acceptance criteria begets → AI-generated code begets → AI code review begets → AI testing. It is a game of telephone played between teams and within teams, with AI participating at every step. Context is compressed, transformed, expanded, and then transformed again, and shared understanding is lost along the way.
I call this The AI Slop Spiral.
The real danger is not that AI produces bad work, to be fair, humans have always produced bad work. The danger is that AI produces work that looks finished, and is powerful enough to finish it, in whatever way that means.
It turns a vague thought into a polished document before the thought has been fully formed. It fills in the gaps, smooths over the contradictions, and confidently supplies the details that nobody actually discussed. Those details may even be reasonable, but reasonable is not the same as agreed upon.
The more polished the output looks, the less likely someone is to challenge it. And the more AI-agents one is managing, the less time there is to monitor closely. A messy two-page document invites questions. A twelve-page document with headings, matrices, diagrams, and a section called “Edge Cases and Failure Modes” creates the impression that the questions have already been answered.
So people skim it. I skim it. I get so much of it lately, I’d be doing nothing but reviewing documentation and code all day if I didn’t. Then we ask AI whether the document is complete, and AI, being very good at evaluating documents as documents, tells us that it is.
This creates a kind of artificial confidence throughout the organization. The PM assumes engineering understands the requirements because the PRD is detailed. Engineering assumes the requirements are settled because the PRD was approved. QA assumes the expected behavior is correct because it was pulled from the PRD. Leadership assumes the process is working because output is up and tickets are moving faster.
We are all supposed to be getting a 100 percent productivity increase from AI, so we need to show results. Any back-and-forth between humans starts to look messy and slow.
Everyone is locally correct. Sometimes we even have the meetings. Everyone walks through the documentation, agrees that it makes sense, and says we are ready to move forward. Because AI told us it was good. Everyone can point to the document they created.
Instead the project fails and the organization has produced a perfect record of nobody understanding the whole thing.
We are all complacent in this to a degree.
There is a fear that exists in every technology hype cycle, but it is especially intense with AI. There is the fear of missing out, the fear that competitors are moving faster, and the fear that another company has discovered some magical new way of working while you are still holding meetings like an idiot.
So the direction comes down from the top:
“Use more AI.” “Maximize tokens.” “AI first.” “Ninety-nine percent of code should be written by AI.” So on and so forth.
These instructions are usually vague because no one actually knows what the correct implementation looks like. If they did, the instruction would be more useful. It would identify which work benefits from AI, which work requires direct human judgment, what success looks like, and what failure modes should be measured.
Instead, AI usage itself becomes the goal.
This is encouraged by the marketing coming from AI companies, consultants, investors, and executives who do not have to live inside the daily reality of making these systems work.
They see the demos. They see the benchmark improvements. They see a person build an application in fifteen minutes. And with everything screaming at you that AI use is the golden key that will unlock all potential in existence, well then of course everyone should be using AI.
Then downtime increases. Rework increases. Costs increase. Projects somehow take longer. People lose track of how the systems work. Nobody can explain why productivity feels worse when every productivity metric appears to be improving.
We are maximizing output and assuming that output is the same thing as progress.
AI is very good at helping us maintain that illusion.
All that said, this is not really an AI problem, it is a human problem that AI is very good at exploiting. Humans are lazy creatures, all told.
We are also creative, intelligent, adaptable, and incredible. But given the choice between deeply understanding a complex problem and generating something that looks like we deeply understand it, many will often choose the second option. And AI makes that choice incredibly easy by letting us be lazy while appearing thorough.
Better yet, the laziness is often mandated by the company. We are not skipping the hard work. We are “leveraging AI.” We’ve been told non-stop that we need to be at the forefront of AI usage. This is a win-win.
Except the feature does not work. Or it works in the happy path but quietly corrupts data. Or it launches in two days and takes the next month to stabilize. Or six months later, nobody understands the generated system well enough to safely change it, and what was once a simple solution has become exponentially more complex.
The problem is not that AI wrote the code or the document. The problem is that a person accepted responsibility for something they did not understand.
Recently, I started asking people to walk me through the documents that they send my way. Whether that's a PRD, or code, a testing plan, or something else. Get in a meeting with me and walk me through it, non-async. Talk it over as a team. Make sure the shared understanding is not implicit but explicit and make sure AI is not involved in at least that part of the process. And every single time, we find something important.
Sometimes the author cannot explain a section. Sometimes they clearly did not read all of it or really understand every piece written. Sometimes the document says something they do not agree with. Sometimes two people have been using the same words while imagining entirely different behavior.
I know this whole thing may make me sound like an old man yelling at clouds about how things aren’t what they used to be. Let me be clear: I am not anti-AI. Quite the opposite. I am pro-AI. Case in point: I used AI to edit this very blog post. The important word there is edit . I wrote the initial version myself and had AI clean it up afterward, and I edited it again after that. I'm still figuring out how to best leverage it with writing, but I figured it was a meaningful exercise to undertake when writing this specific post.
AI is a world-changing technology. It is an extraordinarily powerful tool, and I use it constantly. I want my teams to use it. I expect it to become a normal part of nearly every technical and creative workflow.
Even as it stands, my hyperbolic doomsday examples aside, it has generally been a huge net positive. But, as they say, with great power comes great responsibility. Anything with this much potential for change is going to come with a period of adjustment, along with new problems created by the change itself.
But it is not a magic wand. Even when we build AI-forward shared context systems for teams to use on every project, such that theore

[truncated]
