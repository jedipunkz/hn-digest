---
source: "https://blog.cryptographyengineering.com/2026/06/09/apples-siri-ai-or-more-shouting-into-the-void-about-private-agents/"
hn_url: "https://news.ycombinator.com/item?id=48466088"
title: "Apple's Siri-AI, or more shouting into the void about \"private\" agents"
article_title: "Apple’s Siri-AI, or more shouting into the void about “private” agents – A Few Thoughts on Cryptographic Engineering"
author: "cdrnsf"
captured_at: "2026-06-09T21:40:06Z"
capture_tool: "hn-digest"
hn_id: 48466088
score: 3
comments: 0
posted_at: "2026-06-09T19:10:52Z"
tags:
  - hacker-news
  - translated
---

# Apple's Siri-AI, or more shouting into the void about "private" agents

- HN: [48466088](https://news.ycombinator.com/item?id=48466088)
- Source: [blog.cryptographyengineering.com](https://blog.cryptographyengineering.com/2026/06/09/apples-siri-ai-or-more-shouting-into-the-void-about-private-agents/)
- Score: 3
- Comments: 0
- Posted: 2026-06-09T19:10:52Z

## Translation

タイトル: Apple の Siri-AI、あるいは「民間」エージェントについて虚空に向かって叫ぶ
記事のタイトル: Apple の Siri-AI、あるいは「プライベート」エージェントについて虚空に向かって叫ぶ – 暗号工学に関するいくつかの考察
説明: 昨日、Apple は、Siri エコシステムに本物の AI を導入するための大きな一歩を発表しました。ほとんどの点で、これは良いことであり、避けられないことです。Siri は世界で最も広く使用されている音声エージェントの 1 つであり、それが最悪でなければ良いのですが。 Apple がフロンティアモデルで自社の機能を強化するという考えは、
[切り捨てられた]

記事本文:
Apple の Siri-AI、あるいは「民間」エージェントについて虚空に向かって叫ぶ – 暗号工学に関するいくつかの考察
コンテンツにスキップ
ホーム
メニュー
Apple の Siri-AI、あるいは「民間」エージェントについて虚空に向かって叫ぶ
私の学術ウェブサイト
ブルースカイ
マストドン
ツイッター
人気の投稿
便利な暗号リソース
ビットコインチップジャー
クリプトパルの課題
応用暗号研究: 委員会
暗号工学ジャーナル
(このブログとは関係ありません)
検索:
人気の投稿とページ
暗号化された推論について話しましょう
Apple の Siri-AI、あるいは「民間」エージェントについて虚空に向かって叫ぶ
Apple 様: 今すぐ「消えるメッセージ」を iMessage に追加してください
ゼロ知識証明: 図解による入門書
Telegram は本当に暗号化されたメッセージング アプリですか?
WhatsAppの暗号化、訴訟、そして多くのノイズ
今週の攻撃: RC4 は TLS で壊れているようです
匿名認証情報: 図解による入門書
匿名資格情報: 図解による入門書 (パート 2)
昨日、Apple は、Siri エコシステムに本物の AI を導入するための大きな一歩を発表しました。ほとんどの点で、これは良いことであり、避けられないことです。Siri は世界で最も広く使用されている音声エージェントの 1 つであり、それが最悪でなければ良いのですが。 Apple がフロンティアモデルで自社の機能を強化するという考えは、「かどうか」というよりは、「いつ、誰が」という問題でした。
正体は Google であることが判明: Apple は、Google Gemini モデルを組み合わせて、Google の Confidential Inference とプライベート ホスティング用の Apple 独自のプライベート クラウド コンピューティングを組み合わせて使用​​するようです。これらのシステムは、クエリを処理し、デバイスからのプライベート データを評価します。 Apple のマーケティングでは次のように利点を宣伝しています。
まず、携帯電話にはすでにあなたに関するコンテキスト (個人情報、スケジュール、電子メール、テキスト メッセージなど) が含まれているためです。

es — AI 対応の Siri は、実際のリクエストに対して外部 LLM よりも役立つ回答を提供できる可能性があります。来週の誕生日パーティーの予約を入れたいですか?理論的には、将来の Siri-AI は、誰が来るのか、どんなケーキが好きなのかをすでに知っているかもしれません。
もちろん、Apple が「コンテキスト」と呼ぶものは、あなたの人生の生のデータでもあります。これはすべてのアプリの極秘データであり、そのデータを処理するためにランダムなアドテク企業 (またはサム アルトマン) に送信するだけでは済みません。ユーザーのコンテキストは保護される必要があり、Apple は自らをプライバシー企業として宣伝しています。
これらの目標の間には緊張感があります。 Apple は、Private Cloud Compute (PCC) と呼ばれるサービスをマーケティングすることでこの問題に対処しました。 PCC は、Apple のデータセンターで実行される一連の「信頼できる」ハードウェア セキュリティ モジュールを使用して、完全に Apple Silicon 上で実行されるプライベート モデル推論システムとして 2024 年に導入されました。このシステムの目標は、データが Apple のハードウェアから決して流出しないようにすることです。データは携帯電話から専用サーバーに暗号化され、応答が携帯電話に届くと消去されます。 PCC のステートレス設計により、(理論上は) データが残らないことが保証され、ハードウェアの設計により、Apple でさえ入力を確認できなくなります。
それ以来、Apple は Google のハードウェアも含めるように PCC を「拡張」しました。正直に言うと、新しい「拡張された」PCC の詳細は少し曖昧だと思います。 Apple は主に、このデータを処理するために Google の既存の機密コンピューティング (Google データセンターで実行されている) に依存するつもりのように思えますが、実際に実行されているモデルを制御するために技術的セキュリティの新しい層を追加しています。いずれにせよ、セキュリティ専門家は、これが Cozy Bear をデータから遠ざけるのに十分であるかどうかについて議論することができます。何

Google や Apple があなたのものにアクセスするのを防ぐにはおそらく十分だろうということは認めますが、そもそもほとんどの人がそれを心配しているのです。
民間エージェントが関与する簡単なシナリオ
エージェントがどのように機能するかを説明するには、ユースケースの例を考えると役立ちます。 6 人分のビジネスディナーを計画していると想像してください。これには、いくつかのサブタスクが含まれます。
参加者のスケジュールを調整し、彼らがいつ街にいて、いつ会えるかを把握する必要があります。
メニューと場所に基づいて適切なレストランを選択する必要があります。これは、参加者の好みについてあなたが知っていることに依存するかもしれません。たとえば、マイクは四川山椒に極度のアレルギーを持っているため、かなりの数の選択肢が除外されます。
時間、料理、場所の制約があるため、実際に適切な場所に 6 人掛けのテーブルがあるレストランを探す必要があります。
最後に、予約を入れ、カレンダーにマークを付け、出席者に通知する必要があります。
以前は、この種のスケジューリングには多大な人的労力が必要でした。 AI エージェントの利点は、理論上、まさに自動化できる種類のプロジェクトであることです。エージェントはまず最近の会話をスキャンして、ステップ (1) と (2) に必要な質問に答えてから、ステップ (3) で説明されている検索を実行します。あなたが同意すれば、ステップ (4) を完了するために必要なカレンダーの招待状やテキスト メッセージを作成することもできます。
最初の、そして驚くべきことではありませんが、これらのタスクで役立つためには、エージェントがコンテキストを持っている必要があるということです。これは、プライベート データへの比較的制限のないアクセスを意味します。招待者の空き状況は、招待者からテキストメッセージが送られてくるのでわかります。マイクのアレルギーについて知っているのは、マイクと話したり、どこかに書き留めたりしたからです。

e. (これは、iMessage、電子メール、連絡先、または個人的なメモを意味する可能性があります。) これらすべてのデータをエージェントに再入力するのは面倒で時間がかかりますが、エージェントの目的は時間を節約することです。勝者のパーソナル アシスタントは、単に賢いという理由だけで勝っているわけではありません。机の隣に座っているパーソナル アシスタントのように、ユーザーが知る必要があることを「すでに知っている」から勝ちます。
もう少し詳しく詳細を掘り下げてみましょう。エージェントはメッセージ データベースをスキャンして、夕食のスケジュールに必要なパラメータを学習する場合があります。あるいは、よりトークン効率の高いシステムでは、メッセージを継続的に読み取り、後で必要になる可能性のある有用な事実を抽出する「メモリ」を保存する可能性があります。どちらも機能的には同等ですが、一方は非常に機密性の高いアーティファクトを生成します。また、役立つ可能性のある一連の事実は非常に広範囲にわたることに留意してください。たとえば、マイクのアレルギーもその事実の 1 つです。しかし、他にもたくさんあります。たとえば、マイクが浮気をしていることを発見したときのプライベートな会話は、システムによって保存またはアクセスされる可能性がある別の事実である可能性があります。メモリがあるかどうかに関係なく、このデータはすべてエージェントの視野内にあるため、どれを操作すべきかをエージェントが認識していることを期待する必要があります。
このデータを入手すると、エージェント (実際にはどこかのデータ センターのサーバー上で実行され、多数のローカル状態とプロンプトを組み合わせた LLM) は、このデータに対して推論を実行して、データを要約するか、クエリ自体に応答する必要があります。ここで、Private Cloud Compute と Confidential Inference がユーザーを保護するように設計されています。これらのテクノロジーの目的は、このデータと推論結果がユーザーのみに制限されるようにすることです。入力と出力は推測したらすぐにワイプする必要があります

それが完了すると、そのいずれかの唯一の残りのコピーが携帯電話に存在するはずです。
これまでのところ、推論を超えて何かをするつもりがない限り、これは説得力のある話だと思います。
プライベート推論は便利ですが、役に立つためにはエージェントが物事と対話する必要があります
推論のみを実行する AI は、個人的なファイルを読み取ることができる人間のアシスタントに似ていますが、それ以外の場合は、窓のない部屋に閉じ込められ、インターネット アクセスも電話もありません。データは完全に安全ですが、アシスタントは、受信メッセージを要約して利用したり、テキスト メッセージの下書きを手伝ったりするなど、最も単純なタスク以外には役に立ちません。 (つまり、Apple Intelligence が現在行っていることです。)
では、実際に物事を遂行できるパーソナルアシスタントを想像してみてください。このアシスタントにはインターネット アクセスが必要です。少なくとも、検索エンジンにクエリを実行する機能、または将来的には Gemini や ChatGPT などの LLM を検索する機能が必要です。タスクの後の手順を実行するには、公開カレンダーの招待をスケジュールし、連絡先と共有するメッセージの下書きを行うために必要になります。このアシスタントは現在では便利ですが、「他人が個人データにアクセスできない」という PCC の素晴らしい保証は、もはや適用できません。データのプライバシーは、もはやシリコンの設計に依存するのではなく、アシスタントの裁量と判断に依存します。
仮定のビジネスディナーに戻りましょう。ステップ (3) を実行するには、エージェントは検索エンジン/LLM にアクセスし、おそらくいくつかのクエリを実行する必要があります。各クエリにより、特定の要件に関する情報が漏洩します。データ漏洩の性質は、「プライベート」エージェントがクエリを作成する際にどれだけ慎重であるかによって決まります。完全に合理的なケースは、モデルが単に一連の有用な事実を収集し、それらをすべてより強力なデータベースにアップロードすることです。

次のように、Gemini、ChatGPT、Claude などの LLM を「オープン」検索します。
「ねえ、LLM 検索エンジン、これは私の出席者とこの会議の目的に関する 30 の詳細な事実のリストです。誰にとっても役に立つレストランを見つけてください。」
これにより、当然のことながら、タスクに関するかなりの量の情報が明らかになり、場合によっては、タスクを完了するために必ずしも必要ではない情報も明らかになります。別の言い方をすると、プライベート推論は完全に機能する可能性があり、依然として有用な (そして収益化可能な) データがパブリック検索エンジンや LLM に流出する可能性があります。これは単に、エージェントがプライバシーを保護しない方法で仕事を行うようにプログラムされているからです。
なるほど、検索エンジンは個人データを学習する可能性があります。だから何？
マイクが四川料理にアレルギーがあることが検索エンジンに知られても、おそらくあまり気にしないでしょう。しかし、本当に気をつけなければならないことがあります。セキュリティ用語では、両方とも異なる敵と関係があります。
最も明らかな「敵」から始めましょう。あなたがマーク・ザッカーバーグやサンダー・ピチャイ、あるいは Apple の広告事業を経営する人物だと想像してみてください。非常に有用なデータの山が携帯電話に保存されている何十億人ものユーザーがいます。このデータは、ターゲットを絞った広告にとって非常に価値があり、生成 AI のおかげで収益性が大幅に高まりつつあります。同時に、このデータの大部分はアクセスできません。これは単に、ユーザーがプライベートな会話をスキャンするという考えを好まないためです。したがって、一部の公開データ (Web ブラウジングなど) にはアクセスできるかもしれませんが、多くのユーザーがデバイスに保存している数年分の親密なプライベートな会話を読むことはできません。
次に、ユーザーの電話にエージェントを導入することを想像してください。そのエージェントはそのすべてのデータにアクセスできるようになります。ユーザーが行うすべてのことにアクセスできるようになります。その仕事をするために、それは文字通りになります

各ユーザーの好みを推測し、それを操作して検索エンジンまたは「検索 LLM」に繰り返しヒットするクエリを作成する必要があります。この検索エンジンを運営している人は誰でも、ユーザーの欲求に関する膨大な量の有用な情報を知ることになります。その一部は、最も親密なプライベートな会話から得られるものです。何年も前に起こって忘れてしまった会話も含まれます。検索エンジンを操作する人がモデルとそのプロンプトを設計する人でもある場合、データ抽出の最良のシナリオが実際に得られます。
エージェントが人々と会話できる場合、見知らぬ人もエージェントに話しかけることができます
Googleが自分たちのことをもっと知るという脅しに肩をすくめる人もいるだろう。私はこの見方には同意しませんが、反論するのは難しいと思います。少なくとも一般の側から見ると、Google は私たちのデータをかなり適切に管理してきました。私の知る限り、私たちの最も親密な Google 検索が (AOL の検索侵害のように) インターネット上に放り出されるような重大なデータ侵害はありませんでした。この点で同社は大いに称賛に値します。
したがって、Google やメタ、Apple が私たちの個人データから私たちについてさらに詳しく知る可能性があるという考えには反対ですが、少なくとも私たちの最も機密な秘密が全世界に公開されることはない可能性があります。これは、プライバシーに最も懐疑的な人でもおそらく気にすることでしょう。

[切り捨てられた]

## Original Extract

Yesterday Apple announced a big step towards deploying real AI in their Siri ecosystem. In most ways this is good and inevitable: Siri is one of the world's most widely-used voice agents, and it would be good if it didn't suck. The idea that Apple would boost its capabilities with frontier models wa
[truncated]

Apple’s Siri-AI, or more shouting into the void about “private” agents – A Few Thoughts on Cryptographic Engineering
Skip to content
Home
Menu
Apple’s Siri-AI, or more shouting into the void about “private” agents
My academic website
BlueSky
Mastodon
Twitter
Top Posts
Useful crypto resources
Bitcoin tipjar
Cryptopals challenges
Applied Cryptography Research: A Board
Journal of Cryptographic Engineering
(not related to this blog)
Search for:
Top Posts & Pages
Let's talk about encrypted reasoning
Apple's Siri-AI, or more shouting into the void about "private" agents
Dear Apple: add "Disappearing Messages" to iMessage right now
Zero Knowledge Proofs: An illustrated primer
Is Telegram really an encrypted messaging app?
WhatsApp Encryption, a Lawsuit, and a Lot of Noise
Attack of the week: RC4 is kind of broken in TLS
Anonymous credentials: an illustrated primer
Anonymous credentials: an illustrated primer (Part 2)
Yesterday Apple announced a big step towards deploying real AI in their Siri ecosystem. In most ways this is good and inevitable: Siri is one of the world’s most widely-used voice agents, and it would be good if it didn’t suck. The idea that Apple would boost its capabilities with frontier models wasn’t so much a matter of if , but a question of when and who .
The who turns out to be Google: Apple looks like it will use some combination of Google Gemini models, combined with Google’s Confidential Inference and Apple’s own Private Cloud Compute for private hosting. These systems will process both your queries and evaluate private data from your devices. Apple’s marketing pitches the advantages as follows:
First, since your phone already has context about you — meaning, your private information, schedules, email, text messages — an AI-enabled Siri can potentially offer more useful answers to your practical requests than external LLMs. Want to schedule a reservation for next week’s birthday party? In theory, a future Siri-AI might already know who’s coming, and what kind of cake they like.
Of course, what Apple calls “ context ” is also the raw data of your life. This is deeply private data from all of your apps, and that data can’t just be shipped to random adtech companies (or Sam Altman) for processing. Your context needs to be protected, and Apple bills itself as a privacy company.
There’s some tension between these goals. Apple has addressed this by marketing a service it calls Private Cloud Compute , or PCC. PCC was introduced in 2024 as a private model inference system that ran entirely on Apple Silicon, using a set of “trusted” hardware security modules running in Apple’s datacenters. The goal of this system is to ensure that your data never leaves Apple’s hardware: it’s encrypted from your phone to a dedicated server, and then it disappears once a response reaches your phone. The stateless design of PCC ensures (in theory) that your data doesn’t linger, and the design of the hardware prevents even Apple from seeing the inputs.
Apple has since “expanded” PCC to encompass Google’s hardware as well. I will confess that I find the details of the new “expanded” PCC just a bit vague. It sounds a lot like Apple is primarily going to rely on Google’s existing confidential compute (running in Google datacenters) to process this data, but they’re bolting on a new layer of technical security to control which models are actually running. In any case: security experts can argue about whether this is good enough to keep Cozy Bear away from your data. What I will grant is that it’s probably good enough to keep Google and Apple from accessing your stuff, which is what most people are worried about in the first place.
A brief scenario involving private agents
To illustrate how agents might work, it’s helpful to consider an example use case. Let’s imagine that you’re planning a business dinner for six people. This involves several subtasks:
You need to juggle the participants’ schedules, know when they’re in town and available to meet.
You need to choose the appropriate restaurant based on menu and location. This might depend on what you know about the participants’ preferences: Mike is wildly allergic to szechuan peppercorn, for example, which rules out quite a few options.
With these time/cuisine/location constraints in place, you’ll need to search for a restaurant that actually has a table for six in the right place.
Finally, you’ll need to book the reservation, mark your calendar, and alert your attendees.
In the past, this type of scheduling required a significant amount of human effort. The beauty of AI agents is that, in theory, this is exactly the sort of project that can be automated. The agent can first scan your recent conversations to answer the questions needed for steps (1) & (2), then it can conduct the searches described in step (3). With a nod from you, it can even author the calendar invites and text messages required to complete step (4).
The first and unsurprising observation is that being useful on these tasks requires your agent to have context , which means: relatively unrestricted access to your private data . You know about your invitees’ availability because they texted it to you. You know about Mike’s allergy because you’ve talked about it with him or jotted it down somewhere. (This could mean iMessages, email, contacts, or personal notes.) Re-entering all of this data into an agent would be annoying and time consuming and the whole point of an agent is to save you time. The winning personal assistant doesn’t win just because it’s smart: it wins because it “already knows” the things you need it to know, like a personal assistant who sits next to your desk.
Allow me to dig into the details just a bit deeper. The agent might scan your messages database to learn the parameters needed to schedule your dinner. Or, in a more token-efficient system, it might read your messages continuously and store a “memory” that distills useful facts that it might need later. Both can be functionally equivalent, but one produces an artifact that may be highly sensitive. And keep in mind that the set of facts that might be useful is very broad. For example, Mike’s allergy is one of those facts. But there are many others. For example, the private conversation you had where you discovered that Mike was having an affair is potentially another fact that could be stored or accessed by a system. Memory or not, this data will all be within the agent’s view, and you’ll have to hope that it knows which one to operate on.
With this data at its fingertips, your agent (which is really an LLM running on a server in a data center somewhere, combined with a bunch of local state and prompting) will need to perform inference over this data, either to summarize it, or to respond to the query itself. This is where Private Cloud Compute and Confidential Inference are designed to protect you. The purpose of these technologies is to ensure that this data, and any inference results, are restricted to you alone. The inputs and outputs should be wiped as soon as inference is done, and the only remaining copy of any of it should exist on your phone.
So far I find this to be a compelling story, as long as you never plan to do anything else beyond inference.
Private inference is nice, but to be useful, agents need to talk to things
An AI that performs only inference is like a human assistant that can read your private files, but is otherwise locked in a windowless room, with no Internet access and no phone. Your data is perfectly safe, but your assistant is worthless for all but the simplest tasks: for example, summarizing inbound messages for your consumption, or helping draft text messages. (In short, what Apple Intelligence does today.)
Now imagine a personal assistant who can actually get things done. This assistant will need Internet access: at minimum the ability to query search engines, or in the future, search LLMs like Gemini or ChatGPT. To accomplish the later steps of our task, you’ll need it to schedule public calendar invites and draft messages to share with your contacts. This assistant is now useful, but the wonderful PCC guarantees of “no private data is accessible to others” are no longer so applicable. The privacy of your data no longer depends on the design of some silicon, but rather, on your assistant’s discretion and judgement.
Let’s move back to our hypothetical business dinner. To accomplish step (3) your agent will need to visit an search engine/LLM, perhaps asking it several queries, each of which leaks some information about your specific requirements. The nature of the data leakage really depends on how cautious the “private” agent is in authoring its queries. A perfectly reasonable case would be for the model to simply collect a series of useful facts, and upload them all to a more powerful “open” search LLM like Gemini, ChatGPT or Claude, as follows:
“ Hey, LLM search engine, here is a list of thirty detailed facts about my attendees and the purpose of this meeting, find me a restaurant that works for everyone. “
This would naturally reveal quite a bit of information about the task, and possibly some information that’s not strictly necessary to get it done. Put differently: private inference can work perfectly, and still useful (and monetizable ) data can still flow outward to a public search engine or LLM, simply because the agent was programmed to do its job in a slightly non-privacy-preserving way.
Ok, so search engines may learn some private data. So what?
You probably don’t care very much if a search engine learns that Mike is allergic to Szechuan food. But there are things you really should care about. In security parlance, they both have to do with different adversaries.
Let’s begin with the most obvious “adversary”. Imagine you’re Mark Zuckerberg or Sundar Pichai, or whoever runs Apple’s advertising business. You have billions of users with piles of deeply useful data stored on their phones. This data is extremely valuable for targeted advertising, something that is about to become wildly more lucrative thanks to generative AI. At the same time, a big chunk of this data is inaccessible, simply because users don’t love the idea of you scanning their private conversations. And so while you might have access to some public data (like web browsing) you can’t read those years worth of intimate private conversations that many users store on their devices.
Now imagine deploying an agent to users’ phones. That agent will have access to all that data. It’ll have access to everything the user does. To do its job, it will literally need to divine each user’s preferences and then operationalize them into queries that will repeatedly hit your search engine or “search LLM”. Whoever operates this search engine will learn a vast amount of useful information about the users’ desires, some of which will come from the most intimate private conversations — even conversations that happened years ago, and that you’ve forgotten about. If the person who operates the search engine is also the person who designs the model and its prompting, then you really have a best-case scenario for data extraction.
If your agent can talk to people, then strangers may talk to it
Some folks will shrug at the threat of Google learning more about them. While I don’t subscribe to this viewpoint, I find it hard to argue with. From the public side, at least, Google has been a reasonably good steward of our data. To my knowledge, there have been no major data breaches where our most intimate Google searches were dumped all over the Internet (in the style of AOL ‘s search breach.) The company deserves a lot of credit for this.
So while I object to the idea that Google or Meta or Apple may learn even more about us from our private data, it’s at least possible that our most intimate secrets won’t be revealed to the entire world. This is something that even the most privacy-skeptical folks probably would care a

[truncated]
