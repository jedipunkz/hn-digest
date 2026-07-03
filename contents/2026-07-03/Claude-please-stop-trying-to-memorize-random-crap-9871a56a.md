---
source: "https://12gramsofcarbon.com/p/agentics-memorizing-session-transcripts"
hn_url: "https://news.ycombinator.com/item?id=48776232"
title: "Claude, please stop trying to memorize random crap"
article_title: "Agentics: Memorizing Session Transcripts Isn't Useful"
author: "theahura"
captured_at: "2026-07-03T15:49:50Z"
capture_tool: "hn-digest"
hn_id: 48776232
score: 1
comments: 0
posted_at: "2026-07-03T15:32:51Z"
tags:
  - hacker-news
  - translated
---

# Claude, please stop trying to memorize random crap

- HN: [48776232](https://news.ycombinator.com/item?id=48776232)
- Source: [12gramsofcarbon.com](https://12gramsofcarbon.com/p/agentics-memorizing-session-transcripts)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T15:32:51Z

## Translation

タイトル: クロード、適当に覚えようとするのはやめてください
記事のタイトル: Agentics: セッション記録の暗記は役に立たない
説明: スクラッチではなく、アーティファクトを追跡します。代替タイトル: クロード、ランダムにくだらないことを覚えようとするのはやめてください

記事本文:
エージェントティクス: セッション記録を暗記しても役に立たない
炭素12グラム
Agentics: セッション記録の記憶は役に立たない
スクラッチではなく、アーティファクトを追跡します。代替タイトル: クロード、ランダムにくだらないことを覚えようとするのはやめてください
theahura 7月 02, 2026 1 6 シェア
31 件の「いいね！」という数は多くないように思えるかもしれませんが、実際にはサブスタック ノートの全員です。エージェントが他の形式のコンテキストにアクセスできる場合、エージェントが以前のトランスクリプト セッションへの検索アクセス権を持っている場合、SWE タスクのパフォーマンス上の利点はゼロであることがわかりました。また、人間が関与していない限り、エージェントのコンテキストを改善するためにセッション記録を自動的にトロールしようとしても、あまりメリットは見つかりませんでした。
読んでいただきありがとうございます!無料で購読して新しい投稿を受け取り、私の仕事をサポートしてください。
直感的には、エージェントとエンジニアの間の記録には貴重な情報がたくさん含まれているように感じられます。おそらく、コードが存在する理由やユーザーの意図に関する情報が含まれるでしょう。あるいは、ユーザーが試して破棄した他のアプローチが含まれている可能性があります。少なくとも、エージェントが理解を深めるために使用できる追加のコンテキストがある程度含まれることになります。私はこれを強く信じていたので、私の会社はこのコンセプトに基づいて製品全体を構築しました。私は人々に、「セッションの記録は新しい油だ」、コード自体よりも価値があるとよく言っていました。
他の人も明らかに同様の考えを持っていました。それが、(もちろん) クロード コード自体を含む、セッション バックアップ メモリを実行するためのさまざまなツールが存在する理由です。
最も一般的なアーキテクチャは次のようなものだと思います。
組織全体のすべてのトランスクリプトを DB に保存する
ベクトル検索、エラスティック検索、または SQL 検索レイヤーをその前に置きます。野心的なチームは 3 つすべてを使用します。おそらくグラフが関係するでしょう。

MCP を使用するか、スキルを備えた CLI を公開することで、エージェントがこれを利用できるようにします。
私たちにとって、この追加作業は大きな違いを生むようには思えません。むしろ、セッション検索アクセスの有無にかかわらず何ヶ月にもわたるテストに基づくと、モデルが悪化する可能性があります。
私たちのチームが非常に重視していることの 1 つは、成果物のコーディングです。私たちはもう手作業でコードを書くことはありません。 PR を読みやすくするために、優れたコミット メッセージ、優れた PR メッセージ、包括的なドキュメントを重視します。すべてのコード変更には、コードと一緒にコミットされる広範なメタデータが付属します。私たちのエージェントがコードの一部に取り組むときは、ドキュメントと以前の PR を確認するように指示されます。
言い換えれば、エージェントはすでにトランスクリプトに関する貴重な情報をすべて抽出し、それを必要な場所に簡単にアクセスできる場所に保存しています。したがって、エージェントがトランスクリプト検索サーバーを使用すると、エージェントが最初に書き留めないと決めたすべての内容を取得しながら、エージェントがすでに知っている内容を読み取るためにトークンを費やすことになります。もしかしたら、時々、役に立つ情報がそこにあるかもしれません。しかし、ほとんどの場合、エージェントは疑似無意味なスクラッチ パッドを眺めているだけであり、そのために貴重なトークンを無駄にしています。
また、エージェントは、長期記憶を維持するために重要な機能であるコンテキストを実際に削除することも苦手です。つまり、文字通り何千ものセッションを通じて、一度もそのようなことが起こったことを見たことがありません。
これは、賢明なプロンプトエンジニアリングによって除去できる特性ではありません。エージェントには状態がないため、入力コンテキスト ウィンドウ内のすべてがグランド トゥルースであると想定する必要があります。コードのすべての行、メモリのすべての既存ビット、すべてのトークンは、たとえそのコードやメモリが

以前のエージェント セッションによって行われたランダムな決定であり、人間によって検討されることも、理解されることもありません。この意図のドリフトは、エージェントが自律的にメモリ ベースを構築しようとするほど悪化します。
私の知る限り、入力データが破損していることを前提としたコーディング ベンチマークは存在しません。実際、モデルは入力データが間違っていると仮定することでペナルティを受けます。これは部分的に調整の問題でもあります。エージェントに意図しないことをさせたくないのですが、「コードベースを削除しない」と「入力コンテキストの一部を削除する」という問題を解決する簡単な方法はありません。
モデルは実際には独自のメモリを管理できないため、自動記憶は結局同じ場所に行き着きます。大量のゴミがトークンを食べ、請求書が膨れ上がり、モデルの品質が低下します。
ネットネット、私はインデックスを作成して保存し、セッション記録をエージェントに表示するツールに対して非常に弱気になっています。セッションの記録はチームの可観測性には役立つかもしれませんが、エージェントの能力が向上するわけではありません。
それは、エージェントが時間の経過とともにコンテキストを学習する役割を持たないという意味ではありません。私たちは社内のノリボットを使用して、PR、Slack、ドライブなどで会社で起こったことすべてを毎週レビューします。そして、彼らは組み込みのノリスキルセットに対する一連の変更を提案し、Slack でチームをタグ付けします。これらはすべてデフォルトで拒否されます。変更を受け入れるには、実際に diff を見て、それが意図に合っていることを確認する必要があります。
このうち 20% 未満しか受け入れられません。これは、これらの「自動」更新の 80% がモデルを悪化させていた可能性があることを意味します。数百人の組織がこれらの「更新」を常に自動的に保存していたら、それがどれほど持続不可能になるか想像できません。
読んでいただきありがとうございます!無料で購読して新しい投稿を受け取り、私の仕事をサポートしてください。
1 6 共有 この件についてのディスカッション

コメントを投稿 Restacks B P 17h Liked by theahura そこで、「コード行または TODO につながるビジネス コンテキストの広さはどれくらいだったのか」という興味深い質問があります。これは企業によって異なる場合があります。場合によっては、よく管理されているストーリー トラッカーなどへのコメント リンクや、現代の LLM がフォローするよう奨励されているやる気を起こさせる Slack スレッドを用意するだけで十分な場合があります。しかし、そのやる気を起こさせる例は、多くの同僚が時間をかけてこの分野を探索したことによる、数多くの経験や洞察の広範な探索の一部にすぎない場合もあります。今後の反復では、1 つの例に対して過剰なインデックスが作成される可能性があります。これは、私たちが厳選したマーケットプレイスでよく目にするもので、事前に何十回もリストに手動でパッチを適用したり上書きしたりする中で、統合に対する修正が存在します。そして、その修正を修正している場合、古いコードの動機となったものはすべて見つかるでしょうか?まるでチェスタトンの柵と暗い森が出会ったようです。私はまだ御社の製品を諦めません！
返信 シェア ブライアン シュナイダー 20h 「いいね！」したユーザー: theahura この点で、それは最も人間的です。
返信 シェア さらにコメント 4 件... トップ 最新のディスカッション 投稿はありません

## Original Extract

Keep track of artifacts, not scratch. Alt title: Claude, please stop trying to memorize random crap

Agentics: Memorizing Session Transcripts Isn't Useful
12 Grams of Carbon
Subscribe Sign in Agentics: Memorizing Session Transcripts Isn't Useful
Keep track of artifacts, not scratch. Alt title: Claude, please stop trying to memorize random crap
theahura Jul 02, 2026 1 6 Share
31 likes may not seem like a lot, but that's actually everyone on substack notes We have found zero performance benefit on SWE tasks when agents have search access to their previous transcript sessions, provided they have access to other forms of context. We also have not found much benefit in trying to automatically trawl through session transcripts to improve agent context, unless there is a human in the loop.
Thanks for reading! Subscribe for free to receive new posts and support my work.
Intuitively it feels like there's a lot of valuable information in a transcript between an agent and an engineer. Maybe it would have information about why the code exists, about user intent. Or it might have the other approaches that a user tried and discarded. At the least, it would have some amount of additional context that the agent could use to augment its understanding. I believed this so strongly that my company built an entire product around this concept. I used to tell folks that "session transcripts were the new oil," that they were more valuable than the code itself.
Other people have clearly had similar thoughts, which is why there are so many different tools to do session backed memory, including (of course) Claude Code itself.
I think the most common architecture is to do something like:
Store all transcripts across an organization in a DB
Put a vector search, an elastic search, or a SQL search layer in front of it. Ambitious teams will use all three. Maybe graphs will be involved.
Make this available to the agent using an MCP, or by exposing a cli with skills.
For us, this additional work doesn't seem to make a bit of difference. If anything, based on many months of testing with and without session search access, it may make the models worse.
One thing our team cares a lot about is coding artifacts. We don't really write code by hand anymore. In order to make PRs legible, we emphasize good commit messages, good pr messages, and comprehensive documentation. Every code change comes with extensive metadata that is committed alongside the code. When our agents do work on a piece of code, they are instructed to go look at the docs and the previous PRs.
In other words, the agent is already distilling all of the information that is valuable about a transcript, and storing it where it is needed and easily accessible. So when the agent uses a transcript search server, it ends up spending tokens reading things it already knows, while picking up all the stuff that the agent decided not to write down in the first place. Maybe, every now and then, there's some useful nugget of information in there. But most of the time, the agent is just looking at a pseudo nonsensical scratch pad and wasting precious tokens to do so.
The agents are also terrible at actually removing context, which is a critical capability for maintaining long term memory. I mean, across literal thousands of sessions, I've never seen it happen even once.
This is not a trait that can be removed with some clever prompt engineering. Agents don't have state, so they have to assume everything in their input context window is the ground truth. Every line of code, every existing bit of memory, every token is treated as an expression of intent -- even if that code or that memory was generated from a random decision made by some previous agent session, never reviewed or even understood by a human. This intent drift compounds the more the agent tries to autonomously build up a memory base.
As far as I am aware, there are zero coding benchmarks that assume the input data is corrupt. In fact, the models are penalized for assuming that the input data is wrong. This is partially an alignment issue too -- we don't want to have agents doing unintended things, and there isn't an easy way to thread the needle of “don't delete the codebase” and “do delete some of the input context.”
Since models can't actually garden their own memory, automatic memorization ends up in the same place: a load of garbage eating tokens, bloating bills, and degrading model quality.
Net net, I've become really bearish on tools that index and store and surface in session transcripts to an agent. The session transcript may be useful for team observability, but it won't make your agents better.
That doesn't mean that agents have no role in learning context over time. We use our internal nori bots to review everything that happened at the company each week across PRs, slack, drive, etc. And they then propose a set of changes to our built in nori skillsets, tagging the team in slack. These are all default rejected. In order to accept a change, you have to go in and actually look at the diff and make sure it fits the intent.
We accept less than 20% of these. Which means 80% of these “automatic” updates would've made the model worse. I can't imagine how much more unsustainable that would be if a multi-hundred person org were all saving these “updates” automatically all the time.
Thanks for reading! Subscribe for free to receive new posts and support my work.
1 6 Share Discussion about this post Comments Restacks B P 17h Liked by theahura So there’s an interesting question of “how broad was the business context that led to a line of code or a TODO.” And this can vary between businesses! Sometimes it’s sufficient to have a comment link back to a well maintained story tracker etc., or a motivating slack thread, which a modern LLM is incentivized to follow. But sometimes that motivating example is just part of a broad exploration of numerous experiences and insights, which might have gone through many colleagues exploring the space over time. A future iteration would over-index on the one example. We see this a lot in our curated marketplace, where a fix to an integration exists in the context of dozens of times listings were manually patched or overridden beforehand. And if I were fixing that fix, would I find everything that motivated the old code? It’s like Chesterton’s Fence meets the Dark Forest. I wouldn’t give up on your product just yet!
Reply Share Brian Schneider 20h Liked by theahura In this way it is most human.
Reply Share 4 more comments... Top Latest Discussions No posts
