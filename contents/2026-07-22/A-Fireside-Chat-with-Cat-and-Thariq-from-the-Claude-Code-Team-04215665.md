---
source: "https://simonwillison.net/2026/Jul/21/cat-and-thariq/"
hn_url: "https://news.ycombinator.com/item?id=49000570"
title: "A Fireside Chat with Cat and Thariq from the Claude Code Team"
article_title: "A Fireside Chat with Cat and Thariq from the Claude Code team"
author: "vismit2000"
captured_at: "2026-07-22T01:44:45Z"
capture_tool: "hn-digest"
hn_id: 49000570
score: 2
comments: 0
posted_at: "2026-07-22T01:12:48Z"
tags:
  - hacker-news
  - translated
---

# A Fireside Chat with Cat and Thariq from the Claude Code Team

- HN: [49000570](https://news.ycombinator.com/item?id=49000570)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jul/21/cat-and-thariq/)
- Score: 2
- Comments: 0
- Posted: 2026-07-22T01:12:48Z

## Translation

タイトル: クロード コード チームの Cat と Thariq との炉辺の会話
記事のタイトル: クロード コード チームの Cat と Thariq との炉端チャット
説明: 今月初め、私は Anthropic の Claude Code チームの Cat Wu 氏と Thariq Shihipar 氏とともに AI エンジニア ワールド フェアで炉端でのチャット セッションを主催しました。クロードについて話しました…

記事本文:
Claude Code チームの Cat と Thariq との炉端チャット
サイモン・ウィリソンのウェブログ
Claude Code チームの Cat と Thariq との炉端チャット
今月初め、私は Anthropic の Claude Code チームの Cat Wu 氏と Thariq Shihipar 氏とともに AI エンジニア ワールド フェアで炉端でのチャット セッションを主催しました。クロード コード、クロード タグ、寓話、コーディング エージェントのセキュリティ、評価、ツール設計、そして Anthropic がこれらのツール自体をどのように使用するかについて話しました。
セッションの完全なビデオは YouTube で公開されています。以下はトランスクリプトの編集されたコピーで、追加のリンクと私自身の太字のハイライトが含まれています。
ビデオを見たり、トランスクリプト全体を読みたくない場合は、いくつかのトップレベルのメモを以下に示します。
Claude Tag (Claude の新しい共同 Slack 統合) により、現在、Claude Code チームの製品エンジニアリング PR の 65% が採用されています。
Claude Code は、最初に Anthropic 従業員に機能を出荷し、そのコホートでのユーザー維持を実証する機能のみを出荷します。
クロード コードに対する重要な変更は依然として手動でレビューされていますが、チームは製品の「外側の層」については自動化されたコード レビューにますます依存しています。
システム プロンプトに例を追加することは、Fable 5 や Opus 4.8 のようなモデルではもはやベスト プラクティスではありません。最近、Claude Code システム プロンプトのサイズが 80% 縮小されました。
同様に、「X を実行しない、Y を実行しない」というリストは、最新モデルの結果の品質を低下させる可能性があります。
Anthropic 内でのドッグフーディングは「アリ フーディング」と呼ばれます。
Anthropic は自社の自動モードを心から信じており、それがクロード タグを可能にするテクノロジーであると考えています。
Thariq 氏は、コーディング エージェントによるディープ ブルーの影響を「より野心的に」取り組むことで相殺するようアドバイスしています。
Fable はビデオ編集に優れており、Thariq はそれを使って自社のローンチビデオを編集しました。
人間の労働文化 (

公開 Slack チャンネルでの Claude Tag の使用方法からもわかるように、公開での公開が成功の鍵となります。
Simon: Claude Code は昨年 2 月に公開されました。まだ 1 年半も経っていませんが、もともとは Claude Sonnet 3.7 のリリースに関する単なる箇条書きでした。実際に機能するコーディング エージェントが登場したことで、この 1 年で日常業務はどのように変化しましたか?
Cat: 最初に Claude Code と Sonnet 3.7 を発表したときのことを覚えています。それにタスクを与えると、それが実行しようとするあらゆる小さなことを注意深く監視する必要がありました。私はすべての許可プロンプトを注意深く読みます。私はよく「いいえ」と言いました。「いいえ、いいえ、いいえ、このファイルをチェックしましたか?」そのファイルをチェックしましたか?そして今では、どの世代のモデルでも信じられないほど素晴らしいものになっています。私たち全員が一歩下がって、より多くの単純な実装を Claude に委任する機会を得たような気がします。クロード コードで多くのエクスペリエンスを実装できることがわかったので、ユーザーに提供すべき適切なエクスペリエンスは何か? など、より創造的な作業について考える時間が大幅に解放されました。そして今回の Fable では、まったく異なるステップ変更の改善が行われています。多くのユースケースで、実際に Fable を使用して大量の機能をワンショットで実行できることがわかりました。
タリク: クロード・コードについて最初に受け取ったテキストを覚えています。私の親友の一人は、「クロード・コードを試してみるべきだ」と言ってきました。 Opus 4 がリリースされた頃で、試してみたら、「ああ、くそー。今から Anthropic で働かなければいけない」と思いました。それが Opus 4 でした。素晴らしいモデルですが、許可のプロンプトを読んでいたのです。私たちの記憶喪失の多さはちょっとクレイジーです。ああ、自動モードは常にここにあったのではないかと思います。 「はい」と「許可」を押した覚えすらありません。私にとって、

私が自分自身を追い込もうとしている大きなことは、これまでよりも質の高い仕事をしなければならないということです。出力は信じられないほど高品質です。私はこれを使ってビデオを編集してきましたが、数時間以内にブランド チームの非常に厳しい要求に応えなければならない、そうしないと無理だ、という感じです。それが私が Fable でシフトしようとしている方法です。これまでにやったことのない最高の仕事を、これまでよりも早く実現しようとしています。
Simon: 1 年前には真実だった、この新しい世界ではもう通用しないと思われる従来のソフトウェア エンジニアリングの一部は何ですか?
Cat: エンジニアのスキルセットで私たちが見ている最大の変化の 1 つは、2 年前、プロダクト マネージャーが多くの顧客と話をしに行き、PRD の部門横断的なチームと 6 か月間にわたって連携し、コードの最初の行が書き込まれる前に、これをどのように実装するかについての詳細な仕様を書くのが非常に一般的でした。今では事態は完全に逆になってしまいました。多くのエンジニアに対して、私がこの場にいる人たちに与えるのは、何を構築すべきかについて、ビジネス感覚と製品感覚をもっと養うことです。なぜなら、アイデアを思いついてからそれを構築するまでのタイムラインが非常に短くなり、6 ～ 12 か月から、場合によっては 1 週間にまで短縮されるからです。つまり、私たち全員が、何を構築する価値があるのか​​、私たちが取り組んでいるビジネスに実際に影響を与えるものは何なのかをよりよく見極める必要があるということです。つまり、製品の味とビジネスセンスの価値が増加し、ほとんどの製品領域での実行の価値が少し低くなります。もちろん、インフラの場合は、すべての詳細が正しいことを確認することが依然として非常に重視されています。
Thariq: 私にとって、リライトがうまくなったということです。
サイモン: 最悪のことでも、今では大丈夫です!
タリク：その通りです。ミシカのすべて

l 人月のもの — 決して書き直さない — 私は今、書き直すことに賛成です。優れたテスト スイートを持っている場合は、書き換えによって実際に優れたテスト スイートであることを確認する必要があると思いますが、コードベースの分岐部分をすべて知っている人はいないため、コードベースは仕様であり、おそらくそれが仕様の唯一のコピーであるということを過小評価していると思います。これを成果物として抽出して抽出したり、他のバージョンを作成したりできます。 Bun を Rust で書き直したところ、うまく機能しました。現在、私にとっては有効です。
Simon: まだ Bun-in-Rust でクロード コードを出荷していないんですよね?
(実際、Anthropic は 6 月 17 日に Bun-in-Rust 上のクロード コードを全員に配布し始めたようです。)
Simon: 最近のもう 1 つの大きなリリースは、Claude Tag です。少なくとも残りの私たちにとっては、1 週間前にリリースされたものです。 Anthropic ではエンジニア以外の人が多く使用していると聞いています。エンジニア以外の人はクロード・タグを使ってどのようなことをやっているのでしょうか？
Cat: クロード タグは、チームのコラボレーション ツールに存在するクロードです。先週、Slack 内でこれを開始しました。 Claude Tag が違うのは、デフォルトでマルチプレイヤーであることです。 Claude Tag を Slack チャネルに追加すると、自分が参加したり、チームメイトが参加したり、一緒に PR に協力したりできるようになります。もう 1 つの大きな違いは、事後対応ではなくプロアクティブであることです。 Claude Tag に「このチャネル内のすべてのバグ レポートを監視して、修正するために PR を投稿し、コードベースのこの部分を最後に触ったエンジニアにタグ付けしてください」と指示すると、手動でタグ付けする必要がなく、チャネルの存続期間中、それが実行されます。そして 3 番目の大きな変化は、この .html にチームのメモリを追加したことです。チャンネル内であなたの好みを Claude Tag に伝えると、今後の投稿ごとにその設定が記憶されます。あなたがいつもそうだったら

停止をデバッグするのではなく、警告をデバッグしたくない場合は、チャネル内で自然言語でその旨を伝えるだけで、あなたとチームの他の全員のためにそれを記憶します。
社内では、Claude Tag は Claude Code の進化版であると考えています。私たちはこれを社内の働き方の大きな変化だと考えています。現在、Claude Tag は当社の製品技術 PR の 65% を占めています。
サイモン: Anthropic 全体に対してですか、それともクロード コードだけに対してですか?
Cat: これは当社の製品エンジニアリング チームに限った話です。現在、当社の内部バージョンの Claude Tag が当社の製品 PR の 65% を占めています。そしてこれは大きな変化です。これは当社の PR の 50% 以上です。人々がクロード コードとクロード タグで作業を分割している様子は次のとおりです。エージェントと対話的に反復処理を行う場合、最も複雑なタスクには依然としてクロード コードが最適です。ただし、Claude Tag は、ユーザーに代わって積極的に機能させるのに優れているため、開発中の機能に関して発生したすべてのバグ レポートに対して手動で Claude Code を開始する必要はもうありません。
Thariq: コーディング以外の場合については、たとえば、この講演の前に、私たちはクロード・タグに「ねえ、Fable はいつリリースされるの?」と尋ねました。発表と確実に一致させたかったのです。クロード・タグは私たちの Slack を検索して、誰が何を言ったかを調べます。あなたの会社の検索エンジンとして、これは非常に価値があります。製品に関するすべてのコンテキストが含まれているため、メトリクス関連の質問をすることができます。多くの場合、意思決定を行う際に、メトリクスが示す内容を知らせてもらいたいため、イベント ストアに接続します。私たちのマーケティング チームが「この機能について教えてください」というようなことをしているのを見てきました。彼らはプログラマーではありませんが、クロードはプログラマーです。コードベースをクローンして、「これはこの機能です、これはどのように見えるか、これは私がその機能を使用している記録です」と言うことができます。それはね

さまざまなことを実現できますが、それを理解するのはまだ早い段階だと思います。
Simon: コーディング エージェントに関して私が抱えていた問題の 1 つは、個人としての使い方は理解できても、チーム環境での使い方がよくわからないことです。 Claude Tag が、この問題に関するチームの共同作業層に対する現在の答えのようですね。
猫：その通りです。そして現在、私たちのセッションの大部分は実際にマルチプレイヤーです。おそらく私は、「この新機能を Cowork に実装すべきだと思います」と言って、最初のパスを実行するために Claude Tag でタグ付けします。次に、Claude Tag に「最終実装の記録を共有してください」と指示し、デザインにタグを付けて確認します。彼らはそれを小突いて、それをエンジニアに渡してゴールラインまで運び、突き出すようにします。とても流動的な経験でした。私たちは同じセッションを操作するための社会的力学がどのようなものであるかをまだ解明しようとしていますが、人々は他の人がそれをどのように使用するかを観察し、それらの社会規範に従っているだけであることがわかりました。私たちにとって、Claude Tag をチームに統合することは非常に直感的でした。
Thariq: これは人に教えるのにも、またスロップを減らすのにも最適です。なぜなら、あなたがクロードを一緒に使っているのをみんなが見ているという事実は、クロードの使い方もレベルアップするからです。
これを見て、Midjourney が Discord チャンネルで公の場でプロンプトを強制することで、人々に高度なイメージ プロンプトを教えるという課題をどのように解決したかを思い出しました。
私自身も非常に難しいと感じているのは、実際に機能を構築するコストが大幅に下がった今、機能をいつリリースする価値があるかを知ることです。
Simon: エンジニアリング全体の中で最も難しい問題、つまり優先順位付けにどのように対処しますか?機能を構築するのに非常に時間がかかる場合、どの機能を構築して出荷する価値があるかをどのように判断しますか?

今はもっと安くなりましたか？
猫：これは難しいことだよ。それにはいくつかのアプローチ方法があります。 1 つは、私たちが毎日自社の製品をドッグフードしていることです。製品で実現したいことがあるのにできない場合は、別の解決策を見つける代わりに、そのケースをサポートできるように製品を修正します。私たちの社内には非常に厳しいドッグフーディング文化があります。私たちの製品を世界中の人々と共有する前に、Anthropic 内の全員と、それについて非常に正直なフィードバックをくれる初期の顧客 (残忍であればあるほど良い) と共有し、人々がそれを気に入るまで繰り返します。私たちは、アクティブ ユーザーの数と、その機能を世界と共有する前に保持する必要がある機能の保持量を示す内部バーを設けています。このバーは非常に明確であるため、すべてのエンジニアは自分が何を達成しようとしているのかを知っています。これにより、私たちの磨きもレベルアップすると思います。なぜなら、機能が洗練されていないと、ユーザーは離れてしまうため、その機能をリリースすべきではないからです。
機能を出荷するかどうかを決定するために内部ユーザー維持を使用することは、私にとって非常に理にかなっています。
Simon: 驚いた機能の例はありますか?それを展開したところ、エンゲージメントは予想外でした。出荷される可能性は低いものが実際の製品になりました。
猫：あるよ。私たちのチームの多くの人はリモートコントロールが大好きです。リモート コントロールを使用すると、モバイル デバイスを使用したり、Web ブラウザでクロードを使用したりできます。

[切り捨てられた]

## Original Extract

Earlier this month I hosted a fireside chat session at the AI Engineer World’s Fair with Cat Wu and Thariq Shihipar from Anthropic’s Claude Code team. We talked about Claude …

A Fireside Chat with Cat and Thariq from the Claude Code team
Simon Willison’s Weblog
A Fireside Chat with Cat and Thariq from the Claude Code team
Earlier this month I hosted a fireside chat session at the AI Engineer World’s Fair with Cat Wu and Thariq Shihipar from Anthropic’s Claude Code team. We talked about Claude Code, Claude Tag, Fable, coding agent security, evals, tool design, and how Anthropic use these tools themselves.
The full video of the session is now available on YouTube . Below is an edited copy of the transcript, with extra links and my own bolded highlights.
A few top-level notes if you don’t want to watch the video or wade through the whole transcript:
Claude Tag (Claude’s new collaborative Slack integration) now lands 65% of the product engineering PRs for the Claude Code team.
Claude Code ships features to Anthropic employees first, and only ships the features that demonstrate user retention with that cohort
Critical changes to Claude Code are still reviewed manually, but the team increasingly relies on automated code review for the “outer layers” of the product.
Adding examples to a system prompt is no longer best practice for models like Fable 5 or even Opus 4.8. The Claude Code system prompt recently reduced in size by 80% .
Likewise, lists of " don’t do X and don’t do Y " can reduce the quality of results from the latest models.
Dogfooding inside Anthropic is called " ant fooding ".
Anthropic really believe in their auto mode , and see that as an enabling technology for Claude Tag.
Thariq advises offsetting coding-agent-induced Deep Blue by " being more ambitious " with the work you take on.
Fable is competent at editing video , and Thariq used it to edit its own launch video.
Anthropic’s culture of working (internally) in public is key to their success, as demonstrated by the way they use Claude Tag in their public Slack Channels.
Simon: Claude Code came out in February of last year — it’s under a year and a half old, and it was originally just a bullet point on the Claude Sonnet 3.7 launch . How has what you do on a day-to-day basis changed in the past year , now that we have these coding agents that actually work for us?
Cat: I remember when we first came out with Claude Code and Sonnet 3.7, you would give it a task and you would have to closely monitor every single little thing it tried to do. I would read every permission prompt extremely carefully. I would frequently say no — no, no, no, did you check this file? Did you check that file? And now it’s been incredible with every model generation. I feel like we’ve all gotten a chance to take a step back and delegate a lot more of the menial implementation to Claude . It’s freed up a lot of our time to think about more creative work, like: what is the right experience that we should be providing to our users, now that we know Claude Code can implement a lot of it? And now with Fable it’s a totally different step change improvement. We see for a lot of our use cases that you can actually one-shot a ton of features with Fable now .
Thariq: I remember the first text I got about Claude Code. One of my best friends was like, “You need to go try Claude Code.” It was about when Opus 4 came out, and I tried it and I was like, “Oh, shit. I need to work at Anthropic now.” And that was Opus 4 — great model, but you were reading permission prompts. It’s kind of crazy how much amnesia we have, where I’m like, oh, auto mode has always been here, right? I don’t even remember pressing yes and allow. For me, the big thing I’m trying to push myself on is that we have to do higher quality work than we’ve ever done before . The outputs are incredibly high quality. I’ve been using it to edit videos a bunch , and I’m like, okay, it has to meet the very exacting demands of our brand team in a couple of hours or we just can’t do it. That’s how I’m trying to shift with Fable: the best work we’ve ever done, faster than we’ve ever done it before .
Simon: What’s a piece of conventional software engineering that was true a year ago that you don’t think holds anymore in this new world?
Cat: One of the biggest shifts we’re seeing in the eng skill set: two years ago it was pretty typical for a product manager to go talk to a bunch of customers, align over the course of six months with cross-functional teams on some PRD, and write a thorough spec on exactly how we’ll implement this before the first line of code gets written. Now things are completely turned the opposite way. For a lot of engineers, the push I would give to folks in the room is to develop more of your business sense and product sense on what it is we should build , because the timeline between having an idea and building it is so much shorter — it’s down from six to twelve months to maybe even a week. That means all of us need to have better taste on what is worth building, what will actually inflect the businesses we’re working on. So it’s an increase in value on product taste and business sense , and a bit lower on execution in most product domains. Of course, for infra there’s still a very heavy emphasis on making sure all the details are right.
Thariq: For me, it’s that rewrites are now good .
Simon: The worst thing you could do is now actually fine!
Thariq: Exactly. All the Mythical Man-Month stuff — never rewrite — I’m pro-rewriting now. If you have a good test suite — and I think the rewrite actually forces you to make sure you have a good test suite — but I think what people undercount is that a codebase is a spec, and maybe it’s the only copy of the spec that you have , because no one knows every branching part of the codebase. You can take this as an artifact and distill it or create other versions of it. We rewrote Bun in Rust and it works great — it’s live for me right now.
Simon: You’re not shipping Claude Code on Bun-in-Rust yet, right?
(Actually it looks like Anthropic started shipping Claude Code on Bun-in-Rust to everyone on June 17th .)
Simon: The other big launch recently was Claude Tag — that’s what, a week old now, at least for the rest of us. I understand it’s being used at Anthropic by non-engineers a great deal. What kind of things are non-engineers doing with Claude Tag?
Cat: Claude Tag is a Claude that lives in your team’s collaboration tools. We launched it last week within Slack. The thing that’s different about Claude Tag is it’s multiplayer by default . Once you add Claude Tag to a Slack channel, you can chime in, your teammates can chime in, and you can collaborate together on the PR. The other big difference is that it’s proactive instead of reactive. You can tell Claude Tag, “Hey, monitor every bug report in this channel, put up a PR to fix it, and tag the engineer who last touched this part of the codebase,” and it’ll do it for the lifetime of the channel without you having to manually tag it in. And the third big shift is that we’ve added team memory into this . If you tell Claude Tag your preferences in the channel, it’ll remember them for every future post. If you always want it to debug outages but you don’t want it to debug warnings, just tell it that in natural language in the channel and it’ll remember it for you and everyone else on your team.
Internally, we see Claude Tag as the evolution of Claude Code. We see this as a large shift in how we work internally. Claude Tag currently lands 65% of our product eng PRs.
Simon: For all of Anthropic, or just for Claude Code?
Cat: This is just for our product engineering team — our internal version of Claude Tag lands 65% of our product PRs right now . And this is a huge shift; this is more than 50% of our PRs. The way we see people split work between Claude Code and Claude Tag is: Claude Code is still the best place for your most complex tasks, when you’re interactively iterating with the agent. But Claude Tag is great for having it work proactively on your behalf , so you no longer need to manually kick off Claude Code for all the bug reports that come up for features you’re working on.
Thariq: And for non-coding cases: for example, before this talk we asked Claude Tag, “Hey, when is Fable releasing?” We wanted to make sure we’d line it up with the announcement. Claude Tag would search our Slack and look at who’s been saying what. As a search engine for your company, it’s really valuable. It has all the context for your product, so you can ask it metrics-related questions — often when you’re making decisions you want them informed by what the metrics say, so you hook it up to your event store. I’ve seen our marketing team do things like, "Hey, tell me about this feature." They’re not programmers, but Claude is a programmer — it can clone the codebase and say, "This is the feature, this is what it looks like, this is a recording of me using the feature ." It enables a whole wide variety of things, and I think we’re still early in figuring that out.
Simon: One of the problems I’ve had with coding agents is that I get how to use them as an individual, but I’m not really clear on how to use them in a team environment. It sounds like Claude Tag is your current answer to that team collaborative layer for this stuff.
Cat: Exactly. And a large percentage of our sessions are actually multiplayer right now. Maybe I say, “Hey, I think we should implement this new feature in Cowork,” and I’ll tag in Claude Tag to do a first pass at it. Then I’ll tell Claude Tag, “Share a recording of your final implementation,” and I’ll tag in design to take a look. They’ll nudge it, then pass it on to eng to take it to the finish line and get it out to prod. It’s been this very fluid experience. We’re still trying to iron out what the social dynamics are for steering the same session , but we’ve found that people just observe how others use it and follow those social norms — it’s been pretty intuitive for us to integrate Claude Tag into our teams.
Thariq: It’s great for teaching people, and also for reducing slop, because the fact that everyone is seeing you use Claude together sort of levels up how you use Claude as well .
This reminded me of how Midjourney solved the challenge of teaching people advanced image prompting by enforcing prompting in public in their Discord channels.
Something I’ve found really hard myself is knowing when a feature is worth shipping now that the cost of actually building features has dropped so much.
Simon: How do you deal with the hardest problem in all of engineering — prioritization? How do you decide which features are worth building and shipping when building a feature is so much more inexpensive now?
Cat: This is the hard thing. There are a few ways we approach it. One is we dogfood our products every single day. Whenever there’s something we want to be able to do in our products that we’re not able to, instead of finding a different solution we fix our product so it can support that case. We have a very heavy dogfooding culture internally. Before we share our products with everyone in the world, we share them with everyone within Anthropic, and with some early customers who give us very honest feedback about it — the more brutal the better — and we iterate until people love it. We have an internal bar for the number of active users and the amount of retention a feature has to have before we share it with the world. Because this bar is very clear, every engineer knows what they’re trying to hit. I think this also levels up our polish, because if the feature isn’t polished, people will churn — and then we shouldn’t ship that feature.
Using internal user-retention to decide if a feature should ship makes a whole lot of sense to me.
Simon: Do you have an example of a feature which surprised you? You rolled it out and the engagement was off the charts — something unlikely to be shipped that turned into a real product thing.
Cat: I do have one. A lot of folks on our team love remote control . Remote control lets you use your mobile device, or Claude in the web browser,

[truncated]
