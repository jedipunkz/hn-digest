---
source: "https://www.ardel.io/blog/the-3am-problem"
hn_url: "https://news.ycombinator.com/item?id=48437438"
title: "The AI coding optimized for the part of engineering that hurt the least"
article_title: "Your coding agent goes quiet at 3am. That's the moment that matters. | Ardelio"
author: "srbsa"
captured_at: "2026-06-07T21:32:31Z"
capture_tool: "hn-digest"
hn_id: 48437438
score: 2
comments: 0
posted_at: "2026-06-07T18:36:18Z"
tags:
  - hacker-news
  - translated
---

# The AI coding optimized for the part of engineering that hurt the least

- HN: [48437438](https://news.ycombinator.com/item?id=48437438)
- Source: [www.ardel.io](https://www.ardel.io/blog/the-3am-problem)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T18:36:18Z

## Translation

タイトル: エンジニアリングの中で最も害が少ない部分に最適化された AI コーディング
記事のタイトル: コーディング エージェントは午前 3 時に静かになります。それが重要な瞬間です。 |アルデリオ
説明: 業界が注目しているすべてのツールは、新しいコードを作成するときに現れます。 Sev 1 の午前 3 時に現れる人は誰もいません。このギャップが、エンジニアリング組織で最もコストのかかる問題になりつつあります。

記事本文:
コーディング エージェントは午前 3 時に静かになります。それが重要な瞬間です。 | Ardelio Ardelio 早期アクセスを取得 → ← すべての投稿 June 1, 2026 · 14 min read コーディング エージェントは午前 3 時に静かになります。それが重要な瞬間です。
それらは楽しい部分のために作られました。上級エンジニアが実際に生活している運用上の現実を考慮して構築された人は誰もいません。
構築中のほとんどのツールは、新しいコードを作成しているときに表示されます。午前 3 時に生産を保存しようとすると、それらはどれも表示されません。このギャップは偶然ではなく、エンジニアリング組織内で最も高価なものになりつつあります。
午前 3 時 14 分、ダニエル (名前は変更されました) の電話が何か異常な動作をしています。 Slack でのメンションの礼儀正しいバズ音の代わりに、一度認識されなくなったページに対する絶え間ない警告音です。彼は毎月週に一度の当番なので、今週は先輩です。
チェックアウトの待ち時間が大幅に増加し、エラー率が上昇しています。彼は起きており、ラップトップを開いて、重要な時計、つまり注文の返金や副社長との月曜日の朝の会話がすでに始まっています。
次の90分間が実際にどのようになるかは次のとおりです。サニタイズされたレトロバージョンではなく、本物です。
彼は誰もがするところから始めます。「何が変わったのですか？」この質問は簡単そうに見えますが、簡単ではありません。答えは相互に通信しない 6 つのツールに分散されているからです。彼は展開ダッシュボードをチェックします。過去数時間に 3 つのサービスが出荷されました。彼はチームの Slack を 200 メッセージまでスクロールして、誰かが支払い経路に触れていないかどうかを探しました。彼は会社の内部検索を起動し (良い検索にはお金を払っています)、「チェックアウト再試行タイムアウト」と入力すると、2023 年の Runbook、2 つの設計ドキュメント、および中途半端な Confluence ページという大量のドキュメントが返されます。検索は技術的には機能しましたが、実際には機能しました

質問に答えません。彼は今でも、それぞれを開いてざっと目を通し、最新のものかどうかを判断し、次のページを開くときに走り書きします。彼は午前 3 時に約 3 時間の睡眠をとりながら手作業で合成を行っています。
30分経ってもまだ何も修正できていない。彼は独自のシステムの図を組み立てています。
彼は遅いわけではありません。これがただの仕事です。インシデント対応の分析では、エンジニアが調査を開始する前に散在する情報を収集するのに 15 ～ 25 分を費やし、複数のツールにわたるコンテキストの再構築に 40% ～ 60% の時間が失われていることが明らかになりました。一般的な停止には、Slack で 5 分間、最近のコミットを 10 分間確認し、ダッシュボードを 5 分間確認してから、25 分後にようやく特定のデプロイメントがスパイクの原因であることを理解する必要があります。技術的な修正は、一度何をすべきかを知ってしまえば、多くの場合簡単です。何をすべきかがわかる時点で夜は終わります。
最終的に、ダニエルは時系列の形状を通じてそれを絞り込みます。再試行が積み重なるような匂いがする、ノコギリ状に忍び寄るレイテンシーです。今、彼はそれを確認するためのデータを必要としていますが、これは単に別のツールであるだけでなく、まったく別の作業モードです。彼は現在、メトリクス バックエンドに対してアドホック クエリを作成し、時間枠をマッサージし、必要な数値を取得する使い捨てスクリプトを構築しています。誰も事前に作成したダッシュボードが存在しないためです。ターミナル、ブラウザ、メトリクス ツール、ターミナルに戻ります。ツールを切り替えるたびに流れが中断され、方向転換を余儀なくされます。
彼はそれを見つけます。誰かが午後のデプロイの 1 つで再試行の上限を引き上げました。キャップがその場所に置かれたのには理由があり、常に存在します。しかし、それは 8 か月前のコードレビューのコメントと、彼女にとって幸いにも眠っているスタッフ エンジニアの頭の中に生きています。ダニエルは電話をかけ、修正を発送し、

グラフを正規化し、オールクリアをチャネルにポストします。
次に、彼が最も嫌う部分が来ます。
彼はそれをすべて書き留めなければなりません。次にこれに遭遇した人が一晩中やり直す必要がないように、事後分析 (タイムライン、根本原因、彼が下した決定、およびフォローアップ)。午前5時に、彼は疲れ果てているため、内容の薄い、喜びのないバージョンを書くつもりであることも知っています。そして、今から3か月後、誰かが隣接する何かにぶつかって、自分の午前3時をゼロから書き始めなければならないでしょう。それにしても、彼らはどうやってこれを見つけるのでしょうか？
ここが注目すべき部分です。
その一晩中、彼の会社が投資した AI ツールは彼に何も提供しませんでした。
サービスを数秒でスキャフォールディングできるコーディング エージェントでしょうか?静けさ。 PRD の草案を作成し、会議を要約する同僚のエージェントですか?どこにもない。
「AI がソフトウェア開発を変革する」という約束は、ダニエルの仕事が最も困難で、賭け金が最も高く、会社が積極的に損失を出していた 90 分間には存在しませんでした。エージェントは、ハッピー パス コードを書くという楽しい部分に存在し、上級エンジニアが人生のほとんどを費やす運用の現実には不在です。
それはおそらく実際よりも気になるはずです。
これはコーディングの問題ではありませんでした
ダニエルの夜をもう一度思い出して、各段階で「実際に何が足りなかったのか?」と尋ねてみましょう。
それはコードを書く能力ではありませんでした。彼は文字通り、寝ている間にコードを書くことができます。常に欠けていたのは、組織のどこかに存在していても、必要なときに必要な場所に存在しない知識でした。
この 4 時間で何が変わりましたか?これはデプロイ ログと PR に存在していました。再試行の上限は何に設定されていますか?またその理由は何ですか?レビュースレッドにそれがありました。この形を見たことがある人はいるだろうか

前に？それは誰も表面化できなかった過去の事件に存在した。修正は、彼がそれを確認できるほど十分なコンテキストを再構築した瞬間に存在しました。
事件のあらゆる段階は、事件の服を着た知識の問題でした。検索ツールが悪かったために失敗したのではありません。読むべき文書の山を見つけることは、質問に答えることと同じではないため、失敗しました。午前 3 時に合成税が最も高価な通貨で支払われます。上級エンジニアはプレッシャーの下、時間との戦いで注意を払います。
これは、組織図には示されていない、上級エンジニアに関する静かな真実です。最も経験豊富な人材は、会社の生きた知識層として機能します。チームが「ダニエルに尋ねる」理由は、ダニエルが頭の中で、物事が実際にどのように機能するのか、そしてなぜそうなるのかについて、マッピングされ、調整され、最新の、情報源に基づいた理解を持っているからです。これは、会議で決定され、レビューで議論され、前回の停止中に苦労して学んだ内容であり、マシンや新人が到達できる場所には決して書き留められなかった内容です。彼は知識のベースです。それが素晴らしいのは、彼が眠っている間、または休暇中、または会社を辞めて唯一のコピーを持ち帰るまでだけです。
これは小さな問題や特殊な問題ではなく、エンジニアリングの一日の平均値にすぎません。複数の調査は、開発者が実際にアプリケーション コードの作成に費やすのは時間の最大 16% のみで、残りは監視、メンテナンス、調整、コンテキストの終わりのない探索などの運用および支援作業に費やされるという発見に集約されています。彼らの約 64% は、検索するだけで 1 日に 30 分以上費やしており、3 分の 1 は 1 時間以上費やしていると報告しています。本当のボトルネックは決してコードを書くことではなく、それを知ることでした。
AI は 16% で劇的に向上しましたが、残りの 84% はほとんど変化しませんでした
ここが、「AI が実現するだろう」というロードマップに賭けている人にとっては複雑になるところです。

ix エンジニアリング速度」。
コーディング ハーネスは、時折のスタブや幻覚 API を除けば、非常に速く非常に優れたものになりました。対照試験では、よく知っているコードベースで作業している経験豊富な開発者は、AI ツールを使用すると明らかに遅くなり、より明確に言えば、自分たちの方が速かったと信じていました。研究を運営している人々も開発者自身も、スピードアップを予測していました。現実は逆で、参加者はそれを実感できませんでした。
この発見は、2025 年初頭のツールを使用した成熟したコードベースに関する小規模な調査であることに注意しながら、正直に受け入れる必要があります。モデル、コーディングハーネス、ワークフローが成熟するにつれて、そのギャップはおそらく縮まっています。重要なのは、認識のギャップです。エンジニアは、実際にそうであるかどうかに関係なく、より速く感じるでしょう。つまり、その感覚はエンジニアリング組織を運営できる指標ではありません。
コーディングから一歩戻ると、より大きなパターンが見えてきます。エージェント支援による作業が失敗する理由についての企業分析は、同じ原因にたどり着き続けます。つまり、生のモデルの機能ではなく、コンテキストと計画が欠落しているということです。エージェントは、制約、組織がどのように機能するかについての知識、または明示的に文書化されていない、または指示されていないことを認識していませんでした。能力曲線は垂直に成長しましたが、チーム固有の運用知識の曲線はまったく動きませんでした。再試行の理論的根拠、一緒に実行する必要がある共同デプロイ、運用上の注意点などは、今でもダニエルの頭の中にだけ残っています。人々の心の中に閉じ込められ、ツールに散在します。
したがって、利益は非常に特定の場所、つまり誰かの頭の中に知識が存在する場所で止まります。それはまさに午前3時14分にダニエルが立っていた場所です。
エージェントを採用すればするほど事態は悪化する
本能的には、「だから、エージェントに運用業務も指示しよう」ということになります。良い本能。しかし、何が起こるかに注目してください

エージェントを追加する際の知識のギャップ。
すべてのエージェントは記憶力のない新入社員です。アーキテクチャのレビューに一度も加わったことはなく、支払いと在庫を一緒に展開する必要があることや、昨年の春からの「一時的な」レート リミッターが負荷を伴うようになったことも知りません。そのため、どのエージェントも毎回、新人エンジニアが尋ねるのと同じ質問を繰り返します。そして、その答えは今でも先輩たちの心の中に生き続けています。コンテキストに関するダニエルの負荷は軽減されていません。あなたはそれに、彼に文脈を尋ねる事柄の数を掛けて、それらすべてが同じ文書化されていないボトルネックにあることを示しています。
その間、根本的な衰退は続いています。部族の理解は、出発や再編成のたびに損なわれます。オンボーディングはまさにこの暗黙の知識の伝達であり、その伝達はエージェントを増員したり、より多く設立したりすることで拡大するものではないため、人間もエージェントも同様に、新人エンジニアはゆっくりと成長していきます。運用負荷は5年ぶりに3割増となり、負荷は増大する一方、それに対応するための知識は薄く広がっている。
エンジニアリング組織がエージェント化すればするほど、文書化されていない、調整されていない、個人に囚われた知識がすべてのボトルネックになります。より多くの機能を購入することはできますが、チームが書き留めなかったコンテキストを買い戻すことはできません。
優れたチームがそれに対して取り組み始めていること
成功を収めているチームは、実践的な知識を人々の心にたまたま蓄積される副産物ではなく、実際のインフラストラクチャ層として扱っています。以下にいくつかの動きを示します。
「何を」だけでなく「なぜ」を捉えます。マージされた PR には、変更内容が記録されます。再試行の上限が 2 である理由、および 2 に留まらなければならない理由は、通常、レビュー スレッドに蒸発し、その後人間の記憶に残ります。インシデントを引き起こす決定は、その論理的根拠が次のとおりであるものです。

次の人がどこを見るかは決して書かれていません。
作業が行われる場所に制約を設けます。インシデント中に誰も開かない Confluence ページに存在するルールは存在しない可能性があります。知識はアクションの瞬間、つまり差分が書かれているとき、変更がレビューされているとき、ページが起動されているときに現れる必要があり、参照するために覚えておく必要がある Wiki の中でではありません。
コンテキストを第一級の成果物として扱います。上級者が与えるのと同じ、最新の、情報源に基づいた、システム固有の調和の取れた答えは、チームが意図的に構築し、維持する必要があるものです。
インシデントからナレッジに戻るループを閉じます。死後のダニエルは午前5時に書くのが嫌いで、まさにこの層に栄養を与えるために存在しています。悲劇的なのは、タイムライン、デプロイ、意思決定、スレッドなど、キャプチャする必要のあるもののほとんどが、それが起こるのを監視していたシステムにすでに存在しているにもかかわらず、疲れ果てた人間が手動で支払う税金であるということです。
最後のポイントは伝えることです。オンコールが知識の問題である理由は、それが解決可能な理由でもあります。次の応答者が必要とするものはすべて、すでにどこかで作成され、記録されています。重要な時点で、現在の 1 つの答えに調整する必要があるだけです。
一つだけ変わったダニエルの夜を想像してみてください。
ページが起動します。ダニエルが開く前に

[切り捨てられた]

## Original Extract

Every tool the industry is excited about shows up when you're writing new code. None show up at 3am during a Sev 1. That gap is about to become the most expensive thing in your engineering org.

Your coding agent goes quiet at 3am. That's the moment that matters. | Ardelio Ardelio Get early access → ← All posts June 1, 2026 · 14 min read Your coding agent goes quiet at 3am. That's the moment that matters.
They were built for the joyful part. Nobody built for the operational reality senior engineers actually live in.
Most tools being built show up when you're writing new code. None of them show up when you're trying to save production at 3am. The gap is not an accident, and it is about to become the most expensive thing in your engineering org.
It's 3.14 am, and Daniel's (name changed) phone is doing something unusual; instead of a polite buzz of a Slack mention, it is the incessant alarm of a page that has already gone unacknowledged once. He's the senior on call this week, as he is a week every month.
Checkout latency is through the roof with the error rate rising. He's awake, laptop open, and the clock that matters — the one that turns into refunded orders and a Monday-morning conversation with the VP — has already started.
Here is what the next ninety minutes actually look like. Not the sanitised retro version, the real one.
He starts where everyone does: "what changed?" The question sounds simple but is not, as the answer is scattered across six tools that don't talk to each other. He checks the deployment dashboard; three services shipped in the last few hours. He scrolls the team Slack, two hundred messages deep, looking for whether anyone mentioned touching the payments path. He pulls up the company's internal search (they pay for a good one) and types "checkout retry timeout" , which returns a wall of documents — a runbook from 2023, two design docs, and a half-finished Confluence page. The search technically worked, but it did not answer the question. He still opens each one and glances through, decides if it is current, and scribbles it as he opens the next. He is manually synthesising by hand at 3 am with about three hours of sleep.
Thirty minutes in, he is yet to fix anything. He is assembling a picture of his own system.
He's not slow; this is just the job. Analysis of incident response reveals that 40% to 60% of the time is lost reconstructing context across tools , with engineers burning 15 to 25 minutes gathering scattered information before investigation begins. A typical outage involves 5 minutes on Slack, 10 minutes checking recent commits, 5 minutes reviewing dashboards, and only then — 25 minutes in — finally understanding a specific deployment caused the spike. The technical fix, once you know what to do, is often trivial. Getting to the point where you know what to do is where the night goes.
Eventually Daniel narrows it through time-series shape — latency creeping in a sawtooth that smells like retries stacking up. Now he needs the data to confirm it, which is not just a different tool but a different mode of work entirely: he is now writing ad hoc queries against the metrics backend, massaging time windows, and building the throwaway script that pulls the numbers he needs, because the dashboard that nobody built in advance doesn't exist. Terminal, browser, metrics tool, back to terminal. Every tool switch breaks the flow and forces him to reorient.
He finds it. Someone raised the retry cap in one of those afternoon deploys. There was a reason the cap was where it was — there always is — but it lives in a code-review comment from eight months ago and in the head of the staff engineer who is, mercifully for her, asleep. Daniel makes the call, ships the fix, watches the graph normalise, and posts the all-clear in the channel.
Then comes the part he hates the most.
He has to write it all down. The postmortem — the timeline, root cause, the decisions he made, and the follow-ups — so that the next person who hits this should not have to re-run his entire night. At 5 am, he also knows that he is going to write a thin, joyless version of it because he is exhausted, and three months from now someone will hit something adjacent and will have to start their own 3 am from scratch. How are they going to find this anyway?
Here's the part worth sitting with.
At no point during that entire night did any of the AI tooling his company had invested in have anything to offer him.
The coding agents that can scaffold a service in seconds? Silent. The coworker agents drafting PRDs and summarising meetings? Nowhere.
The entire promise of "AI is transforming software development" was absent during the ninety minutes when Daniel's job was the hardest, the stakes were the highest, and the company was actively losing money. The agents are present for the joyful part of writing happy-path code and are absent for the operational reality that senior engineers spend most of their lives in.
That should bother you more than it probably does.
This was never a coding problem
Let's replay Daniel's night and ask at each step, "What was actually missing?"
It wasn't the ability to write code. He can write code in his sleep, nearly literally. What was missing, every single time, was knowledge that existed somewhere in the organisation but not where he needed it, when he needed it.
What changed in the last four hours? That existed in the deploy logs and the PRs. What the retry cap was set to, and why ? That existed in a review thread. Whether anyone had seen this shape before? That existed in a past incident that no one could surface. The fix existed the moment he reconstructed enough context to see it.
Every step of the incident was a knowledge problem wearing incident's clothes. The search tool didn't fail because it was bad; it failed because finding a pile of documents to read is not the same as answering a question. At 3 am the synthesis tax is paid in the most expensive currency: a senior engineer's attention, under pressure, against the clock.
This is the quiet truth about senior engineers that no org chart shows: the most experienced people function as the company's living knowledge layer. The reason the team "asks Daniel" is that Daniel has, in his head, the mapped, reconciled, current, sourced understanding of how things actually work and why. This is the stuff that was decided in a meeting, argued out in a review, learnt the hard way during a previous outage, and never written down anywhere a machine or a newcomer could reach. He is the knowledge base. Which is only wonderful right up until he's asleep, or on vacation, or has left the company and taken the only copy with him.
This is not a small or exotic problem — just a median engineering day. Multiple studies converge on the finding that developers only spend ~16% of their time actually writing application code , with the rest going to operational and supportive work: monitoring, maintenance, coordination, and the endless hunt for context. Around 64% of them report spending more than 30 minutes a day just searching for stuff, and a third spend over an hour. The real bottleneck was never writing code — it was to know.
AI got dramatically better at the 16% while the other 84% barely moved
This is where it gets complicated for anyone betting their roadmap on "AI will fix engineering velocity".
The coding harnesses have gotten really good very fast, aside from the occasional stub and hallucinated API. In a controlled trial, experienced developers working in codebases they knew well were measurably slower when using AI tools — and, more tellingly, believed that they had been faster. The people running the study and developers themselves had both predicted a speedup. The reality was contrary, and the participants just couldn't feel it.
You should hold that finding honestly, with caveats — a small study on matured codebases with early-2025 tooling. The gap has likely narrowed as models, coding harnesses, and workflows have matured. The point is the perception gap: engineers will feel faster whether or not they really are, which means the feeling is not a metric you can run an engineering org on.
Stepping back from coding, the larger pattern shows up. Enterprise analysis of why agent-assisted work fails keeps arriving at the same culprit: not raw model capability, but missing context and planning — the agent didn't know the constraint, the knowledge of how the org functions, or the thing that wasn't explicitly written down or prompted. The capability curve grew vertically, but the curve for team-specific operational knowledge didn't move at all. The retry rationale, the co-deploys that must be done together, the operational gotchas — they still live only in Daniel's mind. Trapped in the minds of people and scattered across tools.
The gains therefore stall at a very specific place: where knowledge lives in someone's head. Which is exactly where Daniel was standing at 3:14 am.
It gets worse as you adopt more agents
The instinct is: "so we'll point the agent at the operational work too." Good instinct. But notice what happens to the knowledge gap as you add agents.
Every agent is a brand-new hire with no memory. It has never sat in your architecture review, it doesn't know that payments and inventory must deploy together, or that the "temporary" rate limiter from last spring is now load-bearing. So every agent, every time, re-asks the same questions that new engineers ask — and the answers still live in the minds of your seniors. You haven't reduced the load on Daniel for context; you have multiplied it by the number of things asking him for context and pointed them all at the same undocumented bottleneck.
Meanwhile, the underlying decay continues. The tribal understanding erodes with every departure and every reorg. New engineers — humans and agents alike — ramp slowly because onboarding is the transfer of exactly this unwritten knowledge, and that transfer doesn't scale by hiring more or spinning up more agents. With operational load having gone up by 30% for the first time in five years, the load is going up while the knowledge to handle it gets thinner and more thinly spread.
The more agentic the engineering org becomes, the more unwritten, unreconciled, person-trapped knowledge becomes the bottleneck for everything. You can buy more capability, but you cannot buy back the context your team never wrote down.
What good teams are starting to do about it
The teams getting ahead are treating their working knowledge as a real infrastructure layer — not a byproduct that happens to accumulate in the minds of people. Here are a few moves:
Capture the why , not just the what . A merged PR records what changed. The reasoning of why the retry cap is and must stay at 2 usually evaporates into a review thread and then human memory. The decisions that cause incidents are the ones whose rationale was never written where the next person would look.
Put the constraints where the work happens. A rule that lives in a Confluence page nobody opens during an incident may well not exist. Knowledge has to surface at the moment of action — when the diff is being written, when the change is being reviewed, when the page is firing — not in a wiki that you'd need to remember to consult.
Treat context as the first-class artefact. The same reconciled answer that a senior would give — current, sourced, and specific to your system — is what teams need to build and maintain deliberately.
Close the loop from incidents back to knowledge. The postmortem Daniel hates writing at 5 am exists to feed exactly this layer. The tragedy is that it's manual tax paid by an exhausted human, when most of what it needs to capture — timeline, deploys, decisions, threads — already exists in the system that watched it happen.
The last point is the tell. The reason on-call is a knowledge problem is also the reason it is solvable — everything the next responder needs was already produced and recorded somewhere. It just needed to be reconciled into one current answer at the moment it mattered.
Picture Daniel's night with one thing changed.
The page fires. Before Daniel opens

[truncated]
