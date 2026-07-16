---
source: "https://blog.rezero.org/byo-nos"
hn_url: "https://news.ycombinator.com/item?id=48940312"
title: "BYO NOS: Building a maintainable, SRE-friendly switch"
article_title: "BYO NOS: Building a maintainable, SRE-friendly switch — igloo's blog"
author: "mrngm"
captured_at: "2026-07-16T21:57:04Z"
capture_tool: "hn-digest"
hn_id: 48940312
score: 1
comments: 0
posted_at: "2026-07-16T21:11:25Z"
tags:
  - hacker-news
  - translated
---

# BYO NOS: Building a maintainable, SRE-friendly switch

- HN: [48940312](https://news.ycombinator.com/item?id=48940312)
- Source: [blog.rezero.org](https://blog.rezero.org/byo-nos)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T21:11:25Z

## Translation

タイトル: BYO NOS: 保守可能で SRE フレンドリーなスイッチの構築
記事のタイトル: BYO NOS: 保守可能で SRE フレンドリーなスイッチの構築 — igloo のブログ
説明: Mellanox Spectrum シリコンと mlxsw を使用して、SRE フレンドリーで長期保守可能なネットワーク スイッチを開発します。

記事本文:
BYO NOS: 保守可能で SRE フレンドリーなスイッチの構築
パケットがあるので移動する必要があります。
ここ数年、私は Arista 7050QX2-32S をコア スイッチとして使用してきました。それは見事に機能する有能なデバイスです。残念ながら、この製品は 2016 年に発売され、2025 年に製品寿命を迎えたため、少々長くなってきています。
良い代替品を見つけるのは困難です。その理由を理解するには、まず使用しているハードウェアを理解する必要があります。
ハードウェアの観点から見ると、ネットワーク スイッチのコントロール プレーンは通常は複雑ではありません。ざっくり言うと、スイッチ ファブリックは、当時のかなり低電力のサーバーの側面に移植された、非常に大きくて高価なシリコンの塊です。
Arista のデバイスは、カーネル モジュールとユーザー空間拡張機能を備えた Linux ベース上に構築されています。従来の CLI も利用できますが、ベース Linux システム内の一部の操作は ASIC にコピーされます。たとえば、IPv4 または IPv6 ルートがカーネルに挿入されると、Arista によってカーネル ソース ルート ( show Route の「K」ルート起点を持つ) として解釈され、ハードウェアにオフロードされます。その結果、Arista ハードウェア上で代替ルーティング デーモン (BIRD など) を実行することは非常に簡単です。
初めてハードウェアを入手したとき、私はすぐに Arista の寛大な機能を利用しました。BIRD、非 SNMP 提供のメトリクス、および他のサーバーと同様にシステムの健全性を監視する機能を中心にかなりのインフラストラクチャを構築しました。
残念ながら、Arista のハードウェアのオープンで一見アクセスしやすい性質は、ベンダー ロックインのように機能する可能性があります。私が見つけた他のベンダーは、自社のハードウェア上でサードパーティ製の既製デーモンをサポートすることを中心にこれほど広範なエコシステムを構築していないからです。では、他にはどのような選択肢があるのでしょうか?
数年前から

ああ、ハイパースケーラーは独自のスイッチを構築し始めました。これらのデバイスには (通常) 特注の ASIC が付属しておらず、代わりに既存のシリコンに対して実行される社内ソフトウェア実装が選択されています。これにより、NOS の構築と保守の負担がハイパースケーラーに移されましたが、同時に NOS 内で実行されている基盤となるソフトウェアを制御できるようになりました。ベンダーが修正を拒否した長年のソフトウェアのバグの時代は終わり、代わりに、スイッチの構成を汎用のマイクロサービスと対話するように感じる特注のコントロール プレーンに置き換えられました。
この傾向は社内にとどまりませんでした。特に SONiC や DENT など、非ベンダー開発の一般に入手可能な NOS が出現し始めました。
SONiC は元々 Microsoft によって開発され、ASIC 固有の、通常はブラック ボックス BLOB を利用して、実行される個々のベンダー/SKU をサポートします。その結果、どのベンダー (またはベンダー ライセンシー) がマジック バイナリの開発に時間を費やしたかによって、サポートが大きく制限されます。私の既存のハードウェアもサポートされていますが、大きな問題が 1 つあります。多くのオープンソース プロジェクトと同様、さまざまな機能のサポートは個々のニーズに大きく依存しています。 Microsoft がレイヤー 3 機能を必要としていたのは明らかですが、レイヤー 2 機能にはむらがあり、私が使用したすべてのハードウェアで非同期が発生する傾向がありました。さらに、カーネルからの IP ルート挿入などの一部の機能は使用できません。
SONiC とは対照的に、DENT はプロジェクト固有の抽象化レイヤーを利用しません。代わりに、switchdev と呼ばれる、ベンダーが提供する一般化されたカーネル レベルの抽象化レイヤーの上で、NOS のようなインターフェイスとして機能します。
Switchdev は、使い慣れた Linux 機能を ASIC にオフロードします。 Aristaと同様にIPv4/IPv6ルーをコピーする機能が存在します

テス。 Arista とは異なり、ACL やインターフェイス メトリックの収集などの追加の共通機能は、標準ツールを介して実行できます。 DENT は、機能的には、Linux ホストを管理するために構築された、事前に構築された NOS のようなツール セットです。
DENT 自体は私が求めていたものとまったく同じツールを提供していませんでしたが、switchdev が機会を提供してくれました。
DENT (拡張子 switchdev) がサポートするハードウェア リストは比較的少ないです。リストを見てみると、何らかの形で switchdev をサポートしている 3 つの主要な ASIC ファミリ、Marvell の Alleycat3、Marvell の Aldrin 2、Mellanox の Spectrum がわかります。少し調べた結果、Spectrum ベースの switchdev をサポートするシャーシである Mellanox SN2010 を比較的安価で入手することができました。
それを開くと、非常に単純なレイアウトが表示されます。
「ASIC を搭載した小型 PC」というコンセプトは非常に明白です。コンピューター自体は小さなドーターボードです。
スイッチング ハードウェアに加えて、シャーシには 2 つの非シールド電源と 4 つの内部ファンが含まれています。空気の流れを逆方向にしたい場合は、3 本のネジを使って数分でファンを回転させることができます。完全に内部構成の場合は悪くありません。
フロント パネルの SFP/QSFP ポートとは別に、デバイスには 3 つの管理ポートがあります。RS232 シリアル ポート、RJ45 Intel ギガビット インターフェイス、および USB ポートです。シリアル ポートはデフォルトのボー レート 115200 で出力しますが、BIOS は USB 経由のキーボード入力のみを受け入れることができます。
起動して F7 を押し、BIOS パスワード「admin」を入力すると、かなり標準的な AMI BIOS が表示されます。
ここから、デフォルトの起動デバイスを USB に設定し、キーボードを Debian インストーラーに交換して、再起動します。 GRUB に到達すると、シリアル経由のキーボード入力が機能し始め、(「console=tty0 console=ttyS0,1」を追加することで) シリアル対応の Debian インストーラーを起動できるようになります。

15200' を GRUB のブート オプションに設定し、Debian を内蔵 SSD にインストールします。通常の Debian シェルにアクセスすると、インターフェースが欠落していることに気づきます。
igloo@core02:~$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue 状態 UNKNOWN グループのデフォルト qlen 1000
リンク/ループバック 00:00:00:00:00:00 brd 00:00:00:00:00:00
inet 127.0.0.1/8 スコープ ホスト lo
valid_lft 永久にpreferred_lft 永久に
inet6 ::1/128 スコープ ホスト noprefixroute
valid_lft 永久にpreferred_lft 永久に
2: enp0s20f0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP グループのデフォルト qlen 1000
リンク/イーサ XX:XX:XX:XX:XX:XX brd ff:ff:ff:ff:ff:ff
inet6 fe80::XXXX:XXXX:XXXX:XXXX/64 スコープ リンク プロト kernel_ll
valid_lft 永久にpreferred_lft 永久に
Mellanox の switchdev 実装である mlxsw はアップストリームされていますが、ほとんどのカーネル ビルドではデフォルトで有効になっていません。これを修正するには、カーネル オプションを有効にし、カーネルを構築して、デバイスにカーネルをインストールする必要があります。幸いなことに、mlxsw のドキュメントには正確なリストが記載されています。
CONFIG_NET_IPIP=m
CONFIG_NET_IPGRE_DEMUX=m
CONFIG_NET_IPGRE=m
CONFIG_IPV6_GRE=m
CONFIG_IP_MROUTE_MULTIPLE_TABLES=y
CONFIG_IP_MULTIPLE_TABLES=y
CONFIG_IPV6_MULTIPLE_TABLES=y
CONFIG_BRIDGE=m
CONFIG_VLAN_8021Q=m
CONFIG_BRIDGE_VLAN_FILTERING=y
CONFIG_BRIDGE_IGMP_SNOOPING=y
CONFIG_NET_SWITCHDEV=y
CONFIG_NET_DEVLINK=y
CONFIG_MLXFW=m
CONFIG_MLXSW_CORE=m
CONFIG_MLXSW_CORE_HWMON=y
CONFIG_MLXSW_CORE_THERMAL=y
CONFIG_MLXSW_PCI=m
CONFIG_MLXSW_I2C=m
CONFIG_MLXSW_MINIMAL=y
CONFIG_MLXSW_SWITCHX2=m
CONFIG_MLXSW_SPECTRUM=m
CONFIG_MLXSW_SPECTRUM_DCB=y
CONFIG_LEDS_MLXCPLD=m
CONFIG_NET_SCH_PRIO=m
CONFIG_NET_SCH_RED=m
CONFIG_NET_SCH_INGRESS=m
CONFIG_NET_CLS=y
CONFIG_NET_CLS_ACT=y
CONFIG_NET_ACT_MIRRED=m
CONFIG_NET_CLS_MATCHALL=m
CONFIG_NET_CLS_FLOWER=m
CONFIG_NET_ACT_GACT=m
CONFIG_NET_ACT_SAMPLE=m
C

ONFIG_NET_ACT_VLAN=m
CONFIG_NET_L3_MASTER_DEV=y
CONFIG_NET_VRF=m
デフォルトのインターフェース名は簡単には解析できません。 Pin van Pelt の Spectrum ハードウェアに関するガイドを参照すると、udev ルールで問題を解決できます。
igloo@core02:~$ cat /etc/udev/rules.d/10-local.rules
SUBSYSTEM=="net"、ACTION=="add"、DRIVERS=="mlxsw_spectrum*"、NAME="sw$attr{phys_port_name}"
再起動後、ほぼ標準の Debian に移行し、新しいルーターを構築する準備が整います。
単一の QSFP インターフェイスを 4 つの SFP+ インターフェイスにブレークアウトし、ネットワーク内のホストで使用できるブレークアウト ケーブルを多数持っています。新しいケーブルを購入したくないので、devlink 経由でブレークアウト ケーブル用に個々のスイッチポートを再構成できます。
/sbin/devlink ポート分割 $(/sbin/devlink port show | grep swp22 | Cut -d " " -f 1 | Cut -d ":" -f 1,2,3) count 4
ethtool 経由でポートが正常にブレイクアウトされたことを確認できます。
igloo@core02:~$ sudo ethtool swp22s0
swp22s0 の設定:
サポートされているポート: [ ファイバー ]
サポートされているリンク モード: 1000baseKX/フル
10000baseKR/フル
25000baseCR/フル
25000baseSR/フル
サポートされているポーズ フレームの使用: 対称受信専用
オートネゴシエーションのサポート: はい
サポートされている FEC モード: 報告されていません
アドバタイズされたリンク モード: 1000baseKX/フル
10000baseKR/フル
25000baseCR/フル
25000baseSR/フル
アドバタイズされた一時停止フレームの使用: いいえ
アドバタイズされた自動ネゴシエーション: はい
アドバタイズされた FEC モード: 報告されません
速度: 10000Mb/秒
レーン: 1
二重: 全二重
オートネゴシエーション: オン
ポート: 直接接続銅線
フィアド: 0
トランシーバー: 内蔵
リンクが検出されました: はい
L2 および L3 インターフェイス
従来の Debian ホストでは、ネットワーク インターフェイスを管理するツールを決定する際に、選択肢がありません。私はこれまで systemd-networkd を使用してきたので、使い慣れたものに手を伸ばしました。
systemd-networkd の構成は冗長です。あらゆるオペラに対して

大規模な場合は、テンプレートベースの構成ビルド システムを構築するか、大量の構成ファイルをすぐに作成できる systemd-networkd 以外のものを選択することをお勧めします。
igloo@core02:/etc/systemd/network$ ls -la |トイレ -l
83
とはいえ、基本的な構成は比較的簡単です。 systemd-networkd がインストールされたら、ブリッジを定義する必要があります。
igloo@core02:/etc/systemd/network$ cat 40-br0.netdev
[ネット開発]
名前=br0
種類=ブリッジ
# 注: ASIC リソースの理由により、SN2010 上のすべての MAC プレフィックスは同じ 38 ビットである必要があります。ここのアドレスは、使用されるプレフィックスを導出するために使用されます。
MACアドレス=XX:XX:XX:XX:XX:XX
【橋】
VLANフィルタリング=1
VLANプロトコル=802.1q
デフォルトPVID=なし
優先度=8
STP=はい
# OSPF の問題を回避するためにマルチキャスト スヌーピングを無効にする / 動作させるために追加のソフトウェアを実行する必要がある
マルチキャストスヌーピング=false
各 VLAN のデバイス:
igloo@core02:/etc/systemd/network$ cat 40-vlan304.netdev
[ネット開発]
名前=vlan304
種類=vlan
[VLAN]
ID=304
各 VLAN をブリッジに接続します。
igloo@core02:/etc/systemd/network$ cat 50-br0.network
【試合】
名前=br0
【ネットワーク】
VLAN=vlan304
[ブリッジVLAN]
VLAN=304
各物理ポートを定義し、必要な VLAN を接続します。
igloo@core02:/etc/systemd/network$ cat 70-swp20s0.network
【試合】
名前=swp20s0
【橋】
AllowPortToBeRoot=false
優先度=16
【ネットワーク】
ブリッジ=br0
[ブリッジVLAN]
VLAN=304
PVID=304
出口タグなし=304
最後に、各 VLAN の L3 インターフェイスを構成します。
igloo@core02:/etc/systemd/network$ cat 80-vlan304.network
【試合】
名前=vlan304
[リンク]
ActivationPolicy=常時稼働
【ネットワーク】
IPv6AcceptRA=0
キャリアなしで構成=はい
キャリア損失を無視=はい
【住所】
アドレス=2a0X:XXXX:XXXX:XXXX::1/64
リロード時に、systemd-networkd はインターフェイスを設定し、次にインターフェイスを ASIC にロードします。
igloo@core02:/etc/systemd/network$ sudo ブリッジ

VLANショー
ポートvlan-id
swp20s0 304 PVID 出力タグなし
igloo@core02:/etc/systemd/network$ sudo ifconfig vlan304
vlan304: flags=4163<UP,BROADCAST,RUNNING,MULTICAST> mtu 1500
inet6 fe80::XXXX:XXXX:XXXX:XXXX prefixlen 64 スコープ ID 0x20<リンク>
inet6 2a0X:XXXX:XXXX:XXXX::1 プレフィックスレン 64 スコープ ID 0x0<グローバル>
イーサ XX:XX:XX:XX:XX:XX txqueuelen 1000 (イーサネット)
RX パケット 0 バイト 0 (0.0 B)
RX エラー 0 ドロップ 0 オーバーラン 0 フレーム 0
TX パケット数 30384 バイト 4038796 (3.8 MiB)
TX エラー 0 ドロップ 2 オーバーラン 0 キャリア 0 コリジョン 0
ルーティング
ルートがカーネルにインストールされると、mlxsw はそれを ASIC にオフロードしようとします。成功すると、ルートにオフロード済みのフラグが立てられます。
igloo@core02:~$ ip ルート ショー 1.1.1.0/24
1.1.1.0/24 via XXX.XXX.XXX.XXX dev XXXXX proto Bird src XXX.XXX.XXX.XXX metric 32 offload rt_offload
残念ながら、Spectrum 1 は複数の完全なインターネット ルーティング テーブルをネイティブに処理するのにはあまり適していません。デフォルトの mlxsw リソース割り当てでは、94,208 の IPv4 および IPv6 ルートをサポートする余地が与えられます。この記事の執筆時点で、テーブル全体が IPv4 では約 104 万ルート、IPv6 では約 234,000 であることを考えると、理想的には複数のテーブルを含めることはもちろん、1 つのテーブルに含めることもできません。
幸いなことに、私たちにはありません

[切り捨てられた]

## Original Extract

Working with Mellanox Spectrum silicon and mlxsw to develop a SRE-friendly, long-term maintainable network switch

BYO NOS: Building a maintainable, SRE-friendly switch
I've got packets, they need moving.
For the last several years, I've used an Arista 7050QX2-32S as my core switch. It's a capable device that has served admirably. Unfortunately, it's also getting a little long in the tooth, having come out originally in 2016 and reached end of life in 2025.
Good alternatives are hard to come by. To understand why, we first need to understand the hardware we're working with.
From the perspective of hardware, the control plane of a network switch typically isn't complicated. In loose terms, the switch fabric is an extremely large - and expensive - chunk of silicon grafted onto the side of a reasonably low-power server from the era.
Arista's devices are built atop a Linux base with kernel modules and userspace extensions. While a traditional CLI is available, some operations inside of the base Linux system will be copied into the ASIC. For example, if an IPv4 or IPv6 route is inserted into the kernel, it'll be interpreted by Arista as a kernel sourced route (with a 'K' route origin in show route ) and will be offloaded into hardware. As a result, running alternative routing daemons (such as BIRD) on Arista hardware is quite trivial!
When I first acquired the hardware, I quickly took advantage of Arista's generous functionality - I've built considerable infrastructure around BIRD, non-SNMP provided metrics, and the ability to monitor system health like any other server.
Unfortunately, the open and seemingly accessible nature of Arista's hardware can act almost like vendor lock-in, as no other vendors that I could find have built such an extensive ecosystem around supporting third party, off-the-shelf daemons on their hardware. So what other options are there?
Some years ago, hyperscalers started to build their own switches. These devices (generally) didn't ship with bespoke ASICs, instead opting for an in-house software implementation running against existing silicon. While this shifted the burden of building and maintaining a NOS onto the hyperscalers, it simultaneously enabled them to control the underlying software running inside their NOS. Gone were the days of long-standing software bugs that a vendor refused to fix, replaced instead with bespoke control planes that made configuring switches feel more like interacting with a generic microservice.
The trend didn't stay internal. Publicly available, non-vendor developed NOSes began to pop up, notably including SONiC and DENT .
SONiC was originally developed by Microsoft and utilizes ASIC-specific, typically black-box blobs to support each individual vendor/SKU it runs on; as a result, support is largely limited by which vendors (or vendor licensees) have taken the time to develop the magic binary. While it's got support for my existing hardware, it's got one major problem: like many open source projects, support for various functionality is heavily dependent on individual need. It's quite evident that Microsoft had a need for Layer 3 functionality, while Layer 2 functionality has been spotty and prone to de-syncronization on all hardware I've used it on. Additionally, some functionality - like IP route insertion from the kernel - isn't available.
In contrast to SONiC, DENT does not take advantage of a project-specific abstraction layer; instead, it acts as a NOS-like interface on top of a generalized, vendor-provided, kernel level abstraction layer called switchdev.
Switchdev offloads familiar Linux functionality into the ASIC. Like Arista, functionality exists to copy IPv4 / IPv6 routes. Unlike Arista, additional common functionality - such as ACLs and gathering of interface metrics - can be performed via standard tooling. DENT is functionally a pre-built, NOS-like set of tools built to manage a Linux host.
While DENT itself didn't offer the exact tooling I was after, switchdev presented an opportunity.
DENT - and by extension switchdev - has a relatively small supported hardware list . Looking down the list, we see three major ASIC families that have some form of switchdev support: Marvell's Alleycat3, Marvell's Aldrin 2, and Mellanox's Spectrum. After a bit of research, I managed to pick up a Mellanox SN2010 - a Spectrum-based, switchdev-supporting chassis - for relatively cheap.
Opening it up, we see a fairly simple layout:
The concept of a "small PC that grew an ASIC" is fairly obvious; the computer itself is a tiny daughterboard!
In addition to the switching hardware, the chassis contains two unshielded power supplies and 4 internal fans. If you'd rather have the opposite direction of airflow, the fans can be turned around with three screws in a few minutes - not bad for a fully internal config.
Apart from the front panel SFP/QSFP ports, the device has 3 administrative ports: a RS232 serial port, a RJ45 Intel gigabit interface, and a USB port. The serial port outputs at a default baud rate of 115200, though the bios can only accept input keyboard input over USB.
Starting it up, pressing F7 and entering the bios password 'admin', we're dropped into a fairly standard AMI bios:
From here, we can set the default boot device to USB, swap the keyboard for a Debian installer, and reboot. Once we're at GRUB, keyboard input over serial starts to work, allowing us to boot into a serial-enabled Debian installer (by appending 'console=tty0 console=ttyS0,115200' to boot options in GRUB) and install Debian to the internal SSD. Once at a normal debian shell, we notice we're missing our interfaces:
igloo@core02:~$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
inet 127.0.0.1/8 scope host lo
valid_lft forever preferred_lft forever
inet6 ::1/128 scope host noprefixroute
valid_lft forever preferred_lft forever
2: enp0s20f0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
link/ether XX:XX:XX:XX:XX:XX brd ff:ff:ff:ff:ff:ff
inet6 fe80::XXXX:XXXX:XXXX:XXXX/64 scope link proto kernel_ll
valid_lft forever preferred_lft forever
While Mellanox's switchdev implementation, mlxsw, has been upstreamed, it's not enabled by default on most kernel builds. To fix it, we'll need to enable the kernel options, build our kernel, then install the kernel on the device. Fortunately, mlxsw's documentation provides the exact list for us :
CONFIG_NET_IPIP=m
CONFIG_NET_IPGRE_DEMUX=m
CONFIG_NET_IPGRE=m
CONFIG_IPV6_GRE=m
CONFIG_IP_MROUTE_MULTIPLE_TABLES=y
CONFIG_IP_MULTIPLE_TABLES=y
CONFIG_IPV6_MULTIPLE_TABLES=y
CONFIG_BRIDGE=m
CONFIG_VLAN_8021Q=m
CONFIG_BRIDGE_VLAN_FILTERING=y
CONFIG_BRIDGE_IGMP_SNOOPING=y
CONFIG_NET_SWITCHDEV=y
CONFIG_NET_DEVLINK=y
CONFIG_MLXFW=m
CONFIG_MLXSW_CORE=m
CONFIG_MLXSW_CORE_HWMON=y
CONFIG_MLXSW_CORE_THERMAL=y
CONFIG_MLXSW_PCI=m
CONFIG_MLXSW_I2C=m
CONFIG_MLXSW_MINIMAL=y
CONFIG_MLXSW_SWITCHX2=m
CONFIG_MLXSW_SPECTRUM=m
CONFIG_MLXSW_SPECTRUM_DCB=y
CONFIG_LEDS_MLXCPLD=m
CONFIG_NET_SCH_PRIO=m
CONFIG_NET_SCH_RED=m
CONFIG_NET_SCH_INGRESS=m
CONFIG_NET_CLS=y
CONFIG_NET_CLS_ACT=y
CONFIG_NET_ACT_MIRRED=m
CONFIG_NET_CLS_MATCHALL=m
CONFIG_NET_CLS_FLOWER=m
CONFIG_NET_ACT_GACT=m
CONFIG_NET_ACT_SAMPLE=m
CONFIG_NET_ACT_VLAN=m
CONFIG_NET_L3_MASTER_DEV=y
CONFIG_NET_VRF=m
The default interface names are not trivially parseable. Referencing Pin van Pelt's guide to Spectrum hardware , we can solve the issue with a udev rule:
igloo@core02:~$ cat /etc/udev/rules.d/10-local.rules
SUBSYSTEM=="net", ACTION=="add", DRIVERS=="mlxsw_spectrum*", NAME="sw$attr{phys_port_name}"
Following a reboot, we're dropped into mostly standard Debian, ready to build our new router.
I've got a lot of breakout cables that take a single QSFP interface and break them out into 4x SFP+ interfaces that can be used by hosts within my network. Given I'd like to avoid buying new cables, we can re-configure individual switchports for breakout cables via devlink:
/sbin/devlink port split $(/sbin/devlink port show | grep swp22 | cut -d " " -f 1 | cut -d ":" -f 1,2,3) count 4
We can confirm the port has been broken out successfully via ethtool:
igloo@core02:~$ sudo ethtool swp22s0
Settings for swp22s0:
Supported ports: [ FIBRE ]
Supported link modes: 1000baseKX/Full
10000baseKR/Full
25000baseCR/Full
25000baseSR/Full
Supported pause frame use: Symmetric Receive-only
Supports auto-negotiation: Yes
Supported FEC modes: Not reported
Advertised link modes: 1000baseKX/Full
10000baseKR/Full
25000baseCR/Full
25000baseSR/Full
Advertised pause frame use: No
Advertised auto-negotiation: Yes
Advertised FEC modes: Not reported
Speed: 10000Mb/s
Lanes: 1
Duplex: Full
Auto-negotiation: on
Port: Direct Attach Copper
PHYAD: 0
Transceiver: internal
Link detected: yes
L2 & L3 Interfaces
On a traditional Debian host, one is spoiled for choice when deciding on tools for managing network interfaces. I've historically used systemd-networkd, so I reached for the familiar thing.
systemd-networkd's configuration is verbose. For any operation at scale, I'd recommend either building a template-based config build system or to select something other than systemd-networkd, as one can quickly create a lot of config files:
igloo@core02:/etc/systemd/network$ ls -la | wc -l
83
That said, a basic configuration is relatively straightforward. Once systemd-networkd is installed, we need to define a bridge:
igloo@core02:/etc/systemd/network$ cat 40-br0.netdev
[NetDev]
Name=br0
Kind=bridge
# Note: for ASIC resource reasons, all MAC prefixes on the SN2010 must be the same 38-bits. The address here will be used to derive the prefix that will be used.
MACAddress=XX:XX:XX:XX:XX:XX
[Bridge]
VLANFiltering=1
VLANProtocol=802.1q
DefaultPVID=none
Priority=8
STP=yes
# Disable multicast snooping to avoid OSPF pain / having to run extra software to make it work
MulticastSnooping=false
A device for each vlan:
igloo@core02:/etc/systemd/network$ cat 40-vlan304.netdev
[NetDev]
Name=vlan304
Kind=vlan
[VLAN]
Id=304
Attach each vlan to the bridge:
igloo@core02:/etc/systemd/network$ cat 50-br0.network
[Match]
Name=br0
[Network]
VLAN=vlan304
[BridgeVLAN]
VLAN=304
Define each physical port, then attach the vlans we want:
igloo@core02:/etc/systemd/network$ cat 70-swp20s0.network
[Match]
Name=swp20s0
[Bridge]
AllowPortToBeRoot=false
Priority=16
[Network]
Bridge=br0
[BridgeVLAN]
VLAN=304
PVID=304
EgressUntagged=304
And finally, configure the L3 interface for each vlan:
igloo@core02:/etc/systemd/network$ cat 80-vlan304.network
[Match]
Name=vlan304
[Link]
ActivationPolicy=always-up
[Network]
IPv6AcceptRA=0
ConfigureWithoutCarrier=yes
IgnoreCarrierLoss=yes
[Address]
Address=2a0X:XXXX:XXXX:XXXX::1/64
On reload, systemd-networkd will configure the interfaces, which in turn will load them into the ASIC:
igloo@core02:/etc/systemd/network$ sudo bridge vlan show
port vlan-id
swp20s0 304 PVID Egress Untagged
igloo@core02:/etc/systemd/network$ sudo ifconfig vlan304
vlan304: flags=4163<UP,BROADCAST,RUNNING,MULTICAST> mtu 1500
inet6 fe80::XXXX:XXXX:XXXX:XXXX prefixlen 64 scopeid 0x20<link>
inet6 2a0X:XXXX:XXXX:XXXX::1 prefixlen 64 scopeid 0x0<global>
ether XX:XX:XX:XX:XX:XX txqueuelen 1000 (Ethernet)
RX packets 0 bytes 0 (0.0 B)
RX errors 0 dropped 0 overruns 0 frame 0
TX packets 30384 bytes 4038796 (3.8 MiB)
TX errors 0 dropped 2 overruns 0 carrier 0 collisions 0
Routing
When a route is installed into the kernel, mlxsw will attempt to offload it into the ASIC. When it's successful, the route will be flagged as offloaded:
igloo@core02:~$ ip route show 1.1.1.0/24
1.1.1.0/24 via XXX.XXX.XXX.XXX dev XXXXX proto bird src XXX.XXX.XXX.XXX metric 32 offload rt_offload
Unfortunately, Spectrum 1 is not well suited to natively handling multiple full internet routing tables. The default mlxsw resource allocation gives us room to support 94,208 IPv4 and IPv6 routes. Given that, at the time of writing, the full table is ~1.04 million routes on IPv4 / ~234k on IPv6, we can't fit one table, let alone the several we'd ideally want to.
Fortunately, we don't have

[truncated]
