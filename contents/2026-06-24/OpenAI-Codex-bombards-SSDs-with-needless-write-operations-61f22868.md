---
source: "https://www.theregister.com/ai-and-ml/2026/06/23/openai-codex-bombards-ssds-with-needless-write-operations-costing-millions/5260402"
hn_url: "https://news.ycombinator.com/item?id=48665875"
title: "OpenAI Codex bombards SSDs with needless write operations"
article_title: "OpenAI Codex bombards SSDs with needless write operations, costing millions"
author: "jgalt212"
captured_at: "2026-06-24T22:28:18Z"
capture_tool: "hn-digest"
hn_id: 48665875
score: 9
comments: 0
posted_at: "2026-06-24T21:32:15Z"
tags:
  - hacker-news
  - translated
---

# OpenAI Codex bombards SSDs with needless write operations

- HN: [48665875](https://news.ycombinator.com/item?id=48665875)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/06/23/openai-codex-bombards-ssds-with-needless-write-operations-costing-millions/5260402)
- Score: 9
- Comments: 0
- Posted: 2026-06-24T21:32:15Z

## Translation

タイトル: OpenAI Codex が SSD に不必要な書き込み操作を攻撃する
記事のタイトル: OpenAI Codex が SSD を不必要な書き込み操作で攻撃し、数百万ドルの損害を与える
説明: 不器用なロギング実装により、コストを考慮せずにデータが無駄に流出します

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
OpenAI Codex が SSD を不必要な書き込み操作で攻撃し、数百万ドルの損害を与える
不器用なロギング実装により、コストを考慮せずにデータが無駄に失われる
最新の SSD では、有効期限が切れるまでの書き込みサイクル数が制限されています。現在、OpenAI は、過度のデータ書き込みにより Codex ユーザーのソリッド ステート ドライブ (SSD) の寿命を縮め、デバイスの価値を大幅に低下させている欠陥のあるログ実装の修正に急いでいます。
同社の Codex コーディング エージェント向けに先週公開されたバグ レポートでは、「Codex SQLite フィードバック ログは年間最大 640 TB の書き込みがあり、SSD の耐久性 #28224 を急速に消費する可能性がある」というタイトルでその影響について警告しています。
「私のマシンでは、約 21 日間稼働した後、メイン SSD に約 37 TB が書き込まれました」と、Apache Flink のプロジェクト管理委員会メンバーである開発者の Rui Fan 氏は書いています。 「プロセス/ファイル レベルのチェックにより、Codex SQLite ログが主要な継続書き込みであることがわかります。
「これはおよそ 640 TB/年と推定されます。 1 TB SSD では、年間約 640 回のフルドライブ書き込みになります。一部のコンシューマー向け SSD の定格は約 600 TBW であるため、これにより、ドライブ全体の保証された書き込み耐久性が 1 年以内にほぼ消費される可能性があります。」
SSD には寿命があり、多くの場合、書き込みテラバイト (TBW) で測定されます。この数値はモデルと容量によって異なります。たとえば、Samsung の 2025 9100 PRO SSD は、1 TB SSD で 600 TBW を約束します。そして、その時点を過ぎると、パフォーマンスが低下し、失敗する可能性が高くなると予想されます。
Codex の問題は、大量のログ データが SSD ストレージに書き込まれるため、ユーザーがハードウェアの寿命を縮めているのではないかと懸念するようになったことです。
別の開発者が Rui Fan に投稿

スレッドには、「Codex がディスク使用量を分析したところ、このバグにより、Samsung 990 2 TB NVMe のドライブ価値が 38.64 ドル発生したとのことです。」と書かれています。
この開発者はその後、Codex が生成したこのバグの全体的なコストの推定値を引用しました。「この回帰により、3 月から 6 月にかけてユーザー全体で 1 桁台前半の SSD 耐久性が失われたと考えられます。」
Codex の経済的影響評価では、SSD への書き込みコストを 1 TB あたり 0.13 ドルと想定しています。
これは、書き込み TB × (SSD 価格 / SSD TBW) の式に基づいています。したがって、1 TB SSD を想定すると、Rui Fan は 37 TB の無駄なストレージに対して 12.33 ドルのコストを負担したと推定されます。 (TBW あたりのコスト = SSD 価格 / SSD 耐久性 = 200 ドル / 600 TBW = 書き込み TB あたり 0.333 ドル。) TBW 定格が高く、より大容量で高価な SSD は、無駄なバイトあたりのコストが低くなります (たとえば、300 ドル / 1200 TBW 2 TB Samsung 9100 PRO SSD の場合、TB あたり 0.25 ドル)。
2025 年 12 月、Codex 開発者は、デフォルトでテレメトリ (法律で禁止されている場合を除く) を Codex CLI に追加する計画を発表しました。
しかし、この問題はローカル診断ログに関係しています。このログは、アプリが昨年デビューした頃に導入され、デフォルトでもオンになっています。ユーザーがフィードバック レポートに含めない限り、ログはデバイス上に残ります。
OpenAI の Codex を使用した過剰な書き込み操作に関する懸念は、数か月前からプロジェクトの GitHub リポジトリで表面化しています。
OpenAI の広報担当者は、同社のエンジニアがこの問題を認識しており、その修正に取り組んでいることを認めました。これは、この問題に対処することを目的とした最近のいくつかのプル リクエストから明らかです。これらのログは、OpenAI エンジニアが問題を診断するのに役立つことを目的としており、この問題は、予想よりもはるかに多くのディスク アクティビティが発生する方法で保存された大量のデータの結果であると聞いています。
Mythos がメモリ リーク「Squidbleed」を発見

クリントン時代以来、検出されずにいた
宇宙軍、記録的な速さのロケットラボ打ち上げを受けて軌道戦争へ（ふり）
O2 が英国の 2G スイッチオフに加わり、2029 年夏に開始予定
このオープンソース CLI を使用して、古い AI オーバーライド アドバイスを嗅ぎ分けます
修正と称するものがリリースされ、会社はある程度の進歩を遂げていますが、ユーザーは引き続き問題を報告しています。
この問題は、アプリサーバーの SQLite ログを TRACE レベルで書き込むという 2 月の作業に遡るようです。これにより、たとえば ERROR レベルよりも詳細なログが出力されます。
おそらく GPT-5.3 を実行している Codex が、この特定の一連のコミットをレビューしたことに注目します。そのため、コードが非常に不適切であったことはさらに驚くべきことになります。 ®
システム
クアルコムは、Dragonfly がデータセンターに着陸するのに遅すぎることはないと主張
ああ、スナップ（ドラゴン）： DC の責任者は、モバイルチップ大手がビットバーンを次の成長市場と見ていると述べた
AI の最新の流行語であるループ エンジニアリングには、依然として人間が必要です
プロンプトを減らし、より多くの自動化には代償が伴います
ZTE と広東省中国電信がクロスベンダー IP ネットワーク シミュレーションのパイロットを推進し、インテリジェントなネットワーク運用への道を開く
パートナーコンテンツ: 95% を超えるデジタルツインの忠実度およびマルチベンダーのコラボレーションを活用して、ネットワーク変更のリスクを排除し、エラーゼロの O&M を達成します。
OpenAI は Broadcom で不安定になる
ハラペーニョは、OpenAI を底辺への競争モデル メーカー以上のものとして描写しようとする最新の発表です
仮想化
VMwars からの教訓 – Broadcom 対 Tesco のスラグフェストについては仮想的なものは何もありません
アジアでの地上戦には決して巻き込まれないでください。また、芸術のモンスターとの契約試合を選択しないでください
Microsoft、AIを活用して恐喝訴訟における2つのマルウェア活動を結びつける
StealCとAmadeyにリンクされた200以上のC2サーバーがシャットダウン
仮想化
Tesco は VMware と Broadcom の撤退に全力を尽くしている

急速な移行のリスク
科学
AI とブレイン コンピューター インターフェイスにより、言葉を話すことのできない ALS 患者がフルタイムで働けるようになります
セキュリティ
Mythos、クリントン時代以来検出されなかったメモリリーク「Squidbleed」を発見
サイバー犯罪
大規模なパスワード窃取攻撃が 75,000 のフォーティネット ファイアウォールを襲う
パーソナルテクノロジー
インドと中国には 29 億人が住んでいますが、第 1 四半期に購入した PC は合わせてわずか 1,300 万台です
プロンプトを減らし、より多くの自動化には代償が伴います
不器用なロギング実装により、コストを考慮せずにデータが無駄に失われる
年次報告書によると、同社はデータセンターの拡張に数十億ドルをつぎ込み、従業員数が年間16万2,000人から14万1,000人に減少したことが明らかになった
信頼できないコード、AI エージェント、または長時間実行されるタスクの実行に適しています
浮遊ビットバーンまたは地下ビットバーンが大流行しているが、数ギガワットのサイトと競合する可能性は低い
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
GIMP 0.54 が Flatpak 形式で復活し、過去から吹き飛ばされる
GTK の代わりに Motif を使用する最初 (そして最後の) リリースによる、ノスタルジックな人々のためのレトロ コンピューティングの楽しみ
Bcachefs は新しい「パフォーマンス リリース」で実験的ステータスを終了します
Rustは増えるが、AIスロップの問題も増える
フランスのデジタル主権推進はマイクロソフトの重力から逃れるのに苦戦している
Nextcloud のロールアウトは、ローカルで制御されるストレージが 1 つのことであることを示しています。ユーザーを Office から解放することはまったく別のことです
CentOS の歴史: 生化学者の Linu がどのように語ったのか

x Hobby プロジェクトはエンタープライズ世界のデフォルトのオペレーティング システムになりました
Red Hat が Windows が「おそらく適切な製品」であると述べた後、コミュニティが団結したとき
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
Project Headroom は多額の費用も節約できる
OpenBSD 7.9 が登場、あらゆる鋭いエッジを誇るダイヤモンドの原石
60 番目のリリースでは、その禁欲的な努力を失うことなく、より多くのコア、遅延休止状態、および基本的な Wi-Fi 6 が追加されています
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

Clumsy logging implementation squirrels away data without regard for cost

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
OpenAI Codex bombards SSDs with needless write operations, costing millions
Clumsy logging implementation squirrels away data without regard for cost
Modern SSDs have a limited number of write cycles before they expire. Now, OpenAI is scrambling to fix a flawed logging implementation that has been shortening the lives of Codex users' solid state drives (SSDs) with excessive data writes and lowering the devices' value by a significant amount of money.
A bug report opened last week for the company's Codex coding agent warns of the consequences in its title: "Codex SQLite feedback logs can write ~640 TB/year and rapidly consume SSD endurance #28224 ."
"On my machine, after about 21 days of uptime, the main SSD has written about 37 TB," wrote developer Rui Fan, a project management committee member of Apache Flink . "Process/file-level checks show Codex SQLite logs are the main continuous writer.
"That extrapolates to roughly 640 TB/year. On a 1 TB SSD, that is about 640 full-drive writes per year. Some consumer SSDs are rated around 600 TBW, so this could consume roughly a full drive's warranted write endurance in less than a year."
SSDs have a limited lifespan, often measured in terabytes written (TBW). This number varies by model and capacity. Samsung's 2025 9100 PRO SSDs, for example, promise 600 TBW for the 1 TB SSD . And after that point, we expect their performance to degrade and failure becomes more likely.
The problem with Codex is that it has been writing so much logging data to SSD storage that users have become concerned they're shortening the life of their hardware.
Another developer posting in Rui Fan's thread remarked, "Codex analyzed the disk usage and says this bug cost me $38.64 in drive value of my Samsung 990 2 TB NVMe."
This dev subsequently cited the Codex-generated estimate of the overall cost of this bug: "This regression plausibly burned low-single-digit millions of dollars of SSD endurance across users during the March-June Window."
Codex's economic impact assessment assumes a cost of $0.13 per TB written to SSDs.
This is based on this formula: TB written × (SSD price / SSD TBW). So given a 1 TB SSD, we estimate that Rui Fan incurred a cost of $12.33 for 37 TB of squandered storage. (Cost per TBW = SSD price / SSD endurance = $200 / 600 TBW = $0.333 per TB written.) A more spacious and more costly SSD with a higher TBW rating would cost less per wasted byte (e.g. $0.25 per TB for a $300 / 1200 TBW 2 TB Samsung 9100 PRO SSD).
In December 2025, Codex devs announced plans to add telemetry by default (except where disallowed by law) to the Codex CLI.
But this issue has to do with local diagnostic logging, which was introduced around the time the app debuted last year and is also on by default. The logs stay on the device unless included by the user in a feedback report.
Concerns about excessive write operations using OpenAI's Codex have been surfacing in the project's GitHub repo for several months .
A spokesperson for OpenAI confirmed that company engineers are aware of the problem and are working to fix it – something evident from several recent pull requests intended to address the problem. We're told that these logs are intended to help OpenAI engineers diagnose issues and that the problem was the result of high-volume data that was being stored in a way that created far more disk activity than anticipated.
Mythos discovers 'Squidbleed,' a memory leak that's gone undetected since Clinton era
Space Force goes to (pretend) orbital war following record-fast Rocket Lab launch
O2 joins UK 2G switch-off with summer 2029 start date
Sniff out stale AI override advice with this open source CLI
While purported fixes have been landing and the company has made some progress , users continue to file problems .
The issue appears to date back to work done in February to write app-server SQLite logs at TRACE level, which emits more verbose logs than, say, ERROR level.
We note that Codex, presumably running GPT-5.3 , reviewed this particular series of commits . That makes it all the more surprising that the code was so ill-conceived. ®
systems
Qualcomm claims it's not too late for Dragonfly to land in datacenters
Oh, Snap(dragon): DC chief says the mobile-chip giant sees bit barns as its next growth market
Loop engineering, latest AI buzzword, still needs humans in the loop
Prompting less and automating more comes with a price
ZTE and China Telecom Guangdong advance cross‑vendor IP network simulation pilots, paving the way for intelligent network operations
PARTNER CONTENT: Leveraging >95% digital twin fidelity and multi-vendor collaboration to eliminate network change risks and achieve zero-error O&M
OpenAI gets chippy with Broadcom
Jalapeño is the latest announcement that attempts to portray OpenAI as more than a race-to-the-bottom model maker
Virtualization
Lessons from the VMwars – nothing virtual about the Broadcom vs Tesco slugfest
Never get involved in a land war in Asia. Also, don't pick a contract fight with a monster of the art
Microsoft uses AI to link two malware operations in racketeering suit
200+ C2 servers linked to StealC and Amadey shut down
virtualization
Tesco is sprinting to quit VMware and Broadcom despite rapid migration risks
science
AI and brain-computer interface allow speechless ALS patient to work a full-time job
Security
Mythos discovers 'Squidbleed,' a memory leak that's gone undetected since Clinton era
CYBER-CRIME
Massive password-stealing attack hits 75k Fortinet firewalls
Personal tech
India and China are home to 2.9 billion people – and together they bought just 13 million PCs in Q1
Prompting less and automating more comes with a price
Clumsy logging implementation squirrels away data without regard for cost
Annual report reveals workforce fell from 162,000 to 141,000 in a year as company pours billions into datacenter expansion
Suitable for running untrusted code, AI agents, or any long-running task
Floating or sub-surface bit barns are all the rage, but unlikely to compete with multi-gigawatt sites
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Blast from the past as GIMP 0.54 is revived in Flatpak form
Retro-computing fun for the nostalgic with first (and last) release to use Motif instead of GTK
Bcachefs exits experimental status in new 'performance release'
More Rust, but more trouble with AI slop, too
France's digital sovereignty push is struggling to escape the Microsoft gravity well
Nextcloud rollout shows locally controlled storage is one thing; getting users off Office is quite another
History of CentOS: How a biochemist's Linux hobby project became the enterprise world's default operating system
When a community came together after Red Hat said Windows was 'probably the right product'
Netflix wiz creates app to slash AI bills, then open sources it
Project Headroom could save you big money, too
OpenBSD 7.9 arrives, a diamond in the rough proud of every sharp edge
Sixtieth release adds more cores, delayed hibernation, and basic Wi-Fi 6 without losing its ascetic streak
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
