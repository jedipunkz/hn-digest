---
source: "https://netflixtechblog.com/genrec-towards-llm-native-recommendation-at-netflix-f20be6f643e3"
hn_url: "https://news.ycombinator.com/item?id=49146751"
title: "GenRec: Towards LLM-Native Recommendation at Netflix"
article_title: "Medium"
author: "Uriopass"
captured_at: "2026-08-02T18:55:57Z"
capture_tool: "hn-digest"
hn_id: 49146751
score: 2
comments: 1
posted_at: "2026-08-02T18:00:10Z"
tags:
  - hacker-news
  - translated
---

# GenRec: Towards LLM-Native Recommendation at Netflix

- HN: [49146751](https://news.ycombinator.com/item?id=49146751)
- Source: [netflixtechblog.com](https://netflixtechblog.com/genrec-towards-llm-native-recommendation-at-netflix-f20be6f643e3)
- Score: 2
- Comments: 1
- Posted: 2026-08-02T18:00:10Z

## Translation

タイトル: GenRec: Netflix での LLM ネイティブの推奨に向けて
記事タイトル: 中
概要: GenRec: Netflix における LLM ネイティブ レコメンデーションに向けて 著者: Ying Li、Arjun Rao、Shradha Sehgal はじめに レコメンデーションは Netflix エクスペリエンスの中心にあります。現在の量産モデル…

記事本文:
Medium GenRec: Netflix での LLM ネイティブの推奨に向けて | by Netflix テクノロジーブログ | 2026 年 7 月 | Netflix TechBlog サイトマップ アプリで開く サインアップ
Foundation LLM から Recommendation Ranker フェーズ 1 へ — Netflix に適応した Foundation LLM。
会話としてのトレーニング データ
言語化とコンテキストエンジニアリング
目的: ランキング、言語、報酬 1. カタログを意識したランキングの目的
2. 言語モデリングの目的
3. アライメントによる報酬加重損失
モデルのアーキテクチャとサービスのバックボーンとスコアリングヘッド
オフラインとオンラインの実験 GenRec と運用ベースライン
データ、モデル、フェーズの貢献
LLM‑Native の推奨に向けて 特徴エンジニアリングからコンテキスト エンジニアリングまで
カスタマイズされたアーキテクチャから基盤バックボーンまで
RecSys インフラから LLM インフラへ
Netflix の世界クラスのエンジニアリングの取り組み、企業文化、製品開発などについて学びましょう。
レコメンデーション システム Genai Netflix 大規模言語モデル GenRec: Netflix での LLM ネイティブ レコメンデーションに向けて
著者: イン・リー、アルジュン・ラオ、シュラダ・セーガル
おすすめは Netflix エクスペリエンスの中心にあります。当社の現在の運用モデルは、シーケンス モデリング、機能インタラクション、マルチタスク目標に特化したアーキテクチャとともに、ユーザー、アイテム、インタラクションに関する何千もの手作りの機能に依存しています。このスタックは、多様なコンテンツ タイプ (映画、シリーズ、ゲーム、ライブ、ポッドキャスト) と製品サーフェスをサポートするために長年にわたって進化してきましたが、その複雑さにより、新しいユースケースを導入するのにコストがかかります。コンテンツ タイプまたはサーフェスの追加には、大幅な機能エンジニアリング、アーキテクチャの変更、インフラストラクチャの作業、および実験が必要になる場合があります。
同時に、大規模言語モデル (LLM) は、最近の研究で示されているように、レコメンデーションについての考え方を変えています。

PLUM 、 GLIDE 、 OneRec-Think など。彼らの幅広い世界知識と強力な言語理解により、ユーザー履歴とアイテムのメタデータをテキストとして直接表現し、共有された意味論的空間で豊かな関係をキャプチャし、自然言語プロンプトを介して推奨事項を導くことが可能になります。しかし、既製の LLM は、本番環境に対応したレコメンデーションとは程遠いものです。多くの場合、世界的に人気のあるコンテンツを過剰に推奨し、カタログ外の商品を幻覚させ、ビジネス上の制約を無視し、限られたパーソナライゼーションしか提供しません。
これに対処するために、Netflix 固有のデータと目標に基づいて内部基盤 LLM を事後トレーニングする、LLM 支援の推奨ランカーである GenRec を構築しました。 GenRec は、LLM ベースのランカーが、はるかに少ないラベル付きサンプルと入力信号に依存しながら、成熟した運用システムと同等またはそれを超えることができることを示しています。
図 1: GenRec パイプライン。ユーザー履歴、アイテムのメタデータ、およびコンテキストの生のログは、コンテキスト エンジニアリングによって自然言語プロンプトに変換され、GenRec に供給されます。GenRec は、プレフィル専用モードで vLLM 上で実行され、各カタログ アイテムのスコアを出力して、推奨ランキングを生成します。大まかに言うと、GenRec は次のことを行います。
ユーザー履歴、アイテムのメタデータ、コンテキストをテキストとして言語化します。
Netflix に適応した基盤 LLM をランキング用にポストトレーニングします。
Netflix タイトルにカタログ対応のスコアリング ヘッドを追加します。
長期的な会員価値とビジネス目標に合わせて報酬シグナルを使用します。
コスト効率を高めるため、Netflix の LLM サービススタック上で事前入力専用モードで実行されます。
GenRec は、よく調整された製品ランカーに対する大規模な A/B テストで、フェーズ 2 のラベル付けされたデータと入力信号のごく一部のみを使用しながら、短期および長期のオンライン指標の両方で統計的に有意な改善を達成しました。それは私たちの依存を減らします

手作業で機能をエンジニアリングし、焦点を特徴エンジニアリングからコンテキスト エンジニアリングに移します。このブログ投稿では、GenRec がどのように機能し、どのように実行されるか、そして Netflix での推奨において、GenRec がより LLM 中心の未来を指し示していると考える理由について説明します。
私たちは、完全なカタログのランキング タスク (候補セットが提供されている場合は上位 K ランキング) に焦点を当てます。
ユーザー 𝑢、ユーザーのインタラクション履歴 𝐻、および現在のコンテキスト 𝜏 (デバイス、サーフェス、ロケール、時間など) を考慮して、GenRec は各項目にスコアを付け、レコメンデーションを直接強化したり、下流のパーソナライゼーション システムへの入力として機能するパーソナライズされたランキングを生成します。
形式的には、リクエスト ( u 、 τ 、 t 、 H ) (ユーザー、コンテキスト、時間、履歴) をカタログ C 上のランキング 𝜋 にマップします。ここで、 π(i) は項目 i に割り当てられた位置です。私たちは、短期的なエンゲージメントだけでなく、期待される長期的なメンバーの有用性（満足度と定着率の代用）のために π を最適化します。
Foundation LLM から推奨ランカーへ
GenRec は 2 段階のトレーニング フレームワークに従います (図 2)。
図 2: 2 フェーズのフレームワーク。フェーズ 1 では、ユーザーとコンテンツを理解するために Netflix データで基礎的な LLM をトレーニングし、フェーズ 2 では、ランキング固有のデータと目標についてポストトレーニングを行います。フェーズ 1 — Netflix に適応した Foundation LLM 。
私たちはオープンソース LLM からスタートし、それを独自の Netflix コーパスに適応させて、次のような基本的な機能を学習します。
メンバーの行動と好みのパターン
一般的な言語の理解と生成。
フェーズ 1 は比較的頻繁に更新されず、多くのアプリケーションの共有の Netflix 対応バックボーンとして機能します。
次に、ランキング固有のデータと目標に関する事後トレーニングによって、この基礎モデルを高品質のランキング モデルに変えます。フェーズ 2:
ランキングの品質とステアリングに重点を置く
組み込む

報酬加重損失による複数の報酬シグナル
新しいコンテンツと進化する好みを追跡するために、より頻繁に更新されます
サービスコストの制約の下で明示的に最適化されます。
会話としてのトレーニング データ
Netflix メンバーは、さまざまな面 (視聴、再生、再生時間、サムアップ/ダウン、リストへの追加、放棄など) にまたがる何千億ものインタラクション イベントを生成します。このログ データは、ユーザーと推奨者間の 1 ターンまたは複数ターンの「会話」に変換されます。各ターンには次の内容が含まれます。
ユーザーメッセージ: 言語化されたコンテキスト、プロフィール、履歴、アイテムのメタデータ、およびタスク (例: ユーザーが次に視聴するものを推奨したり、親指を立てたりする)。
アシスタント メッセージ: メンバーの実際のエンゲージメント (例: どのタイトルがどのくらいの時間再生され、どのようなフィードバックが提供されたか)。
フェーズ 2 のトレーニング中に、LLM はアシスタント メッセージがユーザー メッセージにどのように依存するかを学習します。これにより、豊富な推奨シグナルをテキストとして表現できるようになり、言語モデリング (LM) とランキング目標の両方を共同でサポートできます。
推論時には、言語化されたコンテキストをフィードし、カタログ対応のスコアリング ヘッドを適用して項目をランク付けします。アシスタント メッセージはデコードしません。会話形式は主にトレーニング中に LM の目的をサポートし、言語化されたテキストに対する強力な言語理解を維持するために使用されます。
言語化とコンテキストエンジニアリング
従来のレコメンダーは、高密度の機能と埋め込みを操作します。 GenRec は異なるアプローチを採用しています。豊富なユーザー履歴とコンテキストを自然言語として言語化し、生のインタラクション信号を LLM の意味空間で直接エンコードします。その際、手動の特徴エンジニアリングではなく、モデルに依存して、アイテムの関係性や進化するユーザーの関心などのより高いレベルのパターンを発見します。
あらゆるインタラクションを素朴に言語化する

ユーザーの履歴にあるものはすぐにトークンの予算を超え、Netflix の規模では高すぎる可能性があります。コンテキスト ウィンドウが新しい「機能バジェット」になるため、コンテキスト エンジニアリングを適用します。
完全に保持: シグナルの高いエンゲージメント (長時間のプレイ、サムズアップなど) をより詳細に保持
省略: 低信号イベント (非常に短いプレイや素早いホバリングなど)
要約または圧縮: 反復的な行動 (例: むしゃぶり見)
選択的に詳しく説明します: 重要な項目またはコールドスタート項目 (例: 新しいリリース)
固定されたトークン予算内で、最近の高信号履歴を優先し、古い履歴を圧縮または削除します。また、プレフィックス キャッシュを改善するために、共有プレフィックスを最大化するようにプロンプ​​トを構成します。目標は、法外なコストをかけずにランキングの品質を維持する、コンパクトで情報量の多いプロンプトです。
目的: ランキング、言語、報酬
GenRec モデル全体は、推奨ランキングの目標、言語モデリングの目標、報酬加重トレーニングによる調整を組み合わせた多目的損失でトレーニングされます。
1. カタログを意識したランキングの目的
主なタスクは、エンゲージメントの質によって項目をスコアリングするようモデルに教えるランキング目標です。しきい値とノイズ除去ロジックを使用して、価値の高いエンゲージメント (十分に長い再生、強力な明示的なフィードバックなど) を使用してポジティブな要素にラベルを付け、言語化されたコンテキストを考慮してこれらのポジティブな要素に高いスコアを割り当てるように、カタログまたは候補セットにわたるクロスエントロピー損失を介してモデルをトレーニングします。
2. 言語モデリングの目的
また、言語化された入力と出力に対する言語モデリング (LM) の目標も保持しています。これにより、モデルの一般的な言語理解が維持され、豊富な自然言語履歴とアイテムのメタデータを解釈する能力が向上し、テキスト生成の用途への扉が開かれたままになります。

おすすめの説明など。
3. アライメントによる報酬加重損失
GenRec は、生のランキング精度を超えて、(1) ビジネス要件 (映画、シリーズ、ゲーム、ライブ、ポッドキャストのバランスなど) を尊重し、(2) 単なる即時のクリックや再生ではなく、長期的なメンバーの満足度を最適化する必要があります。
Netflix テクノロジーブログの記事を受信トレイで入手
無料で Medium に参加して、このライターから最新情報を入手してください。
より速くサインインするために私を覚えておいてください
生のインタラクション シーケンスのみをトレーニングすると、むちゃくちゃ視聴することを好みすぎたり、単一のコンテンツ タイプに過度に集中したりするなど、望ましくない動作が発生する可能性があります。これに対処するために、別の報酬モデルからのシグナルを使用してランキングの損失に重みを付けます。各トレーニング サンプルは、次の 2 種類の信号から導出されたスカラー重みを受け取ります。
長期的な満足度の指標: 短期的なエンゲージメントが、返品行動、カタログの探索、持続的なエンゲージメントなどの長期的な成果にどの程度寄与するかを推定します。
動作の再バランス: コンテンツ タイプとリリース段階 (たとえば、ゲームと映画、新しいリリースと常緑のタイトル) 全体で動作を調整し、ビジネス目標との整合性を高めます。
この例のランキングの損失は、この重みによって調整されます。つまり、価値の高いエンゲージメントにはより大きな重みが与えられ、価値の低いエンゲージメントには重みが低くなります。この報酬重み付けアプローチは、完全な強化学習よりもシンプルでコスト効率が高く、さらに実際に効果的な調整を実現します。 RL スタイルの手法 (GRPO など) からさらなる利益が得られることがわかりましたが、コストが高いため、将来の研究に委ねられます。
モデルのアーキテクチャとサービス提供
GenRec のアーキテクチャは、当社の基本的な LLM に厳密に従っています。つまり、ネクスト トークン予測スタイルの目標でトレーニングされたデコーダー専用の Transformer であり、カタログ対応の RLM で強化されています。

Netflix のカタログ内アイテムのみを採点するアンキング ヘッド。スコアリング パイプラインは次のように機能します。
言語化: 言語化ツール V は、ユーザー履歴 H 、コンテキスト 𝜏 、および関連するアイテムのメタデータを単一のテキスト シーケンス x にシリアル化します。
プールされた表現: LLM は x を処理し、ユーザーの現在の設定とコンテキストを要約したプールされた隠れ状態 h を抽出します。
カタログ認識スコアリング: 各カタログ項目 i には学習された埋め込み eᵢ があります。スコアリング ヘッド ϕ は、h と eᵢ を (たとえば、ドット積または小さな MLP を介して) 組み合わせて、スコア s ᵢ を生成します。スコアにソフトマックスを適用すると確率分布が得られ、これをランキング π に変換します。
すべてのパラメーター (バックボーン、スコアリング ヘッド、アイテムの埋め込み) は、共同してトレーニングされます。非常に大規模なカタログの場合、サンプリングされたソフトマックスまたは候補セットを使用して、効率的なトレーニングと推論を行うことができます。このアーキテクチャは、大規模な候補セットに対する効率的なスコアリングをサポートしながら、推奨事項を Netflix カタログに制限します。
GenRec は、vLLM を使用して Netflix の内部 LLM スタックで提供されます。 Netflix の規模では、サービスコストは主に 1) モデルのサイズによって決まります。 2) コンテキストの長さ。 3) 推論モード (プレフィル vs. 自己回帰デコード)。私たちは次の 3 つの戦略を通じてコストを管理します。
より小規模な/抽出されたモデル: より低いサービスコストでより大きなモデルの品質のほとんどを取得するために、多くの場合、より大規模な、またはよりターゲットを絞ったデータセットを使用して、小規模または抽出された基礎モデルで GenRec をトレーニングします。
アグレッシブc

[切り捨てられた]

## Original Extract

GenRec: Towards LLM-Native Recommendation at Netflix Authors: Ying Li, Arjun Rao, Shradha Sehgal Introduction Recommendations sit at the heart of the Netflix experience. Our current production models …

Medium GenRec: Towards LLM-Native Recommendation at Netflix | by Netflix Technology Blog | Jul, 2026 | Netflix TechBlog Sitemap Open in app Sign up
From Foundation LLM to Recommendation Ranker Phase 1 — Netflix-Adapted Foundation LLM.
Training Data as Conversations
Verbalization and Context Engineering
Objectives: Ranking, Language, and Rewards 1. Catalog‑Aware Ranking Objective
2. Language Modeling Objective
3. Reward‑Weighted Loss for Alignment
Model Architecture and Serving Backbone and Scoring Head
Offline and Online Experiments GenRec vs Production Baseline
Data, Model, and Phase Contributions
Towards LLM‑Native Recommendation From Feature Engineering to Context Engineering
From Customized Architectures to Foundation Backbones
From RecSys Infra to LLM Infra
Learn about Netflix’s world class engineering efforts, company culture, product developments and more.
Recommendation System Genai Netflix Large Language Models GenRec: Towards LLM-Native Recommendation at Netflix
Authors: Ying Li , Arjun Rao , Shradha Sehgal
Recommendations sit at the heart of the Netflix experience. Our current production models rely on thousands of hand‑crafted features over users, items, and interactions, along with specialized architectures for sequence modeling, feature interactions, and multi‑task objectives. This stack has evolved over many years to support diverse content types (movies, series, games, live, podcasts) and product surfaces, but its complexity makes it costly to onboard new use cases: adding a content type or surface can require significant feature engineering, architecture change, infrastructure work, and experimentation.
At the same time, large language models (LLMs) are changing how we think about recommendation, as shown by recent work such as PLUM , GLIDE , and OneRec-Think . Their broad world knowledge and strong language understanding make it possible to represent user histories and item metadata directly as text, capture rich relationships in a shared semantic space, and steer recommendations via natural‑language prompts. However, off‑the‑shelf LLMs are still far from production‑ready recommenders: they often over‑recommend globally popular content, hallucinate out‑of‑catalog items, ignore business constraints, and provide only limited personalization.
To address this, we built GenRec , an LLM‑backed recommendation ranker that post‑trains an internal foundation LLM on Netflix‑specific data and objectives. GenRec shows that an LLM‑based ranker can match or exceed a mature production system while relying on far fewer labeled examples and input signals.
Figure 1: GenRec pipeline. Raw logs of user history, item metadata, and context are transformed via context engineering into natural-language prompts and fed into the GenRec, which runs on vLLM in prefill-only mode and outputs scores for each catalog item, yielding a recommendation ranking. At a high level, GenRec:
Verbalizes user histories, item metadata, and context as text.
Post‑trains a Netflix‑adapted foundation LLM for ranking.
Adds a catalog‑aware scoring head over Netflix titles.
Uses reward signals to align with long‑term member value and business goals.
Runs in prefill‑only mode on Netflix’s LLM serving stack for cost efficiency.
In a large‑scale A/B test against a well‑tuned production ranker, GenRec achieves statistically significant improvements in both short‑term and long‑term online metrics, while using only a small fraction of the Phase‑2 labeled data and input signals. It reduces our reliance on hand‑engineered features and shifts the focus from feature engineering to context engineering . In this blog post, we will describe how GenRec works, how it performs, and why we believe it points toward a more LLM‑centric future for recommendation at Netflix.
We focus on a full‑catalog ranking task (or top‑ K ranking when a candidate set is provided).
Given a user 𝑢, their interaction history 𝐻, and the current context 𝜏 (device, surface, locale, time, etc.), GenRec scores each item and produces a personalized ranking that can directly power recommendations or serve as input for downstream personalization systems.
Formally, we map a request ( u , τ , t , H ) — user, context, time, and history — to a ranking 𝜋 over the catalog C , where π(i) is the position assigned to item i . We optimize π for expected long‑term member utility (a proxy for satisfaction and retention), not just short‑term engagements.
From Foundation LLM to Recommendation Ranker
GenRec follows a two‑phase training framework (Figure 2):
Figure 2: Two Phase Framework. Phase 1 trains a foundational LLM on Netflix data for user and content understanding, and Phase 2 post-trains on ranking-specific data and objectives. Phase 1 — Netflix-Adapted Foundation LLM .
We start from an open‑source LLM and adapt it on proprietary Netflix corpora, so it learns foundational capabilities such as
Member behavior and preference patterns
General language understanding and generation.
Phase 1 is updated relatively infrequently and serves as a shared, Netflix‑aware backbone for many applications.
We then turn this foundation model into a high‑quality ranking model by post‑training on ranking‑specific data and objectives. Phase 2:
Focuses on ranking quality and steering
Incorporates multiple reward signals via reward‑weighted losses
Is refreshed more frequently to track new content and evolving tastes
Is explicitly optimized under serving cost constraints.
Training Data as Conversations
Netflix members generate hundreds of billions of interaction events spanning many surfaces (views, plays, durations, thumbs up/down, add to list, abandons, etc.). We convert this log data into single‑turn or multi‑turn “conversations” between a user and a recommender. Each turn contains:
User message : verbalized context, profile, history, item metadata, and task (e.g., recommend what the user will watch or thumb next).
Assistant message : the member’s actual engagement (e.g., which titles were played, for how long, what feedback they provided).
During Phase‑2 training, the LLM learns how assistant messages depend on user messages. This allows us to express rich recommendation signals as text, jointly supporting both the language-modeling (LM) and ranking objectives.
At inference time, we feed in the verbalized context and apply a catalog‑aware scoring head to rank items; we do not decode assistant messages. The conversational format is primarily used during training to support the LM objective and preserve strong language understanding over the verbalized text.
Verbalization and Context Engineering
Traditional recommenders operate on dense features and embeddings. GenRec takes a different approach: it verbalizes rich user histories and context as natural language, encoding raw interaction signals directly in the LLM’s semantic space. In doing so, it relies on the model to discover higher‑level patterns — such as item relationships and evolving user interests — rather than on manual feature engineering.
Naively verbalizing every interaction in a user’s history can quickly exceed the token budget and be too expensive at Netflix scale. The context window becomes our new “feature budget”, so we apply context engineering :
Retain in full: high‑signal engagements (e.g., long plays, thumbs‑up) with richer details
Omit: low‑signal events (e.g., very short plays or quick hovers)
Summarize or compress: repetitive behaviors (e.g., binge‑watching )
Elaborate selectively: important or cold‑start items (e.g., new releases)
Within a fixed token budget, we prioritize recent, high‑signal history and compress or drop older history. We also structure the prompt to maximize shared prefixes for better prefix caching. The goal is a compact, high‑information prompt that preserves ranking quality without prohibitive costs.
Objectives: Ranking, Language, and Rewards
The overall GenRec model is trained with a multi‑objective loss that combines a recommendation ranking objective, language modeling objectives, and alignment via reward‑weighted training.
1. Catalog‑Aware Ranking Objective
The primary task is a ranking objective that teaches the model to score items by engagement quality. We label positives using high‑value engagements (e.g., sufficiently long plays, strong explicit feedback), with thresholds and denoising logic, and train the model — via a cross‑entropy loss over the catalog or candidate set — to assign higher scores to these positives given a verbalized context.
2. Language Modeling Objective
We also retain a language modeling (LM) objective over the verbalized inputs and outputs. This helps preserve the model’s general language understanding, improves its ability to interpret rich natural‑language histories and item metadata, and keeps the door open for text‑generation use cases such as recommendation explanations.
3. Reward‑Weighted Loss for Alignment
Beyond raw ranking accuracy, GenRec must (1) respect business requirements — for example, balancing movies, series, games, live, and podcasts — and (2) optimize long‑term member satisfaction rather than just immediate clicks or plays.
Get Netflix Technology Blog ’s stories in your inbox
Join Medium for free to get updates from this writer.
Remember me for faster sign in
Training only on raw interaction sequences can lead to undesirable behaviors, such as over‑favoring binge‑watching or over‑focusing on a single content type. To address this, we weight the ranking loss using signals from separate reward models . Each training example receives a scalar weight derived from two types of signals:
Long‑term satisfaction proxies: estimate how much a short‑term engagement contributes to long‑term outcomes, such as return behavior, catalog exploration, or sustained engagement.
Behavior rebalancing: adjust behaviors across content types and launch stages (for example, games vs. movies, new releases vs. evergreen titles) to better align with business goals.
The example’s ranking loss is then scaled by this weight: high‑value engagements receive larger weights, and low‑value ones are down‑weighted. This reward‑weighted approach is simpler and more cost-efficient than full reinforcement learning, yet provides effective alignment in practice. We have seen additional gains from RL‑style methods (e.g., GRPO), but leave them to future work due to their higher cost.
Model Architecture and Serving
GenRec’s architecture closely follows our foundational LLM: a decoder‑only Transformer trained with next‑token‑prediction style objectives, augmented with a catalog‑aware ranking head that scores only Netflix in-catalog items. The scoring pipeline works as follows:
Verbalization: A verbalizer V serializes user history H , context 𝜏 , and relevant item metadata into a single text sequence x .
Pooled representation: The LLM processes x, and we extract a pooled hidden state h that summarizes the user’s current preferences and context.
Catalog‑aware scoring: Each catalog item i has a learned embedding eᵢ . A scoring head ϕ combines h and eᵢ (e.g., via dot product or small MLP) to produce a score s ᵢ . Applying a softmax over scores yields a probability distribution which we convert into a ranking π .
All parameters — the backbone, scoring head, and item embeddings — are trained jointly. For very large catalogs, we can use sampled softmax or candidate sets for efficient training and inference. This architecture constrains recommendations to the Netflix catalog while supporting efficient scoring over large candidate sets.
GenRec is served on Netflix’s internal LLM stack using vLLM. At Netflix scale, serving cost is driven primarily by 1) Model size; 2) Context length; 3) Inference mode (prefill vs. autoregressive decoding). We control cost through three strategies:
Smaller / distilled models: We train GenRec on smaller or distilled foundation models, often with larger or more targeted datasets, to capture most of the quality of larger models at lower serving cost.
Aggressive c

[truncated]
