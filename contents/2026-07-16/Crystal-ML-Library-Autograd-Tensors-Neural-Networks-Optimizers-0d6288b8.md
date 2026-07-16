---
source: "https://github.com/skuznetsov/cogni-ml"
hn_url: "https://news.ycombinator.com/item?id=48930154"
title: "Crystal ML Library: Autograd, Tensors, Neural Networks, Optimizers"
article_title: "GitHub - skuznetsov/cogni-ml: Crystal ML library: Autograd, Tensors, Neural Networks, Optimizers · GitHub"
author: "zX41ZdbW"
captured_at: "2026-07-16T04:52:25Z"
capture_tool: "hn-digest"
hn_id: 48930154
score: 3
comments: 0
posted_at: "2026-07-16T03:52:11Z"
tags:
  - hacker-news
  - translated
---

# Crystal ML Library: Autograd, Tensors, Neural Networks, Optimizers

- HN: [48930154](https://news.ycombinator.com/item?id=48930154)
- Source: [github.com](https://github.com/skuznetsov/cogni-ml)
- Score: 3
- Comments: 0
- Posted: 2026-07-16T03:52:11Z

## Translation

タイトル: Crystal ML ライブラリ: Autograd、Tensors、ニューラル ネットワーク、オプティマイザー
記事のタイトル: GitHub - skuznetsov/cogni-ml: Crystal ML ライブラリ: Autograd、Tensors、Neural Networks、Optimizers · GitHub
説明: Crystal ML ライブラリ: Autograd、Tensors、ニューラル ネットワーク、オプティマイザー - skuznetsov/cogni-ml

記事本文:
GitHub - skuznetsov/cogni-ml: Crystal ML ライブラリ: Autograd、Tensors、ニューラル ネットワーク、オプティマイザー · GitHub
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
スクズネツォフ
/
コグニミリリットル
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
にサインインする必要があります

通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,649 コミット 1,649 コミット .github .github bin bin docs docs 例 例 スクリプト スクリプト spec spec src src .gitignore .gitignore LANDMARKS.md LANDMARKS.md Makefile Makefile README.md README.md TODO.md TODO.md ml ml run_safe.sh run_safe.sh shard.lock shard.lock shard.yml shard.yml test_graph.cr test_graph.cr test_graph_debug.cr test_graph_debug.cr test_metal.cr test_metal.cr test_nr0.cr test_nr0.cr すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ネイティブ Apple Silicon GPU アクセラレーションを備えた Crystal 機械学習ライブラリ。
Cogni-ML には現在、次の 2 つの要素があります。
一般的な Crystal ML ツールキット: テンソル、autograd、NN レイヤー、オプティマイザー、GGUF リーダー、および llama.cpp バインディング。
GGUF モデルのネイティブ Metal 推論ラボ。nomic-embed-text-v2-moe 埋め込みと Qwen 3.5 テキスト生成に関する実稼働指向の作業を行っています。
nomic-embed-text-v2-moe のネイティブ メタル埋め込みパイプライン。
Apple Silicon Metal のネイティブ Qwen 3.5 9B GGUF 推論パス。
Q4_K/Q5_K/Q6_K/Q8_0 量子化された matmul カーネル。
チャンク化された Qwen 3.5 プレフィル、デコード ウェーブ スケジューリング、プロンプト状態キャッシュの復元、正確な投機的デコード ハーネス。
オフセットを意識したバリア最適化による ComputeGraph ウェーブ スケジューリング。
Crystal autograd エンジン、NN レイヤー、および Adam/AdamW オプティマイザー。
llama.cpp 一般的な GGUF モデル アクセス用の FFI バインディング。
ソース/ml/
コア/テンソル、シェイプ、メタルバッファ
autograd/ 変数、GradFn バックワードパス
nn/リニア、LayerNorm、MultiHeadAttendant、ViT
オプティム/アダム/アダムW
llm/llama.cpp FFI バインディング
gguf/GGUF リーダー、トークナイザー、逆量子化、Qwen35、NomicBertMoE
metal/デバイス、ComputeEncoder、ComputeGraph、GraphEncoder
クウェン 3.5 ネイティブ メタル
ネイティブ Qwen パスは Qwen3.5-9B-Q4_K_ をターゲットとしています。

Apple Silicon 上の M.gguf。このコードは以下をサポートしています。
Qwen 3.5 GGUF メタデータとトークナイザーの読み込み。
Q4_K、Q5_K、Q6_K、および Q8_0 の量子化投影。
GQA、部分 RoPE、KV キャッシュ書き込み、および融合出力プロジェクションを備えたフルアテンション レイヤー。
GPU 常駐のリカレント ステートとチャンク プレフィル スキャンを備えた DeltaNet/リカレント レイヤー。
最終トークンのトップ 1 ショートカットを使用したチャンク プレフィル。
ウェーブ スケジューリングをデコードして、コマンド バッファーの境界を削減します。
A/B 用の外部ラマトークン化フォールバックを備えたネイティブ Qwen BPE トークナイザー。
正確なプロンプト状態の保存/復元、トークン化されたプロンプトの再利用、および最長プレフィックスのプロンプト キャッシュ。
正確な投機的デコード ハーネス:
Qwen 3.5 0.8B Q8_0 を使用したニューラル ドラフト、
繰り返し/生成されたテンプレートテキストのn-gram/キャッシュドラフト、
より大きな受け入れられたチャンクの行バッチ処理された top1 を含むターゲット検証チャンク。
9B Q4_K_M パスは、主に検証されたターゲットです。 Qwen 3.6 27B はスケールアップ ターゲットですが、ローカルの正確性とパフォーマンスの実行が完了するまでは実験版として扱う必要があります。
開発者 CLI はデフォルトでローカルの LM Studio / llama.cpp 形式のパスになります。
~/.cache/lm-studio/models/lmstudio-community/Qwen3.5-9B-GGUF/Qwen3.5-9B-Q4_K_M.gguf
~/.cache/lm-studio/models/lmstudio-community/Qwen3.5-0.8B-GGUF/Qwen3.5-0.8B-Q8_0.gguf
~/SrcArchives/AI/llama.cpp/build/bin/llama-tokenize
~/SrcArchives/AI/llama.cpp/build/bin/llama-bench
ほとんどのベンチマーク/プローブ CLI は、 --model 、 --target 、 --draft 、 --tokenizer-bin 、または環境オーバーライドも受け入れます。 bin/qwen35_generate.cr は意図的に小さなデモであり、現在はファイルの先頭でその定数を使用しています。
Linux、CUDA ホスト、または Metal が利用できない環境で CPU のみの GGUF/Qwen メタデータ スモークを構築します。
Crystal build -Dcpu_only bin/qwen35_gguf_info.cr -o build/qwen35_gguf_info
./build/qwen35_gguf_info --model /path/to/Qwen3.5-9B-Q4_K_M.gguf

./build/qwen35_gguf_info --model /path/to/Qwen3.5-0.8B-Q8_0.gguf --load-weights
このエントリポイントは意図的に推論を実行しません。 Linux ビルドに Metal ブリッジを取り込むことなく、GGUF 解析、Qwen 3.5/3.6 hparams、テンソル インベントリ、構造化された Qwen35Weights ローダーを検証します。
NVIDIA/Linux ホスト上に最小限の Crystal CUDA Driver API スモークを構築します。
Crystal ビルド bin/cuda_driver_smoke.cr -o build/cuda_driver_smoke
./build/cuda_driver_smoke 4096
CUDA スモークはバックエンド境界プローブのみです。 libcuda をリンクし、埋め込まれた PTX をロードし、vector-add カーネルを起動して、結果をチェックします。
CUDA プローブ コードは、再利用可能な CUDA コンテキスト、モジュール/関数、起動、コピー、同期、およびデバイス バッファーの所有権に src/ml/cuda/driver.cr を使用します。これは意図的に小さくなっています。生の CUDA ドライバー API のライフサイクルと呼び出しを所有しますが、CUDA バックエンドの分割が促進されるまで、上位レベルのレイヤーの実行は依然としてプローブローカルです。
また、明示的な Upload_weights 、 replace_sequence 、 run_sequence 、および read_outputs フェーズを備えた常駐シーケンス プローブ用のシン ライフサイクル ファサードである ML::CUDA::ResidentSequenceRunner も提供します。
src/ml/cuda/qwen_recurrent_layer_runner.cr は、Qwen 固有の最初のランナー抽出です。これは、1 つの再帰層の CUDA モジュール、デバイス バッファー、カーネル パラメーター、重みアップロード、シーケンス リセット、トークン起動グラフ、および出力リードバックを所有します。 QwenRecurrentLayerRunner::Weights.load は、Q4_K または Q6_K として保存されたリカレント層 ffn_down テンソルを含む、GGUF テンソル ルックアップ、テンソル形状/型検証、およびランナーの生の重み読み取りを所有します。 CPU とリファレンスの比較は意図的にプローブ内に残ります。
NVIDIA/Linux ホスト上で最初の量子化された CUDA 正確性プローブを構築します。
Crystal build -Dcpu_only bin/cuda_q8_gemv_probe.cr -o build/cuda_q8_gemv_probe
./build/cuda_q8_gemv_pro

\になる
--model /path/to/Qwen3.5-0.8B-Q8_0.gguf \
--tensor blk.0.ffn_up.weight \
--カーネルワープ4 \
--担当者 100 \
--ウォームアップ 10
cuda_q8_gemv_probe は実際の GGUF Q8_0 テンソルをロードし、生の GGUF ブロック レイアウト上で Crystal 駆動の CUDA ドライバー API GEMV カーネルを起動し、既存の CPU QuantMatmul リファレンスと比較します。 --kernel スカラーは、最初の出力行あたり 1 スレッドの正確性カーネルを維持します。デフォルトの --kernel warp4 は、スレッド ブロックごとに 4 つの出力行を 4 つのワープにマップし、現在より高速なプローブ形状です。これはまだスタンドアロンのバックエンド境界プローブであり、最適化された Qwen CUDA 推論パスではありません。現在の完全な qwen35_generate CLI はメタルファーストのままです。
Qwen 9B/27B スタイルのターゲット テンソル用の最初の Q4_K CUDA 正確性プローブを構築します。
Crystal build -Dcpu_only bin/cuda_q4k_gemv_probe.cr -o build/cuda_q4k_gemv_probe
./build/cuda_q4k_gemv_probe \
--model /path/to/Qwen3.5-9B-Q4_K_M.gguf \
--tensor blk.0.attn_gate.weight \
--カーネルワープ4 \
--担当者 20 \
--ウォームアップ 3
cuda_q4k_gemv_probe は、生の GGUF Q4_K ブロック レイアウト ( d 、 dmin 、12 バイトのパックされたスケール/分、128 バイトのパックされたニブル) を使用し、CPU QuantMatmul Q4_K リファレンスに対して CUDA 出力をチェックします。 --kernel スカラーは最初の正確性カーネルを保持します。デフォルトの --kernel warp4 は、ブロックごとに 4 つの出力行を 4 つのワープにマップし、現在より高速なプローブ形状です。
Q6_K に残る Q4_K_M テンソルの Q6_K CUDA 正確性/速度プローブを構築します。
Crystal build -Dcpu_only bin/cuda_q6k_gemv_probe.cr -o build/cuda_q6k_gemv_probe
./build/cuda_q6k_gemv_probe \
--model /path/to/Qwen3.5-9B-Q4_K_M.gguf \
--tensor blk.0.ffn_down.weight \
--カーネルワープ4 \
--担当者 10 \
--ウォームアップ 2
cuda_q6k_gemv_probe は、混合量子ターゲット モデルの出力/値/ダウン投影で使用される GGUF Q6_K ブロック レイアウト ( ql 、 qh 、符号付きスケール、 d ) をカバーします。 L

Q4_K/Q8_0 プローブと同様、スタンドアロンのバックエンド プリミティブ チェックです。完全な CUDA Qwen の実行は依然として別のバックエンド分割です。
最初の GPU 常駐 FFN シーケンス プローブを構築します。
Crystal build -Dcpu_only bin/cuda_ffn_sequence_probe.cr -o build/cuda_ffn_sequence_probe
./build/cuda_ffn_sequence_probe \
--model /path/to/Qwen3.5-9B-Q4_K_M.gguf \
-- レイヤー 0 \
--担当者 10 \
--ウォームアップ 2
cuda_ffn_sequence_probe は、入力、中間アクティベーション、および出力投影入力を GPU に常駐させたまま、チェックされた CUDA プリミティブを Q4_K ffn_gate + Q4_K ffn_up -> SwiGLU -> Q6_K ffn_down として構成します。 CPU QuantMatmul FFN 参照と比較するために、最後の隠しベクトルのみがコピーされて戻されます。
全注意入力投影バンドル プローブを構築します。
Crystal build -Dcpu_only bin/cuda_attn_projection_probe.cr -o build/cuda_attn_projection_probe
./build/cuda_attn_projection_probe \
--model /path/to/Qwen3.5-9B-Q4_K_M.gguf \
-- レイヤー 3 \
--担当者 10 \
--ウォームアップ 2
cuda_attn_projection_probe は、GPU に常駐する 1 つの隠れベクトルから Q4_K attn_q + Q4_K attn_k + Q6_K attn_v を実行し、すべての投影が完了した後でのみ Q/K/V をコピーして戻します。 Qwen3.5 9Bのblk.3などのフルアテンション層を対象としています。
プローブは ML::CUDA::QwenFullAttnProjectionRunner を介してルーティングされ、 --tokens N をサポートし、最終的な正確性リードバックまで Q/K/V 出力を GPU に常駐させます。これは将来の全注意/KV CUDA 作業のための再利用可能な入力投影境界であり、まだ完全な全注意レイヤー ランナーではありません。
フルアテンション Q/K 正規化 + RoPE + KV キャッシュ境界プローブを構築します。
Crystal build -Dcpu_only bin/cuda_full_attn_kv_probe.cr -o build/cuda_full_attn_kv_probe
./build/cuda_full_attn_kv_probe \
--model /path/to/Qwen3.5-9B-Q4_K_M.gguf \
-- レイヤー 3 \
--トークン 4 \
--start-pos 2 \
--max-seq 12
cuda_full_attn_kv

_probe は、投影ランナーと ML::CUDA::QwenFullAttnKVRunner の周囲の残りの非表示から最終非表示へのラッパーである ML::CUDA::QwenFullAttnLayerRunner を介してルーティングされるようになりました。投影ランナーは、Q/K/V 投影の前に、残留隠れ状態から CUDA に初期 attn_norm を適用できます。 Q は正規化/RoPE'd Q とゲートに分割され、K は RMSNormed/RoPE'd で、K/V 行は start_pos で CUDA 常駐キャッシュに追加され、正確性優先のシリアル CUDA カーネルが GQA スコア、ソフトマックス、値リダクション、および Q ゲート乗算を計算し、常駐ゲート付きアテンション出力が attn_output.weight を通じて投影され、層末尾で残差加算が実行されます。ポストアテンション RMSNorm、FFN ゲート/アップ/SwiGLU/ダウン、および最終残差。 Q、ゲート、K、ゲート アテンション出力、投影されたアテンション出力、最終非表示、K キャッシュ、および V キャッシュを CPU Qwen リファレンスと照合してチェックします。これは、混合スタック構成用のデバイス入出力フックを備えたクリーンな 1 層セマンティクス プローブになりましたが、まだエンドツーエンドの Linux デコード パスではありません。フル/リカレント スタック スケジューリング、ロジット/top1、トークナイザー/サンプリング、復元された非ゼロ プレフィックス KV、および高速アテンション カーネルは別個のゲートのままです。
Q5_K CUDA 再帰 QKV プローブを構築します。
クリスタルビルド -

[切り捨てられた]

## Original Extract

Crystal ML library: Autograd, Tensors, Neural Networks, Optimizers - skuznetsov/cogni-ml

GitHub - skuznetsov/cogni-ml: Crystal ML library: Autograd, Tensors, Neural Networks, Optimizers · GitHub
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
skuznetsov
/
cogni-ml
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,649 Commits 1,649 Commits .github .github bin bin docs docs examples examples scripts scripts spec spec src src .gitignore .gitignore LANDMARKS.md LANDMARKS.md Makefile Makefile README.md README.md TODO.md TODO.md ml ml run_safe.sh run_safe.sh shard.lock shard.lock shard.yml shard.yml test_graph.cr test_graph.cr test_graph_debug.cr test_graph_debug.cr test_metal.cr test_metal.cr test_nr0.cr test_nr0.cr View all files Repository files navigation
Crystal machine learning library with native Apple Silicon GPU acceleration.
Cogni-ML is currently two things:
A general Crystal ML toolkit: tensors, autograd, NN layers, optimizers, GGUF readers, and llama.cpp bindings.
A native Metal inference lab for GGUF models, with production-oriented work on nomic-embed-text-v2-moe embeddings and Qwen 3.5 text generation.
Native Metal embedding pipeline for nomic-embed-text-v2-moe .
Native Qwen 3.5 9B GGUF inference path for Apple Silicon Metal.
Q4_K/Q5_K/Q6_K/Q8_0 quantized matmul kernels.
Chunked Qwen 3.5 prefill, decode wave scheduling, prompt-state cache restore, and exact speculative decode harnesses.
ComputeGraph wave scheduling with offset-aware barrier optimization.
Crystal autograd engine, NN layers, and Adam/AdamW optimizers.
llama.cpp FFI bindings for general GGUF model access.
src/ml/
core/ Tensor, Shape, MetalBuffer
autograd/ Variable, GradFn backward pass
nn/ Linear, LayerNorm, MultiHeadAttention, ViT
optim/ Adam/AdamW
llm/ llama.cpp FFI bindings
gguf/ GGUF reader, tokenizer, dequantization, Qwen35, NomicBertMoE
metal/ Device, ComputeEncoder, ComputeGraph, GraphEncoder
Qwen 3.5 Native Metal
The native Qwen path targets Qwen3.5-9B-Q4_K_M.gguf on Apple Silicon. The code supports:
Qwen 3.5 GGUF metadata and tokenizer loading.
Q4_K, Q5_K, Q6_K, and Q8_0 quantized projections.
Full-attention layers with GQA, partial RoPE, KV cache writes, and fused output projection.
DeltaNet/recurrent layers with GPU-resident recurrent state and chunked prefill scan.
Chunked prefill with final-token top1 shortcut.
Decode wave scheduling to reduce command-buffer boundaries.
Native Qwen BPE tokenizer with an external llama-tokenize fallback for A/B.
Exact prompt-state save/restore, tokenized-prompt reuse, and longest-prefix prompt cache.
Exact speculative decode harnesses:
neural draft with Qwen 3.5 0.8B Q8_0,
n-gram/cache draft for repeated/generated-template text,
target-verifier chunks with row-batched top1 for larger accepted chunks.
The 9B Q4_K_M path is the primary verified target. Qwen 3.6 27B is a scale-up target, but it should be treated as experimental until local correctness and performance runs are completed.
The developer CLIs default to local LM Studio / llama.cpp-style paths:
~/.cache/lm-studio/models/lmstudio-community/Qwen3.5-9B-GGUF/Qwen3.5-9B-Q4_K_M.gguf
~/.cache/lm-studio/models/lmstudio-community/Qwen3.5-0.8B-GGUF/Qwen3.5-0.8B-Q8_0.gguf
~/SrcArchives/AI/llama.cpp/build/bin/llama-tokenize
~/SrcArchives/AI/llama.cpp/build/bin/llama-bench
Most benchmark/probe CLIs also accept --model , --target , --draft , --tokenizer-bin , or environment overrides. bin/qwen35_generate.cr is intentionally a small demo and currently uses its constants at the top of the file.
Build the CPU-only GGUF/Qwen metadata smoke on Linux, CUDA hosts, or any environment where Metal is unavailable:
crystal build -Dcpu_only bin/qwen35_gguf_info.cr -o build/qwen35_gguf_info
./build/qwen35_gguf_info --model /path/to/Qwen3.5-9B-Q4_K_M.gguf
./build/qwen35_gguf_info --model /path/to/Qwen3.5-0.8B-Q8_0.gguf --load-weights
This entrypoint intentionally does not run inference. It verifies GGUF parsing, Qwen 3.5/3.6 hparams, tensor inventory, and the structured Qwen35Weights loader without pulling the Metal bridge into a Linux build.
Build the minimal Crystal CUDA Driver API smoke on NVIDIA/Linux hosts:
crystal build bin/cuda_driver_smoke.cr -o build/cuda_driver_smoke
./build/cuda_driver_smoke 4096
The CUDA smoke is a backend boundary probe only: it links libcuda , loads embedded PTX, launches a vector-add kernel, and checks the result.
CUDA probe code uses src/ml/cuda/driver.cr for reusable CUDA context, module/function, launch, copy, synchronize, and device-buffer ownership. This is intentionally small: it owns the raw CUDA Driver API lifecycle and calls, while higher-level layer execution is still probe-local until the CUDA backend split is promoted.
It also provides ML::CUDA::ResidentSequenceRunner , a thin lifecycle facade for resident sequence probes with explicit upload_weights , reset_sequence , run_sequence , and read_outputs phases.
src/ml/cuda/qwen_recurrent_layer_runner.cr is the first Qwen-specific runner extraction: it owns one recurrent layer's CUDA modules, device buffers, kernel parameters, weight upload, sequence reset, token launch graph, and output readback. QwenRecurrentLayerRunner::Weights.load owns GGUF tensor lookup, tensor-shape/type validation, and raw weight reads for the runner, including recurrent-layer ffn_down tensors stored as either Q4_K or Q6_K. CPU-reference comparison intentionally remains in the probe.
Build the first quantized CUDA correctness probe on NVIDIA/Linux hosts:
crystal build -Dcpu_only bin/cuda_q8_gemv_probe.cr -o build/cuda_q8_gemv_probe
./build/cuda_q8_gemv_probe \
--model /path/to/Qwen3.5-0.8B-Q8_0.gguf \
--tensor blk.0.ffn_up.weight \
--kernel warp4 \
--reps 100 \
--warmup 10
cuda_q8_gemv_probe loads a real GGUF Q8_0 tensor, launches a Crystal-driven CUDA Driver API GEMV kernel over the raw GGUF block layout, and compares against the existing CPU QuantMatmul reference. --kernel scalar keeps the first one-thread-per-output-row correctness kernel; the default --kernel warp4 maps four output rows to four warps per thread block and is the current faster probe shape. This is still a standalone backend-boundary probe, not an optimized Qwen CUDA inference path yet. The current full qwen35_generate CLI remains Metal-first.
Build the first Q4_K CUDA correctness probe for Qwen 9B/27B-style target tensors:
crystal build -Dcpu_only bin/cuda_q4k_gemv_probe.cr -o build/cuda_q4k_gemv_probe
./build/cuda_q4k_gemv_probe \
--model /path/to/Qwen3.5-9B-Q4_K_M.gguf \
--tensor blk.0.attn_gate.weight \
--kernel warp4 \
--reps 20 \
--warmup 3
cuda_q4k_gemv_probe uses the raw GGUF Q4_K block layout ( d , dmin , 12-byte packed scales/mins, 128-byte packed nibbles) and checks the CUDA output against the CPU QuantMatmul Q4_K reference. --kernel scalar keeps the first correctness kernel; the default --kernel warp4 maps four output rows to four warps per block and is the current faster probe shape.
Build the Q6_K CUDA correctness/speed probe for Q4_K_M tensors that remain in Q6_K:
crystal build -Dcpu_only bin/cuda_q6k_gemv_probe.cr -o build/cuda_q6k_gemv_probe
./build/cuda_q6k_gemv_probe \
--model /path/to/Qwen3.5-9B-Q4_K_M.gguf \
--tensor blk.0.ffn_down.weight \
--kernel warp4 \
--reps 10 \
--warmup 2
cuda_q6k_gemv_probe covers the GGUF Q6_K block layout ( ql , qh , signed scales, d ) used by output/value/down projections in mixed-quant target models. Like the Q4_K/Q8_0 probes, it is a standalone backend primitive check; full CUDA Qwen execution is still a separate backend split.
Build the first GPU-resident FFN sequence probe:
crystal build -Dcpu_only bin/cuda_ffn_sequence_probe.cr -o build/cuda_ffn_sequence_probe
./build/cuda_ffn_sequence_probe \
--model /path/to/Qwen3.5-9B-Q4_K_M.gguf \
--layer 0 \
--reps 10 \
--warmup 2
cuda_ffn_sequence_probe composes the checked CUDA primitives as Q4_K ffn_gate + Q4_K ffn_up -> SwiGLU -> Q6_K ffn_down while keeping the input, intermediate activations, and output projection input GPU-resident. Only the final hidden vector is copied back for comparison against the CPU QuantMatmul FFN reference.
Build the full-attention input projection bundle probe:
crystal build -Dcpu_only bin/cuda_attn_projection_probe.cr -o build/cuda_attn_projection_probe
./build/cuda_attn_projection_probe \
--model /path/to/Qwen3.5-9B-Q4_K_M.gguf \
--layer 3 \
--reps 10 \
--warmup 2
cuda_attn_projection_probe runs Q4_K attn_q + Q4_K attn_k + Q6_K attn_v from one GPU-resident hidden vector and copies Q/K/V back only after all projections complete. It targets full-attention layers such as blk.3 in Qwen3.5 9B.
The probe now routes through ML::CUDA::QwenFullAttnProjectionRunner , supports --tokens N , and keeps Q/K/V outputs GPU-resident until the final correctness readback. It is the reusable input-projection boundary for future full-attention/KV CUDA work, not a complete full-attention layer runner yet.
Build the full-attention Q/K normalization + RoPE + KV-cache boundary probe:
crystal build -Dcpu_only bin/cuda_full_attn_kv_probe.cr -o build/cuda_full_attn_kv_probe
./build/cuda_full_attn_kv_probe \
--model /path/to/Qwen3.5-9B-Q4_K_M.gguf \
--layer 3 \
--tokens 4 \
--start-pos 2 \
--max-seq 12
cuda_full_attn_kv_probe now routes through ML::CUDA::QwenFullAttnLayerRunner , a residual-hidden-to-final-hidden wrapper around the projection runner and ML::CUDA::QwenFullAttnKVRunner . The projection runner can apply the initial attn_norm on CUDA from residual hidden states before Q/K/V projection; Q is split into normalized/RoPE'd Q and gate, K is RMSNormed/RoPE'd, K/V rows are appended to a CUDA-resident cache at start_pos , a correctness-first serial CUDA kernel computes GQA scores, softmax, value reduction, and Q-gate multiplication, the resident gated attention output is projected through attn_output.weight , and the layer tail runs residual add, post-attention RMSNorm, FFN gate/up/SwiGLU/down, and final residual. It checks Q, gate, K, gated attention output, projected attention output, final hidden, K-cache, and V-cache against the CPU Qwen reference. This is now a clean one-layer semantics probe with device input/output hooks for mixed-stack composition, but it is not yet an end-to-end Linux decode path: full/recurrent stack scheduling, logits/top1, tokenizer/sampling, restored nonzero prefix KV, and a faster attention kernel remain separate gates.
Build the Q5_K CUDA recurrent-QKV probe:
crystal build -

[truncated]
