---
source: "https://github.com/johnhalloran321/mcp_safety_training"
hn_url: "https://news.ycombinator.com/item?id=48834106"
title: "Refusal training for LLM agents against disguised MCP attacks"
article_title: "GitHub - johnhalloran321/mcp_safety_training: DPO/SafeDPO/OPAD training + eval for teaching tool-using LLMs to refuse falsely-benign MCP exploits · GitHub"
author: "jhalloran"
captured_at: "2026-07-08T17:20:28Z"
capture_tool: "hn-digest"
hn_id: 48834106
score: 1
comments: 0
posted_at: "2026-07-08T16:37:30Z"
tags:
  - hacker-news
  - translated
---

# Refusal training for LLM agents against disguised MCP attacks

- HN: [48834106](https://news.ycombinator.com/item?id=48834106)
- Source: [github.com](https://github.com/johnhalloran321/mcp_safety_training)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T16:37:30Z

## Translation

タイトル: 偽装 MCP 攻撃に対する LLM エージェントの拒否トレーニング
記事のタイトル: GitHub - johnhalloran321/mcp_safety_training: DPO/SafeDPO/OPAD トレーニング + 教育ツールの評価 - 誤って無害な MCP エクスプロイトを拒否するための LLM 使用 · GitHub
説明: DPO/SafeDPO/OPAD トレーニング + 誤って無害な MCP エクスプロイトを拒否するための LLM を使用する教育ツールの評価 - johnhalloran321/mcp_safety_training

記事本文:
GitHub - johnhalloran321/mcp_safety_training: DPO/SafeDPO/OPAD トレーニング + 誤って無害な MCP エクスプロイトを拒否するための LLM を使用する教育ツールの評価 · GitHub
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
ジョンハロラン321
/
mcp_safety_training
公共
通知
c にサインインする必要があります

ハンゲの通知設定
追加のナビゲーション オプション
コード
ジョンハロラン321/mcp_safety_training
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット アセット .gitignore .gitignore ライセンス ライセンス README.md README.md dpo.py dpo.py make_rag_dbs.py make_rag_dbs.py mcp_judge_cache.py mcp_judge_cache.py mcp_test_cache.py mcp_test_cache.py my_jailbreak_classifiers.py my_jailbreak_classifiers.py opad.py opad.py opad_dataset.py opad_dataset.py opad_utils.py opad_utils.py Prompt.py Prompt.py rag_pref.py rag_pref.py要件.txt要件.txtsafedpo.pysafedpo.pysafedpo_trainer.pysafedpo_trainer.pytools.py tools.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ハロラン、ジョン。 「MCP 安全トレーニング: 改善された優先順位の調整を使用して、誤って無害な MCP エクスプロイトを拒否する方法を学習します。」 arXiv:2505.23634 (2025)。
Halloran、John T.「トレーニング不要の LLM 調整のための RAG の活用」。 arXiv:2605.11217 (2026)。
誤って無害な MCP エクスプロイト (FBA) に対してツールを使用する LLM を調整するための DPO / SafeDPO トレーニングおよび評価コード — CVE 由来のモデル コンテキスト プロトコル ツールを使用した攻撃で、通常の無害な要求のように表現されます。
安全調整されたモデル (1B ～ 14B パラメーター) は、そのままの状態で 35% を超える FBA を拒否しました。標準の DPO/SafeDPO 調整では、この値はせいぜい 48% にとどまりました。このコードと並行して提案されているトレーニング不要の検索ベースの手法である RAG-Pref は、単独で拒否率が最大 3 倍、DPO/SafeDPO と組み合わせると最大 3.7 倍向上します。完全な数値については、arXiv:2505.23634 および arXiv:2605.11217 を参照してください。
arXiv:2605.11217 の図。表示されているすべてのバーはこのリポジトリに実装されています: Base / DPO / SafeDPO ( dpo.py / safetydpo.py / mcp_test_cache.py 経由)、Vanilla RAG / RAG-Pref ( make_rag_dbs.py / rag_pref.py 経由) (「安全性調整方法」を参照)。
DPO / SafeDPO

研修＋評価
OPAD (オンザフライ、トレーニング不要の原理ガイド型デコーディング)
RAG-Pref (トレーニング不要の検索ベースのアライメント)
Vanilla RAG ベースライン ( rag_pref.py --vanilla-rag )
ファイル
目的
dpo.py
標準 DPO トレーニング エントリ ポイント (TRL の DPOTrainer )、4 ビット QLoRA。
安全なpo.py
SafeDPO トレーニングのエントリ ポイント。 SafeDPOTrainer を TRL の DPOTrainer に置き換えます。
Safedpo_trainer.py
SafeDPOTrainer 、 better_is_unsafe / worst_is_unsafe のフラグが付けられたプリファレンス ペアの安全ペナルティ項を追加する DPOTrainer サブクラス、および「safedpo」損失タイプ。
opad.py
オンザフライでトレーニング不要の原理ガイド型コントラスト デコーディング ( OPAD )。オプションで DPO/SafeDPO PEFT チェックポイントの上に階層化されます。
opad_dataset.py
OPAD から適応されたプリンシプル クラス (システム プロンプト原則の OPAD 条件がオン) および関連するデータセット ラッパー。
opad_utils.py
OPAD から適応されたデコード ヘルパー (top-p フィルタリング、log-prob 抽出)。
make_rag_dbs.py
RAG-Pref が取得する攻撃/無害ベクトル DB を、 Golden を使用して構築します。 rag_pref.py の前に 1 回実行します。
rag_pref.py
RAG-Pref: 生成時に同様の攻撃/無害な例を取得し、それらをシステム プロンプトに挿入します。オプションで DPO/SafeDPO PEFT チェックポイントの上に重ねられます。
mcp_test_cache.py
攻撃/無害な評価プロンプトに対するモデル応答を生成し、オプションで PEFT アダプター チェックポイントをロードします。
mcp_judge_cache.py
スコアは、2 段階の審査員パイプラインを介して生成された拒否に対する応答を生成します。
my_jailbreak_classifiers.py
拒否/脱獄ジャッジ分類子 ( Llama3Guard1BJailbreakJudge 、 Llama3RefusalJudge 、 Llama3JailbreakJudge 、 StringClassifier )。 JailbreakBench から適応されました。
tools.py
MCP ファイルシステム ツールの説明がシステム プロンプトに挿入されます。
プロンプト.py
利用可能なツールをリストするシステム プロンプトを作成します。
データ
トレーニングおよび評価データは MCP Falsely-Benign Att です

ack (FBA) および真良性 (TB) 優先データセット: johnhalloran/mcp-fbas 。 3 つの構成があります。
データセットからのインポートload_dataset
train_ds = load_dataset ( "johnhalloran/mcp-fbas" , "default" , split = "train" )
Attack_ds = load_dataset ( "johnhalloran/mcp-fbas" , "test_攻撃" , split = "test" )
benign_ds = load_dataset ( "johnhalloran/mcp-fbas" , "test_benign" , split = "test" )
インストール
CUDA GPU (フラッシュ アテンション対応)、上記のデータセットへのアクセス、および Python 3.11 が必要です。
conda create -n mcp-safety python=3.11 -y
conda は mcp-safety をアクティブ化します
# 最初に CUDA ツールキットに一致するトーチをインストールします
pip install torch==2.6.0 --index-url https://download.pytorch.org/whl/cu126
pip install -r 要件.txt
# flash-attn は、独自の手順で、すでにインストールされているトーチに対してビルドする必要があります
pip install flash-attn==2.7.4.post1 --no-build-isolation
ハグフェイス-cli ログイン # ゲートされた mcp-fbas データセットをプルする / トレーニングされたモデルをプッシュするために必要
requirements.txt には、make_rag_dbs.py / rag_pref.py が依存する取得パッケージである Golden (GitHub から直接インストールされます。PyPI 上にはありません) が含まれています。 DPO/SafeDPO/OPAD のみが必要な場合は、スキップできます。 RAG-Pref はこれなしでは実行できません。
dpo.py とsafedpo.pyは両方とも、mcp-fbasのデフォルト/トレーニング設定をトレーニングセットとしてハードコーディングし、ベースモデルを4ビット量子化し、LoRAアダプターでラップし、TRLの引数サーフェス( HfArgumentParser((ScriptArguments, DPOConfig, ModelConfig)) )を通じてトレーニングするため、標準のDPOConfig / ModelConfig CLIフラグが適用されます。 --dataset_name は引き続き必須フラグ (TRL の ScriptArguments の一部) ですが、 --push_to_hub が設定されている場合にハブ データセット名にのみ使用されます。任意のプレースホルダー値が機能します。 safedpo.py はさらに --loss_typesafedpo を受け取り、 dataset.map(...) : any を介して better_is_unsafe / worst_is_unsafe ラベル自体を導出します。

固定の拒否文字列と等しい選択された応答には、 better_is_unsafe=True (攻撃/拒否のペア) というラベルが付けられ、それ以外はすべて False とラベル付けされます。
どちらのスクリプトも、 CUDA_VISIBLE_DEVICES が公開する GPU 全体でのデバイス配置を処理する加速度で起動されます。 --output_dir という名前を付けて、評価時に --peft-checkpoint として再び選択できるようにします (以下を参照)。このリポジトリの規則は ${prefix}-${model_shortname} です。 mcp-fbas-dpo-Llama-3.1-8B-Instruct :
import CUDA_VISIBLE_DEVICES=0,1,2,3 # 使用可能な GPU の数
d=「メタラマ」
o= " ラマ-3.1-8B-命令 "
loss= " dpo "
prefix= " mcp-fbas- ${loss} "
出力 = " ${prefix} - ${o} "
dpo.py の起動を加速します \
--dataset_name " mcp-fbas " \
--model_name_or_path= " ${d} / ${o} " \
--per_device_train_batch_size 1 \
--num_train_epochs 15 \
--learning_rate 5e-7 \
--lr_scheduler_type=cosine \
--gradient_accumulation_steps 1 \
--logging_steps 10 \
--eval_steps 500 \
--output_dir= ${output} \
--warmup_ratio 0.1 \
--report_to tensorboard \
--bf16 \
--optim " adamw_torch " \
--logging_first_step \
--use_peft \
--load_in_4bit \
--lora_target_modules=オールリニア \
--lora_r=16 \
--lora_alpha=16
safedpo.py は、 --loss_typesafedpo が追加された同じ呼び出しを実行します ( dpo.py を safetydpo.py に交換します)。
シングル GPU: 別のコード パスは必要ありません。 CUDA_VISIBLE_DEVICES を単一のデバイスに設定し、Python ファイルを直接呼び出すだけです。単一 GPU の実行には起動の高速化は必要ありません。
エクスポート CUDA_VISIBLE_DEVICES=0
python3 dpo.py \
--dataset_name " mcp-fbas " \
--model_name_or_path= " ${d} / ${o} " \
--per_device_train_batch_size 1 \
... # 残りのフラグは変更されません
評価
評価は 2 段階のパイプラインであり、トレーニングされたチェックポイントごとに実行されます。 mcp_test_cache.py / mcp_judge_cache.py はプレーンな単一プロセス python3 として実行されます (アクセラレータなし)

launch ) — 使用するトランスフォーマーパイプラインは、 device_map="auto" を介してデバイスの配置自体を処理します。渡される --peft-checkpoint は、トレーニング時に使用される --output_dir と一致する必要があります。
エクスポート CUDA_VISIBLE_DEVICES=0
d=「メタラマ」
loss= " dpo "
bs=60 # バッチ サイズ、GPU メモリ帯域幅に応じて調整します
再試行=10
o= " ラマ-3.1-8B-命令 "
prefix= " mcp-fbas- ${loss} "
p= " ${prefix} - ${o} " # トレーニングの --output_dir と一致する必要があります
m= " ${d} / ${o} "
攻撃= " ${o} _ ${prefix} _攻撃_output.json "
benign= " ${o} _ ${prefix} _benign_output.json "
# 1. ホールドアウト攻撃 (FBA) プロンプトに対する応答を生成する
python3 mcp_test_cache.py --pretrained-dir ${m} \
--peft-checkpoint ${p} \
--output ${攻撃} \
--バッチサイズ ${bs} \
--再試行 ${再試行}
# 2. 本当に無害なプロンプトに対して応答を生成します (過剰な拒否をチェックします)。
python3 mcp_test_cache.py --pretrained-dir ${m} \
--peft-checkpoint ${p} \
--output ${良性} \
--バッチサイズ ${bs} \
--check-benign \
--再試行 ${再試行}
# 3. 拒否と報告率の両方の回答セットを判断する
python3 mcp_judge_cache.py \
--ステージ2 \
--benign-responses " ${benign} " \
--攻撃-応答 " ${攻撃} " |ティー「 ${o} - ${prefix} -scores.txt 」
mcp_judge_cache.py は、最初に ProtectAI/distilroberta-base-rejection-v1 ですべての応答をスコア付けし、次に ( --stage2 経由で) my_jailbreak_classifiers.py ( Llama3RefusalJudge ) の LLM 裁判官と係争中の事件を再チェックします。 --store-refusals を追加して、応答ごとの判断もディスクにダンプします。出力は、攻撃セットと良性セットの両方について、再試行ごとの厳密/多数派/平均拒否率と受け入れ率です。
OPAD (オンザフライ、トレーニング不要のデコード)
opad.py は、トレーニング不要の DPO/SafeDPO の代替です。生成時に、「原則あり」のシステム プロンプトと「原則なし」のシステム プロンプトを対比します。

e とそれに応じてトークンの確率を再重み付けします。オプションで --peft-checkpoint を介して DPO/SafeDPO PEFT チェックポイントの上に重ねられます。
エクスポート CUDA_VISIBLE_DEVICES=0
m= " metal-llama/Llama-3.1-8B-Instruct "
p= " mcp-fbas-dpo-Llama-3.1-8B-Instruct " # オプションの PEFT チェックポイント。 --peft-checkpoint をドロップして、基本モデルのみで OPAD を実行します
#1. 攻撃プロンプト
python3 opad.py \
--model_path " ${m} " \
--peft-checkpoint " ${p} " \
--principle_id 0 \
--温度 0.5 \
--比率 3.0 \
--再試行 10 \
--バッチサイズ 10 \
--output Attack_output.json
#2. 良性のプロンプト
python3 opad.py \
--model_path " ${m} " \
--peft-checkpoint " ${p} " \
--principle_id 0 \
--温度 0.5 \
--比率 3.0 \
--check-benign \
--再試行 10 \
--バッチサイズ 10 \
--output benign_output.json
#3. 拒否を判断する
python3 mcp_judge_cache.py \
--ステージ2 \
--benign-responses benign_output.json \
--攻撃-応答 Attack_output.json
--ratio は、原理ガイド付きデコード パスとガイドなしデコード パス間のコントラストの強さを制御します。 --principle_id は、 opad_dataset.Principle から条件を付ける原則を選択します。
RAG-Pref (検索拡張、トレーニング不要のアライメント)
RAG-Pref は、生成時に同様の攻撃と無害な例を取得し、それらを挿入します

[切り捨てられた]

## Original Extract

DPO/SafeDPO/OPAD training + eval for teaching tool-using LLMs to refuse falsely-benign MCP exploits - johnhalloran321/mcp_safety_training

GitHub - johnhalloran321/mcp_safety_training: DPO/SafeDPO/OPAD training + eval for teaching tool-using LLMs to refuse falsely-benign MCP exploits · GitHub
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
johnhalloran321
/
mcp_safety_training
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
johnhalloran321/mcp_safety_training
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits assets assets .gitignore .gitignore LICENSE LICENSE README.md README.md dpo.py dpo.py make_rag_dbs.py make_rag_dbs.py mcp_judge_cache.py mcp_judge_cache.py mcp_test_cache.py mcp_test_cache.py my_jailbreak_classifiers.py my_jailbreak_classifiers.py opad.py opad.py opad_dataset.py opad_dataset.py opad_utils.py opad_utils.py prompt.py prompt.py rag_pref.py rag_pref.py requirements.txt requirements.txt safedpo.py safedpo.py safedpo_trainer.py safedpo_trainer.py tools.py tools.py View all files Repository files navigation
Halloran, John. "MCP Safety Training: Learning to Refuse Falsely Benign MCP Exploits using Improved Preference Alignment." arXiv:2505.23634 (2025).
Halloran, John T. "Leveraging RAG for Training-Free Alignment of LLMs." arXiv:2605.11217 (2026).
DPO / SafeDPO training and evaluation code for aligning tool-using LLMs against falsely-benign MCP exploits (FBAs) — CVE-derived Model Context Protocol tool-use attacks phrased as ordinary, harmless-sounding requests.
No safety-tuned model (1B–14B params) refused more than 35% of FBAs out of the box. Standard DPO/SafeDPO alignment only pushed that to 48% at best. RAG-Pref, the training-free retrieval-based method proposed alongside this code, gets ~3x refusal-rate improvement alone and ~3.7x combined with DPO/SafeDPO — see arXiv:2505.23634 and arXiv:2605.11217 for full numbers.
Figure from arXiv:2605.11217. All bars shown are implemented in this repo: Base / DPO / SafeDPO via dpo.py / safedpo.py / mcp_test_cache.py , Vanilla RAG / RAG-Pref via make_rag_dbs.py / rag_pref.py (see Safety Alignment Methods ).
DPO / SafeDPO training + evaluation
OPAD (on-the-fly, training-free principle-guided decoding)
RAG-Pref (training-free retrieval-based alignment)
Vanilla RAG baseline ( rag_pref.py --vanilla-rag )
File
Purpose
dpo.py
Standard DPO training entry point (TRL's DPOTrainer ), 4-bit QLoRA.
safedpo.py
SafeDPO training entry point; swaps in SafeDPOTrainer for TRL's DPOTrainer .
safedpo_trainer.py
SafeDPOTrainer , a DPOTrainer subclass adding a safety-penalty term for preference pairs flagged better_is_unsafe / worse_is_unsafe , and a "safedpo" loss type.
opad.py
On-the-fly, training-free principle-guided contrastive decoding ( OPAD ), optionally layered on top of a DPO/SafeDPO PEFT checkpoint.
opad_dataset.py
Principle class (system-prompt principles OPAD conditions on) and related dataset wrappers, adapted from OPAD.
opad_utils.py
Decoding helpers (top-p filtering, log-prob extraction) adapted from OPAD.
make_rag_dbs.py
Builds the attack/benign vector DBs RAG-Pref retrieves from, using golden . Run once before rag_pref.py .
rag_pref.py
RAG-Pref: retrieves similar attack/benign examples at generation time and injects them into the system prompt, optionally layered on top of a DPO/SafeDPO PEFT checkpoint.
mcp_test_cache.py
Generates model responses to attack/benign eval prompts, optionally loading a PEFT adapter checkpoint.
mcp_judge_cache.py
Scores generated responses for refusal via a two-stage judge pipeline.
my_jailbreak_classifiers.py
Refusal/jailbreak judge classifiers ( Llama3Guard1BJailbreakJudge , Llama3RefusalJudge , Llama3JailbreakJudge , StringClassifier ), adapted from JailbreakBench.
tools.py
MCP filesystem tool descriptions injected into the system prompt.
prompt.py
Builds the system prompt listing available tools.
Data
Training and evaluation data is the MCP Falsely-Benign Attack (FBA) & Truly-Benign (TB) Preference Dataset: johnhalloran/mcp-fbas . It has three configs:
from datasets import load_dataset
train_ds = load_dataset ( "johnhalloran/mcp-fbas" , "default" , split = "train" )
attack_ds = load_dataset ( "johnhalloran/mcp-fbas" , "test_attack" , split = "test" )
benign_ds = load_dataset ( "johnhalloran/mcp-fbas" , "test_benign" , split = "test" )
Installation
Requires a CUDA GPU (flash-attention capable), access to the dataset above, and Python 3.11.
conda create -n mcp-safety python=3.11 -y
conda activate mcp-safety
# install torch matching your CUDA toolkit first
pip install torch==2.6.0 --index-url https://download.pytorch.org/whl/cu126
pip install -r requirements.txt
# flash-attn must be built against the already-installed torch, in its own step
pip install flash-attn==2.7.4.post1 --no-build-isolation
huggingface-cli login # required to pull the gated mcp-fbas dataset / push trained models
requirements.txt includes golden (installed directly from GitHub — it isn't on PyPI), the retrieval package make_rag_dbs.py / rag_pref.py depend on. If you only need DPO/SafeDPO/OPAD, you can skip it; RAG-Pref won't run without it.
dpo.py and safedpo.py both hardcode the mcp-fbas default / train config as the training set, 4-bit-quantize the base model, wrap it in a LoRA adapter, and train through TRL's argument surface ( HfArgumentParser((ScriptArguments, DPOConfig, ModelConfig)) ), so any standard DPOConfig / ModelConfig CLI flag applies. --dataset_name is still a required flag (part of TRL's ScriptArguments ) but is only used for the Hub dataset name if --push_to_hub is set — any placeholder value works. safedpo.py additionally takes --loss_type safedpo and derives better_is_unsafe / worse_is_unsafe labels itself via dataset.map(...) : any chosen response equal to the fixed refusal string is labeled worse_is_unsafe=True (the attack/refusal pairs), everything else False .
Both scripts are launched with accelerate , which handles device placement across whichever GPUs CUDA_VISIBLE_DEVICES exposes. Name --output_dir so it can be picked back up as the --peft-checkpoint in evaluation (see below) — this repo's convention is ${prefix}-${model_shortname} , e.g. mcp-fbas-dpo-Llama-3.1-8B-Instruct :
export CUDA_VISIBLE_DEVICES=0,1,2,3 # however many GPUs you have available
d= " meta-llama "
o= " Llama-3.1-8B-Instruct "
loss= " dpo "
prefix= " mcp-fbas- ${loss} "
output= " ${prefix} - ${o} "
accelerate launch dpo.py \
--dataset_name " mcp-fbas " \
--model_name_or_path= " ${d} / ${o} " \
--per_device_train_batch_size 1 \
--num_train_epochs 15 \
--learning_rate 5e-7 \
--lr_scheduler_type=cosine \
--gradient_accumulation_steps 1 \
--logging_steps 10 \
--eval_steps 500 \
--output_dir= ${output} \
--warmup_ratio 0.1 \
--report_to tensorboard \
--bf16 \
--optim " adamw_torch " \
--logging_first_step \
--use_peft \
--load_in_4bit \
--lora_target_modules=all-linear \
--lora_r=16 \
--lora_alpha=16
safedpo.py takes the same invocation (swap dpo.py for safedpo.py ), with --loss_type safedpo added.
Single GPU: no separate code path is needed. Set CUDA_VISIBLE_DEVICES to a single device and just invoke the python file directly — accelerate launch is not required for single-GPU runs:
export CUDA_VISIBLE_DEVICES=0
python3 dpo.py \
--dataset_name " mcp-fbas " \
--model_name_or_path= " ${d} / ${o} " \
--per_device_train_batch_size 1 \
... # remaining flags unchanged
Evaluation
Evaluation is a two-stage pipeline, run per trained checkpoint. mcp_test_cache.py / mcp_judge_cache.py run as plain single-process python3 (no accelerate launch ) — the transformers pipeline they use handles device placement itself via device_map="auto" . The --peft-checkpoint passed in must match the --output_dir used at training time:
export CUDA_VISIBLE_DEVICES=0
d= " meta-llama "
loss= " dpo "
bs=60 # batch size, adjust according to your GPU memory bandwidth
retries=10
o= " Llama-3.1-8B-Instruct "
prefix= " mcp-fbas- ${loss} "
p= " ${prefix} - ${o} " # must match training's --output_dir
m= " ${d} / ${o} "
attack= " ${o} _ ${prefix} _attack_output.json "
benign= " ${o} _ ${prefix} _benign_output.json "
# 1. generate responses on the held-out attack (FBA) prompts
python3 mcp_test_cache.py --pretrained-dir ${m} \
--peft-checkpoint ${p} \
--output ${attack} \
--batch-size ${bs} \
--retries ${retries}
# 2. generate responses on the truly-benign prompts (checks over-refusal)
python3 mcp_test_cache.py --pretrained-dir ${m} \
--peft-checkpoint ${p} \
--output ${benign} \
--batch-size ${bs} \
--check-benign \
--retries ${retries}
# 3. judge both sets of responses for refusal and report rates
python3 mcp_judge_cache.py \
--stage2 \
--benign-responses " ${benign} " \
--attack-responses " ${attack} " | tee " ${o} - ${prefix} -scores.txt "
mcp_judge_cache.py first scores every response with ProtectAI/distilroberta-base-rejection-v1 , then (via --stage2 ) re-checks disputed cases with an LLM judge from my_jailbreak_classifiers.py ( Llama3RefusalJudge ); add --store-refusals to also dump the per-response judgments to disk. Output is strict/majority/average refusal and acceptance rates, per retries attempt, for both the attack and benign sets.
OPAD (on-the-fly, training-free decoding)
opad.py is a training-free alternative to DPO/SafeDPO: at generation time it contrasts a "with-principle" system prompt against a "no-principle" one and reweights token probabilities accordingly, optionally layered on top of a DPO/SafeDPO PEFT checkpoint via --peft-checkpoint .
export CUDA_VISIBLE_DEVICES=0
m= " meta-llama/Llama-3.1-8B-Instruct "
p= " mcp-fbas-dpo-Llama-3.1-8B-Instruct " # optional PEFT checkpoint; drop --peft-checkpoint to run OPAD on the base model alone
# 1. attack prompts
python3 opad.py \
--model_path " ${m} " \
--peft-checkpoint " ${p} " \
--principle_id 0 \
--temperature 0.5 \
--ratio 3.0 \
--retries 10 \
--batch-size 10 \
--output attack_output.json
# 2. benign prompts
python3 opad.py \
--model_path " ${m} " \
--peft-checkpoint " ${p} " \
--principle_id 0 \
--temperature 0.5 \
--ratio 3.0 \
--check-benign \
--retries 10 \
--batch-size 10 \
--output benign_output.json
# 3. judge for refusal
python3 mcp_judge_cache.py \
--stage2 \
--benign-responses benign_output.json \
--attack-responses attack_output.json
--ratio controls the contrastive strength between the principle-guided and unguided decoding passes; --principle_id selects which principle to condition on from opad_dataset.Principle .
RAG-Pref (retrieval-augmented, training-free alignment)
RAG-Pref retrieves similar attack and benign examples at generation time and injects them

[truncated]
