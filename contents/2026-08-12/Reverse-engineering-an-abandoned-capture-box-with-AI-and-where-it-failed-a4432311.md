---
source: "https://chiptron.eu/reviving-old-unsupported-devices-with-ai-avermedia-game-capture-hd-ii/"
hn_url: "https://news.ycombinator.com/item?id=49272650"
title: "Reverse-engineering an abandoned capture box with AI (and where it failed)"
article_title: "Reviving Old Unsupported Devices with AI - AVerMedia Game Capture HD II - Chiptron.eu"
author: "smitka"
captured_at: "2026-08-12T14:13:56Z"
capture_tool: "hn-digest"
hn_id: 49272650
score: 1
comments: 1
posted_at: "2026-08-12T14:08:46Z"
tags:
  - hacker-news
  - translated
---

# Reverse-engineering an abandoned capture box with AI (and where it failed)

- HN: [49272650](https://news.ycombinator.com/item?id=49272650)
- Source: [chiptron.eu](https://chiptron.eu/reviving-old-unsupported-devices-with-ai-avermedia-game-capture-hd-ii/)
- Score: 1
- Comments: 1
- Posted: 2026-08-12T14:08:46Z

## Translation

タイトル: AI を使用した放棄されたキャプチャ ボックスのリバース エンジニアリング (および失敗した場所)
記事のタイトル: AI を使用して古いサポートされていないデバイスを復活させる - AVerMedia Game Capture HD II - Chiptron.eu
説明: AVerMedia Game Capture HD II を小さな Linux ボックスのように開きました。ファームウェア リバース エンジニアリング、ルート シェル、および AI で構築されたカスタム ソフトウェア (RTSP、640×480、60 fps)。

記事本文:
メインコンテンツにスキップ
フッターにスキップ
検索
サポートされていない古いデバイスを AI で復活させる – AVerMedia Game Capture HD II
つまり、AVerMedia Game Capture HD II (C285) をリバース エンジニアリングします。内部では、Texas Instruments DM368 SoC 上で小型の Linux が実行されています。ここにはロックダウンされたハードウェアはありません。 640×480 入力の受け入れの拒否は技術的な制限ではなく、ベンダー アプリケーションのエンコードのホワイトリストにすぎません。シリアル コンソール、復号化されたシステムのダンプ、ルート シェル (bootargs の init=/bin/sh) を介してボックスを開け、そのファームウェアをほぼ完全に AI によって作成されたカスタム ファームウェアに置き換えることができます。その結果、RTSP ストリーム、640×480 から 1080p までの任意の解像度、60 fps を備えたオープン ネットワーク キャプチャ ボックスが実現しました。オリジナルでは死んだ雲と消えたモバイル アプリしか提供されませんでした。
今日 AI が可能にするものを見るのは驚くべきことです。 15 年前ならメーカーの開発部門全体が必要で数か月かかった作業量を、ほぼ誰でも 1 日で処理できるようになります。 AI が代わりに仕事をしてくれるというわけではありません。後でわかるように、私は AI のナンセンスを打ち破るのに多くの時間を費やしました。しかし、一人の人間が週末に達成できることの変化は非常に大きいです。
当時の残りの箱、AVerMedia Game Capture HD II、モデル C285 で試してみました。堅牢なハードウェア – 内部には、432 MHz の Texas Instruments DM368 SoC、256 MB RAM、およびハードウェア H.264 エンコーダが搭載されています。
発売時には、ディスクに録画したり、YouTube に直接ストリーミングしたり、携帯電話から制御したりするなど、賢明なことが実行できます。しかし、YouTube API はとうの昔に消滅し、モバイル アプリはストアから消え、私は数年前にリモコンを紛失しました。サポートは終了しました。残ったのは稼働中のハードウェアであり、そこへの使用可能なパスはありませんでした。
ある理由から埃を落としました。取り組んでいます

マイクロコントローラー用のピコゲーム エンジンであり、ビデオ出力 (DVI 640×480) をキャプチャするために必要でした。ディスクへの記録はハードウェア ボタンでトリガーできるため、このボックスが少なくともそのような単純なことを処理できることを望みました。
それはできませんでした😥。入力は非常に限られた解像度セットのみをサポートしており、VGA 640×480 はその中に含まれていませんでした。ボックスは単純に信号を拒否しました。
そのとき、私は中を覗いてみようと決心しました。オチはこの記事全体の内容を正確に示しているので、すぐに紹介します。制限はハードウェアにあるわけではありません。カーネル ドライバーにはサポートされる入力モードのテーブルがあり、640×480 プログレッシブ エントリは常にそこにありました。ベンダーのアプリからはアクセスできませんでした。
私は、640×480 が解像度が悪いと報告されている理由を知りたかったのです。ハードウェアまたはソフトウェアの制限によるものでした。そこで箱を開けて、実際に何が動いているのかを確認してみました。
分解して初めて明らかになったことが 1 つあります。それは、ITE チップと SoC の間にいくつかの 74LVC244A ドライバーが配置されていることです。これはビデオ バスがどのようにルーティングされているかを示す物理的な証拠であり、後で意味がわかりました。受信機と送信機は共有 16 ビット バスを介してデータを渡し、SoC はそこから読み取るだけです。そのため、SoC を介したアクティブなキャプチャがなくても、モニターに画像が表示されます。
NAND には 2 つの完全なファームウェア セットが保持されます。
mtd0 ブートローダー 0x00000000–0x003c0000 3.75 MB
mtd1 params 0x003c0000–0x00400000 256 KB (U-Boot 環境)
mtd2 カーネル_fw1 0x00400000–0x00800000 4 MB
mtd3 ファイルシステム_fw1 0x00800000–0x04000000 56 MB
mtd4 カーネル_fw2 0x04000000–0x04400000 4 MB
mtd5 ファイルシステム_fw2 0x04400000–0x07c00000 56 MB
mtd6 プログラマー_スペース 0x07c00000–0x08000000 4 MB (書き込み可能な jffs2)
2 つのバンクはブートローダーの 1 つの変数によって切り替えられます。それは判明しました

重要 – 一方の銀行は、私がもう一方の銀行を繰り返し上書きし、何度もブリックした間、いつでも頼れるセーフティネットとして残りました。
ハードウェアでそれができるのに、なぜやらないのでしょうか?
データシートには、これらのチップが 640×480 で問題を起こすというヒントはありませんでした。 IT6604 は、VGA モードを問題なく処理できる標準の HDMI/DVI レシーバーです。 DM368 の IPIPE は、最大 2176 ピクセルの幅で動作します。 640×480 は、ボックスがサポートする 1080p よりも負荷が大幅に軽いです。
そこで私は、制限はソフトウェアにあるに違いないと考えました。そして、おそらくそのボックスではある種の Linux が実行されているのではないかと考えました。簡単な検索でそれが確認されましたが、それは私が期待していたよりも徹底的でした。
Texas Instruments は、このファミリ用の DVSDK (デジタル ビデオ ソフトウェア開発キット) を提供しました。それは単なるドライバーではありません。これはまさにこの種のデバイスのための完全なツールキットです。
codecs-dm365 – HDVICP で実行される H.264 エンコーダー
linuxutils – カーネルモジュール cmemk (連続 DMA メモリ)、 edmak 、 irqk 、 dm365mmap
そして最も重要なのは、 encode と呼ばれるデモを含む dvsdk-demos です。
それが鍵です。このボックス内のベンダー アプリケーションは encode と呼ばれ、初期化スクリプト /etc/init.d/encode-demo によって起動されます。言い換えれば、私がかつて購入した商用製品は、テキサス・インスツルメンツの SDK デモの子孫であり、その名前は今でも残っています。
これらのカーネル モジュールをロードするloadmodules_hd.sh スクリプトを含む、SDK 全体がボックスの /opt/dvsdk/dm368/ にまだ存在しています。バージョンは時代と一致しています。Linux 2.6.32.17-davinci1 カーネルは 2017 年 12 月に gcc 4.3.3 でビルドされ、U-Boot 2009.03 ブートローダーは 2014 年 1 月から使用されています。
それが仕事を変えました。それはもはやハードウェアの制限を回避することではなく、ベンダー アプリが隠れているレイヤーに到達することでした。
もう存在しないモバイルアプリ
より深く掘り下げる前に、シミュレーションを試してみました

一番の問題: このボックスはネットワーク上で実際に何を提供するのでしょうか?モバイル アプリの API を実行している単一ポートで応答しました。そのアプリはストアから消えてしまったので、それ自体では役に立ちませんでした。
なんとか古いAPKを見つけて逆コンパイルしました。プロトコル全体が明らかになりました。それは単純で、ポート 24170 で暗号化されていない HTTP でした。認証もTLSも複雑さもありません。逆方向のトラフィックもあります。ボックスは TCP ポート 1516 でイベントをアプリに送信します。
PIN とのペアリング、入力切り替え、録画の開始と停止、リモート ボタンのエミュレーションなど、合計 117 のエンドポイントを計画しました。
ここで AI が初めて登場しました。これらの発見に基づいて、私は AI に、ボックスを制御するための CLI ツールと小さな Web アプリを作成させました。最初の試行で実質的に機能し、紛失したリモコンと機能しなくなったモバイル アプリの両方を完全に置き換えました。ファームウェアに触れることなく、ボックスを再び制御できるようにしました。
シリアルコンソール: 内部の最初のエントリ
コントロールは戻りましたが、元の問題は残りました。私は、メニューのどこかに、ようやく実行できるようになり、解像度を有効にする設定が見つかることを期待していました。そんな幸運はない。メニューにはそのようなものはありません。
さらなるレベルへ進む時が来ました。ボード上に UART ポートのようなものを見つけました。ピンを注意深くたどり、コンバータを接続すると、すぐに U-Boot が画面に表示されました (115200 8N1、データシートと一致)。
U-Boot 2009.03-ダーティ (2014 年 1 月 14 日 - 09:44:01)
I2C: 準備完了
ドラム: 256MB
NAND: 128 MiB
入力: シリアル
アウト: シリアル
エラー: シリアル
Linux が続き、ベンダー アプリが続きました。最初の障害: アプリが入力を引き継いだため、アプリが起動するとシリアル コンソールが役に立たなくなりました。
CTRL+C を押して起動を中断すると、ログイン プロンプトが表示されます。
| _ |___ ___ ___ ___ | _ |___ ___ |__|___ ___| |_
| | _| .'| 。 | 。 | | __| _| 。 | | | -_| _| _|
|__|__|_| |__、

|_ |___| |__| |_| |__|_| |__|___|_|
|__| |__|
アラゴプロジェクト http://arago-project.org dm368-evm
アラゴ 2011.02 dm368-evm
Arago は、TI がこのファミリー向けに出荷したディストリビューションであり、OpenEmbedded 上に構築されています。ホスト名: dm368-evm 、開発キットの名前に注目してください。繰り返しになりますが、チップ ベンダーから提供されたままのリファレンス デザインです。
そのログインさえあまり役に立ちませんでした。ベンダー アプリは並行して起動し、再びシリアル入力を無視しました。コンソールを取られてしまいました。
ただし、CTRL+C で明らかになったことが 1 つあり、ログイン自体よりも重要でした。それは、機能したということです。ボックスは実際に私のキャラクターを受信して​​いました - シリアル回線は双方向でした。そのことを念頭に置いてください。すぐにパズルになります。
これにより、Linux が起動する前に U-Boot に入るという、下位層が 1 つ残されました。ブートローダー自体は、私が必要としていたものを正確に提供してくれました。
自動起動を停止するには、任意のキーを押します: 1
任意のキーを押します。 1秒のカウントダウン。
だからこそ、シリアル通信は明らかに機能していたため、単なる配線ミスではなく、真のパズルでした。登場人物たちが去っていった。ブートローダーが何らかの理由でそれらを破棄していました。
ダウンロードしたファームウェアは暗号化されています
もちろん、コンソールをクラックする前に、私はより簡単な方法を試しました。メーカーがかつてリリースした古いファームウェアをダウンロードして、ベンチで分析するというものです。運がなかった。
ファイルには読み取り可能なヘッダーがありました。
AVT-CEDxc285WW001.002.0990.0.20.0.32 201710121914 723c8a4f5cacd1cc5a3fabb...
#k_s
2364296
#f_m
16
#f_s
3145752
3145752
...
ただし、そのヘッダー以降はすべて暗号化されます (エントロピー 1 バイトあたり 8 ビット = 本質的にランダムなデータ)。
これにより、実行中のシステム内で復号化された状態に到達する必要があるという目標が明確になりました。フラッシュをチップから直接ダンプすることは考えられませんでした。私はそのための機器を持っていません（単純な SPI ではなく、はんだ除去が必要なパラレル フラッシュです）。
それは残った

コンソールとキーストロークを無視した U-Boot です。私が試したトリックの 1 つは、シリアル入力にノイズを送信することでした。これは、特定のキーを押すのではなく、ランダムなデータのストリームだけです。
そしてそれはうまくいきました🤯。画面には次のことが表示されました。
DM365 EVM >
ブートローダーのプロンプト!
カウントダウン延長。次回キーを押すまでの時間を長くできるように、ブート遅延を 1 秒から 5 秒に設定しました。ネタバレ: 役に立たなかったのですが、なぜこれが物語全体の最良の部分の 1 つであるのかについては、またお話しします。
プロンプト自体はまだ勝利ではありません。この U-Boot でできることとできないことに注意してください。コマンド テーブルには nand 、 setenv 、 saveenv 、 bootm 、 dhcp 、 tftpboot がありますが、 usb 、 fatload 、 mmc はありません。 U-Boot だけでは USB スティックに何も保存できません。ネットワーク経由でのみ通信できます。プロンプトの実際の値は別の場所にあり、ブート時にカーネルに渡される内容を書き直すことができました。そしてまさにそこが次のステップです。
ブートローダーからシェルへ: bootargs を書き換えます
元のブートパラメータは次のようになっていました。
bootargs=mem=97M console=ttyS0,115200n8 Quiet root=/dev/mtdblock5 rootfstype=cramfs ro
video=davincifb:vid0=OFF:vid1=OFF:osd0=1280x720x16,5400K
dm365_imp.oper_mode=0 davinci_capture.device_type=4
davinci_enc_mngr.ch0_output=液晶 davinci_enc_mngr.ch0_mode=720P-60
bootcmd=nboot 0x80700000 0 0x4000000;bootm
この一文から多くのことが分かりました。
システムは mtdblock5 – 2 番目のバンク、 filesystem_fw2 から実行されます。
ルート ファイルシステムは cramfs で、読み取り専用でマウントされます (したがって、簡単に上書きすることはできません)。
カーネルは 256 MB のうち 97 MB を取得します。残りはコーデック用に予約されます。
ビデオ出力は 720p60 にハードコードされており、
そして、パラメーターを含むドライバー名があります: dm365_imp.oper_mode 、 davinci_capture.device_type 、 davinci_enc_mngr.ch0_output 。それらはどれも後で役に立つことがわかりました。
init=/bin を追加するだけで済みました

最後に /sh を付けます。ベンダー アプリは init スクリプトからのみ起動します。init が実行されない場合、アプリは決して起動せず、シリアル コンソールを盗むこともありません。実行中の復号化されたシステム上に初めて実際の root シェルが存在します。
ここからはルーチンです。USB スティックをマウントし、 /dev/mtd* から dd を使用して NAND パーティションを 1 つずつダンプします。ブートローダー、両方のカーネル、両方のファイルシステム。
復号化されたシステム内の宝物
これで、ファイルシステム全体がディスク上に保存され、自由に探索できるようになりました。それはすぐに3回の成果を上げました。
ダウンロードした .bin ファイルのエントロピーがちょうど 8.0 で、解凍できなかったことを覚えていますか?それらを復号化する方法に対する答えは、次のコマンドのメイン エンコード アプリケーション内にあります。
openssl enc -d -des3 -in /tmp/file.en -out /tmp/file.de -pass pass:Av3rMed1a
これは、パスワード Av3rMed1a を持つ 3DES です。このパスワードに何か見つかるかもしれません 🙂。コマンドは 1 つだけです。
openssl enc -d -des3 -md md5 -pass pass:Av3rMed1a -in blob.en -out blob.de
暗号は対称的で、読み取り可能なヘッダーがパッケージを再構築するための完全なレシピとなります。このキーを使用して、アプリのアップデーターが受け入れるカスタムの「公式」ファームウェアを構築することもできます。しかし、私はそれを必要としたことはありませんでした。root シェルを使用すると、変更されたシステムを非アクティブな NAND バンクに直接書き込むだけで、更新メカニズム全体をバイパスできます。
ルートパスワード: DES 暗号化と 30 分

[切り捨てられた]

## Original Extract

I opened the AVerMedia Game Capture HD II like a small Linux box: firmware reverse engineering, a root shell and custom software built with AI – RTSP, 640×480 and 60 fps.

Skip to main content
Skip to footer
Search
Reviving Old Unsupported Devices with AI – AVerMedia Game Capture HD II
In short: Reverse engineering the AVerMedia Game Capture HD II (C285) . Inside runs a small Linux on a Texas Instruments DM368 SoC – no locked-down hardware here. The refusal to accept 640×480 input was never a technical limit, just a whitelist in the vendor application encode . Via the serial console, a dump of the decrypted system and a root shell ( init=/bin/sh in bootargs), the box can be opened up and its firmware replaced with a custom one written almost entirely by AI. The result: an open network capture box with an RTSP stream, any resolution from 640×480 up to 1080p, and 60 fps where the original offered nothing but a dead cloud and a vanished mobile app.
It’s amazing to watch what AI enables today. In one day, almost anyone can handle the amount of work that fifteen years ago would have required an entire manufacturer’s development department and taken months. I don’t mean AI does the work for you – as you’ll see later, I spent a lot of time shooting down its nonsense. But the shift in what’s achievable for one person over a weekend is enormous.
I tried it on a box left over from that era: AVerMedia Game Capture HD II , model C285 . Solid piece of hardware – inside sits a Texas Instruments DM368 SoC at 432 MHz, 256 MB RAM, and a hardware H.264 encoder.
At launch it could do sensible things: record to disk, stream straight to YouTube, be controlled from a phone. But the YouTube API has long since vanished , the mobile apps disappeared from the stores, and I lost the remote years ago. Support ended. What remained was working hardware with no usable path to it.
I dusted it off for a specific reason. I’m working on the picogame engine for microcontrollers and needed to capture their video output – DVI 640×480 . I hoped the box could at least handle something that simple, since recording to disk can be triggered by the hardware button.
It couldn’t 😥. The input only supported a very limited set of resolutions and VGA 640×480 wasn’t among them . The box simply rejected the signal.
That was the moment I decided to look inside. I’ll give away the punchline right away because it shows exactly what this whole piece is about: the limitation wasn’t in the hardware. The kernel driver has a table of supported input modes and the 640×480 progressive entry had been there the whole time . You just couldn’t reach it through the vendor app.
I wanted to know why 640×480 was being reported as a bad resolution – whether it was a hardware or software limit. So I opened the box and looked at what was actually driving it.
One thing only became obvious after disassembly: several 74LVC244A drivers sit between the ITE chips and the SoC. That’s physical proof of how the video bus is routed, and it made sense later: the receiver and transmitter pass data over a shared 16-bit bus and the SoC just reads from it. That’s why the monitor can show the picture even without active capture through the SoC.
The NAND holds two complete firmware sets :
mtd0 bootloader 0x00000000–0x003c0000 3.75 MB
mtd1 params 0x003c0000–0x00400000 256 kB (U-Boot environment)
mtd2 kernel_fw1 0x00400000–0x00800000 4 MB
mtd3 filesystem_fw1 0x00800000–0x04000000 56 MB
mtd4 kernel_fw2 0x04000000–0x04400000 4 MB
mtd5 filesystem_fw2 0x04400000–0x07c00000 56 MB
mtd6 programer_space 0x07c00000–0x08000000 4 MB (writable jffs2)
Two banks switched by a single variable in the bootloader. That turned out to be crucial – one bank stayed as a safety net I could always fall back to while I repeatedly overwrote the other and bricked it several times.
The hardware can do it, so why not?
The datasheets gave no hint that these chips would have trouble with 640×480. The IT6604 is a standard HDMI/DVI receiver that handles VGA modes without issue. The IPIPE in the DM368 works with widths up to 2176 pixels. 640×480 is a dramatically lighter load than the 1080p the box supports.
So I figured the limit had to be in software – and that the box was probably running some kind of Linux. A quick search confirmed it, and more thoroughly than I had hoped.
Texas Instruments supplied a DVSDK for this family – Digital Video Software Development Kit . It’s not just a driver; it’s a complete toolkit for exactly this kind of device:
codecs-dm365 – the H.264 encoder running on the HDVICP
linuxutils – kernel modules cmemk (contiguous DMA memory), edmak , irqk , dm365mmap
and most importantly dvsdk-demos , which includes a demo called encode
That’s the key. The vendor application in this box is called encode and is started by the init script /etc/init.d/encode-demo . In other words, the commercial product I once bought is a descendant of a Texas Instruments SDK demo – and the name is still there today.
The entire SDK is still present on the box in /opt/dvsdk/dm368/ , including the loadmodules_hd.sh script that loads those kernel modules. The versions match the era: Linux 2.6.32.17-davinci1 kernel built in December 2017 with gcc 4.3.3, U-Boot 2009.03 bootloader from January 2014.
That changed the job. It was no longer about bypassing a hardware limit – it was about reaching the layer the vendor app was hiding.
The mobile app that no longer exists
Before digging deeper I tried the simplest thing: what does the box actually offer over the network? It answered on a single port running the API for the mobile app. That app had vanished from the stores, so it was useless on its own.
I managed to find an old APK and decompiled it. The whole protocol came out. It was simple – unencrypted HTTP on port 24170 . No authentication, no TLS, no complications. There’s also traffic the other way: the box sends events to the app on TCP port 1516.
I mapped out 117 endpoints in total – pairing with PIN, input switching, starting and stopping recording, emulating remote buttons.
This is where AI first came in. From those findings I had it write a CLI tool and a small web app to control the box. It worked practically on the first try and completely replaced both the lost remote and the dead mobile app . Without touching the firmware I had the box controllable again.
Serial console: first entry inside
I had control back, but the original problem remained. I hoped that somewhere in the menu – which I could now finally walk through – I’d find a setting that would enable the resolution. No such luck. Nothing like that exists in the menu.
Time to go one level deeper. On the board I found what looked like a UART port . I carefully traced the pins, hooked up a converter, and soon U-Boot appeared on screen (115200 8N1, matching the datasheet).
U-Boot 2009.03-dirty (Jan 14 2014 - 09:44:01)
I2C: ready
DRAM: 256 MB
NAND: 128 MiB
In: serial
Out: serial
Err: serial
Linux followed, then the vendor app. First snag: the app took over the inputs , so the serial console became useless once the app started.
I could interrupt the boot with CTRL+C and get a login prompt:
| _ |___ ___ ___ ___ | _ |___ ___ |_|___ ___| |_
| | _| .'| . | . | | __| _| . | | | -_| _| _|
|__|__|_| |__,|_ |___| |__| |_| |___|_| |___|___|_|
|___| |___|
Arago Project http://arago-project.org dm368-evm
Arago 2011.02 dm368-evm
Arago is the distribution TI shipped for this family – built on OpenEmbedded. And note the hostname: dm368-evm , the name of the development kit. Again, the reference design exactly as it left the chip vendor.
Even that login wasn’t much use – the vendor app started in parallel and ignored the serial input again . The console was taken.
One thing the CTRL+C did reveal, though, and it mattered more than the login itself: it worked. The box really was receiving my characters – the serial line was bidirectional. Keep that in mind; it’ll become a puzzle shortly.
That left one lower layer: getting into U-Boot before Linux even starts. The bootloader itself offered exactly what I needed:
Hit any key to stop autoboot: 1
Press any key. One-second countdown.
And that’s why it was a real puzzle, not just a wiring mistake, because serial was demonstrably working. Characters were leaving. The bootloader was just discarding them for some reason.
Downloaded firmware is encrypted
Before trying to crack the console I of course tried the easier route: download the old firmware the manufacturer once released and dissect it on the bench. No luck.
The files had a readable header:
AVT-CEDxc285WW001.002.0990.0.20.0.32 201710121914 723c8a4f5cacd1cc5a3fabb...
#k_s
2364296
#f_m
16
#f_s
3145752
3145752
...
But everything after that header is encrypted (entropy 8 bits per byte = essentially random data).
That clarified the goal: I had to reach the decrypted state inside the running system. Dumping the flash directly from the chip was off the table – I don’t have the gear for it (it’s parallel flash that would need desoldering, not simple SPI).
That left the console and the U-Boot that ignored keystrokes. One trick I tried was sending noise to the serial input – just a stream of random data instead of specific key presses.
And it worked 🤯. The screen showed:
DM365 EVM >
Bootloader prompt!
Extended countdown. I set bootdelay from one second to five so I’d have more time to hit a key next time. Spoiler: it didn’t help, and why is one of the best parts of the whole story – we’ll come back to it.
The prompt itself isn’t the win yet. Worth noting what this U-Boot can and can’t do: the command table has nand , setenv , saveenv , bootm , dhcp , tftpboot – but no usb , fatload or mmc . From U-Boot alone you couldn’t save anything to a USB stick; it could only reach out over the network. The real value of the prompt was elsewhere – it let me rewrite what gets passed to the kernel at boot . And that’s exactly where the next step lay.
From bootloader to shell: rewrite bootargs
The original boot parameters looked like this:
bootargs=mem=97M console=ttyS0,115200n8 quiet root=/dev/mtdblock5 rootfstype=cramfs ro
video=davincifb:vid0=OFF:vid1=OFF:osd0=1280x720x16,5400K
dm365_imp.oper_mode=0 davinci_capture.device_type=4
davinci_enc_mngr.ch0_output=LCD davinci_enc_mngr.ch0_mode=720P-60
bootcmd=nboot 0x80700000 0 0x4000000;bootm
That line revealed a lot:
the system runs from mtdblock5 – the second bank, filesystem_fw2 ,
the root filesystem is cramfs, mounted read-only (so nothing can be overwritten easily),
the kernel gets 97 MB out of 256 – the rest is reserved for the codec,
video output is hard-coded to 720p60 ,
and there are driver names with parameters: dm365_imp.oper_mode , davinci_capture.device_type , davinci_enc_mngr.ch0_output . Every one of them proved useful later.
All I had to do was append init=/bin/sh at the end. The vendor app only starts from init scripts – if init never runs, the app never starts, doesn’t steal the serial console, and for the first time I have a real root shell on a running, decrypted system .
From here it was routine: mount a USB stick and dump the NAND partitions one by one with dd from /dev/mtd* . Bootloader, both kernels, both filesystems.
Treasures in the decrypted system
Now I had the entire filesystem on disk and could poke around at leisure. It paid off three times immediately.
Remember those downloaded .bin files with exactly 8.0 entropy that wouldn’t unpack? The answer to how to decrypt them sat inside the main encode application in this command:
openssl enc -d -des3 -in /tmp/file.en -out /tmp/file.de -pass pass:Av3rMed1a
That’s 3DES with the password Av3rMed1a – you might spot something in that password 🙂. Now just one command:
openssl enc -d -des3 -md md5 -pass pass:Av3rMed1a -in blob.en -out blob.de
The cipher is symmetric and the readable header is the complete recipe for reassembling the package. The key could even be used to build a custom “official” firmware that the app’s updater would accept. I never needed it though: with a root shell I can just write the modified system straight into the inactive NAND bank and bypass the whole update mechanism.
Root password: DES crypt and half an hour w

[truncated]
