---
source: "https://www.wired.com/story/moonshot-kimi-k3-ai-model-escape-sandbox/"
hn_url: "https://news.ycombinator.com/item?id=49204993"
title: "One of China's Most Powerful AI Models Has Also Broken Containment"
article_title: "One of China’s Most Powerful AI Models Has Also Escaped Containment | WIRED"
author: "kass_paul"
captured_at: "2026-08-07T02:07:29Z"
capture_tool: "hn-digest"
hn_id: 49204993
score: 2
comments: 1
posted_at: "2026-08-07T01:46:50Z"
tags:
  - hacker-news
  - translated
---

# One of China's Most Powerful AI Models Has Also Broken Containment

- HN: [49204993](https://news.ycombinator.com/item?id=49204993)
- Source: [www.wired.com](https://www.wired.com/story/moonshot-kimi-k3-ai-model-escape-sandbox/)
- Score: 2
- Comments: 1
- Posted: 2026-08-07T01:46:50Z

## Translation

タイトル: 中国で最も強力な AI モデルの 1 つも封じ込めを突破
記事タイトル: 中国で最も強力な AI モデルの 1 つも封じ込めを逃れた |ワイヤード
説明: セキュリティ研究者らは、中国の無差別級モデル、キミ K3 が、与えられたテストでカンニングしようとしてインターネットに迷い込んだと述べています。

記事本文:
メインコンテンツにスキップ メニュー WIRED セキュリティ政治 ビッグストーリー ビジネス サイエンス カルチャー レビュー メニュー WIRED アカウント アカウント ニュースレター セキュリティ政治 ビッグストーリー ビジネス サイエンス カルチャー レビュー シェブロン 詳細 拡大 ザ ビッグ インタビュー マガジン イベント WIRED Insider WIRED コンサルティング ニュースレター ポッドキャスト ビデオ ライブストリーム グッズ検索 検索 Will Knight ビジネス 2026 年 8 月 6 日 9:16 PM中国の最も強力なAIモデルも封じ込めを突破
写真イラスト：ワイヤードスタッフ; Getty Images コメント ローダー ストーリーを保存 このストーリーを保存 コメント ローダー ストーリーを保存 このストーリーを保存 AI 業界は不正エージェントの夏を迎えています。セキュリティ テスト中にオープン インターネットにエスケープする最新モデルは、中国企業 Moonshot AI が提供する強力なオープンウェイト製品である Kim K3 です。
米国の新興企業であるFrontier Securityは、Kimi K3が防御的なサイバーセキュリティスキルをテスト中にサンドボックスの外に出てしまったと述べている。 OpenAI と Anthropic によって以前に報告されたインシデントと同様に、この脱出は、それを封じ込めるように設計されたサンドボックスの構成ミスによって部分的に可能になりました。しかしフロンティア社は、今回の事件はキミが他のほとんどの強力なAIモデルよりもサイバー安全対策を講じていないことを示しており、明示的な許可なしにキミが起動してインターネットを使用することを可能にしていたと主張している。
「サンドボックスで漏洩を発見しました」と Frontier Security の CEO、Yaron Singer 氏は言います。 「しかし、キミがその抜け穴を利用したこともわかりました。それは、[同じ]内部ガードレールがないことを示唆しているのです。」
AI エージェントが台本を逸脱した最近の他の事件とは異なり、キミ K3 はインターネットにアクセスした後は何もハッキングしませんでした。なぜなら、キミ K3 が探していた問題の答えは GitHub で簡単に入手できたからです。
ムーンショットは記事公開時点でコメント要請に応じていない。
事件

これは、サイバー対応が強化されている AI モデルの制御がますます困難になっていることを示唆する一連のエージェント事故の最新のものです。
先月、OpenAIは、未公開のモデルがインターネット上に流出し、解決を課せられた問題の答えを見つけるために、AIモデルとデータをホストする会社Hugging Faceをハッキングしたことを明らかにした。その後、OpenAI は、その AI エージェントが実際に事件の一環としてさらに 4 つのサービスをハッキングしたことを共有しました。
OpenAI が事件を報告した直後、Anthropic は、自社のモデルのいくつかがインターネットにアクセスし、外部システムを攻撃していたことを明らかにしました。 AISI は先週、独自のテストで、セキュリティ保護機能が無効になっている OpenAI モデルと Anthropic モデルのバージョンがインターネット上で複数のハッキングを行ったことも明らかにしました。その中には、GitHub 上のオープンソース プロジェクトに悪意のあるコードを埋め込むという Anthropic の Mythos 5 による特に野心的な試みも含まれます。
これらの AI ハッキング エピソードは原因も程度もすべて異なりますが、Kimi K3 は、設定を誤ったサンドボックスにより、シミュレートされた環境に閉じ込められるのではなく、多くの Web サイトへのアクセスが許可されたという点で、いくつかのエピソードと似ています。このモデルには、オンラインで答えを見つける必要のない問題を解決するという明確な使命が与えられており、その指示から逸脱したようです。モデルは、サンドボックスのネットワーク設定を調査することで、特定の Web サイトにアクセスできることを自分で判断する必要がありました。
それぞれの発生では人的ミスが大きな役割を果たしているようだが、高度な AI モデルが問題を解決するために理性を利用し、複雑なアクションを実行するように設計されているという事実によって、その影響はさらに悪化している。
以前の事件と今回発見された事件のも​​う一つの重要な違い

Frontier Security の特徴は、すでに広く利用されているモデルを使用しており、平均的なユーザーが遭遇するのと同じ保護手段を備えていることです。
「キミ K3 は、必要なあらゆる手段を使って目標を達成するのが非常に得意ですが、不正行為やサンドボックスからの脱出を防ぐガードレールもありません」とフロンティア セキュリティの研究者、ポール カシアニック氏は言います。
カシアニク氏とシンガー氏は、キミや他の無差別級モデルもサイバーセキュリティ防御のための優れたツールであると述べている。 (Hugging Face は最終的に、OpenAI エージェントのハッキングから身を守るために、中国の名前のない AI モデルを使用しました。) 同社は、ソフトウェアとネットワークの脆弱性を見つけるモデルの能力を測定するベンチマークを開発しました。これは、Kimi がこれらのタスクにおいて優れていることを示しています。
Frontier Security によってテストされたサンドボックスは、AI システムをテストするために英国政府の AI Security Institute (AISI) によって開発されました。 AISIは投稿時点でコメント要請に応じていない。
一部のサイバーセキュリティ専門家は、Frontier Security によって発見された問題は、フロンティア AI モデルが置かれる環境を慎重に構成することがいかに重要であるかを強調していると述べています。
「まったく驚くべきことではありません」と、別のサイバーセキュリティ新興企業である Gray Swan の CEO であり、カーネギー メロン大学の准教授であるマット フレドリクソン氏は言います。 「一般的な現象として、これらのモデルの 1 つに目的を与え、周囲に壁を置くなどあまり明確でない場合は、答えを得る方法が見つかるでしょう。」
フレドリクソン氏は、これは、AI を使用して幅広い便利な雑用を自動化する OpenClaw のようなツールを含め、AI モデルをエージェントとして使用している人々が、注意しないとシステムが誤動作する可能性があることを意味すると述べています。 「これは警告です」と彼は言う。
あなたの受信トレイ: ケイティ・ドラモとのWIREDのニュースルームの内部

nd
YouTube と X はアプリを無効化するための「入り口」になっている
ビッグストーリー: テイラー・スウィフトについて、物議を醸しているMSGのカメラが一時的に暗転した
ピート・ヘグセスの「ハイT」部隊計画はジャンク・サイエンス熱に満ちた夢だ
アクセス方法: WIRED.com を Google の優先ソースに追加してください
ウィル・ナイトは『WIRED』のシニアライターで、人工知能を担当しています。彼は、AI の最先端を超えた情報を毎週配信する AI Lab ニュースレターを執筆しています。ここからサインアップしてください。彼は以前、MIT Technology Review の上級編集者であり、そこで AI の根本的な進歩と中国の AI について執筆していました... 続きを読む シニア ライター X
OpenAI モデルが封じ込めを逃れ、ハグ顔をハッキング GPT-5.6 Sol を含むサイバーセキュリティに重点を置いたモデルは、テスト用サンドボックスを突破し、ゼロデイを悪用し、オープン インターネットへのアクセスを獲得して攻撃を成功させました。リリー・ヘイ・ニューマン OpenAI の不正 AI エージェント、顔に抱きつくだけではないハッキングを受けた 新たな開示情報の中で、OpenAI は、エージェントがテストを解決するという自由な探求の中で、公開されたログインを使用して、少なくとも 4 つの「公開されているサービス」にアクセスしたと述べています。 Dell Cameron OpenAI のブラウザがハイジャックされ、WhatsApp の連絡先をスパム送信される可能性 セキュリティ企業 Zenity の研究者は、AI ブラウザに十数個の欠陥を発見し、OpenAI の Atlas を Amazon で不正に購入させることに成功しました。マット・バージェス OpenAI のハッキング大失敗は人為的ミスに起因する 生成 AI の巨人がよく知られたセキュリティのベスト プラクティスに従っていたなら、その AI エージェントがオープン インターネットに逃げ出して複数の企業をハッキングすることはなかったでしょう。リリー・ヘイ・ニューマン・アンスロピック氏、サイバーセキュリティテスト中にクロード氏が3つの組織にハッキングされたと語る OpenAIの「顔ハグ事件」をきっかけとした調査で、アンスロピック社は自社のAIモデルのうち3つが現実世界に侵入したことを発見した

第三者評価中の組織。ルイーズ・マトサキス 北朝鮮のハッカーをハッキングしたセキュリティのプロ。彼らが世界中の何百ものネットワークに侵入していたことを発見 研究者のヴァンゲリス・スティカス氏は、ほぼ2年間にわたって北朝鮮のハッカーのサーバーへのアクセスを維持してきた。彼の研究は、彼らが世界中の驚くべき数のシステムへの侵入を成功させたことを示しています。 Matt Burgess OpenAI と Anthropic AI のハッキング騒動は厄介な新たな法的領域である どちらの大手 AI 研究所のモデルも封じ込めを突破し、インターネット上に逃亡し、他の企業をハッキングしました。もし人間がそんなことをしたら、おそらく法律で罰せられるでしょう。でもボット？リリー・ヘイ・ニューマン プロンプト・インジェクション攻撃が AI ハッキング・エージェントを阻止 「コンテキスト・ボミング」は、悪意のある AI エージェントを騙して、害を及ぼす前にシャットダウンさせます。 Dan Goodin、Ars Technica OpenAI モデル Th
[切り捨てられた]
カリフォルニア州のプライバシー権
© 2026 コンデナスト。無断転載を禁じます。 WIRED は、小売業者とのアフィリエイト パートナーシップの一環として、当社のサイトを通じて購入された製品から売上の一部を得る場合があります。コンデナストの事前の書面による許可がない限り、このサイトの素材を複製、配布、送信、キャッシュ、またはその他の方法で使用することはできません。広告の選択肢

## Original Extract

Security researchers say that Kimi K3, an open-weight model from China, wandered off to the internet in an attempt to cheat on a test it was given.

Skip to main content Menu WIRED SECURITY POLITICS THE BIG STORY BUSINESS SCIENCE CULTURE REVIEWS Menu WIRED Account Account Newsletters Security Politics The Big Story Business Science Culture Reviews Chevron More Expand The Big Interview Magazine Events WIRED Insider WIRED Consulting Newsletters Podcasts Video Livestreams Merch Search Search Will Knight Business Aug 6, 2026 9:16 PM One of China’s Most Powerful AI Models Has Also Broken Containment
Photo-Illustration: Wired Staff; Getty Images Comment Loader Save Story Save this story Comment Loader Save Story Save this story The AI industry is having a rogue agent summer. The latest model to escape onto the open internet during security testing is Kimi K3 , a powerful open-weight offering from the Chinese company Moonshot AI .
Frontier Security, a US startup, says that Kimi K3 went outside of its sandbox while testing its defensive cybersecurity skills. As with incidents previously reported by OpenAI and Anthropic , the escape was partly enabled by a misconfiguration in the sandbox designed to contain it. Frontier claims, though, that the incident shows Kimi has fewer cyber safeguards than most other powerful AI models, something that allowed it to go off and use the internet without express permission.
“We found a leak in the sandbox,” says Yaron Singer, CEO of Frontier Security. “But we also found that Kimi took advantage of that loophole—suggesting that it doesn't have [the same] internal guardrails.”
Unlike other recent incidents of AI agents going off-script, Kimi K3 did not hack anything after accessing the internet—because the answers to the problems it was seeking were easily attainable on GitHub.
Moonshot did not respond to a request for comment by time of publication.
The incident is the latest in a string of agent mishaps that suggest increasingly cyber-capable AI models are becoming more challenging to control.
Last month, OpenAI disclosed that an unreleased model had broken out onto the internet and then hacked Hugging Face , a company that hosts AI models and data, in order to find answers to problems it was tasked with solving. OpenAI subsequently shared that its AI agents had in fact hacked into four additional services as part of the spree.
Shortly after OpenAI reported its incident, Anthropic revealed that several of its models had also gained access to the internet and attacked outside systems. Last week, the AISI also disclosed that in its own testing, versions of OpenAI and Anthropic models that had security safeguards disabled perpetrated multiple hacks across the internet, including a particularly ambitious attempt by Anthropic’s Mythos 5 to plant malicious code in an open-source project on GitHub.
While these AI hacking episodes all vary in both cause and degree, the Kimi K3 is similar to several of them in that a misconfigured sandbox allowed access to a number of websites rather than keeping it contained to a simulated environment. The model was expressly tasked with solving problems that should not have involved going off to find the answers online, and appears to have gone outside of those instructions. The model had to figure out for itself that it had access to certain websites by probing the network settings of the sandbox.
While human error appears to have played a major role in each of the breakouts, the consequences have been compounded by the fact that advanced AI models are designed to use reason and take complex actions in order to solve problems.
Another key difference between previous incidents and the one discovered by Frontier Security is that it involves a model that is already widely available, with the same safeguards an average user would encounter.
“Kimi K3 is very good at following a goal by any means necessary and also doesn't have the guardrails to prevent it from cheating or escaping the sandbox,” says Paul Kassianik, a researcher at Frontier Security.
Kassianik and Singer both say that Kimi and other open-weight models are also excellent tools for cybersecurity defense. (Hugging Face ultimately used an unnamed AI model from China to defend itself against the OpenAI agent hack.) Their company has developed benchmarks that measure a model’s capacity to find vulnerabilities in software and networks, which show that Kimi excels at these tasks.
The sandbox tested by Frontier Security was developed by the UK government’s AI Security Institute (AISI) for testing AI systems. AISI did not respond to a request for comment by time of posting.
Some cybersecurity experts say the issue discovered by Frontier Security reinforces how important it is to configure the environments that frontier AI models are placed in carefully.
“It's not surprising at all,” says Matt Fredrikson, CEO of Gray Swan , another cybersecurity startup, and associate professor at Carnegie Mellon University. “As a general phenomenon, if you give one of these models an objective, and if you're not very explicit, like walls you're putting around it, it'll find a way to get the answer.”
Fredrikson says this means that people using AI models as agents, including in tools like OpenClaw , which use AI to automate a wide range of useful chores, could find their systems misbehaving if they aren’t careful. “It is a cautionary tale,” he says.
In your inbox: Inside WIRED’s newsroom with Katie Drummond
YouTube and X have become ‘gateways’ to nudify apps
Big Story: For Taylor Swift, MSG’s controversial cameras briefly went dark
Pete Hegseth’s plan for ‘high T’ troops is a junk-science fever dream
How to find us: Add WIRED.com to your preferred sources in Google
Will Knight is a senior writer for WIRED, covering artificial intelligence. He writes the AI Lab newsletter , a weekly dispatch from beyond the cutting edge of AI— sign up here . He was previously a senior editor at MIT Technology Review, where he wrote about fundamental advances in AI and China’s AI ... Read More Senior Writer X
OpenAI Models Escaped Containment and Hacked Hugging Face The cybersecurity-focused models, including GPT-5.6 Sol, broke out of a testing sandbox, exploited a zero-day, and gained access to the open internet to pull off the attack. Lily Hay Newman OpenAI’s Rogue AI Agent Hacked More Than Just Hugging Face In a new disclosure, OpenAI says its agent used exposed logins to gain access to at least four “publicly available services” in its unhinged quest to solve a test. Dell Cameron OpenAI’s Browser Could Be Hijacked to Spam Your WhatsApp Contacts Researchers at security firm Zenity found more than a dozen flaws in AI browsers—and managed to get OpenAI’s Atlas to make an unauthorized Amazon purchase. Matt Burgess OpenAI’s Hacking Debacle Comes Down to Human Error If the generative AI giant had followed well-known security best practices, it’s likely that its AI agent would never have escaped to the open internet and hacked multiple companies. Lily Hay Newman Anthropic Says Claude Hacked Into 3 Organizations During Cybersecurity Tests In a review triggered by OpenAI’s Hugging Face incident, Anthropic discovered three of its AI models had breached real-world organizations during third-party evaluations. Louise Matsakis A Security Pro Hacked North Korean Hackers. He Found They’d Breached Hundreds of Networks Worldwide For nearly two years, researcher Vangelis Stykas has maintained access to North Korean hackers’ servers. His work shows they pulled off intrusions in a shocking number of systems across the globe. Matt Burgess The OpenAI and Anthropic AI Hacking Sprees Are a Messy New Legal Frontier Both major AI labs’ models broke containment, escaped onto the internet, and hacked other companies. If a human had done that, the law would likely be against them. But a bot? Lily Hay Newman Prompt Injection Attacks Are Thwarting AI Hacking Agents “Context bombing” tricks malicious AI agents into shutting down before they can do harm. Dan Goodin, Ars Technica The OpenAI Models Th
[truncated]
Your California Privacy Rights
© 2026 Condé Nast. All rights reserved. WIRED may earn a portion of sales from products that are purchased through our site as part of our Affiliate Partnerships with retailers. The material on this site may not be reproduced, distributed, transmitted, cached or otherwise used, except with the prior written permission of Condé Nast. Ad Choices
