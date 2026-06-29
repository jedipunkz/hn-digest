---
source: "https://www.servethehome.com/taking-an-up-close-look-at-the-supermicro-gb300-super-ai-station/"
hn_url: "https://news.ycombinator.com/item?id=48722211"
title: "Taking an Up-Close Look at the Supermicro GB300 Super AI Station – ServeTheHome"
article_title: "Taking an Up-Close Look at the Supermicro GB300 Super AI Station - ServeTheHome"
author: "rbanffy"
captured_at: "2026-06-29T17:55:50Z"
capture_tool: "hn-digest"
hn_id: 48722211
score: 1
comments: 0
posted_at: "2026-06-29T17:27:03Z"
tags:
  - hacker-news
  - translated
---

# Taking an Up-Close Look at the Supermicro GB300 Super AI Station – ServeTheHome

- HN: [48722211](https://news.ycombinator.com/item?id=48722211)
- Source: [www.servethehome.com](https://www.servethehome.com/taking-an-up-close-look-at-the-supermicro-gb300-super-ai-station/)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T17:27:03Z

## Translation

タイトル: Supermicro GB300 Super AI ステーションを詳しく見る – ServeTheHome
記事タイトル: Supermicro GB300 Super AI ステーションを詳しく見る - ServeTheHome
説明: NVIDIA DGX Station システムが現在出荷されていることから、Supermicro は自社の Super AI Station を Computex に展示し、2,080 億トランジスタの栄光をすべて備えた GB300 システムを披露しました。

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
Supermicro GB300 Super AI ステーションを間近で見てみる
NVIDIA が Computex 2026 で複数の新しいデスクトップ システム設計を発表したため、NVIDIA のパートナーは当然のことながら、現在および今後の製品を披露したいと考えていました。ほとんどのデスクトップ ユーザーにとって、ショーのハイライトはおそらく、NVIDIA DGX Station システムに遭遇することでしょう。このシステムは、1 台あたり約 125,000 ドル (Newegg アフィリエイト リンク経由) で、入手可能な最も高価なデスクトップ コンピューターとして現在進行中であり、NVIDIA のフラッグシップ GB300 アクセラレータをサーバーよりも小さいもので入手する唯一の方法です。
幸運なことに、私たちは Supermicro Computex 2026 ブースでそれらのシステムの 1 つに偶然遭遇しました。そこでは、NVIDIA のパートナーが DGX Station システムである Super AI Station をフル展示していました。
Super AI ステーションは、Supermicro の DGX ステーション製品です。そして、自社の製品がどれほど興味深いものであるかを理解してくれた同社に敬意を表します。彼らはシステムをオープンにしていなかったが、次善の策として、

参加者が内部をよく見ることができるように、側面にxiglassパネルが設置されています。
GB10 システムと同様に、NVIDIA は DGX Station システムでも OEM を比較的短い時間で拘束してきました。 NVIDIA は重要なマザーボードとプロセッサーを提供するため、OEM はポートの可用性や電力などについて NVIDIA の仕様を満たしながら、それを中心としたシステムを設計することができます。最終的には、DGX Station システムのコア コンポーネントとレイアウトはすべて Supermicro のシステムと非常によく似たものになります。
その中心となるのは、72 コアの Grace CPU ダイと Blackwell GPU を組み合わせた Grace Blackwell GB300 チップです。システムの 1600 ワット TDP のほとんどがそのチップに送られるため、Supermicro はここで液体冷却を選択し、すべての重要なコンポーネントをコールド プレートで覆い、ラジエーター ループを実行して熱を運び去ります。
これには、Grace CPU に 496 GB のメモリを提供する 4 つの LPDDR5X SOCAMM が含まれており、これらは技術的には取り外し可能ですが、簡単にアクセスできるわけではありません。 Blackwell GPU に関しては、NVIDIA は定格の低い部品を使用しています。 DGX Station システムは、チップ上の 8 つの HBM3e スタックのうち 7 つだけが有効になった状態で出荷され、合計 252 GB の GPU メモリと 7.1 TB/秒の合計メモリ帯域幅を実現します。
ストレージ オプションに関しては、すべての DGX Station システムには 4 つの PCIe Gen5 x4 M.2 2280 スロットが付属しており、ショーケース システムの Supermicro には 480GB Micron 7450 Gen4 x4 ドライブが搭載されています。
一方、NVIDIA RTX PRO ビデオ カードもインストールされており、これは NVIDIA が許可するオプションのアップグレードです。 Blackwell Ultra GPU が同梱されていますが、BMC ビデオ出力よりも優れたものが必要な場合は、ワークステーションに別の GPU が必要です。 GB300 はコンピューティングに最適化された設計のためグラフィックス機能を備えていないため、BMC ビデオを超えるグラフィックスを有効にするには RTX PRO カードをインストールする必要があります。これにより、さらに 2 つの PCIe が残ります

x16 (x8 電気) スロットはさらなるアップグレードに利用可能です。
横から見ると、小さなヒートシンクで覆われたシステムの 400GbE QSFP ポートが確認できます。 GB300 システム自体と並んで、NVIDIA の ConnectX-8 搭載イーサネットは、DGX Station システムのもう 1 つの主要な機能であり、2 つの 400Gb ポートにより、複数のシステムを一緒にネットワーク化するため、またはステーションとローカル ネットワーク上の他のマシン間の高速バックホールに十分な帯域幅が提供されます。
GB300 とマザーボードをシステムの背面に配置することで、Supermicro はシステムの大型ラジエーターを前面に配置しました。 3 つのファンがラジエーターから空気を取り込み、システムの残りの部分に空気を押し込み、最終的には背面から排出されます。
この前後構成は、複数のレベルで意図的に選択されたものです。これにより、ハイエンド ワークステーションの通常のセットアップに加えて、Super AI ステーションがラックマウント セットアップにも適していることが保証されます。デスクトップ コンピュータには、必要に応じてラックマウント キットを装備し、5U サーバーとして取り付けることができます。これは、従来の HGX ベースのサーバーと比べてスペース効率の良いオプションではありませんが、単一の GB300 をラックに取り付ける簡単な方法です。これは Supermicro の差別化要因でもあり、人々にとって非常に理にかなっているかもしれません。誰もが自分のデスクで 1.6kW システムを実行したいと思うわけではありません。
スーパーAIステーションの前を見ることになったら、これが目に入るでしょう。オールブラックのシステムの前面は、TDP が大きいため、主に換気用です。それ以外の場合は、前面右隅に向かって、システムの前面からアクセス可能な 2 つの 5Gbps USB 3 ポートがあります。背面には 4 つの 10Gbps USB 3 ポートがありますが、システムには USB-C ポートはありません。
NVIDIA の DGX ステーションは人気です

アイテムと複数の方法で。 GB300 システムをデスクトップ上に置きながらローカルでの開発とテストを行うことを検討している人にとって、DGX Station はそれを行うための方法です。 Supermicro の Super AI Station で使用される GB300 は依然として全負荷時に 1600 ワットを消費するため、専用の電源回路が利用可能であることを確認してください。良くも悪くも、彼らに匹敵するものはありません。
NVIDIA のカスタム Ubuntu ディストリビューションである DGX OS がプリロードされた現在の DGX Station システムに加えて、今年後半に OEM が DGX Station for Windows ブランドで Windows バージョンの販売を開始するときに、このシステムも登場する予定です。
Qualcomm Investor Day 2026 データセンターの発表 CPU、AI アクセラレーターなど
Broadcom を活用した OpenAI Jalapeño インテリジェンス プラットフォームを披露
MiTAC Computex 2026 ブース ツアー: ダイヤモンド冷却、52U ラックなど
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
STH ニュースレターに登録してください!
このサイトはスパムを低減するために Akismet を使っています。コメントデータがどのように処理されるかをご覧ください。
STH の最高の機能を毎週受信箱に配信します。 STH から毎週最高の投稿を厳選し、直接お届けします。
オプトインすると、ニュースレターの送信に同意したことになります。当社ではサードパーティのサービスを使用してサブスクリプションを管理しているため、いつでもサブスクリプションを解除できます。

## Original Extract

With NVIDIA DGX Station systems now shipping, Supermicro had their Super AI Station on display a Computex, showing off the GB300 system in all of its 208 billion transistor glory

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
Taking an Up-Close Look at the Supermicro GB300 Super AI Station
With NVIDIA announcing multiple new desktop system designs at Computex 2026, NVIDIA’s partners were understandably eager to show off their current and forthcoming wares. For most desktop users, the highlight of the show is arguably running into an NVIDIA DGX Station system, which at around $125,000 each (via Newegg Affiliate link) is in the running as the most expensive desktop computer available, and the only way to get NVIDIA’s flagship GB300 accelerator in something smaller than a server.
As luck would have it, we happened to run into one of those systems over at the Supermicro Computex 2026 booth, where NVIDIA’s partner had its DGX Station system, the Super AI Station, on full display.
The Super AI Station is Supermicro’s DGX Station offering. And kudos to the company for understanding just how interesting a product they had on their hands. While they did not have the system open, they did the next best thing by putting a plexiglass panel on its side so that attendees could get a good look inside.
As with GB10 systems, NVIDIA has kept OEMs on a relatively short leash with DGX Station systems. NVIDIA provides the critical motherboard and processors, leaving OEMs to design a system around it while meeting NVIDIA’s specifications for things such as port availability and power. The end result is that the core components and layout of a DGX Station system are all going to look very similar to Supermicro’s system.
At its heart is the Grace Blackwell GB300 chip, which combines a 72-core Grace CPU die with a Blackwell GPU. With most of the system’s 1600 Watt TDP going to that chip, Supermicro has opted for liquid cooling here, covering all critical components with cold plates and running a radiator loop to carry away the heat.
This includes the four LPDDR5X SOCAMMs that provide the 496GB of memory for the Grace CPU, so while those are technically removable, they are not easily accessible. As for the Blackwell GPU, NVIDIA is using a de-rated part here. DGX Station systems ship with only 7 of the 8 HBM3e stacks on the chip enabled, for a total of 252GB of GPU memory and a total memory bandwidth of 7.1TB/second.
As for storage options, all DGX Station systems come with four PCIe Gen5 x4 M.2 2280 slots, which, for its showcase system, Supermicro filled with 480GB Micron 7450 Gen4 x4 drives.
Meanwhile, they also had an NVIDIA RTX PRO video card installed, an optional upgrade that NVIDIA allows. Despite shipping with the Blackwell Ultra GPU, you need another GPU for a workstation if you want something better than BMC video output. The GB300 is not graphics-capable due to its compute-optimized design, so an RTX PRO card must be installed to enable graphics beyond BMC video. This leaves two more PCIe x16 (x8 electrical) slots available for further upgrades.
From the side angle, you can also just make out the system’s 400GbE QSFP ports, which are covered with a small heatsink. Along with the GB300 system itself, NVIDIA’s ConnectX-8-powered Ethernet is the other marquee feature of the DGX Station system, with the two 400Gb ports providing ample bandwidth for networking multiple systems together or for a fast backhaul between the Station and other machines on the local network.
With the GB300 and motherboard at the rear of the system, Supermicro has placed the system’s sizable radiator towards the front. Three fans pull air through that radiator, push it through the rest of the system, and eventually out the rear.
This front-to-back configuration is an intentional choice on multiple levels. Besides being the customary setup for a high-end workstation, this ensures that the Super AI Station is suitable in a rackmount setup. The desktop computer can, if desired, be outfitted with a rackmount kit, allowing it to be mounted as a 5U server. This is not going to be a space-efficient option versus a traditional HGX-based server, but it is an easy way to get a single GB300 installed in a rack. It is also a differentiator for Supermicro that may make a lot of sense to people. Not everyone will want a 1.6kW system running at their desk.
If you ever do find yourself looking at the front of a Super AI Station, this is what you will be looking at. The front of the all-black system is primarily for ventilation, given its significant TDP. Otherwise, up towards the front-right corner, you will find the system’s two front-accessible 5Gbps USB 3 ports. There are four 10Gbps USB 3 ports on the rear, but no USB-C ports to speak of on the system.
NVIDIA’s DGX Stations are a hot item and in more ways than one. For anyone looking to do local development and testing against a GB300 system while wanting said system on their desktop, the DGX Station is the way to do it. Just make sure you have a dedicated power circuit available, as the GB300 used in Supermicro’s Super AI Station still draws 1600 Watts at full load. For better or worse, there is nothing quite like them.
Along with the current DGX Station systems that come pre-loaded with NVIDIA’s custom Ubuntu distro, DGX OS, we will also see the systems later this year when OEMs begin selling Windows versions under the DGX Station for Windows branding.
Qualcomm Investor Day 2026 Data Center Announcements CPUs, AI Accelerators, and More
OpenAI Jalapeño Intelligence Platform Shown Powered by Broadcom
MiTAC Computex 2026 Booth Tour: Diamond Cooling, 52U Racks, and More
Save my name, email, and website in this browser for the next time I comment.
Sign me up for the STH newsletter!
This site uses Akismet to reduce spam. Learn how your comment data is processed.
Get the best of STH delivered weekly to your inbox. We are going to curate a selection of the best posts from STH each week and deliver them directly to you.
By opting-in you agree to have us send you our newsletter. We are using a third party service to manage subscriptions so you can unsubscribe at any time.
