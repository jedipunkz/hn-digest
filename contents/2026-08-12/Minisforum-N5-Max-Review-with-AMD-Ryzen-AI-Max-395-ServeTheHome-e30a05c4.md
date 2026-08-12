---
source: "https://www.servethehome.com/minisforum-n5-max-review-with-amd-ryzen-ai-max-395/"
hn_url: "https://news.ycombinator.com/item?id=49272450"
title: "Minisforum N5 Max Review with AMD Ryzen AI Max+ 395 – ServeTheHome"
article_title: "Minisforum N5 Max Review with AMD Ryzen AI Max+ 395 - ServeTheHome"
author: "rbanffy"
captured_at: "2026-08-12T14:14:39Z"
capture_tool: "hn-digest"
hn_id: 49272450
score: 2
comments: 0
posted_at: "2026-08-12T13:54:13Z"
tags:
  - hacker-news
  - translated
---

# Minisforum N5 Max Review with AMD Ryzen AI Max+ 395 – ServeTheHome

- HN: [49272450](https://news.ycombinator.com/item?id=49272450)
- Source: [www.servethehome.com](https://www.servethehome.com/minisforum-n5-max-review-with-amd-ryzen-ai-max-395/)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T13:54:13Z

## Translation

タイトル: AMD Ryzen AI Max+ 395 を使用した Minisforum N5 Max レビュー – ServeTheHome
記事のタイトル: AMD Ryzen AI Max+ 395 を使用した Minisforum N5 Max レビュー - ServeTheHome
説明: Minisforum N5 Max、10GbE、5 ベイ NAS などを 1 つのボックスに組み合わせた 64GB AMD Strix Halo システムをレビューします

記事本文:
フェイスブック
リンクトイン
RSS
TikTok
×
ユーチューブ
フォーラム
ワークステーション
ワークステーションプロセッサ
TrueNAS / FreeNAS NAS サーバーの上位ハードウェア コンポーネント
pfSense アプライアンスの上位ハードウェア コンポーネント
napp-it および Solarish NAS サーバーの上位ハードウェア コンポーネント
Windows Server 2016 Essentials ハードウェアのおすすめ
DIY WordPress ホスティング サーバー ハードウェア ガイド
ストレージの信頼性
レイドカリキュレーター
RAID 信頼性計算ツール |単純な MTTDL モデル
編集および著作権ポリシー
ワークステーション
ワークステーションプロセッサ
TrueNAS / FreeNAS NAS サーバーの上位ハードウェア コンポーネント
pfSense アプライアンスの上位ハードウェア コンポーネント
napp-it および Solarish NAS サーバーの上位ハードウェア コンポーネント
Windows Server 2016 Essentials ハードウェアのおすすめ
DIY WordPress ホスティング サーバー ハードウェア ガイド
AMD Ryzen AI Max+ 395 を使用した Minisforum N5 Max レビュー
Minisforum は、AMD Ryzen AI Max+ 395 と新しいシャーシ設計を組み合わせ続けており、N5 Max はその最も野心的な成果です。このモデルは、16 コア、32 スレッドの Strix Halo プロセッサと 64 GB の LPDDR5X メモリを、小型デスクトップではなく 5 ベイ NAS にラップします。多くの意味で、これは 2 つの Minisforum 系譜を 1 つのボックスにまとめた集大成のように感じられます。私たちは、Minisforum MS-S1 Max を含め、この CPU に何度も会いました。この CPU は 128GB と非常に異なるポート レイアウトを組み合わせています。もう 1 つの系統は Minisforum N5 Pro に由来しており、これにより、密接に関連したシャーシ ファミリが初めて確認されました。 N5 Max は、これら 2 つの中間にあるハイブリッド階層型ストレージ、デュアル 10GbE ネットワーキング、およびフラッグシップ コンピューティングを 1 つのボックスにスタックします。現時点でテストした Ryzen AI Max+ 395 システムの数を考えると、興味深い疑問は、Minisforum が NAS スタイルのワークロードを中心にプラットフォームをどのように調整しているのか、またストレージの焦点をどこで柔軟性と引き換えにしているのかということです。たぶんもっと大きな問題

estion は、「できるからといって、そうすべきですか？」です。
今日はやるべきことがたくさんあるので、ハードウェアから始めましょう。
Minisforum N5 Max 外部ハードウェアの概要
Minisforum では、N5 Max のサイズが 199 x 202.4 x 252.3 mm、5.8 kg と記載されています。
磁石で取り付けられたフロント カバーが外れて、5 つの SATA ドライブ ベイが露出します。
それぞれのツールは、独自のロック ラッチを備えた工具不要のトレイに 3.5 インチ ディスクを搭載します。
この正面図は、N5 Max のストレージ側を定義するドライブ面を示しています。各トレイは取り付けネジなしでまっすぐにスライドし、磁気カバーがアレイにカチッとはまり込みます。メイン シャーシを開けなくても、5 つのベイすべてにアクセスできます。
埋め込み型のリセット コントロールが電源ボタンの横にあります。
クラスタ化されたステータス インジケータにより、システムの動作が一目でわかります。ラベル付きインジケーターは、両方の LAN ポート、および 5 つのドライブ ベイのそれぞれの一般的なステータスをカバーし、実装されたアレイ内でどのディスクがビジーであるかについての推測を排除します。
ベイごとのアクティビティ ライトは、管理ページを開かなくてもどのディスクがアクティブであるかを識別し、個別の LAN1 および LAN2 インジケータが両方のネットワーク インターフェイスをカバーします。これは、5 ドライブ システムのローカル ステータスの有用なレベルです。
前面にはUSB Type-CポートとType-Aポートもあります。この前面 USB4 接続は、専用の背面ビデオ接続とともにディスプレイ出力もサポートします。
重要な接続の大部分は背面パネルにあり、その上に冷却装置があります。
両方の 10GbE ポートは、積み重ねられた青色の Type-A コネクタ、HDMI 出力、および Type-C ポートの横にあるこの後部セクションに取り付けられます。
1 対の大型軸流ファンが背面 I/O エリアの真上に取り付けられ、シャーシのドライブ セクションに空気を送ります。
このボードのストレージ コントローラーは、分解したときに見つかった JMicron JMB585 SATA コントローラーです。
2 つの青い USB Type-A ポート スタック バージョン

背面パネルにあります。
HDMI 出力の周囲には、Minisforum の定格で最大 80 Gbps の 2 つの背面 USB4 v2 Type-C ポートが配置されています。
この詳細図では、背面の 2 つの Type-C ポートとその横の HDMI ジャックが繰り返されています。
DC バレル ジャックの代わりに、後部は 3 ピン AC 電源インレットを利用しており、その隣にあるロック ポートで内部電源に電力を供給します。
上部のシャーシは普通の金属です。
大きなゴム製の脚、ラベル、通気口が底部の大部分を覆っています。
次に、内部ハードウェアの概要のメイン マザーボード セクションに入りたいと思いました。
デルタの GoCool-150 が大活躍、ASRock ラックの NVIDIA VR NVL72 で 150kW の液対空冷却を実現
FMS 2026 で動作する 10M IOPS Kioxia GP1 SSD が示される
AMD、モデル固有のAI推論チップのTaalasを買収へ
バリー・ピアース
2026 年 8 月 11 日 午後 5 時 52 分
AI の不足に関係なく、これは依然としてキラー ラボ サーバーになります
マイク・アンダーソン
2026 年 8 月 12 日午前 7 時 1 分
Minisforum N5 Max の興味深いレビュー。特に Ryzen AI Max 395 では、コンパクトなシステムにこれだけのパフォーマンスを詰め込むことができるのは印象的です。CPU パフォーマンス、AI 機能、小型フォームファクター設計の組み合わせにより、これはパワー ユーザーやホームラボ愛好家にとって興味深いオプションになります。
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
STH ニュースレターに登録してください!
このサイトはスパムを低減するために Akismet を使っています。コメントデータがどのように処理されるかをご覧ください。
STH の最高の機能を毎週受信箱に配信します。 STH から毎週最高の投稿を厳選し、直接お届けします。
オプトインすると、ニュースレターの送信に同意したことになります。当社ではサードパーティのサービスを使用してサブスクリプションを管理しているため、いつでもサブスクリプションを解除できます。

## Original Extract

We review the Minisforum N5 Max, a 64GB AMD Strix Halo system that combines 10GbE, a 5-bay NAS, and more into a single box

Facebook
Linkedin
RSS
TikTok
X
Youtube
Forums
Workstation
Workstation Processors
Top Hardware Components for TrueNAS / FreeNAS NAS Servers
Top Hardware Components for pfSense Appliances
Top Hardware Components for napp-it and Solarish NAS Servers
Top Picks for Windows Server 2016 Essentials Hardware
The DIY WordPress Hosting Server Hardware Guide
Storage Reliability
Raid Calculator
RAID Reliability Calculator | Simple MTTDL Model
Editorial and Copyright Policies
Workstation
Workstation Processors
Top Hardware Components for TrueNAS / FreeNAS NAS Servers
Top Hardware Components for pfSense Appliances
Top Hardware Components for napp-it and Solarish NAS Servers
Top Picks for Windows Server 2016 Essentials Hardware
The DIY WordPress Hosting Server Hardware Guide
Minisforum N5 Max Review with AMD Ryzen AI Max+ 395
Minisforum keeps pairing the AMD Ryzen AI Max+ 395 with fresh chassis designs, and the N5 Max is its most ambitious result yet. This model wraps that 16-core, 32-thread Strix Halo processor and 64GB of LPDDR5X memory into a five-bay NAS rather than a small desktop. In many ways, this feels like the culmination of two Minisforum lineages into one box. We have met the CPU many times, including in the Minisforum MS-S1 Max , which pairs it with 128GB and a very different port layout. The other lineage comes from the Minisforum N5 Pro , which gave us our first look at the closely related chassis family. Between those two, the N5 Max stacks hybrid tiered storage, dual 10GbE networking, and flagship compute in a single box. Given how many Ryzen AI Max+ 395 systems we have tested at this point, the interesting questions are how Minisforum tunes the platform around NAS-style workloads and where it trades flexibility for that storage focus. Maybe the bigger question is: “Just because you can, should you?”
We have a lot to get into today, so let us start with the hardware.
Minisforum N5 Max External Hardware Overview
Minisforum lists the N5 Max at 199 x 202.4 x 252.3mm and 5.8kg.
A magnetically attached front cover pops off to expose the five SATA drive bays.
Each of which takes a 3.5-inch disk on a tool-less tray with its own locking latch.
That front view shows the drive face that defines the storage side of the N5 Max. Each tray slides straight out without mounting screws, and the magnetic cover snaps back over the array. All five bays remain accessible without opening the main chassis.
A recessed reset control sits beside the power button.
Up front, clustered status indicators let a glance report what the system is doing. Labeled indicators cover general status, both LAN ports, and each of the five drive bays, removing guesswork about which disk is busy in a populated array.
Per-bay activity lights identify which disk is active without opening a management page, while separate LAN1 and LAN2 indicators cover both network interfaces. This is a useful level of local status for a five-drive system.
There are also USB Type-C and Type-A ports up front. That front USB4 connection also supports display output alongside the dedicated rear video connection.
Most of the serious connectivity lives on the rear panel with the cooling above it.
Both 10GbE ports mount in this rear section beside the stacked blue Type-A connectors, the HDMI output, and a Type-C port.
A pair of large axial fans mounts directly above the rear I/O area and routes air through the drive section of the chassis.
The storage controller on this board is the JMicron JMB585 SATA controller which we found when taking it apart.
Two blue USB Type-A ports stack vertically on the rear panel.
Around the HDMI output sit the two rear USB4 v2 Type-C ports, rated by Minisforum for up to 80Gbps.
This closer view repeats the pair of rear Type-C ports and the HDMI jack beside them.
Instead of a DC barrel jack, the rear relies on a three-pin AC mains inlet that feeds the internal power supply with a locking port next to it.
Up top, the chassis is a plain metal.
Large rubber feet, labels, and vents cover most of the bottom.
Next, we wanted to get into the main motherboard section in our internal hardware overview..
Delta’s GoCool-150 Goes Big To Enable 150kW Liquid-To-Air Cooling for ASRock Rack’s NVIDIA VR NVL72
A 10M IOPS Kioxia GP1 SSD Shown Running at FMS 2026
AMD to Acquire Taalas for Model Specific AI Inference Chips
Barry Pearce
August 11, 2026 At 5:52 pm
This will still make a killer lab server regardless of the AI shortfalls
Mike Anderson
August 12, 2026 At 7:01 am
A fascinating review of the Minisforum N5 Max. It’s impressive to see how much performance can be packed into a compact system, especially with the Ryzen AI Max 395. The combination of CPU performance, AI capabilities, and small-form-factor design makes this an interesting option for power users and homelab enthusiasts.
Save my name, email, and website in this browser for the next time I comment.
Sign me up for the STH newsletter!
This site uses Akismet to reduce spam. Learn how your comment data is processed.
Get the best of STH delivered weekly to your inbox. We are going to curate a selection of the best posts from STH each week and deliver them directly to you.
By opting-in you agree to have us send you our newsletter. We are using a third party service to manage subscriptions so you can unsubscribe at any time.
