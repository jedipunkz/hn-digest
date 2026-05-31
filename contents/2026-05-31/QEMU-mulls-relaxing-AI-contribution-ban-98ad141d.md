---
source: "https://www.theregister.com/ai-ml/2026/05/29/qemu-mulls-relaxing-ai-contribution-ban/5248638"
hn_url: "https://news.ycombinator.com/item?id=48335784"
title: "QEMU mulls relaxing AI contribution ban"
article_title: "QEMU mulls relaxing AI contribution ban"
author: "Bender"
captured_at: "2026-05-31T00:30:32Z"
capture_tool: "hn-digest"
hn_id: 48335784
score: 1
comments: 0
posted_at: "2026-05-30T13:11:54Z"
tags:
  - hacker-news
  - translated
---

# QEMU mulls relaxing AI contribution ban

- HN: [48335784](https://news.ycombinator.com/item?id=48335784)
- Source: [www.theregister.com](https://www.theregister.com/ai-ml/2026/05/29/qemu-mulls-relaxing-ai-contribution-ban/5248638)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T13:11:54Z

## Translation

タイトル: QEMU、AI 貢献禁止の緩和を検討
説明: Red Hat エンジニアは、リスクのバランスが変化したと考えていますが、コアコードは立ち入り禁止のままです

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
QEMU、AI貢献禁止の緩和を検討
Red Hat エンジニアは、リスクのバランスが変化したと考えているが、コアコードは立ち入り禁止のままである
主要な Linux 仮想化コンポーネントである QEMU は、ボットからの限定的な支援を許可するために、AI によって生成されたコントリビューションの全面的な禁止を緩和することを検討しています。
この提案は、Red Hat の著名なエンジニアであり、KVM ハイパーバイザーのメンテナーである Paolo Bonzini 氏から来ました。ボンジーニ氏の提案は、「著作権侵害の影響が少なくとも元に戻すのが容易で、広がる可能性が低い場合」にAI支援を許可するというものだ。コアコードは「メンテナからの事前の同意がない限り」立ち入り禁止のままとなる。
QEMU の現在のコード来歴ポリシーでは、AI が生成したコンテンツを含む、またはそこから派生する可能性のあるものはすべて拒否されます。 Bonzini 氏は、「全面的な禁止は維持が容易でしたが、LLM 出力が単独で使用できることはほとんどありませんでしたが、ツールが改良されるにつれて、絶対的な禁止を正当化することが難しくなりました。」と書いています。
ボットやプロジェクト名は話題にならず、AI が作成したオープンソース ソフトウェアを恥じる
Git は AI コーディングの津波に備えることができていない
AI の目でバグをスキャンすることで、Linux のセキュリティに関する憂慮すべき傾向が生じる
FreeBSD プロジェクトはまだ AI にコードをコミットさせる準備ができていません
AI アシスタントからのコードの問題はそのソースにあります。提出者にはコードを提供する法的権利があるのでしょうか?ボンジーニ氏の見解は、著作権とライセンスに関する懸念は依然としてあるものの、「変化したのはリスクのバランスだ」というものだ。
リスクはどれくらいですか?ボンジーニによれば、それはそうではなかったという。このエンジニアは、深刻な法的問題に陥ることなく AI コンテンツを受け入れた他のプロジェクトや、リスクを考慮した組織 (Red Hat を含む) を挙げた。

kは許容範囲でした。
そうは言っても、Red Hat には自由に使える弁護士軍団がいますが、QEMU のようなプロジェクトには同じリソースがありません。そのため、AI 支援コードを取り消しできる領域 (Bonzini は小さなバグ修正やドキュメントなどの例を示しました) に保持することが提案されています。
貢献における LLM 出力の使用は議論の余地があり、ファンと中傷者がいます。 OpenSlopware などのプロジェクトは、LLM で生成されたコードや統合 AI テクノロジーを使用するフリー ソフトウェアやオープン ソース プロジェクトを追跡しました。懸念事項の 1 つは、LLM がどのようなトレーニングを受けているか、また、このテクノロジーによって生成されたコードのチャンクにライセンス上の問題が発生する可能性があるというリスクです。
解決策の 1 つは、投稿で AI の使用を開示することですが、使用が些細な場合には必要ない可能性があります (Red Hat は変数名の自動補完の例を示しました)。
Bonzini 氏はまた、「AI が使用された場所を記録するトレーラーとして『AI-used-for:』を導入し、レビュー担当者が結果を判断するのに役立つ他の提案を含める」ことも提案しました。
「この基準は、作成者がポリシーを読んだことのチェックを兼ねた、より一般的な『Assisted-by』とは少し異なります。」
Bonzini 氏は、「AI の使用は他の貢献要件を緩和するものではない」と述べましたが、この議論は、AI 支援の全面的な禁止が前進する方法ではない可能性があり、より微妙なアプローチが必要であるという認識を示しています。 ®
ソフトウェア
ウィキメディアの解雇後、ウィキペディア編集者がストライキとバナー破壊行為を計画
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
国立宇宙センターでのロケット展示が NASA SLS の意図せぬ印象を引き出す
Postgres における AI とデータ主権: データセンターのエネルギー危機への答え
10億人のAIエージェントが電力網に乗り込む
AWS は、

企業需要がゼロにもかかわらず、イーロン・マスクのGrokをBedrockに押し込む
フロンティアモデルのエナジードリンク
システム
EUのデジタル主権ブーブーは、このプロジェクトに起こった最高の出来事かもしれない
DIY か、死ぬか。 CIAに買わせないでください
単独の攻撃者が、人気のある OpenSearch、Elasticsearch ライブラリを模倣した 14 個の悪意のある npm パッケージを公開
そしてマイクロソフトはそれらすべてを潰した
AI + ML
Googleは最近AIの機能化に本格的に取り組んでいる
セキュリティ
Anthropic、Mythos クラスのモデルを一般公開へ
セキュリティ
レドモンドが警察に通報、マイクロソフトに「屈辱を受けた」と不満を抱く0日ハンターが「骨が砕けるほどの衝撃」を誓う
オペレーティング システム
リーナス・トーバルズ氏、「無意味なプルリクエスト」について「より厳格になり始める」 - そのうちのいくつかはAIからのもの
セキュリティ
軍隊の携帯電話が位置データを外国の敵に提供した
データ主権のトレードオフを克服する
避けられないトレードオフであるデータ主権はネットワークにとって実際に何を意味するのでしょうか?もっと詳しく知る。
プロンプトからエクスプロイトまで: LLM は API 攻撃をどのように変化させているか
最新のアプリケーションは API 主導で相互接続されており、多くの場合、過剰な権限が与えられているため、AI 支援攻撃の理想的な標的となっています。
未来の構築: Kubernetes 向けエンタープライズ データ サービスのロックを解除する
インフラストラクチャのサイロを排除し、標準化されたエンタープライズ グレードのクラウド ネイティブ プラットフォームを確立する方法を発見してください。
Microsoft 365 が見逃す高度な攻撃を Behavioral AI セキュリティで捕捉
Microsoft 365 は企業コミュニケーションのバックボーンであり、そのネイティブ セキュリティが既知のものや騒々しいものを除外します。
ランサムウェア侵害の混乱に足を踏み入れ、対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んでサイバー犯罪者を出し抜こう
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えず、

私たちもそうではありません。 Druva の人気イベント「Escape Ransomware」が完全にバーチャルになりました。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで真の ROI を実現する方法を学びましょう。
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
フロンティアモデルのエナジードリンク
そしてマイクロソフトはそれらすべてを潰した
CEOのトッド・マッキノン氏は、ServiceNowを含む顧客がオフスイッチを望んでいると語る
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
フロンティアモデルのエナジードリンク
そしてマイクロソフトはそれらすべてを潰した
CEOのトッド・マッキノン氏は、ServiceNowを含む顧客がオフスイッチを望んでいると語る
ウィキメディアの解雇後、ウィキペディア編集者がストライキとバナー破壊行為を計画
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
国立宇宙センターでのロケット展示が NASA SLS の意図せぬ印象を引き出す
5、4、3、2、1... pfft
AWSはエンタープライズ需要がゼロにもかかわらず、イーロン・マスク氏のGrokをBedrockに押し込むと報じられている
フロンティアモデルのエナジードリンク
単独の攻撃者が、人気のある OpenSearch、Elasticsearch ライブラリを模倣した 14 個の悪意のある npm パッケージを公開
そしてマイクロソフトはそれらすべてを潰した
Okta は不正な AI エージェントを殺すための独自のライセンスを作成します
CEOのトッド・マッキノン氏は、ServiceNowを含む顧客がオフスイッチを望んでいると語る
ICE、2,500万ドルの生体認証スキャナ契約であなたの目を監視
顔認識アプリは邪魔だと思いましたか?
欧州は米国の支配から逃れるためにソブリンクラウドを構築した。それからプロセッサーのことは忘れました
Canvasをハッキングした「犯罪者と卑劣者」が盗まれた学生データを本当に削除したとは誰も信じていない
欧州は米国のテクノロジー支配下から抜け出したいが、まずは出口を見つけなければならない
GNOMEがUbuntuを支配するかもしれない

毅然としたアライグマ、しかし X.org はまだロードキルではない
OpenClaw、ただしコンテナ内: NanoClaw の紹介
オープンソース レジストリには、基本的なセキュリティを実装するための十分な資金がありません
Linux Windows 内に Windows アプリを含める
Linux 中年の危機は Tux 主導の変革のチャンス
AI が多すぎる人もいれば、少なすぎる人もいる: AMD が投資家に勝てない理由
エージェント AI が現代のメモリ階層に負担をかける仕組み
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
状況出版
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Red Hat engineer reckons the balance of risk has shifted, but core code stays off limits

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
QEMU mulls relaxing AI contribution ban
Red Hat engineer reckons the balance of risk has shifted, but core code stays off limits
A key Linux virtualization component, QEMU, is considering relaxing its blanket ban on AI-generated contributions to allow limited assistance from the bots.
The suggestion came from Paolo Bonzini, distinguished engineer at Red Hat and a maintainer of the KVM hypervisor. Bonzini's suggestion is to allow AI assistance "where the ramifications of copyright violations are at least easy to revert and unlikely to spread." Core code would remain off-limits "without prior agreement from a maintainer."
QEMU's current code provenance policy rejects anything that might include or derive from AI-generated content. "A blanket ban," wrote Bonzini, "was easy to maintain while LLM output was rarely usable on its own, but as the tools improved an absolute prohibition has become harder to justify."
Not hot on bots, project names and shames AI-created open source software
Git is unprepared for the AI coding tsunami
AI eyes scanning for bugs create a worrisome Linux security trend
FreeBSD Project isn't ready to let AI commit code just yet
The problem with code from AI assistants is its source – does the submitter have the legal right to contribute the code? Bonzini's take is that while there remain concerns around copyright and licensing, "what has shifted is the balance of risk."
How big is the risk? Not what it was, according to Bonzini. The engineer cited other projects that had accepted AI content without running into serious legal trouble, and organizations (including Red Hat) that reckoned the risk was acceptable.
That said, while Red Hat has an army of lawyers at its disposal, a project such as QEMU doesn't have the same resources, hence the suggestion to keep AI-assisted code in areas (Bonzini gave examples, including small bug fixes and documentation) where it can be backed out.
The use of LLM output in contributions is a contentious one and has its fans and detractors. Projects such as OpenSlopware tracked free software and open source projects that used LLM-generated code or integrated AI technologies. One concern cited is what LLMs have been trained on and the risk that chunks of code produced by the technology might have licensing issues.
One solution is to disclose the use of AI in a contribution, although this might not be necessary where the use is trivial (Red Hat gave the example of autocompleting a variable name.)
Bonzini also suggested, "Introduce 'AI-used-for:' as a trailer to record where AI was used, and include other suggestions that help reviewers judge the result."
"The standard is slightly different from the more usual 'Assisted-by', which doubles as a check that the author has read the policy."
Although Bonzini noted, "use of AI does not relax any other contribution requirement," the discussion indicates a recognition that blanket bans on AI assistance might not be the way forward and that a more nuanced approach is needed. ®
Software
Wikipedia editors plot strike and banner sabotage after Wikimedia layoffs
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
Rocket exhibit at National Space Centre pulls off unintentional NASA SLS impression
AI and data sovereignty in Postgres: An answer to the datacenter energy crisis
A billion AI agents walk into a power grid
AWS reportedly to tuck Elon Musk's Grok into Bedrock, despite zero enterprise demand
The energy drink of frontier models
Systems
EU's digital sovereignty boo-boo may be the best thing to ever happen to the project
DIY or die. Just don't let the CIA buy it
Lone attacker published 14 malicious npm packages mimicking popular OpenSearch, Elasticsearch libraries
And then Microsoft busted them all
AI + ML
Google has seriously leaned into AI enshittification lately
Security
Anthropic to release Mythos-class models to the public
Security
Disgruntled 0-day hunter 'humiliated' by Microsoft pledges 'bone shattering drop' as Redmond calls cops
Operating Systems
Linus Torvalds to ‘start being more hardnosed’ about ‘pointless pull requests’ – some of which come from AIs
Security
Troops’ phones gave away location data to foreign adversaries
Overcoming the trade-offs in data sovereignty
What does data sovereignty actually mean for your network, which trade-offs are unavoidable? Learn more.
From Prompt to Exploit: How LLMs Are Changing API Attacks
Modern applications are API-driven, interconnected, and often over-permissioned, making them an ideal target for AI-assisted attacks.
Architecting the Future: Unlocking Enterprise Data Services for Kubernetes
Join us to discover how to eliminate infrastructure silos and establish a standardized, enterprise-grade cloud-native platform.
Catch the Advanced Attacks Microsoft 365 Misses with Behavioral AI Security
Microsoft 365 is the backbone of enterprise communication, and its native security filters out the known and the noisy.
Step into the chaos of a live ransomware breach, test your response skills, and team up with other IT and security pros to outsmart cybercriminals
Virtual Cyber Recovery Simulation
Ransomware attacks aren’t slowing down, and neither are we. Druva’s hit event, Escape Ransomware, is now fully virtual.
Agentic AI at Scale: From Pilot to Production
Join us to learn how to unlock real ROI by driving adoption of AI at scale.
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
The energy drink of frontier models
And then Microsoft busted them all
CEO Todd McKinnon says customers including ServiceNow want an off switch
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
The energy drink of frontier models
And then Microsoft busted them all
CEO Todd McKinnon says customers including ServiceNow want an off switch
Wikipedia editors plot strike and banner sabotage after Wikimedia layoffs
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
Rocket exhibit at National Space Centre pulls off unintentional NASA SLS impression
5, 4, 3, 2, 1... pfft
AWS reportedly to tuck Elon Musk's Grok into Bedrock, despite zero enterprise demand
The energy drink of frontier models
Lone attacker published 14 malicious npm packages mimicking popular OpenSearch, Elasticsearch libraries
And then Microsoft busted them all
Okta writes its own license to kill rogue AI agents
CEO Todd McKinnon says customers including ServiceNow want an off switch
ICE to keep an eye on your eyes under $25M biometric scanner deal
And you thought a face recognition app was intrusive?
Europe built sovereign clouds to escape US control. Then forgot about the processors
Nobody believes the 'criminals and scumbags' who hacked Canvas really deleted stolen student data
Europe wants out from under US tech – but first it has to find the exits
GNOME may rule Ubuntu Resolute Raccoon, but X.org isn't roadkill yet
OpenClaw, but in containers: Meet NanoClaw
Open source registries don't have enough money to implement basic security
Contain your Windows apps inside Linux Windows
The Linux mid-life crisis that's an opportunity for Tux-led transformation
Too much AI for some, too little for others: Why AMD can't win with investors
How agentic AI can strain modern memory hierarchies
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
