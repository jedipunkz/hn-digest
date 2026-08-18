---
source: "https://fathohm.dev/comprehension-debt"
hn_url: "https://news.ycombinator.com/item?id=49346757"
title: "Comprehension debt: what AI-written code costs"
article_title: "Comprehension debt: what AI-written code actually costs — Fathohm"
image: "https://fathohm.dev/comprehension-debt/opengraph-image?f9c276481f826ac1"
author: "ashrivastavaa"
captured_at: "2026-08-18T15:21:58Z"
capture_tool: "hn-digest"
hn_id: 49346757
score: 2
comments: 0
posted_at: "2026-08-18T15:08:00Z"
tags:
  - hacker-news
  - translated
---

# Comprehension debt: what AI-written code costs

- HN: [49346757](https://news.ycombinator.com/item?id=49346757)
- Source: [fathohm.dev](https://fathohm.dev/comprehension-debt)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T15:08:00Z

## Translation

タイトル: 理解の負債: AI が作成したコードにかかる費用
記事のタイトル: 理解の負債: AI が作成したコードの実際のコスト — Fathohm
説明: AI エージェントは、作成されるコードと理解されるコードを分離しました。理解負債とはギャップです。それが何なのか、テクノロジー負債がそれをカバーしない理由、それが目に見えない理由、そしてそれを正直に測定する方法。

記事本文:
理解の負債: AI が作成したコードの実際のコスト — Fathohm Fathohm ギャラリー ダーク インデックス その仕組み 理解の負債 サイン イン マップを入手 → ギャラリー ダーク インデックスの仕組み 理解の負債 サイン イン エッセイ
理解負債: AI が作成したコードの実際のコスト
コードベースに先月出荷されたモジュールがあります。それは動作します。テストがあります。審査に通りました。そして、午前 2 時に故障した場合、チームの誰もそれが何をするのか説明できません。
「誰がこれを理解できますか？」と尋ねてください。 AI ネイティブのコードベース内の特定のファイルについて質問しても、誰も答えられないことがますます多くなっています。エンジニアの能力が低下したからではなく、本番環境に入る途中でコードが頭の中を通過しなくなったからです。
70 年間、コードが書かれるということは、誰かがそれを理解していることを意味していました。この含意は非常に信頼できるものであったため、私たちはそれを仮定とは考えませんでした。コードを書くことは、問題を表現するのに十分正確に問題を理解する行為でした。コードがどんなに悪くても、ドキュメントがどれほど欠けていても、それが何をするのか、そしてなぜそうなるのかを知っていた人は少なくとも 1 人、つまり執筆時点では作者でした。チームの方向性をコードベースに保つために私たちが行っているすべてのプラクティスは、そのフロアに静かに寄りかかっています。レビューでは作成者が変更を擁護できると想定され、オンボーディングでは誰かがシステムについて説明できると想定され、デバッグでは同僚が質問できると想定されます。
AI エージェントはその暗示を打ち破りました。コードの作成とコードの理解は別個のイベントになり、そのうちの 1 つだけがスケーリングです。エージェントは、チームが 1 か月かけて書いていたものを午後 1 分で作成できますが、午後には 1 か月分の理解が伴うわけではありません。 「少なくとも作成者は知っている」というフロアはなくなりました。エージェントが作成したコードの場合、作成者はチームの一員ではありません。それは誰でもない。
コーデブとのギャップ

ase が行うことと、それを担当する人間が理解するものには名前が必要です。名前のないものは管理されないためです。それは 1 つあり、しばらくの間も 1 つありました。ジェイソン・ゴーマン氏は2025年9月にこれを理解負債と名付け、「チームが理解できるよりも早くコードを作成した場合に起こること」と名付け、アディ・オスマニ氏は2026年3月にこれをより幅広い聴衆に伝え、それを「システム内に存在するコードの量と人間が真に理解できるコードの量との間の増大するギャップ」と定義した。
私たちは両方の定義からまったく同じ 1 か所で逸脱しており、それがこのエッセイを続ける理由です。彼らが理解できるよりも早く。本当に理解しています。これらは心の状態についての主張であり、心は観察できるものではありません。チェックできない定義は楽器ではなくフレーズであり、そのフレーズは懐疑的なスタッフエンジニアとの最初の議論に耐えられません。したがって、以下のすべては、実際に確認できるバージョンに関するものです。誰かがファイルを理解しているかどうかではなく、誰かが最近そのファイルを作成、レビュー、または説明したかどうかです。それは意図的に弱い主張です。私たちが作品を展示できる唯一の作品でもあります。
なぜ「テクノロジー負債」ではカバーできないのか
明らかな反論は、蓄積されたコードベースの問題を表す言葉がすでにあるということです。しかし、ウォード・カニンガムが造語したように、技術的負債はコードの特性であり、成果物自体に組み込まれたショートカットであり、成果物自体に表示されます。差分で技術的負債を指摘できます。
理解負債はチームの財産です。同じファイルでも、あるチームでは負債がゼロでも、別のチームではまったくの盲点になる可能性があります。その場合、キャラクターは 1 つも異なりません。負債はファイルにあるのではなく、ファイルとその責任を負う人間との関係にあるからです。だからこそ、技術負債の戦略は当てはまりません:y

リファクタリングして理解の負債から抜け出すことはできません。完璧で、クリーンで、十分にテストされた、誰も理解していないモジュールであっても、依然として責任はあります。何も問題がないようなので、おそらくより悪いものです。
また、通常の関心の方向も逆転します。技術的負債は、機能する悪いコードを心配します。理解力の負債は、機能する良いコードについて心配しますが、機能しなくなる瞬間まで心配し、チームは、自分たちが持っていたと思っていた理解が誰にも獲得されていなかったことに気づきます。
最新のツールチェーンには理解を測るものはありません。私たちは、カバレッジ、複雑さ、速度、デプロイ頻度、インシデント数など、コードとプロセスのあらゆる特性を測定しますが、人間による把握の特性は測定しません。
私たちが最も近いのはコードレビューでしたが、レビューは決して測定ではなく、サンプリングイベントでした。マージ時に 1 人の人の理解度を 1 回だけチェックし、「1 人の人が 1 回承認した」から「チームがこれを理解している」と推測しました。その推定は常に寛大でした。 AI ネイティブのスループットでは、それは崩壊します。差分のサイズが 3 倍になり、頻度が 4 倍になると、レビューは深くならずに短くなります。エージェントが作成した 400 行の変更に対してコメントのない承認は、理解の証拠ではありません。それはスループットの証拠です。
一方、最も古いチームリスクヒューリスティックは、静かに新たなレベルに到達しました。バス係数 (誰もシステムを理解できなくなるまでに何人消えるか) は、誰かがそれを書いたため、以前は 1 未満に制限されていました。エージェントが作成したコードはその限界を突破します。誰かがそれを促したので、カウントはゼロではありません。しかし、ファイルをプロンプトして差分を読んだ人は、著者のようにバスに乗っているわけではありませんし、多くの場合、他に誰もバスに乗っていません。正直な単位は全体ではなく部分であることが判明します。これは不正確です

着心地もポイントです。人全体をカウントするだけのヒューリスティックでは、ほとんどの AI ネイティブ コードが実際に置かれている状態を確認することはできません。
金融負債と同様に、理解負債は負担がかからず、返済は過酷です。維持コストは目に見えません。コードは機能し、ダッシュボードは緑色で、速度は素晴らしく見えます。利息は特定の瞬間に支払われます。
事件。デバッグは、考えられる限り最悪の代償を払って理解することです。午前 2 時に、顧客が待たされているというプレッシャーの下で行われます。理解されていないコードのすべてのインシデントは、マージ時に取得されなかったものであることをチームが理解して、プレミアムを付けて買い戻すことになります。
次の変化。理解していないベースラインとの差分を有意義にレビューすることはできません。理解されていないコードは、コードに関わるすべてのレビューの質を低下させます。これが、理解の負債をさらに悪化させる方法です。盲点があると、その周囲のコードを安全に変更することが難しくなり、その変更がエージェントに委任され、盲点が深まります。
出発。誰かが辞めると、チームは常に理解を失います。そして今、彼らは最後の人間を失いました - そして退任インタビューは、一人だけが知っていたことを誰も理解していなかったものを捉えていません。
オンボーディング。新人エンジニアは、コードベースを理解している人からコードベースを学びます。 AI ネイティブのコードベースは、学習する人がいない状態に達する可能性があります。コーパスは人間のどのコードベースよりも速く成長しました。
これらはいずれも、AI が作成したコードに反対するものではありません。この影響力は本物であり、それを拒否したチームは拒否したチームに負けます。レバレッジには現在の手段では示されないコストがあり、何も示されないコストは管理されず、発見されると主張しています。
実際に理解度を測ることができるのでしょうか？
正当な反論: 理解は人間の心の状態であり、心の状態は git には現れません。
正しいので、やめてください。メアス

代わりに記録を記録します。これは、エンジニアリングにおけるすべての重要な指標がすでに行っていることです。観察可能なシグナルは本物です。人間が変更を実質的にレビューしたかどうか (単なる承認ではなくコメントの記録)、人間が意味のあるファイルに書き込んだのはどれだけ最近か、何人の異なる人間がそのファイルに実際に接触したかどうかです。誰もが知っていることを尋ねるものはありません。それらはすべて、すでに git 履歴に含まれています。
グラウンド トゥルースもチェック可能ですが、ゲーム全体に関わる 1 つの制約があります。チェックは自己管理できないということです。失敗したときにファイルがどのような動作をするかを誰かに説明してもらうのは、本当のテストです。自分自身の回答を採点するよう求めるのはアンケートであり、人々が関心を持っている数値に添付されたアンケートは、単に移動できるスライダーにすぎません。答えはそれを書いた人以外の人が読む必要があります。
信号の選択よりも重要なのは、信号に関する規律です。
決定論的。モデルの雰囲気から得られる理解度スコアは反証不可能であり、反証不可能な指標は無視されます。同じ履歴からは常に同じ数値が生成されるはずです。
分解可能。すべてのスコアは、それを生み出した要因を分析する必要があります。それ自体を説明できない数値は信頼されず、信頼されない信頼指標は無意味です。
議論の余地あり。方法論は公開され、ラベルは修正可能である必要があります。間違っているが検査可能なビートは測定されていません。
朽ちていく。理解が薄れていく。 1 年前にモジュールについて知っていたことは、今ではほとんどわかっていません。指標がそう示す必要があります。つまり、理解は蓄えられるものではなく、再び得られるものであることを意味します。
他の人によって検証されました。どのプロキシも最終的にはゲームにさらされることになります。レビューシアターからファームレビュー深度への移行は、スコアが重要な瞬間には明らかな動きです（グッドハートはそれを保証しています）。アンチゲーム メカニズムは、定期的に本物の人間にコードに関する実際の質問をしています。しかし、問いかけているのは、

半分だけです。受け取った人が得点した小切手は、このリストの中で最も簡単にゲームできるものであり、まったく努力を必要としない唯一のものです。検証は、答えを書いていないチームメイトがそれを読んだときにカウントされます。
それに対して何をすべきか（ツールの有無にかかわらず）
新しいツールを使わずに、明日から理解負債の管理を始めることができます。
エージェントが作成したコードには人間の作成者記録を与えます。マージの一環として「これについては説明できます」と受け入れ、それを受け入れていることを知っている人です。
エージェント PR の無言承認を禁止します。レビューに実質的なコメントがない場合、そのコードに対するチームの理解度はレビュー前とまったく同じで、ゼロになります。
大きな声で質問してください。次の計画会議では、「私たちのシステムのどの部分がここにいる誰も理解していませんか?」その後の沈黙は借金の音が聞こえるようになる。
名前付きの死角をバックログ項目として扱います。最悪の死角を含む 1 時間は、意図的に日中の価格で購入されたインシデント対応です。
手動でできないことは、表面全体を一度に確認したり、その動きを観察したり、劣化について正直に保つことです。スプレッドシートでテスト カバレッジを追跡する人がいないのと同じ理由です。
どちらの方法でもコードは書き続けられます。それが理解され続けるかどうかは、今や選択です。
それが私たちが構築した部分です。 Fathohm は、読み取り専用の GitHub アプリから、コードベース全体の理解負債を決定論的、分解可能、議論可能にマッピングします。

## Original Extract

AI agents decoupled code getting written from code getting understood. Comprehension debt is the gap — what it is, why tech debt doesn't cover it, why it's been invisible, and how to measure it honestly.

Comprehension debt: what AI-written code actually costs — Fathohm Fathohm Gallery The dark index How it works Comprehension debt Sign in Get your Map → Gallery The dark index How it works Comprehension debt Sign in Essay
Comprehension debt: what AI-written code actually costs
There's a module in your codebase that shipped last month. It works. It has tests. It passed review. And if it breaks at 2am, nobody on your team can explain what it does.
Ask “who understands this?” about any given file in an AI-native codebase and the honest answer, increasingly often, is no one — not because your engineers got worse, but because the code stopped passing through their heads on its way into production.
For seventy years, code getting written implied that somebody understood it. The implication was so reliable we never thought of it as an assumption: writing code was the act of understanding a problem precisely enough to express it. However bad the code, however absent the docs, there was at minimum one person — the author, at the moment of authorship — who knew what it did and why. Every practice we have for keeping teams oriented in a codebase quietly leans on that floor: review assumes the author can defend the change, onboarding assumes someone can explain the system, debugging assumes a colleague to ask.
AI agents broke the implication. Code getting written and code getting understood are now separate events, and only one of them is scaling. An agent can produce in an afternoon what a team used to write in a month — and the afternoon does not come with a month's worth of understanding attached. The floor of “at least the author knows” is gone: for agent-authored code, the author isn't on your team. It isn't anyone.
The gap between what a codebase does and what the humans responsible for it understand needs a name, because things without names don't get managed. It has one, and it has had one for a while. Jason Gorman named it comprehension debt in September 2025 — what happens “when teams produce code faster than they can understand it” — and Addy Osmani carried it to a much wider audience in March 2026, defining it as “the growing gap between how much code exists in your system and how much of it any human being genuinely understands.”
We depart from both definitions in exactly the same one place, and it is the reason this essay keeps going. Faster than they can understand it ; genuinely understands — those are claims about states of mind, and minds are not observable. A definition you cannot check is a phrase, not an instrument, and a phrase will not survive its first argument with a skeptical staff engineer. So everything below is about the version you can actually check: not whether anyone understands a file, but whether anyone has recently written, reviewed, or explained it. That is a weaker claim on purpose. It is also the only one we can show our work for.
Why “tech debt” doesn't cover it
The obvious objection is that we already have a word for accumulated codebase problems. But technical debt, as Ward Cunningham coined it, is a property of the code — shortcuts embodied in the artifact itself, visible in the artifact itself. You can point at tech debt in a diff.
Comprehension debt is a property of the team . The same file can be zero debt on one team and a total blind spot on another, with not one character different — because the debt isn't in the file, it's in the relationship between the file and the humans accountable for it. That's why the tech-debt playbook doesn't apply: you cannot refactor your way out of comprehension debt. A perfect, clean, well-tested module that nobody understands is still a liability — arguably a worse one, because nothing about it looks wrong.
It also inverts the usual direction of concern. Tech debt worries about bad code that works. Comprehension debt worries about good code that works — right up until the moment it doesn't, and the team discovers the understanding they assumed they had was never acquired by anyone.
Nothing in the modern toolchain measures understanding. We measure coverage, complexity, velocity, deploy frequency, incident counts — every property of the code and the process, and no property of the humans' grasp of it.
The closest thing we had was code review, and review was never a measurement — it was a sampling event. It checked comprehension exactly once, at merge time, in one person, and we extrapolated “the team understands this” from “one person approved it once.” That extrapolation was always generous. Under AI-native throughput it collapses: when the diffs triple in size and quadruple in frequency, reviews get shorter, not deeper. An approval with no comments on a four-hundred-line agent-written change is not evidence of understanding. It's evidence of throughput.
Meanwhile the oldest team-risk heuristic we have quietly hit a new floor. Bus factor — how many people can disappear before nobody understands a system — used to be bounded below by one, because someone wrote the thing. Agent-authored code breaks that floor. Somebody prompted it, so the count is not zero; but a person who prompted a file and read the diff is not on the bus the way an author is, and often nobody else is on it at all. The honest unit turns out to be fractional rather than whole — which is uncomfortable, and is the point. A heuristic that only counts whole people cannot see the state most AI-native code is actually in.
Like financial debt, comprehension debt is cheap to carry and brutal to service. The carrying cost is invisible: the code works, the dashboards are green, velocity looks great. The interest comes due at specific moments:
The incident. Debugging is comprehension paydown at the worst possible price — acquired under pressure, at 2am, with customers waiting. Every incident in un-understood code is the team buying back, at a premium, understanding it never acquired at merge time.
The next change. You cannot meaningfully review a diff against a baseline you don't understand. Un-understood code degrades the review of everything that touches it — which is how comprehension debt compounds: blind spots make the code around them harder to safely change, which gets delegated to the agent, which deepens the blind spot.
The departure. When someone leaves, teams have always lost understanding. Now they lose the last human who had any — and exit interviews don't capture what nobody realized only one person knew.
The onboarding. New engineers learn codebases from people who understand them. An AI-native codebase can reach the state where there is no one to learn from — the corpus grew faster than any human's model of it.
None of this argues against AI-written code. The leverage is real and teams that refuse it will lose to teams that don't. It argues that the leverage has a cost that no current instrument shows, and costs that nothing shows don't get managed — they get discovered.
Can you actually measure understanding?
The fair objection: understanding is a state of a human mind, and states of minds don't show up in git.
Correct — so don't. Measure the record instead, which is what every serious metric in engineering already does. The observable signals are real: whether a human substantively reviewed a change (a comment trail, not a bare approval), how recently a human meaningfully wrote in a file, how many distinct humans have had real contact with it. None of those ask what anyone knows. All of them are in your git history already.
The ground truth is checkable too, with one constraint that turns out to be the whole game: the check cannot be self-administered. Asking someone to explain what a file does on failure is a real test. Asking them to grade their own answer is a survey, and a survey attached to a number people care about is just a slider they can move. The answer has to be read by someone who did not write it.
What matters more than the choice of signals is the discipline around them:
Deterministic. A comprehension score that comes out of a model's vibes is unfalsifiable, and unfalsifiable metrics get ignored. The same history must always produce the same number.
Decomposable. Every score must break into the factors that produced it. A number that can't explain itself won't be trusted, and a trust metric that isn't trusted is nothing.
Disputable. The methodology should be public and the labels correctable. Wrong-but-inspectable beats unmeasured.
Decaying. Understanding fades. Whatever you knew about a module a year ago, you know less now — the metric has to say so, which means comprehension is re-earned, not banked.
Verified by someone else. Any proxy will eventually be gamed — review-theater to farm review-depth is the obvious move the moment a score matters (Goodhart guarantees it). The anti-gaming mechanism is periodically asking a real human a real question about the code. But the asking is only half of it: a check scored by the person taking it is the easiest thing on this list to game, and the only one where gaming it requires no effort at all. Verification counts when a teammate who did not write the answer reads it.
What to do about it (with or without tooling)
You can start managing comprehension debt tomorrow with no new tools:
Give agent-written code a human author-of-record — someone who accepts “I can explain this” as part of merging it, and knows they're accepting it.
Ban the silent approval for agent PRs. If the review has no substantive comment, the team's comprehension of that code is exactly what it was before the review: zero.
Ask the question out loud. In the next planning meeting: “which parts of our system does nobody here understand?” The silence after is the debt making itself audible.
Treat named blind spots as backlog items — an hour with the worst one, on purpose, is incident response bought at daytime prices.
What you can't do by hand is see the whole surface at once, watch it move, or keep yourself honest about decay — the same reason nobody tracks test coverage in a spreadsheet.
The code will keep getting written either way. Whether it keeps getting understood is now a choice.
That's the part we built. Fathohm maps comprehension debt across a codebase — deterministically, decomposably, disputably — from a read-only GitHub App.
