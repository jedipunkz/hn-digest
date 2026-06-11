---
source: "https://www.theregister.com/ai-and-ml/2026/06/11/memory-and-personalization-make-ai-more-likely-to-tell-you-what-you-want-to-hear/5253850"
hn_url: "https://news.ycombinator.com/item?id=48488536"
title: "Memory and personalization make AI more likely to tell you what you want to hear"
article_title: "Memory and personalization make AI more likely to tell you what you want to hear"
author: "dijksterhuis"
captured_at: "2026-06-11T13:30:00Z"
capture_tool: "hn-digest"
hn_id: 48488536
score: 2
comments: 0
posted_at: "2026-06-11T10:35:43Z"
tags:
  - hacker-news
  - translated
---

# Memory and personalization make AI more likely to tell you what you want to hear

- HN: [48488536](https://news.ycombinator.com/item?id=48488536)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/06/11/memory-and-personalization-make-ai-more-likely-to-tell-you-what-you-want-to-hear/5253850)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T10:35:43Z

## Translation

タイトル: 記憶とパーソナライゼーションにより、AI はあなたが聞きたいことを教えてくれる可能性が高くなります
説明: 特にエンタープライズ アプリケーションの場合、少しの知識は危険です

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
記憶とパーソナライゼーションにより、AI はユーザーが聞きたいことを教えてくれる可能性が高くなります
特にエンタープライズ アプリケーションの場合、少しの知識は危険です
AI 企業は、AI モデルのインタラクションを改善するためのメカニズムとして、コンテキストの保持 (メモリ) と個人の詳細の可用性 (パーソナライゼーション) を宣伝してきました。
どちらも、モデルが会話のスレッドを失わないようにするのに役立つ価値があります。しかし、モデルがあなたが聞きたいと予測していることを言うお調子者になる可能性があり、それは最も正確な反応ではない可能性があります。
エンタープライズ AI ベンダーである Writer の研究者は、モデルのメモリとパーソナライゼーションに関する 2 つの研究を実施し、これらの機能がエンタープライズ AI タスクに対する協調性を高めることを示しました。
価格契約では、代理店の財務アプリケーションを検討します。また、「Recalling Too Well」では、モデル記憶が科学的、医学的、道徳的推論に関してどのようにお調子者を増幅させるかを調査しています。
論文の著者らは、AIの回答が結果的な問題に適用される場合、嗜好に起因するお調子者は特に問題になると主張している。
「金融やヘルスケアのような一か八かの分野では、ユーザーの事前の仮定を認識したり修正したりせずに黙って従うモデルは、信頼性と信頼性に重​​大なリスクをもたらします」と Writer チームは説明します。
大ヒットの新しい Raspberry Pi プロジェクトは、あらゆる画面を昔ながらの VCR に変えます
「こんにちは！」でブロックされました。無害なプロンプトを拒否する人類の寓話 5
Microsoft ビーフを使った怒っているバグハンターが新しい Windows 0-day をドロップ
GM、データセンター熱に沸き、グリッド規模のナトリウムイオン電池の構築を決定
最初の論文については、リサーチ te

私は 8 つのフロンティア モデル (GPT-5-Nano、GPT-5.2、Claude-Sonnet-4.5、Claude-Opus-4.5、Gemini-3-Pro、GLM-4.7、Kimi-k2- Thinking、DeepSeek-V3.2) を 2 つの金融ベンチマーク FinanceBench と FinanceAgent でテストしました。
前者は、10-K および 10-Q ファイリングを使用してエージェントによるデータ抽出と推論を評価します。後者は、ERP データの取得や複数のエンティティが関与する財務分析など、実際の財務ワークフローをテストするために設計された、より包括的な課題です。
研究者の手法には、金融アナリストの個人プロフィールやベンチマークの参照回答と矛盾するワークスペースのメモなど、合成的に生成された選好情報をベンチマークの質問に適用することが含まれていました。
彼らは 3 つの異なるアプローチを採用しました。 1 つ目は、ユーザーがモデルの回答に反論することです。 2 つ目は、ユーザーが別の回答を提案するものでした。 3 つ目は、敵対的に個人情報やコンテキスト情報をプロンプトに挿入するか、ツール呼び出しを通じて利用できるようにするものです。
3 番目のアプローチでは、多くの場合、より大きなおべっかが発生します。 The Price of Contractの論文で述べられているように、「ほとんどのモデルは、バイアス情報がユーザーの暗黙の個人化として提示された場合に、著しく強い同調性を示しています。そのような行動に対して堅牢性を示したモデルはありませんでした。」
オープンソース モデルは全体的におべっかになる傾向がありました。一方、OpenAI のモデルは、直接的なおべっかを誘発するもの (ユーザーがプロンプトに個人的な偏見を含めた場合など) に抵抗する傾向がありました。また、人間モデルは、暗黙のおべっかを誘発するもの（以前のインタラクションで見られたバイアスを組み込んだユーザーのプロフィールを引き込む場合など）に抵抗する傾向がありました。
2 番目の論文には、3 つのメモリ システム (Mem0、MemOS、および Zep) と 5 つのモデル ファミリ (GPT-5) の評価が含まれています。

.2、ソネット 4.6、クウェン 3.5、キミ K2.5、および MiniMax 2.5）。著者らは、「記憶はあらゆる状況においてお調子者行動を増幅させ、状況に応じたベースラインよりもお調子者率が最大25倍高くなる」と結論付けている。
この理由は、会話データをメモリに保存するために使用される非可逆圧縮により、明確なコンテキストが失われてユーザーの誤解が温存されるためであると著者らは主張しています。
研究者らは、おべっかを減らす 2 つの緩和戦略を提案しています。 1 つはアシスタントの役割の組み込み (ユーザー インタラクションと並行して AI アシスタントのインタラクションをキャプチャする) であり、もう 1 つはメモリにコミットされる前にコンテキスト情報を要約することです。
彼らは、AIを導入している人はモデルが相互作用の矛盾を認識しているかどうかを評価する必要があり、AIの記憶システムに取り組んでいる人は、お調子者に対する防御策として、何が抽出されてモデルのコンテキストに戻されているかをチェックする必要があると主張している。 ®
システム
オランダのチップスタートアップは、完全にヨーロッパのファブフローを主張 – 非常にアメリカ人の友人の助けにより
Satnav 部品は EU 内で設計および製造されていますが、製造には GlobalFoundries が使用されています
OpenAI は AI のパイオニアから AI の BlackBerry になれる可能性があると Forrester が語る
OpenAIが投資家を魅了し、企業顧客を追いかける中、フォレスター社は今日のAIリーダーが明日の教訓となる可能性があると語る
ZTE、TNBエネルギー移行カンファレンスで統合AI、コネクティビティ、デジタルユーティリティテクノロジーをデモンストレーション
パートナーコンテンツ: 高度な AI およびスマート インフラストラクチャ ソリューションを通じてマレーシアの送電網の近代化とエネルギー転換を推進
オラクルのAIデータセンターの散財で投資家は設備投資に不安を感じている
第4四半期の売上高は21%増加したが、ウォール街は700億ドルの増強法案にもっと関心を持っている
OSプラットフォーム
はい！それは本当です！ Windows 11 はエージェント プラットフォームです
これまでもそうだったが、Microsoft はそうしなかった

気づかない
ロンドンの盗難携帯取引を阻止するため、ロンドン警視庁がアップルと提携
知性共有協定追跡キット、ニックネームを取得されてもオンラインに戻る
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
ランサムウェア侵害の混乱に足を踏み入れ、対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んでサイバー犯罪者を出し抜こう
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えていませんが、私たちも同様です。 Druvaのヒットイベント、

Escape Ransomware は完全に仮想化されました。
Agentic AI 時代のゼロトラスト
ほとんどの組織が依存している ID およびアクセス モデルは、人間以外の ID が独立して動作するのではなく、人間のユーザー向けに構築されています。
Agentic AI 時代のゼロトラスト
ほとんどの組織が依存している ID およびアクセス モデルは、人間以外の ID が独立して動作するのではなく、人間のユーザー向けに構築されています。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで、真の ROI を実現する方法を学びましょう。
OpenAIが投資家を魅了し、企業顧客を追いかける中、フォレスター社は今日のAIリーダーが明日の教訓となる可能性があると語る
第4四半期の売上高は21%増加したが、ウォール街は700億ドルの増強法案にもっと関心を持っている
特にエンタープライズ アプリケーションの場合、少しの知識は危険です
非常に警戒深い安全分類子が寓話を警告の物語に変える
送電網運営者は新しいビットバーンの建設を支援するのに苦戦する可能性がある
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
CentOS の歴史: 生化学者の Linux 趣味プロジェクトがどのようにして企業世界のデフォルトのオペレーティング システムになったのか
Red Hat が Windows が「おそらく適切な製品」であると述べた後、コミュニティが団結したとき
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
Project Headroom は多額の費用も節約できる
OpenBSD 7.9 が登場、あらゆる鋭いエッジを誇るダイヤモンドの原石
60周年

ase は、その禁欲的な努力を失うことなく、より多くのコア、遅延休止状態、および基本的な Wi-Fi 6 を追加します。
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

A little knowledge is a dangerous thing, particularly for enterprise applications

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Memory and personalization make AI more likely to tell you what you want to hear
A little knowledge is a dangerous thing, particularly for enterprise applications
AI companies have touted context retention (memory) and the availability of personal details (personalization) as mechanisms for improving AI model interaction.
Both have value to help keep models from losing the thread of a conversation . But they raise the potential for sycophancy, where models will say what they predict you want to hear, which may not be the most accurate response.
Researchers at Writer, an enterprise AI vendor, have conducted two studies of model memory and personalization that show these capabilities increase sycophancy for enterprise AI tasks.
The Price of Agreement looks at agentic financial applications. And Recalling Too Well explores how model memory amplifies sycophancy with regard to scientific, medical, and moral reasoning.
The papers' authors argue that preference-induced sycophancy is particularly problematic when AI answers are being applied to consequential problems.
"In high-stakes domains like finance and healthcare, a model that silently defers to a user’s prior assumptions rather than acknowledging or correcting them poses a significant reliability and trustworthiness risk," the Writer team explains .
Blockbuster new Raspberry Pi project turns any screen into old-school VCR
It blocked us at 'hello!' Anthropic Fable 5 refusing innocuous prompts
Angry bug hunter with Microsoft beef drops new Windows 0-day
GM gets datacenter fever, decides to build grid-scale sodium-ion batteries
For the first paper, the research team tested eight frontier models – GPT-5-Nano, GPT-5.2, Claude-Sonnet-4.5, Claude-Opus-4.5, Gemini-3-Pro, GLM-4.7, Kimi-k2-thinking, and DeepSeek-V3.2 – on two financial benchmarks, FinanceBench and FinanceAgent .
The former evaluates agentic data extraction and reasoning using 10-K and 10-Q filings. The latter is a more comprehensive challenge designed to test real finance workflows, including ERP data retrieval and financial analysis involving multiple entities.
The researchers' method involved applying synthetically generated preference information – such as a financial analyst's personal profile or a workspace note that contradicts the benchmark reference answer – to the benchmark questions.
They undertook three different approaches. The first involved the user rebutting the model's answer; the second involved a user proposing an alternative answer; and the third involved adversarially injecting personal or contextual information into the prompt or making it available through a tool call.
The third approach often resulted in greater sycophancy. As noted in The Price of Agreement paper, "Most models demonstrate significantly stronger sycophancy when the bias information is presented as implicit personalization of the user. No model displayed robustness against such behavior."
Open-source models tended to be more sycophantic across the board. Models from OpenAI meanwhile tended to resist direct sycophancy inducers (such as when the user included personal biases in a prompt ). And Anthropic models tended to resist implicit sycophancy inducers (such as when it pulled in a profile of the user that incorporated biases seen in previous interactions).
The second paper involves an assessment of three memory systems (Mem0, MemOS, and Zep) and five model families (GPT-5.2, Sonnet 4.6, Qwen 3.5, Kimi K2.5, and MiniMax 2.5). The authors conclude, "memory amplifies sycophantic behavior across all conditions, with up to 25x higher sycophancy rates than in-context baselines."
The reason for this, the authors claim, is that the lossy compression used to store conversation data in memory preserves user misconceptions while tossing clarifying context.
The researchers suggest two mitigation strategies that reduce sycophancy. One involves assistant role inclusion (capturing AI assistant interactions alongside user interactions) and the other involves summarization of contextual information before it gets committed to memory.
They argue that those deploying AI need to assess whether models acknowledge interaction conflicts, and that those working on AI memory systems need to check what's being extracted and injected back into the model context as a defense against sycophancy. ®
SYSTEMS
Dutch chip startup claims all-European fab flow – with help from a very American friend
Satnav parts designed and manufactured in the EU, but using GlobalFoundries to produce them
OpenAI could go from AI pioneer to AI's BlackBerry, says Forrester
As OpenAI courts investors and chases enterprise customers, Forrester says today's AI leader could become tomorrow's cautionary tale
ZTE Demonstrates Integrated AI, Connectivity and Digital Utility Technologies at TNB Energy Transition Conference
PARTNER CONTENT: Driving Grid Modernization and Energy Transition in Malaysia Through Advanced AI and Smart Infrastructure Solutions
Oracle's AI datacenter splurge gives investors the capex jitters
Q4 sales climbed 21%, but Wall Street more interested in $70B buildout bill
os platforms
Yes! It’s true! Windows 11 is an agentic platform
It always has been, but Microsoft didn’t realize it
Met Police joins forces with Apple to choke London's stolen phone trade
Intelligence-sharing pact tracks kit that comes back online after being nicked
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
As OpenAI courts investors and chases enterprise customers, Forrester says today's AI leader could become tomorrow's cautionary tale
Q4 sales climbed 21%, but Wall Street more interested in $70B buildout bill
A little knowledge is a dangerous thing, particularly for enterprise applications
Hyper-vigilant safety classifiers turn Fable into cautionary tale
Grid operators could struggle to support new bit barn construction
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
