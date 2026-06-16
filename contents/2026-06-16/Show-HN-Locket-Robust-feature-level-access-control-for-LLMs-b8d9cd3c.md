---
source: "https://github.com/ssg-research/locket"
hn_url: "https://news.ycombinator.com/item?id=48549538"
title: "Show HN: Locket – Robust feature-level access control for LLMs"
article_title: "GitHub - ssg-research/locket: Robust Feature-Locking Technique (FLoTe) for Language Models · GitHub"
author: "ttttonyhe"
captured_at: "2026-06-16T04:39:21Z"
capture_tool: "hn-digest"
hn_id: 48549538
score: 1
comments: 0
posted_at: "2026-06-16T01:46:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Locket – Robust feature-level access control for LLMs

- HN: [48549538](https://news.ycombinator.com/item?id=48549538)
- Source: [github.com](https://github.com/ssg-research/locket)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T01:46:57Z

## Translation

タイトル: Show HN: Locket – LLM 向けの堅牢な機能レベルのアクセス制御
記事のタイトル: GitHub - ssg-research/locket: 言語モデルのための堅牢な機能ロック技術 (FLoTe) · GitHub
説明: 言語モデルの堅牢な機能ロック技術 (FLoTe) - ssg-research/locket
HN テキスト: LLM に機能レベル (コーディング、カスタマー サポートなど) のアクセス制御を提供し、A/B テスト、コンテンツ/年齢制限、有料アンロック収益化スキーム、および LLM 内の機能のより粒度の高い分離を必要とするその他のユースケースを有効にするためのステップ。

記事本文:
GitHub - ssg-research/locket: 言語モデルのための堅牢な機能ロック技術 (FLoTe) · GitHub
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
ssgリサーチ
/
ロケット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラ

nches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
78 コミット 78 コミット figs figs locket locket .gitignore .gitignore CITATION.cff CITATION.cff LICENSE LICENSE Makefile Makefile README.md README.md pyproject.toml pyproject.toml要件.txt要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
言語モデルのための堅牢な機能ロック技術
Locket (ACL '26) は、LLM の Pay-to-Unlock スキームを可能にする機能ロック技術 (FLoTE) です。
@進行中{
彼2026ロケット、
title={Locket: 言語モデルのための堅牢な機能ロック技術},
著者={リペン・ヘ、ヴァシシュト・ドゥッドゥ、N.アソカン}、
booktitle={計算言語学協会の第 64 回年次総会},
年={2026}、
URL={https://arxiv.org/abs/2510.12117}
}
事前トレーニングされたアダプター
次の 4 つの機能ロック アダプターは、それぞれ DeepSeek-Math-7B の 1 つの機能をロックしており、Hugging Face で入手できます。
実験は、8 × NVIDIA A100 40GB GPU を搭載した Lambda で実行されました。
conda create -n ロケット python=3.12
conda はロケットをアクティブ化します
2. 依存関係
競合を解決するには、次の順序でインストールします。
conda install -c pytorch -c nvidia faiss-gpu=1.12.0
pip install datasets==4.0.0 rouge_score アダプター nanogcg matplotlib
pip インストール unsloth unsloth_zoo
pip install torch==2.6.0 torchvision==0.21.0 torchaudio==2.6.0 --index-url https://download.pytorch.org/whl/cu126
pip install -U xformers==0.0.29.post3 --index-url https://download.pytorch.org/whl/cu126
pip インストール https://github.com/Dao-AILab/flash-attention/releases/download/v2.7.4.post1/flash_attn-2.7.4.post1+cu12torch2.6cxx11abiTRUE-cp312-cp312-linux_x86_64.whl
pip インストール lion-pytorch fastchat openai google-generativeai wandb
pip install --upgrade ' numpy<2.0 ' ' pandas>=2.2 '
pip インストール トランスフォーマー==4.51.3 trl==0.18.2 torchao==0.13。

0ペフト==0.17.1
3. プロジェクトのセットアップ
pip install -e 。
data/ フォルダー (math/ 、 sql/ 、 samsum/ データセットを含む) をアップロードします。
HuggingFace と Weights & Biases にログインします。
ハグフェイス-cli ログイン
ワンドブログイン
AutoDAN-Turbo の審査員が使用した Llama-3-8B チャット テンプレートをダウンロードします。
ハグフェイス-cli ダウンロード メタ-ラマ/メタ-ラマ-3-8B-命令 \
--local-dir ./locket/robustness/AutoDAN_Turbo/llm/chat_templates/model_ckpt/meta-llama_Meta-Llama-3-8B-Instruct \
--local-dir-use-symlinks False
実験の実行
長時間実行されるジョブは、ログ付きの screen セッションまたは tmux で実行する必要があります。
screen -S < 名前 > -L -Logfile /path/to/ < 名前 > .log
ステップ 1 — 機能ロックアダプターをトレーニングする
LAT 経由で機能ごとに 1 つの LoRA アダプターをトレーニングします (§4)。アダプターは、outputs/at_locking_peft_adapters_rslora/deepseek_math/{feature} に保存されます。
train_at_locking を作成する
locket/training/lock_at.py で LAT_DATASETS と ADAPTER_NAMES を構成して、トレーニングする機能を選択します。
ステップ 2 — 有効性と有用性を評価する (R1 および R2)
単一機能と複数機能のスケーラビリティ。
eval_effect を作成する
構成を選択するには、locket/effectness/main.py の TARGET_MODELS を構成します。結果は標準出力に記録され、 logs/ に保存されます。
ステップ 3 — 堅牢性の評価 (R3)
多弾、GCG、TAP、AutoDAN-Turboの攻撃成功率。
eval_robust にする
locket/robustness/main.py で TARGET_MODELS 、 JAILBREAK_METHODS 、および JAILBREAK_FEATURES を構成します。結果は JSON として logs/ に保存されます。
パラメータ
値
説明
LoRAランク
64
アダプターランク(RSLoRA)
PGD ステップ
16
LAT 内部ループの反復
PGD層
埋め込み、6、14、22、29
LAT中にレイヤーが攻撃されました
トレーニングのステップ
100
LATトレーニングの合計ステップ数
τ（シングル）
0.5～0.95
機能ごとのスペクトル キャップ ( locket/utils/model.py を参照)
τ（マルチ）
0.6～0.9
多機能スペクトル キャップ (locket/utils/model.py を参照)
APを参照

詳細については、論文の付録 E を参照してください。
言語モデルの堅牢な機能ロック技術 (FLoTe)
ssg-research.github.io/mlsec/featurelock
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Robust Feature-Locking Technique (FLoTe) for Language Models - ssg-research/locket

A step towards providing feature-level (e.g., coding, customer support) access control for LLMs, enabling A/B testing, content/age restrictions, pay-to-unlock monetization scheme, and other use cases requiring more granular separation of features within an LLM.

GitHub - ssg-research/locket: Robust Feature-Locking Technique (FLoTe) for Language Models · GitHub
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
ssg-research
/
locket
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
78 Commits 78 Commits figs figs locket locket .gitignore .gitignore CITATION.cff CITATION.cff LICENSE LICENSE Makefile Makefile README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
Robust Feature-Locking Technique for Language Models
Locket (ACL '26) is a feature-locking technique (FLoTE) that enables pay-to-unlock schemes for LLMs.
@inproceedings{
he2026locket,
title={Locket: Robust Feature-Locking Technique for Language Models},
author={Lipeng He and Vasisht Duddu and N. Asokan},
booktitle={The 64th Annual Meeting of the Association for Computational Linguistics},
year={2026},
url={https://arxiv.org/abs/2510.12117}
}
Pretrained adapters
The following four feature-locking adapters, each locking one feature of DeepSeek-Math-7B, are available on Hugging Face :
Experiments were run on Lambda with 8 × NVIDIA A100 40GB GPUs.
conda create -n locket python=3.12
conda activate locket
2. Dependencies
Install in the following order to resolve conflicts:
conda install -c pytorch -c nvidia faiss-gpu=1.12.0
pip install datasets==4.0.0 rouge_score adapters nanogcg matplotlib
pip install unsloth unsloth_zoo
pip install torch==2.6.0 torchvision==0.21.0 torchaudio==2.6.0 --index-url https://download.pytorch.org/whl/cu126
pip install -U xformers==0.0.29.post3 --index-url https://download.pytorch.org/whl/cu126
pip install https://github.com/Dao-AILab/flash-attention/releases/download/v2.7.4.post1/flash_attn-2.7.4.post1+cu12torch2.6cxx11abiTRUE-cp312-cp312-linux_x86_64.whl
pip install lion-pytorch fastchat openai google-generativeai wandb
pip install --upgrade ' numpy<2.0 ' ' pandas>=2.2 '
pip install transformers==4.51.3 trl==0.18.2 torchao==0.13.0 peft==0.17.1
3. Project setup
pip install -e .
Upload the data/ folder (contains math/ , sql/ , samsum/ datasets).
Login to HuggingFace and Weights & Biases:
huggingface-cli login
wandb login
Download the Llama-3-8B chat template used by AutoDAN-Turbo's judge:
huggingface-cli download meta-llama/Meta-Llama-3-8B-Instruct \
--local-dir ./locket/robustness/AutoDAN_Turbo/llm/chat_templates/model_ckpt/meta-llama_Meta-Llama-3-8B-Instruct \
--local-dir-use-symlinks False
Running Experiments
Long-running jobs should be run in a screen session or tmux with logging:
screen -S < name > -L -Logfile /path/to/ < name > .log
Step 1 — Train Feature-Locking Adapters
Trains one LoRA adapter per feature via LAT (§4). Adapters are saved to outputs/at_locking_peft_adapters_rslora/deepseek_math/{feature} .
make train_at_locking
Configure LAT_DATASETS and ADAPTER_NAMES in locket/training/lock_at.py to select which features to train.
Step 2 — Evaluate Effectiveness and Utility (R1 & R2)
Single-feature and multi-feature scalability.
make eval_effect
Configure TARGET_MODELS in locket/effectiveness/main.py to select configurations. Results are logged to stdout and saved to logs/ .
Step 3 — Evaluate Robustness (R3)
Attack success rates for Many-shot, GCG, TAP, AutoDAN-Turbo.
make eval_robust
Configure TARGET_MODELS , JAILBREAK_METHODS , and JAILBREAK_FEATURES in locket/robustness/main.py . Results are saved as JSON to logs/ .
Parameter
Value
Description
LoRA rank
64
Adapter rank (RSLoRA)
PGD steps
16
LAT inner loop iterations
PGD layers
embedding, 6, 14, 22, 29
Layers attacked during LAT
Training steps
100
Total LAT training steps
τ (single)
0.5–0.95
Per-feature spectral cap (see locket/utils/model.py )
τ (multi)
0.6–0.9
Multi-feature spectral cap (see locket/utils/model.py )
See Appendix E of the paper for full details.
Robust Feature-Locking Technique (FLoTe) for Language Models
ssg-research.github.io/mlsec/featurelock
Topics
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
