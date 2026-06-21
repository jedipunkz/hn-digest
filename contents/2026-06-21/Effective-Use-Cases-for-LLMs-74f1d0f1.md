---
source: "https://aggressivelyparaphrasing.me/2026/06/21/effective-use-cases-for-llms/"
hn_url: "https://news.ycombinator.com/item?id=48615807"
title: "Effective Use-Cases for LLMs"
article_title: "Effective use-cases for LLMs - Aggressively Paraphrasing Me"
author: "tcbrah"
captured_at: "2026-06-21T07:45:39Z"
capture_tool: "hn-digest"
hn_id: 48615807
score: 2
comments: 0
posted_at: "2026-06-21T04:55:27Z"
tags:
  - hacker-news
  - translated
---

# Effective Use-Cases for LLMs

- HN: [48615807](https://news.ycombinator.com/item?id=48615807)
- Source: [aggressivelyparaphrasing.me](https://aggressivelyparaphrasing.me/2026/06/21/effective-use-cases-for-llms/)
- Score: 2
- Comments: 0
- Posted: 2026-06-21T04:55:27Z

## Translation

タイトル: LLM の効果的な使用例
記事のタイトル: LLM の効果的な使用例 - 積極的に言い換える
説明: LLM の欠点についてはよく言われています。彼らは実際には理屈を言っていません。特にループ内で実行する場合、コストが高くなります。彼らは物事を行うのが非常に遅いです。 LLM が得意とするユースケースには狭いカテゴリがあり、その 1 つは「ノイズをふるい分ける」ことです。ノイズがすべてだ
[切り捨てられた]

記事本文:
LLM の効果的な使用例 - 積極的に言い換える
コンテンツにスキップ
私を積極的に言い換える
マークの考えやアイデアの一部
LLM の欠点についてはよく言われています。彼らは実際には理屈を考えていません。特にループ内で実行する場合、コストが高くなります。彼らは物事を行うのが非常に遅いです。
LLM が得意とするユースケースは狭いカテゴリにありますが、その 1 つは「ノイズを選別する」ことです。ノイズは、私たちが本当に望むものを得るために処理しなければならないすべてです。
ここでは、私がソフトウェア エンジニアとして楽しんでいる、聞いたことのない使用例をいくつか紹介します。
顧客との会話から検索する
PM の同僚が、当社の主要顧客とのすべての通話の記録を埋め込み DB にアップロードしました。現在、彼らの製品提案は証拠によって深く裏付けられています。当社の上位顧客の 40% がこの問題点について言及していることを私たちは知っています。 PM はまた、私たちの新機能を試してみたいという熱心なプライベート ベータ顧客のリストも特定しました。
これは、顧客の問題が抽象的な場合に役立ちます。多くの場合、これらの問題には明確な解決策がないか、解決策に明確な名前がありません。そのため、機能リクエストの提出が難しくなり、整理や重複排除がさらに難しくなります。 LLM が導入される前は、チームの誰かが十分な在職期間を持っていて、これが何度も発生するのを見て、すべてのリンクと接続を見つける方法を覚えていることが最善の策でした。さて、RAGです。
エンドポイントアラート -> ログ分析へ
大規模なシステムは、ほとんどの時間、障害モードで動作することになります。
— ジョン・ガル、ロリン・ホッホシュタイン経由、Netflix
オンコール時の私の責任の 1 つは、チームが所有する API エンドポイントでの障害を優先順位付けすることです。これらの障害は、「HTTP 4XX/5XX の高レート」として報告されます。場合によっては、ポッドの DB 接続の中断などのノイズが発生することがあります。他の

場合によっては、顧客が何かを削除できなくなったなどのバグを示していることもあります。
最初のステップは、特定のエンドポイントに特定の HTTP エラーをマークする正規ログ行を時間でフィルターして検索することです。
アラートをトリガーしたリクエストを見つけたら、リクエスト ID で検索し、リクエストを最初から最後まで確認します。ログとソース コードに基づいて、通常は何が問題だったかを推測できます。
場合によっては、スタック トレースが Typescript ではなく JavaScript でコンパイルされるため、行番号が一致しないことがあります。次の関数呼び出しの名前に基づいて推測する必要があります。
代表的なリクエストを検討していることを再確認します。さらに 2 つまたは 3 つのリクエスト ID をすぐに調べて、それらがすべて同じ根本原因であることを確認します。
DB 接続タイムアウトなどのより困難な問題については、タイムスタンプ、ホスト マシン、顧客 ID に関する正規ログ行にクラスタリングがあるかどうかを確認します。もしかしたら、それは特に私のルートではなく、インフラの問題かもしれません。
全体として、ふるいにかけるべきものがたくさんあります。非常に多くの判断が必要ですが、解決策を考えることはおろか、問題を見つけることさえできていません。
ただし、エージェント ハーネスはこれに対してほぼ完璧です。アラートとタイムスタンプがあれば、ログ、ソース コード、クラスタリングなど、正しい方向を教えてください。これにより、トリアージにかかる時間が、問題ごとに 15 分以上から 1 ～ 2 分に短縮されました。 SotA ($$$$) モデルも必要ありません。お金を節約して、より高速なモデルを使用してください。
実際のヒューマン スキルを共有することを目的として、このワークフローをチームメイトのスキルとして公開しました。出力には、試行したすべてのクエリの名前が記載されており、有益か非有益かに分類され、さらに深く掘り下げるためのリンクが付いています。チームメイトにトリアージについての考え方を知ってもらいたいので、魔法のようなことは望んでいません。ランプにもしたい

独立した発見。
特にこれを要約とは言いませんでした。理由は次のとおりです。
ChatGPT は要約しません。 ChatGPT にこのテキストを要約するように依頼したところ、代わりにテキストが短縮されました。
— https://ea.rna.nl/2024/05/27/when-chatgpt-summarises-it-actually-does-nothing-of-the-kind/
しかし、それにもかかわらず、私は依然としてテキストを短縮することに信じられないほどの価値があると感じています。 1 時間を超えるポッドキャストやビデオを勧められることがあります。最初の5分で夢中になってしまうこともあります。しかし、技術的なコンテンツの場合、私の興味はビデオの奥深くに埋もれていることが多く、録画された講演の場合は 30 分程度かかる場合があります。何かが自分にとって興味深いかどうかを判断するのにそれほど多くの時間を費やしたくないのですが、LLM はそれを大いに助けてくれます。
私の経験では、短縮バージョンに十分な興味深いコンテンツがあれば、短縮されていないバージョンにも十分なコンテンツがあります。あるビデオでは、米国の東海岸と西海岸の番組について何気なく言及していました。短縮していなかったら、興味がなくなって19秒前に見るのをやめていたでしょう。
要約は私にとって非常に便利ですが、ビデオやポッドキャストで要約を機能させるにはどうすればよいでしょうか?私は、リンクを指定すると以下をチェックする小さな自動化を自分で作成しました。
字幕がある場合はそれをダウンロードしてください
ビデオの場合は、文字起こし用に音声をダウンロードします
ゆっくりとした動画や音声をテキストに変換すれば、まとめられる！
私がこれを言っているのは、おそらくこれが私の ADHD の対処スキルである可能性があるということを警告しています。特にオーディオに関しては、注意力を維持するのが難しいです。私は文字通り、聴覚の集中力、一貫性、そしてスタミナの点で下位 1% に入っているというテスト結果を持っています。これらは 3 つの個別のスキルですが、私は統計的にそれらすべてがひどいです。したがって、おそらく最も重要なのは、音声をテキストに変換する能力です。なぜなら、私は音声よりもテキストの方がはるかにうまく処理できるからです (

統計的にはまだ平均的ですが）。
これは私の夢ですが、インスタグラムに投稿したすべての動画をインデックスに登録したいと思っています。画面上の字幕を OCR し、ビデオから音声を転写し、サムネイルからオブジェクトを検出したいと考えています。 3年前の気に入ったビデオを見つけるのがとても難しいので、これが欲しいです。 Instagram や TikTok など、インターネットのどこかに浮かんでいるだけです。どうやって見つければいいのか分かりません。しかし、埋め込みならできると思います。
確かに、Google は「世界中の情報を整理」したいと考えています。 Apple Photos アプリで「馬」を検索すると、見たことのあるすべての馬が表示されるのでありがたいです。しかし、私は自分のミームのためにそれを望んでいるので、Instagramがこの点でどれほど遅れているかにショックを受けています。
私は LLM についてまだ迷っています。オープンウェイトモデルを楽しんで使っています。推論を手頃な価格で実現する取り組みを見るのが楽しみです。そして私は、それらが私たちの経済、社会構造、個人の精神に及ぼす影響を恐れています。しかし、私はそこからいくつかの楽しいグッズを手に入れました。
はい、地球は破壊されました。しかし、素晴らしい瞬間として、私たちは株主のために多くの価値を生み出しました。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
グランドティトンとイエローストーン 2026
空の Bernzomatic プロパン タンクの重量は 14.5 オンスまたは 412 g (風袋)
Slack チャンネルを「何を」ではなく「どのくらいの頻度で」整理するか

## Original Extract

There's a lot of talk about the shortcomings of LLMs. They don't actually reason. They're expensive, especially when running in a loop. They're quite slow at doing things. There's a narrow category of use cases that LLMs excel at, one of which is "sifting through the noise". The noise is everything
[truncated]

Effective use-cases for LLMs - Aggressively Paraphrasing Me
Skip to content
Aggressively Paraphrasing Me
Some of Mark's thoughts and ideas
There’s a lot of talk about the shortcomings of LLMs. They don’t actually reason. They’re expensive, especially when running in a loop. They’re quite slow at doing things.
There’s a narrow category of use cases that LLMs excel at, one of which is “sifting through the noise”. The noise is everything we have to process to get to what we really want.
Here are some use cases I haven’t heard about that I’ve enjoyed as a software engineer.
Searching through Customer Conversations
A PM colleague uploaded the transcript of every call with our top customers into an Embedding DB. Now their product proposals are deeply backed by evidence. We know 40% of our top customers have mentioned this pain point. The PM also identified a list of eager private beta customers to try out our new feature.
This is useful when the customer’s problem is abstract. Often, these issues don’t have clear solutions, or those solutions don’t have clear names. That makes filing Feature Requests hard, and organizing/deduping even harder. Before LLMs, your best bet was that someone on your team had enough tenure to have seen this come up enough times, and that they remembered how to find all the links and connections. Now, it’s RAG.
Going from endpoint alert -> log analysis
Any large system is going to be operating most of the time in failure mode.
— John Gall, via Lorin Hochstein, Netflix
When I’m on-call, one of my responsibilities is to triage failures on API endpoint out team owns. These failures are reported as “high rate of HTTP 4XX/5XX”. Sometimes, it’s noise, like there’s a DB connection hiccup for the pod. Other times, it’s signaling a bug, like customers can’t delete something anymore.
The first step is searching for the canonical log lines that mark the specific endpoint with the specific HTTP failure, filtered by time.
Once I find the request that triggered the alert, I search by request ID, to see the request from start to end. Based on the logs, and my source code, I can usually guess what went wrong.
Sometimes the stack trace is compiled JavaScript, rather than Typescript, so the line numbers don’t line up. I have to guess based on the name of the next function call.
I double-check that I’m looking at a representative request. I quickly look at two or three more request IDs to make sure they’re all the same root cause.
For more difficult issues like DB connection timeouts, I’ll see if there’s clustering on the canonical log lines around timestamp, host machine, customer ID. Maybe it’s not specifically my route, but an infra issue.
All in all, there’s a lot of stuff to sift through. There’s so much judgment required, and I haven’t even found the problem, let alone thought about a solution yet.
Yet, an agent harness is almost perfect for this. Given some alert and timestamp, point me in the right direction: logs, source code, or clustering. This has cut my triaging time from 15+ minutes to 1-2 minutes per issue. You don’t even need the SotA ($$$$) models. Save your money, use a faster model.
I published this workflow as a skill for my teammates with the intention of sharing the actual human skill involved. The output names all the queries it tried, categorized into informative or non-informative, with links to dig deeper. I don’t want it to be magical, because I want my teammates to know how to think about triaging. I also want it to be a ramp to independent discovery.
I specifically didn’t call this summarizing because:
ChatGPT doesn’t summarise. When I asked ChatGPT to summarise this text, it instead shortened the text .
— https://ea.rna.nl/2024/05/27/when-chatgpt-summarises-it-actually-does-nothing-of-the-kind/
But despite all that, I still find incredible value in shortening texts! I’ll sometimes get recommended a podcast or video that’s over 1 hour long. Sometimes, I’m hooked within the first 5 minutes. But for technical content, my interest is often buried deep in the video, maybe 30 minutes in for recorded talks. I don’t want to spend that much time figuring out if something is interesting to me, and LLMs greatly help with that.
In my experience, if there’s enough interesting content in the shortened version, there’s plenty in the unshortened version. One video casually mentioned east-coast vs west-coast programming in the US . Without shortening, I would have stopped watching 19 seconds earlier out of disinterest.
Okay, summarizing is really useful to me, but how do I get it to work on videos and podcasts? I made myself a little automation that, given a link, will check:
If there are subtitles, download that
If it’s a video, download the audio for transcribing
Once I transform slow video or audio formats into text, I can summarize!
I say all this with the caveat that maybe this is a coping skill for my ADHD. Attention is hard for me to maintain, specifically for audio. I literally have a test result saying I’m in the bottom 1% for auditory focus, consistency, AND stamina. These are three separate skills, and I’m statistically awful at all of them. So maybe what matters the most is the ability to transform audio into text, since I’m able to process text much better than audio (though statistically still average).
This is a dream of mine, but I want to index every video I’ve hearted on Instagram. I want to OCR the subtitles on the screen, transcribe the audio from the video, and object-detect from the thumbnails. I want this because it’s so damn hard to find that one video I liked from 3 years ago. It’s just floating around somewhere in the internet, on Instagram or TikTok. I have no idea how to find it. But I think embeddings can .
Sure, Google wants to “organize the world’s information”. I appreciate that I can search “horse” on the Apple Photos app and get all the horses I’ve seen. But I want it for my memes, and I’m shocked at how behind Instagram is on this.
I’m still very on the fence about LLMs. I enjoy using my open-weight models. I’m excited to see work on making inference affordable. And I’m terrified of their impact on our economy, social fabric, and individual psyche. But I got some fun goodies out of it .
Yes, the planet got destroyed. But for a beautiful moment in time we created a lot of value for shareholders.
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
Grand Tetons and Yellowstone 2026
Empty Bernzomatic Propane Tank Weighs 14.5oz or 412g (Tare)
Organize your Slack channels by “How Often”, not “What”
