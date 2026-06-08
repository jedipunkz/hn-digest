---
source: "https://addyosmani.com/blog/intent-debt/"
hn_url: "https://news.ycombinator.com/item?id=48444373"
title: "AI Agents Can't Help with Intent Debt"
article_title: "AddyOsmani.com - The Intent Debt"
author: "moebrowne"
captured_at: "2026-06-08T13:35:22Z"
capture_tool: "hn-digest"
hn_id: 48444373
score: 3
comments: 1
posted_at: "2026-06-08T12:17:15Z"
tags:
  - hacker-news
  - translated
---

# AI Agents Can't Help with Intent Debt

- HN: [48444373](https://news.ycombinator.com/item?id=48444373)
- Source: [addyosmani.com](https://addyosmani.com/blog/intent-debt/)
- Score: 3
- Comments: 1
- Posted: 2026-06-08T12:17:15Z

## Translation

タイトル: AI エージェントは意図的な負債を解決できない
記事のタイトル: AddyOsmani.com - 意図された負債
説明: 技術的負債はコード内に存在します。認知的負債は頭の中に住んでいます。インテント負債は、あなたが書いたことのない成果物、つまり目標、制約、...の中に存在します。

記事本文:
AddyOsmani.com - 意図された負債
ホーム
GitHub
プレス
略歴
話し合い -->
マストドン -->
リンクトイン
ツイッター
購読 -->
ニュースレター
ショップ -->
ブログ
意図された負債
技術的負債はコードの中に存在します。認知的負債は頭の中に住んでいます。インテント負債は、あなたが書いたことのない成果物、つまり、目標、制約、システムが現在のようになる理由の理論的根拠の中に存在します。運が良ければ、その一部がチームのドキュメントやディスカッションに散在している可能性がありますが、おそらく不完全です。これはエージェントが返済できない一種の借金であり、エージェント エンジニアリングが最も費用がかかります。
マーガレット アン ストーリーの三重債務モデルは、ソフトウェアの健全性について考えるための明確な方法です。負債の 3 つのモデルは、技術的、認知的、意図的です。
技術的負債はコードの中に存在します。複雑なモジュール、期限内に取ったショートカット、漏洩した抽象化など、後からシステムを変更するのが難しくなるのは、実装に関する選択の積み重ねです。私たちはこれを何十年も理解してきました。ビルドの遅さ、テストの脆弱さ、特定のファイルに触れる恐怖などを通じて、それが感じられるでしょう。
認知的負債は人々の中に生きています。それは共有理解の侵食であり、存在するコードの量と人間が理解できる量との間のギャップです。私はこれを理解負債と呼んでいます。システムがチームのメンタルモデルよりも速く成長すると、それが蓄積されます。コードが原始的であっても、誰も原始的なコードを理解できないため、致命的な認知的負債を抱え続ける可能性があります。
意図的負債はアーティファクトの中に存在します。それは、システムがなぜそうなっているのかを説明する、外在化された理論的根拠、目標、制約が欠如しているか侵食されているということです。キーワードは外在化されている。理論的根拠は、チームメイト、将来のあなた、またはエージェントがわかる場所に書き留める必要があります。

頭の中に留めずに読んでください。意図した負債が高くなると、システムは本来の目的から逸脱し、いつ逸脱したのか、なぜ逸脱したのかは誰も言えません。
これら 3 つは独立しているため、理解するのに時間がかかりました。
技術的な負債は少なくても、意図的な負債は大きくなる可能性があります。あなたはシステムを完全に自分自身で理解することができます (あなたにとって認知的負債はありません) 一方で、その意図は頭蓋骨の外には存在しません (他の人にとっては莫大な意図的負債があります)。
内部的には同じように感じられますが、それぞれの料金は別々に請求されます。
なぜ意図的な借金はエージェントが助けられないのか
AI はこれまでよりも速くコードを生成するため、技術的負債の負担と返済が安くなります。複雑なモジュールをエージェントに指示すると、それがリファクタリングされます。
認知的負債も、ほとんどのエンジニアが予想するよりも簡単に回復します。システムの一部が理解できない場合は、エージェントに説明を求めます。コードはまだ存在しており、モデルはそれを読み戻すことができるため、必要に応じて失われたメンタル モデルの一部を再構築します。
意図が違います。インテントはユーザーから入力する必要がある 1 つの入力であるため、エージェントはインテントを生成できません。以前のエンジニアが何かをした理由を推測できるのと同じように、モデルはコードからもっともらしい根拠を推測できます。意図についての推測は意図ではありません。モデルには、その 300 ミリ秒のデバウンスが意図的な UX の決定なのか、ベンチマークの結果なのか、それとも誰かが一度入力しただけで再確認されなかった数値なのかはわかりません。自信満々に聞こえる理由をでっち上げることになるが、それは知らないと認めるよりも悪い。
3 つの債務のうち、意図的債務は代理人があなたを救済できない唯一のものです。コードを書いて理解を取り戻すことができます。その理由は、捏造することしかできない唯一のことです。
エージェントは、記載されていないコストをより迅速に計算します
チームは高い志を持って逃げ切った

私たちはそれを頭と古い文書の中に持ち続けていたため、何年も借金を抱えていました。
新しい人間がチームに加わったとき、すべてを書き留めるわけではありません。廊下での会話やコードレビューのコメント、「2023 年にインシデントがあったため、そのようにはしません」など、時間の経過とともに彼らは意図を理解したからです。知識は人から人へと伝わり、蓄積されていきます。 4 年間在籍していたエンジニアが意図の文書化を担当し、費用がかかり損失も大きかったが、うまくいきました。
エージェントはそのモデルを打ち破ります。エージェントをチームに迎え入れると、長期記憶を持たない若手のチームの規模は一夜にして2倍になります。エージェントはほとんどのセッションをコールドで開始します。そこには、人類が何年にもわたって築き上げてきた暗黙の意図はまったく含まれていません。読み取り可能なアーティファクトとして外部化されていないものは何であれ、読み取り可能ではありません。
それは、物事を書き留めないことの経済学を変えます。外在化されていない意図は、新人研修時や退職後に時々コストがかかることがありました。これからは、セッションごとに、実行するエージェントごとに乗じて料金を支払うことになります。
並列化することにとても興奮している 20 人のエージェントを想像してみてください。それぞれがチームメイトであり、あなたに会ったこともなく、あなたの心を読むこともできず、あなたの意図のギャップをもっともらしい推測で埋めてくれます。私が書いたオーケストレーション税は、一部は意図債務税です。多くのエージェントの管理を疲れさせる原因の多くは、書き留めなかった意図を再提供することです。
理解義務の議論の残り半分
理解による負債について書いたとき、意図的な負債がそれをより鮮明にするため、もう一度見直したい点を書きました。
私は、詳細な仕様は完全な答えではないと主張しました。仕様を動作するコードに変換するには、仕様では決して把握できない膨大な数の暗黙の決定が含まれます。また、プログラムとして十分に詳細な仕様は、より遅い言語でのプログラムになります。私は今でもそう信じています。
意図的な負債は補完的な真実です。
いる

すべての意図をキャプチャできないということは、その意図をまったくキャプチャできないということです。エージェントがユーザーに代わって行う暗黙の決定は、仕様書には決して列挙されないものですが、少なくとも負荷のかかる決定を記録しなければ、その論理的根拠が蒸発してしまいます。すべてを書き留めることはできません。
間違えると高くつく選択の背後にある理由を書き留める必要があります。後でそれらを再構築する人は誰もいないからです。
理解義務は、コードが存在するからといってコードが正しいと信じないように警告します。
インテント負債は、コードが存続するからといって理由が存続すると信じないように警告します。コードが答えです。その意図は、解決することを意図された質問でした。 AI は、書き留め忘れた質問に対する答えを導き出す能力に優れています。
多額の負債がどのようなものか
意図的な負債が摩擦として現れることはほとんどありません。それは、ある種の無力感として現れます。
エージェントはガード句を削除することでバグを「修正」しますが、ガードが存在する理由がドキュメントやコミットメッセージに記録されていないため、そのガードが負荷に耐えていたのか、残っていたのかは誰も言えません。
リファクタリングは、ユーザーが依存する動作を変更します。差分がきれいに見え、テストが緑色であったため、レビューは合格しましたが、テストは以前の動作のみをエンコードしており、意図はエンコードされていませんでした。
なぜ 2 つのサービスが直接通話ではなくキューを介して通信するのかと尋ねると、正直な答えは「エージェントが提案したので問題ないと思われました。」です。その答えは、すでに利息が発生している意図的な借金です。
再構築できない設計上の選択を擁護するという、これの認知的放棄バージョンを感じたことがあるなら、意図的負債は、同じ穴をチーム規模で文書化したバージョンです。
降参とは、その瞬間における自分自身の姿勢に関するものです。インテント負債は、これらの瞬間の 100 が、次の人物と次の代理人が継承するためにレポに残されたものです。
お金を払ってやる

wn: インテントをファーストクラスのアーティファクトとして外部化します。
ここ数か月間、私が書いてきたことのほとんどすべてが、意図的な債務管理であることが判明しました。言葉もありませんでした。動きは毎回同じです。意図を頭から取り出し、エージェントが読める場所に置きます。
実装ではなく、意図の仕様を記述します。優れた仕様には、目標、制約、譲れないもの、および完了の明確な定義 (高速、アクセス可能、安全、快適、「機能的に正しい」を超えたもの) が含まれています。この仕様には、コードだけでは実行できない意図が含まれています。
AGENTS.md を構成ではなく、インテント台帳として扱います。これが、私が /init の使用をやめるよう言い続ける理由です。自動生成されたファイルには、コードが何であるかが説明されています。インテント ファイルには、チームの意味、つまり慣例、「なぜならこのようにはしない」、単一のファイルでは目に見えない制約が記述されます。エージェントはそれを推測することができず、それを最も必要としています。
意思決定が発生した場所でそれをキャプチャします。軽量の意思決定ログ (ADR) は、純粋な意図による負債の返済です。あなたが決断した瞬間の理由を記録するのにほとんど費用はかかりません。事情を知った人物がチームを移動させた後、8か月後に再構築するには莫大な費用がかかる。エージェントは伐採をこれまでよりも安価にしたため、古い言い訳はなくなりました。
学習ループに意図を書き戻すようにします。私は自己改善エージェントがセッションの終了時に学習ファイルを更新することを主張してきました。同じループは、意図と負債のポンプが逆回転するものです。根本原因を記録するすべての間違い、「X を試してみましたが、Y のせいでうまくいきませんでした」はすべて意図であり、そうでなければ嫌な午後の記憶の中にだけ生きていたはずです。
これらはいずれも新しいツールではありません。それは、頭がもはや存在しない時代に、なぜを頭の中にだけ存在させることを拒否する規律です。

ほとんどの作業が行われる場所。
長い間、ソフトウェアにおいて希少で価値のあるものは、正しい実装を作成する能力でした。コードは高価だったので、コードを書くために最適化しました。
AI によりコードは安価になり、理解力は回復可能になりました。目的、制約、理由である意図は、依然として人間によって生み出される必要がある入力の 1 つです。それは私たちが外面化するのが最も苦手な問題でもあります。なぜなら、私たちは何十年も頭の中にそれを持ち続けてきたからです。
これは、チームが何年にもわたって共有されたコンテキストを通じて意図を吸収できる少数の人々であったときに機能しました。チームの半分が見知らぬ人としてすべてのセッションを開始するエージェントである場合、これは機能しません。
技術的負債により、システムの変更が困難になります。認知的負債があると理解が難しくなります。インテント負債があると、システムがまだあなたが望んでいたとおりに動作しているかどうかを知ることが難しくなります。また、それはエージェントがあなたに代わって返済できない 3 つのうちの唯一のものです。その部分は心に残ります。理由を書き留めてください。それは、リポジトリに残すことができる最も価値のあるものになるためです。
もっと知りたいですか？私の無料ニュースレターを購読してください:
免責事項: このサイトで表明された見解や意見は私個人のものであり、必ずしも Google またはその関連会社の見解、立場、戦略を反映するものではありません。

## Original Extract

Technical debt lives in your code. Cognitive debt lives in your head. Intent debt lives in the artifacts you may have never wrote: the goals, constraints, an...

AddyOsmani.com - The Intent Debt
Home
GitHub
Press
Biography
Talks -->
Mastodon -->
LinkedIn
Twitter
Subscribe -->
Newsletter
Shop -->
Blog
The Intent Debt
Technical debt lives in your code. Cognitive debt lives in your head. Intent debt lives in the artifacts you may have never wrote: the goals, constraints, and rationale for why the system is the way it is. If you’re lucky, some of this exists scattered in team documents or discussions, but it’s likely incomplete. It’s the one kind of debt your agents can’t pay down for you, and agentic engineering makes it the most expensive.
Margaret-Anne Storey’s Triple Debt Model is a clean way to think about software health. The three models of debt are technical, cognitive, and intent.
Technical debt lives in the code. It’s the accumulation of implementation choices that make the system harder to change later: the tangled module, the shortcut you took under deadline, the abstraction that leaked. We’ve understood this one for decades. You feel it coming through slow builds, fragile tests, and the dread of touching one particular file.
Cognitive debt lives in people. It’s the erosion of shared understanding, the gap between how much code exists and how much any human understands. I’ve been calling this comprehension debt. It builds up when the system grows faster than the team’s mental model of it. Your code can be pristine and you can still carry crippling cognitive debt, because nobody understands the pristine code either.
Intent debt lives in artifacts. It’s the absence or erosion of the externalized rationale, goals, and constraints that explain why the system is the way it is. The key word is externalized. The rationale has to be written down where a teammate, a future you, or an agent can read it, not held in your head. When intent debt runs high, the system drifts from what you meant it to do, and nobody can say when it diverged or why.
These three are independent, which took me a while to internalize.
You can have low technical debt and high intent debt. You can understand a system completely yourself (no cognitive debt for you) while its intent exists nowhere outside your skull (enormous intent debt for everyone else).
From the inside they feel alike, but each one bills you separately.
Why intent debt is the one agents can’t help with
AI generates code faster than ever, which makes technical debt cheaper to take on and cheaper to pay down. Point an agent at a tangled module and it’ll refactor it.
Cognitive debt recovers too, more easily than most engineers expect. When you don’t understand a chunk of the system, you ask the agent to explain it. You rebuild part of the lost mental model on demand, because the code still exists and the model can read it back to you.
Intent is different. An agent can’t generate intent, because intent is the one input that has to come from you. A model can infer a plausible rationale from the code, the same way you can guess why a previous engineer did something. A guess about intent isn’t the intent. The model doesn’t know whether that 300ms debounce was a deliberate UX decision, a benchmark result, or a number someone typed once and never revisited. It will invent a confident-sounding reason, which is worse than admitting it doesn’t know.
Of the three debts, intent debt is the only one where the agent can’t bail you out. It can write the code and restore your comprehension. The why is the one thing it can only fabricate.
Agents make the un-written cost compound much faster
Teams got away with high intent debt for years because we carried it in our head and old docs.
When a new human joined a team, you didn’t write everything down, because they picked up intent over time: hallway conversations, code review comments, “oh, we don’t do it that way because of an incident in 2023.” Knowledge moved person to person and built up. The engineer who’d been there four years was the intent documentation, expensive and lossy, but it worked.
Agents break that model. Bringing agents onto a team doubles its size overnight with junior people who have no long-term memory. An agent starts most sessions cold. It carries none of the tacit intent your humans built up over years. Whatever you haven’t externalized into an artifact it can read, it doesn’t have.
That changes the economics of not writing things down . Un-externalized intent used to cost you once in a while, at onboarding or after someone left. Now you pay it every session, multiplied by every agent you run.
Picture the 20 agents you’re so excited to parallelize. Each one is a teammate who has never met you, can’t read your mind, and will fill any gap in your intent with a plausible guess. The orchestration tax I wrote about is partly an intent-debt tax. Much of what makes managing many agents exhausting is re-supplying the intent you never wrote down.
The other half of the comprehension debt argument
When I wrote about comprehension debt , I made a point I want to revisit, because intent debt sharpens it.
I argued that detailed specs aren’t a complete answer. Translating a spec into working code involves a huge number of implicit decisions no spec ever captures, and a spec detailed enough to be the program is the program in a slower language. I still believe that.
Intent debt is the complementary truth.
Being unable to capture all intent is no license to capture none of it. The implicit decisions an agent now makes on your behalf, the ones a spec will never enumerate, are the decisions whose rationale evaporates if you don’t record at least the load-bearing ones. You can’t write down everything.
You do have to write down the why behind the choices that would be expensive to get wrong, because nobody will reconstruct those later.
Comprehension debt warns you not to trust that code is correct because it exists.
Intent debt warns you not to trust that the reason survives because the code does. Code is the answer; the intent was the question it was meant to solve. AI is brilliant at producing answers to questions you forgot to write down.
What high intent debt looks like
Intent debt rarely shows up as friction. It shows up as a particular kind of helplessness.
An agent “fixes” a bug by deleting a guard clause, and nobody can say whether that guard was load-bearing or leftover, because no doc or commit message ever recorded why it was there.
A refactor changes a behavior users depend on. The review passed because the diff looked clean and the tests were green, but the tests only encoded the previous behavior, never the intent.
You ask why two services talk over a queue instead of a direct call, and the honest answer is “an agent suggested it and it seemed fine.” That answer is intent debt, already accruing interest.
If you’ve felt the cognitive surrender version of this, defending a design choice you can’t reconstruct, intent debt is the team-scale, written-down version of the same hole.
Surrender is about your own posture in the moment. Intent debt is what a hundred of those moments leave in the repo for the next person and the next agent to inherit.
Paying it down: externalize intent as a first-class artifact
Almost everything I’ve been writing about for the last few months turns out to be intent-debt management. I didn’t have the word for it. The move is the same each time: take the intent out of your head and put it somewhere an agent can read.
Write the spec for the intent, not the implementation. A good spec captures the goals, the constraints, the non-negotiables, and an explicit definition of done (fast, accessible, secure, delightful, beyond “functionally correct”). The spec carries the intent the code can’t carry on its own.
Treat AGENTS.md as your intent ledger, not your config. It’s why I keep saying stop using /init . An auto-generated file describes what the code is. An intent file describes what the team means: the conventions, the “we don’t do it this way because,” the constraints invisible in any single file. Agents can’t infer that, and they need it most.
Capture decisions where they happen. Lightweight decision logs (ADRs) are pure intent-debt paydown. Recording why at the moment you decide costs almost nothing. Reconstructing it eight months later, after the person who knew has moved teams, costs a fortune. Agents have made logging cheaper than ever, so the old excuse is gone.
Make the learning loop write intent back down. I’ve argued for self-improving agents that update a learnings file at the end of a session. The same loop is an intent-debt pump running in reverse: every mistake whose root cause you record, every “we tried X and it didn’t work because Y” is intent that would otherwise have lived only in your memory of a bad afternoon.
None of these are new tools. They’re the discipline of refusing to let the why exist only in your head, in an era where your head is no longer where most of the work happens.
For a long time, the scarce, valuable thing in software was the ability to produce a correct implementation. Code was expensive, so we optimized for writing it.
AI made code cheap, and comprehension is recoverable. Intent, the goals and constraints and reasons, is the one input that still has to originate with a human. It’s also the one we’re worst at externalizing, because for decades we got away with carrying it in our heads.
That worked when the team was a handful of people who could absorb intent over years of shared context. It does not work when half the team is agents that start every session as strangers.
Technical debt makes your system hard to change. Cognitive debt makes it hard to understand. Intent debt makes it hard to know whether the system still does what you wanted, and it’s the only one of the three your agents can’t pay back for you. That part stays with you. Write down the why, because it’s becoming the most valuable thing you can leave in the repo.
Want more? Subscribe to my free newsletter:
Disclaimer: The views and opinions expressed on this site are my own and do not necessarily reflect the views, positions, or strategies of Google or any of its affiliates.
