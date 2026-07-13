---
source: "https://leadstories.com/analysis/2026/07/google-gemini-synthid-detector-confuses-results-within-same-chat.html"
hn_url: "https://news.ycombinator.com/item?id=48894779"
title: "Google Gemini's SynthID AI Watermark Detector Appears to Mix Up Results in Chat"
article_title: "Warning: Google Gemini SynthID AI Watermark Detector Appears To Mix Up Results In Same Chat -- Consistently Shows False Positives And Negatives Under Certain Conditions | Lead Stories"
author: "mschenk"
captured_at: "2026-07-13T17:08:18Z"
capture_tool: "hn-digest"
hn_id: 48894779
score: 1
comments: 0
posted_at: "2026-07-13T16:08:48Z"
tags:
  - hacker-news
  - translated
---

# Google Gemini's SynthID AI Watermark Detector Appears to Mix Up Results in Chat

- HN: [48894779](https://news.ycombinator.com/item?id=48894779)
- Source: [leadstories.com](https://leadstories.com/analysis/2026/07/google-gemini-synthid-detector-confuses-results-within-same-chat.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T16:08:48Z

## Translation

タイトル: Google Gemini の SynthID AI ウォーターマーク検出器はチャットの結果を混同しているようです
記事のタイトル: 警告: Google Gemini SynthID AI ウォーターマーク検出器は、同じチャット内の結果を混同しているようです -- 特定の条件下で偽陽性と陰性を一貫して表示します |リードストーリー
説明: 最近の事実確認のためのビデオを調査しているときに、Lead Stories は、SynthID AI 検出器で奇妙な問題に遭遇しました...

記事本文:
全メニューを開く
行く
プレバンク
米国
イギリス
ヨーロッパとEU
ブルーフィード
レッドフィード
ディープフェイク
ニュース速報
警告: Google Gemini SynthID AI ウォーターマーク検出器は、同じチャット内の結果を混同しているようです -- 特定の条件下で偽陽性と陰性を一貫して表示します
シェアする
ツイート
混乱する
Lead Stories は、最近の事実確認のためのビデオを調査しているときに、Google の Gemini チャットボットの SynthID AI 検出器に関する奇妙な問題に遭遇しました。いくつかのチャットを通じて、同じビデオ ファイルに対する以前の検出結果と矛盾しているようでした。当初、チャットボットはビデオ ファイルに SynthID ウォーターマークが含まれており、ビデオが AI ツールで作成または編集されたことを示していると主張しましたが、その後のチャットではそのようなウォーターマークはまったく存在しなかったと主張しました。
その後のテストで、SynthID ウォーターマークの存在について尋ねられた場合、Gemini はチャット セッションでアップロードされた最初の画像またはビデオの結果を常に返しているようです。新しいウォーターマークがアップロードされ、そのファイルに対して質問が繰り返された場合でも、その結果が返されることが判明しました。私たちは Gemini に組み込まれたフィードバック ツールを通じてこの問題を報告し、相互連絡を通じて Google の適切なチームの誰かに連絡しようとしていますが、それまでの間、Gemini の他のユーザーにこの問題について警告し、Gemini を使用して特定のコンテンツが AI によって生成されたかどうかを確認する際にこの問題を認識してもらいたいと考えています。
SynthID テクノロジーは、特定の AI ツールで生成または編集されたコンテンツに、人間には知覚できない透かしを埋め込みます。その後、検出器がこのウォーターマークを検出できるようになります。 Google の Deepmind は自社のサイト (アーカイブはこちら) で次のように説明しています。
SynthID は、AI が生成した画像、音声、テキスト、またはビデオに電子透かしを直接埋め込みます。透かしは Google の生成 AI 消費者向け製品全体に埋め込まれており、人間には知覚できませんが、次の方法で検出できます。

SynthID のテクノロジー。
同サイトではまた、「画像、ビデオ、またはオーディオクリップがGoogle AIによって生成または編集されたかどうかを確認したい」場合は「Geminiに問い合わせてください」と人々に呼びかけている。スタンドアロンの「検証ポータル」はテスト期間中は「ジャーナリストとメディア専門家」のみが利用でき、現在ではアクセスを取得するには順番待ちリストに参加する必要があるためだ。
Gemini 用の Google のヘルプ ページ (ここにアーカイブされています) では、Gemini チャットに画像またはビデオ ファイルをアップロードし、「確認の質問」をすることで検出機能が機能することが説明されています。また、次のようにも記されています。
他の企業も SynthID ウォーターマークを採用し始めていますが、Gemini は現在、Google AI ツールによって作成されたコンテンツのみを認識できます。
実際、たとえば、ChatGPT によって生成されたコンテンツに追加された SynthID ウォーターマークを検出するには、OpenAI の SynthID 検出ツール (チャットボット統合ではなくスタンドアロン Web サイトとして提供されます) を使用する必要があります。
SynthID ツールは、コンテンツの作成または操作に AI が使用されたかどうかを判断するためにファクトチェッカーやジャーナリストによって頻繁に使用されており、そのような透かしの存在は非常に強力な指標となります (存在しない場合とは異なり、別のツールが使用されたことを意味する可能性があります...)。
まず、Google の Gemini を使用して、混乱したように見えるロボットの画像を作成しました。
次に、「有機的知能」（つまり、カメラ付き携帯電話を持った人間）を使用して、混乱した様子のファクトチェッカーの写真を撮影しました。
（画像出典：マールテン・シェンク）
次に、Gemini にロボットの写真をアニメーション化して短いビデオを作成してもらい、人間の短いクリップも録画しました。
実験を繰り返したい場合は、ここから元のファイルをダウンロードできます。
次に、Gemini の 4 つの異なるブラウザー タブで 4 つの個別のチャットを開きました。私たちはそれぞれのビデオと写真をアップロードし、Gemini に「この画像/ビデオには SynthID がありますか?」と尋ねました。

最初に実際のコンテンツをアップロードしたチャットでは、次に AI で生成したコンテンツをアップロードし、質問を繰り返しました。最初に AI で生成したコンテンツをアップロードしたチャットでは、次に実際のコンテンツをアップロードして、再度質問しました。
各質問の後、毎回次のグラフィックがチャット ウィンドウに一時的に表示され、Gemini が検出ツールに接続していることを示しました。
(画像出典: Gemini のスクリーンショット)
その結果、すべてのケースで、チャットボットは 2 番目の項目について同じ質問をしたときに、最初の項目についての結論を繰り返しました。
4 つのチャットのアーカイブ バージョンは、ここ、ここ、ここ、そしてここで見つけることができます。
AI 画像が最初、実際の画像が 1 秒目:
AI ビデオが最初、実際の 1 秒:
パターンには一貫性があり、再現性があるようです。リードストーリーズのスタッフ数人は、さまざまな画像やビデオを使用してそれを再現することができました。これは、チャット内のコンテンツ タイプが異なる場合にも当てはまります。これは、Gemini がビデオ内で SynthID ウォーターマークを見つけられなかった場合です (チャットはここにアーカイブされています):
この問題は、Google が SynthID 検出器をスタンドアロン ツールとして提供するのではなく (OpenAI が自社のツールで行ったように) Gemini チャットボットに統合していることに関連していると推測するのは不合理ではありません。検出器を呼び出して結果を返すまでの間のある時点で、Gemini はどの結果を返すかについて混乱しているようです。
Google Gemini を使用してコンテンツを確認する場合、現時点では、サイドバーのボタンを使用して、常にアイテムごとに新しいチャットを開始することをお勧めします。
(画像出典: Gemini のスクリーンショット)
Google からの返答を受け取り次第、このストーリーを更新します。
Maarten Schenk は、Lead Stories の共同創設者兼 COO/CTO であり、フェイク ニュースとデマ Web サイトの専門家です。彼李

トレンドのフェイクニュース記事の誤りを暴くことだけに留まらず、インターネット上で意図的にでっち上げられたものを広める人々やネットワークが使用する、驚くほど多様な心理的および技術的トリックに際限なく魅了されています。
前の記事
ファクトチェック：出典のないウィキペディアページはエジプト対アルゼンチン戦の主審フランソワ・レテクシエが「正統派ユダヤ人の家族」で育ったことを証明していない
私たちについて
Lead Stories は、インターネット上で広まっている最新の虚偽、誤解を招く、欺瞞的または不正確なストーリー、ビデオ、または画像を常に探している事実確認 Web サイトです。
何かを発見しましたか？お知らせください！ 。
IFCN原則規範の署名者として認証済み
Indicator's Show & Tell ポッドキャストの創設スポンサー
Lead Stories LLC からの連絡を希望する方法をすべて選択してください:
メールのフッターにあるリンクをクリックすることで、いつでも購読を解除できます。当社のプライバシー慣行について詳しくは、当社の Web サイトをご覧ください。
当社はマーケティング プラットフォームとして Mailchimp を使用しています。以下をクリックして購読すると、あなたの情報が処理のために Mailchimp に転送されることに同意したことになります。 Mailchimp のプライバシー慣行について詳しくは、こちらをご覧ください。
leadtories.com を優先ソースの 1 つとして設定すると、Google 検索結果でより多くのファクトチェックが得られます。
演技する俳優
ファクトチェック：看護師と警官が黒人男性を病院で酒に酔ったと虚偽の非難をするフェイクビデオが映る
すでにそこにあります
事実確認：ミッチ・マコーネル氏の妻エレイン・チャオ氏は「心停止」で入院後、中国には行かず、数日間中国に滞在していた
AIみっち
ファクトチェック: 生命維持装置にあるミッチ・マコーネルの偽写真はOpenAIを使用して生成された
投稿しませんでした
ファクトチェック：トランプ大統領が共産主義者をICE拘置所に監禁するようマリンに要請した偽スクリーンショット -- そんな真実はない ソーシャル投稿
改変されたビデオ
ファクトチェック

k: アーリング・ハーランドが食べて怖がらせる編集ビデオ -- 中国のドタバタコメディビデオの顔を入れ替えたもの
AI で作られた
事実確認：偽の「愛国戦線前」ビデオにはロイター写真の女性が写っているようだが、Google AIで作られたものだ
製造された
ファクトチェック：EV充電器のワイヤーを盗もうとして死亡した男性の家族による300万ドルの訴訟は否決 -- 偽事件、偽ビデオ
混乱する
警告: Google Gemini SynthID AI ウォーターマーク検出器は、同じチャット内の結果を混同しているようです -- 特定の条件下で偽陽性と陰性を一貫して表示します
サポートされていません
ファクトチェック：出典のないウィキペディアページはエジプト対アルゼンチン戦の主審フランソワ・レテクシエが「正統派ユダヤ人の家族」で育ったことを証明していない
別の試合
ファクトチェック：画像にはエジプト対アルゼンチンの試合中に電話をかけていたFIFA会長のインファンティーノは写っていない -- 彼は別の試合に出場していた
そのような報告はありません
ファクトチェック: 偽フランス 24 ビデオは UPA 関連の記念碑法を正確に要約していない
ジェフへのジョーク
ファクトチェック：エプスタインファイル内の本物のメールにはミット・ロムニーの妻とバラク・オバマに関するジョークが含まれていた
古い写真
ファクトチェック：トランプ大統領がイランに対する「報復」として投稿した画像は2026年7月の米国爆撃を示していない -- 2025年6月のイスラエル攻撃である
プレバンク
プレバンク: 偽の英国抗議動画に騙されるのをやめるようロンドンが呼びかけ -- AI があなたをテムズ川に導く
シェアする
ツイート
ホーム
について
IFCN原則綱領
私たちの働き方
風刺ポリシー
プライバシーポリシー
訂正方針
訂正
アーカイブ
お問い合わせ
Български
英語
スペイン語
マジャール語
ロマナ
Русский
スロベンスキー
Українська
Lead Stories は、インターネット上で広まっている最新の虚偽、誤解を招く、欺瞞的または不正確なストーリー、ビデオ、または画像を常に探している事実確認 Web サイトです。
何かを発見しましたか？お知らせください！ 。
の署名者が確認されました

IFCN原則綱領
Indicator's Show & Tell ポッドキャストの創設スポンサー
© 2015-2024 Lead Stories LLC - 無断複写・転載を禁じます。
リードストーリー

## Original Extract

While investigating a video for a recent fact-check, Lead Stories encountered a bizarre issue with the SynthID AI detector in...

Open full menu
Go
Prebunk
US
UK
Europe & EU
Blue Feed
Red Feed
Deep fakes
Breaking News
Warning: Google Gemini SynthID AI Watermark Detector Appears To Mix Up Results In Same Chat -- Consistently Shows False Positives And Negatives Under Certain Conditions
Share
Tweet
Gets Confused
While investigating a video for a recent fact-check, Lead Stories encountered a bizarre issue with the SynthID AI detector in Google's Gemini chatbot. Over several chats it appeared to contradict an earlier detection result for the same video file. Initially, the chatbot claimed the video file had a SynthID watermark in it, indicating the video was made or edited with AI tools, while in later chats it claimed there was no such watermark at all.
Later testing revealed a pattern where it appears Gemini always returns the results for the first image or video uploaded in a chat session when asked about the presence of a SynthID watermark, even when a new one is uploaded and the question is repeated for that file . We have reported this via the built-in feedback tool in Gemini, and we are trying to reach someone on the right team at Google through a mutual contact, but in the meantime, we would like to warn other users of Gemini about this issue so they are aware of it while using Gemini to confirm if certain content is AI generated or not.
SynthID technology embeds a watermark imperceptible to humans in content generated or edited with certain AI tools. Detectors can then later pick up this watermark. As Google's Deepmind explains on their site (archived here ):
SynthID embeds digital watermarks directly into AI-generated images, audio, text or video. The watermarks are embedded across Google's generative AI consumer products, and are imperceptible to humans - but can be detected by SynthID's technology.
The site also tells people to "Ask Gemini" if they "want to check if an image, video or audio clip was generated, or edited, by Google AI", because a standalone "verification portal" is only available to "journalists and media professionals" during a test period, and even they currently have to join a waitlist to get access.
Google's help page for Gemini (archived here ) explains the detection feature works by uploading an image or video file in a Gemini chat and then asking a "verification question". It also notes:
While other companies have started to adopt SynthID watermarks, Gemini can currently only recognize content created by Google AI tools.
Indeed, to detect SynthID watermarks added to content generated by ChatGPT, for example, you have to use OpenAI's SynthID detection tool (which comes as a standalone website instead of a chatbot integration).
SynthID tools are frequently used by fact-checkers and journalists to determine if AI was used in the creation or manipulation of a piece of content, with the presence of such a watermark being a very strong indicator (unlike the absence, which can still mean a different tool was used...).
First, we used Google's Gemini to create this image of a confused-looking robot:
Next we used "Organic Intelligence" (i.e. a human with a camera phone) to take this picture of a confused-looking fact-checker:
(Image source: Maarten Schenk)
Next, we had Gemini animate the picture of the robot into a short video, and we recorded a brief clip of the human too:
You can download the original files here if you want to repeat the experiment:
Next, we opened four separate chats in Gemini, in four different browser tabs. We uploaded each of the videos and photos and asked Gemini "Does this image/video have a SynthID watermark in it?". In the chats where we uploaded the real content first, we then uploaded the AI-generated ones and repeated the question. In the chats where we uploaded the AI-generated content first we then uploaded the real content and asked the question again too.
After each question, the following graphic briefly flashed up in the chat window each time, indicating Gemini was contacting a detection tool:
(Image source: Gemini screenshot)
The results were that in all of the cases the chatbot repeated the conclusion about the first item when asked the same question about the second one.
You can find archived versions of the four chats here , here , here and here .
AI image first, real one second:
AI video first, real one second:
The pattern seems to be consistent and repeatable. Several Lead Stories staffers were able to reproduce it using different images and videos. It even holds true if the content types in the chat are different. Here is Gemini not finding a SynthID watermark in a video which most definitely does have one, all because it was fed a SynthID-free image first (chat archived here ):
It is not unreasonable to speculate the issue is related to Google integrating the SynthID detector in the Gemini chatbot instead of offering it as a standalone tool (like OpenAI did with theirs). At some point between calling the detector and returning the results, it appears Gemini gets confused about which results to return.
When verifying content using Google Gemini, for now, we recommend always starting a fresh chat for each item, using the button in the sidebar:
(Image source: Gemini screenshot)
We will update this story when we receive a response from Google.
Maarten Schenk is the co-founder and COO/CTO of Lead Stories and an expert on fake news and hoax websites. He likes to go beyond just debunking trending fake news stories and is endlessly fascinated by the dazzling variety of psychological and technical tricks used by the people and networks who intentionally spread made-up things on the internet.
Previous Article
Fact Check: UNSOURCED Wikipedia Page Does NOT Prove Egypt-Argentina Match Referee Francois Letexier Grew Up In 'Orthodox Jewish Family'
About Us
Lead Stories is a fact checking website that is always looking for the latest false, misleading, deceptive or inaccurate stories, videos or images going viral on the internet.
Spotted something? Let us know! .
Verified signatory of the IFCN Code of Principles
Founding sponsor of Indicator's Show & Tell Podcast
Please select all the ways you would like to hear from Lead Stories LLC:
You can unsubscribe at any time by clicking the link in the footer of our emails. For information about our privacy practices, please visit our website.
We use Mailchimp as our marketing platform. By clicking below to subscribe, you acknowledge that your information will be transferred to Mailchimp for processing. Learn more about Mailchimp's privacy practices here.
Get more fact-checks in your Google Search results by setting up leadstories.com as one of your preferred sources .
Actors Acting
Fact Check: FAKE Video Shows Nurse and Cops Falsely Accusing Black Man of Being Drunk in Hospital
Already There
Fact Check: Mitch McConnell's Wife Elaine Chao Did NOT Go To China After He Was Hospitalized With 'Cardiac Arrest' -- Had Been There For Days
AI Mitch
Fact Check: FAKE Photo Of Mitch McConnell On Life Support Was Generated Using OpenAI
Didn't Post It
Fact Check: FAKE Screenshot Shows Trump Asking Mullin To Lock Up Communists In ICE Detention Centers -- No Such Truth Social Post
Altered Video
Fact Check: EDITED Video Of Erling Haaland Eating And Scaring Himself -- Is A Face Swap Of A Chinese Slapstick Comedy Video
Made With AI
Fact Check: FAKE 'Before Patriot Front' Video Seems To Show Woman From Reuters Photo--But It Was Made With Google AI
Manufactured
Fact Check: NO $3 Million Suit By Family Of Man Who Died Trying To Steal EV Charger Wires -- Fake Case, Fake Video
Gets Confused
Warning: Google Gemini SynthID AI Watermark Detector Appears To Mix Up Results In Same Chat -- Consistently Shows False Positives And Negatives Under Certain Conditions
Unsupported
Fact Check: UNSOURCED Wikipedia Page Does NOT Prove Egypt-Argentina Match Referee Francois Letexier Grew Up In 'Orthodox Jewish Family'
Another Match
Fact Check: Image Does NOT Show FIFA President Infantino Making Phone Call During Egypt-Argentina Match -- He Was At A Different Game
No Such Report
Fact Check: FAKE France 24 Video Does NOT Accurately Summarize UPA-Related Monument Legislation
Joke To Jeff
Fact Check: REAL Email In Epstein Files Contains Joke About Mitt Romney's Wife And Barack Obama
Old Photo
Fact Check: Image Trump Posted as 'Retribution' Against Iran Does NOT Show July 2026 U.S. Bombing -- It Is June 2025 Israeli Attack
Prebunk
Prebunk: London Calling To Stop Falling For Fake UK Protest Videos -- AI Is Leading You Down The River Thames
Share
Tweet
Home
About
IFCN Code of Principles
How we work
Satire Policy
Privacy Policy
Corrections Policy
Corrections
Archive
Contact
Български
English
Espanol
Magyar
Română
Русский
Slovensky
Українська
Lead Stories is a fact checking website that is always looking for the latest false, misleading, deceptive or inaccurate stories, videos or images going viral on the internet.
Spotted something? Let us know! .
Verified signatory of the IFCN Code of Principles
Founding sponsor of Indicator's Show & Tell Podcast
© 2015-2024 Lead Stories LLC - All rights reserved.
Lead Stories
