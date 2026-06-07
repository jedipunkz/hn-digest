---
source: "https://local-llm.utop.workers.dev/"
hn_url: "https://news.ycombinator.com/item?id=48433268"
title: "Show HN: Best setup local LLM found for a 5090 (llama.cpp fork + turboquant)"
article_title: "Running Qwen 35B MoE at 450k Context on a Single 32GB GPU"
author: "utopman"
captured_at: "2026-06-07T10:03:27Z"
capture_tool: "hn-digest"
hn_id: 48433268
score: 1
comments: 0
posted_at: "2026-06-07T09:36:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Best setup local LLM found for a 5090 (llama.cpp fork + turboquant)

- HN: [48433268](https://news.ycombinator.com/item?id=48433268)
- Source: [local-llm.utop.workers.dev](https://local-llm.utop.workers.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T09:36:02Z

## Translation

タイトル: HN を表示: 5090 で見つかった最良のセットアップ ローカル LLM (llama.cpp フォーク + Turboquant)
記事のタイトル: 単一の 32GB GPU で 450k コンテキストで Qwen 35B MoE を実行する
説明: llama.cpp、YaRN スケーリング、および TurboQuant を使用して、単一の 32GB VRAM NVIDIA RTX 5090 上で 450,000 トークンのコンテキスト ウィンドウで Qwen 3.6 35B Mixture of Experts (MoE) を実行するための技術ガイド。
HN テキスト: こんにちは皆さん。
私はこのセットアップを conSummer ハードウェアで見つけましたが、ローカル ハードウェアで素晴らしい結果が得られるようです。
- クウェン 3.6 q6
- Turboquant Turbo3 モード llama.cpp フォークを使用した 450 K コンテキスト
- マルチモーダル サポート この AI によって生成されたブログ記事は、私が何をどのように行ったか、および結果の例についての一種の「レポート」です。これが一部の人々にとって役立つことを願っています。注: 私はこの記事が成功するかどうかにはあまり興味がありません。主に、5090 の興味深い使用法だと思うことを共有したいと思っています。私は AI に hn 「ルール」に準拠し、事実を保つように指示するブログ ページを生成しました。これは間違いなく完璧ではなく、かなり短時間で完了し、265K のコンテキストで適切にテストされていません。私の怠惰を許してください:)。私は今、5090 で何ができるかに熱中しています。

記事本文:
← コントロールセンターに戻る
ℹ
注: この記事はすべて AI (Antigravity) によって生成されました。
単一の 32GB GPU で 450k コンテキストで Qwen 35B MoE を実行
このレポートでは、 llama.cpp を使用して単一の 32GB VRAM GPU (NVIDIA RTX 5090) 上で 450,000 トークンの拡張コンテキスト ウィンドウで Qwen 3.6 35B MoE (35B-A3B) モデルを実行するために必要な構成とメモリ キャリブレーションについて詳しく説明します。
このセットアップは完全に実装され、デュアル ブートのオーバーヘッドを回避するために RTX 5090 ワークステーションの単一ホスト オペレーティング システムとして選択された Windows で実行されます。このガイドでは、Windows ネイティブ スクリプト (バッチおよび PowerShell) とコンパイル済み DLL 管理に焦点を当てていますが、コアの実行パラメータとパフォーマンス調整は、パスを若干調整するだけで Linux 展開にも完全に適用できます。
1. モデルの選択と量子化のトレードオフ
このセットアップで選択したモデルは Qwen3.6-35B-A3B-Q6_K.gguf (28.5GB) です。
アーキテクチャ (MoE) : Qwen 3.5/3.6 MoE は、専門家混合アーキテクチャを利用しており、合計 350 億のうち、トークンごとにアクティブになるパラメータは約 30 億のみです。これにより、大規模モデルの推論機能を維持しながら、小規模モデルに匹敵する実行速度が維持されます。
量子化 (Q6_K) : Q4_K_M または Q5_K_S は VRAM フットプリントをさらに削減しますが、複雑なコード生成の論理精度と構文構造を維持し、ネイティブ BF16 の複雑さの 99% 以上を維持するために Q6_K が選択されました。
2. メモリバジェットとコンテキストの調整
32 GB の物理 VRAM (nvidia-smi 経由で表示される 32,607 MiB) を搭載した GPU では、メモリ バジェットが制限されます。
モデル重量 (Q6_K) : 28.5GB。
静的オーバーヘッドと OS: ~0.8GB。
コンテキストの VRAM マージン: ~2.7GB。
A. TurboQuant による KV キャッシュ量子化 (turbo3)
TurboQuant 形式を使用して、KV キャッシュ テンソルを 3 ビット精度に圧縮します。

s:
この圧縮により、KV キャッシュのフットプリントが約 80% 削減され、450k のコンテキスト ウィンドウが残りの 2.7GB VRAM バッファ内に収まるようになります。キャッシュを 3 ビットに量子化すると、小規模ではあるが測定可能なパープレキシティの低下が生じます。
B. YaRN による RoPE スケーリングとパープレキシティの低下
Qwen 3.5/3.6 MoE のネイティブ コンテキストの長さは 262,144 トークンです。コンテキスト ウィンドウを 450,000 トークンに拡張するには、YaRN (Yet another RoPE extensioN) スケーリングを適用します。
--rope-scaling 糸 --yarn-orig-ctx 262144 --rope-scale 1.72
3. マルチモーダル (ビジョン) プロジェクターのセットアップ
画像入力をサポートするには、モデルには対応するビジョン プロジェクター ( mmproj ) が必要です。
1. 対応するプロジェクターをダウンロードします: Qwen3.6-35B-A3B-mmproj-F16.gguf (899MB)。
2. --mmproj 引数を使用してサーバーにロードします。
llama.cpp にコンパイルされたイメージ デコーダは、stb_image ライブラリを使用します。 PNG および JPEG 形式をサポートしますが、WebP はサポートしません。 API に送信された WebP ファイルは、400 Bad Request (画像または音声ファイルのロードに失敗しました) エラーを返します。
4. テクニカル レプリケーション ガイド
ステップ 1: 安全なモデルの取得
このスクリプトは、ローカル トークンを使用して、Hugging Face Hub から GGUF モデルとそのビジョン プロジェクターを取得します。
# download_model.py
OSをインポートする
ハギングフェイスハブからインポート hf_hub_download
REPO_ID = "jimbothigpen/Qwen3.6-35B-A3B-GGUF"
FILES = ["Qwen3.6-35B-A3B-Q6_K.gguf", "Qwen3.6-35B-A3B-mmproj-F16.gguf"]
LOCAL_DIR = "./モデル"
TOKEN = os.environ.get("HF_TOKEN")
FILES のファイル名:
hf_hub_download(repo_id=REPO_ID、filename=ファイル名、local_dir=LOCAL_DIR、token=TOKEN)
ステップ 2: サーバー呼び出しの引数
次のフラグを使用して llama-server プロセスを開始します。
./llama-server.exe ^
-m "./models/Qwen3.6-35B-A3B-Q6_K.gguf" ^
--mmproj "./models/Qwen3.6-35B-A3B-mmproj-F16.gguf" ^
--mmap なし ^
--ポート 9000 ^
--ホスト 0.0.0.0 ^
-ngl 99

^
-c 450000 ^
--ロープスケール糸 ^
--yarn-orig-ctx 262144 ^
--ロープスケール 1.72 ^
--キャッシュ-タイプ-K ターボ3 ^
--cache-type-v ターボ3 ^
--flash-attn オン ^
-b 512 ^
-ub 512
(Linux では、 ^ を \ に置き換えます)。
--no-mmap : メモリ マッピングを無効にし、すべてのモデルの重みを強制的に VRAM に連続的にロードします。これにより、推論中の OS ページングのボトルネックが回避されます。
--flash-attn : Flash アテンションをアクティブにして、メモリのオーバーヘッドを削減し、互換性のある NVIDIA アーキテクチャでの計算を高速化します。
-ngl 99 : 40 のモデル レイヤーすべてと出力テンソルを GPU にオフロードします。
5. 自動化された VRAM ライフサイクル管理 (Go Wrapper)
モデルとコンテキストは GPU メモリのほぼ 100% を占有するため、サーバーをバックグラウンドで実行したままにすると、他のアプリケーション (グラフィックス ドライバーや Web ブラウザーなど) に重大な VRAM ボトルネックが発生します。
これを解決するために、Windows サービス (または Linux systemd デーモン) として実行される軽量の Go (Gin Gonic) マネージャーを実装しました。
オンデマンド ライフサイクル : Web ダッシュボードを公開して、必要に応じてサーバーを起動し、完全に停止して (プロセスを強制終了して) VRAM を 100% 即座に解放します。
ステータス ポーリング : TCP 経由でローカル ポート 9000 に 3 秒ごとに ping を実行して、HTTP サーバーがいつリッスンしているかを確認し、リアルタイムでログを解析して GPU オフロード ステータスを確認します。
6. パフォーマンス評価とコード生成
この構成を使用すると、モデルは大規模なコードベースまたはログをダイジェストし、複雑なスクリプトをワンショットで生成できます。
テスト ケースとして、モデルに次のプロンプトが与えられました。
「Web 上で見つけたクールなテクスチャのビデオ ゲーム キャラクター (マリオなど) を表示する Three.js アプリを使用して、index.html を書いてください。単一ページ。」
このモデルは、キャンバスで生成されたマリオ、ハテナ ブロック、パイプのテクスチャを使用して、完全なインタラクティブな 3D シーンを生成することに成功しました。コードがコンパイルされて表示される

オービット コントロールと基本的なアニメーション ループを使用して、ワンショットで作成できます。
👉 生成されたアプリケーションをライブで表示およびテストします (mario.html)
7. クライアント統合 (OpenCode)
このローカル エンドポイントをマルチモーダル機能を備えた開発アシスタント (OpenCode など) として使用するには、ローカルの opencode.json 構成ファイルで「modalities」ブロックを構成します。
"ラマックップ": {
"npm": "@ai-sdk/openai-互換",
"名前": "llamacpp",
"オプション": {
"baseURL": "http://127.0.0.1:9000/v1"
}、
「モデル」: {
"qwen3.6-35b-q6-450k": {
"名前": "qwen3.6-35b-q6-450k",
「制限」: {
「コンテキスト」: 450000、
「入力」: 450000、
「出力」: 8192
}、
「モダリティ」: {
"入力": ["テキスト", "画像"],
"出力": ["テキスト"]
}
}
}
}
「モダリティ」を定義すると、OpenCode クライアントはモデルが画像入力をサポートしていることを認識できるようになり、エディターのチャット インターフェイス内で画像のドラッグ アンド ドロップのサポートがロック解除されます。
コンシューマ ハードウェア上の 450k コンテキスト ウィンドウで 35B の専門家混合モデルを実行すると、ローカル推論の最適化がどこまで進歩したかがわかります。 llama.cpp 、 TurboQuant (3 ビット KV キャッシュ) 、および YaRN スケーリング を組み合わせることで、単一の 32GB VRAM GPU で非常に大きなプロンプトを処理できます。
ただし、この設定は 32 GB フレーム バッファの絶対的な物理制限で動作するため、動的メモリ割り当てのマージンは事実上ありません。さらに重要なことは、モデルのネイティブ 262k トークン制限を超えて YaRN スケーリングによってコンテキストを拡張すると、検索の忠実性と論理的推論の精度が大幅に低下することをユーザーが常に認識しておく必要があることです。クリティカルなワークロードの場合、コンテキスト サイズをネイティブの境界内に維持することが最も信頼性の高い方法ですが、450k の制限は探索的な検索や広範な要約タスク用に確保するのが最適です。

## Original Extract

Technical guide on running Qwen 3.6 35B Mixture of Experts (MoE) at a 450,000 token context window on a single 32GB VRAM NVIDIA RTX 5090 using llama.cpp, YaRN scaling, and TurboQuant.

Hi folks,
I found this setup on consummer hardware that seems to have great results on local hardware.
- qwen 3.6 q6
- 450 K context using turboquant turbo3 mode llama.cpp fork
- multimodal support This AI generated blog article is a kind of "report" of what and how I did and result exemples. I hope this can be usefull to some peopole. Note : I am not much intersted in having success with this article, I mainly want to share what I think is an interesting use of a 5090. I generated the blog page telling AI to be compliant with hn "rules" and remain factual. It's definitely not perfect, done rather quickly, not properly tested over 265K context. please forgive my lazyness :) . I am just enthousiast right now about what can be done on a 5090.

← Back to Control Center
ℹ
Note: This article was fully generated by an AI (Antigravity).
Running Qwen 35B MoE at 450k Context on a Single 32GB GPU
This report details the configuration and memory calibration required to run the Qwen 3.6 35B MoE (35B-A3B) model at an extended context window of 450,000 tokens on a single 32GB VRAM GPU (NVIDIA RTX 5090) using llama.cpp .
This setup is fully implemented and running under Windows , which was chosen as the single host operating system for the RTX 5090 workstation to avoid dual-boot overhead. The guide focuses on Windows-native scripts (Batch and PowerShell) and compiled DLL management, but the core execution parameters and performance calibrations remain fully applicable to Linux deployments with minor path adjustments.
1. Model Selection & Quantization Trade-offs
The model selected for this setup is Qwen3.6-35B-A3B-Q6_K.gguf (28.5GB).
Architecture (MoE) : Qwen 3.5/3.6 MoE utilizes a Mixture of Experts architecture where only ~3 billion parameters are active per token out of 35 billion total. This maintains execution speeds comparable to small models while retaining the reasoning capabilities of larger models.
Quantization (Q6_K) : While Q4_K_M or Q5_K_S would reduce the VRAM footprint further, Q6_K was selected to preserve logical accuracy and syntax structure for complex code generation, maintaining over 99% of the native BF16 perplexity.
2. Memory Budget & Context Calibration
On a GPU with 32GB of physical VRAM (32,607 MiB visible via nvidia-smi ), the memory budget is constrained:
Model Weights (Q6_K) : 28.5GB.
Static Overhead & OS : ~0.8GB.
VRAM Margin for Context : ~2.7GB.
A. KV Cache Quantization via TurboQuant ( turbo3 )
We compress the KV cache tensors to 3-bit precision using the TurboQuant formats:
This compression reduces the KV cache footprint by approximately 80%, allowing the 450k context window to fit within the remaining 2.7GB VRAM buffer. Quantizing the cache to 3-bit introduces a minor but measurable perplexity degradation.
B. RoPE Scaling via YaRN and Perplexity Degradation
The native context length of Qwen 3.5/3.6 MoE is 262,144 tokens. To extend the context window to 450,000 tokens, we apply YaRN (Yet another RoPE extensioN) scaling:
--rope-scaling yarn --yarn-orig-ctx 262144 --rope-scale 1.72
3. Multimodal (Vision) Projector Setup
To support image inputs, the model requires its corresponding vision projector ( mmproj ):
1. Download the matching projector: Qwen3.6-35B-A3B-mmproj-F16.gguf (899MB).
2. Load it in the server using the --mmproj argument.
The image decoder compiled into llama.cpp uses the stb_image library. It supports PNG and JPEG formats, but does not support WebP. WebP files sent to the API return a 400 Bad Request (Failed to load image or audio file) error.
4. Technical Replication Guide
Step 1: Secure Model Acquisition
This script retrieves the GGUF model and its vision projector from Hugging Face Hub using your local token:
# download_model.py
import os
from huggingface_hub import hf_hub_download
REPO_ID = "jimbothigpen/Qwen3.6-35B-A3B-GGUF"
FILES = ["Qwen3.6-35B-A3B-Q6_K.gguf", "Qwen3.6-35B-A3B-mmproj-F16.gguf"]
LOCAL_DIR = "./models"
TOKEN = os.environ.get("HF_TOKEN")
for filename in FILES:
hf_hub_download(repo_id=REPO_ID, filename=filename, local_dir=LOCAL_DIR, token=TOKEN)
Step 2: Server Invocation Arguments
Start the llama-server process with the following flags:
./llama-server.exe ^
-m "./models/Qwen3.6-35B-A3B-Q6_K.gguf" ^
--mmproj "./models/Qwen3.6-35B-A3B-mmproj-F16.gguf" ^
--no-mmap ^
--port 9000 ^
--host 0.0.0.0 ^
-ngl 99 ^
-c 450000 ^
--rope-scaling yarn ^
--yarn-orig-ctx 262144 ^
--rope-scale 1.72 ^
--cache-type-k turbo3 ^
--cache-type-v turbo3 ^
--flash-attn on ^
-b 512 ^
-ub 512
(On Linux, replace ^ with \ ).
--no-mmap : Disables memory-mapping, forcing all model weights to load contiguously into VRAM. This prevents OS paging bottlenecks during inference.
--flash-attn : Activates Flash Attention, reducing memory overhead and accelerating computations on compatible NVIDIA architectures.
-ngl 99 : Offloads all 40 model layers plus the output tensor to the GPU.
5. Automated VRAM Lifecycle Management (Go Wrapper)
Since the model and context occupy nearly 100% of the GPU memory, leaving the server running in the background causes severe VRAM bottlenecks for other applications (such as graphics drivers or web browsers).
To solve this, we implemented a lightweight Go (Gin Gonic) manager that runs as a Windows Service (or Linux systemd daemon):
On-Demand Lifecycle : Exposes a web dashboard to start the server when needed and stop it cleanly (killing the process) to release 100% of VRAM instantly.
Status Polling : Pings the local port 9000 via TCP every 3 seconds to confirm when the HTTP server is listening, and parses logs in real-time to check GPU offloading status.
6. Performance Evaluation & Code Generation
With this configuration, the model can digest large codebases or logs, and generate complex scripts in a single shot.
As a test case, the model was given the following prompt:
"write me an index.html with a Three.js app that displays a cool textured video game character (mario or other) that you find on the web. single page"
The model successfully generated a complete, interactive 3D scene using Canvas-generated textures for Mario, question blocks, and pipes. The code compiled and rendered in one shot with orbit controls and basic animation loops.
👉 View and test the generated application live (mario.html)
7. Client Integration (OpenCode)
To use this local endpoint as a development assistant (e.g. in OpenCode) with multimodal capabilities, configure the "modalities" block in your local opencode.json configuration file:
"llamacpp": {
"npm": "@ai-sdk/openai-compatible",
"name": "llamacpp",
"options": {
"baseURL": "http://127.0.0.1:9000/v1"
},
"models": {
"qwen3.6-35b-q6-450k": {
"name": "qwen3.6-35b-q6-450k",
"limit": {
"context": 450000,
"input": 450000,
"output": 8192
},
"modalities": {
"input": ["text", "image"],
"output": ["text"]
}
}
}
}
Defining "modalities" allows the OpenCode client to recognize that the model supports image inputs, unlocking image drag-and-drop support inside the editor chat interface.
Running a 35B Mixture of Experts model at a 450k context window on consumer hardware is a demonstration of how far local inference optimizations have progressed. By combining llama.cpp , TurboQuant (3-bit KV cache) , and YaRN scaling , a single 32GB VRAM GPU can handle extremely large prompts.
However, this setup operates at the absolute physical limits of a 32GB frame buffer, leaving virtually no margin for dynamic memory allocation. More importantly, users must remain aware that expanding context via YaRN scaling beyond the model's native 262k token limit comes with a significant compromise in retrieval fidelity and logical reasoning accuracy. For critical workloads, keeping context sizes within the native bounds remains the most reliable path, while the 450k limit is best reserved for exploratory search and broad summarization tasks.
