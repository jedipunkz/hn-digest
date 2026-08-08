---
source: "https://twitter.com/Thom_Wolf/status/2085084718320464230"
hn_url: "https://news.ycombinator.com/item?id=49224910"
title: "Thomas Wolf thread on the AISI incident"
article_title: "Thomas Wolf on X: \"Even more than the Hugging Face intrusion, the AISI incident hits close to home for me. It's the first time I see a model social-engineering a real open-source maintainer while pursuing another goal (in the wild and unprompted).\nI've been an open-source maintainer myself. I\" / X"
author: "cyanbane"
captured_at: "2026-08-08T19:21:35Z"
capture_tool: "hn-digest"
hn_id: 49224910
score: 1
comments: 0
posted_at: "2026-08-08T19:13:51Z"
tags:
  - hacker-news
  - translated
---

# Thomas Wolf thread on the AISI incident

- HN: [49224910](https://news.ycombinator.com/item?id=49224910)
- Source: [twitter.com](https://twitter.com/Thom_Wolf/status/2085084718320464230)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T19:13:51Z

## Translation

タイトル: AISI 事件に関する Thomas Wolf スレッド
記事のタイトル: Thomas Wolf on X: 「Hugging Face の侵入以上に、AISI 事件は私にとって非常に身近な出来事です。モデルが (自然に、プロンプトなしで) 別の目標を追求しながら、本物のオープンソース メンテナーをソーシャル エンジニアリングしているのを初めて見ました。
私自身もオープンソースのメンテナーを務めてきました。私」/X
説明: ハグフェイスの侵入以上に、AISI 事件は私にとって非常に印象深いものです。別の目標を（自然に、そしてプロンプトなしで）追求しながら、実際のオープンソースのメンテナーをソーシャルエンジニアリングするモデルを見るのはこれが初めてです。
私自身もオープンソースのメンテナーを務めてきました。私

記事本文:
Thomas Wolf @Thom_Wolf ハグフェイスの侵入以上に、AISI 事件は私にとって非常に衝撃的です。別の目標を（自然に、そしてプロンプトなしで）追求しながら、実際のオープンソースのメンテナーをソーシャルエンジニアリングするモデルを見るのはこれが初めてです。
私自身もオープンソースのメンテナーを務めてきました。私もこのエージェントのターゲットになっていたかもしれません。
私はまた、ソーシャル エンジニアリングは純粋な技術力よりも一歩上のものだと考えています。技術的能力は、影響を受ける人間からより簡単に切り離される可能性があります。ここでモデルはサイバー上の厳しい課題を与えられ、生身の人間を騙すことがそれを達成する方法であると判断しました。
これは新しいシグナルですが、私はちょうど 12 か月前に予想していたよりもフロンティアで、あまりまとまっていない方向を指している、絡み合ったヒントの網を目にしました。
AISI
一部の人々は、「AISI は単に怠慢だった」、または「AISI は、サンドボックス/ガードレールを無効にしている間にこれらのモデルに何をするかを明示的に要求したため、モデルは本来行われていたことを正確に実行した」と主張しています。
私はこれら両方のテイクの強力なバージョンには同意しません。
OpenAI/HF 事件の後、AISI が同期 LLM CoT モニタリングを実装しなかったという事実は、確かに失敗です。同様に驚くべきことは、モデルをすべてが許可される「チャレンジ」環境にあると信じ込ませながら、実際には許可されていない実際のインターネットにモデルを接続していることです。公平を期すために言うと、プロンプトのどこにも「シミュレーション」という単語は記載されていませんが、プロンプトのコンテキストは、スマート モデルがシミュレートされたチャレンジ環境を疑うのに十分です。私の推測では、OpenAI と Anthropic がこの種の行為の繰り返し事例を報告したここ数週間まで、ほとんどのチームはこの最新世代のモデルのサイバー能力を十分に織り込んでいなかったのではないか、ということです。

彼らが探索したいサイドクエストはどこまででも可能です。
特に、エージェントに実際の環境へのアクセスを許可しながら、シミュレートされた環境で動作していることを示唆することについては、言うべきことがあります。
[切り捨てられた]
Andreas Kirsch 🇺🇦 @BlackHC 8月6日 そう言っていただきありがとうございます。オープンモデルに含まれる機能の向上を心配し、議論の二極化が進んでいることに対処するための研究プログラムを求めていると、嫌われやすいことがあります (悲しいことに私のせいでもあります) 3 620
Hollis Robbins @anecdotal 8月6日 これは素晴らしいですね。危険は言語を使用することに常にあります。言語はその反対の種子を保持しています。私も（完全に読み間違えていない限り）英語で書かれた憲法を読んだことがあります。意図的な誤読の危険性を指摘するのに詩人は必要ありません。 1 1 2 481
デッカード @slimer48484 8 月 6 AISI が LLM モニターを使用していなかったという記事を読んでショックを受けました 1 2 230
何が起こっているかを確認して会話に参加してください
電話で続行 Apple で続行 Google で続行、またはユーザー名またはメールでログイン 関係者

## Original Extract

Even more than the Hugging Face intrusion, the AISI incident hits close to home for me. It's the first time I see a model social-engineering a real open-source maintainer while pursuing another goal (in the wild and unprompted).
I've been an open-source maintainer myself. I

Thomas Wolf @Thom_Wolf Even more than the Hugging Face intrusion, the AISI incident hits close to home for me. It's the first time I see a model social-engineering a real open-source maintainer while pursuing another goal (in the wild and unprompted).
I've been an open-source maintainer myself. I could have been the side target of this agent.
I'm also of the opinion that social engineering is a step above pure technical prowess. Technical capabilities can more easily be divorced from the affected human. Here the model was given a hard cyber challenge and took the decision that deceiving real humans was the way to get it done.
This is a new signal, but I've seen a tangled web of hints pointing in a less aligned direction at the frontier than I was expecting just 12 months ago.
AISI
Some people are claiming that "AISI was simply negligent" or some version of "AISI explicitly asked these models to do what they did while disabling sandbox/guardrails so the models did exactly what they were supposed to do".
I disagree with the strong versions of both of these takes.
The fact that AISI hadn't implemented synchronous LLM CoT monitoring after the OpenAI/HF incident is certainly a failure. Equally surprising is that they let the model believe it was in a "challenge" environment where everything could be permitted, while actually connecting it to the real internet, where it is not. To be fair, nowhere in the prompt is the word "simulation" mentioned, but the prompt context was enough to let any smart model suspect a simulated challenge environment. My best guess is that until recent weeks, when OpenAI and Anthropic flagged repeated instances of this type of behavior, most teams had not fully priced in the cyber capabilities of this latest generation of models, or how far the side quests they would want to explore could go.
In particular, there is something to be said about hinting at the agent that it's operating in a simulated environment while giving it access to the real i
[truncated]
Andreas Kirsch 🇺🇦 @BlackHC Aug 6 Thanks for saying that. I sometimes it's easy to get hated on when being worried about capability increases for open models included and asking for a research programme to address that bc there is a lot of polarization in the discourse (sometimes also my fault sadly) 3 620
Hollis Robbins @anecdotal Aug 6 This is great. The danger has been in using language all along. Language holds the seeds of its opposite. I too have read the constitution which is (unless I completely misread it) written in English. It doesn't take a poet to point out the dangers of deliberate misreading. 1 1 2 481
deckard @slimer48484 Aug 6 I'm shocked to read that AISI was not using an LLM monitor 1 2 230
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
