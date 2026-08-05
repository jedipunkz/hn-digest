---
source: "https://github.com/mb-14/openstamp/"
hn_url: "https://news.ycombinator.com/item?id=49188741"
title: "OpenStamp – A watermarking method for open-weight LLMs"
article_title: "GitHub - mb-14/openstamp: A watermarking method for open-weight LLMs · GitHub"
author: "mb14"
captured_at: "2026-08-05T20:57:23Z"
capture_tool: "hn-digest"
hn_id: 49188741
score: 1
comments: 0
posted_at: "2026-08-05T20:43:57Z"
tags:
  - hacker-news
  - translated
---

# OpenStamp – A watermarking method for open-weight LLMs

- HN: [49188741](https://news.ycombinator.com/item?id=49188741)
- Source: [github.com](https://github.com/mb-14/openstamp/)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T20:43:57Z

## Translation

タイトル: OpenStamp – オープンウェイト LLM 用の透かし手法
記事のタイトル: GitHub - mb-14/openstamp: オープンウェイト LLM の透かし手法 · GitHub
説明: オープンウェイト LLM の透かし手法。 GitHub でアカウントを作成して、mb-14/openstamp の開発に貢献してください。

記事本文:
GitHub - mb-14/openstamp: オープンウェイト LLM の透かし手法 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
MB-14
/
オープンスタンプ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
160 コミット 160 コミット アセット アセットcolm_results colm_results Experiment_configs Experiment_configs 微調整 微調整 ノートブック ノートブック Saved_models_new Saved_models_new scripts スクリプト src src .gitignore .gitignore LIC

ENSE ライセンス METHOD.md METHOD.md README.md README.md 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オープンウェイト言語モデルの透かし。
ほとんどの LLM ウォーターマークは、生成時にトークン サンプリングを調整することで機能します。 API を制御する場合は問題ありませんが、ウェイトを解放すると、誰でも透かしをオフにして、滑らかなテキストを取得できるようになります。
したがって、オープン モデルの場合、ウォーターマークは重みの中に存在する必要があります。また、リリース後の一般的な変更 (出力の量子化、微調整、言い換え) にも耐える必要があります。既存のオープンウェイト法は、こうした条件下では機能しなくなることがよくあります。
OpenStamp は、非埋め込みレイヤーに小さな因数分解されたオフセットを追加します。リリースされたチェックポイントからサンプリングすると、特別なデコード時のロジックを使用せずに透かし入りのテキストが生成されます。検出するには、そのチェックポイントと基本モデルの非公開コピーの間で長さ正規化された対数尤度比を比較します。詳細は「仕組み」を参照してください。
OpenStamp は、品質と検出可能性のトレードオフのほぼ頂点に位置しており、低偽陽性率でほぼ完璧な検出を実現し、以前のオープンウェイト ベースラインと一致する複雑さを実現します。
攻撃を受けても完璧ではありません。 LLM 言い換え後、TPR@1%FPR は、Llama-2-7B では 1.0 付近から約 0.91、Mistral-7B では 0.79 に低下します。 LoRA の微調整により信号はさらに消耗します。OpenStamp は依然として GaussMark、KGW Distilled、Unremovable よりも優れていますが、微調整が続くと検出が悪化します。
conda create -n openstamp python=3.12 -y
conda は openstamp をアクティブ化します
pip install torch torchvision --index-url https://download.pytorch.org/whl/cu128
pip install flash-attn --no-build-isolation
pip install -r 要件.txt
モデルをダウンロードする
デフォルトでは、パイプラインは PPL オラクルとして metal-llama/Llama-2-13b-hf を使用し、言い換えには Qwen/Qwen2.5-14B-Instruct を使用します。 D

Hugging Face から両方を独自にロードします。
hf ダウンロード メタラマ/ラマ-2-13b-hf
hf ダウンロード Qwen/Qwen2.5-14B-Instruct
テストに使用したモデルをダウンロードします。
hf ダウンロード メタ-ラマ/ラマ-2-7b-hf
HF ダウンロード ミストラライ/Mistral-7B-v0.3
実験を実行する
透かし入りサンプルを生成し、検出を評価します。
Python スクリプト/run_config.py \
--config Experiment_configs/openstamp.yaml \
--base_output_dir 出力/メイン \
--サンプル数 500 \
--言い換え \
--eval_ppl
シード全体のメトリクスを CSV に集約します。
Python スクリプト/aggregate_metrics.py \
--input-dir 出力/メイン \
--output-csv 結果/aggregated_metrics.csv
構成は Experiment_configs/ にあります。
KGW+LLR (Kirchenbauer et al.、OpenStamp の LLR 検出器)
openstamp/phi-4-openstamp-L250-delta1.0-gamma0.25
openstamp/olmo-3-1025-7b-openstamp-L253-delta1.0-gamma0.25
openstamp/smollm2-1.7b-openstamp-L254-delta1.0-gamma0.25
openstamp/mistral-7b-v0.3-openstamp-L254-delta1.0-gamma0.25
openstamp/qwen2.5-7b-openstamp-L251-delta1.0-gamma0.25
openstamp/llama2-7b-openstamp-L254-delta1.0-gamma0.25
透かし入りのモデルをトランスフォーマーでロードし、通常どおり生成を呼び出します。透かしはすでに重みに含まれています。
輸入トーチ
トランスフォーマーからのインポート AutoModelForCausalLM 、 AutoTokenizer
MODEL_ID = "openstamp/llama2-7b-openstamp-L254-delta1.0-gamma0.25"
トークナイザー = AutoTokenizer 。 from_pretrained ( MODEL_ID )
トークナイザーの場合。 Pad_token は None です:
トークナイザー 。 Pad_token = トークナイザー 。 eos_token
モデル = AutoModelForCausalLM 。 from_pretrained (
モデルID 、
torch_dtype = トーチ。 bfloat16 、
device_map = "自動" ,
)
プロンプト = "むかしむかし、賢明な老賢者がいました。"
inputs = tokenizer (プロンプト、return_tensors = "pt")。へ (モデルとデバイス)
出力 = モデル 。生成 (** 入力、max_new_tokens = 256、do_sample = True、温度 = 0.7)
Watermarked_text = トーク

ナイザー。デコード (出力 [ 0 ]、skip_special_tokens = True )
印刷 (ウォーターマーク付きテキスト)
透かし入りテキストの検出
LLR スコアが選択したしきい値 (データに基づいて調整) を超えている場合、テキストは透かし入りとしてマークされます。
jsonをインポートする
輸入トーチ
ハギングフェイスハブからインポート hf_hub_download
トランスフォーマーからのインポート AutoModelForCausalLM 、 AutoTokenizer
ソースから 。 openstamp import OpenStamp 、モード
MODEL_REPO = "openstamp/llama2-7b-openstamp-L254-delta1.0-gamma0.25"
BASE_MODEL_ID = "メタ-ラマ/ラマ-2-7b-hf"
dtype = トーチ 。 bfloat16 の場合トーチ。クダ。 is_available() else トーチ。 float32
トークナイザー = AutoTokenizer 。 from_pretrained ( BASE_MODEL_ID )
トークナイザーの場合。 Pad_token は None です:
トークナイザー 。 Pad_token = トークナイザー 。 eos_token
モデル = AutoModelForCausalLM 。 from_pretrained (
BASE_MODEL_ID 、 torch_dtype = dtype 、 device_map = "auto"
)
open ( hf_hub_download ( MODEL_REPO , "watermark_config.json" )) を f として使用:
wm_cfg = json 。負荷 ( f )
Final_weight = トーチ 。負荷（
hf_hub_download ( MODEL_REPO , "selector_matrix.pth" )、map_location = "cpu"
)
ウォーターマーク = OpenStamp 。 from_config (
デルタ = wm_cfg [ "デルタ" ]、
ガンマ = wm_cfg [ "ガンマ" ],
シード = wm_cfg [「シード」]、
最終的な重み = 最終的な重み 、
モデル = モデル、
トークナイザー = トークナイザー 、
unembedding_param_name = "lm_head" ,
モード = モード。検出、
)
Watermarked_text = "<ここに透かしテキストを挿入>"
THRESHOLD = 0.0 # データセットで調整します
スコア = ウォーターマーク。スコア_テキスト_バッチ ([ウォーターマーク_テキスト])
llr = float (スコア [ 0 ])
is_watermarked = llr > しきい値
print ( llr 、 is_watermarked )
について
オープンウェイト LLM の透かし手法
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A watermarking method for open-weight LLMs. Contribute to mb-14/openstamp development by creating an account on GitHub.

GitHub - mb-14/openstamp: A watermarking method for open-weight LLMs · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
mb-14
/
openstamp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
160 Commits 160 Commits assets assets colm_results colm_results experiment_configs experiment_configs finetuning finetuning notebooks notebooks saved_models_new saved_models_new scripts scripts src src .gitignore .gitignore LICENSE LICENSE METHOD.md METHOD.md README.md README.md requirements.txt requirements.txt View all files Repository files navigation
Watermarking for open-weight language models.
Most LLM watermarks work by tweaking token sampling at generation time. That is fine when you control the API—but once you release the weights, anyone can turn the watermark off and still get fluent text.
So for open models, the watermark has to live in the weights. It also has to survive common post-release changes: quantization, fine-tuning, and paraphrasing of the output. Existing open-weight methods often break under those conditions.
OpenStamp adds a small, factorized offset to the unembedding layer. Sampling from the released checkpoint then produces watermarked text with no special decode-time logic. To detect, we compare a length-normalized log-likelihood ratio between that checkpoint and a privately held copy of the base model. Details are in How it works .
OpenStamp sits near the top of the quality–detectability tradeoff: near-perfect detection at low false-positive rates, with perplexity in line with earlier open-weight baselines.
It is not perfect under attack. After LLM paraphrasing, TPR@1%FPR drops from near 1.0 to about 0.91 on Llama-2-7B and 0.79 on Mistral-7B. LoRA fine-tuning wears the signal down further—OpenStamp still beats GaussMark, KGW Distilled, and Unremovable, but detection gets worse as fine-tuning continues:
conda create -n openstamp python=3.12 -y
conda activate openstamp
pip install torch torchvision --index-url https://download.pytorch.org/whl/cu128
pip install flash-attn --no-build-isolation
pip install -r requirements.txt
Download models
By default the pipeline uses meta-llama/Llama-2-13b-hf as the PPL oracle and Qwen/Qwen2.5-14B-Instruct for paraphrasing. Download both from Hugging Face:
hf download meta-llama/Llama-2-13b-hf
hf download Qwen/Qwen2.5-14B-Instruct
Download the models used for testing:
hf download meta-llama/Llama-2-7b-hf
hf download mistralai/Mistral-7B-v0.3
Run experiments
Generate watermarked samples and evaluate detection:
python scripts/run_config.py \
--config experiment_configs/openstamp.yaml \
--base_output_dir output/main \
--num_samples 500 \
--paraphrase \
--eval_ppl
Aggregate metrics across seeds into a CSV:
python scripts/aggregate_metrics.py \
--input-dir output/main \
--output-csv results/aggregated_metrics.csv
Configs live in experiment_configs/ :
KGW+LLR ( Kirchenbauer et al. ; LLR detector from OpenStamp)
openstamp/phi-4-openstamp-L250-delta1.0-gamma0.25
openstamp/olmo-3-1025-7b-openstamp-L253-delta1.0-gamma0.25
openstamp/smollm2-1.7b-openstamp-L254-delta1.0-gamma0.25
openstamp/mistral-7b-v0.3-openstamp-L254-delta1.0-gamma0.25
openstamp/qwen2.5-7b-openstamp-L251-delta1.0-gamma0.25
openstamp/llama2-7b-openstamp-L254-delta1.0-gamma0.25
Load a watermarked model with transformers and call generate as usual—the watermark is already in the weights.
import torch
from transformers import AutoModelForCausalLM , AutoTokenizer
MODEL_ID = "openstamp/llama2-7b-openstamp-L254-delta1.0-gamma0.25"
tokenizer = AutoTokenizer . from_pretrained ( MODEL_ID )
if tokenizer . pad_token is None :
tokenizer . pad_token = tokenizer . eos_token
model = AutoModelForCausalLM . from_pretrained (
MODEL_ID ,
torch_dtype = torch . bfloat16 ,
device_map = "auto" ,
)
prompt = "Once upon a time there was a wise old sage who"
inputs = tokenizer ( prompt , return_tensors = "pt" ). to ( model . device )
outputs = model . generate ( ** inputs , max_new_tokens = 256 , do_sample = True , temperature = 0.7 )
watermarked_text = tokenizer . decode ( outputs [ 0 ], skip_special_tokens = True )
print ( watermarked_text )
Detect watermarked text
Text is marked as watermarked when the LLR score is above a threshold you choose (calibrate it on your data).
import json
import torch
from huggingface_hub import hf_hub_download
from transformers import AutoModelForCausalLM , AutoTokenizer
from src . openstamp import OpenStamp , Mode
MODEL_REPO = "openstamp/llama2-7b-openstamp-L254-delta1.0-gamma0.25"
BASE_MODEL_ID = "meta-llama/Llama-2-7b-hf"
dtype = torch . bfloat16 if torch . cuda . is_available () else torch . float32
tokenizer = AutoTokenizer . from_pretrained ( BASE_MODEL_ID )
if tokenizer . pad_token is None :
tokenizer . pad_token = tokenizer . eos_token
model = AutoModelForCausalLM . from_pretrained (
BASE_MODEL_ID , torch_dtype = dtype , device_map = "auto"
)
with open ( hf_hub_download ( MODEL_REPO , "watermark_config.json" )) as f :
wm_cfg = json . load ( f )
final_weight = torch . load (
hf_hub_download ( MODEL_REPO , "selector_matrix.pth" ), map_location = "cpu"
)
watermark = OpenStamp . from_config (
delta = wm_cfg [ "delta" ],
gamma = wm_cfg [ "gamma" ],
seed = wm_cfg [ "seed" ],
final_weight = final_weight ,
model = model ,
tokenizer = tokenizer ,
unembedding_param_name = "lm_head" ,
mode = Mode . Detect ,
)
watermarked_text = "<insert watermarked text here>"
THRESHOLD = 0.0 # calibrate on your dataset
scores = watermark . score_text_batch ([ watermarked_text ])
llr = float ( scores [ 0 ])
is_watermarked = llr > THRESHOLD
print ( llr , is_watermarked )
About
A watermarking method for open-weight LLMs
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
