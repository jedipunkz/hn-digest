---
source: "https://techcommunity.microsoft.com/blog/linuxandopensourceblog/announcing-the-open-source-release-of-ml-video-codec-mlvc/4539875"
hn_url: "https://news.ycombinator.com/item?id=49048733"
title: "The Open-Source Release of ML Video Codec (MLVC)"
article_title: "Announcing the Open-Source Release of ML Video Codec (MLVC) | Microsoft Community Hub"
author: "jongalloway2"
captured_at: "2026-07-25T16:50:51Z"
capture_tool: "hn-digest"
hn_id: 49048733
score: 5
comments: 1
posted_at: "2026-07-25T15:59:59Z"
tags:
  - hacker-news
  - translated
---

# The Open-Source Release of ML Video Codec (MLVC)

- HN: [49048733](https://news.ycombinator.com/item?id=49048733)
- Source: [techcommunity.microsoft.com](https://techcommunity.microsoft.com/blog/linuxandopensourceblog/announcing-the-open-source-release-of-ml-video-codec-mlvc/4539875)
- Score: 5
- Comments: 1
- Posted: 2026-07-25T15:59:59Z

## Translation

タイトル: ML ビデオ コーデック (MLVC) のオープンソース リリース
記事のタイトル: ML ビデオ コーデック (MLVC) のオープンソース リリースの発表 |マイクロソフト コミュニティ ハブ
説明:  
ビデオ コーデックは、送信または保存のためにビデオを圧縮し、帯域幅とストレージの要件を軽減します。 MLVC は、最新の機械学習ベースのコーデックです。

記事本文:
ML ビデオ コーデック (MLVC) のオープンソース リリースの発表 | Microsoft コミュニティ ハブ サイド メニューを開く コンテンツにスキップ テクニカル コミュニティ コミュニティ ハブ 製品 トピック ブログ イベント スキル ハブ コミュニティ 登録 サインイン Microsoft コミュニティ ハブ
ML ビデオ コーデック (MLVC) のオープンソース リリースの発表
Microsoft で開発されたニューラル ネットワーク ベースのビデオ コーデックである ML Video Codec (MLVC) が、MIT ライセンスの下でオープンソース コミュニティにリリースされることを発表できることを嬉しく思います。
ビデオ コーデックは、送信または保存のためにビデオを圧縮し、帯域幅とストレージの要件を軽減します。 MLVC は、従来のコーデックよりも大幅に少ない帯域幅を使用する最新の機械学習ベースのコーデックで、特に制約のあるネットワークや信頼性の低いネットワークでのストリーミングとビデオ通話の品質を向上させると同時に、配信とストレージのコストを削減します。
MLVC は、2021 年から Microsoft Research によってオープンソース化されている NVC (ニューラル ビデオ コーデック) の DCVC (Deep Contextual Video Compression) ファミリの製品イテレーションであり、圧縮効率の向上、汎用ニューラル プロセッシング ユニット (NPU) でのリアルタイム パフォーマンス、およびクロスプラットフォーム サポートを備えています。私たちは最近、この成果を論文「MLVC: 実世界展開のためのマルチプラットフォーム学習ビデオ コーデック」で発表しました。私たちがソース コードを公開するのは、次世代のビデオ コーディングがオープンに構築されると信じているためであり、一般の開発者コミュニティだけでなく、研究者、ビデオ コーデック エンジニア、プラットフォーム ベンダー、製品チームなど、より広範なコミュニティに私たちと一緒に構築してもらいたいと考えているからです。
従来のビデオ コーデック (H.264/AVC、H.265/HEVC など) は長い間業界に貢献してきましたが、世代ごとに段階的に利益を得るには多大なエンジニアリング努力が必要であり、一般に利用可能になるまでに数年かかる専用のハードウェアが必要です。 MLVC 担当者

従来のプリミティブ (動き推定、変換、エントロピー モデリング) にエンドツーエンドの学習ニューラル圧縮を組み合わせ、レート歪み目標に対して直接トレーニングし、汎用 NPU デバイスで実行します。
以下の表は、MLVC と一般的なビデオ コーデックを比較し、ビットレートが低く、その結果として帯域幅とストレージが節約されることを示しています。
たとえば、30 fps の 360p ビデオの場合、H.264 では 1 Mbps が必要ですが、MLVC では同等の品質を得るために約 122 kbps が必要になります。これは、リアルタイム条件でのビットレートの約 8 分の 1 です。推論計算は、360p と 540p の解像度でほぼ同じに保たれました。これらの結果は、P.910 主観テストに基づいており、当社が開発し、最近オープンソース プロジェクトとしてリリースしたビデオ会議データセット (VCD) データセットにも基づいています。
次のビデオ デモは、同じ 200kbps のビットレートでの H.265/HEVC と比較して、MLVC によって達成される品質向上の程度を示しています (より良いデモを表示するには、新しいウィンドウを開いてご覧ください)
MLVC は、ビデオ圧縮効率を超えて、次の機能も提供します。
NPU ファーストの設計。 MLVC は、最新のデバイスにすでに搭載されている AI アクセラレータ (Apple Neural Engine、Qualcomm、Intel NPU) でほぼ完全に動作するように構築されており、NPU 使用率は 50% 以下で、システムの残りの部分に NPU のヘッドルームを残します。
ターゲットの動作ポイントでのリアルタイム実行。 Apple、Intel、Qualcomm ハードウェアで 30 fps で 540p を実証しました。
スケーリング則の軌道。経験的に、MLVC のコーディング効率は、モデル容量の増加とトレーニング コンピューティングの追加によって向上します。
手動で調整したレート歪み最適化ヒューリスティックを必要とせず、すぐに使用できるコンテンツ適応型の動作。
すでに Microsoft Teams で実行されています
MLVC は、Microsoft にとって単なる研究コンセプトではありません。私たちは現在ローリング中です

Microsoft Teams では、アクティブなテレメトリと A/B テストを使用した実際のピアツーピア ビデオ通話で検証されています。この統合は、ハードウェアまたは信頼性の制約に対する従来のビデオ コーデックへのフォールバックと並行して実行されます。これは、実際の製品に必要な混合環境の展開です。このロールアウトからのスケーリングと信頼性に関する洞察は、コーデックとそのロードマップを形成しています。
私たちだけですべてのユースケース、すべてのデバイスクラス、すべてのコンテンツドメインをカバーすることはできません。それが、私たちが MLVC をオープンソース化している理由です。次のことに取り組んでいる場合:
ストリーミング、ビデオ オン デマンド (VOD)、またはライブ ブロードキャスト
リアルタイムのコミュニケーションと会議
クラウド ゲームまたはリモート レンダリング
監視、ドローン、ロボット工学
モバイル キャプチャ、拡張現実 (AR) / 仮想現実 (VR)、またはボリューム ビデオ
コーデック ハードウェア、NPU、または推論ランタイム
ぜひご協力をお願いいたします。モデルの改善、トレーニング レシピ、新しいプラットフォーム ポートと変換ターゲット、ランタイム バックエンド、ドメイン固有の微調整、評価ツール、バグ レポート、シナリオに足りないものについてのフィードバックなど、あらゆる種類の貢献を歓迎します。私たちは特に、NPU の適用範囲を拡大するプラットフォームのポートと、レート歪みのフロンティアを押し上げる効率の改善を歓迎します。
MLVC リポジトリは https://github.com/microsoft/mlvc で入手でき、MIT ライセンスに基づいて共有されます。これには次のものが含まれます。
完全なネットワーク アーキテクチャの MLVC モデル ソース コード。
トレーニング済みのモデルの重みを実行する準備ができています。
出荷モデルの作成に使用されるトレーニング スクリプト。
MLVC の再現と改善に役立つトレーニング データ収集のドキュメント。
さまざまな NPU とランタイムをターゲットとするプラットフォーム変換スクリプト。
イシューとプルリクエストは初日からオープンになります。
フォローアップ リリースでは C++ コーデック ライブラリが追加され、実際のアプリケーションへの統合が簡素化されます。

カチオン。
MLVC に関する当社の長期的な目標は、ビデオ コンテンツの全範囲にわたって従来の最高のコーデックのコーディング効率を満たすかそれを超えるオープンな学習済みビデオ コーデックを作成し、クライアントおよびクラウド デバイスにすでに出荷されている AI ハードウェア上で効率的に実行し、最新の ML システムのようにコンピューティングに合わせて拡張し、標準化サイクルのペースではなく ML コミュニティのペースでオープンに進化することです。
短期的には、540p リアルタイム パフォーマンスの安定化、ハードウェア カバレッジの拡大、損失耐性の向上を意味します。中期的には: 1080p などの高解像度、およびより広範なストリーミング シナリオ。長期的には: 必要に応じて従来のスタックを有意義に置き換えるオープン ビデオ コーデック エコシステム。
私たちだけで MLVC の未来を作ることはできません。皆さんがここに来てくれてうれしいです。皆さんと一緒に次世代のビデオ コーデックを構築できることを楽しみにしています。
MLVC は、マイクロソフトでこのプロジェクトに取り組んでいる次の素晴らしい人々によって提供されています: Ross Cutler、Ando Saabas、Tanel Pärnamaa、Ardi Loot、Haiyan Xie、Lauri Ehrenpreis、Andrei Znobishchev、Martin Lumiste、Evgenii Indenbom、Yan Lu、Bin Li、Jiahao Li、Naba Kumar、Babak Naderi、Juhee Cho、Badal Yadav、ジョウ・ジンシン、ディン・ティエンユー、パトリック・グレゴリー。
教育者のトレーニングと能力開発
学生と保護者向けの特典
AI マーケットプレイス アプリのサポート
あなたのプライバシーの選択 消費者の健康 プライバシー サイトマップ

## Original Extract

&nbsp;
Video codecs compress video for transmission or storage, reducing bandwidth and storage requirements. MLVC is a modern machine-learning-based codec...

Announcing the Open-Source Release of ML Video Codec (MLVC) | Microsoft Community Hub Open Side Menu Skip to content Tech Community Community Hubs Products Topics Blogs Events Skills Hub Community Register Sign In Microsoft Community Hub
Announcing the Open-Source Release of ML Video Codec (MLVC)
We are excited to announce that ML Video Codec (MLVC) — a neural-network-based video codec developed at Microsoft — is being released to the open-source community under the MIT License.
Video codecs compress video for transmission or storage, reducing bandwidth and storage requirements. MLVC is a modern machine-learning-based codec that uses substantially less bandwidth than conventional codecs, improving streaming and video-call quality—especially on constrained or unreliable networks—while lowering delivery and storage costs.
MLVC is the product iteration of DCVC (Deep Contextual Video Compression) family of NVC (Neural Video Codec), open sourced by Microsoft Research since 2021, with improved compression efficiency, real-time performance on commodity Neural Processing Units (NPUs), and cross-platform support. We recently published this work in the paper MLVC: Multi-platform Learned Video Codec for Real-World Deployment . We are releasing the source code because we believe the next generation of video coding will be built openly, and we want the broader community — researchers, video codec engineers, platform vendors, product teams, as well as general developer community — to build it with us.
Traditional video codecs (e.g., H.264/AVC, H.265/HEVC) have served the industry for a long time, but each generation requires enormous engineering effort for incremental gains and needs dedicated hardware which takes years to become commonly available. MLVC replaces conventional primitives — motion estimation, transforms, entropy modeling — with end-to-end learned neural compression , trained directly against rate-distortion objectives, and run on general-purpose NPU devices.
The table below compares MLVC to popular video codecs, showing its lower bitrate and resulting savings in bandwidth and storage.
For example, for 360p video at 30 fps, where H.264 requires 1 Mbps, MLVC requires roughly 122 kbps for equivalent quality — about one-eight the bitrate under real-time conditions. The inference compute was kept approximately equal for the 360p and 540p resolutions. These results are based on a P.910 subjective test and are based on the Video Conferencing Dataset (VCD) dataset that we developed and also recently released as an open-source project.
The following video demo illustrates the extent of quality enhancement achieved by MLVC relative to H.265/HEVC at the same bitrate of 200kbps (please watch by opening in a new window for better demonstration)
Beyond video compression efficiency, MLVC also offers:
NPU-first design. MLVC is built to run almost entirely on the AI accelerators already shipping in modern devices — Apple Neural Engine, Qualcomm and Intel NPUs — at no more than 50% NPU utilization, leaving NPU headroom for the rest of the system.
Real-time execution at the targeted operating points. Demonstrated 540p at 30 fps on Apple, Intel, and Qualcomm hardware.
A scaling-law trajectory. Empirically, MLVC's coding efficiency improves with increased model capacity and additional training compute.
Content-adaptive behavior out of the box, without hand-tuned Rate Distortion Optimization heuristics.
Already running in Microsoft Teams
MLVC is more than just a research concept for Microsoft. We are currently rolling it out in Microsoft Teams, where it is being validated on real peer-to-peer video calls with active telemetry and A/B testing. The integration runs alongside fallback to conventional video codecs for hardware or reliability constraints — the kind of mixed-environment deployment that real products need. The scaling and reliability insights from this rollout are shaping the codec and its roadmap.
We can not cover every use case, every device class, or every content domain by ourselves. That is why we are open sourcing MLVC. If you work on:
Streaming, Video On Demand (VOD), or live broadcast
Real-time communication and conferencing
Cloud gaming or remote rendering
Surveillance, drones, or robotics
Mobile capture, Augmented Reality (AR) / Virtual Reality (VR), or volumetric video
Codec hardware, NPUs, or inference runtimes
We would love your help. Contributions of all kinds are welcome, including model improvements, training recipes, new platform ports and conversion targets, runtime backends, domain-specific fine-tunes, evaluation tooling, bug reports, and feedback on what is missing for your scenario. We particularly welcome platform ports that expand NPU coverage and efficiency improvements that push the rate-distortion frontier.
The MLVC repository is available at https://github.com/microsoft/mlvc and shared under the MIT License . It includes:
MLVC model source code of the full network architecture.
Trained model weights ready to run.
Training scripts used to produce shipping models.
Training data collection documentation to help reproduce and improve MLVC.
Platform conversion scripts to target different NPUs and runtimes.
Issues and pull requests will be open from day one.
A follow-up release will add a C++ codec library, simplifying integration into real-world applications.
Our long-term goal with MLVC is to create an open, learned video codec that meets or exceeds the coding efficiency of the best conventional codecs across the full range of video content, runs efficiently on the AI hardware already shipping in client and cloud devices, scales with compute the way modern ML systems do, and evolves in the open at the pace of the ML community rather than the pace of standardization cycles.
In the near term that means stabilizing 540p real-time performance, expanding hardware coverage, and improving loss resilience. In the medium term: higher resolution, e.g., 1080p, and broader streaming scenarios. In the long term: an open video codec ecosystem that meaningfully replaces legacy stacks where it makes sense to.
We can't create the future of MLVC alone. We're glad you're here, and we are looking forward to building the next-generation video codec with you.
MLVC is brought to you by the following awesome folks working on the project at Microsoft: Ross Cutler, Ando Saabas, Tanel Pärnamaa, Ardi Loot, Haiyan Xie, Lauri Ehrenpreis, Andrei Znobishchev, Martin Lumiste, Evgenii Indenbom, Yan Lu, Bin Li, Jiahao Li, Naba Kumar, Babak Naderi, Juhee Cho, Badal Yadav, Jinxin Zhou, Tianyu Ding, Patrick Gregory.
Educator training and development
Deals for students and parents
Support for AI marketplace apps
Your Privacy Choices Consumer Health Privacy Sitemap
