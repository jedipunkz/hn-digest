---
source: "https://blog.includesecurity.com/2026/06/the-smart-tv-in-your-livingroom-is-a-node-in-the-aiscraping-economy/"
hn_url: "https://news.ycombinator.com/item?id=48420264"
title: "The smart TV in your living room is a node in the AI scraping economy"
article_title: "The Smart TV in Your LivingRoom Is a Node in the AIScraping Economy - Include Security Research Blog"
author: "themaxdavitt"
captured_at: "2026-06-06T00:54:13Z"
capture_tool: "hn-digest"
hn_id: 48420264
score: 1
comments: 0
posted_at: "2026-06-06T00:50:50Z"
tags:
  - hacker-news
  - translated
---

# The smart TV in your living room is a node in the AI scraping economy

- HN: [48420264](https://news.ycombinator.com/item?id=48420264)
- Source: [blog.includesecurity.com](https://blog.includesecurity.com/2026/06/the-smart-tv-in-your-livingroom-is-a-node-in-the-aiscraping-economy/)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T00:50:50Z

## Translation

タイトル: リビングルームのスマート TV は AI スクレイピング エコノミーのノードです
記事のタイトル: リビングルームのスマート TV は AIScraping 経済のノードです - セキュリティ研究ブログを含む
説明: この投稿では、BrightData の SDK の内部と、それが一般の消費者向け TV を、AI 業界が Web データをスクレイピングして言語学習モデルをトレーニングするために活用する巨大な商用、住宅用プロキシ ネットワークの出口ノードにどのように変えるのかを見ていきます。

記事本文:
リビングルームのスマート TV は AIScraping 経済のノードです - セキュリティ研究ブログを含む
コンテンツにスキップ
チームリサーチのブログ
リビングルームのスマート TV は AIScraping 経済の結節点です
Include Security の仕事では、私たちは毎日 AI と取り組んでいます (ハッキング、使用、トレーニングなど)。
AI 機能の向上を目的として最近建設されたデータセンターに対してコミュニティレベルの反対が起きていることは誰もが知っています。あなたは気づいていないかもしれませんが、家の中のデバイスを使用して AI をトレーニングするための分散的な取り組みが行われている可能性があります。
この投稿では、Bright Data 社が、自社の住宅用プロキシ ネットワークを使用してインターネットからトレーニング データをスクレイピングする最新の AI モデルをどのように促進するかを探っていきます。
Bright Data は、顧客が Web スクレイピング トラフィックをルーティングする、4 億以上のホーム IP アドレスからなる世界最大の住宅用プロキシ ネットワークとして販売しているものへのアクセスを販売するデータ収集会社です。このネットワークの背後にある供給は SDK から来ます。SDK は、ユーザーの同意を得て、携帯電話やスマート TV を出口ノードの 1 つに変える、消費者向けアプリに組み込まれたソフトウェアです。
携帯電話やスマート TV などのシステム上でこの会社の SDK が何を行うかについて、平均的なユーザーが知っておくべきことを文書化します。私たちは、SDK がどのように機能するのか、どのプラットフォームが SDK を提供しているのか、そしてなぜインターネットに接続されたテレビが、インターネットから収集したデータでトレーニングしようとしている AI モデルの究極のプロキシであるのかを探っていきます。
AI 企業は、事前トレーニング、検索、エージェントのグラウンディング、検索など、Web スクレイピングされたコンテンツに依存しています。しかし、最新の Web はデータセンターから収集できません。 Cloudflare、DataDome、HUMAN などは、既知のクラウド IP からのリクエストを抑制またはブロックします。
回避策は次のとおりです

居住用プロキシ。 Comcast または T-Mobile 加入者の接続を介してルーティングされたスクレイピング ジョブは、有料の住宅顧客に属する IP からターゲット サイトに到着します。クレブス氏は2025年10月、「Aisuruやその他のソースからの過剰なプロキシが、さまざまなAIプロジェクトに関連した大規模なデータ収集の取り組みを促進している」と報告した。 2019 年に遡る学術的な測定では、これらのネットワークが圧倒的に悪用されていることが示されています。 FBIは今年初めに正式な勧告を出した。
既存のマスコミのほとんどは、ボットネット ( Aisuru 、 Kimwolf )、トロイの木馬化アプリ ( HUMAN Security の PROXYLIB 開示 )、感染済みの IoT ハードウェア ( Google/Mandiant の IPIDEA 削除 ) など、違法な住宅用プロキシの供給に焦点を当てています。これらは悪役です。
一方で、合法的な供給側はあまり精査されていません。現在、Bright Data は独自のマーケティングによって世界最大の住宅用プロキシ ネットワークとなり、パートナー アプリに埋め込まれた同意 SDK を介してソースされた「1 億 5,000 万以上の IP」を宣伝しています。この調査では、その SDK がどのように機能するか、どのプラットフォームに同梱されているか、そしてコネクテッド TV が究極の家庭用プロキシである理由が文書化されています。
Connected TV (CTV) が理想的なプロキシである理由
コネクテッド TV、別名スマート TV は、ほぼ完璧な住宅用プロキシです。携帯電話と比較すると:
ユーザーが眠っているときにテレビのバッテリーが 1% に達したり、WiFi ネットワーク間を移動したり、ロックされたりすることはありません。一部のパートナー パブリッシャーは、プライバシー ポリシーで Bright Data との関係を開示しています。PlayWorks はその一例です。しかし、プライバシー ポリシーの開示は、テレビにとって間違ったコントロール サーフェスです。リモコンの矢印キーで法的文書をスクロールするのは難しく、アプリ内の同意ダイアログでは、Bright Data の有料顧客がスクレイピング トラフィックをユーザーの自宅のインターネット経由でルーティングしようとしていることが伝わりません。

。
The Verge によって文書化されている Roku アプリである Petflix は、その代表的なケースです。そのオプトイン画面には、「広告を減らして Petflix を無料で楽しむため、ブライト データがデバイスの無料リソースと IP アドレスを使用して、インターネットから公開 Web データをダウンロードすることを許可しています。ブライト データは、承認されたビジネス関連のユースケースにのみあなたの IP アドレスを使用します。あなたの個人情報は、あなたの IP アドレス以外にアクセスまたは収集されません。期間」と書かれています。 Petflix ダイアログには「時々」と表示されます。 SDK の公開クエリ可能な設定セット max_bw_monthly_wifi: 200,000,000,000 バイト - 200 GB のデフォルトの月間 WiFi 予算。
Who Bright Dataがパートナーとして名を連ねる
Bright Data はパートナー マニフェスト エンドポイントを公開します。エンドポイントは認証されていないため、誰でも取得できます。公的情報源から高い信頼性を持って特定できたマニフェスト内の名前は次のとおりです。
その他 ( desoline、free_time、ott_studio、global_microtrading、m_m_media、easystaff_lp ) は存在しますが、公開ソースからは特定できません。 Bright_screensavers、bright_videos、brightdata は Bright Data 独自のアプリです。
パートナー リストの証明に関するメモ: Bright Data の構成にリストされているということは、ある時点で統合が存在していた可能性があることを意味します。これ自体では、特定の発行元が現在出荷しているアプリに実稼働中の SDK が含まれていることを証明するものではありません。指定された発行者については、アプリごとの検証が必要です。
パートナー リストが直接証明していること:
Bright Data は、この名簿を認証されていないパブリック エンドポイントに送信します。
少なくとも 3 つの CTV を中心とした事業体 (PlayWorks、CloudTV、Longvision) が、ユーザーのデバイスを住宅用プロキシ出口ノードとして収益化していました。特に PlayWorks は、主要な TV プラットフォームと ISP にわたる CTV 配信を報告しており、リーチ数は 1 件あたり数億世帯に上ります。

独自のマーケティング資料。
Bright Data SDK はどのようにしてユーザーのデバイスを住宅用プロキシ出口ノードに変えるのでしょうか?
Bright Data SDK は公的に文書化された商用製品であり、Bright Data の SDK 統合ドキュメント (Web 用の JavaScript バリアントを含む) を通じてパブリッシャーに提供されます。以下の内容は、出荷されている iOS フレームワークのリバース エンジニアリングと 30 日間のランタイム トラフィックの計測から得られた結果を基に、公開された表面に基づいて構築されます。
SDK は、パートナー アプリ内の iOS フレームワーク (brdsdk.framework) として出荷されます。私はバイナリをリバース エンジニアリングし、同意を得てインストールされたパートナー アプリ内で SDK を実行している調査フリートから 30 日間のトラフィックをキャプチャしました。
起動するたびに、SDK は以下を呼び出します。
GET <https://clientsdk.bright-sdk.com/sdk_config_ios.json>?appid=<bundle>&ver=<sdk-version>&uuid=sdk-ios-<32hex>
エンドポイントは意味のある意味で認証されていません。サーバーは、appid (パートナー アプリの App Store リストにあるアプリ バンドル ID) と ver (SDK バージョン文字列) の 2 つのクエリ パラメーターのみをゲートします。これらの UUID とランダムに生成された UUID を指定すると、サーバーは実際のデバイスが取得するのと同じ応答を返します。機能フラグ、アイドル検出しきい値 (バッテリー %、CPU/メモリの上限、WiFi 対セルラーのルール)、国ごとの帯域幅層、上で紹介したパートナー マニフェストなどです。これらの各ブランチは、デバイスがリレーできる時期を決定するアイドル ルール、VPN 経由でピア トラフィックをルーティングするフラグ、プラットフォーム間のインストールを 1 つの ID に統合するマップ、および国ごとの帯域幅の上限など、それぞれについて検討する価値があります。
構成のフェッチ後、SDK は永続的な WebSocket を開き、次のことを行います。
このホスト名は AWS Global Accelerator IP (この記事の執筆時点では 3.33.193.183、15.197.193.114) に解決されます。 TLS 証明書は CN=*.luminatinet.com です。

ブライトデータの2018年以前の社名であるLuminati Networksのメイン。ブランド変更は 2018 年に正式に発表されました。アクティブな SDK インフラストラクチャは依然としてレガシー証明書で実行されており、これは有用な検出ピボットです。現在の顧客向けプロキシ サービスは、brightdata.com ブランドのドメイン上に存在するため、ネットワーク上の luminatinet.com / brdtnet.com トラフィックは、顧客側の Bright Data の使用ではなく、具体的にはピア トンネル プレーンになります。サーバーは自身を uWebSockets: 20 として識別します。
ピア エンドポイントのアップグレードには認証は必要ありません。サーバーは TLS 有効な WebSocket アップグレードを受け入れ、クライアントのパブリック IP がエコーバックされたアプリケーション層フレームを接続中のクライアントに即座にプッシュします。そこから握手が始まります。
サーバー → クライアント:tunnel_init はセッションを確立し、クライアントのパブリック IP を返します。
サーバー → クライアント: cid_set サーバーはクライアントに <IP>-<token>/ls<N>c<M>p443_<IP>_<counter> 形式のセッション追跡識別子を割り当てます。この形式が、実際のデバイスから SDK でキャプチャされたテレメトリ トラフィックに存在する cid フィールドと一致することを確認しました。
サーバー → クライアント: status_get サーバーは、アイドル状態、バッテリー、ネットワーク タイプ、および利用可能な帯域幅についてデバイスをポーリングします。デバイスは、継続的なテレメトリ フィードで応答します: idle、wifi_connected、mobile_connected、mobile_type (LTE/5G)、roaming、battery_level、using_battery、screen_on、on_call、cpu_usage、mem_usage、raw_bw、bw、ipv6_supported、appid (ホスト アプリ)、sdk_version、platform 、および割り当てられた cid 。これは、物理デバイスの状態をサードパーティに継続的にフィードするもので、ホスト アプリの発行者によってテキストが選択される同意ダイアログを介して配信されます。
握手完了。デバイスが良好なステータスを報告すると、サーバーのジョブマッチング層は自由に cmd_tun フレームをプッシュできるようになります。個別のスクレイピング-j

ob 命令は、SDK がユーザーの住宅 IP をソースとして使用して、サードパーティのサイトに対する HTTP リクエストとして実行する命令です。
WebSocket 上のすべてのフレームは、固定エンベロープを持つプレーンな JSON です。
{"タイプ": "ipc_call"|"ipc_post"|"ipc_result"|"ipc_error",
"cmd": <コマンド>、"cookie": <相関 ID>、
"err_code": 0、"msg": { ...ペイロード... }}
バイナリから抽出され、回線上で検証された完全なコマンド語彙:
メッセージ署名、HMAC、クライアント証明書、デバイス認証はありません。 TLS 層とサーバーの IP レピュテーション フィルターのみが、どのピアが実際にジョブを受信するかを制御します。商用マルウェア プロトコルの設計に詳しい読者の場合: これは、一般的な C2 よりも安全性が大幅に低くなります。
この構成には、デバイスが他の人のトラフィックを中継できる場合に関する明示的なルールブックが同梱されています。
"アイドルメトリクス": {
"ignore_screen_on": true, // 画面がオンでも中継します
"ignore_on_call": true, // ユーザーが通話中に中継する
「最大帯域幅比」: 1、
"min_battery": 0.2、
"wifi_on_battery": true、
「min_battery_wifi」: 0.2、
「max_cpu_usage」: 70、
"max_mem_usage": 90、
"mem_screen_off": true、
「アイドルタイムアウト」: 30、
「not_idle_timeout」: 10
}
ignore_screen_on フラグとignore_on_call フラグは注目に値します。「アイドル」はユーザーがデバイスから離れていることを意味するわけではありません。これは、デバイスの CPU、メモリ、バッテリーが SDK のしきい値内にあることを意味します。電話中のユーザーは、画面を積極的に読んでおり、中継目的ではアイドル状態であるとみなされます。
この構成には、dual_pairing マップも同梱されています。
"デュアルペアリング": {
"ios_com.brd.earnapp": ["win_earnapp.com", "mac_com.earnapp"]
}
これは、ユーザーの同じブランドの iOS、Windows、macOS インストールを 1 つのエンティティに結び付けるサーバー側のマップです。これは、クロスプラットフォームの ID ステッチングであり、公開設定ファイル内に文書化されています。

g フィールド: http3_enabled: true。 SDK は、QUIC ベースのピア トランスポート用のフラグをすでに出荷しています。将来のバージョンでは、ピア トンネルが TCP/443 から UDP/443 に移動される可能性があり、これにより、WebSocket の検出に TCP 接続追跡に依存する防御機能が機能しなくなります。
SDK の設定にはフラグ「use_netifs」: true が付属しています。このフラグは、システムのデフォルト ルートを使用するのではなく、特定の必要なインターフェイス en0 (WiFi) または pdp_ip0 (セルラー) を使用して NWConnection を構築する SDK バイナリ内のコードをトリガーします。
iOS では、これにより、設定された VPN の tun0 インターフェイスが完全にバイパスされます。アプリの残りの HTTPS トラフィックが通過する場合でも、ピア トンネルはユーザー設定の VPN を通過しません。
私たちはこれを経験的に観察しました。私の研究設定には透過的な TLS インターセプトが含まれています。ポート 443 が明示的にインスペクターにリダイレクトされているにもかかわらず、proxyjs.brdtnet.com:443 へのピア トンネルを除く、SDK が行ったすべての HTTPS 呼び出しがキャプチャされました。バイパスには、Apple の文書化された NWParameters.requiredInterface API が使用されます。
SDK は 2 つの独立した検査バイパス (プレーンごとに 1 つ) を使用していることを強調する価値があります。
コントロール プレーン (構成フェッチ、テレメトリ ping): URLSession/NSURLConnection ではなく、CFNetwork の CFHTTPMessage プリミティブに基づいて構築されています。これにより、URLSessionlevel のインストルメンテーション (スウィズリング、ネットワーク拡張、URLProtocol サブクラス) が無効になります。

[切り捨てられた]

## Original Extract

In this post we look under the hood of BrightData's SDK and how it turns ordinary consumer TVs into exit nodes of an enormous commercial, residential proxy network leveraged by the AI industry to scrape web data and train language learning models.

The Smart TV in Your LivingRoom Is a Node in the AIScraping Economy - Include Security Research Blog
Skip to content
Team Research blog
The Smart TV in Your LivingRoom Is a Node in the AIScraping Economy
The work at Include Security has us working with AI day in and day out (hacking it, using it, training it, etc).
We’re all aware of the community-level opposition happening against datacenters, aimed at improving AI capabilities, being built recently. What you might not be aware of are the distributed efforts to train AI that could be using the devices inside your home.
In this post, we’re going to explore how the company Bright Data facilitates modern AI models scraping training data from the Internet using its residential proxy network.
Bright Data is a data-collection company that sells access to what it markets as the world’s largest residential proxy network of 400M+ home IP addresses that its customers route web-scraping traffic through. The supply behind that network comes from an SDK: a piece of software embedded in consumer apps that, with the user’s consent, turns their phone or smart TV into one of those exit nodes.
We’ll document what you, the average user, should know about what this company’s SDK does on your systems such as your mobile phone and your smart TV. We’re going to explore how their SDK works, which platforms have shipped it, and why your Internet-connected TV is the ultimate proxy for AI models looking to train on data scraped from the Internet.
AI companies depend on web-scraped content: for pre-training, for retrieval, for agent grounding, for search. But the modern web isn’t scrapeable from a datacenter. Cloudflare, DataDome, HUMAN , among others throttle or block requests from known cloud IPs.
The workaround is residential proxies. A scraping job routed through a Comcast or T-Mobile subscriber’s connection arrives at the target site from an IP that belongs to a paying residential customer. Krebs reported in October 2025 that “a glut of proxies from Aisuru and other sources is fueling large-scale data harvesting efforts tied to various AI projects.” Academic measurement going back to 2019 shows these networks are overwhelmingly misused. The FBI issued a formal advisory earlier this year.
Most of the existing press has focused on the illegal residential-proxy supply: botnets ( Aisuru , Kimwolf ), trojanized apps ( HUMAN Security’s PROXYLIB disclosure ), pre-infected IoT hardware ( Google/Mandiant’s IPIDEA takedown ). These are the bad actors.
On the other hand, the legal supply side has received far less scrutiny. Today Bright Data is the largest residential proxy network in the world by its own marketing, advertising “150M+ IPs” sourced via a consent SDK embedded in partner apps. This research documents how that SDK works, which platforms have shipped it, and why the connected-TV is the ultimate residential proxy.
Why Connected TV (CTV) is the Ideal Proxy
Connected TV, a.k.a Smart TV, is a near-perfect residential proxy. Compared to a mobile phone:
A TV never hits 1% battery, jumps between WiFi networks or gets locked when the user is asleep. Some partner publishers do disclose the Bright Data relationship in their privacy policies PlayWorks is one example . But privacy-policy disclosure is the wrong control surface for a TV. It is hard to scroll through a legal document navigated by arrow keys on a remote, and the in-app consent dialog, doesn’t convey that a paying Bright Data customer is about to route their scraping traffic through the user’s home internet.
Petflix, a Roku app documented by The Verge , is a representative case. Its opt-in screen reads: “To enjoy Petflix for free with fewer ads, you are allowing Bright Data to occasionally use your device’s free resources and IP address to download public web data from the internet. Bright Data will only use your IP address for approved business-related use cases. None of your personal information is accessed or collected except your IP address. Period.” The Petflix dialog says “occasionally.” The SDK’s publicly queryable config sets max_bw_monthly_wifi: 200,000,000,000 bytes — a 200 GB default monthly WiFi budget.
Who Bright Data Names as Partners
Bright Data exposes a partner manifest endpoint. The endpoint is unauthenticated and anyone can fetch it. Names in the manifest that I was able to identify with high confidence from public sources:
Others ( desoline, free_time, ott_studio, global_microtrading, m_m_media, easystaff_lp ) are present but less identifiable from public sources. bright_screensavers, bright_videos , and brightdata are Bright Data’s own apps.
A note on what the partner list proves: Being listed in Bright Data’s config means an integration might have existed at some point. It does not by itself prove that a specific publisher’s currently-shipping app(s) includes the SDK in production. For any named publisher, per-app verification is required.
What the partner list does directly prove:
Bright Data ships this roster in an unauthenticated public endpoint .
At least three CTV-focused entities (PlayWorks, CloudTV, Longvision) monetized their user’s devices as residential proxy exit nodes. PlayWorks in particular reports CTV distribution across major TV platforms and ISPs, with reach figures in the hundreds of millions of households per its own marketing materials.
How does the Bright Data SDK turn a user’s device into a residential proxy exit node?
The Bright Data SDK is a publicly documented commercial product, offered to publishers via Bright Data’s SDK integration docs (with a JavaScript variant for web). What follows builds on that public surface with findings from reverse-engineering the shipping iOS framework and instrumenting 30 days of its runtime traffic.
The SDK ships as an iOS framework (brdsdk.framework) inside partner apps. I reverse-engineered the binary and captured 30 days of traffic from a research fleet running the SDK inside a consent-installed partner app.
On every launch the SDK calls:
GET <https://clientsdk.bright-sdk.com/sdk_config_ios.json>?appid=<bundle>&ver=<sdk-version>&uuid=sdk-ios-<32hex>
The endpoint is unauthenticated in any meaningful sense. The server gates only on two query parameters appid (an app bundle ID, which can be found in the App Store listing of the partner app) and ver (the SDK version string). Supply those and any randomly generated UUID, and the server returns the same response a real device gets: feature flags, idle-detection thresholds (battery %, CPU/memory ceilings, WiFi-vs-cellular rules), per-country bandwidth tiers, and the partner manifest I showcased above. Each of these branches is worth examining on its own: the idle rules that decide when your device is eligible to relay, a flag that routes peer traffic around your VPN, a map that stitches your installs across platforms into one identity, and the per-country bandwidth caps.
After config fetch, the SDK opens a persistent WebSocket to:
This hostname resolves to AWS Global Accelerator IPs (3.33.193.183, 15.197.193.114 as of this writing). The TLS certificate is CN=*.luminatinet.com — the domain for Luminati Networks, Bright Data’s pre-2018 corporate name. The rebrand was publicly announced in 2018 . Active SDK infrastructure still runs on the legacy cert, which is a useful detection pivot: the current customer-facing proxy service lives on brightdata.com-branded domains, so any luminatinet.com / brdtnet.com traffic on your network is specifically the peer-tunnel plane, not customer-side Bright Data usage. The server identifies itself as uWebSockets: 20 .
The peer endpoint requires no authentication to upgrade. The server accepts any TLS-valid WebSocket upgrade and immediately pushes the connecting client an application-layer frame with the client’s public IP echoed back. From there, a handshake unfolds:
Server → client: tunnel_init establishes the session, returns the client’s public IP.
Server → client: cid_set the server assigns the client a session-tracking identifier in the format <IP>-<token>/ls<N>c<M>p443_<IP>_<counter> . We confirmed this format matches the cid field present in the SDK’s captured telemetry traffic from real devices.
Server → client: status_get the server polls the device for its idle state, battery, network type, and available bandwidth. The device responds with a continuous telemetry feed: idle, wifi_connected, mobile_connected, mobile_type (LTE/5G), roaming, battery_level, using_battery, screen_on, on_call, cpu_usage, mem_usage, raw_bw, bw, ipv6_supported, appid (the host app), sdk_version, platform , and the assigned cid . This is a continuous feed of physical-device state to a third party, delivered via a consent dialog whose text is chosen by the host app publisher.
Handshake complete. Once the device reports favorable status, the server’s job-matching layer is free to push cmd_tun frames: individual scraping-job instructions that the SDK executes as HTTP requests against third-party sites, using the user’s residential IP as the source.
Every frame on the WebSocket is plain JSON with a fixed envelope:
{"type": "ipc_call"|"ipc_post"|"ipc_result"|"ipc_error",
"cmd": <command>, "cookie": <correlation-id>,
"err_code": 0, "msg": { ...payload... }}
The full command vocabulary extracted from the binary and verified on the wire:
There’s no message signing, HMAC, client certificate or device attestation. Only the TLS layer and the server’s IP-reputation filter gating which peers actually receive jobs. For readers familiar with commercial malware protocol design: this is substantially less secure than typical C2.
The config ships an explicit rulebook for when the device is eligible to relay someone else’s traffic:
"idle_metrics": {
"ignore_screen_on": true, // relay even with the screen on
"ignore_on_call": true, // relay while the user is on a phone call
"max_bw_ratio": 1,
"min_battery": 0.2,
"wifi_on_battery": true,
"min_battery_wifi": 0.2,
"max_cpu_usage": 70,
"max_mem_usage": 90,
"mem_screen_off": true,
"idle_timeout": 30,
"not_idle_timeout": 10
}
The ignore_screen_on and ignore_on_call flags are notable: “idle” does not mean the user is away from the device. It means the device’s CPU, memory, and battery are within the SDK’s thresholds. A user on a phone call, actively reading the screen, is considered idle for relay purposes.
The config also ships a dual_pairing map:
"dual_pairing": {
"ios_com.brd.earnapp": ["win_earnapp.com", "mac_com.earnapp"]
}
That’s a server-side map tying a user’s iOS, Windows, and macOS installations of the same brand into one entity. It’s cross-platform identity stitching documented inside a public config file.One more forward-looking field: http3_enabled: true. The SDK is already shipping the flag for QUIC-based peer transport. A future version may move the peer tunnel from TCP/443 to UDP/443, which would break any defender relying on TCP connection tracking to detect the WebSocket.
The SDK’s config ships a flag “use_netifs”: true. That flag triggers code in the SDK binary that constructs its NWConnection with a specific required interface: en0 (WiFi) or pdp_ip0 (cellular), rather than using the system default route.
On iOS, this bypasses any configured VPN’s tun0 interface entirely. The peer tunnel does not cross a user-configured VPN, even when the rest of the app’s HTTPS traffic does.
We observed this empirically. My research setup includes transparent TLS interception. It captured every HTTPS call the SDK made, except the peer tunnel to proxyjs.brdtnet.com:443, even though port 443 is explicitly redirected to the inspector. The bypass uses Apple’s documented NWParameters.requiredInterface API.
It’s worth emphasizing that the SDK uses two independent inspection bypasses , one per plane:
Control plane (config fetch, telemetry pings): built on CFNetwork’s CFHTTPMessage primitives rather than URLSession/NSURLConnection. This defeats URLSessionlevel instrumentation (swizzling, network extensions, URLProtocol subclasses) c

[truncated]
