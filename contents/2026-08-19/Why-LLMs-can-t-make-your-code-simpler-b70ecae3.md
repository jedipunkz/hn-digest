---
source: "https://www.answer.ai/posts/2026-08-19-llms-code-simpler.html"
hn_url: "https://news.ycombinator.com/item?id=49360447"
title: "Why LLMs can't make your code simpler"
article_title: "Why LLMs can’t make your code simpler – Answer.AI"
image: "https://www.answer.ai/posts/2026-08-19-llms-code-simpler_assets/naur.jpeg"
author: "tosh"
captured_at: "2026-08-19T12:24:24Z"
capture_tool: "hn-digest"
hn_id: 49360447
score: 1
comments: 0
posted_at: "2026-08-19T12:09:42Z"
tags:
  - hacker-news
  - translated
---

# Why LLMs can't make your code simpler

- HN: [49360447](https://news.ycombinator.com/item?id=49360447)
- Source: [www.answer.ai](https://www.answer.ai/posts/2026-08-19-llms-code-simpler.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T12:09:42Z

## Translation

タイトル: LLM がコードを簡素化できない理由
記事のタイトル: なぜ LLM ではコードを簡素化できないのか – Answer.AI
説明: 古典的な論文「理論構築としてのプログラミング」が LLM について教えてくれること

記事本文:
コードの複雑さが間違った指標である理由
LLM がコードを簡素化できない理由
この投稿はもともと Medium で公開されました。
ピーター・ナウル博士の「理論構築としてのプログラミング」では、実際のプログラム、つまり大文字の T を付けて彼が理論と呼ぶものは、エンジニアの頭の中にあると述べています。コードとドキュメントは単なる下流の (したがって不完全な) 成果物です。 LLM に関する私の主な不満の 1 つは、コードが冗長であり、随所に複雑さが入り込んでいることです。私は常に、LoC や独立したコード パスの数などの指標を使用して制約を設けることで、この問題を改善できるのではないかと信じていました。しかし、Naur を読んだ後、私たちが軽減しようとしている複雑さはコードではなく理論の問題であり、非常に主観的なものであるため、それに関連する対策はないことに気づきました。
私は最近、Peter Naur による素晴らしい論文「理論構築としてのプログラミング」を読みました。現在の LLM で何ができるか、できないか、どのようにプロンプ​​トを表示するか、現在のエージェント システムの主な制限は何か、ペア プログラミングがなぜ非常に効果的であるかなどについて、私の見方は大きく変わりました。
今日はコードの複雑さとの関係に焦点を当てます。
まだこの論文を読んでいないのであれば、ぜひ読むことをお勧めします。実際、この記事を書く私の主な目的は、何人かに原論文を読んでもらうことです。努力する価値はあります。また、Solveit で精読することを練習 (または学習) する絶好の機会でもあります。なぜなら、Solveit のほうがはるかに簡単だからです。本文全体で質問したり、好きなところを調べたり、Solveit に言葉をよりわかりやすくしてもらうこともできます。詳細については、このブログ投稿を参照してください。私自身のダイアログをフォークすることですぐに始めることができます。
それを念頭に置いて、それでも読まないことに決めた場合は、この論文の内容を以下に示します。
プログラムは理論です

プログラムを構築および保守する人々が保持しているもの、つまり、プログラムが現実世界の問題にどのように関係しているか、どのような制約とトレードオフがそれを形作っているのか、なぜ機能するのか、そしてどの変更がその設計に適合するのかについての理解です。コードとドキュメントは理論の下流の成果物であり、理論を完全に捉えることはできません。
エンジニアは、ユーザーと話し、障害を観察し、ドメインを学習し、現実世界でシステムがどのように動作するかを確認するなどの経験を通じて、この理解を深めます。この理解は、関連性、類似性、シンプルさ、優れたデザインの判断に役立ちます。
LLM は、継続的に学習したり、ユーザーと話したり、現実世界で物事を体験したりすることが苦手です。しかし、なぜこれが複雑さに関係するのでしょうか?
一般に LLM を放置するとコードベースの複雑さが増す傾向があるということには誰もが同意できると思います。理由は数多くあります。メソッドがすでに存在していることに気づかずにそれを複製したり、起こり得ないエッジケースから保護するなどの過剰防御的なコードを作成したり、物事を早期に過剰最適化しすぎたりします。余談になりますが、ほとんどのフロンティアラボは、トークンを消費すればするほど多額の利益を得ることができるため、トークンマックス化を促進するよう奨励されています。全体として、LLM が KISS 原則に従うことはほとんどありません。この問題は、出力をチェックせずにバイブコーディングを行う場合に最悪になります。ただし、コードをレビューする場合でも、プログラムを簡潔に保ちたい場合は、複雑さを可能な限り削減するための積極的な努力が必要です。
私の素朴な頃（約 2 週間前）は、いつかこの複雑さの穴から抜け出せるだろうと考えていました。フロンティア ラボは、RL トレーニングに複雑さのペナルティを追加するだけです。最小化するための指標として合計 LoC を使用することもできますが、場合によっては 1 行の方が 2 ～ 3 行よりも複雑になる可能性があることに誰もが同意します。別のオプションはサイクロマチックです

独立したコードパスの総数を測定する複雑さ。 Peter Naur を読んだ後、それらでは問題を実際には解決できないことに気づきました (おそらく、少しは軽減できるかもしれません)。その理由を見てみましょう。
コードの複雑さが間違った指標である理由
以下の (発明された) 例では、OpenAI と Anthropic の呼び出しをサポートしたいと考えています。すべてのメッセージの準備と再試行は同じように処理されますが、リクエスト本文のパラメーターがわずかに異なるため、 call_openai と call_anthropic という 2 つの異なるメソッドがあると想像してください。
最初の単純なバージョンでは、重複したボイラープレート コード (つまり prepare と with_retries ) を持つ 2 つの異なるクラスがあります。
# 分離 — メトリクスが重複にフラグを立てる
クラスOpenAIClient:
def complete( self , プロンプト):
msgs = prepare(prompt) # 共有定型文
y = call_openai(msgs) # 唯一異なる行
return with_retries(y) # 共有定型文
クラス AnthropicClient:
def complete( self , プロンプト):
msgs = 準備(プロンプト)
y = call_anthropic(msgs)
return with_retries(y)
簡単なリファクタリングは、DRY するために単一のクラスを作成することです。多くの指標から見て、これはより良い実装です。コード行が少なく、プログラム長 (N) と語彙 (n) の V=N×log2(n) として計算されるハルステッドのボリュームが優れており、保守性指数も優れています。
# マージ — メトリクス的には本当に優れています: DRY、行数の減少
クラスLLMClient:
def __init__ ( self , プロバイダー): self .provider = プロバイダー
def complete( self , プロンプト):
msgs = 準備(プロンプト)
y = call_openai(msgs) if self .provider == "openai" else call_anthropic(msgs)
return with_retries(y)
一般に、2 番目のバージョン (または LoC と重複を削減する同様のバージョン) はそれほど複雑ではありません。しかし、来月には Anthropic のサポートを停止する可能性が非常に高いと言ったらどうでしょうか。その状況で

私は、必要なときに AnthropicClient を含むファイルを簡単に削除できるように、これらを分離しておいた方がよいと考えています。
もちろん、このような単純な例は、特に LLM がコーディングできるようになったので、簡単に解決できます。しかし、LiteLLM のように 2 つのプロバイダーではなく 165 を超えるプロバイダーをサポートしたい場合はどうすればよいでしょうか?その場合は、LiteLLM を使用してください。しかし、その場合、あなたと最終プロバイダーの間に +120 万を超える Python LoC がかかることになります。それだけの価値はありますか?
適切に設計された API は、複雑さを抽象化する優れた方法です。明確な契約があるので、その背後で何が起こっているかを理解する必要はありません。 LiteLLM が巨大なパッケージであっても、理論の全体的な複雑さにはカウントされない可能性があります。 LLM 推論エンドポイントの初期の頃は、「テキストを送信すると、テキストが出力される」というようなものでした。ただし、契約が信頼できなくなると、これは破綻します。これは、バグが多いため、API の背後に隠れて取得できない状態が多数あるため、または単に新しいプロバイダー機能がサポートされているかどうかが不明なために発生する可能性があります。
API 抽象化レイヤーを超えて深淵を覗き込む必要があるときは常に、その複雑さが問題になります。この問題はますます一般的になってきています。 OpenAI や Anthropic などのプロバイダーは、暗号化された圧縮データや推論トークンなどのデータをサーバー側で非表示にすることが増えています (詳細はこちら)。その上に抽象化レイヤーを提供しようとする人は、プロバイダー固有の詳細を公開するか、ますます不安定になっている互換性ロジックを再実装するか、抽象化がその下のシステムを忠実に表現できないことを受け入れるかのいずれかを余儀なくされます。
基本的な機能を使用して迅速に作業を進め、多くのプロバイダーを試してみたい場合は、LiteLLM または同様のライブラリを使用する価値があるかもしれません。一方、制御、デバッグ、およびインストーラを重視する場合は、

スタックを理解しているなら、おそらくそうではありません。適切な「複雑さ」というものはありません (ただし、私は 2 番目のオプションを非常に好みます)。
どのパスを取るかを決定するための情報はコード内には存在しません。この情報は、ナウルがプログラムの理論と呼ぶものの一部です。この情報は人々の心の中に存在するため、LLM はこれまでのところ、この情報にほとんどアクセスできません。 IM、電子メール、または他の種類の文書でそれに関するヒントを得ることができますが、それらはすべて常に部分的です (最良の場合)。
この情報の一部は、ビジネスの優先順位、予想される製品の変更、運用上の制約、以前の決定の背後にある理由などのコンテキストとして LLM に提供できます。そうすることで選択肢は改善されるかもしれませんが、これらは依然として理論の成果物です。チームがその理論を開発した経験と判断を完全に移転することはできません。
それはそれでいいのですが、もし私がバイブコーディングとトークンマックス化に全力で取り組んでいたらどうなるでしょうか、あるいはコードをまったく気にしなかったらどうなるでしょうか。まあ、そのような状況では、ブログ投稿の残りの部分を読んでも、そうでないことを納得させることはできません。一方、あなたがその中間にいる場合は、Answer.AI で Solveit の請求システムを使用して経験した実際の状況について説明します。
ネタバレに注意してください。Answer.AI では、いかなる犠牲を払ってでも複雑さを最小限に抑えることに多大な投資を行っています。実際、この投資により、コードがより単純になり、頭の中に残りやすくなる傾向があります。
Solveit は、AI を使用してノートブックのような環境で作業できるプラットフォームです。環境は永続的なため、LLM の使用量、CPU、ディスク、メモリ、帯域幅に対して料金が発生します。
当初の計画では、ユーザーに月額サブスクリプション料金 (例: 5 ドル) を請求する予定でした。この毎月のサブスクリプション金額は、その月に消費されるクレジットになります。すべて消費した場合は、期限が切れる前に残高を補充する必要があります。

xt月。私たちはこのアプローチを検証するために、まず小規模なプロジェクトでこのアプローチをテストしました。
このアプローチは、私たちが望んでいたよりも複雑であることが判明しました。まず、クレジット + サブスクリプションの仕組みは、一部の人 (私と同じように) を混乱させるでしょう。第 2 に、コードが複数のレベルでより複雑になりました。最初に「月次サブスクリプション クレジット」を消費し、次に通常のクレジットを消費するか、または残ったクレジットをどうするかというすべてのロジックに対処する必要があります。 Stripe 側には、手動チャージとサブスクリプション サービスという 2 つの異なるコード パスがありました。
馴染みのない方のために説明すると、Stripe サブスクリプションはフルマネージド サービスです。 Stripe はライフサイクルを管理します (再発、請求、再試行はすべてストライプ側で行われます)。
大まかに次のようにサブスクリプションを作成するだけです。
Stripe.Subscription.create(customer = cust_id, items = [{ "価格" : "price_5usd_monthly" }])
そして、サブスクリプションステータスが変更されたとき、または支払いが到着したとき（またはあなたが決定できる他の多くのオプション）に、WebhookをリッスンしてDBを更新します。
直接請求では、ユーザーが Stripe のドメインにアクセスしてカードに情報を記入できるように、チェックアウト セッションを開始する必要があります。顧客 ID を提供する必要があります。どこでストライプ顧客を作成する必要がありますか?ユーザーがサインアップするとき?彼らが支払おうとしたとき、何か別のことがあるでしょうか？すべて有効なオプションです。
Stripe.checkout.Session.create(mode = "支払い" , customer = cust_id,
line_items = [{ "価格" : "price_5usd" , "数量" : 1 }],
success_url = "https://yoursite.com/done" )
ここまでは大丈夫ですか？ Stripe や支払いの経験があまりない場合は、すでにすべてを頭の中に収めるのに苦労している可能性があります。おそらく、次のように頭の中で単純化できたのではないでしょうか。
サブスクリプション -> ストライプ サブスクリプションによって管理されます
明らかではない問題は、Stripe マネージド サービスを使用していることです。

データが重複していることを意味します。データの半分は Stripe のバックエンドに存在するため、ローカル DB がバックエンドと同期していることを確認する必要があります。もう 1 つの問題は、デバッグ中に Stripe サービスにクエリを実行し、独自の DB をチェックする必要があることです。たとえば、支払いが表示されない場合、Stripe が Webhook を送信しなかったため (「彼らのせい」)、それとも DB に Webhook を保存できなかったため (「当社のせい」) なのでしょうか?
Stripe サブスクリプション サービスは優れており、セットアップも簡単ですが、膨大な数のユースケースをサポートするように設計されています。つまり、たとえ適切に設計されていたとしても (確かにそのとおりです)、API の抽象化は非常に複雑になってしまいます。この場合、Stripe API を学習する複雑さと引き換えに、袖をまくって自分でコードを書く複雑さをトレードすることになります。
LiteLLM と Stripe は、同じトレードオフをさまざまなスケールで示しています。外部抽象化は、契約が維持されている間は理論を大幅に簡素化できますが、その契約を超えてデバッグ、変更、または推論する必要がある場合は常に、その隠れた複雑さが自分のものになります。
Answer.AI では、すべてのスタックを所有しようとします。その理由については記事全体 (または複数) を書く価値があるため、ここでは詳しく説明しません。したがって、私たちは上記の状況に満足していませんでした。
何日にもわたる調査と議論の後、私たちは次のことに落ち着きました。

[切り捨てられた]

## Original Extract

What the classic paper “Programming as Theory Building” can teach us about LLMs

Why Code Complexity is the wrong metric
Why LLMs can’t make your code simpler
This post was originally published on Medium .
tl;dr Peter Naur’s “Programming as Theory building” states that the real programs, what he calls Theory, with capital T, are in the mind of the engineers. The code & documentation are just downstream (and thus incomplete) artifacts. One of my main complaints about LLMs is how verbose the code is and how complexity creeps in everywhere. I always had some faith that we might improve that using metrics like LoC or the number of independent code paths to constrain them. However, after reading Naur I realized that the complexity we are trying to reduce is the Theory one, not the code, and there are no relevant measures for that because it’s very subjective.
I recently read the fantastic paper “Programming as Theory building” by Peter Naur . It has changed quite substantially my views on what current LLMs can or can’t do, how to prompt them, what are the main limits to current agentic systems or why pair programming is so effective.
Today I will focus on its relation to code complexity.
If you haven’t read the paper, I really suggest that you do. In fact, my main goal in writing this post is to get some of you to read the original paper. It is worth the effort. It is also the perfect chance to practice (or learn!) to do close reading in Solveit because it’s much easier: you can ask any questions throughout the text, go down any rabbit holes you like, or just ask Solveit to make the language clearer to you. More information on close reading in this blogpost and you can quickly get started by forking my own dialog .
With that in mind, if you still decide not to read it, here’s the paper’s tl;dr:
A program is the Theory held by the people who build and maintain it: an understanding of how the program relates to the real-world problem, which constraints and trade-offs shaped it, why it works, and which changes would fit its design. Code and documentation are downstream artifacts of that Theory, and can never capture it completely.
Engineers develop this understanding through experience: talking to users, observing failures, learning the domain, and seeing how the system behaves in the real world. This understanding guides judgments of relevance, similarity, simplicity, and good design.
LLMs aren’t great at continual learning, talking to users, or experiencing things in the real world. But why is this relevant to complexity?
I think we can all agree that LLMs in general tend to increase the complexity of codebases if left unchecked. The reasons are many: they fail to realize a method already existed and duplicate it, they write overdefensive code like protecting against edge cases that can’t happen, or overoptimizing things too early. Tangentially, most frontier labs make a hefty sum the more tokens you consume, so they are kind of incentivized to promote tokenmaxxing. All in all, LLMs rarely follow the KISS principle. This issue is at its worst when vibe-coding w/o checking the output. But even when you review code, if you want to keep a program concise it requires an active effort to cut down complexity as much as possible.
In my naive days (about two weeks ago), I used to think we would get out of this complexity pit at some point. Frontier labs would just add some complexity penalization to their RL training. They could use total LoC as a metric to minimize, but we can all agree that sometimes a one-liner can be more complex than 2–3 lines. Another option would be cyclomatic complexity which measures the total number of independent code paths. After reading Peter Naur I realized that those things can’t really solve the problem (maybe they can alleviate it a bit though). Let’s see why.
Why Code Complexity is the wrong metric
In the (invented) example below we want to support calling OpenAI & Anthropic. Imagine all the message preparation and retries are handled identically, but the request body params are slightly different so we have 2 different methods call_openai and call_anthropic .
In the first naive version we have 2 different classes with duplicated boilerplate code (ie prepare and with_retries ).
# Separate — duplication a metric would flag
class OpenAIClient:
def complete( self , prompt):
msgs = prepare(prompt) # shared boilerplate
y = call_openai(msgs) # the only line that differs
return with_retries(y) # shared boilerplate
class AnthropicClient:
def complete( self , prompt):
msgs = prepare(prompt)
y = call_anthropic(msgs)
return with_retries(y)
A straightforward refactor would be to create a single class so we DRY. By many metrics this is a better implementation: it has fewer lines of code, better Halstead volume — calculated as V=N×log2(n) for program length (N) and vocabulary (n) — and better Maintainability Index .
# Merged — genuinely better by the metrics: DRY, fewer lines
class LLMClient:
def __init__ ( self , provider): self .provider = provider
def complete( self , prompt):
msgs = prepare(prompt)
y = call_openai(msgs) if self .provider == "openai" else call_anthropic(msgs)
return with_retries(y)
In general, the second version (or similar versions that reduce LoC and duplication) are less complex. However, what if I told you that it is very likely that next month we will stop supporting Anthropic. In that situation I find it preferable to keep them separated so when the time comes I can simply remove the file containing the AnthropicClient .
Of course a simple example like this is easy to solve, especially now that LLMs can code for you. But what if instead of 2 providers you want to support 165+ providers like LiteLLM does? Well in that case just use LiteLLM. But then you will be putting over +1.2M python LoC between you and the final provider. Is it worth it?
Well-designed APIs are a great way to abstract complexity. You have a clear contract and you don’t need to understand what happens behind. Even if the LiteLLM is a huge package, it might not count towards your total Theory complexity. The early days of LLM inference endpoints were like this: “you send text in, you get text out”. However, this breaks down when the contract is not reliable anymore. This can happen because it’s buggy, or because there’s a lot of state hidden behind the API which you can’t get, or simply because you’re not sure if a new provider feature is supported.
Whenever you are forced to peer into the abyss beyond the API abstraction layer, that complexity becomes your problem. This problem is becoming more common. Providers like OpenAI and Anthropic increasingly hide data server-side like encrypted compaction data or reasoning tokens ( more on this ). Anyone aiming to offer an abstraction layer on top of them is forced to either expose provider-specific details, reimplement increasingly brittle compatibility logic, or accept that the abstraction cannot faithfully represent the systems beneath it.
If you are moving fast, using basic features and want to try many providers then LiteLLM or similar libs might be worth it for you. If on the other hand you value control, debugging and in understanding your stack, then probably not. There’s no right “complexity” (although I very much prefer the second option).
The information to decide which path to take does not live in the code. This information is part of what Naur calls the Theory of the program. LLMs so far have little to no access to this information because it lives in people’s minds. They can get hints about it in IM, email, or other kind of written docs but all of those will always be partial (best case).
Some of this information can be supplied to an LLM as context like business priorities, expected product changes, operational constraints, and the reasons behind earlier decisions. Doing so might improve its choices, but these are still artifacts of the Theory. They cannot completely transfer the experience and judgment through which the team developed that Theory.
This is all good, but what if I am all-in with vibe-coding & tokenmaxxing, what if I don’t care at all about the code. Well in that situation the rest of the blog post won’t convince you otherwise. If on the other hand you’re in-between I will describe a real situation we went through at Answer.AI with our billing system for Solveit.
Spoiler alert, at Answer.AI we are heavily invested in minimizing complexity at all costs. This investment actually tends to result in simpler code that is easier to hold in your mind.
Solveit is a platform where you can use AI to work in a notebook-like environment. The environment is persistent so we charge for LLM usage, cpu, disk, memory and bandwidth.
The initial plan was to charge users a monthly subscription fee (e.g. $5). This monthly subscription amount would then become credits to be consumed during the month, if you consume them all then you’d have to top up your balance before next month. We tested this approach in a smaller scale project first to validate it.
This approach turned out to be more complex than we wanted. First, the credit + subscription mechanism would make some people (like myself) confused. Second, it made the code more complex at multiple levels. You have to deal with all the logic of consuming “monthly subscription credits” first, and then normal credits or what to do with leftover credits. On the Stripe side of things it had 2 distinct code paths: manual top up and the subscription service.
For those of you not familiar with it, Stripe subscriptions is a fully managed service. Stripe manages the lifecycle (recurrence, invoicing, retries all happen on their side).
You just create the subscription roughly like:
stripe.Subscription.create(customer = cust_id, items = [{ "price" : "price_5usd_monthly" }])
and listen to their webhooks to update your DB when subscription status changes or when payments arrive (or many other options you can decide on).
The direct charge requires you to start a checkout session so users land in Stripe’s domain and fill in their card. It requires you to provide a customer ID. Where should you create the stripe customer? when your user signs up? when they try to pay, something else? all are valid options.
stripe.checkout.Session.create(mode = "payment" , customer = cust_id,
line_items = [{ "price" : "price_5usd" , "quantity" : 1 }],
success_url = "https://yoursite.com/done" )
Good so far? If you don’t have a lot of experience with Stripe or payments, chances are you will be already struggling to hold it all in your head. Maybe you have managed to simplify it into your mind like:
Subscriptions -> managed by stripe subscriptions
A problem that is not obvious is that using Stripe managed service means that you have duplicated data. Half of the data lives in Stripe’s backend and you have to make sure your local DB is synchronized with it. Another problem is that while debugging, you have to both query Stripe services AND check your own DB. For example, if a payment does not show up is it because Stripe did not send the webhook (“their fault”), or because we failed to store it in our DB (“our fault”)?
Stripe subscription service is great and easy to set up, but it is designed to support a huge number of use cases. That means that, even if well designed (which it certainly is), the API abstraction ends up being quite complex. In this case you are trading the complexity of rolling up your sleeves and writing that code yourself, for the complexity of learning the Stripe API.
LiteLLM and Stripe illustrate the same trade-off at different scales: an external abstraction can greatly simplify your Theory while its contract holds, but its hidden complexity becomes yours whenever you must debug, modify, or reason beyond that contract.
At Answer.AI we try to own all the stack. The reason why merits a full article (or multiple ones) so I won’t delve into it. So we were not pleased with the state of affairs discussed above.
After many days of exploration and discussion we settled o

[truncated]
