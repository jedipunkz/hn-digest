---
source: "https://arstechnica.com/apple/2026/06/apple-says-its-ai-is-still-private-even-when-its-running-on-googles-servers/"
hn_url: "https://news.ycombinator.com/item?id=48461720"
title: "Apple says its AI is still private, even when it's running on Google's servers"
article_title: "Apple says its AI is still private, even when it's running on Google's servers - Ars Technica"
author: "Bender"
captured_at: "2026-06-09T16:05:20Z"
capture_tool: "hn-digest"
hn_id: 48461720
score: 3
comments: 3
posted_at: "2026-06-09T14:35:49Z"
tags:
  - hacker-news
  - translated
---

# Apple says its AI is still private, even when it's running on Google's servers

- HN: [48461720](https://news.ycombinator.com/item?id=48461720)
- Source: [arstechnica.com](https://arstechnica.com/apple/2026/06/apple-says-its-ai-is-still-private-even-when-its-running-on-googles-servers/)
- Score: 3
- Comments: 3
- Posted: 2026-06-09T14:35:49Z

## Translation

タイトル: Apple、自社の AI は Google のサーバー上で実行されている場合でも非公開であると発表
記事タイトル: Apple、自社の AI は Google のサーバー上で実行されている場合でも非公開であると発表 - Ars Technica
説明: 一部のモデルは Google のクラウドで実行されますが、Google にいかなるアクセスも許可しません。

記事本文:
コンテンツにスキップ
アルス テクニカ ホーム
セクション
フォーラム
購読する
検索
AI
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ストーリーにピンを付ける
テーマ
ハイパーライト
検索
サインイン
サインインダイアログ...
サインイン
ほぼ曇り
Apple、自社の AI は Google のサーバー上で実行されている場合でも非公開であると発表
一部のモデルは Google のクラウドで実行されますが、Google にいかなるアクセスも許可しません。
61
Apple ソフトウェアエンジニアリング上級副社長、クレイグ・フェデリギ氏。
クレジット:
アンドリュー・カニンガム
Apple ソフトウェアエンジニアリング上級副社長、クレイグ・フェデリギ氏。
クレジット:
アンドリュー・カニンガム
テキスト
設定
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ナビゲーション用に最小化する
カリフォルニア州クパチーノ -- Apple は今年初め、長らく延期されていた Siri のアップグレード（今週「Siri AI」として発表）が Google の Gemini 言語モデルを使用すると発表した。昨日のWorldwide Developers Conferenceで同社が確認したのは、GoogleサーバーにインストールされたNvidiaハードウェアでも動作するということだった。しかし、同社は、すべての AI モデルがユーザーのデバイス上でローカルに実行されていたか、Apple が管理するサーバー ハードウェア上で実行されていた以前と同じプライバシーの約束を今でも行っています。
Appleは何年もの間、自社プラットフォームを利用する主な利点としてユーザーのプライバシーを宣伝してきた。同社のクラウド サービスでは、Apple 従業員を含む他の人がアクセスできないようにすることを目的とした暗号化が使用されています。そして同社は、画像のスキャンなどにオンデバイス処理を使用し、最初からできるだけ多くのデータがデバイスから流出しないようにすることを長い間宣伝してきました。
しかし、Apple Intelligence のおかげで、Apple は自社のハードウェアの限界に直面しました。ローカルで実行できる言語と推論モデルの種類

iPhone or Mac are relatively small, limiting their capabilities and accuracy. Apple’s Private Cloud Compute system was a partial solution but relied on Apple’s own server hardware; Siri AI をサポートするために必要な容量を確保するには、Apple はこれまで避けてきた大規模なデータセンターの増強に取り組む必要があったでしょう。
Appleのクレイグ・フェデリギ氏と他のApple幹部は、WWDCの基調講演後に小規模なステージに上がり、必要なコンピューティング能力を確保しながらユーザーのプライバシーをどのように保護する計画なのか、そしてGoogleとの提携が何を意味するのかを報道機関や他のメディアに説明した。
プライベート クラウド コンピューティングを路上で活用
「これは私たちが使用している Google アシスタントの量です。まったくありません」とフェデリギ氏は、数時間前に CEO のティム・クック氏を紹介した巨大な屋外講堂よりもはるかに親しみやすい劇場で、何もないスライドの前に立って言いました。
フェデリギは、「従来のチャットボット アーキテクチャ」、つまりデバイス上で実行されるクライアント アプリが、サードパーティのサーバー上で実行されるクラウドベースのモデルにアクセスすることについて概要を説明したところです。 Those models can then reach out to Google Search or something similar “to [ground themselves] in world knowledge.”
Apple’s system still depends on an on-device model for simpler queries. In this year’s OS releases , most Apple Intelligence devices get AFM 3 Core, a new Gemini-based model co-developed by Google and Apple.少なくとも 12 GB の RAM と比較的新しいチップを搭載した新しいデバイス (Mac の場合は M3 以降、iPad の場合は M4 以降、iPhone の場合は A19 Pro のみ) では、代わりに AFM 3 Core Advanced が使用され、追加のハードウェアとデバイスのストレージが機能します (ディクテーションを改善し、Siri のより表現力豊かな音声を強化するために使用されます)。
「より高度な」質問については

、デバイスは、やはり Apple と Google が共同開発したクラウドベースのモデル、つまり AFM 3 Cloud と呼ばれる汎用モデル、ADM 3 Cloud と呼ばれる画像生成モデル、および「エージェント ツールの使用と複雑な推論」用の AFM 3 Cloud Pro と呼ばれる高度なモデルに接続します。 Apple によると、最初の 2 つのモデルは依然として Apple のサーバー上の Apple のシリコン上で動作します。 Cloud Pro モデルは、Google が所有する Nvidia ハードウェアで実行されるモデルです。
これを実現するために、同じプライバシーの約束を守りながら、Apple は Private Cloud Compute の新しいバージョンを導入しました。これは、サードパーティのハードウェアで実行するように設計されています。 Apple は、Nvidia の Confidential Computing 、Intel の Trust Domain Extensions 、Google の Titan セキュリティ チップを使用して、Apple が自社のサーバーに提供しているものと同様の保護層を提供しています。追加の保護を提供するために、Apple は「PCC フリートの一部であるすべての Google Cloud ハードウェアの暗号的に検証可能な追加専用台帳」を保持しており、Apple のデバイスは、Apple によって署名されたこれらのサーバー上のソフトウェアのみを信頼します。
Google Cloud サーバーはまだ、Apple 独自の Private Cloud Compute サーバーと同じ保護機能をすべてサポートしているわけではありませんが、Apple は「夏のプレビュー期間を通じて完全な保護セットに向けて徐々に強化していく予定です」と述べています。
どのモデルを使用するか、どのアプリがどのデータにアクセスできるかなどの重要な決定は、Apple が「システム オーケストレーター」と呼ぶオンデバイス機能によって処理されます。その役割の 1 つは、ユーザーのクエリに答えるために必要なデータのみが最初にデバイス外に送信されるようにすることです (デバイスは、たとえば、メッセージ アプリで送信されたレシピに関する回答を、送信者、送信日時、送信理由などの情報を取得せずに生成できます)。

「PCC に送信されるものは絶対に最小限に抑えますが、PCC で重要なことは、アーキテクチャ上、その時点での効率化の尺度であるということです」とフェデリギ氏は述べています。 「PCC 自体は、根本からの設計により、質問に回答した直後にそのデータの記録を蒸発させることになっているためです…これは保存されません。すべては完全に一時的な形式です。」
Siri AI とその他の新しい Apple Intelligence 機能は、iOS 27、iPadOS 27、macOS 27 Golden Gate、および Apple のその他のオペレーティング システムの今秋リリースの一部としてリリースされる予定です。現在、開発者は最初のベータ版を利用できますが、ほとんどの人は、より安定したパブリック ベータ版がリリースされる 7 月まで待って試したほうがよいでしょう。
61件のコメント
コメント
フォーラムビュー
コメントを読み込んでいます...
前の話
次の話
よく読まれている
1.
ここ数週間で 2 回目、Microsoft パッケージに認証情報スティーラーが混入
2.
Falcon 9 ブースターが 5 周年を迎え、驚くべき再利用記録を樹立しました
3.
ロシアの衛星が大陸規模でGPSを妨害できる可能性があることをテストが示唆
4.
「チャットは死んだ」: OpenAI が ChatGPT のオーバーホールを準備
5.
USB 接続されたスピーカーが触れずに PC に感染する仕組み
カスタマイズ
Ars Technica は信号を分離してきました。
25年以上続く騒音。弊社独自の組み合わせにより、
技術的な知識と技術芸術への幅広い関心
Ars は、情報の海の中で信頼できる情報源です。後
すべてを知る必要はありません。重要なことだけを知っておく必要があります。

## Original Extract

Some models run in Google's cloud, but without giving Google any kind of access.

Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
mostly cloudy
Apple says its AI is still private, even when it’s running on Google’s servers
Some models run in Google’s cloud, but without giving Google any kind of access.
61
Apple software engineering SVP Craig Federighi.
Credit:
Andrew Cunningham
Apple software engineering SVP Craig Federighi.
Credit:
Andrew Cunningham
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
CUPERTINO, California—Apple announced earlier this year that its long-delayed Siri upgrade, announced this week as “Siri AI,” would use Google’s Gemini language models . What the company confirmed at its Worldwide Developers Conference yesterday was that it also ran on Nvidia hardware installed in Google servers. But the company is still making the same privacy promises it did before, when all of its AI models were either running locally on your devices or on Apple-controlled server hardware.
For years, Apple has touted user privacy as a key benefit of using its platforms. Its cloud services use encryption that’s intended to keep other people—including Apple employees—from being able to gain access to it. And the company has long advertised its use of on-device processing for things like scanning images, keeping as much data as possible from leaving your device in the first place.
But with Apple Intelligence, Apple has run up against the limits of its own hardware. The kinds of language and reasoning models that can run locally on an iPhone or Mac are relatively small, limiting their capabilities and accuracy. Apple’s Private Cloud Compute system was a partial solution but relied on Apple’s own server hardware; to get the kind of capacity it would need to support Siri AI, Apple would have had to commit to a huge data center buildout that it has so far avoided .
Apple’s Craig Federighi and other Apple executives got on a smaller stage after the WWDC keynote to explain to the press and other media how it planned to preserve user privacy while still getting the kind of compute capacity it needed and what its partnership with Google meant.
Taking Private Cloud Compute on the road
“This is the amount of the Google Assistant we use, which is none,” says Federighi, standing in front of a blank slide in a much more intimate theater than the giant outdoor auditorium where he had introduced CEO Tim Cook a couple of hours before.
Federighi has just outlined a “traditional chatbot architecture”—a client app running on your device that reaches out to cloud-based models running on third-party servers. Those models can then reach out to Google Search or something similar “to [ground themselves] in world knowledge.”
Apple’s system still depends on an on-device model for simpler queries. In this year’s OS releases , most Apple Intelligence devices get AFM 3 Core, a new Gemini-based model co-developed by Google and Apple. Newer devices with at least 12GB of RAM and a relatively recent chip (M3 and newer for Macs, M4 and newer for iPads, just the A19 Pro for iPhones) use AFM 3 Core Advanced instead, which leverages the extra hardware as well as your device’s storage to function (it’s used to improve dictation and power Siri’s more expressive voice).
For “more sophisticated” questions, your device will contact cloud-based models, again co-developed by Apple and Google: a general-use model called AFM 3 Cloud, an image-generation model called ADM 3 Cloud, and an advanced model called AFM 3 Cloud Pro for “agentic tool use and complex reasoning.” The first two models, Apple says, still run on Apple’s silicon on Apple’s servers. The Cloud Pro model is the one running on Google-owned Nvidia hardware.
To do this while still making the same privacy promises, Apple has introduced a new iteration of Private Cloud Compute, this one designed to run on third-party hardware. Apple is using Nvidia’s Confidential Computing , Intel’s Trust Domain Extensions , and Google’s Titan security chip to provide layers of protection similar to what Apple provides for its own servers. To provide additional protection, Apple keeps “a cryptographically verifiable, append-only ledger of all Google Cloud hardware that is part of the PCC fleet,” and Apple’s devices will only trust software on these servers that is signed by Apple.
The Google Cloud servers don’t yet support all the same protections as Apple’s own Private Cloud Compute servers, but Apple says it “will be gradually ramping towards the complete set of protections throughout the summer preview period.”
Important decisions, like which model to use and what apps have access to what data, are handled by an on-device feature Apple calls the “System Orchestrator.” Among its duties is making sure that only the data needed to answer a user query is sent off-device in the first place (your device could generate an answer about a recipe you were sent in the Messages app, for example, without getting information about the person who sent it to you, when they sent it, or why they were sending it).
“While we absolutely minimize what is sent up to PCC, the critical thing about PCC is, architecturally, that’s at that point an efficiency measure,” said Federighi. “Because PCC itself, by design from the ground up, is going to vaporize any record of that data the moment after it answers your question… This is not stored. It’s all in a form where it’s completely transient.”
Siri AI and the other new Apple Intelligence features will launch as part of iOS 27, iPadOS 27, macOS 27 Golden Gate, and Apple’s other operating system releases this fall. The first beta versions are available to developers now, but most people would be better served by waiting until July to try it, when a more stable public beta version will be released.
61 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
For the 2nd time in weeks, Microsoft packages laced with credential stealer
2.
A Falcon 9 booster turns 5 years old—and just set a remarkable reuse record
3.
Tests suggest Russian satellites can jam GPS on a continental scale
4.
"Chat is dead": OpenAI preps overhaul of ChatGPT
5.
How a USB-connected speaker can infect a PC without ever being touched
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
