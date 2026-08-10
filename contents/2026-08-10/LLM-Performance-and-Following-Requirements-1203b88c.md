---
source: "https://492ab3fb.danpalmer-me.pages.dev/2026-08-09-llms-and-requirements/"
hn_url: "https://news.ycombinator.com/item?id=49246369"
title: "LLM Performance and Following Requirements"
article_title: "LLM Performance and Following Requirements – Dan Palmer"
author: "speckx"
captured_at: "2026-08-10T17:44:50Z"
capture_tool: "hn-digest"
hn_id: 49246369
score: 2
comments: 0
posted_at: "2026-08-10T16:53:39Z"
tags:
  - hacker-news
  - translated
---

# LLM Performance and Following Requirements

- HN: [49246369](https://news.ycombinator.com/item?id=49246369)
- Source: [492ab3fb.danpalmer-me.pages.dev](https://492ab3fb.danpalmer-me.pages.dev/2026-08-09-llms-and-requirements/)
- Score: 2
- Comments: 0
- Posted: 2026-08-10T16:53:39Z

## Translation

タイトル: LLM のパフォーマンスと次の要件
記事タイトル: LLM のパフォーマンスと次の要件 – Dan Palmer

記事本文:
ブログ
プロジェクト
CV LLM のパフォーマンスと次の要件
さまざまな人々が LLM に関して経験していることの間には、乖離があります。ある者はそれらを変革的だと考えるが、他の者は彼らが無能であると考える。これは部分的にはスキルによるものであり、部分的にはテクノロジーに対する嫌悪感によるものであり、社会的影響もありますが、さらに問題があると思います。この投稿は、AI 擁護や反 AI を目指すものではありません。社会問題は現実だと思います。LLM が誇大広告のすべてに応えられるとは思いませんが、LLM を使用すると非常に生産的になる可能性があると思います。
それで何が問題なのでしょうか？それは、LLM が要件をどのように満たすか、そしておそらく人間が要件をどのように表現するかに起因します。 LLM がうまく機能する部分と不十分な部分を確認して、いくつかのパターンを観察してみましょう。
これは、LLM がこれらすべてに対して完璧に機能するということではなく、むしろ、LLM のほうが一般的により適切に機能し、必要な指導が少なくなるということです。
バイブコーディング – これは、すべてのプログラミングを LLM に延期し、コードを書かず、コードをレビューせず、ワンショットまたは数ショットの実装に焦点を当てることを意味します。 LLM がこの点で非常に優れていることは明らかであり、それを中心に製品のエコシステムが出現しています。
セキュリティ テスト – LLM は、コードやシステムの脆弱性を見つけるのにも優れているようです。軌跡が長いほど、LLM は脆弱性を発見する能力が高くなります。
短い創造的な作品 – LLM は詩や作詞作曲などを得意とします。芸術的価値については別の問題があります 1 が、LLM が完全に合理的な詩、歌、短い形式の創造的な文章を作成し、既存の文書の内容や性質を変更できることは否定できません。
LLM がうまく機能する場合と同様に、これは LLM が次の領域で失敗するということではなく、むしろ苦労したり、(品質、時間、トークンの点で) パフォーマンスが悪くなったり、より多くの要求を要求したりするということです。

望ましい結果を達成するためのガイダンスと人間の警戒。
コードのメンテナンス – LLM はコードベースを良好な状態に保つのが苦手で、一般に進行状況とメンテナンスのバランスを取るのが苦手です。
技術的な聴衆に向けた文章 – LLM は技術的な文章が貧弱で、ほとんどの場合、正しいレベルの抽象化を伝えられなかったり、ニュアンスが欠落していたり​​、含まれるべき歴史的文脈の量を誤解していたり​​します。
コードベースの進化 – 時間をかけてコードを作成することは、一度作成することとは大きく異なり、システム間の依存関係、下位互換性、データ移行、デプロイメントのリズム、さらには変更に対する信頼性の構築を考慮する必要があります。 LLM はこれらすべてに対処しています。
研究 – LLM は、コードベースがどのように機能するかを研究する場合でも、トピックを調査してレポートを作成する場合でも、研究が得意であるように見えますが、専門家なら誰でも証言できるように、微妙な点を見逃したり、技術的には有効であるかもしれないテキストを作成したりすることがよくありますが、間違ったことを強調したり、要点を外すなど間違った方法でストーリーを伝えたりします。
私の仮説は、良いことと悪いことの間のパターンは要件に依存するというものです。
LLM は、要件が少ない場合にはうまく機能しますが、要件が多い場合には困難を伴います。要件が明確な場合にも機能しますが、それが主な要因ではないと思います。
バイブコーディングを例に挙げてみましょう。バイブコーディングで与えられる要件はほとんどありません。ユーザーが何らかの高レベルの目標を指定すると、モデルはそこに到達するために必要なことを自由に実行できます。 React フロントエンドはうまく機能しますか? それともトレーニング データとより一致しますか?次に、React フロントエンドを使用してみましょう。 Postgres バックエンドがトレーニング データの中央値に近いので、それを使用しましょう。ユーザーは気にせず、モデルは自由に暴走し、新しい依存関係を導入します。

どこでも発生する可能性、同じことを行う複数の方法、その他あらゆる種類のエンジニアリング上の罪。
セキュリティについてはどうですか?セキュリティの脆弱性を見つけるのは難しいですが、実際には反事実に関するデータがあまりありません。 LLM を実行しても何も見つからなかった回数は何回ですか?このフィールドは基本的に制約がなく、やはりモデルを自由に探索できるため、うまく機能します。実際のエクスプロイトも、うまく構築する必要はなく、機能することだけが必要です。
同様に、クリエイティブな作品にも要件はほとんどなく、すべてユーザー プロンプトで明示的に規定される傾向があります。有効な出力のスペースは大きいです。
一方で、LLM は時間の経過とともに進化するコードベースに苦労しています。コードベースは基本的に、どの API が利用可能で使用する必要があるか、各サブシステムがどのように動作するか、スキーマなどの要件の集合です。コードベース上の制約を網羅的にリストすると、それは膨大な情報の山となり、コンテキスト ウィンドウが現実的な範囲をはるかに超えて拡大してしまう可能性があります。
これらの要件の多くは、非常に暗黙的なものでもあります。下位互換性は微妙な問題であり、測定するのが難しい場合があります。パフォーマンス要件は暗黙的であり、測定するのが困難です。また、コードベースは時間内に修正されるわけではありませんが、開発には時間の経過とともにコードベースを進化させるための複雑な一連のコミット、デプロイ、移行が必要です。
技術者向けの記事を書く場合も同様の点でつまずくことがありますが、ニュアンスが重要です。要件は、聴衆とその経験、知識、精神状態に基づいています。ドキュメントでは、複雑な概念を伝えるために 3 つすべてを使用したスト​​ーリーを伝える必要があります。研究成果も同様です。
…という質問が聞こえてきます。 LLM に要件を検証するツールを提供し、そのツールを実行して合格するまで作業するように指示するだけです。
これは要点を外していると思います。まず、これらの要件の多くは、

測定できないため、それらを検証できるツールはありません。しかし、すべてを測定できる領域 (例: 100% のテスト カバレッジのコードベース、リンティングなど) であっても、実際に要件に取り組むのではなく、ツールのチェックに合格するソリューションに向けて山登りをしているだけではないでしょうか?実際には、要件の数を「チェックに合格する」という 1 つに減らしただけであると主張することもできます。
このアプローチが役に立たないと言っているわけではありません。私は実際その逆を提案しています。LLM は要件を満たすのが苦手なので、LLM から最高のパフォーマンスを得るには要件を削除して削減する必要があり、ツールはそれを行うためのかなり効果的な方法です。
これは単なるコンテキストエンジニアリングではないでしょうか?
問題は単なるコンテキスト エンジニアリングであり、すべての要件をコンテキストに取り込むことができれば、モデルはうまく機能するだろうと主張する人もいるでしょう。これはある意味では真実ですが、コンテキスト内での明示的な要件の方が確実にうまく機能しますが、これも要点を外しています。
人間は、暗黙の要件に暗黙的に従うのがかなり得意です。つまり、私たちは自分が何かを正しい方法で行ったことに気づかないことが多く、たとえ誰に尋ねられなかったとしても、とにかく最初はそれを正しくやってしまうのです 2 。
これを補うコンテキストは、少なくとも現在の LLM テクノロジーで予測できるものではありません。
ソフトウェアエンジニアリングに限らず、要件は厳しいものです。彼らはいつもそうでした。私の経験では、LLM は要件の複雑さに基づいて能力を拡張するようです。明示的なものは暗黙的なものよりも優れており、検証可能なものは検証不可能よりも優れており、時間と人的要因に依存するものは最も困難です。
人間は要件を書くのも苦手で、それがここでの LLM パフォーマンスの一因となっていますが、私は人間は苦手だと思います。

以下の要件、特に暗黙的な要件に取り組みます。私たちは時間の経過とともに直感、つまりシステムや視聴者の要件が何であるかについての直感を養い、その直感を使って意思決定を導きます。一歩下がって考えるまで、理由が分からずに何かをすべきだとわかっていることはよくあります。
LLM は直観力を養いません。さまざまな検索拡張生成 (RAG) 技術を介して知識を発展させることができますが、それはコンテキストの詳細を取得する別の方法にすぎず、要件に従うという同じ問題が発生します。人間の直感に最も近いのは LLM トレーニングですが、継続的にトレーニングするためのテクノロジーはまだありません。私たちがそのような一歩を前進させれば、この問題全体はほとんど時代遅れになると思います。
その間、より良いトレーニングにより LLM のパフォーマンスが向上し、応用可能なスキルのギャップが縮まります 3 が、彼らが独自の直感を養えるようになるまでは、多くの多様な要件を伴うタスクには常に上限が存在します。スキルの出現にもかかわらず、トレーニング データは依然として重要です。
これらすべてに基づいて、私は LLM が成功する可能性と失敗する可能性が高い場所について独自の直感を開発しています。すべては、LLM が従うべき要件がいくつあるかに帰着します。モデルに任せれば任せるほど、良い結果が得られます。残念ながら、これは大規模で成熟したコードベースでの開発には適していませんが、分離されたコンポーネント、プロトタイピング、グリーンフィールド作業、および最適化の検討には最適です。私はまた、より多くの作業をこれらの側面に移す方法でプロジェクトを構築し始めています。
私は AI が生成した詩や歌には興味がありませんが、機械的には適切な言葉を使って必要な形式の文章を生成することができます。

場所。 ↩︎
私がここで話しているのは、大きく複雑な問題についてではなく、筋肉の記憶、私たちが考えもしない小さな学習されたステップや相互作用について話しているのです。 ↩︎
たとえば、LLM は git 操作に苦労していましたが、私たちはそれを LLM に訓練しました。 git を使用すると、「考えず」に多くのステップを自然に正しく実行できるようになります。 git を使用していない場合は、頑張ってください。 ↩︎

## Original Extract

Blog
Projects
CV LLM Performance and Following Requirements
There’s a disconnect between the experiences different people are having with LLMs. Some find them transformational, others find them incapable. Partly this is down to skill, partly this is down to a dislike of the technology and it’s social implications, but I think there’s a further issue. This post is not trying to be pro-AI or anti-AI, I think the social issues are real, I don’t believe LLMs live up to all of the hype, but I do find using them can be extremely productive.
So what’s the problem? It stems from how LLMs meet requirements, and possibly how humans express them. Let’s do a review of where LLMs work well and fall short to observe some patterns…
This is not to say that LLMs work perfectly for all of these things, but rather that they generally work better, and need less guidance.
Vibe coding – by this I mean deferring all programming to an LLM, not writing any code, not reviewing any code, and focusing on one-shot or few-shot implementations. Clearly LLMs are very good at this, and an ecosystem of products has emerged around it.
Security testing – LLMs are also seemingly amazing at finding vulnerabilities in code and systems. The longer the trajectories, the more capable LLMs are at finding vulnerabilities.
Short creative works – LLMs excel at things like poetry and songwriting. There’s a separate question of artistic value 1 , but it’s undeniable that LLMs can produce perfectly reasonable poems, songs, short form creative writing, and change the voice or nature of existing documents.
Similarly to where LLMs work well, this is not to say that LLMs fail on these next areas, but rather that they struggle, perform worse (in quality, time, tokens), or require more guidance and human vigilance to achieve the desired outcome.
Code maintenance – LLMs are not good at keeping codebases in good condition, being generally poor at balancing forward progress with maintenance.
Writing for a technical audience – LLMs produce poor technical writing, almost always failing to convey the correct level abstraction, missing nuances, or misunderstanding the amount of historical context that should be included.
Evolving a codebase – Producing code over time is very different to producing it once, needing to account for cross-system dependencies, backward compatibility, data migrations, deployment cadence, and even building confidence in changes. LLMs struggle with all of these.
Research – LLMs often appear to be great at research, whether it’s researching how a codebase works, or researching a topic and preparing a report, but as any expert can attest to, they often miss subtleties and produce text that may be technically valid, but emphasises the wrong things or tells a story in the wrong way such as to miss the point.
My hypothesis is that the pattern between the good and the bad is down to requirements .
LLMs work well when there are few requirements, and struggle when there are many. They also work well when the requirements are clear, but I don’t think that’s the main driver.
Let’s take vibe coding as an example. There are very few requirements given with vibe coding. The user specifies some high level goal, and the model is free to do whatever it wants to get there. Does a React frontend work well, or match the training data more? Then let’s use a React frontend. Is a Postgres backend close to the median in the training data, let’s use that. The user doesn’t care, and the model is free to run wild, introducing new dependencies anywhere, multiple ways of doing the same thing, and all manner of other engineering sins.
What about security? Finding security vulnerabilities is hard, but we actually lack much data about the counterfactuals. How many runs with LLMs failed to find anything? This field works well because it is fundamentally unconstrained, and again the model is free to explore. The actual exploit also doesn’t need to be well built, it only needs to work.
Creative works similarly tend to have few requirements, all explicit in the user prompt. The space of valid outputs is large.
On the other hand, LLMs struggle with evolving codebases over time. A codebase is essentially a collection of requirements – which APIs are available and should be used, how each subsystem works, schemas, and so on. If one were to exhaustively list the constraints on a codebase, it would be a vast trove of information, and would likely blow up the context window far beyond what is practical.
Many of those requirements are also very implicit. Backwards compatibility is a subtle issue and can be hard to measure. Performance requirements are implicit and hard to measure. And a codebase is not fixed in time, but development requires a complex series of commits, deploys, and migrations to evolve it over time.
Writing for a technical audience trips up on some of the same aspects, nuance is key. The requirements are based on the audience and their experience, knowledge, and state of mind. Documentation requires telling a story that uses all three to convey complex concepts. Research output is also similar.
…I hear you ask. We can give LLMs a tool that verifies requirements, and just tell them to run the tool and work until it passes!
I think this misses the point. Firstly, many of these requirements cannot be measured, there is no tool that can verify them. But even in areas where we could measure everything (e.g. a codebase with 100% test coverage, linting, etc), aren’t we just hill-climbing towards a solution that passes the tool’s checks, rather than actually engaging with the requirements? It could be argued that we have in fact only reduced the number of requirements to one: “pass the checks”.
That’s not to say this approach is useless, I’m in fact proposing the opposite – that LLMs are bad at meeting requirements so to get the best performance from them we must remove and reduce requirements, and tools are a fairly effective way of doing that.
Isn’t this just context engineering?
You could argue that the problem is just context engineering, and that if we can get all the requirements into the context then the model will do fine. This is true in some ways – an explicit requirement in the context is certainly going to work better – but this too misses the point.
Humans are reasonably good at implicitly following implicit requirements. That is to say that we often don’t realise we’ve done something the right way, we just get it right the first time anyway 2 , even when when no one asked.
There’s no amount of context, at least not one we can forsee with current LLM technology, that will make up for this.
Requirements are hard, in software engineering and beyond. They always have been. LLMs seem, in my experience, to scale in ability based on the complexity of requirements. Explicit is better than implicit, verifiable better than unverifiable, and those that depend on time and human factors are the most challenging.
Humans are also bad at writing requirements, and that’s a contributing factor to LLM performance here, but I think humans are better at following requirements, particularly implicit ones. We develop intuition over time, intuition for what the requirements of a system, or audience are, and we use that intuition to guide our decisions. It’s often the case that we know to do something without knowing why until we step back and think about it.
LLMs don’t develop intuition. They can develop knowledge via various retrieval augmented generation (RAG) techniques, but that’s just another way to get more detail in the context and produces the same problem of following requirements. It’s LLM training that is the closest analogy to human intuition, and we don’t yet have the technology for continuous training. I think when we develop such a step forward it’ll make this whole problem largely obsolete.
In the mean time, better training will continue to improve LLM performance and close the gap for transferrable skills 3 , but until they can develop their own intuition, there’s always going to be a ceiling on tasks that have many and varied requirements. Despite the advent of skills, training data is still critical.
Based on all of this, I’m developing my own intuition for where LLMs are likely to succeed and fail, and it all comes down to how many requirements there are for them to follow. The more I can leave up to the model, the better it will do. Unfortunately that doesn’t align well with development in large, mature codebases, but it’s great for isolated components, prototyping, greenfield work, and exploring optimisations. I’m also starting to structure projects in a way that shifts more of the work towards these aspects.
I have no interest in AI generated poetry and songs, but mechanically they are capable of producing the necessary format of writing with the right words in the right place. ↩︎
I’m not talking here about big complex problems, I’m talking about muscle memory, small, learned steps and interactions that we don’t even think about. ↩︎
LLMs used to struggle with git operations for example, but we’ve trained that into them. If you use git then they’ll naturally get a lot of steps right without “thinking” about it. If you don’t use git, good luck. ↩︎
