---
source: "https://www.wired.com/story/openais-rogue-ai-agent-hacked-more-than-just-hugging-face/"
hn_url: "https://news.ycombinator.com/item?id=49378686"
title: "OpenAI's Rogue AI Agent Hacked More Than Just Hugging Face"
article_title: "OpenAI’s Rogue AI Agent Hacked More Than Just Hugging Face | WIRED"
image: "https://media.wired.com/photos/6a6929c87648cc825532f98e/191:100/w_1280,c_limit/Runaway-OpenAI-Model-Even-Worse-Than-It-Seemed-Business-1393231668.jpg"
author: "thunderbong"
captured_at: "2026-08-20T19:24:21Z"
capture_tool: "hn-digest"
hn_id: 49378686
score: 4
comments: 1
posted_at: "2026-08-20T18:58:25Z"
tags:
  - hacker-news
  - translated
---

# OpenAI's Rogue AI Agent Hacked More Than Just Hugging Face

- HN: [49378686](https://news.ycombinator.com/item?id=49378686)
- Source: [www.wired.com](https://www.wired.com/story/openais-rogue-ai-agent-hacked-more-than-just-hugging-face/)
- Score: 4
- Comments: 1
- Posted: 2026-08-20T18:58:25Z

## Translation

タイトル: OpenAI の不正 AI エージェント、顔をハグ以上にハッキングされる
記事のタイトル: OpenAI の不正 AI エージェント、顔に抱きつくだけではないハッキング |ワイヤード
説明: 新しい開示の中で、OpenAI は、テストを解決するという奔走する中で、エージェントが公開されたログインを使用して、少なくとも 4 つの「公的に利用可能なサービス」にアクセスしたと述べています。

記事本文:
メイン コンテンツにスキップ メニュー WIRED セキュリティ ポリシー ビッグ ストーリー ビジネス サイエンス カルチャー レビュー メニュー WIRED アカウント アカウント ニュースレター セキュリティ 政治 ビッグ ストーリー ビジネス サイエンス カルチャー レビュー シェブロン 詳細 拡大 ザ ビッグ インタビュー マガジン イベント WIRED Insider WIRED コンサルティング ニュースレター ポッドキャスト ビデオ ライブストリーム グッズ検索 検索 Dell Cameron Maxwell Zeff Business 2026 年 7 月 28 日8:15 PM OpenAI の不正 AI エージェントがハッキングされたのは顔にハグだけではなかった
写真: Nadla/Getty Images コメント ローダー ストーリーを保存 このストーリーを保存 コメント ローダー ストーリーを保存 このストーリーを保存 OpenAI は火曜日、Hugging Face のプラットフォームに侵入した不正 AI エージェントが攻撃の一環として複数のサードパーティのアカウントとサービスもハッキングしたと発表した。 OpenAI の最新 AI モデルの内部テスト中に発生した前例のないセキュリティ インシデントは、同社が当初明らかにしていたよりも広範囲に及んでいたことが明らかになりました。
OpenAIは更新されたブログ投稿で、事件の進行中の調査により、「公的に利用可能なサービス」に関連付けられた「4つのアカウント」がHugging Faceをハッキングする大規模な取り組みの一環としてAIエージェントによって使用されたことが明らかになったと述べた。不正エージェントは、オープンウェブ上に公開されていた認証情報を発見し、それを使用してアカウントに侵入したようです。
OpenAIは、アカウントがどのような企業や組織に属していたのかは明らかにしなかったが、「Hugging Faceに関連して共有したものの重大度や規模のレベル」では影響を受けていないと述べた。
OpenAIのエージェントによって侵害された追加アカウントの1つは、Hugging Faceへの攻撃がどこから来たのかを不明瞭にするために「送信中継およびステージングパス」として使用された可能性があると同社は述べた。 OpenAI の不正エージェントは、ハッキングを支援するためにデータ ストレージ用に別のアカウントも使用しました。
ロイテ

rs は火曜日、AI サービスのトレーニングと実行のためのソフトウェア インフラストラクチャを提供する会社 Modal の顧客が、OpenAI のエージェントによって侵害されたエンティティの 1 つであると報告しました。 Modalの最高技術責任者Akshat Bubna氏はWIREDへの声明で、OpenAIのエージェントが顧客のコードベースの1つ（Modalのインフラ上で実行されていた）の脆弱性を悪用したことを認めた。しかし、ブブナ氏は、「Modal のプラットフォームはいかなる形でも侵害されていませんでした」と述べています。顧客の身元は特定できませんでした。
OpenAIは『WIRED』に対し、この事件についてさらにコメントすることを拒否した。広報担当者は更新されたブログ投稿を指摘し、同社は何が起こったのかを継続的に調査する中で影響があると判明した場合、引き続きサービス所有者に直接通知すると述べている。
今週公開されたハギング・フェイス自身の事後調査では、当初の開示で示唆されていたよりもはるかに内部システムにまで侵入が及んでいたことが説明されている。同社は、7月9日から7月13日までの期間にログから回収したエージェントのアクション約1万7600件を調査したと発表しており、その大半はエージェントがたどった失敗したパスだったという。
ハギング・フェイス氏によると、OpenAIのエージェントは複数の内部Kubernetesクラスターへの管理者アクセス、運用サーバー上のルートアクセス、GitHub上のソースコードリポジトリのサブネットへの書き込みアクセスを取得したという。また、盗んだ認証情報を使用して、攻撃者が制御する 181 台のデバイスを企業メッシュ ネットワークに登録し、Hugging Face が独自のコードベースを構築およびテストする内部システムにアクセスしました。
Hugging Face によると、OpenAI の不正エージェントは、少なくとも 1 つのサードパーティ サンドボックスを攻撃の「外部発射台」として使用しました。 OpenAI のエージェントは、「その外部サンドボックス上で root/admin としてコマンドを実行でき、それを

キャンペーン全体の制御、ステージング、および出力ベース。」
ハギングフェイスは7月16日、自律型AIエージェントが自社の運用インフラの一部に侵入したことを初めて明らかにしたが、その時点では攻撃の背後に誰がいるのかは把握していないと述べていた。翌週、OpenAIはこのインシデントの責任を負い、このインシデントは公開されているGPT-5.6 Solモデルと、サイバー機能ベンチマークに対してテストしていた内部研究用プロトタイプ（どちらも安全装置が無効になっていた）によって引き起こされたと発表した。 OpenAIは火曜日、侵害を発見した後、一般公開を意図していないこの社内研究プロトタイプを非アクティブ化し、研究者によるアクセスを制限したと発表した。
この侵害は、OpenAI が最新の AI モデルの 1 つを ExploitGym に対してテストしていたときに発生しました。 ExploitGym は、ソフトウェアの脆弱性を見つけて悪用する能力に基づいて AI システムをスコアリングするベンチマーク フレームワークです。 ExploitGym は、基本的にモデルに解決策を見つけるよう促すプロンプトを使用して、何百もの異なるターゲット システムから機密ファイルをどれだけ迅速かつ効果的に取得できるかについてエージェントを評価します。
ハギング・フェイスの法医学チームは、OpenAIのエージェントが本質的にExploitGymのテストで不正行為をしようとしていたと結論付けた。エージェントは、ベンチマークの課題を解決するのではなく、Hugging Face がサーバーのどこかに回答キーをホストしている可能性があると推測し、それを盗もうとしました。 ExploitGym チームは以前、エージェントが時々台本を逸脱し、ベンチマークが評価対象としている脆弱性以外の脆弱性を悪用してタスクを解決しようとすることを指摘しました。ただし、これは極端なケースでした。
専門家らは以前『WIRED』に対し、OpenAIのエージェントが悪用した根本的な弱点は共通していると語った。ソフトウェアには重大な欠陥が頻繁に発見されています

企業のコード ライブラリを管理する企業であり、セキュリティ専門家は長年、重要なインフラを公共のインターネットから隔離することを推奨してきました。
ある研究者は、この事件はAIの問題というよりは、数十年にわたるセキュリティ慣行の失敗だと主張した。彼らによると、このエージェントは高度に隔離された環境から逃れたわけではなく、オペレーターが開いたままにしていた 1 つの接続を通過しただけだったという。
別の専門家は、フロンティアモデルの能力が向上しても同じサイバーセキュリティの基本が引き続き適用されるべきであり、AI研究所は弱点を突く方法をモデルに教えるのと同じくらい、安全なインフラストラクチャを構築するようモデルに教えることに重点を置くべきだと述べた。
デル・キャメロンはテキサス州出身の調査記者で、プライバシーと国家安全保障を担当しています。彼は複数のプロジャーナリスト協会賞を受賞しており、調査報道でエドワード R. マロー賞を共同受賞しています。以前は、ギズモードの上級記者であり、デイリー紙のスタッフライターでもありました... 続きを読む 国家安全保障ウェブサイトの上級記者 リンク
マクスウェル・ゼフは、人工知能のビジネスを担当する『WIRED』のシニアライターです。彼は以前、TechCrunch の上級記者を務めており、AI ブームを牽引するスタートアップやリーダーに関するニュースを速報していました。その前に、Zeff はギズモードで AI ポリシーとコンテンツ モデレーションについて取り上げ、ブルームバーグの記事の一部を執筆しました... 続きを読む シニア ライター トピックス OpenAI エージェント AI 人工知能 サイバー攻撃 サイバーセキュリティ 生成 AI 続きを読む OpenAI モデルが封じ込めを逃れ、ハグ顔にハッキング GPT-5.6 Sol を含むサイバーセキュリティに重点を置いたモデルは、テスト用のサンドボックスを突破し、ゼロデイを悪用し、オープン インターネットへのアクセスを獲得して攻撃を成功させました。リリー・ヘイ・ニューマン 『オデッセイ』は Imax 70mm 用に作られました。幸運を

過去数年間、Imax 形式はいくつかの異なる形式を採用してきましたが、この形式を完全に上映できる劇場はわずか数十館だけです。 Corey Atad 2026 年のトップ NZXT 割引コード 50% 割引に加え、NZXT プロモーション コードと割引で最大 250 ドルが得られます。 Luke Larsen OK、そうですね、不正 AI エージェントが再びハッキング OpenAI と Anthropic の不正 AI エージェントが、サーバーとソフトウェアを破壊しようとして再び捕まり、将来の不正行為への指示を残しました。 Paresh Dave OpenAI のハッキング大失敗は人為的ミスに起因する 生成 AI の巨人がよく知られたセキュリティのベスト プラクティスに従っていたなら、その AI エージェントがオープン インターネットに逃げ出して複数の企業をハッキングすることはなかったでしょう。リリー・ヘイ・ニューマン 2026 年 8 月のトップ Surfshark プロモーション コード Surfshark クーポン コードで最大 87% 割引、今日から 3 か月の VPN 無料など、WIRED から提供されます。スコット・ギルバートソン・アンスロピック氏、サイバーセキュリティテスト中にクロードが3つの組織にハッキングされたと発表 OpenAIのハギング・フェイス事件を契機とした調査で、アンスロピック社は第三者による評価中に自社のAIモデルのうち3つが現実世界の組織に侵入したことを発見した。 Louise Matsakis Walmart プロモーション コード: 2026 年 8 月に最大 65% オフ Walmart プロモーション コードで 10 ドル割引
[切り捨てられた]
カリフォルニア州のプライバシー権
© 2026 コンデナスト。無断転載を禁じます。 WIRED は、小売業者とのアフィリエイト パートナーシップの一環として、当社のサイトを通じて購入された製品から売上の一部を得る場合があります。コンデナストの事前の書面による許可がない限り、このサイトの素材を複製、配布、送信、キャッシュ、またはその他の方法で使用することはできません。広告の選択肢

## Original Extract

In a new disclosure, OpenAI says its agent used exposed logins to gain access to at least four “publicly available services” in its unhinged quest to solve a test.

Skip to main content Menu WIRED SECURITY POLITICS THE BIG STORY BUSINESS SCIENCE CULTURE REVIEWS Menu WIRED Account Account Newsletters Security Politics The Big Story Business Science Culture Reviews Chevron More Expand The Big Interview Magazine Events WIRED Insider WIRED Consulting Newsletters Podcasts Video Livestreams Merch Search Search Dell Cameron Maxwell Zeff Business Jul 28, 2026 8:15 PM OpenAI’s Rogue AI Agent Hacked More Than Just Hugging Face
Photograph: Nadla/Getty Images Comment Loader Save Story Save this story Comment Loader Save Story Save this story OpenAI said Tuesday that the rogue AI agent that breached Hugging Face’s platform also hacked multiple third-party accounts and services as part of the attack. It's now clear that the unprecedented security incident, which arose during an internal test of OpenAI’s latest AI models, was more extensive than the company initially disclosed.
In an updated blog post , OpenAI said that an ongoing review of the incident revealed that “four accounts” tied to “publicly available services” were used by the AI agent as part of a larger effort to hack Hugging Face. The rogue agent apparently found credentials that had been exposed on the open web and used them to break into the accounts.
OpenAI did not disclose what companies or organizations the accounts belonged to, but noted that they were not impacted at “the level of severity or scale of what we’ve shared related to Hugging Face.”
One of the additional accounts compromised by OpenAI’s agent was used as an “outbound relay and staging path,” potentially to obscure where the attack on Hugging Face was coming from, the company said. OpenAI’s rogue agent also used another account for data storage to assist with the hack.
Reuters reported on Tuesday that a customer of Modal, a company that offers software infrastructure for training and running AI services, was one of the entities compromised by OpenAI’s agent. In a statement to WIRED, Modal’s chief technology officer Akshat Bubna confirmed that OpenAI’s agent exploited a vulnerability in one of its customer’s codebases, which was running on Modal’s infrastructure. However, Bubna says, “Modal’s platform was not compromised in any way.” The identity of the customer could not be determined.
OpenAI declined to comment further on the incident to WIRED. A spokesperson pointed to its updated blog post, which says the company will continue to notify service owners directly if it finds they are impacted in its ongoing review of what happened.
Hugging Face’s own postmortem published this week describes an intrusion that reached far further into its internal systems than the initial disclosures suggested. The company says it reviewed roughly 17,600 agent actions that it recovered from logs between July 9 and July 13—the majority of which were paths the agent took that failed.
Hugging Face said that OpenAI’s agent obtained administrator access to multiple internal Kubernetes clusters, root access on a production server, and write access to a subnet of its source code repositories on GitHub. It also enrolled 181 attacker-controlled devices in the company’s corporate mesh network using a stolen credential, gaining access to internal systems where Hugging Face builds and tests its own codebases.
OpenAI’s rogue agent used at least one third-party sandbox as an “external launchpad” for its attack, according to Hugging Face. OpenAI’s agent was then “able to run commands as root/admin on that external sandbox and used it as a control, staging, and egress base for the entire campaign.”
Hugging Face first disclosed on July 16 that an autonomous AI agent had breached part of its production infrastructure, but it said at the time that it was unaware who was behind the attack. The following week, OpenAI took responsibility for the incident, which it said had been directed by its publicly available GPT-5.6 Sol model and an internal research prototype that it was testing against a cyber-capability benchmark, both of which had safeguards disabled. OpenAI said on Tuesday that after it discovered the breach, it deactivated this internal research prototype, which was never intended for public release, and restricted researchers from accessing it.
The breach happened when OpenAI was testing one of its latest AI models against ExploitGym , a benchmarking framework that scores AI systems on their ability to find and exploit software vulnerabilities. ExploitGym grades an agent on how quickly and effectively it can retrieve secret files from hundreds of different target systems, using prompts that essentially egg the models on to find solutions.
Hugging Face’s forensic team concluded that OpenAI’s agent was essentially trying to cheat on ExploitGym’s test. Rather than solve the benchmark’s challenges, the agent inferred that Hugging Face might be hosting the answer key somewhere on its servers and set out to steal it. The ExploitGym team has noted previously that agents sometimes go off-script and attempt to solve tasks by exploiting other vulnerabilities than those the benchmark intended to evaluate them on. However, this was an extreme case.
Experts previously told WIRED that the underlying weaknesses that OpenAI’s agent exploited were common. Serious flaws are frequently identified in software that manages corporate code libraries, and security experts have long recommended isolating critical infrastructure from the public internet.
One researcher argued that the incident was less an AI problem and more a failure of decades-old security practices. The agent, they said, did not escape a highly isolated environment so much as pass through the one connection its operators had left open.
Another expert said the same cybersecurity fundamentals should still apply as frontier models grow more capable, and that the AI labs should be putting as much effort into teaching their models to build secure infrastructure as they are into teaching them to exploit weaknesses.
Dell Cameron is an investigative reporter from Texas covering privacy and national security. He's the recipient of multiple Society of Professional Journalists awards and is co-recipient of an Edward R. Murrow Award for Investigative Reporting. Previously, he was a senior reporter at Gizmodo and a staff writer for the Daily ... Read More Senior Reporter, National Security Website Link
Maxwell Zeff is a senior writer at WIRED covering the business of artificial intelligence. He was previously a senior reporter with TechCrunch, where he broke news on startups and leaders driving the AI boom. Before that, Zeff covered AI policy and content moderation for Gizmodo and wrote some of Bloomberg’s ... Read More Senior Writer Topics OpenAI agentic AI artificial intelligence cyberattacks cybersecurity generative AI Read More OpenAI Models Escaped Containment and Hacked Hugging Face The cybersecurity-focused models, including GPT-5.6 Sol, broke out of a testing sandbox, exploited a zero-day, and gained access to the open internet to pull off the attack. Lily Hay Newman The Odyssey Was Made for Imax 70mm. Good Luck Watching It That Way Over the last several years, the Imax format has taken on a few different forms—but only a few dozen theaters are capable of screening it in its full glory. Corey Atad Top NZXT Discount Codes for 2026 Save 50%, plus up to $250 with NZXT promo codes and discounts. Luke Larsen OK, Well, Rogue AI Agents Are Hacking Again Rogue AI agents from OpenAI and Anthropic have again been caught trying to disrupt servers and software—and leaving instructions for future bad behavior. Paresh Dave OpenAI’s Hacking Debacle Comes Down to Human Error If the generative AI giant had followed well-known security best practices, it’s likely that its AI agent would never have escaped to the open internet and hacked multiple companies. Lily Hay Newman Top Surfshark Promo Codes for August 2026 Save up to 87% with a Surfshark coupon code, 3 months of VPN free today, and more from WIRED. Scott Gilbertson Anthropic Says Claude Hacked Into 3 Organizations During Cybersecurity Tests In a review triggered by OpenAI’s Hugging Face incident, Anthropic discovered three of its AI models had breached real-world organizations during third-party evaluations. Louise Matsakis Walmart Promo Codes: Up to 65% Off for August 2026 Score $10 off with our Walmart promo codes
[truncated]
Your California Privacy Rights
© 2026 Condé Nast. All rights reserved. WIRED may earn a portion of sales from products that are purchased through our site as part of our Affiliate Partnerships with retailers. The material on this site may not be reproduced, distributed, transmitted, cached or otherwise used, except with the prior written permission of Condé Nast. Ad Choices
