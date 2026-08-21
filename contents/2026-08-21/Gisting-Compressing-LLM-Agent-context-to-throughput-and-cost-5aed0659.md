---
source: "https://shopify.engineering/gisting"
hn_url: "https://news.ycombinator.com/item?id=49382646"
title: "Gisting: Compressing LLM Agent context to ↑ throughput and ↓ cost"
article_title: "Gisting: Compressing LLM Agent context to ↑ throughput and ↓ cost (2026) - Shopify"
image: "https://cdn.shopify.com/b/shopify-brochure2-assets/7234876f98a20dbbe33d350d1e7ac41d.png"
author: "mellosouls"
captured_at: "2026-08-21T02:18:12Z"
capture_tool: "hn-digest"
hn_id: 49382646
score: 1
comments: 0
posted_at: "2026-08-21T01:42:09Z"
tags:
  - hacker-news
  - translated
---

# Gisting: Compressing LLM Agent context to ↑ throughput and ↓ cost

- HN: [49382646](https://news.ycombinator.com/item?id=49382646)
- Source: [shopify.engineering](https://shopify.engineering/gisting)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T01:42:09Z

## Translation

タイトル: 要旨: LLM エージェントのコンテキストを圧縮してスループットを↑、コストを↓にする
記事のタイトル: 概要: LLM エージェントのコンテキストを圧縮してスループットを↑、コストを↓にする (2026) - Shopify
説明: Gisting はコンテキストを学習済みトークンのセットに圧縮し、その品質を維持しながらモデルを高速かつ安価にします。

記事本文:
コンテンツにスキップ ソリューション 開始 ビジネスを開始します。ブランドを構築する
ウェブサイトを作成します。オンラインストア編集者
ストアをカスタマイズします。ストアのテーマ
ビジネス アプリを検索します。 Shopify アプリストア
自分のサイトのドメインを所有します。ドメインとホスティング
無料のビジネス ツールを探索してください。ビジネスを運営するためのツール
販売 製品を販売します。オンラインまたは対面で販売する
顧客をチェックしてください。ワールドクラスのチェックアウト
オンラインで販売します。オンラインでビジネスを成長させる
さまざまなチャネルで販売します。何百万人もの買い物客にリーチし、売上を伸ばす
世界中に販売します。海外販売
卸売＆直接販売。企業間 (B2B)
マーケット あなたのビジネスをマーケティングします。顧客にリーチして維持する
ソーシャル全体の市場。ソーシャルメディアの統合
顧客とのチャット。 Shopify 受信箱
顧客を育成する。 Shopify メッセージング
聴衆を知る。顧客の洞察を得る
管理 ビジネスを管理します。売上、注文、分析を追跡する
パフォーマンスを測定します。分析とレポート作成
在庫と注文を管理します。在庫と注文の管理
ビジネスを自動化します。 Shopify フロー
Shopify 開発者。 Shopify の強力な API を使用して構築する
プラス。成長するデジタル ブランドのためのコマース ソリューション
すべての製品。 Shopify のすべての製品と機能を詳しく見る
リソース ヘルプとサポート ヘルプとサポート 。 24時間年中無休のサポートを受ける
ビジネスコース。実績のある専門家から学ぶ
人気のトピック Shopify とは何ですか? 。当社のコマースプラットフォームの仕組み
必須ツール ビジネス名ジェネレーター。
新機能変更ログ。最近の更新情報のソース
ニュースルーム 。すべての企業ニュースとプレスリリース
ブログ | AI と機械学習のギスティング: LLM エージェントのコンテキストを圧縮してスループットを↑、コストを↓にする
Gisting はコンテキストを学習済みトークンのセットに圧縮し、その品質を維持しながらモデルを高速かつ安価にします。
システム プロンプトは、リクエストごとに数千のトークンを処理する可能性があります。長い p

rompt は、より遅く、より高価な推論を意味します。その結果、専用ハードウェアでモデルを提供する場合、同じトラフィックに対応するためにより多くの GPU が必要になります。
「言語モデルにおける制御性と毒性低減のためのプロンプト圧縮と対照的条件付け」で最初に提案された gisting の実装により、短いプロンプトを犠牲にして長いプロンプトの動作上の利点を得ることができます。
Gisting を通じて、予測の品質を損なうことなく、Sidekick GraphQL エージェントのシステム プロンプトを約 6,000 トークンから約 1,500 の Gist トークンに削減しました (4:1 の削減)。これは、知識の蒸留によって特殊なトークンのセットの埋め込みを学習し、推論時にシステム プロンプトと交換することで実現します。
システム プロンプトを Gist トークンに 4:1 圧縮することで、サービスが大幅に向上します。 1 分あたり 350 リクエスト (RPM) では、最初のトークンまでの時間 (TTFT) の中央値は 438 ミリ秒から 354 ミリ秒に低下し、エンドツーエンドのリクエスト遅延の中央値は 6.8 秒から 4.2 秒に低下し、スループットは 1 秒あたり 20.2 クエリから 23.4 クエリ (QPS) に増加しました。これらの利点により、GraphQL エージェントのトラフィックに割り当てられる GPU の数を減らすことができました。
要点トークン n は、モデルの語彙に追加する特別なトークンです。埋め込みのシーケンスをトレーニングして、コンテキストへの埋め込みの置換により、モデルが完全なプロンプトを見たかのように動作するようにします。 4:1 圧縮では、4 つのプロンプト トークンごとに 1 つの gist トークンが追加されます。モデルの重みを凍結し、要点の埋め込みのみをトレーニングします。
私たちは知識の蒸留を通じて要点の埋め込みを学習します。各軌道で前方パスを 2 回実行します。教師のパス中、モデルは完全な自然言語プロンプトを参照して、モデルの応答内の各位置に対する教師のロジットを導き出します。学生証の場合はw

e 完全なプロンプトを要点トークンと交換し、同じモデルを再度実行して、教師の各応答位置に対応する生徒のロジットを導き出します。生徒の予測が教師の予測とほぼ一致するまで、教師のロジットと生徒のロジットの間の KL 発散を使用して要点の埋め込みをトレーニングします。
要点モデルのデプロイは簡単です。トレーニングが終了したら、Gist の埋め込みをモデルの埋め込み行列に直接書き込み、新しい Gist トークンを特別なトークンとしてモデルのトークナイザーに登録します。モデルは、推論時に他のモデルと同様にロードおよび実行されます。カスタム アテンション マスク、追加のエンコーダー、特別なサービング パスはありません。唯一の推論時の変更はリクエスト側であり、プロンプトを gist トークンの文字列に置き換えます。圧縮のコスト全体はトレーニング時に 1 回だけ支払われます。
プレフィックス キャッシュだけでは不十分な理由
gisting とプレフィックス キャッシュの利点は相互に排他的ではありません。すべての最新のサービング エンジンは、以前に確認されたシーケンスのキーと値を保存する KV キャッシュを維持します。新しいリクエストにキャッシュ内に存在するシーケンス (通常はシステム プロンプトを含む) が含まれる場合、そのシーケンスの KV テンソルは再計算されずにフェッチされます。
プレフィックス キャッシュは強力ですが、デコード コストが排除されるわけではありません。モデルがトークンを生成するたびに、そのトークンはキャッシュされているかどうかに関係なく、シーケンス内のすべてのキーを処理します。デコードはメモリ帯域幅に制限されているため、生成されたすべてのトークンは高帯域幅メモリから KV キャッシュ全体をストリーミングする必要があり、その読み取りはキャッシュされたシーケンスの長さに比例して増加します。 Gisting は、アテンション計算と KV キャッシュ読み取りのコストを削減します。後者は、バッチ サイズが大きくなった場合にスループットに特に影響を与えます。
gisting および prefix caching コンパウンドの最適化、および w

実験と本番環境の両方でスタックを使用します。
ハイパーパラメータを調整するために、トレーナーに自動調査ループを向けました。レシピを提案し、Gist の埋め込みをトレーニングし、結果のモデルを評価し、それを繰り返します。
自動調査ループ中に、特に 3 つの最適化が大きな影響を与えました。
初期化 : ランダム ノイズで埋め込みを初期化する代わりに、システム プロンプトを長さ k のシーケンスに分割し (k は k :1 圧縮率から導出されます)、n 番目の gist 埋め込みを n 番目のシステム プロンプト チャンクの平均で初期化します。この最適化により、初期損失が 7 分の 1 に減少しました。
圧縮 : さまざまな圧縮率を試した結果、それを超えると予測品質が低下し始める比率がわかりました (ドメインの複雑さに対して最適な比率は 4:1 でした。他のドメインは異なります)。
データ量と多様性 : 大規模で多様なデータセットをキュレーションすることで、残りのギャップが解消されました。
自動リサーチは、ハイパーパラメータ調整の合理化に加えて、トレーニング インフラストラクチャにいくつかの重要な最適化を実装するのにも役立ちました。 1 つ目は損失を正規化する方法でした。応答トークンごとに平均化するとモデルが幻覚を起こす一方で、バッチ全体で平均化すると長い応答からの信号がより多く保存され、安定した埋め込みが生成されました。 2 つ目は速度です。教師のロジットを事前計算し、データを事前にトークン化することで、フル実行時間が 30 時間から 6 時間に短縮されました。
成果: レイテンシーとスループット
圧縮プロンプトを使用した同じモデルと完全プロンプトを比較する負荷テストを実行しました。
ギスティングにより、完全なプロンプトに比べてレイテンシーが大幅に増加しますが、リクエストの同時実行数が増加し、バッチが大きくなるにつれて、その差はさらに顕著になります。 350 RPM では、TTFT が 19% 低下し、E2E 遅延が約 38% 低下し、スループットが向上しました。

16％でした。
実際には、スループットの 16% の向上は、実稼働ワークロードの GPU 節約に直接つながりました。 GraphQL トラフィックを処理するハードウェア構成により、Gisting で使用する GPU を 14% 減らすことができました。
説明と継続的な学習
Gisting は、Shopify の継続的な学習ループにもうまく組み込まれています。モデルの抽出された Gist 埋め込みを取得したら、そのモデルを継続的な学習の新しい開始点として扱うことができます。 Gist エンベディングをプレフィックスとして使用してポストトレーニングを行い、モデルの重みと Gist エンベディングの両方に勾配更新を適用できます。増分データの重みと要点埋め込みの両方を最適化することで、毎回要点埋め込みを最初から抽出するという計算負荷なしで、モデルを継続的に調整および改善できます。
長いシステム プロンプトは便利ですが、コストがかかります。 Gisting を使用すると、エージェントは少数のトークンで広範な命令の利点をすべて活用できるようになり、待ち時間と GPU の使用量が削減されます。 GPU の需要が供給を上回る最近の時代では、あらゆる推論の最適化が重要になりますが、品質に関して妥協することも拒否します。 Gisting はその両方を提供し、Shopify ではそれが新しい標準です。
1. ウィンゲート、シューエイビ、ソレンセン。言語モデルの制御性と毒性軽減のための即時圧縮と対照的条件付け。 2022年。
当社の募集職種をご覧になり、当社のデジタル・バイ・デザイン文化について詳しく学んでください。

## Original Extract

Gisting compresses context into a set of learned tokens, preserving its quality while making the model faster and cheaper.

Skip to Content Solutions Start Start your business . Build your brand
Create your website . Online store editor
Customize your store . Store themes
Find business apps . Shopify app store
Own your site domain . Domains & hosting
Explore free business tools . Tools to run your business
Sell Sell your products . Sell online or in person
Check out customers . World-class checkout
Sell online . Grow your business online
Sell across channels . Reach millions of shoppers and boost sales
Sell globally . International sales
Sell wholesale & direct . Business-to-business (B2B)
Market Market your business . Reach & retain customers
Market across social . Social media integrations
Chat with customers . Shopify Inbox
Nurture customers . Shopify Messaging
Know your audience . Gain customer insights
Manage Manage your business . Track sales, orders & analytics
Measure your performance . Analytics and Reporting
Manage your stock & orders . Inventory & order management
Automate your business . Shopify Flow
Shopify Developers . Build with Shopify's powerful APIs
Plus . A commerce solution for growing digital brands
All Products . Explore all Shopify products & features
Resources Help and support Help and support . Get 24/7 support
Business courses . Learn from proven experts
Popular topics What is Shopify? . How our commerce platform works
Essential tools Business name generator .
What’s new Changelog . Your source for recent updates
Newsroom . All company news and press releases
blog | AI & Machine Learning Gisting: Compressing LLM Agent context to ↑ throughput and ↓ cost
Gisting compresses context into a set of learned tokens, preserving its quality while making the model faster and cheaper.
System prompts can account for thousands of tokens per request. A longer prompt means slower, more expensive inference. This results in more GPUs required to accommodate the same traffic when serving a model on dedicated hardware.
Our implementation of gisting, as first proposed in Prompt Compression and Contrastive Conditioning for Controllability and Toxicity Reduction in Language Models ,¹ allows us to reap the behavioral advantages of a long prompt at the cost of a short one.
Through gisting, we cut down the Sidekick GraphQL agent’s system prompt from ~6,000 tokens to ~1,500 gist tokens (a 4:1 reduction) without losing prediction quality. We achieve this by learning the embeddings for a set of special tokens via knowledge distillation and swapping them with the system prompt at inference time.
The serving gains from a 4:1 compression of the system prompt into gist tokens are significant. At 350 requests per minute (RPM), the median time to first token (TTFT) dropped from 438ms to 354ms, the median end-to-end request latency dropped from 6.8s to 4.2s, and throughput rose from 20.2 to 23.4 queries per second (QPS). These gains allowed us to reduce the number of GPUs allocated for the GraphQL agent’s traffic.
A gist toke n is a special token that we add to the model's vocabulary. We train the sequence of embeddings so that their substitution into the context induces the model to behave as though it’s seen the full prompt. At 4:1 compression, we add one gist token for every four prompt tokens. We freeze the weights of the model and only train the gist embeddings.
We learn the gist embeddings via knowledge distillation. We run the forward pass twice on each trajectory. During the teacher pass, the model sees the full natural-language prompt to derive the teacher logits for each position in the model’s response. For the student pass, we swap the full prompt for the gist tokens and run the same model again to derive the student logits corresponding to each teacher response position. We train the gist embeddings using the KL divergence between the teacher logits and the student logits, until the student's predictions closely match the teacher's.
Deploying a gisted model is simple. When training finishes, we write the gist embeddings straight into the model's embedding matrix, and register the new gist tokens as special tokens in the model’s tokenizer. The model loads and runs like any other at inference time: no custom attention mask, extra encoder, or special serving path. The only inference-time change is on the request-side: replacing the prompt with the string of gist tokens. The entire cost of compression is paid once, at training time.
Why prefix caching isn't enough
The advantages of gisting and prefix caching are not mutually exclusive. All modern serving engines maintain a KV cache that stores the keys and values for previously-seen sequences. When new requests include a sequence that exists in the cache (including the system prompt, usually), the KV tensors for that sequence are fetched instead of being recomputed.
Though prefix caching is powerful, it doesn't eliminate the decode cost. Every time the model generates a token, that token attends over every key in the sequence, cached or not. Because decode is memory-bandwidth-bound, every generated token must stream the entire KV cache from high bandwidth memory, and that read grows linearly with the cached sequence length. Gisting reduces the cost of attention computations and KV cache reads, the latter of which is especially impactful on throughput when batch sizes grow large.
The optimizations of gisting and prefix caching compound, and we use both in our experiments and production serving stacks.
To tune our hyperparameters, we pointed an autoresearch loop at the trainer. It proposes a recipe, trains gist embeddings, evaluates the resulting model, and repeats.
During the autoresearch loop, three optimizations in particular had high impact:
Initialization : Instead of initializing the embeddings with random noise, we split the system prompt into sequences of length k (where k is derived from the k :1 compression ratio) and initialized the n th gist embedding with the mean of the n th system prompt chunk. This optimization reduced our initial loss by a factor of 7.
Compression : After experimenting with a range of compression ratios, we found the ratio beyond which prediction quality began to degrade (4:1 was the optimal ratio for the complexity of our domain; other domains vary).
Data quantity and diversity : Curating a large and diverse dataset closed the remaining gap.
In addition to streamlining hyperparameter tuning, autoresearch also helped us implement several key optimizations in our training infrastructure. The first was how we normalized the loss: averaging per response token made the model hallucinate, while averaging over the batch preserved more signal from long responses and produced stable embeddings. The second was speed: precomputing the teacher logits and pre-tokenizing the data cut a full run from thirty hours to six.
The payoff: latency and throughput
We ran load tests comparing the same model with the compressed prompt against the full-prompt.
Gisting results in significant latency gains over the full prompt, and the gap becomes even more pronounced as request concurrency rises and batches get larger. At 350 RPM, TTFT dropped 19%, E2E latency dropped about 38%, and throughput increased by 16%.
In practice, that 16% gain in throughput translated directly into GPU savings on our production workload. With the hardware configuration serving our GraphQL traffic, we were able to use 14% fewer GPUs with gisting.
Gisting and continual learning
Gisting also feeds nicely into Shopify’s continual learning loop . Once we have the distilled gist embeddings for a model, we can treat that model as a new starting point for continual learning. We can post-train using the gist embeddings as the prefix and apply the gradient updates to both the model weights and the gist embeddings. By optimizing both the weights and gist embeddings on the incremental data, we can continuously calibrate and improve the model without the computational load of distilling the gist embeddings from scratch each time.
Long system prompts are useful but expensive. Gisting allows agents to leverage all the advantages of extensive instructions with a fraction of the tokens, reducing their latency and GPU spend. In the recent era of GPU demand outpacing supply, every inference optimization matters, but we also refuse to compromise on quality. Gisting gives us both, and at Shopify, it is the new standard.
1. Wingate, Shoeybi, and Sorensen. Prompt Compression and Contrastive Conditioning for Controllability and Toxicity Reduction in Language Models. 2022 .
See our open roles and learn more about our digital by design culture.
