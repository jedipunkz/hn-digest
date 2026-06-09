---
source: "https://livekit.com/blog/why-webrtc-beats-websockets-for-voice-ai-agents"
hn_url: "https://news.ycombinator.com/item?id=48462241"
title: "Why WebRTC beats WebSockets for realtime voice AI"
article_title: "Why WebRTC beats WebSockets for realtime voice AI | LiveKit"
author: "jrm-veris"
captured_at: "2026-06-09T16:04:29Z"
capture_tool: "hn-digest"
hn_id: 48462241
score: 2
comments: 1
posted_at: "2026-06-09T15:17:30Z"
tags:
  - hacker-news
  - translated
---

# Why WebRTC beats WebSockets for realtime voice AI

- HN: [48462241](https://news.ycombinator.com/item?id=48462241)
- Source: [livekit.com](https://livekit.com/blog/why-webrtc-beats-websockets-for-voice-ai-agents)
- Score: 2
- Comments: 1
- Posted: 2026-06-09T15:17:30Z

## Translation

タイトル: リアルタイム音声 AI において WebRTC が WebSockets に勝つ理由
記事のタイトル: なぜ WebRTC がリアルタイム音声 AI で WebSockets に勝るのか |ライブキット
説明: WebSocket は高速にバイトを移動できますが、音声 AI エージェントにはそれ以上のものが必要です。 WebRTC とその背後にある SFU アーキテクチャがリアルタイム音声プロダクションに適したトランスポート層である理由は次のとおりです。

記事本文:
リアルタイム音声 AI において WebRTC が WebSockets に勝る理由 | LiveKit メイン コンテンツにスキップ モバイルでバナーの高さを予約するためのプレースホルダー テキスト 製品 シェブロン アイコン
ライブキット / エージェント 10.9 K エージェント 10.9 K
ブログ / エンジニアリング リアルタイム音声 AI において WebRTC が WebSockets に勝つ理由
開発者が音声 AI エージェントの構築を開始するとき、アーキテクチャ上の最初の決定はトランスポートです。つまり、ユーザーとエージェントの間で音声をどのように伝達するかです。 WebSocket はなじみがあり、十分に文書化されており、すでにほとんどの Web スタックの一部であるため、多くの人が WebSocket に手を伸ばします。これは合理的な選択のように思えます。ソケットを開き、オーディオ バイトを両方向にストリーミングして完了です。
デモでは動作します。制作中に崩れてしまいます。
「音声が流れている」ことと「これは本物の会話のように感じられる」との間のギャップは非常に大きく、それはほぼ完全に通信の問題です。 WebSocket はリアルタイム メディア用に設計されたものではありません。 WebRTC でした。この違いは、ほとんどの開発者が構築を開始するときに期待するよりもはるかに重要です。
WebSocket が実際に提供するもの
WebSocket は、クライアントとサーバーの間に永続的な全二重 TCP 接続を提供します。チャット、通知、構造化データのストリーミングに最適です。このようなユースケースには、これらが適切なツールです。
ただし、生のオーディオを WebSocket 経由でプッシュすると、リアルタイムの会話に対して積極的に機能するプロパティを含む、TCP のすべてのプロパティが継承されます。
TCP は、注文された信頼性の高い配信を保証します。すべてのパケットが順番に到着します。パケットが転送中に失われた場合、TCP はストリームを一時停止し、その後に送信されるものを配信する前にストリームを再送信します。これはヘッドオブラインブロッキングと呼ばれるもので、オーディオにとっては壊滅的なものです。
会話中に 1 つのパケットが失われた場合に何が起こるかを考えてみましょう。 TCP を使用すると、受信機は停止する可能性があり、数百ミリ秒間停止し、TCP を待機します。

彼は再送信します。パケットが失われた後、完全に正常に到着したオーディオは、ギャップが埋まるまで再生されずにバッファに保管されます。ユーザーには沈黙が聞こえ、その後バッファリングされた音声がバーストして聞こえます。会話のリズムが崩れてしまいます。
テキスト チャットでは、200 ミリ秒の遅延は目に見えません。音声会話では、自然な会話と気まずい会話の違いが生じます。
WebSocket にはメディア タイミングの概念がありません。スムーズに再生するには、オーディオ フレームが正確な間隔で到着する必要があります。 WebSocket はバイトを配信します。ジッター バッファー、再生タイミング、早すぎるまたは遅すぎる到着フレームを処理するメカニズムはありません。これらすべてを自分で構築する必要があり、それを適切に構築するには数年にわたるエンジニアリングの取り組みが必要です。
メディア用の輻輳制御機能は組み込まれていません。 TCP の輻輳制御アルゴリズムは、大量のデータ転送用に設計されており、パイプを埋め、損失を検出し、バックオフします。この鋸歯状パターンは、ファイルをダウンロードする場合には問題ありませんが、安定した予測可能なビットレートが必要なリアルタイム オーディオには適していません。ネットワークのパフォーマンスが低下すると、TCP はより多くのデータをバッファリングし、より強力に再試行します。これは、フレームが遅延するよりもドロップされたフレームの方が優れているライブ会話にとっては、まさに間違った戦略です。
TCP ウィンドウは不利に働きます。 TCP は、スライディング ウィンドウを使用して、送信可能な未確認データの量を制御します。パケットが失われるとウィンドウが縮小し、一貫した配信が必要なときにスループットが調整されます。損失が解消された後、ウィンドウはスナップバックしません。ウィンドウはスロースタートと輻輳回避によって控えめに増加し、回復するまでに複数回のラウンドトリップがかかります。高遅延パス (リージョン間接続など) では、各往復に時間がかかるため、この増加は特に苦痛です。その結果、急激な配信不足とその後の回復の遅れが生じます。まさに、まさにそのような状況です。

安定したスループットにより、スムーズな音声会話が途切れ途切れの会話に変わります。
WebRTC は、リアルタイムで人々の間でメディアを移動するという問題のために構築されました。特に会話に最適化する設計上の決定により、上記のすべての欠点に対処します。
損失耐性のある UDP ベースのトランスポート。 WebRTC は、RTP (リアルタイム トランスポート プロトコル) を使用して UDP 経由でメディアを送信します。パケットが失われた場合でも、ストリームは流れ続けます。 20 ミリ秒の音声フレームの欠落は、リスナーにはほとんど知覚されません。 TCP 再送信中の 200 ミリ秒の停止は発生しません。 WebRTC は、完璧な信頼性と引き換えに一貫したタイミングを実現します。これは音声にとってまさに適切なトレードオフです。
内蔵ジッターバッファー。ネットワーク ジッター (パケット到着時間の変動) は、インターネットでは避けられません。 WebRTC クライアントには、この変動を吸収する適応型ジッター バッファーが組み込まれており、再生がスムーズになるため、パケットが不均等に到着した場合でも、リスナーは連続したストリームを聞くことができます。 WebSocket を使用すると、自分で操作できます。
メディアを意識した輻輳制御。 WebRTC は、リアルタイム メディア用に特別に設計された輻輳制御アルゴリズム (Google Congestion Control、GCC など) を実装します。 TCP の積極的なフィルアンドバックオフ パターンの代わりに、GCC は一方向の遅延変動を測定して、パケット損失が発生する前に輻輳を検出します。帯域幅が低下した場合、WebRTC はストリームを停止させるのではなく、スムーズにビットレートを下げることができ、オーディオ品質をスケールダウンしたり、低いビデオ解像度に切り替えたりすることができます。
コーデックのネゴシエーションと適応。 WebRTC は、コーデックの選択、サンプル レート、チャネル構成を接続セットアップの一部として処理します。双方とも、最も効率的なエンコーディングについて合意しています。ネットワークの状態が変化すると、コーデック パラメータが適応できます。 WebSocket を使用すると、ネゴシエーション層を使用せずに、生のバイトまたは事前にエンコードされたバイトをストリーミングします。
ノイズキャンセラ

ションとエコー抑制。 WebRTC クライアントには、メディア パイプラインに組み込まれた音響エコー キャンセル (AEC)、自動ゲイン制御 (AGC)、およびノイズ抑制が含まれています。これらは、音声がネットワークに入る前に実行されます。つまり、エージェントはユーザーの環境に関係なく、クリーンな音声を受信します。 WebSocket では、これらを完全にスキップするか、個別に実装します。
NATトラバーサル。ほとんどのユーザーは NAT とファイアウォールの内側にいます。 WebRTC には、これらの障害物を乗り越えて接続を確実に確立するための ICE (Interactive Connectivity Packaging)、STUN、および TURN が含まれています。 WebSocket 接続は標準の HTTPS ポートを使用するため、同じ NAT の問題に直面することはありませんが、UDP がブロックされている場合に WebRTC が TURN over TCP/443 にフォールバックすることもでき、メディアに最適化されたすべての動作を最上位で維持できることを理解すると、この利点はなくなります。
これらの違いはいずれも、単独で管理できるように思えるかもしれません。ジッターバッファーを構築することもできます。独自の輻輳検出を実装することもできます。エコーキャンセルを前処理ステップとして追加できます。
しかし、これらのシステムは相互作用します。ジッター バッファーは再生タイミングに影響を与えます。輻輳制御はコーデックのビットレートの決定に影響します。エコー キャンセルでは、スピーカーで最近再生された音声を追跡する必要があります。 WebRTC では、これらのコンポーネントは、あらゆるプラットフォームで連携して動作するように共同設計されています。 WebSocket ベースのスタックでは、それらを段階的に統合し、ブラウザー、モバイル プラットフォーム、およびネットワーク状態にわたる相互作用をデバッグします。
これは長年にわたるエンジニアリングの結果であり、すでに解決されています。
音声 AI にとって SFU が重要な理由
WebRTC は、メディアがエンドポイント間を移動する方法を定義します。しかし、音声 AI エージェントは単純な二者間通話ではありません。 SFU (Selective Forwarding Unit) が中心に位置し、メディアをデコードまたは再エンコードせずに参加者間でルーティングします。
考える

それは旅行のようなものです。ピアツーピア WebRTC は地方の道路を運転しています。1 人か 2 人で短距離を移動する場合には問題ありません。しかし、都市を越えて参加者を接続する必要がある場合、または数十の同時セッションを処理する必要がある場合、地方の道路は拡張できません。 SFU は空港です。全員が中央ハブに接続し、目的地に効率的にルーティングされます。ニューヨークからロンドンまでは車で移動するのではなく、飛行機で移動します。そして、世界的なカバレッジが必要な場合、巨大な空港を 1 つ建設する必要はありません。各リージョンにハブを構築し、それらを高速で信頼性の高いリンクで接続します。これがまさに分散 SFU の仕組みです。
AI 音声エージェントにとって、SFU アーキテクチャはいくつかの重要な利点を提供します。
エージェントは一度接続します。各ユーザーと直接ピアツーピア WebRTC 接続を確立する代わりに、エージェントは SFU に接続します。 SFU はファンアウトを処理します。これは、エージェント インフラストラクチャが部屋ごとの同時接続数に合わせて拡張されないことを意味します。SFU がその複雑さを吸収します。
異種ネットワークの処理。ユーザーはさまざまなネットワーク (光ファイバー、携帯電話、混雑した Wi-Fi など) から接続します。 SFU は、サイマルキャストを使用して、他の加入者に影響を与えることなく各ユーザーの帯域幅に適応して、異なる品質レベルを異なる加入者に送信できます。直接接続または WebSocket リレーを使用すると、全員に同じストリームを送信することになります。
トランスコーディングの代わりに選択的に転送します。 MCU (Multipoint Conferencing Unit) は、すべての受信ストリームをデコードし、混合して、単一の出力を再エンコードします。これは CPU を大量に消費し、遅延が増加します。 SFU はパケットを転送するだけで、デコードや再エンコードは行いません。音声 AI では、ミリ秒ごとの遅延が会話の雰囲気に影響を与えるため、これは重要です。
ルーティング層での可観測性。すべてのメディアは SFU を通過するため、

クライアント側の計測を行わずに、すべての参加者の接続品質メトリクス、パケット損失率、ジッター統計、遅延測定値を取得します。このテレメトリは、運用環境でエージェントの動作をデバッグする場合に非常に役立ちます。
マルチリージョン SFU とグローバル音声 AI
音声 AI エージェントは世界中のユーザーにサービスを提供します。シンガポールのユーザーが、インフラストラクチャが米国東部にあるエージェントと通話すると、エージェントが処理を開始する前に 250 ミリ秒を超える往復ネットワーク遅延が発生します。音声会話の場合、これは容認できません。STT-LLM-TTS パイプラインの実行速度に関係なく、すべての交換に遅れを感じさせます。
分散 SFU アーキテクチャは、リージョン全体にノードをデプロイすることでこの問題を解決します。ユーザーは最も近い SFU ノードに接続し、メディア接続のネットワーク遅延を最小限に抑えます。 SFU は、消費者向けインターネット パスよりもはるかに信頼性が高く、遅延が低い最適化されたサーバー間リンクを使用して、リージョン間のルーティングを内部で処理します。
これを WebSocket で再現するのは困難です。独自の地理認識ルーティングを構築し、各リージョンにリレー サーバーを展開し、セッション アフィニティを管理し、フェイルオーバーを処理する必要があります。これは基本的に SFU のルーティング層を再作成しますが、その下にメディアに最適化されたトランスポートはありません。
LiveKit では、マルチリージョン展開はアーキテクチャ プロジェクトではなく構成オプションです。ノードは統計を報告し、リージョン認識セレクターは新しいセッションを最も近い利用可能なノードにルーティングし、スケールダウン中に接続は正常にドレインされます。同じ都市内の 2 人のユーザーを処理する同じアーキテクチャが、世界的に分散された音声 AI 展開を処理します。
ビデオ推論: 同じ議論を増幅したもの
上記のすべてはオーディオに当てはまります。ビデオの場合 — ユーザーのカメラを確認したり、ビジュアル コンテンツを共有したりできるマルチモーダル エージェントで使用されます —

ケースはさらに強力です。
ビデオはオーディオよりも帯域幅を大幅に消費します。 1.5 Mbps の 720p ビデオ ストリームは、一般的なオーディオ ストリームの帯域幅の約 50 倍です。行頭ブロッキング、不十分な輻輳制御、およびジッター バッファーの欠落の影響は、比例して増幅されます。
ここでは WebRTC のサイマルキャスト サポートが不可欠になります。パブリッシャーは複数の解像度レイヤー (720p、360p、180p など) を送信し、SFU は利用可能な帯域幅とビデオをレンダリングするサイズに基づいて各サブスクライバーに適切なレイヤーを選択します。この適応ストリームの動作は自動的に行われるため、パブリッシャーもサブスクライバーもそれを管理する必要はありません。
ビデオ フレームを処理する必要がある視覚対応 AI エージェントの場合、SFU は人間の参加者に配信されるストリーム品質に影響を与えることなく、適切な品質レベルをエージェントの処理パイプラインに転送できます。エージェント側でのリソースの効率的な使用と、ユーザーへの高品質な配信を同時に実現できます。
WebSocket はすべてにおいて間違っているわけではありません。以下のような場合に最適です。
合図。 WebRTC 自体は、シグナリング層に WebSocket (または HTTP) を使用し、メディア接続が確立される前にセッションの説明と ICE 候補を交換します。ライブキット

[切り捨てられた]

## Original Extract

WebSockets can move bytes fast, but voice AI agents need more than that. Here's why WebRTC and the SFU architecture behind it is the right transport layer for production realtime voice.

Why WebRTC beats WebSockets for realtime voice AI | LiveKit Skip to main content Placeholder text for banner height reservation on mobile Products Chevron Icon
livekit / agents 10.9 K agents 10.9 K
Blog / Engineering Why WebRTC beats WebSockets for realtime voice AI
When developers start building voice AI agents, the first architectural decision is transport: how does audio get between the user and the agent? Many reach for WebSockets because they're familiar, well-documented, and already part of most web stacks. It seems like a reasonable choice — open a socket, stream audio bytes in both directions, done.
It works in a demo. It falls apart in production.
The gap between "audio is flowing" and "this feels like a real conversation" is enormous, and it's almost entirely a transport problem. WebSockets weren't designed for realtime media. WebRTC was. That distinction matters far more than most developers expect when they start building.
What WebSockets actually give you
WebSockets provide a persistent, full-duplex TCP connection between a client and server. They're great for chat, notifications, and streaming structured data. For those use cases, they're the right tool.
But when you push raw audio over a WebSocket, you inherit every property of TCP — including the ones that actively work against realtime conversation.
TCP guarantees ordered, reliable delivery. Every packet arrives, and it arrives in sequence. If a packet is lost in transit, TCP pauses the stream and retransmits it before delivering anything that came after. This is called head-of-line blocking, and for audio, it's devastating.
Consider what happens when a single packet is lost during a conversation. With TCP, the receiver stalls — possibly for hundreds of milliseconds — waiting for the retransmission. The audio that arrived perfectly fine after the lost packet sits in a buffer, unplayed, until the gap is filled. The user hears silence, then a burst of buffered audio. The conversational rhythm breaks.
In a text chat, a 200ms delay is invisible. In a voice conversation, it's the difference between a natural exchange and an awkward one.
WebSockets have no concept of media timing. Audio frames need to arrive at precise intervals for smooth playback. WebSockets deliver bytes — there's no jitter buffer, no playout timing, no mechanism to handle frames that arrive too early or too late. You have to build all of that yourself, and building it well is a multi-year engineering effort.
There's no built-in congestion control for media. TCP's congestion control algorithm is designed for bulk data transfer: it fills the pipe, detects loss, and backs off. This sawtooth pattern is fine for downloading files but terrible for realtime audio, where you need a steady, predictable bitrate. When the network degrades, TCP's response is to buffer more data and retry harder — exactly the wrong strategy for a live conversation where a dropped frame is better than a late one.
TCP windowing works against you. TCP uses a sliding window to control how much unacknowledged data can be in flight. When packets are lost, the window shrinks, throttling throughput right when you need consistent delivery. After the loss clears, the window doesn't snap back — it grows conservatively through slow start and congestion avoidance, taking multiple round trips to recover. On high-latency paths (like cross-region connections), this ramp-up is especially painful because each round trip takes longer. The result is bursts of underdelivery followed by slow recovery — exactly the kind of inconsistent throughput that turns a smooth voice conversation into a stuttering one.
WebRTC was purpose-built for the problem of moving media between people in realtime. It addresses every shortcoming above with design decisions that specifically optimize for conversation.
UDP-based transport with loss tolerance. WebRTC sends media over UDP using RTP (Real-time Transport Protocol). When a packet is lost, the stream keeps flowing. A missing 20ms audio frame is nearly imperceptible to a listener; a 200ms stall while TCP retransmits is not. WebRTC trades perfect reliability for consistent timing, which is exactly the right trade-off for voice.
Built-in jitter buffers. Network jitter — variation in packet arrival times — is unavoidable on the internet. WebRTC clients include adaptive jitter buffers that absorb this variation, smoothing out playback so the listener hears a continuous stream even when packets arrive unevenly. With WebSockets, you're on your own.
Media-aware congestion control. WebRTC implements congestion control algorithms (like Google Congestion Control, GCC) that are specifically designed for realtime media. Instead of TCP's aggressive fill-and-backoff pattern, GCC measures one-way delay variation to detect congestion before packet loss occurs. When bandwidth drops, WebRTC can reduce bitrate smoothly — scaling down audio quality or switching to a lower video resolution — rather than stalling the stream.
Codec negotiation and adaptation. WebRTC handles codec selection, sample rates, and channel configuration as part of the connection setup. Both sides agree on the most efficient encoding. When network conditions change, the codec parameters can adapt. With WebSockets, you're streaming raw or pre-encoded bytes with no negotiation layer.
Noise cancellation and echo suppression. WebRTC clients include acoustic echo cancellation (AEC), automatic gain control (AGC), and noise suppression built into the media pipeline. These run before audio enters the network, which means the agent receives clean audio regardless of the user's environment. With WebSockets, you either skip these entirely or implement them separately.
NAT traversal. Most users are behind NATs and firewalls. WebRTC includes ICE (Interactive Connectivity Establishment), STUN, and TURN to reliably establish connections through these obstacles. WebSocket connections don't face the same NAT issues since they use standard HTTPS ports, but this advantage disappears when you realize WebRTC can also fall back to TURN over TCP/443 when UDP is blocked — and still maintain all its media-optimized behavior on top.
Any one of these differences might seem manageable in isolation. You could build a jitter buffer. You could implement your own congestion detection. You could add echo cancellation as a preprocessing step.
But these systems interact. The jitter buffer feeds into playout timing. Congestion control affects codec bitrate decisions. Echo cancellation needs to track what audio was recently played to the speaker. In WebRTC, these components are co-designed to work together across every platform. In a WebSocket-based stack, you're integrating them piecemeal and debugging their interactions across browsers, mobile platforms, and network conditions.
This is years of engineering — and it's already solved.
Why an SFU matters for voice AI
WebRTC defines how media gets between endpoints. But voice AI agents aren't simple two-party calls. An SFU (Selective Forwarding Unit) sits at the center, routing media between participants without decoding or re-encoding it.
Think of it like travel. Peer-to-peer WebRTC is driving local roads — it's fine when you're going a short distance with one or two people. But when you need to connect participants across cities, or handle dozens of concurrent sessions, local roads don't scale. An SFU is the airport. Everyone connects to a central hub that efficiently routes them where they need to go. You don't drive from New York to London — you fly. And when you need global coverage, you don't build one massive airport. You build hubs in each region and connect them with fast, reliable links. That's exactly how a distributed SFU works.
For AI voice agents, the SFU architecture provides several critical advantages.
The agent connects once. Instead of establishing a direct peer-to-peer WebRTC connection with each user, the agent connects to the SFU. The SFU handles the fan-out. This means agent infrastructure doesn't scale with the number of concurrent connections per room — the SFU absorbs that complexity.
Heterogeneous network handling. Users connect from different networks — some on fiber, some on cellular, some on congested Wi-Fi. An SFU can send different quality levels to different subscribers using simulcast , adapting to each user's bandwidth without affecting others. With a direct connection or a WebSocket relay, you're stuck sending the same stream to everyone.
Selective forwarding instead of transcoding. An MCU (Multipoint Conferencing Unit) decodes all incoming streams, mixes them, and re-encodes a single output. This is CPU-intensive and adds latency. An SFU just forwards packets — no decode, no re-encode. For voice AI, where every millisecond of latency affects the feel of the conversation, this matters.
Observability at the routing layer. Because all media flows through the SFU, you get connection quality metrics, packet loss rates, jitter statistics, and latency measurements for every participant without any client-side instrumentation. This telemetry is invaluable for debugging agent behavior in production.
Multi-region SFUs and global voice AI
Voice AI agents serve users globally. A user in Singapore talking to an agent whose infrastructure is in US-East will experience 250+ ms of round-trip network latency before the agent even starts processing. For a voice conversation, that's unacceptable — it makes every exchange feel laggy, regardless of how fast the STT-LLM-TTS pipeline runs.
A distributed SFU architecture solves this by deploying nodes across regions. The user connects to the nearest SFU node, getting the lowest possible network latency for their media connection. The SFU handles routing between regions internally, with optimized server-to-server links that are far more reliable and lower-latency than consumer internet paths.
This is difficult to replicate with WebSockets. You'd need to build your own geo-aware routing, deploy relay servers in each region, manage session affinity, and handle failover — essentially re-creating the SFU's routing layer but without the media-optimized transport underneath.
With LiveKit, multi-region deployment is a configuration option, not an architecture project. Nodes report their stats, a region-aware selector routes new sessions to the closest available node, and connections drain gracefully during scale-down. The same architecture that handles two users in the same city handles a globally distributed voice AI deployment.
Video inference: the same argument, amplified
Everything above applies to audio. For video — used in multimodal agents that can see the user's camera or share visual content — the case is even stronger.
Video is orders of magnitude more bandwidth-intensive than audio. A 720p video stream at 1.5 Mbps is roughly 50x the bandwidth of a typical audio stream. The consequences of head-of-line blocking, poor congestion control, and missing jitter buffers are amplified proportionally.
WebRTC's simulcast support becomes essential here. A publisher sends multiple resolution layers (for example, 720p, 360p, and 180p), and the SFU selects the appropriate layer for each subscriber based on their available bandwidth and the size they're rendering the video at. This adaptive stream behavior is automatic — neither the publisher nor the subscriber needs to manage it.
For vision-capable AI agents that need to process video frames, the SFU can forward the appropriate quality level to the agent's processing pipeline without affecting the stream quality delivered to human participants. You get efficient resource usage on the agent side and high-quality delivery to users, simultaneously.
WebSockets aren't wrong for everything. They're the right choice for:
Signaling. WebRTC itself uses WebSockets (or HTTP) for the signaling layer — exchanging session descriptions and ICE candidates before the media connection is established. LiveKit

[truncated]
