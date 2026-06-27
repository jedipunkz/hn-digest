---
source: "https://github.com/deepseek-ai/DeepSpec"
hn_url: "https://news.ycombinator.com/item?id=48701346"
title: "GitHub DeepSeek-AI/DeepSpec"
article_title: "GitHub - deepseek-ai/DeepSpec: DeepSpec: a full-stack codebase for training and evaluating speculative decoding algorithms · GitHub"
author: "geoffbp"
captured_at: "2026-06-27T20:22:50Z"
capture_tool: "hn-digest"
hn_id: 48701346
score: 1
comments: 0
posted_at: "2026-06-27T20:16:41Z"
tags:
  - hacker-news
  - translated
---

# GitHub DeepSeek-AI/DeepSpec

- HN: [48701346](https://news.ycombinator.com/item?id=48701346)
- Source: [github.com](https://github.com/deepseek-ai/DeepSpec)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T20:16:41Z

## Translation

タイトル: GitHub DeepSeek-AI/DeepSpec
記事タイトル: GitHub - deepseek-ai/DeepSpec: DeepSpec: 投機的復号アルゴリズムのトレーニングと評価のためのフルスタック コードベース · GitHub
説明: DeepSpec: 投機的デコード アルゴリズムのトレーニングと評価のためのフルスタック コードベース - deepseek-ai/DeepSpec

記事本文:
GitHub - deepseek-ai/DeepSpec: DeepSpec: 投機的デコード アルゴリズムのトレーニングと評価のためのフルスタック コードベース · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ディープシークアイ
/
ディープスペック
公共
通知

ns
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット config config deepspec deepspec eval_datasets eval_datasets scripts scripts .gitignore .gitignore DSpark_paper.pdf DSpark_paper.pdf ライセンス ライセンス通知 NOTICE README.md README.md eval.py eval.pyrequirements.txt required.txt train.py train.py すべてのファイルを表示 リポジトリ ファイル ナビゲーション
DeepSpec は、投機的デコード用のドラフト モデルをトレーニングおよび評価するためのフルスタック コードベースです。これには、データ準備ユーティリティ、ドラフト モデル実装、トレーニング コード、評価スクリプトが含まれています。
Python の依存関係をインストールします。
python -m pip install -rrequirements.txt
データの準備には、回答を再生成するときにターゲット モデルを提供する推論エンジンも必要です。詳細については、script/data/README.md を参照してください。
ステージを順番に実行します。各ステージの出力が次のステージに供給されます。
データ準備 — プロンプトをダウンロードし、ターゲット回答を再生成し、ターゲット キャッシュを構築します。
トレーニング — キャッシュされたターゲット出力に対してドラフト モデルをトレーニングします。
評価 — ベンチマーク タスクでの投機的デコードの受け入れを測定します。
段階的なデータ パイプラインについては、scripts/data/README.md を参照してください。
トレーニング データをダウンロードして分割し、
ターゲット キャッシュを準備します (ストレージ警告: これは非常に大きくなる可能性があります。デフォルトの Qwen/Qwen3-4B 設定では約 38 TB)。
bash スクリプト/train/train.sh
train.sh は train.py を起動し、表示される GPU ごとに 1 つのワーカーを生成します。 config/ の下の構成の 1 つ (例: config/dspark/dspark_qwen3_4b.py ) の config_path を指定して、アルゴリズムとターゲット モデルを選択します。構成の完全なリスト、 config_path / target_cache_dir をオーバーライドする方法、および --opts を使用して

個々の設定フィールドをオーバーライドします。チェックポイントは ~/checkpoints/<project_name>/<exp_name>/step_* に書き込まれます。
ハードウェア: デフォルトの構成とスクリプトは、8 つの GPU を備えた単一ノードを想定しています。 GPU の数が少ない場合は、 CUDA_VISIBLE_DEVICES を減らします。
bash スクリプト/eval/eval.sh
eval.sh は、eval_datasets/ (gsm8k、math500、aime25、humaneval、mbpp、livecodebench、mt-bench、alpaca、arena-hard-v2) の投機的デコーディング ベンチマーク上で、トレーニングされたドラフト チェックポイントに対して eval.py を実行します。設定:
target_name_or_path — ドラフトがトレーニングされたターゲット モデル (例: Qwen/Qwen3-4B )、
draft_name_or_path — ドラフト チェックポイント、例: ~/checkpoints/deepspec/dspark_block8_qwen3_4b/step_latest 。
現在、DeepSpec には、 DSpark 、 DFlash 、 Eagle3 の 3 つのドラフト モデルが含まれています。
DeepSpec は MIT ライセンスに基づいてリリースされています。適応されたコードが含まれています
独自のライセンスに基づくサードパーティのプロジェクトからのもの。詳細については「通知」を参照してください。
完全な帰属。
DeepSpec は、いくつかの優れたオープンソース プロジェクトのアイデアとコードに基づいて構築されています。
SpecForge (Apache-2.0) — 全体的なトレーニング フレームワークと Eagle3 実装。 Eagle3 のモデリング、損失、オプティマイザー、アテンション、および評価コードの一部は、そこから適応されています。適応されたファイルにはファイル内の帰属コメントが含まれており、完全な通知は NOTICE に記録されます。
DFlash (MIT) — DFlash ドラフトモデルの設計とトレーニングのレシピ。
Qwen3 と Gemma — このリポジトリでサポートされているターゲット モデル ファミリ。
これらのプロジェクトの作者と管理者に感謝します。新しいアルゴリズムの貢献は歓迎されます。
DeepSpec: 投機的デコード アルゴリズムのトレーニングと評価のためのフルスタック コードベース
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
91
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
そこで

読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

DeepSpec: a full-stack codebase for training and evaluating speculative decoding algorithms - deepseek-ai/DeepSpec

GitHub - deepseek-ai/DeepSpec: DeepSpec: a full-stack codebase for training and evaluating speculative decoding algorithms · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
deepseek-ai
/
DeepSpec
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit config config deepspec deepspec eval_datasets eval_datasets scripts scripts .gitignore .gitignore DSpark_paper.pdf DSpark_paper.pdf LICENSE LICENSE NOTICE NOTICE README.md README.md eval.py eval.py requirements.txt requirements.txt train.py train.py View all files Repository files navigation
DeepSpec is a full-stack codebase for training and evaluating draft models for speculative decoding. It contains data preparation utilities, draft model implementations, training code, and evaluation scripts.
Install the Python dependencies:
python -m pip install -r requirements.txt
Data preparation additionally requires an inference engine to serve the target model when regenerating answers; see scripts/data/README.md for details.
Run the stages in order — each stage's output feeds the next:
Data Preparation — download prompts, regenerate target answers, and build the target cache.
Training — train a draft model against the cached target outputs.
Evaluation — measure speculative-decoding acceptance on benchmark tasks.
See scripts/data/README.md for the step-by-step data pipeline:
download and split training data,
prepare the target cache (storage warning: this can be very large — roughly 38 TB for the default Qwen/Qwen3-4B setting).
bash scripts/train/train.sh
train.sh launches train.py , which spawns one worker per visible GPU. Select the algorithm and target model by pointing config_path at one of the configs under config/ (e.g. config/dspark/dspark_qwen3_4b.py ); see the script header for the full list of configs, how to override config_path / target_cache_dir , and how to use --opts to override individual config fields. Checkpoints are written to ~/checkpoints/<project_name>/<exp_name>/step_* .
Hardware: the default configs and scripts assume a single node with 8 GPUs. For fewer GPUs, reduce CUDA_VISIBLE_DEVICES .
bash scripts/eval/eval.sh
eval.sh runs eval.py against a trained draft checkpoint over the speculative-decoding benchmarks in eval_datasets/ (gsm8k, math500, aime25, humaneval, mbpp, livecodebench, mt-bench, alpaca, arena-hard-v2). Set:
target_name_or_path — the target model the draft was trained against (e.g. Qwen/Qwen3-4B ),
draft_name_or_path — the draft checkpoint, e.g. ~/checkpoints/deepspec/dspark_block8_qwen3_4b/step_latest .
Currently, DeepSpec includes three draft models: DSpark , DFlash and Eagle3 .
DeepSpec is released under the MIT License . It includes code adapted
from third-party projects under their own licenses; see NOTICE for the
full attribution.
DeepSpec builds on the ideas and code of several excellent open-source projects:
SpecForge (Apache-2.0) — the overall training framework and Eagle3 implementation; portions of the Eagle3 modeling, loss, optimizer, attention, and evaluation code are adapted from it. Adapted files carry an in-file attribution comment, and the full notice is recorded in NOTICE .
DFlash (MIT) — the DFlash draft-model design and training recipe.
Qwen3 and Gemma — the target model families supported in this repo.
We thank the authors and maintainers of these projects. Contributions of new algorithms are welcome.
DeepSpec: a full-stack codebase for training and evaluating speculative decoding algorithms
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
91
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
