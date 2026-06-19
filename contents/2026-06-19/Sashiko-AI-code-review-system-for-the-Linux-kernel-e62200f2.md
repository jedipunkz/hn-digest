---
source: "https://www.theregister.com/software/2026/03/20/linux-kernel-engineer-introduces-sashiko-code-review-system/5223725"
hn_url: "https://news.ycombinator.com/item?id=48598499"
title: "Sashiko: AI code review system for the Linux kernel"
article_title: "Linux kernel engineer introduces Sashiko code review system"
author: "root-parent"
captured_at: "2026-06-19T16:05:20Z"
capture_tool: "hn-digest"
hn_id: 48598499
score: 4
comments: 0
posted_at: "2026-06-19T13:43:02Z"
tags:
  - hacker-news
  - translated
---

# Sashiko: AI code review system for the Linux kernel

- HN: [48598499](https://news.ycombinator.com/item?id=48598499)
- Source: [www.theregister.com](https://www.theregister.com/software/2026/03/20/linux-kernel-engineer-introduces-sashiko-code-review-system/5223725)
- Score: 4
- Comments: 0
- Posted: 2026-06-19T13:43:02Z

## Translation

タイトル: Sashiko: Linux カーネル向け AI コードレビューシステム
記事タイトル: Linux カーネルエンジニアが Sashiko コードレビューシステムを導入
説明: : メーリングリストで熱狂するビート

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
Sashiko: Linux カーネル用 AI コード レビュー システムが人間が見逃すバグを発見
メーリングリストでビートが熱狂する
AI は、コードの提出ではなく、コード レビュー システムの形で Linux カーネルに導入されます。
Google の Linux カーネル エンジニアである Roman Gushchin 氏が LinkedIn で発表した Sashiko は、バグを見つけてコードをスクリーニングするための Rust で書かれたツールです。
Gushchin 氏は次のように述べています。「私の測定では、Sasako は、(Gemini 3.1 Pro を使用して) 'Fixes:' タグに基づいた完全にフィルターされていない 1,000 件の最近の上流問題のセットに基づいて、バグの 53 パーセントを見つけることができました。53 パーセントはそれほど印象的ではないと言う人もいるかもしれませんが、これらの問題の 100 パーセントは人間のレビュー担当者によって見逃されていました。」
コードの提出における AI の使用については、オープンソース コミュニティで議論の余地がありますが、Sasako のようなツールは、コード レビューの波に対処するメンテナーの負担を軽減するのに何らかの形で役立つ可能性があります。
Systemd 260 が SysV を強制終了し、AI に不正行為をしないよう指示
Linux Foundation、FOSS メンテナーを AI のスロップバグレポートから守る取り組みを開始
Bcachefsの作成者は、カスタムLLMは女性であり「完全に意識がある」と主張しています
Godot のメンテナーは、AI のスロップの提出を「消耗させて士気を低下させる」ことに苦戦している
Sashiko は、メーリング リストからパッチを取り込むことによって機能します。それらを分析し、メンテナと開発者にフィードバックを提供します。著者らによると、「レビューの質は高い…誤検知率を測定するのは難しいが、限られた手作業によるレビューに基づくと、それは20パーセントの範囲内に十分収まり、その大部分はグレーゾーンだ」という。
著者らはプライバシーとコード共有の側面について率直に述べています。 Sashiko は、LLM プロバイダーにデータとコードを送信します。

用に設定されています。これは Gemini Pro 3.1 で最もテストされていますが、Claude や他の LLM でも動作するはずです。ただし、運営するにはコストがかかります。 Linux カーネル メーリング リストの場合、Google が費用を負担しています。
Gushchin 氏は、「Google 社内でしばらくの間、これを使用してきましたが、多くの実際の問題を発見するのに役立ちました。」と述べました。
Sashiko は Linux Foundation に属しており、コードの提出よりも手間がかからないエージェント AI のアプリケーションである便利なツールのように見えます。 ®
セキュリティ
研究者は、A12 および A13 iPhone 向けの checkm8 スタイルの BootROM エクスプロイトをドロップしました。
影響を受ける iPhone の所有者は、今すぐパッチのチェックを停止できます。この SecureROM バグの修正は新しい携帯端末に含まれています
Tensordyne が Nvidia に勝つために対数計算に大きな賭けをする
対数を加算するだけで済むのに、大量の計算を必要とする乗算が必要になる人はいない
ZTE と広東省中国電信がクロスベンダー IP ネットワーク シミュレーションのパイロットを推進し、インテリジェントなネットワーク運用への道を開く
パートナーコンテンツ: 95% を超えるデジタルツインの忠実度およびマルチベンダーのコラボレーションを活用して、ネットワーク変更のリスクを排除し、エラーゼロの O&M を達成します。
Bcachefs は新しい「パフォーマンス リリース」で実験的ステータスを終了します
Rustは増えるが、AIスロップの問題も増える
PAAS と IAAS
Graviton 5 は感動的ですが、すべての神聖なるものへの愛を込めて、これを「AI チップ」と呼ぶのはやめてください
AWS は口よりもチップ工場の運営のほうが上手い
ロボタクシーが高速道路建設区域の標識を見逃し続けたため、ウェイモがブレーキを踏む
通行止め警報を無視したり、通行止めを示すコーン間を走行したりして約4,000台がリコール
セキュリティ
研究者によると、脱獄ではなく単に「このコードを修正してください」というプロンプトが表示されただけで、連邦当局は寓話5に激怒したとのこと
オンプレミス
アマゾンは昨年自社のビット納屋で最大25億ガロンの水を使用している
科学
AIとB

レインコンピューターインターフェースにより、言葉を失ったALS患者がフルタイムで働けるようになります
公共部門
キャピタは公務員年金制度修正の期限を過ぎて航行しようとしている
仮想化
Tesco は急速な移行リスクにもかかわらず、VMware と Broadcom からの撤退に全力を注いでいます
対数を加算するだけで済むのに、大量の計算を必要とする乗算が必要になる人はいない
Vercel 経由で AWS を間接的に使用する場合のコストプレミアムは、コンピューティング リソースのより効率的な使用により軽減される、と CTO は主張
活動家らは、テクノロジーは使用が計画されている境界で子供と大人を確実に区別することができないと主張している
サンフランシスコがホスティング会社の Localhost カンファレンスの主催者となる
基礎となるテクノロジーは本物です...そして、会社が言及しなかったパートナーから借用したものです
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
Bcachefs は新しい「パフォーマンス リリース」で実験的ステータスを終了します
Rustは増えるが、AIスロップの問題も増える
フランスのデジタル主権推進はマイクロソフトの重力から逃れるのに苦戦している
Nextcloud のロールアウトは、ローカルで制御されるストレージが 1 つのことであることを示しています。ユーザーを Office から解放することはまったく別のことです
CentOS の歴史: 生化学者の Linux 趣味プロジェクトがどのようにして企業世界のデフォルトのオペレーティング システムになったのか
Red Hat が Windows が「おそらく適切な製品」であると述べた後、コミュニティが団結したとき
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
プロジェクトのヘッドルームがあなたを救ってくれる

大金も
OpenBSD 7.9 が登場、あらゆる鋭いエッジを誇るダイヤモンドの原石
60 番目のリリースでは、その禁欲的な努力を失うことなく、より多くのコア、遅延休止状態、および基本的な Wi-Fi 6 が追加されています
Fedora: Microsoft は全力で取り組んでいるが、Deepin は見捨てられている
Red Hat の無料ディストリビューションはデスクトップを失うが、重要な新しい友人を作る
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

: Beats getting roasted on the mailing list

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Sashiko: AI code review system for the Linux kernel spots bugs humans miss
Beats getting roasted on the mailing list
AI is coming to the Linux kernel in the form of a code review system - not code submissions.
Announced on LinkedIn by Roman Gushchin, a Linux kernel engineer at Google, Sashiko is a tool written in Rust for spotting bugs and screening code.
Gushchin said: "In my measurement, Sashiko was able to find 53 percent of bugs based on a completely unfiltered set of 1,000 recent upstream issues based on 'Fixes:' tags (using Gemini 3.1 Pro). Some might say that 53 percent is not that impressive, but 100 percent of these issues were missed by human reviewers."
The use of AI in code submissions is controversial in the open source community, however, a tool like Sashiko could go some way toward easing the burden on maintainers dealing with a wave of code reviews.
Systemd 260 kills SysV, tells AI not to misbehave
Linux Foundation kicks off effort to shield FOSS maintainers from AI slop bug reports
Bcachefs creator insists his custom LLM is female and 'fully conscious'
Godot maintainers struggle with 'draining and demoralizing' AI slop submissions
Sashiko works by ingesting patches from a mailing list. It analyzes them then gives feedback to the maintainers and developers. According to its authors, "the quality of reviews is high... the rate of false positives is harder to measure, but based on limited manual reviews it's well within 20 percent range, and the majority of it is a gray zone."
The authors are upfront about the privacy and code-sharing aspects. Sashiko sends data and code to whatever LLM provider it has been configured for. It has been most tested with Gemini Pro 3.1, but should work with Claude and other LLMs. However, there are costs involved in running it. In the case of the Linux Kernel Mailing List, Google is footing the bill.
Gushchin said: "We've been using it internally at Google for some time, and it helped to discover a large number of real issues."
Sashiko belongs to the Linux Foundation and looks like a useful tool – an application of agentic AI that may provoke less handwringing than code submissions. ®
Security
Researchers drop checkm8-style BootROM exploit for A12 and A13 iPhones
Owners of affected iPhones can stop checking for patches now: the fix for this SecureROM bug comes in a new handset
Tensordyne makes a big bet on log math to beat Nvidia
Who needs compute-hungry multiplications when you can just add logarithms
ZTE and China Telecom Guangdong advance cross‑vendor IP network simulation pilots, paving the way for intelligent network operations
PARTNER CONTENT: Leveraging >95% digital twin fidelity and multi-vendor collaboration to eliminate network change risks and achieve zero-error O&M
Bcachefs exits experimental status in new 'performance release'
More Rust, but more trouble with AI slop, too
PAAS AND IAAS
Graviton 5 impresses, but please, for the love of all that's holy, stop calling them 'AI chips'
AWS better at running chip fabs than their mouths
Waymo hits the brakes after robotaxis keep missing the signs for freeway construction zones
Nearly 4,000 vehicles recalled for driving past closure warnings and between cones marking shut lanes
security
Feds freaked over Fable 5 after simple 'fix this code' prompt, not jailbreak, says researcher
ON-PREM
Amazon owns up to using 2.5bn gallons of H2O in its bit barns last year
science
AI and brain-computer interface allow speechless ALS patient to work a full-time job
PUBLIC SECTOR
Capita is about to sail past deadline to fix civil service pensions scheme
virtualization
Tesco is sprinting to quit VMware and Broadcom despite rapid migration risks
Who needs compute-hungry multiplications when you can just add logarithms
Cost premium of using AWS indirectly via Vercel is mitigated by more efficient use of compute resources, CTO claims
Campaigners say tech is unable to reliably distinguish between kids and adults at the boundary where use is planned
San Francisco plays host to hosting company's Localhost conference
The underlying technology is real...and borrowed from a partner the company failed to mention
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
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
Fedora: Microsoft is all aboard, but Deepin is dumped
Red Hat’s free distro loses a desktop, but makes an important new friend
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
