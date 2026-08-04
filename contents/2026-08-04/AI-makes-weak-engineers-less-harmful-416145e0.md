---
source: "https://www.seangoedecke.com/ai-makes-weak-engineers-less-harmful/"
hn_url: "https://news.ycombinator.com/item?id=49173796"
title: "AI makes weak engineers less harmful"
article_title: "AI makes weak engineers less harmful"
author: "schrodinger"
captured_at: "2026-08-04T20:19:43Z"
capture_tool: "hn-digest"
hn_id: 49173796
score: 2
comments: 0
posted_at: "2026-08-04T19:33:20Z"
tags:
  - hacker-news
  - translated
---

# AI makes weak engineers less harmful

- HN: [49173796](https://news.ycombinator.com/item?id=49173796)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/ai-makes-weak-engineers-less-harmful/)
- Score: 2
- Comments: 0
- Posted: 2026-08-04T19:33:20Z

## Translation

タイトル: AI により弱いエンジニアの害が軽減される

記事本文:
AI により弱いエンジニアの害が軽減される ショーン・ゴーデッケ
AIにより弱いエンジニアの害が軽減される
他の種類の謎解きと同様に、ソフトウェア エンジニアリングの能力は非常にヘビーテールです。最も強いエンジニアは平均よりもはるかに有益な成果を生み出しますが、最も弱いエンジニアは積極的にネットネガティブになることがよくあります。彼らはプロジェクトを進める代わりに、同僚が時間を費やして解決しなければならない問題を作成します。だからこそ、多くのテクノロジー企業が、より平均的なエ​​ンジニアからなる大規模なチームではなく、ばかばかしいほど高給取りの小規模なチームを構築しようとしているのです。また、これが今のところ勝利戦略であるように見える理由もここにあります。
大手テクノロジー企業で有能であるためには、多くの場合、この現象を管理することが重要です。成功させたいプロジェクトに最も有能な人材が着地し、最も有能でない人材は排除されるように物事を調整しようとするのです 1 。たとえば、あなたがプロジェクトの技術リーダーである場合、最も重要な部分が台無しにならない人の手に渡っていることを多かれ少なかれ確認する必要があります (作業を直接割り当てるか、または心配しているエンジニアの「肩に座る」ことができる人を確保するかにかかわらず)。
クロード・コードはこれを変えました。 Frontier LLM には、強いエンジニアのようなセンスやシステムへの精通性はありませんが、弱いエンジニアのレベルを確実に引き上げています。絶対に機能しない、またはすぐに問題を引き起こす可能性のあるプル リクエストを取得するのではなく、標準的な LLM プル リクエストを受け取ることになります。ある意味間違っていて、他の意味では不可解ですが、少なくとも行ごとのレベルでは機能しており、コードベースの知識がない人が指摘できるほど明らかに間違っているわけではありません。それは大きな進歩です！
これは自分でも試してみることができます。仕事中に故意にミスをしようとした場合

コーディング エージェントを使用すると、エージェントが多くの明らかなエラー (ユーザー固有ではないキーを使用したユーザー データのキャッシュ、終了しない可能性のある無限ループの作成、開いているファイルの漏洩など) に対して強力に抵抗することがわかります。もちろん、エージェントは依然として微妙なエラー、特にコードベースの他の部分を理解する必要があるエラーを見逃します。
最も効率の悪いエンジニアと作業することは、Slack 経由で通信する Claude Opus または Codex インスタンスと作業することに似ている場合があります。場合によっては、文字通り、同僚があなたのメッセージをクロード コードに貼り付け、その応答をあなたに貼り付けているだけであることがあります。これは面倒ですが、この種のエンジニアと直接仕事をするよりもずっと良い経験になります。結局のところ、あなたはおそらくすでに多数の LLM インスタンスを使用しているでしょう。 Slack インターフェイスは理想的ではありません。Claude Code を直接使用する場合とは異なり、応答までに数時間または数日かかる場合があり、エージェントの思考プロセスを可視化することはできません。しかし、それでもある程度は役に立ちます。問題に投入されるコンピューティングは、少ないよりも多い方が良いです。
もちろん、これは問題のエンジニアにとって素晴らしい状況ではなく、ほぼ確実に、自分で (悪い) 決定を下した場合よりも学習量が少なくなります。これは、人間の給与を支払い、Copilot のサブスクリプションを取得している (おそらく料金も支払っている) 会社にとっても悪い状況です 3 。 AI がエンジニアにどのような価値を付加しているのかを把握しようとする現在の取り組みの後、エンジニアが AI にどのような価値を付加しているのかを把握しようとする取り組みが行われるのではないかと思います。そして、あまり付加していないエンジニアは職を失うかもしれません。
通常のクロードと話すように、Slack 上でクロードと話すことはできません。 LLM を乱暴に扱う傾向がある場合 (侮辱したり、単純に悪口を言ったりする場合)

とても素っ気ない）の場合は、コミュニケーション スタイルを変える必要があります。結局のところ、実際に LLM とやり取りしている場合でも、人間はメッセージを読むことになります。失礼なことは意味がありません。しかし、私のように、モデル 4 に「よろしくお願いします」と言うのであれば、LLM を使用している同僚を単なる別の Copilot ウィンドウまたは Codex タブとして扱うことができます。彼らを無自覚の妨害者として扱わなければならないよりは、はるかに良いでしょう。
ネットネガティブなエンジニア全員がこのような AI ツールを使用しているわけではありません。多くの人は、優れたソフトウェアの構築方法について自分の間違った意見を強く信じていたり、AI 全般に不信感を持っていたり、LLM に大きく依存することは改善の良い方法ではないと信じていたりしています 5 。しかし、このような AI ツールを使用する優秀なエンジニアはいません。たとえ怠け者だったりずさんだったりしたとしても、有能なエンジニアは、AI によって生成された明らかなエラーを見つけるのに十分な基本的なセンスを備えています。したがって、エンジニア 6 がクロード コードの薄いラッパーになるという現象は、これが作業成果物の改善となる種類のエンジニアに限定されます。
編集: これは最終的に YouTube の Theo ビデオのトピックになりました。
もっと慈悲深いことを言えば、「最も能力の低い」エンジニアの多くは、自分の快適ゾーンから抜け出したばかりで、適切な環境下では問題なく、あるいは優れた能力を発揮することさえあります（ただし、私の考えでは、最高のエンジニアはさまざまな環境で良い仕事をすることができます）。また、私は現在、多くの無能な人たちとは仕事をしていません。これらの多くは、過去の経験や業界の他のエンジニアとの会話に基づいています。
あなたのマネージャーも同じことをしているので、これはマネーボールのように感じることがあるかもしれません。あなたは、上司が別のことを率いるために彼らを引き抜くほど知名度が高くなくても、あなたの勝利を助けるのに十分な力を持った過小評価されている人材を見つけようとしているのです。
ネットネガティブに対してお金を払うよりは、何も支払わない方が良いと思います

出力されましたが、まだ良くないようです。
実際、これがクロード作品 4.7 の正しい持ち方だと思います。
これは本当ですか？ LLM に依存することは、ほとんどのエンジニアにとって改善のための良い方法ではないと思いますが、LLM の出力が自分の出力よりも一貫して優れている場合は、状況が異なる可能性があります。 LLM のどこが優れているかに注意を払っている限り、それは実際に良い学習方法になる可能性があります。
非エンジニアがこの罠に陥った経験 (または逸話) は私にはあまりありませんが、この投稿を読んで、それはもっと悪いかもしれないと確信しました。
この投稿を気に入っていただけた場合は、私の新しい投稿に関する最新情報を電子メールで購読するか、 Hacker News で共有することを検討してください。
これは、この投稿とタグを共有する関連投稿のプレビューです。
エンジニアは単純なコードを書くことで昇進する
ソフトウェア エンジニアの間では、複雑すぎて保守不可能なコードを書くことが雇用を安定させる手段になるというのがよく言われるジョークです。結局のところ、システムに携わることができるのがあなただけであれば、彼らはあなたを解雇することはできません。これに関連して、「単純さのために昇進する人はいない」という見方もあります。つまり、複雑すぎるくだらない仕事を提供するエンジニアは昇進するのです。なぜなら、技術系ではないマネージャーにとって、彼らの仕事はより印象的に見えるからです。
続きを読む...
購読する │ About │ ポッドキャスト │ 人気 │ タグ │ RSS

## Original Extract

AI makes weak engineers less harmful sean goedecke
AI makes weak engineers less harmful
Like other kinds of puzzle-solving, software engineering ability is strongly heavy-tailed. The strongest engineers produce way more useful output than the average, and the weakest engineers often are actively net-negative: instead of moving projects along, they create problems that their colleagues have to spend time solving. That’s why many tech companies try to build a small, ludicrously well-paid team instead of a large team of more average engineers, and why so far this seems to be a winning strategy.
Being effective in a large tech company is often about managing this phenomenon: trying to arrange things so that the most competent people land on projects you want to succeed, and the least competent are shunted out of the way 1 . For instance, if you’re technical lead on a project, you more or less have to ensure 2 that the most critical pieces are in the hands of people who won’t screw them up (whether by directly assigning the work, or by making sure someone can “sit on the shoulder” of the engineer who you’re worried about).
Claude Code changed this. Frontier LLMs don’t have the taste or the system familiarity of a strong engineer, but they have absolutely raised the floor for weak engineers. Instead of getting a pull request that could never possibly work or would cause immediate problems, the worst you’ll now see is a standard LLM pull request: wrong in some ways, baffling in others, but at least functional on the line-by-line level and not so obviously incorrect that someone with no knowledge of the codebase could point it out. That is a huge improvement!
You can try this out yourself. If you attempt to deliberately make mistakes while working with a coding agent, you’ll find that the agent pushes back hard against many obvious errors (i.e. caching user data with a non-user-specific key, writing an infinite loop that might never terminate, or leaking open files). Of course, the agent will still miss subtle errors, particularly ones that require understanding other parts of the codebase.
Working with the least effective engineers is now sometimes like working with a Claude Opus or Codex instance that you communicate with over Slack. Occasionally it’s literally that: your colleague is simply pasting your messages into Claude Code and pasting you the response. This is annoying, but it’s a much better experience than working with this kind of engineer directly. After all, you probably already work with a bunch of LLM instances. The Slack interface is not ideal — unlike using Claude Code directly, you sometimes wait hours or days for a response, and you don’t get visibility into the agent’s thought processes — but it’s still helpful on the margin. More compute being thrown at your problem is better than less.
Of course, this isn’t a great state of affairs for the engineer in question, who is almost certainly learning less than if they were making their own (bad) decisions. It’s also a bad state of affairs for the company, who is paying a human salary and getting a Copilot subscription (which they’re likely also paying for) 3 . After the current push to figure out what value AI is adding to engineers, I suspect there will be a push to figure out what value engineers are adding to AI , and the engineers who aren’t adding much may find themselves out of a job.
You can’t talk to Claude-over-Slack like you’d talk to normal Claude. If you tend to handle LLMs roughly (insulting them, or just being very curt), you’ll have to change your communication style. A human is going to read your messages, after all, even if you’re really interacting with a LLM. There’s no point being rude. But if, like me, you say please-and-thank-you to the models 4 , you can treat your LLM-using coworker as just another Copilot window or Codex tab. It’s far better than having to treat them as an unwitting saboteur.
Not all net-negative engineers use AI tools like this. Many are strongly convinced in their own wrong opinions about how to build good software, or mistrust AI in general, or believe that relying heavily on LLMs is not a good way to improve 5 . But no strong engineers use AI tools like this. Even when they’re being lazy or sloppy, a capable engineer will have enough baseline taste to catch obvious AI-generated errors. So the phenomenon of engineers 6 becoming thin wrappers around Claude Code is limited to the kind of engineers for whom this is an improvement in their work product.
edit: this ended up being the topic of a Theo video on YouTube.
More charitably: many “least competent” engineers are just out of their comfort zone, and can be fine or even excel under the right circumstances (though in my view the best engineers are able to do good work in a wide variety of environments). Also, I don’t currently work with a lot of incompetent people. Much of this is based on past experience or talking to other engineers in the industry.
Since your managers are doing the same thing, this can sometimes feel like Moneyball: you’re trying to identify underappreciated talent who are strong enough to help you win without being so high-profile that your boss poaches them to lead something else.
I suppose it’s better to pay for nothing than to pay for net-negative output, but it still doesn’t seem good .
I think this is actually the right way to hold Claude Opus 4.7.
Is this true? I think relying on LLMs is not a great way for most engineers to improve, but if LLM output is consistently better than your own, it might be different. So long as you’re paying attention to where the LLM does better, it could actually be a good way to learn.
I don’t have as much experience (or anecdotes) about non-engineers falling into this trap, but this post has convinced me that it might be worse.
If you liked this post, consider subscribing to email updates about my new posts, or sharing it on Hacker News .
Here's a preview of a related post that shares tags with this one.
Engineers do get promoted for writing simple code
It’s a popular joke among software engineers that writing overcomplicated, unmaintainable code is a pathway to job security. After all, if you’re the only person who can work on a system, they can’t fire you. There’s a related take that “nobody gets promoted for simplicity” : in other words, engineers who deliver overcomplicated crap will be promoted, because their work looks more impressive to non-technical managers.
Continue reading...
subscribe │ about │ podcasts │ popular │ tags │ rss
