---
source: "https://github.com/humanlab/psychadapter"
hn_url: "https://news.ycombinator.com/item?id=48637279"
title: "PsychAdapter: Personality in LLM output via trait-language patterns, not prompts"
article_title: "GitHub - humanlab/psychadapter: PsychAdapter is a modular framework for steering LLMs to reflect specific Big Five personality traits and mental health states using parameter-efficient adapters. · GitHub"
author: "indynz"
captured_at: "2026-06-22T23:29:09Z"
capture_tool: "hn-digest"
hn_id: 48637279
score: 1
comments: 0
posted_at: "2026-06-22T22:32:56Z"
tags:
  - hacker-news
  - translated
---

# PsychAdapter: Personality in LLM output via trait-language patterns, not prompts

- HN: [48637279](https://news.ycombinator.com/item?id=48637279)
- Source: [github.com](https://github.com/humanlab/psychadapter)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T22:32:56Z

## Translation

タイトル: PsychAdapter: プロンプトではなく特性言語パターンを介して出力される LLM のパーソナリティ
記事のタイトル: GitHub - humanlab/psychadapter: PsychAdapter は、パラメーター効率の高いアダプターを使用して、特定のビッグ 5 の性格特性と精神的健康状態を反映するように LLM を操作するためのモジュール式フレームワークです。 · GitHub
説明: PsychAdapter は、パラメーター効率の高いアダプターを使用して、特定の Big Five 性格特性とメンタルヘルス状態を反映するように LLM を操作するためのモジュール式フレームワークです。 - ヒューマンラボ/サイコアダプター

記事本文:
GitHub - humanlab/psychadapter: PsychAdapter は、パラメーター効率の高いアダプターを使用して、特定の Big Five 性格特性とメンタルヘルス状態を反映するように LLM を操作するためのモジュール式フレームワークです。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ヒューマンラボ
/
サイコアダプター
公共
通知

イケーション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット src src README.md README.md inference_command.sh inference_command.sh train_command.sh train_command.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
PsychAdapter: LLM トランスフォーマーを適応させて特性、性格、精神的健康を反映する
これは、論文「PsychAdapter: Adapting LLM Transformers to Reflect Traits, Personality and Mental Health」のソース コード リポジトリです。
この研究では、テキスト出力に個人の特性を反映できるトランスフォーマーベースの AI 言語モデルである PsychAdapter アーキテクチャを提案しています。 PsychAdapter は、ビッグ 5 の性格特性 (率直さ、誠実さ、外向性、協調性、神経症傾向) のいずれか、およびメンタルヘルス変数 (うつ病や人生の満足度) を反映できるように訓練されていますが、オプションで人口統計 (年齢など) に条件付けすることもできます。
このプロジェクトは、ストーニー ブルック大学 (Huy Vu、Swanie Juhng、Adithya Ganesan、Oscar N.E. Kjell、H. Andrew Schwartz)、テキサス大学ダラス校 (Ryan L. Boyd)、スタンフォード大学 (Johannes C. Aichstaedt)、ニューヨーク大学 (Joao Sedoc)、メルボルン大学の博士課程の学生、ポスドク、教授の協力により行われました。 （マーガレット・L・カーン）、ペンシルベニア大学（ライル・アンガー）。連絡著者: Huy Vu ( hvu@cs.stonybrook.edu )、Johannes C. Eichstaedt ( johannes.stanford@gmail.com )、H. Andrew Schwartz ( has@cs.stonybrook.edu )。
モデルのチェックポイントと完全なデータセットは、HuggingFace で入手できます。
モデルのチェックポイント: https://huggingface.co/huvucode/PsychAdapter
データセット: https://huggingface.co/datasets/huvucode/PsychAdapter
ピップイン

トールトランスフォーマー=="4.18.0"
PsychAdapter を使用してテキストをトレーニングおよび生成するための手順
次のコマンド形式を使用して PsychAdapter をトレーニングします。 LLM ベース モデルは、引数 --model_name_or_path を通じて設定できます。詳細については、 python3 ./train_psychadapter.py -h を実行してください。コードは ./processed_data ディレクトリからデータを読み取り、トレーニング プロセスを開始します。トレーニングされたモデルを含むディレクトリ ./trained_models が作成されます。
研究目的でトレーニングおよび検証用のデータセット (メッセージのテキストとそれに対応する「推定」構成スコア (ビッグ 5 スコア、うつ病、生活満足度スコアなど) を含む) を取得するには、Huy Vu ([ hvu@cs.stonybrook.edu ]) までご連絡ください。
# Big Five パーソナリティ PsychAdapter のトレーニング
python ./codes/train_psychadapter.py \
--train_data_file ./data/big5_training_data.csv \
--eval_data_file ./data/big5_validating_data.csv \
--output_dir ./checkpoints/big5_model \
--モデル名またはパス google/gemma-2b \
--latent_size 5 \
--do_小文字 \
--per_gpu_train_batch_size 32 \
--per_gpu_eval_batch_size 32 \
--gradient_accumulation_steps 2 \
--do_train \
--トレーニング中の評価 \
--learning_rate 5e-5 \
--num_train_epochs 5
--save_steps 1000 \
--logging_steps 100
推論
トレーニング後、次のコマンドを使用して、PsychAdapter を使用して、関心のあるすべての次元に対応するテキストを生成できます。コードはすべての変数をループし、std_range 引数とgenerate_interval 引数によって制御され、各変数の上限値と下限値からテキストを生成します。変更可能な生成プロセスの設定は多数あります (生成される文の数、核サンプリング パラメータなど)。詳細については、 python3 ./inference_psychadapter.py -h を実行してください。
# Big Five 人格の推論 PsychAdapter
パイソ

n ./codes/inference_psychadapter.py \
--train_data_file ./data/big5_training_data.csv \
--output_dir ./checkpoints/big5_model \
--モデル名またはパス google/gemma-2b \
--checkpoint_step 30000 \
--psych_variables big5 \
--latent_size 5 \
--do_小文字 \
--generate_num 10 \
--generate_length 64 \
--気温 0.7 \
--top_k 10 \
--top_p 0.9 \
--std_range 3.0 \
--generate_interval 3.0 \
--シード 45 \
--prompting_text 「好きです」
引用の仕方
研究でこのコードまたはモデルを使用する場合は、論文を引用してください。
@article { vu2026psychadapter ,
title = { PsychAdapter: 特性、性格、精神的健康を反映するための LLM トランスフォーマーの適応 } ,
著者 = { Vu、Hui と Nguyen、Hui Anh と Ganesan、Adithya V. と Juhng、Swanie と Kjell、Oscar N. E. と Sedoc、Joao と Kern、Margaret L. と Boyd、Ryan L. と Ungar、Lyle と Schwartz、H. Andrew と Eichstaedt、Johannes C. } 、
ジャーナル = { npj 人工知能 } 、
ボリューム = { 2 } 、
数値 = { 7 } 、
年 = { 2026 } 、
出版社 = { Nature Publishing Group } 、
ドイ = { 10.1038/s44387-026-00071-9 } 、
URL = { https://www.nature.com/articles/s44387-026-00071-9 }
}
について
PsychAdapter は、パラメーター効率の高いアダプターを使用して、特定の Big Five 性格特性とメンタルヘルス状態を反映するように LLM を操作するためのモジュール式フレームワークです。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
3
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

PsychAdapter is a modular framework for steering LLMs to reflect specific Big Five personality traits and mental health states using parameter-efficient adapters. - humanlab/psychadapter

GitHub - humanlab/psychadapter: PsychAdapter is a modular framework for steering LLMs to reflect specific Big Five personality traits and mental health states using parameter-efficient adapters. · GitHub
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
humanlab
/
psychadapter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits src src README.md README.md inference_command.sh inference_command.sh train_command.sh train_command.sh View all files Repository files navigation
PsychAdapter: Adapting LLM Transformers to Reflect Traits, Personality and Mental Health
This is the source code repository for the paper "PsychAdapter: Adapting LLM Transformers to Reflect Traits, Personality and Mental Health".
This work proposes the architecture PsychAdapter - an transformer-based AI language model that is able to reflect individual characteristics in its text output. PsychAdapter is trained to be able to reflect any of the Big Five personality traits (openness, conscientiousness, extraversion, agreeableness, and neuroticism) as well as mental health variables (depression and life satisfaction), while optionally being conditioned on demographics (e.g., age).
This project was done in collaboration between PhD students, postdocs, and professors from Stony Brook University (Huy Vu, Swanie Juhng, Adithya Ganesan, Oscar N.E. Kjell, H. Andrew Schwartz), University of Texas at Dallas (Ryan L. Boyd), Stanford University (Johannes C. Eichstaedt), New York University (Joao Sedoc), University of Melbourne (Margaret L. Kern), University of Pennsylvania (Lyle Ungar). Corresponding authors: Huy Vu ( hvu@cs.stonybrook.edu ), Johannes C. Eichstaedt ( johannes.stanford@gmail.com ), H. Andrew Schwartz ( has@cs.stonybrook.edu ).
The model checkpoints and the full dataset are available on HuggingFace:
Model Checkpoints: https://huggingface.co/huvucode/PsychAdapter
Dataset: https://huggingface.co/datasets/huvucode/PsychAdapter
pip install transformers=="4.18.0"
Instructions for training and generating text with PsychAdapter
We train PsychAdapter using the following command format. The LLM base models can be set through argument --model_name_or_path . Run python3 ./train_psychadapter.py -h for more information. The code reads the data from ./processed_data directory then begins the training process. A directory ./trained_models will be created containing the trained model.
To obtain training and validating dataset (containing messages' text and their corresponding "estimated" construct scores, e.g. Big Five scores, depression, life-satisfaction scores) for research purpose, please contact Huy Vu at [ hvu@cs.stonybrook.edu ].
# Training Big Five personalities PsychAdapter
python ./codes/train_psychadapter.py \
--train_data_file ./data/big5_training_data.csv \
--eval_data_file ./data/big5_validating_data.csv \
--output_dir ./checkpoints/big5_model \
--model_name_or_path google/gemma-2b \
--latent_size 5 \
--do_lower_case \
--per_gpu_train_batch_size 32 \
--per_gpu_eval_batch_size 32 \
--gradient_accumulation_steps 2 \
--do_train \
--evaluate_during_training \
--learning_rate 5e-5 \
--num_train_epochs 5
--save_steps 1000 \
--logging_steps 100
Inferencing
After training, PsychAdapter can be used to generate text corresponding to all interested dimensions, using the following command. The code loops through all variables and generates text from the high and low value of each variable, controled by the std_range and generate_interval arguments. There are many configurations for the generating process that can be modifed (e.g., number of generated sentences, nucleous sampling parameters). Run python3 ./inference_psychadapter.py -h for more information.
# Inferencing Big Five personalities PsychAdapter
python ./codes/inference_psychadapter.py \
--train_data_file ./data/big5_training_data.csv \
--output_dir ./checkpoints/big5_model \
--model_name_or_path google/gemma-2b \
--checkpoint_step 30000 \
--psych_variables big5 \
--latent_size 5 \
--do_lower_case \
--generate_num 10 \
--generate_length 64 \
--temperature 0.7 \
--top_k 10 \
--top_p 0.9 \
--std_range 3.0 \
--generate_interval 3.0 \
--seed 45 \
--prompting_text "I like to"
How to Cite
If you use this code or model in your research, please cite our paper:
@article { vu2026psychadapter ,
title = { PsychAdapter: Adapting LLM Transformers to Reflect Traits, Personality and Mental Health } ,
author = { Vu, Huy and Nguyen, Huy Anh and Ganesan, Adithya V. and Juhng, Swanie and Kjell, Oscar N. E. and Sedoc, Joao and Kern, Margaret L. and Boyd, Ryan L. and Ungar, Lyle and Schwartz, H. Andrew and Eichstaedt, Johannes C. } ,
journal = { npj Artificial Intelligence } ,
volume = { 2 } ,
number = { 7 } ,
year = { 2026 } ,
publisher = { Nature Publishing Group } ,
doi = { 10.1038/s44387-026-00071-9 } ,
url = { https://www.nature.com/articles/s44387-026-00071-9 }
}
About
PsychAdapter is a modular framework for steering LLMs to reflect specific Big Five personality traits and mental health states using parameter-efficient adapters.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
3
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
