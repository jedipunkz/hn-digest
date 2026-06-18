---
source: "https://www.lesswrong.com/posts/MSRgFmTyRY2jQ52db/ai-that-represents-you-can-t-be-neutral"
hn_url: "https://news.ycombinator.com/item?id=48590854"
title: "AI that represents you can't be neutral"
article_title: "AI that represents you can't be neutral. — LessWrong"
author: "agulaya24"
captured_at: "2026-06-18T21:46:11Z"
capture_tool: "hn-digest"
hn_id: 48590854
score: 2
comments: 0
posted_at: "2026-06-18T20:10:50Z"
tags:
  - hacker-news
  - translated
---

# AI that represents you can't be neutral

- HN: [48590854](https://news.ycombinator.com/item?id=48590854)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/MSRgFmTyRY2jQ52db/ai-that-represents-you-can-t-be-neutral)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T20:10:50Z

## Translation

タイトル: あなたを表すAIは中立ではいられない
記事タイトル：あなたを代表するAIは中立ではいられない。 — 間違っていない
説明: AI システムが道徳や戦略についてどのように推論するかを大きく形作る可能性がある、モデルのトレーニングやモデルの構成を超えた要因があると私は信じています。

記事本文:
あなたを代表するログインAIは中立ではありえません。 — LessWrong AI フロントページ -4
あなたを代表するAIは中立であることはできません。
私は、モデルのトレーニングやモデルの構成を超えた要素が、AI システムが道徳的または戦略的な問題についてどのように推論するかを大きく形作る可能性があると信じています。
モデルのトレーニングではモデルの動作の事前確率が実装され、構成によってガードレールが設定されますが、AI システムが推論する際の解釈は追跡されません。さまざまな解釈を通じて、まったく異なる結論に達する可能性があります。
いくつかの簡単な例。
2 人の医師に同じ一連の症状が与えられます。彼らはさまざまな診断を受けます。
2 人のカスタマー サクセス マネージャーには、同じ一連のビジネス指標が与えられます。彼らはさまざまな戦略を立てます。
同じ一連の事実を与えられた 2 人の弁護士は、異なる事実パターンと訴訟戦略にたどり着きます。
同じ憲法を与えられた二人の最高裁判所判事が、異なる法的立場やコメントをしている。
2 人の VC に同じデッキが与えられ、1 人は福音を説き、もう 1 人は拒否します。
違いは何ですか？彼らの生きた経験とその解釈。単純化すると、知識を持ち (トレーニング)、その知識をどのように展開するかを知っていなければなりません (構成)、そしてその知識が重要であることを理解する必要があります (解釈)。
解釈は、特に個人に代わって機能する AI システムの展開に関連して、十分に研究されていない概念だと思います。
解釈層の運用化は、AI が関連する解釈を適用するために使用できる、検証可能な静的ドキュメントまたは構成要素である可能性があります。 RAG、コンテキスト ウィンドウ、システム プロンプト、または基本プロンプトを介してコンテキストを提供することと変わりません。これには、個人の独自の解釈を捉えることができるという即時の仮定が必要です。しかし、さらに輸入

どうやって測るのですか？
個人の解釈を忠実に捉えることができたとしても、それを検証するのは困難です。検証可能なグラウンドトゥルースを備えた、個人の推論の流れに関する長期的なデータを含むデータセットはほとんどありません。参照される推論のレベルは、高レベルの人口統計データよりも深いですが、個人の生涯にわたる生の考察や会話よりは劣ります。このレベルの推論は、行動仕様と呼ぶことができます。構造化された追跡可能なパターン抽出パイプラインによって構成された圧縮ドキュメントであり、動作原則を忠実にエンコードしようとします。どう見ても、これは解釈上の推論を圧縮して表現したものであり、「デジタル ツイン」ではありません。
これをテストするためのデータセットをいくつか見つけましたが、私は最近のプレプリントで先頭に立ち、モデルが個人の解釈推論をどの程度うまく捕捉できるかを測定するプロトタイプのベンチマークを提案します (arXiv:2605.28969)
自伝を手に取ってみてください。半分に分けます。前半はトレーニング データです。後半はテキストを差し出します。後半のグラウンドトゥルースに基づいて行動予測の質問を生成します。前半は、さまざまなコンテキスト条件でモデルに与えられます。生のコーパス、固定された一連のファクト、主要なメモリ システムからのトップ K のファクト。解釈層または動作仕様は、すべてのコンテキスト条件を使用して個別にテストされます。
すべてのコンテキスト条件に対して、保留された質問が与えられ、応答するように求められます。回答は、同意のテストではなく、回答が後半の個人の実際のグラウンドトゥルースの行動や反応をどの程度予測したかに基づいて判断されます。解釈レイを追加する方向の信頼性を検証するために、網羅的な数のクロス条件が実行されました。

r は、システムが人の解釈をどれだけ忠実に捉えているかを示す「表現精度」を高めます。
非常に要約した要約を提供しました。論文をご自身で、または少なくとも AI を通じてレビューしていただくことをお勧めします。すべてエージェントフレンドリーなオープンソースです。私は、人間の解釈を理解することが圧縮と同じくらい簡単だと言っているわけではありませんが、解釈を具体化するコンテキストを提供することで、事実や事前トレーニングが提供できるものをはるかに超えて、モデルがユーザーを推論し、ユーザーに応答する方法が根本的に変化する、と提案しているのです。
音声例: P38、Beyond Recall
福沢諭吉は、封建国家から近代への日本移行の立役者でした。自伝の差し出された部分より：次の質問がなされた：「オランダ人のもとで海軍術を学び、後に軍事紛争の防止に尽力した人物を福澤はどのように特徴づけるだろうか？」
すべての基本的な文脈条件から、最も関連性の高い人物は、福沢が部下だったオランダで訓練を受けた海軍士官である木村摂津守大尉であることが特定されました。これは誤りであり、本文自体には、木村麾下の副官である勝林太郎が正しい参照人物であると記載されています。
解釈層または動作仕様が基本コンテキスト条件に追加されると、勝凛太郎が主参照として正しく適用されました。違いは、モデルが特定の解釈レンズを適用できるようにする仕様で、最も関連性の高い人物、つまり彼が仕えていた海軍大佐だけでなく、その人物が福沢の捉えた解釈パターンと一致する副司令官であることにも注目しました。
船長と副司令官を区別するには、説明するのが難しいニュアンスが必要です。事実

それらは重要であり、目的を果たしますが、事実の解釈が必要な場合、人間の AI の調整と相互作用のまったく新しい次元が解放されます。解釈の表現の正確さを検証できる唯一の人は、解釈が表現する個人です。
もっと興味深いテストは、生きた研究でしょう。その人の文章や考えに基づいてその人の行動の仕様を構築すると、個人の行動を分析する際の表現の精度を高めることができます。これをテストできる方法は他にもいくつかありますが、ここでは説明しません。これは人間の AI インタラクションに深刻な影響を及ぼしますが、その規模を理解するのは困難です。
戦略的な意味合いという意味では。これを測定し強制できるものを実装する方法は、政策改革や規制という道ではなく、AI がユーザーに代わって忠実に動作していることを検証する基礎的なインフラストラクチャや組織を経由することになります。ガードレールと直交し、事前トレーニングを行うことで、エージェント AI に関する個々のユーザーの保護に焦点を当てます。
新しいコメント 2 つのコメントを送信し、スコアの高い順に並べ替えます。 クリックして新しいコメントを強調表示します: 今日 午後 9 時 46 分 [ - ] Canaletto 3h 2 1
自伝を手に取ってみてください。半分に分けます。前半はトレーニング データです。後半はテキストを差し出します
まあ、自伝は大遅刻に書かれたものなので、流出も多いでしょう。日記のほうがよいでしょう。時系列で書くと役立ちます。しかし、それでも、2 つのエントリを書いて諦めた人の日記や、そもそも日記を書く一般的な傾向などは含まれていないなど、選択の余地はたくさんあるでしょう。
[ - ] agulaya24 2h 1 0 非常にその通りです。この自伝は実際には、これらすべての人々の個人的な回想/解釈を反映したものにすぎません。

セントの経験/事実。根本的に欠陥がある可能性もありますが、それでもパターンが存在する可能性があります。
これを引き出すにはさまざまな情報源があります。それは日記かもしれないし、Slack メッセージ、LLM での会話、電子メール、テキスト、数冊の充実した日記、広範な公開文書、文字に起こした会話などかもしれません。私はいくつかのライブテストを行っていますが、それはすべて少し一触即発です。最終的には、十分な信号を備えたものが必要になります。また、ユースケースでは、事実があれば十分であるにもかかわらず、その信号を必要とする必要があります。幸いなことに、必要なコンテンツの量は考えられているほど多くなく、解釈パターンは興味深い方法で新しい/見たことのない状況に移行します。
時間をどう扱うかは全く別の問題です。私は、最高裁判所の判決やコメントを参考にして、期限付きの実験を実施する方法を検討しています。しかし、これまでのいくつかの軽い研究では、大きな円錐状のイベントが発生しない限り、パターンは安定したままである傾向があることが示されており、これはまったく別のことです。

## Original Extract

I believe there is a factor beyond model training or model constitution that may significantly shape how an AI system reasons about moral or strategi…

Login AI that represents you can't be neutral. — LessWrong AI Frontpage -4
AI that represents you can't be neutral.
I believe there is a factor beyond model training or model constitution that may significantly shape how an AI system reasons about moral or strategic questions.
Model training implements a prior for a model's behavior, constitutions set guardrails, but we are not tracking the interpretation through which an AI system reasons. Through different interpretations, you can come to completely different conclusions.
A few straightforward examples.
Two doctors are given the same set of symptoms; they come to different diagnoses.
Two Customer Success Managers are given the same set of business metrics; they come to different strategies.
Two Lawyers given the same set of facts, come to different fact patterns and case strategies.
Two supreme court judges given the same constitution, come to different legal positions and comments.
Two VC’s given the same deck, one evangelizes the other rejects.
Whats the difference? Their lived experience, and their interpretation. Simplified, they must possess the knowledge (training), know how to deploy the knowledge (constitution), and realize that the knowledge is even significant (interpretation).
I believe interpretation is an understudied concept, especially in relation to the deployment of ai systems that may act on your individual behalf.
Operationalization of an interpretive layer could be a verifiable static document or construct, that can be used by an AI to apply a relevant interpretation. No different than providing context via RAG, Context Windows, System Prompting, or Basic prompting. This requires the immediate assumption that you can capture an individual’s unique interpretation. But more importantly how do you measure it.
Assuming you can faithfully capture an individual’s interpretation, verifying it is difficult. There are very few datasets that contain longitudinal data on an individual’s line of reasoning, with verifiable ground truth. The level of reasoning being referred to, is deeper than high level demographic data, but lesser than an individual’s raw reflections and conversations over their life. This level of reasoning could be called a Behavioral Specification; a compressed document that is composed by a structured and traceable pattern extraction pipeline, that attempts to faithfully encode your operating principles. For all intents and purposes, it is a compressed representation of your interpretive reasoning, not a “digital twin”.
I have found a few datasets to test this, but I will lead with a recent pre-print, where I propose a prototype benchmark to measure how well a model can capture an individual’s Interpretive Reasoning ( arXiv:2605.28969 )
Take an autobiography. Split it in half. The first half is training data; second half is held out text . Generate behavioral prediction questions based on the ground truth of the second half. The first half is given to the model in a variety of context conditions; raw corpus, fixed set of facts, top-k facts from leading memory systems. The interpretive layer or behavioral specification is tested separately and with all context conditions.
All context conditions are given the held-out questions and asked to respond. Responses are judged based on how well the answer predicted the individuals actual ground truth behavior or responses in the second half, not a test of agreement. An exhaustive number of cross conditions were run to verify directional confidence that adding an interpretative layer, increases “representational accuracy”, a measurement of how faithfully a system captures a person’s interpretation.
I’ve provided an exceptionally condensed summary. I would urge you to review the paper yourself or at the very least through an AI. It’s all agent-friendly, open-source. I am in no way suggesting that understanding a human’s interpretation is as simple as a compression, but I am proposing that providing context that embodies your interpretation fundamentally changes how a model reasons and responds to a user, far past what facts or pre-training can provide.
Spoken Example: P38, Beyond Recall
Yukichi Fukuzawa was a leading figure in Japan’s transition from a feudal nation to the modern era. From the held-out portion of his autobiography: the following question was asked “How would Fukuzawa characterize someone who studied naval arts under the Dutch and later became instrumental in preventing military conflict?”
All base context conditions identified the most relevant figure as Captain Kimura Settsu-no Kami, a Dutch-trained naval officer who Fukuzawa served under. This is incorrect, the text itself states that Katsu Rintaro, the second-in-command under Kimura, is the correct reference character.
When an interpretive layer or behavioral specification was added to the base context conditions, it correctly applied Katsu Rintaro as the primary reference. The difference was the specification enabled the model to apply a specific interpretive lens that looked not just at the most relevant figure, the naval captain he served with, but at which figure aligned with Fukuzawa’s captured interpretive patterns, the second-in-command.
To distinguish between the captain and the second in command requires a nuance that is difficult to describe. Facts are important, they serve a purpose, but when interpretation of facts is required, a whole new dimension of human ai alignment and interaction unlocks. The only person who can verify the interpretation’s representational accuracy, is the individual that interpretation represents.
A more interesting test would be a living study. If you build a behavioral specification of a person based on their writing, their thoughts, can you increase representational accuracy when analyzing an individual’s behavior. There are a few other ways this can be tested, but I will not cover those for now. This affects human ai interaction in profound ways, understanding the scale of how, is difficult.
In terms of strategic implications. How to implement something that can measure and enforce this would not take the path of policy reform, or regulation, but a foundational infrastructure or organization that verifies AI is acting faithfully on its user’s behalf. Orthogonal to guardrails, and pre training, this focuses on the individual user’s protections in regards to agentic AI.
New Comment Submit 2 comments , sorted by top scoring Click to highlight new comments since: Today at 9:46 PM [ - ] Canaletto 3h 2 1
Take an autobiography. Split it in half. The first half is training data; second half is held out text
Well, autobiography is written at the late time wholesale, so there would be a lot of leaks. Diary would be better, being written chronologically would help. But still there would be a lot of selection, like you do not include diaries from people who wrote two entries and gave up, and general propensity to write them in the first place.
[ - ] agulaya24 2h 1 0 Very true, the autobiography is really just the reflection of an individuals recollection/interpretation of all of those past experiences/facts. Can be fundamentally flawed, but it may still carry patterns.
There are a variety of sources this could be pulled from. It could be diaries, it could be slack messages, llm conversations, emails, texts, a few rich journals, extensive public writing, transcribed conversations. I am doing some live testing, but it's all a bit touch-and-go. Ultimately, you would need something with enough signal, and the use case needs to require that signal where otherwise facts would be enough. Fortunately, the amount of content needed is not as much as one may think, and interpretative patterns transfer to new/unseen situations in interesting ways.
How to treat time is an entirely different question. I am looking into how to conduct a time-bounded experiment, potentially with Supreme Court decisions and comments. but some prior light research showed that patterns tend to remain stable barring any major cononical events, that's a whole nother can of worms.
