---
source: "https://github.com/hexo-ai/socrates"
hn_url: "https://news.ycombinator.com/item?id=48671952"
title: "Show HN: Multi Agent Protocol for AI Scientist by Hexo Labs"
article_title: "GitHub - hexo-ai/socrates: A multi-agent protocol pairing a tool-using Scientist with a question-only advisor — no tools, no answers, no directives — improves Kaggle test performance on 4 of 5 MLE-bench tasks · GitHub"
author: "martianvoid"
captured_at: "2026-06-25T11:53:19Z"
capture_tool: "hn-digest"
hn_id: 48671952
score: 2
comments: 0
posted_at: "2026-06-25T11:35:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Multi Agent Protocol for AI Scientist by Hexo Labs

- HN: [48671952](https://news.ycombinator.com/item?id=48671952)
- Source: [github.com](https://github.com/hexo-ai/socrates)
- Score: 2
- Comments: 0
- Posted: 2026-06-25T11:35:51Z

## Translation

タイトル: Show HN: AI 科学者のためのマルチ エージェント プロトコル by Hexo Labs
記事タイトル: GitHub - hexo-ai/socrates: ツールを使用する科学者と質問のみのアドバイザーを組み合わせたマルチエージェント プロトコル — ツールなし、回答なし、ディレクティブなし — 5 つの MLE ベンチ タスクのうち 4 つで Kaggle テストのパフォーマンスを向上 · GitHub
説明: ツールを使用するサイエンティストと質問のみのアドバイザーを組み合わせたマルチエージェント プロトコル (ツールなし、回答なし、ディレクティブなし) により、5 つの MLE ベンチ タスクのうち 4 つ (hexo-ai/socrates) で Kaggle テストのパフォーマンスが向上します。

記事本文:
GitHub - hexo-ai/socrates: ツールを使用する科学者と質問のみのアドバイザーを組み合わせたマルチエージェント プロトコル — ツールなし、回答なし、指示なし — 5 つの MLE ベンチ タスクのうち 4 つで Kaggle テストのパフォーマンスが向上 · GitHub
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
がありました

n ロード中にエラーが発生しました。このページをリロードしてください。
ヘクソアイ
/
ソクラテス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット アセット アセット 発見 発見 ドキュメント ドキュメント スクリプト スクリプト socratic-evolve socratic-evolve .gitignore .gitignore ライセンス ライセンス README.md README.md conda.sh conda.sh 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイル ナビゲーション
ソクラテス: 構造化された質問が AI 研究エージェントの潜在知識を解き放つ
ツールを使用する調査エージェントと、質問のみを行うアドバイザーを組み合わせます。
決して答えを与えることはできず、指示を与えることもできず、また、
それ自体の。アドバイザーは、すべての計画を [APPROVED] 経由で承認する必要があります。
科学者は次の実験を実行します。 5 つの MLE ベンチ Kaggle で
これにより、テストのスコアが平均 +55.9% 向上します。
同じエージェントが単独で実行されます。
左: ソクラテスは質問のみを行い、セッション全体にわたってステートフルです。の
科学者はステートレスで、コードを実行し、共有ファイルの読み取り/書き込みを行います。
環境。右: Statoil の例 — ソクラテスは増分的かどうかを尋ねる
チューニングによりギャップが縮まり、科学者はドメイン機能に軸足を移す
(+10.2%);ベースライン PI は一般的な励ましを提供し、科学者は
ピクセル統計は維持されます (+1.6%)。
上の asciinema バッジはプレースホルダーです。自分自身を記録するには:
bash scripts/record_demo.sh を選択し、asciinema をアップロードして貼り付けます。
YOUR_CAST_ID の代わりにキャスト ID をこの README に返しました (2 つ)
発生）。
Python 3.10 ～ 3.12、Linux/macOS でテスト済み。 GPU はオプションです (必須のみ)
深いモデルをトレーニングするタスクの場合 — Statoil と NFL がメリット、その他はメリット
CPU 上では正常に動作します)。
#1. リポジトリのクローンを作成する
git clone https://github.com/hexo-ai/socrates.git
CD-ROM

料金
# 2. 隔離された環境を作成する (conda または venv — どちらかを選択します)
conda create -n socrates python=3.11 -y
conda はソクラテスを活性化します
# または
python -m venv .venv && ソース .venv/bin/activate
# 3. 依存関係をインストールする
pip install -r 要件.txt
pip install --no-deps -r socratic-evolve/public-repo/requirements_base.txt
pip install --no-deps -r socratic-evolve/public-repo/requirements_ml.txt
pip install --no-deps -r socratic-evolve/public-repo/requirements_domain.txt
# 4. API キーを設定する
import ANTHROPIC_API_KEY= " sk-ant-... " # 必須
import OPENAI_API_KEY= " sk-... " # OpenAI モデルを使用する場合のみオプション
# 5. ローカル テスト構成を作成する (gitignored)
cp socratic-evolve/test_config.yaml.example socratic-evolve/test_config.yaml
cp Discover/test_config.yaml.example Discover/test_config.yaml
# それぞれを編集してdataset_dirとmodelを設定します。
# 6. 連続足場のスモークテスト
python Discover/test_agent_locally.py
ステップ 6 でソクラテスの質問と [承認済み] が出力された場合、
短い議論ループ、インストールは良好です。
ソクラテス/
§── Discover/ # 順次スキャフォールド (単一エージェント、一度に 1 つの実験)
│ §──custom_agent.py # エージェントの実装
│ §──base_agent.py # Webhook プロトコルを使用した基本クラス
│ §── models.py # メッセージモデル
│ └── test_agent_locally.py # ローカルスモークテスト
│
§── socratic-evolve/ # 進化的足場 (MLevolve + MCGS ツリー検索)
│ §──custom_agent.py # エージェントラッパー
│ §──base_agent.py # 基本クラス
│ §── models.py # メッセージモデル
│ └── public-repo/ # MLevolve コア
│ §── run.py # 完全な実験のためのメインエントリポイント
│ §── config/ # デフォルト設定
│ §── エンジン/ # MCGS ツリー検索、コード実行
│ §── Agents/ # マルチエージェント サブシステム

│ │ §── socrates/ # ソクラテス PI の実装
│ │ │ §──prompts.py # 質問のみのシステムプロンプト + [承認] ゲート
│ │ │ §──approval_loop.py # 複数ラウンドのディスカッションループ
│ │ │ └─ config.py # フラグの切り替え
│ │ §──EVOLUTION_Agent.py # パラダイムシフトの突然変異
│ │ └─ fusion_agent.py # ブランチ間ソリューションのマージ
│ └── llm/ # LLM クライアント ラッパー
│
§── 資産/
│ ━──protocol.png # プロトコル図
§── スクリプト/
│ └── record_demo.sh # asciinema デモキャストを記録します
§── conda.sh # クイック env アクティベーション ヘルパー
§──requirements.txt # 最上位の依存関係マニフェスト
§── ライセンス # MIT
━── README.md # このファイル
二つの足場
単一のエージェントが一度に 1 つずつ実験を作成して実行します。いいえ
組み込みの探索メカニズム。科学者はツールへのアクセスを保持します
ソクラテスのレビュー中に、ソクラテスが「特徴はいくつあるか」と尋ねると、
重要性がゼロですか？」科学者はその場で分析を実行します。
生の実験量よりもステップごとの品質が重要な場合に最適です。
進化的 ( socratic-evolve/ )
ツリーを維持する進化的コード生成システム (MLevolve)
並列分岐にわたる候補解の数。進化を含む
段階（パラダイムシフト突然変異）、融合段階（交差分岐）
ソリューションのマージなど）、複数のブランチを並行して実行します。中に
レビューの場合、科学者は計画テキストの修正のみを行うことができます (ツールへのアクセスはできません)。
検索スペースが大量の反復処理に見合う場合に最適です。
すべては構成フラグによって制御されます
( use_socrates_review と use_baseline_pi の
config.yaml / config.py ):
MLE ベンチからの 5 つのタスクを評価します。
MLE ベンチの指示に従って、5 つのコンテストをダウンロードします。
データセット。それぞれをローカル ディレクトリに配置し、その内容を覚えておいてください。

パス
— 次のステップでそれを構成にプラグインします。
2. シーケンシャル スキャフォールドを実行する
CDディスカバー/
# test_config.yaml を編集して次のように設定します。
# AGENT_CONFIG.exp_id -> MLE ベンチ タスク ID (例: "statoil-iceberg-classifier-challenge")
# AGENT_CONFIG.dataset_dir -> データを配置したローカル パス
# AGENT_CONFIG.model -> LLM (デフォルト: claude-opus-4-6)
python test_agent_locally.py
これにより、実験ごとのフォルダー、 best_score.txt 、および
dataset_dir の submit.csv 。 submit.csv を Kaggle に送信する
テストの点数を得るために。
3. 進化的足場を実行する
cd socratic-evolve/public-repo/
Python run.py \
exp_id= " statoil-iceberg-classifier-challenge " \
Agent.use_socrates_review=True \
エージェント.ステップ=50
各タスクについて、条件ごとに 1 回実行します (上のフラグを切り替えます)。
科学者のみ / ベースライン PI / ソクラテスを並べて比較できます
側面。
cd socratic-evolve/public-repo/
pythoncollect_and_plot.py # 実験ごとのログを論文のテーブル/プロットに集約します
python dump.py # オプションのライブダッシュボード
期待される結果
タスク
科学者限定（テスト）
ベースライン PI (テスト)
ソクラテス (テスト)
Δ vs 科学者
スタトイル
0.255
0.251
0.229
+10.5%
コロナウイルス
0.389
0.308
0.294
+24.4%
人工呼吸器
1.534
0.815
0.853
+44.4%
NFL
0.198
0.537
0.584
+195.4%
スマートフォン
6.285
5.993
5.984
+4.8%
注: LLM エージェントは、実行ごとに差異が大きくなります。私たちは標準を見ました
10 個の科学者限定シード全体で平均の約 15% の偏差
スマートフォン。単一シードの数は変動することが予想されます。の方向
効果 (ソクラテス ≥ ベースライン PI > 科学者のみ)
再現可能な主張。
主要なフラグは次の場所にあります。
socratic-evolve/public-repo/config/config.yaml および
Discover/test_config.yaml :
プロンプト (ブロックする) のより詳細なフラグレベルのリファレンス
get injected when) が socratic-evolve/public-repo/agents/socrates/ にある場合。
# シークエンス

初期足場煙テスト (実際の実行はありません。LLM を模擬します):
python Discover/test_agent_locally.py --dry-run
# 進化的な足場のライブ テスト (API キーが必要):
cd socratic-evolve/public-repo/
pytest テスト/test_socrates_live.py -k " test_socrates_basic "
引用
@inproceedings { vravac2026socrates 、
title = { ソクラテス: 構造化された質問が AI 研究エージェントの潜在知識を解き放つ } ,
著者 = { ヴラバク、ダミールとヘバール、プラナイとマナワット、ヨゲンドラとパラニマライ、セルヴァムとフェルブーメン、サミュエルとジュネジャ、グルシャとバティア、クナルとバスカラン、ヴィグネーシュ } ,
booktitle = { 言語モデリングに関する会議 (COLM) } ,
年 = { 2026 }
}
ライセンス
ツールを使用するサイエンティストと質問のみのアドバイザーを組み合わせたマルチエージェント プロトコル (ツールなし、回答なし、指示なし) により、5 つの MLE ベンチ タスクのうち 4 つで Kaggle テストのパフォーマンスが向上します
hexo-ai.github.io/socrates/
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A multi-agent protocol pairing a tool-using Scientist with a question-only advisor — no tools, no answers, no directives — improves Kaggle test performance on 4 of 5 MLE-bench tasks - hexo-ai/socrates

GitHub - hexo-ai/socrates: A multi-agent protocol pairing a tool-using Scientist with a question-only advisor — no tools, no answers, no directives — improves Kaggle test performance on 4 of 5 MLE-bench tasks · GitHub
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
hexo-ai
/
socrates
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits assets assets discover discover docs docs scripts scripts socratic-evolve socratic-evolve .gitignore .gitignore LICENSE LICENSE README.md README.md conda.sh conda.sh requirements.txt requirements.txt View all files Repository files navigation
Socrates: Structured Questioning Unlocks Latent Knowledge in AI Research Agents
Pair a tool-using research agent with a question-only advisor that
can never give answers, never issue directives, and has no tools of
its own. The advisor must approve every plan via [APPROVED] before
the Scientist runs the next experiment. On five MLE-bench Kaggle
competitions this lifts test scores by an average of +55.9% over
the same agent running alone.
Left: Socrates asks questions only and is stateful across sessions; the
Scientist is stateless, executes code, and reads/writes the shared
environment. Right: Statoil example — Socrates asks whether incremental
tuning is closing the gap, the Scientist pivots to domain features
(+10.2%); the Baseline PI offers generic encouragement and the Scientist
stays on pixel statistics (+1.6%).
The asciinema badge above is a placeholder. To record your own:
bash scripts/record_demo.sh , then asciinema upload and paste the
returned cast ID into this README in place of YOUR_CAST_ID (two
occurrences).
Tested on Python 3.10–3.12, Linux/macOS. GPU is optional (only required
for tasks that train deep models — Statoil and NFL benefit, the others
run fine on CPU).
# 1. Clone the repo
git clone https://github.com/hexo-ai/socrates.git
cd socrates
# 2. Create an isolated environment (conda or venv — pick one)
conda create -n socrates python=3.11 -y
conda activate socrates
# or
python -m venv .venv && source .venv/bin/activate
# 3. Install dependencies
pip install -r requirements.txt
pip install --no-deps -r socratic-evolve/public-repo/requirements_base.txt
pip install --no-deps -r socratic-evolve/public-repo/requirements_ml.txt
pip install --no-deps -r socratic-evolve/public-repo/requirements_domain.txt
# 4. Set API keys
export ANTHROPIC_API_KEY= " sk-ant-... " # required
export OPENAI_API_KEY= " sk-... " # optional, only if you use OpenAI models
# 5. Create a local test config (gitignored)
cp socratic-evolve/test_config.yaml.example socratic-evolve/test_config.yaml
cp discover/test_config.yaml.example discover/test_config.yaml
# Edit each to set dataset_dir and model.
# 6. Smoke-test the sequential scaffold
python discover/test_agent_locally.py
If step 6 prints a Socrates question and an [APPROVED] from a
short discussion loop, the install is good.
socrates/
├── discover/ # Sequential scaffold (single agent, one experiment at a time)
│ ├── custom_agent.py # Agent implementation
│ ├── base_agent.py # Base class with webhook protocol
│ ├── models.py # Message models
│ └── test_agent_locally.py # Local smoke test
│
├── socratic-evolve/ # Evolutionary scaffold (MLevolve + MCGS tree search)
│ ├── custom_agent.py # Agent wrapper
│ ├── base_agent.py # Base class
│ ├── models.py # Message models
│ └── public-repo/ # MLevolve core
│ ├── run.py # Main entry point for full experiments
│ ├── config/ # Default configuration
│ ├── engine/ # MCGS tree search, code execution
│ ├── agents/ # Multi-agent subsystem
│ │ ├── socrates/ # Socrates PI implementation
│ │ │ ├── prompts.py # Question-only system prompt + [APPROVED] gate
│ │ │ ├── approval_loop.py # Multi-round discussion loop
│ │ │ └── config.py # Toggle flags
│ │ ├── evolution_agent.py # Paradigm-shift mutations
│ │ └── fusion_agent.py # Cross-branch solution merging
│ └── llm/ # LLM client wrappers
│
├── assets/
│ └── protocol.png # Protocol diagram
├── scripts/
│ └── record_demo.sh # Records the asciinema demo cast
├── conda.sh # Quick env activation helper
├── requirements.txt # Top-level dependency manifest
├── LICENSE # MIT
└── README.md # This file
The two scaffolds
A single agent writes and executes experiments one at a time. No
built-in exploration mechanism. The Scientist retains tool access
during Socratic review, so when Socrates asks "how many features
have zero importance?" the Scientist runs the analysis right then.
Best when per-step quality matters more than raw experiment volume.
Evolutionary ( socratic-evolve/ )
An evolutionary code-generation system (MLevolve) maintaining a tree
of candidate solutions across parallel branches. Includes evolution
stages (paradigm-shift mutations), fusion stages (cross-branch
solution merging), and runs multiple branches in parallel. During
review, the Scientist can only revise plan text (no tool access).
Best when the search space rewards high iteration volume.
All controlled via configuration flags
( use_socrates_review and use_baseline_pi in
config.yaml / config.py ):
We evaluate on five tasks from MLE-bench :
Follow the MLE-bench instructions to download the five competition
datasets. Place each one under a local directory and remember its path
— you'll plug it into the config in the next step.
2. Run the sequential scaffold
cd discover/
# Edit test_config.yaml to set:
# AGENT_CONFIG.exp_id -> the MLE-bench task id (e.g. "statoil-iceberg-classifier-challenge")
# AGENT_CONFIG.dataset_dir -> the local path you put the data in
# AGENT_CONFIG.model -> the LLM (default: claude-opus-4-6)
python test_agent_locally.py
This writes per-experiment folders, a best_score.txt , and a
submission.csv in dataset_dir . Submit submission.csv to Kaggle
to get the test score.
3. Run the evolutionary scaffold
cd socratic-evolve/public-repo/
python run.py \
exp_id= " statoil-iceberg-classifier-challenge " \
agent.use_socrates_review=True \
agent.steps=50
For each task, run it once per condition (toggling the flags above)
so you can compare Scientist-only / Baseline PI / Socrates side by
side.
cd socratic-evolve/public-repo/
python collect_and_plot.py # aggregates per-experiment logs into the paper's tables/plots
python dashboard.py # optional live dashboard
Expected results
Task
Scientist-only (test)
Baseline PI (test)
Socrates (test)
Δ vs Scientist
Statoil
0.255
0.251
0.229
+10.5%
COVID
0.389
0.308
0.294
+24.4%
Ventilator
1.534
0.815
0.853
+44.4%
NFL
0.198
0.537
0.584
+195.4%
Smartphone
6.285
5.993
5.984
+4.8%
Note: LLM agents are high-variance run-to-run. We saw a standard
deviation of ~15% of the mean across 10 Scientist-only seeds on
Smartphone. Expect single-seed numbers to vary; the direction of
the effect (Socrates ≥ Baseline PI > Scientist-only) is the
reproducible claim.
The key flags live in
socratic-evolve/public-repo/config/config.yaml and
discover/test_config.yaml :
A more detailed flag-level reference for the prompts (which blocks
get injected when) is in socratic-evolve/public-repo/agents/socrates/ .
# Sequential scaffold smoke test (no real run; mocks the LLM):
python discover/test_agent_locally.py --dry-run
# Evolutionary scaffold live test (requires API key):
cd socratic-evolve/public-repo/
pytest tests/test_socrates_live.py -k " test_socrates_basic "
Citation
@inproceedings { vrabac2026socrates ,
title = { Socrates: Structured Questioning Unlocks Latent Knowledge in AI Research Agents } ,
author = { Vrabac, Damir and Hebbar, Prannay and Manawat, Yogendra and Palanimalai, Selvam and Verboomen, Samuel and Juneja, Gurusha and Bhatia, Kunal and Baskaran, Vignesh } ,
booktitle = { Conference on Language Modeling (COLM) } ,
year = { 2026 }
}
License
A multi-agent protocol pairing a tool-using Scientist with a question-only advisor — no tools, no answers, no directives — improves Kaggle test performance on 4 of 5 MLE-bench tasks
hexo-ai.github.io/socrates/
Resources
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
