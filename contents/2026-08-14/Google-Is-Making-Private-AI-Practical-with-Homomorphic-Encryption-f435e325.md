---
source: "https://blog.google/security/how-google-is-making-private-ai-practical-with-homomorphic-encryption/"
hn_url: "https://news.ycombinator.com/item?id=49300314"
title: "Google Is Making Private AI Practical with Homomorphic Encryption"
article_title: "How Google is Making Private AI Practical with Homomorphic Encryption"
author: "u1hcw9nx"
captured_at: "2026-08-14T16:41:33Z"
capture_tool: "hn-digest"
hn_id: 49300314
score: 12
comments: 7
posted_at: "2026-08-14T15:43:10Z"
tags:
  - hacker-news
  - translated
---

# Google Is Making Private AI Practical with Homomorphic Encryption

- HN: [49300314](https://news.ycombinator.com/item?id=49300314)
- Source: [blog.google](https://blog.google/security/how-google-is-making-private-ai-practical-with-homomorphic-encryption/)
- Score: 12
- Comments: 7
- Posted: 2026-08-14T15:43:10Z

## Translation

タイトル: Google、準同型暗号化でプライベートAIを実用化
記事のタイトル: Google は準同型暗号化でプライベート AI を実用化する方法
説明: 本日は、プライベート コンピューティング ツールキットに追加された最新の強力なツールである HEIR を紹介できることを嬉しく思います。 HEIR は、暗号的に安全なプライベート AI 推論を可能にするオープンソース コンパイラーです。準同型暗号化 AI の成長に伴って新たなメリットが生まれる中、プライバシーとセキュリティのバランスをとることが最優先事項です。
[切り捨てられた]

記事本文:
メインコンテンツにスキップ
セキュリティ
Google は準同型暗号化でプライベート AI をどのように実用化しているか
グローバル（英語）
アフリカ (英語)
オーストラリア (英語)
ブラジル (ポルトガル語)
カナダ (英語)
カナダ (フランス)
チェスコ (チェシュティナ)
ドイチュラント (ドイツ)
エスパーニャ (スペイン語)
フランス (フランセ)
ギリシャ (Ελληνικά)
インド (英語)
インドネシア (インドネシア語)
アイルランド (英語)
イタリア (イタリアーノ)
日本 (日本語)
대한민국 (한국어)
ラテンアメリカ (スペイン語)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
メナ (英語)
オランダ (オランダ)
ニュージーランド (英語)
ポルスカ (ポルスキー)
ポルトガル (ポルトガル人)
ルーマニア (Romană)
スヴェリゲ (スヴェンスカ)
ประเทศไทย (ไทย)
トゥルキエ (トゥルクチェ)
台灣 (中文)
Google は準同型暗号化でプライベート AI をどのように実用化しているか
本日は、プライベート コンピューティング ツールキットに追加された最新の強力なツールである HEIR を紹介できることを嬉しく思います。 HEIR は、暗号的に安全なプライベート AI 推論を可能にするオープンソース コンパイラーです。
AI の成長に伴って新たなメリットが生まれる中、プライバシーとセキュリティのバランスが最優先事項となっています。エンドツーエンド暗号化などの標準的な保護にはトレードオフがあります。ユーザーデータはデータ侵害から保護できますが、その場合、サービスプロバイダーはスパムやウイルスの検出など、データに依存する機能を提供できません。医療や金融などの重要なセクターはこうしたリスクをさらに忌避しており、厳格な規制により機関間のデータ共有が制限されています。ローカル処理など、同じ機能を提供する代替メカニズムは、ローカル デバイスの機能とサービス プロバイダーの IP の感度によって制限されます。独自の AI をデバイスに配布すると、モデルが漏洩する危険があります。
これらの問題の解決策は、準同型暗号化です。

このテクノロジーは、暗号化されたデータに対して直接計算を実行できるようにすることで、このトレードオフを根本的に変更する、急速に成熟しています。サーバーは、基礎となる情報を公開することなく暗号文を処理し、暗号化された結果を返すことができます。たとえば、クラウド サービスは、ユーザーの特徴を確認できなくても、コンテンツの推奨を提供できます。これは誇張ではありません。この投稿で取り上げたデモの 1 つは、まさにこれを実行します。しかし、準同型暗号には少なからぬコストのオーバーヘッドがある一方で、機能とプライバシーのトレードオフはコストの問題に移ります。また、準同型暗号のコストは急速に低下しています。
差分プライバシーやプライベート セット メンバーシップから個人情報の取得や Google Cloud 上の安全なエンクレーブに至るまで、Google のプライバシー テクノロジーにおけるイノベーションの歴史は、常にユーザー データの保護に重点を置いてきました。準同型暗号化は、私たちがプライベート コンピューティング ツールキットに追加しているもう 1 つの強力なツールです。個人情報の検索と同様、またハードウェアベースのソリューションとは対照的に、準同型暗号化の強力なセキュリティとプライバシーの保証は純粋に暗号的なものです。ただし、準同型暗号を効率的に使用するために既存のプログラムを手動で変換するには、暗号学者のチームが必要です。
ユーザビリティの課題を克服し、準同型暗号がもたらす機会を前進させるために、Google の研究者とエンジニアは HEIR コンパイラ プロジェクトを構築しました。 HEIR (Homomorphic Encryption Intermediate Representation) は、準同型暗号化のためのオープンソースのコンパイラ ツールチェーンおよび開発プラットフォームです。特に、HEIR は、暗号化されていないデータで動作する事前トレーニングされた AI モデルを、暗号化された入力で動作するように変換できます。私たちのビジョンは、HEIR をワンクリック ソリューションにして、専門家でなくても暗号化された推論を製品に組み込めるようにすることです。

アプリケーションについて。
2023 年の意向を発表して以来、準同型暗号化コミュニティが HEIR を採用するのを見てきました。当社は、Belfort、Niobium、Cornami、Optalysys など、準同型暗号化用のハードウェア アクセラレータを開発する企業と提携しています。これらの取り組みの成果は以下のデモに示されており、近い将来、これらのアクセラレータのレイテンシの利点を実証する予定です。 HEIR は生産的な研究プラットフォームにもなりました。 HEIR を基盤とすることで、暗号技術者は特定の最適化に集中し、既存のインフラストラクチャをテスト、ベンチマーク、比較に使用できます。これにより、ジョージア工科大学、カーネギーメロン大学、カリフォルニア大学サンタバーバラ校、イリノイ工科大学パデュー校、エディンバラ大学、清華大学などとのコラボレーションが実現しました。現在までに、4 つの査読済み出版物が HEIR に基づいて作成されており、さらに多くの出版物が準備中です。HEIR は数多くの引用を蓄積しています。
準同型暗号化がどこまで進歩したかを示すために、4 つのプライベート推論アプリケーションを共有します。各アプリケーションは HEIR でコンパイルされており、遅延数値はシングルスレッド CPU で示されています。すべての例のソース コードは、GitHub リポジトリで入手できます。
ディープ ラーニング レコメンデーション モデルにより、ベルフォート研究所、LG、ニューヨーク大学との共同作業により、プライベート コンテンツ レコメンデーションの提供が可能になります。
クレジット カード不正検出: Niobium およびhardshell.ai と協力して、クレジット カード不正検出機能をコンパイルしました。
脅威の侵入: Niobium と協力して、暗号化されたネットワーク トラフィックの異常検出用に Kitsune システムをコンパイルしました。これにより、サービス プロバイダーは、ネットワーク パケットの内容をサービス プロバイダーに明かすことなく、異常を検出できます。
Hotword Detector: Belfort Labs と協力して、

ホットワード検出モデルを構築しました。これにより、音声によってトリガーされる AI エージェントが、音声録音のプライバシーを保護しながらホットワードを認識できるようになります。
ソフトウェア業界が AI によるセキュリティとプライバシーの変化に適応する中、当社の研究チームは、開発が容易で実行が速く、業界全体で普及する準同型暗号化の実現に取り組んでいます。
エージェントセキュリティ時代におけるレッドチームの役割の進化
Google では、レッド チームが常にセキュリティの最先端で活動してきました。私たちは過去に私たちの旅を共有しました。Hac で紹介された一か八かの作戦から…
Influence Operations 速報 2026 年第 2 四半期
この速報には、2026 年第 2 四半期に当社のプラットフォームで終了した調整影響作戦キャンペーンが含まれています。最終更新日は 2026 年 6 月 30 日でした。4 月に 13 件の YouTube を終了しました。
アップデートするたびに強化: AI 時代に Chrome とウェブをどのように安全にするか
発見から修正まで: 自動パッチで保守者の負担を軽減
2016 年の発表以来、OSS-Fuzz は数万件のバグを発見して報告することで、オープンソースの安全性を高めることに大きく貢献してきました。しかし、さらなる脆弱性を発見…
ゼロを超える: エンタープライズ セキュリティの新しいパラダイム
AI エージェントを Chrome Enterprise セキュリティ管理に導入
Google の詳細
Google 製品
グローバル（英語）
アフリカ (英語)
オーストラリア (英語)
ブラジル (ポルトガル語)
カナダ (英語)
カナダ (フランス)
チェスコ (チェシュティナ)
ドイチュラント (ドイツ)
エスパーニャ (スペイン語)
フランス (フランセ)
ギリシャ (Ελληνικά)
インド (英語)
インドネシア (インドネシア語)
アイルランド (英語)
イタリア (イタリアーノ)
日本 (日本語)
대한민국 (한국어)
ラテンアメリカ (スペイン語)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
メナ (英語)
オランダ (オランダ)
ニュージーランド (英語)
ポルスカ (ポー

イスキ）
ポルトガル (ポルトガル人)
ルーマニア (Romană)
スヴェリゲ (スヴェンスカ)
タイ（タイ語）
トルコ (テュルクチェ)
台灣 (中文)

## Original Extract

Today we're excited to showcase HEIR, the latest powerful tool added to our Private Computing Toolkit. HEIR is an open source compiler that unlocks cryptographically-secure private AI inference.Homomorphic encryptionAs new benefits emerge with the growth of AI, balancing privacy and security is top
[truncated]

Skip to main content
Security
How Google is Making Private AI Practical with Homomorphic Encryption
Global (English)
Africa (English)
Australia (English)
Brasil (Português)
Canada (English)
Canada (Français)
Česko (Čeština)
Deutschland (Deutsch)
España (Español)
France (Français)
Greece (Ελληνικά)
India (English)
Indonesia (Bahasa Indonesia)
Ireland (English)
Italia (Italiano)
日本 (日本語)
대한민국 (한국어)
Latinoamérica (Español)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
MENA (English)
Nederlands (Nederland)
New Zealand (English)
Polska (Polski)
Portugal (Português)
România (Română)
Sverige (Svenska)
ประเทศไทย (ไทย)
Türkiye (Türkçe)
台灣 (中文)
How Google is Making Private AI Practical with Homomorphic Encryption
Today we're excited to showcase HEIR , the latest powerful tool added to our Private Computing Toolkit. HEIR is an open source compiler that unlocks cryptographically-secure private AI inference.
As new benefits emerge with the growth of AI, balancing privacy and security is top of mind. Standard protections like end-to-end encryption present a trade-off: user-data can be protected from data breaches, but then the service provider cannot provide features that depend on the data, such as spam or virus detection. Critical sectors like healthcare and finance are even more averse to these risks, and strict regulations limit data sharing across institutions. Alternative mechanisms to provide the same features, like local processing, are limited by the capabilities of the local device and the sensitivity of the service provider's IP. Shipping proprietary AI to a device risks leaking the model.
A solution to these issues is homomorphic encryption , a rapidly maturing technology that fundamentally alters this trade-off by allowing computations to be performed directly on encrypted data. Servers can process ciphertexts and return encrypted results without exposing any underlying information. For example, a cloud service can provide content recommendations without being able to see the user's features. This is no exaggeration: one of the demos featured in this post does exactly this. But while homomorphic encryption has a nontrivial cost overhead, it shifts the capability/privacy trade-off to a question of cost. And the cost of homomorphic encryption is rapidly decreasing.
Google’s history of innovations in privacy technology—from differential privacy and private set membership to private information retrieval and secure enclaves on Google Cloud —has always focused on securing user data. Homomorphic encryption is another powerful tool we're adding to our private computing toolkit. Like private information retrieval, and in contrast to hardware-based solutions, homomorphic encryption's strong security and privacy guarantees are purely cryptographic. However, manually converting an existing program to use homomorphic encryption efficiently requires a team of cryptographers.
To overcome the usability challenges and advance the opportunity homomorphic encryption provides, researchers and engineers at Google built the HEIR compiler project . HEIR (Homomorphic Encryption Intermediate Representation) is an open-source compiler toolchain and development platform for homomorphic encryption. In particular, HEIR can convert pre-trained AI models that operate on unencrypted data to operate on encrypted inputs. Our vision is to make HEIR a one-click solution to enable non-experts to incorporate encrypted inference into production applications.
Since announcing our intentions in 2023 , we’ve seen the homomorphic encryption community embrace HEIR. We have partnered with companies developing hardware accelerators for homomorphic encryption, including Belfort , Niobium , Cornami , and Optalysys . The fruits of those efforts are shown in our demos below, and we plan to demonstrate the latency benefits of these accelerators in the near future. HEIR has also become a productive research platform. By building on HEIR, cryptographers can focus on their specific optimization and use the existing infrastructure for testing, benchmarking, and comparisons. This has resulted in collaborations with Georgia Tech, Carnegie Mellon, UC Santa Barbara, Illinois Institute of Technology, Purdue, the University of Edinburgh, Tsinghua University, and others. To date, four peer-reviewed publications were built on HEIR, with more in preparation, and HEIR has accumulated numerous citations .
To demonstrate how far homomorphic encryption has come, we’re sharing four private inference applications. Each application was compiled with HEIR, and latency numbers are presented for a single-threaded CPU. The source code for all examples is available in our GitHub repository .
A Deep Learning Recommendation Model unlocks serving private content recommendations, joint work with Belfort Labs , LG , and New York University .
Credit card fraud detection: Together with Niobium and hardshell.ai , we compiled a credit card fraud detector.
Threat intrusion: Together with Niobium we compiled the Kitsune system for anomaly detection of encrypted network traffic. This allows a service provider to detect anomalies without revealing the contents of network packets to the service provider.
Hotword Detector: Together with Belfort Labs we compiled a hotword detection model, which could allow an audio-triggered AI agent to recognize hotwords while protecting the privacy of the audio recordings.
As the software industry adapts to security and privacy changes amid AI, our research team is working to make homomorphic encryption, easy to develop, fast to run, and ubiquitous across industry.
The Evolving Role of the Red Team in the Era of Agentic Security
At Google, our Red Teams have always operated on the cutting edge of security. We’ve shared our journey in the past: from the high-stakes operations showcased in our Hac…
Influence Operations Bulletin Q2 2026
This bulletin includes coordinated influence operation campaigns terminated on our platforms in Q2 2026. It was last updated on June 30, 2026.AprilWe terminated 13 YouTu…
Stronger with every update: How we’re making Chrome and the web safer in the AI Era
From Finding to Fixing: Reducing maintainer burden with automated patches
Since its launch in 2016, OSS-Fuzz has contributed significantly to making open-source secure by finding and reporting tens of thousands of bugs. But finding more vulner…
Going Beyond Zero: A New Paradigm For Enterprise Security
Bringing AI agents to Chrome Enterprise security management
More of Google
Google Products
Global (English)
Africa (English)
Australia (English)
Brasil (Português)
Canada (English)
Canada (Français)
Česko (Čeština)
Deutschland (Deutsch)
España (Español)
France (Français)
Greece (Ελληνικά)
India (English)
Indonesia (Bahasa Indonesia)
Ireland (English)
Italia (Italiano)
日本 (日本語)
대한민국 (한국어)
Latinoamérica (Español)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
MENA (English)
Nederlands (Nederland)
New Zealand (English)
Polska (Polski)
Portugal (Português)
România (Română)
Sverige (Svenska)
ประเทศไทย (ไทย)
Türkiye (Türkçe)
台灣 (中文)
