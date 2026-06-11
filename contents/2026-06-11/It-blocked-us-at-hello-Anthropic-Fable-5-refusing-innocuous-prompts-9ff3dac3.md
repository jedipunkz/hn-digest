---
source: "https://www.theregister.com/ai-and-ml/2026/06/10/anthropic-claude-fable-5-refuses-innocuous-prompts/5253754"
hn_url: "https://news.ycombinator.com/item?id=48486370"
title: "It blocked us at 'hello ' Anthropic Fable 5 refusing innocuous prompts"
article_title: "Anthropic Claude Fable 5 refuses innocuous prompts"
author: "abliterationai"
captured_at: "2026-06-11T07:50:55Z"
capture_tool: "hn-digest"
hn_id: 48486370
score: 18
comments: 2
posted_at: "2026-06-11T04:54:52Z"
tags:
  - hacker-news
  - translated
---

# It blocked us at 'hello ' Anthropic Fable 5 refusing innocuous prompts

- HN: [48486370](https://news.ycombinator.com/item?id=48486370)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/06/10/anthropic-claude-fable-5-refuses-innocuous-prompts/5253754)
- Score: 18
- Comments: 2
- Posted: 2026-06-11T04:54:52Z

## Translation

タイトル: 「こんにちは」Anthropic Fable 5 で無害なプロンプトを拒否してブロックされました
記事のタイトル: 人間のクロード寓話 5 は無害なプロンプトを拒否します
説明: 非常に警戒深い安全分類子が寓話を警告の物語に変える

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
「こんにちは！」でブロックされました。無害なプロンプトを拒否する人類の寓話 5
非常に警戒深い安全分類子が寓話を警告の物語に変える
更新 Anthropic が新たにリリースした Claude Fable 5 生成 AI モデルは、安全を確保しようと努めすぎて、自身のユーザーベースに損害を与えています。 AI 知識逆流ツールを使用しようとしている顧客は、モデルが無害な質問への回答を拒否していると報告しています。この問題は、過去のモデルのリリース以来、セキュリティ研究者を悩ませてきました。
Anthropic は、Fable 5 のガードレールが保守的に調整されていると警告し、「無害なリクエストをキャッチすることもありますが、平均してセッションの 5% 未満でトリガーされます」と同社は述べ、「誤検知をできるだけ早く減らす」ことを約束しました。
人間が紡ぐ、飼いならされたより安全な神話の寓話
GM、データセンター熱に沸き、グリッド規模のナトリウムイオン電池の構築を決定
データセンターの成長は 2030 年までに電力の壁に突き当たる可能性がある
macOS 27 ベータ版が Apple Silicon から Asahi Linux を起動
同社は、モデルの拒否を定量化するよう求める要請にすぐには応じなかった。したがって、実際の偽陽性率が 5% より大きいか小さいかは不明です。しかし、世界中に 1,800 万から 3,000 万人のユーザーがいると推定されており、妨害されたユーザーのうちのほんの一部でも迷惑になります。
ゲイツ財団のグローバルヘルス部門の一部である疾患モデリング研究所の主任研究員であるマイク・ファミュラーレ氏は、クロード・ファブル 5 が「こんにちは」などの入力を拒否すると報告しています ( #66657 )。
「クロード コードでは、Fable 5 の入力安全性分類器は、基本的にすべてのセッションの最初のターンで、model_refusal_fallback (Opus 4.8 へのサイレント スイッチ) を発行します。

ccount — ユーザー入力が hello! という単語だけであるセッションを含みます。 。起動時にコンテキスト内のリポジトリ コンテンツ、ツール呼び出し、およびファイル読み取りは行われません。」
不満を抱いている顧客は彼だけではありません。 Fable 5 のデビュー以来、他の多くのバグレポートが Anthropic の Claude Code GitHub リポジトリに提出されています。 [バグ] Fable 5 モデルの安全フィルターにより、無害なメッセージで誤検知が発生する #66587 ; Fable 5 は、「Application Security Architect の履歴書」編集の支援を拒否しています #66655 ; [機能リクエスト] 非研究ラボ管理システム #67062 などでの Fable 5 の使用を許可します。
社会的怒りのサイト X.com で、免疫学者でジャクソンゲノム医学研究所の教授であるデリヤ・ウヌトマズ氏は、「『がん』という言葉は、クロード・ファブル 5 によってバイオセキュリティのリスクとして警告されている!」と指摘しています。
同様の苦情が Reddit のスレッドにも現れています。
アンスロピック社がライバルのフロンティアモデル開発を阻止しようとする安全介入を隠すことを選択したため、寓話5は異例である。サイバーセキュリティ、生物学、化学、および蒸留試行を検出するように設計された分類器は、最新の Claude Opus モデルにフォールバックし、ユーザーに通知が届きます。
しかし、同社のシステムカード[PDF]によれば、反競争監視は「迅速な変更、方向転換、パラメータ効率の良い微調整（PEFT）などの方法によって有効性を制限することになる」という。
予告なしの「即時変更」は機能的には中間者攻撃ですが、Anthropic 社は、「トラフィックの約 0.03 パーセントに影響を与え、組織の 0.1 パーセント未満に集中する」と推定しています。
開発者の Clay Merritt 氏は、「Anthropic の Fable 5 は、AI/ML の作業を検出すると、静かに答えを妨害します。拒否はありません。通知はありません。ユーザーには見えない意図的な劣化。」
Anthropic はサイバー防御者と重要な人物を期待しています

インフラストラクチャプロバイダーは、Claude Mythos 5 モデルを使用します。このモデルは、Fable 5 の基礎となるモデルを共有しますが、同じ保護手段はありません。ただし、これを行うには、同社の「Project Glasswing」プログラム、または一部の生物学研究者向けに展開されているトラステッド・アクセス・プログラムに参加する必要がある。
モデルの削除（ガードレールの撤去）を支援するサービス、 Abliteration.ai の創設者であるデボン氏（要請により姓は伏せられている）は、電話インタビューで、大手 AI 研究所からはある程度の恐怖を煽ったりマーケティング上の誇大宣伝が行われている一方で、フロンティアモデルがどのように使用されるかについて正当な懸念があると言っても過言ではないと語った。
「アンスロピックは、人々が彼らのブランドを信頼し、（拒否に）対処するだろうということで、彼らのブランドに大きな賭けをしている」と彼は語った。 「しかし、長期的には、人々は自分たちの生活や情報を一元的に管理するこれらの企業をただ受け入れるわけではありません。」 ®
最新情報：水曜夕方にThe Registerに提供された声明の中で、Anthropicの広報担当者は、同社が安全対策を厳格にしすぎたことを認め、生物学研究における誤検知を減らすためにも取り組んでいると述べた
「私たちは、フロンティア LLM 開発のための Fable 5 の保護手段を変更して、それが見えるようにしています。
「今週から、フラグが立てられたリクエストは明らかに Opus 4.8 に戻ります。 API では、フラグが設定されたリクエストは拒否の理由を返します。それが起こるたびにこれがわかります。
「実際には、我々の現在の安全策は、フロンティアスケールのLLMデータパイプラインや特定の非標準チップ向けのカーネル開発など、少数の狭いタスクをカバーしている。これらの安全策は、外国の敵対者が重大な安全上のリスクを引き起こす方法で我々の最も有能なモデルを使用することを防ぐ。米国とその同盟国は」

es は、フロンティア チップと、それらを最大限に実行する高度に最適化されたソフトウェアで優位性を保っています。これらの保護手段により、たとえば、敵対者が開発したチップを最適化することによって、クロードがその利点を損なうために利用されることがなくなります。また、これらは、競合する AI システムを開発するために当社のモデルを使用することを禁止する当社の利用規約を維持するのにも役立ちます。これは、主要な AI プロバイダーの標準的な制限です。これらは、コーディングや ML 作業の大部分には影響しません。
「それらを表示するか非表示にするかを決定する際に、私たちは選択を迫られました。隠されたセーフガードは調査や回避がより困難です。これは、セーフガードの対象をより狭く設定できることを意味します。現在の使用状況では、分類子がタスクの約 0.05% でトリガーされ、影響を受ける組織は 0.05% 未満であることが示されています。可視のセーフガードをより堅牢にするためには、より広い網を張る必要があり、その結果、より多くのリクエストに誤ってフラグが立てられることになります。
「私たちは間違ったトレードオフを選択し、正しいバランスを取れなかったことをお詫びします。これらの保護手段を構築することは、複雑な技術的な課題です。新しい脅威に対応するためにこれらの分類子を改良するにつれて、ユーザーはより多くの誤検知を経験する可能性があります。私たちはこれらをできるだけ早く減らすよう取り組んでいます。」
セキュリティ
すべての従業員のパスワードは 1 つの Excel ファイルに保存されていました
CEO は、これがいくつかの電子メールの問題に対処する最善の方法であると考えました
中国のエージェントがボットネットを再構築し、AI データセンターの議論をかき回しているのを捕まえた
ZTE、TNBエネルギー移行カンファレンスで統合AI、コネクティビティ、デジタルユーティリティテクノロジーをデモンストレーション
パートナーコンテンツ: 高度な AI およびスマート インフラストラクチャ ソリューションを通じてマレーシアの送電網の近代化とエネルギー転換を推進
記憶とパーソナライゼーションにより、AI はユーザーが聞きたいことを教えてくれる可能性が高くなります
特に企業にとって、少しの知識は危険です

アプリケーション
OSプラットフォーム
はい！それは本当です！ Windows 11 はエージェント プラットフォームです
これまでずっとそうだったが、Microsoft はそれに気づいていなかった
大ヒットの新しい Raspberry Pi プロジェクトは、あらゆる画面を昔ながらの VCR に変えます
派手なメニューや高解像度を必要とする人はいないでしょうか。 240-MP はメディア ファイルを 1999 年のように再生します
セキュリティ
すべてのパスワードは Active Directory の説明フィールドに保存されました
公共部門
GOV.UK が Stripe を売却し、支払いをオランダに変更
セキュリティ
GitHub がワーム感染の疑いで 70 以上の Microsoft リポジトリを破壊し、CI/CD パイプラインを破壊
セキュリティ
シグナル、英国がヌード画像のデバイスをスキャンする計画は「私たち全員を危険にさらす」と主張
データベース
Microsoft は Amazon RDS の BYOL を許可しています。繰り返しますが、Microsoft は Amazon RDS の BYOL を許可しています
ボラティリティを乗り越えて成功する: 不確実な市場における永遠の優位性
消費ベースの運用モデルがどのように柔軟性を提供し、効率を向上させ、インフラストラクチャ投資に予測可能性をもたらすかを学びます。
プロンプトからエクスプロイトまで: LLM は API 攻撃をどのように変化させているか
最新のアプリケーションは API 主導で相互接続されており、多くの場合、過剰な権限が与えられているため、AI 支援攻撃の理想的な標的となっています。
未来の構築: Kubernetes 向けエンタープライズ データ サービスのロックを解除する
インフラストラクチャのサイロを排除し、標準化されたエンタープライズ グレードのクラウド ネイティブ プラットフォームを確立する方法を発見してください。
Microsoft 365 が見逃す高度な攻撃を Behavioral AI セキュリティで捕捉
Microsoft 365 は企業コミュニケーションのバックボーンであり、そのネイティブ セキュリティが既知のものや騒々しいものを除外します。
これは、回復力のある次世代の開発および IT 運用を定義する実用的なツールとテクニックを技術的に深く掘り下げるものです。
ランサムウェア侵害の混乱に足を踏み入れ、自分の対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んで裏をかこう

サイバー犯罪者
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えていませんが、私たちも同様です。 Druva の人気イベント「Escape Ransomware」が完全にバーチャルになりました。
Agentic AI 時代のゼロトラスト
ほとんどの組織が依存している ID およびアクセス モデルは、人間以外の ID が独立して動作するのではなく、人間のユーザー向けに構築されています。
Agentic AI 時代のゼロトラスト
ほとんどの組織が依存している ID およびアクセス モデルは、人間以外の ID が独立して動作するのではなく、人間のユーザー向けに構築されています。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで、真の ROI を実現する方法を学びましょう。
特にエンタープライズ アプリケーションの場合、少しの知識は危険です
非常に警戒深い安全分類子が寓話を警告の物語に変える
送電網運営者は新しいビットバーンの建設を支援するのに苦戦する可能性がある
Firefox は登場していますが、Snap と Flatpak はまだリリースされていませんが、デフォルトの AI ヘルパーは眉をひそめるかもしれません
スタッフが AI にスプーンで餌を与えたり、コックアップを修正したりすることで生産性が失われる
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
CentOS の歴史: 生化学者の Linux 趣味プロジェクトがどのようにして企業世界のデフォルトのオペレーティング システムになったのか
Red Hat が Windows が「おそらく適切な製品」であると述べた後、コミュニティが団結したとき
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
プロジェクトのヘッドルームにより、大幅なコストを節約できる可能性があります。

ああ
OpenBSD 7.9 が登場、あらゆる鋭いエッジを誇るダイヤモンドの原石
60 番目のリリースでは、その禁欲的な努力を失うことなく、より多くのコア、遅延休止状態、および基本的な Wi-Fi 6 が追加されています
Fedora: Microsoft は全力で取り組んでいるが、Deepin は見捨てられている
Red Hat の無料ディストリビューションはデスクトップを失うが、重要な新しい友人を作る
LocalSend はスニーカーネットを廃業にします
Apple ロックインを除いた AirDrop と同様
dBase の劣化: データベースの巨人が 47 年後に黒に退色
ブログ投稿の喪の減少が、このベテランアプリのオンラインでの存在感をオフラインに落とす一因となったようだ
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

Hyper-vigilant safety classifiers turn Fable into cautionary tale

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
It blocked us at 'hello!' Anthropic Fable 5 refusing innocuous prompts
Hyper-vigilant safety classifiers turn Fable into cautionary tale
UPDATED Anthropic's newly released Claude Fable 5 generative AI model is trying so hard to be safe that it's hurting its own userbase. Customers attempting to use the AI knowledge regurgitator are reporting that the model is refusing to answer harmless questions, an issue that has annoyed security researchers following past model releases.
Anthropic warned that it had tuned Fable 5's guardrails conservatively: "they’ll sometimes catch harmless requests, though they trigger, on average, in less than five percent of sessions," the company said , promising to "reduce false positives as quickly as we can."
Anthropic spins a Fable of a tamer, safer Mythos
GM gets datacenter fever, decides to build grid-scale sodium-ion batteries
Datacenter growth may run into a power wall by 2030
macOS 27 beta boots Asahi Linux off Apple Silicon
The company did not immediately respond to a request to quantify model refusals. So it's unclear whether the actual false positive rate is greater or less than five percent. But with an estimated 18 to 30 million users worldwide, even a small percentage of thwarted users makes a racket.
Mike Famulare, principal research scientist at the Institute for Disease Modeling, part of the Global Health Division of the Gates Foundation, reports ( #66657 ) that Claude Fable 5 balks at inputs like "Hello."
"In Claude Code, Fable 5's input safety classifier emits a model_refusal_fallback (silent switch to Opus 4.8) on the first turn of essentially every session on my account — including a session whose only user input is the word hello! . No repo content, no tool calls, and no file reads are in context when it fires."
He is not the only frustrated customer. Many other bug reports have been filed in Anthropic's Claude Code GitHub repo since Fable 5 debuted. These include: [Bug] Fable 5 model safety filters causing false positives on benign messages #66587 ; Fable 5 refuses to assist with 'Application Security Architect resume' editing #66655 ; and [Feature Request] Allow Fable 5 usage for non-research lab management systems #67062 , among others .
On social outrage site X.com, Derya Unutmaz, an immunologist and professor at the Jackson Laboratory for Genomic Medicine, notes , "The word 'cancer' is flagged as a biosecurity risk by Claude Fable 5!"
Similar complaints show up in Reddit threads .
Fable 5 is unusual because Anthropic has chosen to conceal safety interventions that try to block rival frontier model development. The classifiers designed to catch cybersecurity, biology and chemistry, and distillation attempts fall back on the latest Claude Opus model and the user gets notified.
But the counter-competition surveillance, per the company's system card [PDF], "will limit effectiveness through methods such as prompt modification, steering vectors, or parameter-efficient fine-tuning (PEFT)."
"Prompt modification" without notice is functionally a man-in-the-middle attack, though one that Anthropic estimates "will impact ~0.03 percent of traffic, concentrated in fewer than 0.1 percent of organizations."
As developer Clay Merritt fumes , "Anthropic’s Fable 5 silently sabotages its answers when it detects AI/ML work. No refusal. No notice. Purposeful degradation invisible to the user."
Anthropic expects cyber defenders and critical infrastructure providers to use its Claude Mythos 5 model, which shares the underlying model of Fable 5 but without the same safeguards. Doing so, however, requires participating in the company's Project Glasswing program or the trusted access program that's being rolled out for select biology researchers.
Devon (last name withheld by request), founder of Abliteration.ai , a service that assists with model abliteration (guardrail removal), told The Register in a phone interview that while there's some degree of fearmongering and marketing hype coming from the big AI labs, it's also fair to say that there are legitimate concerns about how frontier models get used.
"Anthropic's making a big bet on their brand that people will trust their brand so much they'll just deal with [refusals]," he said. "But in the long term, people are not just going to accept these companies that centralize control over their lives and what they can have information about." ®
Update: In a statement provided to The Register on Wednesday evening, an Anthropic spokesperson acknowledged that the company had made its safeguards too stringent and said it was also working to reduce false positives for biological research
"We’re changing Fable 5’s safeguards for frontier LLM development to make them visible.
"Starting this week, flagged requests will visibly fall back to Opus 4.8. On the API, any flagged requests will return a reason for their refusal. You will see this every time it happens.
"In practice, our current set of safeguards covers a handful of narrow tasks like frontier-scale LLM data pipelines and kernel development for certain non-standard chips. These safeguards prevent foreign adversaries from using our most capable models in ways that pose severe safety risks. The US and its allies hold an edge in frontier chips and the highly optimized software that runs them at full potential. These safeguards ensure Claude isn’t used to erode that advantage—by optimizing chips developed by those adversaries, for example. They also help uphold our terms of service, which prohibit using our models to develop competing AI systems—a standard restriction across major AI providers. They do not affect the vast majority of coding and ML work.
"In deciding whether to make them visible or invisible we faced a choice. A hidden safeguard is harder to probe and work around. This means the safeguards can be targeted much more narrowly. Current usage shows that the classifier triggers on about 0.05% of tasks, affecting less than 0.05% of organizations. A visible safeguard needs to cast a wider net to be more robust, resulting in more requests being incorrectly flagged.
"We made the wrong tradeoff and we apologize for not getting the balance right. Building these safeguards is a complex technical challenge: users may experience more false positives as we refine these classifiers to respond to new threats. We are working to reduce these as fast as possible."
SECURITY
Every employee’s password was stored in a single Excel file
The CEO thought this was the best way to deal with some email issues
Chinese agents caught rebuilding botnets and stirring the pot on AI datacenter debate
ZTE Demonstrates Integrated AI, Connectivity and Digital Utility Technologies at TNB Energy Transition Conference
PARTNER CONTENT: Driving Grid Modernization and Energy Transition in Malaysia Through Advanced AI and Smart Infrastructure Solutions
Memory and personalization make AI more likely to tell you what you want to hear
A little knowledge is a dangerous thing, particularly for enterprise applications
os platforms
Yes! It’s true! Windows 11 is an agentic platform
It always has been, but Microsoft didn’t realize it
Blockbuster new Raspberry Pi project turns any screen into old-school VCR
Who needs fancy menus and high definition? 240-MP will play your media files like it's 1999
SECURITY
All the passwords were stored in Active Directory description fields
public sector
GOV.UK goes Dutch on payments as it dumps Stripe
security
GitHub nukes 70+ Microsoft repos, breaks CI/CD pipelines, following suspected worm infections
Security
Signal says UK plan to scan devices for nude images 'endangers us all'
databases
Microsoft allows BYOL for Amazon RDS. Repeat, Microsoft allows BYOL for Amazon RDS
Thriving Through Volatility: The Everpure Advantage in an Uncertain Market
Learn how a consumption-based operating model provides flexibility, improves efficiency, and brings predictability to infrastructure investments.
From Prompt to Exploit: How LLMs Are Changing API Attacks
Modern applications are API-driven, interconnected, and often over-permissioned, making them an ideal target for AI-assisted attacks.
Architecting the Future: Unlocking Enterprise Data Services for Kubernetes
Join us to discover how to eliminate infrastructure silos and establish a standardized, enterprise-grade cloud-native platform.
Catch the Advanced Attacks Microsoft 365 Misses with Behavioral AI Security
Microsoft 365 is the backbone of enterprise communication, and its native security filters out the known and the noisy.
This is your technical deep-dive into the practical tools and techniques that define the next generation of resilient Dev and IT operations.
Step into the chaos of a live ransomware breach, test your response skills, and team up with other IT and security pros to outsmart cybercriminals
Virtual Cyber Recovery Simulation
Ransomware attacks aren’t slowing down, and neither are we. Druva’s hit event, Escape Ransomware, is now fully virtual.
Zero Trust for the Agentic AI Era
The identity and access models most organizations rely on were built for human users, not non-human identities operating independently.
Zero Trust for the Agentic AI Era
The identity and access models most organizations rely on were built for human users, not non-human identities operating independently.
Agentic AI at Scale: From Pilot to Production
Join us to learn how to unlock real ROI by driving adoption of AI at scale.
A little knowledge is a dangerous thing, particularly for enterprise applications
Hyper-vigilant safety classifiers turn Fable into cautionary tale
Grid operators could struggle to support new bit barn construction
Firefox is in, Snap and Flatpak are still out, but a default AI helper may raise eyebrows
Productivity gains lost as staff spoon-feed AI and correct its cock-ups
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
History of CentOS: How a biochemist's Linux hobby project became the enterprise world's default operating system
When a community came together after Red Hat said Windows was 'probably the right product'
Netflix wiz creates app to slash AI bills, then open sources it
Project Headroom could save you big money, too
OpenBSD 7.9 arrives, a diamond in the rough proud of every sharp edge
Sixtieth release adds more cores, delayed hibernation, and basic Wi-Fi 6 without losing its ascetic streak
Fedora: Microsoft is all aboard, but Deepin is dumped
Red Hat’s free distro loses a desktop, but makes an important new friend
LocalSend puts your sneakernet out of business
Like AirDrop, minus the Apple lock-in
dBase debased: Database titan fades to black after 47 years
Blog post mourning decline appears to have helped knock what was left of the veteran app's online presence offline
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
