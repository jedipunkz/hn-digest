---
source: "https://mlx-optiq.com/blog/humanizer-stacked-lora"
hn_url: "https://news.ycombinator.com/item?id=48353832"
title: "A 1B humanizer that matches human writing on an AI detector"
article_title: "A 1B humanizer that matches human writing on an AI detector · mlx-optiq"
author: "codelion"
captured_at: "2026-06-02T00:44:26Z"
capture_tool: "hn-digest"
hn_id: 48353832
score: 2
comments: 0
posted_at: "2026-06-01T07:55:53Z"
tags:
  - hacker-news
  - translated
---

# A 1B humanizer that matches human writing on an AI detector

- HN: [48353832](https://news.ycombinator.com/item?id=48353832)
- Source: [mlx-optiq.com](https://mlx-optiq.com/blog/humanizer-stacked-lora)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T07:55:53Z

## Translation

タイトル: AI 検出器で人間の書き込みと一致する 1B ヒューマナイザー
記事のタイトル: AI 検出器で人間の書き込みと一致する 1B ヒューマナイザー · mlx-optiq
説明: MiniCPM5-1B-OptiQ-4bit 上のスタックされた SFT + DPO LoRA は、RADAR AI 検出器上の人間の書き込みとのギャップを 100% 近づけます (P(AI) 0.51 ～ 0.37、人間の基準と一致)。レシピは OptIQ 0.1.4 を使用します

記事本文:
mlx-optiq
概要
モデル
ドキュメント
ブログ
ピピ↗
エンジニアリング · 2026 年 6 月 1 日
AI 検出器上で人間の書き込みを照合する 1B ヒューマナイザー。
1B モデルにスタックされた 2 つの LoRA アダプターは、R​​ADAR AI 検出器上の人間の書き込みとのギャップを 100% 近づけます。ソース AI ドラフトは RADAR から 0.51 を取得します。 EditLens 人間の場合は 0.37 になります。 SFT + DPO LoRA を MiniCPM5-1B-OptIQ-4bit に一緒に適用すると、同じ 200 の保留されたドラフトでの書き換えも 0.37 として返されます。スタック全体が 24 GB Mac 上で実行されます。
EditLens ICLR 2026 コーパスからの 200 件の保留された AI ドラフトを各システムによって書き換えた後、 RADAR-Vicuna-7B によってスコア付けしました。 P(AI)が低いほど、より人間らしいことを意味します。
スロップ フレーズの頻度 (「～の証拠」や「～の重要性を強調する」などの定型パターン) は、ソースの 1K トークンあたり 0.6 からスタック出力では 0.0 に低下します。人間の基準セットには 0.1 があります。
基本モデルはディスク上に 875 MB を占有し、各アダプタにはさらに 120 MB が必要です。 70B モデルはなく、API キーも必要ありません。
OptIQ 0.1.4 はすべての部分を出荷します。完全なパイプラインは次のとおりです。
オプションの最初の目的地: ラボのデータセット ビルダー。スタイル転送 + DPO from プリファレンス ペア テンプレートにより、この実行の SFT および DPO データセットが生成されました。
ターミナル・1.bashの量子化
$ pip install 'mlx-optiq>=0.1.4'
$ optiq 変換 openbmb/MiniCPM5-1B \
--target-bpw 5.0 --candidate-bits 4,8 \
--output ./optiq_mixed
感度を考慮した混合精度量子化。ほとんどのレイヤーは 4 ビットに到達し、重要なレイヤーは 8 ビットに到達します。結果は 875 MB で、能力スコア ( eval フレームワーク ) に基づく bf16 ベースにはわずか 1.06 GB 足りません。
$ optiq lora train ./optiq_mixed \
--data ./sft_dataset --method sft \
--presetlarge --iters 600 \
--output ./adapters/humanizer-sft
$ optiq lora train ./optiq_mixed \
--data ./dpo_dataset --method dpo \
--presetlarge --iters 300 \
--mount-adapter ./adapters/humanizer-sf

t \
--output ./adapters/humanizer-dpo
--mount-adapter は、標準の SFT と DPO の継続レシピです。凍結された SFT LoRA を、適応されたすべてのレイヤー上でトレーニング可能な DPO LoRA と並べてスタックします。 DPO 参照フォワード ゼロはトレーニング可能なスケールのみであるため、損失内の KL 項は SFT モデルであるベース + SFT に対して固定されます。これは、すべての最新のアライメント パイプラインが「SFT から続く DPO」を定義する方法と一致します。保存されたアダプターには、DPO デルタのみが保持されます。
OptIQ Lab 微調整ウィザードのハイパーパラメータ ステップ。 DPO を選択すると、学習率が 5e-5 に切り替わり、DPO のデフォルト バナーが表示されます。
ターミナル · 3. 両方を提供し、リクエストごとにスタックされた bash
$ optiqserve --model ./optiq_mixed \
--adapter ./adapters/humanizer-sft \
--adapter ./adapters/humanizer-dpo
# リクエスト本文は「+」演算子の両方でアクティブ化されます。
$カールローカルホスト:8080/v1/chat/completions \
-d '{"モデル":"...","メッセージ":[...],
"アダプター":"ヒューマナイザー-sft+ヒューマナイザー-dpo"}'
optiq サーブは、両方のアダプターを同じベース モデルにマウントします。リクエスト本文のアダプター フィールドは、a+b 形式を指定すると、両方の LoRA 残差を切り替えるのではなく、単一のフォワード パス中に両方の LoRA 残差を適用し、モデルのリロードは行われません。古典的な単一アダプター構文 ( "adapter": "humanizer-sft" ) は引き続き機能します。センチネル "adapter": "base" はアダプターのアクティブ化を完全にバイパスします。同じサービス対象プロセスからの A/B 比較に使用します。
同じマルチアダプターマウントが実験室に現れました。 「設定」→「サーバー」には、登録されているアダプターがリストされます。マウントされると、チャット画面からリクエストごとに選択されます。
スタックがどちらのアダプター単独よりも優れている理由
SFT アダプター単独のスコアは P(AI) = 0.50 で、ソースよりかろうじて優れています。 DPO アダプターだけでは意味がありません。これは、絶対的な LoRA ではなく、SFT からのデルタとしてトレーニングされました。 SFT をアクティブにしないと、小さな摂動が適用されます。

SFT 分布をまったく回復しない基本モデルに変換します。
スタックはトレーニング時間のフォワード パスを再現します。
# トレーニング中 (凍結された SFT + トレーニング可能な DPO):
y = ベース(x) + sft_scale * (x @ sft_a @ sft_b)
+ dpo_scale * ((ドロップアウト(x) @ lora_a) @ lora_b)
# 提供時 (両方ともマウント、「a+b」構文):
y = ベース(x) + sft_scale * (x @ sft_a @ sft_b)
+ dpo_scale * (x @ lora_a @ lora_b)
したがって、トレーニング時の計算と推論時の計算は、同じ重みを使用して同じ式になるため、保持された P(AI) はトレーニング軌跡の予測と一致します。
CLI が苦手な場合は、Lab UI が同じパイプラインを経由します。トレーニングが終了すると、ウィザードの最後のステップで、アダプターの結合、モデルのエクスポート、Hugging-Face へのプッシュが 3 つのオプションのチェックボックスとして表示されます。
微調整ウィザードのステップ 5。 Combine は 2 つのアダプターをマージし、Bundle はモデル ディレクトリをエクスポートし、Push はモデル ディレクトリをハブに送信します。
試してみてください
すべて (ベース、両方のアダプター、モデル カード、ホールドアウト評価) は単一の Hugging Face リポジトリ: mlx-community/humanizer-1B-OptIQ-4bit (~1.1 GB) にバンドルされています。一度ダウンロードすると、両方のアダプターをスタックして使用できます。
$ pip install 'mlx-optiq>=0.1.4'
$ ハグフェイス-cli ダウンロード mlx-community/humanizer-1B-OptIQ-4bit \
--local-dir ./humanizer-1B-OptIQ-4bit
$ オプティクサーブ \
--model ./humanizer-1B-OptIQ-4bit \
--adapter ./humanizer-1B-OptIQ-4bit/adapters/humanizer-sft \
--adapter ./humanizer-1B-OptIQ-4bit/adapters/humanizer-dpo
localhost:8080 上の OpenAI 互換エンドポイント。任意のクライアント (WebUI を開き、続行、独自のスクリプト) を指定し、リクエスト本文で "adapter": "humanizer-sft+humanizer-dpo" を送信すると、ローカルのヒューマナイザーが完成します。
これをロック解除した OptIQ の部分は --mount-adapter とマルチアダプター サーブで、どちらも v0.1.4 に同梱されています。他の SFT でも機能します

それから、単なる人間化ではなく、DPO の継続。 SFT レシピと任意のタスク (コード スタイル、ブランド ボイス、拒否行動など) の設定データセットがある場合、同じ 2 つのコマンドで、最初から開始するのではなく、SFT から継続する DPO LoRA が得られます。
トレーナー、すべての --preset オプション、およびデータセット形式のリファレンスは、LoRA 微調整ガイドにあります。マルチアダプターのserveおよび「a+b」スタッキング構文は、serve docs にあります。
Apple Silicon 上で LLM をネイティブに実行するための最適化コンパイラーとツールキット。 PyPI で配布されます。 pip install mlx-optiq を使用します。

## Original Extract

Stacked SFT + DPO LoRAs on MiniCPM5-1B-OptiQ-4bit close 100% of the gap to human writing on the RADAR AI detector (P(AI) 0.51 to 0.37, matching the human reference). The recipe uses OptIQ 0.1.4

mlx-optiq
overview
models
docs
blog
pypi ↗
Engineering · June 1, 2026
A 1B humanizer that matches human writing on an AI detector.
Two LoRA adapters stacked on a 1B model close 100 % of the gap to human writing on the RADAR AI detector. Source AI drafts get 0.51 from RADAR. EditLens humans get 0.37. With the SFT + DPO LoRAs applied together on MiniCPM5-1B-OptIQ-4bit , the rewrites also come back as 0.37, on the same 200 held-out drafts. Whole stack runs on a 24 GB Mac.
200 held-out AI drafts from the EditLens ICLR 2026 corpus, rewritten by each system, then scored by RADAR-Vicuna-7B . Lower P(AI) means more human-like.
Slop-phrase frequency (boilerplate patterns like "a testament to" and "underscores the importance of") drops from 0.6 per 1K tokens in the source to 0.0 in the stacked output. The human reference set has 0.1.
The base model takes 875 MB on disk, each adapter another 120 MB. No 70B model and no API key required.
OptIQ 0.1.4 ships every piece. The full pipeline is:
Optional first stop: the Lab's dataset builder. The Style transfer + DPO from preference pairs templates produced the SFT and DPO datasets for this run.
terminal · 1. quantize bash
$ pip install 'mlx-optiq>=0.1.4'
$ optiq convert openbmb/MiniCPM5-1B \
--target-bpw 5.0 --candidate-bits 4,8 \
--output ./optiq_mixed
Sensitivity-aware mixed-precision quantization. Most layers land at 4-bit, the sensitive ones at 8-bit. Result is 875 MB and only 1.06 GB short of the bf16 base on Capability Score ( eval framework ).
$ optiq lora train ./optiq_mixed \
--data ./sft_dataset --method sft \
--preset large --iters 600 \
--output ./adapters/humanizer-sft
$ optiq lora train ./optiq_mixed \
--data ./dpo_dataset --method dpo \
--preset large --iters 300 \
--mount-adapter ./adapters/humanizer-sft \
--output ./adapters/humanizer-dpo
--mount-adapter is the standard SFT then DPO continuation recipe. It stacks a frozen SFT LoRA alongside a trainable DPO LoRA on every adapted layer. The DPO reference forward zeroes only the trainable scale, so the KL term in the loss is anchored against base + SFT, which is the SFT model. That matches how every modern alignment pipeline defines "DPO continuing from SFT". The saved adapter holds only the DPO delta.
OptIQ Lab Fine-tune wizard, Hyperparameters step. Picking DPO swaps the learning rate to 5e-5 and surfaces the DPO defaults banner.
terminal · 3. serve both, stacked per-request bash
$ optiq serve --model ./optiq_mixed \
--adapter ./adapters/humanizer-sft \
--adapter ./adapters/humanizer-dpo
# request body activates both with the "+" operator:
$ curl localhost:8080/v1/chat/completions \
-d '{"model":"...","messages":[...],
"adapter":"humanizer-sft+humanizer-dpo"}'
optiq serve mounts both adapters on the same base model. The request body's adapter field, given an a+b form, applies both LoRA residuals during a single forward pass instead of switching between them, and there is no model reload. The classic single-adapter syntax ( "adapter": "humanizer-sft" ) still works. The sentinel "adapter": "base" bypasses adapter activation entirely. Use it for A/B comparisons from the same served process.
The same multi-adapter mount, surfaced in the Lab. Settings → Server lists the registered adapters; once mounted they are picked per-request from the Chat surface.
Why the stack beats either adapter alone
The SFT adapter alone scores P(AI) = 0.50, barely better than the source. The DPO adapter on its own is meaningless. It was trained as a delta from SFT, not an absolute LoRA. Without SFT active, you're applying a small perturbation to the base model that doesn't recover the SFT distribution at all.
The stack reproduces the training-time forward pass:
# during training (frozen SFT + trainable DPO):
y = base(x) + sft_scale * (x @ sft_a @ sft_b)
+ dpo_scale * ((dropout(x) @ lora_a) @ lora_b)
# at serve time (both mounted, "a+b" syntax):
y = base(x) + sft_scale * (x @ sft_a @ sft_b)
+ dpo_scale * (x @ lora_a @ lora_b)
So the math at training time and the math at inference time work out to the same expression with the same weights, which is why the held-out P(AI) lines up with what the training trajectory predicted.
If the CLI isn't your thing, the Lab UI walks through the same pipeline. After training finishes, the wizard's last step exposes adapter combine + model export + push-to-Hugging-Face as three optional checkboxes.
Step 5 of the Fine-tune wizard. Combine merges two adapters, Bundle exports a model directory, Push ships it to the Hub.
Try it
Everything (base, both adapters, model card, held-out eval) is bundled into a single Hugging Face repo: mlx-community/humanizer-1B-OptIQ-4bit (~1.1 GB). Download once, serve with both adapters stacked:
$ pip install 'mlx-optiq>=0.1.4'
$ huggingface-cli download mlx-community/humanizer-1B-OptIQ-4bit \
--local-dir ./humanizer-1B-OptIQ-4bit
$ optiq serve \
--model ./humanizer-1B-OptIQ-4bit \
--adapter ./humanizer-1B-OptIQ-4bit/adapters/humanizer-sft \
--adapter ./humanizer-1B-OptIQ-4bit/adapters/humanizer-dpo
OpenAI-compatible endpoint on localhost:8080 . Point any client (Open WebUI, Continue, your own scripts) at it, send "adapter": "humanizer-sft+humanizer-dpo" in the request body, and you have a local humanizer.
The piece of OptIQ that unlocked this is --mount-adapter plus the multi-adapter serve, both shipped in v0.1.4. They also work for any other SFT then DPO continuation, not just humanization. If you have an SFT recipe and a preference dataset for any task (code style, brand voice, refusal behavior, anything), the same two commands give you a DPO LoRA that continues from your SFT instead of starting from scratch.
Reference for the trainer, every --preset option, and the dataset format is in the LoRA fine-tuning guide . The multi-adapter serve and the "a+b" stacking syntax are in the serve docs .
An optimizing compiler and toolkit for running LLMs natively on Apple Silicon. Distributed on PyPI; pip install mlx-optiq to use.
