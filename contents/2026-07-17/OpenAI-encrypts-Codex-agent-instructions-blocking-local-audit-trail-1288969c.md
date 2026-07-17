---
source: "https://www.theregister.com/ai-and-ml/2026/07/15/openai-hides-codex-agent-instructions-behind-encryption-leaving-developers-in-the-dark/5271484"
hn_url: "https://news.ycombinator.com/item?id=48943426"
title: "OpenAI encrypts Codex agent instructions, blocking local audit trail"
article_title: "OpenAI hides Codex agent instructions behind encryption, leaving developers in the dark"
author: "nyku"
captured_at: "2026-07-17T04:54:56Z"
capture_tool: "hn-digest"
hn_id: 48943426
score: 1
comments: 0
posted_at: "2026-07-17T04:46:38Z"
tags:
  - hacker-news
  - translated
---

# OpenAI encrypts Codex agent instructions, blocking local audit trail

- HN: [48943426](https://news.ycombinator.com/item?id=48943426)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/07/15/openai-hides-codex-agent-instructions-behind-encryption-leaving-developers-in-the-dark/5271484)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T04:46:38Z

## Translation

タイトル: OpenAI が Codex エージェントの命令を暗号化し、ローカル監査証跡をブロック
記事のタイトル: OpenAI が Codex エージェントの指示を暗号化の背後に隠し、開発者を闇の中に放置
説明: 開発者は、暗号化された MultiAgentV2 メッセージによりデバッグと監査が困難になることを懸念しています

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
OpenAI は Codex エージェントの命令を暗号化の背後に隠し、開発者を闇の中に置き去りにする
開発者は、暗号化された MultiAgentV2 メッセージによりデバッグと監査が困難になることを懸念しています
OpenAI はその名前が示すほどオープンではありませんでしたが、ますますオープンではなくなりつつあります。
無償の AI 巨人は最近、Codex コマンド ライン インターフェイスのマルチエージェント オーケストレーションを改訂し、サブエージェントに渡されるメッセージを暗号化しました。
OpenAI の Codex は、親エージェントが子エージェントを生成したり、異なるモデルを呼び出す可能性のある他のエージェントにタスクを委任したりする方法であるマルチエージェント オーケストレーションをサポートしています。
Codex/GPT-5.6 では、マルチエージェント v2 と呼ばれるプロトコルが導入されました。これは、ユーザーが宣言した構成設定に決定を委ねるのではなく、ランタイムに作業を割り当てさせることを目的としているようです。
マルチエージェント v2 はまだ開発中であり、OpenAI はまだ正式に文書化していません。ただし、開発者は、新しいエージェントの配管に対応するために Codex に加えられた変更を観察しており、新しい取り決めに懸念を抱いている人もいます。
先月、OpenAI 開発者はプル リクエスト (提案されたコード変更) をマージして、マルチエージェント v2 メッセージ ペイロード (エージェント間で渡されるテキスト命令) を暗号化しました。プル リクエストでは説明テキストの先頭に「理由」という単語が付けられていますが、実際には変更の理由は示されていません。
「現在、マルチエージェント v2 は通常のツール引数とエージェント間コンテキストを通じてエージェント命令をルーティングします。つまり、親モデルはプレーンテキストのタスク テキストを出力でき、Codex はそれを履歴/ロールアウトに保持でき、受信者はそれを通常のアシスタント メッセージ JSON として受信できます。」
「これ

v2 パスを変更することで、モデル呼び出し間でエージェントの指示が暗号化されたままになります: Responses [An OpenAI API - Ed] はモデルから返されたメッセージ引数を暗号化し、Codex はその暗号文のみを転送し、Responses は受信者モデルに対して内部で暗号化を復号します。
プライバシーとセキュリティを強化したい、またはモデルの抽出に役立つデータを隠蔽したいという要望は、これらの変更の正当な理由です。しかし、OpenAIはなぜ変更を行ったのかについては明らかにしていない。
OpenAIから明確なコミュニケーションがないため、開発者らは同社に対し、その実装が他の懸念事項に対応するために監査可能性を犠牲にしないようにするよう求めている。
決済サービス Zolvat の CTO、Ignat Remizov 氏が提起した問題では、「暗号化された配信パスはプライバシーの強化として理解できますが、ローカルのロールアウト履歴、トレース削減、親側の監査/デバッグ画面から人間が判読できるタスク/メッセージ テキストが削除されます。」と述べています。
Patchpocalypse Now: Microsoft は 622 件の Patch Tuesday CVE で先月の記録を上回りました
クロードに気さくに話しかけてもらいたいなら、ヒンディー語かアラビア語を試してみてください
ニューヨーク州がデータセンターの建設を中止する最初の州となる
DeepMindの大脳はアメリカに対し、手遅れになる前にAI標準を設定するよう呼びかけている
レミゾフ氏の懸念は、開発者やコード管理者がエージェントが受け取った指示やエージェントがとった行動を評価するための情報が少なくなることだ。
「皆さん、私たちはスカイネットを構築した後に、その活動を監査できなくなることは望んでいません」と彼は冗談めかして言います。
他の開発者らは、可観測性の喪失に関するレミゾフ氏の懸念に同調し、OpenAIがマルチエージェントの実装がどのように機能するかを競合他社に分からせないようにエージェントのメッセージングをロックダウンしているのではないかと推測している。
OpenAIはコメント要請にすぐには応じていない。 ®
ネットワーク
過去に遡った NTP サーバーにより、オーストラリアのモバイル ネットワークが大規模になる

タゲ
Telstra はパッチをスキップし、変更を記録せず、それが待っている事故であるとは思いもしませんでした
韓国は独自のセキュリティ中心の AI モデルを作成
既存のローカル LLM プロジェクトを安全保障と主権の目的に適応させ、いつか Mythos と一致することを望んでいます
Gobi X: AI のためのエネルギーを社会から奪うのではなく、より多くのエネルギーを生み出す
パートナーコンテンツ: エンビジョンは、その逆ではなく、コンピューティングが豊富な砂漠の電力を追い求めることで、データセンターの戦略をどのように逆転させているか
中国のメモリ禁止はRAMpocalypse救済を打ち切るだろう
2人の米国議員が中王国の半導体メーカーに対する規制強化を推進
PaaS + IaaS
アマゾン ウェブ サービスの最も声高な顧客は現在 EC2 を実行しています
小売財団のリーダー、デイブ・トレッドウェル氏が上級リーダーに就任し、19年間の退役軍人であるデイブ・ブラウン氏が未知の牧草地へ出発
OpenAIは、GPT-5.6が時々ファイルを削除することを認めていますが、それは「正直な間違い」です
データのパージは、新興企業が回避しようと取り組んでいる「不整合な動作」の一例とみなされる
aiとml
ライナス・トーバルズ氏、AI嫌いの人たちにフォークをやめるよう指示
AIとml
Grok Buildがリポジトリ全体をクラウドに送信していたことが発覚した後、マスク氏は粛清を約束
DevOps
HTTP は QUERY メソッドを取得するため、複雑な検索が POST のふりをするのを防ぐことができます
デボプス
Zigの作者はBunのClaude Rustリライトを「未レビューの駄作」と呼ぶ
ソフトウェア
フレーム: 新しい X11 サーバー – アセンブリに直接実装
データのパージは、新興企業が回避しようと取り組んでいる「不整合な動作」の一例とみなされる
モデルは検証を提供せずに信頼を要求します
プレスリリースを掲載しているファブメーカーに注意してください
Thinking Machines の最初のオープンウェイト モデルは、中国の LLM に代わる 9,750 億パラメータです
ワンツーパンチは、低精度 AI が高精度シミュレーションをどのように補完できるかを垣間見せます
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのか」と尋ねる

エーション？さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
x86-32 での Debian の最終リリースに黙祷を捧げてください
新しい Debian バージョンが 13.6 および 12.15 の形で FOSSland に登場
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
KDE Plasma ユーザーは恐ろしい変化の予兆に直面しています: 6.6.6 の登場
現在は 6.7 が最新版であり、6.8 では好むと好まざるにかかわらず Wayland を入手できるようになります。
FOSS クラウドオフィススイート間の競争が激化する中、Collabora が CODE 26.04 をリリース
Markdown のサポートと、よりスマートな数式エラー処理に加え、AI が統合されました (デフォルトではオフになっています)。
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

Developers worry encrypted MultiAgentV2 messages will make debugging and auditing harder

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
OpenAI hides Codex agent instructions behind encryption, leaving developers in the dark
Developers worry encrypted MultiAgentV2 messages will make debugging and auditing harder
OpenAI has never been as open as its name suggests and is becoming even less so.
The free-spending AI giant recently revised the multi-agent orchestration in its Codex command line interface to encrypt messages passed to subagents.
OpenAI's Codex supports multi-agent orchestration, a way to have a parent agent spawn child agents or delegate tasks to other agents that may call out to different models.
Codex/GPT-5.6 introduced a protocol called multi-agent v2 that appears to be geared toward letting the runtime allocate work instead of leaving those decisions to user-declared configuration settings.
Multi-agent v2 is still under development and OpenAI hasn't formally documented it yet. Developers, however, have observed changes made to Codex to accommodate the new agent plumbing – and some are concerned by the new arrangements.
Last month, OpenAI devs merged a pull request (a suggested code change) to encrypt multi-agent v2 message payloads – the text instruction passed between agents. The pull request prefaces its explanatory text with the word "Why" but doesn't actually offer a reason for the change.
It states: "Multi-agent v2 currently routes agent instructions through normal tool arguments and inter-agent context. That means the parent model can emit plaintext task text, Codex can persist it in history/rollouts, and the recipient can receive it as ordinary assistant-message JSON.
"This changes the v2 path so agent instructions stay encrypted between model calls: Responses [An OpenAI API - Ed ] encrypts the message argument returned by the model, Codex forwards only that ciphertext, and Responses decrypts it internally for the recipient model."
A desire to enhance privacy and security, or conceal data that would be useful for model distillation, are sound reasons for these changes. Yet OpenAI has not said why it made the change.
In the absence of clear communication from OpenAI, developers have urged the company to ensure that its implementation doesn't sacrifice auditability to accommodate other concerns.
An issue opened by Ignat Remizov, CTO at payment service Zolvat, says, "The encrypted delivery path is understandable as privacy hardening, but it also removes the human-readable task/message text from local rollout history, trace reduction, and parent-side audit/debug surfaces."
Patchpocalypse Now: Microsoft tops last month's record with 622 Patch Tuesday CVEs
If you want Claude to speak nicely to you, try Hindi or Arabic
New York becomes first state to halt datacenter buildouts
DeepMind bigbrain calls for America to set AI standards before it's too late
Remizov's concern is that developers and code maintainers will have less information to assess the instructions an agent received and the actions it took.
"Guys, we don’t want to build Skynet and then be unable to audit what it’s doing," he quips .
Other developers, echoing Remizov's concern about the loss of observability, speculate that OpenAI has locked its agent messaging down to keep competitors from seeing how its multi-agent implementation works.
OpenAI did not immediately respond to a request for comment. ®
networks
NTP server that travelled back in time caused massive Aussie mobile outage
Telstra skipped a patch, didn’t record changes, had no idea it was an accident waiting to happen
South Korea making its own security-centric AI model
Adapting existing local LLM project for security and sovereignty purposes and hopes to one day match Mythos
Gobi X: Creating more energy for AI, not taking it from society
PARTNER CONTENT: How Envision is reversing the datacenter playbook by making computing chase abundant desert power, not the other way around
Chinese memory ban would cut off RAMpocalypse relief
Two US lawmakers push tighter curbs on chipmakers from the Middle Kingdom
PaaS + IaaS
Amazon Web Services' most vocal customer now runs EC2
Retail foundation leader Dave Treadwell takes over as senior leader and 19-year vet Dave Brown departs for pastures unknown
OpenAI admits GPT-5.6 occasionally deletes files – but it's an 'honest mistake'
Data purges deemed an example of 'misaligned behavior' that upstart is working to avoid
ai and ml
Linus Torvalds tells AI haters to fork off
AI and ml
Musk promises purge after Grok Build caught sending entire repos to the cloud
DevOps
HTTP gets a QUERY method so complex searches can stop pretending to be POST
DEVOPS
Zig creator calls Bun’s Claude Rust rewrite ‘unreviewed slop’
software
Frame: A new X11 server – implemented directly in assembly
Data purges deemed an example of 'misaligned behavior' that upstart is working to avoid
Models demand trust without offering verification
Beware of fab makers bearing press releases
Thinking Machines' first open weights model is a 975 billion parameter alternative to Chinese LLMs
One-two punch offers a glimpse of how low-precision AI can complement high-precision simulations
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
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
Collabora releases CODE 26.04 as rivalry between FOSS cloudy office suites heats up
Now with Markdown support and smarter formula error handling – plus integrated AI, though it's off by default
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
