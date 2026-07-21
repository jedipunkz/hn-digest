---
source: "https://siliconangle.com/2026/07/20/hugging-face-uses-open-weights-z-ai-glm-5-2-defend-attacker-commercial-frontier-model-refusal/"
hn_url: "https://news.ycombinator.com/item?id=48998910"
title: "Hugging Face uses open-weights Z.ai GLM 5.2 to battle attacker"
article_title: "Hugging Face uses open-weights Z.ai GLM 5.2 to battle attacker after commercial frontier model refusal - SiliconANGLE"
author: "vednig"
captured_at: "2026-07-21T22:45:19Z"
capture_tool: "hn-digest"
hn_id: 48998910
score: 3
comments: 1
posted_at: "2026-07-21T21:57:01Z"
tags:
  - hacker-news
  - translated
---

# Hugging Face uses open-weights Z.ai GLM 5.2 to battle attacker

- HN: [48998910](https://news.ycombinator.com/item?id=48998910)
- Source: [siliconangle.com](https://siliconangle.com/2026/07/20/hugging-face-uses-open-weights-z-ai-glm-5-2-defend-attacker-commercial-frontier-model-refusal/)
- Score: 3
- Comments: 1
- Posted: 2026-07-21T21:57:01Z

## Translation

タイトル: ハギング・フェイスは無差別ウェイト Z.ai GLM 5.2 を使用してアタッカーと戦う
記事タイトル: ハグフェイスは商用フロンティアモデル拒否後の攻撃者と戦うために無差別ウェイト Z.ai GLM 5.2 を使用 - SiliconANGLE
説明: ハグフェイスは商用フロンティアモデル拒否後の攻撃者と戦うためにオープンウェイト Z.ai GLM 5.2 を使用 - SiliconANGLE

記事本文:
古いブラウザを使用しています。エクスペリエンスを向上させるためにブラウザをアップグレードしてください。
シリコンアングル100
マーケットデータ提供：
2026年7月20日12:05 EDT更新
ハグ・フェイスは商用フロンティアモデルの拒否を受けて攻撃者と戦うためにオープンウェイトのZ.ai GLM 5.2を使用
「機械学習の GitHub」とよく言われるオープンソースの人工知能プラットフォームである Hugging Face Inc. は、商用 AI モデルの安全ガードレールがリクエストをブロックしたため、エージェント AI 攻撃に対応するためにオープンウェイト モデルを使用せざるを得なくなったことがわかりました。
先週、Hugging Face は、自律型 AI エージェント システムを使用して、限られた内部データセットと内部サービスで使用されるいくつかの認証情報にアクセスする攻撃者による侵害を検出したと発表しました。同社は防御的に対応し、攻撃者を遮断し、システムを強化しました。
分析の一環として、Hugging Face はフロンティア AI モデルを利用してログ分析を支援しました。これは、Hugging Face が多数の AI モデルの主要なアクセス ポイントであることを考えると、完全に賢明なオプションです。
しかし、それは失敗しました。この種の分析では、膨大な量の実際の攻撃データとコマンド、エクスプロイト ペイロード、攻撃アーティファクトを送信する必要があります。これらのリクエストは、商用フロンティア AI モデルによってブロックされました。そのガードレールは、攻撃者のためのエクスプロイトの構築を要求されたものと、それを検出しようとする防御者とを区別できないためです。
当初、迅速な分析とスピードの必要性があったため、同社は、約 7,530 億のパラメーターを備えた強力なオープンウェイト モデルである Z.ai Co. Ltd. の GLM 5.2 に切り替えました。サードパーティ モデルとは異なり、ローカルまたはクラウド ハードウェア上で、企業の保護されたファイアウォール境界内で完全に実行できます。つまり、制御されたインフラストラクチャからデータが流出することはありません。
GLM 5.2やBeijiなどの中国製AIモデル

ng Moonshot AI Technology Co. Ltd. の Kimi K3 は、オープンソースおよびオープンウェイト モデルが、米国企業が開発した現在の先駆者モデルや主力フロンティア モデルに匹敵するか、それに匹敵することさえできることを証明しました。 GLM 5.2 は、高度な推論、コーディング、さらには脆弱なコードやエクスプロイトの発見が可能な Mythos クラスのモデルである Anthropic PBC の Fable 5 の機能に到達します。また、Anthropic よりもはるかに安価な推論コストで実現します。
ただし、これらのモデルの悪用を防ぐために、Anthropic やその他の主要なクローズドソース AI 開発者は、誤用を避けるために誤検知を引き起こす強力な安全ガードレールを設置しました。これにより、有効なサイバーセキュリティの役割において実際に使用する能力が全体的に大幅に低下します。 Anthropic は、フロンティア モデルを研究者がより安全に使用できるようにするために、偽陽性率が調整されていると述べています。
Anthropic や OpenAI PBC Group などの企業は、自社の最も強力なモデルに対して自発的にこれらの制限を設けていますが、Mythos 5 と Fable 5 は両方とも先月、最初のリリース直後に米国政府の要請により廃止されました。同様に、米国政府は OpenAI に対し、現在一般公開されている独自のフロンティア モデル ファミリ GPT 5.6 のリリースを延期するよう要請しました。
中国のオープンソースモデルとアメリカのクローズドソースモデルの間の緊張
北京に本拠を置くムーンショットAI社の強力なニアフロンティアクラスのオープンウェイトモデルであるキミK3が最近リリースされたことにより、トランプ政権が米国企業による中国製オープンモデルの使用を禁止する可能性があるとの報道が高まっている。
Axios によると、政権の一部はこれまでにも外国のオープンソース モデルを事実上禁止しようと試みているという。これは、多くの米国企業が中国のオープンウェイトMOをますます使用しているときに行われます。

dels は導入コストが安く、競合する商用専用モデルに近いためです。
米国は中国モデルの使用を全面的に禁止する必要はない。その代わりに、政府はセキュリティとガバナンスの欠如を理由に、圧力キャンペーンを利用して企業を企業から遠ざけることができるだろう。関係者によると、米国商務省は昨年、複数の中国のAI研究所を「エンティティリスト」に追加することも検討しており、これにより適切なライセンスのない企業へのアクセスが事実上遮断されることになる。
ホワイトハウスのAIおよび暗号通貨顧問のデイビッド・サックス氏は日曜日、X（以前はTwitter）に「私たちはAI政策の重要な転換点にある。AIモデルの収益という点ですでに複占状態にある大手閉鎖研究所は、政府がオープンソースの競争を排除することを望んでいる」と書いた。
2部構成のXの投稿で、サックス氏は最近のハグ・フェイス事件に触れ、「中国人モデルが問題なくこなす仕事をアメリカ人モデルに制限する理由はない。我々は自分たちの競争力を弱めているだけだ」と付け加えた。
コメントの中で、同氏は、Kimi K3 が 15 件の重大なセキュリティ バグを修正したと述べ、OpenAI と Anthropic のモデルが「サイバー ガードレール」を理由に修正できなかったと指摘しました。 15 件のバグ修正は、このモデルを使用した開発者の主張であり、既知のベンチマークではありませんが、開発者にとっては、中国のモデルが米国のクローズドソース モデルよりも主流になりつつあることに注目する焦点となっています。
開発者は、すべての費用もたったの 250 ドルだと付け加えた。行われた正確な作業内容が分からない場合、コストの比較は不可能であるとしても困難ですが、キミ K3 の価格は Fable 5 の約 3 分の 1 です。もちろん、Fable 5 が作業を拒否した場合、価格の違いは問題ではありません。
中国の習近平国家主席は、世界が自国のモデルに目を向け始めていることに注目し、次のように呼びかけた。

世界的な協力。
習主席は上海で開催された中国年次世界人工知能会議の開会式で、「人工知能の開発は一国による単独のパフォーマンスであるべきではなく、世界的な協力の交響曲であるべきだ」と述べた。
同氏は演説の中で、AIに関連した「国家安全保障の概念の過度の拡大」に反対するよう呼びかけた。中国はこれまで、米国で一部の最先端AIの利用を阻止され、最先端のAIチップの使用も制限されてきた。
中国は、東南アジア諸国連合、アラブ連盟、アフリカ連合、ラテンアメリカ・カリブ海諸国共同体、上海協力機構、BRICS諸国とのAI協力を拡大する意向である。
習氏は今後5年間で、中国は発展途上国に5,000回のAI訓練の機会を提供すると述べた。同氏はまた、中国が開発した早期警戒システム用の気象ツールへのアクセスを30カ国に提供すると述べた。
SiliconANGLE の共同創設者である John Furrier からのメッセージ:
theCUBE コミュニティに参加することで、コンテンツをオープンかつ無料に保つという私たちの使命をサポートしてください。 theCUBE の Alumni Trust Network に参加してください。このネットワークでは、テクノロジー リーダーがつながり、情報を共有し、機会を創出します。
theCUBE ビデオの 1,500 万人以上の視聴者が AI、クラウド、サイバーセキュリティなどにわたる会話を強化
11,400 人以上の theCUBE 卒業生 — 信頼できる独自のネットワークを通じて、未来を形作る 11,400 人を超えるテクノロジーおよびビジネスのリーダーとつながりましょう。
あなたは AWS の顧客ですか? Marketplace ポータル ページとリンクから AWS サービスを購入して、SiliconANGLE を経済的にサポートします。
テクノロジーの先見者である John Furrier と Dave Vellante によって設立された SiliconANGLE Media は、15 マイル以上に及ぶ業界をリードするデジタル メディア ブランドのダイナミックなエコシステムを構築しました。

10 万人のエリート技術専門家。当社の新しい独自の theCUBE AI Video Cloud は、theCUBEai.com ニューラル ネットワークを活用して視聴者との対話に画期的な進歩をもたらし、テクノロジー企業がデータに基づいた意思決定を行い、業界の会話の最前線に留まるのを支援します。
無料コンテンツが好きですか?購読してフォローしてください。
Whop、ビジネスオーナーとAIエージェントがデジタルコマースにアクセスできるようにするコマンドラインインターフェースをリリース
Google、Gemini をより安価なモデルとバグハンターで拡張、常にリードを維持
Box は、エンタープライズ コンテンツを操作する AI エージェントを管理するセキュリティ制御を追加します
Nvidia は、Vera Rubin のパフォーマンスの大幅な向上を示すため、AI ファクトリーを倍増します
Harness は、開発者が使い慣れたプロセスとツールを使用して AI エージェントを展開できるようにエージェント DLC を開始します
独占: Fig が Figaro AI エージェントを起動し、CI/CD をセキュリティ運用に導入
Whop、ビジネスオーナーとAIエージェントがデジタルコマースにアクセスできるようにするコマンドラインインターフェースをリリース
アプリ - KYT DOTSON 作。 10分前
Google、Gemini をより安価なモデルとバグハンターで拡張、常にリードを維持
AI - ダンカン・ライリー著。 10分前
Box は、エンタープライズ コンテンツを操作する AI エージェントを管理するセキュリティ制御を追加します
セキュリティ - ダンカン・ライリー著。 10分前
Nvidia は、Vera Rubin のパフォーマンスの大幅な向上を示すため、AI ファクトリーを倍増します
インフラ - マイク・ウィートリー著。 10分前
Harness は、開発者が使い慣れたプロセスとツールを使用して AI エージェントを展開できるようにエージェント DLC を開始します
AI - マイク・ウィートリー著。 10分前
独占: Fig が Figaro AI エージェントを起動し、CI/CD をセキュリティ運用に導入
セキュリティ - ダンカン・ライリー著。 1時間前
2026 SiliconANGLE Media Inc. 無断複写・転載を禁じます。
提案の種類 * 消費者向けか企業向けか?エンタープライズ スタートアップ ニュース 新興テクノロジー
電子メール このフィールドは検証目的のため、変更しないでください。
元

ノイズから信号を抽出する
無料コンテンツが好きですか?購読してフォローしてください。
分析と測定のために Cookie を使用します。追跡 Cookie を受け入れるか拒否することができます。当社のプライバシーポリシーをご覧ください。

## Original Extract

Hugging Face uses open-weights Z.ai GLM 5.2 to battle attacker after commercial frontier model refusal - SiliconANGLE

You are using an outdated browser. Please upgrade your browser to improve your experience.
THE SILICONANGLE 100
Market Data Provided by
UPDATED 12:05 EDT / JULY 20 2026
Hugging Face uses open-weights Z.ai GLM 5.2 to battle attacker after commercial frontier model refusal
Hugging Face Inc., an open-source artificial intelligence platform often described as the “GitHub of machine learning,” found itself forced to use an open-weights model to respond to an agentic AI attack after the safety guardrails on commercial AI models blocked requests.
Last week, Hugging Face said it detected a breach from an attacker using an autonomous AI agent system to access a limited set of internal datasets and several credentials used by internal services. The company responded defensively, cut off the attacker and hardened the system.
As part of the analysis, Hugging Face went to frontier AI models to assist with log analysis – a completely sensible option given that Hugging Face is a premier access point for numerous AI models.
However, that failed: This kind of analysis requires sending a tremendous amount of real attack data and commands, exploit payloads and attack artifacts. These requests were blocked by commercial frontier AI models because their guardrails cannot distinguish between being asked to build exploits for an attacker and a defender trying to detect them.
With a need for rapid analysis and speed at the onset, the company switched to Z.ai Co. Ltd.’s GLM 5.2, a powerful open-weight model with about 753 billion parameters. Unlike a third-party model, it can be run entirely on local or cloud hardware and within a company’s protected firewall perimeter, meaning no data exits its controlled infrastructure.
Chinese-built AI models such as GLM 5.2 and Beijing Moonshot AI Technology Co. Ltd.’s Kimi K3 have proved that open-source and open-weight models can reach, or even rival, the current forerunner and flagship frontier models built by American companies. GLM 5.2 reaches the capabilities of Anthropic PBC’s Fable 5, a Mythos-class model capable of advanced reasoning, coding and even discovering vulnerable code and exploits. It also does so with far cheaper inference costs than delivered by Anthropic.
However, to prevent the misuse of these models, Anthropic and other leading closed-source AI developers have put strong safety guardrails in place that trigger higher false positives to avoid misuse. This also makes them far less capable overall for real-world usage in valid cybersecurity roles. Anthropic has noted that the false positive rate is being adjusted as it works to make its frontier models safer for use by researchers.
Although companies such as Anthropic and OpenAI PBC Group have voluntarily placed these restrictions on their most powerful models, Mythos 5 and Fable 5 were both pulled last month at the request of the U.S. government shortly after they first launched. Similarly, the U.S. government asked OpenAI to delay the release of its own frontier model family GPT 5.6 , which is now publicly available.
Tension between Chinese open source and American closed source models
The recent release of Kimi K3, a powerful near-frontier-class open-weight model from Beijing-based Moonshot AI, has fueled reports that the Trump administration may ban U.S. companies from using Chinese open models.
Parts of the administration have already attempted to construct de facto bans on foreign open-source models before, according to Axios . This comes at a time when many U.S. companies are increasingly using Chinese open-weight models because they are cheaper to deploy and come close to rivaling commercial-only models.
The U.S. would not need to ban the use of Chinese models outright; instead, the government could use pressure campaigns to steer corporations away from them, citing a lack of security and governance. According to sources, the U.S. Commerce Department also considered adding multiple Chinese AI labs to its “ Entity List ” last year, which would effectively cut off access to companies without proper licensing.
David Sacks, the White House AI and crypto advisor, wrote on X , formerly Twitter, on Sunday: “We are at a critical inflection point in AI policy. The leading closed labs, already a duopoly in terms of AI model revenue, want the government to eliminate their open source competition.”
In a two-part X post, noting the recent Hugging Face event, Sacks added : “There’s no reason to limit American models on tasks that Chinese models handle without issue. We’re only making ourselves less competitive.”
In his commentary, he mentioned that Kimi K3 fixed 15 critical security bugs that he noted OpenAI and Anthropic’s models refused to because of “cyber guardrails.” The 15-bug fix is a claim from a developer who used the model, not a known benchmark, but it has become a focal point for developers to note that Chinese models are becoming a mainstay over U.S. closed-source models.
The developer added that the whole affair also only cost $250. Cost comparisons would be difficult, if impossible, without knowing the exact work done, but Kimi K3 is around one-third the price of Fable 5. Of course, the price distinction doesn’t matter if Fable 5 refuses to do the work.
Noting that the world is beginning to turn toward its models, China’s President Xi Jinping called for global cooperation .
“The development of artificial intelligence should not be a solo performance by any single country but rather a symphony of global cooperation,” Xi said at the opening of China’s annual World Artificial Intelligence Conference in Shanghai.
In his speech, he called for opposing the “overstretching of the concept of national security” as it relates to AI. China has been historically blocked from using some of the most advanced AI in the U.S. and restricted from state-of-the-art AI chips.
China intends to expand AI cooperation with the Association of Southeast Asian Nations, the League of Arab States, the African Union, the Community of Latin American and Caribbean States, the Shanghai Cooperation Organization and the BRICS countries.
Over the next five years, Xi said China will provide 5,000 AI training opportunities for developing countries. He also said China will provide 30 countries access to a Chinese-developed meteorological tool for early-warning systems.
A message from John Furrier, co-founder of SiliconANGLE:
Support our mission to keep content open and free by engaging with theCUBE community. Join theCUBE’s Alumni Trust Network , where technology leaders connect, share intelligence and create opportunities.
15M+ viewers of theCUBE videos , powering conversations across AI, cloud, cybersecurity and more
11.4k+ theCUBE alumni — Connect with more than 11,400 tech and business leaders shaping the future through a unique trusted-based network.
Are you AWS customer? Support SiliconANGLE Financially by buying your AWS services from our Marketplace portal page and links.
Founded by tech visionaries John Furrier and Dave Vellante, SiliconANGLE Media has built a dynamic ecosystem of industry-leading digital media brands that reach 15+ million elite tech professionals. Our new proprietary theCUBE AI Video Cloud is breaking ground in audience interaction, leveraging theCUBEai.com neural network to help technology companies make data-driven decisions and stay at the forefront of industry conversations.
Like Free Content? Subscribe to follow.
Whop releases command-line interface to let business owners and their AI agents access digital commerce
Google expands Gemini with cheaper models and a bug-hunter it keeps on a leash
Box adds security controls to govern AI agents working with enterprise content
Nvidia doubles down on AI factories as it showcases massive Vera Rubin performance gains
Harness launches Agent DLC for developers to deploy AI agents using familiar processes and tools
Exclusive: Fig launches Figaro AI agent and brings CI/CD to security operations
Whop releases command-line interface to let business owners and their AI agents access digital commerce
APPS - BY KYT DOTSON . 10 MINS AGO
Google expands Gemini with cheaper models and a bug-hunter it keeps on a leash
AI - BY DUNCAN RILEY . 10 MINS AGO
Box adds security controls to govern AI agents working with enterprise content
SECURITY - BY DUNCAN RILEY . 10 MINS AGO
Nvidia doubles down on AI factories as it showcases massive Vera Rubin performance gains
INFRA - BY MIKE WHEATLEY . 10 MINS AGO
Harness launches Agent DLC for developers to deploy AI agents using familiar processes and tools
AI - BY MIKE WHEATLEY . 10 MINS AGO
Exclusive: Fig launches Figaro AI agent and brings CI/CD to security operations
SECURITY - BY DUNCAN RILEY . 1 HOUR AGO
2026 SiliconANGLE Media Inc. All rights reserved.
Type of Pitch * Consumer or Enterprise? Enterprise Startup News Emerging Technology
Email This field is for validation purposes and should be left unchanged.
EXTRACT THE SIGNAL FROM THE NOISE
Like Free Content? Subscribe to follow.
We use cookies for analytics and measurement. You can accept or reject tracking cookies. See our Privacy Policy .
