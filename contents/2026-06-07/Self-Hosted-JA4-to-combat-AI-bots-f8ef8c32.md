---
source: "https://blog.miloslavhomer.cz/deploying-ja4/"
hn_url: "https://news.ycombinator.com/item?id=48437214"
title: "Self-Hosted JA4 to combat AI bots"
article_title: "Deploying JA4\n- Miloslav Homer"
author: "ArcHound"
captured_at: "2026-06-07T18:32:49Z"
capture_tool: "hn-digest"
hn_id: 48437214
score: 2
comments: 0
posted_at: "2026-06-07T18:06:23Z"
tags:
  - hacker-news
  - translated
---

# Self-Hosted JA4 to combat AI bots

- HN: [48437214](https://news.ycombinator.com/item?id=48437214)
- Source: [blog.miloslavhomer.cz](https://blog.miloslavhomer.cz/deploying-ja4/)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T18:06:23Z

## Translation

タイトル: AI ボットと戦うセルフホスト型 JA4
記事のタイトル: JA4 の導入
- ミロスラフ・ホーマー

記事本文:
サーバーに来るボットを識別/ブロック/レート制限したいと考えています。ベンダー ロックインを回避し、オープンソース ソリューションを使用したいと考えています。また、PCAP 検査 (Suricata、Zeek など) は暗号化されたクライアント Hello の脅威にさらされるため (セットアップやメンテナンスに手間がかかり、費用がかかる) も避けたいと考えています。
JA4 はこれらの要件に適合します。前のパートでは、JA3/JA4 について一般的に説明しました。
どちらも、ClientHello パラメータ (暗号スイート、拡張機能、ALPN など) を分析する TLS フィンガープリント手法です。
Chrome が拡張機能を並べ替えるため、JA3 は廃止されました。そのため、現在十分にサポートされている JA4 を使用してください。
JA4+ (メソッドの完全なスイート) には混合ライセンス モデルがあるので注意してください。
このパートでは、このテクノロジーを導入してみましょう。 JA4 をサポートするツールのリストはある程度整備されており、そこから選択できます。上記の要件を維持すると、次のオプションのみが見つかりました。
Nginx プラグイン : JA4 作者によって作成されましたが、進行中としてマークされています。
Rama プロキシ : 錆びていますが、まだアルファ版です。
Envoy プロキシ : リストにはありませんが、JA4 をサポートしているようです。
HAProxy プラグイン : 私の選択です。この投稿では、HAProxy をデプロイ/構成して JA4 とユーザー エージェントをログに保存する方法について説明します。
本当に詳しく知りたい場合は、tcpdump によって PoC としてキャプチャされた PCAP を解析するために、scapy を使用して JA4 計算を Python で実装しました。本番環境にまったく対応していませんが、テクノロジーについてはかなり学びました。 PCAP の解析は難しい部分ですが、抽出されたデータから JA4 を計算する方がはるかに簡単です。
まだ興味がありますか?標準と頭字語の難題から始めましょう。
前提条件、テクノロジー、RFC、および短縮表記
この作業を進めるには多くの基準を調整する必要があるため、用語の説明に圧倒されそうになります。ここで簡単に説明します。
TLS (Transport Layer Security) RFC8446 : HTTP およびその他の多くのプロトコルの暗号化層

トコール。これには多くの拡張機能があります。
SSL (Secure Socket Layer): TLS の以前のバージョン。TLS と完全に互換性があります。 HTTPS の参照や SSL/TLS の短縮形が今でも目にされることがあります。
Client Hello : TLS を確立するためのクライアントからサーバーへの最初のメッセージ。
SNI (サーバー名識別) RFC6066 : クライアントは、どのバックエンドに接続するかをサーバーに指示するために、この拡張機能を client hello に含めます。
ALPN (アプリケーション層プロトコル ネゴシエーション) RFC7443 : client hello 拡張機能。TLS が確立された後にどのプロトコルを使用するかをサーバーに説明するヒントです。
GREASE (拡張性を維持するためのランダムな拡張機能の生成) RFC8701 : クライアントが、クラッシュしないように正しく処理する必要があるランダムな場所で、選択したナンセンスな値を使用してマシンをトロールできるようにします。
ECH (Encrypted Client Hello) RFC9849 : SNI や ALPN などを覗き見から隠すためにクライアント Hello を暗号化する新しい標準。これに関する私の記事。
PCAP (パケット キャプチャ): パケット データを保存するための形式。80 年代に libpcap の実装によって確立された事実上の標準。正式な RFC はまだありません。
これらの点は、実際には、これらの略記や標準を理解する際に覚えておくべき単なる参考資料です。それでは、JA4 をデプロイしてみましょう。
多くの大手サービスプロバイダーのいずれかに連絡して、これを行うことができます。ここにリストがありますので、お気に入りを選んでください。
明らかに、それは私がこのブログでやりたいことではありません。このブログの目的は、独立した何かをすることです。テクノロジーに重点を置き、ベンダーロックインを可能な限り回避したいと考えています。
たとえば、Cloudflareは、Enterpriseプランでのみこれを提供します。これは、最高のプランです。多少の値上げは予想されますが、すでにこれらのプランを持っている場合は、迷う必要はありません。プランをアップグレードした場合でも、プロバイダーを切り替えたり、独自に契約したりするよりも安くなる場合があります。
同じことから

リストから、JA4 をサポートできるオープンソース ソリューションをいくつか選択できます。
これらのソリューションの多くは、pcap をキャプチャしてトラフィックを分析します (Suricata、Zeek、Wireshark、Arkime など)。必要な ClientHello プロパティが暗号化されるため、暗号化された Client Hello を使用する場合は注意してください。
私はそれをリバースプロキシにデプロイしたいと思っています。 JA4 作成者からの「公式」nginx モジュールがありますが、その説明によると、それは進行中の作業です。彼らは最近 DB を非表示にしたため、この拡張機能の将来がどうなるかはわかりません。
また、 rama proxy もあります。これは見た目は素晴らしく (新しい、錆びた、機能セット)、随所に LLMism があります。私はこのプロジェクトが嫌いではありませんが、利用していません。
Rust ベースのプロキシといえば、pingora が JA4 をサポートしようとしているようです。
Envoy はすでに存在しているようですが、リストには載っていません。ゴーフィギュア。お気に入りのプロキシ/サーバーがある場合は、まずそのドキュメントを確認してください。実際にそれをサポートしている可能性があります。
上記の問題を解決するために、私は HAProxy JA4 プラグインを選択しました。
haproxy が必要ですが、少なくとも 3.1 が必要です。これは debian 13 のデフォルトではないので、リポジトリを含める方法についてはここを参照してください。 uvx fastapi-new hw を備えたサンプル アプリも必要です。もちろん、標準的なことを行う必要があります 1 。
cd /etc/haproxy
mkdir lua; CDルア
wget https://raw.githubusercontent.com/O-X-L/haproxy-ja4-fingerprint/refs/heads/latest/ja4.lua
これを構成ファイルでも設定する必要があります。
グローバル
...
#JA4関連
tune.ssl.capture-buffer-size 2048
lua-load /etc/haproxy/lua/ja4.lua
フロントエンドhttps-in
バインド *:443 ssl crt /etc/haproxy/certs/haproxy-test.local.pem
デフォルトバックエンド hello_world
http-リクエスト lua.fingerprint_ja4
http リクエスト キャプチャ var(txn.fingerprint_ja4) len 36
http-リクエスト キャプチャ req.hdr(ユーザー エージェント) len 15

0
バックエンド hello_world
サーバーapp1 127.0.0.1:8000チェック
まず、フロントエンドでこれらのチェックと TLS 終端を設定し、トラフィックを正しいバックエンドにルーティングする必要があります。
haproxy に SSL/TLS 情報を取得してプラグインをロードするように指示する必要があります。 tune.ssl.capture-buffer-size 2048 に関する小さなメモ - 拡張機能の作成者はこれを 192 に設定することを推奨していますが、基本的に ClientHello 全体をここに収めようとしているため、1568 バイト長の署名を持つことができる Kyber のようなポスト量子暗号を考慮することができます。
次に、JA4 を計算し、それをユーザー エージェントとともにログにキャプチャするようにフロントエンドに指示する必要があります。
5 月 29 日 20:03:42 haproxy-test haproxy[3319]: [通知] (3319) : 読み込みに成功しました。
5 月 29 日 20:03:42 haproxy-test systemd[1]: haproxy.service - HAProxy ロード バランサを開始しました。
5 月 29 日 20:06:18 haproxy-test haproxy[3322]: 192.168.100.1:46182 [2026 年 5 月 29 日:20:06:18.238] https-in~ hello_world/app1 0/0/0/1/1 404 153 - - ---- 1/1/0/0/0 0/0 {t13d1715h2_5b57614c22b0_7121afd63204|Mozilla/5.0 (X11; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0} "GET https://haproxy-test.local/test HTTP/2.0"
中括弧内で、JA4 とユーザー エージェントの両方がさらなる処理のためにキャプチャされたことがわかります。
ここからは、何でもできます。後のチェックのためにログをキャプチャし、いくつかの IDS/IPS に転送し、アプリケーションのヘッダーを設定します。お気に入りの LLM が必要なものをキャプチャするスクリプトを生成すると確信しています。そうでない場合は、1 時間ごとに実行する必要がある私の cronjob スクリプトを次に示します。
しかし、ああ、Lua は非常に遅く、シングルスレッドであると解釈されます。ああ、これではパフォーマンスが低下します... さて、簡単なテストを実行しました。
クライアント: ffuf (デフォルトでは 40 スレッド) この wordlist 。
ネットワーク: ネットワークの問題はありません - ローカルホスト接続。
JA4 収集を無効にして 1 回、JA4 収集を有効にして 1 回。
$

ffuf -u https://haproxy-test.local/FUZZ -w fav-wordlist.txt -mt ">100"
...
:: 進行状況: [90821/90821] :: ジョブ [1/1] :: 3278 要求/秒 :: 期間: [0:00:40] :: エラー: 0 ::
$ ffuf -u https://haproxy-test.local/FUZZ -w fav-wordlist.txt -mt ">100"
...
:: 進行状況: [90821/90821] :: ジョブ [1/1] :: 3225 要求/秒 :: 期間: [0:00:42] :: エラー: 0 ::
厳密なアプローチとは程遠いですが、その差は非常に小さく (約 5%)、ボトルネックはネットワーク速度になると思います。
私がどれを選んだか当ててみてください。すべてを手作業で行うことは、テクノロジーの限界を探る素晴らしい方法です。私はたくさんのことを学び、実際に手を動かしてきましたが、今では自分に何が欠けているかが分かりました。
サイトの再構築にコミットしたくなかったので、生のパケットをキャプチャするための標準形式である PCAP を収集することを選択しました。
必須の警告: これは運用環境での使用を目的としたものではなく、またそのように構築されたものでもありません。それは壊れやすく、仮定が多く、拡張性がなく、バグがあります。セキュリティの話を始めさせないでください。作成者はこれを root として実行しています。
簡単で汚い方法は次のように実装できます。
#!/bin/bash
IF=eth0
IP=$(ip アドレス show dev $IF | grep inet | grep -o 'inet [0-9]*\.[0-9]*\.[0-9]*\.[0-9]*' | sed 's/inet //')
fname=$(日付 +%s).pcap
tcpdump -ni $IF "dst $IP およびポート 443 および (tcp[((tcp[12] & 0xf0) >> 2)] = 0x16)" -U -w /root/pcaps/$fname
必要に応じて、これを systemd サービスにラップすると、すぐに PCAP の収集を開始できます。
tcpdump フィルターは、フィルターとして設定できる ID を持つクライアント hello パケットのみをターゲットにします。追加の利点として、ディスク容量を使い果たすことなく大量のログを収集できることが挙げられます。
このアプローチではユーザー エージェントをキャプチャできません。これらの場合は、タイムスタンプと IP アドレスを使用して nginx アクセス ログとの関連付けを行う必要があります。不快ですが、PoC ではその必要はありません

もっと良くしてください。
このアプローチのもう 1 つの欠点は、暗号化された Client Hello パケットを処理できないことです。こうして私は ECH のウサギの穴に落ち、記事を持って戻ってきました。幸いなことに、私は ECH を使用していないので、この問題を回避できます。
仕様に従っていれば、JA4 ハッシュの計算はそれほど難しくありません。 Python データクラス形式で必要なものは次のとおりです。
@データクラス
クラスJA4パケット:
プロトコル: ConnProtocol
tls_version: TLSVersion
sni: SNI
暗号: list[CipherSuite] = フィールド(default_factory=lambda: [])
拡張機能: list[拡張機能] = フィールド(default_factory=lambda: [])
署名アルゴリズム: リスト[署名アルゴリズム] = フィールド(default_factory=lambda: [])
alpn: ALPN = ALPN("")
ec: bool = False
問題となるのは、そのデータを抽出、検証し、アルゴリズムに取り込むことです。
すべての RFC を調べて PCAP 解析を最初から実装すれば、PCAP 解析に関する一連の記事を書くことができます。代わりに scapy を使用することにしました。また、参照を容易にし、コードを明確にするために (または少なくともその一部を) 拡張 ID のリストを見つけて準備することをお勧めします。
scapy.all インポートから *
# 既知の拡張機能 ID リストはあなたの友達です
# https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml
ext_ids = {
"サーバー名": 0、
「サポートされるグループ」: 10、
"ec_point_formats": 11、
「署名アルゴリズム」: 13、
"application_layer_protocol_negotiation": 16、# ALPN
「サポートされるバージョン」: 43、
"ech_outer_extensions": 64768、
"encrypted_client_hello": 65037、
}
TLS バージョンの抽出
技術的な複雑さと歴史的な問題を説明するための例として、TLS バージョン抽出の問題を考えてみましょう。簡単そうに聞こえますよね？間違っています:
def get_tls_version(パケット):
tlv = パケット["TLS"].msg[0].version
tlv_e = [
x for x in packet["TLS"].msg[0].ext if x.type

== ext_ids["supported_versions"]
】
len(tlv_e) > 0 かつ len(tlv_e[0].versions) > 0 の場合:
tlv = max(tlv_e[0].versions)
TLVを返す
まず、パケットで指定された TLS バージョンを検索します。今日は、3 つの相反する価値観が見つかります。
TLS レイヤーの TLS バージョン: TLS 1.0 (0x0301)
Client Hello の TLS バージョン (レガシー TLS バージョン): TLS 1.2 (0x0303)
「supported_versions」でクライアントがサポートする TLS バージョン (RFC8846 以降)
実際のバージョンは、サーバーとネゴシエートされたバージョンになります。つまり、ここでは、サーバーが常に利用可能な最も高い有効な TLS バージョンを受け入れると仮定して、ショートカットを行っています (GREASE 値は省略します。もちろん 2 )。
ポイント A からポイント B までデータを取得する必要があります。袖をまくって完了しましょう。
def packet_to_ja4packet(packet) -> JA4Packet:
tlv = get_tls_version(パケット)
cps = packet["TLS"].msg[0].ciphers
sni_l = [x for x in packet["TLS"].msg[0].ext if x.type == ext_ids["server_name"]]
sni = SNI.sni() if len(sni_l) > 0 else SNI.no_sni()
exs = [パケット["TLS"]のxのx.type.msg[0].ext]
sig_sch_l = [
x for x in packet["TLS"].msg[0].ext if x.type == ext_ids["signature_algorithms"]
】
sig_algs = sig_sch_l[0].sig_algs if len(sig_sch_l)>0 else []
alpn_l = [
×
packet["TLS"].msg[0].ext の x の場合
if x.type == ext_ids["application_layer_protocol_negotiat

[切り捨てられた]

## Original Extract

I want to identify/block/rate-limit bots coming to my server. I want to avoid vendor lock-in and use open-source solutions. I also want to avoid PCAP Inspection (Suricata, Zeek...) as those are threatened by Encrypted Client Hello (and are fiddly/expensive to setup/maintain).
JA4 fits these requirements. In the previous part , I’ve covered JA3/JA4 generally:
Both are TLS fingerprinting methods analyzing ClientHello parameters (cipher suites, extensions, ALPN...).
JA3 is now obsolete since Chrome permutes their extensions , so please use JA4 which is now well supported.
JA4+ (the full suite of methods) has a mixed licensing model, beware.
In this part, let’s deploy this tech. There’s a somewhat-maintained list of tools supporting JA4 to choose from. Keeping the requirements above, I’ve found only these options:
Nginx plugin : made by JA4 authors but marked as work in progress,
Rama proxy : in rust, but still in alpha,
Envoy proxy : not on the list, but seems to support JA4 .
HAProxy plugin : my pick, in this post I describe how to deploy/configure HAProxy to store JA4s and User-Agents in logs.
If you really want to dive deep, I’ve implemented JA4 calculation in Python using scapy to parse PCAPs captured by tcpdump as a PoC. It’s not production ready at all, but I've learned a ton about the technology. Parsing PCAPs is the hard part, calculating JA4 from the extracted data is much easier.
Still interested? Let’s get started with the gauntlet of standards and acronyms.
Pre-Requisites, Technologies, RFCs and Shorthands
I’m about to drown you in terminology as we need to juggle a lot of the standards to get this thing rolling. So here’s a brief explainer:
TLS (Transport Layer Security) RFC8446 : the encryption layer for HTTP and many other protocols. There are many extensions to this.
SSL (Secure Socket Layer): previous version of TLS, quite compatible with TLS. Sometimes you still see references or the SSL/TLS shorthand for HTTPS.
Client Hello : First message from client to server to establish TLS.
SNI (Server Name Identification) RFC6066 : the client includes this extension in the client hello to tell the server which backend to connect to.
ALPN (Application Layer Protocol Negotiation) RFC7443 : client hello extension, it’s a hint to the server explaining what protocol to use after TLS is established.
GREASE (Generate Random Extensions to Sustain Extensibility) RFC8701 : enables clients to troll machines with selected nonsense values at random places they have to handle correctly and not crash.
ECH (Encrypted Client Hello) RFC9849 : new standard to encrypt that client hello to hide SNI and ALPN and others from prying eyes. My article on this .
PCAP (Packet Capture): format for storing packet data, de-facto standard established by libpcap implementation in the 80s. No formal RFC yet.
These points are really just a reference to keep as you navigate these shorthands and standards. And now let’s deploy some JA4.
You can reach out to one of the many big service providers to do this for you. Here’s a list , pick your favorite.
Obviously, that’s not something I’d like to do for this blog, the whole purpose of which is to do something independent. Focusing on tech, I want to avoid vendor lock-in as much as possible.
For example Cloudflare gives you this only in the Enterprise plan , the best plan there is! While you can expect some price increases, if you already have these plans, it’s a no-brainer. Even upgrading the plan might be cheaper than switching providers and/or rolling your own.
From the same list we can pick a few open source solutions that can support JA4.
A lot of these solutions capture pcaps to analyze the traffic (Suricata, Zeek, Wireshark, Arkime...). Beware if you’d like to use Encrypted Client Hello as the ClientHello properties you’d need will be... encrypted.
I’d rather deploy it on a reverse-proxy. There is an “official” nginx module from the JA4 creators , but from their descriptions it’s a work in progress. Since they’ve recently hidden their DB , I’m not sure what is the future of this extension.
There is also rama proxy , which looks amazing (new, rust, feature set) but has LLMisms all over the place. I’m not hating on the project, but I’m not using it either.
Speaking of rust-based proxies, pingora seems to be on its way to support JA4 , so fingers crossed.
And Envoy seems to be already there , but it's not on the list. Go figure. If you have a favorite proxy/server, check their docs first, maybe they actually support it.
Eliminating the above is how I’ve picked the HAProxy JA4 plugin .
I need haproxy, but at least 3.1 which is not by default in debian 13, so head here for instructions on how to include the repo. I need also a sample app with uvx fastapi-new hw . Of course, I need to do the standard stuff 1 .
cd /etc/haproxy
mkdir lua; cd lua
wget https://raw.githubusercontent.com/O-X-L/haproxy-ja4-fingerprint/refs/heads/latest/ja4.lua
We'll also need to set this up in the config file:
global
...
# JA4 related stuff
tune.ssl.capture-buffer-size 2048
lua-load /etc/haproxy/lua/ja4.lua
frontend https-in
bind *:443 ssl crt /etc/haproxy/certs/haproxy-test.local.pem
default_backend hello_world
http-request lua.fingerprint_ja4
http-request capture var(txn.fingerprint_ja4) len 36
http-request capture req.hdr(user-agent) len 150
backend hello_world
server app1 127.0.0.1:8000 check
First we need to set these checks and TLS termination on the frontend as well as route the traffic to correct backend.
We need to tell haproxy to capture SSL/TLS info and load the plugin. Small note on the tune.ssl.capture-buffer-size 2048 - the extension author recommends to set this to 192, but as we’re trying to fit basically the entire ClientHello into here we can consider post-quantum ciphers like Kyber , which can have a signature that's 1568 bytes long!
Then we need to tell the front end to compute JA4 and capture it along with the User-Agent to logs.
May 29 20:03:42 haproxy-test haproxy[3319]: [NOTICE] (3319) : Loading success.
May 29 20:03:42 haproxy-test systemd[1]: Started haproxy.service - HAProxy Load Balancer.
May 29 20:06:18 haproxy-test haproxy[3322]: 192.168.100.1:46182 [29/May/2026:20:06:18.238] https-in~ hello_world/app1 0/0/0/1/1 404 153 - - ---- 1/1/0/0/0 0/0 {t13d1715h2_5b57614c22b0_7121afd63204|Mozilla/5.0 (X11; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0} "GET https://haproxy-test.local/test HTTP/2.0"
You can see in the curly braces that we have both the JA4 and the User-Agent captured for further processing.
From here, you can do whatever. Capture the logs for later checks, forward them to some IDS/IPS, set headers for the application.. I'm sure your favourite LLM will generate a script to capture what you need - if not, here's my cronjob script that needs to run every hour .
But oh no, Lua is interpreted it’s so slow, single-threaded, oh no, this will kill performance... Well, I ran a simple test:
Client: ffuf (40 threads by default) with this wordlist .
Network: No network issues - localhost connections.
Once with JA4 collection disabled and once with JA4 collection enabled.
$ ffuf -u https://haproxy-test.local/FUZZ -w fav-wordlist.txt -mt ">100"
...
:: Progress: [90821/90821] :: Job [1/1] :: 3278 req/sec :: Duration: [0:00:40] :: Errors: 0 ::
$ ffuf -u https://haproxy-test.local/FUZZ -w fav-wordlist.txt -mt ">100"
...
:: Progress: [90821/90821] :: Job [1/1] :: 3225 req/sec :: Duration: [0:00:42] :: Errors: 0 ::
While it’s far from a rigorous approach, I think the difference is quite small (~5%) and your bottleneck will be the network speed.
Guess which one I’ve picked. Doing everything by hand is a great way of exploring the limits of the technology. I’ve learned a ton, got my hands dirty and now I know what am I missing.
Since I didn’t want to commit to restructuring my site, I opted for collecting PCAPs, which is a standard format for capturing raw packets.
Obligatory warning: this is not intended nor built for production usage. It’s fragile, riddled with assumptions, doesn’t scale and has bugs. And don’t get me started on security - the author runs this as root!
One quick’n’dirty way can be implemented like so:
#!/bin/bash
IF=eth0
IP=$(ip address show dev $IF | grep inet | grep -o 'inet [0-9]*\.[0-9]*\.[0-9]*\.[0-9]*' | sed 's/inet //')
fname=$(date +%s).pcap
tcpdump -ni $IF "dst $IP and port 443 and (tcp[((tcp[12] & 0xf0) >> 2)] = 0x16)" -U -w /root/pcaps/$fname
If you’d like, wrap it in a systemd service and we can start collecting PCAPs right away.
The tcpdump filter targets only client hello packets, which have an ID we can set as filter. As an added benefit, we can collect a lot of logs without exhausting the disk space.
This approach cannot capture User-Agents. For those you’d need to do correlation against nginx access logs using timestamps and IP addresses. It’s unpleasant, but for a PoC I don’t need to do better.
Another downside of this approach is that you can’t process Encrypted Client Hello packets. This is how I fell into the ECH rabbit hole and came back with an article . Fortunately, I’m not using ECH, so I can get away with this.
Calculating JA4 hashes is not that hard if you follow the spec . Here’s roughly what you need, in the python dataclass form:
@dataclass
class JA4Packet:
protocol: ConnProtocol
tls_version: TLSVersion
sni: SNI
ciphers: list[CipherSuite] = field(default_factory=lambda: [])
extensions: list[Extension] = field(default_factory=lambda: [])
signature_algorithms: list[SignatureAlgorithm] = field(default_factory=lambda: [])
alpn: ALPN = ALPN("")
ech: bool = False
What’s problematic is extracting, validating and then getting that data into your algorithm.
I could write a series of articles on PCAP parsing as I go through all of the RFCs to implement this from the scratch. I’ve opted to use scapy instead. I’d also recommend finding and preparing a list of extension IDs for easier reference and code clarity (or at least some bits of it):
from scapy.all import *
# known extension ids list is your friend
# https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml
ext_ids = {
"server_name": 0,
"supported_groups": 10,
"ec_point_formats": 11,
"signature_algorithms": 13,
"application_layer_protocol_negotiation": 16, # ALPN
"supported_versions": 43,
"ech_outer_extensions": 64768,
"encrypted_client_hello": 65037,
}
TLS Version Extraction
As an example to illustrate the technical complexity and historical baggage, let’s consider the issue of TLS version extraction. Sounds simple, right? Wrong:
def get_tls_version(packet):
tlv = packet["TLS"].msg[0].version
tlv_e = [
x for x in packet["TLS"].msg[0].ext if x.type == ext_ids["supported_versions"]
]
if len(tlv_e) > 0 and len(tlv_e[0].versions) > 0:
tlv = max(tlv_e[0].versions)
return tlv
First you search for the TLS version specified in the packet. Today, you’ll find three conflicting values :
TLS version of the TLS layer: TLS 1.0 (0x0301)
TLS version of the Client Hello (legacy TLS version): TLS 1.2 (0x0303)
TLS versions supported by the client in the “supported_versions” (since RFC8846 )
The actual version is then the one negotiated with the server. That means I’m taking a shortcut here by assuming the server will always accept the highest available valid TLS version (omit the GREASE values, of course 2 ).
We need to get data from point A to point B. Let’s just roll up the sleeves and be done with it:
def packet_to_ja4packet(packet) -> JA4Packet:
tlv = get_tls_version(packet)
cps = packet["TLS"].msg[0].ciphers
sni_l = [x for x in packet["TLS"].msg[0].ext if x.type == ext_ids["server_name"]]
sni = SNI.sni() if len(sni_l) > 0 else SNI.no_sni()
exs = [x.type for x in packet["TLS"].msg[0].ext]
sig_sch_l = [
x for x in packet["TLS"].msg[0].ext if x.type == ext_ids["signature_algorithms"]
]
sig_algs = sig_sch_l[0].sig_algs if len(sig_sch_l)>0 else []
alpn_l = [
x
for x in packet["TLS"].msg[0].ext
if x.type == ext_ids["application_layer_protocol_negotiat

[truncated]
