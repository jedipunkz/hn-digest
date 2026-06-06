---
source: "https://github.com/harshuljain13/llm-inference-at-scale"
hn_url: "https://news.ycombinator.com/item?id=48424467"
title: "Free LLM inference handbook: 100 engineers cloned it in week 1"
article_title: "GitHub - harshuljain13/llm-inference-at-scale: A Practitioner handbook for production llm serving. · GitHub"
author: "harshuljain13"
captured_at: "2026-06-06T15:32:45Z"
capture_tool: "hn-digest"
hn_id: 48424467
score: 2
comments: 0
posted_at: "2026-06-06T12:37:09Z"
tags:
  - hacker-news
  - translated
---

# Free LLM inference handbook: 100 engineers cloned it in week 1

- HN: [48424467](https://news.ycombinator.com/item?id=48424467)
- Source: [github.com](https://github.com/harshuljain13/llm-inference-at-scale)
- Score: 2
- Comments: 0
- Posted: 2026-06-06T12:37:09Z

## Translation

タイトル: 無料の LLM 推論ハンドブック: 100 人のエンジニアが第 1 週でクローンを作成
記事のタイトル: GitHub - Harsuljain13/llm-inference-at-scale: 本番環境の llm サービングのためのプラクティショナー ハンドブック。 · GitHub
説明: プロダクション LLM 提供のためのプラクティショナー ハンドブック。 -hursuljain13/llm-inference-at-scale

記事本文:
GitHub -hursuljain13/llm-inference-at-scale: 本番環境の llm サービングのためのプラクティショナー ハンドブック。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ハーシュルジャイン13
/
llm-大規模推論
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

イオンオプション
コード
Harsuljain13/llm-inference-at-scale
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット アセット アセット コンテンツ コンテンツ リファレンス リファレンス スライド スライド .gitignore .gitignore .python-version .python-version ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイル ナビゲーション
本番環境で大規模な言語モデルを提供するための決定版ガイド。
クイックスタート •
目次 •
研究所 •
コミュニティ •
貢献する
LLM 推論は難しいです。 「ドキュメントを読んで理解する」のは難しくありません。ML の他のすべての作業とは根本的に異なります。
従来の ML 推論は解決された問題です。リクエストをバッチ処理し、転送パスを実行し、結果を返します。レイテンシは予測可能で、メモリは固定されており、スケーリングは線形です。
LLM 推論は、次の前提をすべて破ります。
遅延は予測不可能です。10 トークンの応答には 100 ミリ秒かかり、1000 トークンの応答には 10 秒かかります。
リクエスト中にメモリが増加します - KV キャッシュはトークンが生成されるたびに拡張します
スケーリングは準線形です - GPU を追加すると通信オーバーヘッドが大きくなります
コストは 100 倍になります - リクエストあたり 0.001 ドルがリクエストあたり 0.10 ドルになります
このハンドブックは、私たちが必要としていたのに見つからなかったため存在しています。知識は論文、ブログ投稿、部族の知識、ソース コードのコメントに散在しています。私たちは、長年にわたる制作経験と研究を 1 つの包括的なリソースに統合しました。
これは、私たちが事業を始めたときに存在していればよかったと思っているガイドです。
📬 ビルドをフォローする — 新しい章が、本番環境のコンテキストとともに平易な英語で説明されています。
エンジニア ダイジェストを購読して通知を受け取る
新しいコンテンツがドロップされるとき。無料購読→
💬 ディスカッションに参加する
— 質問、フィードバック、修正を歓迎します
LLM 推論に 100 倍のコストがかかる理由

従来の ML を超える
メモリ帯域幅の壁とその回避方法
プリフィルとデコード: 2 つのまったく異なる問題
バイトレベルでの KV キャッシュの仕組み
vLLM、SGLang、TensorRT-LLM のいずれかを選択する
量子化のトレードオフ (INT8、INT4、FP8、FP4)
Tensor 並列処理とマルチ GPU サービス
キャパシティプランニングとSLO管理
PagedAttend とメモリ効率の高いサービス提供
スループットを高める連続バッチ処理
レイテンシを考慮した投機的デコード
FlashAttendant とカーネルレベルの最適化
細分化されたサービング (プレフィル/デコード分離)
MoE 推論とエキスパート ルーティング
KV キャッシュの圧縮と削除
llama.cpp を使用したエッジ展開
git clone https://github.com/harshuljain13/llm-inference-at-scale.git
cd llm-inference-at-scale
# 仮想環境の作成
Python -m venv .venv
ソース .venv/bin/activate # Windows の場合: .venv\Scripts\activate
# 依存関係をインストールする
pip install -r 要件.txt
読み始める
# 最初の章を開きます
コンテンツ/00_foundations/00.0_what_is_llm_inference/what_is_llm_inference.md を開く
または、以下の目次を参照してください。
章
タイトル
説明
0.0
LLM 推論とは何ですか?
4 つの段階: トークン化 → プレフィル → デコード → 非トークン化。主要な指標: TTFT、ITL、スループット。
0.1
LLM 推論が異なる理由
100倍のコストギャップについて説明しました。メモリ帯域幅の壁。従来の ML ルールが適用されない理由。
0.2
トランスの推論メカニズム
バイトレベルの注意のウォークスルー。 KV キャッシュの計算。 GQA/MQA と実数のトレードオフ。
パート II: GPU の基礎
章
タイトル
説明
1.1
GPUメモリ
HBM アーキテクチャ、メモリ階層、実稼働用の VRAM 予算設定。
1.2
ルーフラインモデル
コンピューティングとメモリ バウンドの分析。算術強度。パフォーマンス予測。
1.3
フラッシュアテンション
それが不可欠な理由、その仕組み、推論エンジンとの統合。
P

アート III: アテンションと KV キャッシュ
章
タイトル
説明
2.1
KV キャッシング
KV キャッシュが存在する理由。記憶の公式。成長パターンと限界。
2.2
注意メカニズム
MHA→MQA→GQAの進化。品質とメモリのトレードオフ。
2.3
ページ化された注意
KV キャッシュ用の仮想メモリ。 vLLM の画期的なイノベーション。
2.4
KV キャッシュの圧縮
量子化された KV、エビクション ポリシー、クロスリクエスト共有。
パート IV: 最適化テクニック
章
タイトル
説明
3.1
量子化
INT8、INT4、FP8 の詳細。それぞれに意味があるとき。品質への影響。
3.2
ターボクアント
FP4 と積極的な量子化のフロンティア。
3.3
連続バッチ処理
スループットのための動的なバッチ処理。反復レベルのスケジューリング。
3.4
投機的デコード
ドラフトモデル、検証、承認率。 2～3倍の高速化テクニック。
3.5
チャンクプレフィル
混合ワークロードのプリフィルとデコードのレイテンシのバランスをとります。
パート V: 推論エンジン
章
タイトル
説明
4.1
vLLM
アーキテクチャを深く掘り下げます。構成ガイド。プロダクションチューニング。
4.2
SGLang
RadixAttention、構造化出力、vLLM ではなく選択する場合。
4.3
TensorRT-LLM
NVIDIA の最適化されたランタイム。コンパイルのトレードオフ。ベストプラクティス。
パート VI: スケーリング
章
タイトル
説明
5.1
テンソル並列処理
モデルを GPU 間で分割します。コミュニケーションパターン。効率の限界。
5.2
MoE 推論
大規模な専門家によるルーティング。負荷分散。 DeepSeek-V2/V3 の洞察。
5.3
蒸留
効率的に提供するために大規模なモデルを圧縮します。
パート VII: 本番サービスの提供
章
タイトル
説明
6.1
レイサーブ
Ray を使用してスケーラブルな LLM サービスを構築します。
6.2
EKS + Kサーブ
AWS 上の Kubernetes ネイティブ LLM デプロイメント。
6.3
セージメーカー
SageMaker LMI によるマネージド推論。
6.4
細分化されたサービス
プリフィルとデコードを分離します。 40～60%のコスト削減。
6.5
コールドスタート
サーバーレス LLM の起動遅延を最小限に抑えます。
パート VIII: 操作

イオン
章
タイトル
説明
7.1
ベンチマーク
TTFT、ITL、スループット測定。ワークロードのリプレイ。
7.2
構造化された出力
JSON スキーマ、文法ガイドに基づいた生成、アウトライン。
7.3
エッジ展開
llama.cpp、GGUF、モバイル推論、Apple Silicon。
🧪研究室
各概念を強化するための実践的な演習。各ラボには、スターター コード、段階的な手順、およびソリューションが含まれています。
ハードウェア要件: ほとんどのラボは単一の GPU (g5.xlarge または同等) で実行されます。ラボ 06 と 08 にはマルチ GPU インスタンスが必要です。
LLM 推論を使用するときに常に使用する公式は次のとおりです。
理論上の最大デコード速度は、モデルの重みを読み取る速度によって制限されます。
秒あたりの最大トークン数 = メモリ帯域幅 / モデルサイズバイト
例: A100 (2 TB/秒) 上の Llama 8B (16GB FP16) → 最大 125 トークン/秒
キーと値のキャッシュに必要なメモリ:
kv_cache_bytes = 2 × num_layers × num_kv_heads × head_dim × seq_len × batch_size × dtype_bytes
例: ラマ 8B、バッチ=1、シーケンス=4096、FP16 → 512 MB
ワークロードがコンピューティング依存かメモリ依存かを決定します。
arithmetic_intensity = FLOPs / bytes_transferred
経験則: リッジポイントより下 (A100 で ~156 FLOPs/バイト) = メモリ制限
llm-大規模推論/
§── コンテンツ/ # 📖 ハンドブックの章
│ §── 00_foundations/ # パート I: 基礎
│ §── 01_gpu_fundamentals/ # パート II: GPU の基礎
│ §── 02_attention_and_kv/ # パート III: アテンションと KV キャッシュ
│ §── 03_optimization/ # パート IV: 最適化テクニック
│ §── 04_engines/ # パート V: 推論エンジン
│ §── 05_scaling/ # パート VI: スケーリング
│ §── 06_serving/ # パート VII: 生産サービング
│ §── 07_operations/ # 第 VIII 部: 作戦
│ └── utils/ # 可視化ユーティリティ
§── labs/ # 🧪 実践演習
§── 参照

ce/ # 📋 クイックリファレンス
│ §── cheat_sheet.md # 1 ページの概要
│ §── Glossary.md # 用語解説
│ §── vllm_quick_reference.md # vLLM コマンド
│ ━──cost_calculator.py # 推論コスト見積もり
§── アセット/ # 🎨 画像と図
└── スライド/ # 📊 プレゼンテーション資料
🎓 学習パス
今週 LLM を導入する必要があるエンジニアの場合:
0.0 LLM 推論とは何ですか? — 15分
0.1 LLM 推論が異なる理由 — 20 分
ラボ 04: vLLM の展開 — 45 分
推論インフラストラクチャを構築するエンジニア向け:
午前: パート I (基礎) + パート II (GPU の基礎)
午後: パート IV (最適化) + パート V (エンジン)
LLM サービスを標準化しているチームの場合:
1 日目: パート I、II、III — KV キャッシュによる基礎
2 日目: パート IV、V、VI — スケーリングによる最適化
3 日目: パート VII、VIII — 本番環境の提供と運用
この資料は次の場所で提供されています。
さらに多くの講演が予定されています — カンファレンスやミートアップでこれを希望する場合は、問題を開いてください。
この資料を研究または内部文書で使用する場合は、以下を引用してください。
@misc { 大規模な llm 推論 ,
title = { 大規模な LLM 推論: 実践者のハンドブック } ,
著者 = { ジャイナ教、ハーシュル } 、
年 = { 2025 } 、
URL = { https://github.com/harshuljain13/llm-inference-at-scale }
}
スターの歴史
これが役立つと思われる場合は、⭐リポジトリを利用してください。他の人がそれを発見するのに役立ちます。
貢献は大歓迎です。これは生きたドキュメントです。
エラーを修正 — タイプミス、古い情報、間違った数式
明確さの向上 — より良い説明、追加の例
コンテンツの追加 — 新しい章、ラボ、または参考資料
機能ブランチを作成します ( git checkout -b Improvement-kv-cache-chapter )
エラーを報告したり、修正を提案したりするには、GitHub Issue を開きます。
Harshul Jain はシニア ML インフラストラクチャです

彼は Audible (Amazon) のエンジニアとして、数百万の顧客にサービスを提供する GenAI セマンティック検索プラットフォームである ML Feature Store と、大規模なリアルタイム ストリーミング パイプラインを所有しています。彼は 4 年以上にわたって本番環境で ML インフラストラクチャを構築および運用しており、eMentoring プログラムを通じて 300 人以上のエンジニアを指導しています。
ニュースレター: エンジニア ダイジェスト — LLM 推論、詳しく説明
このハンドブックに記載されている見解、技術、意見は著者のみのものであり、Audible、Amazon、または関連組織の見解を表すものではありません。独自の、機密の、または内部の Amazon/Audible のシステム、データ、情報は含まれていません。すべてのコンテンツは、公開されている研究、オープンソース ツール、および著者の独立した経験と分析に基づいています。
このハンドブックは教育目的のみに提供されています。運用インフラストラクチャの決定は、特定のワークロード、ハードウェア、組織の制約に照らして検証する必要があります。著者は、ここに記載されている内容の正確性、完全性、または目的への適合性について、いかなる保証も行いません。
© 2026 ハーシュル・ジェイン。無断転載を禁じます。
著者からの事前の書面による許可がない限り、フレームワーク、図、モデル、用語、章の構造、または関連資料を含むこの作品の一部または全部を複製、配布、修正、翻案、または使用することはできません。これには、コース、トレーニングでの使用が含まれますが、これらに限定されません。

[切り捨てられた]

## Original Extract

A Practitioner handbook for production llm serving. - harshuljain13/llm-inference-at-scale

GitHub - harshuljain13/llm-inference-at-scale: A Practitioner handbook for production llm serving. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
harshuljain13
/
llm-inference-at-scale
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
harshuljain13/llm-inference-at-scale
master Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits assets assets content content reference reference slides slides .gitignore .gitignore .python-version .python-version LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
The definitive guide to serving large language models in production.
Quick Start •
Contents •
Labs •
Community •
Contributing
LLM inference is hard. Not "read the docs and figure it out" hard — fundamentally different from everything else in ML hard.
Traditional ML inference is a solved problem. You batch requests, run a forward pass, return results. Latency is predictable, memory is fixed, scaling is linear.
LLM inference breaks all of these assumptions:
Latency is unpredictable — a 10-token response takes 100ms, a 1000-token response takes 10 seconds
Memory grows during requests — the KV cache expands with every generated token
Scaling is sub-linear — communication overhead dominates as you add GPUs
Cost is 100x higher — $0.001/request becomes $0.10/request
This handbook exists because we needed it and couldn't find it. The knowledge is scattered across papers, blog posts, tribal knowledge, and source code comments. We've consolidated years of production experience and research into one comprehensive resource.
This is the guide we wish existed when we started.
📬 Follow the build — New chapters, explained in plain English with production context.
Subscribe to The Engineer's Digest to get notified
when new content drops. Subscribe free →
💬 Join the discussion
— questions, feedback, and corrections welcome
Why LLM inference costs 100x more than traditional ML
The memory bandwidth wall and how to work around it
Prefill vs decode: two completely different problems
KV cache mechanics at the byte level
Choosing between vLLM, SGLang, and TensorRT-LLM
Quantization tradeoffs (INT8, INT4, FP8, FP4)
Tensor parallelism and multi-GPU serving
Capacity planning and SLO management
PagedAttention and memory-efficient serving
Continuous batching for throughput
Speculative decoding for latency
FlashAttention and kernel-level optimizations
Disaggregated serving (prefill/decode separation)
MoE inference and expert routing
KV cache compression and eviction
Edge deployment with llama.cpp
git clone https://github.com/harshuljain13/llm-inference-at-scale.git
cd llm-inference-at-scale
# Create virtual environment
python -m venv .venv
source .venv/bin/activate # On Windows: .venv\Scripts\activate
# Install dependencies
pip install -r requirements.txt
Start Reading
# Open the first chapter
open content/00_foundations/00.0_what_is_llm_inference/what_is_llm_inference.md
Or browse the Table of Contents below.
Chapter
Title
Description
0.0
What is LLM Inference?
The four stages: tokenization → prefill → decode → detokenization. Key metrics: TTFT, ITL, throughput.
0.1
Why LLM Inference is Different
The 100x cost gap explained. Memory bandwidth wall. Why traditional ML rules don't apply.
0.2
Transformer Inference Mechanics
Byte-level attention walkthrough. KV cache math. GQA/MQA tradeoffs with real numbers.
Part II: GPU Fundamentals
Chapter
Title
Description
1.1
GPU Memory
HBM architecture, memory hierarchy, VRAM budgeting for production.
1.2
Roofline Model
Compute vs memory bound analysis. Arithmetic intensity. Performance prediction.
1.3
FlashAttention
Why it's essential, how it works, integration with inference engines.
Part III: Attention & KV Cache
Chapter
Title
Description
2.1
KV Caching
Why KV cache exists. Memory formulas. Growth patterns and limits.
2.2
Attention Mechanisms
MHA → MQA → GQA evolution. Quality vs memory tradeoffs.
2.3
PagedAttention
Virtual memory for KV cache. vLLM's breakthrough innovation.
2.4
KV Cache Compression
Quantized KV, eviction policies, cross-request sharing.
Part IV: Optimization Techniques
Chapter
Title
Description
3.1
Quantization
INT8, INT4, FP8 deep dive. When each makes sense. Quality impact.
3.2
TurboQuant
FP4 and the frontier of aggressive quantization.
3.3
Continuous Batching
Dynamic batching for throughput. Iteration-level scheduling.
3.4
Speculative Decoding
Draft models, verification, acceptance rates. 2-3x speedup techniques.
3.5
Chunked Prefill
Balancing prefill and decode latency for mixed workloads.
Part V: Inference Engines
Chapter
Title
Description
4.1
vLLM
Architecture deep dive. Configuration guide. Production tuning.
4.2
SGLang
RadixAttention, structured output, when to choose it over vLLM.
4.3
TensorRT-LLM
NVIDIA's optimized runtime. Compilation tradeoffs. Best practices.
Part VI: Scaling
Chapter
Title
Description
5.1
Tensor Parallelism
Splitting models across GPUs. Communication patterns. Efficiency limits.
5.2
MoE Inference
Expert routing at scale. Load balancing. DeepSeek-V2/V3 insights.
5.3
Distillation
Compressing large models for efficient serving.
Part VII: Production Serving
Chapter
Title
Description
6.1
Ray Serve
Building scalable LLM services with Ray.
6.2
EKS + KServe
Kubernetes-native LLM deployment on AWS.
6.3
SageMaker
Managed inference with SageMaker LMI.
6.4
Disaggregated Serving
Separating prefill and decode. 40-60% cost reduction.
6.5
Cold Start
Minimizing startup latency for serverless LLMs.
Part VIII: Operations
Chapter
Title
Description
7.1
Benchmarking
TTFT, ITL, throughput measurement. Workload replay.
7.2
Structured Output
JSON schemas, grammar-guided generation, Outlines.
7.3
Edge Deployment
llama.cpp, GGUF, mobile inference, Apple Silicon.
🧪 Labs
Hands-on exercises to reinforce each concept. Each lab includes starter code, step-by-step instructions, and solutions.
Hardware requirements: Most labs run on a single GPU (g5.xlarge or equivalent). Labs 06 and 08 require multi-GPU instances.
Formulas you'll use constantly when working with LLM inference:
The theoretical maximum decode speed, limited by how fast you can read model weights:
max_tokens_per_second = memory_bandwidth / model_size_bytes
Example: Llama 8B (16GB FP16) on A100 (2 TB/s) → 125 tokens/sec maximum
Memory required for the key-value cache:
kv_cache_bytes = 2 × num_layers × num_kv_heads × head_dim × seq_len × batch_size × dtype_bytes
Example: Llama 8B, batch=1, seq=4096, FP16 → 512 MB
Determines whether a workload is compute-bound or memory-bound:
arithmetic_intensity = FLOPs / bytes_transferred
Rule of thumb: Below the ridge point (~156 FLOPs/byte on A100) = memory-bound
llm-inference-at-scale/
├── content/ # 📖 Handbook chapters
│ ├── 00_foundations/ # Part I: Foundations
│ ├── 01_gpu_fundamentals/ # Part II: GPU Fundamentals
│ ├── 02_attention_and_kv/ # Part III: Attention & KV Cache
│ ├── 03_optimization/ # Part IV: Optimization Techniques
│ ├── 04_engines/ # Part V: Inference Engines
│ ├── 05_scaling/ # Part VI: Scaling
│ ├── 06_serving/ # Part VII: Production Serving
│ ├── 07_operations/ # Part VIII: Operations
│ └── utils/ # Visualization utilities
├── labs/ # 🧪 Hands-on exercises
├── reference/ # 📋 Quick references
│ ├── cheat_sheet.md # One-page summary
│ ├── glossary.md # Terminology
│ ├── vllm_quick_reference.md # vLLM commands
│ └── cost_calculator.py # Inference cost estimation
├── assets/ # 🎨 Images and diagrams
└── slides/ # 📊 Presentation materials
🎓 Learning Paths
For engineers who need to deploy an LLM this week:
0.0 What is LLM Inference? — 15 min
0.1 Why LLM Inference is Different — 20 min
Lab 04: vLLM Deployment — 45 min
For engineers building inference infrastructure:
Morning: Part I (Foundations) + Part II (GPU Fundamentals)
Afternoon: Part IV (Optimization) + Part V (Engines)
For teams standardizing on LLM serving:
Day 1: Parts I, II, III — Foundations through KV Cache
Day 2: Parts IV, V, VI — Optimization through Scaling
Day 3: Parts VII, VIII — Production Serving and Operations
This material has been presented at:
More talks coming — if you'd like this at your conference or meetup, open an issue.
If you use this material in research or internal documentation, please cite:
@misc { llm-inference-at-scale ,
title = { LLM Inference at Scale: A Practitioner's Handbook } ,
author = { Jain, Harshul } ,
year = { 2025 } ,
url = { https://github.com/harshuljain13/llm-inference-at-scale }
}
Star History
If you find this useful, please ⭐ the repo — it helps others discover it.
Contributions are welcome. This is a living document.
Fix errors — Typos, outdated information, incorrect formulas
Improve clarity — Better explanations, additional examples
Add content — New chapters, labs, or reference materials
Create a feature branch ( git checkout -b improve-kv-cache-chapter )
To report errors or suggest corrections, open a GitHub Issue.
Harshul Jain is a Senior ML Infrastructure Engineer at Audible (Amazon), where he owns the ML Feature Store, a GenAI semantic search platform serving millions of customers, and real-time streaming pipelines at scale. He has been building and operating ML infrastructure in production for 4+ years and mentors 300+ engineers through an eMentoring program.
Newsletter: The Engineer's Digest — LLM inference, deeply explained
The views, techniques, and opinions expressed in this handbook are solely those of the author and do not represent the views of Audible, Amazon, or any affiliated organization . No proprietary, confidential, or internal Amazon/Audible systems, data, or information has been included. All content is based on publicly available research, open-source tooling, and the author's independent experience and analysis.
This handbook is provided for educational purposes only . Production infrastructure decisions should be validated against your specific workload, hardware, and organizational constraints. The author makes no guarantees about the accuracy, completeness, or fitness for purpose of any content herein.
© 2026 Harshul Jain. All rights reserved.
No part of this work — including the framework, diagrams, models, terminology, chapter structure, or related materials — may be reproduced, distributed, modified, adapted, or used in whole or in part without prior written permission from the author. This includes but is not limited to use in courses, traini

[truncated]
