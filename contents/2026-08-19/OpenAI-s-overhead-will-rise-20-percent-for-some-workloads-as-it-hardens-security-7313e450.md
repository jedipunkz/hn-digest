---
source: "https://www.theregister.com/ai-and-ml/2026/08/19/openais-overhead-will-rise-20-percent-for-some-workloads-as-it-hardens-security/5289303"
hn_url: "https://news.ycombinator.com/item?id=49354828"
title: "OpenAI's overhead will rise 20 percent for some workloads as it hardens security"
article_title: "OpenAI's overhead will rise 20 percent for some workloads as it hardens security"
image: "https://image.theregister.com/260353.jpg?imageId=260353&x=0&y=0&cropw=100&croph=100&panox=0&panoy=0&panow=100&panoh=100&width=1200&height=683"
author: "joebuckwilliams"
captured_at: "2026-08-19T00:39:17Z"
capture_tool: "hn-digest"
hn_id: 49354828
score: 4
comments: 0
posted_at: "2026-08-19T00:14:58Z"
tags:
  - hacker-news
  - translated
---

# OpenAI's overhead will rise 20 percent for some workloads as it hardens security

- HN: [49354828](https://news.ycombinator.com/item?id=49354828)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/08/19/openais-overhead-will-rise-20-percent-for-some-workloads-as-it-hardens-security/5289303)
- Score: 4
- Comments: 0
- Posted: 2026-08-19T00:14:58Z

## Translation

タイトル: セキュリティ強化により、一部のワークロードで OpenAI のオーバーヘッドが 20% 増加する
説明: 多段階の思考連鎖モニタリングの拡張により、フロンティア モデルの作業コストが高くなる

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
2026 年クラウド インフラストラクチャ月間
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
OpenAI はセキュリティを強化するため、一部のワークロードでオーバーヘッドが 20% 増加します
多段階の思考連鎖監視の拡張により、フロンティアモデルの作業コストが高くなる
OpenAIは火曜日、未公開で教師なしのAIモデルがHuggingFaceをハッキングしたことを受けて実施された、モデルトレーニング作業を一時停止するという決定は、AI業界がより強力なセキュリティ対策を導入しようとしている現在も有効であると述べた。これらの対策の中には、計算オーバーヘッドが観測された推論ワークロードの 20% 増加するものもあります。
OpenAIの広報担当者はThe Registerに対し、これらの費用は内部調査を反映したものであり、顧客に直接転嫁されることはないと語った。同社は、現在、または以前の監視体制下で、総推論計算のどの部分がそのような監視の対象となっているかを明らかにしていない。
OpenAIの最高経営責任者（CEO）サム・アルトマン氏はソーシャルメディアへの投稿で、「目の前にある新しいレベルの機能に合わせて、適切な調整、セキュリティ、監視の基準を確実に満たすことができるよう、一部のフロンティアRL（強化学習）トレーニングを一時停止した」と述べた。 「モデルの進歩は現在非常に急速であり、モデルの能力が安全性と整合性のペースを上回っていると感じたら行動を起こすと私たちは常々言っていました。」
アルトマン氏は、おそらく遅れている「アストラ」の新モデルが間もなく出荷されると依然として期待していると語った。トレーニングの一時停止は、さらに先のリリースに影響します。
OpenAIは投稿の中で、先月と同様にモデルの暴走を防ぐため、監視、モデルの調整、セキュリティ対策に重点を置く計画を改めて表明した。 HuggingFace 事件の後、OpenAI は「フロンティアを一時停止した」

コードを実行したり、インターネットにアクセスできるツールを使用したりする実行のための研究クラスター内のモデル推論。」
同企業によると、一部のワークロードの実行は許可されるが、サンドボックス化、ネットワーク分離、継続的なセキュリティテストを含むより厳格なセキュリティ体制の下で移行できるようになるまで、その他のワークロードは一時停止されるという。
「計画されている最大のフロンティアRL（強化学習）の実行は保留されたままだが、モデルの動作を評価し、安全策を検証し、続行する前に整合性のより多くの証拠を確立するために小規模なトレーニングと評価を実施している」と同社は書いている。
強化学習とは、AI エージェントが望ましい結果に対する報酬を得ることで環境について「学習」する試行錯誤のプロセスを指します。
OpenAIはまた、「思考」モデルがタスクを個別のステップに分解し、ステップごとに中間テキスト出力を生成する技術である、思考連鎖プロセスの監視を拡大していると述べた。
期限切れのクレジットカードが研究者によって復活し、不正な支払いが可能に
メタアプリと Google モバイルアプリがユーザーデータを大量に消費: 研究
CISA、現在悪用されている Ray RCE バグの修正に連邦政府に 3 日間の猶予を与える
Apple、スパイウェア悪用の巣窟となった画像処理の穴を塞ぐ
同社の以前のアプローチは、高リスクのワークロード、特にフロンティア モデルの社内展開とフロンティア RL トレーニングの実行に焦点を当てていました。
対照的に、OpenAI によれば、その新しい監視セットアップは、GPT-5.6 Sol 以上の機能レベルのモデルのツールを含むすべての RL トレーニングと評価をカバーします。そして、Astra が重要なサイバー機能を備えていると判断した OpenAI は、RL トレーニングとテストだけでなく、Astra によるすべての推論を対象とする追加の監視要件を追加しました。
「これらの安全策には意味のあるコンピューティングが必要です」とOpenAIは述べた。 「私たちの現在の推定では、モニは

監視対象の推論コンピューティングの約 20% でオーバーヘッドが発生しますが、コストはトレーニングと評価のワークロードによって大幅に異なります。」
OpenAI は、監視スキームの実装についての詳細を今後の投稿で共有する予定です。同社は昨年発表した研究で、思考連鎖の監視はモデルの不正動作を検出する効果的な方法であると述べたが、指示に厳密に従うようにモデルを直接最適化することは「すべての不正動作を排除するわけではなく、モデルがその意図を隠す可能性がある」と警告した。
モデル思考監視のコストを顧客に転嫁しないという同社の保証を信じることにした場合、OpenAI の損失は増加することになる。 OpenAIが上場したとしても、それが持続可能なスタンスになるとは考えにくい。
しかし、同社が AI インフラストラクチャに 6,000 億ドル以上を費やしていると報告されており、少なくとも 2030 年までは採算が取れないと予想されていることを考えると、不確実なセキュリティのためにこれ以上の出費は何でしょうか? ®
システム
Cerebras CS-4 ラック システムは、AI パフォーマンスの最後の一滴までジュース チップを提供します
次世代システムは、ラックに 3 倍のチップを詰め込みながら、チップあたりのパフォーマンスを 2 倍にします
OpenAI はセキュリティを強化するため、一部のワークロードでオーバーヘッドが 20% 増加します
多段階の思考連鎖監視の拡張により、フロンティアモデルの作業コストが高くなる
プラットフォーム エンジニアリング 2.0: プラットフォームは別の時代に構築されました。 AIが暴露しただけ
パートナーのコンテンツ: プラットフォーム エンジニアリングが議論に勝ちました。今では、AI 時代に向けて急速に成長し、進化する必要があります。
期限切れのクレジットカードが研究者によって復活し、不正な支払いが可能に
有効期限チェックにギャップがあると、死んだプラスチックが再び購入される可能性があります
コラムニスト
嫌いになってもいい、AI はここに残る
良いニュースは？最悪の事態の一つ、テクノロジー大手

すべてをコントロールする、もうすぐ終わるかもしれない
メタアプリと Google モバイルアプリがユーザーデータを大量に消費: 研究
ザック帝国はアップルやマイクロソフトの3倍の食欲を誇っている一方、グーグルはリーダーボードを独占していると開発者調査が示した
AIとml
Google、墜落した航空会社スピリットのデータをオークションで購入、理由はAI
AIとML
Excel のコパイロット機能がごみ箱に向かう
デボプス
リポジトリのダウンロードのエラー率が 50% に達するため、GitHub に問題が発生
オフビート
レゴの超巨大ハッブルはもう少し輝くべきだ
aiとml
Anthropicは、テキスト透かしスキームは重要でない単語に依存していると述べています
システム
部品が合わなかったため、技術者はドライバーを取り出しました。するとマザーボードから何かが飛んできた
次世代システムは、ラックに 3 倍のチップを詰め込みながら、チップあたりのパフォーマンスを 2 倍にします
AI の推論エンジンをソーシャル エンジニアリングする方法
企業がモデルのオーケストレーションに苦戦する中、AI ゲートウェイは有望に見える
「あなたを夏の午後に例えてみませんか」というようなことを言うと、他の人も採用する可能性が高いと思われます
中国のAI研究所は前進を続ける一方、米国の研究所は守りを固める
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を持っているのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
FOSS は Microsoft の独占を 1 つ打ち破りました。 20年間の失敗を経て、次の失敗を打ち破る時が来た
一言
GNOME は Windows のように見えることができ、Flashback は拡張機能なしで実行できます
新しい「シンプルタスクバー」はオプションですが、よりシンプルで安定した方法があります
しばらく黙ってください。

x86-32 での Debian の最終リリース
新しい Debian バージョンが 13.6 および 12.15 の形で FOSSland に登場
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
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

Expanded multistage chain of thought monitoring makes frontier model work more expensive

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
OpenAI's overhead will rise 20 percent for some workloads as it hardens security
Expanded multistage chain of thought monitoring makes frontier model work more expensive
OpenAI on Tuesday said its decision to suspend model training work, implemented after unreleased, unsupervised AI models hacked HuggingFace, remains in effect as the AI biz tries to implement stronger security measures. Some of those measures will increase compute overhead by 20 percent of the observed inference workload.
An OpenAI spokesperson told The Register that those costs reflect internal research and won't be passed on directly to customers. The company has not revealed what portion of its total inference compute is subject to such monitoring now, or under its prior monitoring regime.
"We have paused some frontier RL [reinforcement learning] training to ensure that we can meet the appropriate alignment, security and monitoring standards for the new level of capabilities in front of us," OpenAI CEO Sam Altman wrote in a social media post . "Model progress is now extremely rapid, and we always said we would take action if we felt that model capabilities were outstripping the pace of safety and alignment."
Altman said he still expects new models, presumably the delayed Astra , will ship soon. The training pause affects further-out releases.
OpenAI in its post reiterated its plans to focus on monitoring, model alignment, and security measures to prevent its models from running amok as they did last month. Following the HuggingFace incident, OpenAI "paused frontier model inference in research clusters for runs that could execute code or use tools that could access the internet."
The biz said it allows some workloads to run, but paused others until they can be moved under a more stringent security regime that includes sandboxing, network isolation, and continuous security testing.
"Our largest planned frontier RL (reinforcement learning) run remains on hold while we conduct smaller-scale training and evaluations to assess model behavior, validate our safeguards, and establish more evidence of alignment before proceeding," the company wrote.
Reinforcement learning refers to the trial-and-error process by which AI agents "learn" about their environment by being rewarded for desired outcomes.
OpenAI also said it is expanding its monitoring of the chain-of-thought process, the technique that sees "thinking" models break down tasks into discrete steps and produce intermediate text output for each step.
Expired credit cards revived by researchers to make unauthorized payments
Meta and Google mobile apps gorge on user data: Study
CISA gives feds 3 days to fix actively exploited Ray RCE bug
Apple plugs image-processing hole ripe for spyware abuse
The company's prior approach focused on high-risk workloads, specifically internal deployments of frontier models and frontier RL training runs.
In contrast, OpenAI says, its new monitoring setup covers all RL training and evaluations involving tools for models at the capability level of GPT-5.6 Sol or higher. And with the determination that Astra possesses critical cyber capabilities, OpenAI added an additional monitoring requirement that covers all inference with Astra, not just RL training and testing.
"These safeguards require meaningful compute," OpenAI said. "Our current estimates put monitoring overhead at roughly 20 percent of the inference compute being monitored, though the cost varies substantially across training and evaluation workloads."
OpenAI expects to share more details about the implementation of its monitoring scheme in a future post. In research published last year, the company said that chain-of-thought monitoring is an effective way to detect model misbehavior, but cautioned that directly optimizing models to strictly follow instructions "does not eliminate all misbehavior and can cause a model to hide its intent."
If you choose to believe the company's assurance that it will not pass on the cost of model thought policing to customers, it follows that OpenAI's losses will increase. It's difficult to imagine that would be a sustainable stance if OpenAI goes public.
But given the company's reported $ 600+ billion in AI infrastructure commitments and its expectation to remain unprofitable until at least 2030 , what's a bit more expense for the sake of uncertain security? ®
SYSTEMS
Cerebras CS-4 rack systems juice chips for every last drop of AI performance
Next-gen systems double per-chip performance while cramming 3x as many into a rack
OpenAI's overhead will rise 20 percent for some workloads as it hardens security
Expanded multistage chain of thought monitoring makes frontier model work more expensive
Platform Engineering 2.0: your platform was built for a different era. AI just exposed it
PARTNER CONTENT: Platform engineering won the argument. Now it has to grow up fast and evolve for the AI era.
Expired credit cards revived by researchers to make unauthorized payments
Gaps in expiry checks could let dead plastic make purchases again
COLUMNISTS
Be a hater all you want, AI's here to stay
The good news? One of the worst bits, tech giants controlling it all, might soon be over
Meta and Google mobile apps gorge on user data: Study
Zuck's empire declares triple the appetite of Apple or Microsoft, while Google packs the leaderboard, dev survey shows
AI and ml
Google buys crashed airline Spirit’s data at auction, because AI
AI and Ml
Excel's Copilot function is headed for the Recycle Bin
DEVOPS
GitHub has Issues as repo downloads hit 50% error rate
OFFBEAT
Lego's supersized Hubble deserves a little more shine
ai and ml
Anthropic says text watermarking scheme relies on inconsequential words
systems
Part didn't fit so techie got out his screwdriver. Then something flew off the motherboard
Next-gen systems double per-chip performance while cramming 3x as many into a rack
How to social engineer an AI's reasoning engine
AI gateways look promising as companies struggle with model orchestration
'Shall I compare thee to a summer's afternoon' is the sort of thing this will make, and others look likely to adopt it
Chinese AI labs keep moving forward while US labs play defense
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
FOSS smashed one Microsoft monopoly. After 20 years of failure, it's time to smash another
Word up
GNOME can look like Windows – and Flashback can do it without extensions
New 'Simple-taskbar' is an option, but there's a simpler, stabler way
A moment of silence, please, for the final release of Debian on x86-32
New Debian versions hit FOSSland in the form of 13.6 and 12.15
Baddies caught exploiting extensions bugs with perfect 10 scores on vulnerable Joomla websites
Flaws in iCagenda, Balbooa Forms extensions can impact open source CMS that powers a million sites worldwide
Frame: A new X11 server – implemented directly in assembly
Joins yserver, Phoenix, and of course XLibre – and outlier Arcan
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
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
