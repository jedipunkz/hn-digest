---
source: "https://www.liquid.ai/blog/lfm2-5-8b-a1b"
hn_url: "https://news.ycombinator.com/item?id=48325306"
title: "Liquid AI reveals 8B-A1B MoE trained on 38T"
article_title: "LFM2.5-8B-A1B: an Even Better on-Device Mixture-of-Experts | Liquid AI"
author: "simjnd"
captured_at: "2026-05-30T11:43:39Z"
capture_tool: "hn-digest"
hn_id: 48325306
score: 193
comments: 75
posted_at: "2026-05-29T16:19:54Z"
tags:
  - hacker-news
  - translated
---

# Liquid AI reveals 8B-A1B MoE trained on 38T

- HN: [48325306](https://news.ycombinator.com/item?id=48325306)
- Source: [www.liquid.ai](https://www.liquid.ai/blog/lfm2-5-8b-a1b)
- Score: 193
- Comments: 75
- Posted: 2026-05-29T16:19:54Z

## Translation

タイトル: Liquid AI により、38T でトレーニングされた 8B-A1B MoE が明らかに
記事のタイトル: LFM2.5-8B-A1B: より優れたオンデバイス専門家の混合 |リキッドAI
説明: 本日、当社は LFM2.5-8B-A1B をリリースします。LFM2.5-8B-A1B は、コンシューマ ハードウェアでの高速で信頼性の高いツール呼び出しと複雑な命令に続くように最適化された高スループット エッジ モデルで、はるかに大規模なモデルに匹敵する圧縮パフォーマンスと、主要な推論フレームワーク全体での初日サポートを提供します。

記事本文:
LFM2.5-8B-A1B: より優れたオンデバイス専門家の混合 |リキッドAI
製品
リキッドファンデーションモデル
価格設定とライセンスのソリューション
エンタープライズ ソリューション 市場で最も効率的な AI でビジネスの可能性を広げます。スタートアップ ソリューション 専門家によるサポートにより、貴社のスタートアップに最適な LFM を作成できます。自動車 家庭用電化製品 eコマース 金融 ヘルスケアおよびライフ サイエンス 産業およびロボティクス リソース
ブログ ケーススタディ デモ ドキュメント ドキュメント、使用例、クックブック 開発者コミュニティ ハッカソン 会社
私たちのパートナーについて ニュースルーム キャリア
液体を入手
モデルのダウンロード LEAP HuggingFace Amazon Bedrock モデルのカスタマイズとデプロイ LEAP モデルで遊ぶ Liquid Playground Liquid Apollo (App Store) Liquid Apollo (Google Play) デモをリクエストする Liquid を試す
製品
当社の製品 リキッドファンデーションモデル
価格設定とライセンス ソリューション
産業 自動車 家電 電子商取引 金融 ヘルスケア & ライフ サイエンス 産業 & ロボティクス 効率と能力の融合 エンタープライズ ソリューション 市場で最も効率的な AI でビジネスの可能性を解き放ちます。スタートアップ ソリューション 専門家によるサポートにより、貴社のスタートアップに最適な LFM を作成できます。研究リソース
ブログ ケーススタディ デモ ドキュメント ドキュメント、使用例、クックブック 開発者コミュニティ ハッカソン会社
当社のパートナーについて
ニュースルーム
英語 日本語 モデル LFM2.5-8B-A1B: より優れたオンデバイス専門家の混合
本日、消費者向けハードウェアでの高速で信頼性の高いツール呼び出し用に構築されたエッジ モデルである LFM2.5-8B-A1B をリリースします。
これは、2025 年 10 月からの LFM2-8B-A1B リリースに基づいて構築されており、拡張された 128K コンテキスト ウィンドウ、スケールアップされた事前トレーニング (12T から 38T トークン)、および大規模な強化学習が備えられています。また、非ラテン言語のトークン化効率を向上させるために語彙を 2 倍にしました。

歳。その結果、ツール呼び出しを連鎖させ、タスクを達成し、エントリーレベルのラップトップにも快適にフィットするモデルが誕生しました。
ベース (LFM2.5-8B-A1B-Base) モデルとトレーニング後モデル (LFM2.5-8B-A1B) は、Hugging Face と Playground で現在入手可能です。ローカルで実行して微調整する方法については、ドキュメントをご覧ください。
オンデバイスのパーソナルアシスタント。実際のアプリケーションを強化し、ツール呼び出しを連鎖させ、すべてのデバイスで複雑な命令に従うように設計されています。
圧縮されたパフォーマンス。命令フォローやエージェントタスクに関しては、はるかに大規模な高密度モデルや MoE モデルと競合します。
比類のないスループット。 CPU と GPU 推論の両方でそのサイズクラスで最速で、llama.cpp、MLX、vLLM、および SGLang を初日からサポートします。
LFM2-8B-A1Bからの変更点
LFM2-8B-A1B と比較して、この新しいバージョンではコンテキスト ウィンドウが 32,768 トークンから 128,000 トークンに拡張されます。これにより、モデルはより長いドキュメントを処理し、より長い理由を推論できるようになります。非ラテン文字をより効率的にトークン化するために、その語彙サイズも 65,536 から 128,000 にスケールアップされました。特にヒンディー語、タイ語、ベトナム語、インドネシア語、アラビア語で大きな圧縮効果が見られます。次の図に示すように、アーキテクチャの残りの部分は、LFM2-8B-A1B と同じ MoE、GQA、およびゲート ショート コンボリューション ブロックの組み合わせに従います。
前任者とは異なり、LFM2.5-8B-A1B は推論のみのモデルであり、最終的な答えの前に明示的な思考の連鎖を生成します。 MoE モデルは通常、コンピューティング バウンド設定で実行され、アクティブなパラメーターの数が少ないため、各推論トークンが安価になるため、この戦略を採用しました。これにより、速度を損なうことなく品質が大幅に向上します。
推論とスケールアップされたトレーニングのおかげで、この新しいバージョンのパフォーマンスは大幅に向上しました。
トークナイザーの拡張。 LFM2-8B-A1B は元々 65K B でトレーニングされました

初期の言語範囲に合わせて最適化された PE トークナイザー。 LFM2.5 で非ラテン文字をより適切にサポートするために、モデルを最初から再トレーニングするのではなく、既存のトークナイザーを拡張することで語彙を 128K に倍増しました。多言語コーパス上の元のマージから BPE マージ トレーニングを継続しました。これにより、ほとんどの既存のトークン ID が ID マッピングとして保持され、すべての新しいトークンが一連の元のサブトークンに決定論的に分解されます。新しい埋め込み行をそのサブトークン分解の平均として初期化し、共有行を変更せずにコピーします。次に、埋め込みのみのトレーニングと、それに続くフルモデルの継続的な事前トレーニングという、簡単な 2 段階の適応を通じて品質を回復します。
以下の表は、文字数/トークン、各トークンが運ぶおおよそのテキスト量を示しています。高いほど優れており、新しいトークナイザーは 16 言語すべてで効率的です。
コンテキスト拡張。まず、推論、数学、ツールの使用、および長いドキュメントに焦点を当てた 2T トークンの中間トレーニング フェーズを通じて、コンテキスト ウィンドウを 32K に拡張しました。次に、RoPE ベース θ を増やし、長いドキュメントと長い軌道データに焦点を当てたトレーニング段階中に追加の 400B トークンを実行することで、コンテキストを 128K に拡張しました。
ドゥームループ。長い推論トレースにおける破滅ループを削減するために、対象を絞った設定最適化ステージを追加しました。このステージでは、特定のコンテキストでループ動作を引き起こす傾向があるトークンを特定し、次のトークンの分布の残りの部分をほぼそのまま残したまま、確率質量をもっともらしい代替案に再配分します。 RL 中には、「待って…」などのループを誘発する一般的な再起動単語の過剰な使用を防止する軽量の整形報酬も追加しました。パイプライン全体、客観的、実証的な結果の詳細については、専用のブログ投稿で共有します。
母趾

国々。パラメータの数が少ないため、エッジ モデルの知識容量は限られており、これが幻覚の増加につながります。幻覚を軽減するために、多様な知識データセットに対する avg@k ベースの報酬を使用する、ターゲットを絞った RL ステージを追加しました。目標は、既存の知識を維持しながら、信頼できる知識を超えたクエリに対する棄権を強化することです。これにより、知識の境界がより明確になり、不確実性がより明確に表現されます。
私たちは、知識、指示に従い、数学、およびエージェントのワークフローをカバーするベンチマーク全体で LFM2.5-8B-A1B を評価しました。このモデルは、パラメータの総数が同様の高密度の代替案とはるかに大きな MoE の両方と競合します。
avg@k ベースの報酬により、LFM2.5-8B-A1B は妥当な精度を維持しながら、大幅に低い幻覚率を達成できます。また、ベンチマークに従って命令をリードし、アクティブなパラメータ数の一部で Gemma 4-26B のようなより大きな MoE に匹敵します。
エージェントのベンチマークでは、LFM2.5-8B-A1B はより大きなモデルと競争力があり、特に Tau2-Telecom で強力です。エージェント ハーネスがモデルを使用する主な方法になりつつあるため、LFM2.5-8B-A1B は、デバイス上の完全プライベート エージェントをパワーオンするための第一歩となります。
LFM2.5-8B-A1B には、推論エコシステム全体にわたる初日サポートが付属しています。
LEAP — iOS および Android 展開用の Liquid のエッジ AI プラットフォーム
llama.cpp — 効率的なエッジ推論のための GGUF チェックポイント
MLX — Apple Silicon向けに最適化された推論
vLLM — 実稼働スループットを実現する GPU 高速化サービス
SGLang — 実稼働スループットのための GPU 高速化サービス
ONNX — 多様なアクセラレータにわたるクロスプラットフォーム推論
CPU推論。 LFM2.5-8B-A1B には、初日の llama.cpp サポートが付属しており、日常のコンシューマ ハードウェアで動作します。
両方のラップトップクラスのチップ上で最速のモデルです

プロンプトの読み取りと回答の生成をテストし、6 GB 未満に抑えながら、M5 Max では 253 トークン/秒、Ryzen AI Max+ 395 では 146 トークン/秒をデコードしました。電話機には最大 30 トークン/秒も保持できるため、有能なアシスタントが自分のデバイス上で即座に非公開で実行されます。
GPU推論。これらのコードベースへの積極的な貢献を通じて、vLLM および SGLang による推論をサポートします。持続負荷設定を使用して、単一の NVIDIA H100 SXM5 GPU で出力スループット (総出力トークンを経過時間で割ったもの) を測定します。各同時実行レベルで、実行中のリクエストの目標数を継続的に維持し、完了した各リクエストをすぐに置き換えます。
BF16 では、SGLang 0.5.12、1,024 個の入力トークン、最大 256 個の出力トークンを使用して各モデルをベンチマークし、同時実行レベルごとに平均 3 回の実行を実行しました。 LFM2.5-8B-A1B は、そのサイズ クラスで最速のモデルであり、高い同時実行で 1 秒あたり 18.5,000 の出力トークンに達し、単一の H100 で 1 日あたり 16 億トークンを超えます。
オープンソースのデスクトップ エージェント デモである LocalCowork は、LFM2.5-8B-A1B で動作するようになりました。セットアップは 3 月の LFM2-24B-A2B デモで使用したものと同じです。1 台のラップトップ、13 台の MCP サーバーにわたる 67 のツール、クラウドなし、API キーなし、マシンからデータが流出することはありません。ツールの選択は、同じツール メニュー全体でより高速になり、信頼性が大幅に向上しました。
デモのポイントは個々のツールではありません。それは、ツールのディスパッチ ループがコンシューマー ハードウェア上でインタラクティブに感じられることです。質問、提案、確認、実行、繰り返しがすべてディスパッチごとに 1 秒未満で行われ、完全な監査証跡があり、データがデバイスから流出することはありません。
LFM2.5 により、私たちはどこでも実行できる AI というビジョンを実現します。これらのモデルは次のとおりです。
オープンウェイト — 制限なしでダウンロード、微調整、展開します
初日から迅速 — Apple、AMD、Intel、Qualcomm、および Nvidia ハード全体での llama.cpp、MLX、vLLM、SGLang のネイティブ サポート

ウェア
完全なファミリー — カスタマイズ用のベースモデルから特殊なオーディオおよびビジョンのバリアントまで、1 つのアーキテクチャで多様なユースケースをカバーします
オンデバイス エージェントの未来はここから始まります。あなたが何を構築するのか楽しみです。
Liquid AI、「LFM2.5-8B-A1B: ラップトップ上のパーソナル アシスタント」
Liquid AI ブログ、2026 年 5 月。
@article{liquidAI20268BA1B,
著者 = {Liquid AI}、
title = {LFM2.5-8B-A1B: ラップトップ上のパーソナル アシスタント},
ジャーナル = {Liquid AI ブログ}、
年 = {2026}、
注 = {https://www.liquid.ai/blog/lfm2-5-8b-a1b}、
}
AI を体験する準備はできましたか? Liquid AI でビジネス、ワークフロー、エンジニアを強化します。
私たちのチームに相談してください Liquid の最新情報を入手してください ありがとうございます!おっと！フォームの送信中に問題が発生しました。
Liquid Apolloを無料で入手
携帯電話で直接 LFM をバイブチェックしてください。
ソリューション エンタープライズ ソリューション スタートアップ ソリューション 自動車 家庭用電化製品 eコマース 金融サービス ヘルスケア & ライフ サイエンス 産業 & ロボティクス ケース スタディ テクノロジー & 製品 リキッド ファウンデーション モデル LEAP Liquid Apollo 研究 価格設定とライセンス 開発者リソース デモ ドキュメント 開発者コミュニティ ハッカソン 会社概要 ブログ プレス 採用情報
お問い合わせ 法的 LFM ライセンス プライバシー ポリシー プライバシーの選択 利用規約 Liquid Apollo を無料で入手
携帯電話で直接 LFM をバイブチェックしてください。
© {年} 、Liquid AI, Inc. 無断複写・転載を禁じます。
314 Main St、ケンブリッジ、MA 02142
お客様のブラウジング エクスペリエンスを向上させ、トラフィックを分析するために Cookie を使用します。 「すべてを受け入れる」をクリックすると、Cookie の使用に同意したことになります。 Cookie の設定を制御するには、「カスタマイズ」を選択します。
お客様の設定を管理する 当社では、お客様のブラウジング エクスペリエンスを向上させ、トラフィックを分析するために Cookie を使用します。 「すべてを受け入れる」をクリックすると、Cookie の使用に同意したことになります。

## Original Extract

Today, we’re releasing LFM2.5-8B-A1B, a high-throughput edge model optimized for fast, reliable tool calling and complex instruction following on consumer hardware, delivering compressed performance competitive with much larger models and day-one support across major inference frameworks.

LFM2.5-8B-A1B: an Even Better on-Device Mixture-of-Experts | Liquid AI
Products
Liquid Foundation Models
Pricing & Licensing Solutions
Enterprise Solutions Unlock your business with the most efficient AI on the market. Startup Solutions Expert support in creating the LFM that’s right for your startup. Automotive Consumer Electronics Ecommerce Finance Healthcare and Life Sciences Industrial & Robotics Resources
Blog Case Studies Demos Documentation Docs, use cases, and cookbooks Developer Community Hackathons Company
About Our Partners Newsroom Careers
Get Liquid
Download Models LEAP HuggingFace Amazon Bedrock Customize & Deploy Models LEAP Play with Models Liquid Playground Liquid Apollo on App Store Liquid Apollo on Google Play Request a Demo Try Liquid
Products
OUR PRODUCTS Liquid Foundation Models
PRICING Pricing & Licensing Solutions
INDUSTRIES Automotive Consumer Electronics Ecommerce Finance Healthcare & Life Sciences Industrial & Robotics EFFICIENCY MEETS CAPABILITY Enterprise Solutions Unlock your business with the most efficient AI on the market. Startup Solutions Expert support in creating the LFM that’s right for your startup. Research Resources
Blog Case Studies Demos Documentation Docs, use cases, and cookbooks Developer Community Hackathons company
About Our Partners Careers
Newsroom
English Japanese Models LFM2.5-8B-A1B: an Even Better on-Device Mixture-of-Experts
Today, we're releasing LFM2.5-8B-A1B , an edge model built for fast, reliable tool calling on consumer hardware.
It builds on our LFM2-8B-A1B release from October 2025, with an expanded 128K context window, scaled-up pretraining (from 12T to 38T tokens), and large-scale reinforcement learning. We also doubled its vocabulary to improve tokenization efficiency for non-Latin languages. The result is a model that chains tool calls, achieves tasks, and fits comfortably even on an entry-level laptop.
The base (LFM2.5-8B-A1B-Base) and post-trained (LFM2.5-8B-A1B) models are available today on Hugging Face and our Playground . Check out our docs on how to run and fine-tune them locally.
On-device personal assistant. Designed to power real-life applications, chaining tool calls, and following complex instructions on all devices.
Compressed performance. Competitive with much larger dense and MoE models on instruction following and agentic tasks.
Unmatched throughput. Fastest in its size class on both CPU and GPU inference, with day-one support for llama.cpp, MLX, vLLM, and SGLang.
What changed since LFM2-8B-A1B
Compared to LFM2-8B-A1B, this new version expands the context window from 32,768 to 128,000 tokens . This allows the model to process longer documents and reason for longer. Its vocabulary size was also scaled up from 65,536 to 128,000 to tokenize non-Latin scripts more efficiently . We see particularly strong compression gains in Hindi, Thai, Vietnamese, Indonesian, and Arabic. The rest of the architecture follows the same combination of MoE, GQA, and gated short convolution blocks as LFM2-8B-A1B, as shown in the following figure.
Unlike its predecessor, LFM2.5-8B-A1B is a reasoning-only model , producing an explicit chain of thought before its final answer. We adopted this strategy because MoE models generally run in compute-bound settings, where a smaller number of active parameters makes each reasoning token cheap. This provides a significant quality boost without compromising speed.
Thanks to reasoning and scaled-up training, this new version performs significantly better:
Tokenizer expansion. LFM2-8B-A1B was originally trained with a 65K BPE tokenizer optimized for our initial language coverage. To better support non-Latin scripts in LFM2.5, we doubled the vocabulary to 128K by extending the existing tokenizer in place rather than retraining the model from scratch.. We continued BPE merge training from the original merges on a multilingual corpus, which keeps most existing token IDs as identity mappings and makes every new token decompose deterministically into a sequence of original sub-tokens. We initialize the new embedding rows as the mean of their sub-token decompositions and copy the shared rows unchanged. We then recover quality through a brief two-stage adaptation: embedding-only training, followed by full-model continued pretraining.
The table below reports chars/token, roughly how much text each token carries: higher is better, and the new tokenizer is more efficient in all 16 languages
Context extension. We first extended the context window to 32K through a 2T token midtraining phase focused on reasoning, math, tool-use, and longer documents. We then extended the context to 128K by increasing the RoPE base θ and running an additional 400B token midtraining stage focused on long-document and long-trajectory data.
Doom loops. We added a targeted preference optimization stage to reduce doom loops in long reasoning traces. This stage identifies tokens that tend to trigger looping behavior in specific contexts, then redistributes probability mass toward plausible alternatives, while leaving the rest of the next-token distribution largely intact. During RL, we also added a lightweight shaping reward that discourages excessive use of common loop-inducing restart words like “Wait…”. We'll share more details on the full pipeline, objective, and empirical results in a dedicated blog post.
Hallucinations. Because of their small number of parameters, edge models have a limited knowledge capacity, which leads to more hallucinations. To mitigate hallucinations, we added a targeted RL stage that uses an avg@k-based reward over a diverse knowledge dataset. The goal is to reinforce abstention on queries beyond reliable knowledge while preserving existing knowledge. This produces a sharper knowledge boundary and clearer expression of uncertainty.
We evaluated LFM2.5-8B-A1B across benchmarks covering knowledge, instruction following, math, and agentic workflows. The model is competitive with both dense alternatives with a similar total number of parameters and much larger MoEs.
The avg@k-based reward enables LFM2.5-8B-A1B to achieve a significantly lower hallucination rate while maintaining reasonable accuracy. It also leads on instruction following benchmarks, matching bigger MoEs like Gemma 4-26B at a fraction of the active parameter count.
On agentic benchmarks, LFM2.5-8B-A1B is competitive with bigger models and particularly strong on Tau2-Telecom. As agentic harnesses are becoming the main way to consume models, LFM2.5-8B-A1B is a first step towards powering on-device, fully private agents.
LFM2.5-8B-A1B ships with day-one support across the inference ecosystem:
LEAP — Liquid's Edge AI Platform for iOS and Android deployment
llama.cpp — GGUF checkpoints for efficient edge inference
MLX — Optimized inference for Apple Silicon
vLLM — GPU-accelerated serving for production throughput
SGLang — GPU-accelerated serving for production throughput
ONNX — Cross-platform inference across diverse accelerators
CPU inference. LFM2.5-8B-A1B ships with day-one llama.cpp support and runs on everyday consumer hardware.
On both laptop-class chips, it is the fastest model we tested at reading in prompts and generating answers, decoding 253 tokens/s on an M5 Max and 146 on a Ryzen AI Max+ 395 while staying under 6 GB. It even holds ~30 tokens/s on a phone, so a capable assistant runs instantly and privately on your own device.
GPU inference. We support inference via vLLM and SGLang via active contributions to these codebases. We measure output throughput (total output tokens divided by wall time) on a single NVIDIA H100 SXM5 GPU using a sustained-load setting: at each concurrency level, we continuously maintain the target number of in-flight requests, replacing each completed request immediately.
We benchmark each model with SGLang 0.5.12, 1,024 input tokens, up to 256 output tokens, in BF16, averaging 3 runs per concurrency level. LFM2.5-8B-A1B is the fastest model in its size class, reaching 18.5K output tokens per second at high concurrency, over 1.6B tokens per day on a single H100.
Our open-source desktop agent demo, LocalCowork , now runs on LFM2.5-8B-A1B. The setup is the same one we used for LFM2-24B-A2B demo in March: a single laptop, 67 tools across 13 MCP servers, no cloud, no API keys, no data leaving the machine. Tool selection is faster and noticeably more reliable across the same tool menu.
The point of the demo is not the individual tools. It is that the tool-dispatch loop feels interactive on consumer hardware: ask, propose, confirm, run, repeat, all in well under a second per dispatch, with full audit trails and your data never leaving the device.
With LFM2.5, we're delivering on our vision of AI that runs anywhere. These models are:
Open-weight — Download, fine-tune, and deploy without restrictions
Fast from day one — Native support for llama.cpp, MLX, vLLM, SGLang across Apple, AMD, Intel, Qualcomm, and Nvidia hardware
A complete family — From base models for customization to specialized audio and vision variants, one architecture covers diverse use cases
The on-device agentic future starts here. We can't wait to see what you build.
Liquid AI, “LFM2.5-8B-A1B: Personal Assistant On Your Laptop,”
Liquid AI Blog , May 2026.
@article{liquidAI20268BA1B,
author = {Liquid AI},
title = {LFM2.5-8B-A1B: Personal Assistant On Your Laptop},
journal = {Liquid AI Blog},
year = {2026},
note = {https://www.liquid.ai/blog/lfm2-5-8b-a1b},
}
Ready to experience AI? Power your business, workflows, and engineers with Liquid AI.
Talk to our team Stay up to date with Liquid Thank you! Oops! Something went wrong while submitting the form.
Get Liquid Apollo for free
Vibe check LFMs directly on your phone.
Solutions Enterprise Solutions Start Up Solutions Automotive Consumer Electronics Ecommerce Financial Services Healthcare & Life Sciences Industrial & Robotics Case Studies Tech & Products Liquid Foundation Models LEAP Liquid Apollo Research Pricing & Licensing Developer Resources Demos Documentation Developer Community Hackathons Company About Blog Press Careers
Contact Us Shoutouts Legal LFM License Privacy Policy Privacy Choices Terms & Conditions Get Liquid Apollo for free
Vibe check LFMs directly on your phone.
© {year} , Liquid AI, Inc. All rights reserved.
314 Main St, Cambridge, MA 02142
We use cookies to enhance your browsing experience and analyze our traffic. By clicking "Accept All", you consent to our use of cookies. To control your cookie settings, choose "Customize".
Manage your preferences We use cookies to enhance your browsing experience and analyze our traffic. By clicking "Accept All", you consent to our use of cookies.
