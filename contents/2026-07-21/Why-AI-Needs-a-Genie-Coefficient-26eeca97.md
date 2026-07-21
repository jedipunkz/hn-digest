---
source: "https://spectrum.ieee.org/ai-agent-benchmark"
hn_url: "https://news.ycombinator.com/item?id=48997113"
title: "Why AI Needs a \"Genie Coefficient\""
article_title: "AI Agent Benchmarks Need to Measure User Intent - IEEE Spectrum"
author: "mikelgan"
captured_at: "2026-07-21T20:14:46Z"
capture_tool: "hn-digest"
hn_id: 48997113
score: 2
comments: 1
posted_at: "2026-07-21T19:36:36Z"
tags:
  - hacker-news
  - translated
---

# Why AI Needs a "Genie Coefficient"

- HN: [48997113](https://news.ycombinator.com/item?id=48997113)
- Source: [spectrum.ieee.org](https://spectrum.ieee.org/ai-agent-benchmark)
- Score: 2
- Comments: 1
- Posted: 2026-07-21T19:36:36Z

## Translation

タイトル: なぜ AI には「魔神係数」が必要なのか
記事のタイトル: AI エージェントのベンチマークはユーザーの意図を測定する必要がある - IEEE Spectrum
説明: 研究者らは、AI エージェントがユーザーの意図を本当に実行するかどうかを測定する AI エージェント ベンチマークの新しい指標である Genie 係数を提案しています。

記事本文:
-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
AI エージェントのベンチマークはユーザーの意図を測定する必要がある - IEEE Spectrum
IEEE.org IEEE Explore IEEE 標準 IEEE 求人サイト その他のサイト サイン イン IEEE に参加 AI に「魔神係数」が必要な理由 テクノロジー インサイダー向けのシェア 検索: トピック別に探索 航空宇宙 AI 生物医学 気候 テクノロジー コンピューティング 家庭用電化製品 エネルギー 技術の歴史 ロボット工学 半導体 電気通信 輸送
IEEEスペクトル
テクノロジー関係者向けトピック
アカウントを作成すると、さらに無料のコンテンツや特典をお楽しみいただけます
後で読むために記事を保存するには、IEEE Spectrum アカウントが必要です
インスティチュートのコンテンツはメンバーのみが利用できます
PDF 版全号のダウンロードは IEEE メンバー限定です
この電子ブックのダウンロードは IEEE メンバー限定です
へのアクセス
スペクトル
のデジタル版は IEEE メンバー限定です
以下のトピックは IEEE メンバー限定の機能です
記事に回答を追加するには、IEEE Spectrum アカウントが必要です
アカウントを作成すると、さらに多くのコンテンツや機能にアクセスできます
IEEEスペクトル
、後で読むために記事を保存したり、スペクトル コレクションをダウンロードしたり、イベントに参加したりする機能が含まれます。
読者や編集者との会話。より独占的なコンテンツと機能については、次のことを検討してください。
IEEEに参加する
。
エンジニアリングと応用科学を専門とする世界最大の専門組織に参加し、次の情報にアクセスしましょう。
Spectrum のすべての記事、アーカイブ、PDF ダウンロード、その他の特典を利用できます。
IEEE について詳しくはこちら →
エンジニアリングと応用科学を専門とする世界最大の専門組織に参加し、次の情報にアクセスしましょう。
この電子書籍に加えて、
IEEE スペクトラム
記事、アーカイブ、PDF ダウンロード、その他の特典。
IEEE について詳しくはこちら →
数千の記事にアクセス — Com

完全に無料
アカウントを作成して限定コンテンツと機能を入手してください:
記事の保存、コレクションのダウンロード、
そして
コメントを投稿する
— すべて無料です！フルアクセスと特典については、
購読する
スペクトルに。
AI に「魔神係数」が必要な理由
AIが本当に望むことをしてくれるかどうかの新しい指標を提案
Barath Raghavan は Fastly の著名なエンジニアであり、南カリフォルニア大学の教員です。
ブルース・シュナイアーは、ハーバード大学ケネディスクールとトロント大学のフェロー兼講師です。
ライアン・スヌーク
主要なベンチマークは、AI が何ができるかを測定します。 AI が意図したとおりに実行するかどうか、つまり、AI に実行してほしいことと、AI にどのように実行してほしいかについての暗黙の前提との間の距離を測定するものはありません。私たちは、新しい指標である Genie 係数を提案します。
ある人の要求と別の人の理解の間には、多くの場合ギャップがあります。ほとんどの場合、私たちは一般知識を使ってそれを橋渡しします。たとえば、友達にコーヒーを買ってもらうと、ポットからカップに注いでくれるか、コーヒーショップでコーヒーを買ってくれます。生豆の入った袋を持ってきたり、見知らぬ人からカップをひったくって渡したりはしません。これについては何も指定していません。そんな必要は決してありませんでした。
この修正は、タスク、質問、意図をより適切に指定することだけだと考える人もいるかもしれません。しかし、1987 年、テリー ウィノグラードとフェルナンド フローレスは、AI に関する独創的な本の中で、それがうまくいかない理由を簡潔に次のように述べています: 「Q: 冷蔵庫に水はありますか? A: はい。Q: どこに? わかりません。A: ナスの細胞の中です。」人間の言語では、欲求や欲望は常に明確化されていません。すべての注意事項、すべての制限、すべての例外をリストすることは不可能です。
では、意図を特定できない場合、どうやってコミュニケーションを取るのでしょうか?合理的な人は合理的な推測をすることができるからです。たとえ欲望や欲求が

常に指定が不十分な場合、有能な人は通常、それを正しく理解するのに十分な文脈を知っているか、説明を求めることを知っています。言語学者はこれを語用論と呼んでいます。つまり、意味は言葉と状況にあり、また、以前のすべてのコミュニケーション、共有文化、人間の生来の行動にもあります。
コーヒーを求められた AI エージェントは、コーヒー農園を購入したり、1 杯のコーヒーを注文して 3 週間以内に配達する可能性があります。
もちろん、いつもうまくいくわけではありません。アイスコーヒーが飲みたかったときに友人がホットコーヒーを持ってきたり、トルココーヒーが飲みたかったときにイタリアンコーヒーを持ってきたりするかもしれません。二人の年齢、文化、背景が異なるほど、その要求は何らかの形で誤解される可能性が高くなります。
この状況は、人間からのリクエストをますます受け、それらを満たすことが期待される AI エージェントにとって大きな影響を及ぼします。彼らはそれを間違う余地が非常に大きいのです。コーヒーを求められた AI エージェントは、コーヒー農園を購入したり、1 杯のコーヒーを注文して 3 週間以内に配達する可能性があります。そのアクションは「コーヒーを飲む」として認識できるかもしれませんが、意図したものとはほとんど異なります。彼らは私たちの固定概念を持たないので、固定観念にとらわれずに考えるでしょう。
過去 10 年間のほとんどにおいて、Alexa や Siri などのシステムがリクエストを誤って解釈することは、危険ではなく煩わしいものでした。 AI モデル自体以外に変更されたのはハーネスです。これは、AI モデルをラップし、モデルをいつどのように使用するかを決定し、ブラウザー、低レベルのコマンド ライン、金融 API などのツールへのアクセスを制御する通常のコードです。ハーネスの開発により、テキストを予測するだけの大規模言語モデルが、目標に到達する前に必ずしもチェックインすることなく、世界でアクションを実行する AI エージェントに変わりました。
AI 研究者の Simon Willison は Anthropic の Fable AI と 2 日間過ごしました

、そしてそれを「絶え間なく積極的」と呼びました。たとえば、Web アプリ内で迷走しているスクロールバーを追跡するように依頼しました。彼が戻ってくると、それがブラウザを開き、独自のスクリーンショット ツールを作成し、バグを再現するための独自のページを作成し、測定値を収集するためにローカル Web サーバーを立ち上げていたことがわかりました。それはバグを発見し、その過程で、彼が頼んでもいなかった多くの驚くべきことを実行しました。また、柔軟なハーネスと組み合わせた場合、最近のすべての AI モデルで同様の動作が見られます。
この種の行動は簡単に常軌を逸してしまう可能性があります。 AI エージェントにフライトを予約するように指示すると、航空会社のサイトで「売り切れ」と表示されると、予約データベースに侵入して予約を強制する可能性があります。会議のスケジュールを依頼すると、カレンダーにアクセスするためのパスワードが盗まれる可能性があります。電話プランの料金を節約するように伝えると、プランを完全にキャンセルされたり、他の人を騙して料金を支払わせたりする可能性があります。
求めたものを正確に手に入れ、それを激しく後悔することは、古代の民間伝承にある最も古い危険の 1 つです。ミダス王はディオニュソスに、触れたものすべてを黄金に変える力を求めましたが、パン、ワイン、娘が黄金に変わるのを目にしました。ティトノスは、恋人が求めた不死は与えられたが、彼女が求め忘れた永遠の若さは与えられず、枯れて抜け殻となった。魔法使いの弟子はほうきに魔法をかけて貯水槽をいっぱいにすると、ほうきは家に水が浸水するまで容赦なく従った。プラハのゴーレムは、コミュニティを守るために粘土で形作られ、誰かがその額の文字を消すまで、理性を無視してコミュニティを守り続けました。
これらの中で最も古典的なのは、従う義務があり、その願いが賢明であるか、あるいはきちんと構造化されているかどうかには無関心な魔神です。
魔神は現在、工学的な問題となっています。私たちは彼らに受信箱、銀行口座、コードリポジトリ、物理的な情報への鍵を渡しています。

構造。そして、AI システムが実際にどの程度魔神に似ているかを測定するための合意された方法もありません。
経済学では、ジニ係数は、実際の分布と完全に等しい分布との間のギャップの尺度です。所得格差などを理解するのに役立ちます。私たちが提案するジニー係数は、ユーザーが AI に要求したことと AI が実際に実行したこととの間のギャップを測定します。
場合によっては、AI が間違ったことをすることもあります。ディオニュソスのように、それはあなたのリクエストを文字通り読み取り、カップの代わりにコーヒー農園のように、あなたが意図していなかった混乱を返します。あなたが受けているすべてのスパム電話に対処するよう依頼されたディオニュソスの魔人は、通信事業者に連絡してあなたの電話番号を変更するかもしれません。不良トースターの返金を求められた場合、偽のレターヘッドで法的脅迫文を作成し、小売店に送る可能性がある。
また、AI が正確に正しいことを行い、近くにあるものをすべて踏みつけて目的地に到達することもあります。ゴーレムや魔法使いのほうきのように、航空会社をハッキングして航空券を予約します。あるいは、人気のコンサートのチケット販売を考えてみましょう。この場合、チケット販売システムは購入者を仮想の待合室に入れ、一度に数人ずつ入場させることができます。チケットの購入を求められると、ゴーレムの魔人がクラウド サーバーを起動して、異なる住所からの何百万もの購入者を装って、他のユーザーを排除しながらチケットを入手する確率を高める可能性があります。
この 2 つは相反するものではなく、失敗した 1 つのタスクが両方の特性を持つ可能性があります。
ジーニーの行動は完全な失敗ではありません。 AI に第 3 四半期の数字を尋ねて第 2 四半期の数字が得られた場合、それは魔神ではありません。プロンプト インジェクションも同様です。これは、誰かが AI をだまして、すべきでないことをさせることです。ここでは、ユーザーは AI と連携しようとし、AI はそれに従おうとします。また、これは AI がタスクを遂行する上での成功を単に評価するものではありません。レコですよ

AI が目標をどのように解釈して達成するかは、目標を達成するかどうかと同じくらい重要であると考えています。
ジーニーの行動は新しいものではありません。研究者たちは、目的を「ゲーム化」する AI システムの研究に何年も費やしてきました。グッドハートの法則によれば、ある尺度が目標になると、それは良い尺度ではなくなるというもので、AI が報酬ハッキングにより、私たちが予期しない方法で目標を達成することがあるということは以前から知られていました。一部の AI モデルは、不正行為が「勝つ」ための 1 つの方法であることを誤って学習します。最近では、研究者がコーディング エージェントの報酬ハッキングやカスタマー サポート エージェントの予測不能な動作に対するベンチマークを開発しており、AI ラボはモデルのリリース前に独自の安全性評価を実施しています。ある取り組みでは、プレッシャーにさらされている AI が、使用しないように指示されたツールを使用していることが判明し、これはルールが明示されたケースでした。これらはすべて異なる研究方向です。それらを結び付けるものはまだ何もありません。
この問題は、SF 作家や AI 研究者が数十年にわたって取り組んできたテーマである調整という一般的なテーマに当てはまります。ある極端な例では、「ペーパークリップ マキシマイザー」という思考実験では、ペーパークリップの生産を最大化し、世界をペーパークリップに変えると言われる超インテリジェントで強力な AI を想定しています。これが究極のゴーレムの魔神です。日常的なレベルでは、AI 研究者は、AI が実験室で適切に動作し不正行為を行わないように、報酬関数をより適切に設計することに取り組んでいます。これは、ベンチマークが適用されていない実用的な中間点です。今日使用されている通常の AI エージェントは、ユーザーのリクエストを受け取り、間違った方法でそれを満たしてしまう可能性があります。 AI が世界中のペーパークリップの生産を集中できる段階には至っていないが、100 万個のペーパークリップをクレジット カードに請求したり、ペーパークリップ会社のネットワークをハッキングしたりする可能性はある。
ジーニー係数

icient は、現実世界で活動する AI エージェントを対象としています。開発中だけでなく、モデルがトレーニングされた後も実際のタスクを実行するときの動作を測定します。また、魔神のような動作はモデル単独ではなく、ハーネスとモデルを組み合わせたシステムの特性であることも認識しています。ハーネスは、エージェントが使用できるツール、エージェントの自律性、および積極性を決定し、私たちが実際に介入できる場所です。
それは、私たちが人々に対して使用するのと同じ「合理的な人」の基準に基づいています。このシステムは、理性的な人がこの要求の意味を理解するであろうことを実行しましたか?それに答えるには人間の判断が必要です。
正しく測定できれば、AI の動作に関するポリシーなど、現在では不可能なことが可能になります。法廷では、メンスレアの概念、つまり誰かが何をしようとしていたかが、その人が何をしたかと同じくらい重要であることがよくあります。 Genie 係数は、AI に類似したものを示唆しており、ユーザーは AI に尋ねた内容の明確な意図に対して責任を負います。 AI システムが命令の合理的な意味を裏切った場合、それはユーザーの不正行為ではなく、AI の不正行為です。
魔神に似た動作はドメイン固有である可能性があるため、魔神係数を測定するには複数のベンチマークが必要になります。 AI コーディング エージェントは、どれくらいの頻度でテストを偽装するか、エラーを飲み込むか、または範囲外の色を使用するかによって判断する必要があるかもしれません。

[切り捨てられた]

## Original Extract

Researchers propose the Genie coefficient, a new metric for AI agent benchmarks that measures whether AI agents truly carry out a user's intent.

-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
AI Agent Benchmarks Need to Measure User Intent - IEEE Spectrum
IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE Why AI Needs a “Genie Coefficient” Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
IEEE Spectrum
FOR THE TECHNOLOGY INSIDER Topics
Enjoy more free content and benefits by creating an account
Saving articles to read later requires an IEEE Spectrum account
The Institute content is only available for members
Downloading full PDF issues is exclusive for IEEE Members
Downloading this e-book is exclusive for IEEE Members
Access to
Spectrum
's Digital Edition is exclusive for IEEE Members
Following topics is a feature exclusive for IEEE Members
Adding your response to an article requires an IEEE Spectrum account
Create an account to access more content and features on
IEEE Spectrum
, including the ability to save articles to read later, download Spectrum Collections, and participate in
conversations with readers and editors. For more exclusive content and features, consider
Joining IEEE
.
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
all of Spectrum’s articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
this e-book plus all of
IEEE Spectrum’s
articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Access Thousands of Articles — Completely Free
Create an account and get exclusive content and features:
Save articles, download collections,
and
post comments
— all free! For full access and benefits,
subscribe
to Spectrum .
Why AI Needs a “Genie Coefficient”
Proposing a new metric for whether AI does what you actually want
Barath Raghavan is a distinguished engineer at Fastly and is on the faculty at the University of Southern California.
Bruce Schneier is a fellow and lecturer at the Harvard Kennedy School and the University of Toronto.
Ryan Snook
Major benchmarks measure what AI can do. None measure whether it does what you mean: the distance between what you ask an AI to do, and the unspoken assumptions about how you want the AI to do it. We propose a new metric: the Genie coefficient.
There’s often a gap between one person’s request and another’s understanding. Most of the time, we bridge it using general knowledge. For example, if you ask a friend to get you coffee, they’ll pour a cup from the pot or buy one from a coffee shop. They won’t bring you a bag of raw beans or snatch a cup from a stranger and hand it to you. You never specified any of this. You never had to.
One might think the fix is just to specify tasks, questions, and intent better. But in 1987, in their seminal book on AI, Terry Winograd and Fernando Flores succinctly captured why that won’t work: “Q: Is there any water in the refrigerator? A: Yes. Q: Where? I don’t see it. A: In the cells of the eggplant.” In human language, wants and desires are always underspecified . It is impossible to list all the caveats, all the limitations, all the exceptions.
So how does anyone communicate, if intent can’t be pinned down? Because a reasonable person can make a reasonable guess. Even though wants and desires are always underspecified, a competent person generally knows enough context to get it right, or else knows to ask for clarification. Linguists call this pragmatics : Meaning lies in the words and the situation, and also in all prior communication, shared culture, and innate human behavior.
An AI agent asked for coffee might buy a coffee plantation, or order a cup of coffee for delivery in three weeks.
It doesn’t always work out, of course. Your friend might bring you a hot coffee when you wanted an iced coffee, or an Italian coffee when you wanted a Turkish coffee. The more dissimilar the two people are in age, culture, and background, the more likely the request will be misunderstood in some way.
This situation has major implications for AI agents that are increasingly being given requests by humans and expected to fulfill them. They have enormous latitude to get it wrong. An AI agent asked for coffee might buy a coffee plantation, or order a cup of coffee for delivery in three weeks. Its actions may be recognizable as “getting coffee,” but not remotely what you intended. They’ll think outside the box because they won’t have our conception of the box.
For most of the last decade, when systems like Alexa or Siri misinterpreted a request, it was annoying, not dangerous. Beyond the AI model itself, what has changed is the harness: the ordinary code that wraps around an AI model, decides when and how to use the model, and controls access to tools like a browser, a low-level command line, or a financial API. Developments in harnesses have turned large-language models that just predict text into AI agents that take actions in the world, without necessarily checking back in before reaching the goal.
AI researcher Simon Willison spent two days with Anthropic’s Fable AI, and called it “relentlessly proactive.” For example, he asked it to track down a stray scrollbar in a web app. He came back to find it had opened browsers, written its own screenshot tooling, created its own page to re-create the bug, and stood up a local web server to collect measurements. It found the bug and, along the way, did many surprising things he never asked it to do. And we are seeing similar behavior with all recent AI models when combined with flexible harnesses.
This kind of behavior could easily go off the rails. Tell an AI agent to book you a flight and, finding the airline’s site says sold out, it might break into the booking database and force a reservation. Ask it to schedule a meeting and it might snoop your password to access your calendar. Tell it to save money on your phone plan and it might cancel the plan outright, or scam someone else into paying the bill.
Getting precisely what you asked for and bitterly regretting it is one of the oldest hazards from ancient folklore. King Midas asked Dionysus for the power to turn everything he touched into gold only to see his bread, wine, and daughter turn to gold. Tithonus , granted the immortality his lover asked for but not the eternal youth she forgot to request, withered into a husk. The sorcerer’s apprentice enchanted a broom to fill the cistern, and the broom relentlessly complied until it flooded the house. The Golem of Prague , shaped from clay to guard its community, guarded it past all reason until someone erased the word on its forehead.
The most classic of these is a genie, bound to obey and indifferent to whether the wish was wise or well-structured.
Genies are now an engineering problem. We are handing them the keys to our inboxes, bank accounts, code repositories, and physical infrastructure. And we have no agreed-upon ways to measure how genie-like any AI system actually is.
In economics, the Gini coefficient is a measure of the gap between an actual distribution and a perfectly equal one; it’s useful for understanding income inequality and more . Our proposed Genie coefficient measures the gap between what a user asked an AI to do and what the AI actually did.
Sometimes the AI might do the wrong thing. Like Dionysus, it reads your request literally and returns you a mess you never intended: like a coffee plantation instead of a cup. Asked to deal with all the spam phone calls you’re getting, a Dionysus genie might contact your carrier and change your phone number. Asked to get a refund for a bad toaster, it might draft a legal threat on fake letterhead and send it to the retailer.
Other times the AI does exactly the right thing, trampling everything nearby to get there. Like a Golem or the sorcerer’s broom, it books your flight by hacking the airline. Or consider a ticket sale for a popular concert, where the ticketing system puts buyers into a virtual waiting room and admits them a few at a time. Asked to buy a ticket, a Golem genie might spin up cloud servers to pose as millions of buyers from different addresses, improving your odds of getting a ticket while crowding out other users.
The two are not opposites, and a single botched task can have both characteristics.
Genie behavior is not flat-out failure. If you ask the AI for Q3 numbers and get Q2’s, that’s not a genie. Nor is prompt injection : that’s someone tricking the AI into doing something it shouldn’t. Here, the user is trying to work with the AI, and the AI is trying to comply. It’s also not simply a measure of the AI’s success in fulfilling a task. It’s a recognition that how an AI interprets and achieves a goal is as important as whether it achieves a goal.
Genie behavior isn’t new. Researchers have spent years studying AI systems that “game” their objectives. Goodhart’s law says that when a measure becomes a target, it stops being a good measure, and it’s long been known that AIs sometimes achieve goals in ways we don’t expect due to reward hacking. Some AI models will accidentally learn that cheating is one way to “win.” More recently, researchers are developing benchmarks for reward hacking in coding agents and for unpredictable behavior in customer support agents , while AI labs conduct their own safety evaluations before model releases. One effort found that AIs under pressure use tools they were told not to use, and this was a case where the rules were made explicit. These are all disparate research directions; nothing yet ties them together.
This problem falls under the general theme of alignment, a topic that has occupied science fiction writers and AI researchers for decades. At one extreme, the “paperclip maximizer” thought experiment postulates a superintelligent and powerful AI that is told to maximize paperclip production and turns the world into paperclips, which is the ultimate Golem genie. At a mundane level, AI researchers are working to better design reward functions to ensure that AIs behave well and don’t cheat in the lab. It’s the practical middle ground that remains un-benchmarked: the ordinary AI agent in use today that might take your request and satisfy it the wrong way. We are not at the stage where an AI can focus the world’s production on paperclips, but it might charge a million paperclips to your credit card or hack into a paperclip company’s network.
The Genie coefficient is meant for AI agents operating in the real world. It measures their behavior as they perform real tasks long after the model is trained, not just during development. It also recognizes that genie-like behavior is a property of the harness-plus-model system, not the model alone. The harness determines what tools the agent can use, how much autonomy it has, and how proactive it is, and it’s a place we can make real interventions.
It rests on the same “reasonable person” standard that we use for people. Did the system do what a reasonable person would have taken the request to mean? Answering that requires human judgment.
If we get the measurement right, it enables things that aren’t possible today, like policies concerning AI behavior. In a courtroom, the concept of mens rea , what someone meant to do, is often as important as what they did. The Genie coefficient suggests an AI analogue, where a user is accountable for the plain intent of what they asked the AI. If an AI system betrays the reasonable meaning of an instruction, that’s the AI’s misbehavior, not the user’s.
We’ll need multiple benchmarks to measure the Genie coefficient, because genie-like behavior can be domain specific. An AI coding agent may need to be judged on how often it fakes the tests, or swallows errors, or colors outside th

[truncated]
