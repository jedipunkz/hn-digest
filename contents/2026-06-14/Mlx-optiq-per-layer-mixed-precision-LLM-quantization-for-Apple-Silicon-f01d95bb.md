---
source: "https://mlx-optiq.com/"
hn_url: "https://news.ycombinator.com/item?id=48529008"
title: "Mlx-optiq: per-layer mixed-precision LLM quantization for Apple Silicon"
article_title: "Run LLMs locally on your Mac (Apple Silicon) · mlx-optiq"
author: "codelion"
captured_at: "2026-06-14T18:36:55Z"
capture_tool: "hn-digest"
hn_id: 48529008
score: 2
comments: 0
posted_at: "2026-06-14T16:20:43Z"
tags:
  - hacker-news
  - translated
---

# Mlx-optiq: per-layer mixed-precision LLM quantization for Apple Silicon

- HN: [48529008](https://news.ycombinator.com/item?id=48529008)
- Source: [mlx-optiq.com](https://mlx-optiq.com/)
- Score: 2
- Comments: 0
- Posted: 2026-06-14T16:20:43Z

## Translation

タイトル: Mlx-optiq: Apple Silicon 向けのレイヤーごとの混合精度 LLM 量子化
記事のタイトル: Mac (Apple Silicon) で LLM をローカルに実行する · mlx-optiq
説明: Mac 上でローカルに LLM を実行します (M1 ～ M5): 混合精度量子化、LoRA 微調整、ロングコンテキスト KV サービング、画像入力、すべて MLX ネイティブ。 PyPI では、PyTorch、GPU クラスター、クラウドはありません。

記事本文:
mlx-optiq
概要
モデル
ドキュメント
ブログ
ピピ↗
最適化コンパイラ・MLX
クオンタイズ、微調整
LLM にサービスを提供する
完全にAppleシリコン上にあります。
M1 から M5 までの大規模な言語モデルを Mac 上でローカルに実行します。
混合精度の重みに対するレイヤーごとの感度分析。
ビットバジェットを考慮した LoRA 微調整。
OpenAI と Anthropic API の両方を話すサーバー (クロード コードをポイントする)
ローカルクオンツ）。視覚機能を備えたモデルであれば、テキストだけでなく画像も送信できます。
GPU クラスターも API キーもありません。
8ビット・センシティブレイヤー
4ビット・堅牢なレイヤー
$ pip install mlx-optiq
コピー
BF16 と比較して 3.1 倍の平均圧縮率
MTP / ドラフター経由のデコード +1.4 倍
+13.6 最高の能力スコアの増加
140,000 以上の HF ダウンロード/月
OptIQ Lab — Mac、モデル、クラウドなし
OptIQ ラボ · ローカルホスト:8080
ローカル LLM ワークベンチ、1 ピップでインストール可能: サンドボックス ツールとチャットし、モデルを比較し、微調整し、提供します。ここでは、4 ビット OptIQ Qwen3.5 が 90 tok/s で応答し、完全にオフラインです。
01 構築済みモデル
ドロップイン 4 ビット クオント。同じ重さで、よりスマートなビット。
Hugging Face 上の 16 個のプロダクション mlx-optiq 量子化 LLM。 Nemotron 3、MiniCPM5、Qwen3.5、Qwen3.6、および Gemma-4 ファミリ、1 B 密度から 35 B-A3B の専門家の混合物まで。これらはストック mlx-lm に直接ロードされます。特別なランタイムはありません。
Google の統合テキスト + ビジョン Gemma-4 (8.3 GB、画像入力付き)。能力スコア 68.2 (均一 4 ビットに対して +6.4) は、混合精度の最大の向上の 1 つであり、ディスク上で 9 GB 未満で出荷される最も強力なモデルです。
当社が出荷する単一数量の中で最大のもの。 20.8 GB に 31 B パラメータ、能力スコア 79.7 (均一 4 ビットに対して +3.5)。投機的なデコードには、対応する -assistant-bf16 ドラフターと組み合わせます。
17.5 GB でフロンティア クラスの推論を実現し、最高の能力スコア (83.0) を実現します。バンドルされた MTP ヘッドは、 optiqserve --mtp 経由で最大 1.4 倍のデコードを実現します。
デフォルトの日次運転

バージョン6.6 GB に 9 B パラメータ。能力スコア 66.8 (均一 4 ビットに対して +0.2)。混合精度 KV を介して 64 k までのロングコンテキスト。投機的デコード用にバンドルされた MTP ヘッド。
3 つのコマンドでゼロからサービス提供 LLM まで。
各ステップは元に戻すことができ、標準の MLX ツールで動作します。 mlx-optiq は付加的です。これらのいずれかをスキップしても、パイプラインは引き続き動作します。
純粋なパイソン。 mlx 、 mlx-lm 、および hackingface-hub を取り込みます。 Apple Silicon 上の Python 3.11 以降。
$ pip install mlx-optiq
ii
事前に構築されたクオンツを使用する
事前に構築された mlx-optiq quants は、ストック mlx-lm でロードされます。レイヤーごとのビット割り当てはモデルのメタデータに記録されます。特別なローダーは必要ありません。
mlx_lm からインポート読み込み、生成
モデル、tok = ロード ( "mlx-community/Qwen3.5-9B-OptiQ-4bit" )
out =generate (model, tok, Prompt= "混合精度量子化について説明します。" , max_tokens= 200 )
印刷（アウト）
iii
混合精度の KV を使用してサービスを提供する
KV キャッシュにはそれ自体の感度の問題があります。 optiq kv-cache はモデルごとに 1 回測定します。 optiqserve は、OpenAI 互換 API の背後で、結果として得られるレイヤーごとの構成を提供します。
# 1 ～ 2 分、モデルごとに 1 回
$ optiq kv-cache mlx-community/Qwen3.5-9B-OptiQ-4bit \
--target-bits 5.0 -o ./kv
# OpenAI + Anthropic 互換サーバー:8080
# /v1/chat/completions (OpenAI)
# /v1/messages (Anthropic; Claude Code、Anthropic SDK などで動作します)
$ optiqserve --model mlx-community/Qwen3.5-9B-OptiQ-4bit \
--kv-config ./kv/kv_config.json \
--ポート 8080
次はどこへ
各モデル ファミリには、モデル固有のサンプリングのデフォルトと推奨される使用例を記載したスタート ガイドがあります。
エージェントを構築しますか? llms.txt を IDE にドロップします。これは、1 つの Markdown ファイル内のライブラリ参照全体です。
03 機能
1 つの感度信号。それを中心としたツールキット全体。
レイヤーごとの単一の KL ダイバージェンス パスにより、重み、KV キャッシュ、および LoRA ランクの割り当てが制御されます。残りの時間

ツールキット (ホットスワップ アダプター、5 つのテスト済みクライアント統合を備えたマルチプロトコル機能、ビジョン モデルへの画像入力、量子化、微調整、データセット、チャット ワークフロー用の OptIQ Lab GUI) がそのコアの周りにあります。
キャリブレーション データのレイヤーごとの KL がビットを選択します。敏感なレイヤーは高精度のままで、残りは低精度になり、均一 4 と同じ平均サイズになります。
KV キャッシュ上の個別の機密性パス。多くの場合、レイヤー 0 は平均よりも 56 倍感度が高いため、均一な 4 ビット KV は致命的です。混合精度はそうではありません。
各レイヤーのビットによってスケールされたアダプター ランクで微調整し、N 個のアダプターを 1 つのベースにマウントしたままにして、リクエストごとに切り替えます。リロードは必要ありません。
テキスト モデルを実行し、ビジョンを取得するモデルに写真を送信します。ベンダーのビジョン タワーは bf16 サイドカーに搭載されているため、1 つのリポジトリは mlx-lm ではテキストのみをロードするか、OptIQ では完全な画像 + テキストをロードします。ビジョンドキュメント。
optiqserve は、1 つのプロセスから OpenAI プロトコルと Anthropic プロトコルの両方を処理します。地元のクォントの Claude Code、Codex、OpenCode、OpenClaw、または Hermise エージェントを指定します。
ワークフロー全体の Web UI: 量子化ウィザード、データセット デザイナーによる SFT/DPO 微調整、サンドボックス ツール (Web 検索、Python、タ​​ーミナル) と画像アップロードによるチャット。
均一な 4 ビット量子化では、すべてのレイヤーが同じように扱われますが、レイヤーは同じではありません。 mlx-optiq は測定してから割り当てます。
各層について、各候補ビット幅でその層だけをシミュレート量子化します。
フォワードパス校正データ。摂動間の KL 発散を測定する
ロジットと参照ロジット。レイヤーごとに繰り返します。あなたは今持っています
(レイヤー、ビット) → 品質コストテーブル。
テーブルの上の貪欲なナップザック: すべてのレイヤーを最小のビット幅で開始し、
次に、追加ビットごとに最も多くの KL 削減を購入するレイヤーを貪欲にアップグレードします。
平均ビットバジェットが使い果たされるまで。 lm_head のようなレイヤー
そして、

最初/最後のアテンション ブロックはデフォルトで 8 ビットで保護されます。
レイヤーごとのビットマップをクォントとして mlx_lm.convert に渡します。
述語。出力は、どこにでもロードされる標準の MLX チェックポイントです。
ストック mlx-lm ロード、感度メタデータが隠蔽されている
下流の LoRA トレーニングの側。
# bf16 とユニフォーム 4 ビット参照間の自動ルーティング
# 利用可能な RAM に基づきます。
$ optiq 変換 Qwen/Qwen3.5-9B \
--target-bpw 5.0 \
--candidate-bits 4 、 8 \
--自動参照 \
-o optiq_output/Qwen3.5-9B
9 B モデルでは 1 ～ 2 分で実行されますが、27 B+ ではさらに長くなります
なぜこれがスケールするのか
単一のキャリブレーション主導の感度パス。 --自動参照
適合する場合は bf16 を選択し、適合しない場合は均一の 4 ビット ベースラインに戻ります。
bf16 ストリーミング プローブを使用するため、27 B+ モデルでもキャリブレーション主導の
36 GB Mac 上の信号。完全な方法論は次の場所にあります
私たちの研究報告書 →
05 比較してみると
mlx-optiq が Mac LLM オプションの中に位置します。
Apple Silicon 上で実際に品質と速度を向上させるものに人気のパスがどのように積み重なっているかを示すスナップショット。これらはどれも間違っていません。彼らはさまざまな軸を最適化しています。
Mac 上で LLM を実行すると、「」と答えました。
LLM を Mac 上でローカルに実行できますか?
はい。 mlx-optiq は、Apple の MLX フレームワークを使用して、M1 から M5 までの大規模な言語モデルを Apple Silicon 上でネイティブに実行します。 pip install mlx-optiq を使用して PyPI からインストールし、完全にオフラインで Mac 上でモデルを量子化、微調整、提供します。
Mac で LLM を実行するには GPU または PyTorch が必要ですか?
いいえ。パスには PyTorch も個別の GPU もありません。 mlx-optiq は MLX ネイティブで、Apple Silicon のユニファイド メモリを直接使用するため、MacBook、Mac mini、または Mac Studio で十分です。 CUDA、クラウド、API キーはありません。
Mac で LLM を実行するにはどれくらいの RAM が必要ですか?
モデルにより異なります。 4B モデルの 4 ビット OptIQ クオントには約 3 GB が必要です。 9Bn

約 6 GB が必要です。大規模な専門家混合モデルにはさらに多くの必要があります。混合精度 4 ビット量子化により、完全精度に近い品質を維持しながら、より大きなモデルを Mac のメモリに収めることができます。
LLM を Apple Silicon 上でローカルに量子化、微調整、提供するための MLX ネイティブ ツールキット。その核心はデータ駆動型の混合精度量子化です。各レイヤーの感度を測定し、レイヤーごとのビット幅を割り当てるため、量子化は同じサイズで均一な 4 ビットよりも高い品質を維持します。また、ローカル Web UI (OptIQ Lab) と OpenAI および Anthropic 互換サーバーも同梱されています。
Mac を LLM ワークステーションにします。
モデルを選択し、スニペットを取得して、出荷します。このドキュメントでは、サポートされているすべてのファミリー、レシピの微調整、OpenAI 互換のサービング スタックについて説明します。
Apple Silicon 上で LLM をネイティブに実行するための最適化コンパイラーとツールキット。 PyPI で配布されます。 pip install mlx-optiq を使用します。

## Original Extract

Run LLMs locally on your Mac (M1 to M5): mixed-precision quantization, LoRA fine-tuning, long-context KV serving, and image input, all MLX-native. On PyPI, no PyTorch, no GPU cluster, no cloud.

mlx-optiq
overview
models
docs
blog
pypi ↗
Optimizing compiler · MLX
Quantize, fine-tune
and serve LLMs
entirely on Apple Silicon.
Run large language models locally on your Mac, from M1 to M5.
Per-layer sensitivity analysis for mixed-precision weights.
LoRA fine-tuning that respects the bit budget.
A server that speaks both OpenAI and Anthropic APIs (point Claude Code at your
local quant). Send it an image, not just text, on any vision-capable model.
No GPU cluster, no API key.
8-bit · sensitive layers
4-bit · robust layers
$ pip install mlx-optiq
copy
3.1× avg compression vs bf16
+1.4× decode via MTP / drafter
+13.6 best Capability Score gain
140k+ HF downloads / month
OptIQ Lab — your Mac, your models, no cloud
OptIQ Lab · localhost:8080
A local LLM workbench, one pip install away: chat with sandboxed tools, compare models, fine-tune, and serve. Here a 4-bit OptIQ Qwen3.5 answers at 90 tok/s, fully offline .
01 Pre-built models
Drop-in 4-bit quants. Same weights, smarter bits.
Sixteen production mlx-optiq-quantized LLMs on Hugging Face. Nemotron 3, MiniCPM5, Qwen3.5, Qwen3.6 and Gemma-4 families, from 1 B dense to 35 B-A3B mixture-of-experts. They load directly into stock mlx-lm . No special runtime.
Google's unified text+vision Gemma-4, at 8.3 GB, with image input. Capability Score 68.2 (+6.4 vs uniform-4-bit), one of our largest mixed-precision gains, and the strongest model we ship under 9 GB on disk.
The largest single quant we ship. 31 B parameters in 20.8 GB with Capability Score 79.7 (+3.5 vs uniform-4-bit). Pair with the matching -assistant-bf16 drafter for speculative decoding.
Frontier-class reasoning at 17.5 GB with our highest Capability Score (83.0). Bundled MTP head gives ~1.4× decode via optiq serve --mtp .
The default daily-driver. 9 B parameters in 6.6 GB. Capability Score 66.8 (+0.2 vs uniform-4-bit). Long context to 64 k via mixed-precision KV; bundled MTP head for speculative decoding.
From zero to a serving LLM in three commands.
Each step is reversible and works with stock MLX tools. mlx-optiq is additive. Skip any of these and you still have a working pipeline.
Pure Python. Pulls in mlx , mlx-lm and huggingface-hub . Python 3.11+ on Apple Silicon.
$ pip install mlx-optiq
ii
Use a pre-built quant
Pre-built mlx-optiq quants load with stock mlx-lm . Per-layer bit assignment is recorded in the model metadata. No special loader required.
from mlx_lm import load, generate
model, tok = load ( "mlx-community/Qwen3.5-9B-OptiQ-4bit" )
out = generate (model, tok, prompt= "Explain mixed-precision quantization." , max_tokens= 200 )
print (out)
iii
Serve with mixed-precision KV
The KV cache is its own sensitivity problem. optiq kv-cache measures it once per model; optiq serve serves with the resulting per-layer config behind an OpenAI-compatible API.
# 1-2 min, once per model
$ optiq kv-cache mlx-community/Qwen3.5-9B-OptiQ-4bit \
--target-bits 5.0 -o ./kv
# OpenAI + Anthropic compatible server on :8080
# /v1/chat/completions (OpenAI)
# /v1/messages (Anthropic; works with Claude Code, anthropic SDK, etc.)
$ optiq serve --model mlx-community/Qwen3.5-9B-OptiQ-4bit \
--kv-config ./kv/kv_config.json \
--port 8080
Where to next
Each model family has a getting-started guide with model-specific sampling defaults and recommended use cases.
Building an agent? Drop llms.txt into your IDE. It's the entire library reference in one Markdown file.
03 What it does
One sensitivity signal. A whole toolkit around it.
A single per-layer KL-divergence pass drives weight, KV-cache and LoRA-rank allocation. The rest of the toolkit (hot-swap adapters, multi-protocol serving with five tested client integrations, image input on the vision models, and the OptIQ Lab GUI for quantize, fine-tune, dataset, and chat workflows) sits around that core.
Per-layer KL on calibration data picks the bits. Sensitive layers stay high-precision, the rest go low, at the same average size as uniform-4.
A separate sensitivity pass on the KV cache. Layer 0 is often 56× more sensitive than average, so uniform 4-bit KV is catastrophic; mixed-precision is not.
Fine-tune with adapter rank scaled by each layer's bits, then keep N adapters mounted on one base and switch them per request, no reload.
Run text models, and send pictures to the ones that take vision. The vendored vision tower rides in a bf16 sidecar, so one repo loads text-only under mlx-lm or full image+text under OptIQ. Vision docs .
optiq serve speaks both the OpenAI and Anthropic protocols from one process. Point Claude Code, Codex, OpenCode, OpenClaw, or Hermes Agent at a local quant.
A web UI for the whole workflow: quantize wizard, SFT/DPO fine-tuning with a dataset designer, and chat with sandboxed tools (web search, Python, terminal) and image upload.
Uniform 4-bit quantization treats every layer the same, but layers are not the same. mlx-optiq measures, then allocates.
For each layer, simulate-quantize just that layer at each candidate bit-width.
Forward-pass calibration data. Measure KL divergence between the perturbed
logits and the reference logits. Repeat for every layer; you now have a
(layer, bits) → quality cost table.
Greedy knapsack on the table: start every layer at the lowest bit-width,
then greedily upgrade the layer that buys the most KL-reduction per extra bit
until the average bit-budget is exhausted. Layers like lm_head
and the first/last attention blocks are protected at 8-bit by default.
Hand the per-layer bit map to mlx_lm.convert as a quant
predicate. The output is a standard MLX checkpoint that loads anywhere
stock mlx-lm loads, with sensitivity metadata stashed on
the side for downstream LoRA training.
# Auto-routes between bf16 and uniform-4-bit reference
# based on available RAM.
$ optiq convert Qwen/Qwen3.5-9B \
--target-bpw 5.0 \
--candidate-bits 4 , 8 \
--reference auto \
-o optiq_output/Qwen3.5-9B
runs in 1–2 min on a 9 B model · longer for 27 B+
Why this scales
A single calibration-driven sensitivity path. --reference auto
picks bf16 if it fits, otherwise falls back to a uniform-4-bit baseline
with bf16-streaming probes, so 27 B+ models still get a calibration-driven
signal on a 36 GB Mac. The full methodology lives in
our research write-up →
05 How it compares
Where mlx-optiq sits among the Mac LLM options.
A snapshot of how the popular paths stack up on the things that actually move quality and speed on Apple Silicon. None of these are wrong; they're optimizing different axes.
Running LLMs on a Mac, answered .
Can I run LLMs locally on a Mac?
Yes. mlx-optiq runs large language models natively on Apple Silicon, from M1 to M5, using Apple's MLX framework. Install it from PyPI with pip install mlx-optiq , then quantize, fine-tune, and serve models entirely on your Mac, fully offline.
Do I need a GPU or PyTorch to run LLMs on a Mac?
No. There is no PyTorch and no discrete GPU in the path. mlx-optiq is MLX-native and uses the unified memory of Apple Silicon directly, so a MacBook, Mac mini, or Mac Studio is enough. No CUDA, no cloud, no API key.
How much RAM do I need to run an LLM on a Mac?
It depends on the model. A 4-bit OptIQ quant of a 4B model needs roughly 3 GB; a 9B needs about 6 GB; larger mixture-of-experts models need more. Mixed-precision 4-bit quantization is what lets bigger models fit in a Mac's memory while staying close to full-precision quality.
An MLX-native toolkit to quantize, fine-tune, and serve LLMs locally on Apple Silicon. Its core is data-driven mixed-precision quantization: it measures each layer's sensitivity and assigns per-layer bit-widths, so quants keep more quality than uniform 4-bit at the same size. It also ships a local web UI (OptIQ Lab) and an OpenAI and Anthropic compatible server.
Make your Mac an LLM workstation .
Pick a model, get a snippet, ship it. The docs cover every supported family, fine-tuning recipes, and the OpenAI-compatible serving stack.
An optimizing compiler and toolkit for running LLMs natively on Apple Silicon. Distributed on PyPI; pip install mlx-optiq to use.
