---
source: "https://blog.bytebytego.com/p/how-openai-delivers-low-latency-voice"
hn_url: "https://news.ycombinator.com/item?id=48825411"
title: "How OpenAI Delivers Low-Latency Voice AI for 900M Users"
article_title: "How OpenAI Delivers Low-Latency Voice AI for 900M Users"
author: "gmays"
captured_at: "2026-07-07T23:47:32Z"
capture_tool: "hn-digest"
hn_id: 48825411
score: 1
comments: 0
posted_at: "2026-07-07T23:26:38Z"
tags:
  - hacker-news
  - translated
---

# How OpenAI Delivers Low-Latency Voice AI for 900M Users

- HN: [48825411](https://news.ycombinator.com/item?id=48825411)
- Source: [blog.bytebytego.com](https://blog.bytebytego.com/p/how-openai-delivers-low-latency-voice)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T23:26:38Z

## Translation

タイトル: OpenAI が 9 億ユーザーに低遅延の音声 AI を提供する方法
説明: この記事では、その過程全体と、OpenAI エンジニアリング チームが直面した課題について詳しく説明します。

記事本文:
OpenAI が 9 億ユーザーに低遅延の音声 AI を提供する方法
ByteByteGo ニュースレター
OpenAI が 9 億ユーザーに低遅延の音声 AI を提供する方法
ByteByteGo 2026 年 7 月 1 日 224 4 6 シェア 家庭でのインテリジェントな自動化を可能にするロボット掃除機 (スポンサー付き)
Matic は、あなたと同じように家を観察し、思いどおりに掃除できる初の視覚的にインテリジェントなロボット掃除機です。
マティッチのヒーローの特徴:
完全にカメラ上で動作し、障害物を巧みに乗り越えます
床の種類を認識して掃除機掛けとモップ掛けを自動で切り替える
大きなホイールと高さ調節可能なクリーニングヘッドで厚いラグも処理可能
55dBで会話より静か
ペットの毛を詰まらせたり絡ませたりせずに処理します。
1 つのバッグに乾いた排泄物と湿った排泄物を収集します。おむつ用塩ジェルの汚れた水と抗菌パウダーがカビを防ぎます。
すべてのバッグに新鮮な HEPA フィルターが入っているので、洗浄や交換の必要がなく、空気がきれいになります。
180 日間の返金保証付きの Matic で自律清掃を体験してください。
OpenAI は毎週 9 億人のユーザーに対して音声 AI を実行しており、その代替手段としてインターネットがライブ オーディオを処理する方法を再発明する必要があるため、WebRTC を使用しています。
問題は、WebRTC が安定した IP とポートを持つサーバー向けに設計されており、Kubernetes はそれらのアドレスを使い捨てとして扱うことです。この規模における従来の答えは、グループ ビデオ通話などのマルチパーティのワークロードに適した SFU ですが、OpenAI のトラフィックは圧倒的に 1 人のユーザーが 1 つのモデルと話しているものです。
これに対処するために、彼らのアーキテクチャはスタックを 2 つの部分に分割します。
ステートレス リレーは、地理的なエッジでプロトコルを認識したパケット ルーティングを処理します。
ステートフル トランシーバーは、すべての重い WebRTC 状態を所有します。
それらを結びつける秘訣は、プロトコルがセットアップ中にすでに交換しているフィールドである ICE ufrag を、リレーが最初から読み取ることができるルーティング キーとして使用することです。

新しいセッションの t パケット。 Global Relay からユーザー空間の Go 実装、Redis キャッシュ、および慎重なソケットレベルの最適化に至るまで、その他すべては、その中心となるアイデアに基づいて構築されています。
この記事では、その過程全体と、OpenAI エンジニアリング チームが直面した課題について詳しく見ていきます。
免責事項: この投稿は、Open AI エンジニアリング チームから一般に共有された詳細に基づいています。間違いに気づいた場合はコメントしてください。
音声 AI にとって遅延が重要な理由
音声 AI は会話のように感じられるか、トランシーバーのように感じられます。これらのエクスペリエンスの間の境界線はミリ秒単位で測定されます。
ユーザーの声を聞いてから応答するまでの間にネットワークが一時停止すると、錯覚が壊れます。間がぎこちなくなり、中断が切り取られ、ユーザーは AI を文の途中で中断せざるを得なくなりますが、これも一種の失礼です。言い換えれば、音声 AI が自然に感じられるのは、会話が話し言葉のスピードで進んでいる場合に限られます。
その下にあるより厳しい制約は、連続ストリーム プロパティです。音声は、ユーザーが話し終えた後の単一のアップロードとしてではなく、安定した流れとしてモデルに到達する必要があります。このストリームにより、ユーザーが話している間にモデルが文字起こし、推論、ツールの呼び出しを開始できるようになります。エクスペリエンスが中断されると、プッシュツートークに崩壊します。
特に OpenAI の場合、これらの制約は次の 3 つの具体的な要件に変換されます。
このシステムは、毎週 9 億人のアクティブ ユーザーがどこにいてもアクセスできる必要があります。
接続セットアップは、ユーザーがセッションの開始後すぐに話し始めることができるように、十分早く完了する必要があります。
ターンテイクがサクサクと感じられるように、オーディオの往復時間は短く安定している必要があります。
WebRTC は、この種の作業のために業界が構築したプロトコルです。これは、より小さなプロトコルのバンドルです (2 つのエンドポイントがどのように相互に到達するかを理解するための ICE)

ファイアウォール、チャネル暗号化のための DTLS、音声パケットのための SRTP、および品質フィードバックのための RTCP) を越えます。 WebRTC のオリジナル アーキテクトの 1 人である Justin Uberti と、OpenAI の基盤となる Pion ライブラリを管理している Sean DuBois は、現在 OpenAI で働いています。この種のプロトコルの深さは、彼らが出荷したアーキテクチャに現れています。
OpenAI の WebRTC インフラストラクチャの最初のバージョンは、Pion 上に構築された単一の Go サービスでした。両方のジョブを 1 か所で処理しました。
シグナリング側では、サービスは SDP (クライアントとサーバーがセッションを記述するために使用する形式) をネゴシエートし、コーデックを選択し、ICE 資格情報を生成し、セッションをセットアップしました。
メディア側では、このサービスはクライアントからの WebRTC 接続を終了し、推論、文字起こし、音声生成、ツールの使用、オーケストレーションなどの AI モデルを実行するバックエンド サービスへのアップストリーム接続を維持しました。
この複合サービスは今でも ChatGPT 音声、Realtime API の WebRTC エンドポイント、およびいくつかの研究プロジェクトを強化しており、それらの作業をうまく処理しています。 OpenAI が直面した問題は、最新のクラウド インフラストラクチャを実行するコンテナ オーケストレーション システムである Kubernetes に OpenAI をデプロイする方法でした。
Kubernetes は、コンピューティングが安価で移動可能であることを前提としています。ポッドが起動し、容量が存在する限りスケジュールされ、しばらく実行された後、再スケジュールまたは交換されます。標準的な WebRTC デプロイメント パターンはその逆を想定しています。この不一致は、特定の 2 つの場所に現れます。
WebRTC を展開する従来の方法では、セッションごとに 1 つの UDP ポートを使用します。 OpenAI の規模では、サービスごとに数万のパブリック UDP ポートがあることを意味します。クラウド ロード バランサーは、少数の既知のポート用に構築されているため、範囲が追加されるたびに、ロード バランサーの構成、ヘルス チェック、ファイアウォール ポリシー、ロールアウトの安全性の運用が複雑になります。露出した表面

また、セキュリティ監査も難しくなります。 Kubernetes の自動スケーリングは、大きく安定したポート範囲を予約するという要件と衝突するため、弾力性が脆弱になります。
2 つ目は状態のスティッキーです。
サーバーごとに 1 つの UDP ポートを実行し、その背後でセッションを多重化解除することで、ポートの問題が解決されます。ただし、ICE と DTLS はステートフル プロトコルです。セッションを開始したプロセスは、接続チェックを検証し、DTLS ハンドシェイクを完了し、SRTP を復号化し、ICE の再起動などの後のセッション変更を処理するために、パケットを受信し続ける必要があります。既存のセッションのパケットが別のプロセスに到達すると、セットアップが失敗するか、メディアが破損します。
どちらの圧力も同じ答えを示しています。クライアントのエクスペリエンスは同一のままで、導入アーキテクチャを変更する必要があります。
リレーをトランシーバーから分離する
OpenAI が出荷したアーキテクチャは、パケット ルーティングをプロトコル終端から分割します。
無国籍の中継器が最前線に位置し、インターネットへの小さな公共の足跡を示しています。ステートフル トランシーバーはその背後にあり、すべての重い WebRTC 状態を所有します。シグナリングは引き続きトランシーバーに直接送信されます。メディアはまず中継経由で入ります。
リレーの範囲は意図的に狭いです。各パケットを十分に読み取って宛先を選択し、残りを不透明なペイロードとして転送します。オーディオは途中で暗号化されたままとなり、ICE ステート マシンはトランシーバーに留まり、コーデックのネゴシエーションは他の場所で行われます。クライアントの観点から見ると、WebRTC セッションはあらゆる点で正常に見えます。
トランシーバーは、物事を記憶する必要がある WebRTC の部分を所有します。 ICE 接続チェック、DTLS ハンドシェイク、SRTP 暗号化キー、セッション ライフサイクルはすべてそこに存在します。トランシーバーは、ハンドシェイクを完了し、実際のメディアを暗号化または復号化するエンドポイントです。
明らかな代替案があった

チームが評価して選択したことを意味します。
SFU (Selective Forwarding Unit) は、大規模な WebRTC の標準メディア サーバー アーキテクチャです。参加者ごとに 1 つの WebRTC 接続を終了し、参加者間でストリームを選択的に転送します。 AI は別の参加者として参加します。
これは、グループ通話、教室、共同会議など、本質的にマルチパーティ製品に適しています。 OpenAI のワークロードは異なるようです。ほとんどのセッションは 1 対 1 で、1 人のユーザーが 1 人のモデルと会話します。この種のトラフィックの場合、SFU モデルはオーバーヘッドを追加し、バックエンド サービスが WebRTC ピア自体のように動作するように強制します。トランシーバー モデルでは、バックエンドを通常のサービスのままにします。
SFU アプローチを示す以下の図を参照してください。
TURNも検討して脇に置きました。
TURN は、NAT トラバーサルに使用される標準のプロトコル終端リレーです。問題は、TURN 割り当てでは、メディアが流れる前にセットアップのラウンドトリップが追加され、サーバー間での移行や回復が困難であることです。レイテンシの影響を受けやすいワークロードの場合、余分なラウンドトリップが重要になります。
分割により、原則としてポートと状態の問題が解決されます。残りの問題は、リレーが最初のパケットを正しくルーティングすることです。
新しいセッションの最初のパケットは難しいものです。
リレーには、この送信元 IP とポートからのパケットがこのトランシーバーに送信されるというマッピングがあるため、後続のパケットは簡単です。最初のパケットはそのマッピングを作成するものであるため、リレーはパケット自体からマッピングをどこに送信するかを判断する必要があります。
2 つの単純なオプションが存在しました。
ホット パスでのデータベース ルックアップにより、待機時間が追加され、正常な状態を維持する別のサービスへの強い依存関係が生じます。
ランダムなトランシーバーへのルーティングと転送は内部的に機能しますが、ホップ数が 2 倍になります。
答えは、WebRTC がすでに交換しているフィールドの中にあります。すべての WebRTC セッション

sion は、ufrag と呼ばれる ICE ユーザー名のフラグメントを伝送します。これはセッションのセットアップ中に生成され、STUN バインディング リクエストでエコーされます。 STUN バインディング リクエストは、2 つのエンドポイントが実際に相互に到達できることを確認するために ICE が使用するパケットであり、通常はクライアントがメディア パス上で最初に送信するものです。
重要なのは、OpenAI がシグナリング中にサーバー側の ufrag を生成することです。必要なものを何でも入れることができるため、ルーティング メタデータをエンコードします。リレーは、最初の STUN バインディング リクエストを十分に解析して、ufrag を読み取り、ルーティング ヒントをデコードし、セッションを所有するトランシーバーにパケットを転送します。最初のパケット以降のすべてのパケットは、確立されたセッション マッピングを通過し、ufrag 解析ステップを完全にスキップします。
接続確立シーケンスの詳細を示す以下の図を参照してください。
フリート内の各トランシーバーは、内部 IP とポートにバインドされた 1 つのオペレーティング システム エンドポイントである共有 UDP ソケットで待機します。そのトランシーバーのすべてのセッションはその背後で多重化されます。
シグナリング中に、トランシーバーは SDP 応答で共有リレー VIP と UDP ポートを返します。 VIP はリレー フリート全体の先頭にある仮想 IP アドレスであるため、そのアドレスの背後に多くのリレー インスタンスがある場合でも、クライアントには 203.0.113.10:3478 のような 1 つの安定した宛先が見えます。クライアント側から見ると、パケットを送信する場所は 1 つあり、それはセッションが存続する間同じままです。
リレーの状態は意図的に小さくなっています。これは、送信元アドレスからトランシーバー宛先までのメモリ内マップに加えて、監視用のいくつかのカウンターとセッション クリーンアップ用のタイマーを保持します。リレー インスタンスが再起動してマッピングが失われた場合、次の STUN パケットによって ufrag からマッピングが再構築されます。リカバリを高速化するために、Redis キャッシュはソースから宛先へのマッピングを一度保持します。

ユートが設立されました。再起動されたリレーは、Redis からマッピングをすぐに検索できます。
ここでの原則はよく一般化されます。ホット パス上のデータが必要な場合は、プロトコルがすでに何を交換しているかを調べます。ペイロードのフィールドは基本的に自由に解析できます。新しい検索には遅延、依存関係が発生し、さらにもう 1 つ問題が発生する可能性があります。
グローバルリレーとジオステアリングシグナリング
パブリック UDP サーフェスが少数の固定アドレス セットに縮小されると、同じリレー パターンがグローバルに展開可能になりました。
Global Relay は、OpenAI の地理的に分散されたリレー入口ポイントのフリートです。これらはすべて、上記で説明した同一のパケット転送動作を実行します。唯一変わるのは、地図上のどこに位置するかです。
地理的な分散により、クライアントから OpenAI への最初のホップが短縮されます。ユーザーに近いリレーからネットワークに入るパケットは、地理的にもネットワーク トポロジー的にも、最初に遠くの地域に到達するためにパブリック インターネットを通過する必要があるパケットよりもはるかに簡単です。実際の効果は、トラフィックが OpenAI バックボーンに到達する前に、レイテンシが低くなり、タイミングがより安定し、損失プロファイルがきれいになることです。
OpenAI は、シグナリング側での地理的および近接ステアリングに Cloudflare を使用します。初期の HTTP または WebSocket r

[切り捨てられた]

## Original Extract

In this article, we will look at the entire journey in detail and challenges the OpenAI engineering team faced.

How OpenAI Delivers Low-Latency Voice AI for 900M Users
ByteByteGo Newsletter
Subscribe Sign in How OpenAI Delivers Low-Latency Voice AI for 900M Users
ByteByteGo Jul 01, 2026 224 4 6 Share The Robot Vacuum Making Intelligent Automation Possible at Home (Sponsored)
Matic is the first visually intelligent robot vacuum that sees your home like you do, so it can clean how you want it to.
Matic’s hero features:
Runs entirely on cameras to deftly navigate obstacles
Recognizes floor types to auto-switch between vacuuming and mopping
Big wheels and a height-adjustable cleaning head handle thick rugs
Quieter than conversations at 55 dB
Handles pet hair without clogging or tangling
A single bag collects dry and wet waste—diaper-salts gel dirty water, antimicrobial powder prevents mold
A fresh HEPA filter in every bag for cleaner air, no washing or replacing
Experience autonomous cleaning with Matic with a 180-day money-back guarantee.
OpenAI runs voice AI for 900 million users a week, and they use WebRTC for it because the alternative would mean reinventing how the internet handles live audio.
The catch is that WebRTC was designed for servers with stable IPs and ports, and Kubernetes treats those addresses as disposable. The conventional answer at this scale is an SFU, which suits multiparty workloads like group video calls, but OpenAI’s traffic is overwhelmingly one user talking to one model.
To deal with this, their architecture splits the stack into two pieces:
A stateless relay handles protocol-aware packet routing at the geographic edge.
A stateful transceiver owns all the heavy WebRTC state.
The trick that ties them together is using the ICE ufrag, a field the protocol already exchanges during setup, as a routing key that the relay can read off the first packet of a new session. Everything else, from Global Relay to the userspace Go implementation to the Redis cache and the careful socket-level optimizations, builds on top of that core idea.
In this article, we will look at the entire journey in detail and challenges the OpenAI engineering team faced.
Disclaimer: This post is based on publicly shared details from the Open AI Engineering Team. Please comment if you notice any inaccuracies.
Why Latency Matters For Voice AI
Voice AI either feels like a conversation or it feels like a walkie-talkie. The line between those experiences is measured in milliseconds.
When the network pauses between hearing a user and responding, the illusion breaks. Pauses turn awkward, interruptions get clipped, and users are compelled to cut off the AI mid-sentence, which is also kind of rude. In other words, voice AI only feels natural if the conversation moves at the speed of speech.
The harder constraint underneath is the continuous-stream property. Audio has to arrive at the model as a steady flow, rather than as a single upload after the user finishes talking. That stream is what lets the model start transcribing, reasoning, and calling tools while the user is still speaking. The experience collapses into push-to-talk once it breaks.
For OpenAI specifically, those constraints translate into three concrete requirements:
The system has to reach 900 million weekly active users wherever they are.
Connection setup has to be completed quickly enough that users can start speaking as soon as a session begins.
Round-trip time for audio has to stay low and stable so turn-taking feels crisp.
WebRTC is the protocol the industry built for this kind of work. It is a bundle of smaller protocols (ICE for figuring out how two endpoints reach each other across firewalls, DTLS for encrypting the channel, SRTP for the audio packets, and RTCP for quality feedback). Justin Uberti, one of WebRTC’s original architects, and Sean DuBois, who maintains the Pion library OpenAI builds on, both work at OpenAI today. That kind of protocol depth shows up in the architecture they shipped.
The first version of OpenAI’s WebRTC infrastructure was a single Go service built on Pion. It handled both jobs in one place:
On the signaling side, the service negotiated SDP (the format clients and servers use to describe a session), selected codecs, generated ICE credentials, and set up sessions.
On the media side, the service terminated WebRTC connections from clients and maintained upstream connections to the backend services that run the AI models, including inference, transcription, speech generation, tool use, and orchestration.
That combined service still powers ChatGPT voice, the Realtime API’s WebRTC endpoint, and several research projects, and it has handled that work well. The question OpenAI ran into was how to deploy it on Kubernetes, the container orchestration system that runs most modern cloud infrastructure.
Kubernetes assumes compute is cheap and movable. Pods come up, get scheduled wherever capacity exists, run for a while, then get rescheduled or replaced. Standard WebRTC deployment patterns assume the opposite. That mismatch shows up in two specific places.
The conventional way to deploy WebRTC uses one UDP port per session. At OpenAI’s scale, that means tens of thousands of public UDP ports per service. Cloud load balancers were built for a handful of well-known ports, so each additional range adds operational complexity for load balancer config, health checks, firewall policy, and rollout safety. The exposed surface area also makes security audits harder. Kubernetes autoscaling clashes with the requirement to reserve large and stable port ranges, which makes elasticity brittle.
The second is state stickiness.
Running one UDP port per server and demultiplexing sessions behind it solves the port problem. ICE and DTLS, however, are stateful protocols. The process that started a session has to keep receiving its packets to validate connectivity checks, complete the DTLS handshake, decrypt SRTP, and process later session changes like ICE restarts. If a packet for an existing session lands on a different process, setup fails, or media breaks.
Both pressures point to the same answer. The deployment architecture has to change while the client experience stays identical.
Splitting The Relay From The Transceiver
The architecture OpenAI shipped splits packet routing from protocol termination.
A stateless relay sits at the front, presenting a small public footprint to the internet. A stateful transceiver sits behind it, owning all the heavy WebRTC state. Signaling still goes directly to the transceiver. Media enters through the relay first.
The relay’s scope is deliberately narrow. It reads enough of each packet to choose a destination, then forwards the rest as an opaque payload. Audio stays encrypted on the way through, ICE state machines stay with the transceiver, and codec negotiation happens elsewhere. From a client’s perspective, the WebRTC session looks normal in every way.
The transceiver owns the parts of WebRTC that have to remember things. ICE connectivity checks, the DTLS handshake, SRTP encryption keys, and the session lifecycle all live there. The transceiver is the endpoint that completes the handshakes and encrypts or decrypts the actual media.
There was an obvious alternative that the team evaluated and chose against.
An SFU, or Selective Forwarding Unit, is the standard media server architecture for WebRTC at scale. It terminates one WebRTC connection per participant and selectively forwards streams between them. The AI joins as another participant.
This works well for inherently multiparty products like group calls, classrooms, and collaborative meetings. OpenAI’s workload looks different. Most sessions are 1:1, with one user talking to one model. For that kind of traffic, the SFU model adds overhead and forces backend services to behave like WebRTC peers themselves. The transceiver model lets the backend stay an ordinary service.
See the diagram below that shows the SFU approach:
TURN was also considered and set aside.
TURN is the standard protocol-terminating relay used for NAT traversal. The trouble is that TURN allocations add setup round-trips before media can flow, and migrating or recovering them across servers is hard. For a latency-sensitive workload, those extra round-trip matters.
The split solves the port and state problems in principle. The remaining problem is making the relay route the first packet correctly.
The first packet of any new session is the difficult one.
Subsequent packets are easy because the relay has a mapping that says that packets from this source IP and port go to this transceiver. The first packet is what creates that mapping, so the relay has to figure out where to send it from the packet itself.
Two naive options were present:
A database lookup on the hot path adds latency and a hard dependency on another service staying healthy.
Routing to a random transceiver and forwarding internally works, but doubles the hop count.
The answer lives inside a field that WebRTC already exchanges. Every WebRTC session carries an ICE username fragment, called the ufrag, which is produced during session setup and echoed in STUN binding requests. STUN binding requests are the packets ICE uses to verify that two endpoints can actually reach each other, and they are usually the first thing a client sends on the media path.
The trick is that OpenAI generates the server-side ufrag during signaling. They can put whatever they want in it, so they encode routing metadata into it. The relay parses just enough of the first STUN binding request to read the ufrag, decode the routing hint, and forward the packet to the transceiver that owns the session. Every packet after the first one flows through the established session mapping, which skips the ufrag parsing step entirely.
See the diagram below that shows the connection establishment sequence in detail:
Each transceiver in the fleet listens on a shared UDP socket, which is one operating system endpoint bound to an internal IP and port. All sessions for that transceiver multiplex behind it.
During signaling, the transceiver returns a shared relay VIP and UDP port in the SDP answer. A VIP is a virtual IP address that fronts the entire relay fleet, so the client sees one stable destination like 203.0.113.10:3478, even though many relay instances sit behind that address. From the client’s side, there is one place to send packets, and it stays the same for the life of the session.
The relay’s state is purposefully tiny. It holds an in-memory map of source address to transceiver destination, plus some counters for monitoring and timers for session cleanup. If a relay instance restarts and loses the mapping, the next STUN packet rebuilds it from the ufrag. To make recovery faster, a Redis cache holds the source-to-destination mapping once a route is established. A restarted relay can look up the mapping from Redis immediately.
The principle here generalizes well. When we need data on the hot path, look at what the protocol is already exchanging. A field on the payload is essentially free to parse. A new lookup costs latency, a dependency, and one more thing that can break.
Global Relay and Geo-Steered Signaling
Once the public UDP surface was reduced to a small fixed set of addresses, the same relay pattern became deployable globally.
Global Relay is OpenAI’s fleet of geographically distributed relay ingress points. All of them run the identical packet-forwarding behavior described above. The only thing that changes is where on the map they sit.
Geographic distribution shortens the first client-to-OpenAI hop. A packet entering the network at a relay close to the user, in both geography and network topology, has a much easier time than a packet that has to traverse the public internet to reach a distant region first. The practical effect is lower latency, more stable timing, and a cleaner loss profile before traffic reaches the OpenAI backbone.
OpenAI uses Cloudflare for geographic and proximity steering on the signaling side. The initial HTTP or WebSocket r

[truncated]
