---
source: "https://mullvad.net/en/blog/introducing-defense-against-ai-guided-traffic-analysis-daita"
hn_url: "https://news.ycombinator.com/item?id=48572780"
title: "Defense against AI-guided Traffic Analysis (DAITA) (2024)"
article_title: "Introducing Defense against AI-guided Traffic Analysis (DAITA) | Mullvad VPN"
author: "Cider9986"
captured_at: "2026-06-17T18:58:36Z"
capture_tool: "hn-digest"
hn_id: 48572780
score: 2
comments: 0
posted_at: "2026-06-17T16:28:38Z"
tags:
  - hacker-news
  - translated
---

# Defense against AI-guided Traffic Analysis (DAITA) (2024)

- HN: [48572780](https://news.ycombinator.com/item?id=48572780)
- Source: [mullvad.net](https://mullvad.net/en/blog/introducing-defense-against-ai-guided-traffic-analysis-daita)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T16:28:38Z

## Translation

タイトル: AI 誘導トラフィック分析に対する防御 (DAITA) (2024)
記事タイトル: AI 誘導トラフィック分析に対する防御の紹介 (DAITA) |マルバド VPN
説明: VPN (または Tor ネットワーク) でトラフィックを暗号化している場合でも、高度なトラフィック分析はプライバシーに対する脅威として増大しています。そこで今回はDAITAをご紹介します。

記事本文:
AI 誘導トラフィック分析に対する防御 (DAITA) の紹介 |マルバド VPN
メイン コンテンツにスキップ Mullvad ロゴ メニュー 製品とサービス
ログイン はじめる AI 誘導トラフィック分析に対する防御の導入 (DAITA)
2024 年 5 月 7 日 ニュース 特集 プライバシー
VPN (または Tor ネットワーク) でトラフィックを暗号化している場合でも、高度なトラフィック分析はプライバシーに対する脅威として増大しています。そこで今回はDAITAをご紹介します。
一定のパケット サイズ、ランダムなバックグラウンド トラフィック、データ パターンの歪みを通じて、私たちは高度なトラフィック分析との戦いの第一歩を踏み出しています。
VPN (または Tor ネットワーク) 経由でインターネットに接続すると、IP アドレスがマスクされ、トラフィックが暗号化され、インターネット サービス プロバイダーから隠蔽されます。プライバシーを重視した Web ブラウザも使用している場合は、サードパーティの Cookie、ピクセル、ブラウザのフィンガープリントなどの他の追跡テクノロジを通じて、敵対者があなたのアクティビティを監視することが困難になります。
しかし、依然として、今日の大規模監視はかつてないほど洗練されており、高度なトラフィック分析による暗号化通信のパターン分析がプライバシーに対する脅威として増大しています。
これは、暗号化されている場合でも、AI を使用してトラフィックを分析する方法です。
Web サイトにアクセスすると、パケットの交換が行われます。デバイスは訪問先のサイトにネットワーク パケットを送信し、サイトはパケットを返送します。これはインターネットのまさにバックボーンの一部です。 VPN (または Tor ネットワーク) を使用している場合でも、パケットが送信されているという事実、パケットのサイズ、送信頻度は ISP に表示されます。
すべての Web サイトは、その要素の構成に基づいて送受信されるネットワーク パケットのパターンを生成するため (

画像やテキスト ブロックなど）、AI を使用してトラフィック パターンを特定の Web サイトに結び付けることが可能です。これは、ISP または ISP にアクセスできるオブザーバー (当局またはデータ ブローカー) が、デバイスに出入りするすべてのデータ パケットを監視し、この種の分析を行って、アクセスしたサイトだけでなく、相関攻撃 (特定の時間に特定のパターンのメッセージを送信し、別のデバイスに特定のパターンのメッセージを同時に受信する) を使用して通信する相手も追跡できることを意味します。
トラフィック分析とどう戦うか: これが DAITA の仕組みです。
DAITA はカールスタード大学のコンピューター サイエンスと共同で開発され、トラフィック分析に対抗するために 3 種類のカバー トラフィックを使用します。
ネットワーク パケットのサイズ、特に小さなパケットのサイズは明らかになる可能性があるため、DAITA では、VPN 経由で送信されるすべてのパケットを同じ一定のサイズにします。
DAITA は、ダミー パケットをトラフィックに予期せず散在させることで、デバイスとの間で送受信される日常的な信号をマスクします。これにより、観察者が意味のあるアクティビティと背景ノイズを区別することが難しくなります。
Web サイトにアクセスするとき (または、大量のトラフィックを引き起こすその他のアクティビティを行うとき)、DAITA は、クライアントと VPN サーバーの間の両方向にカバー トラフィックを予期せず送信することで、トラフィック パターンを変更します。これにより、Web サイト訪問の認識可能なパターンが歪められ、サイトを正確に識別できなくなります。
トラフィックデータを販売するデータブローカーの未来はすでに到来しています
今日の高度な AI により、トラフィック分析は大規模な監視に使用できる可能性があります。現在、トラフィック分析がどの程度使用されているかを確認するのは困難です。しかし、野心はそこにあります。 2021年にViceは、FBIが90％以上をカバーすると主張するデータブローカーからネットフローデータを購入したと報告した

世界のインターネット トラフィックの中で最大のトラフィックです。
トラフィック分析が将来どのように使用できるかは、概要を把握するのが困難です。だからこそ、私たちは今日、抵抗に取り組む必要があるのです。 DAITA のこの初期バージョンは、オンライン プライバシーの進化する課題に対する最初の対応です。 DAITA はオープンソースとしてリリースされており、フィードバックを収集しながら改良と開発を続け、プライバシー テクノロジーの最前線であり続けることを保証します。
「現在、トラフィック分析がどの程度使用されているかについて推測する必要はありません。私たちは AI の発展と権威主義社会の発展を観察しているだけです。また、将来の大規模監視においてトラフィック分析がどのような役割を果たすかについて推測する必要もありません。私たちがしなければならないのは、脅威と機会を認識し、抵抗に取り組むことです。」と Mullvad VPN の CEO、ヤン ジョンソン氏は述べています。
DAITA の構成要素はオープンソースです
DAITA は、Mullvad が開発資金を援助しているオープンソースのMaybenot防衛フレームワークを使用して構築されています。この研究は学術的に査読され、オープンアクセスとして出版されています。
「トラフィック分析防御の実践は、長い間待ち望まれてきました。AI の急速な発展により、この分野は変化しているため、フレームワークに時間とエネルギーを投資することは完全に理にかなっています。」と、カールスタード大学の研究者、トビアス プルス氏は述べています。
まず、DAITA 2024.3-beta1 は、Windows 10 および 11 の VPN アプリで利用できます。
DAITA の使用を開始するには: Windows 用 Mullvad VPN のベータ版をダウンロードします。 「設定」→「VPN設定」→「WireGuard設定」→「DAITA」をオンにします。
バグまたは脆弱性の報告

## Original Extract

Even if you have encrypted your traffic with a VPN (or the Tor Network), advanced traffic analysis is a growing threat against your privacy. Therefore, we now introduce DAITA.

Introducing Defense against AI-guided Traffic Analysis (DAITA) | Mullvad VPN
Skip to main content Mullvad Logo Menu Products and services
Log in Get started Introducing Defense against AI-guided Traffic Analysis (DAITA)
May 7, 2024 News Features Privacy
Even if you have encrypted your traffic with a VPN (or the Tor Network), advanced traffic analysis is a growing threat against your privacy. Therefore, we now introduce DAITA.
Through constant packet sizes, random background traffic and data pattern distortion we are taking the first step in our battle against sophisticated traffic analysis.
When you connect to the internet through a VPN (or the Tor Network) your IP address is masked, and your traffic is encrypted and hidden from your internet service provider. If you also use a privacy-focused web browser , you make it harder for adversaries to monitor your activity through other tracking technologies such as third-party cookies, pixels or browser fingerprints.
But still, the mass surveillance of today is more sophisticated than ever, and a growing threat against privacy is the analysis of patterns in encrypted communication through advanced traffic analysis .
This is how AI can be used to analyze your traffic – even if it’s encrypted.
When you visit a website, there is an exchange of packets: your device will send network packets to the site you're visiting and the site will send packets back to you. This is a part of the very backbone of the internet. The fact that packets are being sent, the size of the packets, and how often they are sent will still be visible for your ISP, even if you are using a VPN (or the Tor network).
Since every website generates a pattern of network packets being sent back and forth based on the composition of its elements (like images and text blocks), it’s possible to use AI to connect traffic patterns to specific websites. This means your ISP or any observer (authority or data broker) having access to your ISP can monitor all the data packets going in and out of your device and make this kind of analysis to attempt to track the sites you visit, but also who you communicate with using correlation attacks (you sending messages with certain patterns at certain times, to another device receiving messages with a certain pattern at same times).
How we combat traffic analysis: this is how DAITA works.
DAITA has been developed together with Computer Science at Karlstad University and uses three types of cover traffic to resist traffic analysis.
The size of network packets can be particularly revealing, especially small packets, so DAITA makes all packets sent over the VPN the same constant size.
By unpredictably interspersing dummy packets into the traffic, DAITA masks the routine signals to and from your device. This makes it harder for observers to distinguish between meaningful activity and background noise.
When visiting websites (or doing any other activity that causes significant traffic), DAITA modifies the traffic pattern by unpredictably sending cover traffic in both directions between client and VPN server. This distorts the recognizable pattern of a website visit, resisting accurate identification of the site.
The future of data brokers selling traffic data is already here
With the sophisticated AI of today, traffic analysis can potentially be used for mass surveillance. The extent to which traffic analysis is used today is difficult to ascertain. But the ambition is there. In 2021, Vice reported that the FBI purchased netflow data from a data broker claiming to cover over 90 percent of the world’s internet traffic .
How traffic analysis can be used in the future is hard to overview. That’s why we need to work on a resistance today. This initial version of DAITA is our first response to the evolving challenges of online privacy. DAITA is released as open source and as we gather feedback we will continue to refine and develop, ensuring it remains at the forefront of privacy technology.
“ We don't need to speculate on the extent to which traffic analysis is being used today. We just observe the development of AI and the development of authoritarian societies. There is also no need to speculate on which role traffic analysis will play in future mass surveillance. What we must do is to recognize the threats and opportunities – and work on resistance”, says Jan Jonsson, CEO at Mullvad VPN.
The building blocks of DAITA are open source
DAITA is built using the open-source Maybenot defense framework, which Mullvad helps to fund development of. The work has been academically peer reviewed and published as open access .
“ Putting traffic analysis defenses to practice is long overdue. Because the area is changing due to the rapid development of AI, investing time and energy into a framework makes perfect sense”, says Tobias Pulls, researcher at Karlstad University.
To begin with, DAITA 2024.3-beta1 is available in our VPN app on Windows 10 and 11.
To start using DAITA: Download the beta version of Mullvad VPN for Windows. Go to Settings – VPN settings – WireGuard settings – turn on DAITA.
Reporting a bug or vulnerability
