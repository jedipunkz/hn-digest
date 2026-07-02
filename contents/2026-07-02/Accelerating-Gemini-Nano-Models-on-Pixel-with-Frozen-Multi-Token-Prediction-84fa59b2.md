---
source: "https://research.google/blog/accelerating-gemini-nano-models-on-pixel-with-frozen-multi-token-prediction/"
hn_url: "https://news.ycombinator.com/item?id=48755245"
title: "Accelerating Gemini Nano Models on Pixel with Frozen Multi-Token Prediction"
article_title: "Accelerating Gemini Nano models on Pixel with frozen Multi-Token Prediction"
author: "CharlesW"
captured_at: "2026-07-02T01:40:04Z"
capture_tool: "hn-digest"
hn_id: 48755245
score: 2
comments: 0
posted_at: "2026-07-02T01:20:46Z"
tags:
  - hacker-news
  - translated
---

# Accelerating Gemini Nano Models on Pixel with Frozen Multi-Token Prediction

- HN: [48755245](https://news.ycombinator.com/item?id=48755245)
- Source: [research.google](https://research.google/blog/accelerating-gemini-nano-models-on-pixel-with-frozen-multi-token-prediction/)
- Score: 2
- Comments: 0
- Posted: 2026-07-02T01:20:46Z

## Translation

タイトル: 凍結されたマルチトークン予測を使用したピクセル上の Gemini Nano モデルの高速化
記事のタイトル: 凍結されたマルチトークン予測を使用して Pixel で Gemini Nano モデルを高速化する

記事本文:
凍結されたマルチトークン予測を使用して Pixel 上で Gemini Nano モデルを高速化する
メインコンテンツにスキップ
当社が注力している多くの分野をご覧ください
協力的なエコシステムの構築
発見を現実世界への影響に変える
当社が注力している多くの分野をご覧ください
協力的なエコシステムの構築
発見を現実世界への影響に変える
Google
研究
Google AI
当社のすべての AI について学ぶ
Googleディープマインド
AI のフロンティアを探索する
Google Labs
AI 実験を試してみる
研究
リソース
カンファレンスとイベント
キャリア
ブログ
について
検索
ホーム
凍結されたマルチトークン予測を使用して Pixel 上で Gemini Nano モデルを高速化する
Google プラットフォームおよびデバイス担当リサーチ プロダクト マネージャーの Eden Cohen 氏とリサーチ マネージャー Michelle Ramanovich 氏
マルチトークン予測を凍結された実稼働モデルに後付けする方法を導入し、個別の製図者の非効率性を排除してデバイス上の推論を高速化します。
Gemini Nano や Gemma などのオンデバイス モデルを使用すると、強力なラージ言語モデル (LLM) をポケットに入れることが現実になります。このテクノロジーにより、プライベート データをデバイスから送信することなく、大量の通知を即座に要約したり、重要なテキスト メッセージを校正したりするなど、日常的な機能を携帯電話上で実行できるようになります。ただし、これらの機能を日常のユーザーにとって役立つものにするためには、非常に効率的に実行する必要があります。
モバイル デバイスでこの種の速度を実現することは、大きな課題です。広大なサーバー環境とは異なり、携帯電話は厳しいエネルギー バジェットとハード メモリ (RAM) 制限の下で動作します。さらに、標準言語モデルはテキストを「自己回帰的に」生成します。つまり、一度に 1 つの単語 (またはトークン) だけを処理して出力します。この段階的なプロセスによりボトルネックが発生し、電話機の処理能力が十分に活用されず、メモリ帯域幅に負担がかかるため、

最終的にはユーザー エクスペリエンスが低下し、バッテリーが消耗します。
このボトルネックを克服するために、既存の「凍結された」Gemini Nano v3 モデルにマルチトークン予測 (MTP) を改良する新しいアーキテクチャを発表します。 EAGLE フレームワークや Confident Adaptive Language Modeling (CALM) などの以前のアプローチに基づいて、特にモバイル環境向けにこれらの効率向上を最大化する新しいアーキテクチャ コンポーネントを設計しました。最近の発表では、MTP を使用して Gemma 4 を高速化し、開発者が利用できるようにすることが強調されています。
今日の記事では、エッジ コンピューティング特有の極端な制約について取り上げます。最近 Pixel 9 および 10 シリーズに導入されたこのアプローチは、すぐに使える高速化として機能します。これはユーザーにとって、AI 通知概要や校正などの機能により、より少ないエネルギー消費でテキストを大幅に高速に生成できることを意味します。開発者にとっては、新しいタスクごとにメモリを大量に使用する個別の製図モデルを微調整する必要がなく、高速なオンデバイス AI が提供されるという大きな問題点が解消されます。
MTP は、投機的デコーディングの進化に基づいて構築されています。従来のセットアップでは、N 個のトークンを生成するには、大規模モデルの N 回の前方パスが必要です。投機的デコードでは、このプロセスが 2 つの部分に分離されます。
ドラフト: より小さく、より高速な近似モデル (「ドラフト」) は、候補トークンの短いシーケンス (例: 3 つのトークン) を生成します。
検証: 大規模なモデル (「検証者」) がこれらの候補を並行して処理します。候補が大規模モデルの予測と一致する場合、その候補は受け入れられます。そうでない場合、システムは最初の発散にロールバックします。
ただし、これによりいくつかの非効率が生じます。別の「スタンドアロン」ドラフター モデル (例: 128M パラメータ) を実行すると、限られた RAM をめぐって競合します。さらに、スタンドアロンの製図者はメイン モに対して「盲目」です。

del の豊富な内部状態。メイン モデルが既に計算したセマンティック コンテキストを使用せずに、テキスト履歴のみに基づいて次のトークンを予測します。 MTP は、スタンドアロン アーキテクチャから統合アーキテクチャに移行することで、これらの非効率性に対処します。トークンをドラフトするために別の小さな言語モデルをトレーニングする代わりに、軽量の Transformer ヘッドである MTP ヘッドをメイン モデルの最終層に追加します。
このアーキテクチャは、製図に深い出口層を使用し、メイン モデルのバックボーンによってすでに実行されている作業を活用します。 MTP ヘッドは、メイン モデルの最終的な高次元のアクティベーション (隠れた状態) を取得し、それらを使用して将来のトークンのシーケンスを自己回帰的に予測します。
最近のリリースの Gemma 4 モデルのように、MTP ヘッドはバックボーンと並行して事前トレーニングされるのが一般的ですが、すでにデプロイされているオンデバイス基盤モデルを活用する場合、これは法外です。代わりに、私たちの作業は、事前トレーニング パイプラインから独立して動作するようにドラフター ヘッドを改造することに重点を置いています。
完全にトレーニングされた Gemini Nano v3 モデルを使用し、その重みをフリーズし、高密度のトランスフォーマー スタック (MTP ヘッド) を最終層に接続します。将来のトークンの予測誤差を最小限に抑えるために、これらのパラメーターのみをトレーニングします。バックボーンが凍結されている場合、MTP は厳密に効率を最適化するものとなり、基本モデルの機能や安全性の調整が低下しないことが保証されます。
間違ったドラフトは検証中に破棄されるため、最終出力はメイン モデルとビットごとに同一のままであり、完全な下位互換性を備えた効率のアップデートを展開できます。
標準的な MTP 実装は、メイン モデルとドラフター間で静的パラメーター (埋め込み重みなど) を共有することでトレーニング効率を最適化しますが、オンデバイス推論は厳しい問題に直面しています。

ボトルネック: 動的メモリ。共有重みを使用する場合でも、ドラフターがコンテキストを独立して処理する場合、独自のキー/値 (KV) キャッシュを生成および維持することにより、メモリに「二重課税」が発生します。モバイルのメモリが限られているため、この冗長性を回避することが重要です。
これを解決するために、MTP ヘッドがメイン モデルの状態を効果的に活用するゼロコピー アーキテクチャを設計しました。 MTP ヘッドは、自身の履歴を維持するのではなく、メイン モデルのフリーズされた KV キャッシュに直接相互接続するように設計されています。これにより、起草者は、重複することなくバックボーンによってすでに計算された「メモリ」とコンテキストをクエリできるようになります。
この設計により、2 つの効率が向上します。まず、ドラフターの事前入力レイテンシーが排除されます。既存のキャッシュを利用することで、ヘッドはプロンプトを処理するために追加の時間を必要としません。 2 番目に、実行時のメモリ使用量が削減されます。ドラフター埋め込みルックアップ テーブル、プレフィル ドット アテンション バリアント、およびアプリケーション固有のチューニング パラメーターを節約することにより、スタンドアロン ドラフターと比較して、インスタンスごとに 130MB の節約が確認されました。
メイン モデルの隠れステートと KV キャッシュを活用することで、MTP ヘッドはバックボーンによって並行して検証される候補トークンを生成し、冗長なプレフィル レイテンシーを排除し、メモリ使用量を最大 130 MB 削減します。
より豊かな表現を可能にする
私たちの実験では、MTP ドラフターは一貫してより正確なトークン予測を生成し、その結果、同等のパラメーター数の「スタンドアロン ドラフター」と比較して、タスクに応じて Pixel 9 デバイスで 50% 以上の高速化が得られることがわかりました [aef552] 。
このパフォーマンスのギャップは、MTP がよりリッチな表現にアクセスすることに起因しています。メイン モデルをブラック ボックスとして扱うスタンドアロン ドラフターとは異なり、MTP ヘッドは、大規模なモデルによってすでに処理された最終アクティベーションを直接利用します。

えっとバックボーン:
次の指示: 複雑な制約のある要約や書き換えなどのタスクでは、MTP はスタンドアロンの微調整されたドラフターよりも大幅に優れたパフォーマンスを発揮しました。
予測可能なテキスト構造: 構造予測可能性の高いタスク (スマート リプライなど) の場合、MTP ヘッドはメイン モデルの構文パターンを効果的に学習し、トークン受け入れが最大 55% 向上しました。
Pixel 9 および 10 デバイスに MTP を導入するために、検証フェーズと製図フェーズの間の複雑な依存関係を処理できるように、デバイス上の推論スタックを再設計しました。
その結果、アーキテクチャ上の選択が検証されました。 AI 通知サマリーや校正などの実稼働ワークロードでは、MTP は推論パスごとに平均でほぼ 2 つの追加トークンを正確に予測します。さらに、検証ステップが少なくなると、重いプロセッサを起動する時間が短縮され、エネルギー消費が削減され、バッテリ寿命が向上します。
Gemini Nano トークン生成の MTP と、さまざまな Pixel 9 アプリケーションにわたるアプリ固有のスタンドアロン調整ドラフターの影響。
私たちは、将来の Pixel デバイスに MTP を統合することを楽しみにしています。また、ドラフトのレイテンシをさらに短縮し、モバイルの厳しい制約下での同時トークン検証を増やすために、補助ヘッドを使用しない並列デコードやパラダイムなどの代替アーキテクチャを検討することを楽しみにしています。
また、言語生成に固有の曖昧さをより効率的に処理する方法も調査しています。標準的な投機的デコードでは単一の最適な将来パスが想定されていますが、私たちはモデルが分岐の可能性を並行して探索できるようにする技術を開発しています。これは、不確実なコンテキストでも長いシーケンスを受け入れる確率を最大化することを目的としています。さらに、検証の寛容性を研究しています。つまり、ドラフト間の厳密な正確なトークンの一致を緩和することです。

特定のユースケースの検証と検証により、エッジの効率がさらに向上します。
この取り組みは、Filippo Galgani、Omri Homburger、Pooja Consul、Matthew Markwell、Vivek Kumar と協力して、オンデバイス LLM 効率を最適化する取り組みの一環です。特定の要素は、Google DeepMind の Gemini チーム (Tal Schuster、Ziwei ji、Ivan Korotkov、Ganesh Jawahar) の開発に基づいて構築されました。また、Nadav Bar 氏、Utku Evci 氏、Nir Shabat 氏、Joe Zou 氏、および Google Research、Google Deepmind、プラットフォームとデバイスのチームのレビュー、貴重なフィードバック、サポートに深く感謝いたします。
MTP アップデート前の Pixel 9 および Pixel 10 スマートフォンとの比較。

## Original Extract

Accelerating Gemini Nano models on Pixel with frozen Multi-Token Prediction
Skip to main content
Explore our many areas of focus
Building a collaborative ecosystem
Translating discovery into real-world impact
Explore our many areas of focus
Building a collaborative ecosystem
Translating discovery into real-world impact
Google
Research
Google AI
Learn about all our AI
Google DeepMind
Explore the frontier of AI
Google Labs
Try our AI experiments
Research
Resources
Conferences & events
Careers
Blog
About
Search
Home
Accelerating Gemini Nano models on Pixel with frozen Multi-Token Prediction
Eden Cohen, Research Product Manager, and Michelle Ramanovich, Research Manager, Google Platforms and Devices
We introduce a method to retrofit Multi-Token Prediction onto frozen production models, accelerating on-device inference without the inefficiencies of separate drafters.
Having powerful Large Language Models (LLMs) right in your pocket is now a reality with on-device models like Gemini Nano and Gemma . This technology enables everyday features on your phone — such as instantly summarizing a flurry of notifications or proofreading an important text message — all without sending your private data off device. But to make these features useful for everyday users, they need to happen very efficiently.
Delivering this kind of speed on a mobile device is a significant challenge. Unlike vast server environments, mobile phones operate under a strict energy budget and hard memory (RAM) limits. Furthermore, standard language models generate text "autoregressively" — meaning they process and output just one word (or token) at a time. This step-by-step process creates a bottleneck, underutilizing the phone's processing power while straining its memory bandwidth, which can ultimately slow down the user experience and drain the battery.
To overcome this bottleneck, we are announcing a new architecture that retrofits Multi-Token Prediction (MTP) onto existing, "frozen" Gemini Nano v3 models. Building on prior approaches like the EAGLE framework and Confident Adaptive Language Modeling (CALM), we designed new architectural components to maximize these efficiency gains specifically for mobile environments. Our recent announcements highlighted accelerating Gemma 4 with MTP and making it available to developers.
Today's article tackles the unique, extreme constraints of edge computing. Recently rolled out to the Pixel 9 and 10 series, this approach acts as an out-of-the-box speedup. For users, this means that features like AI Notification Summaries and Proofread generate text significantly faster and with less energy consumption. For developers, it eliminates a major friction point: delivering high-speed on-device AI without the need to fine-tune separate, memory-heavy drafting models for every new task.
MTP builds upon the evolution of speculative decoding . In a traditional setup, generating N tokens requires N forward passes of the large model. Speculative decoding decouples this process into two parts:
Draft: a smaller, faster approximation model (the "drafter") generates a short sequence of candidate tokens (e.g., 3 tokens).
Verify: a large model (the "verifier") processes these candidates in parallel. If the candidates match what the large model would have predicted, they are accepted. If not, the system rolls back to the first divergence.
However, this results in some inefficiencies. Running a separate "standalone" drafter model (e.g., 128M parameters) competes for limited RAM. Furthermore, a standalone drafter is "blind" to the main model's rich internal state, predicting next tokens based solely on text history without the semantic context the main model has already computed. MTP addresses these inefficiencies by moving from a standalone architecture to an integrated one. Instead of training a separate small language model to draft tokens, we append a lightweight Transformer head, the MTP head, to the final layers of the main model.
This architecture, which uses a deep exit layer for drafting, leverages the work already performed by the main model’s backbone. The MTP head takes the final high-dimensional activations (hidden states) of the main model and uses them to autoregressively predict a sequence of future tokens.
While MTP heads are commonly pre-trained in tandem with the backbone — such as in our recent releases of Gemma 4 models — this is prohibitive when leveraging already-deployed on-device foundation models. Instead, our work focuses on retrofitting the drafter head to operate independently of the pre-training pipeline.
We take a fully trained Gemini Nano v3 model, freeze its weights, and attach a dense transformer stack — the MTP head — to the final layers. We train only these parameters to minimize the prediction error on future tokens. With a frozen backbone, MTP becomes strictly an efficiency optimization, ensuring no degradation in the base model's capabilities or safety alignment.
Because incorrect drafts are discarded during verification, the final output remains bit-for-bit identical to the main model, allowing us to roll out efficiency updates with full backward compatibility.
While standard MTP implementations optimize for training efficiency by sharing static parameters (like embedding weights) between the main model and the drafter, on-device inference faces a stricter bottleneck: dynamic memory. Even with shared weights, if a drafter processes context independently, it incurs a "double tax" on memory by generating and maintaining its own key-value (KV) cache. Given the limited memory on mobile, avoiding this redundancy is critical.
To solve this, we engineered a zero-copy architecture where the MTP head effectively leverages the main model's state. Instead of maintaining its own history, the MTP head is designed to cross-attend directly to the main model’s frozen KV cache. This allows the drafter to query the "memories" and context already computed by the backbone without duplication.
This design yields two efficiency gains. First, it eliminates drafter prefill latency: by utilizing the existing cache, the head requires no additional time to process the prompt. Second, it reduces the runtime memory footprint. We observed savings of 130MB per instance compared to a standalone drafter by saving drafter embedding lookup tables, prefill dot attention variants, and application specific tuning parameters.
By leveraging the main model’s hidden states and KV cache, the MTP head generates candidate tokens that are verified in parallel by the backbone, eliminating redundant prefill latency and reducing memory usage by up to 130MB.
Unlocking richer representations
In our experiments, we found that MTP drafters consistently produce more accurate token predictions, which results in speedups on Pixel 9 devices of 50% or more [aef552] , depending on the task, compared to "standalone drafters" of comparable parameter counts.
This performance gap stems from MTP’s access to richer representations. Unlike standalone drafters that treat the main model as a black box, the MTP head directly utilizes final activations already processed by the larger backbone:
Instruction following: In tasks like summarization or rewriting with complex constraints, MTP significantly outperformed standalone fine-tuned drafters.
Predictable text structures: For tasks with high structural predictability (e.g., smart replies), the MTP head effectively learned the syntactic patterns of the main model, achieving up to a 55% improvement in token acceptance.
For the deployment of MTP on Pixel 9 and 10 devices, we redesigned the on-device inference stack to handle the complex dependency between the verification and drafting phases.
The results validated the architectural choices. In production workloads, such as AI Notification Summaries and Proofread, MTP correctly predicts an average of nearly two additional tokens per inference pass. Furthermore, fewer verification steps mean less time waking heavy processors, reducing energy consumption and improving battery life.
Gemini Nano token generation impact of MTP vs. app-specific standalone tuned drafter across various Pixel 9 applications.
We look forward to integrating MTP on future Pixel devices, as well as exploring alternative architectures — including parallel decoding and paradigms without auxiliary heads — to further drive down draft latency and increase simultaneous token verification under strict mobile constraints.
We are also investigating ways to handle the inherent ambiguity of language generation more efficiently. While standard speculative decoding assumes a single best future path, we are developing techniques that allow the model to explore branching possibilities in parallel. This aims to maximize the probability of accepting long sequences even in uncertain contexts. Furthermore, we are studying verification leniency: relaxing the strict exact token match between draft and verification for specific use cases to bring further efficiencies to the edge.
This work is part of our efforts for optimizing on-device LLM efficiency, with Filippo Galgani, Omri Homburger, Pooja Consul, Matthew Markwell, and Vivek Kumar. Certain elements were built on developments from the Gemini team in Google DeepMind: Tal Schuster, Ziwei ji, Ivan Korotkov, and Ganesh Jawahar. We’d also like to extend a big thank you for reviews and valuable feedback and support to Nadav Bar, Utku Evci, Nir Shabat, Joe Zou, and teams in Google Research, Google Deepmind, and Platforms & Devices.
Compared to Pixel 9 and Pixel 10 phones prior to MTP update.
