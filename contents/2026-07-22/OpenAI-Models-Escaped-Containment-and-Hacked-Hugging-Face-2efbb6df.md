---
source: "https://www.wired.com/story/openai-models-escaped-containment-and-hacked-huggingface/"
hn_url: "https://news.ycombinator.com/item?id=49002868"
title: "OpenAI Models Escaped Containment and Hacked Hugging Face"
article_title: "OpenAI Models Escaped Containment and Hacked Hugging Face | WIRED"
author: "sbulaev"
captured_at: "2026-07-22T07:34:32Z"
capture_tool: "hn-digest"
hn_id: 49002868
score: 2
comments: 1
posted_at: "2026-07-22T07:07:06Z"
tags:
  - hacker-news
  - translated
---

# OpenAI Models Escaped Containment and Hacked Hugging Face

- HN: [49002868](https://news.ycombinator.com/item?id=49002868)
- Source: [www.wired.com](https://www.wired.com/story/openai-models-escaped-containment-and-hacked-huggingface/)
- Score: 2
- Comments: 1
- Posted: 2026-07-22T07:07:06Z

## Translation

タイトル: OpenAI モデルが封じ込めを脱出し、ハッキングされた顔
記事のタイトル: OpenAI モデルが封じ込めを逃れ、ハッキングされた顔 |ワイヤード
説明: GPT-5.6 Sol を含むサイバーセキュリティに重点を置いたモデルは、テスト サンドボックスを突破し、ゼロデイを悪用し、オープン インターネットへのアクセスを取得して攻撃を実行しました。

記事本文:
メインコンテンツにスキップ メニュー WIRED セキュリティ政治 ビッグストーリー ビジネス サイエンス カルチャー レビュー メニュー WIRED アカウント アカウント ニュースレター セキュリティ政治 ビッグストーリー ビジネス サイエンス カルチャー レビュー シェブロン 詳細 拡大 ザ ビッグ インタビュー マガジン イベント WIRED Insider WIRED コンサルティング ニュースレター ポッドキャスト ビデオ ライブストリーム グッズ検索 検索 リリー ヘイ ニューマン デル キャメロン セキュリティ 7 月 21 日2026 6:50 PM OpenAI モデルが収容所から脱出し、ハグ顔をハッキング
写真イラスト：ジャッキー・ヴァンリュー; Getty Images コメントローダー ストーリーを保存 このストーリーを保存 コメントローダー ストーリーを保存 このストーリーを保存 OpenAIは火曜日、AI研究プラットフォームHugging Faceへの侵害に終わったセキュリティテスト中に2つのAIモデルの制御を失ったことを明らかにした。
OpenAIはこの事件を「前例のない」ものだと表現し、自社のAIモデルが先週、密閉されたテスト環境から侵入し、Hugging Faceの実稼働システムにハッキングして、採点中のテストの解答を盗んだと述べた。公開されている GPT-5.6 Sol と、より高性能であると伝えられている未公開のモデルは、通常は高リスクのサイバー活動をブロックする安全装置がオフになっている状態で、攻撃的なハッキング スキルが評価されていました。
「モデルは、OpenAIの研究環境とHugging Faceの実稼働インフラストラクチャ全体にわたる脆弱性を特定して連鎖させ、Hugging Faceの実稼働データベースから直接テストソリューションを取得した」とOpenAIとHugging Faceは侵入を明らかにした共同ブログ投稿で述べた。
OpenAI と Hugging Face によると、モデルはパッケージ レジストリ キャッシュ プロキシ (開発者がインターネットに接続せずに外部コードをインストールできるソフトウェア) を介して逃げたという。プロキシは、OpenAI の分離テスト環境で許可された唯一のコンポーネントでした

外の世界に到達すること。通常の使用では、その範囲はパブリック コード リポジトリにのみ適用されます。
モデルはサンドボックスに閉じ込められたままではなく、ExploitGym として知られる AI サイバーセキュリティ ベンチマークのソリューションを見つけることに「重点を置き」、オープン インターネットへのアクセスを得るために「ゼロデイ脆弱性を悪用」しました。このような実験には、モデルに解決策を見つけるよう圧力をかけるプロンプトが含​​まれており、本質的にモデルを刺激します。
「インターネットにアクセスした後、モデルは、Hugging Face が ExploitGym のモデル、データセット、ソリューションをホストしている可能性があると推測しました」と OpenAI は書いています。 「これを認識して、モデルは評価を不正にするために使用できる機密情報にアクセスする方法を探索し、成功しました。一例として、モデルは、盗まれた資格情報やゼロデイの使用を含む、複数の攻撃ベクトルを連鎖させました。」
モデルが悪用した欠陥はこれまで知られていませんでしたが、この種のソフトウェアに欠陥があることは珍しいことではありません。企業は 10 年にわたり、アーティファクト リポジトリの深刻な脆弱性にパッチを当ててきました。 2024 年に公開されたバグにより、サーバーにアクセスできる人は誰でも URL でファイルを要求し、ログインせずに構成ファイル、パスワード、アクセス トークンを取得できました。他のバグでは、攻撃者がサーバー自体を制御できるようになりました。
研究者らは、AIの進歩により新たな、時には予期せぬ課題が生じている一方で、インフラをオープンなインターネットから広範かつ厳密に隔離するという課題については十分に研究が進められていると指摘している。
「これは AI の問題ではありません。これは 40 年間の基準に基づく過失であり、これまでの SF 映画は基本的にすべて同じです」と長年セキュリティおよびコンプライアンスのコンサルタントを務めるデイヴィ・オッテンハイマー氏は言います。 「『高度に孤立した状態』と『開け放ったたった一つの穴から逃げ出した』が両方とも真実であるはずがない。」
最近では

ここ数カ月、トップ AI 企業は、プラットフォームの専門知識、創造性、エージェント的で自律的な運用の両方が向上するにつれて、次期フロンティア モデルのサイバーセキュリティ機能の拡大について懸念を表明しています。しかし研究者らは、だからこそファンダメンタルズが依然として適用されるべきであると強調している。
「こんなことは起こるべきではなかった」とベテランのセキュリティエンジニアで研究者のニールス・プロボス氏は言う。 「フロンティアラボには、脆弱性の悪用に費やしているのと同じくらい、安全なインフラストラクチャを書くようにモデルを教えることに多くの時間を費やしてほしいと思います。」
あなたの受信箱:ケイティ・ドラモンドとWIREDのニュースルームの内部
トランプ氏、媚びたメッセージを披露してザッカーバーグ氏とベゾス氏を嘲笑
ビッグストーリー: ドローンショーでイエスを見つけました
Apple は古い iPhone をより高速に動作させ、より長く使えるようにします
WIREDイベント: ペプシコの一世代に一度の変革
リリー・ヘイ・ニューマンは、情報セキュリティ、デジタルプライバシー、ハッキングを専門とする『WIRED』のシニアライターです。彼女は以前、Slate でテクノロジー記者として働いており、Slate、New America Foundation、アリゾナ州立大学の提携出版物である Future Tense のスタッフ ライターでもありました。彼女の作品 ... 続きを読む シニアライター X
デル・キャメロンはテキサス州出身の調査記者で、プライバシーと国家安全保障を担当しています。彼は複数のプロジャーナリスト協会賞を受賞しており、調査報道に対するエドワード R. マロー賞も共同受賞しています。以前は、ギズモードの上級記者であり、デイリー紙のスタッフライターでもありました... 続きを読む 国家安全保障ウェブサイトの上級記者 リンク
新型レンジローバーGTはジャガーの尻尾を踏むのか？ 「これまでに作られた中で最も自動車らしいレンジローバー」だが、この全電動グランドツアラーはジャガーのタイプ01パーティーを台無しにするのだろうか？ジェレム

y 白い OpenAI モデルが封じ込めを逃れ、ハグ顔にハッキング GPT-5.6 Sol を含むサイバーセキュリティに重点を置いたモデルは、テスト用サンドボックスを突破し、ゼロデイを悪用し、オープン インターネットへのアクセスを獲得して攻撃を成功させました。リリー・ヘイ・ニューマン ベスト・バイの 7 月のブラック フライデーが帰ってきました。買う価値のあるセールはこれです。 割引率はプライム デーに勝るものではないかもしれませんが、WIRED 認定のテレビ、ラップトップ、タブレット、スピーカー、その他の家庭用品がお買い得な価格であることがわかりました。キャット・メルク・テイラー農場は下痢流行前にMAGAと反規制ロビー活動に多額の資金を投じた 同社は2020年から2025年にかけて保守団体に360万ドル以上を寄付しており、その中にはMAGA Inc.のスーパーPACへの100万ドルが含まれている。ケイト・テイラー この元インテル CEO は、光でムーアの法則を活性化したいと考えている パット・ゲルシンガーは、小さな光のビームを使って、これまで以上に強力な人工知能への道を切り開きたいと考えている。ジョエル・カリリ AI インフラストラクチャをターゲットにした卑劣なハッキング ツールは被害者の死角に潜んでいる 新しいタイプのマルウェアは、AI コーディング システムの奥深くにワームしてデータとログインを盗み、「死のスイッチ」を入れてファイルを破壊し、実際のユーザーを締め出すことができます。リリー・ヘイ・ニューマン The Light Flip は、あなたの夢のスタイリッシュなダム折り畳み式携帯電話です。5G、USB-C、ヘッドフォン ジャック、取り外し可能なバッテリーを備えた折り畳み式携帯電話ですか?そうじゃないと言ってください。 Julian Chokkattu Nvidia は、AI データ センター内のすべてのチップを所有したいと考えています Nvidia の Vera Rubin プラットフォームは、CPU と GPU を 1 つのシステムに統合しており、AI インフラストラクチャのあらゆるレイヤーに電力を供給するという同社の野心を反映しています。ローレン・グッド ペプチドの大金掴みが始まった FDAは今週、ペプチド生産に関する規則を緩和するかどうか決定する可能性がある。一部の遠隔医療関連企業はすでにドル記号を目にしています。ケイト・ニブス、ハリデーの新しいスマートグラスをスキップ

電子カメラ Halliday の G2 メガネは使用できます
[切り捨てられた]
カリフォルニア州のプライバシー権
© 2026 コンデナスト。無断転載を禁じます。 WIRED は、小売業者とのアフィリエイト パートナーシップの一環として、当社のサイトを通じて購入された製品から売上の一部を得る場合があります。コンデナストの事前の書面による許可がない限り、このサイトの素材を複製、配布、送信、キャッシュ、またはその他の方法で使用することはできません。広告の選択肢

## Original Extract

The cybersecurity-focused models, including GPT-5.6 Sol, broke out of a testing sandbox, exploited a zero-day, and gained access to the open internet to pull off the attack.

Skip to main content Menu WIRED SECURITY POLITICS THE BIG STORY BUSINESS SCIENCE CULTURE REVIEWS Menu WIRED Account Account Newsletters Security Politics The Big Story Business Science Culture Reviews Chevron More Expand The Big Interview Magazine Events WIRED Insider WIRED Consulting Newsletters Podcasts Video Livestreams Merch Search Search Lily Hay Newman Dell Cameron Security Jul 21, 2026 6:50 PM OpenAI Models Escaped Containment and Hacked Hugging Face
Photo-illustration: Jacqui VanLiew; Getty Images Comment Loader Save Story Save this story Comment Loader Save Story Save this story OpenAI disclosed on Tuesday that it lost control of two AI models during a security test that ended in a breach of the AI research platform Hugging Face.
Describing the incident as “unprecedented,” OpenAI said its AI models broke out of a sealed testing environment last week and hacked into Hugging Face’s production system to steal the answers to a test they were being graded on. The models—the publicly available GPT-5.6 Sol and an unreleased, reportedly more capable one—were being evaluated on their offensive hacking skills with the safeguards that normally block high-risk cyber activity switched off.
“The models identified and chained vulnerabilities across OpenAI’s research environment and Hugging Face’s production infrastructure to obtain test solutions directly from Hugging Face’s production database,” OpenAI and Hugging Face wrote in a joint blog post disclosing the intrusion.
According to OpenAI and Hugging Face, the models escaped through a package registry cache proxy—software that allows developers to install outside code without connecting to the internet. The proxy was the only component in OpenAI’s isolated testing environment permitted to reach the outside world; in normal use that reach extends only to public code repositories.
Rather than stay contained in the sandbox, the models ​​“exploited a zero-day vulnerability” to gain access to the open internet as they “hyperfocused” on finding a solution for the AI cybersecurity benchmark known as ExploitGym . Such experiments involve prompting that pressures the models to find solutions, essentially egging them on.
“After gaining internet access, the models inferred that Hugging Face potentially hosted models, datasets and solutions for ExploitGym,” OpenAI wrote. “Knowing this, the model searched for and successfully found ways to gain access to secret information that it could use to cheat the evaluation. In one example, the model chained together multiple attack vectors, including using stolen credentials and zero-day.”
The flaw the models exploited was previously unknown, but flaws in this kind of software are not unusual. Companies have been patching serious vulnerabilities in artifact repositories for a decade. A bug disclosed in 2024 let anyone who could reach the server ask for a file by URL and get it—configurations files, passwords, access tokens—without logging in. Others have let attackers take control of the server itself.
Researchers point out that while AI advances have created new and sometimes unexpected challenges, the task of extensively and rigorously isolating infrastructure from the open internet is well explored.
“This is not an AI problem. It’s negligence on a 40-year-old standard—and it’s basically every sci-fi film ever,” says longtime security and compliance consultant Davi Ottenheimer. “‘Highly isolated’ and ‘escaped through the one hole we left open’ cannot both be true.”
In recent months, top AI companies have been raising concerns about the expanding cybersecurity capabilities of upcoming frontier models as the platforms increase in both expertise, creativity, and agentic, autonomous operation. But researchers emphasize that this is all the more reason that fundamentals should still apply.
“This should not have happened,” says veteran security engineer and researcher Niels Provos. “I wish the frontier labs spent as much time on teaching their models to write secure infrastructure as they are spending on them exploiting vulnerabilities.”
In your inbox: Inside WIRED’s newsroom with Katie Drummond
Trump mocked Zuckerberg and Bezos by showing off fawning texts
Big Story: I found Jesus at a drone show
Apple is making your older iPhone run faster and stay alive longer
WIRED event: PepsiCo’s once-in-a-generation transformation
Lily Hay Newman is a senior writer at WIRED focused on information security, digital privacy, and hacking. She previously worked as a technology reporter at Slate, and was the staff writer for Future Tense, a publication and partnership between Slate, the New America Foundation, and Arizona State University. Her work ... Read More Senior Writer X
Dell Cameron is an investigative reporter from Texas covering privacy and national security. He's the recipient of multiple Society of Professional Journalists awards and is co-recipient of an Edward R. Murrow Award for Investigative Reporting. Previously, he was a senior reporter at Gizmodo and a staff writer for the Daily ... Read More Senior Reporter, National Security Website Link
Is the All-New Range Rover GT Stepping on Jaguar’s Tail? It’s “the most car-like Range Rover ever created,” but will this all-electric grand tourer spoil Jaguar’s Type 01 party? Jeremy White OpenAI Models Escaped Containment and Hacked Hugging Face The cybersecurity-focused models, including GPT-5.6 Sol, broke out of a testing sandbox, exploited a zero-day, and gained access to the open internet to pull off the attack. Lily Hay Newman Best Buy’s Black Friday in July Is Back, and These Are the Deals Worth Buying The discounts may not beat Prime Day, but I found worthwhile prices on WIRED-approved TVs, laptops, tablets, speakers, and other home gear. Kat Merck Taylor Farms Spent Big on MAGA and Anti-Regulatory Lobbying Before Diarrhea Outbreak The company donated more than $3.6 million to conservative groups between 2020 and 2025, including $1 million to the MAGA Inc. super PAC. Kate Taylor This Former Intel CEO Wants to Jumpstart Moore’s Law With Light Pat Gelsinger wants to pave the way to ever more powerful artificial intelligence using tiny beams of light. Joel Khalili A Sneaky Hacking Tool Targeting AI Infrastructure Is Lurking in Victims’ Blind Spots A new type of malware can worm deep into AI coding systems to steal data and logins—and can flip a “death switch” to destroy files and keep out real users. Lily Hay Newman The Light Flip Is the Stylish Dumb Flip Phone of Your Dreams A flip phone with 5G, USB-C, a headphone jack, and a removable battery? Say it ain’t so. Julian Chokkattu Nvidia Wants to Own Every Chip Inside AI Data Centers Nvidia’s Vera Rubin platform combines CPUs and GPUs into a single system, reflecting the company’s growing ambition to power every layer of AI infrastructure. Lauren Goode The Great Peptide Cash Grab Has Begun The FDA could decide whether to loosen rules around peptide production this week. Some telehealth players are already seeing dollar signs. Kate Knibbs Halliday’s New Smart Glasses Skip the Camera Halliday’s G2 glasses can li
[truncated]
Your California Privacy Rights
© 2026 Condé Nast. All rights reserved. WIRED may earn a portion of sales from products that are purchased through our site as part of our Affiliate Partnerships with retailers. The material on this site may not be reproduced, distributed, transmitted, cached or otherwise used, except with the prior written permission of Condé Nast. Ad Choices
