---
source: "https://patrickmccanna.net/inspecting-claude-codes-network-traffic-with-linux-namespaces-and-mitm-proxying-part-1/"
hn_url: "https://news.ycombinator.com/item?id=48746915"
title: "Using network namespaces to discover how Claude Code scrapes"
article_title: "Inspecting Claude Code’s Network Traffic with Linux Namespaces and MITM Proxying (Part 1) – blog"
author: "0o_MrPatrick_o0"
captured_at: "2026-07-01T14:59:47Z"
capture_tool: "hn-digest"
hn_id: 48746915
score: 2
comments: 0
posted_at: "2026-07-01T13:59:20Z"
tags:
  - hacker-news
  - translated
---

# Using network namespaces to discover how Claude Code scrapes

- HN: [48746915](https://news.ycombinator.com/item?id=48746915)
- Source: [patrickmccanna.net](https://patrickmccanna.net/inspecting-claude-codes-network-traffic-with-linux-namespaces-and-mitm-proxying-part-1/)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T13:59:20Z

## Translation

タイトル: ネットワーク名前空間を使用してクロード コードがどのようにスクレイピングするかを発見する
記事のタイトル: Linux 名前空間と MITM プロキシを使用した Claude Code のネットワーク トラフィックの検査 (パート 1) – ブログ

記事本文:
Linux 名前空間と MITM プロキシを使用した Claude Code のネットワーク トラフィックの検査 (パート 1) – ブログ
コンテンツにスキップ
ブログ
モバイル セキュリティのメンターシップとテクノロジー分野で働くよう子供たちを鼓舞する
Linux 名前空間と MITM プロキシを使用した Claude Code のネットワーク トラフィックの検査 (パート 1)
これは、Claude Code が Web スクレイピング リクエストを開始する方法の分析です。クロード コードが API エンドポイント、またはクロード コードを実行しているホストからスクレイピングできることに驚かれるかもしれません。このシリーズでは、ネットワーク名前空間、tcpdump、MITMproxy を使用してクライアントのネットワーク トラフィックを分離し、クロード コードが API エンドポイントからコンテンツをスクレイピングしているのか、それともクロード コードを実行しているホストからコンテンツをスクレイピングしているのかを実験的に検出する方法を学びます。このプロジェクトのコードは私の github で入手できます。
クロードコードをインストールしました。ユーザーに代わってインターネットに接続できる組み込みネットワーク ツール (WebFetch および WebSearch) が付属しています。使用するたびに明示的な承認が必要です (自動承認ポリシーを構成していない場合)。しかし、ドキュメントでは誰も答えていない質問があります。
正確には何を送信し、どこからリクエストを送信しているのでしょうか?
Linux 名前空間と MITM プロキシを使用した Claude Code のネットワーク トラフィックの検査
プロキシフェッチモデルでは、リクエストは Anthropic のサーバーを経由してルーティングされます。ターゲット システムは Anthropic の IP を認識し、Anthropic のインフラストラクチャがフェッチされたコンテンツを処理します。
これら 2 つのモデルには、プライバシーとセキュリティへの影響が大きく異なります。クロードコードは実際にどれを使用しますか?このガイドでは、その質問に経験的に答えます。
クロード コードが「ウェブを検索」するとき、検索エンジンはあなたの IP アドレス、または Anthropic の IP アドレスを認識しますか?どのようなテレメトリが家に送信されますか?
これらは仮説上の懸念ではありません。 AI エージェントを実行している場合

内部サービスにアクセスできるネットワーク ツールがマシンに組み込まれている場合は、そのトラフィック パターンを知る必要があります。このガイドでは、この動作を直接発見する方法と、分析中に発見したことを説明します。
このガイドを終えると、次のことができるようになります。
ネットワーク名前空間を使用して、Linux 上で単一プロセスのネットワーク トラフィックを分離します。この手法により、キャプチャに他のプロセスのパケットが含まれないようになります。
SNI および DNS メタデータを使用して、暗号化された TLS トラフィックから宛先ホスト名を抽出します
mitmproxy による MITM プロキシを使用して特定のプロセスの HTTPS トラフィックを復号化し、1 つのプロセスのみが影響を受けるように範囲を設定します
Claude CLI が生成するすべてのトラフィックを分類します: API 呼び出し、Web フェッチ、Web 検索、テレメトリ、MCP 接続
ローカルフェッチモデルのセキュリティへの影響を理解する: LAN と内部サービスにとってそれが何を意味するか
以下については快適に行うことができます。
Linuxターミナルとシェルスクリプト
基本的なネットワーク概念 (IP アドレス、ポート、TCP、DNS)
tcpdump の機能 (毎日使用しない場合でも)
TLS/HTTPS が提供するもの (暗号化されたトランスポート)
カーネル名前空間をサポートする Linux (最新のディストリビューション)
tcpdump 、 iproute2 ( ip コマンド)、 iptables
Python 3 (分析スクリプト用)
mitmproxy ( pipx install mitmproxy または pip install mitmproxy ) 復号化セクション用
slirp4netns パケットを収集するためのルートレス アプローチ (ほとんどのディストリビューション リポジトリで利用可能)
git clone git@github.com:CaptainMcCrank/claude-netns.git
cd クロードネットンズ
chmod +x クロード-sandbox.sh クロード-mitm.sh rootless-capture.sh mitm-capture.sh
始める前の重要な概念
このガイドは、馴染みのない、または直感に反する可能性がある Linux ネットワークと TLS の概念に依存しています。
Linux にはプロセスごとのパケット キャプチャはありません。 Linux のパケット キャプチャ メカニズム ( AF_PACKET ) はインターフェイスに接続されます

プロセスではなく ces 。 tcpdump --pid はありません。ボックス上のすべてのプロセスで共有されるインターフェイスでキャプチャすると、すべてが混在してしまいます。このガイドの前半では、ネットワーク名前空間を使用したこの問題の回避策を構築します。
HTTPS はすべてのメタデータを非表示にするわけではありません。 TLS はペイロードを暗号化します。ペイロードには、URL パス、ヘッダー、本文が含まれます。しかし、宛先ホスト名はすべての ClientHello で TLS SNI 拡張機能を介して平文で漏洩し、DNS クエリによって標準の暗号化されていないルックアップを介して再び漏洩します。コンテンツを読むことはできませんが、トラフィックがどこに行くのかを確認できます。この部分的な可視性は、復号化をまったく行わなくても、最も重要な質問 (WebFetch はどこに接続するのか?) に答えるのに十分であることがわかります。
MITM プロキシは標準的なデバッグ手法です。敵対的なネットワーク上の攻撃者が使用するのと同じ TLS 傍受メカニズムが、Burp Suite、Charles Proxy、Fiddler、mitmproxy、および企業の TLS 検査アプライアンスによって毎日使用されています。これらのツールには正当なアプリケーションがあるにもかかわらず、多くの企業セキュリティ ツールはそれらをハッキング ツールとして扱います。 MITM プロキシは、検査する権限を与えられたネットワークおよびクライアントに対してのみ実行してください。このような技術に対する承認に疑問がある場合は、続行しないでください。これはすべて正当なデバッグと監査ですが、この作業が攻撃であるか、または許可されているかをサードパーティが独自に区別する方法はありません。多くの企業セキュリティ チームは、MITM プロキシなどのツールがインストールされているのを確認したときに警告を発するように監視ツールを構成しています。
WebFetch が実際にどこに接続するかは経験的な問題です。リクエストは Anthropic のサーバーを経由しますか? それともあなたのマシンから直接接続しますか?このガイドでは、パケット キャプチャでその質問に答えます。
すべてのトラフィックの一般的な tcpdump が失敗する理由
簡単なアプローチは次のとおりです。

メインインターフェイス上のトラフィックをキャプチャします。
sudo tcpdump -i eth0 -n -c 100 -w everything.pcap
別の端末で claude を使用しながらこれを実行するとキャプチャが生成されますが、ブラウザのタブ、システム アップデートのチェック、NTP 同期、SSH キープアライブ、ボックス上のすべてのアプリケーションからの DNS クエリなど、ノイズの壁が含まれています。クロードのトラフィックはどこかにありますが、それを見つけるには何百ものフローを手作業で分解する必要があります。
宛先でフィルタリングしてみるとよいでしょう。
sudo tcpdump -i eth0 -n host api.anthropic.com -w filtered.pcap
より良いですが、これにはすべての宛先 (api.anthropic.com など) を事前に知っておく必要があります。クロードがあなたが予測していない URL にアクセスした場合、あなたはそれを見逃してしまいます。さらに、他の Anthropic 統合を実行している場合は、他のプロセスも api.anthropic.com にアクセスする可能性があります。
tcpdump は、プロセスではなくネットワーク インターフェイスによってフィルタリングします。 Linux には tcpdump --pid 12345 はありません。カーネルのパケット キャプチャ メカニズム (AF_PACKET ソケット) はインターフェイスに接続され、メイン インターフェイス ( eth0 、 wlan0 など) はシステム上のすべてのプロセスからのトラフィックを伝送します。
建物全体の電話線をタップして、1 人の通話を録音しようとします。ブラウザ、アップデータ、バックグラウンド サービスなど、すべてのテナントの声が聞こえます。すべての音声は同じワイヤーを共有しているため、1 つの音声を分離する方法はありません。
クロードにトラフィックのみを伝送し、他には何も伝送しない独自のプライベート ネットワーク インターフェイスを提供する方法が必要です。パート 2 では、特定のプロセスのすべてのネットワーク トラフィックを分離できるようにする Linux ネームスペースの作成について説明します。

## Original Extract

Inspecting Claude Code’s Network Traffic with Linux Namespaces and MITM Proxying (Part 1) – blog
Skip to content
blog
Mobile Security Mentorship & Inspiring Kids to work in tech
Inspecting Claude Code’s Network Traffic with Linux Namespaces and MITM Proxying (Part 1)
This is an analysis of how Claude Code initiates web scraping requests. You might be surprised to find that claude code could scrape from an API endpoint, or from the host that’s running claude code. In this series, you will learn how to use network namespaces, tcpdump and MITMproxy to isolate a client’s network traffic and experimentally discover whether claude code is scraping content from an api endpoint or from the host that’s running claude code. Code for this project is available at my github.
You installed Claude Code. It ships with built-in network tools (WebFetch & WebSearch) that can reach out to the internet on your behalf. Each use requires your explicit approval (unless you’ve configured auto-approve policies). But here’s the question nobody answers in the documentation:
What exactly does it send and where is it sending requests from?
Inspecting Claude Code’s Network Traffic with Linux Namespaces and MITM Proxying
In the proxied fetch model, the request would route through Anthropic’s servers. The target system sees Anthropic’s IP, and Anthropic’s infrastructure handles the fetched content:
These two models have very different privacy and security implications. Which one does Claude Code actually use? This guide answers that question empirically.
When claude code “searches the web,” whose IP address does the search engine see- yours or Anthropic’s? What telemetry does it send home?
These aren’t hypothetical concerns. If you’re running an AI agent that has built-in network tools on a machine that can reach internal services, you need to know its traffic pattern. This guide shows you how you can discover this behavior directly and what I found during my analysis.
By the end of this guide, you’ll be able to:
Isolate a single process’s network traffic on Linux using network namespaces: this technique ensures you won’t have other process’s packets in the capture
Extract destination hostnames from encrypted TLS traffic using SNI and DNS metadata
Decrypt HTTPS traffic for a specific process using MITM proxying with mitmproxy, scoped so only that one process is affected
Classify all traffic the Claude CLI generates: API calls, web fetches, web searches, telemetry, and MCP connections
Understand the security implications of a local-fetch model: What it means for your LAN and internal services
You should be comfortable with:
Linux terminal and shell scripting
Basic networking concepts (IP addresses, ports, TCP, DNS)
What tcpdump does (even if you don’t use it daily)
What TLS/HTTPS provides (encrypted transport)
Linux with kernel namespace support (any modern distro)
tcpdump , iproute2 ( ip command), iptables
Python 3 (for the analysis scripts)
mitmproxy ( pipx install mitmproxy or pip install mitmproxy ) for the decryption sections
slirp4netns for the rootless approach to collecting packets (available in most distro repos)
git clone git@github.com:CaptainMcCrank/claude-netns.git
cd claude-netns
chmod +x claude-sandbox.sh claude-mitm.sh rootless-capture.sh mitm-capture.sh
Key Concepts Before We Start
This guide relies on Linux networking and TLS concepts that may be unfamiliar or counterintuitive.
There is no per-process packet capture on Linux. Linux’s packet capture mechanism ( AF_PACKET ) attaches to interfaces , not processes. There is no tcpdump --pid . If you capture on an interface shared by every process on the box, you get everything mixed together. The first half of this guide builds a workaround for this using network namespaces.
HTTPS does not hide all metadata. TLS encrypts the payload . The payload includes the URL path, headers, and body. But the destination hostname leaks in plaintext via the TLS SNI extension in every ClientHello, and DNS queries leak it again via standard unencrypted lookups. You can’t read the content, but you can see where traffic goes. That partial visibility turns out to be enough to answer the most important question (where does WebFetch connect?) without any decryption at all.
MITM proxying is a standard debugging technique. The same TLS interception mechanism used by attackers on hostile networks is used every day by Burp Suite, Charles Proxy, Fiddler, mitmproxy, and corporate TLS inspection appliances. Despite the fact that these tools have legitimate applications, many corporate security tools treat them as hacking tools. Only perform MITM Proxying on networks and clients you’re authorized to inspect. Do not proceed if you have any doubts about your authorization for such techniques. While this is all legitimate debugging and auditing, there’s no way for a 3rd party to independently distinguish if this work is an attack or authorized. Many corporate security teams configure monitoring tools to alarm when they see tools such as MITM proxy installed.
Where WebFetch actually connects is an empirical question. Does the request go through Anthropic’s servers, or does it connect directly from your machine? This guide answers that with packet captures.
Why a general tcpdump of all traffic fails
The straightforward approach is to capture traffic on your main interface:
sudo tcpdump -i eth0 -n -c 100 -w everything.pcap
Running this while using claude in another terminal produces a capture, but it contains a wall of noise: browser tabs, system update checks, NTP syncs, SSH keepalives, DNS queries from every application on the box. Claude’s traffic is in there somewhere, but finding it requires picking apart hundreds of flows by hand.
You might try filtering by destination:
sudo tcpdump -i eth0 -n host api.anthropic.com -w filtered.pcap
Better- but this requires you to know every destination in advance (e.g. api.anthropic.com). If claude contacts a URL you didn’t predict, you miss it. Additionally- other processes might also contact api.anthropic.com if you have other Anthropic integrations running.
tcpdump filters by network interface, not by process. There is no tcpdump --pid 12345 on Linux. The kernel’s packet capture mechanism ( AF_PACKET sockets) attaches to interfaces and your main interface ( eth0 , wlan0 , etc.) carries traffic from every process on the system.
Trying to record one person’s phone calls by tapping the entire building’s phone line. You hear every tenant: browsers, updaters, background services. There’s no way to isolate one voice because they all share the same wire.
We need a way to give claude its own private network interface that carries only its traffic and nothing else. In Part 2 , we’ll explore creating a Linux Namespace that enables us to isolate all network traffic for a specific process.
