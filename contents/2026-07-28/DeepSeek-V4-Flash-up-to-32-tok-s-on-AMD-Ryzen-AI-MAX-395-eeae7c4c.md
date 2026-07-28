---
source: "https://www.lucebox.com/blog/deepseek-v4-strix-halo"
hn_url: "https://news.ycombinator.com/item?id=49085014"
title: "DeepSeek V4 Flash, up to 32 tok/s on AMD Ryzen AI MAX+ 395"
article_title: "DeepSeek V4 Flash: 284B model, up to 32 tok/s on AMD Ryzen AI MAX+ 395 | lucebox"
author: "GreenGames"
captured_at: "2026-07-28T15:12:40Z"
capture_tool: "hn-digest"
hn_id: 49085014
score: 2
comments: 0
posted_at: "2026-07-28T15:03:04Z"
tags:
  - hacker-news
  - translated
---

# DeepSeek V4 Flash, up to 32 tok/s on AMD Ryzen AI MAX+ 395

- HN: [49085014](https://news.ycombinator.com/item?id=49085014)
- Source: [www.lucebox.com](https://www.lucebox.com/blog/deepseek-v4-strix-halo)
- Score: 2
- Comments: 0
- Posted: 2026-07-28T15:03:04Z

## Translation

タイトル: DeepSeek V4 フラッシュ、AMD Ryzen AI MAX+ 395 で最大 32 tok/s
記事のタイトル: DeepSeek V4 フラッシュ: 284B モデル、AMD Ryzen AI MAX+ 395 で最大 32 tok/s |ルースボックス
説明: AMD 搭載 Lucebox は、128 GB ユニファイド メモリを備えた AMD Ryzen AI MAX+ 395 上で 284B DeepSeek V4 フラッシュ モデルをローカルに実行し、最大 32 tok/s デコードと約 250 tok/s スパース プリフィルに達します。

記事本文:
DeepSeek V4 フラッシュ: 284B モデル、AMD Ryzen AI MAX+ 395 で最大 32 tok/s | lucebox すべての投稿 2026 年 7 月
DeepSeek V4 フラッシュ: 284B モデル、AMD Ryzen AI MAX+ 395 で最大 32 tok/s
AMD-Powered Lucebox は、128 GB ユニファイド メモリを搭載した AMD Ryzen AI MAX+ 395 上で完全な DeepSeek V4 フラッシュ ターゲットをローカルに実行します。デコードは最大 32.0 tok/s、インデックス付きスパース プリフィルでは約 250 tok/s です。両方のパスに使用されるコードは Lucebox main にあります。
DeepSeek V4 フラッシュは、284B の専門家混合モデルです。 CPU と Radeon 8060S が同じ 128 GB LPDDR5X プールを共有しているため、ターゲットとドラフトの両方が適合します。この実行には個別の GPU、2 台目のマシン、またはリモート推論サービスはありません。プロンプトとモデルの実行はローカル システムに残ります。
実行を LocalMaxxing に送信しました。 7 月 18 日、Radeon 8060S 向けの次に速い DeepSeek V4 Flash エントリは、18.99 tok/s の HipFire でした。このサイトの Ryzen AI Max 395 ユニファイド メモリ グループのこれまでの最高速度は、DwarfStar の 15.6 tok/s でした。
これにより、私たちのランは HipFire より 68.5% リードし、DwarfStar の結果の 2.05 倍となります。これらは、上記の公開 LocalMaxxing エントリとの比較であり、管理された A/B テストではありません。
ROCmFPX: 284B の重みを 128 GB に適合
ROCmFPX は 1 つの量子化フォーマットではありません。これは、AMD ROCm/HIP パスを中心に構築されたブロック形式のファミリーです。各ブロックは、パックされた低ビット コードとして 32 の重みと 1 つまたは 2 つの小さなスケールを保持します。 ROCmFP2 は、ブロックを 10 バイト、つまり重みあたり 2.50 ビットで保存します。 ROCmFP3 は重みあたり 3.50 ビットを使用します。高速 ROCmFP4 レイアウトは 4.25 を使用します。
DeepSeek V4 フラッシュの場合、不足している 2 ビット フォーマットとその HIP カーネルを追加し、Strix 固有の混合精度レシピを構築しました。巨大な配線済みエキスパート ゲートおよびアップ マトリックスは ROCmFP2 を使用し、エキスパート ダウン プロジェクションは ROCmFP3 を使用し、高密度またはより高感度のプロジェクションは ROCmFP4 以上の p を維持します。

精度。量子化中に重要度行列を使用し、モデルの MTP ヘッドを維持しました。最終的な 102.3 GB の目標は、パラメータあたり約 2.88 ビットになります。ファイル名に ROCmFP2 と記載されているのは、それが主流の形式であるためであり、すべてのテンソルが 2 ビットであるためではありません。
フォーマットも速度結果の一部です。バッチ 1 では、生成されたすべてのトークンが 43 レイヤーすべてにアクティブな重みをストリーミングするため、デコードはメモリ トラフィックによってほとんど制限されます。 ROCmFPX カーネルは、パックされたブロックを直接読み取り、AMD バイト置換命令を使用してレジスタ内の小さなコードブックを拡張し、別個のコードブックを収集することなく整数ドット積を供給します。実際には、ファイル レイアウトと GPU カーネルは 1 つのパスとして設計されています。
ROCmFPX は重み付けトラフィックを処理します。次に、モデルのハイパー接続、アテンション、ルーティング、専門家の作業のための DeepSeek 固有の HIP デコード パスを追加しました。投機的なドラフトがない場合、そのターゲットは自己回帰で 25.31 tok/s で実行されます。
DSpark は次の層です。 q=4 のバッチでは、その小さなドラフトは最大 3 つの新しいトークンを提案し、284B ターゲットは 1 つの融合パスで現在のシードを含む 4 つのポジションを検証します。
q=4 のキャップと適応幅を無効にすると、パブリック実行は 32.0 tok/s に達し、25.31 tok/s の自己回帰結果を 26.4% 上回りました。利益は、ターゲットが受け入れるドラフトトークンの数によって異なります。
また、ML チームがより高速な GPU カーネルを見つけて出荷するのに役立つ Geometric とも連携しました。 Lucebox の場合、DeanoC は専用の ROCm TOP_K および ARGSORT カーネルを使用して HIP ルーティング パスを最適化しました。 A/B テストでは、カーネル マイクロベンチでエキスパートの選択が 3.5 ～ 7 倍高速になり、エンドツーエンドのデコードが 0.44% 向上しました。この作業は個別に測定されたものであり、パブリックな 32 トーク/秒のメインブランチ実行の一部ではありません。
main では、 q=4 ROCmFP4 パスが各パックされた密な重みを 1 回ずつデコードするようになりました。

nd は、それを 4 回解凍するのではなく、4 つの検証列すべてに適用します。この変化により 2.1 ～ 2.3% 増加しました。 ROCmFPX フォーマット、融合デコード パス、DSpark ベリファイア、スパース プレフィル、および q=4 重みの再利用はすべて、以下の複製に含まれています。統合ブランチやプライベート パッチは必要ありません。
LocalMaxxing リクエストでは、2,048 個のプロンプト トークン、510 個の出力トークン、8,192 個のトークン コンテキスト、温度 0、およびバッチ 1 が使用されました。また、カーネルの変更後に 5 つの GSM プロンプトと 5 つの MATH プロンプトを再実行しました。 10 件すべてが期待どおりの回答を返しました。これは回帰チェックであり、精度のベンチマークではありません。
32 tok/s プロファイルは、モデルのデフォルトである 6 人のエキスパートの代わりに --ds4-expert-top-k 4 を使用します。これによりモデルの実行が変更され、速度と引き換えにある程度の品質マージンが犠牲になります。特定のワークロードに使用する前に、そのワークロードに対する 6 人のエキスパートによる自己回帰出力と比較してください。
スパースプレフィル: 約 250 トーク/秒
パブリック LocalMaxxing リクエストは、 --ds4-prefill sparse を使用した 245 tok/s プレフィルを報告します。別の 7,960 トークンの検証では、インデックス付きスパース プリフィルが 251.79 トークン/秒に達しました。 8K の場合は 246.8 ～ 255.9 tok/s の範囲でした。約 24,000 トークンの場合、スループットは 221.9 トークン/秒でした。
スパース プレフィルは、DeepSeek V4 の学習済みインデクサーを使用して、圧縮履歴への注目を制限します。また、作業をレイヤーごとにバッチ処理し、浮動小数点リダクションの順序を変更します。出力はトークンごとに正確なプリフィルとバイト同一ではないため、スパース モードはオプトインのままです。小型の GSM8K セットでは 10/10 のスコア、HumanEval スモーク セットでは 3/3 のスコアでした。私たちはまだ広範な品質評価を行っていません。
参考までに、トークンごとの正確なプレフィルでは、短い GSM/MATH プロンプトで 22.5 ～ 23 tok/s、ペアの 7,960 トークン テストで 16.46 tok/s が測定されました。これらの図では別の実行パスが使用されています。約 250 トーク/秒の結果をインデックス付きスパにのみ使用します

rse プレフィル。
ROCm 7.2.4 がすでにインストールされている 128 GB Strix Halo マシンから開始します。
sudo apt-get アップデート
sudo apt-get install -y build-essential cmake git ninja-buildcurl \
hipblas-dev hipcub-dev rocblas-dev rocprim-dev rocwmma-dev
git clone --branch main --recurse-submodules \
https://github.com/Luce-Org/lucebox.git
CD ルースボックス
cmake -S サーバー -B サーバー/ビルドヒップ -G Ninja \
-DCMAKE_BUILD_TYPE=リリース \
-DCMAKE_HIP_COMPILER=/opt/rocm/lib/llvm/bin/clang++ \
-DDFLASH27B_GPU_BACKEND=ヒップ \
-DDFLASH27B_HIP_ARCHITECTURES=gfx1151 \
-DDFLASH27B_HIP_SM80_EQUIV=ON \
-DCMAKE_HIP_FLAGS=-DDFLASH_WAVE_SIZE=32 \
-DGGML_HIP_MMQ_MFMA=ON \
-DGGML_HIP_NO_VMM=ON \
-DGGML_HIP_GRAPHS=オフ
cmake --build server/build-hip --target dflash_server -j"$(nproc)" ROCmFPX ターゲットと DSpark ドラフトをダウンロードし、測定されたプロファイルを開始します。
mkdir -p モデル
カール -L -C - --再試行 5 \
-o モデル/DeepSeek-V4-Flash-ROCMFP2-STRIX.gguf \
「https://huggingface.co/Lucebox/DeepSeek-V4-Flash-ROCMFPX/resolve/main/DeepSeek-V4-Flash-ROCMFP2-STRIX.gguf」
カール -L -C - --再試行 5 \
-o モデル/DeepSeek-V4-Flash-DSpark-draft-Q4RMFP4-denseF16.gguf \
「https://huggingface.co/Lucebox/DeepSeek-V4-Flash-DSpark-Drafter-GGUF/resolve/main/DeepSeek-V4-Flash-DSpark-draft-Q4RMFP4-denseF16.gguf」
MODEL="$PWD/モデル/DeepSeek-V4-Flash-ROCMFP2-STRIX.gguf"
DRAFT="$PWD/モデル/DeepSeek-V4-Flash-DSpark-draft-Q4RMFP4-denseF16.gguf"
エコーパフォーマンス | sudo tee /sys/firmware/acpi/platform_profile
sudo /opt/rocm/bin/rocm-smi -d 0 --setperflevel high
printf '0\n' > /tmp/ds4_awidth
printf '4\n' > /tmp/ds4_spec_q
DFLASH_DS4_SPEC=1 \
DFLASH_DS4_FUSED_VERIFY=1 \
DFLASH_DS4_SPEC_Q=4 \
DFLASH_DS4_TIMING=1 \
DFLASH_DS4_DRAFT="$ドラフト" \
LUCE_MMVQ_MAX_NCOLS=4 \
./server/build-hip/dflash_server "$MODEL" \
--ターゲットデバイスのヒップ:0 \
--ホスト 127.0.0.1 --ポート 80

00\
--max-ctx 8192 --default-max-tokens 2048 \
--chunk 2048 --ds4-prefill sparse \
--ds4-fused-decode \
--ds4-expert-top-k 4 \
--prefix-cache-slots 0 --prefill-cache-slots 0 \
--disk-prefix-cache off モデルを一度ウォームアップし、温度: 0 を使用します。サーバーは、[deepseek4] DSpark デコード ラインにデコード速度を出力します。 DFLASH_DS4_SPEC_Q=4 は DS4 検証の上限を設定します。 --verify-width は Laguna オプションであり、ここでは使用されません。この実装により、コンプレッサー境界でバッチが短縮される可能性があります。これは、正しい状態処理に必要です。
2 つのクロック コマンドも測定セットアップの一部です。私たちの検証マシンでは、Radeon High はグラフィックス クロックを 2.9 GHz に維持しました。自動時計には 1 秒あたり数トークンのコストがかかる場合があります。公開されたリクエストでは、245 tok/s のスパース プリフィルと 32.0 tok/s のデコードが報告されました。スループットは、プロンプトの形状と、デコードの場合、ターゲットが受け入れる DSpark プロポーザルの数によって異なります。正確なプレフィルに切り替えるか、モデルの 6 つのエキスパートを復元すると、これらの数値は適用されなくなります。統合ブランチやプライベート パッチは必要ありません。
2026 年 7 月に、Ryzen AI MAX+ 395 / Radeon 8060S ( gfx1151 )、128 GB LPDDR5X、ROCm 7.2.4、プラットフォーム プロファイル パフォーマンス、Radeon パフォーマンス レベル高 (2.9 GHz を測定) で測定。ターゲット: DeepSeek-V4-Flash-ROCMFP2-STRIX.gguf 。ドラフト: DeepSeek-V4-Flash-DSpark-draft-Q4RMFP4-denseF16.gguf 。パブリック LocalMaxxing エントリ: 32.0 tok/s 出力および 245 tok/s スパース プリフィル、2,048 プロンプト トークン、510 出力トークン、8,192 コンテキスト、バッチ 1、温度 0。コミュニティの投稿ではさまざまなプロトコルが使用されます。この記事での比較は、公開されているデコード スループットに明示的に限定されています。
AMD Strix Halo の DFlash + PFlash
Gemma 4 26B と ds4-eval-92 での DeepSeek V4 フラッシュ
antirez/ds4: DeepSeek V4 Flash 実装
AMD Ryzen AI で DeepSeek V4 フラッシュを実行する

マックス+ 395
現在の Lucebox メイン: 128 GB ユニファイド メモリを搭載した AMD Ryzen AI MAX+ 395 上で、最大 32 tok/s のデコードと約 250 tok/s のインデックス付きスパース プリフィル。

## Original Extract

AMD-Powered Lucebox runs the 284B DeepSeek V4 Flash model locally on AMD Ryzen AI MAX+ 395 with 128 GB unified memory, reaching up to 32 tok/s decode and roughly 250 tok/s sparse prefill.

DeepSeek V4 Flash: 284B model, up to 32 tok/s on AMD Ryzen AI MAX+ 395 | lucebox All posts July 2026
DeepSeek V4 Flash: 284B model, up to 32 tok/s on AMD Ryzen AI MAX+ 395
AMD-Powered Lucebox runs the full DeepSeek V4 Flash target locally on AMD Ryzen AI MAX+ 395 with 128 GB unified memory: up to 32.0 tok/s decode and roughly 250 tok/s with indexed sparse prefill. The code used for both paths is on Lucebox main .
DeepSeek V4 Flash is a 284B mixture-of-experts model. Its target and draft both fit because the CPU and Radeon 8060S share the same 128 GB LPDDR5X pool. There is no discrete GPU, second machine, or remote inference service in this run: prompts and model execution stay on the local system.
We submitted the run to LocalMaxxing . On July 18, its next-fastest DeepSeek V4 Flash entry for the Radeon 8060S was HipFire at 18.99 tok/s. The previous best in the site’s Ryzen AI Max 395 unified-memory group was DwarfStar at 15.6 tok/s.
That puts our run 68.5% ahead of HipFire and at 2.05× the DwarfStar result. These are comparisons against the public LocalMaxxing entries shown above, not controlled A/B tests.
ROCmFPX: fitting 284B weights into 128 GB
ROCmFPX is not one quantization format. It is a family of block formats built around the AMD ROCm/HIP path. Each block holds 32 weights as packed low-bit codes plus one or two small scales. ROCmFP2 stores a block in 10 bytes, or 2.50 bits per weight; ROCmFP3 uses 3.50 bits per weight; and the fast ROCmFP4 layout uses 4.25.
For DeepSeek V4 Flash, we added the missing 2-bit format and its HIP kernels, then built a Strix-specific mixed-precision recipe. The enormous routed-expert gate and up matrices use ROCmFP2, expert down projections use ROCmFP3, and dense or more sensitive projections keep ROCmFP4 or higher precision. We used an importance matrix during quantization and kept the model’s MTP head. The final 102.3 GB target works out to roughly 2.88 bits per parameter ; the filename says ROCmFP2 because that is the dominant format, not because every tensor is 2-bit.
The format is also part of the speed result. At batch one, every generated token streams the active weights across all 43 layers, so decode is mostly limited by memory traffic. The ROCmFPX kernels read the packed blocks directly, expand their small codebooks in registers with AMD byte-permute instructions, and feed integer dot products without a separate codebook gather. In practice, the file layout and the GPU kernel are designed as one path.
ROCmFPX handles the weight traffic. We then added a DeepSeek-specific HIP decode path for the model’s hyper-connections, attention, routing, and expert work. With no speculative draft, that target runs at 25.31 tok/s autoregressive.
DSpark is the next layer. With a q=4 batch, its small draft proposes up to three new tokens and the 284B target verifies four positions, including the current seed, in one fused pass.
With a q=4 cap and adaptive width disabled, the public run reached 32.0 tok/s , 26.4% above the 25.31 tok/s autoregressive result. The gain varies with how many draft tokens the target accepts.
We also worked with Geometric , which helps ML teams find and ship faster GPU kernels. For Lucebox, DeanoC optimized the HIP routing path with dedicated ROCm TOP_K and ARGSORT kernels. In our A/B test, they made expert selection 3.5–7× faster in the kernel microbench and improved end-to-end decode by 0.44%. That work was measured separately and is not part of the public 32 tok/s main-branch run.
On main , the q=4 ROCmFP4 path now decodes each packed dense weight once and applies it to all four verification columns instead of unpacking it four times. That change added 2.1–2.3%. The ROCmFPX format, fused decode path, DSpark verifier, sparse prefill, and q=4 weight reuse are all included in the reproduction below; no integration branch or private patch is required.
The LocalMaxxing request used 2,048 prompt tokens, 510 output tokens, an 8,192-token context, temperature zero, and batch one. We also reran five GSM and five MATH prompts after the kernel changes; all ten completed with the expected answer. That is a regression check, not an accuracy benchmark.
The 32 tok/s profile uses --ds4-expert-top-k 4 instead of the model default of six experts. This changes model execution and trades some quality margin for speed. Before using it for a specific workload, compare it with six-expert autoregressive output on that workload.
Sparse prefill: roughly 250 tok/s
The public LocalMaxxing request reports 245 tok/s prefill with --ds4-prefill sparse . In a separate 7,960-token validation, indexed sparse prefill reached 251.79 tok/s; the 8K cases ranged from 246.8 to 255.9 tok/s . At roughly 24K tokens, throughput was 221.9 tok/s.
Sparse prefill uses DeepSeek V4’s learned indexer to limit compressed-history attention. It also batches work layer by layer, which changes floating-point reduction order. The output is not byte-identical to tokenwise exact prefill, so sparse mode remains opt-in. It scored 10/10 on our small GSM8K set and 3/3 on a HumanEval smoke set; we have not run a broad quality evaluation yet.
For reference, tokenwise exact prefill measured 22.5–23 tok/s on our short GSM/MATH prompts and 16.46 tok/s on the paired 7,960-token test. Those figures use a different execution path. Use the roughly 250 tok/s result only for indexed sparse prefill.
Starting from a 128 GB Strix Halo machine with ROCm 7.2.4 already installed:
sudo apt-get update
sudo apt-get install -y build-essential cmake git ninja-build curl \
hipblas-dev hipcub-dev rocblas-dev rocprim-dev rocwmma-dev
git clone --branch main --recurse-submodules \
https://github.com/Luce-Org/lucebox.git
cd lucebox
cmake -S server -B server/build-hip -G Ninja \
-DCMAKE_BUILD_TYPE=Release \
-DCMAKE_HIP_COMPILER=/opt/rocm/lib/llvm/bin/clang++ \
-DDFLASH27B_GPU_BACKEND=hip \
-DDFLASH27B_HIP_ARCHITECTURES=gfx1151 \
-DDFLASH27B_HIP_SM80_EQUIV=ON \
-DCMAKE_HIP_FLAGS=-DDFLASH_WAVE_SIZE=32 \
-DGGML_HIP_MMQ_MFMA=ON \
-DGGML_HIP_NO_VMM=ON \
-DGGML_HIP_GRAPHS=OFF
cmake --build server/build-hip --target dflash_server -j"$(nproc)" Download the ROCmFPX target and DSpark draft , then start the measured profile:
mkdir -p models
curl -L -C - --retry 5 \
-o models/DeepSeek-V4-Flash-ROCMFP2-STRIX.gguf \
"https://huggingface.co/Lucebox/DeepSeek-V4-Flash-ROCMFPX/resolve/main/DeepSeek-V4-Flash-ROCMFP2-STRIX.gguf"
curl -L -C - --retry 5 \
-o models/DeepSeek-V4-Flash-DSpark-draft-Q4RMFP4-denseF16.gguf \
"https://huggingface.co/Lucebox/DeepSeek-V4-Flash-DSpark-Drafter-GGUF/resolve/main/DeepSeek-V4-Flash-DSpark-draft-Q4RMFP4-denseF16.gguf"
MODEL="$PWD/models/DeepSeek-V4-Flash-ROCMFP2-STRIX.gguf"
DRAFT="$PWD/models/DeepSeek-V4-Flash-DSpark-draft-Q4RMFP4-denseF16.gguf"
echo performance | sudo tee /sys/firmware/acpi/platform_profile
sudo /opt/rocm/bin/rocm-smi -d 0 --setperflevel high
printf '0\n' > /tmp/ds4_awidth
printf '4\n' > /tmp/ds4_spec_q
DFLASH_DS4_SPEC=1 \
DFLASH_DS4_FUSED_VERIFY=1 \
DFLASH_DS4_SPEC_Q=4 \
DFLASH_DS4_TIMING=1 \
DFLASH_DS4_DRAFT="$DRAFT" \
LUCE_MMVQ_MAX_NCOLS=4 \
./server/build-hip/dflash_server "$MODEL" \
--target-device hip:0 \
--host 127.0.0.1 --port 8000 \
--max-ctx 8192 --default-max-tokens 2048 \
--chunk 2048 --ds4-prefill sparse \
--ds4-fused-decode \
--ds4-expert-top-k 4 \
--prefix-cache-slots 0 --prefill-cache-slots 0 \
--disk-prefix-cache off Warm the model once and use temperature: 0 . The server prints decode speed on its [deepseek4] DSpark decode line. DFLASH_DS4_SPEC_Q=4 sets the DS4 verification cap; --verify-width is a Laguna option and is not used here. The implementation may shorten a batch at a compressor boundary, which is required for correct state handling.
The two clock commands are also part of the measured setup. On our validation machine, Radeon high held the graphics clock at 2.9 GHz; automatic clocks can cost several tokens per second. The published request reported 245 tok/s sparse prefill and 32.0 tok/s decode. Throughput varies with prompt shape and, for decode, how many DSpark proposals the target accepts. If you switch to exact prefill or restore the model’s six experts, those numbers no longer apply. No integration branch or private patch is required.
Measured July 2026 on a Ryzen AI MAX+ 395 / Radeon 8060S ( gfx1151 ), 128 GB LPDDR5X, ROCm 7.2.4, platform profile performance , and Radeon performance level high (2.9 GHz observed). Target: DeepSeek-V4-Flash-ROCMFP2-STRIX.gguf . Draft: DeepSeek-V4-Flash-DSpark-draft-Q4RMFP4-denseF16.gguf . Public LocalMaxxing entry: 32.0 tok/s output and 245 tok/s sparse prefill, 2,048 prompt tokens, 510 output tokens, 8,192 context, batch one, temperature zero. Community submissions use different protocols; comparisons in this article are explicitly limited to published decode throughput.
DFlash + PFlash on AMD Strix Halo
Gemma 4 26B vs DeepSeek V4 Flash on ds4-eval-92
antirez/ds4: a DeepSeek V4 Flash implementation
Run DeepSeek V4 Flash on AMD Ryzen AI MAX+ 395
Current Lucebox main: up to 32 tok/s decode and roughly 250 tok/s indexed sparse prefill on AMD Ryzen AI MAX+ 395 with 128 GB unified memory.
