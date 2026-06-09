---
source: "https://cloud.google.com/blog/products/containers-kubernetes/gke-inference-gateway-prefix-caching-accelerates-ai-inference"
hn_url: "https://news.ycombinator.com/item?id=48462961"
title: "GKE Inference Gateway delivers up to 92% faster AI responses"
article_title: "GKE Inference Gateway prefix caching accelerates AI inference | Google Cloud Blog"
author: "ilreb"
captured_at: "2026-06-09T18:53:33Z"
capture_tool: "hn-digest"
hn_id: 48462961
score: 1
comments: 0
posted_at: "2026-06-09T16:09:03Z"
tags:
  - hacker-news
  - translated
---

# GKE Inference Gateway delivers up to 92% faster AI responses

- HN: [48462961](https://news.ycombinator.com/item?id=48462961)
- Source: [cloud.google.com](https://cloud.google.com/blog/products/containers-kubernetes/gke-inference-gateway-prefix-caching-accelerates-ai-inference)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T16:09:03Z

## Translation

タイトル: GKE Inference Gateway は AI 応答を最大 92% 高速化します
記事のタイトル: GKE Inference Gateway のプレフィックス キャッシュにより AI 推論が高速化 | Google クラウド ブログ
説明: GKE Inference Gateway は、プレフィックス キャッシュとその他のルーティング アルゴリズムを使用して、ジョブに最適なアクセラレータにリクエストを送信します。

記事本文:
GKE Inference Gateway のプレフィックス キャッシュにより AI 推論が高速化 | Google Cloud ブログ コンテンツに移動 クラウド ブログ 営業担当者へのお問い合わせ 無料で始めましょう クラウド ブログ ソリューションとテクノロジー AI と機械学習
コンテナと Kubernetes レポート: GKE Inference Gateway は AI 応答を最大 92% 高速化します
職場における AI への玄関口
生成 AI が実験的なパイロット環境から大規模な実稼働環境に移行するにつれて、インフラストラクチャの効率が究極の差別化要因になります。これを最大限に活用し、コストのかかるアクセラレータのアイドル時間を最小限に抑える 1 つの方法は、Google Kubernetes Engine (GKE) Inference Gateway を活用することです。これは、リアルタイム モデル サーバー メトリクスに基づいて生成 AI ワークロードをインテリジェントにルーティングします。
従来の単純なラウンドロビン負荷分散に依存する代わりに（頻繁に高価なアクセラレータの再計算がトリガーされ、ユーザーのレイテンシが急増します）、GKE ゲートウェイのこのネイティブ拡張機能は、プレフィックス キャッシュやモデル認識ルーティングなどの高度な機能を利用します。 GKE は、リクエストをすぐに処理できる正確なアクセラレータに確実に到達するようにすることで、優れたハードウェア使用率と超高速な応答時間で大規模言語モデル（LLM）を提供する方法を変革します。
実際、独立したベンチマーク レポートによると、GKE Inference Gateway は、スループットが 15.7% 高く、待機時間が 92.8% 短く、トークン間レイテンシが 62.6% 低く、次の主要なマネージド Kubernetes サービスよりも優れています。このパフォーマンスにより、LLM ベースのアプリケーションが、低速で高価なアプリケーションから、高速で実稼働グレードのアプリケーションに変わります。
そのパフォーマンスは、GKE Inference Gateway を使用した Snap のエクスペリエンスによって追跡されます。
「Snap では、大規模な高性能推論を促進するために、llm-d を実稼働 AI インフラストラクチャに統合しています。

プレフィックス キャッシュを認識したルーティングにより、最大 75 ～ 80% のプレフィックス キャッシュ ヒット率を達成しました。 Envoy ベースのサービス メッシュとのシームレスな統合を可能にする llm-d のオープンソースの性質を高く評価しています。」 - Vinay Kola 氏、Snap Inc. ソフトウェア エンジニアリング担当シニア マネージャー
このブログでは、GKE Inference Gateway のプレフィックス キャッシュについて例を交えて詳しく説明します。ベンチマーク結果についても詳しく説明します。飛び込みましょう。
低遅延 AI の秘密: プレフィックス キャッシュ
プレフィックス キャッシュは、長く繰り返されるプロンプト プレフィックスの KV キャッシュ (アクティブ化状態) を保存することにより、LLM のパフォーマンスを最適化します。連続するユーザー要求が同じシステム命令、コンテキスト、またはドキュメントを共有する場合、モデルはそれらのトークンの再処理を完全にスキップします。 GKE Inference Gateway は、受信リクエストのプレフィックスを読み取り、そのデータをメモリ内にすでに保持している特定のポッドと照合します。これにより、GPU と TPU にかかる「思考」の負担が軽減され、重い推論ループがほぼ瞬時の答えに変わります。
使用例 1: 検索拡張生成 (RAG) を使用したドキュメントとコードベースの Q&A
大規模なエンタープライズ リポジトリにクエリを実行する場合、RAG を使用してドキュメント セット全体を静的なキャッシュされたプレフィックスとして固定することで、遅延を追加することなく LLM の応答を固定できます。
GKE Inference Gateway は、ユーザーの質問ごとに数千行の API リファレンスや企業 Wiki を再度読み取ることを LLM に強制するのではなく、KV キャッシュでウォームアップされた特定のコンテキストをすでに持っているポッドにクエリをルーティングします。 LLM は、ユーザーの簡潔で動的な質問を計算するだけでよく、コストのかかる文書の再評価を完全に回避します。
読み込み中... [静的プレフィックス - キャッシュに留まる] あなたは、技術文書を専門とするエキスパート AI アシスタントです。以下は完全な API ドキュメントです。

私たちのソフトウェアプラットフォーム。このコンテキストを使用して、ユーザーの質問に正確に答えます。答えがドキュメントで見つからない場合は、「提供されたコンテキストではその答えが見つかりません」と言います。
<documentation> [10,000 語以上の API リファレンス ドキュメント、エンドポイント、エラー コードなど] </documentation>
[動的サフィックス - リクエストごとの変更] ユーザーの質問: Python SDK を使用して 429 レート制限エラーを処理するにはどうすればよいですか?図 1: キャッシュされた静的プレフィックスと要求ごとに変化する動的サフィックスを示すソフトウェア トラブルシューティング シナリオのプロンプトの内訳
また、プレフィックス キャッシュを使用すると、コンピューティング コストを増大させることなく、数千の同時セッションにわたって顧客サービスのやり取りを維持できます。これを行うには、永続的なシステム ペルソナとコア ビジネス ルールを LLM サーバーに直接キャッシュします。
エンタープライズ チャット アーキテクチャでは、ベース システム プロンプトと参照テーブルは、何百万もの顧客との対話を通じて完全に同一のままです。 GKE Inference Gateway は、コンテキスト認識ルーティングを使用してこれらのマルチターン会話を処理し、反復的なトークン処理をバイパスするため、トラフィックのピーク時でもチャットボットの応答性は非常に高くなります。
読み込み中... [静的プレフィックス - キャッシュに残る]
- システム ペルソナ: あなたは「FinBot」、ABC バンキング ソリューションの親切で共感力があり、準拠性の高い仮想アシスタントです。次のルールを厳守する必要があります。 1. 具体的な投資アドバイスを決して提供しないでください。 2. ユーザーが小切手または貯蓄について質問しているかどうかを常に確認します。 3. 回答は 3 文以内にしてください。 4. ユーザーが怒っている場合は、人間のマネージャーにつなぐように申し出ます。
2026 年 5 月の現在の金利表は次のとおりです。
- 割引: 年率 4.2%
- 小切手: 0.5% APR
- CD (12 か月): 年率 5.1%
[動的サフィックス - リクエストごとの変更] ユーザー: こんにちは、ロックした場合にいくら稼げるか調べています。

1年で1万ドルもらえる？図 2: 上記のプロンプトは、銀行チャットボット インタラクションの静的なプレフィックスと動的、リクエストごとのコンポーネントを示しています。
GKE は代替のマネージド Kubernetes ソリューションよりも優れたパフォーマンスを発揮します
これらのアーキテクチャ上の利点を検証するために、Principled Technologies は最近、GKE (GKE Inference Gateway を搭載) と従来のラウンドロビン HTTP 負荷分散を利用する標準のサードパーティ マネージド Kubernetes サービスを比較する独立したベンチマーク レポートをリリースしました。
同一のハードウェア (NVIDIA A100 40GB GPU 8 個) を使用して Llama 3.1 8B Instruct 共有プレフィックス ワークロードでテストした結果、2 つの Kubernetes サービス間に大きなパフォーマンスのギャップがあることが明らかになりました。 GKE は単に勝っただけではありません。これにより、次の 3 つの重要な指標にわたる推論効率が完全に再定義されました。
スループットの向上: 1 秒あたりに処理されるトークンの数が 15.7% 増加し、同じワークロードでのリクエスト容量の増加またはハードウェアのニーズの削減が可能になります。
最初のトークンまでの時間 (TTFT) の大幅な短縮: 待ち時間が 92.8% 短縮され、対話型シナリオでの応答開始が劇的に速く認識されます。
トークン間レイテンシ (ITL) の短縮: 62.6% 削減され、最初のトークン後のトークン ストリーミングがよりスムーズかつ高速になります。
図 3: GKE Inference Gateway と Llama 3.1-8B 上のサードパーティ マネージド Kubernetes サービスを使用した GKE の平均レイテンシ（出力トークンあたりの正規化時間） 共有プレフィックスの使用例について LLM を指示します。どちらのソリューションも同じハードウェアを使用しました。出典: Principled Technologies
サードパーティのマネージド Kubernetes サービス
1 秒あたり 7,169.21 出力トークン
1 秒あたり 6,042.05 出力トークン
出力トークンのスループットが 15.7% 増加
最初のトークンまでの平均時間 (TTFT)
平均トークン間レイテンシー (ITL)
図 4: GKE Inference Gateway を使用した GKE は、サードパーティのマネージド Kubernetes サービスと比較して優れた AI 推論を実現しました

標準の HTTP LB を使用します。
生成 AI 推論ワークロードを高速化する準備はできていますか?
リアルタイムのカスタマー サポート エージェント、ダイナミック コーディング アシスタント、または 1 秒未満の不正検出モデルなどの推論ワークロードを導入する場合、インフラストラクチャの遅延がユーザー エクスペリエンスを決定します。 GKE Inference Gateway は、共有プロンプト プレフィックスがほぼ 100% の確率でアクティブ キャッシュにヒットするようにすることで、LLM を低速で高価な推論エンジンから、迅速で資本効率の高い本番レベルの強力なエンジンに変換します。
GKE Inference Gateway が生成 AI ワークロードにもたらすパフォーマンス上の利点を探索する準備はできていますか?詳細については、ここから完全なベンチマーク レポートにアクセスし、この説明ビデオをご覧ください。
Principled Technologies のシニア パフォーマンス アーキテクト、Dan Sullivan に心より感謝いたします。
GKE スタンバイ バッファの紹介: 予算を超過することなくノードの起動時間を短縮します
エヤル・ヤブロンカ著 • 7 分で読めます
GKE 上の Agent Sandbox が誰でも利用できるようになり、Agent Substrate が初めて公開されました
ブランドン・ロイヤル著 • 5 分で読めます
GKE のノード起動が高速化され、コールドスタートのレイテンシに別れを告げます
エヤル・ヤブロンカ著 • 4 分で読めます
ドリュー・ブラッドストック著 • 6 分で読めます
言語 英語 ドイツ語 フランス語 한국어 日本語

## Original Extract

GKE Inference Gateway uses prefix caching and other routing algorithms to land requests on the best accelerator for the job.

GKE Inference Gateway prefix caching accelerates AI inference | Google Cloud Blog Jump to Content Cloud Blog Contact sales Get started for free Cloud Blog Solutions & technology AI & Machine Learning
Containers & Kubernetes Report: GKE Inference Gateway delivers up to 92% faster AI responses
The front door to AI in the workplace
As generative AI moves from experimental pilots to massive production environments, the efficiency of your infrastructure becomes the ultimate differentiator. One way to get the most out of it and minimize costly accelerator idle time is to leverage the Google Kubernetes Engine (GKE) Inference Gateway , which intelligently routes generative AI workloads based on real-time model server metrics.
Instead of relying on traditional, naive round-robin load balancing — which frequently triggers expensive accelerator recomputation and spikes user latency — this native extension of the GKE Gateway utilizes advanced capabilities like prefix caching and model-aware routing . By ensuring requests land on the exact accelerator that is primed to process them right away, GKE transforms how you can serve your large language models (LLMs), with excellent hardware utilization and ultra-fast response times.
In fact, according to an independent benchmark report , GKE Inference Gateway outperforms the next leading managed Kubernetes service with 15.7% higher throughput, 92.8% shorter wait times, and 62.6% lower inter-token latency . This performance takes LLM-based applications from sluggish and expensive to fast and production-grade.
That performance tracks with Snap ’s experience using GKE Inference Gateway.
“At Snap, we are integrating llm-d into our production AI infrastructure to facilitate high-performance inference at scale. By employing prefix-cache-aware routing, we have achieved prefix cache hit rates ranging up to 75-80%. We appreciate the open-source nature of llm-d, as it enables seamless integration with our Envoy-based Service Mesh.” - Vinay Kola, Senior Manager, Software Engineering, Snap Inc.
In this blog, we take a closer look at GKE Inference Gateway’s prefix caching, complete with examples. We also provide more details about its benchmark results. Let’s jump in.
The secret to low-latency AI: Prefix caching
Prefix caching optimizes LLM performance by storing the KV cache (activation states) of long, repetitive prompt prefixes. When consecutive user requests share the same system instructions, context, or documentation, the model entirely skips reprocessing those tokens. GKE Inference Gateway reads incoming request prefixes and matches them to the specific pods that already hold that data in memory. This eliminates the "thinking" tax on your GPUs and TPUs, turning heavy reasoning loops into near-instant answers.
Use case 1: Documentation and codebase Q&A with retrieval-augmented generation (RAG)
When querying massive enterprise repositories, you can ground your LLMs’ responses without any added latency by pinning entire documentation sets as static cached prefixes, using RAG.
Instead of forcing an LLM to re-read thousands of lines of API references or corporate wikis for every single user question, GKE Inference Gateway routes the query to a pod that already has that specific context warmed up in its KV cache. The LLM only has to compute the user's brief, dynamic question, completely bypassing expensive document re-evaluation.
Loading... [STATIC PREFIX - STAYS IN CACHE] You are an expert AI assistant specializing in technical documentation. Below is the complete API documentation for our software platform. Use this context to answer the user's questions accurately. If the answer cannot be found in the documentation, say "I cannot find that in the provided context."
<documentation> [10,000+ words of API reference documentation, endpoints, error codes, etc.] </documentation>
[DYNAMIC SUFFIX - CHANGES PER REQUEST] User Question: How do I handle a 429 rate limit error using the Python SDK? Figure 1: Prompt breakdown for a software troubleshooting scenario showing cached static prefix and dynamic suffix that changes per request
You can also use prefix caching to maintain customer service interactions across thousands of simultaneous sessions without compounding compute costs. You can do so by caching permanent system personas and core business rules directly on the LLM server.
In enterprise chat architectures, the base system prompt and reference tables remain completely identical across millions of customer interactions. GKE Inference Gateway handles these multi-turn conversations using context-aware routing to bypass repetitive token processing, so that your chatbot stays ultra-responsive even under peak traffic.
Loading... [STATIC PREFIX - STAYS IN CACHE]
-System Persona: You are "FinBot", a helpful, empathetic, and compliant virtual assistant for ABC Banking Solutions. You must strictly adhere to the following rules: 1. Never provide concrete investment advice. 2. Always verify if the user is asking about checking or savings. 3. Keep your answers under 3 sentences. 4. If a user is angry, offer to connect them to a human manager.
Here is the current interest rate table for May 2026:
- Savings: 4.2% APR
- Checking: 0.5% APR
- CD (12-month): 5.1% APR
[DYNAMIC SUFFIX - CHANGES PER REQUEST] User: Hi, I'm trying to figure out how much I'd make if I locked away $10,000 for a year? Figure 2: The above prompts illustrate the static prefix and the dynamic, per-request components of a banking chatbot interaction
GKE outperforms alternative managed Kubernetes solutions
To validate these architectural advantages, Principled Technologies recently released an independent benchmark report comparing GKE (equipped with the GKE Inference Gateway) against a standard third-party managed Kubernetes service utilizing conventional round-robin HTTP load balancing.
Tested on a Llama 3.1 8B Instruct shared prefix workload using identical hardware (eight NVIDIA A100 40GB GPUs) the results reveal a massive performance gap between the two Kubernetes services. GKE didn't just win; it completely redefined inference efficiency across three critical metrics:
Higher throughput: 15.7% more tokens processed per second, enabling higher request capacity or reduced hardware needs for the same workload
Much faster time to first token (TTFT): 92.8% shorter wait times, producing dramatically quicker perceived response starts for interactive scenarios
Lower inter-token latency (ITL): 62.6% reduction, resulting in smoother and faster token streaming after the first token
Figure 3: Mean latency (normalized time per output token) of GKE with GKE Inference Gateway and third-party managed Kubernetes service on the Llama 3.1-8B Instruct LLM on the Shared prefix use case. Both solutions used the same hardware. Source: Principled Technologies
3rd party Managed Kubernetes Service
7,169.21 output tokens per second
6,042.05 output tokens per second
15.7% more output token throughput
Mean time to first token (TTFT)
Mean inter-token latency (ITL)
Figure 4: GKE with GKE Inference Gateway delivered superior AI inference compared to a third-party managed Kubernetes service using standard HTTP LB.
Ready to accelerate your gen AI inference workloads?
Whether you’re deploying inference workloads such as real-time customer support agents, dynamic coding assistants, or sub-second fraud detection models, infrastructure latency dictates your user experience. By ensuring shared prompt prefixes hit the active cache nearly 100% of the time, GKE Inference Gateway transforms your LLMs from sluggish, expensive reasoning engines into rapid, capital-efficient, production-grade powerhouses.
Ready to explore the performance advantage that GKE Inference Gateway can bring to your gen AI workloads? Access the full benchmark report here and watch this explainer video to learn more.
A special thanks to Dan Sullivan, Senior Performance Architect , Principled Technologies.
Introducing the GKE standby buffer: Improve node startup times without blowing your budget
By Eyal Yablonka • 7-minute read
Agent Sandbox on GKE is now available for everyone, and a first look at Agent Substrate
By Brandon Royal • 5-minute read
With faster node startup for GKE, say goodbye to cold-start latency
By Eyal Yablonka • 4-minute read
By Drew Bradstock • 6-minute read
Language ‪English‬ ‪Deutsch‬ ‪Français‬ ‪한국어‬ ‪日本語‬
