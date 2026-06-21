---
source: "https://FareedKhan-dev.github.io/train-llm-from-scratch/"
hn_url: "https://news.ycombinator.com/item?id=48615416"
title: "Train LLM from Scratch"
article_title: "Train LLM From Scratch"
author: "vinhnx"
captured_at: "2026-06-21T04:35:39Z"
capture_tool: "hn-digest"
hn_id: 48615416
score: 1
comments: 0
posted_at: "2026-06-21T03:30:18Z"
tags:
  - hacker-news
  - translated
---

# Train LLM from Scratch

- HN: [48615416](https://news.ycombinator.com/item?id=48615416)
- Source: [FareedKhan-dev.github.io](https://FareedKhan-dev.github.io/train-llm-from-scratch/)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T03:30:18Z

## Translation

タイトル: LLM をゼロから訓練する
記事のタイトル: LLM をゼロからトレーニングする
説明: 事前トレーニングから RLHF/GRPO まで、すべてのアルゴリズムは純粋な PyTorch で手書きされています。

記事本文:
LLM を最初からトレーニングする
コンテンツにスキップ
LLM を最初からトレーニングする
ホーム
検索を初期化しています
FareedKhan-dev/train-llm-from-scratch
ホーム
LLM を最初からトレーニングする
FareedKhan-dev/train-llm-from-scratch
ホーム
ホーム
目次
推奨される読む順序
唯一の設計ルール: ラップし、書き換えない
色の凡例 (これらのドキュメントのすべての図で使用されています)
基礎
基礎
トークン化とデータ形状
目的、損失、困惑
最適化とトレーニング システム
理論とパイプライン
理論とパイプライン
データの取り扱い
ハウツー
ハウツー
設定(JSON)
参考資料
参考資料
コマンドチートシート
唯一の設計ルール: ラップし、書き換えない
色の凡例 (これらのドキュメントのすべての図で使用されています)
トレーニング後の調整と調整 — 概要 ¶
このトランスフォーマーを最初からトレーニングしたとき、テキストを継続することはできましたが、それはできませんでした。
指示または理由に従ってください。それはトレーニング後の修正です。このドキュメント/フォルダーは歩きます
基本モデルの上に構築した旅全体を通して、すべての段階をゼロから作成しました
プレーンな PyTorch ( trl 、 peft 、transformers なし) で、実際の公開データセットでトレーニングされ、
単一の GPU で実行することも、DDP を使用して複数の GPU にまたがって拡張することもできます。
LLM トレーニング内部を初めて使用する場合は、新しいものから始めてください。
ステージのページを読む前に、LLM Foundations セクションを参照してください。それは、
トークンの形状、デコーダー専用のトランスフォーマー、アテンション マスク、目標、最適化ループ、および生成
後のすべてのページが依存する仕組み。
まず基礎:
トークン化 ->
変圧器 ->
注意 ->
目的 ->
最適化 ->
世代。
次に、完全なパイプライン:
データ ->
事前トレーニング ->
SFT ->
報酬モデル ->
DPO ->
PPO ->
GRPO 。
最後に実行して検査します。
評価、推論/チャット、そして
コマンドチートシート 。
このパイプラインは、最新の整合/推論モデルが実際にどのように構築されるかを反映しています。
f

ローチャート TD
PILE([パイル<br/>98 億トークン]):::データ --> PRE{{事前トレーニング<br/>~4 億ベース}}:::モデル
PRE --> BASE[(base_pretrained.pt)]:::ckpt
BASE --> SFT{{SFT<br/>アルパカ・ドリー・GSM8K}}:::モデル
SFT --> SFTCK[(sft.pt)]:::ckpt
SFTCK --> RM{{報酬モデル<br/>ブラッドリー・テリー}}:::rl
SFTCK --> DPO{{DPO / ORPO / KTO<br/>設定}}:::rl
RM --> RMCK[(reward.pt)]:::ckpt
RMCK -->|報酬信号| PPO{{PPO<br/>GAE + クリップ + KL}}:::rl
SFTCK --> PPO
SFTCK --> GRPO{{GRPO / RLVR<br/>グループ相対}}:::rl
PPO --> EVAL([GSM8K eval<br/>+ チャット / 推論]):::eval
DPO --> 評価
GRPO --> 評価
classDef データの塗りつぶし:#d6ffd9、ストローク:#27ae60、ストローク幅:2px、色:#143d1a;
classDef モデルの塗りつぶし:#ffe8a3、ストローク:#d48806、ストローク幅:2px、色:#5a3d00;
classDef rl 塗りつぶし:#ffd9b3、ストローク:#e67e22、ストローク幅:2px、色:#6b3500;
classDef ckpt fill:#eeeeee、ストローク:#555555、ストローク幅:2px、色:#222;
classDef eval fill:#e8d6ff、ストローク:#8e44ad、ストローク幅:2px、color:#3d1a5a;
ステージの順序 ¶
1 つの設計ルール: ラップし、書き換えない ¶
ここにあるものはすべて、オリジナルの Transformer の上にあります。変更しました
教育モデルを正確に 1 か所に — forward_hidden を追加しました
lm_head が消費する最終的な隠し状態を返すメソッド。トレーニング後のすべての頭 (値
PPO の場合はヘッド、報酬モデルの場合はスカラー報酬ヘッド)、およびすべての RL 対数確率計算が構成されます。
その 1 つのメソッドを中心にしているため、すでに理解している最初からのモデルはそのまま残ります。
色の凡例 (これらのドキュメントのすべての図で使用) ¶
🟩 データ / コーパス · 🟦 前処理 · ‍⬛ ストレージ (HDF5 / JSONL) · 🟨 モデル / トレーニング ループ
・🟧RL/報酬・🟥損失/目標・🟪評価・⬜チェックポイント
各図は手描きで色分けされた人魚のスケッチで、事前に PNG にレンダリングされ、次のように埋め込まれています。
画像 (GitHub のライブ

e Mermaid は確実に look するわけではありません: handDrawn 、および一部のビューア — 例:の
VS Code プレビュー — SVG をブロックして、埋め込まれた PNG がどこにでも表示されるようにします。編集可能な Mermaid ソースが置かれています
各画像の下にある折りたたみ可能な「マーメイド ソース」ブロック内にあります。編集後の画像を再生成するには、
図/README.md を参照してください。
基本モデルが事前トレーニング ( 02_pretraining.md ) されると、チェーン全体が 1 つのスクリプトになります。
bash scripts/run_posttraining.sh # SFT -> RM -> DPO -> PPO -> GRPO -> 評価テーブル
要約されたコマンド リファレンスについては、POST_TRAINING.md を参照してください。

## Original Extract

From pretraining to RLHF/GRPO — every algorithm hand-written in pure PyTorch.

Train LLM From Scratch
Skip to content
Train LLM From Scratch
Home
Initializing search
FareedKhan-dev/train-llm-from-scratch
Home
Train LLM From Scratch
FareedKhan-dev/train-llm-from-scratch
Home
Home
Table of contents
Recommended reading order
The one design rule: wrap, don't rewrite
Colour legend (used in every diagram in these docs)
Foundations
Foundations
Tokenization & Data Shapes
Objectives, Losses & Perplexity
Optimization & Training Systems
Theory & Pipeline
Theory & Pipeline
Data handling
How-to
How-to
Configure (JSON)
Reference
Reference
Command cheatsheet
The one design rule: wrap, don't rewrite
Colour legend (used in every diagram in these docs)
Post-Training & Alignment — Overview ¶
When I first trained this transformer from scratch, it could continue text but it couldn't
follow instructions or reason . That's what post-training fixes. This docs/ folder walks
through the whole journey I built on top of the base model — every stage written from scratch
in plain PyTorch (no trl , no peft , no transformers ), trained on real public datasets, and
runnable on a single GPU or scaled across multiple GPUs with DDP.
If you are new to LLM training internals, start with the new
LLM Foundations section before reading the stage pages. It explains the
token shapes, decoder-only Transformer, attention masks, objectives, optimization loop, and generation
mechanics that every later page relies on.
Foundations first :
Tokenization ->
Transformer ->
Attention ->
Objectives ->
Optimization ->
Generation .
Then the full pipeline :
Data ->
Pretraining ->
SFT ->
Reward Model ->
DPO ->
PPO ->
GRPO .
Finally run and inspect :
Evaluation , Inference / Chat , and the
command cheatsheet .
The pipeline mirrors how modern aligned/reasoning models are actually built:
flowchart TD
PILE([The Pile<br/>9.8B tokens]):::data --> PRE{{Pretrain<br/>~400M base}}:::model
PRE --> BASE[(base_pretrained.pt)]:::ckpt
BASE --> SFT{{SFT<br/>Alpaca · Dolly · GSM8K}}:::model
SFT --> SFTCK[(sft.pt)]:::ckpt
SFTCK --> RM{{Reward Model<br/>Bradley-Terry}}:::rl
SFTCK --> DPO{{DPO / ORPO / KTO<br/>preference}}:::rl
RM --> RMCK[(reward.pt)]:::ckpt
RMCK -->|reward signal| PPO{{PPO<br/>GAE + clip + KL}}:::rl
SFTCK --> PPO
SFTCK --> GRPO{{GRPO / RLVR<br/>group-relative}}:::rl
PPO --> EVAL([GSM8K eval<br/>+ chat / inference]):::eval
DPO --> EVAL
GRPO --> EVAL
classDef data fill:#d6ffd9,stroke:#27ae60,stroke-width:2px,color:#143d1a;
classDef model fill:#ffe8a3,stroke:#d48806,stroke-width:2px,color:#5a3d00;
classDef rl fill:#ffd9b3,stroke:#e67e22,stroke-width:2px,color:#6b3500;
classDef ckpt fill:#eeeeee,stroke:#555555,stroke-width:2px,color:#222;
classDef eval fill:#e8d6ff,stroke:#8e44ad,stroke-width:2px,color:#3d1a5a;
The stages, in order ¶
The one design rule: wrap, don't rewrite ¶
Everything here sits on top of the original Transformer . I changed the
educational model in exactly one place — I added a forward_hidden
method that returns the final hidden states the lm_head consumes. Every post-training head (a value
head for PPO, a scalar reward head for the reward model) and every RL log-prob computation composes
around that one method, so the from-scratch model you already understand stays intact.
Colour legend (used in every diagram in these docs) ¶
🟩 data / corpus · 🟦 preprocessing · 🟦‍⬛ storage (HDF5 / JSONL) · 🟨 model / training loop
· 🟧 RL / reward · 🟥 loss / objective · 🟪 evaluation · ⬜ checkpoint
Each diagram is a hand-drawn, colour-coded Mermaid sketch, pre-rendered to a PNG and embedded as
an image (GitHub's live Mermaid doesn't reliably do look: handDrawn , and some viewers — e.g. the
VS Code preview — block SVGs, so an embedded PNG shows everywhere). The editable Mermaid source sits
in a collapsible "Mermaid source" block under each image. To regenerate the images after editing,
see diagrams/README.md .
Once the base model has pretrained ( 02_pretraining.md ), the entire chain is one script:
bash scripts/run_posttraining.sh # SFT -> RM -> DPO -> PPO -> GRPO -> eval table
See POST_TRAINING.md for the condensed command reference.
