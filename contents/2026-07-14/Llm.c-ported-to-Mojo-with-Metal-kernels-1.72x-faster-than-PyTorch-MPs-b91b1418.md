---
source: "https://github.com/ulmentflam/llm.mojo"
hn_url: "https://news.ycombinator.com/item?id=48910825"
title: "Llm.c ported to Mojo with Metal kernels, 1.72x faster than PyTorch MPs"
article_title: "GitHub - ulmentflam/llm.mojo: GPT-2 training in pure Mojo with hand-written CUDA and Metal GPU kernels. llm.c parity in bf16 on NVIDIA, 1.72x faster than PyTorch MPS on Apple Silicon. · GitHub"
author: "ulmentflam"
captured_at: "2026-07-14T19:03:32Z"
capture_tool: "hn-digest"
hn_id: 48910825
score: 1
comments: 0
posted_at: "2026-07-14T18:08:09Z"
tags:
  - hacker-news
  - translated
---

# Llm.c ported to Mojo with Metal kernels, 1.72x faster than PyTorch MPs

- HN: [48910825](https://news.ycombinator.com/item?id=48910825)
- Source: [github.com](https://github.com/ulmentflam/llm.mojo)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T18:08:09Z

## Translation

タイトル: Llm.c を Metal カーネルで Mojo に移植、PyTorch MP より 1.72 倍高速
記事のタイトル: GitHub - ulmentflam/llm.mojo: 手書きの CUDA および Metal GPU カーネルを使用した純粋な Mojo での GPT-2 トレーニング。 NVIDIA 上の bf16 の llm.c パリティは、Apple Silicon 上の PyTorch MPS より 1.72 倍高速です。 · GitHub
説明: 手書きの CUDA および Metal GPU カーネルを使用した純粋な Mojo での GPT-2 トレーニング。 NVIDIA 上の bf16 の llm.c パリティは、Apple Silicon 上の PyTorch MPS より 1.72 倍高速です。 - ulmentflam/llm.mojo

記事本文:
GitHub - ulmentflam/llm.mojo: 手書きの CUDA および Metal GPU カーネルを使用した純粋な Mojo での GPT-2 トレーニング。 NVIDIA 上の bf16 の llm.c パリティは、Apple Silicon 上の PyTorch MPS より 1.72 倍高速です。 · GitHub
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
ウルメントフラム
/
llm.モジョ
公共
通知
あなた

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
289 コミット 289 コミット データ データ ドキュメント ドキュメント 数字 数字 llmm llmm スクリプト スクリプト スイープ スイープ テスト テスト third_party third_party .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE Makefile Makefile README.md README.md bench_gemm.mojo bench_gemm.mojo bench_zero_w4_after.json bench_zero_w4_after.json bench_zero_w4_before.json bench_zero_w4_before.json bench_zero_w4_big_after.json bench_zero_w4_big_after.json bench_zero_w4_big_before.json bench_zero_w4_big_before.json bench_zero_world2.json bench_zero_world2.json bench_zero_world3.json bench_zero_world3.json bench_zero_world4.json bench_zero_world4.json bench_zero_world8.json bench_zero_world8.json calibrate_fp8_scales.mojo calibrate_fp8_scales.mojo dump_grads_gpt2.mojo dump_grads_gpt2.mojo infer_gpt2.mojo infer_gpt2.mojo pixi.lock pixi.lock pixi.toml pixi.toml profile_gpt2.mojo profile_gpt2.mojo profile_gpt2.py profile_gpt2.py pyrefly.toml pyrefly.toml要件.txt要件.txt test_gpt2.mojo test_gpt2.mojo train_gpt2.mojo train_gpt2.mojo train_gpt2.py train_gpt2.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
これは Andrej Karpathy の llm.c を私が移植したもので、Mojo の v1.0.0 リリースを記念して @dorjeduck の llm.🔥 の GPU カーネルを拡張しています (このプロジェクトは毎晩 1.0.0b3 を追跡します)。見出しの結果は次のとおりです。
NVIDIA GB10 では、どちらのトレーニング精度でも llm.c の CUDA パスと同等かそれを上回っています (bf16 パリティ、fp32 は TF32 で約 7% 高速)。
Apple M4 Max では、PyTorch MPS bf16 よりも 1.71 倍高速に動作しますが、Apple 独自の MLX は高速です。

まだです (「ベンチマーク」を参照)。
bf16/fp32 とともに、実用的な FP8 (e4m3/e5m2) および NVFP4 (e2m1) の低精度トレーニングを追加します。
元のプロジェクトの詳細な説明については、llm.c を参照してください。
注: このプロジェクトは夜間リリースの Mojo 1.0.0b3 に基づいています。
このプロジェクトは、Karpathy の llm.c を git サブモジュール (CPU/GPU として使用) として提供します。
ベンチマークのリファレンス) なので、 --recurse-submodules を使用してクローンを作成します。
git clone --recurse-submodules https://github.com/ulmentflam/llm.mojo.git
cd llm.mojo
そのフラグを使用せずにすでにクローンを作成している場合は、 git submodule update --init from を実行します。
リポジトリのルート (Makefile は、これを初回実行時に自動的に実行します)
ベンチマーク/profile-llmc-* ターゲットが必要です)。
お持ちでない場合は、pixi をインストールします。
カール -fsSL https://pixi.sh/install.sh |しー
ステップ 3: 依存関係と Git フックをインストールする
クイックセットアップ: Pixi 環境 + git プリコミット / プリプッシュ フック (実行される)
それぞれ make lint と make check; make install-hooks を参照してください。が必要です
すでに PATH 上にあるように事前にコミットし、
見つからない場合はスキップされますが、致命的ではありません):
インストールする
CUDA が利用可能でインストールされている場合 (このファイルの範囲外)、次を使用します。
代わりに CUDA 対応の同等のものを使用します。
インストールを行う-cuda
make install / make install-cuda は Tiny Shakespeare をダウンロードしません
データセットまたは GPT-2 124M スターター ウェイト (~1.5 GB)。それは別のステップです
make train 、 make verify 、または make benchmark* が機能する前に必要です。
データを作る
または、make install-with-data (または
CUDA バリアントの場合は install-cuda-with-data を作成します)。
電車を作る
追加のヘルプについては、「 make help 」を参照してください。
マルチ GPU トレーニング: WORLD_SIZE=N でビルド (GPU ごとに 1 ランク、コンパイル時)
そして、実行時に ZeRO シャーディング ステージを選択します。
ビルドをWORLD_SIZE=8にする
scripts/run_train_gpt2.sh -z 2 # ZeRO: 0=DDP、1=+opt シャード、2=+grad シャード、3=+param シャード
プレトラへ

Tiny Shakespeare の代わりに FineWeb で、最初にトークン化します
( --streaming は、小型ディスク ホスト上で ~60 GB の Arrow の実体化を回避します):
Pixi run python data/fineweb.py -t classic -v 10B -m gpt-2 --streaming
長時間にわたるトレーニングの実行では、クラッシュ/再起動に耐えられるようにする必要があります。
無人、自動監視で監視
(自己修復プロセス スーパーバイザ: 再起動時のチェックポイント再開、OOM
バッチの半分化、認識できない障害に対するクロード エージェントのエスカレーション）を介して
リポジトリの .autosentry/autosentry.yaml 、 scripts/run_train_gpt2_bf16.sh 、および
scripts/ensure_supervisor.sh 。参照
docs/ai/gpt2_124m_fineweb_training_run.md
完全に機能する例については、
ベンチマーク結果: (NVIDIA DGX Spark)
llm.mojo、llm.c、および PyTorch の平均トレーニング ループ時間。すべて、同一の比較で一致するハイパーパラメーターを使用します。 llm.c は、20 スレッドで OpenMP 対応で実行されます。 CPU の比較は float32 で、GPU の比較は float32 と bfloat16 の両方で実行されます。 Apple Silicon では、make benchmark-metal は PyTorch MPS に対して llm.mojo (Metal GPU) を実行します (llm.c には Metal ポートがないため、PyTorch MPS がベースラインとして使用されます)。これらの結果については、「Apple Silicon (Metal GPU)」セクションを参照してください。
GB10 での公式実行 (B=4、T=1024、ウォームアップとして最初の 5 をドロップする 40 ステップ、1 セッションに 6 つのアームすべてをインターリーブ、2026 年 7 月 11 日 15:31、メタル テスト修復と MLX ベンチマークのマージ後、出荷されたツリー HEAD c1a48d5 で直接測定。2026 年 7 月 10 日の結果を確認)ノイズ内のキャンペーン後のテーブル、6 つのアームすべてがその 1.7% 以内):
TF32 注: llm.mojo の fp32 GPU GEMM は、llm.c の fp32 動作と一致して、デフォルトで TF32 tensor コア (CUBLAS_COMPUTE_32F_FAST_TF32) を使用するようになりました。その fp32 ビルドは、計算能力 8.0 以上の GPU で TF32 を自動的に有効にするため、上記の fp32 行は次のようになります。 TF32対TF32。厳密な IEEE fp32 m を復元するには、-D LLMM_NO_TF32=1 を指定してビルドします。

ath (これは verify-gpu ゲートをオンにするものでもあります。デフォルトの TF32 パスには独自のゲートがあり、 make verify-gpu-tf32 )。
バックワード カーネルのメモ: 07-10 の午後の番号では、TF32 に 2 つのバックワード パス最適化が追加されています。再設計された融合 LN バックワード (1 つのレジスタ蓄積カーネルとチャネルごとのブロック ファイナライズ、呼び出しごとに 4 回の起動を置き換えます; -6.9 ms fp32 / -3.1 ms bf16 カーネル ファミリ時間) と融合、 128 ビットベクトル化された matmul dbias 削減 (-1.5 ms fp32 / -1.0 ms bf16)。どちらも上記の完全な検証バッテリーによってゲートされます。
マルチ GPU データ並列トレーニングは、単一プロセス内で GPU ごとに 1 ランクを実行します。 make build WORLD_SIZE=N でビルドし (コンパイル時、 llm.c の -DMULTI_GPU と同様)、実行時に -z 0|1|2|3 で ZeRO ステージを選択します。ステージ 0 はプレーン DDP、ステージ 1 はシャード オプティマイザ状態、ステージ 2 は勾配もシャード、ステージ 3 もパラメータをシャードします。 4 つのステージはすべて、シングル GPU ベースラインに対して等価ゲート処理されています (ワールド サイズ 2 および 8 では testing/test_zero_equivalence.mojo、パラメーターごとに 1e-5 に一致します。ステップごとのトレーニング損失はステージ全体で一致し、ワールド サイズ 8 では ~1e-4 になります)。コレクティブは、ドライバーによってルーティングされたデバイス間のコピーを介してステージングされたリデュース スキャッター/オールギャザーであるため、CUDA P2P または NVLink を使用せずに PCIe ボックスで動作します。設計については docs/ai/zero_multigpu_rewrite_2026-07-14.md を、検証キャンペーンについては docs/ai/zero_world8_verification_2026-07-14.md を参照してください。
4× NVIDIA RTX PRO 6000 Blackwell Max-Q (96 GB、PCIe、P2P なし)、GPT-2 124M、ランクあたり B=4 T=64、12 ステップ (最初の 2 つはウォームアップとしてドロップ)、ステージ 2/3 メモリ パス (レイヤーごとの勾配バケット化およびステージ 3 パラメータ ストリーミング) 後に測定:
各ステージは GPU ごとに追加のメモリを購入するようになりました。ステージ 1 のオプティマイザー状態のシャーディングにより、fp32 で 768 MiB が節約されます (bf16 では 1 GiB、fp32 マスターがシャードの重み付けを行う場合)

o)、ステージ 2 のバケット化された逆方向削減によりさらに 256 MiB が節約され、ステージ 3 のパラメータ ストリーミングによりさらに 256 MiB が節約されます。このモデル サイズのフロアは、結合された WTE 埋め込み (~150 MiB) です。これは、バックワード パスの両端に書き込まれ、LM ヘッド用に全体が集められるため、常駐します。これらのカーネルの語彙チャンク化は文書化されたフォローアップです ( docs/ai/zero_grad_bucketing_design_2026-07-14.md および docs/ai/zero_stage3_param_streaming_2026-07-14.md を参照)。節約効果はステップ時間と引き換えに発生します。ステージ 2 と 3 ではレイヤーごとのコレクティブに料金がかかり、ステージ 3 のジャストインタイム パラメータ収集のコストが最も高くなります。どれくらいの費用がかかるかは形状によって異なります。この小さなベンチマーク形状では、全体的なオーバーヘッドは 50 ミリ秒ステップの大部分ですが、製品形状 (B≥32、T=1024、約 250 ～ 470 ミリ秒ステップ) では数パーセントですが、そこではアクティベーションがピーク メモリを支配しており、ZeRO のモデル状態の節約はごく一部になります。ステージは、状態が支配的なモデル規模で維持されます。すべてのステージは、シングル GPU ベースラインに対して数値的に等価ゲートされたままになります (ワールド サイズ 2 および 8 での 1e-5 へのパラメーターごとの一致)。 make benchmark-zero ( BENCH_ZERO_WORLD=N ; JSON を書き込み、このチャートを Figures/ にレンダリングします) で再現します。
低精度トレーニング(FP8/NVFP4)
FP8 と NVFP4 は、推論形式だけでなく、トレーニング精度も機能しています。 FP8 は、ブロックごとの線形 GEMM (QKV/attn-proj/MLP fc/fc-proj、前方および後方) を遅延スケーリングを使用して一時的な e4m3/e5m2 オペランドに量子化します。計算は FP8 tensor コアで実行されますが、マスターの重みとオプティマイザーの状態は fp32 に残ります。 NVFP4 ブロックは、中央のトランスフォーマー ブロックの MLP GEMM を cuBLASLt 上の e2m1 にスケールします。確率的丸めとランダム アダマール変換 (公開されている NVFP4 トレーニング レシピに従って) を追加して、余分な量を制御します。

化の差異。どちらも GPT-2 124M スケールで bf16 と並行して収束します。以下の損失エンベロープを参照し、 docs/ai/ai_assisted_optimizations_and_benchmarks.md で verify-fp8-grads / fp8/fp4 ゲートを作成します。
ステップ時間測定 (B=4、T=1024、チェックポイント初期化 tinyshakespeare、アーム順序がラウンドごとに交互になる 2 ラウンド、破棄されたフレッシュバイナリウォームアップ実行後のアームあたり 40 ステップの測定値、2026 年 7 月 11 日、最適化キャンペーン後のツリー):
2026-07-10/11 の最適化キャンペーン (合体/融合量子化転置カーネル、永続的な fp8 重み転置キャッシュ、デュアル出力量子化、NVFP4 用の融合タイル RHT-prep) により、FP8 が 150.5 ミリ秒から 146.6 ミリ秒に、NVFP4 が 184.2 ミリ秒から 154.3 ミリ秒に短縮されました。この規模では。幅ではその効果はさらに大きくなります。774M クラスの d36 構成では、FP8 は bf16 より約 12% (0.878 倍) 高速になり、NVFP4 はパリティ (1.002 倍) に達します。オプションの調整済み静的スケール モード ( -D LLMM_FP8_STATIC_SCALES=1 、デフォルトはオフ) は、ステップごとの amax/scale-update カーネルを完全に削除します。構成ごとのオフライン キャリブレーションを犠牲にして、124M でさらに約 1.5%、d36 (0.853x) で約 3% 削減されます。 docs/ai/ai_assisted_optimizations_and_benchmarks.md の A1 の記事と最終的なキャンペーン スコアボードを参照してください。
124M パラメータでは、これらは数値/研究設定です

[切り捨てられた]

## Original Extract

GPT-2 training in pure Mojo with hand-written CUDA and Metal GPU kernels. llm.c parity in bf16 on NVIDIA, 1.72x faster than PyTorch MPS on Apple Silicon. - ulmentflam/llm.mojo

GitHub - ulmentflam/llm.mojo: GPT-2 training in pure Mojo with hand-written CUDA and Metal GPU kernels. llm.c parity in bf16 on NVIDIA, 1.72x faster than PyTorch MPS on Apple Silicon. · GitHub
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
ulmentflam
/
llm.mojo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
289 Commits 289 Commits data data docs docs figures figures llmm llmm scripts scripts sweeps sweeps tests tests third_party third_party .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE Makefile Makefile README.md README.md bench_gemm.mojo bench_gemm.mojo bench_zero_w4_after.json bench_zero_w4_after.json bench_zero_w4_before.json bench_zero_w4_before.json bench_zero_w4_big_after.json bench_zero_w4_big_after.json bench_zero_w4_big_before.json bench_zero_w4_big_before.json bench_zero_world2.json bench_zero_world2.json bench_zero_world3.json bench_zero_world3.json bench_zero_world4.json bench_zero_world4.json bench_zero_world8.json bench_zero_world8.json calibrate_fp8_scales.mojo calibrate_fp8_scales.mojo dump_grads_gpt2.mojo dump_grads_gpt2.mojo infer_gpt2.mojo infer_gpt2.mojo pixi.lock pixi.lock pixi.toml pixi.toml profile_gpt2.mojo profile_gpt2.mojo profile_gpt2.py profile_gpt2.py pyrefly.toml pyrefly.toml requirements.txt requirements.txt test_gpt2.mojo test_gpt2.mojo train_gpt2.mojo train_gpt2.mojo train_gpt2.py train_gpt2.py View all files Repository files navigation
This is my port of Andrej Karpathy's llm.c , extending the GPU kernels of @dorjeduck's llm.🔥 in honor of Mojo's v1.0.0 release (this project tracks the 1.0.0b3 nightly). The headline results:
On an NVIDIA GB10, it matches or beats llm.c's CUDA path at both training precisions (bf16 parity, fp32 ~7% faster with TF32).
On an Apple M4 Max, it runs 1.71× faster than PyTorch MPS bf16, though Apple's own MLX is faster still (see Benchmarks ).
It adds working FP8 (e4m3/e5m2) and NVFP4 (e2m1) low-precision training alongside bf16/fp32.
See llm.c for a detailed explanation of the original project.
Note : This project is based on nightly Mojo 1.0.0b3 release.
This project vendors Karpathy's llm.c as a git submodule (used as the CPU/GPU
reference for benchmarking), so clone with --recurse-submodules :
git clone --recurse-submodules https://github.com/ulmentflam/llm.mojo.git
cd llm.mojo
If you already cloned without that flag, run git submodule update --init from
the repo root (the Makefile also does this automatically the first time a
benchmark / profile-llmc-* target needs it).
If you don't have it, install pixi :
curl -fsSL https://pixi.sh/install.sh | sh
Step 3: Install Dependencies and Git Hooks
Quick setup: pixi environment + git pre-commit / pre-push hooks (which run
make lint and make check respectively; see make install-hooks ; requires
pre-commit to already be on your PATH , and is
skipped, not fatal, if not found):
make install
If you have CUDA available and installed (beyond the scope of this file), use
the CUDA-enabled equivalent instead:
make install-cuda
make install / make install-cuda do not download the Tiny Shakespeare
dataset or GPT-2 124M starter weights (~1.5 GB). That's a separate step,
needed before make train , make verify , or make benchmark* will work:
make data
Or combine both in one shot with make install-with-data (or
make install-cuda-with-data for the CUDA variant).
make train
For additional help, see make help .
Multi-GPU training: build with WORLD_SIZE=N (one rank per GPU, compile-time)
and choose a ZeRO sharding stage at runtime:
make build WORLD_SIZE=8
scripts/run_train_gpt2.sh -z 2 # ZeRO: 0=DDP, 1=+opt shard, 2=+grad shard, 3=+param shard
To pretrain on FineWeb instead of Tiny Shakespeare, tokenize it first
( --streaming avoids the ~60 GB Arrow materialization on small-disk hosts):
pixi run python data/fineweb.py -t classic -v 10B -m gpt-2 --streaming
For a long-running training run you want to survive crashes/reboots
unattended, supervise it with autosentry
(a self-healing process supervisor: checkpoint-resume on restart, OOM
batch-halving, Claude-agent escalation on unrecognized failures) via this
repo's .autosentry/autosentry.yaml , scripts/run_train_gpt2_bf16.sh , and
scripts/ensure_supervisor.sh . See
docs/ai/gpt2_124m_fineweb_training_run.md
for a full worked example.
Benchmark Results: (NVIDIA DGX Spark)
Average training loop times across llm.mojo, llm.c, and PyTorch, all with matched hyperparameters in an apples-to-apples comparison. llm.c runs OpenMP-enabled with 20 threads. CPU comparisons are float32, and GPU comparisons run both float32 and bfloat16. On Apple Silicon, make benchmark-metal runs llm.mojo (Metal GPU) against PyTorch MPS (llm.c has no Metal port, so PyTorch MPS fills in as the baseline). See the Apple Silicon (Metal GPU) section for those results.
Official run on the GB10 (B=4, T=1024, 40 steps with the first 5 dropped as warmup, all six arms interleaved in one session, 2026-07-11 15:31, measured directly on the shipped tree, HEAD c1a48d5 , after the Metal test-restoration + MLX-benchmark merge; confirms the 2026-07-10 post-campaign table within noise, all six arms within 1.7% of it):
TF32 note : llm.mojo's fp32 GPU GEMMs now use TF32 tensor cores by default ( CUBLAS_COMPUTE_32F_FAST_TF32 ), matching llm.c's fp32 behavior: its fp32 build auto-enables TF32 on any compute-capability-8.0+ GPU, so the fp32 rows above are TF32-vs-TF32. Build with -D LLMM_NO_TF32=1 to restore strict IEEE fp32 math (that is also what make verify-gpu gates on; the default TF32 path has its own gate, make verify-gpu-tf32 ).
Backward-kernel note : the 07-10 afternoon numbers add two backward-pass optimizations on top of TF32: a redesigned fused LN-backward (one register-accumulating kernel plus a block-per-channel finalize, replacing 4 launches per invocation; −6.9 ms fp32 / −3.1 ms bf16 kernel-family time) and a fused, 128-bit-vectorized matmul dbias reduction (−1.5 ms fp32 / −1.0 ms bf16). Both are gated by the full verify battery above.
Multi-GPU data-parallel training runs one rank per GPU inside a single process: build with make build WORLD_SIZE=N (compile-time, like llm.c's -DMULTI_GPU ) and pick the ZeRO stage at runtime with -z 0|1|2|3 . Stage 0 is plain DDP, stage 1 shards optimizer state, stage 2 also shards gradients, and stage 3 also shards parameters. All four stages are equivalence-gated against the single-GPU baseline ( tests/test_zero_equivalence.mojo at world sizes 2 and 8, per-parameter match to 1e-5; per-step training losses agree across stages to ~1e-4 at world size 8). The collectives are staged reduce-scatter/all-gather over driver-routed device-to-device copies, so they work on PCIe boxes without CUDA P2P or NVLink. See docs/ai/zero_multigpu_rewrite_2026-07-14.md for the design and docs/ai/zero_world8_verification_2026-07-14.md for the verification campaign.
Measured on 4× NVIDIA RTX PRO 6000 Blackwell Max-Q (96 GB, PCIe, no P2P), GPT-2 124M, B=4 T=64 per rank, 12 steps (first 2 dropped as warmup), after the stage-2/3 memory pass (per-layer gradient bucketing and stage-3 parameter streaming):
Each stage now buys additional per-GPU memory. Stage 1's optimizer-state sharding saves 768 MiB in fp32 (1 GiB in bf16, where the fp32 master weights shard too), stage 2's bucketed backward reduction saves another 256 MiB, and stage 3's parameter streaming another 256 MiB. The floor at this model size is the tied wte embedding (~150 MiB): it is written at both ends of the backward pass and gathered whole for the LM head, so it stays resident. Vocab-chunking those kernels is the documented follow-up (see docs/ai/zero_grad_bucketing_design_2026-07-14.md and docs/ai/zero_stage3_param_streaming_2026-07-14.md ). The savings trade against step time: stages 2 and 3 pay for their per-layer collectives, with stage 3's just-in-time parameter gathers costing the most. How much that costs depends on the shape. At this tiny benchmark shape the collective overhead is a large fraction of the 50 ms step, but at production shapes (B≥32, T=1024, ~250-470 ms steps) it is a few percent, while activations dominate peak memory there and ZeRO's model-state savings become a small fraction. The stages earn their keep at model scales where state dominates. All stages stay numerically equivalence-gated against the single-GPU baseline (per-parameter match to 1e-5 at world sizes 2 and 8). Reproduce with make benchmark-zero ( BENCH_ZERO_WORLD=N ; writes the JSON and renders this chart into figures/ ).
Low-precision training (FP8 / NVFP4)
FP8 and NVFP4 are working training precisions, not just inference formats. FP8 quantizes the per-block linear GEMMs (QKV/attn-proj/MLP fc/fc-proj, forward and backward) to transient e4m3/e5m2 operands with delayed scaling. The math runs in FP8 tensor cores, but the master weights and optimizer state stay in fp32. NVFP4 block-scales the middle transformer blocks' MLP GEMMs to e2m1 on cuBLASLt. It adds stochastic rounding and a random Hadamard transform (per the published NVFP4 training recipe) to control the extra quantization variance. Both converge alongside bf16 at GPT-2 124M scale. See the loss envelopes below and make verify-fp8-grads / the fp8/fp4 gates in docs/ai/ai_assisted_optimizations_and_benchmarks.md .
Step-time measurement (B=4, T=1024, checkpoint-init tinyshakespeare, 2 rounds with arm order alternated per round, 40 measured steps/arm after a discarded fresh-binary warmup run, 2026-07-11, post-optimization-campaign tree):
The 2026-07-10/11 optimization campaign (coalesced/fused quantize-transpose kernels, persistent fp8 weight-transpose caching, dual-output quantize, fused tiled RHT-prep for NVFP4) cut FP8 from 150.5 to 146.6 ms and NVFP4 from 184.2 to 154.3 ms at this scale. It pays off harder at width: at the 774M-class d36 config FP8 is now ~12% faster than bf16 (0.878×), and NVFP4 reaches parity (1.002×). An optional calibrated static-scales mode ( -D LLMM_FP8_STATIC_SCALES=1 , default off) removes the per-step amax/scale-update kernels entirely. It shaves a further ~1.5% at 124M and ~3% at d36 (0.853×), at the cost of per-config offline calibration. See the A1 writeup and the final campaign scoreboard in docs/ai/ai_assisted_optimizations_and_benchmarks.md .
At 124M params, these are numerics/research confi

[truncated]
