---
source: "https://buildforever.com/engineering/posts/an-ai-teammate-that-builds-itself"
hn_url: "https://news.ycombinator.com/item?id=48588584"
title: "An AI teammate that builds itself: how we wired Avery into our company's DNA"
article_title: "An AI teammate that builds itself: how we wired Avery into our company’s DNA | BuildForever Engineering"
author: "fishpham"
captured_at: "2026-06-18T18:55:19Z"
capture_tool: "hn-digest"
hn_id: 48588584
score: 6
comments: 1
posted_at: "2026-06-18T17:22:43Z"
tags:
  - hacker-news
  - translated
---

# An AI teammate that builds itself: how we wired Avery into our company's DNA

- HN: [48588584](https://news.ycombinator.com/item?id=48588584)
- Source: [buildforever.com](https://buildforever.com/engineering/posts/an-ai-teammate-that-builds-itself)
- Score: 6
- Comments: 1
- Posted: 2026-06-18T17:22:43Z

## Translation

タイトル: 自分自身を構築する AI チームメイト: Avery を当社の DNA にどのように組み込んだか
記事のタイトル: 自分自身を構築する AI チームメイト: Avery を当社の DNA に組み込んだ方法 |ビルドフォーエバーエンジニアリング
説明: Avery は、BuildForever のチームが毎日使用するシステムやツールに組み込まれている AI エンジニアリングのチームメイトです。現在、チームが毎週 Extra を構築するために出荷するコードの 3 分の 1 近くを Avery が作成しています。

記事本文:
自らを構築する AI チームメイト: Avery を当社の DNA に組み込んだ方法 | BuildForever Engineering チームについて 採用情報 お問い合わせ AI ファースト エンジニアリング / 2026 年 6 月 18 日
自らを構築する AI チームメイト: Avery を当社の DNA に組み込んだ方法
Avery は、BuildForever のチームが毎日使用するシステムやツールに組み込まれている AI エンジニアリングのチームメイトです。現在、チームが毎週 Extra を構築するために出荷するコードの 3 分の 1 近くを Avery が作成しています。
先週、Avery は、5 人のエンジニア チームが運用環境に出荷したコードの 31% を作成しました。しかし、プロンプトが表示されるだけではなく、問題を特定し、根本原因を診断し、修正を生成し、マージされた運用コードにそれを反映させます。私たちのような小さなチームにとって、これは影響力の大きな変化です。常にチームメイトがバグ、アラート、メンテナンス作業を 24 時間体制で出荷コードに変換します。サイド プロジェクトとして始まったものが、現在では BuildForever のあらゆる部分に組み込まれており、社内の誰もがコードを本番環境に出荷できるようになりました。
Avery は Slack に住んでおり、自分の Mac mini で実行しており、私たちのチームが毎日使用しているシステム (コード、プル リクエスト、レビュー用の GitHub) にアクセスできます。問題に対しては直線的。バグレポートやコミュニティからのフィードバックについては電子メールと WhatsApp をご利用ください。 TestFlight と App Store Connect によるクラッシュ レポートとユーザー レビュー。製品分析用の Mixpanel。実際に何が起こったかを追跡するには、Cloud Logging の本番環境ログを使用します。
コーディングアシスタントと異なる理由
あなたは Slack で Avery と話します。スレッドで言及すると、そのスレッドに関連付けられた永続的なエージェント セッションが起動されるため、会話とその背後にある作業はリンクされたままになります。フォローアップでは、最初からやり直すのではなく同じセッションが再開され、スレッド間で耐久性のあるメモリが保持され、要求を待たずにスケジュールされたジョブが独自に実行されます。そのp

永続性により、エイブリーはクロード・コードやデヴィンのような永続的なエージェントと同じファミリーに属します。
しかし、Avery は単なるコード補完プラグインではありません。Avery の特徴は、その積極性、メモリ、コンテキストです。これは、実際に実行されているすべてのシステムに接続されています。
実際の Mac ハードウェア上で動作します。 Avery は、Extra iOS アプリをコンパイルして実行し、テスト アカウントにログインして、修正の動作を記録します。 Linux ボックス上のエージェントは Swift を読み取ることはできますが、実行したり、結果を表示したりすることはできません。
推測ではなく本番環境からデバッグします。 Avery は本番環境の GCP ログ、データベース、分析、あらゆるバグ チャネルをまとめて収集するため、あいまいなレポートはデータに裏付けられた根本原因になります。
PR ライフサイクル全体を管理します。 1 つの Slack メッセージから、Avery は分離されたワークツリーを起動し、実装、テストし、PR を開き、チケットをリンクし、スレッドに追跡し、要求に応じてマージしてクリーンアップします。各タスクは独自のワークツリーを取得するため、複数のエンジニアが互いに干渉することなく、Avery に並行して作業を引き渡すことができます。
会社に合わせてバージョン管理されています。 Avery のスキル、スケジュールされたジョブ、チームの規則は、コードの隣のリポジトリに存在します。これらは他のすべてと同じプル リクエストを通じて改善され、Avery はそれらの PR を自身のスキルに対してオープンし、他の変更と同様にレビューできます。ノーコードのダッシュボードやレビュー以外の自己編集はありません。 Avery は自らを構築します。
AI エンジニアリング ツールの次の波は、IDE 内にだけ存在するわけではありません。それらは、オペレーティング システム、ツール、人間の会話、ログ、分析、コードを製品に変える儀式など、企業の DNA に組み込まれることになります。
Avery は 1 つのユーザー レポートから運用ログを調べて根本原因を見つけ、修正を含む PR を開き、Mac mini でアプリを実行して検証し、同じ Slack スレッドでレポートを返します。
Avery は常に最初のパスを実行し、t を解放します。

彼はチームを組んで他の重要な仕事に集中します。
Avery は、BugSnag を使用して自動 iOS クラッシュ レポートも設定しました。新しいクラッシュが発生すると、Avery が独自にそれを検出し、根本原因を追跡し、修正を含む PR を開いて検証します。エンジニアが関与する必要はありません。
Avery はまた、バックエンドで API エラーの傾向を常に監視しています。あるクラスのエラーが急増し始めると、チームに警告し、ベースラインに対するワーカー ログを通じて回帰を追跡し、ページになる前に修正を加えます。
私たちがコンピューターから離れているときは、モバイル版 Avery がゲームチェンジャーとなっています。 Extra の使用中に発生したバグは、Avery に直接渡して外出先で修正できます。これにより、SLA に照らしてユーザー向けの修正を行う速度が 10 倍になりました。
Avery の強みの 1 つは、スクリーンショットやビデオによるライブ検証による新機能の実装です。ここでは、BuildForever の共同創設者である Steven が、Avery がどのように新機能を構築できるかを示しています。
Avery は任意の Slack チャネルに所属しているため、チーム全体が他のメンバーが Avery をどのように使用しているかを確認でき、各リクエストが学習の機会に変わります。見ているだけで新しいプロンプトの方法を習得でき、誰でもスレッドの途中でチャイムを鳴らして、起こっていることを小突いたり、修正したり、拡張したりできます。その結果、私たちがこれまでに見た中で最もマルチプレイヤーなプロンプトが生まれました。プライベートな 1 対 1 のセッションではなく、オープンな場で、数人の人々が協力して変化に向けて Avery を導いています。
分析は明らかな利点の 1 つです。新しい機能が採用されているかどうかを尋ねると、Avery は Mixpanel 番号、バックエンド セッション ログを取得し、ユーザーがどこから離脱したかを見つけ、ダッシュボードを更新し、スレッドに回答を投稿します。そのため、製品に関する質問は誰かが空くのを待つ必要はありません。
エイブリーは求められたときだけ行動するわけではありません。最も価値のある作業の一部は、スケジュールに従って実行されます。

ループに参加し、チームがすでに確認している場所に結果を投稿します。
Engineering Daily Pulse - 平日の夕方に、出荷されたもの、低下したもの、そしてまだ注意が必要なものについて読みます。
個人的な朝の集中 - 各チームメイトは、予定していること、待っていること、今日は何に集中するかなど、早い段階で概要を受け取ります。
毎日のバグ スイープ - 1 日 2 回、すべてのレポート チャネルをスイープして、すでに添付されている証拠とともにトリアージされ、重複が排除された問題を調査します。
iOS および Web PR レビュー キュー - まだレビューを待っている最もリスクの高い変更を表面化する毎週のキュー。
セキュリティ ログのレビューと API 侵入テスト - セキュリティ ログを毎日検査し、独自の API の自動侵入テストを行います。
エラー ダイジェスト - Cloud Logging のエラー ダイジェストが数時間ごとに作成されるため、回帰がすぐに表面化します。
リリース レポート - 次の iOS 出荷に向けた週次準備状況レポート。
Today タブの品質レビュー - フィード自体の毎週の製品品質パス。
これらのスキルはリポジトリ内に存在するため、各実行は前回と同じ基準を保持し、Avery は自身で開くことができる PR を通じてスキルを磨きます。
Avery は人々の電子メールやユーザー レポートを扱うため、そのガードレールは明確です。チーム全体がスレッドに参加している場合でも、ユーザーの個人情報や電子メールの内容が Slack に投稿されることはありません。
バグレポートが届くと、Avery はそれを調査し、ユーザーが私たちと共有することを選択した内容のみに基づいて修正を提案します。当社が製品に適用するのと同じプライバシー バーがエージェントにも適用されます。
Slack は Avery の本拠地ですが、そこに到達する唯一の方法ではありません。 Avery は Mac および iOS のネイティブ アプリとしても実行できるため、同じチームメイトがラップトップまたは携帯電話でクリックするだけでアクセスできます。
それによってオンボーディングが変わります。新しく参加する人は、セットアップ ドキュメントを読む必要はありません。Avery をインストールすると、接続するツールや、

チームがすでに作業している場所にプロビジョニングするアカウント。新入社員が最初に出会うのは、会社の運営方法をすでに知っているチームメイトです。
Avery が役立つようになったのは、数週間にわたって非常に具体的なチューニングを行った後です。初期の段階では、診断で止まったり、古いコンテキストを過度に信頼したり、十分な証拠がないまま PR を開いたり、間違った Slack コンテキストで返信したり、根本原因ではなく症状にパッチを当てたりしていました。そこで私たちは、あらゆる失敗を操作手順に変えました。つまり、新しいワークツリーを使用し、スレッド履歴を検査し、運用証拠で主張を根拠付け、UI 修正のスクリーンショットを表示し、焦点を絞ったテストを実行し、セッション ログを書き込み、危険なマージを人間に任せました。最近では、Avery はその作業の多くを Avery 自体で行っています。私たちは Avery を使用して Avery を構築し、独自のスキルを作成し、独自のバグを修正し、独自の改善を出荷しています。そのチューニングは、Avery を他とは違うと感じさせるものでもあります。Avery はすぐに適応し、個性を持ち、会社が実際にどのように運営されているかを学習するので、別の管理ツールになるのではなく、既存のシステムに織り込むことができます。私たちは依然として、好み、アーキテクチャ、プライバシー、判断が重要な呼び出しに関してそれをオーバーライドします。
私たちにとっての教訓は、AI エンジニアリング ツールの次の波は IDE 内にだけ存在するわけではないということです。これらは、オペレーティング システム、ツール、人間の会話、ログ、分析、コードを製品に変える儀式など、企業の DNA に組み込まれることになります。これは、汎用エージェントをワークフローにドロップするよりもはるかに便利です。
このようなものを構築している場合、または会社に Avery を組み込みたい場合は、ぜひチャットしてください。 Avery にメールしてください (avery@buildforever.com)
そんな働き方をしたいエンジニアも募集中です。 Careers@buildforever.com までご連絡ください。
ああ、気になる方は、エイブリーも実在の人物です。当社のエンジニアの 1 人、ルークの信じられないほどかわいい 1 歳半の子供です。人間のエイブリーはまだ

私は彼の最初の PR に取り組んでいますが、彼には時間があります。

## Original Extract

Avery is an AI engineering teammate wired into the systems and tools BuildForever's team uses every day — and it now writes nearly a third of the code the team ships to build Extra every week.

An AI teammate that builds itself: how we wired Avery into our company’s DNA | BuildForever Engineering About Team Careers Contact AI-first engineering / June 18, 2026
An AI teammate that builds itself: how we wired Avery into our company’s DNA
Avery is an AI engineering teammate wired into the systems and tools BuildForever's team uses every day — and it now writes nearly a third of the code the team ships to build Extra every week.
Last week, Avery wrote 31% of the code our five-engineer team shipped to production. But not just by being prompted - it identified the issues, diagnosed the root causes, generated the fixes, and carried them through to merged production code on its own. For the tiny team we are, that's a step-change in leverage: an always-on teammate turning bugs, alerts, and maintenance work into shipped code around the clock. What started as a side project is now woven into every part of BuildForever, and now lets anyone at the company ship code to production.
Avery lives in Slack, runs on its own Mac mini, and has access to the systems our team uses every day: GitHub for code, pull requests, and reviews; Linear for issues; email and WhatsApp for bug reports and community feedback; TestFlight and App Store Connect for crash reports and user reviews; Mixpanel for product analytics; and production logs in Cloud Logging to trace what actually happened.
Why it's different from a coding assistant
You talk to Avery in Slack. Mention it in a thread and it spins up a persistent agent session tied to that thread, so the conversation and the work behind it stay linked. Follow-ups resume the same session instead of starting over, it carries a durable memory across threads, and it runs scheduled jobs on its own without waiting to be asked. That persistence puts Avery in the same family as persistent agents like Claude Code or Devin.
But Avery is not just a code-completion plugin - what sets it apart is its proactivity, memory and context. It's wired into every system we actually run on:
It runs on real Mac hardware. Avery compiles and runs the Extra iOS app, logs into test accounts, and records the fix working. An agent on a Linux box can read Swift, but it can't run it or show you the result.
It debugs from production, not guesses. Avery pulls our production GCP logs, databases, analytics, and every bug channel together, so a vague report becomes a root cause backed by data.
It owns the whole PR lifecycle. From one Slack message, Avery spins up an isolated worktree, implements, tests, opens the PR, links the ticket, and tracks it back to the thread - then merges and cleans up when asked. Each task gets its own worktree, so several engineers can hand Avery work in parallel without stepping on each other.
It's versioned with the company. Avery's skills, scheduled jobs, and team conventions live in the repo next to the code. They improve through the same pull requests as everything else - and Avery can open those PRs against its own skills, reviewed like any other change. There is no no-code dashboard and no self-editing outside of review. Avery builds itself.
The next wave of AI engineering tools won't just live in the IDE. They'll be embedded in the company's DNA - the operating systems, tools, human conversations, logs, analytics, and rituals that turn code into product.
From a single user report, Avery digs through production logs to find the root cause, opens a PR with the fix, runs the app on its Mac mini to verify it, and reports back in the same Slack thread.
Avery always does the first pass, freeing up the team to focus on other critical work.
Avery also set up automatic iOS crash reporting with BugSnag. A new crash shows up, Avery picks it up on its own, traces it to the root cause, opens a PR with the fix, and verifies it - without any engineer needing to be in the loop.
Avery also constantly watches our backend for trends in API errors. When a class of errors starts spiking, it alerts the team, traces the regression through the worker logs against a baseline, and puts up a fix - before it turns into a page.
When we're away from our computers, Avery on mobile has been a game changer. Any bug we hit while using Extra, we can hand straight to Avery and fix it on the go - which has 10x'd how fast we turn around user-facing fixes against our SLAs.
One of Avery's superpowers is implementing new features, with live verification through screenshots and videos. Here's Steven, co-founder of BuildForever, showing how Avery can build new features.
Because Avery lives in any Slack channel, the whole team can see how everyone else uses Avery, which turns each request into a chance to learn - you pick up new ways to prompt just by watching, and anyone can chime in mid-thread to nudge, correct, or extend what's happening. The result is the most multiplayer prompting we've seen: several people steering Avery toward a change together, in the open, instead of working in private one-on-one sessions.
Analytics is one of the clearer wins. Ask whether a new feature is being adopted, and Avery pulls the Mixpanel numbers, our backend session logs, finds where users drop off, updates the dashboard, and posts the answer in-thread - so product questions don't have to wait for someone to free up.
Avery doesn't only act when it's asked. Some of its most valuable work runs on a schedule, with nobody in the loop, posting results where the team already looks:
Engineering Daily Pulse - every weekday evening, a read on what shipped, what regressed, and what still needs attention.
Personal Morning Focus - each teammate gets an early brief: what's on their plate, what's waiting on them, what to focus on today.
Daily Bug Sweep - twice a day, sweeps every report channel into triaged, de-duplicated issues with the evidence already attached.
iOS & Web PR Review Queues - weekly queues that surface the riskiest changes still waiting on review.
Security Log Review & API Pentest - a daily pass over the security logs and an automated penetration test of our own APIs.
Error Digests - a Cloud Logging error digest every few hours, so regressions surface fast.
Release Report - a weekly readiness report for the next iOS ship.
Today Tab Quality Review - a weekly product-quality pass on the feed itself.
Because these skills live in the repo, each run holds the same standard as the last, and Avery sharpens them through PRs it can open itself.
Avery works with people's email and user reports, so its guardrails are explicit. It never posts a user's personal information or email contents into Slack, even when the whole team is in the thread.
When a bug report comes in, Avery investigates it and proposes a fix working only from what that user has chosen to share with us. The same privacy bar we hold the product to applies to the agent.
Slack is Avery's home base, but it isn't the only way to reach it. Avery also runs as native Mac and iOS apps, so the same teammate is a click away on a laptop or a phone.
That changes onboarding. When someone new joins, they don't work through a setup doc - they install Avery, and it walks them through getting set up in seconds: the tools to connect, the accounts to provision, where the team already works. The first thing a new hire meets is the teammate that already knows how the company runs.
Avery only became useful after a few weeks of very specific tuning. Early on it would stop at diagnosis, over-trust stale context, open PRs without enough proof, reply in the wrong Slack context, or patch symptoms instead of root causes. So we turned every failure into operating instructions: use fresh worktrees, inspect thread history, ground claims in production evidence, show screenshots for UI fixes, run focused tests, write session logs, and leave risky merges to humans. These days Avery does a lot of that work on itself - we use Avery to build Avery, drafting its own skills, fixing its own bugs, and shipping its own improvements. That tuning is also what makes Avery feel different: it adapts quickly, has personality, and learns how the company actually operates so it can weave into the existing systems instead of becoming another tool to manage. We still override it on taste, architecture, privacy, and judgment-heavy calls.
The lesson for us is that the next wave of AI engineering tools won't just live in the IDE. They'll be embedded in the company's DNA - the operating systems, tools, human conversations, logs, analytics, and rituals that turn code into product, which is far more useful than dropping a generic agent into a workflow.
If you're building something like this, or would like to embed Avery at your company, we'd love to chat. Just email Avery at avery@buildforever.com
We're also hiring engineers who want to work this way. Reach out to careers@buildforever.com
Oh, and if you're wondering: Avery is a real person too - the incredibly cute one-and-a-half-year-old of one of our engineers, Luke. The human Avery is still working on his first PR, but he's got time.
