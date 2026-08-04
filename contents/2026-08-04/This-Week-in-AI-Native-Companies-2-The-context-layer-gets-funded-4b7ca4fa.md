---
source: "https://www.insideainative.com/p/this-week-in-ai-native-companies-555"
hn_url: "https://news.ycombinator.com/item?id=49162757"
title: "This Week in AI Native Companies #2: The context layer gets funded"
article_title: "This Week in AI Native Companies #2: The context layer gets funded"
author: "mmayernick"
captured_at: "2026-08-04T00:07:30Z"
capture_tool: "hn-digest"
hn_id: 49162757
score: 2
comments: 0
posted_at: "2026-08-03T23:39:00Z"
tags:
  - hacker-news
  - translated
---

# This Week in AI Native Companies #2: The context layer gets funded

- HN: [49162757](https://news.ycombinator.com/item?id=49162757)
- Source: [www.insideainative.com](https://www.insideainative.com/p/this-week-in-ai-native-companies-555)
- Score: 2
- Comments: 0
- Posted: 2026-08-03T23:39:00Z

## Translation

タイトル: 今週の AI ネイティブ企業 #2: コンテキスト層に資金提供
説明: バズは 7 日間で 3 倍になり、2 つの 1,000 万ドルのシードは 48 時間間隔で着地し、動作仕様はオープンソースになります。

記事本文:
今週の AI ネイティブ企業 #2: コンテキスト層に資金提供
今週の AI ネイティブ企業 #2: コンテキスト層に資金提供
バズは 7 日間で 3 倍になり、2 つの 1,000 万ドルのシードは 48 時間間隔で着地し、動作仕様はオープンソースになります。
Michael Mayernick 2026 年 8 月 1 日 1 シェア 私は毎週、AI エージェントを実行している企業の内部から、現場で見ているものについてこの記事を書いています。今週、現場は実際のお金と実際の数字をボードに載せました。
撮影されたコンパニオンとして、Matt Nicosia が加わりました。Grey Swan の GTM のデータおよび運用責任者であり、その前はコンテキスト領域で Zayer の構築者でした。そのため、彼は今週のストーリーのほとんどを直接体験してきました。クロード・タグが実際に同僚としてどのような人物なのか（彼のチームの評決は「より低いリフトのような方法でプロセスを置き換える」だが、ジョークを頑張りすぎたとして楽しいチャンネルから削除した）、コンテキスト製品が内部から区別するのが非常に難しい理由、そしてマットが「長期にわたるAI精神病」と呼ぶものに対する防御策としての動作仕様に迫ります。
1. バズは 1 週間でほぼ 3 倍になりました。 Block の人間とエージェントのワークスペースは、7 日間で約 6,700 個の GitHub スターから 18,400 個に増加し、途中で 6 つのリリース (v0.4.20 から v0.5.2) を出荷しました。キャンバスは到着し、リリースのペースは衰えていない。チャット ルームの形状についてどう考えても、需要のシグナルは大きく、チームは人々とエージェントが一緒に働く場所を望んでいる。しかし、先週からの未解決の疑問は残っています。部屋にはすべてが表示されます。何が真実なのかはわかりません。
Inside AI Native をお読みいただきありがとうございます。新しい投稿を受け取るには無料で購読してください。
2. タイプが扉を開きました。 Type - Halp の背後にあるチームによる人間と AI のワークスペース (Slack 支援、Atlassian に売却) - プライベート ベータから移行: 月額一律 50 ドル、無制限

人間と AI を組み合わせた座席は、独自の Claude または ChatGPT サブスクリプションを持ち込んでください。際立った設計上の選択は「スペース」です。これは、1 つのグローバルな山ではなく、コンテキストごとに独自のメモリ、コネクタ、スキルを運ぶ共有コンテナです。さらに、チームが段階的に移行できるように Slack ブリッジを追加します。彼らの X のプロフィールには「最初のマルチプレイヤー AI 製品」と書かれていますが、これは YC の現在の Request for Startups がカテゴリ名として使用している正確なフレーズです。私はここ 1 週間ベータ版を利用してきましたが、ワークスペースを介したアクセス制御と整理に対する敏感さを高く評価しています。この部屋戦争には、値札を付けた 2 番目の本格的な参入者が現れました。
3. コンテキスト レイヤーの 2 つの $1,000 万シードを 48 時間間隔で配置します。 Modus は、Insight Partners が主導する「Context Warehouse」（エージェントがクエリを実行できるように、企業が実際にどのように運営されているかを継続的に維持するモデル）でステルス状態から抜け出しました。 Credible Data は、オープンソースの Malloy セマンティック モデリング言語に基づいて構築された「AI Context Engine」に対して同額を調達しました。同じ週、同じレイヤー、同じ単語: コンテキスト。 2 つのファンドが 1 週間以内に同じ名詞で同じ小切手を切った場合、市場はその名詞が重要であると判断します。
4. 動作仕様が公開されます。 Braintrust と Basis は、エージェントがどのように動作するかを文書化するためのオープンスタンダードである、コードでバージョン管理されたマークダウン仕様である、agentbehavior.dev をリリースしました。起源のストーリーが興味深い部分です。Basis は実稼働環境で税務担当者を数時間、場合によっては数日間運営していますが、結果の検証がその地平線で破綻することがわかりました。決定が多すぎ、一般化が少なすぎるため、代わりにプロセスを監督し、苦労して得た行動を文書化する方法が仕様です。仕様は意図的に散文的であり (「人間が明確に表現した散文は、エージェント システムを駆動する最も有効な手段である」)、意図的に実行不可能になっています。

それらは評価設計に情報を提供し、その発表は「最終的には報酬が得られる」動作を指しており、仕様トレーニングの合図となるでしょう。この分野がグラウンドトゥルースとして書かれた動作に到達するのは今月で 3 回目です (Google の OKF 信頼フィールド、コードスコープでの Fluent の動作仕様、現在は評価されています)。そして 3 つすべてが、動作の変更を誰が承認し、それを誰がチェックするかという同じ未解決の質問に止まります。例を読む価値があります。
5. オープンソースのソフトウェアファクトリー。 Mrinal Wadhwa (Ockam の CTO) は、オープンソースの「自己改善ソフトウェア ファクトリ」である Fluent をリリースしました。これは、スキルとして Claude Code または Codex にインストールするオープンソースの「自己改善ソフトウェア ファクトリ」です。観察はブリーフになり、ブリーフはテストによる動作仕様になり、作業は分離されたワークツリーでライター、テスター、およびレビュー担当者の役割を通じて実行され、レッスンはマージで完了します。その構造がなければ何が起こるかについての彼のセリフは、「自動化されたもぐらたたきゲーム」です。工場自体が製品になりつつあり、彼はそれを屋外で構築しています。
6. Andrew Ng の OpenWorker がベータ版をオープンしました。 Ng 氏と Rohit Prasad 氏によるローカルファーストのオープンソース デスクトップ エージェントは、チャットの代わりに完成した成果物を返します。25 以上のコネクタ、アクションごとの承認など、すべてがマシン上にあります。ワーカー層は本当に良くなり、急速に進歩しています。何から機能するかは依然として各人独自の設定です。
7. そして今週の議論: 判断はどこにあるのか?マーティン・カサドは、「ハーネスは少ないほど良い、モデルは魔法だ」と「トレーニング後とハーネスが勝利する」の間で揺れ動いていると投稿した。アンドリュー・チェンは、自身のエージェントスキルが「結局、見直すべきことがどんどん増えていくだけだ。今欲しいのは行動だ」と投稿した。 2 人の投資家が同じ週に、反対側から 1 つの質問を回覧しました。エージェントが仕事をするとき、かつては回転にあった判断はどこにあるのですか

実際に行くの？私には考えがあるのですが、月曜日の手紙はまさにこれについてです。
それがその週です。あなたが望むなら、そのパターンです。部屋は増えて値札がつき、お金は背景に、基準は行動に、そして議論は判決に移ります。最後の 3 つは、同じ欠落しているレイヤーの 3 つの名前です。今週の「コンテキストは新しいプロセス」で、そのプロセスの半分についてさらに詳しく説明しました。私たちがプロセスと呼ぶもののほとんどは、常に人々の間でコンテキストを移動させるための機械であり、エージェントはその根底にある経済性を変えただけです。
Inside AI Native をお読みいただきありがとうございます。新しい投稿を受け取るには無料で購読してください。
1 シェア この投稿についてのディスカッション コメント 再スタック トップ 最新 投稿はありません

## Original Extract

Buzz triples in seven days, two $10M seeds land 48 hours apart, and behavior specs go open source.

This Week in AI Native Companies #2: The context layer gets funded
Subscribe Sign in This Week in AI Native Companies #2: The context layer gets funded
Buzz triples in seven days, two $10M seeds land 48 hours apart, and behavior specs go open source.
Michael Mayernick Aug 01, 2026 1 Share Every week I write this from inside a company that runs on AI agents, about what we’re seeing in the field. This week the field put real money and real numbers on the board.
For the filmed companion, Matt Nicosia joins me - Head of Data and Ops for GTM at Gray Swan, and before that the builder of Zayer in the context space, so he’s lived most of this week’s stories firsthand. We get into what Claude Tag is actually like as a coworker (his team’s verdict: “it replaces process in a much lower lift kind of way,” though they did remove it from the fun channel for trying too hard at jokes), why context products are so hard to differentiate from the inside, and behavior specs as a guard against what Matt calls “long-running AI psychosis.”
1. Buzz nearly tripled in a week. Block’s humans-and-agents workspace went from about 6,700 GitHub stars to 18,400 in seven days, shipping six releases along the way (v0.4.20 to v0.5.2). Canvases landed, the release pace hasn’t slowed, and whatever you think of the chat-room shape, the demand signal is loud: teams want a place where people and agents work together. The open question from last week stands, though. A room shows you everything; it doesn’t tell you what’s true.
Thanks for reading Inside AI Native! Subscribe for free to receive new posts.
2. Type opened its doors. Type - the human-plus-AI workspace from the team behind Halp (Slack-backed, sold to Atlassian) - moved out of private beta: $50 a month flat, unlimited seats, human and AI, bring your own Claude or ChatGPT subscription. The design choice that stands out is “spaces”: shared containers that carry their own memory, connectors, and skills per context, instead of one global pile - plus a Slack bridge so teams can migrate gradually. Their X bio now reads “the first multiplayer AI product,” which is the exact phrase YC’s current Request for Startups uses as a category name. I’ve been in the beta for the past week and appreciate the sensitivity to access control and organization through workspaces. The room war now has a second serious entrant with a price tag.
3. Two $10M seeds for the context layer, 48 hours apart. Modus came out of stealth with a “Context Warehouse” - a continuously maintained model of how an enterprise actually operates, for agents to query - led by Insight Partners. Credible Data raised the same amount for an “AI Context Engine” built on the open-source Malloy semantic-modeling language. Same week, same layer, same word: context. When two funds write the same check for the same noun in one week, the market has decided the noun matters.
4. Behavior specs go open. Braintrust and Basis released agentbehavior.dev , an open standard for documenting how an agent should behave - markdown specs, versioned with code. The origin story is the interesting part: Basis runs tax agents in production for hours, sometimes days, and found that verifying outcomes breaks down at that horizon - too many decisions, too little generalization - so they supervise the process instead, and the specs are how those hard-earned behaviors get documented. The specs are deliberately prose (”clear, human-articulated prose is the highest leverage way to drive agentic systems”) and deliberately not executable - they inform eval design, with the announcement pointing at “eventually rewarding” behaviors, which would make the specs training signal. It’s the third time this month the field has landed on written-behavior-as-ground-truth (Google’s OKF trust fields, Fluent’s behavior specs at code scope, now evals) - and all three stop at the same open questions: who approves a behavior change, and what checks it. Worth reading the examples.
5. The software factory, open-sourced. Mrinal Wadhwa (CTO of Ockam) released Fluent , an open-source “self-improving software factory” you install into Claude Code or Codex as a skill: observations become briefs, briefs become behavior specs with tests, work runs through writer, tester, and reviewer roles in isolated worktrees, and the lessons land with the merge. His line for what happens without that structure: “an automated game of whack-a-mole.” The factory itself is becoming the product, and he’s building it in the open.
6. Andrew Ng’s OpenWorker opened its beta . The local-first, open-source desktop agent from Ng and Rohit Prasad returns finished deliverables instead of chat - 25+ connectors, per-action approvals, everything on your machine. The worker layer is getting genuinely good, fast. What it works from is still each person’s own setup.
7. And the week’s argument: where does judgment live? Martin Casado posted that he vacillates between “the less harness the better, models are the magic” and “post-training plus harness wins.” Andrew Chen posted that his agent skills “just end up generating more and more things for me to review - what I want now is actions.” Two investors, same week, circling one question from opposite ends: when agents do the work, where does the judgment that used to live in review actually go? I have thoughts, and Monday’s letter is about exactly this.
That’s the week. The pattern, if you want one: the rooms multiplied and got price tags, the money went to context, the standards went to behavior, and the argument went to judgment. The last three are three names for the same missing layer. I went deeper on the process half of that this week in “Context is the new process” : most of what we call process was always machinery for moving context between people, and agents just changed the economics underneath it.
Thanks for reading Inside AI Native! Subscribe for free to receive new posts.
1 Share Discussion about this post Comments Restacks Top Latest No posts
