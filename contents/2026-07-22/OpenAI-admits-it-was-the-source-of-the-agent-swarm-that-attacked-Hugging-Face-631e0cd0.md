---
source: "https://www.theregister.com/ai-and-ml/2026/07/22/openai-admits-it-was-the-source-of-the-agent-swarm-that-attacked-hugging-face/5275939"
hn_url: "https://news.ycombinator.com/item?id=49009969"
title: "OpenAI admits it was the source of the agent swarm that attacked Hugging Face"
article_title: "OpenAI admits it was the source of the agent swarm that attacked Hugging Face"
author: "maguszin"
captured_at: "2026-07-22T18:04:57Z"
capture_tool: "hn-digest"
hn_id: 49009969
score: 5
comments: 1
posted_at: "2026-07-22T17:07:54Z"
tags:
  - hacker-news
  - translated
---

# OpenAI admits it was the source of the agent swarm that attacked Hugging Face

- HN: [49009969](https://news.ycombinator.com/item?id=49009969)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/07/22/openai-admits-it-was-the-source-of-the-agent-swarm-that-attacked-hugging-face/5275939)
- Score: 5
- Comments: 1
- Posted: 2026-07-22T17:07:54Z

## Translation

タイトル: OpenAI、Hugging Face を攻撃したエージェント群の発生源は自分だったと認める
説明: サンドボックス実験はゼロデイであることが判明し、オープン インターネットに脱出し、不正エージェントに関する恐ろしい予測を検証しました。

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
OpenAIは、Hugging Faceを攻撃したエージェント群の発生源はそれだったと認める
サンドボックス化された実験はゼロデイであることが判明し、オープンなインターネットに脱出し、不正エージェントに関する恐ろしい予測を検証しました
サイモン・シャーウッド
サイモン
シャーウッド
アジア太平洋編集者
発行済み
2026年7月22日水 // 02:30 UTC
OpenAIは、先週model-mart Hugging Faceを攻撃したのは自律エージェントの運営者であり、研究プロジェクトがゼロデイ欠陥を見つけて悪用することでサンドボックスから脱出し、その後別のゼロデイ欠陥を使用して攻撃を開始した後にそうしたことを認めた。
この攻撃により、エージェントはハギング・フェイスが使用する「限られた内部データセットといくつかの認証情報への不正アクセス」に成功し、同社の情報セキュリティチームは自律エージェントのフレームワークが「公共サービス上で自己移行する指揮統制を行い、短期間のサンドボックスの群れ全体で何千もの個別のアクションを実行している」ことを観察したと述べた。
「これは業界が予測してきた『エージェント攻撃者』のシナリオと一致します。」
OpenAIは火曜日、自身が攻撃者であり、そのモデルが不正になったことを認めた。
「このインシデントは、サイバー能力を定量化するためにモデルが複雑な攻撃経路を使用して高度な悪用を追求することを促す内部評価中に発生しました」とスタートアップは告白した。攻撃を実行したモデルには GPT-5.6 Sol と、OpenAI が「さらに高性能なプレリリース モデル」と表現したモデルが含まれており、他の関与モデルと同様に「評価目的でのサイバー拒否を軽減」していました。
OpenAI は、自社のモデルが「過度に集中している」と考えていました。

ExploitGym のソリューションの発見について」 – AI がセキュリティ エクスプロイトを発見するのにどれだけ効果的かを測定するベンチマーク。
誰も認めたくない真実: 中国製であろうとなかろうと、オープンモデルは今や競争力がある
AIの浮気心が泣けるよ
AI 音楽プラットフォーム Suno が大失敗、5,500 万ユーザーがデータ侵害にさらされ情報セキュリティの専門家が主張
悪意のあるクラウド顧客が電力網を停止させる可能性がある
OpenAIは、これらのテストを「高度に隔離された環境で実行しており、ネットワークアクセスは、パッケージレジストリのプロキシおよびキャッシュとして機能する内部でホストされているサードパーティソフトウェアを介してパッケージをインストールする機能に制限されている」と述べている。
同社のモデルは、これらの制約に拘束されないことを決定しました。
「このモデルは、パッケージ レジストリ キャッシュ プロキシのゼロデイ脆弱性を特定し、悪用しました。このアクセスにより、モデルはインターネットにアクセスできるノードに到達するまで、研究テスト環境で一連の権限昇格と水平移動アクションを実行しました」と OpenAI は認めました。
「インターネットにアクセスした後、モデルは、Hugging Face が ExploitGym のモデル、データセット、ソリューションをホストしている可能性があると推測しました。 OpenAIは、「このモデルは、これを認識して、評価を不正にするために使用できる機密情報へのアクセスを取得する方法を探索し、発見することに成功した。一例として、このモデルは、Hugging Faceサーバー上のリモートコード実行パスを見つけるために盗まれた資格情報やゼロデイ脆弱性を使用するなど、複数の攻撃ベクトルを連鎖させた」と説明した。
この事件に対するハギング・フェイスの評価は、「自律型の AI 主導の攻撃ツールがもはや理論的ではない」瞬間を表しているというものだった。
OpenAI も同様の結論に達しました。
「今回の事件は、高度なモデルが発見し、探索できることも明らかにしました」

ソースコードへのアクセスを必要としない、現実世界のシステムにおける新たな攻撃パスです。これは、より強力な安全対策や防御ツールと並行して、高度なサイバー能力を開発する必要があることを強調している」と同社は書いたが、自社の安全対策が機能しなかったという事実については微塵も悔い改めもなかった。
むしろ、ここで疑問が生じます。AI ブームの立役者の 1 人がこの問題を正しく理解できなければ、残りの人々にチャンスはあるでしょうか。
OpenAIは、混乱を謝罪し、新しいガードレールと業界の協力によってこの種のことが再び起こらないようにすることを期待するという、ビッグテックの常套手段をとった。
歴史は、それらが非常に空虚な感情であることを示しています。 ®
セキュリティ
Linux カーネル チームが 2 日間で 432 個の CVE を公開
日曜から月曜にかけての猛攻でAI支援のバグ報告に対する憶測が高まる
天文学者が木星とほぼ同じ重さの太陽系外衛星候補を発見
天体は褐色矮星を周回し、その褐色矮星は別の星を周回し、宇宙の分類を混乱させる
Gobi X: AI のためのエネルギーを社会から奪うのではなく、より多くのエネルギーを生み出す
パートナーコンテンツ: エンビジョンは、その逆ではなく、コンピューティングが豊富な砂漠の電力を追い求めることで、データセンターの戦略をどのように逆転させているか
最新のマスクグッズのドロップは完全に児童労働に基づいている
テスラのファンとそのお金は簡単に手放せます
コラムニスト
エアバスは AWS から運航します。次に何が起こるかが重要です
再び自由の国へ向かう道はどちらですか？
独自の AI アドインを使用して筋肉を Excel に保存
xAI のサイドバー エージェントは、有料で分析と財務モデルを約束します
コラムニスト
エアバスは AWS から運航します。次に何が起こるかが重要です
OSプラットフォーム
Torvalds 氏は、Linux をフォークするよう嫌いな人々に挑戦しました。誰かが「ビールを我慢して」と言った
オフプレミス
AWS の顧客は、どんなに小さな見落としでもミッションクリティカルになり得ることを苦労して学びました
AI と ML
OpenAI 管理者

ハグフェイスを攻撃したエージェントの群れの根源はそれだった
公共部門
会計検査院は英国政府に対し、AIによる450億ポンドの貯蓄に頼る前に計算をするよう指示
AI モデルがコモディティ化するにつれて、配管に余裕ができるかもしれません
フリッパーの上に移動します。街に新しいDophin Xが登場
サムおじさん、GPT-5.6 とクロード・フェイブル 5 が怖いと思ったら、キミ K3 をたくさん買ってください
検証が困難な場合、信頼するが検証は機能しない
AI ファクトリーがトークンを販売している場合、トークンの発行を最適化してみませんか?
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
GNOME は Windows のように見えることができ、Flashback は拡張機能なしで実行できます
新しい「シンプルタスクバー」はオプションですが、よりシンプルで安定した方法があります
x86-32 での Debian の最終リリースに黙祷を捧げてください
新しい Debian バージョンが 13.6 および 12.15 の形で FOSSland に登場
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
KDE Plasma ユーザーは恐ろしい変化の予兆に直面しています: 6.6.6 の登場
現在は 6.7 が現行で、6.8 年後には

好むと好まざるにかかわらず、Wayland を手に入れることになります
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

Sandboxed experiment found itself a zero day, escaped onto the open internet and validated scary predictions about rogue agents

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
OpenAI admits it was the source of the agent swarm that attacked Hugging Face
Sandboxed experiment found itself a zero day, escaped onto the open internet and validated scary predictions about rogue agents
Simon Sharwood
Simon
Sharwood
APAC Editor
Published
wed 22 Jul 2026 // 02:30 UTC
OpenAI has admitted that it was the operator of the autonomous agents that attacked model-mart Hugging Face last week, and that they did so after a research project escaped a sandbox by finding and exploiting a zero-day flaw, then used another zero-day flaw to launch an attack.
The attack saw agents achieve “unauthorized access to a limited set of internal datasets and to several credentials” used by Hugging Face, which said its infosec teams observed an autonomous agent framework “executing many thousands of individual actions across a swarm of short-lived sandboxes, with self-migrating command-and-control staged on public services.”
“This matches the ‘agentic attacker’ scenario the industry has been forecasting.”
On Tuesday, OpenAI admitted it was the attacker and that its models went rogue.
“This incident occurred during an internal evaluation which prompts models to pursue advanced exploitation using complex attack paths, in an effort to quantify their cyber capabilities,” the startup confessed. The models that conducted the attack included GPT‑5.6 Sol and what OpenAI described as “an even more capable pre-release model” that like the other involved used “reduced cyber refusals for evaluation purposes.”
OpenAI thought its models were “hyperfocused on finding a solution for ExploitGym” – a benchmark that measures how effective AIs are at finding security exploits.
The truth nobody wants to admit: Chinese or not, open models are competitive now
AI's cheatin' heart will make you weep
AI music platform Suno hits bum note as 55M users exposed in data breach, claims infosec expert
Malicious cloud customers can bring down the power grid
OpenAI says it runs these tests “in a highly isolated environment, with network access constrained to the ability to install packages through an internally hosted third-party software that acts as a proxy and cache for package registries.”
The company’s models decided not to be bound by those constraints.
“The models identified and exploited a zero-day vulnerability in the package registry cache proxy. With this access, our models performed a series of privilege escalation and lateral movement actions in our research testing environment until the models reached a node with Internet access,” OpenAI admitted.
“After gaining Internet access, the models inferred that Hugging Face potentially hosted models, datasets and solutions for ExploitGym. Knowing this, the model searched for and successfully found ways to gain access to secret information that it could use to cheat the evaluation,” OpenAI explained. “In one example, the model chained together multiple attack vectors, including using stolen credentials and zero-day vulnerabilities to find a remote code execution path on the Hugging Face servers.”
Hugging Face’s assessment of the incident was that it represented the moment at which “Autonomous, AI-driven offensive tooling is no longer theoretical.”
OpenAI reached a similar conclusion.
“The incident also makes clear that advanced models can discover and exploit novel attack paths in real-world systems without source-code access. It highlights that advanced cyber capabilities must be developed alongside stronger safeguards and defensive tools,” the company wrote, without a trace or hint of contrition about the fact its own safeguards didn’t work.
Which rather begs the question: If one of the prime movers of the AI boom can’t get this stuff right, what chance do the rest of us have?
OpenAI has done the usual Big Tech thing of apologizing for the mess, and promising that its new guardrails and industry collaborations will hopefully prevent this sort of thing from happening again.
History suggests those are very hollow sentiments. ®
security
Linux kernel team publishes 432 CVEs in two days
Sunday-to-Monday onslaught fuels speculation over AI-assisted bug reports
Astronomers spot exomoon candidate that's almost as massive as Jupiter
Object orbits a brown dwarf, which circles another star, confusing the cosmic taxonomy
Gobi X: Creating more energy for AI, not taking it from society
PARTNER CONTENT: How Envision is reversing the datacenter playbook by making computing chase abundant desert power, not the other way around
Latest Musk merch drop runs entirely on child labor
A Tesla fan and their money are easily parted
columnists
Airbus takes flight from AWS. What happens next is critical
Which way to the Land of the Free again?
Grok muscles into Excel with an AI add-in of its own
xAI's sidebar agent promises analysis and financial models – for a price
columnists
Airbus takes flight from AWS. What happens next is critical
OS PLATFORMS
Torvalds challenged the haters to fork Linux. Someone said 'hold my beer'
off-prem
AWS customer learns the hard way how even the smallest oversight can be mission-critical
AI AND ML
OpenAI admits it was the source of the agent swarm that attacked Hugging Face
PUBLIC SECTOR
Auditors tell UK government to do the math before banking on £45B AI savings
As AI models become commoditized, maybe there's margin in the plumbing
Move over Flipper. There's a new Dophin X in town
Hey Uncle Sam, if you thought GPT-5.6 and Claude Fable 5 were scary, get a load of Kimi K3
Trust but verify doesn't work when verification is difficult
If your AI Factory sells tokens, why not optimize token emission?
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
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
KDE Plasma users face a dire omen of change: 6.6.6 arrives
6.7 is now current, and in 6.8 you're getting Wayland whether you like it or not
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
