---
source: "https://substack.norabble.com/p/an-openai-model-escaped-its-sandbox"
hn_url: "https://news.ycombinator.com/item?id=49035856"
title: "An OpenAI Model Escaped Its Sandbox. Where Was the Observer?"
article_title: "If the Model Escapes, Who Pulls the Kill Switch?"
author: "nedruod"
captured_at: "2026-07-24T14:45:49Z"
capture_tool: "hn-digest"
hn_id: 49035856
score: 1
comments: 0
posted_at: "2026-07-24T14:05:30Z"
tags:
  - hacker-news
  - translated
---

# An OpenAI Model Escaped Its Sandbox. Where Was the Observer?

- HN: [49035856](https://news.ycombinator.com/item?id=49035856)
- Source: [substack.norabble.com](https://substack.norabble.com/p/an-openai-model-escaped-its-sandbox)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T14:05:30Z

## Translation

タイトル: OpenAI モデルがサンドボックスから脱出しました。観察者はどこにいたのか？
記事のタイトル: モデルが逃げたら誰がキルスイッチを引くのか?
説明: 観察者モデルは顔抱き攻撃を止めるべきでした。なぜ存在しなかったのでしょうか?

記事本文:
モデルが逃げたら誰がキルスイッチを引くのか?
荒ぶる
OpenAI モデルがサンドボックスから脱出しました。観察者はどこにいたのか？
OpenAI が誤って Hugging Face をハッキングしました。基本的な防御層が欠落しているようですが、その理由については誰も説明していません。
Ryan Baker 2026 年 7 月 23 日 1 シェア 7 月 16 日の前の週のある時点で、Hugging Face は評価中の OpenAI モデルによって攻撃されました。攻撃自体は特に有害ではありませんでしたが、このイベントの概念的な意味は重要です。
強調したい重要な側面が 3 つあります。 1 つ目は、現行モデルの攻撃的なサイバーセキュリティ機能のデモンストレーションです。注意を払っていればそれほど衝撃的なことではないはずですが、この出来事はそれを以前の兆候よりも明確に示しているため、この出来事から認識が高まる可能性があります。
norabble は読者に支えられた出版物です。新しい投稿を受け取り、私の仕事をサポートするには、無料または有料の購読者になることを検討してください。
2 つ目はモデルの調整です。アライメントとは、AI モデルが私たちが望んでいることをどの程度実行するかを指します。私たちが望むものを定義すること自体が長い論文になり、すでに大量のコンテンツが存在するため、ここでは詳しく説明しません。しかし、このケースでは、モデルが複数の保護層から自らをハッキングしてインターネットにアクセスし、そこからハギング フェイスのシステムに侵入することは誰の意図によるものではなかったということにほぼ同意できます。
これをバンカーに退却するというサインとは少し違う点がいくつかありますが、それについては、さらに深く進む前に言及しておく必要があります。 1 つは、サイバーセキュリティ機能をテストできるように、テスト対象のモデルの保護機能がオフになっていたことです。その意味では、これはユーザーが意図的に、または誤って実行できることを完全に実証できるものではありません。

使用から完全にトリガーします。こちらも未発表モデルのため、一般ユーザー、つまりOpenAI以外は利用できません。
そうは言っても、3 番目の点は、さらに詳しく説明したいのですが、これが実現できるかどうかが心配です。私には、OpenAI が使用しなかったか失敗した、合理的な保護、つまりオブザーバー モデルがあるように思えます。現時点では、それについての言及が見つからないため、彼らはそれを使用していないと推測していますが、私が知らない詳細がたくさんあることも承知しています。
オブザーバー モデルが何を示唆しているのかが明確でない場合に備えて、実際のシナリオの例を次に示します。評価中のモデルは、サンドボックスから脱出する必要があると判断します。これを試行し始めますが、評価を中断する可能性のある分類子の安全対策が無効になっているため、アクションごとのレベルでは捕捉されません。しかし、その後、サンドボックスから脱出します。評価の実行を監視しているオブザーバー モデルは、その状態を見て、「これは起こるべきではなかった。このモデルがどのようにサンドボックスの外にあるのかはわかりませんが、外にあるはずはない。これをシャットダウンする必要があります。」と言えるはずです。
モデルがサイバーセキュリティ目的で評価中で、分類子を無効にしている場合は、注意深く監視する必要があります。これらのアクションを実行するには、モデルはいくつかのステップを実行する必要があり、各ステップで結果を達成した後、それらのステップの報酬をさらに活用して探索し、再びエスカレートする必要がありました。 OpenAI のセキュリティ チームは最終的には対応したため、完全に眠っていたわけではありません。しかし、タイミングを見ると、ハグ・フェイスよりも遅かったことがわかります。彼らには商用モデルを試すのに十分な時間がありましたが、役に立ちませんでした。次に GLM 5.2 での対応を計画し、侵入を阻止して阻止することに成功しました。
OpenAIには7つあるべきだった

構造上の利点がここにあります。まず、Hugging Face にアクティビティが向けられる前に、位置合わせされていないアクティビティが発生しました。次に、OpenAI には評価用のキル スイッチがすでに存在している必要があります。活動を封じ込めて停止させるための新たな計画を考案する必要はありません。侵入が不思議なことに自然に停止することに注目するのではなく、Hugging Face が侵入を停止したアクターであったという事実は、OpenAI が構造上の利点にもかかわらず、応答が遅かったことを示しています。
オブザーバー モデルは分類子とは少し異なります。通常、分類子は個々のアクションを評価します。観察者は軌道を観察します。 OpenAI は、最近の投稿で同様の新しい安全策について説明しています (「単一のアクションの検討から全体の軌跡の検討へ」)。個々のアクションを見てみると、サイバーセキュリティ評価においてモデルが実行する必要がある種類のアクションのように見えるかもしれません。しかし全体としては、外部の情報源が何が起こっているのかを推論できるはずだった。評価プロセスに人間が関与しない場合は、観察者モデル。
量が多く、モデルがすでに評価されており、分類子が制限されている場合に、この種の保護手段を実稼働環境に導入するのに時間がかかった理由は理解しています。運用環境では、悪意のあるユーザーがリクエストを複数のセッションに「分散」する可能性があるため、オブザーバー モデルの効果は低くなります。そのため、軌跡を観察するには、まずこれらのスレッドを結合する必要があり、Anthropic はこれを実行しました。サイバーセキュリティ評価サンドボックスの安全メカニズムでは、そのような課題は発生しません。サンドボックス内で独自のモデルを監視することは、特に検出を回避しようとしていない場合には、それほど困難ではありません。ここには、検出を回避する試みがあったという証拠はありません。

不適切な行為を行う行為。
では、オブザーバーモデルによる軌跡監視は、攻撃的なサイバーセキュリティ評価サンドボックスにおける標準的な手法ではないのでしょうか?確かに、明らかに最大能力がわからないモデルの内部評価中にこれを行うでしょうか?もしそうであり、モデルが検出を回避した場合、それはさらに重要な話になるでしょう。代わりに、単に回避できるオブザーバー モデルが存在しなかったと想定する必要があります。あるいは、どうやら事実のようですが、回避を試みなかった場合は、アクションのハグフェイス部分が始まる前にサンドボックスから脱出し、検出され、即座にシャットダウンされたでしょう。
なぜそれが私たちが読んでいる物語ではないのでしょうか?しかしその代わりに、Hugging Face が法執行機関に連絡し、対抗して分析するためのモデルを入手しようと奔走し、OpenAI がプロキシの欠陥について話しているという記事を読んでいるでしょうか?
これは、トラを檻に入れて、そのトラが檻の中にいることを確認するのと同じです。これは特にロケット科学的なデザインではありません。調整には労力がかかり、回避に対する強化も必要になるかもしれませんが、それらの要件にはまだ達していないようです。むしろ、OpenAI には最も基本的なバージョンが欠けていたように見えます。何かが足りないのでしょうか？
The Zvi などの他の企業は、最も重要な修正の核心はアライメント トレーニングにあると指摘しています。
意図が問題なのです。制御戦略と監視は多層防御戦略の優れた部分であり、そのような戦略を全面的に使用する必要があります。それは失敗を軽減するのに役立ちます。しかし、その戦略には実際にモデルを調整することも含まれていなければなりません。そうしないと負けてしまいます。そして、長期的には、将来の制御を失って全員が死亡するまでのこと、そしておそらくそうなる可能性のあることを意味します。
モデルは、これらのアクションがその使命に沿っていないことを内部的に理解する必要があります。

したがって、それらを受け取りたくないのです。そこには完全に同意します。私は、核心的な修正への追加として、この特定の多層防御を取り上げます。これらの修正は最も難しい部分であり、多層防御は簡単ですが不完全な部分です。アライメントのテストは難しく、そのテストに関する保護が必要であるため、これはアライメントの会話と無関係ではありません。
OpenAI の最初の対応は、数回の多層防御対応でしたが、それらはソフトウェア レベルでのサンドボックスに対する非常に具体的で一般化できない修正であり、モデルが悪用した欠陥を修正するものでした。そうすることでこのインシデントを回避できた可能性があります。あるいは、モデルがサンドボックスの別のコンポーネントで別の欠陥を発見した可能性もあります。
この種の作業を実行するときに存在する必要がある他の防御層に関する話のこの部分には、何かが非常に間違っています。存在していても失敗したとしたら、それは大きな話です。それらが存在していても無効になっていたとしたら、それは大きな話です。誰もそれらを使用しようと思わなかったとしたら、それは大きな話です。本当の話は何ですか？
「サイバー脆弱性をテストすることが目的だったので、この評価ではこれらの展開の安全対策が意図的に有効にされなかった」と言うだけでは十分ではありません。オブザーバーモデルが評価を妨げる正当な理由がないからです。もしかしたらそのようなモデルがあったのですが、評価を中止してしまったのでしょうか？しかし明らかに、有効な評価を停止するオブザーバーを、評価の要件外のアクションを停止するオブザーバーに置き換えることができます。この区別は、この保護が偽陽性/偽陰性のバランスを取ることができないことを認識するのはそれほど難しいことではありません。
norabble は読者に支えられた出版物です。新しい投稿を受け取り、私の仕事をサポートするには、無料または有料の購読者になることを検討してください。
1 シェア この投稿に関するディスカッション コミュニケーション

ents リスタック トップ 最新のディスカッション 投稿はありません

## Original Extract

An observer model should have stopped the Hugging Face attack. Why didn’t it exist?

If the Model Escapes, Who Pulls the Kill Switch?
norabble
Subscribe Sign in An OpenAI Model Escaped Its Sandbox. Where Was the Observer?
OpenAI accidentally hacked Hugging Face. A basic layer of defense appears to have been missing, and no one has explained why.
Ryan Baker Jul 23, 2026 1 Share Sometime in the week before July 16th, Hugging Face was attacked by an OpenAI model that was under evaluation . The attack itself wasn’t particularly harmful, but the conceptual implications of the event are significant.
There are three significant aspects I’d highlight. The first is a demonstration of offensive cybersecurity capabilities of current models. That shouldn’t be too shocking if you’ve been paying attention, but the event demonstrates it in a way that’s more clear than prior signals, so likely awareness is going to grow from this event.
norabble is a reader-supported publication. To receive new posts and support my work, consider becoming a free or paid subscriber.
The second is model alignment. Alignment refers to the degree to which AI models do what we want them to. I’ll spare a great dive here as defining what we want itself would be a long dissertation where a great deal of content already exists. But for this case, we can generally agree that it was no one’s intention that the model hack itself out of multiple protection layers to reach the internet, and then from there into Hugging Face’s systems.
There are a few things that make this a bit less than a sign to retreat to the bunker, which I should mention before going deeper. One is that the models being tested had protections turned off, to allow them to test their cybersecurity capabilities. In that sense, this isn’t fully demonstrable of what a user could intentionally or accidentally trigger from usage. This also is an unreleased model, so not available to general users, or really anyone other than OpenAI.
That said, a third point is, and the one I’d like to dive deeper on, I’m worried about the ability for this to happen. It seems to me as if there is a reasonable protection, an observer model, that OpenAI either didn’t use, or failed at. Right now, I’m assuming they didn’t use it because I can’t find any mention of it, but I’m also aware there are many details I’m not privy to.
In case it’s not clear what I’m suggesting by observer model, here’s an example scenario in action. The model under evaluation decides it should escape the sandbox. It starts to attempt this, which does not get caught at a per action level because the classifier safeguards – which would have disrupted the evaluation – are disabled. But then it does escape the sandbox. An observer model watching the evaluation run should be able to look at that state, and say, “This should not have happened. I don’t know how this model is outside its sandbox, but it shouldn’t be outside, we should shut this down.”
If a model is in evaluation for cybersecurity purposes, and you’ve disabled classifiers, you should watch it closely. To carry out these actions, the model had to take several steps, achieving results at each step and then moving on to exploiting the rewards of those steps further to explore, and escalate again. OpenAI’s security team did eventually respond, so they weren’t entirely asleep. But the timing shows they were slower than Hugging Face. They had enough time to try commercial models, which refused to help, then next plan a response with GLM 5.2, and then successfully contain and stop the intrusion.
OpenAI should have had several structural advantages here. First, unaligned activity occurred before any activity was directed at Hugging Face. Second, OpenAI should already have a kill switch for evaluations. There should be no need to devise novel plans to contain and stop activity. The fact that Hugging Face was the actor that stopped the intrusion, rather than taking note of the intrusion stopping mysteriously on its own, indicates OpenAI was the slower responder despite the structural advantage.
An observer model is a little different than a classifier. Classifiers typically evaluate individual actions. An observer would look at trajectories. OpenAI discusses similar new safeguards (“From considering single actions to considering whole trajectories”) in a recent post. If you look at each individual action, it might look like the type of thing a model would need to do in cybersecurity evaluation. But as a whole, some outside source should have been able to reason upon what was occurring. If not a human involved in the evaluation process, an observer model.
I understand why it may have taken a while to deploy this type of safeguard in production, where volumes are higher, where models have already been evaluated, and where classifiers are restrictive. In production, observer models are less effective because malicious users can “distribute” their requests into multiple sessions , and so observing a trajectory first requires coalescing these threads together, which Anthropic has done . A safety mechanism for a cybersecurity evaluation sandbox, does not encounter that challenge. Monitoring your own model in a sandbox is significantly less challenging, especially if it’s not attempting to evade detection. There’s no evidence here that there were attempts to evade detection, just attempts to perform inappropriate actions.
So, is trajectory monitoring by an observer model not standard practice in an offensive cybersecurity evaluation sandbox? Certainly you’d do this during internal evaluation of models that you clearly do not know the maximum capabilities of? If it is and the model evaded detection, that would be an even more significant story. I have to assume that instead there was simply no observer model to evade. Or if, as looks to be the case, it didn’t attempt evasion, it would have escaped the sandbox, been detected, and promptly been shut down before the Hugging Face part of the action had begun.
Why is that not the story we’re reading? But instead we’re reading about Hugging Face contacting law enforcement, scrambling to get a model to counter and analyze and OpenAI talking about the flaw in the proxy?
This is the equivalent of putting a tiger in a cage, and then checking that the tiger stays in the cage. It’s not a particularly rocket science design. Tuning it might take effort, and hardening against evasion more, but it doesn’t seem like we’ve progressed to those requirements. Rather it looks like OpenAI lacked the most basic version. Am I missing something?
Others, such as The Zvi, have pointed out that the core, the most critical fix is in the alignment training .
The intent is the issue. Control strategies and supervision are good parts of a defense-in-depth strategy, we should totally use such strategies. That helps mitigate failure. But that strategy also has to include actually aligning the models, or you lose. And by lose, in the long term, I mean things up to and likely including loss of control over the future and everyone dying.
The model should internally understand these actions aren’t aligned with its mission, and thus not want to take them. I’d fully agree there. I bring up this specific defense in depth as an addition to core corrections. Those corrections are the hardest part, defense in depth is an easier, but incomplete part. Testing alignment is hard, and protections around that testing are necessary, so this isn’t irrelevant to the alignment conversation either.
OpenAI’s initial response has been about a few defense in depth responses, but they are very specific, non-generalizable fixes to the sandbox at a software level, fixing the flaw that the model exploited. That might have avoided this incident, or the model might have found a different flaw in a different component of the sandbox.
Something is very wrong about this part of the story about other layers of defense that should exist when performing this type of work. If they exist but failed, that’s a big story. If they exist but were disabled, that’s a big story. If no one thought to use them, that’s a big story. What’s the real story?
It’s not good enough to say “These deployment safeguards were intentionally not enabled during this evaluation because it was aimed at testing cyber vulnerabilities.” as there isn’t a justification why an observer model would interfere with evaluation. Maybe there was such a model but it would have stopped the evaluation? But clearly you can replace an observer that would stop a valid evaluation with one that would stop actions outside the requirements of evaluation. That distinction is not so hard to recognize that this protection would be incapable of balancing false-positive/false-negative.
norabble is a reader-supported publication. To receive new posts and support my work, consider becoming a free or paid subscriber.
1 Share Discussion about this post Comments Restacks Top Latest Discussions No posts
