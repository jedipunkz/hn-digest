---
source: "https://www.intercom.com/blog/ai-is-approving-our-pull-requests-heres-how-we-made-it-safe/"
hn_url: "https://news.ycombinator.com/item?id=49254116"
title: "AI is approving our pull requests"
article_title: "AI is approving our pull requests: Here's how we made it safe - The Intercom Blog"
author: "LexSiga"
captured_at: "2026-08-11T06:48:16Z"
capture_tool: "hn-digest"
hn_id: 49254116
score: 1
comments: 0
posted_at: "2026-08-11T06:30:27Z"
tags:
  - hacker-news
  - translated
---

# AI is approving our pull requests

- HN: [49254116](https://news.ycombinator.com/item?id=49254116)
- Source: [www.intercom.com](https://www.intercom.com/blog/ai-is-approving-our-pull-requests-heres-how-we-made-it-safe/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T06:30:27Z

## Translation

タイトル: AI がプル リクエストを承認しています
記事のタイトル: AI がプル リクエストを承認: 安全にする方法は次のとおりです - The Intercom Blog
説明: Intercom では、これまで以上に多くのコードを作成しています。 PR の承認に AI を安全に使用する方法を紹介します。

記事本文:
メインコンテンツにスキップ
インターコムのブログ
製品
顧客向け チャット、電子メール、音声、ソーシャル機能を備えた世界最高のビジネス メッセンジャーを使用して、顧客がいる場所で顧客に会えます。
サポート エージェント向け AI で強化された受信トレイでチームを強化します。
サポート リーダー向け チーム、AI エージェント、サポート エクスペリエンス全体を、美しく設計された 1 つのシステムで管理します。
リソース
アイデア ブログ 当社のリーダーシップ チームによる製品とデザインの考え
AI 研究ブログ Intercom AI グループからの洞察
The Ticket ポッドキャスト カスタマー サービスの最先端で未来を見据えたリーダーとの会話。
製品
顧客向け チャット、電子メール、音声、ソーシャル機能を備えた世界最高のビジネス メッセンジャーを使用して、顧客がいる場所で顧客に会えます。
サポート エージェント向け AI で強化された受信トレイでチームを強化します。
サポート リーダー向け チーム、AI エージェント、サポート エクスペリエンス全体を、美しく設計された 1 つのシステムで管理します。
リソース
アイデア ブログ 当社のリーダーシップ チームによる製品とデザインの考え
AI 研究ブログ Intercom AI グループからの洞察
The Ticket ポッドキャスト カスタマー サービスの最先端で未来を見据えたリーダーとの会話。
AI がプル リクエストを承認しています: 安全性を確保する方法は次のとおりです
Intercom では、これまで以上に多くのコードを作成しています。 PR の承認に AI を安全に使用する方法をご紹介します。
9分
Intercom では、配送こそが私たちの心です。私たちは 1 日に何百回もコードを本番環境にプッシュします。
エンジニア、エンジニアリング マネージャー、デザイナー、PM は全員、安全にこれに貢献します。コードをマージしてから本番環境で実行されるまでの平均時間は 12 分です。
私たちは、直観に反して聞こえるかもしれないが、スピードは安全の敵ではないという信念を長年抱いてきました。それはそのための前提条件です。コードを蓄積するとリスクが生じます。少量のバッチで出荷することで、それを最小限に抑えることができます。ファス

出荷後は、それぞれの変更が小さくなり、コンテキストがまだ頭の中に残っているため、問題を見つけやすくなり、問題が発生した場合にロールバックしやすくなります。
現在、2 つの主要なコードベースにわたるプル リクエスト (PR) の 93% 以上がエージェント主導型です。また、19% 以上は人間のレビュー担当者が関与することなく自動承認されています。
この投稿では、2 番目の数字と、それが私たちの安全を高めると考える理由について説明します。ほとんどの人は「AI がプルリクエストを承認している」と聞くと、それは無謀だと考えます。データは別のことを物語っていると私たちは考えています。
昨年、当社の CTO ダラー・カランは、12 か月以内に研究開発組織全体の生産性を 2 倍にするという明確な目標を設定しました。構築と出荷が早くなればなるほど、お客様が必要な機能をより早く入手できるようになるからです。
9か月後、私たちはそれを達成しました。結果は全体的に重要なものでしたが、この投稿にとって最も重要な数値は、デプロイメントが 2 倍になったにもかかわらず、重大なコード変更によるダウンタイムが 35% 減少したことです。発送が早くなったことで、より安全になりました。ソフトウェアの構築方法と出荷方法を最新化する中で、私たちは体系的にボトルネックを明らかにし、それに取り組みます。私たちが見つけた中で最大のものの 1 つでしょうか? PRレビュー。
人間には、現在生成されている AI 生成コードの量を適切にレビューする時間も精神力もありません。
AI エージェントが機能する実装を数分で作成できる場合、人間がレビューするのに数時間または数日待つのは、インピーダンスの不一致です。生産ラインは、品質ゲートが追いつかないほどの速さで動いています。
それが起こると、次の 2 つのうちの 1 つが起こります。キューが後退して速度が低下するか、より危険なことに、人間がゴム印を押し始めるかのどちらかです。差分をざっと見て、説明をざっと読んで、「承認」をクリックします。
一部の企業は、静かにこの失敗モードに陥っています。私たちはそれに正面から立ち向かうことを選択し、

厳密な解決策。
PR レビューは、適切に行われたとしても、非常に複雑なプロセスです。優れたレビュー担当者は、PR の説明に基づいて問題の記述の品質を評価します。彼らは、変更セットが実際に指定された意図と一致することを確認します。彼らは、ベスト プラクティスに照らしてコードをレビューし、論理的な問題を探し、個人的な製品コンテキストを適用して変更が意味があることを検証し、パフォーマンスの問題や安全性に関する懸念などをチェックします。現実的には、特に時間的制約がない限り、すべての PR についてこれらすべての側面を適切にカバーする経験を持った人間のレビュー担当者は一人もいません。そして、データを見ればわかるように、私たちが比較してきたベースライン、つまり人間によるレビューは、私たちのほとんどが想定していたよりも弱かったのです。
そこで私たちは、もっとうまくできるとしたらどうなるだろうかと自問しました。
当社の PR レビュー エージェントは、コード レビューを単一のタスクとして扱いません。それを個別のサブジョブに分解し、それぞれが独立したサブエージェントによって処理されます。問題の説明の質を評価します。もう 1 つは、diff が実際に指定された意図と一致しているかどうかをチェックします。安全性に関する懸念についての別のレビュー。もう 1 つは論理的な正しさをチェックします。もう 1 つは、ベスト プラクティスと既知のアンチパターンに対するレビューです。等々。
その結果、すべての PR は、あたかも当社の最も勤続年数があり知識が豊富な 12 人のエンジニア全員が、それぞれ独自の専門レンズを持ち込んで同時に検討しているかのようにレビューされます。以前は、単一の PR についてこれほど広範なレビューを得ることは不可能でした。今ではそれがデフォルトになっています。
人間のレビュー担当者は通常、実際のコード変更、つまり差分に焦点を当てます。私たちのエージェントはさらに深く掘り下げていきます。実行パスをトレースし、コードベース全体の変更の影響を追跡します。これは、人間がやりたくても、なかなか時間が取れなかったことです。
一連の過去の PR で新しい PR レビュー エージェントをテストしているときに、

1 行のテキスト コピーの変更が間違っているとフラグを立てていることがわかりました。表面的には、それはまったく無害で、単なるテキスト更新のように見えました。私たちはそれが間違いだと思っていましたが、そうではありませんでした。私たちのエージェントは、新しいコピーがコードベースの他の場所にある既存の検証メカニズムと矛盾していることを発見しました。人間のレビュー担当者は、たまたまその検証コードをごく最近に作成したのでない限り、これを現実的に発見することはできなかったでしょう。私たちのエージェントは常に実行をトレースしているため、この種の事象を毎回一貫して捕捉します。
レビューも一般的ではありません。これは、当社のエンジニアが構築し、改良を続けている Intercom 固有のガイダンスに基づいており、PR を自分でレビューする場合に適用するのと同じコンテキスト、標準、製品知識をエンコードしています。エージェントが PR をレビューすると、エンジニアはレビュー コメントが役に立ったかどうかにフラグを立て、そのフィードバックによってガイダンスが継続的に強化されます。これはフライホイールです。エンジニアがコードベースの考え方をシステムに教えることに投資すればするほど、その後のレビューはより良くなります。
自動承認が強制されることもありません。エンジニアはいつでも、変更に関する人間によるレビューを要求できます。システムはツールであり、義務ではありません。 Intercom では、配送コードはマージ時に終了しません。変更を出荷するエンジニアは、変更が実際に稼働するのを監視し、本番環境での動作を監視し、何か問題があればロールバックできるようにすることが期待されます。 AI による承認によってそれは変わりません。コードを配布した人間は結果に対して責任を負います。
AI が承認した PR についての素朴な見方は、それはただのゴム印の LLM 呼び出しであり、人間が手間をかける必要はない、というものです。便利な機能です。それは実際に起こっていることを見逃しています。
私たちのエージェントは厳格です。大規模な PR は承認されません。変更が大きすぎる、複雑すぎる、または範囲が広すぎる場合は、

スコープにフラグを立てて、分解する必要があります。エンジニアがより小規模で段階的な、適切な範囲の変更を出荷するための直接的で前向きなインセンティブを生み出します。
これは安全にとって非常に重要です。小さな変更はレビューしやすく、テストしやすく、理解しやすく、そして重要なことに、何か問題が発生した場合のロールバックも簡単です。これは当社の配送文化を常に支えてきた同じ原則ですが、現在では PR レビュー エージェントが積極的にこれを実施しています。
「AI によって承認された PR が 50% 以上」などの目標を見て、結果ではなく指標に合わせて最適化しているのではないかと心配したくなります。私たちはそれを違って見ています。私たちの目標は、チェックを怠った場合、エンジニアが時間のプレッシャーの下でゴム印を押してレビューするように仕向ける危険があるボトルネックを取り除くことでした。
表面的には、当社の PR レビュー エージェントは「人間には PR をレビューする時間がない」という問題に対する解決策にすぎません。しかし、その本質は、本質的には安全機構です。これにより、AI によって生成されたコードの量が増加しても、ハートビートを維持および増加させながら、少量ずつ迅速に出荷し続けることが保証されます。
私たちは AI レビューだけで十分だとは考えず、積極的に実験を行いました。
私たちの仮説は、実際に重要な結果、つまり変更が正しかったかどうかによって測定される、AI レビューが人間によるレビューの品質と同等、またはそれを上回る可能性があるというものでした。それらは生産時に問題を引き起こしましたか?どれくらい早く審査され、承認されましたか?
私たちは、AI 承認パイプラインを通じて 100 人を超える PR の制御されたパイロットから開始しました。その結果、AI によって承認された PR の取り消しはゼロになり、75 パーセンタイルでの承認までの時間が 6 ～ 16 倍改善されました。それ以来、システムは大幅に拡張されました。広範な展開の最初の 4 週間で、497 人の PR が完全に自律的に稼働し、クロードがコードを作成し、AI 承認システムがレビュー、承認、

そして生産への出荷。
承認パイプライン自体を超えて、人間が作成したコードと比較して、AI が作成したコードが本番環境でどのように動作するかについてもより広範囲に調査しました。 AI が作成したバックエンド コードの復帰率は 0.53% でしたが、人間が作成したものは 5.39% でした。フロントエンドでは、0.22% 対 2.00% でした。
AI が作成したコードは、自動化されたパイプラインを通じてレビューおよび承認され、人間が作成し人間が承認したコードの数分の 1 の割合で元に戻されています。これが永遠にゼロに維持されるとは考えていませんが、これまでのデータから、エージェントが保持している品質基準は少なくとも人間が保持していた品質基準と同じくらい高く、多くの場合それより高いことがわかります。
そしておそらく最も重要な視点の変化は、過去にサービス停止を引き起こした製品の変更でしょうか?それらはすべて人間によってレビューされ、承認されました。人間によるレビューは安全性を保証するものではありません。そんなことは決してなかった。これは便利なヒューリスティックですが、実際には限界があり、私たちは何十年も黙って受け入れてきました。
イノベーションを起こしながらコンプライアンスを維持する
この投稿で説明したすべて、サブエージェント アーキテクチャ、トレーサビリティ、ラベル付け、データは、いずれもシステムを高速化するためだけに構築されたものではありません。これは、システムを監査可能にするために構築されました。それは初日からの設計上の制約でした。
AI によって承認されたすべての PR にはラベルが付けられ、ログに記録され、クエリ可能です。レビューコメント、承認決定、テスト結果、マージイベントはすべて記録されます。監査人が確認することを期待する証拠は、変更を承認したのが人間であっても AI であっても同じです。 「誰」は変わるかもしれないが、「何を」は変わらない。
私たちは規模を拡大する前の早い段階で監査人のシェルマン氏と協力しました。私たちは彼らと積極的に協力して、自動化されたレビュープロセスと彼らが生成する証拠がSOを含むコンプライアンスフレームワークの要件を満たしていることを確認しました。

C 2、HIPAA、ISO 27001、ISO 42001、AIUC-1 など。私たちは、AI 主導の変更管理は人間主導のプロセスが設定する基準を満たし、それを超えることができると考えており、それを証明したいと考えています。
私たちはこれを、システムを正しい方法で構築するための機能であり、クリアしなければならない追加のハードルではないと考えています。安全性を重視して構築すると、コンプライアンスも遵守されます。
PR レビューを安全メカニズムとして行うには、人間や AI を問わず、どれほど優れたレビュー担当者であっても限界があります。本番でのみ、未知のことを発見できます。 Intercom の最大規模の機能停止のほとんどは、製品コードの変更によって引き起こされたものではありませんでした。それは、インフラストラクチャの問題、顧客の予期しない使用パターン、またはサードパーティの機能停止でした。 PR レビューは、人間であろうと AI であろうと、決してそれらを捕まえることはできません。そのため、私たちは並行して、本番環境の問題を積極的に診断するエージェントの開発にも取り組んでいます。これについては近々詳しく説明します。
Intercom では常にスピードが当社の構築の中核であり、安全性にもかかわらずではなく、安全性のためにスピードが重視されてきました。そして、AI の活用によりさらに高速化が進んでいます。 AI が承認した PR は品質と安全性の低下につながると容易に想定できますが、私たちのデータはそうではないことを証明しています。私たちの鼓動は強まるばかりです。
物事がうまくいかないときに正しいことをする
AI のスケールに合わせて顧客エクスペリエンスを測定する方法
会話デジ

[切り捨てられた]

## Original Extract

We're producing more code than ever at Intercom. Here's how we're safely using AI for PR approval.

Skip to main content
Intercom Blog
Product
For customers Meet your customers where they already are with the world’s best business messenger for chat, email, voice, social…
For support agents An AI-enhanced inbox to supercharge your team.
For support leaders Manage your team, your AI agents and your entire support experience in one beautifully designed system.
Resources
Ideas blog Product & Design thoughts from our leadership team
AI research blog Insights from the Intercom AI group
The Ticket podcast Conversations with future-focused leaders at the cutting edge of customer service.
Product
For customers Meet your customers where they already are with the world’s best business messenger for chat, email, voice, social…
For support agents An AI-enhanced inbox to supercharge your team.
For support leaders Manage your team, your AI agents and your entire support experience in one beautifully designed system.
Resources
Ideas blog Product & Design thoughts from our leadership team
AI research blog Insights from the Intercom AI group
The Ticket podcast Conversations with future-focused leaders at the cutting edge of customer service.
AI is approving our pull requests: Here’s how we made it safe
We’re producing more code than ever at Intercom. Here’s how we’re safely using AI for PR approval.
9 min
At Intercom, shipping is our heartbeat . We push code to production hundreds of times a day.
Engineers, engineering managers, designers, and PMs all contribute to this, safely. The average time from merging code to it running in production is 12 minutes .
We’ve long held a belief that might sound counterintuitive: speed is not the enemy of safety. It’s a prerequisite for it. Accumulating code creates risk. Shipping small batches minimizes it. The faster you ship, the smaller each change is, and the easier it is to catch problems, and roll back when something goes wrong as the context is still fresh in your head.
Today, over 93% of our pull requests (PRs) across our two main codebases are Agent-driven. And over 19% are auto-approved with no human reviewer in the loop.
This post is about that second number, and why we think it makes us safer. Most people hear “AI is approving our pull requests” and think that’s reckless. We think the data tells a different story.
Last year, our CTO Darragh Curran set an explicit goal: double the productivity of our entire R&D organization within 12 months . Because the faster we can build and ship, the faster our customers get the capabilities they need.
Nine months later, we did it . The results have been significant across the board, but the number that matters most for this post: downtime from breaking code changes dropped 35%, even as our deployments doubled. Shipping faster made us safer. As we modernize how we build and ship software, we systematically surface bottlenecks and tackle them. One of the biggest we’ve found? PR review.
Humans simply don’t have the time or mental capacity to properly review the volume of AI-generated code we’re now producing.
When an AI Agent can produce a working implementation in minutes, waiting hours or days for a human to review it is an impedance mismatch. The production line is moving faster than the quality gate can keep up.
When that happens, one of two things follows: either the queue backs up and velocity drops, or, more dangerously, humans start rubber-stamping. Glancing at a diff, skimming the description, clicking approve.
Some companies are drifting into this failure mode silently. We chose to confront it head-on and built a rigorous solution.
PR review, done properly, is a genuinely complicated process. A good reviewer assesses the quality of the problem statement based on the PR description. They confirm the changeset actually matches the stated intent. They review the code against best practices, look for logical issues, apply their personal product context to validate the changes make sense, and check for performance issues, safety concerns, and more. No single human reviewer realistically has experience to properly cover every one of these dimensions on every PR, especially not under time pressure. And as we’ll see in the data, the baseline we’ve been comparing against, human review, was weaker than most of us assumed.
So we asked ourselves: what if we could do better?
Our PR review Agent doesn’t treat code review as a single task. It decomposes it into separate sub-jobs, each handled by an independent sub-Agent. One assesses the quality of the problem description. Another checks whether the diff actually aligns with the stated intent. Another reviews for safety concerns. Another checks for logical correctness. Another reviews against best practices and known anti-patterns. And so on.
The result is that every PR is reviewed as if a dozen of our most tenured and knowledgeable engineers were all looking at it simultaneously, each bringing their own specialist lens. In the past, getting that breadth of review on a single PR was impossible. Now it’s the default.
A human reviewer typically focuses on the actual code changes, the diff. Our Agent goes deeper. It traces execution paths, following the implications of a change through the codebase. This is something humans rarely had time to do, even when they wanted to.
While testing our new PR review Agent on a set of historical PRs, we found it flagging a one-line text copy change as incorrect. On the surface, it looked completely harmless, just a text update. We assumed it was a mistake, but it wasn’t. Our Agent caught that the new copy contradicted an existing validation mechanism elsewhere in the codebase. No human reviewer would have realistically found this unless they happened to have written that validation code very recently. Our Agent catches this kind of thing consistently, every time, because it’s always tracing execution.
The review isn’t generic either. It’s grounded in Intercom-specific guidance that our engineers have built and continue to refine, encoding the same context, standards, and product knowledge they’d apply if they were reviewing the PR themselves. When the Agent reviews a PR, engineers flag whether the review comments were helpful or not, and that feedback continuously sharpens the guidance. It’s a flywheel: the more our engineers invest in teaching the system how to think about our codebase, the better every subsequent review gets.
Automated approval is also never forced. Any engineer can request a human review on any change, at any time. The system is a tool, not a mandate. At Intercom, shipping code doesn’t end at merge. The engineer who ships a change is expected to watch it go live, monitor its behaviour in production, and be ready to roll back if something isn’t right. AI approval doesn’t change that. The human who ships the code remains accountable for the outcome.
The naive take on AI-approved PRs is that it’s just a rubber-stamp LLM call so that humans don’t have to bother. A convenience feature. That misses what’s actually happening.
Our Agent is strict. It won’t approve large PRs. If a change is too big, too complex, or too broad in scope, it flags it and requires it to be broken down. Creating a direct, positive incentive for engineers to ship smaller, more incremental, well-scoped changes.
This matters enormously for safety. Small changes are easier to review, easier to test, easier to understand, and, critically, easier to roll back when something goes wrong. This is the same principle that has always underpinned our shipping culture , but now the PR review Agent actively enforces it.
It’s tempting to look at a goal like “>50% AI-approved PRs” and worry that we’re optimizing for the metric rather than the outcome. We see it differently. Our goal was to remove a bottleneck that, if left unchecked, risked pushing engineers towards rubber-stamping reviews under time pressure.
On the surface, our PR review Agent is only a solution to “humans don’t have time to review PRs.” But what it really is, at its core, is a safety mechanism. It ensures we continue shipping fast, in small increments, maintaining and increasing our heartbeat, even as the volume of AI-generated code grows.
We didn’t assume AI review would be good enough, we actively ran an experiment.
Our hypothesis was that AI review could match or outperform human review quality, measured by the outcomes that actually matter: were the changes correct? Did they cause problems in production? How quickly were they reviewed and approved?
We started with a controlled pilot of over 100 PRs through the AI approval pipeline. The results: zero reverts of AI-approved PRs, and a 6–16x improvement in time-to-approval at the 75th percentile. Since then, the system has scaled significantly. In the first four weeks of broader rollout, 497 PRs went fully autonomous, with Claude writing the code and our AI approval system reviewing, approving, and shipping to production.
Beyond the approval pipeline itself, we also looked more broadly at how AI-authored code performs in production compared to human-authored code. AI-authored backend code had a revert rate of 0.53%, compared to 5.39% for human-authored. On the frontend, it was 0.22% versus 2.00%.
AI-authored code, reviewed and approved through our automated pipeline, is being reverted at a fraction of the rate of human-authored, human-approved code. We don’t expect that to hold at zero forever, but the data so far tells us that the quality bar our Agent holds is at least as high as the one humans were holding, and in many cases higher.
And perhaps the most important perspective shift: those product changes that did cause outages in the past? They were all reviewed and approved by humans. Human review is not a guarantee of safety. It never was. It’s a useful heuristic, but one with real limitations that we’ve been quietly accepting for decades.
Staying compliant while we innovate
Everything we’ve described in this post, the sub-Agent architecture, the traceability, the labelling, the data, none of it was built just to make the system fast. It was built to make the system auditable. That was a design constraint from day one.
Every AI-approved PR is labelled, logged, and queryable. The review comments, the approval decision, the test results, the merge event: all recorded. The evidence an auditor expects to see is the same whether a human or an AI approved the change. The “who” may change, but the “what” doesn’t.
We engaged our auditors, Schellman, early, before we scaled. We proactively worked with them to confirm that our automated review processes and the evidence they produce meet the requirements of our compliance frameworks , including SOC 2, HIPAA, ISO 27001, ISO 42001, and AIUC-1, among others. We think AI-driven change management can meet and exceed the standards that human-driven processes set, and we want to help prove that.
We see this as a feature of building the system the right way, not an extra hurdle we had to clear. When you build for safety, compliance follows.
You can only go so far with PR review as a safety mechanism, no matter how good the reviewer is, human or AI. Only in production do you discover the unknown unknowns. The majority of Intercom’s largest outages weren’t even caused by changes to product code at all. They were infrastructure issues, unanticipated customer usage patterns, or third-party outages. PR review, whether human or AI, was never going to catch those. That’s why, in parallel, we’re also working on an Agent that proactively diagnoses issues in production. We’ll share more on this soon.
Speed has always been at the core of how we build at Intercom, not in spite of safety, but because of it. And we’re getting even faster with AI. It’s easy to assume that AI-approved PRs would lead to a drop in quality and safety but our data proves otherwise. Our heartbeat is just getting stronger.
Doing the right thing when things go wrong
How to measure the customer experience as AI scales
Conversation desi

[truncated]
