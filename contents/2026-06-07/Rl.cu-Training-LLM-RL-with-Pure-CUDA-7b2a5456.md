---
source: "https://github.com/KJLdefeated/RL.cu"
hn_url: "https://news.ycombinator.com/item?id=48431886"
title: "Rl.cu: Training LLM RL with Pure CUDA"
article_title: "GitHub - KJLdefeated/RL.cu: RLVR training for LLM in CUDA/C++ · GitHub"
author: "KJL0508"
captured_at: "2026-06-07T07:31:40Z"
capture_tool: "hn-digest"
hn_id: 48431886
score: 1
comments: 0
posted_at: "2026-06-07T04:45:22Z"
tags:
  - hacker-news
  - translated
---

# Rl.cu: Training LLM RL with Pure CUDA

- HN: [48431886](https://news.ycombinator.com/item?id=48431886)
- Source: [github.com](https://github.com/KJLdefeated/RL.cu)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T04:45:22Z

## Translation

タイトル: Rl.cu: Pure CUDA を使用した LLM RL のトレーニング
記事のタイトル: GitHub - KJLdefeated/RL.cu: CUDA/C++ での LLM のための RLVR トレーニング · GitHub
説明: CUDA/C++ での LLM の RLVR トレーニング。 GitHub でアカウントを作成して、KJLdefeated/RL.cu の開発に貢献してください。

記事本文:
GitHub - KJLdefeated/RL.cu: CUDA/C++ での LLM の RLVR トレーニング · GitHub
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
KJL敗北
/
RL.cu
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 開く

その他のアクション メニュー フォルダーとファイル
35 コミット 35 コミット .claude .claude 資産 アセット ドキュメント ドキュメント インクルード スクリプト スクリプト src src テスト テスト .gitignore .gitignore CMakeLists.txt CMakeLists.txt LICENSE LICENSE Makefile Makefile README.md README.md Agentic_rl_plan.md Agentic_rl_plan.md すべてのファイルを表示 リポジトリ ファイルナビゲーション
Pure CUDA での LLM 強化学習 — カーネルから GRPO まで
ゼロ PyTorch。完全な RL ループ。 vLLM を使用すると TRL よりも 1.37 倍高速になります。
クイックスタート •
カーネル •
推論エンジン •
トレーニング •
ベンチマーク •
アーキテクチャ •
ライセンス •
クロード・コード
DeepMath-103K (Qwen3-0.6B、同じ GPU) での GRPO トレーニング — 1.37 倍高速なウォールクロック、同等の報酬
完全な LLM RL パイプラインの最初からの実装 - 手書きの CUDA カーネル、継続的なバッチ処理を備えた vLLM スタイルの推論エンジン、および GRPO トレーニング。
レイヤー
何
CUDA カーネル
FlashAttendant-2、RMSNorm、RoPE、SwiGLU、Embedding、Sampler、AdamW、GRPO 損失 - すべて前方および後方パスを使用
モデル
Qwen3-0.6B フルフォワード+バックワードパス、セーフテンサー重量負荷
KVキャッシュ
ブロック マネージャーを使用したページ化された KV キャッシュ (vLLM と同じ設計)
推論エンジン
連続バッチ処理、CUDA グラフ キャプチャ、2 フェーズ スケジューリング
トレーニング
SFT + GRPO、勾配チェックポインティング、混合精度 AdamW
クイックスタート
Python 3.8+ (モデルの重みのダウンロード/データの準備のみ)
git clone https://github.com/KJLdefeated/RL.cu.git && cd RL.cu
# サードパーティのヘッダーと CUTLASS を取得します (下記の「サードパーティの依存関係」を参照)
mkdir -p include/ third_party third_party
curl -L -o include/third_party/json.hpp \
https://github.com/nlohmann/json/releases/download/v3.11.3/json.hpp
curl -L -o include/third_party/xxhash.h \
https://raw.githubusercontent.com/Cyan4973/xxhash/v0.8.1/xxhash.h
git clone -- Depth 1 https://github.com/N

VIDIA/cutlass.git third_party/cutlass
# モデルの重みをダウンロードする
pip インストール ハグフェイス_ハブ
python scripts/download_model.py Qwen/Qwen3-0.6B model_weights/Qwen3-0.6B
# すべてのターゲットをビルドします (デフォルト: sm_120; H100 などの場合は ARCH=90 でオーバーライドします)
作る
# または特定のターゲットをビルドする
make build/test_llmengine
サードパーティの依存関係
これらは、リポジトリにチェックインされないヘッダーのみ/ソースの依存関係です ( .gitignore にあります)。表示されているパスにそれらをドロップします。#include "third_party/..." と CMakeLists.txt がそれらを見つけることを期待する場所です。
上記の Build ブロック内のcurl / git clone コマンドは、3 つすべてをフェッチします。 CUTLASS は大きい (約 250 MB)。 --深さ 1 はクローンを浅く保ちます。
コマンド
説明
作る
すべてのターゲットを構成およびビルドする
make build/<ターゲット>
特定のバイナリをビルドする
test_rmsnorm を作成します
単一のテストをビルドして実行する
テストを行う
すべてのカーネルとモデルのテストをビルドして実行する
train_grpo を作成する
GRPO トレーニングの構築と実行 (デフォルトの引数)
train_sft を作成する
SFT トレーニングの構築と実行 (デフォルトの引数)
きれいにする
ビルドディレクトリを削除する
ARCH=90にする
特定の GPU アーチをターゲットにする (例: 80=A100、89=RTX4090、90=H100)
BUILD_TYPE=デバッグにする
デバッグビルド
推論の実行
# 正確性テスト + スループットベンチマーク
./build/test_llmengine model_weights/Qwen3-0.6B
GRPO トレーニング
# データセットを準備します (DeepMath-103K を JSONL としてダウンロードします)
pip インストール データセット
python scripts/prepare_data.py --mode grpo-text \
--dataset trl-lib/DeepMath-103K --output data/deepmath-103k.jsonl
# デフォルト設定で実行 (8 プロンプト x 8 gen、100 ステップ)
./build/train_grpo
# またはカスタマイズします (すべてのオプションについては --help を参照)
./build/train_grpo \
--model model_weights/Qwen3-0.6B \
--data データ/deepmath-103k.jsonl \
--バッチサイズ 64 --num-gens 8 \
--lr 1e-6 --合計ステップ 500 \
--save-dir チェックポイント/my_run
SFTトレーニング
# データセットを準備する
Python python_scripts/prepare_d

ata.py --mode sft --output data/sft_train.bin
# デフォルト設定で実行
./build/train_sft
# またはカスタマイズする
./build/train_sft \
--model model_weights/Qwen3-0.6B \
--data データ/sft_train.bin \
--バッチサイズ 8 --seq-len 2048 \
--lr 1e-5 --合計ステップ 5000
カーネル
すべてのカーネルは、FP16 I/O および FP32 蓄積を使用して書き込まれます。それぞれに、CPU/PyTorch リファレンスと比較するスタンドアロン テストがあります。
個別のカーネル テストを構築して実行します。
make build/test_attention && ./build/test_attention
make build/test_rmsnorm && ./build/test_rmsnorm
# ... など
推論エンジン
継続的なバッチ処理、ページ化された KV キャッシュ、および CUDA グラフ アクセラレーションを備えた vLLM スタイルのエンジン。
連続バッチ処理 — 2 フェーズのデコード + ステップごとのプレフィル、新しいリクエストはすぐに開始されます
ページ化された KV キャッシュ — ブロックレベルのメモリ管理、無駄な事前割り当てなし
CUDA グラフ — バケット サイズ (1、2、4、8、...、256) でキャプチャされたデコード
融合された投影 - QKV およびゲート + アップ投影が単一の GEMM に融合
トレーニング推論の不一致 = 0 - GRPO ポリシー比率には、log_probs_old (ロールアウト) と log_probs_new (トレーニング) が必要です。どちらも同じ qwen3_forward()、つまり同じカーネル、同じ FP 削減順序から来ています。
エンジンステップ:
schedule_decode() → 新しいシーケンスを事前入力 → サンプル → ポストプロセス
└── 継続的: 終了したスロットは待機中のリクエストによって即座に再利用されます
トレーニング
メモリを制御するためのチャンク化された lm_head を使用した標準のクロスエントロピー トレーニング。
GRPO (グループ相対ポリシー最適化)
各ステップについて:
1. 推論エンジンを使用してプロンプトごとに G 補完を生成します
2. 報酬関数によるスコア (例: ボックス回答マッチング)
3. GRPO の利点を計算する (グループ相対正規化)
4. 勾配チェックポイントを使用したフォワードパス
5. バックワードパス (レイヤーごとのアクティベーションを再計算)
6. 勾配累積 + クリッピングによる AdamW 更新
勾配チェックポイントi

ng はレイヤーごとの入力残差 + FlashAttendant LSE のみを保存し、他のすべてのアクティベーションをバックワード中に再計算します。
スリープ/ウェイクアップのライフサイクル: KV キャッシュ プールはトレーニング中に解放され、生成前に再割り当てされるため、推論フェーズとトレーニング フェーズの間で同じ GPU メモリが共有されます。
重みの転送は必要ありません: 推論フェーズとトレーニング フェーズは同じ重みを維持し、重みの転送を必要とせず、推論とトレーニングの不一致を回避します。
推論スループット (Qwen3-0.6B、RTX PRO 6000)
バッチサイズ
スループット (tok/秒)
256
6,963
nano-vllm スループットの 94% (7,411 tok/秒)。最適化の詳細については、docs/ENGINE.md を参照してください。
GRPO トレーニング (Qwen3-0.6B、8 プロンプト x 8 世代、DeepMath-103K)
同じタスク、同じ GPU (RTX PRO 6000) での RL.cu と TRL (vLLM を使用):
同じステップ数で RL.cu が 1.37 倍高速である理由 (実測値):
15% 高速な生成 (一致するトークン数で 2,992 対 2,602 トークン/秒) — CUDA グラフ、融合された QKV/ゲートアップ プロジェクション、およびゼロの Python オーバーヘッドにより、ステップごとの起動コストが削減されます
重みの転送なし — RL.cu は、重みを共有して同じプロセスで推論とトレーニングを実行します。 TRL は、ステップごとにトレーニング モデルと vLLM の推論コピーの間で重みを同期する必要があります。
トレーニングよりも短い完了数 — モデルが簡潔な答えを学習するにつれて、RL.cu の完了数は 1,889 → 968 トークンに縮小し、生成作業が最大 50% 削減されます。 TRL は全体を通じて約 1,840 トークンに留まります
RL.cu
§── src/kernels/ # 手書き CUDA カーネル (fwd + bwd)
│ §──attention.cu # FlashAttendant-2 with GQA
│ §── rmsnorm.cu # RMSNorm
│ §──rope.cu # ロータリー位置埋め込み (NeoX)
│ §── swiglu.cu # SwiGLU アクティベーション
│ §── embedding.cu # トークン埋め込み
│ §──sampler.cu # Top-k/p サンプリング
│ §── adamw.cu # 混合精度オプティマイザ
│ ━─ 。

..
§── src/model/
│ §── qwen3.cu # フルフォワード+バックワードパス(1,432行)
│ └── kv_cache.cu # ページ化された KV キャッシュ操作
§── include/engine/ # 推論エンジン
│ §── llm_engine.h # LLMEngine: トップレベル API
│ §──scheduler.h # 連続バッチ処理スケジューラ
│ §── block_manager.h # ページ化された KV ブロックの割り当て
│ └─ model_runner.cuh # モデル実行 + CUDA グラフ
§── include/training/ # トレーニングインフラストラクチャ
│ §── GRPO_trainer.h # GRPO トレーニングループ
│ §── SFT_trainer.h # SFT トレーニングループ
│ §── optimizer.h # フラットバッファーを使用した AdamW
│ └── lr_scheduler.h # コサイン + ウォームアップ
§── include/model/
│ §── tokenizer.h # BPE tokenizer (HF tokenizer.json を読み込みます)
│ §── config.h # モデル + エンジンコンフィグ
│ └──weights.h # Safetensors ローダー (mmap)
└── テスト/
§── kernels/ # 個々の CUDA カーネルの単体テスト
§── models/ # エンドツーエンドのモデル テスト (Qwen3、LLMEngine)
└── トレーニング/ # トレーニング ループ テスト (SFT、GRPO)
比較
llm.c
vLLM
TRL
RL.cu
言語
C/CUDA
Python + CUDA
Python + PyTorch
C++/CUDA
推論エンジン
-
はい
vLLM経由
はい
連続バッチ処理
-
はい
vLLM経由
はい
ページングされた KV キャッシュ
-
はい
vLLM経由
はい
CUDA グラフ
-
はい
-
はい
SFTトレーニング
はい (GPT-2)
-
はい
はい (Qwen3)
RL トレーニング (GRPO)
-
-
はい
はい
統合された推論 + トレーニング
該当なし
該当なし
ブリッジが必要です
はい
トレイン推論の不一致がゼロ
該当なし
該当なし
緩和が必要
設計上
実行時の依存関係
なし
Python + PyTorch
Python + PyTorch
なし
テスト
テストは 3 つのディレクトリに編成されています。
# すべてのテストを実行する
テストを行う
# 単一のテストを実行する
make test_attention # FlashAttendant-2 fwd+bwd
make test_llmengine # フルエンジン (11 個の統合テスト)
# Nsight Systems を使用してカーネルをプロファイリングする
nsys プロファイル --trace=c

uda,cublas --stats=true ./build/test_attention
完全なプロファイリング ガイドについては、tests/README.md を参照してください。
このプロジェクトは、Claude Code で動作するように構築されています。リポジトリにはプロジェクト レベルのドキュメントとスキルが含まれているため、クロードは最初のメッセージから効果的に貢献できます。
.claude/CLAUDE.md — クロードに完全なコンテキストを提供する 270 行のプロジェクト メモリ:
GPU アーキテクチャ (sm_120、CUDA 12.8)、モデル寸法 (Qwen3-0.6B)、ビルド システム
すべてのカーネルの API、グリッド/ブロック構成、および既知の落とし穴
スケジューラーの設計、KV キャッシュのレイアウト、カーネルの注意点
私たちが修正したすべてのバグとその理由 (クロードがそれらを再導入しないように)
コーディング規則: FP16 I/O、FP32 蓄積、 #pragma unroll 、テスト パターン
.claude/skills/add_new_kernel/ — 新しい CUDA カーネルを実装するための段階的なスキル:
ヘッダーは include/kernels/*.cuh 、ソースは src/kernels/*.cu 、テストは testing/test_*.cu
カーネル → ランチャー → テスト パターンを示す完全な例 (FlashAttendant-2、RMSNorm) が含まれています。
ワープレベルの削減、共有メモリのタイリング、ベクトル化されたロード、境界チェック
CPU リファレンス + 許容差比較テンプレート
# クロードコードをインストールする
npm install -g @anthropic-ai/claude-code
# 作業を開始します — クロードはすでにプロジェクトを知っています
cd RL.cu
クロード
# クロードがプロジェクト コンテキストでできることの例:
# "ページング アテンション用にフラッシュ デコード カーネルを追加"
# "なぜ私のアテンション カーネルは S > 16 に対して NaN を生成するのでしょうか?"
# "ワープ シャッフルを使用して RMSNorm 後方カーネルを最適化する"
# "線形レイヤーに INT8 量子化サポートを追加"
クロード・アンダースタ

[切り捨てられた]

## Original Extract

RLVR training for LLM in CUDA/C++. Contribute to KJLdefeated/RL.cu development by creating an account on GitHub.

GitHub - KJLdefeated/RL.cu: RLVR training for LLM in CUDA/C++ · GitHub
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
KJLdefeated
/
RL.cu
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .claude .claude assets assets docs docs include include scripts scripts src src tests tests .gitignore .gitignore CMakeLists.txt CMakeLists.txt LICENSE LICENSE Makefile Makefile README.md README.md agentic_rl_plan.md agentic_rl_plan.md View all files Repository files navigation
LLM Reinforcement Learning in Pure CUDA — From Kernels to GRPO
Zero PyTorch. Full RL loop. 1.37x faster than TRL with vLLM.
Quick Start •
Kernels •
Inference Engine •
Training •
Benchmarks •
Architecture •
License •
Claude Code
GRPO training on DeepMath-103K (Qwen3-0.6B, same GPU) — 1.37x faster wall-clock, matching reward
A from-scratch implementation of the complete LLM RL pipeline — hand-written CUDA kernels, a vLLM-style inference engine with continuous batching, and GRPO training.
Layer
What
CUDA Kernels
FlashAttention-2, RMSNorm, RoPE, SwiGLU, Embedding, Sampler, AdamW, GRPO loss — all with forward AND backward passes
Model
Qwen3-0.6B full forward + backward pass, safetensors weight loading
KV Cache
Paged KV cache with block manager (same design as vLLM)
Inference Engine
Continuous batching, CUDA graph capture, two-phase scheduling
Training
SFT + GRPO with gradient checkpointing, mixed-precision AdamW
Quick Start
Python 3.8+ (only for downloading model weights / preparing data)
git clone https://github.com/KJLdefeated/RL.cu.git && cd RL.cu
# Fetch third-party headers and CUTLASS (see "Third-party dependencies" below)
mkdir -p include/third_party third_party
curl -L -o include/third_party/json.hpp \
https://github.com/nlohmann/json/releases/download/v3.11.3/json.hpp
curl -L -o include/third_party/xxhash.h \
https://raw.githubusercontent.com/Cyan4973/xxHash/v0.8.1/xxhash.h
git clone --depth 1 https://github.com/NVIDIA/cutlass.git third_party/cutlass
# Download model weights
pip install huggingface_hub
python scripts/download_model.py Qwen/Qwen3-0.6B model_weights/Qwen3-0.6B
# Build all targets (default: sm_120; override with ARCH=90 for H100, etc.)
make
# Or build a specific target
make build/test_llmengine
Third-party dependencies
These are header-only / source dependencies that aren't checked into the repo (they're in .gitignore ). Drop them at the paths shown — that's where #include "third_party/..." and CMakeLists.txt expect to find them.
The curl / git clone commands in the Build block above fetch all three. CUTLASS is large (~250 MB); --depth 1 keeps the clone shallow.
Command
Description
make
Configure + build all targets
make build/<target>
Build a specific binary
make test_rmsnorm
Build + run a single test
make tests
Build + run all kernel & model tests
make train_grpo
Build + run GRPO training (default args)
make train_sft
Build + run SFT training (default args)
make clean
Remove build directory
make ARCH=90
Target a specific GPU arch (e.g. 80=A100, 89=RTX4090, 90=H100)
make BUILD_TYPE=Debug
Debug build
Run Inference
# Correctness tests + throughput benchmark
./build/test_llmengine model_weights/Qwen3-0.6B
GRPO Training
# Prepare dataset (downloads DeepMath-103K as JSONL)
pip install datasets
python scripts/prepare_data.py --mode grpo-text \
--dataset trl-lib/DeepMath-103K --output data/deepmath-103k.jsonl
# Run with default settings (8 prompts x 8 gens, 100 steps)
./build/train_grpo
# Or customize (see --help for all options)
./build/train_grpo \
--model model_weights/Qwen3-0.6B \
--data data/deepmath-103k.jsonl \
--batch-size 64 --num-gens 8 \
--lr 1e-6 --total-steps 500 \
--save-dir checkpoints/my_run
SFT Training
# Prepare dataset
python python_scripts/prepare_data.py --mode sft --output data/sft_train.bin
# Run with default settings
./build/train_sft
# Or customize
./build/train_sft \
--model model_weights/Qwen3-0.6B \
--data data/sft_train.bin \
--batch-size 8 --seq-len 2048 \
--lr 1e-5 --total-steps 5000
Kernels
Every kernel is written with FP16 I/O and FP32 accumulation. Each has a standalone test comparing against a CPU/PyTorch reference.
Build and run any individual kernel test:
make build/test_attention && ./build/test_attention
make build/test_rmsnorm && ./build/test_rmsnorm
# ... etc
Inference Engine
A vLLM-style engine with continuous batching, paged KV cache, and CUDA graph acceleration.
Continuous batching — two-phase decode + prefill per step, new requests start immediately
Paged KV cache — block-level memory management, no wasted pre-allocation
CUDA graphs — decode captured at bucket sizes (1, 2, 4, 8, ..., 256)
Fused projections — QKV and gate+up projections fused into single GEMMs
Train-inference mismatch = 0 - The GRPO policy ratio needs log_probs_old (rollout) and log_probs_new (training). Both come from the same qwen3_forward() — same kernels, same FP reduction order.
Engine step:
schedule_decode() → prefill new seqs → sample → postprocess
└── continuous: finished slots instantly reused by waiting requests
Training
Standard cross-entropy training with chunked lm_head to control memory.
GRPO (Group Relative Policy Optimization)
For each step:
1. Generate G completions per prompt using the inference engine
2. Score with reward function (e.g., boxed-answer matching)
3. Compute GRPO advantages (group-relative normalization)
4. Forward pass with gradient checkpointing
5. Backward pass (recompute activations per layer)
6. AdamW update with gradient accumulation + clipping
Gradient checkpointing saves only per-layer input residuals + FlashAttention LSE, recomputing all other activations during backward.
Sleep/wakeup lifecycle: KV cache pools are freed during training and re-allocated before generation, so the same GPU memory is shared between inference and training phases.
No weight transfer needed: Inference phase and training phase maitain same weight, don't need weight transfer and avoid of inference-training mismatch.
Inference Throughput (Qwen3-0.6B, RTX PRO 6000)
Batch Size
Throughput (tok/s)
256
6,963
94% of nano-vllm throughput (7,411 tok/s). See docs/ENGINE.md for the full optimization journey.
GRPO Training (Qwen3-0.6B, 8 prompts x 8 generations, DeepMath-103K)
RL.cu vs TRL (w/ vLLM) on the same task, same GPU (RTX PRO 6000):
Why RL.cu is 1.37x faster (wall-clock) for the same number of steps:
15% faster generation (2,992 vs 2,602 tok/s at matched token counts) — CUDA graphs, fused QKV/gate-up projections, and zero Python overhead eliminate per-step launch costs
No weight transfer — RL.cu runs inference and training in the same process with shared weights; TRL must sync weights between the training model and vLLM's inference copy every step
Shorter completions over training — RL.cu's completions shrink from 1,889 → 968 tokens as the model learns concise answers, reducing generation work by ~50%; TRL stays at ~1,840 tokens throughout
RL.cu
├── src/kernels/ # Hand-written CUDA kernels (fwd + bwd)
│ ├── attention.cu # FlashAttention-2 with GQA
│ ├── rmsnorm.cu # RMSNorm
│ ├── rope.cu # Rotary Position Embedding (NeoX)
│ ├── swiglu.cu # SwiGLU activation
│ ├── embedding.cu # Token embedding
│ ├── sampler.cu # Top-k/p sampling
│ ├── adamw.cu # Mixed-precision optimizer
│ └── ...
├── src/model/
│ ├── qwen3.cu # Full forward + backward pass (1,432 lines)
│ └── kv_cache.cu # Paged KV cache operations
├── include/engine/ # Inference engine
│ ├── llm_engine.h # LLMEngine: top-level API
│ ├── scheduler.h # Continuous batching scheduler
│ ├── block_manager.h # Paged KV block allocation
│ └── model_runner.cuh # Model execution + CUDA graphs
├── include/training/ # Training infrastructure
│ ├── GRPO_trainer.h # GRPO training loop
│ ├── SFT_trainer.h # SFT training loop
│ ├── optimizer.h # AdamW with flat buffer
│ └── lr_scheduler.h # Cosine + warmup
├── include/model/
│ ├── tokenizer.h # BPE tokenizer (reads HF tokenizer.json)
│ ├── config.h # Model + engine config
│ └── weights.h # Safetensors loader (mmap)
└── tests/
├── kernels/ # Unit tests for individual CUDA kernels
├── models/ # End-to-end model tests (Qwen3, LLMEngine)
└── training/ # Training loop tests (SFT, GRPO)
Comparison
llm.c
vLLM
TRL
RL.cu
Language
C/CUDA
Python + CUDA
Python + PyTorch
C++/CUDA
Inference engine
-
Yes
via vLLM
Yes
Continuous batching
-
Yes
via vLLM
Yes
Paged KV cache
-
Yes
via vLLM
Yes
CUDA graphs
-
Yes
-
Yes
SFT training
Yes (GPT-2)
-
Yes
Yes (Qwen3)
RL training (GRPO)
-
-
Yes
Yes
Unified inference + training
N/A
N/A
Bridge needed
Yes
Zero train-infer mismatch
N/A
N/A
Requires mitigation
By design
Runtime dependencies
None
Python + PyTorch
Python + PyTorch
None
Tests
Tests are organized into three directories:
# Run all tests
make tests
# Run a single test
make test_attention # FlashAttention-2 fwd+bwd
make test_llmengine # Full engine (11 integration tests)
# Profile a kernel with Nsight Systems
nsys profile --trace=cuda,cublas --stats=true ./build/test_attention
See tests/README.md for full profiling guide.
This project is built to work with Claude Code . The repo includes project-level docs and skills so Claude can contribute effectively from the first message.
.claude/CLAUDE.md — 270-line project memory that gives Claude full context:
GPU architecture (sm_120, CUDA 12.8), model dimensions (Qwen3-0.6B), build system
Every kernel's API, grid/block config, and known pitfalls
Scheduler design, KV cache layout, attention kernel gotchas
All bugs we've fixed and why (so Claude doesn't reintroduce them)
Coding conventions: FP16 I/O, FP32 accumulation, #pragma unroll , test patterns
.claude/skills/add_new_kernel/ — Step-by-step skill for implementing new CUDA kernels:
Header in include/kernels/*.cuh , source in src/kernels/*.cu , test in tests/test_*.cu
Includes full examples (FlashAttention-2, RMSNorm) showing the kernel → launcher → test pattern
Warp-level reduction, shared memory tiling, vectorized loads, bounds checking
CPU reference + tolerance comparison template
# Install Claude Code
npm install -g @anthropic-ai/claude-code
# Start working — Claude already knows the project
cd RL.cu
claude
# Examples of what Claude can do with the project context:
# "Add a flash decoding kernel for paged attention"
# "Why is my attention kernel producing NaN for S > 16?"
# "Optimize the RMSNorm backward kernel with warp shuffle"
# "Add INT8 quantization support for linear layers"
Claude understa

[truncated]
